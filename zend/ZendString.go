// <<generate>>

package zend

import b "sik/builtin"

/**
 * ZendString
 */
type ZendString struct {
	ZendRefcounted
	h    ZendUlong
	len_ int
	val  []byte
}

var _ IRefcounted = &ZendString{}

func NewZendString(str string) *ZendString {
	var zs = &ZendString{val: []byte(str), len_: len(str)}

	zs.SetRefcount(1)
	zs.SetGcTypeInfo(IS_STRING)

	return zs
}

func NewZendStringPersistent(str string, persistent bool) *ZendString {
	var zs = NewZendString(str)
	if persistent {
		zs.AddGcFlags(IS_STR_PERSISTENT)
	}
	return zs
}

func (this *ZendString) GetH() ZendUlong      { return this.h }
func (this *ZendString) SetH(value ZendUlong) { this.h = value }
func (this *ZendString) GetLen() int          { return this.len_ }
func (this *ZendString) SetLen(value int)     { this.len_ = value }
func (this *ZendString) GetVal() []byte       { return this.val }
func (this *ZendString) SetVal(value []byte)  { this.val = value }

func (this *ZendString) GetStr() string {
	return string(this.val[:this.len_])
}

func (this *ZendString) GetHash() ZendUlong {
	if this.h == 0 {
		this.h = b.HashBytes(this.val[:this.len_])
	}

	return this.h
}

func (this *ZendString) Copy() *ZendString {
	this.AddRefcount()
	return this
}

func (this *ZendString) Dup(persistent int) *ZendString {
	return NewZendStringPersistent(this.GetStr(), persistent != 0)
}

func (this *ZendString) Free() {
	b.Free(this)
}

/**
 * InternedStrings
 * 内部字符串缓存，通过将相等字符串替换为内部字符串的方式减少字符串内存占用
 */
type InternedStrings struct {
	cache map[string]string
}

const MIN_INTERNED_STRINGS_SIZE = 1024

func NewInternedStrings() *InternedStrings {
	return &InternedStrings{
		cache: make(map[string]string, MIN_INTERNED_STRINGS_SIZE),
	}
}

func (this *InternedStrings) GetOrInsert(str string) (string, bool) {
	if interned, ok := this.cache[str]; ok {
		return interned, true
	} else {
		this.cache[str] = str
		return str, false
	}
}

func (this *InternedStrings) Get(str string) (string, bool) {
	if interned, ok := this.cache[str]; ok {
		return interned, true
	}
	return "", false
}

func (this *InternedStrings) GetOrInsertZendString(str string) (*ZendString, bool) {
	s, exists := this.GetOrInsert(str)
	return NewZendString(s), exists
}

func (this *InternedStrings) GetZendString(str string) (*ZendString, bool) {
	if s, exists := this.Get(str); exists {
		return NewZendString(s), true
	}
	return nil, false
}

func (this *InternedStrings) LookupZendString(str string) *ZendString {
	if interned, ok := this.Get(str); ok {
		return NewZendString(interned)
	}
	return nil
}

func (this *InternedStrings) Clean() {
	this.cache = make(map[string]string, MIN_INTERNED_STRINGS_SIZE)
}

func (this *InternedStrings) Destroy() {
	this.Clean()
}
