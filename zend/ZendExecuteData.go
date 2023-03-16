// <<generate>>

package zend

import b "sik/builtin"

/**
 * ZendExecuteData
 */
type ZendExecuteData struct {
	opline          *ZendOp
	call            *ZendExecuteData
	return_value    *Zval
	func_           *ZendFunction
	This            Zval
	prevExecuteData *ZendExecuteData
	symbol_table    *ZendArray
	run_time_cache  *any
}

func (this *ZendExecuteData) isStrictTypes() bool {
	return this != nil && this.func_ != nil && this.func_.IsStrictTypes()
}
func (this *ZendExecuteData) IsCallUseStrictTypes() bool { return this.isStrictTypes() }
func (this *ZendExecuteData) IsRetUseStrictTypes() bool  { return this.isStrictTypes() }
func (this *ZendExecuteData) IsArgUseStrictTypes() bool  { return this.prevExecuteData.isStrictTypes() }

func (this *ZendExecuteData) NumArgs() uint32 { return this.This.GetNumArgs() }

func (this *ZendExecuteData) CheckNumArgsError(minNumArgs int, maxNumArgs int) bool {
	return this.CheckNumArgs(minNumArgs, maxNumArgs, false)
}
func (this *ZendExecuteData) CheckNumArgsException(minNumArgs int, maxNumArgs int) bool {
	return this.CheckNumArgs(minNumArgs, maxNumArgs, true)
}
func (this *ZendExecuteData) CheckNumArgs(minNumArgs int, maxNumArgs int, forceStrict bool) bool {
	// 检查参数个数，若检查通过直接返回
	numArgs := int(this.NumArgs())
	if numArgs >= minNumArgs && numArgs <= maxNumArgs {
		return true
	}

	// 构建错误信息
	activeFunc := this.GetFunc()
	var callee string
	if activeFunc.GetScope() != nil {
		className := activeFunc.GetScope().Name()
		callee = className + "::" + activeFunc.GetFunctionName().GetStr()
	} else {
		callee = activeFunc.GetFunctionName().GetStr()
	}

	// 判断是否为 strict 模式
	var throwException = forceStrict || this.IsArgUseStrictTypes()
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
func (this *ZendExecuteData) SetReturnValue(value *Zval)           { this.return_value = value }
func (this *ZendExecuteData) GetFunc() *ZendFunction               { return this.func_ }
func (this *ZendExecuteData) SetFunc(value *ZendFunction)          { this.func_ = value }
func (this *ZendExecuteData) GetThis() *Zval                       { return &this.This }
func (this *ZendExecuteData) GetPrevExecuteData() *ZendExecuteData { return this.prevExecuteData }
func (this *ZendExecuteData) SetPrevExecuteData(value *ZendExecuteData) {
	this.prevExecuteData = value
}
func (this *ZendExecuteData) GetSymbolTable() *ZendArray      { return this.symbol_table }
func (this *ZendExecuteData) SetSymbolTable(value *ZendArray) { this.symbol_table = value }
