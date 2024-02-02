package types

import (
	"github.com/heyuuu/gophp/php/assert"
)

// Object
type Object struct {
	handle     uint
	ce         *Class
	data       ObjectData
	protected  bool
	properties map[string]Zval
}

func NewObject(ce *Class, handle uint, data ObjectData) *Object {
	assert.Assert(ce != nil)
	obj := initObject(ce, handle, data)

	return obj
}

func initObject(ce *Class, handle uint, data ObjectData) *Object {
	obj := &Object{
		handle:     handle,
		ce:         ce,
		data:       data,
		properties: make(map[string]Zval),
	}

	ce.PropertyTable().Each(func(propName string, propInfo *PropertyInfo) {
		obj.data.WriteProperty(ZvalString(propName), propInfo.defaultVal)
	})

	return obj
}

func (o *Object) Handle() uint      { return o.handle }
func (o *Object) Class() *Class     { return o.ce }
func (o *Object) ClassName() string { return o.ce.Name() }

// property
func (o *Object) ReadProperty(name Zval) Zval {
	return o.data.ReadProperty(name)
}
func (o *Object) WriteProperty(member Zval, value Zval) {
	o.data.WriteProperty(member, value)
}
func (o *Object) HasProperty(member Zval, hasSetExists int) bool {
	return o.data.HasProperty(member, hasSetExists)
}
func (o *Object) UnsetProperty(member Zval) {
	o.data.UnsetProperty(member)
}

// properties
func (o *Object) GetPropertiesArray() *Array {
	return o.data.GetPropertiesArray()
}
func (o *Object) GetPropertiesFor(purpose PropPurposeType) *Array {
	return o.data.GetPropertiesFor(purpose)
}

// dimension
func (o *Object) ReadDimension(offset Zval, typ int) Zval {
	return o.data.ReadDimension(offset, typ)
}
func (o *Object) WriteDimension(offset Zval, value Zval) {
	o.data.WriteDimension(offset, value)
}
func (o *Object) HasDimension(offset Zval, checkEmpty int) bool {
	return o.data.HasDimension(offset, checkEmpty)
}
func (o *Object) UnsetDimension(offset Zval) {
	o.data.UnsetDimension(offset)
}

// elements
func (o *Object) CountElements() (int, bool) {
	return o.data.CountElements()
}

// methods
func (o *Object) GetMethod(method string) *Function {
	return o.data.GetMethod(method)
}
func (o *Object) GetConstructor() *Function {
	return o.data.GetConstructor(o)
}

// mixed
func (o *Object) Cast(typ ZvalType) (Zval, bool) {
	return o.data.Cast(typ)
}

func (o *Object) CompareObjectsTo(other *Object) (int, bool) {
	// todo
	return 1, false
}

func (o *Object) CanCompare() bool {
	// todo
	return false
}

// recursive
func (o *Object) IsRecursive() bool   { return o.protected }
func (o *Object) ProtectRecursive()   { o.protected = true }
func (o *Object) UnprotectRecursive() { o.protected = false }
