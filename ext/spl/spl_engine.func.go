package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func SplInstantiateArgEx1(pce *types.ClassEntry, retval *types.Zval, arg1 *types.Zval) int {
	var func_ types.IFunction = pce.GetConstructor()
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, func_.GetFunctionName().GetStr(), nil, 1, arg1, nil)
	return 0
}
func SplInstantiateArgEx2(pce *types.ClassEntry, retval *types.Zval, arg1 *types.Zval, arg2 *types.Zval) int {
	var func_ types.IFunction = pce.GetConstructor()
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, func_.GetFunctionName().GetStr(), nil, 2, arg1, arg2)
	return 0
}
func SplInstantiateArgN(pce *types.ClassEntry, retval *types.Zval, argc int, argv *types.Zval) {
	var func_ types.IFunction = pce.GetConstructor()
	var fci types.ZendFcallInfo
	var fcc types.ZendFcallInfoCache
	var dummy types.Zval
	SplInstantiate(pce, retval)
	fci.SetSize(b.SizeOf("zend_fcall_info"))
	fci.GetFunctionName().SetString(func_.GetFunctionName())
	fci.SetObject(retval.Object())
	fci.SetRetval(&dummy)
	fci.SetParamCount(argc)
	fci.SetParams(argv)
	fci.SetNoSeparation(1)
	fcc.SetFunctionHandler(func_)
	fcc.SetCalledScope(pce)
	fcc.SetObject(retval.Object())
	zend.ZendCallFunction(&fci, &fcc)
}
func SplInstantiate(pce *types.ClassEntry, object *types.Zval) { zend.ObjectInitEx(object, pce) }
func SplOffsetConvertToLong(offset *types.Zval) zend.ZendLong {
	var idx zend.ZendUlong
try_again:
	switch offset.GetType() {
	case types.IS_STRING:
		if types.HandleNumericStr(offset.String().GetStr(), &idx) {
			return idx
		}
	case types.IS_DOUBLE:
		return zend.ZendLong(offset.Double())
	case types.IS_LONG:
		return offset.Long()
	case types.IS_FALSE:
		return 0
	case types.IS_TRUE:
		return 1
	case types.IS_REFERENCE:
		offset = types.Z_REFVAL_P(offset)
		goto try_again
	case types.IS_RESOURCE:
		return offset.ResourceHandle()
	}
	return -1
}
