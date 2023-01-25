// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/zend"
)

func SplInstantiateArgEx1(pce *zend.ZendClassEntry, retval *zend.Zval, arg1 *zend.Zval) int {
	var func_ *zend.ZendFunction = pce.GetConstructor()
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, zend.ZSTR_VAL(func_.GetFunctionName()), zend.ZSTR_LEN(func_.GetFunctionName()), nil, 1, arg1, nil)
	return 0
}
func SplInstantiateArgEx2(pce *zend.ZendClassEntry, retval *zend.Zval, arg1 *zend.Zval, arg2 *zend.Zval) int {
	var func_ *zend.ZendFunction = pce.GetConstructor()
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, zend.ZSTR_VAL(func_.GetFunctionName()), zend.ZSTR_LEN(func_.GetFunctionName()), nil, 2, arg1, arg2)
	return 0
}
func SplInstantiateArgN(pce *zend.ZendClassEntry, retval *zend.Zval, argc int, argv *zend.Zval) {
	var func_ *zend.ZendFunction = pce.GetConstructor()
	var fci zend.ZendFcallInfo
	var fcc zend.ZendFcallInfoCache
	var dummy zend.Zval
	SplInstantiate(pce, retval)
	fci.SetSize(b.SizeOf("zend_fcall_info"))
	zend.ZVAL_STR(&fci.GetFunctionName(), func_.GetFunctionName())
	fci.SetObject(zend.Z_OBJ_P(retval))
	fci.SetRetval(&dummy)
	fci.SetParamCount(argc)
	fci.SetParams(argv)
	fci.SetNoSeparation(1)
	fcc.SetFunctionHandler(func_)
	fcc.SetCalledScope(pce)
	fcc.SetObject(zend.Z_OBJ_P(retval))
	zend.ZendCallFunction(&fci, &fcc)
}
func SplInstantiate(pce *zend.ZendClassEntry, object *zend.Zval) { zend.ObjectInitEx(object, pce) }
func SplOffsetConvertToLong(offset *zend.Zval) zend.ZendLong {
	var idx zend.ZendUlong
try_again:
	switch zend.Z_TYPE_P(offset) {
	case zend.IS_STRING:
		if zend.ZEND_HANDLE_NUMERIC(zend.Z_STR_P(offset), idx) != 0 {
			return idx
		}
		break
	case zend.IS_DOUBLE:
		return zend.ZendLong(zend.Z_DVAL_P(offset))
	case zend.IS_LONG:
		return zend.Z_LVAL_P(offset)
	case zend.IS_FALSE:
		return 0
	case zend.IS_TRUE:
		return 1
	case zend.IS_REFERENCE:
		offset = zend.Z_REFVAL_P(offset)
		goto try_again
	case zend.IS_RESOURCE:
		return zend.Z_RES_HANDLE_P(offset)
	}
	return -1
}
