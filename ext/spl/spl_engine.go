// <<generate>>

package spl

import (
	g "sik/runtime/grammar"
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
	zend.ZendCallMethod(retval, pce, &func_, func_.common.function_name.val, func_.common.function_name.len_, nil, 1, arg1, nil)
	return 0
}

/* }}} */

func SplInstantiateArgEx2(pce *zend.ZendClassEntry, retval *zend.Zval, arg1 *zend.Zval, arg2 *zend.Zval) int {
	var func_ *zend.ZendFunction = pce.constructor
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, func_.common.function_name.val, func_.common.function_name.len_, nil, 2, arg1, arg2)
	return 0
}

/* }}} */

func SplInstantiateArgN(pce *zend.ZendClassEntry, retval *zend.Zval, argc int, argv *zend.Zval) {
	var func_ *zend.ZendFunction = pce.constructor
	var fci zend.ZendFcallInfo
	var fcc zend.ZendFcallInfoCache
	var dummy zend.Zval
	SplInstantiate(pce, retval)
	fci.size = g.SizeOf("zend_fcall_info")
	var __z *zend.Zval = &fci.function_name
	var __s *zend.ZendString = func_.common.function_name
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	fci.object = retval.value.obj
	fci.retval = &dummy
	fci.param_count = argc
	fci.params = argv
	fci.no_separation = 1
	fcc.function_handler = func_
	fcc.called_scope = pce
	fcc.object = retval.value.obj
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
	switch offset.u1.v.type_ {
	case 6:
		if zend._zendHandleNumericStr(offset.value.str.val, offset.value.str.len_, &idx) != 0 {
			return idx
		}
		break
	case 5:
		return zend_long(*offset).value.dval
	case 4:
		return offset.value.lval
	case 2:
		return 0
	case 3:
		return 1
	case 10:
		offset = &(*offset).value.ref.val
		goto try_again
	case 9:
		return offset.value.res.handle
	}
	return -1
}

/* }}} */
