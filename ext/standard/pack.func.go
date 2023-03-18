// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func PhpPack(val *types.Zval, size int, map_ *int, output *byte) {
	var i int
	var v *byte
	if val.GetType() != types.IS_LONG {
		zend.ConvertToLong(val)
	}
	v = (*byte)(&(val.GetLval()))
	for i = 0; i < size; i++ {
		b.PostInc(&(*output)) = v[map_[i]]
	}
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
func PhpPackCopyFloat(is_little_endian int, dst any, f float) {
	var m struct /* union */ {
		f float
		i uint32
	}
	m.f = f
	if is_little_endian == 0 {
		m.i = PhpPackReverseInt32(m.i)
	}
	memcpy(dst, m.f, b.SizeOf("float"))
}
func PhpPackCopyDouble(is_little_endian int, dst any, d float64) {
	var m struct /* union */ {
		d float64
		i uint64
	}
	m.d = d
	if is_little_endian == 0 {
		m.i = PhpPackReverseInt64(m.i)
	}
	memcpy(dst, m.d, b.SizeOf("double"))
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
func ZifPack(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var argv *types.Zval = nil
	var num_args int = 0
	var i int
	var currentarg int
	var format *byte
	var formatlen int
	var formatcodes *byte
	var formatargs *int
	var formatcount int = 0
	var outputpos int = 0
	var outputsize int = 0
	var output *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgString(_arg, &format, &formatlen, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				argv = _real_arg + 1
				num_args = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				argv = nil
				num_args = 0
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* We have a maximum of <formatlen> format codes to deal with */

	formatcodes = zend.SafeEmalloc(formatlen, b.SizeOf("* formatcodes"), 0)
	formatargs = zend.SafeEmalloc(formatlen, b.SizeOf("* formatargs"), 0)
	currentarg = 0

	/* Preprocess format into formatcodes and formatargs */

	for i = 0; i < formatlen; formatcount++ {
		var code byte = format[b.PostInc(&i)]
		var arg int = 1

		/* Handle format arguments if any */

		if i < formatlen {
			var c byte = format[i]
			if c == '*' {
				arg = -1
				i++
			} else if c >= '0' && c <= '9' {
				arg = atoi(&format[i])
				for format[i] >= '0' && format[i] <= '9' && i < formatlen {
					i++
				}
			}
		}

		/* Handle special arg '*' for all codes and check argv overflows */

		switch int(code) {
		case 'x':
			fallthrough
		case 'X':
			fallthrough
		case '@':
			if arg < 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: '*' ignored", code)
				arg = 1
			}
		case 'a':
			fallthrough
		case 'A':
			fallthrough
		case 'Z':
			fallthrough
		case 'h':
			fallthrough
		case 'H':
			if currentarg >= num_args {
				zend.Efree(formatcodes)
				zend.Efree(formatargs)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: not enough arguments", code)
				return_value.SetFalse()
				return
			}
			if arg < 0 {
				if zend.TryConvertToString(&argv[currentarg]) == 0 {
					zend.Efree(formatcodes)
					zend.Efree(formatargs)
					return
				}
				arg = argv[currentarg].GetStr().GetLen()
				if code == 'Z' {

					/* add one because Z is always NUL-terminated:
					 * pack("Z*", "aa") === "aa\0"
					 * pack("Z2", "aa") === "a\0" */

					arg++

					/* add one because Z is always NUL-terminated:
					 * pack("Z*", "aa") === "aa\0"
					 * pack("Z2", "aa") === "a\0" */

				}
			}
			currentarg++
		case 'q':
			fallthrough
		case 'Q':
			fallthrough
		case 'J':
			fallthrough
		case 'P':
			fallthrough
		case 'c':
			fallthrough
		case 'C':
			fallthrough
		case 's':
			fallthrough
		case 'S':
			fallthrough
		case 'i':
			fallthrough
		case 'I':
			fallthrough
		case 'l':
			fallthrough
		case 'L':
			fallthrough
		case 'n':
			fallthrough
		case 'N':
			fallthrough
		case 'v':
			fallthrough
		case 'V':
			fallthrough
		case 'f':
			fallthrough
		case 'g':
			fallthrough
		case 'G':
			fallthrough
		case 'd':
			fallthrough
		case 'e':
			fallthrough
		case 'E':
			if arg < 0 {
				arg = num_args - currentarg
			}
			if currentarg > core.INT_MAX-arg {
				goto too_few_args
			}
			currentarg += arg
			if currentarg > num_args {
			too_few_args:
				zend.Efree(formatcodes)
				zend.Efree(formatargs)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: too few arguments", code)
				return_value.SetFalse()
				return
			}
		default:
			zend.Efree(formatcodes)
			zend.Efree(formatargs)
			core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: unknown format code", code)
			return_value.SetFalse()
			return
		}
		formatcodes[formatcount] = code
		formatargs[formatcount] = arg
	}
	if currentarg < num_args {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%d arguments unused", num_args-currentarg)
	}

	/* Calculate output length and upper bound while processing*/

	for i = 0; i < formatcount; i++ {
		var code int = int(formatcodes[i])
		var arg int = formatargs[i]
		switch int(code) {
		case 'h':
			fallthrough
		case 'H':
			if (arg+arg%2)/2 < 0 || (core.INT_MAX-outputpos)/int(1) < (arg+arg%2)/2 {
				zend.Efree(formatcodes)
				zend.Efree(formatargs)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: integer overflow in format string", code)
				return_value.SetFalse()
				return
			}
			outputpos += (arg + arg%2) / 2 * 1
		case 'a':
			fallthrough
		case 'A':
			fallthrough
		case 'Z':
			fallthrough
		case 'c':
			fallthrough
		case 'C':
			fallthrough
		case 'x':
			if arg < 0 || (core.INT_MAX-outputpos)/int(1) < arg {
				zend.Efree(formatcodes)
				zend.Efree(formatargs)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: integer overflow in format string", code)
				return_value.SetFalse()
				return
			}
			outputpos += arg * 1
		case 's':
			fallthrough
		case 'S':
			fallthrough
		case 'n':
			fallthrough
		case 'v':
			if arg < 0 || (core.INT_MAX-outputpos)/int(2) < arg {
				zend.Efree(formatcodes)
				zend.Efree(formatargs)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: integer overflow in format string", code)
				return_value.SetFalse()
				return
			}
			outputpos += arg * 2
		case 'i':
			fallthrough
		case 'I':
			if arg < 0 || (core.INT_MAX-outputpos)/int(b.SizeOf("int")) < arg {
				zend.Efree(formatcodes)
				zend.Efree(formatargs)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: integer overflow in format string", code)
				return_value.SetFalse()
				return
			}
			outputpos += arg * b.SizeOf("int")
		case 'l':
			fallthrough
		case 'L':
			fallthrough
		case 'N':
			fallthrough
		case 'V':
			if arg < 0 || (core.INT_MAX-outputpos)/int(4) < arg {
				zend.Efree(formatcodes)
				zend.Efree(formatargs)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: integer overflow in format string", code)
				return_value.SetFalse()
				return
			}
			outputpos += arg * 4
		case 'q':
			fallthrough
		case 'Q':
			fallthrough
		case 'J':
			fallthrough
		case 'P':
			if arg < 0 || (core.INT_MAX-outputpos)/int(8) < arg {
				zend.Efree(formatcodes)
				zend.Efree(formatargs)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: integer overflow in format string", code)
				return_value.SetFalse()
				return
			}
			outputpos += arg * 8
		case 'f':
			fallthrough
		case 'g':
			fallthrough
		case 'G':
			if arg < 0 || (core.INT_MAX-outputpos)/int(b.SizeOf("float")) < arg {
				zend.Efree(formatcodes)
				zend.Efree(formatargs)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: integer overflow in format string", code)
				return_value.SetFalse()
				return
			}
			outputpos += arg * b.SizeOf("float")
		case 'd':
			fallthrough
		case 'e':
			fallthrough
		case 'E':
			if arg < 0 || (core.INT_MAX-outputpos)/int(b.SizeOf("double")) < arg {
				zend.Efree(formatcodes)
				zend.Efree(formatargs)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: integer overflow in format string", code)
				return_value.SetFalse()
				return
			}
			outputpos += arg * b.SizeOf("double")
		case 'X':
			outputpos -= arg
			if outputpos < 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: outside of string", code)
				outputpos = 0
			}
		case '@':
			outputpos = arg
		}
		if outputsize < outputpos {
			outputsize = outputpos
		}
	}
	output = types.ZendStringAlloc(outputsize, 0)
	outputpos = 0
	currentarg = 0

	/* Do actual packing */

	for i = 0; i < formatcount; i++ {
		var code int = int(formatcodes[i])
		var arg int = formatargs[i]
		switch int(code) {
		case 'a':
			fallthrough
		case 'A':
			fallthrough
		case 'Z':
			var arg_cp int = b.CondF2(code != 'Z', arg, func() int { return b.Max(0, arg-1) })
			var tmp_str *types.ZendString
			var str *types.ZendString = zend.ZvalGetTmpString(&argv[b.PostInc(&currentarg)], &tmp_str)
			memset(&output.GetVal()[outputpos], b.Cond(code == 'a' || code == 'Z', '0', ' '), arg)
			memcpy(&output.GetVal()[outputpos], str.GetVal(), b.CondF1(str.GetLen() < arg_cp, func() int { return str.GetLen() }, arg_cp))
			outputpos += arg
			zend.ZendTmpStringRelease(tmp_str)
		case 'h':
			fallthrough
		case 'H':
			var nibbleshift int = b.Cond(code == 'h', 0, 4)
			var first int = 1
			var tmp_str *types.ZendString
			var str *types.ZendString = zend.ZvalGetTmpString(&argv[b.PostInc(&currentarg)], &tmp_str)
			var v *byte = str.GetVal()
			outputpos--
			if int(arg > str.GetLen()) != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: not enough characters in string", code)
				arg = str.GetLen()
			}
			for b.PostDec(&arg) > 0 {
				var n byte = b.PostInc(&(*v))
				if n >= '0' && n <= '9' {
					n -= '0'
				} else if n >= 'A' && n <= 'F' {
					n -= 'A' - 10
				} else if n >= 'a' && n <= 'f' {
					n -= 'a' - 10
				} else {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: illegal hex digit %c", code, n)
					n = 0
				}
				if b.PostDec(&first) {
					output.GetVal()[b.PreInc(&outputpos)] = 0
				} else {
					first = 1
				}
				output.GetVal()[outputpos] |= n << nibbleshift
				nibbleshift = nibbleshift + 4&7
			}
			outputpos++
			zend.ZendTmpStringRelease(tmp_str)
		case 'c':
			fallthrough
		case 'C':
			for b.PostDec(&arg) > 0 {
				PhpPack(&argv[b.PostInc(&currentarg)], 1, ByteMap, &output.GetVal()[outputpos])
				outputpos++
			}
		case 's':
			fallthrough
		case 'S':
			fallthrough
		case 'n':
			fallthrough
		case 'v':
			var map_ *int = MachineEndianShortMap
			if code == 'n' {
				map_ = BigEndianShortMap
			} else if code == 'v' {
				map_ = LittleEndianShortMap
			}
			for b.PostDec(&arg) > 0 {
				PhpPack(&argv[b.PostInc(&currentarg)], 2, map_, &output.GetVal()[outputpos])
				outputpos += 2
			}
		case 'i':
			fallthrough
		case 'I':
			for b.PostDec(&arg) > 0 {
				PhpPack(&argv[b.PostInc(&currentarg)], b.SizeOf("int"), IntMap, &output.GetVal()[outputpos])
				outputpos += b.SizeOf("int")
			}
		case 'l':
			fallthrough
		case 'L':
			fallthrough
		case 'N':
			fallthrough
		case 'V':
			var map_ *int = MachineEndianLongMap
			if code == 'N' {
				map_ = BigEndianLongMap
			} else if code == 'V' {
				map_ = LittleEndianLongMap
			}
			for b.PostDec(&arg) > 0 {
				PhpPack(&argv[b.PostInc(&currentarg)], 4, map_, &output.GetVal()[outputpos])
				outputpos += 4
			}
		case 'q':
			fallthrough
		case 'Q':
			fallthrough
		case 'J':
			fallthrough
		case 'P':
			var map_ *int = MachineEndianLonglongMap
			if code == 'J' {
				map_ = BigEndianLonglongMap
			} else if code == 'P' {
				map_ = LittleEndianLonglongMap
			}
			for b.PostDec(&arg) > 0 {
				PhpPack(&argv[b.PostInc(&currentarg)], 8, map_, &output.GetVal()[outputpos])
				outputpos += 8
			}
		case 'f':
			for b.PostDec(&arg) > 0 {
				var v float = float(zend.ZvalGetDouble(&argv[b.PostInc(&currentarg)]))
				memcpy(&output.GetVal()[outputpos], &v, b.SizeOf("v"))
				outputpos += b.SizeOf("v")
			}
		case 'g':

			/* pack little endian float */

			for b.PostDec(&arg) > 0 {
				var v float = float(zend.ZvalGetDouble(&argv[b.PostInc(&currentarg)]))
				PhpPackCopyFloat(1, &output.GetVal()[outputpos], v)
				outputpos += b.SizeOf("v")
			}
		case 'G':

			/* pack big endian float */

			for b.PostDec(&arg) > 0 {
				var v float = float(zend.ZvalGetDouble(&argv[b.PostInc(&currentarg)]))
				PhpPackCopyFloat(0, &output.GetVal()[outputpos], v)
				outputpos += b.SizeOf("v")
			}
		case 'd':
			for b.PostDec(&arg) > 0 {
				var v float64 = float64(zend.ZvalGetDouble(&argv[b.PostInc(&currentarg)]))
				memcpy(&output.GetVal()[outputpos], &v, b.SizeOf("v"))
				outputpos += b.SizeOf("v")
			}
		case 'e':

			/* pack little endian double */

			for b.PostDec(&arg) > 0 {
				var v float64 = float64(zend.ZvalGetDouble(&argv[b.PostInc(&currentarg)]))
				PhpPackCopyDouble(1, &output.GetVal()[outputpos], v)
				outputpos += b.SizeOf("v")
			}
		case 'E':

			/* pack big endian double */

			for b.PostDec(&arg) > 0 {
				var v float64 = float64(zend.ZvalGetDouble(&argv[b.PostInc(&currentarg)]))
				PhpPackCopyDouble(0, &output.GetVal()[outputpos], v)
				outputpos += b.SizeOf("v")
			}
		case 'x':
			memset(&output.GetVal()[outputpos], '0', arg)
			outputpos += arg
		case 'X':
			outputpos -= arg
			if outputpos < 0 {
				outputpos = 0
			}
		case '@':
			if arg > outputpos {
				memset(&output.GetVal()[outputpos], '0', arg-outputpos)
			}
			outputpos = arg
		}
	}
	zend.Efree(formatcodes)
	zend.Efree(formatargs)
	output.GetVal()[outputpos] = '0'
	output.SetLen(outputpos)
	return_value.SetString(output)
	return
}
func PhpUnpack(data *byte, size int, issigned int, map_ *int) zend.ZendLong {
	var result zend.ZendLong
	var cresult *byte = (*byte)(&result)
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
func ZifUnpack(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var format *byte
	var input *byte
	var formatarg *types.ZendString
	var inputarg *types.ZendString
	var formatlen zend.ZendLong
	var inputpos zend.ZendLong
	var inputlen zend.ZendLong
	var i int
	var offset zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &formatarg, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &inputarg, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if !argparse.ZendParseArgLong00(_arg, &offset) {
				_expected_type = argparse.Z_EXPECTED_LONG
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
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
		core.PhpErrorDocref(nil, zend.E_WARNING, "Offset "+zend.ZEND_LONG_FMT+" is out of input range", offset)
		return_value.SetFalse()
		return
	}
	input += offset
	inputlen -= offset
	zend.ArrayInit(return_value)
	for b.PostDec(&formatlen) > 0 {
		var type_ byte = *(b.PostInc(&format))
		var c byte
		var arg int = 1
		var argb int
		var name *byte
		var namelen int
		var size int = 0

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
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: '*' ignored", type_)
				arg = 1
			}
		case '@':
			size = 0
		case 'a':
			fallthrough
		case 'A':
			fallthrough
		case 'Z':
			size = arg
			arg = 1
		case 'h':
			fallthrough
		case 'H':
			if arg > 0 {
				size = (arg + arg%2) / 2
			} else {
				size = arg
			}
			arg = 1
		case 'c':
			fallthrough
		case 'C':
			fallthrough
		case 'x':
			size = 1
		case 's':
			fallthrough
		case 'S':
			fallthrough
		case 'n':
			fallthrough
		case 'v':
			size = 2
		case 'i':
			fallthrough
		case 'I':
			size = b.SizeOf("int")
		case 'l':
			fallthrough
		case 'L':
			fallthrough
		case 'N':
			fallthrough
		case 'V':
			size = 4
		case 'q':
			fallthrough
		case 'Q':
			fallthrough
		case 'J':
			fallthrough
		case 'P':
			size = 8
		case 'f':
			fallthrough
		case 'g':
			fallthrough
		case 'G':
			size = b.SizeOf("float")
		case 'd':
			fallthrough
		case 'e':
			fallthrough
		case 'E':
			size = b.SizeOf("double")
		default:
			core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid format type %c", type_)
			return_value.GetArr().DestroyEx()
			return_value.SetFalse()
			return
		}
		if size != 0 && size != -1 && size < 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: integer overflow", type_)
			return_value.GetArr().DestroyEx()
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
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: integer overflow", type_)
				return_value.GetArr().DestroyEx()
				return_value.SetFalse()
				return
			}
			if inputpos+size <= inputlen {
				switch int(type_) {
				case 'a':

					/* a will not strip any trailing whitespace or null padding */

					var len_ zend.ZendLong = inputlen - inputpos

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
					var len_ zend.ZendLong = inputlen - inputpos

					/* If size was given take minimum of len and size */

					if size >= 0 && len_ > size {
						len_ = size
					}
					size = len_

					/* Remove trailing white space and nulls chars from unpacked data */

					for b.PreDec(&len_) >= 0 {
						if input[inputpos+len_] != padn && input[inputpos+len_] != pads && input[inputpos+len_] != padt && input[inputpos+len_] != padc && input[inputpos+len_] != padl {
							break
						}
					}
					zend.AddAssocStringl(return_value, n, &input[inputpos], len_+1)
				case 'Z':

					/* Z will strip everything after the first null character */

					var pad byte = '0'
					var s zend.ZendLong
					var len_ zend.ZendLong = inputlen - inputpos

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
					var len_ zend.ZendLong = (inputlen - inputpos) * 2
					var nibbleshift int = b.Cond(type_ == 'h', 0, 4)
					var first int = 1
					var buf *types.ZendString
					var ipos zend.ZendLong
					var opos zend.ZendLong

					/* If size was given take minimum of len and size */

					if size >= 0 && len_ > size*2 {
						len_ = size * 2
					}
					if len_ > 0 && argb > 0 {
						len_ -= argb % 2
					}
					buf = types.ZendStringAlloc(len_, 0)
					opos = 0
					ipos = opos
					for ; opos < len_; opos++ {
						var cc byte = input[inputpos+ipos] >> nibbleshift & 0xf
						if cc < 10 {
							cc += '0'
						} else {
							cc += 'a' - 10
						}
						buf.GetVal()[opos] = cc
						nibbleshift = nibbleshift + 4&7
						if b.PostDec(&first) == 0 {
							ipos++
							first = 1
						}
					}
					buf.GetVal()[len_] = '0'
					zend.AddAssocStr(return_value, n, buf.GetStr())
				case 'c':
					fallthrough
				case 'C':
					var issigned int = b.CondF1(type_ == 'c', func() int { return input[inputpos] & 0x80 }, 0)
					var v zend.ZendLong = PhpUnpack(&input[inputpos], 1, issigned, ByteMap)
					zend.AddAssocLong(return_value, n, v)
				case 's':
					fallthrough
				case 'S':
					fallthrough
				case 'n':
					fallthrough
				case 'v':
					var v zend.ZendLong
					var issigned int = 0
					var map_ *int = MachineEndianShortMap
					if type_ == 's' {
						issigned = input[inputpos+b.Cond(MachineLittleEndian, 1, 0)] & 0x80
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
					var issigned int = 0
					if type_ == 'i' {
						issigned = input[inputpos+b.CondF1(MachineLittleEndian, func() int { return b.SizeOf("int") - 1 }, 0)] & 0x80
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
					var issigned int = 0
					var map_ *int = MachineEndianLongMap
					var v zend.ZendLong = 0
					if type_ == 'l' || type_ == 'L' {
						issigned = input[inputpos+b.Cond(MachineLittleEndian, 3, 0)] & 0x80
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
					var issigned int = 0
					var map_ *int = MachineEndianLonglongMap
					var v zend.ZendLong = 0
					if type_ == 'q' || type_ == 'Q' {
						issigned = input[inputpos+b.Cond(MachineLittleEndian, 7, 0)] & 0x80
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
							core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: outside of string", type_)
						}
					}
				case '@':
					if arg <= inputlen {
						inputpos = arg
					} else {
						core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: outside of string", type_)
					}
					i = arg - 1
				}
				inputpos += size
				if inputpos < 0 {
					if size != -1 {
						core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: outside of string", type_)
					}
					inputpos = 0
				}
			} else if arg < 0 {

				/* Reached end of input for '*' repeater */

				break

				/* Reached end of input for '*' repeater */

			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Type %c: not enough input, need %d, have "+zend.ZEND_LONG_FMT, type_, size, inputlen-inputpos)
				return_value.GetArr().DestroyEx()
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
func ZmStartupPack(type_ int, module_number int) int {
	var machine_endian_check int = 1
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
		var val types.Zval
		var size int = b.SizeOf("Z_LVAL ( val )")
		val.SetLval(0)

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
