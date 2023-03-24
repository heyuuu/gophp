package standard

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

const PHP_HRTIME_PLATFORM_POSIX = 0
const PHP_HRTIME_PLATFORM_WINDOWS = 0
const PHP_HRTIME_PLATFORM_APPLE = 0
const PHP_HRTIME_PLATFORM_HPUX = 0
const PHP_HRTIME_PLATFORM_AIX = 0
const HRTIME_AVAILABLE = PHP_HRTIME_PLATFORM_POSIX || PHP_HRTIME_PLATFORM_WINDOWS || PHP_HRTIME_PLATFORM_APPLE || PHP_HRTIME_PLATFORM_HPUX || PHP_HRTIME_PLATFORM_AIX

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

/* {{{ */

const NANO_IN_SEC = 1000000000

/* {{{ */

/* {{{ proto mixed hrtime([bool get_as_number = false])
   Returns an array of integers in form [seconds, nanoseconds] counted
   from an arbitrary point in time. If an optional boolean argument is
   passed, returns an integer on 64-bit platforms or float on 32-bit
   containing the current high-resolution time in nanoseconds. The
   delivered timestamp is monotonic and can not be adjusted. */
