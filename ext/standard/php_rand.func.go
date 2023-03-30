package standard

import (
	"github.com/heyuuu/gophp/zend"
)

func RAND_RANGE_BADSCALING(__n int64, __min zend.ZendLong, __max zend.ZendLong, __tmax float) int64 {
	__n = __min + zend_long(float64(float64(__max-__min+1.0)*(__n/(__tmax+1.0))))
	return __n
}
func GENERATE_SEED() int {
	return zend_long(time(0)*getpid()) ^ zend_long(1000000.0*PhpCombinedLcg())
}
