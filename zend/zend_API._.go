package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

/* End of parameter parsing API -- andrei */

const IS_CALLABLE_CHECK_SYNTAX_ONLY uint32 = 1 << 0
const IS_CALLABLE_CHECK_NO_ACCESS = 1 << 1
const IS_CALLABLE_CHECK_IS_STATIC = 1 << 2
const IS_CALLABLE_CHECK_SILENT uint32 = 1 << 3
const IS_CALLABLE_STRICT uint32 = IS_CALLABLE_CHECK_IS_STATIC

/* these variables are true statics/globals, and have to be mutex'ed on every access */
var ClassCleanupHandlers **types.ClassEntry
