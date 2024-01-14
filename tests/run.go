package tests

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var rawSupportSections = map[string]bool{
	"TEST":   true,
	"FILE":   true,
	"EXPECT": true,

	//"EXPECTF":              true,
	//"EXPECTREGEX":          true,
	//"EXPECTHEADERS":        true,
	//"POST":                 true,
	//"POST_RAW":             true,
	//"GZIP_POST":            true,
	//"DEFLATE_POST":         true,
	//"PUT":                  true,
	//"GET":                  true,
	//"COOKIE":               true,
	//"ARGS":                 true,
	//"REDIRECTTEST":         true,
	//"CAPTURE_STDIO":        true,
	//"STDIN":                true,
	//"CGI":                  true,
	//"PHPDBG":               true,
	//"INI":                  true,
	//"ENV":                  true,
	//"EXTENSIONS":           true,
	//"SKIPIF":               true,
	//"XFAIL":                true,
	//"XLEAK":                true,
	//"CLEAN":                true,
	//"CREDITS":              true,
	//"DESCRIPTION":          true,
	//"CONFLICTS":            true,
	//"WHITESPACE_SENSITIVE": true,
}

func (c Config) runTestRealRaw(testIndex int, tc *TestCase) (*TestResult, error) {
	if !c.rawSupport(testIndex, tc) {
		return c.runTestReal(testIndex, tc)
	}

	php := c.PhpBin
	sections := tc.sections

	var captureStdIn, captureStdOut, captureStdErr bool
	_ = captureStdIn
	if captureText, ok := sections["CAPTURE_STDIO"]; ok {
		ucCaptureText := strings.ToUpper(captureText)
		captureStdIn = strings.Contains(ucCaptureText, "STDIN")
		captureStdOut = strings.Contains(ucCaptureText, "STDOUT")
		captureStdErr = strings.Contains(ucCaptureText, "STDERR")
	} else {
		captureStdIn = true
		captureStdOut = true
		captureStdErr = true
	}
	var cmdRedirect string = ""
	if captureStdOut && captureStdErr {
		cmdRedirect = "2>&1"
	}

	/* For GET/POST/PUT tests, check if cgi sapi is available and if it is, use it. */
	useCgi := existKey(sections, "CGI") || sections["GET"] != "" || sections["POST"] != "" ||
		sections["GZIP_POST"] != "" || sections["DEFLATE_POST"] != "" || sections["POST_RAW"] != "" ||
		sections["PUT"] != "" || sections["COOKIE"] != "" || sections["EXPECTHEADERS"] != ""
	if useCgi {
		if c.PhpCgiBin != "" {
			php = c.PhpCgiBin + " -C "
		} else {
			return NewTestResult(tc, SKIP, "reason: CGI not available", 0), nil
		}
	}

	/* For phpdbg tests, check if phpdbg sapi is available and if it is, use it. */
	extraOptions := ""
	_ = extraOptions
	if existKey(sections, "PHPDBG") {
		if !existKey(sections, "STDIN") {
			sections["STDIN"] = sections["PHPDBG"] + "\n"
		}
		if c.PhpCgiBin != "" {
			php = c.PhpCgiBin + " -qIb"
			// Additional phpdbg command line options for sections that need to
			// be run straight away. For example, EXTENSIONS, SKIPIF, CLEAN.
			extraOptions = "-rr"
		} else {
			return NewTestResult(tc, SKIP, "reason: phpdbg not available", 0), nil
		}
	}

	testDir := filepath.Dir(tc.File)
	tempDir := c.mapTempPath(testDir)
	mainFileName := basename(tc.File, ".phpt")

	diffFilename := filepath.Join(tempDir, mainFileName+".diff")
	logFilename := filepath.Join(tempDir, mainFileName+".log")
	expFilename := filepath.Join(tempDir, mainFileName+".exp")
	outputFilename := filepath.Join(tempDir, mainFileName+".out")
	memcheckFilename := filepath.Join(tempDir, mainFileName+".mem")
	shFilename := filepath.Join(tempDir, mainFileName+".sh")
	tempFile := filepath.Join(tempDir, mainFileName+".php")
	testFile := filepath.Join(testDir, mainFileName+".php")
	tempSkipif := filepath.Join(tempDir, mainFileName+".skip.php")
	testSkipif := filepath.Join(testDir, mainFileName+".skip.php")
	tempClean := filepath.Join(tempDir, mainFileName+".clean.php")
	testClean := filepath.Join(testDir, mainFileName+".clean.php")
	preloadFilename := filepath.Join(tempDir, mainFileName+".preload.php")
	tmpPost := filepath.Join(tempDir, mainFileName+".post")

	if c.TempSource != "" && c.TempTarget != "" {
		tempSkipif += "s"
		tempFile += "s"
		tempClean += "s"

		copyFile := filepath.Join(tempDir, basename(tc.File, "")+".phps")
		if fileText, ok := sections["FILE"]; ok {
			c.saveText(copyFile, fileText)
		}
	}

	// unlink old test results
	unlink(diffFilename)
	unlink(logFilename)
	unlink(expFilename)
	unlink(outputFilename)
	unlink(memcheckFilename)
	unlink(shFilename)
	unlink(tempFile)
	unlink(testFile)
	unlink(tempSkipif)
	unlink(testSkipif)
	unlink(tempClean)
	unlink(testClean)
	unlink(preloadFilename)
	unlink(tmpPost)

	// Reset environment from any previous test.
	var env *Environments = new(Environments)
	env.Set("REDIRECT_STATUS", "")
	env.Set("QUERY_STRING", "")
	env.Set("PATH_TRANSLATED", "")
	env.Set("SCRIPT_FILENAME", "")
	env.Set("REQUEST_METHOD", "")
	env.Set("CONTENT_TYPE", "")
	env.Set("CONTENT_LENGTH", "")
	env.Set("TZ", "")

	if sections["ENV"] != "" {
		for _, e := range strings.Split(strings.TrimSpace(sections["ENV"]), "\n") {
			if k, v, ok := strings.Cut(strings.TrimSpace(e), "="); ok && k != "" {
				env.Set(k, v)
			}
		}
	}

	// Default ini settings
	var iniSettings []IniEntry = c.setting2array(baseIniOverwrites)

	// Additional required extensions
	if existKey(sections, "EXTENSIONS") {
		//extDir := c.ExtensionDir
	}

	// additional ini overwrites
	iniSettings = append(iniSettings, c.IniOverwrites...)
	//oriIniSettingsParam := c.setting2params(iniSettings)

	// Any special ini settings
	// these may overwrite the test defaults...
	if iniText := sections["INI"]; iniText != "" {
		iniText = strings.ReplaceAll(iniText, "{PWD}", filepath.Dir(tc.File))
		iniText = strings.ReplaceAll(iniText, "{TMP}", os.TempDir())
		iniSettings = append(iniSettings, c.setting2array(regexp.MustCompile(`[\n\r]+`).Split(iniText, -1))...)
	}
	iniSettingsParam := c.setting2params(iniSettings)

	// Check if test should be skipped.
	//info := ""
	//warn := false

	if skipIfText := sections["SKIPIF"]; strings.TrimSpace(skipIfText) != "" {
		c.showFileBlock("skip", sections["SKIPIF"], "")
		c.saveTextEx(testSkipif, skipIfText, tempSkipif)
		// todo
	}

	// We've satisfied the preconditions - run the test!
	if fileText, ok := sections["FILE"]; ok {
		c.showFileBlock("php", fileText, "TEST")
		c.saveTextEx(testFile, fileText, tempFile)
	} else {
		testFile = ""
		tempFile = ""
	}

	queryString := strings.TrimSpace(sections["GET"])

	env.Set("REDIRECT_STATUS", "1")
	if env.Get("QUERY_STRING") == "" {
		env.Set("QUERY_STRING", queryString)
	}
	if env.Get("PATH_TRANSLATED") == "" {
		env.Set("PATH_TRANSLATED", testFile)
	}
	if env.Get("SCRIPT_FILENAME") == "" {
		env.Set("SCRIPT_FILENAME", testFile)
	}
	env.Set("HTTP_COOKIE", strings.TrimSpace(sections["COOKIE"]))

	args := ""
	if existKey(sections, "ARGS") {
		args = " -- " + sections["ARGS"]
	}

	var cmd string
	if postRaw := sections["POST_RAW"]; postRaw != "" {
		post := strings.TrimSpace(postRaw)
		var requestLines []string
		for _, line := range strings.Split(post, "\n") {
			if env.Get("CONTENT_TYPE") == "" {
				if match := regexp.MustCompile(`^(?i)Content-Type:(.*)`).FindStringSubmatch(line); len(match) > 0 {
					env.Set("CONTENT_TYPE", strings.TrimSpace(strings.ReplaceAll(match[1], "\r", "")))
					continue
				}
			}
			requestLines = append(requestLines, line)
		}
		request := strings.Join(requestLines, "\n")

		env.Set("CONTENT_LENGTH", strconv.Itoa(len(request)))
		env.Set("REQUEST_METHOD", "POST")

		if request == "" {
			return NewTestResult(tc, BORK, "reason: POST_RAW build request empty", 0), nil
		}
		c.saveText(tmpPost, request)
		cmd = fmt.Sprintf(`%s %s %s -f "%s" %s < "%s"`, php, c.PassOption, iniSettingsParam, testFile, cmdRedirect, tmpPost)
	} else if put := sections["PUT"]; put != "" {
		post := strings.TrimSpace(put)
		var requestLines []string
		for _, line := range strings.Split(post, "\n") {
			if env.Get("CONTENT_TYPE") == "" {
				if match := regexp.MustCompile(`^(?i)Content-Type:(.*)`).FindStringSubmatch(line); len(match) > 0 {
					env.Set("CONTENT_TYPE", strings.TrimSpace(strings.ReplaceAll(match[1], "\r", "")))
					continue
				}
			}
			requestLines = append(requestLines, line)
		}
		request := strings.Join(requestLines, "\n")

		env.Set("CONTENT_LENGTH", strconv.Itoa(len(request)))
		env.Set("REQUEST_METHOD", "PUT")

		if request == "" {
			return NewTestResult(tc, BORK, "reason: PUT build request empty", 0), nil
		}
		c.saveText(tmpPost, request)
		cmd = fmt.Sprintf(`%s %s %s -f "%s" %s < "%s"`, php, c.PassOption, iniSettingsParam, testFile, cmdRedirect, tmpPost)
	} else if postText := sections["POST"]; postText != "" {
		post := strings.TrimSpace(postText)
		c.saveText(tmpPost, post)

		env.Set("REQUEST_METHOD", "POST")
		if env.Get("CONTENT_TYPE") == "" {
			env.Set("CONTENT_TYPE", "application/x-www-form-urlencoded")
		}
		if env.Get("CONTENT_LENGTH") == "" {
			env.Set("CONTENT_LENGTH", strconv.Itoa(len(post)))
		}
		cmd = fmt.Sprintf(`%s %s %s -f "%s" %s < "%s"`, php, c.PassOption, iniSettingsParam, testFile, cmdRedirect, tmpPost)
	} else if gzipPostText := sections["GZIP_POST"]; gzipPostText != "" {
		//todo
		cmd = fmt.Sprintf(`%s %s %s -f "%s" %s < "%s"`, php, c.PassOption, iniSettingsParam, testFile, cmdRedirect, tmpPost)
	} else if gzipPostText := sections["DEFLATE_POST"]; gzipPostText != "" {
		//todo
		cmd = fmt.Sprintf(`%s %s %s -f "%s" %s < "%s"`, php, c.PassOption, iniSettingsParam, testFile, cmdRedirect, tmpPost)
	} else {
		env.Set("REQUEST_METHOD", "GET")
		env.Set("CONTENT_TYPE", "")
		env.Set("CONTENT_LENGTH", "")

		cmd = fmt.Sprintf(`%s %s %s -f "%s" %s%s`, php, c.PassOption, iniSettingsParam, testFile, args, cmdRedirect)
	}

	c.Events.Log(testIndex, "CONTENT_LENGTH  = "+env.Get("CONTENT_LENGTH"))
	c.Events.Log(testIndex, "CONTENT_TYPE    = "+env.Get("CONTENT_TYPE"))
	c.Events.Log(testIndex, "PATH_TRANSLATED = "+env.Get("PATH_TRANSLATED"))
	c.Events.Log(testIndex, "QUERY_STRING    = "+env.Get("QUERY_STRING"))
	c.Events.Log(testIndex, "REDIRECT_STATUS = "+env.Get("REDIRECT_STATUS"))
	c.Events.Log(testIndex, "REQUEST_METHOD  = "+env.Get("REQUEST_METHOD"))
	c.Events.Log(testIndex, "SCRIPT_FILENAME = "+env.Get("SCRIPT_FILENAME"))
	c.Events.Log(testIndex, "HTTP_COOKIE     = "+env.Get("HTTP_COOKIE"))
	c.Events.Log(testIndex, "COMMAND "+cmd)

	return NewTestResult(tc, SKIP, "", 0), nil
}

func (c Config) rawSupport(testIndex int, tc *TestCase) bool {
	// check sections
	for section, _ := range tc.sections {
		if !rawSupportSections[section] {
			//c.Events.Log(testIndex, fmt.Sprintf("unsupport section: %s", section))
			return false
		}
	}
	return true
}

func (c Config) mapTempPath(srcPath string) string {
	// todo
	return srcPath
}

func (c Config) showFileBlock(typ string, block string, section string) {
	// todo
}

func (c Config) saveText(file string, text string) {
	err := filePutContents(file, text)
	if err != nil {
		c.errorf(`Cannot open file '%s' (save_text)`, file)
	}
}

func (c Config) saveTextEx(file string, text string, fileCopy string) {
	if fileCopy != "" && fileCopy != file {
		c.saveText(fileCopy, text)
	}

	c.saveText(file, text)
}

func (c Config) setting2array(settings []string) []IniEntry {
	var result []IniEntry
	for _, ini := range settings {
		if iniName, iniValue, ok := strings.Cut(ini, "="); ok {
			result = append(result, IniEntry{
				name:  strings.TrimSpace(iniName),
				value: strings.TrimSpace(iniValue),
			})
		}
	}
	return result
}

func (c Config) setting2params(settings []IniEntry) string {
	// todo 移除 extension、zend_extension 以外的重复项

	var buf strings.Builder
	for _, entry := range settings {
		buf.WriteString(fmt.Sprintf(` -d "%s=%s"`, entry.name, entry.value))
	}
	return buf.String()
}

func (c Config) errorf(format string, v ...any) {
	log.Panicf(format, v...)
}

func (c Config) systemWithTimeout(cmd []string, env *Environments) string {
	return ""
}
