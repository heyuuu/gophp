// <<generate>>

package standard

import (
	b "sik/builtin"
)

func AsciiToBin(ch byte) int {
	var sch signed__char = ch
	var retval int
	retval = sch - '.'
	if sch >= 'A' {
		retval = sch - ('A' - 12)
		if sch >= 'a' {
			retval = sch - ('a' - 38)
		}
	}
	retval &= 0x3f
	return retval
}
func AsciiIsUnsafe(ch byte) int { return !ch || ch == '\n' || ch == ':' }
func _cryptExtendedInit() {
	var i int
	var j int
	var b int
	var k int
	var inbit int
	var obit int
	var p *uint32
	var il *uint32
	var ir *uint32
	var fl *uint32
	var fr *uint32
	var bits28 *uint32
	var bits24 *uint32
	var inv_key_perm []u_char
	var inv_comp_perm []u_char
	var init_perm []u_char
	var final_perm []u_char
	var u_sbox [][]u_char
	var un_pbox []u_char
	bits24 = b.Assign(&bits28, Bits32+4) + 4

	/*
	 * Invert the S-boxes, reordering the input bits.
	 */

	for i = 0; i < 8; i++ {
		for j = 0; j < 64; j++ {
			b = j&0x20 | (j&1)<<4 | j>>1&0xf
			u_sbox[i][j] = Sbox[i][b]
		}
	}

	/*
	 * Convert the inverted S-boxes into 4 arrays of 8 bits.
	 * Each will handle 12 bits of the S-box input.
	 */

	for b = 0; b < 4; b++ {
		for i = 0; i < 64; i++ {
			for j = 0; j < 64; j++ {
				MSbox[b][i<<6|j] = u_sbox[b<<1][i]<<4 | u_sbox[(b<<1)+1][j]
			}
		}
	}

	/*
	 * Set up the initial & final permutations into a useful form, and
	 * initialise the inverted key permutation.
	 */

	for i = 0; i < 64; i++ {
		init_perm[b.Assign(&final_perm[i], IP[i]-1)] = i
		inv_key_perm[i] = 255
	}

	/*
	 * Invert the key permutation and initialise the inverted key
	 * compression permutation.
	 */

	for i = 0; i < 56; i++ {
		inv_key_perm[KeyPerm[i]-1] = i
		inv_comp_perm[i] = 255
	}

	/*
	 * Invert the key compression permutation.
	 */

	for i = 0; i < 48; i++ {
		inv_comp_perm[CompPerm[i]-1] = i
	}

	/*
	 * Set up the OR-mask arrays for the initial and final permutations,
	 * and for the key initial and compression permutations.
	 */

	for k = 0; k < 8; k++ {
		for i = 0; i < 256; i++ {
			*(b.Assign(&il, &IpMaskl[k][i])) = 0
			*(b.Assign(&ir, &IpMaskr[k][i])) = 0
			*(b.Assign(&fl, &FpMaskl[k][i])) = 0
			*(b.Assign(&fr, &FpMaskr[k][i])) = 0
			for j = 0; j < 8; j++ {
				inbit = 8*k + j
				if (i & Bits8[j]) != 0 {
					if b.Assign(&obit, init_perm[inbit]) < 32 {
						*il |= Bits32[obit]
					} else {
						*ir |= Bits32[obit-32]
					}
					if b.Assign(&obit, final_perm[inbit]) < 32 {
						*fl |= Bits32[obit]
					} else {
						*fr |= Bits32[obit-32]
					}
				}
			}
		}
		for i = 0; i < 128; i++ {
			*(b.Assign(&il, &KeyPermMaskl[k][i])) = 0
			*(b.Assign(&ir, &KeyPermMaskr[k][i])) = 0
			for j = 0; j < 7; j++ {
				inbit = 8*k + j
				if (i & Bits8[j+1]) != 0 {
					if b.Assign(&obit, inv_key_perm[inbit]) == 255 {
						continue
					}
					if obit < 28 {
						*il |= bits28[obit]
					} else {
						*ir |= bits28[obit-28]
					}
				}
			}
			*(b.Assign(&il, &CompMaskl[k][i])) = 0
			*(b.Assign(&ir, &CompMaskr[k][i])) = 0
			for j = 0; j < 7; j++ {
				inbit = 7*k + j
				if (i & Bits8[j+1]) != 0 {
					if b.Assign(&obit, inv_comp_perm[inbit]) == 255 {
						continue
					}
					if obit < 24 {
						*il |= bits24[obit]
					} else {
						*ir |= bits24[obit-24]
					}
				}
			}
		}
	}

	/*
	 * Invert the P-box permutation, and convert into OR-masks for
	 * handling the output of the S-box arrays setup above.
	 */

	for i = 0; i < 32; i++ {
		un_pbox[Pbox[i]-1] = i
	}
	for b = 0; b < 4; b++ {
		for i = 0; i < 256; i++ {
			*(b.Assign(&p, &Psbox[b][i])) = 0
			for j = 0; j < 8; j++ {
				if (i & Bits8[j]) != 0 {
					*p |= Bits32[un_pbox[8*b+j]]
				}
			}
		}
	}
}
func DesInitLocal(data *PhpCryptExtendedData) {
	data.SetOldRawkey1(0)
	data.SetOldRawkey0(data.GetOldRawkey1())
	data.SetSaltbits(0)
	data.SetOldSalt(0)
	data.SetInitialized(1)
}
func SetupSalt(salt uint32, data *PhpCryptExtendedData) {
	var obit uint32
	var saltbit uint32
	var saltbits uint32
	var i int
	if salt == data.GetOldSalt() {
		return
	}
	data.SetOldSalt(salt)
	saltbits = 0
	saltbit = 1
	obit = 0x800000
	for i = 0; i < 24; i++ {
		if (salt & saltbit) != 0 {
			saltbits |= obit
		}
		saltbit <<= 1
		obit >>= 1
	}
	data.SetSaltbits(saltbits)
}
func DesSetkey(key *byte, data *PhpCryptExtendedData) int {
	var k0 uint32
	var k1 uint32
	var rawkey0 uint32
	var rawkey1 uint32
	var shifts int
	var round int
	rawkey0 = uint32(u_char(key[3] | uint32(u_char(key[2]<<8)) | uint32(u_char(key[1]<<16)) | uint32(u_char(key[0]<<24))))
	rawkey1 = uint32(u_char(key[7] | uint32(u_char(key[6]<<8)) | uint32(u_char(key[5]<<16)) | uint32(u_char(key[4]<<24))))
	if (rawkey0|rawkey1) != 0 && rawkey0 == data.GetOldRawkey0() && rawkey1 == data.GetOldRawkey1() {

		/*
		 * Already setup for this key.
		 * This optimisation fails on a zero key (which is weak and
		 * has bad parity anyway) in order to simplify the starting
		 * conditions.
		 */

		return 0

		/*
		 * Already setup for this key.
		 * This optimisation fails on a zero key (which is weak and
		 * has bad parity anyway) in order to simplify the starting
		 * conditions.
		 */

	}
	data.SetOldRawkey0(rawkey0)
	data.SetOldRawkey1(rawkey1)

	/*
	 *    Do key permutation and split into two 28-bit subkeys.
	 */

	k0 = KeyPermMaskl[0][rawkey0>>25] | KeyPermMaskl[1][rawkey0>>17&0x7f] | KeyPermMaskl[2][rawkey0>>9&0x7f] | KeyPermMaskl[3][rawkey0>>1&0x7f] | KeyPermMaskl[4][rawkey1>>25] | KeyPermMaskl[5][rawkey1>>17&0x7f] | KeyPermMaskl[6][rawkey1>>9&0x7f] | KeyPermMaskl[7][rawkey1>>1&0x7f]
	k1 = KeyPermMaskr[0][rawkey0>>25] | KeyPermMaskr[1][rawkey0>>17&0x7f] | KeyPermMaskr[2][rawkey0>>9&0x7f] | KeyPermMaskr[3][rawkey0>>1&0x7f] | KeyPermMaskr[4][rawkey1>>25] | KeyPermMaskr[5][rawkey1>>17&0x7f] | KeyPermMaskr[6][rawkey1>>9&0x7f] | KeyPermMaskr[7][rawkey1>>1&0x7f]

	/*
	 *    Rotate subkeys and do compression permutation.
	 */

	shifts = 0
	for round = 0; round < 16; round++ {
		var t0 uint32
		var t1 uint32
		shifts += KeyShifts[round]
		t0 = k0<<shifts | k0>>28 - shifts
		t1 = k1<<shifts | k1>>28 - shifts
		data.GetEnKeysl()[round] = CompMaskl[0][t0>>21&0x7f] | CompMaskl[1][t0>>14&0x7f] | CompMaskl[2][t0>>7&0x7f] | CompMaskl[3][t0&0x7f] | CompMaskl[4][t1>>21&0x7f] | CompMaskl[5][t1>>14&0x7f] | CompMaskl[6][t1>>7&0x7f] | CompMaskl[7][t1&0x7f]
		data.GetDeKeysl()[15-round] = data.GetEnKeysl()[round]
		data.GetEnKeysr()[round] = CompMaskr[0][t0>>21&0x7f] | CompMaskr[1][t0>>14&0x7f] | CompMaskr[2][t0>>7&0x7f] | CompMaskr[3][t0&0x7f] | CompMaskr[4][t1>>21&0x7f] | CompMaskr[5][t1>>14&0x7f] | CompMaskr[6][t1>>7&0x7f] | CompMaskr[7][t1&0x7f]
		data.GetDeKeysr()[15-round] = data.GetEnKeysr()[round]
	}
	return 0
}
func DoDes(
	l_in uint32,
	r_in uint32,
	l_out *uint32,
	r_out *uint32,
	count int,
	data *PhpCryptExtendedData,
) int {
	/*
	 *    l_in, r_in, l_out, and r_out are in pseudo-"big-endian" format.
	 */

	var l uint32
	var r uint32
	var kl *uint32
	var kr *uint32
	var kl1 *uint32
	var kr1 *uint32
	var f uint32
	var r48l uint32
	var r48r uint32
	var saltbits uint32
	var round int
	if count == 0 {
		return 1
	} else if count > 0 {

		/*
		 * Encrypting
		 */

		kl1 = data.GetEnKeysl()
		kr1 = data.GetEnKeysr()
	} else {

		/*
		 * Decrypting
		 */

		count = -count
		kl1 = data.GetDeKeysl()
		kr1 = data.GetDeKeysr()
	}

	/*
	 *    Do initial permutation (IP).
	 */

	l = IpMaskl[0][l_in>>24] | IpMaskl[1][l_in>>16&0xff] | IpMaskl[2][l_in>>8&0xff] | IpMaskl[3][l_in&0xff] | IpMaskl[4][r_in>>24] | IpMaskl[5][r_in>>16&0xff] | IpMaskl[6][r_in>>8&0xff] | IpMaskl[7][r_in&0xff]
	r = IpMaskr[0][l_in>>24] | IpMaskr[1][l_in>>16&0xff] | IpMaskr[2][l_in>>8&0xff] | IpMaskr[3][l_in&0xff] | IpMaskr[4][r_in>>24] | IpMaskr[5][r_in>>16&0xff] | IpMaskr[6][r_in>>8&0xff] | IpMaskr[7][r_in&0xff]
	saltbits = data.GetSaltbits()
	for b.PostDec(&count) {

		/*
		 * Do each round.
		 */

		kl = kl1
		kr = kr1
		round = 16
		for b.PostDec(&round) {

			/*
			 * Expand R to 48 bits (simulate the E-box).
			 */

			r48l = (r&0x1)<<23 | (r&0xf8000000)>>9 | (r&0x1f800000)>>11 | (r&0x1f80000)>>13 | (r&0x1f8000)>>15
			r48r = (r&0x1f800)<<7 | (r&0x1f80)<<5 | (r&0x1f8)<<3 | (r&0x1f)<<1 | (r&0x80000000)>>31

			/*
			 * Do salting for crypt() and friends, and
			 * XOR with the permuted key.
			 */

			f = (r48l ^ r48r) & saltbits
			r48l ^= f ^ b.PostInc(&(*kl))
			r48r ^= f ^ b.PostInc(&(*kr))

			/*
			 * Do sbox lookups (which shrink it back to 32 bits)
			 * and do the pbox permutation at the same time.
			 */

			f = Psbox[0][MSbox[0][r48l>>12]] | Psbox[1][MSbox[1][r48l&0xfff]] | Psbox[2][MSbox[2][r48r>>12]] | Psbox[3][MSbox[3][r48r&0xfff]]

			/*
			 * Now that we've permuted things, complete f().
			 */

			f ^= l
			l = r
			r = f
		}
		r = l
		l = f
	}

	/*
	 * Do final permutation (inverse of IP).
	 */

	*l_out = FpMaskl[0][l>>24] | FpMaskl[1][l>>16&0xff] | FpMaskl[2][l>>8&0xff] | FpMaskl[3][l&0xff] | FpMaskl[4][r>>24] | FpMaskl[5][r>>16&0xff] | FpMaskl[6][r>>8&0xff] | FpMaskl[7][r&0xff]
	*r_out = FpMaskr[0][l>>24] | FpMaskr[1][l>>16&0xff] | FpMaskr[2][l>>8&0xff] | FpMaskr[3][l&0xff] | FpMaskr[4][r>>24] | FpMaskr[5][r>>16&0xff] | FpMaskr[6][r>>8&0xff] | FpMaskr[7][r&0xff]
	return 0
}
func DesCipher(in *byte, out *byte, salt uint32, count int, data *PhpCryptExtendedData) int {
	var l_out uint32 = 0
	var r_out uint32 = 0
	var rawl uint32
	var rawr uint32
	var retval int
	SetupSalt(salt, data)
	rawl = uint32(u_char(in[3] | uint32(u_char(in[2]<<8)) | uint32(u_char(in[1]<<16)) | uint32(u_char(in[0]<<24))))
	rawr = uint32(u_char(in[7] | uint32(u_char(in[6]<<8)) | uint32(u_char(in[5]<<16)) | uint32(u_char(in[4]<<24))))
	retval = DoDes(rawl, rawr, &l_out, &r_out, count, data)
	out[0] = l_out >> 24
	out[1] = l_out >> 16
	out[2] = l_out >> 8
	out[3] = l_out
	out[4] = r_out >> 24
	out[5] = r_out >> 16
	out[6] = r_out >> 8
	out[7] = r_out
	return retval
}
func _cryptExtendedR(key *uint8, setting *byte, data *PhpCryptExtendedData) *byte {
	var i int
	var count uint32
	var salt uint32
	var l uint32
	var r0 uint32
	var r1 uint32
	var keybuf []uint32
	var p *u_char
	var q *u_char
	if data.GetInitialized() == 0 {
		DesInitLocal(data)
	}

	/*
	 * Copy the key, shifting each character up by one bit
	 * and padding with zeros.
	 */

	q = (*u_char)(keybuf)
	for size_t(q-(*u_char)(keybuf)) < b.SizeOf("keybuf") {
		b.PostInc(&(*q)) = (*key) << 1
		if (*key) != 0 {
			key++
		}
	}
	if DesSetkey((*byte)(keybuf), data) != 0 {
		return nil
	}
	if (*setting) == _PASSWORD_EFMT1 {

		/*
		 * "new"-style:
		 *    setting - underscore, 4 chars of count, 4 chars of salt
		 *    key - unlimited characters
		 */

		i = 1
		count = 0
		for ; i < 5; i++ {
			var value int = AsciiToBin(setting[i])
			if Ascii64[value] != setting[i] {
				return nil
			}
			count |= value << (i - 1) * 6
		}
		if count == 0 {
			return nil
		}
		i = 5
		salt = 0
		for ; i < 9; i++ {
			var value int = AsciiToBin(setting[i])
			if Ascii64[value] != setting[i] {
				return nil
			}
			salt |= value << (i - 5) * 6
		}
		for (*key) != 0 {

			/*
			 * Encrypt the key with itself.
			 */

			if DesCipher((*byte)(keybuf), (*byte)(keybuf), 0, 1, data) != 0 {
				return nil
			}

			/*
			 * And XOR with the next 8 characters of the key.
			 */

			q = (*u_char)(keybuf)
			for size_t(q-(*u_char)(keybuf)) < b.SizeOf("keybuf") && (*key) != 0 {
				b.PostInc(&(*q)) ^= b.PostInc(&(*key)) << 1
			}
			if DesSetkey((*byte)(keybuf), data) != 0 {
				return nil
			}
		}
		memcpy(data.GetOutput(), setting, 9)
		data.GetOutput()[9] = '0'
		p = (*u_char)(data.GetOutput() + 9)
	} else {

		/*
		 * "old"-style:
		 *    setting - 2 chars of salt
		 *    key - up to 8 characters
		 */

		count = 25
		if AsciiIsUnsafe(setting[0]) != 0 || AsciiIsUnsafe(setting[1]) != 0 {
			return nil
		}
		salt = AsciiToBin(setting[1])<<6 | AsciiToBin(setting[0])
		data.GetOutput()[0] = setting[0]
		data.GetOutput()[1] = setting[1]
		p = (*u_char)(data.GetOutput() + 2)
	}
	SetupSalt(salt, data)

	/*
	 * Do it.
	 */

	if DoDes(0, 0, &r0, &r1, count, data) != 0 {
		return nil
	}

	/*
	 * Now encode the result...
	 */

	l = r0 >> 8
	b.PostInc(&(*p)) = Ascii64[l>>18&0x3f]
	b.PostInc(&(*p)) = Ascii64[l>>12&0x3f]
	b.PostInc(&(*p)) = Ascii64[l>>6&0x3f]
	b.PostInc(&(*p)) = Ascii64[l&0x3f]
	l = r0<<16 | r1>>16&0xffff
	b.PostInc(&(*p)) = Ascii64[l>>18&0x3f]
	b.PostInc(&(*p)) = Ascii64[l>>12&0x3f]
	b.PostInc(&(*p)) = Ascii64[l>>6&0x3f]
	b.PostInc(&(*p)) = Ascii64[l&0x3f]
	l = r1 << 2
	b.PostInc(&(*p)) = Ascii64[l>>12&0x3f]
	b.PostInc(&(*p)) = Ascii64[l>>6&0x3f]
	b.PostInc(&(*p)) = Ascii64[l&0x3f]
	*p = 0
	return data.GetOutput()
}
