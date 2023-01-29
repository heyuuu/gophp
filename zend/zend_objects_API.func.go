// <<generate>>

package zend

import (
	b "sik/builtin"
)

func IS_OBJ_VALID(o __auto__) bool {
	return !(zend_uintptr_t(o) & OBJ_BUCKET_INVALID)
}
func SET_OBJ_INVALID(o *ZendObject) *ZendObject {
	return (*ZendObject)(zend_uintptr_t(o) | OBJ_BUCKET_INVALID)
}
func GET_OBJ_BUCKET_NUMBER(o *ZendObject) int { return zend_intptr_t(o) >> 1 }
func SET_OBJ_BUCKET_NUMBER(o *ZendObject, n int) {
	o = (*ZendObject)(zend_uintptr_t(n)<<1 | OBJ_BUCKET_INVALID)
}
func ZEND_OBJECTS_STORE_ADD_TO_FREE_LIST(h int) {
	SET_OBJ_BUCKET_NUMBER(__EG().GetObjectsStore().GetObjectBuckets()[h], __EG().GetObjectsStore().GetFreeListHead())
	__EG().GetObjectsStore().SetFreeListHead(h)
}
func OBJ_RELEASE(obj *ZendObject)               { ZendObjectRelease(obj) }
func ZendObjectStoreCtorFailed(obj *ZendObject) { obj.AddGcFlags(IS_OBJ_DESTRUCTOR_CALLED) }
func ZendObjectRelease(obj *ZendObject) {
	if obj.DelRefcount() == 0 {
		ZendObjectsStoreDel(obj)
	} else if GC_MAY_LEAK((*ZendRefcounted)(obj)) {
		GcPossibleRoot((*ZendRefcounted)(obj))
	}
}
func ZendObjectPropertiesSize(ce *ZendClassEntry) int {
	return b.SizeOf("zval") * (ce.GetDefaultPropertiesCount() - b.Cond(ce.IsUseGuards(), 0, 1))
}
func ZendObjectAlloc(obj_size int, ce *ZendClassEntry) any {
	var obj any = Emalloc(obj_size + ZendObjectPropertiesSize(ce))

	/* Subtraction of sizeof(zval) is necessary, because zend_object_properties_size() may be
	 * -sizeof(zval), if the object has no properties. */

	memset(obj, 0, obj_size-b.SizeOf("zval"))
	return obj
}
func ZendGetPropertyInfoForSlot(obj *ZendObject, slot *Zval) *ZendPropertyInfo {
	var table **ZendPropertyInfo = obj.GetCe().GetPropertiesInfoTable()
	var prop_num intPtr = slot - obj.GetPropertiesTable()
	ZEND_ASSERT(prop_num >= 0 && prop_num < obj.GetCe().GetDefaultPropertiesCount())
	return table[prop_num]
}
func ZendGetTypedPropertyInfoForSlot(obj *ZendObject, slot *Zval) *ZendPropertyInfo {
	var prop_info *ZendPropertyInfo = ZendGetPropertyInfoForSlot(obj, slot)
	if prop_info != nil && prop_info.GetType() != 0 {
		return prop_info
	}
	return nil
}
func ZendObjectsStoreInit(objects *ZendObjectsStore, init_size uint32) {
	objects.SetObjectBuckets((**ZendObject)(Emalloc(init_size * b.SizeOf("zend_object *"))))
	objects.SetTop(1)
	objects.SetSize(init_size)
	objects.SetFreeListHead(-1)
	memset(objects.GetObjectBuckets()[0], 0, b.SizeOf("zend_object *"))
}
func ZendObjectsStoreDestroy(objects *ZendObjectsStore) {
	Efree(objects.GetObjectBuckets())
	objects.SetObjectBuckets(nil)
}
func ZendObjectsStoreCallDestructors(objects *ZendObjectsStore) {
	__EG().SetIsObjectStoreNoReuse(true)
	if objects.GetTop() > 1 {
		var i uint32
		for i = 1; i < objects.GetTop(); i++ {
			var obj *ZendObject = objects.GetObjectBuckets()[i]
			if IS_OBJ_VALID(obj) {
				if (obj.GetGcFlags() & IS_OBJ_DESTRUCTOR_CALLED) == 0 {
					obj.AddGcFlags(IS_OBJ_DESTRUCTOR_CALLED)
					if obj.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || obj.GetCe().GetDestructor() != nil {
						obj.AddRefcount()
						obj.GetHandlers().GetDtorObj()(obj)
						obj.DelRefcount()
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
			if IS_OBJ_VALID(obj) {
				obj.AddGcFlags(IS_OBJ_DESTRUCTOR_CALLED)
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
			if IS_OBJ_VALID(obj) {
				if (obj.GetGcFlags() & IS_OBJ_FREE_CALLED) == 0 {
					obj.AddGcFlags(IS_OBJ_FREE_CALLED)
					if obj.GetHandlers().GetFreeObj() != ZendObjectStdDtor {
						obj.AddRefcount()
						obj.GetHandlers().GetFreeObj()(obj)
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
			if IS_OBJ_VALID(obj) {
				if (obj.GetGcFlags() & IS_OBJ_FREE_CALLED) == 0 {
					obj.AddGcFlags(IS_OBJ_FREE_CALLED)
					obj.AddRefcount()
					obj.GetHandlers().GetFreeObj()(obj)
				}
			}
			if obj_ptr == end {
				break
			}
		}
	}
}
func ZendObjectsStorePutCold(object *ZendObject) {
	var handle int
	var new_size uint32 = 2 * __EG().GetObjectsStore().GetSize()
	__EG().GetObjectsStore().SetObjectBuckets((**ZendObject)(Erealloc(__EG().GetObjectsStore().GetObjectBuckets(), new_size*b.SizeOf("zend_object *"))))

	/* Assign size after realloc, in case it fails */

	__EG().GetObjectsStore().SetSize(new_size)
	__EG().GetObjectsStore().GetTop()++
	handle = __EG().GetObjectsStore().GetTop() - 1
	object.SetHandle(handle)
	__EG().GetObjectsStore().GetObjectBuckets()[handle] = object
}
func ZendObjectsStorePut(object *ZendObject) {
	var handle int

	/* When in shutdown sequence - do not reuse previously freed handles, to make sure
	 * the dtors for newly created objects are called in zend_objects_store_call_destructors() loop
	 */

	if __EG().GetObjectsStore().GetFreeListHead() != -1 && !__EG().IsObjectStoreNoReuse() {
		handle = __EG().GetObjectsStore().GetFreeListHead()
		__EG().GetObjectsStore().SetFreeListHead(GET_OBJ_BUCKET_NUMBER(__EG().GetObjectsStore().GetObjectBuckets()[handle]))
	} else if __EG().GetObjectsStore().GetTop() == __EG().GetObjectsStore().GetSize() {
		ZendObjectsStorePutCold(object)
		return
	} else {
		__EG().GetObjectsStore().GetTop()++
		handle = __EG().GetObjectsStore().GetTop() - 1
	}
	object.SetHandle(handle)
	__EG().GetObjectsStore().GetObjectBuckets()[handle] = object
}
func ZendObjectsStoreDel(object *ZendObject) {
	ZEND_ASSERT(object.GetRefcount() == 0)

	/* GC might have released this object already. */

	if object.GetGcType() == IS_NULL {
		return
	}

	/*    Make sure we hold a reference count during the destructor call
	      otherwise, when the destructor ends the storage might be freed
	      when the refcount reaches 0 a second time
	*/

	if (object.GetGcFlags() & IS_OBJ_DESTRUCTOR_CALLED) == 0 {
		object.AddGcFlags(IS_OBJ_DESTRUCTOR_CALLED)
		if object.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || object.GetCe().GetDestructor() != nil {
			object.SetRefcount(1)
			object.GetHandlers().GetDtorObj()(object)
			object.DelRefcount()
		}
	}
	if object.GetRefcount() == 0 {
		var handle uint32 = object.GetHandle()
		var ptr any
		ZEND_ASSERT(__EG().GetObjectsStore().GetObjectBuckets() != nil)
		ZEND_ASSERT(IS_OBJ_VALID(__EG().GetObjectsStore().GetObjectBuckets()[handle]))
		__EG().GetObjectsStore().GetObjectBuckets()[handle] = SET_OBJ_INVALID(object)
		if (object.GetGcFlags() & IS_OBJ_FREE_CALLED) == 0 {
			object.AddGcFlags(IS_OBJ_FREE_CALLED)
			object.SetRefcount(1)
			object.GetHandlers().GetFreeObj()(object)
		}
		ptr = (*byte)(object) - object.GetHandlers().GetOffset()
		GC_REMOVE_FROM_BUFFER(object)
		Efree(ptr)
		ZEND_OBJECTS_STORE_ADD_TO_FREE_LIST(handle)
	}
}
