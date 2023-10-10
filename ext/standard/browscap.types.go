package standard

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * BrowscapKv
 */
type BrowscapKv struct {
	key   string
	value string
}

/**
 * BrowscapEntry
 */
type BrowscapEntry struct {
	pattern        *types.String
	parent         *types.String
	kv_start       uint32
	kv_end         uint32
	contains_start []uint16
	contains_len   []uint8
	prefix_len     uint8
}

func (this *BrowscapEntry) GetPattern() *types.String      { return this.pattern }
func (this *BrowscapEntry) SetPattern(value *types.String) { this.pattern = value }
func (this *BrowscapEntry) GetParent() *types.String       { return this.parent }
func (this *BrowscapEntry) SetParent(value *types.String)  { this.parent = value }
func (this *BrowscapEntry) GetKvStart() uint32             { return this.kv_start }
func (this *BrowscapEntry) SetKvStart(value uint32)        { this.kv_start = value }
func (this *BrowscapEntry) GetKvEnd() uint32               { return this.kv_end }
func (this *BrowscapEntry) SetKvEnd(value uint32)          { this.kv_end = value }
func (this *BrowscapEntry) GetContainsStart() []uint16     { return this.contains_start }
func (this *BrowscapEntry) GetContainsLen() []uint8        { return this.contains_len }
func (this *BrowscapEntry) GetPrefixLen() uint8            { return this.prefix_len }
func (this *BrowscapEntry) SetPrefixLen(value uint8)       { this.prefix_len = value }

/**
 * BrowserData
 */
type BrowserData struct {
	htab    *types.Array
	kv      []BrowscapKv
	kv_used uint32
	kv_size uint32
}

func NewBrowserData(kvSize uint32) *BrowserData {
	return &BrowserData{
		htab:    types.NewArray(),
		kv:      make([]BrowscapKv, 0, kvSize),
		kv_used: 0,
		kv_size: kvSize,
	}
}
func (d *BrowserData) AddKv(key string, value string) {
	kv := BrowscapKv{key: key, value: value}
	d.kv = append(d.kv, kv)
	d.kv_used++
}

func (d *BrowserData) EachKv(handler func(key string, value string)) {
	for _, kv := range d.kv {
		handler(kv.key, kv.value)
	}
}

func (d *BrowserData) GetHtab() *types.Array { return d.htab }
func (d *BrowserData) GetKvUsed() uint32     { return uint32(len(d.kv)) }

/**
 * BrowscapParserCtx
 */
type BrowscapParserCtx struct {
	bdata                *BrowserData
	current_entry        *BrowscapEntry
	current_section_name *types.String
	strInterned          map[string]*types.String
}

func NewBrowscapParserCtx(bdata *BrowserData) *BrowscapParserCtx {
	return &BrowscapParserCtx{
		bdata:                bdata,
		current_entry:        nil,
		current_section_name: nil,
		strInterned:          make(map[string]*types.String),
	}
}
func (this *BrowscapParserCtx) GetBdata() *BrowserData               { return this.bdata }
func (this *BrowscapParserCtx) GetCurrentEntry() *BrowscapEntry      { return this.current_entry }
func (this *BrowscapParserCtx) SetCurrentEntry(value *BrowscapEntry) { this.current_entry = value }
func (this *BrowscapParserCtx) GetCurrentSectionName() *types.String {
	return this.current_section_name
}
func (this *BrowscapParserCtx) SetCurrentSectionName(value *types.String) {
	this.current_section_name = value
}
func (this *BrowscapParserCtx) GetInternedStr(str string) *types.String {
	if interned, ok := this.strInterned[str]; ok {
		return interned
	}
	interned := types.NewString(str)
	this.strInterned[str] = interned
	return interned
}
