package types

/**
 * PropertyInfo
 */
type PropertyInfo struct {
	offset     uint32
	flags      uint32
	name       string
	docComment *string
	ce         *ClassEntry
	typ        ZendType
}

func NewPropertyInfo(offset uint32, flags uint32, name string, docComment *string, ce *ClassEntry, typ ZendType) *PropertyInfo {
	// 默认访问等级为 public
	if flags&AccPppMask == 0 {
		flags |= AccPublic
	}

	return &PropertyInfo{
		offset:     offset,
		flags:      flags,
		name:       name,
		docComment: docComment,
		ce:         ce,
		typ:        typ,
	}
}

func (this *PropertyInfo) SetOffset(value uint32) { this.offset = value }
func (this *PropertyInfo) SetType(value ZendType) { this.typ = value }

func (this *PropertyInfo) GetOffset() uint32      { return this.offset }
func (this *PropertyInfo) GetFlags() uint32       { return this.flags }
func (this *PropertyInfo) GetName() string        { return this.name }
func (this *PropertyInfo) GetDocComment() *string { return this.docComment }
func (this *PropertyInfo) GetCe() *ClassEntry     { return this.ce }
func (this *PropertyInfo) GetType() ZendType      { return this.typ }

/* PropertyInfo.flags */
func (this *PropertyInfo) AddFlags(value uint32)      { this.flags |= value }
func (this *PropertyInfo) SubFlags(value uint32)      { this.flags &^= value }
func (this *PropertyInfo) HasFlags(value uint32) bool { return this.flags&value != 0 }
func (this *PropertyInfo) SwitchFlags(value uint32, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this PropertyInfo) IsStatic() bool          { return this.HasFlags(AccStatic) }
func (this PropertyInfo) IsProtected() bool       { return this.HasFlags(AccProtected) }
func (this PropertyInfo) IsPrivate() bool         { return this.HasFlags(AccPrivate) }
func (this PropertyInfo) IsPublic() bool          { return this.HasFlags(AccPublic) }
func (this PropertyInfo) IsChanged() bool         { return this.HasFlags(AccChanged) }
func (this *PropertyInfo) SetIsChanged(cond bool) { this.SwitchFlags(AccChanged, cond) }
