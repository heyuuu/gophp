package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_FE_FETCH_R_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var array *types.Zval
	var value *types.Zval
	var fe_ht *types.Array
	var pos types.ArrayPosition
	var p *types.Bucket
	array = opline.Op1()
	if array.IsArray() {
		fe_ht = array.Array()
		pos = array.GetFePos()
		p = fe_ht.Bucket(pos)
		for true {
			if pos >= fe_ht.GetNNumUsed() {

				/* reached end of iteration */

			fe_fetch_r_exit:
				ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
				return 0
			}
			value = p.GetVal().DeIndirect()
			if value.IsNotUndef() {
				break
			}

			pos++
			p++
		}
		array.SetFePos(pos + 1)
		if RETURN_VALUE_USED(opline) {
			if p.GetKey() == nil {
				opline.Result().SetLong(p.GetH())
			} else {
				opline.Result().SetStringVal(p.GetKey().GetStr())
			}
		}
	} else {
		var iter *ZendObjectIterator
		b.Assert(array.IsObject())
		if lang.Assign(&iter, ZendIteratorUnwrap(array)) == nil {

			/* plain object */

			fe_ht = types.Z_OBJPROP_P(array)
			pos = types.ZendHashIteratorPos(array.GetFeIterIdx(), fe_ht)
			p = fe_ht.Bucket(pos)
			for true {
				if pos >= fe_ht.GetNNumUsed() {

					/* reached end of iteration */

					goto fe_fetch_r_exit

					/* reached end of iteration */

				}
				value = p.GetVal()

				if value.IsNotUndef() {
					if value.IsIndirect() {
						value = value.Indirect()
						if value.IsNotUndef() && ZendCheckPropertyAccess(array.Object(), p.GetKey(), false) == types.SUCCESS {
							break
						}
					} else if types.Z_OBJCE_P(array).GetDefaultPropertiesCount() == 0 || p.GetKey() == nil || ZendCheckPropertyAccess(array.Object(), p.GetKey(), true) == types.SUCCESS {
						break
					}
				}
				pos++
				p++
			}
			if RETURN_VALUE_USED(opline) {
				if p.GetKey() == nil {
					opline.Result().SetLong(p.GetH())
				} else if p.GetKey().GetStr()[0] {
					opline.Result().SetStringVal(p.GetKey().GetStr())
				} else {
					var class_name *byte
					var prop_name *byte
					var prop_name_len int
					ZendUnmanglePropertyNameEx(p.GetKey(), &class_name, &prop_name, &prop_name_len)
					opline.Result().SetStringVal(b.CastStr(prop_name, prop_name_len))
				}
			}
			EG__().GetHtIterators()[types.Z_FE_ITER_P(array)].SetPos(pos + 1)
		} else {
			if lang.PreInc(&(iter.GetIndex())) > 0 {

				/* This could cause an endless loop if index becomes zero again.
				 * In case that ever happens we need an additional flag. */

				iter.GetFuncs().GetMoveForward()(iter)
				if EG__().GetException() != nil {
					UNDEF_RESULT()
					return 0
				}
				if iter.GetFuncs().GetValid()(iter) == types.FAILURE {

					/* reached end of iteration */

					if EG__().GetException() != nil {
						UNDEF_RESULT()
						return 0
					}
					goto fe_fetch_r_exit
				}
			}
			value = iter.GetFuncs().GetGetCurrentData()(iter)
			if EG__().GetException() != nil {
				UNDEF_RESULT()
				return 0
			}
			if value == nil {

				/* failure in get_current_data */

				goto fe_fetch_r_exit

				/* failure in get_current_data */

			}
			if RETURN_VALUE_USED(opline) {
				if iter.GetFuncs().GetGetCurrentKey() != nil {
					iter.GetFuncs().GetGetCurrentKey()(iter, opline.Result())
					if EG__().GetException() != nil {
						UNDEF_RESULT()
						return 0
					}
				} else {
					opline.Result().SetLong(iter.GetIndex())
				}
			}
		}
	}
	if opline.GetOp2Type() == IS_CV {
		var variable_ptr *types.Zval = opline.Op2()
		ZendAssignToVariable(variable_ptr, value, executeData.IsCallUseStrictTypes())
	} else {
		var res *types.Zval = opline.Op2()
		res.CopyValueFrom(value)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
