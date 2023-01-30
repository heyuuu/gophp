// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	r "sik/runtime"
	"sik/zend"
)

func ZmStartupImagetypes(type_ int, module_number int) int {
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_GIF", IMAGE_FILETYPE_GIF, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_JPEG", IMAGE_FILETYPE_JPEG, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_PNG", IMAGE_FILETYPE_PNG, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_SWF", IMAGE_FILETYPE_SWF, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_PSD", IMAGE_FILETYPE_PSD, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_BMP", IMAGE_FILETYPE_BMP, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_TIFF_II", IMAGE_FILETYPE_TIFF_II, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_TIFF_MM", IMAGE_FILETYPE_TIFF_MM, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_JPC", IMAGE_FILETYPE_JPC, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_JP2", IMAGE_FILETYPE_JP2, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_JPX", IMAGE_FILETYPE_JPX, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_JB2", IMAGE_FILETYPE_JB2, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_IFF", IMAGE_FILETYPE_IFF, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_WBMP", IMAGE_FILETYPE_WBMP, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_JPEG2000", IMAGE_FILETYPE_JPC, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_XBM", IMAGE_FILETYPE_XBM, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_ICO", IMAGE_FILETYPE_ICO, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_WEBP", IMAGE_FILETYPE_WEBP, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_UNKNOWN", IMAGE_FILETYPE_UNKNOWN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("IMAGETYPE_COUNT", IMAGE_FILETYPE_COUNT, zend.CONST_CS|zend.CONST_PERSISTENT)
	return zend.SUCCESS
}
func PhpHandleGif(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var dim []uint8
	if core.PhpStreamSeek(stream, 3, r.SEEK_CUR) != 0 {
		return nil
	}
	if core.PhpStreamRead(stream, (*byte)(dim), b.SizeOf("dim")) != b.SizeOf("dim") {
		return nil
	}
	result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
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
func PhpHandlePsd(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var dim []uint8
	if core.PhpStreamSeek(stream, 11, r.SEEK_CUR) != 0 {
		return nil
	}
	if core.PhpStreamRead(stream, (*byte)(dim), b.SizeOf("dim")) != b.SizeOf("dim") {
		return nil
	}
	result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
	result.SetHeight((uint(dim[0]) << 24) + (uint(dim[1]) << 16) + (uint(dim[2]) << 8) + uint(dim[3]))
	result.SetWidth((uint(dim[4]) << 24) + (uint(dim[5]) << 16) + (uint(dim[6]) << 8) + uint(dim[7]))
	return result
}
func PhpHandleBmp(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var dim []uint8
	var size int
	if core.PhpStreamSeek(stream, 11, r.SEEK_CUR) != 0 {
		return nil
	}
	if core.PhpStreamRead(stream, (*byte)(dim), b.SizeOf("dim")) != b.SizeOf("dim") {
		return nil
	}
	size = (uint(dim[3]) << 24) + (uint(dim[2]) << 16) + (uint(dim[1]) << 8) + uint(dim[0])
	if size == 12 {
		result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
		result.SetWidth((uint(dim[5]) << 8) + uint(dim[4]))
		result.SetHeight((uint(dim[7]) << 8) + uint(dim[6]))
		result.SetBits(uint(dim[11]))
	} else if size > 12 && (size <= 64 || size == 108 || size == 124) {
		result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
		result.SetWidth((uint(dim[7]) << 24) + (uint(dim[6]) << 16) + (uint(dim[5]) << 8) + uint(dim[4]))
		result.SetHeight((uint(dim[11]) << 24) + (uint(dim[10]) << 16) + (uint(dim[9]) << 8) + uint(dim[8]))
		result.SetHeight(abs(int32(result.GetHeight())))
		result.SetBits((uint(dim[15]) << 8) + uint(dim[14]))
	} else {
		return nil
	}
	return result
}
func PhpSwfGetBits(buffer *uint8, pos uint, count uint) unsigned__long__int {
	var loop uint
	var result unsigned__long__int = 0
	for loop = pos; loop < pos+count; loop++ {
		result = result + ((buffer[loop/8]>>7-loop%8&0x1)<<count - (loop - pos) - 1)
	}
	return result
}
func PhpHandleSwf(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var bits long
	var a []uint8
	if core.PhpStreamSeek(stream, 5, r.SEEK_CUR) != 0 {
		return nil
	}
	if core.PhpStreamRead(stream, (*byte)(a), b.SizeOf("a")) != b.SizeOf("a") {
		return nil
	}
	result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
	bits = PhpSwfGetBits(a, 0, 5)
	result.SetWidth((PhpSwfGetBits(a, 5+bits, bits) - PhpSwfGetBits(a, 5, bits)) / 20)
	result.SetHeight((PhpSwfGetBits(a, 5+3*bits, bits) - PhpSwfGetBits(a, 5+2*bits, bits)) / 20)
	result.SetBits(0)
	result.SetChannels(0)
	return result
}
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

	if core.PhpStreamSeek(stream, 8, r.SEEK_CUR) != 0 {
		return nil
	}
	if core.PhpStreamRead(stream, (*byte)(dim), b.SizeOf("dim")) < b.SizeOf("dim") {
		return nil
	}
	result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
	result.SetWidth((uint(dim[0]) << 24) + (uint(dim[1]) << 16) + (uint(dim[2]) << 8) + uint(dim[3]))
	result.SetHeight((uint(dim[4]) << 24) + (uint(dim[5]) << 16) + (uint(dim[6]) << 8) + uint(dim[7]))
	result.SetBits(uint(dim[8]))
	return result
}
func PhpRead2(stream *core.PhpStream) uint16 {
	var a []uint8

	/* return 0 if we couldn't read enough data */

	if core.PhpStreamRead(stream, (*byte)(a), b.SizeOf("a")) < b.SizeOf("a") {
		return 0
	}
	return (uint16(a[0]) << 8) + uint16(a[1])
}
func PhpNextMarker(stream *core.PhpStream, last_marker int, ff_read int) uint {
	var a int = 0
	var marker int

	/* get marker byte, swallowing possible padding                           */

	if ff_read == 0 {
		var extraneous int = 0
		for b.Assign(&marker, core.PhpStreamGetc(stream)) != 0xff {
			if marker == r.EOF {
				return M_EOI
			}
			extraneous++
		}
		if extraneous != 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "corrupt JPEG data: %zu extraneous bytes before marker", extraneous)
		}
	}
	a = 1
	for {
		if b.Assign(&marker, core.PhpStreamGetc(stream)) == r.EOF {
			return M_EOI
		}
		a++
		if marker != 0xff {
			break
		}
	}
	if a < 2 {
		return M_EOI
	}
	return uint(marker)
}
func PhpSkipVariable(stream *core.PhpStream) int {
	var length zend.ZendOffT = uint(PhpRead2(stream))
	if length < 2 {
		return 0
	}
	length = length - 2
	core.PhpStreamSeek(stream, zend.ZendLong(length), r.SEEK_CUR)
	return 1
}
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
	buffer = zend.Emalloc(int(length))
	if core.PhpStreamRead(stream, buffer, int(length)) != length {
		zend.Efree(buffer)
		return 0
	}
	core.Snprintf(markername, b.SizeOf("markername"), "APP%d", marker-M_APP0)
	if b.Assign(&tmp, info.GetArr().KeyFind(b.CastStr(markername, strlen(markername)))) == nil {

		/* XXX we only catch the 1st tag of it's kind! */

		zend.AddAssocStringl(info, markername, buffer, length)

		/* XXX we only catch the 1st tag of it's kind! */

	}
	zend.Efree(buffer)
	return 1
}
func PhpHandleJpeg(stream *core.PhpStream, info *zend.Zval) *Gfxinfo {
	var result *Gfxinfo = nil
	var marker uint = M_PSEUDO
	var length uint16
	var ff_read uint16 = 1
	for {
		marker = PhpNextMarker(stream, marker, ff_read)
		ff_read = 0
		switch marker {
		case M_SOF0:

		case M_SOF1:

		case M_SOF2:

		case M_SOF3:

		case M_SOF5:

		case M_SOF6:

		case M_SOF7:

		case M_SOF9:

		case M_SOF10:

		case M_SOF11:

		case M_SOF13:

		case M_SOF14:

		case M_SOF15:
			if result == nil {

				/* handle SOFn block */

				result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
				length = PhpRead2(stream)
				result.SetBits(core.PhpStreamGetc(stream))
				result.SetHeight(PhpRead2(stream))
				result.SetWidth(PhpRead2(stream))
				result.SetChannels(core.PhpStreamGetc(stream))
				if info == nil || length < 8 {
					return result
				}
				if core.PhpStreamSeek(stream, length-8, r.SEEK_CUR) != 0 {
					return result
				}
			} else {
				if PhpSkipVariable(stream) == 0 {
					return result
				}
			}
			break
		case M_APP0:

		case M_APP1:

		case M_APP2:

		case M_APP3:

		case M_APP4:

		case M_APP5:

		case M_APP6:

		case M_APP7:

		case M_APP8:

		case M_APP9:

		case M_APP10:

		case M_APP11:

		case M_APP12:

		case M_APP13:

		case M_APP14:

		case M_APP15:
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
		case M_SOS:

		case M_EOI:
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
func PhpRead4(stream *core.PhpStream) uint {
	var a []uint8

	/* just return 0 if we hit the end-of-file */

	if core.PhpStreamRead(stream, (*byte)(a), b.SizeOf("a")) != b.SizeOf("a") {
		return 0
	}
	return (uint(a[0]) << 24) + (uint(a[1]) << 16) + (uint(a[2]) << 8) + uint(a[3])
}
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

	first_marker_id = core.PhpStreamGetc(stream)

	/* Ensure that this marker is SIZ (as is mandated by the standard) */

	if first_marker_id != JPEG2000_MARKER_SIZ {
		core.PhpErrorDocref(nil, zend.E_WARNING, "JPEG2000 codestream corrupt(Expected SIZ marker not found after SOC)")
		return nil
	}
	result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
	PhpRead2(stream)
	PhpRead2(stream)
	result.SetWidth(PhpRead4(stream))
	result.SetHeight(PhpRead4(stream))
	if core.PhpStreamSeek(stream, 24, r.SEEK_CUR) != 0 {
		zend.Efree(result)
		return nil
	}
	result.SetChannels(PhpRead2(stream))
	if result.GetChannels() == 0 && core.PhpStreamEof(stream) != 0 || result.GetChannels() > 256 {
		zend.Efree(result)
		return nil
	}

	/* Collect bit depth info */

	highest_bit_depth = 0
	for i = 0; i < result.GetChannels(); i++ {
		bit_depth = core.PhpStreamGetc(stream)
		bit_depth++
		if bit_depth > highest_bit_depth {
			highest_bit_depth = bit_depth
		}
		core.PhpStreamGetc(stream)
		core.PhpStreamGetc(stream)
	}
	result.SetBits(highest_bit_depth)
	return result
}
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

		if core.PhpStreamRead(stream, any(&box_type), b.SizeOf("box_type")) != b.SizeOf("box_type") {

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

			core.PhpStreamSeek(stream, 3, r.SEEK_CUR)
			result = PhpHandleJpc(stream)
			break
		}

		/* Stop if this was the last box */

		if int(box_length <= 0) != 0 {
			break
		}

		/* Skip over LBox (Which includes both TBox and LBox itself */

		if core.PhpStreamSeek(stream, box_length-8, r.SEEK_CUR) != 0 {
			break
		}

		/* Skip over LBox (Which includes both TBox and LBox itself */

	}
	if result == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "JP2 file has no codestreams at root level")
	}
	return result
}
func PhpIfdGet16u(Short any, motorola_intel int) int {
	if motorola_intel != 0 {
		return (*uint8)(Short)[0]<<8 | (*uint8)(Short)[1]
	} else {
		return (*uint8)(Short)[1]<<8 | (*uint8)(Short)[0]
	}
}
func PhpIfdGet16s(Short any, motorola_intel int) signed__short {
	return signed__short(PhpIfdGet16u(Short, motorola_intel))
}
func PhpIfdGet32s(Long any, motorola_intel int) int {
	if motorola_intel != 0 {
		return (*byte)(Long)[0]<<24 | (*uint8)(Long)[1]<<16 | (*uint8)(Long)[2]<<8 | (*uint8)(Long)[3]<<0
	} else {
		return (*byte)(Long)[3]<<24 | (*uint8)(Long)[2]<<16 | (*uint8)(Long)[1]<<8 | (*uint8)(Long)[0]<<0
	}
}
func PhpIfdGet32u(Long any, motorola_intel int) unsigned {
	return unsigned(PhpIfdGet32s(Long, motorola_intel) & 0xffffffff)
}
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
	if core.PhpStreamRead(stream, ifd_ptr, 4) != 4 {
		return nil
	}
	ifd_addr = PhpIfdGet32u(ifd_ptr, motorola_intel)
	if core.PhpStreamSeek(stream, ifd_addr-8, r.SEEK_CUR) != 0 {
		return nil
	}
	ifd_size = 2
	ifd_data = zend.Emalloc(ifd_size)
	if core.PhpStreamRead(stream, ifd_data, 2) != 2 {
		zend.Efree(ifd_data)
		return nil
	}
	num_entries = PhpIfdGet16u(ifd_data, motorola_intel)
	dir_size = 2 + 12*num_entries + 4
	ifd_size = dir_size
	ifd_data = zend.Erealloc(ifd_data, ifd_size)
	if core.PhpStreamRead(stream, ifd_data+2, dir_size-2) != dir_size-2 {
		zend.Efree(ifd_data)
		return nil
	}

	/* now we have the directory we can look how long it should be */

	ifd_size = dir_size
	for i = 0; i < num_entries; i++ {
		dir_entry = (*uint8)(ifd_data + 2 + i*12)
		entry_tag = PhpIfdGet16u(dir_entry+0, motorola_intel)
		entry_type = PhpIfdGet16u(dir_entry+2, motorola_intel)
		switch entry_type {
		case TAG_FMT_BYTE:

		case TAG_FMT_SBYTE:
			entry_value = size_t(dir_entry[8])
			break
		case TAG_FMT_USHORT:
			entry_value = PhpIfdGet16u(dir_entry+8, motorola_intel)
			break
		case TAG_FMT_SSHORT:
			entry_value = PhpIfdGet16s(dir_entry+8, motorola_intel)
			break
		case TAG_FMT_ULONG:
			entry_value = PhpIfdGet32u(dir_entry+8, motorola_intel)
			break
		case TAG_FMT_SLONG:
			entry_value = PhpIfdGet32s(dir_entry+8, motorola_intel)
			break
		default:
			continue
		}
		switch entry_tag {
		case TAG_IMAGEWIDTH:

		case TAG_COMP_IMAGEWIDTH:
			width = entry_value
			break
		case TAG_IMAGEHEIGHT:

		case TAG_COMP_IMAGEHEIGHT:
			height = entry_value
			break
		}
	}
	zend.Efree(ifd_data)
	if width != 0 && height != 0 {

		/* not the same when in for-loop */

		result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
		result.SetHeight(height)
		result.SetWidth(width)
		result.SetBits(0)
		result.SetChannels(0)
		return result
	}
	return nil
}
func PhpHandleIff(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo
	var a []uint8
	var chunkId int
	var size int
	var width short
	var height short
	var bits short
	if core.PhpStreamRead(stream, (*byte)(a), 8) != 8 {
		return nil
	}
	if strncmp((*byte)(a+4), "ILBM", 4) && strncmp((*byte)(a+4), "PBM ", 4) {
		return nil
	}

	/* loop chunks to find BMHD chunk */

	for {
		if core.PhpStreamRead(stream, (*byte)(a), 8) != 8 {
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
			if size < 9 || core.PhpStreamRead(stream, (*byte)(a), 9) != 9 {
				return nil
			}
			width = PhpIfdGet16s(a+0, 1)
			height = PhpIfdGet16s(a+2, 1)
			bits = a[8] & 0xff
			if width > 0 && height > 0 && bits > 0 && bits < 33 {
				result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
				result.SetWidth(width)
				result.SetHeight(height)
				result.SetBits(bits)
				result.SetChannels(0)
				return result
			}
		} else {
			if core.PhpStreamSeek(stream, size, r.SEEK_CUR) != 0 {
				return nil
			}
		}

	}

	/* loop chunks to find BMHD chunk */
}
func PhpGetWbmp(stream *core.PhpStream, result **Gfxinfo, check int) int {
	var i int
	var width int = 0
	var height int = 0
	if core.PhpStreamRewind(stream) != 0 {
		return 0
	}

	/* get type */

	if core.PhpStreamGetc(stream) != 0 {
		return 0
	}

	/* skip header */

	for {
		i = core.PhpStreamGetc(stream)
		if i < 0 {
			return 0
		}
		if (i & 0x80) == 0 {
			break
		}
	}

	/* get width */

	for {
		i = core.PhpStreamGetc(stream)
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
		i = core.PhpStreamGetc(stream)
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
		result.SetWidth(width)
		result.SetHeight(height)
	}
	return IMAGE_FILETYPE_WBMP
}
func PhpHandleWbmp(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
	if PhpGetWbmp(stream, &result, 0) == 0 {
		zend.Efree(result)
		return nil
	}
	return result
}
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
	if core.PhpStreamRewind(stream) != 0 {
		return 0
	}
	for b.Assign(&fline, core.PhpStreamGets(stream, nil, 0)) != nil {
		iname = zend.Estrdup(fline)
		if sscanf(fline, "#define %s %d", iname, &value) == 2 {
			if !(b.Assign(&type_, strrchr(iname, '_'))) {
				type_ = iname
			} else {
				type_++
			}
			if !(strcmp("width", type_)) {
				width = uint(value)
				if height != 0 {
					zend.Efree(iname)
					break
				}
			}
			if !(strcmp("height", type_)) {
				height = uint(value)
				if width != 0 {
					zend.Efree(iname)
					break
				}
			}
		}
		zend.Efree(fline)
		zend.Efree(iname)
	}
	if fline != nil {
		zend.Efree(fline)
	}
	if width != 0 && height != 0 {
		if result != nil {
			*result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
			result.SetWidth(width)
			result.SetHeight(height)
		}
		return IMAGE_FILETYPE_XBM
	}
	return 0
}
func PhpHandleXbm(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo
	PhpGetXbm(stream, &result)
	return result
}
func PhpHandleIco(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var dim []uint8
	var num_icons int = 0
	if core.PhpStreamRead(stream, (*byte)(dim), 2) != 2 {
		return nil
	}
	num_icons = (uint(dim[1]) << 8) + uint(dim[0])
	if num_icons < 1 || num_icons > 255 {
		return nil
	}
	result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
	for num_icons > 0 {
		if core.PhpStreamRead(stream, (*byte)(dim), b.SizeOf("dim")) != b.SizeOf("dim") {
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
func PhpHandleWebp(stream *core.PhpStream) *Gfxinfo {
	var result *Gfxinfo = nil
	var sig []byte = []byte{'V', 'P', '8'}
	var buf []uint8
	var format byte
	if core.PhpStreamRead(stream, (*byte)(buf), 18) != 18 {
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
	result = (*Gfxinfo)(zend.Ecalloc(1, b.SizeOf("struct gfxinfo")))
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
func ZifImageTypeToMimeType(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var p_image_type zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &p_image_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	zend.ZVAL_STRING(return_value, (*byte)(PhpImageTypeToMimeType(p_image_type)))
}
func ZifImageTypeToExtension(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var image_type zend.ZendLong
	var inc_dot zend.ZendBool = 1
	var imgext *byte = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &image_type, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &inc_dot, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
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
		zend.RETVAL_STRING(&imgext[!inc_dot])
		return
	}
	zend.RETVAL_FALSE
	return
}
func PhpGetimagetype(stream *core.PhpStream, filetype *byte) int {
	var tmp []byte
	var twelve_bytes_read int
	if filetype == nil {
		filetype = tmp
	}
	if core.PhpStreamRead(stream, filetype, 3) != 3 {
		core.PhpErrorDocref(nil, zend.E_NOTICE, "Read error!")
		return IMAGE_FILETYPE_UNKNOWN
	}

	/* BYTES READ: 3 */

	if !(memcmp(filetype, PhpSigGif, 3)) {
		return IMAGE_FILETYPE_GIF
	} else if !(memcmp(filetype, PhpSigJpg, 3)) {
		return IMAGE_FILETYPE_JPEG
	} else if !(memcmp(filetype, PhpSigPng, 3)) {
		if core.PhpStreamRead(stream, filetype+3, 5) != 5 {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "Read error!")
			return IMAGE_FILETYPE_UNKNOWN
		}
		if !(memcmp(filetype, PhpSigPng, 8)) {
			return IMAGE_FILETYPE_PNG
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "PNG file corrupted by ASCII conversion")
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
		if core.PhpStreamRead(stream, filetype+3, 9) != 9 {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "Read error!")
			return IMAGE_FILETYPE_UNKNOWN
		}
		if !(memcmp(filetype+8, PhpSigWebp, 4)) {
			return IMAGE_FILETYPE_WEBP
		} else {
			return IMAGE_FILETYPE_UNKNOWN
		}
	}
	if core.PhpStreamRead(stream, filetype+3, 1) != 1 {
		core.PhpErrorDocref(nil, zend.E_NOTICE, "Read error!")
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

	twelve_bytes_read = core.PhpStreamRead(stream, filetype+4, 8) == 8

	/* BYTES READ: 12 */

	if twelve_bytes_read != 0 && !(memcmp(filetype, PhpSigJp2, 12)) {
		return IMAGE_FILETYPE_JP2
	}

	/* AFTER ALL ABOVE FAILED */

	if PhpGetWbmp(stream, nil, 1) != 0 {
		return IMAGE_FILETYPE_WBMP
	}
	if twelve_bytes_read == 0 {
		core.PhpErrorDocref(nil, zend.E_NOTICE, "Read error!")
		return IMAGE_FILETYPE_UNKNOWN
	}
	if PhpGetXbm(stream, nil) != 0 {
		return IMAGE_FILETYPE_XBM
	}
	return IMAGE_FILETYPE_UNKNOWN
}
func PhpGetimagesizeFromStream(stream *core.PhpStream, info *zend.Zval, execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var itype int = 0
	var result *Gfxinfo = nil
	if stream == nil {
		zend.RETVAL_FALSE
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
		core.PhpErrorDocref(nil, zend.E_NOTICE, "The image is a compressed SWF file, but you do not have a static version of the zlib extension enabled")
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
		zend.ArrayInit(return_value)
		zend.AddIndexLong(return_value, 0, result.GetWidth())
		zend.AddIndexLong(return_value, 1, result.GetHeight())
		zend.AddIndexLong(return_value, 2, itype)
		core.Snprintf(temp, b.SizeOf("temp"), "width=\"%d\" height=\"%d\"", result.GetWidth(), result.GetHeight())
		zend.AddIndexString(return_value, 3, temp)
		if result.GetBits() != 0 {
			zend.AddAssocLong(return_value, "bits", result.GetBits())
		}
		if result.GetChannels() != 0 {
			zend.AddAssocLong(return_value, "channels", result.GetChannels())
		}
		zend.AddAssocString(return_value, "mime", (*byte)(PhpImageTypeToMimeType(itype)))
		zend.Efree(result)
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func PhpGetimagesizeFromAny(execute_data *zend.ZendExecuteData, return_value *zend.Zval, mode int) {
	var info *zend.Zval = nil
	var stream *core.PhpStream = nil
	var input *byte
	var input_len int
	var argc int = zend.ZEND_NUM_ARGS()
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &input, &input_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &info, 0)
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
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
	if mode == FROM_PATH && zend.CHECK_NULL_PATH(input, input_len) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid path")
		return
	}
	if argc == 2 {
		info = zend.ZendTryArrayInit(info)
		if info == nil {
			return
		}
	}
	if mode == FROM_PATH {
		stream = core.PhpStreamOpenWrapper(input, "rb", core.STREAM_MUST_SEEK|core.REPORT_ERRORS|core.IGNORE_PATH, nil)
	} else {
		stream = core.PhpStreamMemoryOpen(core.TEMP_STREAM_READONLY, input, input_len)
	}
	if stream == nil {
		zend.RETVAL_FALSE
		return
	}
	PhpGetimagesizeFromStream(stream, info, execute_data, return_value)
	core.PhpStreamClose(stream)
}
func ZifGetimagesize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpGetimagesizeFromAny(execute_data, return_value, FROM_PATH)
}
func ZifGetimagesizefromstring(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpGetimagesizeFromAny(execute_data, return_value, FROM_DATA)
}
