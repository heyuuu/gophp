package main

import (
	"errors"
	"github.com/heyuuu/gophp/tests"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

var pwd string

func init() {
	pwd, _ = os.Getwd()
}

func runTestCaseInDir(t *testing.T, dir string) {
	files := findFiles(dir, true)
	sort.Strings(files)
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

	tr, err := tests.RunCaseBuiltin(0, testName, testFile)
	if err != nil {
		if errors.Is(err, tests.ErrUnsupportedSection) {
			t.SkipNow()
			return
		}

		t.Errorf("runTestCase() error = %v", err)
		return
	}
	switch tr.Type {
	case tests.PASS:
	default:
		t.Errorf("runTestCase() fail = %s", tr.Reason)
	}
}
