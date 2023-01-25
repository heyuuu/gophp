// <<generate>>

package core

// Source: <main/php_open_temporary_file.h>

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
   | Author: Zeev Suraski <zeev@php.net>                                  |
   +----------------------------------------------------------------------+
*/

const PHP_TMP_FILE_DEFAULT = 0
const PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_FALLBACK uint32 = 1 << 0
const PHP_TMP_FILE_SILENT = 1 << 1
const PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_EXPLICIT_DIR = 1 << 2
const PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ALWAYS uint32 = PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_FALLBACK | PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_EXPLICIT_DIR

/* for compatibility purpose */

const PHP_TMP_FILE_OPEN_BASEDIR_CHECK = PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_FALLBACK

// Source: <main/php_open_temporary_file.c>

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
   | Author: Zeev Suraski <zeev@php.net>                                  |
   +----------------------------------------------------------------------+
*/

const P_tmpdir = ""

/* {{{ php_do_open_temporary_file */

/* {{{ php_open_temporary_file
 *
 * Unlike tempnam(), the supplied dir argument takes precedence
 * over the TMPDIR environment variable
 * This function should do its best to return a file pointer to a newly created
 * unique file, on every platform.
 */
