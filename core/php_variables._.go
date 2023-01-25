// <<generate>>

package core

import (
	"sik/zend"
)

// Source: <main/php_variables.h>

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
   | Authors: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                       |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

const PARSE_POST = 0
const PARSE_GET = 1
const PARSE_COOKIE = 2
const PARSE_STRING = 3
const PARSE_ENV = 4
const PARSE_SERVER = 5
const PARSE_SESSION = 6

/* binary-safe version */

const NUM_TRACK_VARS = 6

// Source: <main/php_variables.c>

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
   | Authors: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                       |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

/* for systems that need to override reading of environment variables */

var PhpImportEnvironmentVariables func(array_ptr *zend.Zval) = _phpImportEnvironmentVariables

/* binary-safe version */

type PostVarData = PostVarDataT

/* {{{ php_build_argv
 */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* Upgly hack to fix HTTP_PROXY issue, see bug #72573 */
