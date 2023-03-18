// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func MakeDigest(md5str *byte, digest *uint8) { MakeDigestEx(md5str, digest, 16) }
func MakeDigestEx(md5str *byte, digest *uint8, len_ int) {
	var hexits []byte = "0123456789abcdef"
	var i int
	for i = 0; i < len_; i++ {
		md5str[i*2] = hexits[digest[i]>>4]
		md5str[i*2+1] = hexits[digest[i]&0xf]
	}
	md5str[len_*2] = '0'
}
func PhpIfMd5(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *types.ZendString
	var raw_output types.ZendBool = 0
	var context PHP_MD5_CTX
	var digest []uint8
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &arg, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			raw_output = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	PHP_MD5Init(&context)
	PHP_MD5Update(&context, arg.GetVal(), arg.GetLen())
	PHP_MD5Final(digest, &context)
	if raw_output != 0 {
		return_value.SetRawString(b.CastStr((*byte)(digest), 16))
		return
	} else {
		return_value.SetString(types.ZendStringAlloc(32, 0))
		MakeDigestEx(return_value.GetStr().GetVal(), digest, 16)
	}
}
func PhpIfMd5File(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var arg *byte
	var arg_len int
	var raw_output types.ZendBool = 0
	var buf []uint8
	var digest []uint8
	var context PHP_MD5_CTX
	var n ssize_t
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			arg, arg_len = fp.ParsePath()
			fp.StartOptional()
			raw_output = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
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
	PHP_MD5Init(&context)
	for b.Assign(&n, core.PhpStreamRead(stream, (*byte)(buf), b.SizeOf("buf"))) > 0 {
		PHP_MD5Update(&context, buf, n)
	}

	/* XXX this probably can be improved with some number of retries */

	if core.PhpStreamEof(stream) == 0 {
		core.PhpStreamClose(stream)
		PHP_MD5Final(digest, &context)
		return_value.SetFalse()
		return
	}
	core.PhpStreamClose(stream)
	PHP_MD5Final(digest, &context)
	if raw_output != 0 {
		return_value.SetRawString(b.CastStr((*byte)(digest), 16))
		return
	} else {
		return_value.SetString(types.ZendStringAlloc(32, 0))
		MakeDigestEx(return_value.GetStr().GetVal(), digest, 16)
	}
}
func F(x uint32, y uint32, z uint32) int { return z ^ x&(y^z) }
func G(x uint32, y uint32, z uint32) int { return y ^ z&(x^y) }
func H(x uint32, y uint32, z uint32) int { return x ^ y ^ z }
func I(x uint32, y uint32, z uint32) int { return y ^ (x | ^z) }
func SET(n int) __auto__ {
	ctx.block[n] = uint32(ptr[n*4] | uint32(ptr[n*4+1]<<8) | uint32(ptr[n*4+2]<<16) | uint32(ptr[n*4+3]<<24))
	return ctx.block[n]
}
func GET(n int) __auto__ { return ctx.block[n] }
func Body(ctx *PHP_MD5_CTX, data any, size int) any {
	var ptr *uint8
	var a uint32
	var b uint32
	var c uint32
	var d uint32
	var saved_a uint32
	var saved_b uint32
	var saved_c uint32
	var saved_d uint32
	ptr = data
	a = ctx.GetA()
	b = ctx.GetB()
	c = ctx.GetC()
	d = ctx.GetD()
	for {
		saved_a = a
		saved_b = b
		saved_c = c
		saved_d = d

		/* Round 1 */

		a += F(b, c, d) + SET(0) + 0xd76aa478
		a = a<<7 | (a&0xffffffff)>>32 - 7
		a += b
		d += F(a, b, c) + SET(1) + 0xe8c7b756
		d = d<<12 | (d&0xffffffff)>>32 - 12
		d += a
		c += F(d, a, b) + SET(2) + 0x242070db
		c = c<<17 | (c&0xffffffff)>>32 - 17
		c += d
		b += F(c, d, a) + SET(3) + 0xc1bdceee
		b = b<<22 | (b&0xffffffff)>>32 - 22
		b += c
		a += F(b, c, d) + SET(4) + 0xf57c0faf
		a = a<<7 | (a&0xffffffff)>>32 - 7
		a += b
		d += F(a, b, c) + SET(5) + 0x4787c62a
		d = d<<12 | (d&0xffffffff)>>32 - 12
		d += a
		c += F(d, a, b) + SET(6) + 0xa8304613
		c = c<<17 | (c&0xffffffff)>>32 - 17
		c += d
		b += F(c, d, a) + SET(7) + 0xfd469501
		b = b<<22 | (b&0xffffffff)>>32 - 22
		b += c
		a += F(b, c, d) + SET(8) + 0x698098d8
		a = a<<7 | (a&0xffffffff)>>32 - 7
		a += b
		d += F(a, b, c) + SET(9) + 0x8b44f7af
		d = d<<12 | (d&0xffffffff)>>32 - 12
		d += a
		c += F(d, a, b) + SET(10) + 0xffff5bb1
		c = c<<17 | (c&0xffffffff)>>32 - 17
		c += d
		b += F(c, d, a) + SET(11) + 0x895cd7be
		b = b<<22 | (b&0xffffffff)>>32 - 22
		b += c
		a += F(b, c, d) + SET(12) + 0x6b901122
		a = a<<7 | (a&0xffffffff)>>32 - 7
		a += b
		d += F(a, b, c) + SET(13) + 0xfd987193
		d = d<<12 | (d&0xffffffff)>>32 - 12
		d += a
		c += F(d, a, b) + SET(14) + 0xa679438e
		c = c<<17 | (c&0xffffffff)>>32 - 17
		c += d
		b += F(c, d, a) + SET(15) + 0x49b40821
		b = b<<22 | (b&0xffffffff)>>32 - 22
		b += c

		/* Round 2 */

		a += G(b, c, d) + GET(1) + 0xf61e2562
		a = a<<5 | (a&0xffffffff)>>32 - 5
		a += b
		d += G(a, b, c) + GET(6) + 0xc040b340
		d = d<<9 | (d&0xffffffff)>>32 - 9
		d += a
		c += G(d, a, b) + GET(11) + 0x265e5a51
		c = c<<14 | (c&0xffffffff)>>32 - 14
		c += d
		b += G(c, d, a) + GET(0) + 0xe9b6c7aa
		b = b<<20 | (b&0xffffffff)>>32 - 20
		b += c
		a += G(b, c, d) + GET(5) + 0xd62f105d
		a = a<<5 | (a&0xffffffff)>>32 - 5
		a += b
		d += G(a, b, c) + GET(10) + 0x2441453
		d = d<<9 | (d&0xffffffff)>>32 - 9
		d += a
		c += G(d, a, b) + GET(15) + 0xd8a1e681
		c = c<<14 | (c&0xffffffff)>>32 - 14
		c += d
		b += G(c, d, a) + GET(4) + 0xe7d3fbc8
		b = b<<20 | (b&0xffffffff)>>32 - 20
		b += c
		a += G(b, c, d) + GET(9) + 0x21e1cde6
		a = a<<5 | (a&0xffffffff)>>32 - 5
		a += b
		d += G(a, b, c) + GET(14) + 0xc33707d6
		d = d<<9 | (d&0xffffffff)>>32 - 9
		d += a
		c += G(d, a, b) + GET(3) + 0xf4d50d87
		c = c<<14 | (c&0xffffffff)>>32 - 14
		c += d
		b += G(c, d, a) + GET(8) + 0x455a14ed
		b = b<<20 | (b&0xffffffff)>>32 - 20
		b += c
		a += G(b, c, d) + GET(13) + 0xa9e3e905
		a = a<<5 | (a&0xffffffff)>>32 - 5
		a += b
		d += G(a, b, c) + GET(2) + 0xfcefa3f8
		d = d<<9 | (d&0xffffffff)>>32 - 9
		d += a
		c += G(d, a, b) + GET(7) + 0x676f02d9
		c = c<<14 | (c&0xffffffff)>>32 - 14
		c += d
		b += G(c, d, a) + GET(12) + 0x8d2a4c8a
		b = b<<20 | (b&0xffffffff)>>32 - 20
		b += c

		/* Round 3 */

		a += H(b, c, d) + GET(5) + 0xfffa3942
		a = a<<4 | (a&0xffffffff)>>32 - 4
		a += b
		d += H(a, b, c) + GET(8) + 0x8771f681
		d = d<<11 | (d&0xffffffff)>>32 - 11
		d += a
		c += H(d, a, b) + GET(11) + 0x6d9d6122
		c = c<<16 | (c&0xffffffff)>>32 - 16
		c += d
		b += H(c, d, a) + GET(14) + 0xfde5380c
		b = b<<23 | (b&0xffffffff)>>32 - 23
		b += c
		a += H(b, c, d) + GET(1) + 0xa4beea44
		a = a<<4 | (a&0xffffffff)>>32 - 4
		a += b
		d += H(a, b, c) + GET(4) + 0x4bdecfa9
		d = d<<11 | (d&0xffffffff)>>32 - 11
		d += a
		c += H(d, a, b) + GET(7) + 0xf6bb4b60
		c = c<<16 | (c&0xffffffff)>>32 - 16
		c += d
		b += H(c, d, a) + GET(10) + 0xbebfbc70
		b = b<<23 | (b&0xffffffff)>>32 - 23
		b += c
		a += H(b, c, d) + GET(13) + 0x289b7ec6
		a = a<<4 | (a&0xffffffff)>>32 - 4
		a += b
		d += H(a, b, c) + GET(0) + 0xeaa127fa
		d = d<<11 | (d&0xffffffff)>>32 - 11
		d += a
		c += H(d, a, b) + GET(3) + 0xd4ef3085
		c = c<<16 | (c&0xffffffff)>>32 - 16
		c += d
		b += H(c, d, a) + GET(6) + 0x4881d05
		b = b<<23 | (b&0xffffffff)>>32 - 23
		b += c
		a += H(b, c, d) + GET(9) + 0xd9d4d039
		a = a<<4 | (a&0xffffffff)>>32 - 4
		a += b
		d += H(a, b, c) + GET(12) + 0xe6db99e5
		d = d<<11 | (d&0xffffffff)>>32 - 11
		d += a
		c += H(d, a, b) + GET(15) + 0x1fa27cf8
		c = c<<16 | (c&0xffffffff)>>32 - 16
		c += d
		b += H(c, d, a) + GET(2) + 0xc4ac5665
		b = b<<23 | (b&0xffffffff)>>32 - 23
		b += c

		/* Round 4 */

		a += I(b, c, d) + GET(0) + 0xf4292244
		a = a<<6 | (a&0xffffffff)>>32 - 6
		a += b
		d += I(a, b, c) + GET(7) + 0x432aff97
		d = d<<10 | (d&0xffffffff)>>32 - 10
		d += a
		c += I(d, a, b) + GET(14) + 0xab9423a7
		c = c<<15 | (c&0xffffffff)>>32 - 15
		c += d
		b += I(c, d, a) + GET(5) + 0xfc93a039
		b = b<<21 | (b&0xffffffff)>>32 - 21
		b += c
		a += I(b, c, d) + GET(12) + 0x655b59c3
		a = a<<6 | (a&0xffffffff)>>32 - 6
		a += b
		d += I(a, b, c) + GET(3) + 0x8f0ccc92
		d = d<<10 | (d&0xffffffff)>>32 - 10
		d += a
		c += I(d, a, b) + GET(10) + 0xffeff47d
		c = c<<15 | (c&0xffffffff)>>32 - 15
		c += d
		b += I(c, d, a) + GET(1) + 0x85845dd1
		b = b<<21 | (b&0xffffffff)>>32 - 21
		b += c
		a += I(b, c, d) + GET(8) + 0x6fa87e4f
		a = a<<6 | (a&0xffffffff)>>32 - 6
		a += b
		d += I(a, b, c) + GET(15) + 0xfe2ce6e0
		d = d<<10 | (d&0xffffffff)>>32 - 10
		d += a
		c += I(d, a, b) + GET(6) + 0xa3014314
		c = c<<15 | (c&0xffffffff)>>32 - 15
		c += d
		b += I(c, d, a) + GET(13) + 0x4e0811a1
		b = b<<21 | (b&0xffffffff)>>32 - 21
		b += c
		a += I(b, c, d) + GET(4) + 0xf7537e82
		a = a<<6 | (a&0xffffffff)>>32 - 6
		a += b
		d += I(a, b, c) + GET(11) + 0xbd3af235
		d = d<<10 | (d&0xffffffff)>>32 - 10
		d += a
		c += I(d, a, b) + GET(2) + 0x2ad7d2bb
		c = c<<15 | (c&0xffffffff)>>32 - 15
		c += d
		b += I(c, d, a) + GET(9) + 0xeb86d391
		b = b<<21 | (b&0xffffffff)>>32 - 21
		b += c
		a += saved_a
		b += saved_b
		c += saved_c
		d += saved_d
		ptr += 64
		if !(b.AssignOp(&size, "-=", 64)) {
			break
		}
	}
	ctx.SetA(a)
	ctx.SetB(b)
	ctx.SetC(c)
	ctx.SetD(d)
	return ptr
}
func PHP_MD5Init(ctx *PHP_MD5_CTX) {
	ctx.SetA(0x67452301)
	ctx.SetB(0xefcdab89)
	ctx.SetC(0x98badcfe)
	ctx.SetD(0x10325476)
	ctx.SetLo(0)
	ctx.SetHi(0)
}
func PHP_MD5Update(ctx *PHP_MD5_CTX, data any, size int) {
	var saved_lo uint32
	var used uint32
	var free uint32
	saved_lo = ctx.GetLo()
	if b.Assign(&(ctx.GetLo()), saved_lo+size&0x1fffffff) < saved_lo {
		ctx.GetHi()++
	}
	ctx.SetHi(ctx.GetHi() + size>>29)
	used = saved_lo & 0x3f
	if used != 0 {
		free = 64 - used
		if size < free {
			memcpy(ctx.GetBuffer()[used], data, size)
			return
		}
		memcpy(ctx.GetBuffer()[used], data, free)
		data = (*uint8)(data + free)
		size -= free
		Body(ctx, ctx.GetBuffer(), 64)
	}
	if size >= 64 {
		data = Body(ctx, data, size & ^int(0x3f))
		size &= 0x3f
	}
	memcpy(ctx.GetBuffer(), data, size)
}
func PHP_MD5Final(result *uint8, ctx *PHP_MD5_CTX) {
	var used uint32
	var free uint32
	used = ctx.GetLo() & 0x3f
	ctx.GetBuffer()[b.PostInc(&used)] = 0x80
	free = 64 - used
	if free < 8 {
		memset(ctx.GetBuffer()[used], 0, free)
		Body(ctx, ctx.GetBuffer(), 64)
		used = 0
		free = 64
	}
	memset(ctx.GetBuffer()[used], 0, free-8)
	ctx.SetLo(ctx.GetLo() << 3)
	ctx.GetBuffer()[56] = ctx.GetLo()
	ctx.GetBuffer()[57] = ctx.GetLo() >> 8
	ctx.GetBuffer()[58] = ctx.GetLo() >> 16
	ctx.GetBuffer()[59] = ctx.GetLo() >> 24
	ctx.GetBuffer()[60] = ctx.GetHi()
	ctx.GetBuffer()[61] = ctx.GetHi() >> 8
	ctx.GetBuffer()[62] = ctx.GetHi() >> 16
	ctx.GetBuffer()[63] = ctx.GetHi() >> 24
	Body(ctx, ctx.GetBuffer(), 64)
	result[0] = ctx.GetA()
	result[1] = ctx.GetA() >> 8
	result[2] = ctx.GetA() >> 16
	result[3] = ctx.GetA() >> 24
	result[4] = ctx.GetB()
	result[5] = ctx.GetB() >> 8
	result[6] = ctx.GetB() >> 16
	result[7] = ctx.GetB() >> 24
	result[8] = ctx.GetC()
	result[9] = ctx.GetC() >> 8
	result[10] = ctx.GetC() >> 16
	result[11] = ctx.GetC() >> 24
	result[12] = ctx.GetD()
	result[13] = ctx.GetD() >> 8
	result[14] = ctx.GetD() >> 16
	result[15] = ctx.GetD() >> 24
	zend.ZEND_SECURE_ZERO(ctx, b.SizeOf("* ctx"))
}
