// <<generate>>

package zend

// Source: <Zend/zend_strtod.h>

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
   | Authors: Derick Rethans <derick@php.net>                             |
   +----------------------------------------------------------------------+
*/

// Source: <Zend/zend_strtod.c>

/****************************************************************
 *
 * The author of this software is David M. Gay.
 *
 * Copyright (c) 1991, 2000, 2001 by Lucent Technologies.
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose without fee is hereby granted, provided that this entire notice
 * is included in all copies of any software which is or includes a copy
 * or modification of this software and in all copies of the supporting
 * documentation for such software.
 *
 * THIS SOFTWARE IS BEING PROVIDED "AS IS", WITHOUT ANY EXPRESS OR IMPLIED
 * WARRANTY.  IN PARTICULAR, NEITHER THE AUTHOR NOR LUCENT MAKES ANY
 * REPRESENTATION OR WARRANTY OF ANY KIND CONCERNING THE MERCHANTABILITY
 * OF THIS SOFTWARE OR ITS FITNESS FOR ANY PARTICULAR PURPOSE.
 *
 ***************************************************************/

const Long = int32_t
const ULong = uint32_t
const MALLOC = Malloc
const STRTOD_DIGLIM = 40
const StrtodDiglim = STRTOD_DIGLIM

/* The following definition of Storeinc is appropriate for MIPS processors.
 * An alternative that might be better on some machines is
 * #define Storeinc(a,b,c) (*a++ = b << 16 | c & 0xffff)
 */

/* #define P __special__  DBL_MANT_DIG */

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

/* The factor of 2^53 in tinytens[4] helps us avoid setting the underflow */

const Scale_Bit = 0x10
const NBigtens = 5
const ULbits = 32
const Kshift = 5
const Kmask = 31

var DtoaResult *byte

/* freedtoa(s) must be used to free values s returned by dtoa
 * when MULTIPLE_THREADS is #defined.  It should be used in all cases,
 * but for consistency with earlier versions of dtoa, it is optional
 * when MULTIPLE_THREADS is not defined.
 */

/* dtoa for IEEE arithmetic (dmg): convert double to ASCII string.
 *
 * Inspired by "How to Print Floating-Point Numbers Accurately" by
 * Guy L. Steele, Jr. and Jon L. White [Proc. ACM SIGPLAN '90, pp. 112-126].
 *
 * Modifications:
 *    1. Rather than iterating, we use a simple numeric overestimate
 *       to determine k = floor(log10(d)).  We scale relevant
 *       quantities using O(log2(k)) rather than O(k) multiplications.
 *    2. For some modes > 2 (corresponding to ecvt and fcvt), we don't
 *       try to generate digits strictly left to right.  Instead, we
 *       compute with fewer bits and propagate the carry if necessary
 *       when rounding the final digit up.  This is often faster.
 *    3. Under the assumption that input will be rounded nearest,
 *       mode 0 renders 1e23 as 1e23 rather than 9.999999999999999e22.
 *       That is, we allow equality in stopping tests when the
 *       round-nearest rule will give the same floating-point value
 *       as would satisfaction of the stopping test with strict
 *       inequality.
 *    4. We remove common factors of powers of 2 from relevant
 *       quantities.
 *    5. When converting floating-point integers less than 1e16,
 *       we use floating-point arithmetic rather than resorting
 *       to multiple-precision integers.
 *    6. When asked to produce fewer than 15 digits, we first try
 *       to get by with floating-point arithmetic; we resort to
 *       multiple-precision integer arithmetic only if we cannot
 *       guarantee that the floating-point calculation has given
 *       the correctly rounded result.  For k requested digits and
 *       "uniformly" distributed input, the probability is
 *       something like 10^(k-15) that we must resort to the Long
 *       calculation.
 */
