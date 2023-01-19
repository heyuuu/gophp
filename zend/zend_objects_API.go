// <<generate>>

package zend

import g "sik/runtime/grammar"

// Source: <Zend/zend_objects_API.h>

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

// #define ZEND_OBJECTS_API_H

// # include "zend.h"

// # include "zend_compile.h"

// #define OBJ_BUCKET_INVALID       ( 1 << 0 )

// #define IS_OBJ_VALID(o) ( ! ( ( ( zend_uintptr_t ) ( o ) ) & OBJ_BUCKET_INVALID ) )

// #define SET_OBJ_INVALID(o) ( ( zend_object * ) ( ( ( ( zend_uintptr_t ) ( o ) ) | OBJ_BUCKET_INVALID ) ) )

// #define GET_OBJ_BUCKET_NUMBER(o) ( ( ( zend_intptr_t ) ( o ) ) >> 1 )

// #define SET_OBJ_BUCKET_NUMBER(o,n) do { ( o ) = ( zend_object * ) ( ( ( ( zend_uintptr_t ) ( n ) ) << 1 ) | OBJ_BUCKET_INVALID ) ; } while ( 0 )

// #define ZEND_OBJECTS_STORE_ADD_TO_FREE_LIST(h) do { SET_OBJ_BUCKET_NUMBER ( EG ( objects_store ) . object_buckets [ ( h ) ] , EG ( objects_store ) . free_list_head ) ; EG ( objects_store ) . free_list_head = ( h ) ; } while ( 0 )

// #define OBJ_RELEASE(obj) zend_object_release ( obj )

// @type ZendObjectsStore struct

/* Global store handling functions */

/* Store API functions */

/* Called when the ctor was terminated by an exception */

func ZendObjectStoreCtorFailed(obj *ZendObject) {
	obj.GetGc().SetTypeInfo(obj.GetGc().GetTypeInfo() | 1<<8<<0)
}
func ZendObjectRelease(obj *ZendObject) {
	if ZendGcDelref(&obj.gc) == 0 {
		ZendObjectsStoreDel(obj)
	} else if ((*ZendRefcounted)(obj).GetGc().GetTypeInfo() & (0xfffffc00 | 1<<4<<0)) == 1<<4<<0 {
		GcPossibleRoot((*ZendRefcounted)(obj))
	}
}
func ZendObjectPropertiesSize(ce *ZendClassEntry) int {
	return g.SizeOf("zval") * (ce.GetDefaultPropertiesCount() - g.Cond((ce.GetCeFlags()&1<<11) != 0, 0, 1))
}

/* Allocates object type and zeros it, but not the properties.
 * Properties MUST be initialized using object_properties_init(). */

func ZendObjectAlloc(obj_size int, ce *ZendClassEntry) any {
	var obj any = _emalloc(obj_size + ZendObjectPropertiesSize(ce))

	/* Subtraction of sizeof(zval) is necessary, because zend_object_properties_size() may be
	 * -sizeof(zval), if the object has no properties. */

	memset(obj, 0, obj_size-g.SizeOf("zval"))
	return obj
}
func ZendGetPropertyInfoForSlot(obj *ZendObject, slot *Zval) *ZendPropertyInfo {
	var table **ZendPropertyInfo = obj.GetCe().GetPropertiesInfoTable()
	var prop_num intPtr = slot - obj.GetPropertiesTable()
	assert(prop_num >= 0 && prop_num < obj.GetCe().GetDefaultPropertiesCount())
	return table[prop_num]
}

/* Helper for cases where we're only interested in property info of typed properties. */

func ZendGetTypedPropertyInfoForSlot(obj *ZendObject, slot *Zval) *ZendPropertyInfo {
	var prop_info *ZendPropertyInfo = ZendGetPropertyInfoForSlot(obj, slot)
	if prop_info != nil && prop_info.GetType() != 0 {
		return prop_info
	}
	return nil
}

// Source: <Zend/zend_objects_API.c>

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

// # include "zend_objects_API.h"

func ZendObjectsStoreInit(objects *ZendObjectsStore, init_size uint32) {
	objects.SetObjectBuckets((**ZendObject)(_emalloc(init_size * g.SizeOf("zend_object *"))))
	objects.SetTop(1)
	objects.SetSize(init_size)
	objects.SetFreeListHead(-1)
	memset(&objects.object_buckets[0], 0, g.SizeOf("zend_object *"))
}
func ZendObjectsStoreDestroy(objects *ZendObjectsStore) {
	_efree(objects.GetObjectBuckets())
	objects.SetObjectBuckets(nil)
}
func ZendObjectsStoreCallDestructors(objects *ZendObjectsStore) {
	EG.SetFlags(EG.GetFlags() | 1<<1)
	if objects.GetTop() > 1 {
		var i uint32
		for i = 1; i < objects.GetTop(); i++ {
			var obj *ZendObject = objects.GetObjectBuckets()[i]
			if (zend_uintptr_t(obj) & 1 << 0) == 0 {
				if (ZvalGcFlags(obj.GetGc().GetTypeInfo()) & 1 << 8) == 0 {
					obj.GetGc().SetTypeInfo(obj.GetGc().GetTypeInfo() | 1<<8<<0)
					if obj.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || obj.GetCe().GetDestructor() != nil {
						ZendGcAddref(&obj.gc)
						obj.GetHandlers().GetDtorObj()(obj)
						ZendGcDelref(&obj.gc)
					}
				}
			}
		}
	}
}
func ZendObjectsStoreMarkDestructed(objects *ZendObjectsStore) {
	if objects.GetObjectBuckets() != nil && objects.GetTop() > 1 {
		var obj_ptr **ZendObject = objects.GetObjectBuckets() + 1
		var end **ZendObject = objects.GetObjectBuckets() + objects.GetTop()
		for {
			var obj *ZendObject = *obj_ptr
			if (zend_uintptr_t(obj) & 1 << 0) == 0 {
				obj.GetGc().SetTypeInfo(obj.GetGc().GetTypeInfo() | 1<<8<<0)
			}
			obj_ptr++
			if obj_ptr == end {
				break
			}
		}
	}
}
func ZendObjectsStoreFreeObjectStorage(objects *ZendObjectsStore, fast_shutdown ZendBool) {
	var obj_ptr **ZendObject
	var end ***ZendObject
	var obj **ZendObject
	if objects.GetTop() <= 1 {
		return
	}

	/* Free object contents, but don't free objects themselves, so they show up as leaks.
	 * Also add a ref to all objects, so the object can't be freed by something else later. */

	end = objects.GetObjectBuckets() + 1
	obj_ptr = objects.GetObjectBuckets() + objects.GetTop()
	if fast_shutdown != 0 {
		for {
			obj_ptr--
			obj = *obj_ptr
			if (zend_uintptr_t(obj) & 1 << 0) == 0 {
				if (ZvalGcFlags(obj.gc.GetTypeInfo()) & 1 << 9) == 0 {
					obj.gc.SetTypeInfo(obj.gc.GetTypeInfo() | 1<<9<<0)
					if obj.handlers.GetFreeObj() != ZendObjectStdDtor {
						ZendGcAddref(&obj.gc)
						obj.handlers.GetFreeObj()(obj)
					}
				}
			}
			if obj_ptr == end {
				break
			}
		}
	} else {
		for {
			obj_ptr--
			obj = *obj_ptr
			if (zend_uintptr_t(obj) & 1 << 0) == 0 {
				if (ZvalGcFlags(obj.gc.GetTypeInfo()) & 1 << 9) == 0 {
					obj.gc.SetTypeInfo(obj.gc.GetTypeInfo() | 1<<9<<0)
					ZendGcAddref(&obj.gc)
					obj.handlers.GetFreeObj()(obj)
				}
			}
			if obj_ptr == end {
				break
			}
		}
	}
}

/* Store objects API */

func ZendObjectsStorePutCold(object *ZendObject) {
	var handle int
	var new_size uint32 = 2 * EG.GetObjectsStore().GetSize()
	EG.GetObjectsStore().SetObjectBuckets((**ZendObject)(_erealloc(EG.GetObjectsStore().GetObjectBuckets(), new_size*g.SizeOf("zend_object *"))))

	/* Assign size after realloc, in case it fails */

	EG.GetObjectsStore().SetSize(new_size)
	EG.GetObjectsStore().GetTop()++
	handle = EG.GetObjectsStore().GetTop() - 1
	object.SetHandle(handle)
	EG.GetObjectsStore().GetObjectBuckets()[handle] = object
}
func ZendObjectsStorePut(object *ZendObject) {
	var handle int

	/* When in shutdown sequence - do not reuse previously freed handles, to make sure
	 * the dtors for newly created objects are called in zend_objects_store_call_destructors() loop
	 */

	if EG.GetObjectsStore().GetFreeListHead() != -1 && (EG.GetFlags()&1<<1) == 0 {
		handle = EG.GetObjectsStore().GetFreeListHead()
		EG.GetObjectsStore().SetFreeListHead(zend_intptr_t(EG.GetObjectsStore().GetObjectBuckets()[handle]) >> 1)
	} else if EG.GetObjectsStore().GetTop() == EG.GetObjectsStore().GetSize() {
		ZendObjectsStorePutCold(object)
		return
	} else {
		EG.GetObjectsStore().GetTop()++
		handle = EG.GetObjectsStore().GetTop() - 1
	}
	object.SetHandle(handle)
	EG.GetObjectsStore().GetObjectBuckets()[handle] = object
}
func ZendObjectsStoreDel(object *ZendObject) {
	assert(ZendGcRefcount(&object.gc) == 0)

	/* GC might have released this object already. */

	if ZvalGcType(object.GetGc().GetTypeInfo()) == 1 {
		return
	}

	/*    Make sure we hold a reference count during the destructor call
	      otherwise, when the destructor ends the storage might be freed
	      when the refcount reaches 0 a second time
	*/

	if (ZvalGcFlags(object.GetGc().GetTypeInfo()) & 1 << 8) == 0 {
		object.GetGc().SetTypeInfo(object.GetGc().GetTypeInfo() | 1<<8<<0)
		if object.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || object.GetCe().GetDestructor() != nil {
			ZendGcSetRefcount(&object.gc, 1)
			object.GetHandlers().GetDtorObj()(object)
			ZendGcDelref(&object.gc)
		}
	}
	if ZendGcRefcount(&object.gc) == 0 {
		var handle uint32 = object.GetHandle()
		var ptr any
		assert(EG.GetObjectsStore().GetObjectBuckets() != nil)
		assert((zend_uintptr_t(EG.GetObjectsStore().GetObjectBuckets()[handle]) & 1 << 0) == 0)
		EG.GetObjectsStore().GetObjectBuckets()[handle] = (*ZendObject)(zend_uintptr_t(object) | 1<<0)
		if (ZvalGcFlags(object.GetGc().GetTypeInfo()) & 1 << 9) == 0 {
			object.GetGc().SetTypeInfo(object.GetGc().GetTypeInfo() | 1<<9<<0)
			ZendGcSetRefcount(&object.gc, 1)
			object.GetHandlers().GetFreeObj()(object)
		}
		ptr = (*byte)(object) - object.GetHandlers().GetOffset()
		var _p *ZendRefcounted = (*ZendRefcounted)(object)
		if (_p.GetGc().GetTypeInfo() & 0xfffffc00) != 0 {
			GcRemoveFromBuffer(_p)
		}
		_efree(ptr)
		EG.GetObjectsStore().GetObjectBuckets()[handle] = (*ZendObject)(zend_uintptr_t(EG.GetObjectsStore().GetFreeListHead())<<1 | 1<<0)
		EG.GetObjectsStore().SetFreeListHead(handle)
	}
}

/* }}} */
