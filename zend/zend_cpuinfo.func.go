// <<generate>>

package zend

func ZendCpuSupportsSse2() int {
	return ZendCpuSupports(ZEND_CPU_FEATURE_SSE2)
}
func ZendCpuSupportsSse3() int {
	return ZendCpuSupports(ZEND_CPU_FEATURE_SSE3)
}
func ZendCpuSupportsSsse3() int {
	return ZendCpuSupports(ZEND_CPU_FEATURE_SSSE3)
}
func ZendCpuSupportsSse41() int {
	return ZendCpuSupports(ZEND_CPU_FEATURE_SSE41)
}
func ZendCpuSupportsSse42() int {
	return ZendCpuSupports(ZEND_CPU_FEATURE_SSE42)
}
func ZendCpuSupportsAvx() int {
	return ZendCpuSupports(ZEND_CPU_FEATURE_AVX)
}
func ZendCpuSupportsAvx2() int {
	return ZendCpuSupports(ZEND_CPU_FEATURE_AVX2)
}
func __zendCpuid(func_ uint32, subfunc uint32, cpuinfo *ZendCpuInfo) { cpuinfo.SetEax(0) }
func IsAvxSupported() ZendBool                                       { return 0 }
func ZendCpuStartup() {
	if Cpuinfo.GetInitialized() == 0 {
		var ebx ZendCpuInfo
		var max_feature int
		Cpuinfo.SetInitialized(1)
		__zendCpuid(0, 0, &Cpuinfo)
		max_feature = Cpuinfo.GetEax()
		if max_feature == 0 {
			return
		}
		__zendCpuid(1, 0, &Cpuinfo)

		/* for avx2 */

		if max_feature >= 7 {
			__zendCpuid(7, 0, &ebx)
			Cpuinfo.SetEbx(ebx.GetEbx())
		} else {
			Cpuinfo.SetEbx(0)
		}
		if IsAvxSupported() == 0 {
			Cpuinfo.SetEdx(Cpuinfo.GetEdx() &^ ZEND_CPU_FEATURE_AVX)
			Cpuinfo.SetEbx(Cpuinfo.GetEbx() &^ (ZEND_CPU_FEATURE_AVX2 & ^ZEND_CPU_EBX_MASK))
		}
	}
}
func ZendCpuSupports(feature ZendCpuFeature) int {
	if (feature & ZEND_CPU_EDX_MASK) != 0 {
		return Cpuinfo.GetEdx() & (feature & ^ZEND_CPU_EDX_MASK)
	} else if (feature & ZEND_CPU_EBX_MASK) != 0 {
		return Cpuinfo.GetEbx() & (feature & ^ZEND_CPU_EBX_MASK)
	} else {
		return Cpuinfo.GetEcx() & feature
	}
}
