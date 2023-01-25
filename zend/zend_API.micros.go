// <<generate>>

package zend

// #define ZEND_API_H

// # include "zend_modules.h"

// # include "zend_list.h"

// # include "zend_operators.h"

// # include "zend_variables.h"

// # include "zend_execute.h"

// #define ZEND_FN(name) zif_ ## name

// #define ZEND_MN(name) zim_ ## name

// #define ZEND_NAMED_FUNCTION(name) void ZEND_FASTCALL name ( INTERNAL_FUNCTION_PARAMETERS )

// #define ZEND_FUNCTION(name) ZEND_NAMED_FUNCTION ( ZEND_FN ( name ) )

// #define ZEND_METHOD(classname,name) ZEND_NAMED_FUNCTION ( ZEND_MN ( classname ## _ ## name ) )

// #define ZEND_FENTRY(zend_name,name,arg_info,flags) { # zend_name , name , arg_info , ( uint32_t ) ( sizeof ( arg_info ) / sizeof ( struct _zend_internal_arg_info ) - 1 ) , flags } ,

// #define ZEND_RAW_FENTRY(zend_name,name,arg_info,flags) { zend_name , name , arg_info , ( uint32_t ) ( sizeof ( arg_info ) / sizeof ( struct _zend_internal_arg_info ) - 1 ) , flags } ,

// #define ZEND_RAW_NAMED_FE(zend_name,name,arg_info) ZEND_RAW_FENTRY ( # zend_name , name , arg_info , 0 )

// #define ZEND_NAMED_FE(zend_name,name,arg_info) ZEND_FENTRY ( zend_name , name , arg_info , 0 )

// #define ZEND_FE(name,arg_info) ZEND_FENTRY ( name , ZEND_FN ( name ) , arg_info , 0 )

// #define ZEND_DEP_FE(name,arg_info) ZEND_FENTRY ( name , ZEND_FN ( name ) , arg_info , ZEND_ACC_DEPRECATED )

// #define ZEND_FALIAS(name,alias,arg_info) ZEND_FENTRY ( name , ZEND_FN ( alias ) , arg_info , 0 )

// #define ZEND_DEP_FALIAS(name,alias,arg_info) ZEND_FENTRY ( name , ZEND_FN ( alias ) , arg_info , ZEND_ACC_DEPRECATED )

// #define ZEND_NAMED_ME(zend_name,name,arg_info,flags) ZEND_FENTRY ( zend_name , name , arg_info , flags )

// #define ZEND_ME(classname,name,arg_info,flags) ZEND_FENTRY ( name , ZEND_MN ( classname ## _ ## name ) , arg_info , flags )

// #define ZEND_DEP_ME(classname,name,arg_info,flags) ZEND_ME ( classname , name , arg_info , flags | ZEND_ACC_DEPRECATED )

// #define ZEND_ABSTRACT_ME(classname,name,arg_info) ZEND_FENTRY ( name , NULL , arg_info , ZEND_ACC_PUBLIC | ZEND_ACC_ABSTRACT )

// #define ZEND_MALIAS(classname,name,alias,arg_info,flags) ZEND_FENTRY ( name , ZEND_MN ( classname ## _ ## alias ) , arg_info , flags )

// #define ZEND_ME_MAPPING(name,func_name,arg_types,flags) ZEND_NAMED_ME ( name , ZEND_FN ( func_name ) , arg_types , flags )

// #define ZEND_NS_FENTRY(ns,zend_name,name,arg_info,flags) ZEND_RAW_FENTRY ( ZEND_NS_NAME ( ns , # zend_name ) , name , arg_info , flags )

// #define ZEND_NS_RAW_FENTRY(ns,zend_name,name,arg_info,flags) ZEND_RAW_FENTRY ( ZEND_NS_NAME ( ns , zend_name ) , name , arg_info , flags )

// #define ZEND_NS_RAW_NAMED_FE(ns,zend_name,name,arg_info) ZEND_NS_RAW_FENTRY ( ns , # zend_name , name , arg_info , 0 )

// #define ZEND_NS_NAMED_FE(ns,zend_name,name,arg_info) ZEND_NS_FENTRY ( ns , zend_name , name , arg_info , 0 )

// #define ZEND_NS_FE(ns,name,arg_info) ZEND_NS_FENTRY ( ns , name , ZEND_FN ( name ) , arg_info , 0 )

// #define ZEND_NS_DEP_FE(ns,name,arg_info) ZEND_NS_FENTRY ( ns , name , ZEND_FN ( name ) , arg_info , ZEND_ACC_DEPRECATED )

// #define ZEND_NS_FALIAS(ns,name,alias,arg_info) ZEND_NS_FENTRY ( ns , name , ZEND_FN ( alias ) , arg_info , 0 )

// #define ZEND_NS_DEP_FALIAS(ns,name,alias,arg_info) ZEND_NS_FENTRY ( ns , name , ZEND_FN ( alias ) , arg_info , ZEND_ACC_DEPRECATED )

// #define ZEND_FE_END       { NULL , NULL , NULL , 0 , 0 }

// #define ZEND_ARG_INFO(pass_by_ref,name) { # name , 0 , pass_by_ref , 0 } ,

// #define ZEND_ARG_PASS_INFO(pass_by_ref) { NULL , 0 , pass_by_ref , 0 } ,

// #define ZEND_ARG_OBJ_INFO(pass_by_ref,name,classname,allow_null) { # name , ZEND_TYPE_ENCODE_CLASS_CONST ( # classname , allow_null ) , pass_by_ref , 0 } ,

// #define ZEND_ARG_ARRAY_INFO(pass_by_ref,name,allow_null) { # name , ZEND_TYPE_ENCODE ( IS_ARRAY , allow_null ) , pass_by_ref , 0 } ,

// #define ZEND_ARG_CALLABLE_INFO(pass_by_ref,name,allow_null) { # name , ZEND_TYPE_ENCODE ( IS_CALLABLE , allow_null ) , pass_by_ref , 0 } ,

// #define ZEND_ARG_TYPE_INFO(pass_by_ref,name,type_hint,allow_null) { # name , ZEND_TYPE_ENCODE ( type_hint , allow_null ) , pass_by_ref , 0 } ,

// #define ZEND_ARG_VARIADIC_INFO(pass_by_ref,name) { # name , 0 , pass_by_ref , 1 } ,

// #define ZEND_ARG_VARIADIC_TYPE_INFO(pass_by_ref,name,type_hint,allow_null) { # name , ZEND_TYPE_ENCODE ( type_hint , allow_null ) , pass_by_ref , 1 } ,

// #define ZEND_ARG_VARIADIC_OBJ_INFO(pass_by_ref,name,classname,allow_null) { # name , ZEND_TYPE_ENCODE_CLASS_CONST ( # classname , allow_null ) , pass_by_ref , 1 } ,

// #define ZEND_BEGIN_ARG_WITH_RETURN_OBJ_INFO_EX(name,return_reference,required_num_args,class_name,allow_null) static const zend_internal_arg_info name [ ] = { { ( const char * ) ( zend_uintptr_t ) ( required_num_args ) , ZEND_TYPE_ENCODE_CLASS_CONST ( # class_name , allow_null ) , return_reference , 0 } ,

// #define ZEND_BEGIN_ARG_WITH_RETURN_OBJ_INFO(name,class_name,allow_null) ZEND_BEGIN_ARG_WITH_RETURN_OBJ_INFO_EX ( name , 0 , - 1 , class_name , allow_null )

// #define ZEND_BEGIN_ARG_WITH_RETURN_TYPE_INFO_EX(name,return_reference,required_num_args,type,allow_null) static const zend_internal_arg_info name [ ] = { { ( const char * ) ( zend_uintptr_t ) ( required_num_args ) , ZEND_TYPE_ENCODE ( type , allow_null ) , return_reference , 0 } ,

// #define ZEND_BEGIN_ARG_WITH_RETURN_TYPE_INFO(name,type,allow_null) ZEND_BEGIN_ARG_WITH_RETURN_TYPE_INFO_EX ( name , 0 , - 1 , type , allow_null )

// #define ZEND_BEGIN_ARG_INFO_EX(name,_unused,return_reference,required_num_args) static const zend_internal_arg_info name [ ] = { { ( const char * ) ( zend_uintptr_t ) ( required_num_args ) , 0 , return_reference , 0 } ,

// #define ZEND_BEGIN_ARG_INFO(name,_unused) ZEND_BEGIN_ARG_INFO_EX ( name , 0 , ZEND_RETURN_VALUE , - 1 )

// #define ZEND_END_ARG_INFO() } ;

// #define ZEND_MODULE_STARTUP_N(module) zm_startup_ ## module

// #define ZEND_MODULE_SHUTDOWN_N(module) zm_shutdown_ ## module

// #define ZEND_MODULE_ACTIVATE_N(module) zm_activate_ ## module

// #define ZEND_MODULE_DEACTIVATE_N(module) zm_deactivate_ ## module

// #define ZEND_MODULE_POST_ZEND_DEACTIVATE_N(module) zm_post_zend_deactivate_ ## module

// #define ZEND_MODULE_INFO_N(module) zm_info_ ## module

// #define ZEND_MODULE_GLOBALS_CTOR_N(module) zm_globals_ctor_ ## module

// #define ZEND_MODULE_GLOBALS_DTOR_N(module) zm_globals_dtor_ ## module

// #define ZEND_MODULE_STARTUP_D(module) int ZEND_MODULE_STARTUP_N ( module ) ( INIT_FUNC_ARGS )

// #define ZEND_MODULE_SHUTDOWN_D(module) int ZEND_MODULE_SHUTDOWN_N ( module ) ( SHUTDOWN_FUNC_ARGS )

// #define ZEND_MODULE_ACTIVATE_D(module) int ZEND_MODULE_ACTIVATE_N ( module ) ( INIT_FUNC_ARGS )

// #define ZEND_MODULE_DEACTIVATE_D(module) int ZEND_MODULE_DEACTIVATE_N ( module ) ( SHUTDOWN_FUNC_ARGS )

// #define ZEND_MODULE_INFO_D(module) ZEND_COLD void ZEND_MODULE_INFO_N ( module ) ( ZEND_MODULE_INFO_FUNC_ARGS )

// #define ZEND_MODULE_GLOBALS_CTOR_D(module) void ZEND_MODULE_GLOBALS_CTOR_N ( module ) ( zend_ ## module ## _globals * module ## _globals )

// #define ZEND_MODULE_GLOBALS_DTOR_D(module) void ZEND_MODULE_GLOBALS_DTOR_N ( module ) ( zend_ ## module ## _globals * module ## _globals )

// #define ZEND_GET_MODULE(name) BEGIN_EXTERN_C ( ) ZEND_DLEXPORT zend_module_entry * get_module ( void ) { return & name ## _module_entry ; } END_EXTERN_C ( )

// #define ZEND_BEGIN_MODULE_GLOBALS(module_name) typedef struct _zend_ ## module_name ## _globals {

// #define ZEND_END_MODULE_GLOBALS(module_name) } zend_ ## module_name ## _globals ;

// #define ZEND_DECLARE_MODULE_GLOBALS(module_name) zend_ ## module_name ## _globals module_name ## _globals ;

// #define ZEND_EXTERN_MODULE_GLOBALS(module_name) extern zend_ ## module_name ## _globals module_name ## _globals ;

// #define ZEND_INIT_MODULE_GLOBALS(module_name,globals_ctor,globals_dtor) globals_ctor ( & module_name ## _globals ) ;

// #define ZEND_MODULE_GLOBALS_ACCESSOR(module_name,v) ( module_name ## _globals . v )

// #define ZEND_MODULE_GLOBALS_BULK(module_name) ( & module_name ## _globals )

// #define INIT_CLASS_ENTRY(class_container,class_name,functions) INIT_CLASS_ENTRY_EX ( class_container , class_name , sizeof ( class_name ) - 1 , functions )

// #define INIT_CLASS_ENTRY_EX(class_container,class_name,class_name_len,functions) { memset ( & class_container , 0 , sizeof ( zend_class_entry ) ) ; class_container . name = zend_string_init_interned ( class_name , class_name_len , 1 ) ; class_container . info . internal . builtin_functions = functions ; }

// #define ZEND_WRONG_PARAM_COUNT() { zend_wrong_param_count ( ) ; return ; }

// #define ZEND_WRONG_PARAM_COUNT_WITH_RETVAL(ret) { zend_wrong_param_count ( ) ; return ret ; }

// #define DLEXPORT

// #define CHECK_ZVAL_STRING(z)

// #define CHECK_ZVAL_STRING_REL(z)

// #define RETURN_BOOL(b) { RETVAL_BOOL ( b ) ; return ; }

// #define RETURN_NULL() { RETVAL_NULL ( ) ; return ; }

// #define RETURN_LONG(l) { RETVAL_LONG ( l ) ; return ; }

// #define RETURN_DOUBLE(d) { RETVAL_DOUBLE ( d ) ; return ; }

// #define RETURN_STR(s) { RETVAL_STR ( s ) ; return ; }

// #define RETURN_INTERNED_STR(s) { RETVAL_INTERNED_STR ( s ) ; return ; }

// #define RETURN_NEW_STR(s) { RETVAL_NEW_STR ( s ) ; return ; }

// #define RETURN_STR_COPY(s) { RETVAL_STR_COPY ( s ) ; return ; }

// #define RETURN_STRING(s) { RETVAL_STRING ( s ) ; return ; }

// #define RETURN_STRINGL(s,l) { RETVAL_STRINGL ( s , l ) ; return ; }

// #define RETURN_EMPTY_STRING() { RETVAL_EMPTY_STRING ( ) ; return ; }

// #define RETURN_RES(r) { RETVAL_RES ( r ) ; return ; }

// #define RETURN_ARR(r) { RETVAL_ARR ( r ) ; return ; }

// #define RETURN_EMPTY_ARRAY() { RETVAL_EMPTY_ARRAY ( ) ; return ; }

// #define RETURN_OBJ(r) { RETVAL_OBJ ( r ) ; return ; }

// #define RETURN_ZVAL(zv,copy,dtor) { RETVAL_ZVAL ( zv , copy , dtor ) ; return ; }

// #define RETURN_FALSE       { RETVAL_FALSE ; return ; }

// #define RETURN_TRUE       { RETVAL_TRUE ; return ; }

// #define ZEND_MINIT_FUNCTION       ZEND_MODULE_STARTUP_D

// #define ZEND_GINIT_FUNCTION       ZEND_MODULE_GLOBALS_CTOR_D

// #define Z_EXPECTED_TYPES(_) _ ( Z_EXPECTED_LONG , "int" ) _ ( Z_EXPECTED_BOOL , "bool" ) _ ( Z_EXPECTED_STRING , "string" ) _ ( Z_EXPECTED_ARRAY , "array" ) _ ( Z_EXPECTED_FUNC , "valid callback" ) _ ( Z_EXPECTED_RESOURCE , "resource" ) _ ( Z_EXPECTED_PATH , "a valid path" ) _ ( Z_EXPECTED_OBJECT , "object" ) _ ( Z_EXPECTED_DOUBLE , "float" )

// #define Z_EXPECTED_TYPE_ENUM(id,str) id ,

// #define Z_EXPECTED_TYPE_STR(id,str) str ,

// #define ZEND_PARSE_PARAMETERS_START_EX(flags,min_num_args,max_num_args) do { const int _flags = ( flags ) ; int _min_num_args = ( min_num_args ) ; int _max_num_args = ( max_num_args ) ; int _num_args = EX_NUM_ARGS ( ) ; int _i = 0 ; zval * _real_arg , * _arg = NULL ; zend_expected_type _expected_type = Z_EXPECTED_LONG ; char * _error = NULL ; zend_bool _dummy ; zend_bool _optional = 0 ; int _error_code = ZPP_ERROR_OK ; ( ( void ) _i ) ; ( ( void ) _real_arg ) ; ( ( void ) _arg ) ; ( ( void ) _expected_type ) ; ( ( void ) _error ) ; ( ( void ) _dummy ) ; ( ( void ) _optional ) ; do { if ( UNEXPECTED ( _num_args < _min_num_args ) || ( UNEXPECTED ( _num_args > _max_num_args ) && EXPECTED ( _max_num_args >= 0 ) ) ) { if ( ! ( _flags & ZEND_PARSE_PARAMS_QUIET ) ) { if ( _flags & ZEND_PARSE_PARAMS_THROW ) { zend_wrong_parameters_count_exception ( _min_num_args , _max_num_args ) ; } else { zend_wrong_parameters_count_error ( _min_num_args , _max_num_args ) ; } } _error_code = ZPP_ERROR_FAILURE ; break ; } _real_arg = ZEND_CALL_ARG ( execute_data , 0 ) ;

// #define ZEND_PARSE_PARAMETERS_START(min_num_args,max_num_args) ZEND_PARSE_PARAMETERS_START_EX ( 0 , min_num_args , max_num_args )

// #define ZEND_PARSE_PARAMETERS_NONE() do { if ( UNEXPECTED ( ZEND_NUM_ARGS ( ) != 0 ) ) { zend_wrong_parameters_none_error ( ) ; return ; } } while ( 0 )

// #define ZEND_PARSE_PARAMETERS_END_EX(failure) } while ( 0 ) ; if ( UNEXPECTED ( _error_code != ZPP_ERROR_OK ) ) { if ( ! ( _flags & ZEND_PARSE_PARAMS_QUIET ) ) { if ( _error_code == ZPP_ERROR_WRONG_CALLBACK ) { if ( _flags & ZEND_PARSE_PARAMS_THROW ) { zend_wrong_callback_exception ( _i , _error ) ; } else { zend_wrong_callback_error ( _i , _error ) ; } } else if ( _error_code == ZPP_ERROR_WRONG_CLASS ) { if ( _flags & ZEND_PARSE_PARAMS_THROW ) { zend_wrong_parameter_class_exception ( _i , _error , _arg ) ; } else { zend_wrong_parameter_class_error ( _i , _error , _arg ) ; } } else if ( _error_code == ZPP_ERROR_WRONG_ARG ) { if ( _flags & ZEND_PARSE_PARAMS_THROW ) { zend_wrong_parameter_type_exception ( _i , _expected_type , _arg ) ; } else { zend_wrong_parameter_type_error ( _i , _expected_type , _arg ) ; } } } failure ; } } while ( 0 )

// #define ZEND_PARSE_PARAMETERS_END() ZEND_PARSE_PARAMETERS_END_EX ( return )

// #define Z_PARAM_OPTIONAL       _optional = 1 ;

// #define Z_PARAM_ARRAY_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_array ( _arg , & dest , check_null , 0 ) ) ) { _expected_type = Z_EXPECTED_ARRAY ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_ARRAY_EX(dest,check_null,separate) Z_PARAM_ARRAY_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_ARRAY(dest) Z_PARAM_ARRAY_EX ( dest , 0 , 0 )

// #define Z_PARAM_ARRAY_OR_OBJECT_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_array ( _arg , & dest , check_null , 1 ) ) ) { _expected_type = Z_EXPECTED_ARRAY ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_ARRAY_OR_OBJECT_EX(dest,check_null,separate) Z_PARAM_ARRAY_OR_OBJECT_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_ARRAY_OR_OBJECT(dest) Z_PARAM_ARRAY_OR_OBJECT_EX ( dest , 0 , 0 )

// #define Z_PARAM_BOOL_EX2(dest,is_null,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_bool ( _arg , & dest , & is_null , check_null ) ) ) { _expected_type = Z_EXPECTED_BOOL ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_BOOL_EX(dest,is_null,check_null,separate) Z_PARAM_BOOL_EX2 ( dest , is_null , check_null , separate , separate )

// #define Z_PARAM_BOOL(dest) Z_PARAM_BOOL_EX ( dest , _dummy , 0 , 0 )

// #define Z_PARAM_CLASS_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_class ( _arg , & dest , _i , check_null ) ) ) { _error_code = ZPP_ERROR_FAILURE ; break ; }

// #define Z_PARAM_CLASS_EX(dest,check_null,separate) Z_PARAM_CLASS_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_CLASS(dest) Z_PARAM_CLASS_EX ( dest , 0 , 0 )

// #define Z_PARAM_DOUBLE_EX2(dest,is_null,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_double ( _arg , & dest , & is_null , check_null ) ) ) { _expected_type = Z_EXPECTED_DOUBLE ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_DOUBLE_EX(dest,is_null,check_null,separate) Z_PARAM_DOUBLE_EX2 ( dest , is_null , check_null , separate , separate )

// #define Z_PARAM_DOUBLE(dest) Z_PARAM_DOUBLE_EX ( dest , _dummy , 0 , 0 )

// #define Z_PARAM_FUNC_EX2(dest_fci,dest_fcc,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_func ( _arg , & dest_fci , & dest_fcc , check_null , & _error ) ) ) { if ( ! _error ) { _expected_type = Z_EXPECTED_FUNC ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; } else { _error_code = ZPP_ERROR_WRONG_CALLBACK ; break ; } } else if ( UNEXPECTED ( _error != NULL ) ) { zend_wrong_callback_deprecated ( _i , _error ) ; }

// #define Z_PARAM_FUNC_EX(dest_fci,dest_fcc,check_null,separate) Z_PARAM_FUNC_EX2 ( dest_fci , dest_fcc , check_null , separate , separate )

// #define Z_PARAM_FUNC(dest_fci,dest_fcc) Z_PARAM_FUNC_EX ( dest_fci , dest_fcc , 0 , 0 )

// #define Z_PARAM_ARRAY_HT_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_array_ht ( _arg , & dest , check_null , 0 , separate ) ) ) { _expected_type = Z_EXPECTED_ARRAY ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_ARRAY_HT_EX(dest,check_null,separate) Z_PARAM_ARRAY_HT_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_ARRAY_HT(dest) Z_PARAM_ARRAY_HT_EX ( dest , 0 , 0 )

// #define Z_PARAM_ARRAY_OR_OBJECT_HT_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_array_ht ( _arg , & dest , check_null , 1 , separate ) ) ) { _expected_type = Z_EXPECTED_ARRAY ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_ARRAY_OR_OBJECT_HT_EX(dest,check_null,separate) Z_PARAM_ARRAY_OR_OBJECT_HT_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_ARRAY_OR_OBJECT_HT(dest) Z_PARAM_ARRAY_OR_OBJECT_HT_EX ( dest , 0 , 0 )

// #define Z_PARAM_LONG_EX2(dest,is_null,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_long ( _arg , & dest , & is_null , check_null , 0 ) ) ) { _expected_type = Z_EXPECTED_LONG ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_LONG_EX(dest,is_null,check_null,separate) Z_PARAM_LONG_EX2 ( dest , is_null , check_null , separate , separate )

// #define Z_PARAM_LONG(dest) Z_PARAM_LONG_EX ( dest , _dummy , 0 , 0 )

// #define Z_PARAM_OBJECT_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_object ( _arg , & dest , NULL , check_null ) ) ) { _expected_type = Z_EXPECTED_OBJECT ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_OBJECT_EX(dest,check_null,separate) Z_PARAM_OBJECT_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_OBJECT(dest) Z_PARAM_OBJECT_EX ( dest , 0 , 0 )

// #define Z_PARAM_PATH_EX2(dest,dest_len,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_path ( _arg , & dest , & dest_len , check_null ) ) ) { _expected_type = Z_EXPECTED_PATH ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_PATH_EX(dest,dest_len,check_null,separate) Z_PARAM_PATH_EX2 ( dest , dest_len , check_null , separate , separate )

// #define Z_PARAM_PATH(dest,dest_len) Z_PARAM_PATH_EX ( dest , dest_len , 0 , 0 )

// #define Z_PARAM_PATH_STR_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_path_str ( _arg , & dest , check_null ) ) ) { _expected_type = Z_EXPECTED_PATH ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_PATH_STR_EX(dest,check_null,separate) Z_PARAM_PATH_STR_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_PATH_STR(dest) Z_PARAM_PATH_STR_EX ( dest , 0 , 0 )

// #define Z_PARAM_RESOURCE_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_resource ( _arg , & dest , check_null ) ) ) { _expected_type = Z_EXPECTED_RESOURCE ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_RESOURCE_EX(dest,check_null,separate) Z_PARAM_RESOURCE_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_RESOURCE(dest) Z_PARAM_RESOURCE_EX ( dest , 0 , 0 )

// #define Z_PARAM_STRING_EX2(dest,dest_len,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_string ( _arg , & dest , & dest_len , check_null ) ) ) { _expected_type = Z_EXPECTED_STRING ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_STRING_EX(dest,dest_len,check_null,separate) Z_PARAM_STRING_EX2 ( dest , dest_len , check_null , separate , separate )

// #define Z_PARAM_STRING(dest,dest_len) Z_PARAM_STRING_EX ( dest , dest_len , 0 , 0 )

// #define Z_PARAM_STR_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_str ( _arg , & dest , check_null ) ) ) { _expected_type = Z_EXPECTED_STRING ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_STR_EX(dest,check_null,separate) Z_PARAM_STR_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_STR(dest) Z_PARAM_STR_EX ( dest , 0 , 0 )

// #define Z_PARAM_ZVAL_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; zend_parse_arg_zval_deref ( _arg , & dest , check_null ) ;

// #define Z_PARAM_ZVAL_EX(dest,check_null,separate) Z_PARAM_ZVAL_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_ZVAL(dest) Z_PARAM_ZVAL_EX ( dest , 0 , 0 )

// #define Z_PARAM_VARIADIC_EX(spec,dest,dest_num,post_varargs) do { int _num_varargs = _num_args - _i - ( post_varargs ) ; if ( EXPECTED ( _num_varargs > 0 ) ) { dest = _real_arg + 1 ; dest_num = _num_varargs ; _i += _num_varargs ; _real_arg += _num_varargs ; } else { dest = NULL ; dest_num = 0 ; } } while ( 0 ) ;

// #define Z_PARAM_VARIADIC(spec,dest,dest_num) Z_PARAM_VARIADIC_EX ( spec , dest , dest_num , 0 )

// # include "zend.h"

// # include "zend_execute.h"

// # include "zend_API.h"

// # include "zend_modules.h"

// # include "zend_extensions.h"

// # include "zend_constants.h"

// # include "zend_interfaces.h"

// # include "zend_exceptions.h"

// # include "zend_closures.h"

// # include "zend_inheritance.h"

// # include "zend_ini.h"

// # include < stdarg . h >
