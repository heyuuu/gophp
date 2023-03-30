package types

import b "github.com/heyuuu/gophp/builtin"

type PackedFillScope struct {
	ht *Array
	zv *Zval
}

func PackedFillStart(arr *Array) *PackedFillScope {
	b.Assert(arr.IsPacked())
	return &PackedFillScope{
		ht: arr,
		zv: NewZvalUndef(),
	}
}

func (s *PackedFillScope) FillNext() {
	s.ht.NextIndexInsertNew(s.zv)
	s.zv = NewZvalUndef()
	//s.bucket.SetH(s.idx)
	//s.bucket.SetKey(nil)
	//s.bucket++
	//s.idx++
}

func (s *PackedFillScope) FillEnd() {
	//s.ht.SetNNumUsed(s.idx)
	//s.ht.SetNNumOfElements(s.idx)
	//s.ht.SetNNextFreeElement(s.idx)
	//s.ht.SetNInternalPointer(0)
}

func (s *PackedFillScope) FillSet(arg *Zval) {
	ZVAL_COPY_VALUE(s.zv, arg)
}

func (s *PackedFillScope) FillSetNull() {
	//s.bucket.GetVal().SetNull()
	s.zv.SetNull()
}

func (s *PackedFillScope) FillSetLong(l int) {
	//s.bucket.GetVal().SetLong(l)
	s.zv.SetLong(l)
}

func (s *PackedFillScope) FillSetDouble(d float64) {
	//s.bucket.GetVal().SetDouble(d)
	s.zv.SetDouble(d)
}

func (s *PackedFillScope) FillSetInternedStr(str *String) {
	//s.bucket.GetVal().SetInternedString(str)
	s.zv.SetInternedString(str)
}

func (s *PackedFillScope) FillSetStringCopy(str *String) {
	//s.bucket.GetVal().SetStringCopy(str)
	s.zv.SetString(str)
}
