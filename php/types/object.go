package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"runtime"
)

type objectCompareFunc func(object1 *Zval, object2 *Zval) int
type objectCastFunc func(readobj *Zval, retval *Zval, type_ ZvalType) int
type objectGetPropertiesFunc func(object *Zval) *Array

/**
 * ZendObject
 */
type ZendObject struct {
	ZendRefcounted
	handle          uint32
	ce              *ClassEntry
	handlers        *ObjectHandlers
	properties      *Array
	propertiesTable []Zval
}

func NewObject(ce *ClassEntry, handle uint32, handlers *ObjectHandlers) *ZendObject {
	propertyCount := ce.GetDefaultPropertiesCount()
	if ce.IsUseGuards() {
		propertyCount++
	}

	o := &ZendObject{}
	o.handlers = handlers
	o.propertiesTable = make([]Zval, propertyCount)

	o.Init(ce, handle)
	return o
}

func (o *ZendObject) Init(ce *ClassEntry, handle uint32) {
	o.SetGcTypeInfo(uint32(IS_OBJECT) | GC_COLLECTABLE<<GC_FLAGS_SHIFT)

	o.handle = handle
	o.ce = ce
	o.properties = nil

	if ce.IsUseGuards() {
		o.propertiesTable[ce.GetDefaultPropertiesCount()].SetUndef()
	}

	runtime.SetFinalizer(o, ObjectAutoFree)
}

func (o *ZendObject) GetHandle() uint32                 { return o.handle }
func (o *ZendObject) GetCe() *ClassEntry                { return o.ce }
func (o *ZendObject) GetHandlers() *ObjectHandlers      { return o.handlers }
func (o *ZendObject) SetHandlers(value *ObjectHandlers) { o.handlers = value }
func (o *ZendObject) GetProperties() *Array             { return o.properties }
func (o *ZendObject) SetProperties(value *Array)        { o.properties = value }
func (o *ZendObject) GetPropertiesTable() []Zval        { return o.propertiesTable }

// object handlers
func (o *ZendObject) Free() { o.handlers.FreeObj(o) }
func (o *ZendObject) Dtor() { o.handlers.DtorObj(o) }

func (o *ZendObject) CanClone() bool { return o.handlers.CloneObj != nil }
func (o *ZendObject) Clone(zv *Zval) *ZendObject {
	b.Assert(o.handlers.CloneObj != nil)
	return o.handlers.CloneObj(zv)
}

// property
func (o *ZendObject) ReadProperty(object *Zval, member *Zval, typ int, cacheSlot *any, rv *Zval) *Zval {
	return o.handlers.ReadProperty(object, member, typ, cacheSlot, rv)
}
func (o *ZendObject) WriteProperty(object *Zval, member *Zval, value *Zval, cacheSlot *any) *Zval {
	return o.handlers.WriteProperty(object, member, value, cacheSlot)
}
func (o *ZendObject) HasProperty(object *Zval, member *Zval, hasSetExists int, cacheSlot *any) int {
	return o.handlers.HasProperty(object, member, hasSetExists, cacheSlot)
}
func (o *ZendObject) UnsetProperty(object *Zval, member *Zval, cacheSlot *any) {
	o.handlers.UnsetProperty(object, member, cacheSlot)
}
func (o *ZendObject) GetPropertyPtr(object *Zval, member *Zval, typ int, cacheSlot *any) *Zval {
	return o.handlers.GetPropertyPtrPtr(object, member, typ, cacheSlot)
}

// properties
func (o *ZendObject) IsStdGetProperties() bool {
	std := zend.ZendStdGetProperties
	return objectGetPropertiesFunc(o.handlers.GetProperties) == objectGetPropertiesFunc(std)
}
func (o *ZendObject) GetPropertiesArray(object *Zval) *Array {
	return o.handlers.GetProperties(object)
}
func (o *ZendObject) CanGetPropertiesFor() bool {
	return o.handlers.GetPropertiesFor != nil
}
func (o *ZendObject) GetPropertiesFor(object *Zval, purpose zend.ZendPropPurpose) *Array {
	return o.handlers.GetPropertiesFor(object, purpose)
}

// get & set
func (o *ZendObject) CanGet() bool { return o.handlers.Get != nil }
func (o *ZendObject) Get(object *Zval, rv *Zval) *Zval {
	return o.handlers.Get(object, rv)
}
func (o *ZendObject) CanSet() bool { return o.handlers.Set != nil }
func (o *ZendObject) Set(object *Zval, value *Zval) {
	o.handlers.Set(object, value)
}

// dimension
func (o *ZendObject) ReadDimension(object *Zval, offset *Zval, typ int, rv *Zval) *Zval {
	return o.handlers.ReadDimension(object, offset, typ, rv)
}
func (o *ZendObject) WriteDimension(object *Zval, offset *Zval, value *Zval) {
	o.handlers.WriteDimension(object, offset, value)
}
func (o *ZendObject) HasDimension(object *Zval, offset *Zval, checkEmpty int) int {
	return o.handlers.HasDimension(object, offset, checkEmpty)
}
func (o *ZendObject) UnsetDimension(object *Zval, offset *Zval) {
	o.handlers.UnsetDimension(object, offset)
}

// elements
func (o *ZendObject) CanCountElements() bool {
	return o.handlers.CountElements != nil
}
func (o *ZendObject) CountElements(object *Zval, count *int) int {
	return o.handlers.CountElements(object, count)
}

// method
func (o *ZendObject) CanGetMethod() bool { return o.handlers.GetMethod != nil }
func (o *ZendObject) GetMethod(object **ZendObject, method *String, key *Zval) IFunction {
	return o.handlers.GetMethod(object, method, key)
}
func (o *ZendObject) CallMethod(method *String, object *ZendObject, executeData *zend.ZendExecuteData, returnValue *Zval) int {
	return o.handlers.CallMethod(method, object, executeData, returnValue)
}
func (o *ZendObject) GetConstructor(object *ZendObject) IFunction {
	return o.handlers.GetConstructor(object)
}

// cast
func (o *ZendObject) CanCast() bool { return o.handlers.CastObject != nil }
func (o *ZendObject) IsStdCast() bool {
	std := zend.ZendStdCastObjectTostring
	return objectCastFunc(o.handlers.CastObject) == objectCastFunc(std)
}
func (o *ZendObject) Cast(readobj *Zval, retval *Zval, type_ ZvalType) int {
	return o.handlers.CastObject(readobj, retval, type_)
}

// mixed
func (o *ZendObject) ClassName() string {
	return o.handlers.GetClassName(o).GetStr()
}

func (o *ZendObject) CanGetClosure() bool {
	return o.handlers.GetClosure != nil
}
func (o *ZendObject) GetClosure(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **ZendObject) int {
	return o.handlers.GetClosure(obj, cePtr, fptrPtr, objPtr)
}

func (o *ZendObject) CanDoOperation() bool {
	return o.handlers.DoOperation != nil
}
func (o *ZendObject) DoOperation(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int {
	return o.handlers.DoOperation(opcode, result, op1, op2)
}

func (o *ZendObject) CanCompareObjectsTo(obj2 *ZendObject) bool {
	return objectCompareFunc(o.handlers.CompareObjects) == objectCompareFunc(obj2.handlers.CompareObjects)
}
func (o *ZendObject) CompareObjects(obj1, obj2 *Zval) int {
	return o.handlers.CompareObjects(obj1, obj2)
}

func (o *ZendObject) CanCompare() bool {
	return o.handlers.Compare != nil
}
func (o *ZendObject) Compare(result *Zval, op1 *Zval, op2 *Zval) int {
	return o.handlers.Compare(result, op1, op2)
}
