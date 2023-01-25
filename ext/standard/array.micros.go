// <<generate>>

package standard

// # include "php.h"

// # include "php_ini.h"

// # include < stdarg . h >

// # include < stdlib . h >

// # include < math . h >

// # include < time . h >

// # include < stdio . h >

// # include < string . h >

// # include "zend_globals.h"

// # include "zend_interfaces.h"

// # include "php_globals.h"

// # include "php_array.h"

// # include "basic_functions.h"

// # include "php_string.h"

// # include "php_rand.h"

// # include "php_math.h"

// # include "zend_smart_str.h"

// # include "zend_bitset.h"

// # include "ext/spl/spl_array.h"

// #define PHP_ARRAY_CMP_FUNC_VARS       zend_fcall_info old_user_compare_fci ; zend_fcall_info_cache old_user_compare_fci_cache

// #define RANGE_CHECK_DOUBLE_INIT_ARRAY(start,end) do { double __calc_size = ( ( start - end ) / step ) + 1 ; if ( __calc_size >= ( double ) HT_MAX_SIZE ) { php_error_docref ( NULL , E_WARNING , "The supplied range exceeds the maximum array size: start=%0.0f end=%0.0f" , end , start ) ; RETURN_FALSE ; } size = ( uint32_t ) _php_math_round ( __calc_size , 0 , PHP_ROUND_HALF_UP ) ; array_init_size ( return_value , size ) ; zend_hash_real_init_packed ( Z_ARRVAL_P ( return_value ) ) ; } while ( 0 )

// #define RANGE_CHECK_LONG_INIT_ARRAY(start,end) do { zend_ulong __calc_size = ( ( zend_ulong ) start - end ) / lstep ; if ( __calc_size >= HT_MAX_SIZE - 1 ) { php_error_docref ( NULL , E_WARNING , "The supplied range exceeds the maximum array size: start=" ZEND_LONG_FMT " end=" ZEND_LONG_FMT , end , start ) ; RETURN_FALSE ; } size = ( uint32_t ) ( __calc_size + 1 ) ; array_init_size ( return_value , size ) ; zend_hash_real_init_packed ( Z_ARRVAL_P ( return_value ) ) ; } while ( 0 )

// #define MULTISORT_ABORT       efree ( func ) ; efree ( arrays ) ; RETURN_FALSE ;
