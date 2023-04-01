package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"strings"
)

func ZendResolveFunctionName(name *types.String, type_ uint32, is_fully_qualified *types.ZendBool) *types.String {
	return ZendResolveNonClassName(name, type_, is_fully_qualified, 0, FC__().GetImportsFunction())
}
func ZendResolveConstName(name *types.String, type_ uint32, is_fully_qualified *types.ZendBool) *types.String {
	return ZendResolveNonClassName(name, type_, is_fully_qualified, 1, FC__().GetImportsConst())
}
func ZendResolveClassName(name *types.String, type_ uint32) *types.String {
	var compound *byte
	if type_ == ZEND_NAME_RELATIVE {
		return ZendPrefixWithNs(name)
	}
	if type_ == ZEND_NAME_FQ || name.GetVal()[0] == '\\' {

		/* Remove \ prefix (only relevant if this is a string rather than a label) */

		if name.GetVal()[0] == '\\' {
			name = types.NewString(b.CastStr(name.GetVal()+1, name.GetLen()-1))
		} else {
			name.AddRefcount()
		}

		/* Ensure that \self, \parent and \static are not used */

		if ZEND_FETCH_CLASS_DEFAULT != ZendGetClassFetchType(name) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'\\%s' is an invalid class name", name.GetVal())
		}
		return name
	}
	if FC__().GetImports() != nil {
		compound = memchr(name.GetVal(), '\\', name.GetLen())
		if compound != nil {

			/* If the first part of a qualified name is an alias, substitute it. */

			var len_ int = compound - name.GetVal()
			var import_name *types.String = ZendHashFindPtrLc(FC__().GetImports(), name.GetVal(), len_)
			if import_name != nil {
				return ZendConcatNames(import_name.GetVal(), import_name.GetLen(), name.GetVal()+len_+1, name.GetLen()-len_-1)
			}
		} else {

			/* If an unqualified name is an alias, replace it. */

			var import_name *types.String = ZendHashFindPtrLc(FC__().GetImports(), name.GetVal(), name.GetLen())
			if import_name != nil {
				return import_name.Copy()
			}
		}
	}

	/* If not fully qualified and not an alias, prepend the current namespace */

	return ZendPrefixWithNs(name)

	/* If not fully qualified and not an alias, prepend the current namespace */
}
func ZendResolveClassNameAst(ast *ZendAst) *types.String {
	var class_name = ZendAstGetZval(ast)
	if class_name.GetType() != types.IS_STRING {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Illegal class name")
	}
	return ZendResolveClassName(class_name.GetStr(), ast.GetAttr())
}
func LabelPtrDtor(zv *types.Zval) {
	EfreeSize(zv.GetPtr(), b.SizeOf("zend_label"))
}
func StrDtor(zv *types.Zval) { types.ZendStringReleaseEx(zv.GetStr(), 0) }
func ZendAddTryElement(try_op uint32) uint32 {
	var op_array = CG__().GetActiveOpArray()
	var try_catch_offset uint32 = b.PostInc(&(op_array.GetLastTryCatch()))
	var elem *ZendTryCatchElement
	op_array.SetTryCatchArray(SafeErealloc(op_array.GetTryCatchArray(), b.SizeOf("zend_try_catch_element"), op_array.GetLastTryCatch(), 0))
	elem = op_array.GetTryCatchArray()[try_catch_offset]
	elem.SetTryOp(try_op)
	elem.SetCatchOp(0)
	elem.SetFinallyOp(0)
	elem.SetFinallyEnd(0)
	return try_catch_offset
}
func FunctionAddRef(function *types.ZendFunction) {
	if function.GetType() == ZEND_USER_FUNCTION {
		var op_array = function.GetOpArray()
		if op_array.GetRefcount() != nil {
			op_array.refcount++
		}
		if op_array.GetStaticVariables() != nil {
			if (op_array.GetStaticVariables().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
				op_array.GetStaticVariables().AddRefcount()
			}
		}
		if (CG__().GetCompilerOptions() & ZEND_COMPILE_PRELOAD) != 0 {
			b.Assert(op_array.IsPreloaded())
			ZEND_MAP_PTR_NEW(op_array.run_time_cache)
			ZEND_MAP_PTR_NEW(op_array.static_variables_ptr)
		} else {
			ZEND_MAP_PTR_INIT(op_array.static_variables_ptr, op_array.GetStaticVariables())
			ZEND_MAP_PTR_INIT(op_array.run_time_cache, ZendArenaAlloc(CG__().GetArena(), b.SizeOf("void *")))
			ZEND_MAP_PTR_SET(op_array.run_time_cache, nil)
		}
	} else if function.GetType() == ZEND_INTERNAL_FUNCTION {
		if function.GetFunctionName() != nil {
			function.GetFunctionName().AddRefcount()
		}
	}
}
func DoBindFunctionError(lcname *types.String, op_array *types.ZendOpArray, compile_time types.ZendBool) {
	var zv = b.CondF(compile_time != 0, func() *types.Array { return CG__().GetFunctionTable() }, func() *types.Array { return EG__().GetFunctionTable() }).KeyFind(lcname.GetStr())
	var error_level = b.Cond(compile_time != 0, faults.E_COMPILE_ERROR, faults.E_ERROR)
	var old_function *types.ZendFunction
	b.Assert(zv != nil)
	old_function = (*types.ZendFunction)(zv.GetPtr())
	if old_function.GetType() == ZEND_USER_FUNCTION && old_function.GetOpArray().GetLast() > 0 {
		faults.ErrorNoreturn(error_level, "Cannot redeclare %s() (previously declared in %s:%d)", b.CondF(op_array != nil, func() []byte { return op_array.GetFunctionName().GetVal() }, func() []byte { return old_function.GetFunctionName().GetVal() }), old_function.GetOpArray().GetFilename().GetVal(), old_function.GetOpArray().GetOpcodes()[0].GetLineno())
	} else {
		faults.ErrorNoreturn(error_level, "Cannot redeclare %s()", b.CondF(op_array != nil, func() []byte { return op_array.GetFunctionName().GetVal() }, func() []byte { return old_function.GetFunctionName().GetVal() }))
	}
}
func DoBindFunction(lcname *types.Zval) int {
	var function *types.ZendFunction
	var rtd_key *types.Zval
	var zv *types.Zval
	rtd_key = lcname + 1
	zv = EG__().GetFunctionTable().KeyFind(rtd_key.GetStr().GetStr())
	if zv == nil {
		DoBindFunctionError(lcname.GetStr(), nil, 0)
		return types.FAILURE
	}
	function = (*types.ZendFunction)(zv.GetPtr())
	if function.IsPreloaded() && (CG__().GetCompilerOptions()&ZEND_COMPILE_PRELOAD) == 0 {
		zv = EG__().GetFunctionTable().KeyAdd(lcname.GetStr().GetStr(), zv)
	} else {
		zv = types.ZendHashSetBucketKey(EG__().GetFunctionTable(), (*types.Bucket)(zv), lcname.GetStr().GetStr())
	}
	if zv == nil {
		DoBindFunctionError(lcname.GetStr(), function.GetOpArray(), 0)
		return types.FAILURE
	}
	return types.SUCCESS
}
func DoBindClass(lcname *types.Zval, lc_parent_name *types.String) int {
	var ce *types.ClassEntry
	var rtd_key *types.Zval
	var zv *types.Zval
	rtd_key = lcname + 1
	zv = EG__().GetClassTable().KeyFind(rtd_key.GetStr().GetStr())
	if zv == nil {
		ce = types.ZendHashFindPtr(EG__().GetClassTable(), lcname.GetStr().GetStr())
		if ce != nil {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.GetName().GetVal())
			return types.FAILURE
		} else {
			for {
				b.Assert(CurrEX().GetFunc().GetOpArray().IsPreloaded())
				if ZendPreloadAutoload != nil && ZendPreloadAutoload(CurrEX().GetFunc().GetOpArray().GetFilename()) == types.SUCCESS {
					zv = EG__().GetClassTable().KeyFind(rtd_key.GetStr().GetStr())
					if zv != nil {
						break
					}
				}
				faults.ErrorNoreturn(faults.E_ERROR, "Class %s wasn't preloaded", lcname.GetStr().GetVal())
				return types.FAILURE
				break
			}
		}
	}

	/* Register the derived class */

	ce = (*types.ClassEntry)(zv.GetPtr())
	zv = types.ZendHashSetBucketKey(EG__().GetClassTable(), (*types.Bucket)(zv), lcname.GetStr().GetStr())
	if zv == nil {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.GetName().GetVal())
		return types.FAILURE
	}
	if ZendDoLinkClass(ce, lc_parent_name) == types.FAILURE {

		/* Reload bucket pointer, the hash table may have been reallocated */

		zv = EG__().GetClassTable().KeyFind(lcname.GetStr().GetStr())
		types.ZendHashSetBucketKey(EG__().GetClassTable(), (*types.Bucket)(zv), rtd_key.GetStr().GetStr())
		return types.FAILURE
	}
	return types.SUCCESS
}
func ZendMarkFunctionAsGenerator() {
	if CG__().GetActiveOpArray().GetFunctionName() == nil {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "The \"yield\" expression can only be used inside a function")
	}
	if CG__().GetActiveOpArray().IsHasReturnType() {
		var return_info ZendArgInfo = CG__().GetActiveOpArray().GetArgInfo()[-1]
		if return_info.GetType().Code() != types.IS_ITERABLE {
			var msg *byte = "Generators may only declare a return type of Generator, Iterator, Traversable, or iterable, %s is not permitted"
			if !(return_info.GetType().IsClass()) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, msg, types.ZendGetTypeByConst(return_info.GetType().Code()))
			}
			if !(types.ZendStringEqualsLiteralCi(return_info.GetType().Name(), "Traversable")) && !(types.ZendStringEqualsLiteralCi(return_info.GetType().Name(), "Iterator")) && !(types.ZendStringEqualsLiteralCi(return_info.GetType().Name(), "Generator")) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, msg, types.ZEND_TYPE_NAME(return_info.GetType()).GetVal())
			}
		}
	}
	CG__().GetActiveOpArray().SetIsGenerator(true)
}
func ZendBuildDelayedEarlyBindingList(op_array *types.ZendOpArray) uint32 {
	if op_array.IsEarlyBinding() {
		var first_early_binding_opline = uint32 - 1
		var prev_opline_num = &first_early_binding_opline
		var opline = op_array.GetOpcodes()
		var end *ZendOp = opline + op_array.GetLast()
		for opline < end {
			if opline.GetOpcode() == ZEND_DECLARE_CLASS_DELAYED {
				*prev_opline_num = opline - op_array.GetOpcodes()
				prev_opline_num = opline.GetResult().GetOplineNum()
			}
			opline++
		}
		*prev_opline_num = -1
		return first_early_binding_opline
	}
	return uint32 - 1
}
func ZendDoDelayedEarlyBinding(op_array *types.ZendOpArray, first_early_binding_opline uint32) {
	if first_early_binding_opline != uint32-1 {
		var orig_in_compilation = CG__().GetInCompilation()
		var opline_num = first_early_binding_opline
		var run_time_cache *any
		if op_array.GetRunTimeCachePtr() == nil {
			var ptr any
			b.Assert(op_array.IsHeapRtCache())
			ptr = Emalloc(op_array.GetCacheSize() + b.SizeOf("void *"))
			ZEND_MAP_PTR_INIT(op_array.run_time_cache, ptr)
			ptr = (*byte)(ptr + b.SizeOf("void *"))
			ZEND_MAP_PTR_SET(op_array.run_time_cache, ptr)
			memset(ptr, 0, op_array.GetCacheSize())
		}
		run_time_cache = RUN_TIME_CACHE(op_array)
		CG__().SetInCompilation(1)
		for opline_num != uint32-1 {
			var opline *ZendOp = op_array.GetOpcodes()[opline_num]
			var lcname = RT_CONSTANT(opline, opline.GetOp1())
			var zv = EG__().GetClassTable().KeyFind((lcname + 1).GetStr().GetStr())
			if zv != nil {
				var ce = zv.GetCe()
				var lc_parent_name = RT_CONSTANT(opline, opline.GetOp2()).GetStr()
				var parent_ce *types.ClassEntry = types.ZendHashFindPtr(EG__().GetClassTable(), lc_parent_name.GetStr())
				if parent_ce != nil {
					if ZendTryEarlyBind(ce, parent_ce, lcname.GetStr(), zv) != 0 {

						/* Store in run-time cache */

						(*any)((*byte)(run_time_cache + opline.GetExtendedValue()))[0] = ce

						/* Store in run-time cache */

					}
				}
			}
			opline_num = op_array.GetOpcodes()[opline_num].GetResult().GetOplineNum()
		}
		CG__().SetInCompilation(orig_in_compilation)
	}
}

func ZendManglePropertyName_Ex(src1 string, src2 string) string {
	propName := "\000" + src1 + "\000" + src2
	return propName
}

func ZendManglePropertyName_ZStr(src1 string, src2 string) *types.String {
	str := ZendManglePropertyName_Ex(src1, src2)
	return types.NewString(str)
}

func ZendUnmanglePropertyName_Ex(name string) (className string, propName string, ok bool) {
	if len(name) == 0 || name[0] != '\000' {
		return "", name, true
	}
	if len(name) < 3 || name[1] == '\000' {
		faults.Error(faults.E_NOTICE, "Illegal member variable name")
		return "", name, false
	}
	/*
	 * 可能的Name结构
	 * -	\0 + {className} + \0 + {$propName}
	 * -	\0 + {className} + \0 + {annoClassSrc} + \0 + {$propName}
	 */
	parts := strings.SplitN(name[1:], "\000", 3)
	switch len(parts) {
	case 2:
		return parts[0], parts[1], true
	case 3:
		return parts[0], parts[2], true
	default:
		faults.Error(faults.E_NOTICE, "Corrupt member variable name")
		return "", name, false
	}
}

func ZendUnmanglePropertyNameEx(name *types.String, class_name **byte, prop_name **byte, prop_len *int) int {
	className, propName, ok := ZendUnmanglePropertyName_Ex(name.GetStr())

	*class_name = className
	*prop_name = propName
	if prop_len != nil {
		*prop_len = len(propName)
	}
	if ok {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func ZendLookupReservedConst(name *byte, len_ int) *ZendConstant {
	var c *ZendConstant = ZendHashFindPtrLc(EG__().GetZendConstants(), name, len_)
	if c != nil && (ZEND_CONSTANT_FLAGS(c)&CONST_CS) == 0 && (ZEND_CONSTANT_FLAGS(c)&CONST_CT_SUBST) != 0 {
		return c
	}
	return nil
}
func ZendTryCtEvalConst(zv *types.Zval, name *types.String, is_fully_qualified types.ZendBool) types.ZendBool {
	var c *ZendConstant

	/* Substitute case-sensitive (or lowercase) constants */

	c = types.ZendHashFindPtr(EG__().GetZendConstants(), name.GetStr())
	if c != nil && ((ZEND_CONSTANT_FLAGS(c)&CONST_PERSISTENT) != 0 && (CG__().GetCompilerOptions()&ZEND_COMPILE_NO_PERSISTENT_CONSTANT_SUBSTITUTION) == 0 && ((ZEND_CONSTANT_FLAGS(c)&CONST_NO_FILE_CACHE) == 0 || (CG__().GetCompilerOptions()&ZEND_COMPILE_WITH_FILE_CACHE) == 0) || c.Value().GetType() < types.IS_OBJECT && (CG__().GetCompilerOptions()&ZEND_COMPILE_NO_CONSTANT_SUBSTITUTION) == 0) {
		types.ZVAL_COPY_OR_DUP(zv, c.Value())
		return 1
	}

	/* Substitute true, false and null (including unqualified usage in namespaces) */

	var lookup_name *byte = name.GetVal()
	var lookup_len = name.GetLen()
	if is_fully_qualified == 0 {
		ZendGetUnqualifiedName(name, &lookup_name, &lookup_len)
	}
	c = ZendLookupReservedConst(lookup_name, lookup_len)
	if c != nil {
		types.ZVAL_COPY_OR_DUP(zv, c.Value())
		return 1
	}
	return 0
}
func ZendIsScopeKnown() types.ZendBool {
	if CG__().GetActiveOpArray().IsClosure() {

		/* Closures can be rebound to a different scope */

		return 0

		/* Closures can be rebound to a different scope */

	}
	if CG__().GetActiveClassEntry() == nil {

		/* The scope is known if we're in a free function (no scope), but not if we're in
		 * a file/eval (which inherits including/eval'ing scope). */

		return CG__().GetActiveOpArray().GetFunctionName() != nil

		/* The scope is known if we're in a free function (no scope), but not if we're in
		 * a file/eval (which inherits including/eval'ing scope). */

	}

	/* For traits self etc refers to the using class, not the trait itself */

	return !CG__().GetActiveClassEntry().IsTrait()

	/* For traits self etc refers to the using class, not the trait itself */
}
func ClassNameRefersToActiveCe(class_name *types.String, fetch_type uint32) types.ZendBool {
	if CG__().GetActiveClassEntry() == nil {
		return 0
	}
	if fetch_type == ZEND_FETCH_CLASS_SELF && ZendIsScopeKnown() != 0 {
		return 1
	}
	return fetch_type == ZEND_FETCH_CLASS_DEFAULT && types.ZendStringEqualsCi(class_name, CG__().GetActiveClassEntry().GetName())
}
func ZendGetClassFetchType(name *types.String) uint32 {
	if types.ZendStringEqualsLiteralCi(name, "self") {
		return ZEND_FETCH_CLASS_SELF
	} else if types.ZendStringEqualsLiteralCi(name, "parent") {
		return ZEND_FETCH_CLASS_PARENT
	} else if types.ZendStringEqualsLiteralCi(name, "static") {
		return ZEND_FETCH_CLASS_STATIC
	} else {
		return ZEND_FETCH_CLASS_DEFAULT
	}
}
func ZendGetClassFetchTypeAst(name_ast *ZendAst) uint32 {
	/* Fully qualified names are always default refs */

	if name_ast.GetAttr() == ZEND_NAME_FQ {
		return ZEND_FETCH_CLASS_DEFAULT
	}
	return ZendGetClassFetchType(ZendAstGetStr(name_ast))
}
func ZendEnsureValidClassFetchType(fetch_type uint32) {
	if fetch_type != ZEND_FETCH_CLASS_DEFAULT && ZendIsScopeKnown() != 0 {
		var ce = CG__().GetActiveClassEntry()
		if ce == nil {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use \"%s\" when no class scope is active", b.Cond(b.Cond(fetch_type == ZEND_FETCH_CLASS_SELF, "self", fetch_type == ZEND_FETCH_CLASS_PARENT), "parent", "static"))
		} else if fetch_type == ZEND_FETCH_CLASS_PARENT && !(ce.GetParentName()) {
			faults.Error(faults.E_DEPRECATED, "Cannot use \"parent\" when current class scope has no parent")
		}
	}
}
func ZendTryCompileConstExprResolveClassName(zv *types.Zval, class_ast *ZendAst) types.ZendBool {
	var fetch_type uint32
	var class_name *types.Zval
	if class_ast.GetKind() != ZEND_AST_ZVAL {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use ::class with dynamic class name")
	}
	class_name = ZendAstGetZval(class_ast)
	if class_name.GetType() != types.IS_STRING {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Illegal class name")
	}
	fetch_type = ZendGetClassFetchType(class_name.GetStr())
	ZendEnsureValidClassFetchType(fetch_type)
	switch fetch_type {
	case ZEND_FETCH_CLASS_SELF:
		if CG__().GetActiveClassEntry() != nil && ZendIsScopeKnown() != 0 {
			zv.SetStringCopy(CG__().GetActiveClassEntry().GetName())
			return 1
		}
		return 0
	case ZEND_FETCH_CLASS_PARENT:
		if CG__().GetActiveClassEntry() != nil && CG__().GetActiveClassEntry().GetParentName() && ZendIsScopeKnown() != 0 {
			zv.SetStringCopy(CG__().GetActiveClassEntry().GetParentName())
			return 1
		}
		return 0
	case ZEND_FETCH_CLASS_STATIC:
		return 0
	case ZEND_FETCH_CLASS_DEFAULT:
		zv.SetString(ZendResolveClassNameAst(class_ast))
		return 1
	default:

	}
}
func ZendVerifyCtConstAccess(c *ZendClassConstant, scope *types.ClassEntry) types.ZendBool {
	if (c.GetValue().GetAccessFlags() & AccPublic) != 0 {
		return 1
	} else if (c.GetValue().GetAccessFlags() & AccPrivate) != 0 {
		return c.GetCe() == scope
	} else {
		var ce = c.GetCe()
		for true {
			if ce == scope {
				return 1
			}
			if !(ce.GetParent()) {
				break
			}
			if ce.IsResolvedParent() {
				ce = ce.GetParent()
			} else {
				ce = CG__().ClassTable().Get(ce.GetParentName().GetStr())
				if ce == nil {
					break
				}
			}
		}

		/* Reverse case cannot be true during compilation */

		return 0

		/* Reverse case cannot be true during compilation */

	}
}
func ZendTryCtEvalClassConst(zv *types.Zval, class_name *types.String, name *types.String) types.ZendBool {
	var fetch_type = ZendGetClassFetchType(class_name)
	var cc *ZendClassConstant
	var c *types.Zval
	if ClassNameRefersToActiveCe(class_name, fetch_type) != 0 {
		cc = types.ZendHashFindPtr(CG__().GetActiveClassEntry().GetConstantsTable(), name.GetStr())
	} else if fetch_type == ZEND_FETCH_CLASS_DEFAULT && (CG__().GetCompilerOptions()&ZEND_COMPILE_NO_CONSTANT_SUBSTITUTION) == 0 {
		ce := CG__().ClassTable().Get(class_name.GetStr())
		if ce != nil {
			cc = types.ZendHashFindPtr(ce.GetConstantsTable(), name.GetStr())
		} else {
			return 0
		}
	} else {
		return 0
	}
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_NO_PERSISTENT_CONSTANT_SUBSTITUTION) != 0 {
		return 0
	}
	if cc == nil || ZendVerifyCtConstAccess(cc, CG__().GetActiveClassEntry()) == 0 {
		return 0
	}
	c = cc.GetValue()

	/* Substitute case-sensitive (or lowercase) persistent class constants */

	if c.GetType() < types.IS_OBJECT {
		types.ZVAL_COPY_OR_DUP(zv, c)
		return 1
	}
	return 0
}
func ZendAddToList(result any, item any) {
	var list *any = *((*any)(result))
	var n = 0
	if list != nil {
		for list[n] {
			n++
		}
	}
	list = Erealloc(list, b.SizeOf("void *")*(n+2))
	list[n] = item
	list[n+1] = nil
	*((*any)(result)) = list
}
func ZendDoExtendedStmt() {
	var opline *ZendOp
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_EXTENDED_STMT) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_EXT_STMT)
}
func ZendDoExtendedFcallBegin() {
	var opline *ZendOp
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_EXTENDED_FCALL) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_EXT_FCALL_BEGIN)
}
func ZendDoExtendedFcallEnd() {
	var opline *ZendOp
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_EXTENDED_FCALL) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_EXT_FCALL_END)
}
func ZendIsAutoGlobalStr(name string) types.ZendBool {
	var autoGlobal = types.ZendHashStrFindPtr(CG__().GetAutoGlobals(), name).(*ZendAutoGlobal)
	if autoGlobal != nil {
		if autoGlobal.GetArmed() != 0 {
			autoGlobal.SetArmed(autoGlobal.GetAutoGlobalCallback()(autoGlobal.GetName()))
		}
		return 1
	}
	return 0
}
func ZendIsAutoGlobal(name *types.String) types.ZendBool {
	var autoGlobal = types.ZendHashFindPtr(CG__().GetAutoGlobals(), name.GetStr()).(*ZendAutoGlobal)
	if autoGlobal != nil {
		if autoGlobal.GetArmed() != 0 {
			autoGlobal.SetArmed(autoGlobal.GetAutoGlobalCallback()(autoGlobal.GetName()))
		}
		return 1
	}
	return 0
}
func ZendRegisterAutoGlobal(name *types.String, jit types.ZendBool, auto_global_callback ZendAutoGlobalCallback) int {
	var auto_global ZendAutoGlobal
	var retval int
	auto_global.SetName(name)
	auto_global.SetAutoGlobalCallback(auto_global_callback)
	auto_global.SetJit(jit)
	if types.ZendHashAddMem(CG__().GetAutoGlobals(), auto_global.GetName().GetStr(), &auto_global, b.SizeOf("zend_auto_global")) != nil {
		retval = types.SUCCESS
	} else {
		retval = types.FAILURE
	}
	return retval
}
func ZendActivateAutoGlobals() {
	var auto_global *ZendAutoGlobal
	var __ht = CG__().GetAutoGlobals()
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()

		auto_global = _z.GetPtr()
		if auto_global.GetJit() != 0 {
			auto_global.SetArmed(1)
		} else if auto_global.GetAutoGlobalCallback() != nil {
			auto_global.SetArmed(auto_global.GetAutoGlobalCallback()(auto_global.GetName()))
		} else {
			auto_global.SetArmed(0)
		}
	}
}
func Zendlex(elem *ZendParserStackElem) int {
	var zv types.Zval
	var ret int
	if CG__().GetIncrementLineno() != 0 {
		CG__().GetZendLineno()++
		CG__().SetIncrementLineno(0)
	}
	ret = LexScan(&zv, elem)
	b.Assert(EG__().GetException() == nil || ret == T_ERROR)
	return ret
}
func ZendInitializeClassData(ce *types.ClassEntry, nullify_handlers types.ZendBool) {
	var persistent_hashes = ce.GetType() == ZEND_INTERNAL_CLASS
	ce.SetRefcount(1)
	ce.SetCeFlags(AccConstantsUpdated)
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_GUARDS) != 0 {
		ce.SetIsUseGuards(true)
	}
	ce.SetDefaultPropertiesTable(nil)
	ce.SetDefaultStaticMembersTable(nil)
	ce.GetPropertiesInfo() = types.MakeArrayEx(8, b.Cond(persistent_hashes, ZendDestroyPropertyInfoInternal, nil), persistent_hashes)
	ce.GetConstantsTable() = types.MakeArrayEx(8, nil, persistent_hashes)
	ce.GetFunctionTable() = types.MakeArrayEx(8, ZEND_FUNCTION_DTOR, persistent_hashes)
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		ZEND_MAP_PTR_INIT(ce.static_members_table, nil)
	} else {
		ZEND_MAP_PTR_INIT(ce.static_members_table, ce.GetDefaultStaticMembersTable())
		ce.SetDocComment(nil)
	}
	ce.SetDefaultPropertiesCount(0)
	ce.SetDefaultStaticMembersCount(0)
	ce.SetPropertiesInfoTable(nil)
	if nullify_handlers != 0 {
		ce.SetConstructor(nil)
		ce.SetDestructor(nil)
		ce.SetClone(nil)
		ce.SetGet(nil)
		ce.SetSet(nil)
		ce.SetUnset(nil)
		ce.SetIsset(nil)
		ce.SetCall(nil)
		ce.SetCallstatic(nil)
		ce.SetTostring(nil)
		ce.SetCreateObject(nil)
		ce.SetGetIterator(nil)
		ce.SetIteratorFuncsPtr(nil)
		ce.SetGetStaticMethod(nil)
		ce.SetParent(nil)
		ce.SetParentName(nil)
		ce.SetNumInterfaces(0)
		ce.SetInterfaces(nil)
		ce.SetNumTraits(0)
		ce.SetTraitNames(nil)
		ce.SetTraitAliases(nil)
		ce.SetTraitPrecedences(nil)
		ce.SetSerialize(nil)
		ce.SetUnserialize(nil)
		ce.SetSerializeFunc(nil)
		ce.SetUnserializeFunc(nil)
		ce.SetDebugInfo(nil)
		if ce.GetType() == ZEND_INTERNAL_CLASS {
			ce.SetModule(nil)
			ce.SetBuiltinFunctions(nil)
		}
	}
}
func ZendGetCompiledVariableName(op_array *types.ZendOpArray, var_ uint32) *types.String {
	return op_array.GetVars()[EX_VAR_TO_NUM(var_)]
}
func ZendAstAppendStr(left_ast *ZendAst, right_ast *ZendAst) *ZendAst {
	var left_zv = ZendAstGetZval(left_ast)
	var left = left_zv.GetStr()
	var right = ZendAstGetStr(right_ast)
	var result *types.String
	var left_len = left.GetLen()
	var len_ = left_len + right.GetLen() + 1
	result = types.ZendStringExtend(left, len_, 0)
	result.GetVal()[left_len] = '\\'
	memcpy(&result.GetVal()[left_len+1], right.GetVal(), right.GetLen())
	result.GetVal()[len_] = '0'
	types.ZendStringReleaseEx(right, 0)
	left_zv.SetString(result)
	return left_ast
}
func ZendNegateNumString(ast *ZendAst) *ZendAst {
	var zv = ZendAstGetZval(ast)
	if zv.IsLong() {
		if zv.GetLval() == 0 {
			zv.SetString(types.NewString("-0"))
		} else {
			b.Assert(zv.GetLval() > 0)
			zv.SetLval(zv.GetLval() * -1)
		}
	} else if zv.IsString() {
		var orig_len = zv.GetStr().GetLen()
		zv.SetStr(types.ZendStringExtend(zv.GetStr(), orig_len+1, 0))
		memmove(zv.GetStr().GetVal()+1, zv.GetStr().GetVal(), orig_len+1)
		zv.GetStr().GetVal()[0] = '-'
	} else {
		b.Assert(false)
	}
	return ast
}
func ZendVerifyNamespace() {
	if FC__().GetHasBracketedNamespaces() != 0 && FC__().GetInNamespace() == 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "No code may exist outside of namespace {}")
	}
}
func ZendDirname(path *byte, len_ int) int {
	var end *byte = path + len_ - 1
	var len_adjust uint = 0
	if len_ == 0 {

		/* Illegal use of this function */

		return 0

		/* Illegal use of this function */

	}

	/* Strip trailing slashes */

	for end >= path && IS_SLASH_P(end) {
		end--
	}
	if end < path {

		/* The path only contained slashes */

		path[0] = DEFAULT_SLASH
		path[1] = '0'
		return 1 + len_adjust
	}

	/* Strip filename */

	for end >= path && !(IS_SLASH_P(end)) {
		end--
	}
	if end < path {

		/* No slash found, therefore return '.' */

		path[0] = '.'
		path[1] = '0'
		return 1 + len_adjust
	}

	/* Strip slashes which came before the file name */

	for end >= path && IS_SLASH_P(end) {
		end--
	}
	if end < path {
		path[0] = DEFAULT_SLASH
		path[1] = '0'
		return 1 + len_adjust
	}
	*(end + 1) = '0'
	return size_t(end+1-path) + len_adjust
}
func ZendAdjustForFetchType(opline *ZendOp, result *Znode, type_ uint32) {
	var factor types.ZendUchar = b.Cond(opline.GetOpcode() == ZEND_FETCH_STATIC_PROP_R, 1, 3)
	switch type_ {
	case BP_VAR_R:
		opline.SetResultType(IS_TMP_VAR)
		result.SetOpType(IS_TMP_VAR)
		return
	case BP_VAR_W:
		opline.SetOpcode(opline.GetOpcode() + 1*factor)
		return
	case BP_VAR_RW:
		opline.SetOpcode(opline.GetOpcode() + 2*factor)
		return
	case BP_VAR_IS:
		opline.SetResultType(IS_TMP_VAR)
		result.SetOpType(IS_TMP_VAR)
		opline.SetOpcode(opline.GetOpcode() + 3*factor)
		return
	case BP_VAR_FUNC_ARG:
		opline.SetOpcode(opline.GetOpcode() + 4*factor)
		return
	case BP_VAR_UNSET:
		opline.SetOpcode(opline.GetOpcode() + 5*factor)
		return
	default:

	}
}
