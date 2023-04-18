package standard

import (
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/php/types"
)

var Rot13From = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var Rot13To = "nopqrstuvwxyzabcdefghijklmNOPQRSTUVWXYZABCDEFGHIJKLM"
var StrfilterRot13Ops = streams.MakePhpStreamFilterOps(StrfilterRot13Filter, nil, "string.rot13")
var StrfilterRot13Factory = streams.MakePhpStreamFilterFactory(StrfilterRot13Create)
var Lowercase = "abcdefghijklmnopqrstuvwxyz"
var Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var StrfilterToupperOps = streams.MakePhpStreamFilterOps(StrfilterToupperFilter, nil, "string.toupper")
var StrfilterTolowerOps = streams.MakePhpStreamFilterOps(StrfilterTolowerFilter, nil, "string.tolower")
var StrfilterToupperFactory = streams.MakePhpStreamFilterFactory(StrfilterToupperCreate)
var StrfilterTolowerFactory = streams.MakePhpStreamFilterFactory(StrfilterTolowerCreate)

var StrfilterStripTagsOps = streams.MakePhpStreamFilterOps(StrfilterStripTagsFilter, StrfilterStripTagsDtor, "string.strip_tags")
var StrfilterStripTagsFactory = streams.MakePhpStreamFilterFactory(StrfilterStripTagsCreate)

type PhpConvErrT = int

const (
	PHP_CONV_ERR_SUCCESS PhpConvErrT = types.SUCCESS
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

var B64TblEnc = []uint8{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/'}

var B64TblDec = []uint{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 62, 64, 64, 64, 63, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 64, 64, 64, 128, 64, 64, 64, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 64, 64, 64, 64, 64, 64, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64}

const PHP_CONV_QPRINT_OPT_BINARY = 0x1
const PHP_CONV_QPRINT_OPT_FORCE_ENCODE_FIRST = 0x2

const PHP_CONV_BASE64_ENCODE = 1
const PHP_CONV_BASE64_DECODE = 2
const PHP_CONV_QPRINT_ENCODE = 3
const PHP_CONV_QPRINT_DECODE = 4

var StrfilterConvertOps = streams.MakePhpStreamFilterOps(StrfilterConvertFilter, StrfilterConvertDtor, "convert.*")
var StrfilterConvertFactory = streams.MakePhpStreamFilterFactory(StrfilterConvertCreate)

var ConsumedFilterOps = streams.MakePhpStreamFilterOps(ConsumedFilterFilter, ConsumedFilterDtor, "consumed")
var ConsumedFilterFactory = streams.MakePhpStreamFilterFactory(ConsumedFilterCreate)

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

var ChunkedFilterOps = streams.MakePhpStreamFilterOps(PhpChunkedFilter, PhpChunkedDtor, "dechunk")
var ChunkedFilterFactory = streams.MakePhpStreamFilterFactory(ChunkedFilterCreate)
var StandardFilters = []struct {
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
