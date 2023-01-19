// <<generate>>

package core

import "sik/zend"

// Source: <main/php_syslog.h>

/*
  +----------------------------------------------------------------------+
  | PHP Version 7                                                        |
  +----------------------------------------------------------------------+
  | Copyright (c) The PHP Group                                          |
  +----------------------------------------------------------------------+
  | This source file is subject to version 3.01 of the PHP license,      |
  | that is bundled with this package in the file LICENSE, and is        |
  | available through the world-wide-web at the following url:           |
  | http://www.php.net/license/3_01.txt                                  |
  | If you did not receive a copy of the PHP license and are unable to   |
  | obtain it through the world-wide-web, please send a note to          |
  | license@php.net so we can mail you a copy immediately.               |
  +----------------------------------------------------------------------+
  | Author:                                                              |
  +----------------------------------------------------------------------+
*/

// #define PHP_SYSLOG_H

// # include "php.h"

// # include < php_config . h >

// # include < syslog . h >

/* Syslog filters */

// #define PHP_SYSLOG_FILTER_ALL       0

// #define PHP_SYSLOG_FILTER_NO_CTRL       1

// #define PHP_SYSLOG_FILTER_ASCII       2

// #define PHP_SYSLOG_FILTER_RAW       3

var PhpOpenlog func(*byte, int, int)

// Source: <main/php_syslog.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Philip Prindeville <philipp@redfish-solutions.com>           |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include < string . h >

// # include < stdlib . h >

// # include "php.h"

// # include "php_syslog.h"

// # include "zend.h"

// # include "zend_smart_string.h"

/*
 * The SCO OpenServer 5 Development System (not the UDK)
 * defines syslog to std_syslog.
 */

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

	if CoreGlobals.GetHaveCalledOpenlog() == 0 {
		PhpOpenlog(CoreGlobals.GetSyslogIdent(), 0, CoreGlobals.GetSyslogFacility())
	}
	va_start(args, format)
	zend.ZendPrintfToSmartString(&fbuf, format, args)
	zend.SmartString0(&fbuf)
	va_end(args)
	if CoreGlobals.GetSyslogFilter() == 3 {

		/* Just send it directly to the syslog */

		syslog(priority, "%.*s", int(fbuf.len_), fbuf.c)
		zend.SmartStringFreeEx(&fbuf, 0)
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
			zend.SmartStringAppendcEx(&sbuf, c, 0)
		} else if c >= 0x80 && CoreGlobals.GetSyslogFilter() != 2 {
			zend.SmartStringAppendcEx(&sbuf, c, 0)
		} else if c == '\n' {
			syslog(priority, "%.*s", int(sbuf.len_), sbuf.c)
			zend.SmartStringReset(&sbuf)
		} else if c < 0x20 && CoreGlobals.GetSyslogFilter() == 0 {
			zend.SmartStringAppendcEx(&sbuf, c, 0)
		} else {
			var xdigits []byte = "0123456789abcdef"
			zend.SmartStringAppendlEx(&sbuf, "\\x", 2, 0)
			zend.SmartStringAppendcEx(&sbuf, xdigits[c/0x10], 0)
			c &= 0xf
			zend.SmartStringAppendcEx(&sbuf, xdigits[c], 0)
		}

		/* check for NVT ASCII only unless test disabled */

	}
	zend.SmartStringFreeEx(&fbuf, 0)
	zend.SmartStringFreeEx(&sbuf, 0)
}

/* }}} */
