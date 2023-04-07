package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"strings"
)

func ZSTR_CHAR(c int) *String { return oneCharStrings[c] }

func ZSTR_ALLOCA_ALLOC(str *String, _len int) {
	*str = *ZendStringAlloc(_len, 0)
}

func ZendStringAlloc(len_ int, persistent int) *String {
	var str_ = b.EmptyString(len_)
	return NewString(str_)
}
func ZendStringSafeAlloc(n int, m int, l int, persistent int) *String {
	// todo 不太明白参数作用，仅从纯代码功能重构
	var len_ = n*m + l
	return ZendStringAlloc(len_, persistent)
}
func ZendStringExtend(s *String, len_ int, persistent int) *String {
	b.Assert(len_ >= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr + b.EmptyString(len_-len(oldStr))
	s.DelRefcount()
	return NewString(newStr)
}
func ZendStringTruncate(s *String, len_ int, persistent int) *String {
	b.Assert(len_ <= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr[:len_]
	s.DelRefcount()
	return NewString(newStr)
}
func ZendStringSafeRealloc(s *String, n int, m int, l int, persistent int) *String {
	var ret *String
	//if s.GetRefcount() == 1 {
	//	ret = (*String)(SafePerealloc(s, n, m, ZEND_MM_ALIGNED_SIZE(_ZSTR_STRUCT_SIZE(l)), persistent))
	//	ret.SetLen(n*m + l)
	//	ZendStringForgetHashVal(ret)
	//	return ret
	//}
	ret = ZendStringSafeAlloc(n, m, l, persistent)
	memcpy(ret.GetVal(), s.GetVal(), b.Min(n*m+l, s.GetLen())+1)
	s.DelRefcount()
	return ret
}
func ZendStringEquals(s1 *String, s2 *String) ZendBool {
	return IntBool(s1.GetStr() == s2.GetStr())
}
func ZendStringEqualsCi(s1 *String, s2 *String) bool {
	return strCaseEquals(s1.GetStr(), s2.GetStr())
}
func ZendStringEqualsLiteralCi(str *String, c string) bool {
	return strCaseEquals(str.GetStr(), c)
}
func ZendStringEqualsLiteral(str *String, literal string) bool {
	return str.GetStr() == literal
}

func strCaseEquals(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	}
	return strings.EqualFold(str1, str2)
}

/**
 * Interned String 相关
 */
var IsInRequestForInternedString bool = false

func ZendInternedStringsInit() {
	InternedStringsPermanent.Clean()
	IsInRequestForInternedString = false
}
func ZendInternedStringsDtor() {
	InternedStringsPermanent.Clean()
}

func InitInternedString(str string) string {
	if IsInRequestForInternedString {
		/* Check for permanent strings, the table is readonly at this point. */
		if ret, ok := InternedStringsPermanent.Get(str); ok {
			return ret
		}

		ret, _ := zend.CG__().InternedStrings.Get(str)
		return ret
	} else {
		ret, _ := InternedStringsPermanent.GetOrInsert(str)
		return ret
	}
}

func InitInternedZendString(str string) *String {
	var interned = InitInternedString(str)
	return NewString(interned)
}

func ZendNewInternedString(str *String) *String {
	return InitInternedZendString(str.GetStr())
}

func ZendStringInitInterned(str *byte, size int, permanent int) *String {
	return InitInternedZendString(b.CastStr(str, size))
}

func ZendInternedStringsActivate() {
	zend.CG__().InternedStrings = NewInternedStrings()
}
func ZendInternedStringsDeactivate() {
	zend.CG__().InternedStrings.Destroy()
}
func ZendInternedStringsSwitchStorage(inRequest bool) {
	IsInRequestForInternedString = inRequest
}
