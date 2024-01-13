package tests

import (
	"log"
	"path/filepath"
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
	if captureText, ok := sections["CAPTURE_STDIO"]; ok {
		ucCaptureText := strings.ToUpper(captureText)
		captureStdIn = strings.Contains(ucCaptureText, "STDIN")
		captureStdOut = strings.Contains(ucCaptureText, "STDOUT")
		captureStdErr = strings.Contains(ucCaptureText, "STDERR")
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

	// We've satisfied the preconditions - run the test!
	if fileText, ok := sections["FILE"]; ok {
		c.showFileBlock("php", fileText, "TEST")
		c.saveTextEx(testFile, fileText, tempFile)
	} else {
		testFile = ""
		tempFile = ""
	}

	if true {
		cmd = append(cmd, "-f", testFile)
	}

	c.Events.Log(testIndex, "CONTENT_LENGTH  =")
	c.Events.Log(testIndex, "CONTENT_TYPE    =")
	c.Events.Log(testIndex, "PATH_TRANSLATED =")
	c.Events.Log(testIndex, "QUERY_STRING    =")
	c.Events.Log(testIndex, "REDIRECT_STATUS =")
	c.Events.Log(testIndex, "REQUEST_METHOD  =")
	c.Events.Log(testIndex, "SCRIPT_FILENAME =")
	c.Events.Log(testIndex, "HTTP_COOKIE     =")
	c.Events.Log(testIndex, "COMMAND "+strings.Join(cmd, " "))

	return nil, nil
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

func (c Config) errorf(format string, v ...any) {
	log.Panicf(format, v...)
}
