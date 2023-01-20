// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/filters.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors:                                                             |
   | Wez Furlong (wez@thebrainroom.com)                                   |
   | Sara Golemon (pollita@php.net)                                       |
   | Moriyoshi Koizumi (moriyoshi@php.net)                                |
   | Marcus Boerger (helly@php.net)                                       |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_globals.h"

// # include "ext/standard/basic_functions.h"

// # include "ext/standard/file.h"

// # include "ext/standard/php_string.h"

// # include "zend_smart_str.h"

/* {{{ rot13 stream filter implementation */

var Rot13From []byte = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var Rot13To []byte = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"

func StrfilterRot13Filter(stream *core.PhpStream, thisfilter *core.PhpStreamFilter, buckets_in *streams.PhpStreamBucketBrigade, buckets_out *streams.PhpStreamBucketBrigade, bytes_consumed *int, flags int) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	for buckets_in.head != nil {
		bucket = streams.PhpStreamBucketMakeWriteable(buckets_in.head)
		PhpStrtr(bucket.buf, bucket.buflen, Rot13From, Rot13To, 52)
		consumed += bucket.buflen
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
}

var StrfilterRot13Ops streams.PhpStreamFilterOps = streams.PhpStreamFilterOps{StrfilterRot13Filter, nil, "string.rot13"}

func StrfilterRot13Create(filtername *byte, filterparams *zend.Zval, persistent uint8) *core.PhpStreamFilter {
	return streams._phpStreamFilterAlloc(&StrfilterRot13Ops, nil, persistent)
}

var StrfilterRot13Factory streams.PhpStreamFilterFactory = streams.PhpStreamFilterFactory{StrfilterRot13Create}

/* }}} */

var Lowercase []byte = "abcdefghijklmnopqrstuvwxyz"
var Uppercase []byte = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func StrfilterToupperFilter(stream *core.PhpStream, thisfilter *core.PhpStreamFilter, buckets_in *streams.PhpStreamBucketBrigade, buckets_out *streams.PhpStreamBucketBrigade, bytes_consumed *int, flags int) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	for buckets_in.head != nil {
		bucket = streams.PhpStreamBucketMakeWriteable(buckets_in.head)
		PhpStrtr(bucket.buf, bucket.buflen, Lowercase, Uppercase, 26)
		consumed += bucket.buflen
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
}
func StrfilterTolowerFilter(stream *core.PhpStream, thisfilter *core.PhpStreamFilter, buckets_in *streams.PhpStreamBucketBrigade, buckets_out *streams.PhpStreamBucketBrigade, bytes_consumed *int, flags int) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	for buckets_in.head != nil {
		bucket = streams.PhpStreamBucketMakeWriteable(buckets_in.head)
		PhpStrtr(bucket.buf, bucket.buflen, Uppercase, Lowercase, 26)
		consumed += bucket.buflen
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
}

var StrfilterToupperOps streams.PhpStreamFilterOps = streams.PhpStreamFilterOps{StrfilterToupperFilter, nil, "string.toupper"}
var StrfilterTolowerOps streams.PhpStreamFilterOps = streams.PhpStreamFilterOps{StrfilterTolowerFilter, nil, "string.tolower"}

func StrfilterToupperCreate(filtername *byte, filterparams *zend.Zval, persistent uint8) *core.PhpStreamFilter {
	return streams._phpStreamFilterAlloc(&StrfilterToupperOps, nil, persistent)
}
func StrfilterTolowerCreate(filtername *byte, filterparams *zend.Zval, persistent uint8) *core.PhpStreamFilter {
	return streams._phpStreamFilterAlloc(&StrfilterTolowerOps, nil, persistent)
}

var StrfilterToupperFactory streams.PhpStreamFilterFactory = streams.PhpStreamFilterFactory{StrfilterToupperCreate}
var StrfilterTolowerFactory streams.PhpStreamFilterFactory = streams.PhpStreamFilterFactory{StrfilterTolowerCreate}

/* }}} */

// @type PhpStripTagsFilter struct

func PhpStripTagsFilterCtor(inst *PhpStripTagsFilter, allowed_tags *zend.ZendString, persistent int) int {
	if allowed_tags != nil {
		if nil == g.Assign(&(inst.GetAllowedTags()), g.CondF(persistent != 0, func() any { return zend.__zendMalloc(allowed_tags.len_ + 1) }, func() any { return zend._emalloc(allowed_tags.len_ + 1) })) {
			return zend.FAILURE
		}
		memcpy((*byte)(inst.GetAllowedTags()), allowed_tags.val, allowed_tags.len_+1)
		inst.SetAllowedTagsLen(int(allowed_tags.len_))
	} else {
		inst.SetAllowedTags(nil)
	}
	inst.SetState(0)
	inst.SetPersistent(persistent)
	return zend.SUCCESS
}
func PhpStripTagsFilterDtor(inst *PhpStripTagsFilter) {
	if inst.GetAllowedTags() != nil {
		g.CondF(inst.GetPersistent() != 0, func() { return zend.Free(any(inst.GetAllowedTags())) }, func() { return zend._efree(any(inst.GetAllowedTags())) })
	}
}
func StrfilterStripTagsFilter(stream *core.PhpStream, thisfilter *core.PhpStreamFilter, buckets_in *streams.PhpStreamBucketBrigade, buckets_out *streams.PhpStreamBucketBrigade, bytes_consumed *int, flags int) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	var inst *PhpStripTagsFilter = (*PhpStripTagsFilter)(thisfilter.abstract.value.ptr)
	for buckets_in.head != nil {
		bucket = streams.PhpStreamBucketMakeWriteable(buckets_in.head)
		consumed = bucket.buflen
		bucket.buflen = PhpStripTags(bucket.buf, bucket.buflen, &(inst.GetState()), inst.GetAllowedTags(), inst.GetAllowedTagsLen())
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
}
func StrfilterStripTagsDtor(thisfilter *core.PhpStreamFilter) {
	assert(thisfilter.abstract.value.ptr != nil)
	PhpStripTagsFilterDtor((*PhpStripTagsFilter)(thisfilter.abstract.value.ptr))
	g.CondF((*PhpStripTagsFilter)(thisfilter.abstract.value.ptr).GetPersistent() != 0, func() { return zend.Free(thisfilter.abstract.value.ptr) }, func() { return zend._efree(thisfilter.abstract.value.ptr) })
}

var StrfilterStripTagsOps streams.PhpStreamFilterOps = streams.PhpStreamFilterOps{StrfilterStripTagsFilter, StrfilterStripTagsDtor, "string.strip_tags"}

func StrfilterStripTagsCreate(filtername *byte, filterparams *zend.Zval, persistent uint8) *core.PhpStreamFilter {
	var inst *PhpStripTagsFilter
	var filter *core.PhpStreamFilter = nil
	var allowed_tags *zend.ZendString = nil
	core.PhpErrorDocref(nil, 1<<13, "The string.strip_tags filter is deprecated")
	if filterparams != nil {
		if filterparams.u1.v.type_ == 7 {
			var tags_ss zend.SmartStr = zend.SmartStr{0}
			var tmp *zend.Zval
			for {
				var __ht *zend.HashTable = filterparams.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					tmp = _z
					if tmp.u1.v.type_ != 6 {
						if tmp.u1.v.type_ != 6 {
							zend._convertToString(tmp)
						}
					}
					zend.SmartStrAppendcEx(&tags_ss, '<', 0)
					zend.SmartStrAppendEx(&tags_ss, tmp.value.str, 0)
					zend.SmartStrAppendcEx(&tags_ss, '>', 0)
				}
				break
			}
			zend.SmartStr0(&tags_ss)
			allowed_tags = tags_ss.s
		} else {
			allowed_tags = zend.ZvalGetString(filterparams)
		}

		/* Exception during string conversion. */

		if zend.EG.exception != nil {
			if allowed_tags != nil {
				zend.ZendStringRelease(allowed_tags)
			}
			return nil
		}

		/* Exception during string conversion. */

	}
	if persistent != 0 {
		inst = zend.__zendMalloc(g.SizeOf("php_strip_tags_filter"))
	} else {
		inst = zend._emalloc(g.SizeOf("php_strip_tags_filter"))
	}
	if PhpStripTagsFilterCtor(inst, allowed_tags, persistent) == zend.SUCCESS {
		filter = streams._phpStreamFilterAlloc(&StrfilterStripTagsOps, inst, persistent)
	} else {
		g.CondF(persistent != 0, func() { return zend.Free(inst) }, func() { return zend._efree(inst) })
	}
	if allowed_tags != nil {
		zend.ZendStringRelease(allowed_tags)
	}
	return filter
}

var StrfilterStripTagsFactory streams.PhpStreamFilterFactory = streams.PhpStreamFilterFactory{StrfilterStripTagsCreate}

/* }}} */

type PhpConvErrT = int

const (
	PHP_CONV_ERR_SUCCESS PhpConvErrT = zend.SUCCESS
	PHP_CONV_ERR_UNKNOWN
	PHP_CONV_ERR_TOO_BIG
	PHP_CONV_ERR_INVALID_SEQ
	PHP_CONV_ERR_UNEXPECTED_EOS
	PHP_CONV_ERR_EXISTS
	PHP_CONV_ERR_MORE
	PHP_CONV_ERR_ALLOC
	PHP_CONV_ERR_NOT_FOUND
)

type PhpConvConvertFunc func(*PhpConv, **byte, *int, **byte, *int) PhpConvErrT
type PhpConvDtorFunc func(*PhpConv)

// @type PhpConv struct

// #define php_conv_convert(a,b,c,d,e) ( ( php_conv * ) ( a ) ) -> convert_op ( ( php_conv * ) ( a ) , ( b ) , ( c ) , ( d ) , ( e ) )

// #define php_conv_dtor(a) ( ( php_conv * ) a ) -> dtor ( ( a ) )

/* {{{ php_conv_base64_encode */

// @type PhpConvBase64Encode struct

var B64TblEnc []uint8 = []uint8{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/'}

func PhpConvBase64EncodeCtor(inst *PhpConvBase64Encode, line_len uint, lbchars *byte, lbchars_len int, lbchars_dup int, persistent int) PhpConvErrT {
	inst.GetSuper().SetConvertOp(PhpConvConvertFunc(PhpConvBase64EncodeConvert))
	inst.GetSuper().SetDtor(PhpConvDtorFunc(PhpConvBase64EncodeDtor))
	inst.SetEremLen(0)
	inst.SetLineCcnt(line_len)
	inst.SetLineLen(line_len)
	if lbchars != nil {
		if lbchars_dup != 0 {
			if persistent != 0 {
				inst.SetLbchars(strdup(lbchars))
			} else {
				inst.SetLbchars(zend._estrdup(lbchars))
			}
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
	assert(inst != nil)
	if inst.GetLbcharsDup() != 0 && inst.GetLbchars() != nil {
		g.CondF(inst.GetPersistent() != 0, func() { return zend.Free(any(inst.GetLbchars())) }, func() { return zend._efree(any(inst.GetLbchars())) })
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

		break
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
		*(g.PostInc(&pd)) = B64TblEnc[inst.GetErem()[0]>>2]
		*(g.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[0]<<4)]
		*(g.PostInc(&pd)) = '='
		*(g.PostInc(&pd)) = '='
		inst.SetEremLen(0)
		ocnt -= 4
		line_ccnt -= 4
		break
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
		*(g.PostInc(&pd)) = B64TblEnc[inst.GetErem()[0]>>2]
		*(g.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[0]<<4|inst.GetErem()[1]>>4)]
		*(g.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[1]<<2)]
		*(g.PostInc(&pd)) = '='
		inst.SetEremLen(0)
		ocnt -= 4
		line_ccnt -= 4
		break
	default:

		/* should not happen... */

		err = PHP_CONV_ERR_UNKNOWN
		break
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
			*(g.PostInc(&pd)) = B64TblEnc[inst.GetErem()[0]>>2]
			*(g.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[0]<<4|ps[0]>>4)]
			*(g.PostInc(&pd)) = B64TblEnc[uint8(ps[0]<<2|ps[1]>>6)]
			*(g.PostInc(&pd)) = B64TblEnc[ps[1]]
			ocnt -= 4
			ps += 2
			icnt -= 2
			inst.SetEremLen(0)
			line_ccnt -= 4
		}
		break
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
			*(g.PostInc(&pd)) = B64TblEnc[inst.GetErem()[0]>>2]
			*(g.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[0]<<4|inst.GetErem()[1]>>4)]
			*(g.PostInc(&pd)) = B64TblEnc[uint8(inst.GetErem()[1]<<2|ps[0]>>6)]
			*(g.PostInc(&pd)) = B64TblEnc[ps[0]]
			ocnt -= 4
			ps += 1
			icnt -= 1
			inst.SetEremLen(0)
			line_ccnt -= 4
		}
		break
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
		*(g.PostInc(&pd)) = B64TblEnc[ps[0]>>2]
		*(g.PostInc(&pd)) = B64TblEnc[uint8(ps[0]<<4|ps[1]>>4)]
		*(g.PostInc(&pd)) = B64TblEnc[uint8(ps[1]<<2|ps[2]>>6)]
		*(g.PostInc(&pd)) = B64TblEnc[ps[2]]
		ps += 3
		icnt -= 3
		ocnt -= 4
		line_ccnt -= 4
	}
	for ; icnt > 0; icnt-- {
		inst.GetErem()[g.PostInc(&(inst.GetEremLen()))] = *(g.PostInc(&ps))
	}
out:
	*in_pp = (*byte)(ps)
	*in_left_p = icnt
	*out_pp = (*byte)(pd)
	*out_left_p = ocnt
	inst.SetLineCcnt(line_ccnt)
	return err
}

/* }}} */

// @type PhpConvBase64Decode struct

var B64TblDec []uint = []uint{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 62, 64, 64, 64, 63, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 64, 64, 64, 128, 64, 64, 64, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 64, 64, 64, 64, 64, 64, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64}

func PhpConvBase64DecodeCtor(inst *PhpConvBase64Decode) int {
	inst.GetSuper().SetConvertOp(PhpConvConvertFunc(PhpConvBase64DecodeConvert))
	inst.GetSuper().SetDtor(PhpConvDtorFunc(PhpConvBase64DecodeDtor))
	inst.SetUrem(0)
	inst.SetUremNbits(0)
	inst.SetUstat(0)
	inst.SetEos(0)
	return zend.SUCCESS
}
func PhpConvBase64DecodeDtor(inst *PhpConvBase64Decode) {}

// #define bmask(a) ( 0xffff >> ( 16 - a ) )

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
			i = B64TblDec[uint(*(g.PostInc(&ps)))]
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
			*(g.PostInc(&pd)) = pack
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

/* }}} */

// @type PhpConvQprintEncode struct

// #define PHP_CONV_QPRINT_OPT_BINARY       0x00000001

// #define PHP_CONV_QPRINT_OPT_FORCE_ENCODE_FIRST       0x00000002

func PhpConvQprintEncodeDtor(inst *PhpConvQprintEncode) {
	assert(inst != nil)
	if inst.GetLbcharsDup() != 0 && inst.GetLbchars() != nil {
		g.CondF(inst.GetPersistent() != 0, func() { return zend.Free(any(inst.GetLbchars())) }, func() { return zend._efree(any(inst.GetLbchars())) })
	}
}

// #define NEXT_CHAR(ps,icnt,lb_ptr,lb_cnt,lbchars) ( ( lb_ptr ) < ( lb_cnt ) ? ( lbchars ) [ ( lb_ptr ) ] : * ( ps ) )

// #define CONSUME_CHAR(ps,icnt,lb_ptr,lb_cnt) if ( ( lb_ptr ) < ( lb_cnt ) ) { ( lb_ptr ) ++ ; } else { ( lb_cnt ) = ( lb_ptr ) = 0 ; -- ( icnt ) ; ( ps ) ++ ; }

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
		if (opts&0x1) == 0 && inst.GetLbchars() != nil && inst.GetLbcharsLen() > 0 {

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
						*(g.PostInc(&pd)) = inst.GetLbchars()[i]
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
		if (opts&0x1) == 0 && trail_ws == 0 && (c == '\t' || c == ' ') {
			if line_ccnt < 2 && inst.GetLbchars() != nil {
				if ocnt < inst.GetLbcharsLen()+1 {
					err = PHP_CONV_ERR_TOO_BIG
					break
				}
				*(g.PostInc(&pd)) = '='
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
					*(g.PostInc(&pd)) = c
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
		} else if ((opts&0x2) == 0 || line_ccnt < inst.GetLineLen()) && (c >= 33 && c <= 60 || c >= 62 && c <= 126) {
			if line_ccnt < 2 && inst.GetLbchars() != nil {
				if ocnt < inst.GetLbcharsLen()+1 {
					err = PHP_CONV_ERR_TOO_BIG
					break
				}
				*(g.PostInc(&pd)) = '='
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
			*(g.PostInc(&pd)) = c
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
				*(g.PostInc(&pd)) = '='
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
			*(g.PostInc(&pd)) = '='
			*(g.PostInc(&pd)) = qp_digits[c>>4]
			*(g.PostInc(&pd)) = qp_digits[c&0xf]
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
func PhpConvQprintEncodeCtor(inst *PhpConvQprintEncode, line_len uint, lbchars *byte, lbchars_len int, lbchars_dup int, opts int, persistent int) PhpConvErrT {
	if line_len < 4 && lbchars != nil {
		return PHP_CONV_ERR_TOO_BIG
	}
	inst.GetSuper().SetConvertOp(PhpConvConvertFunc(PhpConvQprintEncodeConvert))
	inst.GetSuper().SetDtor(PhpConvDtorFunc(PhpConvQprintEncodeDtor))
	inst.SetLineCcnt(line_len)
	inst.SetLineLen(line_len)
	if lbchars != nil {
		if lbchars_dup != 0 {
			if persistent != 0 {
				inst.SetLbchars(strdup(lbchars))
			} else {
				inst.SetLbchars(zend._estrdup(lbchars))
			}
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

/* }}} */

// @type PhpConvQprintDecode struct

func PhpConvQprintDecodeDtor(inst *PhpConvQprintDecode) {
	assert(inst != nil)
	if inst.GetLbcharsDup() != 0 && inst.GetLbchars() != nil {
		g.CondF(inst.GetPersistent() != 0, func() { return zend.Free(any(inst.GetLbchars())) }, func() { return zend._efree(any(inst.GetLbchars())) })
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
				*(g.PostInc(&pd)) = *ps
				ocnt--
			}
			ps++
			icnt--
			break
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
		case 2:
			if icnt == 0 {
				goto out
			}
			if !(isxdigit(int(*ps))) {
				err = PHP_CONV_ERR_INVALID_SEQ
				goto out
			}
			next_char = next_char<<4 | g.Cond((*ps) >= 'A', (*ps)-0x37, (*ps)-0x30)
			scan_stat++
			ps++
			icnt--
			if scan_stat != 3 {
				break
			}
		case 3:
			if ocnt < 1 {
				err = PHP_CONV_ERR_TOO_BIG
				goto out
			}
			*(g.PostInc(&pd)) = next_char
			ocnt--
			scan_stat = 0
			break
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
			break
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
			break
		case 6:
			if lb_ptr < lb_cnt {
				if ocnt < 1 {
					err = PHP_CONV_ERR_TOO_BIG
					goto out
				}
				*(g.PostInc(&pd)) = inst.GetLbchars()[g.PostInc(&lb_ptr)]
				ocnt--
			} else {
				scan_stat = 0
				lb_ptr = 0
				lb_cnt = lb_ptr
			}
			break
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
			if persistent != 0 {
				inst.SetLbchars(strdup(lbchars))
			} else {
				inst.SetLbchars(zend._estrdup(lbchars))
			}
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

/* }}} */

// @type PhpConvertFilter struct

// #define PHP_CONV_BASE64_ENCODE       1

// #define PHP_CONV_BASE64_DECODE       2

// #define PHP_CONV_QPRINT_ENCODE       3

// #define PHP_CONV_QPRINT_DECODE       4

func PhpConvGetStringPropEx(ht *zend.HashTable, pretval **byte, pretval_len *int, field_name string, field_name_len int, persistent int) PhpConvErrT {
	var tmpval *zend.Zval
	*pretval = nil
	*pretval_len = 0
	if g.Assign(&tmpval, zend.ZendHashStrFind((*zend.HashTable)(ht), field_name, field_name_len-1)) != nil {
		var tmp *zend.ZendString
		var str *zend.ZendString = zend.ZvalGetTmpString(tmpval, &tmp)
		if persistent != 0 {
			*pretval = zend.__zendMalloc(str.len_ + 1)
		} else {
			*pretval = zend._emalloc(str.len_ + 1)
		}
		*pretval_len = str.len_
		memcpy(*pretval, str.val, str.len_+1)
		zend.ZendTmpStringRelease(tmp)
	} else {
		return PHP_CONV_ERR_NOT_FOUND
	}
	return PHP_CONV_ERR_SUCCESS
}
func PhpConvGetUlongPropEx(ht *zend.HashTable, pretval *zend.ZendUlong, field_name string, field_name_len int) PhpConvErrT {
	var tmpval *zend.Zval = zend.ZendHashStrFind((*zend.HashTable)(ht), field_name, field_name_len-1)
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
func PhpConvGetBoolPropEx(ht *zend.HashTable, pretval *int, field_name string, field_name_len int) PhpConvErrT {
	var tmpval *zend.Zval = zend.ZendHashStrFind((*zend.HashTable)(ht), field_name, field_name_len-1)
	if tmpval != nil {
		*pretval = zend.ZendIsTrue(tmpval)
		return PHP_CONV_ERR_SUCCESS
	} else {
		*pretval = 0
		return PHP_CONV_ERR_NOT_FOUND
	}
}

/* XXX this might need an additional fix so it uses size_t, whereby unsigned is quite big so leaving as is for now */

func PhpConvGetUintPropEx(ht *zend.HashTable, pretval *uint, field_name string, field_name_len int) int {
	var l zend.ZendUlong
	var err PhpConvErrT
	*pretval = 0
	if g.Assign(&err, PhpConvGetUlongPropEx(ht, &l, field_name, field_name_len)) == PHP_CONV_ERR_SUCCESS {
		*pretval = uint(l)
	}
	return err
}

// #define GET_STR_PROP(ht,var,var_len,fldname,persistent) php_conv_get_string_prop_ex ( ht , & var , & var_len , fldname , sizeof ( fldname ) , persistent )

// #define GET_INT_PROP(ht,var,fldname) php_conv_get_int_prop_ex ( ht , & var , fldname , sizeof ( fldname ) )

// #define GET_UINT_PROP(ht,var,fldname) php_conv_get_uint_prop_ex ( ht , & var , fldname , sizeof ( fldname ) )

// #define GET_BOOL_PROP(ht,var,fldname) php_conv_get_bool_prop_ex ( ht , & var , fldname , sizeof ( fldname ) )

func PhpConvOpen(conv_mode int, options *zend.HashTable, persistent int) *PhpConv {
	/* FIXME: I'll have to replace this ugly code by something neat
	   (factories?) in the near future. */

	var retval *PhpConv = nil
	switch conv_mode {
	case 1:
		var line_len uint = 0
		var lbchars *byte = nil
		var lbchars_len int
		if options != nil {
			PhpConvGetStringPropEx(options, &lbchars, &lbchars_len, "line-break-chars", g.SizeOf("\"line-break-chars\""), 0)
			PhpConvGetUintPropEx(options, &line_len, "line-length", g.SizeOf("\"line-length\""))
			if line_len < 4 {
				if lbchars != nil {
					g.CondF(false, func() { return zend.Free(lbchars) }, func() { return zend._efree(lbchars) })
				}
				lbchars = nil
			} else {
				if lbchars == nil {
					lbchars = zend._estrdup("\r\n")
					lbchars_len = 2
				}
			}
		}
		if persistent != 0 {
			retval = zend.__zendMalloc(g.SizeOf("php_conv_base64_encode"))
		} else {
			retval = zend._emalloc(g.SizeOf("php_conv_base64_encode"))
		}
		if lbchars != nil {
			if PhpConvBase64EncodeCtor((*PhpConvBase64Encode)(retval), line_len, lbchars, lbchars_len, 1, persistent) != 0 {
				if lbchars != nil {
					g.CondF(false, func() { return zend.Free(lbchars) }, func() { return zend._efree(lbchars) })
				}
				goto out_failure
			}
			g.CondF(false, func() { return zend.Free(lbchars) }, func() { return zend._efree(lbchars) })
		} else {
			if PhpConvBase64EncodeCtor((*PhpConvBase64Encode)(retval), 0, nil, 0, 0, persistent) != 0 {
				goto out_failure
			}
		}
		break
	case 2:
		if persistent != 0 {
			retval = zend.__zendMalloc(g.SizeOf("php_conv_base64_decode"))
		} else {
			retval = zend._emalloc(g.SizeOf("php_conv_base64_decode"))
		}
		if PhpConvBase64DecodeCtor((*PhpConvBase64Decode)(retval)) != 0 {
			goto out_failure
		}
		break
	case 3:
		var line_len uint = 0
		var lbchars *byte = nil
		var lbchars_len int
		var opts int = 0
		if options != nil {
			var opt_binary int = 0
			var opt_force_encode_first int = 0
			PhpConvGetStringPropEx(options, &lbchars, &lbchars_len, "line-break-chars", g.SizeOf("\"line-break-chars\""), 0)
			PhpConvGetUintPropEx(options, &line_len, "line-length", g.SizeOf("\"line-length\""))
			PhpConvGetBoolPropEx(options, &opt_binary, "binary", g.SizeOf("\"binary\""))
			PhpConvGetBoolPropEx(options, &opt_force_encode_first, "force-encode-first", g.SizeOf("\"force-encode-first\""))
			if line_len < 4 {
				if lbchars != nil {
					g.CondF(false, func() { return zend.Free(lbchars) }, func() { return zend._efree(lbchars) })
				}
				lbchars = nil
			} else {
				if lbchars == nil {
					lbchars = zend._estrdup("\r\n")
					lbchars_len = 2
				}
			}
			if opt_binary != 0 {
				opts |= 0x1
			} else {
				opts |= 0
			}
			if opt_force_encode_first != 0 {
				opts |= 0x2
			} else {
				opts |= 0
			}
		}
		if persistent != 0 {
			retval = zend.__zendMalloc(g.SizeOf("php_conv_qprint_encode"))
		} else {
			retval = zend._emalloc(g.SizeOf("php_conv_qprint_encode"))
		}
		if lbchars != nil {
			if PhpConvQprintEncodeCtor((*PhpConvQprintEncode)(retval), line_len, lbchars, lbchars_len, 1, opts, persistent) != 0 {
				g.CondF(false, func() { return zend.Free(lbchars) }, func() { return zend._efree(lbchars) })
				goto out_failure
			}
			g.CondF(false, func() { return zend.Free(lbchars) }, func() { return zend._efree(lbchars) })
		} else {
			if PhpConvQprintEncodeCtor((*PhpConvQprintEncode)(retval), 0, nil, 0, 0, opts, persistent) != 0 {
				goto out_failure
			}
		}
		break
	case 4:
		var lbchars *byte = nil
		var lbchars_len int
		if options != nil {

			/* If line-break-chars are not specified, filter will attempt to detect line endings (\r, \n, or \r\n) */

			PhpConvGetStringPropEx(options, &lbchars, &lbchars_len, "line-break-chars", g.SizeOf("\"line-break-chars\""), 0)

			/* If line-break-chars are not specified, filter will attempt to detect line endings (\r, \n, or \r\n) */

		}
		if persistent != 0 {
			retval = zend.__zendMalloc(g.SizeOf("php_conv_qprint_decode"))
		} else {
			retval = zend._emalloc(g.SizeOf("php_conv_qprint_decode"))
		}
		if lbchars != nil {
			if PhpConvQprintDecodeCtor((*PhpConvQprintDecode)(retval), lbchars, lbchars_len, 1, persistent) != 0 {
				g.CondF(false, func() { return zend.Free(lbchars) }, func() { return zend._efree(lbchars) })
				goto out_failure
			}
			g.CondF(false, func() { return zend.Free(lbchars) }, func() { return zend._efree(lbchars) })
		} else {
			if PhpConvQprintDecodeCtor((*PhpConvQprintDecode)(retval), nil, 0, 0, persistent) != 0 {
				goto out_failure
			}
		}
		break
	default:
		retval = nil
		break
	}
	return retval
out_failure:
	if retval != nil {
		g.CondF(persistent != 0, func() { return zend.Free(retval) }, func() { return zend._efree(retval) })
	}
	return nil
}
func PhpConvertFilterCtor(inst *PhpConvertFilter, conv_mode int, conv_opts *zend.HashTable, filtername *byte, persistent int) int {
	inst.SetPersistent(persistent)
	if persistent != 0 {
		inst.SetFiltername(strdup(filtername))
	} else {
		inst.SetFiltername(zend._estrdup(filtername))
	}
	inst.SetStubLen(0)
	if g.Assign(&(inst.GetCd()), PhpConvOpen(conv_mode, conv_opts, persistent)) == nil {
		goto out_failure
	}
	return zend.SUCCESS
out_failure:
	if inst.GetCd() != nil {
		(*PhpConv)(inst.GetCd()).GetDtor()(inst.GetCd())
		g.CondF(persistent != 0, func() { return zend.Free(inst.GetCd()) }, func() { return zend._efree(inst.GetCd()) })
	}
	if inst.GetFiltername() != nil {
		g.CondF(persistent != 0, func() { return zend.Free(inst.GetFiltername()) }, func() { return zend._efree(inst.GetFiltername()) })
	}
	return zend.FAILURE
}
func PhpConvertFilterDtor(inst *PhpConvertFilter) {
	if inst.GetCd() != nil {
		(*PhpConv)(inst.GetCd()).GetDtor()(inst.GetCd())
		g.CondF(inst.GetPersistent() != 0, func() { return zend.Free(inst.GetCd()) }, func() { return zend._efree(inst.GetCd()) })
	}
	if inst.GetFiltername() != nil {
		g.CondF(inst.GetPersistent() != 0, func() { return zend.Free(inst.GetFiltername()) }, func() { return zend._efree(inst.GetFiltername()) })
	}
}

/* {{{ strfilter_convert_append_bucket */

func StrfilterConvertAppendBucket(inst *PhpConvertFilter, stream *core.PhpStream, filter *core.PhpStreamFilter, buckets_out *streams.PhpStreamBucketBrigade, ps *byte, buf_len int, consumed *int, persistent int) int {
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
	if persistent != 0 {
		out_buf = zend.__zendMalloc(out_buf_size)
	} else {
		out_buf = zend._emalloc(out_buf_size)
	}
	pd = out_buf
	if inst.GetStubLen() > 0 {
		pt = inst.GetStub()
		tcnt = inst.GetStubLen()
		for tcnt > 0 {
			err = (*PhpConv)(inst.GetCd()).GetConvertOp()((*PhpConv)(inst.GetCd()), &pt, &tcnt, &pd, &ocnt)
			switch err {
			case PHP_CONV_ERR_INVALID_SEQ:
				core.PhpErrorDocref(nil, 1<<1, "stream filter (%s): invalid byte sequence", inst.GetFiltername())
				goto out_failure
			case PHP_CONV_ERR_MORE:
				if ps != nil {
					if icnt > 0 {
						if inst.GetStubLen() >= g.SizeOf("inst -> stub") {
							core.PhpErrorDocref(nil, 1<<1, "stream filter (%s): insufficient buffer", inst.GetFiltername())
							goto out_failure
						}
						inst.GetStub()[g.PostInc(&(inst.GetStubLen()))] = *(g.PostInc(&ps))
						icnt--
						pt = inst.GetStub()
						tcnt = inst.GetStubLen()
					} else {
						tcnt = 0
						break
					}
				}
				break
			case PHP_CONV_ERR_UNEXPECTED_EOS:
				core.PhpErrorDocref(nil, 1<<1, "stream filter (%s): unexpected end of stream", inst.GetFiltername())
				goto out_failure
			case PHP_CONV_ERR_TOO_BIG:
				var new_out_buf *byte
				var new_out_buf_size int
				new_out_buf_size = out_buf_size << 1
				if new_out_buf_size < out_buf_size {

					/* whoa! no bigger buckets are sold anywhere... */

					if nil == g.Assign(&new_bucket, streams.PhpStreamBucketNew(stream, out_buf, out_buf_size-ocnt, 1, persistent)) {
						goto out_failure
					}
					streams.PhpStreamBucketAppend(buckets_out, new_bucket)
					ocnt = initial_out_buf_size
					out_buf_size = ocnt
					if persistent != 0 {
						out_buf = zend.__zendMalloc(out_buf_size)
					} else {
						out_buf = zend._emalloc(out_buf_size)
					}
					pd = out_buf
				} else {
					if persistent != 0 {
						new_out_buf = zend.__zendRealloc(out_buf, new_out_buf_size)
					} else {
						new_out_buf = zend._erealloc(out_buf, new_out_buf_size)
					}
					pd = new_out_buf + (pd - out_buf)
					ocnt += new_out_buf_size - out_buf_size
					out_buf = new_out_buf
					out_buf_size = new_out_buf_size
				}
				break
			case PHP_CONV_ERR_UNKNOWN:
				core.PhpErrorDocref(nil, 1<<1, "stream filter (%s): unknown error", inst.GetFiltername())
				goto out_failure
			default:
				break
			}
		}
		memmove(inst.GetStub(), pt, tcnt)
		inst.SetStubLen(tcnt)
	}
	for icnt > 0 {
		if ps == nil {
			err = (*PhpConv)(inst.GetCd()).GetConvertOp()((*PhpConv)(inst.GetCd()), nil, nil, &pd, &ocnt)
		} else {
			err = (*PhpConv)(inst.GetCd()).GetConvertOp()((*PhpConv)(inst.GetCd()), &ps, &icnt, &pd, &ocnt)
		}
		switch err {
		case PHP_CONV_ERR_INVALID_SEQ:
			core.PhpErrorDocref(nil, 1<<1, "stream filter (%s): invalid byte sequence", inst.GetFiltername())
			goto out_failure
		case PHP_CONV_ERR_MORE:
			if ps != nil {
				if icnt > g.SizeOf("inst -> stub") {
					core.PhpErrorDocref(nil, 1<<1, "stream filter (%s): insufficient buffer", inst.GetFiltername())
					goto out_failure
				}
				memcpy(inst.GetStub(), ps, icnt)
				inst.SetStubLen(icnt)
				ps += icnt
				icnt = 0
			} else {
				core.PhpErrorDocref(nil, 1<<1, "stream filter (%s): unexpected octet values", inst.GetFiltername())
				goto out_failure
			}
			break
		case PHP_CONV_ERR_TOO_BIG:
			var new_out_buf *byte
			var new_out_buf_size int
			new_out_buf_size = out_buf_size << 1
			if new_out_buf_size < out_buf_size {

				/* whoa! no bigger buckets are sold anywhere... */

				if nil == g.Assign(&new_bucket, streams.PhpStreamBucketNew(stream, out_buf, out_buf_size-ocnt, 1, persistent)) {
					goto out_failure
				}
				streams.PhpStreamBucketAppend(buckets_out, new_bucket)
				ocnt = initial_out_buf_size
				out_buf_size = ocnt
				if persistent != 0 {
					out_buf = zend.__zendMalloc(out_buf_size)
				} else {
					out_buf = zend._emalloc(out_buf_size)
				}
				pd = out_buf
			} else {
				if persistent != 0 {
					new_out_buf = zend.__zendRealloc(out_buf, new_out_buf_size)
				} else {
					new_out_buf = zend._erealloc(out_buf, new_out_buf_size)
				}
				pd = new_out_buf + (pd - out_buf)
				ocnt += new_out_buf_size - out_buf_size
				out_buf = new_out_buf
				out_buf_size = new_out_buf_size
			}
			break
		case PHP_CONV_ERR_UNKNOWN:
			core.PhpErrorDocref(nil, 1<<1, "stream filter (%s): unknown error", inst.GetFiltername())
			goto out_failure
		default:
			if ps == nil {
				icnt = 0
			}
			break
		}
	}
	if out_buf_size > ocnt {
		if nil == g.Assign(&new_bucket, streams.PhpStreamBucketNew(stream, out_buf, out_buf_size-ocnt, 1, persistent)) {
			goto out_failure
		}
		streams.PhpStreamBucketAppend(buckets_out, new_bucket)
	} else {
		g.CondF(persistent != 0, func() { return zend.Free(out_buf) }, func() { return zend._efree(out_buf) })
	}
	*consumed += buf_len - icnt
	return zend.SUCCESS
out_failure:
	g.CondF(persistent != 0, func() { return zend.Free(out_buf) }, func() { return zend._efree(out_buf) })
	return zend.FAILURE
}

/* }}} */

func StrfilterConvertFilter(stream *core.PhpStream, thisfilter *core.PhpStreamFilter, buckets_in *streams.PhpStreamBucketBrigade, buckets_out *streams.PhpStreamBucketBrigade, bytes_consumed *int, flags int) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket = nil
	var consumed int = 0
	var inst *PhpConvertFilter = (*PhpConvertFilter)(thisfilter.abstract.value.ptr)
	for buckets_in.head != nil {
		bucket = buckets_in.head
		streams.PhpStreamBucketUnlink(bucket)
		if StrfilterConvertAppendBucket(inst, stream, thisfilter, buckets_out, bucket.buf, bucket.buflen, &consumed, stream.is_persistent) != zend.SUCCESS {
			goto out_failure
		}
		streams.PhpStreamBucketDelref(bucket)
	}
	if flags != 0 {
		if StrfilterConvertAppendBucket(inst, stream, thisfilter, buckets_out, nil, 0, &consumed, stream.is_persistent) != zend.SUCCESS {
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
	assert(thisfilter.abstract.value.ptr != nil)
	PhpConvertFilterDtor((*PhpConvertFilter)(thisfilter.abstract.value.ptr))
	g.CondF((*PhpConvertFilter)(thisfilter.abstract.value.ptr).GetPersistent() != 0, func() { return zend.Free(thisfilter.abstract.value.ptr) }, func() { return zend._efree(thisfilter.abstract.value.ptr) })
}

var StrfilterConvertOps streams.PhpStreamFilterOps = streams.PhpStreamFilterOps{StrfilterConvertFilter, StrfilterConvertDtor, "convert.*"}

func StrfilterConvertCreate(filtername *byte, filterparams *zend.Zval, persistent uint8) *core.PhpStreamFilter {
	var inst *PhpConvertFilter
	var retval *core.PhpStreamFilter = nil
	var dot *byte
	var conv_mode int = 0
	if filterparams != nil && filterparams.u1.v.type_ != 7 {
		core.PhpErrorDocref(nil, 1<<1, "stream filter (%s): invalid filter parameter", filtername)
		return nil
	}
	if g.Assign(&dot, strchr(filtername, '.')) == nil {
		return nil
	}
	dot++
	if persistent != 0 {
		inst = zend.__zendMalloc(g.SizeOf("php_convert_filter"))
	} else {
		inst = zend._emalloc(g.SizeOf("php_convert_filter"))
	}
	if strcasecmp(dot, "base64-encode") == 0 {
		conv_mode = 1
	} else if strcasecmp(dot, "base64-decode") == 0 {
		conv_mode = 2
	} else if strcasecmp(dot, "quoted-printable-encode") == 0 {
		conv_mode = 3
	} else if strcasecmp(dot, "quoted-printable-decode") == 0 {
		conv_mode = 4
	}
	if PhpConvertFilterCtor(inst, conv_mode, g.CondF1(filterparams != nil, func() *zend.ZendArray { return filterparams.value.arr }, nil), filtername, persistent) != zend.SUCCESS {
		goto out
	}
	retval = streams._phpStreamFilterAlloc(&StrfilterConvertOps, inst, persistent)
out:
	if retval == nil {
		g.CondF(persistent != 0, func() { return zend.Free(inst) }, func() { return zend._efree(inst) })
	}
	return retval
}

var StrfilterConvertFactory streams.PhpStreamFilterFactory = streams.PhpStreamFilterFactory{StrfilterConvertCreate}

/* }}} */

// @type PhpConsumedFilterData struct

func ConsumedFilterFilter(stream *core.PhpStream, thisfilter *core.PhpStreamFilter, buckets_in *streams.PhpStreamBucketBrigade, buckets_out *streams.PhpStreamBucketBrigade, bytes_consumed *int, flags int) streams.PhpStreamFilterStatusT {
	var data *PhpConsumedFilterData = (*PhpConsumedFilterData)(thisfilter.abstract.value.ptr)
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	if data.GetOffset() == ^0 {
		data.SetOffset(streams._phpStreamTell(stream))
	}
	for g.Assign(&bucket, buckets_in.head) != nil {
		streams.PhpStreamBucketUnlink(bucket)
		consumed += bucket.buflen
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	if (flags & 2) != 0 {
		streams._phpStreamSeek(stream, data.GetOffset()+data.GetConsumed(), SEEK_SET)
	}
	data.SetConsumed(data.GetConsumed() + consumed)
	return streams.PSFS_PASS_ON
}
func ConsumedFilterDtor(thisfilter *core.PhpStreamFilter) {
	if thisfilter != nil && thisfilter.abstract.value.ptr {
		var data *PhpConsumedFilterData = (*PhpConsumedFilterData)(thisfilter.abstract.value.ptr)
		g.CondF(data.GetPersistent() != 0, func() { return zend.Free(data) }, func() { return zend._efree(data) })
	}
}

var ConsumedFilterOps streams.PhpStreamFilterOps = streams.PhpStreamFilterOps{ConsumedFilterFilter, ConsumedFilterDtor, "consumed"}

func ConsumedFilterCreate(filtername *byte, filterparams *zend.Zval, persistent uint8) *core.PhpStreamFilter {
	var fops *streams.PhpStreamFilterOps = nil
	var data *PhpConsumedFilterData
	if strcasecmp(filtername, "consumed") {
		return nil
	}

	/* Create this filter */

	if persistent != 0 {
		data = zend.__zendCalloc(1, g.SizeOf("php_consumed_filter_data"))
	} else {
		data = zend._ecalloc(1, g.SizeOf("php_consumed_filter_data"))
	}
	data.SetPersistent(persistent)
	data.SetConsumed(0)
	data.SetOffset(^0)
	fops = &ConsumedFilterOps
	return streams._phpStreamFilterAlloc(fops, data, persistent)
}

var ConsumedFilterFactory streams.PhpStreamFilterFactory = streams.PhpStreamFilterFactory{ConsumedFilterCreate}

/* }}} */

type PhpChunkedFilterState = int

const (
	CHUNK_SIZE_START = iota
	CHUNK_SIZE
	CHUNK_SIZE_EXT
	CHUNK_SIZE_CR
	CHUNK_SIZE_LF
	CHUNK_BODY
	CHUNK_BODY_CR
	CHUNK_BODY_LF
	CHUNK_TRAILER
	CHUNK_ERROR
)

// @type PhpChunkedFilterData struct

func PhpDechunk(buf *byte, len_ int, data *PhpChunkedFilterData) int {
	var p *byte = buf
	var end *byte = p + len_
	var out *byte = buf
	var out_len int = 0
	for p < end {
		switch data.GetState() {
		case CHUNK_SIZE_START:
			data.SetChunkSize(0)
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
		case CHUNK_SIZE_EXT:

			/* skip extension */

			for p < end && (*p) != '\r' && (*p) != '\n' {
				p++
			}
			if p == end {
				return out_len
			}
		case CHUNK_SIZE_CR:
			if (*p) == '\r' {
				p++
				if p == end {
					data.SetState(CHUNK_SIZE_LF)
					return out_len
				}
			}
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
		case CHUNK_BODY_CR:
			if (*p) == '\r' {
				p++
				if p == end {
					data.SetState(CHUNK_BODY_LF)
					return out_len
				}
			}
		case CHUNK_BODY_LF:
			if (*p) == '\n' {
				p++
				data.SetState(CHUNK_SIZE_START)
				continue
			} else {
				data.SetState(CHUNK_ERROR)
				continue
			}
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
func PhpChunkedFilter(stream *core.PhpStream, thisfilter *core.PhpStreamFilter, buckets_in *streams.PhpStreamBucketBrigade, buckets_out *streams.PhpStreamBucketBrigade, bytes_consumed *int, flags int) streams.PhpStreamFilterStatusT {
	var bucket *streams.PhpStreamBucket
	var consumed int = 0
	var data *PhpChunkedFilterData = (*PhpChunkedFilterData)(thisfilter.abstract.value.ptr)
	for buckets_in.head != nil {
		bucket = streams.PhpStreamBucketMakeWriteable(buckets_in.head)
		consumed += bucket.buflen
		bucket.buflen = PhpDechunk(bucket.buf, bucket.buflen, data)
		streams.PhpStreamBucketAppend(buckets_out, bucket)
	}
	if bytes_consumed != nil {
		*bytes_consumed = consumed
	}
	return streams.PSFS_PASS_ON
}
func PhpChunkedDtor(thisfilter *core.PhpStreamFilter) {
	if thisfilter != nil && thisfilter.abstract.value.ptr {
		var data *PhpChunkedFilterData = (*PhpChunkedFilterData)(thisfilter.abstract.value.ptr)
		g.CondF(data.GetPersistent() != 0, func() { return zend.Free(data) }, func() { return zend._efree(data) })
	}
}

var ChunkedFilterOps streams.PhpStreamFilterOps = streams.PhpStreamFilterOps{PhpChunkedFilter, PhpChunkedDtor, "dechunk"}

func ChunkedFilterCreate(filtername *byte, filterparams *zend.Zval, persistent uint8) *core.PhpStreamFilter {
	var fops *streams.PhpStreamFilterOps = nil
	var data *PhpChunkedFilterData
	if strcasecmp(filtername, "dechunk") {
		return nil
	}

	/* Create this filter */

	data = (*PhpChunkedFilterData)(g.CondF(persistent != 0, func() any { return zend.__zendCalloc(1, g.SizeOf("php_chunked_filter_data")) }, func() any { return zend._ecalloc(1, g.SizeOf("php_chunked_filter_data")) }))
	data.SetState(CHUNK_SIZE_START)
	data.SetChunkSize(0)
	data.SetPersistent(persistent)
	fops = &ChunkedFilterOps
	return streams._phpStreamFilterAlloc(fops, data, persistent)
}

var ChunkedFilterFactory streams.PhpStreamFilterFactory = streams.PhpStreamFilterFactory{ChunkedFilterCreate}

/* }}} */

var StandardFilters []struct {
	ops     *streams.PhpStreamFilterOps
	factory *streams.PhpStreamFilterFactory
} = []struct {
	ops     *streams.PhpStreamFilterOps
	factory *streams.PhpStreamFilterFactory
}{
	{&StrfilterRot13Ops, &StrfilterRot13Factory},
	{&StrfilterToupperOps, &StrfilterToupperFactory},
	{&StrfilterTolowerOps, &StrfilterTolowerFactory},
	{&StrfilterStripTagsOps, &StrfilterStripTagsFactory},
	{&StrfilterConvertOps, &StrfilterConvertFactory},
	{&ConsumedFilterOps, &ConsumedFilterFactory},
	{&ChunkedFilterOps, &ChunkedFilterFactory},
	{nil, nil},
}

/* {{{ filter MINIT and MSHUTDOWN */

func ZmStartupStandardFilters(type_ int, module_number int) int {
	var i int
	for i = 0; StandardFilters[i].ops != nil; i++ {
		if zend.FAILURE == streams.PhpStreamFilterRegisterFactory(StandardFilters[i].ops.label, StandardFilters[i].factory) {
			return zend.FAILURE
		}
	}
	return zend.SUCCESS
}
func ZmShutdownStandardFilters(type_ int, module_number int) int {
	var i int
	for i = 0; StandardFilters[i].ops != nil; i++ {
		streams.PhpStreamFilterUnregisterFactory(StandardFilters[i].ops.label)
	}
	return zend.SUCCESS
}

/* }}} */
