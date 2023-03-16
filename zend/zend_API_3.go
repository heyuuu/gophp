// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendParseArg(arg_num int, arg *Zval, va *va_list, spec **byte, flags int) int {
	var expected_type *byte = nil
	var error *byte = nil
	var severity int = 0
	expected_type = ZendParseArgImpl(arg_num, arg, va, spec, &error, &severity)
	if expected_type != nil {
		if EG__().GetException() != nil {
			return FAILURE
		}
		if (flags&ZEND_PARSE_PARAMS_QUIET) == 0 && ((*expected_type) || error != nil) {
			var space *byte
			class_name, space := GetActiveClassName()
			var throw_exception ZendBool = ZEND_ARG_USES_STRICT_TYPES() || (flags&ZEND_PARSE_PARAMS_THROW) != 0
			if error != nil {
				ZendInternalTypeError(throw_exception, "%s%s%s() expects parameter %d %s", class_name, space, GetActiveFunctionName(), arg_num, error)
				Efree(error)
			} else {
				ZendInternalTypeError(throw_exception, "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), arg_num, expected_type, ZendZvalTypeName(arg))
			}
		}
		if severity != E_DEPRECATED {
			return FAILURE
		}
	}
	return SUCCESS
}
func ZendParseParametersDebugError(msg string) {
	var active_function *ZendFunction = CurrEX().GetFunc()
	var class_name *byte = b.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendErrorNoreturn(E_CORE_ERROR, "%s%s%s(): %s", class_name, b.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), msg)
}
func ZendParseVaArgs(num_args int, type_spec *byte, va *va_list, flags int) int {
	var spec_walk *byte
	var c int
	var i int
	var min_num_args int = -1
	var max_num_args int = 0
	var post_varargs int = 0
	var arg *Zval
	var arg_count int
	var have_varargs ZendBool = 0
	var varargs **Zval = nil
	var n_varargs *int = nil
	for spec_walk = type_spec; *spec_walk; spec_walk++ {
		c = *spec_walk
		switch c {
		case 'l':

		case 'd':

		case 's':

		case 'b':

		case 'r':

		case 'a':

		case 'o':

		case 'O':

		case 'z':

		case 'Z':

		case 'C':

		case 'h':

		case 'f':

		case 'A':

		case 'H':

		case 'p':

		case 'S':

		case 'P':

		case 'L':
			max_num_args++
			break
		case '|':
			min_num_args = max_num_args
			break
		case '/':

		case '!':

			/* Pass */

			break
		case '*':

		case '+':
			if have_varargs != 0 {
				ZendParseParametersDebugError("only one varargs specifier (* or +) is permitted")
				return FAILURE
			}
			have_varargs = 1

			/* we expect at least one parameter in varargs */

			if c == '+' {
				max_num_args++
			}

			/* mark the beginning of varargs */

			post_varargs = max_num_args
			break
		default:
			ZendParseParametersDebugError("bad type specifier while parsing parameters")
			return FAILURE
		}
	}
	if min_num_args < 0 {
		min_num_args = max_num_args
	}
	if have_varargs != 0 {

		/* calculate how many required args are at the end of the specifier list */

		post_varargs = max_num_args - post_varargs
		max_num_args = -1
	}
	if num_args < min_num_args || num_args > max_num_args && max_num_args >= 0 {
		if (flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
			var active_function *ZendFunction = CurrEX().GetFunc()
			var class_name *byte = b.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
			var throw_exception = ZEND_ARG_USES_STRICT_TYPES() || (flags&ZEND_PARSE_PARAMS_THROW) != 0
			ZendInternalArgumentCountError(throw_exception, "%s%s%s() expects %s %d parameter%s, %d given", class_name, b.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), b.Cond(b.Cond(min_num_args == max_num_args, "exactly", num_args < min_num_args), "at least", "at most"), b.Cond(num_args < min_num_args, min_num_args, max_num_args), b.Cond(b.Cond(num_args < min_num_args, min_num_args, max_num_args) == 1, "", "s"), num_args)
		}
		return FAILURE
	}
	arg_count = ZEND_CALL_NUM_ARGS(CurrEX())
	if num_args > arg_count {
		ZendParseParametersDebugError("could not obtain parameters for parsing")
		return FAILURE
	}
	i = 0
	for b.PostDec(&num_args) > 0 {
		if (*type_spec) == '|' {
			type_spec++
		}
		if (*type_spec) == '*' || (*type_spec) == '+' {
			var num_varargs int = num_args + 1 - post_varargs

			/* eat up the passed in storage even if it won't be filled in with varargs */

			varargs = __va_arg(*va, (**Zval)(_))
			n_varargs = __va_arg(*va, (*int)(_))
			type_spec++
			if num_varargs > 0 {
				*n_varargs = num_varargs
				*varargs = ZEND_CALL_ARG(CurrEX(), i+1)

				/* adjust how many args we have left and restart loop */

				num_args += 1 - num_varargs
				i += num_varargs
				continue
			} else {
				*varargs = nil
				*n_varargs = 0
			}
		}
		arg = ZEND_CALL_ARG(CurrEX(), i+1)
		if ZendParseArg(i+1, arg, va, &type_spec, flags) == FAILURE {

			/* clean up varargs array if it was used */

			if varargs != nil && (*varargs) != nil {
				*varargs = nil
			}
			return FAILURE
		}
		i++
	}
	return SUCCESS
}
func ZendParseParametersEx(flags int, num_args int, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	va_start(va, type_spec)
	retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
	va_end(va)
	return retval
}
func ZendParseParameters(num_args int, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	var flags int = 0
	va_start(va, type_spec)
	retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
	va_end(va)
	return retval
}
func ZendParseParametersThrow(num_args int, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	var flags int = ZEND_PARSE_PARAMS_THROW
	va_start(va, type_spec)
	retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
	va_end(va)
	return retval
}
func ZendParseMethodParameters(num_args int, this_ptr *Zval, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	var flags int = 0
	var p *byte = type_spec
	var object **Zval
	var ce *ZendClassEntry

	/* Just checking this_ptr is not enough, because fcall_common_helper does not set
	 * Z_OBJ(EG(This)) to NULL when calling an internal function with common.scope == NULL.
	 * In that case EG(This) would still be the $this from the calling code and we'd take the
	 * wrong branch here. */

	var is_method ZendBool = CurrEX().GetFunc().GetScope() != nil
	if is_method == 0 || this_ptr == nil || this_ptr.GetType() != IS_OBJECT {
		va_start(va, type_spec)
		retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
		va_end(va)
	} else {
		p++
		va_start(va, type_spec)
		object = __va_arg(va, (**Zval)(_))
		ce = __va_arg(va, (*ZendClassEntry)(_))
		*object = this_ptr
		if ce != nil && InstanceofFunction(Z_OBJCE_P(this_ptr), ce) == 0 {
			ZendErrorNoreturn(E_CORE_ERROR, "%s::%s() must be derived from %s::%s", Z_OBJCE_P(this_ptr).GetName().GetVal(), GetActiveFunctionName(), ce.GetName().GetVal(), GetActiveFunctionName())
		}
		retval = ZendParseVaArgs(num_args, p, &va, flags)
		va_end(va)
	}
	return retval
}
func ZendParseMethodParametersEx(flags int, num_args int, this_ptr *Zval, type_spec *byte, _ ...any) int {
	var va va_list
	var retval int
	var p *byte = type_spec
	var object **Zval
	var ce *ZendClassEntry
	if this_ptr == nil {
		va_start(va, type_spec)
		retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
		va_end(va)
	} else {
		p++
		va_start(va, type_spec)
		object = __va_arg(va, (**Zval)(_))
		ce = __va_arg(va, (*ZendClassEntry)(_))
		*object = this_ptr
		if ce != nil && InstanceofFunction(Z_OBJCE_P(this_ptr), ce) == 0 {
			if (flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				ZendErrorNoreturn(E_CORE_ERROR, "%s::%s() must be derived from %s::%s", ce.GetName().GetVal(), GetActiveFunctionName(), Z_OBJCE_P(this_ptr).GetName().GetVal(), GetActiveFunctionName())
			}
			va_end(va)
			return FAILURE
		}
		retval = ZendParseVaArgs(num_args, p, &va, flags)
		va_end(va)
	}
	return retval
}
func ZendMergeProperties(obj *Zval, properties *HashTable) {
	var obj_ht *ZendObjectHandlers = Z_OBJ_HT_P(obj)
	var old_scope *ZendClassEntry = EG__().GetFakeScope()
	var key *ZendString
	var value *Zval
	EG__().SetFakeScope(Z_OBJCE_P(obj))
	var __ht *HashTable = properties
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		key = _p.GetKey()
		value = _z
		if key != nil {
			var member Zval
			member.SetString(key)
			obj_ht.GetWriteProperty()(obj, &member, value, nil)
		}
	}
	EG__().SetFakeScope(old_scope)
}
func ZendUpdateClassConstants(class_type *ZendClassEntry) int {
	if !class_type.IsConstantsUpdated() {
		var ce *ZendClassEntry
		var c *ZendClassConstant
		var val *Zval
		var prop_info *ZendPropertyInfo
		if class_type.GetParent() {
			if ZendUpdateClassConstants(class_type.GetParent()) != SUCCESS {
				return FAILURE
			}
		}
		var __ht *HashTable = class_type.GetConstantsTable()
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			c = _z.GetPtr()
			val = c.GetValue()
			if val.IsConstant() {
				if ZvalUpdateConstantEx(val, c.GetCe()) != SUCCESS {
					return FAILURE
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
			var __ht *HashTable = ce.GetPropertiesInfo()
			for _, _p := range __ht.foreachData() {
				var _z *Zval = _p.GetVal()

				prop_info = _z.GetPtr()
				if prop_info.GetCe() == ce {
					if prop_info.IsStatic() {
						val = CE_STATIC_MEMBERS(class_type) + prop_info.GetOffset()
					} else {
						val = (*Zval)((*byte)(class_type.GetDefaultPropertiesTable() + prop_info.GetOffset() - OBJ_PROP_TO_OFFSET(0)))
					}
					if val.IsConstant() {
						if prop_info.GetType() != 0 {
							var tmp Zval
							ZVAL_COPY(&tmp, val)
							if ZvalUpdateConstantEx(&tmp, ce) != SUCCESS {
								ZvalPtrDtor(&tmp)
								return FAILURE
							}
							if ZendVerifyPropertyType(prop_info, &tmp, 1) == 0 {
								ZvalPtrDtor(&tmp)
								return FAILURE
							}
							ZvalPtrDtor(val)
							ZVAL_COPY_VALUE(val, &tmp)
						} else if ZvalUpdateConstantEx(val, ce) != SUCCESS {
							return FAILURE
						}
					}
				}
			}
			ce = ce.GetParent()
		}
		class_type.SetIsConstantsUpdated(true)
	}
	return SUCCESS
}
func _objectPropertiesInit(object *ZendObject, class_type *ZendClassEntry) {
	if class_type.GetDefaultPropertiesCount() != 0 {
		var src *Zval = class_type.GetDefaultPropertiesTable()
		var dst *Zval = object.GetPropertiesTable()
		var end *Zval = src + class_type.GetDefaultPropertiesCount()
		if class_type.GetType() == ZEND_INTERNAL_CLASS {
			for {
				ZVAL_COPY_OR_DUP_PROP(dst, src)
				src++
				dst++
				if src == end {
					break
				}
			}
		} else {
			for {
				ZVAL_COPY_PROP(dst, src)
				src++
				dst++
				if src == end {
					break
				}
			}
		}
	}
}
