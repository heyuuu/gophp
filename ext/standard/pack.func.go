package standard

import (
	"encoding/binary"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/kits/strkit"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func packGetLong(val *types.Zval) int {
	if !val.IsLong() {
		operators.ConvertToLong(val)
	}
	return val.Long()
}
func packInt8(buf *strings.Builder, val *types.Zval) {
	num := packGetLong(val)
	buf.WriteByte(byte(num))
}

func packInt16(buf *strings.Builder, val *types.Zval, order binary.ByteOrder) {
	num := packGetLong(val)
	_ = binary.Write(buf, order, int16(num))
}
func packInt32(buf *strings.Builder, val *types.Zval, order binary.ByteOrder) {
	num := packGetLong(val)
	_ = binary.Write(buf, order, int32(num))
}
func packInt64(buf *strings.Builder, val *types.Zval, order binary.ByteOrder) {
	num := packGetLong(val)
	_ = binary.Write(buf, order, int64(num))
}
func packInt(buf *strings.Builder, val *types.Zval, order binary.ByteOrder) {
	num := packGetLong(val)
	_ = binary.Write(buf, order, int(num))
}

func packGetDouble(val *types.Zval) float64 {
	return operators.ZvalGetDouble(val)
}
func packFloat32(buf *strings.Builder, val *types.Zval, order binary.ByteOrder) {
	num := packGetDouble(val)
	_ = binary.Write(buf, order, float32(num))
}
func packFloat64(buf *strings.Builder, val *types.Zval, order binary.ByteOrder) {
	num := packGetDouble(val)
	_ = binary.Write(buf, order, float64(num))
}

func PhpPackReverseInt32(arg uint32) uint32 {
	var result uint32
	result = (arg&0xff)<<24 | (arg&0xff00)<<8 | arg>>8&0xff00 | arg>>24&0xff
	return result
}
func PhpPackReverseInt64(arg uint64) uint64 {
	var tmp struct /* union */ {
		i  uint64
		ul []uint32
	}
	var result struct /* union */ {
		i  uint64
		ul []uint32
	}
	tmp.i = arg
	result.ul[0] = PhpPackReverseInt32(tmp.ul[1])
	result.ul[1] = PhpPackReverseInt32(tmp.ul[0])
	return result.i
}
func PhpPackParseFloat(is_little_endian int, src any) float {
	var m struct /* union */ {
		f float
		i uint32
	}
	memcpy(m.i, src, b.SizeOf("float"))
	if is_little_endian == 0 {
		m.i = PhpPackReverseInt32(m.i)
	}
	return m.f
}
func PhpPackParseDouble(is_little_endian int, src any) float64 {
	var m struct /* union */ {
		d float64
		i uint64
	}
	memcpy(m.i, src, b.SizeOf("double"))
	if is_little_endian == 0 {
		m.i = PhpPackReverseInt64(m.i)
	}
	return m.d
}
func ZifPack(format_ string, _ zpp.Opt, args []*types.Zval) (string, bool) {
	var num_args = len(args)
	var i int
	var currentarg int
	var formatcodes []byte
	var formatargs []int
	var formatcount = 0
	var outputpos = 0
	var outputsize = 0

	/* We have a maximum of <formatlen> format codes to deal with */
	formatcodes = make([]byte, len(format_))
	formatargs = make([]int, len(format_))
	currentarg = 0

	/* Preprocess format into formatcodes and formatargs */
	for i = 0; i < len(format_); formatcount++ {
		var code = format_[lang.PostInc(&i)]
		var arg = 1

		/* Handle format arguments if any */
		if i < len(format_) {
			var c = format_[i]
			if c == '*' {
				arg = -1
				i++
			} else if c >= '0' && c <= '9' {
				arg = atoi(&format_[i])
				for format_[i] >= '0' && format_[i] <= '9' && i < len(format_) {
					i++
				}
			}
		}

		/* Handle special arg '*' for all codes and check argv overflows */

		switch int(code) {
		case 'x', 'X', '@':
			if arg < 0 {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: '*' ignored", code))
				arg = 1
			}
		case 'a', 'A', 'Z', 'h', 'H':
			if currentarg >= num_args {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: not enough arguments", code))
				return "", false
			}
			if arg < 0 {
				if operators.TryConvertToString(args[currentarg]) == 0 {
					return "", false
				}
				arg = args[currentarg].StringEx().GetLen()
				if code == 'Z' {
					/* add one because Z is always NUL-terminated:
					 * pack("Z*", "aa") === "aa\0"
					 * pack("Z2", "aa") === "a\0" */
					arg++
				}
			}
			currentarg++
		case 'q', 'Q', 'J', 'P', 'c', 'C', 's', 'S', 'i', 'I', 'l', 'L', 'n', 'N', 'v', 'V', 'f', 'g', 'G', 'd', 'e', 'E':
			if arg < 0 {
				arg = num_args - currentarg
			}
			if currentarg > core.INT_MAX-arg {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: too few arguments", code))
				return "", false
			}
			currentarg += arg
			if currentarg > num_args {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: too few arguments", code))
				return "", false
			}
		default:
			core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: unknown format code", code))
			return "", false
		}
		formatcodes[formatcount] = code
		formatargs[formatcount] = arg
	}
	if currentarg < num_args {
		core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("%d arguments unused", num_args-currentarg))
	}

	/* Calculate output length and upper bound while processing*/

	for i = 0; i < formatcount; i++ {
		var code = int(formatcodes[i])
		var arg = formatargs[i]
		switch int(code) {
		case 'h', 'H':
			if (arg+arg%2)/2 < 0 || (core.INT_MAX-outputpos)/int(1) < (arg+arg%2)/2 {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: integer overflow in format string", code))
				return "", false
			}
			outputpos += (arg + arg%2) / 2 * 1
		case 'a', 'A', 'Z', 'c', 'C', 'x':
			if arg < 0 || (core.INT_MAX-outputpos)/int(1) < arg {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: integer overflow in format string", code))
				return "", false
			}
			outputpos += arg * 1
		case 's', 'S', 'n', 'v':
			if arg < 0 || (core.INT_MAX-outputpos)/int(2) < arg {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: integer overflow in format string", code))
				return "", false
			}
			outputpos += arg * 2
		case 'i', 'I':
			if arg < 0 || (core.INT_MAX-outputpos)/int(b.SizeOf("int")) < arg {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: integer overflow in format string", code))
				return "", false
			}
			outputpos += arg * b.SizeOf("int")
		case 'l', 'L', 'N', 'V':
			if arg < 0 || (core.INT_MAX-outputpos)/int(4) < arg {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: integer overflow in format string", code))
				return "", false
			}
			outputpos += arg * 4
		case 'q', 'Q', 'J', 'P':
			if arg < 0 || (core.INT_MAX-outputpos)/int(8) < arg {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: integer overflow in format string", code))
				return "", false
			}
			outputpos += arg * 8
		case 'f', 'g', 'G':
			if arg < 0 || (core.INT_MAX-outputpos)/int(b.SizeOf("float")) < arg {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: integer overflow in format string", code))
				return "", false
			}
			outputpos += arg * b.SizeOf("float")
		case 'd', 'e', 'E':
			if arg < 0 || (core.INT_MAX-outputpos)/int(b.SizeOf("double")) < arg {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: integer overflow in format string", code))
				return "", false
			}
			outputpos += arg * b.SizeOf("double")
		case 'X':
			outputpos -= arg
			if outputpos < 0 {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: outside of string", code))
				outputpos = 0
			}
		case '@':
			outputpos = arg
		}
		if outputsize < outputpos {
			outputsize = outputpos
		}
	}

	var buf strings.Builder
	var output *types.String

	outputpos = 0
	currentarg = 0

	/* Do actual packing */

	for i = 0; i < formatcount; i++ {
		var code = int(formatcodes[i])
		var arg = formatargs[i]
		switch int(code) {
		case 'a':
			str := operators.ZvalGetStrVal(args[lang.PostInc(&currentarg)])
			str = strkit.PadRight(str, arg, '\x00')[:arg]
			buf.WriteString(str)
		case 'A':
			str := operators.ZvalGetStrVal(args[lang.PostInc(&currentarg)])
			str = strkit.PadRight(str, arg, ' ')[:arg]
			buf.WriteString(str)
		case 'Z':
			str := operators.ZvalGetStrVal(args[lang.PostInc(&currentarg)])
			if arg <= 0 {
				str = "\x00"
			} else {
				str = strkit.PadRight(str, arg-1, '\x00')[:arg-1] + "\x00"
			}
			buf.WriteString(str)
		case 'h', 'H':
			var nibbleshift = lang.Cond(code == 'h', 0, 4)
			var first = 1
			var str = operators.ZvalGetString(args[lang.PostInc(&currentarg)])
			var v *byte = str.GetVal()
			outputpos--
			if arg > str.GetLen() {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: not enough characters in string", code))
				arg = str.GetLen()
			}
			for lang.PostDec(&arg) > 0 {
				var n = lang.PostInc(&(*v))
				if n >= '0' && n <= '9' {
					n -= '0'
				} else if n >= 'A' && n <= 'F' {
					n -= 'A' - 10
				} else if n >= 'a' && n <= 'f' {
					n -= 'a' - 10
				} else {
					core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: illegal hex digit %c", code, n))
					n = 0
				}
				if lang.PostDec(&first) {
					output.GetStr()[lang.PreInc(&outputpos)] = 0
				} else {
					first = 1
				}
				output.GetStr()[outputpos] |= n << nibbleshift
				nibbleshift = nibbleshift + 4&7
			}
			outputpos++
		case 'c', 'C':
			for lang.PostDec(&arg) > 0 {
				packInt8(&buf, args[lang.PostInc(&currentarg)])
			}
		case 's', 'S', 'n', 'v':
			order := machineEndian
			if code == 'n' {
				order = bigEndian
			} else if code == 'v' {
				order = litteEndian
			}
			for lang.PostDec(&arg) > 0 {
				packInt16(&buf, args[lang.PostInc(&currentarg)], order)
			}
		case 'i', 'I':
			for lang.PostDec(&arg) > 0 {
				packInt(&buf, args[lang.PostInc(&currentarg)], machineEndian)
			}
		case 'l', 'L', 'N', 'V':
			order := machineEndian
			if code == 'N' {
				order = bigEndian
			} else if code == 'V' {
				order = litteEndian
			}
			for lang.PostDec(&arg) > 0 {
				packInt32(&buf, args[lang.PostInc(&currentarg)], order)
			}
		case 'q', 'Q', 'J', 'P':
			order := machineEndian
			if code == 'J' {
				order = bigEndian
			} else if code == 'P' {
				order = litteEndian
			}
			for lang.PostDec(&arg) > 0 {
				packInt64(&buf, args[lang.PostInc(&currentarg)], order)
			}
		case 'f':
			for lang.PostDec(&arg) > 0 {
				packFloat32(&buf, args[lang.PostInc(&currentarg)], machineEndian)
			}
		case 'g':
			/* pack little endian float */
			for lang.PostDec(&arg) > 0 {
				packFloat32(&buf, args[lang.PostInc(&currentarg)], litteEndian)
			}
		case 'G':
			/* pack big endian float */
			for lang.PostDec(&arg) > 0 {
				packFloat32(&buf, args[lang.PostInc(&currentarg)], bigEndian)
			}
		case 'd':
			for lang.PostDec(&arg) > 0 {
				packFloat64(&buf, args[lang.PostInc(&currentarg)], machineEndian)
			}
		case 'e':
			/* pack little endian double */
			for lang.PostDec(&arg) > 0 {
				packFloat64(&buf, args[lang.PostInc(&currentarg)], litteEndian)
			}
		case 'E':
			/* pack big endian double */
			for lang.PostDec(&arg) > 0 {
				packFloat64(&buf, args[lang.PostInc(&currentarg)], bigEndian)
			}
		case 'x':
			if arg > 0 {
				buf.Write(make([]byte, arg))
			}
		case 'X':
			// 向前裁剪 arg 个字符
			newLen := buf.Len() - arg
			if newLen <= 0 {
				buf.Reset()
			} else {
				leftStr := buf.String()[:newLen]
				buf.Reset()
				buf.WriteString(leftStr)
			}
		case '@':
			// 调整长度到 arg
			if arg > buf.Len() {
				buf.Write(make([]byte, arg-buf.Len()))
			} else if arg < buf.Len() {
				leftStr := buf.String()[:arg]
				buf.Reset()
				buf.WriteString(leftStr)
			}
		}
	}
	return buf.String(), true
}
func PhpUnpack(data *byte, size int, issigned int, map_ *int) zend.ZendLong {
	var result zend.ZendLong
	var cresult = (*byte)(&result)
	var i int
	if issigned != 0 {
		result = -1
	} else {
		result = 0
	}
	for i = 0; i < size; i++ {
		*data++
		cresult[map_[i]] = (*data) - 1
	}
	return result
}
func ZifUnpack(executeData zpp.Ex, return_value zpp.Ret, format *types.Zval, input *types.Zval, _ zpp.Opt, offset *types.Zval) {
	var format *byte
	var input *byte
	var formatarg *types.String
	var inputarg *types.String
	var formatlen zend.ZendLong
	var inputpos zend.ZendLong
	var inputlen zend.ZendLong
	var i int
	var offset = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			formatarg = fp.ParseStr()
			inputarg = fp.ParseStr()
			fp.StartOptional()
			offset = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	format = formatarg.GetVal()
	formatlen = formatarg.GetLen()
	input = inputarg.GetVal()
	inputlen = inputarg.GetLen()
	inputpos = 0
	if offset < 0 || offset > inputlen {
		core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Offset %d is out of input range", offset))
		return_value.SetFalse()
		return
	}
	input += offset
	inputlen -= offset
	zend.ArrayInit(return_value)
	for lang.PostDec(&formatlen) > 0 {
		var type_ byte = *(lang.PostInc(&format))
		var c byte
		var arg = 1
		var argb int
		var name *byte
		var namelen int
		var size = 0

		/* Handle format arguments if any */

		if formatlen > 0 {
			c = *format
			if c >= '0' && c <= '9' {
				arg = atoi(format)
				for formatlen > 0 && (*format) >= '0' && (*format) <= '9' {
					format++
					formatlen--
				}
			} else if c == '*' {
				arg = -1
				format++
				formatlen--
			}
		}

		/* Get of new value in array */

		name = format
		argb = arg
		for formatlen > 0 && (*format) != '/' {
			formatlen--
			format++
		}
		namelen = format - name
		if namelen > 200 {
			namelen = 200
		}
		switch int(type_) {
		case 'X':
			size = -1
			if arg < 0 {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: '*' ignored", type_))
				arg = 1
			}
		case '@':
			size = 0
		case 'a', 'A', 'Z':
			size = arg
			arg = 1
		case 'h', 'H':
			if arg > 0 {
				size = (arg + arg%2) / 2
			} else {
				size = arg
			}
			arg = 1
		case 'c', 'C', 'x':
			size = 1
		case 's', 'S', 'n', 'v':
			size = 2
		case 'i', 'I':
			size = b.SizeOf("int")
		case 'l', 'L', 'N', 'V':
			size = 4
		case 'q', 'Q', 'J', 'P':
			size = 8
		case 'f', 'g', 'G':
			size = b.SizeOf("float")
		case 'd', 'e', 'E':
			size = b.SizeOf("double")
		default:
			core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Invalid format type %c", type_))
			return_value.Array().Destroy()
			return_value.SetFalse()
			return
		}
		if size != 0 && size != -1 && size < 0 {
			core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: integer overflow", type_))
			return_value.Array().Destroy()
			return_value.SetFalse()
			return
		}

		/* Do actual unpacking */

		for i = 0; i != arg; i++ {

			/* Space for name + number, safe as namelen is ensured <= 200 */

			var n []byte
			if arg != 1 || namelen == 0 {

				/* Need to add element number to name */

				core.Snprintf(n, b.SizeOf("n"), "%.*s%d", namelen, name, i+1)

				/* Need to add element number to name */

			} else {

				/* Truncate name to next format code or end of string */

				core.Snprintf(n, b.SizeOf("n"), "%.*s", namelen, name)

				/* Truncate name to next format code or end of string */

			}
			if size != 0 && size != -1 && core.INT_MAX-size+1 < inputpos {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: integer overflow", type_))
				return_value.Array().Destroy()
				return_value.SetFalse()
				return
			}
			if inputpos+size <= inputlen {
				switch int(type_) {
				case 'a':

					/* a will not strip any trailing whitespace or null padding */

					var len_ = inputlen - inputpos

					/* If size was given take minimum of len and size */

					if size >= 0 && len_ > size {
						len_ = size
					}
					size = len_
					zend.AddAssocStringl(return_value, n, &input[inputpos], len_)
				case 'A':

					/* A will strip any trailing whitespace */

					var padn byte = '0'
					var pads byte = ' '
					var padt byte = '\t'
					var padc byte = '\r'
					var padl byte = '\n'
					var len_ = inputlen - inputpos

					/* If size was given take minimum of len and size */

					if size >= 0 && len_ > size {
						len_ = size
					}
					size = len_

					/* Remove trailing white space and nulls chars from unpacked data */

					for lang.PreDec(&len_) >= 0 {
						if input[inputpos+len_] != padn && input[inputpos+len_] != pads && input[inputpos+len_] != padt && input[inputpos+len_] != padc && input[inputpos+len_] != padl {
							break
						}
					}
					zend.AddAssocStringl(return_value, n, &input[inputpos], len_+1)
				case 'Z':

					/* Z will strip everything after the first null character */

					var pad byte = '0'
					var s zend.ZendLong
					var len_ = inputlen - inputpos

					/* If size was given take minimum of len and size */

					if size >= 0 && len_ > size {
						len_ = size
					}
					size = len_

					/* Remove everything after the first null */

					for s = 0; s < len_; s++ {
						if input[inputpos+s] == pad {
							break
						}
					}
					len_ = s
					zend.AddAssocStringl(return_value, n, &input[inputpos], len_)
				case 'h':
					fallthrough
				case 'H':
					var len_ = (inputlen - inputpos) * 2
					var nibbleshift = lang.Cond(type_ == 'h', 0, 4)
					var first = 1
					var ipos zend.ZendLong
					var opos zend.ZendLong

					/* If size was given take minimum of len and size */

					if size >= 0 && len_ > size*2 {
						len_ = size * 2
					}
					if len_ > 0 && argb > 0 {
						len_ -= argb % 2
					}

					buf := make([]byte, len_)
					opos = 0
					ipos = opos
					for ; opos < len_; opos++ {
						var cc byte = input[inputpos+ipos] >> nibbleshift & 0xf
						if cc < 10 {
							cc += '0'
						} else {
							cc += 'a' - 10
						}
						buf[opos] = cc
						nibbleshift = nibbleshift + 4&7
						if lang.PostDec(&first) == 0 {
							ipos++
							first = 1
						}
					}
					zend.AddAssocStr(return_value, n, string(buf))
				case 'c':
					fallthrough
				case 'C':
					var issigned = lang.CondF1(type_ == 'c', func() int { return input[inputpos] & 0x80 }, 0)
					var v = PhpUnpack(&input[inputpos], 1, issigned, ByteMap)
					zend.AddAssocLong(return_value, n, v)
				case 's':
					fallthrough
				case 'S':
					fallthrough
				case 'n':
					fallthrough
				case 'v':
					var v zend.ZendLong
					var issigned = 0
					var map_ *int = MachineEndianShortMap
					if type_ == 's' {
						issigned = input[inputpos+lang.Cond(MachineLittleEndian, 1, 0)] & 0x80
					} else if type_ == 'n' {
						map_ = BigEndianShortMap
					} else if type_ == 'v' {
						map_ = LittleEndianShortMap
					}
					v = PhpUnpack(&input[inputpos], 2, issigned, map_)
					zend.AddAssocLong(return_value, n, v)
				case 'i':
					fallthrough
				case 'I':
					var v zend.ZendLong
					var issigned = 0
					if type_ == 'i' {
						issigned = input[inputpos+lang.CondF1(MachineLittleEndian, func() int { return b.SizeOf("int") - 1 }, 0)] & 0x80
					}
					v = PhpUnpack(&input[inputpos], b.SizeOf("int"), issigned, IntMap)
					zend.AddAssocLong(return_value, n, v)
				case 'l':
					fallthrough
				case 'L':
					fallthrough
				case 'N':
					fallthrough
				case 'V':
					var issigned = 0
					var map_ *int = MachineEndianLongMap
					var v = 0
					if type_ == 'l' || type_ == 'L' {
						issigned = input[inputpos+lang.Cond(MachineLittleEndian, 3, 0)] & 0x80
					} else if type_ == 'N' {
						issigned = input[inputpos] & 0x80
						map_ = BigEndianLongMap
					} else if type_ == 'V' {
						issigned = input[inputpos+3] & 0x80
						map_ = LittleEndianLongMap
					}
					if zend.SIZEOF_ZEND_LONG > 4 && issigned != 0 {
						v = ^core.INT_MAX
					}
					v |= PhpUnpack(&input[inputpos], 4, issigned, map_)
					if zend.SIZEOF_ZEND_LONG > 4 {
						if type_ == 'l' {
							v = signed__int(v)
						} else {
							v = uint(v)
						}
					}
					zend.AddAssocLong(return_value, n, v)
				case 'q':
					fallthrough
				case 'Q':
					fallthrough
				case 'J':
					fallthrough
				case 'P':
					var issigned = 0
					var map_ *int = MachineEndianLonglongMap
					var v = 0
					if type_ == 'q' || type_ == 'Q' {
						issigned = input[inputpos+lang.Cond(MachineLittleEndian, 7, 0)] & 0x80
					} else if type_ == 'J' {
						issigned = input[inputpos] & 0x80
						map_ = BigEndianLonglongMap
					} else if type_ == 'P' {
						issigned = input[inputpos+7] & 0x80
						map_ = LittleEndianLonglongMap
					}
					v = PhpUnpack(&input[inputpos], 8, issigned, map_)
					if type_ == 'q' {
						v = zend.ZendLong(v)
					} else {
						v = zend.ZendUlong(v)
					}
					zend.AddAssocLong(return_value, n, v)
				case 'f':
					fallthrough
				case 'g':
					fallthrough
				case 'G':
					var v float
					if type_ == 'g' {
						v = PhpPackParseFloat(1, &input[inputpos])
					} else if type_ == 'G' {
						v = PhpPackParseFloat(0, &input[inputpos])
					} else {
						memcpy(&v, &input[inputpos], b.SizeOf("float"))
					}
					zend.AddAssocDouble(return_value, n, float64(v))
				case 'd':
					fallthrough
				case 'e':
					fallthrough
				case 'E':
					var v float64
					if type_ == 'e' {
						v = PhpPackParseDouble(1, &input[inputpos])
					} else if type_ == 'E' {
						v = PhpPackParseDouble(0, &input[inputpos])
					} else {
						memcpy(&v, &input[inputpos], b.SizeOf("double"))
					}
					zend.AddAssocDouble(return_value, n, v)
				case 'x':

					/* Do nothing with input, just skip it */

				case 'X':
					if inputpos < size {
						inputpos = -size
						i = arg - 1
						if arg >= 0 {
							core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: outside of string", type_))
						}
					}
				case '@':
					if arg <= inputlen {
						inputpos = arg
					} else {
						core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: outside of string", type_))
					}
					i = arg - 1
				}
				inputpos += size
				if inputpos < 0 {
					if size != -1 {
						core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: outside of string", type_))
					}
					inputpos = 0
				}
			} else if arg < 0 {

				/* Reached end of input for '*' repeater */

				break

				/* Reached end of input for '*' repeater */

			} else {
				core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type %c: not enough input, need %d, have %d", type_, size, inputlen-inputpos))
				return_value.Array().Destroy()
				return_value.SetFalse()
				return
			}
		}
		if formatlen > 0 {
			formatlen--
			format++
		}
	}
}
func ZmStartupPack(moduleNumber int) int {
	var machine_endian_check = 1
	var i int
	MachineLittleEndian = (*byte)(&machine_endian_check)[0]
	if MachineLittleEndian {

		/* Where to get lo to hi bytes from */

		ByteMap[0] = 0
		for i = 0; i < int(b.SizeOf("int")); i++ {
			IntMap[i] = i
		}
		MachineEndianShortMap[0] = 0
		MachineEndianShortMap[1] = 1
		BigEndianShortMap[0] = 1
		BigEndianShortMap[1] = 0
		LittleEndianShortMap[0] = 0
		LittleEndianShortMap[1] = 1
		MachineEndianLongMap[0] = 0
		MachineEndianLongMap[1] = 1
		MachineEndianLongMap[2] = 2
		MachineEndianLongMap[3] = 3
		BigEndianLongMap[0] = 3
		BigEndianLongMap[1] = 2
		BigEndianLongMap[2] = 1
		BigEndianLongMap[3] = 0
		LittleEndianLongMap[0] = 0
		LittleEndianLongMap[1] = 1
		LittleEndianLongMap[2] = 2
		LittleEndianLongMap[3] = 3
		MachineEndianLonglongMap[0] = 0
		MachineEndianLonglongMap[1] = 1
		MachineEndianLonglongMap[2] = 2
		MachineEndianLonglongMap[3] = 3
		MachineEndianLonglongMap[4] = 4
		MachineEndianLonglongMap[5] = 5
		MachineEndianLonglongMap[6] = 6
		MachineEndianLonglongMap[7] = 7
		BigEndianLonglongMap[0] = 7
		BigEndianLonglongMap[1] = 6
		BigEndianLonglongMap[2] = 5
		BigEndianLonglongMap[3] = 4
		BigEndianLonglongMap[4] = 3
		BigEndianLonglongMap[5] = 2
		BigEndianLonglongMap[6] = 1
		BigEndianLonglongMap[7] = 0
		LittleEndianLonglongMap[0] = 0
		LittleEndianLonglongMap[1] = 1
		LittleEndianLonglongMap[2] = 2
		LittleEndianLonglongMap[3] = 3
		LittleEndianLonglongMap[4] = 4
		LittleEndianLonglongMap[5] = 5
		LittleEndianLonglongMap[6] = 6
		LittleEndianLonglongMap[7] = 7
	} else {
		var size = b.SizeOf("Z_LVAL ( val )")

		/* Where to get hi to lo bytes from */

		ByteMap[0] = size - 1
		for i = 0; i < int(b.SizeOf("int")); i++ {
			IntMap[i] = size - (b.SizeOf("int") - i)
		}
		MachineEndianShortMap[0] = size - 2
		MachineEndianShortMap[1] = size - 1
		BigEndianShortMap[0] = size - 2
		BigEndianShortMap[1] = size - 1
		LittleEndianShortMap[0] = size - 1
		LittleEndianShortMap[1] = size - 2
		MachineEndianLongMap[0] = size - 4
		MachineEndianLongMap[1] = size - 3
		MachineEndianLongMap[2] = size - 2
		MachineEndianLongMap[3] = size - 1
		BigEndianLongMap[0] = size - 4
		BigEndianLongMap[1] = size - 3
		BigEndianLongMap[2] = size - 2
		BigEndianLongMap[3] = size - 1
		LittleEndianLongMap[0] = size - 1
		LittleEndianLongMap[1] = size - 2
		LittleEndianLongMap[2] = size - 3
		LittleEndianLongMap[3] = size - 4
		MachineEndianLonglongMap[0] = size - 8
		MachineEndianLonglongMap[1] = size - 7
		MachineEndianLonglongMap[2] = size - 6
		MachineEndianLonglongMap[3] = size - 5
		MachineEndianLonglongMap[4] = size - 4
		MachineEndianLonglongMap[5] = size - 3
		MachineEndianLonglongMap[6] = size - 2
		MachineEndianLonglongMap[7] = size - 1
		BigEndianLonglongMap[0] = size - 8
		BigEndianLonglongMap[1] = size - 7
		BigEndianLonglongMap[2] = size - 6
		BigEndianLonglongMap[3] = size - 5
		BigEndianLonglongMap[4] = size - 4
		BigEndianLonglongMap[5] = size - 3
		BigEndianLonglongMap[6] = size - 2
		BigEndianLonglongMap[7] = size - 1
		LittleEndianLonglongMap[0] = size - 1
		LittleEndianLonglongMap[1] = size - 2
		LittleEndianLonglongMap[2] = size - 3
		LittleEndianLonglongMap[3] = size - 4
		LittleEndianLonglongMap[4] = size - 5
		LittleEndianLonglongMap[5] = size - 6
		LittleEndianLonglongMap[6] = size - 7
		LittleEndianLonglongMap[7] = size - 8
	}
	return types.SUCCESS
}
