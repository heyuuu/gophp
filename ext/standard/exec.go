// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/exec.h>

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
   | Author: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                        |
   +----------------------------------------------------------------------+
*/

// #define EXEC_H

// Source: <ext/standard/exec.c>

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
   | Author: Rasmus Lerdorf <rasmus@php.net>                              |
   |         Ilia Alshanetsky <iliaa@php.net>                             |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include "php.h"

// # include < ctype . h >

// # include "php_string.h"

// # include "ext/standard/head.h"

// # include "ext/standard/file.h"

// # include "basic_functions.h"

// # include "exec.h"

// # include "php_globals.h"

// # include "SAPI.h"

// # include < sys / wait . h >

// # include < signal . h >

// # include < sys / types . h >

// # include < sys / stat . h >

// # include < fcntl . h >

// # include < unistd . h >

// # include < limits . h >

var CmdMaxLen int

/* {{{ PHP_MINIT_FUNCTION(exec) */

func ZmStartupExec(type_ int, module_number int) int {
	/* This is just an arbitrary value for the fallback case. */

	CmdMaxLen = 4096
	return zend.SUCCESS
}

/* }}} */

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
	fp = popen(cmd, "r")
	if fp == nil {
		core.PhpErrorDocref(nil, 1<<1, "Unable to fork [%s]", cmd)
		goto err
	}
	stream = streams._phpStreamFopenFromPipe(fp, "rb")
	buf = (*byte)(zend._emalloc(4096))
	buflen = 4096
	if type_ != 3 {
		b = buf
		for streams._phpStreamGetLine(stream, b, 4096, &bufl) != nil {

			/* no new line found, let's read some more */

			if b[bufl-1] != '\n' && streams._phpStreamEof(stream) == 0 {
				if buflen < bufl+(b-buf)+4096 {
					bufl += b - buf
					buflen = bufl + 4096
					buf = zend._erealloc(buf, buflen)
					b = buf + bufl
				} else {
					b += bufl
				}
				continue
			} else if b != buf {
				bufl += b - buf
			}
			if type_ == 1 {
				core.PhpOutputWrite(buf, bufl)
				if core.PhpOutputGetLevel() < 1 {
					core.SapiFlush()
				}
			} else if type_ == 2 {

				/* strip trailing whitespaces */

				l = bufl
				for g.PostDec(&l) > 0 && isspace((*uint8)(buf)[l]) {

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
				core.PhpOutputWrite(buf, bufl)
				if core.PhpOutputGetLevel() < 1 {
					core.SapiFlush()
				}
			}

			/* strip trailing whitespaces if we have not done so already */

			if type_ == 2 && buf != b || type_ != 2 {
				l = bufl
				for g.PostDec(&l) > 0 && isspace((*uint8)(buf)[l]) {

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

			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(buf, bufl, 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8

			/* Return last line from the shell command */

		} else {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendEmptyString
			__z.value.str = __s
			__z.u1.type_info = 6
		}
	} else {
		var read ssize_t
		for g.Assign(&read, streams._phpStreamRead(stream, buf, 4096)) > 0 {
			core.PhpOutputWrite(buf, read)
		}
	}
	pclose_return = streams._phpStreamFree(stream, 1|2)
	zend._efree(buf)
done:
	if d != nil {
		zend._efree(d)
	}
	return pclose_return
err:
	pclose_return = -1
	goto done
}

/* }}} */

func PhpExecEx(execute_data *zend.ZendExecuteData, return_value *zend.Zval, mode int) {
	var cmd *byte
	var cmd_len int
	var ret_code *zend.Zval = nil
	var ret_array *zend.Zval = nil
	var ret int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = g.Cond(mode != 0, 2, 3)
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &cmd, &cmd_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			if mode == 0 {
				_i++
				r.Assert(_i <= _min_num_args || _optional == 1)
				r.Assert(_i > _min_num_args || _optional == 0)
				if _optional != 0 {
					if _i > _num_args {
						break
					}
				}
				_real_arg++
				_arg = _real_arg

				zend.ZendParseArgZvalDeref(_arg, &ret_array, 0)
			}
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &ret_code, 0)
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
	if cmd_len == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Cannot execute a blank command")
		return_value.u1.type_info = 2
		return
	}
	if strlen(cmd) != cmd_len {
		core.PhpErrorDocref(nil, 1<<1, "NULL byte detected. Possible attack")
		return_value.u1.type_info = 2
		return
	}
	if ret_array == nil {
		ret = PhpExec(mode, cmd, nil, return_value)
	} else {
		if &(*ret_array).value.ref.val.u1.v.type_ == 7 {
			if ret_array.u1.v.type_ == 10 {
				ret_array = &(*ret_array).value.ref.val
			}
			var _zv *zend.Zval = ret_array
			var _arr *zend.ZendArray = _zv.value.arr
			if zend.ZendGcRefcount(&_arr.gc) > 1 {
				if _zv.u1.v.type_flags != 0 {
					zend.ZendGcDelref(&_arr.gc)
				}
				var __arr *zend.ZendArray = zend.ZendArrayDup(_arr)
				var __z *zend.Zval = _zv
				__z.value.arr = __arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			}
		} else {
			ret_array = zend.ZendTryArrayInit(ret_array)
			if ret_array == nil {
				return
			}
		}
		ret = PhpExec(2, cmd, ret_array, return_value)
	}
	if ret_code != nil {
		for {
			r.Assert(ret_code.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = ret_code
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefLong(ref, ret)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				__z.value.lval = ret
				__z.u1.type_info = 4
				break
			}
			break
		}
	}
}

/* }}} */

func ZifExec(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpExecEx(execute_data, return_value, 0)
}

/* }}} */

func ZifSystem(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpExecEx(execute_data, return_value, 1)
}

/* }}} */

func ZifPassthru(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpExecEx(execute_data, return_value, 3)
}

/* }}} */

func PhpEscapeShellCmd(str *byte) *zend.ZendString {
	var x int
	var y int
	var l int = strlen(str)
	var estimate uint64 = 2*uint64(l) + 1
	var cmd *zend.ZendString
	var p *byte = nil

	/* max command line length - two single quotes - \0 byte length */

	if l > CmdMaxLen-2-1 {
		core.PhpErrorDocref(nil, 1<<0, "Command exceeds the allowed length of %zu bytes", CmdMaxLen)
		return zend.ZendEmptyString
	}
	cmd = zend.ZendStringSafeAlloc(2, l, 0, 0)
	x = 0
	y = 0
	for ; x < l; x++ {
		var mb_len int = mblen(str+x, l-x)

		/* skip non-valid multibyte characters */

		if mb_len < 0 {
			continue
		} else if mb_len > 1 {
			memcpy(cmd.val+y, str+x, mb_len)
			y += mb_len
			x += mb_len - 1
			continue
		}
		switch str[x] {
		case '"':

		case '\'':
			if p == nil && g.Assign(&p, memchr(str+x+1, str[x], l-x-1)) {

			} else if p != nil && (*p) == str[x] {
				p = nil
			} else {
				cmd.val[g.PostInc(&y)] = '\\'
			}
			cmd.val[g.PostInc(&y)] = str[x]
			break
		case '#':

		case '&':

		case ';':

		case '`':

		case '|':

		case '*':

		case '?':

		case '~':

		case '<':

		case '>':

		case '^':

		case '(':

		case ')':

		case '[':

		case ']':

		case '{':

		case '}':

		case '$':

		case '\\':

		case 'x':

		case 'x':
			cmd.val[g.PostInc(&y)] = '\\'
		default:
			cmd.val[g.PostInc(&y)] = str[x]
		}
	}
	cmd.val[y] = '0'
	if y > CmdMaxLen+1 {
		core.PhpErrorDocref(nil, 1<<0, "Escaped command exceeds the allowed length of %zu bytes", CmdMaxLen)
		zend.ZendStringReleaseEx(cmd, 0)
		return zend.ZendEmptyString
	}
	if estimate-y > 4096 {

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

		cmd = zend.ZendStringTruncate(cmd, y, 0)

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

	}
	cmd.len_ = y
	return cmd
}

/* }}} */

func PhpEscapeShellArg(str *byte) *zend.ZendString {
	var x int
	var y int = 0
	var l int = strlen(str)
	var cmd *zend.ZendString
	var estimate uint64 = 4*uint64(l) + 3

	/* max command line length - two single quotes - \0 byte length */

	if l > CmdMaxLen-2-1 {
		core.PhpErrorDocref(nil, 1<<0, "Argument exceeds the allowed length of %zu bytes", CmdMaxLen)
		return zend.ZendEmptyString
	}
	cmd = zend.ZendStringSafeAlloc(4, l, 2, 0)
	cmd.val[g.PostInc(&y)] = '\''
	for x = 0; x < l; x++ {
		var mb_len int = mblen(str+x, l-x)

		/* skip non-valid multibyte characters */

		if mb_len < 0 {
			continue
		} else if mb_len > 1 {
			memcpy(cmd.val+y, str+x, mb_len)
			y += mb_len
			x += mb_len - 1
			continue
		}
		switch str[x] {
		case '\'':
			cmd.val[g.PostInc(&y)] = '\''
			cmd.val[g.PostInc(&y)] = '\\'
			cmd.val[g.PostInc(&y)] = '\''
		default:
			cmd.val[g.PostInc(&y)] = str[x]
		}
	}
	cmd.val[g.PostInc(&y)] = '\''
	cmd.val[y] = '0'
	if y > CmdMaxLen+1 {
		core.PhpErrorDocref(nil, 1<<0, "Escaped argument exceeds the allowed length of %zu bytes", CmdMaxLen)
		zend.ZendStringReleaseEx(cmd, 0)
		return zend.ZendEmptyString
	}
	if estimate-y > 4096 {

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

		cmd = zend.ZendStringTruncate(cmd, y, 0)

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

	}
	cmd.len_ = y
	return cmd
}

/* }}} */

func ZifEscapeshellcmd(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var command *byte
	var command_len int
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &command, &command_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	if command_len != 0 {
		if command_len != strlen(command) {
			core.PhpErrorDocref(nil, 1<<0, "Input string contains NULL bytes")
			return
		}
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = PhpEscapeShellCmd(command)
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
	}
}

/* }}} */

func ZifEscapeshellarg(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var argument *byte
	var argument_len int
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &argument, &argument_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	if argument != nil {
		if argument_len != strlen(argument) {
			core.PhpErrorDocref(nil, 1<<0, "Input string contains NULL bytes")
			return
		}
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = PhpEscapeShellArg(argument)
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
	}
}

/* }}} */

func ZifShellExec(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var in *r.FILE
	var command *byte
	var command_len int
	var ret *zend.ZendString
	var stream *core.PhpStream
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &command, &command_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	if command_len == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Cannot execute a blank command")
		return_value.u1.type_info = 2
		return
	}
	if strlen(command) != command_len {
		core.PhpErrorDocref(nil, 1<<1, "NULL byte detected. Possible attack")
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&in, popen(command, "r")) == nil {
		core.PhpErrorDocref(nil, 1<<1, "Unable to execute '%s'", command)
		return_value.u1.type_info = 2
		return
	}
	stream = streams._phpStreamFopenFromPipe(in, "rb")
	ret = streams._phpStreamCopyToMem(stream, size_t-1, 0)
	streams._phpStreamFree(stream, 1|2)
	if ret != nil && ret.len_ > 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = ret
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
	}
}

/* }}} */

/* {{{ proto bool proc_nice(int priority)
   Change the priority of the current process */

func ZifProcNice(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var pri zend.ZendLong
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &pri, &_dummy, 0, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	errno = 0
	void(nice(pri))
	if errno {
		core.PhpErrorDocref(nil, 1<<1, "Only a super user may attempt to increase the priority of a process")
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */
