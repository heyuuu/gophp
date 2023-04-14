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
	opcodes                   []zend.ZendOp
	run_time_cache__ptr       **any
	static_variables_ptr__ptr **Array
	static_variables          *Array
	vars                      []*String
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

func NewOpArray() *ZendOpArray {
	return &ZendOpArray{}
}
func CopyOpArray(array *ZendOpArray) *ZendOpArray {
	// todo 处理 copy 逻辑
	return NewOpArray()
}

func (f *ZendOpArray) init() {
	f.typ = zend.ZEND_USER_FUNCTION
}

//  ZEND_MAP_PTR_GET(f.static_variables_ptr)
func (f *ZendOpArray) GetStaticVariablesPtr() *Array {
	if uintptr(f.static_variables_ptr__ptr)&1 != 0 {
		// todo
		return nil
	} else {
		return *f.static_variables_ptr__ptr
	}
}

//  ZEND_MAP_PTR_SET(f.static_variables_ptr, ht)
func (f *ZendOpArray) SetStaticVariablesPtr(ht *Array) {
	if uintptr(f.static_variables_ptr__ptr)&1 != 0 {
		// todo
	} else {
		*f.static_variables_ptr__ptr = ht
	}
}

func (f *ZendOpArray) GetOpArray() *ZendOpArray { return f }
func (f *ZendOpArray) GetInternalFunction() *InternalFunction {
	panic("*ZendOpArray is not *InternalFunction")
}

func (f *ZendOpArray) GetCacheSize() int                           { return f.cache_size }
func (f *ZendOpArray) SetCacheSize(value int)                      { f.cache_size = value }
func (f *ZendOpArray) GetLastVar() int                             { return f.last_var }
func (f *ZendOpArray) SetLastVar(value int)                        { f.last_var = value }
func (f *ZendOpArray) GetT() uint32                                { return f.T }
func (f *ZendOpArray) SetT(value uint32)                           { f.T = value }
func (f *ZendOpArray) GetLast() uint32                             { return f.last }
func (f *ZendOpArray) SetLast(value uint32)                        { f.last = value }
func (f *ZendOpArray) GetOpcodes() []zend.ZendOp                   { return f.opcodes }
func (f *ZendOpArray) SetOpcodes(value []zend.ZendOp)              { f.opcodes = value }
func (f *ZendOpArray) GetRunTimeCachePtr() **any                   { return f.run_time_cache__ptr }
func (f *ZendOpArray) GetStaticVariables() *Array                  { return f.static_variables }
func (f *ZendOpArray) SetStaticVariables(value *Array)             { f.static_variables = value }
func (f *ZendOpArray) GetVars() []*String                          { return f.vars }
func (f *ZendOpArray) SetVars(value []*String)                     { f.vars = value }
func (f *ZendOpArray) GetRefcount() *uint32                        { return f.refcount }
func (f *ZendOpArray) SetRefcount(value *uint32)                   { f.refcount = value }
func (f *ZendOpArray) GetLastLiveRange() int                       { return f.last_live_range }
func (f *ZendOpArray) SetLastLiveRange(value int)                  { f.last_live_range = value }
func (f *ZendOpArray) GetLastTryCatch() int                        { return f.last_try_catch }
func (f *ZendOpArray) SetLastTryCatch(value int)                   { f.last_try_catch = value }
func (f *ZendOpArray) GetLiveRange() *zend.ZendLiveRange           { return f.live_range }
func (f *ZendOpArray) SetLiveRange(value *zend.ZendLiveRange)      { f.live_range = value }
func (f *ZendOpArray) GetTryCatchArray() *zend.ZendTryCatchElement { return f.try_catch_array }
func (f *ZendOpArray) SetTryCatchArray(value *zend.ZendTryCatchElement) {
	f.try_catch_array = value
}
func (f *ZendOpArray) GetFilename() *String        { return f.filename }
func (f *ZendOpArray) SetFilename(value *String)   { f.filename = value }
func (f *ZendOpArray) GetLineStart() uint32        { return f.line_start }
func (f *ZendOpArray) SetLineStart(value uint32)   { f.line_start = value }
func (f *ZendOpArray) GetLineEnd() uint32          { return f.line_end }
func (f *ZendOpArray) SetLineEnd(value uint32)     { f.line_end = value }
func (f *ZendOpArray) GetDocComment() *String      { return f.doc_comment }
func (f *ZendOpArray) SetDocComment(value *String) { f.doc_comment = value }
func (f *ZendOpArray) GetLastLiteral() int         { return f.last_literal }
func (f *ZendOpArray) SetLastLiteral(value int)    { f.last_literal = value }
func (f *ZendOpArray) GetLiterals() *Zval          { return f.literals }
func (f *ZendOpArray) SetLiterals(value *Zval)     { f.literals = value }
func (f *ZendOpArray) GetReserved() []any          { return f.reserved }
