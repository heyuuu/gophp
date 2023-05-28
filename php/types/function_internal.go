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

func (f *InternalFunction) GetOpArray() *ZendOpArray {
	panic("*InternalFunction is not *ZendOpArray")
}
func (f *InternalFunction) GetInternalFunction() *InternalFunction { return f }

func NewInternalFunction() *InternalFunction {
	return &InternalFunction{}
}

func NewInternalFunctionEx(funcName string, handler zend.ZifHandler) *InternalFunction {
	f := &InternalFunction{}
	f.functionName = funcName
	f.handler = handler
	return f
}

func MakeInternalFunctionSimplify(handler zend.ZifHandler) InternalFunction {
	return InternalFunction{handler: handler}
}

func (f *InternalFunction) InitByEntry(entry *FunctionEntry) {
	f.handler = entry.Handler()
	f.functionName = entry.FuncName()
	f.prototype = nil
}

func (f *InternalFunction) GetHandler() zend.ZifHandler       { return f.handler }
func (f *InternalFunction) SetHandler(value zend.ZifHandler)  { f.handler = value }
func (f *InternalFunction) GetModule() *zend.ModuleEntry      { return f.module }
func (f *InternalFunction) SetModule(value *zend.ModuleEntry) { f.module = value }
func (f *InternalFunction) GetReserved() []any                { return f.reserved }
func (f *InternalFunction) SetReserved(value []any)           { f.reserved = value }
