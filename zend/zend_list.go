// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_list.h>

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

// #define ZEND_LIST_H

// # include "zend_hash.h"

// # include "zend_globals.h"

type RsrcDtorFuncT func(res *ZendResource)

// #define ZEND_RSRC_DTOR_FUNC(name) void name ( zend_resource * res )

var LeIndexPtr int

// Source: <Zend/zend_list.c>

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

// # include "zend.h"

// # include "zend_list.h"

// # include "zend_API.h"

// # include "zend_globals.h"

/* true global */

var ListDestructors HashTable

func ZendListInsert(ptr any, type_ int) *Zval {
	var index int
	var zv Zval
	index = &EG.regular_list.nNextFreeElement
	if index == 0 {
		index = 1
	} else if index == INT_MAX {
		ZendErrorNoreturn(1<<0, "Resource ID space overflow")
	}
	var _res *ZendResource = (*ZendResource)(_emalloc(g.SizeOf("zend_resource")))
	var __z *Zval
	ZendGcSetRefcount(&_res.gc, 1)
	_res.GetGc().SetTypeInfo(9)
	_res.SetHandle(index)
	_res.SetType(type_)
	_res.SetPtr(ptr)
	__z = &zv
	__z.GetValue().SetRes(_res)
	__z.SetTypeInfo(9 | 1<<0<<8)
	return ZendHashIndexAddNew(&EG.regular_list, index, &zv)
}
func ZendListDelete(res *ZendResource) int {
	if ZendGcDelref(&res.gc) <= 0 {
		return ZendHashIndexDel(&EG.regular_list, res.GetHandle())
	} else {
		return SUCCESS
	}
}
func ZendListFree(res *ZendResource) int {
	if ZendGcRefcount(&res.gc) <= 0 {
		return ZendHashIndexDel(&EG.regular_list, res.GetHandle())
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
		ZendError(1<<1, "Unknown list entry type (%d)", r.GetType())
	}
}
func ZendListClose(res *ZendResource) int {
	if ZendGcRefcount(&res.gc) <= 0 {
		return ZendListFree(res)
	} else if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	return SUCCESS
}
func ZendRegisterResource(rsrc_pointer any, rsrc_type int) *ZendResource {
	var zv *Zval
	zv = ZendListInsert(rsrc_pointer, rsrc_type)
	return zv.GetValue().GetRes()
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
		ZendError(1<<1, "%s%s%s(): supplied resource is not a valid %s resource", class_name, space, GetActiveFunctionName(), resource_type_name)
	}
	return nil
}
func ZendFetchResource(res *ZendResource, resource_type_name string, resource_type int) any {
	if resource_type == res.GetType() {
		return res.GetPtr()
	}
	if resource_type_name {
		var space *byte
		var class_name *byte = GetActiveClassName(&space)
		ZendError(1<<1, "%s%s%s(): supplied resource is not a valid %s resource", class_name, space, GetActiveFunctionName(), resource_type_name)
	}
	return nil
}
func ZendFetchResourceEx(res *Zval, resource_type_name string, resource_type int) any {
	var space *byte
	var class_name *byte
	if res == nil {
		if resource_type_name {
			class_name = GetActiveClassName(&space)
			ZendError(1<<1, "%s%s%s(): no %s resource supplied", class_name, space, GetActiveFunctionName(), resource_type_name)
		}
		return nil
	}
	if res.GetType() != 9 {
		if resource_type_name {
			class_name = GetActiveClassName(&space)
			ZendError(1<<1, "%s%s%s(): supplied argument is not a valid %s resource", class_name, space, GetActiveFunctionName(), resource_type_name)
		}
		return nil
	}
	return ZendFetchResource(res.GetValue().GetRes(), resource_type_name, resource_type)
}
func ZendFetchResource2Ex(res *Zval, resource_type_name string, resource_type1 int, resource_type2 int) any {
	var space *byte
	var class_name *byte
	if res == nil {
		if resource_type_name {
			class_name = GetActiveClassName(&space)
			ZendError(1<<1, "%s%s%s(): no %s resource supplied", class_name, space, GetActiveFunctionName(), resource_type_name)
		}
		return nil
	}
	if res.GetType() != 9 {
		if resource_type_name {
			class_name = GetActiveClassName(&space)
			ZendError(1<<1, "%s%s%s(): supplied argument is not a valid %s resource", class_name, space, GetActiveFunctionName(), resource_type_name)
		}
		return nil
	}
	return ZendFetchResource2(res.GetValue().GetRes(), resource_type_name, resource_type1, resource_type2)
}
func ListEntryDestructor(zv *Zval) {
	var res *ZendResource = zv.GetValue().GetRes()
	zv.SetTypeInfo(0)
	if res.GetType() >= 0 {
		ZendResourceDtor(res)
	}
	_efree(res)
}
func PlistEntryDestructor(zv *Zval) {
	var res *ZendResource = zv.GetValue().GetRes()
	if res.GetType() >= 0 {
		var ld *ZendRsrcListDtorsEntry
		ld = ZendHashIndexFindPtr(&ListDestructors, res.GetType())
		if ld != nil {
			if ld.GetPlistDtorEx() != nil {
				ld.GetPlistDtorEx()(res)
			}
		} else {
			ZendError(1<<1, "Unknown list entry type (%d)", res.GetType())
		}
	}
	Free(res)
}
func ZendInitRsrcList() int {
	_zendHashInit(&EG.regular_list, 8, ListEntryDestructor, 0)
	return SUCCESS
}
func ZendInitRsrcPlist() int {
	_zendHashInit(&EG.persistent_list, 8, PlistEntryDestructor, 1)
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
			_z = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			res = _z.GetValue().GetPtr()
			if res.GetType() >= 0 {
				ZendResourceDtor(res)
			}
		}
		break
	}
}
func ZendDestroyRsrcList(ht *HashTable) { ZendHashGracefulReverseDestroy(ht) }
func CleanModuleResource(zv *Zval, arg any) int {
	var resource_id int = *((*int)(arg))
	return zv.GetValue().GetRes().GetType() == resource_id
}
func ZendCleanModuleRsrcDtorsCb(zv *Zval, arg any) int {
	var ld *ZendRsrcListDtorsEntry = (*ZendRsrcListDtorsEntry)(zv.GetValue().GetPtr())
	var module_number int = *((*int)(arg))
	if ld.GetModuleNumber() == module_number {
		ZendHashApplyWithArgument(&EG.persistent_list, CleanModuleResource, any(&(ld.GetResourceId())))
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
	lde = Malloc(g.SizeOf("zend_rsrc_list_dtors_entry"))
	lde.SetListDtorEx(ld)
	lde.SetPlistDtorEx(pld)
	lde.SetModuleNumber(module_number)
	lde.SetResourceId(ListDestructors.GetNNextFreeElement())
	lde.SetTypeName(type_name)
	&zv.GetValue().SetPtr(lde)
	&zv.SetTypeInfo(14)
	if ZendHashNextIndexInsert(&ListDestructors, &zv) == nil {
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
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			lde = _z.GetValue().GetPtr()
			if lde.GetTypeName() != nil && strcmp(type_name, lde.GetTypeName()) == 0 {
				return lde.GetResourceId()
			}
		}
		break
	}
	return 0
}
func ListDestructorsDtor(zv *Zval) { Free(zv.GetValue().GetPtr()) }
func ZendInitRsrcListDtors() int {
	_zendHashInit(&ListDestructors, 64, ListDestructorsDtor, 1)
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
	var _res *ZendResource = (*ZendResource)(Malloc(g.SizeOf("zend_resource")))
	var __z *Zval
	ZendGcSetRefcount(&_res.gc, 1)
	_res.GetGc().SetTypeInfo(9 | 1<<7<<0)
	_res.SetHandle(-1)
	_res.SetType(rsrc_type)
	_res.SetPtr(rsrc_pointer)
	__z = &tmp
	__z.GetValue().SetRes(_res)
	__z.SetTypeInfo(9 | 1<<0<<8)

	zv = ZendHashUpdate(&EG.persistent_list, key, &tmp)
	return zv.GetValue().GetRes()
}
func ZendRegisterPersistentResource(key *byte, key_len int, rsrc_pointer any, rsrc_type int) *ZendResource {
	var str *ZendString = ZendStringInit(key, key_len, 1)
	var ret *ZendResource = ZendRegisterPersistentResourceEx(str, rsrc_pointer, rsrc_type)
	ZendStringReleaseEx(str, 1)
	return ret
}
