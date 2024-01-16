package main

import (
	"errors"
	"github.com/heyuuu/gophp/tests"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	pwd, _ := os.Getwd()
	os.Chdir(filepath.Dir(pwd))
}

func runTestCase(t *testing.T, testName string) {
	testFile := filepath.Join("/Users/heyu/Code/sik/gophp/testcases", testName)
	tr, err := tests.RunCaseBuiltin(0, testName, testFile)
	if err != nil {
		if errors.Is(err, tests.ErrUnsupportedSection) {
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
