// <<generate>>

package standard

import (
	r "sik/builtin/file"
	"sik/core"
	"sik/zend"
	"sik/zend/types"
)

// Source: <ext/standard/file.h>

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

var ZifFdSet func(executeData *zend.ZendExecuteData, return_value *types.Zval)
var ZifFdIsset func(executeData *zend.ZendExecuteData, return_value *types.Zval)
var ZmStartupUserStreams func(type_ int, module_number int) int
var PhpSetSockBlocking func(socketd core.PhpSocketT, block int) int

const PHP_CSV_NO_ESCAPE = r.EOF
const META_DEF_BUFSIZE = 8192
const PHP_FILE_USE_INCLUDE_PATH = 1
const PHP_FILE_IGNORE_NEW_LINES = 2
const PHP_FILE_SKIP_EMPTY_LINES = 4
const PHP_FILE_APPEND = 8
const PHP_FILE_NO_DEFAULT_CONTEXT = 16

type PhpMetaTagsToken = int

const (
	TOK_EOF = 0
	TOK_OPENTAG
	TOK_CLOSETAG
	TOK_SLASH
	TOK_EQUAL
	TOK_SPACE
	TOK_ID
	TOK_STRING
	TOK_OTHER
)

var FileGlobals PhpFileGlobals

// Source: <ext/standard/file.c>

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
   |          Stig Bakken <ssb@php.net>                                   |
   |          Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   | PHP 4.0 patches by Thies C. Arntzen (thies@thieso.net)               |
   | PHP streams by Wez Furlong (wez@thebrainroom.com)                    |
   +----------------------------------------------------------------------+
*/

/* {{{ ZTS-stuff / Globals / Prototypes */

var LeStreamContext int = types.FAILURE
var FlockValues []int = []int{LOCK_SH, LOCK_EX, LOCK_UN}

/* {{{ proto bool flock(resource fp, int operation [, int &wouldblock])
   Portable file locking */

const PHP_META_UNSAFE = ".\\+*?[^]$() "

/* {{{ proto array get_meta_tags(string filename [, bool use_include_path])
   Extracts all meta tag content attributes from a file and returns an array */

const PHP_FILE_BUF_SIZE = 80

/* {{{ proto array file(string filename [, int flags[, resource context]])
   Read entire file into an array */

/* {{{ proto int fputcsv(resource fp, array fields [, string delimiter [, string enclosure [, string escape_char]]])
   Format line as CSV and write to file pointer */

/* {{{ proto string realpath(string path)
   Return the resolved path */

/* See http://www.w3.org/TR/html4/intro/sgmltut.html#h-3.2.2 */

const PHP_META_HTML401_CHARS = "-_.:"

/* {{{ php_next_meta_token
   Tokenizes an HTML file for get_meta_tags */

/* {{{ proto bool fnmatch(string pattern, string filename [, int flags])
   Match filename against pattern */

/* {{{ proto string sys_get_temp_dir()
   Returns directory path used for temporary files */
