// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * UrlAdaptStateExT
 */
type UrlAdaptStateExT struct {
	tag         zend.SmartStr
	arg         zend.SmartStr
	val         zend.SmartStr
	buf         zend.SmartStr
	result      zend.SmartStr
	form_app    zend.SmartStr
	url_app     zend.SmartStr
	active      int
	lookup_data *byte
	state       int
	type_       int
	attr_val    zend.SmartStr
	tag_type    int
	attr_type   int
	tags        *zend.HashTable
}

func (this *UrlAdaptStateExT) GetTag() zend.SmartStr          { return this.tag }
func (this *UrlAdaptStateExT) SetTag(value zend.SmartStr)     { this.tag = value }
func (this *UrlAdaptStateExT) GetArg() zend.SmartStr          { return this.arg }
func (this *UrlAdaptStateExT) SetArg(value zend.SmartStr)     { this.arg = value }
func (this *UrlAdaptStateExT) GetVal() zend.SmartStr          { return this.val }
func (this *UrlAdaptStateExT) SetVal(value zend.SmartStr)     { this.val = value }
func (this *UrlAdaptStateExT) GetBuf() zend.SmartStr          { return this.buf }
func (this *UrlAdaptStateExT) SetBuf(value zend.SmartStr)     { this.buf = value }
func (this *UrlAdaptStateExT) GetResult() zend.SmartStr       { return this.result }
func (this *UrlAdaptStateExT) SetResult(value zend.SmartStr)  { this.result = value }
func (this *UrlAdaptStateExT) GetFormApp() zend.SmartStr      { return this.form_app }
func (this *UrlAdaptStateExT) SetFormApp(value zend.SmartStr) { this.form_app = value }
func (this *UrlAdaptStateExT) GetUrlApp() zend.SmartStr       { return this.url_app }
func (this *UrlAdaptStateExT) SetUrlApp(value zend.SmartStr)  { this.url_app = value }
func (this *UrlAdaptStateExT) GetActive() int                 { return this.active }
func (this *UrlAdaptStateExT) SetActive(value int)            { this.active = value }
func (this *UrlAdaptStateExT) GetLookupData() *byte           { return this.lookup_data }
func (this *UrlAdaptStateExT) SetLookupData(value *byte)      { this.lookup_data = value }
func (this *UrlAdaptStateExT) GetState() int                  { return this.state }
func (this *UrlAdaptStateExT) SetState(value int)             { this.state = value }
func (this *UrlAdaptStateExT) GetType() int                   { return this.type_ }
func (this *UrlAdaptStateExT) SetType(value int)              { this.type_ = value }
func (this *UrlAdaptStateExT) GetAttrVal() zend.SmartStr      { return this.attr_val }
func (this *UrlAdaptStateExT) SetAttrVal(value zend.SmartStr) { this.attr_val = value }
func (this *UrlAdaptStateExT) GetTagType() int                { return this.tag_type }
func (this *UrlAdaptStateExT) SetTagType(value int)           { this.tag_type = value }
func (this *UrlAdaptStateExT) GetAttrType() int               { return this.attr_type }
func (this *UrlAdaptStateExT) SetAttrType(value int)          { this.attr_type = value }
func (this *UrlAdaptStateExT) GetTags() *zend.HashTable       { return this.tags }
func (this *UrlAdaptStateExT) SetTags(value *zend.HashTable)  { this.tags = value }
