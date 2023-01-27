// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * BrowscapKv
 */
type BrowscapKv struct {
	key   *zend.ZendString
	value *zend.ZendString
}

func (this *BrowscapKv) GetKey() *zend.ZendString        { return this.key }
func (this *BrowscapKv) SetKey(value *zend.ZendString)   { this.key = value }
func (this *BrowscapKv) GetValue() *zend.ZendString      { return this.value }
func (this *BrowscapKv) SetValue(value *zend.ZendString) { this.value = value }

/**
 * BrowscapEntry
 */
type BrowscapEntry struct {
	pattern        *zend.ZendString
	parent         *zend.ZendString
	kv_start       uint32
	kv_end         uint32
	contains_start []uint16
	contains_len   []uint8
	prefix_len     uint8
}

func (this *BrowscapEntry) GetPattern() *zend.ZendString      { return this.pattern }
func (this *BrowscapEntry) SetPattern(value *zend.ZendString) { this.pattern = value }
func (this *BrowscapEntry) GetParent() *zend.ZendString       { return this.parent }
func (this *BrowscapEntry) SetParent(value *zend.ZendString)  { this.parent = value }
func (this *BrowscapEntry) GetKvStart() uint32                { return this.kv_start }
func (this *BrowscapEntry) SetKvStart(value uint32)           { this.kv_start = value }
func (this *BrowscapEntry) GetKvEnd() uint32                  { return this.kv_end }
func (this *BrowscapEntry) SetKvEnd(value uint32)             { this.kv_end = value }
func (this *BrowscapEntry) GetContainsStart() []uint16        { return this.contains_start }
func (this *BrowscapEntry) SetContainsStart(value []uint16)   { this.contains_start = value }
func (this *BrowscapEntry) GetContainsLen() []uint8           { return this.contains_len }
func (this *BrowscapEntry) SetContainsLen(value []uint8)      { this.contains_len = value }
func (this *BrowscapEntry) GetPrefixLen() uint8               { return this.prefix_len }
func (this *BrowscapEntry) SetPrefixLen(value uint8)          { this.prefix_len = value }

/**
 * BrowserData
 */
type BrowserData struct {
	htab     *zend.HashTable
	kv       *BrowscapKv
	kv_used  uint32
	kv_size  uint32
	filename []byte
}

func (this *BrowserData) GetHtab() *zend.HashTable      { return this.htab }
func (this *BrowserData) SetHtab(value *zend.HashTable) { this.htab = value }
func (this *BrowserData) GetKv() *BrowscapKv            { return this.kv }
func (this *BrowserData) SetKv(value *BrowscapKv)       { this.kv = value }
func (this *BrowserData) GetKvUsed() uint32             { return this.kv_used }
func (this *BrowserData) SetKvUsed(value uint32)        { this.kv_used = value }
func (this *BrowserData) GetKvSize() uint32             { return this.kv_size }
func (this *BrowserData) SetKvSize(value uint32)        { this.kv_size = value }
func (this *BrowserData) GetFilename() []byte           { return this.filename }
func (this *BrowserData) SetFilename(value []byte)      { this.filename = value }

/**
 * ZendBrowscapGlobals
 */
type ZendBrowscapGlobals struct {
	activation_bdata BrowserData
}

func (this *ZendBrowscapGlobals) GetActivationBdata() BrowserData      { return this.activation_bdata }
func (this *ZendBrowscapGlobals) SetActivationBdata(value BrowserData) { this.activation_bdata = value }

/**
 * BrowscapParserCtx
 */
type BrowscapParserCtx struct {
	bdata                *BrowserData
	current_entry        *BrowscapEntry
	current_section_name *zend.ZendString
	str_interned         zend.HashTable
}

func (this *BrowscapParserCtx) GetBdata() *BrowserData               { return this.bdata }
func (this *BrowscapParserCtx) SetBdata(value *BrowserData)          { this.bdata = value }
func (this *BrowscapParserCtx) GetCurrentEntry() *BrowscapEntry      { return this.current_entry }
func (this *BrowscapParserCtx) SetCurrentEntry(value *BrowscapEntry) { this.current_entry = value }
func (this *BrowscapParserCtx) GetCurrentSectionName() *zend.ZendString {
	return this.current_section_name
}
func (this *BrowscapParserCtx) SetCurrentSectionName(value *zend.ZendString) {
	this.current_section_name = value
}
func (this *BrowscapParserCtx) GetStrInterned() zend.HashTable      { return this.str_interned }
func (this *BrowscapParserCtx) SetStrInterned(value zend.HashTable) { this.str_interned = value }
