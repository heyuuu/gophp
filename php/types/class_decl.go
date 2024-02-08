package types

type ClassDecl struct {
	Name       string
	Flags      uint32
	Parent     string
	Interfaces []string
	Constants  []*ClassConstant
	Properties []*PropertyInfo
	Methods    []*Function

	Filename    string
	StartLineno uint32
	EndLineno   uint32
	DocComment  string
}

type UserClassDecl = ClassDecl
type InternalClassDecl = ClassDecl
