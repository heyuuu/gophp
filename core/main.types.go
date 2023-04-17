package core

import "github.com/heyuuu/gophp/zend"

type PhpStreamForZend struct {
	handle *PhpStream
}

var _ zend.IStream = (*PhpStreamForZend)(nil)

func NewPhpStreamForZend(handle *PhpStream) *PhpStreamForZend {
	return &PhpStreamForZend{handle: handle}
}

func (p *PhpStreamForZend) Isatty() bool {
	return false
}

func (p *PhpStreamForZend) Read(len_ int) []byte {
	buf := make([]byte, len_)
	size := PhpStreamRead(p.handle, buf, len_)
	return buf[:size]
}

func (p *PhpStreamForZend) FileSize() int {
	var stream *PhpStream = p.handle
	var ssb PhpStreamStatbuf

	/* File size reported by stat() may be inaccurate if stream filters are used.
	 * TODO: Should stat() be generally disabled if filters are used? */

	if stream.GetReadfilters().GetHead() != nil {
		return 0
	}
	if PhpStreamStat(stream, &ssb) == 0 {
		return ssb.GetSb().st_size
	}
	return 0
}

func (p *PhpStreamForZend) Close() {
	PhpStreamClose(p.handle)
}

func (p *PhpStreamForZend) GetHandle() any {
	return p.handle
}
