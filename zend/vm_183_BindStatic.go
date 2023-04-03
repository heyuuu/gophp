package zend

func ZEND_BIND_STATIC_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ht *types.Array
	var value *types.Zval
	var variable_ptr *types.Zval
	variable_ptr = opline.GetOp1Zval()
	IZvalPtrDtor(variable_ptr)
	ht = ZEND_MAP_PTR_GET(executeData.GetFunc().op_array.static_variables_ptr)
	if ht == nil {
		b.Assert((executeData.GetFunc().op_array.fn_flags & (AccImmutable | AccPreloaded)) != 0)
		ht = types.ZendArrayDup(executeData.GetFunc().op_array.static_variables)
		ZEND_MAP_PTR_SET(executeData.GetFunc().op_array.static_variables_ptr, ht)
	} else if ht.GetRefcount() > 1 {
		if (ht.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
			ht.DelRefcount()
		}
		ht = types.ZendArrayDup(ht)
		ZEND_MAP_PTR_SET(executeData.GetFunc().op_array.static_variables_ptr, ht)
	}
	value = (*types.Zval)((*byte)(ht.GetArData() + (opline.GetExtendedValue() & ^(ZEND_BIND_REF | ZEND_BIND_IMPLICIT))))
	if (opline.GetExtendedValue() & ZEND_BIND_REF) != 0 {
		if value.IsConstant() {
			if ZvalUpdateConstantEx(value, executeData.GetFunc().op_array.scope) != types.SUCCESS {
				variable_ptr.SetNull()
				return 0
			}
		}
		if !(value.IsReference()) {
			var ref *types.ZendReference = (*types.ZendReference)(Emalloc(b.SizeOf("zend_reference")))
			ref.SetRefcount(2)
			ref.GetGcTypeInfo() = types.IS_REFERENCE
			types.ZVAL_COPY_VALUE(ref.GetVal(), value)
			ref.GetSources().SetPtr(nil)
			value.SetTypeReference()
			variable_ptr.SetReference(ref)
		} else {
			value.AddRefcount()
			variable_ptr.SetReference(value.GetRef())
		}
	} else {
		types.ZVAL_COPY(variable_ptr, value)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
