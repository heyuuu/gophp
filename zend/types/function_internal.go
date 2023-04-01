package types

import "github.com/heyuuu/gophp/zend"

/**
 * InternalFunction
 */
const ZEND_MAX_RESERVED_RESOURCES = 6

var _ IFunction = (*InternalFunction)(nil)

type InternalFunction struct {
	functionHeader
	handler  zend.ZifHandler
	module   *zend.ModuleEntry
	reserved [ZEND_MAX_RESERVED_RESOURCES]any
}

func NewInternalFunction() *InternalFunction {
	return &InternalFunction{}
}

func NewInternalFunctionEx(funcName string, handler zend.ZifHandler) *InternalFunction {
	f := &InternalFunction{}
	f.functionName = NewString(funcName)
	f.handler = handler
	return f
}

func MakeInternalFunctionSimplify(handler zend.ZifHandler) InternalFunction {
	return InternalFunction{handler: handler}
}

func (this *InternalFunction) InitByEntry(entry *FunctionEntry) {
	this.handler = entry.Handler()
	this.functionName = NewString(entry.FuncName())
	this.prototype = nil
}

func (this *InternalFunction) GetHandler() zend.ZifHandler       { return this.handler }
func (this *InternalFunction) SetHandler(value zend.ZifHandler)  { this.handler = value }
func (this *InternalFunction) GetModule() *zend.ModuleEntry      { return this.module }
func (this *InternalFunction) SetModule(value *zend.ModuleEntry) { this.module = value }
func (this *InternalFunction) GetReserved() []any                { return this.reserved }
func (this *InternalFunction) SetReserved(value []any)           { this.reserved = value }
