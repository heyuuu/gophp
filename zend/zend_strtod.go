// <<generate>>

package zend

import g "sik/runtime/grammar"

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

// #define ZEND_STRTOD_H

// # include < zend . h >

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

// # include < zend_operators . h >

// # include < zend_strtod . h >

// # include "zend_strtod_int.h"

// #define Long       int32_t

// #define ULong       uint32_t

// # include "stdlib.h"

// # include "string.h"

// #define MALLOC       malloc

// #define IEEE_Arith

// # include "errno.h"

// # include "float.h"

// # include "math.h"

// #define CONST       const

// @type U struct

// #define word0(x) ( x ) -> L [ 1 ]

// #define word1(x) ( x ) -> L [ 0 ]

// #define dval(x) ( x ) -> d

// #define STRTOD_DIGLIM       40

// #define strtod_diglim       STRTOD_DIGLIM

/* The following definition of Storeinc is appropriate for MIPS processors.
 * An alternative that might be better on some machines is
 * #define Storeinc(a,b,c) (*a++ = b << 16 | c & 0xffff)
 */

// #define Storeinc(a,b,c) ( ( ( unsigned short * ) a ) [ 1 ] = ( unsigned short ) b , ( ( unsigned short * ) a ) [ 0 ] = ( unsigned short ) c , a ++ )

/* #define P __special__  DBL_MANT_DIG */

// #define Exp_shift       20

// #define Exp_shift1       20

// #define Exp_msk1       0x100000

// #define Exp_msk11       0x100000

// #define Exp_mask       0x7ff00000

// #define P       53

// #define Nbits       53

// #define Bias       1023

// #define Emax       1023

// #define Emin       ( - 1022 )

// #define Exp_1       0x3ff00000

// #define Exp_11       0x3ff00000

// #define Ebits       11

// #define Frac_mask       0xfffff

// #define Frac_mask1       0xfffff

// #define Ten_pmax       22

// #define Bletch       0x10

// #define Bndry_mask       0xfffff

// #define Bndry_mask1       0xfffff

// #define LSB       1

// #define Sign_bit       0x80000000

// #define Log2P       1

// #define Tiny0       0

// #define Tiny1       1

// #define Quick_max       14

// #define Int_max       14

// #define Avoid_Underflow

// #define Flt_Rounds       1

// #define Rounding       Flt_Rounds

// #define rounded_product(a,b) a *= b

// #define rounded_quotient(a,b) a /= b

// #define Big0       ( Frac_mask1 | Exp_msk1 * ( DBL_MAX_EXP + Bias - 1 ) )

// #define Big1       0xffffffff

// #define Pack_32

// @type BCinfo struct

// #define FFFFFFFF       0xffffffffUL

// #define Llong       long long

// #define ULLong       unsigned Llong

// #define ACQUIRE_DTOA_LOCK(n)

// #define FREE_DTOA_LOCK(n)

// #define Kmax       7

// @type Bigint struct

var Freelist []*Bigint

func ZendStartupStrtod() int { return 1 }

/* }}} */

func ZendShutdownStrtod() int {
	DestroyFreelist()
	FreeP5s()
	return 1
}

/* }}} */

func Balloc(k int) *Bigint {
	var x int
	var rv *Bigint

	/* The k > Kmax case does not need ACQUIRE_DTOA_LOCK(0), */

	if k <= 7 && g.Assign(&rv, Freelist[k]) {
		Freelist[k] = rv.GetNext()
	} else {
		x = 1 << k
		rv = (*Bigint)(Malloc(g.SizeOf("Bigint") + (x-1)*g.SizeOf("ULong")))
		if rv == nil {
			ZendErrorNoreturn(1<<0, "Balloc() failed to allocate memory")
		}
		rv.SetK(k)
		rv.SetMaxwds(x)
	}
	rv.SetWds(0)
	rv.SetSign(rv.GetWds())
	return rv
}
func Bfree(v *Bigint) {
	if v != nil {
		if v.GetK() > 7 {
			Free(any(v))
		} else {
			v.SetNext(Freelist[v.GetK()])
			Freelist[v.GetK()] = v
		}
	}
}

// #define Bcopy(x,y) memcpy ( ( char * ) & x -> sign , ( char * ) & y -> sign , y -> wds * sizeof ( Long ) + 2 * sizeof ( int ) )

func Multadd(b *Bigint, m int, a int) *Bigint {
	var i int
	var wds int
	var x *uint32
	var carry unsigned__long__long
	var y unsigned__long__long
	var b1 *Bigint
	wds = b.GetWds()
	x = b.GetX()
	i = 0
	carry = a
	for {
		y = (*x) * unsigned__long__long(m+carry)
		carry = y >> 32
		g.PostInc(&(*x)) = y & 0xffffffff
		if g.PreInc(&i) >= wds {
			break
		}
	}
	if carry {
		if wds >= b.GetMaxwds() {
			b1 = Balloc(b.GetK() + 1)
			memcpy((*byte)(&b1.sign), (*byte)(&b.sign), b.GetWds()*g.SizeOf("Long")+2*g.SizeOf("int"))
			Bfree(b)
			b = b1
		}
		b.GetX()[g.PostInc(&wds)] = carry
		b.SetWds(wds)
	}
	return b
}
func S2b(s *byte, nd0 int, nd int, y9 uint32, dplen int) *Bigint {
	var b *Bigint
	var i int
	var k int
	var x int32
	var y int32
	x = (nd + 8) / 9
	k = 0
	y = 1
	for x > y {
		y <<= 1
		k++
	}
	b = Balloc(k)
	b.GetX()[0] = y9
	b.SetWds(1)
	i = 9
	if 9 < nd0 {
		s += 9
		for {
			b = Multadd(b, 10, g.PostInc(&(*s))-'0')
			if g.PreInc(&i) >= nd0 {
				break
			}
		}
		s += dplen
	} else {
		s += dplen + 9
	}
	for ; i < nd; i++ {
		b = Multadd(b, 10, g.PostInc(&(*s))-'0')
	}
	return b
}
func Hi0bits(x uint32) int {
	var k int = 0
	if (x & 0xffff0000) == 0 {
		k = 16
		x <<= 16
	}
	if (x & 0xff000000) == 0 {
		k += 8
		x <<= 8
	}
	if (x & 0xf0000000) == 0 {
		k += 4
		x <<= 4
	}
	if (x & 0xc0000000) == 0 {
		k += 2
		x <<= 2
	}
	if (x & 0x80000000) == 0 {
		k++
		if (x & 0x40000000) == 0 {
			return 32
		}
	}
	return k
}
func Lo0bits(y *uint32) int {
	var k int
	var x uint32 = *y
	if (x & 7) != 0 {
		if (x & 1) != 0 {
			return 0
		}
		if (x & 2) != 0 {
			*y = x >> 1
			return 1
		}
		*y = x >> 2
		return 2
	}
	k = 0
	if (x & 0xffff) == 0 {
		k = 16
		x >>= 16
	}
	if (x & 0xff) == 0 {
		k += 8
		x >>= 8
	}
	if (x & 0xf) == 0 {
		k += 4
		x >>= 4
	}
	if (x & 0x3) == 0 {
		k += 2
		x >>= 2
	}
	if (x & 1) == 0 {
		k++
		x >>= 1
		if x == 0 {
			return 32
		}
	}
	*y = x
	return k
}
func I2b(i int) *Bigint {
	var b *Bigint
	b = Balloc(1)
	b.GetX()[0] = i
	b.SetWds(1)
	return b
}
func Mult(a *Bigint, b *Bigint) *Bigint {
	var c *Bigint
	var k int
	var wa int
	var wb int
	var wc int
	var x *uint32
	var xa *uint32
	var xae *uint32
	var xb *uint32
	var xbe *uint32
	var xc *uint32
	var xc0 *uint32
	var y uint32
	var carry unsigned__long__long
	var z unsigned__long__long
	if a.GetWds() < b.GetWds() {
		c = a
		a = b
		b = c
	}
	k = a.GetK()
	wa = a.GetWds()
	wb = b.GetWds()
	wc = wa + wb
	if wc > a.GetMaxwds() {
		k++
	}
	c = Balloc(k)
	x = c.GetX()
	xa = x + wc
	for ; x < xa; x++ {
		*x = 0
	}
	xa = a.GetX()
	xae = xa + wa
	xb = b.GetX()
	xbe = xb + wb
	xc0 = c.GetX()
	for ; xb < xbe; xc0++ {
		if g.Assign(&y, g.PostInc(&(*xb))) {
			x = xa
			xc = xc0
			carry = 0
			for {
				z = g.PostInc(&(*x)) * unsigned__long__long(y+(*xc)+carry)
				carry = z >> 32
				g.PostInc(&(*xc)) = z & 0xffffffff
				if x >= xae {
					break
				}
			}
			*xc = carry
		}
	}
	xc0 = c.GetX()
	xc = xc0 + wc
	for ; wc > 0 && !(*(g.PreDec(&xc))); wc-- {

	}
	c.SetWds(wc)
	return c
}

var P5s *Bigint

func Pow5mult(b *Bigint, k int) *Bigint {
	var b1 *Bigint
	var p5 *Bigint
	var p51 *Bigint
	var i int
	var p05 []int = []int{5, 25, 125}
	if g.Assign(&i, k&3) {
		b = Multadd(b, p05[i-1], 0)
	}
	if !(g.AssignOp(&k, ">>=", 2)) {
		return b
	}
	if !(g.Assign(&p5, P5s)) {

		/* first time */

		P5s = I2b(625)
		p5 = P5s
		p5.SetNext(0)
	}
	for {
		if (k & 1) != 0 {
			b1 = Mult(b, p5)
			Bfree(b)
			b = b1
		}
		if !(g.AssignOp(&k, ">>=", 1)) {
			break
		}
		if !(g.Assign(&p51, p5.GetNext())) {
			p5.SetNext(Mult(p5, p5))
			p51 = p5.GetNext()
			p51.SetNext(0)
		}
		p5 = p51
	}
	return b
}
func Lshift(b *Bigint, k int) *Bigint {
	var i int
	var k1 int
	var n int
	var n1 int
	var b1 *Bigint
	var x *uint32
	var x1 *uint32
	var xe *uint32
	var z uint32
	n = k >> 5
	k1 = b.GetK()
	n1 = n + b.GetWds() + 1
	for i = b.GetMaxwds(); n1 > i; i <<= 1 {
		k1++
	}
	b1 = Balloc(k1)
	x1 = b1.GetX()
	for i = 0; i < n; i++ {
		g.PostInc(&(*x1)) = 0
	}
	x = b.GetX()
	xe = x + b.GetWds()
	if g.AssignOp(&k, "&=", 0x1f) {
		k1 = 32 - k
		z = 0
		for {
			g.PostInc(&(*x1)) = (*x)<<k | z
			z = g.PostInc(&(*x)) >> k1
			if x >= xe {
				break
			}
		}
		if g.Assign(&(*x1), z) {
			n1++
		}
	} else {
		for {
			*x++
			g.PostInc(&(*x1)) = (*x) - 1
			if x >= xe {
				break
			}
		}
	}
	b1.SetWds(n1 - 1)
	Bfree(b)
	return b1
}
func Cmp(a *Bigint, b *Bigint) int {
	var xa *uint32
	var xa0 *uint32
	var xb *uint32
	var xb0 *uint32
	var i int
	var j int
	i = a.GetWds()
	j = b.GetWds()
	if g.AssignOp(&i, "-=", j) {
		return i
	}
	xa0 = a.GetX()
	xa = xa0 + j
	xb0 = b.GetX()
	xb = xb0 + j
	for {
		if (*(g.PreDec(&xa))) != (*(g.PreDec(&xb))) {
			if (*xa) < (*xb) {
				return -1
			} else {
				return 1
			}
		}
		if xa <= xa0 {
			break
		}
	}
	return 0
}
func Diff(a *Bigint, b *Bigint) *Bigint {
	var c *Bigint
	var i int
	var wa int
	var wb int
	var xa *uint32
	var xae *uint32
	var xb *uint32
	var xbe *uint32
	var xc *uint32
	var borrow unsigned__long__long
	var y unsigned__long__long
	i = Cmp(a, b)
	if i == 0 {
		c = Balloc(0)
		c.SetWds(1)
		c.GetX()[0] = 0
		return c
	}
	if i < 0 {
		c = a
		a = b
		b = c
		i = 1
	} else {
		i = 0
	}
	c = Balloc(a.GetK())
	c.SetSign(i)
	wa = a.GetWds()
	xa = a.GetX()
	xae = xa + wa
	wb = b.GetWds()
	xb = b.GetX()
	xbe = xb + wb
	xc = c.GetX()
	borrow = 0
	for {
		y = unsigned__long__long(g.PostInc(&(*xa)) - g.PostInc(&(*xb)) - borrow)
		borrow = y >> 32 & uint32(1)
		g.PostInc(&(*xc)) = y & 0xffffffff
		if xb >= xbe {
			break
		}
	}
	for xa < xae {
		y = g.PostInc(&(*xa)) - borrow
		borrow = y >> 32 & uint32(1)
		g.PostInc(&(*xc)) = y & 0xffffffff
	}
	for !(*(g.PreDec(&xc))) {
		wa--
	}
	c.SetWds(wa)
	return c
}
func Ulp(x *U) float64 {
	var L int32
	var u U
	L = (x.GetL()[1] & 0x7ff00000) - (53-1)*0x100000
	&u.GetL()[1] = L
	&u.GetL()[0] = 0
	return &u.GetD()
}
func B2d(a *Bigint, e *int) float64 {
	var xa *uint32
	var xa0 *uint32
	var w uint32
	var y uint32
	var z uint32
	var k int
	var d U

	// #define d0       word0 ( & d )

	// #define d1       word1 ( & d )

	xa0 = a.GetX()
	xa = xa0 + a.GetWds()
	y = *(g.PreDec(&xa))
	k = Hi0bits(y)
	*e = 32 - k
	if k < 11 {
		&d.GetL()[1] = 0x3ff00000 | y>>11 - k
		if xa > xa0 {
			w = *(g.PreDec(&xa))
		} else {
			w = 0
		}
		&d.GetL()[0] = y<<32 - 11 + k | w>>11 - k
		goto ret_d
	}
	if xa > xa0 {
		z = *(g.PreDec(&xa))
	} else {
		z = 0
	}
	if g.AssignOp(&k, "-=", 11) {
		&d.GetL()[1] = 0x3ff00000 | y<<k | z>>32 - k
		if xa > xa0 {
			y = *(g.PreDec(&xa))
		} else {
			y = 0
		}
		&d.GetL()[0] = z<<k | y>>32 - k
	} else {
		&d.GetL()[1] = 0x3ff00000 | y
		&d.GetL()[0] = z
	}
ret_d:
	return &d.GetD()
}
func D2b(d *U, e *int, bits *int) *Bigint {
	var b *Bigint
	var de int
	var k int
	var x *uint32
	var y uint32
	var z uint32
	var i int

	// #define d0       word0 ( d )

	// #define d1       word1 ( d )

	b = Balloc(1)
	x = b.GetX()
	z = d.GetL()[1] & 0xfffff
	d.GetL()[1] &= 0x7fffffff
	if g.Assign(&de, int(d.GetL()[1]>>20)) {
		z |= 0x100000
	}
	if g.Assign(&y, d.GetL()[0]) {
		if g.Assign(&k, Lo0bits(&y)) {
			x[0] = y | z<<32 - k
			z >>= k
		} else {
			x[0] = y
		}
		if g.Assign(&x[1], z) {
			b.SetWds(2)
		} else {
			b.SetWds(1)
		}
		i = b.GetWds()
	} else {
		k = Lo0bits(&z)
		x[0] = z
		b.SetWds(1)
		i = b.GetWds()
		k += 32
	}
	if de != 0 {
		*e = de - 1023 - (53 - 1) + k
		*bits = 53 - k
	} else {
		*e = de - 1023 - (53 - 1) + 1 + k
		*bits = 32*i - Hi0bits(x[i-1])
	}
	return b
}
func Ratio(a *Bigint, b *Bigint) float64 {
	var da U
	var db U
	var k int
	var ka int
	var kb int
	&da.SetD(B2d(a, &ka))
	&db.SetD(B2d(b, &kb))
	k = ka - kb + 32*(a.GetWds()-b.GetWds())
	if k > 0 {
		&da.GetL()[1] += k * 0x100000
	} else {
		k = -k
		&db.GetL()[1] += k * 0x100000
	}
	return &da.GetD() / &db.GetD()
}

var Tens []float64 = []float64{1.0, 10.0, 100.0, 1000.0, 10000.0, 100000.0, 1000000.0, 1.0e7, 1.0e8, 1.0e9, 1.0e10, 9.9999998e10, 1.0e12, 9.9999998e12, 1.0e14, 9.9999999e14, 1.00000003e16, 9.9999998e16, 9.9999998e17, 1.0e19, 1.0e20, 1.0e21, 1.0e22}
var Bigtens []float64 = []float64{1.00000003e16, 1.0e32, Infinity, Infinity, Infinity}
var Tinytens []float64 = []float64{1.0e-16, 1.0e-32, 0.0, 0.0, 9.0071993e15 * 0.0}

/* The factor of 2^53 in tinytens[4] helps us avoid setting the underflow */

// #define Scale_Bit       0x10

// #define n_bigtens       5

// #define ULbits       32

// #define kshift       5

// #define kmask       31

func Dshift(b *Bigint, p2 int) int {
	var rv int = Hi0bits(b.GetX()[b.GetWds()-1]) - 4
	if p2 > 0 {
		rv -= p2
	}
	return rv & 31
}
func Quorem(b *Bigint, S *Bigint) int {
	var n int
	var bx *uint32
	var bxe *uint32
	var q uint32
	var sx *uint32
	var sxe *uint32
	var borrow unsigned__long__long
	var carry unsigned__long__long
	var y unsigned__long__long
	var ys unsigned__long__long
	n = S.GetWds()
	if b.GetWds() < n {
		return 0
	}
	sx = S.GetX()
	sxe = sx + g.PreDec(&n)
	bx = b.GetX()
	bxe = bx + n
	q = (*bxe) / ((*sxe) + 1)
	if q != 0 {
		borrow = 0
		carry = 0
		for {
			ys = g.PostInc(&(*sx)) * unsigned__long__long(q+carry)
			carry = ys >> 32
			y = (*bx) - (ys & 0xffffffff) - borrow
			borrow = y >> 32 & uint32(1)
			g.PostInc(&(*bx)) = y & 0xffffffff
			if sx > sxe {
				break
			}
		}
		if (*bxe) == 0 {
			bx = b.GetX()
			for g.PreDec(&bxe) > bx && (*bxe) == 0 {
				n--
			}
			b.SetWds(n)
		}
	}
	if Cmp(b, S) >= 0 {
		q++
		borrow = 0
		carry = 0
		bx = b.GetX()
		sx = S.GetX()
		for {
			ys = g.PostInc(&(*sx)) + carry
			carry = ys >> 32
			y = (*bx) - (ys & 0xffffffff) - borrow
			borrow = y >> 32 & uint32(1)
			g.PostInc(&(*bx)) = y & 0xffffffff
			if sx > sxe {
				break
			}
		}
		bx = b.GetX()
		bxe = bx + n
		if (*bxe) == 0 {
			for g.PreDec(&bxe) > bx && (*bxe) == 0 {
				n--
			}
			b.SetWds(n)
		}
	}
	return q
}
func Sulp(x *U, bc *BCinfo) float64 {
	var u U
	var rv float64
	var i int
	rv = Ulp(x)
	if bc.GetScale() == 0 || g.Assign(&i, 2*53+1-((x.GetL()[1]&0x7ff00000)>>20)) <= 0 {
		return rv
	}
	&u.GetL()[1] = 0x3ff00000 + (i << 20)
	&u.GetL()[0] = 0
	return rv * u.GetD()
}
func Bigcomp(rv *U, s0 *byte, bc *BCinfo) {
	var b *Bigint
	var d *Bigint
	var b2 int
	var bbits int
	var d2 int
	var dd int
	var dig int
	var dsign int
	var i int
	var j int
	var nd int
	var nd0 int
	var p2 int
	var p5 int
	var speccase int
	dsign = bc.GetDsign()
	nd = bc.GetNd()
	nd0 = bc.GetNd0()
	p5 = nd + bc.GetE0() - 1
	speccase = 0
	if rv.GetD() == 0.0 {

		/* threshold was rounded to zero */

		b = I2b(1)
		p2 = -1022 - 53 + 1
		bbits = 1
		rv.GetL()[1] = 53 + 2<<20
		i = 0
		speccase = 1
		p2--
		dsign = 0
		goto have_i
	} else {
		b = D2b(rv, &p2, &bbits)
	}
	p2 -= bc.GetScale()

	/* floor(log2(rv)) == bbits - 1 + p2 */

	i = 53 - bbits
	if i > g.Assign(&j, 53 - -1022 - 1 + p2) {
		i = j
	}
	b = Lshift(b, g.PreInc(&i))
	b.GetX()[0] |= 1
have_i:
	p2 -= p5 + i
	d = I2b(1)

	/* Arrange for convenient computation of quotients:
	 * shift left if necessary so divisor has 4 leading 0 bits.
	 */

	if p5 > 0 {
		d = Pow5mult(d, p5)
	} else if p5 < 0 {
		b = Pow5mult(b, -p5)
	}
	if p2 > 0 {
		b2 = p2
		d2 = 0
	} else {
		b2 = 0
		d2 = -p2
	}
	i = Dshift(d, d2)
	if g.AssignOp(&b2, "+=", i) > 0 {
		b = Lshift(b, b2)
	}
	if g.AssignOp(&d2, "+=", i) > 0 {
		d = Lshift(d, d2)
	}

	/* Now b/d = exactly half-way between the two floating-point values */

	if !(g.Assign(&dig, Quorem(b, d))) {
		b = Multadd(b, 10, 0)
		dig = Quorem(b, d)
	}

	/* Compare b/d with s0 */

	for i = 0; i < nd0; {
		if g.Assign(&dd, s0[g.PostInc(&i)]-'0'-dig) {
			goto ret
		}
		if b.GetX()[0] == 0 && b.GetWds() == 1 {
			if i < nd {
				dd = 1
			}
			goto ret
		}
		b = Multadd(b, 10, 0)
		dig = Quorem(b, d)
	}
	for j = bc.GetDp1(); g.PostInc(&i) < nd; {
		if g.Assign(&dd, s0[g.PostInc(&j)]-'0'-dig) {
			goto ret
		}
		if b.GetX()[0] == 0 && b.GetWds() == 1 {
			if i < nd {
				dd = 1
			}
			goto ret
		}
		b = Multadd(b, 10, 0)
		dig = Quorem(b, d)
	}
	if dig > 0 || b.GetX()[0] != 0 || b.GetWds() > 1 {
		dd = -1
	}
ret:
	Bfree(b)
	Bfree(d)
	if speccase != 0 {
		if dd <= 0 {
			rv.SetD(0.0)
		}
	} else if dd < 0 {
		if dsign == 0 {
		retlow1:
		}
		rv.SetD(rv.GetD() - Sulp(rv, bc))
	} else if dd > 0 {
		if dsign != 0 {
		rethi1:
			rv.SetD(rv.GetD() + Sulp(rv, bc))
		}
	} else {

		/* Exact half-way case:  apply round-even rule. */

		if g.Assign(&j, ((rv.GetL()[1]&0x7ff00000)>>20)-bc.GetScale()) <= 0 {
			i = 1 - j
			if i <= 31 {
				if (rv.GetL()[0] & 0x1 << i) != 0 {
					goto odd
				}
			} else if (rv.GetL()[1]&0x1<<i - 32) != 0 {
				goto odd
			}
		} else if (rv.GetL()[0] & 1) != 0 {
		odd:
			if dsign != 0 {
				goto rethi1
			}
			goto retlow1
		}

		/* Exact half-way case:  apply round-even rule. */

	}
	return
}
func ZendStrtod(s00 *byte, se **byte) float64 {
	var bb2 int
	var bb5 int
	var bbe int
	var bd2 int
	var bd5 int
	var bbbits int
	var bs2 int
	var c int
	var e int
	var e1 int
	var esign int
	var i int
	var j int
	var k int
	var nd int
	var nd0 int
	var nf int
	var nz int
	var nz0 int
	var nz1 int
	var sign int
	var s *byte
	var s0 *byte
	var s1 *byte
	var aadj float64
	var aadj1 float64
	var L int32
	var aadj2 U
	var adj U
	var rv U
	var rv0 U
	var y uint32
	var z uint32
	var bc BCinfo
	var bb *Bigint
	var bb1 *Bigint
	var bd *Bigint
	var bd0 *Bigint
	var bs *Bigint
	var delta *Bigint
	var Lsb uint32
	var Lsb1 uint32
	var req_bigcomp int = 0
	bc.SetUflchk(0)
	bc.SetDplen(bc.GetUflchk())
	nz = bc.GetDplen()
	nz1 = nz
	nz0 = nz1
	sign = nz0
	&rv.SetD(0.0)
	for s = s00; ; s++ {
		switch *s {
		case '-':
			sign = 1
		case '+':
			if *(g.PreInc(&s)) {
				goto break2
			}
		case 0:
			goto ret0
		case '\t':

		case '\n':

		case 'v':

		case 'f':

		case '\r':

		case ' ':
			continue
		default:
			goto break2
		}
	}
break2:
	if (*s) == '0' {
		nz0 = 1
		for (*(g.PreInc(&s))) == '0' {

		}
		if !(*s) {
			goto ret
		}
	}
	s0 = s
	z = 0
	y = z
	nf = 0
	nd = nf
	for g.Assign(&c, *s) >= '0' && c <= '9' {
		if nd < 9 {
			y = 10*y + c - '0'
		} else if nd < DBL_DIG+2 {
			z = 10*z + c - '0'
		}
		nd++
		s++
	}
	nd0 = nd
	bc.SetDp1(s - s0)
	bc.SetDp0(bc.GetDp1())
	for s1 = s; s1 > s0 && (*(g.PreDec(&s1))) == '0'; {
		nz1++
	}
	if c == '.' {
		c = *(g.PreInc(&s))
		bc.SetDp1(s - s0)
		bc.SetDplen(bc.GetDp1() - bc.GetDp0())
		if nd == 0 {
			for ; c == '0'; c = *(g.PreInc(&s)) {
				nz++
			}
			if c > '0' && c <= '9' {
				bc.SetDp0(s0 - s)
				bc.SetDp1(bc.GetDp0() + bc.GetDplen())
				s0 = s
				nf += nz
				nz = 0
				goto have_dig
			}
			goto dig_done
		}
		for ; c >= '0' && c <= '9'; c = *(g.PreInc(&s)) {
		have_dig:
			nz++
			if g.AssignOp(&c, "-=", '0') {
				nf += nz
				for i = 1; i < nz; i++ {
					if g.PostInc(&nd) < 9 {
						y *= 10
					} else if nd <= DBL_DIG+2 {
						z *= 10
					}
				}
				if g.PostInc(&nd) < 9 {
					y = 10*y + c
				} else if nd <= DBL_DIG+2 {
					z = 10*z + c
				}
				nz1 = 0
				nz = nz1
			}
		}
	}
dig_done:
	if nd < 0 {

		/* overflow */

		nd = DBL_DIG + 2

		/* overflow */

	}
	if nf < 0 {

		/* overflow */

		nf = DBL_DIG + 2

		/* overflow */

	}
	e = 0
	if c == 'e' || c == 'E' {
		if nd == 0 && nz == 0 && nz0 == 0 {
			goto ret0
		}
		s00 = s
		esign = 0
		switch g.Assign(&c, *(g.PreInc(&s))) {
		case '-':
			esign = 1
		case '+':
			c = *(g.PreInc(&s))
		}
		if c >= '0' && c <= '9' {
			for c == '0' {
				c = *(g.PreInc(&s))
			}
			if c > '0' && c <= '9' {
				L = c - '0'
				s1 = s
				for g.Assign(&c, *(g.PreInc(&s))) >= '0' && c <= '9' {
					L = int32(10 * uint32(L+(c-'0')))
				}
				if s-s1 > 8 || L > 19999 {

					/* Avoid confusion from exponents
					 * so large that e might overflow.
					 */

					e = 19999
				} else {
					e = int(L)
				}
				if esign != 0 {
					e = -e
				}
			} else {
				e = 0
			}
		} else {
			s = s00
		}
	}
	if nd == 0 {
		if nz == 0 && nz0 == 0 {
		ret0:
			s = s00
			sign = 0
		}
		goto ret
	}
	e -= nf
	e1 = e
	bc.SetE0(e1)

	/* Now we have nd0 digits, starting at s0, followed by a
	 * decimal point, followed by nd-nd0 digits.  The number we're
	 * after is the integer represented by those digits times
	 * 10**e */

	if nd0 == 0 {
		nd0 = nd
	}
	if nd < DBL_DIG+2 {
		k = nd
	} else {
		k = DBL_DIG + 2
	}
	&rv.SetD(y)
	if k > 9 {
		&rv.SetD(Tens[k-9]*&rv.GetD() + z)
	}
	bd0 = 0
	if nd <= DBL_DIG {
		if e == 0 {
			goto ret
		}
		if e > 0 {
			if e <= 22 {

				/* rv = */

				&rv.SetD(&rv.GetD() * Tens[e])
				goto ret
			}
			i = DBL_DIG - nd
			if e <= 22+i {

				/* A fancier test would sometimes let us do
				 * this for larger i values.
				 */

				e -= i
				&rv.SetD(&rv.GetD() * Tens[i])

				/* rv = */

				&rv.SetD(&rv.GetD() * Tens[e])
				goto ret
			}
		} else if e >= -22 {

			/* rv = */

			&rv.SetD(&rv.GetD() / Tens[-e])
			goto ret
		}
	}
	e1 += nd - k
	bc.SetScale(0)

	/* Get starting approximation = rv * 10**e1 */

	if e1 > 0 {
		if g.Assign(&i, e1&15) {
			&rv.SetD(&rv.GetD() * Tens[i])
		}
		if g.AssignOp(&e1, "&=", ^15) {
			if e1 > DBL_MAX_10_EXP {
			ovfl:

				/* Can't trust HUGE_VAL */

				&rv.GetL()[1] = 0x7ff00000
				&rv.GetL()[0] = 0
			range_err:
				if bd0 != nil {
					Bfree(bb)
					Bfree(bd)
					Bfree(bs)
					Bfree(bd0)
					Bfree(delta)
				}
				goto ret
			}
			e1 >>= 4
			for j = 0; e1 > 1; {
				if (e1 & 1) != 0 {
					&rv.SetD(&rv.GetD() * Bigtens[j])
				}
				j++
				e1 >>= 1
			}

			/* The last multiplication could overflow. */

			&rv.GetL()[1] -= 53 * 0x100000
			&rv.SetD(&rv.GetD() * Bigtens[j])
			if g.Assign(&z, &rv.GetL()[1]&0x7ff00000) > 0x100000*(DBL_MAX_EXP+1023-53) {
				goto ovfl
			}
			if z > 0x100000*(DBL_MAX_EXP+1023-1-53) {

				/* set to largest number */

				&rv.GetL()[1] = 0xfffff | 0x100000*(DBL_MAX_EXP+1023-1)
				&rv.GetL()[0] = 0xffffffff
			} else {
				&rv.GetL()[1] += 53 * 0x100000
			}
		}
	} else if e1 < 0 {
		e1 = -e1
		if g.Assign(&i, e1&15) {
			&rv.SetD(&rv.GetD() / Tens[i])
		}
		if g.AssignOp(&e1, ">>=", 4) {
			if e1 >= 1<<5 {
				goto undfl
			}
			if (e1 & 0x10) != 0 {
				bc.SetScale(2 * 53)
			}
			for j = 0; e1 > 0; {
				if (e1 & 1) != 0 {
					&rv.SetD(&rv.GetD() * Tinytens[j])
				}
				j++
				e1 >>= 1
			}
			if bc.GetScale() != 0 && g.Assign(&j, 2*53+1-((&rv.GetL()[1]&0x7ff00000)>>20)) > 0 {

				/* scaled rv is denormal; clear j low bits */

				if j >= 32 {
					if j > 54 {
						goto undfl
					}
					&rv.GetL()[0] = 0
					if j >= 53 {
						&rv.GetL()[1] = (53 + 2) * 0x100000
					} else {
						&rv.GetL()[1] &= 0xffffffff<<j - 32
					}
				} else {
					&rv.GetL()[0] &= 0xffffffff << j
				}

				/* scaled rv is denormal; clear j low bits */

			}
			if !(&rv.GetD()) {
			undfl:
				&rv.SetD(0.0)
				goto range_err
			}
		}
	}

	/* Now the hard part -- adjusting rv to the correct value.*/

	bc.SetNd(nd - nz1)
	bc.SetNd0(nd0)

	/* to silence an erroneous warning about bc.nd0 */

	if nd > 40 {

		/* ASSERT(strtod_diglim >= 18); 18 == one more than the */

		j = 18
		i = j
		if i > nd0 {
			j += bc.GetDplen()
		}
		for {
			if g.PreDec(&j) < bc.GetDp1() && j >= bc.GetDp0() {
				j = bc.GetDp0() - 1
			}
			if s0[j] != '0' {
				break
			}
			i--
		}
		e += nd - i
		nd = i
		if nd0 > nd {
			nd0 = nd
		}
		if nd < 9 {
			y = 0
			for i = 0; i < nd0; i++ {
				y = 10*y + s0[i] - '0'
			}
			for j = bc.GetDp1(); i < nd; i++ {
				y = 10*y + s0[g.PostInc(&j)] - '0'
			}
		}
	}
	bd0 = S2b(s0, nd0, nd, y, bc.GetDplen())
	for {
		bd = Balloc(bd0.GetK())
		memcpy((*byte)(&bd.sign), (*byte)(&bd0.sign), bd0.GetWds()*g.SizeOf("Long")+2*g.SizeOf("int"))
		bb = D2b(&rv, &bbe, &bbbits)
		bs = I2b(1)
		if e >= 0 {
			bb5 = 0
			bb2 = bb5
			bd5 = e
			bd2 = bd5
		} else {
			bb5 = -e
			bb2 = bb5
			bd5 = 0
			bd2 = bd5
		}
		if bbe >= 0 {
			bb2 += bbe
		} else {
			bd2 -= bbe
		}
		bs2 = bb2
		Lsb = 1
		Lsb1 = 0
		j = bbe - bc.GetScale()
		i = j + bbbits - 1
		j = 53 + 1 - bbbits
		if i < -1022 {
			i = -1022 - i
			j -= i
			if i < 32 {
				Lsb <<= i
			} else if i < 52 {
				Lsb1 = Lsb<<i - 32
			} else {
				Lsb1 = 0x7ff00000
			}
		}
		bb2 += j
		bd2 += j
		bd2 += bc.GetScale()
		if bb2 < bd2 {
			i = bb2
		} else {
			i = bd2
		}
		if i > bs2 {
			i = bs2
		}
		if i > 0 {
			bb2 -= i
			bd2 -= i
			bs2 -= i
		}
		if bb5 > 0 {
			bs = Pow5mult(bs, bb5)
			bb1 = Mult(bs, bb)
			Bfree(bb)
			bb = bb1
		}
		if bb2 > 0 {
			bb = Lshift(bb, bb2)
		}
		if bd5 > 0 {
			bd = Pow5mult(bd, bd5)
		}
		if bd2 > 0 {
			bd = Lshift(bd, bd2)
		}
		if bs2 > 0 {
			bs = Lshift(bs, bs2)
		}
		delta = Diff(bb, bd)
		bc.SetDsign(delta.GetSign())
		delta.SetSign(0)
		i = Cmp(delta, bs)
		if bc.GetNd() > nd && i <= 0 {
			if bc.GetDsign() != 0 {

				/* Must use bigcomp(). */

				req_bigcomp = 1
				break
			}
			i = -1
		}
		if i < 0 {

			/* Error is less than half an ulp -- check for
			 * special case of mantissa a power of two.
			 */

			if bc.GetDsign() != 0 || &rv.GetL()[0] != 0 || (&rv.GetL()[1]&0xfffff) != 0 || (&rv.GetL()[1]&0x7ff00000) <= (2*53+1)*0x100000 {
				break
			}
			if delta.GetX()[0] == 0 && delta.GetWds() <= 1 {

				/* exact result */

				break

				/* exact result */

			}
			delta = Lshift(delta, 1)
			if Cmp(delta, bs) > 0 {
				goto drop_down
			}
			break
		}
		if i == 0 {

			/* exactly half-way between */

			if bc.GetDsign() != 0 {
				if (&rv.GetL()[1]&0xfffff) == 0xfffff && &rv.GetL()[0] == g.Cond(bc.GetScale() != 0 && g.Assign(&y, &rv.GetL()[1]&0x7ff00000) <= 2*53*0x100000, 0xffffffff&0xffffffff<<2*53+1-(y>>20), 0xffffffff) {

					/*boundary case -- increment exponent*/

					if &rv.GetL()[1] == (0xfffff|0x100000*(DBL_MAX_EXP+1023-1)) && &rv.GetL()[0] == 0xffffffff {
						goto ovfl
					}
					&rv.GetL()[1] = (&rv.GetL()[1] & 0x7ff00000) + 0x100000
					&rv.GetL()[0] = 0
					bc.SetDsign(0)
					break
				}
			} else if (&rv.GetL()[1]&0xfffff) == 0 && &rv.GetL()[0] == 0 {
			drop_down:

				/* boundary case -- decrement exponent */

				if bc.GetScale() != 0 {
					L = &rv.GetL()[1] & 0x7ff00000
					if L <= (2*53+1)*0x100000 {
						if L > (53+2)*0x100000 {

							/* round even ==> */

							break
						}

						/* rv = smallest denormal */

						if bc.GetNd() > nd {
							bc.SetUflchk(1)
							break
						}
						goto undfl
					}
				}
				L = (&rv.GetL()[1] & 0x7ff00000) - 0x100000
				&rv.GetL()[1] = L | 0xfffff
				&rv.GetL()[0] = 0xffffffff
				if bc.GetNd() > nd {
					goto cont
				}
				break
			}
			if Lsb1 != 0 {
				if (&rv.GetL()[1] & Lsb1) == 0 {
					break
				}
			} else if (&rv.GetL()[0] & Lsb) == 0 {
				break
			}
			if bc.GetDsign() != 0 {
				&rv.SetD(&rv.GetD() + Sulp(&rv, &bc))
			} else {
				&rv.SetD(&rv.GetD() - Sulp(&rv, &bc))
				if !(&rv.GetD()) {
					if bc.GetNd() > nd {
						bc.SetUflchk(1)
						break
					}
					goto undfl
				}
			}
			bc.SetDsign(1 - bc.GetDsign())
			break
		}
		if g.Assign(&aadj, Ratio(delta, bs)) <= 2.0 {
			if bc.GetDsign() != 0 {
				aadj1 = 1.0
				aadj = aadj1
			} else if &rv.GetL()[0] != 0 || (&rv.GetL()[1]&0xfffff) != 0 {
				if &rv.GetL()[0] == 1 && &rv.GetL()[1] == 0 {
					if bc.GetNd() > nd {
						bc.SetUflchk(1)
						break
					}
					goto undfl
				}
				aadj = 1.0
				aadj1 = -1.0
			} else {

				/* special case -- power of FLT_RADIX to be */

				if aadj < 2.0/FLT_RADIX {
					aadj = 1.0 / FLT_RADIX
				} else {
					aadj *= 0.5
				}
				aadj1 = -aadj
			}
		} else {
			aadj *= 0.5
			if bc.GetDsign() != 0 {
				aadj1 = aadj
			} else {
				aadj1 = -aadj
			}

		}
		y = &rv.GetL()[1] & 0x7ff00000

		/* Check for overflow */

		if y == 0x100000*(DBL_MAX_EXP+1023-1) {
			&rv0.SetD(&rv.GetD())
			&rv.GetL()[1] -= 53 * 0x100000
			adj.SetD(aadj1 * Ulp(&rv))
			&rv.SetD(&rv.GetD() + adj.GetD())
			if (&rv.GetL()[1] & 0x7ff00000) >= 0x100000*(DBL_MAX_EXP+1023-53) {
				if &rv0.GetL()[1] == (0xfffff|0x100000*(DBL_MAX_EXP+1023-1)) && &rv0.GetL()[0] == 0xffffffff {
					goto ovfl
				}
				&rv.GetL()[1] = 0xfffff | 0x100000*(DBL_MAX_EXP+1023-1)
				&rv.GetL()[0] = 0xffffffff
				goto cont
			} else {
				&rv.GetL()[1] += 53 * 0x100000
			}
		} else {
			if bc.GetScale() != 0 && y <= 2*53*0x100000 {
				if aadj <= 0x7fffffff {
					if g.Assign(&z, aadj) <= 0 {
						z = 1
					}
					aadj = z
					if bc.GetDsign() != 0 {
						aadj1 = aadj
					} else {
						aadj1 = -aadj
					}
				}
				&aadj2.SetD(aadj1)
				&aadj2.GetL()[1] += (2*53+1)*0x100000 - y
				aadj1 = &aadj2.GetD()
				adj.SetD(aadj1 * Ulp(&rv))
				&rv.SetD(&rv.GetD() + adj.GetD())
				if rv.GetD() == 0.0 {
					req_bigcomp = 1
					break
				}
			} else {
				adj.SetD(aadj1 * Ulp(&rv))
				&rv.SetD(&rv.GetD() + adj.GetD())
			}
		}
		z = &rv.GetL()[1] & 0x7ff00000
		if bc.GetNd() == nd {
			if bc.GetScale() == 0 {
				if y == z {

					/* Can we stop now? */

					L = int32(aadj)
					aadj -= L

					/* The tolerances below are conservative. */

					if bc.GetDsign() != 0 || &rv.GetL()[0] != 0 || (&rv.GetL()[1]&0xfffff) != 0 {
						if aadj < 0.4999999 || aadj > 0.5000001 {
							break
						}
					} else if aadj < 0.4999999/FLT_RADIX {
						break
					}

					/* The tolerances below are conservative. */

				}
			}
		}
	cont:
		Bfree(bb)
		Bfree(bd)
		Bfree(bs)
		Bfree(delta)
	}
	Bfree(bb)
	Bfree(bd)
	Bfree(bs)
	Bfree(bd0)
	Bfree(delta)
	if req_bigcomp != 0 {
		bd0 = 0
		bc.SetE0(bc.GetE0() + nz1)
		Bigcomp(&rv, s0, &bc)
		y = &rv.GetL()[1] & 0x7ff00000
		if y == 0x7ff00000 {
			goto ovfl
		}
		if y == 0 && rv.GetD() == 0.0 {
			goto undfl
		}
	}
	if bc.GetScale() != 0 {
		&rv0.GetL()[1] = 0x3ff00000 - 2*53*0x100000
		&rv0.GetL()[0] = 0
		&rv.SetD(&rv.GetD() * &rv0.GetD())
	}
ret:
	if se != nil {
		*se = (*byte)(s)
	}
	if sign != 0 {
		return -(&rv.GetD())
	} else {
		return &rv.GetD()
	}
}

var DtoaResult *byte

func RvAlloc(i int) *byte {
	var j int
	var k int
	var r *int
	j = g.SizeOf("ULong")
	for k = 0; g.SizeOf("Bigint")-g.SizeOf("ULong")-g.SizeOf("int")+int(j <= int(i)) != 0; j <<= 1 {
		k++
	}
	r = (*int)(Balloc(k))
	*r = k
	DtoaResult = (*byte)(r + 1)
	return DtoaResult
}
func NrvAlloc(s string, rve **byte, n int) *byte {
	var rv *byte
	var t *byte
	rv = RvAlloc(n)
	t = rv
	for g.Assign(&(*t), g.PostInc(&(*s))) {
		t++
	}
	if rve != nil {
		*rve = t
	}
	return rv
}

/* freedtoa(s) must be used to free values s returned by dtoa
 * when MULTIPLE_THREADS is #defined.  It should be used in all cases,
 * but for consistency with earlier versions of dtoa, it is optional
 * when MULTIPLE_THREADS is not defined.
 */

func ZendFreedtoa(s *byte) {
	var b *Bigint = (*Bigint)((*int)(s - 1))
	b.SetMaxwds(1 << g.Assign(&(b.GetK()), *((*int)(b))))
	Bfree(b)
	if s == DtoaResult {
		DtoaResult = 0
	}
}

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

func ZendDtoa(dd float64, mode int, ndigits int, decpt *int, sign *int, rve **byte) *byte {
	/*    Arguments ndigits, decpt, sign are similar to those
	      of ecvt and fcvt; trailing zeros are suppressed from
	      the returned string.  If not null, *rve is set to point
	      to the end of the return value.  If d is +-Infinity or NaN,
	      then *decpt is set to 9999.

	      mode:
	          0 ==> shortest string that yields d when read in
	              and rounded to nearest.
	          1 ==> like 0, but with Steele & White stopping rule;
	              e.g. with IEEE P754 arithmetic , mode 0 gives
	              1e23 whereas mode 1 gives 9.999999999999999e22.
	          2 ==> max(1,ndigits) significant digits.  This gives a
	              return value similar to that of ecvt, except
	              that trailing zeros are suppressed.
	          3 ==> through ndigits past the decimal point.  This
	              gives a return value similar to that from fcvt,
	              except that trailing zeros are suppressed, and
	              ndigits can be negative.
	          4,5 ==> similar to 2 and 3, respectively, but (in
	              round-nearest mode) with the tests of mode 0 to
	              possibly return a shorter string that rounds to d.
	              With IEEE arithmetic and compilation with
	              -DHonor_FLT_ROUNDS, modes 4 and 5 behave the same
	              as modes 2 and 3 when FLT_ROUNDS != 1.
	          6-9 ==> Debugging modes similar to mode - 4:  don't try
	              fast floating-point estimate (if applicable).

	          Values of mode other than 0-9 are treated as mode 0.

	          Sufficient space is allocated to the return value
	          to hold the suppressed trailing zeros.
	*/

	var bbits int
	var b2 int
	var b5 int
	var be int
	var dig int
	var i int
	var ieps int
	var ilim int = 0
	var ilim0 int
	var ilim1 int
	var j int
	var j1 int = 0
	var k int
	var k0 int
	var k_check int
	var leftright int
	var m2 int
	var m5 int
	var s2 int
	var s5 int
	var spec_case int = 0
	var try_quick int
	var L int32
	var denorm int
	var x uint32
	var b *Bigint
	var b1 *Bigint
	var delta *Bigint
	var mlo *Bigint
	var mhi *Bigint
	var S *Bigint
	var d2 U
	var eps U
	var u U
	var ds float64
	var s *byte
	var s0 *byte
	var eps1 U
	if DtoaResult != nil {
		ZendFreedtoa(DtoaResult)
		DtoaResult = 0
	}
	u.SetD(dd)
	if (&u.GetL()[1] & 0x80000000) != 0 {

		/* set sign for everything, including 0's and NaNs */

		*sign = 1
		&u.GetL()[1] &= ^0x80000000
	} else {
		*sign = 0
	}
	if (&u.GetL()[1] & 0x7ff00000) == 0x7ff00000 {

		/* Infinity or NaN */

		*decpt = 9999
		if &u.GetL()[0] == 0 && (&u.GetL()[1]&0xfffff) == 0 {
			return NrvAlloc("Infinity", rve, 8)
		}
		return NrvAlloc("NaN", rve, 3)
	}
	if !(&u.GetD()) {
		*decpt = 1
		return NrvAlloc("0", rve, 1)
	}
	b = D2b(&u, &be, &bbits)
	if g.Assign(&i, int(&u.GetL()[1]>>20&0x7ff00000>>20)) {
		&d2.SetD(&u.GetD())
		&d2.GetL()[1] &= 0xfffff
		&d2.GetL()[1] |= 0x3ff00000

		/* log(x)    ~=~ log(1.5) + (x-1.5)/1.5
		 * log10(x)     =  log(x) / log(10)
		 *        ~=~ log(1.5)/log(10) + (x-1.5)/(1.5*log(10))
		 * log10(d) = (i-Bias)*log(2)/log(10) + log10(d2)
		 *
		 * This suggests computing an approximation k to log10(d) by
		 *
		 * k = (i - Bias)*0.301029995663981
		 *    + ( (d2-1.5)*0.289529654602168 + 0.176091259055681 );
		 *
		 * We want k to be too large rather than too small.
		 * The error in the first-order Taylor series approximation
		 * is in our favor, so we just round up the constant enough
		 * to compensate for any error in the multiplication of
		 * (i - Bias) by 0.301029995663981; since |i - Bias| <= 1077,
		 * and 1077 * 0.30103 * 2^-52 ~=~ 7.2e-14,
		 * adding 1e-13 to the constant term more than suffices.
		 * Hence we adjust the constant term to 0.1760912590558.
		 * (We could get a more accurate k by invoking log10,
		 *  but this is probably not worthwhile.)
		 */

		i -= 1023
		denorm = 0
	} else {

		/* d is denormalized */

		i = bbits + be + (1023 + (53 - 1) - 1)
		if i > 32 {
			x = &u.GetL()[1]<<64 - i | &u.GetL()[0]>>i - 32
		} else {
			x = &u.GetL()[0]<<32 - i
		}
		&d2.SetD(x)
		&d2.GetL()[1] -= 31 * 0x100000
		i -= 1023 + (53 - 1) - 1 + 1
		denorm = 1
	}
	ds = (&d2.GetD()-1.5)*0.28952965 + 0.17609125 + i*0.30103
	k = int(ds)
	if ds < 0.0 && ds != k {
		k--
	}
	k_check = 1
	if k >= 0 && k <= 22 {
		if &u.GetD() < Tens[k] {
			k--
		}
		k_check = 0
	}
	j = bbits - i - 1
	if j >= 0 {
		b2 = 0
		s2 = j
	} else {
		b2 = -j
		s2 = 0
	}
	if k >= 0 {
		b5 = 0
		s5 = k
		s2 += k
	} else {
		b2 -= k
		b5 = -k
		s5 = 0
	}
	if mode < 0 || mode > 9 {
		mode = 0
	}
	try_quick = 1
	if mode > 5 {
		mode -= 4
		try_quick = 0
	}
	leftright = 1
	ilim1 = -1
	ilim = ilim1

	/* silence erroneous "gcc -Wall" warning. */

	switch mode {
	case 0:

	case 1:
		i = 18
		ndigits = 0
		break
	case 2:
		leftright = 0
	case 4:
		if ndigits <= 0 {
			ndigits = 1
		}
		i = ndigits
		ilim1 = i
		ilim = ilim1
		break
	case 3:
		leftright = 0
	case 5:
		i = ndigits + k + 1
		ilim = i
		ilim1 = i - 1
		if i <= 0 {
			i = 1
		}
	}
	s0 = RvAlloc(i)
	s = s0
	if ilim >= 0 && ilim <= 14 && try_quick != 0 {

		/* Try to get by with floating-point arithmetic. */

		i = 0
		&d2.SetD(&u.GetD())
		k0 = k
		ilim0 = ilim
		ieps = 2
		if k > 0 {
			ds = Tens[k&0xf]
			j = k >> 4
			if (j & 0x10) != 0 {

				/* prevent overflows */

				j &= 0x10 - 1
				&u.SetD(&u.GetD() / Bigtens[5-1])
				ieps++
			}
			for j != 0 {
				if (j & 1) != 0 {
					ieps++
					ds *= Bigtens[i]
				}
				j >>= 1
				i++
			}
			&u.SetD(&u.GetD() / ds)
		} else if g.Assign(&j1, -k) {
			&u.SetD(&u.GetD() * Tens[j1&0xf])
			for j = j1 >> 4; j != 0; {
				if (j & 1) != 0 {
					ieps++
					&u.SetD(&u.GetD() * Bigtens[i])
				}
				j >>= 1
				i++
			}
		}
		if k_check != 0 && &u.GetD() < 1.0 && ilim > 0 {
			if ilim1 <= 0 {
				goto fast_failed
			}
			ilim = ilim1
			k--
			&u.SetD(&u.GetD() * 10.0)
			ieps++
		}
		&eps.SetD(ieps*&u.GetD() + 7.0)
		&eps.GetL()[1] -= (53 - 1) * 0x100000
		if ilim == 0 {
			mhi = 0
			S = mhi
			&u.SetD(&u.GetD() - 5.0)
			if &u.GetD() > &eps.GetD() {
				goto one_digit
			}
			if &u.GetD() < -(&eps.GetD()) {
				goto no_digits
			}
			goto fast_failed
		}
		if leftright != 0 {

			/* Use Steele & White method of only
			 * generating digits needed.
			 */

			&eps.SetD(0.5/Tens[ilim-1] - &eps.GetD())
			if k0 < 0 && j1 >= 307 {
				eps1.SetD(Infinity)
				&eps1.GetL()[1] -= 0x100000 * (1023 + 53 - 1)
				&eps1.SetD(&eps1.GetD() * Tens[j1&0xf])
				i = 0
				j = j1 - 256>>4
				for j != 0 {
					if (j & 1) != 0 {
						&eps1.SetD(&eps1.GetD() * Bigtens[i])
					}
					j >>= 1
					i++
				}
				if eps.GetD() < eps1.GetD() {
					eps.SetD(eps1.GetD())
				}
			}
			for i = 0; ; {
				L = &u.GetD()
				&u.SetD(&u.GetD() - L)
				g.PostInc(&(*s)) = '0' + int(L)
				if 1.0-&u.GetD() < &eps.GetD() {
					goto bump_up
				}
				if &u.GetD() < &eps.GetD() {
					goto ret1
				}
				if g.PreInc(&i) >= ilim {
					break
				}
				&eps.SetD(&eps.GetD() * 10.0)
				&u.SetD(&u.GetD() * 10.0)
			}
		} else {

			/* Generate ilim digits, then fix them up. */

			&eps.SetD(&eps.GetD() * Tens[ilim-1])
			for i = 1; ; {
				L = int32(&u.GetD())
				if !(g.AssignOp(&(&u.GetD()), "-=", L)) {
					ilim = i
				}
				g.PostInc(&(*s)) = '0' + int(L)
				if i == ilim {
					if &u.GetD() > 0.5+&eps.GetD() {
						goto bump_up
					} else if &u.GetD() < 0.5-&eps.GetD() {
						for (*(g.PreDec(&s))) == '0' {

						}
						s++
						goto ret1
					}
					break
				}
				i++
				&u.SetD(&u.GetD() * 10.0)
			}
		}
	fast_failed:
		s = s0
		&u.SetD(&d2.GetD())
		k = k0
		ilim = ilim0
	}

	/* Do we have a "small" integer? */

	if be >= 0 && k <= 14 {

		/* Yes. */

		ds = Tens[k]
		if ndigits < 0 && ilim <= 0 {
			mhi = 0
			S = mhi
			if ilim < 0 || &u.GetD() <= 5*ds {
				goto no_digits
			}
			goto one_digit
		}
		for i = 1; ; {
			L = int32(&u.GetD() / ds)
			&u.SetD(&u.GetD() - L*ds)
			g.PostInc(&(*s)) = '0' + int(L)
			if !(&u.GetD()) {
				break
			}
			if i == ilim {
				&u.SetD(&u.GetD() + &u.GetD())
				if &u.GetD() > ds || &u.GetD() == ds && (L&1) != 0 {
				bump_up:
					for (*(g.PreDec(&s))) == '9' {
						if s == s0 {
							k++
							*s = '0'
							break
						}
					}
					g.PreInc(&(*s))++
				}
				break
			}
			i++
			&u.SetD(&u.GetD() * 10.0)
		}
		goto ret1
	}
	m2 = b2
	m5 = b5
	mlo = 0
	mhi = mlo
	if leftright != 0 {
		if denorm != 0 {
			i = be + (1023 + (53 - 1) - 1 + 1)
		} else {
			i = 1 + 53 - bbits
		}
		b2 += i
		s2 += i
		mhi = I2b(1)
	}
	if m2 > 0 && s2 > 0 {
		if m2 < s2 {
			i = m2
		} else {
			i = s2
		}
		b2 -= i
		m2 -= i
		s2 -= i
	}
	if b5 > 0 {
		if leftright != 0 {
			if m5 > 0 {
				mhi = Pow5mult(mhi, m5)
				b1 = Mult(mhi, b)
				Bfree(b)
				b = b1
			}
			if g.Assign(&j, b5-m5) {
				b = Pow5mult(b, j)
			}
		} else {
			b = Pow5mult(b, b5)
		}
	}
	S = I2b(1)
	if s5 > 0 {
		S = Pow5mult(S, s5)
	}

	/* Check for special case that d is a normalized power of 2. */

	spec_case = 0
	if mode < 2 || leftright != 0 {
		if &u.GetL()[0] == 0 && (&u.GetL()[1]&0xfffff) == 0 && (&u.GetL()[1]&(0x7ff00000 & ^0x100000)) != 0 {

			/* The special case */

			b2 += 1
			s2 += 1
			spec_case = 1
		}
	}

	/* Arrange for convenient computation of quotients:
	 * shift left if necessary so divisor has 4 leading 0 bits.
	 *
	 * Perhaps we should just compute leading 28 bits of S once
	 * and for all and pass them and a shift to quorem, so it
	 * can do shifts and ors to compute the numerator for q.
	 */

	i = Dshift(S, s2)
	b2 += i
	m2 += i
	s2 += i
	if b2 > 0 {
		b = Lshift(b, b2)
	}
	if s2 > 0 {
		S = Lshift(S, s2)
	}
	if k_check != 0 {
		if Cmp(b, S) < 0 {
			k--
			b = Multadd(b, 10, 0)
			if leftright != 0 {
				mhi = Multadd(mhi, 10, 0)
			}
			ilim = ilim1
		}
	}
	if ilim <= 0 && (mode == 3 || mode == 5) {
		if ilim < 0 || Cmp(b, g.Assign(&S, Multadd(S, 5, 0))) <= 0 {

			/* no digits, fcvt style */

		no_digits:
			k = -1 - ndigits
			goto ret
		}
	one_digit:
		g.PostInc(&(*s)) = '1'
		k++
		goto ret
	}
	if leftright != 0 {
		if m2 > 0 {
			mhi = Lshift(mhi, m2)
		}

		/* Compute mlo -- check for special case
		 * that d is a normalized power of 2.
		 */

		mlo = mhi
		if spec_case != 0 {
			mhi = Balloc(mhi.GetK())
			memcpy((*byte)(&mhi.sign), (*byte)(&mlo.sign), mlo.GetWds()*g.SizeOf("Long")+2*g.SizeOf("int"))
			mhi = Lshift(mhi, 1)
		}
		for i = 1; ; i++ {
			dig = Quorem(b, S) + '0'

			/* Do we yet have the shortest decimal string
			 * that will round to d?
			 */

			j = Cmp(b, mlo)
			delta = Diff(S, mhi)
			if delta.GetSign() != 0 {
				j1 = 1
			} else {
				j1 = Cmp(b, delta)
			}
			Bfree(delta)
			if j1 == 0 && mode != 1 && (&u.GetL()[0]&1) == 0 {
				if dig == '9' {
					goto round_9_up
				}
				if j > 0 {
					dig++
				}
				g.PostInc(&(*s)) = dig
				goto ret
			}
			if j < 0 || j == 0 && mode != 1 && (&u.GetL()[0]&1) == 0 {
				if b.GetX()[0] == 0 && b.GetWds() <= 1 {
					goto accept_dig
				}
				if j1 > 0 {
					b = Lshift(b, 1)
					j1 = Cmp(b, S)
					if (j1 > 0 || j1 == 0 && (dig&1) != 0) && g.PostInc(&dig) == '9' {
						goto round_9_up
					}
				}
			accept_dig:
				g.PostInc(&(*s)) = dig
				goto ret
			}
			if j1 > 0 {
				if dig == '9' {
				round_9_up:
					g.PostInc(&(*s)) = '9'
					goto roundoff
				}
				g.PostInc(&(*s)) = dig + 1
				goto ret
			}
			g.PostInc(&(*s)) = dig
			if i == ilim {
				break
			}
			b = Multadd(b, 10, 0)
			if mlo == mhi {
				mhi = Multadd(mhi, 10, 0)
				mlo = mhi
			} else {
				mlo = Multadd(mlo, 10, 0)
				mhi = Multadd(mhi, 10, 0)
			}
		}
	} else {
		for i = 1; ; i++ {
			dig = Quorem(b, S) + '0'
			g.PostInc(&(*s)) = dig
			if b.GetX()[0] == 0 && b.GetWds() <= 1 {
				goto ret
			}
			if i >= ilim {
				break
			}
			b = Multadd(b, 10, 0)
		}
	}

	/* Round off last digit */

	b = Lshift(b, 1)
	j = Cmp(b, S)
	if j > 0 || j == 0 && (dig&1) != 0 {
	roundoff:
		for (*(g.PreDec(&s))) == '9' {
			if s == s0 {
				k++
				g.PostInc(&(*s)) = '1'
				goto ret
			}
		}
		g.PreInc(&(*s))++
	} else {
		for (*(g.PreDec(&s))) == '0' {

		}
		s++
	}
ret:
	Bfree(S)
	if mhi != nil {
		if mlo != nil && mlo != mhi {
			Bfree(mlo)
		}
		Bfree(mhi)
	}
ret1:
	Bfree(b)
	*s = 0
	*decpt = k + 1
	if rve != nil {
		*rve = s
	}
	return s0
}
func ZendHexStrtod(str *byte, endptr **byte) float64 {
	var s *byte = str
	var c byte
	var any int = 0
	var value float64 = 0
	if (*s) == '0' && (s[1] == 'x' || s[1] == 'X') {
		s += 2
	}
	for g.Assign(&c, g.PostInc(&(*s))) {
		if c >= '0' && c <= '9' {
			c -= '0'
		} else if c >= 'A' && c <= 'F' {
			c -= 'A' - 10
		} else if c >= 'a' && c <= 'f' {
			c -= 'a' - 10
		} else {
			break
		}
		any = 1
		value = value*16 + c
	}
	if endptr != nil {
		if any != 0 {
			*endptr = s - 1
		} else {
			*endptr = str
		}
	}
	return value
}
func ZendOctStrtod(str *byte, endptr **byte) float64 {
	var s *byte = str
	var c byte
	var value float64 = 0
	var any int = 0
	if str[0] == '0' {
		if endptr != nil {
			*endptr = str
		}
		return 0.0
	}

	/* skip leading zero */

	s++
	for g.Assign(&c, g.PostInc(&(*s))) {
		if c < '0' || c > '7' {

			/* break and return the current value if the number is not well-formed
			 * that's what Linux strtol() does
			 */

			break

			/* break and return the current value if the number is not well-formed
			 * that's what Linux strtol() does
			 */

		}
		value = value*8 + c - '0'
		any = 1
	}
	if endptr != nil {
		if any != 0 {
			*endptr = s - 1
		} else {
			*endptr = str
		}
	}
	return value
}
func ZendBinStrtod(str *byte, endptr **byte) float64 {
	var s *byte = str
	var c byte
	var value float64 = 0
	var any int = 0
	if '0' == (*s) && ('b' == s[1] || 'B' == s[1]) {
		s += 2
	}
	for g.Assign(&c, g.PostInc(&(*s))) {

		/*
		 * Verify the validity of the current character as a base-2 digit.  In
		 * the event that an invalid digit is found, halt the conversion and
		 * return the portion which has been converted thus far.
		 */

		if '0' == c || '1' == c {
			value = value*2 + c - '0'
		} else {
			break
		}
		any = 1
	}

	/*
	 * As with many strtoX implementations, should the subject sequence be
	 * empty or not well-formed, no conversion is performed and the original
	 * value of str is stored in *endptr, provided that endptr is not a null
	 * pointer.
	 */

	if nil != endptr {
		*endptr = (*byte)(g.Cond(any != 0, s-1, str))
	}
	return value
}
func DestroyFreelist() {
	var i int
	var tmp *Bigint
	for i = 0; i <= 7; i++ {
		var listp **Bigint = &Freelist[i]
		for g.Assign(&tmp, *listp) != nil {
			*listp = tmp.GetNext()
			Free(tmp)
		}
		Freelist[i] = nil
	}
}
func FreeP5s() {
	var listp **Bigint
	var tmp **Bigint
	listp = &P5s
	for g.Assign(&tmp, *listp) != nil {
		*listp = tmp.next
		Free(tmp)
	}
}
