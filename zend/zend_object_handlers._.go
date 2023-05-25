package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

const ZEND_DYNAMIC_PROPERTY_OFFSET = uintPtr(intptr_t)(-1)

/* Used to get hash of the properties of the object, as hash of zval's */

type ZendPropPurpose = int

const (
	ZEND_PROP_PURPOSE_DEBUG = iota
	ZEND_PROP_PURPOSE_ARRAY_CAST
	ZEND_PROP_PURPOSE_SERIALIZE
	ZEND_PROP_PURPOSE_VAR_EXPORT
	ZEND_PROP_PURPOSE_JSON
	ZEND_PROP_PURPOSE_ARRAY_KEY_EXISTS
)

/* Used to call methods */
const ZEND_PROPERTY_ISSET = 0x0
const ZEND_PROPERTY_NOT_EMPTY = ZEND_ISEMPTY
const ZEND_PROPERTY_EXISTS = 0x2

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

var StdObjectHandlersPtr = types.NewObjectHandlers(types.ObjectHandlers{
	Offset:            0,
	FreeObj:           ZendObjectStdDtor,
	DtorObj:           ZendObjectsDestroyObject,
	CloneObjEx:        ZendObjectsCloneObjEx,
	ReadPropertyEx:    ZendStdReadPropertyEx,
	WritePropertyEx:   ZendStdWritePropertyEx,
	ReadDimensionEx:   ZendStdReadDimensionEx,
	WriteDimensionEx:  ZendStdWriteDimensionEx,
	GetPropertyPtrPtr: ZendStdGetPropertyPtrPtr,
	HasPropertyEx:     ZendStdHasPropertyEx,
	UnsetPropertyEx:   ZendStdUnsetPropertyEx,
	HasDimensionEx:    ZendStdHasDimensionEx,
	UnsetDimensionEx:  ZendStdUnsetDimensionEx,
	GetPropertiesEx:   ZendStdGetPropertiesEx,
	GetMethod:         ZendStdGetMethod,
	GetConstructor:    ZendStdGetConstructor,
	CompareObjectsEx:  ZendStdCompareObjectsEx,
	CastObjectEx:      ZendStdCastObject,
	GetClosure:        ZendStdGetClosure,
})
