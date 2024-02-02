package types

type UserClassDecl struct {
	Name       string
	Constants  []*ClassConstant
	Properties []*PropertyInfo
	Methods    []*Function
}
