// <<generate>>

package standard

import (
	b "sik/builtin"
	r "sik/builtin/file"
	"sik/core"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func PhpIptcPut1(fp *r.FILE, spool int, c uint8, spoolbuf **uint8) int {
	if spool > 0 {
		core.PUTC(c)
	}
	if spoolbuf != nil {
		b.PostInc(&(*(*spoolbuf))) = c
	}
	return c
}
func PhpIptcGet1(fp *r.FILE, spool int, spoolbuf **uint8) int {
	var c int
	var cc byte
	c = r.Getc(fp)
	if c == r.EOF {
		return r.EOF
	}
	if spool > 0 {
		cc = c
		core.PUTC(cc)
	}
	if spoolbuf != nil {
		b.PostInc(&(*(*spoolbuf))) = c
	}
	return c
}
func PhpIptcReadRemaining(fp *r.FILE, spool int, spoolbuf **uint8) int {
	for PhpIptcGet1(fp, spool, spoolbuf) != r.EOF {
		continue
	}
	return M_EOI
}
func PhpIptcSkipVariable(fp *r.FILE, spool int, spoolbuf **uint8) int {
	var length uint
	var c1 int
	var c2 int
	if b.Assign(&c1, PhpIptcGet1(fp, spool, spoolbuf)) == r.EOF {
		return M_EOI
	}
	if b.Assign(&c2, PhpIptcGet1(fp, spool, spoolbuf)) == r.EOF {
		return M_EOI
	}
	length = (uint8(c1) << 8) + uint8(c2)
	length -= 2
	for b.PostDec(&length) {
		if PhpIptcGet1(fp, spool, spoolbuf) == r.EOF {
			return M_EOI
		}
	}
	return 0
}
func PhpIptcNextMarker(fp *r.FILE, spool int, spoolbuf **uint8) int {
	var c int

	/* skip unimportant stuff */

	c = PhpIptcGet1(fp, spool, spoolbuf)
	if c == r.EOF {
		return M_EOI
	}
	for c != 0xff {
		if b.Assign(&c, PhpIptcGet1(fp, spool, spoolbuf)) == r.EOF {
			return M_EOI
		}
	}

	/* get marker byte, swallowing possible padding */

	for {
		c = PhpIptcGet1(fp, 0, 0)
		if c == r.EOF {
			return M_EOI
		} else if c == 0xff {
			PhpIptcPut1(fp, spool, uint8(c), spoolbuf)
		}
		if c != 0xff {
			break
		}
	}
	return uint(c)
}
func ZifIptcembed(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var iptcdata *byte
	var jpeg_file *byte
	var iptcdata_len int
	var jpeg_file_len int
	var spool zend.ZendLong = 0
	var fp *r.FILE
	var marker uint
	var done uint = 0
	var inx int
	var spoolbuf *types.ZendString = nil
	var poi *uint8 = nil
	var sb zend.ZendStatT
	var written types.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			iptcdata, iptcdata_len = fp.ParseString()
			jpeg_file, jpeg_file_len = fp.ParsePath()
			fp.StartOptional()
			spool = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if core.PhpCheckOpenBasedir(jpeg_file) != 0 {
		return_value.SetFalse()
		return
	}
	if iptcdata_len >= SIZE_MAX-b.SizeOf("psheader")-1025 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "IPTC data too large")
		return_value.SetFalse()
		return
	}
	if b.Assign(&fp, zend.VCWD_FOPEN(jpeg_file, "rb")) == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to open %s", jpeg_file)
		return_value.SetFalse()
		return
	}
	if spool < 2 {
		if zend.ZendFstat(fileno(fp), &sb) != 0 {
			return_value.SetFalse()
			return
		}
		spoolbuf = types.ZendStringSafeAlloc(1, iptcdata_len+b.SizeOf("psheader")+1024+1, sb.st_size, 0)
		poi = (*uint8)(spoolbuf.GetVal())
		memset(poi, 0, iptcdata_len+b.SizeOf("psheader")+sb.st_size+1024+1)
	}
	if PhpIptcGet1(fp, spool, b.Cond(poi != nil, &poi, 0)) != 0xff {
		r.Fclose(fp)
		if spoolbuf != nil {
			types.ZendStringEfree(spoolbuf)
		}
		return_value.SetFalse()
		return
	}
	if PhpIptcGet1(fp, spool, b.Cond(poi != nil, &poi, 0)) != 0xd8 {
		r.Fclose(fp)
		if spoolbuf != nil {
			types.ZendStringEfree(spoolbuf)
		}
		return_value.SetFalse()
		return
	}
	for done == 0 {
		marker = PhpIptcNextMarker(fp, spool, b.Cond(poi != nil, &poi, 0))
		if marker == M_EOI {
			break
		} else if marker != M_APP13 {
			PhpIptcPut1(fp, spool, uint8(marker), b.Cond(poi != nil, &poi, 0))
		}
		switch marker {
		case M_APP13:

			/* we are going to write a new APP13 marker, so don't output the old one */

			PhpIptcSkipVariable(fp, 0, 0)
			r.Fgetc(fp)
			PhpIptcReadRemaining(fp, spool, b.Cond(poi != nil, &poi, 0))
			done = 1
		case M_APP0:
			fallthrough
		case M_APP1:
			if written != 0 {

				/* don't try to write the data twice */

				break

				/* don't try to write the data twice */

			}
			written = 1
			PhpIptcSkipVariable(fp, spool, b.Cond(poi != nil, &poi, 0))
			if (iptcdata_len & 1) != 0 {
				iptcdata_len++
			}
			Psheader[2] = byte(iptcdata_len + 28>>8)
			Psheader[3] = iptcdata_len + 28&0xff
			for inx = 0; inx < 28; inx++ {
				PhpIptcPut1(fp, spool, Psheader[inx], b.Cond(poi != nil, &poi, 0))
			}
			PhpIptcPut1(fp, spool, uint8(iptcdata_len>>8), b.Cond(poi != nil, &poi, 0))
			PhpIptcPut1(fp, spool, uint8(iptcdata_len&0xff), b.Cond(poi != nil, &poi, 0))
			for inx = 0; inx < iptcdata_len; inx++ {
				PhpIptcPut1(fp, spool, iptcdata[inx], b.Cond(poi != nil, &poi, 0))
			}
		case M_SOS:

			/* we hit data, no more marker-inserting can be done! */

			PhpIptcReadRemaining(fp, spool, b.Cond(poi != nil, &poi, 0))
			done = 1
		default:
			PhpIptcSkipVariable(fp, spool, b.Cond(poi != nil, &poi, 0))
		}
	}
	r.Fclose(fp)
	if spool < 2 {
		spoolbuf = types.ZendStringTruncate(spoolbuf, poi-(*uint8)(spoolbuf.GetVal()), 0)
		return_value.SetString(spoolbuf)
		return
	} else {
		return_value.SetTrue()
		return
	}
}
func ZifIptcparse(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var inx int = 0
	var len_ int
	var tagsfound uint = 0
	var buffer *uint8
	var recnum uint8
	var dataset uint8
	var str *byte
	var key []*byte
	var str_len int
	var values types.Zval
	var element *types.Zval
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			str, str_len = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	buffer = (*uint8)(str)
	for inx < str_len {
		if buffer[inx] == 0x1c && (buffer[inx+1] == 0x1 || buffer[inx+1] == 0x2) {
			break
		} else {
			inx++
		}
	}
	for inx < str_len {
		if buffer[b.PostInc(&inx)] != 0x1c {
			break
		}
		if inx+4 >= str_len {
			break
		}
		dataset = buffer[b.PostInc(&inx)]
		recnum = buffer[b.PostInc(&inx)]
		if (buffer[inx] & uint8(0x80)) != 0 {
			if inx+6 >= str_len {
				break
			}
			len_ = (zend.ZendLong(buffer[inx+2]) << 24) + (zend.ZendLong(buffer[inx+3]) << 16) + (zend.ZendLong(buffer[inx+4]) << 8) + zend.ZendLong(buffer[inx+5])
			inx += 6
		} else {
			len_ = uint16(buffer[inx])<<8 | uint16(buffer[inx+1])
			inx += 2
		}
		if len_ > str_len || inx+len_ > str_len {
			break
		}
		core.Snprintf(key, b.SizeOf("key"), "%d#%03d", uint(dataset), uint(recnum))
		if tagsfound == 0 {
			zend.ArrayInit(return_value)
		}
		if b.Assign(&element, return_value.GetArr().KeyFind(b.CastStrAuto(key))) == nil {
			zend.ArrayInit(&values)
			element = return_value.GetArr().KeyUpdate(b.CastStrAuto(key), &values)
		}
		zend.AddNextIndexStringl(element, (*byte)(buffer+inx), len_)
		inx += len_
		tagsfound++
	}
	if tagsfound == 0 {
		return_value.SetFalse()
		return
	}
}
