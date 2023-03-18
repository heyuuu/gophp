// <<generate>>

package zend

import (
	"math"
	b "sik/builtin"
	"sik/zend/types"
)

/**
 * ZendExecuteData
 */
type ZendExecuteData struct {
	opline          *ZendOp
	call            *ZendExecuteData
	returnValue     *types.Zval
	func_           *ZendFunction
	This            types.Zval
	prevExecuteData *ZendExecuteData
	symbolTable     *types.ZendArray
	//runTimeCache    *any
	// Ex
	runTimeCache []types.Zval
}

func (this *ZendExecuteData) FunctionName() string {
	if this == nil {
		return ""
	}

	activeFunc := this.GetFunc()
	if activeFunc == nil {
		return ""
	}
	switch activeFunc.GetType() {
	case ZEND_USER_FUNCTION, ZEND_INTERNAL_FUNCTION:
		funcName := activeFunc.GetFunctionName()
		if funcName != nil {
			return funcName.GetStr()
		} else {
			return "main"
		}
	default:
		return ""
	}
}
func (this *ZendExecuteData) ClassName() string {
	if this == nil {
		return ""
	}

	activeFunc := this.GetFunc()
	if activeFunc == nil {
		return ""
	}
	switch activeFunc.GetType() {
	case ZEND_USER_FUNCTION, ZEND_INTERNAL_FUNCTION:
		ce := activeFunc.GetScope()
		if ce != nil {
			return ce.Name()
		} else {
			return ""
		}
	default:
		return ""
	}
}
func (this *ZendExecuteData) CalleeName() string {
	if this == nil {
		return ""
	}

	activeFunc := this.GetFunc()
	if activeFunc == nil {
		return ""
	}

	switch activeFunc.GetType() {
	case ZEND_USER_FUNCTION, ZEND_INTERNAL_FUNCTION:
		// scopePrefix(className + "::")
		scopePrefix := ""
		if activeFunc.GetScope() != nil {
			scopePrefix = activeFunc.GetScope().Name() + "::"
		}

		// func name
		funcName := activeFunc.GetFunctionName()
		if funcName != nil {
			return scopePrefix + funcName.GetStr()
		} else {
			return "main"
		}
	default:
		return ""
	}
}

func (this *ZendExecuteData) isStrictTypes() bool {
	return this != nil && this.func_ != nil && this.func_.IsStrictTypes()
}
func (this *ZendExecuteData) IsCallUseStrictTypes() bool { return this.isStrictTypes() } // ZEND_RET_USES_STRICT_TYPES
func (this *ZendExecuteData) IsRetUseStrictTypes() bool  { return this.isStrictTypes() }
func (this *ZendExecuteData) IsArgUseStrictTypes() bool  { return this.prevExecuteData.isStrictTypes() } // ZEND_ARG_USES_STRICT_TYPES

func (this *ZendExecuteData) NumArgs() int { return int(this.This.GetNumArgs()) }
func (this *ZendExecuteData) VarNum(n int) *types.Zval {
	if len(this.runTimeCache) > n {
		return &this.runTimeCache[n]
	}
	return nil
}
func (this *ZendExecuteData) Arg(n int) *types.Zval { return this.VarNum(n - 1) }

func (this *ZendExecuteData) CheckNumArgsError(minNumArgs int, maxNumArgs int) bool {
	return this.CheckNumArgs(minNumArgs, maxNumArgs, false)
}
func (this *ZendExecuteData) CheckNumArgsException(minNumArgs int, maxNumArgs int) bool {
	return this.CheckNumArgs(minNumArgs, maxNumArgs, true)
}
func (this *ZendExecuteData) CheckMinNumArgs(minNumArgs int, forceStrict bool) bool {
	return this.CheckNumArgs(minNumArgs, math.MaxInt, forceStrict)
}
func (this *ZendExecuteData) CheckNumArgs(minNumArgs int, maxNumArgs int, forceStrict bool) bool {
	// 检查参数个数，若检查通过直接返回
	return this.CheckNumArgsEx(this.NumArgs(), minNumArgs, maxNumArgs, forceStrict)
}

func (this *ZendExecuteData) CheckNumArgsEx(numArgs int, minNumArgs int, maxNumArgs int, forceStrict bool) bool {
	// 检查参数个数，若检查通过直接返回
	if numArgs >= minNumArgs && numArgs <= maxNumArgs {
		return true
	}

	// 判断是否为 strict 模式
	var throwException = forceStrict || this.IsArgUseStrictTypes()

	// 构建错误信息
	callee := this.CalleeName()
	if minNumArgs == maxNumArgs {
		ZendInternalArgumentCountError(throwException, "%s() expects exactly %d parameter%s, %d given", callee, minNumArgs, b.Cond(minNumArgs == 1, "", "s"), numArgs)
	} else if numArgs < minNumArgs {
		ZendInternalArgumentCountError(throwException, "%s() expects at least %d parameter%s, %d given", callee, minNumArgs, b.Cond(minNumArgs == 1, "", "s"), numArgs)
	} else { // numArgs > maxNumArgs
		ZendInternalArgumentCountError(throwException, "%s() expects at most %d parameter%s, %d given", callee, maxNumArgs, b.Cond(maxNumArgs == 1, "", "s"), numArgs)
	}
	return false
}

/**
 * Getter/Setter
 */
func (this *ZendExecuteData) GetOpline() *ZendOp                   { return this.opline }
func (this *ZendExecuteData) SetOpline(value *ZendOp)              { this.opline = value }
func (this *ZendExecuteData) GetCall() *ZendExecuteData            { return this.call }
func (this *ZendExecuteData) SetCall(value *ZendExecuteData)       { this.call = value }
func (this *ZendExecuteData) GetReturnValue() *types.Zval          { return this.returnValue }
func (this *ZendExecuteData) SetReturnValue(value *types.Zval)     { this.returnValue = value }
func (this *ZendExecuteData) GetFunc() *ZendFunction               { return this.func_ }
func (this *ZendExecuteData) SetFunc(value *ZendFunction)          { this.func_ = value }
func (this *ZendExecuteData) GetThis() *types.Zval                 { return &this.This }
func (this *ZendExecuteData) SetThis(zv *types.Zval)               { this.This = *zv }
func (this *ZendExecuteData) GetPrevExecuteData() *ZendExecuteData { return this.prevExecuteData }
func (this *ZendExecuteData) SetPrevExecuteData(value *ZendExecuteData) {
	this.prevExecuteData = value
}
func (this *ZendExecuteData) GetSymbolTable() *types.ZendArray      { return this.symbolTable }
func (this *ZendExecuteData) SetSymbolTable(value *types.ZendArray) { this.symbolTable = value }

func (this *ZendExecuteData) GetRunTimeCache() any      { return this.runTimeCache }
func (this *ZendExecuteData) SetRunTimeCache(value any) { this.symbolTable = value }
