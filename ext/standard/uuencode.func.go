package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

func PHP_UU_ENC(c __auto__) __auto__ {
	if c {
		return (c & 077) + ' '
	} else {
		return '`'
	}
}
func PHP_UU_ENC_C2(c int) __auto__ {
	return PHP_UU_ENC((*c)<<4&060 | (*(c + 1))>>4&017)
}
func PHP_UU_ENC_C3(c int) __auto__ {
	return PHP_UU_ENC((*(c + 1))<<2&074 | (*(c + 2))>>6&3)
}
func PHP_UU_DEC(c char) int { return c - ' '&077 }
func PhpUuencode(src *byte, src_len int) *types.String {
	var len_ int = 45
	var p *uint8
	var s *uint8
	var e *uint8
	var ee *uint8
	var dest *types.String

	/* encoded length is ~ 38% greater than the original
	   Use 1.5 for easier calculation.
	*/

	dest = types.ZendStringSafeAlloc(src_len/2, 3, 46, 0)
	p = (*uint8)(dest.GetVal())
	s = (*uint8)(src)
	e = s + src_len
	for s+3 < e {
		ee = s + len_
		if ee > e {
			ee = e
			len_ = ee - s
			if len_%3 != 0 {
				ee = s + int(floor(float64(len_/3))*3)
			}
		}
		b.PostInc(&(*p)) = PHP_UU_ENC(len_)
		for s < ee {
			b.PostInc(&(*p)) = PHP_UU_ENC((*s) >> 2)
			b.PostInc(&(*p)) = PHP_UU_ENC_C2(s)
			b.PostInc(&(*p)) = PHP_UU_ENC_C3(s)
			b.PostInc(&(*p)) = PHP_UU_ENC((*(s + 2)) & 077)
			s += 3
		}
		if len_ == 45 {
			b.PostInc(&(*p)) = '\n'
		}
	}
	if s < e {
		if len_ == 45 {
			b.PostInc(&(*p)) = PHP_UU_ENC(e - s)
			len_ = 0
		}
		b.PostInc(&(*p)) = PHP_UU_ENC((*s) >> 2)
		b.PostInc(&(*p)) = PHP_UU_ENC_C2(s)
		if e-s > 1 {
			b.PostInc(&(*p)) = PHP_UU_ENC_C3(s)
		} else {
			b.PostInc(&(*p)) = PHP_UU_ENC('0')
		}
		if e-s > 2 {
			b.PostInc(&(*p)) = PHP_UU_ENC((*(s + 2)) & 077)
		} else {
			b.PostInc(&(*p)) = PHP_UU_ENC('0')
		}
	}
	if len_ < 45 {
		b.PostInc(&(*p)) = '\n'
	}
	b.PostInc(&(*p)) = PHP_UU_ENC('0')
	b.PostInc(&(*p)) = '\n'
	*p = '0'
	dest = types.ZendStringTruncate(dest, (*byte)(p-dest.GetVal()), 0)
	return dest
}
func PhpUudecode(src *byte, src_len int) *types.String {
	var len_ int
	var total_len int = 0
	var s *byte
	var e *byte
	var p *byte
	var ee *byte
	var dest *types.String
	dest = types.ZendStringAlloc(int(ceil(src_len*0.75)), 0)
	p = dest.GetVal()
	s = src
	e = src + src_len
	for s < e {
		if b.Assign(&len_, PHP_UU_DEC(b.PostInc(&(*s)))) == 0 {
			break
		}

		/* sanity check */

		if len_ > src_len {
			goto err
		}
		total_len += len_
		ee = s + b.CondF2(len_ == 45, 60, func() int { return int(floor(len_ * 1.33)) })

		/* sanity check */

		if ee > e {
			goto err
		}
		for s < ee {
			if s+4 > e {
				goto err
			}
			b.PostInc(&(*p)) = PHP_UU_DEC(*s)<<2 | PHP_UU_DEC(*(s + 1))>>4
			b.PostInc(&(*p)) = PHP_UU_DEC(*(s + 1))<<4 | PHP_UU_DEC(*(s + 2))>>2
			b.PostInc(&(*p)) = PHP_UU_DEC(*(s + 2))<<6 | PHP_UU_DEC(*(s + 3))
			s += 4
		}
		if len_ < 45 {
			break
		}

		/* skip \n */

		s++

		/* skip \n */

	}
	b.Assert(p >= dest.GetVal())
	if b.Assign(&len_, total_len) > size_t(p-dest.GetVal()) {
		b.PostInc(&(*p)) = PHP_UU_DEC(*s)<<2 | PHP_UU_DEC(*(s + 1))>>4
		if len_ > 1 {
			b.PostInc(&(*p)) = PHP_UU_DEC(*(s + 1))<<4 | PHP_UU_DEC(*(s + 2))>>2
			if len_ > 2 {
				b.PostInc(&(*p)) = PHP_UU_DEC(*(s + 2))<<6 | PHP_UU_DEC(*(s + 3))
			}
		}
	}
	dest.SetLen(total_len)
	dest.GetVal()[dest.GetLen()] = '0'
	return dest
err:
	// types.ZendStringEfree(dest)
	return nil
}
func ZifConvertUuencode(executeData zpp.Ex, return_value zpp.Ret, data *types.Zval) {
	var src *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			src = fp.ParseStr()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if src.GetLen() < 1 {
		return_value.SetFalse()
		return
	}
	return_value.SetString(PhpUuencode(src.GetVal(), src.GetLen()))
	return
}
func ZifConvertUudecode(executeData zpp.Ex, return_value zpp.Ret, data *types.Zval) {
	var src *types.String
	var dest *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			src = fp.ParseStr()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if src.GetLen() < 1 {
		return_value.SetFalse()
		return
	}
	if b.Assign(&dest, PhpUudecode(src.GetVal(), src.GetLen())) == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The given parameter is not a valid uuencoded string")
		return_value.SetFalse()
		return
	}
	return_value.SetString(dest)
	return
}
