package standard

import (
	b "github.com/heyuuu/gophp/builtin"
)

func __setErrno(val int) __auto__ {
	errno = val
	return errno
}
func BF_decode(dst *BF_word, src *byte, size int) int {
	var dptr *uint8 = (*uint8)(dst)
	var end *uint8 = dptr + size
	var sptr *uint8 = (*uint8)(src)
	var tmp uint
	var c1 uint
	var c2 uint
	var c3 uint
	var c4 uint
	for {
		tmp = uint8(b.PostInc(&(*sptr)))
		if tmp == '$' {
			break
		}
		if uint(b.AssignOp(&tmp, "-=", 0x20) >= 0x60) != 0 {
			return -1
		}
		tmp = BF_atoi64[tmp]
		if tmp > 63 {
			return -1
		}
		c1 = tmp
		tmp = uint8(b.PostInc(&(*sptr)))
		if tmp == '$' {
			break
		}
		if uint(b.AssignOp(&tmp, "-=", 0x20) >= 0x60) != 0 {
			return -1
		}
		tmp = BF_atoi64[tmp]
		if tmp > 63 {
			return -1
		}
		c2 = tmp
		b.PostInc(&(*dptr)) = c1<<2 | (c2&0x30)>>4
		if dptr >= end {
			break
		}
		tmp = uint8(b.PostInc(&(*sptr)))
		if tmp == '$' {
			break
		}
		if uint(b.AssignOp(&tmp, "-=", 0x20) >= 0x60) != 0 {
			return -1
		}
		tmp = BF_atoi64[tmp]
		if tmp > 63 {
			return -1
		}
		c3 = tmp
		b.PostInc(&(*dptr)) = (c2&0xf)<<4 | (c3&0x3c)>>2
		if dptr >= end {
			break
		}
		tmp = uint8(b.PostInc(&(*sptr)))
		if tmp == '$' {
			break
		}
		if uint(b.AssignOp(&tmp, "-=", 0x20) >= 0x60) != 0 {
			return -1
		}
		tmp = BF_atoi64[tmp]
		if tmp > 63 {
			return -1
		}
		c4 = tmp
		b.PostInc(&(*dptr)) = (c3&0x3)<<6 | c4
		if dptr >= end {
			break
		}
	}
	if end-dptr == size {
		return -1
	}
	for dptr < end {
		b.PostInc(&(*dptr)) = 0
	}
	return 0
}
func BF_encode(dst *byte, src *BF_word, size int) {
	var sptr *uint8 = (*uint8)(src)
	var end *uint8 = sptr + size
	var dptr *uint8 = (*uint8)(dst)
	var c1 uint
	var c2 uint
	for {
		*sptr++
		c1 = (*sptr) - 1
		b.PostInc(&(*dptr)) = BF_itoa64[c1>>2]
		c1 = (c1 & 0x3) << 4
		if sptr >= end {
			b.PostInc(&(*dptr)) = BF_itoa64[c1]
			break
		}
		*sptr++
		c2 = (*sptr) - 1
		c1 |= c2 >> 4
		b.PostInc(&(*dptr)) = BF_itoa64[c1]
		c1 = (c2 & 0xf) << 2
		if sptr >= end {
			b.PostInc(&(*dptr)) = BF_itoa64[c1]
			break
		}
		*sptr++
		c2 = (*sptr) - 1
		c1 |= c2 >> 6
		b.PostInc(&(*dptr)) = BF_itoa64[c1]
		b.PostInc(&(*dptr)) = BF_itoa64[c2&0x3f]
		if sptr >= end {
			break
		}
	}
}
func BF_swap(x *BF_word, count int) {
	var endianness_check int = 1
	var is_little_endian *byte = (*byte)(&endianness_check)
	var tmp BF_word
	if *is_little_endian {
		for {
			tmp = *x
			tmp = tmp<<16 | tmp>>16
			b.PostInc(&(*x)) = (tmp&0xff00ff)<<8 | tmp>>8&0xff00ff
			if !(b.PreDec(&count)) {
				break
			}
		}
	}
}
func BF_INDEX(S __auto__, i *uint8) BF_word { return *((*BF_word)((*uint8)(S) + i)) }
func BF_ROUND(L BF_word, R BF_word, N int) {
	tmp1 = L & 0xff
	tmp1 <<= 2
	tmp2 = L >> 6
	tmp2 &= 0x3fc
	tmp3 = L >> 14
	tmp3 &= 0x3fc
	tmp4 = L >> 22
	tmp4 &= 0x3fc
	tmp1 = BF_INDEX(data.ctx.S[3], tmp1)
	tmp2 = BF_INDEX(data.ctx.S[2], tmp2)
	tmp3 = BF_INDEX(data.ctx.S[1], tmp3)
	tmp3 += BF_INDEX(data.ctx.S[0], tmp4)
	tmp3 ^= tmp2
	R ^= data.ctx.P[N+1]
	tmp3 += tmp1
	R ^= tmp3
}
func BF_body() {
	R = 0
	L = R
	ptr = data.ctx.P
	for {
		ptr += 2
		L ^= data.ctx.P[0]
		BF_ROUND(L, R, 0)
		BF_ROUND(R, L, 1)
		BF_ROUND(L, R, 2)
		BF_ROUND(R, L, 3)
		BF_ROUND(L, R, 4)
		BF_ROUND(R, L, 5)
		BF_ROUND(L, R, 6)
		BF_ROUND(R, L, 7)
		BF_ROUND(L, R, 8)
		BF_ROUND(R, L, 9)
		BF_ROUND(L, R, 10)
		BF_ROUND(R, L, 11)
		BF_ROUND(L, R, 12)
		BF_ROUND(R, L, 13)
		BF_ROUND(L, R, 14)
		BF_ROUND(R, L, 15)
		tmp4 = R
		R = L
		L = tmp4 ^ data.ctx.P[BF_N+1]
		*(ptr - 2) = L
		*(ptr - 1) = R
		if ptr >= data.ctx.P[BF_N+2] {
			break
		}
	}
	ptr = data.ctx.S[0]
	for {
		ptr += 2
		L ^= data.ctx.P[0]
		BF_ROUND(L, R, 0)
		BF_ROUND(R, L, 1)
		BF_ROUND(L, R, 2)
		BF_ROUND(R, L, 3)
		BF_ROUND(L, R, 4)
		BF_ROUND(R, L, 5)
		BF_ROUND(L, R, 6)
		BF_ROUND(R, L, 7)
		BF_ROUND(L, R, 8)
		BF_ROUND(R, L, 9)
		BF_ROUND(L, R, 10)
		BF_ROUND(R, L, 11)
		BF_ROUND(L, R, 12)
		BF_ROUND(R, L, 13)
		BF_ROUND(L, R, 14)
		BF_ROUND(R, L, 15)
		tmp4 = R
		R = L
		L = tmp4 ^ data.ctx.P[BF_N+1]
		*(ptr - 2) = L
		*(ptr - 1) = R
		if ptr >= data.ctx.S[3][0xff] {
			break
		}
	}
}
func BF_set_key(key *byte, expanded BF_key, initial BF_key, flags uint8) {
	var ptr *byte = key
	var bug uint
	var i uint
	var j uint
	var safety BF_word
	var sign BF_word
	var diff BF_word
	var tmp []BF_word

	/*
	 * There was a sign extension bug in older revisions of this function. While
	 * we would have liked to simply fix the bug and move on, we have to provide
	 * a backwards compatibility feature (essentially the bug) for some systems and
	 * a safety measure for some others. The latter is needed because for certain
	 * multiple inputs to the buggy algorithm there exist easily found inputs to
	 * the correct algorithm that produce the same hash. Thus, we optionally
	 * deviate from the correct algorithm just enough to avoid such collisions.
	 * While the bug itself affected the majority of passwords containing
	 * characters with the 8th bit set (although only a percentage of those in a
	 * collision-producing way), the anti-collision safety measure affects
	 * only a subset of passwords containing the '\xff' character (not even all of
	 * those passwords, just some of them). This character is not found in valid
	 * UTF-8 sequences and is rarely used in popular 8-bit character encodings.
	 * Thus, the safety measure is unlikely to cause much annoyance, and is a
	 * reasonable tradeoff to use when authenticating against existing hashes that
	 * are not reliably known to have been computed with the correct algorithm.
	 *
	 * We use an approach that tries to minimize side-channel leaks of password
	 * information - that is, we mostly use fixed-cost bitwise operations instead
	 * of branches or table lookups. (One conditional branch based on password
	 * length remains. It is not part of the bug aftermath, though, and is
	 * difficult and possibly unreasonable to avoid given the use of C strings by
	 * the caller, which results in similar timing leaks anyway.)
	 *
	 * For actual implementation, we set an array index in the variable "bug"
	 * (0 means no bug, 1 means sign extension bug emulation) and a flag in the
	 * variable "safety" (bit 16 is set when the safety measure is requested).
	 * Valid combinations of settings are:
	 *
	 * Prefix "$2a$": bug = 0, safety = 0x10000
	 * Prefix "$2b$": bug = 0, safety = 0
	 * Prefix "$2x$": bug = 1, safety = 0
	 * Prefix "$2y$": bug = 0, safety = 0
	 */

	bug = uint(flags & 1)
	safety = BF_word(flags&2) << 15
	diff = 0
	sign = diff
	for i = 0; i < BF_N+2; i++ {
		tmp[1] = 0
		tmp[0] = tmp[1]
		for j = 0; j < 4; j++ {
			tmp[0] <<= 8
			tmp[0] |= uint8(*ptr)
			tmp[1] <<= 8
			tmp[1] |= BF_word_signed(signed__char(*ptr))

			/*
			 * Sign extension in the first char has no effect - nothing to overwrite yet,
			 * and those extra 24 bits will be fully shifted out of the 32-bit word. For
			 * chars 2, 3, 4 in each four-char block, we set bit 7 of "sign" if sign
			 * extension in tmp[1] occurs. Once this flag is set, it remains set.
			 */

			if j != 0 {
				sign |= tmp[1] & 0x80
			}
			if !(*ptr) {
				ptr = key
			} else {
				ptr++
			}
		}
		diff |= tmp[0] ^ tmp[1]
		expanded[i] = tmp[bug]
		initial[i] = BF_init_state.GetP()[i] ^ tmp[bug]
	}

	/*
	 * At this point, "diff" is zero iff the correct and buggy algorithms produced
	 * exactly the same result. If so and if "sign" is non-zero, which indicates
	 * that there was a non-benign sign extension, this means that we have a
	 * collision between the correctly computed hash for this password and a set of
	 * passwords that could be supplied to the buggy algorithm. Our safety measure
	 * is meant to protect from such many-buggy to one-correct collisions, by
	 * deviating from the correct algorithm in such cases. Let's check for this.
	 */

	diff |= diff >> 16
	diff &= 0xffff
	diff += 0xffff
	sign <<= 9
	sign &= ^diff & safety

	/*
	 * If we have determined that we need to deviate from the correct algorithm,
	 * flip bit 16 in initial expanded key. (The choice of 16 is arbitrary, but
	 * let's stick to it now. It came out of the approach we used above, and it's
	 * not any worse than any other choice we could make.)
	 *
	 * It is crucial that we don't do the same to the expanded key used in the main
	 * Eksblowfish loop. By doing it to only one of these two, we deviate from a
	 * state that could be directly specified by a password to the buggy algorithm
	 * (and to the fully correct one as well, but that's a side-effect).
	 */

	initial[0] ^= sign

	/*
	 * If we have determined that we need to deviate from the correct algorithm,
	 * flip bit 16 in initial expanded key. (The choice of 16 is arbitrary, but
	 * let's stick to it now. It came out of the approach we used above, and it's
	 * not any worse than any other choice we could make.)
	 *
	 * It is crucial that we don't do the same to the expanded key used in the main
	 * Eksblowfish loop. By doing it to only one of these two, we deviate from a
	 * state that could be directly specified by a password to the buggy algorithm
	 * (and to the fully correct one as well, but that's a side-effect).
	 */
}
func BF_crypt(key *byte, setting *byte, output *byte, size int, min BF_word) *byte {
	var data struct {
		ctx          BF_ctx
		expanded_key BF_key
		binary       struct /* union */ {
			salt   []BF_word
			output []BF_word
		}
	}
	var L BF_word
	var R BF_word
	var tmp1 BF_word
	var tmp2 BF_word
	var tmp3 BF_word
	var tmp4 BF_word
	var ptr *BF_word
	var count BF_word
	var i int
	if size < 7+22+31+1 {
		__setErrno(ERANGE)
		return nil
	}
	if setting[0] != '$' || setting[1] != '2' || setting[2] < 'a' || setting[2] > 'z' || FlagsBySubtype[uint(uint8(setting[2]-'a'))] == 0 || setting[3] != '$' || setting[4] < '0' || setting[4] > '3' || setting[5] < '0' || setting[5] > '9' || setting[4] == '3' && setting[5] > '1' || setting[6] != '$' {
		__setErrno(EINVAL)
		return nil
	}
	count = BF_word(1<<(setting[4]-'0')*10 + (setting[5] - '0'))
	if count < min || BF_decode(data.binary.salt, &setting[7], 16) != 0 {
		__setErrno(EINVAL)
		return nil
	}
	BF_swap(data.binary.salt, 4)
	BF_set_key(key, data.expanded_key, data.ctx.GetP(), FlagsBySubtype[uint(uint8(setting[2]-'a'))])
	memcpy(data.ctx.GetS(), BF_init_state.GetS(), b.SizeOf("data . ctx . S"))
	R = 0
	L = R
	for i = 0; i < BF_N+2; i += 2 {
		L ^= data.binary.salt[i&2]
		R ^= data.binary.salt[(i&2)+1]
		L ^= data.ctx.GetP()[0]
		BF_ROUND(L, R, 0)
		BF_ROUND(R, L, 1)
		BF_ROUND(L, R, 2)
		BF_ROUND(R, L, 3)
		BF_ROUND(L, R, 4)
		BF_ROUND(R, L, 5)
		BF_ROUND(L, R, 6)
		BF_ROUND(R, L, 7)
		BF_ROUND(L, R, 8)
		BF_ROUND(R, L, 9)
		BF_ROUND(L, R, 10)
		BF_ROUND(R, L, 11)
		BF_ROUND(L, R, 12)
		BF_ROUND(R, L, 13)
		BF_ROUND(L, R, 14)
		BF_ROUND(R, L, 15)
		tmp4 = R
		R = L
		L = tmp4 ^ data.ctx.GetP()[BF_N+1]
		data.ctx.GetP()[i] = L
		data.ctx.GetP()[i+1] = R
	}
	ptr = data.ctx.GetS()[0]
	for {
		ptr += 4
		L ^= data.binary.salt[BF_N+2&3]
		R ^= data.binary.salt[BF_N+3&3]
		L ^= data.ctx.GetP()[0]
		BF_ROUND(L, R, 0)
		BF_ROUND(R, L, 1)
		BF_ROUND(L, R, 2)
		BF_ROUND(R, L, 3)
		BF_ROUND(L, R, 4)
		BF_ROUND(R, L, 5)
		BF_ROUND(L, R, 6)
		BF_ROUND(R, L, 7)
		BF_ROUND(L, R, 8)
		BF_ROUND(R, L, 9)
		BF_ROUND(L, R, 10)
		BF_ROUND(R, L, 11)
		BF_ROUND(L, R, 12)
		BF_ROUND(R, L, 13)
		BF_ROUND(L, R, 14)
		BF_ROUND(R, L, 15)
		tmp4 = R
		R = L
		L = tmp4 ^ data.ctx.GetP()[BF_N+1]
		*(ptr - 4) = L
		*(ptr - 3) = R
		L ^= data.binary.salt[BF_N+4&3]
		R ^= data.binary.salt[BF_N+5&3]
		L ^= data.ctx.GetP()[0]
		BF_ROUND(L, R, 0)
		BF_ROUND(R, L, 1)
		BF_ROUND(L, R, 2)
		BF_ROUND(R, L, 3)
		BF_ROUND(L, R, 4)
		BF_ROUND(R, L, 5)
		BF_ROUND(L, R, 6)
		BF_ROUND(R, L, 7)
		BF_ROUND(L, R, 8)
		BF_ROUND(R, L, 9)
		BF_ROUND(L, R, 10)
		BF_ROUND(R, L, 11)
		BF_ROUND(L, R, 12)
		BF_ROUND(R, L, 13)
		BF_ROUND(L, R, 14)
		BF_ROUND(R, L, 15)
		tmp4 = R
		R = L
		L = tmp4 ^ data.ctx.GetP()[BF_N+1]
		*(ptr - 2) = L
		*(ptr - 1) = R
		if ptr >= data.ctx.GetS()[3][0xff] {
			break
		}
	}
	for {
		var done int
		for i = 0; i < BF_N+2; i += 2 {
			data.ctx.GetP()[i] ^= data.expanded_key[i]
			data.ctx.GetP()[i+1] ^= data.expanded_key[i+1]
		}
		done = 0
		for {
			BF_body()
			if done != 0 {
				break
			}
			done = 1
			tmp1 = data.binary.salt[0]
			tmp2 = data.binary.salt[1]
			tmp3 = data.binary.salt[2]
			tmp4 = data.binary.salt[3]
			for i = 0; i < BF_N; i += 4 {
				data.ctx.GetP()[i] ^= tmp1
				data.ctx.GetP()[i+1] ^= tmp2
				data.ctx.GetP()[i+2] ^= tmp3
				data.ctx.GetP()[i+3] ^= tmp4
			}
			data.ctx.GetP()[16] ^= tmp1
			data.ctx.GetP()[17] ^= tmp2

		}
		if !(b.PreDec(&count)) {
			break
		}
	}
	for i = 0; i < 6; i += 2 {
		L = BF_magic_w[i]
		R = BF_magic_w[i+1]
		count = 64
		for {
			L ^= data.ctx.GetP()[0]
			BF_ROUND(L, R, 0)
			BF_ROUND(R, L, 1)
			BF_ROUND(L, R, 2)
			BF_ROUND(R, L, 3)
			BF_ROUND(L, R, 4)
			BF_ROUND(R, L, 5)
			BF_ROUND(L, R, 6)
			BF_ROUND(R, L, 7)
			BF_ROUND(L, R, 8)
			BF_ROUND(R, L, 9)
			BF_ROUND(L, R, 10)
			BF_ROUND(R, L, 11)
			BF_ROUND(L, R, 12)
			BF_ROUND(R, L, 13)
			BF_ROUND(L, R, 14)
			BF_ROUND(R, L, 15)
			tmp4 = R
			R = L
			L = tmp4 ^ data.ctx.GetP()[BF_N+1]
			if !(b.PreDec(&count)) {
				break
			}
		}
		data.binary.output[i] = L
		data.binary.output[i+1] = R
	}
	memcpy(output, setting, 7+22-1)
	output[7+22-1] = BF_itoa64[int(BF_atoi64[int(setting[7+22-1]-0x20)]&0x30)]

	/* This has to be bug-compatible with the original implementation, so
	 * only encode 23 of the 24 bytes. :-) */

	BF_swap(data.binary.output, 6)
	BF_encode(&output[7+22], data.binary.output, 23)
	output[7+22+31] = '0'
	return output
}
func _cryptOutputMagic(setting *byte, output *byte, size int) int {
	if size < 3 {
		return -1
	}
	output[0] = '*'
	output[1] = '0'
	output[2] = '0'
	if setting[0] == '*' && setting[1] == '0' {
		output[1] = '1'
	}
	return 0
}
func PhpCryptBlowfishRn(key *byte, setting *byte, output *byte, size int) *byte {
	var test_key *byte = "8b xd0xc1xd2xcfxccxd8"
	var test_setting *byte = "$2a$00$abcdefghijklmnopqrstuu"
	var test_hashes []*byte = []*byte{"i1D709vfamulimlGcq0qq3UvuUasvEa0x55", "VUrPmXD6q/nVSSp7pNDhCR9071IfIRe0x55"}
	var test_hash *byte = test_hashes[0]
	var retval *byte
	var p *byte
	var save_errno int
	var ok int
	var buf struct {
		s []byte
		o []byte
	}

	/* Hash the supplied password */

	_cryptOutputMagic(setting, output, size)
	retval = BF_crypt(key, setting, output, size, 16)
	save_errno = errno

	/*
	 * Do a quick self-test. It is important that we make both calls to BF_crypt()
	 * from the same scope such that they likely use the same stack locations,
	 * which makes the second call overwrite the first call's sensitive data on the
	 * stack and makes it more likely that any alignment related issues would be
	 * detected by the self-test.
	 */

	memcpy(buf.s, test_setting, b.SizeOf("buf . s"))
	if retval != nil {
		var flags uint = FlagsBySubtype[uint(uint8(setting[2]-'a'))]
		test_hash = test_hashes[flags&1]
		buf.s[2] = setting[2]
	}
	memset(buf.o, 0x55, b.SizeOf("buf . o"))
	buf.o[b.SizeOf("buf . o")-1] = 0
	p = BF_crypt(test_key, buf.s, buf.o, b.SizeOf("buf . o")-(1+1), 1)
	ok = p == buf.o && !(memcmp(p, buf.s, 7+22)) && !(memcmp(p+(7+22), test_hash, 31+1+1+1))
	var k *byte = "xffxa3" + "34" + "xffxffxffxa3" + "345"
	var ae BF_key
	var ai BF_key
	var ye BF_key
	var yi BF_key
	BF_set_key(k, ae, ai, 2)
	BF_set_key(k, ye, yi, 4)
	ai[0] ^= 0x10000
	ok = ok != 0 && ai[0] == 0xdb9c59bc && ye[17] == 0x33343500 && !(memcmp(ae, ye, b.SizeOf("ae"))) && !(memcmp(ai, yi, b.SizeOf("ai")))
	__setErrno(save_errno)
	if ok != 0 {
		return retval
	}

	/* Should not happen */

	_cryptOutputMagic(setting, output, size)
	__setErrno(EINVAL)
	return nil
}
