// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/image.c>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include < stdio . h >

// # include < fcntl . h >

// # include "fopen_wrappers.h"

// # include "ext/standard/fsock.h"

// # include < unistd . h >

// # include "php_image.h"

/* file type markers */

var PhpSigGif []byte = []byte{'G', 'I', 'F'}
var PhpSigPsd []byte = []byte{'8', 'B', 'P', 'S'}
var PhpSigBmp []byte = []byte{'B', 'M'}
var PhpSigSwf []byte = []byte{'F', 'W', 'S'}
var PhpSigSwc []byte = []byte{'C', 'W', 'S'}
var PhpSigJpg []byte = []byte{byte(0xff), byte(0xd8), byte(0xff)}
var PhpSigPng []byte = []byte{byte(0x89), byte(0x50), byte(0x4e), byte(0x47), byte(0xd), byte(0xa), byte(0x1a), byte(0xa)}
var PhpSigTifIi []byte = []byte{'I', 'I', byte(0x2a), byte(0x0)}
var PhpSigTifMm []byte = []byte{'M', 'M', byte(0x0), byte(0x2a)}
var PhpSigJpc []byte = []byte{byte(0xff), byte(0x4f), byte(0xff)}
var PhpSigJp2 []byte = []byte{byte(0x0), byte(0x0), byte(0x0), byte(0xc), byte(0x6a), byte(0x50), byte(0x20), byte(0x20), byte(0xd), byte(0xa), byte(0x87), byte(0xa)}
var PhpSigIff []byte = []byte{'F', 'O', 'R', 'M'}
var PhpSigIco []byte = []byte{byte(0x0), byte(0x0), byte(0x1), byte(0x0)}
var PhpSigRiff []byte = []byte{'R', 'I', 'F', 'F'}
var PhpSigWebp []byte = []byte{'W', 'E', 'B', 'P'}

/* REMEMBER TO ADD MIME-TYPE TO FUNCTION php_image_type_to_mime_type */

// @type Gfxinfo struct

/* {{{ PHP_MINIT_FUNCTION(imagetypes)
 * Register IMAGETYPE_<xxx> constants used by GetImageSize(), image_type_to_mime_type, ext/exif */

func ZmStartupImagetypes(type_ int, module_number int) int {
	zend.ZendRegisterLongConstant("IMAGETYPE_GIF", g.SizeOf("\"IMAGETYPE_GIF\"")-1, IMAGE_FILETYPE_GIF, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_JPEG", g.SizeOf("\"IMAGETYPE_JPEG\"")-1, IMAGE_FILETYPE_JPEG, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_PNG", g.SizeOf("\"IMAGETYPE_PNG\"")-1, IMAGE_FILETYPE_PNG, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_SWF", g.SizeOf("\"IMAGETYPE_SWF\"")-1, IMAGE_FILETYPE_SWF, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_PSD", g.SizeOf("\"IMAGETYPE_PSD\"")-1, IMAGE_FILETYPE_PSD, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_BMP", g.SizeOf("\"IMAGETYPE_BMP\"")-1, IMAGE_FILETYPE_BMP, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_TIFF_II", g.SizeOf("\"IMAGETYPE_TIFF_II\"")-1, IMAGE_FILETYPE_TIFF_II, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_TIFF_MM", g.SizeOf("\"IMAGETYPE_TIFF_MM\"")-1, IMAGE_FILETYPE_TIFF_MM, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_JPC", g.SizeOf("\"IMAGETYPE_JPC\"")-1, IMAGE_FILETYPE_JPC, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_JP2", g.SizeOf("\"IMAGETYPE_JP2\"")-1, IMAGE_FILETYPE_JP2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_JPX", g.SizeOf("\"IMAGETYPE_JPX\"")-1, IMAGE_FILETYPE_JPX, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_JB2", g.SizeOf("\"IMAGETYPE_JB2\"")-1, IMAGE_FILETYPE_JB2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_IFF", g.SizeOf("\"IMAGETYPE_IFF\"")-1, IMAGE_FILETYPE_IFF, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_WBMP", g.SizeOf("\"IMAGETYPE_WBMP\"")-1, IMAGE_FILETYPE_WBMP, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_JPEG2000", g.SizeOf("\"IMAGETYPE_JPEG2000\"")-1, IMAGE_FILETYPE_JPC, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_XBM", g.SizeOf("\"IMAGETYPE_XBM\"")-1, IMAGE_FILETYPE_XBM, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_ICO", g.SizeOf("\"IMAGETYPE_ICO\"")-1, IMAGE_FILETYPE_ICO, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_WEBP", g.SizeOf("\"IMAGETYPE_WEBP\"")-1, IMAGE_FILETYPE_WEBP, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_UNKNOWN", g.SizeOf("\"IMAGETYPE_UNKNOWN\"")-1, IMAGE_FILETYPE_UNKNOWN, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("IMAGETYPE_COUNT", g.SizeOf("\"IMAGETYPE_COUNT\"")-1, IMAGE_FILETYPE_COUNT, 1<<0|1<<1, module_number)
	return zend.SUCCESS
}

/* }}} */

func PhpHandleGif(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var dim []uint8
	if streams._phpStreamSeek(stream, 3, SEEK_CUR) != 0 {
		return nil
	}
	if streams._phpStreamRead(stream, (*byte)(dim), g.SizeOf("dim")) != g.SizeOf("dim") {
		return nil
	}
	result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
	result.SetWidth(uint(dim[0] | uint(dim[1])<<8))
	result.SetHeight(uint(dim[2] | uint(dim[3])<<8))
	if (dim[4] & 0x80) != 0 {
		result.SetBits((uint(dim[4]) & 0x7) + 1)
	} else {
		result.SetBits(0)
	}
	result.SetChannels(3)
	return result
}

/* }}} */

func PhpHandlePsd(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var dim []uint8
	if streams._phpStreamSeek(stream, 11, SEEK_CUR) != 0 {
		return nil
	}
	if streams._phpStreamRead(stream, (*byte)(dim), g.SizeOf("dim")) != g.SizeOf("dim") {
		return nil
	}
	result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
	result.SetHeight((uint(dim[0]) << 24) + (uint(dim[1]) << 16) + (uint(dim[2]) << 8) + uint(dim[3]))
	result.SetWidth((uint(dim[4]) << 24) + (uint(dim[5]) << 16) + (uint(dim[6]) << 8) + uint(dim[7]))
	return result
}

/* }}} */

func PhpHandleBmp(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var dim []uint8
	var size int
	if streams._phpStreamSeek(stream, 11, SEEK_CUR) != 0 {
		return nil
	}
	if streams._phpStreamRead(stream, (*byte)(dim), g.SizeOf("dim")) != g.SizeOf("dim") {
		return nil
	}
	size = (uint(dim[3]) << 24) + (uint(dim[2]) << 16) + (uint(dim[1]) << 8) + uint(dim[0])
	if size == 12 {
		result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
		result.SetWidth((uint(dim[5]) << 8) + uint(dim[4]))
		result.SetHeight((uint(dim[7]) << 8) + uint(dim[6]))
		result.SetBits(uint(dim[11]))
	} else if size > 12 && (size <= 64 || size == 108 || size == 124) {
		result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
		result.SetWidth((uint(dim[7]) << 24) + (uint(dim[6]) << 16) + (uint(dim[5]) << 8) + uint(dim[4]))
		result.SetHeight((uint(dim[11]) << 24) + (uint(dim[10]) << 16) + (uint(dim[9]) << 8) + uint(dim[8]))
		result.SetHeight(abs(int32(result.GetHeight())))
		result.SetBits((uint(dim[15]) << 8) + uint(dim[14]))
	} else {
		return nil
	}
	return result
}

/* }}} */

func PhpSwfGetBits(buffer *uint8, pos uint, count uint) unsigned__long__int {
	var loop uint
	var result unsigned__long__int = 0
	for loop = pos; loop < pos+count; loop++ {
		result = result + ((buffer[loop/8]>>7-loop%8&0x1)<<count - (loop - pos) - 1)
	}
	return result
}

/* }}} */

/* {{{ php_handle_swf
 */

func PhpHandleSwf(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var bits long
	var a []uint8
	if streams._phpStreamSeek(stream, 5, SEEK_CUR) != 0 {
		return nil
	}
	if streams._phpStreamRead(stream, (*byte)(a), g.SizeOf("a")) != g.SizeOf("a") {
		return nil
	}
	result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
	bits = PhpSwfGetBits(a, 0, 5)
	result.SetWidth((PhpSwfGetBits(a, 5+bits, bits) - PhpSwfGetBits(a, 5, bits)) / 20)
	result.SetHeight((PhpSwfGetBits(a, 5+3*bits, bits) - PhpSwfGetBits(a, 5+2*bits, bits)) / 20)
	result.SetBits(0)
	result.SetChannels(0)
	return result
}

/* }}} */

func PhpHandlePng(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var dim []uint8

	/* Width:              4 bytes
	 * Height:             4 bytes
	 * Bit depth:          1 byte
	 * Color type:         1 byte
	 * Compression method: 1 byte
	 * Filter method:      1 byte
	 * Interlace method:   1 byte
	 */

	if streams._phpStreamSeek(stream, 8, SEEK_CUR) != 0 {
		return nil
	}
	if streams._phpStreamRead(stream, (*byte)(dim), g.SizeOf("dim")) < g.SizeOf("dim") {
		return nil
	}
	result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
	result.SetWidth((uint(dim[0]) << 24) + (uint(dim[1]) << 16) + (uint(dim[2]) << 8) + uint(dim[3]))
	result.SetHeight((uint(dim[4]) << 24) + (uint(dim[5]) << 16) + (uint(dim[6]) << 8) + uint(dim[7]))
	result.SetBits(uint(dim[8]))
	return result
}

/* }}} */

// #define M_SOF0       0xC0

// #define M_SOF1       0xC1

// #define M_SOF2       0xC2

// #define M_SOF3       0xC3

// #define M_SOF5       0xC5

// #define M_SOF6       0xC6

// #define M_SOF7       0xC7

// #define M_SOF9       0xC9

// #define M_SOF10       0xCA

// #define M_SOF11       0xCB

// #define M_SOF13       0xCD

// #define M_SOF14       0xCE

// #define M_SOF15       0xCF

// #define M_SOI       0xD8

// #define M_EOI       0xD9

// #define M_SOS       0xDA

// #define M_APP0       0xe0

// #define M_APP1       0xe1

// #define M_APP2       0xe2

// #define M_APP3       0xe3

// #define M_APP4       0xe4

// #define M_APP5       0xe5

// #define M_APP6       0xe6

// #define M_APP7       0xe7

// #define M_APP8       0xe8

// #define M_APP9       0xe9

// #define M_APP10       0xea

// #define M_APP11       0xeb

// #define M_APP12       0xec

// #define M_APP13       0xed

// #define M_APP14       0xee

// #define M_APP15       0xef

// #define M_COM       0xFE

// #define M_PSEUDO       0xFFD8

/* {{{ php_read2
 */

func PhpRead2(stream *core.PhpStream) uint16 {
	var a []uint8

	/* return 0 if we couldn't read enough data */

	if streams._phpStreamRead(stream, (*byte)(a), g.SizeOf("a")) < g.SizeOf("a") {
		return 0
	}
	return (uint16(a[0]) << 8) + uint16(a[1])
}

/* }}} */

func PhpNextMarker(stream *core.PhpStream, last_marker int, ff_read int) uint {
	var a int = 0
	var marker int

	/* get marker byte, swallowing possible padding                           */

	if ff_read == 0 {
		var extraneous int = 0
		for g.Assign(&marker, streams._phpStreamGetc(stream)) != 0xff {
			if marker == EOF {
				return 0xd9
			}
			extraneous++
		}
		if extraneous != 0 {
			core.PhpErrorDocref(nil, 1<<1, "corrupt JPEG data: %zu extraneous bytes before marker", extraneous)
		}
	}
	a = 1
	for {
		if g.Assign(&marker, streams._phpStreamGetc(stream)) == EOF {
			return 0xd9
		}
		a++
		if marker != 0xff {
			break
		}
	}
	if a < 2 {
		return 0xd9
	}
	return uint(marker)
}

/* }}} */

func PhpSkipVariable(stream *core.PhpStream) int {
	var length zend.ZendOffT = uint(PhpRead2(stream))
	if length < 2 {
		return 0
	}
	length = length - 2
	streams._phpStreamSeek(stream, zend.ZendLong(length), SEEK_CUR)
	return 1
}

/* }}} */

func php_read_APP(stream *core.PhpStream, marker uint, info *zend.Zval) int {
	var length uint16
	var buffer *byte
	var markername []byte
	var tmp *zend.Zval
	length = PhpRead2(stream)
	if length < 2 {
		return 0
	}
	length -= 2
	buffer = zend._emalloc(int(length))
	if streams._phpStreamRead(stream, buffer, int(length)) != length {
		zend._efree(buffer)
		return 0
	}
	core.ApPhpSnprintf(markername, g.SizeOf("markername"), "APP%d", marker-0xe0)
	if g.Assign(&tmp, zend.ZendHashStrFind(info.value.arr, markername, strlen(markername))) == nil {

		/* XXX we only catch the 1st tag of it's kind! */

		zend.AddAssocStringlEx(info, markername, strlen(markername), buffer, length)

		/* XXX we only catch the 1st tag of it's kind! */

	}
	zend._efree(buffer)
	return 1
}

/* }}} */

func PhpHandleJpeg(stream *core.PhpStream, info *zend.Zval) *Gfxinfo {
	var result *Gfxinfo = nil
	var marker uint = 0xffd8
	var length uint16
	var ff_read uint16 = 1
	for {
		marker = PhpNextMarker(stream, marker, ff_read)
		ff_read = 0
		switch marker {
		case 0xc0:

		case 0xc1:

		case 0xc2:

		case 0xc3:

		case 0xc5:

		case 0xc6:

		case 0xc7:

		case 0xc9:

		case 0xca:

		case 0xcb:

		case 0xcd:

		case 0xce:

		case 0xcf:
			if result == nil {

				/* handle SOFn block */

				result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
				length = PhpRead2(stream)
				result.SetBits(streams._phpStreamGetc(stream))
				result.SetHeight(PhpRead2(stream))
				result.SetWidth(PhpRead2(stream))
				result.SetChannels(streams._phpStreamGetc(stream))
				if info == nil || length < 8 {
					return result
				}
				if streams._phpStreamSeek(stream, length-8, SEEK_CUR) != 0 {
					return result
				}
			} else {
				if PhpSkipVariable(stream) == 0 {
					return result
				}
			}
			break
		case 0xe0:

		case 0xe1:

		case 0xe2:

		case 0xe3:

		case 0xe4:

		case 0xe5:

		case 0xe6:

		case 0xe7:

		case 0xe8:

		case 0xe9:

		case 0xea:

		case 0xeb:

		case 0xec:

		case 0xed:

		case 0xee:

		case 0xef:
			if info != nil {
				if php_read_APP(stream, marker, info) == 0 {
					return result
				}
			} else {
				if PhpSkipVariable(stream) == 0 {
					return result
				}
			}
			break
		case 0xda:

		case 0xd9:
			return result
		default:
			if PhpSkipVariable(stream) == 0 {
				return result
			}
			break
		}
	}
	return result
}

/* }}} */

func PhpRead4(stream *core.PhpStream) uint {
	var a []uint8

	/* just return 0 if we hit the end-of-file */

	if streams._phpStreamRead(stream, (*byte)(a), g.SizeOf("a")) != g.SizeOf("a") {
		return 0
	}
	return (uint(a[0]) << 24) + (uint(a[1]) << 16) + (uint(a[2]) << 8) + uint(a[3])
}

/* }}} */

// #define JPEG2000_MARKER_PREFIX       0xFF

// #define JPEG2000_MARKER_SOC       0x4F

// #define JPEG2000_MARKER_SOT       0x90

// #define JPEG2000_MARKER_SOD       0x93

// #define JPEG2000_MARKER_EOC       0xD9

// #define JPEG2000_MARKER_SIZ       0x51

// #define JPEG2000_MARKER_COD       0x52

// #define JPEG2000_MARKER_COC       0x53

// #define JPEG2000_MARKER_RGN       0x5E

// #define JPEG2000_MARKER_QCD       0x5C

// #define JPEG2000_MARKER_QCC       0x5D

// #define JPEG2000_MARKER_POC       0x5F

// #define JPEG2000_MARKER_TLM       0x55

// #define JPEG2000_MARKER_PLM       0x57

// #define JPEG2000_MARKER_PLT       0x58

// #define JPEG2000_MARKER_PPM       0x60

// #define JPEG2000_MARKER_PPT       0x61

// #define JPEG2000_MARKER_SOP       0x91

// #define JPEG2000_MARKER_EPH       0x92

// #define JPEG2000_MARKER_CRG       0x63

// #define JPEG2000_MARKER_COM       0x64

/* }}} */

func PhpHandleJpc(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var highest_bit_depth int
	var bit_depth int
	var first_marker_id uint8
	var i uint

	/* JPEG 2000 components can be vastly different from one another.
	   Each component can be sampled at a different resolution, use
	   a different colour space, have a separate colour depth, and
	   be compressed totally differently! This makes giving a single
	   "bit depth" answer somewhat problematic. For this implementation
	   we'll use the highest depth encountered. */

	first_marker_id = streams._phpStreamGetc(stream)

	/* Ensure that this marker is SIZ (as is mandated by the standard) */

	if first_marker_id != 0x51 {
		core.PhpErrorDocref(nil, 1<<1, "JPEG2000 codestream corrupt(Expected SIZ marker not found after SOC)")
		return nil
	}
	result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
	PhpRead2(stream)
	PhpRead2(stream)
	result.SetWidth(PhpRead4(stream))
	result.SetHeight(PhpRead4(stream))
	if streams._phpStreamSeek(stream, 24, SEEK_CUR) != 0 {
		zend._efree(result)
		return nil
	}
	result.SetChannels(PhpRead2(stream))
	if result.GetChannels() == 0 && streams._phpStreamEof(stream) != 0 || result.GetChannels() > 256 {
		zend._efree(result)
		return nil
	}

	/* Collect bit depth info */

	highest_bit_depth = 0
	for i = 0; i < result.GetChannels(); i++ {
		bit_depth = streams._phpStreamGetc(stream)
		bit_depth++
		if bit_depth > highest_bit_depth {
			highest_bit_depth = bit_depth
		}
		streams._phpStreamGetc(stream)
		streams._phpStreamGetc(stream)
	}
	result.SetBits(highest_bit_depth)
	return result
}

/* }}} */

func PhpHandleJp2(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var box_length uint
	var box_type uint
	var jp2c_box_id []byte = []byte{byte(0x6a), byte(0x70), byte(0x32), byte(0x63)}

	/* JP2 is a wrapper format for JPEG 2000. Data is contained within "boxes".
	   Boxes themselves can be contained within "super-boxes". Super-Boxes can
	   contain super-boxes which provides us with a hierarchical storage system.

	   It is valid for a JP2 file to contain multiple individual codestreams.
	   We'll just look for the first codestream at the root of the box structure
	   and handle that.
	*/

	for {
		box_length = PhpRead4(stream)

		/* TBox */

		if streams._phpStreamRead(stream, any(&box_type), g.SizeOf("box_type")) != g.SizeOf("box_type") {

			/* Use this as a general "out of stream" error */

			break

			/* Use this as a general "out of stream" error */

		}
		if box_length == 1 {

			/* We won't handle XLBoxes */

			return nil

			/* We won't handle XLBoxes */

		}
		if !(memcmp(&box_type, jp2c_box_id, 4)) {

			/* Skip the first 3 bytes to emulate the file type examination */

			streams._phpStreamSeek(stream, 3, SEEK_CUR)
			result = PhpHandleJpc(stream)
			break
		}

		/* Stop if this was the last box */

		if int(box_length <= 0) != 0 {
			break
		}

		/* Skip over LBox (Which includes both TBox and LBox itself */

		if streams._phpStreamSeek(stream, box_length-8, SEEK_CUR) != 0 {
			break
		}

		/* Skip over LBox (Which includes both TBox and LBox itself */

	}
	if result == nil {
		core.PhpErrorDocref(nil, 1<<1, "JP2 file has no codestreams at root level")
	}
	return result
}

/* }}} */

var PhpTiffBytesPerFormat []int = []int{0, 1, 1, 2, 4, 8, 1, 1, 2, 4, 8, 4, 8}

/* uncompressed only */

// #define TAG_IMAGEWIDTH       0x0100

// #define TAG_IMAGEHEIGHT       0x0101

/* compressed images only */

// #define TAG_COMP_IMAGEWIDTH       0xA002

// #define TAG_COMP_IMAGEHEIGHT       0xA003

// #define TAG_FMT_BYTE       1

// #define TAG_FMT_STRING       2

// #define TAG_FMT_USHORT       3

// #define TAG_FMT_ULONG       4

// #define TAG_FMT_URATIONAL       5

// #define TAG_FMT_SBYTE       6

// #define TAG_FMT_UNDEFINED       7

// #define TAG_FMT_SSHORT       8

// #define TAG_FMT_SLONG       9

// #define TAG_FMT_SRATIONAL       10

// #define TAG_FMT_SINGLE       11

// #define TAG_FMT_DOUBLE       12

/* }}} */

func PhpIfdGet16u(Short any, motorola_intel int) int {
	if motorola_intel != 0 {
		return (*uint8)(Short)[0]<<8 | (*uint8)(Short)[1]
	} else {
		return (*uint8)(Short)[1]<<8 | (*uint8)(Short)[0]
	}
}

/* }}} */

func PhpIfdGet16s(Short any, motorola_intel int) signed__short {
	return signed__short(PhpIfdGet16u(Short, motorola_intel))
}

/* }}} */

func PhpIfdGet32s(Long any, motorola_intel int) int {
	if motorola_intel != 0 {
		return (*byte)(Long)[0]<<24 | (*uint8)(Long)[1]<<16 | (*uint8)(Long)[2]<<8 | (*uint8)(Long)[3]<<0
	} else {
		return (*byte)(Long)[3]<<24 | (*uint8)(Long)[2]<<16 | (*uint8)(Long)[1]<<8 | (*uint8)(Long)[0]<<0
	}
}

/* }}} */

func PhpIfdGet32u(Long any, motorola_intel int) unsigned {
	return unsigned(PhpIfdGet32s(Long, motorola_intel) & 0xffffffff)
}

/* }}} */

func PhpHandleTiff(stream *core.PhpStream, info *zend.Zval, motorola_intel int) *Gfxinfo {
	var result *Gfxinfo = nil
	var i int
	var num_entries int
	var dir_entry *uint8
	var ifd_size int
	var dir_size int
	var entry_value int
	var width int = 0
	var height int = 0
	var ifd_addr int
	var entry_tag int
	var entry_type int
	var ifd_data *byte
	var ifd_ptr []*byte
	if streams._phpStreamRead(stream, ifd_ptr, 4) != 4 {
		return nil
	}
	ifd_addr = PhpIfdGet32u(ifd_ptr, motorola_intel)
	if streams._phpStreamSeek(stream, ifd_addr-8, SEEK_CUR) != 0 {
		return nil
	}
	ifd_size = 2
	ifd_data = zend._emalloc(ifd_size)
	if streams._phpStreamRead(stream, ifd_data, 2) != 2 {
		zend._efree(ifd_data)
		return nil
	}
	num_entries = PhpIfdGet16u(ifd_data, motorola_intel)
	dir_size = 2 + 12*num_entries + 4
	ifd_size = dir_size
	ifd_data = zend._erealloc(ifd_data, ifd_size)
	if streams._phpStreamRead(stream, ifd_data+2, dir_size-2) != dir_size-2 {
		zend._efree(ifd_data)
		return nil
	}

	/* now we have the directory we can look how long it should be */

	ifd_size = dir_size
	for i = 0; i < num_entries; i++ {
		dir_entry = (*uint8)(ifd_data + 2 + i*12)
		entry_tag = PhpIfdGet16u(dir_entry+0, motorola_intel)
		entry_type = PhpIfdGet16u(dir_entry+2, motorola_intel)
		switch entry_type {
		case 1:

		case 6:
			entry_value = size_t(dir_entry[8])
			break
		case 3:
			entry_value = PhpIfdGet16u(dir_entry+8, motorola_intel)
			break
		case 8:
			entry_value = PhpIfdGet16s(dir_entry+8, motorola_intel)
			break
		case 4:
			entry_value = PhpIfdGet32u(dir_entry+8, motorola_intel)
			break
		case 9:
			entry_value = PhpIfdGet32s(dir_entry+8, motorola_intel)
			break
		default:
			continue
		}
		switch entry_tag {
		case 0x100:

		case 0xa002:
			width = entry_value
			break
		case 0x101:

		case 0xa003:
			height = entry_value
			break
		}
	}
	zend._efree(ifd_data)
	if width != 0 && height != 0 {

		/* not the same when in for-loop */

		result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
		result.SetHeight(height)
		result.SetWidth(width)
		result.SetBits(0)
		result.SetChannels(0)
		return result
	}
	return nil
}

/* }}} */

func PhpHandleIff(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo
	var a []uint8
	var chunkId int
	var size int
	var width short
	var height short
	var bits short
	if streams._phpStreamRead(stream, (*byte)(a), 8) != 8 {
		return nil
	}
	if strncmp((*byte)(a+4), "ILBM", 4) && strncmp((*byte)(a+4), "PBM ", 4) {
		return nil
	}

	/* loop chunks to find BMHD chunk */

	for {
		if streams._phpStreamRead(stream, (*byte)(a), 8) != 8 {
			return nil
		}
		chunkId = PhpIfdGet32s(a+0, 1)
		size = PhpIfdGet32s(a+4, 1)
		if size < 0 {
			return nil
		}
		if (size & 1) == 1 {
			size++
		}
		if chunkId == 0x424d4844 {
			if size < 9 || streams._phpStreamRead(stream, (*byte)(a), 9) != 9 {
				return nil
			}
			width = PhpIfdGet16s(a+0, 1)
			height = PhpIfdGet16s(a+2, 1)
			bits = a[8] & 0xff
			if width > 0 && height > 0 && bits > 0 && bits < 33 {
				result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
				result.SetWidth(width)
				result.SetHeight(height)
				result.SetBits(bits)
				result.SetChannels(0)
				return result
			}
		} else {
			if streams._phpStreamSeek(stream, size, SEEK_CUR) != 0 {
				return nil
			}
		}

	}

	/* loop chunks to find BMHD chunk */
}

/* }}} */

func PhpGetWbmp(stream *core.PhpStream, result **Gfxinfo, check int) int {
	var i int
	var width int = 0
	var height int = 0
	if streams._phpStreamSeek(stream, 0, SEEK_SET) != 0 {
		return 0
	}

	/* get type */

	if streams._phpStreamGetc(stream) != 0 {
		return 0
	}

	/* skip header */

	for {
		i = streams._phpStreamGetc(stream)
		if i < 0 {
			return 0
		}
		if (i & 0x80) == 0 {
			break
		}
	}

	/* get width */

	for {
		i = streams._phpStreamGetc(stream)
		if i < 0 {
			return 0
		}
		width = width<<7 | i&0x7f

		/* maximum valid width for wbmp (although 127 may be a more accurate one) */

		if width > 2048 {
			return 0
		}

		/* maximum valid width for wbmp (although 127 may be a more accurate one) */

		if (i & 0x80) == 0 {
			break
		}
	}

	/* get height */

	for {
		i = streams._phpStreamGetc(stream)
		if i < 0 {
			return 0
		}
		height = height<<7 | i&0x7f

		/* maximum valid height for wbmp (although 127 may be a more accurate one) */

		if height > 2048 {
			return 0
		}

		/* maximum valid height for wbmp (although 127 may be a more accurate one) */

		if (i & 0x80) == 0 {
			break
		}
	}
	if height == 0 || width == 0 {
		return 0
	}
	if check == 0 {
		(*result).SetWidth(width)
		(*result).SetHeight(height)
	}
	return IMAGE_FILETYPE_WBMP
}

/* }}} */

func PhpHandleWbmp(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
	if PhpGetWbmp(stream, &result, 0) == 0 {
		zend._efree(result)
		return nil
	}
	return result
}

/* }}} */

func PhpGetXbm(stream *core.PhpStream, result **Gfxinfo) int {
	var fline *byte
	var iname *byte
	var type_ *byte
	var value int
	var width uint = 0
	var height uint = 0
	if result != nil {
		*result = nil
	}
	if streams._phpStreamSeek(stream, 0, SEEK_SET) != 0 {
		return 0
	}
	for g.Assign(&fline, streams._phpStreamGetLine(stream, nil, 0, nil)) != nil {
		iname = zend._estrdup(fline)
		if sscanf(fline, "#define %s %d", iname, &value) == 2 {
			if !(g.Assign(&type_, strrchr(iname, '_'))) {
				type_ = iname
			} else {
				type_++
			}
			if !(strcmp("width", type_)) {
				width = uint(value)
				if height != 0 {
					zend._efree(iname)
					break
				}
			}
			if !(strcmp("height", type_)) {
				height = uint(value)
				if width != 0 {
					zend._efree(iname)
					break
				}
			}
		}
		zend._efree(fline)
		zend._efree(iname)
	}
	if fline != nil {
		zend._efree(fline)
	}
	if width != 0 && height != 0 {
		if result != nil {
			*result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
			(*result).SetWidth(width)
			(*result).SetHeight(height)
		}
		return IMAGE_FILETYPE_XBM
	}
	return 0
}

/* }}} */

func PhpHandleXbm(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo
	PhpGetXbm(stream, &result)
	return result
}

/* }}} */

func PhpHandleIco(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var dim []uint8
	var num_icons int = 0
	if streams._phpStreamRead(stream, (*byte)(dim), 2) != 2 {
		return nil
	}
	num_icons = (uint(dim[1]) << 8) + uint(dim[0])
	if num_icons < 1 || num_icons > 255 {
		return nil
	}
	result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
	for num_icons > 0 {
		if streams._phpStreamRead(stream, (*byte)(dim), g.SizeOf("dim")) != g.SizeOf("dim") {
			break
		}
		if (uint(dim[7])<<8)+uint(dim[6]) >= result.GetBits() {
			result.SetWidth(uint(dim[0]))
			result.SetHeight(uint(dim[1]))
			result.SetBits((uint(dim[7]) << 8) + uint(dim[6]))
		}
		num_icons--
	}
	if 0 == result.GetWidth() {
		result.SetWidth(256)
	}
	if 0 == result.GetHeight() {
		result.SetHeight(256)
	}
	return result
}

/* }}} */

func PhpHandleWebp(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var sig []byte = []byte{'V', 'P', '8'}
	var buf []uint8
	var format byte
	if streams._phpStreamRead(stream, (*byte)(buf), 18) != 18 {
		return nil
	}
	if memcmp(buf, sig, 3) {
		return nil
	}
	switch buf[3] {
	case ' ':

	case 'L':

	case 'X':
		format = buf[3]
		break
	default:
		return nil
	}
	result = (*Gfxinfo)(zend._ecalloc(1, g.SizeOf("struct gfxinfo")))
	switch format {
	case ' ':
		result.SetWidth(buf[14] + ((buf[15] & 0x3f) << 8))
		result.SetHeight(buf[16] + ((buf[17] & 0x3f) << 8))
		break
	case 'L':
		result.SetWidth(buf[9] + ((buf[10] & 0x3f) << 8) + 1)
		result.SetHeight((buf[10] >> 6) + (buf[11] << 2) + ((buf[12] & 0xf) << 10) + 1)
		break
	case 'X':
		result.SetWidth(buf[12] + (buf[13] << 8) + (buf[14] << 16) + 1)
		result.SetHeight(buf[15] + (buf[16] << 8) + (buf[17] << 16) + 1)
		break
	}
	result.SetBits(8)
	return result
}

/* }}} */

func PhpImageTypeToMimeType(image_type int) *byte {
	switch image_type {
	case IMAGE_FILETYPE_GIF:
		return "image/gif"
	case IMAGE_FILETYPE_JPEG:
		return "image/jpeg"
	case IMAGE_FILETYPE_PNG:
		return "image/png"
	case IMAGE_FILETYPE_SWF:

	case IMAGE_FILETYPE_SWC:
		return "application/x-shockwave-flash"
	case IMAGE_FILETYPE_PSD:
		return "image/psd"
	case IMAGE_FILETYPE_BMP:
		return "image/bmp"
	case IMAGE_FILETYPE_TIFF_II:

	case IMAGE_FILETYPE_TIFF_MM:
		return "image/tiff"
	case IMAGE_FILETYPE_IFF:
		return "image/iff"
	case IMAGE_FILETYPE_WBMP:
		return "image/vnd.wap.wbmp"
	case IMAGE_FILETYPE_JPC:
		return "application/octet-stream"
	case IMAGE_FILETYPE_JP2:
		return "image/jp2"
	case IMAGE_FILETYPE_XBM:
		return "image/xbm"
	case IMAGE_FILETYPE_ICO:
		return "image/vnd.microsoft.icon"
	case IMAGE_FILETYPE_WEBP:
		return "image/webp"
	default:

	case IMAGE_FILETYPE_UNKNOWN:
		return "application/octet-stream"
	}
}

/* }}} */

func ZifImageTypeToMimeType(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var p_image_type zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &p_image_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	var _s *byte = (*byte)(PhpImageTypeToMimeType(p_image_type))
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
}

/* }}} */

func ZifImageTypeToExtension(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var image_type zend.ZendLong
	var inc_dot zend.ZendBool = 1
	var imgext *byte = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &image_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &inc_dot, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	switch image_type {
	case IMAGE_FILETYPE_GIF:
		imgext = ".gif"
		break
	case IMAGE_FILETYPE_JPEG:
		imgext = ".jpeg"
		break
	case IMAGE_FILETYPE_PNG:
		imgext = ".png"
		break
	case IMAGE_FILETYPE_SWF:

	case IMAGE_FILETYPE_SWC:
		imgext = ".swf"
		break
	case IMAGE_FILETYPE_PSD:
		imgext = ".psd"
		break
	case IMAGE_FILETYPE_BMP:

	case IMAGE_FILETYPE_WBMP:
		imgext = ".bmp"
		break
	case IMAGE_FILETYPE_TIFF_II:

	case IMAGE_FILETYPE_TIFF_MM:
		imgext = ".tiff"
		break
	case IMAGE_FILETYPE_IFF:
		imgext = ".iff"
		break
	case IMAGE_FILETYPE_JPC:
		imgext = ".jpc"
		break
	case IMAGE_FILETYPE_JP2:
		imgext = ".jp2"
		break
	case IMAGE_FILETYPE_JPX:
		imgext = ".jpx"
		break
	case IMAGE_FILETYPE_JB2:
		imgext = ".jb2"
		break
	case IMAGE_FILETYPE_XBM:
		imgext = ".xbm"
		break
	case IMAGE_FILETYPE_ICO:
		imgext = ".ico"
		break
	case IMAGE_FILETYPE_WEBP:
		imgext = ".webp"
		break
	}
	if imgext != nil {
		var _s *byte = &imgext[!inc_dot]
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func PhpGetimagetype(stream *core.PhpStream, filetype *byte) int {
	var tmp []byte
	var twelve_bytes_read int
	if filetype == nil {
		filetype = tmp
	}
	if streams._phpStreamRead(stream, filetype, 3) != 3 {
		core.PhpErrorDocref(nil, 1<<3, "Read error!")
		return IMAGE_FILETYPE_UNKNOWN
	}

	/* BYTES READ: 3 */

	if !(memcmp(filetype, PhpSigGif, 3)) {
		return IMAGE_FILETYPE_GIF
	} else if !(memcmp(filetype, PhpSigJpg, 3)) {
		return IMAGE_FILETYPE_JPEG
	} else if !(memcmp(filetype, PhpSigPng, 3)) {
		if streams._phpStreamRead(stream, filetype+3, 5) != 5 {
			core.PhpErrorDocref(nil, 1<<3, "Read error!")
			return IMAGE_FILETYPE_UNKNOWN
		}
		if !(memcmp(filetype, PhpSigPng, 8)) {
			return IMAGE_FILETYPE_PNG
		} else {
			core.PhpErrorDocref(nil, 1<<1, "PNG file corrupted by ASCII conversion")
			return IMAGE_FILETYPE_UNKNOWN
		}
	} else if !(memcmp(filetype, PhpSigSwf, 3)) {
		return IMAGE_FILETYPE_SWF
	} else if !(memcmp(filetype, PhpSigSwc, 3)) {
		return IMAGE_FILETYPE_SWC
	} else if !(memcmp(filetype, PhpSigPsd, 3)) {
		return IMAGE_FILETYPE_PSD
	} else if !(memcmp(filetype, PhpSigBmp, 2)) {
		return IMAGE_FILETYPE_BMP
	} else if !(memcmp(filetype, PhpSigJpc, 3)) {
		return IMAGE_FILETYPE_JPC
	} else if !(memcmp(filetype, PhpSigRiff, 3)) {
		if streams._phpStreamRead(stream, filetype+3, 9) != 9 {
			core.PhpErrorDocref(nil, 1<<3, "Read error!")
			return IMAGE_FILETYPE_UNKNOWN
		}
		if !(memcmp(filetype+8, PhpSigWebp, 4)) {
			return IMAGE_FILETYPE_WEBP
		} else {
			return IMAGE_FILETYPE_UNKNOWN
		}
	}
	if streams._phpStreamRead(stream, filetype+3, 1) != 1 {
		core.PhpErrorDocref(nil, 1<<3, "Read error!")
		return IMAGE_FILETYPE_UNKNOWN
	}

	/* BYTES READ: 4 */

	if !(memcmp(filetype, PhpSigTifIi, 4)) {
		return IMAGE_FILETYPE_TIFF_II
	} else if !(memcmp(filetype, PhpSigTifMm, 4)) {
		return IMAGE_FILETYPE_TIFF_MM
	} else if !(memcmp(filetype, PhpSigIff, 4)) {
		return IMAGE_FILETYPE_IFF
	} else if !(memcmp(filetype, PhpSigIco, 4)) {
		return IMAGE_FILETYPE_ICO
	}

	/* WBMP may be smaller than 12 bytes, so delay error */

	twelve_bytes_read = streams._phpStreamRead(stream, filetype+4, 8) == 8

	/* BYTES READ: 12 */

	if twelve_bytes_read != 0 && !(memcmp(filetype, PhpSigJp2, 12)) {
		return IMAGE_FILETYPE_JP2
	}

	/* AFTER ALL ABOVE FAILED */

	if PhpGetWbmp(stream, nil, 1) != 0 {
		return IMAGE_FILETYPE_WBMP
	}
	if twelve_bytes_read == 0 {
		core.PhpErrorDocref(nil, 1<<3, "Read error!")
		return IMAGE_FILETYPE_UNKNOWN
	}
	if PhpGetXbm(stream, nil) != 0 {
		return IMAGE_FILETYPE_XBM
	}
	return IMAGE_FILETYPE_UNKNOWN
}

/* }}} */

func PhpGetimagesizeFromStream(stream *core.PhpStream, info *zend.Zval, execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var itype int = 0
	var result *Gfxinfo = nil
	if stream == nil {
		return_value.u1.type_info = 2
		return
	}
	itype = PhpGetimagetype(stream, nil)
	switch itype {
	case IMAGE_FILETYPE_GIF:
		result = PhpHandleGif(stream)
		break
	case IMAGE_FILETYPE_JPEG:
		if info != nil {
			result = PhpHandleJpeg(stream, info)
		} else {
			result = PhpHandleJpeg(stream, nil)
		}
		break
	case IMAGE_FILETYPE_PNG:
		result = PhpHandlePng(stream)
		break
	case IMAGE_FILETYPE_SWF:
		result = PhpHandleSwf(stream)
		break
	case IMAGE_FILETYPE_SWC:
		core.PhpErrorDocref(nil, 1<<3, "The image is a compressed SWF file, but you do not have a static version of the zlib extension enabled")
		break
	case IMAGE_FILETYPE_PSD:
		result = PhpHandlePsd(stream)
		break
	case IMAGE_FILETYPE_BMP:
		result = PhpHandleBmp(stream)
		break
	case IMAGE_FILETYPE_TIFF_II:
		result = PhpHandleTiff(stream, nil, 0)
		break
	case IMAGE_FILETYPE_TIFF_MM:
		result = PhpHandleTiff(stream, nil, 1)
		break
	case IMAGE_FILETYPE_JPC:
		result = PhpHandleJpc(stream)
		break
	case IMAGE_FILETYPE_JP2:
		result = PhpHandleJp2(stream)
		break
	case IMAGE_FILETYPE_IFF:
		result = PhpHandleIff(stream)
		break
	case IMAGE_FILETYPE_WBMP:
		result = PhpHandleWbmp(stream)
		break
	case IMAGE_FILETYPE_XBM:
		result = PhpHandleXbm(stream)
		break
	case IMAGE_FILETYPE_ICO:
		result = PhpHandleIco(stream)
		break
	case IMAGE_FILETYPE_WEBP:
		result = PhpHandleWebp(stream)
		break
	default:

	case IMAGE_FILETYPE_UNKNOWN:
		break
	}
	if result != nil {
		var temp []byte
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		zend.AddIndexLong(return_value, 0, result.GetWidth())
		zend.AddIndexLong(return_value, 1, result.GetHeight())
		zend.AddIndexLong(return_value, 2, itype)
		core.ApPhpSnprintf(temp, g.SizeOf("temp"), "width=\"%d\" height=\"%d\"", result.GetWidth(), result.GetHeight())
		zend.AddIndexString(return_value, 3, temp)
		if result.GetBits() != 0 {
			zend.AddAssocLongEx(return_value, "bits", strlen("bits"), result.GetBits())
		}
		if result.GetChannels() != 0 {
			zend.AddAssocLongEx(return_value, "channels", strlen("channels"), result.GetChannels())
		}
		zend.AddAssocStringEx(return_value, "mime", strlen("mime"), (*byte)(PhpImageTypeToMimeType(itype)))
		zend._efree(result)
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

// #define FROM_DATA       0

// #define FROM_PATH       1

func PhpGetimagesizeFromAny(execute_data *zend.ZendExecuteData, return_value *zend.Zval, mode int) {
	var info *zend.Zval = nil
	var stream *core.PhpStream = nil
	var input *byte
	var input_len int
	var argc int = execute_data.This.u2.num_args
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &input, &input_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &info, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if mode == 1 && strlen(input) != size_t(input_len) {
		core.PhpErrorDocref(nil, 1<<1, "Invalid path")
		return
	}
	if argc == 2 {
		info = zend.ZendTryArrayInit(info)
		if info == nil {
			return
		}
	}
	if mode == 1 {
		stream = streams._phpStreamOpenWrapperEx(input, "rb", 0x10|0x8|0x0, nil, nil)
	} else {
		stream = streams._phpStreamMemoryOpen(0x1, input, input_len)
	}
	if stream == nil {
		return_value.u1.type_info = 2
		return
	}
	PhpGetimagesizeFromStream(stream, info, execute_data, return_value)
	streams._phpStreamFree(stream, 1|2)
}

/* }}} */

func ZifGetimagesize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpGetimagesizeFromAny(execute_data, return_value, 1)
}

/* }}} */

func ZifGetimagesizefromstring(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpGetimagesizeFromAny(execute_data, return_value, 0)
}

/* }}} */
