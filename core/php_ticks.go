// <<generate>>

package core

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/php_ticks.h>

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
   | Author: Stig Bakken <ssb@php.net>                                    |
   +----------------------------------------------------------------------+
*/

// #define PHP_TICKS_H

// Source: <main/php_ticks.c>

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
   | Author: Stig Bakken <ssb@php.net>                                    |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_ticks.h"

// @type StTickFunction struct
func PhpStartupTicks() int {
	zend.ZendLlistInit(&(CoreGlobals.GetTickFunctions()), g.SizeOf("struct st_tick_function"), nil, 1)
	return zend.SUCCESS
}
func PhpDeactivateTicks() {
	zend.ZendLlistClean(&(CoreGlobals.GetTickFunctions()))
}
func PhpShutdownTicks() {
	zend.ZendLlistDestroy(&(CoreGlobals.GetTickFunctions()))
}
func PhpCompareTickFunctions(elem1 any, elem2 any) int {
	var e1 *StTickFunction = (*StTickFunction)(elem1)
	var e2 *StTickFunction = (*StTickFunction)(elem2)
	return e1.GetFunc() == e2.GetFunc() && e1.GetArg() == e2.GetArg()
}
func PhpAddTickFunction(func_ func(int, any), arg any) {
	var tmp StTickFunction = StTickFunction{func_, arg}
	zend.ZendLlistAddElement(&(CoreGlobals.GetTickFunctions()), any(&tmp))
}
func PhpRemoveTickFunction(func_ func(int, any), arg any) {
	var tmp StTickFunction = StTickFunction{func_, arg}
	zend.ZendLlistDelElement(&(CoreGlobals.GetTickFunctions()), any(&tmp), (func(any, any) int)(PhpCompareTickFunctions))
}
func PhpTickIterator(d any, arg any) {
	var data *StTickFunction = (*StTickFunction)(d)
	data.GetFunc()(*((*int)(arg)), data.GetArg())
}
func PhpRunTicks(count int) {
	zend.ZendLlistApplyWithArgument(&(CoreGlobals.GetTickFunctions()), zend.LlistApplyWithArgFuncT(PhpTickIterator), &count)
}
