package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_FE_FETCH_RW_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var array *types.Zval
	var value *types.Zval
	var value_type types.ZvalType
	var fe_ht *types.Array
	var pos types.ArrayPosition
	var p *types.Bucket
	array = opline.Op1()
	array = types.ZVAL_DEREF(array)
	if array.IsArray() {
		pos = types.ZendHashIteratorPos(opline.Op1().GetFeIterIdx(), array.Array())
		fe_ht = array.GetArr()
		p = fe_ht.Bucket(pos)
		for true {
			if pos >= fe_ht.GetNNumUsed() {

				/* reached end of iteration */

				goto fe_fetch_w_exit

				/* reached end of iteration */

			}
			value = p.GetVal()
			value_type = value.Type()
			if value.IsNotUndef() {
				if value.IsIndirect() {
					value = value.Indirect()
					value_type = value.Type()
					if value.IsNotUndef() {
						break
					}
				} else {
					break
				}
			}

			pos++
			p++
		}
		if RETURN_VALUE_USED(opline) {
			if p.GetKey() == nil {
				opline.Result().SetLong(p.GetH())
			} else {
				opline.Result().SetString(p.GetKey().GetStr())
			}
		}
		EG__().GetHtIterators()[types.Z_FE_ITER_P(opline.Op1())].SetPos(pos + 1)
	} else if array.IsObject() {
		var iter *ZendObjectIterator
		if lang.Assign(&iter, ZendIteratorUnwrap(array)) == nil {

			/* plain object */

			fe_ht = types.Z_OBJPROP_P(array)
			pos = types.ZendHashIteratorPos(opline.Op1().GetFeIterIdx(), fe_ht)
			p = fe_ht.Bucket(pos)
			for true {
				if pos >= fe_ht.GetNNumUsed() {

					/* reached end of iteration */

					goto fe_fetch_w_exit

					/* reached end of iteration */

				}
				value = p.GetVal()
				value_type = value.Type()
				if value.IsNotUndef() {
					if value.IsIndirect() {
						value = value.Indirect()
						value_type = value.Type()
						if value.IsNotUndef() && ZendCheckPropertyAccess(array.Object(), p.GetKey(), false) == types.SUCCESS {
							if value.IsRef() {
								var prop_info *types.PropertyInfo = ZendGetTypedPropertyInfoForSlot(array.Object(), value)
								if prop_info != nil {
									value.SetNewRef(value)
									ZEND_REF_ADD_TYPE_SOURCE(value.Ref(), prop_info)
									value_type = types.IsRef
								}
							}
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
					opline.Result().SetString(p.GetKey().GetStr())
				} else {
					var class_name *byte
					var prop_name *byte
					var prop_name_len int
					ZendUnmanglePropertyNameEx(p.GetKey(), &class_name, &prop_name, &prop_name_len)
					opline.Result().SetString(b.CastStr(prop_name, prop_name_len))
				}
			}
			EG__().GetHtIterators()[types.Z_FE_ITER_P(opline.Op1())].SetPos(pos + 1)
		} else {
			if lang.PreInc(&(iter.GetIndex())) > 0 {

				/* This could cause an endless loop if index becomes zero again.
				 * In case that ever happens we need an additional flag. */

				iter.GetFuncs().GetMoveForward()(iter)
				if EG__().HasException() {
					UNDEF_RESULT()
					return 0
				}
				if iter.GetFuncs().GetValid()(iter) == types.FAILURE {

					/* reached end of iteration */

					if EG__().HasException() {
						UNDEF_RESULT()
						return 0
					}
					goto fe_fetch_w_exit
				}
			}
			value = iter.GetFuncs().GetGetCurrentData()(iter)
			if EG__().HasException() {
				UNDEF_RESULT()
				return 0
			}
			if value == nil {

				/* failure in get_current_data */

				goto fe_fetch_w_exit

				/* failure in get_current_data */

			}
			if RETURN_VALUE_USED(opline) {
				if iter.GetFuncs().GetGetCurrentKey() != nil {
					iter.GetFuncs().GetGetCurrentKey()(iter, opline.Result())
					if EG__().HasException() {
						UNDEF_RESULT()
						return 0
					}
				} else {
					opline.Result().SetLong(iter.GetIndex())
				}
			}
			value_type = value.Type()
		}
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		if EG__().HasException() {
			UNDEF_RESULT()
			return 0
		}
	fe_fetch_w_exit:
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	}
	if value_type != types.IsRef {
		var ref *types.Zval
		value.SetNewEmptyRef()
		ref = types.Z_REFVAL_P(value)
		ref.CopyValueFrom(value)
	}
	if opline.GetOp2Type() == IS_CV {
		var variable_ptr *types.Zval = opline.Op2()
		if variable_ptr != value {
			var ref *types.Reference
			ref = value.Ref()
			variable_ptr.SetReference(ref)
		}
	} else {
		opline.Op2().SetReference(value.Ref())
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
