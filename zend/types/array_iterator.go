package types

/**
 * HashTableIterator
 */
type HashTableIterator struct {
	ht  *Array
	pos ArrayPosition
}

func (this *HashTableIterator) GetHt() *Array              { return this.ht }
func (this *HashTableIterator) SetHt(value *Array)         { this.ht = value }
func (this *HashTableIterator) GetPos() ArrayPosition      { return this.pos }
func (this *HashTableIterator) SetPos(value ArrayPosition) { this.pos = value }
