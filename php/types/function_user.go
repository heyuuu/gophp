package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
)

const initialOpArraySize = 64

/**
 * ZendOpArray
 */
type UserFunction = ZendOpArray
type ZendOpArray struct {
	functionHeader
	cache_size                int
	T                         uint32
	last                      uint32
	opcodes                   []zend.ZendOp
	run_time_cache__ptr       *[]any
	static_variables_ptr__ptr **Array
	static_variables          *Array
	vars                      []string
	refcount                  *uint32
	last_live_range           int
	last_try_catch            int
	live_range                *zend.ZendLiveRange
	try_catch_array           *zend.ZendTryCatchElement
	filename                  *String
	line_start                uint32
	line_end                  uint32
	docComment                string // 块注释，默认值空字符表示注释不存在
	last_literal              int
	literals                  *Zval
	reserved                  []any
}

var _ IFunction = (*ZendOpArray)(nil)

func NewOpArray() *ZendOpArray {
	var refcount uint32 = 1
	return &ZendOpArray{
		functionHeader: functionHeader{
			typ: zend.ZEND_USER_FUNCTION,
		},
		opcodes:  make([]zend.ZendOp, initialOpArraySize),
		refcount: &refcount,
		// memset(op_array.GetReserved(), 0, types.ZEND_MAX_RESERVED_RESOURCES*b.SizeOf("void *"))
		reserved: make([]any, ZEND_MAX_RESERVED_RESOURCES),
	}
}
func CopyOpArray(array *ZendOpArray) *ZendOpArray {
	// todo 处理 copy 逻辑
	return NewOpArray()
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

// run_time_cache__ptr
func (f *ZendOpArray) GetRunTimeCache() []any    { return *f.run_time_cache__ptr }
func (f *ZendOpArray) HasInitRunTimeCache() bool { return f.run_time_cache__ptr != nil }
func (f *ZendOpArray) InitRunTimeCache() {
	if f.run_time_cache__ptr != nil {
		return
	}

	b.Assert(f.IsHeapRtCache())

	f.InitRunTimeCacheEx(true)
}
func (f *ZendOpArray) InitRunTimeCacheEx(heapRtCache bool) {
	if f.run_time_cache__ptr != nil {
		return
	}

	f.SetIsHeapRtCache(heapRtCache)

	var ptr []any
	if heapRtCache {
		// todo 注意此处多留了1个 void* 位置，ptr 不指向数组开端
		ptr = make([]any, 1+f.GetCacheSize())
		ptr = ptr[1:]
	} else {
		ptr = make([]any, f.GetCacheSize())
	}

	f.run_time_cache__ptr = ptr
}

func (f *ZendOpArray) InitPtr(preload bool) {
	if preload {
		f.SetIsPreloaded(true)
		zend.ZEND_MAP_PTR_NEW(f.run_time_cache)
		zend.ZEND_MAP_PTR_NEW(f.static_variables_ptr)
	} else {
		zend.ZEND_MAP_PTR_INIT(f.run_time_cache, zend.ZendArenaAlloc(zend.CG__().GetArena(), b.SizeOf("void *")))
		zend.ZEND_MAP_PTR_SET(f.run_time_cache, nil)
	}
}

func (f *ZendOpArray) InitPtr2(preload bool) {
	if preload {
		b.Assert(f.IsPreloaded())
		zend.ZEND_MAP_PTR_NEW(f.run_time_cache)
		zend.ZEND_MAP_PTR_NEW(f.static_variables_ptr)
	} else {
		zend.ZEND_MAP_PTR_INIT(f.run_time_cache, zend.ZendArenaAlloc(zend.CG__().GetArena(), b.SizeOf("void *")))
		zend.ZEND_MAP_PTR_SET(f.run_time_cache, nil)

		zend.ZEND_MAP_PTR_INIT(f.static_variables_ptr, f.GetStaticVariables())
	}
}

func (f *ZendOpArray) GetRefcount() *uint32 { return f.refcount }
func (f *ZendOpArray) TryIncRefCount() {
	if f.refcount != nil {
		*f.refcount++
	}
}

// vars
func (f *ZendOpArray) GetLastVar() int     { return len(f.vars) }
func (f *ZendOpArray) Vars() []string      { return f.vars }
func (f *ZendOpArray) GetVar(i int) string { return f.vars[i] }
func (f *ZendOpArray) AppendVar(name string) int {
	i := len(f.vars)
	f.vars = append(f.vars, name)
	return i
}
func (f *ZendOpArray) FindVar(name string) int {
	for i, varName := range f.vars {
		if varName == name {
			return i
		}
	}
	return -1
}

// fields
func (f *ZendOpArray) GetCacheSize() int              { return f.cache_size }
func (f *ZendOpArray) SetCacheSize(value int)         { f.cache_size = value }
func (f *ZendOpArray) GetT() uint32                   { return f.T }
func (f *ZendOpArray) SetT(value uint32)              { f.T = value }
func (f *ZendOpArray) GetLast() uint32                { return f.last }
func (f *ZendOpArray) SetLast(value uint32)           { f.last = value }
func (f *ZendOpArray) GetOpcodes() []zend.ZendOp      { return f.opcodes }
func (f *ZendOpArray) SetOpcodes(value []zend.ZendOp) { f.opcodes = value }

func (f *ZendOpArray) GetStaticVariables() *Array                  { return f.static_variables }
func (f *ZendOpArray) SetStaticVariables(value *Array)             { f.static_variables = value }
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

func (f *ZendOpArray) GetFilename() string        { return f.filename.GetStr() }
func (f *ZendOpArray) SetFilename(value string)   { f.filename = NewString(value) }
func (f *ZendOpArray) GetLineStart() uint32       { return f.line_start }
func (f *ZendOpArray) SetLineStart(value uint32)  { f.line_start = value }
func (f *ZendOpArray) GetLineEnd() uint32         { return f.line_end }
func (f *ZendOpArray) SetLineEnd(value uint32)    { f.line_end = value }
func (f *ZendOpArray) GetDocComment() string      { return f.docComment }
func (f *ZendOpArray) SetDocComment(value string) { f.docComment = value }
func (f *ZendOpArray) GetLastLiteral() int        { return f.last_literal }
func (f *ZendOpArray) SetLastLiteral(value int)   { f.last_literal = value }
func (f *ZendOpArray) GetLiterals() *Zval         { return f.literals }
func (f *ZendOpArray) SetLiterals(value *Zval)    { f.literals = value }
func (f *ZendOpArray) GetReserved() []any         { return f.reserved }
