package builtin

type StrPtr SlicePtr[byte]

func NewStrPtr(str string) *StrPtr {
	return &StrPtr{items: []byte(str)}
}

// StrPtrReader

type StrPtrReader struct {
	str  string
	len_ int
	idx  int
}

func NewStrReader(str string) *StrPtrReader {
	return &StrPtrReader{str: str, len_: len(str), idx: 0}
}

func (r StrPtrReader) Copy() *StrPtrReader {
	return &r
}

func (r *StrPtrReader) IsValid() bool {
	return r.idx < r.len_
}

func (r *StrPtrReader) Curr() byte {
	if r.idx < r.len_ {
		return r.str[r.idx]
	}
	return 0
}

func (r *StrPtrReader) Read() byte {
	c := r.Curr()
	r.Next()
	return c
}

func (r *StrPtrReader) Next() {
	r.idx++
}
