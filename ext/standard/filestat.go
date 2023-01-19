// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/filestat.c>

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
   | Author:  Jim Winstead <jimw@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "fopen_wrappers.h"

// # include "php_globals.h"

// # include < stdlib . h >

// # include < sys / stat . h >

// # include < string . h >

// # include < errno . h >

// # include < ctype . h >

// # include < time . h >

// # include < unistd . h >

// # include < sys / param . h >

// # include < sys / statvfs . h >

// # include < pwd . h >

// # include < grp . h >

// # include < utime . h >

// # include "basic_functions.h"

// # include "php_filestat.h"

func ZmActivateFilestat(type_ int, module_number int) int {
	BasicGlobals.SetCurrentStatFile(nil)
	BasicGlobals.SetCurrentLStatFile(nil)
	return zend.SUCCESS
}

/* }}} */

func ZmDeactivateFilestat(type_ int, module_number int) int {
	if BasicGlobals.GetCurrentStatFile() != nil {
		zend._efree(BasicGlobals.GetCurrentStatFile())
		BasicGlobals.SetCurrentStatFile(nil)
	}
	if BasicGlobals.GetCurrentLStatFile() != nil {
		zend._efree(BasicGlobals.GetCurrentLStatFile())
		BasicGlobals.SetCurrentLStatFile(nil)
	}
	return zend.SUCCESS
}

/* }}} */

func PhpDiskTotalSpace(path *byte, space *float64) int {
	var bytestotal float64 = 0
	var buf __struct__statvfs
	if statvfs(path, &buf) {
		core.PhpErrorDocref(nil, 1<<1, "%s", strerror(errno))
		return zend.FAILURE
	}
	if buf.f_frsize {
		bytestotal = float64(buf.f_blocks) * float64(buf.f_frsize)
	} else {
		bytestotal = float64(buf.f_blocks) * float64(buf.f_bsize)
	}
	*space = bytestotal
	return zend.SUCCESS
}

/* }}} */

func ZifDiskTotalSpace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var bytestotal float64
	var path *byte
	var fullpath []*byte
	var path_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &path, &path_len, 0) == 0 {
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
			return
		}
		break
	}
	if core.ExpandFilepath(path, fullpath) == nil {
		return_value.u1.type_info = 2
		return
	}
	if core.PhpCheckOpenBasedir(fullpath) != 0 {
		return_value.u1.type_info = 2
		return
	}
	if PhpDiskTotalSpace(fullpath, &bytestotal) == zend.SUCCESS {
		var __z *zend.Zval = return_value
		__z.value.dval = bytestotal
		__z.u1.type_info = 5
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func PhpDiskFreeSpace(path *byte, space *float64) int {
	var bytesfree float64 = 0
	var buf __struct__statvfs
	if statvfs(path, &buf) {
		core.PhpErrorDocref(nil, 1<<1, "%s", strerror(errno))
		return zend.FAILURE
	}
	if buf.f_frsize {
		bytesfree = float64(buf.f_bavail) * float64(buf.f_frsize)
	} else {
		bytesfree = float64(buf.f_bavail) * float64(buf.f_bsize)
	}
	*space = bytesfree
	return zend.SUCCESS
}

/* }}} */

func ZifDiskFreeSpace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var bytesfree float64
	var path *byte
	var fullpath []*byte
	var path_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &path, &path_len, 0) == 0 {
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
			return
		}
		break
	}
	if core.ExpandFilepath(path, fullpath) == nil {
		return_value.u1.type_info = 2
		return
	}
	if core.PhpCheckOpenBasedir(fullpath) != 0 {
		return_value.u1.type_info = 2
		return
	}
	if PhpDiskFreeSpace(fullpath, &bytesfree) == zend.SUCCESS {
		var __z *zend.Zval = return_value
		__z.value.dval = bytesfree
		__z.u1.type_info = 5
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func PhpGetGidByName(name *byte, gid *gid_t) int {
	var gr *__struct__group = getgrnam(name)
	if gr == nil {
		return zend.FAILURE
	}
	*gid = gr.gr_gid
	return zend.SUCCESS
}
func PhpDoChgrp(execute_data *zend.ZendExecuteData, return_value *zend.Zval, do_lchgrp int) {
	var filename *byte
	var filename_len int
	var group *zend.Zval
	var gid gid_t
	var ret int
	var wrapper *core.PhpStreamWrapper
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &group, 0)
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
	wrapper = streams.PhpStreamLocateUrlWrapper(filename, nil, 0)
	if wrapper != &PhpPlainFilesWrapper || strncasecmp("file://", filename, 7) == 0 {
		if wrapper != nil && wrapper.wops.stream_metadata != nil {
			var option int
			var value any
			if group.u1.v.type_ == 4 {
				option = 5
				value = &(*group).value.lval
			} else if group.u1.v.type_ == 6 {
				option = 4
				value = group.value.str.val
			} else {
				core.PhpErrorDocref(nil, 1<<1, "parameter 2 should be string or int, %s given", zend.ZendZvalTypeName(group))
				return_value.u1.type_info = 2
				return
			}
			if wrapper.wops.stream_metadata(wrapper, filename, option, value, nil) != 0 {
				return_value.u1.type_info = 3
				return
			} else {
				return_value.u1.type_info = 2
				return
			}
		} else {

			/* On Windows, we expect regular chgrp to fail silently by default */

			core.PhpErrorDocref(nil, 1<<1, "Can not call chgrp() for a non-standard stream")
			return_value.u1.type_info = 2
			return
		}
	}
	if group.u1.v.type_ == 4 {
		gid = gid_t(*group).value.lval
	} else if group.u1.v.type_ == 6 {
		if PhpGetGidByName(group.value.str.val, &gid) != zend.SUCCESS {
			core.PhpErrorDocref(nil, 1<<1, "Unable to find gid for %s", group.value.str.val)
			return_value.u1.type_info = 2
			return
		}
	} else {
		core.PhpErrorDocref(nil, 1<<1, "parameter 2 should be string or int, %s given", zend.ZendZvalTypeName(group))
		return_value.u1.type_info = 2
		return
	}

	/* Check the basedir */

	if core.PhpCheckOpenBasedir(filename) != 0 {
		return_value.u1.type_info = 2
		return
	}
	if do_lchgrp != 0 {
		ret = lchown(filename, -1, gid)
	} else {
		ret = chown(filename, -1, gid)
	}
	if ret == -1 {
		core.PhpErrorDocref(nil, 1<<1, "%s", strerror(errno))
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifChgrp(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpDoChgrp(execute_data, return_value, 0)
}

/* }}} */

func ZifLchgrp(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpDoChgrp(execute_data, return_value, 1)
}

/* }}} */

func PhpGetUidByName(name *byte, uid *uid_t) uid_t {
	var pw *__struct__passwd = getpwnam(name)
	if pw == nil {
		return zend.FAILURE
	}
	*uid = pw.pw_uid
	return zend.SUCCESS
}
func PhpDoChown(execute_data *zend.ZendExecuteData, return_value *zend.Zval, do_lchown int) {
	var filename *byte
	var filename_len int
	var user *zend.Zval
	var uid uid_t
	var ret int
	var wrapper *core.PhpStreamWrapper
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &user, 0)
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
	wrapper = streams.PhpStreamLocateUrlWrapper(filename, nil, 0)
	if wrapper != &PhpPlainFilesWrapper || strncasecmp("file://", filename, 7) == 0 {
		if wrapper != nil && wrapper.wops.stream_metadata != nil {
			var option int
			var value any
			if user.u1.v.type_ == 4 {
				option = 3
				value = &(*user).value.lval
			} else if user.u1.v.type_ == 6 {
				option = 2
				value = user.value.str.val
			} else {
				core.PhpErrorDocref(nil, 1<<1, "parameter 2 should be string or int, %s given", zend.ZendZvalTypeName(user))
				return_value.u1.type_info = 2
				return
			}
			if wrapper.wops.stream_metadata(wrapper, filename, option, value, nil) != 0 {
				return_value.u1.type_info = 3
				return
			} else {
				return_value.u1.type_info = 2
				return
			}
		} else {

			/* On Windows, we expect regular chown to fail silently by default */

			core.PhpErrorDocref(nil, 1<<1, "Can not call chown() for a non-standard stream")
			return_value.u1.type_info = 2
			return
		}
	}
	if user.u1.v.type_ == 4 {
		uid = uid_t(*user).value.lval
	} else if user.u1.v.type_ == 6 {
		if PhpGetUidByName(user.value.str.val, &uid) != zend.SUCCESS {
			core.PhpErrorDocref(nil, 1<<1, "Unable to find uid for %s", user.value.str.val)
			return_value.u1.type_info = 2
			return
		}
	} else {
		core.PhpErrorDocref(nil, 1<<1, "parameter 2 should be string or int, %s given", zend.ZendZvalTypeName(user))
		return_value.u1.type_info = 2
		return
	}

	/* Check the basedir */

	if core.PhpCheckOpenBasedir(filename) != 0 {
		return_value.u1.type_info = 2
		return
	}
	if do_lchown != 0 {
		ret = lchown(filename, uid, -1)
	} else {
		ret = chown(filename, uid, -1)
	}
	if ret == -1 {
		core.PhpErrorDocref(nil, 1<<1, "%s", strerror(errno))
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifChown(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpDoChown(execute_data, return_value, 0)
}

/* }}} */

func ZifLchown(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	return_value.u1.type_info = 3
	PhpDoChown(execute_data, return_value, 1)
}

/* }}} */

func ZifChmod(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var mode zend.ZendLong
	var ret int
	var imode mode_t
	var wrapper *core.PhpStreamWrapper
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &mode, &_dummy, 0, 0) == 0 {
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
	wrapper = streams.PhpStreamLocateUrlWrapper(filename, nil, 0)
	if wrapper != &PhpPlainFilesWrapper || strncasecmp("file://", filename, 7) == 0 {
		if wrapper != nil && wrapper.wops.stream_metadata != nil {
			if wrapper.wops.stream_metadata(wrapper, filename, 6, &mode, nil) != 0 {
				return_value.u1.type_info = 3
				return
			} else {
				return_value.u1.type_info = 2
				return
			}
		} else {
			core.PhpErrorDocref(nil, 1<<1, "Can not call chmod() for a non-standard stream")
			return_value.u1.type_info = 2
			return
		}
	}

	/* Check the basedir */

	if core.PhpCheckOpenBasedir(filename) != 0 {
		return_value.u1.type_info = 2
		return
	}
	imode = mode_t(mode)
	ret = chmod(filename, imode)
	if ret == -1 {
		core.PhpErrorDocref(nil, 1<<1, "%s", strerror(errno))
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

/* {{{ proto bool touch(string filename [, int time [, int atime]])
   Set modification time of file */

func ZifTouch(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var filetime zend.ZendLong = 0
	var fileatime zend.ZendLong = 0
	var ret int
	var argc int = execute_data.This.u2.num_args
	var file *FILE
	var newtimebuf __struct__utimbuf
	var newtime *__struct__utimbuf = &newtimebuf
	var wrapper *core.PhpStreamWrapper
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &filetime, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &fileatime, &_dummy, 0, 0) == 0 {
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
	if filename_len == 0 {
		return_value.u1.type_info = 2
		return
	}
	switch argc {
	case 1:
		newtime = nil
		break
	case 2:
		newtime.actime = filetime
		newtime.modtime = newtime.actime
		break
	case 3:
		newtime.modtime = filetime
		newtime.actime = fileatime
		break
	default:

		/* Never reached */

		zend.ZendWrongParamCount()
		return
	}
	wrapper = streams.PhpStreamLocateUrlWrapper(filename, nil, 0)
	if wrapper != &PhpPlainFilesWrapper || strncasecmp("file://", filename, 7) == 0 {
		if wrapper != nil && wrapper.wops.stream_metadata != nil {
			if wrapper.wops.stream_metadata(wrapper, filename, 1, newtime, nil) != 0 {
				return_value.u1.type_info = 3
				return
			} else {
				return_value.u1.type_info = 2
				return
			}
		} else {
			var stream *core.PhpStream
			if argc > 1 {
				core.PhpErrorDocref(nil, 1<<1, "Can not call touch() for a non-standard stream")
				return_value.u1.type_info = 2
				return
			}
			stream = streams._phpStreamOpenWrapperEx(filename, "c", 0x8, nil, nil)
			if stream != nil {
				streams._phpStreamFree(stream, 1|2)
				return_value.u1.type_info = 3
				return
			} else {
				return_value.u1.type_info = 2
				return
			}
		}
	}

	/* Check the basedir */

	if core.PhpCheckOpenBasedir(filename) != 0 {
		return_value.u1.type_info = 2
		return
	}

	/* create the file if it doesn't exist already */

	if access(filename, F_OK) != 0 {
		file = fopen(filename, "w")
		if file == nil {
			core.PhpErrorDocref(nil, 1<<1, "Unable to create file %s because %s", filename, strerror(errno))
			return_value.u1.type_info = 2
			return
		}
		fclose(file)
	}
	ret = utime(filename, newtime)
	if ret == -1 {
		core.PhpErrorDocref(nil, 1<<1, "Utime failed: %s", strerror(errno))
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

/* {{{ php_clear_stat_cache()
 */

func PhpClearStatCache(clear_realpath_cache zend.ZendBool, filename *byte, filename_len int) {
	/* always clear CurrentStatFile and CurrentLStatFile even if filename is not NULL
	 * as it may contain outdated data (e.g. "nlink" for a directory when deleting a file
	 * in this directory, as shown by lstat_stat_variation9.phpt) */

	if BasicGlobals.GetCurrentStatFile() != nil {
		zend._efree(BasicGlobals.GetCurrentStatFile())
		BasicGlobals.SetCurrentStatFile(nil)
	}
	if BasicGlobals.GetCurrentLStatFile() != nil {
		zend._efree(BasicGlobals.GetCurrentLStatFile())
		BasicGlobals.SetCurrentLStatFile(nil)
	}
	if clear_realpath_cache != 0 {
		if filename != nil {
			zend.RealpathCacheDel(filename, filename_len)
		} else {
			zend.RealpathCacheClean()
		}
	}
}

/* }}} */

func ZifClearstatcache(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var clear_realpath_cache zend.ZendBool = 0
	var filename *byte = nil
	var filename_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &clear_realpath_cache, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpClearStatCache(clear_realpath_cache, filename, filename_len)
}

/* }}} */

// #define IS_LINK_OPERATION(__t) ( ( __t ) == FS_TYPE || ( __t ) == FS_IS_LINK || ( __t ) == FS_LSTAT )

// #define IS_EXISTS_CHECK(__t) ( ( __t ) == FS_EXISTS || ( __t ) == FS_IS_W || ( __t ) == FS_IS_R || ( __t ) == FS_IS_X || ( __t ) == FS_IS_FILE || ( __t ) == FS_IS_DIR || ( __t ) == FS_IS_LINK )

// #define IS_ABLE_CHECK(__t) ( ( __t ) == FS_IS_R || ( __t ) == FS_IS_W || ( __t ) == FS_IS_X )

// #define IS_ACCESS_CHECK(__t) ( IS_ABLE_CHECK ( type ) || ( __t ) == FS_EXISTS )

/* {{{ php_stat
 */

func PhpStat(filename *byte, filename_length int, type_ int, return_value *zend.Zval) {
	var stat_dev zend.Zval
	var stat_ino zend.Zval
	var stat_mode zend.Zval
	var stat_nlink zend.Zval
	var stat_uid zend.Zval
	var stat_gid zend.Zval
	var stat_rdev zend.Zval
	var stat_size zend.Zval
	var stat_atime zend.Zval
	var stat_mtime zend.Zval
	var stat_ctime zend.Zval
	var stat_blksize zend.Zval
	var stat_blocks zend.Zval
	var stat_sb *zend.ZendStatT
	var ssb core.PhpStreamStatbuf
	var flags int = 0
	var rmask int = S_IROTH
	var wmask int = S_IWOTH
	var xmask int = S_IXOTH
	var stat_sb_names []*byte = []*byte{"dev", "ino", "mode", "nlink", "uid", "gid", "rdev", "size", "atime", "mtime", "ctime", "blksize", "blocks"}
	var local *byte
	var wrapper *core.PhpStreamWrapper
	if filename_length == 0 {
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&wrapper, streams.PhpStreamLocateUrlWrapper(filename, &local, 0)) == &PhpPlainFilesWrapper && core.PhpCheckOpenBasedir(local) != 0 {
		return_value.u1.type_info = 2
		return
	}
	if type_ == 10 || type_ == 9 || type_ == 11 || type_ == 15 {
		if wrapper == &PhpPlainFilesWrapper {
			switch type_ {

			}
		}
	}
	if type_ == 8 || type_ == 14 || type_ == 16 {
		flags |= 1
	}
	if type_ == 15 || type_ == 9 || type_ == 10 || type_ == 11 || type_ == 12 || type_ == 13 || type_ == 14 {
		flags |= 2
	}
	if streams._phpStreamStatPath((*byte)(filename), flags, &ssb, nil) != 0 {

		/* Error Occurred */

		if !(type_ == 15 || type_ == 9 || type_ == 10 || type_ == 11 || type_ == 12 || type_ == 13 || type_ == 14) {
			core.PhpErrorDocref(nil, 1<<1, "%sstat failed for %s", g.Cond(type_ == 8 || type_ == 14 || type_ == 16, "L", ""), filename)
		}
		return_value.u1.type_info = 2
		return
	}
	stat_sb = &ssb.sb
	if type_ >= 9 && type_ <= 11 {
		if ssb.sb.st_uid == getuid() {
			rmask = S_IRUSR
			wmask = S_IWUSR
			xmask = S_IXUSR
		} else if ssb.sb.st_gid == getgid() {
			rmask = S_IRGRP
			wmask = S_IWGRP
			xmask = S_IXGRP
		} else {
			var groups int
			var n int
			var i int
			var gids *gid_t
			groups = getgroups(0, nil)
			if groups > 0 {
				gids = (*gid_t)(zend._safeEmalloc(groups, g.SizeOf("gid_t"), 0))
				n = getgroups(groups, gids)
				for i = 0; i < n; i++ {
					if ssb.sb.st_gid == gids[i] {
						rmask = S_IRGRP
						wmask = S_IWGRP
						xmask = S_IXGRP
						break
					}
				}
				zend._efree(gids)
			}
		}
	}
	if (type_ == 10 || type_ == 9 || type_ == 11) && getuid() == 0 {

		/* root has special perms on plain_wrapper */

		if wrapper == &PhpPlainFilesWrapper {
			if type_ == 11 {
				xmask = S_IXUSR | S_IXGRP | S_IXOTH
			} else {
				return_value.u1.type_info = 3
				return
			}
		}

		/* root has special perms on plain_wrapper */

	}
	switch type_ {
	case 0:
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZendLong(ssb.sb.st_mode)
		__z.u1.type_info = 4
		return
	case 1:
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZendLong(ssb.sb.st_ino)
		__z.u1.type_info = 4
		return
	case 2:
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZendLong(ssb.sb.st_size)
		__z.u1.type_info = 4
		return
	case 3:
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZendLong(ssb.sb.st_uid)
		__z.u1.type_info = 4
		return
	case 4:
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZendLong(ssb.sb.st_gid)
		__z.u1.type_info = 4
		return
	case 5:
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZendLong(ssb.sb.st_atime)
		__z.u1.type_info = 4
		return
	case 6:
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZendLong(ssb.sb.st_mtime)
		__z.u1.type_info = 4
		return
	case 7:
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZendLong(ssb.sb.st_ctime)
		__z.u1.type_info = 4
		return
	case 8:
		if (ssb.sb.st_mode & S_IFMT) == 0120000 {
			var _s *byte = "link"
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		}
		switch ssb.sb.st_mode & S_IFMT {
		case 010000:
			var _s *byte = "fifo"
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		case S_IFCHR:
			var _s *byte = "char"
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		case S_IFDIR:
			var _s *byte = "dir"
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		case 060000:
			var _s *byte = "block"
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		case S_IFREG:
			var _s *byte = "file"
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		}
		core.PhpErrorDocref(nil, 1<<3, "Unknown file type (%d)", ssb.sb.st_mode&S_IFMT)
		var _s *byte = "unknown"
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	case 9:
		if (ssb.sb.st_mode & wmask) != 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	case 10:
		if (ssb.sb.st_mode & rmask) != 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	case 11:
		if (ssb.sb.st_mode & xmask) != 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	case 12:
		if (ssb.sb.st_mode & S_IFMT) == S_IFREG {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	case 13:
		if (ssb.sb.st_mode & S_IFMT) == S_IFDIR {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	case 14:
		if (ssb.sb.st_mode & S_IFMT) == 0120000 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	case 15:
		return_value.u1.type_info = 3
		return
	case 16:

	case 17:
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		var __z *zval = &stat_dev
		__z.value.lval = stat_sb.st_dev
		__z.u1.type_info = 4
		var __z *zval = &stat_ino
		__z.value.lval = stat_sb.st_ino
		__z.u1.type_info = 4
		var __z *zval = &stat_mode
		__z.value.lval = stat_sb.st_mode
		__z.u1.type_info = 4
		var __z *zval = &stat_nlink
		__z.value.lval = stat_sb.st_nlink
		__z.u1.type_info = 4
		var __z *zval = &stat_uid
		__z.value.lval = stat_sb.st_uid
		__z.u1.type_info = 4
		var __z *zval = &stat_gid
		__z.value.lval = stat_sb.st_gid
		__z.u1.type_info = 4
		var __z *zval = &stat_rdev
		__z.value.lval = stat_sb.st_rdev
		__z.u1.type_info = 4
		var __z *zval = &stat_size
		__z.value.lval = stat_sb.st_size
		__z.u1.type_info = 4
		var __z *zval = &stat_atime
		__z.value.lval = stat_sb.st_atime
		__z.u1.type_info = 4
		var __z *zval = &stat_mtime
		__z.value.lval = stat_sb.st_mtime
		__z.u1.type_info = 4
		var __z *zval = &stat_ctime
		__z.value.lval = stat_sb.st_ctime
		__z.u1.type_info = 4
		var __z *zval = &stat_blksize
		__z.value.lval = stat_sb.st_blksize
		__z.u1.type_info = 4
		var __z *zend.Zval = &stat_blocks
		__z.value.lval = stat_sb.st_blocks
		__z.u1.type_info = 4

		/* Store numeric indexes in proper order */

		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_dev)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_ino)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_mode)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_nlink)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_uid)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_gid)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_rdev)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_size)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_atime)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_mtime)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_ctime)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_blksize)
		zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_blocks)

		/* Store string indexes referencing the same zval*/

		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[0], strlen(stat_sb_names[0]), &stat_dev)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[1], strlen(stat_sb_names[1]), &stat_ino)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[2], strlen(stat_sb_names[2]), &stat_mode)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[3], strlen(stat_sb_names[3]), &stat_nlink)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[4], strlen(stat_sb_names[4]), &stat_uid)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[5], strlen(stat_sb_names[5]), &stat_gid)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[6], strlen(stat_sb_names[6]), &stat_rdev)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[7], strlen(stat_sb_names[7]), &stat_size)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[8], strlen(stat_sb_names[8]), &stat_atime)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[9], strlen(stat_sb_names[9]), &stat_mtime)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[10], strlen(stat_sb_names[10]), &stat_ctime)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[11], strlen(stat_sb_names[11]), &stat_blksize)
		zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[12], strlen(stat_sb_names[12]), &stat_blocks)
		return
	}
	core.PhpErrorDocref(nil, 1<<1, "Didn't understand stat call")
	return_value.u1.type_info = 2
	return
}

/* }}} */

// #define FileFunction(name,funcnum) ZEND_NAMED_FUNCTION ( name ) { char * filename ; size_t filename_len ; ZEND_PARSE_PARAMETERS_START ( 1 , 1 ) Z_PARAM_PATH ( filename , filename_len ) ZEND_PARSE_PARAMETERS_END ( ) ; php_stat ( filename , filename_len , funcnum , return_value ) ; }

/* }}} */

func ZifFileperms(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 0, return_value)
}

/* }}} */

func ZifFileinode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 1, return_value)
}

/* }}} */

func ZifFilesize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 2, return_value)
}

/* }}} */

func ZifFileowner(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 3, return_value)
}

/* }}} */

func ZifFilegroup(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 4, return_value)
}

/* }}} */

func ZifFileatime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 5, return_value)
}

/* }}} */

func ZifFilemtime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 6, return_value)
}

/* }}} */

func ZifFilectime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 7, return_value)
}

/* }}} */

func ZifFiletype(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 8, return_value)
}

/* }}} */

func ZifIsWritable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 9, return_value)
}

/* }}} */

func ZifIsReadable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 10, return_value)
}

/* }}} */

func ZifIsExecutable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 11, return_value)
}

/* }}} */

func ZifIsFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 12, return_value)
}

/* }}} */

func ZifIsDir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 13, return_value)
}

/* }}} */

func ZifIsLink(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 14, return_value)
}

/* }}} */

func ZifFileExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 15, return_value)
}

/* }}} */

func PhpIfLstat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 16, return_value)
}

/* }}} */

func PhpIfStat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return
		}
		break
	}
	PhpStat(filename, filename_len, 17, return_value)
}

/* }}} */

func ZifRealpathCacheSize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = zend.RealpathCacheSize()
	__z.u1.type_info = 4
	return
}

/* {{{ proto bool realpath_cache_get()
   Get current size of realpath cache */

func ZifRealpathCacheGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var buckets **zend.RealpathCacheBucket = zend.RealpathCacheGetBuckets()
	var end ***zend.RealpathCacheBucket = buckets + zend.RealpathCacheMaxBuckets()
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for buckets < end {
		var bucket *zend.RealpathCacheBucket = *buckets
		for bucket != nil {
			var entry zend.Zval
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = &entry
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

			/* bucket->key is unsigned long */

			if INT64_MAX >= bucket.key {
				zend.AddAssocLongEx(&entry, "key", g.SizeOf("\"key\"")-1, bucket.key)
			} else {
				zend.AddAssocDoubleEx(&entry, "key", g.SizeOf("\"key\"")-1, float64(bucket.key))
			}
			zend.AddAssocBoolEx(&entry, "is_dir", g.SizeOf("\"is_dir\"")-1, bucket.is_dir)
			zend.AddAssocStringlEx(&entry, "realpath", g.SizeOf("\"realpath\"")-1, bucket.realpath, bucket.realpath_len)
			zend.AddAssocLongEx(&entry, "expires", g.SizeOf("\"expires\"")-1, bucket.expires)
			zend.ZendHashStrUpdate(return_value.value.arr, bucket.path, bucket.path_len, &entry)
			bucket = bucket.next
		}
		buckets++
	}
}
