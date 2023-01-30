// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func ZendListInsert(ptr any, type_ int) *Zval {
	var index int
	var zv Zval
	index = __EG().GetRegularList().GetNNextFreeElement()
	if index == 0 {
		index = 1
	} else if index == core.INT_MAX {
		ZendErrorNoreturn(E_ERROR, "Resource ID space overflow")
	}
	ZVAL_NEW_RES(&zv, index, ptr, type_)
	return ZendHashIndexAddNew(__EG().GetRegularList(), index, &zv)
}
func ZendListDelete(res *ZendResource) int {
	if res.DelRefcount() <= 0 {
		return ZendHashIndexDel(__EG().GetRegularList(), res.GetHandle())
	} else {
		return SUCCESS
	}
}
func ZendListFree(res *ZendResource) int {
	if res.GetRefcount() <= 0 {
		return ZendHashIndexDel(__EG().GetRegularList(), res.GetHandle())
	} else {
		return SUCCESS
	}
}
func ZendResourceDtor(res *ZendResource) {
	var ld *ZendRsrcListDtorsEntry
	var r ZendResource = *res
	res.SetType(-1)
	res.SetPtr(nil)
	ld = ZendHashIndexFindPtr(&ListDestructors, r.GetType())
	if ld != nil {
		if ld.GetListDtorEx() != nil {
			ld.GetListDtorEx()(&r)
		}
	} else {
		ZendError(E_WARNING, "Unknown list entry type (%d)", r.GetType())
	}
}
func ZendListClose(res *ZendResource) int {
	if res.GetRefcount() <= 0 {
		return ZendListFree(res)
	} else if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	return SUCCESS
}
func ZendRegisterResource(rsrc_pointer any, rsrc_type int) *ZendResource {
	var zv *Zval
	zv = ZendListInsert(rsrc_pointer, rsrc_type)
	return zv.GetRes()
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
	return ZendFetchResource(res.GetRes(), resource_type_name, resource_type)
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
	return ZendFetchResource2(res.GetRes(), resource_type_name, resource_type1, resource_type2)
}
func ListEntryDestructor(zv *Zval) {
	var res *ZendResource = zv.GetRes()
	ZVAL_UNDEF(zv)
	if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	EfreeSize(res, b.SizeOf("zend_resource"))
}
func PlistEntryDestructor(zv *Zval) {
	var res *ZendResource = zv.GetRes()
	if res.GetType() >= 0 {
		var ld *ZendRsrcListDtorsEntry
		ld = ZendHashIndexFindPtr(&ListDestructors, res.GetType())
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
	ZendHashInit(__EG().GetRegularList(), 8, nil, ListEntryDestructor, 0)
	return SUCCESS
}
func ZendInitRsrcPlist() int {
	ZendHashInitEx(__EG().GetPersistentList(), 8, nil, PlistEntryDestructor, 1, 0)
	return SUCCESS
}
func ZendCloseRsrcList(ht *HashTable) {
	var res *ZendResource
	var __ht *HashTable = ht
	for _, _p := range __ht.foreachDataReserve() {
		var _z Zval = _p.GetVal()

		res = _z.GetPtr()
		if res.GetType() >= 0 {
			ZendResourceDtor(res)
		}
	}
}
func ZendDestroyRsrcList(ht *HashTable) { ZendHashGracefulReverseDestroy(ht) }
func CleanModuleResource(zv *Zval, arg any) int {
	var resource_id int = *((*int)(arg))
	return Z_RES_TYPE_P(zv) == resource_id
}
func ZendCleanModuleRsrcDtorsCb(zv *Zval, arg any) int {
	var ld *ZendRsrcListDtorsEntry = (*ZendRsrcListDtorsEntry)(zv.GetPtr())
	var module_number int = *((*int)(arg))
	if ld.GetModuleNumber() == module_number {
		ZendHashApplyWithArgument(__EG().GetPersistentList(), CleanModuleResource, any(&(ld.GetResourceId())))
		return 1
	} else {
		return 0
	}
}
func ZendCleanModuleRsrcDtors(module_number int) {
	ZendHashApplyWithArgument(&ListDestructors, ZendCleanModuleRsrcDtorsCb, any(&module_number))
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
	if ZendHashNextIndexInsert(&ListDestructors, &zv) == nil {
		return FAILURE
	}
	return ListDestructors.GetNNextFreeElement() - 1
}
func ZendFetchListDtorId(type_name *byte) int {
	var lde *ZendRsrcListDtorsEntry
	var __ht *HashTable = &ListDestructors
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		lde = _z.GetPtr()
		if lde.GetTypeName() != nil && strcmp(type_name, lde.GetTypeName()) == 0 {
			return lde.GetResourceId()
		}
	}
	return 0
}
func ListDestructorsDtor(zv *Zval) { Free(zv.GetPtr()) }
func ZendInitRsrcListDtors() int {
	ZendHashInit(&ListDestructors, 64, nil, ListDestructorsDtor, 1)
	ListDestructors.SetNNextFreeElement(1)
	return SUCCESS
}
func ZendDestroyRsrcListDtors() { ZendHashDestroy(&ListDestructors) }
func ZendRsrcListGetRsrcType(res *ZendResource) *byte {
	var lde *ZendRsrcListDtorsEntry
	lde = ZendHashIndexFindPtr(&ListDestructors, res.GetType())
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
	zv = ZendHashUpdate(__EG().GetPersistentList(), key, &tmp)
	return zv.GetRes()
}
func ZendRegisterPersistentResource(key *byte, key_len int, rsrc_pointer any, rsrc_type int) *ZendResource {
	var str *ZendString = ZendStringInit(key, key_len, 1)
	var ret *ZendResource = ZendRegisterPersistentResourceEx(str, rsrc_pointer, rsrc_type)
	ZendStringReleaseEx(str, 1)
	return ret
}
