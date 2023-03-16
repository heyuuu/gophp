package zend

import b "sik/builtin"

func CE_STATIC_MEMBERS(ce *ZendClassEntry) *Zval {
	// todo
	return (ZEND_MAP_PTR_GET(ce.static_members_table__ptr)).(*Zval)
}
func ZEND_FCI_INITIALIZED(fci ZendFcallInfo) bool { return fci.GetSize() != 0 }
func ZendGetParametersArray(ht uint32, param_count int, argument_array *Zval) int {
	return _zendGetParametersArrayEx(param_count, argument_array)
}
func ZendGetParametersArrayEx(param_count int, argument_array *Zval) int {
	return _zendGetParametersArrayEx(param_count, argument_array)
}
func ZendParseParametersNone() ZEND_RESULT_CODE {
	if ZEND_NUM_ARGS() == 0 {
		return SUCCESS
	} else {
		CheckNumArgsNoneError()
		return FAILURE
	}
}
func ZendParseParametersNoneThrow() ZEND_RESULT_CODE {
	if ZEND_NUM_ARGS() == 0 {
		return SUCCESS
	} else {
		CheckNumArgsNoneException()
		return FAILURE
	}
}
func getThis(executeData *ZendExecuteData) *Zval {
	if ZEND_THIS(executeData).IsObject() {
		return ZEND_THIS(executeData)
	} else {
		return nil
	}
}
func ZEND_NUM_ARGS() uint32                { return EX_NUM_ARGS() }
func ArrayInit(arg *Zval)                  { arg.SetArray(ZendNewArray(0)) }
func ArrayInitSize(arg *Zval, size uint32) { arg.SetArray(ZendNewArray(size)) }
func AddAssocLong(__arg *Zval, __key string, __n ZendLong) int {
	return AddAssocLongEx(__arg, __key, __n)
}
func AddAssocNull(__arg *Zval, __key string) int {
	return AddAssocNullEx(__arg, __key)
}
func AddAssocBool(__arg *Zval, __key string, __b int) int {
	return AddAssocBoolEx(__arg, __key, __b)
}
func AddAssocDouble(__arg *Zval, __key string, __d float64) int {
	return AddAssocDoubleEx(__arg, __key, __d)
}
func AddAssocStr(__arg *Zval, __key string, __str string) int {
	return AddAssocStrEx(__arg, __key, __str)
}
func AddAssocString(__arg *Zval, __key string, __str *byte) int {
	return AddAssocStrEx(__arg, __key, b.CastStrAuto(__str))
}
func AddAssocStringl(__arg *Zval, __key string, __str *byte, __length int) int {
	return AddAssocStringlEx(__arg, __key, b.CastStr(__str, __length))
}
func AddAssocZval(__arg *Zval, __key string, __value *Zval) int {
	return AddAssocZvalEx(__arg, __key, __value)
}
func AddIndexZval(arg *Zval, index ZendUlong, value *Zval) int {
	if arg.GetArr().IndexUpdateH(index, value) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexZval(arg *Zval, value *Zval) int {
	if arg.GetArr().NextIndexInsert(value) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddPropertyLong(__arg *Zval, __key string, __n ZendLong) int {
	return AddPropertyLongEx(__arg, __key, __n)
}
func AddPropertyNull(__arg *Zval, __key string) int {
	return AddPropertyNullEx(__arg, __key)
}
func AddPropertyResource(__arg *Zval, __key string, __r *ZendResource) int {
	return AddPropertyResourceEx(__arg, __key, __r)
}
func AddPropertyString(__arg *Zval, __key string, __str string) int {
	return AddPropertyStrEx(__arg, __key, __str)
}
func AddPropertyStringl(__arg *Zval, __key string, __str string) int {
	return AddPropertyStrEx(__arg, __key, __str)
}
func AddPropertyZval(__arg *Zval, __key string, __value *Zval) int {
	return AddPropertyZvalEx(__arg, __key, __value)
}
func CallUserFunction(object *Zval, function_name *Zval, retval_ptr *Zval, param_count uint32, params []Zval) int {
	return _callUserFunctionEx(object, function_name, retval_ptr, param_count, params, 1)
}
func CallUserFunctionEx(object *Zval, function_name *Zval, retval_ptr *Zval, param_count uint32, params []Zval, no_separation int) int {
	return _callUserFunctionEx(object, function_name, retval_ptr, param_count, params, no_separation)
}
func ZendForbidDynamicCall(func_name string) int {
	var ex *ZendExecuteData = CurrEX()
	ZEND_ASSERT(ex != nil && ex.GetFunc() != nil)
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_DYNAMIC) != 0 {
		ZendError(E_WARNING, "Cannot call %s dynamically", func_name)
		return FAILURE
	}
	return SUCCESS
}
func CHECK_NULL_PATH(p []byte, l int) bool { return len(p) != l }
func ZVAL_EMPTY_STRING(z *Zval)            { z.SetInternedString(ZSTR_EMPTY_ALLOC()) }
func ZVAL_ZVAL(z *Zval, zv *Zval, copy int, dtor int) {
	var __z *Zval = z
	var __zv *Zval = zv
	if !(__zv.IsReference()) {
		if copy != 0 && dtor == 0 {
			ZVAL_COPY(__z, __zv)
		} else {
			ZVAL_COPY_VALUE(__z, __zv)
		}
	} else {
		ZVAL_COPY(__z, Z_REFVAL_P(__zv))
		if dtor != 0 || copy == 0 {
			ZvalPtrDtor(__zv)
		}
	}
}
func HASH_OF(p *Zval) *ZendArray {
	if p.IsArray() {
		return p.GetArr()
	} else {
		if p.IsObject() {
			return Z_OBJ_HT_P(p).GetGetProperties()(p)
		} else {
			return nil
		}
	}
}
func _ZEND_TRY_ASSIGN_NULL(zv *Zval, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *ZendReference = _zv.GetRef()
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
func ZEND_TRY_ASSIGN_REF_NULL(zv *Zval) {
	ZEND_ASSERT(zv.IsReference())
	_ZEND_TRY_ASSIGN_NULL(zv, 1)
}
func _ZEND_TRY_ASSIGN_LONG(zv *Zval, lval ZendLong, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *ZendReference = _zv.GetRef()
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
func ZEND_TRY_ASSIGN_REF_LONG(zv *Zval, lval ZendLong) {
	ZEND_ASSERT(zv.IsReference())
	_ZEND_TRY_ASSIGN_LONG(zv, lval, 1)
}
func _ZEND_TRY_ASSIGN_DOUBLE(zv *Zval, dval float64, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *ZendReference = _zv.GetRef()
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
func ZEND_TRY_ASSIGN_REF_DOUBLE(zv *Zval, dval float64) {
	ZEND_ASSERT(zv.IsReference())
	_ZEND_TRY_ASSIGN_DOUBLE(zv, dval, 1)
}
func _ZEND_TRY_ASSIGN_EMPTY_STRING(zv *Zval, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *ZendReference = _zv.GetRef()
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
func ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zv *Zval) {
	ZEND_ASSERT(zv.IsReference())
	_ZEND_TRY_ASSIGN_EMPTY_STRING(zv, 1)
}
func _ZEND_TRY_ASSIGN_STR(zv *Zval, str *ZendString, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *ZendReference = _zv.GetRef()
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
func ZEND_TRY_ASSIGN_REF_STR(zv *Zval, str *ZendString) {
	ZEND_ASSERT(zv.IsReference())
	_ZEND_TRY_ASSIGN_STR(zv, str, 1)
}
func _ZEND_TRY_ASSIGN_STRING(zv *Zval, string *byte, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *ZendReference = _zv.GetRef()
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
func ZEND_TRY_ASSIGN_REF_STRING(zv *Zval, string *byte) {
	ZEND_ASSERT(zv.IsReference())
	_ZEND_TRY_ASSIGN_STRING(zv, string, 1)
}
func _ZEND_TRY_ASSIGN_STRINGL(zv *Zval, string *byte, len_ int, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *ZendReference = _zv.GetRef()
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
func ZEND_TRY_ASSIGN_REF_STRINGL(zv *Zval, string *byte, len_ int) {
	ZEND_ASSERT(zv.IsReference())
	_ZEND_TRY_ASSIGN_STRINGL(zv, string, len_, 1)
}
