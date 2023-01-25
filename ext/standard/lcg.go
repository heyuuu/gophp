// <<generate>>

package standard

import (
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

func MODMULT(a int, b int, c int, m int, s int) {
	q = s / a
	s = b*(s-a*q) - c*q
	if s < 0 {
		s += m
	}
}
func PhpCombinedLcg() float64 {
	var q int32
	var z int32
	if !(LCG(seeded)) {
		LcgSeed()
	}
	MODMULT(53668, 40014, 12211, 2147483563, LCG(s1))
	MODMULT(52774, 40692, 3791, 2147483399, LCG(s2))
	z = LCG(s1) - LCG(s2)
	if z < 1 {
		z += 2147483562
	}
	return z * 4.656613e-10
}

/* }}} */

func LcgSeed() {
	var tv __struct__timeval
	if gettimeofday(&tv, nil) == 0 {
		LCG(s1) = tv.tv_sec ^ tv.tv_usec<<11
	} else {
		LCG(s1) = 1
	}
	LCG(s2) = zend.ZendLong(getpid())

	/* Add entropy to s2 by calling gettimeofday() again */

	if gettimeofday(&tv, nil) == 0 {
		LCG(s2) ^= tv.tv_usec << 11
	}
	LCG(seeded) = 1
}

/* }}} */

func LcgInitGlobals(lcg_globals_p *PhpLcgGlobals) { LCG(seeded) = 0 }

/* }}} */

func ZmStartupLcg(type_ int, module_number int) int {
	LcgInitGlobals(&LcgGlobals)
	return zend.SUCCESS
}

/* }}} */

func ZifLcgValue(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_DOUBLE(PhpCombinedLcg())
	return
}

/* }}} */
