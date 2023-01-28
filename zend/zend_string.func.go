// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZSTR_VAL(zstr *ZendString) []byte { return zstr.GetVal() }

var ZendEmptyString *ZendString = nil

func STR_EMPTY_ALLOC() *ZendString  { return ZendEmptyString }
func ZSTR_EMPTY_ALLOC() *ZendString { return ZendEmptyString }

func ZSTR_CHAR(c int) *ZendString                  { return ZendOneCharString[c] }
func ZSTR_KNOWN(idx ZendKnownStringId) *ZendString { return ZendKnownStrings[idx] }

func _ZSTR_STRUCT_SIZE(len_ int) int { return _ZSTR_HEADER_SIZE + len_ + 1 }

func ZSTR_ALLOCA_ALLOC(str *ZendString, _len int, use_heap __auto__) {
	str = (*ZendString)(DoAlloca(ZEND_MM_ALIGNED_SIZE_EX(_ZSTR_STRUCT_SIZE(_len), 8), use_heap))
	str.SetRefcount(1)
	str.GetGcTypeInfo() = IS_STRING
	str.SetH(0)
	str.SetLen(_len)
}

func ZSTR_ALLOCA_FREE(str any, use_heap any) { b.Free(str) }
func ZendStringForgetHashVal(s *ZendString) {
	s.SetH(0)
	s.DelGcFlags(IS_STR_VALID_UTF8)
}

func ZendStringAlloc(len_ int, persistent int) *ZendString {
	var str_ = b.EmptyString(len_)
	return ZendStringNew(str_, persistent != 0)
}
func ZendStringSafeAlloc(n int, m int, l int, persistent int) *ZendString {
	// todo 不太明白参数作用，仅从纯代码功能重构
	var len_ = n*m + l
	return ZendStringAlloc(len_, persistent)
}
func ZendStringInit(str *byte, len_ int, persistent int) *ZendString {
	var str_ = b.CastStr(str, len_)
	return ZendStringNew(str_, persistent != 0)
}
func ZendStringExtend(s *ZendString, len_ int, persistent int) *ZendString {
	ZEND_ASSERT(len_ >= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr + b.EmptyString(len_-len(oldStr))
	s.DelRefcount()
	return ZendStringNew(newStr, persistent != 0)
}
func ZendStringTruncate(s *ZendString, len_ int, persistent int) *ZendString {
	ZEND_ASSERT(len_ <= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr[:len_]
	s.DelRefcount()
	return ZendStringNew(newStr, persistent != 0)
}
func ZendStringSafeRealloc(s *ZendString, n int, m int, l int, persistent int) *ZendString {
	var ret *ZendString
	if s.GetRefcount() == 1 {
		ret = (*ZendString)(SafePerealloc(s, n, m, ZEND_MM_ALIGNED_SIZE(_ZSTR_STRUCT_SIZE(l)), persistent))
		ret.SetLen(n*m + l)
		ZendStringForgetHashVal(ret)
		return ret
	}
	ret = ZendStringSafeAlloc(n, m, l, persistent)
	memcpy(ret.GetVal(), s.GetVal(), MIN(n*m+l, s.GetLen())+1)
	s.DelRefcount()
	return ret
}
func ZendStringFree(s *ZendString) {
	ZEND_ASSERT(s.GetRefcount() <= 1)
	b.Free(s)
}
func ZendStringEfree(s *ZendString) {
	ZEND_ASSERT(s.GetRefcount() <= 1)
	ZEND_ASSERT((s.GetGcFlags() & IS_STR_PERSISTENT) == 0)
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
	return strCaseEquals(s1.GetStr(), s2.GetStr())
}
func ZendStringEqualsLiteralCi(str *ZendString, c string) bool {
	return strCaseEquals(str.GetStr(), c)
}
func ZendStringEqualsLiteral(str *ZendString, literal string) bool {
	return str.GetStr() == literal
}
func ZendInlineHashFunc(str *byte, len_ int) ZendUlong {
	var str_ = b.CastStr(str, len_)
	return b.HashStr(str_)
}
func ZendHashFunc(str *byte, len_ int) ZendUlong { return ZendInlineHashFunc(str, len_) }
func _strDtor(zv *Zval) {
	var str *ZendString = zv.GetStr()
	Pefree(str, str.GetGcFlags()&IS_STR_PERSISTENT)
}
func ZendInitInternedStringsHt(interned_strings *HashTable, permanent int) {
	ZendHashInit(interned_strings, 1024, nil, _strDtor, permanent)
	if permanent != 0 {
		ZendHashRealInitMixed(interned_strings)
	}
}
func ZendInternedStringsInit() {
	var s []byte
	var i uint
	var str *ZendString
	InternedStringRequestHandler = ZendNewInternedStringRequest
	InternedStringInitRequestHandler = ZendStringInitInternedRequest
	ZendEmptyString = nil
	ZendKnownStrings = nil
	ZendInitInternedStringsHt(&InternedStringsPermanent, 1)
	ZendNewInternedString = ZendNewInternedStringPermanent
	ZendStringInitInterned = ZendStringInitInternedPermanent

	/* interned empty string */

	str = ZendStringAlloc(b.SizeOf("\"\"")-1, 1)
	str.GetVal()[0] = '0'
	ZendEmptyString = ZendNewInternedStringPermanent(str)
	s[1] = 0
	for i = 0; i < 256; i++ {
		s[0] = i
		ZendOneCharString[i] = ZendNewInternedStringPermanent(ZendStringInit(s, 1, 1))
	}

	/* known strings */

	ZendKnownStrings = Pemalloc(b.SizeOf("zend_string *")*(b.SizeOf("known_strings")/b.SizeOf("known_strings [ 0 ]")-1), 1)
	for i = 0; i < b.SizeOf("known_strings")/b.SizeOf("known_strings [ 0 ]")-1; i++ {
		str = ZendStringInit(KnownStrings[i], strlen(KnownStrings[i]), 1)
		ZendKnownStrings[i] = ZendNewInternedStringPermanent(str)
	}
}
func ZendInternedStringsDtor() {
	ZendHashDestroy(&InternedStringsPermanent)
	Free(ZendKnownStrings)
	ZendKnownStrings = nil
}
func ZendInternedStringHtLookupEx(h ZendUlong, str *byte, size int, interned_strings *HashTable) *ZendString {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	nIndex = h | interned_strings.GetNTableMask()
	idx = HT_HASH(interned_strings, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(interned_strings, idx)
		if p.GetH() == h && p.GetKey().GetLen() == size {
			if !(memcmp(p.GetKey().GetVal(), str, size)) {
				return p.GetKey()
			}
		}
		idx = p.GetVal().GetNext()
	}
	return nil
}
func ZendInternedStringHtLookup(str *ZendString, interned_strings *HashTable) *ZendString {
	var h ZendUlong = str.GetH()
	var nIndex uint32
	var idx uint32
	var p *Bucket
	nIndex = h | interned_strings.GetNTableMask()
	idx = HT_HASH(interned_strings, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(interned_strings, idx)
		if p.GetH() == h && ZendStringEqualContent(p.GetKey(), str) != 0 {
			return p.GetKey()
		}
		idx = p.GetVal().GetNext()
	}
	return nil
}
func ZendAddInternedString(str *ZendString, interned_strings *HashTable, flags uint32) *ZendString {
	var val Zval
	str.SetRefcount(1)
	str.AddGcFlags(IS_STR_INTERNED | flags)
	ZVAL_INTERNED_STR(&val, str)
	ZendHashAddNew(interned_strings, str, &val)
	return str
}
func ZendInternedStringFindPermanent(str *ZendString) *ZendString {
	str.GetHash()
	return ZendInternedStringHtLookup(str, &InternedStringsPermanent)
}
func ZendNewInternedStringPermanent(str *ZendString) *ZendString {
	var ret *ZendString

	str.GetHash()
	ret = ZendInternedStringHtLookup(str, &InternedStringsPermanent)
	if ret != nil {
		ZendStringRelease(str)
		return ret
	}
	ZEND_ASSERT((str.GetGcFlags() & GC_PERSISTENT) != 0)
	if str.GetRefcount() > 1 {
		var h ZendUlong = str.GetH()
		str.DelRefcount()
		str = ZendStringInit(str.GetVal(), str.GetLen(), 1)
		str.SetH(h)
	}
	return ZendAddInternedString(str, &InternedStringsPermanent, IS_STR_PERMANENT)
}
func ZendNewInternedStringRequest(str *ZendString) *ZendString {
	var ret *ZendString

	str.GetHash()

	/* Check for permanent strings, the table is readonly at this point. */

	ret = ZendInternedStringHtLookup(str, &InternedStringsPermanent)
	if ret != nil {
		ZendStringRelease(str)
		return ret
	}
	ret = ZendInternedStringHtLookup(str, &(CompilerGlobals.GetInternedStrings()))
	if ret != nil {
		ZendStringRelease(str)
		return ret
	}

	/* Create a short living interned, freed after the request. */

	if str.GetRefcount() > 1 {
		var h ZendUlong = str.GetH()
		str.DelRefcount()
		str = ZendStringInit(str.GetVal(), str.GetLen(), 0)
		str.SetH(h)
	}
	ret = ZendAddInternedString(str, &(CompilerGlobals.GetInternedStrings()), 0)
	return ret
}
func ZendStringInitInternedPermanent(str *byte, size int, permanent int) *ZendString {
	var ret *ZendString
	var h ZendUlong = ZendInlineHashFunc(str, size)
	ret = ZendInternedStringHtLookupEx(h, str, size, &InternedStringsPermanent)
	if ret != nil {
		return ret
	}
	ZEND_ASSERT(permanent != 0)
	ret = ZendStringInit(str, size, permanent)
	ret.SetH(h)
	return ZendAddInternedString(ret, &InternedStringsPermanent, IS_STR_PERMANENT)
}
func ZendStringInitInternedRequest(str *byte, size int, permanent int) *ZendString {
	var ret *ZendString
	var h ZendUlong = ZendInlineHashFunc(str, size)

	/* Check for permanent strings, the table is readonly at this point. */

	ret = ZendInternedStringHtLookupEx(h, str, size, &InternedStringsPermanent)
	if ret != nil {
		return ret
	}
	ret = ZendInternedStringHtLookupEx(h, str, size, &(CompilerGlobals.GetInternedStrings()))
	if ret != nil {
		return ret
	}
	ret = ZendStringInit(str, size, permanent)
	ret.SetH(h)

	/* Create a short living interned, freed after the request. */

	return ZendAddInternedString(ret, &(CompilerGlobals.GetInternedStrings()), 0)

	/* Create a short living interned, freed after the request. */
}
func ZendInternedStringsActivate() {
	ZendInitInternedStringsHt(&(CompilerGlobals.GetInternedStrings()), 0)
}
func ZendInternedStringsDeactivate() {
	ZendHashDestroy(&(CompilerGlobals.GetInternedStrings()))
}
func ZendInternedStringsSwitchStorage(request ZendBool) {
	if request != 0 {
		ZendNewInternedString = InternedStringRequestHandler
		ZendStringInitInterned = InternedStringInitRequestHandler
	} else {
		ZendNewInternedString = ZendNewInternedStringPermanent
		ZendStringInitInterned = ZendStringInitInternedPermanent
	}
}
