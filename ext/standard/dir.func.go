// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/zend"
)

func DIRG(v __auto__) __auto__ { return DirGlobals.v }
func PhpSetDefaultDir(res *zend.ZendResource) {
	if DIRG(default_dir) {
		zend.ZendListDelete(DIRG(default_dir))
	}
	if res != nil {
		res.AddRefcount()
	}
	DIRG(default_dir) = res
}
func ZmActivateDir(type_ int, module_number int) int {
	DIRG(default_dir) = nil
	return zend.SUCCESS
}
func ZmStartupDir(type_ int, module_number int) int {
	var dirsep_str []byte
	var pathsep_str []byte
	var dir_class_entry zend.ZendClassEntry
	memset(&dir_class_entry, 0, b.SizeOf("zend_class_entry"))
	dir_class_entry.SetName(zend.ZendStringInitInterned("Directory", b.SizeOf("\"Directory\"")-1, 1))
	dir_class_entry.SetBuiltinFunctions(PhpDirClassFunctions)
	DirClassEntryPtr = zend.ZendRegisterInternalClass(&dir_class_entry)
	dirsep_str[0] = zend.DEFAULT_SLASH
	dirsep_str[1] = '0'
	zend.REGISTER_STRING_CONSTANT("DIRECTORY_SEPARATOR", dirsep_str, zend.CONST_CS|zend.CONST_PERSISTENT)
	pathsep_str[0] = zend.ZEND_PATHS_SEPARATOR
	pathsep_str[1] = '0'
	zend.REGISTER_STRING_CONSTANT("PATH_SEPARATOR", pathsep_str, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SCANDIR_SORT_ASCENDING", PHP_SCANDIR_SORT_ASCENDING, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SCANDIR_SORT_DESCENDING", PHP_SCANDIR_SORT_DESCENDING, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SCANDIR_SORT_NONE", PHP_SCANDIR_SORT_NONE, zend.CONST_CS|zend.CONST_PERSISTENT)
	const GLOB_BRACE = 0
	const GLOB_MARK = 0
	const GLOB_NOSORT = 0
	const GLOB_NOCHECK = 0
	const GLOB_NOESCAPE = 0
	const GLOB_ERR = 0
	const GLOB_ONLYDIR zend.ZendLong = 1 << 30

	// #define GLOB_EMULATE_ONLYDIR

	const GLOB_FLAGMASK = ^GLOB_ONLYDIR

	/* This is used for checking validity of passed flags (passing invalid flags causes segfault in glob()!! */

	const GLOB_AVAILABLE_FLAGS zend.ZendLong = 0 | GLOB_BRACE | GLOB_MARK | GLOB_NOSORT | GLOB_NOCHECK | GLOB_NOESCAPE | GLOB_ERR | GLOB_ONLYDIR
	zend.REGISTER_LONG_CONSTANT("GLOB_ONLYDIR", GLOB_ONLYDIR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("GLOB_AVAILABLE_FLAGS", GLOB_AVAILABLE_FLAGS, zend.CONST_CS|zend.CONST_PERSISTENT)
	return zend.SUCCESS
}
func _phpDoOpendir(execute_data *zend.ZendExecuteData, return_value *zend.Zval, createobject int) {
	var dirname *byte
	var dir_len int
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext = nil
	var dirp *core.PhpStream
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
			if zend.ZendParseArgPath(_arg, &dirname, &dir_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
			return
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	dirp = core.PhpStreamOpendir(dirname, core.REPORT_ERRORS, context)
	if dirp == nil {
		zend.RETVAL_FALSE
		return
	}
	dirp.AddFlags(core.PHP_STREAM_FLAG_NO_FCLOSE)
	PhpSetDefaultDir(dirp.GetRes())
	if createobject != 0 {
		zend.ObjectInitEx(return_value, DirClassEntryPtr)
		zend.AddPropertyStringl(return_value, "path", dirname, dir_len)
		zend.AddPropertyResource(return_value, "handle", dirp.GetRes())
		core.PhpStreamAutoCleanup(dirp)
	} else {
		core.PhpStreamToZval(dirp, return_value)
	}
}
func ZifOpendir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	_phpDoOpendir(execute_data, return_value, 0)
}
func ZifGetdir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	_phpDoOpendir(execute_data, return_value, 1)
}
func ZifClosedir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var id *zend.Zval = nil
	var tmp *zend.Zval
	var myself *zend.Zval
	var dirp *core.PhpStream
	var res *zend.ZendResource
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &id, 0) == 0 {
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
			return
		}
		break
	}
	if zend.ZEND_NUM_ARGS() == 0 {
		myself = zend.getThis()
		if myself != nil {
			if b.Assign(&tmp, zend.Z_OBJPROP_P(myself).KeyFind("handle")) == nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to find my handle property")
				zend.RETVAL_FALSE
				return
			}
			if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResourceEx(tmp, "Directory", streams.PhpFileLeStream()))) == nil {
				zend.RETVAL_FALSE
				return
			}
		} else {
			if !(DIRG(default_dir)) || b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(DIRG(default_dir), "Directory", streams.PhpFileLeStream()))) == nil {
				zend.RETVAL_FALSE
				return
			}
		}
	} else {
		if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(id.GetRes(), "Directory", streams.PhpFileLeStream()))) == nil {
			zend.RETVAL_FALSE
			return
		}
	}
	if !dirp.HasFlags(core.PHP_STREAM_FLAG_IS_DIR) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%d is not a valid Directory resource", dirp.GetRes().GetHandle())
		zend.RETVAL_FALSE
		return
	}
	res = dirp.GetRes()
	zend.ZendListClose(dirp.GetRes())
	if res == DIRG(default_dir) {
		PhpSetDefaultDir(nil)
	}
}
func ZifChroot(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	var ret int
	var str_len int
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
			if zend.ZendParseArgPath(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
	ret = chroot(str)
	if ret != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s (errno %d)", strerror(errno), errno)
		zend.RETVAL_FALSE
		return
	}
	PhpClearStatCache(1, nil, 0)
	ret = chdir("/")
	if ret != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s (errno %d)", strerror(errno), errno)
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifChdir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	var ret int
	var str_len int
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
			if zend.ZendParseArgPath(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
	if core.PhpCheckOpenBasedir(str) != 0 {
		zend.RETVAL_FALSE
		return
	}
	ret = zend.VCWD_CHDIR(str)
	if ret != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s (errno %d)", strerror(errno), errno)
		zend.RETVAL_FALSE
		return
	}
	if BG(CurrentStatFile) && !(zend.IS_ABSOLUTE_PATH(BG(CurrentStatFile), strlen(BG(CurrentStatFile)))) {
		zend.Efree(BG(CurrentStatFile))
		BG(CurrentStatFile) = nil
	}
	if BG(CurrentLStatFile) && !(zend.IS_ABSOLUTE_PATH(BG(CurrentLStatFile), strlen(BG(CurrentLStatFile)))) {
		zend.Efree(BG(CurrentLStatFile))
		BG(CurrentLStatFile) = nil
	}
	zend.RETVAL_TRUE
	return
}
func ZifGetcwd(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var path []byte
	var ret *byte = nil
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	ret = zend.VCWD_GETCWD(path, core.MAXPATHLEN)
	if ret != nil {
		zend.RETVAL_STRING(path)
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifRewinddir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var id *zend.Zval = nil
	var tmp *zend.Zval
	var myself *zend.Zval
	var dirp *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &id, 0) == 0 {
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
			return
		}
		break
	}
	if zend.ZEND_NUM_ARGS() == 0 {
		myself = zend.getThis()
		if myself != nil {
			if b.Assign(&tmp, zend.Z_OBJPROP_P(myself).KeyFind("handle")) == nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to find my handle property")
				zend.RETVAL_FALSE
				return
			}
			if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResourceEx(tmp, "Directory", streams.PhpFileLeStream()))) == nil {
				zend.RETVAL_FALSE
				return
			}
		} else {
			if !(DIRG(default_dir)) || b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(DIRG(default_dir), "Directory", streams.PhpFileLeStream()))) == nil {
				zend.RETVAL_FALSE
				return
			}
		}
	} else {
		if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(id.GetRes(), "Directory", streams.PhpFileLeStream()))) == nil {
			zend.RETVAL_FALSE
			return
		}
	}
	if !dirp.HasFlags(core.PHP_STREAM_FLAG_IS_DIR) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%d is not a valid Directory resource", dirp.GetRes().GetHandle())
		zend.RETVAL_FALSE
		return
	}
	core.PhpStreamRewinddir(dirp)
}
func PhpIfReaddir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var id *zend.Zval = nil
	var tmp *zend.Zval
	var myself *zend.Zval
	var dirp *core.PhpStream
	var entry core.PhpStreamDirent
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &id, 0) == 0 {
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
			return
		}
		break
	}
	if zend.ZEND_NUM_ARGS() == 0 {
		myself = zend.getThis()
		if myself != nil {
			if b.Assign(&tmp, zend.Z_OBJPROP_P(myself).KeyFind("handle")) == nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to find my handle property")
				zend.RETVAL_FALSE
				return
			}
			if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResourceEx(tmp, "Directory", streams.PhpFileLeStream()))) == nil {
				zend.RETVAL_FALSE
				return
			}
		} else {
			if !(DIRG(default_dir)) || b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(DIRG(default_dir), "Directory", streams.PhpFileLeStream()))) == nil {
				zend.RETVAL_FALSE
				return
			}
		}
	} else {
		if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(id.GetRes(), "Directory", streams.PhpFileLeStream()))) == nil {
			zend.RETVAL_FALSE
			return
		}
	}
	if !dirp.HasFlags(core.PHP_STREAM_FLAG_IS_DIR) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%d is not a valid Directory resource", dirp.GetRes().GetHandle())
		zend.RETVAL_FALSE
		return
	}
	if core.PhpStreamReaddir(dirp, &entry) != nil {
		zend.RETVAL_STRINGL(entry.GetDName(), strlen(entry.GetDName()))
		return
	}
	zend.RETVAL_FALSE
	return
}
func ZifGlob(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var cwd_skip int = 0
	var pattern *byte = nil
	var pattern_len int
	var flags zend.ZendLong = 0
	var globbuf glob_t
	var n int
	var ret int
	var basedir_limit zend.ZendBool = 0
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
			if zend.ZendParseArgPath(_arg, &pattern, &pattern_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
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
	if pattern_len >= core.MAXPATHLEN {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Pattern exceeds the maximum allowed length of %d characters", core.MAXPATHLEN)
		zend.RETVAL_FALSE
		return
	}
	if (GLOB_AVAILABLE_FLAGS & flags) != flags {
		core.PhpErrorDocref(nil, zend.E_WARNING, "At least one of the passed flags is invalid or not supported on this platform")
		zend.RETVAL_FALSE
		return
	}
	memset(&globbuf, 0, b.SizeOf("glob_t"))
	globbuf.gl_offs = 0
	if 0 != b.Assign(&ret, glob(pattern, flags&streams.GLOB_FLAGMASK, nil, &globbuf)) {
		zend.RETVAL_FALSE
		return
	}

	/* now catch the FreeBSD style of "no matches" */

	if !(globbuf.gl_pathc) || !(globbuf.gl_pathv) {

		/* Paths containing '*', '?' and some other chars are
		   illegal on Windows but legit on other platforms. For
		   this reason the direct basedir check against the glob
		   query is senseless on windows. For instance while *.txt
		   is a pretty valid filename on EXT3, it's invalid on NTFS. */

		if core.PG(open_basedir) && (*core.PG)(open_basedir) {
			if core.PhpCheckOpenBasedirEx(pattern, 0) != 0 {
				zend.RETVAL_FALSE
				return
			}
		}
		zend.ArrayInit(return_value)
		return
	}
	zend.ArrayInit(return_value)
	for n = 0; n < int(globbuf.gl_pathc); n++ {
		if core.PG(open_basedir) && (*core.PG)(open_basedir) {
			if core.PhpCheckOpenBasedirEx(globbuf.gl_pathv[n], 0) != 0 {
				basedir_limit = 1
				continue
			}
		}

		/* we need to do this every time since GLOB_ONLYDIR does not guarantee that
		 * all directories will be filtered. GNU libc documentation states the
		 * following:
		 * If the information about the type of the file is easily available
		 * non-directories will be rejected but no extra work will be done to
		 * determine the information for each file. I.e., the caller must still be
		 * able to filter directories out.
		 */

		if (flags & streams.GLOB_ONLYDIR) != 0 {
			var s zend.ZendStatT
			if 0 != zend.VCWD_STAT(globbuf.gl_pathv[n], &s) {
				continue
			}
			if S_IFDIR != (s.st_mode & S_IFMT) {
				continue
			}
		}
		zend.AddNextIndexString(return_value, globbuf.gl_pathv[n]+cwd_skip)
	}
	globfree(&globbuf)
	if basedir_limit != 0 && !(zend.Z_ARRVAL_P(return_value).GetNNumOfElements()) {
		return_value.GetArr().DestroyEx()
		zend.RETVAL_FALSE
		return
	}
}
func ZifScandir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var dirn *byte
	var dirn_len int
	var flags zend.ZendLong = 0
	var namelist **zend.ZendString
	var n int
	var i int
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
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
			if zend.ZendParseArgPath(_arg, &dirn, &dirn_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
			return
		}
		break
	}
	if dirn_len < 1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Directory name cannot be empty")
		zend.RETVAL_FALSE
		return
	}
	if zcontext != nil {
		context = streams.PhpStreamContextFromZval(zcontext, 0)
	}
	if flags == PHP_SCANDIR_SORT_ASCENDING {
		n = core.PhpStreamScandir(dirn, &namelist, context, any(streams.PhpStreamDirentAlphasort))
	} else if flags == PHP_SCANDIR_SORT_NONE {
		n = core.PhpStreamScandir(dirn, &namelist, context, nil)
	} else {
		n = core.PhpStreamScandir(dirn, &namelist, context, any(streams.PhpStreamDirentAlphasortr))
	}
	if n < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "(errno %d): %s", errno, strerror(errno))
		zend.RETVAL_FALSE
		return
	}
	zend.ArrayInit(return_value)
	for i = 0; i < n; i++ {
		zend.AddNextIndexStr(return_value, namelist[i])
	}
	if n != 0 {
		zend.Efree(namelist)
	}
}
