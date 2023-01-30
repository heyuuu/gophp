// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/sapi/cli"
	"sik/zend"
)

func BG(v *uint32) __auto__ { return BasicGlobals.v }
func PhpPutenvDestructor(zv *zend.Zval) {
	var pe *PutenvEntry = zv.GetPtr()
	if pe.GetPreviousValue() != nil {
		putenv(pe.GetPreviousValue())
	} else {
		unsetenv(pe.GetKey())
	}

	/* don't forget to reset the various libc globals that
	 * we might have changed by an earlier call to tzset(). */

	if !(strncmp(pe.GetKey(), "TZ", pe.GetKeyLen())) {
		tzset()
	}
	zend.Efree(pe.GetPutenvString())
	zend.Efree(pe.GetKey())
	zend.Efree(pe)
}
func BasicGlobalsCtor(basic_globals_p *PhpBasicGlobals) {
	BG(mt_rand_is_seeded) = 0
	BG(mt_rand_mode) = MT_RAND_MT19937
	BG(umask) = -1
	BG(next) = nil
	BG(left) = -1
	BG(user_tick_functions) = nil
	BG(user_filter_map) = nil
	BG(serialize_lock) = 0
	memset(&(BG(serialize)), 0, b.SizeOf("BG ( serialize )"))
	memset(&(BG(unserialize)), 0, b.SizeOf("BG ( unserialize )"))
	memset(&(BG(url_adapt_session_ex)), 0, b.SizeOf("BG ( url_adapt_session_ex )"))
	memset(&(BG(url_adapt_output_ex)), 0, b.SizeOf("BG ( url_adapt_output_ex )"))
	BG(url_adapt_session_ex).type_ = 1
	BG(url_adapt_output_ex).type_ = 0
	zend.ZendHashInit(&(BG(url_adapt_session_hosts_ht)), 0, nil, nil, 1)
	zend.ZendHashInit(&(BG(url_adapt_output_hosts_ht)), 0, nil, nil, 1)
	BG(incomplete_class) = IncompleteClassEntry
	BG(page_uid) = -1
	BG(page_gid) = -1
}
func BasicGlobalsDtor(basic_globals_p *PhpBasicGlobals) {
	if basic_globals_p.GetUrlAdaptSessionEx().GetTags() != nil {
		zend.ZendHashDestroy(basic_globals_p.GetUrlAdaptSessionEx().GetTags())
		zend.Free(basic_globals_p.GetUrlAdaptSessionEx().GetTags())
	}
	if basic_globals_p.GetUrlAdaptOutputEx().GetTags() != nil {
		zend.ZendHashDestroy(basic_globals_p.GetUrlAdaptOutputEx().GetTags())
		zend.Free(basic_globals_p.GetUrlAdaptOutputEx().GetTags())
	}
	zend.ZendHashDestroy(basic_globals_p.GetUrlAdaptSessionHostsHt())
	zend.ZendHashDestroy(basic_globals_p.GetUrlAdaptOutputHostsHt())
}
func PhpGetNan() float64 { return zend.ZEND_NAN }
func PhpGetInf() float64 { return zend.ZEND_INFINITY }
func ZmStartupBasic(type_ int, module_number int) int {
	BasicGlobalsCtor(&BasicGlobals)
	IncompleteClassEntry = PhpCreateIncompleteClass()
	BG(incomplete_class) = IncompleteClassEntry
	zend.REGISTER_LONG_CONSTANT("CONNECTION_ABORTED", core.PHP_CONNECTION_ABORTED, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CONNECTION_NORMAL", core.PHP_CONNECTION_NORMAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CONNECTION_TIMEOUT", core.PHP_CONNECTION_TIMEOUT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_USER", zend.ZEND_INI_USER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_PERDIR", zend.ZEND_INI_PERDIR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_SYSTEM", zend.ZEND_INI_SYSTEM, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_ALL", zend.ZEND_INI_ALL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_SCANNER_NORMAL", zend.ZEND_INI_SCANNER_NORMAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_SCANNER_RAW", zend.ZEND_INI_SCANNER_RAW, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("INI_SCANNER_TYPED", zend.ZEND_INI_SCANNER_TYPED, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_SCHEME", PHP_URL_SCHEME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_HOST", PHP_URL_HOST, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_PORT", PHP_URL_PORT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_USER", PHP_URL_USER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_PASS", PHP_URL_PASS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_PATH", PHP_URL_PATH, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_QUERY", PHP_URL_QUERY, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_URL_FRAGMENT", PHP_URL_FRAGMENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_QUERY_RFC1738", PHP_QUERY_RFC1738, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_QUERY_RFC3986", PHP_QUERY_RFC3986, zend.CONST_CS|zend.CONST_PERSISTENT)

	// #define REGISTER_MATH_CONSTANT(x) REGISTER_DOUBLE_CONSTANT ( # x , x , CONST_CS | CONST_PERSISTENT )

	zend.REGISTER_DOUBLE_CONSTANT("M_E", M_E, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_LOG2E", M_LOG2E, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_LOG10E", M_LOG10E, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_LN2", M_LN2, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_LN10", M_LN10, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_PI", M_PI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_PI_2", M_PI_2, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_PI_4", M_PI_4, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_1_PI", M_1_PI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_2_PI", M_2_PI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_SQRTPI", M_SQRTPI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_2_SQRTPI", M_2_SQRTPI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_LNPI", M_LNPI, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_EULER", M_EULER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_SQRT2", M_SQRT2, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_SQRT1_2", M_SQRT1_2, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("M_SQRT3", M_SQRT3, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("INF", zend.ZEND_INFINITY, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_DOUBLE_CONSTANT("NAN", zend.ZEND_NAN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_ROUND_HALF_UP", PHP_ROUND_HALF_UP, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_ROUND_HALF_DOWN", PHP_ROUND_HALF_DOWN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_ROUND_HALF_EVEN", PHP_ROUND_HALF_EVEN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PHP_ROUND_HALF_ODD", PHP_ROUND_HALF_ODD, zend.CONST_CS|zend.CONST_PERSISTENT)
	RegisterPhpinfoConstants(type_, module_number)
	RegisterHtmlConstants(type_, module_number)
	RegisterStringConstants(type_, module_number)
	if ZmStartupVar(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupFile(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupPack(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupBrowscap(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupStandardFilters(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupUserFilters(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupPassword(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupMtRand(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if zm_startup_nl_langinfo(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupCrypt(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupLcg(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupDir(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupSyslog(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupArray(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupAssert(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupUrlScannerEx(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupProcOpen(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupExec(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupUserStreams(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupImagetypes(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	streams.PhpRegisterUrlStreamWrapper("php", &PhpStreamPhpWrapper)
	streams.PhpRegisterUrlStreamWrapper("file", &PhpPlainFilesWrapper)
	streams.PhpRegisterUrlStreamWrapper("glob", &streams.PhpGlobStreamWrapper)
	streams.PhpRegisterUrlStreamWrapper("data", &streams.PhpStreamRfc2397Wrapper)
	streams.PhpRegisterUrlStreamWrapper("http", &PhpStreamHttpWrapper)
	streams.PhpRegisterUrlStreamWrapper("ftp", &PhpStreamFtpWrapper)
	if ZmStartupDns(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupRandom(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	if ZmStartupHrtime(type_, module_number) != zend.SUCCESS {
		return zend.FAILURE
	}
	return zend.SUCCESS
}
func ZmShutdownBasic(type_ int, module_number int) int {
	ZmShutdownSyslog(type_, module_number)
	BasicGlobalsDtor(&BasicGlobals)
	streams.PhpUnregisterUrlStreamWrapper("php")
	streams.PhpUnregisterUrlStreamWrapper("http")
	streams.PhpUnregisterUrlStreamWrapper("ftp")
	ZmShutdownBrowscap(type_, module_number)
	ZmShutdownArray(type_, module_number)
	ZmShutdownAssert(type_, module_number)
	ZmShutdownUrlScannerEx(type_, module_number)
	ZmShutdownFile(type_, module_number)
	ZmShutdownStandardFilters(type_, module_number)
	ZmShutdownCrypt(type_, module_number)
	ZmShutdownRandom(type_, module_number)
	ZmShutdownPassword(type_, module_number)
	return zend.SUCCESS
}
func ZmActivateBasic(type_ int, module_number int) int {
	memset(BG(strtok_table), 0, 256)
	BG(serialize_lock) = 0
	memset(&(BG(serialize)), 0, b.SizeOf("BG ( serialize )"))
	memset(&(BG(unserialize)), 0, b.SizeOf("BG ( unserialize )"))
	BG(strtok_string) = nil
	zend.ZVAL_UNDEF(&(BG(strtok_zval)))
	BG(strtok_last) = nil
	BG(locale_string) = nil
	BG(locale_changed) = 0
	BG(array_walk_fci) = zend.EmptyFcallInfo
	BG(array_walk_fci_cache) = zend.EmptyFcallInfoCache
	BG(user_compare_fci) = zend.EmptyFcallInfo
	BG(user_compare_fci_cache) = zend.EmptyFcallInfoCache
	BG(page_uid) = -1
	BG(page_gid) = -1
	BG(page_inode) = -1
	BG(page_mtime) = -1
	zend.ZendHashInit(&(BG(putenv_ht)), 1, nil, PhpPutenvDestructor, 0)
	BG(user_shutdown_function_names) = nil
	ZmActivateFilestat(type_, module_number)
	ZmActivateSyslog(type_, module_number)
	ZmActivateDir(type_, module_number)
	ZmActivateUrlScannerEx(type_, module_number)

	/* Setup default context */

	FG(default_context) = nil

	/* Default to global wrappers only */

	FG(stream_wrappers) = nil

	/* Default to global filters only */

	FG(stream_filters) = nil
	return zend.SUCCESS
}
func ZmDeactivateBasic(type_ int, module_number int) int {
	zend.ZvalPtrDtor(&(BG(strtok_zval)))
	zend.ZVAL_UNDEF(&(BG(strtok_zval)))
	BG(strtok_string) = nil
	tsrm_env_lock()
	zend.ZendHashDestroy(&(BG(putenv_ht)))
	tsrm_env_unlock()
	BG(mt_rand_is_seeded) = 0
	if BG(umask) != -1 {
		umask(BG(umask))
	}

	/* Check if locale was changed and change it back
	 * to the value in startup environment */

	if BG(locale_changed) {
		setlocale(LC_ALL, "C")
		setlocale(LC_CTYPE, "")
		if BG(locale_string) {
			zend.ZendStringReleaseEx(BG(locale_string), 0)
			BG(locale_string) = nil
		}
	}

	/* FG(stream_wrappers) and FG(stream_filters) are destroyed
	 * during php_request_shutdown() */

	ZmDeactivateFilestat(type_, module_number)
	ZmDeactivateAssert(type_, module_number)
	ZmDeactivateUrlScannerEx(type_, module_number)
	streams.ZmDeactivateStreams(type_, module_number)
	if BG(user_tick_functions) {
		zend.ZendLlistDestroy(BG(user_tick_functions))
		zend.Efree(BG(user_tick_functions))
		BG(user_tick_functions) = nil
	}
	ZmDeactivateUserFilters(type_, module_number)
	ZmDeactivateBrowscap(type_, module_number)
	BG(page_uid) = -1
	BG(page_gid) = -1
	return zend.SUCCESS
}
func ZmInfoBasic(ZEND_MODULE_INFO_FUNC_ARGS) {
	PhpInfoPrintTableStart()
	ZmInfoDl(zend.ZEND_MODULE_INFO_FUNC_ARGS_PASSTHRU)
	ZmInfoMail(zend.ZEND_MODULE_INFO_FUNC_ARGS_PASSTHRU)
	PhpInfoPrintTableEnd()
	ZmInfoAssert(zend.ZEND_MODULE_INFO_FUNC_ARGS_PASSTHRU)
}
func ZifConstant(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var const_name *zend.ZendString
	var c *zend.Zval
	var scope *zend.ZendClassEntry
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
			if zend.ZendParseArgStr(_arg, &const_name, 0) == 0 {
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
	scope = zend.ZendGetExecutedScope()
	c = zend.ZendGetConstantEx(const_name, scope, zend.ZEND_FETCH_CLASS_SILENT)
	if c != nil {
		zend.ZVAL_COPY_OR_DUP(return_value, c)
		if return_value.IsType(zend.IS_CONSTANT_AST) {
			if zend.ZvalUpdateConstantEx(return_value, scope) != zend.SUCCESS {
				return
			}
		}
	} else {
		if zend.__EG().GetException() == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Couldn't find constant %s", const_name.GetVal())
		}
		zend.RETVAL_NULL()
		return
	}
}
func ZifInetNtop(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var address *byte
	var address_len int
	var af int = AF_INET
	var buffer []byte
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
			if zend.ZendParseArgString(_arg, &address, &address_len, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if address_len == 16 {
		af = AF_INET6
	} else if address_len != 4 {
		zend.RETVAL_FALSE
		return
	}
	if !(inet_ntop(af, address, buffer, b.SizeOf("buffer"))) {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRING(buffer)
	return
}
func PhpInetPton(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ret int
	var af int = AF_INET
	var address *byte
	var address_len int
	var buffer []byte
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
			if zend.ZendParseArgString(_arg, &address, &address_len, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	memset(buffer, 0, b.SizeOf("buffer"))
	if strchr(address, ':') {
		af = AF_INET6
	} else if !(strchr(address, '.')) {
		zend.RETVAL_FALSE
		return
	}
	ret = inet_pton(af, address, buffer)
	if ret <= 0 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRINGL(buffer, b.Cond(af == AF_INET, 4, 16))
	return
}
func ZifIp2long(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var addr *byte
	var addr_len int
	var ip __struct__in_addr
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
			if zend.ZendParseArgString(_arg, &addr, &addr_len, 0) == 0 {
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
	if addr_len == 0 || inet_pton(AF_INET, addr, &ip) != 1 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(ntohl(ip.s_addr))
	return
}
func ZifLong2ip(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ip zend.ZendUlong
	var sip zend.ZendLong
	var myaddr __struct__in_addr
	var str []byte
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
			if zend.ZendParseArgLong(_arg, &sip, &_dummy, 0, 0) == 0 {
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

	/* autoboxes on 32bit platforms, but that's expected */

	ip = zend.ZendUlong(sip)
	myaddr.s_addr = htonl(ip)
	if inet_ntop(AF_INET, &myaddr, str, b.SizeOf("str")) {
		zend.RETVAL_STRING(str)
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifGetenv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ptr *byte
	var str *byte = nil
	var str_len int
	var local_only zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &local_only, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	if str == nil {
		zend.ArrayInit(return_value)
		core.PhpImportEnvironmentVariables(return_value)
		return
	}
	if local_only == 0 {

		/* SAPI method returns an emalloc()'d string */

		ptr = core.SapiGetenv(str, str_len)
		if ptr != nil {

			// TODO: avoid realocation ???

			zend.RETVAL_STRING(ptr)
			zend.Efree(ptr)
			return
		}
	}
	tsrm_env_lock()

	/* system method returns a const */

	ptr = getenv(str)
	if ptr != nil {
		zend.RETVAL_STRING(ptr)
	}
	tsrm_env_unlock()
	if ptr != nil {
		return
	}
	zend.RETVAL_FALSE
	return
}
func ZifPutenv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var setting *byte
	var setting_len int
	var p *byte
	var env **byte
	var pe PutenvEntry
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
			if zend.ZendParseArgString(_arg, &setting, &setting_len, 0) == 0 {
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
	if setting_len == 0 || setting[0] == '=' {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid parameter syntax")
		zend.RETVAL_FALSE
		return
	}
	pe.SetPutenvString(zend.Estrndup(setting, setting_len))
	pe.SetKey(zend.Estrndup(setting, setting_len))
	if b.Assign(&p, strchr(pe.GetKey(), '=')) {
		*p = '0'
	}
	pe.SetKeyLen(strlen(pe.GetKey()))
	tsrm_env_lock()
	zend.ZendHashStrDel(&(BG(putenv_ht)), pe.GetKey(), pe.GetKeyLen())

	/* find previous value */

	pe.SetPreviousValue(nil)
	for env = cli.Environ; env != nil && (*env) != nil; env++ {
		if !(strncmp(*env, pe.GetKey(), pe.GetKeyLen())) && (*env)[pe.GetKeyLen()] == '=' {
			pe.SetPreviousValue(*env)
			break
		}
	}
	if p == nil {
		unsetenv(pe.GetPutenvString())
	}
	if p == nil || putenv(pe.GetPutenvString()) == 0 {
		zend.ZendHashStrAddMem(&(BG(putenv_ht)), pe.GetKey(), pe.GetKeyLen(), &pe, b.SizeOf("putenv_entry"))
		if !(strncmp(pe.GetKey(), "TZ", pe.GetKeyLen())) {
			tzset()
		}
		tsrm_env_unlock()
		zend.RETVAL_TRUE
		return
	} else {
		zend.Efree(pe.GetPutenvString())
		zend.Efree(pe.GetKey())
		zend.RETVAL_FALSE
		return
	}
}
func FreeArgv(argv **byte, argc int) {
	var i int
	if argv != nil {
		for i = 0; i < argc; i++ {
			if argv[i] != nil {
				zend.Efree(argv[i])
			}
		}
		zend.Efree(argv)
	}
}
func FreeLongopts(longopts *core.Opt) {
	var p *core.Opt
	if longopts != nil {
		for p = longopts; p != nil && p.GetOptChar() != '-'; p++ {
			if p.GetOptName() != nil {
				zend.Efree((*byte)(p.GetOptName()))
			}
		}
	}
}
func ParseOpts(opts *byte, result **core.Opt) int {
	var paras *core.Opt = nil
	var i uint
	var count uint = 0
	var opts_len uint = uint(strlen(opts))
	for i = 0; i < opts_len; i++ {
		if opts[i] >= 48 && opts[i] <= 57 || opts[i] >= 65 && opts[i] <= 90 || opts[i] >= 97 && opts[i] <= 122 {
			count++
		}
	}
	paras = zend.SafeEmalloc(b.SizeOf("opt_struct"), count, 0)
	memset(paras, 0, b.SizeOf("opt_struct")*count)
	*result = paras
	for (*opts) >= 48 && (*opts) <= 57 || (*opts) >= 65 && (*opts) <= 90 || (*opts) >= 97 && (*opts) <= 122 {
		paras.SetOptChar(*opts)
		paras.SetNeedParam((*(b.PreInc(&opts))) == ':')
		paras.SetOptName(nil)
		if paras.GetNeedParam() == 1 {
			opts++
			if (*opts) == ':' {
				paras.GetNeedParam()++
				opts++
			}
		}
		paras++
	}
	return count
}
func ZifGetopt(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var options *byte = nil
	var argv **byte = nil
	var opt []byte = []byte{'0'}
	var optname *byte
	var argc int = 0
	var o int
	var options_len int = 0
	var len_ int
	var php_optarg *byte = nil
	var php_optind int = 1
	var val zend.Zval
	var args *zend.Zval = nil
	var p_longopts *zend.Zval = nil
	var zoptind *zend.Zval = nil
	var optname_len int = 0
	var opts *core.Opt
	var orig_opts *core.Opt
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
			if zend.ZendParseArgString(_arg, &options, &options_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &p_longopts, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zoptind, 0)
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

	/* Init zoptind to 1 */

	if zoptind != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zoptind, 1)
	}

	/* Get argv from the global symbol table. We calculate argc ourselves
	 * in order to be on the safe side, even though it is also available
	 * from the symbol table. */

	if (core.PG(http_globals)[core.TRACK_VARS_SERVER].u1.v.type_ == zend.IS_ARRAY || zend.ZendIsAutoGlobalStr(zend.ZEND_STRL("_SERVER")) != 0) && (b.Assign(&args, zend.ZendHashFindExInd(core.PG(http_globals)[core.TRACK_VARS_SERVER].GetArr(), zend.ZSTR_KNOWN(zend.ZEND_STR_ARGV), 1)) != nil || b.Assign(&args, zend.ZendHashFindExInd(zend.__EG().GetSymbolTable(), zend.ZSTR_KNOWN(zend.ZEND_STR_ARGV), 1)) != nil) {
		var pos int = 0
		var entry *zend.Zval
		if args.GetType() != zend.IS_ARRAY {
			zend.RETVAL_FALSE
			return
		}
		argc = zend.Z_ARRVAL_P(args).GetNNumOfElements()

		/* Attempt to allocate enough memory to hold all of the arguments
		 * and a trailing NULL */

		argv = (**byte)(zend.SafeEmalloc(b.SizeOf("char *"), argc+1, 0))

		/* Iterate over the hash to construct the argv array. */

		var __ht *zend.HashTable = args.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()

			entry = _z
			var tmp_arg_str *zend.ZendString
			var arg_str *zend.ZendString = zend.ZvalGetTmpString(entry, &tmp_arg_str)
			argv[b.PostInc(&pos)] = zend.Estrdup(arg_str.GetVal())
			zend.ZendTmpStringRelease(tmp_arg_str)
		}

		/* The C Standard requires argv[argc] to be NULL - this might
		 * keep some getopt implementations happy. */

		argv[argc] = nil

		/* The C Standard requires argv[argc] to be NULL - this might
		 * keep some getopt implementations happy. */

	} else {

		/* Return false if we can't find argv. */

		zend.RETVAL_FALSE
		return
	}
	len_ = ParseOpts(options, &opts)
	if p_longopts != nil {
		var count int
		var entry *zend.Zval
		count = zend.Z_ARRVAL_P(p_longopts).GetNNumOfElements()

		/* the first <len> slots are filled by the one short ops
		 * we now extend our array and jump to the new added structs */

		opts = (*core.Opt)(zend.Erealloc(opts, b.SizeOf("opt_struct")*(len_+count+1)))
		orig_opts = opts
		opts += len_
		memset(opts, 0, count*b.SizeOf("opt_struct"))

		/* Iterate over the hash to construct the argv array. */

		var __ht *zend.HashTable = p_longopts.GetArr()
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()

			entry = _z
			var tmp_arg_str *zend.ZendString
			var arg_str *zend.ZendString = zend.ZvalGetTmpString(entry, &tmp_arg_str)
			opts.SetNeedParam(0)
			opts.SetOptName(zend.Estrdup(arg_str.GetVal()))
			len_ = strlen(opts.GetOptName())
			if len_ > 0 && opts.GetOptName()[len_-1] == ':' {
				opts.GetNeedParam()++
				opts.GetOptName()[len_-1] = '0'
				if len_ > 1 && opts.GetOptName()[len_-2] == ':' {
					opts.GetNeedParam()++
					opts.GetOptName()[len_-2] = '0'
				}
			}
			opts.SetOptChar(0)
			opts++
			zend.ZendTmpStringRelease(tmp_arg_str)
		}

		/* Iterate over the hash to construct the argv array. */

	} else {
		opts = (*core.Opt)(zend.Erealloc(opts, b.SizeOf("opt_struct")*(len_+1)))
		orig_opts = opts
		opts += len_
	}

	/* php_getopt want to identify the last param */

	opts.SetOptChar('-')
	opts.SetNeedParam(0)
	opts.SetOptName(nil)

	/* Initialize the return value as an array. */

	zend.ArrayInit(return_value)

	/* after our pointer arithmetic jump back to the first element */

	opts = orig_opts
	for b.Assign(&o, core.PhpGetopt(argc, argv, opts, &php_optarg, &php_optind, 0, 1)) != -1 {

		/* Skip unknown arguments. */

		if o == core.PHP_GETOPT_INVALID_ARG {
			continue
		}

		/* Prepare the option character and the argument string. */

		if o == 0 {
			optname = opts[core.PhpOptidx].GetOptName()
		} else {
			if o == 1 {
				o = '-'
			}
			opt[0] = o
			optname = opt
		}
		if php_optarg != nil {

			/* keep the arg as binary, since the encoding is not known */

			zend.ZVAL_STRING(&val, php_optarg)

			/* keep the arg as binary, since the encoding is not known */

		} else {
			zend.ZVAL_FALSE(&val)
		}

		/* Add this option / argument pair to the result hash. */

		optname_len = strlen(optname)
		if !(optname_len > 1 && optname[0] == '0') && zend.IsNumericString(optname, optname_len, nil, nil, 0) == zend.IS_LONG {

			/* numeric string */

			var optname_int int = atoi(optname)
			if b.Assign(&args, zend.ZendHashIndexFind(return_value.GetArr(), optname_int)) != nil {
				if args.GetType() != zend.IS_ARRAY {
					zend.ConvertToArrayEx(args)
				}
				args.GetArr().NextIndexInsert(&val)
			} else {
				return_value.GetArr().IndexUpdateH(optname_int, &val)
			}
		} else {

			/* other strings */

			if b.Assign(&args, return_value.GetArr().FindByStrPtr(optname, strlen(optname))) != nil {
				if args.GetType() != zend.IS_ARRAY {
					zend.ConvertToArrayEx(args)
				}
				args.GetArr().NextIndexInsert(&val)
			} else {
				zend.ZendHashStrAdd(return_value.GetArr(), b.CastStr(optname, strlen(optname)), &val)
			}

			/* other strings */

		}
		php_optarg = nil
	}

	/* Set zoptind to php_optind */

	if zoptind != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zoptind, php_optind)
	}
	FreeLongopts(orig_opts)
	zend.Efree(orig_opts)
	FreeArgv(argv, argc)
}
func ZifFlush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	core.SapiFlush()
}
func ZifSleep(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var num zend.ZendLong
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
			if zend.ZendParseArgLong(_arg, &num, &_dummy, 0, 0) == 0 {
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
	if num < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Number of seconds must be greater than or equal to 0")
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(core.PhpSleep(uint(num)))
	return
}
func ZifUsleep(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var num zend.ZendLong
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
			if zend.ZendParseArgLong(_arg, &num, &_dummy, 0, 0) == 0 {
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
	if num < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Number of microseconds must be greater than or equal to 0")
		zend.RETVAL_FALSE
		return
	}
	usleep(uint(num))
}
func ZifTimeNanosleep(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var tv_sec zend.ZendLong
	var tv_nsec zend.ZendLong
	var php_req __struct__timespec
	var php_rem __struct__timespec
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.ZendParseArgLong(_arg, &tv_sec, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &tv_nsec, &_dummy, 0, 0) == 0 {
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
	if tv_sec < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The seconds value must be greater than 0")
		zend.RETVAL_FALSE
		return
	}
	if tv_nsec < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The nanoseconds value must be greater than 0")
		zend.RETVAL_FALSE
		return
	}
	php_req.tv_sec = int64(tv_sec)
	php_req.tv_nsec = long(tv_nsec)
	if !(nanosleep(&php_req, &php_rem)) {
		zend.RETVAL_TRUE
		return
	} else if errno == EINTR {
		zend.ArrayInit(return_value)
		zend.AddAssocLongEx(return_value, "seconds", b.SizeOf("\"seconds\"")-1, php_rem.tv_sec)
		zend.AddAssocLongEx(return_value, "nanoseconds", b.SizeOf("\"nanoseconds\"")-1, php_rem.tv_nsec)
		return
	} else if errno == EINVAL {
		core.PhpErrorDocref(nil, zend.E_WARNING, "nanoseconds was not in the range 0 to 999 999 999 or seconds was negative")
	}
	zend.RETVAL_FALSE
	return
}
func ZifTimeSleepUntil(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var target_secs float64
	var tm __struct__timeval
	var php_req __struct__timespec
	var php_rem __struct__timespec
	var current_ns uint64
	var target_ns uint64
	var diff_ns uint64
	var ns_per_sec uint64 = 1000000000
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
			if zend.ZendParseArgDouble(_arg, &target_secs, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_DOUBLE
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
	if gettimeofday((*__struct__timeval)(&tm), nil) != 0 {
		zend.RETVAL_FALSE
		return
	}
	target_ns = uint64(target_secs * ns_per_sec)
	current_ns = uint64(tm.tv_sec)*ns_per_sec + uint64(tm.tv_usec)*1000
	if target_ns < current_ns {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Sleep until to time is less than current time")
		zend.RETVAL_FALSE
		return
	}
	diff_ns = target_ns - current_ns
	php_req.tv_sec = time_t(diff_ns / ns_per_sec)
	php_req.tv_nsec = long(diff_ns % ns_per_sec)
	for nanosleep(&php_req, &php_rem) {
		if errno == EINTR {
			php_req.tv_sec = php_rem.tv_sec
			php_req.tv_nsec = php_rem.tv_nsec
		} else {
			zend.RETVAL_FALSE
			return
		}
	}
	zend.RETVAL_TRUE
	return
}
func ZifGetCurrentUser(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_STRING(core.PhpGetCurrentUser())
	return
}
func AddConfigEntry(h zend.ZendUlong, key *zend.ZendString, entry *zend.Zval, retval *zend.Zval) {
	if entry.IsType(zend.IS_STRING) {
		var str *zend.ZendString = entry.GetStr()
		if (str.GetGcFlags() & zend.GC_PERSISTENT) == 0 {
			str.AddRefcount()
		} else {
			str = zend.ZendStringInit(str.GetVal(), str.GetLen(), 0)
		}
		if key != nil {
			zend.AddAssocStrEx(retval, key.GetVal(), key.GetLen(), str)
		} else {
			zend.AddIndexStr(retval, h, str)
		}
	} else if entry.IsType(zend.IS_ARRAY) {
		var tmp zend.Zval
		zend.ArrayInit(&tmp)
		AddConfigEntries(entry.GetArr(), &tmp)
		zend.ZendHashUpdate(retval.GetArr(), key.GetStr(), &tmp)
	}
}
func AddConfigEntries(hash *zend.HashTable, return_value *zend.Zval) {
	var h zend.ZendUlong
	var key *zend.ZendString
	var zv *zend.Zval
	var __ht *zend.HashTable = hash
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		h = _p.GetH()
		key = _p.GetKey()
		zv = _z
		AddConfigEntry(h, key, zv, return_value)
	}
}
func ZifGetCfgVar(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var varname *byte
	var varname_len int
	var retval *zend.Zval
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
			if zend.ZendParseArgString(_arg, &varname, &varname_len, 0) == 0 {
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
	retval = core.CfgGetEntry(varname, uint32(varname_len))
	if retval != nil {
		if retval.IsType(zend.IS_ARRAY) {
			zend.ArrayInit(return_value)
			AddConfigEntries(retval.GetArr(), return_value)
			return
		} else {
			zend.RETVAL_STRING(zend.Z_STRVAL_P(retval))
			return
		}
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifGetMagicQuotesRuntime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_FALSE
	return
}
func ZifGetMagicQuotesGpc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_FALSE
	return
}
func ZifErrorLog(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var message *byte
	var opt *byte = nil
	var headers *byte = nil
	var message_len int
	var opt_len int = 0
	var headers_len int = 0
	var opt_err int = 0
	var argc int = zend.ZEND_NUM_ARGS()
	var erropt zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 4
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
			if zend.ZendParseArgString(_arg, &message, &message_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &erropt, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &opt, &opt_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &headers, &headers_len, 0) == 0 {
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
	if argc > 1 {
		opt_err = int(erropt)
	}
	if _phpErrorLogEx(opt_err, message, message_len, opt, headers) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func _phpErrorLog(opt_err int, message *byte, opt *byte, headers *byte) int {
	return _phpErrorLogEx(opt_err, message, b.CondF1(opt_err == 3, func() __auto__ { return strlen(message) }, 0), opt, headers)
}
func _phpErrorLogEx(opt_err int, message *byte, message_len int, opt *byte, headers *byte) int {
	var stream *core.PhpStream = nil
	var nbytes int
	switch opt_err {
	case 1:
		if PhpMail(opt, "PHP error_log message", message, headers, nil) == 0 {
			return zend.FAILURE
		}
		break
	case 2:
		core.PhpErrorDocref(nil, zend.E_WARNING, "TCP/IP option not available!")
		return zend.FAILURE
		break
	case 3:
		stream = core.PhpStreamOpenWrapper(opt, "a", core.IGNORE_URL_WIN|core.REPORT_ERRORS, nil)
		if stream == nil {
			return zend.FAILURE
		}
		nbytes = core.PhpStreamWrite(stream, message, message_len)
		core.PhpStreamClose(stream)
		if nbytes != message_len {
			return zend.FAILURE
		}
		break
	case 4:
		if core.sapi_module.GetLogMessage() != nil {
			core.sapi_module.GetLogMessage()(message, -1)
		} else {
			return zend.FAILURE
		}
		break
	default:
		core.PhpLogErrWithSeverity(message, LOG_NOTICE)
		break
	}
	return zend.SUCCESS
}
func ZifErrorGetLast(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if core.PG(last_error_message) {
		zend.ArrayInit(return_value)
		zend.AddAssocLongEx(return_value, "type", b.SizeOf("\"type\"")-1, core.PG(last_error_type))
		zend.AddAssocStringEx(return_value, "message", b.SizeOf("\"message\"")-1, core.PG(last_error_message))
		zend.AddAssocStringEx(return_value, "file", b.SizeOf("\"file\"")-1, b.CondF1(core.PG(last_error_file), func() __auto__ { return core.PG(last_error_file) }, "-"))
		zend.AddAssocLongEx(return_value, "line", b.SizeOf("\"line\"")-1, core.PG(last_error_lineno))
	}
}
func ZifErrorClearLast(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if core.PG(last_error_message) {
		core.PG(last_error_type) = 0
		core.PG(last_error_lineno) = 0
		zend.Free(core.PG(last_error_message))
		core.PG(last_error_message) = nil
		if core.PG(last_error_file) {
			zend.Free(core.PG(last_error_file))
			core.PG(last_error_file) = nil
		}
	}
}
func ZifCallUserFunc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var retval zend.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
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
			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				fci.SetParams(_real_arg + 1)
				fci.SetParamCount(_num_varargs)
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				fci.SetParams(nil)
				fci.SetParamCount(0)
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
	fci.SetRetval(&retval)
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && retval.GetType() != zend.IS_UNDEF {
		if zend.Z_ISREF(retval) {
			zend.ZendUnwrapReference(&retval)
		}
		zend.ZVAL_COPY_VALUE(return_value, &retval)
	}
}
func ZifCallUserFuncArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var params *zend.Zval
	var retval zend.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &params, 0, 0) == 0 {
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
			return
		}
		break
	}
	zend.ZendFcallInfoArgs(&fci, params)
	fci.SetRetval(&retval)
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && retval.GetType() != zend.IS_UNDEF {
		if zend.Z_ISREF(retval) {
			zend.ZendUnwrapReference(&retval)
		}
		zend.ZVAL_COPY_VALUE(return_value, &retval)
	}
	zend.ZendFcallInfoArgsClear(&fci, 1)
}
func ZifForwardStaticCall(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var retval zend.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache
	var called_scope *zend.ZendClassEntry
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
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
			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				fci.SetParams(_real_arg + 1)
				fci.SetParamCount(_num_varargs)
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				fci.SetParams(nil)
				fci.SetParamCount(0)
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
	if !(zend.EX(prev_execute_data).func_.common.scope) {
		zend.ZendThrowError(nil, "Cannot call forward_static_call() when no class scope is active")
		return
	}
	fci.SetRetval(&retval)
	called_scope = zend.ZendGetCalledScope(execute_data)
	if called_scope != nil && fci_cache.GetCallingScope() != nil && zend.InstanceofFunction(called_scope, fci_cache.GetCallingScope()) != 0 {
		fci_cache.SetCalledScope(called_scope)
	}
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && retval.GetType() != zend.IS_UNDEF {
		if zend.Z_ISREF(retval) {
			zend.ZendUnwrapReference(&retval)
		}
		zend.ZVAL_COPY_VALUE(return_value, &retval)
	}
}
func ZifForwardStaticCallArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var params *zend.Zval
	var retval zend.Zval
	var fci zend.ZendFcallInfo
	var fci_cache zend.ZendFcallInfoCache
	var called_scope *zend.ZendClassEntry
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.ZendParseArgFunc(_arg, &fci, &fci_cache, 0, &_error) == 0 {
				if _error == nil {
					_expected_type = zend.Z_EXPECTED_FUNC
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				} else {
					_error_code = zend.ZPP_ERROR_WRONG_CALLBACK
					break
				}
			} else if _error != nil {
				zend.ZendWrongCallbackDeprecated(_i, _error)
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &params, 0, 0) == 0 {
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
			return
		}
		break
	}
	zend.ZendFcallInfoArgs(&fci, params)
	fci.SetRetval(&retval)
	called_scope = zend.ZendGetCalledScope(execute_data)
	if called_scope != nil && fci_cache.GetCallingScope() != nil && zend.InstanceofFunction(called_scope, fci_cache.GetCallingScope()) != 0 {
		fci_cache.SetCalledScope(called_scope)
	}
	if zend.ZendCallFunction(&fci, &fci_cache) == zend.SUCCESS && retval.GetType() != zend.IS_UNDEF {
		if zend.Z_ISREF(retval) {
			zend.ZendUnwrapReference(&retval)
		}
		zend.ZVAL_COPY_VALUE(return_value, &retval)
	}
	zend.ZendFcallInfoArgsClear(&fci, 1)
}
func UserShutdownFunctionDtor(zv *zend.Zval) {
	var i int
	var shutdown_function_entry *PhpShutdownFunctionEntry = zv.GetPtr()
	for i = 0; i < shutdown_function_entry.GetArgCount(); i++ {
		zend.ZvalPtrDtor(shutdown_function_entry.GetArguments()[i])
	}
	zend.Efree(shutdown_function_entry.GetArguments())
	zend.Efree(shutdown_function_entry)
}
func UserTickFunctionDtor(tick_function_entry *UserTickFunctionEntry) {
	var i int
	for i = 0; i < tick_function_entry.GetArgCount(); i++ {
		zend.ZvalPtrDtor(tick_function_entry.GetArguments()[i])
	}
	zend.Efree(tick_function_entry.GetArguments())
}
func UserShutdownFunctionCall(zv *zend.Zval) int {
	var shutdown_function_entry *PhpShutdownFunctionEntry = zv.GetPtr()
	var retval zend.Zval
	if zend.ZendIsCallable(shutdown_function_entry.GetArguments()[0], 0, nil) == 0 {
		var function_name *zend.ZendString = zend.ZendGetCallableName(shutdown_function_entry.GetArguments()[0])
		core.PhpError(zend.E_WARNING, "(Registered shutdown functions) Unable to call %s() - function does not exist", function_name.GetVal())
		zend.ZendStringReleaseEx(function_name, 0)
		return 0
	}
	if zend.CallUserFunction(nil, nil, shutdown_function_entry.GetArguments()[0], &retval, shutdown_function_entry.GetArgCount()-1, shutdown_function_entry.GetArguments()+1) == zend.SUCCESS {
		zend.ZvalPtrDtor(&retval)
	}
	return 0
}
func UserTickFunctionCall(tick_fe *UserTickFunctionEntry) {
	var retval zend.Zval
	var function *zend.Zval = tick_fe.GetArguments()[0]

	/* Prevent reentrant calls to the same user ticks function */

	if tick_fe.GetCalling() == 0 {
		tick_fe.SetCalling(1)
		if zend.CallUserFunction(nil, nil, function, &retval, tick_fe.GetArgCount()-1, tick_fe.GetArguments()+1) == zend.SUCCESS {
			zend.ZvalPtrDtor(&retval)
		} else {
			var obj *zend.Zval
			var method *zend.Zval
			if function.IsType(zend.IS_STRING) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to call %s() - function does not exist", zend.Z_STRVAL_P(function))
			} else if function.IsType(zend.IS_ARRAY) && b.Assign(&obj, zend.ZendHashIndexFind(function.GetArr(), 0)) != nil && b.Assign(&method, zend.ZendHashIndexFind(function.GetArr(), 1)) != nil && obj.IsType(zend.IS_OBJECT) && method.IsType(zend.IS_STRING) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to call %s::%s() - function does not exist", zend.Z_OBJCE_P(obj).GetName().GetVal(), zend.Z_STRVAL_P(method))
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to call tick function")
			}
		}
		tick_fe.SetCalling(0)
	}

	/* Prevent reentrant calls to the same user ticks function */
}
func RunUserTickFunctions(tick_count int, arg any) {
	zend.ZendLlistApply(BG(user_tick_functions), zend.LlistApplyFuncT(UserTickFunctionCall))
}
func UserTickFunctionCompare(tick_fe1 *UserTickFunctionEntry, tick_fe2 *UserTickFunctionEntry) int {
	var func1 *zend.Zval = tick_fe1.GetArguments()[0]
	var func2 *zend.Zval = tick_fe2.GetArguments()[0]
	var ret int
	if func1.IsType(zend.IS_STRING) && func2.IsType(zend.IS_STRING) {
		ret = zend.ZendBinaryZvalStrcmp(func1, func2) == 0
	} else if func1.IsType(zend.IS_ARRAY) && func2.IsType(zend.IS_ARRAY) {
		ret = zend.ZendCompareArrays(func1, func2) == 0
	} else if func1.IsType(zend.IS_OBJECT) && func2.IsType(zend.IS_OBJECT) {
		ret = zend.ZendCompareObjects(func1, func2) == 0
	} else {
		ret = 0
	}
	if ret != 0 && tick_fe1.GetCalling() != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to delete tick function executed at the moment")
		return 0
	}
	return ret
}
func PhpCallShutdownFunctions() {
	if BG(user_shutdown_function_names) {
		var __orig_bailout *JMP_BUF = zend.__EG().GetBailout()
		var __bailout JMP_BUF
		zend.__EG().SetBailout(&__bailout)
		if zend.SETJMP(__bailout) == 0 {
			zend.ZendHashApply(BG(user_shutdown_function_names), UserShutdownFunctionCall)
		}
		zend.__EG().SetBailout(__orig_bailout)
	}
}
func PhpFreeShutdownFunctions() {
	if BG(user_shutdown_function_names) {
		var __orig_bailout *JMP_BUF = zend.__EG().GetBailout()
		var __bailout JMP_BUF
		zend.__EG().SetBailout(&__bailout)
		if zend.SETJMP(__bailout) == 0 {
			zend.ZendHashDestroy(BG(user_shutdown_function_names))
			zend.FREE_HASHTABLE(BG(user_shutdown_function_names))
			BG(user_shutdown_function_names) = nil
		} else {
			zend.__EG().SetBailout(__orig_bailout)

			/* maybe shutdown method call exit, we just ignore it */

			zend.FREE_HASHTABLE(BG(user_shutdown_function_names))
			BG(user_shutdown_function_names) = nil
		}
		zend.__EG().SetBailout(__orig_bailout)
	}
}
func ZifRegisterShutdownFunction(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var shutdown_function_entry PhpShutdownFunctionEntry
	var i int
	shutdown_function_entry.SetArgCount(zend.ZEND_NUM_ARGS())
	if shutdown_function_entry.GetArgCount() < 1 {
		zend.WRONG_PARAM_COUNT
	}
	shutdown_function_entry.SetArguments((*zend.Zval)(zend.SafeEmalloc(b.SizeOf("zval"), shutdown_function_entry.GetArgCount(), 0)))
	if zend.ZendGetParametersArray(zend.ZEND_NUM_ARGS(), shutdown_function_entry.GetArgCount(), shutdown_function_entry.GetArguments()) == zend.FAILURE {
		zend.Efree(shutdown_function_entry.GetArguments())
		zend.RETVAL_FALSE
		return
	}

	/* Prevent entering of anything but valid callback (syntax check only!) */

	if zend.ZendIsCallable(shutdown_function_entry.GetArguments()[0], 0, nil) == 0 {
		var callback_name *zend.ZendString = zend.ZendGetCallableName(shutdown_function_entry.GetArguments()[0])
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid shutdown callback '%s' passed", callback_name.GetVal())
		zend.Efree(shutdown_function_entry.GetArguments())
		zend.ZendStringReleaseEx(callback_name, 0)
		zend.RETVAL_FALSE
	} else {
		if !(BG(user_shutdown_function_names)) {
			zend.ALLOC_HASHTABLE(BG(user_shutdown_function_names))
			zend.ZendHashInit(BG(user_shutdown_function_names), 0, nil, UserShutdownFunctionDtor, 0)
		}
		for i = 0; i < shutdown_function_entry.GetArgCount(); i++ {
			zend.Z_TRY_ADDREF(shutdown_function_entry.GetArguments()[i])
		}
		zend.ZendHashNextIndexInsertMem(BG(user_shutdown_function_names), &shutdown_function_entry, b.SizeOf("php_shutdown_function_entry"))
	}

	/* Prevent entering of anything but valid callback (syntax check only!) */
}
func RegisterUserShutdownFunction(function_name *byte, function_len int, shutdown_function_entry *PhpShutdownFunctionEntry) zend.ZendBool {
	if !(BG(user_shutdown_function_names)) {
		zend.ALLOC_HASHTABLE(BG(user_shutdown_function_names))
		zend.ZendHashInit(BG(user_shutdown_function_names), 0, nil, UserShutdownFunctionDtor, 0)
	}
	zend.ZendHashStrUpdateMem(BG(user_shutdown_function_names), function_name, function_len, shutdown_function_entry, b.SizeOf("php_shutdown_function_entry"))
	return 1
}
func RemoveUserShutdownFunction(function_name *byte, function_len int) zend.ZendBool {
	if BG(user_shutdown_function_names) {
		return zend.ZendHashStrDel(BG(user_shutdown_function_names), function_name, function_len) != zend.FAILURE
	}
	return 0
}
func AppendUserShutdownFunction(shutdown_function_entry PhpShutdownFunctionEntry) zend.ZendBool {
	if !(BG(user_shutdown_function_names)) {
		zend.ALLOC_HASHTABLE(BG(user_shutdown_function_names))
		zend.ZendHashInit(BG(user_shutdown_function_names), 0, nil, UserShutdownFunctionDtor, 0)
	}
	return zend.ZendHashNextIndexInsertMem(BG(user_shutdown_function_names), &shutdown_function_entry, b.SizeOf("php_shutdown_function_entry")) != nil
}
func PhpGetHighlight(syntax_highlighter_ini *zend.ZendSyntaxHighlighterIni) {
	syntax_highlighter_ini.SetHighlightComment(zend.INI_STR("highlight.comment"))
	syntax_highlighter_ini.SetHighlightDefault(zend.INI_STR("highlight.default"))
	syntax_highlighter_ini.SetHighlightHtml(zend.INI_STR("highlight.html"))
	syntax_highlighter_ini.SetHighlightKeyword(zend.INI_STR("highlight.keyword"))
	syntax_highlighter_ini.SetHighlightString(zend.INI_STR("highlight.string"))
}
func ZifHighlightFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var ret int
	var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
	var i zend.ZendBool = 0
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
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &i, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	if core.PhpCheckOpenBasedir(filename) != 0 {
		zend.RETVAL_FALSE
		return
	}
	if i != 0 {
		core.PhpOutputStartDefault()
	}
	PhpGetHighlight(&syntax_highlighter_ini)
	ret = zend.HighlightFile(filename, &syntax_highlighter_ini)
	if ret == zend.FAILURE {
		if i != 0 {
			core.PhpOutputEnd()
		}
		zend.RETVAL_FALSE
		return
	}
	if i != 0 {
		core.PhpOutputGetContents(return_value)
		core.PhpOutputDiscard()
	} else {
		zend.RETVAL_TRUE
		return
	}
}
func ZifPhpStripWhitespace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var original_lex_state zend.ZendLexState
	var file_handle zend.ZendFileHandle
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
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
	core.PhpOutputStartDefault()
	zend.ZendStreamInitFilename(&file_handle, filename)
	zend.ZendSaveLexicalState(&original_lex_state)
	if zend.OpenFileForScanning(&file_handle) == zend.FAILURE {
		zend.ZendRestoreLexicalState(&original_lex_state)
		core.PhpOutputEnd()
		zend.RETVAL_EMPTY_STRING()
		return
	}
	zend.ZendStrip()
	zend.ZendDestroyFileHandle(&file_handle)
	zend.ZendRestoreLexicalState(&original_lex_state)
	core.PhpOutputGetContents(return_value)
	core.PhpOutputDiscard()
}
func ZifHighlightString(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var expr *zend.Zval
	var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
	var hicompiled_string_description *byte
	var i zend.ZendBool = 0
	var old_error_reporting int = zend.__EG().GetErrorReporting()
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
			zend.ZendParseArgZvalDeref(_arg, &expr, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &i, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	if zend.TryConvertToString(expr) == 0 {
		return
	}
	if i != 0 {
		core.PhpOutputStartDefault()
	}
	zend.__EG().SetErrorReporting(zend.E_ERROR)
	PhpGetHighlight(&syntax_highlighter_ini)
	hicompiled_string_description = zend.ZendMakeCompiledStringDescription("highlighted code")
	if zend.HighlightString(expr, &syntax_highlighter_ini, hicompiled_string_description) == zend.FAILURE {
		zend.Efree(hicompiled_string_description)
		zend.__EG().SetErrorReporting(old_error_reporting)
		if i != 0 {
			core.PhpOutputEnd()
		}
		zend.RETVAL_FALSE
		return
	}
	zend.Efree(hicompiled_string_description)
	zend.__EG().SetErrorReporting(old_error_reporting)
	if i != 0 {
		core.PhpOutputGetContents(return_value)
		core.PhpOutputDiscard()
	} else {
		zend.RETVAL_TRUE
		return
	}
}
func ZifIniGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var varname *zend.ZendString
	var val *zend.ZendString
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
			if zend.ZendParseArgStr(_arg, &varname, 0) == 0 {
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
	val = zend.ZendIniGetValue(varname)
	if val == nil {
		zend.RETVAL_FALSE
		return
	}
	if val.GetLen() == 0 {
		zend.RETVAL_EMPTY_STRING()
	} else if val.GetLen() == 1 {
		zend.RETVAL_INTERNED_STR(zend.ZSTR_CHAR(zend.ZendUchar(val.GetVal()[0])))
	} else if (val.GetGcFlags() & zend.GC_PERSISTENT) == 0 {
		zend.ZVAL_NEW_STR(return_value, val.Copy())
	} else {
		zend.ZVAL_NEW_STR(return_value, zend.ZendStringInit(val.GetVal(), val.GetLen(), 0))
	}
}
func ZifIniGetAll(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var extname *byte = nil
	var extname_len int = 0
	var module_number int = 0
	var module *zend.ZendModuleEntry
	var details zend.ZendBool = 1
	var key *zend.ZendString
	var ini_entry *zend.ZendIniEntry
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &extname, &extname_len, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &details, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	zend.ZendIniSortEntries()
	if extname != nil {
		if b.Assign(&module, zend.ZendHashStrFindPtr(&zend.ModuleRegistry, extname, extname_len)) == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to find extension '%s'", extname)
			zend.RETVAL_FALSE
			return
		}
		module_number = module.GetModuleNumber()
	}
	zend.ArrayInit(return_value)
	var __ht *zend.HashTable = zend.__EG().GetIniDirectives()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		key = _p.GetKey()
		ini_entry = _z.GetPtr()
		var option zend.Zval
		if module_number != 0 && ini_entry.GetModuleNumber() != module_number {
			continue
		}
		if key == nil || key.GetVal()[0] != 0 {
			if details != 0 {
				zend.ArrayInit(&option)
				if ini_entry.GetOrigValue() != nil {
					zend.AddAssocStr(&option, "global_value", ini_entry.GetOrigValue().Copy())
				} else if ini_entry.GetValue() != nil {
					zend.AddAssocStr(&option, "global_value", ini_entry.GetValue().Copy())
				} else {
					zend.AddAssocNull(&option, "global_value")
				}
				if ini_entry.GetValue() != nil {
					zend.AddAssocStr(&option, "local_value", ini_entry.GetValue().Copy())
				} else {
					zend.AddAssocNull(&option, "local_value")
				}
				zend.AddAssocLong(&option, "access", ini_entry.GetModifiable())
				zend.ZendSymtableUpdate(return_value.GetArr(), ini_entry.GetName(), &option)
			} else {
				if ini_entry.GetValue() != nil {
					var zv zend.Zval
					zend.ZVAL_STR_COPY(&zv, ini_entry.GetValue())
					zend.ZendSymtableUpdate(return_value.GetArr(), ini_entry.GetName(), &zv)
				} else {
					zend.ZendSymtableUpdate(return_value.GetArr(), ini_entry.GetName(), zend.__EG().GetUninitializedZval())
				}
			}
		}
	}
}
func PhpIniCheckPath(option_name *byte, option_len int, new_option_name string, new_option_len int) int {
	if option_len+1 != new_option_len {
		return 0
	}
	return !(strncmp(option_name, new_option_name, option_len))
}
func ZifIniSet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var varname *zend.ZendString
	var new_value *zend.ZendString
	var val *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.ZendParseArgStr(_arg, &varname, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &new_value, 0) == 0 {
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
	val = zend.ZendIniGetValue(varname)

	/* copy to return here, because alter might free it! */

	if val != nil {
		if val.GetLen() == 0 {
			zend.RETVAL_EMPTY_STRING()
		} else if val.GetLen() == 1 {
			zend.RETVAL_INTERNED_STR(zend.ZSTR_CHAR(zend.ZendUchar(val.GetVal()[0])))
		} else if (val.GetGcFlags() & zend.GC_PERSISTENT) == 0 {
			zend.ZVAL_NEW_STR(return_value, val.Copy())
		} else {
			zend.ZVAL_NEW_STR(return_value, zend.ZendStringInit(val.GetVal(), val.GetLen(), 0))
		}
	} else {
		zend.RETVAL_FALSE
	}

	// #define _CHECK_PATH(var,var_len,ini) php_ini_check_path ( var , var_len , ini , sizeof ( ini ) )

	/* open basedir check */

	if core.PG(open_basedir) {
		if PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "error_log", b.SizeOf("\"error_log\"")) != 0 || PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "java.class.path", b.SizeOf("\"java.class.path\"")) != 0 || PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "java.home", b.SizeOf("\"java.home\"")) != 0 || PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "mail.log", b.SizeOf("\"mail.log\"")) != 0 || PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "java.library.path", b.SizeOf("\"java.library.path\"")) != 0 || PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "vpopmail.directory", b.SizeOf("\"vpopmail.directory\"")) != 0 {
			if core.PhpCheckOpenBasedir(new_value.GetVal()) != 0 {
				zend.ZvalPtrDtorStr(return_value)
				zend.RETVAL_FALSE
				return
			}
		}
	}
	if zend.ZendAlterIniEntryEx(varname, new_value, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0) == zend.FAILURE {
		zend.ZvalPtrDtorStr(return_value)
		zend.RETVAL_FALSE
		return
	}
}
func ZifIniRestore(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var varname *zend.ZendString
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
			if zend.ZendParseArgStr(_arg, &varname, 0) == 0 {
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
	zend.ZendRestoreIniEntry(varname, core.PHP_INI_STAGE_RUNTIME)
}
func ZifSetIncludePath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var new_value *zend.ZendString
	var old_value *byte
	var key *zend.ZendString
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
			if zend.ZendParseArgPathStr(_arg, &new_value, 0) == 0 {
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
			return
		}
		break
	}
	old_value = zend.ZendIniString("include_path", b.SizeOf("\"include_path\"")-1, 0)

	/* copy to return here, because alter might free it! */

	if old_value != nil {
		zend.RETVAL_STRING(old_value)
	} else {
		zend.RETVAL_FALSE
	}
	key = zend.ZendStringInit("include_path", b.SizeOf("\"include_path\"")-1, 0)
	if zend.ZendAlterIniEntryEx(key, new_value, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0) == zend.FAILURE {
		zend.ZendStringReleaseEx(key, 0)
		zend.ZvalPtrDtorStr(return_value)
		zend.RETVAL_FALSE
		return
	}
	zend.ZendStringReleaseEx(key, 0)
}
func ZifGetIncludePath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	str = zend.ZendIniString("include_path", b.SizeOf("\"include_path\"")-1, 0)
	if str == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRING(str)
	return
}
func ZifRestoreIncludePath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var key *zend.ZendString
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	key = zend.ZendStringInit("include_path", b.SizeOf("\"include_path\"")-1, 0)
	zend.ZendRestoreIniEntry(key, core.PHP_INI_STAGE_RUNTIME)
	zend.ZendStringEfree(key)
}
func ZifPrintR(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var var_ *zend.Zval
	var do_return zend.ZendBool = 0
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
			zend.ZendParseArgZvalDeref(_arg, &var_, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &do_return, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	if do_return != 0 {
		zend.RETVAL_STR(zend.ZendPrintZvalRToStr(var_, 0))
		return
	} else {
		zend.ZendPrintZvalR(var_, 0)
		zend.RETVAL_TRUE
		return
	}
}
func ZifConnectionAborted(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	zend.RETVAL_LONG(core.PG(connection_status) & core.PHP_CONNECTION_ABORTED)
	return
}
func ZifConnectionStatus(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	zend.RETVAL_LONG(core.PG(connection_status))
	return
}
func ZifIgnoreUserAbort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg zend.ZendBool = 0
	var old_setting int
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
			if zend.ZendParseArgBool(_arg, &arg, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	old_setting = uint16(core.PG(ignore_user_abort))
	if zend.ZEND_NUM_ARGS() != 0 {
		var key *zend.ZendString = zend.ZendStringInit("ignore_user_abort", b.SizeOf("\"ignore_user_abort\"")-1, 0)
		zend.ZendAlterIniEntryChars(key, b.Cond(arg != 0, "1", "0"), 1, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME)
		zend.ZendStringReleaseEx(key, 0)
	}
	zend.RETVAL_LONG(old_setting)
	return
}
func ZifGetservbyname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var name *byte
	var proto *byte
	var name_len int
	var proto_len int
	var serv *__struct__servent
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.ZendParseArgString(_arg, &name, &name_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &proto, &proto_len, 0) == 0 {
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

	/* empty string behaves like NULL on windows implementation of
	   getservbyname. Let be portable instead. */

	serv = getservbyname(name, proto)
	if serv == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(ntohs(serv.s_port))
	return
}
func ZifGetservbyport(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var proto *byte
	var proto_len int
	var port zend.ZendLong
	var serv *__struct__servent
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.ZendParseArgLong(_arg, &port, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &proto, &proto_len, 0) == 0 {
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
	serv = getservbyport(htons(uint16(port)), proto)
	if serv == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRING(serv.s_name)
	return
}
func ZifGetprotobyname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var name *byte
	var name_len int
	var ent *__struct__protoent
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
			if zend.ZendParseArgString(_arg, &name, &name_len, 0) == 0 {
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
	ent = getprotobyname(name)
	if ent == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(ent.p_proto)
	return
}
func ZifGetprotobynumber(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var proto zend.ZendLong
	var ent *__struct__protoent
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
			if zend.ZendParseArgLong(_arg, &proto, &_dummy, 0, 0) == 0 {
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
	ent = getprotobynumber(int(proto))
	if ent == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRING(ent.p_name)
	return
}
func ZifRegisterTickFunction(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var tick_fe UserTickFunctionEntry
	var i int
	var function_name *zend.ZendString = nil
	tick_fe.SetCalling(0)
	tick_fe.SetArgCount(zend.ZEND_NUM_ARGS())
	if tick_fe.GetArgCount() < 1 {
		zend.WRONG_PARAM_COUNT
	}
	tick_fe.SetArguments((*zend.Zval)(zend.SafeEmalloc(b.SizeOf("zval"), tick_fe.GetArgCount(), 0)))
	if zend.ZendGetParametersArray(zend.ZEND_NUM_ARGS(), tick_fe.GetArgCount(), tick_fe.GetArguments()) == zend.FAILURE {
		zend.Efree(tick_fe.GetArguments())
		zend.RETVAL_FALSE
		return
	}
	if zend.ZendIsCallable(tick_fe.GetArguments()[0], 0, &function_name) == 0 {
		zend.Efree(tick_fe.GetArguments())
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid tick callback '%s' passed", function_name.GetVal())
		zend.ZendStringReleaseEx(function_name, 0)
		zend.RETVAL_FALSE
		return
	} else if function_name != nil {
		zend.ZendStringReleaseEx(function_name, 0)
	}
	if tick_fe.GetArguments()[0].GetType() != zend.IS_ARRAY && tick_fe.GetArguments()[0].GetType() != zend.IS_OBJECT {
		zend.ConvertToStringEx(tick_fe.GetArguments()[0])
	}
	if !(BG(user_tick_functions)) {
		BG(user_tick_functions) = (*zend.ZendLlist)(zend.Emalloc(b.SizeOf("zend_llist")))
		zend.ZendLlistInit(BG(user_tick_functions), b.SizeOf("user_tick_function_entry"), zend.LlistDtorFuncT(UserTickFunctionDtor), 0)
		core.PhpAddTickFunction(RunUserTickFunctions, nil)
	}
	for i = 0; i < tick_fe.GetArgCount(); i++ {
		zend.Z_TRY_ADDREF(tick_fe.GetArguments()[i])
	}
	zend.ZendLlistAddElement(BG(user_tick_functions), &tick_fe)
	zend.RETVAL_TRUE
	return
}
func ZifUnregisterTickFunction(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var function *zend.Zval
	var tick_fe UserTickFunctionEntry
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
			zend.ZendParseArgZvalDeref(_arg, &function, 0)
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
	if !(BG(user_tick_functions)) {
		return
	}
	if function.GetType() != zend.IS_ARRAY && function.GetType() != zend.IS_OBJECT {
		zend.ConvertToString(function)
	}
	tick_fe.SetArguments((*zend.Zval)(zend.Emalloc(b.SizeOf("zval"))))
	zend.ZVAL_COPY_VALUE(tick_fe.GetArguments()[0], function)
	tick_fe.SetArgCount(1)
	zend.ZendLlistDelElement(BG(user_tick_functions), &tick_fe, (func(any, any) int)(UserTickFunctionCompare))
	zend.Efree(tick_fe.GetArguments())
}
func ZifIsUploadedFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var path *byte
	var path_len int
	if !(core.SG(rfc1867_uploaded_files)) {
		zend.RETVAL_FALSE
		return
	}
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
			if zend.ZendParseArgString(_arg, &path, &path_len, 0) == 0 {
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
	if zend.ZendHashStrExists(core.SG(rfc1867_uploaded_files), path, path_len) != 0 {
		zend.RETVAL_TRUE
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifMoveUploadedFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var path *byte
	var new_path *byte
	var path_len int
	var new_path_len int
	var successful zend.ZendBool = 0
	var oldmask int
	var ret int
	if !(core.SG(rfc1867_uploaded_files)) {
		zend.RETVAL_FALSE
		return
	}
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.ZendParseArgString(_arg, &path, &path_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &new_path, &new_path_len, 0) == 0 {
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
			return
		}
		break
	}
	if zend.ZendHashStrExists(core.SG(rfc1867_uploaded_files), path, path_len) == 0 {
		zend.RETVAL_FALSE
		return
	}
	if core.PhpCheckOpenBasedir(new_path) != 0 {
		zend.RETVAL_FALSE
		return
	}
	if zend.VCWD_RENAME(path, new_path) == 0 {
		successful = 1
		oldmask = umask(077)
		umask(oldmask)
		ret = zend.VCWD_CHMOD(new_path, 0666 & ^oldmask)
		if ret == -1 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s", strerror(errno))
		}
	} else if PhpCopyFileEx(path, new_path, core.STREAM_DISABLE_OPEN_BASEDIR) == zend.SUCCESS {
		zend.VCWD_UNLINK(path)
		successful = 1
	}
	if successful != 0 {
		zend.ZendHashStrDel(core.SG(rfc1867_uploaded_files), path, path_len)
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to move '%s' to '%s'", path, new_path)
	}
	zend.RETVAL_BOOL(successful != 0)
	return
}
func PhpSimpleIniParserCb(arg1 *zend.Zval, arg2 *zend.Zval, arg3 *zend.Zval, callback_type int, arr *zend.Zval) {
	switch callback_type {
	case zend.ZEND_INI_PARSER_ENTRY:
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}
		zend.Z_TRY_ADDREF_P(arg2)
		zend.ZendSymtableUpdate(arr.GetArr(), arg1.GetStr(), arg2)
		break
	case zend.ZEND_INI_PARSER_POP_ENTRY:
		var hash zend.Zval
		var find_hash *zend.Zval
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}
		if !(zend.Z_STRLEN_P(arg1) > 1 && zend.Z_STRVAL_P(arg1)[0] == '0') && zend.IsNumericString(zend.Z_STRVAL_P(arg1), zend.Z_STRLEN_P(arg1), nil, nil, 0) == zend.IS_LONG {
			var key zend.ZendUlong = zend.ZendUlong(zend.ZendAtol(zend.Z_STRVAL_P(arg1), zend.Z_STRLEN_P(arg1)))
			if b.Assign(&find_hash, zend.ZendHashIndexFind(arr.GetArr(), key)) == nil {
				zend.ArrayInit(&hash)
				find_hash = arr.GetArr().IndexAddNewH(key, &hash)
			}
		} else {
			if b.Assign(&find_hash, arr.GetArr().FindByZendString(arg1.GetStr())) == nil {
				zend.ArrayInit(&hash)
				find_hash = zend.ZendHashAddNew(arr.GetArr(), arg1.GetStr().GetStr(), &hash)
			}
		}
		if find_hash.GetType() != zend.IS_ARRAY {
			zend.ZvalPtrDtorNogc(find_hash)
			zend.ArrayInit(find_hash)
		}
		if arg3 == nil || arg3.IsType(zend.IS_STRING) && zend.Z_STRLEN_P(arg3) == 0 {
			zend.Z_TRY_ADDREF_P(arg2)
			zend.AddNextIndexZval(find_hash, arg2)
		} else {
			zend.ArraySetZvalKey(find_hash.GetArr(), arg3, arg2)
		}
		break
	case zend.ZEND_INI_PARSER_SECTION:
		break
	}
}
func PhpIniParserCbWithSections(arg1 *zend.Zval, arg2 *zend.Zval, arg3 *zend.Zval, callback_type int, arr *zend.Zval) {
	if callback_type == zend.ZEND_INI_PARSER_SECTION {
		zend.ArrayInit(&(BG(active_ini_file_section)))
		zend.ZendSymtableUpdate(arr.GetArr(), arg1.GetStr(), &(BG(active_ini_file_section)))
	} else if arg2 != nil {
		var active_arr *zend.Zval
		if BG(active_ini_file_section).u1.v.type_ != zend.IS_UNDEF {
			active_arr = &(BG(active_ini_file_section))
		} else {
			active_arr = arr
		}
		PhpSimpleIniParserCb(arg1, arg2, arg3, callback_type, active_arr)
	}
}
func ZifParseIniFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte = nil
	var filename_len int = 0
	var process_sections zend.ZendBool = 0
	var scanner_mode zend.ZendLong = zend.ZEND_INI_SCANNER_NORMAL
	var fh zend.ZendFileHandle
	var ini_parser_cb zend.ZendIniParserCbT
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
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &process_sections, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &scanner_mode, &_dummy, 0, 0) == 0 {
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
	if filename_len == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Filename cannot be empty!")
		zend.RETVAL_FALSE
		return
	}

	/* Set callback function */

	if process_sections != 0 {
		zend.ZVAL_UNDEF(&(BG(active_ini_file_section)))
		ini_parser_cb = zend.ZendIniParserCbT(PhpIniParserCbWithSections)
	} else {
		ini_parser_cb = zend.ZendIniParserCbT(PhpSimpleIniParserCb)
	}

	/* Setup filehandle */

	zend.ZendStreamInitFilename(&fh, filename)
	zend.ArrayInit(return_value)
	if zend.ZendParseIniFile(&fh, 0, int(scanner_mode), ini_parser_cb, return_value) == zend.FAILURE {
		zend.ZendArrayDestroy(return_value.GetArr())
		zend.RETVAL_FALSE
		return
	}
}
func ZifParseIniString(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var string *byte = nil
	var str *byte = nil
	var str_len int = 0
	var process_sections zend.ZendBool = 0
	var scanner_mode zend.ZendLong = zend.ZEND_INI_SCANNER_NORMAL
	var ini_parser_cb zend.ZendIniParserCbT
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
			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &process_sections, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &scanner_mode, &_dummy, 0, 0) == 0 {
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
	if core.INT_MAX-str_len < zend.ZEND_MMAP_AHEAD {
		zend.RETVAL_FALSE
	}

	/* Set callback function */

	if process_sections != 0 {
		zend.ZVAL_UNDEF(&(BG(active_ini_file_section)))
		ini_parser_cb = zend.ZendIniParserCbT(PhpIniParserCbWithSections)
	} else {
		ini_parser_cb = zend.ZendIniParserCbT(PhpSimpleIniParserCb)
	}

	/* Setup string */

	string = (*byte)(zend.Emalloc(str_len + zend.ZEND_MMAP_AHEAD))
	memcpy(string, str, str_len)
	memset(string+str_len, 0, zend.ZEND_MMAP_AHEAD)
	zend.ArrayInit(return_value)
	if zend.ZendParseIniString(string, 0, int(scanner_mode), ini_parser_cb, return_value) == zend.FAILURE {
		zend.ZendArrayDestroy(return_value.GetArr())
		zend.RETVAL_FALSE
	}
	zend.Efree(string)
}
func ZifSysGetloadavg(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var load []float64
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if getloadavg(load, 3) == -1 {
		zend.RETVAL_FALSE
		return
	} else {
		zend.ArrayInit(return_value)
		zend.AddIndexDouble(return_value, 0, load[0])
		zend.AddIndexDouble(return_value, 1, load[1])
		zend.AddIndexDouble(return_value, 2, load[2])
	}
}
