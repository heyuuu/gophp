package zend

import (
	"fmt"
	"sik/zend/argparse"
	"sik/zend/faults"
	"sik/zend/types"
)

func CheckNumArgsNoneError() bool { return argparse.CheckNumArgsNone(CurrEX(), 0) }
func CheckNumArgsNoneException() bool {
	return argparse.CheckNumArgsNone(CurrEX(), argparse.ZEND_PARSE_PARAMS_THROW)
}

func WrongParamTypeError(num int, expectedType string, arg *types.Zval, forceStrict bool) {
	if EG__().GetException() != nil {
		return
	}
	message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", GetActiveCalleeName(), num, expectedType, ZendZvalTypeName(arg))
	throwException := forceStrict || CurrEX().IsArgUseStrictTypes()
	faults.InternalTypeErrorEx(throwException, message)
}

func WrongParamClassError(num int, name string, arg *types.Zval, forceStrict bool) {
	if EG__().GetException() != nil {
		return
	}
	message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", GetActiveCalleeName(), num, name, ZendZvalTypeName(arg))
	throwException := forceStrict || CurrEX().IsArgUseStrictTypes()
	faults.InternalTypeErrorEx(throwException, message)
}

func WrongCallbackError(num int, error string, forceStrict bool) {
	if EG__().GetException() != nil {
		return
	}
	message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", GetActiveCalleeName(), num, error)
	throwException := forceStrict || CurrEX().IsArgUseStrictTypes()
	faults.InternalTypeErrorEx(throwException, message)
}

func ZendWrongCallbackDeprecated(num int, error string) {
	message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", GetActiveCalleeName(), num, error)
	faults.ErrorEx(faults.E_DEPRECATED, message)
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
