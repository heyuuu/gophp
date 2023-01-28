// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func ZEND_NS_NAME(ns string, name string) string { return ns + "\\" + name }
func ZEND_MODULE_POST_ZEND_DEACTIVATE_D(module __auto__) {
	var zm_post_zend_deactivate_module func() int
}
func INIT_CLASS_ENTRY_INIT_METHODS(class_container ZendClassEntry, functions []ZendFunctionEntry) {
	class_container.SetConstructor(nil)
	class_container.SetDestructor(nil)
	class_container.SetClone(nil)
	class_container.SetSerialize(nil)
	class_container.SetUnserialize(nil)
	class_container.create_object = nil
	class_container.SetGetStaticMethod(nil)
	class_container.SetCall(nil)
	class_container.SetCallstatic(nil)
	class_container.SetTostring(nil)
	class_container.SetGet(nil)
	class_container.SetSet(nil)
	class_container.SetUnset(nil)
	class_container.SetIsset(nil)
	class_container.SetDebugInfo(nil)
	class_container.SetSerializeFunc(nil)
	class_container.SetUnserializeFunc(nil)
	class_container.parent = nil
	class_container.SetNumInterfaces(0)
	class_container.SetTraitNames(nil)
	class_container.SetNumTraits(0)
	class_container.SetTraitAliases(nil)
	class_container.SetTraitPrecedences(nil)
	class_container.interfaces = nil
	class_container.SetGetIterator(nil)
	class_container.SetIteratorFuncsPtr(nil)
	class_container.SetModule(nil)
	class_container.SetBuiltinFunctions(functions)
}
func INIT_NS_CLASS_ENTRY(class_container __auto__, ns string, class_name string, functions __auto__) {
	memset(&class_container, 0, b.SizeOf("zend_class_entry"))
	class_container.name = ZendStringInitInterned(ZEND_NS_NAME(ns, class_name), b.SizeOf("ZEND_NS_NAME ( ns , class_name )")-1, 1)
	class_container.info.internal.builtin_functions = functions
}
func CE_STATIC_MEMBERS(ce *ZendClassEntry) *Zval {
	return (*Zval)(ZEND_MAP_PTR_GET(ce.static_members_table))
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
		ZendWrongParametersNoneError()
		return FAILURE
	}
}
func ZendParseParametersNoneThrow() ZEND_RESULT_CODE {
	if ZEND_NUM_ARGS() == 0 {
		return SUCCESS
	} else {
		ZendWrongParametersNoneException()
		return FAILURE
	}
}
func ZendRegisterClassAlias(name *byte, ce *ZendClassEntry) int {
	return ZendRegisterClassAliasEx(name, b.SizeOf("name")-1, ce, 1)
}
func ZendRegisterNsClassAlias(ns string, name string, ce *ZendClassEntry) int {
	return ZendRegisterClassAliasEx(ZEND_NS_NAME(ns, name), b.SizeOf("ZEND_NS_NAME ( ns , name )")-1, ce, 1)
}
func getThis() *Zval {
	if ZEND_THIS.IsType(IS_OBJECT) {
		return ZEND_THIS
	} else {
		return nil
	}
}
func ZEND_IS_METHOD_CALL() bool { return EX(func_).common.scope != nil }
func WRONG_PARAM_COUNT_WITH_RETVAL(ret ZEND_RESULT_CODE) __auto__ {
	return ZEND_WRONG_PARAM_COUNT_WITH_RETVAL(ret)
}
func ARG_COUNT(dummy __auto__) uint32      { return EX_NUM_ARGS() }
func ZEND_NUM_ARGS() uint32                { return EX_NUM_ARGS() }
func ArrayInit(arg *Zval)                  { ZVAL_ARR(arg, ZendNewArray(0)) }
func ArrayInitSize(arg *Zval, size uint32) { ZVAL_ARR(arg, ZendNewArray(size)) }
func AddAssocLong(__arg *Zval, __key string, __n ZendLong) int {
	return AddAssocLongEx(__arg, __key, strlen(__key), __n)
}
func AddAssocNull(__arg *Zval, __key string) int {
	return AddAssocNullEx(__arg, __key, strlen(__key))
}
func AddAssocBool(__arg *Zval, __key string, __b int) int {
	return AddAssocBoolEx(__arg, __key, strlen(__key), __b)
}
func AddAssocResource(__arg *Zval, __key *byte, __r *ZendResource) int {
	return AddAssocResourceEx(__arg, __key, strlen(__key), __r)
}
func AddAssocDouble(__arg *Zval, __key *byte, __d float64) int {
	return AddAssocDoubleEx(__arg, __key, strlen(__key), __d)
}
func AddAssocStr(__arg *Zval, __key string, __str *ZendString) int {
	return AddAssocStrEx(__arg, __key, strlen(__key), __str)
}
func AddAssocString(__arg *Zval, __key *byte, __str *byte) int {
	return AddAssocStringEx(__arg, __key, strlen(__key), __str)
}
func AddAssocStringl(__arg *Zval, __key string, __str *byte, __length int) int {
	return AddAssocStringlEx(__arg, __key, strlen(__key), __str, __length)
}
func AddAssocZval(__arg *Zval, __key string, __value *Zval) int {
	return AddAssocZvalEx(__arg, __key, strlen(__key), __value)
}
func AddIndexZval(arg *Zval, index ZendUlong, value *Zval) int {
	if ZendHashIndexUpdate(arg.GetArr(), index, value) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexZval(arg *Zval, value *Zval) int {
	if ZendHashNextIndexInsert(arg.GetArr(), value) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddPropertyLong(__arg *Zval, __key string, __n ZendLong) int {
	return AddPropertyLongEx(__arg, __key, strlen(__key), __n)
}
func AddPropertyNull(__arg *Zval, __key string) int {
	return AddPropertyNullEx(__arg, __key, strlen(__key))
}
func AddPropertyBool(__arg *Zval, __key *byte, __b ZendLong) int {
	return AddPropertyBoolEx(__arg, __key, strlen(__key), __b)
}
func AddPropertyResource(__arg *Zval, __key string, __r *ZendResource) int {
	return AddPropertyResourceEx(__arg, __key, strlen(__key), __r)
}
func AddPropertyDouble(__arg *Zval, __key *byte, __d float64) int {
	return AddPropertyDoubleEx(__arg, __key, strlen(__key), __d)
}
func AddPropertyStr(__arg *Zval, __key *byte, __str *ZendString) int {
	return AddPropertyStrEx(__arg, __key, strlen(__key), __str)
}
func AddPropertyString(__arg *Zval, __key string, __str *byte) int {
	return AddPropertyStringEx(__arg, __key, strlen(__key), __str)
}
func AddPropertyStringl(__arg *Zval, __key string, __str *byte, __length int) int {
	return AddPropertyStringlEx(__arg, __key, strlen(__key), __str, __length)
}
func AddPropertyZval(__arg *Zval, __key string, __value *Zval) int {
	return AddPropertyZvalEx(__arg, __key, strlen(__key), __value)
}
func CallUserFunction(function_table *HashTable, object *Zval, function_name *Zval, retval_ptr *Zval, param_count uint32, params []Zval) int {
	return _callUserFunctionEx(object, function_name, retval_ptr, param_count, params, 1)
}
func CallUserFunctionEx(function_table __auto__, object *Zval, function_name *Zval, retval_ptr *Zval, param_count uint32, params []Zval, no_separation int, symbol_table __auto__) int {
	return _callUserFunctionEx(object, function_name, retval_ptr, param_count, params, no_separation)
}
func ZendForbidDynamicCall(func_name string) int {
	var ex *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData()
	ZEND_ASSERT(ex != nil && ex.GetFunc() != nil)
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_DYNAMIC) != 0 {
		ZendError(E_WARNING, "Cannot call %s dynamically", func_name)
		return FAILURE
	}
	return SUCCESS
}
func CHECK_ZVAL_NULL_PATH(p *Zval) bool {
	return Z_STRLEN_P(p) != strlen(Z_STRVAL_P(p))
}
func CHECK_NULL_PATH(p []byte, l int) bool { return strlen(p) != size_t(l) }
func ZVAL_STRINGL(z *Zval, s *byte, l int) { ZVAL_NEW_STR(z, ZendStringInit(s, l, 0)) }
func ZVAL_STRING(z *Zval, s *byte) {
	var _s *byte = s
	ZVAL_STRINGL(z, _s, strlen(_s))
}
func ZVAL_EMPTY_STRING(z *Zval)             { ZVAL_INTERNED_STR(z, ZSTR_EMPTY_ALLOC()) }
func ZVAL_PSTRINGL(z *Zval, s *byte, l int) { ZVAL_NEW_STR(z, ZendStringInit(s, l, 1)) }
func ZVAL_PSTRING(z *Zval, s *byte) {
	var _s *byte = s
	ZVAL_PSTRINGL(z, _s, strlen(_s))
}
func ZVAL_EMPTY_PSTRING(z *Zval) { ZVAL_PSTRINGL(z, "", 0) }
func ZVAL_ZVAL(z *Zval, zv *Zval, copy int, dtor int) {
	var __z *Zval = z
	var __zv *Zval = zv
	if !(Z_ISREF_P(__zv)) {
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
func RETVAL_BOOL(b bool)                       { ZVAL_BOOL(return_value, b) }
func RETVAL_NULL()                             { ZVAL_NULL(return_value) }
func RETVAL_LONG(l int)                        { ZVAL_LONG(return_value, l) }
func RETVAL_DOUBLE(d float64)                  { ZVAL_DOUBLE(return_value, d) }
func RETVAL_STR(s *ZendString)                 { ZVAL_STR(return_value, s) }
func RETVAL_INTERNED_STR(s *ZendString)        { ZVAL_INTERNED_STR(return_value, s) }
func RETVAL_NEW_STR(s *ZendString)             { ZVAL_NEW_STR(return_value, s) }
func RETVAL_STR_COPY(s *ZendString)            { ZVAL_STR_COPY(return_value, s) }
func RETVAL_STRING(s *byte)                    { ZVAL_STRING(return_value, s) }
func RETVAL_STRINGL(s string, l int)           { ZVAL_STRINGL(return_value, s, l) }
func RETVAL_EMPTY_STRING()                     { ZVAL_EMPTY_STRING(return_value) }
func RETVAL_RES(r *ZendResource)               { ZVAL_RES(return_value, r) }
func RETVAL_ARR(r *HashTable)                  { ZVAL_ARR(return_value, r) }
func RETVAL_EMPTY_ARRAY()                      { ZVAL_EMPTY_ARRAY(return_value) }
func RETVAL_OBJ(r *ZendObject)                 { ZVAL_OBJ(return_value, r) }
func RETVAL_ZVAL(zv *Zval, copy int, dtor int) { ZVAL_ZVAL(return_value, zv, copy, dtor) }
func HASH_OF(p *Zval) __auto__ {
	if p.IsType(IS_ARRAY) {
		return p.GetArr()
	} else {
		if p.IsType(IS_OBJECT) {
			return Z_OBJ_HT_P(p).GetGetProperties()(p)
		} else {
			return nil
		}
	}
}
func ZVAL_IS_NULL(z *Zval) bool { return z.IsType(IS_NULL) }
func ZEND_GINIT(module __auto__) func(any) {
	return (func(any))(zm_globals_ctor_module)
}
func ZEND_GSHUTDOWN(module __auto__) func(any) {
	return (func(any))(zm_globals_dtor_module)
}
func _ZEND_TRY_ASSIGN_NULL(zv *Zval, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefNull(ref)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_NULL(_zv)
		break
	}
}
func ZEND_TRY_ASSIGN_NULL(zv *Zval) { _ZEND_TRY_ASSIGN_NULL(zv, 0) }
func ZEND_TRY_ASSIGN_REF_NULL(zv *Zval) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_NULL(zv, 1)
}
func _ZEND_TRY_ASSIGN_FALSE(zv *Zval, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefBool(ref, 0)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_FALSE(_zv)
		break
	}
}
func ZEND_TRY_ASSIGN_FALSE(zv *Zval) { _ZEND_TRY_ASSIGN_FALSE(zv, 0) }
func ZEND_TRY_ASSIGN_REF_FALSE(zv *Zval) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_FALSE(zv, 1)
}
func _ZEND_TRY_ASSIGN_TRUE(zv *Zval, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefBool(ref, 1)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_TRUE(_zv)
		break
	}
}
func ZEND_TRY_ASSIGN_TRUE(zv *Zval) { _ZEND_TRY_ASSIGN_TRUE(zv, 0) }
func ZEND_TRY_ASSIGN_REF_TRUE(zv *Zval) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_TRUE(zv, 1)
}
func _ZEND_TRY_ASSIGN_BOOL(zv *Zval, bval __auto__, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefBool(ref, 1)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_BOOL(_zv, bval)
		break
	}
}
func ZEND_TRY_ASSIGN_BOOL(zv *Zval, bval __auto__) { _ZEND_TRY_ASSIGN_BOOL(zv, bval, 0) }
func ZEND_TRY_ASSIGN_REF_BOOL(zv *Zval, bval __auto__) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_BOOL(zv, bval, 1)
}
func _ZEND_TRY_ASSIGN_LONG(zv *Zval, lval ZendLong, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefLong(ref, lval)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_LONG(_zv, lval)
		break
	}
}
func ZEND_TRY_ASSIGN_LONG(zv *Zval, lval ZendLong) { _ZEND_TRY_ASSIGN_LONG(zv, lval, 0) }
func ZEND_TRY_ASSIGN_REF_LONG(zv *Zval, lval ZendLong) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_LONG(zv, lval, 1)
}
func _ZEND_TRY_ASSIGN_DOUBLE(zv *Zval, dval float64, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefDouble(ref, dval)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_DOUBLE(_zv, dval)
		break
	}
}
func ZEND_TRY_ASSIGN_DOUBLE(zv *Zval, dval float64) { _ZEND_TRY_ASSIGN_DOUBLE(zv, dval, 0) }
func ZEND_TRY_ASSIGN_REF_DOUBLE(zv *Zval, dval float64) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_DOUBLE(zv, dval, 1)
}
func _ZEND_TRY_ASSIGN_EMPTY_STRING(zv *Zval, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
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
func ZEND_TRY_ASSIGN_EMPTY_STRING(zv *Zval) { _ZEND_TRY_ASSIGN_EMPTY_STRING(zv, 0) }
func ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zv *Zval) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_EMPTY_STRING(zv, 1)
}
func _ZEND_TRY_ASSIGN_STR(zv *Zval, str *ZendString, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefStr(ref, str)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_STR(_zv, str)
		break
	}
}
func ZEND_TRY_ASSIGN_STR(zv *Zval, str *ZendString) { _ZEND_TRY_ASSIGN_STR(zv, str, 0) }
func ZEND_TRY_ASSIGN_REF_STR(zv *Zval, str *ZendString) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_STR(zv, str, 1)
}
func _ZEND_TRY_ASSIGN_NEW_STR(zv *Zval, str *ZendString, is_str int) {
	for {
		var _zv *Zval = zv
		if is_str != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefStr(ref, str)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_NEW_STR(_zv, str)
		break
	}
}
func ZEND_TRY_ASSIGN_NEW_STR(zv *Zval, str *ZendString) { _ZEND_TRY_ASSIGN_NEW_STR(zv, str, 0) }
func ZEND_TRY_ASSIGN_REF_NEW_STR(zv *Zval, str *ZendString) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_NEW_STR(zv, str, 1)
}
func _ZEND_TRY_ASSIGN_STRING(zv *Zval, string *byte, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefString(ref, string)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_STRING(_zv, string)
		break
	}
}
func ZEND_TRY_ASSIGN_STRING(zv *Zval, string *byte) { _ZEND_TRY_ASSIGN_STRING(zv, string, 0) }
func ZEND_TRY_ASSIGN_REF_STRING(zv *Zval, string *byte) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_STRING(zv, string, 1)
}
func _ZEND_TRY_ASSIGN_STRINGL(zv *Zval, string *byte, len_ int, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefStringl(ref, string, len_)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_STRINGL(_zv, string, len_)
		break
	}
}
func ZEND_TRY_ASSIGN_STRINGL(zv *Zval, string *byte, len_ int) {
	_ZEND_TRY_ASSIGN_STRINGL(zv, string, len_, 0)
}
func ZEND_TRY_ASSIGN_REF_STRINGL(zv *Zval, string *byte, len_ int) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_STRINGL(zv, string, len_, 1)
}
func _ZEND_TRY_ASSIGN_ARR(zv *Zval, arr *ZendArray, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefArr(ref, arr)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_ARR(_zv, arr)
		break
	}
}
func ZEND_TRY_ASSIGN_ARR(zv *Zval, arr *ZendArray) { _ZEND_TRY_ASSIGN_ARR(zv, arr, 0) }
func ZEND_TRY_ASSIGN_REF_ARR(zv *Zval, arr *ZendArray) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_ARR(zv, arr, 1)
}
func _ZEND_TRY_ASSIGN_RES(zv *Zval, res *ZendResource, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefRes(ref, res)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_RES(_zv, res)
		break
	}
}
func ZEND_TRY_ASSIGN_RES(zv *Zval, res *ZendResource) { _ZEND_TRY_ASSIGN_RES(zv, res, 0) }
func ZEND_TRY_ASSIGN_REF_RES(zv *Zval, res *ZendResource) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_RES(zv, res, 1)
}
func _ZEND_TRY_ASSIGN_TMP(zv *Zval, other_zv *Zval, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRef(ref, other_zv)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_COPY_VALUE(_zv, other_zv)
		break
	}
}
func ZEND_TRY_ASSIGN_TMP(zv *Zval, other_zv *Zval) { _ZEND_TRY_ASSIGN_TMP(zv, other_zv, 0) }
func ZEND_TRY_ASSIGN_REF_TMP(zv *Zval, other_zv *Zval) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_TMP(zv, other_zv, 1)
}
func _ZEND_TRY_ASSIGN_VALUE(zv *Zval, other_zv *Zval, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefZval(ref, other_zv)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_COPY_VALUE(_zv, other_zv)
		break
	}
}
func ZEND_TRY_ASSIGN_VALUE(zv *Zval, other_zv *Zval) { _ZEND_TRY_ASSIGN_VALUE(zv, other_zv, 0) }
func ZEND_TRY_ASSIGN_REF_VALUE(zv *Zval, other_zv *Zval) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_VALUE(zv, other_zv, 1)
}
func ZEND_TRY_ASSIGN_COPY(zv *Zval, other_zv *Zval) {
	Z_TRY_ADDREF_P(other_zv)
	ZEND_TRY_ASSIGN_VALUE(zv, other_zv)
}
func ZEND_TRY_ASSIGN_REF_COPY(zv *Zval, other_zv *Zval) {
	Z_TRY_ADDREF_P(other_zv)
	ZEND_TRY_ASSIGN_REF_VALUE(zv, other_zv)
}
func _ZEND_TRY_ASSIGN_VALUE_EX(zv *Zval, other_zv *Zval, strict ZendBool, is_ref int) {
	for {
		var _zv *Zval = zv
		if is_ref != 0 || Z_ISREF_P(_zv) {
			var ref *ZendReference = _zv.GetRef()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefZvalEx(ref, other_zv, strict)
				break
			}
			_zv = ref.GetVal()
		}
		ZvalPtrDtor(_zv)
		ZVAL_COPY_VALUE(_zv, other_zv)
		break
	}
}
func ZEND_TRY_ASSIGN_VALUE_EX(zv *Zval, other_zv *Zval, strict ZendBool) {
	_ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict, 0)
}
func ZEND_TRY_ASSIGN_REF_VALUE_EX(zv *Zval, other_zv *Zval, strict ZendBool) {
	ZEND_ASSERT(Z_ISREF_P(zv))
	_ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict, 1)
}
func ZEND_TRY_ASSIGN_COPY_EX(zv *Zval, other_zv *Zval, strict ZendBool) {
	Z_TRY_ADDREF_P(other_zv)
	ZEND_TRY_ASSIGN_VALUE_EX(zv, other_zv, strict)
}
func ZEND_TRY_ASSIGN_REF_COPY_EX(zv *Zval, other_zv *Zval, strict ZendBool) {
	Z_TRY_ADDREF_P(other_zv)
	ZEND_TRY_ASSIGN_REF_VALUE_EX(zv, other_zv, strict)
}
func ZendTryArrayInitSize(zv *Zval, size uint32) *Zval {
	var arr *ZendArray = ZendNewArray(size)
	if Z_ISREF_P(zv) {
		var ref *ZendReference = zv.GetRef()
		if ZEND_REF_HAS_TYPE_SOURCES(ref) {
			if ZendTryAssignTypedRefArr(ref, arr) != SUCCESS {
				return nil
			}
			return ref.GetVal()
		}
		zv = ref.GetVal()
	}
	ZvalPtrDtor(zv)
	ZVAL_ARR(zv, arr)
	return zv
}
func ZendTryArrayInit(zv *Zval) *Zval { return ZendTryArrayInitSize(zv, 0) }
func Z_PARAM_PROLOGUE(deref int, separate int) {
	_i++
	ZEND_ASSERT(_i <= _min_num_args || _optional == 1)
	ZEND_ASSERT(_i > _min_num_args || _optional == 0)
	if _optional {
		if _i > _num_args {
			break
		}
	}
	_real_arg++
	_arg = _real_arg
	if deref != 0 {
		if Z_ISREF_P(_arg) {
			_arg = Z_REFVAL_P(_arg)
		}
	}
	if separate != 0 {
		SEPARATE_ZVAL_NOREF(_arg)
	}
}
func Z_PARAM_STRICT_LONG_EX2(dest ZendLong, is_null ZendBool, check_null int, deref int, separate int) {
	Z_PARAM_PROLOGUE(deref, separate)
	if ZendParseArgLong(_arg, &dest, &is_null, check_null, 1) == 0 {
		_expected_type = Z_EXPECTED_LONG
		_error_code = ZPP_ERROR_WRONG_ARG
		break
	}
}
func Z_PARAM_STRICT_LONG_EX(dest ZendLong, is_null ZendBool, check_null int, separate int) {
	Z_PARAM_STRICT_LONG_EX2(dest, is_null, check_null, separate, separate)
}
func Z_PARAM_STRICT_LONG(dest ZendLong) {
	Z_PARAM_STRICT_LONG_EX(dest, _dummy, 0, 0)
}
func Z_PARAM_OBJECT_OF_CLASS_EX2(dest *Zval, _ce *ZendClassEntry, check_null int, deref int, separate int) {
	Z_PARAM_PROLOGUE(deref, separate)
	if ZendParseArgObject(_arg, &dest, _ce, check_null) == 0 {
		if _ce != nil {
			_error = _ce.GetName().GetVal()
			_error_code = ZPP_ERROR_WRONG_CLASS
			break
		} else {
			_expected_type = Z_EXPECTED_OBJECT
			_error_code = ZPP_ERROR_WRONG_ARG
			break
		}
	}
}
func Z_PARAM_OBJECT_OF_CLASS_EX(dest *Zval, _ce *ZendClassEntry, check_null int, separate int) {
	Z_PARAM_OBJECT_OF_CLASS_EX2(dest, _ce, check_null, separate, separate)
}
func Z_PARAM_OBJECT_OF_CLASS(dest *Zval, _ce *ZendClassEntry) {
	Z_PARAM_OBJECT_OF_CLASS_EX(dest, _ce, 0, 0)
}
func Z_PARAM_ZVAL_DEREF_EX(dest *Zval, check_null int, separate int) {
	Z_PARAM_PROLOGUE(1, separate)
	ZendParseArgZvalDeref(_arg, &dest, check_null)
}
func Z_PARAM_ZVAL_DEREF(dest *Zval) { Z_PARAM_ZVAL_DEREF_EX(dest, 0, 0) }
func ZendParseArgBool(arg *Zval, dest *ZendBool, is_null *ZendBool, check_null int) int {
	if check_null != 0 {
		*is_null = 0
	}
	if arg.IsType(IS_TRUE) {
		*dest = 1
	} else if arg.IsType(IS_FALSE) {
		*dest = 0
	} else if check_null != 0 && arg.IsType(IS_NULL) {
		*is_null = 1
		*dest = 0
	} else {
		return ZendParseArgBoolSlow(arg, dest)
	}
	return 1
}
func ZendParseArgLong(arg *Zval, dest *ZendLong, is_null *ZendBool, check_null int, cap int) int {
	if check_null != 0 {
		*is_null = 0
	}
	if arg.IsType(IS_LONG) {
		*dest = arg.GetLval()
	} else if check_null != 0 && arg.IsType(IS_NULL) {
		*is_null = 1
		*dest = 0
	} else if cap != 0 {
		return ZendParseArgLongCapSlow(arg, dest)
	} else {
		return ZendParseArgLongSlow(arg, dest)
	}
	return 1
}
func ZendParseArgDouble(arg *Zval, dest *float64, is_null *ZendBool, check_null int) int {
	if check_null != 0 {
		*is_null = 0
	}
	if arg.IsType(IS_DOUBLE) {
		*dest = arg.GetDval()
	} else if check_null != 0 && arg.IsType(IS_NULL) {
		*is_null = 1
		*dest = 0.0
	} else {
		return ZendParseArgDoubleSlow(arg, dest)
	}
	return 1
}
func ZendParseArgStr(arg *Zval, dest **ZendString, check_null int) int {
	if arg.IsType(IS_STRING) {
		*dest = arg.GetStr()
	} else if check_null != 0 && arg.IsType(IS_NULL) {
		*dest = nil
	} else {
		return ZendParseArgStrSlow(arg, dest)
	}
	return 1
}
func ZendParseArgString(arg *Zval, dest **byte, dest_len *int, check_null int) int {
	var str *ZendString
	if ZendParseArgStr(arg, &str, check_null) == 0 {
		return 0
	}
	if check_null != 0 && str == nil {
		*dest = nil
		*dest_len = 0
	} else {
		*dest = str.GetVal()
		*dest_len = str.GetLen()
	}
	return 1
}
func ZendParseArgPathStr(arg *Zval, dest **ZendString, check_null int) int {
	if ZendParseArgStr(arg, dest, check_null) == 0 || (*dest) != nil && CHECK_NULL_PATH(dest.GetVal(), dest.GetLen()) {
		return 0
	}
	return 1
}
func ZendParseArgPath(arg *Zval, dest **byte, dest_len *int, check_null int) int {
	var str *ZendString
	if ZendParseArgPathStr(arg, &str, check_null) == 0 {
		return 0
	}
	if check_null != 0 && str == nil {
		*dest = nil
		*dest_len = 0
	} else {
		*dest = str.GetVal()
		*dest_len = str.GetLen()
	}
	return 1
}
func ZendParseArgArray(arg *Zval, dest **Zval, check_null int, or_object int) int {
	if arg.IsType(IS_ARRAY) || or_object != 0 && arg.IsType(IS_OBJECT) {
		*dest = arg
	} else if check_null != 0 && arg.IsType(IS_NULL) {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgArrayHt(arg *Zval, dest **HashTable, check_null int, or_object int, separate int) int {
	if arg.IsType(IS_ARRAY) {
		*dest = arg.GetArr()
	} else if or_object != 0 && arg.IsType(IS_OBJECT) {
		if separate != 0 && Z_OBJ_P(arg).GetProperties() != nil && GC_REFCOUNT(Z_OBJ_P(arg).GetProperties()) > 1 {
			if (GC_FLAGS(Z_OBJ_P(arg).GetProperties()) & IS_ARRAY_IMMUTABLE) == 0 {
				GC_DELREF(Z_OBJ_P(arg).GetProperties())
			}
			Z_OBJ_P(arg).SetProperties(ZendArrayDup(Z_OBJ_P(arg).GetProperties()))
		}
		*dest = Z_OBJ_HT_P(arg).GetGetProperties()(arg)
	} else if check_null != 0 && arg.IsType(IS_NULL) {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgObject(arg *Zval, dest **Zval, ce *ZendClassEntry, check_null int) int {
	if arg.IsType(IS_OBJECT) && (ce == nil || InstanceofFunction(Z_OBJCE_P(arg), ce) != 0) {
		*dest = arg
	} else if check_null != 0 && arg.IsType(IS_NULL) {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgResource(arg *Zval, dest **Zval, check_null int) int {
	if arg.IsType(IS_RESOURCE) {
		*dest = arg
	} else if check_null != 0 && arg.IsType(IS_NULL) {
		*dest = nil
	} else {
		return 0
	}
	return 1
}
func ZendParseArgFunc(arg *Zval, dest_fci *ZendFcallInfo, dest_fcc *ZendFcallInfoCache, check_null int, error **byte) int {
	if check_null != 0 && arg.IsType(IS_NULL) {
		dest_fci.SetSize(0)
		dest_fcc.SetFunctionHandler(nil)
		*error = nil
	} else if ZendFcallInfoInit(arg, 0, dest_fci, dest_fcc, nil, error) != SUCCESS {
		return 0
	}
	return 1
}
func ZendParseArgZval(arg *Zval, dest **Zval, check_null int) {
	if check_null != 0 && (arg.IsType(IS_NULL) || Z_ISREF_P(arg) && Z_REFVAL_P(arg).IsType(IS_NULL)) {
		*dest = nil
	} else {
		*dest = arg
	}
}
func ZendParseArgZvalDeref(arg *Zval, dest **Zval, check_null int) {
	if check_null != 0 && arg.IsType(IS_NULL) {
		*dest = nil
	} else {
		*dest = arg
	}
}
func _zendGetParametersArrayEx(param_count int, argument_array *Zval) int {
	var param_ptr *Zval
	var arg_count int
	param_ptr = ZEND_CALL_ARG(ExecutorGlobals.GetCurrentExecuteData(), 1)
	arg_count = ZEND_CALL_NUM_ARGS(ExecutorGlobals.GetCurrentExecuteData())
	if param_count > arg_count {
		return FAILURE
	}
	for b.PostDec(&param_count) > 0 {
		ZVAL_COPY_VALUE(argument_array, param_ptr)
		argument_array++
		param_ptr++
	}
	return SUCCESS
}
func ZendCopyParametersArray(param_count int, argument_array *Zval) int {
	var param_ptr *Zval
	var arg_count int
	param_ptr = ZEND_CALL_ARG(ExecutorGlobals.GetCurrentExecuteData(), 1)
	arg_count = ZEND_CALL_NUM_ARGS(ExecutorGlobals.GetCurrentExecuteData())
	if param_count > arg_count {
		return FAILURE
	}
	for b.PostDec(&param_count) > 0 {
		Z_TRY_ADDREF_P(param_ptr)
		ZendHashNextIndexInsertNew(argument_array.GetArr(), param_ptr)
		param_ptr++
	}
	return SUCCESS
}
func ZendWrongParamCount() {
	var space *byte
	var class_name *byte = GetActiveClassName(&space)
	ZendInternalArgumentCountError(ZEND_ARG_USES_STRICT_TYPES(), "Wrong parameter count for %s%s%s()", class_name, space, GetActiveFunctionName())
}
func ZendGetTypeByConst(type_ int) *byte {
	switch type_ {
	case IS_FALSE:

	case IS_TRUE:

	case _IS_BOOL:
		return "bool"
	case IS_LONG:
		return "int"
	case IS_DOUBLE:
		return "float"
	case IS_STRING:
		return "string"
	case IS_OBJECT:
		return "object"
	case IS_RESOURCE:
		return "resource"
	case IS_NULL:
		return "null"
	case IS_CALLABLE:
		return "callable"
	case IS_ITERABLE:
		return "iterable"
	case IS_ARRAY:
		return "array"
	case IS_VOID:
		return "void"
	case _IS_NUMBER:
		return "number"
	default:
		return "unknown"
	}
}
func ZendZvalTypeName(arg *Zval) *byte {
	ZVAL_DEREF(arg)
	return ZendGetTypeByConst(arg.GetType())
}
func ZendZvalGetType(arg *Zval) *ZendString {
	switch arg.GetType() {
	case IS_NULL:
		return ZSTR_KNOWN(ZEND_STR_NULL)
	case IS_FALSE:

	case IS_TRUE:
		return ZSTR_KNOWN(ZEND_STR_BOOLEAN)
	case IS_LONG:
		return ZSTR_KNOWN(ZEND_STR_INTEGER)
	case IS_DOUBLE:
		return ZSTR_KNOWN(ZEND_STR_DOUBLE)
	case IS_STRING:
		return ZSTR_KNOWN(ZEND_STR_STRING)
	case IS_ARRAY:
		return ZSTR_KNOWN(ZEND_STR_ARRAY)
	case IS_OBJECT:
		return ZSTR_KNOWN(ZEND_STR_OBJECT)
	case IS_RESOURCE:
		if ZendRsrcListGetRsrcType(arg.GetRes()) != nil {
			return ZSTR_KNOWN(ZEND_STR_RESOURCE)
		} else {
			return ZSTR_KNOWN(ZEND_STR_CLOSED_RESOURCE)
		}
	default:
		return nil
	}
}
func ZendWrongParametersNoneError() int {
	var num_args int = ZEND_CALL_NUM_ARGS(ExecutorGlobals.GetCurrentExecuteData())
	var active_function *ZendFunction = ExecutorGlobals.GetCurrentExecuteData().GetFunc()
	var class_name *byte = b.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendInternalArgumentCountError(ZEND_ARG_USES_STRICT_TYPES(), "%s%s%s() expects %s %d parameter%s, %d given", class_name, b.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), "exactly", 0, "s", num_args)
	return FAILURE
}
func ZendWrongParametersNoneException() int {
	var num_args int = ZEND_CALL_NUM_ARGS(ExecutorGlobals.GetCurrentExecuteData())
	var active_function *ZendFunction = ExecutorGlobals.GetCurrentExecuteData().GetFunc()
	var class_name *byte = b.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendInternalArgumentCountError(1, "%s%s%s() expects %s %d parameter%s, %d given", class_name, b.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), "exactly", 0, "s", num_args)
	return FAILURE
}
func ZendWrongParametersCountError(min_num_args int, max_num_args int) {
	var num_args int = ZEND_CALL_NUM_ARGS(ExecutorGlobals.GetCurrentExecuteData())
	var active_function *ZendFunction = ExecutorGlobals.GetCurrentExecuteData().GetFunc()
	var class_name *byte = b.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendInternalArgumentCountError(ZEND_ARG_USES_STRICT_TYPES(), "%s%s%s() expects %s %d parameter%s, %d given", class_name, b.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), b.Cond(b.Cond(min_num_args == max_num_args, "exactly", num_args < min_num_args), "at least", "at most"), b.Cond(num_args < min_num_args, min_num_args, max_num_args), b.Cond(b.Cond(num_args < min_num_args, min_num_args, max_num_args) == 1, "", "s"), num_args)
}
func ZendWrongParametersCountException(min_num_args int, max_num_args int) {
	var num_args int = ZEND_CALL_NUM_ARGS(ExecutorGlobals.GetCurrentExecuteData())
	var active_function *ZendFunction = ExecutorGlobals.GetCurrentExecuteData().GetFunc()
	var class_name *byte = b.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendInternalArgumentCountError(1, "%s%s%s() expects %s %d parameter%s, %d given", class_name, b.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), b.Cond(b.Cond(min_num_args == max_num_args, "exactly", num_args < min_num_args), "at least", "at most"), b.Cond(num_args < min_num_args, min_num_args, max_num_args), b.Cond(b.Cond(num_args < min_num_args, min_num_args, max_num_args) == 1, "", "s"), num_args)
}
func ZendWrongParameterTypeError(num int, expected_type ZendExpectedType, arg *Zval) {
	var space *byte
	var class_name *byte
	var expected_error []*byte = []*byte{"int", "bool", "string", "array", "valid callback", "resource", "a valid path", "object", "float", nil}
	if ExecutorGlobals.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(ZEND_ARG_USES_STRICT_TYPES(), "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), num, expected_error[expected_type], ZendZvalTypeName(arg))
}
func ZendWrongParameterTypeException(num int, expected_type ZendExpectedType, arg *Zval) {
	var space *byte
	var class_name *byte
	var expected_error []*byte = []*byte{"int", "bool", "string", "array", "valid callback", "resource", "a valid path", "object", "float", nil}
	if ExecutorGlobals.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(1, "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), num, expected_error[expected_type], ZendZvalTypeName(arg))
}
func ZendWrongParameterClassError(num int, name *byte, arg *Zval) {
	var space *byte
	var class_name *byte
	if ExecutorGlobals.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(ZEND_ARG_USES_STRICT_TYPES(), "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), num, name, ZendZvalTypeName(arg))
}
func ZendWrongParameterClassException(num int, name *byte, arg *Zval) {
	var space *byte
	var class_name *byte
	if ExecutorGlobals.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(1, "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), num, name, ZendZvalTypeName(arg))
}
func ZendWrongCallbackError(num int, error *byte) {
	var space *byte
	var class_name *byte
	if ExecutorGlobals.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(ZEND_ARG_USES_STRICT_TYPES(), "%s%s%s() expects parameter %d to be a valid callback, %s", class_name, space, GetActiveFunctionName(), num, error)
	Efree(error)
}
func ZendWrongCallbackException(num int, error *byte) {
	var space *byte
	var class_name *byte
	if ExecutorGlobals.GetException() != nil {
		return
	}
	class_name = GetActiveClassName(&space)
	ZendInternalTypeError(1, "%s%s%s() expects parameter %d to be a valid callback, %s", class_name, space, GetActiveFunctionName(), num, error)
	Efree(error)
}
func ZendWrongCallbackDeprecated(num int, error *byte) {
	var space *byte
	var class_name *byte = GetActiveClassName(&space)
	ZendError(E_DEPRECATED, "%s%s%s() expects parameter %d to be a valid callback, %s", class_name, space, GetActiveFunctionName(), num, error)
	Efree(error)
}
func ZendParseArgClass(arg *Zval, pce **ZendClassEntry, num int, check_null int) int {
	var ce_base *ZendClassEntry = *pce
	if check_null != 0 && arg.IsType(IS_NULL) {
		*pce = nil
		return 1
	}
	if TryConvertToString(arg) == 0 {
		*pce = nil
		return 0
	}
	*pce = ZendLookupClass(arg.GetStr())
	if ce_base != nil {
		if (*pce) == nil || InstanceofFunction(*pce, ce_base) == 0 {
			var space *byte
			var class_name *byte = GetActiveClassName(&space)
			ZendInternalTypeError(ZEND_ARG_USES_STRICT_TYPES(), "%s%s%s() expects parameter %d to be a class name derived from %s, '%s' given", class_name, space, GetActiveFunctionName(), num, ce_base.GetName().GetVal(), Z_STRVAL_P(arg))
			*pce = nil
			return 0
		}
	}
	if (*pce) == nil {
		var space *byte
		var class_name *byte = GetActiveClassName(&space)
		ZendInternalTypeError(ZEND_ARG_USES_STRICT_TYPES(), "%s%s%s() expects parameter %d to be a valid class name, '%s' given", class_name, space, GetActiveFunctionName(), num, Z_STRVAL_P(arg))
		return 0
	}
	return 1
}
func ZendParseArgBoolWeak(arg *Zval, dest *ZendBool) int {
	if arg.GetType() <= IS_STRING {
		*dest = ZendIsTrue(arg)
	} else {
		return 0
	}
	return 1
}
func ZendParseArgBoolSlow(arg *Zval, dest *ZendBool) int {
	if ZEND_ARG_USES_STRICT_TYPES() {
		return 0
	}
	return ZendParseArgBoolWeak(arg, dest)
}
func ZendParseArgLongWeak(arg *Zval, dest *ZendLong) int {
	if arg.IsType(IS_DOUBLE) {
		if core.ZendIsnan(arg.GetDval()) {
			return 0
		}
		if !(ZEND_DOUBLE_FITS_LONG(arg.GetDval())) {
			return 0
		} else {
			*dest = ZendDvalToLval(arg.GetDval())
		}
	} else if arg.IsType(IS_STRING) {
		var d float64
		var type_ int
		if b.Assign(&type_, IsNumericStrFunction(arg.GetStr(), dest, &d)) != IS_LONG {
			if type_ != 0 {
				if core.ZendIsnan(d) {
					return 0
				}
				if !(ZEND_DOUBLE_FITS_LONG(d)) {
					return 0
				} else {
					*dest = ZendDvalToLval(d)
				}
			} else {
				return 0
			}
		}
		if ExecutorGlobals.GetException() != nil {
			return 0
		}
	} else if arg.GetType() < IS_TRUE {
		*dest = 0
	} else if arg.IsType(IS_TRUE) {
		*dest = 1
	} else {
		return 0
	}
	return 1
}
func ZendParseArgLongSlow(arg *Zval, dest *ZendLong) int {
	if ZEND_ARG_USES_STRICT_TYPES() {
		return 0
	}
	return ZendParseArgLongWeak(arg, dest)
}
func ZendParseArgLongCapWeak(arg *Zval, dest *ZendLong) int {
	if arg.IsType(IS_DOUBLE) {
		if core.ZendIsnan(arg.GetDval()) {
			return 0
		}
		*dest = ZendDvalToLvalCap(arg.GetDval())
	} else if arg.IsType(IS_STRING) {
		var d float64
		var type_ int
		if b.Assign(&type_, IsNumericStrFunction(arg.GetStr(), dest, &d)) != IS_LONG {
			if type_ != 0 {
				if core.ZendIsnan(d) {
					return 0
				}
				*dest = ZendDvalToLvalCap(d)
			} else {
				return 0
			}
		}
		if ExecutorGlobals.GetException() != nil {
			return 0
		}
	} else if arg.GetType() < IS_TRUE {
		*dest = 0
	} else if arg.IsType(IS_TRUE) {
		*dest = 1
	} else {
		return 0
	}
	return 1
}
func ZendParseArgLongCapSlow(arg *Zval, dest *ZendLong) int {
	if ZEND_ARG_USES_STRICT_TYPES() {
		return 0
	}
	return ZendParseArgLongCapWeak(arg, dest)
}
func ZendParseArgDoubleWeak(arg *Zval, dest *float64) int {
	if arg.IsType(IS_LONG) {
		*dest = float64(arg.GetLval())
	} else if arg.IsType(IS_STRING) {
		var l ZendLong
		var type_ int
		if b.Assign(&type_, IsNumericStrFunction(arg.GetStr(), &l, dest)) != IS_DOUBLE {
			if type_ != 0 {
				*dest = float64(l)
			} else {
				return 0
			}
		}
		if ExecutorGlobals.GetException() != nil {
			return 0
		}
	} else if arg.GetType() < IS_TRUE {
		*dest = 0.0
	} else if arg.IsType(IS_TRUE) {
		*dest = 1.0
	} else {
		return 0
	}
	return 1
}
func ZendParseArgDoubleSlow(arg *Zval, dest *float64) int {
	if arg.IsType(IS_LONG) {

		/* SSTH Exception: IS_LONG may be accepted instead as IS_DOUBLE */

		*dest = float64(arg.GetLval())

		/* SSTH Exception: IS_LONG may be accepted instead as IS_DOUBLE */

	} else if ZEND_ARG_USES_STRICT_TYPES() {
		return 0
	}
	return ZendParseArgDoubleWeak(arg, dest)
}
func ZendParseArgStrWeak(arg *Zval, dest **ZendString) int {
	if arg.GetType() < IS_STRING {
		ConvertToString(arg)
		*dest = arg.GetStr()
	} else if arg.IsType(IS_OBJECT) {
		if Z_OBJ_HT(*arg).GetCastObject() != nil {
			var obj Zval
			if Z_OBJ_HT(*arg).GetCastObject()(arg, &obj, IS_STRING) == SUCCESS {
				ZvalPtrDtor(arg)
				ZVAL_COPY_VALUE(arg, &obj)
				*dest = arg.GetStr()
				return 1
			}
		} else if Z_OBJ_HT(*arg).GetGet() != nil {
			var rv Zval
			var z *Zval = Z_OBJ_HT(*arg).GetGet()(arg, &rv)
			if z.GetType() != IS_OBJECT {
				ZvalPtrDtor(arg)
				if z.IsType(IS_STRING) {
					ZVAL_COPY_VALUE(arg, z)
				} else {
					ZVAL_STR(arg, ZvalGetStringFunc(z))
					ZvalPtrDtor(z)
				}
				*dest = arg.GetStr()
				return 1
			}
			ZvalPtrDtor(z)
		}
		return 0
	} else {
		return 0
	}
	return 1
}
func ZendParseArgStrSlow(arg *Zval, dest **ZendString) int {
	if ZEND_ARG_USES_STRICT_TYPES() {
		return 0
	}
	return ZendParseArgStrWeak(arg, dest)
}
func ZendParseArgImpl(arg_num int, arg *Zval, va *va_list, spec **byte, error **byte, severity *int) *byte {
	var spec_walk *byte = *spec
	var c byte = b.PostInc(&(*spec_walk))
	var check_null int = 0
	var separate int = 0
	var real_arg *Zval = arg

	/* scan through modifiers */

	ZVAL_DEREF(arg)
	for true {
		if (*spec_walk) == '/' {
			SEPARATE_ZVAL_NOREF(arg)
			real_arg = arg
			separate = 1
		} else if (*spec_walk) == '!' {
			check_null = 1
		} else {
			break
		}
		spec_walk++
	}
	switch c {
	case 'l':

	case 'L':
		var p *ZendLong = __va_arg(*va, (*ZendLong)(_))
		var is_null *ZendBool = nil
		if check_null != 0 {
			is_null = __va_arg(*va, (*ZendBool)(_))
		}
		if ZendParseArgLong(arg, p, is_null, check_null, c == 'L') == 0 {
			return "int"
		}
		break
	case 'd':
		var p *float64 = __va_arg(*va, (*float64)(_))
		var is_null *ZendBool = nil
		if check_null != 0 {
			is_null = __va_arg(*va, (*ZendBool)(_))
		}
		if ZendParseArgDouble(arg, p, is_null, check_null) == 0 {
			return "float"
		}
		break
	case 's':
		var p **byte = __va_arg(*va, (**byte)(_))
		var pl *int = __va_arg(*va, (*int)(_))
		if ZendParseArgString(arg, p, pl, check_null) == 0 {
			return "string"
		}
		break
	case 'p':
		var p **byte = __va_arg(*va, (**byte)(_))
		var pl *int = __va_arg(*va, (*int)(_))
		if ZendParseArgPath(arg, p, pl, check_null) == 0 {
			return "a valid path"
		}
		break
	case 'P':
		var str **ZendString = __va_arg(*va, (**ZendString)(_))
		if ZendParseArgPathStr(arg, str, check_null) == 0 {
			return "a valid path"
		}
		break
	case 'S':
		var str **ZendString = __va_arg(*va, (**ZendString)(_))
		if ZendParseArgStr(arg, str, check_null) == 0 {
			return "string"
		}
		break
	case 'b':
		var p *ZendBool = __va_arg(*va, (*ZendBool)(_))
		var is_null *ZendBool = nil
		if check_null != 0 {
			is_null = __va_arg(*va, (*ZendBool)(_))
		}
		if ZendParseArgBool(arg, p, is_null, check_null) == 0 {
			return "bool"
		}
		break
	case 'r':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		if ZendParseArgResource(arg, p, check_null) == 0 {
			return "resource"
		}
		break
	case 'A':

	case 'a':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		if ZendParseArgArray(arg, p, check_null, c == 'A') == 0 {
			return "array"
		}
		break
	case 'H':

	case 'h':
		var p **HashTable = __va_arg(*va, (**HashTable)(_))
		if ZendParseArgArrayHt(arg, p, check_null, c == 'H', separate) == 0 {
			return "array"
		}
		break
	case 'o':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		if ZendParseArgObject(arg, p, nil, check_null) == 0 {
			return "object"
		}
		break
	case 'O':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		var ce *ZendClassEntry = __va_arg(*va, (*ZendClassEntry)(_))
		if ZendParseArgObject(arg, p, ce, check_null) == 0 {
			if ce != nil {
				return ce.GetName().GetVal()
			} else {
				return "object"
			}
		}
		break
	case 'C':
		var lookup *ZendClassEntry
		var pce **ZendClassEntry = __va_arg(*va, (**ZendClassEntry)(_))
		var ce_base *ZendClassEntry = *pce
		if check_null != 0 && arg.IsType(IS_NULL) {
			*pce = nil
			break
		}
		if TryConvertToString(arg) == 0 {
			*pce = nil
			return "valid class name"
		}
		if b.Assign(&lookup, ZendLookupClass(arg.GetStr())) == nil {
			*pce = nil
		} else {
			*pce = lookup
		}
		if ce_base != nil {
			if (*pce) == nil || InstanceofFunction(*pce, ce_base) == 0 {
				ZendSpprintf(error, 0, "to be a class name derived from %s, '%s' given", ce_base.GetName().GetVal(), Z_STRVAL_P(arg))
				*pce = nil
				return ""
			}
		}
		if (*pce) == nil {
			ZendSpprintf(error, 0, "to be a valid class name, '%s' given", Z_STRVAL_P(arg))
			return ""
		}
		break
		break
	case 'f':
		var fci *ZendFcallInfo = __va_arg(*va, (*ZendFcallInfo)(_))
		var fcc *ZendFcallInfoCache = __va_arg(*va, (*ZendFcallInfoCache)(_))
		var is_callable_error *byte = nil
		if check_null != 0 && arg.IsType(IS_NULL) {
			fci.SetSize(0)
			fcc.SetFunctionHandler(0)
			break
		}
		if ZendFcallInfoInit(arg, 0, fci, fcc, nil, &is_callable_error) == SUCCESS {
			if is_callable_error != nil {
				*severity = E_DEPRECATED
				ZendSpprintf(error, 0, "to be a valid callback, %s", is_callable_error)
				Efree(is_callable_error)
				*spec = spec_walk
				return ""
			}
			break
		} else {
			if is_callable_error != nil {
				*severity = E_ERROR
				ZendSpprintf(error, 0, "to be a valid callback, %s", is_callable_error)
				Efree(is_callable_error)
				return ""
			} else {
				return "valid callback"
			}
		}
	case 'z':
		var p **Zval = __va_arg(*va, (**Zval)(_))
		ZendParseArgZvalDeref(real_arg, p, check_null)
		break
	case 'Z':

		/* 'Z' iz not supported anymore and should be replaced with 'z' */

		ZEND_ASSERT(c != 'Z')
	default:
		return "unknown"
	}
	*spec = spec_walk
	return nil
}
func ZendParseArg(arg_num int, arg *Zval, va *va_list, spec **byte, flags int) int {
	var expected_type *byte = nil
	var error *byte = nil
	var severity int = 0
	expected_type = ZendParseArgImpl(arg_num, arg, va, spec, &error, &severity)
	if expected_type != nil {
		if ExecutorGlobals.GetException() != nil {
			return FAILURE
		}
		if (flags&ZEND_PARSE_PARAMS_QUIET) == 0 && ((*expected_type) || error != nil) {
			var space *byte
			var class_name *byte = GetActiveClassName(&space)
			var throw_exception ZendBool = ZEND_ARG_USES_STRICT_TYPES() || (flags&ZEND_PARSE_PARAMS_THROW) != 0
			if error != nil {
				ZendInternalTypeError(throw_exception, "%s%s%s() expects parameter %d %s", class_name, space, GetActiveFunctionName(), arg_num, error)
				Efree(error)
			} else {
				ZendInternalTypeError(throw_exception, "%s%s%s() expects parameter %d to be %s, %s given", class_name, space, GetActiveFunctionName(), arg_num, expected_type, ZendZvalTypeName(arg))
			}
		}
		if severity != E_DEPRECATED {
			return FAILURE
		}
	}
	return SUCCESS
}
func ZendParseParameter(flags int, arg_num int, arg *Zval, spec *byte, _ ...any) int {
	var va va_list
	var ret int
	va_start(va, spec)
	ret = ZendParseArg(arg_num, arg, &va, &spec, flags)
	va_end(va)
	return ret
}
func ZendParseParametersDebugError(msg string) {
	var active_function *ZendFunction = ExecutorGlobals.GetCurrentExecuteData().GetFunc()
	var class_name *byte = b.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
	ZendErrorNoreturn(E_CORE_ERROR, "%s%s%s(): %s", class_name, b.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), msg)
}
func ZendParseVaArgs(num_args int, type_spec *byte, va *va_list, flags int) int {
	var spec_walk *byte
	var c int
	var i int
	var min_num_args int = -1
	var max_num_args int = 0
	var post_varargs int = 0
	var arg *Zval
	var arg_count int
	var have_varargs ZendBool = 0
	var varargs **Zval = nil
	var n_varargs *int = nil
	for spec_walk = type_spec; *spec_walk; spec_walk++ {
		c = *spec_walk
		switch c {
		case 'l':

		case 'd':

		case 's':

		case 'b':

		case 'r':

		case 'a':

		case 'o':

		case 'O':

		case 'z':

		case 'Z':

		case 'C':

		case 'h':

		case 'f':

		case 'A':

		case 'H':

		case 'p':

		case 'S':

		case 'P':

		case 'L':
			max_num_args++
			break
		case '|':
			min_num_args = max_num_args
			break
		case '/':

		case '!':

			/* Pass */

			break
		case '*':

		case '+':
			if have_varargs != 0 {
				ZendParseParametersDebugError("only one varargs specifier (* or +) is permitted")
				return FAILURE
			}
			have_varargs = 1

			/* we expect at least one parameter in varargs */

			if c == '+' {
				max_num_args++
			}

			/* mark the beginning of varargs */

			post_varargs = max_num_args
			break
		default:
			ZendParseParametersDebugError("bad type specifier while parsing parameters")
			return FAILURE
		}
	}
	if min_num_args < 0 {
		min_num_args = max_num_args
	}
	if have_varargs != 0 {

		/* calculate how many required args are at the end of the specifier list */

		post_varargs = max_num_args - post_varargs
		max_num_args = -1
	}
	if num_args < min_num_args || num_args > max_num_args && max_num_args >= 0 {
		if (flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
			var active_function *ZendFunction = ExecutorGlobals.GetCurrentExecuteData().GetFunc()
			var class_name *byte = b.CondF1(active_function.GetScope() != nil, func() []byte { return active_function.GetScope().GetName().GetVal() }, "")
			var throw_exception ZendBool = ZEND_ARG_USES_STRICT_TYPES() || (flags&ZEND_PARSE_PARAMS_THROW) != 0
			ZendInternalArgumentCountError(throw_exception, "%s%s%s() expects %s %d parameter%s, %d given", class_name, b.Cond(class_name[0], "::", ""), active_function.GetFunctionName().GetVal(), b.Cond(b.Cond(min_num_args == max_num_args, "exactly", num_args < min_num_args), "at least", "at most"), b.Cond(num_args < min_num_args, min_num_args, max_num_args), b.Cond(b.Cond(num_args < min_num_args, min_num_args, max_num_args) == 1, "", "s"), num_args)
		}
		return FAILURE
	}
	arg_count = ZEND_CALL_NUM_ARGS(ExecutorGlobals.GetCurrentExecuteData())
	if num_args > arg_count {
		ZendParseParametersDebugError("could not obtain parameters for parsing")
		return FAILURE
	}
	i = 0
	for b.PostDec(&num_args) > 0 {
		if (*type_spec) == '|' {
			type_spec++
		}
		if (*type_spec) == '*' || (*type_spec) == '+' {
			var num_varargs int = num_args + 1 - post_varargs

			/* eat up the passed in storage even if it won't be filled in with varargs */

			varargs = __va_arg(*va, (**Zval)(_))
			n_varargs = __va_arg(*va, (*int)(_))
			type_spec++
			if num_varargs > 0 {
				*n_varargs = num_varargs
				*varargs = ZEND_CALL_ARG(ExecutorGlobals.GetCurrentExecuteData(), i+1)

				/* adjust how many args we have left and restart loop */

				num_args += 1 - num_varargs
				i += num_varargs
				continue
			} else {
				*varargs = nil
				*n_varargs = 0
			}
		}
		arg = ZEND_CALL_ARG(ExecutorGlobals.GetCurrentExecuteData(), i+1)
		if ZendParseArg(i+1, arg, va, &type_spec, flags) == FAILURE {

			/* clean up varargs array if it was used */

			if varargs != nil && (*varargs) != nil {
				*varargs = nil
			}
			return FAILURE
		}
		i++
	}
	return SUCCESS
}
func ZendParseParametersEx(flags int, num_args int, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	va_start(va, type_spec)
	retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
	va_end(va)
	return retval
}
func ZendParseParameters(num_args int, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	var flags int = 0
	va_start(va, type_spec)
	retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
	va_end(va)
	return retval
}
func ZendParseParametersThrow(num_args int, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	var flags int = ZEND_PARSE_PARAMS_THROW
	va_start(va, type_spec)
	retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
	va_end(va)
	return retval
}
func ZendParseMethodParameters(num_args int, this_ptr *Zval, type_spec string, _ ...any) int {
	var va va_list
	var retval int
	var flags int = 0
	var p *byte = type_spec
	var object **Zval
	var ce *ZendClassEntry

	/* Just checking this_ptr is not enough, because fcall_common_helper does not set
	 * Z_OBJ(EG(This)) to NULL when calling an internal function with common.scope == NULL.
	 * In that case EG(This) would still be the $this from the calling code and we'd take the
	 * wrong branch here. */

	var is_method ZendBool = ExecutorGlobals.GetCurrentExecuteData().GetFunc().GetScope() != nil
	if is_method == 0 || this_ptr == nil || this_ptr.GetType() != IS_OBJECT {
		va_start(va, type_spec)
		retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
		va_end(va)
	} else {
		p++
		va_start(va, type_spec)
		object = __va_arg(va, (**Zval)(_))
		ce = __va_arg(va, (*ZendClassEntry)(_))
		*object = this_ptr
		if ce != nil && InstanceofFunction(Z_OBJCE_P(this_ptr), ce) == 0 {
			ZendErrorNoreturn(E_CORE_ERROR, "%s::%s() must be derived from %s::%s", Z_OBJCE_P(this_ptr).GetName().GetVal(), GetActiveFunctionName(), ce.GetName().GetVal(), GetActiveFunctionName())
		}
		retval = ZendParseVaArgs(num_args, p, &va, flags)
		va_end(va)
	}
	return retval
}
func ZendParseMethodParametersEx(flags int, num_args int, this_ptr *Zval, type_spec *byte, _ ...any) int {
	var va va_list
	var retval int
	var p *byte = type_spec
	var object **Zval
	var ce *ZendClassEntry
	if this_ptr == nil {
		va_start(va, type_spec)
		retval = ZendParseVaArgs(num_args, type_spec, &va, flags)
		va_end(va)
	} else {
		p++
		va_start(va, type_spec)
		object = __va_arg(va, (**Zval)(_))
		ce = __va_arg(va, (*ZendClassEntry)(_))
		*object = this_ptr
		if ce != nil && InstanceofFunction(Z_OBJCE_P(this_ptr), ce) == 0 {
			if (flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
				ZendErrorNoreturn(E_CORE_ERROR, "%s::%s() must be derived from %s::%s", ce.GetName().GetVal(), GetActiveFunctionName(), Z_OBJCE_P(this_ptr).GetName().GetVal(), GetActiveFunctionName())
			}
			va_end(va)
			return FAILURE
		}
		retval = ZendParseVaArgs(num_args, p, &va, flags)
		va_end(va)
	}
	return retval
}
func ZendMergeProperties(obj *Zval, properties *HashTable) {
	var obj_ht *ZendObjectHandlers = Z_OBJ_HT_P(obj)
	var old_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	var key *ZendString
	var value *Zval
	ExecutorGlobals.SetFakeScope(Z_OBJCE_P(obj))
	for {
		var __ht *HashTable = properties
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			key = _p.GetKey()
			value = _z
			if key != nil {
				var member Zval
				ZVAL_STR(&member, key)
				obj_ht.GetWriteProperty()(obj, &member, value, nil)
			}
		}
		break
	}
	ExecutorGlobals.SetFakeScope(old_scope)
}
func ZendUpdateClassConstants(class_type *ZendClassEntry) int {
	if !class_type.IsConstantsUpdated() {
		var ce *ZendClassEntry
		var c *ZendClassConstant
		var val *Zval
		var prop_info *ZendPropertyInfo
		if class_type.parent {
			if ZendUpdateClassConstants(class_type.parent) != SUCCESS {
				return FAILURE
			}
		}
		for {
			var __ht *HashTable = class_type.GetConstantsTable()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = _p.GetVal()

				if _z.IsType(IS_UNDEF) {
					continue
				}
				c = _z.GetPtr()
				val = c.GetValue()
				if val.IsType(IS_CONSTANT_AST) {
					if ZvalUpdateConstantEx(val, c.GetCe()) != SUCCESS {
						return FAILURE
					}
				}
			}
			break
		}
		if class_type.GetDefaultStaticMembersCount() != 0 && CE_STATIC_MEMBERS(class_type) == nil {
			if class_type.GetType() == ZEND_INTERNAL_CLASS || class_type.HasCeFlags(ZEND_ACC_IMMUTABLE|ZEND_ACC_PRELOADED) {
				ZendClassInitStatics(class_type)
			}
		}
		ce = class_type
		for ce != nil {
			for {
				var __ht *HashTable = ce.GetPropertiesInfo()
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = _p.GetVal()

					if _z.IsType(IS_UNDEF) {
						continue
					}
					prop_info = _z.GetPtr()
					if prop_info.GetCe() == ce {
						if prop_info.IsStatic() {
							val = CE_STATIC_MEMBERS(class_type) + prop_info.GetOffset()
						} else {
							val = (*Zval)((*byte)(class_type.GetDefaultPropertiesTable() + prop_info.GetOffset() - OBJ_PROP_TO_OFFSET(0)))
						}
						if val.IsType(IS_CONSTANT_AST) {
							if prop_info.GetType() != 0 {
								var tmp Zval
								ZVAL_COPY(&tmp, val)
								if ZvalUpdateConstantEx(&tmp, ce) != SUCCESS {
									ZvalPtrDtor(&tmp)
									return FAILURE
								}
								if ZendVerifyPropertyType(prop_info, &tmp, 1) == 0 {
									ZvalPtrDtor(&tmp)
									return FAILURE
								}
								ZvalPtrDtor(val)
								ZVAL_COPY_VALUE(val, &tmp)
							} else if ZvalUpdateConstantEx(val, ce) != SUCCESS {
								return FAILURE
							}
						}
					}
				}
				break
			}
			ce = ce.parent
		}
		class_type.SetIsConstantsUpdated(true)
	}
	return SUCCESS
}
func _objectPropertiesInit(object *ZendObject, class_type *ZendClassEntry) {
	if class_type.GetDefaultPropertiesCount() != 0 {
		var src *Zval = class_type.GetDefaultPropertiesTable()
		var dst *Zval = object.GetPropertiesTable()
		var end *Zval = src + class_type.GetDefaultPropertiesCount()
		if class_type.GetType() == ZEND_INTERNAL_CLASS {
			for {
				ZVAL_COPY_OR_DUP_PROP(dst, src)
				src++
				dst++
				if src == end {
					break
				}
			}
		} else {
			for {
				ZVAL_COPY_PROP(dst, src)
				src++
				dst++
				if src == end {
					break
				}
			}
		}
	}
}
func ObjectPropertiesInit(object *ZendObject, class_type *ZendClassEntry) {
	object.SetProperties(nil)
	_objectPropertiesInit(object, class_type)
}
func ObjectPropertiesInitEx(object *ZendObject, properties *HashTable) {
	object.SetProperties(properties)
	if object.GetCe().GetDefaultPropertiesCount() != 0 {
		var prop *Zval
		var key *ZendString
		var property_info *ZendPropertyInfo
		for {
			var __ht *HashTable = properties
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = _p.GetVal()

				if _z.IsType(IS_UNDEF) {
					continue
				}
				key = _p.GetKey()
				prop = _z
				property_info = ZendGetPropertyInfo(object.GetCe(), key, 1)
				if property_info != ZEND_WRONG_PROPERTY_INFO && property_info != nil && !property_info.IsStatic() {
					var slot *Zval = OBJ_PROP(object, property_info.GetOffset())
					if property_info.GetType() != 0 {
						var tmp Zval
						ZVAL_COPY_VALUE(&tmp, prop)
						if ZendVerifyPropertyType(property_info, &tmp, 0) == 0 {
							continue
						}
						ZVAL_COPY_VALUE(slot, &tmp)
					} else {
						ZVAL_COPY_VALUE(slot, prop)
					}
					ZVAL_INDIRECT(prop, slot)
				}
			}
			break
		}
	}
}
func ObjectPropertiesLoad(object *ZendObject, properties *HashTable) {
	var prop *Zval
	var tmp Zval
	var key *ZendString
	var h ZendLong
	var property_info *ZendPropertyInfo
	for {
		var __ht *HashTable = properties
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			h = _p.GetH()
			key = _p.GetKey()
			prop = _z
			if key != nil {
				if key.GetVal()[0] == '0' {
					var class_name *byte
					var prop_name *byte
					var prop_name_len int
					if ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, &prop_name_len) == SUCCESS {
						var pname *ZendString = ZendStringInit(prop_name, prop_name_len, 0)
						var prev_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
						if class_name != nil && class_name[0] != '*' {
							var cname *ZendString = ZendStringInit(class_name, strlen(class_name), 0)
							ExecutorGlobals.SetFakeScope(ZendLookupClass(cname))
							ZendStringReleaseEx(cname, 0)
						}
						property_info = ZendGetPropertyInfo(object.GetCe(), pname, 1)
						ZendStringReleaseEx(pname, 0)
						ExecutorGlobals.SetFakeScope(prev_scope)
					} else {
						property_info = ZEND_WRONG_PROPERTY_INFO
					}
				} else {
					property_info = ZendGetPropertyInfo(object.GetCe(), key, 1)
				}
				if property_info != ZEND_WRONG_PROPERTY_INFO && property_info != nil && !property_info.IsStatic() {
					var slot *Zval = OBJ_PROP(object, property_info.GetOffset())
					ZvalPtrDtor(slot)
					ZVAL_COPY_VALUE(slot, prop)
					ZvalAddRef(slot)
					if object.GetProperties() != nil {
						ZVAL_INDIRECT(&tmp, slot)
						ZendHashUpdate(object.GetProperties(), key, &tmp)
					}
				} else {
					if object.GetProperties() == nil {
						RebuildObjectProperties(object)
					}
					prop = ZendHashUpdate(object.GetProperties(), key, prop)
					ZvalAddRef(prop)
				}
			} else {
				if object.GetProperties() == nil {
					RebuildObjectProperties(object)
				}
				prop = ZendHashIndexUpdate(object.GetProperties(), h, prop)
				ZvalAddRef(prop)
			}
		}
		break
	}
}
func _objectAndPropertiesInit(arg *Zval, class_type *ZendClassEntry, properties *HashTable) int {
	if class_type.HasCeFlags(ZEND_ACC_INTERFACE | ZEND_ACC_TRAIT | ZEND_ACC_IMPLICIT_ABSTRACT_CLASS | ZEND_ACC_EXPLICIT_ABSTRACT_CLASS) {
		if class_type.IsInterface() {
			ZendThrowError(nil, "Cannot instantiate interface %s", class_type.GetName().GetVal())
		} else if class_type.IsTrait() {
			ZendThrowError(nil, "Cannot instantiate trait %s", class_type.GetName().GetVal())
		} else {
			ZendThrowError(nil, "Cannot instantiate abstract class %s", class_type.GetName().GetVal())
		}
		ZVAL_NULL(arg)
		arg.SetObj(nil)
		return FAILURE
	}
	if !class_type.IsConstantsUpdated() {
		if ZendUpdateClassConstants(class_type) != SUCCESS {
			ZVAL_NULL(arg)
			arg.SetObj(nil)
			return FAILURE
		}
	}
	if class_type.create_object == nil {
		var obj *ZendObject = ZendObjectsNew(class_type)
		ZVAL_OBJ(arg, obj)
		if properties != nil {
			ObjectPropertiesInitEx(obj, properties)
		} else {
			_objectPropertiesInit(obj, class_type)
		}
	} else {
		ZVAL_OBJ(arg, class_type.create_object(class_type))
	}
	return SUCCESS
}
func ObjectAndPropertiesInit(arg *Zval, class_type *ZendClassEntry, properties *HashTable) int {
	return _objectAndPropertiesInit(arg, class_type, properties)
}
func ObjectInitEx(arg *Zval, class_type *ZendClassEntry) int {
	return _objectAndPropertiesInit(arg, class_type, nil)
}
func ObjectInit(arg *Zval) int {
	ZVAL_OBJ(arg, ZendObjectsNew(ZendStandardClassDef))
	return SUCCESS
}
func AddAssocLongEx(arg *Zval, key string, key_len int, n ZendLong) int {
	var tmp Zval
	ZVAL_LONG(&tmp, n)
	ZendSymtableStrUpdate(arg.GetArr(), key, key_len, &tmp)
	return SUCCESS
}
func AddAssocNullEx(arg *Zval, key *byte, key_len int) int {
	var tmp Zval
	ZVAL_NULL(&tmp)
	ZendSymtableStrUpdate(arg.GetArr(), key, key_len, &tmp)
	return SUCCESS
}
func AddAssocBoolEx(arg *Zval, key string, key_len int, b int) int {
	var tmp Zval
	ZVAL_BOOL(&tmp, b)
	ZendSymtableStrUpdate(arg.GetArr(), key, key_len, &tmp)
	return SUCCESS
}
func AddAssocResourceEx(arg *Zval, key *byte, key_len int, r *ZendResource) int {
	var tmp Zval
	ZVAL_RES(&tmp, r)
	ZendSymtableStrUpdate(arg.GetArr(), key, key_len, &tmp)
	return SUCCESS
}
func AddAssocDoubleEx(arg *Zval, key string, key_len int, d float64) int {
	var tmp Zval
	ZVAL_DOUBLE(&tmp, d)
	ZendSymtableStrUpdate(arg.GetArr(), key, key_len, &tmp)
	return SUCCESS
}
func AddAssocStrEx(arg *Zval, key *byte, key_len int, str *ZendString) int {
	var tmp Zval
	ZVAL_STR(&tmp, str)
	ZendSymtableStrUpdate(arg.GetArr(), key, key_len, &tmp)
	return SUCCESS
}
func AddAssocStringEx(arg *Zval, key string, key_len int, str *byte) int {
	var tmp Zval
	ZVAL_STRING(&tmp, str)
	ZendSymtableStrUpdate(arg.GetArr(), key, key_len, &tmp)
	return SUCCESS
}
func AddAssocStringlEx(arg *Zval, key *byte, key_len int, str *byte, length int) int {
	var tmp Zval
	ZVAL_STRINGL(&tmp, str, length)
	ZendSymtableStrUpdate(arg.GetArr(), key, key_len, &tmp)
	return SUCCESS
}
func AddAssocZvalEx(arg *Zval, key string, key_len int, value *Zval) int {
	ZendSymtableStrUpdate(arg.GetArr(), key, key_len, value)
	return SUCCESS
}
func AddIndexLong(arg *Zval, index ZendUlong, n ZendLong) int {
	var tmp Zval
	ZVAL_LONG(&tmp, n)
	ZendHashIndexUpdate(arg.GetArr(), index, &tmp)
	return SUCCESS
}
func AddIndexNull(arg *Zval, index ZendUlong) int {
	var tmp Zval
	ZVAL_NULL(&tmp)
	ZendHashIndexUpdate(arg.GetArr(), index, &tmp)
	return SUCCESS
}
func AddIndexBool(arg *Zval, index ZendUlong, b int) int {
	var tmp Zval
	ZVAL_BOOL(&tmp, b)
	ZendHashIndexUpdate(arg.GetArr(), index, &tmp)
	return SUCCESS
}
func AddIndexResource(arg *Zval, index ZendUlong, r *ZendResource) int {
	var tmp Zval
	ZVAL_RES(&tmp, r)
	ZendHashIndexUpdate(arg.GetArr(), index, &tmp)
	return SUCCESS
}
func AddIndexDouble(arg *Zval, index ZendUlong, d float64) int {
	var tmp Zval
	ZVAL_DOUBLE(&tmp, d)
	ZendHashIndexUpdate(arg.GetArr(), index, &tmp)
	return SUCCESS
}
func AddIndexStr(arg *Zval, index ZendUlong, str *ZendString) int {
	var tmp Zval
	ZVAL_STR(&tmp, str)
	ZendHashIndexUpdate(arg.GetArr(), index, &tmp)
	return SUCCESS
}
func AddIndexString(arg *Zval, index ZendUlong, str *byte) int {
	var tmp Zval
	ZVAL_STRING(&tmp, str)
	ZendHashIndexUpdate(arg.GetArr(), index, &tmp)
	return SUCCESS
}
func AddIndexStringl(arg *Zval, index ZendUlong, str *byte, length int) int {
	var tmp Zval
	ZVAL_STRINGL(&tmp, str, length)
	ZendHashIndexUpdate(arg.GetArr(), index, &tmp)
	return SUCCESS
}
func AddNextIndexLong(arg *Zval, n ZendLong) int {
	var tmp Zval
	ZVAL_LONG(&tmp, n)
	if ZendHashNextIndexInsert(arg.GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexNull(arg *Zval) int {
	var tmp Zval
	ZVAL_NULL(&tmp)
	if ZendHashNextIndexInsert(arg.GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexBool(arg *Zval, b int) int {
	var tmp Zval
	ZVAL_BOOL(&tmp, b)
	if ZendHashNextIndexInsert(arg.GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexResource(arg *Zval, r *ZendResource) int {
	var tmp Zval
	ZVAL_RES(&tmp, r)
	if ZendHashNextIndexInsert(arg.GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexDouble(arg *Zval, d float64) int {
	var tmp Zval
	ZVAL_DOUBLE(&tmp, d)
	if ZendHashNextIndexInsert(arg.GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexStr(arg *Zval, str *ZendString) int {
	var tmp Zval
	ZVAL_STR(&tmp, str)
	if ZendHashNextIndexInsert(arg.GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexString(arg *Zval, str *byte) int {
	var tmp Zval
	ZVAL_STRING(&tmp, str)
	if ZendHashNextIndexInsert(arg.GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexStringl(arg *Zval, str *byte, length int) int {
	var tmp Zval
	ZVAL_STRINGL(&tmp, str, length)
	if ZendHashNextIndexInsert(arg.GetArr(), &tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func ArraySetZvalKey(ht *HashTable, key *Zval, value *Zval) int {
	var result *Zval
	switch key.GetType() {
	case IS_STRING:
		result = ZendSymtableUpdate(ht, key.GetStr(), value)
		break
	case IS_NULL:
		result = ZendSymtableUpdate(ht, ZSTR_EMPTY_ALLOC(), value)
		break
	case IS_RESOURCE:
		ZendError(E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", Z_RES_HANDLE_P(key), Z_RES_HANDLE_P(key))
		result = ZendHashIndexUpdate(ht, Z_RES_HANDLE_P(key), value)
		break
	case IS_FALSE:
		result = ZendHashIndexUpdate(ht, 0, value)
		break
	case IS_TRUE:
		result = ZendHashIndexUpdate(ht, 1, value)
		break
	case IS_LONG:
		result = ZendHashIndexUpdate(ht, key.GetLval(), value)
		break
	case IS_DOUBLE:
		result = ZendHashIndexUpdate(ht, ZendDvalToLval(key.GetDval()), value)
		break
	default:
		ZendError(E_WARNING, "Illegal offset type")
		result = nil
	}
	if result != nil {
		Z_TRY_ADDREF_P(result)
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddPropertyLongEx(arg *Zval, key *byte, key_len int, n ZendLong) int {
	var tmp Zval
	ZVAL_LONG(&tmp, n)
	return AddPropertyZvalEx(arg, key, key_len, &tmp)
}
func AddPropertyBoolEx(arg *Zval, key *byte, key_len int, b ZendLong) int {
	var tmp Zval
	ZVAL_BOOL(&tmp, b)
	return AddPropertyZvalEx(arg, key, key_len, &tmp)
}
func AddPropertyNullEx(arg *Zval, key *byte, key_len int) int {
	var tmp Zval
	ZVAL_NULL(&tmp)
	return AddPropertyZvalEx(arg, key, key_len, &tmp)
}
func AddPropertyResourceEx(arg *Zval, key *byte, key_len int, r *ZendResource) int {
	var tmp Zval
	ZVAL_RES(&tmp, r)
	AddPropertyZvalEx(arg, key, key_len, &tmp)
	ZvalPtrDtor(&tmp)
	return SUCCESS
}
func AddPropertyDoubleEx(arg *Zval, key *byte, key_len int, d float64) int {
	var tmp Zval
	ZVAL_DOUBLE(&tmp, d)
	return AddPropertyZvalEx(arg, key, key_len, &tmp)
}
func AddPropertyStrEx(arg *Zval, key *byte, key_len int, str *ZendString) int {
	var tmp Zval
	ZVAL_STR(&tmp, str)
	AddPropertyZvalEx(arg, key, key_len, &tmp)
	ZvalPtrDtor(&tmp)
	return SUCCESS
}
func AddPropertyStringEx(arg *Zval, key *byte, key_len int, str *byte) int {
	var tmp Zval
	ZVAL_STRING(&tmp, str)
	AddPropertyZvalEx(arg, key, key_len, &tmp)
	ZvalPtrDtor(&tmp)
	return SUCCESS
}
func AddPropertyStringlEx(arg *Zval, key *byte, key_len int, str *byte, length int) int {
	var tmp Zval
	ZVAL_STRINGL(&tmp, str, length)
	AddPropertyZvalEx(arg, key, key_len, &tmp)
	ZvalPtrDtor(&tmp)
	return SUCCESS
}
func AddPropertyZvalEx(arg *Zval, key *byte, key_len int, value *Zval) int {
	var z_key Zval
	ZVAL_STRINGL(&z_key, key, key_len)
	Z_OBJ_HT(*arg).GetWriteProperty()(arg, &z_key, value, nil)
	ZvalPtrDtor(&z_key)
	return SUCCESS
}
func ZendStartupModuleEx(module *ZendModuleEntry) int {
	var name_len int
	var lcname *ZendString
	if module.GetModuleStarted() != 0 {
		return SUCCESS
	}
	module.SetModuleStarted(1)

	/* Check module dependencies */

	if module.GetDeps() != nil {
		var dep *ZendModuleDep = module.GetDeps()
		for dep.GetName() != nil {
			if dep.GetType() == MODULE_DEP_REQUIRED {
				var req_mod *ZendModuleEntry
				name_len = strlen(dep.GetName())
				lcname = ZendStringAlloc(name_len, 0)
				ZendStrTolowerCopy(lcname.GetVal(), dep.GetName(), name_len)
				if b.Assign(&req_mod, ZendHashFindPtr(&ModuleRegistry, lcname)) == nil || req_mod.GetModuleStarted() == 0 {
					ZendStringEfree(lcname)

					/* TODO: Check version relationship */

					ZendError(E_CORE_WARNING, "Cannot load module '%s' because required module '%s' is not loaded", module.GetName(), dep.GetName())
					module.SetModuleStarted(0)
					return FAILURE
				}
				ZendStringEfree(lcname)
			}
			dep++
		}
	}

	/* Initialize module globals */

	if module.GetGlobalsSize() != 0 {
		if module.GetGlobalsCtor() != nil {
			module.GetGlobalsCtor()(module.GetGlobalsPtr())
		}
	}
	if module.GetModuleStartupFunc() != nil {
		ExecutorGlobals.SetCurrentModule(module)
		if module.GetModuleStartupFunc()(module.GetType(), module.GetModuleNumber()) == FAILURE {
			ZendErrorNoreturn(E_CORE_ERROR, "Unable to start %s module", module.GetName())
			ExecutorGlobals.SetCurrentModule(nil)
			return FAILURE
		}
		ExecutorGlobals.SetCurrentModule(nil)
	}
	return SUCCESS
}
func ZendStartupModuleZval(zv *Zval) int {
	var module *ZendModuleEntry = zv.GetPtr()
	if ZendStartupModuleEx(module) == SUCCESS {
		return ZEND_HASH_APPLY_KEEP
	} else {
		return ZEND_HASH_APPLY_REMOVE
	}
}
func ZendSortModules(base any, count int, siz int, compare CompareFuncT, swp SwapFuncT) {
	var b1 *Bucket = base
	var b2 *Bucket
	var end *Bucket = b1 + count
	var tmp Bucket
	var m *ZendModuleEntry
	var r *ZendModuleEntry
	for b1 < end {
	try_again:
		m = (*ZendModuleEntry)(b1.GetVal().GetPtr())
		if m.GetModuleStarted() == 0 && m.GetDeps() != nil {
			var dep *ZendModuleDep = m.GetDeps()
			for dep.GetName() != nil {
				if dep.GetType() == MODULE_DEP_REQUIRED || dep.GetType() == MODULE_DEP_OPTIONAL {
					b2 = b1 + 1
					for b2 < end {
						r = (*ZendModuleEntry)(b2.GetVal().GetPtr())
						if strcasecmp(dep.GetName(), r.GetName()) == 0 {
							tmp = *b1
							*b1 = *b2
							*b2 = tmp
							goto try_again
						}
						b2++
					}
				}
				dep++
			}
		}
		b1++
	}
}
func ZendCollectModuleHandlers() {
	var module *ZendModuleEntry
	var startup_count int = 0
	var shutdown_count int = 0
	var post_deactivate_count int = 0
	var ce *ZendClassEntry
	var class_count int = 0

	/* Collect extensions with request startup/shutdown handlers */

	for {
		var __ht *HashTable = &ModuleRegistry
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			module = _z.GetPtr()
			if module.GetRequestStartupFunc() != nil {
				startup_count++
			}
			if module.GetRequestShutdownFunc() != nil {
				shutdown_count++
			}
			if module.GetPostDeactivateFunc() != nil {
				post_deactivate_count++
			}
		}
		break
	}
	ModuleRequestStartupHandlers = (**ZendModuleEntry)(Malloc(b.SizeOf("zend_module_entry *") * (startup_count + 1 + shutdown_count + 1 + post_deactivate_count + 1)))
	ModuleRequestStartupHandlers[startup_count] = nil
	ModuleRequestShutdownHandlers = ModuleRequestStartupHandlers + startup_count + 1
	ModuleRequestShutdownHandlers[shutdown_count] = nil
	ModulePostDeactivateHandlers = ModuleRequestShutdownHandlers + shutdown_count + 1
	ModulePostDeactivateHandlers[post_deactivate_count] = nil
	startup_count = 0
	for {
		var __ht *HashTable = &ModuleRegistry
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			module = _z.GetPtr()
			if module.GetRequestStartupFunc() != nil {
				ModuleRequestStartupHandlers[b.PostInc(&startup_count)] = module
			}
			if module.GetRequestShutdownFunc() != nil {
				ModuleRequestShutdownHandlers[b.PreDec(&shutdown_count)] = module
			}
			if module.GetPostDeactivateFunc() != nil {
				ModulePostDeactivateHandlers[b.PreDec(&post_deactivate_count)] = module
			}
		}
		break
	}

	/* Collect internal classes with static members */

	for {
		var __ht *HashTable = CompilerGlobals.GetClassTable()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			ce = _z.GetPtr()
			if ce.GetType() == ZEND_INTERNAL_CLASS && ce.GetDefaultStaticMembersCount() > 0 {
				class_count++
			}
		}
		break
	}
	ClassCleanupHandlers = (**ZendClassEntry)(Malloc(b.SizeOf("zend_class_entry *") * (class_count + 1)))
	ClassCleanupHandlers[class_count] = nil
	if class_count != 0 {
		for {
			var __ht *HashTable = CompilerGlobals.GetClassTable()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = _p.GetVal()

				if _z.IsType(IS_UNDEF) {
					continue
				}
				ce = _z.GetPtr()
				if ce.GetType() == ZEND_INTERNAL_CLASS && ce.GetDefaultStaticMembersCount() > 0 {
					ClassCleanupHandlers[b.PreDec(&class_count)] = ce
				}
			}
			break
		}
	}
}
func ZendStartupModules() int {
	ZendHashSortEx(&ModuleRegistry, ZendSortModules, nil, 0)
	ZendHashApply(&ModuleRegistry, ZendStartupModuleZval)
	return SUCCESS
}
func ZendDestroyModules() {
	Free(ClassCleanupHandlers)
	Free(ModuleRequestStartupHandlers)
	ZendHashGracefulReverseDestroy(&ModuleRegistry)
}
func ZendRegisterModuleEx(module *ZendModuleEntry) *ZendModuleEntry {
	var name_len int
	var lcname *ZendString
	var module_ptr *ZendModuleEntry
	if module == nil {
		return nil
	}

	/* Check module dependencies */

	if module.GetDeps() != nil {
		var dep *ZendModuleDep = module.GetDeps()
		for dep.GetName() != nil {
			if dep.GetType() == MODULE_DEP_CONFLICTS {
				name_len = strlen(dep.GetName())
				lcname = ZendStringAlloc(name_len, 0)
				ZendStrTolowerCopy(lcname.GetVal(), dep.GetName(), name_len)
				if ZendHashExists(&ModuleRegistry, lcname) != 0 || ZendGetExtension(dep.GetName()) != nil {
					ZendStringEfree(lcname)

					/* TODO: Check version relationship */

					ZendError(E_CORE_WARNING, "Cannot load module '%s' because conflicting module '%s' is already loaded", module.GetName(), dep.GetName())
					return nil
				}
				ZendStringEfree(lcname)
			}
			dep++
		}
	}
	name_len = strlen(module.GetName())
	lcname = ZendStringAlloc(name_len, module.GetType() == MODULE_PERSISTENT)
	ZendStrTolowerCopy(lcname.GetVal(), module.GetName(), name_len)
	lcname = ZendNewInternedString(lcname)
	if b.Assign(&module_ptr, ZendHashAddMem(&ModuleRegistry, lcname, module, b.SizeOf("zend_module_entry"))) == nil {
		ZendError(E_CORE_WARNING, "Module '%s' already loaded", module.GetName())
		ZendStringRelease(lcname)
		return nil
	}
	module = module_ptr
	ExecutorGlobals.SetCurrentModule(module)
	if module.GetFunctions() != nil && ZendRegisterFunctions(nil, module.GetFunctions(), nil, module.GetType()) == FAILURE {
		ZendHashDel(&ModuleRegistry, lcname)
		ZendStringRelease(lcname)
		ExecutorGlobals.SetCurrentModule(nil)
		ZendError(E_CORE_WARNING, "%s: Unable to register functions, unable to load", module.GetName())
		return nil
	}
	ExecutorGlobals.SetCurrentModule(nil)
	ZendStringRelease(lcname)
	return module
}
func ZendRegisterInternalModule(module *ZendModuleEntry) *ZendModuleEntry {
	module.SetModuleNumber(ZendNextFreeModule())
	module.SetType(MODULE_PERSISTENT)
	return ZendRegisterModuleEx(module)
}
func ZendCheckMagicMethodImplementation(ce *ZendClassEntry, fptr *ZendFunction, error_type int) {
	var lcname []byte
	var name_len int
	if fptr.GetFunctionName().GetVal()[0] != '_' || fptr.GetFunctionName().GetVal()[1] != '_' {
		return
	}

	/* we don't care if the function name is longer, in fact lowercasing only
	 * the beginning of the name speeds up the check process */

	name_len = fptr.GetFunctionName().GetLen()
	ZendStrTolowerCopy(lcname, fptr.GetFunctionName().GetVal(), MIN(name_len, b.SizeOf("lcname")-1))
	lcname[b.SizeOf("lcname")-1] = '0'
	if name_len == b.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_DESTRUCTOR_FUNC_NAME, b.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Destructor %s::%s() cannot take arguments", ce.GetName().GetVal(), ZEND_DESTRUCTOR_FUNC_NAME)
	} else if name_len == b.SizeOf("ZEND_CLONE_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_CLONE_FUNC_NAME, b.SizeOf("ZEND_CLONE_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Method %s::%s() cannot accept any arguments", ce.GetName().GetVal(), ZEND_CLONE_FUNC_NAME)
	} else if name_len == b.SizeOf("ZEND_GET_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_GET_FUNC_NAME, b.SizeOf("ZEND_GET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 1 {
			ZendError(error_type, "Method %s::%s() must take exactly 1 argument", ce.GetName().GetVal(), ZEND_GET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), ZEND_GET_FUNC_NAME)
		}
	} else if name_len == b.SizeOf("ZEND_SET_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_SET_FUNC_NAME, b.SizeOf("ZEND_SET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 2 {
			ZendError(error_type, "Method %s::%s() must take exactly 2 arguments", ce.GetName().GetVal(), ZEND_SET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 || QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 2) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), ZEND_SET_FUNC_NAME)
		}
	} else if name_len == b.SizeOf("ZEND_UNSET_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_UNSET_FUNC_NAME, b.SizeOf("ZEND_UNSET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 1 {
			ZendError(error_type, "Method %s::%s() must take exactly 1 argument", ce.GetName().GetVal(), ZEND_UNSET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), ZEND_UNSET_FUNC_NAME)
		}
	} else if name_len == b.SizeOf("ZEND_ISSET_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_ISSET_FUNC_NAME, b.SizeOf("ZEND_ISSET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 1 {
			ZendError(error_type, "Method %s::%s() must take exactly 1 argument", ce.GetName().GetVal(), ZEND_ISSET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), ZEND_ISSET_FUNC_NAME)
		}
	} else if name_len == b.SizeOf("ZEND_CALL_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_CALL_FUNC_NAME, b.SizeOf("ZEND_CALL_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 2 {
			ZendError(error_type, "Method %s::%s() must take exactly 2 arguments", ce.GetName().GetVal(), ZEND_CALL_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 || QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 2) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), ZEND_CALL_FUNC_NAME)
		}
	} else if name_len == b.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_CALLSTATIC_FUNC_NAME, b.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 2 {
			ZendError(error_type, "Method %s::__callStatic() must take exactly 2 arguments", ce.GetName().GetVal())
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 || QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 2) != 0 {
			ZendError(error_type, "Method %s::__callStatic() cannot take arguments by reference", ce.GetName().GetVal())
		}
	} else if name_len == b.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_TOSTRING_FUNC_NAME, b.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Method %s::%s() cannot take arguments", ce.GetName().GetVal(), ZEND_TOSTRING_FUNC_NAME)
	} else if name_len == b.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_DEBUGINFO_FUNC_NAME, b.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Method %s::%s() cannot take arguments", ce.GetName().GetVal(), ZEND_DEBUGINFO_FUNC_NAME)
	}
}
func ZendRegisterFunctions(scope *ZendClassEntry, functions *ZendFunctionEntry, function_table *HashTable, type_ int) int {
	var ptr *ZendFunctionEntry = functions
	var function ZendFunction
	var reg_function *ZendFunction
	var internal_function *ZendInternalFunction = (*ZendInternalFunction)(&function)
	var count int = 0
	var unload int = 0
	var target_function_table *HashTable = function_table
	var error_type int
	var ctor *ZendFunction = nil
	var dtor *ZendFunction = nil
	var clone *ZendFunction = nil
	var __get *ZendFunction = nil
	var __set *ZendFunction = nil
	var __unset *ZendFunction = nil
	var __isset *ZendFunction = nil
	var __call *ZendFunction = nil
	var __callstatic *ZendFunction = nil
	var __tostring *ZendFunction = nil
	var __debugInfo *ZendFunction = nil
	var serialize_func *ZendFunction = nil
	var unserialize_func *ZendFunction = nil
	var lowercase_name *ZendString
	var fname_len int
	var lc_class_name *byte = nil
	var class_name_len int = 0
	if type_ == MODULE_PERSISTENT {
		error_type = E_CORE_WARNING
	} else {
		error_type = E_WARNING
	}
	if target_function_table == nil {
		target_function_table = CompilerGlobals.GetFunctionTable()
	}
	internal_function.SetType(ZEND_INTERNAL_FUNCTION)
	internal_function.SetModule(ExecutorGlobals.GetCurrentModule())
	memset(internal_function.GetReserved(), 0, ZEND_MAX_RESERVED_RESOURCES*b.SizeOf("void *"))
	if scope != nil {
		class_name_len = scope.GetName().GetLen()
		if b.Assign(&lc_class_name, ZendMemrchr(scope.GetName().GetVal(), '\\', class_name_len)) {
			lc_class_name++
			class_name_len -= lc_class_name - scope.GetName().GetVal()
			lc_class_name = ZendStrTolowerDup(lc_class_name, class_name_len)
		} else {
			lc_class_name = ZendStrTolowerDup(scope.GetName().GetVal(), class_name_len)
		}
	}
	for ptr.GetFname() != nil {
		fname_len = strlen(ptr.GetFname())
		internal_function.SetHandler(ptr.GetHandler())
		internal_function.SetFunctionName(ZendStringInitInterned(ptr.GetFname(), fname_len, 1))
		internal_function.SetScope(scope)
		internal_function.SetPrototype(nil)
		if ptr.GetFlags() != 0 {
			if !ptr.IsPppMask() {
				if ptr.GetFlags() != ZEND_ACC_DEPRECATED && scope != nil {
					ZendError(error_type, "Invalid access level for %s%s%s() - access must be exactly one of public, protected or private", b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), b.Cond(scope != nil, "::", ""), ptr.GetFname())
				}
				internal_function.SetFnFlags(ZEND_ACC_PUBLIC | ptr.GetFlags())
			} else {
				internal_function.SetFnFlags(ptr.GetFlags())
			}
		} else {
			internal_function.SetFnFlags(ZEND_ACC_PUBLIC)
		}
		if ptr.GetArgInfo() != nil {
			var info *ZendInternalFunctionInfo = (*ZendInternalFunctionInfo)(ptr.GetArgInfo())
			internal_function.SetArgInfo((*ZendInternalArgInfo)(ptr.GetArgInfo() + 1))
			internal_function.SetNumArgs(ptr.GetNumArgs())

			/* Currently you cannot denote that the function can accept less arguments than num_args */

			if info.GetRequiredNumArgs() == zend_uintptr_t-1 {
				internal_function.SetRequiredNumArgs(ptr.GetNumArgs())
			} else {
				internal_function.SetRequiredNumArgs(info.GetRequiredNumArgs())
			}
			if info.GetReturnReference() != 0 {
				internal_function.SetIsReturnReference(true)
			}
			if ptr.GetArgInfo()[ptr.GetNumArgs()].GetIsVariadic() != 0 {
				internal_function.SetIsVariadic(true)

				/* Don't count the variadic argument */

				internal_function.GetNumArgs()--

				/* Don't count the variadic argument */

			}
			if info.GetType().IsSet() {
				if info.GetType().IsClass() {
					var type_name *byte = (*byte)(info.GetType())
					if type_name[0] == '?' {
						type_name++
					}
					if scope == nil && (!(strcasecmp(type_name, "self")) || !(strcasecmp(type_name, "parent"))) {
						ZendErrorNoreturn(E_CORE_ERROR, "Cannot declare a return type of %s outside of a class scope", type_name)
					}
				}
				internal_function.SetIsHasReturnType(true)
			}
		} else {
			internal_function.SetArgInfo(nil)
			internal_function.SetNumArgs(0)
			internal_function.SetRequiredNumArgs(0)
		}
		ZendSetFunctionArgFlags((*ZendFunction)(internal_function))
		if ptr.IsAbstract() {
			if scope != nil {

				/* This is a class that must be abstract itself. Here we set the check info. */

				scope.SetIsImplicitAbstractClass(true)
				if !scope.IsInterface() {

					/* Since the class is not an interface it needs to be declared as a abstract class. */

					scope.SetIsExplicitAbstractClass(true)

					/* Since the class is not an interface it needs to be declared as a abstract class. */

				}
			}
			if ptr.IsStatic() && (scope == nil || !scope.IsInterface()) {
				ZendError(error_type, "Static function %s%s%s() cannot be abstract", b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), b.Cond(scope != nil, "::", ""), ptr.GetFname())
			}
		} else {
			if scope != nil && scope.IsInterface() {
				Efree((*byte)(lc_class_name))
				ZendError(error_type, "Interface %s cannot contain non abstract method %s()", scope.GetName().GetVal(), ptr.GetFname())
				return FAILURE
			}
			if internal_function.GetHandler() == nil {
				if scope != nil {
					Efree((*byte)(lc_class_name))
				}
				ZendError(error_type, "Method %s%s%s() cannot be a NULL function", b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), b.Cond(scope != nil, "::", ""), ptr.GetFname())
				ZendUnregisterFunctions(functions, count, target_function_table)
				return FAILURE
			}
		}
		lowercase_name = ZendStringTolowerEx(internal_function.GetFunctionName(), type_ == MODULE_PERSISTENT)
		lowercase_name = ZendNewInternedString(lowercase_name)
		reg_function = Malloc(b.SizeOf("zend_internal_function"))
		memcpy(reg_function, &function, b.SizeOf("zend_internal_function"))
		if ZendHashAddPtr(target_function_table, lowercase_name, reg_function) == nil {
			unload = 1
			Free(reg_function)
			ZendStringRelease(lowercase_name)
			break
		}

		/* If types of arguments have to be checked */

		if reg_function.GetArgInfo() != nil && reg_function.GetNumArgs() != 0 {
			var i uint32
			for i = 0; i < reg_function.GetNumArgs(); i++ {
				if reg_function.GetArgInfo()[i].GetType().IsSet() {
					reg_function.SetIsHasTypeHints(true)
					break
				}
			}
		}
		if reg_function.GetArgInfo() != nil && reg_function.HasFnFlags(ZEND_ACC_HAS_RETURN_TYPE|ZEND_ACC_HAS_TYPE_HINTS) {

			/* convert "const char*" class type names into "zend_string*" */

			var i uint32
			var num_args uint32 = reg_function.GetNumArgs() + 1
			var arg_info *ZendArgInfo = reg_function.GetArgInfo() - 1
			var new_arg_info *ZendArgInfo
			if reg_function.IsVariadic() {
				num_args++
			}
			new_arg_info = Malloc(b.SizeOf("zend_arg_info") * num_args)
			memcpy(new_arg_info, arg_info, b.SizeOf("zend_arg_info")*num_args)
			reg_function.SetArgInfo(new_arg_info + 1)
			for i = 0; i < num_args; i++ {
				if new_arg_info[i].GetType().IsClass() {
					var class_name *byte = (*byte)(new_arg_info[i].GetType())
					var allow_null ZendBool = 0
					var str *ZendString
					if class_name[0] == '?' {
						class_name++
						allow_null = 1
					}
					str = ZendStringInitInterned(class_name, strlen(class_name), 1)
					new_arg_info[i].SetType(ZEND_TYPE_ENCODE_CLASS(str, allow_null))
				}
			}
		}
		if scope != nil {

			/* Look for ctor, dtor, clone
			 * If it's an old-style constructor, store it only if we don't have
			 * a constructor already.
			 */

			if fname_len == class_name_len && ctor == nil && !(memcmp(lowercase_name.GetVal(), lc_class_name, class_name_len+1)) {
				ctor = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, "serialize") {
				serialize_func = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, "unserialize") {
				unserialize_func = reg_function
			} else if lowercase_name.GetVal()[0] != '_' || lowercase_name.GetVal()[1] != '_' {
				reg_function = nil
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_CONSTRUCTOR_FUNC_NAME) {
				ctor = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_DESTRUCTOR_FUNC_NAME) {
				dtor = reg_function
				if internal_function.GetNumArgs() != 0 {
					ZendError(error_type, "Destructor %s::%s() cannot take arguments", scope.GetName().GetVal(), ptr.GetFname())
				}
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_CLONE_FUNC_NAME) {
				clone = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_CALL_FUNC_NAME) {
				__call = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_CALLSTATIC_FUNC_NAME) {
				__callstatic = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_TOSTRING_FUNC_NAME) {
				__tostring = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_GET_FUNC_NAME) {
				__get = reg_function
				scope.SetIsUseGuards(true)
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_SET_FUNC_NAME) {
				__set = reg_function
				scope.SetIsUseGuards(true)
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_UNSET_FUNC_NAME) {
				__unset = reg_function
				scope.SetIsUseGuards(true)
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_ISSET_FUNC_NAME) {
				__isset = reg_function
				scope.SetIsUseGuards(true)
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_DEBUGINFO_FUNC_NAME) {
				__debugInfo = reg_function
			} else {
				reg_function = nil
			}
			if reg_function != nil {
				ZendCheckMagicMethodImplementation(scope, reg_function, error_type)
			}
		}
		ptr++
		count++
		ZendStringRelease(lowercase_name)
	}
	if unload != 0 {
		if scope != nil {
			Efree((*byte)(lc_class_name))
		}
		for ptr.GetFname() != nil {
			fname_len = strlen(ptr.GetFname())
			lowercase_name = ZendStringAlloc(fname_len, 0)
			ZendStrTolowerCopy(lowercase_name.GetVal(), ptr.GetFname(), fname_len)
			if ZendHashExists(target_function_table, lowercase_name) != 0 {
				ZendError(error_type, "Function registration failed - duplicate name - %s%s%s", b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), b.Cond(scope != nil, "::", ""), ptr.GetFname())
			}
			ZendStringEfree(lowercase_name)
			ptr++
		}
		ZendUnregisterFunctions(functions, count, target_function_table)
		return FAILURE
	}
	if scope != nil {
		scope.SetConstructor(ctor)
		scope.SetDestructor(dtor)
		scope.SetClone(clone)
		scope.SetCall(__call)
		scope.SetCallstatic(__callstatic)
		scope.SetTostring(__tostring)
		scope.SetGet(__get)
		scope.SetSet(__set)
		scope.SetUnset(__unset)
		scope.SetIsset(__isset)
		scope.SetDebugInfo(__debugInfo)
		scope.SetSerializeFunc(serialize_func)
		scope.SetUnserializeFunc(unserialize_func)
		if ctor != nil {
			ctor.SetIsCtor(true)
			if ctor.IsStatic() {
				ZendError(error_type, "Constructor %s::%s() cannot be static", scope.GetName().GetVal(), ctor.GetFunctionName().GetVal())
			}
			ctor.SetIsAllowStatic(false)
		}
		if dtor != nil {
			dtor.SetIsDtor(true)
			if dtor.IsStatic() {
				ZendError(error_type, "Destructor %s::%s() cannot be static", scope.GetName().GetVal(), dtor.GetFunctionName().GetVal())
			}
			dtor.SetIsAllowStatic(false)
		}
		if clone != nil {
			if clone.IsStatic() {
				ZendError(error_type, "%s::%s() cannot be static", scope.GetName().GetVal(), clone.GetFunctionName().GetVal())
			}
			clone.SetIsAllowStatic(false)
		}
		if __call != nil {
			if __call.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __call.GetFunctionName().GetVal())
			}
			__call.SetIsAllowStatic(false)
		}
		if __callstatic != nil {
			if !__callstatic.IsStatic() {
				ZendError(error_type, "Method %s::%s() must be static", scope.GetName().GetVal(), __callstatic.GetFunctionName().GetVal())
			}
			__callstatic.SetIsStatic(true)
		}
		if __tostring != nil {
			if __tostring.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __tostring.GetFunctionName().GetVal())
			}
			__tostring.SetIsAllowStatic(false)
		}
		if __get != nil {
			if __get.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __get.GetFunctionName().GetVal())
			}
			__get.SetIsAllowStatic(false)
		}
		if __set != nil {
			if __set.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __set.GetFunctionName().GetVal())
			}
			__set.SetIsAllowStatic(false)
		}
		if __unset != nil {
			if __unset.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __unset.GetFunctionName().GetVal())
			}
			__unset.SetIsAllowStatic(false)
		}
		if __isset != nil {
			if __isset.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __isset.GetFunctionName().GetVal())
			}
			__isset.SetIsAllowStatic(false)
		}
		if __debugInfo != nil {
			if __debugInfo.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __debugInfo.GetFunctionName().GetVal())
			}
		}
		if ctor != nil && ctor.IsHasReturnType() {
			ZendErrorNoreturn(E_CORE_ERROR, "Constructor %s::%s() cannot declare a return type", scope.GetName().GetVal(), ctor.GetFunctionName().GetVal())
		}
		if dtor != nil && dtor.IsHasReturnType() {
			ZendErrorNoreturn(E_CORE_ERROR, "Destructor %s::%s() cannot declare a return type", scope.GetName().GetVal(), dtor.GetFunctionName().GetVal())
		}
		if clone != nil && clone.IsHasReturnType() {
			ZendErrorNoreturn(E_CORE_ERROR, "%s::%s() cannot declare a return type", scope.GetName().GetVal(), clone.GetFunctionName().GetVal())
		}
		Efree((*byte)(lc_class_name))
	}
	return SUCCESS
}
func ZendUnregisterFunctions(functions *ZendFunctionEntry, count int, function_table *HashTable) {
	var ptr *ZendFunctionEntry = functions
	var i int = 0
	var target_function_table *HashTable = function_table
	var lowercase_name *ZendString
	var fname_len int
	if target_function_table == nil {
		target_function_table = CompilerGlobals.GetFunctionTable()
	}
	for ptr.GetFname() != nil {
		if count != -1 && i >= count {
			break
		}
		fname_len = strlen(ptr.GetFname())
		lowercase_name = ZendStringAlloc(fname_len, 0)
		ZendStrTolowerCopy(lowercase_name.GetVal(), ptr.GetFname(), fname_len)
		ZendHashDel(target_function_table, lowercase_name)
		ZendStringEfree(lowercase_name)
		ptr++
		i++
	}
}
func ZendStartupModule(module *ZendModuleEntry) int {
	if b.Assign(&module, ZendRegisterInternalModule(module)) != nil && ZendStartupModuleEx(module) == SUCCESS {
		return SUCCESS
	}
	return FAILURE
}
func ZendGetModuleStarted(module_name *byte) int {
	var module *ZendModuleEntry
	module = ZendHashStrFindPtr(&ModuleRegistry, module_name, strlen(module_name))
	if module != nil && module.GetModuleStarted() != 0 {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func CleanModuleClass(el *Zval, arg any) int {
	var ce *ZendClassEntry = (*ZendClassEntry)(el.GetPtr())
	var module_number int = *((*int)(arg))
	if ce.GetType() == ZEND_INTERNAL_CLASS && ce.GetModule().GetModuleNumber() == module_number {
		return ZEND_HASH_APPLY_REMOVE
	} else {
		return ZEND_HASH_APPLY_KEEP
	}
}
func CleanModuleClasses(module_number int) {
	ZendHashApplyWithArgument(ExecutorGlobals.GetClassTable(), CleanModuleClass, any(&module_number))
}
func ModuleDestructor(module *ZendModuleEntry) {
	if module.GetType() == MODULE_TEMPORARY {
		ZendCleanModuleRsrcDtors(module.GetModuleNumber())
		CleanModuleConstants(module.GetModuleNumber())
		CleanModuleClasses(module.GetModuleNumber())
	}
	if module.GetModuleStarted() != 0 && module.GetModuleShutdownFunc() != nil {
		module.GetModuleShutdownFunc()(module.GetType(), module.GetModuleNumber())
	}
	if module.GetModuleStarted() != 0 && module.GetModuleShutdownFunc() == nil && module.GetType() == MODULE_TEMPORARY {
		ZendUnregisterIniEntries(module.GetModuleNumber())
	}

	/* Deinitilaise module globals */

	if module.GetGlobalsSize() != 0 {
		if module.GetGlobalsDtor() != nil {
			module.GetGlobalsDtor()(module.GetGlobalsPtr())
		}
	}
	module.SetModuleStarted(0)
	if module.GetType() == MODULE_TEMPORARY && module.GetFunctions() != nil {
		ZendUnregisterFunctions(module.GetFunctions(), -1, nil)
	}
	if module.GetHandle() && !(getenv("ZEND_DONT_UNLOAD_MODULES")) {
		DL_UNLOAD(module.GetHandle())
	}
}
func ZendActivateModules() {
	var p **ZendModuleEntry = ModuleRequestStartupHandlers
	for (*p) != nil {
		var module *ZendModuleEntry = *p
		if module.GetRequestStartupFunc()(module.GetType(), module.GetModuleNumber()) == FAILURE {
			ZendError(E_WARNING, "request_startup() for %s module failed", module.GetName())
			exit(1)
		}
		p++
	}
}
func ZendDeactivateModules() {
	ExecutorGlobals.SetCurrentExecuteData(nil)
	var __orig_bailout *JMP_BUF = ExecutorGlobals.GetBailout()
	var __bailout JMP_BUF
	ExecutorGlobals.SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		if ExecutorGlobals.GetFullTablesCleanup() != 0 {
			var module *ZendModuleEntry
			for {
				var __ht *HashTable = &ModuleRegistry
				var _idx uint32 = __ht.GetNNumUsed()
				var _p *Bucket = __ht.GetArData() + _idx
				var _z *Zval
				for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
					_p--
					_z = _p.GetVal()

					if _z.IsType(IS_UNDEF) {
						continue
					}
					module = _z.GetPtr()
					if module.GetRequestShutdownFunc() != nil {
						module.GetRequestShutdownFunc()(module.GetType(), module.GetModuleNumber())
					}
				}
				break
			}
		} else {
			var p **ZendModuleEntry = ModuleRequestShutdownHandlers
			for (*p) != nil {
				var module *ZendModuleEntry = *p
				module.GetRequestShutdownFunc()(module.GetType(), module.GetModuleNumber())
				p++
			}
		}
	}
	ExecutorGlobals.SetBailout(__orig_bailout)
}
func ZendCleanupInternalClasses() {
	var p **ZendClassEntry = ClassCleanupHandlers
	for (*p) != nil {
		ZendCleanupInternalClassData(*p)
		p++
	}
}
func ZendPostDeactivateModules() {
	if ExecutorGlobals.GetFullTablesCleanup() != 0 {
		var module *ZendModuleEntry
		var zv *Zval
		var key *ZendString
		for {
			var __ht *HashTable = &ModuleRegistry
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = _p.GetVal()

				if _z.IsType(IS_UNDEF) {
					continue
				}
				module = _z.GetPtr()
				if module.GetPostDeactivateFunc() != nil {
					module.GetPostDeactivateFunc()()
				}
			}
			break
		}
		for {
			var __ht *HashTable = &ModuleRegistry
			var _idx uint32 = __ht.GetNNumUsed()
			var _p *Bucket = __ht.GetArData() + _idx
			var _z *Zval
			for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
				_p--
				_z = _p.GetVal()

				if _z.IsType(IS_UNDEF) {
					continue
				}
				key = _p.GetKey()
				zv = _z
				module = zv.GetPtr()
				if module.GetType() != MODULE_TEMPORARY {
					break
				}
				ModuleDestructor(module)
				Free(module)
				ZendStringReleaseEx(key, 0)
				__ht.GetNNumOfElements()--
				var j uint32 = HT_IDX_TO_HASH(_idx - 1)
				var nIndex uint32 = _p.GetH() | __ht.GetNTableMask()
				var i uint32 = HT_HASH(__ht, nIndex)
				if j != i {
					var prev *Bucket = HT_HASH_TO_BUCKET(__ht, i)
					for prev.GetVal().GetNext() != j {
						i = prev.GetVal().GetNext()
						prev = HT_HASH_TO_BUCKET(__ht, i)
					}
					prev.GetVal().GetNext() = _p.GetVal().GetNext()
				} else {
					HT_HASH(__ht, nIndex) = _p.GetVal().GetNext()
				}
			}
			__ht.SetNNumUsed(_idx)
			break
		}
	} else {
		var p **ZendModuleEntry = ModulePostDeactivateHandlers
		for (*p) != nil {
			var module *ZendModuleEntry = *p
			module.GetPostDeactivateFunc()()
			p++
		}
	}
}
func ZendNextFreeModule() int {
	return ModuleRegistry.GetNNumOfElements() + 1
}
func DoRegisterInternalClass(orig_class_entry *ZendClassEntry, ce_flags uint32) *ZendClassEntry {
	var class_entry *ZendClassEntry = Malloc(b.SizeOf("zend_class_entry"))
	var lowercase_name *ZendString
	*class_entry = *orig_class_entry
	class_entry.SetType(ZEND_INTERNAL_CLASS)
	ZendInitializeClassData(class_entry, 0)
	class_entry.SetCeFlags(ce_flags | ZEND_ACC_CONSTANTS_UPDATED | ZEND_ACC_LINKED | ZEND_ACC_RESOLVED_PARENT | ZEND_ACC_RESOLVED_INTERFACES)
	class_entry.SetModule(ExecutorGlobals.GetCurrentModule())
	if class_entry.GetBuiltinFunctions() != nil {
		ZendRegisterFunctions(class_entry, class_entry.GetBuiltinFunctions(), class_entry.GetFunctionTable(), ExecutorGlobals.GetCurrentModule().GetType())
	}
	lowercase_name = ZendStringTolowerEx(orig_class_entry.GetName(), ExecutorGlobals.GetCurrentModule().GetType() == MODULE_PERSISTENT)
	lowercase_name = ZendNewInternedString(lowercase_name)
	ZendHashUpdatePtr(CompilerGlobals.GetClassTable(), lowercase_name, class_entry)
	ZendStringReleaseEx(lowercase_name, 1)
	return class_entry
}
func ZendRegisterInternalClassEx(class_entry *ZendClassEntry, parent_ce *ZendClassEntry) *ZendClassEntry {
	var register_class *ZendClassEntry
	register_class = ZendRegisterInternalClass(class_entry)
	if parent_ce != nil {
		ZendDoInheritance(register_class, parent_ce)
		ZendBuildPropertiesInfoTable(register_class)
	}
	return register_class
}
func ZendClassImplements(class_entry *ZendClassEntry, num_interfaces int, _ ...any) {
	var interface_entry *ZendClassEntry
	var interface_list va_list
	va_start(interface_list, num_interfaces)
	for b.PostDec(&num_interfaces) {
		interface_entry = __va_arg(interface_list, (*ZendClassEntry)(_))
		ZendDoImplementInterface(class_entry, interface_entry)
	}
	va_end(interface_list)
}
func ZendRegisterInternalClass(orig_class_entry *ZendClassEntry) *ZendClassEntry {
	return DoRegisterInternalClass(orig_class_entry, 0)
}
func ZendRegisterInternalInterface(orig_class_entry *ZendClassEntry) *ZendClassEntry {
	return DoRegisterInternalClass(orig_class_entry, ZEND_ACC_INTERFACE)
}
func ZendRegisterClassAliasEx(name *byte, name_len int, ce *ZendClassEntry, persistent int) int {
	var lcname *ZendString
	var zv Zval
	var ret *Zval

	/* TODO: Move this out of here in 7.4. */

	if persistent != 0 && ExecutorGlobals.GetCurrentModule() != nil && ExecutorGlobals.GetCurrentModule().GetType() == MODULE_TEMPORARY {
		persistent = 0
	}
	if name[0] == '\\' {
		lcname = ZendStringAlloc(name_len-1, persistent)
		ZendStrTolowerCopy(lcname.GetVal(), name+1, name_len-1)
	} else {
		lcname = ZendStringAlloc(name_len, persistent)
		ZendStrTolowerCopy(lcname.GetVal(), name, name_len)
	}
	ZendAssertValidClassName(lcname)
	lcname = ZendNewInternedString(lcname)
	ZVAL_ALIAS_PTR(&zv, ce)
	ret = ZendHashAdd(CompilerGlobals.GetClassTable(), lcname, &zv)
	ZendStringReleaseEx(lcname, 0)
	if ret != nil {
		if !ce.IsImmutable() {
			ce.GetRefcount()++
		}
		return SUCCESS
	}
	return FAILURE
}
func ZendSetHashSymbol(symbol *Zval, name *byte, name_length int, is_ref ZendBool, num_symbol_tables int, _ ...any) int {
	var symbol_table *HashTable
	var symbol_table_list va_list
	if num_symbol_tables <= 0 {
		return FAILURE
	}
	if is_ref != 0 {
		ZVAL_MAKE_REF(symbol)
	}
	va_start(symbol_table_list, num_symbol_tables)
	for b.PostDec(&num_symbol_tables) > 0 {
		symbol_table = __va_arg(symbol_table_list, (*HashTable)(_))
		ZendHashStrUpdate(symbol_table, name, name_length, symbol)
		Z_TRY_ADDREF_P(symbol)
	}
	va_end(symbol_table_list)
	return SUCCESS
}
func ZifDisplayDisabledFunction(execute_data *ZendExecuteData, return_value *Zval) {
	ZendError(E_WARNING, "%s() has been disabled for security reasons", GetActiveFunctionName())
}
func ZendDisableFunction(function_name *byte, function_name_length int) int {
	var func_ *ZendInternalFunction
	if b.Assign(&func_, ZendHashStrFindPtr(CompilerGlobals.GetFunctionTable(), function_name, function_name_length)) {
		ZendFreeInternalArgInfo(func_)
		func_.SubFnFlags(ZEND_ACC_VARIADIC | ZEND_ACC_HAS_TYPE_HINTS | ZEND_ACC_HAS_RETURN_TYPE)
		func_.SetNumArgs(0)
		func_.SetArgInfo(nil)
		func_.SetHandler(ZifDisplayDisabledFunction)
		return SUCCESS
	}
	return FAILURE
}
func DisplayDisabledClass(class_type *ZendClassEntry) *ZendObject {
	var intern *ZendObject
	intern = ZendObjectsNew(class_type)

	/* Initialize default properties */

	if class_type.GetDefaultPropertiesCount() != 0 {
		var p *Zval = intern.GetPropertiesTable()
		var end *Zval = p + class_type.GetDefaultPropertiesCount()
		for {
			ZVAL_UNDEF(p)
			p++
			if p == end {
				break
			}
		}
	}
	ZendError(E_WARNING, "%s() has been disabled for security reasons", class_type.GetName().GetVal())
	return intern
}
func ZendDisableClass(class_name *byte, class_name_length int) int {
	var disabled_class *ZendClassEntry
	var key *ZendString
	var fn *ZendFunction
	key = ZendStringAlloc(class_name_length, 0)
	ZendStrTolowerCopy(key.GetVal(), class_name, class_name_length)
	disabled_class = ZendHashFindPtr(CompilerGlobals.GetClassTable(), key)
	ZendStringReleaseEx(key, 0)
	if disabled_class == nil {
		return FAILURE
	}
	INIT_CLASS_ENTRY_INIT_METHODS(*disabled_class, DisabledClassNew)
	disabled_class.create_object = DisplayDisabledClass
	for {
		var __ht *HashTable = disabled_class.GetFunctionTable()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			fn = _z.GetPtr()
			if fn.HasFnFlags(ZEND_ACC_HAS_RETURN_TYPE|ZEND_ACC_HAS_TYPE_HINTS) && fn.GetScope() == disabled_class {
				ZendFreeInternalArgInfo(fn.GetInternalFunction())
			}
		}
		break
	}
	ZendHashClean(disabled_class.GetFunctionTable())
	return SUCCESS
}
func ZendIsCallableCheckClass(name *ZendString, scope *ZendClassEntry, fcc *ZendFcallInfoCache, strict_class *int, error **byte) int {
	var ret int = 0
	var ce *ZendClassEntry
	var name_len int = name.GetLen()
	var lcname *ZendString
	ZSTR_ALLOCA_ALLOC(lcname, name_len, use_heap)
	ZendStrTolowerCopy(lcname.GetVal(), name.GetVal(), name_len)
	*strict_class = 0
	if ZendStringEqualsLiteral(lcname, "self") {
		if scope == nil {
			if error != nil {
				*error = Estrdup("cannot access self:: when no class scope is active")
			}
		} else {
			fcc.SetCalledScope(ZendGetCalledScope(ExecutorGlobals.GetCurrentExecuteData()))
			if fcc.GetCalledScope() == nil || InstanceofFunction(fcc.GetCalledScope(), scope) == 0 {
				fcc.SetCalledScope(scope)
			}
			fcc.SetCallingScope(scope)
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(ExecutorGlobals.GetCurrentExecuteData()))
			}
			ret = 1
		}
	} else if ZendStringEqualsLiteral(lcname, "parent") {
		if scope == nil {
			if error != nil {
				*error = Estrdup("cannot access parent:: when no class scope is active")
			}
		} else if !(scope.parent) {
			if error != nil {
				*error = Estrdup("cannot access parent:: when current class scope has no parent")
			}
		} else {
			fcc.SetCalledScope(ZendGetCalledScope(ExecutorGlobals.GetCurrentExecuteData()))
			if fcc.GetCalledScope() == nil || InstanceofFunction(fcc.GetCalledScope(), scope.parent) == 0 {
				fcc.SetCalledScope(scope.parent)
			}
			fcc.SetCallingScope(scope.parent)
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(ExecutorGlobals.GetCurrentExecuteData()))
			}
			*strict_class = 1
			ret = 1
		}
	} else if ZendStringEqualsLiteral(lcname, "static") {
		var called_scope *ZendClassEntry = ZendGetCalledScope(ExecutorGlobals.GetCurrentExecuteData())
		if called_scope == nil {
			if error != nil {
				*error = Estrdup("cannot access static:: when no class scope is active")
			}
		} else {
			fcc.SetCalledScope(called_scope)
			fcc.SetCallingScope(called_scope)
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(ExecutorGlobals.GetCurrentExecuteData()))
			}
			*strict_class = 1
			ret = 1
		}
	} else if b.Assign(&ce, ZendLookupClass(name)) != nil {
		var scope *ZendClassEntry
		var ex *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData()
		for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
			ex = ex.GetPrevExecuteData()
		}
		if ex != nil {
			scope = ex.GetFunc().GetScope()
		} else {
			scope = nil
		}
		fcc.SetCallingScope(ce)
		if scope != nil && fcc.GetObject() == nil {
			var object *ZendObject = ZendGetThisObject(ExecutorGlobals.GetCurrentExecuteData())
			if object != nil && InstanceofFunction(object.GetCe(), scope) != 0 && InstanceofFunction(scope, ce) != 0 {
				fcc.SetObject(object)
				fcc.SetCalledScope(object.GetCe())
			} else {
				fcc.SetCalledScope(ce)
			}
		} else {
			if fcc.GetObject() != nil {
				fcc.SetCalledScope(fcc.GetObject().GetCe())
			} else {
				fcc.SetCalledScope(ce)
			}
		}
		*strict_class = 1
		ret = 1
	} else {
		if error != nil {
			ZendSpprintf(error, 0, "class '%.*s' not found", int(name_len), name.GetVal())
		}
	}
	ZSTR_ALLOCA_FREE(lcname, use_heap)
	return ret
}
func ZendReleaseFcallInfoCache(fcc *ZendFcallInfoCache) {
	if fcc.GetFunctionHandler() != nil && (fcc.GetFunctionHandler().IsCallViaTrampoline() || fcc.GetFunctionHandler().GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY || fcc.GetFunctionHandler().GetType() == ZEND_OVERLOADED_FUNCTION) {
		if fcc.GetFunctionHandler().GetType() != ZEND_OVERLOADED_FUNCTION && fcc.GetFunctionHandler().GetFunctionName() != nil {
			ZendStringReleaseEx(fcc.GetFunctionHandler().GetFunctionName(), 0)
		}
		ZendFreeTrampoline(fcc.GetFunctionHandler())
	}
	fcc.SetFunctionHandler(nil)
}
func ZendIsCallableCheckFunc(check_flags int, callable *Zval, fcc *ZendFcallInfoCache, strict_class int, error **byte) int {
	var ce_org *ZendClassEntry = fcc.GetCallingScope()
	var retval int = 0
	var mname *ZendString
	var cname *ZendString
	var lmname *ZendString
	var colon *byte
	var clen int
	var ftable *HashTable
	var call_via_handler int = 0
	var scope *ZendClassEntry
	var zv *Zval
	fcc.SetCallingScope(nil)
	if ce_org == nil {
		var func_ *ZendFunction
		var lmname *ZendString

		/* Check if function with given name exists.
		 * This may be a compound name that includes namespace name */

		if Z_STRVAL_P(callable)[0] == '\\' {

			/* Skip leading \ */

			ZSTR_ALLOCA_ALLOC(lmname, Z_STRLEN_P(callable)-1, use_heap)
			ZendStrTolowerCopy(lmname.GetVal(), Z_STRVAL_P(callable)+1, Z_STRLEN_P(callable)-1)
			func_ = ZendFetchFunction(lmname)
			ZSTR_ALLOCA_FREE(lmname, use_heap)
		} else {
			lmname = callable.GetStr()
			func_ = ZendFetchFunction(lmname)
			if func_ == nil {
				ZSTR_ALLOCA_ALLOC(lmname, Z_STRLEN_P(callable), use_heap)
				ZendStrTolowerCopy(lmname.GetVal(), Z_STRVAL_P(callable), Z_STRLEN_P(callable))
				func_ = ZendFetchFunction(lmname)
				ZSTR_ALLOCA_FREE(lmname, use_heap)
			}
		}
		if func_ != nil {
			fcc.SetFunctionHandler(func_)
			return 1
		}
	}

	/* Split name into class/namespace and method/function names */

	if b.Assign(&colon, ZendMemrchr(Z_STRVAL_P(callable), ':', Z_STRLEN_P(callable))) != nil && colon > Z_STRVAL_P(callable) && (*(colon - 1)) == ':' {
		var mlen int
		colon--
		clen = colon - Z_STRVAL_P(callable)
		mlen = Z_STRLEN_P(callable) - clen - 2
		if colon == Z_STRVAL_P(callable) {
			if error != nil {
				*error = Estrdup("invalid function name")
			}
			return 0
		}

		/* This is a compound name.
		 * Try to fetch class and then find static method. */

		if ce_org != nil {
			scope = ce_org
		} else {
			scope = ZendGetExecutedScope()
		}
		cname = ZendStringInit(Z_STRVAL_P(callable), clen, 0)
		if ZendIsCallableCheckClass(cname, scope, fcc, &strict_class, error) == 0 {
			ZendStringReleaseEx(cname, 0)
			return 0
		}
		ZendStringReleaseEx(cname, 0)
		ftable = fcc.GetCallingScope().GetFunctionTable()
		if ce_org != nil && InstanceofFunction(ce_org, fcc.GetCallingScope()) == 0 {
			if error != nil {
				ZendSpprintf(error, 0, "class '%s' is not a subclass of '%s'", ce_org.GetName().GetVal(), fcc.GetCallingScope().GetName().GetVal())
			}
			return 0
		}
		mname = ZendStringInit(Z_STRVAL_P(callable)+clen+2, mlen, 0)
	} else if ce_org != nil {

		/* Try to fetch find static method of given class. */

		mname = callable.GetStr()
		ZendStringAddref(mname)
		ftable = ce_org.GetFunctionTable()
		fcc.SetCallingScope(ce_org)
	} else {

		/* We already checked for plain function before. */

		if error != nil && (check_flags&IS_CALLABLE_CHECK_SILENT) == 0 {
			ZendSpprintf(error, 0, "function '%s' not found or invalid function name", Z_STRVAL_P(callable))
		}
		return 0
	}
	lmname = ZendStringTolower(mname)
	if strict_class != 0 && fcc.GetCallingScope() != nil && ZendStringEqualsLiteral(lmname, ZEND_CONSTRUCTOR_FUNC_NAME) {
		fcc.SetFunctionHandler(fcc.GetCallingScope().GetConstructor())
		if fcc.GetFunctionHandler() != nil {
			retval = 1
		}
	} else if b.Assign(&zv, ZendHashFind(ftable, lmname)) != nil {
		fcc.SetFunctionHandler(zv.GetPtr())
		retval = 1
		if fcc.GetFunctionHandler().GetOpArray().IsChanged() && strict_class == 0 {
			scope = ZendGetExecutedScope()
			if scope != nil && InstanceofFunction(fcc.GetFunctionHandler().GetScope(), scope) != 0 {
				zv = ZendHashFind(scope.GetFunctionTable(), lmname)
				if zv != nil {
					var priv_fbc *ZendFunction = zv.GetPtr()
					if priv_fbc.IsPrivate() && priv_fbc.GetScope() == scope {
						fcc.SetFunctionHandler(priv_fbc)
					}
				}
			}
		}
		if !fcc.GetFunctionHandler().IsPublic() && (check_flags&IS_CALLABLE_CHECK_NO_ACCESS) == 0 && (fcc.GetCallingScope() != nil && (fcc.GetObject() != nil && fcc.GetCallingScope().GetCall() != nil || fcc.GetObject() == nil && fcc.GetCallingScope().GetCallstatic() != nil)) {
			scope = ZendGetExecutedScope()
			if fcc.GetFunctionHandler().GetScope() != scope {
				if fcc.GetFunctionHandler().IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(fcc.GetFunctionHandler()), scope) == 0 {
					retval = 0
					fcc.SetFunctionHandler(nil)
					goto get_function_via_handler
				}
			}
		}
	} else {
	get_function_via_handler:
		if fcc.GetObject() != nil && fcc.GetCallingScope() == ce_org {
			if strict_class != 0 && ce_org.GetCall() != nil {
				fcc.SetFunctionHandler(ZendGetCallTrampolineFunc(ce_org, mname, 0))
				call_via_handler = 1
				retval = 1
			} else {
				fcc.SetFunctionHandler(fcc.GetObject().GetHandlers().GetGetMethod()(fcc.GetObject(), mname, nil))
				if fcc.GetFunctionHandler() != nil {
					if strict_class != 0 && (fcc.GetFunctionHandler().GetScope() == nil || InstanceofFunction(ce_org, fcc.GetFunctionHandler().GetScope()) == 0) {
						ZendReleaseFcallInfoCache(fcc)
					} else {
						retval = 1
						call_via_handler = fcc.GetFunctionHandler().IsCallViaTrampoline()
					}
				}
			}
		} else if fcc.GetCallingScope() != nil {
			if fcc.GetCallingScope().GetGetStaticMethod() != nil {
				fcc.SetFunctionHandler(fcc.GetCallingScope().GetGetStaticMethod()(fcc.GetCallingScope(), mname))
			} else {
				fcc.SetFunctionHandler(ZendStdGetStaticMethod(fcc.GetCallingScope(), mname, nil))
			}
			if fcc.GetFunctionHandler() != nil {
				retval = 1
				call_via_handler = fcc.GetFunctionHandler().IsCallViaTrampoline()
				if call_via_handler != 0 && fcc.GetObject() == nil {
					var object *ZendObject = ZendGetThisObject(ExecutorGlobals.GetCurrentExecuteData())
					if object != nil && InstanceofFunction(object.GetCe(), fcc.GetCallingScope()) != 0 {
						fcc.SetObject(object)
					}
				}
			}
		}
	}
	if retval != 0 {
		if fcc.GetCallingScope() != nil && call_via_handler == 0 {
			if fcc.GetFunctionHandler().IsAbstract() {
				retval = 0
				if error != nil {
					ZendSpprintf(error, 0, "cannot call abstract method %s::%s()", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal())
				}
			} else if fcc.GetObject() == nil && !fcc.GetFunctionHandler().IsStatic() {
				var severity int
				var verb *byte
				if fcc.GetFunctionHandler().IsAllowStatic() {
					severity = E_DEPRECATED
					verb = "should not"
				} else {

					/* An internal function assumes $this is present and won't check that. So PHP would crash by allowing the call. */

					severity = E_ERROR
					verb = "cannot"
				}
				if (check_flags & IS_CALLABLE_CHECK_IS_STATIC) != 0 {
					retval = 0
				}
				if error != nil {
					ZendSpprintf(error, 0, "non-static method %s::%s() %s be called statically", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal(), verb)
					if severity != E_DEPRECATED {
						retval = 0
					}
				} else if retval != 0 {
					if severity == E_ERROR {
						ZendThrowError(nil, "Non-static method %s::%s() %s be called statically", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal(), verb)
					} else {
						ZendError(severity, "Non-static method %s::%s() %s be called statically", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal(), verb)
					}
				}
			}
			if retval != 0 && !fcc.GetFunctionHandler().IsPublic() && (check_flags&IS_CALLABLE_CHECK_NO_ACCESS) == 0 {
				scope = ZendGetExecutedScope()
				if fcc.GetFunctionHandler().GetScope() != scope {
					if fcc.GetFunctionHandler().IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(fcc.GetFunctionHandler()), scope) == 0 {
						if error != nil {
							if (*error) != nil {
								Efree(*error)
							}
							ZendSpprintf(error, 0, "cannot access %s method %s::%s()", ZendVisibilityString(fcc.GetFunctionHandler().GetFnFlags()), fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal())
						}
						retval = 0
					}
				}
			}
		}
	} else if error != nil && (check_flags&IS_CALLABLE_CHECK_SILENT) == 0 {
		if fcc.GetCallingScope() != nil {
			if error != nil {
				ZendSpprintf(error, 0, "class '%s' does not have a method '%s'", fcc.GetCallingScope().GetName().GetVal(), mname.GetVal())
			}
		} else {
			if error != nil {
				ZendSpprintf(error, 0, "function '%s' does not exist", mname.GetVal())
			}
		}
	}
	ZendStringReleaseEx(lmname, 0)
	ZendStringReleaseEx(mname, 0)
	if fcc.GetObject() != nil {
		fcc.SetCalledScope(fcc.GetObject().GetCe())
		if fcc.GetFunctionHandler() != nil && fcc.GetFunctionHandler().IsStatic() {
			fcc.SetObject(nil)
		}
	}
	return retval
}
func ZendCreateMethodString(class_name *ZendString, method_name *ZendString) *ZendString {
	var callable_name *ZendString = ZendStringAlloc(class_name.GetLen()+method_name.GetLen()+b.SizeOf("\"::\"")-1, 0)
	var ptr *byte = callable_name.GetVal()
	memcpy(ptr, class_name.GetVal(), class_name.GetLen())
	ptr += class_name.GetLen()
	memcpy(ptr, "::", b.SizeOf("\"::\"")-1)
	ptr += b.SizeOf("\"::\"") - 1
	memcpy(ptr, method_name.GetVal(), method_name.GetLen()+1)
	return callable_name
}
func ZendGetCallableNameEx(callable *Zval, object *ZendObject) *ZendString {
try_again:
	switch callable.GetType() {
	case IS_STRING:
		if object != nil {
			return ZendCreateMethodString(object.GetCe().GetName(), callable.GetStr())
		}
		return ZendStringCopy(callable.GetStr())
	case IS_ARRAY:
		var method *Zval = nil
		var obj *Zval = nil
		if Z_ARRVAL_P(callable).GetNNumOfElements() == 2 {
			obj = ZendHashIndexFindDeref(callable.GetArr(), 0)
			method = ZendHashIndexFindDeref(callable.GetArr(), 1)
		}
		if obj == nil || method == nil || method.GetType() != IS_STRING {
			return ZSTR_KNOWN(ZEND_STR_ARRAY_CAPITALIZED)
		}
		if obj.IsType(IS_STRING) {
			return ZendCreateMethodString(obj.GetStr(), method.GetStr())
		} else if obj.IsType(IS_OBJECT) {
			return ZendCreateMethodString(Z_OBJCE_P(obj).GetName(), method.GetStr())
		} else {
			return ZSTR_KNOWN(ZEND_STR_ARRAY_CAPITALIZED)
		}
	case IS_OBJECT:
		var calling_scope *ZendClassEntry
		var fptr *ZendFunction
		var object *ZendObject
		if Z_OBJ_HT(*callable).GetGetClosure() != nil && Z_OBJ_HT(*callable).GetGetClosure()(callable, &calling_scope, &fptr, &object) == SUCCESS {
			var ce *ZendClassEntry = Z_OBJCE_P(callable)
			var callable_name *ZendString = ZendStringAlloc(ce.GetName().GetLen()+b.SizeOf("\"::__invoke\"")-1, 0)
			memcpy(callable_name.GetVal(), ce.GetName().GetVal(), ce.GetName().GetLen())
			memcpy(callable_name.GetVal()+ce.GetName().GetLen(), "::__invoke", b.SizeOf("\"::__invoke\""))
			return callable_name
		}
		return ZvalGetString(callable)
	case IS_REFERENCE:
		callable = Z_REFVAL_P(callable)
		goto try_again
	default:
		return ZvalGetStringFunc(callable)
	}
}
func ZendGetCallableName(callable *Zval) *ZendString {
	return ZendGetCallableNameEx(callable, nil)
}
func ZendIsCallableImpl(callable *Zval, object *ZendObject, check_flags uint32, fcc *ZendFcallInfoCache, error **byte) ZendBool {
	var ret ZendBool
	var fcc_local ZendFcallInfoCache
	var strict_class int = 0
	if fcc == nil {
		fcc = &fcc_local
	}
	if error != nil {
		*error = nil
	}
	fcc.SetCallingScope(nil)
	fcc.SetCalledScope(nil)
	fcc.SetFunctionHandler(nil)
	fcc.SetObject(nil)
again:
	switch callable.GetType() {
	case IS_STRING:
		if object != nil {
			fcc.SetObject(object)
			fcc.SetCallingScope(object.GetCe())
		}
		if (check_flags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
			fcc.SetCalledScope(fcc.GetCallingScope())
			return 1
		}
	check_func:
		ret = ZendIsCallableCheckFunc(check_flags, callable, fcc, strict_class, error)
		if fcc == &fcc_local {
			ZendReleaseFcallInfoCache(fcc)
		}
		return ret
	case IS_ARRAY:
		var method *Zval = nil
		var obj *Zval = nil
		if Z_ARRVAL_P(callable).GetNNumOfElements() == 2 {
			obj = ZendHashIndexFind(callable.GetArr(), 0)
			method = ZendHashIndexFind(callable.GetArr(), 1)
		}
		for {
			if obj == nil || method == nil {
				break
			}
			ZVAL_DEREF(method)
			if method.GetType() != IS_STRING {
				break
			}
			ZVAL_DEREF(obj)
			if obj.IsType(IS_STRING) {
				if (check_flags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
					return 1
				}
				if ZendIsCallableCheckClass(obj.GetStr(), ZendGetExecutedScope(), fcc, &strict_class, error) == 0 {
					return 0
				}
			} else if obj.IsType(IS_OBJECT) {
				fcc.SetCallingScope(Z_OBJCE_P(obj))
				fcc.SetObject(obj.GetObj())
				if (check_flags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
					fcc.SetCalledScope(fcc.GetCallingScope())
					return 1
				}
			} else {
				break
			}
			callable = method
			goto check_func
			break
		}
		if Z_ARRVAL_P(callable).GetNNumOfElements() == 2 {
			if obj == nil || b.CondF(!(Z_ISREF_P(obj)), func() bool { return obj.GetType() != IS_STRING && obj.GetType() != IS_OBJECT }, func() bool { return Z_REFVAL_P(obj).GetType() != IS_STRING && Z_REFVAL_P(obj).GetType() != IS_OBJECT }) {
				if error != nil {
					*error = Estrdup("first array member is not a valid class name or object")
				}
			} else {
				if error != nil {
					*error = Estrdup("second array member is not a valid method")
				}
			}
		} else {
			if error != nil {
				*error = Estrdup("array must have exactly two members")
			}
		}
		return 0
	case IS_OBJECT:
		if Z_OBJ_HT(*callable).GetGetClosure() != nil {
			if Z_OBJ_HT(*callable).GetGetClosure()(callable, fcc.GetCallingScope(), fcc.GetFunctionHandler(), fcc.GetObject()) == SUCCESS {
				fcc.SetCalledScope(fcc.GetCallingScope())
				if fcc == &fcc_local {
					ZendReleaseFcallInfoCache(fcc)
				}
				return 1
			} else {

				/* Discard exceptions thrown from Z_OBJ_HANDLER_P(callable, get_closure) */

				ZendClearException()

				/* Discard exceptions thrown from Z_OBJ_HANDLER_P(callable, get_closure) */

			}
		}
		if error != nil {
			*error = Estrdup("no array or string given")
		}
		return 0
	case IS_REFERENCE:
		callable = Z_REFVAL_P(callable)
		goto again
	default:
		if error != nil {
			*error = Estrdup("no array or string given")
		}
		return 0
	}
}
func ZendIsCallableEx(callable *Zval, object *ZendObject, check_flags uint32, callable_name **ZendString, fcc *ZendFcallInfoCache, error **byte) ZendBool {
	var ret ZendBool = ZendIsCallableImpl(callable, object, check_flags, fcc, error)
	if callable_name != nil {
		*callable_name = ZendGetCallableNameEx(callable, object)
	}
	return ret
}
func ZendIsCallable(callable *Zval, check_flags uint32, callable_name **ZendString) ZendBool {
	return ZendIsCallableEx(callable, nil, check_flags, callable_name, nil, nil)
}
func ZendMakeCallable(callable *Zval, callable_name **ZendString) ZendBool {
	var fcc ZendFcallInfoCache
	if ZendIsCallableEx(callable, nil, IS_CALLABLE_STRICT, callable_name, &fcc, nil) != 0 {
		if callable.IsType(IS_STRING) && fcc.GetCallingScope() != nil {
			ZvalPtrDtorStr(callable)
			ArrayInit(callable)
			AddNextIndexStr(callable, ZendStringCopy(fcc.GetCallingScope().GetName()))
			AddNextIndexStr(callable, ZendStringCopy(fcc.GetFunctionHandler().GetFunctionName()))
		}
		ZendReleaseFcallInfoCache(&fcc)
		return 1
	}
	return 0
}
func ZendFcallInfoInit(callable *Zval, check_flags uint32, fci *ZendFcallInfo, fcc *ZendFcallInfoCache, callable_name **ZendString, error **byte) int {
	if ZendIsCallableEx(callable, nil, check_flags, callable_name, fcc, error) == 0 {
		return FAILURE
	}
	fci.SetSize(b.SizeOf("* fci"))
	fci.SetObject(fcc.GetObject())
	ZVAL_COPY_VALUE(fci.GetFunctionName(), callable)
	fci.SetRetval(nil)
	fci.SetParamCount(0)
	fci.SetParams(nil)
	fci.SetNoSeparation(1)
	return SUCCESS
}
func ZendFcallInfoArgsClear(fci *ZendFcallInfo, free_mem int) {
	if fci.GetParams() != nil {
		var p *Zval = fci.GetParams()
		var end *Zval = p + fci.GetParamCount()
		for p != end {
			IZvalPtrDtor(p)
			p++
		}
		if free_mem != 0 {
			Efree(fci.GetParams())
			fci.SetParams(nil)
		}
	}
	fci.SetParamCount(0)
}
func ZendFcallInfoArgsSave(fci *ZendFcallInfo, param_count *int, params **Zval) {
	*param_count = fci.GetParamCount()
	*params = fci.GetParams()
	fci.SetParamCount(0)
	fci.SetParams(nil)
}
func ZendFcallInfoArgsRestore(fci *ZendFcallInfo, param_count int, params *Zval) {
	ZendFcallInfoArgsClear(fci, 1)
	fci.SetParamCount(param_count)
	fci.SetParams(params)
}
func ZendFcallInfoArgsEx(fci *ZendFcallInfo, func_ *ZendFunction, args *Zval) int {
	var arg *Zval
	var params *Zval
	var n uint32 = 1
	ZendFcallInfoArgsClear(fci, !args)
	if args == nil {
		return SUCCESS
	}
	if args.GetType() != IS_ARRAY {
		return FAILURE
	}
	fci.SetParamCount(Z_ARRVAL_P(args).GetNNumOfElements())
	params = (*Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval")))
	fci.SetParams(params)
	for {
		var __ht *HashTable = args.GetArr()
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			arg = _z
			if func_ != nil && !(Z_ISREF_P(arg)) && ARG_SHOULD_BE_SENT_BY_REF(func_, n) != 0 {
				ZVAL_NEW_REF(params, arg)
				Z_TRY_ADDREF_P(arg)
			} else {
				ZVAL_COPY(params, arg)
			}
			params++
			n++
		}
		break
	}
	return SUCCESS
}
func ZendFcallInfoArgs(fci *ZendFcallInfo, args *Zval) int {
	return ZendFcallInfoArgsEx(fci, nil, args)
}
func ZendFcallInfoArgp(fci *ZendFcallInfo, argc int, argv *Zval) int {
	var i int
	if argc < 0 {
		return FAILURE
	}
	ZendFcallInfoArgsClear(fci, !argc)
	if argc != 0 {
		fci.SetParamCount(argc)
		fci.SetParams((*Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval"))))
		for i = 0; i < argc; i++ {
			ZVAL_COPY(fci.GetParams()[i], &argv[i])
		}
	}
	return SUCCESS
}
func ZendFcallInfoArgv(fci *ZendFcallInfo, argc int, argv *va_list) int {
	var i int
	var arg *Zval
	if argc < 0 {
		return FAILURE
	}
	ZendFcallInfoArgsClear(fci, !argc)
	if argc != 0 {
		fci.SetParamCount(argc)
		fci.SetParams((*Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval"))))
		for i = 0; i < argc; i++ {
			arg = __va_arg(*argv, (*Zval)(_))
			ZVAL_COPY(fci.GetParams()[i], arg)
		}
	}
	return SUCCESS
}
func ZendFcallInfoArgn(fci *ZendFcallInfo, argc int, _ ...any) int {
	var ret int
	var argv va_list
	va_start(argv, argc)
	ret = ZendFcallInfoArgv(fci, argc, &argv)
	va_end(argv)
	return ret
}
func ZendFcallInfoCall(fci *ZendFcallInfo, fcc *ZendFcallInfoCache, retval_ptr *Zval, args *Zval) int {
	var retval Zval
	var org_params *Zval = nil
	var result int
	var org_count int = 0
	if retval_ptr != nil {
		fci.SetRetval(retval_ptr)
	} else {
		fci.SetRetval(&retval)
	}
	if args != nil {
		ZendFcallInfoArgsSave(fci, &org_count, &org_params)
		ZendFcallInfoArgs(fci, args)
	}
	result = ZendCallFunction(fci, fcc)
	if retval_ptr == nil && retval.GetType() != IS_UNDEF {
		ZvalPtrDtor(&retval)
	}
	if args != nil {
		ZendFcallInfoArgsRestore(fci, org_count, org_params)
	}
	return result
}
func ZendGetModuleVersion(module_name *byte) *byte {
	var lname *ZendString
	var name_len int = strlen(module_name)
	var module *ZendModuleEntry
	lname = ZendStringAlloc(name_len, 0)
	ZendStrTolowerCopy(lname.GetVal(), module_name, name_len)
	module = ZendHashFindPtr(&ModuleRegistry, lname)
	ZendStringEfree(lname)
	if module != nil {
		return module.GetVersion()
	} else {
		return nil
	}
}
func ZvalMakeInternedString(zv *Zval) *ZendString {
	ZEND_ASSERT(zv.IsType(IS_STRING))
	zv.SetStr(ZendNewInternedString(zv.GetStr()))

	return zv.GetStr()
}
func IsPersistentClass(ce *ZendClassEntry) ZendBool {
	return (ce.GetType()&ZEND_INTERNAL_CLASS) != 0 && ce.GetModule().GetType() == MODULE_PERSISTENT
}
func ZendDeclareTypedProperty(ce *ZendClassEntry, name *ZendString, property *Zval, access_type int, doc_comment *ZendString, type_ ZendType) int {
	var property_info *ZendPropertyInfo
	var property_info_ptr *ZendPropertyInfo
	if type_.IsSet() {
		ce.SetIsHasTypeHints(true)
	}
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		property_info = Pemalloc(b.SizeOf("zend_property_info"), 1)
	} else {
		property_info = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_property_info"))
		if property.IsType(IS_CONSTANT_AST) {
			ce.SetIsConstantsUpdated(false)
		}
	}
	if property.IsType(IS_STRING) {
		ZvalMakeInternedString(property)
	}
	if (access_type & ZEND_ACC_PPP_MASK) == 0 {
		access_type |= ZEND_ACC_PUBLIC
	}
	if (access_type & ZEND_ACC_STATIC) != 0 {
		if b.Assign(&property_info_ptr, ZendHashFindPtr(ce.GetPropertiesInfo(), name)) != nil && property_info_ptr.IsStatic() {
			property_info.SetOffset(property_info_ptr.GetOffset())
			ZvalPtrDtor(ce.GetDefaultStaticMembersTable()[property_info.GetOffset()])
			ZendHashDel(ce.GetPropertiesInfo(), name)
		} else {
			ce.GetDefaultStaticMembersCount()++
			property_info.SetOffset(ce.GetDefaultStaticMembersCount() - 1)
			ce.SetDefaultStaticMembersTable(Perealloc(ce.GetDefaultStaticMembersTable(), b.SizeOf("zval")*ce.GetDefaultStaticMembersCount(), ce.GetType() == ZEND_INTERNAL_CLASS))
		}
		ZVAL_COPY_VALUE(ce.GetDefaultStaticMembersTable()[property_info.GetOffset()], property)
		if ce.GetStaticMembersTablePtr() == nil {
			ZEND_ASSERT(ce.GetType() == ZEND_INTERNAL_CLASS)
			if ExecutorGlobals.GetCurrentExecuteData() == nil {
				ZEND_MAP_PTR_NEW(ce.static_members_table)
			} else {

				/* internal class loaded by dl() */

				ZEND_MAP_PTR_INIT(ce.static_members_table, ce.GetDefaultStaticMembersTable())

				/* internal class loaded by dl() */

			}
		}
	} else {
		var property_default_ptr *Zval
		if b.Assign(&property_info_ptr, ZendHashFindPtr(ce.GetPropertiesInfo(), name)) != nil && !property_info_ptr.IsStatic() {
			property_info.SetOffset(property_info_ptr.GetOffset())
			ZvalPtrDtor(ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())])
			ZendHashDel(ce.GetPropertiesInfo(), name)
			ZEND_ASSERT(ce.GetType() == ZEND_INTERNAL_CLASS)
			ZEND_ASSERT(ce.GetPropertiesInfoTable() != nil)
			ce.GetPropertiesInfoTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())] = property_info
		} else {
			property_info.SetOffset(OBJ_PROP_TO_OFFSET(ce.GetDefaultPropertiesCount()))
			ce.GetDefaultPropertiesCount()++
			ce.SetDefaultPropertiesTable(Perealloc(ce.GetDefaultPropertiesTable(), b.SizeOf("zval")*ce.GetDefaultPropertiesCount(), ce.GetType() == ZEND_INTERNAL_CLASS))

			/* For user classes this is handled during linking */

			if ce.GetType() == ZEND_INTERNAL_CLASS {
				ce.SetPropertiesInfoTable(Perealloc(ce.GetPropertiesInfoTable(), b.SizeOf("zend_property_info *")*ce.GetDefaultPropertiesCount(), 1))
				ce.GetPropertiesInfoTable()[ce.GetDefaultPropertiesCount()-1] = property_info
			}

			/* For user classes this is handled during linking */

		}
		property_default_ptr = ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())]
		ZVAL_COPY_VALUE(property_default_ptr, property)
		if Z_ISUNDEF_P(property) {
			property_default_ptr.SetU2Extra(IS_PROP_UNINIT)
		} else {
			property_default_ptr.SetU2Extra(0)
		}
	}
	if (ce.GetType() & ZEND_INTERNAL_CLASS) != 0 {
		switch property.GetType() {
		case IS_ARRAY:

		case IS_OBJECT:

		case IS_RESOURCE:
			ZendErrorNoreturn(E_CORE_ERROR, "Internal zval's can't be arrays, objects or resources")
			break
		default:
			break
		}

		/* Must be interned to avoid ZTS data races */

		if IsPersistentClass(ce) != 0 {
			name = ZendNewInternedString(ZendStringCopy(name))
		}

		/* Must be interned to avoid ZTS data races */

	}
	if (access_type & ZEND_ACC_PUBLIC) != 0 {
		property_info.SetName(ZendStringCopy(name))
	} else if (access_type & ZEND_ACC_PRIVATE) != 0 {
		property_info.SetName(ZendManglePropertyName(ce.GetName().GetVal(), ce.GetName().GetLen(), name.GetVal(), name.GetLen(), IsPersistentClass(ce)))
	} else {
		ZEND_ASSERT((access_type & ZEND_ACC_PROTECTED) != 0)
		property_info.SetName(ZendManglePropertyName("*", 1, name.GetVal(), name.GetLen(), IsPersistentClass(ce)))
	}
	property_info.SetName(ZendNewInternedString(property_info.GetName()))
	property_info.SetFlags(access_type)
	property_info.SetDocComment(doc_comment)
	property_info.SetCe(ce)
	property_info.SetType(type_)
	ZendHashUpdatePtr(ce.GetPropertiesInfo(), name, property_info)
	return SUCCESS
}
func ZendTryAssignTypedRefEx(ref *ZendReference, val *Zval, strict ZendBool) int {
	if ZendVerifyRefAssignableZval(ref, val, strict) == 0 {
		ZvalPtrDtor(val)
		return FAILURE
	} else {
		ZvalPtrDtor(ref.GetVal())
		ZVAL_COPY_VALUE(ref.GetVal(), val)
		return SUCCESS
	}
}
func ZendTryAssignTypedRef(ref *ZendReference, val *Zval) int {
	return ZendTryAssignTypedRefEx(ref, val, ZEND_ARG_USES_STRICT_TYPES())
}
func ZendTryAssignTypedRefNull(ref *ZendReference) int {
	var tmp Zval
	ZVAL_NULL(&tmp)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefBool(ref *ZendReference, val ZendBool) int {
	var tmp Zval
	ZVAL_BOOL(&tmp, val)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefLong(ref *ZendReference, lval ZendLong) int {
	var tmp Zval
	ZVAL_LONG(&tmp, lval)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefDouble(ref *ZendReference, dval float64) int {
	var tmp Zval
	ZVAL_DOUBLE(&tmp, dval)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefEmptyString(ref *ZendReference) int {
	var tmp Zval
	ZVAL_EMPTY_STRING(&tmp)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefStr(ref *ZendReference, str *ZendString) int {
	var tmp Zval
	ZVAL_STR(&tmp, str)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefString(ref *ZendReference, string *byte) int {
	var tmp Zval
	ZVAL_STRING(&tmp, string)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefStringl(ref *ZendReference, string *byte, len_ int) int {
	var tmp Zval
	ZVAL_STRINGL(&tmp, string, len_)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefArr(ref *ZendReference, arr *ZendArray) int {
	var tmp Zval
	ZVAL_ARR(&tmp, arr)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefRes(ref *ZendReference, res *ZendResource) int {
	var tmp Zval
	ZVAL_RES(&tmp, res)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefZval(ref *ZendReference, zv *Zval) int {
	var tmp Zval
	ZVAL_COPY_VALUE(&tmp, zv)
	return ZendTryAssignTypedRef(ref, &tmp)
}
func ZendTryAssignTypedRefZvalEx(ref *ZendReference, zv *Zval, strict ZendBool) int {
	var tmp Zval
	ZVAL_COPY_VALUE(&tmp, zv)
	return ZendTryAssignTypedRefEx(ref, &tmp, strict)
}
func ZendDeclarePropertyEx(ce *ZendClassEntry, name *ZendString, property *Zval, access_type int, doc_comment *ZendString) int {
	return ZendDeclareTypedProperty(ce, name, property, access_type, doc_comment, 0)
}
func ZendDeclareProperty(ce *ZendClassEntry, name *byte, name_length int, property *Zval, access_type int) int {
	var key *ZendString = ZendStringInit(name, name_length, IsPersistentClass(ce))
	var ret int = ZendDeclarePropertyEx(ce, key, property, access_type, nil)
	ZendStringRelease(key)
	return ret
}
func ZendDeclarePropertyNull(ce *ZendClassEntry, name string, name_length int, access_type int) int {
	var property Zval
	ZVAL_NULL(&property)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyBool(ce *ZendClassEntry, name *byte, name_length int, value ZendLong, access_type int) int {
	var property Zval
	ZVAL_BOOL(&property, value)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyLong(ce *ZendClassEntry, name string, name_length int, value ZendLong, access_type int) int {
	var property Zval
	ZVAL_LONG(&property, value)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyDouble(ce *ZendClassEntry, name *byte, name_length int, value float64, access_type int) int {
	var property Zval
	ZVAL_DOUBLE(&property, value)
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyString(ce *ZendClassEntry, name string, name_length int, value string, access_type int) int {
	var property Zval
	ZVAL_NEW_STR(&property, ZendStringInit(value, strlen(value), ce.GetType()&ZEND_INTERNAL_CLASS))
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclarePropertyStringl(ce *ZendClassEntry, name *byte, name_length int, value *byte, value_len int, access_type int) int {
	var property Zval
	ZVAL_NEW_STR(&property, ZendStringInit(value, value_len, ce.GetType()&ZEND_INTERNAL_CLASS))
	return ZendDeclareProperty(ce, name, name_length, &property, access_type)
}
func ZendDeclareClassConstantEx(ce *ZendClassEntry, name *ZendString, value *Zval, access_type int, doc_comment *ZendString) int {
	var c *ZendClassConstant
	if ce.IsInterface() {
		if access_type != ZEND_ACC_PUBLIC {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Access type for interface constant %s::%s must be public", ce.GetName().GetVal(), name.GetVal())
		}
	}
	if ZendStringEqualsLiteralCi(name, "class") {
		ZendErrorNoreturn(b.Cond(ce.GetType() == ZEND_INTERNAL_CLASS, E_CORE_ERROR, E_COMPILE_ERROR), "A class constant must not be called 'class'; it is reserved for class name fetching")
	}
	if value.IsType(IS_STRING) {
		ZvalMakeInternedString(value)
	}
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		c = Pemalloc(b.SizeOf("zend_class_constant"), 1)
	} else {
		c = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_class_constant"))
	}
	ZVAL_COPY_VALUE(c.GetValue(), value)
	c.GetValue().GetAccessFlags() = access_type
	c.SetDocComment(doc_comment)
	c.SetCe(ce)
	if value.IsType(IS_CONSTANT_AST) {
		ce.SetIsConstantsUpdated(false)
	}
	if !(ZendHashAddPtr(ce.GetConstantsTable(), name, c)) {
		ZendErrorNoreturn(b.Cond(ce.GetType() == ZEND_INTERNAL_CLASS, E_CORE_ERROR, E_COMPILE_ERROR), "Cannot redefine class __special__  constant %s::%s", ce.GetName().GetVal(), name.GetVal())
	}
	return SUCCESS
}
func ZendDeclareClassConstant(ce *ZendClassEntry, name *byte, name_length int, value *Zval) int {
	var ret int
	var key *ZendString
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		key = ZendStringInitInterned(name, name_length, 1)
	} else {
		key = ZendStringInit(name, name_length, 0)
	}
	ret = ZendDeclareClassConstantEx(ce, key, value, ZEND_ACC_PUBLIC, nil)
	ZendStringRelease(key)
	return ret
}
func ZendDeclareClassConstantNull(ce *ZendClassEntry, name *byte, name_length int) int {
	var constant Zval
	ZVAL_NULL(&constant)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantLong(ce *ZendClassEntry, name string, name_length int, value ZendLong) int {
	var constant Zval
	ZVAL_LONG(&constant, value)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantBool(ce *ZendClassEntry, name *byte, name_length int, value ZendBool) int {
	var constant Zval
	ZVAL_BOOL(&constant, value)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantDouble(ce *ZendClassEntry, name *byte, name_length int, value float64) int {
	var constant Zval
	ZVAL_DOUBLE(&constant, value)
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantStringl(ce *ZendClassEntry, name *byte, name_length int, value *byte, value_length int) int {
	var constant Zval
	ZVAL_NEW_STR(&constant, ZendStringInit(value, value_length, ce.GetType()&ZEND_INTERNAL_CLASS))
	return ZendDeclareClassConstant(ce, name, name_length, &constant)
}
func ZendDeclareClassConstantString(ce *ZendClassEntry, name *byte, name_length int, value *byte) int {
	return ZendDeclareClassConstantStringl(ce, name, name_length, value, strlen(value))
}
func ZendUpdatePropertyEx(scope *ZendClassEntry, object *Zval, name *ZendString, value *Zval) {
	var property Zval
	var old_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	ExecutorGlobals.SetFakeScope(scope)
	ZVAL_STR(&property, name)
	Z_OBJ_HT_P(object).GetWriteProperty()(object, &property, value, nil)
	ExecutorGlobals.SetFakeScope(old_scope)
}
func ZendUpdateProperty(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *Zval) {
	var property Zval
	var old_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	ExecutorGlobals.SetFakeScope(scope)
	ZVAL_STRINGL(&property, name, name_length)
	Z_OBJ_HT_P(object).GetWriteProperty()(object, &property, value, nil)
	ZvalPtrDtor(&property)
	ExecutorGlobals.SetFakeScope(old_scope)
}
func ZendUpdatePropertyNull(scope *ZendClassEntry, object *Zval, name *byte, name_length int) {
	var tmp Zval
	ZVAL_NULL(&tmp)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUnsetProperty(scope *ZendClassEntry, object *Zval, name string, name_length int) {
	var property Zval
	var old_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	ExecutorGlobals.SetFakeScope(scope)
	ZVAL_STRINGL(&property, name, name_length)
	Z_OBJ_HT_P(object).GetUnsetProperty()(object, &property, 0)
	ZvalPtrDtor(&property)
	ExecutorGlobals.SetFakeScope(old_scope)
}
func ZendUpdatePropertyBool(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value ZendLong) {
	var tmp Zval
	ZVAL_BOOL(&tmp, value)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyLong(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value ZendLong) {
	var tmp Zval
	ZVAL_LONG(&tmp, value)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyDouble(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value float64) {
	var tmp Zval
	ZVAL_DOUBLE(&tmp, value)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyStr(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *ZendString) {
	var tmp Zval
	ZVAL_STR(&tmp, value)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyString(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *byte) {
	var tmp Zval
	ZVAL_STRING(&tmp, value)
	Z_SET_REFCOUNT(tmp, 0)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdatePropertyStringl(scope *ZendClassEntry, object *Zval, name *byte, name_length int, value *byte, value_len int) {
	var tmp Zval
	ZVAL_STRINGL(&tmp, value, value_len)
	Z_SET_REFCOUNT(tmp, 0)
	ZendUpdateProperty(scope, object, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyEx(scope *ZendClassEntry, name *ZendString, value *Zval) int {
	var property *Zval
	var tmp Zval
	var prop_info *ZendPropertyInfo
	var old_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	if !scope.IsConstantsUpdated() {
		if ZendUpdateClassConstants(scope) != SUCCESS {
			return FAILURE
		}
	}
	ExecutorGlobals.SetFakeScope(scope)
	property = ZendStdGetStaticPropertyWithInfo(scope, name, BP_VAR_W, &prop_info)
	ExecutorGlobals.SetFakeScope(old_scope)
	if property == nil {
		return FAILURE
	}
	ZEND_ASSERT(!(Z_ISREF_P(value)))
	Z_TRY_ADDREF_P(value)
	if prop_info.GetType() != 0 {
		ZVAL_COPY_VALUE(&tmp, value)
		if ZendVerifyPropertyType(prop_info, &tmp, 0) == 0 {
			Z_TRY_DELREF_P(value)
			return FAILURE
		}
		value = &tmp
	}
	ZendAssignToVariable(property, value, IS_TMP_VAR, 0)
	return SUCCESS
}
func ZendUpdateStaticProperty(scope *ZendClassEntry, name *byte, name_length int, value *Zval) int {
	var key *ZendString = ZendStringInit(name, name_length, 0)
	var retval int = ZendUpdateStaticPropertyEx(scope, key, value)
	ZendStringEfree(key)
	return retval
}
func ZendUpdateStaticPropertyNull(scope *ZendClassEntry, name *byte, name_length int) int {
	var tmp Zval
	ZVAL_NULL(&tmp)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyBool(scope *ZendClassEntry, name *byte, name_length int, value ZendLong) int {
	var tmp Zval
	ZVAL_BOOL(&tmp, value)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyLong(scope *ZendClassEntry, name *byte, name_length int, value ZendLong) int {
	var tmp Zval
	ZVAL_LONG(&tmp, value)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyDouble(scope *ZendClassEntry, name *byte, name_length int, value float64) int {
	var tmp Zval
	ZVAL_DOUBLE(&tmp, value)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyString(scope *ZendClassEntry, name *byte, name_length int, value *byte) int {
	var tmp Zval
	ZVAL_STRING(&tmp, value)
	Z_SET_REFCOUNT(tmp, 0)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendUpdateStaticPropertyStringl(scope *ZendClassEntry, name *byte, name_length int, value *byte, value_len int) int {
	var tmp Zval
	ZVAL_STRINGL(&tmp, value, value_len)
	Z_SET_REFCOUNT(tmp, 0)
	return ZendUpdateStaticProperty(scope, name, name_length, &tmp)
}
func ZendReadPropertyEx(scope *ZendClassEntry, object *Zval, name *ZendString, silent ZendBool, rv *Zval) *Zval {
	var property Zval
	var value *Zval
	var old_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	ExecutorGlobals.SetFakeScope(scope)
	ZVAL_STR(&property, name)
	value = Z_OBJ_HT_P(object).GetReadProperty()(object, &property, b.Cond(silent != 0, BP_VAR_IS, BP_VAR_R), nil, rv)
	ExecutorGlobals.SetFakeScope(old_scope)
	return value
}
func ZendReadProperty(scope *ZendClassEntry, object *Zval, name string, name_length int, silent ZendBool, rv *Zval) *Zval {
	var value *Zval
	var str *ZendString
	str = ZendStringInit(name, name_length, 0)
	value = ZendReadPropertyEx(scope, object, str, silent, rv)
	ZendStringReleaseEx(str, 0)
	return value
}
func ZendReadStaticPropertyEx(scope *ZendClassEntry, name *ZendString, silent ZendBool) *Zval {
	var property *Zval
	var old_scope *ZendClassEntry = ExecutorGlobals.GetFakeScope()
	ExecutorGlobals.SetFakeScope(scope)
	property = ZendStdGetStaticProperty(scope, name, b.Cond(silent != 0, BP_VAR_IS, BP_VAR_R))
	ExecutorGlobals.SetFakeScope(old_scope)
	return property
}
func ZendReadStaticProperty(scope *ZendClassEntry, name *byte, name_length int, silent ZendBool) *Zval {
	var key *ZendString = ZendStringInit(name, name_length, 0)
	var property *Zval = ZendReadStaticPropertyEx(scope, key, silent)
	ZendStringEfree(key)
	return property
}
func ZendSaveErrorHandling(current *ZendErrorHandling) {
	current.SetHandling(ExecutorGlobals.GetErrorHandling())
	current.SetException(ExecutorGlobals.GetExceptionClass())
	ZVAL_UNDEF(current.GetUserHandler())
}
func ZendReplaceErrorHandling(error_handling ZendErrorHandlingT, exception_class *ZendClassEntry, current *ZendErrorHandling) {
	if current != nil {
		ZendSaveErrorHandling(current)
	}
	ZEND_ASSERT(error_handling == EH_THROW || exception_class == nil)
	ExecutorGlobals.SetErrorHandling(error_handling)
	ExecutorGlobals.SetExceptionClass(exception_class)
}
func ZendRestoreErrorHandling(saved *ZendErrorHandling) {
	ExecutorGlobals.SetErrorHandling(saved.GetHandling())
	ExecutorGlobals.SetExceptionClass(saved.GetException())
}
func ZendFindAliasName(ce *ZendClassEntry, name *ZendString) *ZendString {
	var alias *ZendTraitAlias
	var alias_ptr **ZendTraitAlias
	if b.Assign(&alias_ptr, ce.GetTraitAliases()) {
		alias = *alias_ptr
		for alias != nil {
			if alias.GetAlias() != nil && ZendStringEqualsCi(alias.GetAlias(), name) {
				return alias.GetAlias()
			}
			alias_ptr++
			alias = *alias_ptr
		}
	}
	return name
}
func ZendResolveMethodName(ce *ZendClassEntry, f *ZendFunction) *ZendString {
	var func_ *ZendFunction
	var function_table *HashTable
	var name *ZendString
	if f.GetCommonType() != ZEND_USER_FUNCTION || f.GetOpArray().GetRefcount() != nil && (*(f.GetOpArray().GetRefcount())) < 2 || f.GetScope() == nil || f.GetScope().GetTraitAliases() == nil {
		return f.GetFunctionName()
	}
	function_table = ce.GetFunctionTable()
	for {
		var __ht *HashTable = function_table
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			name = _p.GetKey()
			func_ = _z.GetPtr()
			if func_ == f {
				if name == nil {
					return f.GetFunctionName()
				}
				if name.GetLen() == f.GetFunctionName().GetLen() && !(strncasecmp(name.GetVal(), f.GetFunctionName().GetVal(), f.GetFunctionName().GetLen())) {
					return f.GetFunctionName()
				}
				return ZendFindAliasName(f.GetScope(), name)
			}
		}
		break
	}
	return f.GetFunctionName()
}
func ZendGetObjectType(ce *ZendClassEntry) *byte {
	if ce.IsTrait() {
		return "trait"
	} else if ce.IsInterface() {
		return "interface"
	} else {
		return "class"
	}
}
func ZendIsIterable(iterable *Zval) ZendBool {
	switch iterable.GetType() {
	case IS_ARRAY:
		return 1
	case IS_OBJECT:
		return InstanceofFunction(Z_OBJCE_P(iterable), ZendCeTraversable)
	default:
		return 0
	}
}
func ZendIsCountable(countable *Zval) ZendBool {
	switch countable.GetType() {
	case IS_ARRAY:
		return 1
	case IS_OBJECT:
		if Z_OBJ_HT_P(countable).GetCountElements() != nil {
			return 1
		}
		return InstanceofFunction(Z_OBJCE_P(countable), ZendCeCountable)
	default:
		return 0
	}
}
