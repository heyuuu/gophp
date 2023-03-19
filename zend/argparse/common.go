package argparse

import (
	b "sik/builtin"
	"sik/zend"
)

func CheckNumArgsNone(executeData ExecuteData, flags int) bool {
	return CheckNumArgs(executeData, 0, 0, flags)
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
	if (flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
		// 判断是否强制抛出异常或为 strict 模式
		var throwException = (flags&ZEND_PARSE_PARAMS_THROW) != 0 || executeData.IsArgUseStrictTypes()

		// 构建错误信息
		callee := executeData.CalleeName()
		if minNumArgs == maxNumArgs {
			zend.ZendInternalArgumentCountError(throwException, "%s() expects exactly %d parameter%s, %d given", callee, minNumArgs, b.Cond(minNumArgs == 1, "", "s"), numArgs)
		} else if numArgs < minNumArgs {
			zend.ZendInternalArgumentCountError(throwException, "%s() expects at least %d parameter%s, %d given", callee, minNumArgs, b.Cond(minNumArgs == 1, "", "s"), numArgs)
		} else { // numArgs > maxNumArgs
			zend.ZendInternalArgumentCountError(throwException, "%s() expects at most %d parameter%s, %d given", callee, maxNumArgs, b.Cond(maxNumArgs == 1, "", "s"), numArgs)
		}
	}

	return false
}
