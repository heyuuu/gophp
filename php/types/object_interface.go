package types

import "github.com/heyuuu/gophp/zend"

type IObject interface {
	ClassName() string
	GetCe() *ClassEntry

	Free()
	Dtor()
	CanClone() bool
	Clone() *ZendObject

	// property
	ReadProperty(member *Zval, typ int, cacheSlot *any, rv *Zval) *Zval
	WriteProperty(member *Zval, value *Zval, cacheSlot *any) *Zval
	HasProperty(member *Zval, hasSetExists int, cacheSlot *any) int
	UnsetPropertyEx(member *Zval, cacheSlot *any)
	GetPropertyPtr(member *Zval, typ int, cacheSlot *any) *Zval

	// properties
	IsStdGetProperties() bool
	GetPropertiesArray() *Array
	CanGetPropertiesFor() bool
	GetPropertiesFor(purpose zend.ZendPropPurpose) *Array

	// get & set
	CanGet() bool
	Get(rv *Zval) *Zval
	CanSet() bool
	Set(value *Zval)

	// dimension
	ReadDimension(offset *Zval, typ int, rv *Zval) *Zval
	WriteDimension(offset *Zval, value *Zval)
	HasDimension(offset *Zval, checkEmpty int) int
	UnsetDimension(offset *Zval)

	// elements
	CanCountElements() bool
	CountElements(count *int) int

	// method
	CanGetMethod() bool
	GetMethod(object **ZendObject, method *String, key *Zval) IFunction
	CallMethod(method *String, object *ZendObject, executeData *zend.ZendExecuteData, returnValue *Zval) int
	GetConstructor(object *ZendObject) IFunction
	// cast
	CanCast() bool
	Cast(retval *Zval, type_ ZvalType) int

	// mixed
	CanGetClosure() bool
	GetClosure(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **ZendObject) int
	CanDoOperation() bool
	DoOperation(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int
	CanCompareObjectsTo(obj2 *ZendObject) bool
	CompareObjectsTo(another *ZendObject) int
	CanCompare() bool
	Compare(result *Zval, op1 *Zval, op2 *Zval) int
}
