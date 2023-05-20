package types

import (
	"github.com/heyuuu/gophp/zend"
)

/**
 * ObjectHandlers
 */
type ObjectHandlersSetting = ObjectHandlers
type ObjectHandlers struct {
	Offset  int                      // 指向 Object 的偏移量
	FreeObj func(object *ZendObject) // todo free函数,在释放时若无free_obj则调用
	DtorObj func(object *ZendObject) // todo 析构函数,在释放时调用，优先级高于 freeObj

	CloneObj   func(object *Zval) *ZendObject
	CloneObjEx func(object *ZendObject) *ZendObject

	ReadProperty   func(object *Zval, member *Zval, type_ int, cacheSlot *any, rv *Zval) *Zval
	ReadPropertyEx func(object *ZendObject, member *Zval, type_ int, cacheSlot *any, rv *Zval) *Zval

	WriteProperty   func(object *Zval, member *Zval, value *Zval, cacheSlot *any) *Zval
	WritePropertyEx func(object *ZendObject, member *Zval, value *Zval, cacheSlot *any) *Zval

	ReadDimension   func(object *Zval, offset *Zval, type_ int, rv *Zval) *Zval
	ReadDimensionEx func(object *ZendObject, offset *Zval, type_ int, rv *Zval) *Zval

	WriteDimension   func(object *Zval, offset *Zval, value *Zval)
	WriteDimensionEx func(object *ZendObject, offset *Zval, value *Zval)

	GetPropertyPtrPtr   func(object *Zval, member *Zval, type_ int, cacheSlot *any) *Zval
	GetPropertyPtrPtrEx func(object *ZendObject, member *Zval, type_ int, cacheSlot *any) *Zval

	Get   func(object *Zval, rv *Zval) *Zval
	GetEx func(object *ZendObject, rv *Zval) *Zval

	Set   func(object *Zval, value *Zval)
	SetEx func(object *ZendObject, value *Zval)

	HasProperty   func(object *Zval, member *Zval, hasSetExists int, cacheSlot *any) int
	HasPropertyEx func(object *ZendObject, member *Zval, hasSetExists int, cacheSlot *any) int

	UnsetProperty   func(object *Zval, member *Zval, cacheSlot *any)
	UnsetPropertyEx func(object *ZendObject, member *Zval, cacheSlot *any)

	HasDimension   func(object *Zval, member *Zval, checkEmpty int) int
	HasDimensionEx func(object *ZendObject, member *Zval, checkEmpty int) int

	UnsetDimension   func(object *Zval, offset *Zval)
	UnsetDimensionEx func(object *ZendObject, offset *Zval)

	GetProperties   func(object *Zval) *Array
	GetPropertiesEx func(object *ZendObject) *Array

	GetPropertiesFor   func(object *Zval, purpose zend.ZendPropPurpose) *Array
	GetPropertiesForEx func(object *ZendObject, purpose zend.ZendPropPurpose) *Array

	GetMethod      func(object **ZendObject, method *String, key *Zval) IFunction
	CallMethod     func(method *String, object *ZendObject, executeData *zend.ZendExecuteData, returnValue *Zval) int
	GetConstructor func(object *ZendObject) IFunction

	CompareObjects   func(object1 *Zval, object2 *Zval) int
	CompareObjectsEx func(object1 *ZendObject, object2 *ZendObject) int

	CastObject   func(readobj *Zval, retval *Zval, type_ ZvalType) int
	CastObjectEx func(readobj *ZendObject, retval *Zval, type_ ZvalType) int

	CountElements   func(object *Zval, count *int) int
	CountElementsEx func(object *ZendObject, count *int) int

	GetClosure   func(obj *Zval, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **ZendObject) int
	GetClosureEx func(obj *ZendObject, cePtr **ClassEntry, fptrPtr *IFunction, objPtr **ZendObject) int

	DoOperation func(opcode uint8, result *Zval, op1 *Zval, op2 *Zval) int
	Compare     func(result *Zval, op1 *Zval, op2 *Zval) int
}

func NewObjectHandlers(s ObjectHandlersSetting) *ObjectHandlers { return &s }
func NewObjectHandlersEx(base *ObjectHandlers, s ObjectHandlersSetting) *ObjectHandlers {
	// todo settings 覆盖 base 产生新 handlers，后续用接口替换
	panic("todo")
}
