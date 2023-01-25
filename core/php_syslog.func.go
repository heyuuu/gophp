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
	zend.SmartString0(&fbuf)
	va_end(args)
	if PG(syslog_filter) == PHP_SYSLOG_FILTER_RAW {

		/* Just send it directly to the syslog */

		syslog(priority, "%.*s", int(fbuf.len_), fbuf.c)
		zend.SmartStringFree(&fbuf)
		return
	}
	for ptr = fbuf.c; ; ptr++ {
		c = *ptr
		if c == '0' {
			syslog(priority, "%.*s", int(sbuf.len_), sbuf.c)
			break
		}

		/* check for NVT ASCII only unless test disabled */

		if 0x20 <= c && c <= 0x7e {
			zend.SmartStringAppendc(&sbuf, c)
		} else if c >= 0x80 && PG(syslog_filter) != PHP_SYSLOG_FILTER_ASCII {
			zend.SmartStringAppendc(&sbuf, c)
		} else if c == '\n' {
			syslog(priority, "%.*s", int(sbuf.len_), sbuf.c)
			zend.SmartStringReset(&sbuf)
		} else if c < 0x20 && PG(syslog_filter) == PHP_SYSLOG_FILTER_ALL {
			zend.SmartStringAppendc(&sbuf, c)
		} else {
			var xdigits []byte = "0123456789abcdef"
			zend.SmartStringAppendl(&sbuf, "\\x", 2)
			zend.SmartStringAppendc(&sbuf, xdigits[c/0x10])
			c &= 0xf
			zend.SmartStringAppendc(&sbuf, xdigits[c])
		}

		/* check for NVT ASCII only unless test disabled */

	}
	zend.SmartStringFree(&fbuf)
	zend.SmartStringFree(&sbuf)
}
