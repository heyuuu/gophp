package tests

import (
	"time"
)

type TestCase struct {
	File          string
	ShortFileName string
	TestName      string

	sections map[string]string
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
