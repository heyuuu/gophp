package types

import "github.com/heyuuu/gophp/zend"

/**
 * ClassConstant
 */
type ClassConstant struct {
	value      Zval
	docComment *String
	ce         *ClassEntry
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

func (c *ClassConstant) GetCe() *ClassEntry     { return c.ce }
func (c *ClassConstant) GetValue() *Zval        { return &c.value }
func (c *ClassConstant) GetDocComment() *String { return c.docComment }
func (c *ClassConstant) IsVisited() bool {
	return c.value.GetAccessFlags()&zend.IS_CONSTANT_VISITED_MARK != 0
}
func (c *ClassConstant) MarkVisited() {
	c.value.AddAccessFlags(zend.IS_CONSTANT_VISITED_MARK)
}
func (c *ClassConstant) ResetVisited() {
	c.value.SubAccessFlags(zend.IS_CONSTANT_VISITED_MARK)
}
