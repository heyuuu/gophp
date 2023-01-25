// <<generate>>

package standard

import (
	"sik/zend"
)

// Source: <ext/standard/php_random.h>

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
   | Authors: Sammy Kaye Powers <me@sammyk.me>                            |
   +----------------------------------------------------------------------+
*/

// #define PHP_RANDOM_H

func PhpRandomBytesThrow(b any, s int) int  { return PhpRandomBytes(b, s, 1) }
func PhpRandomBytesSilent(b any, s int) int { return PhpRandomBytes(b, s, 0) }
func PhpRandomIntThrow(min zend.ZendLong, max zend.ZendLong, result *zend.ZendLong) int {
	return PhpRandomInt(min, max, result, 1)
}
func PhpRandomIntSilent(min zend.ZendLong, max zend.ZendLong, result *zend.ZendLong) int {
	return PhpRandomInt(min, max, result, 0)
}
func RANDOM_G(v int) __auto__ { return RandomGlobals.v }

var RandomGlobals PhpRandomGlobals
