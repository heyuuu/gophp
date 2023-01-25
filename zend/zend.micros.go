// <<generate>>

package zend

// #define ZEND_H

// #define ZEND_ENGINE_3

// # include "zend_types.h"

// # include "zend_map_ptr.h"

// # include "zend_errors.h"

// # include "zend_alloc.h"

// # include "zend_llist.h"

// # include "zend_string.h"

// # include "zend_hash.h"

// # include "zend_ast.h"

// # include "zend_gc.h"

// # include "zend_variables.h"

// # include "zend_iterators.h"

// # include "zend_stream.h"

// # include "zend_smart_str_public.h"

// # include "zend_smart_string_public.h"

// # include "zend_signal.h"

// #define INTERNAL_FUNCTION_PARAMETERS       zend_execute_data * execute_data , zval * return_value

// #define INTERNAL_FUNCTION_PARAM_PASSTHRU       execute_data , return_value

// #define ZEND_TSRMLS_CACHE_EXTERN()

// #define ZEND_TSRMLS_CACHE_DEFINE()

// #define ZEND_TSRMLS_CACHE_UPDATE()

// #define ZEND_TSRMLS_CACHE

// #define zend_try       { JMP_BUF * __orig_bailout = EG ( bailout ) ; JMP_BUF __bailout ; EG ( bailout ) = & __bailout ; if ( SETJMP ( __bailout ) == 0 ) {

// #define zend_catch       } else { EG ( bailout ) = __orig_bailout ;

// #define zend_end_try() } EG ( bailout ) = __orig_bailout ; }

// #define zend_first_try       EG ( bailout ) = NULL ; zend_try

// # include "zend_object_handlers.h"

// # include "zend_operators.h"

// # include "zend.h"

// # include "zend_extensions.h"

// # include "zend_modules.h"

// # include "zend_constants.h"

// # include "zend_list.h"

// # include "zend_API.h"

// # include "zend_exceptions.h"

// # include "zend_builtin_functions.h"

// # include "zend_ini.h"

// # include "zend_vm.h"

// # include "zend_dtrace.h"

// # include "zend_virtual_cwd.h"

// # include "zend_smart_str.h"

// # include "zend_smart_string.h"

// # include "zend_cpuinfo.h"
