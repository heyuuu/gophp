package standard

// Source: <ext/standard/crypt.c>

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
   | Authors: Stig Bakken <ssb@php.net>                                   |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Rasmus Lerdorf <rasmus@php.net>                             |
   |          Pierre Joye <pierre@php.net>                                |
   +----------------------------------------------------------------------+
*/

/* sha512 crypt has the maximal salt length of 123 characters */

const PHP_MAX_SALT_LEN = 123

/* Used to check DES salts to ensure that they contain only valid characters */

const DES_INVALID_SALT_ERROR = "Supplied salt is not valid for DES. Possible bug in provided salt format."

var Itoa64 []uint8 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
