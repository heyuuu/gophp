package types

type ArrayData interface {
	Len() int
	Exists(key ArrayKey) bool
	Find(key ArrayKey) *Zval

	Add(key ArrayKey, value *Zval) bool
	Update(key ArrayKey, value *Zval)
	Delete(key ArrayKey) bool
	Push(value *Zval) int
	Clean()
}
