// <<generate>>

package zend

import g "sik/runtime/grammar"

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

// #define ZEND_API_H

// # include "zend_modules.h"

// # include "zend_list.h"

// # include "zend_operators.h"

// # include "zend_variables.h"

// # include "zend_execute.h"

// @type ZendFunctionEntry struct

// @type ZendFcallInfo struct

// @type ZendFcallInfoCache struct

// #define ZEND_NS_NAME(ns,name) ns "\\" name

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

/* Name macros */

// #define ZEND_MODULE_STARTUP_N(module) zm_startup_ ## module

// #define ZEND_MODULE_SHUTDOWN_N(module) zm_shutdown_ ## module

// #define ZEND_MODULE_ACTIVATE_N(module) zm_activate_ ## module

// #define ZEND_MODULE_DEACTIVATE_N(module) zm_deactivate_ ## module

// #define ZEND_MODULE_POST_ZEND_DEACTIVATE_N(module) zm_post_zend_deactivate_ ## module

// #define ZEND_MODULE_INFO_N(module) zm_info_ ## module

// #define ZEND_MODULE_GLOBALS_CTOR_N(module) zm_globals_ctor_ ## module

// #define ZEND_MODULE_GLOBALS_DTOR_N(module) zm_globals_dtor_ ## module

/* Declaration macros */

// #define ZEND_MODULE_STARTUP_D(module) int ZEND_MODULE_STARTUP_N ( module ) ( INIT_FUNC_ARGS )

// #define ZEND_MODULE_SHUTDOWN_D(module) int ZEND_MODULE_SHUTDOWN_N ( module ) ( SHUTDOWN_FUNC_ARGS )

// #define ZEND_MODULE_ACTIVATE_D(module) int ZEND_MODULE_ACTIVATE_N ( module ) ( INIT_FUNC_ARGS )

// #define ZEND_MODULE_DEACTIVATE_D(module) int ZEND_MODULE_DEACTIVATE_N ( module ) ( SHUTDOWN_FUNC_ARGS )

// #define ZEND_MODULE_POST_ZEND_DEACTIVATE_D(module) int ZEND_MODULE_POST_ZEND_DEACTIVATE_N ( module ) ( void )

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

// #define INIT_CLASS_ENTRY_INIT_METHODS(class_container,functions) { class_container . constructor = NULL ; class_container . destructor = NULL ; class_container . clone = NULL ; class_container . serialize = NULL ; class_container . unserialize = NULL ; class_container . create_object = NULL ; class_container . get_static_method = NULL ; class_container . __call = NULL ; class_container . __callstatic = NULL ; class_container . __tostring = NULL ; class_container . __get = NULL ; class_container . __set = NULL ; class_container . __unset = NULL ; class_container . __isset = NULL ; class_container . __debugInfo = NULL ; class_container . serialize_func = NULL ; class_container . unserialize_func = NULL ; class_container . parent = NULL ; class_container . num_interfaces = 0 ; class_container . trait_names = NULL ; class_container . num_traits = 0 ; class_container . trait_aliases = NULL ; class_container . trait_precedences = NULL ; class_container . interfaces = NULL ; class_container . get_iterator = NULL ; class_container . iterator_funcs_ptr = NULL ; class_container . info . internal . module = NULL ; class_container . info . internal . builtin_functions = functions ; }

// #define INIT_NS_CLASS_ENTRY(class_container,ns,class_name,functions) INIT_CLASS_ENTRY ( class_container , ZEND_NS_NAME ( ns , class_name ) , functions )

// #define CE_STATIC_MEMBERS(ce) ( ( zval * ) ZEND_MAP_PTR_GET ( ( ce ) -> static_members_table ) )

// #define ZEND_FCI_INITIALIZED(fci) ( ( fci ) . size != 0 )

/* internal function to efficiently copy parameters when executing __call() */

// #define zend_get_parameters_array(ht,param_count,argument_array) _zend_get_parameters_array_ex ( param_count , argument_array )

// #define zend_get_parameters_array_ex(param_count,argument_array) _zend_get_parameters_array_ex ( param_count , argument_array )

// #define zend_parse_parameters_none() ( EXPECTED ( ZEND_NUM_ARGS ( ) == 0 ) ? SUCCESS : ( zend_wrong_parameters_none_error ( ) , FAILURE ) )

// #define zend_parse_parameters_none_throw() ( EXPECTED ( ZEND_NUM_ARGS ( ) == 0 ) ? SUCCESS : ( zend_wrong_parameters_none_exception ( ) , FAILURE ) )

/* Parameter parsing API -- andrei */

// #define ZEND_PARSE_PARAMS_QUIET       ( 1 << 1 )

// #define ZEND_PARSE_PARAMS_THROW       ( 1 << 2 )

/* End of parameter parsing API -- andrei */

// #define zend_register_class_alias(name,ce) zend_register_class_alias_ex ( name , sizeof ( name ) - 1 , ce , 1 )

// #define zend_register_ns_class_alias(ns,name,ce) zend_register_class_alias_ex ( ZEND_NS_NAME ( ns , name ) , sizeof ( ZEND_NS_NAME ( ns , name ) ) - 1 , ce , 1 )

// #define IS_CALLABLE_CHECK_SYNTAX_ONLY       ( 1 << 0 )

// #define IS_CALLABLE_CHECK_NO_ACCESS       ( 1 << 1 )

// #define IS_CALLABLE_CHECK_IS_STATIC       ( 1 << 2 )

// #define IS_CALLABLE_CHECK_SILENT       ( 1 << 3 )

// #define IS_CALLABLE_STRICT       ( IS_CALLABLE_CHECK_IS_STATIC )

// #define ZEND_THIS       ( & EX ( This ) )

// #define getThis() ( ( Z_TYPE_P ( ZEND_THIS ) == IS_OBJECT ) ? ZEND_THIS : NULL )

// #define ZEND_IS_METHOD_CALL() ( EX ( func ) -> common . scope != NULL )

// #define WRONG_PARAM_COUNT       ZEND_WRONG_PARAM_COUNT ( )

// #define WRONG_PARAM_COUNT_WITH_RETVAL(ret) ZEND_WRONG_PARAM_COUNT_WITH_RETVAL ( ret )

// #define ARG_COUNT(dummy) EX_NUM_ARGS ( )

// #define ZEND_NUM_ARGS() EX_NUM_ARGS ( )

// #define ZEND_WRONG_PARAM_COUNT() { zend_wrong_param_count ( ) ; return ; }

// #define ZEND_WRONG_PARAM_COUNT_WITH_RETVAL(ret) { zend_wrong_param_count ( ) ; return ret ; }

// #define DLEXPORT

// #define array_init(arg) ZVAL_ARR ( ( arg ) , zend_new_array ( 0 ) )

// #define array_init_size(arg,size) ZVAL_ARR ( ( arg ) , zend_new_array ( size ) )

// #define add_assoc_long(__arg,__key,__n) add_assoc_long_ex ( __arg , __key , strlen ( __key ) , __n )

// #define add_assoc_null(__arg,__key) add_assoc_null_ex ( __arg , __key , strlen ( __key ) )

// #define add_assoc_bool(__arg,__key,__b) add_assoc_bool_ex ( __arg , __key , strlen ( __key ) , __b )

// #define add_assoc_resource(__arg,__key,__r) add_assoc_resource_ex ( __arg , __key , strlen ( __key ) , __r )

// #define add_assoc_double(__arg,__key,__d) add_assoc_double_ex ( __arg , __key , strlen ( __key ) , __d )

// #define add_assoc_str(__arg,__key,__str) add_assoc_str_ex ( __arg , __key , strlen ( __key ) , __str )

// #define add_assoc_string(__arg,__key,__str) add_assoc_string_ex ( __arg , __key , strlen ( __key ) , __str )

// #define add_assoc_stringl(__arg,__key,__str,__length) add_assoc_stringl_ex ( __arg , __key , strlen ( __key ) , __str , __length )

// #define add_assoc_zval(__arg,__key,__value) add_assoc_zval_ex ( __arg , __key , strlen ( __key ) , __value )

func AddIndexZval(arg *Zval, index ZendUlong, value *Zval) int {
	if ZendHashIndexUpdate(arg.GetValue().GetArr(), index, value) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexZval(arg *Zval, value *Zval) int {
	if ZendHashNextIndexInsert(arg.GetValue().GetArr(), value) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}

// #define add_property_long(__arg,__key,__n) add_property_long_ex ( __arg , __key , strlen ( __key ) , __n )

// #define add_property_null(__arg,__key) add_property_null_ex ( __arg , __key , strlen ( __key ) )

// #define add_property_bool(__arg,__key,__b) add_property_bool_ex ( __arg , __key , strlen ( __key ) , __b )

// #define add_property_resource(__arg,__key,__r) add_property_resource_ex ( __arg , __key , strlen ( __key ) , __r )

// #define add_property_double(__arg,__key,__d) add_property_double_ex ( __arg , __key , strlen ( __key ) , __d )

// #define add_property_str(__arg,__key,__str) add_property_str_ex ( __arg , __key , strlen ( __key ) , __str )

// #define add_property_string(__arg,__key,__str) add_property_string_ex ( __arg , __key , strlen ( __key ) , __str )

// #define add_property_stringl(__arg,__key,__str,__length) add_property_stringl_ex ( __arg , __key , strlen ( __key ) , __str , __length )

// #define add_property_zval(__arg,__key,__value) add_property_zval_ex ( __arg , __key , strlen ( __key ) , __value )

// #define call_user_function(function_table,object,function_name,retval_ptr,param_count,params) _call_user_function_ex ( object , function_name , retval_ptr , param_count , params , 1 )

// #define call_user_function_ex(function_table,object,function_name,retval_ptr,param_count,params,no_separation,symbol_table) _call_user_function_ex ( object , function_name , retval_ptr , param_count , params , no_separation )

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

func ZendForbidDynamicCall(func_name string) int {
	var ex *ZendExecuteData = EG.GetCurrentExecuteData()
	assert(ex != nil && ex.GetFunc() != nil)
	if (ex.GetThis().GetTypeInfo() & 1 << 25) != 0 {
		ZendError(1<<1, "Cannot call %s dynamically", func_name)
		return FAILURE
	}
	return SUCCESS
}

// #define CHECK_ZVAL_STRING(z)

// #define CHECK_ZVAL_STRING_REL(z)

// #define CHECK_ZVAL_NULL_PATH(p) ( Z_STRLEN_P ( p ) != strlen ( Z_STRVAL_P ( p ) ) )

// #define CHECK_NULL_PATH(p,l) ( strlen ( p ) != ( size_t ) ( l ) )

// #define ZVAL_STRINGL(z,s,l) do { ZVAL_NEW_STR ( z , zend_string_init ( s , l , 0 ) ) ; } while ( 0 )

// #define ZVAL_STRING(z,s) do { const char * _s = ( s ) ; ZVAL_STRINGL ( z , _s , strlen ( _s ) ) ; } while ( 0 )

// #define ZVAL_EMPTY_STRING(z) do { ZVAL_INTERNED_STR ( z , ZSTR_EMPTY_ALLOC ( ) ) ; } while ( 0 )

// #define ZVAL_PSTRINGL(z,s,l) do { ZVAL_NEW_STR ( z , zend_string_init ( s , l , 1 ) ) ; } while ( 0 )

// #define ZVAL_PSTRING(z,s) do { const char * _s = ( s ) ; ZVAL_PSTRINGL ( z , _s , strlen ( _s ) ) ; } while ( 0 )

// #define ZVAL_EMPTY_PSTRING(z) do { ZVAL_PSTRINGL ( z , "" , 0 ) ; } while ( 0 )

// #define ZVAL_ZVAL(z,zv,copy,dtor) do { zval * __z = ( z ) ; zval * __zv = ( zv ) ; if ( EXPECTED ( ! Z_ISREF_P ( __zv ) ) ) { if ( copy && ! dtor ) { ZVAL_COPY ( __z , __zv ) ; } else { ZVAL_COPY_VALUE ( __z , __zv ) ; } } else { ZVAL_COPY ( __z , Z_REFVAL_P ( __zv ) ) ; if ( dtor || ! copy ) { zval_ptr_dtor ( __zv ) ; } } } while ( 0 )

// #define RETVAL_BOOL(b) ZVAL_BOOL ( return_value , b )

// #define RETVAL_NULL() ZVAL_NULL ( return_value )

// #define RETVAL_LONG(l) ZVAL_LONG ( return_value , l )

// #define RETVAL_DOUBLE(d) ZVAL_DOUBLE ( return_value , d )

// #define RETVAL_STR(s) ZVAL_STR ( return_value , s )

// #define RETVAL_INTERNED_STR(s) ZVAL_INTERNED_STR ( return_value , s )

// #define RETVAL_NEW_STR(s) ZVAL_NEW_STR ( return_value , s )

// #define RETVAL_STR_COPY(s) ZVAL_STR_COPY ( return_value , s )

// #define RETVAL_STRING(s) ZVAL_STRING ( return_value , s )

// #define RETVAL_STRINGL(s,l) ZVAL_STRINGL ( return_value , s , l )

// #define RETVAL_EMPTY_STRING() ZVAL_EMPTY_STRING ( return_value )

// #define RETVAL_RES(r) ZVAL_RES ( return_value , r )

// #define RETVAL_ARR(r) ZVAL_ARR ( return_value , r )

// #define RETVAL_EMPTY_ARRAY() ZVAL_EMPTY_ARRAY ( return_value )

// #define RETVAL_OBJ(r) ZVAL_OBJ ( return_value , r )

// #define RETVAL_ZVAL(zv,copy,dtor) ZVAL_ZVAL ( return_value , zv , copy , dtor )

// #define RETVAL_FALSE       ZVAL_FALSE ( return_value )

// #define RETVAL_TRUE       ZVAL_TRUE ( return_value )

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

// #define HASH_OF(p) ( Z_TYPE_P ( p ) == IS_ARRAY ? Z_ARRVAL_P ( p ) : ( ( Z_TYPE_P ( p ) == IS_OBJECT ? Z_OBJ_HT_P ( p ) -> get_properties ( ( p ) ) : NULL ) ) )

// #define ZVAL_IS_NULL(z) ( Z_TYPE_P ( z ) == IS_NULL )

/* For compatibility */

// #define ZEND_MINIT       ZEND_MODULE_STARTUP_N

// #define ZEND_MSHUTDOWN       ZEND_MODULE_SHUTDOWN_N

// #define ZEND_RINIT       ZEND_MODULE_ACTIVATE_N

// #define ZEND_RSHUTDOWN       ZEND_MODULE_DEACTIVATE_N

// #define ZEND_MINFO       ZEND_MODULE_INFO_N

// #define ZEND_GINIT(module) ( ( void ( * ) ( void * ) ) ( ZEND_MODULE_GLOBALS_CTOR_N ( module ) ) )

// #define ZEND_GSHUTDOWN(module) ( ( void ( * ) ( void * ) ) ( ZEND_MODULE_GLOBALS_DTOR_N ( module ) ) )

// #define ZEND_MINIT_FUNCTION       ZEND_MODULE_STARTUP_D

// #define ZEND_MSHUTDOWN_FUNCTION       ZEND_MODULE_SHUTDOWN_D

// #define ZEND_RINIT_FUNCTION       ZEND_MODULE_ACTIVATE_D

// #define ZEND_RSHUTDOWN_FUNCTION       ZEND_MODULE_DEACTIVATE_D

// #define ZEND_MINFO_FUNCTION       ZEND_MODULE_INFO_D

// #define ZEND_GINIT_FUNCTION       ZEND_MODULE_GLOBALS_CTOR_D

// #define ZEND_GSHUTDOWN_FUNCTION       ZEND_MODULE_GLOBALS_DTOR_D

/* May modify arg in-place. Will free arg in failure case (and take ownership in success case).
 * Prefer using the ZEND_TRY_ASSIGN_* macros over these APIs. */

// #define _ZEND_TRY_ASSIGN_NULL(zv,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_null ( ref ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_NULL ( _zv ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_NULL(zv) _ZEND_TRY_ASSIGN_NULL ( zv , 0 )

// #define ZEND_TRY_ASSIGN_REF_NULL(zv) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_NULL ( zv , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_FALSE(zv,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_bool ( ref , 0 ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_FALSE ( _zv ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_FALSE(zv) _ZEND_TRY_ASSIGN_FALSE ( zv , 0 )

// #define ZEND_TRY_ASSIGN_REF_FALSE(zv) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_FALSE ( zv , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_TRUE(zv,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_bool ( ref , 1 ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_TRUE ( _zv ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_TRUE(zv) _ZEND_TRY_ASSIGN_TRUE ( zv , 0 )

// #define ZEND_TRY_ASSIGN_REF_TRUE(zv) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_TRUE ( zv , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_BOOL(zv,bval,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_bool ( ref , 1 ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_BOOL ( _zv , bval ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_BOOL(zv,bval) _ZEND_TRY_ASSIGN_BOOL ( zv , bval , 0 )

// #define ZEND_TRY_ASSIGN_REF_BOOL(zv,bval) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_BOOL ( zv , bval , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_LONG(zv,lval,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_long ( ref , lval ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_LONG ( _zv , lval ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_LONG(zv,lval) _ZEND_TRY_ASSIGN_LONG ( zv , lval , 0 )

// #define ZEND_TRY_ASSIGN_REF_LONG(zv,lval) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_LONG ( zv , lval , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_DOUBLE(zv,dval,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_double ( ref , dval ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_DOUBLE ( _zv , dval ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_DOUBLE(zv,dval) _ZEND_TRY_ASSIGN_DOUBLE ( zv , dval , 0 )

// #define ZEND_TRY_ASSIGN_REF_DOUBLE(zv,dval) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_DOUBLE ( zv , dval , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_EMPTY_STRING(zv,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_empty_string ( ref ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_EMPTY_STRING ( _zv ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_EMPTY_STRING(zv) _ZEND_TRY_ASSIGN_EMPTY_STRING ( zv , 0 )

// #define ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zv) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_EMPTY_STRING ( zv , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_STR(zv,str,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_str ( ref , str ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_STR ( _zv , str ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_STR(zv,str) _ZEND_TRY_ASSIGN_STR ( zv , str , 0 )

// #define ZEND_TRY_ASSIGN_REF_STR(zv,str) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_STR ( zv , str , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_NEW_STR(zv,str,is_str) do { zval * _zv = zv ; if ( is_str || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_str ( ref , str ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_NEW_STR ( _zv , str ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_NEW_STR(zv,str) _ZEND_TRY_ASSIGN_NEW_STR ( zv , str , 0 )

// #define ZEND_TRY_ASSIGN_REF_NEW_STR(zv,str) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_NEW_STR ( zv , str , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_STRING(zv,string,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_string ( ref , string ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_STRING ( _zv , string ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_STRING(zv,string) _ZEND_TRY_ASSIGN_STRING ( zv , string , 0 )

// #define ZEND_TRY_ASSIGN_REF_STRING(zv,string) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_STRING ( zv , string , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_STRINGL(zv,string,len,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_stringl ( ref , string , len ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_STRINGL ( _zv , string , len ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_STRINGL(zv,string,len) _ZEND_TRY_ASSIGN_STRINGL ( zv , string , len , 0 )

// #define ZEND_TRY_ASSIGN_REF_STRINGL(zv,string,len) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_STRINGL ( zv , string , len , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_ARR(zv,arr,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_arr ( ref , arr ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_ARR ( _zv , arr ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_ARR(zv,arr) _ZEND_TRY_ASSIGN_ARR ( zv , arr , 0 )

// #define ZEND_TRY_ASSIGN_REF_ARR(zv,arr) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_ARR ( zv , arr , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_RES(zv,res,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_res ( ref , res ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_RES ( _zv , res ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_RES(zv,res) _ZEND_TRY_ASSIGN_RES ( zv , res , 0 )

// #define ZEND_TRY_ASSIGN_REF_RES(zv,res) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_RES ( zv , res , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_TMP(zv,other_zv,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref ( ref , other_zv ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_COPY_VALUE ( _zv , other_zv ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_TMP(zv,other_zv) _ZEND_TRY_ASSIGN_TMP ( zv , other_zv , 0 )

// #define ZEND_TRY_ASSIGN_REF_TMP(zv,other_zv) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_TMP ( zv , other_zv , 1 ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_VALUE(zv,other_zv,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_zval ( ref , other_zv ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_COPY_VALUE ( _zv , other_zv ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_VALUE(zv,other_zv) _ZEND_TRY_ASSIGN_VALUE ( zv , other_zv , 0 )

// #define ZEND_TRY_ASSIGN_REF_VALUE(zv,other_zv) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_VALUE ( zv , other_zv , 1 ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_COPY(zv,other_zv) do { Z_TRY_ADDREF_P ( other_zv ) ; ZEND_TRY_ASSIGN_VALUE ( zv , other_zv ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_REF_COPY(zv,other_zv) do { Z_TRY_ADDREF_P ( other_zv ) ; ZEND_TRY_ASSIGN_REF_VALUE ( zv , other_zv ) ; } while ( 0 )

// #define _ZEND_TRY_ASSIGN_VALUE_EX(zv,other_zv,strict,is_ref) do { zval * _zv = zv ; if ( is_ref || UNEXPECTED ( Z_ISREF_P ( _zv ) ) ) { zend_reference * ref = Z_REF_P ( _zv ) ; if ( UNEXPECTED ( ZEND_REF_HAS_TYPE_SOURCES ( ref ) ) ) { zend_try_assign_typed_ref_zval_ex ( ref , other_zv , strict ) ; break ; } _zv = & ref -> val ; } zval_ptr_dtor ( _zv ) ; ZVAL_COPY_VALUE ( _zv , other_zv ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_VALUE_EX(zv,other_zv,strict) _ZEND_TRY_ASSIGN_VALUE_EX ( zv , other_zv , strict , 0 )

// #define ZEND_TRY_ASSIGN_REF_VALUE_EX(zv,other_zv,strict) do { ZEND_ASSERT ( Z_ISREF_P ( zv ) ) ; _ZEND_TRY_ASSIGN_VALUE_EX ( zv , other_zv , strict , 1 ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_COPY_EX(zv,other_zv,strict) do { Z_TRY_ADDREF_P ( other_zv ) ; ZEND_TRY_ASSIGN_VALUE_EX ( zv , other_zv , strict ) ; } while ( 0 )

// #define ZEND_TRY_ASSIGN_REF_COPY_EX(zv,other_zv,strict) do { Z_TRY_ADDREF_P ( other_zv ) ; ZEND_TRY_ASSIGN_REF_VALUE_EX ( zv , other_zv , strict ) ; } while ( 0 )

/* Initializes a reference to an empty array and returns dereferenced zval,
 * or NULL if the initialization failed. */

func ZendTryArrayInitSize(zv *Zval, size uint32) *Zval {
	var arr *ZendArray = _zendNewArray(size)
	if zv.GetType() == 10 {
		var ref *ZendReference = zv.GetValue().GetRef()
		if ref.GetSources().GetPtr() != nil {
			if ZendTryAssignTypedRefArr(ref, arr) != SUCCESS {
				return nil
			}
			return &ref.val
		}
		zv = &ref.val
	}
	ZvalPtrDtor(zv)
	var __arr *ZendArray = arr
	var __z *Zval = zv
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	return zv
}
func ZendTryArrayInit(zv *Zval) *Zval { return ZendTryArrayInitSize(zv, 0) }

/* Fast parameter parsing API */

// #define FAST_ZPP       1

// #define Z_EXPECTED_TYPES(_) _ ( Z_EXPECTED_LONG , "int" ) _ ( Z_EXPECTED_BOOL , "bool" ) _ ( Z_EXPECTED_STRING , "string" ) _ ( Z_EXPECTED_ARRAY , "array" ) _ ( Z_EXPECTED_FUNC , "valid callback" ) _ ( Z_EXPECTED_RESOURCE , "resource" ) _ ( Z_EXPECTED_PATH , "a valid path" ) _ ( Z_EXPECTED_OBJECT , "object" ) _ ( Z_EXPECTED_DOUBLE , "float" )

// #define Z_EXPECTED_TYPE_ENUM(id,str) id ,

// #define Z_EXPECTED_TYPE_STR(id,str) str ,

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

// #define ZPP_ERROR_OK       0

// #define ZPP_ERROR_FAILURE       1

// #define ZPP_ERROR_WRONG_CALLBACK       2

// #define ZPP_ERROR_WRONG_CLASS       3

// #define ZPP_ERROR_WRONG_ARG       4

// #define ZPP_ERROR_WRONG_COUNT       5

// #define ZEND_PARSE_PARAMETERS_START_EX(flags,min_num_args,max_num_args) do { const int _flags = ( flags ) ; int _min_num_args = ( min_num_args ) ; int _max_num_args = ( max_num_args ) ; int _num_args = EX_NUM_ARGS ( ) ; int _i = 0 ; zval * _real_arg , * _arg = NULL ; zend_expected_type _expected_type = Z_EXPECTED_LONG ; char * _error = NULL ; zend_bool _dummy ; zend_bool _optional = 0 ; int _error_code = ZPP_ERROR_OK ; ( ( void ) _i ) ; ( ( void ) _real_arg ) ; ( ( void ) _arg ) ; ( ( void ) _expected_type ) ; ( ( void ) _error ) ; ( ( void ) _dummy ) ; ( ( void ) _optional ) ; do { if ( UNEXPECTED ( _num_args < _min_num_args ) || ( UNEXPECTED ( _num_args > _max_num_args ) && EXPECTED ( _max_num_args >= 0 ) ) ) { if ( ! ( _flags & ZEND_PARSE_PARAMS_QUIET ) ) { if ( _flags & ZEND_PARSE_PARAMS_THROW ) { zend_wrong_parameters_count_exception ( _min_num_args , _max_num_args ) ; } else { zend_wrong_parameters_count_error ( _min_num_args , _max_num_args ) ; } } _error_code = ZPP_ERROR_FAILURE ; break ; } _real_arg = ZEND_CALL_ARG ( execute_data , 0 ) ;

// #define ZEND_PARSE_PARAMETERS_START(min_num_args,max_num_args) ZEND_PARSE_PARAMETERS_START_EX ( 0 , min_num_args , max_num_args )

// #define ZEND_PARSE_PARAMETERS_NONE() do { if ( UNEXPECTED ( ZEND_NUM_ARGS ( ) != 0 ) ) { zend_wrong_parameters_none_error ( ) ; return ; } } while ( 0 )

// #define ZEND_PARSE_PARAMETERS_END_EX(failure) } while ( 0 ) ; if ( UNEXPECTED ( _error_code != ZPP_ERROR_OK ) ) { if ( ! ( _flags & ZEND_PARSE_PARAMS_QUIET ) ) { if ( _error_code == ZPP_ERROR_WRONG_CALLBACK ) { if ( _flags & ZEND_PARSE_PARAMS_THROW ) { zend_wrong_callback_exception ( _i , _error ) ; } else { zend_wrong_callback_error ( _i , _error ) ; } } else if ( _error_code == ZPP_ERROR_WRONG_CLASS ) { if ( _flags & ZEND_PARSE_PARAMS_THROW ) { zend_wrong_parameter_class_exception ( _i , _error , _arg ) ; } else { zend_wrong_parameter_class_error ( _i , _error , _arg ) ; } } else if ( _error_code == ZPP_ERROR_WRONG_ARG ) { if ( _flags & ZEND_PARSE_PARAMS_THROW ) { zend_wrong_parameter_type_exception ( _i , _expected_type , _arg ) ; } else { zend_wrong_parameter_type_error ( _i , _expected_type , _arg ) ; } } } failure ; } } while ( 0 )

// #define ZEND_PARSE_PARAMETERS_END() ZEND_PARSE_PARAMETERS_END_EX ( return )

// #define Z_PARAM_PROLOGUE(deref,separate) ++ _i ; ZEND_ASSERT ( _i <= _min_num_args || _optional == 1 ) ; ZEND_ASSERT ( _i > _min_num_args || _optional == 0 ) ; if ( _optional ) { if ( UNEXPECTED ( _i > _num_args ) ) break ; } _real_arg ++ ; _arg = _real_arg ; if ( deref ) { if ( EXPECTED ( Z_ISREF_P ( _arg ) ) ) { _arg = Z_REFVAL_P ( _arg ) ; } } if ( separate ) { SEPARATE_ZVAL_NOREF ( _arg ) ; }

/* old "|" */

// #define Z_PARAM_OPTIONAL       _optional = 1 ;

/* old "a" */

// #define Z_PARAM_ARRAY_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_array ( _arg , & dest , check_null , 0 ) ) ) { _expected_type = Z_EXPECTED_ARRAY ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_ARRAY_EX(dest,check_null,separate) Z_PARAM_ARRAY_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_ARRAY(dest) Z_PARAM_ARRAY_EX ( dest , 0 , 0 )

/* old "A" */

// #define Z_PARAM_ARRAY_OR_OBJECT_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_array ( _arg , & dest , check_null , 1 ) ) ) { _expected_type = Z_EXPECTED_ARRAY ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_ARRAY_OR_OBJECT_EX(dest,check_null,separate) Z_PARAM_ARRAY_OR_OBJECT_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_ARRAY_OR_OBJECT(dest) Z_PARAM_ARRAY_OR_OBJECT_EX ( dest , 0 , 0 )

/* old "b" */

// #define Z_PARAM_BOOL_EX2(dest,is_null,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_bool ( _arg , & dest , & is_null , check_null ) ) ) { _expected_type = Z_EXPECTED_BOOL ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_BOOL_EX(dest,is_null,check_null,separate) Z_PARAM_BOOL_EX2 ( dest , is_null , check_null , separate , separate )

// #define Z_PARAM_BOOL(dest) Z_PARAM_BOOL_EX ( dest , _dummy , 0 , 0 )

/* old "C" */

// #define Z_PARAM_CLASS_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_class ( _arg , & dest , _i , check_null ) ) ) { _error_code = ZPP_ERROR_FAILURE ; break ; }

// #define Z_PARAM_CLASS_EX(dest,check_null,separate) Z_PARAM_CLASS_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_CLASS(dest) Z_PARAM_CLASS_EX ( dest , 0 , 0 )

/* old "d" */

// #define Z_PARAM_DOUBLE_EX2(dest,is_null,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_double ( _arg , & dest , & is_null , check_null ) ) ) { _expected_type = Z_EXPECTED_DOUBLE ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_DOUBLE_EX(dest,is_null,check_null,separate) Z_PARAM_DOUBLE_EX2 ( dest , is_null , check_null , separate , separate )

// #define Z_PARAM_DOUBLE(dest) Z_PARAM_DOUBLE_EX ( dest , _dummy , 0 , 0 )

/* old "f" */

// #define Z_PARAM_FUNC_EX2(dest_fci,dest_fcc,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_func ( _arg , & dest_fci , & dest_fcc , check_null , & _error ) ) ) { if ( ! _error ) { _expected_type = Z_EXPECTED_FUNC ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; } else { _error_code = ZPP_ERROR_WRONG_CALLBACK ; break ; } } else if ( UNEXPECTED ( _error != NULL ) ) { zend_wrong_callback_deprecated ( _i , _error ) ; }

// #define Z_PARAM_FUNC_EX(dest_fci,dest_fcc,check_null,separate) Z_PARAM_FUNC_EX2 ( dest_fci , dest_fcc , check_null , separate , separate )

// #define Z_PARAM_FUNC(dest_fci,dest_fcc) Z_PARAM_FUNC_EX ( dest_fci , dest_fcc , 0 , 0 )

/* old "h" */

// #define Z_PARAM_ARRAY_HT_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_array_ht ( _arg , & dest , check_null , 0 , separate ) ) ) { _expected_type = Z_EXPECTED_ARRAY ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_ARRAY_HT_EX(dest,check_null,separate) Z_PARAM_ARRAY_HT_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_ARRAY_HT(dest) Z_PARAM_ARRAY_HT_EX ( dest , 0 , 0 )

/* old "H" */

// #define Z_PARAM_ARRAY_OR_OBJECT_HT_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_array_ht ( _arg , & dest , check_null , 1 , separate ) ) ) { _expected_type = Z_EXPECTED_ARRAY ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_ARRAY_OR_OBJECT_HT_EX(dest,check_null,separate) Z_PARAM_ARRAY_OR_OBJECT_HT_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_ARRAY_OR_OBJECT_HT(dest) Z_PARAM_ARRAY_OR_OBJECT_HT_EX ( dest , 0 , 0 )

/* old "l" */

// #define Z_PARAM_LONG_EX2(dest,is_null,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_long ( _arg , & dest , & is_null , check_null , 0 ) ) ) { _expected_type = Z_EXPECTED_LONG ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_LONG_EX(dest,is_null,check_null,separate) Z_PARAM_LONG_EX2 ( dest , is_null , check_null , separate , separate )

// #define Z_PARAM_LONG(dest) Z_PARAM_LONG_EX ( dest , _dummy , 0 , 0 )

/* old "L" */

// #define Z_PARAM_STRICT_LONG_EX2(dest,is_null,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_long ( _arg , & dest , & is_null , check_null , 1 ) ) ) { _expected_type = Z_EXPECTED_LONG ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_STRICT_LONG_EX(dest,is_null,check_null,separate) Z_PARAM_STRICT_LONG_EX2 ( dest , is_null , check_null , separate , separate )

// #define Z_PARAM_STRICT_LONG(dest) Z_PARAM_STRICT_LONG_EX ( dest , _dummy , 0 , 0 )

/* old "o" */

// #define Z_PARAM_OBJECT_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_object ( _arg , & dest , NULL , check_null ) ) ) { _expected_type = Z_EXPECTED_OBJECT ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_OBJECT_EX(dest,check_null,separate) Z_PARAM_OBJECT_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_OBJECT(dest) Z_PARAM_OBJECT_EX ( dest , 0 , 0 )

/* old "O" */

// #define Z_PARAM_OBJECT_OF_CLASS_EX2(dest,_ce,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_object ( _arg , & dest , _ce , check_null ) ) ) { if ( _ce ) { _error = ZSTR_VAL ( ( _ce ) -> name ) ; _error_code = ZPP_ERROR_WRONG_CLASS ; break ; } else { _expected_type = Z_EXPECTED_OBJECT ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; } }

// #define Z_PARAM_OBJECT_OF_CLASS_EX(dest,_ce,check_null,separate) Z_PARAM_OBJECT_OF_CLASS_EX2 ( dest , _ce , check_null , separate , separate )

// #define Z_PARAM_OBJECT_OF_CLASS(dest,_ce) Z_PARAM_OBJECT_OF_CLASS_EX ( dest , _ce , 0 , 0 )

/* old "p" */

// #define Z_PARAM_PATH_EX2(dest,dest_len,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_path ( _arg , & dest , & dest_len , check_null ) ) ) { _expected_type = Z_EXPECTED_PATH ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_PATH_EX(dest,dest_len,check_null,separate) Z_PARAM_PATH_EX2 ( dest , dest_len , check_null , separate , separate )

// #define Z_PARAM_PATH(dest,dest_len) Z_PARAM_PATH_EX ( dest , dest_len , 0 , 0 )

/* old "P" */

// #define Z_PARAM_PATH_STR_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_path_str ( _arg , & dest , check_null ) ) ) { _expected_type = Z_EXPECTED_PATH ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_PATH_STR_EX(dest,check_null,separate) Z_PARAM_PATH_STR_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_PATH_STR(dest) Z_PARAM_PATH_STR_EX ( dest , 0 , 0 )

/* old "r" */

// #define Z_PARAM_RESOURCE_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_resource ( _arg , & dest , check_null ) ) ) { _expected_type = Z_EXPECTED_RESOURCE ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_RESOURCE_EX(dest,check_null,separate) Z_PARAM_RESOURCE_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_RESOURCE(dest) Z_PARAM_RESOURCE_EX ( dest , 0 , 0 )

/* old "s" */

// #define Z_PARAM_STRING_EX2(dest,dest_len,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_string ( _arg , & dest , & dest_len , check_null ) ) ) { _expected_type = Z_EXPECTED_STRING ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_STRING_EX(dest,dest_len,check_null,separate) Z_PARAM_STRING_EX2 ( dest , dest_len , check_null , separate , separate )

// #define Z_PARAM_STRING(dest,dest_len) Z_PARAM_STRING_EX ( dest , dest_len , 0 , 0 )

/* old "S" */

// #define Z_PARAM_STR_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; if ( UNEXPECTED ( ! zend_parse_arg_str ( _arg , & dest , check_null ) ) ) { _expected_type = Z_EXPECTED_STRING ; _error_code = ZPP_ERROR_WRONG_ARG ; break ; }

// #define Z_PARAM_STR_EX(dest,check_null,separate) Z_PARAM_STR_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_STR(dest) Z_PARAM_STR_EX ( dest , 0 , 0 )

/* old "z" */

// #define Z_PARAM_ZVAL_EX2(dest,check_null,deref,separate) Z_PARAM_PROLOGUE ( deref , separate ) ; zend_parse_arg_zval_deref ( _arg , & dest , check_null ) ;

// #define Z_PARAM_ZVAL_EX(dest,check_null,separate) Z_PARAM_ZVAL_EX2 ( dest , check_null , separate , separate )

// #define Z_PARAM_ZVAL(dest) Z_PARAM_ZVAL_EX ( dest , 0 , 0 )

/* old "z" (with dereference) */

// #define Z_PARAM_ZVAL_DEREF_EX(dest,check_null,separate) Z_PARAM_PROLOGUE ( 1 , separate ) ; zend_parse_arg_zval_deref ( _arg , & dest , check_null ) ;

// #define Z_PARAM_ZVAL_DEREF(dest) Z_PARAM_ZVAL_DEREF_EX ( dest , 0 , 0 )

/* old "+" and "*" */

// #define Z_PARAM_VARIADIC_EX(spec,dest,dest_num,post_varargs) do { int _num_varargs = _num_args - _i - ( post_varargs ) ; if ( EXPECTED ( _num_varargs > 0 ) ) { dest = _real_arg + 1 ; dest_num = _num_varargs ; _i += _num_varargs ; _real_arg += _num_varargs ; } else { dest = NULL ; dest_num = 0 ; } } while ( 0 ) ;

// #define Z_PARAM_VARIADIC(spec,dest,dest_num) Z_PARAM_VARIADIC_EX ( spec , dest , dest_num , 0 )

/* End of new parameter parsing API */

func ZendParseArgBool(arg *Zval, dest *ZendBool, is_null *ZendBool, check_null int) int {
	if check_null != 0 {
		*is_null = 0
	}
	if arg.GetType() == 3 {
		*dest = 1
	} else if arg.GetType() == 2 {
		*dest = 0
	} else if check_null != 0 && arg.GetType() == 1 {
		*is_null = 1
		*dest = 0
	} else {
		return ZendParseArgBoolSlow(arg, dest)
	}
	return 1
}
func ZendParseArgLong(arg *Zval, dest *ZendLong, is_null *ZendBool, check_null int, cap int) int {
	if check_null != 0 {
		*is_null = 0
	}
	if arg.GetType() == 4 {
		*dest = arg.GetValue().GetLval()
	} else if check_null != 0 && arg.GetType() == 1 {
		*is_null = 1
		*dest = 0
	} else if cap != 0 {
		return ZendParseArgLongCapSlow(arg, dest)
	} else {
		return ZendParseArgLongSlow(arg, dest)
	}
	return 1
}
func ZendParseArgDouble(arg *Zval, dest *float64, is_null *ZendBool, check_null int) int {
	if check_null != 0 {
		*is_null = 0
	}
	if arg.GetType() == 5 {
		*dest = arg.GetValue().GetDval()
	} else if check_null != 0 && arg.GetType() == 1 {
		*is_null = 1
		*dest = 0.0
	} else {
		return ZendParseArgDoubleSlow(arg, dest)
	}
	return 1
}
func ZendParseArgStr(arg *Zval, dest **ZendString, check_null int) int {
	if arg.GetType() == 6 {
		*dest = arg.GetValue().GetStr()
	} else if check_null != 0 && arg.GetType() == 1 {
		*dest = nil
	} else {
		return ZendParseArgStrSlow(arg, dest)
	}
	return 1
}
func ZendParseArgString(arg *Zval, dest **byte, dest_len *int, check_null int) int {
	var str *ZendString
	if ZendParseArgStr(arg, &str, check_null) == 0 {
		return 0
	}
	if check_null != 0 && str == nil {
		*dest = nil
		*dest_len = 0
	} else {
		*dest = str.GetVal()
		*dest_len = str.GetLen()
	}
	return 1
}
func ZendParseArgPathStr(arg *Zval, dest **ZendString, check_null int) int {
	if ZendParseArgStr(arg, dest, check_null) == 0 || (*dest) != nil && strlen((*dest).GetVal()) != size_t((*dest).GetLen()) {
		return 0
	}
	return 1
}
func ZendParseArgPath(arg *Zval, dest **byte, dest_len *int, check_null int) int {
	var str *ZendString
	if ZendParseArgPathStr(arg, &str, check_null) == 0 {
		return 0
	}
	if check_null != 0 && str == nil {
		*dest = nil
		*dest_len = 0
	} else {
		*dest = str.GetVal()
		*dest_len = str.GetLen()
	}
	return 1
}
func ZendParseArgArray(arg *Zval, dest **Zval, check_null int, or_object int) int {
	if arg.GetType() == 7 || or_object != 0 && arg.GetType() == 8 {
		*dest = arg
	} else if check_null != 0 && arg.GetType() == 1 {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgArrayHt(arg *Zval, dest **HashTable, check_null int, or_object int, separate int) int {
	if arg.GetType() == 7 {
		*dest = arg.GetValue().GetArr()
	} else if or_object != 0 && arg.GetType() == 8 {
		if separate != 0 && arg.GetValue().GetObj().GetProperties() != nil && ZendGcRefcount(&(arg.GetValue().GetObj().GetProperties()).gc) > 1 {
			if (ZvalGcFlags(arg.GetValue().GetObj().GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
				ZendGcDelref(&(arg.GetValue().GetObj().GetProperties()).gc)
			}
			arg.GetValue().GetObj().SetProperties(ZendArrayDup(arg.GetValue().GetObj().GetProperties()))
		}
		*dest = arg.GetValue().GetObj().GetHandlers().GetGetProperties()(arg)
	} else if check_null != 0 && arg.GetType() == 1 {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgObject(arg *Zval, dest **Zval, ce *ZendClassEntry, check_null int) int {
	if arg.GetType() == 8 && (ce == nil || InstanceofFunction(arg.GetValue().GetObj().GetCe(), ce) != 0) {
		*dest = arg
	} else if check_null != 0 && arg.GetType() == 1 {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgResource(arg *Zval, dest **Zval, check_null int) int {
	if arg.GetType() == 9 {
		*dest = arg
	} else if check_null != 0 && arg.GetType() == 1 {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgFunc(arg *Zval, dest_fci *ZendFcallInfo, dest_fcc *ZendFcallInfoCache, check_null int, error **byte) int {
	if check_null != 0 && arg.GetType() == 1 {
		dest_fci.SetSize(0)
		dest_fcc.SetFunctionHandler(nil)
		*error = nil
	} else if ZendFcallInfoInit(arg, 0, dest_fci, dest_fcc, nil, error) != SUCCESS {
		return 0
	}
	return 1
}
func ZendParseArgZval(arg *Zval, dest **Zval, check_null int) {
	if check_null != 0 && (arg.GetType() == 1 || arg.GetType() == 10 && &(*arg).value.GetRef().GetVal().u1.v.type_ == 1) {
		*dest = nil
	} else {
		*dest = arg
	}
}
func ZendParseArgZvalDeref(arg *Zval, dest **Zval, check_null int) {
	if check_null != 0 && arg.GetType() == 1 {
		*dest = nil
	} else {
		*dest = arg
	}
}

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

/* these variables are true statics/globals, and have to be mutex'ed on every access */

var ModuleRegistry HashTable
var ModuleRequestStartupHandlers **ZendModuleEntry
var ModuleRequestShutdownHandlers **ZendModuleEntry
var ModulePostDeactivateHandlers **ZendModuleEntry
var ClassCleanupHandlers **ZendClassEntry

func _zendGetParametersArrayEx(param_count int, argument_array *Zval) int {
	var param_ptr *Zval
	var arg_count int
	param_ptr = (*Zval)(EG.GetCurrentExecuteData()) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
	arg_count = EG.GetCurrentExecuteData().GetThis().GetNumArgs()
	if param_count > arg_count {
		return FAILURE
	}
	for g.PostDec(&param_count) > 0 {
		var _z1 *Zval = argument_array
		var _z2 *Zval = param_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		argument_array++
		param_ptr++
	}
	return SUCCESS
}

/* }}} */

func ZendCopyParametersArray(param_count int, argument_array *Zval) int {
	var param_ptr *Zval
	var arg_count int
	param_ptr = (*Zval)(EG.GetCurrentExecuteData()) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(1)-1))
	arg_count = EG.GetCurrentExecuteData().GetThis().GetNumArgs()
	if param_count > arg_count {
		return FAILURE
	}
	for g.PostDec(&param_count) > 0 {
		if param_ptr.GetTypeFlags() != 0 {
			ZvalAddrefP(param_ptr)
		}
		ZendHashNextIndexInsertNew(argument_array.GetValue().GetArr(), param_ptr)
		param_ptr++
	}
	return SUCCESS
}

/* }}} */

func ZendWrongParamCount() {
	var space *byte
	var class_name *byte = GetActiveClassName(&space)
	ZendInternalArgumentCountError(EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0, "Wrong parameter count for %s%s%s()", class_name, space, GetActiveFunctionName())
}

/* }}} */

func ZendGetTypeByConst(type_ int) *byte {
	switch type_ {
	case 2:

	case 3:

	case 16:
		return "bool"
	case 4:
		return "int"
	case 5:
		return "float"
	case 6:
		return "string"
	case 8:
		return "object"
	case 9:
		return "resource"
	case 1:
		return "null"
	case 17:
		return "callable"
	case 18:
		return "iterable"
	case 7:
		return "array"
	case 19:
		return "void"
	case 20:
		return "number"
	default:
		return "unknown"
	}
}

/* }}} */

func ZendZvalTypeName(arg *Zval) *byte {
	if arg.GetType() == 10 {
		arg = &(*arg).value.GetRef().GetVal()
	}
	return ZendGetTypeByConst(arg.GetType())
}

/* }}} */

func ZendZvalGetType(arg *Zval) *ZendString {
	switch arg.GetType() {
	case 1:
		return ZendKnownStrings[ZEND_STR_NULL]
	case 2:

	case 3:
		return ZendKnownStrings[ZEND_STR_BOOLEAN]
	case 4:
		return ZendKnownStrings[ZEND_STR_INTEGER]
	case 5:
		return ZendKnownStrings[ZEND_STR_DOUBLE]
	case 6:
		return ZendKnownStrings[ZEND_STR_STRING]
	case 7:
		return ZendKnownStrings[ZEND_STR_ARRAY]
	case 8:
		return ZendKnownStrings[ZEND_STR_OBJECT]
	case 9:
		if ZendRsrcListGetRsrcType(arg.GetValue().GetRes()) != nil {
			return ZendKnownStrings[ZEND_STR_RESOURCE]
		} else {
			return ZendKnownStrings[ZEND_STR_CLOSED_RESOURCE]
		}
	default:
		return nil
	}
}

/* }}} */

func ZendWrongParametersNoneError() int {
	var num_args int = EG.GetCurrentExecuteData().GetThis().GetNumArgs()
	var active_function *ZendFunction = EG.GetCurrentExecuteData().GetFunc()
	var class_name *byte = g.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendInternalArgumentCountError(EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0, "%s%s%s() expects %s %d parameter%s, %d given", class_name, g.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), "exactly", 0, "s", num_args)
	return FAILURE
}

/* }}} */

func ZendWrongParametersNoneException() int {
	var num_args int = EG.GetCurrentExecuteData().GetThis().GetNumArgs()
	var active_function *ZendFunction = EG.GetCurrentExecuteData().GetFunc()
	var class_name *byte = g.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendInternalArgumentCountError(1, "%s%s%s() expects %s %d parameter%s, %d given", class_name, g.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), "exactly", 0, "s", num_args)
	return FAILURE
}

/* }}} */

func ZendWrongParametersCountError(min_num_args int, max_num_args int) {
	var num_args int = EG.GetCurrentExecuteData().GetThis().GetNumArgs()
	var active_function *ZendFunction = EG.GetCurrentExecuteData().GetFunc()
	var class_name *byte = g.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendInternalArgumentCountError(EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0, "%s%s%s() expects %s %d parameter%s, %d given", class_name, g.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), g.Cond(g.Cond(min_num_args == max_num_args, "exactly", num_args < min_num_args), "at least", "at most"), g.Cond(num_args < min_num_args, min_num_args, max_num_args), g.Cond(g.Cond(num_args < min_num_args, min_num_args, max_num_args) == 1, "", "s"), num_args)
}

/* }}} */

func ZendWrongParametersCountException(min_num_args int, max_num_args int) {
	var num_args int = EG.GetCurrentExecuteData().GetThis().GetNumArgs()
	var active_function *ZendFunction = EG.GetCurrentExecuteData().GetFunc()
	var class_name *byte = g.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendInternalArgumentCountError(1, "%s%s%s() expects %s %d parameter%s, %d given", class_name, g.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), g.Cond(g.Cond(min_num_args == max_num_args, "exactly", num_args < min_num_args), "at least", "at most"), g.Cond(num_args < min_num_args, min_num_args, max_num_args), g.Cond(g.Cond(num_args < min_num_args, min_num_args, max_num_args) == 1, "", "s"), num_args)
}

/* }}} */

func ZendWrongParameterTypeError(num int, expected_type ZendExpectedType, arg *Zval) {
	var space *byte
	var class_name *byte
	var expected_error []*byte = []*byte{"int", "bool", "string", "array", "valid callback", "resource", "a valid path", "object", "float", nil}
	if EG.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0, "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), num, expected_error[expected_type], ZendZvalTypeName(arg))
}

/* }}} */

func ZendWrongParameterTypeException(num int, expected_type ZendExpectedType, arg *Zval) {
	var space *byte
	var class_name *byte
	var expected_error []*byte = []*byte{"int", "bool", "string", "array", "valid callback", "resource", "a valid path", "object", "float", nil}
	if EG.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(1, "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), num, expected_error[expected_type], ZendZvalTypeName(arg))
}

/* }}} */

func ZendWrongParameterClassError(num int, name *byte, arg *Zval) {
	var space *byte
	var class_name *byte
	if EG.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0, "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), num, name, ZendZvalTypeName(arg))
}

/* }}} */

func ZendWrongParameterClassException(num int, name *byte, arg *Zval) {
	var space *byte
	var class_name *byte
	if EG.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(1, "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), num, name, ZendZvalTypeName(arg))
}

/* }}} */

func ZendWrongCallbackError(num int, error *byte) {
	var space *byte
	var class_name *byte
	if EG.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0, "%s%s%s() expects parameter %d to be a valid callback, %s", class_name, space, GetActiveFunctionName(), num, error)
	_efree(error)
}

/* }}} */

func ZendWrongCallbackException(num int, error *byte) {
	var space *byte
	var class_name *byte
	if EG.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(1, "%s%s%s() expects parameter %d to be a valid callback, %s", class_name, space, GetActiveFunctionName(), num, error)
	_efree(error)
}

/* }}} */

func ZendWrongCallbackDeprecated(num int, error *byte) {
	var space *byte
	var class_name *byte = GetActiveClassName(&space)
	ZendError(1<<13, "%s%s%s() expects parameter %d to be a valid callback, %s", class_name, space, GetActiveFunctionName(), num, error)
	_efree(error)
}

/* }}} */

func ZendParseArgClass(arg *Zval, pce **ZendClassEntry, num int, check_null int) int {
	var ce_base *ZendClassEntry = *pce
	if check_null != 0 && arg.GetType() == 1 {
		*pce = nil
		return 1
	}
	if TryConvertToString(arg) == 0 {
		*pce = nil
		return 0
	}
	*pce = ZendLookupClass(arg.GetValue().GetStr())
	if ce_base != nil {
		if (*pce) == nil || InstanceofFunction(*pce, ce_base) == 0 {
			var space *byte
			var class_name *byte = GetActiveClassName(&space)
			ZendInternalTypeError(EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0, "%s%s%s() expects parameter %d to be a class name derived from %s, '%s' given", class_name, space, GetActiveFunctionName(), num, ce_base.GetName().GetVal(), arg.GetValue().GetStr().GetVal())
			*pce = nil
			return 0
		}
	}
	if (*pce) == nil {
		var space *byte
		var class_name *byte = GetActiveClassName(&space)
		ZendInternalTypeError(EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0, "%s%s%s() expects parameter %d to be a valid class name, '%s' given", class_name, space, GetActiveFunctionName(), num, arg.GetValue().GetStr().GetVal())
		return 0
	}
	return 1
}

/* }}} */

func ZendParseArgBoolWeak(arg *Zval, dest *ZendBool) int {
	if arg.GetType() <= 6 {
		*dest = ZendIsTrue(arg)
	} else {
		return 0
	}
	return 1
}

/* }}} */

func ZendParseArgBoolSlow(arg *Zval, dest *ZendBool) int {
	if EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0 {
		return 0
	}
	return ZendParseArgBoolWeak(arg, dest)
}

/* }}} */

func ZendParseArgLongWeak(arg *Zval, dest *ZendLong) int {
	if arg.GetType() == 5 {
		if isnan(arg.GetValue().GetDval()) {
			return 0
		}
		if arg.GetValue().GetDval() >= float64(INT64_MAX || arg.GetValue().GetDval() < float64(INT64_MIN)) {
			return 0
		} else {
			*dest = ZendDvalToLval(arg.GetValue().GetDval())
		}
	} else if arg.GetType() == 6 {
		var d float64
		var type_ int
		if g.Assign(&type_, IsNumericStrFunction(arg.GetValue().GetStr(), dest, &d)) != 4 {
			if type_ != 0 {
				if isnan(d) {
					return 0
				}
				if d >= float64(INT64_MAX || d < float64(INT64_MIN)) {
					return 0
				} else {
					*dest = ZendDvalToLval(d)
				}
			} else {
				return 0
			}
		}
		if EG.GetException() != nil {
			return 0
		}
	} else if arg.GetType() < 3 {
		*dest = 0
	} else if arg.GetType() == 3 {
		*dest = 1
	} else {
		return 0
	}
	return 1
}

/* }}} */

func ZendParseArgLongSlow(arg *Zval, dest *ZendLong) int {
	if EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0 {
		return 0
	}
	return ZendParseArgLongWeak(arg, dest)
}

/* }}} */

func ZendParseArgLongCapWeak(arg *Zval, dest *ZendLong) int {
	if arg.GetType() == 5 {
		if isnan(arg.GetValue().GetDval()) {
			return 0
		}
		*dest = ZendDvalToLvalCap(arg.GetValue().GetDval())
	} else if arg.GetType() == 6 {
		var d float64
		var type_ int
		if g.Assign(&type_, IsNumericStrFunction(arg.GetValue().GetStr(), dest, &d)) != 4 {
			if type_ != 0 {
				if isnan(d) {
					return 0
				}
				*dest = ZendDvalToLvalCap(d)
			} else {
				return 0
			}
		}
		if EG.GetException() != nil {
			return 0
		}
	} else if arg.GetType() < 3 {
		*dest = 0
	} else if arg.GetType() == 3 {
		*dest = 1
	} else {
		return 0
	}
	return 1
}

/* }}} */

func ZendParseArgLongCapSlow(arg *Zval, dest *ZendLong) int {
	if EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0 {
		return 0
	}
	return ZendParseArgLongCapWeak(arg, dest)
}

/* }}} */

func ZendParseArgDoubleWeak(arg *Zval, dest *float64) int {
	if arg.GetType() == 4 {
		*dest = float64(arg.GetValue().GetLval())
	} else if arg.GetType() == 6 {
		var l ZendLong
		var type_ int
		if g.Assign(&type_, IsNumericStrFunction(arg.GetValue().GetStr(), &l, dest)) != 5 {
			if type_ != 0 {
				*dest = float64(l)
			} else {
				return 0
			}
		}
		if EG.GetException() != nil {
			return 0
		}
	} else if arg.GetType() < 3 {
		*dest = 0.0
	} else if arg.GetType() == 3 {
		*dest = 1.0
	} else {
		return 0
	}
	return 1
}

/* }}} */

func ZendParseArgDoubleSlow(arg *Zval, dest *float64) int {
	if arg.GetType() == 4 {

		/* SSTH Exception: IS_LONG may be accepted instead as IS_DOUBLE */

		*dest = float64(arg.GetValue().GetLval())

		/* SSTH Exception: IS_LONG may be accepted instead as IS_DOUBLE */

	} else if EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0 {
		return 0
	}
	return ZendParseArgDoubleWeak(arg, dest)
}

/* }}} */

func ZendParseArgStrWeak(arg *Zval, dest **ZendString) int {
	if arg.GetType() < 6 {
		if arg.GetType() != 6 {
			_convertToString(arg)
		}
		*dest = arg.GetValue().GetStr()
	} else if arg.GetType() == 8 {
		if arg.GetValue().GetObj().GetHandlers().GetCastObject() != nil {
			var obj Zval
			if arg.GetValue().GetObj().GetHandlers().GetCastObject()(arg, &obj, 6) == SUCCESS {
				ZvalPtrDtor(arg)
				var _z1 *Zval = arg
				var _z2 *Zval = &obj
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				*dest = arg.GetValue().GetStr()
				return 1
			}
		} else if arg.GetValue().GetObj().GetHandlers().GetGet() != nil {
			var rv Zval
			var z *Zval = arg.GetValue().GetObj().GetHandlers().GetGet()(arg, &rv)
			if z.GetType() != 8 {
				ZvalPtrDtor(arg)
				if z.GetType() == 6 {
					var _z1 *Zval = arg
					var _z2 *Zval = z
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
				} else {
					var __z *Zval = arg
					var __s *ZendString = ZvalGetStringFunc(z)
					__z.GetValue().SetStr(__s)
					if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
						__z.SetTypeInfo(6)
					} else {
						__z.SetTypeInfo(6 | 1<<0<<8)
					}
					ZvalPtrDtor(z)
				}
				*dest = arg.GetValue().GetStr()
				return 1
			}
			ZvalPtrDtor(z)
		}
		return 0
	} else {
		return 0
	}
	return 1
}

/* }}} */

func ZendParseArgStrSlow(arg *Zval, dest **ZendString) int {
	if EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0 {
		return 0
	}
	return ZendParseArgStrWeak(arg, dest)
}

/* }}} */

func ZendParseArgImpl(arg_num int, arg *Zval, va *va_list, spec **byte, error **byte, severity *int) *byte {
	var spec_walk *byte = *spec
	var c byte = g.PostInc(&(*spec_walk))
	var check_null int = 0
	var separate int = 0
	var real_arg *Zval = arg

	/* scan through modifiers */

	if arg.GetType() == 10 {
		arg = &(*arg).value.GetRef().GetVal()
	}
	for true {
		if (*spec_walk) == '/' {
			var _zv *Zval = arg
			assert(_zv.GetType() != 10)
			var __zv *Zval = _zv
			if __zv.GetType() == 7 {
				if ZvalRefcountP(__zv) > 1 {
					if __zv.GetTypeFlags() != 0 {
						ZvalDelrefP(__zv)
					}
					var __arr *ZendArray = ZendArrayDup(__zv.GetValue().GetArr())
					var __z *Zval = __zv
					__z.GetValue().SetArr(__arr)
					__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				}
			}
			real_arg = arg
			separate = 1
		} else if (*spec_walk) == '!' {
			check_null = 1
		} else {
			break
		}
		spec_walk++
	}
	switch c {
	case 'l':

	case 'L':
		var p *ZendLong = __va_arg(*va, (*ZendLong)(_))
		var is_null *ZendBool = nil
		if check_null != 0 {
			is_null = __va_arg(*va, (*ZendBool)(_))
		}
		if ZendParseArgLong(arg, p, is_null, check_null, c == 'L') == 0 {
			return "int"
		}
		break
	case 'd':
		var p *float64 = __va_arg(*va, (*float64)(_))
		var is_null *ZendBool = nil
		if check_null != 0 {
			is_null = __va_arg(*va, (*ZendBool)(_))
		}
		if ZendParseArgDouble(arg, p, is_null, check_null) == 0 {
			return "float"
		}
		break
	case 's':
		var p **byte = __va_arg(*va, (**byte)(_))
		var pl *int = __va_arg(*va, (*int)(_))
		if ZendParseArgString(arg, p, pl, check_null) == 0 {
			return "string"
		}
		break
	case 'p':
		var p **byte = __va_arg(*va, (**byte)(_))
		var pl *int = __va_arg(*va, (*int)(_))
		if ZendParseArgPath(arg, p, pl, check_null) == 0 {
			return "a valid path"
		}
		break
	case 'P':
		var str **ZendString = __va_arg(*va, (**ZendString)(_))
		if ZendParseArgPathStr(arg, str, check_null) == 0 {
			return "a valid path"
		}
		break
	case 'S':
		var str **ZendString = __va_arg(*va, (**ZendString)(_))
		if ZendParseArgStr(arg, str, check_null) == 0 {
			return "string"
		}
		break
	case 'b':
		var p *ZendBool = __va_arg(*va, (*ZendBool)(_))
		var is_null *ZendBool = nil
		if check_null != 0 {
			is_null = __va_arg(*va, (*ZendBool)(_))
		}
		if ZendParseArgBool(arg, p, is_null, check_null) == 0 {
			return "bool"
		}
		break
	case 'r':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		if ZendParseArgResource(arg, p, check_null) == 0 {
			return "resource"
		}
		break
	case 'A':

	case 'a':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		if ZendParseArgArray(arg, p, check_null, c == 'A') == 0 {
			return "array"
		}
		break
	case 'H':

	case 'h':
		var p **HashTable = __va_arg(*va, (**HashTable)(_))
		if ZendParseArgArrayHt(arg, p, check_null, c == 'H', separate) == 0 {
			return "array"
		}
		break
	case 'o':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		if ZendParseArgObject(arg, p, nil, check_null) == 0 {
			return "object"
		}
		break
	case 'O':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		var ce *ZendClassEntry = __va_arg(*va, (*ZendClassEntry)(_))
		if ZendParseArgObject(arg, p, ce, check_null) == 0 {
			if ce != nil {
				return ce.GetName().GetVal()
			} else {
				return "object"
			}
		}
		break
	case 'C':
		var lookup *ZendClassEntry
		var pce **ZendClassEntry = __va_arg(*va, (**ZendClassEntry)(_))
		var ce_base *ZendClassEntry = *pce
		if check_null != 0 && arg.GetType() == 1 {
			*pce = nil
			break
		}
		if TryConvertToString(arg) == 0 {
			*pce = nil
			return "valid class name"
		}
		if g.Assign(&lookup, ZendLookupClass(arg.GetValue().GetStr())) == nil {
			*pce = nil
		} else {
			*pce = lookup
		}
		if ce_base != nil {
			if (*pce) == nil || InstanceofFunction(*pce, ce_base) == 0 {
				ZendSpprintf(error, 0, "to be a class name derived from %s, '%s' given", ce_base.GetName().GetVal(), arg.GetValue().GetStr().GetVal())
				*pce = nil
				return ""
			}
		}
		if (*pce) == nil {
			ZendSpprintf(error, 0, "to be a valid class name, '%s' given", arg.GetValue().GetStr().GetVal())
			return ""
		}
		break
		break
	case 'f':
		var fci *ZendFcallInfo = __va_arg(*va, (*ZendFcallInfo)(_))
		var fcc *ZendFcallInfoCache = __va_arg(*va, (*ZendFcallInfoCache)(_))
		var is_callable_error *byte = nil
		if check_null != 0 && arg.GetType() == 1 {
			fci.SetSize(0)
			fcc.SetFunctionHandler(0)
			break
		}
		if ZendFcallInfoInit(arg, 0, fci, fcc, nil, &is_callable_error) == SUCCESS {
			if is_callable_error != nil {
				*severity = 1 << 13
				ZendSpprintf(error, 0, "to be a valid callback, %s", is_callable_error)
				_efree(is_callable_error)
				*spec = spec_walk
				return ""
			}
			break
		} else {
			if is_callable_error != nil {
				*severity = 1 << 0
				ZendSpprintf(error, 0, "to be a valid callback, %s", is_callable_error)
				_efree(is_callable_error)
				return ""
			} else {
				return "valid callback"
			}
		}
	case 'z':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		ZendParseArgZvalDeref(real_arg, p, check_null)
		break
	case 'Z':

		/* 'Z' iz not supported anymore and should be replaced with 'z' */

		assert(c != 'Z')
	default:
		return "unknown"
	}
	*spec = spec_walk
	return nil
}

/* }}} */

func ZendParseArg(arg_num int, arg *Zval, va *va_list, spec **byte, flags int) int {
	var expected_type *byte = nil
	var error *byte = nil
	var severity int = 0
	expected_type = ZendParseArgImpl(arg_num, arg, va, spec, &error, &severity)
	if expected_type != nil {
		if EG.GetException() != nil {
			return FAILURE
		}
		if (flags&1<<1) == 0 && ((*expected_type) || error != nil) {
			var space *byte
			var class_name *byte = GetActiveClassName(&space)
			var throw_exception ZendBool = EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0 || (flags&1<<2) != 0
			if error != nil {
				ZendInternalTypeError(throw_exception, "%s%s%s() expects parameter %d %s", class_name, space, GetActiveFunctionName(), arg_num, error)
				_efree(error)
			} else {
				ZendInternalTypeError(throw_exception, "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), arg_num, expected_type, ZendZvalTypeName(arg))
			}
		}
		if severity != 1<<13 {
			return FAILURE
		}
	}
	return SUCCESS
}

/* }}} */

func ZendParseParameter(flags int, arg_num int, arg *Zval, spec *byte, _ ...any) int {
	var va va_list
	var ret int
	va_start(va, spec)
	ret = ZendParseArg(arg_num, arg, &va, &spec, flags)
	va_end(va)
	return ret
}
func ZendParseParametersDebugError(msg string) {
	var active_function *ZendFunction = EG.GetCurrentExecuteData().GetFunc()
	var class_name *byte = g.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendErrorNoreturn(1<<4, "%s%s%s(): %s", class_name, g.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), msg)
}
func ZendParseVaArgs(num_args int, type_spec *byte, va *va_list, flags int) int {
	var spec_walk *byte
	var c int
	var i int
	var min_num_args int = -1
	var max_num_args int = 0
	var post_varargs int = 0
	var arg *Zval
	var arg_count int
	var have_varargs ZendBool = 0
	var varargs **Zval = nil
	var n_varargs *int = nil
	for spec_walk = type_spec; *spec_walk; spec_walk++ {
		c = *spec_walk
		switch c {
		case 'l':

		case 'd':

		case 's':

		case 'b':

		case 'r':

		case 'a':

		case 'o':

		case 'O':

		case 'z':

		case 'Z':

		case 'C':

		case 'h':

		case 'f':

		case 'A':

		case 'H':

		case 'p':

		case 'S':

		case 'P':

		case 'L':
			max_num_args++
			break
		case '|':
			min_num_args = max_num_args
			break
		case '/':

		case '!':

			/* Pass */

			break
		case '*':

		case '+':
			if have_varargs != 0 {
				ZendParseParametersDebugError("only one varargs specifier (* or +) is permitted")
				return FAILURE
			}
			have_varargs = 1

			/* we expect at least one parameter in varargs */

			if c == '+' {
				max_num_args++
			}

			/* mark the beginning of varargs */

			post_varargs = max_num_args
			break
		default:
			ZendParseParametersDebugError("bad type specifier while parsing parameters")
			return FAILURE
		}
	}
	if min_num_args < 0 {
		min_num_args = max_num_args
	}
	if have_varargs != 0 {

		/* calculate how many required args are at the end of the specifier list */

		post_varargs = max_num_args - post_varargs
		max_num_args = -1
	}
	if num_args < min_num_args || num_args > max_num_args && max_num_args >= 0 {
		if (flags & 1 << 1) == 0 {
			var active_function *ZendFunction = EG.GetCurrentExecuteData().GetFunc()
			var class_name *byte = g.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
			var throw_exception ZendBool = EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0 || (flags&1<<2) != 0
			ZendInternalArgumentCountError(throw_exception, "%s%s%s() expects %s %d parameter%s, %d given", class_name, g.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), g.Cond(g.Cond(min_num_args == max_num_args, "exactly", num_args < min_num_args), "at least", "at most"), g.Cond(num_args < min_num_args, min_num_args, max_num_args), g.Cond(g.Cond(num_args < min_num_args, min_num_args, max_num_args) == 1, "", "s"), num_args)
		}
		return FAILURE
	}
	arg_count = EG.GetCurrentExecuteData().GetThis().GetNumArgs()
	if num_args > arg_count {
		ZendParseParametersDebugError("could not obtain parameters for parsing")
		return FAILURE
	}
	i = 0
	for g.PostDec(&num_args) > 0 {
		if (*type_spec) == '|' {
			type_spec++
		}
		if (*type_spec) == '*' || (*type_spec) == '+' {
			var num_varargs int = num_args + 1 - post_varargs

			/* eat up the passed in storage even if it won't be filled in with varargs */

			varargs = __va_arg(*va, (**Zval)(_))
			n_varargs = __va_arg(*va, (*int)(_))
			type_spec++
			if num_varargs > 0 {
				*n_varargs = num_varargs
				*varargs = (*Zval)(EG.GetCurrentExecuteData()) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(i+1)-1))

				/* adjust how many args we have left and restart loop */

				num_args += 1 - num_varargs
				i += num_varargs
				continue
			} else {
				*varargs = nil
				*n_varargs = 0
			}
		}
		arg = (*Zval)(EG.GetCurrentExecuteData()) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(i+1)-1))
		if ZendParseArg(i+1, arg, va, &type_spec, flags) == FAILURE {

			/* clean up varargs array if it was used */

			if varargs != nil && (*varargs) != nil {
				*varargs = nil
			}
			return FAILURE
		}
		i++
	}
	return SUCCESS
}

/* }}} */

func ZendParseParametersEx(flags int, num_args int, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	va_start(va, type_spec)
	retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
	va_end(va)
	return retval
}

/* }}} */

func ZendParseParameters(num_args int, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	var flags int = 0
	va_start(va, type_spec)
	retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
	va_end(va)
	return retval
}

/* }}} */

func ZendParseParametersThrow(num_args int, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	var flags int = 1 << 2
	va_start(va, type_spec)
	retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
	va_end(va)
	return retval
}

/* }}} */

func ZendParseMethodParameters(num_args int, this_ptr *Zval, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	var flags int = 0
	var p *byte = type_spec
	var object **Zval
	var ce *ZendClassEntry

	/* Just checking this_ptr is not enough, because fcall_common_helper does not set
	 * Z_OBJ(EG(This)) to NULL when calling an internal function with common.scope == NULL.
	 * In that case EG(This) would still be the $this from the calling code and we'd take the
	 * wrong branch here. */

	var is_method ZendBool = EG.GetCurrentExecuteData().GetFunc().GetScope() != nil
	if is_method == 0 || this_ptr == nil || this_ptr.GetType() != 8 {
		va_start(va, type_spec)
		retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
		va_end(va)
	} else {
		p++
		va_start(va, type_spec)
		object = __va_arg(va, (**Zval)(_))
		ce = __va_arg(va, (*ZendClassEntry)(_))
		*object = this_ptr
		if ce != nil && InstanceofFunction(this_ptr.GetValue().GetObj().GetCe(), ce) == 0 {
			ZendErrorNoreturn(1<<4, "%s::%s() must be derived from %s::%s", this_ptr.GetValue().GetObj().GetCe().GetName().GetVal(), GetActiveFunctionName(), ce.GetName().GetVal(), GetActiveFunctionName())
		}
		retval = ZendParseVaArgs(num_args, p, &va, flags)
		va_end(va)
	}
	return retval
}

/* }}} */

func ZendParseMethodParametersEx(flags int, num_args int, this_ptr *Zval, type_spec *byte, _ ...any) int {
	var va va_list
	var retval int
	var p *byte = type_spec
	var object **Zval
	var ce *ZendClassEntry
	if this_ptr == nil {
		va_start(va, type_spec)
		retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
		va_end(va)
	} else {
		p++
		va_start(va, type_spec)
		object = __va_arg(va, (**Zval)(_))
		ce = __va_arg(va, (*ZendClassEntry)(_))
		*object = this_ptr
		if ce != nil && InstanceofFunction(this_ptr.GetValue().GetObj().GetCe(), ce) == 0 {
			if (flags & 1 << 1) == 0 {
				ZendErrorNoreturn(1<<4, "%s::%s() must be derived from %s::%s", ce.GetName().GetVal(), GetActiveFunctionName(), this_ptr.GetValue().GetObj().GetCe().GetName().GetVal(), GetActiveFunctionName())
			}
			va_end(va)
			return FAILURE
		}
		retval = ZendParseVaArgs(num_args, p, &va, flags)
		va_end(va)
	}
	return retval
}

/* }}} */

func ZendMergeProperties(obj *Zval, properties *HashTable) {
	var obj_ht *ZendObjectHandlers = obj.GetValue().GetObj().GetHandlers()
	var old_scope *ZendClassEntry = EG.GetFakeScope()
	var key *ZendString
	var value *Zval
	EG.SetFakeScope(obj.GetValue().GetObj().GetCe())
	for {
		var __ht *HashTable = properties
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			key = _p.GetKey()
			value = _z
			if key != nil {
				var member Zval
				var __z *Zval = &member
				var __s *ZendString = key
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				obj_ht.GetWriteProperty()(obj, &member, value, nil)
			}
		}
		break
	}
	EG.SetFakeScope(old_scope)
}

/* }}} */

func ZendUpdateClassConstants(class_type *ZendClassEntry) int {
	if (class_type.GetCeFlags() & 1 << 12) == 0 {
		var ce *ZendClassEntry
		var c *ZendClassConstant
		var val *Zval
		var prop_info *ZendPropertyInfo
		if class_type.parent {
			if ZendUpdateClassConstants(class_type.parent) != SUCCESS {
				return FAILURE
			}
		}
		for {
			var __ht *HashTable = &class_type.constants_table
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				c = _z.GetValue().GetPtr()
				val = &c.value
				if val.GetType() == 11 {
					if ZvalUpdateConstantEx(val, c.GetCe()) != SUCCESS {
						return FAILURE
					}
				}
			}
			break
		}
		if class_type.GetDefaultStaticMembersCount() != 0 && (*Zval)(g.CondF((uintptr_t(class_type).static_members_table__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(class_type).static_members_table__ptr - 1)))
		}, func() any { return any(*(class_type.GetStaticMembersTablePtr())) })) == nil {
			if class_type.GetType() == 1 || (class_type.GetCeFlags()&(1<<7|1<<10)) != 0 {
				ZendClassInitStatics(class_type)
			}
		}
		ce = class_type
		for ce != nil {
			for {
				var __ht *HashTable = &ce.properties_info
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					prop_info = _z.GetValue().GetPtr()
					if prop_info.GetCe() == ce {
						if (prop_info.GetFlags() & 1 << 4) != 0 {
							val = (*Zval)(g.CondF((uintptr_t(class_type).static_members_table__ptr&1) != 0, func() any {
								return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(class_type).static_members_table__ptr - 1)))
							}, func() any { return any(*(class_type.GetStaticMembersTablePtr())) })) + prop_info.GetOffset()
						} else {
							val = (*Zval)((*byte)(class_type.GetDefaultPropertiesTable() + prop_info.GetOffset() - uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0)))
						}
						if val.GetType() == 11 {
							if prop_info.GetType() != 0 {
								var tmp Zval
								var _z1 *Zval = &tmp
								var _z2 *Zval = val
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								if (_t & 0xff00) != 0 {
									ZendGcAddref(&_gc.gc)
								}
								if ZvalUpdateConstantEx(&tmp, ce) != SUCCESS {
									ZvalPtrDtor(&tmp)
									return FAILURE
								}
								if ZendVerifyPropertyType(prop_info, &tmp, 1) == 0 {
									ZvalPtrDtor(&tmp)
									return FAILURE
								}
								ZvalPtrDtor(val)
								var _z1 *Zval = val
								var _z2 *Zval = &tmp
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
							} else if ZvalUpdateConstantEx(val, ce) != SUCCESS {
								return FAILURE
							}
						}
					}
				}
				break
			}
			ce = ce.parent
		}
		class_type.SetCeFlags(class_type.GetCeFlags() | 1<<12)
	}
	return SUCCESS
}

/* }}} */

func _objectPropertiesInit(object *ZendObject, class_type *ZendClassEntry) {
	if class_type.GetDefaultPropertiesCount() != 0 {
		var src *Zval = class_type.GetDefaultPropertiesTable()
		var dst *Zval = object.GetPropertiesTable()
		var end *Zval = src + class_type.GetDefaultPropertiesCount()
		if class_type.GetType() == 1 {
			for {
				var _z1 *Zval = dst
				var _z2 *Zval = src
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
						ZendGcAddref(&_gc.gc)
					} else {
						ZvalCopyCtorFunc(_z1)
					}
				}
				dst.SetU2Extra(src.GetU2Extra())
				src++
				dst++
				if src == end {
					break
				}
			}
		} else {
			for {
				var _z1 *Zval = dst
				var _z2 *Zval = src
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
				dst.SetU2Extra(src.GetU2Extra())
				src++
				dst++
				if src == end {
					break
				}
			}
		}
	}
}

/* }}} */

func ObjectPropertiesInit(object *ZendObject, class_type *ZendClassEntry) {
	object.SetProperties(nil)
	_objectPropertiesInit(object, class_type)
}

/* }}} */

func ObjectPropertiesInitEx(object *ZendObject, properties *HashTable) {
	object.SetProperties(properties)
	if object.GetCe().GetDefaultPropertiesCount() != 0 {
		var prop *Zval
		var key *ZendString
		var property_info *ZendPropertyInfo
		for {
			var __ht *HashTable = properties
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				key = _p.GetKey()
				prop = _z
				property_info = ZendGetPropertyInfo(object.GetCe(), key, 1)
				if property_info != (*ZendPropertyInfo)(intptr_t-1) && property_info != nil && (property_info.GetFlags()&1<<4) == 0 {
					var slot *Zval = (*Zval)((*byte)(object + property_info.GetOffset()))
					if property_info.GetType() != 0 {
						var tmp Zval
						var _z1 *Zval = &tmp
						var _z2 *Zval = prop
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if ZendVerifyPropertyType(property_info, &tmp, 0) == 0 {
							continue
						}
						var _z1 *Zval = slot
						var _z2 *Zval = &tmp
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
					} else {
						var _z1 *Zval = slot
						var _z2 *Zval = prop
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
					}
					prop.GetValue().SetZv(slot)
					prop.SetTypeInfo(13)
				}
			}
			break
		}
	}
}

/* }}} */

func ObjectPropertiesLoad(object *ZendObject, properties *HashTable) {
	var prop *Zval
	var tmp Zval
	var key *ZendString
	var h ZendLong
	var property_info *ZendPropertyInfo
	for {
		var __ht *HashTable = properties
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			h = _p.GetH()
			key = _p.GetKey()
			prop = _z
			if key != nil {
				if key.GetVal()[0] == '0' {
					var class_name *byte
					var prop_name *byte
					var prop_name_len int
					if ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, &prop_name_len) == SUCCESS {
						var pname *ZendString = ZendStringInit(prop_name, prop_name_len, 0)
						var prev_scope *ZendClassEntry = EG.GetFakeScope()
						if class_name != nil && class_name[0] != '*' {
							var cname *ZendString = ZendStringInit(class_name, strlen(class_name), 0)
							EG.SetFakeScope(ZendLookupClass(cname))
							ZendStringReleaseEx(cname, 0)
						}
						property_info = ZendGetPropertyInfo(object.GetCe(), pname, 1)
						ZendStringReleaseEx(pname, 0)
						EG.SetFakeScope(prev_scope)
					} else {
						property_info = (*ZendPropertyInfo)(intptr_t - 1)
					}
				} else {
					property_info = ZendGetPropertyInfo(object.GetCe(), key, 1)
				}
				if property_info != (*ZendPropertyInfo)(intptr_t-1) && property_info != nil && (property_info.GetFlags()&1<<4) == 0 {
					var slot *Zval = (*Zval)((*byte)(object + property_info.GetOffset()))
					ZvalPtrDtor(slot)
					var _z1 *Zval = slot
					var _z2 *Zval = prop
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					ZvalAddRef(slot)
					if object.GetProperties() != nil {
						&tmp.GetValue().SetZv(slot)
						&tmp.SetTypeInfo(13)
						ZendHashUpdate(object.GetProperties(), key, &tmp)
					}
				} else {
					if object.GetProperties() == nil {
						RebuildObjectProperties(object)
					}
					prop = ZendHashUpdate(object.GetProperties(), key, prop)
					ZvalAddRef(prop)
				}
			} else {
				if object.GetProperties() == nil {
					RebuildObjectProperties(object)
				}
				prop = ZendHashIndexUpdate(object.GetProperties(), h, prop)
				ZvalAddRef(prop)
			}
		}
		break
	}
}

/* }}} */

func _objectAndPropertiesInit(arg *Zval, class_type *ZendClassEntry, properties *HashTable) int {
	if (class_type.GetCeFlags() & (1<<0 | 1<<1 | 1<<4 | 1<<6)) != 0 {
		if (class_type.GetCeFlags() & 1 << 0) != 0 {
			ZendThrowError(nil, "Cannot instantiate interface %s", class_type.GetName().GetVal())
		} else if (class_type.GetCeFlags() & 1 << 1) != 0 {
			ZendThrowError(nil, "Cannot instantiate trait %s", class_type.GetName().GetVal())
		} else {
			ZendThrowError(nil, "Cannot instantiate abstract class %s", class_type.GetName().GetVal())
		}
		arg.SetTypeInfo(1)
		arg.GetValue().SetObj(nil)
		return FAILURE
	}
	if (class_type.GetCeFlags() & 1 << 12) == 0 {
		if ZendUpdateClassConstants(class_type) != SUCCESS {
			arg.SetTypeInfo(1)
			arg.GetValue().SetObj(nil)
			return FAILURE
		}
	}
	if class_type.create_object == nil {
		var obj *ZendObject = ZendObjectsNew(class_type)
		var __z *Zval = arg
		__z.GetValue().SetObj(obj)
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		if properties != nil {
			ObjectPropertiesInitEx(obj, properties)
		} else {
			_objectPropertiesInit(obj, class_type)
		}
	} else {
		var __z *Zval = arg
		__z.GetValue().SetObj(class_type.create_object(class_type))
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	}
	return SUCCESS
}

/* }}} */

func ObjectAndPropertiesInit(arg *Zval, class_type *ZendClassEntry, properties *HashTable) int {
	return _objectAndPropertiesInit(arg, class_type, properties)
}

/* }}} */

func ObjectInitEx(arg *Zval, class_type *ZendClassEntry) int {
	return _objectAndPropertiesInit(arg, class_type, nil)
}

/* }}} */

func ObjectInit(arg *Zval) int {
	var __z *Zval = arg
	__z.GetValue().SetObj(ZendObjectsNew(ZendStandardClassDef))
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	return SUCCESS
}

/* }}} */

func AddAssocLongEx(arg *Zval, key string, key_len int, n ZendLong) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetLval(n)
	__z.SetTypeInfo(4)
	ZendSymtableStrUpdate(arg.GetValue().GetArr(), key, key_len, &tmp)
	return SUCCESS
}

/* }}} */

func AddAssocNullEx(arg *Zval, key string, key_len int) int {
	var tmp Zval
	&tmp.SetTypeInfo(1)
	ZendSymtableStrUpdate(arg.GetValue().GetArr(), key, key_len, &tmp)
	return SUCCESS
}

/* }}} */

func AddAssocBoolEx(arg *Zval, key string, key_len int, b int) int {
	var tmp Zval
	if b != 0 {
		&tmp.SetTypeInfo(3)
	} else {
		&tmp.SetTypeInfo(2)
	}
	ZendSymtableStrUpdate(arg.GetValue().GetArr(), key, key_len, &tmp)
	return SUCCESS
}

/* }}} */

func AddAssocResourceEx(arg *Zval, key *byte, key_len int, r *ZendResource) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetRes(r)
	__z.SetTypeInfo(9 | 1<<0<<8)
	ZendSymtableStrUpdate(arg.GetValue().GetArr(), key, key_len, &tmp)
	return SUCCESS
}

/* }}} */

func AddAssocDoubleEx(arg *Zval, key *byte, key_len int, d float64) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetDval(d)
	__z.SetTypeInfo(5)
	ZendSymtableStrUpdate(arg.GetValue().GetArr(), key, key_len, &tmp)
	return SUCCESS
}

/* }}} */

func AddAssocStrEx(arg *Zval, key string, key_len int, str *ZendString) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	ZendSymtableStrUpdate(arg.GetValue().GetArr(), key, key_len, &tmp)
	return SUCCESS
}

/* }}} */

func AddAssocStringEx(arg *Zval, key *byte, key_len int, str *byte) int {
	var tmp Zval
	var _s *byte = str
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	ZendSymtableStrUpdate(arg.GetValue().GetArr(), key, key_len, &tmp)
	return SUCCESS
}

/* }}} */

func AddAssocStringlEx(arg *Zval, key *byte, key_len int, str *byte, length int) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(str, length, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	ZendSymtableStrUpdate(arg.GetValue().GetArr(), key, key_len, &tmp)
	return SUCCESS
}

/* }}} */

func AddAssocZvalEx(arg *Zval, key string, key_len int, value *Zval) int {
	ZendSymtableStrUpdate(arg.GetValue().GetArr(), key, key_len, value)
	return SUCCESS
}

/* }}} */

func AddIndexLong(arg *Zval, index ZendUlong, n ZendLong) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetLval(n)
	__z.SetTypeInfo(4)
	ZendHashIndexUpdate(arg.GetValue().GetArr(), index, &tmp)
	return SUCCESS
}

/* }}} */

func AddIndexNull(arg *Zval, index ZendUlong) int {
	var tmp Zval
	&tmp.SetTypeInfo(1)
	ZendHashIndexUpdate(arg.GetValue().GetArr(), index, &tmp)
	return SUCCESS
}

/* }}} */

func AddIndexBool(arg *Zval, index ZendUlong, b int) int {
	var tmp Zval
	if b != 0 {
		&tmp.SetTypeInfo(3)
	} else {
		&tmp.SetTypeInfo(2)
	}
	ZendHashIndexUpdate(arg.GetValue().GetArr(), index, &tmp)
	return SUCCESS
}

/* }}} */

func AddIndexResource(arg *Zval, index ZendUlong, r *ZendResource) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetRes(r)
	__z.SetTypeInfo(9 | 1<<0<<8)
	ZendHashIndexUpdate(arg.GetValue().GetArr(), index, &tmp)
	return SUCCESS
}

/* }}} */

func AddIndexDouble(arg *Zval, index ZendUlong, d float64) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetDval(d)
	__z.SetTypeInfo(5)
	ZendHashIndexUpdate(arg.GetValue().GetArr(), index, &tmp)
	return SUCCESS
}

/* }}} */

func AddIndexStr(arg *Zval, index ZendUlong, str *ZendString) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	ZendHashIndexUpdate(arg.GetValue().GetArr(), index, &tmp)
	return SUCCESS
}

/* }}} */

func AddIndexString(arg *Zval, index ZendUlong, str *byte) int {
	var tmp Zval
	var _s *byte = str
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	ZendHashIndexUpdate(arg.GetValue().GetArr(), index, &tmp)
	return SUCCESS
}

/* }}} */

func AddIndexStringl(arg *Zval, index ZendUlong, str *byte, length int) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(str, length, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	ZendHashIndexUpdate(arg.GetValue().GetArr(), index, &tmp)
	return SUCCESS
}

/* }}} */

func AddNextIndexLong(arg *Zval, n ZendLong) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetLval(n)
	__z.SetTypeInfo(4)
	if ZendHashNextIndexInsert(arg.GetValue().GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func AddNextIndexNull(arg *Zval) int {
	var tmp Zval
	&tmp.SetTypeInfo(1)
	if ZendHashNextIndexInsert(arg.GetValue().GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func AddNextIndexBool(arg *Zval, b int) int {
	var tmp Zval
	if b != 0 {
		&tmp.SetTypeInfo(3)
	} else {
		&tmp.SetTypeInfo(2)
	}
	if ZendHashNextIndexInsert(arg.GetValue().GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func AddNextIndexResource(arg *Zval, r *ZendResource) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetRes(r)
	__z.SetTypeInfo(9 | 1<<0<<8)
	if ZendHashNextIndexInsert(arg.GetValue().GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func AddNextIndexDouble(arg *Zval, d float64) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetDval(d)
	__z.SetTypeInfo(5)
	if ZendHashNextIndexInsert(arg.GetValue().GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func AddNextIndexStr(arg *Zval, str *ZendString) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	if ZendHashNextIndexInsert(arg.GetValue().GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func AddNextIndexString(arg *Zval, str *byte) int {
	var tmp Zval
	var _s *byte = str
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	if ZendHashNextIndexInsert(arg.GetValue().GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func AddNextIndexStringl(arg *Zval, str *byte, length int) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(str, length, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	if ZendHashNextIndexInsert(arg.GetValue().GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func ArraySetZvalKey(ht *HashTable, key *Zval, value *Zval) int {
	var result *Zval
	switch key.GetType() {
	case 6:
		result = ZendSymtableUpdate(ht, key.GetValue().GetStr(), value)
		break
	case 1:
		result = ZendSymtableUpdate(ht, ZendEmptyString, value)
		break
	case 9:
		ZendError(1<<3, "Resource ID#%d used as offset, casting to integer (%d)", key.GetValue().GetRes().GetHandle(), key.GetValue().GetRes().GetHandle())
		result = ZendHashIndexUpdate(ht, key.GetValue().GetRes().GetHandle(), value)
		break
	case 2:
		result = ZendHashIndexUpdate(ht, 0, value)
		break
	case 3:
		result = ZendHashIndexUpdate(ht, 1, value)
		break
	case 4:
		result = ZendHashIndexUpdate(ht, key.GetValue().GetLval(), value)
		break
	case 5:
		result = ZendHashIndexUpdate(ht, ZendDvalToLval(key.GetValue().GetDval()), value)
		break
	default:
		ZendError(1<<1, "Illegal offset type")
		result = nil
	}
	if result != nil {
		if result.GetTypeFlags() != 0 {
			ZvalAddrefP(result)
		}
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func AddPropertyLongEx(arg *Zval, key string, key_len int, n ZendLong) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetLval(n)
	__z.SetTypeInfo(4)
	return AddPropertyZvalEx(arg, key, key_len, &tmp)
}

/* }}} */

func AddPropertyBoolEx(arg *Zval, key *byte, key_len int, b ZendLong) int {
	var tmp Zval
	if b != 0 {
		&tmp.SetTypeInfo(3)
	} else {
		&tmp.SetTypeInfo(2)
	}
	return AddPropertyZvalEx(arg, key, key_len, &tmp)
}

/* }}} */

func AddPropertyNullEx(arg *Zval, key string, key_len int) int {
	var tmp Zval
	&tmp.SetTypeInfo(1)
	return AddPropertyZvalEx(arg, key, key_len, &tmp)
}

/* }}} */

func AddPropertyResourceEx(arg *Zval, key string, key_len int, r *ZendResource) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetRes(r)
	__z.SetTypeInfo(9 | 1<<0<<8)
	AddPropertyZvalEx(arg, key, key_len, &tmp)
	ZvalPtrDtor(&tmp)
	return SUCCESS
}

/* }}} */

func AddPropertyDoubleEx(arg *Zval, key *byte, key_len int, d float64) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetDval(d)
	__z.SetTypeInfo(5)
	return AddPropertyZvalEx(arg, key, key_len, &tmp)
}

/* }}} */

func AddPropertyStrEx(arg *Zval, key *byte, key_len int, str *ZendString) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	AddPropertyZvalEx(arg, key, key_len, &tmp)
	ZvalPtrDtor(&tmp)
	return SUCCESS
}

/* }}} */

func AddPropertyStringEx(arg *Zval, key string, key_len int, str *byte) int {
	var tmp Zval
	var _s *byte = str
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	AddPropertyZvalEx(arg, key, key_len, &tmp)
	ZvalPtrDtor(&tmp)
	return SUCCESS
}

/* }}} */

func AddPropertyStringlEx(arg *Zval, key string, key_len int, str *byte, length int) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(str, length, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	AddPropertyZvalEx(arg, key, key_len, &tmp)
	ZvalPtrDtor(&tmp)
	return SUCCESS
}

/* }}} */

func AddPropertyZvalEx(arg *Zval, key string, key_len int, value *Zval) int {
	var z_key Zval
	var __z *Zval = &z_key
	var __s *ZendString = ZendStringInit(key, key_len, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	arg.GetValue().GetObj().GetHandlers().GetWriteProperty()(arg, &z_key, value, nil)
	ZvalPtrDtor(&z_key)
	return SUCCESS
}

/* }}} */

func ZendStartupModuleEx(module *ZendModuleEntry) int {
	var name_len int
	var lcname *ZendString
	if module.GetModuleStarted() != 0 {
		return SUCCESS
	}
	module.SetModuleStarted(1)

	/* Check module dependencies */

	if module.GetDeps() != nil {
		var dep *ZendModuleDep = module.GetDeps()
		for dep.GetName() != nil {
			if dep.GetType() == 1 {
				var req_mod *ZendModuleEntry
				name_len = strlen(dep.GetName())
				lcname = ZendStringAlloc(name_len, 0)
				ZendStrTolowerCopy(lcname.GetVal(), dep.GetName(), name_len)
				if g.Assign(&req_mod, ZendHashFindPtr(&ModuleRegistry, lcname)) == nil || req_mod.GetModuleStarted() == 0 {
					ZendStringEfree(lcname)

					/* TODO: Check version relationship */

					ZendError(1<<5, "Cannot load module '%s' because required module '%s' is not loaded", module.GetName(), dep.GetName())
					module.SetModuleStarted(0)
					return FAILURE
				}
				ZendStringEfree(lcname)
			}
			dep++
		}
	}

	/* Initialize module globals */

	if module.GetGlobalsSize() != 0 {
		if module.GetGlobalsCtor() != nil {
			module.GetGlobalsCtor()(module.GetGlobalsPtr())
		}
	}
	if module.GetModuleStartupFunc() != nil {
		EG.SetCurrentModule(module)
		if module.GetModuleStartupFunc()(module.GetType(), module.GetModuleNumber()) == FAILURE {
			ZendErrorNoreturn(1<<4, "Unable to start %s module", module.GetName())
			EG.SetCurrentModule(nil)
			return FAILURE
		}
		EG.SetCurrentModule(nil)
	}
	return SUCCESS
}

/* }}} */

func ZendStartupModuleZval(zv *Zval) int {
	var module *ZendModuleEntry = zv.GetValue().GetPtr()
	if ZendStartupModuleEx(module) == SUCCESS {
		return 0
	} else {
		return 1 << 0
	}
}

/* }}} */

func ZendSortModules(base any, count int, siz int, compare CompareFuncT, swp SwapFuncT) {
	var b1 *Bucket = base
	var b2 *Bucket
	var end *Bucket = b1 + count
	var tmp Bucket
	var m *ZendModuleEntry
	var r *ZendModuleEntry
	for b1 < end {
	try_again:
		m = (*ZendModuleEntry)(b1.GetVal().GetValue().GetPtr())
		if m.GetModuleStarted() == 0 && m.GetDeps() != nil {
			var dep *ZendModuleDep = m.GetDeps()
			for dep.GetName() != nil {
				if dep.GetType() == 1 || dep.GetType() == 3 {
					b2 = b1 + 1
					for b2 < end {
						r = (*ZendModuleEntry)(b2.GetVal().GetValue().GetPtr())
						if strcasecmp(dep.GetName(), r.GetName()) == 0 {
							tmp = *b1
							*b1 = *b2
							*b2 = tmp
							goto try_again
						}
						b2++
					}
				}
				dep++
			}
		}
		b1++
	}
}

/* }}} */

func ZendCollectModuleHandlers() {
	var module *ZendModuleEntry
	var startup_count int = 0
	var shutdown_count int = 0
	var post_deactivate_count int = 0
	var ce *ZendClassEntry
	var class_count int = 0

	/* Collect extensions with request startup/shutdown handlers */

	for {
		var __ht *HashTable = &ModuleRegistry
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			module = _z.GetValue().GetPtr()
			if module.GetRequestStartupFunc() != nil {
				startup_count++
			}
			if module.GetRequestShutdownFunc() != nil {
				shutdown_count++
			}
			if module.GetPostDeactivateFunc() != nil {
				post_deactivate_count++
			}
		}
		break
	}
	ModuleRequestStartupHandlers = (**ZendModuleEntry)(Malloc(g.SizeOf("zend_module_entry *") * (startup_count + 1 + shutdown_count + 1 + post_deactivate_count + 1)))
	ModuleRequestStartupHandlers[startup_count] = nil
	ModuleRequestShutdownHandlers = ModuleRequestStartupHandlers + startup_count + 1
	ModuleRequestShutdownHandlers[shutdown_count] = nil
	ModulePostDeactivateHandlers = ModuleRequestShutdownHandlers + shutdown_count + 1
	ModulePostDeactivateHandlers[post_deactivate_count] = nil
	startup_count = 0
	for {
		var __ht *HashTable = &ModuleRegistry
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			module = _z.GetValue().GetPtr()
			if module.GetRequestStartupFunc() != nil {
				ModuleRequestStartupHandlers[g.PostInc(&startup_count)] = module
			}
			if module.GetRequestShutdownFunc() != nil {
				ModuleRequestShutdownHandlers[g.PreDec(&shutdown_count)] = module
			}
			if module.GetPostDeactivateFunc() != nil {
				ModulePostDeactivateHandlers[g.PreDec(&post_deactivate_count)] = module
			}
		}
		break
	}

	/* Collect internal classes with static members */

	for {
		var __ht *HashTable = CG.GetClassTable()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			ce = _z.GetValue().GetPtr()
			if ce.GetType() == 1 && ce.GetDefaultStaticMembersCount() > 0 {
				class_count++
			}
		}
		break
	}
	ClassCleanupHandlers = (**ZendClassEntry)(Malloc(g.SizeOf("zend_class_entry *") * (class_count + 1)))
	ClassCleanupHandlers[class_count] = nil
	if class_count != 0 {
		for {
			var __ht *HashTable = CG.GetClassTable()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				ce = _z.GetValue().GetPtr()
				if ce.GetType() == 1 && ce.GetDefaultStaticMembersCount() > 0 {
					ClassCleanupHandlers[g.PreDec(&class_count)] = ce
				}
			}
			break
		}
	}
}

/* }}} */

func ZendStartupModules() int {
	ZendHashSortEx(&ModuleRegistry, ZendSortModules, nil, 0)
	ZendHashApply(&ModuleRegistry, ZendStartupModuleZval)
	return SUCCESS
}

/* }}} */

func ZendDestroyModules() {
	Free(ClassCleanupHandlers)
	Free(ModuleRequestStartupHandlers)
	ZendHashGracefulReverseDestroy(&ModuleRegistry)
}

/* }}} */

func ZendRegisterModuleEx(module *ZendModuleEntry) *ZendModuleEntry {
	var name_len int
	var lcname *ZendString
	var module_ptr *ZendModuleEntry
	if module == nil {
		return nil
	}

	/* Check module dependencies */

	if module.GetDeps() != nil {
		var dep *ZendModuleDep = module.GetDeps()
		for dep.GetName() != nil {
			if dep.GetType() == 2 {
				name_len = strlen(dep.GetName())
				lcname = ZendStringAlloc(name_len, 0)
				ZendStrTolowerCopy(lcname.GetVal(), dep.GetName(), name_len)
				if ZendHashExists(&ModuleRegistry, lcname) != 0 || ZendGetExtension(dep.GetName()) != nil {
					ZendStringEfree(lcname)

					/* TODO: Check version relationship */

					ZendError(1<<5, "Cannot load module '%s' because conflicting module '%s' is already loaded", module.GetName(), dep.GetName())
					return nil
				}
				ZendStringEfree(lcname)
			}
			dep++
		}
	}
	name_len = strlen(module.GetName())
	lcname = ZendStringAlloc(name_len, module.GetType() == 1)
	ZendStrTolowerCopy(lcname.GetVal(), module.GetName(), name_len)
	lcname = ZendNewInternedString(lcname)
	if g.Assign(&module_ptr, ZendHashAddMem(&ModuleRegistry, lcname, module, g.SizeOf("zend_module_entry"))) == nil {
		ZendError(1<<5, "Module '%s' already loaded", module.GetName())
		ZendStringRelease(lcname)
		return nil
	}
	module = module_ptr
	EG.SetCurrentModule(module)
	if module.GetFunctions() != nil && ZendRegisterFunctions(nil, module.GetFunctions(), nil, module.GetType()) == FAILURE {
		ZendHashDel(&ModuleRegistry, lcname)
		ZendStringRelease(lcname)
		EG.SetCurrentModule(nil)
		ZendError(1<<5, "%s: Unable to register functions, unable to load", module.GetName())
		return nil
	}
	EG.SetCurrentModule(nil)
	ZendStringRelease(lcname)
	return module
}

/* }}} */

func ZendRegisterInternalModule(module *ZendModuleEntry) *ZendModuleEntry {
	module.SetModuleNumber(ZendNextFreeModule())
	module.SetType(1)
	return ZendRegisterModuleEx(module)
}

/* }}} */

func ZendCheckMagicMethodImplementation(ce *ZendClassEntry, fptr *ZendFunction, error_type int) {
	var lcname []byte
	var name_len int
	if fptr.GetFunctionName().GetVal()[0] != '_' || fptr.GetFunctionName().GetVal()[1] != '_' {
		return
	}

	/* we don't care if the function name is longer, in fact lowercasing only
	 * the beginning of the name speeds up the check process */

	name_len = fptr.GetFunctionName().GetLen()
	ZendStrTolowerCopy(lcname, fptr.GetFunctionName().GetVal(), g.CondF2(name_len < g.SizeOf("lcname")-1, name_len, func() int { return g.SizeOf("lcname") - 1 }))
	lcname[g.SizeOf("lcname")-1] = '0'
	if name_len == g.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1 && !(memcmp(lcname, "__destruct", g.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Destructor %s::%s() cannot take arguments", ce.GetName().GetVal(), "__destruct")
	} else if name_len == g.SizeOf("ZEND_CLONE_FUNC_NAME")-1 && !(memcmp(lcname, "__clone", g.SizeOf("ZEND_CLONE_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Method %s::%s() cannot accept any arguments", ce.GetName().GetVal(), "__clone")
	} else if name_len == g.SizeOf("ZEND_GET_FUNC_NAME")-1 && !(memcmp(lcname, "__get", g.SizeOf("ZEND_GET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 1 {
			ZendError(error_type, "Method %s::%s() must take exactly 1 argument", ce.GetName().GetVal(), "__get")
		} else if (fptr.GetQuickArgFlags() >> (1 + 3) * 2 & (1 | 2)) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), "__get")
		}
	} else if name_len == g.SizeOf("ZEND_SET_FUNC_NAME")-1 && !(memcmp(lcname, "__set", g.SizeOf("ZEND_SET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 2 {
			ZendError(error_type, "Method %s::%s() must take exactly 2 arguments", ce.GetName().GetVal(), "__set")
		} else if (fptr.GetQuickArgFlags()>>(1+3)*2&(1|2)) != 0 || (fptr.GetQuickArgFlags()>>(2+3)*2&(1|2)) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), "__set")
		}
	} else if name_len == g.SizeOf("ZEND_UNSET_FUNC_NAME")-1 && !(memcmp(lcname, "__unset", g.SizeOf("ZEND_UNSET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 1 {
			ZendError(error_type, "Method %s::%s() must take exactly 1 argument", ce.GetName().GetVal(), "__unset")
		} else if (fptr.GetQuickArgFlags() >> (1 + 3) * 2 & (1 | 2)) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), "__unset")
		}
	} else if name_len == g.SizeOf("ZEND_ISSET_FUNC_NAME")-1 && !(memcmp(lcname, "__isset", g.SizeOf("ZEND_ISSET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 1 {
			ZendError(error_type, "Method %s::%s() must take exactly 1 argument", ce.GetName().GetVal(), "__isset")
		} else if (fptr.GetQuickArgFlags() >> (1 + 3) * 2 & (1 | 2)) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), "__isset")
		}
	} else if name_len == g.SizeOf("ZEND_CALL_FUNC_NAME")-1 && !(memcmp(lcname, "__call", g.SizeOf("ZEND_CALL_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 2 {
			ZendError(error_type, "Method %s::%s() must take exactly 2 arguments", ce.GetName().GetVal(), "__call")
		} else if (fptr.GetQuickArgFlags()>>(1+3)*2&(1|2)) != 0 || (fptr.GetQuickArgFlags()>>(2+3)*2&(1|2)) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), "__call")
		}
	} else if name_len == g.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1 && !(memcmp(lcname, "__callstatic", g.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 2 {
			ZendError(error_type, "Method %s::__callStatic() must take exactly 2 arguments", ce.GetName().GetVal())
		} else if (fptr.GetQuickArgFlags()>>(1+3)*2&(1|2)) != 0 || (fptr.GetQuickArgFlags()>>(2+3)*2&(1|2)) != 0 {
			ZendError(error_type, "Method %s::__callStatic() cannot take arguments by reference", ce.GetName().GetVal())
		}
	} else if name_len == g.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1 && !(memcmp(lcname, "__tostring", g.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Method %s::%s() cannot take arguments", ce.GetName().GetVal(), "__tostring")
	} else if name_len == g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1 && !(memcmp(lcname, "__debuginfo", g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Method %s::%s() cannot take arguments", ce.GetName().GetVal(), "__debuginfo")
	}
}

/* }}} */

func ZendRegisterFunctions(scope *ZendClassEntry, functions *ZendFunctionEntry, function_table *HashTable, type_ int) int {
	var ptr *ZendFunctionEntry = functions
	var function ZendFunction
	var reg_function *ZendFunction
	var internal_function *ZendInternalFunction = (*ZendInternalFunction)(&function)
	var count int = 0
	var unload int = 0
	var target_function_table *HashTable = function_table
	var error_type int
	var ctor *ZendFunction = nil
	var dtor *ZendFunction = nil
	var clone *ZendFunction = nil
	var __get *ZendFunction = nil
	var __set *ZendFunction = nil
	var __unset *ZendFunction = nil
	var __isset *ZendFunction = nil
	var __call *ZendFunction = nil
	var __callstatic *ZendFunction = nil
	var __tostring *ZendFunction = nil
	var __debugInfo *ZendFunction = nil
	var serialize_func *ZendFunction = nil
	var unserialize_func *ZendFunction = nil
	var lowercase_name *ZendString
	var fname_len int
	var lc_class_name *byte = nil
	var class_name_len int = 0
	if type_ == 1 {
		error_type = 1 << 5
	} else {
		error_type = 1 << 1
	}
	if target_function_table == nil {
		target_function_table = CG.GetFunctionTable()
	}
	internal_function.SetType(1)
	internal_function.SetModule(EG.GetCurrentModule())
	memset(internal_function.GetReserved(), 0, 6*g.SizeOf("void *"))
	if scope != nil {
		class_name_len = scope.GetName().GetLen()
		if g.Assign(&lc_class_name, ZendMemrchr(scope.GetName().GetVal(), '\\', class_name_len)) {
			lc_class_name++
			class_name_len -= lc_class_name - scope.GetName().GetVal()
			lc_class_name = ZendStrTolowerDup(lc_class_name, class_name_len)
		} else {
			lc_class_name = ZendStrTolowerDup(scope.GetName().GetVal(), class_name_len)
		}
	}
	for ptr.GetFname() != nil {
		fname_len = strlen(ptr.GetFname())
		internal_function.SetHandler(ptr.GetHandler())
		internal_function.SetFunctionName(ZendStringInitInterned(ptr.GetFname(), fname_len, 1))
		internal_function.SetScope(scope)
		internal_function.SetPrototype(nil)
		if ptr.GetFlags() != 0 {
			if (ptr.GetFlags() & (1<<0 | 1<<1 | 1<<2)) == 0 {
				if ptr.GetFlags() != 1<<11 && scope != nil {
					ZendError(error_type, "Invalid access level for %s%s%s() - access must be exactly one of public, protected or private", g.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), g.Cond(scope != nil, "::", ""), ptr.GetFname())
				}
				internal_function.SetFnFlags(1<<0 | ptr.GetFlags())
			} else {
				internal_function.SetFnFlags(ptr.GetFlags())
			}
		} else {
			internal_function.SetFnFlags(1 << 0)
		}
		if ptr.GetArgInfo() != nil {
			var info *ZendInternalFunctionInfo = (*ZendInternalFunctionInfo)(ptr.GetArgInfo())
			internal_function.SetArgInfo((*ZendInternalArgInfo)(ptr.GetArgInfo() + 1))
			internal_function.SetNumArgs(ptr.GetNumArgs())

			/* Currently you cannot denote that the function can accept less arguments than num_args */

			if info.GetRequiredNumArgs() == zend_uintptr_t-1 {
				internal_function.SetRequiredNumArgs(ptr.GetNumArgs())
			} else {
				internal_function.SetRequiredNumArgs(info.GetRequiredNumArgs())
			}
			if info.GetReturnReference() != 0 {
				internal_function.SetFnFlags(internal_function.GetFnFlags() | 1<<12)
			}
			if ptr.GetArgInfo()[ptr.GetNumArgs()].GetIsVariadic() != 0 {
				internal_function.SetFnFlags(internal_function.GetFnFlags() | 1<<14)

				/* Don't count the variadic argument */

				internal_function.GetNumArgs()--

				/* Don't count the variadic argument */

			}
			if info.GetType() > 0x3 {
				if info.GetType() > 0x3ff {
					var type_name *byte = (*byte)(info.GetType())
					if type_name[0] == '?' {
						type_name++
					}
					if scope == nil && (!(strcasecmp(type_name, "self")) || !(strcasecmp(type_name, "parent"))) {
						ZendErrorNoreturn(1<<4, "Cannot declare a return type of %s outside of a class scope", type_name)
					}
				}
				internal_function.SetFnFlags(internal_function.GetFnFlags() | 1<<13)
			}
		} else {
			internal_function.SetArgInfo(nil)
			internal_function.SetNumArgs(0)
			internal_function.SetRequiredNumArgs(0)
		}
		ZendSetFunctionArgFlags((*ZendFunction)(internal_function))
		if (ptr.GetFlags() & 1 << 6) != 0 {
			if scope != nil {

				/* This is a class that must be abstract itself. Here we set the check info. */

				scope.SetCeFlags(scope.GetCeFlags() | 1<<4)
				if (scope.GetCeFlags() & 1 << 0) == 0 {

					/* Since the class is not an interface it needs to be declared as a abstract class. */

					scope.SetCeFlags(scope.GetCeFlags() | 1<<6)

					/* Since the class is not an interface it needs to be declared as a abstract class. */

				}
			}
			if (ptr.GetFlags()&1<<4) != 0 && (scope == nil || (scope.GetCeFlags()&1<<0) == 0) {
				ZendError(error_type, "Static function %s%s%s() cannot be abstract", g.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), g.Cond(scope != nil, "::", ""), ptr.GetFname())
			}
		} else {
			if scope != nil && (scope.GetCeFlags()&1<<0) != 0 {
				_efree((*byte)(lc_class_name))
				ZendError(error_type, "Interface %s cannot contain non abstract method %s()", scope.GetName().GetVal(), ptr.GetFname())
				return FAILURE
			}
			if internal_function.GetHandler() == nil {
				if scope != nil {
					_efree((*byte)(lc_class_name))
				}
				ZendError(error_type, "Method %s%s%s() cannot be a NULL function", g.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), g.Cond(scope != nil, "::", ""), ptr.GetFname())
				ZendUnregisterFunctions(functions, count, target_function_table)
				return FAILURE
			}
		}
		lowercase_name = ZendStringTolowerEx(internal_function.GetFunctionName(), type_ == 1)
		lowercase_name = ZendNewInternedString(lowercase_name)
		reg_function = Malloc(g.SizeOf("zend_internal_function"))
		memcpy(reg_function, &function, g.SizeOf("zend_internal_function"))
		if ZendHashAddPtr(target_function_table, lowercase_name, reg_function) == nil {
			unload = 1
			Free(reg_function)
			ZendStringRelease(lowercase_name)
			break
		}

		/* If types of arguments have to be checked */

		if reg_function.GetArgInfo() != nil && reg_function.GetNumArgs() != 0 {
			var i uint32
			for i = 0; i < reg_function.GetNumArgs(); i++ {
				if reg_function.GetArgInfo()[i].GetType() > 0x3 {
					reg_function.SetFnFlags(reg_function.GetFnFlags() | 1<<8)
					break
				}
			}
		}
		if reg_function.GetArgInfo() != nil && (reg_function.GetFnFlags()&(1<<13|1<<8)) != 0 {

			/* convert "const char*" class type names into "zend_string*" */

			var i uint32
			var num_args uint32 = reg_function.GetNumArgs() + 1
			var arg_info *ZendArgInfo = reg_function.GetArgInfo() - 1
			var new_arg_info *ZendArgInfo
			if (reg_function.GetFnFlags() & 1 << 14) != 0 {
				num_args++
			}
			new_arg_info = Malloc(g.SizeOf("zend_arg_info") * num_args)
			memcpy(new_arg_info, arg_info, g.SizeOf("zend_arg_info")*num_args)
			reg_function.SetArgInfo(new_arg_info + 1)
			for i = 0; i < num_args; i++ {
				if new_arg_info[i].GetType() > 0x3ff {
					var class_name *byte = (*byte)(new_arg_info[i].GetType())
					var allow_null ZendBool = 0
					var str *ZendString
					if class_name[0] == '?' {
						class_name++
						allow_null = 1
					}
					str = ZendStringInitInterned(class_name, strlen(class_name), 1)
					new_arg_info[i].SetType(uintptr_t(str) | g.Cond(allow_null != 0, 0x1, 0x0))
				}
			}
		}
		if scope != nil {

			/* Look for ctor, dtor, clone
			 * If it's an old-style constructor, store it only if we don't have
			 * a constructor already.
			 */

			if fname_len == class_name_len && ctor == nil && !(memcmp(lowercase_name.GetVal(), lc_class_name, class_name_len+1)) {
				ctor = reg_function
			} else if lowercase_name.GetLen() == g.SizeOf("\"serialize\"")-1 && !(memcmp(lowercase_name.GetVal(), "serialize", g.SizeOf("\"serialize\"")-1)) {
				serialize_func = reg_function
			} else if lowercase_name.GetLen() == g.SizeOf("\"unserialize\"")-1 && !(memcmp(lowercase_name.GetVal(), "unserialize", g.SizeOf("\"unserialize\"")-1)) {
				unserialize_func = reg_function
			} else if lowercase_name.GetVal()[0] != '_' || lowercase_name.GetVal()[1] != '_' {
				reg_function = nil
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_CONSTRUCTOR_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__construct", g.SizeOf("ZEND_CONSTRUCTOR_FUNC_NAME")-1)) {
				ctor = reg_function
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__destruct", g.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1)) {
				dtor = reg_function
				if internal_function.GetNumArgs() != 0 {
					ZendError(error_type, "Destructor %s::%s() cannot take arguments", scope.GetName().GetVal(), ptr.GetFname())
				}
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_CLONE_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__clone", g.SizeOf("ZEND_CLONE_FUNC_NAME")-1)) {
				clone = reg_function
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_CALL_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__call", g.SizeOf("ZEND_CALL_FUNC_NAME")-1)) {
				__call = reg_function
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__callstatic", g.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1)) {
				__callstatic = reg_function
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__tostring", g.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1)) {
				__tostring = reg_function
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_GET_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__get", g.SizeOf("ZEND_GET_FUNC_NAME")-1)) {
				__get = reg_function
				scope.SetCeFlags(scope.GetCeFlags() | 1<<11)
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_SET_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__set", g.SizeOf("ZEND_SET_FUNC_NAME")-1)) {
				__set = reg_function
				scope.SetCeFlags(scope.GetCeFlags() | 1<<11)
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_UNSET_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__unset", g.SizeOf("ZEND_UNSET_FUNC_NAME")-1)) {
				__unset = reg_function
				scope.SetCeFlags(scope.GetCeFlags() | 1<<11)
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_ISSET_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__isset", g.SizeOf("ZEND_ISSET_FUNC_NAME")-1)) {
				__isset = reg_function
				scope.SetCeFlags(scope.GetCeFlags() | 1<<11)
			} else if lowercase_name.GetLen() == g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1 && !(memcmp(lowercase_name.GetVal(), "__debuginfo", g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1)) {
				__debugInfo = reg_function
			} else {
				reg_function = nil
			}
			if reg_function != nil {
				ZendCheckMagicMethodImplementation(scope, reg_function, error_type)
			}
		}
		ptr++
		count++
		ZendStringRelease(lowercase_name)
	}
	if unload != 0 {
		if scope != nil {
			_efree((*byte)(lc_class_name))
		}
		for ptr.GetFname() != nil {
			fname_len = strlen(ptr.GetFname())
			lowercase_name = ZendStringAlloc(fname_len, 0)
			ZendStrTolowerCopy(lowercase_name.GetVal(), ptr.GetFname(), fname_len)
			if ZendHashExists(target_function_table, lowercase_name) != 0 {
				ZendError(error_type, "Function registration failed - duplicate name - %s%s%s", g.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), g.Cond(scope != nil, "::", ""), ptr.GetFname())
			}
			ZendStringEfree(lowercase_name)
			ptr++
		}
		ZendUnregisterFunctions(functions, count, target_function_table)
		return FAILURE
	}
	if scope != nil {
		scope.SetConstructor(ctor)
		scope.SetDestructor(dtor)
		scope.SetClone(clone)
		scope.SetCall(__call)
		scope.SetCallstatic(__callstatic)
		scope.SetTostring(__tostring)
		scope.SetGet(__get)
		scope.SetSet(__set)
		scope.SetUnset(__unset)
		scope.SetIsset(__isset)
		scope.SetDebugInfo(__debugInfo)
		scope.SetSerializeFunc(serialize_func)
		scope.SetUnserializeFunc(unserialize_func)
		if ctor != nil {
			ctor.SetFnFlags(ctor.GetFnFlags() | 1<<28)
			if (ctor.GetFnFlags() & 1 << 4) != 0 {
				ZendError(error_type, "Constructor %s::%s() cannot be static", scope.GetName().GetVal(), ctor.GetFunctionName().GetVal())
			}
			ctor.SetFnFlags(ctor.GetFnFlags() &^ (1 << 17))
		}
		if dtor != nil {
			dtor.SetFnFlags(dtor.GetFnFlags() | 1<<29)
			if (dtor.GetFnFlags() & 1 << 4) != 0 {
				ZendError(error_type, "Destructor %s::%s() cannot be static", scope.GetName().GetVal(), dtor.GetFunctionName().GetVal())
			}
			dtor.SetFnFlags(dtor.GetFnFlags() &^ (1 << 17))
		}
		if clone != nil {
			if (clone.GetFnFlags() & 1 << 4) != 0 {
				ZendError(error_type, "%s::%s() cannot be static", scope.GetName().GetVal(), clone.GetFunctionName().GetVal())
			}
			clone.SetFnFlags(clone.GetFnFlags() &^ (1 << 17))
		}
		if __call != nil {
			if (__call.GetFnFlags() & 1 << 4) != 0 {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __call.GetFunctionName().GetVal())
			}
			__call.SetFnFlags(__call.GetFnFlags() &^ (1 << 17))
		}
		if __callstatic != nil {
			if (__callstatic.GetFnFlags() & 1 << 4) == 0 {
				ZendError(error_type, "Method %s::%s() must be static", scope.GetName().GetVal(), __callstatic.GetFunctionName().GetVal())
			}
			__callstatic.SetFnFlags(__callstatic.GetFnFlags() | 1<<4)
		}
		if __tostring != nil {
			if (__tostring.GetFnFlags() & 1 << 4) != 0 {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __tostring.GetFunctionName().GetVal())
			}
			__tostring.SetFnFlags(__tostring.GetFnFlags() &^ (1 << 17))
		}
		if __get != nil {
			if (__get.GetFnFlags() & 1 << 4) != 0 {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __get.GetFunctionName().GetVal())
			}
			__get.SetFnFlags(__get.GetFnFlags() &^ (1 << 17))
		}
		if __set != nil {
			if (__set.GetFnFlags() & 1 << 4) != 0 {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __set.GetFunctionName().GetVal())
			}
			__set.SetFnFlags(__set.GetFnFlags() &^ (1 << 17))
		}
		if __unset != nil {
			if (__unset.GetFnFlags() & 1 << 4) != 0 {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __unset.GetFunctionName().GetVal())
			}
			__unset.SetFnFlags(__unset.GetFnFlags() &^ (1 << 17))
		}
		if __isset != nil {
			if (__isset.GetFnFlags() & 1 << 4) != 0 {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __isset.GetFunctionName().GetVal())
			}
			__isset.SetFnFlags(__isset.GetFnFlags() &^ (1 << 17))
		}
		if __debugInfo != nil {
			if (__debugInfo.GetFnFlags() & 1 << 4) != 0 {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __debugInfo.GetFunctionName().GetVal())
			}
		}
		if ctor != nil && (ctor.GetFnFlags()&1<<13) != 0 {
			ZendErrorNoreturn(1<<4, "Constructor %s::%s() cannot declare a return type", scope.GetName().GetVal(), ctor.GetFunctionName().GetVal())
		}
		if dtor != nil && (dtor.GetFnFlags()&1<<13) != 0 {
			ZendErrorNoreturn(1<<4, "Destructor %s::%s() cannot declare a return type", scope.GetName().GetVal(), dtor.GetFunctionName().GetVal())
		}
		if clone != nil && (clone.GetFnFlags()&1<<13) != 0 {
			ZendErrorNoreturn(1<<4, "%s::%s() cannot declare a return type", scope.GetName().GetVal(), clone.GetFunctionName().GetVal())
		}
		_efree((*byte)(lc_class_name))
	}
	return SUCCESS
}

/* }}} */

func ZendUnregisterFunctions(functions *ZendFunctionEntry, count int, function_table *HashTable) {
	var ptr *ZendFunctionEntry = functions
	var i int = 0
	var target_function_table *HashTable = function_table
	var lowercase_name *ZendString
	var fname_len int
	if target_function_table == nil {
		target_function_table = CG.GetFunctionTable()
	}
	for ptr.GetFname() != nil {
		if count != -1 && i >= count {
			break
		}
		fname_len = strlen(ptr.GetFname())
		lowercase_name = ZendStringAlloc(fname_len, 0)
		ZendStrTolowerCopy(lowercase_name.GetVal(), ptr.GetFname(), fname_len)
		ZendHashDel(target_function_table, lowercase_name)
		ZendStringEfree(lowercase_name)
		ptr++
		i++
	}
}

/* }}} */

func ZendStartupModule(module *ZendModuleEntry) int {
	if g.Assign(&module, ZendRegisterInternalModule(module)) != nil && ZendStartupModuleEx(module) == SUCCESS {
		return SUCCESS
	}
	return FAILURE
}

/* }}} */

func ZendGetModuleStarted(module_name *byte) int {
	var module *ZendModuleEntry
	module = ZendHashStrFindPtr(&ModuleRegistry, module_name, strlen(module_name))
	if module != nil && module.GetModuleStarted() != 0 {
		return SUCCESS
	} else {
		return FAILURE
	}
}

/* }}} */

func CleanModuleClass(el *Zval, arg any) int {
	var ce *ZendClassEntry = (*ZendClassEntry)(el.GetValue().GetPtr())
	var module_number int = *((*int)(arg))
	if ce.GetType() == 1 && ce.GetModule().GetModuleNumber() == module_number {
		return 1 << 0
	} else {
		return 0
	}
}

/* }}} */

func CleanModuleClasses(module_number int) {
	ZendHashApplyWithArgument(EG.GetClassTable(), CleanModuleClass, any(&module_number))
}

/* }}} */

func ModuleDestructor(module *ZendModuleEntry) {
	if module.GetType() == 2 {
		ZendCleanModuleRsrcDtors(module.GetModuleNumber())
		CleanModuleConstants(module.GetModuleNumber())
		CleanModuleClasses(module.GetModuleNumber())
	}
	if module.GetModuleStarted() != 0 && module.GetModuleShutdownFunc() != nil {
		module.GetModuleShutdownFunc()(module.GetType(), module.GetModuleNumber())
	}
	if module.GetModuleStarted() != 0 && module.GetModuleShutdownFunc() == nil && module.GetType() == 2 {
		ZendUnregisterIniEntries(module.GetModuleNumber())
	}

	/* Deinitilaise module globals */

	if module.GetGlobalsSize() != 0 {
		if module.GetGlobalsDtor() != nil {
			module.GetGlobalsDtor()(module.GetGlobalsPtr())
		}
	}
	module.SetModuleStarted(0)
	if module.GetType() == 2 && module.GetFunctions() != nil {
		ZendUnregisterFunctions(module.GetFunctions(), -1, nil)
	}
	if module.GetHandle() && !(getenv("ZEND_DONT_UNLOAD_MODULES")) {
		dlclose(module.GetHandle())
	}
}

/* }}} */

func ZendActivateModules() {
	var p **ZendModuleEntry = ModuleRequestStartupHandlers
	for (*p) != nil {
		var module *ZendModuleEntry = *p
		if module.GetRequestStartupFunc()(module.GetType(), module.GetModuleNumber()) == FAILURE {
			ZendError(1<<1, "request_startup() for %s module failed", module.GetName())
			exit(1)
		}
		p++
	}
}

/* }}} */

func ZendDeactivateModules() {
	EG.SetCurrentExecuteData(nil)
	var __orig_bailout *sigjmp_buf = EG.GetBailout()
	var __bailout sigjmp_buf
	EG.SetBailout(&__bailout)
	if sigsetjmp(__bailout, 0) == 0 {
		if EG.GetFullTablesCleanup() != 0 {
			var module *ZendModuleEntry
			for {
				var __ht *HashTable = &ModuleRegistry
				var _idx uint32 = __ht.GetNNumUsed()
				var _p *Bucket = __ht.GetArData() + _idx
				var _z *Zval
				for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
					_p--
					_z = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					module = _z.GetValue().GetPtr()
					if module.GetRequestShutdownFunc() != nil {
						module.GetRequestShutdownFunc()(module.GetType(), module.GetModuleNumber())
					}
				}
				break
			}
		} else {
			var p **ZendModuleEntry = ModuleRequestShutdownHandlers
			for (*p) != nil {
				var module *ZendModuleEntry = *p
				module.GetRequestShutdownFunc()(module.GetType(), module.GetModuleNumber())
				p++
			}
		}
	}
	EG.SetBailout(__orig_bailout)
}

/* }}} */

func ZendCleanupInternalClasses() {
	var p **ZendClassEntry = ClassCleanupHandlers
	for (*p) != nil {
		ZendCleanupInternalClassData(*p)
		p++
	}
}

/* }}} */

func ZendPostDeactivateModules() {
	if EG.GetFullTablesCleanup() != 0 {
		var module *ZendModuleEntry
		var zv *Zval
		var key *ZendString
		for {
			var __ht *HashTable = &ModuleRegistry
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				module = _z.GetValue().GetPtr()
				if module.GetPostDeactivateFunc() != nil {
					module.GetPostDeactivateFunc()()
				}
			}
			break
		}
		for {
			var __ht *HashTable = &ModuleRegistry
			var _idx uint32 = __ht.GetNNumUsed()
			var _p *Bucket = __ht.GetArData() + _idx
			var _z *Zval
			for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
				_p--
				_z = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				key = _p.GetKey()
				zv = _z
				module = zv.GetValue().GetPtr()
				if module.GetType() != 2 {
					break
				}
				ModuleDestructor(module)
				Free(module)
				ZendStringReleaseEx(key, 0)
				ZEND_HASH_FOREACH_END_DEL_ITEM(__ht, _idx, _p)
			}
			__ht.SetNNumUsed(_idx)
			break
		}
	} else {
		var p **ZendModuleEntry = ModulePostDeactivateHandlers
		for (*p) != nil {
			var module *ZendModuleEntry = *p
			module.GetPostDeactivateFunc()()
			p++
		}
	}
}

/* }}} */

func ZendNextFreeModule() int {
	return &ModuleRegistry.GetNNumOfElements() + 1
}

/* }}} */

func DoRegisterInternalClass(orig_class_entry *ZendClassEntry, ce_flags uint32) *ZendClassEntry {
	var class_entry *ZendClassEntry = Malloc(g.SizeOf("zend_class_entry"))
	var lowercase_name *ZendString
	*class_entry = *orig_class_entry
	class_entry.SetType(1)
	ZendInitializeClassData(class_entry, 0)
	class_entry.SetCeFlags(ce_flags | 1<<12 | 1<<3 | 1<<19 | 1<<20)
	class_entry.SetModule(EG.GetCurrentModule())
	if class_entry.GetBuiltinFunctions() != nil {
		ZendRegisterFunctions(class_entry, class_entry.GetBuiltinFunctions(), &class_entry.function_table, EG.GetCurrentModule().GetType())
	}
	lowercase_name = ZendStringTolowerEx(orig_class_entry.GetName(), EG.GetCurrentModule().GetType() == 1)
	lowercase_name = ZendNewInternedString(lowercase_name)
	ZendHashUpdatePtr(CG.GetClassTable(), lowercase_name, class_entry)
	ZendStringReleaseEx(lowercase_name, 1)
	return class_entry
}

/* }}} */

func ZendRegisterInternalClassEx(class_entry *ZendClassEntry, parent_ce *ZendClassEntry) *ZendClassEntry {
	var register_class *ZendClassEntry
	register_class = ZendRegisterInternalClass(class_entry)
	if parent_ce != nil {
		ZendDoInheritanceEx(register_class, parent_ce, 0)
		ZendBuildPropertiesInfoTable(register_class)
	}
	return register_class
}

/* }}} */

func ZendClassImplements(class_entry *ZendClassEntry, num_interfaces int, _ ...any) {
	var interface_entry *ZendClassEntry
	var interface_list va_list
	va_start(interface_list, num_interfaces)
	for g.PostDec(&num_interfaces) {
		interface_entry = __va_arg(interface_list, (*ZendClassEntry)(_))
		ZendDoImplementInterface(class_entry, interface_entry)
	}
	va_end(interface_list)
}

/* }}} */

func ZendRegisterInternalClass(orig_class_entry *ZendClassEntry) *ZendClassEntry {
	return DoRegisterInternalClass(orig_class_entry, 0)
}

/* }}} */

func ZendRegisterInternalInterface(orig_class_entry *ZendClassEntry) *ZendClassEntry {
	return DoRegisterInternalClass(orig_class_entry, 1<<0)
}

/* }}} */

func ZendRegisterClassAliasEx(name *byte, name_len int, ce *ZendClassEntry, persistent int) int {
	var lcname *ZendString
	var zv Zval
	var ret *Zval

	/* TODO: Move this out of here in 7.4. */

	if persistent != 0 && EG.GetCurrentModule() != nil && EG.GetCurrentModule().GetType() == 2 {
		persistent = 0
	}
	if name[0] == '\\' {
		lcname = ZendStringAlloc(name_len-1, persistent)
		ZendStrTolowerCopy(lcname.GetVal(), name+1, name_len-1)
	} else {
		lcname = ZendStringAlloc(name_len, persistent)
		ZendStrTolowerCopy(lcname.GetVal(), name, name_len)
	}
	ZendAssertValidClassName(lcname)
	lcname = ZendNewInternedString(lcname)
	&zv.GetValue().SetPtr(ce)
	&zv.SetTypeInfo(15)
	ret = ZendHashAdd(CG.GetClassTable(), lcname, &zv)
	ZendStringReleaseEx(lcname, 0)
	if ret != nil {
		if (ce.GetCeFlags() & 1 << 7) == 0 {
			ce.GetRefcount()++
		}
		return SUCCESS
	}
	return FAILURE
}

/* }}} */

func ZendSetHashSymbol(symbol *Zval, name *byte, name_length int, is_ref ZendBool, num_symbol_tables int, _ ...any) int {
	var symbol_table *HashTable
	var symbol_table_list va_list
	if num_symbol_tables <= 0 {
		return FAILURE
	}
	if is_ref != 0 {
		var __zv *Zval = symbol
		if __zv.GetType() != 10 {
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 1)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = __zv
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			__zv.GetValue().SetRef(_ref)
			__zv.SetTypeInfo(10 | 1<<0<<8)
		}
	}
	va_start(symbol_table_list, num_symbol_tables)
	for g.PostDec(&num_symbol_tables) > 0 {
		symbol_table = __va_arg(symbol_table_list, (*HashTable)(_))
		ZendHashStrUpdate(symbol_table, name, name_length, symbol)
		if symbol.GetTypeFlags() != 0 {
			ZvalAddrefP(symbol)
		}
	}
	va_end(symbol_table_list)
	return SUCCESS
}

/* }}} */

func ZifDisplayDisabledFunction(execute_data *ZendExecuteData, return_value *Zval) {
	ZendError(1<<1, "%s() has been disabled for security reasons", GetActiveFunctionName())
}

/* }}} */

func ZendDisableFunction(function_name *byte, function_name_length int) int {
	var func_ *ZendInternalFunction
	if g.Assign(&func_, ZendHashStrFindPtr(CG.GetFunctionTable(), function_name, function_name_length)) {
		ZendFreeInternalArgInfo(func_)
		func_.SetFnFlags(func_.GetFnFlags() &^ (1<<14 | 1<<8 | 1<<13))
		func_.SetNumArgs(0)
		func_.SetArgInfo(nil)
		func_.SetHandler(ZifDisplayDisabledFunction)
		return SUCCESS
	}
	return FAILURE
}

/* }}} */

func DisplayDisabledClass(class_type *ZendClassEntry) *ZendObject {
	var intern *ZendObject
	intern = ZendObjectsNew(class_type)

	/* Initialize default properties */

	if class_type.GetDefaultPropertiesCount() != 0 {
		var p *Zval = intern.GetPropertiesTable()
		var end *Zval = p + class_type.GetDefaultPropertiesCount()
		for {
			p.SetTypeInfo(0)
			p++
			if p == end {
				break
			}
		}
	}
	ZendError(1<<1, "%s() has been disabled for security reasons", class_type.GetName().GetVal())
	return intern
}

/* }}} */

var DisabledClassNew []ZendFunctionEntry = []ZendFunctionEntry{{nil, nil, nil, 0, 0}}

func ZendDisableClass(class_name *byte, class_name_length int) int {
	var disabled_class *ZendClassEntry
	var key *ZendString
	var fn *ZendFunction
	key = ZendStringAlloc(class_name_length, 0)
	ZendStrTolowerCopy(key.GetVal(), class_name, class_name_length)
	disabled_class = ZendHashFindPtr(CG.GetClassTable(), key)
	ZendStringReleaseEx(key, 0)
	if disabled_class == nil {
		return FAILURE
	}
	disabled_class.SetConstructor(nil)
	disabled_class.SetDestructor(nil)
	disabled_class.SetClone(nil)
	disabled_class.SetSerialize(nil)
	disabled_class.SetUnserialize(nil)
	disabled_class.create_object = nil
	disabled_class.SetGetStaticMethod(nil)
	disabled_class.SetCall(nil)
	disabled_class.SetCallstatic(nil)
	disabled_class.SetTostring(nil)
	disabled_class.SetGet(nil)
	disabled_class.SetSet(nil)
	disabled_class.SetUnset(nil)
	disabled_class.SetIsset(nil)
	disabled_class.SetDebugInfo(nil)
	disabled_class.SetSerializeFunc(nil)
	disabled_class.SetUnserializeFunc(nil)
	disabled_class.parent = nil
	disabled_class.SetNumInterfaces(0)
	disabled_class.SetTraitNames(nil)
	disabled_class.SetNumTraits(0)
	disabled_class.SetTraitAliases(nil)
	disabled_class.SetTraitPrecedences(nil)
	disabled_class.interfaces = nil
	disabled_class.SetGetIterator(nil)
	disabled_class.SetIteratorFuncsPtr(nil)
	disabled_class.SetModule(nil)
	disabled_class.SetBuiltinFunctions(DisabledClassNew)
	disabled_class.create_object = DisplayDisabledClass
	for {
		var __ht *HashTable = &disabled_class.function_table
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			fn = _z.GetValue().GetPtr()
			if (fn.GetFnFlags()&(1<<13|1<<8)) != 0 && fn.GetScope() == disabled_class {
				ZendFreeInternalArgInfo(&fn.internal_function)
			}
		}
		break
	}
	ZendHashClean(&disabled_class.function_table)
	return SUCCESS
}

/* }}} */

func ZendIsCallableCheckClass(name *ZendString, scope *ZendClassEntry, fcc *ZendFcallInfoCache, strict_class *int, error **byte) int {
	var ret int = 0
	var ce *ZendClassEntry
	var name_len int = name.GetLen()
	var lcname *ZendString
	lcname = (*ZendString)(_emalloc(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + name_len + 1 + (8-1) & ^(8-1)))
	ZendGcSetRefcount(&lcname.gc, 1)
	lcname.GetGc().SetTypeInfo(6)
	lcname.SetH(0)
	lcname.SetLen(name_len)
	ZendStrTolowerCopy(lcname.GetVal(), name.GetVal(), name_len)
	*strict_class = 0
	if lcname.GetLen() == g.SizeOf("\"self\"")-1 && !(memcmp(lcname.GetVal(), "self", g.SizeOf("\"self\"")-1)) {
		if scope == nil {
			if error != nil {
				*error = _estrdup("cannot access self:: when no class scope is active")
			}
		} else {
			fcc.SetCalledScope(ZendGetCalledScope(EG.GetCurrentExecuteData()))
			if fcc.GetCalledScope() == nil || InstanceofFunction(fcc.GetCalledScope(), scope) == 0 {
				fcc.SetCalledScope(scope)
			}
			fcc.SetCallingScope(scope)
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(EG.GetCurrentExecuteData()))
			}
			ret = 1
		}
	} else if lcname.GetLen() == g.SizeOf("\"parent\"")-1 && !(memcmp(lcname.GetVal(), "parent", g.SizeOf("\"parent\"")-1)) {
		if scope == nil {
			if error != nil {
				*error = _estrdup("cannot access parent:: when no class scope is active")
			}
		} else if !(scope.parent) {
			if error != nil {
				*error = _estrdup("cannot access parent:: when current class scope has no parent")
			}
		} else {
			fcc.SetCalledScope(ZendGetCalledScope(EG.GetCurrentExecuteData()))
			if fcc.GetCalledScope() == nil || InstanceofFunction(fcc.GetCalledScope(), scope.parent) == 0 {
				fcc.SetCalledScope(scope.parent)
			}
			fcc.SetCallingScope(scope.parent)
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(EG.GetCurrentExecuteData()))
			}
			*strict_class = 1
			ret = 1
		}
	} else if lcname.GetLen() == g.SizeOf("\"static\"")-1 && !(memcmp(lcname.GetVal(), "static", g.SizeOf("\"static\"")-1)) {
		var called_scope *ZendClassEntry = ZendGetCalledScope(EG.GetCurrentExecuteData())
		if called_scope == nil {
			if error != nil {
				*error = _estrdup("cannot access static:: when no class scope is active")
			}
		} else {
			fcc.SetCalledScope(called_scope)
			fcc.SetCallingScope(called_scope)
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(EG.GetCurrentExecuteData()))
			}
			*strict_class = 1
			ret = 1
		}
	} else if g.Assign(&ce, ZendLookupClass(name)) != nil {
		var scope *ZendClassEntry
		var ex *ZendExecuteData = EG.GetCurrentExecuteData()
		for ex != nil && (ex.GetFunc() == nil || (ex.GetFunc().GetType()&1) != 0) {
			ex = ex.GetPrevExecuteData()
		}
		if ex != nil {
			scope = ex.GetFunc().GetScope()
		} else {
			scope = nil
		}
		fcc.SetCallingScope(ce)
		if scope != nil && fcc.GetObject() == nil {
			var object *ZendObject = ZendGetThisObject(EG.GetCurrentExecuteData())
			if object != nil && InstanceofFunction(object.GetCe(), scope) != 0 && InstanceofFunction(scope, ce) != 0 {
				fcc.SetObject(object)
				fcc.SetCalledScope(object.GetCe())
			} else {
				fcc.SetCalledScope(ce)
			}
		} else {
			if fcc.GetObject() != nil {
				fcc.SetCalledScope(fcc.GetObject().GetCe())
			} else {
				fcc.SetCalledScope(ce)
			}
		}
		*strict_class = 1
		ret = 1
	} else {
		if error != nil {
			ZendSpprintf(error, 0, "class '%.*s' not found", int(name_len), name.GetVal())
		}
	}
	_efree(lcname)
	return ret
}

/* }}} */

func ZendReleaseFcallInfoCache(fcc *ZendFcallInfoCache) {
	if fcc.GetFunctionHandler() != nil && ((fcc.GetFunctionHandler().GetFnFlags()&1<<18) != 0 || fcc.GetFunctionHandler().GetType() == 5 || fcc.GetFunctionHandler().GetType() == 3) {
		if fcc.GetFunctionHandler().GetType() != 3 && fcc.GetFunctionHandler().GetFunctionName() != nil {
			ZendStringReleaseEx(fcc.GetFunctionHandler().GetFunctionName(), 0)
		}
		if fcc.GetFunctionHandler() == &EG.trampoline {
			EG.GetTrampoline().SetFunctionName(nil)
		} else {
			_efree(fcc.GetFunctionHandler())
		}
	}
	fcc.SetFunctionHandler(nil)
}
func ZendIsCallableCheckFunc(check_flags int, callable *Zval, fcc *ZendFcallInfoCache, strict_class int, error **byte) int {
	var ce_org *ZendClassEntry = fcc.GetCallingScope()
	var retval int = 0
	var mname *ZendString
	var cname *ZendString
	var lmname *ZendString
	var colon *byte
	var clen int
	var ftable *HashTable
	var call_via_handler int = 0
	var scope *ZendClassEntry
	var zv *Zval
	fcc.SetCallingScope(nil)
	if ce_org == nil {
		var func_ *ZendFunction
		var lmname *ZendString

		/* Check if function with given name exists.
		 * This may be a compound name that includes namespace name */

		if callable.GetValue().GetStr().GetVal()[0] == '\\' {

			/* Skip leading \ */

			lmname = (*ZendString)(_emalloc(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + callable.GetValue().GetStr().GetLen() - 1 + 1 + (8-1) & ^(8-1)))
			ZendGcSetRefcount(&lmname.gc, 1)
			lmname.GetGc().SetTypeInfo(6)
			lmname.SetH(0)
			lmname.SetLen(callable.GetValue().GetStr().GetLen() - 1)
			ZendStrTolowerCopy(lmname.GetVal(), callable.GetValue().GetStr().GetVal()+1, callable.GetValue().GetStr().GetLen()-1)
			func_ = ZendFetchFunction(lmname)
			_efree(lmname)
		} else {
			lmname = callable.GetValue().GetStr()
			func_ = ZendFetchFunction(lmname)
			if func_ == nil {
				lmname = (*ZendString)(_emalloc(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil)) + callable.GetValue().GetStr().GetLen() + 1 + (8-1) & ^(8-1)))
				ZendGcSetRefcount(&lmname.gc, 1)
				lmname.GetGc().SetTypeInfo(6)
				lmname.SetH(0)
				lmname.SetLen(callable.GetValue().GetStr().GetLen())
				ZendStrTolowerCopy(lmname.GetVal(), callable.GetValue().GetStr().GetVal(), callable.GetValue().GetStr().GetLen())
				func_ = ZendFetchFunction(lmname)
				_efree(lmname)
			}
		}
		if func_ != nil {
			fcc.SetFunctionHandler(func_)
			return 1
		}
	}

	/* Split name into class/namespace and method/function names */

	if g.Assign(&colon, ZendMemrchr(callable.GetValue().GetStr().GetVal(), ':', callable.GetValue().GetStr().GetLen())) != nil && colon > callable.GetValue().GetStr().GetVal() && (*(colon - 1)) == ':' {
		var mlen int
		colon--
		clen = colon - callable.GetValue().GetStr().GetVal()
		mlen = callable.GetValue().GetStr().GetLen() - clen - 2
		if colon == callable.GetValue().GetStr().GetVal() {
			if error != nil {
				*error = _estrdup("invalid function name")
			}
			return 0
		}

		/* This is a compound name.
		 * Try to fetch class and then find static method. */

		if ce_org != nil {
			scope = ce_org
		} else {
			scope = ZendGetExecutedScope()
		}
		cname = ZendStringInit(callable.GetValue().GetStr().GetVal(), clen, 0)
		if ZendIsCallableCheckClass(cname, scope, fcc, &strict_class, error) == 0 {
			ZendStringReleaseEx(cname, 0)
			return 0
		}
		ZendStringReleaseEx(cname, 0)
		ftable = &fcc.calling_scope.GetFunctionTable()
		if ce_org != nil && InstanceofFunction(ce_org, fcc.GetCallingScope()) == 0 {
			if error != nil {
				ZendSpprintf(error, 0, "class '%s' is not a subclass of '%s'", ce_org.GetName().GetVal(), fcc.GetCallingScope().GetName().GetVal())
			}
			return 0
		}
		mname = ZendStringInit(callable.GetValue().GetStr().GetVal()+clen+2, mlen, 0)
	} else if ce_org != nil {

		/* Try to fetch find static method of given class. */

		mname = callable.GetValue().GetStr()
		ZendStringAddref(mname)
		ftable = &ce_org.function_table
		fcc.SetCallingScope(ce_org)
	} else {

		/* We already checked for plain function before. */

		if error != nil && (check_flags&1<<3) == 0 {
			ZendSpprintf(error, 0, "function '%s' not found or invalid function name", callable.GetValue().GetStr().GetVal())
		}
		return 0
	}
	lmname = ZendStringTolowerEx(mname, 0)
	if strict_class != 0 && fcc.GetCallingScope() != nil && (lmname.GetLen() == g.SizeOf("ZEND_CONSTRUCTOR_FUNC_NAME")-1 && !(memcmp(lmname.GetVal(), "__construct", g.SizeOf("ZEND_CONSTRUCTOR_FUNC_NAME")-1))) {
		fcc.SetFunctionHandler(fcc.GetCallingScope().GetConstructor())
		if fcc.GetFunctionHandler() != nil {
			retval = 1
		}
	} else if g.Assign(&zv, ZendHashFind(ftable, lmname)) != nil {
		fcc.SetFunctionHandler(zv.GetValue().GetPtr())
		retval = 1
		if (fcc.GetFunctionHandler().GetOpArray().GetFnFlags()&1<<3) != 0 && strict_class == 0 {
			scope = ZendGetExecutedScope()
			if scope != nil && InstanceofFunction(fcc.GetFunctionHandler().GetScope(), scope) != 0 {
				zv = ZendHashFind(&scope.function_table, lmname)
				if zv != nil {
					var priv_fbc *ZendFunction = zv.GetValue().GetPtr()
					if (priv_fbc.GetFnFlags()&1<<2) != 0 && priv_fbc.GetScope() == scope {
						fcc.SetFunctionHandler(priv_fbc)
					}
				}
			}
		}
		if (fcc.GetFunctionHandler().GetFnFlags()&1<<0) == 0 && (check_flags&1<<1) == 0 && (fcc.GetCallingScope() != nil && (fcc.GetObject() != nil && fcc.GetCallingScope().GetCall() != nil || fcc.GetObject() == nil && fcc.GetCallingScope().GetCallstatic() != nil)) {
			scope = ZendGetExecutedScope()
			if fcc.GetFunctionHandler().GetScope() != scope {
				if (fcc.GetFunctionHandler().GetFnFlags()&1<<2) != 0 || ZendCheckProtected(g.CondF(fcc.GetFunctionHandler().GetPrototype() != nil, func() *ZendClassEntry { return fcc.GetFunctionHandler().GetPrototype().GetScope() }, func() *ZendClassEntry { return fcc.GetFunctionHandler().GetScope() }), scope) == 0 {
					retval = 0
					fcc.SetFunctionHandler(nil)
					goto get_function_via_handler
				}
			}
		}
	} else {
	get_function_via_handler:
		if fcc.GetObject() != nil && fcc.GetCallingScope() == ce_org {
			if strict_class != 0 && ce_org.GetCall() != nil {
				fcc.SetFunctionHandler(ZendGetCallTrampolineFunc(ce_org, mname, 0))
				call_via_handler = 1
				retval = 1
			} else {
				fcc.SetFunctionHandler(fcc.GetObject().GetHandlers().GetGetMethod()(&fcc.object, mname, nil))
				if fcc.GetFunctionHandler() != nil {
					if strict_class != 0 && (fcc.GetFunctionHandler().GetScope() == nil || InstanceofFunction(ce_org, fcc.GetFunctionHandler().GetScope()) == 0) {
						ZendReleaseFcallInfoCache(fcc)
					} else {
						retval = 1
						call_via_handler = (fcc.GetFunctionHandler().GetFnFlags() & 1 << 18) != 0
					}
				}
			}
		} else if fcc.GetCallingScope() != nil {
			if fcc.GetCallingScope().GetGetStaticMethod() != nil {
				fcc.SetFunctionHandler(fcc.GetCallingScope().GetGetStaticMethod()(fcc.GetCallingScope(), mname))
			} else {
				fcc.SetFunctionHandler(ZendStdGetStaticMethod(fcc.GetCallingScope(), mname, nil))
			}
			if fcc.GetFunctionHandler() != nil {
				retval = 1
				call_via_handler = (fcc.GetFunctionHandler().GetFnFlags() & 1 << 18) != 0
				if call_via_handler != 0 && fcc.GetObject() == nil {
					var object *ZendObject = ZendGetThisObject(EG.GetCurrentExecuteData())
					if object != nil && InstanceofFunction(object.GetCe(), fcc.GetCallingScope()) != 0 {
						fcc.SetObject(object)
					}
				}
			}
		}
	}
	if retval != 0 {
		if fcc.GetCallingScope() != nil && call_via_handler == 0 {
			if (fcc.GetFunctionHandler().GetFnFlags() & 1 << 6) != 0 {
				retval = 0
				if error != nil {
					ZendSpprintf(error, 0, "cannot call abstract method %s::%s()", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal())
				}
			} else if fcc.GetObject() == nil && (fcc.GetFunctionHandler().GetFnFlags()&1<<4) == 0 {
				var severity int
				var verb *byte
				if (fcc.GetFunctionHandler().GetFnFlags() & 1 << 17) != 0 {
					severity = 1 << 13
					verb = "should not"
				} else {

					/* An internal function assumes $this is present and won't check that. So PHP would crash by allowing the call. */

					severity = 1 << 0
					verb = "cannot"
				}
				if (check_flags & 1 << 2) != 0 {
					retval = 0
				}
				if error != nil {
					ZendSpprintf(error, 0, "non-static method %s::%s() %s be called statically", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal(), verb)
					if severity != 1<<13 {
						retval = 0
					}
				} else if retval != 0 {
					if severity == 1<<0 {
						ZendThrowError(nil, "Non-static method %s::%s() %s be called statically", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal(), verb)
					} else {
						ZendError(severity, "Non-static method %s::%s() %s be called statically", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal(), verb)
					}
				}
			}
			if retval != 0 && (fcc.GetFunctionHandler().GetFnFlags()&1<<0) == 0 && (check_flags&1<<1) == 0 {
				scope = ZendGetExecutedScope()
				if fcc.GetFunctionHandler().GetScope() != scope {
					if (fcc.GetFunctionHandler().GetFnFlags()&1<<2) != 0 || ZendCheckProtected(g.CondF(fcc.GetFunctionHandler().GetPrototype() != nil, func() *ZendClassEntry { return fcc.GetFunctionHandler().GetPrototype().GetScope() }, func() *ZendClassEntry { return fcc.GetFunctionHandler().GetScope() }), scope) == 0 {
						if error != nil {
							if (*error) != nil {
								_efree(*error)
							}
							ZendSpprintf(error, 0, "cannot access %s method %s::%s()", ZendVisibilityString(fcc.GetFunctionHandler().GetFnFlags()), fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal())
						}
						retval = 0
					}
				}
			}
		}
	} else if error != nil && (check_flags&1<<3) == 0 {
		if fcc.GetCallingScope() != nil {
			if error != nil {
				ZendSpprintf(error, 0, "class '%s' does not have a method '%s'", fcc.GetCallingScope().GetName().GetVal(), mname.GetVal())
			}
		} else {
			if error != nil {
				ZendSpprintf(error, 0, "function '%s' does not exist", mname.GetVal())
			}
		}
	}
	ZendStringReleaseEx(lmname, 0)
	ZendStringReleaseEx(mname, 0)
	if fcc.GetObject() != nil {
		fcc.SetCalledScope(fcc.GetObject().GetCe())
		if fcc.GetFunctionHandler() != nil && (fcc.GetFunctionHandler().GetFnFlags()&1<<4) != 0 {
			fcc.SetObject(nil)
		}
	}
	return retval
}

/* }}} */

func ZendCreateMethodString(class_name *ZendString, method_name *ZendString) *ZendString {
	var callable_name *ZendString = ZendStringAlloc(class_name.GetLen()+method_name.GetLen()+g.SizeOf("\"::\"")-1, 0)
	var ptr *byte = callable_name.GetVal()
	memcpy(ptr, class_name.GetVal(), class_name.GetLen())
	ptr += class_name.GetLen()
	memcpy(ptr, "::", g.SizeOf("\"::\"")-1)
	ptr += g.SizeOf("\"::\"") - 1
	memcpy(ptr, method_name.GetVal(), method_name.GetLen()+1)
	return callable_name
}
func ZendGetCallableNameEx(callable *Zval, object *ZendObject) *ZendString {
try_again:
	switch callable.GetType() {
	case 6:
		if object != nil {
			return ZendCreateMethodString(object.GetCe().GetName(), callable.GetValue().GetStr())
		}
		return ZendStringCopy(callable.GetValue().GetStr())
	case 7:
		var method *Zval = nil
		var obj *Zval = nil
		if callable.GetValue().GetArr().GetNNumOfElements() == 2 {
			obj = ZendHashIndexFindDeref(callable.GetValue().GetArr(), 0)
			method = ZendHashIndexFindDeref(callable.GetValue().GetArr(), 1)
		}
		if obj == nil || method == nil || method.GetType() != 6 {
			return ZendKnownStrings[ZEND_STR_ARRAY_CAPITALIZED]
		}
		if obj.GetType() == 6 {
			return ZendCreateMethodString(obj.GetValue().GetStr(), method.GetValue().GetStr())
		} else if obj.GetType() == 8 {
			return ZendCreateMethodString(obj.GetValue().GetObj().GetCe().GetName(), method.GetValue().GetStr())
		} else {
			return ZendKnownStrings[ZEND_STR_ARRAY_CAPITALIZED]
		}
	case 8:
		var calling_scope *ZendClassEntry
		var fptr *ZendFunction
		var object *ZendObject
		if callable.GetValue().GetObj().GetHandlers().GetGetClosure() != nil && callable.GetValue().GetObj().GetHandlers().GetGetClosure()(callable, &calling_scope, &fptr, &object) == SUCCESS {
			var ce *ZendClassEntry = callable.GetValue().GetObj().GetCe()
			var callable_name *ZendString = ZendStringAlloc(ce.GetName().GetLen()+g.SizeOf("\"::__invoke\"")-1, 0)
			memcpy(callable_name.GetVal(), ce.GetName().GetVal(), ce.GetName().GetLen())
			memcpy(callable_name.GetVal()+ce.GetName().GetLen(), "::__invoke", g.SizeOf("\"::__invoke\""))
			return callable_name
		}
		return ZvalGetString(callable)
	case 10:
		callable = &(*callable).value.GetRef().GetVal()
		goto try_again
	default:
		return ZvalGetStringFunc(callable)
	}
}

/* }}} */

func ZendGetCallableName(callable *Zval) *ZendString {
	return ZendGetCallableNameEx(callable, nil)
}

/* }}} */

func ZendIsCallableImpl(callable *Zval, object *ZendObject, check_flags uint32, fcc *ZendFcallInfoCache, error **byte) ZendBool {
	var ret ZendBool
	var fcc_local ZendFcallInfoCache
	var strict_class int = 0
	if fcc == nil {
		fcc = &fcc_local
	}
	if error != nil {
		*error = nil
	}
	fcc.SetCallingScope(nil)
	fcc.SetCalledScope(nil)
	fcc.SetFunctionHandler(nil)
	fcc.SetObject(nil)
again:
	switch callable.GetType() {
	case 6:
		if object != nil {
			fcc.SetObject(object)
			fcc.SetCallingScope(object.GetCe())
		}
		if (check_flags & 1 << 0) != 0 {
			fcc.SetCalledScope(fcc.GetCallingScope())
			return 1
		}
	check_func:
		ret = ZendIsCallableCheckFunc(check_flags, callable, fcc, strict_class, error)
		if fcc == &fcc_local {
			ZendReleaseFcallInfoCache(fcc)
		}
		return ret
	case 7:
		var method *Zval = nil
		var obj *Zval = nil
		if callable.GetValue().GetArr().GetNNumOfElements() == 2 {
			obj = ZendHashIndexFind(callable.GetValue().GetArr(), 0)
			method = ZendHashIndexFind(callable.GetValue().GetArr(), 1)
		}
		for {
			if obj == nil || method == nil {
				break
			}
			if method.GetType() == 10 {
				method = &(*method).value.GetRef().GetVal()
			}
			if method.GetType() != 6 {
				break
			}
			if obj.GetType() == 10 {
				obj = &(*obj).value.GetRef().GetVal()
			}
			if obj.GetType() == 6 {
				if (check_flags & 1 << 0) != 0 {
					return 1
				}
				if ZendIsCallableCheckClass(obj.GetValue().GetStr(), ZendGetExecutedScope(), fcc, &strict_class, error) == 0 {
					return 0
				}
			} else if obj.GetType() == 8 {
				fcc.SetCallingScope(obj.GetValue().GetObj().GetCe())
				fcc.SetObject(obj.GetValue().GetObj())
				if (check_flags & 1 << 0) != 0 {
					fcc.SetCalledScope(fcc.GetCallingScope())
					return 1
				}
			} else {
				break
			}
			callable = method
			goto check_func
			break
		}
		if callable.GetValue().GetArr().GetNNumOfElements() == 2 {
			if obj == nil || g.CondF(obj.GetType() != 10, func() bool { return obj.GetType() != 6 && obj.GetType() != 8 }, func() bool {
				return &(*obj).value.GetRef().GetVal().u1.v.type_ != 6 && &(*obj).value.GetRef().GetVal().u1.v.type_ != 8
			}) {
				if error != nil {
					*error = _estrdup("first array member is not a valid class name or object")
				}
			} else {
				if error != nil {
					*error = _estrdup("second array member is not a valid method")
				}
			}
		} else {
			if error != nil {
				*error = _estrdup("array must have exactly two members")
			}
		}
		return 0
	case 8:
		if callable.GetValue().GetObj().GetHandlers().GetGetClosure() != nil {
			if callable.GetValue().GetObj().GetHandlers().GetGetClosure()(callable, &fcc.calling_scope, &fcc.function_handler, &fcc.object) == SUCCESS {
				fcc.SetCalledScope(fcc.GetCallingScope())
				if fcc == &fcc_local {
					ZendReleaseFcallInfoCache(fcc)
				}
				return 1
			} else {

				/* Discard exceptions thrown from Z_OBJ_HANDLER_P(callable, get_closure) */

				ZendClearException()

				/* Discard exceptions thrown from Z_OBJ_HANDLER_P(callable, get_closure) */

			}
		}
		if error != nil {
			*error = _estrdup("no array or string given")
		}
		return 0
	case 10:
		callable = &(*callable).value.GetRef().GetVal()
		goto again
	default:
		if error != nil {
			*error = _estrdup("no array or string given")
		}
		return 0
	}
}

/* }}} */

func ZendIsCallableEx(callable *Zval, object *ZendObject, check_flags uint32, callable_name **ZendString, fcc *ZendFcallInfoCache, error **byte) ZendBool {
	var ret ZendBool = ZendIsCallableImpl(callable, object, check_flags, fcc, error)
	if callable_name != nil {
		*callable_name = ZendGetCallableNameEx(callable, object)
	}
	return ret
}
func ZendIsCallable(callable *Zval, check_flags uint32, callable_name **ZendString) ZendBool {
	return ZendIsCallableEx(callable, nil, check_flags, callable_name, nil, nil)
}

/* }}} */

func ZendMakeCallable(callable *Zval, callable_name **ZendString) ZendBool {
	var fcc ZendFcallInfoCache
	if ZendIsCallableEx(callable, nil, 1<<2, callable_name, &fcc, nil) != 0 {
		if callable.GetType() == 6 && fcc.GetCallingScope() != nil {
			ZvalPtrDtorStr(callable)
			var __arr *ZendArray = _zendNewArray(0)
			var __z *Zval = callable
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
			AddNextIndexStr(callable, ZendStringCopy(fcc.GetCallingScope().GetName()))
			AddNextIndexStr(callable, ZendStringCopy(fcc.GetFunctionHandler().GetFunctionName()))
		}
		ZendReleaseFcallInfoCache(&fcc)
		return 1
	}
	return 0
}

/* }}} */

func ZendFcallInfoInit(callable *Zval, check_flags uint32, fci *ZendFcallInfo, fcc *ZendFcallInfoCache, callable_name **ZendString, error **byte) int {
	if ZendIsCallableEx(callable, nil, check_flags, callable_name, fcc, error) == 0 {
		return FAILURE
	}
	fci.SetSize(g.SizeOf("* fci"))
	fci.SetObject(fcc.GetObject())
	var _z1 *Zval = &fci.function_name
	var _z2 *Zval = callable
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	fci.SetRetval(nil)
	fci.SetParamCount(0)
	fci.SetParams(nil)
	fci.SetNoSeparation(1)
	return SUCCESS
}

/* }}} */

func ZendFcallInfoArgsClear(fci *ZendFcallInfo, free_mem int) {
	if fci.GetParams() != nil {
		var p *Zval = fci.GetParams()
		var end *Zval = p + fci.GetParamCount()
		for p != end {
			IZvalPtrDtor(p)
			p++
		}
		if free_mem != 0 {
			_efree(fci.GetParams())
			fci.SetParams(nil)
		}
	}
	fci.SetParamCount(0)
}

/* }}} */

func ZendFcallInfoArgsSave(fci *ZendFcallInfo, param_count *int, params **Zval) {
	*param_count = fci.GetParamCount()
	*params = fci.GetParams()
	fci.SetParamCount(0)
	fci.SetParams(nil)
}

/* }}} */

func ZendFcallInfoArgsRestore(fci *ZendFcallInfo, param_count int, params *Zval) {
	ZendFcallInfoArgsClear(fci, 1)
	fci.SetParamCount(param_count)
	fci.SetParams(params)
}

/* }}} */

func ZendFcallInfoArgsEx(fci *ZendFcallInfo, func_ *ZendFunction, args *Zval) int {
	var arg *Zval
	var params *Zval
	var n uint32 = 1
	ZendFcallInfoArgsClear(fci, !args)
	if args == nil {
		return SUCCESS
	}
	if args.GetType() != 7 {
		return FAILURE
	}
	fci.SetParamCount(args.GetValue().GetArr().GetNNumOfElements())
	params = (*Zval)(_erealloc(fci.GetParams(), fci.GetParamCount()*g.SizeOf("zval")))
	fci.SetParams(params)
	for {
		var __ht *HashTable = args.GetValue().GetArr()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			arg = _z
			if func_ != nil && arg.GetType() != 10 && ZendCheckArgSendType(func_, n, 1|2) != 0 {
				var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
				ZendGcSetRefcount(&_ref.gc, 1)
				_ref.GetGc().SetTypeInfo(10)
				var _z1 *Zval = &_ref.val
				var _z2 *Zval = arg
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				_ref.GetSources().SetPtr(nil)
				params.GetValue().SetRef(_ref)
				params.SetTypeInfo(10 | 1<<0<<8)
				if arg.GetTypeFlags() != 0 {
					ZvalAddrefP(arg)
				}
			} else {
				var _z1 *Zval = params
				var _z2 *Zval = arg
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
			}
			params++
			n++
		}
		break
	}
	return SUCCESS
}

/* }}} */

func ZendFcallInfoArgs(fci *ZendFcallInfo, args *Zval) int {
	return ZendFcallInfoArgsEx(fci, nil, args)
}

/* }}} */

func ZendFcallInfoArgp(fci *ZendFcallInfo, argc int, argv *Zval) int {
	var i int
	if argc < 0 {
		return FAILURE
	}
	ZendFcallInfoArgsClear(fci, !argc)
	if argc != 0 {
		fci.SetParamCount(argc)
		fci.SetParams((*Zval)(_erealloc(fci.GetParams(), fci.GetParamCount()*g.SizeOf("zval"))))
		for i = 0; i < argc; i++ {
			var _z1 *Zval = &fci.params[i]
			var _z2 *Zval = &argv[i]
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	}
	return SUCCESS
}

/* }}} */

func ZendFcallInfoArgv(fci *ZendFcallInfo, argc int, argv *va_list) int {
	var i int
	var arg *Zval
	if argc < 0 {
		return FAILURE
	}
	ZendFcallInfoArgsClear(fci, !argc)
	if argc != 0 {
		fci.SetParamCount(argc)
		fci.SetParams((*Zval)(_erealloc(fci.GetParams(), fci.GetParamCount()*g.SizeOf("zval"))))
		for i = 0; i < argc; i++ {
			arg = __va_arg(*argv, (*Zval)(_))
			var _z1 *Zval = &fci.params[i]
			var _z2 *Zval = arg
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	}
	return SUCCESS
}

/* }}} */

func ZendFcallInfoArgn(fci *ZendFcallInfo, argc int, _ ...any) int {
	var ret int
	var argv va_list
	va_start(argv, argc)
	ret = ZendFcallInfoArgv(fci, argc, &argv)
	va_end(argv)
	return ret
}

/* }}} */

func ZendFcallInfoCall(fci *ZendFcallInfo, fcc *ZendFcallInfoCache, retval_ptr *Zval, args *Zval) int {
	var retval Zval
	var org_params *Zval = nil
	var result int
	var org_count int = 0
	if retval_ptr != nil {
		fci.SetRetval(retval_ptr)
	} else {
		fci.SetRetval(&retval)
	}
	if args != nil {
		ZendFcallInfoArgsSave(fci, &org_count, &org_params)
		ZendFcallInfoArgs(fci, args)
	}
	result = ZendCallFunction(fci, fcc)
	if retval_ptr == nil && retval.GetType() != 0 {
		ZvalPtrDtor(&retval)
	}
	if args != nil {
		ZendFcallInfoArgsRestore(fci, org_count, org_params)
	}
	return result
}

/* }}} */

func ZendGetModuleVersion(module_name *byte) *byte {
	var lname *ZendString
	var name_len int = strlen(module_name)
	var module *ZendModuleEntry
	lname = ZendStringAlloc(name_len, 0)
	ZendStrTolowerCopy(lname.GetVal(), module_name, name_len)
	module = ZendHashFindPtr(&ModuleRegistry, lname)
	ZendStringEfree(lname)
	if module != nil {
		return module.GetVersion()
	} else {
		return nil
	}
}

/* }}} */

func ZvalMakeInternedString(zv *Zval) *ZendString {
	assert(zv.GetType() == 6)
	zv.GetValue().SetStr(ZendNewInternedString(zv.GetValue().GetStr()))
	if (ZvalGcFlags(zv.GetValue().GetStr().GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		zv.SetTypeFlags(0)
	}
	return zv.GetValue().GetStr()
}
func IsPersistentClass(ce *ZendClassEntry) ZendBool {
	return (ce.GetType()&1) != 0 && ce.GetModule().GetType() == 1
}
func ZendDeclareTypedProperty(ce *ZendClassEntry, name *ZendString, property *Zval, access_type int, doc_comment *ZendString, type_ ZendType) int {
	var property_info *ZendPropertyInfo
	var property_info_ptr *ZendPropertyInfo
	if type_ > 0x3 {
		ce.SetCeFlags(ce.GetCeFlags() | 1<<8)
	}
	if ce.GetType() == 1 {
		property_info = __zendMalloc(g.SizeOf("zend_property_info"))
	} else {
		property_info = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_property_info"))
		if property.GetType() == 11 {
			ce.SetCeFlags(ce.GetCeFlags() &^ (1 << 12))
		}
	}
	if property.GetType() == 6 && (ZvalGcFlags(property.GetValue().GetStr().GetGc().GetTypeInfo())&1<<6) == 0 {
		ZvalMakeInternedString(property)
	}
	if (access_type & (1<<0 | 1<<1 | 1<<2)) == 0 {
		access_type |= 1 << 0
	}
	if (access_type & 1 << 4) != 0 {
		if g.Assign(&property_info_ptr, ZendHashFindPtr(&ce.properties_info, name)) != nil && (property_info_ptr.GetFlags()&1<<4) != 0 {
			property_info.SetOffset(property_info_ptr.GetOffset())
			ZvalPtrDtor(&ce.default_static_members_table[property_info.GetOffset()])
			ZendHashDel(&ce.properties_info, name)
		} else {
			ce.GetDefaultStaticMembersCount()++
			property_info.SetOffset(ce.GetDefaultStaticMembersCount() - 1)
			if ce.GetType() == 1 {
				ce.SetDefaultStaticMembersTable(__zendRealloc(ce.GetDefaultStaticMembersTable(), g.SizeOf("zval")*ce.GetDefaultStaticMembersCount()))
			} else {
				ce.SetDefaultStaticMembersTable(_erealloc(ce.GetDefaultStaticMembersTable(), g.SizeOf("zval")*ce.GetDefaultStaticMembersCount()))
			}
		}
		var _z1 *Zval = &ce.default_static_members_table[property_info.GetOffset()]
		var _z2 *Zval = property
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if ce.GetStaticMembersTablePtr() == nil {
			assert(ce.GetType() == 1)
			if EG.GetCurrentExecuteData() == nil {
				ce.SetStaticMembersTablePtr(ZendMapPtrNew())
			} else {

				/* internal class loaded by dl() */

				ce.SetStaticMembersTablePtr(&ce.default_static_members_table)

				/* internal class loaded by dl() */

			}
		}
	} else {
		var property_default_ptr *Zval
		if g.Assign(&property_info_ptr, ZendHashFindPtr(&ce.properties_info, name)) != nil && (property_info_ptr.GetFlags()&1<<4) == 0 {
			property_info.SetOffset(property_info_ptr.GetOffset())
			ZvalPtrDtor(&ce.default_properties_table[(property_info.GetOffset()-uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0))/g.SizeOf("zval")])
			ZendHashDel(&ce.properties_info, name)
			assert(ce.GetType() == 1)
			assert(ce.GetPropertiesInfoTable() != nil)
			ce.GetPropertiesInfoTable()[(property_info.GetOffset()-uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0))/g.SizeOf("zval")] = property_info
		} else {
			property_info.SetOffset(uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil)) + g.SizeOf("zval")*ce.GetDefaultPropertiesCount()))
			ce.GetDefaultPropertiesCount()++
			if ce.GetType() == 1 {
				ce.SetDefaultPropertiesTable(__zendRealloc(ce.GetDefaultPropertiesTable(), g.SizeOf("zval")*ce.GetDefaultPropertiesCount()))
			} else {
				ce.SetDefaultPropertiesTable(_erealloc(ce.GetDefaultPropertiesTable(), g.SizeOf("zval")*ce.GetDefaultPropertiesCount()))
			}

			/* For user classes this is handled during linking */

			if ce.GetType() == 1 {
				ce.SetPropertiesInfoTable(__zendRealloc(ce.GetPropertiesInfoTable(), g.SizeOf("zend_property_info *")*ce.GetDefaultPropertiesCount()))
				ce.GetPropertiesInfoTable()[ce.GetDefaultPropertiesCount()-1] = property_info
			}

			/* For user classes this is handled during linking */

		}
		property_default_ptr = &ce.default_properties_table[(property_info.GetOffset()-uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0))/g.SizeOf("zval")]
		var _z1 *Zval = property_default_ptr
		var _z2 *Zval = property
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if property.GetType() == 0 {
			property_default_ptr.SetU2Extra(1)
		} else {
			property_default_ptr.SetU2Extra(0)
		}
	}
	if (ce.GetType() & 1) != 0 {
		switch property.GetType() {
		case 7:

		case 8:

		case 9:
			ZendErrorNoreturn(1<<4, "Internal zval's can't be arrays, objects or resources")
			break
		default:
			break
		}

		/* Must be interned to avoid ZTS data races */

		if IsPersistentClass(ce) != 0 {
			name = ZendNewInternedString(ZendStringCopy(name))
		}

		/* Must be interned to avoid ZTS data races */

	}
	if (access_type & 1 << 0) != 0 {
		property_info.SetName(ZendStringCopy(name))
	} else if (access_type & 1 << 2) != 0 {
		property_info.SetName(ZendManglePropertyName(ce.GetName().GetVal(), ce.GetName().GetLen(), name.GetVal(), name.GetLen(), IsPersistentClass(ce)))
	} else {
		assert((access_type & 1 << 1) != 0)
		property_info.SetName(ZendManglePropertyName("*", 1, name.GetVal(), name.GetLen(), IsPersistentClass(ce)))
	}
	property_info.SetName(ZendNewInternedString(property_info.GetName()))
	property_info.SetFlags(access_type)
	property_info.SetDocComment(doc_comment)
	property_info.SetCe(ce)
	property_info.SetType(type_)
	ZendHashUpdatePtr(&ce.properties_info, name, property_info)
	return SUCCESS
}

/* }}} */

func ZendTryAssignTypedRefEx(ref *ZendReference, val *Zval, strict ZendBool) int {
	if ZendVerifyRefAssignableZval(ref, val, strict) == 0 {
		ZvalPtrDtor(val)
		return FAILURE
	} else {
		ZvalPtrDtor(&ref.val)
		var _z1 *Zval = &ref.val
		var _z2 *Zval = val
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		return SUCCESS
	}
}

/* }}} */

func ZendTryAssignTypedRef(ref *ZendReference, val *Zval) int {
	return ZendTryAssignTypedRefEx(ref, val, EG.GetCurrentExecuteData().GetPrevExecuteData() != nil && EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetPrevExecuteData().GetFunc().GetFnFlags()&1<<31) != 0)
}

/* }}} */

func ZendTryAssignTypedRefNull(ref *ZendReference) int {
	var tmp Zval
	&tmp.SetTypeInfo(1)
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefBool(ref *ZendReference, val ZendBool) int {
	var tmp Zval
	if val != 0 {
		&tmp.SetTypeInfo(3)
	} else {
		&tmp.SetTypeInfo(2)
	}
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefLong(ref *ZendReference, lval ZendLong) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetLval(lval)
	__z.SetTypeInfo(4)
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefDouble(ref *ZendReference, dval float64) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetDval(dval)
	__z.SetTypeInfo(5)
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefEmptyString(ref *ZendReference) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = ZendEmptyString
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6)
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefStr(ref *ZendReference, str *ZendString) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = str
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefString(ref *ZendReference, string *byte) int {
	var tmp Zval
	var _s *byte = string
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefStringl(ref *ZendReference, string *byte, len_ int) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(string, len_, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefArr(ref *ZendReference, arr *ZendArray) int {
	var tmp Zval
	var __arr *ZendArray = arr
	var __z *Zval = &tmp
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefRes(ref *ZendReference, res *ZendResource) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetRes(res)
	__z.SetTypeInfo(9 | 1<<0<<8)
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefZval(ref *ZendReference, zv *Zval) int {
	var tmp Zval
	var _z1 *Zval = &tmp
	var _z2 *Zval = zv
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	return ZendTryAssignTypedRef(ref, &tmp)
}

/* }}} */

func ZendTryAssignTypedRefZvalEx(ref *ZendReference, zv *Zval, strict ZendBool) int {
	var tmp Zval
	var _z1 *Zval = &tmp
	var _z2 *Zval = zv
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	return ZendTryAssignTypedRefEx(ref, &tmp, strict)
}

/* }}} */

func ZendDeclarePropertyEx(ce *ZendClassEntry, name *ZendString, property *Zval, access_type int, doc_comment *ZendString) int {
	return ZendDeclareTypedProperty(ce, name, property, access_type, doc_comment, 0)
}

/* }}} */

func ZendDeclareProperty(ce *ZendClassEntry, name *byte, name_length int, property *Zval, access_type int) int {
	var key *ZendString = ZendStringInit(name, name_length, IsPersistentClass(ce))
	var ret int = ZendDeclarePropertyEx(ce, key, property, access_type, nil)
	ZendStringRelease(key)
	return ret
}

/* }}} */

func ZendDeclarePropertyNull(ce *ZendClassEntry, name string, name_length int, access_type int) int {
	var property Zval
	&property.SetTypeInfo(1)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}

/* }}} */

func ZendDeclarePropertyBool(ce *ZendClassEntry, name *byte, name_length int, value ZendLong, access_type int) int {
	var property Zval
	if value != 0 {
		&property.SetTypeInfo(3)
	} else {
		&property.SetTypeInfo(2)
	}
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}

/* }}} */

func ZendDeclarePropertyLong(ce *ZendClassEntry, name string, name_length int, value ZendLong, access_type int) int {
	var property Zval
	var __z *Zval = &property
	__z.GetValue().SetLval(value)
	__z.SetTypeInfo(4)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}

/* }}} */

func ZendDeclarePropertyDouble(ce *ZendClassEntry, name *byte, name_length int, value float64, access_type int) int {
	var property Zval
	var __z *Zval = &property
	__z.GetValue().SetDval(value)
	__z.SetTypeInfo(5)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}

/* }}} */

func ZendDeclarePropertyString(ce *ZendClassEntry, name string, name_length int, value string, access_type int) int {
	var property Zval
	var __z *Zval = &property
	var __s *ZendString = ZendStringInit(value, strlen(value), ce.GetType()&1)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}

/* }}} */

func ZendDeclarePropertyStringl(ce *ZendClassEntry, name *byte, name_length int, value *byte, value_len int, access_type int) int {
	var property Zval
	var __z *Zval = &property
	var __s *ZendString = ZendStringInit(value, value_len, ce.GetType()&1)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}

/* }}} */

func ZendDeclareClassConstantEx(ce *ZendClassEntry, name *ZendString, value *Zval, access_type int, doc_comment *ZendString) int {
	var c *ZendClassConstant
	if (ce.GetCeFlags() & 1 << 0) != 0 {
		if access_type != 1<<0 {
			ZendErrorNoreturn(1<<6, "Access type for interface constant %s::%s must be public", ce.GetName().GetVal(), name.GetVal())
		}
	}
	if name.GetLen() == g.SizeOf("\"class\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "class", g.SizeOf("\"class\"")-1) == 0 {
		ZendErrorNoreturn(g.Cond(ce.GetType() == 1, 1<<4, 1<<6), "A class constant must not be called 'class'; it is reserved for class name fetching")
	}
	if value.GetType() == 6 && (ZvalGcFlags(value.GetValue().GetStr().GetGc().GetTypeInfo())&1<<6) == 0 {
		ZvalMakeInternedString(value)
	}
	if ce.GetType() == 1 {
		c = __zendMalloc(g.SizeOf("zend_class_constant"))
	} else {
		c = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_class_constant"))
	}
	var _z1 *Zval = &c.value
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	c.GetValue().SetAccessFlags(access_type)
	c.SetDocComment(doc_comment)
	c.SetCe(ce)
	if value.GetType() == 11 {
		ce.SetCeFlags(ce.GetCeFlags() &^ (1 << 12))
	}
	if !(ZendHashAddPtr(&ce.constants_table, name, c)) {
		ZendErrorNoreturn(g.Cond(ce.GetType() == 1, 1<<4, 1<<6), "Cannot redefine class __special__  constant %s::%s", ce.GetName().GetVal(), name.GetVal())
	}
	return SUCCESS
}

/* }}} */

func ZendDeclareClassConstant(ce *ZendClassEntry, name *byte, name_length int, value *Zval) int {
	var ret int
	var key *ZendString
	if ce.GetType() == 1 {
		key = ZendStringInitInterned(name, name_length, 1)
	} else {
		key = ZendStringInit(name, name_length, 0)
	}
	ret = ZendDeclareClassConstantEx(ce, key, value, 1<<0, nil)
	ZendStringRelease(key)
	return ret
}

/* }}} */

func ZendDeclareClassConstantNull(ce *ZendClassEntry, name *byte, name_length int) int {
	var constant Zval
	&constant.SetTypeInfo(1)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}

/* }}} */

func ZendDeclareClassConstantLong(ce *ZendClassEntry, name string, name_length int, value ZendLong) int {
	var constant Zval
	var __z *Zval = &constant
	__z.GetValue().SetLval(value)
	__z.SetTypeInfo(4)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}

/* }}} */

func ZendDeclareClassConstantBool(ce *ZendClassEntry, name *byte, name_length int, value ZendBool) int {
	var constant Zval
	if value != 0 {
		&constant.SetTypeInfo(3)
	} else {
		&constant.SetTypeInfo(2)
	}
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}

/* }}} */

func ZendDeclareClassConstantDouble(ce *ZendClassEntry, name *byte, name_length int, value float64) int {
	var constant Zval
	var __z *Zval = &constant
	__z.GetValue().SetDval(value)
	__z.SetTypeInfo(5)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}

/* }}} */

func ZendDeclareClassConstantStringl(ce *ZendClassEntry, name *byte, name_length int, value *byte, value_length int) int {
	var constant Zval
	var __z *Zval = &constant
	var __s *ZendString = ZendStringInit(value, value_length, ce.GetType()&1)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}

/* }}} */

func ZendDeclareClassConstantString(ce *ZendClassEntry, name *byte, name_length int, value *byte) int {
	return ZendDeclareClassConstantStringl(ce, name, name_length, value, strlen(value))
}

/* }}} */

func ZendUpdatePropertyEx(scope *ZendClassEntry, object *Zval, name *ZendString, value *Zval) {
	var property Zval
	var old_scope *ZendClassEntry = EG.GetFakeScope()
	EG.SetFakeScope(scope)
	var __z *Zval = &property
	var __s *ZendString = name
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, &property, value, nil)
	EG.SetFakeScope(old_scope)
}

/* }}} */

func ZendUpdateProperty(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *Zval) {
	var property Zval
	var old_scope *ZendClassEntry = EG.GetFakeScope()
	EG.SetFakeScope(scope)
	var __z *Zval = &property
	var __s *ZendString = ZendStringInit(name, name_length, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, &property, value, nil)
	ZvalPtrDtor(&property)
	EG.SetFakeScope(old_scope)
}

/* }}} */

func ZendUpdatePropertyNull(scope *ZendClassEntry, object *Zval, name *byte, name_length int) {
	var tmp Zval
	&tmp.SetTypeInfo(1)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}

/* }}} */

func ZendUnsetProperty(scope *ZendClassEntry, object *Zval, name *byte, name_length int) {
	var property Zval
	var old_scope *ZendClassEntry = EG.GetFakeScope()
	EG.SetFakeScope(scope)
	var __z *Zval = &property
	var __s *ZendString = ZendStringInit(name, name_length, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	object.GetValue().GetObj().GetHandlers().GetUnsetProperty()(object, &property, 0)
	ZvalPtrDtor(&property)
	EG.SetFakeScope(old_scope)
}

/* }}} */

func ZendUpdatePropertyBool(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value ZendLong) {
	var tmp Zval
	if value != 0 {
		&tmp.SetTypeInfo(3)
	} else {
		&tmp.SetTypeInfo(2)
	}
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}

/* }}} */

func ZendUpdatePropertyLong(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value ZendLong) {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetLval(value)
	__z.SetTypeInfo(4)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}

/* }}} */

func ZendUpdatePropertyDouble(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value float64) {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetDval(value)
	__z.SetTypeInfo(5)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}

/* }}} */

func ZendUpdatePropertyStr(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *ZendString) {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = value
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}

/* }}} */

func ZendUpdatePropertyString(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *byte) {
	var tmp Zval
	var _s *byte = value
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	ZvalSetRefcountP(&tmp, 0)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}

/* }}} */

func ZendUpdatePropertyStringl(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *byte, value_len int) {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(value, value_len, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	ZvalSetRefcountP(&tmp, 0)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}

/* }}} */

func ZendUpdateStaticPropertyEx(scope *ZendClassEntry, name *ZendString, value *Zval) int {
	var property *Zval
	var tmp Zval
	var prop_info *ZendPropertyInfo
	var old_scope *ZendClassEntry = EG.GetFakeScope()
	if (scope.GetCeFlags() & 1 << 12) == 0 {
		if ZendUpdateClassConstants(scope) != SUCCESS {
			return FAILURE
		}
	}
	EG.SetFakeScope(scope)
	property = ZendStdGetStaticPropertyWithInfo(scope, name, 1, &prop_info)
	EG.SetFakeScope(old_scope)
	if property == nil {
		return FAILURE
	}
	assert(value.GetType() != 10)
	if value.GetTypeFlags() != 0 {
		ZvalAddrefP(value)
	}
	if prop_info.GetType() != 0 {
		var _z1 *Zval = &tmp
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if ZendVerifyPropertyType(prop_info, &tmp, 0) == 0 {
			if value.GetTypeFlags() != 0 {
				ZvalDelrefP(value)
			}
			return FAILURE
		}
		value = &tmp
	}
	ZendAssignToVariable(property, value, 1<<1, 0)
	return SUCCESS
}

/* }}} */

func ZendUpdateStaticProperty(scope *ZendClassEntry, name *byte, name_length int, value *Zval) int {
	var key *ZendString = ZendStringInit(name, name_length, 0)
	var retval int = ZendUpdateStaticPropertyEx(scope, key, value)
	ZendStringEfree(key)
	return retval
}

/* }}} */

func ZendUpdateStaticPropertyNull(scope *ZendClassEntry, name *byte, name_length int) int {
	var tmp Zval
	&tmp.SetTypeInfo(1)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}

/* }}} */

func ZendUpdateStaticPropertyBool(scope *ZendClassEntry, name *byte, name_length int, value ZendLong) int {
	var tmp Zval
	if value != 0 {
		&tmp.SetTypeInfo(3)
	} else {
		&tmp.SetTypeInfo(2)
	}
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}

/* }}} */

func ZendUpdateStaticPropertyLong(scope *ZendClassEntry, name *byte, name_length int, value ZendLong) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetLval(value)
	__z.SetTypeInfo(4)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}

/* }}} */

func ZendUpdateStaticPropertyDouble(scope *ZendClassEntry, name *byte, name_length int, value float64) int {
	var tmp Zval
	var __z *Zval = &tmp
	__z.GetValue().SetDval(value)
	__z.SetTypeInfo(5)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}

/* }}} */

func ZendUpdateStaticPropertyString(scope *ZendClassEntry, name *byte, name_length int, value *byte) int {
	var tmp Zval
	var _s *byte = value
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	ZvalSetRefcountP(&tmp, 0)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}

/* }}} */

func ZendUpdateStaticPropertyStringl(scope *ZendClassEntry, name *byte, name_length int, value *byte, value_len int) int {
	var tmp Zval
	var __z *Zval = &tmp
	var __s *ZendString = ZendStringInit(value, value_len, 0)
	__z.GetValue().SetStr(__s)
	__z.SetTypeInfo(6 | 1<<0<<8)
	ZvalSetRefcountP(&tmp, 0)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}

/* }}} */

func ZendReadPropertyEx(scope *ZendClassEntry, object *Zval, name *ZendString, silent ZendBool, rv *Zval) *Zval {
	var property Zval
	var value *Zval
	var old_scope *ZendClassEntry = EG.GetFakeScope()
	EG.SetFakeScope(scope)
	var __z *Zval = &property
	var __s *ZendString = name
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	value = object.GetValue().GetObj().GetHandlers().GetReadProperty()(object, &property, g.Cond(silent != 0, 3, 0), nil, rv)
	EG.SetFakeScope(old_scope)
	return value
}

/* }}} */

func ZendReadProperty(scope *ZendClassEntry, object *Zval, name string, name_length int, silent ZendBool, rv *Zval) *Zval {
	var value *Zval
	var str *ZendString
	str = ZendStringInit(name, name_length, 0)
	value = ZendReadPropertyEx(scope, object, str, silent, rv)
	ZendStringReleaseEx(str, 0)
	return value
}

/* }}} */

func ZendReadStaticPropertyEx(scope *ZendClassEntry, name *ZendString, silent ZendBool) *Zval {
	var property *Zval
	var old_scope *ZendClassEntry = EG.GetFakeScope()
	EG.SetFakeScope(scope)
	property = ZendStdGetStaticProperty(scope, name, g.Cond(silent != 0, 3, 0))
	EG.SetFakeScope(old_scope)
	return property
}

/* }}} */

func ZendReadStaticProperty(scope *ZendClassEntry, name *byte, name_length int, silent ZendBool) *Zval {
	var key *ZendString = ZendStringInit(name, name_length, 0)
	var property *Zval = ZendReadStaticPropertyEx(scope, key, silent)
	ZendStringEfree(key)
	return property
}

/* }}} */

func ZendSaveErrorHandling(current *ZendErrorHandling) {
	current.SetHandling(EG.GetErrorHandling())
	current.SetException(EG.GetExceptionClass())
	&current.user_handler.u1.type_info = 0
}

/* }}} */

func ZendReplaceErrorHandling(error_handling ZendErrorHandlingT, exception_class *ZendClassEntry, current *ZendErrorHandling) {
	if current != nil {
		ZendSaveErrorHandling(current)
	}
	assert(error_handling == EH_THROW || exception_class == nil)
	EG.SetErrorHandling(error_handling)
	EG.SetExceptionClass(exception_class)
}

/* }}} */

func ZendRestoreErrorHandling(saved *ZendErrorHandling) {
	EG.SetErrorHandling(saved.GetHandling())
	EG.SetExceptionClass(saved.GetException())
}

/* }}} */

func ZendFindAliasName(ce *ZendClassEntry, name *ZendString) *ZendString {
	var alias *ZendTraitAlias
	var alias_ptr **ZendTraitAlias
	if g.Assign(&alias_ptr, ce.GetTraitAliases()) {
		alias = *alias_ptr
		for alias != nil {
			if alias.GetAlias() != nil && (alias.GetAlias().GetLen() == name.GetLen() && ZendBinaryStrcasecmp(alias.GetAlias().GetVal(), alias.GetAlias().GetLen(), name.GetVal(), name.GetLen()) == 0) {
				return alias.GetAlias()
			}
			alias_ptr++
			alias = *alias_ptr
		}
	}
	return name
}

/* }}} */

func ZendResolveMethodName(ce *ZendClassEntry, f *ZendFunction) *ZendString {
	var func_ *ZendFunction
	var function_table *HashTable
	var name *ZendString
	if f.GetCommonType() != 2 || f.GetOpArray().GetRefcount() != nil && (*(f.GetOpArray().GetRefcount())) < 2 || f.GetScope() == nil || f.GetScope().GetTraitAliases() == nil {
		return f.GetFunctionName()
	}
	function_table = &ce.function_table
	for {
		var __ht *HashTable = function_table
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			name = _p.GetKey()
			func_ = _z.GetValue().GetPtr()
			if func_ == f {
				if name == nil {
					return f.GetFunctionName()
				}
				if name.GetLen() == f.GetFunctionName().GetLen() && !(strncasecmp(name.GetVal(), f.GetFunctionName().GetVal(), f.GetFunctionName().GetLen())) {
					return f.GetFunctionName()
				}
				return ZendFindAliasName(f.GetScope(), name)
			}
		}
		break
	}
	return f.GetFunctionName()
}

/* }}} */

func ZendGetObjectType(ce *ZendClassEntry) *byte {
	if (ce.GetCeFlags() & 1 << 1) != 0 {
		return "trait"
	} else if (ce.GetCeFlags() & 1 << 0) != 0 {
		return "interface"
	} else {
		return "class"
	}
}

/* }}} */

func ZendIsIterable(iterable *Zval) ZendBool {
	switch iterable.GetType() {
	case 7:
		return 1
	case 8:
		return InstanceofFunction(iterable.GetValue().GetObj().GetCe(), ZendCeTraversable)
	default:
		return 0
	}
}

/* }}} */

func ZendIsCountable(countable *Zval) ZendBool {
	switch countable.GetType() {
	case 7:
		return 1
	case 8:
		if countable.GetValue().GetObj().GetHandlers().GetCountElements() != nil {
			return 1
		}
		return InstanceofFunction(countable.GetValue().GetObj().GetCe(), ZendCeCountable)
	default:
		return 0
	}
}

/* }}} */
