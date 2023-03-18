// <<generate>>

package zend

import "sik/zend/types"

/**
 * ZendUserIterator
 */
type ZendUserIterator struct {
	it    ZendObjectIterator
	ce    *ZendClassEntry
	value types.Zval
}

// func MakeZendUserIterator(it ZendObjectIterator, ce *ZendClassEntry, value Zval) ZendUserIterator {
//     return ZendUserIterator{
//         it:it,
//         ce:ce,
//         value:value,
//     }
// }
func (this *ZendUserIterator) GetIt() ZendObjectIterator { return this.it }

// func (this *ZendUserIterator) SetIt(value ZendObjectIterator) { this.it = value }
func (this *ZendUserIterator) GetCe() *ZendClassEntry      { return this.ce }
func (this *ZendUserIterator) SetCe(value *ZendClassEntry) { this.ce = value }
func (this *ZendUserIterator) GetValue() types.Zval        { return this.value }

// func (this *ZendUserIterator) SetValue(value Zval) { this.value = value }
