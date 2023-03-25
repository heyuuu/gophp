package standard

import (
	"sik/zend"
	"sik/zend/types"
	"sik/zend/zpp"
)

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
func LcgInitGlobals(lcg_globals_p *PhpLcgGlobals) { LCG(seeded) = 0 }
func ZmStartupLcg(type_ int, module_number int) int {
	LcgInitGlobals(&LcgGlobals)
	return types.SUCCESS
}
func ZifLcgValue(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetDouble(PhpCombinedLcg())
	return
}
