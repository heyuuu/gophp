package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZendListInsert(ptr any, type_ int) *types.Zval {
	var index int
	var zv types.Zval
	index = EG__().GetRegularList().GetNNextFreeElement()
	if index == 0 {
		index = 1
	} else if index == core.INT_MAX {
		faults.ErrorNoreturn(faults.E_ERROR, "Resource ID space overflow")
	}
	zv.SetNewResource(index, ptr, type_)
	return EG__().GetRegularList().IndexAddNew(index, &zv)
}
func ZendListDelete(res *types.ZendResource) int {
	if res.DelRefcount() <= 0 {
		return types.ZendHashIndexDel(EG__().GetRegularList(), res.GetHandle())
	} else {
		return types.SUCCESS
	}
}
func ZendListFree(res *types.ZendResource) int {
	if res.GetRefcount() <= 0 {
		return types.ZendHashIndexDel(EG__().GetRegularList(), res.GetHandle())
	} else {
		return types.SUCCESS
	}
}
func ZendResourceDtor(res *types.ZendResource) {
	var ld *ZendRsrcListDtorsEntry
	var r types.ZendResource = *res
	res.SetType(-1)
	res.SetPtr(nil)
	ld = types.ZendHashIndexFindPtr(&ListDestructors, r.GetType())
	if ld != nil {
		if ld.GetListDtorEx() != nil {
			ld.GetListDtorEx()(&r)
		}
	} else {
		faults.Error(faults.E_WARNING, "Unknown list entry type (%d)", r.GetType())
	}
}
func ZendListClose(res *types.ZendResource) int {
	if res.GetRefcount() <= 0 {
		return ZendListFree(res)
	} else if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	return types.SUCCESS
}
func ZendRegisterResource(rsrc_pointer any, rsrc_type int) *types.ZendResource {
	var zv *types.Zval
	zv = ZendListInsert(rsrc_pointer, rsrc_type)
	return zv.GetRes()
}
func ZendFetchResource2(res *types.ZendResource, resource_type_name string, resource_type1 int, resource_type2 int) any {
	if res != nil {
		if resource_type1 == res.GetType() {
			return res.GetPtr()
		}
		if resource_type2 == res.GetType() {
			return res.GetPtr()
		}
	}
	if resource_type_name {
		faults.Error(faults.E_WARNING, "%s(): supplied resource is not a valid %s resource", GetActiveCalleeName(), resource_type_name)
	}
	return nil
}
func ZendFetchResource(res *types.ZendResource, resource_type_name *byte, resource_type int) any {
	if resource_type == res.GetType() {
		return res.GetPtr()
	}
	if resource_type_name != nil {
		faults.Error(faults.E_WARNING, "%s(): supplied resource is not a valid %s resource", GetActiveCalleeName(), resource_type_name)
	}
	return nil
}
func ZendFetchResourceEx(res *types.Zval, resource_type_name string, resource_type int) any {
	if res == nil {
		if resource_type_name {
			faults.Error(faults.E_WARNING, "%s(): no %s resource supplied", GetActiveCalleeName(), resource_type_name)
		}
		return nil
	}
	if res.GetType() != types.IS_RESOURCE {
		if resource_type_name {
			faults.Error(faults.E_WARNING, "%s(): supplied argument is not a valid %s resource", GetActiveCalleeName(), resource_type_name)
		}
		return nil
	}
	return ZendFetchResource(res.GetRes(), resource_type_name, resource_type)
}
func ZendFetchResource2Ex(res *types.Zval, resource_type_name string, resource_type1 int, resource_type2 int) any {
	if res == nil {
		if resource_type_name {
			faults.Error(faults.E_WARNING, "%s(): no %s resource supplied", GetActiveCalleeName(), resource_type_name)
		}
		return nil
	}
	if res.GetType() != types.IS_RESOURCE {
		if resource_type_name {
			faults.Error(faults.E_WARNING, "%s(): supplied argument is not a valid %s resource", GetActiveCalleeName(), resource_type_name)
		}
		return nil
	}
	return ZendFetchResource2(res.GetRes(), resource_type_name, resource_type1, resource_type2)
}
func ListEntryDestructor(zv *types.Zval) {
	var res *types.ZendResource = zv.GetRes()
	zv.SetUndef()
	if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	EfreeSize(res, b.SizeOf("zend_resource"))
}
func PlistEntryDestructor(zv *types.Zval) {
	var res *types.ZendResource = zv.GetRes()
	if res.GetType() >= 0 {
		var ld *ZendRsrcListDtorsEntry
		ld = types.ZendHashIndexFindPtr(&ListDestructors, res.GetType())
		if ld != nil {
			if ld.GetPlistDtorEx() != nil {
				ld.GetPlistDtorEx()(res)
			}
		} else {
			faults.Error(faults.E_WARNING, "Unknown list entry type (%d)", res.GetType())
		}
	}
	Free(res)
}
func ZendInitRsrcList() int {
	EG__().GetRegularList() = types.MakeArrayEx(8, ListEntryDestructor, 0)
	return types.SUCCESS
}
func ZendInitRsrcPlist() int {
	EG__().GetPersistentList() = types.MakeArrayEx(8, PlistEntryDestructor, 1)
	return types.SUCCESS
}
func ZendCloseRsrcList(ht *types.Array) {
	var res *types.ZendResource
	var __ht *types.Array = ht
	for _, _p := range __ht.ForeachDataReserve() {
		var _z types.Zval = _p.GetVal()

		res = _z.GetPtr()
		if res.GetType() >= 0 {
			ZendResourceDtor(res)
		}
	}
}
func CleanModuleResource(zv *types.Zval, arg any) int {
	var resource_id int = *((*int)(arg))
	return types.Z_RES_TYPE_P(zv) == resource_id
}
func ZendCleanModuleRsrcDtorsCb(zv *types.Zval, arg any) int {
	var ld *ZendRsrcListDtorsEntry = (*ZendRsrcListDtorsEntry)(zv.GetPtr())
	var module_number int = *((*int)(arg))
	if ld.GetModuleNumber() == module_number {
		types.ZendHashApplyWithArgument(EG__().GetPersistentList(), CleanModuleResource, any(&(ld.GetResourceId())))
		return 1
	} else {
		return 0
	}
}
func ZendCleanModuleRsrcDtors(module_number int) {
	types.ZendHashApplyWithArgument(&ListDestructors, ZendCleanModuleRsrcDtorsCb, any(&module_number))
}
func ZendRegisterListDestructorsEx(ld RsrcDtorFuncT, pld RsrcDtorFuncT, type_name string, module_number int) int {
	var lde *ZendRsrcListDtorsEntry
	var zv types.Zval
	lde = Malloc(b.SizeOf("zend_rsrc_list_dtors_entry"))
	lde.SetListDtorEx(ld)
	lde.SetPlistDtorEx(pld)
	lde.SetModuleNumber(module_number)
	lde.SetResourceId(ListDestructors.GetNNextFreeElement())
	lde.SetTypeName(type_name)
	zv.SetAsPtr(lde)
	if ListDestructors.NextIndexInsert(&zv) == nil {
		return types.FAILURE
	}
	return ListDestructors.GetNNextFreeElement() - 1
}
func ZendFetchListDtorId(type_name *byte) int {
	var lde *ZendRsrcListDtorsEntry
	var __ht *types.Array = &ListDestructors
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		lde = _z.GetPtr()
		if lde.GetTypeName() != nil && strcmp(type_name, lde.GetTypeName()) == 0 {
			return lde.GetResourceId()
		}
	}
	return 0
}
func ListDestructorsDtor(zv *types.Zval) { Free(zv.GetPtr()) }
func ZendInitRsrcListDtors() int {
	&ListDestructors = types.MakeArrayEx(64, ListDestructorsDtor, 1)
	ListDestructors.SetNNextFreeElement(1)
	return types.SUCCESS
}
func ZendDestroyRsrcListDtors() { ListDestructors.Destroy() }
func ZendRsrcListGetRsrcType(res *types.ZendResource) *byte {
	var lde *ZendRsrcListDtorsEntry
	lde = types.ZendHashIndexFindPtr(&ListDestructors, res.GetType())
	if lde != nil {
		return lde.GetTypeName()
	} else {
		return nil
	}
}
func ZendRegisterPersistentResourceEx(key *types.String, rsrc_pointer any, rsrc_type int) *types.ZendResource {
	var zv *types.Zval
	var tmp types.Zval
	tmp.SetNewResourcePersistent(-1, rsrc_pointer, rsrc_type)
	types.GC_MAKE_PERSISTENT_LOCAL(tmp.GetCounted())
	types.GC_MAKE_PERSISTENT_LOCAL(key)
	zv = EG__().GetPersistentList().KeyUpdate(key.GetStr(), &tmp)
	return zv.GetRes()
}
func ZendRegisterPersistentResource(key *byte, key_len int, rsrc_pointer any, rsrc_type int) *types.ZendResource {
	var str *types.String = types.NewString(b.CastStr(key, key_len))
	var ret *types.ZendResource = ZendRegisterPersistentResourceEx(str, rsrc_pointer, rsrc_type)
	types.ZendStringReleaseEx(str, 1)
	return ret
}
