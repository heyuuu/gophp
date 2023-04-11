package ascii

/**
 * Ascii 字符Set
 * 功能包含标准库 strings.asciiSet 功能，因后者非公开所有单独提出来
 */
type AsciiSet [8]uint32

func NewAsciiSet() *AsciiSet {
	return &AsciiSet{}
}

func (as *AsciiSet) Mark(c byte) {
	as[c/32] |= 1 << (c % 32)
}
func (as *AsciiSet) MarkString(s string) {
	for _, c := range []byte(s) {
		as[c/32] |= 1 << (c % 32)
	}
}
func (as *AsciiSet) Unmark(c byte) {
	as[c/32] &^= 1 << (c % 32)
}

func (as *AsciiSet) Contains(c byte) bool {
	return as[c/32]&1<<(c%32) != 0
}
