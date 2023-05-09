package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendCompileMagicConst(result *Znode, ast *ZendAst) {
	var opline *ZendOp
	if ZendTryCtEvalMagicConst(result.GetConstant(), ast) != 0 {
		result.SetOpType(IS_CONST)
		return
	}
	b.Assert(ast.GetAttr() == T_CLASS_C && CG__().GetActiveClassEntry() != nil && CG__().GetActiveClassEntry().IsTrait())
	opline = ZendEmitOpTmp(result, ZEND_FETCH_CLASS_NAME, nil, nil)
	opline.GetOp1().SetNum(ZEND_FETCH_CLASS_SELF)
}
func ZendIsAllowedInConstExpr(kind ZendAstKind) types.ZendBool {
	return kind == ZEND_AST_ZVAL || kind == ZEND_AST_BINARY_OP || kind == ZEND_AST_GREATER || kind == ZEND_AST_GREATER_EQUAL || kind == ZEND_AST_AND || kind == ZEND_AST_OR || kind == ZEND_AST_UNARY_OP || kind == ZEND_AST_UNARY_PLUS || kind == ZEND_AST_UNARY_MINUS || kind == ZEND_AST_CONDITIONAL || kind == ZEND_AST_DIM || kind == ZEND_AST_ARRAY || kind == ZEND_AST_ARRAY_ELEM || kind == ZEND_AST_UNPACK || kind == ZEND_AST_CONST || kind == ZEND_AST_CLASS_CONST || kind == ZEND_AST_CLASS_NAME || kind == ZEND_AST_MAGIC_CONST || kind == ZEND_AST_COALESCE
}
func ZendCompileConstExprClassConst(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var class_ast *ZendAst = ast.GetChild()[0]
	var const_ast *ZendAst = ast.GetChild()[1]
	var class_name *types.String
	var const_name *types.String = ZendAstGetStr(const_ast)
	var name *types.String
	var fetch_type int
	if class_ast.GetKind() != ZEND_AST_ZVAL {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Dynamic class names are not allowed in compile-time class constant references")
	}
	class_name = ZendAstGetStr(class_ast)
	fetch_type = ZendGetClassFetchType(class_name.GetStr())
	if ZEND_FETCH_CLASS_STATIC == fetch_type {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "\"static::\" is not allowed in compile-time constants")
	}
	if ZEND_FETCH_CLASS_DEFAULT == fetch_type {
		class_name = ZendResolveClassNameAst(class_ast)
	} else {
		//class_name.AddRefcount()
	}
	name = ZendConcat3(class_name.GetVal(), class_name.GetLen(), "::", 2, const_name.GetVal(), const_name.GetLen())
	ZendAstDestroy(ast)
	// types.ZendStringReleaseEx(class_name, 0)
	*ast_ptr = ZendAstCreateConstant(name, fetch_type|ZEND_FETCH_CLASS_EXCEPTION)
}
func ZendCompileConstExprClassName(ast_ptr **ZendAst) {
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
func ZendCompileConstExprConst(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var name_ast *ZendAst = ast.GetChild()[0]
	var orig_name *types.String = ZendAstGetStr(name_ast)
	var is_fully_qualified types.ZendBool
	var result types.Zval
	var resolved_name *types.String
	resolved_name = ZendResolveConstName(orig_name, name_ast.GetAttr(), &is_fully_qualified)
	if ZendTryCtEvalConst(&result, resolved_name.GetStr(), is_fully_qualified) != 0 {
		// types.ZendStringReleaseEx(resolved_name, 0)
		ZendAstDestroy(ast)
		*ast_ptr = ZendAstCreateZval(&result)
		return
	}
	ZendAstDestroy(ast)
	*ast_ptr = ZendAstCreateConstant(resolved_name, b.Cond(is_fully_qualified == 0, IS_CONSTANT_UNQUALIFIED, 0))
}
func ZendCompileConstExprMagicConst(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr

	/* Other cases already resolved by constant folding */

	b.Assert(ast.GetAttr() == T_CLASS_C)
	ZendAstDestroy(ast)
	*ast_ptr = ZendAstCreate(ZEND_AST_CONSTANT_CLASS)
}
func ZendCompileConstExpr(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	if ast == nil || ast.GetKind() == ZEND_AST_ZVAL {
		return
	}
	if ZendIsAllowedInConstExpr(ast.GetKind()) == 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Constant expression contains invalid operations")
	}
	switch ast.GetKind() {
	case ZEND_AST_CLASS_CONST:
		ZendCompileConstExprClassConst(ast_ptr)
	case ZEND_AST_CLASS_NAME:
		ZendCompileConstExprClassName(ast_ptr)
	case ZEND_AST_CONST:
		ZendCompileConstExprConst(ast_ptr)
	case ZEND_AST_MAGIC_CONST:
		ZendCompileConstExprMagicConst(ast_ptr)
	default:
		ZendAstApply(ast, ZendCompileConstExpr)
	}
}
func ZendConstExprToZval(result *types.Zval, ast *ZendAst) {
	var orig_ast *ZendAst = ast
	ZendEvalConstExpr(&ast)
	ZendCompileConstExpr(&ast)
	if ast.GetKind() == ZEND_AST_ZVAL {
		types.ZVAL_COPY_VALUE(result, ZendAstGetZval(ast))
	} else {
		result.SetConstantAst(types.NewAstRef(ast))

		/* destroy the ast here, it might have been replaced */

		ZendAstDestroy(ast)

		/* destroy the ast here, it might have been replaced */

	}

	/* Kill this branch of the original AST, as it was already destroyed.
	 * It would be nice to find a better solution to this problem in the
	 * future. */

	orig_ast.SetKind(0)

	/* Kill this branch of the original AST, as it was already destroyed.
	 * It would be nice to find a better solution to this problem in the
	 * future. */
}
func ZendCompileTopStmt(ast *ZendAst) {
	if ast == nil {
		return
	}
	if ast.GetKind() == ZEND_AST_STMT_LIST {
		var list *ZendAstList = ZendAstGetList(ast)
		var i uint32
		for i = 0; i < list.GetChildren(); i++ {
			ZendCompileTopStmt(list.GetChild()[i])
		}
		return
	}
	if ast.GetKind() == ZEND_AST_FUNC_DECL {
		CG__().SetZendLineno(ast.GetLineno())
		ZendCompileFuncDecl(nil, ast, 1)
		CG__().SetZendLineno((*ZendAstDecl)(ast).GetEndLineno())
	} else if ast.GetKind() == ZEND_AST_CLASS {
		CG__().SetZendLineno(ast.GetLineno())
		ZendCompileClassDecl(ast, 1)
		CG__().SetZendLineno((*ZendAstDecl)(ast).GetEndLineno())
	} else {
		ZendCompileStmt(ast)
	}
	if ast.GetKind() != ZEND_AST_NAMESPACE && ast.GetKind() != ZEND_AST_HALT_COMPILER {
		ZendVerifyNamespace()
	}
}
func ZendCompileStmt(ast *ZendAst) {
	if ast == nil {
		return
	}
	CG__().SetZendLineno(ast.GetLineno())
	if (CG__().GetCompilerOptions()&ZEND_COMPILE_EXTENDED_STMT) != 0 && ZendIsUntickedStmt(ast) == 0 {
		ZendDoExtendedStmt()
	}
	switch ast.GetKind() {
	case ZEND_AST_STMT_LIST:
		ZendCompileStmtList(ast)
	case ZEND_AST_GLOBAL:
		ZendCompileGlobalVar(ast)
	case ZEND_AST_STATIC:
		ZendCompileStaticVar(ast)
	case ZEND_AST_UNSET:
		ZendCompileUnset(ast)
	case ZEND_AST_RETURN:
		ZendCompileReturn(ast)
	case ZEND_AST_ECHO:
		ZendCompileEcho(ast)
	case ZEND_AST_THROW:
		ZendCompileThrow(ast)
	case ZEND_AST_BREAK:
		fallthrough
	case ZEND_AST_CONTINUE:
		ZendCompileBreakContinue(ast)
	case ZEND_AST_GOTO:
		ZendCompileGoto(ast)
	case ZEND_AST_LABEL:
		ZendCompileLabel(ast)
	case ZEND_AST_WHILE:
		ZendCompileWhile(ast)
	case ZEND_AST_DO_WHILE:
		ZendCompileDoWhile(ast)
	case ZEND_AST_FOR:
		ZendCompileFor(ast)
	case ZEND_AST_FOREACH:
		ZendCompileForeach(ast)
	case ZEND_AST_IF:
		ZendCompileIf(ast)
	case ZEND_AST_SWITCH:
		ZendCompileSwitch(ast)
	case ZEND_AST_TRY:
		ZendCompileTry(ast)
	case ZEND_AST_DECLARE:
		ZendCompileDeclare(ast)
	case ZEND_AST_FUNC_DECL:
		fallthrough
	case ZEND_AST_METHOD:
		ZendCompileFuncDecl(nil, ast, 0)
	case ZEND_AST_PROP_GROUP:
		ZendCompilePropGroup(ast)
	case ZEND_AST_CLASS_CONST_DECL:
		ZendCompileClassConstDecl(ast)
	case ZEND_AST_USE_TRAIT:
		ZendCompileUseTrait(ast)
	case ZEND_AST_CLASS:
		ZendCompileClassDecl(ast, 0)
	case ZEND_AST_GROUP_USE:
		ZendCompileGroupUse(ast)
	case ZEND_AST_USE:
		ZendCompileUse(ast)
	case ZEND_AST_CONST_DECL:
		ZendCompileConstDecl(ast)
	case ZEND_AST_NAMESPACE:
		ZendCompileNamespace(ast)
	case ZEND_AST_HALT_COMPILER:
		ZendCompileHaltCompiler(ast)
	default:
		var result Znode
		ZendCompileExpr(&result, ast)
		ZendDoFree(&result)
	}
}
func ZendCompileExpr(result *Znode, ast *ZendAst) {
	/* CG(zend_lineno) = ast->lineno; */

	CG__().SetZendLineno(ZendAstGetLineno(ast))
	if CG__().GetMemoizeMode() != ZEND_MEMOIZE_NONE {
		ZendCompileMemoizedExpr(result, ast)
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
		ZendCompileVar(result, ast, BP_VAR_R, 0)
		return
	case ZEND_AST_ASSIGN:
		ZendCompileAssign(result, ast)
		return
	case ZEND_AST_ASSIGN_REF:
		ZendCompileAssignRef(result, ast)
		return
	case ZEND_AST_NEW:
		ZendCompileNew(result, ast)
		return
	case ZEND_AST_CLONE:
		ZendCompileClone(result, ast)
		return
	case ZEND_AST_ASSIGN_OP:
		ZendCompileCompoundAssign(result, ast)
		return
	case ZEND_AST_BINARY_OP:
		ZendCompileBinaryOp(result, ast)
		return
	case ZEND_AST_GREATER:
		fallthrough
	case ZEND_AST_GREATER_EQUAL:
		ZendCompileGreater(result, ast)
		return
	case ZEND_AST_UNARY_OP:
		ZendCompileUnaryOp(result, ast)
		return
	case ZEND_AST_UNARY_PLUS:
		fallthrough
	case ZEND_AST_UNARY_MINUS:
		ZendCompileUnaryPm(result, ast)
		return
	case ZEND_AST_AND:
		fallthrough
	case ZEND_AST_OR:
		ZendCompileShortCircuiting(result, ast)
		return
	case ZEND_AST_POST_INC:
		fallthrough
	case ZEND_AST_POST_DEC:
		ZendCompilePostIncdec(result, ast)
		return
	case ZEND_AST_PRE_INC:
		fallthrough
	case ZEND_AST_PRE_DEC:
		ZendCompilePreIncdec(result, ast)
		return
	case ZEND_AST_CAST:
		ZendCompileCast(result, ast)
		return
	case ZEND_AST_CONDITIONAL:
		ZendCompileConditional(result, ast)
		return
	case ZEND_AST_COALESCE:
		ZendCompileCoalesce(result, ast)
		return
	case ZEND_AST_ASSIGN_COALESCE:
		ZendCompileAssignCoalesce(result, ast)
		return
	case ZEND_AST_PRINT:
		ZendCompilePrint(result, ast)
		return
	case ZEND_AST_EXIT:
		ZendCompileExit(result, ast)
		return
	case ZEND_AST_YIELD:
		ZendCompileYield(result, ast)
		return
	case ZEND_AST_YIELD_FROM:
		ZendCompileYieldFrom(result, ast)
		return
	case ZEND_AST_INSTANCEOF:
		ZendCompileInstanceof(result, ast)
		return
	case ZEND_AST_INCLUDE_OR_EVAL:
		ZendCompileIncludeOrEval(result, ast)
		return
	case ZEND_AST_ISSET:
		fallthrough
	case ZEND_AST_EMPTY:
		ZendCompileIssetOrEmpty(result, ast)
		return
	case ZEND_AST_SILENCE:
		ZendCompileSilence(result, ast)
		return
	case ZEND_AST_SHELL_EXEC:
		ZendCompileShellExec(result, ast)
		return
	case ZEND_AST_ARRAY:
		ZendCompileArray(result, ast)
		return
	case ZEND_AST_CONST:
		ZendCompileConst(result, ast)
		return
	case ZEND_AST_CLASS_CONST:
		ZendCompileClassConst(result, ast)
		return
	case ZEND_AST_CLASS_NAME:
		ZendCompileClassName(result, ast)
		return
	case ZEND_AST_ENCAPS_LIST:
		ZendCompileEncapsList(result, ast)
		return
	case ZEND_AST_MAGIC_CONST:
		ZendCompileMagicConst(result, ast)
		return
	case ZEND_AST_CLOSURE:
		fallthrough
	case ZEND_AST_ARROW_FUNC:
		ZendCompileFuncDecl(result, ast, 0)
		return
	default:
		b.Assert(false)
	}
}
func ZendCompileVar(result *Znode, ast *ZendAst, type_ uint32, by_ref int) *ZendOp {
	CG__().SetZendLineno(ZendAstGetLineno(ast))
	switch ast.GetKind() {
	case ZEND_AST_VAR:
		return ZendCompileSimpleVar(result, ast, type_, 0)
	case ZEND_AST_DIM:
		return ZendCompileDim(result, ast, type_)
	case ZEND_AST_PROP:
		return ZendCompileProp(result, ast, type_, by_ref)
	case ZEND_AST_STATIC_PROP:
		return ZendCompileStaticProp(result, ast, type_, by_ref, 0)
	case ZEND_AST_CALL:
		ZendCompileCall(result, ast, type_)
		return nil
	case ZEND_AST_METHOD_CALL:
		ZendCompileMethodCall(result, ast, type_)
		return nil
	case ZEND_AST_STATIC_CALL:
		ZendCompileStaticCall(result, ast, type_)
		return nil
	case ZEND_AST_ZNODE:
		*result = (*ZendAstGetZnode)(ast)
		return nil
	default:
		if type_ == BP_VAR_W || type_ == BP_VAR_RW || type_ == BP_VAR_UNSET {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use temporary expression in write context")
		}
		ZendCompileExpr(result, ast)
		return nil
	}
}
func ZendDelayedCompileVar(result *Znode, ast *ZendAst, type_ uint32, by_ref types.ZendBool) *ZendOp {
	switch ast.GetKind() {
	case ZEND_AST_VAR:
		return ZendCompileSimpleVar(result, ast, type_, 1)
	case ZEND_AST_DIM:
		return ZendDelayedCompileDim(result, ast, type_)
	case ZEND_AST_PROP:
		var opline *ZendOp = ZendDelayedCompileProp(result, ast, type_)
		if by_ref != 0 {
			opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_REF)
		}
		return opline
	case ZEND_AST_STATIC_PROP:
		return ZendCompileStaticProp(result, ast, type_, by_ref, 1)
	default:
		return ZendCompileVar(result, ast, type_, 0)
	}
}
func ZendEvalConstExpr(ast_ptr **ZendAst) {
	var ast *ZendAst = *ast_ptr
	var result types.Zval
	if ast == nil {
		return
	}
	switch ast.GetKind() {
	case ZEND_AST_BINARY_OP:
		ZendEvalConstExpr(ast.GetChild()[0])
		ZendEvalConstExpr(ast.GetChild()[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		if ZendTryCtEvalBinaryOp(&result, ast.GetAttr(), ZendAstGetZval(ast.GetChild()[0]), ZendAstGetZval(ast.GetChild()[1])) == 0 {
			return
		}
	case ZEND_AST_GREATER:
		fallthrough
	case ZEND_AST_GREATER_EQUAL:
		ZendEvalConstExpr(ast.GetChild()[0])
		ZendEvalConstExpr(ast.GetChild()[1])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ast.GetChild()[1].GetKind() != ZEND_AST_ZVAL {
			return
		}
		ZendCtEvalGreater(&result, ast.GetKind(), ZendAstGetZval(ast.GetChild()[0]), ZendAstGetZval(ast.GetChild()[1]))
	case ZEND_AST_AND:
		fallthrough
	case ZEND_AST_OR:
		var child0_is_true types.ZendBool
		var child1_is_true types.ZendBool
		ZendEvalConstExpr(ast.GetChild()[0])
		ZendEvalConstExpr(ast.GetChild()[1])
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
		ZendEvalConstExpr(ast.GetChild()[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
			return
		}
		ZendCtEvalUnaryOp(&result, ast.GetAttr(), ZendAstGetZval(ast.GetChild()[0]))
	case ZEND_AST_UNARY_PLUS:
		fallthrough
	case ZEND_AST_UNARY_MINUS:
		ZendEvalConstExpr(ast.GetChild()[0])
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
		ZendEvalConstExpr(ast.GetChild()[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {

			/* ensure everything was compile-time evaluated at least once */

			ZendEvalConstExpr(ast.GetChild()[1])
			return
		}
		if ZendAstGetZval(ast.GetChild()[0]).IsNull() {
			ZendEvalConstExpr(ast.GetChild()[1])
			*ast_ptr = ast.GetChild()[1]
			ast.GetChild()[1] = nil
			ZendAstDestroy(ast)
		} else {
			*ast_ptr = ast.GetChild()[0]
			ast.GetChild()[0] = nil
			ZendAstDestroy(ast)
		}
		return
	case ZEND_AST_CONDITIONAL:
		var child **ZendAst
		var child_ast **ZendAst
		ZendEvalConstExpr(ast.GetChild()[0])
		if ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL {

			/* ensure everything was compile-time evaluated at least once */

			if ast.GetChild()[1] != nil {
				ZendEvalConstExpr(ast.GetChild()[1])
			}
			ZendEvalConstExpr(ast.GetChild()[2])
			return
		}
		child = ast.GetChild()[2-operators.IZendIsTrue(ZendAstGetZval(ast.GetChild()[0]))]
		if (*child) == nil {
			child--
		}
		child_ast = *child
		*child = nil
		ZendAstDestroy(ast)
		*ast_ptr = child_ast
		ZendEvalConstExpr(ast_ptr)
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
		ZendEvalConstExpr(ast.GetChild()[0])
		ZendEvalConstExpr(ast.GetChild()[1])
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
			} else if dim.GetType() != types.IS_STRING || operators.IsNumericString(dim.String().GetStr(), &offset, nil, 1) != types.IS_LONG {
				return
			}
			if offset < 0 || int(offset >= container.String().GetLen()) != 0 {
				return
			}
			c = uint8(container.String().GetStr()[offset])
			result.SetStringVal(string(c))
		} else if container.GetType() <= types.IS_FALSE {
			result.SetNull()
		} else {
			return
		}
	case ZEND_AST_ARRAY:
		if ZendTryCtEvalArray(&result, ast) == 0 {
			return
		}
	case ZEND_AST_MAGIC_CONST:
		if ZendTryCtEvalMagicConst(&result, ast) == 0 {
			return
		}
	case ZEND_AST_CONST:
		var name_ast *ZendAst = ast.GetChild()[0]
		var is_fully_qualified types.ZendBool
		var resolved_name *types.String = ZendResolveConstName(ZendAstGetStr(name_ast), name_ast.GetAttr(), &is_fully_qualified)
		if ZendTryCtEvalConst(&result, resolved_name.GetStr(), is_fully_qualified) == 0 {
			// types.ZendStringReleaseEx(resolved_name, 0)
			return
		}
		// types.ZendStringReleaseEx(resolved_name, 0)
	case ZEND_AST_CLASS_CONST:
		var class_ast *ZendAst
		var name_ast *ZendAst
		var resolved_name *types.String
		ZendEvalConstExpr(ast.GetChild()[0])
		ZendEvalConstExpr(ast.GetChild()[1])
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
	ZendAstDestroy(ast)
	*ast_ptr = ZendAstCreateZval(&result)
}
