// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/types"
	"sik/zend/zpp"
)

func PhpBase64EncodeStr(str *types.String) *types.String {
	return PhpBase64Encode((*uint8)(str.GetVal()), str.GetLen())
}
func PhpBase64Decode(str *uint8, len_ int) *types.String { return PhpBase64DecodeEx(str, len_, 0) }
func PhpBase64DecodeStr(str *types.String) *types.String {
	return PhpBase64DecodeEx((*uint8)(str.GetVal()), str.GetLen(), 0)
}
func PhpBase64EncodeImpl(in *uint8, inl int, out *uint8) *uint8 {
	for inl > 2 {
		b.PostInc(&(*out)) = Base64Table[in[0]>>2]
		b.PostInc(&(*out)) = Base64Table[((in[0]&0x3)<<4)+(in[1]>>4)]
		b.PostInc(&(*out)) = Base64Table[((in[1]&0xf)<<2)+(in[2]>>6)]
		b.PostInc(&(*out)) = Base64Table[in[2]&0x3f]
		in += 3
		inl -= 3
	}

	/* now deal with the tail end of things */

	if inl != 0 {
		b.PostInc(&(*out)) = Base64Table[in[0]>>2]
		if inl > 1 {
			b.PostInc(&(*out)) = Base64Table[((in[0]&0x3)<<4)+(in[1]>>4)]
			b.PostInc(&(*out)) = Base64Table[(in[1]&0xf)<<2]
			b.PostInc(&(*out)) = Base64Pad
		} else {
			b.PostInc(&(*out)) = Base64Table[(in[0]&0x3)<<4]
			b.PostInc(&(*out)) = Base64Pad
			b.PostInc(&(*out)) = Base64Pad
		}
	}
	*out = '0'
	return out
}
func PhpBase64DecodeImpl(in *uint8, inl int, out *uint8, outl *int, strict types.ZendBool) int {
	var ch int
	var i int = 0
	var padding int = 0
	var j int = *outl

	/* run through the whole string, converting as we go */

	for b.PostDec(&inl) > 0 {
		*in++
		ch = (*in) - 1
		if ch == Base64Pad {
			padding++
			continue
		}
		ch = Base64ReverseTable[ch]
		if strict == 0 {

			/* skip unknown characters and whitespace */

			if ch < 0 {
				continue
			}

			/* skip unknown characters and whitespace */

		} else {

			/* skip whitespace */

			if ch == -1 {
				continue
			}

			/* fail on bad characters or if any data follows padding */

			if ch == -2 || padding != 0 {
				goto fail
			}

			/* fail on bad characters or if any data follows padding */

		}
		switch i % 4 {
		case 0:
			out[j] = ch << 2
		case 1:
			out[b.PostInc(&j)] |= ch >> 4
			out[j] = (ch & 0xf) << 4
		case 2:
			out[b.PostInc(&j)] |= ch >> 2
			out[j] = (ch & 0x3) << 6
		case 3:
			out[b.PostInc(&j)] |= ch
		}
		i++
	}

	/* fail if the input is truncated (only one char in last group) */

	if strict != 0 && i%4 == 1 {
		goto fail
	}

	/* fail if the padding length is wrong (not VV==, VVV=), but accept zero padding
	 * RFC 4648: "In some circumstances, the use of padding [--] is not required" */

	if strict != 0 && padding != 0 && (padding > 2 || (i+padding)%4 != 0) {
		goto fail
	}
	*outl = j
	out[j] = '0'
	return 1
fail:
	return 0
}
func PhpBase64Encode(str *uint8, length int) *types.String {
	var p *uint8
	var result *types.String
	result = types.ZendStringSafeAlloc((length+2)/3, 4*b.SizeOf("char"), 0, 0)
	p = (*uint8)(result.GetVal())
	p = PhpBase64EncodeImpl(str, length, p)
	result.SetLen(p - (*uint8)(result.GetVal()))
	return result
}
func PhpBase64DecodeEx(str *uint8, length int, strict types.ZendBool) *types.String {
	var result *types.String
	var outl int = 0
	result = types.ZendStringAlloc(length, 0)
	if PhpBase64DecodeImpl(str, length, (*uint8)(result.GetVal()), &outl, strict) == 0 {
		types.ZendStringEfree(result)
		return nil
	}
	result.SetLen(outl)
	return result
}
func ZifBase64Encode(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var str *byte
	var str_len int
	var result *types.String
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			str, str_len = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	result = PhpBase64Encode((*uint8)(str), str_len)
	return_value.SetString(result)
	return
}
func ZifBase64Decode(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var str *byte
	var strict types.ZendBool = 0
	var str_len int
	var result *types.String
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			str, str_len = fp.ParseString()
			fp.StartOptional()
			strict = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	result = PhpBase64DecodeEx((*uint8)(str), str_len, strict)
	if result != nil {
		return_value.SetString(result)
		return
	} else {
		return_value.SetFalse()
		return
	}
}
