package types

import (
	"github.com/heyuuu/gophp/zend"
)

var _ ObjectData = (*StdObjectData)(nil)

type StdObjectData struct {
	ce              *ClassEntry
	properties      *Array // 动态属性
	propertiesTable []Zval // 静态属性
}

func (o *StdObjectData) ClassName() string { return o.ce.Name() }
func (o *StdObjectData) Free() {
	zend.ZendObjectStdDtorEx(o.propertiesTable, o.ce)
}

func (o *StdObjectData) Dtor() {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanClone() bool { return true }
func (o *StdObjectData) Clone(zv *Zval) *ZendObject {

}

func (o *StdObjectData) ReadProperty(object *Zval, member *Zval, typ int, cacheSlot *any, rv *Zval) *Zval {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) WriteProperty(object *Zval, member *Zval, value *Zval, cacheSlot *any) *Zval {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) HasProperty(object *Zval, member *Zval, hasSetExists int, cacheSlot *any) int {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) UnsetProperty(object *Zval, member *Zval, cacheSlot *any) {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) GetPropertyPtr(object *Zval, member *Zval, typ int, cacheSlot *any) *Zval {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) GetPropertiesArray(object *Zval) *Array {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanGetPropertiesFor() bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) GetPropertiesFor(object *Zval, purpose zend.ZendPropPurpose) *Array {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanGet() bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) Get(object *Zval, rv *Zval) *Zval {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanSet() bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) Set(object *Zval, value *Zval) {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) ReadDimension(object *Zval, offset *Zval, typ int, rv *Zval) *Zval {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) WriteDimension(object *Zval, offset *Zval, value *Zval) {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) HasDimension(object *Zval, offset *Zval, checkEmpty int) int {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) UnsetDimension(object *Zval, offset *Zval) {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanCountElements() bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CountElements(object *Zval, count *int) int {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanGetMethod() bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) GetMethod(object **ZendObject, method *String, key *Zval) IFunction {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CallMethod(method *String, object *ZendObject, executeData *zend.ZendExecuteData, returnValue *Zval) int {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) GetConstructor(object *ZendObject) IFunction {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanCast() bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) Cast(readobj *Zval, retval *Zval, type_ ZvalType) int {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanGetClosure() bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) GetClosure(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **ZendObject) int {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanDoOperation() bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) DoOperation(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanCompareObjectsTo(obj2 *ZendObject) bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CompareObjects(obj1, obj2 *Zval) int {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) CanCompare() bool {
	//TODO implement me
	panic("implement me")
}

func (o *StdObjectData) Compare(result *Zval, op1 *Zval, op2 *Zval) int {
	//TODO implement me
	panic("implement me")
}
