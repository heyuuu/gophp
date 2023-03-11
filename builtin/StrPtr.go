package builtin

type StrPtr SlicePtr[byte]

func NewStrPtr(str string) *StrPtr {
	return &StrPtr{items: []byte(str)}
}

// StrPtrReader

type StrReader struct {
	str  string
	len_ int
	idx  int
}

func NewStrReader(str string) *StrReader {
	return &StrReader{str: str, len_: len(str), idx: 0}
}

func (r *StrReader) IsValid() bool {
	return r.idx < r.len_
}

func (r *StrReader) Curr() byte {
	if r.idx < r.len_ {
		return r.str[r.idx]
	}
	return 0
}

func (r *StrReader) Next() {
	r.idx++
}
