// <<generate>>

package standard

import (
	"sik/zend"
)

// Source: <ext/standard/php_string.h>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Stig Sæther Bakken <ssb@php.net>                            |
   +----------------------------------------------------------------------+
*/

// #define PHP_STRING_H

var ZifNlLanginfo func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)

func Strnatcmp(a *byte, b *byte) int {
	return StrnatcmpEx(a, strlen(a), b, strlen(b), 0)
}
func Strnatcasecmp(a *byte, b *byte) int {
	return StrnatcmpEx(a, strlen(a), b, strlen(b), 1)
}
func PhpMblen(ptr *byte, len_ int) __auto__ { return mblen(ptr, len_) }

// #define php_mb_reset() php_ignore_value ( mblen ( NULL , 0 ) )
