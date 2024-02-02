package types

import (
	"github.com/heyuuu/gophp/php/assert"
)

// Object
type Object struct {
	handle uint
	ce     *Class

	protected bool

	properties map[string]Zval
}

func NewObject(ce *Class, handle uint) *Object {
	assert.Assert(ce != nil)
	obj := initObject(ce, handle)

	return obj
}

func initObject(ce *Class, handle uint) *Object {
	obj := &Object{
		handle:     handle,
		ce:         ce,
		properties: make(map[string]Zval),
	}

	ce.PropertyTable().Each(func(propName string, propInfo *PropertyInfo) {
		obj.properties[propName] = propInfo.defaultVal
	})

	return obj
}

func (o *Object) Handle() uint      { return o.handle }
func (o *Object) Ce() *Class        { return o.ce }
func (o *Object) ClassName() string { return o.ce.Name() }
func (o *Object) CeName() string    { return o.ce.Name() }

func (o *Object) ReadPropertyR(name string) Zval {
	return o.properties[name]
}

func (o *Object) PropertiesFor(typ PropPurposeType) *Array {
	arr := NewArrayCap(len(o.properties))
	for name, zval := range o.properties {
		arr.KeyAdd(name, zval)
	}
	return arr
}

func (o *Object) CanCast() bool {
	// todo
	return false
}

func (o *Object) Cast(typ ZvalType) (Zval, bool) {
	// todo
	return Undef, false
}

func (o *Object) CanCompare() bool {
	// todo
	return false
}

func (o *Object) CompareObjectsTo(other *Object) (int, bool) {
	// todo
	return 1, false
}

// recursive
func (o *Object) IsRecursive() bool   { return o.protected }
func (o *Object) ProtectRecursive()   { o.protected = true }
func (o *Object) UnprotectRecursive() { o.protected = false }
