package pfmt

import (
	"github.com/heyuuu/gophp/builtin/ascii"
	"strings"
)

type byteReader struct {
	s string
	l int
	i int
}

func newByteReader(s string) *byteReader {
	return &byteReader{
		s: s,
		l: len(s),
		i: 0,
	}
}

func (r *byteReader) valid() bool {
	return r.i < r.l
}

func (r *byteReader) curr() byte {
	if r.i < r.l {
		return r.s[r.i]
	}
	return 0
}

func (r *byteReader) offset(offset int) byte {
	if r.i+offset < r.l {
		return r.s[r.i+offset]
	}
	return 0
}

func (r *byteReader) read() byte {
	if r.i < r.l {
		c := r.s[r.i]
		r.i++
		return c
	}
	return 0
}

func (r *byteReader) isAscii() bool {
	return ascii.IsAscii(r.curr())
}

func (r *byteReader) isDigit() bool {
	return ascii.IsDigit(r.curr())
}

func (r *byteReader) isLower() bool {
	return ascii.IsLower(r.curr())
}

func (r *byteReader) readDec() int {
	num := 0
	for r.isDigit() {
		digit := int(r.read() - '0')
		num = num*10 + digit
	}
	return num
}

func (r *byteReader) tryRead(str string) bool {
	if strings.HasPrefix(r.s[r.l:], str) {
		r.l += len(str)
		return true
	}
	return false
}
