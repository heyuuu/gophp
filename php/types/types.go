package types

import (
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/perr"
)

// ZvalType
type ZvalType uint8

const (
	/* regular data types */
	IsUndef ZvalType = iota
	IsNull
	IsFalse
	IsTrue
	IsLong
	IsDouble
	IsString
	IsArray
	IsObject
	IsResource
	IsRef

	/* fake types used only for type hinting (Z_TYPE(zv) can not use them) */
	IsBool     ZvalType = 16 // _IS_BOOL
	IsCallable ZvalType = 17
	IsIterable ZvalType = 18
	IsVoid     ZvalType = 19
	IsNumber   ZvalType = 20 // _IS_NUMBER
)

// Zval value
var (
	Undef = Zval{nil}
	Null  = Zval{IsNull}
	False = Zval{false}
	True  = Zval{true}
)

func ZvalUndef() Zval                 { return Zval{nil} }
func ZvalNull() Zval                  { return Zval{IsNull} }
func ZvalFalse() Zval                 { return Zval{false} }
func ZvalTrue() Zval                  { return Zval{true} }
func ZvalBool(b bool) Zval            { return Zval{b} }
func ZvalLong(l int) Zval             { return Zval{l} }
func ZvalDouble(d float64) Zval       { return Zval{d} }
func ZvalString(s string) Zval        { return Zval{s} }
func ZvalArray(arr *Array) Zval       { assert.Assert(arr != nil); return Zval{arr} }
func ZvalObject(obj *Object) Zval     { assert.Assert(obj != nil); return Zval{obj} }
func ZvalResource(res *Resource) Zval { assert.Assert(res != nil); return Zval{res} }
func ZvalRef(ref *Reference) Zval     { assert.Assert(ref != nil); return Zval{ref} }

func ZvalArrayInit() Zval                 { return Zval{NewArray()} }
func ZvalArrayOfInt(arr []int) Zval       { return ZvalArray(NewArrayOfInt(arr)) }
func ZvalArrayOfString(arr []string) Zval { return ZvalArray(NewArrayOfString(arr)) }
func ZvalArrayOfZval(arr []Zval) Zval     { return ZvalArray(NewArrayOfZval(arr)) }

/**
 * Zval
 *
 * tips:
 * - 零值为合法的 Undef 类型
 */
type Zval struct {
	v any
}

// Zval new
func NewZvalUndef() *Zval           { return &Zval{v: nil} }
func NewZvalNull() *Zval            { return &Zval{v: IsNull} }
func NewZvalFalse() *Zval           { return &Zval{v: false} }
func NewZvalTrue() *Zval            { return &Zval{v: true} }
func NewZvalBool(b bool) *Zval      { return &Zval{v: b} }
func NewZvalLong(l int) *Zval       { return &Zval{v: l} }
func NewZvalDouble(d float64) *Zval { return &Zval{v: d} }
func NewZvalString(s string) *Zval  { return &Zval{v: s} }
func NewZvalArray(arr *Array) *Zval {
	if arr == nil {
		arr = NewArray()
	}
	return &Zval{v: arr}
}
func NewZvalObject(obj *Object) *Zval     { return &Zval{v: obj} }
func NewZvalResource(res *Resource) *Zval { return &Zval{v: res} }

// Zval setter
func (zv *Zval) SetUndef()           { zv.v = nil }
func (zv *Zval) SetNull()            { zv.v = IsNull }
func (zv *Zval) SetFalse()           { zv.v = false }
func (zv *Zval) SetTrue()            { zv.v = true }
func (zv *Zval) SetBool(b bool)      { zv.v = b }
func (zv *Zval) SetLong(l int)       { zv.v = l }
func (zv *Zval) SetDouble(d float64) { zv.v = d }
func (zv *Zval) SetString(s string)  { zv.v = s }
func (zv *Zval) SetArray(arr *Array) {
	if arr == nil {
		arr = NewArray()
	}
	zv.v = arr
}
func (zv *Zval) SetObject(obj *Object)     { zv.v = obj }
func (zv *Zval) SetResource(res *Resource) { zv.v = res }

func (zv *Zval) SetEmptyArray()                { zv.v = NewArray() }
func (zv *Zval) SetArrayOfInt(arr []int)       { zv.SetArray(NewArrayOfInt(arr)) }
func (zv *Zval) SetArrayOfString(arr []string) { zv.SetArray(NewArrayOfString(arr)) }
func (zv *Zval) SetArrayOfZval(arr []Zval)     { zv.SetArray(NewArrayOfZval(arr)) }
func (zv *Zval) SetReference(ref *Reference)   { zv.v = ref }
func (zv *Zval) SetNewEmptyRef()               { zv.SetReference(NewReference(Undef)) }
func (zv *Zval) SetNewRef(val Zval)            { zv.SetReference(NewReference(val)) }

func (zv *Zval) SetBy(val Zval) {
	zv.v = val.v
}
func (zv *Zval) SetByPtr(val *Zval) {
	if val == nil {
		*zv = Undef
	}
	zv.v = val.v
}

func (zv *Zval) Clone() *Zval {
	var tmp Zval
	tmp.SetByPtr(zv)
	return &tmp
}

func (zv Zval) CopyOrDup() Zval {
	// 除数组外，基础类型都复制了值，引用类型都复制了指针；仅数组需要做写时复制
	if zv.IsArray() {
		zv.v = zv.Array().Dup()
	}
	return zv
}

// Zval getter
func (zv Zval) Type() ZvalType {
	switch v := zv.v.(type) {
	case nil:
		return IsUndef
	case ZvalType:
		return v
	case bool:
		if v {
			return IsTrue
		} else {
			return IsFalse
		}
	case int:
		return IsLong
	case float64:
		return IsDouble
	case string:
		return IsString
	case *Array:
		return IsArray
	case *Object:
		return IsObject
	case *Resource:
		return IsResource
	case *Reference:
		return IsRef
	default:
		panic(perr.Unreachable())
	}
}

func (zv Zval) IsType(typ ZvalType) bool { return zv.Type() == typ }
func (zv Zval) IsUndef() bool            { return zv.v == nil }
func (zv Zval) IsNotUndef() bool         { return zv.v != nil }
func (zv Zval) IsNull() bool             { return zv.v == IsNull }
func (zv Zval) IsFalse() bool            { return zv.v == false }
func (zv Zval) IsTrue() bool             { return zv.v == true }
func (zv Zval) IsBool() bool             { _, ok := zv.v.(bool); return ok }
func (zv Zval) IsLong() bool             { _, ok := zv.v.(int); return ok }
func (zv Zval) IsDouble() bool           { _, ok := zv.v.(float64); return ok }
func (zv Zval) IsString() bool           { _, ok := zv.v.(string); return ok }
func (zv Zval) IsArray() bool            { _, ok := zv.v.(*Array); return ok }
func (zv Zval) IsObject() bool           { _, ok := zv.v.(*Object); return ok }
func (zv Zval) IsResource() bool         { _, ok := zv.v.(*Resource); return ok }
func (zv Zval) IsRef() bool              { _, ok := zv.v.(*Reference); return ok }

// 返回是否为 undef、null、false，用于快速类型判断
func (zv Zval) IsSignFalse() bool { return zv.Type() <= IsFalse }

// 返回是否为 undef、null、false 或 true，用于快速类型判断
func (zv Zval) IsSignType() bool { return zv.Type() <= IsTrue }

func zvalValue[T any](zv Zval) T {
	if v, ok := zv.v.(T); ok {
		return v
	}
	panic("Get Zval value by a mismatched type")
}
func (zv Zval) Bool() bool          { return zvalValue[bool](zv) }
func (zv Zval) Long() int           { return zvalValue[int](zv) }
func (zv Zval) Double() float64     { return zvalValue[float64](zv) }
func (zv Zval) String() string      { return zvalValue[string](zv) }
func (zv Zval) Array() *Array       { return zvalValue[*Array](zv) }
func (zv Zval) Object() *Object     { return zvalValue[*Object](zv) }
func (zv Zval) Resource() *Resource { return zvalValue[*Resource](zv) }
func (zv Zval) Ref() *Reference     { return zvalValue[*Reference](zv) }
func (zv Zval) DeRef() Zval {
	if ref, ok := zv.v.(*Reference); ok {
		return ref.Val()
	}
	return zv
}

// fast property
func (zv Zval) ResourceHandle() int { return zv.Resource().Handle() }
func (zv Zval) RefVal() Zval        { return zv.Ref().Val() }


// -- Resource 相关

type ResourceType string

func (n ResourceType) Name() string {
	return string(n)
}

type ResourceCloser interface {
	ResourceClose() int // 返回运行的进程的终止状态。发生错误时会返回 -1。主要用于 `pclose` 方法.
}

type Resource struct {
	handle int          `get:""`
	ptr    any          `get:""`
	typ    ResourceType `get:"Type"`
}

func NewResource(handle int, typ ResourceType, ptr any) *Resource {
	return &Resource{handle: handle, typ: typ, ptr: ptr}
}

func (res *Resource) TypeName() string { return string(res.typ) }
func (res *Resource) Closed() bool     { return res.typ == "" }
func (res *Resource) Close() int {
	code := 0
	if res.typ != "" {
		// 先置空 ptr 后处理，避免循环调用
		ptr := res.ptr
		res.typ, res.ptr = "", nil

		// 处理定义了接口的指针
		if closer, ok := ptr.(ResourceCloser); ok {
			code = closer.ResourceClose()
		}
	}
	return code
}

// Reference
type Reference struct {
	val Zval
	// todo
}

func NewReference(val Zval) *Reference {
	var ref = &Reference{val: val}
	return ref
}

func (ref *Reference) Val() Zval {
	return ref.val
}
func (ref *Reference) SetVal(v Zval) {
	// todo check type
	ref.val = v.DeRef()
}

// RefZval
type RefZval interface {
	Val() Zval
	SetVal(v Zval)
}
