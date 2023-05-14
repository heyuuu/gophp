package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"runtime"
	"unsafe"
)

type objectCompareFunc func(object1 *Zval, object2 *Zval) int
type objectCastFunc func(readobj *Zval, retval *Zval, type_ ZvalType) int
type objectGetPropertiesFunc func(object *Zval) *Array

/**
 * ZendObject
 */
type ZendObject struct {
	handle          uint
	ce              *ClassEntry
	handlers        *ObjectHandlers
	properties      *Array // 动态属性
	propertiesTable []Zval // 静态属性

	// flags todo 待合并
	protected    bool
	isDtorCalled bool
	isFreeCalled bool
}

func NewStdObject(ce *ClassEntry) *ZendObject {
	return NewObject(ce, zend.StdObjectHandlersPtr)
}

func NewStdObjectEx(ce *ClassEntry) *ZendObject {
	return NewObjectEx(ce, zend.StdObjectHandlersPtr)
}

func NewStdObjectExEx(ce *ClassEntry, properties *Array) *ZendObject {
	o := NewObject(ce, zend.StdObjectHandlersPtr)
	o.PropertiesInitEx(properties)
	return o
}

func NewObject(ce *ClassEntry, handlers *ObjectHandlers) *ZendObject {
	propertyCount := ce.GetDefaultPropertiesCount()
	if ce.IsUseGuards() {
		propertyCount++
	}

	o := &ZendObject{
		ce:              ce,
		handlers:        handlers,
		properties:      nil,
		propertiesTable: make([]Zval, propertyCount),
	}

	o.handle = uint(uintptr(unsafe.Pointer(o)))

	if ce.IsUseGuards() {
		o.propertiesTable[ce.GetDefaultPropertiesCount()].SetUndef()
	}

	runtime.SetFinalizer(o, ObjectAutoFree)
	return o
}

func NewObjectEx(ce *ClassEntry, handlers *ObjectHandlers) *ZendObject {
	o := NewObject(ce, handlers)

	// init properties
	defaultPropertiesCount := ce.GetDefaultPropertiesCount()
	if defaultPropertiesCount != 0 {
		src := ce.GetDefaultPropertiesTable()
		dst := o.propertiesTable
		if ce.GetType() == zend.ZEND_INTERNAL_CLASS {
			for i := 0; i < defaultPropertiesCount; i++ {
				ZVAL_COPY_OR_DUP_PROP(&dst[i], &src[i])
			}
		} else {
			for i := 0; i < defaultPropertiesCount; i++ {
				ZVAL_COPY_PROP(&dst[i], &src[i])
			}
		}
	}

	return o
}

func (o *ZendObject) PropertiesInitEx(properties *Array) {
	o.properties = properties
	defaultPropertiesCount := o.ce.GetDefaultPropertiesCount()
	if defaultPropertiesCount != 0 {
		properties.Foreach(func(key_ ArrayKey, prop *Zval) {
			propertyInfo := zend.ZendGetPropertyInfo(o.GetCe(), key_.StrKey(), 1)
			if propertyInfo != zend.ZEND_WRONG_PROPERTY_INFO && propertyInfo != nil && !propertyInfo.IsStatic() {
				var slot *Zval = zend.OBJ_PROP(o, propertyInfo.GetOffset())
				if propertyInfo.GetType() != 0 {
					var tmp Zval
					ZVAL_COPY_VALUE(&tmp, prop)
					if zend.ZendVerifyPropertyType(propertyInfo, &tmp, 0) == 0 {
						return
					}
					ZVAL_COPY_VALUE(slot, &tmp)
				} else {
					slot.CopyValueFrom(prop)
				}
				prop.SetIndirect(slot)
			}
		})
	}
}

func (o *ZendObject) GetHandle() uint                   { return o.handle }
func (o *ZendObject) GetCe() *ClassEntry                { return o.ce }
func (o *ZendObject) GetHandlers() *ObjectHandlers      { return o.handlers }
func (o *ZendObject) SetHandlers(value *ObjectHandlers) { o.handlers = value }
func (o *ZendObject) GetProperties() *Array             { return o.properties }
func (o *ZendObject) SetProperties(value *Array)        { o.properties = value }
func (o *ZendObject) DupProperties() {
	o.properties = o.properties.LazyDup()
}
func (o *ZendObject) GetPropertiesTable() []Zval { return o.propertiesTable }

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

// object
func (o *ZendObject) IsObjDtorCalled() bool {
	return o.isDtorCalled
}
func (o *ZendObject) MarkObjDtorCalled() {
	o.isDtorCalled = true
}
func (o *ZendObject) IsObjFreeCalled() bool {
	return o.isFreeCalled
}
func (o *ZendObject) MarkObjFreeCalled() {
	o.isFreeCalled = true
}

// recursive
func (o *ZendObject) IsRecursive() bool   { return o.protected }
func (o *ZendObject) ProtectRecursive()   { o.protected = true }
func (o *ZendObject) UnprotectRecursive() { o.protected = false }
