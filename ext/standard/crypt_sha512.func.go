// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/sapi/cli"
	"sik/zend"
	"sik/zend/types"
)

func Sha512ProcessBlock(buffer any, len_ int, ctx *Sha512Ctx) {
	var words *uint64 = buffer
	var nwords int = len_ / b.SizeOf("uint64_t")
	var a uint64 = ctx.GetH()[0]
	var b uint64 = ctx.GetH()[1]
	var c uint64 = ctx.GetH()[2]
	var d uint64 = ctx.GetH()[3]
	var e uint64 = ctx.GetH()[4]
	var f uint64 = ctx.GetH()[5]
	var g uint64 = ctx.GetH()[6]
	var h uint64 = ctx.GetH()[7]

	/* First increment the byte count.  FIPS 180-2 specifies the possible
	   length of the file up to 2^128 bits.  Here we only compute the
	   number of bytes.  Do a double word increment.  */

	ctx.GetTotal()[0] += len_
	if ctx.GetTotal()[0] < len_ {
		ctx.GetTotal()[1]++
	}

	/* Process all bytes in the buffer with 128 bytes in each round of
	   the loop.  */

	for nwords > 0 {
		var W []uint64
		var a_save uint64 = a
		var b_save uint64 = b
		var c_save uint64 = c
		var d_save uint64 = d
		var e_save uint64 = e
		var f_save uint64 = f
		var g_save uint64 = g
		var h_save uint64 = h
		var t uint

		/* Operators defined in FIPS 180-2:4.1.2.  */

		var Ch func(x uint64, y uint64, z uint64) int = func(x uint64, y uint64, z uint64) int { return x&y ^ ^x&z }
		var Maj func(x uint64, y uint64, z uint64) int = func(x uint64, y uint64, z uint64) int { return x&y ^ x&z ^ y&z }
		var S0 func(x uint64) int = func(x uint64) int {
			return CYCLIC(x, 28) ^ CYCLIC(x, 34) ^ CYCLIC(x, 39)
		}
		var S1 func(x uint64) int = func(x uint64) int {
			return CYCLIC(x, 14) ^ CYCLIC(x, 18) ^ CYCLIC(x, 41)
		}
		var R0 func(x uint64) int = func(x uint64) int {
			return CYCLIC(x, 1) ^ CYCLIC(x, 8) ^ x>>7
		}
		var R1 func(x uint64) int = func(x uint64) int {
			return CYCLIC(x, 19) ^ CYCLIC(x, 61) ^ x>>6
		}

		/* It is unfortunate that C does not provide an operator for
		   cyclic rotation.  Hope the C compiler is smart enough.  */

		var CYCLIC func(w uint64, s int) int = func(w uint64, s int) int { return w>>s | w<<64 - s }

		/* Compute the message schedule according to FIPS 180-2:6.3.2 step 2.  */

		for t = 0; t < 16; t++ {
			W[t] = SWAP(*words)
			words++
		}
		for t = 16; t < 80; t++ {
			W[t] = R1(W[t-2]) + W[t-7] + R0(W[t-15]) + W[t-16]
		}

		/* The actual computation according to FIPS 180-2:6.3.2 step 3.  */

		for t = 0; t < 80; t++ {
			var T1 uint64 = h + S1(e) + Ch(e, f, g) + K[t] + W[t]
			var T2 uint64 = S0(a) + Maj(a, b, c)
			h = g
			g = f
			f = e
			e = d + T1
			d = c
			c = b
			b = a
			a = T1 + T2
		}

		/* Add the starting values of the context according to FIPS 180-2:6.3.2
		   step 4.  */

		a += a_save
		b += b_save
		c += c_save
		d += d_save
		e += e_save
		f += f_save
		g += g_save
		h += h_save

		/* Prepare for the next round.  */

		nwords -= 16

		/* Prepare for the next round.  */

	}

	/* Put checksum in context given as argument.  */

	ctx.GetH()[0] = a
	ctx.GetH()[1] = b
	ctx.GetH()[2] = c
	ctx.GetH()[3] = d
	ctx.GetH()[4] = e
	ctx.GetH()[5] = f
	ctx.GetH()[6] = g
	ctx.GetH()[7] = h
}
func Sha512InitCtx(ctx *Sha512Ctx) {
	ctx.GetH()[0] = uint64(0x6a09e667f3bcc908)
	ctx.GetH()[1] = uint64(-0x4498517a7b3558c5)
	ctx.GetH()[2] = uint64(0x3c6ef372fe94f82b)
	ctx.GetH()[3] = uint64(-0x5ab00ac5a0e2c90f)
	ctx.GetH()[4] = uint64(0x510e527fade682d1)
	ctx.GetH()[5] = uint64(-0x64fa9773d4c193e1)
	ctx.GetH()[6] = uint64(0x1f83d9abfb41bd6b)
	ctx.GetH()[7] = uint64(0x5be0cd19137e2179)
	ctx.GetTotal()[1] = 0
	ctx.GetTotal()[0] = ctx.GetTotal()[1]
	ctx.SetBuflen(0)
}
func Sha512FinishCtx(ctx *Sha512Ctx, resbuf any) any {
	/* Take yet unprocessed bytes into account.  */

	var bytes uint64 = ctx.GetBuflen()
	var pad int
	var i uint

	/* Now count remaining bytes.  */

	ctx.GetTotal()[0] += bytes
	if ctx.GetTotal()[0] < bytes {
		ctx.GetTotal()[1]++
	}
	if bytes >= 112 {
		pad = 128 + 112 - int(bytes)
	} else {
		pad = 112 - int(bytes)
	}
	memcpy(ctx.GetBuffer()[bytes], Fillbuf, pad)

	/* Put the 128-bit file length in *bits* at the end of the buffer.  */

	*((*uint64)(ctx.GetBuffer()[bytes+pad+8])) = SWAP(ctx.GetTotal()[0] << 3)
	*((*uint64)(ctx.GetBuffer()[bytes+pad])) = SWAP(ctx.GetTotal()[1]<<3 | ctx.GetTotal()[0]>>61)

	/* Process last bytes.  */

	Sha512ProcessBlock(ctx.GetBuffer(), size_t(bytes+pad+16), ctx)

	/* Put result from CTX in first 64 bytes following RESBUF.  */

	for i = 0; i < 8; i++ {
		(*uint64)(resbuf)[i] = SWAP(ctx.GetH()[i])
	}
	return resbuf
}
func Sha512ProcessBytes(buffer any, len_ int, ctx *Sha512Ctx) {
	/* When we already have some bits in our internal buffer concatenate
	   both inputs first.  */

	if ctx.GetBuflen() != 0 {
		var left_over int = int(ctx.GetBuflen())
		var add int = size_t(b.Cond(256-left_over > len_, len_, 256-left_over))
		memcpy(ctx.GetBuffer()[left_over], buffer, add)
		ctx.SetBuflen(ctx.GetBuflen() + add)
		if ctx.GetBuflen() > 128 {
			Sha512ProcessBlock(ctx.GetBuffer(), ctx.GetBuflen() & ^127, ctx)
			ctx.SetBuflen(ctx.GetBuflen() & 127)

			/* The regions in the following copy operation cannot overlap.  */

			memcpy(ctx.GetBuffer(), ctx.GetBuffer()[left_over + add & ^127], int(ctx.GetBuflen()))

			/* The regions in the following copy operation cannot overlap.  */

		}
		buffer = (*byte)(buffer + add)
		len_ -= add
	}

	/* Process available complete blocks.  */

	if len_ >= 128 {

		/* To check alignment gcc has an appropriate operator.  Other
		   compilers don't.  */

		var UNALIGNED_P func(p any) bool = func(p any) bool {
			return uintPtr(p)%b.SizeOf("uint64_t") != 0
		}
		if UNALIGNED_P(buffer) {
			for len_ > 128 {
				Sha512ProcessBlock(memcpy(ctx.GetBuffer(), buffer, 128), 128, ctx)
				buffer = (*byte)(buffer + 128)
				len_ -= 128
			}
		} else {
			Sha512ProcessBlock(buffer, len_ & ^127, ctx)
			buffer = (*byte)(buffer + (len_ & ^127))
			len_ &= 127
		}
	}

	/* Move remaining bytes into internal buffer.  */

	if len_ > 0 {
		var left_over int = int(ctx.GetBuflen())
		memcpy(ctx.GetBuffer()[left_over], buffer, len_)
		left_over += len_
		if left_over >= 128 {
			Sha512ProcessBlock(ctx.GetBuffer(), 128, ctx)
			left_over -= 128
			memcpy(ctx.GetBuffer(), ctx.GetBuffer()[128], left_over)
		}
		ctx.SetBuflen(left_over)
	}

	/* Move remaining bytes into internal buffer.  */
}
func PhpSha512CryptR(key *byte, salt *byte, buffer *byte, buflen int) *byte {
	var alt_result []uint8
	var temp_result []uint8
	var ctx Sha512Ctx
	var alt_ctx Sha512Ctx
	var salt_len int
	var key_len int
	var cnt int
	var cp *byte
	var copied_key *byte = nil
	var copied_salt *byte = nil
	var p_bytes *byte
	var s_bytes *byte

	/* Default number of rounds.  */

	var rounds int = ROUNDS_DEFAULT
	var rounds_custom types.ZendBool = 0

	/* Find beginning of salt string.  The prefix should normally always
	   be present.  Just in case it is not.  */

	if strncmp(Sha512SaltPrefix, salt, b.SizeOf("sha512_salt_prefix")-1) == 0 {

		/* Skip salt prefix.  */

		salt += b.SizeOf("sha512_salt_prefix") - 1

		/* Skip salt prefix.  */

	}
	if strncmp(salt, Sha512RoundsPrefix, b.SizeOf("sha512_rounds_prefix")-1) == 0 {
		var num *byte = salt + b.SizeOf("sha512_rounds_prefix") - 1
		var endp *byte
		var srounds zend.ZendUlong = zend.ZEND_STRTOUL(num, &endp, 10)
		if (*endp) == '$' {
			salt = endp + 1
			rounds = b.Max(ROUNDS_MIN, cli.MIN(srounds, ROUNDS_MAX))
			rounds_custom = 1
		}
	}
	salt_len = cli.MIN(strcspn(salt, "$"), SALT_LEN_MAX)
	key_len = strlen(key)
	if (key-(*byte)(0))%__alignof__(uint64) != 0 {
		var tmp *byte = (*byte)(alloca(key_len + __alignof__(uint64)))
		copied_key = memcpy(tmp+__alignof__(uint64)-(tmp-(*byte)(0))%__alignof__(uint64), key, key_len)
		key = copied_key
	}
	if (salt-(*byte)(0))%__alignof__(uint64) != 0 {
		var tmp *byte = (*byte)(alloca(salt_len + 1 + __alignof__(uint64)))
		copied_salt = memcpy(tmp+__alignof__(uint64)-(tmp-(*byte)(0))%__alignof__(uint64), salt, salt_len)
		salt = copied_salt
		copied_salt[salt_len] = 0
	}

	/* Prepare for the real work.  */

	Sha512InitCtx(&ctx)

	/* Add the key string.  */

	Sha512ProcessBytes(key, key_len, &ctx)

	/* The last part is the salt string.  This must be at most 16
	   characters and it ends at the first `$' character (for
	   compatibility with existing implementations).  */

	Sha512ProcessBytes(salt, salt_len, &ctx)

	/* Compute alternate SHA512 sum with input KEY, SALT, and KEY.  The
	   final result will be added to the first context.  */

	Sha512InitCtx(&alt_ctx)

	/* Add key.  */

	Sha512ProcessBytes(key, key_len, &alt_ctx)

	/* Add salt.  */

	Sha512ProcessBytes(salt, salt_len, &alt_ctx)

	/* Add key again.  */

	Sha512ProcessBytes(key, key_len, &alt_ctx)

	/* Now get result of this (64 bytes) and add it to the other
	   context.  */

	Sha512FinishCtx(&alt_ctx, alt_result)

	/* Add for any character in the key one byte of the alternate sum.  */

	for cnt = key_len; cnt > 64; cnt -= 64 {
		Sha512ProcessBytes(alt_result, 64, &ctx)
	}
	Sha512ProcessBytes(alt_result, cnt, &ctx)

	/* Take the binary representation of the length of the key and for every
	   1 add the alternate sum, for every 0 the key.  */

	for cnt = key_len; cnt > 0; cnt >>= 1 {
		if (cnt & 1) != 0 {
			Sha512ProcessBytes(alt_result, 64, &ctx)
		} else {
			Sha512ProcessBytes(key, key_len, &ctx)
		}
	}

	/* Create intermediate result.  */

	Sha512FinishCtx(&ctx, alt_result)

	/* Start computation of P byte sequence.  */

	Sha512InitCtx(&alt_ctx)

	/* For every character in the password add the entire password.  */

	for cnt = 0; cnt < key_len; cnt++ {
		Sha512ProcessBytes(key, key_len, &alt_ctx)
	}

	/* Finish the digest.  */

	Sha512FinishCtx(&alt_ctx, temp_result)

	/* Create byte sequence P.  */

	p_bytes = alloca(key_len)
	cp = p_bytes
	for cnt = key_len; cnt >= 64; cnt -= 64 {
		cp = __phpMempcpy(any(cp), any(temp_result), 64)
	}
	memcpy(cp, temp_result, cnt)

	/* Start computation of S byte sequence.  */

	Sha512InitCtx(&alt_ctx)

	/* For every character in the password add the entire password.  */

	for cnt = 0; cnt < size_t(16+alt_result[0]); cnt++ {
		Sha512ProcessBytes(salt, salt_len, &alt_ctx)
	}

	/* Finish the digest.  */

	Sha512FinishCtx(&alt_ctx, temp_result)

	/* Create byte sequence S.  */

	s_bytes = alloca(salt_len)
	cp = s_bytes
	for cnt = salt_len; cnt >= 64; cnt -= 64 {
		cp = __phpMempcpy(cp, temp_result, 64)
	}
	memcpy(cp, temp_result, cnt)

	/* Repeatedly run the collected hash value through SHA512 to burn
	   CPU cycles.  */

	for cnt = 0; cnt < rounds; cnt++ {

		/* New context.  */

		Sha512InitCtx(&ctx)

		/* Add key or last result.  */

		if (cnt & 1) != 0 {
			Sha512ProcessBytes(p_bytes, key_len, &ctx)
		} else {
			Sha512ProcessBytes(alt_result, 64, &ctx)
		}

		/* Add salt for numbers not divisible by 3.  */

		if cnt%3 != 0 {
			Sha512ProcessBytes(s_bytes, salt_len, &ctx)
		}

		/* Add key for numbers not divisible by 7.  */

		if cnt%7 != 0 {
			Sha512ProcessBytes(p_bytes, key_len, &ctx)
		}

		/* Add key or last result.  */

		if (cnt & 1) != 0 {
			Sha512ProcessBytes(alt_result, 64, &ctx)
		} else {
			Sha512ProcessBytes(p_bytes, key_len, &ctx)
		}

		/* Create intermediate result.  */

		Sha512FinishCtx(&ctx, alt_result)

		/* Create intermediate result.  */

	}

	/* Now we can construct the result string.  It consists of three
	   parts.  */

	cp = __phpStpncpy(buffer, Sha512SaltPrefix, b.Max(0, buflen))
	buflen -= b.SizeOf("sha512_salt_prefix") - 1
	if rounds_custom != 0 {
		var n int = core.Snprintf(cp, b.Max(0, buflen), "%s%zu$", Sha512RoundsPrefix, rounds)
		cp += n
		buflen -= n
	}
	cp = __phpStpncpy(cp, salt, cli.MIN(int(b.Max(0, buflen)), salt_len))
	buflen -= int(cli.MIN(int(b.Max(0, buflen)), salt_len))
	if buflen > 0 {
		b.PostInc(&(*cp)) = '$'
		buflen--
	}
	var b64_from_24bit func(B2 __auto__, B1 __auto__, B0 uint8, N int) = func(B2 __auto__, B1 __auto__, B0 uint8, N int) {
		var w uint = B2<<16 | B1<<8 | B0
		var n int = N
		for b.PostDec(&n) > 0 && buflen > 0 {
			b.PostInc(&(*cp)) = B64t[w&0x3f]
			buflen--
			w >>= 6
		}
	}
	b64_from_24bit(alt_result[0], alt_result[21], alt_result[42], 4)
	b64_from_24bit(alt_result[22], alt_result[43], alt_result[1], 4)
	b64_from_24bit(alt_result[44], alt_result[2], alt_result[23], 4)
	b64_from_24bit(alt_result[3], alt_result[24], alt_result[45], 4)
	b64_from_24bit(alt_result[25], alt_result[46], alt_result[4], 4)
	b64_from_24bit(alt_result[47], alt_result[5], alt_result[26], 4)
	b64_from_24bit(alt_result[6], alt_result[27], alt_result[48], 4)
	b64_from_24bit(alt_result[28], alt_result[49], alt_result[7], 4)
	b64_from_24bit(alt_result[50], alt_result[8], alt_result[29], 4)
	b64_from_24bit(alt_result[9], alt_result[30], alt_result[51], 4)
	b64_from_24bit(alt_result[31], alt_result[52], alt_result[10], 4)
	b64_from_24bit(alt_result[53], alt_result[11], alt_result[32], 4)
	b64_from_24bit(alt_result[12], alt_result[33], alt_result[54], 4)
	b64_from_24bit(alt_result[34], alt_result[55], alt_result[13], 4)
	b64_from_24bit(alt_result[56], alt_result[14], alt_result[35], 4)
	b64_from_24bit(alt_result[15], alt_result[36], alt_result[57], 4)
	b64_from_24bit(alt_result[37], alt_result[58], alt_result[16], 4)
	b64_from_24bit(alt_result[59], alt_result[17], alt_result[38], 4)
	b64_from_24bit(alt_result[18], alt_result[39], alt_result[60], 4)
	b64_from_24bit(alt_result[40], alt_result[61], alt_result[19], 4)
	b64_from_24bit(alt_result[62], alt_result[20], alt_result[41], 4)
	b64_from_24bit(0, 0, alt_result[63], 2)
	if buflen <= 0 {
		errno = ERANGE
		buffer = nil
	} else {
		*cp = '0'
	}

	/* Clear the buffer for the intermediate result so that people
	   attaching to processes or reading core dumps cannot get any
	   information.  We do it in this way to clear correct_words[]
	   inside the SHA512 implementation as well.  */

	Sha512InitCtx(&ctx)
	Sha512FinishCtx(&ctx, alt_result)
	zend.ZEND_SECURE_ZERO(temp_result, b.SizeOf("temp_result"))
	zend.ZEND_SECURE_ZERO(p_bytes, key_len)
	zend.ZEND_SECURE_ZERO(s_bytes, salt_len)
	zend.ZEND_SECURE_ZERO(&ctx, b.SizeOf("ctx"))
	zend.ZEND_SECURE_ZERO(&alt_ctx, b.SizeOf("alt_ctx"))
	if copied_key != nil {
		zend.ZEND_SECURE_ZERO(copied_key, key_len)
	}
	if copied_salt != nil {
		zend.ZEND_SECURE_ZERO(copied_salt, salt_len)
	}
	return buffer
}
func PhpSha512Crypt(key *byte, salt *byte) *byte {
	/* We don't want to have an arbitrary limit in the size of the
	   password.  We can compute an upper bound for the size of the
	   result in advance and so we can prepare the buffer we pass to
	   `sha512_crypt_r'.  */

	var buffer *byte
	var buflen int = 0
	var needed int = int(b.SizeOf("sha512_salt_prefix") - 1 + b.SizeOf("sha512_rounds_prefix") + 9 + 1 + strlen(salt) + 1 + 86 + 1)
	if buflen < needed {
		var new_buffer *byte = (*byte)(realloc(buffer, needed))
		if new_buffer == nil {
			return nil
		}
		buffer = new_buffer
		buflen = needed
	}
	return PhpSha512CryptR(key, salt, buffer, buflen)
}
