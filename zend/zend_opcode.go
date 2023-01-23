// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_opcode.c>

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

// # include < stdio . h >

// # include "zend.h"

// # include "zend_alloc.h"

// # include "zend_compile.h"

// # include "zend_extensions.h"

// # include "zend_API.h"

// # include "zend_sort.h"

// # include "zend_vm.h"

func ZendExtensionOpArrayCtorHandler(extension *ZendExtension, op_array *ZendOpArray) {
	if extension.GetOpArrayCtor() != nil {
		extension.GetOpArrayCtor()(op_array)
	}
}
func ZendExtensionOpArrayDtorHandler(extension *ZendExtension, op_array *ZendOpArray) {
	if extension.GetOpArrayDtor() != nil {
		extension.GetOpArrayDtor()(op_array)
	}
}
func InitOpArray(op_array *ZendOpArray, type_ ZendUchar, initial_ops_size int) {
	op_array.SetType(type_)
	op_array.GetArgFlags()[0] = 0
	op_array.GetArgFlags()[1] = 0
	op_array.GetArgFlags()[2] = 0
	op_array.SetRefcount((*uint32)(_emalloc(g.SizeOf("uint32_t"))))
	(*op_array).refcount = 1
	op_array.SetLast(0)
	op_array.SetOpcodes(_emalloc(initial_ops_size * g.SizeOf("zend_op")))
	op_array.SetLastVar(0)
	op_array.SetVars(nil)
	op_array.SetT(0)
	op_array.SetFunctionName(nil)
	op_array.SetFilename(ZendGetCompiledFilename())
	op_array.SetDocComment(nil)
	op_array.SetArgInfo(nil)
	op_array.SetNumArgs(0)
	op_array.SetRequiredNumArgs(0)
	op_array.SetScope(nil)
	op_array.SetPrototype(nil)
	op_array.SetLiveRange(nil)
	op_array.SetTryCatchArray(nil)
	op_array.SetLastLiveRange(0)
	op_array.SetStaticVariables(nil)
	op_array.SetStaticVariablesPtrPtr(&op_array.static_variables)
	op_array.SetLastTryCatch(0)
	op_array.SetFnFlags(0)
	op_array.SetLastLiteral(0)
	op_array.SetLiterals(nil)
	op_array.SetRunTimeCachePtr(nil)
	op_array.SetCacheSize(ZendOpArrayExtensionHandles * g.SizeOf("void *"))
	memset(op_array.GetReserved(), 0, 6*g.SizeOf("void *"))
	if (ZendExtensionFlags & 1 << 0) != 0 {
		ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionOpArrayCtorHandler), op_array)
	}
}
func DestroyZendFunction(function *ZendFunction) {
	var tmp Zval
	&tmp.GetValue().SetPtr(function)
	&tmp.SetTypeInfo(14)
	ZendFunctionDtor(&tmp)
}
func ZendFreeInternalArgInfo(function *ZendInternalFunction) {
	if (function.GetFnFlags()&(1<<13|1<<8)) != 0 && function.GetArgInfo() != nil {
		var i uint32
		var num_args uint32 = function.GetNumArgs() + 1
		var arg_info *ZendInternalArgInfo = function.GetArgInfo() - 1
		if (function.GetFnFlags() & 1 << 14) != 0 {
			num_args++
		}
		for i = 0; i < num_args; i++ {
			if arg_info[i].GetType() > 0x3ff {
				ZendStringReleaseEx((*ZendString)(arg_info[i].GetType() & ^0x3), 1)
			}
		}
		Free(arg_info)
	}
}
func ZendFunctionDtor(zv *Zval) {
	var function *ZendFunction = zv.GetValue().GetPtr()
	if function.GetType() == 2 {
		r.Assert(function.GetFunctionName() != nil)
		DestroyOpArray(&function.op_array)
	} else {
		r.Assert(function.GetType() == 1)
		r.Assert(function.GetFunctionName() != nil)
		ZendStringReleaseEx(function.GetFunctionName(), 1)

		/* For methods this will be called explicitly. */

		if function.GetScope() == nil {
			ZendFreeInternalArgInfo(&function.internal_function)
		}
		if (function.GetFnFlags() & 1 << 25) == 0 {
			g.CondF(true, func() { return Free(function) }, func() { return _efree(function) })
		}
	}
}
func ZendCleanupInternalClassData(ce *ZendClassEntry) {
	if (*Zval)(g.CondF((uintptr_t(ce).static_members_table__ptr&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(ce).static_members_table__ptr - 1)))
	}, func() any { return any(*(ce.GetStaticMembersTablePtr())) })) != nil {
		var static_members *Zval = (*Zval)(g.CondF((uintptr_t(ce).static_members_table__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(ce).static_members_table__ptr - 1)))
		}, func() any { return any(*(ce.GetStaticMembersTablePtr())) }))
		var p *Zval = static_members
		var end *Zval = p + ce.GetDefaultStaticMembersCount()
		if ce.GetStaticMembersTablePtr() == &ce.default_static_members_table {

			/* Special case: If this is a static property on a dl'ed internal class, then the
			 * static property table and the default property table are the same. In this case we
			 * destroy the values here, but leave behind valid UNDEF zvals and don't free the
			 * table itself. */

			for p != end {
				if p.GetType() == 10 {
					var prop_info *ZendPropertyInfo
					for {
						var _source_list *ZendPropertyInfoSourceList = &(p.GetValue().GetRef()).sources
						var _prop **ZendPropertyInfo
						var _end ***ZendPropertyInfo
						var _list *ZendPropertyInfoList
						if _source_list.GetPtr() != nil {
							if (_source_list.GetList() & 0x1) != 0 {
								_list = (*ZendPropertyInfoList)(_source_list.GetList() & ^0x1)
								_prop = _list.GetPtr()
								_end = _list.GetPtr() + _list.GetNum()
							} else {
								_prop = &_source_list.ptr
								_end = _prop + 1
							}
							for ; _prop < _end; _prop++ {
								prop_info = *_prop
								if prop_info.GetCe() == ce && p-static_members == prop_info.GetOffset() {
									ZendRefDelTypeSource(&(p.GetValue().GetRef()).sources, prop_info)
									break
								}
							}
						}
						break
					}
				}
				IZvalPtrDtor(p)
				p.SetTypeInfo(0)
				p++
			}

			/* Special case: If this is a static property on a dl'ed internal class, then the
			 * static property table and the default property table are the same. In this case we
			 * destroy the values here, but leave behind valid UNDEF zvals and don't free the
			 * table itself. */

		} else {
			if (uintPtr(ce.GetStaticMembersTablePtr()) & 1) != 0 {
				*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(ce.GetStaticMembersTablePtr()-1)))) = nil
			} else {
				*(ce.GetStaticMembersTablePtr()) = nil
			}
			for p != end {
				if p.GetType() == 10 {
					var prop_info *ZendPropertyInfo
					for {
						var _source_list *ZendPropertyInfoSourceList = &(p.GetValue().GetRef()).sources
						var _prop **ZendPropertyInfo
						var _end ***ZendPropertyInfo
						var _list *ZendPropertyInfoList
						if _source_list.GetPtr() != nil {
							if (_source_list.GetList() & 0x1) != 0 {
								_list = (*ZendPropertyInfoList)(_source_list.GetList() & ^0x1)
								_prop = _list.GetPtr()
								_end = _list.GetPtr() + _list.GetNum()
							} else {
								_prop = &_source_list.ptr
								_end = _prop + 1
							}
							for ; _prop < _end; _prop++ {
								prop_info = *_prop
								if prop_info.GetCe() == ce && p-static_members == prop_info.GetOffset() {
									ZendRefDelTypeSource(&(p.GetValue().GetRef()).sources, prop_info)
									break
								}
							}
						}
						break
					}
				}
				IZvalPtrDtor(p)
				p++
			}
			_efree(static_members)
		}
	}
}
func _destroyZendClassTraitsInfo(ce *ZendClassEntry) {
	var i uint32
	for i = 0; i < ce.GetNumTraits(); i++ {
		ZendStringReleaseEx(ce.GetTraitNames()[i].GetName(), 0)
		ZendStringReleaseEx(ce.GetTraitNames()[i].GetLcName(), 0)
	}
	_efree(ce.GetTraitNames())
	if ce.GetTraitAliases() != nil {
		i = 0
		for ce.GetTraitAliases()[i] != nil {
			if ce.GetTraitAliases()[i].GetTraitMethod().GetMethodName() != nil {
				ZendStringReleaseEx(ce.GetTraitAliases()[i].GetTraitMethod().GetMethodName(), 0)
			}
			if ce.GetTraitAliases()[i].GetTraitMethod().GetClassName() != nil {
				ZendStringReleaseEx(ce.GetTraitAliases()[i].GetTraitMethod().GetClassName(), 0)
			}
			if ce.GetTraitAliases()[i].GetAlias() != nil {
				ZendStringReleaseEx(ce.GetTraitAliases()[i].GetAlias(), 0)
			}
			_efree(ce.GetTraitAliases()[i])
			i++
		}
		_efree(ce.GetTraitAliases())
	}
	if ce.GetTraitPrecedences() != nil {
		var j uint32
		i = 0
		for ce.GetTraitPrecedences()[i] != nil {
			ZendStringReleaseEx(ce.GetTraitPrecedences()[i].GetTraitMethod().GetMethodName(), 0)
			ZendStringReleaseEx(ce.GetTraitPrecedences()[i].GetTraitMethod().GetClassName(), 0)
			for j = 0; j < ce.GetTraitPrecedences()[i].GetNumExcludes(); j++ {
				ZendStringReleaseEx(ce.GetTraitPrecedences()[i].GetExcludeClassNames()[j], 0)
			}
			_efree(ce.GetTraitPrecedences()[i])
			i++
		}
		_efree(ce.GetTraitPrecedences())
	}
}
func DestroyZendClass(zv *Zval) {
	var prop_info *ZendPropertyInfo
	var ce *ZendClassEntry = zv.GetValue().GetPtr()
	var fn *ZendFunction
	if (ce.GetCeFlags() & (1<<7 | 1<<10)) != 0 {
		var op_array *ZendOpArray
		if ce.GetDefaultStaticMembersCount() != 0 {
			ZendCleanupInternalClassData(ce)
		}
		if (ce.GetCeFlags() & 1 << 16) != 0 {
			for {
				var __ht *HashTable = &ce.function_table
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					op_array = _z.GetValue().GetPtr()
					if op_array.GetType() == 2 {
						DestroyOpArray(op_array)
					}
				}
				break
			}
		}
		return
	} else if g.PreDec(&(ce.GetRefcount())) > 0 {
		return
	}
	switch ce.GetType() {
	case 2:
		if ce.parent_name && (ce.GetCeFlags()&1<<19) == 0 {
			ZendStringReleaseEx(ce.parent_name, 0)
		}
		if ce.GetDefaultPropertiesTable() != nil {
			var p *Zval = ce.GetDefaultPropertiesTable()
			var end *Zval = p + ce.GetDefaultPropertiesCount()
			for p != end {
				IZvalPtrDtor(p)
				p++
			}
			_efree(ce.GetDefaultPropertiesTable())
		}
		if ce.GetDefaultStaticMembersTable() != nil {
			var p *Zval = ce.GetDefaultStaticMembersTable()
			var end *Zval = p + ce.GetDefaultStaticMembersCount()
			for p != end {
				if p.GetType() == 10 {
					var prop_info *ZendPropertyInfo
					for {
						var _source_list *ZendPropertyInfoSourceList = &(p.GetValue().GetRef()).sources
						var _prop **ZendPropertyInfo
						var _end ***ZendPropertyInfo
						var _list *ZendPropertyInfoList
						if _source_list.GetPtr() != nil {
							if (_source_list.GetList() & 0x1) != 0 {
								_list = (*ZendPropertyInfoList)(_source_list.GetList() & ^0x1)
								_prop = _list.GetPtr()
								_end = _list.GetPtr() + _list.GetNum()
							} else {
								_prop = &_source_list.ptr
								_end = _prop + 1
							}
							for ; _prop < _end; _prop++ {
								prop_info = *_prop
								if prop_info.GetCe() == ce && p-ce.GetDefaultStaticMembersTable() == prop_info.GetOffset() {
									ZendRefDelTypeSource(&(p.GetValue().GetRef()).sources, prop_info)
									break
								}
							}
						}
						break
					}
				}
				IZvalPtrDtor(p)
				p++
			}
			_efree(ce.GetDefaultStaticMembersTable())
		}
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
				if prop_info.GetCe() == ce {
					ZendStringReleaseEx(prop_info.GetName(), 0)
					if prop_info.GetDocComment() != nil {
						ZendStringReleaseEx(prop_info.GetDocComment(), 0)
					}
					if prop_info.GetType() > 0x3ff && (prop_info.GetType()&0x2) == 0 {
						ZendStringRelease((*ZendString)(prop_info.GetType() & ^0x3))
					}
				}
			}
			break
		}
		ZendHashDestroy(&ce.properties_info)
		ZendStringReleaseEx(ce.GetName(), 0)
		ZendHashDestroy(&ce.function_table)
		if &ce.constants_table.nNumOfElements {
			var c *ZendClassConstant
			for {
				var __ht *HashTable = &ce.constants_table
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					c = _z.GetValue().GetPtr()
					if c.GetCe() == ce {
						ZvalPtrDtorNogc(&c.value)
						if c.GetDocComment() != nil {
							ZendStringReleaseEx(c.GetDocComment(), 0)
						}
					}
				}
				break
			}
		}
		ZendHashDestroy(&ce.constants_table)
		if ce.GetNumInterfaces() > 0 {
			if (ce.GetCeFlags() & 1 << 20) == 0 {
				var i uint32
				for i = 0; i < ce.GetNumInterfaces(); i++ {
					ZendStringReleaseEx(ce.interface_names[i].name, 0)
					ZendStringReleaseEx(ce.interface_names[i].lc_name, 0)
				}
			}
			_efree(ce.interfaces)
		}
		if ce.GetDocComment() != nil {
			ZendStringReleaseEx(ce.GetDocComment(), 0)
		}
		if ce.GetNumTraits() > 0 {
			_destroyZendClassTraitsInfo(ce)
		}
		break
	case 1:
		if ce.GetDefaultPropertiesTable() != nil {
			var p *Zval = ce.GetDefaultPropertiesTable()
			var end *Zval = p + ce.GetDefaultPropertiesCount()
			for p != end {
				ZvalInternalPtrDtor(p)
				p++
			}
			Free(ce.GetDefaultPropertiesTable())
		}
		if ce.GetDefaultStaticMembersTable() != nil {
			var p *Zval = ce.GetDefaultStaticMembersTable()
			var end *Zval = p + ce.GetDefaultStaticMembersCount()
			for p != end {
				ZvalInternalPtrDtor(p)
				p++
			}
			Free(ce.GetDefaultStaticMembersTable())
			if ce.GetStaticMembersTablePtr() != &ce.default_static_members_table {
				ZendCleanupInternalClassData(ce)
			}
		}
		ZendHashDestroy(&ce.properties_info)
		ZendStringReleaseEx(ce.GetName(), 1)

		/* TODO: eliminate this loop for classes without functions with arg_info */

		for {
			var __ht *HashTable = &ce.function_table
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				fn = _z.GetValue().GetPtr()
				if (fn.GetFnFlags()&(1<<13|1<<8)) != 0 && fn.GetScope() == ce {
					ZendFreeInternalArgInfo(&fn.internal_function)
				}
			}
			break
		}
		ZendHashDestroy(&ce.function_table)
		if &ce.constants_table.nNumOfElements {
			var c *ZendClassConstant
			for {
				var __ht *HashTable = &ce.constants_table
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					c = _z.GetValue().GetPtr()
					if c.GetCe() == ce {
						ZvalInternalPtrDtor(&c.value)
						if c.GetDocComment() != nil {
							ZendStringReleaseEx(c.GetDocComment(), 1)
						}
					}
					Free(c)
				}
				break
			}
			ZendHashDestroy(&ce.constants_table)
		}
		if ce.GetIteratorFuncsPtr() != nil {
			Free(ce.GetIteratorFuncsPtr())
		}
		if ce.GetNumInterfaces() > 0 {
			Free(ce.interfaces)
		}
		if ce.GetPropertiesInfoTable() != nil {
			Free(ce.GetPropertiesInfoTable())
		}
		Free(ce)
		break
	}
}
func ZendClassAddRef(zv *Zval) {
	var ce *ZendClassEntry = zv.GetValue().GetPtr()
	if (ce.GetCeFlags() & 1 << 7) == 0 {
		ce.GetRefcount()++
	}
}
func DestroyOpArray(op_array *ZendOpArray) {
	var i uint32
	if op_array.GetStaticVariables() != nil {
		var ht *HashTable = g.CondF((uintPtr(op_array.GetStaticVariablesPtrPtr())&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(op_array.GetStaticVariablesPtrPtr()-1))))
		}, func() any { return any(*(op_array.GetStaticVariablesPtrPtr())) })
		if ht != nil && (ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<6) == 0 {
			if ZendGcDelref(&ht.gc) == 0 {
				ZendArrayDestroy(ht)
			}
		}
	}
	if (op_array.GetFnFlags()&1<<22) != 0 && op_array.GetRunTimeCachePtr() != nil {
		_efree(op_array.GetRunTimeCachePtr())
	}
	if op_array.GetRefcount() == nil || g.PreDec(&((*op_array).refcount)) > 0 {
		return
	}
	_efree(op_array.GetRefcount())
	if op_array.GetVars() != nil {
		i = op_array.GetLastVar()
		for i > 0 {
			i--
			ZendStringReleaseEx(op_array.GetVars()[i], 0)
		}
		_efree(op_array.GetVars())
	}
	if op_array.GetLiterals() != nil {
		var literal *Zval = op_array.GetLiterals()
		var end *Zval = literal + op_array.GetLastLiteral()
		for literal < end {
			ZvalPtrDtorNogc(literal)
			literal++
		}
		if (op_array.GetFnFlags() & 1 << 25) == 0 {
			_efree(op_array.GetLiterals())
		}
	}
	_efree(op_array.GetOpcodes())
	if op_array.GetFunctionName() != nil {
		ZendStringReleaseEx(op_array.GetFunctionName(), 0)
	}
	if op_array.GetDocComment() != nil {
		ZendStringReleaseEx(op_array.GetDocComment(), 0)
	}
	if op_array.GetLiveRange() != nil {
		_efree(op_array.GetLiveRange())
	}
	if op_array.GetTryCatchArray() != nil {
		_efree(op_array.GetTryCatchArray())
	}
	if (ZendExtensionFlags & 1 << 1) != 0 {
		if (op_array.GetFnFlags() & 1 << 25) != 0 {
			ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionOpArrayDtorHandler), op_array)
		}
	}
	if op_array.GetArgInfo() != nil {
		var num_args uint32 = op_array.GetNumArgs()
		var arg_info *ZendArgInfo = op_array.GetArgInfo()
		if (op_array.GetFnFlags() & 1 << 13) != 0 {
			arg_info--
			num_args++
		}
		if (op_array.GetFnFlags() & 1 << 14) != 0 {
			num_args++
		}
		for i = 0; i < num_args; i++ {
			if arg_info[i].GetName() != nil {
				ZendStringReleaseEx(arg_info[i].GetName(), 0)
			}
			if arg_info[i].GetType() > 0x3ff {
				ZendStringReleaseEx((*ZendString)(arg_info[i].GetType() & ^0x3), 0)
			}
		}
		_efree(arg_info)
	}
}
func ZendUpdateExtendedStmts(op_array *ZendOpArray) {
	var opline *ZendOp = op_array.GetOpcodes()
	var end *ZendOp = opline + op_array.GetLast()
	for opline < end {
		if opline.GetOpcode() == 101 {
			if opline+1 < end {
				if (opline + 1).GetOpcode() == 101 {
					opline.SetOpcode(0)
					opline++
					continue
				}
				if opline+1 < end {
					opline.SetLineno((opline + 1).GetLineno())
				}
			} else {
				opline.SetOpcode(0)
			}
		}
		opline++
	}
}
func ZendExtensionOpArrayHandler(extension *ZendExtension, op_array *ZendOpArray) {
	if extension.GetOpArrayHandler() != nil {
		extension.GetOpArrayHandler()(op_array)
	}
}
func ZendCheckFinallyBreakout(op_array *ZendOpArray, op_num uint32, dst_num uint32) {
	var i int
	for i = 0; i < op_array.GetLastTryCatch(); i++ {
		if (op_num < op_array.GetTryCatchArray()[i].GetFinallyOp() || op_num >= op_array.GetTryCatchArray()[i].GetFinallyEnd()) && (dst_num >= op_array.GetTryCatchArray()[i].GetFinallyOp() && dst_num <= op_array.GetTryCatchArray()[i].GetFinallyEnd()) {
			CG.SetInCompilation(1)
			CG.SetActiveOpArray(op_array)
			CG.SetZendLineno(op_array.GetOpcodes()[op_num].GetLineno())
			ZendErrorNoreturn(1<<6, "jump into a finally block is disallowed")
		} else if op_num >= op_array.GetTryCatchArray()[i].GetFinallyOp() && op_num <= op_array.GetTryCatchArray()[i].GetFinallyEnd() && (dst_num > op_array.GetTryCatchArray()[i].GetFinallyEnd() || dst_num < op_array.GetTryCatchArray()[i].GetFinallyOp()) {
			CG.SetInCompilation(1)
			CG.SetActiveOpArray(op_array)
			CG.SetZendLineno(op_array.GetOpcodes()[op_num].GetLineno())
			ZendErrorNoreturn(1<<6, "jump out of a finally block is disallowed")
		}
	}
}
func ZendGetBrkContTarget(op_array *ZendOpArray, opline *ZendOp) uint32 {
	var nest_levels int = opline.GetOp2().GetNum()
	var array_offset int = opline.GetOp1().GetNum()
	var jmp_to *ZendBrkContElement
	for {
		jmp_to = &CG.context.GetBrkContArray()[array_offset]
		if nest_levels > 1 {
			array_offset = jmp_to.GetParent()
		}
		if g.PreDec(&nest_levels) <= 0 {
			break
		}
	}
	if opline.GetOpcode() == 254 {
		return jmp_to.GetBrk()
	} else {
		return jmp_to.GetCont()
	}
}
func EmitLiveRangeRaw(op_array *ZendOpArray, var_num uint32, kind uint32, start uint32, end uint32) {
	var range_ *ZendLiveRange
	op_array.GetLastLiveRange()++
	op_array.SetLiveRange(_erealloc(op_array.GetLiveRange(), g.SizeOf("zend_live_range")*op_array.GetLastLiveRange()))
	r.Assert(start < end)
	range_ = &op_array.live_range[op_array.GetLastLiveRange()-1]
	range_.SetVar(uint32(intptr_t)((*Zval)(nil) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(op_array.GetLastVar()+var_num))))
	range_.SetVar(range_.GetVar() | kind)
	range_.SetStart(start)
	range_.SetEnd(end)
}
func EmitLiveRange(op_array *ZendOpArray, var_num uint32, start uint32, end uint32, needs_live_range ZendNeedsLiveRangeCb) {
	var def_opline *ZendOp = &op_array.opcodes[start]
	var orig_def_opline *ZendOp = def_opline
	var use_opline *ZendOp = &op_array.opcodes[end]
	var kind uint32
	switch def_opline.GetOpcode() {
	case 72:

	case 147:

	case 55:
		r.Assert(false)
		return
	case 46:

	case 47:

	case 52:

	case 14:

	case 109:

	case 146:

	case 162:
		return
	case 57:
		kind = 2
		start++
		break
	case 54:
		kind = 3

		/* ROPE live ranges include the generating opcode. */

		def_opline--
		break
	case 77:

	case 125:
		kind = 1
		start++
		break
	case 68:
		var level int = 0
		var orig_start uint32 = start
		for def_opline+1 < use_opline {
			def_opline++
			start++
			if def_opline.GetOpcode() == 60 {
				if level == 0 {
					break
				}
				level--
			} else {
				switch def_opline.GetOpcode() {
				case 61:

				case 59:

				case 69:

				case 128:

				case 118:

				case 112:

				case 113:

				case 68:
					level++
					break
				case 129:

				case 130:

				case 131:
					level--
					break
				}
			}
		}
		EmitLiveRangeRaw(op_array, var_num, 4, orig_start+1, start+1)
		if start+1 == end {

			/* Trivial live-range, no need to store it. */

			return

			/* Trivial live-range, no need to store it. */

		}
	default:
		start++
		kind = 0

		/* Check hook to determine whether a live range is necessary,
		 * e.g. based on type info. */

		if needs_live_range != nil && needs_live_range(op_array, orig_def_opline) == 0 {
			return
		}
		break
	case 167:

		/* COPY_TMP has a split live-range: One from the definition until the use in
		 * "null" branch, and another from the start of the "non-null" branch to the
		 * FREE opcode. */

		var rt_var_num uint32 = uint32(intptr_t)((*Zval)(nil) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(op_array.GetLastVar()+var_num)))
		var block_start_op *ZendOp = use_opline
		if needs_live_range != nil && needs_live_range(op_array, orig_def_opline) == 0 {
			return
		}
		for (block_start_op - 1).opcode == 70 {
			block_start_op--
		}
		kind = 0
		start = block_start_op - op_array.GetOpcodes()
		if start != end {
			EmitLiveRangeRaw(op_array, var_num, kind, start, end)
		}
		for {
			use_opline--
			if (use_opline.GetOp1Type()&(1<<1|1<<2)) != 0 && use_opline.GetOp1().GetVar() == rt_var_num || (use_opline.GetOp2Type()&(1<<1|1<<2)) != 0 && use_opline.GetOp2().GetVar() == rt_var_num {
				break
			}
		}
		start = def_opline + 1 - op_array.GetOpcodes()
		end = use_opline - op_array.GetOpcodes()
		EmitLiveRangeRaw(op_array, var_num, kind, start, end)
		return
	}
	EmitLiveRangeRaw(op_array, var_num, kind, start, end)
}
func IsFakeDef(opline *ZendOp) ZendBool {
	/* These opcodes only modify the result, not create it. */

	return opline.GetOpcode() == 55 || opline.GetOpcode() == 72 || opline.GetOpcode() == 147

	/* These opcodes only modify the result, not create it. */
}
func KeepsOp1Alive(opline *ZendOp) ZendBool {
	/* These opcodes don't consume their OP1 operand,
	 * it is later freed by something else. */

	if opline.GetOpcode() == 48 || opline.GetOpcode() == 187 || opline.GetOpcode() == 98 || opline.GetOpcode() == 167 {
		return 1
	}
	r.Assert(opline.GetOpcode() != 188 && opline.GetOpcode() != 78 && opline.GetOpcode() != 126 && opline.GetOpcode() != 155 && opline.GetOpcode() != 124 && opline.GetOpcode() != 182 && opline.GetOpcode() != 55)
	return 0
}

/* Live ranges must be sorted by increasing start opline */

func CmpLiveRange(a *ZendLiveRange, b *ZendLiveRange) int { return a.GetStart() - b.GetStart() }
func SwapLiveRange(a *ZendLiveRange, b *ZendLiveRange) {
	var tmp uint32
	tmp = a.GetVar()
	a.SetVar(b.GetVar())
	b.SetVar(tmp)
	tmp = a.GetStart()
	a.SetStart(b.GetStart())
	b.SetStart(tmp)
	tmp = a.GetEnd()
	a.SetEnd(b.GetEnd())
	b.SetEnd(tmp)
}
func ZendCalcLiveRanges(op_array *ZendOpArray, needs_live_range ZendNeedsLiveRangeCb) {
	var opnum uint32 = op_array.GetLast()
	var opline *ZendOp = &op_array.opcodes[opnum]
	var var_offset uint32 = op_array.GetLastVar()
	var last_use *uint32 = _emalloc(g.SizeOf("uint32_t") * op_array.GetT())
	memset(last_use, -1, g.SizeOf("uint32_t")*op_array.GetT())
	r.Assert(op_array.GetLiveRange() == nil)
	for opnum > 0 {
		opnum--
		opline--
		if (opline.GetResultType()&(1<<1|1<<2)) != 0 && IsFakeDef(opline) == 0 {
			var var_num uint32 = uint32((*Zval)((*byte)(nil)+int(opline.GetResult().GetVar()))-((*Zval)(nil)+(int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1)))+int(0)))) - var_offset

			/* Defs without uses can occur for two reasons: Either because the result is
			 * genuinely unused (e.g. omitted FREE opcode for an unused boolean result), or
			 * because there are multiple defining opcodes (e.g. JMPZ_EX and QM_ASSIGN), in
			 * which case the last one starts the live range. As such, we can simply ignore
			 * missing uses here. */

			if last_use[var_num] != uint32-1 {

				/* Skip trivial live-range */

				if opnum+1 != last_use[var_num] {
					var num uint32

					/* OP_DATA uses only op1 operand */

					r.Assert(opline.GetOpcode() != 137)
					num = opnum
					EmitLiveRange(op_array, var_num, num, last_use[var_num], needs_live_range)
				}
				last_use[var_num] = uint32 - 1
			}

			/* Defs without uses can occur for two reasons: Either because the result is
			 * genuinely unused (e.g. omitted FREE opcode for an unused boolean result), or
			 * because there are multiple defining opcodes (e.g. JMPZ_EX and QM_ASSIGN), in
			 * which case the last one starts the live range. As such, we can simply ignore
			 * missing uses here. */

		}
		if (opline.GetOp1Type() & (1<<1 | 1<<2)) != 0 {
			var var_num uint32 = uint32((*Zval)((*byte)(nil)+int(opline.GetOp1().GetVar()))-((*Zval)(nil)+(int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1)))+int(0)))) - var_offset
			if last_use[var_num] == uint32-1 {
				if KeepsOp1Alive(opline) == 0 {

					/* OP_DATA is really part of the previous opcode. */

					last_use[var_num] = opnum - (opline.GetOpcode() == 137)

					/* OP_DATA is really part of the previous opcode. */

				}
			}
		}
		if (opline.GetOp2Type() & (1<<1 | 1<<2)) != 0 {
			var var_num uint32 = uint32((*Zval)((*byte)(nil)+int(opline.GetOp2().GetVar()))-((*Zval)(nil)+(int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1)))+int(0)))) - var_offset
			if opline.GetOpcode() == 78 || opline.GetOpcode() == 126 {

				/* OP2 of FE_FETCH is actually a def, not a use. */

				if last_use[var_num] != uint32-1 {
					if opnum+1 != last_use[var_num] {
						EmitLiveRange(op_array, var_num, opnum, last_use[var_num], needs_live_range)
					}
					last_use[var_num] = uint32 - 1
				}

				/* OP2 of FE_FETCH is actually a def, not a use. */

			} else if last_use[var_num] == uint32-1 {

				/* OP_DATA uses only op1 operand */

				r.Assert(opline.GetOpcode() != 137)
				last_use[var_num] = opnum
			}
		}
	}
	if op_array.GetLastLiveRange() > 1 {
		var r1 *ZendLiveRange = op_array.GetLiveRange()
		var r2 *ZendLiveRange = r1 + op_array.GetLastLiveRange() - 1

		/* In most cases we need just revert the array */

		for r1 < r2 {
			SwapLiveRange(r1, r2)
			r1++
			r2--
		}
		r1 = op_array.GetLiveRange()
		r2 = r1 + op_array.GetLastLiveRange() - 1
		for r1 < r2 {
			if r1.GetStart() > (r1 + 1).GetStart() {
				ZendSort(r1, r2-r1+1, g.SizeOf("zend_live_range"), CompareFuncT(CmpLiveRange), SwapFuncT(SwapLiveRange))
				break
			}
			r1++
		}
	}
	_efree(last_use)
}
func ZendRecalcLiveRanges(op_array *ZendOpArray, needs_live_range ZendNeedsLiveRangeCb) {
	/* We assume that we never create live-ranges where there were none before. */

	r.Assert(op_array.GetLiveRange() != nil)
	_efree(op_array.GetLiveRange())
	op_array.SetLiveRange(nil)
	op_array.SetLastLiveRange(0)
	ZendCalcLiveRanges(op_array, needs_live_range)
}
func PassTwo(op_array *ZendOpArray) int {
	var opline *ZendOp
	var end *ZendOp
	if (op_array.GetType() & 1) != 0 {
		return 0
	}
	if (CG.GetCompilerOptions() & 1 << 0) != 0 {
		ZendUpdateExtendedStmts(op_array)
	}
	if (CG.GetCompilerOptions() & 1 << 2) != 0 {
		if (ZendExtensionFlags & 1 << 2) != 0 {
			ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionOpArrayHandler), op_array)
		}
	}
	if CG.GetContext().GetVarsSize() != op_array.GetLastVar() {
		op_array.SetVars((**ZendString)(_erealloc(op_array.GetVars(), g.SizeOf("zend_string *")*op_array.GetLastVar())))
		CG.GetContext().SetVarsSize(op_array.GetLastVar())
	}
	op_array.SetOpcodes((*ZendOp)(_erealloc(op_array.GetOpcodes(), (g.SizeOf("zend_op")*op_array.GetLast() + (16-1) & ^(16-1))+g.SizeOf("zval")*op_array.GetLastLiteral())))
	if op_array.GetLiterals() != nil {
		memcpy((*byte)(op_array.GetOpcodes())+(g.SizeOf("zend_op")*op_array.GetLast() + (16-1) & ^(16-1)), op_array.GetLiterals(), g.SizeOf("zval")*op_array.GetLastLiteral())
		_efree(op_array.GetLiterals())
		op_array.SetLiterals((*Zval)((*byte)(op_array.GetOpcodes()) + (g.SizeOf("zend_op")*op_array.GetLast() + (16-1) & ^(16-1))))
	}
	CG.GetContext().SetOpcodesSize(op_array.GetLast())
	CG.GetContext().SetLiteralsSize(op_array.GetLastLiteral())

	/* Needs to be set directly after the opcode/literal reallocation, to ensure destruction
	 * happens correctly if any of the following fixups generate a fatal error. */

	op_array.SetFnFlags(op_array.GetFnFlags() | 1<<25)
	opline = op_array.GetOpcodes()
	end = opline + op_array.GetLast()
	for opline < end {
		switch opline.GetOpcode() {
		case 64:
			var val *Zval = CG.GetActiveOpArray().GetLiterals() + opline.GetOp2().GetConstant()
			if val.GetType() == 11 {
				var slot uint32 = op_array.GetCacheSize() + (8-1) & ^(8-1)
				val.SetCacheSlot(slot)
				op_array.SetCacheSize(op_array.GetCacheSize() + g.SizeOf("zval"))
			}
			break
		case 162:
			opline.GetOp1().SetOplineNum(op_array.GetTryCatchArray()[opline.GetOp1().GetNum()].GetFinallyOp())
			opline.GetOp1().SetJmpOffset((*byte)(&op_array.opcodes[opline.GetOp1().GetOplineNum()] - (*byte)(opline)))
			break
		case 254:

		case 255:
			var jmp_target uint32 = ZendGetBrkContTarget(op_array, opline)
			if (op_array.GetFnFlags() & 1 << 15) != 0 {
				ZendCheckFinallyBreakout(op_array, opline-op_array.GetOpcodes(), jmp_target)
			}
			opline.SetOpcode(42)
			opline.GetOp1().SetOplineNum(jmp_target)
			opline.GetOp2().SetNum(0)
			opline.GetOp1().SetJmpOffset((*byte)(&op_array.opcodes[opline.GetOp1().GetOplineNum()] - (*byte)(opline)))
			break
		case 253:
			ZendResolveGotoLabel(op_array, opline)
			if (op_array.GetFnFlags() & 1 << 15) != 0 {
				ZendCheckFinallyBreakout(op_array, opline-op_array.GetOpcodes(), opline.GetOp1().GetOplineNum())
			}
		case 42:
			opline.GetOp1().SetJmpOffset((*byte)(&op_array.opcodes[opline.GetOp1().GetOplineNum()] - (*byte)(opline)))
			break
		case 45:

			/* absolute index to relative offset */

			opline.SetExtendedValue((*byte)(&op_array.opcodes[opline.GetExtendedValue()] - (*byte)(opline)))
		case 43:

		case 44:

		case 46:

		case 47:

		case 152:

		case 169:

		case 77:

		case 125:
			opline.GetOp2().SetJmpOffset((*byte)(&op_array.opcodes[opline.GetOp2().GetOplineNum()] - (*byte)(opline)))
			break
		case 151:

			/* If result of assert is unused, result of check is unused as well */

			var call *ZendOp = &op_array.opcodes[opline.GetOp2().GetOplineNum()-1]
			if call.GetOpcode() == 103 {
				call--
			}
			if call.GetResultType() == 0 {
				opline.SetResultType(0)
			}
			opline.GetOp2().SetJmpOffset((*byte)(&op_array.opcodes[opline.GetOp2().GetOplineNum()] - (*byte)(opline)))
			break
		case 78:

		case 126:

			/* absolute index to relative offset */

			opline.SetExtendedValue((*byte)(&op_array.opcodes[opline.GetExtendedValue()] - (*byte)(opline)))
			break
		case 107:
			if (opline.GetExtendedValue() & 1 << 0) == 0 {
				opline.GetOp2().SetJmpOffset((*byte)(&op_array.opcodes[opline.GetOp2().GetOplineNum()] - (*byte)(opline)))
			}
			break
		case 62:

		case 111:
			if (op_array.GetFnFlags() & 1 << 24) != 0 {
				opline.SetOpcode(161)
			}
			break
		case 187:

		case 188:

			/* absolute indexes to relative offsets */

			var jumptable *HashTable = (CG.GetActiveOpArray().GetLiterals() + opline.GetOp2().GetConstant()).value.arr
			var zv *Zval
			for {
				var __ht *HashTable = jumptable
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					zv = _z
					zv.GetValue().SetLval((*byte)(&op_array.opcodes[zv.GetValue().GetLval()] - (*byte)(opline)))
				}
				break
			}
			opline.SetExtendedValue((*byte)(&op_array.opcodes[opline.GetExtendedValue()] - (*byte)(opline)))
			break
		}
		if opline.GetOp1Type() == 1<<0 {
			opline.GetOp1().SetConstant((*byte)(op_array.GetLiterals()+opline.GetOp1().GetConstant()) - (*byte)(opline))
		} else if (opline.GetOp1Type() & (1<<2 | 1<<1)) != 0 {
			opline.GetOp1().SetVar(uint32(zend_intptr_t)((*Zval)(nil) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(op_array.GetLastVar()+opline.GetOp1().GetVar()))))
		}
		if opline.GetOp2Type() == 1<<0 {
			opline.GetOp2().SetConstant((*byte)(op_array.GetLiterals()+opline.GetOp2().GetConstant()) - (*byte)(opline))
		} else if (opline.GetOp2Type() & (1<<2 | 1<<1)) != 0 {
			opline.GetOp2().SetVar(uint32(zend_intptr_t)((*Zval)(nil) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(op_array.GetLastVar()+opline.GetOp2().GetVar()))))
		}
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			opline.GetResult().SetVar(uint32(zend_intptr_t)((*Zval)(nil) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(op_array.GetLastVar()+opline.GetResult().GetVar()))))
		}
		ZendVmSetOpcodeHandler(opline)
		opline++
	}
	ZendCalcLiveRanges(op_array, nil)
	return 0
}
func GetUnaryOp(opcode int) UnaryOpType {
	switch opcode {
	case 13:
		return UnaryOpType(BitwiseNotFunction)
	case 14:
		return UnaryOpType(BooleanNotFunction)
	default:
		return UnaryOpType(nil)
	}
}
func GetBinaryOp(opcode int) BinaryOpType {
	switch opcode {
	case 1:
		return BinaryOpType(AddFunction)
	case 2:
		return BinaryOpType(SubFunction)
	case 3:
		return BinaryOpType(MulFunction)
	case 12:
		return BinaryOpType(PowFunction)
	case 4:
		return BinaryOpType(DivFunction)
	case 5:
		return BinaryOpType(ModFunction)
	case 6:
		return BinaryOpType(ShiftLeftFunction)
	case 7:
		return BinaryOpType(ShiftRightFunction)
	case 252:

	case 53:

	case 8:
		return BinaryOpType(ConcatFunction)
	case 16:
		return BinaryOpType(IsIdenticalFunction)
	case 17:
		return BinaryOpType(IsNotIdenticalFunction)
	case 18:

	case 48:
		return BinaryOpType(IsEqualFunction)
	case 19:
		return BinaryOpType(IsNotEqualFunction)
	case 20:
		return BinaryOpType(IsSmallerFunction)
	case 21:
		return BinaryOpType(IsSmallerOrEqualFunction)
	case 170:
		return BinaryOpType(CompareFunction)
	case 9:
		return BinaryOpType(BitwiseOrFunction)
	case 10:
		return BinaryOpType(BitwiseAndFunction)
	case 11:
		return BinaryOpType(BitwiseXorFunction)
	case 15:
		return BinaryOpType(BooleanXorFunction)
	default:
		r.Assert(false)
		return BinaryOpType(nil)
	}
}
