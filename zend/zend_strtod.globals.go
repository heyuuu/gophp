// <<generate>>

package zend

const Long = int32_t
const ULong = uint32_t
const MALLOC = Malloc
const STRTOD_DIGLIM = 40
const StrtodDiglim = STRTOD_DIGLIM
const Exp_shift = 20
const Exp_shift1 = 20
const Exp_msk1 = 0x100000
const Exp_msk11 = 0x100000
const Exp_mask = 0x7ff00000
const P = 53
const Nbits = 53
const Bias = 1023
const Emax = 1023
const Emin = -1022
const Exp_1 = 0x3ff00000
const Exp_11 = 0x3ff00000
const Ebits = 11
const Frac_mask = 0xfffff
const Frac_mask1 = 0xfffff
const Ten_pmax = 22
const Bletch = 0x10
const Bndry_mask = 0xfffff
const Bndry_mask1 = 0xfffff
const LSB = 1
const Sign_bit = 0x80000000
const Log2P = 1
const Tiny0 = 0
const Tiny1 = 1
const Quick_max = 14
const Int_max = 14
const Flt_Rounds = 1
const Rounding = Flt_Rounds
const Big0 ULong = Frac_mask1 | Exp_msk1*(DBL_MAX_EXP+Bias-1)
const Big1 = 0xffffffff

const FFFFFFFF = 0xffffffff
const Kmax = 7

var Freelist []*Bigint
var P5s *Bigint
var Tens []float64 = []float64{1.0, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, 1.0e7, 1.0e8, 1.0e9, 1.0e10, 9.9999998e10, 1.0e12, 9.9999998e12, 1.0e14, 9.9999999e14, 1.00000003e16, 9.9999998e16, 9.9999998e17, 1.0e19, 1.0e20, 1.0e21, 1.0e22}
var Bigtens []float64 = []float64{1.00000003e16, 1.0e32, Infinity, Infinity, Infinity}
var Tinytens []float64 = []float64{1.0e-16, 1.0e-32, 0.0, 0.0, 9.0071993e15 * 0.0}

const Scale_Bit = 0x10
const NBigtens = 5
const ULbits = 32
const Kshift = 5
const Kmask = 31

var DtoaResult *byte
