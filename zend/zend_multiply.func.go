package zend

import (
	"github.com/heyuuu/gophp/zend/faults"
)

func SignedMultiplyLong(a int, b int) (iVal int, dVal float64, overflow bool) {
	// ZEND_SIGNED_MULTIPLY_LONG
	iVal = a * b
	dVal = float64(a) * float64(b)
	delta := dVal - float64(iVal)
	if dVal+delta != dVal {
		return 0, dVal, true
	} else {
		return iVal, 0, false
	}
}

func ZendSafeAddress(nmemb int, size int, offset int, overflow *int) int {
	var res int = nmemb*size + offset
	var _d float64 = float64(nmemb)*float64(size) + float64(offset)
	var _delta float64 = float64(res) - _d
	if _d+_delta != _d {
		*overflow = 1
		return 0
	}
	*overflow = 0
	return res
}
func ZendSafeAddmult(nmemb int, size int, offset int, message string) int {
	var overflow int
	var ret int = ZendSafeAddress(nmemb, size, offset, &overflow)
	if overflow != 0 {
		faults.ErrorNoreturn(faults.E_ERROR, fmt.Sprintf("Possible integer overflow in %s (%zu * %zu + %zu)", message, nmemb, size, offset))
		return 0
	}
	return ret
}
