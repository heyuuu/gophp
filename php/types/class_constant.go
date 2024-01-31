package types

// ClassConstant
type ClassConstant struct {
	name       string
	value      Zval
	docComment string
	ce         *Class
	flags      uint32
}

func (c ClassConstant) Name() string       { return c.name }
func (c ClassConstant) Value() Zval        { return c.value }
func (c ClassConstant) DocComment() string { return c.docComment }
func (c ClassConstant) Ce() *Class         { return c.ce }
func (c ClassConstant) Flags() uint32      { return c.flags }
