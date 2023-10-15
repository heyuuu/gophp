package spl

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func SplInstantiateArgEx1(pce *types.ClassEntry, retval *types.Zval, arg1 *types.Zval) int {
	var func_ types.IFunction = pce.GetConstructor()
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, func_.FunctionName(), nil, arg1)
	return 0
}
func SplInstantiateArgEx2(pce *types.ClassEntry, retval *types.Zval, arg1 *types.Zval, arg2 *types.Zval) int {
	var func_ types.IFunction = pce.GetConstructor()
	SplInstantiate(pce, retval)
	zend.ZendCallMethod(retval, pce, &func_, func_.FunctionName(), nil, arg1, arg2)
	return 0
}
func SplInstantiateArgN(pce *types.ClassEntry, retval *types.Zval, args []*types.Zval) {
	SplInstantiate(pce, retval)

	var func_ types.IFunction = pce.GetConstructor()
	var fci *types.ZendFcallInfo = types.InitFCallInfo(retval.Object(), nil, args...)
	fci.SetFunctionName(func_.FunctionName())

	var fcc types.ZendFcallInfoCache
	fcc.SetFunctionHandler(func_)
	fcc.SetCalledScope(pce)
	fcc.SetObject(retval.Object())
	zend.ZendCallFunction(fci, &fcc)
}
func SplInstantiate(pce *types.ClassEntry, object *types.Zval) { zend.ObjectInitEx(object, pce) }
func SplOffsetConvertToLong(offset *types.Zval) zend.ZendLong {
	var idx zend.ZendUlong
try_again:
	switch offset.Type() {
	case types.IsString:
		if types.HandleNumericStr(offset.String(), &idx) {
			return idx
		}
	case types.IsDouble:
		return zend.ZendLong(offset.Double())
	case types.IsLong:
		return offset.Long()
	case types.IsFalse:
		return 0
	case types.IsTrue:
		return 1
	case types.IsRef:
		offset = types.Z_REFVAL_P(offset)
		goto try_again
	case types.IsResource:
		return offset.ResourceHandle()
	}
	return -1
}
