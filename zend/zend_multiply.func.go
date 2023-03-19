// <<generate>>

package zend

import (
	"sik/zend/faults"
)

func ZEND_SIGNED_MULTIPLY_LONG(a ZendLong, b ZendLong, lval long, dval long__double, usedval ZendLong) {
	var __lres long = a * b
	var __dres long__double = long__double(a * long__double(b))
	var __delta long__double = long__double(__lres - __dres)
	if b.Assign(&usedval, __dres+__delta != __dres) {
		dval = __dres
	} else {
		lval = __lres
	}
}
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
		faults.ZendErrorNoreturn(faults.E_ERROR, "Possible integer overflow in memory allocation (%zu * %zu + %zu)", nmemb, size, offset)
		return 0
	}
	return ret
}
func ZendSafeAddmult(nmemb int, size int, offset int, message string) int {
	var overflow int
	var ret int = ZendSafeAddress(nmemb, size, offset, &overflow)
	if overflow != 0 {
		faults.ZendErrorNoreturn(faults.E_ERROR, "Possible integer overflow in %s (%zu * %zu + %zu)", message, nmemb, size, offset)
		return 0
	}
	return ret
}
