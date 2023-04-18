package types

import b "github.com/heyuuu/gophp/builtin"

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

/**
 * ArrayIterator 数组迭代器，HashTableIterator 的不完全兼容替代版本
 */
type ArrayIterator struct {
	arr *Array
	pos int
}

var _ b.Iterator[ArrayKey, *Zval] = (*ArrayIterator)(nil)

func (a ArrayIterator) Key() ArrayKey {
	//TODO implement me
	panic("implement me")
}

func (a ArrayIterator) Current() *Zval {
	//TODO implement me
	panic("implement me")
}

func (a ArrayIterator) Valid() bool {
	//TODO implement me
	panic("implement me")
}

func (a ArrayIterator) Next() {
	//TODO implement me
	panic("implement me")
}
