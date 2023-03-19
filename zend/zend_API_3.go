package zend

import (
	"sik/zend/argparse"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZendParseParametersEx(flags int, num_args int, type_spec string, args ...any) int {
	return argparse.ParseVaArgs(num_args, type_spec, args, flags)
}
func ZendParseParameters(num_args int, type_spec string, args ...any) int {
	return argparse.ParseVaArgs(num_args, type_spec, args, 0)
}
func ZendParseParametersThrow(num_args int, type_spec string, args ...any) int {
	return argparse.ParseVaArgs(num_args, type_spec, args, argparse.ZEND_PARSE_PARAMS_THROW)
}
func ZendParseMethodParameters(num_args int, this_ptr *types.Zval, type_spec string, args ...any) int {
	/* Just checking this_ptr is not enough, because fcall_common_helper does not set
	 * Z_OBJ(EG(This)) to NULL when calling an internal function with common.scope == NULL.
	 * In that case EG(This) would still be the $this from the calling code and we'd take the
	 * wrong branch here. */

	var is_method = CurrEX().GetFunc().GetScope() != nil
	if !is_method || this_ptr == nil || this_ptr.GetType() != types.IS_OBJECT {
		return argparse.ParseVaArgs(num_args, type_spec, args, 0)
	} else {
		object := args[0].(**types.Zval)
		ce := args[1].(*types.ClassEntry)
		*object = this_ptr
		if ce != nil && InstanceofFunction(types.Z_OBJCE_P(this_ptr), ce) == 0 {
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "%s::%s() must be derived from %s::%s", types.Z_OBJCE_P(this_ptr).GetName().GetVal(), GetActiveFunctionName(), ce.GetName().GetVal(), GetActiveFunctionName())
		}
		return argparse.ParseVaArgs(num_args, type_spec[1:], args[2:], 0)
	}
}
func ZendParseMethodParametersEx(flags int, num_args int, this_ptr *types.Zval, type_spec string, args ...any) int {
	if this_ptr == nil {
		return argparse.ParseVaArgs(num_args, type_spec, args, flags)
	} else {
		object := args[0].(**types.Zval)
		ce := args[1].(*types.ClassEntry)
		*object = this_ptr
		if ce != nil && InstanceofFunction(types.Z_OBJCE_P(this_ptr), ce) == 0 {
			if (flags & argparse.ZEND_PARSE_PARAMS_QUIET) == 0 {
				faults.ErrorNoreturn(faults.E_CORE_ERROR, "%s::%s() must be derived from %s::%s", ce.GetName().GetVal(), GetActiveFunctionName(), types.Z_OBJCE_P(this_ptr).GetName().GetVal(), GetActiveFunctionName())
			}
			return types.FAILURE
		}
		return argparse.ParseVaArgs(num_args, type_spec[1:], args[2:], flags)
	}
}
func ZendMergeProperties(obj *types.Zval, properties *types.HashTable) {
	var obj_ht *ZendObjectHandlers = types.Z_OBJ_HT_P(obj)
	var old_scope *types.ClassEntry = EG__().GetFakeScope()
	var key *types.ZendString
	var value *types.Zval
	EG__().SetFakeScope(types.Z_OBJCE_P(obj))

	var __ht *types.HashTable = properties
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		key = _p.GetKey()
		value = _z
		if key != nil {
			var member types.Zval
			member.SetString(key)
			obj_ht.GetWriteProperty()(obj, &member, value, nil)
		}
	}
	EG__().SetFakeScope(old_scope)
}
func ZendUpdateClassConstants(class_type *types.ClassEntry) int {
	if !class_type.IsConstantsUpdated() {
		var ce *types.ClassEntry
		var c *ZendClassConstant
		var val *types.Zval
		var prop_info *ZendPropertyInfo
		if class_type.GetParent() {
			if ZendUpdateClassConstants(class_type.GetParent()) != types.SUCCESS {
				return types.FAILURE
			}
		}
		var __ht *types.HashTable = class_type.GetConstantsTable()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			c = _z.GetPtr()
			val = c.GetValue()
			if val.IsConstant() {
				if ZvalUpdateConstantEx(val, c.GetCe()) != types.SUCCESS {
					return types.FAILURE
				}
			}
		}
		if class_type.GetDefaultStaticMembersCount() != 0 && CE_STATIC_MEMBERS(class_type) == nil {
			if class_type.GetType() == ZEND_INTERNAL_CLASS || class_type.HasCeFlags(ZEND_ACC_IMMUTABLE|ZEND_ACC_PRELOADED) {
				ZendClassInitStatics(class_type)
			}
		}
		ce = class_type
		for ce != nil {
			var __ht *types.HashTable = ce.GetPropertiesInfo()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()

				prop_info = _z.GetPtr()
				if prop_info.GetCe() == ce {
					if prop_info.IsStatic() {
						val = CE_STATIC_MEMBERS(class_type) + prop_info.GetOffset()
					} else {
						val = (*types.Zval)((*byte)(class_type.GetDefaultPropertiesTable() + prop_info.GetOffset() - OBJ_PROP_TO_OFFSET(0)))
					}
					if val.IsConstant() {
						if prop_info.GetType() != 0 {
							var tmp types.Zval
							types.ZVAL_COPY(&tmp, val)
							if ZvalUpdateConstantEx(&tmp, ce) != types.SUCCESS {
								ZvalPtrDtor(&tmp)
								return types.FAILURE
							}
							if ZendVerifyPropertyType(prop_info, &tmp, 1) == 0 {
								ZvalPtrDtor(&tmp)
								return types.FAILURE
							}
							ZvalPtrDtor(val)
							types.ZVAL_COPY_VALUE(val, &tmp)
						} else if ZvalUpdateConstantEx(val, ce) != types.SUCCESS {
							return types.FAILURE
						}
					}
				}
			}
			ce = ce.GetParent()
		}
		class_type.SetIsConstantsUpdated(true)
	}
	return types.SUCCESS
}
func _objectPropertiesInit(object *types.ZendObject, class_type *types.ClassEntry) {
	if class_type.GetDefaultPropertiesCount() != 0 {
		var src *types.Zval = class_type.GetDefaultPropertiesTable()
		var dst *types.Zval = object.GetPropertiesTable()
		var end *types.Zval = src + class_type.GetDefaultPropertiesCount()
		if class_type.GetType() == ZEND_INTERNAL_CLASS {
			for {
				types.ZVAL_COPY_OR_DUP_PROP(dst, src)
				src++
				dst++
				if src == end {
					break
				}
			}
		} else {
			for {
				types.ZVAL_COPY_PROP(dst, src)
				src++
				dst++
				if src == end {
					break
				}
			}
		}
	}
}
