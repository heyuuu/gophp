package builtin

type StrPtr SlicePtr[byte]

func NewStrPtr(str string) *StrPtr {
	return &StrPtr{items: []byte(str)}
}
