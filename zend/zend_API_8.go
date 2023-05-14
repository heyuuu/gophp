package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZendIsCallableImpl(callable *types.Zval, object *types.ZendObject, check_flags uint32, fcc *types.ZendFcallInfoCache, error **byte) types.ZendBool {
	var ret types.ZendBool
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
	case types.IS_STRING:
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
	case types.IS_ARRAY:
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
			if method.GetType() != types.IS_STRING {
				break
			}
			obj = types.ZVAL_DEREF(obj)
			if obj.IsString() {
				if (check_flags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
					return 1
				}
				if ZendIsCallableCheckClass(obj.String(), ZendGetExecutedScope(), fcc, &strict_class, error) == 0 {
					return 0
				}
			} else if obj.IsObject() {
				fcc.SetCallingScope(types.Z_OBJCE_P(obj))
				fcc.SetObject(obj.Object())
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
		if callable.Array().Len() == 2 {
			if obj == nil || b.CondF(!(obj.IsReference()), func() bool { return obj.GetType() != types.IS_STRING && obj.GetType() != types.IS_OBJECT }, func() bool {
				return types.Z_REFVAL_P(obj).GetType() != types.IS_STRING && types.Z_REFVAL_P(obj).GetType() != types.IS_OBJECT
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
	case types.IS_OBJECT:
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
	case types.IS_REFERENCE:
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
	object *types.ZendObject,
	check_flags uint32,
	callable_name **types.String,
	fcc *types.ZendFcallInfoCache,
	error **byte,
) types.ZendBool {
	var ret types.ZendBool = ZendIsCallableImpl(callable, object, check_flags, fcc, error)
	if callable_name != nil {
		*callable_name = ZendGetCallableNameEx(callable, object)
	}
	return ret
}
func ZendIsCallable(callable *types.Zval, check_flags uint32, callable_name **types.String) types.ZendBool {
	return ZendIsCallableEx(callable, nil, check_flags, callable_name, nil, nil)
}
func ZendFcallInfoInit(
	callable *types.Zval,
	check_flags uint32,
	fci *types.ZendFcallInfo,
	fcc *types.ZendFcallInfoCache,
	callable_name **types.String,
	error **byte,
) int {
	if ZendIsCallableEx(callable, nil, check_flags, callable_name, fcc, error) == 0 {
		return types.FAILURE
	}
	fci.SetSize(b.SizeOf("* fci"))
	fci.SetObject(fcc.GetObject())
	types.ZVAL_COPY_VALUE(fci.GetFunctionName(), callable)
	fci.SetRetval(nil)
	fci.SetParamCount(0)
	fci.SetParams(nil)
	fci.SetNoSeparation(1)
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
	if args.GetType() != types.IS_ARRAY {
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
func ZendFcallInfoArgp(fci *types.ZendFcallInfo, argc int, argv *types.Zval) int {
	var i int
	if argc < 0 {
		return types.FAILURE
	}
	ZendFcallInfoArgsClear(fci, !argc)
	if argc != 0 {
		fci.SetParamCount(argc)
		fci.SetParams((*types.Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval"))))
		for i = 0; i < argc; i++ {
			types.ZVAL_COPY(fci.GetParams()[i], &argv[i])
		}
	}
	return types.SUCCESS
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
		propName = ZendManglePropertyName_Ex(ce.GetName().GetStr(), name.GetStr())
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
		var property_default_ptr *types.Zval
		propInfoPtr = ce.PropertyTable().Get(name.GetStr())
		if propInfoPtr != nil && !propInfoPtr.IsStatic() {
			propOffset = propInfoPtr.GetOffset()
			// ZvalPtrDtor(ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())])
			ce.PropertyTable().Del(name.GetStr())
			b.Assert(ce.IsInternalClass())
			b.Assert(ce.GetPropertiesInfoTable() != nil)
			ce.GetPropertiesInfoTable()[OBJ_PROP_TO_NUM(propOffset)] = propInfo
		} else {
			propOffset = OBJ_PROP_TO_OFFSET(ce.GetDefaultPropertiesCount())
			ce.GetDefaultPropertiesCount()++
			ce.SetDefaultPropertiesTable(Perealloc(ce.GetDefaultPropertiesTable(), b.SizeOf("zval")*ce.GetDefaultPropertiesCount()))

			/* For user classes this is handled during linking */

			if ce.IsInternalClass() {
				ce.SetPropertiesInfoTable(Perealloc(ce.GetPropertiesInfoTable(), b.SizeOf("zend_property_info *")*ce.GetDefaultPropertiesCount()))
				ce.GetPropertiesInfoTable()[ce.GetDefaultPropertiesCount()-1] = propInfo
			}
		}
		property_default_ptr = ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(propOffset)]
		property_default_ptr.CopyValueFrom(property)
		if property.IsUndef() {
			property_default_ptr.SetU2Extra(types.IS_PROP_UNINIT)
		} else {
			property_default_ptr.SetU2Extra(0)
		}
	}
	if ce.IsInternalClass() {
		switch property.GetType() {
		case types.IS_ARRAY, types.IS_OBJECT, types.IS_RESOURCE:
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Internal zval's can't be arrays, objects or resources")
		}
	}

	propInfo.SetOffset(propOffset)
	ce.PropertyTable().Update(name.GetStr(), propInfo)
	return types.SUCCESS
}
func ZendTryAssignTypedRefEx(ref *types.ZendReference, val *types.Zval, strict types.ZendBool) int {
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
