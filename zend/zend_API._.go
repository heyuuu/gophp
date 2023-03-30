package zend

import (
	"github.com/heyuuu/gophp/zend/types"
)

/* End of parameter parsing API -- andrei */

const IS_CALLABLE_CHECK_SYNTAX_ONLY uint32 = 1 << 0
const IS_CALLABLE_CHECK_NO_ACCESS = 1 << 1
const IS_CALLABLE_CHECK_IS_STATIC = 1 << 2
const IS_CALLABLE_CHECK_SILENT uint32 = 1 << 3
const IS_CALLABLE_STRICT uint32 = IS_CALLABLE_CHECK_IS_STATIC

//const ZEND_THIS *Zval = &(EX(This))

func ZEND_THIS(executeData *ZendExecuteData) *types.Zval {
	return executeData.GetThis()
}

/* these variables are true statics/globals, and have to be mutex'ed on every access */
var ModuleRequestStartupHandlers **ModuleEntry
var ModuleRequestShutdownHandlers **ModuleEntry
var ModulePostDeactivateHandlers **ModuleEntry
var ClassCleanupHandlers **types.ClassEntry
var DisabledClassNew = []types.ZendFunctionEntry{}
