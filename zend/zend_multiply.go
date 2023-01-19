// <<generate>>

package zend

// Source: <Zend/zend_multiply.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Sascha Schumann <sascha@schumann.cx>                        |
   |          Ard Biesheuvel <ard.biesheuvel@linaro.org>                  |
   +----------------------------------------------------------------------+
*/

// # include "zend_portability.h"

// #define ZEND_MULTIPLY_H

// #define ZEND_SIGNED_MULTIPLY_LONG(a,b,lval,dval,usedval) do { long __lres = ( a ) * ( b ) ; long double __dres = ( long double ) ( a ) * ( long double ) ( b ) ; long double __delta = ( long double ) __lres - __dres ; if ( ( ( usedval ) = ( ( __dres + __delta ) != __dres ) ) ) { ( dval ) = __dres ; } else { ( lval ) = __lres ; } } while ( 0 )

func ZendSafeAddress(nmemb int, size int, offset int, overflow *int) int {
	var res int = nmemb*size + offset
	var _d float64 = float64(nmemb * float64(size+float64(offset)))
	var _delta float64 = float64(res - _d)
	if _d+_delta != _d {
		*overflow = 1
		return 0
	}
	*overflow = 0
	return res
}
func ZendSafeAddressGuarded(nmemb int, size int, offset int) int {
	var overflow int
	var ret int = ZendSafeAddress(nmemb, size, offset, &overflow)
	if overflow != 0 {
		ZendErrorNoreturn(1<<0, "Possible integer overflow in memory allocation (%zu * %zu + %zu)", nmemb, size, offset)
		return 0
	}
	return ret
}

/* A bit more generic version of the same */

func ZendSafeAddmult(nmemb int, size int, offset int, message string) int {
	var overflow int
	var ret int = ZendSafeAddress(nmemb, size, offset, &overflow)
	if overflow != 0 {
		ZendErrorNoreturn(1<<0, "Possible integer overflow in %s (%zu * %zu + %zu)", message, nmemb, size, offset)
		return 0
	}
	return ret
}
