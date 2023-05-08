package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

func StrfilterRot13Filter(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *streams.PhpStreamBucketBrigade,
	buckets_out *streams.PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	for buckets_in.GetHead() != nil {
		bucket = streams.PhpStreamBucketMakeWriteable(buckets_in.GetHead())
		str.PhpStrtr(bucket.GetBuf(), bucket.GetBuflen(), Rot13From, Rot13To)
		consumed += bucket.GetBuflen()
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
}
func StrfilterRot13Create(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter {
	return streams.PhpStreamFilterAlloc(&StrfilterRot13Ops, nil, persistent)
}
func StrfilterToupperFilter(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *streams.PhpStreamBucketBrigade,
	buckets_out *streams.PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	for buckets_in.GetHead() != nil {
		bucket = streams.PhpStreamBucketMakeWriteable(buckets_in.GetHead())
		str.PhpStrtr(bucket.GetBuf(), bucket.GetBuflen(), Lowercase, Uppercase)
		consumed += bucket.GetBuflen()
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
}
func StrfilterTolowerFilter(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *streams.PhpStreamBucketBrigade,
	buckets_out *streams.PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	for buckets_in.GetHead() != nil {
		bucket = streams.PhpStreamBucketMakeWriteable(buckets_in.GetHead())
		str.PhpStrtr(bucket.GetBuf(), bucket.GetBuflen(), Uppercase, Lowercase)
		consumed += bucket.GetBuflen()
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
}
func StrfilterToupperCreate(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter {
	return streams.PhpStreamFilterAlloc(&StrfilterToupperOps, nil, persistent)
}
func StrfilterTolowerCreate(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter {
	return streams.PhpStreamFilterAlloc(&StrfilterTolowerOps, nil, persistent)
}
func PhpStripTagsFilterCtor(inst *PhpStripTagsFilter, allowed_tags *types.String, persistent int) int {
	if allowed_tags != nil {
		if nil == b.Assign(&(inst.GetAllowedTags()), zend.Pemalloc(allowed_tags.GetLen()+1)) {
			return types.FAILURE
		}
		memcpy((*byte)(inst.GetAllowedTags()), allowed_tags.GetVal(), allowed_tags.GetLen()+1)
		inst.SetAllowedTagsLen(int(allowed_tags.GetLen()))
	} else {
		inst.SetAllowedTags(nil)
	}
	inst.SetState(0)
	inst.SetPersistent(persistent)
	return types.SUCCESS
}
func PhpStripTagsFilterDtor(inst *PhpStripTagsFilter) {
	if inst.GetAllowedTags() != nil {
		zend.Pefree(any(inst.GetAllowedTags()), inst.GetPersistent())
	}
}
func StrfilterStripTagsFilter(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *streams.PhpStreamBucketBrigade,
	buckets_out *streams.PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	var inst *PhpStripTagsFilter = (*PhpStripTagsFilter)(thisfilter.GetAbstract().Ptr())
	for buckets_in.GetHead() != nil {
		bucket = streams.PhpStreamBucketMakeWriteable(buckets_in.GetHead())
		consumed = bucket.GetBuflen()

		result, state := str.PhpStripTags(b.CastStr(bucket.GetBuf(), bucket.GetBuflen()), inst.GetState(), b.CastStr(inst.GetAllowedTags(), inst.GetAllowedTagsLen()))
		inst.SetState(state)
		bucket.SetBuf_([]byte(result))

		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
}
func StrfilterStripTagsDtor(thisfilter *core.PhpStreamFilter) {
	b.Assert(thisfilter.GetAbstract().Ptr() != nil)
	PhpStripTagsFilterDtor((*PhpStripTagsFilter)(thisfilter.GetAbstract().Ptr()))
	zend.Pefree(thisfilter.GetAbstract().Ptr(), (*PhpStripTagsFilter)(types.Z_PTR(thisfilter.GetAbstract())).GetPersistent())
}
func StrfilterStripTagsCreate(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter {
	var inst *PhpStripTagsFilter
	var filter *core.PhpStreamFilter = nil
	var allowed_tags *types.String = nil
	core.PhpErrorDocref(nil, faults.E_DEPRECATED, "The string.strip_tags filter is deprecated")
	if filterparams != nil {
		if filterparams.IsType(types.IS_ARRAY) {
			var tags_ss zend.SmartStr = zend.MakeSmartStr(0)
			filterparams.Array().Foreach(func(_ types.ArrayKey, tmp *types.Zval) {
				zend.ConvertToStringEx(tmp)
				tags_ss.AppendByte('<')
				tags_ss.AppendString(tmp.String().GetStr())
				tags_ss.AppendByte('>')
			})
			tags_ss.ZeroTail()
			allowed_tags = tags_ss.GetS()
		} else {
			allowed_tags = zend.ZvalGetString(filterparams)
		}

		/* Exception during string conversion. */

		if zend.EG__().GetException() != nil {
			if allowed_tags != nil {
				// types.ZendStringRelease(allowed_tags)
			}
			return nil
		}

		/* Exception during string conversion. */

	}
	inst = zend.Pemalloc(b.SizeOf("php_strip_tags_filter"))
	if PhpStripTagsFilterCtor(inst, allowed_tags, persistent) == types.SUCCESS {
		filter = streams.PhpStreamFilterAlloc(&StrfilterStripTagsOps, inst, persistent)
	} else {
		zend.Pefree(inst, persistent)
	}
	if allowed_tags != nil {
		// types.ZendStringRelease(allowed_tags)
	}
	return filter
}
func PhpConvConvert(a *PhpConv, b **byte, c *int, d **byte, e *int) PhpConvErrT {
	return (*PhpConv)(a).GetConvertOp()((*PhpConv)(a), b, c, d, e)
}
func PhpConvDtor(a *PhpConv) { (*PhpConv)(a).GetDtor()(a) }
func PhpConvBase64EncodeCtor(
	inst *PhpConvBase64Encode,
	line_len uint,
	lbchars *byte,
	lbchars_len int,
	lbchars_dup int,
	persistent int,
) PhpConvErrT {
	inst.GetSuper().SetConvertOp(PhpConvConvertFunc(PhpConvBase64EncodeConvert))
	inst.GetSuper().SetDtor(PhpConvDtorFunc(PhpConvBase64EncodeDtor))
	inst.SetEremLen(0)
	inst.SetLineCcnt(line_len)
	inst.SetLineLen(line_len)
	if lbchars != nil {
		if lbchars_dup != 0 {
			inst.SetLbchars(zend.Pestrdup(lbchars))
		} else {
			inst.SetLbchars(lbchars)
		}
		inst.SetLbcharsLen(lbchars_len)
	} else {
		inst.SetLbchars(nil)
	}
	inst.SetLbcharsDup(lbchars_dup)
	inst.SetPersistent(persistent)
	return PHP_CONV_ERR_SUCCESS
}
func PhpConvBase64EncodeDtor(inst *PhpConvBase64Encode) {
	b.Assert(inst != nil)
	if inst.GetLbcharsDup() != 0 && inst.GetLbchars() != nil {
		zend.Pefree(any(inst.GetLbchars()), inst.GetPersistent())
	}
}
func PhpConvBase64EncodeFlush(inst *PhpConvBase64Encode, in_pp **byte, in_left_p *int, out_pp **byte, out_left_p *int) PhpConvErrT {
	var err PhpConvErrT = PHP_CONV_ERR_SUCCESS
	var pd *uint8
	var ocnt int
	var line_ccnt uint
	pd = (*uint8)(*out_pp)
	ocnt = *out_left_p
	line_ccnt = inst.GetLineCcnt()
	switch inst.GetEremLen() {
	case 0:

		/* do nothing */

	case 1:
		if line_ccnt < 4 && inst.GetLbchars() != nil {
			if ocnt < inst.GetLbcharsLen() {
				return PHP_CONV_ERR_TOO_BIG
			}
			memcpy(pd, inst.GetLbchars(), inst.GetLbcharsLen())
			pd += inst.GetLbcharsLen()
			ocnt -= inst.GetLbcharsLen()
			line_ccnt = inst.GetLineLen()
		}
		if ocnt < 4 {
			err = PHP_CONV_ERR_TOO_BIG
			goto out
		}
		*(b.PostInc(&pd)) = B64TblEnc[inst.GetErem()[0]>>2]
		*(b.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[0]<<4)]
		*(b.PostInc(&pd)) = '='
		*(b.PostInc(&pd)) = '='
		inst.SetEremLen(0)
		ocnt -= 4
		line_ccnt -= 4
	case 2:
		if line_ccnt < 4 && inst.GetLbchars() != nil {
			if ocnt < inst.GetLbcharsLen() {
				return PHP_CONV_ERR_TOO_BIG
			}
			memcpy(pd, inst.GetLbchars(), inst.GetLbcharsLen())
			pd += inst.GetLbcharsLen()
			ocnt -= inst.GetLbcharsLen()
			line_ccnt = inst.GetLineLen()
		}
		if ocnt < 4 {
			err = PHP_CONV_ERR_TOO_BIG
			goto out
		}
		*(b.PostInc(&pd)) = B64TblEnc[inst.GetErem()[0]>>2]
		*(b.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[0]<<4|inst.GetErem()[1]>>4)]
		*(b.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[1]<<2)]
		*(b.PostInc(&pd)) = '='
		inst.SetEremLen(0)
		ocnt -= 4
		line_ccnt -= 4
	default:

		/* should not happen... */

		err = PHP_CONV_ERR_UNKNOWN
	}
out:
	*out_pp = (*byte)(pd)
	*out_left_p = ocnt
	inst.SetLineCcnt(line_ccnt)
	return err
}
func PhpConvBase64EncodeConvert(inst *PhpConvBase64Encode, in_pp **byte, in_left_p *int, out_pp **byte, out_left_p *int) PhpConvErrT {
	var err PhpConvErrT = PHP_CONV_ERR_SUCCESS
	var ocnt int
	var icnt int
	var ps *uint8
	var pd *uint8
	var line_ccnt uint
	if in_pp == nil || in_left_p == nil {
		return PhpConvBase64EncodeFlush(inst, in_pp, in_left_p, out_pp, out_left_p)
	}
	pd = (*uint8)(*out_pp)
	ocnt = *out_left_p
	ps = (*uint8)(*in_pp)
	icnt = *in_left_p
	line_ccnt = inst.GetLineCcnt()

	/* consume the remainder first */

	switch inst.GetEremLen() {
	case 1:
		if icnt >= 2 {
			if line_ccnt < 4 && inst.GetLbchars() != nil {
				if ocnt < inst.GetLbcharsLen() {
					return PHP_CONV_ERR_TOO_BIG
				}
				memcpy(pd, inst.GetLbchars(), inst.GetLbcharsLen())
				pd += inst.GetLbcharsLen()
				ocnt -= inst.GetLbcharsLen()
				line_ccnt = inst.GetLineLen()
			}
			if ocnt < 4 {
				err = PHP_CONV_ERR_TOO_BIG
				goto out
			}
			*(b.PostInc(&pd)) = B64TblEnc[inst.GetErem()[0]>>2]
			*(b.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[0]<<4|ps[0]>>4)]
			*(b.PostInc(&pd)) = B64TblEnc[uint8(ps[0]<<2|ps[1]>>6)]
			*(b.PostInc(&pd)) = B64TblEnc[ps[1]]
			ocnt -= 4
			ps += 2
			icnt -= 2
			inst.SetEremLen(0)
			line_ccnt -= 4
		}
	case 2:
		if icnt >= 1 {
			if inst.GetLineCcnt() < 4 && inst.GetLbchars() != nil {
				if ocnt < inst.GetLbcharsLen() {
					return PHP_CONV_ERR_TOO_BIG
				}
				memcpy(pd, inst.GetLbchars(), inst.GetLbcharsLen())
				pd += inst.GetLbcharsLen()
				ocnt -= inst.GetLbcharsLen()
				line_ccnt = inst.GetLineLen()
			}
			if ocnt < 4 {
				err = PHP_CONV_ERR_TOO_BIG
				goto out
			}
			*(b.PostInc(&pd)) = B64TblEnc[inst.GetErem()[0]>>2]
			*(b.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[0]<<4|inst.GetErem()[1]>>4)]
			*(b.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[1]<<2|ps[0]>>6)]
			*(b.PostInc(&pd)) = B64TblEnc[ps[0]]
			ocnt -= 4
			ps += 1
			icnt -= 1
			inst.SetEremLen(0)
			line_ccnt -= 4
		}
	}
	for icnt >= 3 {
		if line_ccnt < 4 && inst.GetLbchars() != nil {
			if ocnt < inst.GetLbcharsLen() {
				err = PHP_CONV_ERR_TOO_BIG
				goto out
			}
			memcpy(pd, inst.GetLbchars(), inst.GetLbcharsLen())
			pd += inst.GetLbcharsLen()
			ocnt -= inst.GetLbcharsLen()
			line_ccnt = inst.GetLineLen()
		}
		if ocnt < 4 {
			err = PHP_CONV_ERR_TOO_BIG
			goto out
		}
		*(b.PostInc(&pd)) = B64TblEnc[ps[0]>>2]
		*(b.PostInc(&pd)) = B64TblEnc[uint8(ps[0]<<4|ps[1]>>4)]
		*(b.PostInc(&pd)) = B64TblEnc[uint8(ps[1]<<2|ps[2]>>6)]
		*(b.PostInc(&pd)) = B64TblEnc[ps[2]]
		ps += 3
		icnt -= 3
		ocnt -= 4
		line_ccnt -= 4
	}
	for ; icnt > 0; icnt-- {
		inst.GetErem()[b.PostInc(&(inst.GetEremLen()))] = *(b.PostInc(&ps))
	}
out:
	*in_pp = (*byte)(ps)
	*in_left_p = icnt
	*out_pp = (*byte)(pd)
	*out_left_p = ocnt
	inst.SetLineCcnt(line_ccnt)
	return err
}
func PhpConvBase64DecodeCtor(inst *PhpConvBase64Decode) int {
	inst.GetSuper().SetConvertOp(PhpConvConvertFunc(PhpConvBase64DecodeConvert))
	inst.GetSuper().SetDtor(PhpConvDtorFunc(PhpConvBase64DecodeDtor))
	inst.SetUrem(0)
	inst.SetUremNbits(0)
	inst.SetUstat(0)
	inst.SetEos(0)
	return types.SUCCESS
}
func PhpConvBase64DecodeDtor(inst *PhpConvBase64Decode) {}
func PhpConvBase64DecodeConvert(inst *PhpConvBase64Decode, in_pp **byte, in_left_p *int, out_pp **byte, out_left_p *int) PhpConvErrT {
	var err PhpConvErrT
	var urem uint
	var urem_nbits uint
	var pack uint
	var pack_bcnt uint
	var ps *uint8
	var pd *uint8
	var icnt int
	var ocnt int
	var ustat uint
	var nbitsof_pack uint = 8
	if in_pp == nil || in_left_p == nil {
		if inst.GetEos() != 0 || inst.GetUremNbits() == 0 {
			return PHP_CONV_ERR_SUCCESS
		}
		return PHP_CONV_ERR_UNEXPECTED_EOS
	}
	err = PHP_CONV_ERR_SUCCESS
	ps = (*uint8)(*in_pp)
	pd = (*uint8)(*out_pp)
	icnt = *in_left_p
	ocnt = *out_left_p
	urem = inst.GetUrem()
	urem_nbits = inst.GetUremNbits()
	ustat = inst.GetUstat()
	pack = 0
	pack_bcnt = nbitsof_pack
	for {
		if pack_bcnt >= urem_nbits {
			pack_bcnt -= urem_nbits
			pack |= urem << pack_bcnt
			urem_nbits = 0
		} else {
			urem_nbits -= pack_bcnt
			pack |= urem >> urem_nbits
			urem &= 0xffff>>16 - urem_nbits
			pack_bcnt = 0
		}
		if pack_bcnt > 0 {
			var i uint
			if icnt < 1 {
				break
			}
			i = B64TblDec[uint(*(b.PostInc(&ps)))]
			icnt--
			ustat |= i & 0x80
			if (i & 0xc0) == 0 {
				if ustat != 0 {
					err = PHP_CONV_ERR_INVALID_SEQ
					break
				}
				if 6 <= pack_bcnt {
					pack_bcnt -= 6
					pack |= i << pack_bcnt
					urem = 0
				} else {
					urem_nbits = 6 - pack_bcnt
					pack |= i >> urem_nbits
					urem = i&0xffff>>16 - urem_nbits
					pack_bcnt = 0
				}
			} else if ustat != 0 {
				if pack_bcnt == 8 || pack_bcnt == 2 {
					err = PHP_CONV_ERR_INVALID_SEQ
					break
				}
				inst.SetEos(1)
			}
		}
		if (pack_bcnt | ustat) == 0 {
			if ocnt < 1 {
				err = PHP_CONV_ERR_TOO_BIG
				break
			}
			*(b.PostInc(&pd)) = pack
			ocnt--
			pack = 0
			pack_bcnt = nbitsof_pack
		}
	}
	if urem_nbits >= pack_bcnt {
		urem |= pack<<urem_nbits - pack_bcnt
		urem_nbits += nbitsof_pack - pack_bcnt
	} else {
		urem |= pack>>pack_bcnt - urem_nbits
		urem_nbits += nbitsof_pack - pack_bcnt
	}
	inst.SetUrem(urem)
	inst.SetUremNbits(urem_nbits)
	inst.SetUstat(ustat)
	*in_pp = (*byte)(ps)
	*in_left_p = icnt
	*out_pp = (*byte)(pd)
	*out_left_p = ocnt
	return err
}
func PhpConvQprintEncodeDtor(inst *PhpConvQprintEncode) {
	b.Assert(inst != nil)
	if inst.GetLbcharsDup() != 0 && inst.GetLbchars() != nil {
		zend.Pefree(any(inst.GetLbchars()), inst.GetPersistent())
	}
}
func PhpConvQprintEncodeConvert(inst *PhpConvQprintEncode, in_pp **byte, in_left_p *int, out_pp **byte, out_left_p *int) PhpConvErrT {
	var err PhpConvErrT = PHP_CONV_ERR_SUCCESS
	var ps *uint8
	var pd *uint8
	var icnt int
	var ocnt int
	var c uint
	var line_ccnt uint
	var lb_ptr uint
	var lb_cnt uint
	var trail_ws uint
	var opts int
	var qp_digits []byte = "0123456789ABCDEF"
	line_ccnt = inst.GetLineCcnt()
	opts = inst.GetOpts()
	lb_ptr = inst.GetLbPtr()
	lb_cnt = inst.GetLbCnt()
	if in_pp == nil || in_left_p == nil {
		return PHP_CONV_ERR_SUCCESS
	}
	ps = (*uint8)(*in_pp)
	icnt = *in_left_p
	pd = (*uint8)(*out_pp)
	ocnt = *out_left_p
	trail_ws = 0
	for {
		if (opts&PHP_CONV_QPRINT_OPT_BINARY) == 0 && inst.GetLbchars() != nil && inst.GetLbcharsLen() > 0 {

			/* look ahead for the line break chars to make a right decision
			 * how to consume incoming characters */

			if icnt > 0 && (*ps) == inst.GetLbchars()[lb_cnt] {
				lb_cnt++
				if lb_cnt >= inst.GetLbcharsLen() {
					var i uint
					if ocnt < lb_cnt {
						lb_cnt--
						err = PHP_CONV_ERR_TOO_BIG
						break
					}
					for i = 0; i < lb_cnt; i++ {
						*(b.PostInc(&pd)) = inst.GetLbchars()[i]
						ocnt--
					}
					line_ccnt = inst.GetLineLen()
					lb_cnt = 0
					lb_ptr = lb_cnt
				}
				ps++
				icnt--
				continue
			}

			/* look ahead for the line break chars to make a right decision
			 * how to consume incoming characters */

		}
		if lb_ptr >= lb_cnt && icnt == 0 {
			break
		}
		if lb_ptr < lb_cnt {
			c = inst.GetLbchars()[lb_ptr]
		} else {
			c = *ps
		}
		if (opts&PHP_CONV_QPRINT_OPT_BINARY) == 0 && trail_ws == 0 && (c == '\t' || c == ' ') {
			if line_ccnt < 2 && inst.GetLbchars() != nil {
				if ocnt < inst.GetLbcharsLen()+1 {
					err = PHP_CONV_ERR_TOO_BIG
					break
				}
				*(b.PostInc(&pd)) = '='
				ocnt--
				line_ccnt--
				memcpy(pd, inst.GetLbchars(), inst.GetLbcharsLen())
				pd += inst.GetLbcharsLen()
				ocnt -= inst.GetLbcharsLen()
				line_ccnt = inst.GetLineLen()
			} else {
				if ocnt < 1 {
					err = PHP_CONV_ERR_TOO_BIG
					break
				}

				/* Check to see if this is EOL whitespace. */

				if inst.GetLbchars() != nil {
					var ps2 *uint8
					var lb_cnt2 uint
					var j int
					lb_cnt2 = 0
					ps2 = ps
					trail_ws = 1
					for j = icnt - 1; j > 0; {
						if (*ps2) == inst.GetLbchars()[lb_cnt2] {
							lb_cnt2++
							if lb_cnt2 >= inst.GetLbcharsLen() {

								/* Found trailing ws. Reset to top of main
								 * for loop to allow for code to do necessary
								 * wrapping/encoding. */

								break

								/* Found trailing ws. Reset to top of main
								 * for loop to allow for code to do necessary
								 * wrapping/encoding. */

							}
						} else if lb_cnt2 != 0 || (*ps2) != '\t' && (*ps2) != ' ' {

							/* At least one non-EOL character following, so
							 * don't need to encode ws. */

							trail_ws = 0
							break
						} else {
							trail_ws++
						}
						j--
						ps2++
					}
				}
				if trail_ws == 0 {
					*(b.PostInc(&pd)) = c
					ocnt--
					line_ccnt--
					if lb_ptr < lb_cnt {
						lb_ptr++
					} else {
						lb_ptr = 0
						lb_cnt = lb_ptr
						icnt--
						ps++
					}
				}
			}
		} else if ((opts&PHP_CONV_QPRINT_OPT_FORCE_ENCODE_FIRST) == 0 || line_ccnt < inst.GetLineLen()) && (c >= 33 && c <= 60 || c >= 62 && c <= 126) {
			if line_ccnt < 2 && inst.GetLbchars() != nil {
				if ocnt < inst.GetLbcharsLen()+1 {
					err = PHP_CONV_ERR_TOO_BIG
					break
				}
				*(b.PostInc(&pd)) = '='
				ocnt--
				line_ccnt--
				memcpy(pd, inst.GetLbchars(), inst.GetLbcharsLen())
				pd += inst.GetLbcharsLen()
				ocnt -= inst.GetLbcharsLen()
				line_ccnt = inst.GetLineLen()
			}
			if ocnt < 1 {
				err = PHP_CONV_ERR_TOO_BIG
				break
			}
			*(b.PostInc(&pd)) = c
			ocnt--
			line_ccnt--
			if lb_ptr < lb_cnt {
				lb_ptr++
			} else {
				lb_ptr = 0
				lb_cnt = lb_ptr
				icnt--
				ps++
			}
		} else {
			if line_ccnt < 4 && inst.GetLbchars() != nil {
				if ocnt < inst.GetLbcharsLen()+1 {
					err = PHP_CONV_ERR_TOO_BIG
					break
				}
				*(b.PostInc(&pd)) = '='
				ocnt--
				line_ccnt--
				memcpy(pd, inst.GetLbchars(), inst.GetLbcharsLen())
				pd += inst.GetLbcharsLen()
				ocnt -= inst.GetLbcharsLen()
				line_ccnt = inst.GetLineLen()
			}
			if ocnt < 3 {
				err = PHP_CONV_ERR_TOO_BIG
				break
			}
			*(b.PostInc(&pd)) = '='
			*(b.PostInc(&pd)) = qp_digits[c>>4]
			*(b.PostInc(&pd)) = qp_digits[c&0xf]
			ocnt -= 3
			line_ccnt -= 3
			if trail_ws > 0 {
				trail_ws--
			}
			if lb_ptr < lb_cnt {
				lb_ptr++
			} else {
				lb_ptr = 0
				lb_cnt = lb_ptr
				icnt--
				ps++
			}
		}
	}
	*in_pp = (*byte)(ps)
	*in_left_p = icnt
	*out_pp = (*byte)(pd)
	*out_left_p = ocnt
	inst.SetLineCcnt(line_ccnt)
	inst.SetLbPtr(lb_ptr)
	inst.SetLbCnt(lb_cnt)
	return err
}
func PhpConvQprintEncodeCtor(
	inst *PhpConvQprintEncode,
	line_len uint,
	lbchars *byte,
	lbchars_len int,
	lbchars_dup int,
	opts int,
	persistent int,
) PhpConvErrT {
	if line_len < 4 && lbchars != nil {
		return PHP_CONV_ERR_TOO_BIG
	}
	inst.GetSuper().SetConvertOp(PhpConvConvertFunc(PhpConvQprintEncodeConvert))
	inst.GetSuper().SetDtor(PhpConvDtorFunc(PhpConvQprintEncodeDtor))
	inst.SetLineCcnt(line_len)
	inst.SetLineLen(line_len)
	if lbchars != nil {
		if lbchars_dup != 0 {
			inst.SetLbchars(zend.Pestrdup(lbchars))
		} else {
			inst.SetLbchars(lbchars)
		}
		inst.SetLbcharsLen(lbchars_len)
	} else {
		inst.SetLbchars(nil)
	}
	inst.SetLbcharsDup(lbchars_dup)
	inst.SetPersistent(persistent)
	inst.SetOpts(opts)
	inst.SetLbPtr(0)
	inst.SetLbCnt(inst.GetLbPtr())
	return PHP_CONV_ERR_SUCCESS
}
func PhpConvQprintDecodeDtor(inst *PhpConvQprintDecode) {
	b.Assert(inst != nil)
	if inst.GetLbcharsDup() != 0 && inst.GetLbchars() != nil {
		zend.Pefree(any(inst.GetLbchars()), inst.GetPersistent())
	}
}
func PhpConvQprintDecodeConvert(inst *PhpConvQprintDecode, in_pp **byte, in_left_p *int, out_pp **byte, out_left_p *int) PhpConvErrT {
	var err PhpConvErrT = PHP_CONV_ERR_SUCCESS
	var icnt int
	var ocnt int
	var ps *uint8
	var pd *uint8
	var scan_stat uint
	var next_char uint
	var lb_ptr uint
	var lb_cnt uint
	lb_ptr = inst.GetLbPtr()
	lb_cnt = inst.GetLbCnt()
	if in_pp == nil || in_left_p == nil {
		if inst.GetScanStat() != 0 {
			return PHP_CONV_ERR_UNEXPECTED_EOS
		}
		return PHP_CONV_ERR_SUCCESS
	}
	ps = (*uint8)(*in_pp)
	icnt = *in_left_p
	pd = (*uint8)(*out_pp)
	ocnt = *out_left_p
	scan_stat = inst.GetScanStat()
	next_char = inst.GetNextChar()
	for {
		switch scan_stat {
		case 0:
			if icnt == 0 {
				goto out
			}
			if (*ps) == '=' {
				scan_stat = 1
			} else {
				if ocnt < 1 {
					err = PHP_CONV_ERR_TOO_BIG
					goto out
				}
				*(b.PostInc(&pd)) = *ps
				ocnt--
			}
			ps++
			icnt--
		case 1:
			if icnt == 0 {
				goto out
			}
			if (*ps) == ' ' || (*ps) == '\t' {
				scan_stat = 4
				ps++
				icnt--
				break
			} else if inst.GetLbchars() == nil && lb_cnt == 0 && (*ps) == '\r' {

				/* auto-detect line endings, looks like network line ending \r\n (could be mac \r) */

				lb_cnt++
				scan_stat = 5
				ps++
				icnt--
				break
			} else if inst.GetLbchars() == nil && lb_cnt == 0 && (*ps) == '\n' {

				/* auto-detect line endings, looks like unix-lineendings, not to spec, but it is seem in the wild, a lot */

				lb_ptr = 0
				lb_cnt = lb_ptr
				scan_stat = 0
				ps++
				icnt--
				break
			} else if lb_cnt < inst.GetLbcharsLen() && (*ps) == uint8(inst.GetLbchars()[lb_cnt]) {
				lb_cnt++
				scan_stat = 5
				ps++
				icnt--
				break
			}
			fallthrough
		case 2:
			if icnt == 0 {
				goto out
			}
			if !(isxdigit(int(*ps))) {
				err = PHP_CONV_ERR_INVALID_SEQ
				goto out
			}
			next_char = next_char<<4 | b.Cond((*ps) >= 'A', (*ps)-0x37, (*ps)-0x30)
			scan_stat++
			ps++
			icnt--
			if scan_stat != 3 {
				break
			}
			fallthrough
		case 3:
			if ocnt < 1 {
				err = PHP_CONV_ERR_TOO_BIG
				goto out
			}
			*(b.PostInc(&pd)) = next_char
			ocnt--
			scan_stat = 0
		case 4:
			if icnt == 0 {
				goto out
			}
			if lb_cnt < inst.GetLbcharsLen() && (*ps) == uint8(inst.GetLbchars()[lb_cnt]) {
				lb_cnt++
				scan_stat = 5
			} else if (*ps) != '\t' && (*ps) != ' ' {
				err = PHP_CONV_ERR_INVALID_SEQ
				goto out
			}
			ps++
			icnt--
		case 5:
			if inst.GetLbchars() == nil && lb_cnt == 1 && (*ps) == '\n' {

				/* auto-detect soft line breaks, found network line break */

				lb_ptr = 0
				lb_cnt = lb_ptr
				scan_stat = 0
				ps++
				icnt--
			} else if inst.GetLbchars() == nil && lb_cnt > 0 {

				/* auto-detect soft line breaks, found mac line break */

				lb_ptr = 0
				lb_cnt = lb_ptr
				scan_stat = 0
			} else if lb_cnt >= inst.GetLbcharsLen() {

				/* soft line break */

				lb_ptr = 0
				lb_cnt = lb_ptr
				scan_stat = 0
			} else if icnt > 0 {
				if (*ps) == uint8(inst.GetLbchars()[lb_cnt]) {
					lb_cnt++
					ps++
					icnt--
				} else {
					scan_stat = 6
				}
			} else {
				goto out
			}
		case 6:
			if lb_ptr < lb_cnt {
				if ocnt < 1 {
					err = PHP_CONV_ERR_TOO_BIG
					goto out
				}
				*(b.PostInc(&pd)) = inst.GetLbchars()[b.PostInc(&lb_ptr)]
				ocnt--
			} else {
				scan_stat = 0
				lb_ptr = 0
				lb_cnt = lb_ptr
			}
		}
	}
out:
	*in_pp = (*byte)(ps)
	*in_left_p = icnt
	*out_pp = (*byte)(pd)
	*out_left_p = ocnt
	inst.SetScanStat(scan_stat)
	inst.SetLbPtr(lb_ptr)
	inst.SetLbCnt(lb_cnt)
	inst.SetNextChar(next_char)
	return err
}
func PhpConvQprintDecodeCtor(inst *PhpConvQprintDecode, lbchars *byte, lbchars_len int, lbchars_dup int, persistent int) PhpConvErrT {
	inst.GetSuper().SetConvertOp(PhpConvConvertFunc(PhpConvQprintDecodeConvert))
	inst.GetSuper().SetDtor(PhpConvDtorFunc(PhpConvQprintDecodeDtor))
	inst.SetScanStat(0)
	inst.SetNextChar(0)
	inst.SetLbCnt(0)
	inst.SetLbPtr(inst.GetLbCnt())
	if lbchars != nil {
		if lbchars_dup != 0 {
			inst.SetLbchars(zend.Pestrdup(lbchars))
		} else {
			inst.SetLbchars(lbchars)
		}
		inst.SetLbcharsLen(lbchars_len)
	} else {
		inst.SetLbchars(nil)
		inst.SetLbcharsLen(0)
	}
	inst.SetLbcharsDup(lbchars_dup)
	inst.SetPersistent(persistent)
	return PHP_CONV_ERR_SUCCESS
}
func PhpConvGetStringPropEx(
	ht *types.Array,
	pretval **byte,
	pretval_len *int,
	field_name string,
	field_name_len int,
	persistent int,
) PhpConvErrT {
	var tmpval *types.Zval
	*pretval = nil
	*pretval_len = 0
	if b.Assign(&tmpval, (*types.Array)(ht).KeyFind(b.CastStr(field_name, field_name_len-1))) != nil {
		var tmp *types.String
		var str *types.String = zend.ZvalGetTmpString(tmpval, &tmp)
		*pretval = zend.Pemalloc(str.GetLen() + 1)
		*pretval_len = str.GetLen()
		memcpy(*pretval, str.GetVal(), str.GetLen()+1)
		// zend.ZendTmpStringRelease(tmp)
	} else {
		return PHP_CONV_ERR_NOT_FOUND
	}
	return PHP_CONV_ERR_SUCCESS
}
func PhpConvGetUlongPropEx(ht *types.Array, pretval *zend.ZendUlong, field_name string, field_name_len int) PhpConvErrT {
	var tmpval *types.Zval = (*types.Array)(ht).KeyFind(b.CastStr(field_name, field_name_len-1))
	if tmpval != nil {
		var lval zend.ZendLong = zend.ZvalGetLong(tmpval)
		if lval < 0 {
			*pretval = 0
		} else {
			*pretval = lval
		}
		return PHP_CONV_ERR_SUCCESS
	} else {
		*pretval = 0
		return PHP_CONV_ERR_NOT_FOUND
	}
}
func PhpConvGetBoolPropEx(ht *types.Array, pretval *int, field_name string, field_name_len int) PhpConvErrT {
	var tmpval *types.Zval = (*types.Array)(ht).KeyFind(b.CastStr(field_name, field_name_len-1))
	if tmpval != nil {
		*pretval = zend.IZendIsTrue(tmpval)
		return PHP_CONV_ERR_SUCCESS
	} else {
		*pretval = 0
		return PHP_CONV_ERR_NOT_FOUND
	}
}
func PhpConvGetUintPropEx(ht *types.Array, pretval *uint, field_name string, field_name_len int) int {
	var l zend.ZendUlong
	var err PhpConvErrT
	*pretval = 0
	if b.Assign(&err, PhpConvGetUlongPropEx(ht, &l, field_name, field_name_len)) == PHP_CONV_ERR_SUCCESS {
		*pretval = uint(l)
	}
	return err
}
func PhpConvOpen(conv_mode int, options *types.Array, persistent int) *PhpConv {
	/* FIXME: I'll have to replace this ugly code by something neat
	   (factories?) in the near future. */

	var retval *PhpConv = nil
	switch conv_mode {
	case PHP_CONV_BASE64_ENCODE:
		var line_len uint = 0
		var lbchars *byte = nil
		var lbchars_len int
		if options != nil {
			PhpConvGetStringPropEx(options, &lbchars, &lbchars_len, "line-break-chars", b.SizeOf("\"line-break-chars\""), 0)
			PhpConvGetUintPropEx(options, &line_len, "line-length", b.SizeOf("\"line-length\""))
			if line_len < 4 {
				if lbchars != nil {
					zend.Pefree(lbchars, 0)
				}
				lbchars = nil
			} else {
				if lbchars == nil {
					lbchars = zend.Pestrdup("\r\n")
					lbchars_len = 2
				}
			}
		}
		retval = zend.Pemalloc(b.SizeOf("php_conv_base64_encode"))
		if lbchars != nil {
			if PhpConvBase64EncodeCtor((*PhpConvBase64Encode)(retval), line_len, lbchars, lbchars_len, 1, persistent) != 0 {
				if lbchars != nil {
					zend.Pefree(lbchars, 0)
				}
				goto out_failure
			}
			zend.Pefree(lbchars, 0)
		} else {
			if PhpConvBase64EncodeCtor((*PhpConvBase64Encode)(retval), 0, nil, 0, 0, persistent) != 0 {
				goto out_failure
			}
		}
	case PHP_CONV_BASE64_DECODE:
		retval = zend.Pemalloc(b.SizeOf("php_conv_base64_decode"))
		if PhpConvBase64DecodeCtor((*PhpConvBase64Decode)(retval)) != 0 {
			goto out_failure
		}
	case PHP_CONV_QPRINT_ENCODE:
		var line_len uint = 0
		var lbchars *byte = nil
		var lbchars_len int
		var opts int = 0
		if options != nil {
			var opt_binary int = 0
			var opt_force_encode_first int = 0
			PhpConvGetStringPropEx(options, &lbchars, &lbchars_len, "line-break-chars", b.SizeOf("\"line-break-chars\""), 0)
			PhpConvGetUintPropEx(options, &line_len, "line-length", b.SizeOf("\"line-length\""))
			PhpConvGetBoolPropEx(options, &opt_binary, "binary", b.SizeOf("\"binary\""))
			PhpConvGetBoolPropEx(options, &opt_force_encode_first, "force-encode-first", b.SizeOf("\"force-encode-first\""))
			if line_len < 4 {
				if lbchars != nil {
					zend.Pefree(lbchars, 0)
				}
				lbchars = nil
			} else {
				if lbchars == nil {
					lbchars = zend.Pestrdup("\r\n")
					lbchars_len = 2
				}
			}
			if opt_binary != 0 {
				opts |= PHP_CONV_QPRINT_OPT_BINARY
			} else {
				opts |= 0
			}
			if opt_force_encode_first != 0 {
				opts |= PHP_CONV_QPRINT_OPT_FORCE_ENCODE_FIRST
			} else {
				opts |= 0
			}
		}
		retval = zend.Pemalloc(b.SizeOf("php_conv_qprint_encode"))
		if lbchars != nil {
			if PhpConvQprintEncodeCtor((*PhpConvQprintEncode)(retval), line_len, lbchars, lbchars_len, 1, opts, persistent) != 0 {
				zend.Pefree(lbchars, 0)
				goto out_failure
			}
			zend.Pefree(lbchars, 0)
		} else {
			if PhpConvQprintEncodeCtor((*PhpConvQprintEncode)(retval), 0, nil, 0, 0, opts, persistent) != 0 {
				goto out_failure
			}
		}
	case PHP_CONV_QPRINT_DECODE:
		var lbchars *byte = nil
		var lbchars_len int
		if options != nil {

			/* If line-break-chars are not specified, filter will attempt to detect line endings (\r, \n, or \r\n) */

			PhpConvGetStringPropEx(options, &lbchars, &lbchars_len, "line-break-chars", b.SizeOf("\"line-break-chars\""), 0)

			/* If line-break-chars are not specified, filter will attempt to detect line endings (\r, \n, or \r\n) */

		}
		retval = zend.Pemalloc(b.SizeOf("php_conv_qprint_decode"))
		if lbchars != nil {
			if PhpConvQprintDecodeCtor((*PhpConvQprintDecode)(retval), lbchars, lbchars_len, 1, persistent) != 0 {
				zend.Pefree(lbchars, 0)
				goto out_failure
			}
			zend.Pefree(lbchars, 0)
		} else {
			if PhpConvQprintDecodeCtor((*PhpConvQprintDecode)(retval), nil, 0, 0, persistent) != 0 {
				goto out_failure
			}
		}
	default:
		retval = nil
	}
	return retval
out_failure:
	if retval != nil {
		zend.Pefree(retval, persistent)
	}
	return nil
}
func PhpConvertFilterCtor(inst *PhpConvertFilter, conv_mode int, conv_opts *types.Array, filtername *byte, persistent int) int {
	inst.SetPersistent(persistent)
	inst.SetFiltername(zend.Pestrdup(filtername))
	inst.SetStubLen(0)
	if b.Assign(&(inst.GetCd()), PhpConvOpen(conv_mode, conv_opts, persistent)) == nil {
		goto out_failure
	}
	return types.SUCCESS
out_failure:
	if inst.GetCd() != nil {
		PhpConvDtor(inst.GetCd())
		zend.Pefree(inst.GetCd(), persistent)
	}
	if inst.GetFiltername() != nil {
		zend.Pefree(inst.GetFiltername(), persistent)
	}
	return types.FAILURE
}
func PhpConvertFilterDtor(inst *PhpConvertFilter) {
	if inst.GetCd() != nil {
		PhpConvDtor(inst.GetCd())
		zend.Pefree(inst.GetCd(), inst.GetPersistent())
	}
	if inst.GetFiltername() != nil {
		zend.Pefree(inst.GetFiltername(), inst.GetPersistent())
	}
}
func StrfilterConvertAppendBucket(
	inst *PhpConvertFilter,
	stream *core.PhpStream,
	filter *core.PhpStreamFilter,
	buckets_out *streams.PhpStreamBucketBrigade,
	ps *byte,
	buf_len int,
	consumed *int,
	persistent int,
) int {
	var err PhpConvErrT
	var new_bucket *streams.PhpStreamBucket
	var out_buf *byte = nil
	var out_buf_size int
	var pd *byte
	var pt *byte
	var ocnt int
	var icnt int
	var tcnt int
	var initial_out_buf_size int
	if ps == nil {
		initial_out_buf_size = 64
		icnt = 1
	} else {
		initial_out_buf_size = buf_len
		icnt = buf_len
	}
	ocnt = initial_out_buf_size
	out_buf_size = ocnt
	out_buf = zend.Pemalloc(out_buf_size)
	pd = out_buf
	if inst.GetStubLen() > 0 {
		pt = inst.GetStub()
		tcnt = inst.GetStubLen()
		for tcnt > 0 {
			err = PhpConvConvert(inst.GetCd(), &pt, &tcnt, &pd, &ocnt)
			switch err {
			case PHP_CONV_ERR_INVALID_SEQ:
				core.PhpErrorDocref(nil, faults.E_WARNING, "stream filter (%s): invalid byte sequence", inst.GetFiltername())
				goto out_failure
			case PHP_CONV_ERR_MORE:
				if ps != nil {
					if icnt > 0 {
						if inst.GetStubLen() >= b.SizeOf("inst -> stub") {
							core.PhpErrorDocref(nil, faults.E_WARNING, "stream filter (%s): insufficient buffer", inst.GetFiltername())
							goto out_failure
						}
						inst.GetStub()[b.PostInc(&(inst.GetStubLen()))] = *(b.PostInc(&ps))
						icnt--
						pt = inst.GetStub()
						tcnt = inst.GetStubLen()
					} else {
						tcnt = 0
						break
					}
				}
			case PHP_CONV_ERR_UNEXPECTED_EOS:
				core.PhpErrorDocref(nil, faults.E_WARNING, "stream filter (%s): unexpected end of stream", inst.GetFiltername())
				goto out_failure
			case PHP_CONV_ERR_TOO_BIG:
				var new_out_buf *byte
				var new_out_buf_size int
				new_out_buf_size = out_buf_size << 1
				if new_out_buf_size < out_buf_size {

					/* whoa! no bigger buckets are sold anywhere... */

					if nil == b.Assign(&new_bucket, streams.PhpStreamBucketNew(stream, out_buf, out_buf_size-ocnt, 1, persistent)) {
						goto out_failure
					}
					streams.PhpStreamBucketAppend(buckets_out, new_bucket)
					ocnt = initial_out_buf_size
					out_buf_size = ocnt
					out_buf = zend.Pemalloc(out_buf_size)
					pd = out_buf
				} else {
					new_out_buf = zend.Perealloc(out_buf, new_out_buf_size)
					pd = new_out_buf + (pd - out_buf)
					ocnt += new_out_buf_size - out_buf_size
					out_buf = new_out_buf
					out_buf_size = new_out_buf_size
				}
			case PHP_CONV_ERR_UNKNOWN:
				core.PhpErrorDocref(nil, faults.E_WARNING, "stream filter (%s): unknown error", inst.GetFiltername())
				goto out_failure
			default:

			}
		}
		memmove(inst.GetStub(), pt, tcnt)
		inst.SetStubLen(tcnt)
	}
	for icnt > 0 {
		if ps == nil {
			err = PhpConvConvert(inst.GetCd(), nil, nil, &pd, &ocnt)
		} else {
			err = PhpConvConvert(inst.GetCd(), &ps, &icnt, &pd, &ocnt)
		}
		switch err {
		case PHP_CONV_ERR_INVALID_SEQ:
			core.PhpErrorDocref(nil, faults.E_WARNING, "stream filter (%s): invalid byte sequence", inst.GetFiltername())
			goto out_failure
		case PHP_CONV_ERR_MORE:
			if ps != nil {
				if icnt > b.SizeOf("inst -> stub") {
					core.PhpErrorDocref(nil, faults.E_WARNING, "stream filter (%s): insufficient buffer", inst.GetFiltername())
					goto out_failure
				}
				memcpy(inst.GetStub(), ps, icnt)
				inst.SetStubLen(icnt)
				ps += icnt
				icnt = 0
			} else {
				core.PhpErrorDocref(nil, faults.E_WARNING, "stream filter (%s): unexpected octet values", inst.GetFiltername())
				goto out_failure
			}
		case PHP_CONV_ERR_TOO_BIG:
			var new_out_buf *byte
			var new_out_buf_size int
			new_out_buf_size = out_buf_size << 1
			if new_out_buf_size < out_buf_size {

				/* whoa! no bigger buckets are sold anywhere... */

				if nil == b.Assign(&new_bucket, streams.PhpStreamBucketNew(stream, out_buf, out_buf_size-ocnt, 1, persistent)) {
					goto out_failure
				}
				streams.PhpStreamBucketAppend(buckets_out, new_bucket)
				ocnt = initial_out_buf_size
				out_buf_size = ocnt
				out_buf = zend.Pemalloc(out_buf_size)
				pd = out_buf
			} else {
				new_out_buf = zend.Perealloc(out_buf, new_out_buf_size)
				pd = new_out_buf + (pd - out_buf)
				ocnt += new_out_buf_size - out_buf_size
				out_buf = new_out_buf
				out_buf_size = new_out_buf_size
			}
		case PHP_CONV_ERR_UNKNOWN:
			core.PhpErrorDocref(nil, faults.E_WARNING, "stream filter (%s): unknown error", inst.GetFiltername())
			goto out_failure
		default:
			if ps == nil {
				icnt = 0
			}
		}
	}
	if out_buf_size > ocnt {
		if nil == b.Assign(&new_bucket, streams.PhpStreamBucketNew(stream, out_buf, out_buf_size-ocnt, 1, persistent)) {
			goto out_failure
		}
		streams.PhpStreamBucketAppend(buckets_out, new_bucket)
	} else {
		zend.Pefree(out_buf, persistent)
	}
	*consumed += buf_len - icnt
	return types.SUCCESS
out_failure:
	zend.Pefree(out_buf, persistent)
	return types.FAILURE
}
func StrfilterConvertFilter(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *streams.PhpStreamBucketBrigade,
	buckets_out *streams.PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket = nil
	var consumed int = 0
	var inst *PhpConvertFilter = (*PhpConvertFilter)(thisfilter.GetAbstract().Ptr())
	for buckets_in.GetHead() != nil {
		bucket = buckets_in.GetHead()
		streams.PhpStreamBucketUnlink(bucket)
		if StrfilterConvertAppendBucket(inst, stream, thisfilter, buckets_out, bucket.GetBuf(), bucket.GetBuflen(), &consumed, stream.GetIsPersistent()) != types.SUCCESS {
			goto out_failure
		}
		streams.PhpStreamBucketDelref(bucket)
	}
	if flags != streams.PSFS_FLAG_NORMAL {
		if StrfilterConvertAppendBucket(inst, stream, thisfilter, buckets_out, nil, 0, &consumed, stream.GetIsPersistent()) != types.SUCCESS {
			goto out_failure
		}
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
out_failure:
	if bucket != nil {
		streams.PhpStreamBucketDelref(bucket)
	}
	return streams.PSFS_ERR_FATAL
}
func StrfilterConvertDtor(thisfilter *core.PhpStreamFilter) {
	b.Assert(thisfilter.GetAbstract().Ptr() != nil)
	PhpConvertFilterDtor((*PhpConvertFilter)(thisfilter.GetAbstract().Ptr()))
	zend.Pefree(thisfilter.GetAbstract().Ptr(), (*PhpConvertFilter)(types.Z_PTR(thisfilter.GetAbstract())).GetPersistent())
}
func StrfilterConvertCreate(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter {
	var inst *PhpConvertFilter
	var retval *core.PhpStreamFilter = nil
	var dot *byte
	var conv_mode int = 0
	if filterparams != nil && filterparams.GetType() != types.IS_ARRAY {
		core.PhpErrorDocref(nil, faults.E_WARNING, "stream filter (%s): invalid filter parameter", filtername)
		return nil
	}
	if b.Assign(&dot, strchr(filtername, '.')) == nil {
		return nil
	}
	dot++
	inst = zend.Pemalloc(b.SizeOf("php_convert_filter"))
	if strcasecmp(dot, "base64-encode") == 0 {
		conv_mode = PHP_CONV_BASE64_ENCODE
	} else if strcasecmp(dot, "base64-decode") == 0 {
		conv_mode = PHP_CONV_BASE64_DECODE
	} else if strcasecmp(dot, "quoted-printable-encode") == 0 {
		conv_mode = PHP_CONV_QPRINT_ENCODE
	} else if strcasecmp(dot, "quoted-printable-decode") == 0 {
		conv_mode = PHP_CONV_QPRINT_DECODE
	}
	if PhpConvertFilterCtor(inst, conv_mode, b.CondF1(filterparams != nil, func() *types.Array { return filterparams.Array() }, nil), filtername, persistent) != types.SUCCESS {
		goto out
	}
	retval = streams.PhpStreamFilterAlloc(&StrfilterConvertOps, inst, persistent)
out:
	if retval == nil {
		zend.Pefree(inst, persistent)
	}
	return retval
}
func ConsumedFilterFilter(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *streams.PhpStreamBucketBrigade,
	buckets_out *streams.PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) streams.PhpStreamFilterStatusT {
	var data *PhpConsumedFilterData = (*PhpConsumedFilterData)(thisfilter.GetAbstract().Ptr())
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	if data.GetOffset() == ^0 {
		data.SetOffset(stream.GetPosition())
	}
	for b.Assign(&bucket, buckets_in.GetHead()) != nil {
		streams.PhpStreamBucketUnlink(bucket)
		consumed += bucket.GetBuflen()
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	if (flags & streams.PSFS_FLAG_FLUSH_CLOSE) != 0 {
		core.PhpStreamSeek(stream, data.GetOffset()+data.GetConsumed(), r.SEEK_SET)
	}
	data.SetConsumed(data.GetConsumed() + consumed)
	return streams.PSFS_PASS_ON
}
func ConsumedFilterDtor(thisfilter *core.PhpStreamFilter) {
	if thisfilter != nil && thisfilter.GetAbstract().Ptr() {
		var data *PhpConsumedFilterData = (*PhpConsumedFilterData)(thisfilter.GetAbstract().Ptr())
		zend.Pefree(data, data.GetPersistent())
	}
}
func ConsumedFilterCreate(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter {
	var fops *streams.PhpStreamFilterOps = nil
	var data *PhpConsumedFilterData
	if strcasecmp(filtername, "consumed") {
		return nil
	}

	/* Create this filter */

	data = zend.Pecalloc(1, b.SizeOf("php_consumed_filter_data"))
	data.SetPersistent(persistent)
	data.SetConsumed(0)
	data.SetOffset(^0)
	fops = &ConsumedFilterOps
	return streams.PhpStreamFilterAlloc(fops, data, persistent)
}
func PhpDechunk(buf *byte, len_ int, data *PhpChunkedFilterData) int {
	var p *byte = buf
	var end *byte = p + len_
	var out *byte = buf
	var out_len int = 0
	for p < end {
		switch data.GetState() {
		case CHUNK_SIZE_START:
			data.SetChunkSize(0)
			fallthrough
		case CHUNK_SIZE:
			for p < end {
				if (*p) >= '0' && (*p) <= '9' {
					data.SetChunkSize(data.GetChunkSize()*16 + ((*p) - '0'))
				} else if (*p) >= 'A' && (*p) <= 'F' {
					data.SetChunkSize(data.GetChunkSize()*16 + ((*p) - 'A' + 10))
				} else if (*p) >= 'a' && (*p) <= 'f' {
					data.SetChunkSize(data.GetChunkSize()*16 + ((*p) - 'a' + 10))
				} else if data.GetState() == CHUNK_SIZE_START {
					data.SetState(CHUNK_ERROR)
					break
				} else {
					data.SetState(CHUNK_SIZE_EXT)
					break
				}
				data.SetState(CHUNK_SIZE)
				p++
			}
			if data.GetState() == CHUNK_ERROR {
				continue
			} else if p == end {
				return out_len
			}
			fallthrough
		case CHUNK_SIZE_EXT:

			/* skip extension */

			for p < end && (*p) != '\r' && (*p) != '\n' {
				p++
			}
			if p == end {
				return out_len
			}
			fallthrough
		case CHUNK_SIZE_CR:
			if (*p) == '\r' {
				p++
				if p == end {
					data.SetState(CHUNK_SIZE_LF)
					return out_len
				}
			}
			fallthrough
		case CHUNK_SIZE_LF:
			if (*p) == '\n' {
				p++
				if data.GetChunkSize() == 0 {

					/* last chunk */

					data.SetState(CHUNK_TRAILER)
					continue
				} else if p == end {
					data.SetState(CHUNK_BODY)
					return out_len
				}
			} else {
				data.SetState(CHUNK_ERROR)
				continue
			}
			fallthrough
		case CHUNK_BODY:
			if size_t(end-p) >= data.GetChunkSize() {
				if p != out {
					memmove(out, p, data.GetChunkSize())
				}
				out += data.GetChunkSize()
				out_len += data.GetChunkSize()
				p += data.GetChunkSize()
				if p == end {
					data.SetState(CHUNK_BODY_CR)
					return out_len
				}
			} else {
				if p != out {
					memmove(out, p, end-p)
				}
				data.SetChunkSize(data.GetChunkSize() - end - p)
				data.SetState(CHUNK_BODY)
				out_len += end - p
				return out_len
			}
			fallthrough
		case CHUNK_BODY_CR:
			if (*p) == '\r' {
				p++
				if p == end {
					data.SetState(CHUNK_BODY_LF)
					return out_len
				}
			}
			fallthrough
		case CHUNK_BODY_LF:
			if (*p) == '\n' {
				p++
				data.SetState(CHUNK_SIZE_START)
				continue
			} else {
				data.SetState(CHUNK_ERROR)
				continue
			}
			fallthrough
		case CHUNK_TRAILER:

			/* ignore trailer */

			p = end
			continue
		case CHUNK_ERROR:
			if p != out {
				memmove(out, p, end-p)
			}
			out_len += end - p
			return out_len
		}
	}
	return out_len
}
func PhpChunkedFilter(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *streams.PhpStreamBucketBrigade,
	buckets_out *streams.PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	var data *PhpChunkedFilterData = (*PhpChunkedFilterData)(thisfilter.GetAbstract().Ptr())
	for buckets_in.GetHead() != nil {
		bucket = streams.PhpStreamBucketMakeWriteable(buckets_in.GetHead())
		consumed += bucket.GetBuflen()
		bucket.SetBuflen(PhpDechunk(bucket.GetBuf(), bucket.GetBuflen(), data))
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
}
func PhpChunkedDtor(thisfilter *core.PhpStreamFilter) {
	if thisfilter != nil && thisfilter.GetAbstract().Ptr() {
		var data *PhpChunkedFilterData = (*PhpChunkedFilterData)(thisfilter.GetAbstract().Ptr())
		zend.Pefree(data, data.GetPersistent())
	}
}
func ChunkedFilterCreate(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter {
	var fops *streams.PhpStreamFilterOps = nil
	var data *PhpChunkedFilterData
	if strcasecmp(filtername, "dechunk") {
		return nil
	}

	/* Create this filter */

	data = (*PhpChunkedFilterData)(zend.Pecalloc(1, b.SizeOf("php_chunked_filter_data")))
	data.SetState(CHUNK_SIZE_START)
	data.SetChunkSize(0)
	data.SetPersistent(persistent)
	fops = &ChunkedFilterOps
	return streams.PhpStreamFilterAlloc(fops, data, persistent)
}
func ZmStartupStandardFilters(type_ int, module_number int) int {
	var i int
	for i = 0; StandardFilters[i].ops != nil; i++ {
		if types.FAILURE == streams.PhpStreamFilterRegisterFactory(StandardFilters[i].ops.GetLabel(), StandardFilters[i].factory) {
			return types.FAILURE
		}
	}
	return types.SUCCESS
}
func ZmShutdownStandardFilters(type_ int, module_number int) int {
	var i int
	for i = 0; StandardFilters[i].ops != nil; i++ {
		streams.PhpStreamFilterUnregisterFactory(StandardFilters[i].ops.GetLabel())
	}
	return types.SUCCESS
}
