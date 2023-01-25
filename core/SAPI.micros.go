// <<generate>>

package core

// #define SAPI_H

// # include "php.h"

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_llist.h"

// # include "zend_operators.h"

// # include < sys / stat . h >

// #define SAPI_API

// #define SAPI_POST_READER_FUNC(post_reader) void post_reader ( void )

// #define SAPI_POST_HANDLER_FUNC(post_handler) void post_handler ( char * content_type_dup , void * arg )

// #define SAPI_TREAT_DATA_FUNC(treat_data) void treat_data ( int arg , char * str , zval * destArray )

// #define SAPI_INPUT_FILTER_FUNC(input_filter) unsigned int input_filter ( int arg , char * var , char * * val , size_t val_len , size_t * new_val_len )

// #define STANDARD_SAPI_MODULE_PROPERTIES       NULL , NULL , NULL , NULL , 0 , 0 , NULL , NULL , NULL , NULL , NULL , NULL , 0 , NULL , NULL , NULL

// # include < ctype . h >

// # include < sys / stat . h >

// # include "php.h"

// # include "SAPI.h"

// # include "php_variables.h"

// # include "php_ini.h"

// # include "ext/standard/php_string.h"

// # include "ext/standard/pageinfo.h"

// # include < sys / time . h >

// # include "rfc1867.h"

// # include "php_content_types.h"
