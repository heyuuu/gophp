package perr

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

type stack []uintptr

const pkgPrefix = "github.com/heyuuu/gophp/php/perr."

func (s stack) Format(w io.Writer) {
	frames := runtime.CallersFrames(s)
	for i := 0; i < len(s); i++ {
		frame, _ := frames.Next()
		if strings.HasPrefix(frame.Function, pkgPrefix) {
			continue
		}
		_, _ = fmt.Fprintf(w, "\n%s\n\t%s:%d", frame.Function, frame.File, frame.Line)
	}
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var callStack stack = pcs[0:n]
	return &callStack
}

type withStack struct {
	err   error
	stack *stack
}

func (e withStack) Error() string { return e.err.Error() }
func (e withStack) Unwrap() error { return e.err }
func (e withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		fmt.Fprintf(s, "%+v", e.err)
		if s.Flag('+') {
			e.stack.Format(s)
		}
	case 's':
		fmt.Fprintf(s, "%s", e.err)
	case 'q':
		fmt.Fprintf(s, "%q", e.err)
	}
}

func WithStack(err error) error {
	return &withStack{
		err:   err,
		stack: callers(),
	}
}
