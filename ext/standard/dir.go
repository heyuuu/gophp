// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/dir.c>

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
   | Author: Thies C. Arntzen <thies@thieso.net>                          |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "fopen_wrappers.h"

// # include "file.h"

// # include "php_dir.h"

// # include "php_string.h"

// # include "php_scandir.h"

// # include "basic_functions.h"

// # include < unistd . h >

// # include < errno . h >

// # include < glob . h >

// @type PhpDirGlobals struct

// #define DIRG(v) ( dir_globals . v )

var DirGlobals PhpDirGlobals
var DirClassEntryPtr *zend.ZendClassEntry

// #define FETCH_DIRP() ZEND_PARSE_PARAMETERS_START ( 0 , 1 ) Z_PARAM_OPTIONAL Z_PARAM_RESOURCE ( id ) ZEND_PARSE_PARAMETERS_END ( ) ; if ( ZEND_NUM_ARGS ( ) == 0 ) { myself = getThis ( ) ; if ( myself ) { if ( ( tmp = zend_hash_str_find ( Z_OBJPROP_P ( myself ) , "handle" , sizeof ( "handle" ) - 1 ) ) == NULL ) { php_error_docref ( NULL , E_WARNING , "Unable to find my handle property" ) ; RETURN_FALSE ; } if ( ( dirp = ( php_stream * ) zend_fetch_resource_ex ( tmp , "Directory" , php_file_le_stream ( ) ) ) == NULL ) { RETURN_FALSE ; } } else { if ( ! DIRG ( default_dir ) || ( dirp = ( php_stream * ) zend_fetch_resource ( DIRG ( default_dir ) , "Directory" , php_file_le_stream ( ) ) ) == NULL ) { RETURN_FALSE ; } } } else { if ( ( dirp = ( php_stream * ) zend_fetch_resource ( Z_RES_P ( id ) , "Directory" , php_file_le_stream ( ) ) ) == NULL ) { RETURN_FALSE ; } }

/* {{{ arginfo */

/* }}} */

var PhpDirClassFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"close",
		ZifClosedir,
		ArginfoDir,
		uint32(g.SizeOf("arginfo_dir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rewind",
		ZifRewinddir,
		ArginfoDir,
		uint32(g.SizeOf("arginfo_dir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"read",
		PhpIfReaddir,
		ArginfoDir,
		uint32(g.SizeOf("arginfo_dir")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

func PhpSetDefaultDir(res *zend.ZendResource) {
	if DirGlobals.GetDefaultDir() != nil {
		zend.ZendListDelete(DirGlobals.GetDefaultDir())
	}
	if res != nil {
		zend.ZendGcAddref(&res.gc)
	}
	DirGlobals.SetDefaultDir(res)
}
func ZmActivateDir(type_ int, module_number int) int {
	DirGlobals.SetDefaultDir(nil)
	return zend.SUCCESS
}
func ZmStartupDir(type_ int, module_number int) int {
	var dirsep_str []byte
	var pathsep_str []byte
	var dir_class_entry zend.ZendClassEntry
	memset(&dir_class_entry, 0, g.SizeOf("zend_class_entry"))
	dir_class_entry.name = zend.ZendStringInitInterned("Directory", g.SizeOf("\"Directory\"")-1, 1)
	dir_class_entry.info.internal.builtin_functions = PhpDirClassFunctions
	DirClassEntryPtr = zend.ZendRegisterInternalClass(&dir_class_entry)
	dirsep_str[0] = '/'
	dirsep_str[1] = '0'
	zend.ZendRegisterStringConstant("DIRECTORY_SEPARATOR", g.SizeOf("\"DIRECTORY_SEPARATOR\"")-1, dirsep_str, 1<<0|1<<1, module_number)
	pathsep_str[0] = ':'
	pathsep_str[1] = '0'
	zend.ZendRegisterStringConstant("PATH_SEPARATOR", g.SizeOf("\"PATH_SEPARATOR\"")-1, pathsep_str, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SCANDIR_SORT_ASCENDING", g.SizeOf("\"SCANDIR_SORT_ASCENDING\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SCANDIR_SORT_DESCENDING", g.SizeOf("\"SCANDIR_SORT_DESCENDING\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SCANDIR_SORT_NONE", g.SizeOf("\"SCANDIR_SORT_NONE\"")-1, 2, 1<<0|1<<1, module_number)

	// #define GLOB_BRACE       0

	// #define GLOB_MARK       0

	// #define GLOB_NOSORT       0

	// #define GLOB_NOCHECK       0

	// #define GLOB_NOESCAPE       0

	// #define GLOB_ERR       0

	// #define GLOB_ONLYDIR       ( 1 << 30 )

	// #define GLOB_EMULATE_ONLYDIR

	// #define GLOB_FLAGMASK       ( ~ GLOB_ONLYDIR )

	/* This is used for checking validity of passed flags (passing invalid flags causes segfault in glob()!! */

	// #define GLOB_AVAILABLE_FLAGS       ( 0 | GLOB_BRACE | GLOB_MARK | GLOB_NOSORT | GLOB_NOCHECK | GLOB_NOESCAPE | GLOB_ERR | GLOB_ONLYDIR )

	zend.ZendRegisterLongConstant("GLOB_ONLYDIR", g.SizeOf("\"GLOB_ONLYDIR\"")-1, 1<<30, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("GLOB_AVAILABLE_FLAGS", g.SizeOf("\"GLOB_AVAILABLE_FLAGS\"")-1, 0|0|0|0|0|0|0|1<<30, 1<<0|1<<1, module_number)
	return zend.SUCCESS
}

/* }}} */

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

			if zend.ZendParseArgPath(_arg, &dirname, &dir_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
			return
		}
		break
	}
	if g.CondF2(g.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", PhpLeStreamContext()) }, 0), nil, func() *core.PhpStreamContext { return FileGlobals.GetDefaultContext() }) {
		context = FileGlobals.GetDefaultContext()
	} else {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
		context = FileGlobals.GetDefaultContext()
	}
	dirp = streams._phpStreamOpendir(dirname, 0x8, context)
	if dirp == nil {
		return_value.u1.type_info = 2
		return
	}
	dirp.flags |= 0x80
	PhpSetDefaultDir(dirp.res)
	if createobject != 0 {
		zend.ObjectInitEx(return_value, DirClassEntryPtr)
		zend.AddPropertyStringlEx(return_value, "path", strlen("path"), dirname, dir_len)
		zend.AddPropertyResourceEx(return_value, "handle", strlen("handle"), dirp.res)
		dirp.__exposed = 1
	} else {
		var __z *zend.Zval = return_value
		__z.value.res = dirp.res
		__z.u1.type_info = 9 | 1<<0<<8
		dirp.__exposed = 1
	}
}

/* }}} */

func ZifOpendir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	_phpDoOpendir(execute_data, return_value, 0)
}

/* }}} */

func ZifGetdir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	_phpDoOpendir(execute_data, return_value, 1)
}

/* }}} */

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

			if zend.ZendParseArgResource(_arg, &id, 0) == 0 {
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
			return
		}
		break
	}
	if execute_data.This.u2.num_args == 0 {
		if &(execute_data.This).u1.v.type_ == 8 {
			myself = &(execute_data.This)
		} else {
			myself = nil
		}
		if myself != nil {
			if g.Assign(&tmp, zend.ZendHashStrFind(myself.value.obj.handlers.get_properties(&(*myself)), "handle", g.SizeOf("\"handle\"")-1)) == nil {
				core.PhpErrorDocref(nil, 1<<1, "Unable to find my handle property")
				return_value.u1.type_info = 2
				return
			}
			if g.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResourceEx(tmp, "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.u1.type_info = 2
				return
			}
		} else {
			if DirGlobals.GetDefaultDir() == nil || g.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(DirGlobals.GetDefaultDir(), "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.u1.type_info = 2
				return
			}
		}
	} else {
		if g.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(id.value.res, "Directory", streams.PhpFileLeStream()))) == nil {
			return_value.u1.type_info = 2
			return
		}
	}
	if (dirp.flags & 0x40) == 0 {
		core.PhpErrorDocref(nil, 1<<1, "%d is not a valid Directory resource", dirp.res.handle)
		return_value.u1.type_info = 2
		return
	}
	res = dirp.res
	zend.ZendListClose(dirp.res)
	if res == DirGlobals.GetDefaultDir() {
		PhpSetDefaultDir(nil)
	}
}

/* }}} */

/* {{{ proto bool chroot(string directory)
   Change root directory */

func ZifChroot(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	var ret int
	var str_len int
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

			if zend.ZendParseArgPath(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
	ret = chroot(str)
	if ret != 0 {
		core.PhpErrorDocref(nil, 1<<1, "%s (errno %d)", strerror(errno), errno)
		return_value.u1.type_info = 2
		return
	}
	PhpClearStatCache(1, nil, 0)
	ret = chdir("/")
	if ret != 0 {
		core.PhpErrorDocref(nil, 1<<1, "%s (errno %d)", strerror(errno), errno)
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

/* {{{ proto bool chdir(string directory)
   Change the current directory */

func ZifChdir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	var ret int
	var str_len int
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

			if zend.ZendParseArgPath(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
	if core.PhpCheckOpenBasedir(str) != 0 {
		return_value.u1.type_info = 2
		return
	}
	ret = chdir(str)
	if ret != 0 {
		core.PhpErrorDocref(nil, 1<<1, "%s (errno %d)", strerror(errno), errno)
		return_value.u1.type_info = 2
		return
	}
	if BasicGlobals.GetCurrentStatFile() != nil && BasicGlobals.GetCurrentStatFile()[0] != '/' {
		zend._efree(BasicGlobals.GetCurrentStatFile())
		BasicGlobals.SetCurrentStatFile(nil)
	}
	if BasicGlobals.GetCurrentLStatFile() != nil && BasicGlobals.GetCurrentLStatFile()[0] != '/' {
		zend._efree(BasicGlobals.GetCurrentLStatFile())
		BasicGlobals.SetCurrentLStatFile(nil)
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifGetcwd(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var path []byte
	var ret *byte = nil
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	ret = getcwd(path, 256)
	if ret != nil {
		var _s *byte = path
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifRewinddir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var id *zend.Zval = nil
	var tmp *zend.Zval
	var myself *zend.Zval
	var dirp *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 0
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

			if zend.ZendParseArgResource(_arg, &id, 0) == 0 {
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
			return
		}
		break
	}
	if execute_data.This.u2.num_args == 0 {
		if &(execute_data.This).u1.v.type_ == 8 {
			myself = &(execute_data.This)
		} else {
			myself = nil
		}
		if myself != nil {
			if g.Assign(&tmp, zend.ZendHashStrFind(myself.value.obj.handlers.get_properties(&(*myself)), "handle", g.SizeOf("\"handle\"")-1)) == nil {
				core.PhpErrorDocref(nil, 1<<1, "Unable to find my handle property")
				return_value.u1.type_info = 2
				return
			}
			if g.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResourceEx(tmp, "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.u1.type_info = 2
				return
			}
		} else {
			if DirGlobals.GetDefaultDir() == nil || g.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(DirGlobals.GetDefaultDir(), "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.u1.type_info = 2
				return
			}
		}
	} else {
		if g.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(id.value.res, "Directory", streams.PhpFileLeStream()))) == nil {
			return_value.u1.type_info = 2
			return
		}
	}
	if (dirp.flags & 0x40) == 0 {
		core.PhpErrorDocref(nil, 1<<1, "%d is not a valid Directory resource", dirp.res.handle)
		return_value.u1.type_info = 2
		return
	}
	streams._phpStreamSeek(dirp, 0, 0)
}

/* }}} */

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

			if zend.ZendParseArgResource(_arg, &id, 0) == 0 {
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
			return
		}
		break
	}
	if execute_data.This.u2.num_args == 0 {
		if &(execute_data.This).u1.v.type_ == 8 {
			myself = &(execute_data.This)
		} else {
			myself = nil
		}
		if myself != nil {
			if g.Assign(&tmp, zend.ZendHashStrFind(myself.value.obj.handlers.get_properties(&(*myself)), "handle", g.SizeOf("\"handle\"")-1)) == nil {
				core.PhpErrorDocref(nil, 1<<1, "Unable to find my handle property")
				return_value.u1.type_info = 2
				return
			}
			if g.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResourceEx(tmp, "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.u1.type_info = 2
				return
			}
		} else {
			if DirGlobals.GetDefaultDir() == nil || g.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(DirGlobals.GetDefaultDir(), "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.u1.type_info = 2
				return
			}
		}
	} else {
		if g.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(id.value.res, "Directory", streams.PhpFileLeStream()))) == nil {
			return_value.u1.type_info = 2
			return
		}
	}
	if (dirp.flags & 0x40) == 0 {
		core.PhpErrorDocref(nil, 1<<1, "%d is not a valid Directory resource", dirp.res.handle)
		return_value.u1.type_info = 2
		return
	}
	if streams._phpStreamReaddir(dirp, &entry) != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(entry.d_name, strlen(entry.d_name), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

/* {{{ proto array glob(string pattern [, int flags])
   Find pathnames matching a pattern */

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

			if zend.ZendParseArgPath(_arg, &pattern, &pattern_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
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
			return
		}
		break
	}
	if pattern_len >= 256 {
		core.PhpErrorDocref(nil, 1<<1, "Pattern exceeds the maximum allowed length of %d characters", 256)
		return_value.u1.type_info = 2
		return
	}
	if ((0 | 0 | 0 | 0 | 0 | 0 | 0 | 1<<30) & flags) != flags {
		core.PhpErrorDocref(nil, 1<<1, "At least one of the passed flags is invalid or not supported on this platform")
		return_value.u1.type_info = 2
		return
	}
	memset(&globbuf, 0, g.SizeOf("glob_t"))
	globbuf.gl_offs = 0
	if 0 != g.Assign(&ret, glob(pattern, flags & ^(1<<30), nil, &globbuf)) {
		return_value.u1.type_info = 2
		return
	}

	/* now catch the FreeBSD style of "no matches" */

	if !(globbuf.gl_pathc) || !(globbuf.gl_pathv) {

		/* Paths containing '*', '?' and some other chars are
		   illegal on Windows but legit on other platforms. For
		   this reason the direct basedir check against the glob
		   query is senseless on windows. For instance while *.txt
		   is a pretty valid filename on EXT3, it's invalid on NTFS. */

		if core.CoreGlobals.open_basedir != nil && (*(core.CoreGlobals.open_basedir)) {
			if core.PhpCheckOpenBasedirEx(pattern, 0) != 0 {
				return_value.u1.type_info = 2
				return
			}
		}
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for n = 0; n < int(globbuf.gl_pathc); n++ {
		if core.CoreGlobals.open_basedir != nil && (*(core.CoreGlobals.open_basedir)) {
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

		if (flags & 1 << 30) != 0 {
			var s zend.ZendStatT
			if 0 != stat(globbuf.gl_pathv[n], &s) {
				continue
			}
			if S_IFDIR != (s.st_mode & S_IFMT) {
				continue
			}
		}
		zend.AddNextIndexString(return_value, globbuf.gl_pathv[n]+cwd_skip)
	}
	globfree(&globbuf)
	if basedir_limit != 0 && return_value.value.arr.nNumOfElements == 0 {
		zend.ZendArrayDestroy(return_value.value.arr)
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

/* {{{ proto array scandir(string dir [, int sorting_order [, resource context]])
   List files & directories inside the specified path */

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

			if zend.ZendParseArgPath(_arg, &dirn, &dirn_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
			return
		}
		break
	}
	if dirn_len < 1 {
		core.PhpErrorDocref(nil, 1<<1, "Directory name cannot be empty")
		return_value.u1.type_info = 2
		return
	}
	if zcontext != nil {
		if g.CondF2(g.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", PhpLeStreamContext()) }, 0), nil, func() *core.PhpStreamContext { return FileGlobals.GetDefaultContext() }) {
			context = FileGlobals.GetDefaultContext()
		} else {
			FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
			context = FileGlobals.GetDefaultContext()
		}
	}
	if flags == 0 {
		n = streams._phpStreamScandir(dirn, &namelist, 0, context, any(streams.PhpStreamDirentAlphasort))
	} else if flags == 2 {
		n = streams._phpStreamScandir(dirn, &namelist, 0, context, nil)
	} else {
		n = streams._phpStreamScandir(dirn, &namelist, 0, context, any(streams.PhpStreamDirentAlphasortr))
	}
	if n < 0 {
		core.PhpErrorDocref(nil, 1<<1, "(errno %d): %s", errno, strerror(errno))
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for i = 0; i < n; i++ {
		zend.AddNextIndexStr(return_value, namelist[i])
	}
	if n != 0 {
		zend._efree(namelist)
	}
}

/* }}} */
