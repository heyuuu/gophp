// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * PhpUserFilterData
 */
type PhpUserFilterData struct {
	ce        *zend.ZendClassEntry
	classname *zend.ZendString
}

func (this PhpUserFilterData) GetCe() *zend.ZendClassEntry          { return this.ce }
func (this *PhpUserFilterData) SetCe(value *zend.ZendClassEntry)    { this.ce = value }
func (this PhpUserFilterData) GetClassname() *zend.ZendString       { return this.classname }
func (this *PhpUserFilterData) SetClassname(value *zend.ZendString) { this.classname = value }
