package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

/**
 * ZendExecuteData
 */
type ZendExecuteData struct {
	opline          *types.ZendOp
	call            *ZendExecuteData
	returnValue     *types.Zval
	func_           types.IFunction
	prevExecuteData *ZendExecuteData
	symbolTable     *types.Array

	// info
	callInfo  uint32
	numArgs   uint32
	thisClass *types.ClassEntry
	thisObj   *types.ZendObject

	//runtimeCache    *any
	runtimeCache []types.Zval
}

func NewExecuteData(callInfo uint32, fun types.IFunction, numArgs uint32, objectOrCalledScope any, runtimeCacheSize uint32) *ZendExecuteData {
	ex := &ZendExecuteData{}

	if callInfo&ZEND_CALL_HAS_THIS != 0 {
		b.Assert(objectOrCalledScope != nil)
	}

	ex.func_ = fun
	ex.SetCallInfo(callInfo)
	ex.SetScope(objectOrCalledScope)
	ex.SetNumArgs(numArgs)

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
		funcName := activeFunc.FunctionName()
		if funcName != "" {
			return funcName
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
		funcName := activeFunc.FunctionName()
		if funcName != "" {
			return scopePrefix + funcName
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

func (ex *ZendExecuteData) NumArgs() int                { return int(ex.numArgs) }
func (ex *ZendExecuteData) SetNumArgs(value uint32)     { ex.numArgs = value }
func (ex *ZendExecuteData) CallInfo() uint32            { return ex.callInfo }
func (ex *ZendExecuteData) SetCallInfo(callInfo uint32) { ex.callInfo = callInfo }

// scope
func (ex *ZendExecuteData) Scope() any {
	if ex.thisObj != nil {
		return ex.thisObj
	} else if ex.thisClass != nil {
		return ex.thisClass
	} else {
		return nil
	}
}

func (ex *ZendExecuteData) SetScope(scope any) {
	if scope == nil {
		ex.thisClass, ex.thisObj = nil, nil
		return
	}

	switch s := scope.(type) {
	case *types.ZendObject:
		ex.thisClass, ex.thisObj = s.GetCe(), s
	case *types.ClassEntry:
		ex.thisClass, ex.thisObj = s, nil
	default:
		panic("ZendExecuteData.SetScope() 只支持 *types.ZendObject、*types.ClassEntry 或 nil 参数")
	}
}
func (ex *ZendExecuteData) InScope() bool                 { return ex.thisClass != nil }
func (ex *ZendExecuteData) ThisObject() *types.ZendObject { return ex.thisObj }
func (ex *ZendExecuteData) ThisClass() *types.ClassEntry  { return ex.thisClass }

// 临时兼容，后续使用 ThisObject 替代
func (ex *ZendExecuteData) ThisObjectZval() *types.Zval {
	if ex.thisObj != nil {
		return types.NewZvalObject(ex.thisObj)
	}
	return nil
}

/**
 * Getter/Setter
 */
func (ex *ZendExecuteData) GetOpline() *types.ZendOp             { return ex.opline }
func (ex *ZendExecuteData) SetOpline(value *types.ZendOp)        { ex.opline = value }
func (ex *ZendExecuteData) GetCall() *ZendExecuteData            { return ex.call }
func (ex *ZendExecuteData) SetCall(value *ZendExecuteData)       { ex.call = value }
func (ex *ZendExecuteData) GetReturnValue() *types.Zval          { return ex.returnValue }
func (ex *ZendExecuteData) SetReturnValue(value *types.Zval)     { ex.returnValue = value }
func (ex *ZendExecuteData) GetFunc() types.IFunction             { return ex.func_ }
func (ex *ZendExecuteData) SetFunc(value types.IFunction)        { ex.func_ = value }
func (ex *ZendExecuteData) GetPrevExecuteData() *ZendExecuteData { return ex.prevExecuteData }
func (ex *ZendExecuteData) SetPrevExecuteData(value *ZendExecuteData) {
	ex.prevExecuteData = value
}
func (ex *ZendExecuteData) GetSymbolTable() *types.Array      { return ex.symbolTable }
func (ex *ZendExecuteData) SetSymbolTable(value *types.Array) { ex.symbolTable = value }

func (ex *ZendExecuteData) RuntimeCache() []types.Zval { return ex.runtimeCache }
func (ex *ZendExecuteData) GetRuntimeCache() any       { return ex.runtimeCache }
func (ex *ZendExecuteData) RuntimeCacheLen() int       { return len(ex.runtimeCache) }
