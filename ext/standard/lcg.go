// <<generate>>

package standard

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/lcg.c>

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
   | Author: Sascha Schumann <sascha@schumann.cx>                         |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_lcg.h"

// # include < unistd . h >

// # include < sys / time . h >

var LcgGlobals PhpLcgGlobals

/*
 * combinedLCG() returns a pseudo random number in the range of (0, 1).
 * The function combines two CGs with periods of
 * 2^31 - 85 and 2^31 - 249. The period of this function
 * is equal to the product of both primes.
 */

// #define MODMULT(a,b,c,m,s) q = s / a ; s = b * ( s - a * q ) - c * q ; if ( s < 0 ) s += m

func PhpCombinedLcg() float64 {
	var q int32
	var z int32
	if LcgGlobals.GetSeeded() == 0 {
		LcgSeed()
	}
	q = LcgGlobals.GetS1() / 53668
	LcgGlobals.SetS1(40014*(LcgGlobals.GetS1()-53668*q) - 12211*q)
	if LcgGlobals.GetS1() < 0 {
		LcgGlobals.SetS1(LcgGlobals.GetS1() + 2147483563)
	}
	q = LcgGlobals.GetS2() / 52774
	LcgGlobals.SetS2(40692*(LcgGlobals.GetS2()-52774*q) - 3791*q)
	if LcgGlobals.GetS2() < 0 {
		LcgGlobals.SetS2(LcgGlobals.GetS2() + 2147483399)
	}
	z = LcgGlobals.GetS1() - LcgGlobals.GetS2()
	if z < 1 {
		z += 2147483562
	}
	return z * 4.656613e-10
}

/* }}} */

func LcgSeed() {
	var tv __struct__timeval
	if gettimeofday(&tv, nil) == 0 {
		LcgGlobals.SetS1(tv.tv_sec ^ tv.tv_usec<<11)
	} else {
		LcgGlobals.SetS1(1)
	}
	LcgGlobals.SetS2(zend.ZendLong(getpid()))

	/* Add entropy to s2 by calling gettimeofday() again */

	if gettimeofday(&tv, nil) == 0 {
		LcgGlobals.SetS2(LcgGlobals.GetS2() ^ tv.tv_usec<<11)
	}
	LcgGlobals.SetSeeded(1)
}

/* }}} */

func LcgInitGlobals(lcg_globals_p *PhpLcgGlobals) { LcgGlobals.SetSeeded(0) }

/* }}} */

func ZmStartupLcg(type_ int, module_number int) int {
	LcgInitGlobals(&LcgGlobals)
	return zend.SUCCESS
}

/* }}} */

func ZifLcgValue(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.dval = PhpCombinedLcg()
	__z.u1.type_info = 5
	return
}

/* }}} */
