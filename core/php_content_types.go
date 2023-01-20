// <<generate>>

package core

import (
	g "sik/runtime/grammar"
	"sik/zend"
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

// #define PHP_CONTENT_TYPES_H

// #define DEFAULT_POST_CONTENT_TYPE       "application/x-www-form-urlencoded"

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

// # include "php.h"

// # include "SAPI.h"

// # include "rfc1867.h"

// # include "php_content_types.h"

/* {{{ php_post_entries[]
 */

var PhpPostEntries []SapiPostEntry = []SapiPostEntry{
	{
		"application/x-www-form-urlencoded",
		g.SizeOf("DEFAULT_POST_CONTENT_TYPE") - 1,
		SapiReadStandardFormData,
		PhpStdPostHandler,
	},
	{"multipart/form-data", g.SizeOf("MULTIPART_CONTENT_TYPE") - 1, nil, Rfc1867PostHandler},
	{nil, 0, nil, nil},
}

/* }}} */

func PhpDefaultPostReader() {
	if !(strcmp(sapi_globals.GetRequestInfo().GetRequestMethod(), "POST")) {
		if nil == sapi_globals.GetRequestInfo().GetPostEntry() {

			/* no post handler registered, so we just swallow the data */

			SapiReadStandardFormData()

			/* no post handler registered, so we just swallow the data */

		}
	}
}

/* }}} */

func PhpStartupSapiContentTypes() int {
	SapiRegisterDefaultPostReader(PhpDefaultPostReader)
	SapiRegisterTreatData(PhpDefaultTreatData)
	SapiRegisterInputFilter(PhpDefaultInputFilter, nil)
	return zend.SUCCESS
}

/* }}} */

func PhpSetupSapiContentTypes() int {
	SapiRegisterPostEntries(PhpPostEntries)
	return zend.SUCCESS
}

/* }}} */
