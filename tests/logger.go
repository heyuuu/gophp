package tests

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Logger interface {
	OnAllStart(startTime time.Time, testCount int)
	OnAllEnd(endTime time.Time, summary *Summary)
	OnTestStart(tc *TestCase)
	OnTestEnd(tc *TestCase)
	Log(tc *TestCase, message string)
	Logf(tc *TestCase, format string, a ...any)
}

type DumpLogger struct {
	dumpRoot string
	channels []*strings.Builder
}

func NewDumpLogger(dumpRoot string) *DumpLogger {
	return &DumpLogger{dumpRoot: dumpRoot}
}

func (l *DumpLogger) OnAllStart(startTime time.Time, testCount int) {
	l.channels = make([]*strings.Builder, testCount)

	l.Log(nil, "=====================================================================\n")
	l.Log(nil, "TIME START "+timeFormat(startTime, "Y-m-d H:i:s")+"\n")
	l.Log(nil, "=====================================================================\n")
}

func (l *DumpLogger) OnAllEnd(endTime time.Time, summary *Summary) {
	l.channels = nil

	l.Log(nil, "=====================================================================\n")
	l.Log(nil, "TIME END "+timeFormat(endTime, "Y-m-d H:i:s")+"\n")
	l.Log(nil, "=====================================================================\n")

	l.Log(nil, summary.Summary())
}

func (l *DumpLogger) getWriter(tc *TestCase) io.Writer {
	if tc == nil {
		return os.Stdout
	}

	testIndex := tc.index
	if testIndex <= 0 || testIndex > len(l.channels) {
		panic(fmt.Sprintf("index(%d) must in 1~testCount(%d)", testIndex, len(l.channels)))
	}

	index := testIndex - 1
	if l.channels[index] == nil {
		l.channels[index] = new(strings.Builder)
	}
	return l.channels[index]
}
func (l *DumpLogger) closeWriter(tc *TestCase) {
	if tc == nil {
		return
	}

	testIndex := tc.index
	if testIndex <= 0 || testIndex > len(l.channels) {
		panic(fmt.Sprintf("index(%d) must in 1~testCount(%d)", testIndex, len(l.channels)))
	}

	index := testIndex - 1
	w := l.channels[index]
	l.channels[index] = nil

	if w != nil {
		dumpFile := filepath.Join(l.dumpRoot, tc.shortFileName)
		_ = filePutContents(dumpFile, w.String())
	}
}

func (l *DumpLogger) OnTestStart(tc *TestCase) {
	// 触发 channel 初始化
	l.getWriter(tc)
}

func (l *DumpLogger) OnTestEnd(tc *TestCase) {
	l.closeWriter(tc)
}

func (l *DumpLogger) Log(tc *TestCase, message string) {
	w := l.getWriter(tc)
	_, _ = fmt.Fprint(w, message)
}

func (l *DumpLogger) Logf(tc *TestCase, format string, a ...any) {
	w := l.getWriter(tc)
	_, _ = fmt.Fprintf(w, format, a...)
}
