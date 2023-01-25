// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/sapi/cli"
	"sik/zend"
)

// Source: <ext/standard/crypt_sha256.c>

/* SHA256-based Unix crypt implementation.
   Released into the Public Domain by Ulrich Drepper <drepper@redhat.com>.  */

// # include "php.h"

// # include "php_main.h"

// # include < errno . h >

// # include < limits . h >

// # include < stdio . h >

// # include < stdlib . h >

// # include < sys / param . h >

// # include < sys / types . h >

// # include < string . h >

func __phpStpncpy(dst *byte, src *byte, len_ int) *byte {
	var n int = strlen(src)
	if n > len_ {
		n = len_
	}
	return strncpy(dst, src, len_) + n
}
func __phpMempcpy(dst any, src any, len_ int) any {
	return (*byte)(memcpy(dst, src, len_)) + len_
}

/* Structure to save state of computation between the single steps.  */

func SWAP(n __auto__) int {
	return n<<24 | (n&0xff00)<<8 | n>>8&0xff00 | n>>24
}

/* This array contains the bytes used to pad the buffer to the next
   64-byte boundary.  (FIPS 180-2:5.1.1)  */

var Fillbuf []uint8 = []uint8{0x80, 0}

/* Constants for SHA256 from FIPS 180-2:4.2.2.  */

var K32 []uint32 = []uint32{0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5, 0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174, 0xe49b69c1, 0xefbe4786, 0xfc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da, 0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x6ca6351, 0x14292967, 0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85, 0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070, 0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3, 0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2}

/* Process LEN bytes of BUFFER, accumulating context into CTX.
   It is assumed that LEN % 64 == 0.  */

func Sha256ProcessBlock(buffer any, len_ int, ctx *Sha256Ctx) {
	var words *uint32 = buffer
	var nwords int = len_ / b.SizeOf("uint32_t")
	var t uint
	var a uint32 = ctx.GetH()[0]
	var b uint32 = ctx.GetH()[1]
	var c uint32 = ctx.GetH()[2]
	var d uint32 = ctx.GetH()[3]
	var e uint32 = ctx.GetH()[4]
	var f uint32 = ctx.GetH()[5]
	var g uint32 = ctx.GetH()[6]
	var h uint32 = ctx.GetH()[7]

	/* First increment the byte count.  FIPS 180-2 specifies the possible
	   length of the file up to 2^64 bits.  Here we only compute the
	   number of bytes.  Do a double word increment.  */

	ctx.GetTotal()[0] += uint32(len_)
	if ctx.GetTotal()[0] < len_ {
		ctx.GetTotal()[1]++
	}

	/* Process all bytes in the buffer with 64 bytes in each round of
	   the loop.  */

	for nwords > 0 {
		var W []uint32
		var a_save uint32 = a
		var b_save uint32 = b
		var c_save uint32 = c
		var d_save uint32 = d
		var e_save uint32 = e
		var f_save uint32 = f
		var g_save uint32 = g
		var h_save uint32 = h

		/* Operators defined in FIPS 180-2:4.1.2.  */

		var Ch func(x uint32, y uint32, z uint32) int = func(x uint32, y uint32, z uint32) int { return x&y ^ ^x&z }
		var Maj func(x uint32, y uint32, z uint32) int = func(x uint32, y uint32, z uint32) int { return x&y ^ x&z ^ y&z }
		var S0 func(x uint32) int = func(x uint32) int {
			return CYCLIC(x, 2) ^ CYCLIC(x, 13) ^ CYCLIC(x, 22)
		}
		var S1 func(x uint32) int = func(x uint32) int {
			return CYCLIC(x, 6) ^ CYCLIC(x, 11) ^ CYCLIC(x, 25)
		}
		var R0 func(x uint32) int = func(x uint32) int {
			return CYCLIC(x, 7) ^ CYCLIC(x, 18) ^ x>>3
		}
		var R1 func(x uint32) int = func(x uint32) int {
			return CYCLIC(x, 17) ^ CYCLIC(x, 19) ^ x>>10
		}

		/* It is unfortunate that C does not provide an operator for
		   cyclic rotation.  Hope the C compiler is smart enough.  */

		var CYCLIC func(w uint32, s int) int = func(w uint32, s int) int { return w>>s | w<<32 - s }

		/* Compute the message schedule according to FIPS 180-2:6.2.2 step 2.  */

		for t = 0; t < 16; t++ {
			W[t] = SWAP(*words)
			words++
		}
		for t = 16; t < 64; t++ {
			W[t] = R1(W[t-2]) + W[t-7] + R0(W[t-15]) + W[t-16]
		}

		/* The actual computation according to FIPS 180-2:6.2.2 step 3.  */

		for t = 0; t < 64; t++ {
			var T1 uint32 = h + S1(e) + Ch(e, f, g) + K[t] + W[t]
			var T2 uint32 = S0(a) + Maj(a, b, c)
			h = g
			g = f
			f = e
			e = d + T1
			d = c
			c = b
			b = a
			a = T1 + T2
		}

		/* Add the starting values of the context according to FIPS 180-2:6.2.2
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

/* Initialize structure containing state of computation.
   (FIPS 180-2:5.3.2)  */

func Sha256InitCtx(ctx *Sha256Ctx) {
	ctx.GetH()[0] = 0x6a09e667
	ctx.GetH()[1] = 0xbb67ae85
	ctx.GetH()[2] = 0x3c6ef372
	ctx.GetH()[3] = 0xa54ff53a
	ctx.GetH()[4] = 0x510e527f
	ctx.GetH()[5] = 0x9b05688c
	ctx.GetH()[6] = 0x1f83d9ab
	ctx.GetH()[7] = 0x5be0cd19
	ctx.GetTotal()[1] = 0
	ctx.GetTotal()[0] = ctx.GetTotal()[1]
	ctx.SetBuflen(0)
}

/* Process the remaining bytes in the internal buffer and the usual
   prolog according to the standard and write the result to RESBUF.

   IMPORTANT: On some systems it is required that RESBUF is correctly
   aligned for a 32 bits value.  */

func Sha256FinishCtx(ctx *Sha256Ctx, resbuf any) any {
	/* Take yet unprocessed bytes into account.  */

	var bytes uint32 = ctx.GetBuflen()
	var pad int
	var i uint

	/* Now count remaining bytes.  */

	ctx.GetTotal()[0] += bytes
	if ctx.GetTotal()[0] < bytes {
		ctx.GetTotal()[1]++
	}
	if bytes >= 56 {
		pad = 64 + 56 - bytes
	} else {
		pad = 56 - bytes
	}
	memcpy(&ctx.buffer[bytes], Fillbuf, pad)

	/* Put the 64-bit file length in *bits* at the end of the buffer.  */

	*((*uint32)(&ctx.buffer[bytes+pad+4])) = SWAP(ctx.GetTotal()[0] << 3)
	*((*uint32)(&ctx.buffer[bytes+pad])) = SWAP(ctx.GetTotal()[1]<<3 | ctx.GetTotal()[0]>>29)

	/* Process last bytes.  */

	Sha256ProcessBlock(ctx.GetBuffer(), bytes+pad+8, ctx)

	/* Put result from CTX in first 32 bytes following RESBUF.  */

	for i = 0; i < 8; i++ {
		(*uint32)(resbuf)[i] = SWAP(ctx.GetH()[i])
	}
	return resbuf
}
func Sha256ProcessBytes(buffer any, len_ int, ctx *Sha256Ctx) {
	/* When we already have some bits in our internal buffer concatenate
	   both inputs first.  */

	if ctx.GetBuflen() != 0 {
		var left_over int = ctx.GetBuflen()
		var add int = b.Cond(128-left_over > len_, len_, 128-left_over)
		memcpy(&ctx.buffer[left_over], buffer, add)
		ctx.SetBuflen(ctx.GetBuflen() + uint32(add))
		if ctx.GetBuflen() > 64 {
			Sha256ProcessBlock(ctx.GetBuffer(), ctx.GetBuflen() & ^63, ctx)
			ctx.SetBuflen(ctx.GetBuflen() & 63)

			/* The regions in the following copy operation cannot overlap.  */

			memcpy(ctx.GetBuffer(), &ctx.buffer[left_over + add & ^63], ctx.GetBuflen())

			/* The regions in the following copy operation cannot overlap.  */

		}
		buffer = (*byte)(buffer + add)
		len_ -= add
	}

	/* Process available complete blocks.  */

	if len_ >= 64 {

		/* To check alignment gcc has an appropriate operator.  Other
		   compilers don't.  */

		var UNALIGNED_P func(p any) bool = func(p any) bool {
			return uintPtr(p)%b.SizeOf("uint32_t") != 0
		}
		if UNALIGNED_P(buffer) {
			for len_ > 64 {
				Sha256ProcessBlock(memcpy(ctx.GetBuffer(), buffer, 64), 64, ctx)
				buffer = (*byte)(buffer + 64)
				len_ -= 64
			}
		} else {
			Sha256ProcessBlock(buffer, len_ & ^63, ctx)
			buffer = (*byte)(buffer + (len_ & ^63))
			len_ &= 63
		}
	}

	/* Move remaining bytes into internal buffer.  */

	if len_ > 0 {
		var left_over int = ctx.GetBuflen()
		memcpy(&ctx.buffer[left_over], buffer, len_)
		left_over += len_
		if left_over >= 64 {
			Sha256ProcessBlock(ctx.GetBuffer(), 64, ctx)
			left_over -= 64
			memcpy(ctx.GetBuffer(), &ctx.buffer[64], left_over)
		}
		ctx.SetBuflen(uint32(left_over))
	}

	/* Move remaining bytes into internal buffer.  */
}

/* Define our magic string to mark salt for SHA256 "encryption"
   replacement.  */

var Sha256SaltPrefix []byte = "$5$"

/* Prefix for optional rounds specification.  */

var Sha256RoundsPrefix []byte = "rounds="

/* Maximum salt string length.  */

const SALT_LEN_MAX = 16

/* Default number of rounds if not explicitly specified.  */

const ROUNDS_DEFAULT = 5000

/* Minimum number of rounds.  */

const ROUNDS_MIN = 1000

/* Maximum number of rounds.  */

const ROUNDS_MAX = 999999999

/* Table with characters for base64 transformation.  */

var B64t []byte = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func PhpSha256CryptR(key *byte, salt *byte, buffer *byte, buflen int) *byte {
	var alt_result []uint8
	var temp_result []uint8
	var ctx Sha256Ctx
	var alt_ctx Sha256Ctx
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
	var rounds_custom zend.ZendBool = 0

	/* Find beginning of salt string.  The prefix should normally always
	   be present.  Just in case it is not.  */

	if strncmp(Sha256SaltPrefix, salt, b.SizeOf("sha256_salt_prefix")-1) == 0 {

		/* Skip salt prefix.  */

		salt += b.SizeOf("sha256_salt_prefix") - 1

		/* Skip salt prefix.  */

	}
	if strncmp(salt, Sha256RoundsPrefix, b.SizeOf("sha256_rounds_prefix")-1) == 0 {
		var num *byte = salt + b.SizeOf("sha256_rounds_prefix") - 1
		var endp *byte
		var srounds zend.ZendUlong = zend.ZEND_STRTOUL(num, &endp, 10)
		if (*endp) == '$' {
			salt = endp + 1
			rounds = zend.MAX(ROUNDS_MIN, cli.MIN(srounds, ROUNDS_MAX))
			rounds_custom = 1
		}
	}
	salt_len = cli.MIN(strcspn(salt, "$"), SALT_LEN_MAX)
	key_len = strlen(key)
	if (key-(*byte)(0))%__alignof__(uint32_t) != 0 {
		var tmp *byte = (*byte)(alloca(key_len + __alignof__(uint32_t)))
		copied_key = memcpy(tmp+__alignof__(uint32_t)-(tmp-(*byte)(0))%__alignof__(uint32_t), key, key_len)
		key = copied_key
	}
	if (salt-(*byte)(0))%__alignof__(uint32_t) != 0 {
		var tmp *byte = (*byte)(alloca(salt_len + 1 + __alignof__(uint32_t)))
		copied_salt = memcpy(tmp+__alignof__(uint32_t)-(tmp-(*byte)(0))%__alignof__(uint32_t), salt, salt_len)
		salt = copied_salt
		copied_salt[salt_len] = 0
	}

	/* Prepare for the real work.  */

	Sha256InitCtx(&ctx)

	/* Add the key string.  */

	Sha256ProcessBytes(key, key_len, &ctx)

	/* The last part is the salt string.  This must be at most 16
	   characters and it ends at the first `$' character (for
	   compatibility with existing implementations).  */

	Sha256ProcessBytes(salt, salt_len, &ctx)

	/* Compute alternate SHA256 sum with input KEY, SALT, and KEY.  The
	   final result will be added to the first context.  */

	Sha256InitCtx(&alt_ctx)

	/* Add key.  */

	Sha256ProcessBytes(key, key_len, &alt_ctx)

	/* Add salt.  */

	Sha256ProcessBytes(salt, salt_len, &alt_ctx)

	/* Add key again.  */

	Sha256ProcessBytes(key, key_len, &alt_ctx)

	/* Now get result of this (32 bytes) and add it to the other
	   context.  */

	Sha256FinishCtx(&alt_ctx, alt_result)

	/* Add for any character in the key one byte of the alternate sum.  */

	for cnt = key_len; cnt > 32; cnt -= 32 {
		Sha256ProcessBytes(alt_result, 32, &ctx)
	}
	Sha256ProcessBytes(alt_result, cnt, &ctx)

	/* Take the binary representation of the length of the key and for every
	   1 add the alternate sum, for every 0 the key.  */

	for cnt = key_len; cnt > 0; cnt >>= 1 {
		if (cnt & 1) != 0 {
			Sha256ProcessBytes(alt_result, 32, &ctx)
		} else {
			Sha256ProcessBytes(key, key_len, &ctx)
		}
	}

	/* Create intermediate result.  */

	Sha256FinishCtx(&ctx, alt_result)

	/* Start computation of P byte sequence.  */

	Sha256InitCtx(&alt_ctx)

	/* For every character in the password add the entire password.  */

	for cnt = 0; cnt < key_len; cnt++ {
		Sha256ProcessBytes(key, key_len, &alt_ctx)
	}

	/* Finish the digest.  */

	Sha256FinishCtx(&alt_ctx, temp_result)

	/* Create byte sequence P.  */

	p_bytes = alloca(key_len)
	cp = p_bytes
	for cnt = key_len; cnt >= 32; cnt -= 32 {
		cp = __phpMempcpy(any(cp), any(temp_result), 32)
	}
	memcpy(cp, temp_result, cnt)

	/* Start computation of S byte sequence.  */

	Sha256InitCtx(&alt_ctx)

	/* For every character in the password add the entire password.  */

	for cnt = 0; cnt < size_t(16+alt_result[0]); cnt++ {
		Sha256ProcessBytes(salt, salt_len, &alt_ctx)
	}

	/* Finish the digest.  */

	Sha256FinishCtx(&alt_ctx, temp_result)

	/* Create byte sequence S.  */

	s_bytes = alloca(salt_len)
	cp = s_bytes
	for cnt = salt_len; cnt >= 32; cnt -= 32 {
		cp = __phpMempcpy(cp, temp_result, 32)
	}
	memcpy(cp, temp_result, cnt)

	/* Repeatedly run the collected hash value through SHA256 to burn
	   CPU cycles.  */

	for cnt = 0; cnt < rounds; cnt++ {

		/* New context.  */

		Sha256InitCtx(&ctx)

		/* Add key or last result.  */

		if (cnt & 1) != 0 {
			Sha256ProcessBytes(p_bytes, key_len, &ctx)
		} else {
			Sha256ProcessBytes(alt_result, 32, &ctx)
		}

		/* Add salt for numbers not divisible by 3.  */

		if cnt%3 != 0 {
			Sha256ProcessBytes(s_bytes, salt_len, &ctx)
		}

		/* Add key for numbers not divisible by 7.  */

		if cnt%7 != 0 {
			Sha256ProcessBytes(p_bytes, key_len, &ctx)
		}

		/* Add key or last result.  */

		if (cnt & 1) != 0 {
			Sha256ProcessBytes(alt_result, 32, &ctx)
		} else {
			Sha256ProcessBytes(p_bytes, key_len, &ctx)
		}

		/* Create intermediate result.  */

		Sha256FinishCtx(&ctx, alt_result)

		/* Create intermediate result.  */

	}

	/* Now we can construct the result string.  It consists of three
	   parts.  */

	cp = __phpStpncpy(buffer, Sha256SaltPrefix, zend.MAX(0, buflen))
	buflen -= b.SizeOf("sha256_salt_prefix") - 1
	if rounds_custom != 0 {
		var n int = core.Snprintf(cp, zend.MAX(0, buflen), "%s%zu$", Sha256RoundsPrefix, rounds)
		cp += n
		buflen -= n
	}
	cp = __phpStpncpy(cp, salt, cli.MIN(int(zend.MAX(0, buflen)), salt_len))
	buflen -= cli.MIN(zend.MAX(0, buflen), int(salt_len))
	if buflen > 0 {
		b.PostInc(&(*cp)) = '$'
		buflen--
	}
	var b64_from_24bit func(B2 __auto__, B1 uint8, B0 uint8, N int) = func(B2 __auto__, B1 uint8, B0 uint8, N int) {
		var w uint = B2<<16 | B1<<8 | B0
		var n int = N
		for b.PostDec(&n) > 0 && buflen > 0 {
			b.PostInc(&(*cp)) = B64t[w&0x3f]
			buflen--
			w >>= 6
		}
	}
	b64_from_24bit(alt_result[0], alt_result[10], alt_result[20], 4)
	b64_from_24bit(alt_result[21], alt_result[1], alt_result[11], 4)
	b64_from_24bit(alt_result[12], alt_result[22], alt_result[2], 4)
	b64_from_24bit(alt_result[3], alt_result[13], alt_result[23], 4)
	b64_from_24bit(alt_result[24], alt_result[4], alt_result[14], 4)
	b64_from_24bit(alt_result[15], alt_result[25], alt_result[5], 4)
	b64_from_24bit(alt_result[6], alt_result[16], alt_result[26], 4)
	b64_from_24bit(alt_result[27], alt_result[7], alt_result[17], 4)
	b64_from_24bit(alt_result[18], alt_result[28], alt_result[8], 4)
	b64_from_24bit(alt_result[9], alt_result[19], alt_result[29], 4)
	b64_from_24bit(0, alt_result[31], alt_result[30], 3)
	if buflen <= 0 {
		errno = ERANGE
		buffer = nil
	} else {
		*cp = '0'
	}

	/* Clear the buffer for the intermediate result so that people
	   attaching to processes or reading core dumps cannot get any
	   information.  We do it in this way to clear correct_words[]
	   inside the SHA256 implementation as well.  */

	Sha256InitCtx(&ctx)
	Sha256FinishCtx(&ctx, alt_result)
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

/* This entry point is equivalent to the `crypt' function in Unix
   libcs.  */

func PhpSha256Crypt(key *byte, salt *byte) *byte {
	/* We don't want to have an arbitrary limit in the size of the
	   password.  We can compute an upper bound for the size of the
	   result in advance and so we can prepare the buffer we pass to
	   `sha256_crypt_r'.  */

	var buffer *byte
	var buflen int = 0
	var needed int = b.SizeOf("sha256_salt_prefix") - 1 + b.SizeOf("sha256_rounds_prefix") + 9 + 1 + int(strlen(salt)+1+43+1)
	if buflen < needed {
		var new_buffer *byte = (*byte)(realloc(buffer, needed))
		if new_buffer == nil {
			return nil
		}
		buffer = new_buffer
		buflen = needed
	}
	return PhpSha256CryptR(key, salt, buffer, buflen)
}
