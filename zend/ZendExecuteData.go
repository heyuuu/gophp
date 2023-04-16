package zend

import (
	b "github.com/heyuuu/gophp/builtin"
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
	func_           types.IFunction
	This            types.Zval
	prevExecuteData *ZendExecuteData
	symbolTable     *types.Array
	//runtimeCache    *any
	// Ex
	runtimeCache []types.Zval
}

func NewExecuteData(callInfo uint32, fun types.IFunction, numArgs uint32, objectOrCalledScope any, runtimeCacheSize uint32) *ZendExecuteData {
	ex := &ZendExecuteData{}

	ex.func_ = fun
	ex.This.SetPtr(objectOrCalledScope)
	ex.This.SetTypeInfo(callInfo)
	ex.This.SetNumArgs(numArgs)

	ex.runtimeCache = make([]types.Zval, runtimeCacheSize)

	return ex
}
func (ex *ZendExecuteData) Extend(passedArgs uint32, additionalArgs uint32) {
	b.Assert(passedArgs == uint32(len(ex.runtimeCache)))
	newRuntimeCache := make([]types.Zval, len(ex.runtimeCache)+int(additionalArgs))
	copy(newRuntimeCache, ex.runtimeCache)
	ex.runtimeCache = newRuntimeCache
}

func (ex *ZendExecuteData) FunctionName() string {
	if ex == nil {
		return ""
	}

	activeFunc := ex.GetFunc()
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
func (ex *ZendExecuteData) ClassName() string {
	if ex == nil {
		return ""
	}

	activeFunc := ex.GetFunc()
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
func (ex *ZendExecuteData) CalleeName() string {
	if ex == nil {
		return ""
	}

	activeFunc := ex.GetFunc()
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

func (ex *ZendExecuteData) isStrictTypes() bool {
	return ex != nil && ex.func_ != nil && ex.func_.IsStrictTypes()
}
func (ex *ZendExecuteData) IsCallUseStrictTypes() bool { return ex.isStrictTypes() } // ZEND_RET_USES_STRICT_TYPES
func (ex *ZendExecuteData) IsRetUseStrictTypes() bool  { return ex.isStrictTypes() }
func (ex *ZendExecuteData) IsArgUseStrictTypes() bool  { return ex.prevExecuteData.isStrictTypes() } // ZEND_ARG_USES_STRICT_TYPES

func (ex *ZendExecuteData) NumArgs() int { return int(ex.This.GetNumArgs()) }
func (ex *ZendExecuteData) VarNum(n int) *types.Zval {
	if len(ex.runtimeCache) > n {
		return &ex.runtimeCache[n]
	}
	return nil
}
func (ex *ZendExecuteData) Arg(n int) *types.Zval { return ex.VarNum(n - 1) }
func (ex *ZendExecuteData) Args(start int, len_ int) []*types.Zval {
	if len_ <= 0 {
		return nil
	}

	// todo 确认是否可简化为 slice 操作
	result := make([]*types.Zval, len_)
	for i := 0; i < len_; i++ {
		result[i] = ex.Arg(start + i)
	}
	return result
}
func (ex *ZendExecuteData) AllArgs() []*types.Zval {
	return ex.Args(1, ex.NumArgs())
}

func (ex *ZendExecuteData) CheckNumArgsNone(forceStrict bool) bool {
	if forceStrict {
		return zpp.CheckNumArgsNoneException(ex)
	} else {
		return zpp.CheckNumArgsNoneError(ex)
	}
}

func (ex *ZendExecuteData) CallInfo() uint32 {
	return ex.This.GetTypeInfo()
}

/**
 * Getter/Setter
 */
func (ex *ZendExecuteData) GetOpline() *ZendOp                   { return ex.opline }
func (ex *ZendExecuteData) SetOpline(value *ZendOp)              { ex.opline = value }
func (ex *ZendExecuteData) GetCall() *ZendExecuteData            { return ex.call }
func (ex *ZendExecuteData) SetCall(value *ZendExecuteData)       { ex.call = value }
func (ex *ZendExecuteData) GetReturnValue() *types.Zval          { return ex.returnValue }
func (ex *ZendExecuteData) SetReturnValue(value *types.Zval)     { ex.returnValue = value }
func (ex *ZendExecuteData) GetFunc() types.IFunction             { return ex.func_ }
func (ex *ZendExecuteData) SetFunc(value types.IFunction)        { ex.func_ = value }
func (ex *ZendExecuteData) GetThis() *types.Zval                 { return &ex.This }
func (ex *ZendExecuteData) SetThis(zv *types.Zval)               { ex.This = *zv }
func (ex *ZendExecuteData) GetPrevExecuteData() *ZendExecuteData { return ex.prevExecuteData }
func (ex *ZendExecuteData) SetPrevExecuteData(value *ZendExecuteData) {
	ex.prevExecuteData = value
}
func (ex *ZendExecuteData) GetSymbolTable() *types.Array      { return ex.symbolTable }
func (ex *ZendExecuteData) SetSymbolTable(value *types.Array) { ex.symbolTable = value }

func (ex *ZendExecuteData) GetRuntimeCache() any { return ex.runtimeCache }
func (ex *ZendExecuteData) RuntimeCacheLen() int { return len(ex.runtimeCache) }
