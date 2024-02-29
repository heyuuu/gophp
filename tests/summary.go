package tests

import (
	"fmt"
	"github.com/heyuuu/gophp/shim/cmp"
	"github.com/heyuuu/gophp/shim/slices"
	"strings"
	"time"
)

type Summary struct {
	StartTime time.Time
	EndTime   time.Time

	//
	total       int
	sumResults  map[ResultType]int
	failedTests map[ResultType][]*Result
}

func NewSummary() *Summary {
	return &Summary{
		sumResults:  map[ResultType]int{},
		failedTests: map[ResultType][]*Result{},
	}
}

func (s *Summary) AddResult(tc *TestCase, result *Result) {
	s.total++

	s.sumResults[result.MainType()]++

	if result.slow {
		s.failedTests[SLOW] = append(s.failedTests[SLOW], result)
	}
	for _, typ := range result.types {
		if typ == BORK || typ == WARN || typ == XFAIL || typ == FAIL {
			s.failedTests[typ] = append(s.failedTests[typ], result)
		}
	}
}

func (s *Summary) HasResult(targetTypes ...ResultType) bool {
	for _, typ := range targetTypes {
		if s.sumResults[typ] > 0 {
			return true
		}
	}
	return false
}

func (s *Summary) Summary() string {
	total, sumResults := s.total, s.sumResults

	percents := make(map[ResultType]float64, len(sumResults))
	for typ, typCount := range sumResults {
		percents[typ] = float64(typCount) * 100 / float64(total)
	}

	// exec
	xTotal := total - sumResults[SKIP] - sumResults[BORK]
	xPercents := make(map[ResultType]float64, len(sumResults))
	if xTotal > 0 {
		for typ, typCount := range sumResults {
			xPercents[typ] = float64(typCount) * 100 / float64(xTotal)
		}
	}

	// print
	var p summaryPrinter

	p.tableStart("TEST RESULT SUMMARY")
	p.linef("Number of tests : %4d          %8d", total, xTotal)
	p.linef("Tests borked    : %4d (%5.1f%%) --------", sumResults[BORK], percents[BORK])
	p.linef("Tests skipped   : %4d (%5.1f%%) --------", sumResults[SKIP], percents[SKIP])
	p.linef("Tests warned    : %4d (%5.1f%%) (%5.1f%%)", sumResults[WARN], percents[WARN], xPercents[WARN])
	p.linef("Tests failed    : %4d (%5.1f%%) (%5.1f%%)", sumResults[FAIL], percents[FAIL], xPercents[FAIL])
	p.linef("Expected fail   : %4d (%5.1f%%) (%5.1f%%)", sumResults[XFAIL], percents[XFAIL], xPercents[XFAIL])
	p.linef("Tests passed    : %4d (%5.1f%%) (%5.1f%%)", sumResults[PASS], percents[PASS], xPercents[PASS])
	p.singleLine()
	p.linef("Time taken      : %4.2f seconds", s.EndTime.Sub(s.StartTime).Seconds())
	p.doubleLine()

	s.buildFailedTestSummary(&p)

	return p.String()
}

func (s *Summary) buildFailedTestSummary(p *summaryPrinter) {
	if slowResults := s.failedTests[SLOW]; len(slowResults) > 0 {
		slices.SortFunc(slowResults, func(a, b *Result) int {
			return cmp.Compare(b.useTime, a.useTime)
		})

		p.tableStart("SLOW TEST SUMMARY")
		for _, result := range slowResults {
			p.linef("(%.3f s) %s", result.useTime.Seconds(), result.tc.ShowName())
		}
		p.doubleLine()
	}

	s.buildFailedTestSummaryByType(p, XFAIL, "EXPECTED FAILED TEST SUMMARY")
	s.buildFailedTestSummaryByType(p, BORK, "BORKED TEST SUMMARY")
	s.buildFailedTestSummaryByType(p, FAIL, "FAILED TEST SUMMARY")
	s.buildFailedTestSummaryByType(p, WARN, "WARNED TEST SUMMARY")
	s.buildFailedTestSummaryByType(p, LEAK, "LEAKED TEST SUMMARY")
	s.buildFailedTestSummaryByType(p, XLEAK, "XLEAKED TEST SUMMARY")
}

func (s *Summary) buildFailedTestSummaryByType(p *summaryPrinter, typ ResultType, title string) {
	if typeResults := s.failedTests[typ]; len(typeResults) > 0 {
		p.tableStart(title)
		for _, result := range typeResults {
			p.linef("%s%s", result.tc.ShowName(), result.info)
		}
		p.doubleLine()
	}
}

type summaryPrinter struct {
	buf strings.Builder
}

func (p *summaryPrinter) linef(format string, a ...any) {
	_, _ = fmt.Fprintf(&p.buf, format, a...)
	p.buf.WriteByte('\n')
}

func (p *summaryPrinter) doubleLine() {
	p.buf.WriteString("=====================================================================\n")
}
func (p *summaryPrinter) singleLine() {
	p.buf.WriteString("---------------------------------------------------------------------\n")
}

func (p *summaryPrinter) tableStart(title string) {
	p.doubleLine()
	p.linef(title)
	p.singleLine()
}

func (p *summaryPrinter) String() string {
	return p.buf.String()
}
