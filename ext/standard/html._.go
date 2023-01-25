// <<generate>>

package standard

import (
	"sik/zend"
)

// Source: <ext/standard/html.h>

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

const ENT_HTML_QUOTE_NONE = 0
const ENT_HTML_QUOTE_SINGLE = 1
const ENT_HTML_QUOTE_DOUBLE = 2
const ENT_HTML_IGNORE_ERRORS = 4
const ENT_HTML_SUBSTITUTE_ERRORS = 8
const ENT_HTML_DOC_TYPE_MASK = 16 | 32
const ENT_HTML_DOC_HTML401 = 0
const ENT_HTML_DOC_XML1 = 16
const ENT_HTML_DOC_XHTML = 32
const ENT_HTML_DOC_HTML5 = 16 | 32

/* reserve bit 6 */

const ENT_HTML_SUBSTITUTE_DISALLOWED_CHARS = 128
const ENT_COMPAT zend.ZendLong = ENT_HTML_QUOTE_DOUBLE
const ENT_QUOTES zend.ZendLong = ENT_HTML_QUOTE_DOUBLE | ENT_HTML_QUOTE_SINGLE
const ENT_NOQUOTES zend.ZendLong = ENT_HTML_QUOTE_NONE
const ENT_IGNORE zend.ZendLong = ENT_HTML_IGNORE_ERRORS
const ENT_SUBSTITUTE zend.ZendLong = ENT_HTML_SUBSTITUTE_ERRORS
const ENT_HTML401 = 0
const ENT_XML1 = 16
const ENT_XHTML = 32
const ENT_HTML5 zend.ZendLong = 16 | 32
const ENT_DISALLOWED = 128

// Source: <ext/standard/html.c>

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
   |          Jaakko Hyvätti <jaakko.hyvatti@iki.fi>                      |
   |          Wez Furlong    <wez@thebrainroom.com>                       |
   |          Gustavo Lopes  <cataphract@php.net>                         |
   +----------------------------------------------------------------------+
*/

/* Macro for disabling flag of translation of non-basic entities where this isn't supported.
 * Not appropriate for html_entity_decode/htmlspecialchars_decode */

/* valid as single byte character or leading byte */

/* whether it's actually valid depends on other stuff;
 * this macro cannot check for non-shortest forms, surrogates or
 * code points above 0x10FFFF */

/* {{{ get_default_charset
 */

/* {{{ traverse_for_entities
 * Auxiliary function to php_unescape_html_entities().
 * - The argument "all" determines if all numeric entities are decode or only those
 *   that correspond to quotes (depending on quote_style).
 */

/* {{{ find_entity_for_char */

const HTML_SPECIALCHARS = 0
const HTML_ENTITIES = 1

/* {{{ register_html_constants
 */
