package types

// Class
type Class struct {
	name string

	// info.internal
	moduleNumber int
	// info.user
	blockInfo
}

func (c *Class) Name() string { return c.name }

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

// PropertyInfo
type PropertyInfo struct {
	name  string
	flags uint32
	ce    *Class
}

func (p PropertyInfo) Flags() uint32 { return p.flags }
func (p PropertyInfo) Name() string  { return p.name }
func (p PropertyInfo) Ce() *Class    { return p.ce }
