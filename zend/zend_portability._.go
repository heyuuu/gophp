package zend

import "math"

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

const ZTS_V = 0
const LONG_MAX = 2147483647

var ZEND_INFINITY = math.Inf(1)

var ZEND_NAN = math.NaN()
