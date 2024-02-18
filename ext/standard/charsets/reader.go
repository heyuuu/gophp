package charsets

type CharReader struct {
	charset Charset
	str     string
	cursor  int
	decoder CharDecoder
}

func NewCharReader(charset Charset, str string) *CharReader {
	return &CharReader{str: str, decoder: GetCharDecoder(charset)}
}
func (r *CharReader) Cursor() int {
	return r.cursor
}
func (r *CharReader) Left() string {
	if r.cursor >= len(r.str) {
		return ""
	}
	return r.str[r.cursor:]
}
func (r *CharReader) Skip(n int) {
	r.cursor += n
}
func (r *CharReader) Valid() bool {
	return r.cursor < len(r.str)
}
func (r *CharReader) NextChar() (uint, bool) {
	if r.cursor >= len(r.str) {
		return 0, false
	}

	c, n, ok := r.decoder(r.str[r.cursor:])
	r.cursor += n
	return c, ok
}
func (r *CharReader) NextCharEx() (uint, string, bool) {
	if r.cursor >= len(r.str) {
		return 0, "", false
	}

	c, n, ok := r.decoder(r.str[r.cursor:])
	match := r.str[r.cursor : r.cursor+n]
	r.cursor += n
	return c, match, ok
}
