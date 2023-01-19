// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
)

// Source: <ext/standard/hrtime.h>

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
   | Author: Niklas Keller <kelunik@php.net>                              |
   | Author: Anatol Belski <ab@php.net>                                   |
   +----------------------------------------------------------------------+
*/

// #define HRTIME_H

// #define PHP_HRTIME_PLATFORM_POSIX       0

// #define PHP_HRTIME_PLATFORM_WINDOWS       0

// #define PHP_HRTIME_PLATFORM_APPLE       0

// #define PHP_HRTIME_PLATFORM_HPUX       0

// #define PHP_HRTIME_PLATFORM_AIX       0

// #define HRTIME_AVAILABLE       ( PHP_HRTIME_PLATFORM_POSIX || PHP_HRTIME_PLATFORM_WINDOWS || PHP_HRTIME_PLATFORM_APPLE || PHP_HRTIME_PLATFORM_HPUX || PHP_HRTIME_PLATFORM_AIX )

type PhpHrtimeT = uint64

// Source: <ext/standard/hrtime.c>

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
   | Author: Niklas Keller <kelunik@php.net>                              |
   | Author: Anatol Belski <ab@php.net>                                   |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "hrtime.h"

/* {{{ */

// #define NANO_IN_SEC       1000000000

/* }}} */

func _timerInit() int {
	/* Timer unavailable. */

	return -1
	return 0
}

/* {{{ */

func ZmStartupHrtime(type_ int, module_number int) int {
	if 0 > _timerInit() {
		core.PhpErrorDocref(nil, 1<<1, "Failed to initialize high-resolution timer")
		return zend.FAILURE
	}
	return zend.SUCCESS
}

/* }}} */

func _timerCurrent() PhpHrtimeT { return 0 }

// #define PHP_RETURN_HRTIME(t) RETURN_LONG ( ( zend_long ) t )

/* {{{ proto mixed hrtime([bool get_as_number = false])
   Returns an array of integers in form [seconds, nanoseconds] counted
   from an arbitrary point in time. If an optional boolean argument is
   passed, returns an integer on 64-bit platforms or float on 32-bit
   containing the current high-resolution time in nanoseconds. The
   delivered timestamp is monotonic and can not be adjusted. */

func ZifHrtime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	return_value.u1.type_info = 2
	return
}

/* }}} */

func PhpHrtimeCurrent() PhpHrtimeT { return _timerCurrent() }
