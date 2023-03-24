// <<generate>>

package standard

import (
	"sik/builtin"
	b "sik/builtin"
	r "sik/builtin/file"
	"sik/core"
	"sik/core/streams"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZmStartupExec(type_ int, module_number int) int {
	/* This is just an arbitrary value for the fallback case. */

	CmdMaxLen = 4096
	return types.SUCCESS
}
func PhpExec(type_ int, cmd *byte, array *types.Zval, return_value *types.Zval) int {
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
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to fork [%s]", cmd)
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
func PhpExecEx(executeData *zend.ZendExecuteData, return_value *types.Zval, mode int) {
	var cmd *byte
	var cmd_len int
	var ret_code *types.Zval = nil
	var ret_array *types.Zval = nil
	var ret int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = b.Cond(mode != 0, 2, 3)

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			cmd, cmd_len = fp.ParseString()
			fp.StartOptional()
			if mode == 0 {
				ret_array = fp.ParseZval()
			}
			ret_code = fp.ParseZval()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if cmd_len == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Cannot execute a blank command")
		return_value.SetFalse()
		return
	}
	if strlen(cmd) != cmd_len {
		core.PhpErrorDocref(nil, faults.E_WARNING, "NULL byte detected. Possible attack")
		return_value.SetFalse()
		return
	}
	if ret_array == nil {
		ret = PhpExec(mode, cmd, nil, return_value)
	} else {
		if types.Z_REFVAL_P(ret_array).IsType(types.IS_ARRAY) {
			ret_array = types.ZVAL_DEREF(ret_array)
			types.SEPARATE_ARRAY(ret_array)
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
func ZifExec(executeData zpp.DefEx, return_value zpp.DefReturn, command *types.Zval, _ zpp.DefOpt, output zpp.DefRef, returnValue zpp.DefRef) {
	PhpExecEx(executeData, return_value, 0)
}
func ZifSystem(executeData zpp.DefEx, return_value zpp.DefReturn, command *types.Zval, _ zpp.DefOpt, returnValue zpp.DefRef) {
	PhpExecEx(executeData, return_value, 1)
}
func ZifPassthru(executeData zpp.DefEx, return_value zpp.DefReturn, command *types.Zval, _ zpp.DefOpt, returnValue zpp.DefRef) {
	PhpExecEx(executeData, return_value, 3)
}
func PhpEscapeShellCmd(str *byte) *types.String {
	var x int
	var y int
	var l int = strlen(str)
	var estimate uint64 = 2*uint64(l) + 1
	var cmd *types.String
	var p *byte = nil

	/* max command line length - two single quotes - \0 byte length */

	if l > CmdMaxLen-2-1 {
		core.PhpErrorDocref(nil, faults.E_ERROR, "Command exceeds the allowed length of %zu bytes", CmdMaxLen)
		return types.ZSTR_EMPTY_ALLOC()
	}
	cmd = types.ZendStringSafeAlloc(2, l, 0, 0)
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
		core.PhpErrorDocref(nil, faults.E_ERROR, "Escaped command exceeds the allowed length of %zu bytes", CmdMaxLen)
		types.ZendStringReleaseEx(cmd, 0)
		return types.ZSTR_EMPTY_ALLOC()
	}
	if estimate-y > 4096 {

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

		cmd = types.ZendStringTruncate(cmd, y, 0)

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

	}
	cmd.SetLen(y)
	return cmd
}
func PhpEscapeShellArg(str *byte) *types.String {
	var x int
	var y int = 0
	var l int = strlen(str)
	var cmd *types.String
	var estimate uint64 = 4*uint64(l) + 3

	/* max command line length - two single quotes - \0 byte length */

	if l > CmdMaxLen-2-1 {
		core.PhpErrorDocref(nil, faults.E_ERROR, "Argument exceeds the allowed length of %zu bytes", CmdMaxLen)
		return types.ZSTR_EMPTY_ALLOC()
	}
	cmd = types.ZendStringSafeAlloc(4, l, 2, 0)
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
		core.PhpErrorDocref(nil, faults.E_ERROR, "Escaped argument exceeds the allowed length of %zu bytes", CmdMaxLen)
		types.ZendStringReleaseEx(cmd, 0)
		return types.ZSTR_EMPTY_ALLOC()
	}
	if estimate-y > 4096 {

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

		cmd = types.ZendStringTruncate(cmd, y, 0)

		/* realloc if the estimate was way overill
		 * Arbitrary cutoff point of 4096 */

	}
	cmd.SetLen(y)
	return cmd
}
func ZifEscapeshellcmd(executeData zpp.DefEx, return_value zpp.DefReturn, command *types.Zval) {
	var command *byte
	var command_len int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			command, command_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if command_len != 0 {
		if command_len != strlen(command) {
			core.PhpErrorDocref(nil, faults.E_ERROR, "Input string contains NULL bytes")
			return
		}
		return_value.SetString(PhpEscapeShellCmd(command))
	} else {
		zend.ZVAL_EMPTY_STRING(return_value)
	}
}
func ZifEscapeshellarg(executeData zpp.DefEx, return_value zpp.DefReturn, arg *types.Zval) {
	var argument *byte
	var argument_len int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			argument, argument_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if argument != nil {
		if argument_len != strlen(argument) {
			core.PhpErrorDocref(nil, faults.E_ERROR, "Input string contains NULL bytes")
			return
		}
		return_value.SetString(PhpEscapeShellArg(argument))
	}
}
func ZifShellExec(executeData zpp.DefEx, return_value zpp.DefReturn, cmd *types.Zval) {
	var in *r.FILE
	var command *byte
	var command_len int
	var ret *types.String
	var stream *core.PhpStream
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			command, command_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if command_len == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Cannot execute a blank command")
		return_value.SetFalse()
		return
	}
	if strlen(command) != command_len {
		core.PhpErrorDocref(nil, faults.E_WARNING, "NULL byte detected. Possible attack")
		return_value.SetFalse()
		return
	}
	if b.Assign(&in, zend.VCWD_POPEN(command, "r")) == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to execute '%s'", command)
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
func ZifProcNice(executeData zpp.DefEx, return_value zpp.DefReturn, priority *types.Zval) {
	var pri zend.ZendLong
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			pri = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	errno = 0
	core.PhpIgnoreValue(nice(pri))
	if errno {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Only a super user may attempt to increase the priority of a process")
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
