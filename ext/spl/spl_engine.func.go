// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/zend"
)

func SplInstantiateArgEx1(pce *zend.ZendClassEntry, retval *zend.Zval, arg1 *zend.Zval) int {
	var func_ *zend.ZendFunction = pce.GetConstructor()
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, func_.GetFunctionName().GetVal(), func_.GetFunctionName().GetLen(), nil, 1, arg1, nil)
	return 0
}
func SplInstantiateArgEx2(pce *zend.ZendClassEntry, retval *zend.Zval, arg1 *zend.Zval, arg2 *zend.Zval) int {
	var func_ *zend.ZendFunction = pce.GetConstructor()
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, func_.GetFunctionName().GetVal(), func_.GetFunctionName().GetLen(), nil, 2, arg1, arg2)
	return 0
}
func SplInstantiateArgN(pce *zend.ZendClassEntry, retval *zend.Zval, argc int, argv *zend.Zval) {
	var func_ *zend.ZendFunction = pce.GetConstructor()
	var fci zend.ZendFcallInfo
	var fcc zend.ZendFcallInfoCache
	var dummy zend.Zval
	SplInstantiate(pce, retval)
	fci.SetSize(b.SizeOf("zend_fcall_info"))
	fci.GetFunctionName().SetString(func_.GetFunctionName())
	fci.SetObject(retval.GetObj())
	fci.SetRetval(&dummy)
	fci.SetParamCount(argc)
	fci.SetParams(argv)
	fci.SetNoSeparation(1)
	fcc.SetFunctionHandler(func_)
	fcc.SetCalledScope(pce)
	fcc.SetObject(retval.GetObj())
	zend.ZendCallFunction(&fci, &fcc)
}
func SplInstantiate(pce *zend.ZendClassEntry, object *zend.Zval) { zend.ObjectInitEx(object, pce) }
func SplOffsetConvertToLong(offset *zend.Zval) zend.ZendLong {
	var idx zend.ZendUlong
try_again:
	switch offset.GetType() {
	case zend.IS_STRING:
		if zend.ZEND_HANDLE_NUMERIC(offset.GetStr(), &idx) {
			return idx
		}
	case zend.IS_DOUBLE:
		return zend.ZendLong(offset.GetDval())
	case zend.IS_LONG:
		return offset.GetLval()
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
