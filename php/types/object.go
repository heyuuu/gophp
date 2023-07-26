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
type PropertyGuard uint8

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
var _ IObject = (*Object)(nil)

type Object struct {
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
func (o *Object) Guard(member string) *PropertyGuard {
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
func NewStdObject(ce *ClassEntry) *Object {
	return NewObject(ce, zend.StdObjectHandlersPtr)
}

func NewStdObjectSkipPropertiesInit(ce *ClassEntry) *Object {
	return _newObject(ce, zend.StdObjectHandlersPtr)
}

func NewStdObjectExEx(ce *ClassEntry, properties *Array) *Object {
	o := _newObject(ce, zend.StdObjectHandlersPtr)
	o.propertiesInitEx(properties)
	return o
}

func NewObject(ce *ClassEntry, handlers *ObjectHandlers) *Object {
	o := _newObject(ce, handlers)
	o.propertiesInit()
	return o
}

func NewObjectEx(data IObject) *Object {
	// todo 冗余细节待调整
	o := _newObject(data.GetCe(), nil)
	o.data = data
	return o
}

func _newObject(ce *ClassEntry, handlers *ObjectHandlers) *Object {
	propertyCount := ce.GetDefaultPropertiesCount()
	if ce.IsUseGuards() {
		propertyCount++
	}

	o := &Object{
		ce:              ce,
		handlers:        handlers,
		properties:      nil,
		propertiesTable: make([]Zval, propertyCount),
	}
	o.handle = uint(uintptr(unsafe.Pointer(o)))

	runtime.SetFinalizer(o, ObjectAutoFree)
	return o
}

func (o *Object) propertiesInit() {
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

func (o *Object) propertiesInitEx(properties *Array) {
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

func (o *Object) GetHandle() uint              { return o.handle }
func (o *Object) ClassName() string            { return o.data.ClassName() }
func (o *Object) GetCe() *ClassEntry           { return o.data.GetCe() }
func (o *Object) GetHandlers() *ObjectHandlers { return o.handlers }
func (o *Object) GetProperties() *Array        { return o.properties }
func (o *Object) SetProperties(value *Array)   { o.properties = value }
func (o *Object) DupProperties() {
	o.properties = o.properties.LazyDup()
}
func (o *Object) GetPropertiesTable() []Zval { return o.propertiesTable }

func (o *Object) GetData() IObject { return o.data }

// object handlers
func (o *Object) Free()          { o.data.Free() }
func (o *Object) Dtor()          { o.data.Dtor() }
func (o *Object) CanClone() bool { return o.data.CanClone() }
func (o *Object) Clone() *Object {
	b.Assert(o.data.CanClone())
	return o.data.Clone()
}

// property
func (o *Object) ReadPropertyEx(member *Zval, typ int, rv *Zval) *Zval {
	return o.data.ReadProperty(member, typ, nil, rv)
}
func (o *Object) WritePropertyEx(member *Zval, value *Zval) *Zval {
	return o.data.WriteProperty(member, value, nil)
}
func (o *Object) HasPropertyEx(member *Zval, hasSetExists int) bool {
	return o.data.HasProperty(member, hasSetExists, nil) != 0
}
func (o *Object) UnsetPropertyEx(member *Zval) {
	o.data.UnsetProperty(member, nil)
}
func (o *Object) GetPropertyPtrEx(member *Zval, typ int) *Zval {
	return o.data.GetPropertyPtr(member, typ, nil)
}

// properties
func (o *Object) IsStdGetProperties() bool {
	return o.data.IsStdGetProperties()
}
func (o *Object) GetPropertiesArray() *Array {
	return o.data.GetPropertiesArray()
}
func (o *Object) CanGetPropertiesFor() bool {
	return o.data.CanGetPropertiesFor()
}
func (o *Object) GetPropertiesFor(purpose zend.ZendPropPurpose) *Array {
	return o.data.GetPropertiesFor(purpose)
}

// get & set
func (o *Object) CanGet() bool {
	return o.data.CanGet()
}
func (o *Object) Get(rv *Zval) *Zval {
	return o.data.Get(rv)
}
func (o *Object) CanSet() bool {
	return o.data.CanSet()
}
func (o *Object) Set(value *Zval) {
	o.data.Set(value)
}

// dimension
func (o *Object) ReadDimension(offset *Zval, typ int, rv *Zval) *Zval {
	return o.data.ReadDimension(offset, typ, rv)
}
func (o *Object) WriteDimension(offset *Zval, value *Zval) {
	o.data.WriteDimension(offset, value)
}
func (o *Object) HasDimension(offset *Zval, checkEmpty int) int {
	return o.data.HasDimension(offset, checkEmpty)
}
func (o *Object) UnsetDimension(offset *Zval) {
	o.data.UnsetDimension(offset)
}

// elements
func (o *Object) CanCountElements() bool {
	return o.data.CanCountElements()
}
func (o *Object) CountElements() (int, bool) {
	return o.data.CountElements()
}

// method
func (o *Object) CanGetMethod() bool {
	return o.CanGetMethod()
}
func (o *Object) GetMethod(method string, key *Zval) IFunction {
	return o.data.GetMethod(method, key)
}
func (o *Object) CallMethod(method string, object *Object, executeData *zend.ZendExecuteData, returnValue *Zval) int {
	return o.data.CallMethod(method, object, executeData, returnValue)
}
func (o *Object) GetConstructor(object *Object) IFunction {
	return o.data.GetConstructor(object)
}

// cast
func (o *Object) CanCast() bool { return o.data.CanCast() }
func (o *Object) Cast(retval *Zval, type_ ZvalType) int {
	return o.data.Cast(retval, type_)
}

// mixed
func (o *Object) CanGetClosure() bool { return o.data.CanGetClosure() }
func (o *Object) GetClosure(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **Object) int {
	return o.data.GetClosure(obj, cePtr, fptrPtr, objPtr)
}

func (o *Object) CanDoOperation() bool { return o.data.CanDoOperation() }
func (o *Object) DoOperation(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int {
	return o.data.DoOperation(opcode, result, op1, op2)
}

func (o *Object) CanCompareObjectsTo(obj2 *Object) bool {
	return o.data.CanCompareObjectsTo(obj2)
}
func (o *Object) CompareObjectsTo(another *Object) int {
	return o.data.CompareObjectsTo(another)
}

func (o *Object) CanCompare() bool { return o.data.CanCompare() }
func (o *Object) Compare(result *Zval, op1 *Zval, op2 *Zval) int {
	return o.data.Compare(result, op1, op2)
}

// object
func (o *Object) IsObjDtorCalled() bool {
	return o.isDtorCalled
}
func (o *Object) MarkObjDtorCalled() {
	o.isDtorCalled = true
}
func (o *Object) IsObjFreeCalled() bool {
	return o.isFreeCalled
}
func (o *Object) MarkObjFreeCalled() {
	o.isFreeCalled = true
}

// recursive
func (o *Object) IsRecursive() bool   { return o.protected }
func (o *Object) ProtectRecursive()   { o.protected = true }
func (o *Object) UnprotectRecursive() { o.protected = false }
