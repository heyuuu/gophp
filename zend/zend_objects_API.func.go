package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
)

func IS_OBJ_VALID(o __auto__) bool {
	return !(zend_uintptr_t(o) & OBJ_BUCKET_INVALID)
}
func SET_OBJ_INVALID(o *types.ZendObject) *types.ZendObject {
	return (*types.ZendObject)(zend_uintptr_t(o) | OBJ_BUCKET_INVALID)
}
func GET_OBJ_BUCKET_NUMBER(o *types.ZendObject) int { return zend_intptr_t(o) >> 1 }
func SET_OBJ_BUCKET_NUMBER(o *types.ZendObject, n int) {
	o = (*types.ZendObject)(zend_uintptr_t(n)<<1 | OBJ_BUCKET_INVALID)
}
func ZEND_OBJECTS_STORE_ADD_TO_FREE_LIST(h int) {
	SET_OBJ_BUCKET_NUMBER(EG__().GetObjectsStore().GetObjectBuckets()[h], EG__().GetObjectsStore().GetFreeListHead())
	EG__().GetObjectsStore().SetFreeListHead(h)
}
func OBJ_RELEASE(obj *types.ZendObject)               { ZendObjectRelease(obj) }
func ZendObjectStoreCtorFailed(obj *types.ZendObject) { obj.AddGcFlags(types.IS_OBJ_DESTRUCTOR_CALLED) }
func ZendObjectRelease(obj *types.ZendObject) {
	if obj.DelRefcount() == 0 {
		ZendObjectsStoreDel(obj)
		//} else if GC_MAY_LEAK((*types.ZendRefcounted)(obj)) {
		//	GcPossibleRoot((*types.ZendRefcounted)(obj))
	}
}
func ZendObjectPropertiesSize(ce *types.ClassEntry) int {
	return b.SizeOf("zval") * (ce.GetDefaultPropertiesCount() - b.Cond(ce.IsUseGuards(), 0, 1))
}
func ZendObjectAlloc(obj_size int, ce *types.ClassEntry) any {
	var obj any = Emalloc(obj_size + ZendObjectPropertiesSize(ce))

	/* Subtraction of sizeof(zval) is necessary, because zend_object_properties_size() may be
	 * -sizeof(zval), if the object has no properties. */

	memset(obj, 0, obj_size-b.SizeOf("zval"))
	return obj
}
func ZendGetPropertyInfoForSlot(obj *types.ZendObject, slot *types.Zval) *ZendPropertyInfo {
	var table **ZendPropertyInfo = obj.GetCe().GetPropertiesInfoTable()
	var prop_num intPtr = slot - obj.GetPropertiesTable()
	b.Assert(prop_num >= 0 && prop_num < obj.GetCe().GetDefaultPropertiesCount())
	return table[prop_num]
}
func ZendGetTypedPropertyInfoForSlot(obj *types.ZendObject, slot *types.Zval) *ZendPropertyInfo {
	var prop_info *ZendPropertyInfo = ZendGetPropertyInfoForSlot(obj, slot)
	if prop_info != nil && prop_info.GetType() != 0 {
		return prop_info
	}
	return nil
}
func ZendObjectsStoreInit(objects *ZendObjectsStore, init_size uint32) {
	objects.SetObjectBuckets((**types.ZendObject)(Emalloc(init_size * b.SizeOf("zend_object *"))))
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
	EG__().SetIsObjectStoreNoReuse(true)
	if objects.GetTop() > 1 {
		var i uint32
		for i = 1; i < objects.GetTop(); i++ {
			var obj *types.ZendObject = objects.GetObjectBuckets()[i]
			if IS_OBJ_VALID(obj) {
				if (obj.GetGcFlags() & types.IS_OBJ_DESTRUCTOR_CALLED) == 0 {
					obj.AddGcFlags(types.IS_OBJ_DESTRUCTOR_CALLED)
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
		var obj_ptr **types.ZendObject = objects.GetObjectBuckets() + 1
		var end **types.ZendObject = objects.GetObjectBuckets() + objects.GetTop()
		for {
			var obj *types.ZendObject = *obj_ptr
			if IS_OBJ_VALID(obj) {
				obj.AddGcFlags(types.IS_OBJ_DESTRUCTOR_CALLED)
			}
			obj_ptr++
			if obj_ptr == end {
				break
			}
		}
	}
}
func ZendObjectsStoreFreeObjectStorage(objects *ZendObjectsStore, fast_shutdown types.ZendBool) {
	var obj_ptr **types.ZendObject
	var end ***types.ZendObject
	var obj **types.ZendObject
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
				if (obj.GetGcFlags() & types.IS_OBJ_FREE_CALLED) == 0 {
					obj.AddGcFlags(types.IS_OBJ_FREE_CALLED)
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
				if (obj.GetGcFlags() & types.IS_OBJ_FREE_CALLED) == 0 {
					obj.AddGcFlags(types.IS_OBJ_FREE_CALLED)
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
func ZendObjectsStorePutCold(object *types.ZendObject) {
	var handle int
	var new_size uint32 = 2 * EG__().GetObjectsStore().GetSize()
	EG__().GetObjectsStore().SetObjectBuckets((**types.ZendObject)(Erealloc(EG__().GetObjectsStore().GetObjectBuckets(), new_size*b.SizeOf("zend_object *"))))

	/* Assign size after realloc, in case it fails */

	EG__().GetObjectsStore().SetSize(new_size)
	EG__().GetObjectsStore().GetTop()++
	handle = EG__().GetObjectsStore().GetTop() - 1
	object.SetHandle(handle)
	EG__().GetObjectsStore().GetObjectBuckets()[handle] = object
}
func ZendObjectsStorePut(object *types.ZendObject) {
	var handle int

	/* When in shutdown sequence - do not reuse previously freed handles, to make sure
	 * the dtors for newly created objects are called in zend_objects_store_call_destructors() loop
	 */

	if EG__().GetObjectsStore().GetFreeListHead() != -1 && !EG__().IsObjectStoreNoReuse() {
		handle = EG__().GetObjectsStore().GetFreeListHead()
		EG__().GetObjectsStore().SetFreeListHead(GET_OBJ_BUCKET_NUMBER(EG__().GetObjectsStore().GetObjectBuckets()[handle]))
	} else if EG__().GetObjectsStore().GetTop() == EG__().GetObjectsStore().GetSize() {
		ZendObjectsStorePutCold(object)
		return
	} else {
		EG__().GetObjectsStore().GetTop()++
		handle = EG__().GetObjectsStore().GetTop() - 1
	}
	object.SetHandle(handle)
	EG__().GetObjectsStore().GetObjectBuckets()[handle] = object
}
func ZendObjectsStoreDel(object *types.ZendObject) {
	b.Assert(object.GetRefcount() == 0)

	/* GC might have released this object already. */

	if object.GetGcType() == types.IS_NULL {
		return
	}

	/*    Make sure we hold a reference count during the destructor call
	      otherwise, when the destructor ends the storage might be freed
	      when the refcount reaches 0 a second time
	*/

	if (object.GetGcFlags() & types.IS_OBJ_DESTRUCTOR_CALLED) == 0 {
		object.AddGcFlags(types.IS_OBJ_DESTRUCTOR_CALLED)
		if object.GetHandlers().GetDtorObj() != ZendObjectsDestroyObject || object.GetCe().GetDestructor() != nil {
			object.SetRefcount(1)
			object.GetHandlers().GetDtorObj()(object)
			object.DelRefcount()
		}
	}
	if object.GetRefcount() == 0 {
		var handle uint32 = object.GetHandle()
		var ptr any
		b.Assert(EG__().GetObjectsStore().GetObjectBuckets() != nil)
		b.Assert(IS_OBJ_VALID(EG__().GetObjectsStore().GetObjectBuckets()[handle]))
		EG__().GetObjectsStore().GetObjectBuckets()[handle] = SET_OBJ_INVALID(object)
		if (object.GetGcFlags() & types.IS_OBJ_FREE_CALLED) == 0 {
			object.AddGcFlags(types.IS_OBJ_FREE_CALLED)
			object.SetRefcount(1)
			object.GetHandlers().GetFreeObj()(object)
		}
		ptr = (*byte)(object) - object.GetHandlers().GetOffset()
		//GC_REMOVE_FROM_BUFFER(object)
		Efree(ptr)
		ZEND_OBJECTS_STORE_ADD_TO_FREE_LIST(handle)
	}
}
