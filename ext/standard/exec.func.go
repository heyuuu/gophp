// <<generate>>

package standard

import (
	"sik/builtin"
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	"sik/zend"
)

func ZmStartupExec(type_ int, module_number int) int {
	/* This is just an arbitrary value for the fallback case. */

	CmdMaxLen = 4096
	return zend.SUCCESS
}
func PhpExec(type_ int, cmd *byte, array *zend.Zval, return_value *zend.Zval) int {
	var fp *r.FILE
	var buf *byte
	var l int = 0
	var pclose_return int
	var b *byte
	var d *byte = nil
	var stream *core.PhpStream
	var buflen int
	var bufl int = 0
	fp = zend.VCWD_POPEN(cmd, "r")
	if fp == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to fork [%s]", cmd)
		goto err
	}
	stream = streams.PhpStreamFopenFromPipe(fp, "rb")
	buf = (*byte)(zend.Emalloc(core.EXEC_INPUT_BUF))
	buflen = core.EXEC_INPUT_BUF
	if type_ != 3 {
		b = buf
		for core.PhpStreamGetLine(stream, b, core.EXEC_INPUT_BUF, &bufl) != nil {

			/* no new line found, let's read some more */

			if b[bufl-1] != '\n' && core.PhpStreamEof(stream) == 0 {
				if buflen < bufl+(b-buf)+core.EXEC_INPUT_BUF {
					bufl += b - buf
					buflen = bufl + core.EXEC_INPUT_BUF
					buf = zend.Erealloc(buf, buflen)
					b = buf + bufl
				} else {
					b += bufl
				}
				continue
			} else if b != buf {
				bufl += b - buf
			}
			if type_ == 1 {
				core.PHPWRITE(buf, bufl)
				if core.PhpOutputGetLevel() < 1 {
					core.SapiFlush()
				}
			} else if type_ == 2 {

				/* strip trailing whitespaces */

				l = bufl
				for b.PostDec(&l) > 0 && isspace((*uint8)(buf)[l]) {

				}
				if l != bufl-1 {
					bufl = l + 1
					buf[bufl] = '0'
				}
				zend.AddNextIndexStringl(array, buf, bufl)
			}
			b = buf
		}
		if bufl != 0 {

			/* output remaining data in buffer */

			if type_ == 1 && buf != b {
				core.PHPWRITE(buf, bufl)
				if core.PhpOutputGetLevel() < 1 {
					core.SapiFlush()
				}
			}

			/* strip trailing whitespaces if we have not done so already */

			if type_ == 2 && buf != b || type_ != 2 {
				l = bufl
				for b.PostDec(&l) > 0 && isspace((*uint8)(buf)[l]) {

				}
				if l != bufl-1 {
					bufl = l + 1
					buf[bufl] = '0'
				}
				if type_ == 2 {
					zend.AddNextIndexStringl(array, buf, bufl)
				}
			}

			/* Return last line from the shell command */

			return_value.SetRawString(builtin.CastStr(buf, bufl))

		} else {
			zend.ZVAL_EMPTY_STRING(return_value)
		}
	} else {
		var read ssize_t
		for b.Assign(&read, core.PhpStreamRead(stream, buf, core.EXEC_INPUT_BUF)) > 0 {
			core.PHPWRITE(buf, read)
		}
	}
	pclose_return = core.PhpStreamClose(stream)
	zend.Efree(buf)
done:
	if d != nil {
		zend.Efree(d)
	}
	return pclose_return
err:
	pclose_return = -1
	goto done
}
func PhpExecEx(executeData *zend.ZendExecuteData, return_value *zend.Zval, mode int) {
	var cmd *byte
	var cmd_len int
	var ret_code *zend.Zval = nil
	var ret_array *zend.Zval = nil
	var ret int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = b.Cond(mode != 0, 2, 3)
		var _num_args int = executeData.NumArgs()
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &cmd, &cmd_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			if mode == 0 {
				zend.Z_PARAM_PROLOGUE(0, 0)
				zend.ZendParseArgZvalDeref(_arg, &ret_array, 0)
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &ret_code, 0)
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
			return_value.SetFalse()
			return
		}
		break
	}
	if cmd_len == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Cannot execute a blank command")
		return_value.SetFalse()
		return
	}
	if strlen(cmd) != cmd_len {
		core.PhpErrorDocref(nil, zend.E_WARNING, "NULL byte detected. Possible attack")
		return_value.SetFalse()
		return
	}
	if ret_array == nil {
		ret = PhpExec(mode, cmd, nil, return_value)
	} else {
		if zend.Z_REFVAL_P(ret_array).IsType(zend.IS_ARRAY) {
			ret_array = zend.ZVAL_DEREF(ret_array)
			zend.SEPARATE_ARRAY(ret_array)
		} else {
			ret_array = zend.ZendTryArrayInit(ret_array)
			if ret_array == nil {
				return
			}
		}
		ret = PhpExec(2, cmd, ret_array, return_value)
	}
	if ret_code != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(ret_code, ret)
	}
}
func ZifExec(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpExecEx(executeData, return_value, 0)
}
func ZifSystem(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpExecEx(executeData, return_value, 1)
}
func ZifPassthru(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpExecEx(executeData, return_value, 3)
}
func PhpEscapeShellCmd(str *byte) *zend.ZendString {
	var x int
	var y int
	var l int = strlen(str)
	var estimate uint64 = 2*uint64(l) + 1
	var cmd *zend.ZendString
	var p *byte = nil

	/* max command line length - two single quotes - \0 byte length */

	if l > CmdMaxLen-2-1 {
		core.PhpErrorDocref(nil, zend.E_ERROR, "Command exceeds the allowed length of %zu bytes", CmdMaxLen)
		return zend.ZSTR_EMPTY_ALLOC()
	}
	cmd = zend.ZendStringSafeAlloc(2, l, 0, 0)
	x = 0
	y = 0
	for ; x < l; x++ {
		var mb_len int = PhpMblen(str+x, l-x)

		/* skip non-valid multibyte characters */

		if mb_len < 0 {
			continue
		} else if mb_len > 1 {
			memcpy(cmd.GetVal()+y, str+x, mb_len)
			y += mb_len
			x += mb_len - 1
			continue
		}
		switch str[x] {
		case '"':
			fallthrough
		case '\'':
			if p == nil && b.Assign(&p, memchr(str+x+1, str[x], l-x-1)) {

			} else if p != nil && (*p) == str[x] {
				p = nil
			} else {
				cmd.GetVal()[b.PostInc(&y)] = '\\'
			}
			cmd.GetVal()[b.PostInc(&y)] = str[x]
		case '#':
			fallthrough
		case '&':
			fallthrough
		case ';':
			fallthrough
		case '`':
			fallthrough
		case '|':
			fallthrough
		case '*':
			fallthrough
		case '?':
			fallthrough
		case '~':
			fallthrough
		case '<':
			fallthrough
		case '>':
			fallthrough
		case '^':
			fallthrough
		case '(':
			fallthrough
		case ')':
			fallthrough
		case '[':
			fallthrough
		case ']':
			fallthrough
		case '{':
			fallthrough
		case '}':
			fallthrough
		case '$':
			fallthrough
		case '\\':
			fallthrough
		case 'x':
			fallthrough
		case 'x':
			cmd.GetVal()[b.PostInc(&y)] = '\\'
			fallthrough
		default:
			cmd.GetVal()[b.PostInc(&y)] = str[x]
		}
	}
	cmd.GetVal()[y] = '0'
	if y > CmdMaxLen+1 {
		core.PhpErrorDocref(nil, zend.E_ERROR, "Escaped command exceeds the allowed length of %zu bytes", CmdMaxLen)
		zend.ZendStringReleaseEx(cmd, 0)
		return zend.ZSTR_EMPTY_ALLOC()
	}
	if estimate-y > 4096 {

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

		cmd = zend.ZendStringTruncate(cmd, y, 0)

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

	}
	cmd.SetLen(y)
	return cmd
}
func PhpEscapeShellArg(str *byte) *zend.ZendString {
	var x int
	var y int = 0
	var l int = strlen(str)
	var cmd *zend.ZendString
	var estimate uint64 = 4*uint64(l) + 3

	/* max command line length - two single quotes - \0 byte length */

	if l > CmdMaxLen-2-1 {
		core.PhpErrorDocref(nil, zend.E_ERROR, "Argument exceeds the allowed length of %zu bytes", CmdMaxLen)
		return zend.ZSTR_EMPTY_ALLOC()
	}
	cmd = zend.ZendStringSafeAlloc(4, l, 2, 0)
	cmd.GetVal()[b.PostInc(&y)] = '\''
	for x = 0; x < l; x++ {
		var mb_len int = PhpMblen(str+x, l-x)

		/* skip non-valid multibyte characters */

		if mb_len < 0 {
			continue
		} else if mb_len > 1 {
			memcpy(cmd.GetVal()+y, str+x, mb_len)
			y += mb_len
			x += mb_len - 1
			continue
		}
		switch str[x] {
		case '\'':
			cmd.GetVal()[b.PostInc(&y)] = '\''
			cmd.GetVal()[b.PostInc(&y)] = '\\'
			cmd.GetVal()[b.PostInc(&y)] = '\''
			fallthrough
		default:
			cmd.GetVal()[b.PostInc(&y)] = str[x]
		}
	}
	cmd.GetVal()[b.PostInc(&y)] = '\''
	cmd.GetVal()[y] = '0'
	if y > CmdMaxLen+1 {
		core.PhpErrorDocref(nil, zend.E_ERROR, "Escaped argument exceeds the allowed length of %zu bytes", CmdMaxLen)
		zend.ZendStringReleaseEx(cmd, 0)
		return zend.ZSTR_EMPTY_ALLOC()
	}
	if estimate-y > 4096 {

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

		cmd = zend.ZendStringTruncate(cmd, y, 0)

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

	}
	cmd.SetLen(y)
	return cmd
}
func ZifEscapeshellcmd(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var command *byte
	var command_len int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &command, &command_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	if command_len != 0 {
		if command_len != strlen(command) {
			core.PhpErrorDocref(nil, zend.E_ERROR, "Input string contains NULL bytes")
			return
		}
		return_value.SetString(PhpEscapeShellCmd(command))
	} else {
		zend.ZVAL_EMPTY_STRING(return_value)
	}
}
func ZifEscapeshellarg(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var argument *byte
	var argument_len int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &argument, &argument_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	if argument != nil {
		if argument_len != strlen(argument) {
			core.PhpErrorDocref(nil, zend.E_ERROR, "Input string contains NULL bytes")
			return
		}
		return_value.SetString(PhpEscapeShellArg(argument))
	}
}
func ZifShellExec(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var in *r.FILE
	var command *byte
	var command_len int
	var ret *zend.ZendString
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &command, &command_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	if command_len == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Cannot execute a blank command")
		return_value.SetFalse()
		return
	}
	if strlen(command) != command_len {
		core.PhpErrorDocref(nil, zend.E_WARNING, "NULL byte detected. Possible attack")
		return_value.SetFalse()
		return
	}
	if b.Assign(&in, zend.VCWD_POPEN(command, "r")) == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to execute '%s'", command)
		return_value.SetFalse()
		return
	}
	stream = streams.PhpStreamFopenFromPipe(in, "rb")
	ret = core.PhpStreamCopyToMem(stream, core.PHP_STREAM_COPY_ALL, 0)
	core.PhpStreamClose(stream)
	if ret != nil && ret.GetLen() > 0 {
		return_value.SetString(ret)
	}
}
func ZifProcNice(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var pri zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &pri, &_dummy, 0, 0) == 0 {
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
			return_value.SetFalse()
			return
		}
		break
	}
	errno = 0
	core.PhpIgnoreValue(nice(pri))
	if errno {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Only a super user may attempt to increase the priority of a process")
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
