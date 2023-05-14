package types

import "github.com/heyuuu/gophp/zend"

/**
 * PropertyInfo
 */
type PropertyInfo struct {
	offset     uint32
	flags      uint32
	name       *String
	docComment *String
	ce         *ClassEntry
	type_      ZendType
}

func (this *PropertyInfo) GetOffset() uint32           { return this.offset }
func (this *PropertyInfo) SetOffset(value uint32)      { this.offset = value }
func (this *PropertyInfo) GetFlags() uint32            { return this.flags }
func (this *PropertyInfo) SetFlags(value uint32)       { this.flags = value }
func (this *PropertyInfo) GetName() *String            { return this.name }
func (this *PropertyInfo) SetName(value *String)       { this.name = value }
func (this *PropertyInfo) GetDocComment() *String      { return this.docComment }
func (this *PropertyInfo) SetDocComment(value *String) { this.docComment = value }
func (this *PropertyInfo) GetCe() *ClassEntry          { return this.ce }
func (this *PropertyInfo) SetCe(value *ClassEntry)     { this.ce = value }
func (this *PropertyInfo) GetType() ZendType           { return this.type_ }
func (this *PropertyInfo) SetType(value ZendType)      { this.type_ = value }

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
func (this PropertyInfo) IsStatic() bool            { return this.HasFlags(zend.AccStatic) }
func (this PropertyInfo) IsProtected() bool         { return this.HasFlags(zend.AccProtected) }
func (this PropertyInfo) IsPrivate() bool           { return this.HasFlags(zend.AccPrivate) }
func (this PropertyInfo) IsPublic() bool            { return this.HasFlags(zend.AccPublic) }
func (this PropertyInfo) IsChanged() bool           { return this.HasFlags(zend.AccChanged) }
func (this *PropertyInfo) SetIsStatic(cond bool)    { this.SwitchFlags(zend.AccStatic, cond) }
func (this *PropertyInfo) SetIsProtected(cond bool) { this.SwitchFlags(zend.AccProtected, cond) }
func (this *PropertyInfo) SetIsPrivate(cond bool)   { this.SwitchFlags(zend.AccPrivate, cond) }
func (this *PropertyInfo) SetIsPublic(cond bool)    { this.SwitchFlags(zend.AccPublic, cond) }
func (this *PropertyInfo) SetIsChanged(cond bool)   { this.SwitchFlags(zend.AccChanged, cond) }
