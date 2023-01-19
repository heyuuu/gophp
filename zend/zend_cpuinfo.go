// <<generate>>

package zend

// Source: <Zend/zend_cpuinfo.h>

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
   | Authors: Xinchen Hui <xinchen.h@zend.com>                            |
   +----------------------------------------------------------------------+
*/

// #define ZEND_CPU_INFO_H

// # include "zend.h"

// #define ZEND_CPU_EBX_MASK       ( 1 << 30 )

// #define ZEND_CPU_EDX_MASK       ( 1U << 31 )

type ZendCpuFeature = int

const (
	ZEND_CPU_FEATURE_SSE3         ZendCpuFeature = 1 << 0
	ZEND_CPU_FEATURE_PCLMULQDQ    ZendCpuFeature = 1 << 1
	ZEND_CPU_FEATURE_DTES64       ZendCpuFeature = 1 << 2
	ZEND_CPU_FEATURE_MONITOR      ZendCpuFeature = 1 << 3
	ZEND_CPU_FEATURE_DSCPL        ZendCpuFeature = 1 << 4
	ZEND_CPU_FEATURE_VMX          ZendCpuFeature = 1 << 5
	ZEND_CPU_FEATURE_SMX          ZendCpuFeature = 1 << 6
	ZEND_CPU_FEATURE_EST          ZendCpuFeature = 1 << 7
	ZEND_CPU_FEATURE_TM2          ZendCpuFeature = 1 << 8
	ZEND_CPU_FEATURE_SSSE3        ZendCpuFeature = 1 << 9
	ZEND_CPU_FEATURE_CID          ZendCpuFeature = 1 << 10
	ZEND_CPU_FEATURE_SDBG         ZendCpuFeature = 1 << 11
	ZEND_CPU_FEATURE_FMA          ZendCpuFeature = 1 << 12
	ZEND_CPU_FEATURE_CX16         ZendCpuFeature = 1 << 13
	ZEND_CPU_FEATURE_XTPR         ZendCpuFeature = 1 << 14
	ZEND_CPU_FEATURE_PDCM         ZendCpuFeature = 1 << 15
	ZEND_CPU_FEATURE_PCID         ZendCpuFeature = 1 << 17
	ZEND_CPU_FEATURE_DCA          ZendCpuFeature = 1 << 18
	ZEND_CPU_FEATURE_SSE41        ZendCpuFeature = 1 << 19
	ZEND_CPU_FEATURE_SSE42        ZendCpuFeature = 1 << 20
	ZEND_CPU_FEATURE_X2APIC       ZendCpuFeature = 1 << 21
	ZEND_CPU_FEATURE_MOVBE        ZendCpuFeature = 1 << 22
	ZEND_CPU_FEATURE_POPCNT       ZendCpuFeature = 1 << 23
	ZEND_CPU_FEATURE_TSC_DEADLINE ZendCpuFeature = 1 << 24
	ZEND_CPU_FEATURE_AES          ZendCpuFeature = 1 << 25
	ZEND_CPU_FEATURE_XSAVE        ZendCpuFeature = 1 << 26
	ZEND_CPU_FEATURE_OSXSAVE      ZendCpuFeature = 1 << 27
	ZEND_CPU_FEATURE_AVX          ZendCpuFeature = 1 << 28
	ZEND_CPU_FEATURE_F16C         ZendCpuFeature = 1 << 29
	ZEND_CPU_FEATURE_AVX2         ZendCpuFeature = 1<<5 | 1<<30
	ZEND_CPU_FEATURE_FPU          ZendCpuFeature = 1<<0 | 1<<31
	ZEND_CPU_FEATURE_VME          ZendCpuFeature = 1<<1 | 1<<31
	ZEND_CPU_FEATURE_DE           ZendCpuFeature = 1<<2 | 1<<31
	ZEND_CPU_FEATURE_PSE          ZendCpuFeature = 1<<3 | 1<<31
	ZEND_CPU_FEATURE_TSC          ZendCpuFeature = 1<<4 | 1<<31
	ZEND_CPU_FEATURE_MSR          ZendCpuFeature = 1<<5 | 1<<31
	ZEND_CPU_FEATURE_PAE          ZendCpuFeature = 1<<6 | 1<<31
	ZEND_CPU_FEATURE_MCE          ZendCpuFeature = 1<<7 | 1<<31
	ZEND_CPU_FEATURE_CX8          ZendCpuFeature = 1<<8 | 1<<31
	ZEND_CPU_FEATURE_APIC         ZendCpuFeature = 1<<9 | 1<<31
	ZEND_CPU_FEATURE_SEP          ZendCpuFeature = 1<<11 | 1<<31
	ZEND_CPU_FEATURE_MTRR         ZendCpuFeature = 1<<12 | 1<<31
	ZEND_CPU_FEATURE_PGE          ZendCpuFeature = 1<<13 | 1<<31
	ZEND_CPU_FEATURE_MCA          ZendCpuFeature = 1<<14 | 1<<31
	ZEND_CPU_FEATURE_CMOV         ZendCpuFeature = 1<<15 | 1<<31
	ZEND_CPU_FEATURE_PAT          ZendCpuFeature = 1<<16 | 1<<31
	ZEND_CPU_FEATURE_PSE36        ZendCpuFeature = 1<<17 | 1<<31
	ZEND_CPU_FEATURE_PN           ZendCpuFeature = 1<<18 | 1<<31
	ZEND_CPU_FEATURE_CLFLUSH      ZendCpuFeature = 1<<19 | 1<<31
	ZEND_CPU_FEATURE_DS           ZendCpuFeature = 1<<21 | 1<<31
	ZEND_CPU_FEATURE_ACPI         ZendCpuFeature = 1<<22 | 1<<31
	ZEND_CPU_FEATURE_MMX          ZendCpuFeature = 1<<23 | 1<<31
	ZEND_CPU_FEATURE_FXSR         ZendCpuFeature = 1<<24 | 1<<31
	ZEND_CPU_FEATURE_SSE          ZendCpuFeature = 1<<25 | 1<<31
	ZEND_CPU_FEATURE_SSE2         ZendCpuFeature = 1<<26 | 1<<31
	ZEND_CPU_FEATURE_SS           ZendCpuFeature = 1<<27 | 1<<31
	ZEND_CPU_FEATURE_HT           ZendCpuFeature = 1<<28 | 1<<31
	ZEND_CPU_FEATURE_TM           ZendCpuFeature = 1<<29 | 1<<31
)

/* Address sanitizer is incompatible with ifunc resolvers, so exclude the
 * CPU support helpers from asan.
 * See also https://github.com/google/sanitizers/issues/342. */

// #define ZEND_NO_SANITIZE_ADDRESS

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

// Source: <Zend/zend_cpuinfo.c>

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
   | Authors: Xinchen Hui <xinchen.h@zend.com>                            |
   +----------------------------------------------------------------------+
*/

// # include "zend_cpuinfo.h"

// @type ZendCpuInfo struct

var Cpuinfo ZendCpuInfo = ZendCpuInfo{0}

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
			Cpuinfo.SetEbx(Cpuinfo.GetEbx() &^ (ZEND_CPU_FEATURE_AVX2 & ^(1 << 30)))
		}
	}
}
func ZendCpuSupports(feature ZendCpuFeature) int {
	if (feature & 1 << 31) != 0 {
		return Cpuinfo.GetEdx() & (feature & ^(1 << 31))
	} else if (feature & 1 << 30) != 0 {
		return Cpuinfo.GetEbx() & (feature & ^(1 << 30))
	} else {
		return Cpuinfo.GetEcx() & feature
	}
}
