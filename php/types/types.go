package types

import "github.com/heyuuu/gophp/php/lang"

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

	/* fake types */
	IsBool
	IsCallable
	IsIterable
	IsVoid
	IsNumber
)

// Zval
type Zval struct {
	v any
}

// Zval new
func NewZvalUndef() *Zval                 { var zv Zval; zv.SetUndef(); return &zv }
func NewZvalNull() *Zval                  { var zv Zval; zv.SetNull(); return &zv }
func NewZvalFalse() *Zval                 { var zv Zval; zv.SetFalse(); return &zv }
func NewZvalTrue() *Zval                  { var zv Zval; zv.SetTrue(); return &zv }
func NewZvalBool(b bool) *Zval            { var zv Zval; zv.SetBool(b); return &zv }
func NewZvalLong(l int) *Zval             { var zv Zval; zv.SetLong(l); return &zv }
func NewZvalDouble(d float64) *Zval       { var zv Zval; zv.SetDouble(d); return &zv }
func NewZvalString(s string) *Zval        { var zv Zval; zv.SetString(s); return &zv }
func NewZvalArray(a *Array) *Zval         { var zv Zval; zv.SetArray(a); return &zv }
func NewZvalObject(o *Object) *Zval       { var zv Zval; zv.SetObject(o); return &zv }
func NewZvalResource(res *Resource) *Zval { var zv Zval; zv.SetResource(res); return &zv }

// Zval setter
func (zv *Zval) SetUndef()                 { zv.v = nil }
func (zv *Zval) SetNull()                  { zv.v = IsNull }
func (zv *Zval) SetFalse()                 { zv.v = IsFalse }
func (zv *Zval) SetTrue()                  { zv.v = IsTrue }
func (zv *Zval) SetBool(b bool)            { zv.v = lang.Cond(b, IsTrue, IsFalse) }
func (zv *Zval) SetLong(l int)             { zv.v = l }
func (zv *Zval) SetDouble(d float64)       { zv.v = d }
func (zv *Zval) SetString(s string)        { zv.v = s }
func (zv *Zval) SetArray(arr *Array)       { zv.v = arr }
func (zv *Zval) SetObject(obj *Object)     { zv.v = obj }
func (zv *Zval) SetResource(res *Resource) { zv.v = res }

// Zval getter
func (zv *Zval) Type() ZvalType {
	switch v := zv.v.(type) {
	case nil:
		return IsUndef
	case ZvalType:
		return v
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
		panic("unreachable")
	}
}
func (zv *Zval) IsType(value ZvalType) bool { return zv.Type() == value }
func (zv *Zval) IsUndef() bool              { return zv.v == nil }
func (zv *Zval) IsNotUndef() bool           { return zv.v != nil }
func (zv *Zval) IsNull() bool               { return zv.v == IsNull }
func (zv *Zval) IsFalse() bool              { return zv.v == IsFalse }
func (zv *Zval) IsTrue() bool               { return zv.v == IsTrue }
func (zv *Zval) IsBool() bool               { return zv.v == IsFalse || zv.v == IsTrue }
func (zv *Zval) IsLong() bool               { _, ok := zv.v.(int); return ok }
func (zv *Zval) IsDouble() bool             { _, ok := zv.v.(float64); return ok }
func (zv *Zval) IsString() bool             { _, ok := zv.v.(string); return ok }
func (zv *Zval) IsArray() bool              { _, ok := zv.v.(*Array); return ok }
func (zv *Zval) IsObject() bool             { _, ok := zv.v.(*Object); return ok }
func (zv *Zval) IsResource() bool           { _, ok := zv.v.(*Resource); return ok }
func (zv *Zval) IsRef() bool                { _, ok := zv.v.(*Reference); return ok }

func zvalValue[T any](zv *Zval) T {
	if v, ok := zv.v.(T); ok {
		return v
	}
	panic("Get Zval value by a mismatched type")
}
func (zv *Zval) Long() int           { return zvalValue[int](zv) }
func (zv *Zval) Double() float64     { return zvalValue[float64](zv) }
func (zv *Zval) String() string      { return zvalValue[string](zv) }
func (zv *Zval) Array() *Array       { return zvalValue[*Array](zv) }
func (zv *Zval) Object() *Object     { return zvalValue[*Object](zv) }
func (zv *Zval) Resource() *Resource { return zvalValue[*Resource](zv) }
func (zv *Zval) Ref() *Reference     { return zvalValue[*Reference](zv) }

// Array
type Array struct {
	// todo
}

// Object
type Object struct {
	// todo
}

// Resource
type Resource struct {
	// todo
}

// Reference
type Reference struct {
	val Zval
	// todo
}
