// <<generate>>

package zend

import (
	b "sik/builtin"
)

func IZvalPtrDtorNoref(zval_ptr *Zval) {
	if zval_ptr.IsRefcounted() {
		var ref *ZendRefcounted = zval_ptr.GetCounted()
		ZEND_ASSERT(zval_ptr.GetType() != IS_REFERENCE)
		if ref.DelRefcount() == 0 {
			RcDtorFunc(ref)
		} else if GC_MAY_LEAK(ref) {
			GcPossibleRoot(ref)
		}
	}
}
func ZendAssignToTypedRef(variable_ptr *Zval, orig_value *Zval, value_type ZendUchar, strict ZendBool, ref *ZendRefcounted) *Zval {
	var ret ZendBool
	var value Zval
	ZVAL_COPY(&value, orig_value)
	ret = ZendVerifyRefAssignableZval(variable_ptr.GetRef(), &value, strict)
	variable_ptr = Z_REFVAL_P(variable_ptr)
	if ret != 0 {
		IZvalPtrDtorNoref(variable_ptr)
		ZVAL_COPY_VALUE(variable_ptr, &value)
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
func ZendVerifyPropAssignableByRef(prop_info *ZendPropertyInfo, orig_val *Zval, strict ZendBool) ZendBool {
	var val *Zval = orig_val
	if val.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(val.GetRef()) {
		var result int
		val = Z_REFVAL_P(val)
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
		val = ZVAL_DEREF(val)
		if IZendCheckPropertyType(prop_info, val, strict) != 0 {
			return 1
		}
	}
	ZendVerifyPropertyTypeError(prop_info, val)
	return 0
}
func ZendRefAddTypeSource(source_list *ZendPropertyInfoSourceList, prop *ZendPropertyInfo) {
	var list *ZendPropertyInfoList
	if source_list.GetPtr() == nil {
		source_list.SetPtr(prop)
		return
	}
	list = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(source_list.GetList())
	if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(source_list.GetList()) == 0 {
		list = Emalloc(b.SizeOf("zend_property_info_list") + (4-1)*b.SizeOf("zend_property_info *"))
		list.GetPtr()[0] = source_list.GetPtr()
		list.SetNumAllocated(4)
		list.SetNum(1)
	} else if list.GetNumAllocated() == list.GetNum() {
		list.SetNumAllocated(list.GetNum() * 2)
		list = Erealloc(list, b.SizeOf("zend_property_info_list")+(list.GetNumAllocated()-1)*b.SizeOf("zend_property_info *"))
	}
	list.GetPtr()[b.PostInc(&(list.GetNum()))] = prop
	source_list.SetList(ZEND_PROPERTY_INFO_SOURCE_FROM_LIST(list))
}
func ZendRefDelTypeSource(source_list *ZendPropertyInfoSourceList, prop *ZendPropertyInfo) {
	var list *ZendPropertyInfoList = ZEND_PROPERTY_INFO_SOURCE_TO_LIST(source_list.GetList())
	var ptr **ZendPropertyInfo
	var end ***ZendPropertyInfo
	if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(source_list.GetList()) == 0 {
		ZEND_ASSERT(source_list.GetPtr() == prop)
		source_list.SetPtr(nil)
		return
	}
	if list.GetNum() == 1 {
		ZEND_ASSERT(list.ptr == prop)
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
	ZEND_ASSERT((*ptr) == prop)

	/* Copy the last list element into the deleted slot. */

	*ptr = list.GetPtr()[b.PreDec(&(list.GetNum()))]
	if list.GetNum() >= 4 && list.GetNum()*4 == list.GetNumAllocated() {
		list.SetNumAllocated(list.GetNum() * 2)
		source_list.SetList(ZEND_PROPERTY_INFO_SOURCE_FROM_LIST(Erealloc(list, b.SizeOf("zend_property_info_list")+(list.GetNumAllocated()-1)*b.SizeOf("zend_property_info *"))))
	}
}
func ZendFetchThisVar(type_ int, opline *ZendOp, executeData *ZendExecuteData) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	switch type_ {
	case BP_VAR_R:
		if EX(This).u1.v.type_ == IS_OBJECT {
			result.SetObject(EX(This).GetObj())
			result.AddRefcount()
		} else {
			result.SetNull()
			ZendError(E_NOTICE, "Undefined variable: this")
		}
	case BP_VAR_IS:
		if EX(This).u1.v.type_ == IS_OBJECT {
			result.SetObject(EX(This).GetObj())
			result.AddRefcount()
		} else {
			result.SetNull()
		}
	case BP_VAR_RW:
		fallthrough
	case BP_VAR_W:
		result.SetUndef()
		ZendThrowError(nil, "Cannot re-assign $this")
	case BP_VAR_UNSET:
		result.SetUndef()
		ZendThrowError(nil, "Cannot unset $this")
	default:

	}
}
func ZendWrongCloneCall(clone *ZendFunction, scope *ZendClassEntry) {
	ZendThrowError(nil, "Call to %s %s::__clone() from context '%s'", ZendVisibilityString(clone.GetFnFlags()), clone.GetScope().GetName().GetVal(), b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
}
func ExecuteInternal(executeData *ZendExecuteData, return_value *Zval) {
	executeData.GetFunc().GetInternalFunction().GetHandler()(executeData, return_value)
}
func ZendCleanAndCacheSymbolTable(symbol_table *ZendArray) {
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
	var cv *Zval = EX_VAR_NUM(0)
	var count int = EX(func_).op_array.last_var
	for count != 0 {
		if cv.IsRefcounted() {
			var r *ZendRefcounted = cv.GetCounted()
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
	var op_array *ZendOpArray = EX(func_).op_array
	var first_extra_arg uint32 = op_array.GetNumArgs()
	var num_args uint32 = EX_NUM_ARGS()
	var src *Zval
	var delta int
	var count uint32
	var type_flags uint32 = 0
	if !op_array.IsHasTypeHints() {

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

		EX(opline) += first_extra_arg

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

	}

	/* move extra args into separate array after all CV and TMP vars */

	src = EX_VAR_NUM(num_args - 1)
	delta = op_array.GetLastVar() + op_array.GetT() - first_extra_arg
	count = num_args - first_extra_arg
	if delta != 0 {
		delta *= b.SizeOf("zval")
		for {
			type_flags |= src.GetTypeInfo()
			ZVAL_COPY_VALUE((*Zval)((*byte)(src)+delta), src)
			src.SetUndef()
			src--
			if !(b.PreDec(&count)) {
				break
			}
		}
		if Z_TYPE_INFO_REFCOUNTED(type_flags) {
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
		var var_ *Zval = EX_VAR_NUM(first)
		for {
			var_.SetUndef()
			var_++
			if !(b.PreDec(&count)) {
				break
			}
		}
	}
}
func IInitFuncExecuteData(op_array *ZendOpArray, return_value *Zval, may_be_trampoline ZendBool, executeData *ZendExecuteData) {
	var first_extra_arg uint32
	var num_args uint32
	ZEND_ASSERT(EX(func_) == (*ZendFunction)(op_array))
	EX(opline) = op_array.GetOpcodes()
	EX(call) = nil
	EX(return_value) = return_value

	/* Handle arguments */

	first_extra_arg = op_array.GetNumArgs()
	num_args = EX_NUM_ARGS()
	if num_args > first_extra_arg {
		if may_be_trampoline == 0 || !op_array.IsCallViaTrampoline() {
			ZendCopyExtraArgs(executeData)
		}
	} else if !op_array.IsHasTypeHints() {

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

		EX(opline) += num_args

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

	}

	/* Initialize CV variables (skip arguments) */

	ZendInitCvs(num_args, op_array.GetLastVar(), executeData)
	EX(run_time_cache) = RUN_TIME_CACHE(op_array)
	EG__().SetCurrentExecuteData(executeData)
}
func InitFuncRunTimeCacheI(op_array *ZendOpArray) {
	var run_time_cache *any
	ZEND_ASSERT(RUN_TIME_CACHE(op_array) == nil)
	run_time_cache = ZendArenaAlloc(CG__().GetArena(), op_array.GetCacheSize())
	memset(run_time_cache, 0, op_array.GetCacheSize())
	ZEND_MAP_PTR_SET(op_array.run_time_cache, run_time_cache)
}
func InitFuncRunTimeCache(op_array *ZendOpArray) { InitFuncRunTimeCacheI(op_array) }
func ZendFetchFunction(name *ZendString) *ZendFunction {
	var zv *Zval = EG__().GetFunctionTable().KeyFind(name.GetStr())
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
	var zv *Zval = EG__().GetFunctionTable().KeyFind(b.CastStr(name, len_))
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
func IInitCodeExecuteData(executeData *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	ZEND_ASSERT(EX(func_) == (*ZendFunction)(op_array))
	EX(opline) = op_array.GetOpcodes()
	EX(call) = nil
	EX(return_value) = return_value
	ZendAttachSymbolTable(executeData)
	if op_array.GetRunTimeCachePtr() == nil {
		var ptr any
		ZEND_ASSERT(op_array.IsHeapRtCache())
		ptr = Emalloc(op_array.GetCacheSize() + b.SizeOf("void *"))
		ZEND_MAP_PTR_INIT(op_array.run_time_cache, ptr)
		ptr = (*byte)(ptr + b.SizeOf("void *"))
		ZEND_MAP_PTR_SET(op_array.run_time_cache, ptr)
		memset(ptr, 0, op_array.GetCacheSize())
	}
	EX(run_time_cache) = RUN_TIME_CACHE(op_array)
	EG__().SetCurrentExecuteData(executeData)
}
func ZendInitFuncExecuteData(ex *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	var executeData *ZendExecuteData = ex
	EX(prev_execute_data) = CurrEX()
	if !(RUN_TIME_CACHE(op_array)) {
		InitFuncRunTimeCache(op_array)
	}
	IInitFuncExecuteData(op_array, return_value, 1, executeData)
}
func ZendInitCodeExecuteData(executeData *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	EX(prev_execute_data) = CurrEX()
	IInitCodeExecuteData(executeData, op_array, return_value)
}
func ZendInitExecuteData(executeData *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
		ZendInitCodeExecuteData(executeData, op_array, return_value)
	} else {
		ZendInitFuncExecuteData(executeData, op_array, return_value)
	}
}
func ZendVmStackCopyCallFrame(call *ZendExecuteData, passed_args uint32, additional_args uint32) *ZendExecuteData {
	var new_call *ZendExecuteData
	var used_stack int = EG__().GetVmStackTop() - (*Zval)(call) + additional_args

	/* copy call frame into new stack segment */

	new_call = ZendVmStackExtend(used_stack * b.SizeOf("zval"))
	*new_call = *call
	ZEND_ADD_CALL_FLAG(new_call, ZEND_CALL_ALLOCATED)
	if passed_args != 0 {
		var src *Zval = call.Arg(1)
		var dst *Zval = new_call.Arg(1)
		for {
			ZVAL_COPY_VALUE(dst, src)
			passed_args--
			src++
			dst++
			if passed_args == 0 {
				break
			}
		}
	}

	/* delete old call_frame from previous stack segment */

	EG__().GetVmStack().GetPrev().SetTop((*Zval)(call))

	/* delete previous stack segment if it became empty */

	if EG__().GetVmStack().GetPrev().GetTop() == ZEND_VM_STACK_ELEMENTS(EG__().GetVmStack().GetPrev()) {
		var r ZendVmStack = EG__().GetVmStack().GetPrev()
		EG__().GetVmStack().SetPrev(r.GetPrev())
		Efree(r)
	}
	return new_call
}
