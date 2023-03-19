// <<generate>>

package standard

import (
	"sik/zend/types"
)

/**
 * PhpUserFilterData
 */
type PhpUserFilterData struct {
	ce        *types.ClassEntry
	classname *types.ZendString
}

// func MakePhpUserFilterData(ce *zend.ClassEntry, classname *zend.ZendString) PhpUserFilterData {
//     return PhpUserFilterData{
//         ce:ce,
//         classname:classname,
//     }
// }
func (this *PhpUserFilterData) GetCe() *types.ClassEntry { return this.ce }

// func (this *PhpUserFilterData) SetCe(value *zend.ClassEntry) { this.ce = value }
func (this *PhpUserFilterData) GetClassname() *types.ZendString      { return this.classname }
func (this *PhpUserFilterData) SetClassname(value *types.ZendString) { this.classname = value }
