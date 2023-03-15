// <<generate>>

package zend

import "math"

// Source: <Zend/zend_portability.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <zeev@php.net>                                |
   +----------------------------------------------------------------------+
*/

/*
 * general definitions
 */

const ZEND_PATHS_SEPARATOR = ':'

/* all HAVE_XXX test have to be after the include of zend_config above */

const RTLD_LAZY = 1
const RTLD_GLOBAL = 0
const PHP_RTLD_MODE = RTLD_LAZY
const DL_UNLOAD = dlclose
const DL_FETCH_SYMBOL = dlsym
const DL_ERROR = dlerror

/* AIX requires this to be the first thing in the file.  */

//const JMP_BUF = sigjmp_buf
type JMP_BUF struct{}

const ZTS_V = 0
const LONG_MAX = 2147483647
const LONG_MIN = -LONG_MAX - 1
const MAX_LENGTH_OF_DOUBLE = 32

/* We always define a __special__  function, even if there's a macro or expression we could
 * alias, so that using it in contexts where we can't make function calls
 * won't fail to compile on some machines and not others.
 */

var ZEND_INFINITY = math.Inf(1)

var ZEND_NAN = math.NaN()
