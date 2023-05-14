package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_BIND_STATIC_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ht *types.Array
	var value *types.Zval
	var variable_ptr *types.Zval
	variable_ptr = opline.Op1()
	// IZvalPtrDtor(variable_ptr)
	ht = executeData.GetFunc().GetOpArray().GetStaticVariablesPtr()
	if ht == nil {
		b.Assert((executeData.GetFunc().GetOpArray().GetFnFlags() & (types.AccImmutable | types.AccPreloaded)) != 0)
		ht = types.ZendArrayDup(executeData.GetFunc().GetOpArray().static_variables)
		executeData.GetFunc().GetOpArray().SetStaticVariablesPtr(ht)
	} else {
		ht = ht.LazyDup()
		executeData.GetFunc().GetOpArray().SetStaticVariablesPtr(ht)
	}
	value = (*types.Zval)((*byte)(ht.GetArData() + (opline.GetExtendedValue() & ^(ZEND_BIND_REF | ZEND_BIND_IMPLICIT))))
	if (opline.GetExtendedValue() & ZEND_BIND_REF) != 0 {
		if value.IsConstantAst() {
			if ZvalUpdateConstantEx(value, executeData.GetFunc().GetOpArray().scope) != types.SUCCESS {
				variable_ptr.SetNull()
				return 0
			}
		}
		if !(value.IsReference()) {
			var ref = (*types.ZendReference)(Emalloc(b.SizeOf("zend_reference")))
			types.ZVAL_COPY_VALUE(ref.GetVal(), value)
			ref.GetSources().SetPtr(nil)
			value.SetTypeReference()
			variable_ptr.SetReference(ref)
		} else {
			// 			value.AddRefcount()
			variable_ptr.SetReference(value.Reference())
			variable_ptr.SetReference(value.Reference())
		}
	} else {
		types.ZVAL_COPY(variable_ptr, value)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
