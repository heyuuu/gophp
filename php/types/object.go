package types

import (
	"github.com/heyuuu/gophp/php/assert"
)

const (
	guardInGet = 1 << iota
	guardInSet
	guardInUnset
	guardInIsset
)

// PropertyGuard
type PropertyGuard uint8

func (guard *PropertyGuard) InGet() bool        { return *guard&guardInGet != 0 }
func (guard *PropertyGuard) InSet() bool        { return *guard&guardInSet != 0 }
func (guard *PropertyGuard) InIsset() bool      { return *guard&guardInIsset != 0 }
func (guard *PropertyGuard) InUnset() bool      { return *guard&guardInUnset != 0 }
func (guard *PropertyGuard) MarkInGet(v bool)   { guard.mark(guardInGet, v) }
func (guard *PropertyGuard) MarkInSet(v bool)   { guard.mark(guardInSet, v) }
func (guard *PropertyGuard) MarkInIsset(v bool) { guard.mark(guardInIsset, v) }
func (guard *PropertyGuard) MarkInUnset(v bool) { guard.mark(guardInUnset, v) }
func (guard *PropertyGuard) mark(sign PropertyGuard, v bool) {
	if v {
		*guard |= sign
	} else {
		*guard &^= sign
	}
}

// ClosureData
type ClosureData struct {
	ce  *Class    `prop:""`
	fn  *Function `prop:""`
	obj *Object   `prop:""`
}

// IObject 内部对象接口，用于适配各种对象实现
type IObject interface {
	Class() *Class

	// clone
	CanClone() bool
	Clone() *Object

	// property
	ReadProperty(member Zval, typ int) Zval
	WriteProperty(member Zval, value Zval)
	HasProperty(member Zval, hasSetExists int) bool
	UnsetProperty(member Zval)
	GetPropertyPtr(member Zval, typ int) *Zval

	// properties
	GetPropertiesArray() *Array
	GetPropertiesFor(purpose PropPurposeType) *Array

	// dimension
	ReadDimension(offset Zval, typ int) Zval
	WriteDimension(offset Zval, value Zval)
	HasDimension(offset Zval, checkEmpty int) bool
	UnsetDimension(offset Zval)

	// elements
	CountElements() (int, bool)

	// methods
	GetMethod(method string) *Function
	GetConstructor(object *Object) *Function

	// mixed
	Cast(typ ZvalType) (Zval, bool)
	CompareObjectTo(another *Object) (int, bool)
	CompareTo(another Zval) (int, bool)
	GetClosure() (*ClosureData, bool)
}

// Object
type Object struct {
	handle         uint
	intern         IObject
	protected      bool
	propertyGuards map[string]*PropertyGuard // 属性 guard
}

func NewObject(handle uint, intern IObject) *Object {
	obj := initObject(handle, intern)
	obj.propertiesInit()
	return obj
}

func initObject(handle uint, intern IObject) *Object {
	return &Object{
		handle: handle,
		intern: intern,
	}
}

func (o *Object) propertiesInit() {
	ce := o.intern.Class()

	ce.PropertyTable().Each(func(propName string, propInfo *PropertyInfo) {
		o.intern.WriteProperty(ZvalString(propName), propInfo.defaultVal)
	})
}

// guard
func (o *Object) Guard(member string) *PropertyGuard {
	assert.Assert(o.Class().IsUseGuards())
	// php 原版设计，将 guard 附加在 propertiesTable 最后一位; gophp 中，将其抽出作为单独 map
	if o.propertyGuards == nil {
		var tmp PropertyGuard
		o.propertyGuards = map[string]*PropertyGuard{
			member: &tmp,
		}
		return &tmp
	}

	if guard := o.propertyGuards[member]; guard != nil {
		return guard
	}

	var tmp PropertyGuard
	o.propertyGuards[member] = &tmp
	return &tmp
}

func (o *Object) Handle() uint      { return o.handle }
func (o *Object) Class() *Class     { return o.intern.Class() }
func (o *Object) ClassName() string { return o.intern.Class().Name() }
func (o *Object) Data() IObject     { return o.intern }

// clone
func (o *Object) CanClone() bool { return o.intern.CanClone() }
func (o *Object) Clone() *Object {
	assert.Assert(o.intern.CanClone())
	return o.intern.Clone()
}

// property
func (o *Object) ReadProperty(member Zval, typ int) Zval {
	return o.intern.ReadProperty(member, typ)
}
func (o *Object) WriteProperty(member Zval, value Zval) {
	o.intern.WriteProperty(member, value)
}
func (o *Object) HasProperty(member Zval, hasSetExists int) bool {
	return o.intern.HasProperty(member, hasSetExists)
}
func (o *Object) UnsetProperty(member Zval) {
	o.intern.UnsetProperty(member)
}
func (o *Object) GetPropertyPtr(member Zval, typ int) *Zval {
	return o.intern.GetPropertyPtr(member, typ)
}

// property
func (o *Object) ReadKeyProperty(member string, typ int) Zval {
	return o.intern.ReadProperty(ZvalString(member), typ)
}
func (o *Object) WriteKeyProperty(member string, value Zval) {
	o.intern.WriteProperty(ZvalString(member), value)
}
func (o *Object) HasKeyProperty(member string, hasSetExists int) bool {
	return o.intern.HasProperty(ZvalString(member), hasSetExists)
}
func (o *Object) UnsetKeyProperty(member string) {
	o.intern.UnsetProperty(ZvalString(member))
}

// properties
func (o *Object) GetPropertiesArray() *Array {
	return o.intern.GetPropertiesArray()
}
func (o *Object) GetPropertiesFor(purpose PropPurposeType) *Array {
	return o.intern.GetPropertiesFor(purpose)
}

// dimension
func (o *Object) ReadDimension(offset Zval, typ int) Zval {
	return o.intern.ReadDimension(offset, typ)
}
func (o *Object) WriteDimension(offset Zval, value Zval) {
	o.intern.WriteDimension(offset, value)
}
func (o *Object) HasDimension(offset Zval, checkEmpty int) bool {
	return o.intern.HasDimension(offset, checkEmpty)
}
func (o *Object) UnsetDimension(offset Zval) {
	o.intern.UnsetDimension(offset)
}

// elements
func (o *Object) CountElements() (int, bool) {
	return o.intern.CountElements()
}

// methods
func (o *Object) GetMethod(method string) *Function {
	return o.intern.GetMethod(method)
}
func (o *Object) GetConstructor() *Function {
	return o.intern.GetConstructor(o)
}

// cast
func (o *Object) Cast(typ ZvalType) (Zval, bool) {
	if typ == IsTrue || typ == IsFalse {
		typ = IsBool
	}
	return o.intern.Cast(typ)
}

// mixed
func (o *Object) CompareObjectTo(another *Object) (int, bool) {
	return o.intern.CompareObjectTo(another)
}
func (o *Object) CompareTo(another Zval) (int, bool) {
	return o.intern.CompareTo(another)
}
func (o *Object) GetClosure() (*ClosureData, bool) {
	return o.intern.GetClosure()
}

// recursive
func (o *Object) IsRecursive() bool   { return o.protected }
func (o *Object) ProtectRecursive()   { o.protected = true }
func (o *Object) UnprotectRecursive() { o.protected = false }
