// <<generate>>

package standard

// # include "php.h"

// # include "fopen_wrappers.h"

// # include "file.h"

// # include "php_dir.h"

// # include "php_string.h"

// # include "php_scandir.h"

// # include "basic_functions.h"

// # include < unistd . h >

// # include < errno . h >

// # include < glob . h >

// #define FETCH_DIRP() ZEND_PARSE_PARAMETERS_START ( 0 , 1 ) Z_PARAM_OPTIONAL Z_PARAM_RESOURCE ( id ) ZEND_PARSE_PARAMETERS_END ( ) ; if ( ZEND_NUM_ARGS ( ) == 0 ) { myself = getThis ( ) ; if ( myself ) { if ( ( tmp = zend_hash_str_find ( Z_OBJPROP_P ( myself ) , "handle" , sizeof ( "handle" ) - 1 ) ) == NULL ) { php_error_docref ( NULL , E_WARNING , "Unable to find my handle property" ) ; RETURN_FALSE ; } if ( ( dirp = ( php_stream * ) zend_fetch_resource_ex ( tmp , "Directory" , php_file_le_stream ( ) ) ) == NULL ) { RETURN_FALSE ; } } else { if ( ! DIRG ( default_dir ) || ( dirp = ( php_stream * ) zend_fetch_resource ( DIRG ( default_dir ) , "Directory" , php_file_le_stream ( ) ) ) == NULL ) { RETURN_FALSE ; } } } else { if ( ( dirp = ( php_stream * ) zend_fetch_resource ( Z_RES_P ( id ) , "Directory" , php_file_le_stream ( ) ) ) == NULL ) { RETURN_FALSE ; } }
