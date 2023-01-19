// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * PhpStripTagsFilter
 */
type PhpStripTagsFilter struct {
	allowed_tags     *byte
	allowed_tags_len int
	state            uint8
	persistent       uint8
}

func (this PhpStripTagsFilter) GetAllowedTags() *byte        { return this.allowed_tags }
func (this *PhpStripTagsFilter) SetAllowedTags(value *byte)  { this.allowed_tags = value }
func (this PhpStripTagsFilter) GetAllowedTagsLen() int       { return this.allowed_tags_len }
func (this *PhpStripTagsFilter) SetAllowedTagsLen(value int) { this.allowed_tags_len = value }
func (this PhpStripTagsFilter) GetState() uint8              { return this.state }
func (this *PhpStripTagsFilter) SetState(value uint8)        { this.state = value }
func (this PhpStripTagsFilter) GetPersistent() uint8         { return this.persistent }
func (this *PhpStripTagsFilter) SetPersistent(value uint8)   { this.persistent = value }

/**
 * PhpConv
 */
type PhpConv struct {
	convert_op PhpConvConvertFunc
	dtor       PhpConvDtorFunc
}

func (this PhpConv) GetConvertOp() PhpConvConvertFunc       { return this.convert_op }
func (this *PhpConv) SetConvertOp(value PhpConvConvertFunc) { this.convert_op = value }
func (this PhpConv) GetDtor() PhpConvDtorFunc               { return this.dtor }
func (this *PhpConv) SetDtor(value PhpConvDtorFunc)         { this.dtor = value }

/**
 * PhpConvBase64Encode
 */
type PhpConvBase64Encode struct {
	_super      PhpConv
	lbchars     *byte
	lbchars_len int
	erem_len    int
	line_ccnt   uint
	line_len    uint
	lbchars_dup int
	persistent  int
	erem        []uint8
}

func (this PhpConvBase64Encode) GetSuper() PhpConv        { return this._super }
func (this *PhpConvBase64Encode) SetSuper(value PhpConv)  { this._super = value }
func (this PhpConvBase64Encode) GetLbchars() *byte        { return this.lbchars }
func (this *PhpConvBase64Encode) SetLbchars(value *byte)  { this.lbchars = value }
func (this PhpConvBase64Encode) GetLbcharsLen() int       { return this.lbchars_len }
func (this *PhpConvBase64Encode) SetLbcharsLen(value int) { this.lbchars_len = value }
func (this PhpConvBase64Encode) GetEremLen() int          { return this.erem_len }
func (this *PhpConvBase64Encode) SetEremLen(value int)    { this.erem_len = value }
func (this PhpConvBase64Encode) GetLineCcnt() uint        { return this.line_ccnt }
func (this *PhpConvBase64Encode) SetLineCcnt(value uint)  { this.line_ccnt = value }
func (this PhpConvBase64Encode) GetLineLen() uint         { return this.line_len }
func (this *PhpConvBase64Encode) SetLineLen(value uint)   { this.line_len = value }
func (this PhpConvBase64Encode) GetLbcharsDup() int       { return this.lbchars_dup }
func (this *PhpConvBase64Encode) SetLbcharsDup(value int) { this.lbchars_dup = value }
func (this PhpConvBase64Encode) GetPersistent() int       { return this.persistent }
func (this *PhpConvBase64Encode) SetPersistent(value int) { this.persistent = value }
func (this PhpConvBase64Encode) GetErem() []uint8         { return this.erem }
func (this *PhpConvBase64Encode) SetErem(value []uint8)   { this.erem = value }

/**
 * PhpConvBase64Decode
 */
type PhpConvBase64Decode struct {
	_super     PhpConv
	urem       uint
	urem_nbits uint
	ustat      uint
	eos        int
}

func (this PhpConvBase64Decode) GetSuper() PhpConv        { return this._super }
func (this *PhpConvBase64Decode) SetSuper(value PhpConv)  { this._super = value }
func (this PhpConvBase64Decode) GetUrem() uint            { return this.urem }
func (this *PhpConvBase64Decode) SetUrem(value uint)      { this.urem = value }
func (this PhpConvBase64Decode) GetUremNbits() uint       { return this.urem_nbits }
func (this *PhpConvBase64Decode) SetUremNbits(value uint) { this.urem_nbits = value }
func (this PhpConvBase64Decode) GetUstat() uint           { return this.ustat }
func (this *PhpConvBase64Decode) SetUstat(value uint)     { this.ustat = value }
func (this PhpConvBase64Decode) GetEos() int              { return this.eos }
func (this *PhpConvBase64Decode) SetEos(value int)        { this.eos = value }

/**
 * PhpConvQprintEncode
 */
type PhpConvQprintEncode struct {
	_super      PhpConv
	lbchars     *byte
	lbchars_len int
	opts        int
	line_ccnt   uint
	line_len    uint
	lbchars_dup int
	persistent  int
	lb_ptr      uint
	lb_cnt      uint
}

func (this PhpConvQprintEncode) GetSuper() PhpConv        { return this._super }
func (this *PhpConvQprintEncode) SetSuper(value PhpConv)  { this._super = value }
func (this PhpConvQprintEncode) GetLbchars() *byte        { return this.lbchars }
func (this *PhpConvQprintEncode) SetLbchars(value *byte)  { this.lbchars = value }
func (this PhpConvQprintEncode) GetLbcharsLen() int       { return this.lbchars_len }
func (this *PhpConvQprintEncode) SetLbcharsLen(value int) { this.lbchars_len = value }
func (this PhpConvQprintEncode) GetOpts() int             { return this.opts }
func (this *PhpConvQprintEncode) SetOpts(value int)       { this.opts = value }
func (this PhpConvQprintEncode) GetLineCcnt() uint        { return this.line_ccnt }
func (this *PhpConvQprintEncode) SetLineCcnt(value uint)  { this.line_ccnt = value }
func (this PhpConvQprintEncode) GetLineLen() uint         { return this.line_len }
func (this *PhpConvQprintEncode) SetLineLen(value uint)   { this.line_len = value }
func (this PhpConvQprintEncode) GetLbcharsDup() int       { return this.lbchars_dup }
func (this *PhpConvQprintEncode) SetLbcharsDup(value int) { this.lbchars_dup = value }
func (this PhpConvQprintEncode) GetPersistent() int       { return this.persistent }
func (this *PhpConvQprintEncode) SetPersistent(value int) { this.persistent = value }
func (this PhpConvQprintEncode) GetLbPtr() uint           { return this.lb_ptr }
func (this *PhpConvQprintEncode) SetLbPtr(value uint)     { this.lb_ptr = value }
func (this PhpConvQprintEncode) GetLbCnt() uint           { return this.lb_cnt }
func (this *PhpConvQprintEncode) SetLbCnt(value uint)     { this.lb_cnt = value }

/**
 * PhpConvQprintDecode
 */
type PhpConvQprintDecode struct {
	_super      PhpConv
	lbchars     *byte
	lbchars_len int
	scan_stat   int
	next_char   uint
	lbchars_dup int
	persistent  int
	lb_ptr      uint
	lb_cnt      uint
}

func (this PhpConvQprintDecode) GetSuper() PhpConv        { return this._super }
func (this *PhpConvQprintDecode) SetSuper(value PhpConv)  { this._super = value }
func (this PhpConvQprintDecode) GetLbchars() *byte        { return this.lbchars }
func (this *PhpConvQprintDecode) SetLbchars(value *byte)  { this.lbchars = value }
func (this PhpConvQprintDecode) GetLbcharsLen() int       { return this.lbchars_len }
func (this *PhpConvQprintDecode) SetLbcharsLen(value int) { this.lbchars_len = value }
func (this PhpConvQprintDecode) GetScanStat() int         { return this.scan_stat }
func (this *PhpConvQprintDecode) SetScanStat(value int)   { this.scan_stat = value }
func (this PhpConvQprintDecode) GetNextChar() uint        { return this.next_char }
func (this *PhpConvQprintDecode) SetNextChar(value uint)  { this.next_char = value }
func (this PhpConvQprintDecode) GetLbcharsDup() int       { return this.lbchars_dup }
func (this *PhpConvQprintDecode) SetLbcharsDup(value int) { this.lbchars_dup = value }
func (this PhpConvQprintDecode) GetPersistent() int       { return this.persistent }
func (this *PhpConvQprintDecode) SetPersistent(value int) { this.persistent = value }
func (this PhpConvQprintDecode) GetLbPtr() uint           { return this.lb_ptr }
func (this *PhpConvQprintDecode) SetLbPtr(value uint)     { this.lb_ptr = value }
func (this PhpConvQprintDecode) GetLbCnt() uint           { return this.lb_cnt }
func (this *PhpConvQprintDecode) SetLbCnt(value uint)     { this.lb_cnt = value }

/**
 * PhpConvertFilter
 */
type PhpConvertFilter struct {
	cd         *PhpConv
	persistent int
	filtername *byte
	stub       []byte
	stub_len   int
}

func (this PhpConvertFilter) GetCd() *PhpConv            { return this.cd }
func (this *PhpConvertFilter) SetCd(value *PhpConv)      { this.cd = value }
func (this PhpConvertFilter) GetPersistent() int         { return this.persistent }
func (this *PhpConvertFilter) SetPersistent(value int)   { this.persistent = value }
func (this PhpConvertFilter) GetFiltername() *byte       { return this.filtername }
func (this *PhpConvertFilter) SetFiltername(value *byte) { this.filtername = value }
func (this PhpConvertFilter) GetStub() []byte            { return this.stub }
func (this *PhpConvertFilter) SetStub(value []byte)      { this.stub = value }
func (this PhpConvertFilter) GetStubLen() int            { return this.stub_len }
func (this *PhpConvertFilter) SetStubLen(value int)      { this.stub_len = value }

/**
 * PhpConsumedFilterData
 */
type PhpConsumedFilterData struct {
	consumed   int
	offset     zend.ZendOffT
	persistent uint8
}

func (this PhpConsumedFilterData) GetConsumed() int               { return this.consumed }
func (this *PhpConsumedFilterData) SetConsumed(value int)         { this.consumed = value }
func (this PhpConsumedFilterData) GetOffset() zend.ZendOffT       { return this.offset }
func (this *PhpConsumedFilterData) SetOffset(value zend.ZendOffT) { this.offset = value }
func (this PhpConsumedFilterData) GetPersistent() uint8           { return this.persistent }
func (this *PhpConsumedFilterData) SetPersistent(value uint8)     { this.persistent = value }

/**
 * PhpChunkedFilterData
 */
type PhpChunkedFilterData struct {
	chunk_size int
	state      PhpChunkedFilterState
	persistent int
}

func (this PhpChunkedFilterData) GetChunkSize() int                     { return this.chunk_size }
func (this *PhpChunkedFilterData) SetChunkSize(value int)               { this.chunk_size = value }
func (this PhpChunkedFilterData) GetState() PhpChunkedFilterState       { return this.state }
func (this *PhpChunkedFilterData) SetState(value PhpChunkedFilterState) { this.state = value }
func (this PhpChunkedFilterData) GetPersistent() int                    { return this.persistent }
func (this *PhpChunkedFilterData) SetPersistent(value int)              { this.persistent = value }
