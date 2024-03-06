package tests

import (
	"errors"
	"fmt"
	"github.com/heyuuu/gophp/kits/slicekit"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func TestAll(conf *Config) error {
	return newRunner(conf).TestAll()
}
func TestOneFile(conf *Config, testFile string) *Result {
	return newRunner(conf).TestOnce(testFile)
}

type runner struct {
	conf        *Config
	testFiles   []string
	logger      Logger
	summary     *Summary
	passOption  string
	passOptions []commandArg
	env         *Env
	oneFileMode bool
}

func newRunner(conf *Config) *runner {
	r := &runner{conf: conf}
	r.init(conf)
	return r
}

func (r *runner) TestAll() (err error) {
	r.testFiles, err = FindTestFilesInSrcDir(r.conf.SrcDir, true)
	if err != nil {
		return err
	}
	if limit := r.conf.Limit; limit > 0 && len(r.testFiles) > limit {
		r.testFiles = r.testFiles[:limit]
	}
	r.runTests()
	return nil
}

func (r *runner) TestOnce(testFile string) *Result {
	r.testFiles = []string{testFile}
	r.oneFileMode = true
	lastResult := r.runTests()
	return lastResult
}

func (r *runner) init(conf *Config) {
	r.conf = conf
	if conf.Logger != nil {
		r.logger = conf.Logger
	} else {
		r.logger = ConsoleLogger
	}
	// 并发时，使用并发安全的 Logger
	if conf.Workers > 1 {
		r.logger = NewSyncLogger(r.logger)
	}
	r.summary = NewSummary()

	// init env
	r.env = NewEnv() // $_ENV ?: getenv()
	r.env.Set("TEMP", os.TempDir())

	// init conf
	r.passOptions = nil
	if conf.PassOptionN {
		r.passOptions = append(r.passOptions, arg("-n"))
	}
	if conf.PassOptionE {
		r.passOptions = append(r.passOptions, arg("-e"))
	}
	if conf.ConfPassed != "" {
		r.passOptions = append(r.passOptions, arg("-c"), quoteArg(conf.ConfPassed))
	}
	if conf.Quite {
		r.env.Set("NO_INTERACTION", "1")
	}
	if conf.TestTimeout != "" {
		r.env.Set("TEST_TIMEOUT", conf.TestTimeout)
	}
	if conf.X {
		r.env.Set("SKIP_SLOW_TESTS", "1")
	}
	if conf.Offline {
		r.env.Set("SKIP_ONLINE_TESTS", "1")
	}

	// verify config
	if conf.PhpBin == "" || !fileExists(conf.PhpBin) {
		panic(errors.New("Config.PhpBin must be set to specify PHP executable"))
	}
}

func (r *runner) runTests() (lastResult *Result) {
	// write information
	// todo

	// run all tests
	r.onAllStart()
	if r.conf.Workers > 1 {
		lastResult = r.parallelRunTests()
	} else {
		lastResult = r.simpleRunTests()
	}
	r.onAllEnd()

	return
}

func (r *runner) simpleRunTests() (lastResult *Result) {
	for i, file := range r.testFiles {
		testIndex := i + 1
		lastResult = r.runTest(testIndex, file)
		r.summary.AddResult(lastResult)
	}
	return
}

func (r *runner) parallelRunTests() (lastResult *Result) {
	testFiles := r.testFiles
	testCount := len(testFiles)
	workers := r.conf.Workers

	var wg sync.WaitGroup
	wg.Add(testCount)

	// 结果chan，串行处理结果
	var resultChan = make(chan *Result, workers)
	defer close(resultChan)

	// 单独 goroutine 串行处理结果(因为 Summary 非并发安全)
	go func() {
		for i := 0; i < testCount; i++ {
			lastResult = <-resultChan
			if lastResult != nil {
				r.summary.AddResult(lastResult)
			}
		}
	}()

	// 用于限制并发数
	var limitChan = make(chan struct{}, workers)
	defer close(limitChan)

	// 并发遍历任务
	for i, testFile := range testFiles {
		limitChan <- struct{}{}
		go func(index int, file string) {
			var result *Result
			defer func() {
				if e := recover(); e != nil {
					log.Printf("parallelRunTests painc: %v\n", e)
				}
				wg.Done()

				resultChan <- result
				<-limitChan
			}()

			result = r.runTest(index, file)
		}(i+1, testFile)
	}

	wg.Wait()
	return
}

func (r *runner) runTest(testIndex int, testFile string) *Result {
	shortFileName := testFile
	if strings.HasPrefix(testFile, r.conf.SrcDir+"/") {
		shortFileName = testFile[len(r.conf.SrcDir)+1:]
	}

	tc := NewTestCase(testIndex, testFile, shortFileName)

	r.onTestStart(tc)
	r.cleanTempFiles(tc, true)
	result := r.runTestReal(tc)
	//r.cleanTempFiles(tc, false)
	if r.conf.SlowMinTime > 0 && r.conf.SlowMinTime < result.useTime {
		result.slow = true
	}
	r.onTestEnd(result)

	return result
}

func (r *runner) runTestReal(tc *TestCase) *Result {
	// Load the sections of the test file.
	err := tc.parse()
	if err != nil {
		return SimpleResult(tc, BORK, err.Error())
	}
	sections := tc.sections

	r.logger.Log(tc, fmt.Sprintf("TEST %d/%d [%s]\n", tc.index, len(r.testFiles), tc.shortFileName))

	// stdio
	var captureStdIn, captureStdOut, captureStdErr bool
	if existKey(sections, "CAPTURE_STDIO") {
		captureStdioText := strings.ToUpper(sections["CAPTURE_STDIO"])
		captureStdIn = strings.Contains(captureStdioText, "STDIN")
		captureStdOut = strings.Contains(captureStdioText, "STDOUT")
		captureStdErr = strings.Contains(captureStdioText, "STDERR")
	} else {
		captureStdIn, captureStdOut, captureStdErr = true, true, true
	}

	/* For GET/POST/PUT tests, check if cgi sapi is available and if it is, use it. */
	useCgi := existAnyKey(sections, "CGI", "GET", "POST", "GZIP_POST", "DEFLATE_POST", "POST_RAW", "PUT", "COOKIE", "EXPECTHEADERS")

	var baseCmd *command
	if useCgi {
		if r.conf.PhpCgiBin == "" {
			return SimpleResult(tc, SKIP, "reason: CGI not available")
		}

		baseCmd = newCommand(r.conf.PhpCgiBin, arg("-C"))
	} else {
		baseCmd = newCommand(r.conf.PhpBin)
	}

	// Reset environment from any previous test.
	env := r.env.Clone()
	env.Set("REDIRECT_STATUS", "")
	env.Set("QUERY_STRING", "")
	env.Set("PATH_TRANSLATED", "")
	env.Set("SCRIPT_FILENAME", "")
	env.Set("REQUEST_METHOD", "")
	env.Set("CONTENT_TYPE", "")
	env.Set("CONTENT_LENGTH", "")
	env.Set("TZ", "")

	if envText := sections["ENV"]; envText != "" {
		for _, envLine := range strings.Split(envText, "\n") {
			if envKey, envValue, ok := strings.Cut(envLine, "="); ok && envKey != "" {
				env.Set(envKey, envValue)
			}
		}
	}

	// ini settings
	var iniSettings = NewIniSettings()
	if extText := strings.TrimSpace(sections["EXTENSIONS"]); extText != "" {
		for _, ext := range strings.Split(extText, "\n") {
			ext = strings.TrimSpace(ext)
			if ext == "opcache" {
				iniSettings.Add("zend_extension", filepath.Join(r.conf.ExtDir, ext+".so"))
			} else {
				iniSettings.Add("extension", filepath.Join(r.conf.ExtDir, ext+".so"))
			}
		}
	}
	iniSettings.Merge(baseIniOverwrites)
	iniSettings.Merge(r.conf.IniOverwrites)
	origIniSettingsParam := iniSettings.ToArgs()

	// Any special ini settings
	// these may overwrite the test defaults...
	if iniText := sections["INI"]; iniText != "" {
		iniText = strings.NewReplacer(
			"{PWD}", filepath.Dir(tc.file),
			"{TMP}", os.TempDir(),
		).Replace(iniText)
		iniSettings.Merge(strings.Split(iniText, "\n"))
	}
	iniSettingsParam := iniSettings.ToArgs()

	extraArgs := slicekit.Concat(r.passOptions, iniSettingsParam)
	env.Set("TEST_PHP_EXTRA_ARGS", CommandArgsToString(extraArgs))

	// Check if test should be skipped.
	if skipifText := sections["SKIPIF"]; strings.TrimSpace(skipifText) != "" {
		r.showFileBlock(tc, blockSkip, skipifText)
		r.saveText(tc, tc.testSkipif, skipifText)

		//extra := "unset REQUEST_METHOD; unset QUERY_STRING; unset PATH_TRANSLATED; unset SCRIPT_FILENAME; unset REQUEST_METHOD;"
		skipifCmd := baseCmd.clone().
			add(r.passOptions...).
			add(arg("-q")).
			add(origIniSettingsParam...).
			add(noFileCacheArgs...).
			option("-d", "display_errors=0").
			add(quoteArg(tc.testSkipif))
		output := r.runCommand(skipifCmd)
		if len(output) >= 4 && strings.ToLower(output[:4]) == "skip" {
			reason := ""
			if match := regexp.MustCompile(`^\s*skip\s*(.+)\s*`).FindStringSubmatch(output); match != nil {
				reason = "reason: " + match[1]
			}
			return SimpleResult(tc, SKIP, reason)
		}
		// todo
	}

	if existAnyKey(sections, "GZIP_POST", "DEFLATE_POST") {
		return SimpleResult(tc, SKIP, "reason: ext/zlib required")
	}

	// We've satisfied the preconditions - run the test!
	var testFile, queryString string
	if existKey(sections, "FILE") {
		testFile = tc.testFile
		r.showFileBlock(tc, blockTest, sections["FILE"])
		r.saveText(tc, testFile, sections["FILE"])
	}
	queryString = strings.TrimSpace(sections["GET"])

	// env
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

	// args
	var args string
	if existKey(sections, "ARGS") {
		args = " -- " + sections["ARGS"]
	}

	// body
	execCmd := baseCmd.clone()
	execCmd.add(r.passOptions...).addIniSettings(iniSettings).add(arg("-f"), quoteArg(testFile))
	execCmd.capture(captureStdIn, captureStdOut, captureStdErr)
	if request, ok := r.tryBuildRequestContent(sections, env); ok {
		r.saveText(tc, tc.testPost, request)
		execCmd.option("<", tc.testPost)
	} else {
		execCmd.add(arg(args))
	}
	execCmd.stdin = sections["STDIN"]

	// show before test exec
	if r.conf.Verbose {
		r.logger.Log(tc, "\nCONTENT_LENGTH  = "+env.Get("CONTENT_LENGTH"))
		r.logger.Log(tc, "\nCONTENT_TYPE    = "+env.Get("CONTENT_TYPE"))
		r.logger.Log(tc, "\nPATH_TRANSLATED = "+env.Get("PATH_TRANSLATED"))
		r.logger.Log(tc, "\nQUERY_STRING    = "+env.Get("QUERY_STRING"))
		r.logger.Log(tc, "\nREDIRECT_STATUS = "+env.Get("REDIRECT_STATUS"))
		r.logger.Log(tc, "\nREQUEST_METHOD  = "+env.Get("REQUEST_METHOD"))
		r.logger.Log(tc, "\nSCRIPT_FILENAME = "+env.Get("SCRIPT_FILENAME"))
		r.logger.Log(tc, "\nHTTP_COOKIE     = "+env.Get("HTTP_COOKIE"))
		r.logger.Log(tc, "\nCOMMAND "+execCmd.String()+"\n")
	}

	startTime := time.Now()
	output := r.runCommand(execCmd)
	endTime := time.Now()
	useTime := endTime.Sub(startTime)

	if !r.conf.NoClean || r.conf.IsKeep("clean") {
		if cleanText := strings.TrimSpace(sections["CLEAN"]); cleanText != "" {
			r.showFileBlock(tc, blockClean, cleanText)
			r.saveText(tc, tc.testClean, cleanText)

			if !r.conf.NoClean {
				//extra := "unset REQUEST_METHOD; unset QUERY_STRING; unset PATH_TRANSLATED; unset SCRIPT_FILENAME; unset REQUEST_METHOD;"
				cleanCmd := baseCmd.clone().add(r.passOptions...).
					add(arg("-q")).
					add(origIniSettingsParam...).
					add(noFileCacheArgs...).
					add(quoteArg(tc.testClean))
				r.runCommand(cleanCmd)
			}
		}
	}

	// Does the output match what is expected?
	output = strings.ReplaceAll(strings.TrimSpace(output), "\r\n", "\n")

	/* when using CGI, strip the headers from the output */
	var headers = map[string]string{}
	if useCgi {
		if match := regexp.MustCompile(`^(.*?)\r?\n\r?\n(.*)`).FindStringSubmatch(output); match != nil {
			output = strings.TrimSpace(match[2])
			rh := strings.Split(match[1], "\n")

			for _, line := range rh {
				if name, value, ok := strings.Cut(line, ":"); ok {
					headers[name] = value
				}
			}
		}
	}

	var failedHeaders = false
	var wantedHeaders, outputHeaders []string
	if existKey(sections, "EXPECTHEADERS") {
		var want = map[string]string{}
		for _, line := range strings.Split(sections["EXPECTHEADERS"], "\n") {
			if key, value, ok := strings.Cut(line, ":"); ok {
				key, value = strings.TrimSpace(key), strings.TrimSpace(value)
				want[key] = value
				// wanted header
				wantedHeaders = append(wantedHeaders, key+": "+value)
				// output header
				if existKey(headers, key) {
					outputHeaders = append(outputHeaders, key+": "+headers[key])
				}
				if !existKey(headers, key) || headers[key] != value {
					failedHeaders = true
				}
			}
		}
	}

	r.showFileBlock(tc, blockOut, output)

	var wanted, wantedReg, info string
	var passed, warn bool
	if existAnyKey(sections, "EXPECTF", "EXPECTREGEX") {
		if existKey(sections, "EXPECTF") {
			wanted = strings.TrimSpace(sections["EXPECTF"])
		} else {
			wanted = strings.TrimSpace(sections["EXPECTREGEX"])
		}

		r.showFileBlock(tc, blockExp, wanted)
		wantedReg = strings.ReplaceAll(wanted, "\r\n", "\n")

		if existKey(sections, "EXPECTF") {
			wantedReg = convertExpectFormat2Regex(wantedReg)
		}

		passed, err = safeExpectRegexCompare(wantedReg, output)
		if err != nil {
			return SimpleResult(tc, BORK, err.Error())
		}
	} else {
		wanted = strings.TrimSpace(sections["EXPECT"])
		wanted = strings.ReplaceAll(wanted, "\r\n", "\n")
		r.showFileBlock(tc, blockExp, wanted)

		passed = output == wanted
	}
	if passed && !failedHeaders {
		if existKey(sections, "XFAIL") {
			warn = true
			info = " (warn: XFAIL section but test passes)"
		} else {
			return PassResult(tc, output, useTime)
		}
	}

	// Test failed so we need to report details.
	if failedHeaders {
		passed = false
		wanted = strings.Join(wantedHeaders, "\n") + "\n--HEADERS--\n" + wanted
		output = strings.Join(outputHeaders, "\n") + "\n--HEADERS--\n" + output

		if wantedReg != "" {
			wantedReg = pregQuote(strings.Join(wantedHeaders, "\n")+"\n--HEADERS--\n") + wantedReg
		}
	}

	// result types
	var resultTypes []ResultType
	if warn {
		resultTypes = append(resultTypes, WARN)
	}
	if !passed {
		if existKey(sections, "XFAIL") {
			resultTypes = append(resultTypes, XFAIL)
			info = "  XFAIL REASON: " + strings.TrimSpace(sections["XFAIL"])
		} else {
			resultTypes = append(resultTypes, FAIL)
		}
	}
	if !passed {
		diff := generateDiff(wanted, wantedReg, output)
		r.showFileBlock(tc, blockDiff, diff)
	}

	return ComplexResult(tc, resultTypes, output, info, 0)
}

func (r *runner) showFileBlock(tc *TestCase, typ string, block string) {
	if r.conf.IsShow(typ) {
		r.logger.Log(tc, fmt.Sprintf("\n========%s========\n%s\n========DONE========\n", typ, strings.TrimSpace(block)))
	}
}

func (r *runner) cleanTempFiles(tc *TestCase, force bool) {
	if force || !r.conf.IsKeep("php") {
		_ = unlink(tc.testFile)
	}
	if force || !r.conf.IsKeep("skip") {
		_ = unlink(tc.testSkipif)
	}
	_ = unlink(tc.testClean)
	_ = unlink(tc.testPost)
}

func (r *runner) saveText(tc *TestCase, file string, content string) {
	err := filePutContents(file, content)
	if err != nil {
		panic(fmt.Errorf("cannot open file '%s' (save_text): %w", file, err))
	}
}

func (r *runner) runCommand(cmd *command) string {
	output, err := cmd.Run()
	if err != nil && output == "" {
		return output + "\nRun Error: " + err.Error()
	}
	return output
}

func (r *runner) tryBuildRequestContent(sectionText map[string]string, env *Env) (request string, ok bool) {
	if postRawText := sectionText["POST_RAW"]; postRawText != "" {
		request = r.buildBody(postRawText, env)

		env.Set("CONTENT_LENGTH", strconv.Itoa(len(request)))
		env.Set("REQUEST_METHOD", "POST")

		return request, true
	} else if putText := sectionText["PUT"]; putText != "" {
		request = r.buildBody(postRawText, env)

		env.Set("CONTENT_LENGTH", strconv.Itoa(len(request)))
		env.Set("REQUEST_METHOD", "PUT")

		return request, true
	} else if postText := sectionText["POST"]; postText != "" {
		request := strings.TrimSpace(postText)

		env.Set("REQUEST_METHOD", "POST")
		env.SetIfEmpty("CONTENT_TYPE", "application/x-www-form-urlencoded")
		env.SetIfEmpty("CONTENT_LENGTH", strconv.Itoa(len(request)))

		return request, true
	} else if gzipPostText := sectionText["GZIP_POST"]; gzipPostText != "" {
		request, _ := gzencode(strings.TrimSpace(gzipPostText))

		env.Set("REQUEST_METHOD", "POST")
		env.Set("CONTENT_TYPE", "application/x-www-form-urlencoded")
		env.Set("CONTENT_LENGTH", strconv.Itoa(len(request)))
		env.Set("HTTP_CONTENT_ENCODING", "gzip")

		return request, true
	} else if deflatePostText := sectionText["DEFLATE_POST"]; deflatePostText != "" {
		request, _ := gzcompress(strings.TrimSpace(deflatePostText))

		env.Set("REQUEST_METHOD", "POST")
		env.Set("CONTENT_TYPE", "application/x-www-form-urlencoded")
		env.Set("CONTENT_LENGTH", strconv.Itoa(len(request)))
		env.Set("HTTP_CONTENT_ENCODING", "deflate")

		return request, true
	} else {
		env.Set("REQUEST_METHOD", "GET")
		env.Set("CONTENT_TYPE", "")
		env.Set("CONTENT_LENGTH", "")

		return "", false
	}
}

func (r *runner) buildBody(text string, env *Env) string {
	rawLines := strings.Split(strings.TrimSpace(text), "\n")

	var buf strings.Builder
	buf.Grow(len(text))
	for _, line := range rawLines {
		line = strings.TrimSpace(line)
		if env.Get("CONTENT_TYPE") == "" && strings.HasPrefix(line, "Content-Type:") {
			env.Set("CONTENT_TYPE", strings.TrimSpace(line[len("Content-Type:"):]))
			continue
		}

		if buf.Len() > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(line)
	}

	return buf.String()
}

// -- events

func (r *runner) onAllStart() {
	startTime := time.Now()

	r.summary.StartTime = startTime

	r.logger.OnAllStart()
	r.logger.Log(nil, "=====================================================================\n")
	r.logger.Log(nil, "TIME START "+timeFormat(startTime, "Y-m-d H:i:s")+"\n")
	r.logger.Log(nil, "=====================================================================\n")
}

func (r *runner) onAllEnd() {
	endTime := time.Now()

	r.summary.EndTime = endTime

	r.logger.Log(nil, "=====================================================================\n")
	r.logger.Log(nil, "TIME END "+timeFormat(endTime, "Y-m-d H:i:s")+"\n")
	r.logger.Log(nil, "=====================================================================\n")
	if !r.oneFileMode {
		r.logger.Log(nil, r.summary.Summary())
	}
	r.logger.OnAllEnd()
}

func (r *runner) onTestStart(tc *TestCase) {
	r.logger.Log(nil, fmt.Sprintf("RUN: %d %s\n", tc.index, tc.shortFileName))

	r.logger.OnTestStart(tc)
	if r.conf.Verbose {
		r.logger.Log(tc, fmt.Sprintf("\n=================\nTEST %s\n", tc.file))
	}
}

func (r *runner) onTestEnd(result *Result) {
	tc := result.tc
	r.logger.Log(tc, fmt.Sprintf("%s %s %s\n", result.ShowTypeNames(), tc.ShowName(), result.info))
	r.logger.OnTestEnd(tc)
}
