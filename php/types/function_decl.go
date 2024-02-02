package types

type FunctionDecl struct {
	name     string
	handler  any
	argInfos []ArgInfo
}

func DefFunc(name string, handler any, argInfos []ArgInfo) FunctionDecl {
	return FunctionDecl{
		name:     name,
		handler:  handler,
		argInfos: argInfos,
	}
}

func (f FunctionDecl) Name() string { return f.name }
func (f FunctionDecl) Handler() any { return f.handler }
