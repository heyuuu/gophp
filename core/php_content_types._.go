// <<generate>>

package core

import (
	b "sik/builtin"
)

// Source: <main/php_content_types.h>

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
  | Author:                                                              |
  +----------------------------------------------------------------------+
*/

const DEFAULT_POST_CONTENT_TYPE = "application/x-www-form-urlencoded"

// Source: <main/php_content_types.c>

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
   | Author:                                                              |
   +----------------------------------------------------------------------+
*/

/* {{{ php_post_entries[]
 */

var PhpPostEntries []SapiPostEntry = []SapiPostEntry{
	{
		DEFAULT_POST_CONTENT_TYPE,
		b.SizeOf("DEFAULT_POST_CONTENT_TYPE") - 1,
		SapiReadStandardFormData,
		PhpStdPostHandler,
	},
	{MULTIPART_CONTENT_TYPE, b.SizeOf("MULTIPART_CONTENT_TYPE") - 1, nil, Rfc1867PostHandler},
	{nil, 0, nil, nil},
}

/* }}} */

/* }}} */

/* }}} */

/* }}} */
