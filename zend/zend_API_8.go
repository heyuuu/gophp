// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendIsCallableImpl(callable *Zval, object *ZendObject, check_flags uint32, fcc *ZendFcallInfoCache, error **byte) ZendBool {
	var ret ZendBool
	var fcc_local ZendFcallInfoCache
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
	case IS_STRING:
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
	case IS_ARRAY:
		var method *Zval = nil
		var obj *Zval = nil
		if Z_ARRVAL_P(callable).GetNNumOfElements() == 2 {
			obj = callable.GetArr().IndexFindH(0)
			method = callable.GetArr().IndexFindH(1)
		}
		for {
			if obj == nil || method == nil {
				break
			}
			ZVAL_DEREF(method)
			if method.GetType() != IS_STRING {
				break
			}
			ZVAL_DEREF(obj)
			if obj.IsString() {
				if (check_flags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
					return 1
				}
				if ZendIsCallableCheckClass(obj.GetStr(), ZendGetExecutedScope(), fcc, &strict_class, error) == 0 {
					return 0
				}
			} else if obj.IsObject() {
				fcc.SetCallingScope(Z_OBJCE_P(obj))
				fcc.SetObject(obj.GetObj())
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
		if Z_ARRVAL_P(callable).GetNNumOfElements() == 2 {
			if obj == nil || b.CondF(!(obj.IsReference()), func() bool { return obj.GetType() != IS_STRING && obj.GetType() != IS_OBJECT }, func() bool { return Z_REFVAL_P(obj).GetType() != IS_STRING && Z_REFVAL_P(obj).GetType() != IS_OBJECT }) {
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
	case IS_OBJECT:
		if Z_OBJ_HT(*callable).GetGetClosure() != nil {
			if Z_OBJ_HT(*callable).GetGetClosure()(callable, fcc.GetCallingScope(), fcc.GetFunctionHandler(), fcc.GetObject()) == SUCCESS {
				fcc.SetCalledScope(fcc.GetCallingScope())
				if fcc == &fcc_local {
					ZendReleaseFcallInfoCache(fcc)
				}
				return 1
			} else {

				/* Discard exceptions thrown from Z_OBJ_HANDLER_P(callable, get_closure) */

				ZendClearException()

				/* Discard exceptions thrown from Z_OBJ_HANDLER_P(callable, get_closure) */

			}
		}
		if error != nil {
			*error = Estrdup("no array or string given")
		}
		return 0
	case IS_REFERENCE:
		callable = Z_REFVAL_P(callable)
		goto again
	default:
		if error != nil {
			*error = Estrdup("no array or string given")
		}
		return 0
	}
}
func ZendIsCallableEx(
	callable *Zval,
	object *ZendObject,
	check_flags uint32,
	callable_name **ZendString,
	fcc *ZendFcallInfoCache,
	error **byte,
) ZendBool {
	var ret ZendBool = ZendIsCallableImpl(callable, object, check_flags, fcc, error)
	if callable_name != nil {
		*callable_name = ZendGetCallableNameEx(callable, object)
	}
	return ret
}
func ZendIsCallable(callable *Zval, check_flags uint32, callable_name **ZendString) ZendBool {
	return ZendIsCallableEx(callable, nil, check_flags, callable_name, nil, nil)
}
func ZendMakeCallable(callable *Zval, callable_name **ZendString) ZendBool {
	var fcc ZendFcallInfoCache
	if ZendIsCallableEx(callable, nil, IS_CALLABLE_STRICT, callable_name, &fcc, nil) != 0 {
		if callable.IsString() && fcc.GetCallingScope() != nil {
			ZvalPtrDtorStr(callable)
			ArrayInit(callable)
			AddNextIndexStr(callable, fcc.GetCallingScope().GetName().Copy())
			AddNextIndexStr(callable, fcc.GetFunctionHandler().GetFunctionName().Copy())
		}
		ZendReleaseFcallInfoCache(&fcc)
		return 1
	}
	return 0
}
func ZendFcallInfoInit(
	callable *Zval,
	check_flags uint32,
	fci *ZendFcallInfo,
	fcc *ZendFcallInfoCache,
	callable_name **ZendString,
	error **byte,
) int {
	if ZendIsCallableEx(callable, nil, check_flags, callable_name, fcc, error) == 0 {
		return FAILURE
	}
	fci.SetSize(b.SizeOf("* fci"))
	fci.SetObject(fcc.GetObject())
	ZVAL_COPY_VALUE(fci.GetFunctionName(), callable)
	fci.SetRetval(nil)
	fci.SetParamCount(0)
	fci.SetParams(nil)
	fci.SetNoSeparation(1)
	return SUCCESS
}
func ZendFcallInfoArgsClear(fci *ZendFcallInfo, free_mem int) {
	if fci.GetParams() != nil {
		var p *Zval = fci.GetParams()
		var end *Zval = p + fci.GetParamCount()
		for p != end {
			IZvalPtrDtor(p)
			p++
		}
		if free_mem != 0 {
			Efree(fci.GetParams())
			fci.SetParams(nil)
		}
	}
	fci.SetParamCount(0)
}
func ZendFcallInfoArgsSave(fci *ZendFcallInfo, param_count *int, params **Zval) {
	*param_count = fci.GetParamCount()
	*params = fci.GetParams()
	fci.SetParamCount(0)
	fci.SetParams(nil)
}
func ZendFcallInfoArgsRestore(fci *ZendFcallInfo, param_count int, params *Zval) {
	ZendFcallInfoArgsClear(fci, 1)
	fci.SetParamCount(param_count)
	fci.SetParams(params)
}
func ZendFcallInfoArgsEx(fci *ZendFcallInfo, func_ *ZendFunction, args *Zval) int {
	var arg *Zval
	var params *Zval
	var n uint32 = 1
	ZendFcallInfoArgsClear(fci, !args)
	if args == nil {
		return SUCCESS
	}
	if args.GetType() != IS_ARRAY {
		return FAILURE
	}
	fci.SetParamCount(Z_ARRVAL_P(args).GetNNumOfElements())
	params = (*Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval")))
	fci.SetParams(params)
	var __ht *HashTable = args.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		arg = _z
		if func_ != nil && !(arg.IsReference()) && ARG_SHOULD_BE_SENT_BY_REF(func_, n) != 0 {
			params.SetNewRef(arg)
			arg.TryAddRefcount()
		} else {
			ZVAL_COPY(params, arg)
		}
		params++
		n++
	}
	return SUCCESS
}
func ZendFcallInfoArgs(fci *ZendFcallInfo, args *Zval) int {
	return ZendFcallInfoArgsEx(fci, nil, args)
}
func ZendFcallInfoArgp(fci *ZendFcallInfo, argc int, argv *Zval) int {
	var i int
	if argc < 0 {
		return FAILURE
	}
	ZendFcallInfoArgsClear(fci, !argc)
	if argc != 0 {
		fci.SetParamCount(argc)
		fci.SetParams((*Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval"))))
		for i = 0; i < argc; i++ {
			ZVAL_COPY(fci.GetParams()[i], &argv[i])
		}
	}
	return SUCCESS
}
func ZendFcallInfoArgv(fci *ZendFcallInfo, argc int, argv *va_list) int {
	var i int
	var arg *Zval
	if argc < 0 {
		return FAILURE
	}
	ZendFcallInfoArgsClear(fci, !argc)
	if argc != 0 {
		fci.SetParamCount(argc)
		fci.SetParams((*Zval)(Erealloc(fci.GetParams(), fci.GetParamCount()*b.SizeOf("zval"))))
		for i = 0; i < argc; i++ {
			arg = __va_arg(*argv, (*Zval)(_))
			ZVAL_COPY(fci.GetParams()[i], arg)
		}
	}
	return SUCCESS
}
func ZendFcallInfoArgn(fci *ZendFcallInfo, argc int, _ ...any) int {
	var ret int
	var argv va_list
	va_start(argv, argc)
	ret = ZendFcallInfoArgv(fci, argc, &argv)
	va_end(argv)
	return ret
}
func ZendFcallInfoCall(fci *ZendFcallInfo, fcc *ZendFcallInfoCache, retval_ptr *Zval, args *Zval) int {
	var retval Zval
	var org_params *Zval = nil
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
	if retval_ptr == nil && retval.GetType() != IS_UNDEF {
		ZvalPtrDtor(&retval)
	}
	if args != nil {
		ZendFcallInfoArgsRestore(fci, org_count, org_params)
	}
	return result
}
func ZendGetModuleVersion(module_name *byte) *byte {
	var lname *ZendString
	var name_len int = strlen(module_name)
	var module *ZendModuleEntry
	lname = ZendStringAlloc(name_len, 0)
	ZendStrTolowerCopy(lname.GetVal(), module_name, name_len)
	module = ZendHashFindPtr(&ModuleRegistry, lname)
	ZendStringEfree(lname)
	if module != nil {
		return module.GetVersion()
	} else {
		return nil
	}
}
func ZvalMakeInternedString(zv *Zval) *ZendString {
	ZEND_ASSERT(zv.IsString())
	zv.SetStr(ZendNewInternedString(zv.GetStr()))

	return zv.GetStr()
}
func IsPersistentClass(ce *ZendClassEntry) ZendBool {
	return (ce.GetType()&ZEND_INTERNAL_CLASS) != 0 && ce.GetModule().GetType() == MODULE_PERSISTENT
}
func ZendDeclareTypedProperty(
	ce *ZendClassEntry,
	name *ZendString,
	property *Zval,
	access_type int,
	doc_comment *ZendString,
	type_ ZendType,
) int {
	var property_info *ZendPropertyInfo
	var property_info_ptr *ZendPropertyInfo
	if type_.IsSet() {
		ce.SetIsHasTypeHints(true)
	}
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		property_info = Pemalloc(b.SizeOf("zend_property_info"), 1)
	} else {
		property_info = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_property_info"))
		if property.IsConstant() {
			ce.SetIsConstantsUpdated(false)
		}
	}
	if property.IsString() {
		ZvalMakeInternedString(property)
	}
	if (access_type & ZEND_ACC_PPP_MASK) == 0 {
		access_type |= ZEND_ACC_PUBLIC
	}
	if (access_type & ZEND_ACC_STATIC) != 0 {
		if b.Assign(&property_info_ptr, ZendHashFindPtr(ce.GetPropertiesInfo(), name)) != nil && property_info_ptr.IsStatic() {
			property_info.SetOffset(property_info_ptr.GetOffset())
			ZvalPtrDtor(ce.GetDefaultStaticMembersTable()[property_info.GetOffset()])
			ZendHashDel(ce.GetPropertiesInfo(), name)
		} else {
			ce.GetDefaultStaticMembersCount()++
			property_info.SetOffset(ce.GetDefaultStaticMembersCount() - 1)
			ce.SetDefaultStaticMembersTable(Perealloc(ce.GetDefaultStaticMembersTable(), b.SizeOf("zval")*ce.GetDefaultStaticMembersCount(), ce.GetType() == ZEND_INTERNAL_CLASS))
		}
		ZVAL_COPY_VALUE(ce.GetDefaultStaticMembersTable()[property_info.GetOffset()], property)
		if ce.GetStaticMembersTablePtr() == nil {
			ZEND_ASSERT(ce.GetType() == ZEND_INTERNAL_CLASS)
			if EG__().GetCurrentExecuteData() == nil {
				ZEND_MAP_PTR_NEW(ce.static_members_table)
			} else {

				/* internal class loaded by dl() */

				ZEND_MAP_PTR_INIT(ce.static_members_table, ce.GetDefaultStaticMembersTable())

				/* internal class loaded by dl() */

			}
		}
	} else {
		var property_default_ptr *Zval
		if b.Assign(&property_info_ptr, ZendHashFindPtr(ce.GetPropertiesInfo(), name)) != nil && !property_info_ptr.IsStatic() {
			property_info.SetOffset(property_info_ptr.GetOffset())
			ZvalPtrDtor(ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())])
			ZendHashDel(ce.GetPropertiesInfo(), name)
			ZEND_ASSERT(ce.GetType() == ZEND_INTERNAL_CLASS)
			ZEND_ASSERT(ce.GetPropertiesInfoTable() != nil)
			ce.GetPropertiesInfoTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())] = property_info
		} else {
			property_info.SetOffset(OBJ_PROP_TO_OFFSET(ce.GetDefaultPropertiesCount()))
			ce.GetDefaultPropertiesCount()++
			ce.SetDefaultPropertiesTable(Perealloc(ce.GetDefaultPropertiesTable(), b.SizeOf("zval")*ce.GetDefaultPropertiesCount(), ce.GetType() == ZEND_INTERNAL_CLASS))

			/* For user classes this is handled during linking */

			if ce.GetType() == ZEND_INTERNAL_CLASS {
				ce.SetPropertiesInfoTable(Perealloc(ce.GetPropertiesInfoTable(), b.SizeOf("zend_property_info *")*ce.GetDefaultPropertiesCount(), 1))
				ce.GetPropertiesInfoTable()[ce.GetDefaultPropertiesCount()-1] = property_info
			}

			/* For user classes this is handled during linking */

		}
		property_default_ptr = ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())]
		ZVAL_COPY_VALUE(property_default_ptr, property)
		if property.IsUndef() {
			property_default_ptr.SetU2Extra(IS_PROP_UNINIT)
		} else {
			property_default_ptr.SetU2Extra(0)
		}
	}
	if (ce.GetType() & ZEND_INTERNAL_CLASS) != 0 {
		switch property.GetType() {
		case IS_ARRAY:

		case IS_OBJECT:

		case IS_RESOURCE:
			ZendErrorNoreturn(E_CORE_ERROR, "Internal zval's can't be arrays, objects or resources")
			break
		default:
			break
		}

		/* Must be interned to avoid ZTS data races */

		if IsPersistentClass(ce) != 0 {
			name = ZendNewInternedString(name.Copy())
		}

		/* Must be interned to avoid ZTS data races */

	}
	if (access_type & ZEND_ACC_PUBLIC) != 0 {
		property_info.SetName(name.Copy())
	} else if (access_type & ZEND_ACC_PRIVATE) != 0 {
		property_info.SetName(ZendManglePropertyName_ZStr(ce.GetName().GetStr(), name.GetStr(), IsPersistentClass(ce) != 0))
	} else {
		ZEND_ASSERT((access_type & ZEND_ACC_PROTECTED) != 0)
		property_info.SetName(ZendManglePropertyName_ZStr("*", name.GetStr(), IsPersistentClass(ce) != 0))
	}
	property_info.SetName(ZendNewInternedString(property_info.GetName()))
	property_info.SetFlags(access_type)
	property_info.SetDocComment(doc_comment)
	property_info.SetCe(ce)
	property_info.SetType(type_)
	ZendHashUpdatePtr(ce.GetPropertiesInfo(), name, property_info)
	return SUCCESS
}
func ZendTryAssignTypedRefEx(ref *ZendReference, val *Zval, strict ZendBool) int {
	if ZendVerifyRefAssignableZval(ref, val, strict) == 0 {
		ZvalPtrDtor(val)
		return FAILURE
	} else {
		ZvalPtrDtor(ref.GetVal())
		ZVAL_COPY_VALUE(ref.GetVal(), val)
		return SUCCESS
	}
}
func ZendTryAssignTypedRef(ref *ZendReference, val *Zval) int {
	return ZendTryAssignTypedRefEx(ref, val, ZEND_ARG_USES_STRICT_TYPES())
}
