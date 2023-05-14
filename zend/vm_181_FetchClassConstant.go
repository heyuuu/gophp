package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_FETCH_CLASS_CONSTANT_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var c *types.ClassConstant
	var value *types.Zval
	var opline *ZendOp = executeData.GetOpline()
	for {
		{
			if CACHED_PTR(opline.GetExtendedValue() + b.SizeOf("void *")) {
				value = CACHED_PTR(opline.GetExtendedValue() + b.SizeOf("void *"))
				break
			} else if CACHED_PTR(opline.GetExtendedValue()) {
				ce = CACHED_PTR(opline.GetExtendedValue())
			} else {
				ce = ZendFetchClassByName(opline.Const1().String(), (opline.Const1() + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
				if ce == nil {
					b.Assert(EG__().GetException() != nil)
					opline.Result().SetUndef()
					return 0
				}
			}
		}

		c = ce.ConstantsTable().Get(opline.Const2().StringVal())
		if c != nil {
			scope = executeData.GetFunc().GetOpArray().GetScope()
			if !ZendVerifyConstAccess(c, scope) {
				faults.ThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), ce.GetName().GetVal(), opline.Const2().String().GetVal())
				opline.Result().SetUndef()
				return 0
			}
			value = c.GetValue()
			if value.IsConstantAst() {
				ZvalUpdateConstantEx(value, c.GetCe())
				if EG__().GetException() != nil {
					opline.Result().SetUndef()
					return 0
				}
			}
			CACHE_POLYMORPHIC_PTR(opline.GetExtendedValue(), ce, value)
		} else {
			faults.ThrowError(nil, "Undefined class constant '%s'", opline.Const2().String().GetVal())
			opline.Result().SetUndef()
			return 0
		}
		break
	}
	types.ZVAL_COPY_OR_DUP(opline.Result(), value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_FETCH_CLASS_CONSTANT_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var c *types.ClassConstant
	var value *types.Zval
	var zv *types.Zval
	var opline *ZendOp = executeData.GetOpline()
	for {

		{

			{
				ce = opline.Op1().Class()
			}
			if CACHED_PTR(opline.GetExtendedValue()) == ce {
				value = CACHED_PTR(opline.GetExtendedValue() + b.SizeOf("void *"))
				break
			}
		}
		c = ce.ConstantsTable().Get(opline.Const2().String().GetStr())
		if c != nil {
			scope = executeData.GetFunc().GetOpArray().GetScope()
			if !ZendVerifyConstAccess(c, scope) {
				faults.ThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), ce.GetName().GetVal(), opline.Const2().String().GetVal())
				opline.Result().SetUndef()
				return 0
			}
			value = c.GetValue()
			if value.IsConstantAst() {
				ZvalUpdateConstantEx(value, c.GetCe())
				if EG__().GetException() != nil {
					opline.Result().SetUndef()
					return 0
				}
			}
			CACHE_POLYMORPHIC_PTR(opline.GetExtendedValue(), ce, value)
		} else {
			faults.ThrowError(nil, "Undefined class constant '%s'", opline.Const2().String().GetVal())
			opline.Result().SetUndef()
			return 0
		}
		break
	}
	types.ZVAL_COPY_OR_DUP(opline.Result(), value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_FETCH_CLASS_CONSTANT_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	var c *types.ClassConstant
	var value *types.Zval
	var opline *ZendOp = executeData.GetOpline()
	for {

		{
			{
				ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
				if ce == nil {
					b.Assert(EG__().GetException() != nil)
					opline.Result().SetUndef()
					return 0
				}
			}

			if CACHED_PTR(opline.GetExtendedValue()) == ce {
				value = CACHED_PTR(opline.GetExtendedValue() + b.SizeOf("void *"))
				break
			}
		}
		c = ce.ConstantsTable().Get(opline.Const2().String().GetStr())
		if c != nil {
			scope = executeData.GetFunc().GetOpArray().GetScope()
			if !ZendVerifyConstAccess(c, scope) {
				faults.ThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), ce.GetName().GetVal(), opline.Const2().String().GetVal())
				opline.Result().SetUndef()
				return 0
			}
			value = c.GetValue()
			if value.IsConstantAst() {
				ZvalUpdateConstantEx(value, c.GetCe())
				if EG__().GetException() != nil {
					opline.Result().SetUndef()
					return 0
				}
			}
			CACHE_POLYMORPHIC_PTR(opline.GetExtendedValue(), ce, value)
		} else {
			faults.ThrowError(nil, "Undefined class constant '%s'", opline.Const2().String().GetVal())
			opline.Result().SetUndef()
			return 0
		}
		break
	}
	types.ZVAL_COPY_OR_DUP(opline.Result(), value)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
