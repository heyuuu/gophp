package types

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

/**
 * Zval
 *
 * tips:
 * - 零值为合法的 Undef 类型，无需 SetUndef() 初始化
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

// Zval getter
func (zv *Zval) Type() ZvalType {
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
	case *Ref:
		return IsRef
	default:
		panic("unreachable")
	}
}

func (zv *Zval) IsType(typ ZvalType) bool { return zv.Type() == typ }
func (zv *Zval) IsUndef() bool            { return zv.v == nil }
func (zv *Zval) IsNotUndef() bool         { return zv.v != nil }
func (zv *Zval) IsNull() bool             { return zv.v == IsNull }
func (zv *Zval) IsFalse() bool            { return zv.v == false }
func (zv *Zval) IsTrue() bool             { return zv.v == true }
func (zv *Zval) IsBool() bool             { _, ok := zv.v.(bool); return ok }
func (zv *Zval) IsLong() bool             { _, ok := zv.v.(int); return ok }
func (zv *Zval) IsDouble() bool           { _, ok := zv.v.(float64); return ok }
func (zv *Zval) IsString() bool           { _, ok := zv.v.(string); return ok }
func (zv *Zval) IsArray() bool            { _, ok := zv.v.(*Array); return ok }
func (zv *Zval) IsObject() bool           { _, ok := zv.v.(*Object); return ok }
func (zv *Zval) IsResource() bool         { _, ok := zv.v.(*Resource); return ok }
func (zv *Zval) IsRef() bool              { _, ok := zv.v.(*Ref); return ok }

// 返回是否为 undef、null、false，用于快速类型判断
func (zv *Zval) IsSignFalse() bool { return zv.Type() <= IsFalse }

// 返回是否为 undef、null、false 或 true，用于快速类型判断
func (zv *Zval) IsSignType() bool { return zv.Type() <= IsTrue }

func zvalValue[T any](zv *Zval) T {
	if v, ok := zv.v.(T); ok {
		return v
	}
	panic("Get Zval value by a mismatched type")
}
func (zv *Zval) Bool() bool          { return zvalValue[bool](zv) }
func (zv *Zval) Long() int           { return zvalValue[int](zv) }
func (zv *Zval) Double() float64     { return zvalValue[float64](zv) }
func (zv *Zval) String() string      { return zvalValue[string](zv) }
func (zv *Zval) Array() *Array       { return zvalValue[*Array](zv) }
func (zv *Zval) Object() *Object     { return zvalValue[*Object](zv) }
func (zv *Zval) Resource() *Resource { return zvalValue[*Resource](zv) }
func (zv *Zval) Ref() *Ref           { return zvalValue[*Ref](zv) }
func (zv *Zval) DeRef() *Zval {
	if zv == nil {
		return nil
	}
	if ref, ok := zv.v.(*Ref); ok {
		return ref.Val()
	}
	return zv
}

// fast property
func (zv *Zval) ResourceHandle() int { return zv.Resource().Handle() }
func (zv *Zval) ResourceType() int   { return zv.Resource().Type() }
func (zv *Zval) RefVal() *Zval       { return zv.Ref().Val() }

// Resource
type Resource struct {
	handle int
	typ    int
	ptr    any
}

func NewResource(handle int, typ int, ptr any) *Resource {
	return &Resource{handle: handle, typ: typ, ptr: ptr}
}

func (res *Resource) Handle() int { return res.handle }
func (res *Resource) Type() int   { return res.typ }
func (res *Resource) Ptr() any    { return res.ptr }

// Reference
type Ref = Reference
type Reference struct {
	val Zval
	// todo
}

func (ref *Reference) Val() *Zval { return &ref.val }
