// <<generate>>

package types

import (
	b "sik/builtin"
	"sik/zend"
)

func ZSTR_VAL(zstr *ZendString) []byte { return zstr.GetVal() }

var ZendEmptyString *ZendString = nil

func STR_EMPTY_ALLOC() *ZendString  { return ZendEmptyString }
func ZSTR_EMPTY_ALLOC() *ZendString { return ZendEmptyString }

func ZSTR_CHAR(c int) *ZendString { return ZendOneCharString[c] }
func ZSTR_KNOWN(str string) *ZendString {
	return NewZendStringPersistent(str, true)
}

func ZSTR_ALLOCA_ALLOC(str *ZendString, _len int, use_heap any) {
	*str = *ZendStringAlloc(_len, 0)
}

func ZendStringForgetHashVal(s *ZendString) {
	s.SetH(0)
	s.DelGcFlags(IS_STR_VALID_UTF8)
}

func ZendStringAlloc(len_ int, persistent int) *ZendString {
	var str_ = b.EmptyString(len_)
	return NewZendStringPersistent(str_, persistent != 0)
}
func ZendStringSafeAlloc(n int, m int, l int, persistent int) *ZendString {
	// todo 不太明白参数作用，仅从纯代码功能重构
	var len_ = n*m + l
	return ZendStringAlloc(len_, persistent)
}
func ZendStringInit(str *byte, len_ int, persistent int) *ZendString {
	var str_ = b.CastStr(str, len_)
	return NewZendStringPersistent(str_, persistent != 0)
}
func ZendStringExtend(s *ZendString, len_ int, persistent int) *ZendString {
	zend.ZEND_ASSERT(len_ >= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr + b.EmptyString(len_-len(oldStr))
	s.DelRefcount()
	return NewZendStringPersistent(newStr, persistent != 0)
}
func ZendStringTruncate(s *ZendString, len_ int, persistent int) *ZendString {
	zend.ZEND_ASSERT(len_ <= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr[:len_]
	s.DelRefcount()
	return NewZendStringPersistent(newStr, persistent != 0)
}
func ZendStringSafeRealloc(s *ZendString, n int, m int, l int, persistent int) *ZendString {
	var ret *ZendString
	//if s.GetRefcount() == 1 {
	//	ret = (*ZendString)(SafePerealloc(s, n, m, ZEND_MM_ALIGNED_SIZE(_ZSTR_STRUCT_SIZE(l)), persistent))
	//	ret.SetLen(n*m + l)
	//	ZendStringForgetHashVal(ret)
	//	return ret
	//}
	ret = ZendStringSafeAlloc(n, m, l, persistent)
	memcpy(ret.GetVal(), s.GetVal(), b.Min(n*m+l, s.GetLen())+1)
	s.DelRefcount()
	return ret
}
func ZendStringFree(s *ZendString) {
	zend.ZEND_ASSERT(s.GetRefcount() <= 1)
	b.Free(s)
}
func ZendStringEfree(s *ZendString) {
	zend.ZEND_ASSERT(s.GetRefcount() <= 1)
	zend.ZEND_ASSERT((s.GetGcFlags() & IS_STR_PERSISTENT) == 0)
	b.Free(s)
}
func ZendStringRelease(s *ZendString) {
	if s.DelRefcount() == 0 {
		b.Free(s)
	}
}
func ZendStringReleaseEx(s *ZendString, persistent int) {
	if s.DelRefcount() == 0 {
		b.Free(s)
	}
}
func ZendStringEqualContent(s1 *ZendString, s2 *ZendString) ZendBool {
	return intBool(s1.GetStr() == s2.GetStr())
}
func ZendStringEquals(s1 *ZendString, s2 *ZendString) ZendBool {
	return intBool(s1.GetStr() == s2.GetStr())
}
func ZendStringEqualsCi(s1 *ZendString, s2 *ZendString) bool {
	return zend.strCaseEquals(s1.GetStr(), s2.GetStr())
}
func ZendStringEqualsLiteralCi(str *ZendString, c string) bool {
	return zend.strCaseEquals(str.GetStr(), c)
}
func ZendStringEqualsLiteral(str *ZendString, literal string) bool {
	return str.GetStr() == literal
}
func ZendInlineHashFunc(str *byte, len_ int) zend.ZendUlong {
	var str_ = b.CastStr(str, len_)
	return b.HashStr(str_)
}
func ZendHashFunc(str *byte, len_ int) zend.ZendUlong { return ZendInlineHashFunc(str, len_) }

/**
 * Interned String 相关
 */
var IsInRequestForInternedString bool = false

func ZendInternedStringsInit() {
	ZendEmptyString = nil
	InternedStringsPermanent.Clean()
	IsInRequestForInternedString = false

	/* interned empty string */
	ZendEmptyString = InitInternedZendString("")
	for i := 0; i < 256; i++ {
		ZendOneCharString[i] = InitInternedZendString(string([]byte{byte(i)}))
	}
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

func InitInternedZendString(str string) *ZendString {
	var interned = InitInternedString(str)
	return NewZendString(interned)
}

func ZendNewInternedString(str *ZendString) *ZendString {
	return InitInternedZendString(str.GetStr())
}

func ZendStringInitInterned(str *byte, size int, permanent int) *ZendString {
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
