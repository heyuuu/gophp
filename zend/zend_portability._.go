package zend

import "math"

/*
 * general definitions
 */

const ZEND_PATHS_SEPARATOR = ':'

/* all HAVE_XXX test have to be after the include of zend_config above */

const DL_UNLOAD = dlclose
const DL_FETCH_SYMBOL = dlsym
const DL_ERROR = dlerror

/* AIX requires this to be the first thing in the file.  */

const LONG_MAX = math.MaxInt32
