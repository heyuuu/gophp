package argparse

import (
	"fmt"
	b "sik/builtin"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
)

func CheckNumArgsNoneError() bool { return CheckNumArgsNone(zend.CurrEX(), 0) }
func CheckNumArgsNoneException() bool {
	return CheckNumArgsNone(zend.CurrEX(), ZEND_PARSE_PARAMS_THROW)
}

func ZendWrongCallbackDeprecated(num int, error string) {
	message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", zend.GetActiveCalleeName(), num, error)
	faults.ErrorEx(faults.E_DEPRECATED, message)
}
func ZendParseArgBoolWeak(arg *types.Zval, dest *types.ZendBool) int {
	if val, ok := ParseBoolWeak(arg); ok {
		*dest = types.IntBool(val)
		return 1
	}
	return 0
}
func ZendParseArgLongWeak(arg *types.Zval, dest *zend.ZendLong) int {
	if val, ok := ParseLongWeak(arg, false); ok {
		*dest = val
		return 1
	}
	return 0
}
func ZendParseArgDoubleWeak(arg *types.Zval, dest *float64) int {
	if val, ok := ParseDoubleWeak(arg); ok {
		*dest = val
		return 1
	}
	return 0
}

func ZendParseArgStrWeak(arg *types.Zval, dest **types.ZendString) int {
	if val, ok := ParseZStrWeak(arg); ok {
		*dest = val
		return 1
	}
	return 0
}

func ZendParseParametersDebugError(msg string) {
	var active_function *zend.ZendFunction = zend.CurrEX().GetFunc()
	var class_name *byte = b.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	faults.ErrorNoreturn(faults.E_CORE_ERROR, "%s%s%s(): %s", class_name, b.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), msg)
}
