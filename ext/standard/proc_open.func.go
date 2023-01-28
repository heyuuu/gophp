// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	"sik/sapi/cli"
	"sik/zend"
)

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
	memset(&env, 0, b.SizeOf("env"))
	if environment == nil {
		return env
	}
	cnt = zend.Z_ARRVAL_P(environment).GetNNumOfElements()
	if cnt < 1 {
		env.SetEnvarray((**byte)(zend.Pecalloc(1, b.SizeOf("char *"), is_persistent)))
		env.SetEnvp((*byte)(zend.Pecalloc(4, 1, is_persistent)))
		return env
	}
	zend.ALLOC_HASHTABLE(env_hash)
	zend.ZendHashInit(env_hash, cnt, nil, nil, 0)

	/* first, we have to get the size of all the elements in the hash */

	for {
		var __ht *zend.HashTable = environment.GetArr()
		var _p *zend.Bucket = __ht.GetArData()
		var _end *zend.Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *zend.Zval = _p.GetVal()

			if _z.IsType(zend.IS_UNDEF) {
				continue
			}
			key = _p.GetKey()
			element = _z
			str = zend.ZvalGetString(element)
			if str.GetLen() == 0 {
				zend.ZendStringReleaseEx(str, 0)
				continue
			}
			sizeenv += str.GetLen() + 1
			if key != nil && key.GetLen() != 0 {
				sizeenv += key.GetLen() + 1
				zend.ZendHashAddPtr(env_hash, key, str)
			} else {
				zend.ZendHashNextIndexInsertPtr(env_hash, str)
			}
		}
		break
	}
	env.SetEnvarray((**byte)(zend.Pecalloc(cnt+1, b.SizeOf("char *"), is_persistent)))
	ep = env.GetEnvarray()
	env.SetEnvp((*byte)(zend.Pecalloc(sizeenv+4, 1, is_persistent)))
	p = env.GetEnvp()
	for {
		var __ht *zend.HashTable = env_hash
		var _p *zend.Bucket = __ht.GetArData()
		var _end *zend.Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *zend.Zval = _p.GetVal()

			if _z.IsType(zend.IS_UNDEF) {
				continue
			}
			key = _p.GetKey()
			str = _z.GetPtr()
			*ep = p
			ep++
			if key != nil {
				memcpy(p, key.GetVal(), key.GetLen())
				p += key.GetLen()
				b.PostInc(&(*p)) = '='
			}
			memcpy(p, str.GetVal(), str.GetLen())
			p += str.GetLen()
			b.PostInc(&(*p)) = '0'
			zend.ZendStringReleaseEx(str, 0)
		}
		break
	}
	r.Assert(uint32(p-env.GetEnvp()) <= sizeenv)
	zend.ZendHashDestroy(env_hash)
	zend.FREE_HASHTABLE(env_hash)
	return env
}
func _phpFreeEnvp(env PhpProcessEnvT, is_persistent int) {
	if env.GetEnvarray() != nil {
		zend.Pefree(env.GetEnvarray(), is_persistent)
	}
	if env.GetEnvp() != nil {
		zend.Pefree(env.GetEnvp(), is_persistent)
	}
}
func ProcOpenRsrcDtor(rsrc *zend.ZendResource) {
	var proc *PhpProcessHandle = (*PhpProcessHandle)(rsrc.GetPtr())
	var i int
	var wstatus int
	var waitpid_options int = 0
	var wait_pid pid_t

	/* Close all handles to avoid a deadlock */

	for i = 0; i < proc.GetNpipes(); i++ {
		if proc.GetPipes()[i] != 0 {
			proc.GetPipes()[i].DelRefcount()
			zend.ZendListClose(proc.GetPipes()[i])
			proc.GetPipes()[i] = 0
		}
	}
	if !(FG(pclose_wait)) {
		waitpid_options = WNOHANG
	}
	for {
		wait_pid = waitpid(proc.GetChild(), &wstatus, waitpid_options)
		if !(wait_pid == -1 && errno == EINTR) {
			break
		}
	}
	if wait_pid <= 0 {
		FG(pclose_ret) = -1
	} else {
		if WIFEXITED(wstatus) {
			wstatus = WEXITSTATUS(wstatus)
		}
		FG(pclose_ret) = wstatus
	}
	_phpFreeEnvp(proc.GetEnv(), proc.GetIsPersistent())
	zend.Pefree(proc.GetPipes(), proc.GetIsPersistent())
	zend.Pefree(proc.GetCommand(), proc.GetIsPersistent())
	zend.Pefree(proc, proc.GetIsPersistent())
}
func ZmStartupProcOpen(type_ int, module_number int) int {
	LeProcOpen = zend.ZendRegisterListDestructorsEx(ProcOpenRsrcDtor, nil, "process", module_number)
	return zend.SUCCESS
}
func ZifProcTerminate(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zproc *zend.Zval
	var proc *PhpProcessHandle
	var sig_no zend.ZendLong = SIGTERM
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
			if zend.ZendParseArgResource(_arg, &zproc, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &sig_no, &_dummy, 0, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if b.Assign(&proc, (*PhpProcessHandle)(zend.ZendFetchResource(zproc.GetRes(), "process", LeProcOpen))) == nil {
		zend.RETVAL_FALSE
		return
	}
	if kill(proc.GetChild(), sig_no) == 0 {
		zend.RETVAL_TRUE
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifProcClose(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zproc *zend.Zval
	var proc *PhpProcessHandle
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
			if zend.ZendParseArgResource(_arg, &zproc, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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
	if b.Assign(&proc, (*PhpProcessHandle)(zend.ZendFetchResource(zproc.GetRes(), "process", LeProcOpen))) == nil {
		zend.RETVAL_FALSE
		return
	}
	FG(pclose_wait) = 1
	zend.ZendListClose(zproc.GetRes())
	FG(pclose_wait) = 0
	zend.RETVAL_LONG(FG(pclose_ret))
	return
}
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
			if zend.ZendParseArgResource(_arg, &zproc, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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
	if b.Assign(&proc, (*PhpProcessHandle)(zend.ZendFetchResource(zproc.GetRes(), "process", LeProcOpen))) == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.ArrayInit(return_value)
	zend.AddAssocString(return_value, "command", proc.GetCommand())
	zend.AddAssocLong(return_value, "pid", zend.ZendLong(proc.GetChild()))
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
	zend.AddAssocBool(return_value, "running", running)
	zend.AddAssocBool(return_value, "signaled", signaled)
	zend.AddAssocBool(return_value, "stopped", stopped)
	zend.AddAssocLong(return_value, "exitcode", exitcode)
	zend.AddAssocLong(return_value, "termsig", termsig)
	zend.AddAssocLong(return_value, "stopsig", stopsig)
}
func CloseDescriptor(fd PhpFileDescriptorT) __auto__ { return close(fd) }
func GetValidArgString(zv *zend.Zval, elem_num int) *zend.ZendString {
	var str *zend.ZendString = zend.ZvalGetString(zv)
	if str == nil {
		return nil
	}
	if strlen(str.GetVal()) != str.GetLen() {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Command array element %d contains a null byte", elem_num)
		zend.ZendStringRelease(str)
		return nil
	}
	return str
}
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
			zend.ZendParseArgZvalDeref(_arg, &command_zv, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &descriptorspec, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &pipes, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &cwd, &cwd_len, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &environment, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &other_options, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	memset(&env, 0, b.SizeOf("env"))
	if command_zv.IsType(zend.IS_ARRAY) {
		var arg_zv *zend.Zval
		var num_elems uint32 = zend.Z_ARRVAL_P(command_zv).GetNNumOfElements()
		if num_elems == 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Command array must have at least one element")
			zend.RETVAL_FALSE
			return
		}
		argv = zend.SafeEmalloc(b.SizeOf("char *"), num_elems+1, 0)
		i = 0
		for {
			var __ht *zend.HashTable = command_zv.GetArr()
			var _p *zend.Bucket = __ht.GetArData()
			var _end *zend.Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *zend.Zval = _p.GetVal()

				if _z.IsType(zend.IS_UNDEF) {
					continue
				}
				arg_zv = _z
				var arg_str *zend.ZendString = GetValidArgString(arg_zv, i+1)
				if arg_str == nil {
					argv[i] = nil
					goto exit_fail
				}
				if i == 0 {
					command = zend.Pestrdup(arg_str.GetVal(), is_persistent)
				}
				argv[b.PostInc(&i)] = zend.Estrdup(arg_str.GetVal())
				zend.ZendStringRelease(arg_str)
			}
			break
		}
		argv[i] = nil

		/* As the array is non-empty, we should have found a command. */

		zend.ZEND_ASSERT(command != nil)

		/* As the array is non-empty, we should have found a command. */

	} else {
		zend.ConvertToString(command_zv)
		command = zend.Pestrdup(zend.Z_STRVAL_P(command_zv), is_persistent)
	}
	if environment != nil {
		env = _phpArrayToEnvp(environment, is_persistent)
	}
	ndescriptors_array = zend.Z_ARRVAL_P(descriptorspec).GetNNumOfElements()
	descriptors = zend.SafeEmalloc(b.SizeOf("struct php_proc_open_descriptor_item"), ndescriptors_array, 0)
	memset(descriptors, 0, b.SizeOf("struct php_proc_open_descriptor_item")*ndescriptors_array)

	/* walk the descriptor spec and set up files/pipes */

	for {
		var __ht *zend.HashTable = descriptorspec.GetArr()
		var _p *zend.Bucket = __ht.GetArData()
		var _end *zend.Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *zend.Zval = _p.GetVal()

			if _z.IsType(zend.IS_UNDEF) {
				continue
			}
			nindex = _p.GetH()
			str_index = _p.GetKey()
			descitem = _z
			var ztype *zend.Zval
			if str_index != nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "descriptor spec must be an integer indexed array")
				goto exit_fail
			}
			descriptors[ndesc].SetIndex(int(nindex))
			if descitem.IsType(zend.IS_RESOURCE) {

				/* should be a stream - try and dup the descriptor */

				var stream *core.PhpStream
				var fd core.PhpSocketT
				core.PhpStreamFromZval(stream, descitem)
				if zend.FAILURE == core.PhpStreamCast(stream, core.PHP_STREAM_AS_FD, (*any)(&fd), core.REPORT_ERRORS) {
					goto exit_fail
				}
				descriptors[ndesc].SetChildend(dup(fd))
				if descriptors[ndesc].GetChildend() < 0 {
					core.PhpErrorDocref(nil, zend.E_WARNING, "unable to dup File-Handle for descriptor "+zend.ZEND_ULONG_FMT+" - %s", nindex, strerror(errno))
					goto exit_fail
				}
				descriptors[ndesc].SetMode(DESC_FILE)
			} else if descitem.GetType() != zend.IS_ARRAY {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Descriptor item must be either an array or a File-Handle")
				goto exit_fail
			} else {
				if b.Assign(&ztype, zend.ZendHashIndexFind(descitem.GetArr(), 0)) != nil {
					if zend.TryConvertToString(ztype) == 0 {
						goto exit_fail
					}
				} else {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Missing handle qualifier in array")
					goto exit_fail
				}
				if strcmp(zend.Z_STRVAL_P(ztype), "pipe") == 0 {
					var newpipe []PhpFileDescriptorT
					var zmode *zend.Zval
					if b.Assign(&zmode, zend.ZendHashIndexFind(descitem.GetArr(), 1)) != nil {
						if zend.TryConvertToString(zmode) == 0 {
							goto exit_fail
						}
					} else {
						core.PhpErrorDocref(nil, zend.E_WARNING, "Missing mode parameter for 'pipe'")
						goto exit_fail
					}
					descriptors[ndesc].SetMode(DESC_PIPE)
					if 0 != pipe(newpipe) {
						core.PhpErrorDocref(nil, zend.E_WARNING, "unable to create pipe %s", strerror(errno))
						goto exit_fail
					}
					if strncmp(zend.Z_STRVAL_P(zmode), "w", 1) != 0 {
						descriptors[ndesc].SetParentend(newpipe[1])
						descriptors[ndesc].SetChildend(newpipe[0])
						descriptors[ndesc].SetMode(descriptors[ndesc].GetMode() | DESC_PARENT_MODE_WRITE)
					} else {
						descriptors[ndesc].SetParentend(newpipe[0])
						descriptors[ndesc].SetChildend(newpipe[1])
					}
					if (descriptors[ndesc].GetMode() & DESC_PARENT_MODE_WRITE) != 0 {
						descriptors[ndesc].SetModeFlags(O_WRONLY)
					} else {
						descriptors[ndesc].SetModeFlags(O_RDONLY)
					}
				} else if strcmp(zend.Z_STRVAL_P(ztype), "file") == 0 {
					var zfile *zend.Zval
					var zmode *zend.Zval
					var fd core.PhpSocketT
					var stream *core.PhpStream
					descriptors[ndesc].SetMode(DESC_FILE)
					if b.Assign(&zfile, zend.ZendHashIndexFind(descitem.GetArr(), 1)) != nil {
						if zend.TryConvertToString(zfile) == 0 {
							goto exit_fail
						}
					} else {
						core.PhpErrorDocref(nil, zend.E_WARNING, "Missing file name parameter for 'file'")
						goto exit_fail
					}
					if b.Assign(&zmode, zend.ZendHashIndexFind(descitem.GetArr(), 2)) != nil {
						if zend.TryConvertToString(zmode) == 0 {
							goto exit_fail
						}
					} else {
						core.PhpErrorDocref(nil, zend.E_WARNING, "Missing mode parameter for 'file'")
						goto exit_fail
					}

					/* try a wrapper */

					stream = core.PhpStreamOpenWrapper(zend.Z_STRVAL_P(zfile), zend.Z_STRVAL_P(zmode), core.REPORT_ERRORS|core.STREAM_WILL_CAST, nil)

					/* force into an fd */

					if stream == nil || zend.FAILURE == core.PhpStreamCast(stream, core.PHP_STREAM_CAST_RELEASE|core.PHP_STREAM_AS_FD, (*any)(&fd), core.REPORT_ERRORS) {
						goto exit_fail
					}
					descriptors[ndesc].SetChildend(fd)
				} else if strcmp(zend.Z_STRVAL_P(ztype), "redirect") == 0 {
					var ztarget *zend.Zval = zend.ZendHashIndexFindDeref(descitem.GetArr(), 1)
					var target *PhpProcOpenDescriptorItem = nil
					var childend PhpFileDescriptorT
					if ztarget == nil {
						core.PhpErrorDocref(nil, zend.E_WARNING, "Missing redirection target")
						goto exit_fail
					}
					if ztarget.GetType() != zend.IS_LONG {
						core.PhpErrorDocref(nil, zend.E_WARNING, "Redirection target must be an integer")
						goto exit_fail
					}
					for i = 0; i < ndesc; i++ {
						if descriptors[i].GetIndex() == ztarget.GetLval() {
							target = &descriptors[i]
							break
						}
					}
					if target != nil {
						childend = target.GetChildend()
					} else {
						if ztarget.GetLval() < 0 || ztarget.GetLval() > 2 {
							core.PhpErrorDocref(nil, zend.E_WARNING, "Redirection target "+zend.ZEND_LONG_FMT+" not found", ztarget.GetLval())
							goto exit_fail
						}

						/* Support referring to a stdin/stdout/stderr pipe adopted from the parent,
						 * which happens whenever an explicit override is not provided. */

						childend = ztarget.GetLval()

						/* Support referring to a stdin/stdout/stderr pipe adopted from the parent,
						 * which happens whenever an explicit override is not provided. */

					}
					descriptors[ndesc].SetChildend(dup(childend))
					if descriptors[ndesc].GetChildend() < 0 {
						core.PhpErrorDocref(nil, zend.E_WARNING, "Failed to dup() for descriptor "+zend.ZEND_LONG_FMT+" - %s", nindex, strerror(errno))
						goto exit_fail
					}
					descriptors[ndesc].SetMode(DESC_REDIRECT)
				} else if strcmp(zend.Z_STRVAL_P(ztype), "null") == 0 {
					descriptors[ndesc].SetChildend(open("/dev/null", O_RDWR))
					if descriptors[ndesc].GetChildend() < 0 {
						core.PhpErrorDocref(nil, zend.E_WARNING, "Failed to open /dev/null - %s", strerror(errno))
						goto exit_fail
					}
					descriptors[ndesc].SetMode(DESC_FILE)
				} else if strcmp(zend.Z_STRVAL_P(ztype), "pty") == 0 {
					core.PhpErrorDocref(nil, zend.E_WARNING, "pty pseudo terminal not supported on this system")
					goto exit_fail
				} else {
					core.PhpErrorDocref(nil, zend.E_WARNING, "%s is not a valid descriptor spec/mode", zend.Z_STRVAL_P(ztype))
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
			switch descriptors[i].GetMode() & ^DESC_PARENT_MODE_WRITE {
			case DESC_PIPE:
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
			core.PhpIgnoreValue(chdir(cwd))
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
		core.PhpErrorDocref(nil, zend.E_WARNING, "fork failed - %s", strerror(errno))
		goto exit_fail
	}

	/* we forked/spawned and this is the parent */

	pipes = zend.ZendTryArrayInit(pipes)
	if pipes == nil {
		goto exit_fail
	}
	proc = (*PhpProcessHandle)(zend.Pemalloc(b.SizeOf("struct php_process_handle"), is_persistent))
	proc.SetIsPersistent(is_persistent)
	proc.SetCommand(command)
	proc.SetPipes(zend.Pemalloc(b.SizeOf("zend_resource *")*ndesc, is_persistent))
	proc.SetNpipes(ndesc)
	proc.SetChild(child)
	proc.SetEnv(env)

	/* clean up all the child ends and then open streams on the parent
	 * ends, where appropriate */

	for i = 0; i < ndesc; i++ {
		var mode_string *byte = nil
		var stream *core.PhpStream = nil
		CloseDescriptor(descriptors[i].GetChildend())
		switch descriptors[i].GetMode() & ^DESC_PARENT_MODE_WRITE {
		case DESC_PIPE:
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
			stream = streams.PhpStreamFopenFromFd(descriptors[i].GetParentend(), mode_string, nil)
			if stream != nil {
				var retfp zend.Zval

				/* nasty hack; don't copy it */

				stream.AddFlags(core.PHP_STREAM_FLAG_NO_SEEK)
				core.PhpStreamToZval(stream, &retfp)
				zend.AddIndexZval(pipes, descriptors[i].GetIndex(), &retfp)
				proc.GetPipes()[i] = retfp.GetRes()
				zend.Z_ADDREF(retfp)
			}
			break
		default:
			proc.GetPipes()[i] = nil
		}
	}
	if argv != nil {
		var arg **byte = argv
		for (*arg) != nil {
			zend.Efree(*arg)
			arg++
		}
		zend.Efree(argv)
	}
	zend.Efree(descriptors)
	zend.ZVAL_RES(return_value, zend.ZendRegisterResource(proc, LeProcOpen))
	return
exit_fail:
	if descriptors != nil {
		zend.Efree(descriptors)
	}
	_phpFreeEnvp(env, is_persistent)
	if command != nil {
		zend.Pefree(command, is_persistent)
	}
	if argv != nil {
		var arg **byte = argv
		for (*arg) != nil {
			zend.Efree(*arg)
			arg++
		}
		zend.Efree(argv)
	}
	zend.RETVAL_FALSE
	return
}
