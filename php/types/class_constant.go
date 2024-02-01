package types

// ClassConstant
type ClassConstant struct {
	name       string `get:""`
	value      Zval   `get:""`
	docComment string `get:""`
	ce         *Class `get:""`
	flags      uint32 `get:""`
}
