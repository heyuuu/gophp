package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_FE_FETCH_RW_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var array *types2.Zval
	var value *types2.Zval
	var value_type uint32
	var fe_ht *types2.Array
	var pos types2.ArrayPosition
	var p *types2.Bucket
	array = opline.Op1()
	array = types2.ZVAL_DEREF(array)
	if array.IsArray() {
		pos = types2.ZendHashIteratorPosEx(opline.Op1().GetFeIterIdx(), array)
		fe_ht = array.GetArr()
		p = fe_ht.Bucket(pos)
		for true {
			if pos >= fe_ht.GetNNumUsed() {

				/* reached end of iteration */

				goto fe_fetch_w_exit

				/* reached end of iteration */

			}
			value = p.GetVal()
			value_type = value.GetTypeInfo()
			if value_type != types2.IS_UNDEF {
				if value_type == types2.IS_INDIRECT {
					value = value.GetZv()
					value_type = value.GetTypeInfo()
					if value_type != types2.IS_UNDEF {
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
				opline.Result().SetStringCopy(p.GetKey())
			}
		}
		EG__().GetHtIterators()[types.Z_FE_ITER_P(opline.Op1())].SetPos(pos + 1)
	} else if array.IsObject() {
		var iter *ZendObjectIterator
		if b.Assign(&iter, ZendIteratorUnwrap(array)) == nil {

			/* plain object */

			fe_ht = types2.Z_OBJPROP_P(array)
			pos = types2.ZendHashIteratorPos(opline.Op1().GetFeIterIdx(), fe_ht)
			p = fe_ht.Bucket(pos)
			for true {
				if pos >= fe_ht.GetNNumUsed() {

					/* reached end of iteration */

					goto fe_fetch_w_exit

					/* reached end of iteration */

				}
				value = p.GetVal()
				value_type = value.GetTypeInfo()
				if value_type != types2.IS_UNDEF {
					if value_type == types2.IS_INDIRECT {
						value = value.GetZv()
						value_type = value.GetTypeInfo()
						if value_type != types2.IS_UNDEF && ZendCheckPropertyAccess(array.GetObj(), p.GetKey(), 0) == types2.SUCCESS {
							if (value_type & types2.Z_TYPE_MASK) != types2.IS_REFERENCE {
								var prop_info *ZendPropertyInfo = ZendGetTypedPropertyInfoForSlot(array.GetObj(), value)
								if prop_info != nil {
									value.SetNewRef(value)
									ZEND_REF_ADD_TYPE_SOURCE(value.Reference(), prop_info)
									value_type = types2.IS_REFERENCE_EX
								}
							}
							break
						}
					} else if types2.Z_OBJCE_P(array).GetDefaultPropertiesCount() == 0 || p.GetKey() == nil || ZendCheckPropertyAccess(array.GetObj(), p.GetKey(), 1) == types2.SUCCESS {
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
					opline.Result().SetStringCopy(p.GetKey())
				} else {
					var class_name *byte
					var prop_name *byte
					var prop_name_len int
					ZendUnmanglePropertyNameEx(p.GetKey(), &class_name, &prop_name, &prop_name_len)
					opline.Result().SetStringVal(b.CastStr(prop_name, prop_name_len))
				}
			}
			EG__().GetHtIterators()[types.Z_FE_ITER_P(opline.Op1())].SetPos(pos + 1)
		} else {
			if b.PreInc(&(iter.GetIndex())) > 0 {

				/* This could cause an endless loop if index becomes zero again.
				 * In case that ever happens we need an additional flag. */

				iter.GetFuncs().GetMoveForward()(iter)
				if EG__().GetException() != nil {
					UNDEF_RESULT()
					return 0
				}
				if iter.GetFuncs().GetValid()(iter) == types2.FAILURE {

					/* reached end of iteration */

					if EG__().GetException() != nil {
						UNDEF_RESULT()
						return 0
					}
					goto fe_fetch_w_exit
				}
			}
			value = iter.GetFuncs().GetGetCurrentData()(iter)
			if EG__().GetException() != nil {
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
					if EG__().GetException() != nil {
						UNDEF_RESULT()
						return 0
					}
				} else {
					opline.Result().SetLong(iter.GetIndex())
				}
			}
			value_type = value.GetTypeInfo()
		}
	} else {
		faults.Error(faults.E_WARNING, "Invalid argument supplied for foreach()")
		if EG__().GetException() != nil {
			UNDEF_RESULT()
			return 0
		}
	fe_fetch_w_exit:
		ZEND_VM_SET_RELATIVE_OPCODE(executeData, opline, opline.GetExtendedValue())
		return 0
	}
	if (value_type & types2.Z_TYPE_MASK) != types2.IS_REFERENCE {
		var gc *types2.ZendRefcounted = value.GetCounted()
		var ref *types2.Zval
		value.SetNewEmptyRef()
		ref = types2.Z_REFVAL_P(value)
		types.ZVAL_COPY_VALUE_EX(ref, value, gc, value_type)
	}
	if opline.GetOp2Type() == IS_CV {
		var variable_ptr *types2.Zval = opline.Op2()
		if variable_ptr != value {
			var ref *types2.ZendReference
			ref = value.Reference()
			// 			ref.AddRefcount()
			// IZvalPtrDtor(variable_ptr)
			variable_ptr.SetReference(ref)
		}
	} else {
		// 		value.AddRefcount()
		opline.Op2().SetReference(value.Reference())
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
