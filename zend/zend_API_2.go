package zend

import (
	"fmt"
	"sik/zend/argparse"
	"sik/zend/types"
)

func CheckNumArgsNoneError() bool     { return CurrEX().CheckNumArgsError(0, 0) }
func CheckNumArgsNoneException() bool { return CurrEX().CheckNumArgsException(0, 0) }
func CheckNumArgsError(minNumArgs int, maxNumArgs int) bool {
	return CurrEX().CheckNumArgsError(minNumArgs, maxNumArgs)
}
func CheckNumArgsException(minNumArgs int, maxNumArgs int) bool {
	return CurrEX().CheckNumArgsException(minNumArgs, maxNumArgs)
}

func WrongParamTypeError(num int, expectedType argparse.ZendExpectedType, arg *types.Zval, forceStrict bool) {
	if EG__().GetException() != nil {
		return
	}
	message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", GetActiveCalleeName(), num, expectedType, ZendZvalTypeName(arg))
	throwException := forceStrict || CurrEX().IsArgUseStrictTypes()
	ZendInternalTypeErrorEx(throwException, message)
}

func ZendWrongParameterTypeError(num int, expected_type argparse.ZendExpectedType, arg *types.Zval) {
	WrongParamTypeError(num, expected_type, arg, false)
}
func ZendWrongParameterTypeException(num int, expected_type argparse.ZendExpectedType, arg *types.Zval) {
	WrongParamTypeError(num, expected_type, arg, true)
}

func WrongParamClassError(num int, name string, arg *types.Zval, forceStrict bool) {
	if EG__().GetException() != nil {
		return
	}
	message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", GetActiveCalleeName(), num, name, ZendZvalTypeName(arg))
	throwException := forceStrict || CurrEX().IsArgUseStrictTypes()
	ZendInternalTypeErrorEx(throwException, message)
}

func ZendWrongParameterClassError(num int, name string, arg *types.Zval) {
	WrongParamClassError(num, name, arg, false)
}
func ZendWrongParameterClassException(num int, name string, arg *types.Zval) {
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
func ZendParseArgClass(arg *types.Zval, pce **ZendClassEntry, num int, check_null int) int {
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
func ZendParseArgBoolWeak(arg *types.Zval, dest *types.ZendBool) int {
	if val, ok := argparse.ParseBoolWeak(arg); ok {
		*dest = types.IntBool(val)
		return 1
	}
	return 0
}
func ZendParseArgLongWeak(arg *types.Zval, dest *ZendLong) int {
	if val, ok := argparse.ParseLongWeak(arg, false); ok {
		*dest = val
		return 1
	}
	return 0
}
func ZendParseArgDoubleWeak(arg *types.Zval, dest *float64) int {
	if val, ok := argparse.ParseDoubleWeak(arg); ok {
		*dest = val
		return 1
	}
	return 0
}

func ZendParseArgStrWeak(arg *types.Zval, dest **types.ZendString) int {
	if val, ok := argparse.ParseZStrWeak(arg); ok {
		*dest = val
		return 1
	}
	return 0
}
