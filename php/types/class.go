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
	name  string
	value Zval
	flags uint32
	ce    *Class
}

func (c ClassConstant) Flags() uint32 { return c.flags }
func (c ClassConstant) Name() string  { return c.name }
func (c ClassConstant) Ce() *Class    { return c.ce }
func (c ClassConstant) Value() Zval   { return c.value }

// Property
type Property struct {
	name  string
	flags uint32
	ce    *Class
}

func (p Property) Flags() uint32 { return p.flags }
func (p Property) Name() string  { return p.name }
func (p Property) Ce() *Class    { return p.ce }
