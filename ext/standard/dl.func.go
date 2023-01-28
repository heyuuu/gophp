// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func GET_DL_ERROR() __auto__ { return zend.DL_ERROR() }
func ZifDl(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
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
			if zend.ZendParseArgString(_arg, &filename, &filename_len, 0) == 0 {
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
	if !(core.PG(enable_dl)) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Dynamically loaded extensions aren't enabled")
		zend.RETVAL_FALSE
		return
	}
	if filename_len >= core.MAXPATHLEN {
		core.PhpErrorDocref(nil, zend.E_WARNING, "File name exceeds the maximum allowed length of %d characters", core.MAXPATHLEN)
		zend.RETVAL_FALSE
		return
	}
	PhpDl(filename, zend.MODULE_TEMPORARY, return_value, 0)
	if return_value.GetType() == zend.IS_TRUE {
		zend.ExecutorGlobals.SetFullTablesCleanup(1)
	}
}
func PhpLoadShlib(path *byte, errp **byte) any {
	var handle any
	var err *byte
	handle = zend.DL_LOAD(path)
	if !handle {
		err = GET_DL_ERROR()
		*errp = zend.Estrdup(err)
		GET_DL_ERROR()
	}
	return handle
}
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
	if type_ == zend.MODULE_PERSISTENT {
		extension_dir = zend.INI_STR("extension_dir")
	} else {
		extension_dir = core.PG(extension_dir)
	}
	if type_ == zend.MODULE_TEMPORARY {
		error_type = zend.E_WARNING
	} else {
		error_type = zend.E_CORE_WARNING
	}

	/* Check if passed filename contains directory separators */

	if strchr(filename, '/') != nil || strchr(filename, zend.DEFAULT_SLASH) != nil {

		/* Passing modules with full path is not supported for dynamically loaded extensions */

		if type_ == zend.MODULE_TEMPORARY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Temporary module name should contain only filename")
			return zend.FAILURE
		}
		libpath = zend.Estrdup(filename)
	} else if extension_dir != nil && extension_dir[0] {
		slash_suffix = zend.IS_SLASH(extension_dir[strlen(extension_dir)-1])

		/* Try as filename first */

		if slash_suffix != 0 {
			core.Spprintf(&libpath, 0, "%s%s", extension_dir, filename)
		} else {
			core.Spprintf(&libpath, 0, "%s%c%s", extension_dir, zend.DEFAULT_SLASH, filename)
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
			core.Spprintf(&libpath, 0, "%s"+core.PHP_SHLIB_EXT_PREFIX+"%s."+core.PHP_SHLIB_SUFFIX, extension_dir, filename)
		} else {
			core.Spprintf(&libpath, 0, "%s%c"+core.PHP_SHLIB_EXT_PREFIX+"%s."+core.PHP_SHLIB_SUFFIX, extension_dir, zend.DEFAULT_SLASH, filename)
		}
		handle = PhpLoadShlib(libpath, &err2)
		if !handle {
			core.PhpErrorDocref(nil, error_type, "Unable to load dynamic library '%s' (tried: %s (%s), %s (%s))", filename, orig_libpath, err1, libpath, err2)
			zend.Efree(orig_libpath)
			zend.Efree(err1)
			zend.Efree(libpath)
			zend.Efree(err2)
			return zend.FAILURE
		}
		zend.Efree(orig_libpath)
		zend.Efree(err1)
	}
	zend.Efree(libpath)
	get_module = (func() *zend.ZendModuleEntry)(zend.DL_FETCH_SYMBOL(handle, "get_module"))

	/* Some OS prepend _ to symbol names while their dynamic linker
	 * does not do that automatically. Thus we check manually for
	 * _get_module. */

	if get_module == nil {
		get_module = (func() *zend.ZendModuleEntry)(zend.DL_FETCH_SYMBOL(handle, "_get_module"))
	}
	if get_module == nil {
		if zend.DL_FETCH_SYMBOL(handle, "zend_extension_entry") || zend.DL_FETCH_SYMBOL(handle, "_zend_extension_entry") {
			zend.DL_UNLOAD(handle)
			core.PhpErrorDocref(nil, error_type, "Invalid library (appears to be a Zend Extension, try loading using zend_extension=%s from php.ini)", filename)
			return zend.FAILURE
		}
		zend.DL_UNLOAD(handle)
		core.PhpErrorDocref(nil, error_type, "Invalid library (maybe not a PHP library) '%s'", filename)
		return zend.FAILURE
	}
	module_entry = get_module()
	if module_entry.GetZendApi() != zend.ZEND_MODULE_API_NO {
		core.PhpErrorDocref(nil, error_type, "%s: Unable to initialize module\n"+"Module compiled with module API=%d\n"+"PHP    compiled with module API=%d\n"+"These options need to match\n", module_entry.GetName(), module_entry.GetZendApi(), zend.ZEND_MODULE_API_NO)
		zend.DL_UNLOAD(handle)
		return zend.FAILURE
	}
	if strcmp(module_entry.GetBuildId(), "API"+"ZEND_MODULE_API_NO"+zend.ZEND_BUILD_TS) {
		core.PhpErrorDocref(nil, error_type, "%s: Unable to initialize module\n"+"Module compiled with build ID=%s\n"+"PHP    compiled with build ID=%s\n"+"These options need to match\n", module_entry.GetName(), module_entry.GetBuildId(), "API"+"ZEND_MODULE_API_NO"+zend.ZEND_BUILD_TS)
		zend.DL_UNLOAD(handle)
		return zend.FAILURE
	}
	module_entry.SetType(type_)
	module_entry.SetModuleNumber(zend.ZendNextFreeModule())
	module_entry.SetHandle(handle)
	if b.Assign(&module_entry, zend.ZendRegisterModuleEx(module_entry)) == nil {
		zend.DL_UNLOAD(handle)
		return zend.FAILURE
	}
	if (type_ == zend.MODULE_TEMPORARY || start_now != 0) && zend.ZendStartupModuleEx(module_entry) == zend.FAILURE {
		zend.DL_UNLOAD(handle)
		return zend.FAILURE
	}
	if (type_ == zend.MODULE_TEMPORARY || start_now != 0) && module_entry.GetRequestStartupFunc() != nil {
		if module_entry.GetRequestStartupFunc()(type_, module_entry.GetModuleNumber()) == zend.FAILURE {
			core.PhpErrorDocref(nil, error_type, "Unable to initialize module '%s'", module_entry.GetName())
			zend.DL_UNLOAD(handle)
			return zend.FAILURE
		}
	}
	return zend.SUCCESS
}
func PhpDl(file *byte, type_ int, return_value *zend.Zval, start_now int) {
	/* Load extension */

	if PhpLoadExtension(file, type_, start_now) == zend.FAILURE {
		zend.RETVAL_FALSE
	} else {
		zend.RETVAL_TRUE
	}

	/* Load extension */
}
func ZmInfoDl(ZEND_MODULE_INFO_FUNC_ARGS) {
	PhpInfoPrintTableRow(2, "Dynamic Library Support", "enabled")
}
