package spl

// #define SPL_ITERATORS_H

// # include "php.h"

// # include "php_spl.h"

// # include "php.h"

// # include "php_ini.h"

// # include "ext/standard/info.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_iterators.h"

// # include "spl_directory.h"

// # include "spl_array.h"

// # include "spl_exceptions.h"

// # include "zend_smart_str.h"

// #define SPL_FETCH_AND_CHECK_DUAL_IT(var,objzval) do { spl_dual_it_object * it = Z_SPLDUAL_IT_P ( objzval ) ; if ( it -> dit_type == DIT_Unknown ) { zend_throw_exception_ex ( spl_ce_LogicException , 0 , "The object is in an invalid state as the parent constructor was not called" ) ; return ; } ( var ) = it ; } while ( 0 )

// #define SPL_FETCH_SUB_ELEMENT(var,object,element) do { if ( ! ( object ) -> iterators ) { zend_throw_exception_ex ( spl_ce_LogicException , 0 , "The object is in an invalid state as the parent constructor was not called" ) ; return ; } ( var ) = ( object ) -> iterators [ ( object ) -> level ] . element ; } while ( 0 )

// #define SPL_FETCH_SUB_ELEMENT_ADDR(var,object,element) do { if ( ! ( object ) -> iterators ) { zend_throw_exception_ex ( spl_ce_LogicException , 0 , "The object is in an invalid state as the parent constructor was not called" ) ; return ; } ( var ) = & ( object ) -> iterators [ ( object ) -> level ] . element ; } while ( 0 )

// #define SPL_CHECK_CTOR(intern,classname) if ( intern -> dit_type == DIT_Unknown ) { zend_throw_exception_ex ( spl_ce_BadMethodCallException , 0 , "Classes derived from %s must call %s::__construct()" , ZSTR_VAL ( ( spl_ce_ ## classname ) -> name ) , ZSTR_VAL ( ( spl_ce_ ## classname ) -> name ) ) ; return ; }
