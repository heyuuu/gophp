package types

import (
	"github.com/heyuuu/gophp/zend"
)

/**
 * ObjectHandlers
 */
type ObjectHandlersSetting = ObjectHandlers
type ObjectHandlers struct {
	Offset  int                  // 指向 Object 的偏移量
	FreeObj func(object *Object) // todo free函数,在释放时若无free_obj则调用
	DtorObj func(object *Object) // todo 析构函数,在释放时调用，优先级高于 freeObj

	CloneObj   func(object *Zval) *Object
	CloneObjEx func(object *Object) *Object

	ReadProperty   func(object *Zval, member *Zval, type_ int, cacheSlot *any, rv *Zval) *Zval
	ReadPropertyEx func(object *Object, member *Zval, type_ int, cacheSlot *any, rv *Zval) *Zval

	WriteProperty   func(object *Zval, member *Zval, value *Zval, cacheSlot *any) *Zval
	WritePropertyEx func(object *Object, member *Zval, value *Zval, cacheSlot *any) *Zval

	ReadDimension   func(object *Zval, offset *Zval, type_ int, rv *Zval) *Zval
	ReadDimensionEx func(object *Object, offset *Zval, type_ int, rv *Zval) *Zval

	WriteDimension   func(object *Zval, offset *Zval, value *Zval)
	WriteDimensionEx func(object *Object, offset *Zval, value *Zval)

	GetPropertyPtrPtr   func(object *Zval, member *Zval, type_ int, cacheSlot *any) *Zval
	GetPropertyPtrPtrEx func(object *Object, member *Zval, type_ int, cacheSlot *any) *Zval

	Get   func(object *Zval, rv *Zval) *Zval
	GetEx func(object *Object, rv *Zval) *Zval

	Set   func(object *Zval, value *Zval)
	SetEx func(object *Object, value *Zval)

	HasProperty   func(object *Zval, member *Zval, hasSetExists int, cacheSlot *any) int
	HasPropertyEx func(object *Object, member *Zval, hasSetExists int, cacheSlot *any) int

	UnsetProperty   func(object *Zval, member *Zval, cacheSlot *any)
	UnsetPropertyEx func(object *Object, member *Zval, cacheSlot *any)

	HasDimension   func(object *Zval, member *Zval, checkEmpty int) int
	HasDimensionEx func(object *Object, member *Zval, checkEmpty int) int

	UnsetDimension   func(object *Zval, offset *Zval)
	UnsetDimensionEx func(object *Object, offset *Zval)

	GetProperties   func(object *Zval) *Array
	GetPropertiesEx func(object *Object) *Array

	GetPropertiesFor   func(object *Zval, purpose zend.ZendPropPurpose) *Array
	GetPropertiesForEx func(object *Object, purpose zend.ZendPropPurpose) *Array

	GetMethod      func(object **Object, method *String, key *Zval) IFunction
	CallMethod     func(method *String, object *Object, executeData *zend.ZendExecuteData, returnValue *Zval) int
	GetConstructor func(object *Object) IFunction

	CompareObjects   func(object1 *Zval, object2 *Zval) int
	CompareObjectsEx func(object1 *Object, object2 *Object) int

	CastObject   func(readobj *Zval, retval *Zval, type_ ZvalType) int
	CastObjectEx func(readobj *Object, retval *Zval, type_ ZvalType) int

	CountElements   func(object *Zval, count *int) int
	CountElementsEx func(object *Object, count *int) int

	GetClosure   func(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **Object) int
	GetClosureEx func(obj *Object, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **Object) int

	DoOperation func(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int
	Compare     func(result *Zval, op1 *Zval, op2 *Zval) int
}

func NewObjectHandlers(s ObjectHandlersSetting) *ObjectHandlers { return &s }
func NewObjectHandlersEx(base *ObjectHandlers, s ObjectHandlersSetting) *ObjectHandlers {
	// todo settings 覆盖 base 产生新 handlers，后续用接口替换
	panic("todo")
}
