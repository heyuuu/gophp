// <<generate>>

package types

import (
	b "sik/builtin"
)

/**
 * ZendString
 */
type ZendString struct {
	str  string
	hash uint
	len_ int
	val  []byte
}

func NewZendString(str string) *ZendString {
	return &ZendString{str: str}
}

func initZendString(str string) *ZendString {
	return NewZendString(str)
}

func NewZendStringPersistent(str string) *ZendString {
	return NewZendString(str)
}

func (zs ZendString) Copy() *ZendString { return &zs }

func (zs *ZendString) GetStr() string { return zs.str }
func (zs *ZendString) GetLen() int    { return len(zs.str) }
func (zs *ZendString) GetH() uint     { return zs.hash }
func (zs *ZendString) GetHash() uint {
	if zs.hash == 0 {
		zs.hash = b.HashStr(zs.str)
	}

	return zs.hash
}

func (zs *ZendString) GetVal() []byte      { return zs.val }           // todo remove
func (zs *ZendString) GetValPtr() *byte    { return zs.val }           // todo remove
func (zs *ZendString) SetLen(value int)    { zs.str = zs.str[:value] } // todo remove
func (zs *ZendString) SetVal(value []byte) { zs.str = string(value) }  // todo remove
func (zs *ZendString) Free()               {}                          // todo remove
func (zs *ZendString) GetRefcount() uint32 { panic("implement me") }   // todo remove
func (zs *ZendString) AddRefcount() uint32 { panic("implement me") }   // todo remove
func (zs *ZendString) DelRefcount() uint32 { panic("implement me") }   // todo remove

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
