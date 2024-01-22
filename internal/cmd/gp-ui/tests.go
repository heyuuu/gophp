package main

import (
	"errors"
	"github.com/heyuuu/gophp/tests"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func apiTestListHandler(request *http.Request) (data any, err error) {
	path := request.FormValue("path")
	if path == "" {
		return nil, errors.New("path is empty")
	}
	testNames, err := loadTests(path)
	if err != nil {
		return nil, err
	}
	return testNames, nil
}

func loadTests(path string) ([]string, error) {
	realpath := realTestCasePath(path)
	if _, err := os.Stat(realpath); err != nil {
		return nil, err
	}
	testFiles, err := tests.FindTestFiles(realpath, true)
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
	testFile := realTestCasePath(name)
	testResult := tests.RunTestFile(0, name, testFile)

	data = map[string]any{
		// case
		"case": testResult.Case,
		// result
		"status":  testResult.Type,
		"reason":  testResult.Reason,
		"useTime": testResult.UseTime.Nanoseconds(),
	}
	return data, nil
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
