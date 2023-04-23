package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_FUNC_GET_ARGS_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ht *types.Array
	var arg_count uint32
	var result_size uint32
	var skip uint32
	arg_count = executeData.NumArgs()
	{
		skip = opline.Const1().Long()
		if arg_count < skip {
			result_size = 0
		} else {
			result_size = arg_count - skip
		}
	}

	if result_size != 0 {
		var first_extra_arg uint32 = executeData.GetFunc().GetOpArray().GetNumArgs()
		ht = types.NewArray(result_size)
		opline.Result().SetArray(ht)
		var p *types.Zval
		var q *types.Zval
		var i uint32 = skip
		p = executeData.VarNum(i)
		if arg_count > first_extra_arg {
			for i < first_extra_arg {
				q = p
				if !q.IsUndef() {
					q = types.ZVAL_DEREF(q)
					ht.Append(q)
				} else {
					ht.Append(types.NewZvalNull())
				}
				p++
				i++
			}
			if skip < first_extra_arg {
				skip = 0
			} else {
				skip -= first_extra_arg
			}
			p = executeData.VarNum(executeData.GetFunc().GetOpArray().last_var + executeData.GetFunc().GetOpArray().T + skip)
		}
		for i < arg_count {
			q = p
			if !q.IsUndef() {
				q = types.ZVAL_DEREF(q)
				ht.Append(q)
			} else {
				ht.Append(types.NewZvalNull())
			}
			p++
			i++
		}
	} else {
		opline.Result().SetEmptyArray()
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_FUNC_GET_ARGS_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ht *types.Array
	var arg_count uint32
	var result_size uint32
	var skip uint32
	arg_count = executeData.NumArgs()

	{
		skip = 0
		result_size = arg_count
	}
	if result_size != 0 {
		var first_extra_arg uint32 = executeData.GetFunc().GetOpArray().GetNumArgs()
		ht = types.NewArray(result_size)
		opline.Result().SetArray(ht)
		var p *types.Zval
		var q *types.Zval
		var i uint32 = skip
		p = executeData.VarNum(i)
		if arg_count > first_extra_arg {
			for i < first_extra_arg {
				q = p
				if !q.IsUndef() {
					q = types.ZVAL_DEREF(q)
					ht.Append(q)
				} else {
					ht.Append(types.NewZvalNull())
				}
				p++
				i++
			}
			if skip < first_extra_arg {
				skip = 0
			} else {
				skip -= first_extra_arg
			}
			p = executeData.VarNum(executeData.GetFunc().GetOpArray().last_var + executeData.GetFunc().GetOpArray().T + skip)
		}
		for i < arg_count {
			q = p
			if !q.IsUndef() {
				q = types.ZVAL_DEREF(q)
				ht.Append(q)
			} else {
				ht.Append(types.NewZvalNull())
			}
			p++
			i++
		}
		ht.SetNNumOfElements(result_size)
	} else {
		opline.Result().SetEmptyArray()
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
