package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZendIsCallableImpl(callable *types.Zval, object *types.Object, checkFlags uint32, fcc *types.ZendFcallInfoCache, error **byte) bool {
	var ret bool
	var fcc_local types.ZendFcallInfoCache
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
	case types.IsString:
		if object != nil {
			fcc.SetObject(object)
			fcc.SetCallingScope(object.GetCe())
		}
		if (checkFlags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
			fcc.SetCalledScope(fcc.GetCallingScope())
			return 1
		}
	check_func:
		ret = ZendIsCallableCheckFunc(checkFlags, callable, fcc, strict_class, error)
		if fcc == &fcc_local {
			ZendReleaseFcallInfoCache(fcc)
		}
		return ret
	case types.IsArray:
		var method *types.Zval = nil
		var obj *types.Zval = nil
		if callable.Array().Len() == 2 {
			obj = callable.Array().IndexFind(0)
			method = callable.Array().IndexFind(1)
		}
		for {
			if obj == nil || method == nil {
				break
			}
			method = types.ZVAL_DEREF(method)
			if !method.IsString() {
				break
			}
			obj = types.ZVAL_DEREF(obj)
			if obj.IsString() {
				if (checkFlags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
					return 1
				}
				if ZendIsCallableCheckClass(obj.String(), ZendGetExecutedScope(), fcc, &strict_class, error) == 0 {
					return 0
				}
			} else if obj.IsObject() {
				fcc.SetCallingScope(types.Z_OBJCE_P(obj))
				fcc.SetObject(obj.Object())
				if (checkFlags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
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
		if callable.Array().Len() == 2 {
			if obj == nil || lang.CondF(!(obj.IsReference()), func() bool { return !obj.IsString() && !obj.IsObject() }, func() bool {
				return types.Z_REFVAL_P(obj).GetType() != types.IsString && types.Z_REFVAL_P(obj).GetType() != types.IsObject
			}) {
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
	case types.IsObject:
		if callable.Object().CanGetClosure() {
			if callable.Object().GetClosure(callable, fcc.GetCallingScope(), fcc.GetFunctionHandler(), fcc.GetObject()) == types.SUCCESS {
				fcc.SetCalledScope(fcc.GetCallingScope())
				if fcc == &fcc_local {
					ZendReleaseFcallInfoCache(fcc)
				}
				return 1
			} else {

				/* Discard exceptions thrown from Z_OBJ_HANDLER_P(callable, get_closure) */

				faults.ClearException()
			}
		}
		if error != nil {
			*error = Estrdup("no array or string given")
		}
		return 0
	case types.IsRef:
		callable = types.Z_REFVAL_P(callable)
		goto again
	default:
		if error != nil {
			*error = Estrdup("no array or string given")
		}
		return 0
	}
}
func ZendIsCallableEx(
	callable *types.Zval,
	object *types.Object,
	checkFlags uint32,
	callableName **types.String,
	fcc *types.ZendFcallInfoCache,
	error **byte,
) bool {
	var ret bool = ZendIsCallableImpl(callable, object, checkFlags, fcc, error)
	if callableName != nil {
		*callableName = types.NewString(ZendGetCallableNameEx(callable, object))
	}
	return ret
}
func ZendIsCallable(callable *types.Zval, check_flags uint32, callable_name **types.String) bool {
	return ZendIsCallableEx(callable, nil, check_flags, callable_name, nil, nil)
}
func ZendFcallInfoInit(
	callable *types.Zval,
	checkFlags uint32,
	fci *types.ZendFcallInfo,
	fcc *types.ZendFcallInfoCache,
	callableName **types.String,
	error **byte,
) int {
	if !ZendIsCallableEx(callable, nil, checkFlags, callableName, fcc, error) {
		return types.FAILURE
	}
	*fci = *types.InitFCallInfo(fcc.GetObject(), nil)
	fci.SetFunctionNameZval(callable)
	return types.SUCCESS
}
func ZendFcallInfoArgsClear(fci *types.ZendFcallInfo, free_mem int) {
	if fci.GetParams() != nil {
		var p *types.Zval = fci.GetParams()
		var end *types.Zval = p + fci.GetParamCount()
		for p != end {
			// IZvalPtrDtor(p)
			p++
		}
		if free_mem != 0 {
			Efree(fci.GetParams())
			fci.SetParams(nil)
		}
	}
	fci.SetParamCount(0)
}
func ZendFcallInfoArgsSave(fci *types.ZendFcallInfo, param_count *int, params **types.Zval) {
	*param_count = fci.GetParamCount()
	*params = fci.GetParams()
	fci.SetParamCount(0)
	fci.SetParams(nil)
}
func ZendFcallInfoArgsRestore(fci *types.ZendFcallInfo, param_count int, params *types.Zval) {
	ZendFcallInfoArgsClear(fci, 1)
	fci.SetParamCount(param_count)
	fci.SetParams(params)
}
func ZendFcallInfoArgsEx(fci *types.ZendFcallInfo, func_ types.IFunction, args *types.Zval) int {
	var arg *types.Zval
	var params *types.Zval
	var n uint32 = 1
	ZendFcallInfoArgsClear(fci, !args)
	if args == nil {
		return types.SUCCESS
	}
	if !args.IsArray() {
		return types.FAILURE
	}
	fci.SetParamCount(args.Array().Len())
	params = (*types.Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval")))
	fci.SetParams(params)
	var __ht *types.Array = args.Array()
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		arg = _z
		if func_ != nil && !(arg.IsReference()) && ARG_SHOULD_BE_SENT_BY_REF(func_, n) != 0 {
			params.SetNewRef(arg)
			// arg.TryAddRefcount()
		} else {
			types.ZVAL_COPY(params, arg)
		}
		params++
		n++
	}
	return types.SUCCESS
}

func ZendFcallInfoArgs(fci *types.ZendFcallInfo, args *types.Zval) int {
	return ZendFcallInfoArgsEx(fci, nil, args)
}
func ZendFcallInfoArgv(fci *types.ZendFcallInfo, argc int, argv *va_list) int {
	var i int
	var arg *types.Zval
	if argc < 0 {
		return types.FAILURE
	}
	ZendFcallInfoArgsClear(fci, !argc)
	if argc != 0 {
		fci.SetParamCount(argc)
		fci.SetParams((*types.Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval"))))
		for i = 0; i < argc; i++ {
			arg = __va_arg(*argv, (*types.Zval)(_))
			types.ZVAL_COPY(fci.GetParams()[i], arg)
		}
	}
	return types.SUCCESS
}
func ZendFcallInfoArgn(fci *types.ZendFcallInfo, argc int, _ ...any) int {
	var ret int
	var argv va_list
	va_start(argv, argc)
	ret = ZendFcallInfoArgv(fci, argc, &argv)
	va_end(argv)
	return ret
}
func ZendFcallInfoCall(fci *types.ZendFcallInfo, fcc *types.ZendFcallInfoCache, retval_ptr *types.Zval, args *types.Zval) int {
	var retval types.Zval
	var org_params *types.Zval = nil
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
	if retval_ptr == nil && retval.IsNotUndef() {
		// ZvalPtrDtor(&retval)
	}
	if args != nil {
		ZendFcallInfoArgsRestore(fci, org_count, org_params)
	}
	return result
}
func ZendDeclareTypedProperty(
	ce *types.ClassEntry,
	name *types.String,
	property *types.Zval,
	accessType uint32,
	docComment *string,
	typ types.TypeHint,
) int {
	// calc prop name
	var propName string
	if accessType&types.AccPrivate != 0 {
		propName = ZendManglePropertyName_Ex(ce.Name(), name.GetStr())
	} else if accessType&types.AccProtected != 0 {
		propName = ZendManglePropertyName_Ex("*", name.GetStr())
	} else { // public
		//b.Assert(accessType&types.AccPublic != 0 || accessType&types.AccPppMask == 0)
		propName = name.GetStr()
	}

	//
	var propInfo = types.NewPropertyInfo(0, accessType, propName, docComment, ce, typ)
	var propInfoPtr *types.PropertyInfo
	var propOffset uint32 = 0

	if typ.IsSet() {
		ce.SetIsHasTypeHints(true)
	}
	if ce.IsUserClass() && property.IsConstantAst() {
		ce.SetIsConstantsUpdated(false)
	}
	if propInfo.IsStatic() {
		propInfoPtr = ce.PropertyTable().Get(name.GetStr())
		if propInfoPtr != nil && propInfoPtr.IsStatic() {
			propOffset = propInfoPtr.GetOffset()
			ce.PropertyTable().Del(name.GetStr())
		} else {
			ce.GetDefaultStaticMembersCount()++
			propOffset = ce.GetDefaultStaticMembersCount() - 1
			ce.SetDefaultStaticMembersTable(Perealloc(ce.GetDefaultStaticMembersTable(), b.SizeOf("zval")*ce.GetDefaultStaticMembersCount()))
		}
		types.ZVAL_COPY_VALUE(ce.GetDefaultStaticMembersTable()[propOffset], property)
		if ce.GetStaticMembersTablePtr() == nil {
			b.Assert(ce.IsInternalClass())
			if CurrEX() == nil {
				ZEND_MAP_PTR_NEW(ce.static_members_table)
			} else {
				/* internal class loaded by dl() */
				ZEND_MAP_PTR_INIT(ce.static_members_table, ce.GetDefaultStaticMembersTable())
			}
		}
	} else {
		var propertyDefaultPtr *types.Zval
		propInfoPtr = ce.PropertyTable().Get(name.GetStr())
		if propInfoPtr != nil && !propInfoPtr.IsStatic() {
			propOffset = propInfoPtr.GetOffset()
			ce.PropertyTable().Del(name.GetStr())
			b.Assert(ce.IsInternalClass())
			b.Assert(ce.GetPropertiesInfoTable() != nil)

			ce.GetPropertiesInfoTable()[OBJ_PROP_TO_NUM(propOffset)] = propInfo
			propertyDefaultPtr = ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(propOffset)]
			propertyDefaultPtr.CopyValueFrom(property)

			if property.IsUndef() {
				propertyDefaultPtr.SetU2Extra(types.IS_PROP_UNINIT)
			} else {
				propertyDefaultPtr.SetU2Extra(0)
			}
		} else {
			propOffset = OBJ_PROP_TO_OFFSET(ce.GetDefaultPropertiesCount())

			ce.AddDefaultProperty(property)

			/* For user classes this is handled during linking */
			if ce.IsInternalClass() {
				ce.AddPropertiesInfo(propInfo)
			}
		}
	}
	if ce.IsInternalClass() {
		switch property.GetType() {
		case types.IsArray, types.IsObject, types.IsResource:
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Internal zval's can't be arrays, objects or resources")
		}
	}

	propInfo.SetOffset(propOffset)
	ce.PropertyTable().Update(name.GetStr(), propInfo)
	return types.SUCCESS
}
func ZendTryAssignTypedRefEx(ref *types.ZendReference, val *types.Zval, strict bool) int {
	if ZendVerifyRefAssignableZval(ref, val, strict) == 0 {
		// ZvalPtrDtor(val)
		return types.FAILURE
	} else {
		// ZvalPtrDtor(ref.GetVal())
		types.ZVAL_COPY_VALUE(ref.GetVal(), val)
		return types.SUCCESS
	}
}
func ZendTryAssignTypedRef(ref *types.ZendReference, val *types.Zval) int {
	return ZendTryAssignTypedRefEx(ref, val, CurrEX().IsArgUseStrictTypes())
}
