// <<generate>>

package spl

// #define PHP_FUNCTIONS_H

// # include "php.h"

// #define REGISTER_SPL_STD_CLASS(class_name,obj_ctor) spl_register_std_class ( & spl_ce_ ## class_name , # class_name , obj_ctor , NULL ) ;

// #define REGISTER_SPL_STD_CLASS_EX(class_name,obj_ctor,funcs) spl_register_std_class ( & spl_ce_ ## class_name , # class_name , obj_ctor , funcs ) ;

// #define REGISTER_SPL_SUB_CLASS_EX(class_name,parent_class_name,obj_ctor,funcs) spl_register_sub_class ( & spl_ce_ ## class_name , spl_ce_ ## parent_class_name , # class_name , obj_ctor , funcs ) ;

// #define REGISTER_SPL_INTERFACE(class_name) spl_register_interface ( & spl_ce_ ## class_name , # class_name , spl_funcs_ ## class_name ) ;

// #define REGISTER_SPL_IMPLEMENTS(class_name,interface_name) zend_class_implements ( spl_ce_ ## class_name , 1 , spl_ce_ ## interface_name ) ;

// #define REGISTER_SPL_ITERATOR(class_name) zend_class_implements ( spl_ce_ ## class_name , 1 , zend_ce_iterator ) ;

// #define REGISTER_SPL_PROPERTY(class_name,prop_name,prop_flags) spl_register_property ( spl_ce_ ## class_name , prop_name , sizeof ( prop_name ) - 1 , prop_flags ) ;

// #define REGISTER_SPL_CLASS_CONST_LONG(class_name,const_name,value) zend_declare_class_constant_long ( spl_ce_ ## class_name , const_name , sizeof ( const_name ) - 1 , ( zend_long ) value ) ;

// #define SPL_ME(class_name,function_name,arg_info,flags) PHP_ME ( spl_ ## class_name , function_name , arg_info , flags )

// #define SPL_ABSTRACT_ME(class_name,function_name,arg_info) ZEND_ABSTRACT_ME ( spl_ ## class_name , function_name , arg_info )

// #define SPL_METHOD(class_name,function_name) PHP_METHOD ( spl_ ## class_name , function_name )

// #define SPL_MA(class_name,function_name,alias_class,alias_function,arg_info,flags) PHP_MALIAS ( spl_ ## alias_class , function_name , alias_function , arg_info , flags )

// # include "php.h"

// # include "php_ini.h"

// # include "ext/standard/info.h"

// # include "php_spl.h"
