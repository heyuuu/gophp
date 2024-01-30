package standard

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/kits/strkit"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

/* Whether machine is little endian */
var machineEndian binary.ByteOrder
var bigEndian binary.ByteOrder = binary.BigEndian
var littleEndian binary.ByteOrder = binary.LittleEndian

const sizeofInt = 4
const sizeofFloat = 4
const sizeofDouble = 8

func init() {
	if isMachineLittleEndian() {
		machineEndian = littleEndian
	} else {
		machineEndian = bigEndian
	}
}

func isMachineLittleEndian() bool {
	n := 0x12345678
	return *(*byte)(unsafe.Pointer(&n)) == 0x78
}

func packReadToken(format string, unpack bool) (nextFormat string, typ byte, arg int, name string) {
	if format == "" {
		return "", 0, 0, ""
	}

	// type
	typ = format[0]
	format = format[1:]

	// arg
	arg = 1
	if len(format) > 0 {
		/* Handle format arguments if any */
		if format[0] == '*' {
			arg = -1
			format = format[1:]
		} else if ascii.IsDigit(format[0]) {
			numLen := 1
			for numLen < len(format) && ascii.IsDigit(format[numLen]) {
				numLen++
			}
			arg, _ = strconv.Atoi(format[:numLen])
			if arg > math.MaxInt32 {
				arg = math.MinInt32 // 用于触发长度错误(integer overflow)
			}
			format = format[numLen:]
		}
	}

	// name
	if unpack {
		if idx := strings.IndexByte(format, '/'); idx >= 0 {
			name = format[:idx]
			format = format[idx+1:]
		} else {
			name = format
			format = ""
		}
	}

	return format, typ, arg, name
}

func packCheckSize(ctx *php.Context, codes []byte, args []int) (int, bool) {
	var outputPos = 0
	var outputSize = 0
	addOutputPos := func(code byte, arg int, inSize int, outSize int) bool {
		inChunk := (arg + arg%inSize) / inSize
		if inChunk < 0 || (types.MaxStrLen-outputPos)/outSize < inChunk {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: integer overflow in format string", code))
			return false
		}
		outputPos += inChunk * outSize
		return true
	}

	/* Calculate output length and upper bound while processing*/
	for i := 0; i < len(codes); i++ {
		var code = codes[i]
		var arg = args[i]
		var ok bool
		switch code {
		case 'h', 'H':
			ok = addOutputPos(code, arg, 2, 1)
		case 'a', 'A', 'Z', 'c', 'C', 'x':
			ok = addOutputPos(code, arg, 1, 1)
		case 's', 'S', 'n', 'v':
			ok = addOutputPos(code, arg, 1, 2)
		case 'i', 'I':
			ok = addOutputPos(code, arg, 1, sizeofInt)
		case 'l', 'L', 'N', 'V':
			ok = addOutputPos(code, arg, 1, 4)
		case 'q', 'Q', 'J', 'P':
			ok = addOutputPos(code, arg, 1, 8)
		case 'f', 'g', 'G':
			ok = addOutputPos(code, arg, 1, sizeofFloat)
		case 'd', 'e', 'E':
			ok = addOutputPos(code, arg, 1, sizeofDouble)
		case 'X':
			outputPos -= arg
			if outputPos < 0 {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: outside of string", code))
				outputPos = 0
			}
		case '@':
			outputPos = arg
		}
		if !ok {
			return 0, false
		}
		if outputSize < outputPos {
			outputSize = outputPos
		}
	}
	return outputSize, true
}

func ZifPack(ctx *php.Context, format_ string, _ zpp.Opt, args []types.Zval) (string, bool) {
	var numArgs = len(args)
	var currentarg int
	var formatcodes []byte
	var formatargs []int
	var outputSize = 0

	/* We have a maximum of <formatlen> format codes to deal with */
	formatcodes = make([]byte, 0, len(format_))
	formatargs = make([]int, 0, len(format_))
	currentarg = 0

	/* Preprocess format into formatcodes and formatargs */
	parsingFormat := format_
	for len(parsingFormat) > 0 {
		var code byte
		var arg int
		parsingFormat, code, arg, _ = packReadToken(parsingFormat, false)

		/* Handle special arg '*' for all codes and check argv overflows */
		switch int(code) {
		case 'x', 'X', '@':
			if arg < 0 {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: '*' ignored", code))
				arg = 1
			}
		case 'a', 'A', 'Z', 'h', 'H':
			if currentarg >= numArgs {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: not enough arguments", code))
				return "", false
			}
			if arg < 0 {
				if str, ok := php.ZvalTryGetStr(ctx, args[currentarg]); ok {
					args[currentarg].SetString(str)
				} else {
					return "", false
				}
				arg = len(args[currentarg].String())
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
				arg = numArgs - currentarg
			}
			if currentarg > types.MaxLong-arg {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: too few arguments", code))
				return "", false
			}
			currentarg += arg
			if currentarg > numArgs {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: too few arguments", code))
				return "", false
			}
		default:
			php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: unknown format code", code))
			return "", false
		}

		formatcodes = append(formatcodes, code)
		formatargs = append(formatargs, arg)
	}
	if currentarg < numArgs {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("%d arguments unused", numArgs-currentarg))
	}

	/* Check size */
	outputSize, ok := packCheckSize(ctx, formatcodes, formatargs)
	if !ok {
		return "", false
	}

	/* Do actual packing */
	var buf strings.Builder
	buf.Grow(outputSize)

	argIdx := 0
	nextArg := func() types.Zval {
		if argIdx >= len(args) {
			return types.Undef
		}

		result := args[argIdx]
		argIdx++
		return result
	}

	for i := 0; i < len(formatcodes); i++ {
		var code = int(formatcodes[i])
		var arg = formatargs[i]
		switch code {
		case 'a':
			str := php.ZvalGetStrVal(ctx, nextArg())
			str = strkit.PadRight(str, arg, '\x00')[:arg]
			buf.WriteString(str)
		case 'A':
			str := php.ZvalGetStrVal(ctx, nextArg())
			str = strkit.PadRight(str, arg, ' ')[:arg]
			buf.WriteString(str)
		case 'Z':
			str := php.ZvalGetStrVal(ctx, nextArg())
			if arg <= 0 {
				str = ""
			} else {
				str = strkit.PadRight(str, arg-1, '\x00')[:arg-1] + "\x00"
			}
			buf.WriteString(str)
		case 'h', 'H':
			var isLittleEndian = code == 'h'
			var str = php.ZvalGetStrVal(ctx, nextArg())
			if arg > len(str) {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: not enough characters in string", code))
				arg = len(str)
			}

			for j := 0; j < arg; j += 2 {
				var digit0, digit1 byte
				var ok bool
				digit0, ok = ascii.ParseXDigit(str[j])
				if !ok {
					php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: illegal hex digit %c", code, str[j]))
				}
				if j+1 < arg {
					digit1, ok = ascii.ParseXDigit(str[j+1])
					if !ok {
						php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: illegal hex digit %c", code, str[j+1]))
					}
				}
				if isLittleEndian {
					buf.WriteByte((digit0 << 4) | digit1)
				} else {
					buf.WriteByte((digit1 << 4) | digit0)
				}
			}
		case 'c', 'C':
			for j := 0; j < arg; j++ {
				num := php.ZvalGetLong(ctx, nextArg())
				buf.WriteByte(byte(num))
			}
		case 's', 'S', 'n', 'v':
			order := machineEndian
			if code == 'n' {
				order = bigEndian
			} else if code == 'v' {
				order = littleEndian
			}
			for j := 0; j < arg; j++ {
				num := php.ZvalGetLong(ctx, nextArg())
				_ = binary.Write(&buf, order, int16(num))
			}
		case 'i', 'I':
			for j := 0; j < arg; j++ {
				num := php.ZvalGetLong(ctx, nextArg())
				_ = binary.Write(&buf, machineEndian, int(num))
			}
		case 'l', 'L', 'N', 'V':
			order := machineEndian
			if code == 'N' {
				order = bigEndian
			} else if code == 'V' {
				order = littleEndian
			}
			for j := 0; j < arg; j++ {
				num := php.ZvalGetLong(ctx, nextArg())
				_ = binary.Write(&buf, order, int32(num))
			}
		case 'q', 'Q', 'J', 'P':
			order := machineEndian
			if code == 'J' {
				order = bigEndian
			} else if code == 'P' {
				order = littleEndian
			}
			for j := 0; j < arg; j++ {
				num := php.ZvalGetLong(ctx, nextArg())
				_ = binary.Write(&buf, order, int64(num))
			}
		case 'f', 'g', 'G':
			order := machineEndian
			if code == 'g' {
				order = littleEndian
			} else if code == 'G' {
				order = bigEndian
			}
			for j := 0; j < arg; j++ {
				num := php.ZvalGetDouble(ctx, nextArg())
				_ = binary.Write(&buf, order, float32(num))
			}
		case 'd', 'e', 'E':
			order := machineEndian
			if code == 'e' {
				order = littleEndian
			} else if code == 'E' {
				order = bigEndian
			}
			for j := 0; j < arg; j++ {
				num := php.ZvalGetDouble(ctx, nextArg())
				_ = binary.Write(&buf, order, float64(num))
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

func ZifUnpack(ctx *php.Context, format string, input string, _ zpp.Opt, offset int) (*types.Array, bool) {
	if offset < 0 || offset > len(input) {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Offset %d is out of input range", offset))
		return nil, false
	}
	input = input[offset:]

	inputlen := len(input)
	inputpos := 0

	result := types.NewArray()

	var typ byte
	var arg int
	var name string
	for len(format) > 0 {
		format, typ, arg, name = packReadToken(format, true)

		var argb = arg
		var size = 0
		if len(name) > 200 {
			name = name[:200]
		}
		switch typ {
		case 'X':
			size = -1
			if arg < 0 {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: '*' ignored", typ))
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
			size = sizeofInt
		case 'l', 'L', 'N', 'V':
			size = 4
		case 'q', 'Q', 'J', 'P':
			size = 8
		case 'f', 'g', 'G':
			size = sizeofFloat
		case 'd', 'e', 'E':
			size = sizeofDouble
		default:
			php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Invalid format type %c", typ))
			return nil, false
		}
		if size != 0 && size != -1 && size < 0 {
			php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: integer overflow", typ))
			return nil, false
		}

		inputBytes := []byte(input)

		/* Do actual unpacking */
		for i := 0; i != arg; i++ {
			/* Space for name + number, safe as namelen is ensured <= 200 */
			var n string
			if arg != 1 || name == "" {
				/* Need to add element number to name */
				n = fmt.Sprintf("%s%d", name, i+1)
			} else {
				/* Truncate name to next format code or end of string */
				n = name
			}
			if size != 0 && size != -1 && types.MaxLong-size+1 < inputpos {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: integer overflow", typ))
				return nil, false
			}
			if inputpos+size <= inputlen {
				switch int(typ) {
				case 'a':
					/* a will not strip any trailing whitespace or null padding */
					var len_ = inputlen - inputpos

					/* If size was given take minimum of len and size */
					if size >= 0 && len_ > size {
						len_ = size
					}
					size = len_
					result.AddAssocStr(n, input[inputpos:inputpos+len_])
				case 'A':
					/* A will strip any trailing whitespace */
					var len_ = inputlen - inputpos

					/* If size was given take minimum of len and size */
					if size >= 0 && len_ > size {
						len_ = size
					}
					size = len_

					/* Remove trailing white space and nulls chars from unpacked data */
					str := strings.TrimRight(input[inputpos:inputpos+len_], "\000 \t\r\n")
					result.AddAssocStr(n, str)
				case 'Z':
					/* Z will strip everything after the first null character */
					var len_ = inputlen - inputpos

					/* If size was given take minimum of len and size */
					if size >= 0 && len_ > size {
						len_ = size
					}
					size = len_

					/* Remove everything after the first null */
					str := input[inputpos : inputpos+len_]
					if idx := strings.IndexByte(str, 0); idx >= 0 {
						str = str[:idx]
					}
					len_ = len(str)
					result.AddAssocStr(n, str)
				case 'h', 'H':
					var isLittleEndian = typ == 'h'
					var len_ = (inputlen - inputpos) * 2

					/* If size was given take minimum of len and size */
					if size >= 0 && len_ > size*2 {
						len_ = size * 2
					}
					if len_ > 0 && argb > 0 {
						len_ -= argb % 2
					}

					buf := make([]byte, len_+1)
					hex.Encode(buf, inputBytes[inputpos:inputpos+len(buf)/2])
					if !isLittleEndian {
						for j := 0; j < len_; j += 2 {
							buf[j], buf[j+1] = buf[j+1], buf[j]
						}
					}
					result.AddAssocStr(n, string(buf[:len_]))
				case 'c':
					v := int8(input[inputpos])
					result.AddAssocLong(n, int(v))
				case 'C':
					v := input[inputpos]
					result.AddAssocLong(n, int(v))
				case 's':
					v := int16(machineEndian.Uint16(inputBytes[inputpos:]))
					result.AddAssocLong(n, int(v))
				case 'S', 'n', 'v':
					order := machineEndian
					if typ == 'n' {
						order = bigEndian
					} else if typ == 'v' {
						order = littleEndian
					}
					v := order.Uint16(inputBytes[inputpos:])
					result.AddAssocLong(n, int(v))
				case 'i':
					v := int(machineEndian.Uint64(inputBytes[inputpos:]))
					result.AddAssocLong(n, v)
				case 'I':
					v := uint(machineEndian.Uint64(inputBytes[inputpos:]))
					result.AddAssocLong(n, int(v))
				case 'l':
					v := int32(machineEndian.Uint32(inputBytes[inputpos:]))
					result.AddAssocLong(n, int(v))
				case 'L', 'N', 'V':
					order := machineEndian
					if typ == 'N' {
						order = bigEndian
					} else if typ == 'V' {
						order = littleEndian
					}
					v := order.Uint32(inputBytes[inputpos:])
					result.AddAssocLong(n, int(v))
				case 'q':
					v := int64(machineEndian.Uint64(inputBytes[inputpos:]))
					result.AddAssocLong(n, int(v))
				case 'Q', 'J', 'P':
					order := machineEndian
					if typ == 'J' {
						order = bigEndian
					} else if typ == 'P' {
						order = littleEndian
					}
					v := order.Uint64(inputBytes[inputpos:])
					result.AddAssocLong(n, int(v))
				case 'f', 'g', 'G':
					order := machineEndian
					if typ == 'g' {
						order = littleEndian
					} else if typ == 'G' {
						order = bigEndian
					}
					v := math.Float32frombits(order.Uint32(inputBytes[inputpos:]))
					result.AddAssocDouble(n, float64(v))
				case 'd', 'e', 'E':
					order := machineEndian
					if typ == 'e' {
						order = littleEndian
					} else if typ == 'E' {
						order = bigEndian
					}
					v := math.Float64frombits(order.Uint64(inputBytes[inputpos:]))
					result.AddAssocDouble(n, v)
				case 'x':
					/* Do nothing with input, just skip it */
				case 'X':
					if inputpos < size {
						inputpos = -size
						i = arg - 1
						if arg >= 0 {
							php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: outside of string", typ))
						}
					}
				case '@':
					if arg <= inputlen {
						inputpos = arg
					} else {
						php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: outside of string", typ))
					}
					i = arg - 1
				}
				inputpos += size
				if inputpos < 0 {
					if size != -1 {
						php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: outside of string", typ))
					}
					inputpos = 0
				}
			} else if arg < 0 {
				/* Reached end of input for '*' repeater */
				break
			} else {
				php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Type %c: not enough input, need %d, have %d", typ, size, inputlen-inputpos))
				return nil, false
			}
		}
	}
	return result, true
}
