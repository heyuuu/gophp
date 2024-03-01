package main

import (
	"errors"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/tests"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getFormValueInt(request *http.Request, name string, defaultValue int) int {
	value := request.FormValue(name)
	if value == "" {
		return defaultValue
	}
	if intVal, err := strconv.Atoi(value); err == nil {
		return intVal
	} else {
		return defaultValue
	}
}

func apiTestListHandler(request *http.Request) (data any, err error) {
	path := request.FormValue("path")
	if path == "" {
		return nil, errors.New("path is empty")
	}
	testNames, err := loadTests(path)
	if err != nil {
		return nil, err
	}

	offset := getFormValueInt(request, "offset", 0)
	limit := getFormValueInt(request, "limit", -1)

	//
	offset = lang.FixRange(offset, 0, len(testNames))
	if limit < 0 || offset+limit >= len(testNames) {
		testNames = testNames[offset:]
	} else {
		testNames = testNames[offset : offset+limit]
	}

	return testNames, nil
}

func loadTests(path string) ([]string, error) {
	realpath := realTestCasePath(path)
	if _, err := os.Stat(realpath); err != nil {
		return nil, err
	}
	testFiles, err := tests.FindTestFiles(realpath)
	if err != nil {
		return nil, err
	}
	for i, testFile := range testFiles {
		testFiles[i] = relativeTestCasePath(testFile)
	}
	return testFiles, nil
}

func apiTestRunHandler(request *http.Request) (data any, err error) {
	name := request.FormValue("name")
	if name == "" {
		return nil, errors.New("name is empty")
	}
	testResult, log := apiRunTestCase(name)

	data = map[string]any{
		// case
		"case": testResult.Case(),
		// result
		"status":  testResult.MainType(),
		"output":  testResult.Output(),
		"reason":  testResult.Info() + "\n" + log,
		"useTime": testResult.UseTime().Nanoseconds(),
	}
	return data, nil
}

func apiRunTestCase(name string) (*tests.Result, string) {
	testFile := realTestCasePath(name)

	conf := tests.DefaultConfig()
	conf.SrcDir = "/Users/heyu/Code/src/php-7.4.33"
	conf.ExtDir = "/__ext__"
	conf.PhpBin = "/opt/homebrew/Cellar/php@7.4/7.4.33_6/bin/php"
	conf.PhpCgiBin = "/opt/homebrew/Cellar/php@7.4/7.4.33_6/bin/php-cgi"
	conf.Verbose = true

	var buf strings.Builder
	conf.Logger = tests.LoggerFunc(func(tc *tests.TestCase, event int, message string) {
		if tc != nil {
			buf.WriteString(message)
		}
	})

	return tests.TestOneFile(&conf, testFile), buf.String()
}

var testCasesPath string

func init() {
	pwd, _ := os.Getwd()
	testCasesPath = filepath.Join(pwd, "testcases") + string(filepath.Separator)
}

func realTestCasePath(path string) string {
	return filepath.Join(testCasesPath, path)
}
func relativeTestCasePath(path string) string {
	if strings.HasPrefix(path, testCasesPath) {
		return path[len(testCasesPath):]
	}
	return path
}
