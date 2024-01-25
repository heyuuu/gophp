package types

type FunctionEntry struct {
	name     string
	handler  any
	argInfos []ArgInfo
}

func DefFunctionEntry(name string, handler any, argInfos []ArgInfo) FunctionEntry {
	return FunctionEntry{
		name:     name,
		handler:  handler,
		argInfos: argInfos,
	}
}

func (f FunctionEntry) Name() string { return f.name }
func (f FunctionEntry) Handler() any { return f.handler }
