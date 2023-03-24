package spl

// #define SPL_DIRECTORY_H

// # include "php.h"

// # include "php_spl.h"

// # include "php.h"

// # include "php_ini.h"

// # include "ext/standard/info.h"

// # include "ext/standard/file.h"

// # include "ext/standard/php_string.h"

// # include "zend_compile.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_iterators.h"

// # include "spl_directory.h"

// # include "spl_exceptions.h"

// # include "php.h"

// # include "fopen_wrappers.h"

// # include "ext/standard/basic_functions.h"

// # include "ext/standard/php_filestat.h"

// #define FileInfoFunction(func_name,func_num) SPL_METHOD ( SplFileInfo , func_name ) { spl_filesystem_object * intern = Z_SPLFILESYSTEM_P ( ZEND_THIS(executeData) ) ; zend_error_handling error_handling ; if ( zend_parse_parameters_none ( ) == FAILURE ) { return ; } zend_replace_error_handling ( EH_THROW , spl_ce_RuntimeException , & error_handling ) ; spl_filesystem_object_get_file_name ( intern ) ; php_stat ( intern -> file_name , intern -> file_name_len , func_num , return_value ) ; zend_restore_error_handling ( & error_handling ) ; }

// #define FileFunctionCall(func_name,pass_num_args,arg2) { zend_function * func_ptr ; func_ptr = ( zend_function * ) zend_hash_str_find_ptr ( EG ( function_table ) , # func_name , sizeof ( # func_name ) - 1 ) ; if ( func_ptr == NULL ) { zend_throw_exception_ex ( spl_ce_RuntimeException , 0 , "Internal error, function '%s' not found. Please report" , # func_name ) ; return ; } spl_filesystem_file_call ( intern , func_ptr , pass_num_args , return_value , arg2 ) ; }

// #define FileFunction(func_name) SPL_METHOD ( SplFileObject , func_name ) { spl_filesystem_object * intern = Z_SPLFILESYSTEM_P ( ZEND_THIS(executeData) ) ; FileFunctionCall ( func_name , ZEND_NUM_ARGS ( ) , NULL ) ; }
