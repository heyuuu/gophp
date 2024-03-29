package main

import (
	"github.com/heyuuu/gophp/tests"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var pwd string

func init() {
	pwd, _ = os.Getwd()
}

func runTestCaseInDir(t *testing.T, dir string) {
	files, _ := tests.FindTestFiles(dir, true)
	for _, file := range files {
		name := file
		if strings.HasPrefix(name, pwd) {
			name = name[len(pwd):]
		}
		t.Run(name, func(t *testing.T) {
			runTestCaseReal(t, name, file)
		})
	}
}

func runTestCase(t *testing.T, testName string) {
	testFile := filepath.Join(pwd, testName)
	runTestCaseReal(t, testName, testFile)
}

func runTestCaseReal(t *testing.T, testName string, testFile string) {
	defer func() {
		if e := recover(); e != nil {
			t.Errorf("runTestCase() panic = %v", e)
		}
	}()

	result := tests.RunTestFile(0, testName, testFile)
	switch result.Type {
	case tests.PASS:
		// pass
	case tests.SKIP:
		t.SkipNow()
	default:
		t.Errorf("runTestCase() fail = %s", result.Reason)
	}
}
