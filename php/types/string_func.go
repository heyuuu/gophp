package types

import (
	b "github.com/heyuuu/gophp/builtin"
)

func ZstrAlloc(str *String, _len int) {
	*str = *ZendStringAlloc(_len, 0)
}

func emptyString(len_ int) string {
	return string(make([]byte, len_))
}

func ZendStringAlloc(len_ int, persistent int) *String {
	var str_ = emptyString(len_)
	return NewString(str_)
}
func ZendStringSafeAlloc(n int, m int, l int, persistent int) *String {
	// todo 不太明白参数作用，仅从纯代码功能重构
	var len_ = n*m + l
	return ZendStringAlloc(len_, persistent)
}
func ZendStringExtend(s *String, len_ int) *String {
	b.Assert(len_ >= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr + emptyString(len_-len(oldStr))
	//s.DelRefcount()
	return NewString(newStr)
}
func ZendStringTruncate(s *String, len_ int) *String {
	b.Assert(len_ <= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr[:len_]
	return NewString(newStr)
}
func ZendStringSafeRealloc(s *String, n int, m int, l int, persistent int) *String {
	var ret *String
	ret = ZendStringSafeAlloc(n, m, l, persistent)
	memcpy(ret.GetVal(), s.GetVal(), b.Min(n*m+l, s.GetLen())+1)
	return ret
}
