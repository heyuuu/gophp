package types

type UserClassEntry struct {
	Name       string
	Constants  []*ClassConstant
	Properties []*PropertyInfo
	Methods    []*Function
}
