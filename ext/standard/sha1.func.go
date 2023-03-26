package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/types"
	"sik/zend/zpp"
)

func MakeSha1Digest(sha1str *byte, digest *uint8) { MakeDigestEx(sha1str, digest, 20) }
func ZifSha1(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval, _ zpp.Opt, rawOutput *types.Zval) {
	var arg *types.String
	var raw_output types.ZendBool = 0
	var context PHP_SHA1_CTX
	var digest []uint8
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			arg = fp.ParseStr()
			fp.StartOptional()
			raw_output = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	PHP_SHA1Init(&context)
	PHP_SHA1Update(&context, (*uint8)(arg.GetVal()), arg.GetLen())
	PHP_SHA1Final(digest, &context)
	if raw_output != 0 {
		return_value.SetStringVal(b.CastStr((*byte)(digest), 20))
		return
	} else {
		return_value.SetString(types.ZendStringAlloc(40, 0))
		MakeDigestEx(return_value.GetStr().GetVal(), digest, 20)
	}
}
func ZifSha1File(executeData zpp.Ex, return_value zpp.Ret, filename *types.Zval, _ zpp.Opt, rawOutput *types.Zval) {
	var arg *byte
	var arg_len int
	var raw_output types.ZendBool = 0
	var buf []uint8
	var digest []uint8
	var context PHP_SHA1_CTX
	var n ssize_t
	var stream *core.PhpStream
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			arg, arg_len = fp.ParsePath()
			fp.StartOptional()
			raw_output = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	stream = core.PhpStreamOpenWrapper(arg, "rb", core.REPORT_ERRORS, nil)
	if stream == nil {
		return_value.SetFalse()
		return
	}
	PHP_SHA1Init(&context)
	for b.Assign(&n, core.PhpStreamRead(stream, (*byte)(buf), b.SizeOf("buf"))) > 0 {
		PHP_SHA1Update(&context, buf, n)
	}
	PHP_SHA1Final(digest, &context)
	core.PhpStreamClose(stream)
	if raw_output != 0 {
		return_value.SetStringVal(b.CastStr((*byte)(digest), 20))
		return
	} else {
		return_value.SetString(types.ZendStringAlloc(40, 0))
		MakeDigestEx(return_value.GetStr().GetVal(), digest, 20)
	}
}
func ROTATE_LEFT(x uint32, n int) int { return x<<n | x>>32 - n }
func W(i int) __auto__ {
	tmp = x[i-3&15] ^ x[i-8&15] ^ x[i-14&15] ^ x[i&15]
	x[i&15] = ROTATE_LEFT(tmp, 1)
	return x[i&15]
}
func FF(
	a uint32,
	b uint32,
	c uint32,
	d uint32,
	e __auto__,
	w int,
) {
	e += F(b, c, d) + w + uint32(0x5a827999)
	e += ROTATE_LEFT(a, 5)
	b = ROTATE_LEFT(b, 30)
}
func GG(
	a uint32,
	b uint32,
	c uint32,
	d uint32,
	e __auto__,
	w int,
) {
	e += G(b, c, d) + w + uint32(0x6ed9eba1)
	e += ROTATE_LEFT(a, 5)
	b = ROTATE_LEFT(b, 30)
}
func HH(
	a uint32,
	b uint32,
	c uint32,
	d uint32,
	e __auto__,
	w int,
) {
	e += H(b, c, d) + w + uint32(0x8f1bbcdc)
	e += ROTATE_LEFT(a, 5)
	b = ROTATE_LEFT(b, 30)
}
func II(
	a uint32,
	b uint32,
	c uint32,
	d uint32,
	e __auto__,
	w int,
) {
	e += I(b, c, d) + w + uint32(0xca62c1d6)
	e += ROTATE_LEFT(a, 5)
	b = ROTATE_LEFT(b, 30)
}
func PHP_SHA1Init(context *PHP_SHA1_CTX) {
	context.GetCount()[1] = 0
	context.GetCount()[0] = context.GetCount()[1]

	/* Load magic initialization constants.
	 */

	context.GetState()[0] = 0x67452301
	context.GetState()[1] = 0xefcdab89
	context.GetState()[2] = 0x98badcfe
	context.GetState()[3] = 0x10325476
	context.GetState()[4] = 0xc3d2e1f0
}
func PHP_SHA1Update(context *PHP_SHA1_CTX, input *uint8, inputLen int) {
	var i uint
	var index uint
	var partLen uint

	/* Compute number of bytes mod 64 */

	index = uint(context.GetCount()[0] >> 3 & 0x3f)

	/* Update number of bits */

	if b.AssignOp(&context.GetCount()[0], "+=", uint32(inputLen<<3)) < uint32(inputLen<<3) {
		context.GetCount()[1]++
	}
	context.GetCount()[1] += uint32(inputLen >> 29)
	partLen = 64 - index

	/* Transform as many times as possible.
	 */

	if inputLen >= partLen {
		memcpy((*uint8)(context.GetBuffer()[index]), (*uint8)(input), partLen)
		SHA1Transform(context.GetState(), context.GetBuffer())
		for i = partLen; i+63 < inputLen; i += 64 {
			SHA1Transform(context.GetState(), &input[i])
		}
		index = 0
	} else {
		i = 0
	}

	/* Buffer remaining input */

	memcpy((*uint8)(context.GetBuffer()[index]), (*uint8)(&input[i]), inputLen-i)

	/* Buffer remaining input */
}
func PHP_SHA1Final(digest []uint8, context *PHP_SHA1_CTX) {
	var bits []uint8
	var index uint
	var padLen uint

	/* Save number of bits */

	bits[7] = context.GetCount()[0] & 0xff
	bits[6] = context.GetCount()[0] >> 8 & 0xff
	bits[5] = context.GetCount()[0] >> 16 & 0xff
	bits[4] = context.GetCount()[0] >> 24 & 0xff
	bits[3] = context.GetCount()[1] & 0xff
	bits[2] = context.GetCount()[1] >> 8 & 0xff
	bits[1] = context.GetCount()[1] >> 16 & 0xff
	bits[0] = context.GetCount()[1] >> 24 & 0xff

	/* Pad out to 56 mod 64.
	 */

	index = uint(context.GetCount()[0] >> 3 & 0x3f)
	if index < 56 {
		padLen = 56 - index
	} else {
		padLen = 120 - index
	}
	PHP_SHA1Update(context, PADDING, padLen)

	/* Append length (before padding) */

	PHP_SHA1Update(context, bits, 8)

	/* Store state in digest */

	SHA1Encode(digest, context.GetState(), 20)

	/* Zeroize sensitive information.
	 */

	zend.ZEND_SECURE_ZERO((*uint8)(context), b.SizeOf("* context"))

	/* Zeroize sensitive information.
	 */
}
