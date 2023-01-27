// <<generate>>

package zend

/**
 * TsHashTable
 */
type TsHashTable struct {
	hash   HashTable
	reader uint32
}

func (this *TsHashTable) GetHash() HashTable      { return this.hash }
func (this *TsHashTable) SetHash(value HashTable) { this.hash = value }
func (this *TsHashTable) GetReader() uint32       { return this.reader }
func (this *TsHashTable) SetReader(value uint32)  { this.reader = value }
