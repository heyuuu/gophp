package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

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
		if func_ != nil && !(arg.IsRef()) && ARG_SHOULD_BE_SENT_BY_REF(func_, n) != 0 {
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
func ZendFcallInfoCall(fci *types.ZendFcallInfo, fcc *types.ZendFcallInfoCache, retval_ptr *types.Zval, args *types.Zval) bool {
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
	return result == types.SUCCESS
}
func ZendDeclareTypedProperty(ce *types.ClassEntry, name string, property *types.Zval, accessType uint32, docComment string, typ *types.TypeHint) int {
	// calc prop name
	var propName string
	if accessType&types.AccPrivate != 0 {
		propName = ZendManglePropertyName_Ex(ce.Name(), name)
	} else if accessType&types.AccProtected != 0 {
		propName = ZendManglePropertyName_Ex("*", name)
	} else { // public
		b.Assert(accessType&types.AccPublic != 0 || accessType&types.AccPppMask == 0)
		propName = name
	}

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
		propInfoPtr = ce.PropertyTable().Get(name)
		if propInfoPtr != nil && propInfoPtr.IsStatic() {
			propOffset = propInfoPtr.GetOffset()
			ce.PropertyTable().Del(name)
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
		propInfoPtr = ce.PropertyTable().Get(name)
		if propInfoPtr != nil && !propInfoPtr.IsStatic() {
			propOffset = propInfoPtr.GetOffset()
			ce.PropertyTable().Del(name)
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
		switch property.Type() {
		case types.IsArray, types.IsObject, types.IsResource:
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Internal zval's can't be arrays, objects or resources")
		}
	}

	propInfo.SetOffset(propOffset)
	ce.PropertyTable().Update(name, propInfo)
	return types.SUCCESS
}
func ZendTryAssignTypedRefEx(ref *types.Reference, val *types.Zval, strict bool) int {
	if !ZendVerifyRefAssignableZval(ref, val, strict) {
		return types.FAILURE
	} else {
		types.ZVAL_COPY_VALUE(ref.GetVal(), val)
		return types.SUCCESS
	}
}
func ZendTryAssignTypedRef(ref *types.Reference, val *types.Zval) int {
	return ZendTryAssignTypedRefEx(ref, val, CurrEX().IsArgUseStrictTypes())
}
