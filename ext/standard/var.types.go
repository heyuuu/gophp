// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * PhpSerializeData
 */
type PhpSerializeData struct {
	ht zend.HashTable
	n  uint32
}

// func NewPhpSerializeData(ht zend.HashTable, n uint32) *PhpSerializeData {
//     return &PhpSerializeData{
//         ht:ht,
//         n:n,
//     }
// }
// func MakePhpSerializeData(ht zend.HashTable, n uint32) PhpSerializeData {
//     return PhpSerializeData{
//         ht:ht,
//         n:n,
//     }
// }
func (this *PhpSerializeData) GetHt() zend.HashTable { return this.ht }

// func (this *PhpSerializeData) SetHt(value zend.HashTable) { this.ht = value }
func (this *PhpSerializeData) GetN() uint32      { return this.n }
func (this *PhpSerializeData) SetN(value uint32) { this.n = value }
