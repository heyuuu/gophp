package standard

import (
	"sik/zend"
)

// Source: <ext/standard/string.c>

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
   |          Stig Sæther Bakken <ssb@php.net>                          |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

/*
 * This define is __special__  here because some versions of libintl redefine setlocale __special__
 * to point to libintl_setlocale.  That's a ridiculous thing to do as far
 * as I am concerned, but with this define and __special__  the subsequent undef we
 * limit the damage to just the actual setlocale() call in this file
 * without turning zif_setlocale into zif_libintl_setlocale.  -Rasmus
 */

const PhpMySetlocale = setlocale

/* For str_getcsv() support */

/* For php_next_utf8_char() */

const STR_PAD_LEFT = 0
const STR_PAD_RIGHT = 1
const STR_PAD_BOTH = 2
const PHP_PATHINFO_DIRNAME = 1
const PHP_PATHINFO_BASENAME = 2
const PHP_PATHINFO_EXTENSION = 4
const PHP_PATHINFO_FILENAME = 8
const PHP_PATHINFO_ALL zend.ZendLong = PHP_PATHINFO_DIRNAME | PHP_PATHINFO_BASENAME | PHP_PATHINFO_EXTENSION | PHP_PATHINFO_FILENAME
const STR_STRSPN = 0
const STR_STRCSPN = 1

/* {{{ register_string_constants
 */

/* this is read-only, so it's ok */

var Hexconvtab []byte = "0123456789abcdef"

/* localeconv mutex */

/* {{{ php_bin2hex
 */

/* {{{ proto string bin2hex(string data)
   Converts the binary representation of data to hex */

/* {{{ proto string strtok([string str,] string token)
   Tokenize a string */

const _HEB_BLOCK_TYPE_ENG = 1
const _HEB_BLOCK_TYPE_HEB = 2

/* {{{ php_str_replace_in_subject
 */

const PHP_TAG_BUF_SIZE = 1023

/* {{{ php_tag_find
 *
 * Check if tag is in a set of tags
 *
 * states:
 *
 * 0 start tag
 * 1 first non-whitespace char seen
 */

/* {{{ proto string money_format(string format , float value)
   Convert monetary value(s) to string */

/* {{{ proto array str_split(string str [, int split_length])
   Convert a string to an array. If split_length is specified, break the string down into chunks each split_length characters long. */
