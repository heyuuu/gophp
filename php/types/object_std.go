package types

import "github.com/heyuuu/gophp/zend"

var _ = zend.StdObjectHandlersPtr
var _ IObject = (*ObjectStd)(nil)

type ObjectStd struct {
	ce              *ClassEntry
	properties      *Array // 动态属性
	propertiesTable []Zval // 静态属性
}

func (o *ObjectStd) obj() *Object { return o }

func (o *ObjectStd) ClassName() string          { return o.ce.Name() }
func (o *ObjectStd) GetCe() *ClassEntry         { return o.ce }
func (o *ObjectStd) GetProperties() *Array      { return o.properties }
func (o *ObjectStd) GetPropertiesTable() []Zval { return o.propertiesTable }

func (o *ObjectStd) Free() { zend.ZendObjectStdDtor(o.obj()) }
func (o *ObjectStd) Dtor() { zend.ZendObjectsDestroyObject(o.obj()) }

func (o *ObjectStd) CanClone() bool { return true }
func (o *ObjectStd) Clone() *Object {
	return zend.ZendObjectsCloneObjEx(o.obj())
}

func (o *ObjectStd) ReadProperty(member *Zval, typ int, cacheSlot *any, rv *Zval) *Zval {
	return zend.ZendStdReadPropertyEx(o.obj(), member, typ, cacheSlot, rv)
}

func (o *ObjectStd) WriteProperty(member *Zval, value *Zval, cacheSlot *any) *Zval {
	return zend.ZendStdWritePropertyEx(o.obj(), member, value, cacheSlot)
}

func (o *ObjectStd) HasProperty(member *Zval, hasSetExists int, cacheSlot *any) int {
	return zend.ZendStdHasPropertyEx(o.obj(), member, hasSetExists, cacheSlot)
}

func (o *ObjectStd) UnsetProperty(member *Zval, cacheSlot *any) {
	zend.ZendStdUnsetPropertyEx(o.obj(), member, cacheSlot)
}

func (o *ObjectStd) GetPropertyPtr(member *Zval, typ int, cacheSlot *any) *Zval {
	return zend.ZendStdGetPropertyPtrPtrEx(o.obj(), member, typ, cacheSlot)
}

func (o *ObjectStd) IsStdGetProperties() bool { return true }
func (o *ObjectStd) GetPropertiesArray() *Array {
	return zend.ZendStdGetPropertiesEx(o.obj())
}

func (o *ObjectStd) CanGetPropertiesFor() bool                            { return false }
func (o *ObjectStd) GetPropertiesFor(purpose zend.ZendPropPurpose) *Array { panic("implement me") }

func (o *ObjectStd) CanGet() bool       { return false }
func (o *ObjectStd) Get(rv *Zval) *Zval { panic("implement me") }

func (o *ObjectStd) CanSet() bool    { return false }
func (o *ObjectStd) Set(value *Zval) { panic("implement me") }

func (o *ObjectStd) ReadDimension(offset *Zval, typ int, rv *Zval) *Zval {
	return zend.ZendStdReadDimensionEx(o.obj(), offset, typ, rv)
}

func (o *ObjectStd) WriteDimension(offset *Zval, value *Zval) {
	zend.ZendStdWriteDimensionEx(o.obj(), offset, value)
}

func (o *ObjectStd) HasDimension(offset *Zval, checkEmpty int) int {
	return zend.ZendStdHasDimensionEx(o.obj(), offset, checkEmpty)
}

func (o *ObjectStd) UnsetDimension(offset *Zval) {
	zend.ZendStdUnsetDimensionEx(o.obj(), offset)
}

func (o *ObjectStd) CanCountElements() bool     { return false }
func (o *ObjectStd) CountElements() (int, bool) { panic("implement me") }

func (o *ObjectStd) CanGetMethod() bool { return true }

func (o *ObjectStd) GetMethod(method string, key *Zval) IFunction {
	return zend.ZendStdGetMethod_Ex(o.obj(), method, key)
}

func (o *ObjectStd) CallMethod(method string, object *Object, executeData *zend.ZendExecuteData, returnValue *Zval) int {
	panic("implement me")
}

func (o *ObjectStd) GetConstructor(object *Object) IFunction {
	return zend.ZendStdGetConstructor(object)
}

func (o *ObjectStd) CanCast() bool { return true }

func (o *ObjectStd) Cast(retval *Zval, type_ ZvalType) int {
	return zend.ZendStdCastObject(o.obj(), retval, type_)
}

func (o *ObjectStd) CanGetClosure() bool { return false }
func (o *ObjectStd) GetClosure(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **Object) int {
	panic("implement me")
}

func (o *ObjectStd) CanDoOperation() bool { return false }
func (o *ObjectStd) DoOperation(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int {
	panic("implement me")
}

func (o *ObjectStd) CanCompareObjectsTo(obj2 *Object) bool {
	// todo
	return objectCompareFunc(o.obj().handlers.CompareObjects) == objectCompareFunc(obj2.handlers.CompareObjects)
}

func (o *ObjectStd) CompareObjectsTo(another *Object) int {
	return zend.ZendStdCompareObjectsEx(o.obj(), another)
}

func (o *ObjectStd) CanCompare() bool { return false }
func (o *ObjectStd) Compare(result *Zval, op1 *Zval, op2 *Zval) int {
	panic("implement me")
}
