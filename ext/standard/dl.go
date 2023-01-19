// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/dl.h>

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
   | Authors: Brian Schaffner <brian@tool.net>                            |
   |          Shane Caraveo <shane@caraveo.com>                           |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define DL_H

/* dynamic loading functions */

// Source: <ext/standard/dl.c>

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
   | Authors: Brian Schaffner <brian@tool.net>                            |
   |          Shane Caraveo <shane@caraveo.com>                           |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "dl.h"

// # include "php_globals.h"

// # include "php_ini.h"

// # include "ext/standard/info.h"

// # include "SAPI.h"

// # include < stdlib . h >

// # include < stdio . h >

// # include < string . h >

// # include < sys / param . h >

// #define GET_DL_ERROR() DL_ERROR ( )

/* {{{ proto int dl(string extension_filename)
   Load a PHP extension at runtime */

func ZifDl(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
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

			if zend.ZendParseArgString(_arg, &filename, &filename_len, 0) == 0 {
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
	if core.CoreGlobals.enable_dl == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Dynamically loaded extensions aren't enabled")
		return_value.u1.type_info = 2
		return
	}
	if filename_len >= 256 {
		core.PhpErrorDocref(nil, 1<<1, "File name exceeds the maximum allowed length of %d characters", 256)
		return_value.u1.type_info = 2
		return
	}
	PhpDl(filename, 2, return_value, 0)
	if return_value.u1.v.type_ == 3 {
		zend.EG.full_tables_cleanup = 1
	}
}

/* }}} */

/* {{{ php_load_shlib
 */

func PhpLoadShlib(path *byte, errp **byte) any {
	var handle any
	var err *byte
	handle = dlopen(path, 1|0)
	if !handle {
		err = dlerror()
		*errp = zend._estrdup(err)
		dlerror()
	}
	return handle
}

/* }}} */

func PhpLoadExtension(filename *byte, type_ int, start_now int) int {
	var handle any
	var libpath *byte
	var module_entry *zend.ZendModuleEntry
	var get_module func() *zend.ZendModuleEntry
	var error_type int
	var slash_suffix int = 0
	var extension_dir *byte
	var err1 *byte
	var err2 *byte
	if type_ == 1 {
		extension_dir = zend.ZendIniStringEx("extension_dir", g.SizeOf("\"extension_dir\"")-1, 0, nil)
	} else {
		extension_dir = core.CoreGlobals.extension_dir
	}
	if type_ == 2 {
		error_type = 1 << 1
	} else {
		error_type = 1 << 5
	}

	/* Check if passed filename contains directory separators */

	if strchr(filename, '/') != nil || strchr(filename, '/') != nil {

		/* Passing modules with full path is not supported for dynamically loaded extensions */

		if type_ == 2 {
			core.PhpErrorDocref(nil, 1<<1, "Temporary module name should contain only filename")
			return zend.FAILURE
		}
		libpath = zend._estrdup(filename)
	} else if extension_dir != nil && extension_dir[0] {
		slash_suffix = extension_dir[strlen(extension_dir)-1] == '/'

		/* Try as filename first */

		if slash_suffix != 0 {
			zend.ZendSpprintf(&libpath, 0, "%s%s", extension_dir, filename)
		} else {
			zend.ZendSpprintf(&libpath, 0, "%s%c%s", extension_dir, '/', filename)
		}

		/* Try as filename first */

	} else {
		return zend.FAILURE
	}
	handle = PhpLoadShlib(libpath, &err1)
	if !handle {

		/* Now, consider 'filename' as extension name and build file name */

		var orig_libpath *byte = libpath
		if slash_suffix != 0 {
			zend.ZendSpprintf(&libpath, 0, "%s"+""+"%s."+"so", extension_dir, filename)
		} else {
			zend.ZendSpprintf(&libpath, 0, "%s%c"+""+"%s."+"so", extension_dir, '/', filename)
		}
		handle = PhpLoadShlib(libpath, &err2)
		if !handle {
			core.PhpErrorDocref(nil, error_type, "Unable to load dynamic library '%s' (tried: %s (%s), %s (%s))", filename, orig_libpath, err1, libpath, err2)
			zend._efree(orig_libpath)
			zend._efree(err1)
			zend._efree(libpath)
			zend._efree(err2)
			return zend.FAILURE
		}
		zend._efree(orig_libpath)
		zend._efree(err1)
	}
	zend._efree(libpath)
	get_module = (func() *zend.ZendModuleEntry)(dlsym(handle, "get_module"))

	/* Some OS prepend _ to symbol names while their dynamic linker
	 * does not do that automatically. Thus we check manually for
	 * _get_module. */

	if get_module == nil {
		get_module = (func() *zend.ZendModuleEntry)(dlsym(handle, "_get_module"))
	}
	if get_module == nil {
		if dlsym(handle, "zend_extension_entry") || dlsym(handle, "_zend_extension_entry") {
			dlclose(handle)
			core.PhpErrorDocref(nil, error_type, "Invalid library (appears to be a Zend Extension, try loading using zend_extension=%s from php.ini)", filename)
			return zend.FAILURE
		}
		dlclose(handle)
		core.PhpErrorDocref(nil, error_type, "Invalid library (maybe not a PHP library) '%s'", filename)
		return zend.FAILURE
	}
	module_entry = get_module()
	if module_entry.zend_api != 20190902 {
		core.PhpErrorDocref(nil, error_type, "%s: Unable to initialize module\n"+"Module compiled with module API=%d\n"+"PHP    compiled with module API=%d\n"+"These options need to match\n", module_entry.name, module_entry.zend_api, 20190902)
		dlclose(handle)
		return zend.FAILURE
	}
	if strcmp(module_entry.build_id, "API"+"20190902"+",NTS") {
		core.PhpErrorDocref(nil, error_type, "%s: Unable to initialize module\n"+"Module compiled with build ID=%s\n"+"PHP    compiled with build ID=%s\n"+"These options need to match\n", module_entry.name, module_entry.build_id, "API"+"20190902"+",NTS")
		dlclose(handle)
		return zend.FAILURE
	}
	module_entry.type_ = type_
	module_entry.module_number = zend.ZendNextFreeModule()
	module_entry.handle = handle
	if g.Assign(&module_entry, zend.ZendRegisterModuleEx(module_entry)) == nil {
		dlclose(handle)
		return zend.FAILURE
	}
	if (type_ == 2 || start_now != 0) && zend.ZendStartupModuleEx(module_entry) == zend.FAILURE {
		dlclose(handle)
		return zend.FAILURE
	}
	if (type_ == 2 || start_now != 0) && module_entry.request_startup_func != nil {
		if module_entry.request_startup_func(type_, module_entry.module_number) == zend.FAILURE {
			core.PhpErrorDocref(nil, error_type, "Unable to initialize module '%s'", module_entry.name)
			dlclose(handle)
			return zend.FAILURE
		}
	}
	return zend.SUCCESS
}

/* }}} */

func PhpDl(file *byte, type_ int, return_value *zend.Zval, start_now int) {
	/* Load extension */

	if PhpLoadExtension(file, type_, start_now) == zend.FAILURE {
		return_value.u1.type_info = 2
	} else {
		return_value.u1.type_info = 3
	}

	/* Load extension */
}

/* }}} */

func ZmInfoDl(zend_module *zend.ZendModuleEntry) {
	PhpInfoPrintTableRow(2, "Dynamic Library Support", "enabled")
}
