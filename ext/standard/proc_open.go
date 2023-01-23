// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/sapi/cli"
	"sik/zend"
)

// Source: <ext/standard/proc_open.h>

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
   | Author: Wez Furlong <wez@thebrainroom.com>                           |
   +----------------------------------------------------------------------+
*/

type PhpFileDescriptorT = int
type PhpProcessIdT = pid_t

/* Environment block under win32 is a NUL terminated sequence of NUL terminated
 * name=value strings.
 * Under unix, it is an argv style array.
 * */

// @type PhpProcessEnvT struct
type _phpProcessEnv = PhpProcessEnvT

// @type PhpProcessHandle struct

// Source: <ext/standard/proc_open.c>

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
   | Author: Wez Furlong <wez@thebrainroom.com>                           |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include < stdio . h >

// # include < ctype . h >

// # include "php_string.h"

// # include "ext/standard/head.h"

// # include "ext/standard/basic_functions.h"

// # include "ext/standard/file.h"

// # include "exec.h"

// # include "php_globals.h"

// # include "SAPI.h"

// # include "main/php_network.h"

// # include "zend_smart_string.h"

// # include < sys / wait . h >

// # include < signal . h >

// # include < sys / stat . h >

// # include < fcntl . h >

/* This symbol is defined in ext/standard/config.m4.
 * Essentially, it is set if you HAVE_FORK || PHP_WIN32
 * Other platforms may modify that configure check and add suitable #ifdefs
 * around the alternate code.
 * */

// # include "proc_open.h"

var LeProcOpen int

/* {{{ _php_array_to_envp */

func _phpArrayToEnvp(environment *zend.Zval, is_persistent int) PhpProcessEnvT {
	var element *zend.Zval
	var env PhpProcessEnvT
	var key *zend.ZendString
	var str *zend.ZendString
	var ep **byte
	var p *byte
	var cnt int
	var sizeenv int = 0
	var env_hash *zend.HashTable
	memset(&env, 0, g.SizeOf("env"))
	if environment == nil {
		return env
	}
	cnt = environment.value.arr.nNumOfElements
	if cnt < 1 {
		env.SetEnvarray((**byte)(g.CondF(is_persistent != 0, func() any { return zend.__zendCalloc(1, g.SizeOf("char *")) }, func() any { return zend._ecalloc(1, g.SizeOf("char *")) })))
		env.SetEnvp((*byte)(g.CondF(is_persistent != 0, func() any { return zend.__zendCalloc(4, 1) }, func() any { return zend._ecalloc(4, 1) })))
		return env
	}
	env_hash = (*zend.HashTable)(zend._emalloc(g.SizeOf("HashTable")))
	zend._zendHashInit(env_hash, cnt, nil, 0)

	/* first, we have to get the size of all the elements in the hash */

	for {
		var __ht *zend.HashTable = environment.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			key = _p.key
			element = _z
			str = zend.ZvalGetString(element)
			if str.len_ == 0 {
				zend.ZendStringReleaseEx(str, 0)
				continue
			}
			sizeenv += str.len_ + 1
			if key != nil && key.len_ != 0 {
				sizeenv += key.len_ + 1
				zend.ZendHashAddPtr(env_hash, key, str)
			} else {
				zend.ZendHashNextIndexInsertPtr(env_hash, str)
			}
		}
		break
	}
	env.SetEnvarray((**byte)(g.CondF(is_persistent != 0, func() any { return zend.__zendCalloc(cnt+1, g.SizeOf("char *")) }, func() any { return zend._ecalloc(cnt+1, g.SizeOf("char *")) })))
	ep = env.GetEnvarray()
	env.SetEnvp((*byte)(g.CondF(is_persistent != 0, func() any { return zend.__zendCalloc(sizeenv+4, 1) }, func() any { return zend._ecalloc(sizeenv+4, 1) })))
	p = env.GetEnvp()
	for {
		var __ht *zend.HashTable = env_hash
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			key = _p.key
			str = _z.value.ptr
			*ep = p
			ep++
			if key != nil {
				memcpy(p, key.val, key.len_)
				p += key.len_
				g.PostInc(&(*p)) = '='
			}
			memcpy(p, str.val, str.len_)
			p += str.len_
			g.PostInc(&(*p)) = '0'
			zend.ZendStringReleaseEx(str, 0)
		}
		break
	}
	r.Assert(uint32(p-env.GetEnvp()) <= sizeenv)
	zend.ZendHashDestroy(env_hash)
	zend._efree(env_hash)
	return env
}

/* }}} */

func _phpFreeEnvp(env PhpProcessEnvT, is_persistent int) {
	if env.GetEnvarray() != nil {
		g.CondF(is_persistent != 0, func() { return zend.Free(env.GetEnvarray()) }, func() { return zend._efree(env.GetEnvarray()) })
	}
	if env.GetEnvp() != nil {
		g.CondF(is_persistent != 0, func() { return zend.Free(env.GetEnvp()) }, func() { return zend._efree(env.GetEnvp()) })
	}
}

/* }}} */

func ProcOpenRsrcDtor(rsrc *zend.ZendResource) {
	var proc *PhpProcessHandle = (*PhpProcessHandle)(rsrc.ptr)
	var i int
	var wstatus int
	var waitpid_options int = 0
	var wait_pid pid_t

	/* Close all handles to avoid a deadlock */

	for i = 0; i < proc.GetNpipes(); i++ {
		if proc.GetPipes()[i] != 0 {
			zend.ZendGcDelref(&proc.GetPipes()[i].gc)
			zend.ZendListClose(proc.GetPipes()[i])
			proc.GetPipes()[i] = 0
		}
	}
	if FileGlobals.GetPcloseWait() == 0 {
		waitpid_options = WNOHANG
	}
	for {
		wait_pid = waitpid(proc.GetChild(), &wstatus, waitpid_options)
		if !(wait_pid == -1 && errno == EINTR) {
			break
		}
	}
	if wait_pid <= 0 {
		FileGlobals.SetPcloseRet(-1)
	} else {
		if WIFEXITED(wstatus) {
			wstatus = WEXITSTATUS(wstatus)
		}
		FileGlobals.SetPcloseRet(wstatus)
	}
	_phpFreeEnvp(proc.GetEnv(), proc.GetIsPersistent())
	g.CondF(proc.GetIsPersistent() != 0, func() { return zend.Free(proc.GetPipes()) }, func() { return zend._efree(proc.GetPipes()) })
	g.CondF(proc.GetIsPersistent() != 0, func() { return zend.Free(proc.GetCommand()) }, func() { return zend._efree(proc.GetCommand()) })
	g.CondF(proc.GetIsPersistent() != 0, func() { return zend.Free(proc) }, func() { return zend._efree(proc) })
}

/* }}} */

func ZmStartupProcOpen(type_ int, module_number int) int {
	LeProcOpen = zend.ZendRegisterListDestructorsEx(ProcOpenRsrcDtor, nil, "process", module_number)
	return zend.SUCCESS
}

/* }}} */

func ZifProcTerminate(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zproc *zend.Zval
	var proc *PhpProcessHandle
	var sig_no zend.ZendLong = SIGTERM
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
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

			if zend.ZendParseArgResource(_arg, &zproc, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgLong(_arg, &sig_no, &_dummy, 0, 0) == 0 {
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
	if g.Assign(&proc, (*PhpProcessHandle)(zend.ZendFetchResource(zproc.value.res, "process", LeProcOpen))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if kill(proc.GetChild(), sig_no) == 0 {
		return_value.u1.type_info = 3
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifProcClose(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zproc *zend.Zval
	var proc *PhpProcessHandle
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

			if zend.ZendParseArgResource(_arg, &zproc, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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
	if g.Assign(&proc, (*PhpProcessHandle)(zend.ZendFetchResource(zproc.value.res, "process", LeProcOpen))) == nil {
		return_value.u1.type_info = 2
		return
	}
	FileGlobals.SetPcloseWait(1)
	zend.ZendListClose(zproc.value.res)
	FileGlobals.SetPcloseWait(0)
	var __z *zend.Zval = return_value
	__z.value.lval = FileGlobals.GetPcloseRet()
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifProcGetStatus(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zproc *zend.Zval
	var proc *PhpProcessHandle
	var wstatus int
	var wait_pid pid_t
	var running int = 1
	var signaled int = 0
	var stopped int = 0
	var exitcode int = -1
	var termsig int = 0
	var stopsig int = 0
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

			if zend.ZendParseArgResource(_arg, &zproc, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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
	if g.Assign(&proc, (*PhpProcessHandle)(zend.ZendFetchResource(zproc.value.res, "process", LeProcOpen))) == nil {
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.AddAssocStringEx(return_value, "command", strlen("command"), proc.GetCommand())
	zend.AddAssocLongEx(return_value, "pid", strlen("pid"), zend.ZendLong(proc.GetChild()))
	errno = 0
	wait_pid = waitpid(proc.GetChild(), &wstatus, WNOHANG|WUNTRACED)
	if wait_pid == proc.GetChild() {
		if WIFEXITED(wstatus) {
			running = 0
			exitcode = WEXITSTATUS(wstatus)
		}
		if WIFSIGNALED(wstatus) {
			running = 0
			signaled = 1
			termsig = WTERMSIG(wstatus)
		}
		if WIFSTOPPED(wstatus) {
			stopped = 1
			stopsig = WSTOPSIG(wstatus)
		}
	} else if wait_pid == -1 {
		running = 0
	}
	zend.AddAssocBoolEx(return_value, "running", strlen("running"), running)
	zend.AddAssocBoolEx(return_value, "signaled", strlen("signaled"), signaled)
	zend.AddAssocBoolEx(return_value, "stopped", strlen("stopped"), stopped)
	zend.AddAssocLongEx(return_value, "exitcode", strlen("exitcode"), exitcode)
	zend.AddAssocLongEx(return_value, "termsig", strlen("termsig"), termsig)
	zend.AddAssocLongEx(return_value, "stopsig", strlen("stopsig"), stopsig)
}

/* }}} */

// #define close_descriptor(fd) close ( fd )

// #define DESC_PIPE       1

// #define DESC_FILE       2

// #define DESC_REDIRECT       3

// #define DESC_PARENT_MODE_WRITE       8

// @type PhpProcOpenDescriptorItem struct

/* }}} */

func GetValidArgString(zv *zend.Zval, elem_num int) *zend.ZendString {
	var str *zend.ZendString = zend.ZvalGetString(zv)
	if str == nil {
		return nil
	}
	if strlen(str.val) != str.len_ {
		core.PhpErrorDocref(nil, 1<<1, "Command array element %d contains a null byte", elem_num)
		zend.ZendStringRelease(str)
		return nil
	}
	return str
}

/* {{{ proto resource proc_open(string|array command, array descriptorspec, array &pipes [, string cwd [, array env [, array other_options]]])
   Run a process with more control over it's file descriptors */

func ZifProcOpen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var command_zv *zend.Zval
	var command *byte = nil
	var cwd *byte = nil
	var cwd_len int = 0
	var descriptorspec *zend.Zval
	var pipes *zend.Zval
	var environment *zend.Zval = nil
	var other_options *zend.Zval = nil
	var env PhpProcessEnvT
	var ndesc int = 0
	var i int
	var descitem *zend.Zval = nil
	var str_index *zend.ZendString
	var nindex zend.ZendUlong
	var descriptors *PhpProcOpenDescriptorItem = nil
	var ndescriptors_array int
	var argv **byte = nil
	var child PhpProcessIdT
	var proc *PhpProcessHandle
	var is_persistent int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 6
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

			zend.ZendParseArgZvalDeref(_arg, &command_zv, 0)
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

			if zend.ZendParseArgArray(_arg, &descriptorspec, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
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

			zend.ZendParseArgZvalDeref(_arg, &pipes, 0)
			_optional = 1
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

			if zend.ZendParseArgString(_arg, &cwd, &cwd_len, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
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

			if zend.ZendParseArgArray(_arg, &environment, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
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

			if zend.ZendParseArgArray(_arg, &other_options, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	memset(&env, 0, g.SizeOf("env"))
	if command_zv.u1.v.type_ == 7 {
		var arg_zv *zend.Zval
		var num_elems uint32 = command_zv.value.arr.nNumOfElements
		if num_elems == 0 {
			core.PhpErrorDocref(nil, 1<<1, "Command array must have at least one element")
			return_value.u1.type_info = 2
			return
		}
		argv = zend._safeEmalloc(g.SizeOf("char *"), num_elems+1, 0)
		i = 0
		for {
			var __ht *zend.HashTable = command_zv.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				arg_zv = _z
				var arg_str *zend.ZendString = GetValidArgString(arg_zv, i+1)
				if arg_str == nil {
					argv[i] = nil
					goto exit_fail
				}
				if i == 0 {
					if is_persistent != 0 {
						command = strdup(arg_str.val)
					} else {
						command = zend._estrdup(arg_str.val)
					}
				}
				argv[g.PostInc(&i)] = zend._estrdup(arg_str.val)
				zend.ZendStringRelease(arg_str)
			}
			break
		}
		argv[i] = nil

		/* As the array is non-empty, we should have found a command. */

		r.Assert(command != nil)

		/* As the array is non-empty, we should have found a command. */

	} else {
		if command_zv.u1.v.type_ != 6 {
			zend._convertToString(command_zv)
		}
		if is_persistent != 0 {
			command = strdup(command_zv.value.str.val)
		} else {
			command = zend._estrdup(command_zv.value.str.val)
		}
	}
	if environment != nil {
		env = _phpArrayToEnvp(environment, is_persistent)
	}
	ndescriptors_array = descriptorspec.value.arr.nNumOfElements
	descriptors = zend._safeEmalloc(g.SizeOf("struct php_proc_open_descriptor_item"), ndescriptors_array, 0)
	memset(descriptors, 0, g.SizeOf("struct php_proc_open_descriptor_item")*ndescriptors_array)

	/* walk the descriptor spec and set up files/pipes */

	for {
		var __ht *zend.HashTable = descriptorspec.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			nindex = _p.h
			str_index = _p.key
			descitem = _z
			var ztype *zend.Zval
			if str_index != nil {
				core.PhpErrorDocref(nil, 1<<1, "descriptor spec must be an integer indexed array")
				goto exit_fail
			}
			descriptors[ndesc].SetIndex(int(nindex))
			if descitem.u1.v.type_ == 9 {

				/* should be a stream - try and dup the descriptor */

				var stream *core.PhpStream
				var fd core.PhpSocketT
				if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(descitem, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
					return_value.u1.type_info = 2
					return
				}
				if zend.FAILURE == streams._phpStreamCast(stream, 1, (*any)(&fd), 0x8) {
					goto exit_fail
				}
				descriptors[ndesc].SetChildend(dup(fd))
				if descriptors[ndesc].GetChildend() < 0 {
					core.PhpErrorDocref(nil, 1<<1, "unable to dup File-Handle for descriptor "+"%"+"llu"+" - %s", nindex, strerror(errno))
					goto exit_fail
				}
				descriptors[ndesc].SetMode(2)
			} else if descitem.u1.v.type_ != 7 {
				core.PhpErrorDocref(nil, 1<<1, "Descriptor item must be either an array or a File-Handle")
				goto exit_fail
			} else {
				if g.Assign(&ztype, zend.ZendHashIndexFind(descitem.value.arr, 0)) != nil {
					if zend.TryConvertToString(ztype) == 0 {
						goto exit_fail
					}
				} else {
					core.PhpErrorDocref(nil, 1<<1, "Missing handle qualifier in array")
					goto exit_fail
				}
				if strcmp(ztype.value.str.val, "pipe") == 0 {
					var newpipe []PhpFileDescriptorT
					var zmode *zend.Zval
					if g.Assign(&zmode, zend.ZendHashIndexFind(descitem.value.arr, 1)) != nil {
						if zend.TryConvertToString(zmode) == 0 {
							goto exit_fail
						}
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Missing mode parameter for 'pipe'")
						goto exit_fail
					}
					descriptors[ndesc].SetMode(1)
					if 0 != pipe(newpipe) {
						core.PhpErrorDocref(nil, 1<<1, "unable to create pipe %s", strerror(errno))
						goto exit_fail
					}
					if strncmp(zmode.value.str.val, "w", 1) != 0 {
						descriptors[ndesc].SetParentend(newpipe[1])
						descriptors[ndesc].SetChildend(newpipe[0])
						descriptors[ndesc].SetMode(descriptors[ndesc].GetMode() | 8)
					} else {
						descriptors[ndesc].SetParentend(newpipe[0])
						descriptors[ndesc].SetChildend(newpipe[1])
					}
					if (descriptors[ndesc].GetMode() & 8) != 0 {
						descriptors[ndesc].SetModeFlags(O_WRONLY)
					} else {
						descriptors[ndesc].SetModeFlags(O_RDONLY)
					}
				} else if strcmp(ztype.value.str.val, "file") == 0 {
					var zfile *zend.Zval
					var zmode *zend.Zval
					var fd core.PhpSocketT
					var stream *core.PhpStream
					descriptors[ndesc].SetMode(2)
					if g.Assign(&zfile, zend.ZendHashIndexFind(descitem.value.arr, 1)) != nil {
						if zend.TryConvertToString(zfile) == 0 {
							goto exit_fail
						}
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Missing file name parameter for 'file'")
						goto exit_fail
					}
					if g.Assign(&zmode, zend.ZendHashIndexFind(descitem.value.arr, 2)) != nil {
						if zend.TryConvertToString(zmode) == 0 {
							goto exit_fail
						}
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Missing mode parameter for 'file'")
						goto exit_fail
					}

					/* try a wrapper */

					stream = streams._phpStreamOpenWrapperEx(zfile.value.str.val, zmode.value.str.val, 0x8|0x20, nil, nil)

					/* force into an fd */

					if stream == nil || zend.FAILURE == streams._phpStreamCast(stream, 0x40000000|1, (*any)(&fd), 0x8) {
						goto exit_fail
					}
					descriptors[ndesc].SetChildend(fd)
				} else if strcmp(ztype.value.str.val, "redirect") == 0 {
					var ztarget *zend.Zval = zend.ZendHashIndexFindDeref(descitem.value.arr, 1)
					var target *PhpProcOpenDescriptorItem = nil
					var childend PhpFileDescriptorT
					if ztarget == nil {
						core.PhpErrorDocref(nil, 1<<1, "Missing redirection target")
						goto exit_fail
					}
					if ztarget.u1.v.type_ != 4 {
						core.PhpErrorDocref(nil, 1<<1, "Redirection target must be an integer")
						goto exit_fail
					}
					for i = 0; i < ndesc; i++ {
						if descriptors[i].GetIndex() == ztarget.value.lval {
							target = &descriptors[i]
							break
						}
					}
					if target != nil {
						childend = target.GetChildend()
					} else {
						if ztarget.value.lval < 0 || ztarget.value.lval > 2 {
							core.PhpErrorDocref(nil, 1<<1, "Redirection target "+"%"+"lld"+" not found", ztarget.value.lval)
							goto exit_fail
						}

						/* Support referring to a stdin/stdout/stderr pipe adopted from the parent,
						 * which happens whenever an explicit override is not provided. */

						childend = ztarget.value.lval

						/* Support referring to a stdin/stdout/stderr pipe adopted from the parent,
						 * which happens whenever an explicit override is not provided. */

					}
					descriptors[ndesc].SetChildend(dup(childend))
					if descriptors[ndesc].GetChildend() < 0 {
						core.PhpErrorDocref(nil, 1<<1, "Failed to dup() for descriptor "+"%"+"lld"+" - %s", nindex, strerror(errno))
						goto exit_fail
					}
					descriptors[ndesc].SetMode(3)
				} else if strcmp(ztype.value.str.val, "null") == 0 {
					descriptors[ndesc].SetChildend(open("/dev/null", O_RDWR))
					if descriptors[ndesc].GetChildend() < 0 {
						core.PhpErrorDocref(nil, 1<<1, "Failed to open /dev/null - %s", strerror(errno))
						goto exit_fail
					}
					descriptors[ndesc].SetMode(2)
				} else if strcmp(ztype.value.str.val, "pty") == 0 {
					core.PhpErrorDocref(nil, 1<<1, "pty pseudo terminal not supported on this system")
					goto exit_fail
				} else {
					core.PhpErrorDocref(nil, 1<<1, "%s is not a valid descriptor spec/mode", ztype.value.str.val)
					goto exit_fail
				}
			}
			ndesc++
		}
		break
	}

	/* the unix way */

	child = fork()
	if child == 0 {

		/* this is the child process */

		/* close those descriptors that we just opened for the parent stuff,
		 * dup new descriptors into required descriptors and close the original
		 * cruft */

		for i = 0; i < ndesc; i++ {
			switch descriptors[i].GetMode() & ^8 {
			case 1:
				close(descriptors[i].GetParentend())
				break
			}
			if dup2(descriptors[i].GetChildend(), descriptors[i].GetIndex()) < 0 {
				r.Perror("dup2")
			}
			if descriptors[i].GetChildend() != descriptors[i].GetIndex() {
				close(descriptors[i].GetChildend())
			}
		}
		if cwd != nil {
			void(chdir(cwd))
		}
		if argv != nil {

			/* execvpe() is non-portable, use environ instead. */

			if env.GetEnvarray() != nil {
				cli.Environ = env.GetEnvarray()
			}
			execvp(command, argv)
		} else {
			if env.GetEnvarray() != nil {
				execle("/bin/sh", "sh", "-c", command, nil, env.GetEnvarray())
			} else {
				execl("/bin/sh", "sh", "-c", command, nil)
			}
		}
		_exit(127)
	} else if child < 0 {

		/* failed to fork() */

		for i = 0; i < ndesc; i++ {
			close(descriptors[i].GetChildend())
			if descriptors[i].GetParentend() != 0 {
				close(descriptors[i].GetParentend())
			}
		}
		core.PhpErrorDocref(nil, 1<<1, "fork failed - %s", strerror(errno))
		goto exit_fail
	}

	/* we forked/spawned and this is the parent */

	pipes = zend.ZendTryArrayInit(pipes)
	if pipes == nil {
		goto exit_fail
	}
	proc = (*PhpProcessHandle)(g.CondF(is_persistent != 0, func() any { return zend.__zendMalloc(g.SizeOf("struct php_process_handle")) }, func() any { return zend._emalloc(g.SizeOf("struct php_process_handle")) }))
	proc.SetIsPersistent(is_persistent)
	proc.SetCommand(command)
	if is_persistent != 0 {
		proc.SetPipes(zend.__zendMalloc(g.SizeOf("zend_resource *") * ndesc))
	} else {
		proc.SetPipes(zend._emalloc(g.SizeOf("zend_resource *") * ndesc))
	}
	proc.SetNpipes(ndesc)
	proc.SetChild(child)
	proc.SetEnv(env)

	/* clean up all the child ends and then open streams on the parent
	 * ends, where appropriate */

	for i = 0; i < ndesc; i++ {
		var mode_string *byte = nil
		var stream *core.PhpStream = nil
		close(descriptors[i].GetChildend())
		switch descriptors[i].GetMode() & ^8 {
		case 1:
			switch descriptors[i].GetModeFlags() {
			case O_WRONLY:
				mode_string = "w"
				break
			case O_RDONLY:
				mode_string = "r"
				break
			case O_RDWR:
				mode_string = "r+"
				break
			}
			stream = streams._phpStreamFopenFromFd(descriptors[i].GetParentend(), mode_string, nil)
			if stream != nil {
				var retfp zend.Zval

				/* nasty hack; don't copy it */

				stream.flags |= 0x1
				var __z *zend.Zval = &retfp
				__z.value.res = stream.res
				__z.u1.type_info = 9 | 1<<0<<8
				stream.__exposed = 1
				zend.AddIndexZval(pipes, descriptors[i].GetIndex(), &retfp)
				proc.GetPipes()[i] = retfp.value.res
				zend.ZvalAddrefP(&retfp)
			}
			break
		default:
			proc.GetPipes()[i] = nil
		}
	}
	if argv != nil {
		var arg **byte = argv
		for (*arg) != nil {
			zend._efree(*arg)
			arg++
		}
		zend._efree(argv)
	}
	zend._efree(descriptors)
	var __z *zend.Zval = return_value
	__z.value.res = zend.ZendRegisterResource(proc, LeProcOpen)
	__z.u1.type_info = 9 | 1<<0<<8
	return
exit_fail:
	if descriptors != nil {
		zend._efree(descriptors)
	}
	_phpFreeEnvp(env, is_persistent)
	if command != nil {
		g.CondF(is_persistent != 0, func() { return zend.Free(command) }, func() { return zend._efree(command) })
	}
	if argv != nil {
		var arg **byte = argv
		for (*arg) != nil {
			zend._efree(*arg)
			arg++
		}
		zend._efree(argv)
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */
