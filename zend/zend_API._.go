package zend

/* Parameter parsing API -- andrei */

const ZEND_PARSE_PARAMS_QUIET = 1 << 1
const ZEND_PARSE_PARAMS_THROW = 1 << 2

/* End of parameter parsing API -- andrei */

const IS_CALLABLE_CHECK_SYNTAX_ONLY uint32 = 1 << 0
const IS_CALLABLE_CHECK_NO_ACCESS = 1 << 1
const IS_CALLABLE_CHECK_IS_STATIC = 1 << 2
const IS_CALLABLE_CHECK_SILENT uint32 = 1 << 3
const IS_CALLABLE_STRICT uint32 = IS_CALLABLE_CHECK_IS_STATIC

//const ZEND_THIS *Zval = &(EX(This))

func ZEND_THIS(executeData *ZendExecuteData) *Zval {
	return &executeData.This
}

/* Fast parameter parsing API */

type ZendExpectedType = int

const (
	Z_EXPECTED_LONG = iota
	Z_EXPECTED_BOOL
	Z_EXPECTED_STRING
	Z_EXPECTED_ARRAY
	Z_EXPECTED_FUNC
	Z_EXPECTED_RESOURCE
	Z_EXPECTED_PATH
	Z_EXPECTED_OBJECT
	Z_EXPECTED_DOUBLE
	Z_EXPECTED_LAST
)
const ZPP_ERROR_OK = 0
const ZPP_ERROR_FAILURE = 1
const ZPP_ERROR_WRONG_CALLBACK = 2
const ZPP_ERROR_WRONG_CLASS = 3
const ZPP_ERROR_WRONG_ARG = 4
const ZPP_ERROR_WRONG_COUNT = 5

/* these variables are true statics/globals, and have to be mutex'ed on every access */

var ModuleRegistry HashTable
var ModuleRequestStartupHandlers **ZendModuleEntry
var ModuleRequestShutdownHandlers **ZendModuleEntry
var ModulePostDeactivateHandlers **ZendModuleEntry
var ClassCleanupHandlers **ZendClassEntry
var DisabledClassNew []ZendFunctionEntry = []ZendFunctionEntry{}
