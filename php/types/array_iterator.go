package types

import b "github.com/heyuuu/gophp/builtin"

/**
 * ArrayIterator 数组迭代器
 * 约定:
 * - ht 指向一个不可变数组，即要求 pos 对应的数组中的位置不会发生变化
 */
type ArrayIterator struct {
	ht  *Array
	pos ArrayPosition
}

var _ b.Iterator[ArrayKey, *Zval] = (*ArrayIterator)(nil)

func (iter *ArrayIterator) GetHt() *Array              { return iter.ht }
func (iter *ArrayIterator) SetHt(value *Array)         { iter.ht = value }
func (iter *ArrayIterator) GetPos() ArrayPosition      { return iter.pos }
func (iter *ArrayIterator) SetPos(value ArrayPosition) { iter.pos = value }

func (iter *ArrayIterator) Key() ArrayKey {
	//TODO implement me
	panic("implement me")
}

func (iter *ArrayIterator) Current() *Zval {
	//TODO implement me
	panic("implement me")
}

func (iter *ArrayIterator) Valid() bool {
	//TODO implement me
	panic("implement me")
}

func (iter *ArrayIterator) Next() {
	//TODO implement me
	panic("implement me")
}
