package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func CE_STATIC_MEMBERS(ce *types.ClassEntry) *types.Zval {
	// todo
	return (ZEND_MAP_PTR_GET(ce.static_members_table__ptr)).(*types.Zval)
}
func ArrayInit(arg *types.Zval)                  { arg.SetArray(types.NewArray(0)) }
func ArrayInitSize(arg *types.Zval, size uint32) { arg.SetArray(types.NewArray(size)) }
func AddAssocLong(arg *types.Zval, key string, n ZendLong) int {
	return AddAssocLongEx(arg, key, n)
}
func AddAssocNull(arg *types.Zval, key string) int {
	return AddAssocNullEx(arg, key)
}
func AddAssocBool(arg *types.Zval, key string, __b int) int {
	return AddAssocBoolEx(arg, key, __b)
}
func AddAssocDouble(arg *types.Zval, key string, __d float64) int {
	return AddAssocDoubleEx(arg, key, __d)
}
func AddAssocStr(arg *types.Zval, key string, __str string) int {
	return AddAssocStrEx(arg, key, __str)
}
func AddAssocString(arg *types.Zval, key string, __str *byte) int {
	return AddAssocStrEx(arg, key, b.CastStrAuto(__str))
}
func AddAssocStringl(arg *types.Zval, key string, __str *byte, __length int) int {
	return AddAssocStringlEx(arg, key, b.CastStr(__str, __length))
}
func AddAssocZval(arg *types.Zval, key string, __value *types.Zval) int {
	return AddAssocZvalEx(arg, key, __value)
}
func AddIndexZval(arg *types.Zval, index ZendUlong, value *types.Zval) int {
	if arg.Array().IndexUpdate(index, value) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexZval(arg *types.Zval, value *types.Zval) int {
	if arg.Array().Append(value) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddPropertyLong(arg *types.Zval, key string, n ZendLong) int {
	return AddPropertyLongEx(arg, key, n)
}
func AddPropertyNull(arg *types.Zval, key string) int {
	return AddPropertyNullEx(arg, key)
}
func AddPropertyResource(arg *types.Zval, key string, r *types.ZendResource) int {
	return AddPropertyResourceEx(arg, key, r)
}
func AddPropertyString(arg *types.Zval, key string, __str string) int {
	return AddPropertyStrEx(arg, key, __str)
}
func AddPropertyStringl(arg *types.Zval, key string, __str string) int {
	return AddPropertyStrEx(arg, key, __str)
}
func AddPropertyZval(arg *types.Zval, key string, __value *types.Zval) int {
	return AddPropertyZvalEx(arg, key, __value)
}
func CallUserFunction(object *types.Zval, functionName *types.Zval, retvalPtr *types.Zval, paramCount uint32, params []types.Zval) int {
	return CallUserFunctionEx(object, functionName, retvalPtr, paramCount, params, 1)
}
func CallUserFunction_Ex(object *types.Zval, functionName *types.Zval, retvalPtr *types.Zval, params []types.Zval) int {
	return CallUserFunctionEx(object, functionName, retvalPtr, uint32(len(params)), params, 1)
}
func CallUserFunctionEx(object *types.Zval, functionName *types.Zval, retvalPtr *types.Zval, paramCount uint32, params []types.Zval, no_separation int) int {
	b.Assert(int(paramCount) == len(params))

	var objPtr *types.ZendObject
	if objPtr != nil {
		objPtr = object.Object()
	}

	var fci = types.InitFCallInfo(objPtr, retvalPtr, params...)
	fci.SetFunctionNameZval(functionName)
	fci.SetNoSeparation(no_separation != 0)

	return ZendCallFunction(fci, nil)
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
func ZVAL_ZVAL(z *types.Zval, zv *types.Zval, copy int, dtor int) {
	var __z *types.Zval = z
	var __zv *types.Zval = zv
	if !(__zv.IsReference()) {
		if copy != 0 && dtor == 0 {
			types.ZVAL_COPY(__z, __zv)
		} else {
			__z.CopyValueFrom(__zv)
		}
	} else {
		types.ZVAL_COPY(__z, types.Z_REFVAL_P(__zv))
	}
}

func ZvalZval(zv *types.Zval, copy bool, dtor bool) *types.Zval {
	var z types.Zval
	if zv.IsReference() {
		z.CopyFrom(zv.DeRef())
	} else {
		if copy && !dtor {
			z.CopyFrom(zv)
		} else {
			z.CopyValueFrom(zv)
		}
	}

	return &z
}

func HASH_OF(p *types.Zval) *types.Array {
	if p.IsArray() {
		return p.Array()
	} else {
		if p.IsObject() {
			return p.Object().GetPropertiesArray()
		} else {
			return nil
		}
	}
}
func _ZEND_TRY_ASSIGN_NULL(zv *types.Zval, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.Reference()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefNull(ref)
				break
			}
			_zv = ref.GetVal()
		}
		// ZvalPtrDtor(_zv)
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
			var ref *types.ZendReference = _zv.Reference()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefLong(ref, lval)
				break
			}
			_zv = ref.GetVal()
		}
		// ZvalPtrDtor(_zv)
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
			var ref *types.ZendReference = _zv.Reference()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefDouble(ref, dval)
				break
			}
			_zv = ref.GetVal()
		}
		// ZvalPtrDtor(_zv)
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
			var ref *types.ZendReference = _zv.Reference()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefEmptyString(ref)
				break
			}
			_zv = ref.GetVal()
		}
		// ZvalPtrDtor(_zv)
		_zv.SetStringVal("")
		break
	}
}
func ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zv *types.Zval) {
	b.Assert(zv.IsReference())
	_ZEND_TRY_ASSIGN_EMPTY_STRING(zv, 1)
}
func _ZEND_TRY_ASSIGN_STR(zv *types.Zval, str *types.String, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.Reference()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefStr(ref, str)
				break
			}
			_zv = ref.GetVal()
		}
		// ZvalPtrDtor(_zv)
		_zv.SetString(str)
		break
	}
}
func ZEND_TRY_ASSIGN_REF_STR(zv *types.Zval, str *types.String) {
	b.Assert(zv.IsReference())
	_ZEND_TRY_ASSIGN_STR(zv, str, 1)
}
func _ZEND_TRY_ASSIGN_STRING(zv *types.Zval, string *byte, is_ref int) {
	for {
		var _zv *types.Zval = zv
		if is_ref != 0 || _zv.IsReference() {
			var ref *types.ZendReference = _zv.Reference()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefString(ref, string)
				break
			}
			_zv = ref.GetVal()
		}
		// ZvalPtrDtor(_zv)
		_zv.SetStringVal(b.CastStrAuto(string))
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
			var ref *types.ZendReference = _zv.Reference()
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendTryAssignTypedRefStringl(ref, string, len_)
				break
			}
			_zv = ref.GetVal()
		}
		// ZvalPtrDtor(_zv)
		_zv.SetStringVal(b.CastStr(string, len_))
		break
	}
}
func ZEND_TRY_ASSIGN_REF_STRINGL(zv *types.Zval, string *byte, len_ int) {
	b.Assert(zv.IsReference())
	_ZEND_TRY_ASSIGN_STRINGL(zv, string, len_, 1)
}
