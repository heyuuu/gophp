package zend

import "sik/zend/types"

/**
 * ZendUserIterator
 */
type ZendUserIterator struct {
	it    ZendObjectIterator
	ce    *types.ClassEntry
	value types.Zval
}

// func MakeZendUserIterator(it ZendObjectIterator, ce *ClassEntry, value Zval) ZendUserIterator {
//     return ZendUserIterator{
//         it:it,
//         ce:ce,
//         value:value,
//     }
// }
func (this *ZendUserIterator) GetIt() ZendObjectIterator { return this.it }

// func (this *ZendUserIterator) SetIt(value ZendObjectIterator) { this.it = value }
func (this *ZendUserIterator) GetCe() *types.ClassEntry      { return this.ce }
func (this *ZendUserIterator) SetCe(value *types.ClassEntry) { this.ce = value }
func (this *ZendUserIterator) GetValue() types.Zval          { return this.value }

// func (this *ZendUserIterator) SetValue(value Zval) { this.value = value }
