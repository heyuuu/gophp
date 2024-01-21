package tests

import (
	"strings"
	"time"
)

type TestCase struct {
	// Case 名及路径
	FileName string
	File     string
	// Case 文件解析的信息
	Desc          string
	CaptureStdin  bool
	CaptureStdout bool
	CaptureStderr bool
	UseCgi        bool
	UseDbg        bool
	Sections      map[string]string
}

func NewTestCase(file string, fileName string, sections map[string]string) *TestCase {
	tc := &TestCase{File: file, FileName: fileName, Sections: sections}
	tc.init()
	return tc
}

func (tc *TestCase) init() {
	tc.Desc = strings.TrimSpace(tc.Sections["TEST"])

	if capture, ok := tc.Sections["CAPTURE_STDIO"]; ok {
		lcCapture := strings.ToLower(capture)
		tc.CaptureStdin = strings.Contains(lcCapture, "stdin")
		tc.CaptureStdout = strings.Contains(lcCapture, "stdout")
		tc.CaptureStderr = strings.Contains(lcCapture, "stderr")
	}

	/* For GET/POST/PUT tests, check if cgi sapi is available and if it is, use it. */

}

func (tc *TestCase) exists(key string) bool {
	_, ok := tc.Sections[key]
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

type IniEntry struct {
	name  string
	value string
}

type Environments struct {
	m map[string]string
}

func (env *Environments) Set(key string, value string) {
	if env.m == nil {
		env.m = map[string]string{}
	}
	env.m[key] = value
}
func (env *Environments) Get(key string) string {
	if env.m == nil {
		return ""
	}
	return env.m[key]
}
