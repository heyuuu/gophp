// <<generate>>

package core

// Source: <main/SAPI.h>

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
   | Author:  Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

const SAPI_OPTION_NO_CHDIR = 1
const SAPI_POST_BLOCK_SIZE = 0x4000

type sapi_module_struct = _sapiModule

var sapi_module sapi_module_struct

/* Some values in this structure needs to be filled in before
 * calling sapi_activate(). We WILL change the `char *' entries,
 * so make sure that you allocate a separate buffer for them
 * and that you free them after sapi_deactivate().
 */

type _sapiGlobals = sapi_globals_struct

var sapi_globals sapi_globals_struct

/*
 * This is the preferred and maintained API for
 * operating on HTTP headers.
 */

type SapiHeaderOpEnum = int

const (
	SAPI_HEADER_REPLACE = iota
	SAPI_HEADER_ADD
	SAPI_HEADER_DELETE
	SAPI_HEADER_DELETE_ALL
	SAPI_HEADER_SET_STATUS
)

/* Deprecated functions. Use sapi_header_op instead. */

/* header_handler() constants */

const SAPI_HEADER_ADD SapiHeaderOpEnum = 1 << 0
const SAPI_HEADER_SENT_SUCCESSFULLY = 1
const SAPI_HEADER_DO_SEND = 2
const SAPI_HEADER_SEND_FAILED = 3
const SAPI_DEFAULT_MIMETYPE = "text/html"
const SAPI_DEFAULT_CHARSET *byte = PHP_DEFAULT_CHARSET
const SAPI_PHP_VERSION_HEADER *byte = "X-Powered-By: PHP/" + PHP_VERSION

// Source: <main/SAPI.c>

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
   | Original design:  Shane Caraveo <shane@caraveo.com>                  |
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// failed # include "ext/pcre/php_pcre.h"

/* True globals (no need for thread safety) */

/* {{{ proto bool header_register_callback(mixed callback)
   call a header function */

/*
 * Add charset on content-type header if the MIME type starts with
 * "text/", the default_charset directive is not empty and
 * there is not already a charset option in there.
 *
 * If "mimetype" is non-NULL, it should point to a pointer allocated
 * with emalloc().  If a charset is added, the string will be
 * re-allocated and the new length is returned.  If mimetype is
 * unchanged, 0 is returned.
 *
 */

/*
 * Called from php_request_startup() for every request.
 */

/*
 * since zend_llist_del_element only remove one matched item once,
 * we should remove them by ourself
 */
