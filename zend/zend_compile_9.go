package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func (compiler *Compiler) CompileMagicConst(result *Znode, ast *ZendAst) {
	var opline *types.ZendOp
	if ZendTryCtEvalMagicConst(result.GetConstant(), ast) != 0 {
		result.SetOpType(IS_CONST)
		return
	}
	b.Assert(ast.GetAttr() == T_CLASS_C && CG__().GetActiveClassEntry() != nil && CG__().GetActiveClassEntry().IsTrait())
	opline = ZendEmitOpTmp(result, ZEND_FETCH_CLASS_NAME, nil, nil)
	opline.GetOp1().SetNum(ZEND_FETCH_CLASS_SELF)
}
func ZendIsAllowedInConstExpr(kind ZendAstKind) bool {
	return kind == ZEND_AST_ZVAL || kind == ZEND_AST_BINARY_OP || kind == ZEND_AST_GREATER || kind == ZEND_AST_GREATER_EQUAL || kind == ZEND_AST_AND || kind == ZEND_AST_OR || kind == ZEND_AST_UNARY_OP || kind == ZEND_AST_UNARY_PLUS || kind == ZEND_AST_UNARY_MINUS || kind == ZEND_AST_CONDITIONAL || kind == ZEND_AST_DIM || kind == ZEND_AST_ARRAY || kind == ZEND_AST_ARRAY_ELEM || kind == ZEND_AST_UNPACK || kind == ZEND_AST_CONST || kind == ZEND_AST_CLASS_CONST || kind == ZEND_AST_CLASS_NAME || kind == ZEND_AST_MAGIC_CONST || kind == ZEND_AST_COALESCE
}
func (compiler *Compiler) CompileConstExprClassConst(astPtr **ZendAst) {
	var ast *ZendAst = *astPtr
	var classAst *ZendAst = ast.GetChild()[0]
	var constAst *ZendAst = ast.GetChild()[1]
	var className *types.String
	var constName *types.String = ZendAstGetStr(constAst)
	var fetchType int
	if classAst.GetKind() != ZEND_AST_ZVAL {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Dynamic class names are not allowed in compile-time class constant references")
	}
	className = ZendAstGetStr(classAst)
	fetchType = ZendGetClassFetchType(className.GetStr())
	if ZEND_FETCH_CLASS_STATIC == fetchType {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "\"static::\" is not allowed in compile-time constants")
	}
	if ZEND_FETCH_CLASS_DEFAULT == fetchType {
		className = ZendResolveClassNameAst(classAst)
	}
	name := className.GetStr() + "::" + constName.GetStr()
	// ZendAstDestroy(ast)
	*astPtr = ZendAstCreateConstant(types.NewString(name), fetchType|ZEND_FETCH_CLASS_EXCEPTION)
}
func (compiler *Compiler) CompileConstExprClassName(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var class_ast *ZendAst = ast.GetChild()[0]
	var class_name *types.String = ZendAstGetStr(class_ast)
	var fetch_type uint32 = ZendGetClassFetchType(class_name.GetStr())
	switch fetch_type {
	case ZEND_FETCH_CLASS_SELF:
		fallthrough
	case ZEND_FETCH_CLASS_PARENT:

		/* For the const-eval representation store the fetch type instead of the name. */

		// types.ZendStringRelease(class_name)
		ast.GetChild()[0] = nil
		ast.SetAttr(fetch_type)
		return
	case ZEND_FETCH_CLASS_STATIC:
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "static::class cannot be used for compile-time class name resolution")
		return
	default:

	}
}
func (compiler *Compiler) CompileConstExprConst(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var name_ast *ZendAst = ast.GetChild()[0]
	var orig_name = ZendAstGetStrVal(name_ast)
	var result types.Zval
	resolved_name, isFullyQualified := ZendResolveConstName(orig_name, name_ast.GetAttr())
	if ZendTryCtEvalConst(&result, resolved_name, isFullyQualified) != 0 {
		// types.ZendStringReleaseEx(resolved_name, 0)
		// ZendAstDestroy(ast)
		*ast_ptr = ZendAstCreateZval(&result)
		return
	}
	// ZendAstDestroy(ast)
	*ast_ptr = ZendAstCreateConstant(resolved_name, b.Cond(!isFullyQualified, IS_CONSTANT_UNQUALIFIED, 0))
}
func (compiler *Compiler) CompileConstExprMagicConst(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr

	/* Other cases already resolved by constant folding */

	b.Assert(ast.GetAttr() == T_CLASS_C)
	// ZendAstDestroy(ast)
	*ast_ptr = AstCreate(ZEND_AST_CONSTANT_CLASS)
}
func (compiler *Compiler) CompileConstExpr(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	if ast == nil || ast.GetKind() == ZEND_AST_ZVAL {
		return
	}
	if !ZendIsAllowedInConstExpr(ast.GetKind()) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Constant expression contains invalid operations")
	}
	switch ast.GetKind() {
	case ZEND_AST_CLASS_CONST:
		compiler.CompileConstExprClassConst(ast_ptr)
	case ZEND_AST_CLASS_NAME:
		compiler.CompileConstExprClassName(ast_ptr)
	case ZEND_AST_CONST:
		compiler.CompileConstExprConst(ast_ptr)
	case ZEND_AST_MAGIC_CONST:
		compiler.CompileConstExprMagicConst(ast_ptr)
	default:
		ZendAstApply(ast, compiler.CompileConstExpr)
	}
}
func (compiler *Compiler) ConstExprToZval(result *types.Zval, ast *ZendAst) {
	var orig_ast *ZendAst = ast
	compiler.EvalConstExpr(&ast)
	compiler.CompileConstExpr(&ast)
	if ast.GetKind() == ZEND_AST_ZVAL {
		types.ZVAL_COPY_VALUE(result, ZendAstGetZval(ast))
	} else {
		result.SetConstantAst(types.NewAstRef(ast))

		/* destroy the ast here, it might have been replaced */
		// ZendAstDestroy(ast)
	}

	/* Kill this branch of the original AST, as it was already destroyed.
	 * It would be nice to find a better solution to this problem in the
	 * future. */

	orig_ast.SetKind(0)
}
func (compiler *Compiler) CompileTopStmt(ast *ZendAst) {
	if ast == nil {
		return
	}
	if ast.GetKind() == ZEND_AST_STMT_LIST {
		var list *ZendAstList = ast.AsAstList()
		var i uint32
		for i = 0; i < list.GetChildren(); i++ {
			compiler.CompileTopStmt(list.GetChild()[i])
		}
		return
	}
	if ast.GetKind() == ZEND_AST_FUNC_DECL {
		compiler.setLinenoByAst(ast)
		compiler.CompileFuncDecl(nil, ast, 1)
		compiler.setLinenoByDeclEnd((*ZendAstDecl)(ast))
	} else if ast.GetKind() == ZEND_AST_CLASS {
		compiler.setLinenoByAst(ast)
		compiler.CompileClassDecl(ast, 1)
		compiler.setLinenoByDeclEnd((*ZendAstDecl)(ast))
	} else {
		compiler.CompileStmt(ast)
	}
	if ast.GetKind() != ZEND_AST_NAMESPACE && ast.GetKind() != ZEND_AST_HALT_COMPILER {
		ZendVerifyNamespace()
	}
}
func (compiler *Compiler) CompileStmt(ast *ZendAst) {
	if ast == nil {
		return
	}
	compiler.setLinenoByAst(ast)
	if (CG__().GetCompilerOptions()&ZEND_COMPILE_EXTENDED_STMT) != 0 && ZendIsUntickedStmt(ast) == 0 {
		ZendDoExtendedStmt()
	}
	switch ast.GetKind() {
	case ZEND_AST_STMT_LIST:
		compiler.CompileStmtList(ast)
	case ZEND_AST_GLOBAL:
		compiler.CompileGlobalVar(ast)
	case ZEND_AST_STATIC:
		compiler.CompileStaticVar(ast)
	case ZEND_AST_UNSET:
		compiler.CompileUnset(ast)
	case ZEND_AST_RETURN:
		compiler.CompileReturn(ast)
	case ZEND_AST_ECHO:
		compiler.CompileEcho(ast)
	case ZEND_AST_THROW:
		compiler.CompileThrow(ast)
	case ZEND_AST_BREAK:
		fallthrough
	case ZEND_AST_CONTINUE:
		compiler.CompileBreakContinue(ast)
	case ZEND_AST_GOTO:
		compiler.CompileGoto(ast)
	case ZEND_AST_LABEL:
		compiler.CompileLabel(ast)
	case ZEND_AST_WHILE:
		compiler.CompileWhile(ast)
	case ZEND_AST_DO_WHILE:
		compiler.CompileDoWhile(ast)
	case ZEND_AST_FOR:
		compiler.CompileFor(ast)
	case ZEND_AST_FOREACH:
		compiler.CompileForeach(ast)
	case ZEND_AST_IF:
		compiler.CompileIf(ast)
	case ZEND_AST_SWITCH:
		compiler.CompileSwitch(ast)
	case ZEND_AST_TRY:
		compiler.CompileTry(ast)
	case ZEND_AST_DECLARE:
		compiler.CompileDeclare(ast)
	case ZEND_AST_FUNC_DECL:
		fallthrough
	case ZEND_AST_METHOD:
		compiler.CompileFuncDecl(nil, ast, 0)
	case ZEND_AST_PROP_GROUP:
		compiler.CompilePropGroup(ast)
	case ZEND_AST_CLASS_CONST_DECL:
		compiler.CompileClassConstDecl(ast)
	case ZEND_AST_USE_TRAIT:
		compiler.CompileUseTrait(ast)
	case ZEND_AST_CLASS:
		compiler.CompileClassDecl(ast, 0)
	case ZEND_AST_GROUP_USE:
		compiler.CompileGroupUse(ast)
	case ZEND_AST_USE:
		compiler.CompileUse(ast)
	case ZEND_AST_CONST_DECL:
		compiler.CompileConstDecl(ast)
	case ZEND_AST_NAMESPACE:
		compiler.CompileNamespace(ast)
	case ZEND_AST_HALT_COMPILER:
		compiler.CompileHaltCompiler(ast)
	default:
		var result Znode
		compiler.CompileExpr(&result, ast)
		ZendDoFree(&result)
	}
}

func (compiler *Compiler) CompileExpr(result *Znode, ast *ZendAst) {
	/* CG(zend_lineno) = ast->lineno; */

	compiler.setLinenoByAstEx(ast)
	if CG__().GetMemoizeMode() != ZEND_MEMOIZE_NONE {
		compiler.CompileMemoizedExpr(result, ast)
		return
	}
	switch ast.GetKind() {
	case ZEND_AST_ZVAL:
		types.ZVAL_COPY(result.GetConstant(), ZendAstGetZval(ast))
		result.SetOpType(IS_CONST)
		return
	case ZEND_AST_ZNODE:
		*result = (*ZendAstGetZnode)(ast)
		return
	case ZEND_AST_VAR:
		fallthrough
	case ZEND_AST_DIM:
		fallthrough
	case ZEND_AST_PROP:
		fallthrough
	case ZEND_AST_STATIC_PROP:
		fallthrough
	case ZEND_AST_CALL:
		fallthrough
	case ZEND_AST_METHOD_CALL:
		fallthrough
	case ZEND_AST_STATIC_CALL:
		compiler.CompileVar(result, ast, BP_VAR_R, 0)
		return
	case ZEND_AST_ASSIGN:
		compiler.CompileAssign(result, ast)
		return
	case ZEND_AST_ASSIGN_REF:
		compiler.CompileAssignRef(result, ast)
		return
	case ZEND_AST_NEW:
		compiler.CompileNew(result, ast)
		return
	case ZEND_AST_CLONE:
		compiler.CompileClone(result, ast)
		return
	case ZEND_AST_ASSIGN_OP:
		compiler.CompileCompoundAssign(result, ast)
		return
	case ZEND_AST_BINARY_OP:
		compiler.CompileBinaryOp(result, ast)
		return
	case ZEND_AST_GREATER:
		fallthrough
	case ZEND_AST_GREATER_EQUAL:
		compiler.CompileGreater(result, ast)
		return
	case ZEND_AST_UNARY_OP:
		compiler.CompileUnaryOp(result, ast)
		return
	case ZEND_AST_UNARY_PLUS:
		fallthrough
	case ZEND_AST_UNARY_MINUS:
		compiler.CompileUnaryPm(result, ast)
		return
	case ZEND_AST_AND:
		fallthrough
	case ZEND_AST_OR:
		compiler.CompileShortCircuiting(result, ast)
		return
	case ZEND_AST_POST_INC:
		fallthrough
	case ZEND_AST_POST_DEC:
		compiler.CompilePostIncdec(result, ast)
		return
	case ZEND_AST_PRE_INC:
		fallthrough
	case ZEND_AST_PRE_DEC:
		compiler.CompilePreIncdec(result, ast)
		return
	case ZEND_AST_CAST:
		compiler.CompileCast(result, ast)
		return
	case ZEND_AST_CONDITIONAL:
		compiler.CompileConditional(result, ast)
		return
	case ZEND_AST_COALESCE:
		compiler.CompileCoalesce(result, ast)
		return
	case ZEND_AST_ASSIGN_COALESCE:
		compiler.CompileAssignCoalesce(result, ast)
		return
	case ZEND_AST_PRINT:
		compiler.CompilePrint(result, ast)
		return
	case ZEND_AST_EXIT:
		compiler.CompileExit(result, ast)
		return
	case ZEND_AST_YIELD:
		compiler.CompileYield(result, ast)
		return
	case ZEND_AST_YIELD_FROM:
		compiler.CompileYieldFrom(result, ast)
		return
	case ZEND_AST_INSTANCEOF:
		compiler.CompileInstanceof(result, ast)
		return
	case ZEND_AST_INCLUDE_OR_EVAL:
		compiler.CompileIncludeOrEval(result, ast)
		return
	case ZEND_AST_ISSET:
		fallthrough
	case ZEND_AST_EMPTY:
		compiler.CompileIssetOrEmpty(result, ast)
		return
	case ZEND_AST_SILENCE:
		compiler.CompileSilence(result, ast)
		return
	case ZEND_AST_SHELL_EXEC:
		compiler.CompileShellExec(result, ast)
		return
	case ZEND_AST_ARRAY:
		compiler.CompileArray(result, ast)
		return
	case ZEND_AST_CONST:
		compiler.CompileConst(result, ast)
		return
	case ZEND_AST_CLASS_CONST:
		compiler.CompileClassConst(result, ast)
		return
	case ZEND_AST_CLASS_NAME:
		compiler.CompileClassName(result, ast)
		return
	case ZEND_AST_ENCAPS_LIST:
		compiler.CompileEncapsList(result, ast)
		return
	case ZEND_AST_MAGIC_CONST:
		compiler.CompileMagicConst(result, ast)
		return
	case ZEND_AST_CLOSURE:
		fallthrough
	case ZEND_AST_ARROW_FUNC:
		compiler.CompileFuncDecl(result, ast, 0)
		return
	default:
		b.Assert(false)
	}
}
func (compiler *Compiler) CompileVar(result *Znode, ast *ZendAst, type_ uint32, by_ref int) *types.ZendOp {
	compiler.setLinenoByAstEx(ast)
	switch ast.GetKind() {
	case ZEND_AST_VAR:
		return compiler.CompileSimpleVar(result, ast, type_, 0)
	case ZEND_AST_DIM:
		return compiler.CompileDim(result, ast, type_)
	case ZEND_AST_PROP:
		return compiler.CompileProp(result, ast, type_, by_ref)
	case ZEND_AST_STATIC_PROP:
		return compiler.CompileStaticProp(result, ast, type_, by_ref, 0)
	case ZEND_AST_CALL:
		compiler.CompileCall(result, ast, type_)
		return nil
	case ZEND_AST_METHOD_CALL:
		compiler.CompileMethodCall(result, ast, type_)
		return nil
	case ZEND_AST_STATIC_CALL:
		compiler.CompileStaticCall(result, ast, type_)
		return nil
	case ZEND_AST_ZNODE:
		*result = (*ZendAstGetZnode)(ast)
		return nil
	default:
		if type_ == BP_VAR_W || type_ == BP_VAR_RW || type_ == BP_VAR_UNSET {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use temporary expression in write context")
		}
		compiler.CompileExpr(result, ast)
		return nil
	}
}
func (compiler *Compiler) DelayedCompileVar(result *Znode, ast *ZendAst, type_ uint32, by_ref bool) *types.ZendOp {
	switch ast.GetKind() {
	case ZEND_AST_VAR:
		return compiler.CompileSimpleVar(result, ast, type_, 1)
	case ZEND_AST_DIM:
		return compiler.DelayedCompileDim(result, ast, type_)
	case ZEND_AST_PROP:
		var opline *types.ZendOp = compiler.DelayedCompileProp(result, ast, type_)
		if by_ref != 0 {
			opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_REF)
		}
		return opline
	case ZEND_AST_STATIC_PROP:
		return compiler.CompileStaticProp(result, ast, type_, by_ref, 1)
	default:
		return compiler.CompileVar(result, ast, type_, 0)
	}
}

func (compiler *Compiler) EvalConstExpr(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var result types.Zval
	if ast == nil {
		return
	}
	switch ast.GetKind() {
	case ZEND_AST_BINARY_OP:
		compiler.EvalConstExpr(ast.GetChild()[0])
		compiler.EvalConstExpr(ast.GetChild()[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		if ZendTryCtEvalBinaryOp(&result, ast.GetAttr(), ZendAstGetZval(ast.GetChild()[0]), ZendAstGetZval(ast.GetChild()[1])) == 0 {
			return
		}
	case ZEND_AST_GREATER:
		fallthrough
	case ZEND_AST_GREATER_EQUAL:
		compiler.EvalConstExpr(ast.GetChild()[0])
		compiler.EvalConstExpr(ast.GetChild()[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		ZendCtEvalGreater(&result, ast.GetKind(), ZendAstGetZval(ast.GetChild()[0]), ZendAstGetZval(ast.GetChild()[1]))
	case ZEND_AST_AND:
		fallthrough
	case ZEND_AST_OR:
		var child0_is_true bool
		var child1_is_true bool
		compiler.EvalConstExpr(ast.GetChild()[0])
		compiler.EvalConstExpr(ast.GetChild()[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
			return
		}
		child0_is_true = operators.IZendIsTrue(ZendAstGetZval(ast.GetChild()[0]))
		if child0_is_true == (ast.GetKind() == ZEND_AST_OR) {
			result.SetBool(ast.GetKind() == ZEND_AST_OR)
			break
		}
		if ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		child1_is_true = operators.IZendIsTrue(ZendAstGetZval(ast.GetChild()[1]))
		if ast.GetKind() == ZEND_AST_OR {
			result.SetBool(child0_is_true != 0 || child1_is_true != 0)
		} else {
			result.SetBool(child0_is_true != 0 && child1_is_true != 0)
		}
	case ZEND_AST_UNARY_OP:
		compiler.EvalConstExpr(ast.GetChild()[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
			return
		}
		ZendCtEvalUnaryOp(&result, ast.GetAttr(), ZendAstGetZval(ast.GetChild()[0]))
	case ZEND_AST_UNARY_PLUS:
		fallthrough
	case ZEND_AST_UNARY_MINUS:
		compiler.EvalConstExpr(ast.GetChild()[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
			return
		}
		if ZendTryCtEvalUnaryPm(&result, ast.GetKind(), ZendAstGetZval(ast.GetChild()[0])) == 0 {
			return
		}
	case ZEND_AST_COALESCE:

		/* Set isset fetch indicator here, opcache disallows runtime altering of the AST */

		if ast.GetChild()[0].GetKind() == ZEND_AST_DIM {
			ast.GetChild()[0].SetAttr(ast.GetChild()[0].GetAttr() | ZEND_DIM_IS)
		}
		compiler.EvalConstExpr(ast.GetChild()[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {

			/* ensure everything was compile-time evaluated at least once */

			compiler.EvalConstExpr(ast.GetChild()[1])
			return
		}
		if ZendAstGetZval(ast.GetChild()[0]).IsNull() {
			compiler.EvalConstExpr(ast.GetChild()[1])
			*ast_ptr = ast.GetChild()[1]
			ast.GetChild()[1] = nil
			// ZendAstDestroy(ast)
		} else {
			*ast_ptr = ast.GetChild()[0]
			ast.GetChild()[0] = nil
			// ZendAstDestroy(ast)
		}
		return
	case ZEND_AST_CONDITIONAL:
		var child **ZendAst
		var child_ast **ZendAst
		compiler.EvalConstExpr(ast.GetChild()[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {

			/* ensure everything was compile-time evaluated at least once */

			if ast.GetChild()[1] != nil {
				compiler.EvalConstExpr(ast.GetChild()[1])
			}
			compiler.EvalConstExpr(ast.GetChild()[2])
			return
		}
		child = ast.GetChild()[2-operators.IZendIsTrue(ZendAstGetZval(ast.GetChild()[0]))]
		if (*child) == nil {
			child--
		}
		child_ast = *child
		*child = nil
		// ZendAstDestroy(ast)
		*ast_ptr = child_ast
		compiler.EvalConstExpr(ast_ptr)
		return
	case ZEND_AST_DIM:

		/* constant expression should be always read context ... */

		var container *types.Zval
		var dim *types.Zval
		if ast.GetChild()[1] == nil {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use [] for reading")
		}
		if (ast.GetAttr() & ZEND_DIM_ALTERNATIVE_SYNTAX) != 0 {
			ast.SetAttr(ast.GetAttr() &^ ZEND_DIM_ALTERNATIVE_SYNTAX)
			faults.Error(faults.E_DEPRECATED, "Array and string offset access syntax with curly braces is deprecated")
		}

		/* Set isset fetch indicator here, opcache disallows runtime altering of the AST */

		if (ast.GetAttr()&ZEND_DIM_IS) != 0 && ast.GetChild()[0].GetKind() == ZEND_AST_DIM {
			ast.GetChild()[0].SetAttr(ast.GetChild()[0].GetAttr() | ZEND_DIM_IS)
		}
		compiler.EvalConstExpr(ast.GetChild()[0])
		compiler.EvalConstExpr(ast.GetChild()[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		container = ZendAstGetZval(ast.GetChild()[0])
		dim = ZendAstGetZval(ast.GetChild()[1])
		if container.IsArray() {
			var el *types.Zval
			if dim.IsLong() {
				el = container.Array().IndexFind(dim.Long())
				if el != nil {
					types.ZVAL_COPY(&result, el)
				} else {
					return
				}
			} else if dim.IsString() {
				el = container.Array().SymtableFind(dim.String().GetStr())
				if el != nil {
					types.ZVAL_COPY(&result, el)
				} else {
					return
				}
			} else {
				return
			}
		} else if container.IsString() {
			var offset ZendLong
			var c uint8
			if dim.IsLong() {
				offset = dim.Long()
			} else if !dim.IsString() || operators.IsNumericString(dim.String().GetStr(), &offset, nil, 1) != types.IS_LONG {
				return
			}
			if offset < 0 || int(offset >= container.String().GetLen()) != 0 {
				return
			}
			c = uint8(container.String().GetStr()[offset])
			result.SetStringVal(string(c))
		} else if container.IsSignFalse() {
			result.SetNull()
		} else {
			return
		}
	case ZEND_AST_ARRAY:
		if compiler.TryCtEvalArray(&result, ast) == 0 {
			return
		}
	case ZEND_AST_MAGIC_CONST:
		if ZendTryCtEvalMagicConst(&result, ast) == 0 {
			return
		}
	case ZEND_AST_CONST:
		var name_ast *ZendAst = ast.GetChild()[0]
		var resolved_name, isFullyQualified = ZendResolveConstName(ZendAstGetStr(name_ast), name_ast.GetAttr())
		if ZendTryCtEvalConst(&result, resolved_name, isFullyQualified) == 0 {
			return
		}
	case ZEND_AST_CLASS_CONST:
		var class_ast *ZendAst
		var name_ast *ZendAst
		var resolved_name *types.String
		compiler.EvalConstExpr(ast.GetChild()[0])
		compiler.EvalConstExpr(ast.GetChild()[1])
		class_ast = ast.GetChild()[0]
		name_ast = ast.GetChild()[1]
		if class_ast.GetKind() != ZEND_AST_ZVAL || name_ast.GetKind() != ZEND_AST_ZVAL {
			return
		}
		resolved_name = ZendResolveClassNameAst(class_ast)
		if ZendTryCtEvalClassConst(&result, resolved_name, ZendAstGetStr(name_ast)) == 0 {
			// types.ZendStringReleaseEx(resolved_name, 0)
			return
		}
		// types.ZendStringReleaseEx(resolved_name, 0)
	case ZEND_AST_CLASS_NAME:
		var class_ast *ZendAst = ast.GetChild()[0]
		if ZendTryCompileConstExprResolveClassName(&result, class_ast) == 0 {
			return
		}
	default:
		return
	}
	// ZendAstDestroy(ast)
	*ast_ptr = ZendAstCreateZval(&result)
}
