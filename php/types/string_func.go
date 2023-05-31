package types

import (
	b "github.com/heyuuu/gophp/builtin"
)

func ZendStringAlloc(len_ int, persistent int) *String {
	var str_ = string(make([]byte, len_))
	return NewString(str_)
}
func ZendStringRealloc(s *String, size int) *String {
	ret := ZendStringAlloc(size, 0)
	memcpy(ret.GetVal(), s.GetVal(), b.Min(size, s.GetLen())+1)
	return ret
}
func ZendStringTruncate(s *String, len_ int) *String {
	b.Assert(len_ <= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr[:len_]
	return NewString(newStr)
}
