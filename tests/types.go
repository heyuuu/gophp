package tests

import (
	"strings"
	"time"
)

type TestCase struct {
	File          string
	ShortFileName string

	TestName      string
	CaptureStdin  bool
	CaptureStdout bool
	CaptureStderr bool
	UseCgi        bool
	UseDbg        bool

	sections map[string]string
}

func NewTestCase(file string, shortFileName string, sections map[string]string) *TestCase {
	tc := &TestCase{File: file, ShortFileName: shortFileName, sections: sections}
	tc.init()
	return tc
}

func (tc *TestCase) init() {
	tc.TestName = strings.TrimSpace(tc.sections["TEST"])

	if capture, ok := tc.sections["CAPTURE_STDIO"]; ok {
		lcCapture := strings.ToLower(capture)
		tc.CaptureStdin = strings.Contains(lcCapture, "stdin")
		tc.CaptureStdout = strings.Contains(lcCapture, "stdout")
		tc.CaptureStderr = strings.Contains(lcCapture, "stderr")
	}

	/* For GET/POST/PUT tests, check if cgi sapi is available and if it is, use it. */

}

func (tc *TestCase) exists(key string) bool {
	_, ok := tc.sections[key]
	return ok
}

type ResultType string

const (
	PASS  ResultType = "PASS"
	BORK  ResultType = "BORK"
	FAIL  ResultType = "FAIL"
	WARN  ResultType = "WARN"
	LEAK  ResultType = "LEAK"
	XFAIL ResultType = "XFAIL"
	XLEAK ResultType = "XLEAK"
	SLOW  ResultType = "SLOW"
	SKIP  ResultType = "SKIP"
)

func (t ResultType) ShowName() string {
	switch t {
	case PASS, BORK, FAIL, WARN, LEAK, XFAIL, XLEAK, SLOW:
		return string(t) + "ED"
	case SKIP:
		return "SKIPPED"
	default:
		return string(t)
	}
}

func ValidResultType(t ResultType) bool {
	switch t {
	case PASS, BORK, FAIL, WARN, LEAK, XFAIL, XLEAK, SLOW, SKIP:
		return true
	default:
		return false
	}
}

type TestResult struct {
	Case    *TestCase
	Type    ResultType
	Reason  string
	UseTime time.Duration
}

func NewTestResult(Case *TestCase, Type ResultType, reason string, useTime time.Duration) *TestResult {
	return &TestResult{Case: Case, Type: Type, Reason: reason, UseTime: useTime}
}

type Summary struct {
	StartTime    time.Time
	EndTime      time.Time
	Results      []TestResult
	ExtSkipped   int
	ExtTested    int
	IgnoredByExt int
}

func (summary *Summary) AddResult(tc *TestCase, r *TestResult) {
	summary.Results = append(summary.Results, *r)
}

type Environments struct {
}
