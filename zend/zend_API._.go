// <<generate>>

package zend

// Source: <Zend/zend_API.h>

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
   |          Andrei Zmievski <andrei@php.net>                            |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* Name macros */

/* Declaration macros */

/* internal function to efficiently copy parameters when executing __call() */

/* Parameter parsing API -- andrei */

const ZEND_PARSE_PARAMS_QUIET = 1 << 1
const ZEND_PARSE_PARAMS_THROW = 1 << 2

/* End of parameter parsing API -- andrei */

const IS_CALLABLE_CHECK_SYNTAX_ONLY uint32 = 1 << 0
const IS_CALLABLE_CHECK_NO_ACCESS = 1 << 1
const IS_CALLABLE_CHECK_IS_STATIC = 1 << 2
const IS_CALLABLE_CHECK_SILENT uint32 = 1 << 3
const IS_CALLABLE_STRICT uint32 = IS_CALLABLE_CHECK_IS_STATIC
const ZEND_THIS *Zval = &(EX(This))
const WRONG_PARAM_COUNT = ZEND_WRONG_PARAM_COUNT()

/** Build zend_call_info/cache from a zval*
 *
 * Caller is responsible to provide a return value (fci->retval), otherwise the we will crash.
 * In order to pass parameters the following members need to be set:
 * fci->param_count = 0;
 * fci->params = NULL;
 * The callable_name argument may be NULL.
 * Set check_flags to IS_CALLABLE_STRICT for every new usage!
 */

/** Clear arguments connected with zend_fcall_info *fci
 * If free_mem is not zero then the params array gets free'd as well
 */

/** Save current arguments from zend_fcall_info *fci
 * params array will be set to NULL
 */

/** Free arguments connected with zend_fcall_info *fci andset back saved ones.
 */

/** Set or clear the arguments in the zend_call_info struct taking care of
 * refcount. If args is NULL and arguments are set then those are cleared.
 */

/** Set arguments in the zend_fcall_info struct taking care of refcount.
 * If argc is 0 the arguments which are set will be cleared, else pass
 * a variable amount of zval** arguments.
 */

/** Set arguments in the zend_fcall_info struct taking care of refcount.
 * If argc is 0 the arguments which are set will be cleared, else pass
 * a variable amount of zval** arguments.
 */

/** Set arguments in the zend_fcall_info struct taking care of refcount.
 * If argc is 0 the arguments which are set will be cleared, else pass
 * a variable amount of zval** arguments.
 */

/** Call a function using information created by zend_fcall_info_init()/args().
 * If args is given then those replace the argument info in fci is temporarily.
 */

/* For compatibility */

const ZEND_MINIT = ZEND_MODULE_STARTUP_N
const ZEND_MSHUTDOWN = ZEND_MODULE_SHUTDOWN_N
const ZEND_RINIT = ZEND_MODULE_ACTIVATE_N
const ZEND_RSHUTDOWN = ZEND_MODULE_DEACTIVATE_N
const ZEND_MINFO = ZEND_MODULE_INFO_N
const ZEND_MSHUTDOWN_FUNCTION = ZEND_MODULE_SHUTDOWN_D
const ZEND_RINIT_FUNCTION = ZEND_MODULE_ACTIVATE_D
const ZEND_RSHUTDOWN_FUNCTION = ZEND_MODULE_DEACTIVATE_D
const ZEND_MINFO_FUNCTION = ZEND_MODULE_INFO_D
const ZEND_GSHUTDOWN_FUNCTION = ZEND_MODULE_GLOBALS_DTOR_D

/* May modify arg in-place. Will free arg in failure case (and take ownership in success case).
 * Prefer using the ZEND_TRY_ASSIGN_* macros over these APIs. */

/* Initializes a reference to an empty array and returns dereferenced zval,
 * or NULL if the initialization failed. */

/* Fast parameter parsing API */

const FAST_ZPP = 1

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

/* old "|" */

/* old "a" */

/* old "A" */

/* old "b" */

/* old "C" */

/* old "d" */

/* old "f" */

/* old "h" */

/* old "H" */

/* old "l" */

/* old "L" */

/* old "o" */

/* old "O" */

/* old "p" */

/* old "P" */

/* old "r" */

/* old "s" */

/* old "S" */

/* old "z" */

/* old "z" (with dereference) */

/* old "+" and "*" */

/* End of new parameter parsing API */

// Source: <Zend/zend_API.c>

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
   |          Andrei Zmievski <andrei@php.net>                            |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* these variables are true statics/globals, and have to be mutex'ed on every access */

var ModuleRegistry HashTable
var ModuleRequestStartupHandlers **ZendModuleEntry
var ModuleRequestShutdownHandlers **ZendModuleEntry
var ModulePostDeactivateHandlers **ZendModuleEntry
var ClassCleanupHandlers **ZendClassEntry
var DisabledClassNew []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
