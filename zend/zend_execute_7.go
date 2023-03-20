// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func IZvalPtrDtorNoref(zval_ptr *types.Zval) {
	if zval_ptr.IsRefcounted() {
		var ref *types.ZendRefcounted = zval_ptr.GetCounted()
		b.Assert(zval_ptr.GetType() != types.IS_REFERENCE)
		if ref.DelRefcount() == 0 {
			RcDtorFunc(ref)
		} else if GC_MAY_LEAK(ref) {
			GcPossibleRoot(ref)
		}
	}
}
func ZendAssignToTypedRef(variable_ptr *types.Zval, orig_value *types.Zval, value_type types.ZendUchar, strict types.ZendBool, ref *types.ZendRefcounted) *types.Zval {
	var ret types.ZendBool
	var value types.Zval
	types.ZVAL_COPY(&value, orig_value)
	ret = ZendVerifyRefAssignableZval(variable_ptr.GetRef(), &value, strict)
	variable_ptr = types.Z_REFVAL_P(variable_ptr)
	if ret != 0 {
		IZvalPtrDtorNoref(variable_ptr)
		types.ZVAL_COPY_VALUE(variable_ptr, &value)
	} else {
		ZvalPtrDtorNogc(&value)
	}
	if (value_type & (IS_VAR | IS_TMP_VAR)) != 0 {
		if ref != nil {
			if ref.DelRefcount() == 0 {
				ZvalPtrDtor(orig_value)
				EfreeSize(ref, b.SizeOf("zend_reference"))
			}
		} else {
			IZvalPtrDtorNoref(orig_value)
		}
	}
	return variable_ptr
}
func ZendVerifyPropAssignableByRef(prop_info *ZendPropertyInfo, orig_val *types.Zval, strict types.ZendBool) types.ZendBool {
	var val *types.Zval = orig_val
	if val.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(val.GetRef()) {
		var result int
		val = types.Z_REFVAL_P(val)
		result = IZendVerifyTypeAssignableZval(prop_info.GetType(), prop_info.GetCe(), val, strict)
		if result > 0 {
			return 1
		}
		if result < 0 {
			var ref_prop *ZendPropertyInfo = ZEND_REF_FIRST_SOURCE(orig_val.GetRef())
			if prop_info.GetType().Code() != ref_prop.GetType().Code() {

				/* Invalid due to conflicting coercion */

				ZendThrowRefTypeErrorType(ref_prop, prop_info, val)
				return 0
			}
			if ZendVerifyWeakScalarTypeHint(prop_info.GetType().Code(), val) != 0 {
				return 1
			}
		}
	} else {
		val = types.ZVAL_DEREF(val)
		if IZendCheckPropertyType(prop_info, val, strict) != 0 {
			return 1
		}
	}
	ZendVerifyPropertyTypeError(prop_info, val)
	return 0
}
func ZendRefAddTypeSource(source_list *types.ZendPropertyInfoSourceList, prop *ZendPropertyInfo) {
	var list *types.ZendPropertyInfoList
	if source_list.GetPtr() == nil {
		source_list.SetPtr(prop)
		return
	}
	list = types.ZEND_PROPERTY_INFO_SOURCE_TO_LIST(source_list.GetList())
	if types.ZEND_PROPERTY_INFO_SOURCE_IS_LIST(source_list.GetList()) == 0 {
		list = Emalloc(b.SizeOf("zend_property_info_list") + (4-1)*b.SizeOf("zend_property_info *"))
		list.GetPtr()[0] = source_list.GetPtr()
		list.SetNumAllocated(4)
		list.SetNum(1)
	} else if list.GetNumAllocated() == list.GetNum() {
		list.SetNumAllocated(list.GetNum() * 2)
		list = Erealloc(list, b.SizeOf("zend_property_info_list")+(list.GetNumAllocated()-1)*b.SizeOf("zend_property_info *"))
	}
	list.GetPtr()[b.PostInc(&(list.GetNum()))] = prop
	source_list.SetList(types.ZEND_PROPERTY_INFO_SOURCE_FROM_LIST(list))
}
func ZendRefDelTypeSource(source_list *types.ZendPropertyInfoSourceList, prop *ZendPropertyInfo) {
	var list *types.ZendPropertyInfoList = types.ZEND_PROPERTY_INFO_SOURCE_TO_LIST(source_list.GetList())
	var ptr **ZendPropertyInfo
	var end ***ZendPropertyInfo
	if types.ZEND_PROPERTY_INFO_SOURCE_IS_LIST(source_list.GetList()) == 0 {
		b.Assert(source_list.GetPtr() == prop)
		source_list.SetPtr(nil)
		return
	}
	if list.GetNum() == 1 {
		b.Assert(list.ptr == prop)
		Efree(list)
		source_list.SetPtr(nil)
		return
	}

	/* Checking against end here to get a more graceful failure mode if we missed adding a type
	 * source at some point. */

	ptr = list.GetPtr()
	end = ptr + list.GetNum()
	for ptr < end && (*ptr) != prop {
		ptr++
	}
	b.Assert((*ptr) == prop)

	/* Copy the last list element into the deleted slot. */

	*ptr = list.GetPtr()[b.PreDec(&(list.GetNum()))]
	if list.GetNum() >= 4 && list.GetNum()*4 == list.GetNumAllocated() {
		list.SetNumAllocated(list.GetNum() * 2)
		source_list.SetList(types.ZEND_PROPERTY_INFO_SOURCE_FROM_LIST(Erealloc(list, b.SizeOf("zend_property_info_list")+(list.GetNumAllocated()-1)*b.SizeOf("zend_property_info *"))))
	}
}
func ZendFetchThisVar(type_ int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *types.Zval = EX_VAR(opline.GetResult().GetVar())
	switch type_ {
	case BP_VAR_R:
		if executeData.GetThis().u1.v.type_ == types.IS_OBJECT {
			result.SetObject(executeData.GetThis().GetObj())
			result.AddRefcount()
		} else {
			result.SetNull()
			faults.Error(faults.E_NOTICE, "Undefined variable: this")
		}
	case BP_VAR_IS:
		if executeData.GetThis().u1.v.type_ == types.IS_OBJECT {
			result.SetObject(executeData.GetThis().GetObj())
			result.AddRefcount()
		} else {
			result.SetNull()
		}
	case BP_VAR_RW:
		fallthrough
	case BP_VAR_W:
		result.SetUndef()
		faults.ThrowError(nil, "Cannot re-assign $this")
	case BP_VAR_UNSET:
		result.SetUndef()
		faults.ThrowError(nil, "Cannot unset $this")
	default:

	}
}
func ZendWrongCloneCall(clone *ZendFunction, scope *types.ClassEntry) {
	faults.ThrowError(nil, "Call to %s %s::__clone() from context '%s'", ZendVisibilityString(clone.GetFnFlags()), clone.GetScope().GetName().GetVal(), b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
}
func ExecuteInternal(executeData *ZendExecuteData, return_value *types.Zval) {
	executeData.GetFunc().GetInternalFunction().GetHandler()(executeData, return_value)
}
func ZendCleanAndCacheSymbolTable(symbol_table *types.ZendArray) {
	/* Clean before putting into the cache, since clean could call dtors,
	 * which could use the cached hash. Also do this before the check for
	 * available cache slots, as those may be used by a dtor as well. */

	symbol_table.SymtableClean()
	if EG__().GetSymtableCachePtr() >= EG__().GetSymtableCacheLimit() {
		symbol_table.DestroyEx()
	} else {
		*(b.PostInc(&(EG__().GetSymtableCachePtr()))) = symbol_table
	}
}
func IFreeCompiledVariables(executeData *ZendExecuteData) {
	var cv *types.Zval = executeData.VarNum(0)
	var count int = executeData.GetFunc().op_array.last_var
	for count != 0 {
		if cv.IsRefcounted() {
			var r *types.ZendRefcounted = cv.GetCounted()
			if r.DelRefcount() == 0 {
				cv.SetNull()
				RcDtorFunc(r)
			} else {
				GcCheckPossibleRoot(r)
			}
		}
		cv++
		count--
	}
}
func ZendFreeCompiledVariables(executeData *ZendExecuteData) { IFreeCompiledVariables(executeData) }
func ZEND_VM_INTERRUPT_CHECK() {
	if EG__().GetVmInterrupt() != 0 {
		ZEND_VM_INTERRUPT()
	}
}
func ZEND_VM_LOOP_INTERRUPT_CHECK() {
	if EG__().GetVmInterrupt() != 0 {
		ZEND_VM_LOOP_INTERRUPT()
	}
}
func ZendCopyExtraArgs(executeData *ZendExecuteData) {
	var op_array *ZendOpArray = executeData.GetFunc().op_array
	var first_extra_arg uint32 = op_array.GetNumArgs()
	var num_args uint32 = executeData.NumArgs()
	var src *types.Zval
	var delta int
	var count uint32
	var type_flags uint32 = 0
	if !op_array.IsHasTypeHints() {
		executeData.GetOpline() += first_extra_arg

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

	}

	/* move extra args into separate array after all CV and TMP vars */

	src = executeData.VarNum(num_args - 1)
	delta = op_array.GetLastVar() + op_array.GetT() - first_extra_arg
	count = num_args - first_extra_arg
	if delta != 0 {
		delta *= b.SizeOf("zval")
		for {
			type_flags |= src.GetTypeInfo()
			types.ZVAL_COPY_VALUE((*types.Zval)((*byte)(src)+delta), src)
			src.SetUndef()
			src--
			if !(b.PreDec(&count)) {
				break
			}
		}
		if types.Z_TYPE_INFO_REFCOUNTED(type_flags) {
			ZEND_ADD_CALL_FLAG(executeData, ZEND_CALL_FREE_EXTRA_ARGS)
		}
	} else {
		for {
			if src.IsRefcounted() {
				ZEND_ADD_CALL_FLAG(executeData, ZEND_CALL_FREE_EXTRA_ARGS)
				break
			}
			src--
			if !(b.PreDec(&count)) {
				break
			}
		}
	}
}
func ZendInitCvs(first uint32, last uint32, executeData *ZendExecuteData) {
	if first < last {
		var count uint32 = last - first
		var var_ *types.Zval = executeData.VarNum(first)
		for {
			var_.SetUndef()
			var_++
			if !(b.PreDec(&count)) {
				break
			}
		}
	}
}
func IInitFuncExecuteData(op_array *ZendOpArray, return_value *types.Zval, may_be_trampoline types.ZendBool, executeData *ZendExecuteData) {
	var first_extra_arg uint32
	var num_args uint32
	b.Assert(executeData.GetFunc() == (*ZendFunction)(op_array))
	executeData.GetOpline() = op_array.GetOpcodes()
	executeData.GetCall() = nil
	executeData.GetReturnValue() = return_value

	/* Handle arguments */

	first_extra_arg = op_array.GetNumArgs()
	num_args = executeData.NumArgs()
	if num_args > first_extra_arg {
		if may_be_trampoline == 0 || !op_array.IsCallViaTrampoline() {
			ZendCopyExtraArgs(executeData)
		}
	} else if !op_array.IsHasTypeHints() {
		executeData.GetOpline() += num_args

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

	}

	/* Initialize CV variables (skip arguments) */

	ZendInitCvs(num_args, op_array.GetLastVar(), executeData)
	executeData.GetRunTimeCache() = RUN_TIME_CACHE(op_array)
	EG__().SetCurrentExecuteData(executeData)
}
func InitFuncRunTimeCacheI(op_array *ZendOpArray) {
	var run_time_cache *any
	b.Assert(RUN_TIME_CACHE(op_array) == nil)
	run_time_cache = ZendArenaAlloc(CG__().GetArena(), op_array.GetCacheSize())
	memset(run_time_cache, 0, op_array.GetCacheSize())
	ZEND_MAP_PTR_SET(op_array.run_time_cache, run_time_cache)
}
func InitFuncRunTimeCache(op_array *ZendOpArray) { InitFuncRunTimeCacheI(op_array) }
func ZendFetchFunction(name *types.String) *ZendFunction {
	var zv *types.Zval = EG__().GetFunctionTable().KeyFind(name.GetStr())
	if zv != nil {
		var fbc *ZendFunction = zv.GetFunc()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCacheI(fbc.GetOpArray())
		}
		return fbc
	}
	return nil
}
func ZendFetchFunctionStr(name string, len_ int) *ZendFunction {
	var zv *types.Zval = EG__().GetFunctionTable().KeyFind(b.CastStr(name, len_))
	if zv != nil {
		var fbc *ZendFunction = zv.GetFunc()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCacheI(fbc.GetOpArray())
		}
		return fbc
	}
	return nil
}
func ZendInitFuncRunTimeCache(op_array *ZendOpArray) {
	if !(RUN_TIME_CACHE(op_array)) {
		InitFuncRunTimeCacheI(op_array)
	}
}
func IInitCodeExecuteData(executeData *ZendExecuteData, op_array *ZendOpArray, return_value *types.Zval) {
	b.Assert(executeData.GetFunc() == (*ZendFunction)(op_array))
	executeData.GetOpline() = op_array.GetOpcodes()
	executeData.GetCall() = nil
	executeData.GetReturnValue() = return_value
	ZendAttachSymbolTable(executeData)
	if op_array.GetRunTimeCachePtr() == nil {
		var ptr any
		b.Assert(op_array.IsHeapRtCache())
		ptr = Emalloc(op_array.GetCacheSize() + b.SizeOf("void *"))
		ZEND_MAP_PTR_INIT(op_array.run_time_cache, ptr)
		ptr = (*byte)(ptr + b.SizeOf("void *"))
		ZEND_MAP_PTR_SET(op_array.run_time_cache, ptr)
		memset(ptr, 0, op_array.GetCacheSize())
	}
	executeData.GetRunTimeCache() = RUN_TIME_CACHE(op_array)
	EG__().SetCurrentExecuteData(executeData)
}
func ZendInitFuncExecuteData(ex *ZendExecuteData, op_array *ZendOpArray, return_value *types.Zval) {
	var executeData *ZendExecuteData = ex
	executeData.GetPrevExecuteData() = CurrEX()
	if !(RUN_TIME_CACHE(op_array)) {
		InitFuncRunTimeCache(op_array)
	}
	IInitFuncExecuteData(op_array, return_value, 1, executeData)
}
func ZendInitCodeExecuteData(executeData *ZendExecuteData, op_array *ZendOpArray, return_value *types.Zval) {
	executeData.GetPrevExecuteData() = CurrEX()
	IInitCodeExecuteData(executeData, op_array, return_value)
}
func ZendInitExecuteData(executeData *ZendExecuteData, op_array *ZendOpArray, return_value *types.Zval) {
	if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
		ZendInitCodeExecuteData(executeData, op_array, return_value)
	} else {
		ZendInitFuncExecuteData(executeData, op_array, return_value)
	}
}
func ZendVmStackCopyCallFrame(call *ZendExecuteData, passed_args uint32, additional_args uint32) *ZendExecuteData {
	var new_call *ZendExecuteData
	var used_stack int = EG__().GetVmStackTop() - (*types.Zval)(call) + additional_args

	/* copy call frame into new stack segment */

	new_call = ZendVmStackExtend(used_stack * b.SizeOf("zval"))
	*new_call = *call
	ZEND_ADD_CALL_FLAG(new_call, ZEND_CALL_ALLOCATED)
	if passed_args != 0 {
		var src *types.Zval = call.Arg(1)
		var dst *types.Zval = new_call.Arg(1)
		for {
			types.ZVAL_COPY_VALUE(dst, src)
			passed_args--
			src++
			dst++
			if passed_args == 0 {
				break
			}
		}
	}

	/* delete old call_frame from previous stack segment */

	EG__().GetVmStack().GetPrev().SetTop((*types.Zval)(call))

	/* delete previous stack segment if it became empty */

	if EG__().GetVmStack().GetPrev().GetTop() == ZEND_VM_STACK_ELEMENTS(EG__().GetVmStack().GetPrev()) {
		var r ZendVmStack = EG__().GetVmStack().GetPrev()
		EG__().GetVmStack().SetPrev(r.GetPrev())
		Efree(r)
	}
	return new_call
}
