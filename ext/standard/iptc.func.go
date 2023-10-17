package standard

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"io"
	"strings"
)

type iptcWriter struct {
	fp    *r.File
	spool int
	buf   io.Writer
}

func (w *iptcWriter) Write(data []byte) {
	if w.spool > 0 {
		core.PUTS(string(data))
	}
	if w.buf != nil {
		w.buf.Write(data)
	}
}
func (w *iptcWriter) WriteByte(c byte) {
	w.Write([]byte{c})
}
func (w *iptcWriter) ReadByte() (byte, bool) {
	c, ok := w.fp.GetC()
	if !ok {
		return 0, false
	}
	return c, true
}
func (w *iptcWriter) ReadWrite() (byte, bool) {
	c, ok := w.fp.GetC()
	if !ok {
		return 0, false
	}
	w.WriteByte(c)
	return c, true
}
func (w *iptcWriter) NextMarker() byte {
	var c byte
	var ok bool
	/* skip unimportant stuff */
	c, ok = w.ReadWrite()
	if !ok {
		return M_EOI
	}
	for c != 0xff {
		if c, ok = w.ReadWrite(); ok {
			return M_EOI
		}
	}

	/* get marker byte, swallowing possible padding */
	for {
		c, ok = w.ReadWrite()
		if !ok {
			return M_EOI
		} else if c == 0xff {
			w.WriteByte(c)
		}
		if c != 0xff {
			break
		}
	}
	return c
}
func (w *iptcWriter) SkipVariable() byte {
	var c1, c2 byte
	var ok bool
	if c1, ok = w.ReadWrite(); !ok {
		return M_EOI
	}
	if c2, ok = w.ReadWrite(); !ok {
		return M_EOI
	}
	length := int(c1)<<8 + int(c2) - 2
	for ; length > 0; length-- {
		if _, ok = w.ReadWrite(); !ok {
			return M_EOI
		}
	}
	return 0
}
func (w *iptcWriter) SkipVariableNoOutput() byte {
	var c1, c2 byte
	var ok bool
	if c1, ok = w.ReadByte(); !ok {
		return M_EOI
	}
	if c2, ok = w.ReadByte(); !ok {
		return M_EOI
	}
	length := int(c1)<<8 + int(c2) - 2
	for ; length > 0; length-- {
		if _, ok = w.ReadByte(); !ok {
			return M_EOI
		}
	}
	return 0
}
func (w *iptcWriter) ReadRemaining() {
	for {
		if _, ok := w.ReadWrite(); !ok {
			break
		}
	}
}

func ZifIptcembed(iptcData string, filename zpp.Path, _ zpp.Opt, spool int) *types.Zval {
	var fp *r.File
	var marker byte
	var done uint = 0
	var spoolBuf strings.Builder
	var sb zend.ZendStatT
	var written bool

	if core.PhpCheckOpenBasedir(filename) != 0 {
		return types.NewZvalFalse()
	}
	if len(iptcData) >= SIZE_MAX-b.SizeOf("psheader")-1025 {
		core.PhpErrorDocref("", faults.E_WARNING, "IPTC data too large")
		return types.NewZvalFalse()
	}
	if fp = zend.VCWD_FOPEN(filename, "rb"); fp == nil {
		core.PhpErrorDocref("", faults.E_WARNING, "Unable to open %s", filename)
		return types.NewZvalFalse()
	}
	defer func() { fp.Close() }()

	w := iptcWriter{fp: fp, spool: spool}
	if spool < 2 {
		if zend.ZendFstat(fileno(fp), &sb) != 0 {
			return types.NewZvalFalse()
		}
		spoolBuf.Grow(len(iptcData) + b.SizeOf(Psheader) + 1024 + 1 + sb.st_size)
		w.buf = &spoolBuf
	}

	if c, ok := w.ReadWrite(); !ok || c != 0xff {
		return types.NewZvalFalse()
	}
	if c, ok := w.ReadWrite(); !ok || c != 0xd8 {
		return types.NewZvalFalse()
	}

	for done == 0 {
		marker = w.NextMarker()
		if marker == M_EOI {
			break
		} else if marker != M_APP13 {
			w.WriteByte(marker)
		}
		switch marker {
		case M_APP13:
			/* we are going to write a new APP13 marker, so don't output the old one */
			w.SkipVariableNoOutput()
			fp.GetC()
			w.ReadRemaining()
			done = 1
		case M_APP0:
			fallthrough
		case M_APP1:
			if written {
				/* don't try to write the data twice */
				break
			}
			written = true
			w.SkipVariable()

			iptcdataLen := len(iptcData)
			if (iptcdataLen & 1) != 0 {
				iptcdataLen++
			}
			psheader := []byte(Psheader)
			psheader[2] = byte((iptcdataLen + 28) >> 8)
			psheader[3] = byte((iptcdataLen + 28) & 0xff)
			w.Write(psheader)
			w.WriteByte(uint8(iptcdataLen >> 8))
			w.WriteByte(uint8(iptcdataLen & 0xff))
			w.Write([]byte(iptcData))
		case M_SOS:
			/* we hit data, no more marker-inserting can be done! */
			w.ReadRemaining()
			done = 1
		default:
			w.SkipVariable()
		}
	}
	if spool < 2 {
		return types.NewZvalString(spoolBuf.String())
	} else {
		return types.NewZvalTrue()
	}
}
func ZifIptcparse(iptcdata string) (*types.Array, bool) {
	var idx uint = 0
	var len_ uint
	var tagsfound uint = 0
	var recnum uint8
	var dataset uint8

	data := iptcdata
	dataLen := uint(len(iptcdata))
	for idx < dataLen {
		if idx+1 < dataLen && data[idx] == 0x1c && (data[idx+1] == 0x1 || data[idx+1] == 0x2) {
			break
		} else {
			idx++
		}
	}

	arr := types.NewArray()
	for idx < dataLen {
		if data[lang.PostInc(&idx)] != 0x1c {
			break
		}
		if idx+4 >= dataLen {
			break
		}
		dataset = data[lang.PostInc(&idx)]
		recnum = data[lang.PostInc(&idx)]
		if (data[idx] & 0x80) != 0 {
			if idx+6 >= dataLen {
				break
			}
			len_ = uint(data[idx+2])<<24 + uint(data[idx+3])<<16 + uint(data[idx+4])<<8 + uint(data[idx+5])
			idx += 6
		} else {
			len_ = uint(data[idx])<<8 | uint(data[idx+1])
			idx += 2
		}
		if len_ > dataLen || idx+len_ > dataLen {
			break
		}

		key := fmt.Sprintf("%d#%03d", uint(dataset), uint(recnum))
		element := arr.KeyFind(key)
		if element == nil {
			element = types.NewZvalArray(nil)
			arr.KeyUpdate(key, element)
		}
		element.Array().Append(types.NewZvalString(data[idx : idx+len_]))

		idx += len_
		tagsfound++
	}
	if tagsfound == 0 {
		return nil, false
	}

	return arr, true
}
