package builtin

import "math"

const MaxBitsetSize = math.MaxInt32

type Bitset struct {
	slots []uint32
}

func NewBitset(size int) *Bitset {
	Assert(0 < size && size <= MaxBitsetSize)
	slotSize := (size + 31) / 32
	slots := make([]uint32, slotSize)
	return &Bitset{
		slots: slots,
	}
}

func (b *Bitset) Mark(pos int) {
	Assert(0 < pos && pos < len(b.slots)*32)
	b.slots[pos/32] |= 1 << (pos % 32)
}

func (b *Bitset) Unmark(pos int) {
	Assert(0 < pos && pos < len(b.slots)*32)
	b.slots[pos/32] &^= 1 << (pos % 32)
}

func (b *Bitset) Marked(pos int) bool {
	Assert(0 < pos && pos < len(b.slots)*32)
	return b.slots[pos/32]&1<<(pos%32) != 0
}
