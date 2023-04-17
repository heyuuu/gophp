package core

import (
	"github.com/heyuuu/gophp/core/pfmt"
	"strings"
)

func doSyslog(priority int, message string) {
	// todo syslog
}

func PhpSyslog(priority int, format string, args ...any) {
	/*
	 * don't rely on openlog() being called by syslog() if it's
	 * not already been done; call it ourselves and pass the
	 * correct parameters!
	 */
	if !(PG__().have_called_openlog) {
		PhpOpenlog(PG__().syslog_ident, 0, PG__().syslog_facility)
	}
	log := pfmt.Sprintf(format, args)
	if PG__().syslog_filter == PHP_SYSLOG_FILTER_RAW {
		/* Just send it directly to the syslog */
		doSyslog(priority, log)
		return
	}

	var buf strings.Builder
	for _, c := range []byte(log) {
		/* check for NVT ASCII only unless test disabled */

		if 0x20 <= c && c <= 0x7e {
			buf.WriteByte(c)
		} else if c >= 0x80 && PG__().syslog_filter != PHP_SYSLOG_FILTER_ASCII {
			buf.WriteByte(c)
		} else if c == '\n' {
			doSyslog(priority, buf.String())
			buf.Reset()
		} else if c < 0x20 && PG__().syslog_filter == PHP_SYSLOG_FILTER_ALL {
			buf.WriteByte(c)
		} else {
			var xdigits = "0123456789abcdef"
			buf.WriteString("\\x")
			buf.WriteByte(xdigits[c/0x10])
			c &= 0xf
			buf.WriteByte(xdigits[c])
		}
	}
	doSyslog(priority, buf.String())
}
