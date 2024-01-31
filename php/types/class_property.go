package types

// PropertyInfo
type PropertyInfo struct {
	name  string
	flags uint32
	ce    *Class
	typ   *TypeHint
}

func (p PropertyInfo) Flags() uint32   { return p.flags }
func (p PropertyInfo) Name() string    { return p.name }
func (p PropertyInfo) Ce() *Class      { return p.ce }
func (p PropertyInfo) Type() *TypeHint { return p.typ }
