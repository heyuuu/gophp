package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_INLINE_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types2.Zval
	var offset *types2.Zval
	var cache_slot *any = nil
	container = &(executeData.GetThis())
	if container.IsUndef() {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	offset = opline.Const2()
	/* here we are sure we are dealing with an object */

	var zobj *types2.ZendObject = container.Object()
	var retval *types2.Zval
	{
		cache_slot = CACHE_ADDR(opline.GetExtendedValue() & ^ZEND_FETCH_REF)
		if zobj.GetCe() == CACHED_PTR_EX(cache_slot) {
			var prop_offset uintPtr = uintPtr(CACHED_PTR_EX(cache_slot + 1))
			if IS_VALID_PROPERTY_OFFSET(prop_offset) {
				retval = OBJ_PROP(zobj, prop_offset)
				if !retval.IsUndef() {

					{
					fetch_obj_r_fast_copy:
						types2.ZVAL_COPY_DEREF(opline.Result(), retval)
						return ZEND_VM_NEXT_OPCODE(executeData, opline)
					}
				}
			} else if zobj.GetProperties() != nil {
				if !(IS_UNKNOWN_DYNAMIC_PROPERTY_OFFSET(prop_offset)) {
					var idx uintPtr = ZEND_DECODE_DYN_PROP_OFFSET(prop_offset)
					if idx < zobj.GetProperties().GetNNumUsed()*b.SizeOf("Bucket") {
						var p *types2.Bucket = (*types2.Bucket)((*byte)(zobj.GetProperties().Bucket(idx)))
						if p.GetVal().IsNotUndef() && (p.GetKey() == offset.String() || (p.IsStrKey() && p.StrKey() == offset.StringVal())) {
							retval = p.GetVal()

							{
								goto fetch_obj_r_fast_copy
							}
						}
					}
					CACHE_PTR_EX(cache_slot+1, any(ZEND_DYNAMIC_PROPERTY_OFFSET))
				}
				retval = zobj.GetProperties().KeyFind(offset.String().GetStr())
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					CACHE_PTR_EX(cache_slot+1, any(ZEND_ENCODE_DYN_PROP_OFFSET(idx)))

					{
						goto fetch_obj_r_fast_copy
					}
				}
			}
		}
	}

	retval = zobj.GetHandlers().GetReadProperty()(container, offset, BP_VAR_R, cache_slot, opline.Result())
	if retval != opline.Result() {
	fetch_obj_r_copy:
		types2.ZVAL_COPY_DEREF(opline.Result(), retval)
	} else if retval.IsReference() {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
