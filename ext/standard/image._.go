// <<generate>>

package standard

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

/* {{{ PHP_MINIT_FUNCTION(imagetypes)
 * Register IMAGETYPE_<xxx> constants used by GetImageSize(), image_type_to_mime_type, ext/exif */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ php_handle_swf
 */

/* }}} */

/* }}} */

const M_SOF0 = 0xc0
const M_SOF1 = 0xc1
const M_SOF2 = 0xc2
const M_SOF3 = 0xc3
const M_SOF5 = 0xc5
const M_SOF6 = 0xc6
const M_SOF7 = 0xc7
const M_SOF9 = 0xc9
const M_SOF10 = 0xca
const M_SOF11 = 0xcb
const M_SOF13 = 0xcd
const M_SOF14 = 0xce
const M_SOF15 = 0xcf
const M_SOI = 0xd8
const M_EOI = 0xd9
const M_SOS = 0xda
const M_APP0 = 0xe0
const M_APP1 = 0xe1
const M_APP2 = 0xe2
const M_APP3 = 0xe3
const M_APP4 = 0xe4
const M_APP5 = 0xe5
const M_APP6 = 0xe6
const M_APP7 = 0xe7
const M_APP8 = 0xe8
const M_APP9 = 0xe9
const M_APP10 = 0xea
const M_APP11 = 0xeb
const M_APP12 = 0xec
const M_APP13 = 0xed
const M_APP14 = 0xee
const M_APP15 = 0xef
const M_COM = 0xfe
const M_PSEUDO = 0xffd8

/* {{{ php_read2
 */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

const JPEG2000_MARKER_PREFIX = 0xff
const JPEG2000_MARKER_SOC = 0x4f
const JPEG2000_MARKER_SOT = 0x90
const JPEG2000_MARKER_SOD = 0x93
const JPEG2000_MARKER_EOC = 0xd9
const JPEG2000_MARKER_SIZ = 0x51
const JPEG2000_MARKER_COD = 0x52
const JPEG2000_MARKER_COC = 0x53
const JPEG2000_MARKER_RGN = 0x5e
const JPEG2000_MARKER_QCD = 0x5c
const JPEG2000_MARKER_QCC = 0x5d
const JPEG2000_MARKER_POC = 0x5f
const JPEG2000_MARKER_TLM = 0x55
const JPEG2000_MARKER_PLM = 0x57
const JPEG2000_MARKER_PLT = 0x58
const JPEG2000_MARKER_PPM = 0x60
const JPEG2000_MARKER_PPT = 0x61
const JPEG2000_MARKER_SOP = 0x91
const JPEG2000_MARKER_EPH = 0x92
const JPEG2000_MARKER_CRG = 0x63
const JPEG2000_MARKER_COM = 0x64

/* }}} */

/* }}} */

/* }}} */

var PhpTiffBytesPerFormat []int = []int{0, 1, 1, 2, 4, 8, 1, 1, 2, 4, 8, 4, 8}

/* uncompressed only */

const TAG_IMAGEWIDTH = 0x100
const TAG_IMAGEHEIGHT = 0x101

/* compressed images only */

const TAG_COMP_IMAGEWIDTH = 0xa002
const TAG_COMP_IMAGEHEIGHT = 0xa003
const TAG_FMT_BYTE = 1
const TAG_FMT_STRING = 2
const TAG_FMT_USHORT = 3
const TAG_FMT_ULONG = 4
const TAG_FMT_URATIONAL = 5
const TAG_FMT_SBYTE = 6
const TAG_FMT_UNDEFINED = 7
const TAG_FMT_SSHORT = 8
const TAG_FMT_SLONG = 9
const TAG_FMT_SRATIONAL = 10
const TAG_FMT_SINGLE = 11
const TAG_FMT_DOUBLE = 12

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

const FROM_DATA = 0
const FROM_PATH = 1

/* }}} */

/* }}} */

/* }}} */
