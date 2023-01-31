// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func ZmStartupCore(type_ int, module_number int) int {
	var class_entry ZendClassEntry
	memset(&class_entry, 0, b.SizeOf("zend_class_entry"))
	class_entry.SetName(ZendStringInitInterned("stdClass", b.SizeOf("\"stdClass\"")-1, 1))
	class_entry.SetBuiltinFunctions(nil)
	ZendStandardClassDef = ZendRegisterInternalClass(&class_entry)
	ZendRegisterDefaultClasses()
	return SUCCESS
}
func ZendStartupBuiltinFunctions() int {
	ZendBuiltinModule.SetModuleNumber(0)
	ZendBuiltinModule.SetType(MODULE_PERSISTENT)
	if b.Assign(&(EG__().GetCurrentModule()), ZendRegisterModuleEx(&ZendBuiltinModule)) == nil {
		return FAILURE
	} else {
		return SUCCESS
	}
}
func ZifZendVersion(execute_data *ZendExecuteData, return_value *Zval) {
	if ZendParseParametersNone() == FAILURE {
		return
	}
	RETVAL_STRINGL(ZEND_VERSION, b.SizeOf("ZEND_VERSION")-1)
	return
}
func ZifGcMemCaches(execute_data *ZendExecuteData, return_value *Zval) {
	if ZendParseParametersNone() == FAILURE {
		return
	}
	RETVAL_LONG(ZendMmGc(ZendMmGetHeap()))
	return
}
func ZifGcCollectCycles(execute_data *ZendExecuteData, return_value *Zval) {
	if ZendParseParametersNone() == FAILURE {
		return
	}
	RETVAL_LONG(GcCollectCycles())
	return
}
func ZifGcEnabled(execute_data *ZendExecuteData, return_value *Zval) {
	if ZendParseParametersNone() == FAILURE {
		return
	}
	RETVAL_BOOL(GcEnabled() != 0)
	return
}
func ZifGcEnable(execute_data *ZendExecuteData, return_value *Zval) {
	var key *ZendString
	if ZendParseParametersNone() == FAILURE {
		return
	}
	key = ZendStringInit("zend.enable_gc", b.SizeOf("\"zend.enable_gc\"")-1, 0)
	ZendAlterIniEntryChars(key, "1", b.SizeOf("\"1\"")-1, ZEND_INI_USER, ZEND_INI_STAGE_RUNTIME)
	ZendStringReleaseEx(key, 0)
}
func ZifGcDisable(execute_data *ZendExecuteData, return_value *Zval) {
	var key *ZendString
	if ZendParseParametersNone() == FAILURE {
		return
	}
	key = ZendStringInit("zend.enable_gc", b.SizeOf("\"zend.enable_gc\"")-1, 0)
	ZendAlterIniEntryChars(key, "0", b.SizeOf("\"0\"")-1, ZEND_INI_USER, ZEND_INI_STAGE_RUNTIME)
	ZendStringReleaseEx(key, 0)
}
func ZifGcStatus(execute_data *ZendExecuteData, return_value *Zval) {
	var status ZendGcStatus
	if ZendParseParametersNone() == FAILURE {
		return
	}
	ZendGcGetStatus(&status)
	ArrayInitSize(return_value, 3)
	AddAssocLongEx(return_value, "runs", b.SizeOf("\"runs\"")-1, long(status.GetRuns()))
	AddAssocLongEx(return_value, "collected", b.SizeOf("\"collected\"")-1, long(status.GetCollected()))
	AddAssocLongEx(return_value, "threshold", b.SizeOf("\"threshold\"")-1, long(status.GetThreshold()))
	AddAssocLongEx(return_value, "roots", b.SizeOf("\"roots\"")-1, long(status.GetNumRoots()))
}
func ZifFuncNumArgs(execute_data *ZendExecuteData, return_value *Zval) {
	var ex *ZendExecuteData = EX(prev_execute_data)
	if ZendParseParametersNone() == FAILURE {
		return
	}
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_CODE) != 0 {
		ZendError(E_WARNING, "func_num_args():  Called from the global scope - no function context")
		RETVAL_LONG(-1)
		return
	}
	if ZendForbidDynamicCall("func_num_args()") == FAILURE {
		RETVAL_LONG(-1)
		return
	}
	RETVAL_LONG(ZEND_CALL_NUM_ARGS(ex))
	return
}
func ZifFuncGetArg(execute_data *ZendExecuteData, return_value *Zval) {
	var arg_count uint32
	var first_extra_arg uint32
	var arg *Zval
	var requested_offset ZendLong
	var ex *ZendExecuteData
	if ZendParseParameters(ZEND_NUM_ARGS(), "l", &requested_offset) == FAILURE {
		return
	}
	if requested_offset < 0 {
		ZendError(E_WARNING, "func_get_arg():  The argument number should be >= 0")
		RETVAL_FALSE
		return
	}
	ex = EX(prev_execute_data)
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_CODE) != 0 {
		ZendError(E_WARNING, "func_get_arg():  Called from the global scope - no function context")
		RETVAL_FALSE
		return
	}
	if ZendForbidDynamicCall("func_get_arg()") == FAILURE {
		RETVAL_FALSE
		return
	}
	arg_count = ZEND_CALL_NUM_ARGS(ex)
	if ZendUlong(requested_offset >= arg_count) != 0 {
		ZendError(E_WARNING, "func_get_arg():  Argument "+ZEND_LONG_FMT+" not passed to function", requested_offset)
		RETVAL_FALSE
		return
	}
	first_extra_arg = ex.GetFunc().GetOpArray().GetNumArgs()
	if ZendUlong(requested_offset >= first_extra_arg && ZEND_CALL_NUM_ARGS(ex) > first_extra_arg) != 0 {
		arg = ZEND_CALL_VAR_NUM(ex, ex.GetFunc().GetOpArray().GetLastVar()+ex.GetFunc().GetOpArray().GetT()) + (requested_offset - first_extra_arg)
	} else {
		arg = ZEND_CALL_ARG(ex, requested_offset+1)
	}
	if !(arg.IsUndef()) {
		ZVAL_COPY_DEREF(return_value, arg)
	}
}
func ZifFuncGetArgs(execute_data *ZendExecuteData, return_value *Zval) {
	var p *Zval
	var q *Zval
	var arg_count uint32
	var first_extra_arg uint32
	var i uint32
	var ex *ZendExecuteData = EX(prev_execute_data)
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_CODE) != 0 {
		ZendError(E_WARNING, "func_get_args():  Called from the global scope - no function context")
		RETVAL_FALSE
		return
	}
	if ZendForbidDynamicCall("func_get_args()") == FAILURE {
		RETVAL_FALSE
		return
	}
	arg_count = ZEND_CALL_NUM_ARGS(ex)
	if arg_count != 0 {
		ArrayInitSize(return_value, arg_count)
		first_extra_arg = ex.GetFunc().GetOpArray().GetNumArgs()
		ZendHashRealInitPacked(return_value.GetArr())
		var __fill_ht *HashTable = return_value.GetArr()
		var __fill_bkt *Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		ZEND_ASSERT(__fill_ht.HasUFlags(HASH_FLAG_PACKED))
		i = 0
		p = ZEND_CALL_ARG(ex, 1)
		if arg_count > first_extra_arg {
			for i < first_extra_arg {
				q = p
				if q.GetTypeInfo() != IS_UNDEF {
					ZVAL_DEREF(q)
					if Z_OPT_REFCOUNTED_P(q) {
						Z_ADDREF_P(q)
					}
					ZVAL_COPY_VALUE(__fill_bkt.GetVal(), q)
				} else {
					__fill_bkt.GetVal().SetNull()
				}
				__fill_bkt.SetH(__fill_idx)
				__fill_bkt.SetKey(nil)
				__fill_bkt++
				__fill_idx++
				p++
				i++
			}
			p = ZEND_CALL_VAR_NUM(ex, ex.GetFunc().GetOpArray().GetLastVar()+ex.GetFunc().GetOpArray().GetT())
		}
		for i < arg_count {
			q = p
			if q.GetTypeInfo() != IS_UNDEF {
				ZVAL_DEREF(q)
				if Z_OPT_REFCOUNTED_P(q) {
					Z_ADDREF_P(q)
				}
				ZVAL_COPY_VALUE(__fill_bkt.GetVal(), q)
			} else {
				__fill_bkt.GetVal().SetNull()
			}
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
			p++
			i++
		}
		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
		Z_ARRVAL_P(return_value).SetNNumOfElements(arg_count)
	} else {
		RETVAL_EMPTY_ARRAY()
		return
	}
}
func ZifStrlen(execute_data *ZendExecuteData, return_value *Zval) {
	var s *ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &s, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	RETVAL_LONG(s.GetLen())
}
func ZifStrcmp(execute_data *ZendExecuteData, return_value *Zval) {
	var s1 *ZendString
	var s2 *ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	RETVAL_LONG(ZendBinaryStrcmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen()))
	return
}
func ZifStrncmp(execute_data *ZendExecuteData, return_value *Zval) {
	var s1 *ZendString
	var s2 *ZendString
	var len_ ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgLong(_arg, &len_, &_dummy, 0, 0) == 0 {
				_expected_type = Z_EXPECTED_LONG
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if len_ < 0 {
		ZendError(E_WARNING, "Length must be greater than or equal to 0")
		RETVAL_FALSE
		return
	}
	RETVAL_LONG(ZendBinaryStrncmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen(), len_))
	return
}
func ZifStrcasecmp(execute_data *ZendExecuteData, return_value *Zval) {
	var s1 *ZendString
	var s2 *ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	RETVAL_LONG(ZendBinaryStrcasecmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen()))
	return
}
func ZifStrncasecmp(execute_data *ZendExecuteData, return_value *Zval) {
	var s1 *ZendString
	var s2 *ZendString
	var len_ ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgLong(_arg, &len_, &_dummy, 0, 0) == 0 {
				_expected_type = Z_EXPECTED_LONG
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if len_ < 0 {
		ZendError(E_WARNING, "Length must be greater than or equal to 0")
		RETVAL_FALSE
		return
	}
	RETVAL_LONG(ZendBinaryStrncasecmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen(), len_))
	return
}
func ZifEach(execute_data *ZendExecuteData, return_value *Zval) {
	var array *Zval
	var entry *Zval
	var tmp Zval
	var num_key ZendUlong
	var target_hash *HashTable
	var key *ZendString
	if ZendParseParameters(ZEND_NUM_ARGS(), "z/", &array) == FAILURE {
		return
	}
	if EG__().GetEachDeprecationThrown() == 0 {
		ZendError(E_DEPRECATED, "The each() function is deprecated. This message will be suppressed on further calls")
		EG__().SetEachDeprecationThrown(1)
	}
	target_hash = HASH_OF(array)
	if target_hash == nil {
		ZendError(E_WARNING, "Variable passed to each() is not an array or object")
		return
	}
	for true {
		entry = ZendHashGetCurrentData(target_hash)
		if entry == nil {
			RETVAL_FALSE
			return
		} else if entry.IsIndirect() {
			entry = entry.GetZv()
			if entry.IsUndef() {
				ZendHashMoveForward(target_hash)
				continue
			}
		}
		break
	}
	ArrayInitSize(return_value, 4)
	ZendHashRealInitMixed(return_value.GetArr())

	/* add value elements */

	ZVAL_DEREF(entry)
	if Z_REFCOUNTED_P(entry) {
		entry.GetCounted().AddRefcountEx(2)
	}
	return_value.GetArr().IndexAddNewH(1, entry)
	return_value.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_VALUE).GetStr(), entry)

	/* add the key elements */

	if ZendHashGetCurrentKey(target_hash, &key, &num_key) == HASH_KEY_IS_STRING {
		ZVAL_STR_COPY(&tmp, key)
		Z_TRY_ADDREF(tmp)
	} else {
		tmp.SetLong(num_key)
	}
	return_value.GetArr().IndexAddNewH(0, &tmp)
	return_value.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_KEY).GetStr(), &tmp)
	ZendHashMoveForward(target_hash)
}
func ZifErrorReporting(execute_data *ZendExecuteData, return_value *Zval) {
	var err *Zval = nil
	var old_error_reporting int
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			_optional = 1
			Z_PARAM_PROLOGUE(0, 0)
			ZendParseArgZvalDeref(_arg, &err, 0)
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	old_error_reporting = EG__().GetErrorReporting()
	if ZEND_NUM_ARGS() != 0 {
		var new_val *ZendString = ZvalTryGetString(err)
		if new_val == nil {
			return
		}
		for {
			var p *ZendIniEntry = EG__().GetErrorReportingIniEntry()
			if p == nil {
				var zv *Zval = EG__().GetIniDirectives().KeyFind(ZSTR_KNOWN(ZEND_STR_ERROR_REPORTING).GetStr())
				if zv != nil {
					EG__().SetErrorReportingIniEntry((*ZendIniEntry)(zv.GetPtr()))
					p = EG__().GetErrorReportingIniEntry()
				} else {
					break
				}
			}
			if p.GetModified() == 0 {
				if EG__().GetModifiedIniDirectives() == nil {
					ALLOC_HASHTABLE(EG__().GetModifiedIniDirectives())
					ZendHashInit(EG__().GetModifiedIniDirectives(), 8, nil, nil, 0)
				}
				if ZendHashAddPtr(EG__().GetModifiedIniDirectives(), ZSTR_KNOWN(ZEND_STR_ERROR_REPORTING), p) != nil {
					p.SetOrigValue(p.GetValue())
					p.SetOrigModifiable(p.GetModifiable())
					p.SetModified(1)
				}
			} else if p.GetOrigValue() != p.GetValue() {
				ZendStringReleaseEx(p.GetValue(), 0)
			}
			p.SetValue(new_val)
			if err.IsLong() {
				EG__().SetErrorReporting(err.GetLval())
			} else {
				EG__().SetErrorReporting(atoi(p.GetValue().GetVal()))
			}
			break
		}
	}
	RETVAL_LONG(old_error_reporting)
}
func ValidateConstantArray(ht *HashTable) int {
	var ret int = 1
	var val *Zval
	ht.ProtectRecursive()
	var __ht *HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.GetZv()
			if _z.IsUndef() {
				continue
			}
		}
		val = _z
		ZVAL_DEREF(val)
		if Z_REFCOUNTED_P(val) {
			if val.IsArray() {
				if Z_REFCOUNTED_P(val) {
					if Z_IS_RECURSIVE_P(val) {
						ZendError(E_WARNING, "Constants cannot be recursive arrays")
						ret = 0
						break
					} else if ValidateConstantArray(val.GetArr()) == 0 {
						ret = 0
						break
					}
				}
			} else if val.GetType() != IS_STRING && val.GetType() != IS_RESOURCE {
				ZendError(E_WARNING, "Constants may only evaluate to scalar values, arrays or resources")
				ret = 0
				break
			}
		}
	}
	ht.UnprotectRecursive()
	return ret
}
func CopyConstantArray(dst *Zval, src *Zval) {
	var key *ZendString
	var idx ZendUlong
	var new_val *Zval
	var val *Zval
	ArrayInitSize(dst, Z_ARRVAL_P(src).GetNNumOfElements())
	var __ht *HashTable = src.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.GetZv()
			if _z.IsUndef() {
				continue
			}
		}
		idx = _p.GetH()
		key = _p.GetKey()
		val = _z

		/* constant arrays can't contain references */

		ZVAL_DEREF(val)
		if key != nil {
			new_val = dst.GetArr().KeyAddNew(key.GetStr(), val)
		} else {
			new_val = dst.GetArr().IndexAddNewH(idx, val)
		}
		if val.IsArray() {
			if Z_REFCOUNTED_P(val) {
				CopyConstantArray(new_val, val)
			}
		} else {
			Z_TRY_ADDREF_P(val)
		}
	}
}
func ZifDefine(execute_data *ZendExecuteData, return_value *Zval) {
	var name *ZendString
	var val *Zval
	var val_free Zval
	var non_cs ZendBool = 0
	var case_sensitive int = CONST_CS
	var c ZendConstant
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			Z_PARAM_PROLOGUE(0, 0)
			ZendParseArgZvalDeref(_arg, &val, 0)
			_optional = 1
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgBool(_arg, &non_cs, &_dummy, 0) == 0 {
				_expected_type = Z_EXPECTED_BOOL
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if non_cs != 0 {
		case_sensitive = 0
	}
	if ZendMemnstr(name.GetVal(), "::", b.SizeOf("\"::\"")-1, name.GetVal()+name.GetLen()) != nil {
		ZendError(E_WARNING, "Class constants cannot be defined or redefined")
		RETVAL_FALSE
		return
	}
	val_free.SetUndef()
repeat:
	switch val.GetType() {
	case IS_LONG:

	case IS_DOUBLE:

	case IS_STRING:

	case IS_FALSE:

	case IS_TRUE:

	case IS_NULL:

	case IS_RESOURCE:
		break
	case IS_ARRAY:
		if Z_REFCOUNTED_P(val) {
			if ValidateConstantArray(val.GetArr()) == 0 {
				RETVAL_FALSE
				return
			} else {
				CopyConstantArray(c.GetValue(), val)
				goto register_constant
			}
		}
		break
	case IS_OBJECT:
		if val_free.IsUndef() {
			if Z_OBJ_HT_P(val).GetGet() != nil {
				val = Z_OBJ_HT_P(val).GetGet()(val, &val_free)
				goto repeat
			} else if Z_OBJ_HT_P(val).GetCastObject() != nil {
				if Z_OBJ_HT_P(val).GetCastObject()(val, &val_free, IS_STRING) == SUCCESS {
					val = &val_free
					break
				}
			}
		}
	default:
		ZendError(E_WARNING, "Constants may only evaluate to scalar values, arrays or resources")
		ZvalPtrDtor(&val_free)
		RETVAL_FALSE
		return
	}
	ZVAL_COPY(c.GetValue(), val)
	ZvalPtrDtor(&val_free)
register_constant:
	if non_cs != 0 {
		ZendError(E_DEPRECATED, "define(): Declaration of case-insensitive constants is deprecated")
	}

	/* non persistent */

	ZEND_CONSTANT_SET_FLAGS(&c, case_sensitive, PHP_USER_CONSTANT)
	c.SetName(name.Copy())
	if ZendRegisterConstant(&c) == SUCCESS {
		RETVAL_TRUE
		return
	} else {
		RETVAL_FALSE
		return
	}
}
func ZifDefined(execute_data *ZendExecuteData, return_value *Zval) {
	var name *ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if ZendGetConstantEx(name, ZendGetExecutedScope(), ZEND_FETCH_CLASS_SILENT|ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK) != nil {
		RETVAL_TRUE
		return
	} else {
		RETVAL_FALSE
		return
	}
}
func ZifGetClass(execute_data *ZendExecuteData, return_value *Zval) {
	var obj *Zval = nil
	if ZendParseParameters(ZEND_NUM_ARGS(), "|o", &obj) == FAILURE {
		RETVAL_FALSE
		return
	}
	if obj == nil {
		var scope *ZendClassEntry = ZendGetExecutedScope()
		if scope != nil {
			RETVAL_STR_COPY(scope.GetName())
			return
		} else {
			ZendError(E_WARNING, "get_class() called without object from outside a class")
			RETVAL_FALSE
			return
		}
	}
	RETVAL_STR_COPY(Z_OBJCE_P(obj).GetName())
	return
}
func ZifGetCalledClass(execute_data *ZendExecuteData, return_value *Zval) {
	var called_scope *ZendClassEntry
	if ZendParseParametersNone() == FAILURE {
		return
	}
	called_scope = ZendGetCalledScope(execute_data)
	if called_scope != nil {
		RETVAL_STR_COPY(called_scope.GetName())
		return
	} else {
		var scope *ZendClassEntry = ZendGetExecutedScope()
		if scope == nil {
			ZendError(E_WARNING, "get_called_class() called from outside a class")
		}
	}
	RETVAL_FALSE
	return
}
func ZifGetParentClass(execute_data *ZendExecuteData, return_value *Zval) {
	var arg *Zval
	var ce *ZendClassEntry = nil
	if ZendParseParameters(ZEND_NUM_ARGS(), "|z", &arg) == FAILURE {
		return
	}
	if ZEND_NUM_ARGS() == 0 {
		ce = ZendGetExecutedScope()
		if ce != nil && ce.parent {
			RETVAL_STR_COPY(ce.parent.name)
			return
		} else {
			RETVAL_FALSE
			return
		}
	}
	if arg.IsObject() {
		ce = Z_OBJ_P(arg).GetCe()
	} else if arg.IsString() {
		ce = ZendLookupClass(arg.GetStr())
	}
	if ce != nil && ce.parent {
		RETVAL_STR_COPY(ce.parent.name)
		return
	} else {
		RETVAL_FALSE
		return
	}
}
func IsAImpl(execute_data *ZendExecuteData, return_value *Zval, only_subclass ZendBool) {
	var obj *Zval
	var class_name *ZendString
	var instance_ce *ZendClassEntry
	var ce *ZendClassEntry
	var allow_string ZendBool = only_subclass
	var retval ZendBool
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			ZendParseArgZvalDeref(_arg, &obj, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &class_name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgBool(_arg, &allow_string, &_dummy, 0) == 0 {
				_expected_type = Z_EXPECTED_BOOL
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}

	/*
	 * allow_string - is_a default is no, is_subclass_of is yes.
	 *   if it's allowed, then the autoloader will be called if the class does not exist.
	 *   default behaviour is different, as 'is_a' used to be used to test mixed return values
	 *   and there is no easy way to deprecate this.
	 */

	if allow_string != 0 && obj.IsString() {
		instance_ce = ZendLookupClass(obj.GetStr())
		if instance_ce == nil {
			RETVAL_FALSE
			return
		}
	} else if obj.IsObject() {
		instance_ce = Z_OBJCE_P(obj)
	} else {
		RETVAL_FALSE
		return
	}
	if only_subclass == 0 && ZendStringEquals(instance_ce.GetName(), class_name) != 0 {
		retval = 1
	} else {
		ce = ZendLookupClassEx(class_name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if ce == nil {
			retval = 0
		} else {
			if only_subclass != 0 && instance_ce == ce {
				retval = 0
			} else {
				retval = InstanceofFunction(instance_ce, ce)
			}
		}
	}
	RETVAL_BOOL(retval != 0)
	return
}
func ZifIsSubclassOf(execute_data *ZendExecuteData, return_value *Zval) {
	IsAImpl(execute_data, return_value, 1)
}
func ZifIsA(execute_data *ZendExecuteData, return_value *Zval) {
	IsAImpl(execute_data, return_value, 0)
}
func AddClassVars(scope *ZendClassEntry, ce *ZendClassEntry, statics int, return_value *Zval) {
	var prop_info *ZendPropertyInfo
	var prop *Zval
	var prop_copy Zval
	var key *ZendString
	var __ht *HashTable = ce.GetPropertiesInfo()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		key = _p.GetKey()
		prop_info = _z.GetPtr()
		if prop_info.IsProtected() && ZendCheckProtected(prop_info.GetCe(), scope) == 0 || prop_info.IsPrivate() && prop_info.GetCe() != scope {
			continue
		}
		prop = nil
		if statics != 0 && prop_info.IsStatic() {
			prop = ce.GetDefaultStaticMembersTable()[prop_info.GetOffset()]
			ZVAL_DEINDIRECT(prop)
		} else if statics == 0 && !prop_info.IsStatic() {
			prop = ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(prop_info.GetOffset())]
		}
		if prop == nil {
			continue
		}
		if prop.IsUndef() {

			/* Return uninitialized typed properties as a null value */

			prop_copy.SetNull()

			/* Return uninitialized typed properties as a null value */

		} else {

			/* copy: enforce read only access */

			ZVAL_COPY_OR_DUP(&prop_copy, prop)

			/* copy: enforce read only access */

		}
		prop = &prop_copy

		/* this is necessary to make it able to work with default array
		 * properties, returned to user */

		if Z_OPT_TYPE_P(prop) == IS_CONSTANT_AST {
			if ZvalUpdateConstantEx(prop, nil) != SUCCESS {
				return
			}
		}
		return_value.GetArr().KeyAddNew(key.GetStr(), prop)
	}
}
func ZifGetClassVars(execute_data *ZendExecuteData, return_value *Zval) {
	var class_name *ZendString
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	if ZendParseParameters(ZEND_NUM_ARGS(), "S", &class_name) == FAILURE {
		return
	}
	ce = ZendLookupClass(class_name)
	if ce == nil {
		RETVAL_FALSE
		return
	} else {
		ArrayInit(return_value)
		if !ce.IsConstantsUpdated() {
			if ZendUpdateClassConstants(ce) != SUCCESS {
				return
			}
		}
		scope = ZendGetExecutedScope()
		AddClassVars(scope, ce, 0, return_value)
		AddClassVars(scope, ce, 1, return_value)
	}
}
func ZifGetObjectVars(execute_data *ZendExecuteData, return_value *Zval) {
	var obj *Zval
	var value *Zval
	var properties *HashTable
	var key *ZendString
	var zobj *ZendObject
	var num_key ZendUlong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgObject(_arg, &obj, nil, 0) == 0 {
				_expected_type = Z_EXPECTED_OBJECT
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	properties = Z_OBJ_HT_P(obj).GetGetProperties()(obj)
	if properties == nil {
		RETVAL_FALSE
		return
	}
	zobj = obj.GetObj()
	if zobj.GetCe().GetDefaultPropertiesCount() == 0 && properties == zobj.GetProperties() && !(properties.IsRecursive()) {

		/* fast copy */

		if zobj.GetHandlers() == &StdObjectHandlers {
			RETVAL_ARR(ZendProptableToSymtable(properties, 0))
			return
		}
		RETVAL_ARR(ZendProptableToSymtable(properties, 1))
		return
	} else {
		ArrayInitSize(return_value, properties.GetNNumOfElements())
		var __ht *HashTable = properties
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			num_key = _p.GetH()
			key = _p.GetKey()
			value = _z
			var is_dynamic ZendBool = 1
			if value.IsIndirect() {
				value = value.GetZv()
				if value.IsUndef() {
					continue
				}
				is_dynamic = 0
			}
			if key != nil && ZendCheckPropertyAccess(zobj, key, is_dynamic) == FAILURE {
				continue
			}
			if value.IsReference() && Z_REFCOUNT_P(value) == 1 {
				value = Z_REFVAL_P(value)
			}
			Z_TRY_ADDREF_P(value)
			if key == nil {

				/* This case is only possible due to loopholes, e.g. ArrayObject */

				return_value.GetArr().IndexAddH(num_key, value)

				/* This case is only possible due to loopholes, e.g. ArrayObject */

			} else if is_dynamic == 0 && key.GetVal()[0] == 0 {
				var prop_name *byte
				var class_name *byte
				var prop_len int
				ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, &prop_len)

				/* We assume here that a mangled property name is never
				 * numeric. This is probably a safe assumption, but
				 * theoretically someone might write an extension with
				 * private, numeric properties. Well, too bad.
				 */

				return_value.GetArr().KeyAddNew(b.CastStr(prop_name, prop_len), value)

				/* We assume here that a mangled property name is never
				 * numeric. This is probably a safe assumption, but
				 * theoretically someone might write an extension with
				 * private, numeric properties. Well, too bad.
				 */

			} else {
				ZendSymtableAddNew(return_value.GetArr(), key, value)
			}
		}
	}
}
func ZifGetMangledObjectVars(execute_data *ZendExecuteData, return_value *Zval) {
	var obj *Zval
	var properties *HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgObject(_arg, &obj, nil, 0) == 0 {
				_expected_type = Z_EXPECTED_OBJECT
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	properties = Z_OBJ_HT_P(obj).GetGetProperties()(obj)
	if properties == nil {
		ZVAL_EMPTY_ARRAY(return_value)
		return
	}
	properties = ZendProptableToSymtable(properties, Z_OBJCE_P(obj).GetDefaultPropertiesCount() != 0 || Z_OBJ_P(obj).GetHandlers() != &StdObjectHandlers || properties.IsRecursive())
	RETVAL_ARR(properties)
	return
}
func SameName(key *ZendString, name *ZendString) int {
	var lcname *ZendString
	var ret int
	if key == name {
		return 1
	}
	if key.GetLen() != name.GetLen() {
		return 0
	}
	lcname = ZendStringTolower(name)
	ret = memcmp(lcname.GetVal(), key.GetVal(), key.GetLen()) == 0
	ZendStringReleaseEx(lcname, 0)
	return ret
}
func ZifGetClassMethods(execute_data *ZendExecuteData, return_value *Zval) {
	var klass *Zval
	var method_name Zval
	var ce *ZendClassEntry = nil
	var scope *ZendClassEntry
	var mptr *ZendFunction
	var key *ZendString
	if ZendParseParameters(ZEND_NUM_ARGS(), "z", &klass) == FAILURE {
		return
	}
	if klass.IsObject() {
		ce = Z_OBJCE_P(klass)
	} else if klass.IsString() {
		ce = ZendLookupClass(klass.GetStr())
	}
	if ce == nil {
		RETVAL_NULL()
		return
	}
	ArrayInit(return_value)
	scope = ZendGetExecutedScope()
	var __ht *HashTable = ce.GetFunctionTable()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		key = _p.GetKey()
		mptr = _z.GetPtr()
		if mptr.IsPublic() || scope != nil && (mptr.IsProtected() && ZendCheckProtected(mptr.GetScope(), scope) != 0 || mptr.IsPrivate() && scope == mptr.GetScope()) {
			if mptr.GetType() == ZEND_USER_FUNCTION && (mptr.GetOpArray().GetRefcount() == nil || mptr.op_array.refcount > 1) && key != nil && SameName(key, mptr.GetFunctionName()) == 0 {
				ZVAL_STR_COPY(&method_name, ZendFindAliasName(mptr.GetScope(), key))
				return_value.GetArr().NextIndexInsertNew(&method_name)
			} else {
				ZVAL_STR_COPY(&method_name, mptr.GetFunctionName())
				return_value.GetArr().NextIndexInsertNew(&method_name)
			}
		}
	}
}
func ZifMethodExists(execute_data *ZendExecuteData, return_value *Zval) {
	var klass *Zval
	var method_name *ZendString
	var lcname *ZendString
	var ce *ZendClassEntry
	var func_ *ZendFunction
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			ZendParseArgZvalDeref(_arg, &klass, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &method_name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if klass.IsObject() {
		ce = Z_OBJCE_P(klass)
	} else if klass.IsString() {
		if b.Assign(&ce, ZendLookupClass(klass.GetStr())) == nil {
			RETVAL_FALSE
			return
		}
	} else {
		RETVAL_FALSE
		return
	}
	lcname = ZendStringTolower(method_name)
	func_ = ZendHashFindPtr(ce.GetFunctionTable(), lcname)
	ZendStringReleaseEx(lcname, 0)
	if func_ != nil {

		/* Exclude shadow properties when checking a method on a specific class. Include
		 * them when checking an object, as method_exists() generally ignores visibility.
		 * TODO: Should we use EG(scope) for the object case instead? */

		RETVAL_BOOL(klass.IsObject() || !func_.IsPrivate() || func_.GetScope() == ce)
		return
	}
	if klass.IsObject() {
		var obj *ZendObject = klass.GetObj()
		func_ = Z_OBJ_HT_P(klass).GetGetMethod()(&obj, method_name, nil)
		if func_ != nil {
			if func_.IsCallViaTrampoline() {

				/* Returns true to the fake Closure's __invoke */

				RETVAL_BOOL(func_.GetScope() == ZendCeClosure && ZendStringEqualsLiteral(method_name, ZEND_INVOKE_FUNC_NAME))
				ZendStringReleaseEx(func_.GetFunctionName(), 0)
				ZendFreeTrampoline(func_)
				return
			}
			RETVAL_TRUE
			return
		}
	}
	RETVAL_FALSE
	return
}
func ZifPropertyExists(execute_data *ZendExecuteData, return_value *Zval) {
	var object *Zval
	var property *ZendString
	var ce *ZendClassEntry
	var property_info *ZendPropertyInfo
	var property_z Zval
	if ZendParseParameters(ZEND_NUM_ARGS(), "zS", &object, &property) == FAILURE {
		return
	}
	if property == nil {
		RETVAL_FALSE
		return
	}
	if object.IsString() {
		ce = ZendLookupClass(object.GetStr())
		if ce == nil {
			RETVAL_FALSE
			return
		}
	} else if object.IsObject() {
		ce = Z_OBJCE_P(object)
	} else {
		ZendError(E_WARNING, "First parameter must either be an object or the name of an existing class")
		RETVAL_NULL()
		return
	}
	property_info = ZendHashFindPtr(ce.GetPropertiesInfo(), property)
	if property_info != nil && (!property_info.IsPrivate() || property_info.GetCe() == ce) {
		RETVAL_TRUE
		return
	}
	ZVAL_STR(&property_z, property)
	if object.IsObject() && Z_OBJ_HT(*object).GetHasProperty()(object, &property_z, 2, nil) != 0 {
		RETVAL_TRUE
		return
	}
	RETVAL_FALSE
	return
}
func ClassExistsImpl(execute_data *ZendExecuteData, return_value *Zval, flags int, skip_flags int) {
	var name *ZendString
	var lcname *ZendString
	var ce *ZendClassEntry
	var autoload ZendBool = 1
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgBool(_arg, &autoload, &_dummy, 0) == 0 {
				_expected_type = Z_EXPECTED_BOOL
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if autoload == 0 {
		if name.GetVal()[0] == '\\' {

			/* Ignore leading "\" */

			lcname = ZendStringAlloc(name.GetLen()-1, 0)
			ZendStrTolowerCopy(lcname.GetVal(), name.GetVal()+1, name.GetLen()-1)
		} else {
			lcname = ZendStringTolower(name)
		}
		ce = ZendHashFindPtr(EG__().GetClassTable(), lcname)
		ZendStringReleaseEx(lcname, 0)
	} else {
		ce = ZendLookupClass(name)
	}
	if ce != nil {
		RETVAL_BOOL((ce.GetCeFlags()&flags) == flags && !ce.HasCeFlags(skip_flags))
		return
	} else {
		RETVAL_FALSE
		return
	}
}
func ZifClassExists(execute_data *ZendExecuteData, return_value *Zval) {
	ClassExistsImpl(execute_data, return_value, ZEND_ACC_LINKED, ZEND_ACC_INTERFACE|ZEND_ACC_TRAIT)
}
func ZifInterfaceExists(execute_data *ZendExecuteData, return_value *Zval) {
	ClassExistsImpl(execute_data, return_value, ZEND_ACC_LINKED|ZEND_ACC_INTERFACE, 0)
}
func ZifTraitExists(execute_data *ZendExecuteData, return_value *Zval) {
	ClassExistsImpl(execute_data, return_value, ZEND_ACC_TRAIT, 0)
}
func ZifFunctionExists(execute_data *ZendExecuteData, return_value *Zval) {
	var name *ZendString
	var func_ *ZendFunction
	var lcname *ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *Zval
		var _arg *Zval = nil
		var _expected_type ZendExpectedType = Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy ZendBool
		var _optional ZendBool = 0
		var _error_code int = ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = ZPP_ERROR_FAILURE
				break
			}
			_real_arg = ZEND_CALL_ARG(execute_data, 0)
			Z_PARAM_PROLOGUE(0, 0)
			if ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = Z_EXPECTED_STRING
				_error_code = ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != ZPP_ERROR_OK {
			if (_flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongCallbackException(_i, _error)
					} else {
						ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == ZPP_ERROR_WRONG_CLASS {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == ZPP_ERROR_WRONG_ARG {
					if (_flags & ZEND_PARSE_PARAMS_THROW) != 0 {
						ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	if name.GetVal()[0] == '\\' {

		/* Ignore leading "\" */

		lcname = ZendStringAlloc(name.GetLen()-1, 0)
		ZendStrTolowerCopy(lcname.GetVal(), name.GetVal()+1, name.GetLen()-1)
	} else {
		lcname = ZendStringTolower(name)
	}
	func_ = ZendHashFindPtr(EG__().GetFunctionTable(), lcname)
	ZendStringReleaseEx(lcname, 0)

	/*
	 * A bit of a hack, but not a bad one: we see if the handler of the function
	 * is actually one that displays "function is disabled" message.
	 */

	RETVAL_BOOL(func_ != nil && (func_.GetType() != ZEND_INTERNAL_FUNCTION || func_.GetInternalFunction().GetHandler() != ZifDisplayDisabledFunction))
	return
}
func ZifClassAlias(execute_data *ZendExecuteData, return_value *Zval) {
	var class_name *ZendString
	var alias_name *byte
	var ce *ZendClassEntry
	var alias_name_len int
	var autoload ZendBool = 1
	if ZendParseParameters(ZEND_NUM_ARGS(), "Ss|b", &class_name, &alias_name, &alias_name_len, &autoload) == FAILURE {
		return
	}
	ce = ZendLookupClassEx(class_name, nil, b.Cond(autoload == 0, ZEND_FETCH_CLASS_NO_AUTOLOAD, 0))
	if ce != nil {
		if ce.GetType() == ZEND_USER_CLASS {
			if ZendRegisterClassAliasEx(alias_name, alias_name_len, ce, 0) == SUCCESS {
				RETVAL_TRUE
				return
			} else {
				ZendError(E_WARNING, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), alias_name)
				RETVAL_FALSE
				return
			}
		} else {
			ZendError(E_WARNING, "First argument of class_alias() must be a name of user defined class")
			RETVAL_FALSE
			return
		}
	} else {
		ZendError(E_WARNING, "Class '%s' not found", class_name.GetVal())
		RETVAL_FALSE
		return
	}
}
func ZifGetIncludedFiles(execute_data *ZendExecuteData, return_value *Zval) {
	var entry *ZendString
	if ZendParseParametersNone() == FAILURE {
		return
	}
	ArrayInit(return_value)
	var __ht *HashTable = EG__().GetIncludedFiles()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		entry = _p.GetKey()
		if entry != nil {
			AddNextIndexStr(return_value, entry.Copy())
		}
	}
}
func ZifTriggerError(execute_data *ZendExecuteData, return_value *Zval) {
	var error_type ZendLong = E_USER_NOTICE
	var message *byte
	var message_len int
	if ZendParseParameters(ZEND_NUM_ARGS(), "s|l", &message, &message_len, &error_type) == FAILURE {
		return
	}
	switch error_type {
	case E_USER_ERROR:

	case E_USER_WARNING:

	case E_USER_NOTICE:

	case E_USER_DEPRECATED:
		break
	default:
		ZendError(E_WARNING, "Invalid error type specified")
		RETVAL_FALSE
		return
		break
	}
	ZendError(int(error_type), "%s", message)
	RETVAL_TRUE
	return
}
func ZifSetErrorHandler(execute_data *ZendExecuteData, return_value *Zval) {
	var error_handler *Zval
	var error_type ZendLong = E_ALL
	if ZendParseParameters(ZEND_NUM_ARGS(), "z|l", &error_handler, &error_type) == FAILURE {
		return
	}
	if error_handler.GetType() != IS_NULL {
		if ZendIsCallable(error_handler, 0, nil) == 0 {
			var error_handler_name *ZendString = ZendGetCallableName(error_handler)
			ZendError(E_WARNING, "%s() expects the argument (%s) to be a valid callback", GetActiveFunctionName(), b.CondF1(error_handler_name != nil, func() []byte { return error_handler_name.GetVal() }, "unknown"))
			ZendStringReleaseEx(error_handler_name, 0)
			return
		}
	}
	if EG__().GetUserErrorHandler().GetType() != IS_UNDEF {
		ZVAL_COPY(return_value, EG__().GetUserErrorHandler())
	}
	ZendStackPush(EG__().GetUserErrorHandlersErrorReporting(), EG__().GetUserErrorHandlerErrorReporting())
	ZendStackPush(EG__().GetUserErrorHandlers(), EG__().GetUserErrorHandler())
	if error_handler.IsNull() {
		EG__().GetUserErrorHandler().SetUndef()
		return
	}
	ZVAL_COPY(EG__().GetUserErrorHandler(), error_handler)
	EG__().SetUserErrorHandlerErrorReporting(int(error_type))
}
func ZifRestoreErrorHandler(execute_data *ZendExecuteData, return_value *Zval) {
	if ZendParseParametersNone() == FAILURE {
		return
	}
	if EG__().GetUserErrorHandler().GetType() != IS_UNDEF {
		var zeh Zval
		ZVAL_COPY_VALUE(&zeh, EG__().GetUserErrorHandler())
		EG__().GetUserErrorHandler().SetUndef()
		ZvalPtrDtor(&zeh)
	}
	if ZendStackIsEmpty(EG__().GetUserErrorHandlers()) != 0 {
		EG__().GetUserErrorHandler().SetUndef()
	} else {
		var tmp *Zval
		EG__().SetUserErrorHandlerErrorReporting(ZendStackIntTop(EG__().GetUserErrorHandlersErrorReporting()))
		ZendStackDelTop(EG__().GetUserErrorHandlersErrorReporting())
		tmp = ZendStackTop(EG__().GetUserErrorHandlers())
		ZVAL_COPY_VALUE(EG__().GetUserErrorHandler(), tmp)
		ZendStackDelTop(EG__().GetUserErrorHandlers())
	}
	RETVAL_TRUE
	return
}
func ZifSetExceptionHandler(execute_data *ZendExecuteData, return_value *Zval) {
	var exception_handler *Zval
	if ZendParseParameters(ZEND_NUM_ARGS(), "z", &exception_handler) == FAILURE {
		return
	}
	if exception_handler.GetType() != IS_NULL {
		if ZendIsCallable(exception_handler, 0, nil) == 0 {
			var exception_handler_name *ZendString = ZendGetCallableName(exception_handler)
			ZendError(E_WARNING, "%s() expects the argument (%s) to be a valid callback", GetActiveFunctionName(), b.CondF1(exception_handler_name != nil, func() []byte { return exception_handler_name.GetVal() }, "unknown"))
			ZendStringReleaseEx(exception_handler_name, 0)
			return
		}
	}
	if EG__().GetUserExceptionHandler().GetType() != IS_UNDEF {
		ZVAL_COPY(return_value, EG__().GetUserExceptionHandler())
	}
	ZendStackPush(EG__().GetUserExceptionHandlers(), EG__().GetUserExceptionHandler())
	if exception_handler.IsNull() {
		EG__().GetUserExceptionHandler().SetUndef()
		return
	}
	ZVAL_COPY(EG__().GetUserExceptionHandler(), exception_handler)
}
func ZifRestoreExceptionHandler(execute_data *ZendExecuteData, return_value *Zval) {
	if ZendParseParametersNone() == FAILURE {
		return
	}
	if EG__().GetUserExceptionHandler().GetType() != IS_UNDEF {
		ZvalPtrDtor(EG__().GetUserExceptionHandler())
	}
	if ZendStackIsEmpty(EG__().GetUserExceptionHandlers()) != 0 {
		EG__().GetUserExceptionHandler().SetUndef()
	} else {
		var tmp *Zval = ZendStackTop(EG__().GetUserExceptionHandlers())
		ZVAL_COPY_VALUE(EG__().GetUserExceptionHandler(), tmp)
		ZendStackDelTop(EG__().GetUserExceptionHandlers())
	}
	RETVAL_TRUE
	return
}
func CopyClassOrInterfaceName(array *Zval, key *ZendString, ce *ZendClassEntry) {
	if ce.GetRefcount() == 1 && !ce.IsImmutable() || SameName(key, ce.GetName()) != 0 {
		key = ce.GetName()
	}
	AddNextIndexStr(array, key.Copy())
}
func GetDeclaredClassImpl(execute_data *ZendExecuteData, return_value *Zval, flags int, skip_flags int) {
	var key *ZendString
	var ce *ZendClassEntry
	if ZendParseParametersNone() == FAILURE {
		return
	}
	ArrayInit(return_value)
	var __ht *HashTable = EG__().GetClassTable()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		key = _p.GetKey()
		ce = _z.GetPtr()
		if key != nil && key.GetVal()[0] != 0 && ce.HasCeFlags(flags) && !ce.HasCeFlags(skip_flags) {
			CopyClassOrInterfaceName(return_value, key, ce)
		}
	}
}
func ZifGetDeclaredTraits(execute_data *ZendExecuteData, return_value *Zval) {
	GetDeclaredClassImpl(execute_data, return_value, ZEND_ACC_TRAIT, 0)
}
func ZifGetDeclaredClasses(execute_data *ZendExecuteData, return_value *Zval) {
	GetDeclaredClassImpl(execute_data, return_value, ZEND_ACC_LINKED, ZEND_ACC_INTERFACE|ZEND_ACC_TRAIT)
}
func ZifGetDeclaredInterfaces(execute_data *ZendExecuteData, return_value *Zval) {
	GetDeclaredClassImpl(execute_data, return_value, ZEND_ACC_INTERFACE, 0)
}
func ZifGetDefinedFunctions(execute_data *ZendExecuteData, return_value *Zval) {
	var internal Zval
	var user Zval
	var key *ZendString
	var func_ *ZendFunction
	var exclude_disabled ZendBool = 0
	if ZendParseParameters(ZEND_NUM_ARGS(), "|b", &exclude_disabled) == FAILURE {
		return
	}
	ArrayInit(&internal)
	ArrayInit(&user)
	ArrayInit(return_value)
	var __ht *HashTable = EG__().GetFunctionTable()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		key = _p.GetKey()
		func_ = _z.GetPtr()
		if key != nil && key.GetVal()[0] != 0 {
			if func_.GetType() == ZEND_INTERNAL_FUNCTION && (exclude_disabled == 0 || func_.GetInternalFunction().GetHandler() != ZifDisplayDisabledFunction) {
				AddNextIndexStr(&internal, key.Copy())
			} else if func_.GetType() == ZEND_USER_FUNCTION {
				AddNextIndexStr(&user, key.Copy())
			}
		}
	}
	return_value.GetArr().KeyAddNew(b.CastStr("internal", b.SizeOf("\"internal\"")-1), &internal)
	return_value.GetArr().KeyAddNew(b.CastStr("user", b.SizeOf("\"user\"")-1), &user)
}
func ZifGetDefinedVars(execute_data *ZendExecuteData, return_value *Zval) {
	var symbol_table *ZendArray
	if ZendForbidDynamicCall("get_defined_vars()") == FAILURE {
		return
	}
	symbol_table = ZendRebuildSymbolTable()
	if symbol_table == nil {
		return
	}
	RETVAL_ARR(ZendArrayDup(symbol_table))
	return
}
func ZifCreateFunction(execute_data *ZendExecuteData, return_value *Zval) {
	var function_name *ZendString
	var eval_code *byte
	var function_args *byte
	var function_code *byte
	var eval_code_length int
	var function_args_len int
	var function_code_len int
	var retval int
	var eval_name *byte
	if ZendParseParameters(ZEND_NUM_ARGS(), "ss", &function_args, &function_args_len, &function_code, &function_code_len) == FAILURE {
		return
	}
	eval_code = (*byte)(Emalloc(b.SizeOf("\"function \" LAMBDA_TEMP_FUNCNAME") + function_args_len + 2 + 2 + function_code_len))
	eval_code_length = b.SizeOf("\"function \" LAMBDA_TEMP_FUNCNAME \"(\"") - 1
	memcpy(eval_code, "function "+LAMBDA_TEMP_FUNCNAME+"(", eval_code_length)
	memcpy(eval_code+eval_code_length, function_args, function_args_len)
	eval_code_length += function_args_len
	eval_code[b.PostInc(&eval_code_length)] = ')'
	eval_code[b.PostInc(&eval_code_length)] = '{'
	memcpy(eval_code+eval_code_length, function_code, function_code_len)
	eval_code_length += function_code_len
	eval_code[b.PostInc(&eval_code_length)] = '}'
	eval_code[eval_code_length] = '0'
	eval_name = ZendMakeCompiledStringDescription("runtime-created function")
	retval = ZendEvalStringl(eval_code, eval_code_length, nil, eval_name)
	Efree(eval_code)
	Efree(eval_name)
	if retval == SUCCESS {
		var func_ *ZendOpArray
		var static_variables *HashTable
		func_ = ZendHashStrFindPtr(EG__().GetFunctionTable(), LAMBDA_TEMP_FUNCNAME, b.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		if func_ == nil {
			ZendErrorNoreturn(E_CORE_ERROR, "Unexpected inconsistency in create_function()")
			RETVAL_FALSE
			return
		}
		if func_.GetRefcount() != nil {
			func_.refcount++
		}
		static_variables = func_.GetStaticVariables()
		func_.SetStaticVariables(nil)
		ZendHashStrDel(EG__().GetFunctionTable(), LAMBDA_TEMP_FUNCNAME, b.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		func_.SetStaticVariables(static_variables)
		function_name = ZendStringAlloc(b.SizeOf("\"0lambda_\"")+MAX_LENGTH_OF_LONG, 0)
		function_name.GetVal()[0] = '0'
		for {
			function_name.SetLen(core.Snprintf(function_name.GetVal()+1, b.SizeOf("\"lambda_\"")+MAX_LENGTH_OF_LONG, "lambda_%d", b.PreInc(&(EG__().GetLambdaCount()))) + 1)
			if ZendHashAddPtr(EG__().GetFunctionTable(), function_name, func_) != nil {
				break
			}
		}
		RETVAL_NEW_STR(function_name)
		return
	} else {
		ZendHashStrDel(EG__().GetFunctionTable(), LAMBDA_TEMP_FUNCNAME, b.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		RETVAL_FALSE
		return
	}
}
func ZifGetResourceType(execute_data *ZendExecuteData, return_value *Zval) {
	var resource_type *byte
	var z_resource_type *Zval
	if ZendParseParameters(ZEND_NUM_ARGS(), "r", &z_resource_type) == FAILURE {
		return
	}
	resource_type = ZendRsrcListGetRsrcType(z_resource_type.GetRes())
	if resource_type != nil {
		RETVAL_STRING(resource_type)
		return
	} else {
		RETVAL_STRING("Unknown")
		return
	}
}
func ZifGetResources(execute_data *ZendExecuteData, return_value *Zval) {
	var type_ *ZendString = nil
	var key *ZendString
	var index ZendUlong
	var val *Zval
	if ZendParseParameters(ZEND_NUM_ARGS(), "|S", &type_) == FAILURE {
		return
	}
	if type_ == nil {
		ArrayInit(return_value)
		var __ht *HashTable = EG__().GetRegularList()
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			if key == nil {
				Z_ADDREF_P(val)
				return_value.GetArr().IndexAddNewH(index, val)
			}
		}
	} else if ZendStringEqualsLiteral(type_, "Unknown") {
		ArrayInit(return_value)
		var __ht *HashTable = EG__().GetRegularList()
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			if key == nil && Z_RES_TYPE_P(val) <= 0 {
				Z_ADDREF_P(val)
				return_value.GetArr().IndexAddNewH(index, val)
			}
		}
	} else {
		var id int = ZendFetchListDtorId(type_.GetVal())
		if id <= 0 {
			ZendError(E_WARNING, "get_resources():  Unknown resource type '%s'", type_.GetVal())
			RETVAL_FALSE
			return
		}
		ArrayInit(return_value)
		var __ht *HashTable = EG__().GetRegularList()
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			if key == nil && Z_RES_TYPE_P(val) == id {
				Z_ADDREF_P(val)
				return_value.GetArr().IndexAddNewH(index, val)
			}
		}
	}
}
func AddZendextInfo(ext *ZendExtension, arg any) int {
	var name_array *Zval = (*Zval)(arg)
	AddNextIndexString(name_array, ext.GetName())
	return 0
}
func ZifGetLoadedExtensions(execute_data *ZendExecuteData, return_value *Zval) {
	var zendext ZendBool = 0
	if ZendParseParameters(ZEND_NUM_ARGS(), "|b", &zendext) == FAILURE {
		return
	}
	ArrayInit(return_value)
	if zendext != 0 {
		ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(AddZendextInfo), return_value)
	} else {
		var module *ZendModuleEntry
		var __ht *HashTable = &ModuleRegistry
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			module = _z.GetPtr()
			AddNextIndexString(return_value, module.GetName())
		}
	}
}
func ZifGetDefinedConstants(execute_data *ZendExecuteData, return_value *Zval) {
	var categorize ZendBool = 0
	if ZendParseParameters(ZEND_NUM_ARGS(), "|b", &categorize) == FAILURE {
		return
	}
	ArrayInit(return_value)
	if categorize != 0 {
		var val *ZendConstant
		var module_number int
		var modules *Zval
		var const_val Zval
		var module_names **byte
		var module *ZendModuleEntry
		var i int = 1
		modules = Ecalloc(ModuleRegistry.GetNNumOfElements()+2, b.SizeOf("zval"))
		module_names = Emalloc((ModuleRegistry.GetNNumOfElements() + 2) * b.SizeOf("char *"))
		module_names[0] = "internal"
		var __ht *HashTable = &ModuleRegistry
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			module = _z.GetPtr()
			module_names[module.GetModuleNumber()] = (*byte)(module.GetName())
			i++
		}
		module_names[i] = "user"
		var __ht__1 *HashTable = EG__().GetZendConstants()
		for _, _p := range __ht__1.foreachData() {
			var _z *Zval = _p.GetVal()

			val = _z.GetPtr()
			if val.GetName() == nil {

				/* skip special constants */

				continue

				/* skip special constants */

			}
			if ZEND_CONSTANT_MODULE_NUMBER(val) == PHP_USER_CONSTANT {
				module_number = i
			} else if ZEND_CONSTANT_MODULE_NUMBER(val) > i {

				/* should not happen */

				continue

				/* should not happen */

			} else {
				module_number = ZEND_CONSTANT_MODULE_NUMBER(val)
			}
			if modules[module_number].IsUndef() {
				ArrayInit(&modules[module_number])
				AddAssocZval(return_value, module_names[module_number], &modules[module_number])
			}
			ZVAL_COPY_OR_DUP(&const_val, val.GetValue())
			modules[module_number].GetArr().KeyAddNew(val.GetName().GetStr(), &const_val)
		}
		Efree(module_names)
		Efree(modules)
	} else {
		var constant *ZendConstant
		var const_val Zval
		var __ht *HashTable = EG__().GetZendConstants()
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			constant = _z.GetPtr()
			if constant.GetName() == nil {

				/* skip special constants */

				continue

				/* skip special constants */

			}
			ZVAL_COPY_OR_DUP(&const_val, constant.GetValue())
			return_value.GetArr().KeyAddNew(constant.GetName().GetStr(), &const_val)
		}
	}
}
func DebugBacktraceGetArgs(call *ZendExecuteData, arg_array *Zval) {
	var num_args uint32 = ZEND_CALL_NUM_ARGS(call)
	if num_args != 0 {
		var i uint32 = 0
		var p *Zval = ZEND_CALL_ARG(call, 1)
		ArrayInitSize(arg_array, num_args)
		ZendHashRealInitPacked(arg_array.GetArr())
		var __fill_ht *HashTable = arg_array.GetArr()
		var __fill_bkt *Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		ZEND_ASSERT(__fill_ht.HasUFlags(HASH_FLAG_PACKED))
		if call.GetFunc().GetType() == ZEND_USER_FUNCTION {
			var first_extra_arg uint32 = MIN(num_args, call.GetFunc().GetOpArray().GetNumArgs())
			if (ZEND_CALL_INFO(call) & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {

				/* In case of attached symbol_table, values on stack may be invalid
				 * and we have to access them through symbol_table
				 * See: https://bugs.php.net/bug.php?id=73156
				 */

				var arg_name *ZendString
				var arg *Zval
				for i < first_extra_arg {
					arg_name = call.GetFunc().GetOpArray().GetVars()[i]
					arg = ZendHashFindExInd(call.GetSymbolTable(), arg_name, 1)
					if arg != nil {
						if Z_OPT_REFCOUNTED_P(arg) {
							Z_ADDREF_P(arg)
						}
						ZVAL_COPY_VALUE(__fill_bkt.GetVal(), arg)
					} else {
						__fill_bkt.GetVal().SetNull()
					}
					__fill_bkt.SetH(__fill_idx)
					__fill_bkt.SetKey(nil)
					__fill_bkt++
					__fill_idx++
					i++
				}
			} else {
				for i < first_extra_arg {
					if p.GetTypeInfo() != IS_UNDEF {
						if Z_OPT_REFCOUNTED_P(p) {
							Z_ADDREF_P(p)
						}
						ZVAL_COPY_VALUE(__fill_bkt.GetVal(), p)
					} else {
						__fill_bkt.GetVal().SetNull()
					}
					__fill_bkt.SetH(__fill_idx)
					__fill_bkt.SetKey(nil)
					__fill_bkt++
					__fill_idx++
					p++
					i++
				}
			}
			p = ZEND_CALL_VAR_NUM(call, call.GetFunc().GetOpArray().GetLastVar()+call.GetFunc().GetOpArray().GetT())
		}
		for i < num_args {
			if p.GetTypeInfo() != IS_UNDEF {
				if Z_OPT_REFCOUNTED_P(p) {
					Z_ADDREF_P(p)
				}
				ZVAL_COPY_VALUE(__fill_bkt.GetVal(), p)
			} else {
				__fill_bkt.GetVal().SetNull()
			}
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
			p++
			i++
		}
		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
		Z_ARRVAL_P(arg_array).SetNNumOfElements(num_args)
	} else {
		ZVAL_EMPTY_ARRAY(arg_array)
	}
}
func DebugPrintBacktraceArgs(arg_array *Zval) {
	var tmp *Zval
	var i int = 0
	var __ht *HashTable = arg_array.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		tmp = _z
		if b.PostInc(&i) {
			ZEND_PUTS(", ")
		}
		ZendPrintFlatZvalR(tmp)
	}
}
func SkipInternalHandler(skip *ZendExecuteData) ZendBool {
	return !(skip.GetFunc() != nil && ZEND_USER_CODE(skip.GetFunc().GetCommonType())) && skip.GetPrevExecuteData() != nil && skip.GetPrevExecuteData().GetFunc() != nil && ZEND_USER_CODE(skip.GetPrevExecuteData().GetFunc().GetCommonType()) && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_ICALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_UCALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL_BY_NAME && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL
}
func ZifDebugPrintBacktrace(execute_data *ZendExecuteData, return_value *Zval) {
	var call *ZendExecuteData
	var ptr *ZendExecuteData
	var skip *ZendExecuteData
	var object *ZendObject
	var lineno int
	var frameno int = 0
	var func_ *ZendFunction
	var function_name *byte
	var filename *byte
	var class_name *ZendString = nil
	var call_type *byte
	var include_filename *byte = nil
	var arg_array Zval
	var indent int = 0
	var options ZendLong = 0
	var limit ZendLong = 0
	if ZendParseParameters(ZEND_NUM_ARGS(), "|ll", &options, &limit) == FAILURE {
		return
	}
	arg_array.SetUndef()
	ptr = EX(prev_execute_data)

	/* skip debug_backtrace() */

	call = ptr
	ptr = ptr.GetPrevExecuteData()
	for ptr != nil && (limit == 0 || frameno < limit) {
		frameno++
		class_name = nil
		call_type = nil
		arg_array.SetUndef()
		ptr = ZendGeneratorCheckPlaceholderFrame(ptr)
		skip = ptr

		/* skip internal handler */

		if SkipInternalHandler(skip) != 0 {
			skip = skip.GetPrevExecuteData()
		}
		if skip.GetFunc() != nil && ZEND_USER_CODE(skip.GetFunc().GetCommonType()) {
			filename = skip.GetFunc().GetOpArray().GetFilename().GetVal()
			if skip.GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION {
				if EG__().GetOplineBeforeException() != nil {
					lineno = EG__().GetOplineBeforeException().GetLineno()
				} else {
					lineno = skip.GetFunc().GetOpArray().GetLineEnd()
				}
			} else {
				lineno = skip.GetOpline().GetLineno()
			}
		} else {
			filename = nil
			lineno = 0
		}

		/* $this may be passed into regular internal functions */

		if call.GetThis().IsObject() {
			object = call.GetThis().GetObj()
		} else {
			object = nil
		}
		if call.GetFunc() != nil {
			var zend_function_name *ZendString
			func_ = call.GetFunc()
			if func_.GetScope() != nil && func_.GetScope().GetTraitAliases() != nil {
				zend_function_name = ZendResolveMethodName(b.CondF(object != nil, func() *ZendClassEntry { return object.GetCe() }, func() *ZendClassEntry { return func_.GetScope() }), func_)
			} else {
				zend_function_name = func_.GetFunctionName()
			}
			if zend_function_name != nil {
				function_name = zend_function_name.GetVal()
			} else {
				function_name = nil
			}
		} else {
			func_ = nil
			function_name = nil
		}
		if function_name != nil {
			if object != nil {
				if func_.GetScope() != nil {
					class_name = func_.GetScope().GetName()
				} else if object.GetHandlers().GetGetClassName() == ZendStdGetClassName {
					class_name = object.GetCe().GetName()
				} else {
					class_name = object.GetHandlers().GetGetClassName()(object)
				}
				call_type = "->"
			} else if func_.GetScope() != nil {
				class_name = func_.GetScope().GetName()
				call_type = "::"
			} else {
				class_name = nil
				call_type = nil
			}
			if func_.GetType() != ZEND_EVAL_CODE {
				if (options & DEBUG_BACKTRACE_IGNORE_ARGS) == 0 {
					DebugBacktraceGetArgs(call, &arg_array)
				}
			}
		} else {

			/* i know this is kinda ugly, but i'm trying to avoid extra cycles in the main execution loop */

			var build_filename_arg ZendBool = 1
			if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetCommonType())) || ptr.GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL {

				/* can happen when calling eval from a custom sapi */

				function_name = "unknown"
				build_filename_arg = 0
			} else {
				switch ptr.GetOpline().GetExtendedValue() {
				case ZEND_EVAL:
					function_name = "eval"
					build_filename_arg = 0
					break
				case ZEND_INCLUDE:
					function_name = "include"
					break
				case ZEND_REQUIRE:
					function_name = "require"
					break
				case ZEND_INCLUDE_ONCE:
					function_name = "include_once"
					break
				case ZEND_REQUIRE_ONCE:
					function_name = "require_once"
					break
				default:

					/* this can actually happen if you use debug_backtrace() in your error_handler and
					 * you're in the top-scope */

					function_name = "unknown"
					build_filename_arg = 0
					break
				}
			}
			if build_filename_arg != 0 && include_filename != nil {
				ArrayInit(&arg_array)
				AddNextIndexString(&arg_array, (*byte)(include_filename))
			}
			call_type = nil
		}
		ZendPrintf("#%-2d ", indent)
		if class_name != nil {
			ZEND_PUTS(class_name.GetVal())
			ZEND_PUTS(call_type)
			if object != nil && func_.GetScope() == nil && object.GetHandlers().GetGetClassName() != ZendStdGetClassName {
				ZendStringReleaseEx(class_name, 0)
			}
		}
		ZendPrintf("%s(", function_name)
		if arg_array.GetType() != IS_UNDEF {
			DebugPrintBacktraceArgs(&arg_array)
			ZvalPtrDtor(&arg_array)
		}
		if filename != nil {
			ZendPrintf(") called at [%s:%d]\n", filename, lineno)
		} else {
			var prev_call *ZendExecuteData = skip
			var prev *ZendExecuteData = skip.GetPrevExecuteData()
			for prev != nil {
				if prev_call != nil && prev_call.GetFunc() != nil && !(ZEND_USER_CODE(prev_call.GetFunc().GetCommonType())) {
					prev = nil
					break
				}
				if prev.GetFunc() != nil && ZEND_USER_CODE(prev.GetFunc().GetCommonType()) {
					ZendPrintf(") called at [%s:%d]\n", prev.GetFunc().GetOpArray().GetFilename().GetVal(), prev.GetOpline().GetLineno())
					break
				}
				prev_call = prev
				prev = prev.GetPrevExecuteData()
			}
			if prev == nil {
				ZEND_PUTS(")\n")
			}
		}
		include_filename = filename
		call = skip
		ptr = skip.GetPrevExecuteData()
		indent++
	}
}
func ZendFetchDebugBacktrace(return_value *Zval, skip_last int, options int, limit int) {
	var ptr *ZendExecuteData
	var skip *ZendExecuteData
	var call *ZendExecuteData = nil
	var object *ZendObject
	var lineno int
	var frameno int = 0
	var func_ *ZendFunction
	var function_name *ZendString
	var filename *ZendString
	var include_filename *ZendString = nil
	var stack_frame Zval
	var tmp Zval
	ArrayInit(return_value)
	if !(b.Assign(&ptr, EG__().GetCurrentExecuteData())) {
		return
	}
	if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetCommonType())) {
		call = ptr
		ptr = ptr.GetPrevExecuteData()
	}
	if ptr != nil {
		if skip_last != 0 {

			/* skip debug_backtrace() */

			call = ptr
			ptr = ptr.GetPrevExecuteData()
		} else {

			/* skip "new Exception()" */

			if ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetCommonType()) && ptr.GetOpline().GetOpcode() == ZEND_NEW {
				call = ptr
				ptr = ptr.GetPrevExecuteData()
			}

			/* skip "new Exception()" */

		}
		if call == nil {
			call = ptr
			ptr = ptr.GetPrevExecuteData()
		}
	}
	for ptr != nil && (limit == 0 || frameno < limit) {
		frameno++
		ArrayInit(&stack_frame)
		ptr = ZendGeneratorCheckPlaceholderFrame(ptr)
		skip = ptr

		/* skip internal handler */

		if SkipInternalHandler(skip) != 0 {
			skip = skip.GetPrevExecuteData()
		}
		if skip.GetFunc() != nil && ZEND_USER_CODE(skip.GetFunc().GetCommonType()) {
			filename = skip.GetFunc().GetOpArray().GetFilename()
			if skip.GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION {
				if EG__().GetOplineBeforeException() != nil {
					lineno = EG__().GetOplineBeforeException().GetLineno()
				} else {
					lineno = skip.GetFunc().GetOpArray().GetLineEnd()
				}
			} else {
				lineno = skip.GetOpline().GetLineno()
			}
			ZVAL_STR_COPY(&tmp, filename)
			stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_FILE).GetStr(), &tmp)
			tmp.SetLong(lineno)
			stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_LINE).GetStr(), &tmp)
		} else {
			var prev_call *ZendExecuteData = skip
			var prev *ZendExecuteData = skip.GetPrevExecuteData()
			for prev != nil {
				if prev_call != nil && prev_call.GetFunc() != nil && !(ZEND_USER_CODE(prev_call.GetFunc().GetCommonType())) && !prev_call.GetFunc().IsCallViaTrampoline() {
					break
				}
				if prev.GetFunc() != nil && ZEND_USER_CODE(prev.GetFunc().GetCommonType()) {
					ZVAL_STR_COPY(&tmp, prev.GetFunc().GetOpArray().GetFilename())
					stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_FILE).GetStr(), &tmp)
					tmp.SetLong(prev.GetOpline().GetLineno())
					stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_LINE).GetStr(), &tmp)
					break
				}
				prev_call = prev
				prev = prev.GetPrevExecuteData()
			}
			filename = nil
		}

		/* $this may be passed into regular internal functions */

		if call != nil && call.GetThis().IsObject() {
			object = call.GetThis().GetObj()
		} else {
			object = nil
		}
		if call != nil && call.GetFunc() != nil {
			func_ = call.GetFunc()
			if func_.GetScope() != nil && func_.GetScope().GetTraitAliases() != nil {
				function_name = ZendResolveMethodName(b.CondF(object != nil, func() *ZendClassEntry { return object.GetCe() }, func() *ZendClassEntry { return func_.GetScope() }), func_)
			} else {
				function_name = func_.GetFunctionName()
			}
		} else {
			func_ = nil
			function_name = nil
		}
		if function_name != nil {
			ZVAL_STR_COPY(&tmp, function_name)
			stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_FUNCTION).GetStr(), &tmp)
			if object != nil {
				if func_.GetScope() != nil {
					ZVAL_STR_COPY(&tmp, func_.GetScope().GetName())
				} else if object.GetHandlers().GetGetClassName() == ZendStdGetClassName {
					ZVAL_STR_COPY(&tmp, object.GetCe().GetName())
				} else {
					ZVAL_STR(&tmp, object.GetHandlers().GetGetClassName()(object))
				}
				stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_CLASS).GetStr(), &tmp)
				if (options & DEBUG_BACKTRACE_PROVIDE_OBJECT) != 0 {
					ZVAL_OBJ(&tmp, object)
					stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_OBJECT).GetStr(), &tmp)
					Z_ADDREF(tmp)
				}
				ZVAL_INTERNED_STR(&tmp, ZSTR_KNOWN(ZEND_STR_OBJECT_OPERATOR))
				stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_TYPE).GetStr(), &tmp)
			} else if func_.GetScope() != nil {
				ZVAL_STR_COPY(&tmp, func_.GetScope().GetName())
				stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_CLASS).GetStr(), &tmp)
				ZVAL_INTERNED_STR(&tmp, ZSTR_KNOWN(ZEND_STR_PAAMAYIM_NEKUDOTAYIM))
				stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_TYPE).GetStr(), &tmp)
			}
			if (options&DEBUG_BACKTRACE_IGNORE_ARGS) == 0 && func_.GetType() != ZEND_EVAL_CODE {
				DebugBacktraceGetArgs(call, &tmp)
				stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_ARGS).GetStr(), &tmp)
			}
		} else {

			/* i know this is kinda ugly, but i'm trying to avoid extra cycles in the main execution loop */

			var build_filename_arg ZendBool = 1
			var pseudo_function_name *ZendString
			if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetCommonType())) || ptr.GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL {

				/* can happen when calling eval from a custom sapi */

				pseudo_function_name = ZSTR_KNOWN(ZEND_STR_UNKNOWN)
				build_filename_arg = 0
			} else {
				switch ptr.GetOpline().GetExtendedValue() {
				case ZEND_EVAL:
					pseudo_function_name = ZSTR_KNOWN(ZEND_STR_EVAL)
					build_filename_arg = 0
					break
				case ZEND_INCLUDE:
					pseudo_function_name = ZSTR_KNOWN(ZEND_STR_INCLUDE)
					break
				case ZEND_REQUIRE:
					pseudo_function_name = ZSTR_KNOWN(ZEND_STR_REQUIRE)
					break
				case ZEND_INCLUDE_ONCE:
					pseudo_function_name = ZSTR_KNOWN(ZEND_STR_INCLUDE_ONCE)
					break
				case ZEND_REQUIRE_ONCE:
					pseudo_function_name = ZSTR_KNOWN(ZEND_STR_REQUIRE_ONCE)
					break
				default:

					/* this can actually happen if you use debug_backtrace() in your error_handler and
					 * you're in the top-scope */

					pseudo_function_name = ZSTR_KNOWN(ZEND_STR_UNKNOWN)
					build_filename_arg = 0
					break
				}
			}
			if build_filename_arg != 0 && include_filename != nil {
				var arg_array Zval
				ArrayInit(&arg_array)

				/* include_filename always points to the last filename of the last last called-function.
				   if we have called include in the frame above - this is the file we have included.
				*/

				ZVAL_STR_COPY(&tmp, include_filename)
				arg_array.GetArr().NextIndexInsertNew(&tmp)
				stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_ARGS).GetStr(), &arg_array)
			}
			ZVAL_INTERNED_STR(&tmp, pseudo_function_name)
			stack_frame.GetArr().KeyAddNew(ZSTR_KNOWN(ZEND_STR_FUNCTION).GetStr(), &tmp)
		}
		return_value.GetArr().NextIndexInsertNew(&stack_frame)
		include_filename = filename
		call = skip
		ptr = skip.GetPrevExecuteData()
	}
}
func ZifDebugBacktrace(execute_data *ZendExecuteData, return_value *Zval) {
	var options ZendLong = DEBUG_BACKTRACE_PROVIDE_OBJECT
	var limit ZendLong = 0
	if ZendParseParameters(ZEND_NUM_ARGS(), "|ll", &options, &limit) == FAILURE {
		return
	}
	ZendFetchDebugBacktrace(return_value, 1, options, limit)
}
func ZifExtensionLoaded(execute_data *ZendExecuteData, return_value *Zval) {
	var extension_name *ZendString
	var lcname *ZendString
	if ZendParseParameters(ZEND_NUM_ARGS(), "S", &extension_name) == FAILURE {
		return
	}
	lcname = ZendStringTolower(extension_name)
	if ZendHashExists(&ModuleRegistry, lcname) != 0 {
		RETVAL_TRUE
	} else {
		RETVAL_FALSE
	}
	ZendStringReleaseEx(lcname, 0)
}
func ZifGetExtensionFuncs(execute_data *ZendExecuteData, return_value *Zval) {
	var extension_name *ZendString
	var lcname *ZendString
	var array int
	var module *ZendModuleEntry
	var zif *ZendFunction
	if ZendParseParameters(ZEND_NUM_ARGS(), "S", &extension_name) == FAILURE {
		return
	}
	if strncasecmp(extension_name.GetVal(), "zend", b.SizeOf("\"zend\"")) {
		lcname = ZendStringTolower(extension_name)
		module = ZendHashFindPtr(&ModuleRegistry, lcname)
		ZendStringReleaseEx(lcname, 0)
	} else {
		module = ZendHashStrFindPtr(&ModuleRegistry, "core", b.SizeOf("\"core\"")-1)
	}
	if module == nil {
		RETVAL_FALSE
		return
	}
	if module.GetFunctions() != nil {

		/* avoid BC break, if functions list is empty, will return an empty array */

		ArrayInit(return_value)
		array = 1
	} else {
		array = 0
	}
	var __ht *HashTable = CG__().GetFunctionTable()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		zif = _z.GetPtr()
		if zif.GetCommonType() == ZEND_INTERNAL_FUNCTION && zif.GetInternalFunction().GetModule() == module {
			if array == 0 {
				ArrayInit(return_value)
				array = 1
			}
			AddNextIndexStr(return_value, zif.GetFunctionName().Copy())
		}
	}
	if array == 0 {
		RETVAL_FALSE
		return
	}
}
