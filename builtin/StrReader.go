package builtin

import "strings"

// StrReader 用于替代 C 中使用 *byte 读取字符串的相关操作
type StrReader struct {
	str string
}

func NewStrReader(str string) *StrReader { return &StrReader{str} }
func (r StrReader) Copy() *StrReader     { return NewStrReader(r.str) }

func (r *StrReader) IsValid() bool { return r.str != "" }
func (r *StrReader) Curr() byte {
	if r.str == "" {
		return 0
	}
	return r.str[0]
}

func (r *StrReader) Read() byte {
	if r.str == "" {
		return 0
	}

	c := r.str[0]
	r.str = r.str[1:]
	return c
}

func (r *StrReader) Next() {
	if r.str == "" {
		return
	}
	r.str = r.str[1:]
}

func (r *StrReader) SkipBytes(cutSet string) {
	r.str = strings.TrimLeft(r.str, cutSet)
}

func (r *StrReader) SkipSpace() {
	r.str = strings.TrimLeft(r.str, " ")
}
