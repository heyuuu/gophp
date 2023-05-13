package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
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
	zv.SetResource(types.NewZendResource(index, ptr, type_))
	return EG__().GetRegularList().IndexAddNew(index, &zv)
}
func ZendListDelete(res *types.ZendResource) int {
	// todo 移除机制待处理
	//if res.DelRefcount() <= 0 {
	//	return types.ZendHashIndexDel(EG__().GetRegularList(), res.GetHandle())
	//} else {
	//	return types.SUCCESS
	//}
	return types.SUCCESS
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
	var r = *res
	res.SetType(-1)
	res.SetPtr(nil)
	ld = ListDestructors.Find(&r)
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
	return zv.Resource()
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
	return ZendFetchResource(res.Resource(), resource_type_name, resource_type)
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
	return ZendFetchResource2(res.Resource(), resource_type_name, resource_type1, resource_type2)
}
func ListEntryDtor(res *types.ZendResource) {
	if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	EfreeSize(res, b.SizeOf("zend_resource"))
}
func PlistEntryDtor(res *types.ZendResource) {
	if res.GetType() >= 0 {
		var ld = ListDestructors.Find(res)
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
	EG__().InitRegularList()
	return types.SUCCESS
}
func ZendInitRsrcPlist() int {
	EG__().InitPersistentList()
	return types.SUCCESS
}
func ZendCleanModuleRsrcDtors(moduleNumber int) {
	ListDestructors.CleanByModule(moduleNumber)
}
func ZendRegisterListDestructorsEx(ld RsrcDtorFuncT, pld RsrcDtorFuncT, typeName string, moduleNumber int) int {
	var lde = NewZendRsrcListDtorsEntry(ld, pld, typeName, moduleNumber)
	return ListDestructors.Append(lde)
}
func ZendFetchListDtorId(typeName string) int {
	return ListDestructors.GetResourceIdByTypeName(typeName)
}
func ZendInitRsrcListDtors() int {
	ListDestructors.Init()
	return types.SUCCESS
}
func ZendDestroyRsrcListDtors() { ListDestructors.Destroy() }
func ZendRsrcListGetRsrcType(res *types.ZendResource) *byte {
	var lde = ListDestructors.Find(res)
	if lde != nil {
		return lde.GetTypeName()
	} else {
		return nil
	}
}
func ZendRsrcListGetRsrcTypeEx(res *types.ZendResource) *string {
	var lde = ListDestructors.Find(res)
	if lde == nil {
		return nil
	}
	var typeName = lde.TypeName()
	return &typeName
}
