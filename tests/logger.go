package tests

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Logger interface {
	OnAllStart()
	OnAllEnd()
	OnTestStart(tc *TestCase)
	OnTestEnd(tc *TestCase)
	Log(tc *TestCase, message string)
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

func (fn LoggerFunc) OnAllStart() {
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
	fn(tc, LoggerEventMessage, message)
}

// DumpLogger
type DumpLogger struct {
	dumpRoot string
	channels map[int]*strings.Builder
}

func NewDumpLogger(dumpRoot string) *DumpLogger {
	if dumpRoot == "" || !filepath.IsAbs(dumpRoot) {
		panic("dumpRoot 必须不为空且是个绝对路径")
	}

	return &DumpLogger{dumpRoot: dumpRoot}
}

func (l *DumpLogger) OnAllStart() {
	l.channels = make(map[int]*strings.Builder)
}

func (l *DumpLogger) OnAllEnd() {
	l.channels = nil
}

func (l *DumpLogger) getWriter(tc *TestCase) io.Writer {
	if tc == nil {
		return os.Stdout
	}
	w := l.channels[tc.index]
	if w == nil {
		w = new(strings.Builder)
		l.channels[tc.index] = w
	}
	return w
}
func (l *DumpLogger) closeWriter(tc *TestCase) {
	if tc == nil {
		return
	}

	w := l.channels[tc.index]
	if w == nil {
		return
	}

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

// SyncLogger
type syncLog struct {
	tc      *TestCase
	event   int
	message string
}
type SyncLogger struct {
	inner Logger
	logCh chan syncLog
	wg    sync.WaitGroup
}

func NewSyncLogger(inner Logger) *SyncLogger {
	return &SyncLogger{inner: inner}
}

func (l *SyncLogger) OnAllStart() {
	if l.logCh == nil {
		l.logCh = make(chan syncLog, 10)
		go func() {
			for log := range l.logCh {
				l.handle(log)
			}
		}()
	}
	l.push(nil, LoggerEventStart, "")
}

func (l *SyncLogger) OnAllEnd() {
	l.push(nil, LoggerEventEnd, "")
	if l.logCh != nil {
		close(l.logCh)
		l.logCh = nil
	}
}

func (l *SyncLogger) OnTestStart(tc *TestCase) {
	l.push(tc, LoggerEventStart, "")
}

func (l *SyncLogger) OnTestEnd(tc *TestCase) {
	l.push(tc, LoggerEventEnd, "")
}

func (l *SyncLogger) Log(tc *TestCase, message string) {
	l.push(tc, LoggerEventMessage, message)
}

func (l *SyncLogger) push(tc *TestCase, event int, message string) {
	if c := l.logCh; c != nil {
		c <- syncLog{tc: tc, event: event, message: message}
		l.wg.Add(1)
	}
}

func (l *SyncLogger) handle(log syncLog) {
	switch log.event {
	case LoggerEventStart:
		if log.tc == nil {
			l.inner.OnAllStart()
		} else {
			l.inner.OnTestStart(log.tc)
		}
	case LoggerEventEnd:
		if log.tc == nil {
			l.inner.OnAllEnd()
		} else {
			l.inner.OnTestEnd(log.tc)
		}
	default:
		l.inner.Log(log.tc, log.message)
	}
}
