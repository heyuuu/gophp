package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZendListInsert(ptr any, type_ int) *types2.Zval {
	var index int
	var zv types2.Zval
	index = EG__().GetRegularList().GetNNextFreeElement()
	if index == 0 {
		index = 1
	} else if index == core.INT_MAX {
		faults.ErrorNoreturn(faults.E_ERROR, "Resource ID space overflow")
	}
	zv.SetResource(types2.NewZendResource(index, ptr, type_))
	return EG__().GetRegularList().IndexAddNew(index, &zv)
}
func ZendListDelete(res *types2.ZendResource) int {
	if res.DelRefcount() <= 0 {
		return types2.ZendHashIndexDel(EG__().GetRegularList(), res.GetHandle())
	} else {
		return types2.SUCCESS
	}
}
func ZendListFree(res *types2.ZendResource) int {
	if res.GetRefcount() <= 0 {
		return types2.ZendHashIndexDel(EG__().GetRegularList(), res.GetHandle())
	} else {
		return types2.SUCCESS
	}
}
func ZendResourceDtor(res *types2.ZendResource) {
	var ld *ZendRsrcListDtorsEntry
	var r types2.ZendResource = *res
	res.SetType(-1)
	res.SetPtr(nil)
	ld = types2.ZendHashIndexFindPtr(&ListDestructors, r.GetType())
	if ld != nil {
		if ld.GetListDtorEx() != nil {
			ld.GetListDtorEx()(&r)
		}
	} else {
		faults.Error(faults.E_WARNING, "Unknown list entry type (%d)", r.GetType())
	}
}
func ZendListClose(res *types2.ZendResource) int {
	if res.GetRefcount() <= 0 {
		return ZendListFree(res)
	} else if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	return types2.SUCCESS
}
func ZendRegisterResource(rsrc_pointer any, rsrc_type int) *types2.ZendResource {
	var zv *types2.Zval
	zv = ZendListInsert(rsrc_pointer, rsrc_type)
	return zv.Resource()
}
func ZendFetchResource2(res *types2.ZendResource, resource_type_name string, resource_type1 int, resource_type2 int) any {
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
func ZendFetchResource(res *types2.ZendResource, resource_type_name *byte, resource_type int) any {
	if resource_type == res.GetType() {
		return res.GetPtr()
	}
	if resource_type_name != nil {
		faults.Error(faults.E_WARNING, "%s(): supplied resource is not a valid %s resource", GetActiveCalleeName(), resource_type_name)
	}
	return nil
}
func ZendFetchResourceEx(res *types2.Zval, resource_type_name string, resource_type int) any {
	if res == nil {
		if resource_type_name {
			faults.Error(faults.E_WARNING, "%s(): no %s resource supplied", GetActiveCalleeName(), resource_type_name)
		}
		return nil
	}
	if res.GetType() != types2.IS_RESOURCE {
		if resource_type_name {
			faults.Error(faults.E_WARNING, "%s(): supplied argument is not a valid %s resource", GetActiveCalleeName(), resource_type_name)
		}
		return nil
	}
	return ZendFetchResource(res.Resource(), resource_type_name, resource_type)
}
func ZendFetchResource2Ex(res *types2.Zval, resource_type_name string, resource_type1 int, resource_type2 int) any {
	if res == nil {
		if resource_type_name {
			faults.Error(faults.E_WARNING, "%s(): no %s resource supplied", GetActiveCalleeName(), resource_type_name)
		}
		return nil
	}
	if res.GetType() != types2.IS_RESOURCE {
		if resource_type_name {
			faults.Error(faults.E_WARNING, "%s(): supplied argument is not a valid %s resource", GetActiveCalleeName(), resource_type_name)
		}
		return nil
	}
	return ZendFetchResource2(res.Resource(), resource_type_name, resource_type1, resource_type2)
}
func ListEntryDestructor(zv *types2.Zval) {
	var res *types2.ZendResource = zv.Resource()
	zv.SetUndef()
	if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	EfreeSize(res, b.SizeOf("zend_resource"))
}
func PlistEntryDestructor(zv *types2.Zval) {
	var res *types2.ZendResource = zv.Resource()
	if res.GetType() >= 0 {
		var ld *ZendRsrcListDtorsEntry
		ld = types2.ZendHashIndexFindPtr(&ListDestructors, res.GetType())
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
	EG__().GetRegularList().Init(8, ListEntryDestructor)
	return types2.SUCCESS
}
func ZendInitRsrcPlist() int {
	EG__().GetPersistentList().Init(8, PlistEntryDestructor)
	return types2.SUCCESS
}
func ZendCloseRsrcList(ht *types2.Array) {
	var res *types2.ZendResource
	var __ht *types2.Array = ht
	for _, _p := range __ht.ForeachDataReserve() {
		var _z types2.Zval = _p.GetVal()

		res = _z.Ptr()
		if res.GetType() >= 0 {
			ZendResourceDtor(res)
		}
	}
}
func CleanModuleResource(zv *types2.Zval, arg any) int {
	var resource_id int = *((*int)(arg))
	return types2.Z_RES_TYPE_P(zv) == resource_id
}
func ZendCleanModuleRsrcDtorsCb(zv *types2.Zval, arg any) int {
	var ld *ZendRsrcListDtorsEntry = (*ZendRsrcListDtorsEntry)(zv.Ptr())
	var module_number int = *((*int)(arg))
	if ld.GetModuleNumber() == module_number {
		types2.ZendHashApplyWithArgument(EG__().GetPersistentList(), CleanModuleResource, any(&(ld.GetResourceId())))
		return 1
	} else {
		return 0
	}
}
func ZendCleanModuleRsrcDtors(module_number int) {
	types2.ZendHashApplyWithArgument(&ListDestructors, ZendCleanModuleRsrcDtorsCb, any(&module_number))
}
func ZendRegisterListDestructorsEx(ld RsrcDtorFuncT, pld RsrcDtorFuncT, type_name string, module_number int) int {
	var lde *ZendRsrcListDtorsEntry
	var zv types2.Zval
	lde = Malloc(b.SizeOf("zend_rsrc_list_dtors_entry"))
	lde.SetListDtorEx(ld)
	lde.SetPlistDtorEx(pld)
	lde.SetModuleNumber(module_number)
	lde.SetResourceId(ListDestructors.GetNNextFreeElement())
	lde.SetTypeName(type_name)
	zv.SetPtr(lde)
	if ListDestructors.NextIndexInsert(&zv) == nil {
		return types2.FAILURE
	}
	return ListDestructors.GetNNextFreeElement() - 1
}
func ZendFetchListDtorId(type_name *byte) int {
	var lde *ZendRsrcListDtorsEntry
	var __ht *types2.Array = &ListDestructors
	for _, _p := range __ht.ForeachData() {
		var _z *types2.Zval = _p.GetVal()

		lde = _z.Ptr()
		if lde.GetTypeName() != nil && strcmp(type_name, lde.GetTypeName()) == 0 {
			return lde.GetResourceId()
		}
	}
	return 0
}
func ListDestructorsDtor(zv *types2.Zval) { Free(zv.Ptr()) }
func ZendInitRsrcListDtors() int {
	ListDestructors.Init(64, ListDestructorsDtor)
	ListDestructors.SetNNextFreeElement(1)
	return types2.SUCCESS
}
func ZendDestroyRsrcListDtors() { ListDestructors.Destroy() }
func ZendRsrcListGetRsrcType(res *types2.ZendResource) *byte {
	var lde *ZendRsrcListDtorsEntry
	lde = types2.ZendHashIndexFindPtr(&ListDestructors, res.GetType())
	if lde != nil {
		return lde.GetTypeName()
	} else {
		return nil
	}
}
func ZendRsrcListGetRsrcTypeEx(res *types2.ZendResource) *string {
	var lde *ZendRsrcListDtorsEntry
	lde = types2.ZendHashIndexFindPtr(&ListDestructors, res.GetType())
	if lde == nil {
		return nil
	}
	var typeName = lde.TypeName()
	return &typeName
}
func ZendRegisterPersistentResourceEx(key *types2.String, rsrc_pointer any, rsrc_type int) *types2.ZendResource {
	var zv *types2.Zval
	var tmp types2.Zval
	tmp.SetResource(types2.NewZendResourcePersistent(-1, rsrc_pointer, rsrc_type, true))
	//types.GC_MAKE_PERSISTENT_LOCAL(tmp.RefCounted())
	//types.GC_MAKE_PERSISTENT_LOCAL(key)
	zv = EG__().GetPersistentList().KeyUpdate(key.GetStr(), &tmp)
	return zv.Resource()
}
func ZendRegisterPersistentResource(key *byte, key_len int, rsrc_pointer any, rsrc_type int) *types2.ZendResource {
	var str *types2.String = types2.NewString(b.CastStr(key, key_len))
	var ret *types2.ZendResource = ZendRegisterPersistentResourceEx(str, rsrc_pointer, rsrc_type)
	// types.ZendStringReleaseEx(str, 1)
	return ret
}
