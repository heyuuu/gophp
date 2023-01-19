// <<generate>>

package spl

import "sik/zend"

// Source: <ext/spl/spl_exceptions.h>

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

// #define SPL_EXCEPTIONS_H

// # include "php.h"

// # include "php_spl.h"

var spl_ce_LogicException *zend.ZendClassEntry
var spl_ce_BadFunctionCallException *zend.ZendClassEntry
var spl_ce_BadMethodCallException *zend.ZendClassEntry
var spl_ce_DomainException *zend.ZendClassEntry
var spl_ce_InvalidArgumentException *zend.ZendClassEntry
var spl_ce_LengthException *zend.ZendClassEntry
var spl_ce_OutOfRangeException *zend.ZendClassEntry
var spl_ce_RuntimeException *zend.ZendClassEntry
var spl_ce_OutOfBoundsException *zend.ZendClassEntry
var spl_ce_OverflowException *zend.ZendClassEntry
var spl_ce_RangeException *zend.ZendClassEntry
var spl_ce_UnderflowException *zend.ZendClassEntry
var spl_ce_UnexpectedValueException *zend.ZendClassEntry

// Source: <ext/spl/spl_exceptions.c>

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

// # include "zend_exceptions.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_exceptions.h"

// #define spl_ce_Exception       zend_ce_exception

/* {{{ PHP_MINIT_FUNCTION(spl_exceptions) */

func ZmStartupSplExceptions(type_ int, module_number int) int {
	SplRegisterSubClass(&spl_ce_LogicException, zend.ZendCeException, "LogicException", nil, nil)
	SplRegisterSubClass(&spl_ce_BadFunctionCallException, spl_ce_LogicException, "BadFunctionCallException", nil, nil)
	SplRegisterSubClass(&spl_ce_BadMethodCallException, spl_ce_BadFunctionCallException, "BadMethodCallException", nil, nil)
	SplRegisterSubClass(&spl_ce_DomainException, spl_ce_LogicException, "DomainException", nil, nil)
	SplRegisterSubClass(&spl_ce_InvalidArgumentException, spl_ce_LogicException, "InvalidArgumentException", nil, nil)
	SplRegisterSubClass(&spl_ce_LengthException, spl_ce_LogicException, "LengthException", nil, nil)
	SplRegisterSubClass(&spl_ce_OutOfRangeException, spl_ce_LogicException, "OutOfRangeException", nil, nil)
	SplRegisterSubClass(&spl_ce_RuntimeException, zend.ZendCeException, "RuntimeException", nil, nil)
	SplRegisterSubClass(&spl_ce_OutOfBoundsException, spl_ce_RuntimeException, "OutOfBoundsException", nil, nil)
	SplRegisterSubClass(&spl_ce_OverflowException, spl_ce_RuntimeException, "OverflowException", nil, nil)
	SplRegisterSubClass(&spl_ce_RangeException, spl_ce_RuntimeException, "RangeException", nil, nil)
	SplRegisterSubClass(&spl_ce_UnderflowException, spl_ce_RuntimeException, "UnderflowException", nil, nil)
	SplRegisterSubClass(&spl_ce_UnexpectedValueException, spl_ce_RuntimeException, "UnexpectedValueException", nil, nil)
	return zend.SUCCESS
}

/* }}} */
