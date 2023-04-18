package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_BIND_STATIC_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ht *types2.Array
	var value *types2.Zval
	var variable_ptr *types2.Zval
	variable_ptr = opline.Op1()
	// IZvalPtrDtor(variable_ptr)
	ht = executeData.GetFunc().GetOpArray().GetStaticVariablesPtr()
	if ht == nil {
		b.Assert((executeData.GetFunc().GetOpArray().GetFnFlags() & (AccImmutable | AccPreloaded)) != 0)
		ht = types2.ZendArrayDup(executeData.GetFunc().GetOpArray().static_variables)
		executeData.GetFunc().GetOpArray().SetStaticVariablesPtr(ht)
	} else if ht.GetRefcount() > 1 {
		if (ht.GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) == 0 {
			ht.DelRefcount()
		}
		ht = types2.ZendArrayDup(ht)
		executeData.GetFunc().GetOpArray().SetStaticVariablesPtr(ht)
	}
	value = (*types2.Zval)((*byte)(ht.GetArData() + (opline.GetExtendedValue() & ^(ZEND_BIND_REF | ZEND_BIND_IMPLICIT))))
	if (opline.GetExtendedValue() & ZEND_BIND_REF) != 0 {
		if value.IsConstant() {
			if ZvalUpdateConstantEx(value, executeData.GetFunc().GetOpArray().scope) != types2.SUCCESS {
				variable_ptr.SetNull()
				return 0
			}
		}
		if !(value.IsReference()) {
			var ref *types2.ZendReference = (*types2.ZendReference)(Emalloc(b.SizeOf("zend_reference")))
			ref.SetRefcount(2)
			ref.GetGcTypeInfo() = types2.IS_REFERENCE
			types2.ZVAL_COPY_VALUE(ref.GetVal(), value)
			ref.GetSources().SetPtr(nil)
			value.SetTypeReference()
			variable_ptr.SetReference(ref)
		} else {
			// 			value.AddRefcount()
			variable_ptr.SetReference(value.Reference())
			variable_ptr.SetReference(value.Reference())
		}
	} else {
		types2.ZVAL_COPY(variable_ptr, value)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
