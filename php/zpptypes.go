package php

import (
	"github.com/heyuuu/gophp/php/types"
	"math"
)

// -- parse each types

func zppParseBoolWeak(ctx *Context, arg types.Zval) (dest bool, ok bool) {
	if arg.Type() <= types.IsString {
		return ZvalIsTrue(ctx, arg), true
	}
	return false, false
}

func zppParseLongWeak(ctx *Context, arg types.Zval, cap bool) (dest int, ok bool) {
	// 字符串类型尝试转数字
	if arg.IsString() {
		arg = opParseNumberPrefix(ctx, arg.String(), false)
		if arg.IsUndef() || ctx.EG().HasException() {
			return // fail
		}
	}

	switch arg.Type() {
	case types.IsUndef, types.IsNull, types.IsFalse:
		dest = 0
	case types.IsTrue:
		dest = 1
	case types.IsLong:
		dest = arg.Long()
	case types.IsDouble:
		return zppParseArgWeak_DvalToLval(arg.Double(), cap)
	default:
		return // fail
	}
	// success
	return dest, true
}

func zppParseArgWeak_DvalToLval(dval float64, cap bool) (int, bool) {
	if math.IsNaN(dval) {
		return 0, false
	}
	if cap {
		return DoubleToLongCap(dval), true
	} else {
		if dval < math.MinInt || dval >= math.MaxInt {
			return 0, false
		}
		return DoubleToLong(dval), true
	}
}

func zppParseDoubleWeak(ctx *Context, arg types.Zval) (dest float64, ok bool) {
	// 字符串类型尝试转数字
	if arg.IsString() {
		arg, _ = ParseNumberPrefix(arg.String())
		if arg.IsUndef() || ctx.EG().HasException() {
			return // fail
		}
	}

	switch arg.Type() {
	case types.IsUndef, types.IsNull, types.IsFalse:
		dest = 0
	case types.IsTrue:
		dest = 1
	case types.IsLong:
		dest = float64(arg.Long())
	case types.IsDouble:
		dest = arg.Double()
	default:
		return // fail
	}
	// success
	return dest, true
}

func zppParseStrWeak(ctx *Context, arg types.Zval) (string, bool) {
	if arg.Type() < types.IsString {
		return ZvalTryGetStr(ctx, arg)
	} else if arg.IsString() {
		return arg.String(), true
	} else if arg.IsObject() {
		if tmp, ok := arg.Object().Cast(types.IsString); ok && tmp.IsString() {
			return tmp.String(), true
		}
		return "", false
	} else {
		return "", false
	}
}

func zppParseArray(arg types.Zval, checkNull bool, orObject bool) (dest types.Zval, ok bool) {
	if arg.IsArray() || (orObject && arg.IsObject()) {
		return arg, true
	} else if checkNull && arg.IsNull() {
		return types.Undef, true
	} else {
		return types.Undef, false
	}
}

func zppParseArrayHt(arg types.Zval, checkNull bool, orObject bool, separate bool) (dest *types.Array, ok bool) {
	if arg.IsArray() {
		return arg.Array(), true
		//} else if orObject && arg.IsObject() {
		//	if separate && arg.Object().GetProperties() != nil {
		//		arg.Object().DupProperties()
		//	}
		//	return arg.Object().GetPropertiesArray(), true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

func zppParseObject(arg types.Zval, ce *types.Class, checkNull bool) (dest *types.Object, ok bool) {
	if arg.IsObject() { //&& (ce == nil || InstanceofFunction(arg.Object().Class(), ce)) {
		return arg.Object(), true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

func zppParseResource(arg *types.Zval, checkNull bool) (dest *types.Zval, ok bool) {
	if arg.IsResource() {
		return arg, true
	} else if checkNull && arg.IsNull() {
		return nil, true
	} else {
		return nil, false
	}
}

//
//func zppParseFunc(ctx *Context, arg types.Zval, dest_fci *types.ZendFcallInfo, dest_fcc *types.ZendFcallInfoCache, checkNull bool) (error *string, ok bool) {
//	if checkNull && arg.IsNull() {
//		dest_fci.UnInit()
//		dest_fcc.SetFunctionHandler(nil)
//		return nil, true
//	}
//
//	// notice: 此处在成功时 error 也有可能不为 nil (例如在产生 Deprecated 信息时)
//	state := ZendFcallInfoInit(ctx, &arg, 0, dest_fci, dest_fcc, nil, &error)
//	return error, state == types.SUCCESS
//}

func zppParseZvalDeref(arg *types.Zval, checkNull bool) (dest *types.Zval) {
	if checkNull && arg.IsNull() {
		return nil
	} else {
		return arg
	}
}

//
//func zppParseClass(ctx *Context, arg types.Zval, baseCe *types.Class, num int, checkNull bool) (ce *types.Class, ok bool) {
//	if checkNull && arg.IsNull() {
//		return nil, true
//	}
//
//	if s, ok := ZvalTryGetStr(ctx, arg); ok {
//		arg.SetString(s)
//	} else {
//		return nil, false
//	}
//	ce = ZendLookupClass(ctx, arg.String())
//	if baseCe != nil {
//		if ce == nil || !InstanceofFunction(ce, baseCe) {
//			faults.InternalTypeError(ctx, CurrEX(ctx).IsArgUseStrictTypes(), fmt.Sprintf("%s() expects parameter %d to be a class name derived from %s, '%s' given", CurrEX(ctx).CalleeName(), num, baseCe.Name(), arg.String()))
//			return nil, false
//		}
//	}
//	if ce == nil {
//		faults.InternalTypeError(ctx, CurrEX(ctx).IsArgUseStrictTypes(), fmt.Sprintf("%s() expects parameter %d to be a valid class name, '%s' given", CurrEX(ctx).CalleeName(), num, arg.String()))
//		return nil, false
//	}
//	return ce, true
//}
