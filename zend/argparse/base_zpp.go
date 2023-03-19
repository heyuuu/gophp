package argparse

import (
	"fmt"
	"sik/zend/faults"
	"sik/zend/types"
)

type baseParser struct {
	executeData ExecuteData
	numArgs     int // 所需参数个数 (注意: numArgs 与 executeData.NumArgs() 不一定相等)
	minNumArgs  int
	maxNumArgs  int
	flags       int
	errorCode   int
	finish      bool // 解析已终止 (可能是已解析完成或出现错误)
	idx         int  // 已读取的 Arg 位置，需注意它的值是从 1 开始的，范围是 [1, numArgs]
	arg         *types.Zval
}

func makeBaseParser(executeData ExecuteData, numArgs int, minNumArgs int, maxNumArgs int, flags int) baseParser {
	return baseParser{executeData: executeData, numArgs: numArgs, minNumArgs: minNumArgs, maxNumArgs: maxNumArgs, flags: flags}
}

func (p *baseParser) start() {
	// check num args
	if !CheckNumArgsEx(p.numArgs, p.executeData, p.minNumArgs, p.maxNumArgs, p.flags) {
		p.triggerError(ZPP_ERROR_FAILURE, "")
	}
}

func (p *baseParser) currArg() *types.Zval {
	return p.executeData.Arg(p.idx)
}

func (p *baseParser) isQuiet() bool { return p.flags&ZEND_PARSE_PARAMS_QUIET != 0 }
func (p *baseParser) isThrow() bool { return p.flags&ZEND_PARSE_PARAMS_THROW != 0 }
func (p *baseParser) isThrowEx() bool {
	return p.flags&ZEND_PARSE_PARAMS_THROW != 0 || p.executeData.IsArgUseStrictTypes()
}

func (p *baseParser) IsFinish() bool {
	return p.finish
}

func (p *baseParser) HasError() bool {
	return p.errorCode != ZPP_ERROR_OK
}

func (p *baseParser) triggerError(errorCode int, err string) {
	// 记录错误信息
	p.errorCode = errorCode
	if errorCode != ZPP_ERROR_OK {
		p.finish = true
	}

	// 若已有异常，不做error报错
	if existException() {
		return
	}

	// 触发错误或异常
	if !p.isQuiet() {
		switch errorCode {
		case ZPP_ERROR_FAILURE:
			// pass
		case ZPP_ERROR_WRONG_CALLBACK:
			message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", p.executeData.CalleeName(), p.idx, err)
			faults.InternalTypeErrorEx(p.isThrowEx(), message)
		case ZPP_ERROR_WRONG_CLASS:
			name := err
			message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", p.executeData.CalleeName(), p.idx, name, types.ZendZvalTypeName(p.arg))
			faults.InternalTypeErrorEx(p.isThrowEx(), message)
		case ZPP_ERROR_WRONG_ARG:
			expectedType := err
			message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", p.executeData.CalleeName(), p.idx, expectedType, types.ZendZvalTypeName(p.arg))
			faults.InternalTypeErrorEx(p.isThrowEx(), message)
		}
	}
}

func (p *baseParser) triggerDeprecated(errorCode int, err string) {
	switch errorCode {
	case ZPP_ERROR_WRONG_CALLBACK:
		message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", p.executeData.CalleeName(), p.idx, err)
		faults.ErrorEx(faults.E_DEPRECATED, message)
	}
}
