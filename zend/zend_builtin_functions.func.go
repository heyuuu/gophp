package zend

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend/argparse"
	"sik/zend/types"
)

func ZmStartupCore(type_ int, module_number int) int {
	var class_entry ZendClassEntry
	memset(&class_entry, 0, b.SizeOf("zend_class_entry"))
	class_entry.SetName(types.ZendStringInitInterned("stdClass", b.SizeOf("\"stdClass\"")-1, 1))
	class_entry.SetBuiltinFunctions(nil)
	ZendStandardClassDef = ZendRegisterInternalClass(&class_entry)
	ZendRegisterDefaultClasses()
	return types.SUCCESS
}
func ZendStartupBuiltinFunctions() int {
	ZendBuiltinModule.SetModuleNumber(0)
	ZendBuiltinModule.SetType(MODULE_PERSISTENT)
	if b.Assign(&(EG__().GetCurrentModule()), ZendRegisterModuleEx(&ZendBuiltinModule)) == nil {
		return types.FAILURE
	} else {
		return types.SUCCESS
	}
}

//@zif -c 0
func ZifZendVersion() string { return ZEND_VERSION }

//@zif -c 0
func ZifGcMemCaches() int { return ZendMmGc(ZendMmGetHeap()) }

//@zif -c 0
func ZifGcCollectCycles() int { return 0 }

//@zif -c 0
func ZifGcEnabled() bool { return true }

//@zif -c 0
func ZifGcEnable(executeData *ZendExecuteData, return_value *types.Zval) {
	var key *types.ZendString
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	key = types.ZendStringInit("zend.enable_gc", b.SizeOf("\"zend.enable_gc\"")-1, 0)
	ZendAlterIniEntryChars(key, "1", b.SizeOf("\"1\"")-1, ZEND_INI_USER, ZEND_INI_STAGE_RUNTIME)
	types.ZendStringReleaseEx(key, 0)
}
func ZifGcDisable(executeData *ZendExecuteData, return_value *types.Zval) {
	var key *types.ZendString
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	key = types.ZendStringInit("zend.enable_gc", b.SizeOf("\"zend.enable_gc\"")-1, 0)
	ZendAlterIniEntryChars(key, "0", b.SizeOf("\"0\"")-1, ZEND_INI_USER, ZEND_INI_STAGE_RUNTIME)
	types.ZendStringReleaseEx(key, 0)
}
func ZifGcStatus(executeData *ZendExecuteData, return_value *types.Zval) {
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	ArrayInitSize(return_value, 3)
	AddAssocLongEx(return_value, "runs", 0)
	AddAssocLongEx(return_value, "collected", 0)
	AddAssocLongEx(return_value, "threshold", 0)
	AddAssocLongEx(return_value, "roots", 0)
}
func ZifFuncNumArgs(executeData *ZendExecuteData, return_value *types.Zval) {
	var ex *ZendExecuteData = executeData.GetPrevExecuteData()
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_CODE) != 0 {
		ZendError(E_WARNING, "func_num_args():  Called from the global scope - no function context")
		return_value.SetLong(-1)
		return
	}
	if ZendForbidDynamicCall("func_num_args()") == types.FAILURE {
		return_value.SetLong(-1)
		return
	}
	return_value.SetLong(ex.NumArgs())
	return
}
func ZifFuncGetArg(executeData *ZendExecuteData, return_value *types.Zval) {
	var arg_count uint32
	var first_extra_arg uint32
	var arg *types.Zval
	var requested_offset ZendLong
	var ex *ZendExecuteData
	if ZendParseParameters(executeData.NumArgs(), "l", &requested_offset) == types.FAILURE {
		return
	}
	if requested_offset < 0 {
		ZendError(E_WARNING, "func_get_arg():  The argument number should be >= 0")
		return_value.SetFalse()
		return
	}
	ex = executeData.GetPrevExecuteData()
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_CODE) != 0 {
		ZendError(E_WARNING, "func_get_arg():  Called from the global scope - no function context")
		return_value.SetFalse()
		return
	}
	if ZendForbidDynamicCall("func_get_arg()") == types.FAILURE {
		return_value.SetFalse()
		return
	}
	arg_count = ex.NumArgs()
	if ZendUlong(requested_offset >= arg_count) != 0 {
		ZendError(E_WARNING, "func_get_arg():  Argument "+ZEND_LONG_FMT+" not passed to function", requested_offset)
		return_value.SetFalse()
		return
	}
	first_extra_arg = ex.GetFunc().GetOpArray().GetNumArgs()
	if ZendUlong(requested_offset >= first_extra_arg && ex.NumArgs() > first_extra_arg) != 0 {
		arg = ex.VarNum(ex.GetFunc().GetOpArray().GetLastVar()+ex.GetFunc().GetOpArray().GetT()) + (requested_offset - first_extra_arg)
	} else {
		arg = ex.Arg(requested_offset + 1)
	}
	if !(arg.IsUndef()) {
		types.ZVAL_COPY_DEREF(return_value, arg)
	}
}
func ZifFuncGetArgs(executeData *ZendExecuteData, return_value *types.Zval) {
	var p *types.Zval
	var q *types.Zval
	var arg_count uint32
	var first_extra_arg uint32
	var i uint32
	var ex *ZendExecuteData = executeData.GetPrevExecuteData()
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_CODE) != 0 {
		ZendError(E_WARNING, "func_get_args():  Called from the global scope - no function context")
		return_value.SetFalse()
		return
	}
	if ZendForbidDynamicCall("func_get_args()") == types.FAILURE {
		return_value.SetFalse()
		return
	}
	arg_count = ex.NumArgs()
	if arg_count != 0 {
		ArrayInitSize(return_value, arg_count)
		first_extra_arg = ex.GetFunc().GetOpArray().GetNumArgs()
		ZendHashRealInitPacked(return_value.GetArr())
		var __fill_ht *types.HashTable = return_value.GetArr()
		var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
		i = 0
		p = ex.Arg(1)
		if arg_count > first_extra_arg {
			for i < first_extra_arg {
				q = p
				if q.GetTypeInfo() != types.IS_UNDEF {
					q = types.ZVAL_DEREF(q)
					if q.IsRefcounted() {
						q.AddRefcount()
					}
					types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), q)
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
			p = ex.VarNum(ex.GetFunc().GetOpArray().GetLastVar() + ex.GetFunc().GetOpArray().GetT())
		}
		for i < arg_count {
			q = p
			if q.GetTypeInfo() != types.IS_UNDEF {
				q = types.ZVAL_DEREF(q)
				if q.IsRefcounted() {
					q.AddRefcount()
				}
				types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), q)
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
		types.Z_ARRVAL_P(return_value).SetNNumOfElements(arg_count)
	} else {
		ZVAL_EMPTY_ARRAY(return_value)
		return
	}
}
func ZifStrlen(executeData *ZendExecuteData, return_value *types.Zval) {
	var s *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &s, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetLong(s.GetLen())
}
func ZifStrcmp(executeData *ZendExecuteData, return_value *types.Zval) {
	var s1 *types.ZendString
	var s2 *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetLong(ZendBinaryStrcmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen()))
	return
}
func ZifStrncmp(executeData *ZendExecuteData, return_value *types.Zval) {
	var s1 *types.ZendString
	var s2 *types.ZendString
	var len_ ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			len_ = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if len_ < 0 {
		ZendError(E_WARNING, "Length must be greater than or equal to 0")
		return_value.SetFalse()
		return
	}
	return_value.SetLong(ZendBinaryStrncmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen(), len_))
	return
}
func ZifStrcasecmp(executeData *ZendExecuteData, return_value *types.Zval) {
	var s1 *types.ZendString
	var s2 *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	return_value.SetLong(ZendBinaryStrcasecmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen()))
	return
}
func ZifStrncasecmp(executeData *ZendExecuteData, return_value *types.Zval) {
	var s1 *types.ZendString
	var s2 *types.ZendString
	var len_ ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 3
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &s1, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &s2, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			len_ = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if len_ < 0 {
		ZendError(E_WARNING, "Length must be greater than or equal to 0")
		return_value.SetFalse()
		return
	}
	return_value.SetLong(ZendBinaryStrncasecmp(s1.GetVal(), s1.GetLen(), s2.GetVal(), s2.GetLen(), len_))
	return
}
func ZifEach(executeData *ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var entry *types.Zval
	var tmp types.Zval
	var num_key ZendUlong
	var target_hash *types.HashTable
	var key *types.ZendString
	if ZendParseParameters(executeData.NumArgs(), "z/", &array) == types.FAILURE {
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
			return_value.SetFalse()
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

	entry = types.ZVAL_DEREF(entry)
	if entry.IsRefcounted() {
		entry.GetCounted().AddRefcountEx(2)
	}
	return_value.GetArr().IndexAddNewH(1, entry)
	return_value.GetArr().KeyAddNew(types.ZSTR_VALUE.GetStr(), entry)

	/* add the key elements */

	if ZendHashGetCurrentKey(target_hash, &key, &num_key) == HASH_KEY_IS_STRING {
		tmp.SetStringCopy(key)
		tmp.TryAddRefcount()
	} else {
		tmp.SetLong(num_key)
	}
	return_value.GetArr().IndexAddNewH(0, &tmp)
	return_value.GetArr().KeyAddNew(types.ZSTR_KEY.GetStr(), &tmp)
	ZendHashMoveForward(target_hash)
}
func ZifErrorReporting(executeData *ZendExecuteData, return_value *types.Zval) {
	var err *types.Zval = nil
	var old_error_reporting int
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.StartOptional()
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &err, 0)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	old_error_reporting = EG__().GetErrorReporting()
	if executeData.NumArgs() != 0 {
		var new_val *types.ZendString = ZvalTryGetString(err)
		if new_val == nil {
			return
		}
		for {
			var p *ZendIniEntry = EG__().GetErrorReportingIniEntry()
			if p == nil {
				var zv *types.Zval = EG__().GetIniDirectives().KeyFind(types.ZSTR_ERROR_REPORTING.GetStr())
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
				if ZendHashAddPtr(EG__().GetModifiedIniDirectives(), types.ZSTR_ERROR_REPORTING, p) != nil {
					p.SetOrigValue(p.GetValue())
					p.SetOrigModifiable(p.GetModifiable())
					p.SetModified(1)
				}
			} else if p.GetOrigValue() != p.GetValue() {
				types.ZendStringReleaseEx(p.GetValue(), 0)
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
	return_value.SetLong(old_error_reporting)
}
func ValidateConstantArray(ht *types.HashTable) int {
	var ret int = 1
	var val *types.Zval
	ht.ProtectRecursive()
	var __ht *types.HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.GetZv()
			if _z.IsUndef() {
				continue
			}
		}
		val = _z
		val = types.ZVAL_DEREF(val)
		if val.IsRefcounted() {
			if val.IsArray() {
				if val.IsRefcounted() {
					if val.IsRecursive() {
						ZendError(E_WARNING, "Constants cannot be recursive arrays")
						ret = 0
						break
					} else if ValidateConstantArray(val.GetArr()) == 0 {
						ret = 0
						break
					}
				}
			} else if val.GetType() != types.IS_STRING && val.GetType() != types.IS_RESOURCE {
				ZendError(E_WARNING, "Constants may only evaluate to scalar values, arrays or resources")
				ret = 0
				break
			}
		}
	}
	ht.UnprotectRecursive()
	return ret
}
func CopyConstantArray(dst *types.Zval, src *types.Zval) {
	var key *types.ZendString
	var idx ZendUlong
	var new_val *types.Zval
	var val *types.Zval
	ArrayInitSize(dst, types.Z_ARRVAL_P(src).GetNNumOfElements())
	var __ht *types.HashTable = src.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()
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

		val = types.ZVAL_DEREF(val)
		if key != nil {
			new_val = dst.GetArr().KeyAddNew(key.GetStr(), val)
		} else {
			new_val = dst.GetArr().IndexAddNewH(idx, val)
		}
		if val.IsArray() {
			if val.IsRefcounted() {
				CopyConstantArray(new_val, val)
			}
		} else {
			val.TryAddRefcount()
		}
	}
}
func ZifDefine(executeData *ZendExecuteData, return_value *types.Zval) {
	var name *types.ZendString
	var val *types.Zval
	var val_free types.Zval
	var non_cs types.ZendBool = 0
	var case_sensitive int = CONST_CS
	var c ZendConstant
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &val, 0)
			fp.StartOptional()
			non_cs = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if non_cs != 0 {
		case_sensitive = 0
	}
	if ZendMemnstr(name.GetVal(), "::", b.SizeOf("\"::\"")-1, name.GetVal()+name.GetLen()) != nil {
		ZendError(E_WARNING, "Class constants cannot be defined or redefined")
		return_value.SetFalse()
		return
	}
	val_free.SetUndef()
repeat:
	switch val.GetType() {
	case types.IS_LONG:
		fallthrough
	case types.IS_DOUBLE:
		fallthrough
	case types.IS_STRING:
		fallthrough
	case types.IS_FALSE:
		fallthrough
	case types.IS_TRUE:
		fallthrough
	case types.IS_NULL:
		fallthrough
	case types.IS_RESOURCE:

	case types.IS_ARRAY:
		if val.IsRefcounted() {
			if ValidateConstantArray(val.GetArr()) == 0 {
				return_value.SetFalse()
				return
			} else {
				CopyConstantArray(c.GetValue(), val)
				goto register_constant
			}
		}
	case types.IS_OBJECT:
		if val_free.IsUndef() {
			if types.Z_OBJ_HT_P(val).GetGet() != nil {
				val = types.Z_OBJ_HT_P(val).GetGet()(val, &val_free)
				goto repeat
			} else if types.Z_OBJ_HT_P(val).GetCastObject() != nil {
				if types.Z_OBJ_HT_P(val).GetCastObject()(val, &val_free, types.IS_STRING) == types.SUCCESS {
					val = &val_free
					break
				}
			}
		}
		fallthrough
	default:
		ZendError(E_WARNING, "Constants may only evaluate to scalar values, arrays or resources")
		ZvalPtrDtor(&val_free)
		return_value.SetFalse()
		return
	}
	types.ZVAL_COPY(c.GetValue(), val)
	ZvalPtrDtor(&val_free)
register_constant:
	if non_cs != 0 {
		ZendError(E_DEPRECATED, "define(): Declaration of case-insensitive constants is deprecated")
	}

	/* non persistent */

	ZEND_CONSTANT_SET_FLAGS(&c, case_sensitive, PHP_USER_CONSTANT)
	c.SetName(name.Copy())
	if ZendRegisterConstant(&c) == types.SUCCESS {
		return_value.SetTrue()
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifDefined(executeData *ZendExecuteData, return_value *types.Zval) {
	var name *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if ZendGetConstantEx(name, ZendGetExecutedScope(), ZEND_FETCH_CLASS_SILENT|ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK) != nil {
		return_value.SetTrue()
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifGetClass(executeData *ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval = nil
	if ZendParseParameters(executeData.NumArgs(), "|o", &obj) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	if obj == nil {
		var scope *ZendClassEntry = ZendGetExecutedScope()
		if scope != nil {
			return_value.SetStringCopy(scope.GetName())
			return
		} else {
			ZendError(E_WARNING, "get_class() called without object from outside a class")
			return_value.SetFalse()
			return
		}
	}
	return_value.SetStringCopy(types.Z_OBJCE_P(obj).GetName())
	return
}
func ZifGetCalledClass(executeData *ZendExecuteData, return_value *types.Zval) {
	var called_scope *ZendClassEntry
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	called_scope = ZendGetCalledScope(executeData)
	if called_scope != nil {
		return_value.SetStringCopy(called_scope.GetName())
		return
	} else {
		var scope *ZendClassEntry = ZendGetExecutedScope()
		if scope == nil {
			ZendError(E_WARNING, "get_called_class() called from outside a class")
		}
	}
	return_value.SetFalse()
	return
}
func ZifGetParentClass(executeData *ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	var ce *ZendClassEntry = nil
	if ZendParseParameters(executeData.NumArgs(), "|z", &arg) == types.FAILURE {
		return
	}
	if executeData.NumArgs() == 0 {
		ce = ZendGetExecutedScope()
		if ce != nil && ce.GetParent() {
			return_value.SetStringCopy(ce.GetParent().name)
			return
		} else {
			return_value.SetFalse()
			return
		}
	}
	if arg.IsObject() {
		ce = types.Z_OBJ_P(arg).GetCe()
	} else if arg.IsString() {
		ce = ZendLookupClass(arg.GetStr())
	}
	if ce != nil && ce.GetParent() {
		return_value.SetStringCopy(ce.GetParent().name)
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func IsAImpl(executeData *ZendExecuteData, return_value *types.Zval, only_subclass types.ZendBool) {
	var obj *types.Zval
	var class_name *types.ZendString
	var instance_ce *ZendClassEntry
	var ce *ZendClassEntry
	var allow_string types.ZendBool = only_subclass
	var retval types.ZendBool
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &obj, 0)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &class_name, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			allow_string = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
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
			return_value.SetFalse()
			return
		}
	} else if obj.IsObject() {
		instance_ce = types.Z_OBJCE_P(obj)
	} else {
		return_value.SetFalse()
		return
	}
	if only_subclass == 0 && types.ZendStringEquals(instance_ce.GetName(), class_name) != 0 {
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
	types.ZVAL_BOOL(return_value, retval != 0)
	return
}
func ZifIsSubclassOf(executeData *ZendExecuteData, return_value *types.Zval) {
	IsAImpl(executeData, return_value, 1)
}
func ZifIsA(executeData *ZendExecuteData, return_value *types.Zval) {
	IsAImpl(executeData, return_value, 0)
}
func AddClassVars(scope *ZendClassEntry, ce *ZendClassEntry, statics int, return_value *types.Zval) {
	var prop_info *ZendPropertyInfo
	var prop *types.Zval
	var prop_copy types.Zval
	var key *types.ZendString
	var __ht *types.HashTable = ce.GetPropertiesInfo()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		key = _p.GetKey()
		prop_info = _z.GetPtr()
		if prop_info.IsProtected() && ZendCheckProtected(prop_info.GetCe(), scope) == 0 || prop_info.IsPrivate() && prop_info.GetCe() != scope {
			continue
		}
		prop = nil
		if statics != 0 && prop_info.IsStatic() {
			prop = ce.GetDefaultStaticMembersTable()[prop_info.GetOffset()]
			prop = types.ZVAL_DEINDIRECT(prop)
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

			types.ZVAL_COPY_OR_DUP(&prop_copy, prop)

			/* copy: enforce read only access */

		}
		prop = &prop_copy

		/* this is necessary to make it able to work with default array
		 * properties, returned to user */

		if prop.IsConstant() {
			if ZvalUpdateConstantEx(prop, nil) != types.SUCCESS {
				return
			}
		}
		return_value.GetArr().KeyAddNew(key.GetStr(), prop)
	}
}
func ZifGetClassVars(executeData *ZendExecuteData, return_value *types.Zval) {
	var class_name *types.ZendString
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	if ZendParseParameters(executeData.NumArgs(), "S", &class_name) == types.FAILURE {
		return
	}
	ce = ZendLookupClass(class_name)
	if ce == nil {
		return_value.SetFalse()
		return
	} else {
		ArrayInit(return_value)
		if !ce.IsConstantsUpdated() {
			if ZendUpdateClassConstants(ce) != types.SUCCESS {
				return
			}
		}
		scope = ZendGetExecutedScope()
		AddClassVars(scope, ce, 0, return_value)
		AddClassVars(scope, ce, 1, return_value)
	}
}
func ZifGetObjectVars(executeData *ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var value *types.Zval
	var properties *types.HashTable
	var key *types.ZendString
	var zobj *types.ZendObject
	var num_key ZendUlong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			obj = fp.ParseObject()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	properties = types.Z_OBJ_HT_P(obj).GetGetProperties()(obj)
	if properties == nil {
		return_value.SetFalse()
		return
	}
	zobj = obj.GetObj()
	if zobj.GetCe().GetDefaultPropertiesCount() == 0 && properties == zobj.GetProperties() && !(properties.IsRecursive()) {

		/* fast copy */

		if zobj.GetHandlers() == &StdObjectHandlers {
			return_value.SetArray(ZendProptableToSymtable(properties, 0))
			return
		}
		return_value.SetArray(ZendProptableToSymtable(properties, 1))
		return
	} else {
		ArrayInitSize(return_value, properties.GetNNumOfElements())
		var __ht *types.HashTable = properties
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			num_key = _p.GetH()
			key = _p.GetKey()
			value = _z
			var is_dynamic types.ZendBool = 1
			if value.IsIndirect() {
				value = value.GetZv()
				if value.IsUndef() {
					continue
				}
				is_dynamic = 0
			}
			if key != nil && ZendCheckPropertyAccess(zobj, key, is_dynamic) == types.FAILURE {
				continue
			}
			if value.IsReference() && value.GetRefcount() == 1 {
				value = types.Z_REFVAL_P(value)
			}
			value.TryAddRefcount()
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
				return_value.GetArr().SymtableAddNew(key.GetStr(), value)
			}
		}
	}
}
func ZifGetMangledObjectVars(executeData *ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var properties *types.HashTable
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			obj = fp.ParseObject()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	properties = types.Z_OBJ_HT_P(obj).GetGetProperties()(obj)
	if properties == nil {
		ZVAL_EMPTY_ARRAY(return_value)
		return
	}
	properties = ZendProptableToSymtable(properties, types.Z_OBJCE_P(obj).GetDefaultPropertiesCount() != 0 || types.Z_OBJ_P(obj).GetHandlers() != &StdObjectHandlers || properties.IsRecursive())
	return_value.SetArray(properties)
	return
}
func SameName(key *types.ZendString, name *types.ZendString) int {
	var lcname *types.ZendString
	var ret int
	if key == name {
		return 1
	}
	if key.GetLen() != name.GetLen() {
		return 0
	}
	lcname = ZendStringTolower(name)
	ret = memcmp(lcname.GetVal(), key.GetVal(), key.GetLen()) == 0
	types.ZendStringReleaseEx(lcname, 0)
	return ret
}
func ZifGetClassMethods(executeData *ZendExecuteData, return_value *types.Zval) {
	var klass *types.Zval
	var method_name types.Zval
	var ce *ZendClassEntry = nil
	var scope *ZendClassEntry
	var mptr *ZendFunction
	var key *types.ZendString
	if ZendParseParameters(executeData.NumArgs(), "z", &klass) == types.FAILURE {
		return
	}
	if klass.IsObject() {
		ce = types.Z_OBJCE_P(klass)
	} else if klass.IsString() {
		ce = ZendLookupClass(klass.GetStr())
	}
	if ce == nil {
		return_value.SetNull()
		return
	}
	ArrayInit(return_value)
	scope = ZendGetExecutedScope()
	var __ht *types.HashTable = ce.GetFunctionTable()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		key = _p.GetKey()
		mptr = _z.GetPtr()
		if mptr.IsPublic() || scope != nil && (mptr.IsProtected() && ZendCheckProtected(mptr.GetScope(), scope) != 0 || mptr.IsPrivate() && scope == mptr.GetScope()) {
			if mptr.GetType() == ZEND_USER_FUNCTION && (mptr.GetOpArray().GetRefcount() == nil || mptr.op_array.refcount > 1) && key != nil && SameName(key, mptr.GetFunctionName()) == 0 {
				method_name.SetStringCopy(ZendFindAliasName(mptr.GetScope(), key))
				return_value.GetArr().NextIndexInsertNew(&method_name)
			} else {
				method_name.SetStringCopy(mptr.GetFunctionName())
				return_value.GetArr().NextIndexInsertNew(&method_name)
			}
		}
	}
}
func ZifMethodExists(executeData *ZendExecuteData, return_value *types.Zval) {
	var klass *types.Zval
	var method_name *types.ZendString
	var lcname *types.ZendString
	var ce *ZendClassEntry
	var func_ *ZendFunction
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			argparse.ZendParseArgZvalDeref(_arg, &klass, 0)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &method_name, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if klass.IsObject() {
		ce = types.Z_OBJCE_P(klass)
	} else if klass.IsString() {
		if b.Assign(&ce, ZendLookupClass(klass.GetStr())) == nil {
			return_value.SetFalse()
			return
		}
	} else {
		return_value.SetFalse()
		return
	}
	lcname = ZendStringTolower(method_name)
	func_ = ZendHashFindPtr(ce.GetFunctionTable(), lcname)
	types.ZendStringReleaseEx(lcname, 0)
	if func_ != nil {

		/* Exclude shadow properties when checking a method on a specific class. Include
		 * them when checking an object, as method_exists() generally ignores visibility.
		 * TODO: Should we use EG(scope) for the object case instead? */

		types.ZVAL_BOOL(return_value, klass.IsObject() || !func_.IsPrivate() || func_.GetScope() == ce)
		return
	}
	if klass.IsObject() {
		var obj *types.ZendObject = klass.GetObj()
		func_ = types.Z_OBJ_HT_P(klass).GetGetMethod()(&obj, method_name, nil)
		if func_ != nil {
			if func_.IsCallViaTrampoline() {

				/* Returns true to the fake Closure's __invoke */

				types.ZVAL_BOOL(return_value, func_.GetScope() == ZendCeClosure && types.ZendStringEqualsLiteral(method_name, ZEND_INVOKE_FUNC_NAME))
				types.ZendStringReleaseEx(func_.GetFunctionName(), 0)
				ZendFreeTrampoline(func_)
				return
			}
			return_value.SetTrue()
			return
		}
	}
	return_value.SetFalse()
	return
}
func ZifPropertyExists(executeData *ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval
	var property *types.ZendString
	var ce *ZendClassEntry
	var property_info *ZendPropertyInfo
	var property_z types.Zval
	if ZendParseParameters(executeData.NumArgs(), "zS", &object, &property) == types.FAILURE {
		return
	}
	if property == nil {
		return_value.SetFalse()
		return
	}
	if object.IsString() {
		ce = ZendLookupClass(object.GetStr())
		if ce == nil {
			return_value.SetFalse()
			return
		}
	} else if object.IsObject() {
		ce = types.Z_OBJCE_P(object)
	} else {
		ZendError(E_WARNING, "First parameter must either be an object or the name of an existing class")
		return_value.SetNull()
		return
	}
	property_info = ZendHashFindPtr(ce.GetPropertiesInfo(), property)
	if property_info != nil && (!property_info.IsPrivate() || property_info.GetCe() == ce) {
		return_value.SetTrue()
		return
	}
	property_z.SetString(property)
	if object.IsObject() && types.Z_OBJ_HT(*object).GetHasProperty()(object, &property_z, 2, nil) != 0 {
		return_value.SetTrue()
		return
	}
	return_value.SetFalse()
	return
}
func ClassExistsImpl(executeData *ZendExecuteData, return_value *types.Zval, flags int, skip_flags int) {
	var name *types.ZendString
	var lcname *types.ZendString
	var ce *ZendClassEntry
	var autoload types.ZendBool = 1
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			autoload = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if autoload == 0 {
		if name.GetVal()[0] == '\\' {

			/* Ignore leading "\" */

			lcname = types.ZendStringAlloc(name.GetLen()-1, 0)
			ZendStrTolowerCopy(lcname.GetVal(), name.GetVal()+1, name.GetLen()-1)
		} else {
			lcname = ZendStringTolower(name)
		}
		ce = ZendHashFindPtr(EG__().GetClassTable(), lcname)
		types.ZendStringReleaseEx(lcname, 0)
	} else {
		ce = ZendLookupClass(name)
	}
	if ce != nil {
		types.ZVAL_BOOL(return_value, (ce.GetCeFlags()&flags) == flags && !ce.HasCeFlags(skip_flags))
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifClassExists(executeData *ZendExecuteData, return_value *types.Zval) {
	ClassExistsImpl(executeData, return_value, ZEND_ACC_LINKED, ZEND_ACC_INTERFACE|ZEND_ACC_TRAIT)
}
func ZifInterfaceExists(executeData *ZendExecuteData, return_value *types.Zval) {
	ClassExistsImpl(executeData, return_value, ZEND_ACC_LINKED|ZEND_ACC_INTERFACE, 0)
}
func ZifTraitExists(executeData *ZendExecuteData, return_value *types.Zval) {
	ClassExistsImpl(executeData, return_value, ZEND_ACC_TRAIT, 0)
}
func ZifFunctionExists(executeData *ZendExecuteData, return_value *types.Zval) {
	var name *types.ZendString
	var func_ *ZendFunction
	var lcname *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			argparse.Z_PARAM_PROLOGUE(0, 0)
			if argparse.ZendParseArgStr(_arg, &name, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if name.GetVal()[0] == '\\' {

		/* Ignore leading "\" */

		lcname = types.ZendStringAlloc(name.GetLen()-1, 0)
		ZendStrTolowerCopy(lcname.GetVal(), name.GetVal()+1, name.GetLen()-1)
	} else {
		lcname = ZendStringTolower(name)
	}
	func_ = ZendHashFindPtr(EG__().GetFunctionTable(), lcname)
	types.ZendStringReleaseEx(lcname, 0)

	/*
	 * A bit of a hack, but not a bad one: we see if the handler of the function
	 * is actually one that displays "function is disabled" message.
	 */

	types.ZVAL_BOOL(return_value, func_ != nil && (func_.GetType() != ZEND_INTERNAL_FUNCTION || func_.GetInternalFunction().GetHandler() != ZifDisplayDisabledFunction))
	return
}
func ZifClassAlias(executeData *ZendExecuteData, return_value *types.Zval) {
	var class_name *types.ZendString
	var alias_name *byte
	var ce *ZendClassEntry
	var alias_name_len int
	var autoload types.ZendBool = 1
	if ZendParseParameters(executeData.NumArgs(), "Ss|b", &class_name, &alias_name, &alias_name_len, &autoload) == types.FAILURE {
		return
	}
	ce = ZendLookupClassEx(class_name, nil, b.Cond(autoload == 0, ZEND_FETCH_CLASS_NO_AUTOLOAD, 0))
	if ce != nil {
		if ce.GetType() == ZEND_USER_CLASS {
			if ZendRegisterClassAliasEx(alias_name, alias_name_len, ce, 0) == types.SUCCESS {
				return_value.SetTrue()
				return
			} else {
				ZendError(E_WARNING, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), alias_name)
				return_value.SetFalse()
				return
			}
		} else {
			ZendError(E_WARNING, "First argument of class_alias() must be a name of user defined class")
			return_value.SetFalse()
			return
		}
	} else {
		ZendError(E_WARNING, "Class '%s' not found", class_name.GetVal())
		return_value.SetFalse()
		return
	}
}
func ZifGetIncludedFiles(executeData *ZendExecuteData, return_value *types.Zval) {
	var entry *types.ZendString
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	ArrayInit(return_value)
	var __ht *types.HashTable = EG__().GetIncludedFiles()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		entry = _p.GetKey()
		if entry != nil {
			AddNextIndexStr(return_value, entry.Copy())
		}
	}
}
func ZifTriggerError(executeData *ZendExecuteData, return_value *types.Zval) {
	var error_type ZendLong = E_USER_NOTICE
	var message *byte
	var message_len int
	if ZendParseParameters(executeData.NumArgs(), "s|l", &message, &message_len, &error_type) == types.FAILURE {
		return
	}
	switch error_type {
	case E_USER_ERROR:
		fallthrough
	case E_USER_WARNING:
		fallthrough
	case E_USER_NOTICE:
		fallthrough
	case E_USER_DEPRECATED:

	default:
		ZendError(E_WARNING, "Invalid error type specified")
		return_value.SetFalse()
		return
	}
	ZendError(int(error_type), "%s", message)
	return_value.SetTrue()
	return
}
func ZifSetErrorHandler(executeData *ZendExecuteData, return_value *types.Zval) {
	var error_handler *types.Zval
	var error_type ZendLong = E_ALL
	if ZendParseParameters(executeData.NumArgs(), "z|l", &error_handler, &error_type) == types.FAILURE {
		return
	}
	if error_handler.GetType() != types.IS_NULL {
		if ZendIsCallable(error_handler, 0, nil) == 0 {
			var error_handler_name *types.ZendString = ZendGetCallableName(error_handler)
			ZendError(E_WARNING, "%s() expects the argument (%s) to be a valid callback", GetActiveFunctionName(), b.CondF1(error_handler_name != nil, func() []byte { return error_handler_name.GetVal() }, "unknown"))
			types.ZendStringReleaseEx(error_handler_name, 0)
			return
		}
	}
	if EG__().GetUserErrorHandler().GetType() != types.IS_UNDEF {
		types.ZVAL_COPY(return_value, EG__().GetUserErrorHandler())
	}
	EG__().GetUserErrorHandlersErrorReporting().Push(EG__().GetUserErrorHandlerErrorReporting())
	EG__().GetUserErrorHandlers().Push(EG__().GetUserErrorHandler())
	if error_handler.IsNull() {
		EG__().GetUserErrorHandler().SetUndef()
		return
	}
	types.ZVAL_COPY(EG__().GetUserErrorHandler(), error_handler)
	EG__().SetUserErrorHandlerErrorReporting(int(error_type))
}
func ZifRestoreErrorHandler(executeData *ZendExecuteData, return_value *types.Zval) {
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	if EG__().GetUserErrorHandler().GetType() != types.IS_UNDEF {
		var zeh types.Zval
		types.ZVAL_COPY_VALUE(&zeh, EG__().GetUserErrorHandler())
		EG__().GetUserErrorHandler().SetUndef()
		ZvalPtrDtor(&zeh)
	}
	if ZendStackIsEmpty(EG__().GetUserErrorHandlers()) != 0 {
		EG__().GetUserErrorHandler().SetUndef()
	} else {
		var tmp *types.Zval
		EG__().SetUserErrorHandlerErrorReporting(ZendStackIntTop(EG__().GetUserErrorHandlersErrorReporting()))
		EG__().GetUserErrorHandlersErrorReporting().DelTop()
		tmp = EG__().GetUserErrorHandlers().Top()
		types.ZVAL_COPY_VALUE(EG__().GetUserErrorHandler(), tmp)
		EG__().GetUserErrorHandlers().DelTop()
	}
	return_value.SetTrue()
	return
}
func ZifSetExceptionHandler(executeData *ZendExecuteData, return_value *types.Zval) {
	var exception_handler *types.Zval
	if ZendParseParameters(executeData.NumArgs(), "z", &exception_handler) == types.FAILURE {
		return
	}
	if exception_handler.GetType() != types.IS_NULL {
		if ZendIsCallable(exception_handler, 0, nil) == 0 {
			var exception_handler_name *types.ZendString = ZendGetCallableName(exception_handler)
			ZendError(E_WARNING, "%s() expects the argument (%s) to be a valid callback", GetActiveFunctionName(), b.CondF1(exception_handler_name != nil, func() []byte { return exception_handler_name.GetVal() }, "unknown"))
			types.ZendStringReleaseEx(exception_handler_name, 0)
			return
		}
	}
	if EG__().GetUserExceptionHandler().GetType() != types.IS_UNDEF {
		types.ZVAL_COPY(return_value, EG__().GetUserExceptionHandler())
	}
	EG__().GetUserExceptionHandlers().Push(EG__().GetUserExceptionHandler())
	if exception_handler.IsNull() {
		EG__().GetUserExceptionHandler().SetUndef()
		return
	}
	types.ZVAL_COPY(EG__().GetUserExceptionHandler(), exception_handler)
}
func ZifRestoreExceptionHandler(executeData *ZendExecuteData, return_value *types.Zval) {
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	if EG__().GetUserExceptionHandler().GetType() != types.IS_UNDEF {
		ZvalPtrDtor(EG__().GetUserExceptionHandler())
	}
	if ZendStackIsEmpty(EG__().GetUserExceptionHandlers()) != 0 {
		EG__().GetUserExceptionHandler().SetUndef()
	} else {
		var tmp *types.Zval = EG__().GetUserExceptionHandlers().Top()
		types.ZVAL_COPY_VALUE(EG__().GetUserExceptionHandler(), tmp)
		EG__().GetUserExceptionHandlers().DelTop()
	}
	return_value.SetTrue()
	return
}
func CopyClassOrInterfaceName(array *types.Zval, key *types.ZendString, ce *ZendClassEntry) {
	if ce.GetRefcount() == 1 && !ce.IsImmutable() || SameName(key, ce.GetName()) != 0 {
		key = ce.GetName()
	}
	AddNextIndexStr(array, key.Copy())
}
func GetDeclaredClassImpl(executeData *ZendExecuteData, return_value *types.Zval, flags int, skip_flags int) {
	var key *types.ZendString
	var ce *ZendClassEntry
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	ArrayInit(return_value)
	var __ht *types.HashTable = EG__().GetClassTable()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		key = _p.GetKey()
		ce = _z.GetPtr()
		if key != nil && key.GetVal()[0] != 0 && ce.HasCeFlags(flags) && !ce.HasCeFlags(skip_flags) {
			CopyClassOrInterfaceName(return_value, key, ce)
		}
	}
}
func ZifGetDeclaredTraits(executeData *ZendExecuteData, return_value *types.Zval) {
	GetDeclaredClassImpl(executeData, return_value, ZEND_ACC_TRAIT, 0)
}
func ZifGetDeclaredClasses(executeData *ZendExecuteData, return_value *types.Zval) {
	GetDeclaredClassImpl(executeData, return_value, ZEND_ACC_LINKED, ZEND_ACC_INTERFACE|ZEND_ACC_TRAIT)
}
func ZifGetDeclaredInterfaces(executeData *ZendExecuteData, return_value *types.Zval) {
	GetDeclaredClassImpl(executeData, return_value, ZEND_ACC_INTERFACE, 0)
}
func ZifGetDefinedFunctions(executeData *ZendExecuteData, return_value *types.Zval) {
	var internal types.Zval
	var user types.Zval
	var key *types.ZendString
	var func_ *ZendFunction
	var exclude_disabled types.ZendBool = 0
	if ZendParseParameters(executeData.NumArgs(), "|b", &exclude_disabled) == types.FAILURE {
		return
	}
	ArrayInit(&internal)
	ArrayInit(&user)
	ArrayInit(return_value)
	var __ht *types.HashTable = EG__().GetFunctionTable()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

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
	return_value.GetArr().KeyAddNew("internal", &internal)
	return_value.GetArr().KeyAddNew("user", &user)
}
func ZifGetDefinedVars(executeData *ZendExecuteData, return_value *types.Zval) {
	var symbol_table *types.ZendArray
	if ZendForbidDynamicCall("get_defined_vars()") == types.FAILURE {
		return
	}
	symbol_table = ZendRebuildSymbolTable()
	if symbol_table == nil {
		return
	}
	return_value.SetArray(ZendArrayDup(symbol_table))
	return
}
func ZifCreateFunction(executeData *ZendExecuteData, return_value *types.Zval) {
	var function_name *types.ZendString
	var eval_code *byte
	var function_args *byte
	var function_code *byte
	var eval_code_length int
	var function_args_len int
	var function_code_len int
	var retval int
	var eval_name *byte
	if ZendParseParameters(executeData.NumArgs(), "ss", &function_args, &function_args_len, &function_code, &function_code_len) == types.FAILURE {
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
	if retval == types.SUCCESS {
		var func_ *ZendOpArray
		var static_variables *types.HashTable
		func_ = ZendHashStrFindPtr(EG__().GetFunctionTable(), LAMBDA_TEMP_FUNCNAME, b.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		if func_ == nil {
			ZendErrorNoreturn(E_CORE_ERROR, "Unexpected inconsistency in create_function()")
			return_value.SetFalse()
			return
		}
		if func_.GetRefcount() != nil {
			func_.refcount++
		}
		static_variables = func_.GetStaticVariables()
		func_.SetStaticVariables(nil)
		ZendHashStrDel(EG__().GetFunctionTable(), LAMBDA_TEMP_FUNCNAME, b.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		func_.SetStaticVariables(static_variables)
		function_name = types.ZendStringAlloc(b.SizeOf("\"0lambda_\"")+MAX_LENGTH_OF_LONG, 0)
		function_name.GetVal()[0] = '0'
		for {
			function_name.SetLen(core.Snprintf(function_name.GetVal()+1, b.SizeOf("\"lambda_\"")+MAX_LENGTH_OF_LONG, "lambda_%d", b.PreInc(&(EG__().GetLambdaCount()))) + 1)
			if ZendHashAddPtr(EG__().GetFunctionTable(), function_name, func_) != nil {
				break
			}
		}
		return_value.SetString(function_name)
		return
	} else {
		ZendHashStrDel(EG__().GetFunctionTable(), LAMBDA_TEMP_FUNCNAME, b.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		return_value.SetFalse()
		return
	}
}
func ZifGetResourceType(executeData *ZendExecuteData, return_value *types.Zval) {
	var resource_type *byte
	var z_resource_type *types.Zval
	if ZendParseParameters(executeData.NumArgs(), "r", &z_resource_type) == types.FAILURE {
		return
	}
	resource_type = ZendRsrcListGetRsrcType(z_resource_type.GetRes())
	if resource_type != nil {
		return_value.SetRawString(b.CastStrAuto(resource_type))
		return
	} else {
		return_value.SetRawString("Unknown")
		return
	}
}
func ZifGetResources(executeData *ZendExecuteData, return_value *types.Zval) {
	var type_ *types.ZendString = nil
	var key *types.ZendString
	var index ZendUlong
	var val *types.Zval
	if ZendParseParameters(executeData.NumArgs(), "|S", &type_) == types.FAILURE {
		return
	}
	if type_ == nil {
		ArrayInit(return_value)
		var __ht *types.HashTable = EG__().GetRegularList()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			if key == nil {
				val.AddRefcount()
				return_value.GetArr().IndexAddNewH(index, val)
			}
		}
	} else if types.ZendStringEqualsLiteral(type_, "Unknown") {
		ArrayInit(return_value)
		var __ht *types.HashTable = EG__().GetRegularList()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			if key == nil && types.Z_RES_TYPE_P(val) <= 0 {
				val.AddRefcount()
				return_value.GetArr().IndexAddNewH(index, val)
			}
		}
	} else {
		var id int = ZendFetchListDtorId(type_.GetVal())
		if id <= 0 {
			ZendError(E_WARNING, "get_resources():  Unknown resource type '%s'", type_.GetVal())
			return_value.SetFalse()
			return
		}
		ArrayInit(return_value)
		var __ht *types.HashTable = EG__().GetRegularList()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			if key == nil && types.Z_RES_TYPE_P(val) == id {
				val.AddRefcount()
				return_value.GetArr().IndexAddNewH(index, val)
			}
		}
	}
}
func AddZendextInfo(ext *ZendExtension, arg any) int {
	var name_array *types.Zval = (*types.Zval)(arg)
	AddNextIndexString(name_array, ext.GetName())
	return 0
}
func ZifGetLoadedExtensions(executeData *ZendExecuteData, return_value *types.Zval) {
	var zendext types.ZendBool = 0
	if ZendParseParameters(executeData.NumArgs(), "|b", &zendext) == types.FAILURE {
		return
	}
	ArrayInit(return_value)
	if zendext != 0 {
		ZendExtensions.ApplyWithArgument(LlistApplyWithArgFuncT(AddZendextInfo), return_value)
	} else {
		var module *ZendModuleEntry
		var __ht *types.HashTable = &ModuleRegistry
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			module = _z.GetPtr()
			AddNextIndexString(return_value, module.GetName())
		}
	}
}
func ZifGetDefinedConstants(executeData *ZendExecuteData, return_value *types.Zval) {
	var categorize types.ZendBool = 0
	if ZendParseParameters(executeData.NumArgs(), "|b", &categorize) == types.FAILURE {
		return
	}
	ArrayInit(return_value)
	if categorize != 0 {
		var val *ZendConstant
		var module_number int
		var modules *types.Zval
		var const_val types.Zval
		var module_names **byte
		var module *ZendModuleEntry
		var i int = 1
		modules = Ecalloc(ModuleRegistry.GetNNumOfElements()+2, b.SizeOf("zval"))
		module_names = Emalloc((ModuleRegistry.GetNNumOfElements() + 2) * b.SizeOf("char *"))
		module_names[0] = "internal"
		var __ht *types.HashTable = &ModuleRegistry
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			module = _z.GetPtr()
			module_names[module.GetModuleNumber()] = (*byte)(module.GetName())
			i++
		}
		module_names[i] = "user"
		var __ht__1 *types.HashTable = EG__().GetZendConstants()
		for _, _p := range __ht__1.foreachData() {
			var _z *types.Zval = _p.GetVal()

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
			types.ZVAL_COPY_OR_DUP(&const_val, val.GetValue())
			modules[module_number].GetArr().KeyAddNew(val.GetName().GetStr(), &const_val)
		}
		Efree(module_names)
		Efree(modules)
	} else {
		var constant *ZendConstant
		var const_val types.Zval
		var __ht *types.HashTable = EG__().GetZendConstants()
		for _, _p := range __ht.foreachData() {
			var _z *types.Zval = _p.GetVal()

			constant = _z.GetPtr()
			if constant.GetName() == nil {

				/* skip special constants */

				continue

				/* skip special constants */

			}
			types.ZVAL_COPY_OR_DUP(&const_val, constant.GetValue())
			return_value.GetArr().KeyAddNew(constant.GetName().GetStr(), &const_val)
		}
	}
}
func DebugBacktraceGetArgs(call *ZendExecuteData, arg_array *types.Zval) {
	var num_args uint32 = call.NumArgs()
	if num_args != 0 {
		var i uint32 = 0
		var p *types.Zval = call.Arg(1)
		ArrayInitSize(arg_array, num_args)
		ZendHashRealInitPacked(arg_array.GetArr())
		var __fill_ht *types.HashTable = arg_array.GetArr()
		var __fill_bkt *types.Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		b.Assert(__fill_ht.HasUFlags(types.HASH_FLAG_PACKED))
		if call.GetFunc().GetType() == ZEND_USER_FUNCTION {
			var first_extra_arg uint32 = b.Min(num_args, call.GetFunc().GetOpArray().GetNumArgs())
			if (ZEND_CALL_INFO(call) & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {

				/* In case of attached symbol_table, values on stack may be invalid
				 * and we have to access them through symbol_table
				 * See: https://bugs.php.net/bug.php?id=73156
				 */

				var arg_name *types.ZendString
				var arg *types.Zval
				for i < first_extra_arg {
					arg_name = call.GetFunc().GetOpArray().GetVars()[i]
					arg = ZendHashFindExInd(call.GetSymbolTable(), arg_name, 1)
					if arg != nil {
						if arg.IsRefcounted() {
							arg.AddRefcount()
						}
						types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), arg)
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
					if p.GetTypeInfo() != types.IS_UNDEF {
						if p.IsRefcounted() {
							p.AddRefcount()
						}
						types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), p)
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
			p = call.VarNum(call.GetFunc().GetOpArray().GetLastVar() + call.GetFunc().GetOpArray().GetT())
		}
		for i < num_args {
			if p.GetTypeInfo() != types.IS_UNDEF {
				if p.IsRefcounted() {
					p.AddRefcount()
				}
				types.ZVAL_COPY_VALUE(__fill_bkt.GetVal(), p)
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
		types.Z_ARRVAL_P(arg_array).SetNNumOfElements(num_args)
	} else {
		ZVAL_EMPTY_ARRAY(arg_array)
	}
}
func DebugPrintBacktraceArgs(arg_array *types.Zval) {
	var tmp *types.Zval
	var i int = 0
	var __ht *types.HashTable = arg_array.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		tmp = _z
		if b.PostInc(&i) {
			ZEND_PUTS(", ")
		}
		ZendPrintFlatZvalR(tmp)
	}
}
func SkipInternalHandler(skip *ZendExecuteData) types.ZendBool {
	return !(skip.GetFunc() != nil && ZEND_USER_CODE(skip.GetFunc().GetCommonType())) && skip.GetPrevExecuteData() != nil && skip.GetPrevExecuteData().GetFunc() != nil && ZEND_USER_CODE(skip.GetPrevExecuteData().GetFunc().GetCommonType()) && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_ICALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_UCALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL_BY_NAME && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL
}
func ZifDebugPrintBacktrace(executeData *ZendExecuteData, return_value *types.Zval) {
	var call *ZendExecuteData
	var ptr *ZendExecuteData
	var skip *ZendExecuteData
	var object *types.ZendObject
	var lineno int
	var frameno int = 0
	var func_ *ZendFunction
	var function_name *byte
	var filename *byte
	var class_name *types.ZendString = nil
	var call_type *byte
	var include_filename *byte = nil
	var arg_array types.Zval
	var indent int = 0
	var options ZendLong = 0
	var limit ZendLong = 0
	if ZendParseParameters(executeData.NumArgs(), "|ll", &options, &limit) == types.FAILURE {
		return
	}
	arg_array.SetUndef()
	ptr = executeData.GetPrevExecuteData()

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
			var zend_function_name *types.ZendString
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

			var build_filename_arg types.ZendBool = 1
			if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetCommonType())) || ptr.GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL {

				/* can happen when calling eval from a custom sapi */

				function_name = "unknown"
				build_filename_arg = 0
			} else {
				switch ptr.GetOpline().GetExtendedValue() {
				case ZEND_EVAL:
					function_name = "eval"
					build_filename_arg = 0
				case ZEND_INCLUDE:
					function_name = "include"
				case ZEND_REQUIRE:
					function_name = "require"
				case ZEND_INCLUDE_ONCE:
					function_name = "include_once"
				case ZEND_REQUIRE_ONCE:
					function_name = "require_once"
				default:

					/* this can actually happen if you use debug_backtrace() in your error_handler and
					 * you're in the top-scope */

					function_name = "unknown"
					build_filename_arg = 0
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
			ZEND_PUTS(class_name.GetStr())
			ZEND_PUTS(call_type)
			if object != nil && func_.GetScope() == nil && object.GetHandlers().GetGetClassName() != ZendStdGetClassName {
				types.ZendStringReleaseEx(class_name, 0)
			}
		}
		ZendPrintf("%s(", function_name)
		if arg_array.GetType() != types.IS_UNDEF {
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
func ZendFetchDebugBacktrace(return_value *types.Zval, skip_last int, options int, limit int) {
	var ptr *ZendExecuteData
	var skip *ZendExecuteData
	var call *ZendExecuteData = nil
	var object *types.ZendObject
	var lineno int
	var frameno int = 0
	var func_ *ZendFunction
	var function_name *types.ZendString
	var filename *types.ZendString
	var include_filename *types.ZendString = nil
	var stack_frame types.Zval
	var tmp types.Zval
	ArrayInit(return_value)
	if !(b.Assign(&ptr, CurrEX())) {
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
			tmp.SetStringCopy(filename)
			stack_frame.GetArr().KeyAddNew(types.ZSTR_FILE.GetStr(), &tmp)
			tmp.SetLong(lineno)
			stack_frame.GetArr().KeyAddNew(types.ZSTR_LINE.GetStr(), &tmp)
		} else {
			var prev_call *ZendExecuteData = skip
			var prev *ZendExecuteData = skip.GetPrevExecuteData()
			for prev != nil {
				if prev_call != nil && prev_call.GetFunc() != nil && !(ZEND_USER_CODE(prev_call.GetFunc().GetCommonType())) && !prev_call.GetFunc().IsCallViaTrampoline() {
					break
				}
				if prev.GetFunc() != nil && ZEND_USER_CODE(prev.GetFunc().GetCommonType()) {
					tmp.SetStringCopy(prev.GetFunc().GetOpArray().GetFilename())
					stack_frame.GetArr().KeyAddNew(types.ZSTR_FILE.GetStr(), &tmp)
					tmp.SetLong(prev.GetOpline().GetLineno())
					stack_frame.GetArr().KeyAddNew(types.ZSTR_LINE.GetStr(), &tmp)
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
			tmp.SetStringCopy(function_name)
			stack_frame.GetArr().KeyAddNew(types.ZSTR_FUNCTION.GetStr(), &tmp)
			if object != nil {
				if func_.GetScope() != nil {
					tmp.SetStringCopy(func_.GetScope().GetName())
				} else if object.GetHandlers().GetGetClassName() == ZendStdGetClassName {
					tmp.SetStringCopy(object.GetCe().GetName())
				} else {
					tmp.SetString(object.GetHandlers().GetGetClassName()(object))
				}
				stack_frame.GetArr().KeyAddNew(types.ZSTR_CLASS.GetStr(), &tmp)
				if (options & DEBUG_BACKTRACE_PROVIDE_OBJECT) != 0 {
					tmp.SetObject(object)
					stack_frame.GetArr().KeyAddNew(types.ZSTR_OBJECT.GetStr(), &tmp)
					tmp.AddRefcount()
				}
				tmp.SetInternedString(types.ZSTR_OBJECT_OPERATOR)
				stack_frame.GetArr().KeyAddNew(types.ZSTR_TYPE.GetStr(), &tmp)
			} else if func_.GetScope() != nil {
				tmp.SetStringCopy(func_.GetScope().GetName())
				stack_frame.GetArr().KeyAddNew(types.ZSTR_CLASS.GetStr(), &tmp)
				tmp.SetInternedString(types.ZSTR_PAAMAYIM_NEKUDOTAYIM)
				stack_frame.GetArr().KeyAddNew(types.ZSTR_TYPE.GetStr(), &tmp)
			}
			if (options&DEBUG_BACKTRACE_IGNORE_ARGS) == 0 && func_.GetType() != ZEND_EVAL_CODE {
				DebugBacktraceGetArgs(call, &tmp)
				stack_frame.GetArr().KeyAddNew(types.ZSTR_ARGS.GetStr(), &tmp)
			}
		} else {

			/* i know this is kinda ugly, but i'm trying to avoid extra cycles in the main execution loop */

			var build_filename_arg types.ZendBool = 1
			var pseudo_function_name *types.ZendString
			if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetCommonType())) || ptr.GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL {

				/* can happen when calling eval from a custom sapi */

				pseudo_function_name = types.ZSTR_UNKNOWN
				build_filename_arg = 0
			} else {
				switch ptr.GetOpline().GetExtendedValue() {
				case ZEND_EVAL:
					pseudo_function_name = types.ZSTR_EVAL
					build_filename_arg = 0
				case ZEND_INCLUDE:
					pseudo_function_name = types.ZSTR_INCLUDE
				case ZEND_REQUIRE:
					pseudo_function_name = types.ZSTR_REQUIRE
				case ZEND_INCLUDE_ONCE:
					pseudo_function_name = types.ZSTR_INCLUDE_ONCE
				case ZEND_REQUIRE_ONCE:
					pseudo_function_name = types.ZSTR_REQUIRE_ONCE
				default:

					/* this can actually happen if you use debug_backtrace() in your error_handler and
					 * you're in the top-scope */

					pseudo_function_name = types.ZSTR_UNKNOWN
					build_filename_arg = 0
				}
			}
			if build_filename_arg != 0 && include_filename != nil {
				var arg_array types.Zval
				ArrayInit(&arg_array)

				/* include_filename always points to the last filename of the last last called-function.
				   if we have called include in the frame above - this is the file we have included.
				*/

				tmp.SetStringCopy(include_filename)
				arg_array.GetArr().NextIndexInsertNew(&tmp)
				stack_frame.GetArr().KeyAddNew(types.ZSTR_ARGS.GetStr(), &arg_array)
			}
			tmp.SetInternedString(pseudo_function_name)
			stack_frame.GetArr().KeyAddNew(types.ZSTR_FUNCTION.GetStr(), &tmp)
		}
		return_value.GetArr().NextIndexInsertNew(&stack_frame)
		include_filename = filename
		call = skip
		ptr = skip.GetPrevExecuteData()
	}
}
func ZifDebugBacktrace(executeData *ZendExecuteData, return_value *types.Zval) {
	var options ZendLong = DEBUG_BACKTRACE_PROVIDE_OBJECT
	var limit ZendLong = 0
	if ZendParseParameters(executeData.NumArgs(), "|ll", &options, &limit) == types.FAILURE {
		return
	}
	ZendFetchDebugBacktrace(return_value, 1, options, limit)
}
func ZifExtensionLoaded(executeData *ZendExecuteData, return_value *types.Zval) {
	var extension_name *types.ZendString
	var lcname *types.ZendString
	if ZendParseParameters(executeData.NumArgs(), "S", &extension_name) == types.FAILURE {
		return
	}
	lcname = ZendStringTolower(extension_name)
	if ZendHashExists(&ModuleRegistry, lcname) != 0 {
		return_value.SetTrue()
	} else {
		return_value.SetFalse()
	}
	types.ZendStringReleaseEx(lcname, 0)
}
func ZifGetExtensionFuncs(executeData *ZendExecuteData, return_value *types.Zval) {
	var extension_name *types.ZendString
	var lcname *types.ZendString
	var array int
	var module *ZendModuleEntry
	var zif *ZendFunction
	if ZendParseParameters(executeData.NumArgs(), "S", &extension_name) == types.FAILURE {
		return
	}
	if strncasecmp(extension_name.GetVal(), "zend", b.SizeOf("\"zend\"")) {
		lcname = ZendStringTolower(extension_name)
		module = ZendHashFindPtr(&ModuleRegistry, lcname)
		types.ZendStringReleaseEx(lcname, 0)
	} else {
		module = ZendHashStrFindPtr(&ModuleRegistry, "core", b.SizeOf("\"core\"")-1)
	}
	if module == nil {
		return_value.SetFalse()
		return
	}
	if module.GetFunctions() != nil {

		/* avoid BC break, if functions list is empty, will return an empty array */

		ArrayInit(return_value)
		array = 1
	} else {
		array = 0
	}
	var __ht *types.HashTable = CG__().GetFunctionTable()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

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
		return_value.SetFalse()
		return
	}
}
