package tests

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Logger interface {
	OnAllStart(testCount int)
	OnAllEnd()
	OnTestStart(tc *TestCase)
	OnTestEnd(tc *TestCase)
	Log(tc *TestCase, message string)
	Logf(tc *TestCase, format string, a ...any)
}

var (
	EmptyLogger   = LoggerFunc(func(tc *TestCase, event int, message string) {})
	ConsoleLogger = LoggerFunc(func(tc *TestCase, event int, message string) { fmt.Print(message) })
)

// LoggerFunc
const (
	LoggerEventMessage = 0
	LoggerEventStart   = 1
	LoggerEventEnd     = 2
)

type LoggerFunc func(tc *TestCase, event int, message string)

func (fn LoggerFunc) OnAllStart(testCount int) {
	fn(nil, LoggerEventStart, "")
}

func (fn LoggerFunc) OnAllEnd() {
	fn(nil, LoggerEventEnd, "")
}

func (fn LoggerFunc) OnTestStart(tc *TestCase) {
	fn(tc, LoggerEventStart, "")
}

func (fn LoggerFunc) OnTestEnd(tc *TestCase) {
	fn(tc, LoggerEventEnd, "")
}

func (fn LoggerFunc) Log(tc *TestCase, message string) {
	fn(tc, 0, message)
}

func (fn LoggerFunc) Logf(tc *TestCase, format string, a ...any) {
	fn(tc, 0, fmt.Sprintf(format, a...))
}

// DumpLogger
type DumpLogger struct {
	dumpRoot string
	channels []strings.Builder
}

func NewDumpLogger(dumpRoot string) *DumpLogger {
	if dumpRoot == "" || !filepath.IsAbs(dumpRoot) {
		panic("dumpRoot 必须不为空且是个绝对路径")
	}

	return &DumpLogger{dumpRoot: dumpRoot}
}

func (l *DumpLogger) OnAllStart(testCount int) {
	l.channels = make([]strings.Builder, testCount+1)
}

func (l *DumpLogger) OnAllEnd() {
	l.channels = nil
}

func (l *DumpLogger) checkIndexRange(index int) {
	if index <= 0 || index > len(l.channels) {
		panic(fmt.Sprintf("index(%d) must in 1~testCount(%d)", index, len(l.channels)))
	}
}

func (l *DumpLogger) getWriter(tc *TestCase) io.Writer {
	if tc == nil {
		return os.Stdout
	}

	l.checkIndexRange(tc.index)
	return &l.channels[tc.index]
}
func (l *DumpLogger) closeWriter(tc *TestCase) {
	if tc == nil {
		return
	}

	l.checkIndexRange(tc.index)

	w := &l.channels[tc.index]
	dumpFile := filepath.Join(l.dumpRoot, tc.shortFileName)
	_ = filePutContents(dumpFile, w.String())
	w.Reset()
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
