// <<generate>>

package standard

// Source: <ext/standard/head.h>

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
   | Author: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                        |
   +----------------------------------------------------------------------+
*/

const COOKIE_EXPIRES = "; expires="
const COOKIE_MAX_AGE = "; Max-Age="
const COOKIE_DOMAIN = "; domain="
const COOKIE_PATH = "; path="
const COOKIE_SECURE = "; secure"
const COOKIE_HTTPONLY = "; HttpOnly"
const COOKIE_SAMESITE = "; SameSite="

var ZmActivateHead func(type_ int, module_number int) int

// Source: <ext/standard/head.c>

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
   | Author: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                        |
   +----------------------------------------------------------------------+
*/

// failed # include "ext/date/php_date.h"

/* Implementation of the language Header() function */

/* {{{ proto bool setcookie(string name [, string value [, int expires [, string path [, string domain [, bool secure[, bool httponly]]]]]])
               setcookie(string name [, string value [, array options]])
Send a cookie */

/* {{{ proto array headers_list(void)
   Return list of headers to be sent / already sent */
