package types

import "github.com/heyuuu/gophp/php/lang"

// ZvalType
type ZvalType uint8

const (
	/* real types */
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

// new
func NewZvalUndef() *Zval                 { var zv Zval; zv.SetUndef(); return &zv }
func NewZvalNull() *Zval                  { var zv Zval; zv.SetNull(); return &zv }
func NewZvalFalse() *Zval                 { var zv Zval; zv.SetFalse(); return &zv }
func NewZvalTrue() *Zval                  { var zv Zval; zv.SetTrue(); return &zv }
func NewZvalBool(b bool) *Zval            { var zv Zval; zv.SetBool(b); return &zv }
func NewZvalLong(l int) *Zval             { var zv Zval; zv.SetLong(l); return &zv }
func NewZvalDouble(f float64) *Zval       { var zv Zval; zv.SetDouble(f); return &zv }
func NewZvalString(s string) *Zval        { var zv Zval; zv.SetString(s); return &zv }
func NewZvalArray(a *Array) *Zval         { var zv Zval; zv.SetArray(a); return &zv }
func NewZvalObject(o *Object) *Zval       { var zv Zval; zv.SetObject(o); return &zv }
func NewZvalResource(res *Resource) *Zval { var zv Zval; zv.SetResource(res); return &zv }

// getter
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
	default:
		panic("unreachable")
	}
}

// setter
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
