package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
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
func ZendCopyToVariable(variable_ptr *types.Zval, value *types.Zval) {
	variable_ptr.CopyValueFrom(value)
}
func ZendAssignToVariable(variable_ptr *types.Zval, value *types.Zval, value_type uint8, strict types.ZendBool) *types.Zval {
	var ref *types.ZendRefcounted = nil
	if value.IsReference() {
		ref = value.RefCounted()
		value = types.Z_REFVAL_P(value)
	}
	for {
		if variable_ptr.IsRefcounted() {
			var garbage *types.ZendRefcounted
			if variable_ptr.IsReference() {
				if ZEND_REF_HAS_TYPE_SOURCES(variable_ptr.Reference()) {
					return ZendAssignToTypedRef(variable_ptr, value, value_type, strict, ref)
				}
				variable_ptr = types.Z_REFVAL_P(variable_ptr)
				if !(variable_ptr.IsRefcounted()) {
					break
				}
			}
			if variable_ptr.IsObject() && variable_ptr.Object().CanSet() {
				variable_ptr.Object().Set(variable_ptr, value)
				return variable_ptr
			}
			garbage = variable_ptr.RefCounted()
			ZendCopyToVariable(variable_ptr, value)
			return variable_ptr
		}
		break
	}
	ZendCopyToVariable(variable_ptr, value)
	return variable_ptr
}
func ZendVmStackFreeExtraArgsEx(call_info uint32, call *ZendExecuteData) {
	if (call_info & ZEND_CALL_FREE_EXTRA_ARGS) != 0 {
		var count uint32 = call.NumArgs() - call.GetFunc().GetOpArray().GetNumArgs()
		var p *types.Zval = call.VarNum(call.GetFunc().GetOpArray().GetLastVar() + call.GetFunc().GetOpArray().GetT())
		for {
			if p.IsRefcounted() {
				//var r *types.ZendRefcounted = p.RefCounted()
				//if r.DelRefcount() == 0 {
				//	p.SetNull()
				//	//RcDtorFunc(r)
				//} else {
				//	//GcCheckPossibleRoot(r)
				//}
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
			//r := p.RefCounted()
			//if r.DelRefcount() == 0 {
			p.SetNull()
			//RcDtorFunc(r)
			//}
		}
	}
}
func ZendVmStackFreeCallFrame(call *ZendExecuteData) {
	EG__().VmStack().PopCheck(call)
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
func GetZvalPtrUndef(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *types.Zval {
	return _getZvalPtrUndef(op_type, node, should_free, type_, executeData, opline)
}
func GetOpDataZvalPtrR(op_type int, node ZnodeOp, should_free *ZendFreeOp) *types.Zval {
	return _getOpDataZvalPtrR(op_type, node, should_free, executeData, opline)
}
func GetZvalPtrPtr(op_type int, node ZnodeOp, should_free *ZendFreeOp, type_ int) *types.Zval {
	return _getZvalPtrPtr(op_type, node, should_free, type_, executeData)
}
func RETURN_VALUE_USED(opline *ZendOp) bool {
	return opline.GetResultType() != IS_UNUSED
}
func ZifPass(executeData *ZendExecuteData, return_value *types.Zval) {}
func FREE_VAR_PTR_AND_EXTRACT_RESULT_IF_NECESSARY(free_op *types.Zval, result *types.Zval) {
	var __container_to_free *types.Zval = free_op
	if __container_to_free != nil && __container_to_free.IsRefcounted() {
		//var __ref *types.ZendRefcounted = __container_to_free.RefCounted()
		//if __ref.DelRefcount() == 0 {
		//	var __zv *types.Zval = result
		//	if __zv.IsIndirect() {
		//		types.ZVAL_COPY(__zv, __zv.Indirect())
		//	}
		//	//RcDtorFunc(__ref)
		//}
	}
}
func CV_DEF_OF(i __auto__) __auto__ { return executeData.GetFunc().GetOpArray().vars[i] }

func ZendVmStackInit() {
	EG__().VmStack().Reset()
}
func ZendVmStackDestroy() {
	EG__().VmStack().Reset()
}
func _getZvalPtrTmp(var_ uint32, should_free *ZendFreeOp, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(executeData, var_)
	*should_free = ret
	b.Assert(ret.GetType() != types.IS_REFERENCE)
	return ret
}
func _getZvalPtrVar(var_ uint32, should_free *ZendFreeOp, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(executeData, var_)
	*should_free = ret
	return ret
}
func _getZvalPtrVarDeref(var_ uint32, should_free *ZendFreeOp, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(executeData, var_)
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
func ZVAL_UNDEFINED_OP1(executeData *ZendExecuteData) *types.Zval {
	return _zvalUndefinedOp1(executeData)
}
func ZVAL_UNDEFINED_OP2(executeData *ZendExecuteData) *types.Zval {
	return _zvalUndefinedOp2(executeData)
}
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
	var ret *types.Zval = EX_VAR(executeData, var_)
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
	var ret *types.Zval = EX_VAR(executeData, var_)
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
	var ret *types.Zval = EX_VAR(executeData, var_)
	if ret.IsUndef() {
		return ZvalUndefinedCv(var_, executeData)
	}
	return ret
}
func _get_zval_ptr_cv_deref_BP_VAR_R(var_ uint32, executeData *ZendExecuteData) *types.Zval {
	var ret *types.Zval = EX_VAR(executeData, var_)
	if ret.IsUndef() {
		return ZvalUndefinedCv(var_, executeData)
	}
	ret = types.ZVAL_DEREF(ret)
	return ret
}
