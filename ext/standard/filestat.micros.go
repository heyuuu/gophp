package standard

// # include "php.h"

// # include "fopen_wrappers.h"

// # include "php_globals.h"

// # include < stdlib . h >

// # include < sys / stat . h >

// # include < string . h >

// # include < errno . h >

// # include < ctype . h >

// # include < time . h >

// # include < unistd . h >

// # include < sys / param . h >

// # include < sys / statvfs . h >

// # include < pwd . h >

// # include < grp . h >

// # include < utime . h >

// # include "basic_functions.h"

// # include "php_filestat.h"

// #define FileFunction(name,funcnum) ZEND_NAMED_FUNCTION ( name ) { char * filename ; size_t filename_len ; ZEND_PARSE_PARAMETERS_START ( 1 , 1 ) Z_PARAM_PATH ( filename , filename_len ) ZEND_PARSE_PARAMETERS_END ( ) ; php_stat ( filename , filename_len , funcnum , return_value ) ; }
