package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func CE_STATIC_MEMBERS(ce *types.ClassEntry) *types.Zval {
	// todo
	return (ZEND_MAP_PTR_GET(ce.static_members_table__ptr)).(*types.Zval)
}
func ZEND_FCI_INITIALIZED(fci ZendFcallInfo) bool { return fci.GetSize() != 0 }
func ZendGetParametersArray(ht uint32, param_count int, argument_array *types.Zval) int {
	return _zendGetParametersArrayEx(param_count, argument_array)
}
func ZendGetParametersArrayEx(param_count int, argument_array *types.Zval) int {
	return _zendGetParametersArrayEx(param_count, argument_array)
}
func ZendParseParametersNone() types.ZEND_RESULT_CODE {
	if executeData.NumArgs() == 0 {
		return types.SUCCESS
	} else {
		CheckNumArgsNoneError()
		return types.FAILURE
	}
}
func ZendParseParametersNoneThrow() types.ZEND_RESULT_CODE {
	if executeData.NumArgs() == 0 {
		return types.SUCCESS
	} else {
		CheckNumArgsNoneException()
		return types.FAILURE
	}
}
func getThis(executeData *ZendExecuteData) *types.Zval {
	if ZEND_THIS(executeData).IsObject() {
		return ZEND_THIS(executeData)
	} else {
		return nil
	}
}
func ArrayInit(arg *types.Zval)                  { arg.SetArray(ZendNewArray(0)) }
func ArrayInitSize(arg *types.Zval, size uint32) { arg.SetArray(ZendNewArray(size)) }
func AddAssocLong(__arg *types.Zval, __key string, __n ZendLong) int {
	return AddAssocLongEx(__arg, __key, __n)
}
func AddAssocNull(__arg *types.Zval, __key string) int {
	return AddAssocNullEx(__arg, __key)
}
func AddAssocBool(__arg *types.Zval, __key string, __b int) int {
	return AddAssocBoolEx(__arg, __key, __b)
}
func AddAssocDouble(__arg *types.Zval, __key string, __d float64) int {
	return AddAssocDoubleEx(__arg, __key, __d)
}
func AddAssocStr(__arg *types.Zval, __key string, __str string) int {
	return AddAssocStrEx(__arg, __key, __str)
}
func AddAssocString(__arg *types.Zval, __key string, __str *byte) int {
	return AddAssocStrEx(__arg, __key, b.CastStrAuto(__str))
}
func AddAssocStringl(__arg *types.Zval, __key string, __str *byte, __length int) int {
	return AddAssocStringlEx(__arg, __key, b.CastStr(__str, __length))
}
func AddAssocZval(__arg *types.Zval, __key string, __value *types.Zval) int {
	return AddAssocZvalEx(__arg, __key, __value)
}
func AddIndexZval(arg *types.Zval, index ZendUlong, value *types.Zval) int {
	if arg.GetArr().IndexUpdateH(index, value) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexZval(arg *types.Zval, value *types.Zval) int {
	if arg.GetArr().NextIndexInsert(value) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddPropertyLong(__arg *types.Zval, __key string, __n ZendLong) int {
	return AddPropertyLongEx(__arg, __key, __n)
}
func AddPropertyNull(__arg *types.Zval, __key string) int {
	return AddPropertyNullEx(__arg, __key)
}
func AddPropertyResource(__arg *types.Zval, __key string, __r *types.ZendResource) int {
	return AddPropertyResourceEx(__arg, __key, __r)
}
func AddPropertyString(__arg *types.Zval, __key string, __str string) int {
	return AddPropertyStrEx(__arg, __key, __str)
}
func AddPropertyStringl(__arg *types.Zval, __key string, __str string) int {
	return AddPropertyStrEx(__arg, __key, __str)
}
func AddPropertyZval(__arg *types.Zval, __key string, __value *types.Zval) int {
	return AddPropertyZvalEx(__arg, __key, __value)
}
func CallUserFunction(object *types.Zval, function_name *types.Zval, retval_ptr *types.Zval, param_count uint32, params []types.Zval) int {
	return _callUserFunctionEx(object, function_name, retval_ptr, param_count, params, 1)
}
func CallUserFunctionEx(object *types.Zval, function_name *types.Zval, retval_ptr *types.Zval, param_count uint32, params []types.Zval, no_separation int) int {
	return _callUserFunctionEx(object, function_name, retval_ptr, param_count, params, no_separation)
}
func ZendForbidDynamicCall(func_name string) int {
	var ex *ZendExecuteData = CurrEX()
	b.Assert(ex != nil && ex.GetFunc() != nil)
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_DYNAMIC) != 0 {
		faults.Error(faults.E_WARNING, "Cannot call %s dynamically", func_name)
		return types.FAILURE
	}
	return types.SUCCESS
}
func CHECK_NULL_PATH(p []byte, l int) bool { return len(p) != l }
func ZVAL_EMPTY_STRING(z *types.Zval)      { z.SetInternedString(types.ZSTR_EMPTY_ALLOC()) }
func ZVAL_ZVAL(z *types.Zval, zv *types.Zval, copy int, dtor int) {
	var __z *types.Zval = z
	var __zv *types.Zval = zv
	if !(__zv.IsReference()) {
		if copy != 0 && dtor == 0 {
			types.ZVAL_COPY(__z, __zv)
		} else {
			types.ZVAL_COPY_VALUE(__z, __zv)
		}
	} else {
		types.ZVAL_COPY(__z, types.Z_REFVAL_P(__zv))
		if dtor != 0 || copy == 0 {
			ZvalPtrDtor(__zv)
		}
	}
}
func HASH_OF(p *types.Zval) *types.ZendArray {
	if p.IsArray() {
		return p.GetArr()
	} else {
		if p.IsObject() {
			return types.Z_OBJ_HT_P(p).GetGetProperties()(p)
		} else {
			return nil
		}
	}
}
func _ZEND_TRY_ASSIGN_NULL(zv *types.Zval, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefNull(ref)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		_zv.SetNull()
		break
	}
}
func ZEND_TRY_ASSIGN_REF_NULL(zv *types.Zval) {
	b.Assert(zv.IsReference())
	_ZEND_TRY_ASSIGN_NULL(zv, 1)
}
func _ZEND_TRY_ASSIGN_LONG(zv *types.Zval, lval ZendLong, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefLong(ref, lval)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		_zv.SetLong(lval)
		break
	}
}
func ZEND_TRY_ASSIGN_REF_LONG(zv *types.Zval, lval ZendLong) {
	b.Assert(zv.IsReference())
	_ZEND_TRY_ASSIGN_LONG(zv, lval, 1)
}
func _ZEND_TRY_ASSIGN_DOUBLE(zv *types.Zval, dval float64, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefDouble(ref, dval)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		_zv.SetDouble(dval)
		break
	}
}
func ZEND_TRY_ASSIGN_REF_DOUBLE(zv *types.Zval, dval float64) {
	b.Assert(zv.IsReference())
	_ZEND_TRY_ASSIGN_DOUBLE(zv, dval, 1)
}
func _ZEND_TRY_ASSIGN_EMPTY_STRING(zv *types.Zval, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefEmptyString(ref)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_EMPTY_STRING(_zv)
		break
	}
}
func ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zv *types.Zval) {
	b.Assert(zv.IsReference())
	_ZEND_TRY_ASSIGN_EMPTY_STRING(zv, 1)
}
func _ZEND_TRY_ASSIGN_STR(zv *types.Zval, str *types.ZendString, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefStr(ref, str)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		_zv.SetString(str)
		break
	}
}
func ZEND_TRY_ASSIGN_REF_STR(zv *types.Zval, str *types.ZendString) {
	b.Assert(zv.IsReference())
	_ZEND_TRY_ASSIGN_STR(zv, str, 1)
}
func _ZEND_TRY_ASSIGN_STRING(zv *types.Zval, string *byte, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefString(ref, string)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		_zv.SetRawString(b.CastStrAuto(string))
		break
	}
}
func ZEND_TRY_ASSIGN_REF_STRING(zv *types.Zval, string *byte) {
	b.Assert(zv.IsReference())
	_ZEND_TRY_ASSIGN_STRING(zv, string, 1)
}
func _ZEND_TRY_ASSIGN_STRINGL(zv *types.Zval, string *byte, len_ int, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefStringl(ref, string, len_)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		_zv.SetRawString(b.CastStr(string, len_))
		break
	}
}
func ZEND_TRY_ASSIGN_REF_STRINGL(zv *types.Zval, string *byte, len_ int) {
	b.Assert(zv.IsReference())
	_ZEND_TRY_ASSIGN_STRINGL(zv, string, len_, 1)
}
