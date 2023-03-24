package spl

import (
	"sik/zend/faults"
	"sik/zend/types"
)

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

var spl_ce_LogicException *types.ClassEntry
var spl_ce_BadFunctionCallException *types.ClassEntry
var spl_ce_BadMethodCallException *types.ClassEntry
var spl_ce_DomainException *types.ClassEntry
var spl_ce_InvalidArgumentException *types.ClassEntry
var spl_ce_LengthException *types.ClassEntry
var spl_ce_OutOfRangeException *types.ClassEntry
var spl_ce_RuntimeException *types.ClassEntry
var spl_ce_OutOfBoundsException *types.ClassEntry
var spl_ce_OverflowException *types.ClassEntry
var spl_ce_RangeException *types.ClassEntry
var spl_ce_UnderflowException *types.ClassEntry
var spl_ce_UnexpectedValueException *types.ClassEntry

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

const spl_ce_Exception *types.ClassEntry = faults.ZendCeException

/* {{{ PHP_MINIT_FUNCTION(spl_exceptions) */
