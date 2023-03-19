package argparse

import (
	"sik/zend"
	"sik/zend/types"
)

func CurrExecuteData() ExecuteData { return zend.CurrEX() }

type OldParser struct {
	numArgs  int
	typeSpec string
	va       []any
	flags    int

	executeData ExecuteData
	minNumArgs  int
	maxNumArgs  int
	postVarargs int

	errorCode int
	finish    bool // 解析已终止 (可能是已解析完成或出现错误)
	idx       int
	arg       *types.Zval

	err *parseArgError

	vaReceiver *VaArgsReceiver
}

func OldParseStart(numArgs int, typeSpec string, va []any, flags int) *OldParser {
	minNumArgs, maxNumArgs, postVarargs, ok := checkTypeSpec(typeSpec)
	if !ok {
		return nil
	}
	executeData := CurrExecuteData()
	if numArgs > executeData.NumArgs() {
		zend.ZendParseParametersDebugError("could not obtain parameters for parsing")
		return nil
	}

	p := &OldParser{
		numArgs:     numArgs,
		typeSpec:    typeSpec,
		va:          va,
		flags:       flags,
		executeData: executeData,

		minNumArgs:  minNumArgs,
		maxNumArgs:  maxNumArgs,
		postVarargs: postVarargs,
	}

	if !CheckNumArgsEx(numArgs, p.executeData, p.minNumArgs, p.maxNumArgs, flags) {
		p.triggerError(ZPP_ERROR_FAILURE, "")
	}

	return p
}

func (p *OldParser) currArg() *types.Zval {
	return p.executeData.Arg(p.idx)
}

func (p *OldParser) isQuiet() bool { return p.flags&ZEND_PARSE_PARAMS_QUIET != 0 }
func (p *OldParser) isThrow() bool { return p.flags&ZEND_PARSE_PARAMS_THROW != 0 }

func (p *OldParser) triggerError(errorCode int, err string) {
	// 记录错误信息
	p.errorCode = errorCode
	if errorCode != ZPP_ERROR_OK {
		p.finish = true
	}

	// 触发错误或异常
	if !p.isQuiet() {
		switch errorCode {
		case ZPP_ERROR_WRONG_CALLBACK:
			zend.WrongCallbackError(p.idx, err, p.isThrow())
		case ZPP_ERROR_WRONG_CLASS:
			name := err
			zend.WrongParamClassError(p.idx, name, p.arg, p.isThrow())
		case ZPP_ERROR_WRONG_ARG:
			expectedType := err
			zend.WrongParamTypeError(p.idx, expectedType, p.arg, p.isThrow())
		}
	}
}

func (p *OldParser) HasError() bool {
	return p.errorCode != ZPP_ERROR_OK
}

func (p *OldParser) IsFinish() bool {
	return p.finish
}

func (p *OldParser) parsePrologue(deref bool, separate bool) {
	// todo
}
