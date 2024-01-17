package types

import (
	"github.com/heyuuu/gophp/php/assert"
	"unsafe"
)

// Object
type Object struct {
	handle uint
	ce     *Class

	protected bool
}

func NewObject(ce *Class) *Object {
	assert.Assert(ce != nil)
	return initObject(ce)
}

func initObject(ce *Class) *Object {
	o := &Object{
		ce: ce,
	}
	o.handle = uint(uintptr(unsafe.Pointer(o)))
	return o
}

func (o *Object) Handle() uint      { return o.handle }
func (o *Object) Ce() *Class        { return o.ce }
func (o *Object) ClassName() string { return o.ce.Name() }
func (o *Object) CeName() string    { return o.ce.Name() }

func (o *Object) PropertiesFor(typ PropPurposeType) *Array {
	return nil
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
