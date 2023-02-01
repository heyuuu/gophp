// <<generate>>

package core

import (
	"sik/zend"
)

func PhpSyslog(priority int, format string, _ ...any) {
	var ptr *byte
	var c uint8
	var fbuf zend.SmartString = zend.SmartString{0}
	var sbuf zend.SmartString = zend.SmartString{0}
	var args va_list

	/*
	 * don't rely on openlog() being called by syslog() if it's
	 * not already been done; call it ourselves and pass the
	 * correct parameters!
	 */

	if !(PG(have_called_openlog)) {
		PhpOpenlog(PG(syslog_ident), 0, PG(syslog_facility))
	}
	va_start(args, format)
	zend.ZendPrintfToSmartString(&fbuf, format, args)
	fbuf.ZeroTail()
	va_end(args)
	if PG(syslog_filter) == PHP_SYSLOG_FILTER_RAW {

		/* Just send it directly to the syslog */

		syslog(priority, "%.*s", int(fbuf.GetLen()), fbuf.GetC())
		fbuf.Free()
		return
	}
	for ptr = fbuf.GetC(); ; ptr++ {
		c = *ptr
		if c == '0' {
			syslog(priority, "%.*s", int(sbuf.GetLen()), sbuf.GetC())
			break
		}

		/* check for NVT ASCII only unless test disabled */

		if 0x20 <= c && c <= 0x7e {
			sbuf.AppendByte(c)
		} else if c >= 0x80 && PG(syslog_filter) != PHP_SYSLOG_FILTER_ASCII {
			sbuf.AppendByte(c)
		} else if c == '\n' {
			syslog(priority, "%.*s", int(sbuf.GetLen()), sbuf.GetC())
			sbuf.Reset()
		} else if c < 0x20 && PG(syslog_filter) == PHP_SYSLOG_FILTER_ALL {
			sbuf.AppendByte(c)
		} else {
			var xdigits []byte = "0123456789abcdef"
			sbuf.AppendString("\\x")
			sbuf.AppendByte(xdigits[c/0x10])
			c &= 0xf
			sbuf.AppendByte(xdigits[c])
		}

		/* check for NVT ASCII only unless test disabled */

	}
	fbuf.Free()
	sbuf.Free()
}
