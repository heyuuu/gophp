// <<generate>>

package spl

import (
	"sik/zend"
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

const spl_ce_Exception *zend.ZendClassEntry = zend.ZendCeException

/* {{{ PHP_MINIT_FUNCTION(spl_exceptions) */
