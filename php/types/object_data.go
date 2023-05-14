package types

import "github.com/heyuuu/gophp/zend"

type ObjectData interface {
	ClassName() string

	//
	Free()
	Dtor()
	CanClone() bool
	Clone(zv *Zval) *ZendObject

	// property
	ReadProperty(object *Zval, member *Zval, typ int, cacheSlot *any, rv *Zval) *Zval
	WriteProperty(object *Zval, member *Zval, value *Zval, cacheSlot *any) *Zval
	HasProperty(object *Zval, member *Zval, hasSetExists int, cacheSlot *any) int
	UnsetProperty(object *Zval, member *Zval, cacheSlot *any)
	GetPropertyPtr(object *Zval, member *Zval, typ int, cacheSlot *any) *Zval

	// properties
	GetPropertiesArray(object *Zval) *Array
	CanGetPropertiesFor() bool
	GetPropertiesFor(object *Zval, purpose zend.ZendPropPurpose) *Array

	// get & set
	CanGet() bool
	Get(object *Zval, rv *Zval) *Zval

	CanSet() bool
	Set(object *Zval, value *Zval)

	// dimmension
	ReadDimension(object *Zval, offset *Zval, typ int, rv *Zval) *Zval
	WriteDimension(object *Zval, offset *Zval, value *Zval)
	HasDimension(object *Zval, offset *Zval, checkEmpty int) int
	UnsetDimension(object *Zval, offset *Zval)

	// elements
	CanCountElements() bool
	CountElements(object *Zval, count *int) int

	// method
	CanGetMethod() bool
	GetMethod(object **ZendObject, method *String, key *Zval) IFunction
	CallMethod(method *String, object *ZendObject, executeData *zend.ZendExecuteData, returnValue *Zval) int
	GetConstructor(object *ZendObject) IFunction

	// cast
	CanCast() bool
	Cast(readobj *Zval, retval *Zval, type_ ZvalType) int

	// mixed
	CanGetClosure() bool
	GetClosure(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **ZendObject) int

	CanDoOperation() bool
	DoOperation(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int

	CanCompareObjectsTo(obj2 *ZendObject) bool
	CompareObjects(obj1, obj2 *Zval) int

	CanCompare() bool
	Compare(result *Zval, op1 *Zval, op2 *Zval) int
}
