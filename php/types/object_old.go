package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
)

var _ IObject = (*ObjectOld)(nil)

type ObjectOld struct {
	ce              *ClassEntry
	handlers        *ObjectHandlers
	properties      *Array // 动态属性
	propertiesTable []Zval // 静态属性
}

func (o *ObjectOld) obj() *ZendObject { return o }

func (o *ObjectOld) ClassName() string  { return o.ce.Name() }
func (o *ObjectOld) GetCe() *ClassEntry { return o.ce }

func (o *ObjectOld) Free() { o.handlers.FreeObj(o.obj()) }
func (o *ObjectOld) Dtor() { o.handlers.DtorObj(o.obj()) }

func (o *ObjectOld) CanClone() bool { return o.handlers.CloneObjEx != nil }
func (o *ObjectOld) Clone() *ZendObject {
	b.Assert(o.handlers.CloneObjEx != nil)
	return o.handlers.CloneObjEx(o.obj())
}

func (o *ObjectOld) ReadProperty(member *Zval, typ int, cacheSlot *any, rv *Zval) *Zval {
	return o.handlers.ReadPropertyEx(o.obj(), member, typ, cacheSlot, rv)
}

func (o *ObjectOld) WriteProperty(member *Zval, value *Zval, cacheSlot *any) *Zval {
	return o.handlers.WritePropertyEx(o.obj(), member, value, cacheSlot)
}

func (o *ObjectOld) HasProperty(member *Zval, hasSetExists int, cacheSlot *any) int {
	return o.handlers.HasPropertyEx(o.obj(), member, hasSetExists, cacheSlot)
}

func (o *ObjectOld) UnsetProperty(member *Zval, cacheSlot *any) {
	o.handlers.UnsetPropertyEx(o.obj(), member, cacheSlot)
}

func (o *ObjectOld) GetPropertyPtr(member *Zval, typ int, cacheSlot *any) *Zval {
	return o.handlers.GetPropertyPtrPtrEx(o.obj(), member, typ, cacheSlot)
}

func (o *ObjectOld) IsStdGetProperties() bool {
	// todo
	std := zend.ZendStdGetProperties
	return objectGetPropertiesFunc(o.handlers.GetProperties) == objectGetPropertiesFunc(std)
}

func (o *ObjectOld) GetPropertiesArray() *Array {
	return o.handlers.GetPropertiesEx(o.obj())
}

func (o *ObjectOld) CanGetPropertiesFor() bool { return o.handlers.GetPropertiesForEx != nil }
func (o *ObjectOld) GetPropertiesFor(purpose zend.ZendPropPurpose) *Array {
	return o.handlers.GetPropertiesForEx(o.obj(), purpose)
}

func (o *ObjectOld) CanGet() bool       { return o.handlers.GetEx != nil }
func (o *ObjectOld) Get(rv *Zval) *Zval { return o.handlers.GetEx(o.obj(), rv) }

func (o *ObjectOld) CanSet() bool    { return o.handlers.SetEx != nil }
func (o *ObjectOld) Set(value *Zval) { o.handlers.SetEx(o.obj(), value) }

func (o *ObjectOld) ReadDimension(offset *Zval, typ int, rv *Zval) *Zval {
	return o.handlers.ReadDimensionEx(o.obj(), offset, typ, rv)
}

func (o *ObjectOld) WriteDimension(offset *Zval, value *Zval) {
	o.handlers.WriteDimensionEx(o.obj(), offset, value)
}

func (o *ObjectOld) HasDimension(offset *Zval, checkEmpty int) int {
	return o.handlers.HasDimensionEx(o.obj(), offset, checkEmpty)
}

func (o *ObjectOld) UnsetDimension(offset *Zval) { o.handlers.UnsetDimensionEx(o.obj(), offset) }

func (o *ObjectOld) CanCountElements() bool { return o.handlers.CountElementsEx != nil }
func (o *ObjectOld) CountElements() (int, bool) {
	var count int
	ret := o.handlers.CountElementsEx(o.obj(), &count)
	return count, ret != 0
}

func (o *ObjectOld) CanGetMethod() bool { return o.handlers.GetMethod != nil }
func (o *ObjectOld) GetMethod(method string, key *Zval) IFunction {
	obj := o.obj()
	return o.handlers.GetMethod(&obj, NewString(method), key)
}

func (o *ObjectOld) CallMethod(method string, object *ZendObject, executeData *zend.ZendExecuteData, returnValue *Zval) int {
	return o.handlers.CallMethod(NewString(method), object, executeData, returnValue)
}

func (o *ObjectOld) GetConstructor(object *ZendObject) IFunction {
	return o.handlers.GetConstructor(object)
}

func (o *ObjectOld) CanCast() bool { return o.handlers.CastObjectEx != nil }
func (o *ObjectOld) Cast(retval *Zval, type_ ZvalType) int {
	return o.handlers.CastObjectEx(o.obj(), retval, type_)
}

func (o *ObjectOld) CanGetClosure() bool { return o.handlers.GetClosureEx != nil }
func (o *ObjectOld) GetClosure(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **ZendObject) int {
	return o.handlers.GetClosureEx(o.obj(), cePtr, fptrPtr, objPtr)
}

func (o *ObjectOld) CanDoOperation() bool { return o.handlers.GetClosureEx != nil }

func (o *ObjectOld) DoOperation(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int {
	return o.handlers.DoOperation(opcode, result, op1, op2)
}

func (o *ObjectOld) CanCompareObjectsTo(obj2 *ZendObject) bool {
	return objectCompareFunc(o.handlers.CompareObjects) == objectCompareFunc(obj2.handlers.CompareObjects)
}

func (o *ObjectOld) CompareObjectsTo(another *ZendObject) int {
	return o.handlers.CompareObjectsEx(o.obj(), another)
}

func (o *ObjectOld) CanCompare() bool { return o.handlers.Compare != nil }
func (o *ObjectOld) Compare(result *Zval, op1 *Zval, op2 *Zval) int {
	return o.handlers.Compare(result, op1, op2)
}
