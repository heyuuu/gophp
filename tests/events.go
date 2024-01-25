package tests

import (
	"fmt"
	"sync"
	"time"
)

type EventHandler interface {
	OnAllStart(startTime time.Time, testCount int, extSkipped int, ignoreByExt int)
	OnAllEnd(endTime time.Time)
	OnTestStart(testIndex int, tc *TestCase)
	OnTestEnd(testIndex int, tc *TestCase, tr *TestResult)
	Log(testIndex int, message string)
}

type DefaultEventHandler struct {
	w         func(verbose int, message string)
	testCount int
	summary   *Summary
}

func NewDefaultEventHandler(w func(verbose int, message string)) *DefaultEventHandler {
	return &DefaultEventHandler{w: w}
}

func (l *DefaultEventHandler) OnAllStart(startTime time.Time, testCount int, extSkipped int, ignoreByExt int) {
	l.summary = &Summary{}
	l.summary.StartTime = startTime
	l.summary.ExtSkipped = extSkipped
	l.summary.IgnoredByExt = ignoreByExt

	l.testCount = testCount

	l.printLn("=====================================================================")
	l.printLn("TIME START " + timeFormat(startTime, "Y-m-d H:i:s"))
	l.printLn("=====================================================================")
}

func (l *DefaultEventHandler) OnAllEnd(endTime time.Time) {
	l.summary.EndTime = endTime

	l.printLn("=====================================================================")
	l.printLn("TIME END " + timeFormat(endTime, "Y-m-d H:i:s"))
	l.printLn("=====================================================================")

	l.logSummary(l.summary)
	l.logExtSummary(l.summary)
}

func (l *DefaultEventHandler) OnTestStart(testIndex int, tc *TestCase) {
	l.printLn("")
	l.printLn("=================")
	l.printLn("TEST " + tc.File)
	l.printLn(fmt.Sprintf("TEST %d/%d [%s]", testIndex+1, l.testCount, tc.FileName))
}

func (l *DefaultEventHandler) OnTestEnd(testIndex int, tc *TestCase, tr *TestResult) {
	if tr != nil {
		l.summary.AddResult(tc, tr)
		l.printLn(fmt.Sprintf("%s %s [%s] %s", tr.Type, tr.Case.TestName, tc.FileName, tr.Reason))
	}
}

func (l *DefaultEventHandler) Log(testIndex int, message string) {
	l.printLn(message)
}

func (l *DefaultEventHandler) printLn(message string) {
	l.w(0, message)
	l.w(0, "\n")
}

func (l *DefaultEventHandler) logSummary(summary *Summary) {
	total := len(summary.Results) + summary.IgnoredByExt

	sumResults := map[ResultType]int{}
	for _, result := range summary.Results {
		sumResults[result.Type]++
	}
	sumResults[SKIP] += summary.IgnoredByExt

	totalPrec := func(v int) float64 {
		if total == 0 {
			return 0
		}
		return float64(v) * 100 / float64(total)
	}
	precResults := map[ResultType]float64{}
	for typ, v := range sumResults {
		precResults[typ] = (float64(v) * 100.0) / float64(total)
	}

	execTotal := total - sumResults[SKIP] - sumResults[BORK]
	execPrec := func(v int) float64 {
		if execTotal == 0 {
			return 0
		}
		return float64(v) * 100 / float64(execTotal)
	}

	useTime := summary.EndTime.Sub(summary.StartTime).Seconds()

	l.printLn("=====================================================================")
	l.printLn("TEST RESULT SUMMARY")
	l.printLn("---------------------------------------------------------------------")
	l.printLn(fmt.Sprintf("Exts skipped    : %4d", summary.ExtSkipped))
	l.printLn(fmt.Sprintf("Exts tested     : %4d", summary.ExtTested))
	l.printLn("---------------------------------------------------------------------")
	l.printLn(fmt.Sprintf("Number of tests : %4d          %8d", total, execTotal))
	l.printLn(fmt.Sprintf("Tests borked    : %4d (%5.1f%%) --------", sumResults[BORK], totalPrec(sumResults[BORK])))
	l.printLn(fmt.Sprintf("Tests skipped   : %4d (%5.1f%%) --------", sumResults[SKIP], totalPrec(sumResults[SKIP])))
	l.printLn(fmt.Sprintf("Tests warned    : %4d (%5.1f%%) (%5.1f%%)", sumResults[WARN], totalPrec(sumResults[WARN]), execPrec(sumResults[WARN])))
	l.printLn(fmt.Sprintf("Tests failed    : %4d (%5.1f%%) (%5.1f%%)", sumResults[FAIL], totalPrec(sumResults[FAIL]), execPrec(sumResults[FAIL])))
	l.printLn(fmt.Sprintf("Expected fail   : %4d (%5.1f%%) (%5.1f%%)", sumResults[XFAIL], totalPrec(sumResults[XFAIL]), execPrec(sumResults[XFAIL])))
	l.printLn(fmt.Sprintf("Tests passed    : %4d (%5.1f%%) (%5.1f%%)", sumResults[PASS], totalPrec(sumResults[PASS]), execPrec(sumResults[PASS])))
	l.printLn("---------------------------------------------------------------------")
	l.printLn(fmt.Sprintf("Time taken      : %4.2f seconds", useTime))
	l.printLn("=====================================================================")
}

func (l *DefaultEventHandler) logExtSummary(summary *Summary) {
	var m = map[ResultType][]TestResult{}
	for _, result := range summary.Results {
		m[result.Type] = append(m[result.Type], result)
	}
	showType := func(typ ResultType, title string) {
		typeResults := m[typ]
		if len(typeResults) == 0 {
			return
		}

		l.printLn("=====================================================================")
		l.printLn(title)
		l.printLn("---------------------------------------------------------------------")
		for _, result := range typeResults {
			l.printLn(fmt.Sprintf("%s [%s] %s", result.Case.TestName, result.Case.FileName, result.Reason))
		}
		l.printLn("=====================================================================")
	}

	showType(SLOW, "SLOW TEST SUMMARY")
	showType(XFAIL, "EXPECTED FAILED TEST SUMMARY")
	showType(BORK, "BORKED TEST SUMMARY")
	showType(FAIL, "FAILED TEST SUMMARY")
	showType(WARN, "WARNED TEST SUMMARY")
	showType(LEAK, "LEAKED TEST SUMMARY")
	showType(XLEAK, "XLEAKED TEST SUMMARY")
}

type ParallelHandler struct {
	inner        EventHandler
	testIndex    int
	testCount    int
	testDone     []bool
	testDoneChan chan int
	testEvents   [][]func()
	wg           sync.WaitGroup
}

func NewParallelHandler(inner EventHandler) *ParallelHandler {
	return &ParallelHandler{inner: inner}
}

func (p *ParallelHandler) OnAllStart(startTime time.Time, testCount int, extSkipped int, ignoreByExt int) {
	p.inner.OnAllStart(startTime, testCount, extSkipped, ignoreByExt)
	p.testCount = testCount
	p.testDone = make([]bool, testCount)
	p.testEvents = make([][]func(), testCount)
	p.testDoneChan = make(chan int, 10)
	p.wg.Add(testCount)

	go func() {
		var index int
		for {
			select {
			case index = <-p.testDoneChan:
				p.testDone[index] = true
				if index == p.testIndex {
					for p.testIndex < p.testCount && p.testDone[p.testIndex] {
						for _, event := range p.testEvents[p.testIndex] {
							event()
						}
						p.wg.Done()
						p.testIndex++
					}
					if p.testIndex == p.testCount {
						return
					}
				}
			}
		}
	}()
}

func (p *ParallelHandler) OnAllEnd(endTime time.Time) {
	p.wg.Wait()
	p.inner.OnAllEnd(endTime)
}

func (p *ParallelHandler) sortedEvent(index int, done bool, handler func()) {
	p.testEvents[index] = append(p.testEvents[index], handler)
	if done {
		p.testDoneChan <- index
	}
}

func (p *ParallelHandler) OnTestStart(testIndex int, tc *TestCase) {
	p.sortedEvent(testIndex, false, func() {
		p.inner.OnTestStart(testIndex, tc)
	})
}

func (p *ParallelHandler) OnTestEnd(testIndex int, tc *TestCase, tr *TestResult) {
	p.sortedEvent(testIndex, true, func() {
		p.inner.OnTestEnd(testIndex, tc, tr)
	})
}

func (p *ParallelHandler) Log(testIndex int, message string) {
	p.sortedEvent(testIndex, false, func() {
		p.inner.Log(testIndex, message)
	})
}
