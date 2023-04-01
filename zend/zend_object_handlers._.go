package zend

import "github.com/heyuuu/gophp/zend/types"

const ZEND_WRONG_PROPERTY_INFO *ZendPropertyInfo = (*ZendPropertyInfo)(intptr_t - 1)
const ZEND_DYNAMIC_PROPERTY_OFFSET = uintPtr(intptr_t)(-1)

/* The following rule applies to read_property() and read_dimension() implementations:
   If you return a zval which is not otherwise referenced by the extension or the engine's
   symbol table, its reference count should be 0.
*/

type ZendObjectReadPropertyT func(object *types.Zval, member *types.Zval, type_ int, cache_slot *any, rv *types.Zval) *types.Zval

/* Used to fetch dimension from the object, read-only */

type ZendObjectReadDimensionT func(object *types.Zval, offset *types.Zval, type_ int, rv *types.Zval) *types.Zval

/* The following rule applies to write_property() and write_dimension() implementations:
   If you receive a value zval in write_property/write_dimension, you may only modify it if
   its reference count is 1.  Otherwise, you must create a copy of that zval before making
   any changes.  You should NOT modify the reference count of the value passed to you.
   You must return the final value of the assigned property.
*/

type ZendObjectWritePropertyT func(object *types.Zval, member *types.Zval, value *types.Zval, cache_slot *any) *types.Zval

/* Used to set dimension of the object */

type ZendObjectWriteDimensionT func(object *types.Zval, offset *types.Zval, value *types.Zval)

/* Used to create pointer to the property of the object, for future direct r/w access */

type ZendObjectGetPropertyPtrPtrT func(object *types.Zval, member *types.Zval, type_ int, cache_slot *any) *types.Zval

/* Used to set object value. Can be used to override assignments and scalar
   write ops (like ++, +=) on the object */

type ZendObjectSetT func(object *types.Zval, value *types.Zval)

/* Used to get object value. Can be used when converting object value to
 * one of the basic types and when using scalar ops (like ++, +=) on the object
 */

type ZendObjectGetT func(object *types.Zval, rv *types.Zval) *types.Zval

/* Used to check if a property of the object exists */

type ZendObjectHasPropertyT func(object *types.Zval, member *types.Zval, has_set_exists int, cache_slot *any) int

/* Used to check if a dimension of the object exists */

type ZendObjectHasDimensionT func(object *types.Zval, member *types.Zval, check_empty int) int

/* Used to remove a property of the object */

type ZendObjectUnsetPropertyT func(object *types.Zval, member *types.Zval, cache_slot *any)

/* Used to remove a dimension of the object */

type ZendObjectUnsetDimensionT func(object *types.Zval, offset *types.Zval)

/* Used to get hash of the properties of the object, as hash of zval's */

type ZendObjectGetPropertiesT func(object *types.Zval) *types.Array
type ZendObjectGetDebugInfoT func(object *types.Zval, is_temp *int) *types.Array
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

/* The return value must be released using zend_release_properties(). */

type ZendObjectGetPropertiesForT func(object *types.Zval, purpose ZendPropPurpose) *types.Array

/* Used to call methods */

type ZendObjectCallMethodT func(method *types.String, object *types.ZendObject, executeData *ZendExecuteData, return_value *types.Zval) int
type ZendObjectGetMethodT func(object **types.ZendObject, method *types.String, key *types.Zval) types.IFunction
type ZendObjectGetConstructorT func(object *types.ZendObject) types.IFunction

/* Object maintenance/destruction */

type ZendObjectDtorObjT func(object *types.ZendObject)
type ZendObjectFreeObjT func(object *types.ZendObject)
type ZendObjectCloneObjT func(object *types.Zval) *types.ZendObject

/* Get class name for display in var_dump and other debugging functions.
 * Must be defined and must return a non-NULL value. */

type ZendObjectGetClassNameT func(object *types.ZendObject) *types.String
type ZendObjectCompareT func(object1 *types.Zval, object2 *types.Zval) int
type ZendObjectCompareZvalsT func(result *types.Zval, op1 *types.Zval, op2 *types.Zval) int

/* Cast an object to some other type.
 * readobj and retval must point to distinct zvals.
 */

type ZendObjectCastT func(readobj *types.Zval, retval *types.Zval, type_ int) int

/* updates *count to hold the number of elements present and returns SUCCESS.
 * Returns FAILURE if the object does not have any sense of overloaded dimensions */

type ZendObjectCountElementsT func(object *types.Zval, count *ZendLong) int
type ZendObjectGetClosureT func(obj *types.Zval, ce_ptr **types.ClassEntry, fptr_ptr *types.IFunction, obj_ptr **types.ZendObject) int
type ZendObjectGetGcT func(object *types.Zval, table **types.Zval, n *int) *types.Array
type ZendObjectDoOperationT func(opcode types.ZendUchar, result *types.Zval, op1 *types.Zval, op2 *types.Zval) int

const ZEND_PROPERTY_ISSET = 0x0
const ZEND_PROPERTY_NOT_EMPTY = ZEND_ISEMPTY
const ZEND_PROPERTY_EXISTS = 0x2

/* Default behavior for get_properties_for. For use as a fallback in custom
 * get_properties_for implementations. */

/* Will call get_properties_for handler or use default behavior. For use by
 * consumers of the get_properties_for API. */

const DEBUG_OBJECT_HANDLERS = 0
const ZEND_WRONG_PROPERTY_OFFSET = 0

/* guard flags */

const IN_GET = 1 << 0
const IN_SET = 1 << 1
const IN_UNSET = 1 << 2
const IN_ISSET = 1 << 3

/*
  __X accessors explanation:

  if we have __get and property that is not part of the properties array is
  requested, we call __get handler. If it fails, we return uninitialized.

  if we have __set and property that is not part of the properties array is
  set, we call __set handler. If it fails, we do not change the array.

  for both handlers above, when we are inside __get/__set, no further calls for
  __get/__set for this property of this object will be made, to prevent endless
  recursion and enable accessors to change properties array.

  if we have __call and method which is not part of the class function table is
  called, we cal __call handler.
*/

var StdObjectHandlers ZendObjectHandlers = MakeZendObjectHandlers(0, ZendObjectStdDtor, ZendObjectsDestroyObject, ZendObjectsCloneObj, ZendStdReadProperty, ZendStdWriteProperty, ZendStdReadDimension, ZendStdWriteDimension, ZendStdGetPropertyPtrPtr, nil, nil, ZendStdHasProperty, ZendStdUnsetProperty, ZendStdHasDimension, ZendStdUnsetDimension, ZendStdGetProperties, ZendStdGetMethod, nil, ZendStdGetConstructor, ZendStdGetClassName, ZendStdCompareObjects, ZendStdCastObjectTostring, nil, ZendStdGetDebugInfo, ZendStdGetClosure, ZendStdGetGc, nil, nil, nil)
