// <<generate>>

package zend

import (
	b "sik/builtin"
)

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
	op_array.SetRefcount((*uint32)(Emalloc(b.SizeOf("uint32_t"))))
	op_array.refcount = 1
	op_array.SetLast(0)
	op_array.SetOpcodes(Emalloc(initial_ops_size * b.SizeOf("zend_op")))
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
	ZEND_MAP_PTR_INIT(op_array.static_variables_ptr, op_array.GetStaticVariables())
	op_array.SetLastTryCatch(0)
	op_array.SetFnFlags(0)
	op_array.SetLastLiteral(0)
	op_array.SetLiterals(nil)
	ZEND_MAP_PTR_INIT(op_array.run_time_cache, nil)
	op_array.SetCacheSize(ZendOpArrayExtensionHandles * b.SizeOf("void *"))
	memset(op_array.GetReserved(), 0, ZEND_MAX_RESERVED_RESOURCES*b.SizeOf("void *"))
	if (ZendExtensionFlags & ZEND_EXTENSIONS_HAVE_OP_ARRAY_CTOR) != 0 {
		ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionOpArrayCtorHandler), op_array)
	}
}
func DestroyZendFunction(function *ZendFunction) {
	var tmp Zval
	ZVAL_PTR(&tmp, function)
	ZendFunctionDtor(&tmp)
}
func ZendFreeInternalArgInfo(function *ZendInternalFunction) {
	if function.HasFnFlags(ZEND_ACC_HAS_RETURN_TYPE|ZEND_ACC_HAS_TYPE_HINTS) && function.GetArgInfo() != nil {
		var i uint32
		var num_args uint32 = function.GetNumArgs() + 1
		var arg_info *ZendInternalArgInfo = function.GetArgInfo() - 1
		if function.IsVariadic() {
			num_args++
		}
		for i = 0; i < num_args; i++ {
			if arg_info[i].GetType().IsClass() {
				ZendStringReleaseEx(arg_info[i].GetType().Name(), 1)
			}
		}
		Free(arg_info)
	}
}
func ZendFunctionDtor(zv *Zval) {
	var function *ZendFunction = zv.GetPtr()
	if function.GetType() == ZEND_USER_FUNCTION {
		ZEND_ASSERT(function.GetFunctionName() != nil)
		DestroyOpArray(function.GetOpArray())
	} else {
		ZEND_ASSERT(function.GetType() == ZEND_INTERNAL_FUNCTION)
		ZEND_ASSERT(function.GetFunctionName() != nil)
		ZendStringReleaseEx(function.GetFunctionName(), 1)

		/* For methods this will be called explicitly. */

		if function.GetScope() == nil {
			ZendFreeInternalArgInfo(function.GetInternalFunction())
		}
		if !function.IsArenaAllocated() {
			Pefree(function, 1)
		}
	}
}
func ZendCleanupInternalClassData(ce *ZendClassEntry) {
	if CE_STATIC_MEMBERS(ce) != nil {
		var static_members *Zval = CE_STATIC_MEMBERS(ce)
		var p *Zval = static_members
		var end *Zval = p + ce.GetDefaultStaticMembersCount()
		if ce.GetStaticMembersTablePtr() == ce.GetDefaultStaticMembersTable() {

			/* Special case: If this is a static property on a dl'ed internal class, then the
			 * static property table and the default property table are the same. In this case we
			 * destroy the values here, but leave behind valid UNDEF zvals and don't free the
			 * table itself. */

			for p != end {
				if Z_ISREF_P(p) {
					var prop_info *ZendPropertyInfo
					for {
						var _source_list *ZendPropertyInfoSourceList = &(Z_REF_P(p).GetSources())
						var _prop **ZendPropertyInfo
						var _end ***ZendPropertyInfo
						var _list *ZendPropertyInfoList
						if _source_list.GetPtr() != nil {
							if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
								_list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
								_prop = _list.GetPtr()
								_end = _list.GetPtr() + _list.GetNum()
							} else {
								_prop = _source_list.GetPtr()
								_end = _prop + 1
							}
							for ; _prop < _end; _prop++ {
								prop_info = *_prop
								if prop_info.GetCe() == ce && p-static_members == prop_info.GetOffset() {
									ZEND_REF_DEL_TYPE_SOURCE(p.GetRef(), prop_info)
									break
								}
							}
						}
						break
					}
				}
				IZvalPtrDtor(p)
				ZVAL_UNDEF(p)
				p++
			}

			/* Special case: If this is a static property on a dl'ed internal class, then the
			 * static property table and the default property table are the same. In this case we
			 * destroy the values here, but leave behind valid UNDEF zvals and don't free the
			 * table itself. */

		} else {
			ZEND_MAP_PTR_SET(ce.static_members_table, nil)
			for p != end {
				if Z_ISREF_P(p) {
					var prop_info *ZendPropertyInfo
					for {
						var _source_list *ZendPropertyInfoSourceList = &(Z_REF_P(p).GetSources())
						var _prop **ZendPropertyInfo
						var _end ***ZendPropertyInfo
						var _list *ZendPropertyInfoList
						if _source_list.GetPtr() != nil {
							if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
								_list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
								_prop = _list.GetPtr()
								_end = _list.GetPtr() + _list.GetNum()
							} else {
								_prop = _source_list.GetPtr()
								_end = _prop + 1
							}
							for ; _prop < _end; _prop++ {
								prop_info = *_prop
								if prop_info.GetCe() == ce && p-static_members == prop_info.GetOffset() {
									ZEND_REF_DEL_TYPE_SOURCE(p.GetRef(), prop_info)
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
			Efree(static_members)
		}
	}
}
func _destroyZendClassTraitsInfo(ce *ZendClassEntry) {
	var i uint32
	for i = 0; i < ce.GetNumTraits(); i++ {
		ZendStringReleaseEx(ce.GetTraitNames()[i].GetName(), 0)
		ZendStringReleaseEx(ce.GetTraitNames()[i].GetLcName(), 0)
	}
	Efree(ce.GetTraitNames())
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
			Efree(ce.GetTraitAliases()[i])
			i++
		}
		Efree(ce.GetTraitAliases())
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
			Efree(ce.GetTraitPrecedences()[i])
			i++
		}
		Efree(ce.GetTraitPrecedences())
	}
}
func DestroyZendClass(zv *Zval) {
	var prop_info *ZendPropertyInfo
	var ce *ZendClassEntry = zv.GetPtr()
	var fn *ZendFunction
	if ce.HasCeFlags(ZEND_ACC_IMMUTABLE | ZEND_ACC_PRELOADED) {
		var op_array *ZendOpArray
		if ce.GetDefaultStaticMembersCount() != 0 {
			ZendCleanupInternalClassData(ce)
		}
		if ce.IsHasStaticInMethods() {
			var __ht *HashTable = ce.GetFunctionTable()
			for _, _p := range __ht.foreachData() {
				var _z *Zval = _p.GetVal()

				op_array = _z.GetPtr()
				if op_array.GetType() == ZEND_USER_FUNCTION {
					DestroyOpArray(op_array)
				}
			}
		}
		return
	} else if b.PreDec(&(ce.GetRefcount())) > 0 {
		return
	}
	switch ce.GetType() {
	case ZEND_USER_CLASS:
		if ce.parent_name && !ce.IsResolvedParent() {
			ZendStringReleaseEx(ce.parent_name, 0)
		}
		if ce.GetDefaultPropertiesTable() != nil {
			var p *Zval = ce.GetDefaultPropertiesTable()
			var end *Zval = p + ce.GetDefaultPropertiesCount()
			for p != end {
				IZvalPtrDtor(p)
				p++
			}
			Efree(ce.GetDefaultPropertiesTable())
		}
		if ce.GetDefaultStaticMembersTable() != nil {
			var p *Zval = ce.GetDefaultStaticMembersTable()
			var end *Zval = p + ce.GetDefaultStaticMembersCount()
			for p != end {
				if Z_ISREF_P(p) {
					var prop_info *ZendPropertyInfo
					for {
						var _source_list *ZendPropertyInfoSourceList = &(Z_REF_P(p).GetSources())
						var _prop **ZendPropertyInfo
						var _end ***ZendPropertyInfo
						var _list *ZendPropertyInfoList
						if _source_list.GetPtr() != nil {
							if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(_source_list.GetList()) != 0 {
								_list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(_source_list.GetList())
								_prop = _list.GetPtr()
								_end = _list.GetPtr() + _list.GetNum()
							} else {
								_prop = _source_list.GetPtr()
								_end = _prop + 1
							}
							for ; _prop < _end; _prop++ {
								prop_info = *_prop
								if prop_info.GetCe() == ce && p-ce.GetDefaultStaticMembersTable() == prop_info.GetOffset() {
									ZEND_REF_DEL_TYPE_SOURCE(p.GetRef(), prop_info)
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
			Efree(ce.GetDefaultStaticMembersTable())
		}
		var __ht *HashTable = ce.GetPropertiesInfo()
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			prop_info = _z.GetPtr()
			if prop_info.GetCe() == ce {
				ZendStringReleaseEx(prop_info.GetName(), 0)
				if prop_info.GetDocComment() != nil {
					ZendStringReleaseEx(prop_info.GetDocComment(), 0)
				}
				if prop_info.GetType().IsName() {
					ZendStringRelease(prop_info.GetType().Name())
				}
			}
		}
		ZendHashDestroy(ce.GetPropertiesInfo())
		ZendStringReleaseEx(ce.GetName(), 0)
		ZendHashDestroy(ce.GetFunctionTable())
		if ce.GetConstantsTable().GetNNumOfElements() {
			var c *ZendClassConstant
			var __ht *HashTable = ce.GetConstantsTable()
			for _, _p := range __ht.foreachData() {
				var _z *Zval = _p.GetVal()

				c = _z.GetPtr()
				if c.GetCe() == ce {
					ZvalPtrDtorNogc(c.GetValue())
					if c.GetDocComment() != nil {
						ZendStringReleaseEx(c.GetDocComment(), 0)
					}
				}
			}
		}
		ZendHashDestroy(ce.GetConstantsTable())
		if ce.GetNumInterfaces() > 0 {
			if !ce.IsResolvedInterfaces() {
				var i uint32
				for i = 0; i < ce.GetNumInterfaces(); i++ {
					ZendStringReleaseEx(ce.interface_names[i].name, 0)
					ZendStringReleaseEx(ce.interface_names[i].lc_name, 0)
				}
			}
			Efree(ce.interfaces)
		}
		if ce.GetDocComment() != nil {
			ZendStringReleaseEx(ce.GetDocComment(), 0)
		}
		if ce.GetNumTraits() > 0 {
			_destroyZendClassTraitsInfo(ce)
		}
		break
	case ZEND_INTERNAL_CLASS:
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
			if ce.GetStaticMembersTablePtr() != ce.GetDefaultStaticMembersTable() {
				ZendCleanupInternalClassData(ce)
			}
		}
		ZendHashDestroy(ce.GetPropertiesInfo())
		ZendStringReleaseEx(ce.GetName(), 1)

		/* TODO: eliminate this loop for classes without functions with arg_info */

		var __ht *HashTable = ce.GetFunctionTable()
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			fn = _z.GetPtr()
			if fn.HasFnFlags(ZEND_ACC_HAS_RETURN_TYPE|ZEND_ACC_HAS_TYPE_HINTS) && fn.GetScope() == ce {
				ZendFreeInternalArgInfo(fn.GetInternalFunction())
			}
		}
		ZendHashDestroy(ce.GetFunctionTable())
		if ce.GetConstantsTable().GetNNumOfElements() {
			var c *ZendClassConstant
			var __ht *HashTable = ce.GetConstantsTable()
			for _, _p := range __ht.foreachData() {
				var _z *Zval = _p.GetVal()

				c = _z.GetPtr()
				if c.GetCe() == ce {
					ZvalInternalPtrDtor(c.GetValue())
					if c.GetDocComment() != nil {
						ZendStringReleaseEx(c.GetDocComment(), 1)
					}
				}
				Free(c)
			}
			ZendHashDestroy(ce.GetConstantsTable())
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
	var ce *ZendClassEntry = zv.GetPtr()
	if !ce.IsImmutable() {
		ce.GetRefcount()++
	}
}
func DestroyOpArray(op_array *ZendOpArray) {
	var i uint32
	if op_array.GetStaticVariables() != nil {
		var ht *HashTable = ZEND_MAP_PTR_GET(op_array.static_variables_ptr)
		if ht != nil && (ht.GetGcFlags()&IS_ARRAY_IMMUTABLE) == 0 {
			if ht.DelRefcount() == 0 {
				ZendArrayDestroy(ht)
			}
		}
	}
	if op_array.IsHeapRtCache() && op_array.GetRunTimeCachePtr() != nil {
		Efree(op_array.GetRunTimeCachePtr())
	}
	if op_array.GetRefcount() == nil || b.PreDec(&(op_array.refcount)) > 0 {
		return
	}
	EfreeSize(op_array.GetRefcount(), b.SizeOf("* ( op_array -> refcount )"))
	if op_array.GetVars() != nil {
		i = op_array.GetLastVar()
		for i > 0 {
			i--
			ZendStringReleaseEx(op_array.GetVars()[i], 0)
		}
		Efree(op_array.GetVars())
	}
	if op_array.GetLiterals() != nil {
		var literal *Zval = op_array.GetLiterals()
		var end *Zval = literal + op_array.GetLastLiteral()
		for literal < end {
			ZvalPtrDtorNogc(literal)
			literal++
		}
		if ZEND_USE_ABS_CONST_ADDR || !op_array.IsDonePassTwo() {
			Efree(op_array.GetLiterals())
		}
	}
	Efree(op_array.GetOpcodes())
	if op_array.GetFunctionName() != nil {
		ZendStringReleaseEx(op_array.GetFunctionName(), 0)
	}
	if op_array.GetDocComment() != nil {
		ZendStringReleaseEx(op_array.GetDocComment(), 0)
	}
	if op_array.GetLiveRange() != nil {
		Efree(op_array.GetLiveRange())
	}
	if op_array.GetTryCatchArray() != nil {
		Efree(op_array.GetTryCatchArray())
	}
	if (ZendExtensionFlags & ZEND_EXTENSIONS_HAVE_OP_ARRAY_DTOR) != 0 {
		if op_array.IsDonePassTwo() {
			ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionOpArrayDtorHandler), op_array)
		}
	}
	if op_array.GetArgInfo() != nil {
		var num_args uint32 = op_array.GetNumArgs()
		var arg_info *ZendArgInfo = op_array.GetArgInfo()
		if op_array.IsHasReturnType() {
			arg_info--
			num_args++
		}
		if op_array.IsVariadic() {
			num_args++
		}
		for i = 0; i < num_args; i++ {
			if arg_info[i].GetName() != nil {
				ZendStringReleaseEx(arg_info[i].GetName(), 0)
			}
			if arg_info[i].GetType().IsClass() {
				ZendStringReleaseEx(arg_info[i].GetType().Name(), 0)
			}
		}
		Efree(arg_info)
	}
}
func ZendUpdateExtendedStmts(op_array *ZendOpArray) {
	var opline *ZendOp = op_array.GetOpcodes()
	var end *ZendOp = opline + op_array.GetLast()
	for opline < end {
		if opline.GetOpcode() == ZEND_EXT_STMT {
			if opline+1 < end {
				if (opline + 1).GetOpcode() == ZEND_EXT_STMT {
					opline.SetOpcode(ZEND_NOP)
					opline++
					continue
				}
				if opline+1 < end {
					opline.SetLineno((opline + 1).GetLineno())
				}
			} else {
				opline.SetOpcode(ZEND_NOP)
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
			__CG().SetInCompilation(1)
			__CG().SetActiveOpArray(op_array)
			__CG().SetZendLineno(op_array.GetOpcodes()[op_num].GetLineno())
			ZendErrorNoreturn(E_COMPILE_ERROR, "jump into a finally block is disallowed")
		} else if op_num >= op_array.GetTryCatchArray()[i].GetFinallyOp() && op_num <= op_array.GetTryCatchArray()[i].GetFinallyEnd() && (dst_num > op_array.GetTryCatchArray()[i].GetFinallyEnd() || dst_num < op_array.GetTryCatchArray()[i].GetFinallyOp()) {
			__CG().SetInCompilation(1)
			__CG().SetActiveOpArray(op_array)
			__CG().SetZendLineno(op_array.GetOpcodes()[op_num].GetLineno())
			ZendErrorNoreturn(E_COMPILE_ERROR, "jump out of a finally block is disallowed")
		}
	}
}
func ZendGetBrkContTarget(op_array *ZendOpArray, opline *ZendOp) uint32 {
	var nest_levels int = opline.GetOp2().GetNum()
	var array_offset int = opline.GetOp1().GetNum()
	var jmp_to *ZendBrkContElement
	for {
		jmp_to = __CG().GetContext().GetBrkContArray()[array_offset]
		if nest_levels > 1 {
			array_offset = jmp_to.GetParent()
		}
		if b.PreDec(&nest_levels) <= 0 {
			break
		}
	}
	if opline.GetOpcode() == ZEND_BRK {
		return jmp_to.GetBrk()
	} else {
		return jmp_to.GetCont()
	}
}
func EmitLiveRangeRaw(op_array *ZendOpArray, var_num uint32, kind uint32, start uint32, end uint32) {
	var range_ *ZendLiveRange
	op_array.GetLastLiveRange()++
	op_array.SetLiveRange(Erealloc(op_array.GetLiveRange(), b.SizeOf("zend_live_range")*op_array.GetLastLiveRange()))
	ZEND_ASSERT(start < end)
	range_ = op_array.GetLiveRange()[op_array.GetLastLiveRange()-1]
	range_.SetVar(uint32(intPtr(ZEND_CALL_VAR_NUM(nil, op_array.GetLastVar()+var_num))))
	range_.SetVar(range_.GetVar() | kind)
	range_.SetStart(start)
	range_.SetEnd(end)
}
func EmitLiveRange(op_array *ZendOpArray, var_num uint32, start uint32, end uint32, needs_live_range ZendNeedsLiveRangeCb) {
	var def_opline *ZendOp = op_array.GetOpcodes()[start]
	var orig_def_opline *ZendOp = def_opline
	var use_opline *ZendOp = op_array.GetOpcodes()[end]
	var kind uint32
	switch def_opline.GetOpcode() {
	case ZEND_ADD_ARRAY_ELEMENT:

	case ZEND_ADD_ARRAY_UNPACK:

	case ZEND_ROPE_ADD:
		ZEND_ASSERT(false)
		return
	case ZEND_JMPZ_EX:

	case ZEND_JMPNZ_EX:

	case ZEND_BOOL:

	case ZEND_BOOL_NOT:

	case ZEND_FETCH_CLASS:

	case ZEND_DECLARE_ANON_CLASS:

	case ZEND_FAST_CALL:
		return
	case ZEND_BEGIN_SILENCE:
		kind = ZEND_LIVE_SILENCE
		start++
		break
	case ZEND_ROPE_INIT:
		kind = ZEND_LIVE_ROPE

		/* ROPE live ranges include the generating opcode. */

		def_opline--
		break
	case ZEND_FE_RESET_R:

	case ZEND_FE_RESET_RW:
		kind = ZEND_LIVE_LOOP
		start++
		break
	case ZEND_NEW:
		var level int = 0
		var orig_start uint32 = start
		for def_opline+1 < use_opline {
			def_opline++
			start++
			if def_opline.GetOpcode() == ZEND_DO_FCALL {
				if level == 0 {
					break
				}
				level--
			} else {
				switch def_opline.GetOpcode() {
				case ZEND_INIT_FCALL:

				case ZEND_INIT_FCALL_BY_NAME:

				case ZEND_INIT_NS_FCALL_BY_NAME:

				case ZEND_INIT_DYNAMIC_CALL:

				case ZEND_INIT_USER_CALL:

				case ZEND_INIT_METHOD_CALL:

				case ZEND_INIT_STATIC_METHOD_CALL:

				case ZEND_NEW:
					level++
					break
				case ZEND_DO_ICALL:

				case ZEND_DO_UCALL:

				case ZEND_DO_FCALL_BY_NAME:
					level--
					break
				}
			}
		}
		EmitLiveRangeRaw(op_array, var_num, ZEND_LIVE_NEW, orig_start+1, start+1)
		if start+1 == end {

			/* Trivial live-range, no need to store it. */

			return

			/* Trivial live-range, no need to store it. */

		}
	default:
		start++
		kind = ZEND_LIVE_TMPVAR

		/* Check hook to determine whether a live range is necessary,
		 * e.g. based on type info. */

		if needs_live_range != nil && needs_live_range(op_array, orig_def_opline) == 0 {
			return
		}
		break
	case ZEND_COPY_TMP:

		/* COPY_TMP has a split live-range: One from the definition until the use in
		 * "null" branch, and another from the start of the "non-null" branch to the
		 * FREE opcode. */

		var rt_var_num uint32 = uint32(intPtr(ZEND_CALL_VAR_NUM(nil, op_array.GetLastVar()+var_num)))
		var block_start_op *ZendOp = use_opline
		if needs_live_range != nil && needs_live_range(op_array, orig_def_opline) == 0 {
			return
		}
		for (block_start_op - 1).opcode == ZEND_FREE {
			block_start_op--
		}
		kind = ZEND_LIVE_TMPVAR
		start = block_start_op - op_array.GetOpcodes()
		if start != end {
			EmitLiveRangeRaw(op_array, var_num, kind, start, end)
		}
		for {
			use_opline--
			if (use_opline.GetOp1Type()&(IS_TMP_VAR|IS_VAR)) != 0 && use_opline.GetOp1().GetVar() == rt_var_num || (use_opline.GetOp2Type()&(IS_TMP_VAR|IS_VAR)) != 0 && use_opline.GetOp2().GetVar() == rt_var_num {
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

	return opline.GetOpcode() == ZEND_ROPE_ADD || opline.GetOpcode() == ZEND_ADD_ARRAY_ELEMENT || opline.GetOpcode() == ZEND_ADD_ARRAY_UNPACK

	/* These opcodes only modify the result, not create it. */
}
func KeepsOp1Alive(opline *ZendOp) ZendBool {
	/* These opcodes don't consume their OP1 operand,
	 * it is later freed by something else. */

	if opline.GetOpcode() == ZEND_CASE || opline.GetOpcode() == ZEND_SWITCH_LONG || opline.GetOpcode() == ZEND_FETCH_LIST_R || opline.GetOpcode() == ZEND_COPY_TMP {
		return 1
	}
	ZEND_ASSERT(opline.GetOpcode() != ZEND_SWITCH_STRING && opline.GetOpcode() != ZEND_FE_FETCH_R && opline.GetOpcode() != ZEND_FE_FETCH_RW && opline.GetOpcode() != ZEND_FETCH_LIST_W && opline.GetOpcode() != ZEND_VERIFY_RETURN_TYPE && opline.GetOpcode() != ZEND_BIND_LEXICAL && opline.GetOpcode() != ZEND_ROPE_ADD)
	return 0
}
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
	var opline *ZendOp = op_array.GetOpcodes()[opnum]
	var var_offset uint32 = op_array.GetLastVar()
	var last_use *uint32 = DoAlloca(b.SizeOf("uint32_t")*op_array.GetT(), use_heap)
	memset(last_use, -1, b.SizeOf("uint32_t")*op_array.GetT())
	ZEND_ASSERT(op_array.GetLiveRange() == nil)
	for opnum > 0 {
		opnum--
		opline--
		if (opline.GetResultType()&(IS_TMP_VAR|IS_VAR)) != 0 && IsFakeDef(opline) == 0 {
			var var_num uint32 = EX_VAR_TO_NUM(opline.GetResult().GetVar()) - var_offset

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

					ZEND_ASSERT(opline.GetOpcode() != ZEND_OP_DATA)
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
		if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
			var var_num uint32 = EX_VAR_TO_NUM(opline.GetOp1().GetVar()) - var_offset
			if last_use[var_num] == uint32-1 {
				if KeepsOp1Alive(opline) == 0 {

					/* OP_DATA is really part of the previous opcode. */

					last_use[var_num] = opnum - (opline.GetOpcode() == ZEND_OP_DATA)

					/* OP_DATA is really part of the previous opcode. */

				}
			}
		}
		if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
			var var_num uint32 = EX_VAR_TO_NUM(opline.GetOp2().GetVar()) - var_offset
			if opline.GetOpcode() == ZEND_FE_FETCH_R || opline.GetOpcode() == ZEND_FE_FETCH_RW {

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

				ZEND_ASSERT(opline.GetOpcode() != ZEND_OP_DATA)
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
				ZendSort(r1, r2-r1+1, b.SizeOf("zend_live_range"), CompareFuncT(CmpLiveRange), SwapFuncT(SwapLiveRange))
				break
			}
			r1++
		}
	}
	FreeAlloca(last_use, use_heap)
}
func ZendRecalcLiveRanges(op_array *ZendOpArray, needs_live_range ZendNeedsLiveRangeCb) {
	/* We assume that we never create live-ranges where there were none before. */

	ZEND_ASSERT(op_array.GetLiveRange() != nil)
	Efree(op_array.GetLiveRange())
	op_array.SetLiveRange(nil)
	op_array.SetLastLiveRange(0)
	ZendCalcLiveRanges(op_array, needs_live_range)
}
func PassTwo(op_array *ZendOpArray) int {
	var opline *ZendOp
	var end *ZendOp
	if !(ZEND_USER_CODE(op_array.GetType())) {
		return 0
	}
	if (__CG().GetCompilerOptions() & ZEND_COMPILE_EXTENDED_STMT) != 0 {
		ZendUpdateExtendedStmts(op_array)
	}
	if (__CG().GetCompilerOptions() & ZEND_COMPILE_HANDLE_OP_ARRAY) != 0 {
		if (ZendExtensionFlags & ZEND_EXTENSIONS_HAVE_OP_ARRAY_HANDLER) != 0 {
			ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionOpArrayHandler), op_array)
		}
	}
	if __CG().GetContext().GetVarsSize() != op_array.GetLastVar() {
		op_array.SetVars((**ZendString)(Erealloc(op_array.GetVars(), b.SizeOf("zend_string *")*op_array.GetLastVar())))
		__CG().GetContext().SetVarsSize(op_array.GetLastVar())
	}
	op_array.SetOpcodes((*ZendOp)(Erealloc(op_array.GetOpcodes(), ZEND_MM_ALIGNED_SIZE_EX(b.SizeOf("zend_op")*op_array.GetLast(), 16)+b.SizeOf("zval")*op_array.GetLastLiteral())))
	if op_array.GetLiterals() != nil {
		memcpy((*byte)(op_array.GetOpcodes())+ZEND_MM_ALIGNED_SIZE_EX(b.SizeOf("zend_op")*op_array.GetLast(), 16), op_array.GetLiterals(), b.SizeOf("zval")*op_array.GetLastLiteral())
		Efree(op_array.GetLiterals())
		op_array.SetLiterals((*Zval)((*byte)(op_array.GetOpcodes()) + ZEND_MM_ALIGNED_SIZE_EX(b.SizeOf("zend_op")*op_array.GetLast(), 16)))
	}
	__CG().GetContext().SetOpcodesSize(op_array.GetLast())
	__CG().GetContext().SetLiteralsSize(op_array.GetLastLiteral())

	/* Needs to be set directly after the opcode/literal reallocation, to ensure destruction
	 * happens correctly if any of the following fixups generate a fatal error. */

	op_array.SetIsDonePassTwo(true)
	opline = op_array.GetOpcodes()
	end = opline + op_array.GetLast()
	for opline < end {
		switch opline.GetOpcode() {
		case ZEND_RECV_INIT:
			var val *Zval = CT_CONSTANT(opline.GetOp2())
			if val.IsType(IS_CONSTANT_AST) {
				var slot uint32 = ZEND_MM_ALIGNED_SIZE_EX(op_array.GetCacheSize(), 8)
				val.SetCacheSlot(slot)
				op_array.SetCacheSize(op_array.GetCacheSize() + b.SizeOf("zval"))
			}
			break
		case ZEND_FAST_CALL:
			opline.GetOp1().SetOplineNum(op_array.GetTryCatchArray()[opline.GetOp1().GetNum()].GetFinallyOp())
			ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp1())
			break
		case ZEND_BRK:

		case ZEND_CONT:
			var jmp_target uint32 = ZendGetBrkContTarget(op_array, opline)
			if op_array.IsHasFinallyBlock() {
				ZendCheckFinallyBreakout(op_array, opline-op_array.GetOpcodes(), jmp_target)
			}
			opline.SetOpcode(ZEND_JMP)
			opline.GetOp1().SetOplineNum(jmp_target)
			opline.GetOp2().SetNum(0)
			ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp1())
			break
		case ZEND_GOTO:
			ZendResolveGotoLabel(op_array, opline)
			if op_array.IsHasFinallyBlock() {
				ZendCheckFinallyBreakout(op_array, opline-op_array.GetOpcodes(), opline.GetOp1().GetOplineNum())
			}
		case ZEND_JMP:
			ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp1())
			break
		case ZEND_JMPZNZ:

			/* absolute index to relative offset */

			opline.SetExtendedValue(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, opline.GetExtendedValue()))
		case ZEND_JMPZ:

		case ZEND_JMPNZ:

		case ZEND_JMPZ_EX:

		case ZEND_JMPNZ_EX:

		case ZEND_JMP_SET:

		case ZEND_COALESCE:

		case ZEND_FE_RESET_R:

		case ZEND_FE_RESET_RW:
			ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp2())
			break
		case ZEND_ASSERT_CHECK:

			/* If result of assert is unused, result of check is unused as well */

			var call *ZendOp = op_array.GetOpcodes()[opline.GetOp2().GetOplineNum()-1]
			if call.GetOpcode() == ZEND_EXT_FCALL_END {
				call--
			}
			if call.GetResultType() == IS_UNUSED {
				opline.SetResultType(IS_UNUSED)
			}
			ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp2())
			break
		case ZEND_FE_FETCH_R:

		case ZEND_FE_FETCH_RW:

			/* absolute index to relative offset */

			opline.SetExtendedValue(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, opline.GetExtendedValue()))
			break
		case ZEND_CATCH:
			if (opline.GetExtendedValue() & ZEND_LAST_CATCH) == 0 {
				ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp2())
			}
			break
		case ZEND_RETURN:

		case ZEND_RETURN_BY_REF:
			if op_array.IsGenerator() {
				opline.SetOpcode(ZEND_GENERATOR_RETURN)
			}
			break
		case ZEND_SWITCH_LONG:

		case ZEND_SWITCH_STRING:

			/* absolute indexes to relative offsets */

			var jumptable *HashTable = CT_CONSTANT(opline.GetOp2()).GetArr()
			var zv *Zval
			var __ht *HashTable = jumptable
			for _, _p := range __ht.foreachData() {
				var _z *Zval = _p.GetVal()

				zv = _z
				zv.SetLval(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, zv.GetLval()))
			}
			opline.SetExtendedValue(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, opline.GetExtendedValue()))
			break
		}
		if opline.GetOp1Type() == IS_CONST {
			ZEND_PASS_TWO_UPDATE_CONSTANT(op_array, opline, opline.GetOp1())
		} else if (opline.GetOp1Type() & (IS_VAR | IS_TMP_VAR)) != 0 {
			opline.GetOp1().SetVar(uint32(ZendIntptrT(ZEND_CALL_VAR_NUM(nil, op_array.GetLastVar()+opline.GetOp1().GetVar()))))
		}
		if opline.GetOp2Type() == IS_CONST {
			ZEND_PASS_TWO_UPDATE_CONSTANT(op_array, opline, opline.GetOp2())
		} else if (opline.GetOp2Type() & (IS_VAR | IS_TMP_VAR)) != 0 {
			opline.GetOp2().SetVar(uint32(ZendIntptrT(ZEND_CALL_VAR_NUM(nil, op_array.GetLastVar()+opline.GetOp2().GetVar()))))
		}
		if (opline.GetResultType() & (IS_VAR | IS_TMP_VAR)) != 0 {
			opline.GetResult().SetVar(uint32(ZendIntptrT(ZEND_CALL_VAR_NUM(nil, op_array.GetLastVar()+opline.GetResult().GetVar()))))
		}
		ZEND_VM_SET_OPCODE_HANDLER(opline)
		opline++
	}
	ZendCalcLiveRanges(op_array, nil)
	return 0
}
func GetUnaryOp(opcode int) UnaryOpType {
	switch opcode {
	case ZEND_BW_NOT:
		return UnaryOpType(BitwiseNotFunction)
	case ZEND_BOOL_NOT:
		return UnaryOpType(BooleanNotFunction)
	default:
		return UnaryOpType(nil)
	}
}
func GetBinaryOp(opcode int) BinaryOpType {
	switch opcode {
	case ZEND_ADD:
		return BinaryOpType(AddFunction)
	case ZEND_SUB:
		return BinaryOpType(SubFunction)
	case ZEND_MUL:
		return BinaryOpType(MulFunction)
	case ZEND_POW:
		return BinaryOpType(PowFunction)
	case ZEND_DIV:
		return BinaryOpType(DivFunction)
	case ZEND_MOD:
		return BinaryOpType(ModFunction)
	case ZEND_SL:
		return BinaryOpType(ShiftLeftFunction)
	case ZEND_SR:
		return BinaryOpType(ShiftRightFunction)
	case ZEND_PARENTHESIZED_CONCAT:

	case ZEND_FAST_CONCAT:

	case ZEND_CONCAT:
		return BinaryOpType(ConcatFunction)
	case ZEND_IS_IDENTICAL:
		return BinaryOpType(IsIdenticalFunction)
	case ZEND_IS_NOT_IDENTICAL:
		return BinaryOpType(IsNotIdenticalFunction)
	case ZEND_IS_EQUAL:

	case ZEND_CASE:
		return BinaryOpType(IsEqualFunction)
	case ZEND_IS_NOT_EQUAL:
		return BinaryOpType(IsNotEqualFunction)
	case ZEND_IS_SMALLER:
		return BinaryOpType(IsSmallerFunction)
	case ZEND_IS_SMALLER_OR_EQUAL:
		return BinaryOpType(IsSmallerOrEqualFunction)
	case ZEND_SPACESHIP:
		return BinaryOpType(CompareFunction)
	case ZEND_BW_OR:
		return BinaryOpType(BitwiseOrFunction)
	case ZEND_BW_AND:
		return BinaryOpType(BitwiseAndFunction)
	case ZEND_BW_XOR:
		return BinaryOpType(BitwiseXorFunction)
	case ZEND_BOOL_XOR:
		return BinaryOpType(BooleanXorFunction)
	default:
		ZEND_ASSERT(false)
		return BinaryOpType(nil)
	}
}
