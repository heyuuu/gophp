package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_BIND_STATIC_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var ht *types.Array
	var value *types.Zval
	var variable_ptr *types.Zval
	variable_ptr = opline.Op1()
	ht = executeData.GetFunc().GetOpArray().GetStaticVariablesPtr()
	if ht == nil {
		b.Assert((executeData.GetFunc().GetOpArray().GetFnFlags() & (types.AccImmutable | types.AccPreloaded)) != 0)
		ht = types.ZendArrayDup(executeData.GetFunc().GetOpArray().static_variables)
		executeData.GetFunc().GetOpArray().SetStaticVariablesPtr(ht)
	} else {
		ht = ht.LazyDup()
		executeData.GetFunc().GetOpArray().SetStaticVariablesPtr(ht)
	}
	itemPos := opline.GetExtendedValue() &^ (ZEND_BIND_REF | ZEND_BIND_IMPLICIT)
	value = ht.PosValue(itemPos)
	if (opline.GetExtendedValue() & ZEND_BIND_REF) != 0 {
		if value.IsConstantAst() {
			if ZvalUpdateConstantEx(value, executeData.GetFunc().GetOpArray().scope) != types.SUCCESS {
				variable_ptr.SetNull()
				return 0
			}
		}
		if !(value.IsReference()) {
			var ref = types.NewZendReference(value)
			value.SetReference(ref)
			variable_ptr.SetReference(ref)
		} else {
			variable_ptr.SetReference(value.Reference())
		}
	} else {
		types.ZVAL_COPY(variable_ptr, value)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
