// <<generate>>

package standard

// Source: <ext/standard/php_image.h>

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

/* {{{ enum image_filetype
   This enum is used to have ext/standard/image.c and ext/exif/exif.c use
   the same constants for file types.
*/

type ImageFiletype = int

const (
	IMAGE_FILETYPE_UNKNOWN = 0
	IMAGE_FILETYPE_GIF     = 1
	IMAGE_FILETYPE_JPEG
	IMAGE_FILETYPE_PNG
	IMAGE_FILETYPE_SWF
	IMAGE_FILETYPE_PSD
	IMAGE_FILETYPE_BMP
	IMAGE_FILETYPE_TIFF_II
	IMAGE_FILETYPE_TIFF_MM
	IMAGE_FILETYPE_JPC
	IMAGE_FILETYPE_JP2
	IMAGE_FILETYPE_JPX
	IMAGE_FILETYPE_JB2
	IMAGE_FILETYPE_SWC
	IMAGE_FILETYPE_IFF
	IMAGE_FILETYPE_WBMP
	IMAGE_FILETYPE_XBM
	IMAGE_FILETYPE_ICO
	IMAGE_FILETYPE_WEBP
	IMAGE_FILETYPE_COUNT
)
