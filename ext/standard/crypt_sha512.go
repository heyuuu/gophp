// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/crypt_sha512.c>

/* SHA512-based Unix crypt implementation.
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

/* See #51582 */

/* Structure to save state of computation between the single steps.  */

// @type Sha512Ctx struct

// #define SWAP(n) ( ( ( n ) << 56 ) | ( ( ( n ) & 0xff00 ) << 40 ) | ( ( ( n ) & 0xff0000 ) << 24 ) | ( ( ( n ) & 0xff000000 ) << 8 ) | ( ( ( n ) >> 8 ) & 0xff000000 ) | ( ( ( n ) >> 24 ) & 0xff0000 ) | ( ( ( n ) >> 40 ) & 0xff00 ) | ( ( n ) >> 56 ) )

/* This array contains the bytes used to pad the buffer to the next
   64-byte boundary.  (FIPS 180-2:5.1.2)  */

/* Constants for SHA512 from FIPS 180-2:4.2.3.  */

var K64 []uint64 = []uint64{uint64(0x428a2f98d728ae22), uint64(0x7137449123ef65cd), uint64(-0x4a3f043013b2c4d1), uint64(-0x164a245a7e762444), uint64(0x3956c25bf348b538), uint64(0x59f111f1b605d019), uint64(-0x6dc07d5b50e6b065), uint64(-0x54e3a12a25927ee8), uint64(-0x27f855675cfcfdbe), uint64(0x12835b0145706fbe), uint64(0x243185be4ee4b28c), uint64(0x550c7dc3d5ffb4e2), uint64(0x72be5d74f27b896f), uint64(-0x7f214e01c4e9694f), uint64(-0x6423f958da38edcb), uint64(-0x3e640e8b3096d96c), uint64(-0x1b64963e610eb52e), uint64(-0x1041b879c7b0da1d), uint64(0xfc19dc68b8cd5b5), uint64(0x240ca1cc77ac9c65), uint64(0x2de92c6f592b0275), uint64(0x4a7484aa6ea6e483), uint64(0x5cb0a9dcbd41fbd4), uint64(0x76f988da831153b5), uint64(-0x67c1aead11992055), uint64(-0x57ce3992d24bcdf0), uint64(-0x4ffcd8376704dec1), uint64(-0x40a680384110f11c), uint64(-0x391ff40cc257703e), uint64(-0x2a586eb86cf558db), uint64(0x6ca6351e003826f), uint64(0x142929670a0e6e70), uint64(0x27b70a8546d22ffc), uint64(0x2e1b21385c26c926), uint64(0x4d2c6dfc5ac42aed), uint64(0x53380d139d95b3df), uint64(0x650a73548baf63de), uint64(0x766a0abb3c77b2a8), uint64(-0x7e3d36d1b812511a), uint64(-0x6d8dd37aeb7dcac5), uint64(-0x5d40175eb30efc9c), uint64(-0x57e599b443bdcfff), uint64(-0x3db4748f2f07686f), uint64(-0x3893ae5cf9ab41d0), uint64(-0x2e6d17e62910ade8), uint64(-0x2966f9dbaa9a56f0), uint64(-0xbf1ca7aa88edfd6), uint64(0x106aa07032bbd1b8), uint64(0x19a4c116b8d2d0c8), uint64(0x1e376c085141ab53), uint64(0x2748774cdf8eeb99), uint64(0x34b0bcb5e19b48a8), uint64(0x391c0cb3c5c95a63), uint64(0x4ed8aa4ae3418acb), uint64(0x5b9cca4f7763e373), uint64(0x682e6ff3d6b2b8a3), uint64(0x748f82ee5defb2fc), uint64(0x78a5636f43172f60), uint64(-0x7b3787eb5e0f548e), uint64(-0x7338fdf7e59bc614), uint64(-0x6f410005dc9ce1d8), uint64(-0x5baf9314217d4217), uint64(-0x41065c084d3986eb), uint64(-0x398e870d1c8dacd5), uint64(-0x35d8c13115d99e64), uint64(-0x2e794738de3f3df9), uint64(-0x15258229321f14e2), uint64(-0xa82b08011912e88), uint64(0x6f067aa72176fba), uint64(0xa637dc5a2c898a6), uint64(0x113f9804bef90dae), uint64(0x1b710b35131c471b), uint64(0x28db77f523047d84), uint64(0x32caab7b40c72493), uint64(0x3c9ebe0a15c9bebc), uint64(0x431d67c49c100d4c), uint64(0x4cc5d4becb3e42b6), uint64(0x597f299cfc657e2a), uint64(0x5fcb6fab3ad6faec), uint64(0x6c44198c4a475817)}

/* Process LEN bytes of BUFFER, accumulating context into CTX.
   It is assumed that LEN % 128 == 0.  */

func Sha512ProcessBlock(buffer any, len_ int, ctx *Sha512Ctx) {
	var words *uint64 = buffer
	var nwords int = len_ / g.SizeOf("uint64_t")
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

		// #define Ch(x,y,z) ( ( x & y ) ^ ( ~ x & z ) )

		// #define Maj(x,y,z) ( ( x & y ) ^ ( x & z ) ^ ( y & z ) )

		// #define S0(x) ( CYCLIC ( x , 28 ) ^ CYCLIC ( x , 34 ) ^ CYCLIC ( x , 39 ) )

		// #define S1(x) ( CYCLIC ( x , 14 ) ^ CYCLIC ( x , 18 ) ^ CYCLIC ( x , 41 ) )

		// #define R0(x) ( CYCLIC ( x , 1 ) ^ CYCLIC ( x , 8 ) ^ ( x >> 7 ) )

		// #define R1(x) ( CYCLIC ( x , 19 ) ^ CYCLIC ( x , 61 ) ^ ( x >> 6 ) )

		/* It is unfortunate that C does not provide an operator for
		   cyclic rotation.  Hope the C compiler is smart enough.  */

		// #define CYCLIC(w,s) ( ( w >> s ) | ( w << ( 64 - s ) ) )

		/* Compute the message schedule according to FIPS 180-2:6.3.2 step 2.  */

		for t = 0; t < 16; t++ {
			W[t] = (*words)<<56 | ((*words)&0xff00)<<40 | ((*words)&0xff0000)<<24 | ((*words)&0xff000000)<<8 | (*words)>>8&0xff000000 | (*words)>>24&0xff0000 | (*words)>>40&0xff00 | (*words)>>56
			words++
		}
		for t = 16; t < 80; t++ {
			W[t] = ((W[t-2]>>19 | W[t-2]<<64 - 19) ^ (W[t-2]>>61 | W[t-2]<<64 - 61) ^ W[t-2]>>6) + W[t-7] + ((W[t-15]>>1 | W[t-15]<<64 - 1) ^ (W[t-15]>>8 | W[t-15]<<64 - 8) ^ W[t-15]>>7) + W[t-16]
		}

		/* The actual computation according to FIPS 180-2:6.3.2 step 3.  */

		for t = 0; t < 80; t++ {
			var T1 uint64 = h + ((e>>14 | e<<64 - 14) ^ (e>>18 | e<<64 - 18) ^ (e>>41 | e<<64 - 41)) + (e&f ^ ^e&g) + K[t] + W[t]
			var T2 uint64 = ((a>>28 | a<<64 - 28) ^ (a>>34 | a<<64 - 34) ^ (a>>39 | a<<64 - 39)) + (a&b ^ a&c ^ b&c)
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

/* Initialize structure containing state of computation.
   (FIPS 180-2:5.3.3)  */

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

/* Process the remaining bytes in the internal buffer and the usual
   prolog according to the standard and write the result to RESBUF.

   IMPORTANT: On some systems it is required that RESBUF is correctly
   aligned for a 32 bits value. */

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
	memcpy(&ctx.buffer[bytes], Fillbuf, pad)

	/* Put the 128-bit file length in *bits* at the end of the buffer.  */

	*((*uint64)(&ctx.buffer[bytes+pad+8])) = ctx.GetTotal()[0]<<3<<56 | (ctx.GetTotal()[0]<<3&0xff00)<<40 | (ctx.GetTotal()[0]<<3&0xff0000)<<24 | (ctx.GetTotal()[0]<<3&0xff000000)<<8 | ctx.GetTotal()[0]<<3>>8&0xff000000 | ctx.GetTotal()[0]<<3>>24&0xff0000 | ctx.GetTotal()[0]<<3>>40&0xff00 | ctx.GetTotal()[0]<<3>>56
	*((*uint64)(&ctx.buffer[bytes+pad])) = (ctx.GetTotal()[1]<<3|ctx.GetTotal()[0]>>61)<<56 | ((ctx.GetTotal()[1]<<3|ctx.GetTotal()[0]>>61)&0xff00)<<40 | ((ctx.GetTotal()[1]<<3|ctx.GetTotal()[0]>>61)&0xff0000)<<24 | ((ctx.GetTotal()[1]<<3|ctx.GetTotal()[0]>>61)&0xff000000)<<8 | (ctx.GetTotal()[1]<<3|ctx.GetTotal()[0]>>61)>>8&0xff000000 | (ctx.GetTotal()[1]<<3|ctx.GetTotal()[0]>>61)>>24&0xff0000 | (ctx.GetTotal()[1]<<3|ctx.GetTotal()[0]>>61)>>40&0xff00 | (ctx.GetTotal()[1]<<3|ctx.GetTotal()[0]>>61)>>56

	/* Process last bytes.  */

	Sha512ProcessBlock(ctx.GetBuffer(), size_t(bytes+pad+16), ctx)

	/* Put result from CTX in first 64 bytes following RESBUF.  */

	for i = 0; i < 8; i++ {
		(*uint64)(resbuf)[i] = ctx.GetH()[i]<<56 | (ctx.GetH()[i]&0xff00)<<40 | (ctx.GetH()[i]&0xff0000)<<24 | (ctx.GetH()[i]&0xff000000)<<8 | ctx.GetH()[i]>>8&0xff000000 | ctx.GetH()[i]>>24&0xff0000 | ctx.GetH()[i]>>40&0xff00 | ctx.GetH()[i]>>56
	}
	return resbuf
}
func Sha512ProcessBytes(buffer any, len_ int, ctx *Sha512Ctx) {
	/* When we already have some bits in our internal buffer concatenate
	   both inputs first.  */

	if ctx.GetBuflen() != 0 {
		var left_over int = int(ctx.GetBuflen())
		var add int = size_t(g.Cond(256-left_over > len_, len_, 256-left_over))
		memcpy(&ctx.buffer[left_over], buffer, add)
		ctx.SetBuflen(ctx.GetBuflen() + add)
		if ctx.GetBuflen() > 128 {
			Sha512ProcessBlock(ctx.GetBuffer(), ctx.GetBuflen() & ^127, ctx)
			ctx.SetBuflen(ctx.GetBuflen() & 127)

			/* The regions in the following copy operation cannot overlap.  */

			memcpy(ctx.GetBuffer(), &ctx.buffer[left_over + add & ^127], int(ctx.GetBuflen()))

			/* The regions in the following copy operation cannot overlap.  */

		}
		buffer = (*byte)(buffer + add)
		len_ -= add
	}

	/* Process available complete blocks.  */

	if len_ >= 128 {

		/* To check alignment gcc has an appropriate operator.  Other
		   compilers don't.  */

		// #define UNALIGNED_P(p) ( ( ( uintptr_t ) p ) % sizeof ( uint64_t ) != 0 )

		if uintPtr(buffer)%g.SizeOf("uint64_t") != 0 {
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

		/* To check alignment gcc has an appropriate operator.  Other
		   compilers don't.  */

		// #define UNALIGNED_P(p) ( ( ( uintptr_t ) p ) % sizeof ( uint64_t ) != 0 )

	}

	/* Move remaining bytes into internal buffer.  */

	if len_ > 0 {
		var left_over int = int(ctx.GetBuflen())
		memcpy(&ctx.buffer[left_over], buffer, len_)
		left_over += len_
		if left_over >= 128 {
			Sha512ProcessBlock(ctx.GetBuffer(), 128, ctx)
			left_over -= 128
			memcpy(ctx.GetBuffer(), &ctx.buffer[128], left_over)
		}
		ctx.SetBuflen(left_over)
	}

	/* Move remaining bytes into internal buffer.  */
}

/* Define our magic string to mark salt for SHA512 "encryption"
   replacement.  */

var Sha512SaltPrefix []byte = "$6$"

/* Prefix for optional rounds specification.  */

var Sha512RoundsPrefix []byte = "rounds="

/* Maximum salt string length.  */

// #define SALT_LEN_MAX       16

/* Default number of rounds if not explicitly specified.  */

// #define ROUNDS_DEFAULT       5000

/* Minimum number of rounds.  */

// #define ROUNDS_MIN       1000

/* Maximum number of rounds.  */

// #define ROUNDS_MAX       999999999

/* Table with characters for base64 transformation.  */

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

	var rounds int = 5000
	var rounds_custom zend.ZendBool = 0

	/* Find beginning of salt string.  The prefix should normally always
	   be present.  Just in case it is not.  */

	if strncmp(Sha512SaltPrefix, salt, g.SizeOf("sha512_salt_prefix")-1) == 0 {

		/* Skip salt prefix.  */

		salt += g.SizeOf("sha512_salt_prefix") - 1

		/* Skip salt prefix.  */

	}
	if strncmp(salt, Sha512RoundsPrefix, g.SizeOf("sha512_rounds_prefix")-1) == 0 {
		var num *byte = salt + g.SizeOf("sha512_rounds_prefix") - 1
		var endp *byte
		var srounds zend.ZendUlong = strtoull(num, &endp, 10)
		if (*endp) == '$' {
			salt = endp + 1
			if 1000 > g.Cond(srounds < 999999999, srounds, 999999999) {
				rounds = 1000
			} else {
				if srounds < 999999999 {
					rounds = srounds
				} else {
					rounds = 999999999
				}
			}
			rounds_custom = 1
		}
	}
	if strcspn(salt, "$") < 16 {
		salt_len = strcspn(salt, "$")
	} else {
		salt_len = 16
	}
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

	cp = __phpStpncpy(buffer, Sha512SaltPrefix, g.Cond(0 > buflen, 0, buflen))
	buflen -= g.SizeOf("sha512_salt_prefix") - 1
	if rounds_custom != 0 {
		var n int = core.ApPhpSnprintf(cp, g.Cond(0 > buflen, 0, buflen), "%s%zu$", Sha512RoundsPrefix, rounds)
		cp += n
		buflen -= n
	}
	cp = __phpStpncpy(cp, salt, g.CondF1(size_t(g.Cond(0 > buflen, 0, buflen)) < salt_len, func() __auto__ { return size_t(g.Cond(0 > buflen, 0, buflen)) }, salt_len))
	buflen -= int(g.CondF1(size_t(g.Cond(0 > buflen, 0, buflen)) < salt_len, func() __auto__ { return size_t(g.Cond(0 > buflen, 0, buflen)) }, salt_len))
	if buflen > 0 {
		g.PostInc(&(*cp)) = '$'
		buflen--
	}

	// #define b64_from_24bit(B2,B1,B0,N) do { unsigned int w = ( ( B2 ) << 16 ) | ( ( B1 ) << 8 ) | ( B0 ) ; int n = ( N ) ; while ( n -- > 0 && buflen > 0 ) { * cp ++ = b64t [ w & 0x3f ] ; -- buflen ; w >>= 6 ; } } while ( 0 )

	var w uint = alt_result[0]<<16 | alt_result[21]<<8 | alt_result[42]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[22]<<16 | alt_result[43]<<8 | alt_result[1]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[44]<<16 | alt_result[2]<<8 | alt_result[23]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[3]<<16 | alt_result[24]<<8 | alt_result[45]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[25]<<16 | alt_result[46]<<8 | alt_result[4]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[47]<<16 | alt_result[5]<<8 | alt_result[26]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[6]<<16 | alt_result[27]<<8 | alt_result[48]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[28]<<16 | alt_result[49]<<8 | alt_result[7]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[50]<<16 | alt_result[8]<<8 | alt_result[29]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[9]<<16 | alt_result[30]<<8 | alt_result[51]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[31]<<16 | alt_result[52]<<8 | alt_result[10]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[53]<<16 | alt_result[11]<<8 | alt_result[32]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[12]<<16 | alt_result[33]<<8 | alt_result[54]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[34]<<16 | alt_result[55]<<8 | alt_result[13]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[56]<<16 | alt_result[14]<<8 | alt_result[35]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[15]<<16 | alt_result[36]<<8 | alt_result[57]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[37]<<16 | alt_result[58]<<8 | alt_result[16]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[59]<<16 | alt_result[17]<<8 | alt_result[38]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[18]<<16 | alt_result[39]<<8 | alt_result[60]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[40]<<16 | alt_result[61]<<8 | alt_result[19]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = alt_result[62]<<16 | alt_result[20]<<8 | alt_result[41]
	var n int = 4
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
	var w uint = 0<<16 | 0<<8 | alt_result[63]
	var n int = 2
	for g.PostDec(&n) > 0 && buflen > 0 {
		g.PostInc(&(*cp)) = B64t[w&0x3f]
		buflen--
		w >>= 6
	}
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
	core.PhpExplicitBzero(temp_result, g.SizeOf("temp_result"))
	core.PhpExplicitBzero(p_bytes, key_len)
	core.PhpExplicitBzero(s_bytes, salt_len)
	core.PhpExplicitBzero(&ctx, g.SizeOf("ctx"))
	core.PhpExplicitBzero(&alt_ctx, g.SizeOf("alt_ctx"))
	if copied_key != nil {
		core.PhpExplicitBzero(copied_key, key_len)
	}
	if copied_salt != nil {
		core.PhpExplicitBzero(copied_salt, salt_len)
	}
	return buffer
}

/* This entry point is equivalent to the `crypt' function in Unix
   libcs.  */

func PhpSha512Crypt(key *byte, salt *byte) *byte {
	/* We don't want to have an arbitrary limit in the size of the
	   password.  We can compute an upper bound for the size of the
	   result in advance and so we can prepare the buffer we pass to
	   `sha512_crypt_r'.  */

	var buffer *byte
	var buflen int = 0
	var needed int = int(g.SizeOf("sha512_salt_prefix") - 1 + g.SizeOf("sha512_rounds_prefix") + 9 + 1 + strlen(salt) + 1 + 86 + 1)
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
