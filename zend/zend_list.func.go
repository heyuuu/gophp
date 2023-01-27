// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func ZendListInsert(ptr any, type_ int) *Zval {
	var index int
	var zv Zval
	index = ExecutorGlobals.GetRegularList().GetNNextFreeElement()
	if index == 0 {
		index = 1
	} else if index == core.INT_MAX {
		ZendErrorNoreturn(E_ERROR, "Resource ID space overflow")
	}
	ZVAL_NEW_RES(&zv, index, ptr, type_)
	return ExecutorGlobals.GetRegularList().IndexAddNew(index, &zv)
}
func ZendListDelete(res *ZendResource) int {
	if GC_DELREF(res) <= 0 {
		return ExecutorGlobals.GetRegularList().IndexDel(res.GetHandle())
	} else {
		return SUCCESS
	}
}
func ZendListFree(res *ZendResource) int {
	if GC_REFCOUNT(res) <= 0 {
		return ExecutorGlobals.GetRegularList().IndexDel(res.GetHandle())
	} else {
		return SUCCESS
	}
}
func ZendResourceDtor(res *ZendResource) {
	var ld *ZendRsrcListDtorsEntry
	var r ZendResource = *res
	res.SetType(-1)
	res.SetPtr(nil)
	ld = ListDestructors.IndexFindPtr(r.GetType())
	if ld != nil {
		if ld.GetListDtorEx() != nil {
			ld.GetListDtorEx()(&r)
		}
	} else {
		ZendError(E_WARNING, "Unknown list entry type (%d)", r.GetType())
	}
}
func ZendListClose(res *ZendResource) int {
	if GC_REFCOUNT(res) <= 0 {
		return ZendListFree(res)
	} else if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	return SUCCESS
}
func ZendRegisterResource(rsrc_pointer any, rsrc_type int) *ZendResource {
	var zv *Zval
	zv = ZendListInsert(rsrc_pointer, rsrc_type)
	return Z_RES_P(zv)
}
func ZendFetchResource2(res *ZendResource, resource_type_name string, resource_type1 int, resource_type2 int) any {
	if res != nil {
		if resource_type1 == res.GetType() {
			return res.GetPtr()
		}
		if resource_type2 == res.GetType() {
			return res.GetPtr()
		}
	}
	if resource_type_name {
		var space *byte
		var class_name *byte = GetActiveClassName(&space)
		ZendError(E_WARNING, "%s%s%s(): supplied resource is not a valid %s resource", class_name, space, GetActiveFunctionName(), resource_type_name)
	}
	return nil
}
func ZendFetchResource(res *ZendResource, resource_type_name *byte, resource_type int) any {
	if resource_type == res.GetType() {
		return res.GetPtr()
	}
	if resource_type_name != nil {
		var space *byte
		var class_name *byte = GetActiveClassName(&space)
		ZendError(E_WARNING, "%s%s%s(): supplied resource is not a valid %s resource", class_name, space, GetActiveFunctionName(), resource_type_name)
	}
	return nil
}
func ZendFetchResourceEx(res *Zval, resource_type_name string, resource_type int) any {
	var space *byte
	var class_name *byte
	if res == nil {
		if resource_type_name {
			class_name = GetActiveClassName(&space)
			ZendError(E_WARNING, "%s%s%s(): no %s resource supplied", class_name, space, GetActiveFunctionName(), resource_type_name)
		}
		return nil
	}
	if res.GetType() != IS_RESOURCE {
		if resource_type_name {
			class_name = GetActiveClassName(&space)
			ZendError(E_WARNING, "%s%s%s(): supplied argument is not a valid %s resource", class_name, space, GetActiveFunctionName(), resource_type_name)
		}
		return nil
	}
	return ZendFetchResource(Z_RES_P(res), resource_type_name, resource_type)
}
func ZendFetchResource2Ex(res *Zval, resource_type_name string, resource_type1 int, resource_type2 int) any {
	var space *byte
	var class_name *byte
	if res == nil {
		if resource_type_name {
			class_name = GetActiveClassName(&space)
			ZendError(E_WARNING, "%s%s%s(): no %s resource supplied", class_name, space, GetActiveFunctionName(), resource_type_name)
		}
		return nil
	}
	if res.GetType() != IS_RESOURCE {
		if resource_type_name {
			class_name = GetActiveClassName(&space)
			ZendError(E_WARNING, "%s%s%s(): supplied argument is not a valid %s resource", class_name, space, GetActiveFunctionName(), resource_type_name)
		}
		return nil
	}
	return ZendFetchResource2(Z_RES_P(res), resource_type_name, resource_type1, resource_type2)
}
func ListEntryDestructor(zv *Zval) {
	var res *ZendResource = Z_RES_P(zv)
	ZVAL_UNDEF(zv)
	if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	EfreeSize(res, b.SizeOf("zend_resource"))
}
func PlistEntryDestructor(zv *Zval) {
	var res *ZendResource = Z_RES_P(zv)
	if res.GetType() >= 0 {
		var ld *ZendRsrcListDtorsEntry
		ld = ListDestructors.IndexFindPtr(res.GetType())
		if ld != nil {
			if ld.GetPlistDtorEx() != nil {
				ld.GetPlistDtorEx()(res)
			}
		} else {
			ZendError(E_WARNING, "Unknown list entry type (%d)", res.GetType())
		}
	}
	Free(res)
}
func ZendInitRsrcList() int {
	ExecutorGlobals.GetRegularList().Init(8, nil, ListEntryDestructor, 0)
	return SUCCESS
}
func ZendInitRsrcPlist() int {
	ExecutorGlobals.GetPersistentList().InitEx(8, nil, PlistEntryDestructor, 1, 0)
	return SUCCESS
}
func ZendCloseRsrcList(ht *HashTable) {
	var res *ZendResource
	for {
		var __ht *HashTable = ht
		var _idx uint32 = __ht.GetNNumUsed()
		var _p *Bucket = __ht.GetArData() + _idx
		var _z *Zval
		for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
			_p--
			_z = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			res = Z_PTR_P(_z)
			if res.GetType() >= 0 {
				ZendResourceDtor(res)
			}
		}
		break
	}
}
func ZendDestroyRsrcList(ht *HashTable) { ht.GracefulReverseDestroy() }
func CleanModuleResource(zv *Zval, arg any) int {
	var resource_id int = *((*int)(arg))
	return Z_RES_TYPE_P(zv) == resource_id
}
func ZendCleanModuleRsrcDtorsCb(zv *Zval, arg any) int {
	var ld *ZendRsrcListDtorsEntry = (*ZendRsrcListDtorsEntry)(Z_PTR_P(zv))
	var module_number int = *((*int)(arg))
	if ld.GetModuleNumber() == module_number {
		ExecutorGlobals.GetPersistentList().ApplyWithArgument(CleanModuleResource, any(&(ld.GetResourceId())))
		return 1
	} else {
		return 0
	}
}
func ZendCleanModuleRsrcDtors(module_number int) {
	ListDestructors.ApplyWithArgument(ZendCleanModuleRsrcDtorsCb, any(&module_number))
}
func ZendRegisterListDestructorsEx(ld RsrcDtorFuncT, pld RsrcDtorFuncT, type_name string, module_number int) int {
	var lde *ZendRsrcListDtorsEntry
	var zv Zval
	lde = Malloc(b.SizeOf("zend_rsrc_list_dtors_entry"))
	lde.SetListDtorEx(ld)
	lde.SetPlistDtorEx(pld)
	lde.SetModuleNumber(module_number)
	lde.SetResourceId(ListDestructors.GetNNextFreeElement())
	lde.SetTypeName(type_name)
	ZVAL_PTR(&zv, lde)
	if ListDestructors.NextIndexInsert(&zv) == nil {
		return FAILURE
	}
	return ListDestructors.GetNNextFreeElement() - 1
}
func ZendFetchListDtorId(type_name *byte) int {
	var lde *ZendRsrcListDtorsEntry
	for {
		var __ht *HashTable = &ListDestructors
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			lde = Z_PTR_P(_z)
			if lde.GetTypeName() != nil && strcmp(type_name, lde.GetTypeName()) == 0 {
				return lde.GetResourceId()
			}
		}
		break
	}
	return 0
}
func ListDestructorsDtor(zv *Zval) { Free(Z_PTR_P(zv)) }
func ZendInitRsrcListDtors() int {
	ListDestructors.Init(64, nil, ListDestructorsDtor, 1)
	ListDestructors.SetNNextFreeElement(1)
	return SUCCESS
}
func ZendDestroyRsrcListDtors() { ListDestructors.Destroy() }
func ZendRsrcListGetRsrcType(res *ZendResource) *byte {
	var lde *ZendRsrcListDtorsEntry
	lde = ListDestructors.IndexFindPtr(res.GetType())
	if lde != nil {
		return lde.GetTypeName()
	} else {
		return nil
	}
}
func ZendRegisterPersistentResourceEx(key *ZendString, rsrc_pointer any, rsrc_type int) *ZendResource {
	var zv *Zval
	var tmp Zval
	ZVAL_NEW_PERSISTENT_RES(&tmp, -1, rsrc_pointer, rsrc_type)
	GC_MAKE_PERSISTENT_LOCAL(tmp.GetCounted())
	GC_MAKE_PERSISTENT_LOCAL(key)
	zv = ExecutorGlobals.GetPersistentList().Update(key, &tmp)
	return Z_RES_P(zv)
}
func ZendRegisterPersistentResource(key *byte, key_len int, rsrc_pointer any, rsrc_type int) *ZendResource {
	var str *ZendString = ZendStringInit(key, key_len, 1)
	var ret *ZendResource = ZendRegisterPersistentResourceEx(str, rsrc_pointer, rsrc_type)
	ZendStringReleaseEx(str, 1)
	return ret
}
