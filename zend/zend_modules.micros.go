// <<generate>>

package zend

// #define MODULES_H

// # include "zend.h"

// # include "zend_compile.h"

// # include "zend_build.h"

// #define INIT_FUNC_ARGS       int type , int module_number

// #define INIT_FUNC_ARGS_PASSTHRU       type , module_number

// #define SHUTDOWN_FUNC_ARGS       int type , int module_number

// #define SHUTDOWN_FUNC_ARGS_PASSTHRU       type , module_number

// #define STANDARD_MODULE_HEADER_EX       sizeof ( zend_module_entry ) , ZEND_MODULE_API_NO , ZEND_DEBUG , USING_ZTS

// #define STANDARD_MODULE_HEADER       STANDARD_MODULE_HEADER_EX , NULL , NULL

// #define ZE2_STANDARD_MODULE_HEADER       STANDARD_MODULE_HEADER_EX , ini_entries , NULL

// #define ZEND_MODULE_BUILD_ID       "API" ZEND_TOSTR ( ZEND_MODULE_API_NO ) ZEND_BUILD_TS ZEND_BUILD_DEBUG ZEND_BUILD_SYSTEM ZEND_BUILD_EXTRA

// #define STANDARD_MODULE_PROPERTIES_EX       0 , 0 , NULL , 0 , ZEND_MODULE_BUILD_ID

// #define NO_MODULE_GLOBALS       0 , NULL , NULL , NULL

// #define ZEND_MODULE_GLOBALS(module_name) sizeof ( zend_ ## module_name ## _globals ) , & module_name ## _globals

// #define STANDARD_MODULE_PROPERTIES       NO_MODULE_GLOBALS , NULL , STANDARD_MODULE_PROPERTIES_EX

// #define ZEND_MOD_REQUIRED_EX(name,rel,ver) { name , rel , ver , MODULE_DEP_REQUIRED } ,

// #define ZEND_MOD_CONFLICTS_EX(name,rel,ver) { name , rel , ver , MODULE_DEP_CONFLICTS } ,

// #define ZEND_MOD_OPTIONAL_EX(name,rel,ver) { name , rel , ver , MODULE_DEP_OPTIONAL } ,

// #define ZEND_MOD_REQUIRED(name) ZEND_MOD_REQUIRED_EX ( name , NULL , NULL )

// #define ZEND_MOD_CONFLICTS(name) ZEND_MOD_CONFLICTS_EX ( name , NULL , NULL )

// #define ZEND_MOD_OPTIONAL(name) ZEND_MOD_OPTIONAL_EX ( name , NULL , NULL )

// #define ZEND_MOD_END       { NULL , NULL , NULL , 0 }
