package types

import "github.com/heyuuu/gophp/zend"

/**
 * ZendInternalFunction
 */
const ZEND_MAX_RESERVED_RESOURCES = 6

var _ IFunction = (*ZendInternalFunction)(nil)

type ZendInternalFunction struct {
	functionHeader
	handler  zend.ZifHandler
	module   *zend.ModuleEntry
	reserved [ZEND_MAX_RESERVED_RESOURCES]any
}

func NewInternalFunction() *ZendInternalFunction {
	return &ZendInternalFunction{}
}

func NewInternalFunctionEx(funcName string, handler zend.ZifHandler) *ZendInternalFunction {
	f := &ZendInternalFunction{}
	f.functionName = NewString(funcName)
	f.handler = handler
	return f
}

func MakeInternalFunctionSimplify(handler zend.ZifHandler) ZendInternalFunction {
	return ZendInternalFunction{handler: handler}
}

func (this *ZendInternalFunction) InitByEntry(entry *ZendFunctionEntry) {
	this.handler = entry.Handler()
	this.functionName = NewString(entry.FuncName())
	this.prototype = nil
}

func (this *ZendInternalFunction) GetHandler() zend.ZifHandler       { return this.handler }
func (this *ZendInternalFunction) SetHandler(value zend.ZifHandler)  { this.handler = value }
func (this *ZendInternalFunction) GetModule() *zend.ModuleEntry      { return this.module }
func (this *ZendInternalFunction) SetModule(value *zend.ModuleEntry) { this.module = value }
func (this *ZendInternalFunction) GetReserved() []any                { return this.reserved }
func (this *ZendInternalFunction) SetReserved(value []any)           { this.reserved = value }
