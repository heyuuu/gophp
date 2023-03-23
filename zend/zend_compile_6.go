// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZendCompileClosureBinding(closure *Znode, op_array *ZendOpArray, uses_ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(uses_ast)
	var i uint32
	if list.GetChildren() == 0 {
		return
	}
	if op_array.GetStaticVariables() == nil {
		op_array.SetStaticVariables(types.NewArray(8))
	}
	for i = 0; i < list.GetChildren(); i++ {
		var var_name_ast *ZendAst = list.GetChild()[i]
		var var_name *types.String = ZvalMakeInternedString(ZendAstGetZval(var_name_ast))
		var mode uint32 = var_name_ast.GetAttr()
		var opline *ZendOp
		var value *types.Zval
		if types.ZendStringEqualsLiteral(var_name, "this") {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use $this as lexical variable")
		}
		if ZendIsAutoGlobal(var_name) != 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use auto-global as lexical variable")
		}
		value = op_array.GetStaticVariables().KeyAdd(var_name.GetStr(), EG__().GetUninitializedZval())
		if value == nil {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use variable $%s twice", var_name.GetVal())
		}
		CG__().SetZendLineno(ZendAstGetLineno(var_name_ast))
		opline = ZendEmitOp(nil, ZEND_BIND_LEXICAL, closure, nil)
		opline.SetOp2Type(IS_CV)
		opline.GetOp2().SetVar(LookupCv(var_name))
		opline.SetExtendedValue(uint32((*byte)(value-(*byte)(op_array.GetStaticVariables().GetArData()))) | mode)
	}
}
func FindImplicitBindsRecursively(info *ClosureInfo, ast *ZendAst) {
	if ast == nil {
		return
	}
	if ast.GetKind() == ZEND_AST_VAR {
		var name_ast *ZendAst = ast.GetChild()[0]
		if name_ast.GetKind() == ZEND_AST_ZVAL && ZendAstGetZval(name_ast).IsString() {
			var name *types.String = ZendAstGetStr(name_ast)
			if ZendIsAutoGlobal(name) != 0 {

				/* These is no need to explicitly import auto-globals. */

				return

				/* These is no need to explicitly import auto-globals. */

			}
			if types.ZendStringEqualsLiteral(name, "this") {

				/* $this does not need to be explicitly imported. */

				return

				/* $this does not need to be explicitly imported. */

			}
			types.ZendHashAddEmptyElement(info.GetUses(), name.GetStr())
		} else {
			info.SetVarvarsUsed(1)
			FindImplicitBindsRecursively(info, name_ast)
		}
	} else if ZendAstIsList(ast) != 0 {
		var list *ZendAstList = ZendAstGetList(ast)
		var i uint32
		for i = 0; i < list.GetChildren(); i++ {
			FindImplicitBindsRecursively(info, list.GetChild()[i])
		}
	} else if ast.GetKind() == ZEND_AST_CLOSURE {

		/* For normal closures add the use() list. */

		var closure_ast *ZendAstDecl = (*ZendAstDecl)(ast)
		var uses_ast *ZendAst = closure_ast.GetChild()[1]
		if uses_ast != nil {
			var uses_list *ZendAstList = ZendAstGetList(uses_ast)
			var i uint32
			for i = 0; i < uses_list.GetChildren(); i++ {
				types.ZendHashAddEmptyElement(info.GetUses(), ZendAstGetStr(uses_list.GetChild()[i]).GetStr())
			}
		}
	} else if ast.GetKind() == ZEND_AST_ARROW_FUNC {

		/* For arrow functions recursively check the expression. */

		var closure_ast *ZendAstDecl = (*ZendAstDecl)(ast)
		FindImplicitBindsRecursively(info, closure_ast.GetChild()[2])
	} else if ZendAstIsSpecial(ast) == 0 {
		var i uint32
		var children uint32 = ZendAstGetNumChildren(ast)
		for i = 0; i < children; i++ {
			FindImplicitBindsRecursively(info, ast.GetChild()[i])
		}
	}
}
func FindImplicitBinds(info *ClosureInfo, params_ast *ZendAst, stmt_ast *ZendAst) {
	var param_list *ZendAstList = ZendAstGetList(params_ast)
	var i uint32
	info.GetUses() = types.MakeArrayEx(param_list.GetChildren(), nil, 0)
	FindImplicitBindsRecursively(info, stmt_ast)

	/* Remove variables that are parameters */

	for i = 0; i < param_list.GetChildren(); i++ {
		var param_ast *ZendAst = param_list.GetChild()[i]
		types.ZendHashDel(info.GetUses(), ZendAstGetStr(param_ast.GetChild()[1]).GetStr())
	}
}
func CompileImplicitLexicalBinds(info *ClosureInfo, closure *Znode, op_array *ZendOpArray) {
	var var_name *types.String
	var opline *ZendOp

	/* TODO We might want to use a special binding mode if varvars_used is set. */

	if info.GetUses().CountElements() == 0 {
		return
	}
	if op_array.GetStaticVariables() == nil {
		op_array.SetStaticVariables(types.NewArray(8))
	}
	var __ht *types.Array = info.GetUses()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		var_name = _p.GetKey()
		var value *types.Zval = op_array.GetStaticVariables().KeyAdd(var_name.GetStr(), EG__().GetUninitializedZval())
		var offset uint32 = uint32((*byte)(value - (*byte)(op_array.GetStaticVariables().GetArData())))
		opline = ZendEmitOp(nil, ZEND_BIND_LEXICAL, closure, nil)
		opline.SetOp2Type(IS_CV)
		opline.GetOp2().SetVar(LookupCv(var_name))
		opline.SetExtendedValue(offset | ZEND_BIND_IMPLICIT)
	}
}
func ZendCompileClosureUses(ast *ZendAst) {
	var op_array *ZendOpArray = CG__().GetActiveOpArray()
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		var var_ast *ZendAst = list.GetChild()[i]
		var var_name *types.String = ZendAstGetStr(var_ast)
		var zv types.Zval
		zv.SetNull()
		var i int
		for i = 0; i < op_array.GetLastVar(); i++ {
			if types.ZendStringEquals(op_array.GetVars()[i], var_name) != 0 {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use lexical variable $%s as a parameter name", var_name.GetVal())
			}
		}
		CG__().SetZendLineno(ZendAstGetLineno(var_ast))
		ZendCompileStaticVarCommon(var_name, &zv, b.Cond(var_ast.GetAttr() != 0, ZEND_BIND_REF, 0))
	}
}
func ZendCompileImplicitClosureUses(info *ClosureInfo) {
	var var_name *types.String
	var __ht *types.Array = info.GetUses()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		var_name = _p.GetKey()
		var zv types.Zval
		zv.SetNull()
		ZendCompileStaticVarCommon(var_name, &zv, ZEND_BIND_IMPLICIT)
	}
}
func ZendBeginMethodDecl(op_array *ZendOpArray, name *types.String, has_body types.ZendBool) {
	var ce *types.ClassEntry = CG__().GetActiveClassEntry()
	var in_interface types.ZendBool = ce.IsInterface()
	var in_trait types.ZendBool = ce.IsTrait()
	var is_public types.ZendBool = op_array.IsPublic()
	var is_static types.ZendBool = op_array.IsStatic()
	var lcname *types.String
	if in_interface != 0 {
		if is_public == 0 || op_array.HasFnFlags(ZEND_ACC_FINAL|ZEND_ACC_ABSTRACT) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Access type for interface method "+"%s::%s() must be omitted", ce.GetName().GetVal(), name.GetVal())
		}
		op_array.SetIsAbstract(true)
	}
	if op_array.IsAbstract() {
		if op_array.IsPrivate() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "%s function %s::%s() cannot be declared private", b.Cond(in_interface != 0, "Interface", "Abstract"), ce.GetName().GetVal(), name.GetVal())
		}
		if has_body != 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "%s function %s::%s() cannot contain body", b.Cond(in_interface != 0, "Interface", "Abstract"), ce.GetName().GetVal(), name.GetVal())
		}
		ce.SetIsImplicitAbstractClass(true)
	} else if has_body == 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Non-abstract method %s::%s() must contain body", ce.GetName().GetVal(), name.GetVal())
	}
	op_array.SetScope(ce)
	op_array.SetFunctionName(name.Copy())
	lcname = ZendStringTolower(name)
	lcname = types.ZendNewInternedString(lcname)
	if types.ZendHashAddPtr(ce.GetFunctionTable(), lcname.GetStr(), op_array) == nil {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot redeclare %s::%s()", ce.GetName().GetVal(), name.GetVal())
	}
	if in_interface != 0 {
		if lcname.GetVal()[0] != '_' || lcname.GetVal()[1] != '_' {

		} else if types.ZendStringEqualsLiteral(lcname, ZEND_CALL_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __call() must have "+"public visibility and cannot be static")
			}
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_CALLSTATIC_FUNC_NAME) {
			if is_public == 0 || is_static == 0 {
				faults.Error(faults.E_WARNING, "The magic method __callStatic() must have "+"public visibility and be static")
			}
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_GET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __get() must have "+"public visibility and cannot be static")
			}
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_SET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __set() must have "+"public visibility and cannot be static")
			}
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_UNSET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __unset() must have "+"public visibility and cannot be static")
			}
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_ISSET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __isset() must have "+"public visibility and cannot be static")
			}
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_TOSTRING_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __toString() must have "+"public visibility and cannot be static")
			}
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_INVOKE_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __invoke() must have "+"public visibility and cannot be static")
			}
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_DEBUGINFO_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __debugInfo() must have "+"public visibility and cannot be static")
			}
		}
	} else {
		if in_trait == 0 && types.ZendStringEqualsCi(lcname, ce.GetName()) {
			if ce.GetConstructor() == nil {
				ce.SetConstructor((*ZendFunction)(op_array))
			}
		} else if types.ZendStringEqualsLiteral(lcname, "serialize") {
			ce.SetSerializeFunc((*ZendFunction)(op_array))
			if is_static == 0 {
				op_array.SetIsAllowStatic(true)
			}
		} else if types.ZendStringEqualsLiteral(lcname, "unserialize") {
			ce.SetUnserializeFunc((*ZendFunction)(op_array))
			if is_static == 0 {
				op_array.SetIsAllowStatic(true)
			}
		} else if lcname.GetVal()[0] != '_' || lcname.GetVal()[1] != '_' {
			if is_static == 0 {
				op_array.SetIsAllowStatic(true)
			}
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_CONSTRUCTOR_FUNC_NAME) {
			ce.SetConstructor((*ZendFunction)(op_array))
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_DESTRUCTOR_FUNC_NAME) {
			ce.SetDestructor((*ZendFunction)(op_array))
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_CLONE_FUNC_NAME) {
			ce.SetClone((*ZendFunction)(op_array))
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_CALL_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __call() must have "+"public visibility and cannot be static")
			}
			ce.SetCall((*ZendFunction)(op_array))
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_CALLSTATIC_FUNC_NAME) {
			if is_public == 0 || is_static == 0 {
				faults.Error(faults.E_WARNING, "The magic method __callStatic() must have "+"public visibility and be static")
			}
			ce.SetCallstatic((*ZendFunction)(op_array))
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_GET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __get() must have "+"public visibility and cannot be static")
			}
			ce.SetGet((*ZendFunction)(op_array))
			ce.SetIsUseGuards(true)
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_SET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __set() must have "+"public visibility and cannot be static")
			}
			ce.SetSet((*ZendFunction)(op_array))
			ce.SetIsUseGuards(true)
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_UNSET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __unset() must have "+"public visibility and cannot be static")
			}
			ce.SetUnset((*ZendFunction)(op_array))
			ce.SetIsUseGuards(true)
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_ISSET_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __isset() must have "+"public visibility and cannot be static")
			}
			ce.SetIsset((*ZendFunction)(op_array))
			ce.SetIsUseGuards(true)
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_TOSTRING_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __toString() must have "+"public visibility and cannot be static")
			}
			ce.SetTostring((*ZendFunction)(op_array))
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_INVOKE_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __invoke() must have "+"public visibility and cannot be static")
			}
		} else if types.ZendStringEqualsLiteral(lcname, ZEND_DEBUGINFO_FUNC_NAME) {
			if is_public == 0 || is_static != 0 {
				faults.Error(faults.E_WARNING, "The magic method __debugInfo() must have "+"public visibility and cannot be static")
			}
			ce.SetDebugInfo((*ZendFunction)(op_array))
		} else if is_static == 0 {
			op_array.SetIsAllowStatic(true)
		}
	}
	types.ZendStringReleaseEx(lcname, 0)
}
func ZendBeginFuncDecl(result *Znode, op_array *ZendOpArray, decl *ZendAstDecl, toplevel types.ZendBool) {
	var params_ast *ZendAst = decl.GetChild()[0]
	var unqualified_name *types.String
	var name *types.String
	var lcname *types.String
	var key *types.String
	var opline *ZendOp
	unqualified_name = decl.GetName()
	name = ZendPrefixWithNs(unqualified_name)
	op_array.SetFunctionName(name)
	lcname = ZendStringTolower(name)
	if FC__().GetImportsFunction() != nil {
		var import_name *types.String = ZendHashFindPtrLc(FC__().GetImportsFunction(), unqualified_name.GetVal(), unqualified_name.GetLen())
		if import_name != nil && !(types.ZendStringEqualsCi(lcname, import_name)) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare function %s "+"because the name is already in use", name.GetVal())
		}
	}
	if types.ZendStringEqualsLiteral(lcname, ZEND_AUTOLOAD_FUNC_NAME) {
		if ZendAstGetList(params_ast).GetChildren() != 1 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "%s() must take exactly 1 argument", ZEND_AUTOLOAD_FUNC_NAME)
		}
		faults.Error(faults.E_DEPRECATED, "__autoload() is deprecated, use spl_autoload_register() instead")
	}
	if types.ZendStringEqualsLiteralCi(unqualified_name, "assert") {
		faults.Error(faults.E_DEPRECATED, "Defining a custom assert() function is deprecated, "+"as the function has special semantics")
	}
	ZendRegisterSeenSymbol(lcname, ZEND_SYMBOL_FUNCTION)
	if toplevel != 0 {
		if types.ZendHashAddPtr(CG__().GetFunctionTable(), lcname.GetStr(), op_array) == nil {
			DoBindFunctionError(lcname, op_array, 1)
		}
		types.ZendStringReleaseEx(lcname, 0)
		return
	}

	/* Generate RTD keys until we find one that isn't in use yet. */

	key = nil
	for {
		ZendTmpStringRelease(key)
		key = ZendBuildRuntimeDefinitionKey(lcname, decl.GetStartLineno())
		if types.ZendHashAddPtr(CG__().GetFunctionTable(), key.GetStr(), op_array) {
			break
		}
	}
	if op_array.IsClosure() {
		opline = ZendEmitOpTmp(result, ZEND_DECLARE_LAMBDA_FUNCTION, nil, nil)
		opline.SetExtendedValue(ZendAllocCacheSlot())
		opline.SetOp1Type(IS_CONST)
		LITERAL_STR(opline.GetOp1(), key)
	} else {
		opline = GetNextOp()
		opline.SetOpcode(ZEND_DECLARE_FUNCTION)
		opline.SetOp1Type(IS_CONST)
		LITERAL_STR(opline.GetOp1(), lcname.Copy())

		/* RTD key is placed after lcname literal in op1 */

		ZendAddLiteralString(&key)

		/* RTD key is placed after lcname literal in op1 */

	}
	types.ZendStringReleaseEx(lcname, 0)
}
func ZendCompileFuncDecl(result *Znode, ast *ZendAst, toplevel types.ZendBool) {
	var decl *ZendAstDecl = (*ZendAstDecl)(ast)
	var params_ast *ZendAst = decl.GetChild()[0]
	var uses_ast *ZendAst = decl.GetChild()[1]
	var stmt_ast *ZendAst = decl.GetChild()[2]
	var return_type_ast *ZendAst = decl.GetChild()[3]
	var is_method types.ZendBool = decl.GetKind() == ZEND_AST_METHOD
	var orig_class_entry *types.ClassEntry = CG__().GetActiveClassEntry()
	var orig_op_array *ZendOpArray = CG__().GetActiveOpArray()
	var op_array *ZendOpArray = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_op_array"))
	var orig_oparray_context ZendOparrayContext
	var info ClosureInfo
	memset(&info, 0, b.SizeOf("closure_info"))
	InitOpArray(op_array, ZEND_USER_FUNCTION, INITIAL_OP_ARRAY_SIZE)
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_PRELOAD) != 0 {
		op_array.SetIsPreloaded(true)
		ZEND_MAP_PTR_NEW(op_array.run_time_cache)
		ZEND_MAP_PTR_NEW(op_array.static_variables_ptr)
	} else {
		ZEND_MAP_PTR_INIT(op_array.run_time_cache, ZendArenaAlloc(CG__().GetArena(), b.SizeOf("void *")))
		ZEND_MAP_PTR_SET(op_array.run_time_cache, nil)
	}
	op_array.AddFnFlags(orig_op_array.GetFnFlags() & ZEND_ACC_STRICT_TYPES)
	op_array.AddFnFlags(decl.GetFlags())
	op_array.SetLineStart(decl.GetStartLineno())
	op_array.SetLineEnd(decl.GetEndLineno())
	if decl.GetDocComment() != nil {
		op_array.SetDocComment(decl.GetDocComment().Copy())
	}
	if decl.GetKind() == ZEND_AST_CLOSURE || decl.GetKind() == ZEND_AST_ARROW_FUNC {
		op_array.SetIsClosure(true)
	}
	if is_method != 0 {
		var has_body types.ZendBool = stmt_ast != nil
		ZendBeginMethodDecl(op_array, decl.GetName(), has_body)
	} else {
		ZendBeginFuncDecl(result, op_array, decl, toplevel)
		if decl.GetKind() == ZEND_AST_ARROW_FUNC {
			FindImplicitBinds(&info, params_ast, stmt_ast)
			CompileImplicitLexicalBinds(&info, result, op_array)
		} else if uses_ast != nil {
			ZendCompileClosureBinding(result, op_array, uses_ast)
		}
	}
	CG__().SetActiveOpArray(op_array)

	/* Do not leak the class scope into free standing functions, even if they are dynamically
	 * defined inside a class method. This is necessary for correct handling of magic constants.
	 * For example __CLASS__ should always be "" inside a free standing function. */

	if decl.GetKind() == ZEND_AST_FUNC_DECL {
		CG__().SetActiveClassEntry(nil)
	}
	if toplevel != 0 {
		op_array.SetIsTopLevel(true)
	}
	ZendOparrayContextBegin(&orig_oparray_context)
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_EXTENDED_STMT) != 0 {
		var opline_ext *ZendOp = ZendEmitOp(nil, ZEND_EXT_NOP, nil, nil)
		opline_ext.SetLineno(decl.GetStartLineno())
	}

	/* Push a separator to the loop variable stack */

	var dummy_var ZendLoopVar
	dummy_var.SetOpcode(ZEND_RETURN)
	CG__().GetLoopVarStack().Push(any(&dummy_var))
	ZendCompileParams(params_ast, return_type_ast)
	if CG__().GetActiveOpArray().IsGenerator() {
		ZendMarkFunctionAsGenerator()
		ZendEmitOp(nil, ZEND_GENERATOR_CREATE, nil, nil)
	}
	if decl.GetKind() == ZEND_AST_ARROW_FUNC {
		ZendCompileImplicitClosureUses(&info)
		info.GetUses().Destroy()
	} else if uses_ast != nil {
		ZendCompileClosureUses(uses_ast)
	}
	ZendCompileStmt(stmt_ast)
	if is_method != 0 {
		ZendCheckMagicMethodImplementation(CG__().GetActiveClassEntry(), (*ZendFunction)(op_array), faults.E_COMPILE_ERROR)
	}

	/* put the implicit return on the really last line */

	CG__().SetZendLineno(decl.GetEndLineno())
	ZendDoExtendedStmt()
	ZendEmitFinalReturn(0)
	PassTwo(CG__().GetActiveOpArray())
	ZendOparrayContextEnd(&orig_oparray_context)

	/* Pop the loop variable stack separator */

	CG__().GetLoopVarStack().DelTop()
	CG__().SetActiveOpArray(orig_op_array)
	CG__().SetActiveClassEntry(orig_class_entry)
}
func ZendCompilePropDecl(ast *ZendAst, type_ast *ZendAst, flags uint32) {
	var list *ZendAstList = ZendAstGetList(ast)
	var ce *types.ClassEntry = CG__().GetActiveClassEntry()
	var i uint32
	var children uint32 = list.GetChildren()
	if ce.IsInterface() {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Interfaces may not include member variables")
	}
	if (flags & ZEND_ACC_ABSTRACT) != 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Properties cannot be declared abstract")
	}
	for i = 0; i < children; i++ {
		var prop_ast *ZendAst = list.GetChild()[i]
		var name_ast *ZendAst = prop_ast.GetChild()[0]
		var value_ast *ZendAst = prop_ast.GetChild()[1]
		var doc_comment_ast *ZendAst = prop_ast.GetChild()[2]
		var name *types.String = ZvalMakeInternedString(ZendAstGetZval(name_ast))
		var doc_comment *types.String = nil
		var value_zv types.Zval
		var type_ types.ZendType = 0
		if type_ast != nil {
			type_ = ZendCompileTypename(type_ast, 0)
			if type_.Code() == types.IS_VOID || type_.Code() == types.IS_CALLABLE {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Property %s::$%s cannot have type %s", ce.GetName().GetVal(), name.GetVal(), types.ZendGetTypeByConst(type_.Code()))
			}
		}

		/* Doc comment has been appended as last element in ZEND_AST_PROP_ELEM ast */

		if doc_comment_ast != nil {
			doc_comment = ZendAstGetStr(doc_comment_ast).Copy()
		}
		if (flags & ZEND_ACC_FINAL) != 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare property %s::$%s final, "+"the final modifier is allowed only for methods and classes", ce.GetName().GetVal(), name.GetVal())
		}
		if ce.GetPropertiesInfo().KeyExistsname.GetStr()) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot redeclare %s::$%s", ce.GetName().GetVal(), name.GetVal())
		}
		if value_ast != nil {
			ZendConstExprToZval(&value_zv, value_ast)
			if type_.IsSet() && !(value_zv.IsConstant()) {
				if value_zv.IsNull() {
					if !(type_.AllowNull()) {
						var name *byte = b.CondF(type_.IsClass(), func() []byte { return types.ZEND_TYPE_NAME(type_).GetVal() }, func() *byte { return types.ZendGetTypeByConst(type_.Code()) })
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for property of type %s may not be null. "+"Use the nullable type ?%s to allow null default value", name, name)
					}
				} else if type_.IsClass() {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Property of type %s may not have default value", types.ZEND_TYPE_NAME(type_).GetVal())
				} else if type_.Code() == types.IS_ARRAY || type_.Code() == types.IS_ITERABLE {
					if value_zv.GetType() != types.IS_ARRAY {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for property of type %s can only be an array", types.ZendGetTypeByConst(type_.Code()))
					}
				} else if type_.Code() == types.IS_DOUBLE {
					if value_zv.GetType() != types.IS_DOUBLE && value_zv.GetType() != types.IS_LONG {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for property of type float can only be float or int")
					}
					ConvertToDouble(&value_zv)
				} else if !(types.ZEND_SAME_FAKE_TYPE(type_.Code(), value_zv.GetType())) {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for property of type %s can only be %s", types.ZendGetTypeByConst(type_.Code()), types.ZendGetTypeByConst(type_.Code()))
				}
			}
		} else if !(type_.IsSet()) {
			value_zv.SetNull()
		} else {
			value_zv.SetUndef()
		}
		ZendDeclareTypedProperty(ce, name, &value_zv, flags, doc_comment, type_)
	}
}
func ZendCompilePropGroup(list *ZendAst) {
	var type_ast *ZendAst = list.GetChild()[0]
	var prop_ast *ZendAst = list.GetChild()[1]
	ZendCompilePropDecl(prop_ast, type_ast, list.GetAttr())
}
func ZendCompileClassConstDecl(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var ce *types.ClassEntry = CG__().GetActiveClassEntry()
	var i uint32
	if ce.IsTrait() {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Traits cannot have constants")
		return
	}
	for i = 0; i < list.GetChildren(); i++ {
		var const_ast *ZendAst = list.GetChild()[i]
		var name_ast *ZendAst = const_ast.GetChild()[0]
		var value_ast *ZendAst = const_ast.GetChild()[1]
		var doc_comment_ast *ZendAst = const_ast.GetChild()[2]
		var name *types.String = ZvalMakeInternedString(ZendAstGetZval(name_ast))
		var doc_comment *types.String = b.CondF1(doc_comment_ast != nil, func() *types.String { return ZendAstGetStr(doc_comment_ast).Copy() }, nil)
		var value_zv types.Zval
		if (ast.GetAttr() & (ZEND_ACC_STATIC | ZEND_ACC_ABSTRACT | ZEND_ACC_FINAL)) != 0 {
			if (ast.GetAttr() & ZEND_ACC_STATIC) != 0 {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use 'static' as constant modifier")
			} else if (ast.GetAttr() & ZEND_ACC_ABSTRACT) != 0 {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use 'abstract' as constant modifier")
			} else if (ast.GetAttr() & ZEND_ACC_FINAL) != 0 {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use 'final' as constant modifier")
			}
		}
		ZendConstExprToZval(&value_zv, value_ast)
		ZendDeclareClassConstantEx(ce, name, &value_zv, ast.GetAttr(), doc_comment)
	}
}
func ZendCompileMethodRef(ast *ZendAst, method_ref *ZendTraitMethodReference) {
	var class_ast *ZendAst = ast.GetChild()[0]
	var method_ast *ZendAst = ast.GetChild()[1]
	method_ref.SetMethodName(ZendAstGetStr(method_ast).Copy())
	if class_ast != nil {
		method_ref.SetClassName(ZendResolveClassNameAst(class_ast))
	} else {
		method_ref.SetClassName(nil)
	}
}
func ZendCompileTraitPrecedence(ast *ZendAst) {
	var method_ref_ast *ZendAst = ast.GetChild()[0]
	var insteadof_ast *ZendAst = ast.GetChild()[1]
	var insteadof_list *ZendAstList = ZendAstGetList(insteadof_ast)
	var i uint32
	var precedence *ZendTraitPrecedence = Emalloc(b.SizeOf("zend_trait_precedence") + (insteadof_list.GetChildren()-1)*b.SizeOf("zend_string *"))
	ZendCompileMethodRef(method_ref_ast, precedence.GetTraitMethod())
	precedence.SetNumExcludes(insteadof_list.GetChildren())
	for i = 0; i < insteadof_list.GetChildren(); i++ {
		var name_ast *ZendAst = insteadof_list.GetChild()[i]
		precedence.GetExcludeClassNames()[i] = ZendResolveClassNameAst(name_ast)
	}
	ZendAddToList(CG__().GetActiveClassEntry().GetTraitPrecedences(), precedence)
}
func ZendCompileTraitAlias(ast *ZendAst) {
	var method_ref_ast *ZendAst = ast.GetChild()[0]
	var alias_ast *ZendAst = ast.GetChild()[1]
	var modifiers uint32 = ast.GetAttr()
	var alias *ZendTraitAlias
	if modifiers == ZEND_ACC_STATIC {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use 'static' as method modifier")
	} else if modifiers == ZEND_ACC_ABSTRACT {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use 'abstract' as method modifier")
	} else if modifiers == ZEND_ACC_FINAL {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use 'final' as method modifier")
	}
	alias = Emalloc(b.SizeOf("zend_trait_alias"))
	ZendCompileMethodRef(method_ref_ast, alias.GetTraitMethod())
	alias.SetModifiers(modifiers)
	if alias_ast != nil {
		alias.SetAlias(ZendAstGetStr(alias_ast).Copy())
	} else {
		alias.SetAlias(nil)
	}
	ZendAddToList(CG__().GetActiveClassEntry().GetTraitAliases(), alias)
}
func ZendCompileUseTrait(ast *ZendAst) {
	var traits *ZendAstList = ZendAstGetList(ast.GetChild()[0])
	var adaptations *ZendAstList = b.CondF1(ast.GetChild()[1] != nil, func() *ZendAstList { return ZendAstGetList(ast.GetChild()[1]) }, nil)
	var ce *types.ClassEntry = CG__().GetActiveClassEntry()
	var i uint32
	ce.SetIsImplementTraits(true)
	ce.SetTraitNames(Erealloc(ce.GetTraitNames(), b.SizeOf("zend_class_name")*(ce.GetNumTraits()+traits.GetChildren())))
	for i = 0; i < traits.GetChildren(); i++ {
		var trait_ast *ZendAst = traits.GetChild()[i]
		var name *types.String = ZendAstGetStr(trait_ast)
		if ce.IsInterface() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use traits inside of interfaces. "+"%s is used in %s", name.GetVal(), ce.GetName().GetVal())
		}
		switch ZendGetClassFetchType(name) {
		case ZEND_FETCH_CLASS_SELF:
			fallthrough
		case ZEND_FETCH_CLASS_PARENT:
			fallthrough
		case ZEND_FETCH_CLASS_STATIC:
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use '%s' as trait name "+"as it is reserved", name.GetVal())
		}
		ce.GetTraitNames()[ce.GetNumTraits()].SetName(ZendResolveClassNameAst(trait_ast))
		ce.GetTraitNames()[ce.GetNumTraits()].SetLcName(ZendStringTolower(ce.GetTraitNames()[ce.GetNumTraits()].GetName()))
		ce.GetNumTraits()++
	}
	if adaptations == nil {
		return
	}
	for i = 0; i < adaptations.GetChildren(); i++ {
		var adaptation_ast *ZendAst = adaptations.GetChild()[i]
		switch adaptation_ast.GetKind() {
		case ZEND_AST_TRAIT_PRECEDENCE:
			ZendCompileTraitPrecedence(adaptation_ast)
		case ZEND_AST_TRAIT_ALIAS:
			ZendCompileTraitAlias(adaptation_ast)
		default:

		}
	}
}
func ZendCompileImplements(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var ce *types.ClassEntry = CG__().GetActiveClassEntry()
	var interface_names *ZendClassName
	var i uint32
	interface_names = Emalloc(b.SizeOf("zend_class_name") * list.GetChildren())
	for i = 0; i < list.GetChildren(); i++ {
		var class_ast *ZendAst = list.GetChild()[i]
		var name *types.String = ZendAstGetStr(class_ast)
		if ZendIsConstDefaultClassRef(class_ast) == 0 {
			Efree(interface_names)
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use '%s' as interface name as it is reserved", name.GetVal())
		}
		interface_names[i].SetName(ZendResolveClassNameAst(class_ast))
		interface_names[i].SetLcName(ZendStringTolower(interface_names[i].GetName()))
	}
	ce.SetIsImplementInterfaces(true)
	ce.SetNumInterfaces(list.GetChildren())
	ce.SetInterfaceNames(interface_names)
}
func ZendGenerateAnonClassName(start_lineno uint32) *types.String {
	var filename *types.String = CG__().GetActiveOpArray().GetFilename()
	var result *types.String = ZendStrpprintf(0, "class@anonymous%c%s:%"+"u"+"$%"+PRIx32, '0', filename.GetVal(), start_lineno, b.PostInc(&(CG__().GetRtdKeyCounter())))
	return types.ZendNewInternedString(result)
}
func ZendCompileClassDecl(ast *ZendAst, toplevel types.ZendBool) *ZendOp {
	var decl *ZendAstDecl = (*ZendAstDecl)(ast)
	var extends_ast *ZendAst = decl.GetChild()[0]
	var implements_ast *ZendAst = decl.GetChild()[1]
	var stmt_ast *ZendAst = decl.GetChild()[2]
	var name *types.String
	var lcname *types.String
	var ce *types.ClassEntry = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_class_entry"))
	var opline *ZendOp
	var original_ce *types.ClassEntry = CG__().GetActiveClassEntry()
	if !decl.IsAnonClass() {
		var unqualified_name *types.String = decl.GetName()
		if CG__().GetActiveClassEntry() != nil {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class declarations may not be nested")
		}
		ZendAssertValidClassName(unqualified_name)
		name = ZendPrefixWithNs(unqualified_name)
		name = types.ZendNewInternedString(name)
		lcname = ZendStringTolower(name)
		if FC__().GetImports() != nil {
			var import_name *types.String = ZendHashFindPtrLc(FC__().GetImports(), unqualified_name.GetVal(), unqualified_name.GetLen())
			if import_name != nil && !(types.ZendStringEqualsCi(lcname, import_name)) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare class %s "+"because the name is already in use", name.GetVal())
			}
		}
		ZendRegisterSeenSymbol(lcname, ZEND_SYMBOL_CLASS)
	} else {

		/* Find an anon class name that is not in use yet. */

		name = nil
		lcname = nil
		for {
			ZendTmpStringRelease(name)
			ZendTmpStringRelease(lcname)
			name = ZendGenerateAnonClassName(decl.GetStartLineno())
			lcname = ZendStringTolower(name)
			if !CG__().GetClassTable().KeyExistslcname.GetStr()) {
				break
			}
		}
	}
	lcname = types.ZendNewInternedString(lcname)
	ce.SetType(ZEND_USER_CLASS)
	ce.SetName(name)
	ZendInitializeClassData(ce, 1)
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_PRELOAD) != 0 {
		ce.SetIsPreloaded(true)
		ZEND_MAP_PTR_NEW(ce.static_members_table)
	}
	ce.AddCeFlags(decl.GetFlags())
	ce.SetFilename(ZendGetCompiledFilename())
	ce.SetLineStart(decl.GetStartLineno())
	ce.SetLineEnd(decl.GetEndLineno())
	if decl.GetDocComment() != nil {
		ce.SetDocComment(decl.GetDocComment().Copy())
	}
	if decl.IsAnonClass() {

		/* Serialization is not supported for anonymous classes */

		ce.SetSerialize(ZendClassSerializeDeny)
		ce.SetUnserialize(ZendClassUnserializeDeny)
	}
	if extends_ast != nil {
		var extends_node Znode
		var extends_name *types.String
		if ZendIsConstDefaultClassRef(extends_ast) == 0 {
			extends_name = ZendAstGetStr(extends_ast)
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use '%s' as class name as it is reserved", extends_name.GetVal())
		}
		ZendCompileExpr(&extends_node, extends_ast)
		if extends_node.GetOpType() != IS_CONST || extends_node.GetConstant().GetType() != types.IS_STRING {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Illegal class name")
		}
		extends_name = extends_node.GetConstant().GetStr()
		ce.SetParentName(ZendResolveClassName(extends_name, b.CondF1(extends_ast.GetKind() == ZEND_AST_ZVAL, func() ZendAstAttr { return extends_ast.GetAttr() }, ZEND_NAME_FQ)))
		types.ZendStringReleaseEx(extends_name, 0)
		ce.SetIsInherited(true)
	}
	CG__().SetActiveClassEntry(ce)
	ZendCompileStmt(stmt_ast)

	/* Reset lineno for final opcodes and errors */

	CG__().SetZendLineno(ast.GetLineno())
	if !ce.IsImplementTraits() {

		/* For traits this check is delayed until after trait binding */

		ZendCheckDeprecatedConstructor(ce)

		/* For traits this check is delayed until after trait binding */

	}
	if ce.GetConstructor() != nil {
		ce.GetConstructor().SetIsCtor(true)
		if ce.GetConstructor().IsStatic() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Constructor %s::%s() cannot be static", ce.GetName().GetVal(), ce.GetConstructor().GetFunctionName().GetVal())
		}
		if ce.GetConstructor().IsHasReturnType() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Constructor %s::%s() cannot declare a return type", ce.GetName().GetVal(), ce.GetConstructor().GetFunctionName().GetVal())
		}
	}
	if ce.GetDestructor() != nil {
		ce.GetDestructor().SetIsDtor(true)
		if ce.GetDestructor().IsStatic() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Destructor %s::%s() cannot be static", ce.GetName().GetVal(), ce.GetDestructor().GetFunctionName().GetVal())
		} else if ce.GetDestructor().IsHasReturnType() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Destructor %s::%s() cannot declare a return type", ce.GetName().GetVal(), ce.GetDestructor().GetFunctionName().GetVal())
		}
	}
	if ce.GetClone() != nil {
		if ce.GetClone().IsStatic() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Clone method %s::%s() cannot be static", ce.GetName().GetVal(), ce.GetClone().GetFunctionName().GetVal())
		} else if ce.GetClone().IsHasReturnType() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Clone method %s::%s() cannot declare a return type", ce.GetName().GetVal(), ce.GetClone().GetFunctionName().GetVal())
		}
	}
	if implements_ast != nil {
		ZendCompileImplements(implements_ast)
	}
	if (ce.GetCeFlags() & (ZEND_ACC_IMPLICIT_ABSTRACT_CLASS | ZEND_ACC_INTERFACE | ZEND_ACC_TRAIT | ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)) == ZEND_ACC_IMPLICIT_ABSTRACT_CLASS {
		ZendVerifyAbstractClass(ce)
	}
	CG__().SetActiveClassEntry(original_ce)
	if toplevel != 0 {
		ce.SetIsTopLevel(true)
	}
	if toplevel != 0 && !ce.HasCeFlags(ZEND_ACC_IMPLEMENT_INTERFACES|ZEND_ACC_IMPLEMENT_TRAITS) && (CG__().GetCompilerOptions()&ZEND_COMPILE_PRELOAD) == 0 {
		if extends_ast != nil {
			var parent_ce *types.ClassEntry = ZendLookupClassEx(ce.GetParentName(), nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if parent_ce != nil && (parent_ce.GetType() != ZEND_INTERNAL_CLASS || (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_CLASSES) == 0) && (parent_ce.GetType() != ZEND_USER_CLASS || (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_OTHER_FILES) == 0 || parent_ce.GetFilename() == ce.GetFilename()) {
				CG__().SetZendLineno(decl.GetEndLineno())
				if ZendTryEarlyBind(ce, parent_ce, lcname, nil) != 0 {
					CG__().SetZendLineno(ast.GetLineno())
					types.ZendStringRelease(lcname)
					return nil
				}
				CG__().SetZendLineno(ast.GetLineno())
			}
		} else {
			if types.ZendHashAddPtr(CG__().GetClassTable(), lcname.GetStr(), ce) != nil {
				types.ZendStringRelease(lcname)
				ZendBuildPropertiesInfoTable(ce)
				ce.SetIsLinked(true)
				return nil
			}
		}
	}
	opline = GetNextOp()
	if ce.GetParentName() {

		/* Lowercased parent name */

		var lc_parent_name *types.String = ZendStringTolower(ce.GetParentName())
		opline.SetOp2Type(IS_CONST)
		LITERAL_STR(opline.GetOp2(), lc_parent_name)
	}
	opline.SetOp1Type(IS_CONST)
	LITERAL_STR(opline.GetOp1(), lcname)
	if decl.IsAnonClass() {
		opline.SetOpcode(ZEND_DECLARE_ANON_CLASS)
		opline.SetExtendedValue(ZendAllocCacheSlot())
		opline.SetResultType(IS_VAR)
		opline.GetResult().SetVar(GetTemporaryVariable())
		if !(types.ZendHashAddPtr(CG__().GetClassTable(), lcname.GetStr(), ce)) {

			/* We checked above that the class name is not used. This really shouldn't happen. */

			faults.ErrorNoreturn(faults.E_ERROR, "Runtime definition key collision for %s. This is a bug", name.GetVal())

			/* We checked above that the class name is not used. This really shouldn't happen. */

		}
	} else {

		/* Generate RTD keys until we find one that isn't in use yet. */

		var key *types.String = nil
		for {
			ZendTmpStringRelease(key)
			key = ZendBuildRuntimeDefinitionKey(lcname, decl.GetStartLineno())
			if types.ZendHashAddPtr(CG__().GetClassTable(), key.GetStr(), ce) {
				break
			}
		}

		/* RTD key is placed after lcname literal in op1 */

		ZendAddLiteralString(&key)
		opline.SetOpcode(ZEND_DECLARE_CLASS)
		if extends_ast != nil && toplevel != 0 && (CG__().GetCompilerOptions()&ZEND_COMPILE_DELAYED_BINDING) != 0 && !ce.HasCeFlags(ZEND_ACC_IMPLEMENT_INTERFACES|ZEND_ACC_IMPLEMENT_TRAITS) {
			CG__().GetActiveOpArray().SetIsEarlyBinding(true)
			opline.SetOpcode(ZEND_DECLARE_CLASS_DELAYED)
			opline.SetExtendedValue(ZendAllocCacheSlot())
			opline.SetResultType(IS_UNUSED)
			opline.GetResult().SetOplineNum(-1)
		}
	}
	return opline
}
