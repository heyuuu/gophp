package zend

import (
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

/**
 * ZendExecuteData
 */
type ZendExecuteData struct {
	opline          *ZendOp
	call            *ZendExecuteData
	returnValue     *types.Zval
	func_           *types.ZendFunction
	This            types.Zval
	prevExecuteData *ZendExecuteData
	symbolTable     *types.Array
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

func (this *ZendExecuteData) Args(start int, len_ int) []*types.Zval {
	if len_ <= 0 {
		return nil
	}

	// todo 确认是否可简化为 slice 操作
	result := make([]*types.Zval, len_)
	for i := 0; i < len_; i++ {
		result[i] = this.Arg(start + i)
	}
	return result
}
func (this *ZendExecuteData) AllArgs() []*types.Zval {
	return this.Args(1, this.NumArgs())
}

func (this *ZendExecuteData) CheckNumArgsNone(forceStrict bool) bool {
	if forceStrict {
		zpp.CheckNumArgsNoneException(this)
	} else {
		zpp.CheckNumArgsNoneError(this)
	}
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
func (this *ZendExecuteData) GetFunc() *types.ZendFunction         { return this.func_ }
func (this *ZendExecuteData) SetFunc(value *types.ZendFunction)    { this.func_ = value }
func (this *ZendExecuteData) GetThis() *types.Zval                 { return &this.This }
func (this *ZendExecuteData) SetThis(zv *types.Zval)               { this.This = *zv }
func (this *ZendExecuteData) GetPrevExecuteData() *ZendExecuteData { return this.prevExecuteData }
func (this *ZendExecuteData) SetPrevExecuteData(value *ZendExecuteData) {
	this.prevExecuteData = value
}
func (this *ZendExecuteData) GetSymbolTable() *types.Array      { return this.symbolTable }
func (this *ZendExecuteData) SetSymbolTable(value *types.Array) { this.symbolTable = value }

func (this *ZendExecuteData) GetRunTimeCache() any      { return this.runTimeCache }
func (this *ZendExecuteData) SetRunTimeCache(value any) { this.symbolTable = value }
