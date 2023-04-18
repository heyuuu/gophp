package standard

import (
	"github.com/heyuuu/gophp/builtin"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"math"
	"strings"
)

func PhpExec(type_ int, cmd *byte, array *types.Zval, return_value *types.Zval) int {
	var fp *r.File
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
				core.PUTS(builtin.CastStr(buf, bufl))
				if core.PhpOutputGetLevel() < 1 {
					core.SapiFlush()
				}
			} else if type_ == 2 {

				/* strip trailing whitespaces */

				l = bufl
				for builtin.PostDec(&l) > 0 && isspace((*uint8)(buf)[l]) {

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
				core.PUTS(b.CastStr(buf, bufl))
				if core.PhpOutputGetLevel() < 1 {
					core.SapiFlush()
				}
			}

			/* strip trailing whitespaces if we have not done so already */

			if type_ == 2 && buf != b || type_ != 2 {
				l = bufl
				for builtin.PostDec(&l) > 0 && isspace((*uint8)(buf)[l]) {

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

			return_value.SetStringVal(builtin.CastStr(buf, bufl))

		} else {
			return_value.SetStringVal("")
		}
	} else {
		var read ssize_t
		for builtin.Assign(&read, core.PhpStreamRead(stream, buf, core.EXEC_INPUT_BUF)) > 0 {
			core.PUTS(b.CastStr(buf, read))
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
			types.SeparateArray(ret_array)
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
func ZifExec(executeData zpp.Ex, return_value zpp.Ret, command *types.Zval, _ zpp.Opt, output zpp.RefZval, returnValue zpp.RefZval) {
	PhpExecEx(executeData, return_value, 0)
}
func ZifSystem(executeData zpp.Ex, return_value zpp.Ret, command *types.Zval, _ zpp.Opt, returnValue zpp.RefZval) {
	PhpExecEx(executeData, return_value, 1)
}
func ZifPassthru(executeData zpp.Ex, return_value zpp.Ret, command *types.Zval, _ zpp.Opt, returnValue zpp.RefZval) {
	PhpExecEx(executeData, return_value, 3)
}
func PhpEscapeShellCmd(str string) string {
	/* max command line length - two single quotes - \0 byte length */
	if len(str) > CmdMaxLen-2-1 {
		core.PhpErrorDocref(nil, faults.E_ERROR, "Command exceeds the allowed length of %zu bytes", CmdMaxLen)
		return ""
	}

	var buf strings.Builder
	var p byte = 0 // 可选值 0, ', "
	for i, r := range str {
		/* skip multibyte characters */
		if r > math.MaxUint8 {
			buf.WriteRune(r)
			continue
		}
		//
		c := byte(r)
		switch c {
		case '"', '\'':
			if p == 0 {
				if pos := strings.IndexByte(str[i+1:], c); pos >= 0 {
					p = c
				} else {
					buf.WriteByte('\\')
				}
			} else {
				if p == c {
					p = 0
				} else {
					buf.WriteRune('\\')
				}
			}
			buf.WriteByte(c)
		case '#', '&', ';', '`', '|', '*', '?', '~', '<', '>', '^', '(', ')', '[', ']', '{', '}', '$', '\\', '\x0A', '\xFF':
			buf.WriteByte('\\')
			buf.WriteByte(c)
		default:
			buf.WriteByte(c)
		}
	}
	if buf.Len() > CmdMaxLen+1 {
		core.PhpErrorDocref(nil, faults.E_ERROR, "Escaped command exceeds the allowed length of %zu bytes", CmdMaxLen)
		return ""
	}
	return buf.String()
}
func PhpEscapeShellArg(str string) string {
	/* max command line length - two single quotes - \0 byte length */
	if len(str) > CmdMaxLen-2-1 {
		core.PhpErrorDocref(nil, faults.E_ERROR, "Argument exceeds the allowed length of %zu bytes", CmdMaxLen)
		return ""
	}

	var buf strings.Builder
	buf.WriteByte('\'')
	for _, r := range str {
		/* skip multibyte characters */
		if !ascii.IsAsciiRune(r) {
			buf.WriteRune(r)
			continue
		}

		c := byte(r)
		switch c {
		case '\'':
			buf.WriteString(`'\'`)
			fallthrough
		default:
			buf.WriteByte(c)
		}
	}
	buf.WriteByte('\'')
	if buf.Len() > CmdMaxLen+1 {
		core.PhpErrorDocref(nil, faults.E_ERROR, "Escaped argument exceeds the allowed length of %zu bytes", CmdMaxLen)
		return ""
	}
	return buf.String()
}
func ZifEscapeshellcmd(command string) string {
	if command == "" {
		return ""
	}
	if pos := strings.IndexByte(command, 0); pos >= 0 {
		core.PhpErrorDocref(nil, faults.E_ERROR, "Input string contains NULL bytes")
	}

	return PhpEscapeShellCmd(command)
}
func ZifEscapeshellarg(arg string) string {
	if pos := strings.IndexByte(arg, 0); pos >= 0 {
		core.PhpErrorDocref(nil, faults.E_ERROR, "Input string contains NULL bytes")
	}
	return PhpEscapeShellArg(arg)
}
func ZifShellExec(executeData zpp.Ex, return_value zpp.Ret, cmd *types.Zval) {
	var in *r.File
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
