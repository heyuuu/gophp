package zend

import (
	"fmt"
	b "sik/builtin"
)

func CheckNumArgsNoneError() bool     { return CurrEX().CheckNumArgsError(0, 0) }
func CheckNumArgsNoneException() bool { return CurrEX().CheckNumArgsException(0, 0) }
func CheckNumArgsError(minNumArgs int, maxNumArgs int) bool {
	return CurrEX().CheckNumArgsError(minNumArgs, maxNumArgs)
}
func CheckNumArgsException(minNumArgs int, maxNumArgs int) bool {
	return CurrEX().CheckNumArgsException(minNumArgs, maxNumArgs)
}

func WrongParamTypeError(num int, expectedType ZendExpectedType, arg *Zval, forceStrict bool) {
	if EG__().GetException() != nil {
		return
	}
	message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", GetActiveCalleeName(), num, expectedType, ZendZvalTypeName(arg))
	throwException := forceStrict || CurrEX().IsArgUseStrictTypes()
	ZendInternalTypeErrorEx(throwException, message)
}

func ZendWrongParameterTypeError(num int, expected_type ZendExpectedType, arg *Zval) {
	WrongParamTypeError(num, expected_type, arg, false)
}
func ZendWrongParameterTypeException(num int, expected_type ZendExpectedType, arg *Zval) {
	WrongParamTypeError(num, expected_type, arg, true)
}

func WrongParamClassError(num int, name string, arg *Zval, forceStrict bool) {
	if EG__().GetException() != nil {
		return
	}
	message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", GetActiveCalleeName(), num, name, ZendZvalTypeName(arg))
	throwException := forceStrict || CurrEX().IsArgUseStrictTypes()
	ZendInternalTypeErrorEx(throwException, message)
}

func ZendWrongParameterClassError(num int, name string, arg *Zval) {
	WrongParamClassError(num, name, arg, false)
}
func ZendWrongParameterClassException(num int, name string, arg *Zval) {
	WrongParamClassError(num, name, arg, true)
}

func WrongCallbackError(num int, error string, forceStrict bool) {
	if EG__().GetException() != nil {
		return
	}
	message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", GetActiveCalleeName(), num, error)
	throwException := forceStrict || CurrEX().IsArgUseStrictTypes()
	ZendInternalTypeErrorEx(throwException, message)
}

func ZendWrongCallbackError(num int, error string) {
	WrongCallbackError(num, error, false)
}
func ZendWrongCallbackException(num int, error string) {
	WrongCallbackError(num, error, true)
}
func ZendWrongCallbackDeprecated(num int, error string) {
	message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", GetActiveCalleeName(), num, error)
	ZendErrorEx(E_DEPRECATED, message)
}
func ZendParseArgClass(arg *Zval, pce **ZendClassEntry, num int, check_null int) int {
	var ce_base *ZendClassEntry = *pce
	if check_null != 0 && arg.IsNull() {
		*pce = nil
		return 1
	}
	if TryConvertToString(arg) == 0 {
		*pce = nil
		return 0
	}
	*pce = ZendLookupClass(arg.GetStr())
	if ce_base != nil {
		if (*pce) == nil || InstanceofFunction(*pce, ce_base) == 0 {
			ZendInternalTypeError(CurrEX().IsArgUseStrictTypes(), "%s() expects parameter %d to be a class name derived from %s, '%s' given", GetActiveCalleeName(), num, ce_base.GetName().GetVal(), arg.GetStr().GetVal())
			*pce = nil
			return 0
		}
	}
	if (*pce) == nil {
		ZendInternalTypeError(CurrEX().IsArgUseStrictTypes(), "%s() expects parameter %d to be a valid class name, '%s' given", GetActiveCalleeName(), num, arg.GetStr().GetVal())
		return 0
	}
	return 1
}
func ZendParseArgBoolWeak(arg *Zval, dest *ZendBool) int {
	if val, ok := ParseArgBoolWeak(arg); ok {
		*dest = intBool(val)
		return 1
	}
	return 0
}
func ZendParseArgLongWeak(arg *Zval, dest *ZendLong) int {
	if val, ok := ParseArgLongWeak(arg, false); ok {
		*dest = val
		return 1
	}
	return 0
}
func ZendParseArgDoubleWeak(arg *Zval, dest *float64) int {
	if val, ok := ParseArgDoubleWeak(arg); ok {
		*dest = val
		return 1
	}
	return 0
}

func ZendParseArgStrWeak(arg *Zval, dest **ZendString) int {
	if val, ok := ParseArgStrWeak(arg); ok {
		*dest = val
		return 1
	}
	return 0
}
func ZendParseArgStrSlow(arg *Zval, dest **ZendString) int {
	if CurrEX().IsArgUseStrictTypes() {
		return 0
	}
	return ZendParseArgStrWeak(arg, dest)
}

func ZendParseArgImpl(
	arg_num int,
	arg *Zval,
	va *va_list,
	spec **byte,
	error **byte,
	severity *int,
) *byte {
	return ZendParseArgImpl_Ex(arg, va, spec, error, severity)
}

func ZendParseArgImpl_Ex(arg *Zval, va *b.VaList, spec *b.StrReader, error **byte, severity *int) string {
	specWalk := spec.Copy()
	c := specWalk.Read()
	var check_null int = 0
	var separate int = 0
	var real_arg *Zval = arg

	/* scan through modifiers */

	arg = ZVAL_DEREF(arg)
	for true {
		if specWalk.Curr() == '/' {
			SEPARATE_ZVAL_NOREF(arg)
			real_arg = arg
			separate = 1
		} else if specWalk.Curr() == '!' {
			check_null = 1
		} else {
			break
		}
		specWalk.Next()
	}
	switch c {
	case 'l', 'L':
		var p = va.Pop().(*ZendLong)
		var is_null *ZendBool = nil
		if check_null != 0 {
			is_null = va.Pop().(*ZendBool)
		}
		if ZendParseArgLong(arg, p, is_null, check_null, c == 'L') == 0 {
			return "int"
		}
		break
	case 'd':
		var p *float64 = __va_arg(*va, (*float64)(_))
		var is_null *ZendBool = nil
		if check_null != 0 {
			is_null = __va_arg(*va, (*ZendBool)(_))
		}
		if ZendParseArgDouble(arg, p, is_null, check_null) == 0 {
			return "float"
		}
		break
	case 's':
		var p **byte = __va_arg(*va, (**byte)(_))
		var pl *int = __va_arg(*va, (*int)(_))
		if ZendParseArgString(arg, p, pl, check_null) == 0 {
			return "string"
		}
		break
	case 'p':
		var p **byte = __va_arg(*va, (**byte)(_))
		var pl *int = __va_arg(*va, (*int)(_))
		if ZendParseArgPath(arg, p, pl, check_null) == 0 {
			return "a valid path"
		}
		break
	case 'P':
		var str **ZendString = __va_arg(*va, (**ZendString)(_))
		if ZendParseArgPathStr(arg, str, check_null) == 0 {
			return "a valid path"
		}
		break
	case 'S':
		var str **ZendString = __va_arg(*va, (**ZendString)(_))
		if ZendParseArgStr(arg, str, check_null) == 0 {
			return "string"
		}
		break
	case 'b':
		var p *ZendBool = __va_arg(*va, (*ZendBool)(_))
		var is_null *ZendBool = nil
		if check_null != 0 {
			is_null = __va_arg(*va, (*ZendBool)(_))
		}
		if ZendParseArgBool(arg, p, is_null, check_null) == 0 {
			return "bool"
		}
		break
	case 'r':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		if ZendParseArgResource(arg, p, check_null) == 0 {
			return "resource"
		}
		break
	case 'A':

	case 'a':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		if ZendParseArgArray(arg, p, check_null, c == 'A') == 0 {
			return "array"
		}
		break
	case 'H':

	case 'h':
		var p **HashTable = __va_arg(*va, (**HashTable)(_))
		if ZendParseArgArrayHt(arg, p, check_null, c == 'H', separate) == 0 {
			return "array"
		}
		break
	case 'o':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		if ZendParseArgObject(arg, p, nil, check_null) == 0 {
			return "object"
		}
		break
	case 'O':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		var ce *ZendClassEntry = __va_arg(*va, (*ZendClassEntry)(_))
		if ZendParseArgObject(arg, p, ce, check_null) == 0 {
			if ce != nil {
				return ce.GetName().GetVal()
			} else {
				return "object"
			}
		}
		break
	case 'C':
		var lookup *ZendClassEntry
		var pce **ZendClassEntry = __va_arg(*va, (**ZendClassEntry)(_))
		var ce_base *ZendClassEntry = *pce
		if check_null != 0 && arg.IsNull() {
			*pce = nil
			break
		}
		if TryConvertToString(arg) == 0 {
			*pce = nil
			return "valid class name"
		}
		if b.Assign(&lookup, ZendLookupClass(arg.GetStr())) == nil {
			*pce = nil
		} else {
			*pce = lookup
		}
		if ce_base != nil {
			if (*pce) == nil || InstanceofFunction(*pce, ce_base) == 0 {
				ZendSpprintf(error, 0, "to be a class name derived from %s, '%s' given", ce_base.GetName().GetVal(), arg.GetStr().GetVal())
				*pce = nil
				return ""
			}
		}
		if (*pce) == nil {
			ZendSpprintf(error, 0, "to be a valid class name, '%s' given", arg.GetStr().GetVal())
			return ""
		}
		break
		break
	case 'f':
		var fci *ZendFcallInfo = __va_arg(*va, (*ZendFcallInfo)(_))
		var fcc *ZendFcallInfoCache = __va_arg(*va, (*ZendFcallInfoCache)(_))
		var is_callable_error *byte = nil
		if check_null != 0 && arg.IsNull() {
			fci.SetSize(0)
			fcc.SetFunctionHandler(0)
			break
		}
		if ZendFcallInfoInit(arg, 0, fci, fcc, nil, &is_callable_error) == SUCCESS {
			if is_callable_error != nil {
				*severity = E_DEPRECATED
				ZendSpprintf(error, 0, "to be a valid callback, %s", is_callable_error)
				Efree(is_callable_error)
				*spec = specWalk
				return ""
			}
			break
		} else {
			if is_callable_error != nil {
				*severity = E_ERROR
				ZendSpprintf(error, 0, "to be a valid callback, %s", is_callable_error)
				Efree(is_callable_error)
				return ""
			} else {
				return "valid callback"
			}
		}
	case 'z':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		ZendParseArgZvalDeref(real_arg, p, check_null)
		break
	case 'Z':

		/* 'Z' iz not supported anymore and should be replaced with 'z' */

		ZEND_ASSERT(c != 'Z')
	default:
		return "unknown"
	}
	*spec = specWalk
	return nil
}
