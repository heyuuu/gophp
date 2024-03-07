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

// -- START 临时写死的路径，后续需移除 --
const SrcDir = "/Users/heyu/Code/src/php-7.4.33"
const ExtDir = "/__ext__"
const PhpBin = "/opt/homebrew/Cellar/php@7.4/7.4.33_6/bin/php"
const PhpCgiBin = "/opt/homebrew/Cellar/php@7.4/7.4.33_6/bin/php-cgi"

// -- END 临时写死的路径，后续需移除 --

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

	testCases, err := tests.FindTestCases(SrcDir, realpath)
	if err != nil {
		return nil, err
	}

	var testNames = make([]string, len(testCases))
	for i, tc := range testCases {
		testNames[i] = tc.TestName()
	}
	return testNames, nil
}

func apiTestRunHandler(request *http.Request) (data any, err error) {
	name := request.FormValue("name")
	if name == "" {
		return nil, errors.New("name is empty")
	}
	testResult, log := apiRunTestCase(name)

	sections := testResult.Case().Sections()

	data = map[string]any{
		// case
		"code":   sections["FILE"],
		"expect": sections["EXPECT"] + sections["EXPECTF"] + sections["EXPECTREGEX"],

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
	conf.SrcDir = SrcDir
	conf.ExtDir = ExtDir
	conf.PhpBin = PhpBin
	conf.PhpCgiBin = PhpCgiBin
	conf.Verbose = true

	var buf strings.Builder
	conf.Logger = tests.LoggerFunc(func(tc *tests.TestCase, event int, message string) {
		if tc != nil {
			buf.WriteString(message)
		}
	})

	return tests.TestOneFile(conf, testFile), buf.String()
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
