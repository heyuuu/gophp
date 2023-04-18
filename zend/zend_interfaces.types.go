package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

/**
 * ZendUserIterator
 */
type ZendUserIterator struct {
	it    ZendObjectIterator
	ce    *types2.ClassEntry
	value types2.Zval
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
func (this *ZendUserIterator) GetCe() *types2.ClassEntry      { return this.ce }
func (this *ZendUserIterator) SetCe(value *types2.ClassEntry) { this.ce = value }
func (this *ZendUserIterator) GetValue() types2.Zval          { return this.value }

// func (this *ZendUserIterator) SetValue(value Zval) { this.value = value }
