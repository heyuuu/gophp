package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_REF_HAS_TYPE_SOURCES(ref *types.ZendReference) bool {
	return ref.GetSources().GetPtr() != nil
}
func ZEND_REF_FIRST_SOURCE(ref *types.ZendReference) *ZendPropertyInfo {
	if types.ZEND_PROPERTY_INFO_SOURCE_IS_LIST(ref.GetSources().GetList()) != 0 {
		return types.ZEND_PROPERTY_INFO_SOURCE_TO_LIST(ref.GetSources().GetList()).GetPtr()[0]
	} else {
		return ref.GetSources().GetPtr()
	}
}
func ZendCopyToVariable(variable_ptr *types.Zval, value *types.Zval, value_type types.ZendUchar, ref *types.ZendRefcounted) {
	types.ZVAL_COPY_VALUE(variable_ptr, value)

	if (value_type & (IS_CONST | IS_CV)) != 0 {

		variable_ptr.TryAddRefcount()

	} else if ref != nil {
		if ref.DelRefcount() == 0 {
			EfreeSize(ref, b.SizeOf("zend_reference"))
		} else {
			variable_ptr.TryAddRefcount()
		}

	}
}
func ZendAssignToVariable(variable_ptr *types.Zval, value *types.Zval, value_type types.ZendUchar, strict types.ZendBool) *types.Zval {
	var ref *types.ZendRefcounted = nil
	if value.IsReference() {
		ref = value.GetCounted()
		value = types.Z_REFVAL_P(value)
	}
	for {
		if variable_ptr.IsRefcounted() {
			var garbage *types.ZendRefcounted
			if variable_ptr.IsReference() {
				if ZEND_REF_HAS_TYPE_SOURCES(variable_ptr.GetRef()) {
					return ZendAssignToTypedRef(variable_ptr, value, value_type, strict, ref)
				}
				variable_ptr = types.Z_REFVAL_P(variable_ptr)
				if !(variable_ptr.IsRefcounted()) {
					break
				}
			}
			if variable_ptr.IsObject() && types.Z_OBJ_HT(*variable_ptr).GetSet() != nil {
				types.Z_OBJ_HT(*variable_ptr).GetSet()(variable_ptr, value)
				return variable_ptr
			}
			garbage = variable_ptr.GetCounted()
			ZendCopyToVariable(variable_ptr, value, value_type, ref)
			if garbage.DelRefcount() == 0 {
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
	return (*types.Zval)(stack) + ZEND_VM_STACK_HEADER_SLOTS
}
func ZendVmInitCallFrame(call *ZendExecuteData, call_info uint32, func_ types.IFunction, num_args uint32, object_or_called_scope any) {
	call.SetFunc(func_)
	call.GetThis().GetPtr() = object_or_called_scope
	ZEND_CALL_INFO(call) = call_info
	call.NumArgs() = num_args
}
func ZendVmStackPushCallFrameEx(used_stack uint32, call_info uint32, func_ types.IFunction, num_args uint32, object_or_called_scope any) *ZendExecuteData {
	var call *ZendExecuteData = (*ZendExecuteData)(EG__().GetVmStackTop())
	if used_stack > size_t((*byte)(EG__().GetVmStackEnd())-(*byte)(call)) {
		call = (*ZendExecuteData)(ZendVmStackExtend(used_stack))
		ZendVmInitCallFrame(call, call_info|ZEND_CALL_ALLOCATED, func_, num_args, object_or_called_scope)
		return call
	} else {
		EG__().SetVmStackTop((*types.Zval)((*byte)(call + used_stack)))
		ZendVmInitCallFrame(call, call_info, func_, num_args, object_or_called_scope)
		return call
	}
}
func ZendVmCalcUsedStack(num_args uint32, func_ types.IFunction) uint32 {
	var used_stack uint32 = ZEND_CALL_FRAME_SLOT + num_args
	if ZEND_USER_CODE(func_.GetType()) {
		used_stack += func_.GetOpArray().GetLastVar() + func_.GetOpArray().GetT() - b.Min(func_.GetOpArray().GetNumArgs(), num_args)
	}
	return used_stack * b.SizeOf("zval")
}
func ZendVmStackPushCallFrame(call_info uint32, func_ types.IFunction, num_args uint32, object_or_called_scope any) *ZendExecuteData {
	var used_stack uint32 = ZendVmCalcUsedStack(num_args, func_)
	return ZendVmStackPushCallFrameEx(used_stack, call_info, func_, num_args, object_or_called_scope)
}
func ZendVmStackFreeExtraArgsEx(call_info uint32, call *ZendExecuteData) {
	if (call_info & ZEND_CALL_FREE_EXTRA_ARGS) != 0 {
		var count uint32 = call.NumArgs() - call.GetFunc().GetOpArray().GetNumArgs()
		var p *types.Zval = call.VarNum(call.GetFunc().GetOpArray().GetLastVar() + call.GetFunc().GetOpArray().GetT())
		for {
			if p.IsRefcounted() {
				var r *types.ZendRefcounted = p.GetCounted()
				if r.DelRefcount() == 0 {
					p.SetNull()
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
	for _, p := range call.AllArgs() {
		if p.IsRefcounted() {
			r := p.GetCounted()
			if r.DelRefcount() == 0 {
				p.SetNull()
				RcDtorFunc(r)
			}
		}
	}
}
func ZendVmStackFreeCallFrameEx(call_info uint32, call *ZendExecuteData) {
	if (call_info & ZEND_CALL_ALLOCATED) != 0 {
		var p ZendVmStack = EG__().GetVmStack()
		var prev ZendVmStack = p.GetPrev()
		b.Assert(call == p.ElementsAsEx())
		EG__().SetVmStackTop(prev.GetTop())
		EG__().SetVmStackEnd(prev.GetEnd())
		EG__().SetVmStack(prev)
		Efree(p)
	} else {
		EG__().SetVmStackTop((*types.Zval)(call))
	}
}
func ZendVmStackFreeCallFrame(call *ZendExecuteData) {
	ZendVmStackFreeCallFrameEx(ZEND_CALL_INFO(call), call)
}
func CACHE_ADDR(num __auto__) *any {
	return (*any)((*byte)(executeData.GetRunTimeCache() + num))
}
func CACHED_PTR(num __auto__) any {
	return (*any)((*byte)(executeData.GetRunTimeCache() + num))[0]
}
func CACHE_PTR(num __auto__, ptr any) {
	(*any)((*byte)(executeData.GetRunTimeCache() + num))[0] = ptr
}
func CACHED_POLYMORPHIC_PTR(num __auto__, ce __auto__) bool {
	return (*any)((*byte)(executeData.GetRunTimeCache() + num))[0] == any(b.CondF1(ce, func() any { return (*any)((*byte)(executeData.GetRunTimeCache() + num))[1] }, nil))
}
func CACHE_POLYMORPHIC_PTR(num uint32, ce any, ptr any) {
	var slot *any = (*any)((*byte)(executeData.GetRunTimeCache() + num))
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
func CACHE_POLYMORPHIC_PTR_EX(slot *any, ce *types.ClassEntry, ptr any) {
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
func ZEND_CLASS_HAS_TYPE_HINTS(ce *types.ClassEntry) bool {
	return (ce.GetCeFlags() & AccHasTypeHints) == AccHasTypeHints
}
func ZEND_REF_ADD_TYPE_SOURCE(ref *types.ZendReference, source *ZendPropertyInfo) {
	ZendRefAddTypeSource(&(ref.GetSources()), source)
}
func ZEND_REF_DEL_TYPE_SOURCE(ref *types.ZendReference, source *ZendPropertyInfo) {
	ZendRefDelTypeSource(&(ref.GetSources()), source)
}
func GetZvalPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *types.Zval {
	return _getZvalPtr(op_type, node, should_free, type_, executeData, opline)
}
func GetZvalPtrDeref(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *types.Zval {
	return _getZvalPtrDeref(op_type, node, should_free, type_, executeData, opline)
}
func GetZvalPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *types.Zval {
	return _getZvalPtrUndef(op_type, node, should_free, type_, executeData, opline)
}
func GetOpDataZvalPtrR(op_type int, node ZnodeOp, should_free *ZendFreeOp) *types.Zval {
	return _getOpDataZvalPtrR(op_type, node, should_free, executeData, opline)
}
func GetOpDataZvalPtrDerefR(op_type int, node ZnodeOp, should_free *ZendFreeOp) *types.Zval {
	return _getOpDataZvalPtrDerefR(op_type, node, should_free, executeData, opline)
}
func GetZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *types.Zval {
	return _getZvalPtrPtr(op_type, node, should_free, type_, executeData)
}
func GetZvalPtrPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *types.Zval {
	return _getZvalPtrPtr(op_type, node, should_free, type_, executeData)
}
func GetObjZvalPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *types.Zval {
	return _getObjZvalPtr(op_type, node, should_free, type_, executeData, opline)
}
func GetObjZvalPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *types.Zval {
	return _getObjZvalPtrUndef(op_type, node, should_free, type_, executeData, opline)
}
func GetObjZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *types.Zval {
	return _getObjZvalPtrPtr(op_type, node, should_free, type_, executeData)
}
func RETURN_VALUE_USED(opline *ZendOp) bool {
	return opline.GetResultType() != IS_UNUSED
}
func ZifPass(executeData *ZendExecuteData, return_value *types.Zval) {}
func FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op *types.Zval, result *types.Zval) {
	var __container_to_free *types.Zval = free_op
	if __container_to_free != nil && __container_to_free.IsRefcounted() {
		var __ref *types.ZendRefcounted = __container_to_free.GetCounted()
		if __ref.DelRefcount() == 0 {
			var __zv *types.Zval = result
			if __zv.IsIndirect() {
				types.ZVAL_COPY(__zv, __zv.GetZv())
			}
			RcDtorFunc(__ref)
		}
	}
}
func FREE_OP(should_free *types.Zval) {
	if should_free != nil {
		ZvalPtrDtorNogc(should_free)
	}
}
func FREE_UNFETCHED_OP(type_ types.ZendUchar, var_ uint32) {
	if (type_ & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(EX_VAR(var_))
	}
}
func FREE_OP_VAR_PTR(should_free *types.Zval) {
	if should_free != nil {
		ZvalPtrDtorNogc(should_free)
	}
}
func CV_DEF_OF(i __auto__) __auto__ { return executeData.GetFunc().op_array.vars[i] }
func ZEND_VM_STACK_PAGE_ALIGNED_SIZE(size int, page_size int) int {
	return size + ZEND_VM_STACK_HEADER_SLOTS*b.SizeOf("zval") + (page_size-1) & ^(page_size-1)
}
func ZendVmStackNewPage(size int, prev ZendVmStack) ZendVmStack {
	var page ZendVmStack = ZendVmStack(Emalloc(size))
	page.SetTop(ZEND_VM_STACK_ELEMENTS(page))
	page.SetEnd((*types.Zval)((*byte)(page + size)))
	page.SetPrev(prev)
	return page
}
func ZendVmStackInit() {
	EG__().SetVmStackPageSize(ZEND_VM_STACK_PAGE_SIZE)
	EG__().SetVmStack(ZendVmStackNewPage(ZEND_VM_STACK_PAGE_SIZE, nil))
	EG__().SetVmStackTop(EG__().GetVmStack().GetTop())
	EG__().SetVmStackEnd(EG__().GetVmStack().GetEnd())
}
func ZendVmStackInitEx(page_size int) {
	/* page_size must be a power of 2 */

	b.Assert(page_size > 0 && (page_size&page_size-1) == 0)
	EG__().SetVmStackPageSize(page_size)
	EG__().SetVmStack(ZendVmStackNewPage(page_size, nil))
	EG__().SetVmStackTop(EG__().GetVmStack().GetTop())
	EG__().SetVmStackEnd(EG__().GetVmStack().GetEnd())
}
func ZendVmStackDestroy() {
	var stack ZendVmStack = EG__().GetVmStack()
	for stack != nil {
		var p ZendVmStack = stack.GetPrev()
		Efree(stack)
		stack = p
	}
}
func ZendVmStackExtend(size int) any {
	var stack ZendVmStack
	var ptr any
	stack = EG__().GetVmStack()
	stack.SetTop(EG__().GetVmStackTop())
	stack = ZendVmStackNewPage(b.CondF(size < EG__().GetVmStackPageSize()-ZEND_VM_STACK_HEADER_SLOTS*b.SizeOf("zval"), func() int { return EG__().GetVmStackPageSize() }, func() int { return ZEND_VM_STACK_PAGE_ALIGNED_SIZE(size, EG__().GetVmStackPageSize()) }), stack)
	EG__().SetVmStack(stack)
	ptr = stack.GetTop()
	EG__().SetVmStackTop(any((*byte)(ptr) + size))
	EG__().SetVmStackEnd(stack.GetEnd())
	return ptr
}
func _getZvalPtrTmp(var_ uint32, should_free *ZendFreeOp, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(var_)
	*should_free = ret
	b.Assert(ret.GetType() != types.IS_REFERENCE)
	return ret
}
func _getZvalPtrVar(var_ uint32, should_free *ZendFreeOp, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(var_)
	*should_free = ret
	return ret
}
func _getZvalPtrVarDeref(var_ uint32, should_free *ZendFreeOp, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(var_)
	*should_free = ret
	ret = types.ZVAL_DEREF(ret)
	return ret
}
func ZvalUndefinedCv(var_ uint32, executeData *ZendExecuteData) *types.Zval {
	if EG__().GetException() == nil {
		var cv *types.String = CV_DEF_OF(EX_VAR_TO_NUM(var_))
		faults.Error(faults.E_NOTICE, "Undefined variable: %s", cv.GetVal())
	}
	return EG__().GetUninitializedZval()
}
func _zvalUndefinedOp1(executeData *ZendExecuteData) *types.Zval {
	return ZvalUndefinedCv(executeData.GetOpline().op1.var_, executeData)
}
func _zvalUndefinedOp2(executeData *ZendExecuteData) *types.Zval {
	return ZvalUndefinedCv(executeData.GetOpline().op2.var_, executeData)
}
func ZVAL_UNDEFINED_OP1() *types.Zval { return _zvalUndefinedOp1(executeData) }
func ZVAL_UNDEFINED_OP2() *types.Zval { return _zvalUndefinedOp2(executeData) }
func _getZvalCvLookup(ptr *types.Zval, var_ uint32, type_ int, executeData *ZendExecuteData) *types.Zval {
	switch type_ {
	case BP_VAR_R:
		fallthrough
	case BP_VAR_UNSET:
		ptr = ZvalUndefinedCv(var_, executeData)
	case BP_VAR_IS:
		ptr = EG__().GetUninitializedZval()
	case BP_VAR_RW:
		ZvalUndefinedCv(var_, executeData)
		fallthrough
	case BP_VAR_W:
		ptr.SetNull()
	}
	return ptr
}
func _getZvalPtrCv(var_ uint32, type_ int, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(var_)
	if ret.IsUndef() {
		if type_ == BP_VAR_W {
			ret.SetNull()
		} else {
			return _getZvalCvLookup(ret, var_, type_, executeData)
		}
	}
	return ret
}
func _getZvalPtrCvDeref(var_ uint32, type_ int, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(var_)
	if ret.IsUndef() {
		if type_ == BP_VAR_W {
			ret.SetNull()
			return ret
		} else {
			return _getZvalCvLookup(ret, var_, type_, executeData)
		}
	}
	ret = types.ZVAL_DEREF(ret)
	return ret
}
func _get_zval_ptr_cv_BP_VAR_R(var_ uint32, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(var_)
	if ret.IsUndef() {
		return ZvalUndefinedCv(var_, executeData)
	}
	return ret
}
func _get_zval_ptr_cv_deref_BP_VAR_R(var_ uint32, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(var_)
	if ret.IsUndef() {
		return ZvalUndefinedCv(var_, executeData)
	}
	ret = types.ZVAL_DEREF(ret)
	return ret
}
