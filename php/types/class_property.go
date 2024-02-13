package types

// PropertyInfo
type PropertyInfo struct {
	offset     uint32    `prop:""`
	flags      uint32    `get:""`
	name       string    `get:""`
	docComment string    `get:""`
	ce         *Class    `get:""`
	typ        *TypeHint `get:"Type"`
	defaultVal Zval
}

func NewPropertyInfo(flags uint32, name string, typ *TypeHint, defaultVal Zval) *PropertyInfo {
	// 默认访问等级为 public
	if flags&AccPppMask == 0 {
		flags |= AccPublic
	}

	return &PropertyInfo{
		flags:      flags,
		name:       name,
		typ:        typ,
		defaultVal: defaultVal,
	}
}

/* PropertyInfo.flags */
func (prop *PropertyInfo) AddFlags(value uint32)      { prop.flags |= value }
func (prop *PropertyInfo) HasFlags(value uint32) bool { return prop.flags&value != 0 }

func (prop *PropertyInfo) IsPublic() bool    { return prop.HasFlags(AccPublic) }
func (prop *PropertyInfo) IsProtected() bool { return prop.HasFlags(AccProtected) }
func (prop *PropertyInfo) IsPrivate() bool   { return prop.HasFlags(AccPrivate) }
func (prop *PropertyInfo) IsChanged() bool   { return prop.HasFlags(AccChanged) }
func (prop *PropertyInfo) IsStatic() bool    { return prop.HasFlags(AccStatic) }
func (prop *PropertyInfo) MarkIsChanged()    { prop.flags |= AccChanged }
