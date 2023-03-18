// <<generate>>

package standard

import (
	"sik/zend"
	"sik/zend/types"
)

/**
 * PhpUserFilterData
 */
type PhpUserFilterData struct {
	ce        *zend.ZendClassEntry
	classname *types.ZendString
}

// func MakePhpUserFilterData(ce *zend.ZendClassEntry, classname *zend.ZendString) PhpUserFilterData {
//     return PhpUserFilterData{
//         ce:ce,
//         classname:classname,
//     }
// }
func (this *PhpUserFilterData) GetCe() *zend.ZendClassEntry { return this.ce }

// func (this *PhpUserFilterData) SetCe(value *zend.ZendClassEntry) { this.ce = value }
func (this *PhpUserFilterData) GetClassname() *types.ZendString      { return this.classname }
func (this *PhpUserFilterData) SetClassname(value *types.ZendString) { this.classname = value }
