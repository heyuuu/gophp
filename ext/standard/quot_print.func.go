package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/zpp"
)

func PhpHex2int(c int) byte {
	if isdigit(c) {
		return c - '0'
	} else if c >= 'A' && c <= 'F' {
		return c - 'A' + 10
	} else if c >= 'a' && c <= 'f' {
		return c - 'a' + 10
	} else {
		return -1
	}
}
func PhpQuotPrintDecode(str *uint8, length int, replace_us_by_ws int) *types.String {
	var i int
	var p1 *uint8
	var p2 *uint8
	var h_nbl uint
	var l_nbl uint
	var decoded_len int
	var buf_size int
	var retval *types.String
	var hexval_tbl []uint = []uint{64, 64, 64, 64, 64, 64, 64, 64, 64, 32, 16, 64, 64, 16, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 32, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 64, 64, 64, 64, 64, 64, 64, 10, 11, 12, 13, 14, 15, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 10, 11, 12, 13, 14, 15, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64}
	if replace_us_by_ws != 0 {
		replace_us_by_ws = '_'
	}
	i = length
	p1 = str
	buf_size = length
	for i > 1 && (*p1) != '0' {
		if (*p1) == '=' {
			buf_size -= 2
			p1++
			i--
		}
		p1++
		i--
	}
	retval = types.ZendStringAlloc(buf_size, 0)
	i = length
	p1 = str
	p2 = (*uint8)(retval.GetVal())
	decoded_len = 0
	for i > 0 && (*p1) != '0' {
		if (*p1) == '=' {
			i--
			p1++
			if i == 0 || (*p1) == '0' {
				break
			}
			h_nbl = hexval_tbl[*p1]
			if h_nbl < 16 {

				/* next char should be a hexadecimal digit */

				if b.PreDec(&i) == 0 || b.Assign(&l_nbl, hexval_tbl[*(b.PreInc(&p1))]) >= 16 {
					zend.Efree(retval)
					return nil
				}
				*(b.PostInc(&p2)) = h_nbl<<4 | l_nbl
				decoded_len++
				i--
				p1++
			} else if h_nbl < 64 {

				/* soft line break */

				for h_nbl == 32 {
					if b.PreDec(&i) == 0 || b.Assign(&h_nbl, hexval_tbl[*(b.PreInc(&p1))]) == 64 {
						zend.Efree(retval)
						return nil
					}
				}
				if p1[0] == '\r' && i >= 2 && p1[1] == '\n' {
					i--
					p1++
				}
				i--
				p1++
			} else {
				zend.Efree(retval)
				return nil
			}
		} else {
			if replace_us_by_ws == (*p1) {
				*(b.PostInc(&p2)) = 'x'
			} else {
				*(b.PostInc(&p2)) = *p1
			}
			i--
			p1++
			decoded_len++
		}
	}
	*p2 = '0'
	retval.SetLen(decoded_len)
	return retval
}
func PhpQuotPrintEncode(str *uint8, length int) *types.String {
	var lp zend.ZendUlong = 0
	var c uint8
	var d *uint8
	var hex *byte = "0123456789ABCDEF"
	var ret *types.String
	ret = types.ZendStringSafeAlloc(3, length+(3*length/(PHP_QPRINT_MAXL-9)+1), 0, 0)
	d = (*uint8)(ret.GetVal())
	for b.PostDec(&length) {
		if b.Assign(&c, b.PostInc(&(*str))) == '0' && (*str) == '0' && length > 0 {
			b.PostInc(&(*d)) = '0'
			*str++
			b.PostInc(&(*d)) = (*str) - 1
			length--
			lp = 0
		} else {
			if iscntrl(c) || c == 0x7f || (c&0x80) != 0 || c == '=' || c == ' ' && (*str) == '0' {
				if b.AssignOp(&lp, "+=", 3) > PHP_QPRINT_MAXL && c <= 0x7f || c > 0x7f && c <= 0xdf && lp+3 > PHP_QPRINT_MAXL || c > 0xdf && c <= 0xef && lp+6 > PHP_QPRINT_MAXL || c > 0xef && c <= 0xf4 && lp+9 > PHP_QPRINT_MAXL {
					b.PostInc(&(*d)) = '='
					b.PostInc(&(*d)) = '0'
					b.PostInc(&(*d)) = '0'
					lp = 3
				}
				b.PostInc(&(*d)) = '='
				b.PostInc(&(*d)) = hex[c>>4]
				b.PostInc(&(*d)) = hex[c&0xf]
			} else {
				if b.PreInc(&lp) > PHP_QPRINT_MAXL {
					b.PostInc(&(*d)) = '='
					b.PostInc(&(*d)) = '0'
					b.PostInc(&(*d)) = '0'
					lp = 1
				}
				b.PostInc(&(*d)) = c
			}
		}
	}
	*d = '0'
	ret = types.ZendStringTruncate(ret, d-(*uint8)(ret.GetVal()))
	return ret
}
func ZifQuotedPrintableDecode(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval) {
	var arg1 *types.String
	var str_in *byte
	var str_out *types.String
	var i int = 0
	var j int = 0
	var k int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			arg1 = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if arg1.GetLen() == 0 {

		/* shortcut */

		return_value.SetStringVal("")
		return
	}
	str_in = arg1.GetVal()
	str_out = types.ZendStringAlloc(arg1.GetLen(), 0)
	for str_in[i] {
		switch str_in[i] {
		case '=':
			if str_in[i+1] && str_in[i+2] && isxdigit(int(str_in[i+1])) && isxdigit(int(str_in[i+2])) {
				str_out.GetStr()[b.PostInc(&j)] = (PhpHex2int(int(str_in[i+1])) << 4) + PhpHex2int(int(str_in[i+2]))
				i += 3
			} else {
				k = 1
				for str_in[i+k] && (str_in[i+k] == 32 || str_in[i+k] == 9) {

					/* Possibly, skip spaces/tabs at the end of line */

					k++

					/* Possibly, skip spaces/tabs at the end of line */

				}
				if !(str_in[i+k]) {

					/* End of line reached */

					i += k

					/* End of line reached */

				} else if str_in[i+k] == 13 && str_in[i+k+1] == 10 {

					/* CRLF */

					i += k + 2

					/* CRLF */

				} else if str_in[i+k] == 13 || str_in[i+k] == 10 {

					/* CR or LF */

					i += k + 1

					/* CR or LF */

				} else {
					str_out.GetStr()[b.PostInc(&j)] = str_in[b.PostInc(&i)]
				}
			}
		default:
			str_out.GetStr()[b.PostInc(&j)] = str_in[b.PostInc(&i)]
		}
	}
	str_out.GetStr()[j] = '0'
	str_out.SetLen(j)
	return_value.SetString(str_out)
}
func ZifQuotedPrintableEncode(executeData zpp.Ex, return_value zpp.Ret, str *types.Zval) {
	var str *types.String
	var new_str *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			str = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if str.GetLen() == 0 {
		return_value.SetStringVal("")
		return
	}
	new_str = PhpQuotPrintEncode((*uint8)(str.GetVal()), str.GetLen())
	return_value.SetString(new_str)
	return
}
