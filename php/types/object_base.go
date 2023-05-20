package types

import (
	"github.com/heyuuu/gophp/zend"
)

var _ IObject = (*ObjectBase)(nil)

type ObjectBase struct {
	ce *ClassEntry
}

func (o *ObjectBase) ClassName() string  { return o.ce.Name() }
func (o *ObjectBase) GetCe() *ClassEntry { return o.ce }

func (o *ObjectBase) Free() {}

func (o *ObjectBase) Dtor() {}

func (o *ObjectBase) CanClone() bool { return false }
func (o *ObjectBase) Clone() *ZendObject {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) ReadProperty(member *Zval, typ int, cacheSlot *any, rv *Zval) *Zval {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) WriteProperty(member *Zval, value *Zval, cacheSlot *any) *Zval {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) HasProperty(member *Zval, hasSetExists int, cacheSlot *any) int {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) UnsetPropertyEx(member *Zval, cacheSlot *any) {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) GetPropertyPtr(member *Zval, typ int, cacheSlot *any) *Zval {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) IsStdGetProperties() bool { return false }
func (o *ObjectBase) GetPropertiesArray() *Array {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CanGetPropertiesFor() bool { return false }
func (o *ObjectBase) GetPropertiesFor(purpose zend.ZendPropPurpose) *Array {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CanGet() bool { return false }
func (o *ObjectBase) Get(rv *Zval) *Zval {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CanSet() bool { return false }
func (o *ObjectBase) Set(value *Zval) {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) ReadDimension(offset *Zval, typ int, rv *Zval) *Zval {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) WriteDimension(offset *Zval, value *Zval) {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) HasDimension(offset *Zval, checkEmpty int) int {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) UnsetDimension(offset *Zval) {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CanCountElements() bool {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CountElements(count *int) int {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CanGetMethod() bool {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) GetMethod(object **ZendObject, method *String, key *Zval) IFunction {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CallMethod(method *String, object *ZendObject, executeData *zend.ZendExecuteData, returnValue *Zval) int {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) GetConstructor(object *ZendObject) IFunction {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CanCast() bool {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) Cast(retval *Zval, type_ ZvalType) int {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CanGetClosure() bool {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) GetClosure(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **ZendObject) int {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CanDoOperation() bool {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) DoOperation(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CanCompareObjectsTo(obj2 *ZendObject) bool {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CompareObjectsTo(another *ZendObject) int {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) CanCompare() bool {
	//TODO implement me
	panic("implement me")
}

func (o *ObjectBase) Compare(result *Zval, op1 *Zval, op2 *Zval) int {
	//TODO implement me
	panic("implement me")
}
