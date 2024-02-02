package types

// ClassConstant
type ClassConstant struct {
	name       string `get:""`
	value      Zval   `get:""`
	docComment string `get:""`
	ce         *Class `get:""`
	flags      uint32 `get:""`
}

func NewClassConstant(name string, value Zval, docComment string, flags uint32) *ClassConstant {
	return &ClassConstant{name: name, value: value, docComment: docComment, flags: flags}
}
