package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"runtime"
	"unsafe"
)

type objectCompareFunc func(object1 *Zval, object2 *Zval) int
type objectGetPropertiesFunc func(object *Zval) *Array

const (
	guardInGet = 1 << iota
	guardInSet
	guardInUnset
	guardInIsset
)

/**
 * PropertyGuard
 */
type PropertyGuard uint32

func (guard PropertyGuard) InGet() bool         { return guard&guardInGet != 0 }
func (guard PropertyGuard) InSet() bool         { return guard&guardInSet != 0 }
func (guard PropertyGuard) InIsset() bool       { return guard&guardInIsset != 0 }
func (guard PropertyGuard) InUnset() bool       { return guard&guardInUnset != 0 }
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

/**
 * ZendObject
 */
var _ IObject = (*ZendObject)(nil)

type ZendObject struct {
	handle          uint
	ce              *ClassEntry
	handlers        *ObjectHandlers
	properties      *Array                    // 动态属性
	propertiesTable []Zval                    // 静态属性
	propertyGuards  map[string]*PropertyGuard // 属性 guard

	data IObject // 封装 Object 数据，便于扩展

	// flags todo 待合并
	protected    bool
	isDtorCalled bool
	isFreeCalled bool
}

// guard
func (o *ZendObject) Guard(member string) *PropertyGuard {
	b.Assert(o.ce.IsUseGuards())
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

//
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

func NewObjectEx(data IObject) *ZendObject {
	// todo 冗余细节待调整
	o := _newObject(data.GetCe(), nil)
	o.data = data
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
				if propertyInfo.GetType() != nil {
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
func (o *ZendObject) ClassName() string            { return o.data.ClassName() }
func (o *ZendObject) GetCe() *ClassEntry           { return o.data.GetCe() }
func (o *ZendObject) GetHandlers() *ObjectHandlers { return o.handlers }
func (o *ZendObject) GetProperties() *Array        { return o.properties }
func (o *ZendObject) SetProperties(value *Array)   { o.properties = value }
func (o *ZendObject) DupProperties() {
	o.properties = o.properties.LazyDup()
}
func (o *ZendObject) GetPropertiesTable() []Zval { return o.propertiesTable }

func (o *ZendObject) GetData() IObject { return o.data }

// object handlers
func (o *ZendObject) Free()          { o.data.Free() }
func (o *ZendObject) Dtor()          { o.data.Dtor() }
func (o *ZendObject) CanClone() bool { return o.data.CanClone() }
func (o *ZendObject) Clone() *ZendObject {
	b.Assert(o.data.CanClone())
	return o.data.Clone()
}

// property
func (o *ZendObject) ReadProperty(member *Zval, typ int, cacheSlot *any, rv *Zval) *Zval {
	return o.data.ReadProperty(member, typ, cacheSlot, rv)
}
func (o *ZendObject) WriteProperty(member *Zval, value *Zval, cacheSlot *any) *Zval {
	return o.data.WriteProperty(member, value, cacheSlot)
}
func (o *ZendObject) HasProperty(member *Zval, hasSetExists int, cacheSlot *any) int {
	return o.data.HasProperty(member, hasSetExists, cacheSlot)
}
func (o *ZendObject) UnsetProperty(member *Zval, cacheSlot *any) {
	o.data.UnsetProperty(member, cacheSlot)
}
func (o *ZendObject) GetPropertyPtr(member *Zval, typ int, cacheSlot *any) *Zval {
	return o.data.GetPropertyPtr(member, typ, cacheSlot)
}

// properties
func (o *ZendObject) IsStdGetProperties() bool {
	return o.data.IsStdGetProperties()
}
func (o *ZendObject) GetPropertiesArray() *Array {
	return o.data.GetPropertiesArray()
}
func (o *ZendObject) CanGetPropertiesFor() bool {
	return o.data.CanGetPropertiesFor()
}
func (o *ZendObject) GetPropertiesFor(purpose zend.ZendPropPurpose) *Array {
	return o.data.GetPropertiesFor(purpose)
}

// get & set
func (o *ZendObject) CanGet() bool {
	return o.data.CanGet()
}
func (o *ZendObject) Get(rv *Zval) *Zval {
	return o.data.Get(rv)
}
func (o *ZendObject) CanSet() bool {
	return o.data.CanSet()
}
func (o *ZendObject) Set(value *Zval) {
	o.data.Set(value)
}

// dimension
func (o *ZendObject) ReadDimension(offset *Zval, typ int, rv *Zval) *Zval {
	return o.data.ReadDimension(offset, typ, rv)
}
func (o *ZendObject) WriteDimension(offset *Zval, value *Zval) {
	o.data.WriteDimension(offset, value)
}
func (o *ZendObject) HasDimension(offset *Zval, checkEmpty int) int {
	return o.data.HasDimension(offset, checkEmpty)
}
func (o *ZendObject) UnsetDimension(offset *Zval) {
	o.data.UnsetDimension(offset)
}

// elements
func (o *ZendObject) CanCountElements() bool {
	return o.data.CanCountElements()
}
func (o *ZendObject) CountElements() (int, bool) {
	return o.data.CountElements()
}

// method
func (o *ZendObject) CanGetMethod() bool {
	return o.CanGetMethod()
}
func (o *ZendObject) GetMethod(method string, key *Zval) IFunction {
	return o.data.GetMethod(method, key)
}
func (o *ZendObject) CallMethod(method string, object *ZendObject, executeData *zend.ZendExecuteData, returnValue *Zval) int {
	return o.data.CallMethod(method, object, executeData, returnValue)
}
func (o *ZendObject) GetConstructor(object *ZendObject) IFunction {
	return o.data.GetConstructor(object)
}

// cast
func (o *ZendObject) CanCast() bool { return o.data.CanCast() }
func (o *ZendObject) Cast(retval *Zval, type_ ZvalType) int {
	return o.data.Cast(retval, type_)
}

// mixed
func (o *ZendObject) CanGetClosure() bool { return o.data.CanGetClosure() }
func (o *ZendObject) GetClosure(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **ZendObject) int {
	return o.data.GetClosure(obj, cePtr, fptrPtr, objPtr)
}

func (o *ZendObject) CanDoOperation() bool { return o.data.CanDoOperation() }
func (o *ZendObject) DoOperation(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int {
	return o.data.DoOperation(opcode, result, op1, op2)
}

func (o *ZendObject) CanCompareObjectsTo(obj2 *ZendObject) bool {
	return o.data.CanCompareObjectsTo(obj2)
}
func (o *ZendObject) CompareObjectsTo(another *ZendObject) int {
	return o.data.CompareObjectsTo(another)
}

func (o *ZendObject) CanCompare() bool { return o.data.CanCompare() }
func (o *ZendObject) Compare(result *Zval, op1 *Zval, op2 *Zval) int {
	return o.data.Compare(result, op1, op2)
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
