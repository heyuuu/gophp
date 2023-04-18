package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZendIsCallableImpl(callable *types2.Zval, object *types2.ZendObject, check_flags uint32, fcc *types2.ZendFcallInfoCache, error **byte) types2.ZendBool {
	var ret types2.ZendBool
	var fcc_local types2.ZendFcallInfoCache
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
	case types2.IS_STRING:
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
	case types2.IS_ARRAY:
		var method *types2.Zval = nil
		var obj *types2.Zval = nil
		if callable.Array().Len() == 2 {
			obj = callable.Array().IndexFind(0)
			method = callable.Array().IndexFind(1)
		}
		for {
			if obj == nil || method == nil {
				break
			}
			method = types2.ZVAL_DEREF(method)
			if method.GetType() != types2.IS_STRING {
				break
			}
			obj = types2.ZVAL_DEREF(obj)
			if obj.IsString() {
				if (check_flags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
					return 1
				}
				if ZendIsCallableCheckClass(obj.String(), ZendGetExecutedScope(), fcc, &strict_class, error) == 0 {
					return 0
				}
			} else if obj.IsObject() {
				fcc.SetCallingScope(types2.Z_OBJCE_P(obj))
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
			if obj == nil || b.CondF(!(obj.IsReference()), func() bool { return obj.GetType() != types2.IS_STRING && obj.GetType() != types2.IS_OBJECT }, func() bool {
				return types2.Z_REFVAL_P(obj).GetType() != types2.IS_STRING && types2.Z_REFVAL_P(obj).GetType() != types2.IS_OBJECT
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
	case types2.IS_OBJECT:
		if types2.Z_OBJ_HT(*callable).GetGetClosure() != nil {
			if types2.Z_OBJ_HT(*callable).GetGetClosure()(callable, fcc.GetCallingScope(), fcc.GetFunctionHandler(), fcc.GetObject()) == types2.SUCCESS {
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
	case types2.IS_REFERENCE:
		callable = types2.Z_REFVAL_P(callable)
		goto again
	default:
		if error != nil {
			*error = Estrdup("no array or string given")
		}
		return 0
	}
}
func ZendIsCallableEx(
	callable *types2.Zval,
	object *types2.ZendObject,
	check_flags uint32,
	callable_name **types2.String,
	fcc *types2.ZendFcallInfoCache,
	error **byte,
) types2.ZendBool {
	var ret types2.ZendBool = ZendIsCallableImpl(callable, object, check_flags, fcc, error)
	if callable_name != nil {
		*callable_name = ZendGetCallableNameEx(callable, object)
	}
	return ret
}
func ZendIsCallable(callable *types2.Zval, check_flags uint32, callable_name **types2.String) types2.ZendBool {
	return ZendIsCallableEx(callable, nil, check_flags, callable_name, nil, nil)
}
func ZendFcallInfoInit(
	callable *types2.Zval,
	check_flags uint32,
	fci *types2.ZendFcallInfo,
	fcc *types2.ZendFcallInfoCache,
	callable_name **types2.String,
	error **byte,
) int {
	if ZendIsCallableEx(callable, nil, check_flags, callable_name, fcc, error) == 0 {
		return types2.FAILURE
	}
	fci.SetSize(b.SizeOf("* fci"))
	fci.SetObject(fcc.GetObject())
	types2.ZVAL_COPY_VALUE(fci.GetFunctionName(), callable)
	fci.SetRetval(nil)
	fci.SetParamCount(0)
	fci.SetParams(nil)
	fci.SetNoSeparation(1)
	return types2.SUCCESS
}
func ZendFcallInfoArgsClear(fci *types2.ZendFcallInfo, free_mem int) {
	if fci.GetParams() != nil {
		var p *types2.Zval = fci.GetParams()
		var end *types2.Zval = p + fci.GetParamCount()
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
func ZendFcallInfoArgsSave(fci *types2.ZendFcallInfo, param_count *int, params **types2.Zval) {
	*param_count = fci.GetParamCount()
	*params = fci.GetParams()
	fci.SetParamCount(0)
	fci.SetParams(nil)
}
func ZendFcallInfoArgsRestore(fci *types2.ZendFcallInfo, param_count int, params *types2.Zval) {
	ZendFcallInfoArgsClear(fci, 1)
	fci.SetParamCount(param_count)
	fci.SetParams(params)
}
func ZendFcallInfoArgsEx(fci *types2.ZendFcallInfo, func_ types2.IFunction, args *types2.Zval) int {
	var arg *types2.Zval
	var params *types2.Zval
	var n uint32 = 1
	ZendFcallInfoArgsClear(fci, !args)
	if args == nil {
		return types2.SUCCESS
	}
	if args.GetType() != types2.IS_ARRAY {
		return types2.FAILURE
	}
	fci.SetParamCount(args.Array().Len())
	params = (*types2.Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval")))
	fci.SetParams(params)
	var __ht *types2.Array = args.Array()
	for _, _p := range __ht.ForeachData() {
		var _z *types2.Zval = _p.GetVal()

		arg = _z
		if func_ != nil && !(arg.IsReference()) && ARG_SHOULD_BE_SENT_BY_REF(func_, n) != 0 {
			params.SetNewRef(arg)
			// arg.TryAddRefcount()
		} else {
			types2.ZVAL_COPY(params, arg)
		}
		params++
		n++
	}
	return types2.SUCCESS
}
func ZendFcallInfoArgs(fci *types2.ZendFcallInfo, args *types2.Zval) int {
	return ZendFcallInfoArgsEx(fci, nil, args)
}
func ZendFcallInfoArgp(fci *types2.ZendFcallInfo, argc int, argv *types2.Zval) int {
	var i int
	if argc < 0 {
		return types2.FAILURE
	}
	ZendFcallInfoArgsClear(fci, !argc)
	if argc != 0 {
		fci.SetParamCount(argc)
		fci.SetParams((*types2.Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval"))))
		for i = 0; i < argc; i++ {
			types2.ZVAL_COPY(fci.GetParams()[i], &argv[i])
		}
	}
	return types2.SUCCESS
}
func ZendFcallInfoArgv(fci *types2.ZendFcallInfo, argc int, argv *va_list) int {
	var i int
	var arg *types2.Zval
	if argc < 0 {
		return types2.FAILURE
	}
	ZendFcallInfoArgsClear(fci, !argc)
	if argc != 0 {
		fci.SetParamCount(argc)
		fci.SetParams((*types2.Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval"))))
		for i = 0; i < argc; i++ {
			arg = __va_arg(*argv, (*types2.Zval)(_))
			types2.ZVAL_COPY(fci.GetParams()[i], arg)
		}
	}
	return types2.SUCCESS
}
func ZendFcallInfoArgn(fci *types2.ZendFcallInfo, argc int, _ ...any) int {
	var ret int
	var argv va_list
	va_start(argv, argc)
	ret = ZendFcallInfoArgv(fci, argc, &argv)
	va_end(argv)
	return ret
}
func ZendFcallInfoCall(fci *types2.ZendFcallInfo, fcc *types2.ZendFcallInfoCache, retval_ptr *types2.Zval, args *types2.Zval) int {
	var retval types2.Zval
	var org_params *types2.Zval = nil
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
	ce *types2.ClassEntry,
	name *types2.String,
	property *types2.Zval,
	access_type int,
	doc_comment *types2.String,
	type_ types2.ZendType,
) int {
	var property_info *ZendPropertyInfo
	var property_info_ptr *ZendPropertyInfo
	if type_.IsSet() {
		ce.SetIsHasTypeHints(true)
	}
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		property_info = Pemalloc(b.SizeOf("zend_property_info"))
	} else {
		property_info = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_property_info"))
		if property.IsConstantAst() {
			ce.SetIsConstantsUpdated(false)
		}
	}
	//if property.IsString() {
	//	ZvalMakeInternedString(property)
	//}
	if (access_type & AccPppMask) == 0 {
		access_type |= AccPublic
	}
	if (access_type & AccStatic) != 0 {
		property_info_ptr = ce.PropertyTable().Get(name.GetStr())
		if property_info_ptr != nil && property_info_ptr.IsStatic() {
			property_info.SetOffset(property_info_ptr.GetOffset())
			// ZvalPtrDtor(ce.GetDefaultStaticMembersTable()[property_info.GetOffset()])
			ce.PropertyTable().Del(name.GetStr())
		} else {
			ce.GetDefaultStaticMembersCount()++
			property_info.SetOffset(ce.GetDefaultStaticMembersCount() - 1)
			ce.SetDefaultStaticMembersTable(Perealloc(ce.GetDefaultStaticMembersTable(), b.SizeOf("zval")*ce.GetDefaultStaticMembersCount()))
		}
		types2.ZVAL_COPY_VALUE(ce.GetDefaultStaticMembersTable()[property_info.GetOffset()], property)
		if ce.GetStaticMembersTablePtr() == nil {
			b.Assert(ce.GetType() == ZEND_INTERNAL_CLASS)
			if CurrEX() == nil {
				ZEND_MAP_PTR_NEW(ce.static_members_table)
			} else {

				/* internal class loaded by dl() */

				ZEND_MAP_PTR_INIT(ce.static_members_table, ce.GetDefaultStaticMembersTable())

				/* internal class loaded by dl() */

			}
		}
	} else {
		var property_default_ptr *types2.Zval
		property_info_ptr = ce.PropertyTable().Get(name.GetStr())
		if property_info_ptr != nil && !property_info_ptr.IsStatic() {
			property_info.SetOffset(property_info_ptr.GetOffset())
			// ZvalPtrDtor(ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())])
			ce.PropertyTable().Del(name.GetStr())
			b.Assert(ce.GetType() == ZEND_INTERNAL_CLASS)
			b.Assert(ce.GetPropertiesInfoTable() != nil)
			ce.GetPropertiesInfoTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())] = property_info
		} else {
			property_info.SetOffset(OBJ_PROP_TO_OFFSET(ce.GetDefaultPropertiesCount()))
			ce.GetDefaultPropertiesCount()++
			ce.SetDefaultPropertiesTable(Perealloc(ce.GetDefaultPropertiesTable(), b.SizeOf("zval")*ce.GetDefaultPropertiesCount()))

			/* For user classes this is handled during linking */

			if ce.GetType() == ZEND_INTERNAL_CLASS {
				ce.SetPropertiesInfoTable(Perealloc(ce.GetPropertiesInfoTable(), b.SizeOf("zend_property_info *")*ce.GetDefaultPropertiesCount()))
				ce.GetPropertiesInfoTable()[ce.GetDefaultPropertiesCount()-1] = property_info
			}

			/* For user classes this is handled during linking */

		}
		property_default_ptr = ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())]
		property_default_ptr.CopyValueFrom(property)
		if property.IsUndef() {
			property_default_ptr.SetU2Extra(types2.IS_PROP_UNINIT)
		} else {
			property_default_ptr.SetU2Extra(0)
		}
	}
	if (ce.GetType() & ZEND_INTERNAL_CLASS) != 0 {
		switch property.GetType() {
		case types2.IS_ARRAY:

		case types2.IS_OBJECT:

		case types2.IS_RESOURCE:
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Internal zval's can't be arrays, objects or resources")
			break
		default:
			break
		}

		/* Must be interned to avoid ZTS data races */

		//if IsPersistentClass(ce) != 0 {
		//	name = types.ZendNewInternedString(name.Copy())
		//}

		/* Must be interned to avoid ZTS data races */

	}
	if (access_type & AccPublic) != 0 {
		property_info.SetName(name.Copy())
	} else if (access_type & AccPrivate) != 0 {
		property_info.SetName(ZendManglePropertyName_ZStr(ce.GetName().GetStr(), name.GetStr()))
	} else {
		b.Assert((access_type & AccProtected) != 0)
		property_info.SetName(ZendManglePropertyName_ZStr("*", name.GetStr()))
	}
	//property_info.SetName(types.ZendNewInternedString(property_info.GetName()))
	property_info.SetFlags(access_type)
	property_info.SetDocComment(doc_comment)
	property_info.SetCe(ce)
	property_info.SetType(type_)
	ce.PropertyTable().Update(name.GetStr(), property_info)
	return types2.SUCCESS
}
func ZendTryAssignTypedRefEx(ref *types2.ZendReference, val *types2.Zval, strict types2.ZendBool) int {
	if ZendVerifyRefAssignableZval(ref, val, strict) == 0 {
		// ZvalPtrDtor(val)
		return types2.FAILURE
	} else {
		// ZvalPtrDtor(ref.GetVal())
		types2.ZVAL_COPY_VALUE(ref.GetVal(), val)
		return types2.SUCCESS
	}
}
func ZendTryAssignTypedRef(ref *types2.ZendReference, val *types2.Zval) int {
	return ZendTryAssignTypedRefEx(ref, val, CurrEX().IsArgUseStrictTypes())
}
