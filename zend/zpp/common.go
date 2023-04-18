package zpp

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func CheckNumArgsNoneError(executeData ExecuteData) bool {
	return CheckNumArgs(executeData, 0, 0, 0)
}
func CheckNumArgsNoneException(executeData ExecuteData) bool {
	return CheckNumArgs(executeData, 0, 0, FlagThrow)
}
func CheckNumArgs(executeData ExecuteData, minNumArgs int, maxNumArgs int, flags int) bool {
	numArgs := executeData.NumArgs()
	return CheckNumArgsEx(numArgs, executeData, minNumArgs, maxNumArgs, flags)
}

func CheckNumArgsEx(numArgs int, executeData ExecuteData, minNumArgs int, maxNumArgs int, flags int) bool {
	// 检查参数个数，若检查通过直接返回
	if numArgs >= minNumArgs && (numArgs <= maxNumArgs || maxNumArgs < 0) {
		return true
	}

	// 非 Quiet 模式下，触发 PHP Error
	if (flags & FlagQuiet) == 0 {
		// 判断是否强制抛出异常或为 strict 模式
		var throwException = (flags&FlagThrow) != 0 || executeData.IsArgUseStrictTypes()

		// 构建错误信息
		callee := executeData.CalleeName()
		if minNumArgs == maxNumArgs {
			faults.InternalArgumentCountError(throwException, "%s() expects exactly %d parameter%s, %d given", callee, minNumArgs, b.Cond(minNumArgs == 1, "", "s"), numArgs)
		} else if numArgs < minNumArgs {
			faults.InternalArgumentCountError(throwException, "%s() expects at least %d parameter%s, %d given", callee, minNumArgs, b.Cond(minNumArgs == 1, "", "s"), numArgs)
		} else { // numArgs > maxNumArgs
			faults.InternalArgumentCountError(throwException, "%s() expects at most %d parameter%s, %d given", callee, maxNumArgs, b.Cond(maxNumArgs == 1, "", "s"), numArgs)
		}
	}

	return false
}

func ZendParseArgStrWeak(arg *types2.Zval, dest **types2.String) int {
	if val, ok := ParseZStrWeak(arg); ok {
		*dest = val
		return 1
	}
	return 0
}

func ZendParseParametersDebugError(msg string) {
	faults.ErrorNoreturn(faults.E_CORE_ERROR, "%s(): %s", currExecuteData().CalleeName(), msg)
}
