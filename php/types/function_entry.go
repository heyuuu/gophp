package types

type FunctionEntry struct {
	name    string
	handler any
}

func DefFunctionEntry(name string, handler any) FunctionEntry {
	return FunctionEntry{
		name:    name,
		handler: handler,
	}
}

func (f FunctionEntry) Name() string { return f.name }
func (f FunctionEntry) Handler() any { return f.handler }
