// <<generate>>

package zend

const ZEND_CPU_EBX_MASK = 1 << 30
const ZEND_CPU_EDX_MASK = 1 << 31

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
	ZEND_CPU_FEATURE_AVX2         ZendCpuFeature = 1<<5 | ZEND_CPU_EBX_MASK
	ZEND_CPU_FEATURE_FPU          ZendCpuFeature = 1<<0 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_VME          ZendCpuFeature = 1<<1 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_DE           ZendCpuFeature = 1<<2 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_PSE          ZendCpuFeature = 1<<3 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_TSC          ZendCpuFeature = 1<<4 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_MSR          ZendCpuFeature = 1<<5 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_PAE          ZendCpuFeature = 1<<6 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_MCE          ZendCpuFeature = 1<<7 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_CX8          ZendCpuFeature = 1<<8 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_APIC         ZendCpuFeature = 1<<9 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_SEP          ZendCpuFeature = 1<<11 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_MTRR         ZendCpuFeature = 1<<12 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_PGE          ZendCpuFeature = 1<<13 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_MCA          ZendCpuFeature = 1<<14 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_CMOV         ZendCpuFeature = 1<<15 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_PAT          ZendCpuFeature = 1<<16 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_PSE36        ZendCpuFeature = 1<<17 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_PN           ZendCpuFeature = 1<<18 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_CLFLUSH      ZendCpuFeature = 1<<19 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_DS           ZendCpuFeature = 1<<21 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_ACPI         ZendCpuFeature = 1<<22 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_MMX          ZendCpuFeature = 1<<23 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_FXSR         ZendCpuFeature = 1<<24 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_SSE          ZendCpuFeature = 1<<25 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_SSE2         ZendCpuFeature = 1<<26 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_SS           ZendCpuFeature = 1<<27 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_HT           ZendCpuFeature = 1<<28 | ZEND_CPU_EDX_MASK
	ZEND_CPU_FEATURE_TM           ZendCpuFeature = 1<<29 | ZEND_CPU_EDX_MASK
)

var Cpuinfo ZendCpuInfo = ZendCpuInfo{0}
