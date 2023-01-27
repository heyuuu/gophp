// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func ZEND_REF_TYPE_SOURCES(ref *ZendReference) ZendPropertyInfoSourceList { return ref.GetSources() }
func ZEND_REF_HAS_TYPE_SOURCES(ref *ZendReference) bool {
	return ZEND_REF_TYPE_SOURCES(ref).GetPtr() != nil
}
func ZEND_REF_FIRST_SOURCE(ref *ZendReference) *ZendPropertyInfo {
	if ZEND_PROPERTY_INFO_SOURCE_IS_LIST(ref.GetSources().GetList()) != 0 {
		return ZEND_PROPERTY_INFO_SOURCE_TO_LIST(ref.GetSources().GetList()).GetPtr()[0]
	} else {
		return ref.GetSources().GetPtr()
	}
}
func ZendCopyToVariable(variable_ptr *Zval, value *Zval, value_type ZendUchar, ref *ZendRefcounted) {
	ZVAL_COPY_VALUE(variable_ptr, value)
	if ZEND_CONST_COND(value_type == IS_CONST, 0) {
		if Z_OPT_REFCOUNTED_P(variable_ptr) {
			Z_ADDREF_P(variable_ptr)
		}
	} else if (value_type & (IS_CONST | IS_CV)) != 0 {
		if Z_OPT_REFCOUNTED_P(variable_ptr) {
			Z_ADDREF_P(variable_ptr)
		}
	} else if ZEND_CONST_COND(value_type == IS_VAR, 1) && ref != nil {
		if ref.DecGcRefcount() == 0 {
			EfreeSize(ref, b.SizeOf("zend_reference"))
		} else if Z_OPT_REFCOUNTED_P(variable_ptr) {
			Z_ADDREF_P(variable_ptr)
		}
	}
}
func ZendAssignToVariable(variable_ptr *Zval, value *Zval, value_type ZendUchar, strict ZendBool) *Zval {
	var ref *ZendRefcounted = nil
	if ZEND_CONST_COND(value_type&(IS_VAR|IS_CV), 1) && Z_ISREF_P(value) {
		ref = value.GetCounted()
		value = Z_REFVAL_P(value)
	}
	for {
		if Z_REFCOUNTED_P(variable_ptr) {
			var garbage *ZendRefcounted
			if Z_ISREF_P(variable_ptr) {
				if ZEND_REF_HAS_TYPE_SOURCES(variable_ptr.GetRef()) {
					return ZendAssignToTypedRef(variable_ptr, value, value_type, strict, ref)
				}
				variable_ptr = Z_REFVAL_P(variable_ptr)
				if !(Z_REFCOUNTED_P(variable_ptr)) {
					break
				}
			}
			if variable_ptr.IsType(IS_OBJECT) && Z_OBJ_HANDLER_P(variable_ptr, set) != nil {
				Z_OBJ_HANDLER_P(variable_ptr, set)(variable_ptr, value)
				return variable_ptr
			}
			garbage = variable_ptr.GetCounted()
			ZendCopyToVariable(variable_ptr, value, value_type, ref)
			if garbage.DecGcRefcount() == 0 {
				RcDtorFunc(garbage)
			} else {

				/* optimized version of GC_ZVAL_CHECK_POSSIBLE_ROOT(variable_ptr) */

				if GC_MAY_LEAK(garbage) {
					GcPossibleRoot(garbage)
				}

				/* optimized version of GC_ZVAL_CHECK_POSSIBLE_ROOT(variable_ptr) */

			}
			return variable_ptr
		}
		break
	}
	ZendCopyToVariable(variable_ptr, value, value_type, ref)
	return variable_ptr
}
func ZEND_VM_STACK_ELEMENTS(stack ZendVmStack) __auto__ {
	return (*Zval)(stack) + ZEND_VM_STACK_HEADER_SLOTS
}
func ZendVmInitCallFrame(call *ZendExecuteData, call_info uint32, func_ *ZendFunction, num_args uint32, object_or_called_scope any) {
	call.SetFunc(func_)
	call.GetThis().SetPtr(object_or_called_scope)
	ZEND_CALL_INFO(call) = call_info
	ZEND_CALL_NUM_ARGS(call) = num_args
}
func ZendVmStackPushCallFrameEx(used_stack uint32, call_info uint32, func_ *ZendFunction, num_args uint32, object_or_called_scope any) *ZendExecuteData {
	var call *ZendExecuteData = (*ZendExecuteData)(ExecutorGlobals.GetVmStackTop())
	if used_stack > size_t((*byte)(ExecutorGlobals.GetVmStackEnd())-(*byte)(call)) {
		call = (*ZendExecuteData)(ZendVmStackExtend(used_stack))
		ZendVmInitCallFrame(call, call_info|ZEND_CALL_ALLOCATED, func_, num_args, object_or_called_scope)
		return call
	} else {
		ExecutorGlobals.SetVmStackTop((*Zval)((*byte)(call + used_stack)))
		ZendVmInitCallFrame(call, call_info, func_, num_args, object_or_called_scope)
		return call
	}
}
func ZendVmCalcUsedStack(num_args uint32, func_ *ZendFunction) uint32 {
	var used_stack uint32 = ZEND_CALL_FRAME_SLOT + num_args
	if ZEND_USER_CODE(func_.GetType()) {
		used_stack += func_.GetOpArray().GetLastVar() + func_.GetOpArray().GetT() - MIN(func_.GetOpArray().GetNumArgs(), num_args)
	}
	return used_stack * b.SizeOf("zval")
}
func ZendVmStackPushCallFrame(call_info uint32, func_ *ZendFunction, num_args uint32, object_or_called_scope any) *ZendExecuteData {
	var used_stack uint32 = ZendVmCalcUsedStack(num_args, func_)
	return ZendVmStackPushCallFrameEx(used_stack, call_info, func_, num_args, object_or_called_scope)
}
func ZendVmStackFreeExtraArgsEx(call_info uint32, call *ZendExecuteData) {
	if (call_info & ZEND_CALL_FREE_EXTRA_ARGS) != 0 {
		var count uint32 = ZEND_CALL_NUM_ARGS(call) - call.GetFunc().GetOpArray().GetNumArgs()
		var p *Zval = ZEND_CALL_VAR_NUM(call, call.GetFunc().GetOpArray().GetLastVar()+call.GetFunc().GetOpArray().GetT())
		for {
			if Z_REFCOUNTED_P(p) {
				var r *ZendRefcounted = p.GetCounted()
				if r.DecGcRefcount() == 0 {
					ZVAL_NULL(p)
					RcDtorFunc(r)
				} else {
					GcCheckPossibleRoot(r)
				}
			}
			p++
			if !(b.PreDec(&count)) {
				break
			}
		}
	}
}
func ZendVmStackFreeExtraArgs(call *ZendExecuteData) {
	ZendVmStackFreeExtraArgsEx(ZEND_CALL_INFO(call), call)
}
func ZendVmStackFreeArgs(call *ZendExecuteData) {
	var num_args uint32 = ZEND_CALL_NUM_ARGS(call)
	if num_args > 0 {
		var p *Zval = ZEND_CALL_ARG(call, 1)
		for {
			if Z_REFCOUNTED_P(p) {
				var r *ZendRefcounted = p.GetCounted()
				if r.DecGcRefcount() == 0 {
					ZVAL_NULL(p)
					RcDtorFunc(r)
				}
			}
			p++
			if !(b.PreDec(&num_args)) {
				break
			}
		}
	}
}
func ZendVmStackFreeCallFrameEx(call_info uint32, call *ZendExecuteData) {
	if (call_info & ZEND_CALL_ALLOCATED) != 0 {
		var p ZendVmStack = ExecutorGlobals.GetVmStack()
		var prev ZendVmStack = p.GetPrev()
		ZEND_ASSERT(call == (*ZendExecuteData)(ZEND_VM_STACK_ELEMENTS(ExecutorGlobals.GetVmStack())))
		ExecutorGlobals.SetVmStackTop(prev.GetTop())
		ExecutorGlobals.SetVmStackEnd(prev.GetEnd())
		ExecutorGlobals.SetVmStack(prev)
		Efree(p)
	} else {
		ExecutorGlobals.SetVmStackTop((*Zval)(call))
	}
}
func ZendVmStackFreeCallFrame(call *ZendExecuteData) {
	ZendVmStackFreeCallFrameEx(ZEND_CALL_INFO(call), call)
}
func CACHE_ADDR(num __auto__) *any {
	return (*any)((*byte)(EX(run_time_cache) + num))
}
func CACHED_PTR(num __auto__) any {
	return (*any)((*byte)(EX(run_time_cache) + num))[0]
}
func CACHE_PTR(num __auto__, ptr any) {
	(*any)((*byte)(EX(run_time_cache) + num))[0] = ptr
}
func CACHED_POLYMORPHIC_PTR(num __auto__, ce __auto__) bool {
	return (*any)((*byte)(EX(run_time_cache) + num))[0] == any(b.CondF1(ce, func() any { return (*any)((*byte)(EX(run_time_cache) + num))[1] }, nil))
}
func CACHE_POLYMORPHIC_PTR(num uint32, ce any, ptr any) {
	var slot *any = (*any)((*byte)(EX(run_time_cache) + num))
	slot[0] = ce
	slot[1] = ptr
}
func CACHED_PTR_EX(slot *any) any     { return slot[0] }
func CACHE_PTR_EX(slot *any, ptr any) { slot[0] = ptr }
func CACHED_POLYMORPHIC_PTR_EX(slot __auto__, ce __auto__) __auto__ {
	if slot[0] == ce {
		return slot[1]
	} else {
		return nil
	}
}
func CACHE_POLYMORPHIC_PTR_EX(slot *any, ce *ZendClassEntry, ptr any) {
	slot[0] = ce
	slot[1] = ptr
}
func IS_SPECIAL_CACHE_VAL(ptr *ZendConstant) int { return uintPtr(ptr) & CACHE_SPECIAL }
func ENCODE_SPECIAL_CACHE_NUM(num __auto__) any {
	return any(uintPtr(num)<<1 | CACHE_SPECIAL)
}
func DECODE_SPECIAL_CACHE_NUM(ptr *ZendConstant) int { return uintPtr(ptr) >> 1 }
func ENCODE_SPECIAL_CACHE_PTR(ptr __auto__) any      { return any(uintPtr(ptr) | CACHE_SPECIAL) }
func DECODE_SPECIAL_CACHE_PTR(ptr __auto__) any {
	return any(uintPtr(ptr) & ^CACHE_SPECIAL)
}
func SKIP_EXT_OPLINE(opline __auto__) {
	for opline.opcode >= ZEND_EXT_STMT && opline.opcode <= ZEND_TICKS {
		opline--
	}
}
func ZEND_CLASS_HAS_TYPE_HINTS(ce *ZendClassEntry) bool {
	return (ce.GetCeFlags() & ZEND_ACC_HAS_TYPE_HINTS) == ZEND_ACC_HAS_TYPE_HINTS
}
func ZEND_REF_ADD_TYPE_SOURCE(ref *ZendReference, source *ZendPropertyInfo) {
	ZendRefAddTypeSource(&(ref.GetSources()), source)
}
func ZEND_REF_DEL_TYPE_SOURCE(ref *ZendReference, source *ZendPropertyInfo) {
	ZendRefDelTypeSource(&(ref.GetSources()), source)
}
func GetZvalPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getZvalPtr(op_type, node, should_free, type_, EXECUTE_DATA_C, OPLINE_C)
}
func GetZvalPtrDeref(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getZvalPtrDeref(op_type, node, should_free, type_, EXECUTE_DATA_C, OPLINE_C)
}
func GetZvalPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getZvalPtrUndef(op_type, node, should_free, type_, EXECUTE_DATA_C, OPLINE_C)
}
func GetOpDataZvalPtrR(op_type int, node ZnodeOp, should_free *ZendFreeOp) *Zval {
	return _getOpDataZvalPtrR(op_type, node, should_free, EXECUTE_DATA_C, OPLINE_C)
}
func GetOpDataZvalPtrDerefR(op_type int, node ZnodeOp, should_free *ZendFreeOp) *Zval {
	return _getOpDataZvalPtrDerefR(op_type, node, should_free, EXECUTE_DATA_C, OPLINE_C)
}
func GetZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getZvalPtrPtr(op_type, node, should_free, type_, EXECUTE_DATA_C)
}
func GetZvalPtrPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getZvalPtrPtr(op_type, node, should_free, type_, EXECUTE_DATA_C)
}
func GetObjZvalPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getObjZvalPtr(op_type, node, should_free, type_, EXECUTE_DATA_C, OPLINE_C)
}
func GetObjZvalPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getObjZvalPtrUndef(op_type, node, should_free, type_, EXECUTE_DATA_C, OPLINE_C)
}
func GetObjZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *Zval {
	return _getObjZvalPtrPtr(op_type, node, should_free, type_, EXECUTE_DATA_C)
}
func RETURN_VALUE_USED(opline *ZendOp) bool {
	return opline.GetResultType() != IS_UNUSED
}
func ZifPass(execute_data *ZendExecuteData, return_value *Zval) {}
func FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op *Zval, result *Zval) {
	var __container_to_free *Zval = free_op
	if __container_to_free != nil && Z_REFCOUNTED_P(__container_to_free) {
		var __ref *ZendRefcounted = __container_to_free.GetCounted()
		if __ref.DecGcRefcount() == 0 {
			var __zv *Zval = result
			if __zv.IsType(IS_INDIRECT) {
				ZVAL_COPY(__zv, __zv.GetZv())
			}
			RcDtorFunc(__ref)
		}
	}
}
func FREE_OP(should_free *Zval) {
	if should_free != nil {
		ZvalPtrDtorNogc(should_free)
	}
}
func FREE_UNFETCHED_OP(type_ ZendUchar, var_ uint32) {
	if (type_ & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(EX_VAR(var_))
	}
}
func FREE_OP_VAR_PTR(should_free *Zval) {
	if should_free != nil {
		ZvalPtrDtorNogc(should_free)
	}
}
func CV_DEF_OF(i __auto__) __auto__ { return EX(func_).op_array.vars[i] }
func ZEND_VM_STACK_PAGE_ALIGNED_SIZE(size int, page_size int) int {
	return size + ZEND_VM_STACK_HEADER_SLOTS*b.SizeOf("zval") + (page_size-1) & ^(page_size-1)
}
func ZendVmStackNewPage(size int, prev ZendVmStack) ZendVmStack {
	var page ZendVmStack = ZendVmStack(Emalloc(size))
	page.SetTop(ZEND_VM_STACK_ELEMENTS(page))
	page.SetEnd((*Zval)((*byte)(page + size)))
	page.SetPrev(prev)
	return page
}
func ZendVmStackInit() {
	ExecutorGlobals.SetVmStackPageSize(ZEND_VM_STACK_PAGE_SIZE)
	ExecutorGlobals.SetVmStack(ZendVmStackNewPage(ZEND_VM_STACK_PAGE_SIZE, nil))
	ExecutorGlobals.SetVmStackTop(ExecutorGlobals.GetVmStack().GetTop())
	ExecutorGlobals.SetVmStackEnd(ExecutorGlobals.GetVmStack().GetEnd())
}
func ZendVmStackInitEx(page_size int) {
	/* page_size must be a power of 2 */

	ZEND_ASSERT(page_size > 0 && (page_size&page_size-1) == 0)
	ExecutorGlobals.SetVmStackPageSize(page_size)
	ExecutorGlobals.SetVmStack(ZendVmStackNewPage(page_size, nil))
	ExecutorGlobals.SetVmStackTop(ExecutorGlobals.GetVmStack().GetTop())
	ExecutorGlobals.SetVmStackEnd(ExecutorGlobals.GetVmStack().GetEnd())
}
func ZendVmStackDestroy() {
	var stack ZendVmStack = ExecutorGlobals.GetVmStack()
	for stack != nil {
		var p ZendVmStack = stack.GetPrev()
		Efree(stack)
		stack = p
	}
}
func ZendVmStackExtend(size int) any {
	var stack ZendVmStack
	var ptr any
	stack = ExecutorGlobals.GetVmStack()
	stack.SetTop(ExecutorGlobals.GetVmStackTop())
	stack = ZendVmStackNewPage(b.CondF(size < ExecutorGlobals.GetVmStackPageSize()-ZEND_VM_STACK_HEADER_SLOTS*b.SizeOf("zval"), func() int { return ExecutorGlobals.GetVmStackPageSize() }, func() int { return ZEND_VM_STACK_PAGE_ALIGNED_SIZE(size, ExecutorGlobals.GetVmStackPageSize()) }), stack)
	ExecutorGlobals.SetVmStack(stack)
	ptr = stack.GetTop()
	ExecutorGlobals.SetVmStackTop(any((*byte)(ptr) + size))
	ExecutorGlobals.SetVmStackEnd(stack.GetEnd())
	return ptr
}
func ZendGetCompiledVariableValue(execute_data *ZendExecuteData, var_ uint32) *Zval {
	return EX_VAR(var_)
}
func _getZvalPtrTmp(var_ uint32, should_free *ZendFreeOp, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	*should_free = ret
	ZEND_ASSERT(ret.GetType() != IS_REFERENCE)
	return ret
}
func _getZvalPtrVar(var_ uint32, should_free *ZendFreeOp, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	*should_free = ret
	return ret
}
func _getZvalPtrVarDeref(var_ uint32, should_free *ZendFreeOp, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	*should_free = ret
	ZVAL_DEREF(ret)
	return ret
}
func ZvalUndefinedCv(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	if ExecutorGlobals.GetException() == nil {
		var cv *ZendString = CV_DEF_OF(EX_VAR_TO_NUM(var_))
		ZendError(E_NOTICE, "Undefined variable: %s", cv.GetVal())
	}
	return &(ExecutorGlobals.GetUninitializedZval())
}
func _zvalUndefinedOp1(EXECUTE_DATA_D) *Zval {
	return ZvalUndefinedCv(EX(opline).op1.var_, EXECUTE_DATA_C)
}
func _zvalUndefinedOp2(EXECUTE_DATA_D) *Zval {
	return ZvalUndefinedCv(EX(opline).op2.var_, EXECUTE_DATA_C)
}
func ZVAL_UNDEFINED_OP1() *Zval { return _zvalUndefinedOp1(EXECUTE_DATA_C) }
func ZVAL_UNDEFINED_OP2() *Zval { return _zvalUndefinedOp2(EXECUTE_DATA_C) }
func _getZvalCvLookup(ptr *Zval, var_ uint32, type_ int, _ EXECUTE_DATA_D) *Zval {
	switch type_ {
	case BP_VAR_R:

	case BP_VAR_UNSET:
		ptr = ZvalUndefinedCv(var_, EXECUTE_DATA_C)
		break
	case BP_VAR_IS:
		ptr = &(ExecutorGlobals.GetUninitializedZval())
		break
	case BP_VAR_RW:
		ZvalUndefinedCv(var_, EXECUTE_DATA_C)
	case BP_VAR_W:
		ZVAL_NULL(ptr)
		break
	}
	return ptr
}
func _getZvalPtrCv(var_ uint32, type_ int, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if ret.IsType(IS_UNDEF) {
		if type_ == BP_VAR_W {
			ZVAL_NULL(ret)
		} else {
			return _getZvalCvLookup(ret, var_, type_, EXECUTE_DATA_C)
		}
	}
	return ret
}
func _getZvalPtrCvDeref(var_ uint32, type_ int, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if ret.IsType(IS_UNDEF) {
		if type_ == BP_VAR_W {
			ZVAL_NULL(ret)
			return ret
		} else {
			return _getZvalCvLookup(ret, var_, type_, EXECUTE_DATA_C)
		}
	}
	ZVAL_DEREF(ret)
	return ret
}
func _get_zval_ptr_cv_BP_VAR_R(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if ret.IsType(IS_UNDEF) {
		return ZvalUndefinedCv(var_, EXECUTE_DATA_C)
	}
	return ret
}
func _get_zval_ptr_cv_deref_BP_VAR_R(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if ret.IsType(IS_UNDEF) {
		return ZvalUndefinedCv(var_, EXECUTE_DATA_C)
	}
	ZVAL_DEREF(ret)
	return ret
}
func _get_zval_ptr_cv_BP_VAR_IS(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	return ret
}
func _get_zval_ptr_cv_BP_VAR_RW(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if ret.IsType(IS_UNDEF) {
		ZVAL_NULL(ret)
		ZvalUndefinedCv(var_, EXECUTE_DATA_C)
		return ret
	}
	return ret
}
func _get_zval_ptr_cv_BP_VAR_W(var_ uint32, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if ret.IsType(IS_UNDEF) {
		ZVAL_NULL(ret)
	}
	return ret
}
func _getZvalPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if core.ZEND_DEBUG == 0 || op_type == IS_VAR {
			return _getZvalPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_TMP_VAR)
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return _getZvalPtrCv(node.GetVar(), type_, EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getOpDataZvalPtrR(op_type int, node ZnodeOp, should_free *ZendFreeOp, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if core.ZEND_DEBUG == 0 || op_type == IS_VAR {
			return _getZvalPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_TMP_VAR)
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline+1, node)
		} else if op_type == IS_CV {
			return _get_zval_ptr_cv_BP_VAR_R(node.GetVar(), EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getZvalPtrDeref(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if op_type == IS_TMP_VAR {
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_VAR)
			return _getZvalPtrVarDeref(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return _getZvalPtrCvDeref(node.GetVar(), type_, EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getOpDataZvalPtrDerefR(op_type int, node ZnodeOp, should_free *ZendFreeOp, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if op_type == IS_TMP_VAR {
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_VAR)
			return _getZvalPtrVarDeref(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline+1, node)
		} else if op_type == IS_CV {
			return _get_zval_ptr_cv_deref_BP_VAR_R(node.GetVar(), EXECUTE_DATA_C)
		} else {
			return nil
		}
	}
}
func _getZvalPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if (op_type & (IS_TMP_VAR | IS_VAR)) != 0 {
		if core.ZEND_DEBUG == 0 || op_type == IS_VAR {
			return _getZvalPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
		} else {
			ZEND_ASSERT(op_type == IS_TMP_VAR)
			return _getZvalPtrTmp(node.GetVar(), should_free, EXECUTE_DATA_C)
		}
	} else {
		*should_free = nil
		if op_type == IS_CONST {
			return RT_CONSTANT(opline, node)
		} else if op_type == IS_CV {
			return EX_VAR(node.GetVar())
		} else {
			return nil
		}
	}
}
func _getZvalPtrPtrVar(var_ uint32, should_free *ZendFreeOp, _ EXECUTE_DATA_D) *Zval {
	var ret *Zval = EX_VAR(var_)
	if ret.IsType(IS_INDIRECT) {
		*should_free = nil
		ret = ret.GetZv()
	} else {
		*should_free = ret
	}
	return ret
}
func _getZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D) *Zval {
	if op_type == IS_CV {
		*should_free = nil
		return _getZvalPtrCv(node.GetVar(), type_, EXECUTE_DATA_C)
	} else {
		ZEND_ASSERT(op_type == IS_VAR)
		return _getZvalPtrPtrVar(node.GetVar(), should_free, EXECUTE_DATA_C)
	}
}
func _getObjZvalPtr(op_type int, op ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &(EX(This))
	}
	return GetZvalPtr(op_type, op, should_free, type_)
}
func _getObjZvalPtrUndef(op_type int, op ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D, opline *ZendOp) *Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &(EX(This))
	}
	return GetZvalPtrUndef(op_type, op, should_free, type_)
}
func _getObjZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int, _ EXECUTE_DATA_D) *Zval {
	if op_type == IS_UNUSED {
		*should_free = nil
		return &(EX(This))
	}
	return GetZvalPtrPtr(op_type, node, should_free, type_)
}
func ZendAssignToVariableReference(variable_ptr *Zval, value_ptr *Zval) {
	var ref *ZendReference
	if !(Z_ISREF_P(value_ptr)) {
		ZVAL_NEW_REF(value_ptr, value_ptr)
	} else if variable_ptr == value_ptr {
		return
	}
	ref = value_ptr.GetRef()
	ref.IncGcRefcount()
	if Z_REFCOUNTED_P(variable_ptr) {
		var garbage *ZendRefcounted = variable_ptr.GetCounted()
		if garbage.DecGcRefcount() == 0 {
			ZVAL_REF(variable_ptr, ref)
			RcDtorFunc(garbage)
			return
		} else {
			GcCheckPossibleRoot(garbage)
		}
	}
	ZVAL_REF(variable_ptr, ref)
}
func ZendAssignToTypedPropertyReference(prop_info *ZendPropertyInfo, prop *Zval, value_ptr *Zval, _ EXECUTE_DATA_D) *Zval {
	if ZendVerifyPropAssignableByRef(prop_info, value_ptr, EX_USES_STRICT_TYPES()) == 0 {
		return &(ExecutorGlobals.GetUninitializedZval())
	}
	if Z_ISREF_P(prop) {
		ZEND_REF_DEL_TYPE_SOURCE(prop.GetRef(), prop_info)
	}
	ZendAssignToVariableReference(prop, value_ptr)
	ZEND_REF_ADD_TYPE_SOURCE(prop.GetRef(), prop_info)
	return prop
}
func ZendWrongAssignToVariableReference(variable_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) *Zval {
	ZendError(E_NOTICE, "Only variables should be assigned by reference")
	if ExecutorGlobals.GetException() != nil {
		return &(ExecutorGlobals.GetUninitializedZval())
	}

	/* Use IS_TMP_VAR instead of IS_VAR to avoid ISREF check */

	Z_TRY_ADDREF_P(value_ptr)
	return ZendAssignToVariable(variable_ptr, value_ptr, IS_TMP_VAR, EX_USES_STRICT_TYPES())
}
func ZendFormatType(type_ ZendType, part1 **byte, part2 **byte) {
	if type_.AllowNull() {
		*part1 = "?"
	} else {
		*part1 = ""
	}
	if type_.IsClass() {
		if type_.IsCe() {
			*part2 = type_.Ce().GetName().GetVal()
		} else {
			*part2 = type_.Name().GetVal()
		}
	} else {
		*part2 = ZendGetTypeByConst(type_.Code())
	}
}
func ZendThrowAutoInitInPropError(prop *ZendPropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot auto-initialize an %s inside property %s::$%s of type %s%s", type_, prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowAutoInitInRefError(prop *ZendPropertyInfo, type_ string) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot auto-initialize an %s inside a reference held by property %s::$%s of type %s%s", type_, prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowAccessUninitPropByRefError(prop *ZendPropertyInfo) {
	ZendThrowError(nil, "Cannot access uninitialized non-nullable property %s::$%s by reference", prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()))
}
func MakeRealObject(object *Zval, property *Zval, opline *ZendOp, _ EXECUTE_DATA_D) *Zval {
	var obj *ZendObject
	var ref *Zval = nil
	if Z_ISREF_P(object) {
		ref = object
		object = Z_REFVAL_P(object)
	}
	if object.GetType() > IS_FALSE && (object.GetType() != IS_STRING || Z_STRLEN_P(object) != 0) {
		if opline.GetOp1Type() != IS_VAR || !(Z_ISERROR_P(object)) {
			var tmp_property_name *ZendString
			var property_name *ZendString = ZvalGetTmpString(property, &tmp_property_name)
			if opline.GetOpcode() == ZEND_PRE_INC_OBJ || opline.GetOpcode() == ZEND_PRE_DEC_OBJ || opline.GetOpcode() == ZEND_POST_INC_OBJ || opline.GetOpcode() == ZEND_POST_DEC_OBJ {
				ZendError(E_WARNING, "Attempt to increment/decrement property '%s' of non-object", property_name.GetVal())
			} else if opline.GetOpcode() == ZEND_FETCH_OBJ_W || opline.GetOpcode() == ZEND_FETCH_OBJ_RW || opline.GetOpcode() == ZEND_FETCH_OBJ_FUNC_ARG || opline.GetOpcode() == ZEND_ASSIGN_OBJ_REF {
				ZendError(E_WARNING, "Attempt to modify property '%s' of non-object", property_name.GetVal())
			} else {
				ZendError(E_WARNING, "Attempt to assign property '%s' of non-object", property_name.GetVal())
			}
			ZendTmpStringRelease(tmp_property_name)
		}
		if RETURN_VALUE_USED(opline) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
		return nil
	}
	if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(ref.GetRef()) {
		if zend_verify_ref_stdClass_assignable(ref.GetRef()) == 0 {
			if RETURN_VALUE_USED(opline) {
				ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
			}
			return nil
		}
	}
	ZvalPtrDtorNogc(object)
	ObjectInit(object)
	Z_ADDREF_P(object)
	obj = object.GetObj()
	ZendError(E_WARNING, "Creating default object from empty value")
	if obj.GetGcRefcount() == 1 {

		/* the enclosing container was deleted, obj is unreferenced */

		OBJ_RELEASE(obj)
		if RETURN_VALUE_USED(opline) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
		return nil
	}
	Z_DELREF_P(object)
	return object
}
func ZendVerifyTypeErrorCommon(zf *ZendFunction, arg_info *ZendArgInfo, ce *ZendClassEntry, value *Zval, fname **byte, fsep **byte, fclass **byte, need_msg **byte, need_kind **byte, need_or_null **byte, given_msg **byte, given_kind **byte) {
	var is_interface ZendBool = 0
	*fname = zf.GetFunctionName().GetVal()
	if zf.GetScope() != nil {
		*fsep = "::"
		*fclass = zf.GetScope().GetName().GetVal()
	} else {
		*fsep = ""
		*fclass = ""
	}
	if arg_info.GetType().IsClass() {
		if ce != nil {
			if ce.IsInterface() {
				*need_msg = "implement interface "
				is_interface = 1
			} else {
				*need_msg = "be an instance of "
			}
			*need_kind = ce.GetName().GetVal()
		} else {

			/* We don't know whether it's a class or interface, assume it's a class */

			*need_msg = "be an instance of "
			*need_kind = arg_info.GetType().Name().GetVal()
		}
	} else {
		switch arg_info.GetType().Code() {
		case IS_OBJECT:
			*need_msg = "be an "
			*need_kind = "object"
			break
		case IS_CALLABLE:
			*need_msg = "be callable"
			*need_kind = ""
			break
		case IS_ITERABLE:
			*need_msg = "be iterable"
			*need_kind = ""
			break
		default:
			*need_msg = "be of the type "
			*need_kind = ZendGetTypeByConst(arg_info.GetType().Code())
			break
		}
	}
	if arg_info.GetType().AllowNull() {
		if is_interface != 0 {
			*need_or_null = " or be null"
		} else {
			*need_or_null = " or null"
		}
	} else {
		*need_or_null = ""
	}
	if value != nil {
		if arg_info.GetType().IsClass() && value.IsType(IS_OBJECT) {
			*given_msg = "instance of "
			*given_kind = Z_OBJCE_P(value).GetName().GetVal()
		} else {
			*given_msg = ZendZvalTypeName(value)
			*given_kind = ""
		}
	} else {
		*given_msg = "none"
		*given_kind = ""
	}
}
func ZendVerifyArgError(zf *ZendFunction, arg_info *ZendArgInfo, arg_num int, ce *ZendClassEntry, value *Zval) {
	var ptr *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData().GetPrevExecuteData()
	var fname *byte
	var fsep *byte
	var fclass *byte
	var need_msg *byte
	var need_kind *byte
	var need_or_null *byte
	var given_msg *byte
	var given_kind *byte
	if ExecutorGlobals.GetException() != nil {

		/* The type verification itself might have already thrown an exception
		 * through a promoted warning. */

		return

		/* The type verification itself might have already thrown an exception
		 * through a promoted warning. */

	}
	if value != nil {
		ZendVerifyTypeErrorCommon(zf, arg_info, ce, value, &fname, &fsep, &fclass, &need_msg, &need_kind, &need_or_null, &given_msg, &given_kind)
		if zf.GetCommonType() == ZEND_USER_FUNCTION {
			if ptr != nil && ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetCommonType()) {
				ZendTypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given, called in %s on line %d", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind, ptr.GetFunc().GetOpArray().GetFilename().GetVal(), ptr.GetOpline().GetLineno())
			} else {
				ZendTypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
			}
		} else {
			ZendTypeError("Argument %d passed to %s%s%s() must %s%s%s, %s%s given", arg_num, fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
		}
	} else {
		ZendMissingArgError(ptr)
	}
}
func IsNullConstant(scope *ZendClassEntry, default_value *Zval) int {
	if default_value.IsType(IS_CONSTANT_AST) {
		var constant Zval
		ZVAL_COPY(&constant, default_value)
		if ZvalUpdateConstantEx(&constant, scope) != SUCCESS {
			return 0
		}
		if constant.IsType(IS_NULL) {
			return 1
		}
		ZvalPtrDtorNogc(&constant)
	}
	return 0
}
func ZendVerifyWeakScalarTypeHint(type_hint ZendUchar, arg *Zval) ZendBool {
	switch type_hint {
	case _IS_BOOL:
		var dest ZendBool
		if ZendParseArgBoolWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		ZVAL_BOOL(arg, dest)
		return 1
	case IS_LONG:
		var dest ZendLong
		if ZendParseArgLongWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		ZVAL_LONG(arg, dest)
		return 1
	case IS_DOUBLE:
		var dest float64
		if ZendParseArgDoubleWeak(arg, &dest) == 0 {
			return 0
		}
		ZvalPtrDtor(arg)
		ZVAL_DOUBLE(arg, dest)
		return 1
	case IS_STRING:
		var dest *ZendString

		/* on success "arg" is converted to IS_STRING */

		return ZendParseArgStrWeak(arg, &dest)

	/* on success "arg" is converted to IS_STRING */

	default:
		return 0
	}
}
func ZendVerifyScalarTypeHint(type_hint ZendUchar, arg *Zval, strict ZendBool) ZendBool {
	if strict != 0 {

		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

		if type_hint != IS_DOUBLE || arg.GetType() != IS_LONG {
			return 0
		}

		/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

	} else if arg.IsType(IS_NULL) {

		/* NULL may be accepted only by nullable hints (this is already checked) */

		return 0

		/* NULL may be accepted only by nullable hints (this is already checked) */

	}
	return ZendVerifyWeakScalarTypeHint(type_hint, arg)
}
func ZendVerifyPropertyTypeError(info *ZendPropertyInfo, property *Zval) {
	var prop_type1 *byte
	var prop_type2 *byte

	/* we _may_ land here in case reading already errored and runtime cache thus has not been updated (i.e. it contains a valid but unrelated info) */

	if ExecutorGlobals.GetException() != nil {
		return
	}

	// TODO Switch to a more standard error message?

	ZendFormatType(info.GetType(), &prop_type1, &prop_type2)
	void(prop_type1)
	if info.GetType().IsClass() {
		ZendTypeError("Typed property %s::$%s must be an instance of %s%s, %s used", info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, b.Cond(info.GetType().AllowNull(), " or null", ""), b.CondF(property.IsType(IS_OBJECT), func() []byte { return Z_OBJCE_P(property).GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(property.GetType()) }))
	} else {
		ZendTypeError("Typed property %s::$%s must be %s%s, %s used", info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(info.GetName()), prop_type2, b.Cond(info.GetType().AllowNull(), " or null", ""), b.CondF(property.IsType(IS_OBJECT), func() []byte { return Z_OBJCE_P(property).GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(property.GetType()) }))
	}
}
func ZendResolveClassType(type_ *ZendType, self_ce *ZendClassEntry) ZendBool {
	var ce *ZendClassEntry
	var name *ZendString = type_.Name()
	if ZendStringEqualsLiteralCi(name, "self") {

		/* We need to explicitly check for this here, to avoid updating the type in the trait and
		 * later using the wrong "self" when the trait is used in a class. */

		if self_ce.IsTrait() {
			ZendThrowError(nil, "Cannot write a%s value to a 'self' typed static property of a trait", b.Cond(type_.AllowNull(), " non-null", ""))
			return 0
		}
		ce = self_ce
	} else if ZendStringEqualsLiteralCi(name, "parent") {
		if !(self_ce.parent) {
			ZendThrowError(nil, "Cannot access parent:: when current class scope has no parent")
			return 0
		}
		ce = self_ce.parent
	} else {
		ce = ZendLookupClassEx(name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if ce == nil {
			return 0
		}
	}
	ZendStringRelease(name)
	*type_ = ZEND_TYPE_ENCODE_CE(ce, type_.AllowNull())
	return 1
}
func IZendCheckPropertyType(info *ZendPropertyInfo, property *Zval, strict ZendBool) ZendBool {
	ZEND_ASSERT(!(Z_ISREF_P(property)))
	if info.GetType().IsClass() {
		if property.GetType() != IS_OBJECT {
			return property.IsType(IS_NULL) && info.GetType().AllowNull()
		}
		if !(info.GetType().IsCe()) && ZendResolveClassType(info.GetType(), info.GetCe()) == 0 {
			return 0
		}
		return InstanceofFunction(Z_OBJCE_P(property), info.GetType().Ce())
	}
	ZEND_ASSERT(info.GetType().Code() != IS_CALLABLE)
	if info.GetType().Code() == property.GetType() {
		return 1
	} else if property.IsType(IS_NULL) {
		return info.GetType().AllowNull()
	} else if info.GetType().Code() == _IS_BOOL && property.IsType(IS_FALSE) || property.IsType(IS_TRUE) {
		return 1
	} else if info.GetType().Code() == IS_ITERABLE {
		return ZendIsIterable(property)
	} else {
		return ZendVerifyScalarTypeHint(info.GetType().Code(), property, strict)
	}
}
func IZendVerifyPropertyType(info *ZendPropertyInfo, property *Zval, strict ZendBool) ZendBool {
	if IZendCheckPropertyType(info, property, strict) != 0 {
		return 1
	}
	ZendVerifyPropertyTypeError(info, property)
	return 0
}
func ZendVerifyPropertyType(info *ZendPropertyInfo, property *Zval, strict ZendBool) ZendBool {
	return IZendVerifyPropertyType(info, property, strict)
}
func ZendAssignToTypedProp(info *ZendPropertyInfo, property_val *Zval, value *Zval, _ EXECUTE_DATA_D) *Zval {
	var tmp Zval
	ZVAL_DEREF(value)
	ZVAL_COPY(&tmp, value)
	if IZendVerifyPropertyType(info, &tmp, EX_USES_STRICT_TYPES()) == 0 {
		ZvalPtrDtor(&tmp)
		return &(ExecutorGlobals.GetUninitializedZval())
	}
	return ZendAssignToVariable(property_val, &tmp, IS_TMP_VAR, EX_USES_STRICT_TYPES())
}
func ZendCheckType(type_ ZendType, arg *Zval, ce **ZendClassEntry, cache_slot *any, default_value *Zval, scope *ZendClassEntry, is_return_type ZendBool) ZendBool {
	var ref *ZendReference = nil
	if !(type_.IsSet()) {
		return 1
	}
	if Z_ISREF_P(arg) {
		ref = arg.GetRef()
		arg = Z_REFVAL_P(arg)
	}
	if type_.IsClass() {
		if *cache_slot {
			*ce = (*ZendClassEntry)(*cache_slot)
		} else {
			*ce = ZendFetchClass(type_.Name(), ZEND_FETCH_CLASS_AUTO|ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if (*ce) == nil {
				return arg.IsType(IS_NULL) && (type_.AllowNull() || default_value != nil && IsNullConstant(scope, default_value) != 0)
			}
			*cache_slot = any(*ce)
		}
		if arg.IsType(IS_OBJECT) {
			return InstanceofFunction(Z_OBJCE_P(arg), *ce)
		}
		return arg.IsType(IS_NULL) && (type_.AllowNull() || default_value != nil && IsNullConstant(scope, default_value) != 0)
	} else if type_.Code() == arg.GetType() {
		return 1
	}
	if arg.IsType(IS_NULL) && (type_.AllowNull() || default_value != nil && IsNullConstant(scope, default_value) != 0) {

		/* Null passed to nullable type */

		return 1

		/* Null passed to nullable type */

	}
	if type_.Code() == IS_CALLABLE {
		return ZendIsCallable(arg, IS_CALLABLE_CHECK_SILENT, nil)
	} else if type_.Code() == IS_ITERABLE {
		return ZendIsIterable(arg)
	} else if type_.Code() == _IS_BOOL && arg.IsType(IS_FALSE) || arg.IsType(IS_TRUE) {
		return 1
	} else if ref != nil && ZEND_REF_HAS_TYPE_SOURCES(ref) {
		return 0
	} else {
		return ZendVerifyScalarTypeHint(type_.Code(), arg, b.CondF(is_return_type != 0, func() bool { return ZEND_RET_USES_STRICT_TYPES() }, func() bool { return ZEND_ARG_USES_STRICT_TYPES() }))
	}
}
func ZendVerifyArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *ZendClassEntry
	if arg_num <= zf.GetNumArgs() {
		cur_arg_info = zf.GetArgInfo()[arg_num-1]
	} else if zf.IsVariadic() {
		cur_arg_info = zf.GetArgInfo()[zf.GetNumArgs()]
	} else {
		return 1
	}
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyRecvArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo = zf.GetArgInfo()[arg_num-1]
	var ce *ZendClassEntry
	ZEND_ASSERT(arg_num <= zf.GetNumArgs())
	cur_arg_info = zf.GetArgInfo()[arg_num-1]
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyVariadicArgType(zf *ZendFunction, arg_num uint32, arg *Zval, default_value *Zval, cache_slot *any) int {
	var cur_arg_info *ZendArgInfo
	var ce *ZendClassEntry
	ZEND_ASSERT(arg_num > zf.GetNumArgs())
	ZEND_ASSERT(zf.IsVariadic())
	cur_arg_info = zf.GetArgInfo()[zf.GetNumArgs()]
	ce = nil
	if ZendCheckType(cur_arg_info.GetType(), arg, &ce, cache_slot, default_value, zf.GetScope(), 0) == 0 {
		ZendVerifyArgError(zf, cur_arg_info, arg_num, ce, arg)
		return 0
	}
	return 1
}
func ZendVerifyInternalArgTypes(fbc *ZendFunction, call *ZendExecuteData) int {
	var i uint32
	var num_args uint32 = ZEND_CALL_NUM_ARGS(call)
	var p *Zval = ZEND_CALL_ARG(call, 1)
	var dummy_cache_slot any
	for i = 0; i < num_args; i++ {
		dummy_cache_slot = nil
		if ZendVerifyArgType(fbc, i+1, p, nil, &dummy_cache_slot) == 0 {
			ExecutorGlobals.SetCurrentExecuteData(call.GetPrevExecuteData())
			return 0
		}
		p++
	}
	return 1
}
func ZendMissingArgError(execute_data *ZendExecuteData) {
	var ptr *ZendExecuteData = EX(prev_execute_data)
	if ptr != nil && ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetCommonType()) {
		ZendThrowError(ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed in %s on line %d and %s %d expected", b.CondF1(EX(func_).common.scope, func() []byte { return EX(func_).common.scope.name.GetVal() }, ""), b.Cond(EX(func_).common.scope, "::", ""), EX(func_).common.function_name.GetVal(), EX_NUM_ARGS(), ptr.GetFunc().GetOpArray().GetFilename().GetVal(), ptr.GetOpline().GetLineno(), b.Cond(EX(func_).common.required_num_args == EX(func_).common.num_args, "exactly", "at least"), EX(func_).common.required_num_args)
	} else {
		ZendThrowError(ZendCeArgumentCountError, "Too few arguments to function %s%s%s(), %d passed and %s %d expected", b.CondF1(EX(func_).common.scope, func() []byte { return EX(func_).common.scope.name.GetVal() }, ""), b.Cond(EX(func_).common.scope, "::", ""), EX(func_).common.function_name.GetVal(), EX_NUM_ARGS(), b.Cond(EX(func_).common.required_num_args == EX(func_).common.num_args, "exactly", "at least"), EX(func_).common.required_num_args)
	}
}
func ZendVerifyReturnError(zf *ZendFunction, ce *ZendClassEntry, value *Zval) {
	var arg_info *ZendArgInfo = zf.GetArgInfo()[-1]
	var fname *byte
	var fsep *byte
	var fclass *byte
	var need_msg *byte
	var need_kind *byte
	var need_or_null *byte
	var given_msg *byte
	var given_kind *byte
	ZendVerifyTypeErrorCommon(zf, arg_info, ce, value, &fname, &fsep, &fclass, &need_msg, &need_kind, &need_or_null, &given_msg, &given_kind)
	ZendTypeError("Return value of %s%s%s() must %s%s%s, %s%s returned", fclass, fsep, fname, need_msg, need_kind, need_or_null, given_msg, given_kind)
}
func ZendVerifyReturnType(zf *ZendFunction, ret *Zval, cache_slot *any) {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	var ce *ZendClassEntry = nil
	if ZendCheckType(ret_info.GetType(), ret, &ce, cache_slot, nil, nil, 1) == 0 {
		ZendVerifyReturnError(zf, ce, ret)
	}
}
func ZendVerifyMissingReturnType(zf *ZendFunction, cache_slot *any) int {
	var ret_info *ZendArgInfo = zf.GetArgInfo() - 1
	if ret_info.GetType().IsSet() && ret_info.GetType().Code() != IS_VOID {
		var ce *ZendClassEntry = nil
		if ret_info.GetType().IsClass() {
			if *cache_slot {
				ce = (*ZendClassEntry)(*cache_slot)
			} else {
				ce = ZendFetchClass(ret_info.GetType().Name(), ZEND_FETCH_CLASS_AUTO|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if ce != nil {
					*cache_slot = any(ce)
				}
			}
		}
		ZendVerifyReturnError(zf, ce, nil)
		return 0
	}
	return 1
}
func ZendUseObjectAsArray() {
	ZendThrowError(nil, "Cannot use object as array")
}
func ZendIllegalOffset() {
	ZendError(E_WARNING, "Illegal offset type")
}
func ZendAssignToObjectDim(object *Zval, dim *Zval, value *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	Z_OBJ_HT_P(object).GetWriteDimension()(object, dim, value)
	if RETURN_VALUE_USED(opline) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), value)
	}
}
func ZendBinaryOp(ret *Zval, op1 *Zval, op2 *Zval, opline *ZendOp) int {
	var zend_binary_ops []BinaryOpType = []BinaryOpType{AddFunction, SubFunction, MulFunction, DivFunction, ModFunction, ShiftLeftFunction, ShiftRightFunction, ConcatFunction, BitwiseOrFunction, BitwiseAndFunction, BitwiseXorFunction, PowFunction}

	/* size_t cast makes GCC to better optimize 64-bit PIC code */

	var opcode int = int(opline.GetExtendedValue())
	return zend_binary_ops[opcode-ZEND_ADD](ret, op1, op2)
}
func ZendBinaryAssignOpObjDim(object *Zval, property *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var free_op_data1 ZendFreeOp
	var value *Zval
	var z *Zval
	var rv Zval
	var res Zval
	value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1)
	if b.Assign(&z, Z_OBJ_HT_P(object).GetReadDimension()(object, property, BP_VAR_R, &rv)) != nil {
		if z.IsType(IS_OBJECT) && Z_OBJ_HT_P(z).GetGet() != nil {
			var rv2 Zval
			var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
			if z == &rv {
				ZvalPtrDtor(&rv)
			}
			ZVAL_COPY_VALUE(z, value)
		}
		if ZendBinaryOp(&res, z, value, OPLINE_C) == SUCCESS {
			Z_OBJ_HT_P(object).GetWriteDimension()(object, property, &res)
		}
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		if RETURN_VALUE_USED(opline) {
			ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &res)
		}
		ZvalPtrDtor(&res)
	} else {
		ZendUseObjectAsArray()
		if RETURN_VALUE_USED(opline) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
	}
	FREE_OP(free_op_data1)
}
func ZendBinaryAssignOpTypedRef(ref *ZendReference, value *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var z_copy Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && ref.GetVal().IsType(IS_STRING) {
		ConcatFunction(ref.GetVal(), ref.GetVal(), value)
		ZEND_ASSERT(ref.GetVal().IsType(IS_STRING) && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, ref.GetVal(), value, OPLINE_C)
	if ZendVerifyRefAssignableZval(ref, &z_copy, EX_USES_STRICT_TYPES()) != 0 {
		ZvalPtrDtor(ref.GetVal())
		ZVAL_COPY_VALUE(ref.GetVal(), &z_copy)
	} else {
		ZvalPtrDtor(&z_copy)
	}
}
func ZendBinaryAssignOpTypedProp(prop_info *ZendPropertyInfo, zptr *Zval, value *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var z_copy Zval

	/* Make sure that in-place concatenation is used if the LHS is a string. */

	if opline.GetExtendedValue() == ZEND_CONCAT && zptr.IsType(IS_STRING) {
		ConcatFunction(zptr, zptr, value)
		ZEND_ASSERT(zptr.IsType(IS_STRING) && "Concat should return string")
		return
	}
	ZendBinaryOp(&z_copy, zptr, value, OPLINE_C)
	if ZendVerifyPropertyType(prop_info, &z_copy, EX_USES_STRICT_TYPES()) != 0 {
		ZvalPtrDtor(zptr)
		ZVAL_COPY_VALUE(zptr, &z_copy)
	} else {
		ZvalPtrDtor(&z_copy)
	}
}
func ZendCheckStringOffset(dim *Zval, type_ int, _ EXECUTE_DATA_D) ZendLong {
	var offset ZendLong
try_again:
	if dim.GetType() != IS_LONG {
		switch dim.GetType() {
		case IS_STRING:
			if IS_LONG == IsNumericString(Z_STRVAL_P(dim), Z_STRLEN_P(dim), nil, nil, -1) {
				break
			}
			if type_ != BP_VAR_UNSET {
				ZendError(E_WARNING, "Illegal string offset '%s'", Z_STRVAL_P(dim))
			}
			break
		case IS_UNDEF:
			ZVAL_UNDEFINED_OP2()
		case IS_DOUBLE:

		case IS_NULL:

		case IS_FALSE:

		case IS_TRUE:
			ZendError(E_NOTICE, "String offset cast occurred")
			break
		case IS_REFERENCE:
			dim = Z_REFVAL_P(dim)
			goto try_again
		default:
			ZendIllegalOffset()
			break
		}
		offset = ZvalGetLongFunc(dim)
	} else {
		offset = dim.GetLval()
	}
	return offset
}
func ZendWrongStringOffset(EXECUTE_DATA_D) {
	var msg *byte = nil
	var opline *ZendOp = EX(opline)
	var end *ZendOp
	var var_ uint32
	if ExecutorGlobals.GetException() != nil {
		return
	}
	switch opline.GetOpcode() {
	case ZEND_ASSIGN_OP:

	case ZEND_ASSIGN_DIM_OP:

	case ZEND_ASSIGN_OBJ_OP:

	case ZEND_ASSIGN_STATIC_PROP_OP:
		msg = "Cannot use assign-op operators with string offsets"
		break
	case ZEND_FETCH_DIM_W:

	case ZEND_FETCH_DIM_RW:

	case ZEND_FETCH_DIM_FUNC_ARG:

	case ZEND_FETCH_DIM_UNSET:

	case ZEND_FETCH_LIST_W:

		/* TODO: Encode the "reason" into opline->extended_value??? */

		var_ = opline.GetResult().GetVar()
		opline++
		end = ExecutorGlobals.GetCurrentExecuteData().GetFunc().GetOpArray().GetOpcodes() + ExecutorGlobals.GetCurrentExecuteData().GetFunc().GetOpArray().GetLast()
		for opline < end {
			if opline.GetOp1Type() == IS_VAR && opline.GetOp1().GetVar() == var_ {
				switch opline.GetOpcode() {
				case ZEND_FETCH_OBJ_W:

				case ZEND_FETCH_OBJ_RW:

				case ZEND_FETCH_OBJ_FUNC_ARG:

				case ZEND_FETCH_OBJ_UNSET:

				case ZEND_ASSIGN_OBJ:

				case ZEND_ASSIGN_OBJ_OP:

				case ZEND_ASSIGN_OBJ_REF:
					msg = "Cannot use string offset as an object"
					break
				case ZEND_FETCH_DIM_W:

				case ZEND_FETCH_DIM_RW:

				case ZEND_FETCH_DIM_FUNC_ARG:

				case ZEND_FETCH_DIM_UNSET:

				case ZEND_FETCH_LIST_W:

				case ZEND_ASSIGN_DIM:

				case ZEND_ASSIGN_DIM_OP:
					msg = "Cannot use string offset as an array"
					break
				case ZEND_ASSIGN_STATIC_PROP_OP:

				case ZEND_ASSIGN_OP:
					msg = "Cannot use assign-op operators with string offsets"
					break
				case ZEND_PRE_INC_OBJ:

				case ZEND_PRE_DEC_OBJ:

				case ZEND_POST_INC_OBJ:

				case ZEND_POST_DEC_OBJ:

				case ZEND_PRE_INC:

				case ZEND_PRE_DEC:

				case ZEND_POST_INC:

				case ZEND_POST_DEC:
					msg = "Cannot increment/decrement string offsets"
					break
				case ZEND_ASSIGN_REF:

				case ZEND_ADD_ARRAY_ELEMENT:

				case ZEND_INIT_ARRAY:

				case ZEND_MAKE_REF:
					msg = "Cannot create references to/from string offsets"
					break
				case ZEND_RETURN_BY_REF:

				case ZEND_VERIFY_RETURN_TYPE:
					msg = "Cannot return string offsets by reference"
					break
				case ZEND_UNSET_DIM:

				case ZEND_UNSET_OBJ:
					msg = "Cannot unset string offsets"
					break
				case ZEND_YIELD:
					msg = "Cannot yield string offsets by reference"
					break
				case ZEND_SEND_REF:

				case ZEND_SEND_VAR_EX:

				case ZEND_SEND_FUNC_ARG:
					msg = "Only variables can be passed by reference"
					break
				case ZEND_FE_RESET_RW:
					msg = "Cannot iterate on string offsets by reference"
					break
				default:
					break
				}
				break
			}
			if opline.GetOp2Type() == IS_VAR && opline.GetOp2().GetVar() == var_ {
				ZEND_ASSERT(opline.GetOpcode() == ZEND_ASSIGN_REF)
				msg = "Cannot create references to/from string offsets"
				break
			}
			opline++
		}
		break
	default:
		break
	}
	ZEND_ASSERT(msg != nil)
	ZendThrowError(nil, "%s", msg)
}
func ZendWrongPropertyRead(property *Zval) {
	var tmp_property_name *ZendString
	var property_name *ZendString = ZvalGetTmpString(property, &tmp_property_name)
	ZendError(E_NOTICE, "Trying to get property '%s' of non-object", property_name.GetVal())
	ZendTmpStringRelease(tmp_property_name)
}
func ZendDeprecatedFunction(fbc *ZendFunction) {
	ZendError(E_DEPRECATED, "Function %s%s%s() is deprecated", b.CondF1(fbc.GetScope() != nil, func() []byte { return fbc.GetScope().GetName().GetVal() }, ""), b.Cond(fbc.GetScope() != nil, "::", ""), fbc.GetFunctionName().GetVal())
}
func ZendAbstractMethod(fbc *ZendFunction) {
	ZendThrowError(nil, "Cannot call abstract method %s::%s()", fbc.GetScope().GetName().GetVal(), fbc.GetFunctionName().GetVal())
}
func ZendAssignToStringOffset(str *Zval, dim *Zval, value *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var c ZendUchar
	var string_len int
	var offset ZendLong
	offset = ZendCheckStringOffset(dim, BP_VAR_W, EXECUTE_DATA_C)
	if ExecutorGlobals.GetException() != nil {
		if RETURN_VALUE_USED(opline) {
			ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
		}
		return
	}
	if offset < -ZendLong(Z_STRLEN_P(str)) {

		/* Error on negative offset */

		ZendError(E_WARNING, "Illegal string offset:  "+ZEND_LONG_FMT, offset)
		if RETURN_VALUE_USED(opline) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
		return
	}
	if value.GetType() != IS_STRING {

		/* Convert to string, just the time to pick the 1st byte */

		var tmp *ZendString = ZvalTryGetStringFunc(value)
		if tmp == nil {
			if RETURN_VALUE_USED(opline) {
				ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
			}
			return
		}
		string_len = tmp.GetLen()
		c = ZendUchar(tmp.GetVal()[0])
		ZendStringReleaseEx(tmp, 0)
	} else {
		string_len = Z_STRLEN_P(value)
		c = ZendUchar(Z_STRVAL_P(value)[0])
	}
	if string_len == 0 {

		/* Error on empty input string */

		ZendError(E_WARNING, "Cannot assign an empty string to a string offset")
		if RETURN_VALUE_USED(opline) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
		return
	}
	if offset < 0 {
		offset += ZendLong(Z_STRLEN_P(str))
	}
	if int(offset >= Z_STRLEN_P(str)) != 0 {

		/* Extend string if needed */

		var old_len ZendLong = Z_STRLEN_P(str)
		ZVAL_NEW_STR(str, ZendStringExtend(str.GetStr(), offset+1, 0))
		memset(Z_STRVAL_P(str)+old_len, ' ', offset-old_len)
		Z_STRVAL_P(str)[offset+1] = 0
	} else if !(Z_REFCOUNTED_P(str)) {
		ZVAL_NEW_STR(str, ZendStringInit(Z_STRVAL_P(str), Z_STRLEN_P(str), 0))
	} else if Z_REFCOUNT_P(str) > 1 {
		Z_DELREF_P(str)
		ZVAL_NEW_STR(str, ZendStringInit(Z_STRVAL_P(str), Z_STRLEN_P(str), 0))
	} else {
		ZendStringForgetHashVal(str.GetStr())
	}
	Z_STRVAL_P(str)[offset] = c
	if RETURN_VALUE_USED(opline) {

		/* Return the new character */

		ZVAL_INTERNED_STR(EX_VAR(opline.GetResult().GetVar()), ZSTR_CHAR(c))

		/* Return the new character */

	}
}
func ZendGetPropNotAcceptingDouble(ref *ZendReference) *ZendPropertyInfo {
	var prop *ZendPropertyInfo
	var _source_list *ZendPropertyInfoSourceList = &(ref.GetSources())
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
			prop = *_prop
			if prop.GetType().Code() != IS_DOUBLE {
				return prop
			}
		}
	}
	return nil
}
func ZendThrowIncdecRefError(ref *ZendReference, opline *ZendOp) ZendLong {
	var error_prop *ZendPropertyInfo = ZendGetPropNotAcceptingDouble(ref)

	/* Currently there should be no way for a typed reference to accept both int and double.
	 * Generalize this and the related property code once this becomes possible. */

	ZEND_ASSERT(error_prop != nil)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		ZendTypeError("Cannot increment a reference held by property %s::$%s of type %sint past its maximal value", error_prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(error_prop.GetName()), b.Cond(error_prop.GetType().AllowNull(), "?", ""))
		return ZEND_LONG_MAX
	} else {
		ZendTypeError("Cannot decrement a reference held by property %s::$%s of type %sint past its minimal value", error_prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(error_prop.GetName()), b.Cond(error_prop.GetType().AllowNull(), "?", ""))
		return ZEND_LONG_MIN
	}
}
func ZendThrowIncdecPropError(prop *ZendPropertyInfo, opline *ZendOp) ZendLong {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		ZendTypeError("Cannot increment property %s::$%s of type %s%s past its maximal value", prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
		return ZEND_LONG_MAX
	} else {
		ZendTypeError("Cannot decrement property %s::$%s of type %s%s past its minimal value", prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
		return ZEND_LONG_MIN
	}
}
func ZendIncdecTypedRef(ref *ZendReference, copy *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var tmp Zval
	var var_ptr *Zval = ref.GetVal()
	if copy == nil {
		copy = &tmp
	}
	ZVAL_COPY(copy, var_ptr)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		IncrementFunction(var_ptr)
	} else {
		DecrementFunction(var_ptr)
	}
	if var_ptr.IsType(IS_DOUBLE) && copy.IsType(IS_LONG) {
		var val ZendLong = ZendThrowIncdecRefError(ref, OPLINE_C)
		ZVAL_LONG(var_ptr, val)
	} else if ZendVerifyRefAssignableZval(ref, var_ptr, EX_USES_STRICT_TYPES()) == 0 {
		ZvalPtrDtor(var_ptr)
		ZVAL_COPY_VALUE(var_ptr, copy)
		ZVAL_UNDEF(copy)
	} else if copy == &tmp {
		ZvalPtrDtor(&tmp)
	}
}
func ZendIncdecTypedProp(prop_info *ZendPropertyInfo, var_ptr *Zval, copy *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var tmp Zval
	if copy == nil {
		copy = &tmp
	}
	ZVAL_COPY(copy, var_ptr)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		IncrementFunction(var_ptr)
	} else {
		DecrementFunction(var_ptr)
	}
	if var_ptr.IsType(IS_DOUBLE) && copy.IsType(IS_LONG) {
		var val ZendLong = ZendThrowIncdecPropError(prop_info, OPLINE_C)
		ZVAL_LONG(var_ptr, val)
	} else if ZendVerifyPropertyType(prop_info, var_ptr, EX_USES_STRICT_TYPES()) == 0 {
		ZvalPtrDtor(var_ptr)
		ZVAL_COPY_VALUE(var_ptr, copy)
		ZVAL_UNDEF(copy)
	} else if copy == &tmp {
		ZvalPtrDtor(&tmp)
	}
}
func ZendPreIncdecPropertyZval(prop *Zval, prop_info *ZendPropertyInfo, opline *ZendOp, _ EXECUTE_DATA_D) {
	if prop.IsType(IS_LONG) {
		if ZEND_IS_INCREMENT(opline.GetOpcode()) {
			FastLongIncrementFunction(prop)
		} else {
			FastLongDecrementFunction(prop)
		}
		if prop.GetType() != IS_LONG && prop_info != nil {
			var val ZendLong = ZendThrowIncdecPropError(prop_info, OPLINE_C)
			ZVAL_LONG(prop, val)
		}
	} else {
		for {
			if Z_ISREF_P(prop) {
				var ref *ZendReference = prop.GetRef()
				prop = Z_REFVAL_P(prop)
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					ZendIncdecTypedRef(ref, nil, OPLINE_C, EXECUTE_DATA_C)
					break
				}
			}
			if prop_info != nil {
				ZendIncdecTypedProp(prop_info, prop, nil, OPLINE_C, EXECUTE_DATA_C)
			} else if ZEND_IS_INCREMENT(opline.GetOpcode()) {
				IncrementFunction(prop)
			} else {
				DecrementFunction(prop)
			}
			break
		}
	}
	if RETURN_VALUE_USED(opline) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), prop)
	}
}
func ZendPostIncdecPropertyZval(prop *Zval, prop_info *ZendPropertyInfo, opline *ZendOp, _ EXECUTE_DATA_D) {
	if prop.IsType(IS_LONG) {
		ZVAL_LONG(EX_VAR(opline.GetResult().GetVar()), prop.GetLval())
		if ZEND_IS_INCREMENT(opline.GetOpcode()) {
			FastLongIncrementFunction(prop)
		} else {
			FastLongDecrementFunction(prop)
		}
		if prop.GetType() != IS_LONG && prop_info != nil {
			var val ZendLong = ZendThrowIncdecPropError(prop_info, OPLINE_C)
			ZVAL_LONG(prop, val)
		}
	} else {
		if Z_ISREF_P(prop) {
			var ref *ZendReference = prop.GetRef()
			prop = Z_REFVAL_P(prop)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendIncdecTypedRef(ref, EX_VAR(opline.GetResult().GetVar()), OPLINE_C, EXECUTE_DATA_C)
				return
			}
		}
		if prop_info != nil {
			ZendIncdecTypedProp(prop_info, prop, EX_VAR(opline.GetResult().GetVar()), OPLINE_C, EXECUTE_DATA_C)
		} else {
			ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), prop)
			if ZEND_IS_INCREMENT(opline.GetOpcode()) {
				IncrementFunction(prop)
			} else {
				DecrementFunction(prop)
			}
		}
	}
}
func ZendPostIncdecOverloadedProperty(object *Zval, property *Zval, cache_slot *any, opline *ZendOp, _ EXECUTE_DATA_D) {
	var rv Zval
	var obj Zval
	var z *Zval
	var z_copy Zval
	ZVAL_OBJ(&obj, object.GetObj())
	Z_ADDREF(obj)
	z = Z_OBJ_HT(obj).GetReadProperty()(&obj, property, BP_VAR_R, cache_slot, &rv)
	if ExecutorGlobals.GetException() != nil {
		OBJ_RELEASE(obj.GetObj())
		ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
		return
	}
	if z.IsType(IS_OBJECT) && Z_OBJ_HT_P(z).GetGet() != nil {
		var rv2 Zval
		var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		ZVAL_COPY_VALUE(z, value)
	}
	ZVAL_COPY_DEREF(&z_copy, z)
	ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &z_copy)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		IncrementFunction(&z_copy)
	} else {
		DecrementFunction(&z_copy)
	}
	Z_OBJ_HT(obj).GetWriteProperty()(&obj, property, &z_copy, cache_slot)
	OBJ_RELEASE(obj.GetObj())
	ZvalPtrDtor(&z_copy)
	ZvalPtrDtor(z)
}
func ZendPreIncdecOverloadedProperty(object *Zval, property *Zval, cache_slot *any, opline *ZendOp, _ EXECUTE_DATA_D) {
	var rv Zval
	var z *Zval
	var obj Zval
	var z_copy Zval
	ZVAL_OBJ(&obj, object.GetObj())
	Z_ADDREF(obj)
	z = Z_OBJ_HT(obj).GetReadProperty()(&obj, property, BP_VAR_R, cache_slot, &rv)
	if ExecutorGlobals.GetException() != nil {
		OBJ_RELEASE(obj.GetObj())
		if RETURN_VALUE_USED(opline) {
			ZVAL_NULL(EX_VAR(opline.GetResult().GetVar()))
		}
		return
	}
	if z.IsType(IS_OBJECT) && Z_OBJ_HT_P(z).GetGet() != nil {
		var rv2 Zval
		var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		ZVAL_COPY_VALUE(z, value)
	}
	ZVAL_COPY_DEREF(&z_copy, z)
	if ZEND_IS_INCREMENT(opline.GetOpcode()) {
		IncrementFunction(&z_copy)
	} else {
		DecrementFunction(&z_copy)
	}
	if RETURN_VALUE_USED(opline) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &z_copy)
	}
	Z_OBJ_HT(obj).GetWriteProperty()(&obj, property, &z_copy, cache_slot)
	OBJ_RELEASE(obj.GetObj())
	ZvalPtrDtor(&z_copy)
	ZvalPtrDtor(z)
}
func ZendAssignOpOverloadedProperty(object *Zval, property *Zval, cache_slot *any, value *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var z *Zval
	var rv Zval
	var obj Zval
	var res Zval
	ZVAL_OBJ(&obj, object.GetObj())
	Z_ADDREF(obj)
	z = Z_OBJ_HT(obj).GetReadProperty()(&obj, property, BP_VAR_R, cache_slot, &rv)
	if ExecutorGlobals.GetException() != nil {
		OBJ_RELEASE(obj.GetObj())
		if RETURN_VALUE_USED(opline) {
			ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
		}
		return
	}
	if z.IsType(IS_OBJECT) && Z_OBJ_HT_P(z).GetGet() != nil {
		var rv2 Zval
		var value *Zval = Z_OBJ_HT_P(z).GetGet()(z, &rv2)
		if z == &rv {
			ZvalPtrDtor(&rv)
		}
		ZVAL_COPY_VALUE(z, value)
	}
	if ZendBinaryOp(&res, z, value, OPLINE_C) == SUCCESS {
		Z_OBJ_HT(obj).GetWriteProperty()(&obj, property, &res, cache_slot)
	}
	if RETURN_VALUE_USED(opline) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), &res)
	}
	ZvalPtrDtor(z)
	ZvalPtrDtor(&res)
	OBJ_RELEASE(obj.GetObj())
}
func ZendExtensionStatementHandler(extension *ZendExtension, frame *ZendExecuteData) {
	if extension.GetStatementHandler() != nil {
		extension.GetStatementHandler()(frame)
	}
}
func ZendExtensionFcallBeginHandler(extension *ZendExtension, frame *ZendExecuteData) {
	if extension.GetFcallBeginHandler() != nil {
		extension.GetFcallBeginHandler()(frame)
	}
}
func ZendExtensionFcallEndHandler(extension *ZendExtension, frame *ZendExecuteData) {
	if extension.GetFcallEndHandler() != nil {
		extension.GetFcallEndHandler()(frame)
	}
}
func ZendGetTargetSymbolTable(fetch_type int, _ EXECUTE_DATA_D) *HashTable {
	var ht *HashTable
	if (fetch_type & (ZEND_FETCH_GLOBAL_LOCK | ZEND_FETCH_GLOBAL)) != 0 {
		ht = &(ExecutorGlobals.GetSymbolTable())
	} else {
		ZEND_ASSERT((fetch_type & ZEND_FETCH_LOCAL) != 0)
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			ZendRebuildSymbolTable()
		}
		ht = EX(symbol_table)
	}
	return ht
}
func ZendUndefinedOffset(lval ZendLong) {
	ZendError(E_NOTICE, "Undefined offset: "+ZEND_LONG_FMT, lval)
}
func ZendUndefinedIndex(offset *ZendString) {
	ZendError(E_NOTICE, "Undefined index: %s", offset.GetVal())
}
func ZendUndefinedOffsetWrite(ht *HashTable, lval ZendLong) int {
	/* The array may be destroyed while throwing the notice.
	 * Temporarily increase the refcount to detect this situation. */

	if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
		ht.IncGcRefcount()
	}
	ZendUndefinedOffset(lval)
	if (ht.GetGcFlags()&IS_ARRAY_IMMUTABLE) == 0 && ht.DecGcRefcount() == 0 {
		ZendArrayDestroy(ht)
		return FAILURE
	}
	if ExecutorGlobals.GetException() != nil {
		return FAILURE
	}
	return SUCCESS
}
func ZendUndefinedIndexWrite(ht *HashTable, offset *ZendString) int {
	/* The array may be destroyed while throwing the notice.
	 * Temporarily increase the refcount to detect this situation. */

	if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
		ht.IncGcRefcount()
	}
	ZendUndefinedIndex(offset)
	if (ht.GetGcFlags()&IS_ARRAY_IMMUTABLE) == 0 && ht.DecGcRefcount() == 0 {
		ZendArrayDestroy(ht)
		return FAILURE
	}
	if ExecutorGlobals.GetException() != nil {
		return FAILURE
	}
	return SUCCESS
}
func ZendUndefinedMethod(ce *ZendClassEntry, method *ZendString) {
	ZendThrowError(nil, "Call to undefined method %s::%s()", ce.GetName().GetVal(), method.GetVal())
}
func ZendInvalidMethodCall(object *Zval, function_name *Zval) {
	ZendThrowError(nil, "Call to a member function %s() on %s", Z_STRVAL_P(function_name), ZendGetTypeByConst(object.GetType()))
}
func ZendNonStaticMethodCall(fbc *ZendFunction) {
	if fbc.IsAllowStatic() {
		ZendError(E_DEPRECATED, "Non-static method %s::%s() should not be called statically", fbc.GetScope().GetName().GetVal(), fbc.GetFunctionName().GetVal())
	} else {
		ZendThrowError(ZendCeError, "Non-static method %s::%s() cannot be called statically", fbc.GetScope().GetName().GetVal(), fbc.GetFunctionName().GetVal())
	}
}
func ZendParamMustBeRef(func_ *ZendFunction, arg_num uint32) {
	ZendError(E_WARNING, "Parameter %d to %s%s%s() expected to be a reference, value given", arg_num, b.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().GetName().GetVal() }, ""), b.Cond(func_.GetScope() != nil, "::", ""), func_.GetFunctionName().GetVal())
}
func ZendUseScalarAsArray() {
	ZendError(E_WARNING, "Cannot use a scalar value as an array")
}
func ZendCannotAddElement() {
	ZendError(E_WARNING, "Cannot add element to the array as the next element is already occupied")
}
func ZendUseResourceAsOffset(dim *Zval) {
	ZendError(E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", Z_RES_HANDLE_P(dim), Z_RES_HANDLE_P(dim))
}
func ZendUseNewElementForString() {
	ZendThrowError(nil, "[] operator not supported for strings")
}
func ZendBinaryAssignOpDimSlow(container *Zval, dim *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	if container.IsType(IS_STRING) {
		if opline.GetOp2Type() == IS_UNUSED {
			ZendUseNewElementForString()
		} else {
			ZendCheckStringOffset(dim, BP_VAR_RW, EXECUTE_DATA_C)
			ZendWrongStringOffset(EXECUTE_DATA_C)
		}
	} else if !(Z_ISERROR_P(container)) {
		ZendUseScalarAsArray()
	}
}
func SlowIndexConvert(ht *HashTable, dim *Zval, value *ZendValue, _ EXECUTE_DATA_D) ZendUchar {
	switch dim.GetType() {
	case IS_UNDEF:

		/* The array may be destroyed while throwing the notice.
		 * Temporarily increase the refcount to detect this situation. */

		if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
			ht.IncGcRefcount()
		}
		ZVAL_UNDEFINED_OP2()
		if (ht.GetGcFlags()&IS_ARRAY_IMMUTABLE) == 0 && ht.DecGcRefcount() == 0 {
			ZendArrayDestroy(ht)
			return IS_NULL
		}
		if ExecutorGlobals.GetException() != nil {
			return IS_NULL
		}
	case IS_NULL:
		value.SetStr(ZSTR_EMPTY_ALLOC())
		return IS_STRING
	case IS_DOUBLE:
		value.SetLval(ZendDvalToLval(dim.GetDval()))
		return IS_LONG
	case IS_RESOURCE:
		ZendUseResourceAsOffset(dim)
		value.SetLval(Z_RES_HANDLE_P(dim))
		return IS_LONG
	case IS_FALSE:
		value.SetLval(0)
		return IS_LONG
	case IS_TRUE:
		value.SetLval(1)
		return IS_LONG
	default:
		ZendIllegalOffset()
		return IS_NULL
	}
}
func ZendFetchDimensionAddressInner(ht *HashTable, dim *Zval, dim_type int, type_ int, _ EXECUTE_DATA_D) *Zval {
	var retval *Zval = nil
	var offset_key *ZendString
	var hval ZendUlong
try_again:
	if dim.IsType(IS_LONG) {
		hval = dim.GetLval()
	num_index:
		ZEND_HASH_INDEX_FIND(ht, hval, retval, num_undef)
		return retval
	num_undef:
		switch type_ {
		case BP_VAR_R:
			ZendUndefinedOffset(hval)
		case BP_VAR_UNSET:

		case BP_VAR_IS:
			retval = &(ExecutorGlobals.GetUninitializedZval())
			break
		case BP_VAR_RW:
			if ZendUndefinedOffsetWrite(ht, hval) == FAILURE {
				return nil
			}
		case BP_VAR_W:
			retval = ht.IndexAddNew(hval, &(ExecutorGlobals.GetUninitializedZval()))
			break
		}
	} else if dim.IsType(IS_STRING) {
		offset_key = dim.GetStr()
		if ZEND_CONST_COND(dim_type != IS_CONST, 1) {
			if ZEND_HANDLE_NUMERIC(offset_key, hval) != 0 {
				goto num_index
			}
		}
	str_index:
		retval = ht.FindEx(offset_key, ZEND_CONST_COND(dim_type == IS_CONST, 0))
		if retval != nil {

			/* support for $GLOBALS[...] */

			if retval.IsType(IS_INDIRECT) {
				retval = retval.GetZv()
				if retval.IsType(IS_UNDEF) {
					switch type_ {
					case BP_VAR_R:
						ZendUndefinedIndex(offset_key)
					case BP_VAR_UNSET:

					case BP_VAR_IS:
						retval = &(ExecutorGlobals.GetUninitializedZval())
						break
					case BP_VAR_RW:
						if ZendUndefinedIndexWrite(ht, offset_key) != 0 {
							return nil
						}
					case BP_VAR_W:
						ZVAL_NULL(retval)
						break
					}
				}
			}

			/* support for $GLOBALS[...] */

		} else {
			switch type_ {
			case BP_VAR_R:
				ZendUndefinedIndex(offset_key)
			case BP_VAR_UNSET:

			case BP_VAR_IS:
				retval = &(ExecutorGlobals.GetUninitializedZval())
				break
			case BP_VAR_RW:

				/* Key may be released while throwing the undefined index warning. */

				offset_key.IncGcRefcount()
				if ZendUndefinedIndexWrite(ht, offset_key) == FAILURE {
					ZendStringRelease(offset_key)
					return nil
				}
				retval = ht.AddNew(offset_key, &(ExecutorGlobals.GetUninitializedZval()))
				ZendStringRelease(offset_key)
				break
			case BP_VAR_W:
				retval = ht.AddNew(offset_key, &(ExecutorGlobals.GetUninitializedZval()))
				break
			}
		}
	} else if dim.IsType(IS_REFERENCE) {
		dim = Z_REFVAL_P(dim)
		goto try_again
	} else {
		var val ZendValue
		var t ZendUchar = SlowIndexConvert(ht, dim, &val, EXECUTE_DATA_C)
		if t == IS_STRING {
			offset_key = val.GetStr()
			goto str_index
		} else if t == IS_LONG {
			hval = val.GetLval()
			goto num_index
		} else {
			if type_ == BP_VAR_W || type_ == BP_VAR_RW {
				retval = nil
			} else {
				retval = &(ExecutorGlobals.GetUninitializedZval())
			}
		}
	}
	return retval
}
func zend_fetch_dimension_address_inner_W(ht *HashTable, dim *Zval, _ EXECUTE_DATA_D) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_TMP_VAR, BP_VAR_W, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_inner_W_CONST(ht *HashTable, dim *Zval, _ EXECUTE_DATA_D) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_CONST, BP_VAR_W, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_inner_RW(ht *HashTable, dim *Zval, _ EXECUTE_DATA_D) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_TMP_VAR, BP_VAR_RW, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_inner_RW_CONST(ht *HashTable, dim *Zval, _ EXECUTE_DATA_D) *Zval {
	return ZendFetchDimensionAddressInner(ht, dim, IS_CONST, BP_VAR_RW, EXECUTE_DATA_C)
}
func ZendFetchDimensionAddress(result *Zval, container *Zval, dim *Zval, dim_type int, type_ int, _ EXECUTE_DATA_D) {
	var retval *Zval
	if container.IsType(IS_ARRAY) {
	try_array:
		SEPARATE_ARRAY(container)
	fetch_from_array:
		if dim == nil {
			retval = container.GetArr().NextIndexInsert(&(ExecutorGlobals.GetUninitializedZval()))
			if retval == nil {
				ZendCannotAddElement()
				ZVAL_ERROR(result)
				return
			}
		} else {
			retval = ZendFetchDimensionAddressInner(container.GetArr(), dim, dim_type, type_, EXECUTE_DATA_C)
			if retval == nil {
				ZVAL_ERROR(result)
				return
			}
		}
		ZVAL_INDIRECT(result, retval)
		return
	} else if container.IsType(IS_REFERENCE) {
		var ref *ZendReference = container.GetRef()
		container = Z_REFVAL_P(container)
		if container.IsType(IS_ARRAY) {
			goto try_array
		} else if container.GetType() <= IS_FALSE {
			if type_ != BP_VAR_UNSET {
				if ZEND_REF_HAS_TYPE_SOURCES(ref) {
					if ZendVerifyRefArrayAssignable(ref) == 0 {
						ZVAL_ERROR(result)
						return
					}
				}
				ArrayInit(container)
				goto fetch_from_array
			} else {
				goto return_null
			}
		}
	}
	if container.IsType(IS_STRING) {
		if dim == nil {
			ZendUseNewElementForString()
		} else {
			ZendCheckStringOffset(dim, type_, EXECUTE_DATA_C)
			ZendWrongStringOffset(EXECUTE_DATA_C)
		}
		ZVAL_ERROR(result)
	} else if container.IsType(IS_OBJECT) {
		if ZEND_CONST_COND(dim_type == IS_CV, dim != nil) && dim.IsType(IS_UNDEF) {
			dim = ZVAL_UNDEFINED_OP2()
		}
		if dim_type == IS_CONST && dim.GetU2Extra() == ZEND_EXTRA_VALUE {
			dim++
		}
		retval = Z_OBJ_HT_P(container).GetReadDimension()(container, dim, type_, result)
		if retval == &(ExecutorGlobals.GetUninitializedZval()) {
			var ce *ZendClassEntry = Z_OBJCE_P(container)
			ZVAL_NULL(result)
			ZendError(E_NOTICE, "Indirect modification of overloaded element of %s has no effect", ce.GetName().GetVal())
		} else if retval != nil && retval.GetType() != IS_UNDEF {
			if !(Z_ISREF_P(retval)) {
				if result != retval {
					ZVAL_COPY(result, retval)
					retval = result
				}
				if retval.GetType() != IS_OBJECT {
					var ce *ZendClassEntry = Z_OBJCE_P(container)
					ZendError(E_NOTICE, "Indirect modification of overloaded element of %s has no effect", ce.GetName().GetVal())
				}
			} else if Z_REFCOUNT_P(retval) == 1 {
				ZVAL_UNREF(retval)
			}
			if result != retval {
				ZVAL_INDIRECT(result, retval)
			}
		} else {
			ZVAL_ERROR(result)
		}
	} else {
		if container.GetType() <= IS_FALSE {
			if type_ != BP_VAR_W && container.IsType(IS_UNDEF) {
				ZVAL_UNDEFINED_OP1()
			}
			if type_ != BP_VAR_UNSET {
				ArrayInit(container)
				goto fetch_from_array
			} else {
			return_null:

				/* for read-mode only */

				if ZEND_CONST_COND(dim_type == IS_CV, dim != nil) && dim.IsType(IS_UNDEF) {
					ZVAL_UNDEFINED_OP2()
				}
				ZVAL_NULL(result)
			}
		} else if Z_ISERROR_P(container) {
			ZVAL_ERROR(result)
		} else {
			if type_ == BP_VAR_UNSET {
				ZendError(E_WARNING, "Cannot unset offset in a non-array variable")
				ZVAL_NULL(result)
			} else {
				ZendUseScalarAsArray()
				ZVAL_ERROR(result)
			}
		}
	}
}
func zend_fetch_dimension_address_W(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_W, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_RW(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_RW, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_UNSET(container_ptr *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddress(result, container_ptr, dim, dim_type, BP_VAR_UNSET, EXECUTE_DATA_C)
}
func ZendFetchDimensionAddressRead(result *Zval, container *Zval, dim *Zval, dim_type int, type_ int, is_list int, slow int, _ EXECUTE_DATA_D) {
	var retval *Zval
	if slow == 0 {
		if container.IsType(IS_ARRAY) {
		try_array:
			retval = ZendFetchDimensionAddressInner(container.GetArr(), dim, dim_type, type_, EXECUTE_DATA_C)
			ZVAL_COPY_DEREF(result, retval)
			return
		} else if container.IsType(IS_REFERENCE) {
			container = Z_REFVAL_P(container)
			if container.IsType(IS_ARRAY) {
				goto try_array
			}
		}
	}
	if is_list == 0 && container.IsType(IS_STRING) {
		var offset ZendLong
	try_string_offset:
		if dim.GetType() != IS_LONG {
			switch dim.GetType() {
			case IS_STRING:
				if IS_LONG == IsNumericString(Z_STRVAL_P(dim), Z_STRLEN_P(dim), nil, nil, -1) {
					break
				}
				if type_ == BP_VAR_IS {
					ZVAL_NULL(result)
					return
				}
				ZendError(E_WARNING, "Illegal string offset '%s'", Z_STRVAL_P(dim))
				break
			case IS_UNDEF:
				ZVAL_UNDEFINED_OP2()
			case IS_DOUBLE:

			case IS_NULL:

			case IS_FALSE:

			case IS_TRUE:
				if type_ != BP_VAR_IS {
					ZendError(E_NOTICE, "String offset cast occurred")
				}
				break
			case IS_REFERENCE:
				dim = Z_REFVAL_P(dim)
				goto try_string_offset
			default:
				ZendIllegalOffset()
				break
			}
			offset = ZvalGetLongFunc(dim)
		} else {
			offset = dim.GetLval()
		}
		if Z_STRLEN_P(container) < b.CondF(offset < 0, func() int { return -int(offset) }, func() int { return int(offset + 1) }) {
			if type_ != BP_VAR_IS {
				ZendError(E_NOTICE, "Uninitialized string offset: "+ZEND_LONG_FMT, offset)
				ZVAL_EMPTY_STRING(result)
			} else {
				ZVAL_NULL(result)
			}
		} else {
			var c ZendUchar
			var real_offset ZendLong
			if offset < 0 {
				real_offset = ZendLong(Z_STRLEN_P(container) + offset)
			} else {
				real_offset = offset
			}
			c = ZendUchar(Z_STRVAL_P(container)[real_offset])
			ZVAL_INTERNED_STR(result, ZSTR_CHAR(c))
		}
	} else if container.IsType(IS_OBJECT) {
		if ZEND_CONST_COND(dim_type == IS_CV, 1) && dim.IsType(IS_UNDEF) {
			dim = ZVAL_UNDEFINED_OP2()
		}
		if dim_type == IS_CONST && dim.GetU2Extra() == ZEND_EXTRA_VALUE {
			dim++
		}
		retval = Z_OBJ_HT_P(container).GetReadDimension()(container, dim, type_, result)
		ZEND_ASSERT(result != nil)
		if retval != nil {
			if result != retval {
				ZVAL_COPY_DEREF(result, retval)
			} else if Z_ISREF_P(retval) {
				ZendUnwrapReference(result)
			}
		} else {
			ZVAL_NULL(result)
		}
	} else {
		if type_ != BP_VAR_IS && container.IsType(IS_UNDEF) {
			container = ZVAL_UNDEFINED_OP1()
		}
		if ZEND_CONST_COND(dim_type == IS_CV, 1) && dim.IsType(IS_UNDEF) {
			ZVAL_UNDEFINED_OP2()
		}
		if is_list == 0 && type_ != BP_VAR_IS {
			ZendError(E_NOTICE, "Trying to access array offset on value of type %s", ZendZvalTypeName(container))
		}
		ZVAL_NULL(result)
	}
}
func zend_fetch_dimension_address_read_R(container *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_R, 0, 0, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_read_R_slow(container *Zval, dim *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, IS_CV, BP_VAR_R, 0, 1, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_read_IS(container *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_IS, 0, 0, EXECUTE_DATA_C)
}
func zend_fetch_dimension_address_LIST_r(container *Zval, dim *Zval, dim_type int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	ZendFetchDimensionAddressRead(result, container, dim, dim_type, BP_VAR_R, 1, 0, EXECUTE_DATA_C)
}
func ZendFetchDimensionConst(result *Zval, container *Zval, dim *Zval, type_ int) {
	ZendFetchDimensionAddressRead(result, container, dim, IS_TMP_VAR, type_, 0, 0, nil)
}
func ZendFindArrayDimSlow(ht *HashTable, offset *Zval, _ EXECUTE_DATA_D) *Zval {
	var hval ZendUlong
	if offset.IsType(IS_DOUBLE) {
		hval = ZendDvalToLval(offset.GetDval())
	num_idx:
		return ht.IndexFind(hval)
	} else if offset.IsType(IS_NULL) {
	str_idx:
		return ht.FindExInd(ZSTR_EMPTY_ALLOC(), 1)
	} else if offset.IsType(IS_FALSE) {
		hval = 0
		goto num_idx
	} else if offset.IsType(IS_TRUE) {
		hval = 1
		goto num_idx
	} else if offset.IsType(IS_RESOURCE) {
		hval = Z_RES_HANDLE_P(offset)
		goto num_idx
	} else if offset.IsType(IS_UNDEF) {
		ZVAL_UNDEFINED_OP2()
		goto str_idx
	} else {
		ZendError(E_WARNING, "Illegal offset type in isset or empty")
		return nil
	}
}
func ZendIssetDimSlow(container *Zval, offset *Zval, _ EXECUTE_DATA_D) int {
	if offset.IsType(IS_UNDEF) {
		offset = ZVAL_UNDEFINED_OP2()
	}
	if container.IsType(IS_OBJECT) {
		return Z_OBJ_HT_P(container).GetHasDimension()(container, offset, 0)
	} else if container.IsType(IS_STRING) {
		var lval ZendLong
		if offset.IsType(IS_LONG) {
			lval = offset.GetLval()
		str_offset:
			if lval < 0 {
				lval += ZendLong(Z_STRLEN_P(container))
			}
			if lval >= 0 && int(lval < Z_STRLEN_P(container)) != 0 {
				return 1
			} else {
				return 0
			}
		} else {

			/*if (OP2_TYPE & (IS_CV|IS_VAR)) {*/

			ZVAL_DEREF(offset)

			/*}*/

			if offset.GetType() < IS_STRING || offset.IsType(IS_STRING) && IS_LONG == IsNumericString(Z_STRVAL_P(offset), Z_STRLEN_P(offset), nil, nil, 0) {
				lval = ZvalGetLong(offset)
				goto str_offset
			}
			return 0
		}
	} else {
		return 0
	}
}
func ZendIsemptyDimSlow(container *Zval, offset *Zval, _ EXECUTE_DATA_D) int {
	if offset.IsType(IS_UNDEF) {
		offset = ZVAL_UNDEFINED_OP2()
	}
	if container.IsType(IS_OBJECT) {
		return !(Z_OBJ_HT_P(container).GetHasDimension()(container, offset, 1))
	} else if container.IsType(IS_STRING) {
		var lval ZendLong
		if offset.IsType(IS_LONG) {
			lval = offset.GetLval()
		str_offset:
			if lval < 0 {
				lval += ZendLong(Z_STRLEN_P(container))
			}
			if lval >= 0 && int(lval < Z_STRLEN_P(container)) != 0 {
				return Z_STRVAL_P(container)[lval] == '0'
			} else {
				return 1
			}
		} else {

			/*if (OP2_TYPE & (IS_CV|IS_VAR)) {*/

			ZVAL_DEREF(offset)

			/*}*/

			if offset.GetType() < IS_STRING || offset.IsType(IS_STRING) && IS_LONG == IsNumericString(Z_STRVAL_P(offset), Z_STRLEN_P(offset), nil, nil, 0) {
				lval = ZvalGetLong(offset)
				goto str_offset
			}
			return 1
		}
	} else {
		return 1
	}
}
func ZendArrayKeyExistsFast(ht *HashTable, key *Zval, opline *ZendOp, _ EXECUTE_DATA_D) uint32 {
	var str *ZendString
	var hval ZendUlong
try_again:
	if key.IsType(IS_STRING) {
		str = key.GetStr()
		if ZEND_HANDLE_NUMERIC(str, hval) != 0 {
			goto num_key
		}
	str_key:
		if ht.FindInd(str) != nil {
			return IS_TRUE
		} else {
			return IS_FALSE
		}
	} else if key.IsType(IS_LONG) {
		hval = key.GetLval()
	num_key:
		if ht.IndexFind(hval) != nil {
			return IS_TRUE
		} else {
			return IS_FALSE
		}
	} else if Z_ISREF_P(key) {
		key = Z_REFVAL_P(key)
		goto try_again
	} else if key.GetType() <= IS_NULL {
		if key.IsType(IS_UNDEF) {
			ZVAL_UNDEFINED_OP1()
		}
		str = ZSTR_EMPTY_ALLOC()
		goto str_key
	} else {
		ZendError(E_WARNING, "array_key_exists(): The first argument should be either a string or an integer")
		return IS_FALSE
	}
}
func ZendArrayKeyExistsSlow(subject *Zval, key *Zval, opline *ZendOp, _ EXECUTE_DATA_D) uint32 {
	if subject.IsType(IS_OBJECT) {
		ZendError(E_DEPRECATED, "array_key_exists(): "+"Using array_key_exists() on objects is deprecated. "+"Use isset() or property_exists() instead")
		var ht *HashTable = ZendGetPropertiesFor(subject, ZEND_PROP_PURPOSE_ARRAY_CAST)
		var result uint32 = ZendArrayKeyExistsFast(ht, key, OPLINE_C, EXECUTE_DATA_C)
		ZendReleaseProperties(ht)
		return result
	} else {
		if key.IsType(IS_UNDEF) {
			ZVAL_UNDEFINED_OP1()
		}
		if subject.GetTypeInfo() == IS_UNDEF {
			ZVAL_UNDEFINED_OP2()
		}
		ZendInternalTypeError(EX_USES_STRICT_TYPES(), "array_key_exists() expects parameter 2 to be array, %s given", ZendGetTypeByConst(subject.GetType()))
		return IS_NULL
	}
}
func PromotesToArray(val *Zval) ZendBool {
	return val.GetType() <= IS_FALSE || Z_ISREF_P(val) && Z_REFVAL_P(val).GetType() <= IS_FALSE
}
func PromotesToObject(val *Zval) ZendBool {
	ZVAL_DEREF(val)
	return val.GetType() <= IS_FALSE || val.IsType(IS_STRING) && Z_STRLEN_P(val) == 0
}
func CheckTypeArrayAssignable(type_ ZendType) ZendBool {
	if type_ == 0 {
		return 1
	}
	return type_.IsCode() && (type_.Code() == IS_ARRAY || type_.Code() == IS_ITERABLE)
}
func check_type_stdClass_assignable(type_ ZendType) ZendBool {
	if type_ == 0 {
		return 1
	}
	if type_.IsClass() {
		if type_.IsCe() {
			return type_.Ce() == ZendStandardClassDef
		} else {
			return ZendStringEqualsLiteralCi(type_.Name(), "stdclass")
		}
	} else {
		return type_.Code() == IS_OBJECT
	}
}
func ZendVerifyRefArrayAssignable(ref *ZendReference) ZendBool {
	var prop *ZendPropertyInfo
	ZEND_ASSERT(ZEND_REF_HAS_TYPE_SOURCES(ref))
	var _source_list *ZendPropertyInfoSourceList = &(ref.GetSources())
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
			prop = *_prop
			if CheckTypeArrayAssignable(prop.GetType()) == 0 {
				ZendThrowAutoInitInRefError(prop, "array")
				return 0
			}
		}
	}
	return 1
}
func zend_verify_ref_stdClass_assignable(ref *ZendReference) ZendBool {
	var prop *ZendPropertyInfo
	ZEND_ASSERT(ZEND_REF_HAS_TYPE_SOURCES(ref))
	var _source_list *ZendPropertyInfoSourceList = &(ref.GetSources())
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
			prop = *_prop
			if check_type_stdClass_assignable(prop.GetType()) == 0 {
				ZendThrowAutoInitInRefError(prop, "stdClass")
				return 0
			}
		}
	}
	return 1
}
func ZendObjectFetchPropertyTypeInfo(obj *ZendObject, slot *Zval) *ZendPropertyInfo {
	if !(ZEND_CLASS_HAS_TYPE_HINTS(obj.GetCe())) {
		return nil
	}

	/* Not a declared property */

	if slot < obj.GetPropertiesTable() || slot >= obj.GetPropertiesTable()+obj.GetCe().GetDefaultPropertiesCount() {
		return nil
	}
	return ZendGetTypedPropertyInfoForSlot(obj, slot)
}
func ZendHandleFetchObjFlags(result *Zval, ptr *Zval, obj *ZendObject, prop_info *ZendPropertyInfo, flags uint32) ZendBool {
	switch flags {
	case ZEND_FETCH_DIM_WRITE:
		if PromotesToArray(ptr) != 0 {
			if prop_info == nil {
				prop_info = ZendObjectFetchPropertyTypeInfo(obj, ptr)
				if prop_info == nil {
					break
				}
			}
			if CheckTypeArrayAssignable(prop_info.GetType()) == 0 {
				ZendThrowAutoInitInPropError(prop_info, "array")
				if result != nil {
					ZVAL_ERROR(result)
				}
				return 0
			}
		}
		break
	case ZEND_FETCH_OBJ_WRITE:
		if PromotesToObject(ptr) != 0 {
			if prop_info == nil {
				prop_info = ZendObjectFetchPropertyTypeInfo(obj, ptr)
				if prop_info == nil {
					break
				}
			}
			if check_type_stdClass_assignable(prop_info.GetType()) == 0 {
				ZendThrowAutoInitInPropError(prop_info, "stdClass")
				if result != nil {
					ZVAL_ERROR(result)
				}
				return 0
			}
		}
		break
	case ZEND_FETCH_REF:
		if ptr.GetType() != IS_REFERENCE {
			if prop_info == nil {
				prop_info = ZendObjectFetchPropertyTypeInfo(obj, ptr)
				if prop_info == nil {
					break
				}
			}
			if ptr.IsType(IS_UNDEF) {
				if !(prop_info.GetType().AllowNull()) {
					ZendThrowAccessUninitPropByRefError(prop_info)
					if result != nil {
						ZVAL_ERROR(result)
					}
					return 0
				}
				ZVAL_NULL(ptr)
			}
			ZVAL_NEW_REF(ptr, ptr)
			ZEND_REF_ADD_TYPE_SOURCE(ptr.GetRef(), prop_info)
		}
		break
	default:
		break
	}
	return 1
}
func ZendFetchPropertyAddress(result *Zval, container *Zval, container_op_type uint32, prop_ptr *Zval, prop_op_type uint32, cache_slot *any, type_ int, flags uint32, init_undef ZendBool, opline *ZendOp, _ EXECUTE_DATA_D) {
	var ptr *Zval
	if container_op_type != IS_UNUSED && container.GetType() != IS_OBJECT {
		for {
			if Z_ISREF_P(container) && Z_REFVAL_P(container).IsType(IS_OBJECT) {
				container = Z_REFVAL_P(container)
				break
			}
			if container_op_type == IS_CV && type_ != BP_VAR_W && container.IsType(IS_UNDEF) {
				ZVAL_UNDEFINED_OP1()
			}

			/* this should modify object only if it's empty */

			if type_ == BP_VAR_UNSET {
				ZVAL_NULL(result)
				return
			}
			container = MakeRealObject(container, prop_ptr, OPLINE_C, EXECUTE_DATA_C)
			if container == nil {
				ZVAL_ERROR(result)
				return
			}
			break
		}
	}
	if prop_op_type == IS_CONST && Z_OBJCE_P(container) == CACHED_PTR_EX(cache_slot) {
		var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
		var zobj *ZendObject = container.GetObj()
		if IS_VALID_PROPERTY_OFFSET(prop_offset) {
			ptr = OBJ_PROP(zobj, prop_offset)
			if ptr.GetType() != IS_UNDEF {
				ZVAL_INDIRECT(result, ptr)
				if flags != 0 {
					var prop_info *ZendPropertyInfo = CACHED_PTR_EX(cache_slot + 2)
					if prop_info != nil {
						ZendHandleFetchObjFlags(result, ptr, nil, prop_info, flags)
					}
				}
				return
			}
		} else if zobj.GetProperties() != nil {
			if zobj.GetProperties().GetGcRefcount() > 1 {
				if (zobj.GetProperties().GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
					zobj.GetProperties().DecGcRefcount()
				}
				zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
			}
			ptr = zobj.GetProperties().FindEx(prop_ptr.GetStr(), 1)
			if ptr != nil {
				ZVAL_INDIRECT(result, ptr)
				return
			}
		}
	}
	ptr = Z_OBJ_HT_P(container).GetGetPropertyPtrPtr()(container, prop_ptr, type_, cache_slot)
	if nil == ptr {
		ptr = Z_OBJ_HT_P(container).GetReadProperty()(container, prop_ptr, type_, cache_slot, result)
		if ptr == result {
			if Z_ISREF_P(ptr) && Z_REFCOUNT_P(ptr) == 1 {
				ZVAL_UNREF(ptr)
			}
			return
		}
		if ExecutorGlobals.GetException() != nil {
			ZVAL_ERROR(result)
			return
		}
	} else if Z_ISERROR_P(ptr) {
		ZVAL_ERROR(result)
		return
	}
	ZVAL_INDIRECT(result, ptr)
	if flags != 0 {
		var prop_info *ZendPropertyInfo
		if prop_op_type == IS_CONST {
			prop_info = CACHED_PTR_EX(cache_slot + 2)
			if prop_info != nil {
				if ZendHandleFetchObjFlags(result, ptr, nil, prop_info, flags) == 0 {
					return
				}
			}
		} else {
			if ZendHandleFetchObjFlags(result, ptr, container.GetObj(), nil, flags) == 0 {
				return
			}
		}
	}
	if init_undef != 0 && ptr.IsType(IS_UNDEF) {
		ZVAL_NULL(ptr)
	}
}
func ZendAssignToPropertyReference(container *Zval, container_op_type uint32, prop_ptr *Zval, prop_op_type uint32, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	var variable Zval
	var variable_ptr *Zval = &variable
	var cache_addr *any = b.CondF1(prop_op_type == IS_CONST, func() *any { return CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_RETURNS_FUNCTION) }, nil)
	ZendFetchPropertyAddress(variable_ptr, container, container_op_type, prop_ptr, prop_op_type, cache_addr, BP_VAR_W, 0, 0, OPLINE_C, EXECUTE_DATA_C)
	if variable_ptr.IsType(IS_INDIRECT) {
		variable_ptr = variable_ptr.GetZv()
	}
	if Z_ISERROR_P(variable_ptr) {
		variable_ptr = &(ExecutorGlobals.GetUninitializedZval())
	} else if variable.GetType() != IS_INDIRECT {
		ZendThrowError(nil, "Cannot assign by reference to overloaded object")
		ZvalPtrDtor(&variable)
		variable_ptr = &(ExecutorGlobals.GetUninitializedZval())
	} else if Z_ISERROR_P(value_ptr) {
		variable_ptr = &(ExecutorGlobals.GetUninitializedZval())
	} else if (opline.GetExtendedValue()&ZEND_RETURNS_FUNCTION) != 0 && !(Z_ISREF_P(value_ptr)) {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, OPLINE_C, EXECUTE_DATA_C)
	} else {
		var prop_info *ZendPropertyInfo = nil
		if prop_op_type == IS_CONST {
			prop_info = (*ZendPropertyInfo)(CACHED_PTR_EX(cache_addr + 2))
		} else {
			ZVAL_DEREF(container)
			prop_info = ZendObjectFetchPropertyTypeInfo(container.GetObj(), variable_ptr)
		}
		if prop_info != nil {
			variable_ptr = ZendAssignToTypedPropertyReference(prop_info, variable_ptr, value_ptr, EXECUTE_DATA_C)
		} else {
			ZendAssignToVariableReference(variable_ptr, value_ptr)
		}
	}
	if RETURN_VALUE_USED(opline) {
		ZVAL_COPY(EX_VAR(opline.GetResult().GetVar()), variable_ptr)
	}
}
func ZendAssignToPropertyReferenceThisConst(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_UNUSED, prop_ptr, IS_CONST, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendAssignToPropertyReferenceVarConst(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_VAR, prop_ptr, IS_CONST, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendAssignToPropertyReferenceThisVar(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_UNUSED, prop_ptr, IS_VAR, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendAssignToPropertyReferenceVarVar(container *Zval, prop_ptr *Zval, value_ptr *Zval, opline *ZendOp, _ EXECUTE_DATA_D) {
	ZendAssignToPropertyReference(container, IS_VAR, prop_ptr, IS_VAR, value_ptr, OPLINE_C, EXECUTE_DATA_C)
}
func ZendFetchStaticPropertyAddressEx(retval **Zval, prop_info **ZendPropertyInfo, cache_slot uint32, fetch_type int, opline *ZendOp, _ EXECUTE_DATA_D) int {
	var free_op1 ZendFreeOp
	var name *ZendString
	var tmp_name *ZendString
	var ce *ZendClassEntry
	var property_info *ZendPropertyInfo
	var op1_type ZendUchar = opline.GetOp1Type()
	var op2_type ZendUchar = opline.GetOp2Type()
	if op2_type == IS_CONST {
		var class_name *Zval = RT_CONSTANT(opline, opline.GetOp2())
		ZEND_ASSERT(op1_type != IS_CONST || CACHED_PTR(cache_slot) == nil)
		if b.Assign(&ce, CACHED_PTR(cache_slot)) == nil {
			ce = ZendFetchClassByName(class_name.GetStr(), (class_name + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				FREE_UNFETCHED_OP(op1_type, opline.GetOp1().GetVar())
				return FAILURE
			}
			if op1_type != IS_CONST {
				CACHE_PTR(cache_slot, ce)
			}
		}
	} else {
		if op2_type == IS_UNUSED {
			ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
			if ce == nil {
				FREE_UNFETCHED_OP(op1_type, opline.GetOp1().GetVar())
				return FAILURE
			}
		} else {
			ce = EX_VAR(opline.GetOp2().GetVar()).GetCe()
		}
		if op1_type == IS_CONST && CACHED_PTR(cache_slot) == ce {
			*retval = CACHED_PTR(cache_slot + b.SizeOf("void *"))
			*prop_info = CACHED_PTR(cache_slot + b.SizeOf("void *")*2)
			return SUCCESS
		}
	}
	if op1_type == IS_CONST {
		name = RT_CONSTANT(opline, opline.GetOp1()).GetStr()
	} else {
		var varname *Zval = GetZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
		if varname.IsType(IS_STRING) {
			name = varname.GetStr()
			tmp_name = nil
		} else {
			if op1_type == IS_CV && varname.IsType(IS_UNDEF) {
				ZvalUndefinedCv(opline.GetOp1().GetVar(), EXECUTE_DATA_C)
			}
			name = ZvalGetTmpString(varname, &tmp_name)
		}
	}
	*retval = ZendStdGetStaticPropertyWithInfo(ce, name, fetch_type, &property_info)
	if op1_type != IS_CONST {
		ZendTmpStringRelease(tmp_name)
		if op1_type != IS_CV {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	if (*retval) == nil {
		return FAILURE
	}
	*prop_info = property_info
	if op1_type == IS_CONST {
		CACHE_POLYMORPHIC_PTR(cache_slot, ce, *retval)
		CACHE_PTR(cache_slot+b.SizeOf("void *")*2, property_info)
	}
	return SUCCESS
}
func ZendFetchStaticPropertyAddress(retval **Zval, prop_info **ZendPropertyInfo, cache_slot uint32, fetch_type int, flags int, opline *ZendOp, _ EXECUTE_DATA_D) int {
	var success int
	var property_info *ZendPropertyInfo
	if opline.GetOp1Type() == IS_CONST && (opline.GetOp2Type() == IS_CONST || opline.GetOp2Type() == IS_UNUSED && (opline.GetOp2().GetNum() == ZEND_FETCH_CLASS_SELF || opline.GetOp2().GetNum() == ZEND_FETCH_CLASS_PARENT)) && CACHED_PTR(cache_slot) != nil {
		*retval = CACHED_PTR(cache_slot + b.SizeOf("void *"))
		property_info = CACHED_PTR(cache_slot + b.SizeOf("void *")*2)
		if (fetch_type == BP_VAR_R || fetch_type == BP_VAR_RW) && retval.IsType(IS_UNDEF) && property_info.GetType() != 0 {
			ZendThrowError(nil, "Typed static property %s::$%s must not be accessed before initialization", property_info.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(property_info.GetName()))
			return FAILURE
		}
	} else {
		success = ZendFetchStaticPropertyAddressEx(retval, &property_info, cache_slot, fetch_type, OPLINE_C, EXECUTE_DATA_C)
		if success != SUCCESS {
			return FAILURE
		}
	}
	if flags != 0 && property_info.GetType() != 0 {
		ZendHandleFetchObjFlags(nil, *retval, nil, property_info, flags)
	}
	if prop_info != nil {
		*prop_info = property_info
	}
	return SUCCESS
}
func ZendThrowRefTypeErrorType(prop1 *ZendPropertyInfo, prop2 *ZendPropertyInfo, zv *Zval) {
	var prop1_type1 *byte
	var prop1_type2 *byte
	var prop2_type1 *byte
	var prop2_type2 *byte
	ZendFormatType(prop1.GetType(), &prop1_type1, &prop1_type2)
	ZendFormatType(prop2.GetType(), &prop2_type1, &prop2_type2)
	ZendTypeError("Reference with value of type %s held by property %s::$%s of type %s%s is not compatible with property %s::$%s of type %s%s", b.CondF(zv.IsType(IS_OBJECT), func() []byte { return Z_OBJCE_P(zv).GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(zv.GetType()) }), prop1.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop1.GetName()), prop1_type1, prop1_type2, prop2.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop2.GetName()), prop2_type1, prop2_type2)
}
func ZendThrowRefTypeErrorZval(prop *ZendPropertyInfo, zv *Zval) {
	var prop_type1 *byte
	var prop_type2 *byte
	ZendFormatType(prop.GetType(), &prop_type1, &prop_type2)
	ZendTypeError("Cannot assign %s to reference held by property %s::$%s of type %s%s", b.CondF(zv.IsType(IS_OBJECT), func() []byte { return Z_OBJCE_P(zv).GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(zv.GetType()) }), prop.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop.GetName()), prop_type1, prop_type2)
}
func ZendThrowConflictingCoercionError(prop1 *ZendPropertyInfo, prop2 *ZendPropertyInfo, zv *Zval) {
	var prop1_type1 *byte
	var prop1_type2 *byte
	var prop2_type1 *byte
	var prop2_type2 *byte
	ZendFormatType(prop1.GetType(), &prop1_type1, &prop1_type2)
	ZendFormatType(prop2.GetType(), &prop2_type1, &prop2_type2)
	ZendTypeError("Cannot assign %s to reference held by property %s::$%s of type %s%s and property %s::$%s of type %s%s, as this would result in an inconsistent type conversion", b.CondF(zv.IsType(IS_OBJECT), func() []byte { return Z_OBJCE_P(zv).GetName().GetVal() }, func() *byte { return ZendGetTypeByConst(zv.GetType()) }), prop1.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop1.GetName()), prop1_type1, prop1_type2, prop2.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(prop2.GetName()), prop2_type1, prop2_type2)
}
func IZendVerifyTypeAssignableZval(type_ptr *ZendType, self_ce *ZendClassEntry, zv *Zval, strict ZendBool) int {
	var type_ ZendType = *type_ptr
	var type_code ZendUchar
	var zv_type ZendUchar = zv.GetType()
	if type_.AllowNull() && zv_type == IS_NULL {
		return 1
	}
	if type_.IsClass() {
		if !(type_.IsCe()) {
			if ZendResolveClassType(type_ptr, self_ce) == 0 {
				return 0
			}
			type_ = *type_ptr
		}
		return zv_type == IS_OBJECT && InstanceofFunction(Z_OBJCE_P(zv), type_.Ce()) != 0
	}
	type_code = type_.Code()
	if type_code == zv_type || type_code == _IS_BOOL && (zv_type == IS_FALSE || zv_type == IS_TRUE) {
		return 1
	}
	if type_code == IS_ITERABLE {
		return ZendIsIterable(zv)
	}

	/* SSTH Exception: IS_LONG may be accepted as IS_DOUBLE (converted) */

	if strict != 0 {
		if type_code == IS_DOUBLE && zv_type == IS_LONG {
			return -1
		}
		return 0
	}

	/* No weak conversions for arrays and objects */

	if type_code == IS_ARRAY || type_code == IS_OBJECT {
		return 0
	}

	/* NULL may be accepted only by nullable hints (this is already checked) */

	if zv_type == IS_NULL {
		return 0
	}

	/* Coercion may be necessary, check separately */

	return -1

	/* Coercion may be necessary, check separately */
}
func ZendVerifyRefAssignableZval(ref *ZendReference, zv *Zval, strict ZendBool) ZendBool {
	var prop *ZendPropertyInfo

	/* The value must satisfy each property type, and coerce to the same value for each property
	 * type. Right now, the latter rule means that *if* coercion is necessary, then all types
	 * must be the same (modulo nullability). To handle this, remember the first type we see and
	 * compare against it when coercion becomes necessary. */

	var seen_prop *ZendPropertyInfo = nil
	var seen_type ZendUchar
	var needs_coercion ZendBool = 0
	ZEND_ASSERT(zv.GetType() != IS_REFERENCE)
	var _source_list *ZendPropertyInfoSourceList = &(ref.GetSources())
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
			prop = *_prop
			var result int = IZendVerifyTypeAssignableZval(prop.GetType(), prop.GetCe(), zv, strict)
			if result == 0 {
				ZendThrowRefTypeErrorZval(prop, zv)
				return 0
			}
			if result < 0 {
				needs_coercion = 1
			}
			if seen_prop == nil {
				seen_prop = prop
				if prop.GetType().IsClass() {
					seen_type = IS_OBJECT
				} else {
					seen_type = prop.GetType().Code()
				}
			} else if needs_coercion != 0 && seen_type != prop.GetType().Code() {
				ZendThrowConflictingCoercionError(seen_prop, prop, zv)
				return 0
			}
		}
	}
	if needs_coercion != 0 && ZendVerifyWeakScalarTypeHint(seen_type, zv) == 0 {
		ZendThrowRefTypeErrorZval(seen_prop, zv)
		return 0
	}
	return 1
}
func IZvalPtrDtorNoref(zval_ptr *Zval) {
	if Z_REFCOUNTED_P(zval_ptr) {
		var ref *ZendRefcounted = zval_ptr.GetCounted()
		ZEND_ASSERT(zval_ptr.GetType() != IS_REFERENCE)
		if ref.DecGcRefcount() == 0 {
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
			if ref.DecGcRefcount() == 0 {
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
	if Z_ISREF_P(val) && ZEND_REF_HAS_TYPE_SOURCES(val.GetRef()) {
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
		ZVAL_DEREF(val)
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
func ZendFetchThisVar(type_ int, opline *ZendOp, _ EXECUTE_DATA_D) {
	var result *Zval = EX_VAR(opline.GetResult().GetVar())
	switch type_ {
	case BP_VAR_R:
		if EX(This).u1.v.type_ == IS_OBJECT {
			ZVAL_OBJ(result, EX(This).GetObj())
			Z_ADDREF_P(result)
		} else {
			ZVAL_NULL(result)
			ZendError(E_NOTICE, "Undefined variable: this")
		}
		break
	case BP_VAR_IS:
		if EX(This).u1.v.type_ == IS_OBJECT {
			ZVAL_OBJ(result, EX(This).GetObj())
			Z_ADDREF_P(result)
		} else {
			ZVAL_NULL(result)
		}
		break
	case BP_VAR_RW:

	case BP_VAR_W:
		ZVAL_UNDEF(result)
		ZendThrowError(nil, "Cannot re-assign $this")
		break
	case BP_VAR_UNSET:
		ZVAL_UNDEF(result)
		ZendThrowError(nil, "Cannot unset $this")
		break
	default:
		break
	}
}
func ZendWrongCloneCall(clone *ZendFunction, scope *ZendClassEntry) {
	ZendThrowError(nil, "Call to %s %s::__clone() from context '%s'", ZendVisibilityString(clone.GetFnFlags()), clone.GetScope().GetName().GetVal(), b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
}
func ExecuteInternal(execute_data *ZendExecuteData, return_value *Zval) {
	execute_data.GetFunc().GetInternalFunction().GetHandler()(execute_data, return_value)
}
func ZendCleanAndCacheSymbolTable(symbol_table *ZendArray) {
	/* Clean before putting into the cache, since clean could call dtors,
	 * which could use the cached hash. Also do this before the check for
	 * available cache slots, as those may be used by a dtor as well. */

	ZendSymtableClean(symbol_table)
	if ExecutorGlobals.GetSymtableCachePtr() >= ExecutorGlobals.GetSymtableCacheLimit() {
		ZendArrayDestroy(symbol_table)
	} else {
		*(b.PostInc(&(ExecutorGlobals.GetSymtableCachePtr()))) = symbol_table
	}
}
func IFreeCompiledVariables(execute_data *ZendExecuteData) {
	var cv *Zval = EX_VAR_NUM(0)
	var count int = EX(func_).op_array.last_var
	for count != 0 {
		if Z_REFCOUNTED_P(cv) {
			var r *ZendRefcounted = cv.GetCounted()
			if r.DecGcRefcount() == 0 {
				ZVAL_NULL(cv)
				RcDtorFunc(r)
			} else {
				GcCheckPossibleRoot(r)
			}
		}
		cv++
		count--
	}
}
func ZendFreeCompiledVariables(execute_data *ZendExecuteData) { IFreeCompiledVariables(execute_data) }
func ZEND_VM_INTERRUPT_CHECK() {
	if ExecutorGlobals.GetVmInterrupt() != 0 {
		ZEND_VM_INTERRUPT()
	}
}
func ZEND_VM_LOOP_INTERRUPT_CHECK() {
	if ExecutorGlobals.GetVmInterrupt() != 0 {
		ZEND_VM_LOOP_INTERRUPT()
	}
}
func ZendCopyExtraArgs(EXECUTE_DATA_D) {
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
			ZVAL_UNDEF(src)
			src--
			if !(b.PreDec(&count)) {
				break
			}
		}
		if Z_TYPE_INFO_REFCOUNTED(type_flags) {
			ZEND_ADD_CALL_FLAG(execute_data, ZEND_CALL_FREE_EXTRA_ARGS)
		}
	} else {
		for {
			if Z_REFCOUNTED_P(src) {
				ZEND_ADD_CALL_FLAG(execute_data, ZEND_CALL_FREE_EXTRA_ARGS)
				break
			}
			src--
			if !(b.PreDec(&count)) {
				break
			}
		}
	}
}
func ZendInitCvs(first uint32, last uint32, _ EXECUTE_DATA_D) {
	if first < last {
		var count uint32 = last - first
		var var_ *Zval = EX_VAR_NUM(first)
		for {
			ZVAL_UNDEF(var_)
			var_++
			if !(b.PreDec(&count)) {
				break
			}
		}
	}
}
func IInitFuncExecuteData(op_array *ZendOpArray, return_value *Zval, may_be_trampoline ZendBool, _ EXECUTE_DATA_D) {
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
			ZendCopyExtraArgs(EXECUTE_DATA_C)
		}
	} else if !op_array.IsHasTypeHints() {

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

		EX(opline) += num_args

		/* Skip useless ZEND_RECV and ZEND_RECV_INIT opcodes */

	}

	/* Initialize CV variables (skip arguments) */

	ZendInitCvs(num_args, op_array.GetLastVar(), EXECUTE_DATA_C)
	EX(run_time_cache) = RUN_TIME_CACHE(op_array)
	ExecutorGlobals.SetCurrentExecuteData(execute_data)
}
func InitFuncRunTimeCacheI(op_array *ZendOpArray) {
	var run_time_cache *any
	ZEND_ASSERT(RUN_TIME_CACHE(op_array) == nil)
	run_time_cache = ZendArenaAlloc(&(CompilerGlobals.GetArena()), op_array.GetCacheSize())
	memset(run_time_cache, 0, op_array.GetCacheSize())
	ZEND_MAP_PTR_SET(op_array.run_time_cache, run_time_cache)
}
func InitFuncRunTimeCache(op_array *ZendOpArray) { InitFuncRunTimeCacheI(op_array) }
func ZendFetchFunction(name *ZendString) *ZendFunction {
	var zv *Zval = ExecutorGlobals.GetFunctionTable().Find(name)
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
	var zv *Zval = ExecutorGlobals.GetFunctionTable().StrFind(name, len_)
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
func IInitCodeExecuteData(execute_data *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	ZEND_ASSERT(EX(func_) == (*ZendFunction)(op_array))
	EX(opline) = op_array.GetOpcodes()
	EX(call) = nil
	EX(return_value) = return_value
	ZendAttachSymbolTable(execute_data)
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
	ExecutorGlobals.SetCurrentExecuteData(execute_data)
}
func ZendInitFuncExecuteData(ex *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	var execute_data *ZendExecuteData = ex
	EX(prev_execute_data) = ExecutorGlobals.GetCurrentExecuteData()
	if !(RUN_TIME_CACHE(op_array)) {
		InitFuncRunTimeCache(op_array)
	}
	IInitFuncExecuteData(op_array, return_value, 1, EXECUTE_DATA_C)
}
func ZendInitCodeExecuteData(execute_data *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	EX(prev_execute_data) = ExecutorGlobals.GetCurrentExecuteData()
	IInitCodeExecuteData(execute_data, op_array, return_value)
}
func ZendInitExecuteData(execute_data *ZendExecuteData, op_array *ZendOpArray, return_value *Zval) {
	if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
		ZendInitCodeExecuteData(execute_data, op_array, return_value)
	} else {
		ZendInitFuncExecuteData(execute_data, op_array, return_value)
	}
}
func ZendVmStackCopyCallFrame(call *ZendExecuteData, passed_args uint32, additional_args uint32) *ZendExecuteData {
	var new_call *ZendExecuteData
	var used_stack int = ExecutorGlobals.GetVmStackTop() - (*Zval)(call) + additional_args

	/* copy call frame into new stack segment */

	new_call = ZendVmStackExtend(used_stack * b.SizeOf("zval"))
	*new_call = *call
	ZEND_ADD_CALL_FLAG(new_call, ZEND_CALL_ALLOCATED)
	if passed_args != 0 {
		var src *Zval = ZEND_CALL_ARG(call, 1)
		var dst *Zval = ZEND_CALL_ARG(new_call, 1)
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

	ExecutorGlobals.GetVmStack().GetPrev().SetTop((*Zval)(call))

	/* delete previous stack segment if it became empty */

	if ExecutorGlobals.GetVmStack().GetPrev().GetTop() == ZEND_VM_STACK_ELEMENTS(ExecutorGlobals.GetVmStack().GetPrev()) {
		var r ZendVmStack = ExecutorGlobals.GetVmStack().GetPrev()
		ExecutorGlobals.GetVmStack().SetPrev(r.GetPrev())
		Efree(r)
	}
	return new_call
}
func ZendVmStackExtendCallFrame(call **ZendExecuteData, passed_args uint32, additional_args uint32) {
	if uint32(ExecutorGlobals.GetVmStackEnd()-ExecutorGlobals.GetVmStackTop()) > additional_args {
		ExecutorGlobals.SetVmStackTop(ExecutorGlobals.GetVmStackTop() + additional_args)
	} else {
		*call = ZendVmStackCopyCallFrame(*call, passed_args, additional_args)
	}
}
func ZendGetRunningGenerator(EXECUTE_DATA_D) *ZendGenerator {
	/* The generator object is stored in EX(return_value) */

	var generator *ZendGenerator = (*ZendGenerator)(EX(return_value))

	/* However control may currently be delegated to another generator.
	 * That's the one we're interested in. */

	return generator

	/* However control may currently be delegated to another generator.
	 * That's the one we're interested in. */
}
func CleanupUnfinishedCalls(execute_data *ZendExecuteData, op_num uint32) {
	if EX(call) {
		var call *ZendExecuteData = EX(call)
		var opline *ZendOp = EX(func_).op_array.opcodes + op_num
		var level int
		var do_exit int
		if opline.GetOpcode() == ZEND_INIT_FCALL || opline.GetOpcode() == ZEND_INIT_FCALL_BY_NAME || opline.GetOpcode() == ZEND_INIT_NS_FCALL_BY_NAME || opline.GetOpcode() == ZEND_INIT_DYNAMIC_CALL || opline.GetOpcode() == ZEND_INIT_USER_CALL || opline.GetOpcode() == ZEND_INIT_METHOD_CALL || opline.GetOpcode() == ZEND_INIT_STATIC_METHOD_CALL || opline.GetOpcode() == ZEND_NEW {
			ZEND_ASSERT(op_num != 0)
			opline--
		}
		for {

			/* If the exception was thrown during a function call there might be
			 * arguments pushed to the stack that have to be dtor'ed. */

			level = 0
			do_exit = 0
			for {
				switch opline.GetOpcode() {
				case ZEND_DO_FCALL:

				case ZEND_DO_ICALL:

				case ZEND_DO_UCALL:

				case ZEND_DO_FCALL_BY_NAME:
					level++
					break
				case ZEND_INIT_FCALL:

				case ZEND_INIT_FCALL_BY_NAME:

				case ZEND_INIT_NS_FCALL_BY_NAME:

				case ZEND_INIT_DYNAMIC_CALL:

				case ZEND_INIT_USER_CALL:

				case ZEND_INIT_METHOD_CALL:

				case ZEND_INIT_STATIC_METHOD_CALL:

				case ZEND_NEW:
					if level == 0 {
						ZEND_CALL_NUM_ARGS(call) = 0
						do_exit = 1
					}
					level--
					break
				case ZEND_SEND_VAL:

				case ZEND_SEND_VAL_EX:

				case ZEND_SEND_VAR:

				case ZEND_SEND_VAR_EX:

				case ZEND_SEND_FUNC_ARG:

				case ZEND_SEND_REF:

				case ZEND_SEND_VAR_NO_REF:

				case ZEND_SEND_VAR_NO_REF_EX:

				case ZEND_SEND_USER:
					if level == 0 {
						ZEND_CALL_NUM_ARGS(call) = opline.GetOp2().GetNum()
						do_exit = 1
					}
					break
				case ZEND_SEND_ARRAY:

				case ZEND_SEND_UNPACK:
					if level == 0 {
						do_exit = 1
					}
					break
				}
				if do_exit == 0 {
					opline--
				}
				if do_exit != 0 {
					break
				}
			}
			if call.GetPrevExecuteData() != nil {

				/* skip current call region */

				level = 0
				do_exit = 0
				for {
					switch opline.GetOpcode() {
					case ZEND_DO_FCALL:

					case ZEND_DO_ICALL:

					case ZEND_DO_UCALL:

					case ZEND_DO_FCALL_BY_NAME:
						level++
						break
					case ZEND_INIT_FCALL:

					case ZEND_INIT_FCALL_BY_NAME:

					case ZEND_INIT_NS_FCALL_BY_NAME:

					case ZEND_INIT_DYNAMIC_CALL:

					case ZEND_INIT_USER_CALL:

					case ZEND_INIT_METHOD_CALL:

					case ZEND_INIT_STATIC_METHOD_CALL:

					case ZEND_NEW:
						if level == 0 {
							do_exit = 1
						}
						level--
						break
					}
					opline--
					if do_exit != 0 {
						break
					}
				}
			}
			ZendVmStackFreeArgs(EX(call))
			if (ZEND_CALL_INFO(call) & ZEND_CALL_RELEASE_THIS) != 0 {
				OBJ_RELEASE(call.GetThis().GetObj())
			}
			if call.GetFunc().IsClosure() {
				ZendObjectRelease(ZEND_CLOSURE_OBJECT(call.GetFunc()))
			} else if call.GetFunc().IsCallViaTrampoline() {
				ZendStringReleaseEx(call.GetFunc().GetFunctionName(), 0)
				ZendFreeTrampoline(call.GetFunc())
			}
			EX(call) = call.GetPrevExecuteData()
			ZendVmStackFreeCallFrame(call)
			call = EX(call)
			if call == nil {
				break
			}
		}
	}
}
func FindLiveRange(op_array *ZendOpArray, op_num uint32, var_num uint32) *ZendLiveRange {
	var i int
	for i = 0; i < op_array.GetLastLiveRange(); i++ {
		var range_ *ZendLiveRange = op_array.GetLiveRange()[i]
		if op_num >= range_.GetStart() && op_num < range_.GetEnd() && var_num == (range_.GetVar() & ^ZEND_LIVE_MASK) {
			return range_
		}
	}
	return nil
}
func CleanupLiveVars(execute_data *ZendExecuteData, op_num uint32, catch_op_num uint32) {
	var i int
	for i = 0; i < EX(func_).op_array.last_live_range; i++ {
		var range_ *ZendLiveRange = EX(func_).op_array.live_range[i]
		if range_.GetStart() > op_num {

			/* further blocks will not be relevant... */

			break

			/* further blocks will not be relevant... */

		} else if op_num < range_.GetEnd() {
			if catch_op_num == 0 || catch_op_num >= range_.GetEnd() {
				var kind uint32 = range_.GetVar() & ZEND_LIVE_MASK
				var var_num uint32 = range_.GetVar() & ^ZEND_LIVE_MASK
				var var_ *Zval = EX_VAR(var_num)
				if kind == ZEND_LIVE_TMPVAR {
					ZvalPtrDtorNogc(var_)
				} else if kind == ZEND_LIVE_NEW {
					var obj *ZendObject
					ZEND_ASSERT(var_.IsType(IS_OBJECT))
					obj = var_.GetObj()
					ZendObjectStoreCtorFailed(obj)
					OBJ_RELEASE(obj)
				} else if kind == ZEND_LIVE_LOOP {
					if var_.GetType() != IS_ARRAY && var_.GetFeIterIdx() != uint32-1 {
						ZendHashIteratorDel(var_.GetFeIterIdx())
					}
					ZvalPtrDtorNogc(var_)
				} else if kind == ZEND_LIVE_ROPE {
					var rope **ZendString = (**ZendString)(var_)
					var last *ZendOp = EX(func_).op_array.opcodes + op_num
					for last.GetOpcode() != ZEND_ROPE_ADD && last.GetOpcode() != ZEND_ROPE_INIT || last.GetResult().GetVar() != var_num {
						ZEND_ASSERT(last >= EX(func_).op_array.opcodes)
						last--
					}
					if last.GetOpcode() == ZEND_ROPE_INIT {
						ZendStringReleaseEx(*rope, 0)
					} else {
						var j int = last.GetExtendedValue()
						for {
							ZendStringReleaseEx(rope[j], 0)
							if !(b.PostDec(&j)) {
								break
							}
						}
					}
				} else if kind == ZEND_LIVE_SILENCE {

					/* restore previous error_reporting value */

					if ExecutorGlobals.GetErrorReporting() == 0 && var_.GetLval() != 0 {
						ExecutorGlobals.SetErrorReporting(var_.GetLval())
					}

					/* restore previous error_reporting value */

				}
			}
		}
	}
}
func ZendCleanupUnfinishedExecution(execute_data *ZendExecuteData, op_num uint32, catch_op_num uint32) {
	CleanupUnfinishedCalls(execute_data, op_num)
	CleanupLiveVars(execute_data, op_num, catch_op_num)
}
func ZendSwapOperands(op *ZendOp) {
	var tmp ZnodeOp
	var tmp_type ZendUchar
	tmp = op.GetOp1()
	tmp_type = op.GetOp1Type()
	op.SetOp1(op.GetOp2())
	op.SetOp1Type(op.GetOp2Type())
	op.SetOp2(tmp)
	op.SetOp2Type(tmp_type)
}
func ZendInitDynamicCallString(function *ZendString, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var func_ *Zval
	var called_scope *ZendClassEntry
	var lcname *ZendString
	var colon *byte
	if b.Assign(&colon, ZendMemrchr(function.GetVal(), ':', function.GetLen())) != nil && colon > function.GetVal() && (*(colon - 1)) == ':' {
		var mname *ZendString
		var cname_length int = colon - function.GetVal() - 1
		var mname_length int = function.GetLen() - cname_length - (b.SizeOf("\"::\"") - 1)
		lcname = ZendStringInit(function.GetVal(), cname_length, 0)
		called_scope = ZendFetchClassByName(lcname, nil, ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
		if called_scope == nil {
			ZendStringReleaseEx(lcname, 0)
			return nil
		}
		mname = ZendStringInit(function.GetVal()+(cname_length+b.SizeOf("\"::\"")-1), mname_length, 0)
		if called_scope.GetGetStaticMethod() != nil {
			fbc = called_scope.GetGetStaticMethod()(called_scope, mname)
		} else {
			fbc = ZendStdGetStaticMethod(called_scope, mname, nil)
		}
		if fbc == nil {
			if ExecutorGlobals.GetException() == nil {
				ZendUndefinedMethod(called_scope, mname)
			}
			ZendStringReleaseEx(lcname, 0)
			ZendStringReleaseEx(mname, 0)
			return nil
		}
		ZendStringReleaseEx(lcname, 0)
		ZendStringReleaseEx(mname, 0)
		if !fbc.IsStatic() {
			ZendNonStaticMethodCall(fbc)
			if ExecutorGlobals.GetException() != nil {
				return nil
			}
		}
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
	} else {
		if function.GetVal()[0] == '\\' {
			lcname = ZendStringAlloc(function.GetLen()-1, 0)
			ZendStrTolowerCopy(lcname.GetVal(), function.GetVal()+1, function.GetLen()-1)
		} else {
			lcname = ZendStringTolower(function)
		}
		if b.Assign(&func_, ExecutorGlobals.GetFunctionTable().Find(lcname)) == nil {
			ZendThrowError(nil, "Call to undefined function %s()", function.GetVal())
			ZendStringReleaseEx(lcname, 0)
			return nil
		}
		ZendStringReleaseEx(lcname, 0)
		fbc = func_.GetFunc()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		called_scope = nil
	}
	return ZendVmStackPushCallFrame(ZEND_CALL_NESTED_FUNCTION|ZEND_CALL_DYNAMIC, fbc, num_args, called_scope)
}
func ZendInitDynamicCallObject(function *Zval, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var object_or_called_scope any
	var called_scope *ZendClassEntry
	var object *ZendObject
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	if Z_OBJ_HANDLER_P(function, get_closure) && Z_OBJ_HANDLER_P(function, get_closure)(function, &called_scope, &fbc, &object) == SUCCESS {
		object_or_called_scope = called_scope
		if fbc.IsClosure() {

			/* Delay closure destruction until its invocation */

			ZEND_CLOSURE_OBJECT(fbc).IncGcRefcount()
			call_info |= ZEND_CALL_CLOSURE
			if fbc.IsFakeClosure() {
				call_info |= ZEND_CALL_FAKE_CLOSURE
			}
			if object != nil {
				call_info |= ZEND_CALL_HAS_THIS
				object_or_called_scope = object
			}
		} else if object != nil {
			call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
			object.IncGcRefcount()
			object_or_called_scope = object
		}
	} else {
		ZendThrowError(nil, "Function name must be a string")
		return nil
	}
	if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
		InitFuncRunTimeCache(fbc.GetOpArray())
	}
	return ZendVmStackPushCallFrame(call_info, fbc, num_args, object_or_called_scope)
}
func ZendInitDynamicCallArray(function *ZendArray, num_args uint32) *ZendExecuteData {
	var fbc *ZendFunction
	var object_or_called_scope any
	var call_info uint32 = ZEND_CALL_NESTED_FUNCTION | ZEND_CALL_DYNAMIC
	if function.GetNNumOfElements() == 2 {
		var obj *Zval
		var method *Zval
		obj = function.IndexFind(0)
		method = function.IndexFind(1)
		if obj == nil || method == nil {
			ZendThrowError(nil, "Array callback has to contain indices 0 and 1")
			return nil
		}
		ZVAL_DEREF(obj)
		if obj.GetType() != IS_STRING && obj.GetType() != IS_OBJECT {
			ZendThrowError(nil, "First array member is not a valid class name or object")
			return nil
		}
		ZVAL_DEREF(method)
		if method.GetType() != IS_STRING {
			ZendThrowError(nil, "Second array member is not a valid method")
			return nil
		}
		if obj.IsType(IS_STRING) {
			var called_scope *ZendClassEntry = ZendFetchClassByName(obj.GetStr(), nil, ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if called_scope == nil {
				return nil
			}
			if called_scope.GetGetStaticMethod() != nil {
				fbc = called_scope.GetGetStaticMethod()(called_scope, method.GetStr())
			} else {
				fbc = ZendStdGetStaticMethod(called_scope, method.GetStr(), nil)
			}
			if fbc == nil {
				if ExecutorGlobals.GetException() == nil {
					ZendUndefinedMethod(called_scope, method.GetStr())
				}
				return nil
			}
			if !fbc.IsStatic() {
				ZendNonStaticMethodCall(fbc)
				if ExecutorGlobals.GetException() != nil {
					return nil
				}
			}
			object_or_called_scope = called_scope
		} else {
			var object *ZendObject = obj.GetObj()
			fbc = Z_OBJ_HT_P(obj).GetGetMethod()(&object, method.GetStr(), nil)
			if fbc == nil {
				if ExecutorGlobals.GetException() == nil {
					ZendUndefinedMethod(object.GetCe(), method.GetStr())
				}
				return nil
			}
			if fbc.IsStatic() {
				object_or_called_scope = object.GetCe()
			} else {
				call_info |= ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS
				object.IncGcRefcount()
				object_or_called_scope = object
			}
		}
	} else {
		ZendThrowError(nil, "Function name must be a string")
		return nil
	}
	if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
		InitFuncRunTimeCache(fbc.GetOpArray())
	}
	return ZendVmStackPushCallFrame(call_info, fbc, num_args, object_or_called_scope)
}
func ZendIncludeOrEval(inc_filename *Zval, type_ int) *ZendOpArray {
	var new_op_array *ZendOpArray = nil
	var tmp_inc_filename Zval
	ZVAL_UNDEF(&tmp_inc_filename)
	if inc_filename.GetType() != IS_STRING {
		var tmp *ZendString = ZvalTryGetStringFunc(inc_filename)
		if tmp == nil {
			return nil
		}
		ZVAL_STR(&tmp_inc_filename, tmp)
		inc_filename = &tmp_inc_filename
	}
	switch type_ {
	case ZEND_INCLUDE_ONCE:

	case ZEND_REQUIRE_ONCE:
		var file_handle ZendFileHandle
		var resolved_path *ZendString
		resolved_path = ZendResolvePath(Z_STRVAL_P(inc_filename), Z_STRLEN_P(inc_filename))
		if resolved_path != nil {
			if ExecutorGlobals.GetIncludedFiles().Exists(resolved_path) != 0 {
				goto already_compiled
			}
		} else if ExecutorGlobals.GetException() != nil {
			break
		} else if strlen(Z_STRVAL_P(inc_filename)) != Z_STRLEN_P(inc_filename) {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE_ONCE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), Z_STRVAL_P(inc_filename))
			break
		} else {
			resolved_path = ZendStringCopy(inc_filename.GetStr())
		}
		if SUCCESS == ZendStreamOpen(resolved_path.GetVal(), &file_handle) {
			if file_handle.GetOpenedPath() == nil {
				file_handle.SetOpenedPath(ZendStringCopy(resolved_path))
			}
			if ExecutorGlobals.GetIncludedFiles().AddEmptyElement(file_handle.GetOpenedPath()) != nil {
				var op_array *ZendOpArray = ZendCompileFile(&file_handle, b.Cond(type_ == ZEND_INCLUDE_ONCE, ZEND_INCLUDE, ZEND_REQUIRE))
				ZendDestroyFileHandle(&file_handle)
				ZendStringReleaseEx(resolved_path, 0)
				if tmp_inc_filename.GetType() != IS_UNDEF {
					ZvalPtrDtorStr(&tmp_inc_filename)
				}
				return op_array
			} else {
				ZendFileHandleDtor(&file_handle)
			already_compiled:
				new_op_array = ZEND_FAKE_OP_ARRAY
			}
		} else {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE_ONCE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), Z_STRVAL_P(inc_filename))
		}
		ZendStringReleaseEx(resolved_path, 0)
		break
	case ZEND_INCLUDE:

	case ZEND_REQUIRE:
		if strlen(Z_STRVAL_P(inc_filename)) != Z_STRLEN_P(inc_filename) {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), Z_STRVAL_P(inc_filename))
			break
		}
		new_op_array = CompileFilename(type_, inc_filename)
		break
	case ZEND_EVAL:
		var eval_desc *byte = ZendMakeCompiledStringDescription("eval()'d code")
		new_op_array = ZendCompileString(inc_filename, eval_desc)
		Efree(eval_desc)
		break
	default:
		break
	}
	if tmp_inc_filename.GetType() != IS_UNDEF {
		ZvalPtrDtorStr(&tmp_inc_filename)
	}
	return new_op_array
}
func ZendDoFcallOverloaded(call *ZendExecuteData, ret *Zval) int {
	var fbc *ZendFunction = call.GetFunc()
	var object *ZendObject

	/* Not sure what should be done here if it's a static method */

	if call.GetThis().GetType() != IS_OBJECT {
		ZendVmStackFreeArgs(call)
		if fbc.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
			ZendStringReleaseEx(fbc.GetFunctionName(), 0)
		}
		Efree(fbc)
		ZendVmStackFreeCallFrame(call)
		ZendThrowError(nil, "Cannot call overloaded function for non-object")
		return 0
	}
	object = call.GetThis().GetObj()
	ZVAL_NULL(ret)
	ExecutorGlobals.SetCurrentExecuteData(call)
	object.GetHandlers().GetCallMethod()(fbc.GetFunctionName(), object, call, ret)
	ExecutorGlobals.SetCurrentExecuteData(call.GetPrevExecuteData())
	ZendVmStackFreeArgs(call)
	if fbc.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
		ZendStringReleaseEx(fbc.GetFunctionName(), 0)
	}
	Efree(fbc)
	return 1
}
func ZendFeResetIterator(array_ptr *Zval, by_ref int, opline *ZendOp, _ EXECUTE_DATA_D) ZendBool {
	var ce *ZendClassEntry = Z_OBJCE_P(array_ptr)
	var iter *ZendObjectIterator = ce.GetGetIterator()(ce, array_ptr, by_ref)
	var is_empty ZendBool
	if iter == nil || ExecutorGlobals.GetException() != nil {
		if iter != nil {
			OBJ_RELEASE(iter.GetStd())
		}
		if ExecutorGlobals.GetException() == nil {
			ZendThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
		}
		ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
		return 1
	}
	iter.SetIndex(0)
	if iter.GetFuncs().GetRewind() != nil {
		iter.GetFuncs().GetRewind()(iter)
		if ExecutorGlobals.GetException() != nil {
			OBJ_RELEASE(iter.GetStd())
			ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
			return 1
		}
	}
	is_empty = iter.GetFuncs().GetValid()(iter) != SUCCESS
	if ExecutorGlobals.GetException() != nil {
		OBJ_RELEASE(iter.GetStd())
		ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
		return 1
	}
	iter.SetIndex(-1)
	ZVAL_OBJ(EX_VAR(opline.GetResult().GetVar()), iter.GetStd())
	EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(uint32 - 1)
	return is_empty
}
func _zendQuickGetConstant(key *Zval, flags uint32, check_defined_only int, opline *ZendOp, _ EXECUTE_DATA_D) int {
	var zv *Zval
	var orig_key *Zval = key
	var c *ZendConstant = nil
	zv = ExecutorGlobals.GetZendConstants().FindEx(key.GetStr(), 1)
	if zv != nil {
		c = (*ZendConstant)(zv.GetPtr())
	} else {
		key++
		zv = ExecutorGlobals.GetZendConstants().FindEx(key.GetStr(), 1)
		if zv != nil && (ZEND_CONSTANT_FLAGS((*ZendConstant)(zv.GetPtr()))&CONST_CS) == 0 {
			c = (*ZendConstant)(zv.GetPtr())
		} else {
			if (flags & (IS_CONSTANT_IN_NAMESPACE | IS_CONSTANT_UNQUALIFIED)) == (IS_CONSTANT_IN_NAMESPACE | IS_CONSTANT_UNQUALIFIED) {
				key++
				zv = ExecutorGlobals.GetZendConstants().FindEx(key.GetStr(), 1)
				if zv != nil {
					c = (*ZendConstant)(zv.GetPtr())
				} else {
					key++
					zv = ExecutorGlobals.GetZendConstants().FindEx(key.GetStr(), 1)
					if zv != nil && (ZEND_CONSTANT_FLAGS((*ZendConstant)(zv.GetPtr()))&CONST_CS) == 0 {
						c = (*ZendConstant)(zv.GetPtr())
					}
				}
			}
		}
	}
	if c == nil {
		if check_defined_only == 0 {
			if (opline.GetOp1().GetNum() & IS_CONSTANT_UNQUALIFIED) != 0 {
				var actual *byte = (*byte)(ZendMemrchr(Z_STRVAL_P(RT_CONSTANT(opline, opline.GetOp2())), '\\', Z_STRLEN_P(RT_CONSTANT(opline, opline.GetOp2()))))
				if actual == nil {
					ZVAL_STR_COPY(EX_VAR(opline.GetResult().GetVar()), RT_CONSTANT(opline, opline.GetOp2()).GetStr())
				} else {
					actual++
					ZVAL_STRINGL(EX_VAR(opline.GetResult().GetVar()), actual, Z_STRLEN_P(RT_CONSTANT(opline, opline.GetOp2()))-(actual-Z_STRVAL_P(RT_CONSTANT(opline, opline.GetOp2()))))
				}

				/* non-qualified constant - allow text substitution */

				ZendError(E_WARNING, "Use of undefined constant %s - assumed '%s' (this will throw an Error in a future version of PHP)", Z_STRVAL_P(EX_VAR(opline.GetResult().GetVar())), Z_STRVAL_P(EX_VAR(opline.GetResult().GetVar())))

				/* non-qualified constant - allow text substitution */

			} else {
				ZendThrowError(nil, "Undefined constant '%s'", Z_STRVAL_P(RT_CONSTANT(opline, opline.GetOp2())))
				ZVAL_UNDEF(EX_VAR(opline.GetResult().GetVar()))
			}
		}
		return FAILURE
	}
	if check_defined_only == 0 {
		ZVAL_COPY_OR_DUP(EX_VAR(opline.GetResult().GetVar()), c.GetValue())
		if (ZEND_CONSTANT_FLAGS(c) & (CONST_CS | CONST_CT_SUBST)) == 0 {
			var ns_sep *byte
			var shortname_offset int
			var shortname_len int
			var is_deprecated ZendBool
			if (flags & IS_CONSTANT_UNQUALIFIED) != 0 {
				var access_key *Zval
				if (flags & IS_CONSTANT_IN_NAMESPACE) == 0 {
					access_key = orig_key - 1
				} else {
					if key < orig_key+2 {
						goto check_short_name
					} else {
						access_key = orig_key + 2
					}
				}
				is_deprecated = !(ZendStringEquals(c.GetName(), access_key.GetStr()))
			} else {
			check_short_name:

				/* Namespaces are always case-insensitive. Only compare shortname. */

				ns_sep = ZendMemrchr(c.GetName().GetVal(), '\\', c.GetName().GetLen())
				if ns_sep != nil {
					shortname_offset = ns_sep - c.GetName().GetVal() + 1
					shortname_len = c.GetName().GetLen() - shortname_offset
				} else {
					shortname_offset = 0
					shortname_len = c.GetName().GetLen()
				}
				is_deprecated = memcmp(c.GetName().GetVal()+shortname_offset, Z_STRVAL_P(orig_key-1)+shortname_offset, shortname_len) != 0
			}
			if is_deprecated != 0 {
				ZendError(E_DEPRECATED, "Case-insensitive constants are deprecated. "+"The correct casing for this constant is \"%s\"", c.GetName().GetVal())
				return SUCCESS
			}
		}
	}
	CACHE_PTR(opline.GetExtendedValue(), c)
	return SUCCESS
}
func ZendQuickGetConstant(key *Zval, flags uint32, opline *ZendOp, _ EXECUTE_DATA_D) {
	_zendQuickGetConstant(key, flags, 0, OPLINE_C, EXECUTE_DATA_C)
}
func ZendQuickCheckConstant(key *Zval, opline *ZendOp, _ EXECUTE_DATA_D) int {
	return _zendQuickGetConstant(key, 0, 1, OPLINE_C, EXECUTE_DATA_C)
}
func ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION() {
	OPLINE = EX(opline) + 1
	ZEND_VM_CONTINUE()
}
func ZEND_VM_NEXT_OPCODE() {
	ZEND_ASSERT(ExecutorGlobals.GetException() == nil)
	OPLINE = opline + 1
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SET_RELATIVE_OPCODE(opline *ZendOp, offset uint32) {
	OPLINE = ZEND_OFFSET_TO_OPLINE(opline, offset)
	ZEND_VM_INTERRUPT_CHECK()
}
func ZEND_VM_JMP_EX(new_op *ZendOp, check_exception int) {
	if check_exception != 0 && ExecutorGlobals.GetException() != nil {
		HANDLE_EXCEPTION()
	}
	OPLINE = new_op
	ZEND_VM_INTERRUPT_CHECK()
	ZEND_VM_CONTINUE()
}
func ZEND_VM_JMP(new_op *ZendOp) { ZEND_VM_JMP_EX(new_op, 1) }
func ZEND_VM_INC_OPCODE() int {
	OPLINE++
	return OPLINE - 1
}
func ZEND_VM_SMART_BRANCH(_result __auto__, _check int) {
	for {
		if _check != 0 && ExecutorGlobals.GetException() != nil {
			break
		}
		if (opline + 1).opcode == ZEND_JMPZ {
			if _result {
				OPLINE = opline + 2
			} else {
				OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
				ZEND_VM_INTERRUPT_CHECK()
			}
		} else if (opline + 1).opcode == ZEND_JMPNZ {
			if !_result {
				OPLINE = opline + 2
			} else {
				OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
				ZEND_VM_INTERRUPT_CHECK()
			}
		} else {
			break
		}
		ZEND_VM_CONTINUE()
		break
	}
}
func ZEND_VM_SMART_BRANCH_JMPZ(_result int, _check int) {
	for {
		if _check != 0 && ExecutorGlobals.GetException() != nil {
			break
		}
		if _result != 0 {
			OPLINE = opline + 2
		} else {
			OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
			ZEND_VM_INTERRUPT_CHECK()
		}
		ZEND_VM_CONTINUE()
		break
	}
}
func ZEND_VM_SMART_BRANCH_JMPNZ(_result int, _check int) {
	for {
		if _check != 0 && ExecutorGlobals.GetException() != nil {
			break
		}
		if _result == 0 {
			OPLINE = opline + 2
		} else {
			OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
			ZEND_VM_INTERRUPT_CHECK()
		}
		ZEND_VM_CONTINUE()
		break
	}
}
func ZEND_VM_SMART_BRANCH_TRUE() {
	if (opline + 1).opcode == ZEND_JMPNZ {
		OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
		ZEND_VM_INTERRUPT_CHECK()
		ZEND_VM_CONTINUE()
	} else if (opline + 1).opcode == ZEND_JMPZ {
		OPLINE = opline + 2
		ZEND_VM_CONTINUE()
	}
}
func ZEND_VM_SMART_BRANCH_TRUE_JMPZ() {
	OPLINE = opline + 2
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SMART_BRANCH_TRUE_JMPNZ() {
	OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
	ZEND_VM_INTERRUPT_CHECK()
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SMART_BRANCH_FALSE() {
	if (opline + 1).opcode == ZEND_JMPNZ {
		OPLINE = opline + 2
		ZEND_VM_CONTINUE()
	} else if (opline + 1).opcode == ZEND_JMPZ {
		OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
		ZEND_VM_INTERRUPT_CHECK()
		ZEND_VM_CONTINUE()
	}
}
func ZEND_VM_SMART_BRANCH_FALSE_JMPZ() {
	OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
	ZEND_VM_INTERRUPT_CHECK()
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SMART_BRANCH_FALSE_JMPNZ() {
	OPLINE = opline + 2
	ZEND_VM_CONTINUE()
}
func UNDEF_RESULT() {
	if (opline.result_type & (IS_VAR | IS_TMP_VAR)) != 0 {
		ZVAL_UNDEF(EX_VAR(opline.result.var_))
	}
}
func ZendSetUserOpcodeHandler(opcode ZendUchar, handler UserOpcodeHandlerT) int {
	if opcode != ZEND_USER_OPCODE {
		if handler == nil {

			/* restore the original handler */

			ZendUserOpcodes[opcode] = opcode

			/* restore the original handler */

		} else {
			ZendUserOpcodes[opcode] = ZEND_USER_OPCODE
		}
		ZendUserOpcodeHandlers[opcode] = handler
		return SUCCESS
	}
	return FAILURE
}
func ZendGetUserOpcodeHandler(opcode ZendUchar) UserOpcodeHandlerT {
	return ZendUserOpcodeHandlers[opcode]
}
func ZendGetZvalPtr(opline *ZendOp, op_type int, node *ZnodeOp, execute_data *ZendExecuteData, should_free *ZendFreeOp, type_ int) *Zval {
	var ret *Zval
	switch op_type {
	case IS_CONST:
		ret = RT_CONSTANT(opline, *node)
		*should_free = nil
		break
	case IS_TMP_VAR:

	case IS_VAR:
		ret = EX_VAR(node.GetVar())
		*should_free = ret
		break
	case IS_CV:
		ret = EX_VAR(node.GetVar())
		*should_free = nil
		break
	default:
		ret = nil
		*should_free = ret
		break
	}
	return ret
}
