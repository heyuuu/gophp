// <<generate>>

package zend

const ZEND_WRONG_PROPERTY_INFO *ZendPropertyInfo = (*ZendPropertyInfo)(intptr_t - 1)
const ZEND_DYNAMIC_PROPERTY_OFFSET = uintptr_t(intptr_t)(-1)

type ZendObjectReadPropertyT func(object *Zval, member *Zval, type_ int, cache_slot *any, rv *Zval) *Zval
type ZendObjectReadDimensionT func(object *Zval, offset *Zval, type_ int, rv *Zval) *Zval
type ZendObjectWritePropertyT func(object *Zval, member *Zval, value *Zval, cache_slot *any) *Zval
type ZendObjectWriteDimensionT func(object *Zval, offset *Zval, value *Zval)
type ZendObjectGetPropertyPtrPtrT func(object *Zval, member *Zval, type_ int, cache_slot *any) *Zval
type ZendObjectSetT func(object *Zval, value *Zval)
type ZendObjectGetT func(object *Zval, rv *Zval) *Zval
type ZendObjectHasPropertyT func(object *Zval, member *Zval, has_set_exists int, cache_slot *any) int
type ZendObjectHasDimensionT func(object *Zval, member *Zval, check_empty int) int
type ZendObjectUnsetPropertyT func(object *Zval, member *Zval, cache_slot *any)
type ZendObjectUnsetDimensionT func(object *Zval, offset *Zval)
type ZendObjectGetPropertiesT func(object *Zval) *HashTable
type ZendObjectGetDebugInfoT func(object *Zval, is_temp *int) *HashTable
type ZendPropPurpose = int

const (
	ZEND_PROP_PURPOSE_DEBUG = iota
	ZEND_PROP_PURPOSE_ARRAY_CAST
	ZEND_PROP_PURPOSE_SERIALIZE
	ZEND_PROP_PURPOSE_VAR_EXPORT
	ZEND_PROP_PURPOSE_JSON
	_ZEND_PROP_PURPOSE_ARRAY_KEY_EXISTS
	_ZEND_PROP_PURPOSE_NON_EXHAUSTIVE_ENUM
)

type ZendObjectGetPropertiesForT func(object *Zval, purpose ZendPropPurpose) *ZendArray
type ZendObjectCallMethodT func(method *ZendString, object *ZendObject, execute_data *ZendExecuteData, return_value *Zval) int
type ZendObjectGetMethodT func(object **ZendObject, method *ZendString, key *Zval) *ZendFunction
type ZendObjectGetConstructorT func(object *ZendObject) *ZendFunction
type ZendObjectDtorObjT func(object *ZendObject)
type ZendObjectFreeObjT func(object *ZendObject)
type ZendObjectCloneObjT func(object *Zval) *ZendObject
type ZendObjectGetClassNameT func(object *ZendObject) *ZendString
type ZendObjectCompareT func(object1 *Zval, object2 *Zval) int
type ZendObjectCompareZvalsT func(result *Zval, op1 *Zval, op2 *Zval) int
type ZendObjectCastT func(readobj *Zval, retval *Zval, type_ int) int
type ZendObjectCountElementsT func(object *Zval, count *ZendLong) int
type ZendObjectGetClosureT func(obj *Zval, ce_ptr **ZendClassEntry, fptr_ptr **ZendFunction, obj_ptr **ZendObject) int
type ZendObjectGetGcT func(object *Zval, table **Zval, n *int) *HashTable
type ZendObjectDoOperationT func(opcode ZendUchar, result *Zval, op1 *Zval, op2 *Zval) int

const ZEND_PROPERTY_ISSET = 0x0
const ZEND_PROPERTY_NOT_EMPTY = ZEND_ISEMPTY
const ZEND_PROPERTY_EXISTS = 0x2
const DEBUG_OBJECT_HANDLERS = 0
const ZEND_WRONG_PROPERTY_OFFSET = 0
const IN_GET = 1 << 0
const IN_SET = 1 << 1
const IN_UNSET = 1 << 2
const IN_ISSET = 1 << 3

var StdObjectHandlers ZendObjectHandlers = ZendObjectHandlers{0, ZendObjectStdDtor, ZendObjectsDestroyObject, ZendObjectsCloneObj, ZendStdReadProperty, ZendStdWriteProperty, ZendStdReadDimension, ZendStdWriteDimension, ZendStdGetPropertyPtrPtr, nil, nil, ZendStdHasProperty, ZendStdUnsetProperty, ZendStdHasDimension, ZendStdUnsetDimension, ZendStdGetProperties, ZendStdGetMethod, nil, ZendStdGetConstructor, ZendStdGetClassName, ZendStdCompareObjects, ZendStdCastObjectTostring, nil, ZendStdGetDebugInfo, ZendStdGetClosure, ZendStdGetGc, nil, nil, nil}
