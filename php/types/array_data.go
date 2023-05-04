package types

type ArrayData interface {
	Len() int
	Cap() int
	Exists(key ArrayKey) bool
	Find(key ArrayKey) *Zval
	Add(key ArrayKey, data *Zval) bool
	Update(key ArrayKey, data *Zval)
	Delete(key ArrayKey) bool
	Push(data *Zval) int
	Clean()
}
