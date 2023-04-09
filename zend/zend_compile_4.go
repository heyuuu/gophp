package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZendCompileFuncGetArgs(result *Znode, args *ZendAstList) int {
	if CG__().GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 0 {
		ZendEmitOpTmp(result, ZEND_FUNC_GET_ARGS, nil, nil)
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func ZendCompileFuncArrayKeyExists(result *Znode, args *ZendAstList) int {
	var subject Znode
	var needle Znode
	if args.GetChildren() != 2 {
		return types.FAILURE
	}
	ZendCompileExpr(&needle, args.GetChild()[0])
	ZendCompileExpr(&subject, args.GetChild()[1])
	ZendEmitOpTmp(result, ZEND_ARRAY_KEY_EXISTS, &needle, &subject)
	return types.SUCCESS
}
func ZendCompileFuncArraySlice(result *Znode, args *ZendAstList) int {
	if CG__().GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 2 && args.GetChild()[0].GetKind() == ZEND_AST_CALL && args.GetChild()[0].GetChild()[0].GetKind() == ZEND_AST_ZVAL && ZendAstGetZval(args.GetChild()[0].GetChild()[0]).IsString() && args.GetChild()[0].GetChild()[1].GetKind() == ZEND_AST_ARG_LIST && args.GetChild()[1].GetKind() == ZEND_AST_ZVAL {
		var orig_name *types.String = ZendAstGetStr(args.GetChild()[0].GetChild()[0])
		var is_fully_qualified types.ZendBool
		var name *types.String = ZendResolveFunctionName(orig_name, args.GetChild()[0].GetChild()[0].GetAttr(), &is_fully_qualified)
		var list *ZendAstList = ZendAstGetList(args.GetChild()[0].GetChild()[1])
		var zv *types.Zval = ZendAstGetZval(args.GetChild()[1])
		var first Znode
		if ascii.StrCaseEquals(name.GetStr(), "func_get_args") && list.GetChildren() == 0 && zv.IsLong() && zv.GetLval() >= 0 {
			first.SetOpType(IS_CONST)
			first.GetConstant().SetLong(zv.GetLval())
			ZendEmitOpTmp(result, ZEND_FUNC_GET_ARGS, &first, nil)
			// types.ZendStringReleaseEx(name, 0)
			return types.SUCCESS
		}
		// types.ZendStringReleaseEx(name, 0)
	}
	return types.FAILURE
}
func ZendTryCompileSpecialFunc(result *Znode, lcname *types.String, args *ZendAstList, fbc types.IFunction, type_ uint32) int {
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
		return ZendCompileFuncStrlen(result, args)
	} else if lcname.GetStr() == "is_null" {
		return ZendCompileFuncTypecheck(result, args, types.IS_NULL)
	} else if lcname.GetStr() == "is_bool" {
		return ZendCompileFuncTypecheck(result, args, types.IS_BOOL)
	} else if lcname.GetStr() == "is_long" || lcname.GetStr() == "is_int" || lcname.GetStr() == "is_integer" {
		return ZendCompileFuncTypecheck(result, args, types.IS_LONG)
	} else if lcname.GetStr() == "is_float" || lcname.GetStr() == "is_double" {
		return ZendCompileFuncTypecheck(result, args, types.IS_DOUBLE)
	} else if lcname.GetStr() == "is_string" {
		return ZendCompileFuncTypecheck(result, args, types.IS_STRING)
	} else if lcname.GetStr() == "is_array" {
		return ZendCompileFuncTypecheck(result, args, types.IS_ARRAY)
	} else if lcname.GetStr() == "is_object" {
		return ZendCompileFuncTypecheck(result, args, types.IS_OBJECT)
	} else if lcname.GetStr() == "is_resource" {
		return ZendCompileFuncTypecheck(result, args, types.IS_RESOURCE)
	} else if lcname.GetStr() == "boolval" {
		return ZendCompileFuncCast(result, args, types.IS_BOOL)
	} else if lcname.GetStr() == "intval" {
		return ZendCompileFuncCast(result, args, types.IS_LONG)
	} else if lcname.GetStr() == "floatval" || lcname.GetStr() == "doubleval" {
		return ZendCompileFuncCast(result, args, types.IS_DOUBLE)
	} else if lcname.GetStr() == "strval" {
		return ZendCompileFuncCast(result, args, types.IS_STRING)
	} else if lcname.GetStr() == "defined" {
		return ZendCompileFuncDefined(result, args)
	} else if lcname.GetStr() == "chr" && type_ == BP_VAR_R {
		return ZendCompileFuncChr(result, args)
	} else if lcname.GetStr() == "ord" && type_ == BP_VAR_R {
		return ZendCompileFuncOrd(result, args)
	} else if lcname.GetStr() == "call_user_func_array" {
		return ZendCompileFuncCufa(result, args, lcname)
	} else if lcname.GetStr() == "call_user_func" {
		return ZendCompileFuncCuf(result, args, lcname)
	} else if lcname.GetStr() == "in_array" {
		return ZendCompileFuncInArray(result, args)
	} else if lcname.GetStr() == "count" || lcname.GetStr() == "sizeof" {
		return ZendCompileFuncCount(result, args, lcname)
	} else if lcname.GetStr() == "get_class" {
		return ZendCompileFuncGetClass(result, args)
	} else if lcname.GetStr() == "get_called_class" {
		return ZendCompileFuncGetCalledClass(result, args)
	} else if lcname.GetStr() == "gettype" {
		return ZendCompileFuncGettype(result, args)
	} else if lcname.GetStr() == "func_num_args" {
		return ZendCompileFuncNumArgs(result, args)
	} else if lcname.GetStr() == "func_get_args" {
		return ZendCompileFuncGetArgs(result, args)
	} else if lcname.GetStr() == "array_slice" {
		return ZendCompileFuncArraySlice(result, args)
	} else if lcname.GetStr() == "array_key_exists" {
		return ZendCompileFuncArrayKeyExists(result, args)
	} else {
		return types.FAILURE
	}
}
func ZendCompileCall(result *Znode, ast *ZendAst, type_ uint32) {
	var name_ast *ZendAst = ast.GetChild()[0]
	var args_ast *ZendAst = ast.GetChild()[1]
	var name_node Znode
	if name_ast.GetKind() != ZEND_AST_ZVAL || ZendAstGetZval(name_ast).GetType() != types.IS_STRING {
		ZendCompileExpr(&name_node, name_ast)
		ZendCompileDynamicCall(result, &name_node, args_ast)
		return
	}
	var runtime_resolution types.ZendBool = ZendCompileFunctionName(&name_node, name_ast)
	if runtime_resolution != 0 {
		if ascii.StrCaseEquals(ZendAstGetStr(name_ast).GetStr(), "assert") {
			ZendCompileAssert(result, ZendAstGetList(args_ast), name_node.GetConstant().GetStr(), nil)
		} else {
			ZendCompileNsCall(result, &name_node, args_ast)
		}
		return
	}
	var name *types.Zval = name_node.GetConstant()
	var lcname *types.String
	var fbc types.IFunction
	var opline *ZendOp
	lcname = ZendStringTolower(name.GetStr())

	fbc = CG__().FunctionTable().Get(lcname.GetStr())
	if fbc != nil && lcname.GetStr() == "assert" {
		ZendCompileAssert(result, ZendAstGetList(args_ast), lcname, fbc)
		// types.ZendStringRelease(lcname)
		ZvalPtrDtor(name_node.GetConstant())
		return
	}
	if fbc == nil || FbcIsFinalized(fbc) == 0 || fbc.GetType() == ZEND_INTERNAL_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_USER_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_OTHER_FILES) != 0 && fbc.GetOpArray().GetFilename() != CG__().GetActiveOpArray().GetFilename() {
		// types.ZendStringReleaseEx(lcname, 0)
		ZendCompileDynamicCall(result, &name_node, args_ast)
		return
	}
	if ZendTryCompileSpecialFunc(result, lcname, ZendAstGetList(args_ast), fbc, type_) == types.SUCCESS {
		// types.ZendStringReleaseEx(lcname, 0)
		ZvalPtrDtor(name_node.GetConstant())
		return
	}
	ZvalPtrDtor(name_node.GetConstant())
	name_node.GetConstant().SetString(lcname)
	opline = ZendEmitOp(nil, ZEND_INIT_FCALL, nil, &name_node)
	opline.GetResult().SetNum(ZendAllocCacheSlot())
	ZendCompileCallCommon(result, args_ast, fbc)
}
func ZendCompileMethodCall(result *Znode, ast *ZendAst, type_ uint32) {
	var obj_ast *ZendAst = ast.GetChild()[0]
	var method_ast *ZendAst = ast.GetChild()[1]
	var args_ast *ZendAst = ast.GetChild()[2]
	var obj_node Znode
	var method_node Znode
	var opline *ZendOp
	var fbc types.IFunction = nil
	if IsThisFetch(obj_ast) {
		obj_node.SetOpType(IS_UNUSED)
		CG__().GetActiveOpArray().SetIsUsesThis(true)
	} else {
		ZendCompileExpr(&obj_node, obj_ast)
	}
	ZendCompileExpr(&method_node, method_ast)
	opline = ZendEmitOp(nil, ZEND_INIT_METHOD_CALL, &obj_node, nil)
	if method_node.GetOpType() == IS_CONST {
		if method_node.GetConstant().GetType() != types.IS_STRING {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Method name must be a string")
		}
		opline.SetOp2Type(IS_CONST)
		opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(method_node.GetConstant().GetStr()))
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

	if opline.GetOp1Type() == IS_UNUSED && opline.GetOp2Type() == IS_CONST && CG__().GetActiveClassEntry() != nil && ZendIsScopeKnown() != 0 {
		var lcname *types.String = (CT_CONSTANT(opline.GetOp2()) + 1).GetStr()
		fbc = CG__().GetActiveClassEntry().FunctionTable().Get(lcname.GetStr())
		if fbc != nil && !fbc.HasFnFlags(AccPrivate|AccFinal) {
			fbc = nil
		}

		/* We only know the exact method that is being called if it is either private or final.
		 * Otherwise an overriding method in a child class may be called. */

	}
	ZendCompileCallCommon(result, args_ast, fbc)
}
func ZendIsConstructor(name *types.String) types.ZendBool {
	return ascii.StrCaseEquals(name.GetStr(), ZEND_CONSTRUCTOR_FUNC_NAME)
}
func ZendCompileStaticCall(result *Znode, ast *ZendAst, type_ uint32) {
	var class_ast *ZendAst = ast.GetChild()[0]
	var method_ast *ZendAst = ast.GetChild()[1]
	var args_ast *ZendAst = ast.GetChild()[2]
	var class_node Znode
	var method_node Znode
	var opline *ZendOp
	var fbc types.IFunction = nil
	ZendCompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_EXCEPTION)
	ZendCompileExpr(&method_node, method_ast)
	if method_node.GetOpType() == IS_CONST {
		var name *types.Zval = method_node.GetConstant()
		if name.GetType() != types.IS_STRING {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Method name must be a string")
		}
		if ZendIsConstructor(name.GetStr()) != 0 {
			ZvalPtrDtor(name)
			method_node.SetOpType(IS_UNUSED)
		}
	}
	opline = GetNextOp()
	opline.SetOpcode(ZEND_INIT_STATIC_METHOD_CALL)
	ZendSetClassNameOp1(opline, &class_node)
	if method_node.GetOpType() == IS_CONST {
		opline.SetOp2Type(IS_CONST)
		opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(method_node.GetConstant().GetStr()))
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
		} else if opline.GetOp1Type() == IS_UNUSED && (opline.GetOp1().GetNum()&ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_SELF && ZendIsScopeKnown() != 0 {
			ce = CG__().GetActiveClassEntry()
		}
		if ce != nil {
			var lcname *types.String = (CT_CONSTANT(opline.GetOp2()) + 1).GetStr()
			fbc = ce.FunctionTable().Get(lcname.GetStr())
			if fbc != nil && !fbc.IsPublic() {
				if ce != CG__().GetActiveClassEntry() && (fbc.IsPrivate() || !fbc.GetScope().IsLinked() || CG__().GetActiveClassEntry() != nil && !CG__().GetActiveClassEntry().IsLinked() || ZendCheckProtected(ZendGetFunctionRootClass(fbc), CG__().GetActiveClassEntry()) == 0) {

					/* incompatibe function */

					fbc = nil

					/* incompatibe function */

				}
			}
		}
	}
	ZendCompileCallCommon(result, args_ast, fbc)
}
func ZendCompileNew(result *Znode, ast *ZendAst) {
	var class_ast *ZendAst = ast.GetChild()[0]
	var args_ast *ZendAst = ast.GetChild()[1]
	var class_node Znode
	var ctor_result Znode
	var opline *ZendOp
	if class_ast.GetKind() == ZEND_AST_CLASS {

		/* anon class declaration */

		opline = ZendCompileClassDecl(class_ast, 0)
		class_node.SetOpType(opline.GetResultType())
		class_node.GetOp().SetVar(opline.GetResult().GetVar())
	} else {
		ZendCompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_EXCEPTION)
	}
	opline = ZendEmitOp(result, ZEND_NEW, nil, nil)
	if class_node.GetOpType() == IS_CONST {
		opline.SetOp1Type(IS_CONST)
		opline.GetOp1().SetConstant(ZendAddClassNameLiteral(class_node.GetConstant().GetStr()))
		opline.GetOp2().SetNum(ZendAllocCacheSlot())
	} else {
		opline.SetOp1Type(class_node.GetOpType())
		if class_node.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(class_node.GetConstant()))
		} else {
			opline.SetOp1(class_node.GetOp())
		}
	}
	ZendCompileCallCommon(&ctor_result, args_ast, nil)
	ZendDoFree(&ctor_result)
}
func ZendCompileClone(result *Znode, ast *ZendAst) {
	var obj_ast *ZendAst = ast.GetChild()[0]
	var obj_node Znode
	ZendCompileExpr(&obj_node, obj_ast)
	ZendEmitOpTmp(result, ZEND_CLONE, &obj_node, nil)
}
func ZendCompileGlobalVar(ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var name_ast *ZendAst = var_ast.GetChild()[0]
	var name_node Znode
	var result Znode
	ZendCompileExpr(&name_node, name_ast)
	if name_node.GetOpType() == IS_CONST {
		ConvertToString(name_node.GetConstant())
	}
	if IsThisFetch(var_ast) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use $this as global variable")
	} else if ZendTryCompileCv(&result, var_ast) == types.SUCCESS {
		var opline *ZendOp = ZendEmitOp(nil, ZEND_BIND_GLOBAL, &result, &name_node)
		opline.SetExtendedValue(ZendAllocCacheSlot())
	} else {

		/* name_ast should be evaluated only. FETCH_GLOBAL_LOCK instructs FETCH_W
		 * to not free the name_node operand, so it can be reused in the following
		 * ASSIGN_REF, which then frees it. */

		var opline *ZendOp = ZendEmitOp(&result, ZEND_FETCH_W, &name_node, nil)
		opline.SetExtendedValue(ZEND_FETCH_GLOBAL_LOCK)
		if name_node.GetOpType() == IS_CONST {
			//name_node.GetConstant().GetStr().AddRefcount()
		}
		ZendEmitAssignRefZnode(ZendAstCreate(ZEND_AST_VAR, ZendAstCreateZnode(&name_node)), &result)
	}
}
func ZendCompileStaticVarCommon(var_name *types.String, value *types.Zval, mode uint32) {
	var opline *ZendOp
	if CG__().GetActiveOpArray().GetStaticVariables() == nil {
		if CG__().GetActiveOpArray().GetScope() != nil {
			CG__().GetActiveOpArray().GetScope().SetIsHasStaticInMethods(true)
		}
		CG__().GetActiveOpArray().SetStaticVariables(types.NewArray(8))
	}
	value = CG__().GetActiveOpArray().GetStaticVariables().KeyUpdate(var_name.GetStr(), value)
	if var_name.GetStr() == "this" {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use $this as static variable")
	}
	opline = ZendEmitOp(nil, ZEND_BIND_STATIC, nil, nil)
	opline.SetOp1Type(IS_CV)
	opline.GetOp1().SetVar(LookupCv(var_name))
	opline.SetExtendedValue(uint32((*byte)(value-(*byte)(CG__().GetActiveOpArray().GetStaticVariables().GetArData()))) | mode)
}
func ZendCompileStaticVar(ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var value_ast *ZendAst = ast.GetChild()[1]
	var value_zv types.Zval
	if value_ast != nil {
		ZendConstExprToZval(&value_zv, value_ast)
	} else {
		value_zv.SetNull()
	}
	ZendCompileStaticVarCommon(ZendAstGetStr(var_ast), &value_zv, ZEND_BIND_REF)
}
func ZendCompileUnset(ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var var_node Znode
	var opline *ZendOp
	ZendEnsureWritableVariable(var_ast)
	switch var_ast.GetKind() {
	case ZEND_AST_VAR:
		if IsThisFetch(var_ast) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot unset $this")
		} else if ZendTryCompileCv(&var_node, var_ast) == types.SUCCESS {
			opline = ZendEmitOp(nil, ZEND_UNSET_CV, &var_node, nil)
		} else {
			opline = ZendCompileSimpleVarNoCv(nil, var_ast, BP_VAR_UNSET, 0)
			opline.SetOpcode(ZEND_UNSET_VAR)
		}
		return
	case ZEND_AST_DIM:
		opline = ZendCompileDim(nil, var_ast, BP_VAR_UNSET)
		opline.SetOpcode(ZEND_UNSET_DIM)
		return
	case ZEND_AST_PROP:
		opline = ZendCompileProp(nil, var_ast, BP_VAR_UNSET, 0)
		opline.SetOpcode(ZEND_UNSET_OBJ)
		return
	case ZEND_AST_STATIC_PROP:
		opline = ZendCompileStaticProp(nil, var_ast, BP_VAR_UNSET, 0, 0)
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
			var opline *ZendOp = GetNextOp()
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
			var opline *ZendOp = GetNextOp()
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
			var opline *ZendOp
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
func ZendCompileReturn(ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var is_generator types.ZendBool = CG__().GetActiveOpArray().IsGenerator()
	var by_ref types.ZendBool = CG__().GetActiveOpArray().IsReturnReference()
	var expr_node Znode
	var opline *ZendOp
	if is_generator != 0 {

		/* For generators the by-ref flag refers to yields, not returns */

		by_ref = 0

		/* For generators the by-ref flag refers to yields, not returns */

	}
	if expr_ast == nil {
		expr_node.SetOpType(IS_CONST)
		expr_node.GetConstant().SetNull()
	} else if by_ref != 0 && ZendIsVariable(expr_ast) != 0 {
		ZendCompileVar(&expr_node, expr_ast, BP_VAR_W, 1)
	} else {
		ZendCompileExpr(&expr_node, expr_ast)
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
		ZendEmitReturnTypeCheck(b.Cond(expr_ast != nil, &expr_node, nil), CG__().GetActiveOpArray().GetArgInfo()-1, 0)
	}
	ZendHandleLoopsAndFinally(b.Cond((expr_node.GetOpType()&(IS_TMP_VAR|IS_VAR)) != 0, &expr_node, nil))
	opline = ZendEmitOp(nil, b.Cond(by_ref != 0, ZEND_RETURN_BY_REF, ZEND_RETURN), &expr_node, nil)
	if by_ref != 0 && expr_ast != nil {
		if ZendIsCall(expr_ast) != 0 {
			opline.SetExtendedValue(ZEND_RETURNS_FUNCTION)
		} else if ZendIsVariable(expr_ast) == 0 {
			opline.SetExtendedValue(ZEND_RETURNS_VALUE)
		}
	}
}
func ZendCompileEcho(ast *ZendAst) {
	var opline *ZendOp
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOp(nil, ZEND_ECHO, &expr_node, nil)
	opline.SetExtendedValue(0)
}
func ZendCompileThrow(ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var expr_node Znode
	ZendCompileExpr(&expr_node, expr_ast)
	ZendEmitOp(nil, ZEND_THROW, &expr_node, nil)
}
func ZendCompileBreakContinue(ast *ZendAst) {
	var depth_ast *ZendAst = ast.GetChild()[0]
	var opline *ZendOp
	var depth ZendLong
	b.Assert(ast.GetKind() == ZEND_AST_BREAK || ast.GetKind() == ZEND_AST_CONTINUE)
	if depth_ast != nil {
		var depth_zv *types.Zval
		if depth_ast.GetKind() != ZEND_AST_ZVAL {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'%s' operator with non-integer operand "+"is no longer supported", b.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"))
		}
		depth_zv = ZendAstGetZval(depth_ast)
		if depth_zv.GetType() != types.IS_LONG || depth_zv.GetLval() < 1 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'%s' operator accepts only positive integers", b.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"))
		}
		depth = depth_zv.GetLval()
	} else {
		depth = 1
	}
	if CG__().GetContext().GetCurrentBrkCont() == -1 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'%s' not in the 'loop' or 'switch' context", b.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"))
	} else {
		if ZendHandleLoopsAndFinallyEx(depth, nil) == 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot '%s' "+ZEND_LONG_FMT+" level%s", b.Cond(ast.GetKind() == ZEND_AST_BREAK, "break", "continue"), depth, b.Cond(depth == 1, "", "s"))
		}
	}
	if ast.GetKind() == ZEND_AST_CONTINUE {
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
	opline = ZendEmitOp(nil, b.Cond(ast.GetKind() == ZEND_AST_BREAK, ZEND_BRK, ZEND_CONT), nil, nil)
	opline.GetOp1().SetNum(CG__().GetContext().GetCurrentBrkCont())
	opline.GetOp2().SetNum(depth)
}
func ZendResolveGotoLabel(op_array *types.ZendOpArray, opline *ZendOp) {
	var dest *ZendLabel
	var current int
	var remove_oplines int = opline.GetOp1().GetNum()
	var label *types.Zval
	var opnum uint32 = opline - op_array.GetOpcodes()
	label = CT_CONSTANT_EX(op_array, opline.GetOp2().GetConstant())
	if CG__().GetContext().GetLabels() == nil || b.Assign(&dest, types.ZendHashFindPtr(CG__().GetContext().GetLabels(), label.GetStr().GetStr())) == nil {
		CG__().SetInCompilation(1)
		CG__().SetActiveOpArray(op_array)
		CG__().SetZendLineno(opline.GetLineno())
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "'goto' to undefined label '%s'", label.GetStr().GetVal())
	}

	label.SetNull()
	current = opline.GetExtendedValue()
	for ; current != dest.GetBrkCont(); current = CG__().GetContext().GetBrkContArray()[current].GetParent() {
		if current == -1 {
			CG__().SetInCompilation(1)
			CG__().SetActiveOpArray(op_array)
			CG__().SetZendLineno(opline.GetLineno())
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
	for b.PostDec(&remove_oplines) {
		opline--
		MAKE_NOP(opline)
		ZendVmSetOpcodeHandler(opline)
	}
}
func ZendCompileGoto(ast *ZendAst) {
	var label_ast *ZendAst = ast.GetChild()[0]
	var label_node Znode
	var opline *ZendOp
	var opnum_start uint32 = GetNextOpNumber()
	ZendCompileExpr(&label_node, label_ast)

	/* Label resolution and unwinding adjustments happen in pass two. */

	ZendHandleLoopsAndFinally(nil)
	opline = ZendEmitOp(nil, ZEND_GOTO, nil, &label_node)
	opline.GetOp1().SetNum(GetNextOpNumber() - opnum_start - 1)
	opline.SetExtendedValue(CG__().GetContext().GetCurrentBrkCont())
}
func ZendCompileLabel(ast *ZendAst) {
	var label *types.String = ZendAstGetStr(ast.GetChild()[0])
	var dest ZendLabel
	if CG__().GetContext().GetLabels() == nil {
		ALLOC_HASHTABLE(CG__().GetContext().GetLabels())
		CG__().GetContext().GetLabels() = types.MakeArrayEx(8, LabelPtrDtor, 0)
	}
	dest.SetBrkCont(CG__().GetContext().GetCurrentBrkCont())
	dest.SetOplineNum(GetNextOpNumber())
	if !(types.ZendHashAddMem(CG__().GetContext().GetLabels(), label.GetStr(), &dest, b.SizeOf("zend_label"))) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Label '%s' already defined", label.GetVal())
	}
}
func ZendCompileWhile(ast *ZendAst) {
	var cond_ast *ZendAst = ast.GetChild()[0]
	var stmt_ast *ZendAst = ast.GetChild()[1]
	var cond_node Znode
	var opnum_start uint32
	var opnum_jmp uint32
	var opnum_cond uint32
	opnum_jmp = ZendEmitJump(0)
	ZendBeginLoop(ZEND_NOP, nil, 0)
	opnum_start = GetNextOpNumber()
	ZendCompileStmt(stmt_ast)
	opnum_cond = GetNextOpNumber()
	ZendUpdateJumpTarget(opnum_jmp, opnum_cond)
	ZendCompileExpr(&cond_node, cond_ast)
	ZendEmitCondJump(ZEND_JMPNZ, &cond_node, opnum_start)
	ZendEndLoop(opnum_cond, nil)
}
func ZendCompileDoWhile(ast *ZendAst) {
	var stmt_ast *ZendAst = ast.GetChild()[0]
	var cond_ast *ZendAst = ast.GetChild()[1]
	var cond_node Znode
	var opnum_start uint32
	var opnum_cond uint32
	ZendBeginLoop(ZEND_NOP, nil, 0)
	opnum_start = GetNextOpNumber()
	ZendCompileStmt(stmt_ast)
	opnum_cond = GetNextOpNumber()
	ZendCompileExpr(&cond_node, cond_ast)
	ZendEmitCondJump(ZEND_JMPNZ, &cond_node, opnum_start)
	ZendEndLoop(opnum_cond, nil)
}
func ZendCompileExprList(result *Znode, ast *ZendAst) {
	var list *ZendAstList
	var i uint32
	result.SetOpType(IS_CONST)
	result.GetConstant().SetTrue()
	if ast == nil {
		return
	}
	list = ZendAstGetList(ast)
	for i = 0; i < list.GetChildren(); i++ {
		var expr_ast *ZendAst = list.GetChild()[i]
		ZendDoFree(result)
		ZendCompileExpr(result, expr_ast)
	}
}
func ZendCompileFor(ast *ZendAst) {
	var init_ast *ZendAst = ast.GetChild()[0]
	var cond_ast *ZendAst = ast.GetChild()[1]
	var loop_ast *ZendAst = ast.GetChild()[2]
	var stmt_ast *ZendAst = ast.GetChild()[3]
	var result Znode
	var opnum_start uint32
	var opnum_jmp uint32
	var opnum_loop uint32
	ZendCompileExprList(&result, init_ast)
	ZendDoFree(&result)
	opnum_jmp = ZendEmitJump(0)
	ZendBeginLoop(ZEND_NOP, nil, 0)
	opnum_start = GetNextOpNumber()
	ZendCompileStmt(stmt_ast)
	opnum_loop = GetNextOpNumber()
	ZendCompileExprList(&result, loop_ast)
	ZendDoFree(&result)
	ZendUpdateJumpTargetToNext(opnum_jmp)
	ZendCompileExprList(&result, cond_ast)
	ZendDoExtendedStmt()
	ZendEmitCondJump(ZEND_JMPNZ, &result, opnum_start)
	ZendEndLoop(opnum_loop, nil)
}
func ZendCompileForeach(ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var value_ast *ZendAst = ast.GetChild()[1]
	var key_ast *ZendAst = ast.GetChild()[2]
	var stmt_ast *ZendAst = ast.GetChild()[3]
	var by_ref types.ZendBool = value_ast.GetKind() == ZEND_AST_REF
	var is_variable types.ZendBool = ZendIsVariable(expr_ast) != 0 && ZendCanWriteToVariable(expr_ast) != 0
	var expr_node Znode
	var reset_node Znode
	var value_node Znode
	var key_node Znode
	var opline *ZendOp
	var opnum_reset uint32
	var opnum_fetch uint32
	if key_ast != nil {
		if key_ast.GetKind() == ZEND_AST_REF {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Key element cannot be a reference")
		}
		if key_ast.GetKind() == ZEND_AST_ARRAY {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use list as key element")
		}
	}
	if by_ref != 0 {
		value_ast = value_ast.GetChild()[0]
	}
	if value_ast.GetKind() == ZEND_AST_ARRAY && ZendPropagateListRefs(value_ast) != 0 {
		by_ref = 1
	}
	if by_ref != 0 && is_variable != 0 {
		ZendCompileVar(&expr_node, expr_ast, BP_VAR_W, 1)
	} else {
		ZendCompileExpr(&expr_node, expr_ast)
	}
	if by_ref != 0 {
		ZendSeparateIfCallAndWrite(&expr_node, expr_ast, BP_VAR_W)
	}
	opnum_reset = GetNextOpNumber()
	opline = ZendEmitOp(&reset_node, b.Cond(by_ref != 0, ZEND_FE_RESET_RW, ZEND_FE_RESET_R), &expr_node, nil)
	ZendBeginLoop(ZEND_FE_FREE, &reset_node, 0)
	opnum_fetch = GetNextOpNumber()
	opline = ZendEmitOp(nil, b.Cond(by_ref != 0, ZEND_FE_FETCH_RW, ZEND_FE_FETCH_R), &reset_node, nil)
	if IsThisFetch(value_ast) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot re-assign $this")
	} else if value_ast.GetKind() == ZEND_AST_VAR && ZendTryCompileCv(&value_node, value_ast) == types.SUCCESS {
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
		if value_ast.GetKind() == ZEND_AST_ARRAY {
			ZendCompileListAssign(nil, value_ast, &value_node, value_ast.GetAttr())
		} else if by_ref != 0 {
			ZendEmitAssignRefZnode(value_ast, &value_node)
		} else {
			ZendEmitAssignZnode(value_ast, &value_node)
		}
	}
	if key_ast != nil {
		opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_fetch]
		ZendMakeTmpResult(&key_node, opline)
		ZendEmitAssignZnode(key_ast, &key_node)
	}
	ZendCompileStmt(stmt_ast)

	/* Place JMP and FE_FREE on the line where foreach starts. It would be
	 * better to use the end line, but this information is not available
	 * currently. */

	CG__().SetZendLineno(ast.GetLineno())
	ZendEmitJump(opnum_fetch)
	opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_reset]
	opline.GetOp2().SetOplineNum(GetNextOpNumber())
	opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_fetch]
	opline.SetExtendedValue(GetNextOpNumber())
	ZendEndLoop(opnum_fetch, &reset_node)
	opline = ZendEmitOp(nil, ZEND_FE_FREE, &reset_node, nil)
}
