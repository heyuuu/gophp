// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZSTR_VAL(str *ZendString) []byte { return str.GetVal() }

func STR_EMPTY_ALLOC() *ZendString                 { return ZendEmptyString }
func ZSTR_EMPTY_ALLOC() *ZendString                { return ZendEmptyString }
func ZSTR_CHAR(c int) *ZendString                  { return ZendOneCharString[c] }
func ZSTR_KNOWN(idx ZendKnownStringId) *ZendString { return ZendKnownStrings[idx] }

func _ZSTR_STRUCT_SIZE(len_ int) int { return _ZSTR_HEADER_SIZE + len_ + 1 }
func ZSTR_ALLOCA_ALLOC(str *ZendString, _len int, use_heap any) {
	str = ZendStringAlloc(_len, 0)
}
func ZSTR_ALLOCA_FREE(str any, use_heap any) { b.Free(str) }
func ZendStringForgetHashVal(s *ZendString) {
	s.SetH(0)
	s.DelGcFlags(IS_STR_VALID_UTF8)
}

func ZendStringSafeAlloc(n int, m int, l int, persistent int) *ZendString {
	// todo 没懂
	var ret *ZendString = (*ZendString)(SafePemalloc(n, m, ZEND_MM_ALIGNED_SIZE(_ZSTR_STRUCT_SIZE(l)), persistent))
	GC_SET_REFCOUNT(ret, 1)
	ret.GetGcTypeInfo() = IS_STRING | b.Cond(persistent != 0, IS_STR_PERSISTENT, 0)<<GC_FLAGS_SHIFT
	ret.SetH(0)
	ret.SetLen(n*m + l)
	return ret
}
func ZendStringCopy(s *ZendString) *ZendString                { return s.Copy() }
func ZendStringDup(s *ZendString, persistent int) *ZendString { return s.Dup(persistent) }
func ZendStringExtend(s *ZendString, len_ int, persistent int) *ZendString {
	ZEND_ASSERT(len_ >= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr + string(make([]byte, len_-len(oldStr)))
	s.DecGcRefcount()
	return ZendStringNew(newStr, persistent != 0)
}
func ZendStringTruncate(s *ZendString, len_ int, persistent int) *ZendString {
	ZEND_ASSERT(len_ <= s.GetLen())
	var oldStr = s.GetStr()
	var newStr = oldStr[:len_]
	s.DecGcRefcount()
	return ZendStringNew(newStr, persistent != 0)
}
func ZendStringSafeRealloc(s *ZendString, n int, m int, l int, persistent int) *ZendString {
	var ret *ZendString
	ret = ZendStringSafeAlloc(n, m, l, persistent)
	memcpy(ret.GetVal(), s.GetVal(), MIN(n*m+l, s.GetLen())+1)
	s.DecGcRefcount()
	return ret
}
func ZendStringFree(s *ZendString) {
	ZEND_ASSERT(s.GetGcRefcount() <= 1)
	b.Free(s)
}
func ZendStringEfree(s *ZendString) {
	ZEND_ASSERT(s.GetGcRefcount() <= 1)
	ZEND_ASSERT(!s.IsPersistent())
	b.Free(s)
}
func ZendStringRelease(s *ZendString) {
	if s.DecGcRefcount() == 0 {
		b.Free(s)
	}
}
func ZendStringReleaseEx(s *ZendString, persistent int) {
	if s.DecGcRefcount() == 0 {
		b.Free(s)
	}
}
func ZendStringEqualContent(s1 *ZendString, s2 *ZendString) ZendBool {
	if s1.GetStr() == s2.GetStr() {
		return 1
	} else {
		return 0
	}
}
func ZendStringEquals(s1 *ZendString, s2 *ZendString) ZendBool {
	if s1.GetStr() == s2.GetStr() {
		return 1
	} else {
		return 0
	}
}
func ZendStringEqualsCi(s1 *ZendString, s2 *ZendString) bool {
	return s1.GetLen() == s2.GetLen() && ZendBinaryStrcasecmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen()) == 0
}
func ZendStringEqualsLiteralCi(str *ZendString, c string) bool {
	return str.GetLen() == b.SizeOf("c")-1 && ZendBinaryStrcasecmp(str.GetVal(), str.GetLen(), c, b.SizeOf("c")-1) == 0
}
func ZendStringEqualsLiteral(str *ZendString, literal string) bool {
	return str.GetStr() == literal
}
func ZendInlineHashFunc(str *byte, len_ int) ZendUlong {
	return b.NewStrArg(str, uint(len_)).Hash()
}
func ZendHashFunc(str *byte, len_ int) ZendUlong { return ZendInlineHashFunc(str, len_) }
func _strDtor(zv *Zval) {
	var str *ZendString = zv.GetStr()
	b.Free(str)
}
func ZendInitInternedStringsHt(interned_strings *HashTable, permanent int) {
	interned_strings.Init(1024, nil, _strDtor, permanent)
	if permanent != 0 {
		interned_strings.RealInitMixed()
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
	InternedStringsPermanent.Destroy()
	Free(ZendKnownStrings)
	ZendKnownStrings = nil
}
func ZendInternedStringHtLookupEx(h ZendUlong, str *byte, size int, interned_strings *HashTable) *ZendString {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	nIndex = h | interned_strings.GetNTableMask()
	idx = interned_strings.Hash(nIndex)
	for idx != HT_INVALID_IDX {
		p = interned_strings.HashToBucket(idx)
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
	idx = interned_strings.Hash(nIndex)
	for idx != HT_INVALID_IDX {
		p = interned_strings.HashToBucket(idx)
		if p.GetH() == h && ZendStringEqualContent(p.GetKey(), str) != 0 {
			return p.GetKey()
		}
		idx = p.GetVal().GetNext()
	}
	return nil
}
func ZendAddInternedString(str *ZendString, interned_strings *HashTable, flags uint32) *ZendString {
	var val Zval
	str.SetGcRefcount(1)
	str.AddGcFlags(IS_STR_INTERNED | flags)
	ZVAL_INTERNED_STR(&val, str)
	interned_strings.AddNew(str, &val)
	return str
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
	if str.GetGcRefcount() > 1 {
		var h ZendUlong = str.GetH()
		str.DecGcRefcount()
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

	if str.GetGcRefcount() > 1 {
		var h ZendUlong = str.GetH()
		str.DecGcRefcount()
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
	CompilerGlobals.GetInternedStrings().Destroy()
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
