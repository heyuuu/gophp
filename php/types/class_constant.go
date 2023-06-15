package types

import "github.com/heyuuu/gophp/zend"

/**
 * ClassConstant
 */
type ClassConstant struct {
	value       Zval
	docComment  *String
	ce          *ClassEntry
	accessFlags uint32
}

func NewClassConstant(ce *ClassEntry, value *Zval, docComment *String) *ClassConstant {
	c := &ClassConstant{
		ce:         ce,
		docComment: docComment,
	}
	ZVAL_COPY_VALUE(&c.value, value)
	return c
}

func CopyClassConstant(c *ClassConstant) *ClassConstant {
	return &ClassConstant{
		ce:         c.ce,
		value:      c.value,
		docComment: c.docComment,
	}
}

func (c *ClassConstant) GetCe() *ClassEntry          { return c.ce }
func (c *ClassConstant) GetValue() *Zval             { return &c.value }
func (c *ClassConstant) GetDocComment() *String      { return c.docComment }
func (c *ClassConstant) GetAccessFlags() uint32      { return c.accessFlags }
func (c *ClassConstant) SetAccessFlags(value uint32) { c.accessFlags = value }

func (c *ClassConstant) IsVisited() bool { return c.accessFlags&zend.IS_CONSTANT_VISITED_MARK != 0 }
func (c *ClassConstant) MarkVisited()    { c.accessFlags |= zend.IS_CONSTANT_VISITED_MARK }
func (c *ClassConstant) ResetVisited()   { c.accessFlags &^= zend.IS_CONSTANT_VISITED_MARK }

func (c *ClassConstant) PriorLevel() uint32 { return c.accessFlags & AccPppMask }
func (c *ClassConstant) IsProtected() bool  { return c.accessFlags&AccProtected != 0 }
func (c *ClassConstant) IsPrivate() bool    { return c.accessFlags&AccPrivate != 0 }
func (c *ClassConstant) IsPublic() bool     { return c.accessFlags&AccPublic != 0 }
