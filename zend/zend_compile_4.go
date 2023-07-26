package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func (compiler *Compiler) CompileFuncGetArgs(result *Znode, args *ZendAstList) int {
	if CG__().GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 0 {
		ZendEmitOpTmp(result, ZEND_FUNC_GET_ARGS, nil, nil)
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func (compiler *Compiler) CompileFuncArrayKeyExists(result *Znode, args *ZendAstList) int {
	var subject Znode
	var needle Znode
	if args.GetChildren() != 2 {
		return types.FAILURE
	}
	compiler.CompileExpr(&needle, args.Children()[0])
	compiler.CompileExpr(&subject, args.Children()[1])
	ZendEmitOpTmp(result, ZEND_ARRAY_KEY_EXISTS, &needle, &subject)
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncArraySlice(result *Znode, args *ZendAstList) int {
	if CG__().GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 2 && args.Children()[0].Kind() == ZEND_AST_CALL && args.Children()[0].Children()[0].Kind() == ZEND_AST_ZVAL && ZendAstGetZval(args.Children()[0].Children()[0]).IsString() && args.Children()[0].Children()[1].Kind() == ZEND_AST_ARG_LIST && args.Children()[1].Kind() == ZEND_AST_ZVAL {
		var orig_name *types.String = ZendAstGetStr(args.Children()[0].Children()[0])
		name, _ := ZendResolveFunctionName(orig_name.GetStr(), args.Children()[0].Children()[0].Attr())
		var list *ZendAstList = args.Children()[0].Children()[1].AsAstList()
		var zv *types.Zval = ZendAstGetZval(args.Children()[1])
		var first Znode
		if ascii.StrCaseEquals(name, "func_get_args") && list.GetChildren() == 0 && zv.IsLong() && zv.Long() >= 0 {
			first.SetOpType(IS_CONST)
			first.GetConstant().SetLong(zv.Long())
			ZendEmitOpTmp(result, ZEND_FUNC_GET_ARGS, &first, nil)
			// types.ZendStringReleaseEx(name, 0)
			return types.SUCCESS
		}
		// types.ZendStringReleaseEx(name, 0)
	}
	return types.FAILURE
}
func (compiler *Compiler) TryCompileSpecialFunc(result *Znode, lcname *types.String, args *ZendAstList, fbc types.IFunction, type_ uint32) int {
	if fbc.GetInternalFunction().GetHandler() == ZifDisplayDisabledFunction {
		return types.FAILURE
	}
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_NO_BUILTINS) != 0 {
		return types.FAILURE
	}
	if ZendArgsContainUnpack(args) != 0 {
		return types.FAILURE
	}
	if lcname.GetStr() == "strlen" {
		return compiler.CompileFuncStrlen(result, args)
	} else if lcname.GetStr() == "is_null" {
		return compiler.CompileFuncTypecheck(result, args, types.IsNull)
	} else if lcname.GetStr() == "is_bool" {
		return compiler.CompileFuncTypecheck(result, args, types.IsBool)
	} else if lcname.GetStr() == "is_long" || lcname.GetStr() == "is_int" || lcname.GetStr() == "is_integer" {
		return compiler.CompileFuncTypecheck(result, args, types.IsLong)
	} else if lcname.GetStr() == "is_float" || lcname.GetStr() == "is_double" {
		return compiler.CompileFuncTypecheck(result, args, types.IsDouble)
	} else if lcname.GetStr() == "is_string" {
		return compiler.CompileFuncTypecheck(result, args, types.IsString)
	} else if lcname.GetStr() == "is_array" {
		return compiler.CompileFuncTypecheck(result, args, types.IsArray)
	} else if lcname.GetStr() == "is_object" {
		return compiler.CompileFuncTypecheck(result, args, types.IsObject)
	} else if lcname.GetStr() == "is_resource" {
		return compiler.CompileFuncTypecheck(result, args, types.IsResource)
	} else if lcname.GetStr() == "boolval" {
		return compiler.CompileFuncCast(result, args, types.IsBool)
	} else if lcname.GetStr() == "intval" {
		return compiler.CompileFuncCast(result, args, types.IsLong)
	} else if lcname.GetStr() == "floatval" || lcname.GetStr() == "doubleval" {
		return compiler.CompileFuncCast(result, args, types.IsDouble)
	} else if lcname.GetStr() == "strval" {
		return compiler.CompileFuncCast(result, args, types.IsString)
	} else if lcname.GetStr() == "defined" {
		return compiler.CompileFuncDefined(result, args)
	} else if lcname.GetStr() == "chr" && type_ == BP_VAR_R {
		return compiler.CompileFuncChr(result, args)
	} else if lcname.GetStr() == "ord" && type_ == BP_VAR_R {
		return compiler.CompileFuncOrd(result, args)
	} else if lcname.GetStr() == "call_user_func_array" {
		return compiler.CompileFuncCufa(result, args, lcname)
	} else if lcname.GetStr() == "call_user_func" {
		return compiler.CompileFuncCuf(result, args, lcname)
	} else if lcname.GetStr() == "in_array" {
		return compiler.CompileFuncInArray(result, args)
	} else if lcname.GetStr() == "count" || lcname.GetStr() == "sizeof" {
		return compiler.CompileFuncCount(result, args, lcname)
	} else if lcname.GetStr() == "get_class" {
		return compiler.CompileFuncGetClass(result, args)
	} else if lcname.GetStr() == "get_called_class" {
		return compiler.CompileFuncGetCalledClass(result, args)
	} else if lcname.GetStr() == "gettype" {
		return compiler.CompileFuncGettype(result, args)
	} else if lcname.GetStr() == "func_num_args" {
		return compiler.CompileFuncNumArgs(result, args)
	} else if lcname.GetStr() == "func_get_args" {
		return compiler.CompileFuncGetArgs(result, args)
	} else if lcname.GetStr() == "array_slice" {
		return compiler.CompileFuncArraySlice(result, args)
	} else if lcname.GetStr() == "array_key_exists" {
		return compiler.CompileFuncArrayKeyExists(result, args)
	} else {
		return types.FAILURE
	}
}
func (compiler *Compiler) CompileCall(result *Znode, ast *ZendAst, type_ uint32) {
	var name_ast *ZendAst = ast.Child(0)
	var args_ast *ZendAst = ast.Child(1)
	var name_node Znode
	if name_ast.Kind() != ZEND_AST_ZVAL || name_ast.Val().GetType() != types.IsString {
		compiler.CompileExpr(&name_node, name_ast)
		compiler.CompileDynamicCall(result, &name_node, args_ast)
		return
	}
	var runtime_resolution = compiler.CompileFunctionName(&name_node, name_ast)
	if runtime_resolution {
		if ascii.StrCaseEquals(ZendAstGetStrVal(name_ast), "assert") {
			compiler.CompileAssert(result, args_ast.AsAstList(), name_node.GetConstant().String(), nil)
		} else {
			compiler.CompileNsCall(result, &name_node, args_ast)
		}
		return
	}
	var name *types.Zval = name_node.GetConstant()
	var lcname *types.String
	var fbc types.IFunction
	var opline *types.ZendOp
	lcname = operators.ZendStringTolower(name.String())

	fbc = CG__().FunctionTable().Get(lcname.GetStr())
	if fbc != nil && lcname.GetStr() == "assert" {
		compiler.CompileAssert(result, args_ast.AsAstList(), lcname, fbc)
		// types.ZendStringRelease(lcname)
		// ZvalPtrDtor(name_node.GetConstant())
		return
	}
	if fbc == nil || FbcIsFinalized(fbc) == 0 || fbc.GetType() == ZEND_INTERNAL_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_USER_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_OTHER_FILES) != 0 && fbc.GetOpArray().GetFilename() != CG__().GetActiveOpArray().GetFilename() {
		// types.ZendStringReleaseEx(lcname, 0)
		compiler.CompileDynamicCall(result, &name_node, args_ast)
		return
	}
	if compiler.TryCompileSpecialFunc(result, lcname, args_ast.AsAstList(), fbc, type_) == types.SUCCESS {
		// types.ZendStringReleaseEx(lcname, 0)
		// ZvalPtrDtor(name_node.GetConstant())
		return
	}
	// ZvalPtrDtor(name_node.GetConstant())
	name_node.GetConstant().SetString(lcname)
	opline = ZendEmitOp(nil, ZEND_INIT_FCALL, nil, &name_node)
	opline.GetResult().SetNum(ZendAllocCacheSlot())
	compiler.CompileCallCommon(result, args_ast, fbc)
}
func (compiler *Compiler) CompileMethodCall(result *Znode, ast *ZendAst, type_ uint32) {
	var obj_ast *ZendAst = ast.Child(0)
	var method_ast *ZendAst = ast.Child(1)
	var args_ast *ZendAst = ast.Children()[2]
	var obj_node Znode
	var method_node Znode
	var opline *types.ZendOp
	var fbc types.IFunction = nil
	if IsThisFetch(obj_ast) {
		obj_node.SetOpType(IS_UNUSED)
		CG__().GetActiveOpArray().SetIsUsesThis(true)
	} else {
		compiler.CompileExpr(&obj_node, obj_ast)
	}
	compiler.CompileExpr(&method_node, method_ast)
	opline = ZendEmitOp(nil, ZEND_INIT_METHOD_CALL, &obj_node, nil)
	if method_node.GetOpType() == IS_CONST {
		if method_node.GetConstant().GetType() != types.IsString {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Method name must be a string")
		}
		opline.SetOp2Type(IS_CONST)
		opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(method_node.GetConstant().String()))
		opline.GetResult().SetNum(ZendAllocCacheSlots(2))
	} else {
		opline.SetOp2Type(method_node.GetOpType())
		if method_node.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(method_node.GetConstant()))
		} else {
			opline.SetOp2(method_node.GetOp())
		}
	}

	/* Check if this calls a known method on $this */

	if opline.GetOp1Type() == IS_UNUSED && opline.GetOp2Type() == IS_CONST && CG__().GetActiveClassEntry() != nil && ZendIsScopeKnown() {
		var lcname *types.String = (CT_CONSTANT(opline.GetOp2()) + 1).GetStr()
		fbc = CG__().GetActiveClassEntry().FunctionTable().Get(lcname.GetStr())
		if fbc != nil && !fbc.HasFnFlags(types.AccPrivate|types.AccFinal) {
			fbc = nil
		}

		/* We only know the exact method that is being called if it is either private or final.
		 * Otherwise an overriding method in a child class may be called. */

	}
	compiler.CompileCallCommon(result, args_ast, fbc)
}
func ZendIsConstructor(name *types.String) bool {
	return ascii.StrCaseEquals(name.GetStr(), ZEND_CONSTRUCTOR_FUNC_NAME)
}
func (compiler *Compiler) CompileStaticCall(result *Znode, ast *ZendAst, type_ uint32) {
	var class_ast *ZendAst = ast.Child(0)
	var method_ast *ZendAst = ast.Child(1)
	var args_ast *ZendAst = ast.Children()[2]
	var class_node Znode
	var method_node Znode
	var opline *types.ZendOp
	var fbc types.IFunction = nil
	compiler.CompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_EXCEPTION)
	compiler.CompileExpr(&method_node, method_ast)
	if method_node.GetOpType() == IS_CONST {
		var name *types.Zval = method_node.GetConstant()
		if !name.IsString() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Method name must be a string")
		}
		if ZendIsConstructor(name.String()) != 0 {
			// ZvalPtrDtor(name)
			method_node.SetOpType(IS_UNUSED)
		}
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_INIT_STATIC_METHOD_CALL)
	ZendSetClassNameOp1(opline, &class_node)
	if method_node.GetOpType() == IS_CONST {
		opline.SetOp2Type(IS_CONST)
		opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(method_node.GetConstant().String()))
		opline.GetResult().SetNum(ZendAllocCacheSlots(2))
	} else {
		if opline.GetOp1Type() == IS_CONST {
			opline.GetResult().SetNum(ZendAllocCacheSlot())
		}
		opline.SetOp2Type(method_node.GetOpType())
		if method_node.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(method_node.GetConstant()))
		} else {
			opline.SetOp2(method_node.GetOp())
		}
	}

	/* Check if we already know which method we're calling */

	if opline.GetOp2Type() == IS_CONST {
		var ce *types.ClassEntry = nil
		if opline.GetOp1Type() == IS_CONST {
			var lcname *types.String = (CT_CONSTANT(opline.GetOp1()) + 1).GetStr()
			ce = CG__().ClassTable().Get(lcname.GetStr())
			if ce == nil && CG__().GetActiveClassEntry() != nil && ascii.StrCaseEquals(CG__().GetActiveClassEntry().GetName().GetStr(), lcname.GetStr()) {
				ce = CG__().GetActiveClassEntry()
			}
		} else if opline.GetOp1Type() == IS_UNUSED && (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_SELF && ZendIsScopeKnown() {
			ce = CG__().GetActiveClassEntry()
		}
		if ce != nil {
			var lcname *types.String = (CT_CONSTANT(opline.GetOp2()) + 1).GetStr()
			fbc = ce.FunctionTable().Get(lcname.GetStr())
			if fbc != nil && !fbc.IsPublic() {
				if ce != CG__().GetActiveClassEntry() && (fbc.IsPrivate() || !fbc.GetScope().IsLinked() || CG__().GetActiveClassEntry() != nil && !CG__().GetActiveClassEntry().IsLinked() || !ZendCheckProtected(ZendGetFunctionRootClass(fbc), CG__().GetActiveClassEntry())) {
					/* incompatibe function */
					fbc = nil
				}
			}
		}
	}
	compiler.CompileCallCommon(result, args_ast, fbc)
}
func (compiler *Compiler) CompileNew(result *Znode, ast *ZendAst) {
	var class_ast *ZendAst = ast.Child(0)
	var args_ast *ZendAst = ast.Child(1)
	var class_node Znode
	var ctor_result Znode
	var opline *types.ZendOp
	if class_ast.Kind() == ZEND_AST_CLASS {

		/* anon class declaration */

		opline = compiler.CompileClassDecl(class_ast, 0)
		class_node.SetOpType(opline.GetResultType())
		class_node.GetOp().SetVar(opline.GetResult().GetVar())
	} else {
		compiler.CompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_EXCEPTION)
	}
	opline = ZendEmitOp(result, ZEND_NEW, nil, nil)
	if class_node.GetOpType() == IS_CONST {
		opline.SetOp1Type(IS_CONST)
		opline.GetOp1().SetConstant(ZendAddClassNameLiteral(class_node.GetConstant().String()))
		opline.GetOp2().SetNum(ZendAllocCacheSlot())
	} else {
		opline.SetOp1Type(class_node.GetOpType())
		if class_node.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(class_node.GetConstant()))
		} else {
			opline.SetOp1(class_node.GetOp())
		}
	}
	compiler.CompileCallCommon(&ctor_result, args_ast, nil)
	ZendDoFree(&ctor_result)
}
func (compiler *Compiler) CompileClone(result *Znode, ast *ZendAst) {
	var obj_ast *ZendAst = ast.Child(0)
	var obj_node Znode
	compiler.CompileExpr(&obj_node, obj_ast)
	ZendEmitOpTmp(result, ZEND_CLONE, &obj_node, nil)
}
func (compiler *Compiler) CompileGlobalVar(ast *ZendAst) {
	var var_ast *ZendAst = ast.Child(0)
	var name_ast *ZendAst = var_ast.Child(0)
	var name_node Znode
	var result Znode
	compiler.CompileExpr(&name_node, name_ast)
	if name_node.GetOpType() == IS_CONST {
		operators.ConvertToString(name_node.GetConstant())
	}
	if IsThisFetch(var_ast) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use $this as global variable")
	} else if ZendTryCompileCv(&result, var_ast) == types.SUCCESS {
		var opline *types.ZendOp = ZendEmitOp(nil, ZEND_BIND_GLOBAL, &result, &name_node)
		opline.SetExtendedValue(ZendAllocCacheSlot())
	} else {

		/* name_ast should be evaluated only. FETCH_GLOBAL_LOCK instructs FETCH_W
		 * to not free the name_node operand, so it can be reused in the following
		 * ASSIGN_REF, which then frees it. */

		var opline *types.ZendOp = ZendEmitOp(&result, ZEND_FETCH_W, &name_node, nil)
		opline.SetExtendedValue(ZEND_FETCH_GLOBAL_LOCK)
		if name_node.GetOpType() == IS_CONST {
			//name_node.GetConstant().String().AddRefcount()
		}
		compiler.EmitAssignRefZnode(AstCreate(ZEND_AST_VAR, ZendAstCreateZnode(&name_node)), &result)
	}
}
func (compiler *Compiler) CompileStaticVarCommon(var_name *types.String, value *types.Zval, mode uint32) {
	var opline *types.ZendOp
	if CG__().GetActiveOpArray().GetStaticVariables() == nil {
		if CG__().GetActiveOpArray().GetScope() != nil {
			CG__().GetActiveOpArray().GetScope().SetIsHasStaticInMethods(true)
		}
		CG__().GetActiveOpArray().SetStaticVariables(types.NewArray(8))
	}

	_, offset := CG__().GetActiveOpArray().GetStaticVariables().KeyUpdateValAndPos(var_name.GetStr(), value)
	if var_name.GetStr() == "this" {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use $this as static variable")
	}
	opline = ZendEmitOp(nil, ZEND_BIND_STATIC, nil, nil)
	opline.SetOp1Type(IS_CV)
	opline.GetOp1().SetVar(LookupCv(var_name))
	opline.SetExtendedValue(offset | mode)
}
func (compiler *Compiler) CompileStaticVar(ast *ZendAst) {
	var var_ast *ZendAst = ast.Child(0)
	var value_ast *ZendAst = ast.Child(1)
	var value_zv types.Zval
	if value_ast != nil {
		compiler.ConstExprToZval(&value_zv, value_ast)
	} else {
		value_zv.SetNull()
	}
	compiler.CompileStaticVarCommon(ZendAstGetStr(var_ast), &value_zv, ZEND_BIND_REF)
}
func (compiler *Compiler) CompileUnset(ast *ZendAst) {
	var var_ast *ZendAst = ast.Child(0)
	var var_node Znode
	var opline *types.ZendOp
	ZendEnsureWritableVariable(var_ast)
	switch var_ast.Kind() {
	case ZEND_AST_VAR:
		if IsThisFetch(var_ast) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot unset $this")
		} else if ZendTryCompileCv(&var_node, var_ast) == types.SUCCESS {
			opline = ZendEmitOp(nil, ZEND_UNSET_CV, &var_node, nil)
		} else {
			opline = compiler.CompileSimpleVarNoCv(nil, var_ast, BP_VAR_UNSET, 0)
			opline.SetOpcode(ZEND_UNSET_VAR)
		}
		return
	case ZEND_AST_DIM:
		opline = compiler.CompileDim(nil, var_ast, BP_VAR_UNSET)
		opline.SetOpcode(ZEND_UNSET_DIM)
		return
	case ZEND_AST_PROP:
		opline = compiler.CompileProp(nil, var_ast, BP_VAR_UNSET, 0)
		opline.SetOpcode(ZEND_UNSET_OBJ)
		return
	case ZEND_AST_STATIC_PROP:
		opline = compiler.CompileStaticProp(nil, var_ast, BP_VAR_UNSET, 0, 0)
		opline.SetOpcode(ZEND_UNSET_STATIC_PROP)
		return
	default:

	}
}
func ZendHandleLoopsAndFinallyEx(depth ZendLong, return_value *Znode) int {
	var base *ZendLoopVar
	var loop_var *ZendLoopVar = CG__().GetLoopVarStack().Top()
	if loop_var == nil {
		return 1
	}
	base = CG__().GetLoopVarStack().GetElements()
	for ; loop_var >= base; loop_var-- {
		if loop_var.GetOpcode() == ZEND_FAST_CALL {
			var opline *types.ZendOp = GetNextOp()
			opline.SetOpcode(ZEND_FAST_CALL)
			opline.SetResultType(IS_TMP_VAR)
			opline.GetResult().SetVar(loop_var.GetVarNum())
			if return_value != nil {
				opline.SetOp2Type(return_value.GetOpType())
				if return_value.GetOpType() == IS_CONST {
					opline.GetOp2().SetConstant(ZendAddLiteral(return_value.GetConstant()))
				} else {
					opline.SetOp2(return_value.GetOp())
				}
			}
			opline.GetOp1().SetNum(loop_var.GetTryCatchOffset())
		} else if loop_var.GetOpcode() == ZEND_DISCARD_EXCEPTION {
			var opline *types.ZendOp = GetNextOp()
			opline.SetOpcode(ZEND_DISCARD_EXCEPTION)
			opline.SetOp1Type(IS_TMP_VAR)
			opline.GetOp1().SetVar(loop_var.GetVarNum())
		} else if loop_var.GetOpcode() == ZEND_RETURN {

			/* Stack separator */

			break

			/* Stack separator */

		} else if depth <= 1 {
			return 1
		} else if loop_var.GetOpcode() == ZEND_NOP {

			/* Loop doesn't have freeable variable */

			depth--

			/* Loop doesn't have freeable variable */

		} else {
			var opline *types.ZendOp
			b.Assert((loop_var.GetVarType() & (IS_VAR | IS_TMP_VAR)) != 0)
			opline = GetNextOp()
			opline.SetOpcode(loop_var.GetOpcode())
			opline.SetOp1Type(loop_var.GetVarType())
			opline.GetOp1().SetVar(loop_var.GetVarNum())
			opline.SetExtendedValue(ZEND_FREE_ON_RETURN)
			depth--
		}
	}
	return depth == 0
}
func ZendHandleLoopsAndFinally(return_value *Znode) int {
	return ZendHandleLoopsAndFinallyEx(CG__().GetLoopVarStack().GetTop()+1, return_value)
}
func ZendHasFinallyEx(depth ZendLong) int {
	var base *ZendLoopVar
	var loop_var *ZendLoopVar = CG__().GetLoopVarStack().Top()
	if loop_var == nil {
		return 0
	}
	base = CG__().GetLoopVarStack().GetElements()
	for ; loop_var >= base; loop_var-- {
		if loop_var.GetOpcode() == ZEND_FAST_CALL {
			return 1
		} else if loop_var.GetOpcode() == ZEND_DISCARD_EXCEPTION {

		} else if loop_var.GetOpcode() == ZEND_RETURN {

			/* Stack separator */

			return 0

			/* Stack separator */

		} else if depth <= 1 {
			return 0
		} else {
			depth--
		}
	}
	return 0
}
func ZendHasFinally() int {
	return ZendHasFinallyEx(CG__().GetLoopVarStack().GetTop() + 1)
}
func (compiler *Compiler) CompileReturn(ast *ZendAst) {
	var expr_ast *ZendAst = ast.Child(0)
	var is_generator bool = CG__().GetActiveOpArray().IsGenerator()
	var by_ref bool = CG__().GetActiveOpArray().IsReturnReference()
	var expr_node Znode
	var opline *types.ZendOp
	if is_generator != 0 {

		/* For generators the by-ref flag refers to yields, not returns */

		by_ref = 0

		/* For generators the by-ref flag refers to yields, not returns */

	}
	if expr_ast == nil {
		expr_node.SetOpType(IS_CONST)
		expr_node.GetConstant().SetNull()
	} else if by_ref != 0 && ZendIsVariable(expr_ast) != 0 {
		compiler.CompileVar(&expr_node, expr_ast, BP_VAR_W, 1)
	} else {
		compiler.CompileExpr(&expr_node, expr_ast)
	}
	if CG__().GetActiveOpArray().IsHasFinallyBlock() && (expr_node.GetOpType() == IS_CV || by_ref != 0 && expr_node.GetOpType() == IS_VAR) && ZendHasFinally() != 0 {

		/* Copy return value into temporary VAR to avoid modification in finally code */

		if by_ref != 0 {
			ZendEmitOp(&expr_node, ZEND_MAKE_REF, &expr_node, nil)
		} else {
			ZendEmitOpTmp(&expr_node, ZEND_QM_ASSIGN, &expr_node, nil)
		}

		/* Copy return value into temporary VAR to avoid modification in finally code */

	}

	/* Generator return types are handled separately */

	if is_generator == 0 && CG__().GetActiveOpArray().IsHasReturnType() {
		ZendEmitReturnTypeCheck(lang.Cond(expr_ast != nil, &expr_node, nil), CG__().GetActiveOpArray().GetArgInfo()-1, 0)
	}
	ZendHandleLoopsAndFinally(lang.Cond((expr_node.GetOpType()&(IS_TMP_VAR|IS_VAR)) != 0, &expr_node, nil))
	opline = ZendEmitOp(nil, lang.Cond(by_ref != 0, ZEND_RETURN_BY_REF, ZEND_RETURN), &expr_node, nil)
	if by_ref != 0 && expr_ast != nil {
		if ZendIsCall(expr_ast) != 0 {
			opline.SetExtendedValue(ZEND_RETURNS_FUNCTION)
		} else if ZendIsVariable(expr_ast) == 0 {
			opline.SetExtendedValue(ZEND_RETURNS_VALUE)
		}
	}
}
func (compiler *Compiler) CompileEcho(ast *ZendAst) {
	var opline *types.ZendOp
	var expr_ast *ZendAst = ast.Child(0)
	var expr_node Znode
	compiler.CompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOp(nil, ZEND_ECHO, &expr_node, nil)
	opline.SetExtendedValue(0)
}
func (compiler *Compiler) CompileThrow(ast *ZendAst) {
	var expr_ast *ZendAst = ast.Child(0)
	var expr_node Znode
	compiler.CompileExpr(&expr_node, expr_ast)
	ZendEmitOp(nil, ZEND_THROW, &expr_node, nil)
}
func (compiler *Compiler) CompileBreakContinue(ast *ZendAst) {
	var depth_ast *ZendAst = ast.Child(0)
	var opline *types.ZendOp
	var depth ZendLong
	b.Assert(ast.Kind() == ZEND_AST_BREAK || ast.Kind() == ZEND_AST_CONTINUE)
	if depth_ast != nil {
		var depth_zv *types.Zval
		if depth_ast.Kind() != ZEND_AST_ZVAL {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'%s' operator with non-integer operand "+"is no longer supported", lang.Cond(ast.Kind() == ZEND_AST_BREAK, "break", "continue"))
		}
		depth_zv = depth_ast.Val()
		if !depth_zv.IsLong() || depth_zv.Long() < 1 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'%s' operator accepts only positive integers", lang.Cond(ast.Kind() == ZEND_AST_BREAK, "break", "continue"))
		}
		depth = depth_zv.Long()
	} else {
		depth = 1
	}
	if CG__().GetContext().GetCurrentBrkCont() == -1 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'%s' not in the 'loop' or 'switch' context", lang.Cond(ast.Kind() == ZEND_AST_BREAK, "break", "continue"))
	} else {
		if ZendHandleLoopsAndFinallyEx(depth, nil) == 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot '%s' "+ZEND_LONG_FMT+" level%s", lang.Cond(ast.Kind() == ZEND_AST_BREAK, "break", "continue"), depth, lang.Cond(depth == 1, "", "s"))
		}
	}
	if ast.Kind() == ZEND_AST_CONTINUE {
		var d int
		var cur int = CG__().GetContext().GetCurrentBrkCont()
		for d = depth - 1; d > 0; d-- {
			cur = CG__().GetContext().GetBrkContArray()[cur].GetParent()
			b.Assert(cur != -1)
		}
		if CG__().GetContext().GetBrkContArray()[cur].GetIsSwitch() != 0 {
			if depth == 1 {
				faults.Error(faults.E_WARNING, "\"continue\" targeting switch is equivalent to \"break\". "+"Did you mean to use \"continue "+ZEND_LONG_FMT+"\"?", depth+1)
			} else {
				faults.Error(faults.E_WARNING, "\"continue "+ZEND_LONG_FMT+"\" targeting switch is equivalent to \"break "+ZEND_LONG_FMT+"\". "+"Did you mean to use \"continue "+ZEND_LONG_FMT+"\"?", depth, depth, depth+1)
			}
		}
	}
	opline = ZendEmitOp(nil, lang.Cond(ast.Kind() == ZEND_AST_BREAK, ZEND_BRK, ZEND_CONT), nil, nil)
	opline.GetOp1().SetNum(CG__().GetContext().GetCurrentBrkCont())
	opline.GetOp2().SetNum(depth)
}
func (compiler *Compiler) ResolveGotoLabel(op_array *types.ZendOpArray, opline *types.ZendOp) {
	var dest *ZendLabel
	var current int
	var remove_oplines int = opline.GetOp1().GetNum()
	var label *types.Zval
	var opnum uint32 = opline - op_array.GetOpcodes()
	label = CT_CONSTANT_EX(op_array, opline.GetOp2().GetConstant())
	if dest = CG__().GetContext().GetLabel(label.StringVal()); dest == nil {
		CG__().SetInCompilation(1)
		CG__().SetActiveOpArray(op_array)
		compiler.setLinenoByOpline(opline)
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'goto' to undefined label '%s'", label.String().GetVal())
	}

	label.SetNull()
	current = opline.GetExtendedValue()
	for ; current != dest.GetBrkCont(); current = CG__().GetContext().GetBrkContArray()[current].GetParent() {
		if current == -1 {
			CG__().SetInCompilation(1)
			CG__().SetActiveOpArray(op_array)
			compiler.setLinenoByOpline(opline)
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'goto' into loop or switch statement is disallowed")
		}
		if CG__().GetContext().GetBrkContArray()[current].GetStart() >= 0 {
			remove_oplines--
		}
	}
	for current = 0; current < op_array.GetLastTryCatch(); current++ {
		var elem *ZendTryCatchElement = op_array.GetTryCatchArray()[current]
		if elem.GetTryOp() > opnum {
			break
		}
		if elem.GetFinallyOp() != 0 && opnum < elem.GetFinallyOp()-1 && (dest.GetOplineNum() > elem.GetFinallyEnd() || dest.GetOplineNum() < elem.GetTryOp()) {
			remove_oplines--
		}
	}
	opline.SetOpcode(ZEND_JMP)
	opline.GetOp1().SetOplineNum(dest.GetOplineNum())
	opline.SetExtendedValue(0)
	opline.SetOp1Type(IS_UNUSED)
	opline.SetOp2Type(IS_UNUSED)
	opline.SetResultType(IS_UNUSED)
	b.Assert(remove_oplines >= 0)
	for lang.PostDec(&remove_oplines) {
		opline--
		opline.SetNop()
		ZendVmSetOpcodeHandler(opline)
	}
}
func (compiler *Compiler) CompileGoto(ast *ZendAst) {
	var label_ast *ZendAst = ast.Child(0)
	var label_node Znode
	var opline *types.ZendOp
	var opnum_start uint32 = GetNextOpNumber()
	compiler.CompileExpr(&label_node, label_ast)

	/* Label resolution and unwinding adjustments happen in pass two. */

	ZendHandleLoopsAndFinally(nil)
	opline = ZendEmitOp(nil, ZEND_GOTO, nil, &label_node)
	opline.GetOp1().SetNum(GetNextOpNumber() - opnum_start - 1)
	opline.SetExtendedValue(CG__().GetContext().GetCurrentBrkCont())
}
func (compiler *Compiler) CompileLabel(ast *ZendAst) {
	var label = ZendAstGetStr(ast.Child(0)).GetStr()
	dest := NewZendLabel(label, CG__().GetContext().GetCurrentBrkCont(), GetNextOpNumber())
	if !CG__().GetContext().AddLabel(dest) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Label '%s' already defined", label)
	}
}
func (compiler *Compiler) CompileWhile(ast *ZendAst) {
	var cond_ast *ZendAst = ast.Child(0)
	var stmt_ast *ZendAst = ast.Child(1)
	var cond_node Znode
	var opnum_start uint32
	var opnum_jmp uint32
	var opnum_cond uint32
	opnum_jmp = ZendEmitJump(0)
	ZendBeginLoop(ZEND_NOP, nil, 0)
	opnum_start = GetNextOpNumber()
	compiler.CompileStmt(stmt_ast)
	opnum_cond = GetNextOpNumber()
	ZendUpdateJumpTarget(opnum_jmp, opnum_cond)
	compiler.CompileExpr(&cond_node, cond_ast)
	ZendEmitCondJump(ZEND_JMPNZ, &cond_node, opnum_start)
	ZendEndLoop(opnum_cond, nil)
}
func (compiler *Compiler) CompileDoWhile(ast *ZendAst) {
	var stmt_ast *ZendAst = ast.Child(0)
	var cond_ast *ZendAst = ast.Child(1)
	var cond_node Znode
	var opnum_start uint32
	var opnum_cond uint32
	ZendBeginLoop(ZEND_NOP, nil, 0)
	opnum_start = GetNextOpNumber()
	compiler.CompileStmt(stmt_ast)
	opnum_cond = GetNextOpNumber()
	compiler.CompileExpr(&cond_node, cond_ast)
	ZendEmitCondJump(ZEND_JMPNZ, &cond_node, opnum_start)
	ZendEndLoop(opnum_cond, nil)
}
func (compiler *Compiler) CompileExprList(result *Znode, ast *ZendAst) {
	var list *ZendAstList
	var i uint32
	result.SetOpType(IS_CONST)
	result.GetConstant().SetTrue()
	if ast == nil {
		return
	}
	list = ast.AsAstList()
	for i = 0; i < list.GetChildren(); i++ {
		var expr_ast *ZendAst = list.Children()[i]
		ZendDoFree(result)
		compiler.CompileExpr(result, expr_ast)
	}
}
func (compiler *Compiler) CompileFor(ast *ZendAst) {
	var init_ast *ZendAst = ast.Children()[0]
	var cond_ast *ZendAst = ast.Child(1)
	var loop_ast *ZendAst = ast.Children()[2]
	var stmt_ast *ZendAst = ast.Children()[3]
	var result Znode
	var opnum_start uint32
	var opnum_jmp uint32
	var opnum_loop uint32
	compiler.CompileExprList(&result, init_ast)
	ZendDoFree(&result)
	opnum_jmp = ZendEmitJump(0)
	ZendBeginLoop(ZEND_NOP, nil, 0)
	opnum_start = GetNextOpNumber()
	compiler.CompileStmt(stmt_ast)
	opnum_loop = GetNextOpNumber()
	compiler.CompileExprList(&result, loop_ast)
	ZendDoFree(&result)
	ZendUpdateJumpTargetToNext(opnum_jmp)
	compiler.CompileExprList(&result, cond_ast)
	ZendDoExtendedStmt()
	ZendEmitCondJump(ZEND_JMPNZ, &result, opnum_start)
	ZendEndLoop(opnum_loop, nil)
}
func (compiler *Compiler) CompileForeach(ast *ZendAst) {
	var expr_ast *ZendAst = ast.Children()[0]
	var value_ast *ZendAst = ast.Child(1)
	var key_ast *ZendAst = ast.Children()[2]
	var stmt_ast *ZendAst = ast.Children()[3]
	var by_ref bool = value_ast.Kind() == ZEND_AST_REF
	var is_variable bool = ZendIsVariable(expr_ast) != 0 && ZendCanWriteToVariable(expr_ast) != 0
	var expr_node Znode
	var reset_node Znode
	var value_node Znode
	var key_node Znode
	var opline *types.ZendOp
	var opnum_reset uint32
	var opnum_fetch uint32
	if key_ast != nil {
		if key_ast.Kind() == ZEND_AST_REF {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Key element cannot be a reference")
		}
		if key_ast.Kind() == ZEND_AST_ARRAY {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use list as key element")
		}
	}
	if by_ref != 0 {
		value_ast = value_ast.Children()[0]
	}
	if value_ast.Kind() == ZEND_AST_ARRAY && ZendPropagateListRefs(value_ast) != 0 {
		by_ref = 1
	}
	if by_ref != 0 && is_variable != 0 {
		compiler.CompileVar(&expr_node, expr_ast, BP_VAR_W, 1)
	} else {
		compiler.CompileExpr(&expr_node, expr_ast)
	}
	if by_ref != 0 {
		ZendSeparateIfCallAndWrite(&expr_node, expr_ast, BP_VAR_W)
	}
	opnum_reset = GetNextOpNumber()
	opline = ZendEmitOp(&reset_node, lang.Cond(by_ref != 0, ZEND_FE_RESET_RW, ZEND_FE_RESET_R), &expr_node, nil)
	ZendBeginLoop(ZEND_FE_FREE, &reset_node, 0)
	opnum_fetch = GetNextOpNumber()
	opline = ZendEmitOp(nil, lang.Cond(by_ref != 0, ZEND_FE_FETCH_RW, ZEND_FE_FETCH_R), &reset_node, nil)
	if IsThisFetch(value_ast) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot re-assign $this")
	} else if value_ast.Kind() == ZEND_AST_VAR && ZendTryCompileCv(&value_node, value_ast) == types.SUCCESS {
		opline.SetOp2Type(value_node.GetOpType())
		if value_node.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(value_node.GetConstant()))
		} else {
			opline.SetOp2(value_node.GetOp())
		}
	} else {
		opline.SetOp2Type(IS_VAR)
		opline.GetOp2().SetVar(GetTemporaryVariable())
		value_node.SetOpType(opline.GetOp2Type())
		if value_node.GetOpType() == IS_CONST {
			types.ZVAL_COPY_VALUE(value_node.GetConstant(), CT_CONSTANT(opline.GetOp2()))
		} else {
			value_node.SetOp(opline.GetOp2())
		}
		if value_ast.Kind() == ZEND_AST_ARRAY {
			compiler.CompileListAssign(nil, value_ast, &value_node, value_ast.Attr())
		} else if by_ref != 0 {
			compiler.EmitAssignRefZnode(value_ast, &value_node)
		} else {
			compiler.EmitAssignZnode(value_ast, &value_node)
		}
	}
	if key_ast != nil {
		opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_fetch]
		ZendMakeTmpResult(&key_node, opline)
		compiler.EmitAssignZnode(key_ast, &key_node)
	}
	compiler.CompileStmt(stmt_ast)

	/* Place JMP and FE_FREE on the line where foreach starts. It would be
	 * better to use the end line, but this information is not available
	 * currently. */

	compiler.setLinenoByAst(ast)
	ZendEmitJump(opnum_fetch)
	opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_reset]
	opline.GetOp2().SetOplineNum(GetNextOpNumber())
	opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_fetch]
	opline.SetExtendedValue(GetNextOpNumber())
	ZendEndLoop(opnum_fetch, &reset_node)
	opline = ZendEmitOp(nil, ZEND_FE_FREE, &reset_node, nil)
}
