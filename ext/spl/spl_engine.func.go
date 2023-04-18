package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func SplInstantiateArgEx1(pce *types2.ClassEntry, retval *types2.Zval, arg1 *types2.Zval) int {
	var func_ types2.IFunction = pce.GetConstructor()
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, func_.GetFunctionName().GetStr(), nil, 1, arg1, nil)
	return 0
}
func SplInstantiateArgEx2(pce *types2.ClassEntry, retval *types2.Zval, arg1 *types2.Zval, arg2 *types2.Zval) int {
	var func_ types2.IFunction = pce.GetConstructor()
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, func_.GetFunctionName().GetStr(), nil, 2, arg1, arg2)
	return 0
}
func SplInstantiateArgN(pce *types2.ClassEntry, retval *types2.Zval, argc int, argv *types2.Zval) {
	var func_ types2.IFunction = pce.GetConstructor()
	var fci types2.ZendFcallInfo
	var fcc types2.ZendFcallInfoCache
	var dummy types2.Zval
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
func SplInstantiate(pce *types2.ClassEntry, object *types2.Zval) { zend.ObjectInitEx(object, pce) }
func SplOffsetConvertToLong(offset *types2.Zval) zend.ZendLong {
	var idx zend.ZendUlong
try_again:
	switch offset.GetType() {
	case types2.IS_STRING:
		if types2.HandleNumericStr(offset.String().GetStr(), &idx) {
			return idx
		}
	case types2.IS_DOUBLE:
		return zend.ZendLong(offset.Double())
	case types2.IS_LONG:
		return offset.Long()
	case types2.IS_FALSE:
		return 0
	case types2.IS_TRUE:
		return 1
	case types2.IS_REFERENCE:
		offset = types2.Z_REFVAL_P(offset)
		goto try_again
	case types2.IS_RESOURCE:
		return types2.Z_RES_HANDLE_P(offset)
	}
	return -1
}
