package types

import b "github.com/heyuuu/gophp/builtin"

/**
 * ZendFcallInfo
 */
type ZendFcallInfo struct {
	size         int
	isInit       bool
	functionName Zval
	retval       *Zval
	params       []Zval
	paramCount   uint32
	object       *Object
	noSeparation bool
}

func EmptyFCallInfo() ZendFcallInfo { return ZendFcallInfo{} }

func InitFCallInfo(object *Object, retval *Zval, args ...*Zval) *ZendFcallInfo {
	if retval == nil {
		retval = NewZvalUndef()
	}
	var params []Zval
	if len(args) > 0 {
		params = make([]Zval, len(args))
		for i, arg := range args {
			(&params[i]).CopyValueFrom(arg)
		}
	}

	return &ZendFcallInfo{
		isInit: true,
		size:   b.SizeOf("ZendFcallInfo"),
		//functionName: Zval{}, // 默认值为 Undef 类型
		retval:       retval,
		params:       params,
		paramCount:   uint32(len(params)),
		object:       object,
		noSeparation: true,
	}
}

func (fci *ZendFcallInfo) IsInit() bool { return fci.isInit }
func (fci *ZendFcallInfo) UnInit() {
	fci.isInit = false
	fci.size = 0
}
func (fci *ZendFcallInfo) GetFunctionName() *Zval      { return &fci.functionName }
func (fci *ZendFcallInfo) ClearFunctionName()          { fci.functionName.SetUndef() }
func (fci *ZendFcallInfo) SetFunctionName(name string) { fci.functionName.SetStringVal(name) }
func (fci *ZendFcallInfo) SetFunctionNameZval(name *Zval) {
	fci.functionName.CopyValueFrom(name)
}
func (fci *ZendFcallInfo) GetRetval() *Zval           { return fci.retval }
func (fci *ZendFcallInfo) SetRetval(value *Zval)      { fci.retval = value }
func (fci *ZendFcallInfo) GetParams() []Zval          { return fci.params }
func (fci *ZendFcallInfo) SetParams(value []Zval)     { fci.params = value }
func (fci *ZendFcallInfo) GetObject() *Object         { return fci.object }
func (fci *ZendFcallInfo) SetObject(value *Object)    { fci.object = value }
func (fci *ZendFcallInfo) GetNoSeparation() bool      { return fci.noSeparation }
func (fci *ZendFcallInfo) SetNoSeparation(value bool) { fci.noSeparation = value }
func (fci *ZendFcallInfo) GetParamCount() uint32      { return fci.paramCount }
func (fci *ZendFcallInfo) SetParamCount(value uint32) { fci.paramCount = value }

/**
 * ZendFcallInfoCache
 */
type ZendFcallInfoCache struct {
	functionHandler IFunction
	callingScope    *ClassEntry
	calledScope     *ClassEntry
	object          *Object
}

func EmptyFcallInfoCache() ZendFcallInfoCache { return ZendFcallInfoCache{} }

func (fcc *ZendFcallInfoCache) GetFunctionHandler() IFunction { return fcc.functionHandler }
func (fcc *ZendFcallInfoCache) SetFunctionHandler(value IFunction) {
	fcc.functionHandler = value
}
func (fcc *ZendFcallInfoCache) GetCallingScope() *ClassEntry      { return fcc.callingScope }
func (fcc *ZendFcallInfoCache) SetCallingScope(value *ClassEntry) { fcc.callingScope = value }
func (fcc *ZendFcallInfoCache) GetCalledScope() *ClassEntry       { return fcc.calledScope }
func (fcc *ZendFcallInfoCache) SetCalledScope(value *ClassEntry)  { fcc.calledScope = value }
func (fcc *ZendFcallInfoCache) GetObject() *Object                { return fcc.object }
func (fcc *ZendFcallInfoCache) SetObject(value *Object)           { fcc.object = value }
