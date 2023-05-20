package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"runtime"
	"unsafe"
)

type objectCompareFunc func(object1 *Zval, object2 *Zval) int
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

func NewStdObjectSkipPropertiesInit(ce *ClassEntry) *ZendObject {
	return _newObject(ce, zend.StdObjectHandlersPtr)
}

func NewStdObjectExEx(ce *ClassEntry, properties *Array) *ZendObject {
	o := _newObject(ce, zend.StdObjectHandlersPtr)
	o.propertiesInitEx(properties)
	return o
}

func NewObject(ce *ClassEntry, handlers *ObjectHandlers) *ZendObject {
	o := _newObject(ce, handlers)
	o.propertiesInit()
	return o
}

func _newObject(ce *ClassEntry, handlers *ObjectHandlers) *ZendObject {
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

	runtime.SetFinalizer(o, ObjectAutoFree)
	return o
}

func (o *ZendObject) propertiesInit() {
	defaultPropertiesCount := o.ce.GetDefaultPropertiesCount()
	if defaultPropertiesCount != 0 {
		src := o.ce.GetDefaultPropertiesTable()
		dst := o.propertiesTable
		if o.ce.IsInternalClass() {
			for i := 0; i < defaultPropertiesCount; i++ {
				dst[i].CopyOrDupPropFrom(&src[i])
			}
		} else {
			for i := 0; i < defaultPropertiesCount; i++ {
				dst[i].CopyPropFrom(&src[i])
			}
		}
	}
}

func (o *ZendObject) propertiesInitEx(properties *Array) {
	o.properties = properties
	defaultPropertiesCount := o.ce.GetDefaultPropertiesCount()
	if defaultPropertiesCount != 0 {
		properties.Foreach(func(key_ ArrayKey, prop *Zval) {
			propertyInfo := zend.ZendGetPropertyInfo(o.GetCe(), key_.StrKey())
			if propertyInfo != nil && !propertyInfo.IsStatic() {
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

func (o *ZendObject) GetHandle() uint              { return o.handle }
func (o *ZendObject) GetCe() *ClassEntry           { return o.ce }
func (o *ZendObject) GetHandlers() *ObjectHandlers { return o.handlers }
func (o *ZendObject) GetProperties() *Array        { return o.properties }
func (o *ZendObject) SetProperties(value *Array)   { o.properties = value }
func (o *ZendObject) DupProperties() {
	o.properties = o.properties.LazyDup()
}
func (o *ZendObject) GetPropertiesTable() []Zval { return o.propertiesTable }

// object handlers
func (o *ZendObject) Free() { o.handlers.FreeObj(o) }
func (o *ZendObject) Dtor() { o.handlers.DtorObj(o) }

func (o *ZendObject) CanClone() bool { return o.handlers.CloneObjEx != nil }
func (o *ZendObject) Clone() *ZendObject {
	b.Assert(o.handlers.CloneObjEx != nil)
	return o.handlers.CloneObjEx(o)
}

// property
func (o *ZendObject) ReadProperty(member *Zval, typ int, cacheSlot *any, rv *Zval) *Zval {
	return o.handlers.ReadPropertyEx(o, member, typ, cacheSlot, rv)
}
func (o *ZendObject) WriteProperty(member *Zval, value *Zval, cacheSlot *any) *Zval {
	return o.handlers.WritePropertyEx(o, member, value, cacheSlot)
}
func (o *ZendObject) HasProperty(member *Zval, hasSetExists int, cacheSlot *any) int {
	return o.handlers.HasPropertyEx(o, member, hasSetExists, cacheSlot)
}
func (o *ZendObject) UnsetPropertyEx(member *Zval, cacheSlot *any) {
	o.handlers.UnsetPropertyEx(o, member, cacheSlot)
}
func (o *ZendObject) GetPropertyPtr(member *Zval, typ int, cacheSlot *any) *Zval {
	return o.handlers.GetPropertyPtrPtrEx(o, member, typ, cacheSlot)
}

// properties
func (o *ZendObject) IsStdGetProperties() bool {
	// todo
	std := zend.ZendStdGetProperties
	return objectGetPropertiesFunc(o.handlers.GetProperties) == objectGetPropertiesFunc(std)
}
func (o *ZendObject) GetPropertiesArray() *Array {
	return o.handlers.GetPropertiesEx(o)
}
func (o *ZendObject) CanGetPropertiesFor() bool {
	return o.handlers.GetPropertiesForEx != nil
}
func (o *ZendObject) GetPropertiesFor(purpose zend.ZendPropPurpose) *Array {
	return o.handlers.GetPropertiesForEx(o, purpose)
}

// get & set
func (o *ZendObject) CanGet() bool { return o.handlers.GetEx != nil }
func (o *ZendObject) Get(rv *Zval) *Zval {
	return o.handlers.GetEx(o, rv)
}
func (o *ZendObject) CanSet() bool { return o.handlers.SetEx != nil }
func (o *ZendObject) Set(value *Zval) {
	o.handlers.SetEx(o, value)
}

// dimension
func (o *ZendObject) ReadDimension(offset *Zval, typ int, rv *Zval) *Zval {
	return o.handlers.ReadDimensionEx(o, offset, typ, rv)
}
func (o *ZendObject) WriteDimension(offset *Zval, value *Zval) {
	o.handlers.WriteDimensionEx(o, offset, value)
}
func (o *ZendObject) HasDimension(offset *Zval, checkEmpty int) int {
	return o.handlers.HasDimensionEx(o, offset, checkEmpty)
}
func (o *ZendObject) UnsetDimension(offset *Zval) {
	o.handlers.UnsetDimensionEx(o, offset)
}

// elements
func (o *ZendObject) CanCountElements() bool {
	return o.handlers.CountElementsEx != nil
}
func (o *ZendObject) CountElements(count *int) int {
	return o.handlers.CountElementsEx(o, count)
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
func (o *ZendObject) CanCast() bool { return o.handlers.CastObjectEx != nil }
func (o *ZendObject) Cast(retval *Zval, type_ ZvalType) int {
	return o.handlers.CastObjectEx(o, retval, type_)
}

// mixed
func (o *ZendObject) ClassName() string {
	return o.ce.Name()
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
func (o *ZendObject) CompareObjectsTo(another *ZendObject) int {
	return o.handlers.CompareObjectsEx(o, another)
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
