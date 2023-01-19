// <<generate>>

package standard

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

// @type PhpRandomGlobals struct

// #define php_random_bytes_throw(b,s) php_random_bytes ( ( b ) , ( s ) , 1 )

// #define php_random_bytes_silent(b,s) php_random_bytes ( ( b ) , ( s ) , 0 )

// #define php_random_int_throw(min,max,result) php_random_int ( ( min ) , ( max ) , ( result ) , 1 )

// #define php_random_int_silent(min,max,result) php_random_int ( ( min ) , ( max ) , ( result ) , 0 )

// #define RANDOM_G(v) random_globals . v

var RandomGlobals PhpRandomGlobals
