package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"runtime"
)

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
func (o *ZendObject) Free() { o.handlers.freeObj(o) }
func (o *ZendObject) Dtor() { o.handlers.dtorObj(o) }

func (o *ZendObject) CanClone() bool { return o.handlers.cloneObj != nil }
func (o *ZendObject) Clone(zv *Zval) *ZendObject {
	b.Assert(o.handlers.cloneObj != nil)
	return o.handlers.cloneObj(zv)
}

// property
func (o *ZendObject) ReadProperty(object *Zval, member *Zval, typ int, cacheSlot *any, rv *Zval) *Zval {
	return o.handlers.readProperty(object, member, typ, cacheSlot, rv)
}
func (o *ZendObject) WriteProperty(object *Zval, member *Zval, value *Zval, cacheSlot *any) *Zval {
	return o.handlers.writeProperty(object, member, value, cacheSlot)
}
func (o *ZendObject) HasProperty(object *Zval, member *Zval, hasSetExists int, cacheSlot *any) int {
	return o.handlers.hasProperty(object, member, hasSetExists, cacheSlot)
}
func (o *ZendObject) UnsetProperty(object *Zval, member *Zval, cacheSlot *any) {
	o.handlers.unsetProperty(object, member, cacheSlot)
}
func (o *ZendObject) GetPropertyPtr(object *Zval, member *Zval, typ int, cacheSlot *any) *Zval {
	return o.handlers.getPropertyPtrPtr(object, member, typ, cacheSlot)
}

// properties
func (o *ZendObject) IsStdGetProperties() bool {
	std := zend.ZendStdGetProperties
	return o.handlers.getProperties == zend.ZendObjectGetPropertiesT(std)
}
func (o *ZendObject) GetPropertiesArray(object *Zval) *Array {
	return o.handlers.getProperties(object)
}
func (o *ZendObject) CanGetPropertiesFor() bool {
	return o.handlers.getPropertiesFor != nil
}
func (o *ZendObject) GetPropertiesFor(object *Zval, purpose zend.ZendPropPurpose) *Array {
	return o.handlers.getPropertiesFor(object, purpose)
}

// get & set
func (o *ZendObject) CanGet() bool { return o.handlers.get != nil }
func (o *ZendObject) Get(object *Zval, rv *Zval) *Zval {
	return o.handlers.get(object, rv)
}
func (o *ZendObject) CanSet() bool { return o.handlers.set != nil }
func (o *ZendObject) Set(object *Zval, value *Zval) {
	o.handlers.set(object, value)
}

// dimension
func (o *ZendObject) ReadDimension(object *Zval, offset *Zval, typ int, rv *Zval) *Zval {
	return o.handlers.readDimension(object, offset, typ, rv)
}
func (o *ZendObject) WriteDimension(object *Zval, offset *Zval, value *Zval) {
	o.handlers.writeDimension(object, offset, value)
}
func (o *ZendObject) HasDimension(object *Zval, offset *Zval, checkEmpty int) int {
	return o.handlers.hasDimension(object, offset, checkEmpty)
}
func (o *ZendObject) UnsetDimension(object *Zval, offset *Zval) {
	o.handlers.unsetDimension(object, offset)
}

// elements
func (o *ZendObject) CanCountElements() bool {
	return o.handlers.countElements != nil
}
func (o *ZendObject) CountElements(object *Zval, count *int) int {
	return o.handlers.countElements(object, count)
}

// method
func (o *ZendObject) CanGetMethod() bool { return o.handlers.getMethod != nil }
func (o *ZendObject) GetMethod(object **ZendObject, method *String, key *Zval) IFunction {
	return o.handlers.getMethod(object, method, key)
}
func (o *ZendObject) CallMethod(method *String, object *ZendObject, executeData *zend.ZendExecuteData, returnValue *Zval) int {
	return o.handlers.callMethod(method, object, executeData, returnValue)
}
func (o *ZendObject) GetConstructor(object *ZendObject) IFunction {
	return o.handlers.getConstructor(object)
}

// cast
func (o *ZendObject) CanCast() bool { return o.handlers.castObject != nil }
func (o *ZendObject) IsStdCast() bool {
	std := zend.ZendStdCastObjectTostring
	return o.handlers.castObject == zend.ZendObjectCastT(std)
}
func (o *ZendObject) Cast(readobj *Zval, retval *Zval, type_ ZvalType) int {
	return o.handlers.castObject(readobj, retval, type_)
}

// mixed
func (o *ZendObject) ClassName() string {
	return o.handlers.getClassName(o).GetStr()
}

func (o *ZendObject) CanGetClosure() bool {
	return o.handlers.getClosure != nil
}
func (o *ZendObject) GetClosure(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **ZendObject) int {
	return o.handlers.getClosure(obj, cePtr, fptrPtr, objPtr)
}

func (o *ZendObject) CanDoOperation() bool {
	return o.handlers.doOperation != nil
}
func (o *ZendObject) DoOperation(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int {
	return o.handlers.doOperation(opcode, result, op1, op2)
}

func (o *ZendObject) CanCompareObjectsTo(obj2 *ZendObject) bool {
	return o.handlers.compareObjects == obj2.handlers.compareObjects
}
func (o *ZendObject) CompareObjects(obj1, obj2 *Zval) int {
	return o.handlers.compareObjects(obj1, obj2)
}

func (o *ZendObject) CanCompare() bool {
	return o.handlers.compare != nil
}
func (o *ZendObject) Compare(result *Zval, op1 *Zval, op2 *Zval) int {
	return o.handlers.compare(result, op1, op2)
}
