package standard

// Source: <ext/standard/url.h>

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
   | Author: Jim Winstead <jimw@php.net>                                  |
   +----------------------------------------------------------------------+
*/

const PHP_URL_SCHEME = 0
const PHP_URL_HOST = 1
const PHP_URL_PORT = 2
const PHP_URL_USER = 3
const PHP_URL_PASS = 4
const PHP_URL_PATH = 5
const PHP_URL_QUERY = 6
const PHP_URL_FRAGMENT = 7
const PHP_QUERY_RFC1738 = 1
const PHP_QUERY_RFC3986 = 2

// Source: <ext/standard/url.c>

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
   | Author: Jim Winstead <jimw@php.net>                                  |
   +----------------------------------------------------------------------+
*/

/* {{{ free_url
 */

/* {{{ php_url_parse
 */

/* {{{ php_url_parse_ex2
 */

var UrlHexchars []uint8 = "0123456789ABCDEF"

/* {{{ php_url_encode
 */
