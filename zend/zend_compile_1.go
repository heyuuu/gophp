package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"strings"
)

func ZendResolveNonClassName(name string, typ uint32, caseSensitive bool, currentImportSub ImportNames) (string, bool) {
	isFullyQualified := false
	if name[0] == '\\' {
		/* Remove \ prefix (only relevant if this is a string rather than a label) */
		return name[1:], true
	}
	if typ == ZEND_NAME_FQ {
		return name, true
	}
	if typ == ZEND_NAME_RELATIVE {
		return ZendPrefixWithNsEx(name), true
	}
	if currentImportSub != nil {

		/* If an unqualified name is a function/const alias, replace it. */

		var importName string
		if caseSensitive {
			importName = currentImportSub.Get(name)
		} else {
			importName = currentImportSub.Get(ascii.StrToLower(name))
		}
		if importName != "" {
			return importName, true
		}
	}

	compoundPos := strings.IndexByte(name, '\\')
	if compoundPos >= 0 {
		isFullyQualified = true
	}

	if compoundPos >= 0 && FC__().GetImports() != nil {
		/* If the first part of a qualified name is an alias, substitute it. */
		var importName = FC__().GetImports().Get(ascii.StrToLower(name[:compoundPos]))
		if importName != "" {
			return ZendConcatNames(importName, name[compoundPos+1:]), isFullyQualified
		}
	}
	return ZendPrefixWithNsEx(name), isFullyQualified
}

func ZendResolveFunctionName(name string, typ uint32) (resolveName string, isFullyQualified bool) {
	return ZendResolveNonClassName(name, typ, false, FC__().ImportsFunction())
}
func ZendResolveConstName(name string, typ uint32) (resolveName string, isFullyQualified bool) {
	return ZendResolveNonClassName(name, typ, true, FC__().ImportsConst())
}

func ZendResolveClassName(name string, typ uint32) string {
	if typ == ZEND_NAME_RELATIVE {
		return ZendPrefixWithNsEx(name)
	}
	if typ == ZEND_NAME_FQ || (name != "" && name[0] == '\\') {
		/* Remove \ prefix (only relevant if this is a string rather than a label) */
		if name != "" && name[0] == '\\' {
			name = name[1:]
		}

		/* Ensure that \self, \parent and \static are not used */
		if ZEND_FETCH_CLASS_DEFAULT != ZendGetClassFetchType(name) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'\\%s' is an invalid class name", name)
		}
		return name
	}
	if FC__().GetImports().Len() != 0 {
		pos := strings.IndexByte(name, '\\')
		if pos >= 0 {
			/* If the first part of a qualified name is an alias, substitute it. */
			var importName = FC__().Imports().Get(name[:pos])
			if importName != "" {
				return ZendConcatNames(importName, name[pos+1:])
			}
		} else {
			/* If an unqualified name is an alias, replace it. */
			var importName = FC__().Imports().Get(name)
			if importName != "" {
				return importName
			}
		}
	}

	/* If not fully qualified and not an alias, prepend the current namespace */
	return ZendPrefixWithNsEx(name)
}
func ZendResolveClassNameAst(ast *ZendAst) *types.String {
	var class_name = ZendAstGetZval(ast)
	if class_name.GetType() != types.IS_STRING {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Illegal class name")
	}
	resolveName := ZendResolveClassName(class_name.StringVal(), ast.GetAttr())
	return types.NewString(resolveName)
}
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
func FunctionAddRef(function types.IFunction) {
	if function.GetType() == ZEND_USER_FUNCTION {
		var op_array = function.GetOpArray()
		op_array.TryIncRefCount()
		preload := CG__().IsCompilePreload()
		op_array.InitPtr2(preload)
	} else if function.GetType() == ZEND_INTERNAL_FUNCTION {
		//if function.GetFunctionName() != nil {
		//function.GetFunctionName().AddRefcount()
		//}
	}
}
func DoBindFunctionError(lcname *types.String, op_array *types.ZendOpArray, compile_time types.ZendBool) {
	var oldFunction types.IFunction
	var errorLevel int
	if compile_time != 0 {
		errorLevel = faults.E_COMPILE_ERROR
		oldFunction = CG__().FunctionTable().Get(lcname.GetStr())
	} else {
		errorLevel = faults.E_ERROR
		oldFunction = EG__().FunctionTable().Get(lcname.GetStr())
	}

	b.Assert(oldFunction != nil)
	var functionName string
	if op_array != nil {
		functionName = op_array.FunctionName()
	} else {
		functionName = oldFunction.FunctionName()
	}

	if oldFunction.GetType() == ZEND_USER_FUNCTION && oldFunction.GetOpArray().GetLast() > 0 {
		faults.ErrorNoreturn(errorLevel, "Cannot redeclare %s() (previously declared in %s:%d)", functionName, oldFunction.GetOpArray().GetFilename(), oldFunction.GetOpArray().GetOpcodes()[0].GetLineno())
	} else {
		faults.ErrorNoreturn(errorLevel, "Cannot redeclare %s()", functionName)
	}
}
func DoBindFunction(lcname *types.Zval) int {
	var function types.IFunction
	var rtd_key *types.Zval
	rtd_key = lcname + 1
	function = EG__().FunctionTable().Get(rtd_key.String().GetStr())
	if function == nil {
		DoBindFunctionError(lcname.String(), nil, 0)
		return types.FAILURE
	}

	if EG__().FunctionTable().Exists(lcname.StringVal()) {
		DoBindFunctionError(lcname.String(), function.GetOpArray(), 0)
		return types.FAILURE
	}

	if function.IsPreloaded() && !CG__().IsCompilePreload() {
		EG__().FunctionTable().Add(lcname.StringVal(), function)
	} else {
		EG__().FunctionTable().Del(rtd_key.StringVal())
		EG__().FunctionTable().Add(lcname.StringVal(), function)
	}
	return types.SUCCESS
}
func DoBindClass(lcname *types.Zval, lc_parent_name *types.String) int {
	var rtd_key *types.Zval = lcname + 1

	ce := EG__().ClassTable().Get(rtd_key.StringVal())
	if ce == nil {
		if EG__().ClassTable().Exists(lcname.StringVal()) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.Name())
			return types.FAILURE
		} else {
			b.Assert(CurrEX().GetFunc().GetOpArray().IsPreloaded())
			faults.ErrorNoreturn(faults.E_ERROR, "Class %s wasn't preloaded", lcname.String().GetVal())
			return types.FAILURE
		}
	}

	if EG__().ClassTable().Exists(lcname.StringVal()) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.Name())
		return types.FAILURE
	}

	if ZendDoLinkClass(ce, lc_parent_name) == types.FAILURE {
		return types.FAILURE
	}

	EG__().ClassTable().Del(rtd_key.StringVal())
	EG__().ClassTable().Add(lcname.StringVal(), ce)

	return types.SUCCESS
}
func ZendMarkFunctionAsGenerator() {
	if CG__().GetActiveOpArray().GetFunctionName() == nil {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "The \"yield\" expression can only be used inside a function")
	}
	if CG__().GetActiveOpArray().IsHasReturnType() {
		var return_info ZendArgInfo = CG__().GetActiveOpArray().GetArgInfo()[-1]
		if return_info.GetType().Code() != types.IS_ITERABLE {
			var msg = "Generators may only declare a return type of Generator, Iterator, Traversable, or iterable, %s is not permitted"
			if !(return_info.GetType().IsClass()) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, msg, types.ZendGetTypeByConst(return_info.GetType().Code()))
			}
			if !(ascii.StrCaseEquals(return_info.GetType().Name(), "Traversable")) && !(ascii.StrCaseEquals(return_info.GetType().Name(), "Iterator")) && !(ascii.StrCaseEquals(return_info.GetType().Name(), "Generator")) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, msg, return_info.GetType().Name())
			}
		}
	}
	CG__().GetActiveOpArray().SetIsGenerator(true)
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
func ZendLookupReservedConst(name string) *ZendConstant {
	var c *ZendConstant = EG__().ConstantTable().Get(name)
	if c != nil && !c.IsCaseSensitive() && c.IsCtSubst() {
		return c
	}
	return nil
}
func ZendTryCtEvalConst(zv *types.Zval, name string, is_fully_qualified bool) types.ZendBool {
	var c *ZendConstant = EG__().ConstantTable().Get(name)

	/* Substitute case-sensitive (or lowercase) constants */
	if c != nil && ((c.IsPersistent() && (CG__().GetCompilerOptions()&ZEND_COMPILE_NO_PERSISTENT_CONSTANT_SUBSTITUTION) == 0 && !c.IsNoFileCache() || (CG__().GetCompilerOptions()&ZEND_COMPILE_WITH_FILE_CACHE) == 0) || c.Value().GetType() < types.IS_OBJECT && (CG__().GetCompilerOptions()&ZEND_COMPILE_NO_CONSTANT_SUBSTITUTION) == 0) {
		types.ZVAL_COPY_OR_DUP(zv, c.Value())
		return 1
	}

	/* Substitute true, false and null (including unqualified usage in namespaces) */
	if is_fully_qualified {
		name, _ = ZendGetUnqualifiedNameEx(name)
	}
	c = ZendLookupReservedConst(name)
	if c != nil {
		types.ZVAL_COPY_OR_DUP(zv, c.Value())
		return 1
	}
	return 0
}
func ZendIsScopeKnown() bool {
	if CG__().GetActiveOpArray().IsClosure() {
		/* Closures can be rebound to a different scope */
		return false
	}
	if CG__().GetActiveClassEntry() == nil {
		/* The scope is known if we're in a free function (no scope), but not if we're in
		 * a file/eval (which inherits including/eval'ing scope). */
		return CG__().GetActiveOpArray().FunctionName() != ""
	}

	/* For traits self etc refers to the using class, not the trait itself */
	return !CG__().GetActiveClassEntry().IsTrait()
}
func ClassNameRefersToActiveCe(class_name *types.String, fetch_type uint32) bool {
	if CG__().GetActiveClassEntry() == nil {
		return false
	}
	if fetch_type == ZEND_FETCH_CLASS_SELF && ZendIsScopeKnown() {
		return true
	}
	return fetch_type == ZEND_FETCH_CLASS_DEFAULT && ascii.StrCaseEquals(class_name.GetStr(), CG__().GetActiveClassEntry().Name())
}
func ZendGetClassFetchType(name string) uint32 {
	if ascii.StrCaseEquals(name, "self") {
		return ZEND_FETCH_CLASS_SELF
	} else if ascii.StrCaseEquals(name, "parent") {
		return ZEND_FETCH_CLASS_PARENT
	} else if ascii.StrCaseEquals(name, "static") {
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
	return ZendGetClassFetchType(ZendAstGetStrVal(name_ast))
}
func ZendEnsureValidClassFetchType(fetch_type uint32) {
	if fetch_type != ZEND_FETCH_CLASS_DEFAULT && ZendIsScopeKnown() {
		var ce = CG__().GetActiveClassEntry()
		if ce == nil {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use \"%s\" when no class scope is active", b.Cond(b.Cond(fetch_type == ZEND_FETCH_CLASS_SELF, "self", fetch_type == ZEND_FETCH_CLASS_PARENT), "parent", "static"))
		} else if fetch_type == ZEND_FETCH_CLASS_PARENT && ce.GetParentName() == nil {
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
	fetch_type = ZendGetClassFetchType(class_name.String().GetStr())
	ZendEnsureValidClassFetchType(fetch_type)
	switch fetch_type {
	case ZEND_FETCH_CLASS_SELF:
		if CG__().GetActiveClassEntry() != nil && ZendIsScopeKnown() {
			zv.SetStringVal(CG__().GetActiveClassEntry().GetName().GetStr())
			return 1
		}
		return 0
	case ZEND_FETCH_CLASS_PARENT:
		if CG__().GetActiveClassEntry() != nil && CG__().GetActiveClassEntry().GetParentName() != nil && ZendIsScopeKnown() {
			zv.SetStringVal(CG__().GetActiveClassEntry().GetParentName().GetStr())
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
func ZendVerifyCtConstAccess(c *types.ClassConstant, scope *types.ClassEntry) types.ZendBool {
	if (c.GetValue().GetAccessFlags() & types.AccPublic) != 0 {
		return 1
	} else if (c.GetValue().GetAccessFlags() & types.AccPrivate) != 0 {
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
	var fetch_type = ZendGetClassFetchType(class_name.GetStr())
	var cc *types.ClassConstant
	var c *types.Zval
	if ClassNameRefersToActiveCe(class_name, fetch_type) {
		cc = CG__().GetActiveClassEntry().ConstantsTable().Get(name.GetStr())
	} else if fetch_type == ZEND_FETCH_CLASS_DEFAULT && (CG__().GetCompilerOptions()&ZEND_COMPILE_NO_CONSTANT_SUBSTITUTION) == 0 {
		ce := CG__().ClassTable().Get(class_name.GetStr())
		if ce != nil {
			cc = ce.ConstantsTable().Get(name.GetStr())
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
	var opline *types.ZendOp
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_EXTENDED_STMT) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_EXT_STMT)
}
func ZendDoExtendedFcallBegin() {
	var opline *types.ZendOp
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_EXTENDED_FCALL) == 0 {
		return
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_EXT_FCALL_BEGIN)
}
func ZendDoExtendedFcallEnd() {
	var opline *types.ZendOp
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
	CG__().GetAutoGlobals().Foreach(func(_ types.ArrayKey, value *types.Zval) {
		var autoGlobal *ZendAutoGlobal = value.Ptr()
		if autoGlobal.GetJit() != 0 {
			autoGlobal.SetArmed(1)
		} else if autoGlobal.GetAutoGlobalCallback() != nil {
			autoGlobal.SetArmed(autoGlobal.GetAutoGlobalCallback()(autoGlobal.GetName()))
		} else {
			autoGlobal.SetArmed(0)
		}
	})
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

func ZendVerifyNamespace() {
	if FC__().GetHasBracketedNamespaces() != 0 && FC__().GetInNamespace() == 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "No code may exist outside of namespace {}")
	}
}
func ZendDirname(path string) string {
	if path == "" {
		return ""
	}

	/* Strip trailing slashes */
	path = strings.TrimRight(path, "/")
	if path == "" {
		/* The path only contained slashes */
		return "/"
	}

	/* Strip filename */
	if pos := strings.LastIndexByte(path, '/'); pos >= 0 {
		path = path[:pos]
	} else {
		/* No slash found, therefore return '.' */
		return "."
	}

	/* Strip slashes which came before the file name */
	path = strings.TrimRight(path, "/")
	if path == "" {
		return "."
	}
	return path
}
func ZendAdjustForFetchType(opline *types.ZendOp, result *Znode, type_ uint32) {
	var factor uint8 = b.Cond(opline.GetOpcode() == ZEND_FETCH_STATIC_PROP_R, 1, 3)
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
