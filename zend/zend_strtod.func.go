// <<generate>>

package zend

import (
	b "sik/builtin"
)

func Word0(x *U) ULong  { return x.GetL()[1] }
func Word1(x *U) ULong  { return x.GetL()[0] }
func Dval(x *U) float64 { return x.GetD() }
func Storeinc(a int, b __auto__, c __auto__) int {
	(*uint16)(a)[1] = uint16(b)
	(*uint16)(a)[0] = uint16(c)
	a++
	return a - 1
}
func RoundedProduct(a float64, b float64) float64 {
	a *= b
	return a
}
func RoundedQuotient(a float64, b float64) float64 {
	a /= b
	return a
}
func ZendStartupStrtod() int { return 1 }
func ZendShutdownStrtod() int {
	DestroyFreelist()
	FreeP5s()
	return 1
}
func Balloc(k int) *Bigint {
	var x int
	var rv *Bigint

	/* The k > Kmax case does not need ACQUIRE_DTOA_LOCK(0), */

	if k <= Kmax && b.Assign(&rv, Freelist[k]) {
		Freelist[k] = rv.GetNext()
	} else {
		x = 1 << k
		rv = (*Bigint)(MALLOC(b.SizeOf("Bigint") + (x-1)*b.SizeOf("ULong")))
		if rv == nil {
			ZendErrorNoreturn(E_ERROR, "Balloc() failed to allocate memory")
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
		if v.GetK() > Kmax {
			Free(any(v))
		} else {
			v.SetNext(Freelist[v.GetK()])
			Freelist[v.GetK()] = v
		}
	}
}
func Bcopy(x *Bigint, y *Bigint) __auto__ {
	return memcpy((*byte)(x.GetSign()), (*byte)(y.GetSign()), y.GetWds()*b.SizeOf("Long")+2*b.SizeOf("int"))
}
func Multadd(b *Bigint, m int, a int) *Bigint {
	var i int
	var wds int
	var x *ULong
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
		b.PostInc(&(*x)) = y & FFFFFFFF
		if b.PreInc(&i) >= wds {
			break
		}
	}
	if carry {
		if wds >= b.GetMaxwds() {
			b1 = Balloc(b.GetK() + 1)
			Bcopy(b1, b)
			Bfree(b)
			b = b1
		}
		b.GetX()[b.PostInc(&wds)] = carry
		b.SetWds(wds)
	}
	return b
}
func S2b(s *byte, nd0 int, nd int, y9 ULong, dplen int) *Bigint {
	var b *Bigint
	var i int
	var k int
	var x Long
	var y Long
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
			b = Multadd(b, 10, b.PostInc(&(*s))-'0')
			if b.PreInc(&i) >= nd0 {
				break
			}
		}
		s += dplen
	} else {
		s += dplen + 9
	}
	for ; i < nd; i++ {
		b = Multadd(b, 10, b.PostInc(&(*s))-'0')
	}
	return b
}
func Hi0bits(x ULong) int {
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
func Lo0bits(y *ULong) int {
	var k int
	var x ULong = *y
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
		if !x {
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
	var x *ULong
	var xa *ULong
	var xae *ULong
	var xb *ULong
	var xbe *ULong
	var xc *ULong
	var xc0 *ULong
	var y ULong
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
		if b.Assign(&y, b.PostInc(&(*xb))) {
			x = xa
			xc = xc0
			carry = 0
			for {
				z = b.PostInc(&(*x)) * unsigned__long__long(y+(*xc)+carry)
				carry = z >> 32
				b.PostInc(&(*xc)) = z & FFFFFFFF
				if x >= xae {
					break
				}
			}
			*xc = carry
		}
	}
	xc0 = c.GetX()
	xc = xc0 + wc
	for ; wc > 0 && !(*(b.PreDec(&xc))); wc-- {

	}
	c.SetWds(wc)
	return c
}
func Pow5mult(b *Bigint, k int) *Bigint {
	var b1 *Bigint
	var p5 *Bigint
	var p51 *Bigint
	var i int
	var p05 []int = []int{5, 25, 125}
	if b.Assign(&i, k&3) {
		b = Multadd(b, p05[i-1], 0)
	}
	if !(b.AssignOp(&k, ">>=", 2)) {
		return b
	}
	if !(b.Assign(&p5, P5s)) {

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
		if !(b.AssignOp(&k, ">>=", 1)) {
			break
		}
		if !(b.Assign(&p51, p5.GetNext())) {
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
	var x *ULong
	var x1 *ULong
	var xe *ULong
	var z ULong
	n = k >> 5
	k1 = b.GetK()
	n1 = n + b.GetWds() + 1
	for i = b.GetMaxwds(); n1 > i; i <<= 1 {
		k1++
	}
	b1 = Balloc(k1)
	x1 = b1.GetX()
	for i = 0; i < n; i++ {
		b.PostInc(&(*x1)) = 0
	}
	x = b.GetX()
	xe = x + b.GetWds()
	if b.AssignOp(&k, "&=", 0x1f) {
		k1 = 32 - k
		z = 0
		for {
			b.PostInc(&(*x1)) = (*x)<<k | z
			z = b.PostInc(&(*x)) >> k1
			if x >= xe {
				break
			}
		}
		if b.Assign(&(*x1), z) {
			n1++
		}
	} else {
		for {
			*x++
			b.PostInc(&(*x1)) = (*x) - 1
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
	var xa *ULong
	var xa0 *ULong
	var xb *ULong
	var xb0 *ULong
	var i int
	var j int
	i = a.GetWds()
	j = b.GetWds()
	if b.AssignOp(&i, "-=", j) {
		return i
	}
	xa0 = a.GetX()
	xa = xa0 + j
	xb0 = b.GetX()
	xb = xb0 + j
	for {
		if (*(b.PreDec(&xa))) != (*(b.PreDec(&xb))) {
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
	var xa *ULong
	var xae *ULong
	var xb *ULong
	var xbe *ULong
	var xc *ULong
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
		y = unsigned__long__long(b.PostInc(&(*xa)) - b.PostInc(&(*xb)) - borrow)
		borrow = y >> 32 & ULong(1)
		b.PostInc(&(*xc)) = y & FFFFFFFF
		if xb >= xbe {
			break
		}
	}
	for xa < xae {
		y = b.PostInc(&(*xa)) - borrow
		borrow = y >> 32 & ULong(1)
		b.PostInc(&(*xc)) = y & FFFFFFFF
	}
	for !(*(b.PreDec(&xc))) {
		wa--
	}
	c.SetWds(wa)
	return c
}
func Ulp(x *U) float64 {
	var L Long
	var u U
	L = (Word0(x) & Exp_mask) - (P-1)*Exp_msk1
	Word0(&u) = L
	Word1(&u) = 0
	return u.GetD()
}
func B2d(a *Bigint, e *int) float64 {
	var xa *ULong
	var xa0 *ULong
	var w ULong
	var y ULong
	var z ULong
	var k int
	var d U

	// #define d0       word0 ( & d )

	// #define d1       word1 ( & d )

	xa0 = a.GetX()
	xa = xa0 + a.GetWds()
	y = *(b.PreDec(&xa))
	k = Hi0bits(y)
	*e = 32 - k
	if k < Ebits {
		Word0(&d) = Exp_1 | y>>Ebits - k
		if xa > xa0 {
			w = *(b.PreDec(&xa))
		} else {
			w = 0
		}
		Word1(&d) = y<<32 - Ebits + k | w>>Ebits - k
		goto ret_d
	}
	if xa > xa0 {
		z = *(b.PreDec(&xa))
	} else {
		z = 0
	}
	if b.AssignOp(&k, "-=", Ebits) {
		Word0(&d) = Exp_1 | y<<k | z>>32 - k
		if xa > xa0 {
			y = *(b.PreDec(&xa))
		} else {
			y = 0
		}
		Word1(&d) = z<<k | y>>32 - k
	} else {
		Word0(&d) = Exp_1 | y
		Word1(&d) = z
	}
ret_d:
	return d.GetD()
}
func D2b(d *U, e *int, bits *int) *Bigint {
	var b *Bigint
	var de int
	var k int
	var x *ULong
	var y ULong
	var z ULong
	var i int

	// #define d0       word0 ( d )

	// #define d1       word1 ( d )

	b = Balloc(1)
	x = b.GetX()
	z = Word0(d) & Frac_mask
	Word0(d) &= 0x7fffffff
	if b.Assign(&de, int(Word0(d)>>Exp_shift)) {
		z |= Exp_msk1
	}
	if b.Assign(&y, Word1(d)) {
		if b.Assign(&k, Lo0bits(&y)) {
			x[0] = y | z<<32 - k
			z >>= k
		} else {
			x[0] = y
		}
		if b.Assign(&x[1], z) {
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
		*e = de - Bias - (P - 1) + k
		*bits = P - k
	} else {
		*e = de - Bias - (P - 1) + 1 + k
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
	da.SetD(B2d(a, &ka))
	db.SetD(B2d(b, &kb))
	k = ka - kb + 32*(a.GetWds()-b.GetWds())
	if k > 0 {
		Word0(&da) += k * Exp_msk1
	} else {
		k = -k
		Word0(&db) += k * Exp_msk1
	}
	return da.GetD() / db.GetD()
}
func Dshift(b *Bigint, p2 int) int {
	var rv int = Hi0bits(b.GetX()[b.GetWds()-1]) - 4
	if p2 > 0 {
		rv -= p2
	}
	return rv & Kmask
}
func Quorem(b *Bigint, S *Bigint) int {
	var n int
	var bx *ULong
	var bxe *ULong
	var q ULong
	var sx *ULong
	var sxe *ULong
	var borrow unsigned__long__long
	var carry unsigned__long__long
	var y unsigned__long__long
	var ys unsigned__long__long
	n = S.GetWds()
	if b.GetWds() < n {
		return 0
	}
	sx = S.GetX()
	sxe = sx + b.PreDec(&n)
	bx = b.GetX()
	bxe = bx + n
	q = (*bxe) / ((*sxe) + 1)
	if q {
		borrow = 0
		carry = 0
		for {
			ys = b.PostInc(&(*sx)) * unsigned__long__long(q+carry)
			carry = ys >> 32
			y = (*bx) - (ys & FFFFFFFF) - borrow
			borrow = y >> 32 & ULong(1)
			b.PostInc(&(*bx)) = y & FFFFFFFF
			if sx > sxe {
				break
			}
		}
		if !(*bxe) {
			bx = b.GetX()
			for b.PreDec(&bxe) > bx && !(*bxe) {
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
			ys = b.PostInc(&(*sx)) + carry
			carry = ys >> 32
			y = (*bx) - (ys & FFFFFFFF) - borrow
			borrow = y >> 32 & ULong(1)
			b.PostInc(&(*bx)) = y & FFFFFFFF
			if sx > sxe {
				break
			}
		}
		bx = b.GetX()
		bxe = bx + n
		if !(*bxe) {
			for b.PreDec(&bxe) > bx && !(*bxe) {
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
	if bc.GetScale() == 0 || b.Assign(&i, 2*P+1-((Word0(x)&Exp_mask)>>Exp_shift)) <= 0 {
		return rv
	}
	Word0(&u) = Exp_1 + (i << Exp_shift)
	Word1(&u) = 0
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
		p2 = Emin - P + 1
		bbits = 1
		Word0(rv) = P + 2<<Exp_shift
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

	i = P - bbits
	if i > b.Assign(&j, P-Emin-1+p2) {
		i = j
	}
	b = Lshift(b, b.PreInc(&i))
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
	if b.AssignOp(&b2, "+=", i) > 0 {
		b = Lshift(b, b2)
	}
	if b.AssignOp(&d2, "+=", i) > 0 {
		d = Lshift(d, d2)
	}

	/* Now b/d = exactly half-way between the two floating-point values */

	if !(b.Assign(&dig, Quorem(b, d))) {
		b = Multadd(b, 10, 0)
		dig = Quorem(b, d)
	}

	/* Compare b/d with s0 */

	for i = 0; i < nd0; {
		if b.Assign(&dd, s0[b.PostInc(&i)]-'0'-dig) {
			goto ret
		}
		if !(b.GetX()[0]) && b.GetWds() == 1 {
			if i < nd {
				dd = 1
			}
			goto ret
		}
		b = Multadd(b, 10, 0)
		dig = Quorem(b, d)
	}
	for j = bc.GetDp1(); b.PostInc(&i) < nd; {
		if b.Assign(&dd, s0[b.PostInc(&j)]-'0'-dig) {
			goto ret
		}
		if !(b.GetX()[0]) && b.GetWds() == 1 {
			if i < nd {
				dd = 1
			}
			goto ret
		}
		b = Multadd(b, 10, 0)
		dig = Quorem(b, d)
	}
	if dig > 0 || b.GetX()[0] || b.GetWds() > 1 {
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

		if b.Assign(&j, ((Word0(rv)&Exp_mask)>>Exp_shift)-bc.GetScale()) <= 0 {
			i = 1 - j
			if i <= 31 {
				if (Word1(rv) & 0x1 << i) != 0 {
					goto odd
				}
			} else if (Word0(rv)&0x1<<i - 32) != 0 {
				goto odd
			}
		} else if (Word1(rv) & 1) != 0 {
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
	var L Long
	var aadj2 U
	var adj U
	var rv U
	var rv0 U
	var y ULong
	var z ULong
	var bc BCinfo
	var bb *Bigint
	var bb1 *Bigint
	var bd *Bigint
	var bd0 *Bigint
	var bs *Bigint
	var delta *Bigint
	var Lsb ULong
	var Lsb1 ULong
	var req_bigcomp int = 0
	bc.SetUflchk(0)
	bc.SetDplen(bc.GetUflchk())
	nz = bc.GetDplen()
	nz1 = nz
	nz0 = nz1
	sign = nz0
	rv.SetD(0.0)
	for s = s00; ; s++ {
		switch *s {
		case '-':
			sign = 1
			fallthrough
		case '+':
			if *(b.PreInc(&s)) {
				goto break2
			}
			fallthrough
		case 0:
			goto ret0
			fallthrough
		case '\t':
			fallthrough
		case '\n':
			fallthrough
		case 'v':
			fallthrough
		case 'f':
			fallthrough
		case '\r':
			fallthrough
		case ' ':
			continue
		default:
			goto break2
		}
	}
break2:
	if (*s) == '0' {
		nz0 = 1
		for (*(b.PreInc(&s))) == '0' {

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
	for b.Assign(&c, *s) >= '0' && c <= '9' {
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
	for s1 = s; s1 > s0 && (*(b.PreDec(&s1))) == '0'; {
		nz1++
	}
	if c == '.' {
		c = *(b.PreInc(&s))
		bc.SetDp1(s - s0)
		bc.SetDplen(bc.GetDp1() - bc.GetDp0())
		if nd == 0 {
			for ; c == '0'; c = *(b.PreInc(&s)) {
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
		for ; c >= '0' && c <= '9'; c = *(b.PreInc(&s)) {
		have_dig:
			nz++
			if b.AssignOp(&c, "-=", '0') {
				nf += nz
				for i = 1; i < nz; i++ {
					if b.PostInc(&nd) < 9 {
						y *= 10
					} else if nd <= DBL_DIG+2 {
						z *= 10
					}
				}
				if b.PostInc(&nd) < 9 {
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
		switch b.Assign(&c, *(b.PreInc(&s))) {
		case '-':
			esign = 1
			fallthrough
		case '+':
			c = *(b.PreInc(&s))
		}
		if c >= '0' && c <= '9' {
			for c == '0' {
				c = *(b.PreInc(&s))
			}
			if c > '0' && c <= '9' {
				L = c - '0'
				s1 = s
				for b.Assign(&c, *(b.PreInc(&s))) >= '0' && c <= '9' {
					L = Long(10 * ULong(L+(c-'0')))
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
	rv.SetD(y)
	if k > 9 {
		rv.SetD(Tens[k-9]*rv.GetD() + z)
	}
	bd0 = 0
	if nd <= DBL_DIG && Flt_Rounds == 1 {
		if e == 0 {
			goto ret
		}
		if e > 0 {
			if e <= Ten_pmax {

				/* rv = */

				RoundedProduct(rv.GetD(), Tens[e])
				goto ret
			}
			i = DBL_DIG - nd
			if e <= Ten_pmax+i {

				/* A fancier test would sometimes let us do
				 * this for larger i values.
				 */

				e -= i
				rv.SetD(rv.GetD() * Tens[i])

				/* rv = */

				RoundedProduct(rv.GetD(), Tens[e])
				goto ret
			}
		} else if e >= -Ten_pmax {

			/* rv = */

			RoundedQuotient(rv.GetD(), Tens[-e])
			goto ret
		}
	}
	e1 += nd - k
	bc.SetScale(0)

	/* Get starting approximation = rv * 10**e1 */

	if e1 > 0 {
		if b.Assign(&i, e1&15) {
			rv.SetD(rv.GetD() * Tens[i])
		}
		if b.AssignOp(&e1, "&=", ^15) {
			if e1 > DBL_MAX_10_EXP {
			ovfl:

				/* Can't trust HUGE_VAL */

				Word0(&rv) = Exp_mask
				Word1(&rv) = 0
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
					rv.SetD(rv.GetD() * Bigtens[j])
				}
				j++
				e1 >>= 1
			}

			/* The last multiplication could overflow. */

			Word0(&rv) -= P * Exp_msk1
			rv.SetD(rv.GetD() * Bigtens[j])
			if b.Assign(&z, Word0(&rv)&Exp_mask) > Exp_msk1*(DBL_MAX_EXP+Bias-P) {
				goto ovfl
			}
			if z > Exp_msk1*(DBL_MAX_EXP+Bias-1-P) {

				/* set to largest number */

				Word0(&rv) = Big0
				Word1(&rv) = Big1
			} else {
				Word0(&rv) += P * Exp_msk1
			}
		}
	} else if e1 < 0 {
		e1 = -e1
		if b.Assign(&i, e1&15) {
			rv.SetD(rv.GetD() / Tens[i])
		}
		if b.AssignOp(&e1, ">>=", 4) {
			if e1 >= 1<<NBigtens {
				goto undfl
			}
			if (e1 & Scale_Bit) != 0 {
				bc.SetScale(2 * P)
			}
			for j = 0; e1 > 0; {
				if (e1 & 1) != 0 {
					rv.SetD(rv.GetD() * Tinytens[j])
				}
				j++
				e1 >>= 1
			}
			if bc.GetScale() != 0 && b.Assign(&j, 2*P+1-((Word0(&rv)&Exp_mask)>>Exp_shift)) > 0 {

				/* scaled rv is denormal; clear j low bits */

				if j >= 32 {
					if j > 54 {
						goto undfl
					}
					Word1(&rv) = 0
					if j >= 53 {
						Word0(&rv) = (P + 2) * Exp_msk1
					} else {
						Word0(&rv) &= 0xffffffff<<j - 32
					}
				} else {
					Word1(&rv) &= 0xffffffff << j
				}

				/* scaled rv is denormal; clear j low bits */

			}
			if !(rv.GetD()) {
			undfl:
				rv.SetD(0.0)
				goto range_err
			}
		}
	}

	/* Now the hard part -- adjusting rv to the correct value.*/

	bc.SetNd(nd - nz1)
	bc.SetNd0(nd0)

	/* to silence an erroneous warning about bc.nd0 */

	if nd > StrtodDiglim {

		/* ASSERT(strtod_diglim >= 18); 18 == one more than the */

		j = 18
		i = j
		if i > nd0 {
			j += bc.GetDplen()
		}
		for {
			if b.PreDec(&j) < bc.GetDp1() && j >= bc.GetDp0() {
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
				y = 10*y + s0[b.PostInc(&j)] - '0'
			}
		}
	}
	bd0 = S2b(s0, nd0, nd, y, bc.GetDplen())
	for {
		bd = Balloc(bd0.GetK())
		Bcopy(bd, bd0)
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
		Lsb = LSB
		Lsb1 = 0
		j = bbe - bc.GetScale()
		i = j + bbbits - 1
		j = P + 1 - bbbits
		if i < Emin {
			i = Emin - i
			j -= i
			if i < 32 {
				Lsb <<= i
			} else if i < 52 {
				Lsb1 = Lsb<<i - 32
			} else {
				Lsb1 = Exp_mask
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

			if bc.GetDsign() != 0 || Word1(&rv) || (Word0(&rv)&Bndry_mask) != 0 || (Word0(&rv)&Exp_mask) <= (2*P+1)*Exp_msk1 {
				break
			}
			if !(delta.GetX()[0]) && delta.GetWds() <= 1 {

				/* exact result */

				break

				/* exact result */

			}
			delta = Lshift(delta, Log2P)
			if Cmp(delta, bs) > 0 {
				goto drop_down
			}
			break
		}
		if i == 0 {

			/* exactly half-way between */

			if bc.GetDsign() != 0 {
				if (Word0(&rv)&Bndry_mask1) == Bndry_mask1 && Word1(&rv) == b.Cond(bc.GetScale() != 0 && b.Assign(&y, Word0(&rv)&Exp_mask) <= 2*P*Exp_msk1, 0xffffffff&0xffffffff<<2*P+1-(y>>Exp_shift), 0xffffffff) {

					/*boundary case -- increment exponent*/

					if Word0(&rv) == Big0 && Word1(&rv) == Big1 {
						goto ovfl
					}
					Word0(&rv) = (Word0(&rv) & Exp_mask) + Exp_msk1
					Word1(&rv) = 0
					bc.SetDsign(0)
					break
				}
			} else if (Word0(&rv)&Bndry_mask) == 0 && !(Word1(&rv)) {
			drop_down:

				/* boundary case -- decrement exponent */

				if bc.GetScale() != 0 {
					L = Word0(&rv) & Exp_mask
					if L <= (2*P+1)*Exp_msk1 {
						if L > (P+2)*Exp_msk1 {

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
				L = (Word0(&rv) & Exp_mask) - Exp_msk1
				Word0(&rv) = L | Bndry_mask1
				Word1(&rv) = 0xffffffff
				if bc.GetNd() > nd {
					goto cont
				}
				break
			}
			if Lsb1 {
				if (Word0(&rv) & Lsb1) == 0 {
					break
				}
			} else if (Word1(&rv) & Lsb) == 0 {
				break
			}
			if bc.GetDsign() != 0 {
				rv.SetD(rv.GetD() + Sulp(&rv, &bc))
			} else {
				rv.SetD(rv.GetD() - Sulp(&rv, &bc))
				if !(rv.GetD()) {
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
		if b.Assign(&aadj, Ratio(delta, bs)) <= 2.0 {
			if bc.GetDsign() != 0 {
				aadj1 = 1.0
				aadj = aadj1
			} else if Word1(&rv) || (Word0(&rv)&Bndry_mask) != 0 {
				if Word1(&rv) == Tiny1 && !(Word0(&rv)) {
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
			if Flt_Rounds == 0 {
				aadj1 += 0.5
			}
		}
		y = Word0(&rv) & Exp_mask

		/* Check for overflow */

		if y == Exp_msk1*(DBL_MAX_EXP+Bias-1) {
			rv0.SetD(rv.GetD())
			Word0(&rv) -= P * Exp_msk1
			adj.SetD(aadj1 * Ulp(&rv))
			rv.SetD(rv.GetD() + adj.GetD())
			if (Word0(&rv) & Exp_mask) >= Exp_msk1*(DBL_MAX_EXP+Bias-P) {
				if Word0(&rv0) == Big0 && Word1(&rv0) == Big1 {
					goto ovfl
				}
				Word0(&rv) = Big0
				Word1(&rv) = Big1
				goto cont
			} else {
				Word0(&rv) += P * Exp_msk1
			}
		} else {
			if bc.GetScale() != 0 && y <= 2*P*Exp_msk1 {
				if aadj <= 0x7fffffff {
					if b.Assign(&z, aadj) <= 0 {
						z = 1
					}
					aadj = z
					if bc.GetDsign() != 0 {
						aadj1 = aadj
					} else {
						aadj1 = -aadj
					}
				}
				aadj2.SetD(aadj1)
				Word0(&aadj2) += (2*P+1)*Exp_msk1 - y
				aadj1 = aadj2.GetD()
				adj.SetD(aadj1 * Ulp(&rv))
				rv.SetD(rv.GetD() + adj.GetD())
				if rv.GetD() == 0.0 {
					req_bigcomp = 1
					break
				}
			} else {
				adj.SetD(aadj1 * Ulp(&rv))
				rv.SetD(rv.GetD() + adj.GetD())
			}
		}
		z = Word0(&rv) & Exp_mask
		if bc.GetNd() == nd {
			if bc.GetScale() == 0 {
				if y == z {

					/* Can we stop now? */

					L = Long(aadj)
					aadj -= L

					/* The tolerances below are conservative. */

					if bc.GetDsign() != 0 || Word1(&rv) || (Word0(&rv)&Bndry_mask) != 0 {
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
		y = Word0(&rv) & Exp_mask
		if y == Exp_mask {
			goto ovfl
		}
		if y == 0 && rv.GetD() == 0.0 {
			goto undfl
		}
	}
	if bc.GetScale() != 0 {
		Word0(&rv0) = Exp_1 - 2*P*Exp_msk1
		Word1(&rv0) = 0
		rv.SetD(rv.GetD() * rv0.GetD())
	}
ret:
	if se != nil {
		*se = (*byte)(s)
	}
	if sign != 0 {
		return -(rv.GetD())
	} else {
		return rv.GetD()
	}
}
func RvAlloc(i int) *byte {
	var j int
	var k int
	var r *int
	j = b.SizeOf("ULong")
	for k = 0; b.SizeOf("Bigint")-b.SizeOf("ULong")-b.SizeOf("int")+int(j <= int(i)) != 0; j <<= 1 {
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
	for b.Assign(&(*t), b.PostInc(&(*s))) {
		t++
	}
	if rve != nil {
		*rve = t
	}
	return rv
}
func ZendFreedtoa(s *byte) {
	var b *Bigint = (*Bigint)((*int)(s - 1))
	b.SetMaxwds(1 << b.Assign(&(b.GetK()), *((*int)(b))))
	Bfree(b)
	if s == DtoaResult {
		DtoaResult = 0
	}
}
func ZendDtoa(
	dd float64,
	mode int,
	ndigits int,
	decpt *int,
	sign *int,
	rve **byte,
) *byte {
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
	var L Long
	var denorm int
	var x ULong
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
	if (Word0(&u) & Sign_bit) != 0 {

		/* set sign for everything, including 0's and NaNs */

		*sign = 1
		Word0(&u) &= ^Sign_bit
	} else {
		*sign = 0
	}
	if (Word0(&u) & Exp_mask) == Exp_mask {

		/* Infinity or NaN */

		*decpt = 9999
		if !(Word1(&u)) && (Word0(&u)&0xfffff) == 0 {
			return NrvAlloc("Infinity", rve, 8)
		}
		return NrvAlloc("NaN", rve, 3)
	}
	if !(u.GetD()) {
		*decpt = 1
		return NrvAlloc("0", rve, 1)
	}
	b = D2b(&u, &be, &bbits)
	if b.Assign(&i, int(Word0(&u)>>Exp_shift1&Exp_mask>>Exp_shift1)) {
		d2.SetD(u.GetD())
		Word0(&d2) &= Frac_mask1
		Word0(&d2) |= Exp_11

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

		i -= Bias
		denorm = 0
	} else {

		/* d is denormalized */

		i = bbits + be + (Bias + (P - 1) - 1)
		if i > 32 {
			x = Word0(&u)<<64 - i | Word1(&u)>>i - 32
		} else {
			x = Word1(&u)<<32 - i
		}
		d2.SetD(x)
		Word0(&d2) -= 31 * Exp_msk1
		i -= Bias + (P - 1) - 1 + 1
		denorm = 1
	}
	ds = (d2.GetD()-1.5)*0.28952965 + 0.17609125 + i*0.30103
	k = int(ds)
	if ds < 0.0 && ds != k {
		k--
	}
	k_check = 1
	if k >= 0 && k <= Ten_pmax {
		if u.GetD() < Tens[k] {
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
		fallthrough
	case 1:
		i = 18
		ndigits = 0
	case 2:
		leftright = 0
		fallthrough
	case 4:
		if ndigits <= 0 {
			ndigits = 1
		}
		i = ndigits
		ilim1 = i
		ilim = ilim1
	case 3:
		leftright = 0
		fallthrough
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
	if ilim >= 0 && ilim <= Quick_max && try_quick != 0 {

		/* Try to get by with floating-point arithmetic. */

		i = 0
		d2.SetD(u.GetD())
		k0 = k
		ilim0 = ilim
		ieps = 2
		if k > 0 {
			ds = Tens[k&0xf]
			j = k >> 4
			if (j & Bletch) != 0 {

				/* prevent overflows */

				j &= Bletch - 1
				u.SetD(u.GetD() / Bigtens[NBigtens-1])
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
			u.SetD(u.GetD() / ds)
		} else if b.Assign(&j1, -k) {
			u.SetD(u.GetD() * Tens[j1&0xf])
			for j = j1 >> 4; j != 0; {
				if (j & 1) != 0 {
					ieps++
					u.SetD(u.GetD() * Bigtens[i])
				}
				j >>= 1
				i++
			}
		}
		if k_check != 0 && u.GetD() < 1.0 && ilim > 0 {
			if ilim1 <= 0 {
				goto fast_failed
			}
			ilim = ilim1
			k--
			u.SetD(u.GetD() * 10.0)
			ieps++
		}
		eps.SetD(ieps*u.GetD() + 7.0)
		Word0(&eps) -= (P - 1) * Exp_msk1
		if ilim == 0 {
			mhi = 0
			S = mhi
			u.SetD(u.GetD() - 5.0)
			if u.GetD() > eps.GetD() {
				goto one_digit
			}
			if u.GetD() < -(eps.GetD()) {
				goto no_digits
			}
			goto fast_failed
		}
		if leftright != 0 {

			/* Use Steele & White method of only
			 * generating digits needed.
			 */

			eps.SetD(0.5/Tens[ilim-1] - eps.GetD())
			if k0 < 0 && j1 >= 307 {
				eps1.SetD(Infinity)
				Word0(&eps1) -= Exp_msk1 * (Bias + P - 1)
				eps1.SetD(eps1.GetD() * Tens[j1&0xf])
				i = 0
				j = j1 - 256>>4
				for j != 0 {
					if (j & 1) != 0 {
						eps1.SetD(eps1.GetD() * Bigtens[i])
					}
					j >>= 1
					i++
				}
				if eps.GetD() < eps1.GetD() {
					eps.SetD(eps1.GetD())
				}
			}
			for i = 0; ; {
				L = u.GetD()
				u.SetD(u.GetD() - L)
				b.PostInc(&(*s)) = '0' + int(L)
				if 1.0-u.GetD() < eps.GetD() {
					goto bump_up
				}
				if u.GetD() < eps.GetD() {
					goto ret1
				}
				if b.PreInc(&i) >= ilim {
					break
				}
				eps.SetD(eps.GetD() * 10.0)
				u.SetD(u.GetD() * 10.0)
			}
		} else {

			/* Generate ilim digits, then fix them up. */

			eps.SetD(eps.GetD() * Tens[ilim-1])
			for i = 1; ; {
				L = Long(u.GetD())
				if !(b.AssignOp(&(u.GetD()), "-=", L)) {
					ilim = i
				}
				b.PostInc(&(*s)) = '0' + int(L)
				if i == ilim {
					if u.GetD() > 0.5+eps.GetD() {
						goto bump_up
					} else if u.GetD() < 0.5-eps.GetD() {
						for (*(b.PreDec(&s))) == '0' {

						}
						s++
						goto ret1
					}
					break
				}
				i++
				u.SetD(u.GetD() * 10.0)
			}
		}
	fast_failed:
		s = s0
		u.SetD(d2.GetD())
		k = k0
		ilim = ilim0
	}

	/* Do we have a "small" integer? */

	if be >= 0 && k <= Int_max {

		/* Yes. */

		ds = Tens[k]
		if ndigits < 0 && ilim <= 0 {
			mhi = 0
			S = mhi
			if ilim < 0 || u.GetD() <= 5*ds {
				goto no_digits
			}
			goto one_digit
		}
		for i = 1; ; {
			L = Long(u.GetD() / ds)
			u.SetD(u.GetD() - L*ds)
			b.PostInc(&(*s)) = '0' + int(L)
			if !(u.GetD()) {
				break
			}
			if i == ilim {
				u.SetD(u.GetD() + u.GetD())
				if u.GetD() > ds || u.GetD() == ds && (L&1) != 0 {
				bump_up:
					for (*(b.PreDec(&s))) == '9' {
						if s == s0 {
							k++
							*s = '0'
							break
						}
					}
					b.PreInc(&(*s))++
				}
				break
			}
			i++
			u.SetD(u.GetD() * 10.0)
		}
		goto ret1
	}
	m2 = b2
	m5 = b5
	mlo = 0
	mhi = mlo
	if leftright != 0 {
		if denorm != 0 {
			i = be + (Bias + (P - 1) - 1 + 1)
		} else {
			i = 1 + P - bbits
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
			if b.Assign(&j, b5-m5) {
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
		if !(Word1(&u)) && (Word0(&u)&Bndry_mask) == 0 && (Word0(&u)&(Exp_mask & ^Exp_msk1)) != 0 {

			/* The special case */

			b2 += Log2P
			s2 += Log2P
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
		if ilim < 0 || Cmp(b, b.Assign(&S, Multadd(S, 5, 0))) <= 0 {

			/* no digits, fcvt style */

		no_digits:
			k = -1 - ndigits
			goto ret
		}
	one_digit:
		b.PostInc(&(*s)) = '1'
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
			Bcopy(mhi, mlo)
			mhi = Lshift(mhi, Log2P)
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
			if j1 == 0 && mode != 1 && (Word1(&u)&1) == 0 {
				if dig == '9' {
					goto round_9_up
				}
				if j > 0 {
					dig++
				}
				b.PostInc(&(*s)) = dig
				goto ret
			}
			if j < 0 || j == 0 && mode != 1 && (Word1(&u)&1) == 0 {
				if !(b.GetX()[0]) && b.GetWds() <= 1 {
					goto accept_dig
				}
				if j1 > 0 {
					b = Lshift(b, 1)
					j1 = Cmp(b, S)
					if (j1 > 0 || j1 == 0 && (dig&1) != 0) && b.PostInc(&dig) == '9' {
						goto round_9_up
					}
				}
			accept_dig:
				b.PostInc(&(*s)) = dig
				goto ret
			}
			if j1 > 0 {
				if dig == '9' {
				round_9_up:
					b.PostInc(&(*s)) = '9'
					goto roundoff
				}
				b.PostInc(&(*s)) = dig + 1
				goto ret
			}
			b.PostInc(&(*s)) = dig
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
			b.PostInc(&(*s)) = dig
			if !(b.GetX()[0]) && b.GetWds() <= 1 {
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
		for (*(b.PreDec(&s))) == '9' {
			if s == s0 {
				k++
				b.PostInc(&(*s)) = '1'
				goto ret
			}
		}
		b.PreInc(&(*s))++
	} else {
		for (*(b.PreDec(&s))) == '0' {

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
	for b.Assign(&c, b.PostInc(&(*s))) {
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
	for b.Assign(&c, b.PostInc(&(*s))) {
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
	for b.Assign(&c, b.PostInc(&(*s))) {

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
		*endptr = (*byte)(b.Cond(any != 0, s-1, str))
	}
	return value
}
func DestroyFreelist() {
	var i int
	var tmp *Bigint
	for i = 0; i <= Kmax; i++ {
		var listp **Bigint = &Freelist[i]
		for b.Assign(&tmp, *listp) != nil {
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
	for b.Assign(&tmp, *listp) != nil {
		*listp = tmp.GetNext()
		Free(tmp)
	}
}
