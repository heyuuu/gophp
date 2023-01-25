// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/zend"
)

// Source: <ext/spl/spl_engine.h>

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

// #define SPL_ENGINE_H

// # include "php.h"

// # include "php_spl.h"

// # include "zend_interfaces.h"

/* {{{ spl_instantiate_arg_ex1 */

func SplInstantiateArgEx1(pce *zend.ZendClassEntry, retval *zend.Zval, arg1 *zend.Zval) int {
	var func_ *zend.ZendFunction = pce.constructor
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, zend.ZSTR_VAL(func_.common.function_name), zend.ZSTR_LEN(func_.common.function_name), nil, 1, arg1, nil)
	return 0
}

/* }}} */

func SplInstantiateArgEx2(pce *zend.ZendClassEntry, retval *zend.Zval, arg1 *zend.Zval, arg2 *zend.Zval) int {
	var func_ *zend.ZendFunction = pce.constructor
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, zend.ZSTR_VAL(func_.common.function_name), zend.ZSTR_LEN(func_.common.function_name), nil, 2, arg1, arg2)
	return 0
}

/* }}} */

func SplInstantiateArgN(pce *zend.ZendClassEntry, retval *zend.Zval, argc int, argv *zend.Zval) {
	var func_ *zend.ZendFunction = pce.constructor
	var fci zend.ZendFcallInfo
	var fcc zend.ZendFcallInfoCache
	var dummy zend.Zval
	SplInstantiate(pce, retval)
	fci.size = b.SizeOf("zend_fcall_info")
	zend.ZVAL_STR(&fci.function_name, func_.common.function_name)
	fci.object = zend.Z_OBJ_P(retval)
	fci.retval = &dummy
	fci.param_count = argc
	fci.params = argv
	fci.no_separation = 1
	fcc.function_handler = func_
	fcc.called_scope = pce
	fcc.object = zend.Z_OBJ_P(retval)
	zend.ZendCallFunction(&fci, &fcc)
}

/* }}} */

// Source: <ext/spl/spl_engine.c>

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

// # include "zend_interfaces.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_array.h"

/* {{{ spl_instantiate */

func SplInstantiate(pce *zend.ZendClassEntry, object *zend.Zval) { zend.ObjectInitEx(object, pce) }

/* }}} */

func SplOffsetConvertToLong(offset *zend.Zval) zend.ZendLong {
	var idx zend.ZendUlong
try_again:
	switch zend.Z_TYPE_P(offset) {
	case zend.IS_STRING:
		if zend.ZEND_HANDLE_NUMERIC(zend.Z_STR_P(offset), idx) != 0 {
			return idx
		}
		break
	case zend.IS_DOUBLE:
		return zend.ZendLong(zend.Z_DVAL_P(offset))
	case zend.IS_LONG:
		return zend.Z_LVAL_P(offset)
	case zend.IS_FALSE:
		return 0
	case zend.IS_TRUE:
		return 1
	case zend.IS_REFERENCE:
		offset = zend.Z_REFVAL_P(offset)
		goto try_again
	case zend.IS_RESOURCE:
		return zend.Z_RES_HANDLE_P(offset)
	}
	return -1
}

/* }}} */
