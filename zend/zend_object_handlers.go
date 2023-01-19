// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_object_handlers.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define ZEND_OBJECT_HANDLERS_H

// #define ZEND_WRONG_PROPERTY_INFO       ( ( struct _zend_property_info * ) ( ( intptr_t ) - 1 ) )

// #define ZEND_DYNAMIC_PROPERTY_OFFSET       ( ( uintptr_t ) ( intptr_t ) ( - 1 ) )

// #define IS_VALID_PROPERTY_OFFSET(offset) ( ( intptr_t ) ( offset ) > 0 )

// #define IS_WRONG_PROPERTY_OFFSET(offset) ( ( intptr_t ) ( offset ) == 0 )

// #define IS_DYNAMIC_PROPERTY_OFFSET(offset) ( ( intptr_t ) ( offset ) < 0 )

// #define IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(offset) ( offset == ZEND_DYNAMIC_PROPERTY_OFFSET )

// #define ZEND_DECODE_DYN_PROP_OFFSET(offset) ( ( uintptr_t ) ( - ( intptr_t ) ( offset ) - 2 ) )

// #define ZEND_ENCODE_DYN_PROP_OFFSET(offset) ( ( uintptr_t ) ( - ( ( intptr_t ) ( offset ) + 2 ) ) )

/* The following rule applies to read_property() and read_dimension() implementations:
   If you return a zval which is not otherwise referenced by the extension or the engine's
   symbol table, its reference count should be 0.
*/

type ZendObjectReadPropertyT func(object *Zval, member *Zval, type_ int, cache_slot *any, rv *Zval) *Zval

/* Used to fetch dimension from the object, read-only */

type ZendObjectReadDimensionT func(object *Zval, offset *Zval, type_ int, rv *Zval) *Zval

/* The following rule applies to write_property() and write_dimension() implementations:
   If you receive a value zval in write_property/write_dimension, you may only modify it if
   its reference count is 1.  Otherwise, you must create a copy of that zval before making
   any changes.  You should NOT modify the reference count of the value passed to you.
   You must return the final value of the assigned property.
*/

type ZendObjectWritePropertyT func(object *Zval, member *Zval, value *Zval, cache_slot *any) *Zval

/* Used to set dimension of the object */

type ZendObjectWriteDimensionT func(object *Zval, offset *Zval, value *Zval)

/* Used to create pointer to the property of the object, for future direct r/w access */

type ZendObjectGetPropertyPtrPtrT func(object *Zval, member *Zval, type_ int, cache_slot *any) *Zval

/* Used to set object value. Can be used to override assignments and scalar
   write ops (like ++, +=) on the object */

type ZendObjectSetT func(object *Zval, value *Zval)

/* Used to get object value. Can be used when converting object value to
 * one of the basic types and when using scalar ops (like ++, +=) on the object
 */

type ZendObjectGetT func(object *Zval, rv *Zval) *Zval

/* Used to check if a property of the object exists */

type ZendObjectHasPropertyT func(object *Zval, member *Zval, has_set_exists int, cache_slot *any) int

/* Used to check if a dimension of the object exists */

type ZendObjectHasDimensionT func(object *Zval, member *Zval, check_empty int) int

/* Used to remove a property of the object */

type ZendObjectUnsetPropertyT func(object *Zval, member *Zval, cache_slot *any)

/* Used to remove a dimension of the object */

type ZendObjectUnsetDimensionT func(object *Zval, offset *Zval)

/* Used to get hash of the properties of the object, as hash of zval's */

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

/* The return value must be released using zend_release_properties(). */

type ZendObjectGetPropertiesForT func(object *Zval, purpose ZendPropPurpose) *ZendArray

/* Used to call methods */

type ZendObjectCallMethodT func(method *ZendString, object *ZendObject, execute_data *ZendExecuteData, return_value *Zval) int
type ZendObjectGetMethodT func(object **ZendObject, method *ZendString, key *Zval) *ZendFunction
type ZendObjectGetConstructorT func(object *ZendObject) *ZendFunction

/* Object maintenance/destruction */

type ZendObjectDtorObjT func(object *ZendObject)
type ZendObjectFreeObjT func(object *ZendObject)
type ZendObjectCloneObjT func(object *Zval) *ZendObject

/* Get class name for display in var_dump and other debugging functions.
 * Must be defined and must return a non-NULL value. */

type ZendObjectGetClassNameT func(object *ZendObject) *ZendString
type ZendObjectCompareT func(object1 *Zval, object2 *Zval) int
type ZendObjectCompareZvalsT func(result *Zval, op1 *Zval, op2 *Zval) int

/* Cast an object to some other type.
 * readobj and retval must point to distinct zvals.
 */

type ZendObjectCastT func(readobj *Zval, retval *Zval, type_ int) int

/* updates *count to hold the number of elements present and returns SUCCESS.
 * Returns FAILURE if the object does not have any sense of overloaded dimensions */

type ZendObjectCountElementsT func(object *Zval, count *ZendLong) int
type ZendObjectGetClosureT func(obj *Zval, ce_ptr **ZendClassEntry, fptr_ptr **ZendFunction, obj_ptr **ZendObject) int
type ZendObjectGetGcT func(object *Zval, table **Zval, n *int) *HashTable
type ZendObjectDoOperationT func(opcode ZendUchar, result *Zval, op1 *Zval, op2 *Zval) int

// @type ZendObjectHandlers struct

// #define zend_get_std_object_handlers() ( & std_object_handlers )

// #define zend_get_function_root_class(fbc) ( ( fbc ) -> common . prototype ? ( fbc ) -> common . prototype -> common . scope : ( fbc ) -> common . scope )

// #define ZEND_PROPERTY_ISSET       0x0

// #define ZEND_PROPERTY_NOT_EMPTY       ZEND_ISEMPTY

// #define ZEND_PROPERTY_EXISTS       0x2

/* Default behavior for get_properties_for. For use as a fallback in custom
 * get_properties_for implementations. */

/* Will call get_properties_for handler or use default behavior. For use by
 * consumers of the get_properties_for API. */

// #define zend_release_properties(ht) do { if ( ( ht ) && ! ( GC_FLAGS ( ht ) & GC_IMMUTABLE ) && ! GC_DELREF ( ht ) ) { zend_array_destroy ( ht ) ; } } while ( 0 )

// #define zend_free_trampoline(func) do { if ( ( func ) == & EG ( trampoline ) ) { EG ( trampoline ) . common . function_name = NULL ; } else { efree ( func ) ; } } while ( 0 )

// Source: <Zend/zend_object_handlers.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_globals.h"

// # include "zend_variables.h"

// # include "zend_API.h"

// # include "zend_objects.h"

// # include "zend_objects_API.h"

// # include "zend_object_handlers.h"

// # include "zend_interfaces.h"

// # include "zend_exceptions.h"

// # include "zend_closures.h"

// # include "zend_compile.h"

// # include "zend_hash.h"

// #define DEBUG_OBJECT_HANDLERS       0

// #define ZEND_WRONG_PROPERTY_OFFSET       0

/* guard flags */

// #define IN_GET       ( 1 << 0 )

// #define IN_SET       ( 1 << 1 )

// #define IN_UNSET       ( 1 << 2 )

// #define IN_ISSET       ( 1 << 3 )

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

func RebuildObjectProperties(zobj *ZendObject) {
	if zobj.GetProperties() == nil {
		var prop_info *ZendPropertyInfo
		var ce *ZendClassEntry = zobj.GetCe()
		var flags uint32 = 0
		zobj.SetProperties(_zendNewArray(ce.GetDefaultPropertiesCount()))
		if ce.GetDefaultPropertiesCount() != 0 {
			ZendHashRealInitMixed(zobj.GetProperties())
			for {
				var __ht *HashTable = &ce.properties_info
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					prop_info = _z.GetValue().GetPtr()
					if (prop_info.GetFlags() & 1 << 4) == 0 {
						flags |= prop_info.GetFlags()
						if (*Zval)((*byte)(zobj+prop_info.GetOffset())).GetType() == 0 {
							zobj.GetProperties().SetUFlags(zobj.GetProperties().GetUFlags() | 1<<5)
						}
						_zendHashAppendInd(zobj.GetProperties(), prop_info.GetName(), (*Zval)((*byte)(zobj+prop_info.GetOffset())))
					}
				}
				break
			}
			if (flags & 1 << 3) != 0 {
				for ce.parent && ce.parent.default_properties_count {
					ce = ce.parent
					for {
						var __ht *HashTable = &ce.properties_info
						var _p *Bucket = __ht.GetArData()
						var _end *Bucket = _p + __ht.GetNNumUsed()
						for ; _p != _end; _p++ {
							var _z *Zval = &_p.val

							if _z.GetType() == 0 {
								continue
							}
							prop_info = _z.GetValue().GetPtr()
							if prop_info.GetCe() == ce && (prop_info.GetFlags()&1<<4) == 0 && (prop_info.GetFlags()&1<<2) != 0 {
								var zv Zval
								if (*Zval)((*byte)(zobj+prop_info.GetOffset())).GetType() == 0 {
									zobj.GetProperties().SetUFlags(zobj.GetProperties().GetUFlags() | 1<<5)
								}
								&zv.GetValue().SetZv((*Zval)((*byte)(zobj + prop_info.GetOffset())))
								&zv.SetTypeInfo(13)
								ZendHashAdd(zobj.GetProperties(), prop_info.GetName(), &zv)
							}
						}
						break
					}
				}
			}
		}
	}
}

/* }}} */

func ZendStdGetProperties(object *Zval) *HashTable {
	var zobj *ZendObject
	zobj = object.GetValue().GetObj()
	if zobj.GetProperties() == nil {
		RebuildObjectProperties(zobj)
	}
	return zobj.GetProperties()
}

/* }}} */

func ZendStdGetGc(object *Zval, table **Zval, n *int) *HashTable {
	if object.GetValue().GetObj().GetHandlers().GetGetProperties() != ZendStdGetProperties {
		*table = nil
		*n = 0
		return object.GetValue().GetObj().GetHandlers().GetGetProperties()(object)
	} else {
		var zobj *ZendObject = object.GetValue().GetObj()
		if zobj.GetProperties() != nil {
			*table = nil
			*n = 0
			if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 && (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo())&1<<6) == 0 {
				ZendGcDelref(&(zobj.GetProperties()).gc)
				zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
			}
			return zobj.GetProperties()
		} else {
			*table = zobj.GetPropertiesTable()
			*n = zobj.GetCe().GetDefaultPropertiesCount()
			return nil
		}
	}
}

/* }}} */

func ZendStdGetDebugInfo(object *Zval, is_temp *int) *HashTable {
	var ce *ZendClassEntry = object.GetValue().GetObj().GetCe()
	var retval Zval
	var ht *HashTable
	if ce.GetDebugInfo() == nil {
		*is_temp = 0
		return object.GetValue().GetObj().GetHandlers().GetGetProperties()(object)
	}
	ZendCallMethod(object, ce, &ce.__debugInfo, "__debuginfo", g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1, &retval, 0, nil, nil)
	if retval.GetType() == 7 {
		if retval.GetTypeFlags() == 0 {
			*is_temp = 1
			return ZendArrayDup(retval.GetValue().GetArr())
		} else if ZvalRefcountP(&retval) <= 1 {
			*is_temp = 1
			ht = retval.GetValue().GetArr()
			return ht
		} else {
			*is_temp = 0
			ZvalPtrDtor(&retval)
			return retval.GetValue().GetArr()
		}
	} else if retval.GetType() == 1 {
		*is_temp = 1
		ht = _zendNewArray(0)
		return ht
	}
	ZendErrorNoreturn(1<<0, "__debuginfo"+"() must return an array")
	return nil
}

/* }}} */

func ZendStdCallGetter(zobj *ZendObject, prop_name *ZendString, retval *Zval) {
	var ce *ZendClassEntry = zobj.GetCe()
	var orig_fake_scope *ZendClassEntry = EG.GetFakeScope()
	var fci ZendFcallInfo
	var fcic ZendFcallInfoCache
	var member Zval
	EG.SetFakeScope(nil)

	/* __get handler is called with one argument:
	      property name

	   it should return whether the call was successful or not
	*/

	var __z *Zval = &member
	var __s *ZendString = prop_name
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	fci.SetSize(g.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(retval)
	fci.SetParamCount(1)
	fci.SetParams(&member)
	fci.SetNoSeparation(1)
	&fci.function_name.u1.type_info = 0
	fcic.SetFunctionHandler(ce.GetGet())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	EG.SetFakeScope(orig_fake_scope)
}

/* }}} */

func ZendStdCallSetter(zobj *ZendObject, prop_name *ZendString, value *Zval) {
	var ce *ZendClassEntry = zobj.GetCe()
	var orig_fake_scope *ZendClassEntry = EG.GetFakeScope()
	var fci ZendFcallInfo
	var fcic ZendFcallInfoCache
	var args []Zval
	var ret Zval
	EG.SetFakeScope(nil)

	/* __set handler is called with two arguments:
	   property name
	   value to be set
	*/

	var __z *Zval = &args[0]
	var __s *ZendString = prop_name
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	var _z1 *Zval = &args[1]
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	&ret.SetTypeInfo(0)
	fci.SetSize(g.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(&ret)
	fci.SetParamCount(2)
	fci.SetParams(args)
	fci.SetNoSeparation(1)
	&fci.function_name.u1.type_info = 0
	fcic.SetFunctionHandler(ce.GetSet())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	ZvalPtrDtor(&ret)
	EG.SetFakeScope(orig_fake_scope)
}

/* }}} */

func ZendStdCallUnsetter(zobj *ZendObject, prop_name *ZendString) {
	var ce *ZendClassEntry = zobj.GetCe()
	var orig_fake_scope *ZendClassEntry = EG.GetFakeScope()
	var fci ZendFcallInfo
	var fcic ZendFcallInfoCache
	var ret Zval
	var member Zval
	EG.SetFakeScope(nil)

	/* __unset handler is called with one argument:
	   property name
	*/

	var __z *Zval = &member
	var __s *ZendString = prop_name
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	&ret.SetTypeInfo(0)
	fci.SetSize(g.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(&ret)
	fci.SetParamCount(1)
	fci.SetParams(&member)
	fci.SetNoSeparation(1)
	&fci.function_name.u1.type_info = 0
	fcic.SetFunctionHandler(ce.GetUnset())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	ZvalPtrDtor(&ret)
	EG.SetFakeScope(orig_fake_scope)
}

/* }}} */

func ZendStdCallIssetter(zobj *ZendObject, prop_name *ZendString, retval *Zval) {
	var ce *ZendClassEntry = zobj.GetCe()
	var orig_fake_scope *ZendClassEntry = EG.GetFakeScope()
	var fci ZendFcallInfo
	var fcic ZendFcallInfoCache
	var member Zval
	EG.SetFakeScope(nil)

	/* __isset handler is called with one argument:
	      property name

	   it should return whether the property is set or not
	*/

	var __z *Zval = &member
	var __s *ZendString = prop_name
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	fci.SetSize(g.SizeOf("fci"))
	fci.SetObject(zobj)
	fci.SetRetval(retval)
	fci.SetParamCount(1)
	fci.SetParams(&member)
	fci.SetNoSeparation(1)
	&fci.function_name.u1.type_info = 0
	fcic.SetFunctionHandler(ce.GetIsset())
	fcic.SetCalledScope(ce)
	fcic.SetObject(zobj)
	ZendCallFunction(&fci, &fcic)
	EG.SetFakeScope(orig_fake_scope)
}

/* }}} */

func IsDerivedClass(child_class *ZendClassEntry, parent_class *ZendClassEntry) ZendBool {
	child_class = child_class.parent
	for child_class != nil {
		if child_class == parent_class {
			return 1
		}
		child_class = child_class.parent
	}
	return 0
}

/* }}} */

func IsProtectedCompatibleScope(ce *ZendClassEntry, scope *ZendClassEntry) int {
	return scope != nil && (IsDerivedClass(ce, scope) != 0 || IsDerivedClass(scope, ce) != 0)
}

/* }}} */

func ZendGetParentPrivateProperty(scope *ZendClassEntry, ce *ZendClassEntry, member *ZendString) *ZendPropertyInfo {
	var zv *Zval
	var prop_info *ZendPropertyInfo
	if scope != ce && scope != nil && IsDerivedClass(ce, scope) != 0 {
		zv = ZendHashFind(&scope.properties_info, member)
		if zv != nil {
			prop_info = (*ZendPropertyInfo)(zv.GetValue().GetPtr())
			if (prop_info.GetFlags()&1<<2) != 0 && prop_info.GetCe() == scope {
				return prop_info
			}
		}
	}
	return nil
}

/* }}} */

func ZendBadPropertyAccess(property_info *ZendPropertyInfo, ce *ZendClassEntry, member *ZendString) {
	ZendThrowError(nil, "Cannot access %s property %s::$%s", ZendVisibilityString(property_info.GetFlags()), ce.GetName().GetVal(), member.GetVal())
}

/* }}} */

func ZendBadPropertyName() {
	ZendThrowError(nil, "Cannot access property started with '\\0'")
}

/* }}} */

func ZendGetPropertyOffset(ce *ZendClassEntry, member *ZendString, silent int, cache_slot *any, info_ptr **ZendPropertyInfo) uintPtr {
	var zv *Zval
	var property_info *ZendPropertyInfo
	var flags uint32
	var scope *ZendClassEntry
	var offset uintPtr
	if cache_slot != nil && ce == cache_slot[0] {
		*info_ptr = (cache_slot + 2)[0]
		return uintptr_t(cache_slot + 1)[0]
	}
	if &ce.properties_info.nNumOfElements == 0 || g.Assign(&zv, ZendHashFind(&ce.properties_info, member)) == nil {
		if member.GetVal()[0] == '0' && member.GetLen() != 0 {
			if silent == 0 {
				ZendBadPropertyName()
			}
			return 0
		}
	dynamic:
		if cache_slot != nil {
			cache_slot[0] = ce
			cache_slot[1] = any(uintptr_t(intptr_t)(-1))
			(cache_slot + 2)[0] = nil
		}
		return uintptr_t(intptr_t)(-1)
	}
	property_info = (*ZendPropertyInfo)(zv.GetValue().GetPtr())
	flags = property_info.GetFlags()
	if (flags & (1<<3 | 1<<2 | 1<<1)) != 0 {
		if EG.GetFakeScope() != nil {
			scope = EG.GetFakeScope()
		} else {
			scope = ZendGetExecutedScope()
		}
		if property_info.GetCe() != scope {
			if (flags & 1 << 3) != 0 {
				var p *ZendPropertyInfo = ZendGetParentPrivateProperty(scope, ce, member)

				/* If there is a public/protected instance property on ce, don't try to use a
				 * private static property on scope. If both are static, prefer the static
				 * property on scope. This will throw a static property notice, rather than
				 * a visibility error. */

				if p != nil && ((p.GetFlags()&1<<4) == 0 || (flags&1<<4) != 0) {
					property_info = p
					flags = property_info.GetFlags()
					goto found
				} else if (flags & 1 << 0) != 0 {
					goto found
				}

				/* If there is a public/protected instance property on ce, don't try to use a
				 * private static property on scope. If both are static, prefer the static
				 * property on scope. This will throw a static property notice, rather than
				 * a visibility error. */

			}
			if (flags & 1 << 2) != 0 {
				if property_info.GetCe() != ce {
					goto dynamic
				} else {
				wrong:

					/* Information was available, but we were denied access.  Error out. */

					if silent == 0 {
						ZendBadPropertyAccess(property_info, ce, member)
					}
					return 0
				}
			} else {
				assert((flags & 1 << 1) != 0)
				if IsProtectedCompatibleScope(property_info.GetCe(), scope) == 0 {
					goto wrong
				}
			}
		}
	}
found:
	if (flags & 1 << 4) != 0 {
		if silent == 0 {
			ZendError(1<<3, "Accessing static property %s::$%s as non static", ce.GetName().GetVal(), member.GetVal())
		}
		return uintptr_t(intptr_t)(-1)
	}
	offset = property_info.GetOffset()
	if property_info.GetType() == 0 {
		property_info = nil
	} else {
		*info_ptr = property_info
	}
	if cache_slot != nil {
		cache_slot[0] = ce
		cache_slot[1] = any(uintPtr(offset))
		(cache_slot + 2)[0] = property_info
	}
	return offset
}

/* }}} */

func ZendWrongOffset(ce *ZendClassEntry, member *ZendString) {
	var dummy *ZendPropertyInfo

	/* Trigger the correct error */

	ZendGetPropertyOffset(ce, member, 0, nil, &dummy)

	/* Trigger the correct error */
}

/* }}} */

func ZendGetPropertyInfo(ce *ZendClassEntry, member *ZendString, silent int) *ZendPropertyInfo {
	var zv *Zval
	var property_info *ZendPropertyInfo
	var flags uint32
	var scope *ZendClassEntry
	if &ce.properties_info.nNumOfElements == 0 || g.Assign(&zv, ZendHashFind(&ce.properties_info, member)) == nil {
		if member.GetVal()[0] == '0' && member.GetLen() != 0 {
			if silent == 0 {
				ZendBadPropertyName()
			}
			return (*ZendPropertyInfo)(intptr_t - 1)
		}
	dynamic:
		return nil
	}
	property_info = (*ZendPropertyInfo)(zv.GetValue().GetPtr())
	flags = property_info.GetFlags()
	if (flags & (1<<3 | 1<<2 | 1<<1)) != 0 {
		if EG.GetFakeScope() != nil {
			scope = EG.GetFakeScope()
		} else {
			scope = ZendGetExecutedScope()
		}
		if property_info.GetCe() != scope {
			if (flags & 1 << 3) != 0 {
				var p *ZendPropertyInfo = ZendGetParentPrivateProperty(scope, ce, member)
				if p != nil {
					property_info = p
					flags = property_info.GetFlags()
					goto found
				} else if (flags & 1 << 0) != 0 {
					goto found
				}
			}
			if (flags & 1 << 2) != 0 {
				if property_info.GetCe() != ce {
					goto dynamic
				} else {
				wrong:

					/* Information was available, but we were denied access.  Error out. */

					if silent == 0 {
						ZendBadPropertyAccess(property_info, ce, member)
					}
					return (*ZendPropertyInfo)(intptr_t - 1)
				}
			} else {
				assert((flags & 1 << 1) != 0)
				if IsProtectedCompatibleScope(property_info.GetCe(), scope) == 0 {
					goto wrong
				}
			}
		}
	}
found:
	if (flags & 1 << 4) != 0 {
		if silent == 0 {
			ZendError(1<<3, "Accessing static property %s::$%s as non static", ce.GetName().GetVal(), member.GetVal())
		}
	}
	return property_info
}

/* }}} */

func ZendCheckPropertyAccess(zobj *ZendObject, prop_info_name *ZendString, is_dynamic ZendBool) int {
	var property_info *ZendPropertyInfo
	var class_name *byte = nil
	var prop_name *byte
	var member *ZendString
	var prop_name_len int
	if prop_info_name.GetVal()[0] == 0 {
		if is_dynamic != 0 {
			return SUCCESS
		}
		ZendUnmanglePropertyNameEx(prop_info_name, &class_name, &prop_name, &prop_name_len)
		member = ZendStringInit(prop_name, prop_name_len, 0)
		property_info = ZendGetPropertyInfo(zobj.GetCe(), member, 1)
		ZendStringReleaseEx(member, 0)
		if property_info == nil || property_info == (*ZendPropertyInfo)(intptr_t-1) {
			return FAILURE
		}
		if class_name[0] != '*' {
			if (property_info.GetFlags() & 1 << 2) == 0 {

				/* we we're looking for a private prop but found a non private one of the same name */

				return FAILURE

				/* we we're looking for a private prop but found a non private one of the same name */

			} else if strcmp(prop_info_name.GetVal()+1, property_info.GetName().GetVal()+1) {

				/* we we're looking for a private prop but found a private one of the same name but another class */

				return FAILURE

				/* we we're looking for a private prop but found a private one of the same name but another class */

			}
		} else {
			assert((property_info.GetFlags() & 1 << 1) != 0)
		}
		return SUCCESS
	} else {
		property_info = ZendGetPropertyInfo(zobj.GetCe(), prop_info_name, 1)
		if property_info == nil {
			assert(is_dynamic != 0)
			return SUCCESS
		} else if property_info == (*ZendPropertyInfo)(intptr_t-1) {
			return FAILURE
		}
		if (property_info.GetFlags() & 1 << 0) != 0 {
			return SUCCESS
		} else {
			return FAILURE
		}
	}
}

/* }}} */

func ZendPropertyGuardDtor(el *Zval) {
	var ptr *uint32 = (*uint32)(el.GetValue().GetPtr())
	if (ZendUintptrT(ptr) & 1) == 0 {
		_efree(ptr)
	}
}

/* }}} */

func ZendGetPropertyGuard(zobj *ZendObject, member *ZendString) *uint32 {
	var guards *HashTable
	var zv *Zval
	var ptr *uint32
	assert((zobj.GetCe().GetCeFlags() & 1 << 11) != 0)
	zv = zobj.GetPropertiesTable() + zobj.GetCe().GetDefaultPropertiesCount()
	if zv.GetType() == 6 {
		var str *ZendString = zv.GetValue().GetStr()
		if str == member || str.GetH() == ZendStringHashVal(member) && ZendStringEqualContent(str, member) != 0 {
			return &(*zv).u2.property_guard
		} else if zv.GetPropertyGuard() == 0 {
			ZvalPtrDtorStr(zv)
			var __z *Zval = zv
			var __s *ZendString = member
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			return &(*zv).u2.property_guard
		} else {
			guards = (*HashTable)(_emalloc(g.SizeOf("HashTable")))
			_zendHashInit(guards, 8, ZendPropertyGuardDtor, 0)

			/* mark pointer as "special" using low bit */

			ZendHashAddNewPtr(guards, str, any(zend_uintptr_t&zv.GetPropertyGuard()|1))
			ZvalPtrDtorStr(zv)
			var __arr *ZendArray = guards
			var __z *Zval = zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
	} else if zv.GetType() == 7 {
		guards = zv.GetValue().GetArr()
		assert(guards != nil)
		zv = ZendHashFind(guards, member)
		if zv != nil {
			return (*uint32)(zend_uintptr_t(*zv).value.ptr & ^1)
		}
	} else {
		assert(zv.GetType() == 0)
		var __z *Zval = zv
		var __s *ZendString = member
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		zv.SetPropertyGuard(0)
		return &(*zv).u2.property_guard
	}

	/* we have to allocate uint32_t separately because ht->arData may be reallocated */

	ptr = (*uint32)(_emalloc(g.SizeOf("uint32_t")))
	*ptr = 0
	return (*uint32)(ZendHashAddNewPtr(guards, member, ptr))
}

/* }}} */

func ZendStdReadProperty(object *Zval, member *Zval, type_ int, cache_slot *any, rv *Zval) *Zval {
	var zobj *ZendObject
	var name *ZendString
	var tmp_name *ZendString
	var retval *Zval
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	var guard *uint32 = nil
	zobj = object.GetValue().GetObj()
	name = ZvalTryGetTmpString(member, &tmp_name)
	if name == nil {
		return &EG.uninitialized_zval
	}

	/* make zend_get_property_info silent if we have getter - we may want to use it */

	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, type_ == 3 || zobj.GetCe().GetGet() != nil, cache_slot, &prop_info)
	if intptr_t(property_offset) > 0 {
		retval = (*Zval)((*byte)(zobj + property_offset))
		if retval.GetType() != 0 {
			goto exit
		}
		if retval.GetU2Extra() == 1 {

			/* Skip __get() for uninitialized typed properties */

			goto uninit_error

			/* Skip __get() for uninitialized typed properties */

		}
	} else if intptr_t(property_offset) < 0 {
		if zobj.GetProperties() != nil {
			if property_offset != uintptr_t(intptr_t)(-1) {
				var idx uintPtr = uintptr_t(-(intptr_t(property_offset)) - 2)
				if idx < zobj.GetProperties().GetNNumUsed()*g.SizeOf("Bucket") {
					var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
					if p.GetVal().GetType() != 0 && (p.GetKey() == name || p.GetH() == name.GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), name) != 0) {
						retval = &p.val
						goto exit
					}
				}
				(cache_slot + 1)[0] = any(uintptr_t(intptr_t)(-1))
			}
			retval = ZendHashFind(zobj.GetProperties(), name)
			if retval != nil {
				if cache_slot != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					(cache_slot + 1)[0] = any(uintptr_t(-(intptr_t(idx) + 2)))
				}
				goto exit
			}
		}
	} else if EG.GetException() != nil {
		retval = &EG.uninitialized_zval
		goto exit
	}

	/* magic isset */

	if type_ == 3 && zobj.GetCe().GetIsset() != nil {
		var tmp_result Zval
		guard = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & 1 << 3) == 0 {
			if tmp_name == nil && (ZvalGcFlags(name.GetGc().GetTypeInfo())&1<<6) == 0 {
				tmp_name = ZendStringCopy(name)
			}
			ZendGcAddref(&zobj.gc)
			&tmp_result.SetTypeInfo(0)
			*guard |= 1 << 3
			ZendStdCallIssetter(zobj, name, &tmp_result)
			*guard &= ^(1 << 3)
			if ZendIsTrue(&tmp_result) == 0 {
				retval = &EG.uninitialized_zval
				ZendObjectRelease(zobj)
				ZvalPtrDtor(&tmp_result)
				goto exit
			}
			ZvalPtrDtor(&tmp_result)
			if zobj.GetCe().GetGet() != nil && ((*guard)&1<<0) == 0 {
				goto call_getter
			}
			ZendObjectRelease(zobj)
		} else if zobj.GetCe().GetGet() != nil && ((*guard)&1<<0) == 0 {
			goto call_getter_addref
		}
	} else if zobj.GetCe().GetGet() != nil {

		/* magic get */

		guard = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & 1 << 0) == 0 {

			/* have getter - try with it! */

		call_getter_addref:
			ZendGcAddref(&zobj.gc)
		call_getter:
			*guard |= 1 << 0
			ZendStdCallGetter(zobj, name, rv)
			*guard &= ^(1 << 0)
			if rv.GetType() != 0 {
				retval = rv
				if rv.GetType() != 10 && (type_ == 1 || type_ == 2 || type_ == 5) {
					if rv.GetType() != 8 {
						ZendError(1<<3, "Indirect modification of overloaded property %s::$%s has no effect", zobj.GetCe().GetName().GetVal(), name.GetVal())
					}
				}
			} else {
				retval = &EG.uninitialized_zval
			}
			if prop_info != nil {
				ZendVerifyPropAssignableByRef(prop_info, retval, (zobj.GetCe().GetGet().GetFnFlags()&1<<31) != 0)
			}
			ZendObjectRelease(zobj)
			goto exit
		} else if intptr_t(property_offset) == 0 {

			/* Trigger the correct error */

			ZendGetPropertyOffset(zobj.GetCe(), name, 0, nil, &prop_info)
			assert(EG.GetException() != nil)
			retval = &EG.uninitialized_zval
			goto exit
		}
	}
uninit_error:
	if type_ != 3 {
		if prop_info != nil {
			ZendThrowError(nil, "Typed property %s::$%s must not be accessed before initialization", prop_info.GetCe().GetName().GetVal(), name.GetVal())
		} else {
			ZendError(1<<3, "Undefined property: %s::$%s", zobj.GetCe().GetName().GetVal(), name.GetVal())
		}
	}
	retval = &EG.uninitialized_zval
exit:
	ZendTmpStringRelease(tmp_name)
	return retval
}

/* }}} */

func PropertyUsesStrictTypes() ZendBool {
	var execute_data *ZendExecuteData = EG.GetCurrentExecuteData()
	return execute_data != nil && execute_data.GetFunc() != nil && (EG.GetCurrentExecuteData().GetFunc().GetFnFlags()&1<<31) != 0
}
func ZendStdWriteProperty(object *Zval, member *Zval, value *Zval, cache_slot *any) *Zval {
	var zobj *ZendObject
	var name *ZendString
	var tmp_name *ZendString
	var variable_ptr *Zval
	var tmp Zval
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	assert(value.GetType() != 10)
	zobj = object.GetValue().GetObj()
	name = ZvalTryGetTmpString(member, &tmp_name)
	if name == nil {
		return value
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, zobj.GetCe().GetSet() != nil, cache_slot, &prop_info)
	if intptr_t(property_offset) > 0 {
		variable_ptr = (*Zval)((*byte)(zobj + property_offset))
		if variable_ptr.GetType() != 0 {
			if value.GetTypeFlags() != 0 {
				ZvalAddrefP(value)
			}
			if prop_info != nil {
				var _z1 *Zval = &tmp
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if ZendVerifyPropertyType(prop_info, &tmp, PropertyUsesStrictTypes()) == 0 {
					if value.GetTypeFlags() != 0 {
						ZvalDelrefP(value)
					}
					variable_ptr = &EG.error_zval
					goto exit
				}
				value = &tmp
			}
		found:
			variable_ptr = ZendAssignToVariable(variable_ptr, value, 1<<1, PropertyUsesStrictTypes())
			goto exit
		}
		if variable_ptr.GetU2Extra() == 1 {

			/* Writes to uninitializde typed properties bypass __set(). */

			variable_ptr.SetU2Extra(0)
			goto write_std_property
		}
	} else if intptr_t(property_offset) < 0 {
		if zobj.GetProperties() != nil {
			if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
				if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
					ZendGcDelref(&(zobj.GetProperties()).gc)
				}
				zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
			}
			if g.Assign(&variable_ptr, ZendHashFind(zobj.GetProperties(), name)) != nil {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
				goto found
			}
		}
	} else if EG.GetException() != nil {
		variable_ptr = &EG.error_zval
		goto exit
	}

	/* magic set */

	if zobj.GetCe().GetSet() != nil {
		var guard *uint32 = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & 1 << 1) == 0 {
			ZendGcAddref(&zobj.gc)
			*guard |= 1 << 1
			ZendStdCallSetter(zobj, name, value)
			*guard &= ^(1 << 1)
			ZendObjectRelease(zobj)
			variable_ptr = value
		} else if intptr_t(property_offset) != 0 {
			goto write_std_property
		} else {

			/* Trigger the correct error */

			ZendWrongOffset(zobj.GetCe(), name)
			assert(EG.GetException() != nil)
			variable_ptr = &EG.error_zval
			goto exit
		}
	} else {
		assert(intptr_t(property_offset) != 0)
	write_std_property:
		if value.GetTypeFlags() != 0 {
			ZvalAddrefP(value)
		}
		if intptr_t(property_offset) > 0 {
			variable_ptr = (*Zval)((*byte)(zobj + property_offset))
			if prop_info != nil {
				var _z1 *Zval = &tmp
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if ZendVerifyPropertyType(prop_info, &tmp, PropertyUsesStrictTypes()) == 0 {
					ZvalPtrDtor(value)
					goto exit
				}
				value = &tmp
				goto found
			}
			var _z1 *Zval = variable_ptr
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else {
			if zobj.GetProperties() == nil {
				RebuildObjectProperties(zobj)
			}
			variable_ptr = ZendHashAddNew(zobj.GetProperties(), name, value)
		}
	}
exit:
	ZendTmpStringRelease(tmp_name)
	return variable_ptr
}

/* }}} */

func ZendBadArrayAccess(ce *ZendClassEntry) {
	ZendThrowError(nil, "Cannot use object of type %s as array", ce.GetName().GetVal())
}

/* }}} */

func ZendStdReadDimension(object *Zval, offset *Zval, type_ int, rv *Zval) *Zval {
	var ce *ZendClassEntry = object.GetValue().GetObj().GetCe()
	var tmp_offset Zval
	var tmp_object Zval
	if InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		if offset == nil {

			/* [] construct */

			&tmp_offset.SetTypeInfo(1)

			/* [] construct */

		} else {
			var _z3 *Zval = offset
			if (_z3.GetTypeInfo() & 0xff00) != 0 {
				if (_z3.GetTypeInfo() & 0xff) == 10 {
					_z3 = &(*_z3).value.GetRef().GetVal()
					if (_z3.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(_z3)
					}
				} else {
					ZvalAddrefP(_z3)
				}
			}
			var _z1 *Zval = &tmp_offset
			var _z2 *Zval = _z3
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
		ZvalAddrefP(object)
		var __z *Zval = &tmp_object
		__z.GetValue().SetObj(object.GetValue().GetObj())
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		if type_ == 3 {
			ZendCallMethod(&tmp_object, ce, nil, "offsetexists", g.SizeOf("\"offsetexists\"")-1, rv, 1, &tmp_offset, nil)
			if rv.GetType() == 0 {
				ZvalPtrDtor(&tmp_object)
				ZvalPtrDtor(&tmp_offset)
				return nil
			}
			if IZendIsTrue(rv) == 0 {
				ZvalPtrDtor(&tmp_object)
				ZvalPtrDtor(&tmp_offset)
				ZvalPtrDtor(rv)
				return &EG.uninitialized_zval
			}
			ZvalPtrDtor(rv)
		}
		ZendCallMethod(&tmp_object, ce, nil, "offsetget", g.SizeOf("\"offsetget\"")-1, rv, 1, &tmp_offset, nil)
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
		if rv.GetType() == 0 {
			if EG.GetException() == nil {
				ZendThrowError(nil, "Undefined offset for object of type %s used as array", ce.GetName().GetVal())
			}
			return nil
		}
		return rv
	} else {
		ZendBadArrayAccess(ce)
		return nil
	}
}

/* }}} */

func ZendStdWriteDimension(object *Zval, offset *Zval, value *Zval) {
	var ce *ZendClassEntry = object.GetValue().GetObj().GetCe()
	var tmp_offset Zval
	var tmp_object Zval
	if InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		if offset == nil {
			&tmp_offset.SetTypeInfo(1)
		} else {
			var _z3 *Zval = offset
			if (_z3.GetTypeInfo() & 0xff00) != 0 {
				if (_z3.GetTypeInfo() & 0xff) == 10 {
					_z3 = &(*_z3).value.GetRef().GetVal()
					if (_z3.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(_z3)
					}
				} else {
					ZvalAddrefP(_z3)
				}
			}
			var _z1 *Zval = &tmp_offset
			var _z2 *Zval = _z3
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
		ZvalAddrefP(object)
		var __z *Zval = &tmp_object
		__z.GetValue().SetObj(object.GetValue().GetObj())
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		ZendCallMethod(&tmp_object, ce, nil, "offsetset", g.SizeOf("\"offsetset\"")-1, nil, 2, &tmp_offset, value)
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
	} else {
		ZendBadArrayAccess(ce)
	}
}

/* }}} */

func ZendStdHasDimension(object *Zval, offset *Zval, check_empty int) int {
	var ce *ZendClassEntry = object.GetValue().GetObj().GetCe()
	var retval Zval
	var tmp_offset Zval
	var tmp_object Zval
	var result int
	if InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		var _z3 *Zval = offset
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = &tmp_offset
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		ZvalAddrefP(object)
		var __z *Zval = &tmp_object
		__z.GetValue().SetObj(object.GetValue().GetObj())
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		ZendCallMethod(&tmp_object, ce, nil, "offsetexists", g.SizeOf("\"offsetexists\"")-1, &retval, 1, &tmp_offset, nil)
		result = IZendIsTrue(&retval)
		ZvalPtrDtor(&retval)
		if check_empty != 0 && result != 0 && EG.GetException() == nil {
			ZendCallMethod(&tmp_object, ce, nil, "offsetget", g.SizeOf("\"offsetget\"")-1, &retval, 1, &tmp_offset, nil)
			result = IZendIsTrue(&retval)
			ZvalPtrDtor(&retval)
		}
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
	} else {
		ZendBadArrayAccess(ce)
		return 0
	}
	return result
}

/* }}} */

func ZendStdGetPropertyPtrPtr(object *Zval, member *Zval, type_ int, cache_slot *any) *Zval {
	var zobj *ZendObject
	var name *ZendString
	var tmp_name *ZendString
	var retval *Zval = nil
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	zobj = object.GetValue().GetObj()
	name = ZvalTryGetTmpString(member, &tmp_name)
	if name == nil {
		return &EG.error_zval
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, zobj.GetCe().GetGet() != nil, cache_slot, &prop_info)
	if intptr_t(property_offset) > 0 {
		retval = (*Zval)((*byte)(zobj + property_offset))
		if retval.GetType() == 0 {
			if zobj.GetCe().GetGet() == nil || ((*ZendGetPropertyGuard)(zobj, name)&1<<0) != 0 || prop_info != nil && retval.GetU2Extra() == 1 {
				if type_ == 2 || type_ == 0 {
					if prop_info != nil {
						ZendThrowError(nil, "Typed property %s::$%s must not be accessed before initialization", prop_info.GetCe().GetName().GetVal(), name.GetVal())
						retval = &EG.error_zval
					} else {
						retval.SetTypeInfo(1)
						ZendError(1<<3, "Undefined property: %s::$%s", zobj.GetCe().GetName().GetVal(), name.GetVal())
					}
				}
			} else {

				/* we do have getter - fail and let it try again with usual get/set */

				retval = nil

				/* we do have getter - fail and let it try again with usual get/set */

			}
		}
	} else if intptr_t(property_offset) < 0 {
		if zobj.GetProperties() != nil {
			if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
				if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
					ZendGcDelref(&(zobj.GetProperties()).gc)
				}
				zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
			}
			if g.Assign(&retval, ZendHashFind(zobj.GetProperties(), name)) != nil {
				ZendTmpStringRelease(tmp_name)
				return retval
			}
		}
		if zobj.GetCe().GetGet() == nil || ((*ZendGetPropertyGuard)(zobj, name)&1<<0) != 0 {
			if zobj.GetProperties() == nil {
				RebuildObjectProperties(zobj)
			}
			retval = ZendHashUpdate(zobj.GetProperties(), name, &EG.uninitialized_zval)

			/* Notice is thrown after creation of the property, to avoid EG(std_property_info)
			 * being overwritten in an error handler. */

			if type_ == 2 || type_ == 0 {
				ZendError(1<<3, "Undefined property: %s::$%s", zobj.GetCe().GetName().GetVal(), name.GetVal())
			}

			/* Notice is thrown after creation of the property, to avoid EG(std_property_info)
			 * being overwritten in an error handler. */

		}
	} else if zobj.GetCe().GetGet() == nil {
		retval = &EG.error_zval
	}
	ZendTmpStringRelease(tmp_name)
	return retval
}

/* }}} */

func ZendStdUnsetProperty(object *Zval, member *Zval, cache_slot *any) {
	var zobj *ZendObject
	var name *ZendString
	var tmp_name *ZendString
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	zobj = object.GetValue().GetObj()
	name = ZvalTryGetTmpString(member, &tmp_name)
	if name == nil {
		return
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, zobj.GetCe().GetUnset() != nil, cache_slot, &prop_info)
	if intptr_t(property_offset) > 0 {
		var slot *Zval = (*Zval)((*byte)(zobj + property_offset))
		if slot.GetType() != 0 {
			if slot.GetType() == 10 && slot.GetValue().GetRef().GetSources().GetPtr() != nil {
				if prop_info != nil {
					ZendRefDelTypeSource(&(slot.GetValue().GetRef()).sources, prop_info)
				}
			}
			var tmp Zval
			var _z1 *Zval = &tmp
			var _z2 *Zval = slot
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			slot.SetTypeInfo(0)
			ZvalPtrDtor(&tmp)
			if zobj.GetProperties() != nil {
				zobj.GetProperties().SetUFlags(zobj.GetProperties().GetUFlags() | 1<<5)
			}
			goto exit
		}
		if slot.GetU2Extra() == 1 {

			/* Reset the IS_PROP_UNINIT flag, if it exists and bypass __unset(). */

			slot.SetU2Extra(0)
			goto exit
		}
	} else if intptr_t(property_offset) < 0 && zobj.GetProperties() != nil {
		if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
			if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
				ZendGcDelref(&(zobj.GetProperties()).gc)
			}
			zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
		}
		if ZendHashDel(zobj.GetProperties(), name) != FAILURE {
			goto exit
		}
	} else if EG.GetException() != nil {
		goto exit
	}

	/* magic unset */

	if zobj.GetCe().GetUnset() != nil {
		var guard *uint32 = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & 1 << 2) == 0 {

			/* have unseter - try with it! */

			*guard |= 1 << 2
			ZendStdCallUnsetter(zobj, name)
			*guard &= ^(1 << 2)
		} else if intptr_t(property_offset) == 0 {

			/* Trigger the correct error */

			ZendWrongOffset(zobj.GetCe(), name)
			assert(EG.GetException() != nil)
			goto exit
		}
	}
exit:
	ZendTmpStringRelease(tmp_name)
}

/* }}} */

func ZendStdUnsetDimension(object *Zval, offset *Zval) {
	var ce *ZendClassEntry = object.GetValue().GetObj().GetCe()
	var tmp_offset Zval
	var tmp_object Zval
	if InstanceofFunctionEx(ce, ZendCeArrayaccess, 1) != 0 {
		var _z3 *Zval = offset
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = &tmp_offset
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		ZvalAddrefP(object)
		var __z *Zval = &tmp_object
		__z.GetValue().SetObj(object.GetValue().GetObj())
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		ZendCallMethod(&tmp_object, ce, nil, "offsetunset", g.SizeOf("\"offsetunset\"")-1, nil, 1, &tmp_offset, nil)
		ZvalPtrDtor(&tmp_object)
		ZvalPtrDtor(&tmp_offset)
	} else {
		ZendBadArrayAccess(ce)
	}
}

/* }}} */

func ZendGetParentPrivateMethod(scope *ZendClassEntry, ce *ZendClassEntry, function_name *ZendString) *ZendFunction {
	var func_ *Zval
	var fbc *ZendFunction
	if scope != ce && scope != nil && IsDerivedClass(ce, scope) != 0 {
		func_ = ZendHashFind(&scope.function_table, function_name)
		if func_ != nil {
			fbc = func_.GetValue().GetFunc()
			if (fbc.GetFnFlags()&1<<2) != 0 && fbc.GetScope() == scope {
				return fbc
			}
		}
	}
	return nil
}

/* }}} */

func ZendCheckProtected(ce *ZendClassEntry, scope *ZendClassEntry) int {
	var fbc_scope *ZendClassEntry = ce

	/* Is the context that's calling the function, the same as one of
	 * the function's parents?
	 */

	for fbc_scope != nil {
		if fbc_scope == scope {
			return 1
		}
		fbc_scope = fbc_scope.parent
	}

	/* Is the function's scope the same as our current object context,
	 * or any of the parents of our context?
	 */

	for scope != nil {
		if scope == ce {
			return 1
		}
		scope = scope.parent
	}
	return 0
}

/* }}} */

func ZendGetCallTrampolineFunc(ce *ZendClassEntry, method_name *ZendString, is_static int) *ZendFunction {
	var mname_len int
	var func_ *ZendOpArray
	var fbc *ZendFunction = g.CondF(is_static != 0, func() *ZendFunction { return ce.GetCallstatic() }, func() *ZendFunction { return ce.GetCall() })

	/* We use non-NULL value to avoid useless run_time_cache allocation.
	 * The low bit must be zero, to not be interpreted as a MAP_PTR offset.
	 */

	var dummy any = any(intPtr(2))
	assert(fbc != nil)
	if EG.GetTrampoline().GetFunctionName() == nil {
		func_ = &EG.trampoline.GetOpArray()
	} else {
		func_ = _ecalloc(1, g.SizeOf("zend_op_array"))
	}
	func_.SetType(2)
	func_.GetArgFlags()[0] = 0
	func_.GetArgFlags()[1] = 0
	func_.GetArgFlags()[2] = 0
	func_.SetFnFlags(1<<18 | 1<<0)
	if is_static != 0 {
		func_.SetFnFlags(func_.GetFnFlags() | 1<<4)
	}
	func_.SetOpcodes(&EG.call_trampoline_op)
	func_.SetRunTimeCachePtr((**any)(&dummy))
	func_.SetScope(fbc.GetScope())

	/* reserve space for arguments, local and temporary variables */

	if fbc.GetType() == 2 {
		if fbc.GetOpArray().GetLastVar()+fbc.GetOpArray().GetT() > 2 {
			func_.SetT(fbc.GetOpArray().GetLastVar() + fbc.GetOpArray().GetT())
		} else {
			func_.SetT(2)
		}
	} else {
		func_.SetT(2)
	}
	if fbc.GetType() == 2 {
		func_.SetFilename(fbc.GetOpArray().GetFilename())
	} else {
		func_.SetFilename(ZendEmptyString)
	}
	if fbc.GetType() == 2 {
		func_.SetLineStart(fbc.GetOpArray().GetLineStart())
	} else {
		func_.SetLineStart(0)
	}
	if fbc.GetType() == 2 {
		func_.SetLineEnd(fbc.GetOpArray().GetLineEnd())
	} else {
		func_.SetLineEnd(0)
	}

	//??? keep compatibility for "\0" characters

	if g.Assign(&mname_len, strlen(method_name.GetVal())) != method_name.GetLen() {
		func_.SetFunctionName(ZendStringInit(method_name.GetVal(), mname_len, 0))
	} else {
		func_.SetFunctionName(ZendStringCopy(method_name))
	}
	func_.SetPrototype(nil)
	func_.SetNumArgs(0)
	func_.SetRequiredNumArgs(0)
	func_.SetArgInfo(0)
	return (*ZendFunction)(func_)
}

/* }}} */

func ZendGetUserCallFunction(ce *ZendClassEntry, method_name *ZendString) *ZendFunction {
	return ZendGetCallTrampolineFunc(ce, method_name, 0)
}

/* }}} */

func ZendBadMethodCall(fbc *ZendFunction, method_name *ZendString, scope *ZendClassEntry) {
	ZendThrowError(nil, "Call to %s method %s::%s() from context '%s'", ZendVisibilityString(fbc.GetFnFlags()), g.CondF1(fbc != nil && fbc.GetScope() != nil, func() []byte { return fbc.GetScope().GetName().GetVal() }, ""), method_name.GetVal(), g.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
}

/* }}} */

func ZendStdGetMethod(obj_ptr **ZendObject, method_name *ZendString, key *Zval) *ZendFunction {
	var zobj *ZendObject = *obj_ptr
	var func_ *Zval
	var fbc *ZendFunction
	var lc_method_name *ZendString
	var scope *ZendClassEntry
	if key != nil {
		lc_method_name = key.GetValue().GetStr()
	} else {
		lc_method_name = (*ZendString)(_emalloc(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + method_name.GetLen() + 1 + (8-1) & ^(8-1)))
		ZendGcSetRefcount(&lc_method_name.gc, 1)
		lc_method_name.GetGc().SetTypeInfo(6)
		lc_method_name.SetH(0)
		lc_method_name.SetLen(method_name.GetLen())
		ZendStrTolowerCopy(lc_method_name.GetVal(), method_name.GetVal(), method_name.GetLen())
	}
	if g.Assign(&func_, ZendHashFind(&zobj.ce.GetFunctionTable(), lc_method_name)) == nil {
		if key == nil {
			_efree(lc_method_name)
		}
		if zobj.GetCe().GetCall() != nil {
			return ZendGetUserCallFunction(zobj.GetCe(), method_name)
		} else {
			return nil
		}
	}
	fbc = func_.GetValue().GetFunc()

	/* Check access level */

	if (fbc.GetOpArray().GetFnFlags() & (1<<3 | 1<<2 | 1<<1)) != 0 {
		scope = ZendGetExecutedScope()
		if fbc.GetScope() != scope {
			if (fbc.GetOpArray().GetFnFlags() & 1 << 3) != 0 {
				var updated_fbc *ZendFunction = ZendGetParentPrivateMethod(scope, zobj.GetCe(), lc_method_name)
				if updated_fbc != nil {
					fbc = updated_fbc
					goto exit
				} else if (fbc.GetOpArray().GetFnFlags() & 1 << 0) != 0 {
					goto exit
				}
			}
			if (fbc.GetOpArray().GetFnFlags()&1<<2) != 0 || ZendCheckProtected(g.CondF(fbc.GetPrototype() != nil, func() *ZendClassEntry { return fbc.GetPrototype().GetScope() }, func() *ZendClassEntry { return fbc.GetScope() }), scope) == 0 {
				if zobj.GetCe().GetCall() != nil {
					fbc = ZendGetUserCallFunction(zobj.GetCe(), method_name)
				} else {
					ZendBadMethodCall(fbc, method_name, scope)
					fbc = nil
				}
			}
		}
	}
exit:
	if key == nil {
		_efree(lc_method_name)
	}
	return fbc
}

/* }}} */

func ZendGetUserCallstaticFunction(ce *ZendClassEntry, method_name *ZendString) *ZendFunction {
	return ZendGetCallTrampolineFunc(ce, method_name, 1)
}

/* }}} */

func ZendStdGetStaticMethod(ce *ZendClassEntry, function_name *ZendString, key *Zval) *ZendFunction {
	var fbc *ZendFunction = nil
	var lc_function_name *ZendString
	var object *ZendObject
	var scope *ZendClassEntry
	if key != nil {
		lc_function_name = key.GetValue().GetStr()
	} else {
		lc_function_name = ZendStringTolowerEx(function_name, 0)
	}
	var func_ *Zval = ZendHashFind(&ce.function_table, lc_function_name)
	if func_ != nil {
		fbc = func_.GetValue().GetFunc()
	} else if ce.GetConstructor() != nil && lc_function_name.GetLen() == ce.GetName().GetLen() && ZendBinaryStrncasecmp(lc_function_name.GetVal(), lc_function_name.GetLen(), ce.GetName().GetVal(), lc_function_name.GetLen(), lc_function_name.GetLen()) == 0 && (ce.GetConstructor().GetFunctionName().GetVal()[0] != '_' || ce.GetConstructor().GetFunctionName().GetVal()[1] != '_') {
		fbc = ce.GetConstructor()
	} else {
		if key == nil {
			ZendStringReleaseEx(lc_function_name, 0)
		}
		if ce.GetCall() != nil && g.Assign(&object, ZendGetThisObject(EG.GetCurrentExecuteData())) != nil && InstanceofFunction(object.GetCe(), ce) != 0 {

			/* Call the top-level defined __call().
			 * see: tests/classes/__call_004.phpt  */

			var call_ce *ZendClassEntry = object.GetCe()
			for call_ce.GetCall() == nil {
				call_ce = call_ce.parent
			}
			return ZendGetUserCallFunction(call_ce, function_name)
		} else if ce.GetCallstatic() != nil {
			return ZendGetUserCallstaticFunction(ce, function_name)
		} else {
			return nil
		}
	}
	if (fbc.GetOpArray().GetFnFlags() & 1 << 0) == 0 {
		scope = ZendGetExecutedScope()
		if fbc.GetScope() != scope {
			if (fbc.GetOpArray().GetFnFlags()&1<<2) != 0 || ZendCheckProtected(g.CondF(fbc.GetPrototype() != nil, func() *ZendClassEntry { return fbc.GetPrototype().GetScope() }, func() *ZendClassEntry { return fbc.GetScope() }), scope) == 0 {
				if ce.GetCallstatic() != nil {
					fbc = ZendGetUserCallstaticFunction(ce, function_name)
				} else {
					ZendBadMethodCall(fbc, function_name, scope)
					fbc = nil
				}
			}
		}
	}
	if key == nil {
		ZendStringReleaseEx(lc_function_name, 0)
	}
	return fbc
}

/* }}} */

func ZendClassInitStatics(class_type *ZendClassEntry) {
	var i int
	var p *Zval
	if class_type.GetDefaultStaticMembersCount() != 0 && (*Zval)(g.CondF((uintptr_t(class_type).static_members_table__ptr&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(class_type).static_members_table__ptr - 1)))
	}, func() any { return any(*(class_type.GetStaticMembersTablePtr())) })) == nil {
		if class_type.parent {
			ZendClassInitStatics(class_type.parent)
		}
		if (uintPtr(class_type.GetStaticMembersTablePtr()) & 1) != 0 {
			*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(class_type.GetStaticMembersTablePtr()-1)))) = _emalloc(g.SizeOf("zval") * class_type.GetDefaultStaticMembersCount())
		} else {
			*(class_type.GetStaticMembersTablePtr()) = _emalloc(g.SizeOf("zval") * class_type.GetDefaultStaticMembersCount())
		}
		for i = 0; i < class_type.GetDefaultStaticMembersCount(); i++ {
			p = &class_type.default_static_members_table[i]
			if p.GetType() == 13 {
				var q *Zval = &(*Zval)(g.CondF((uintptr_t(class_type.parent).static_members_table__ptr&1) != 0, func() any {
					return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(class_type.parent).static_members_table__ptr - 1)))
				}, func() any { return any(*(class_type.parent.static_members_table__ptr)) }))[i]
				if q.GetType() == 13 {
					q = q.GetValue().GetZv()
				}
				&(*Zval)(g.CondF((uintptr_t(class_type).static_members_table__ptr&1) != 0, func() any {
					return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(class_type).static_members_table__ptr - 1)))
				}, func() any { return any(*(class_type.GetStaticMembersTablePtr())) }))[i].GetValue().SetZv(q)
				&(*Zval)(g.CondF((uintptr_t(class_type).static_members_table__ptr&1) != 0, func() any {
					return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(class_type).static_members_table__ptr - 1)))
				}, func() any { return any(*(class_type.GetStaticMembersTablePtr())) }))[i].SetTypeInfo(13)
			} else {
				var _z1 *Zval = &(*Zval)(g.CondF((uintptr_t(class_type).static_members_table__ptr&1) != 0, func() any {
					return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(class_type).static_members_table__ptr - 1)))
				}, func() any { return any(*(class_type.GetStaticMembersTablePtr())) }))[i]
				var _z2 *Zval = p
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
						ZendGcAddref(&_gc.gc)
					} else {
						ZvalCopyCtorFunc(_z1)
					}
				}
			}
		}
	}
}
func ZendStdGetStaticPropertyWithInfo(ce *ZendClassEntry, property_name *ZendString, type_ int, property_info_ptr **ZendPropertyInfo) *Zval {
	var ret *Zval
	var scope *ZendClassEntry
	var property_info *ZendPropertyInfo = ZendHashFindPtr(&ce.properties_info, property_name)
	*property_info_ptr = property_info
	if property_info == nil {
		goto undeclared_property
	}
	if (property_info.GetFlags() & 1 << 0) == 0 {
		if EG.GetFakeScope() != nil {
			scope = EG.GetFakeScope()
		} else {
			scope = ZendGetExecutedScope()
		}
		if property_info.GetCe() != scope {
			if (property_info.GetFlags()&1<<2) != 0 || IsProtectedCompatibleScope(property_info.GetCe(), scope) == 0 {
				if type_ != 3 {
					ZendBadPropertyAccess(property_info, ce, property_name)
				}
				return nil
			}
		}
	}
	if (property_info.GetFlags() & 1 << 4) == 0 {
		goto undeclared_property
	}
	if (ce.GetCeFlags() & 1 << 12) == 0 {
		if ZendUpdateClassConstants(ce) != SUCCESS {
			return nil
		}
	}

	/* check if static properties were destroyed */

	if (*Zval)(g.CondF((uintptr_t(ce).static_members_table__ptr&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(ce).static_members_table__ptr - 1)))
	}, func() any { return any(*(ce.GetStaticMembersTablePtr())) })) == nil {
		if ce.GetType() == 1 || (ce.GetCeFlags()&(1<<7|1<<10)) != 0 {
			ZendClassInitStatics(ce)
		} else {
		undeclared_property:
			if type_ != 3 {
				ZendThrowError(nil, "Access to undeclared static property: %s::$%s", ce.GetName().GetVal(), property_name.GetVal())
			}
			return nil
		}
	}
	ret = (*Zval)(g.CondF((uintptr_t(ce).static_members_table__ptr&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(ce).static_members_table__ptr - 1)))
	}, func() any { return any(*(ce.GetStaticMembersTablePtr())) })) + property_info.GetOffset()
	if ret.GetType() == 13 {
		ret = ret.GetValue().GetZv()
	}
	if (type_ == 0 || type_ == 2) && ret.GetType() == 0 && property_info.GetType() != 0 {
		ZendThrowError(nil, "Typed static property %s::$%s must not be accessed before initialization", property_info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(property_name))
		return nil
	}
	return ret
}

/* }}} */

func ZendStdGetStaticProperty(ce *ZendClassEntry, property_name *ZendString, type_ int) *Zval {
	var prop_info *ZendPropertyInfo
	return ZendStdGetStaticPropertyWithInfo(ce, property_name, type_, &prop_info)
}
func ZendStdUnsetStaticProperty(ce *ZendClassEntry, property_name *ZendString) ZendBool {
	ZendThrowError(nil, "Attempt to unset static property %s::$%s", ce.GetName().GetVal(), property_name.GetVal())
	return 0
}

/* }}} */

func ZendBadConstructorCall(constructor *ZendFunction, scope *ZendClassEntry) {
	if scope != nil {
		ZendThrowError(nil, "Call to %s %s::%s() from context '%s'", ZendVisibilityString(constructor.GetFnFlags()), constructor.GetScope().GetName().GetVal(), constructor.GetFunctionName().GetVal(), scope.GetName().GetVal())
	} else {
		ZendThrowError(nil, "Call to %s %s::%s() from invalid context", ZendVisibilityString(constructor.GetFnFlags()), constructor.GetScope().GetName().GetVal(), constructor.GetFunctionName().GetVal())
	}
}

/* }}} */

func ZendStdGetConstructor(zobj *ZendObject) *ZendFunction {
	var constructor *ZendFunction = zobj.GetCe().GetConstructor()
	var scope *ZendClassEntry
	if constructor != nil {
		if (constructor.GetOpArray().GetFnFlags() & 1 << 0) == 0 {
			if EG.GetFakeScope() != nil {
				scope = EG.GetFakeScope()
			} else {
				scope = ZendGetExecutedScope()
			}
			if constructor.GetScope() != scope {
				if (constructor.GetOpArray().GetFnFlags()&1<<2) != 0 || ZendCheckProtected(g.CondF(constructor.GetPrototype() != nil, func() *ZendClassEntry { return constructor.GetPrototype().GetScope() }, func() *ZendClassEntry { return constructor.GetScope() }), scope) == 0 {
					ZendBadConstructorCall(constructor, scope)
					constructor = nil
				}
			}
		}
	}
	return constructor
}

/* }}} */

func ZendStdCompareObjects(o1 *Zval, o2 *Zval) int {
	var zobj1 *ZendObject
	var zobj2 *ZendObject
	zobj1 = o1.GetValue().GetObj()
	zobj2 = o2.GetValue().GetObj()
	if zobj1 == zobj2 {
		return 0
	}
	if zobj1.GetCe() != zobj2.GetCe() {
		return 1
	}
	if zobj1.GetProperties() == nil && zobj2.GetProperties() == nil {
		var info *ZendPropertyInfo
		if zobj1.GetCe().GetDefaultPropertiesCount() == 0 {
			return 0
		}

		/* It's enough to protect only one of the objects.
		 * The second one may be referenced from the first and this may cause
		 * false recursion detection.
		 */

		if (ZvalGcFlags(o1.GetValue().GetCounted().GetGc().GetTypeInfo()) & 1 << 5) != 0 {
			ZendErrorNoreturn(1<<0, "Nesting level too deep - recursive dependency?")
		}
		o1.GetValue().GetCounted().GetGc().SetTypeInfo(o1.GetValue().GetCounted().GetGc().GetTypeInfo() | 1<<5<<0)
		for {
			var __ht *HashTable = &zobj1.ce.GetPropertiesInfo()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				info = _z.GetValue().GetPtr()
				var p1 *Zval = (*Zval)((*byte)(zobj1 + info.GetOffset()))
				var p2 *Zval = (*Zval)((*byte)(zobj2 + info.GetOffset()))
				if (info.GetFlags() & 1 << 4) != 0 {
					continue
				}
				if p1.GetType() != 0 {
					if p2.GetType() != 0 {
						var result Zval
						if CompareFunction(&result, p1, p2) == FAILURE {
							o1.GetValue().GetCounted().GetGc().SetTypeInfo(o1.GetValue().GetCounted().GetGc().GetTypeInfo() &^ (1 << 5 << 0))
							return 1
						}
						if result.GetValue().GetLval() != 0 {
							o1.GetValue().GetCounted().GetGc().SetTypeInfo(o1.GetValue().GetCounted().GetGc().GetTypeInfo() &^ (1 << 5 << 0))
							return result.GetValue().GetLval()
						}
					} else {
						o1.GetValue().GetCounted().GetGc().SetTypeInfo(o1.GetValue().GetCounted().GetGc().GetTypeInfo() &^ (1 << 5 << 0))
						return 1
					}
				} else {
					if p2.GetType() != 0 {
						o1.GetValue().GetCounted().GetGc().SetTypeInfo(o1.GetValue().GetCounted().GetGc().GetTypeInfo() &^ (1 << 5 << 0))
						return 1
					}
				}
			}
			break
		}
		o1.GetValue().GetCounted().GetGc().SetTypeInfo(o1.GetValue().GetCounted().GetGc().GetTypeInfo() &^ (1 << 5 << 0))
		return 0
	} else {
		if zobj1.GetProperties() == nil {
			RebuildObjectProperties(zobj1)
		}
		if zobj2.GetProperties() == nil {
			RebuildObjectProperties(zobj2)
		}
		return ZendCompareSymbolTables(zobj1.GetProperties(), zobj2.GetProperties())
	}
}

/* }}} */

func ZendStdHasProperty(object *Zval, member *Zval, has_set_exists int, cache_slot *any) int {
	var zobj *ZendObject
	var result int
	var value *Zval = nil
	var name *ZendString
	var tmp_name *ZendString
	var property_offset uintPtr
	var prop_info *ZendPropertyInfo = nil
	zobj = object.GetValue().GetObj()
	name = ZvalTryGetTmpString(member, &tmp_name)
	if name == nil {
		return 0
	}
	property_offset = ZendGetPropertyOffset(zobj.GetCe(), name, 1, cache_slot, &prop_info)
	if intptr_t(property_offset) > 0 {
		value = (*Zval)((*byte)(zobj + property_offset))
		if value.GetType() != 0 {
			goto found
		}
		if value.GetU2Extra() == 1 {

			/* Skip __isset() for uninitialized typed properties */

			result = 0
			goto exit
		}
	} else if intptr_t(property_offset) < 0 {
		if zobj.GetProperties() != nil {
			if property_offset != uintptr_t(intptr_t)(-1) {
				var idx uintPtr = uintptr_t(-(intptr_t(property_offset)) - 2)
				if idx < zobj.GetProperties().GetNNumUsed()*g.SizeOf("Bucket") {
					var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
					if p.GetVal().GetType() != 0 && (p.GetKey() == name || p.GetH() == name.GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), name) != 0) {
						value = &p.val
						goto found
					}
				}
				(cache_slot + 1)[0] = any(uintptr_t(intptr_t)(-1))
			}
			value = ZendHashFind(zobj.GetProperties(), name)
			if value != nil {
				if cache_slot != nil {
					var idx uintPtr = (*byte)(value - (*byte)(zobj.GetProperties().GetArData()))
					(cache_slot + 1)[0] = any(uintptr_t(-(intptr_t(idx) + 2)))
				}
			found:
				if has_set_exists == 1<<0 {
					result = ZendIsTrue(value)
				} else if has_set_exists < 1<<0 {
					assert(has_set_exists == 0x0)
					if value.GetType() == 10 {
						value = &(*value).value.GetRef().GetVal()
					}
					result = value.GetType() != 1
				} else {
					assert(has_set_exists == 0x2)
					result = 1
				}
				goto exit
			}
		}
	} else if EG.GetException() != nil {
		result = 0
		goto exit
	}
	result = 0
	if has_set_exists != 0x2 && zobj.GetCe().GetIsset() != nil {
		var guard *uint32 = ZendGetPropertyGuard(zobj, name)
		if ((*guard) & 1 << 3) == 0 {
			var rv Zval

			/* have issetter - try with it! */

			if tmp_name == nil && (ZvalGcFlags(name.GetGc().GetTypeInfo())&1<<6) == 0 {
				tmp_name = ZendStringCopy(name)
			}
			ZendGcAddref(&zobj.gc)
			*guard |= 1 << 3
			ZendStdCallIssetter(zobj, name, &rv)
			result = ZendIsTrue(&rv)
			ZvalPtrDtor(&rv)
			if has_set_exists == 1<<0 && result != 0 {
				if EG.GetException() == nil && zobj.GetCe().GetGet() != nil && ((*guard)&1<<0) == 0 {
					*guard |= 1 << 0
					ZendStdCallGetter(zobj, name, &rv)
					*guard &= ^(1 << 0)
					result = IZendIsTrue(&rv)
					ZvalPtrDtor(&rv)
				} else {
					result = 0
				}
			}
			*guard &= ^(1 << 3)
			ZendObjectRelease(zobj)
		}
	}
exit:
	ZendTmpStringRelease(tmp_name)
	return result
}

/* }}} */

func ZendStdGetClassName(zobj *ZendObject) *ZendString {
	return ZendStringCopy(zobj.GetCe().GetName())
}

/* }}} */

func ZendStdCastObjectTostring(readobj *Zval, writeobj *Zval, type_ int) int {
	var retval Zval
	var ce *ZendClassEntry
	switch type_ {
	case 6:
		ce = readobj.GetValue().GetObj().GetCe()
		if ce.GetTostring() != nil {
			var fake_scope *ZendClassEntry = EG.GetFakeScope()
			EG.SetFakeScope(nil)
			ZendCallMethod(readobj, ce, &ce.__tostring, "__tostring", g.SizeOf("\"__tostring\"")-1, &retval, 0, nil, nil)
			EG.SetFakeScope(fake_scope)
			if retval.GetType() == 6 {
				var _z1 *Zval = writeobj
				var _z2 *Zval = &retval
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				return SUCCESS
			}
			ZvalPtrDtor(&retval)
			if EG.GetException() == nil {
				ZendThrowError(nil, "Method %s::__toString() must return a string value", ce.GetName().GetVal())
			}
		}
		return FAILURE
	case 16:
		writeobj.SetTypeInfo(3)
		return SUCCESS
	case 4:
		ce = readobj.GetValue().GetObj().GetCe()
		ZendError(1<<3, "Object of class %s could not be converted to int", ce.GetName().GetVal())
		var __z *Zval = writeobj
		__z.GetValue().SetLval(1)
		__z.SetTypeInfo(4)
		return SUCCESS
	case 5:
		ce = readobj.GetValue().GetObj().GetCe()
		ZendError(1<<3, "Object of class %s could not be converted to float", ce.GetName().GetVal())
		var __z *Zval = writeobj
		__z.GetValue().SetDval(1)
		__z.SetTypeInfo(5)
		return SUCCESS
	case 20:
		ce = readobj.GetValue().GetObj().GetCe()
		ZendError(1<<3, "Object of class %s could not be converted to number", ce.GetName().GetVal())
		var __z *Zval = writeobj
		__z.GetValue().SetLval(1)
		__z.SetTypeInfo(4)
		return SUCCESS
	default:
		writeobj.SetTypeInfo(1)
		break
	}
	return FAILURE
}

/* }}} */

func ZendStdGetClosure(obj *Zval, ce_ptr **ZendClassEntry, fptr_ptr **ZendFunction, obj_ptr **ZendObject) int {
	var func_ *Zval
	var ce *ZendClassEntry = obj.GetValue().GetObj().GetCe()
	if g.Assign(&func_, ZendHashFindEx(&ce.function_table, ZendKnownStrings[ZEND_STR_MAGIC_INVOKE], 1)) == nil {
		return FAILURE
	}
	*fptr_ptr = func_.GetValue().GetFunc()
	*ce_ptr = ce
	if ((*fptr_ptr).GetFnFlags() & 1 << 4) != 0 {
		if obj_ptr != nil {
			*obj_ptr = nil
		}
	} else {
		if obj_ptr != nil {
			*obj_ptr = obj.GetValue().GetObj()
		}
	}
	return SUCCESS
}

/* }}} */

func ZendStdGetPropertiesFor(obj *Zval, purpose ZendPropPurpose) *HashTable {
	var ht *HashTable
	switch purpose {
	case ZEND_PROP_PURPOSE_DEBUG:
		if obj.GetValue().GetObj().GetHandlers().GetGetDebugInfo() != nil {
			var is_temp int
			ht = obj.GetValue().GetObj().GetHandlers().GetGetDebugInfo()(obj, &is_temp)
			if ht != nil && is_temp == 0 && (ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<6) == 0 {
				ZendGcAddref(&ht.gc)
			}
			return ht
		}
	case ZEND_PROP_PURPOSE_ARRAY_CAST:

	case ZEND_PROP_PURPOSE_SERIALIZE:

	case ZEND_PROP_PURPOSE_VAR_EXPORT:

	case ZEND_PROP_PURPOSE_JSON:

	case _ZEND_PROP_PURPOSE_ARRAY_KEY_EXISTS:
		ht = obj.GetValue().GetObj().GetHandlers().GetGetProperties()(obj)
		if ht != nil && (ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<6) == 0 {
			ZendGcAddref(&ht.gc)
		}
		return ht
	default:
		assert(false)
		return nil
	}
}
func ZendGetPropertiesFor(obj *Zval, purpose ZendPropPurpose) *HashTable {
	if obj.GetValue().GetObj().GetHandlers().GetGetPropertiesFor() != nil {
		return obj.GetValue().GetObj().GetHandlers().GetGetPropertiesFor()(obj, purpose)
	}
	return ZendStdGetPropertiesFor(obj, purpose)
}

var StdObjectHandlers ZendObjectHandlers = ZendObjectHandlers{0, ZendObjectStdDtor, ZendObjectsDestroyObject, ZendObjectsCloneObj, ZendStdReadProperty, ZendStdWriteProperty, ZendStdReadDimension, ZendStdWriteDimension, ZendStdGetPropertyPtrPtr, nil, nil, ZendStdHasProperty, ZendStdUnsetProperty, ZendStdHasDimension, ZendStdUnsetDimension, ZendStdGetProperties, ZendStdGetMethod, nil, ZendStdGetConstructor, ZendStdGetClassName, ZendStdCompareObjects, ZendStdCastObjectTostring, nil, ZendStdGetDebugInfo, ZendStdGetClosure, ZendStdGetGc, nil, nil, nil}
