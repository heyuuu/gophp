package types

import "github.com/heyuuu/gophp/zend"

/**
 * ZendOpArray
 */
type UserFunction = ZendOpArray
type ZendOpArray struct {
	functionHeader
	cache_size                int
	last_var                  int
	T                         uint32
	last                      uint32
	opcodes                   *zend.ZendOp
	run_time_cache__ptr       **any
	static_variables_ptr__ptr **Array
	static_variables          *Array
	vars                      **String
	refcount                  *uint32
	last_live_range           int
	last_try_catch            int
	live_range                *zend.ZendLiveRange
	try_catch_array           *zend.ZendTryCatchElement
	filename                  *String
	line_start                uint32
	line_end                  uint32
	doc_comment               *String
	last_literal              int
	literals                  *Zval
	reserved                  []any
}

var _ IFunction = (*ZendOpArray)(nil)

func (this *ZendOpArray) init() {
	this.typ = zend.ZEND_USER_FUNCTION
}

func (this *ZendOpArray) GetCacheSize() int                           { return this.cache_size }
func (this *ZendOpArray) SetCacheSize(value int)                      { this.cache_size = value }
func (this *ZendOpArray) GetLastVar() int                             { return this.last_var }
func (this *ZendOpArray) SetLastVar(value int)                        { this.last_var = value }
func (this *ZendOpArray) GetT() uint32                                { return this.T }
func (this *ZendOpArray) SetT(value uint32)                           { this.T = value }
func (this *ZendOpArray) GetLast() uint32                             { return this.last }
func (this *ZendOpArray) SetLast(value uint32)                        { this.last = value }
func (this *ZendOpArray) GetOpcodes() *zend.ZendOp                    { return this.opcodes }
func (this *ZendOpArray) SetOpcodes(value *zend.ZendOp)               { this.opcodes = value }
func (this *ZendOpArray) GetRunTimeCachePtr() **any                   { return this.run_time_cache__ptr }
func (this *ZendOpArray) GetStaticVariables() *Array                  { return this.static_variables }
func (this *ZendOpArray) SetStaticVariables(value *Array)             { this.static_variables = value }
func (this *ZendOpArray) GetVars() **String                           { return this.vars }
func (this *ZendOpArray) SetVars(value **String)                      { this.vars = value }
func (this *ZendOpArray) GetRefcount() *uint32                        { return this.refcount }
func (this *ZendOpArray) SetRefcount(value *uint32)                   { this.refcount = value }
func (this *ZendOpArray) GetLastLiveRange() int                       { return this.last_live_range }
func (this *ZendOpArray) SetLastLiveRange(value int)                  { this.last_live_range = value }
func (this *ZendOpArray) GetLastTryCatch() int                        { return this.last_try_catch }
func (this *ZendOpArray) SetLastTryCatch(value int)                   { this.last_try_catch = value }
func (this *ZendOpArray) GetLiveRange() *zend.ZendLiveRange           { return this.live_range }
func (this *ZendOpArray) SetLiveRange(value *zend.ZendLiveRange)      { this.live_range = value }
func (this *ZendOpArray) GetTryCatchArray() *zend.ZendTryCatchElement { return this.try_catch_array }
func (this *ZendOpArray) SetTryCatchArray(value *zend.ZendTryCatchElement) {
	this.try_catch_array = value
}
func (this *ZendOpArray) GetFilename() *String        { return this.filename }
func (this *ZendOpArray) SetFilename(value *String)   { this.filename = value }
func (this *ZendOpArray) GetLineStart() uint32        { return this.line_start }
func (this *ZendOpArray) SetLineStart(value uint32)   { this.line_start = value }
func (this *ZendOpArray) GetLineEnd() uint32          { return this.line_end }
func (this *ZendOpArray) SetLineEnd(value uint32)     { this.line_end = value }
func (this *ZendOpArray) GetDocComment() *String      { return this.doc_comment }
func (this *ZendOpArray) SetDocComment(value *String) { this.doc_comment = value }
func (this *ZendOpArray) GetLastLiteral() int         { return this.last_literal }
func (this *ZendOpArray) SetLastLiteral(value int)    { this.last_literal = value }
func (this *ZendOpArray) GetLiterals() *Zval          { return this.literals }
func (this *ZendOpArray) SetLiterals(value *Zval)     { this.literals = value }
func (this *ZendOpArray) GetReserved() []any          { return this.reserved }
