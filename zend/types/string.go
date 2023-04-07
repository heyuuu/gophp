package types

/**
 * String
 */
type String struct{ str string }

func NewString(str string) *String { return &String{str: str} }

// 内部使用
func initString(str string) *String { return NewString(str) }

func (zs String) Copy() *String   { return &zs }
func (zs *String) GetStr() string { return zs.str }
func (zs *String) GetLen() int    { return len(zs.str) }

func (zs *String) GetVal() []byte   { return []byte(zs.str) }         // todo remove
func (zs *String) GetValPtr() *byte { tmp := zs.str[0]; return &tmp } // todo remove
func (zs *String) SetLen(value int) { zs.str = zs.str[:value] }       // todo remove

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

func (this *InternedStrings) GetOrInsertZendString(str string) (*String, bool) {
	s, exists := this.GetOrInsert(str)
	return NewString(s), exists
}

func (this *InternedStrings) GetZendString(str string) (*String, bool) {
	if s, exists := this.Get(str); exists {
		return NewString(s), true
	}
	return nil, false
}

func (this *InternedStrings) LookupZendString(str string) *String {
	if interned, ok := this.Get(str); ok {
		return NewString(interned)
	}
	return nil
}

func (this *InternedStrings) Clean() {
	this.cache = make(map[string]string, MIN_INTERNED_STRINGS_SIZE)
}

func (this *InternedStrings) Destroy() {
	this.Clean()
}
