package pfmt

import (
	"io"
)

func Fprintf(f io.Writer, fmt string, args ...any) (int, error) {
	p := printer{}
	p.doPrint(fmt, args)
	return f.Write(p.buf)
}

func Sprintf(fmt string, args ...any) string {
	p := printer{}
	p.doPrint(fmt, args)
	return string(p.buf)
}

func Snprintf(size int, fmt string, args ...any) string {
	if size <= 0 {
		return ""
	}

	s := Sprintf(fmt, args)
	if len(s) > size {
		return s[:size]
	}
	return s
}
