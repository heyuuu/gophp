package php

type ZifHandler func(ex *ExecuteData, returnValue Val)

type FunctionEntry struct {
	name    string
	handler ZifHandler
}

func DefFunctionEntry(name string, handler ZifHandler) FunctionEntry {
	return FunctionEntry{
		name:    name,
		handler: handler,
	}
}

func (f FunctionEntry) Name() string        { return f.name }
func (f FunctionEntry) Handler() ZifHandler { return f.handler }
