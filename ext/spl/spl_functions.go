// <<generate>>

package spl

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/spl/spl_functions.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define PHP_FUNCTIONS_H

// # include "php.h"

type CreateObjectFuncT func(class_type *zend.ZendClassEntry) *zend.ZendObject

// #define REGISTER_SPL_STD_CLASS(class_name,obj_ctor) spl_register_std_class ( & spl_ce_ ## class_name , # class_name , obj_ctor , NULL ) ;

// #define REGISTER_SPL_STD_CLASS_EX(class_name,obj_ctor,funcs) spl_register_std_class ( & spl_ce_ ## class_name , # class_name , obj_ctor , funcs ) ;

// #define REGISTER_SPL_SUB_CLASS_EX(class_name,parent_class_name,obj_ctor,funcs) spl_register_sub_class ( & spl_ce_ ## class_name , spl_ce_ ## parent_class_name , # class_name , obj_ctor , funcs ) ;

// #define REGISTER_SPL_INTERFACE(class_name) spl_register_interface ( & spl_ce_ ## class_name , # class_name , spl_funcs_ ## class_name ) ;

// #define REGISTER_SPL_IMPLEMENTS(class_name,interface_name) zend_class_implements ( spl_ce_ ## class_name , 1 , spl_ce_ ## interface_name ) ;

// #define REGISTER_SPL_ITERATOR(class_name) zend_class_implements ( spl_ce_ ## class_name , 1 , zend_ce_iterator ) ;

// #define REGISTER_SPL_PROPERTY(class_name,prop_name,prop_flags) spl_register_property ( spl_ce_ ## class_name , prop_name , sizeof ( prop_name ) - 1 , prop_flags ) ;

// #define REGISTER_SPL_CLASS_CONST_LONG(class_name,const_name,value) zend_declare_class_constant_long ( spl_ce_ ## class_name , const_name , sizeof ( const_name ) - 1 , ( zend_long ) value ) ;

/* sub: whether to allow subclasses/interfaces
   allow = 0: allow all classes and interfaces
   allow > 0: allow all that match and mask ce_flags
   allow < 0: disallow all that match and mask ce_flags
*/

/* caller must efree(return) */

// #define SPL_ME(class_name,function_name,arg_info,flags) PHP_ME ( spl_ ## class_name , function_name , arg_info , flags )

// #define SPL_ABSTRACT_ME(class_name,function_name,arg_info) ZEND_ABSTRACT_ME ( spl_ ## class_name , function_name , arg_info )

// #define SPL_METHOD(class_name,function_name) PHP_METHOD ( spl_ ## class_name , function_name )

// #define SPL_MA(class_name,function_name,alias_class,alias_function,arg_info,flags) PHP_MALIAS ( spl_ ## alias_class , function_name , alias_function , arg_info , flags )

// Source: <ext/spl/spl_functions.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_ini.h"

// # include "ext/standard/info.h"

// # include "php_spl.h"

/* {{{ spl_register_interface */

func SplRegisterInterface(ppce **zend.ZendClassEntry, class_name string, functions *zend.ZendFunctionEntry) {
	var ce zend.ZendClassEntry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.name = zend.ZendStringInitInterned(class_name, strlen(class_name), 1)
	ce.info.internal.builtin_functions = functions
	*ppce = zend.ZendRegisterInternalInterface(&ce)
}

/* }}} */

func SplRegisterStdClass(ppce **zend.ZendClassEntry, class_name string, obj_ctor any, function_list *zend.ZendFunctionEntry) {
	var ce zend.ZendClassEntry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.name = zend.ZendStringInitInterned(class_name, strlen(class_name), 1)
	ce.info.internal.builtin_functions = function_list
	*ppce = zend.ZendRegisterInternalClass(&ce)

	/* entries changed by initialize */

	if obj_ctor {
		(*ppce).create_object = obj_ctor
	}

	/* entries changed by initialize */
}

/* }}} */

func SplRegisterSubClass(ppce **zend.ZendClassEntry, parent_ce *zend.ZendClassEntry, class_name string, obj_ctor any, function_list *zend.ZendFunctionEntry) {
	var ce zend.ZendClassEntry
	memset(&ce, 0, g.SizeOf("zend_class_entry"))
	ce.name = zend.ZendStringInitInterned(class_name, strlen(class_name), 1)
	ce.info.internal.builtin_functions = function_list
	*ppce = zend.ZendRegisterInternalClassEx(&ce, parent_ce)

	/* entries changed by initialize */

	if obj_ctor {
		(*ppce).create_object = obj_ctor
	} else {
		(*ppce).create_object = parent_ce.create_object
	}

	/* entries changed by initialize */
}

/* }}} */

func SplRegisterProperty(class_entry *zend.ZendClassEntry, prop_name string, prop_name_len int, prop_flags int) {
	zend.ZendDeclarePropertyNull(class_entry, prop_name, prop_name_len, prop_flags)
}

/* }}} */

func SplAddClassName(list *zend.Zval, pce *zend.ZendClassEntry, allow int, ce_flags int) {
	if allow == 0 || allow > 0 && (pce.ce_flags&ce_flags) != 0 || allow < 0 && (pce.ce_flags&ce_flags) == 0 {
		var tmp *zend.Zval
		if g.Assign(&tmp, zend.ZendHashFind(list.value.arr, pce.name)) == nil {
			var t zend.Zval
			var __z *zend.Zval = &t
			var __s *zend.ZendString = pce.name
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				zend.ZendGcAddref(&__s.gc)
				__z.u1.type_info = 6 | 1<<0<<8
			}
			zend.ZendHashAdd(list.value.arr, pce.name, &t)
		}
	}
}

/* }}} */

func SplAddInterfaces(list *zend.Zval, pce *zend.ZendClassEntry, allow int, ce_flags int) {
	var num_interfaces uint32
	if pce.num_interfaces != 0 {
		r.Assert((pce.ce_flags & 1 << 3) != 0)
		for num_interfaces = 0; num_interfaces < pce.num_interfaces; num_interfaces++ {
			SplAddClassName(list, pce.interfaces[num_interfaces], allow, ce_flags)
		}
	}
}

/* }}} */

func SplAddTraits(list *zend.Zval, pce *zend.ZendClassEntry, allow int, ce_flags int) {
	var num_traits uint32
	var trait *zend.ZendClassEntry
	for num_traits = 0; num_traits < pce.num_traits; num_traits++ {
		trait = zend.ZendFetchClassByName(pce.trait_names[num_traits].name, pce.trait_names[num_traits].lc_name, 6)
		r.Assert(trait != nil)
		SplAddClassName(list, trait, allow, ce_flags)
	}
}

/* }}} */

func SplAddClasses(pce *zend.ZendClassEntry, list *zend.Zval, sub int, allow int, ce_flags int) int {
	if pce == nil {
		return 0
	}
	SplAddClassName(list, pce, allow, ce_flags)
	if sub != 0 {
		SplAddInterfaces(list, pce, allow, ce_flags)
		for pce.parent {
			pce = pce.parent
			SplAddClasses(pce, list, sub, allow, ce_flags)
		}
	}
	return 0
}

/* }}} */

func SplGenPrivatePropName(ce *zend.ZendClassEntry, prop_name string, prop_len int) *zend.ZendString {
	return zend.ZendManglePropertyName(ce.name.val, ce.name.len_, prop_name, prop_len, 0)
}

/* }}} */
