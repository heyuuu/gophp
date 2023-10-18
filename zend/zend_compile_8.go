package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func (compiler *Compiler) CompileUnaryPm(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.Child(0)
	var expr_node Znode
	var lefthand_node Znode
	b.Assert(ast.Kind() == ZEND_AST_UNARY_PLUS || ast.Kind() == ZEND_AST_UNARY_MINUS)
	compiler.CompileExpr(&expr_node, expr_ast)
	if expr_node.GetOpType() == IS_CONST {
		if ZendTryCtEvalUnaryPm(result.GetConstant(), ast.Kind(), expr_node.GetConstant()) != 0 {
			result.SetOpType(IS_CONST)
			// ZvalPtrDtor(expr_node.GetConstant())
			return
		}
	}
	lefthand_node.SetOpType(IS_CONST)
	lefthand_node.GetConstant().SetLong(lang.Cond(ast.Kind() == ZEND_AST_UNARY_PLUS, 1, -1))
	ZendEmitOpTmp(result, ZEND_MUL, &lefthand_node, &expr_node)
}
func (compiler *Compiler) CompileShortCircuiting(result *Znode, ast *ZendAst) {
	var left_ast *ZendAst = ast.Child(0)
	var right_ast *ZendAst = ast.Child(1)
	var left_node Znode
	var right_node Znode
	var opline_jmpz *types.ZendOp
	var opline_bool *types.ZendOp
	var opnum_jmpz uint32
	b.Assert(ast.Kind() == ZEND_AST_AND || ast.Kind() == ZEND_AST_OR)
	compiler.CompileExpr(&left_node, left_ast)
	if left_node.GetOpType() == IS_CONST {
		if ast.Kind() == ZEND_AST_AND && operators.IZendIsTrue(left_node.GetConstant()) == 0 || ast.Kind() == ZEND_AST_OR && operators.IZendIsTrue(left_node.GetConstant()) != 0 {
			result.SetOpType(IS_CONST)
			result.GetConstant().SetBool(operators.IZendIsTrue(left_node.GetConstant()) != 0)
		} else {
			compiler.CompileExpr(&right_node, right_ast)
			if right_node.GetOpType() == IS_CONST {
				result.SetOpType(IS_CONST)
				result.GetConstant().SetBool(operators.IZendIsTrue(right_node.GetConstant()) != 0)
				// ZvalPtrDtor(right_node.GetConstant())
			} else {
				ZendEmitOpTmp(result, ZEND_BOOL, &right_node, nil)
			}
		}
		// ZvalPtrDtor(left_node.GetConstant())
		return
	}
	opnum_jmpz = GetNextOpNumber()
	opline_jmpz = ZendEmitOp(nil, lang.Cond(ast.Kind() == ZEND_AST_AND, ZEND_JMPZ_EX, ZEND_JMPNZ_EX), &left_node, nil)
	if left_node.GetOpType() == IS_TMP_VAR {
		opline_jmpz.SetResultType(left_node.GetOpType())
		if left_node.GetOpType() == IS_CONST {
			opline_jmpz.GetResult().SetConstant(ZendAddLiteral(left_node.GetConstant()))
		} else {
			opline_jmpz.SetResult(left_node.GetOp())
		}
	} else {
		opline_jmpz.GetResult().SetVar(GetTemporaryVariable())
		opline_jmpz.SetResultType(IS_TMP_VAR)
	}
	result.SetOpType(opline_jmpz.GetResultType())
	if result.GetOpType() == IS_CONST {
		types.ZVAL_COPY_VALUE(result.GetConstant(), CT_CONSTANT(opline_jmpz.GetResult()))
	} else {
		result.SetOp(opline_jmpz.GetResult())
	}
	compiler.CompileExpr(&right_node, right_ast)
	opline_bool = ZendEmitOp(nil, ZEND_BOOL, &right_node, nil)
	opline_bool.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
		opline_bool.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
	} else {
		opline_bool.SetResult(result.GetOp())
	}
	ZendUpdateJumpTargetToNext(opnum_jmpz)
}
func (compiler *Compiler) CompilePostIncdec(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.Child(0)
	b.Assert(ast.Kind() == ZEND_AST_POST_INC || ast.Kind() == ZEND_AST_POST_DEC)
	ZendEnsureWritableVariable(var_ast)
	if var_ast.Kind() == ZEND_AST_PROP {
		var opline *types.ZendOp = compiler.CompileProp(nil, var_ast, BP_VAR_RW, 0)
		if ast.Kind() == ZEND_AST_POST_INC {
			opline.SetOpcode(ZEND_POST_INC_OBJ)
		} else {
			opline.SetOpcode(ZEND_POST_DEC_OBJ)
		}
		ZendMakeTmpResult(result, opline)
	} else if var_ast.Kind() == ZEND_AST_STATIC_PROP {
		var opline *types.ZendOp = compiler.CompileStaticProp(nil, var_ast, BP_VAR_RW, 0, 0)
		if ast.Kind() == ZEND_AST_POST_INC {
			opline.SetOpcode(ZEND_POST_INC_STATIC_PROP)
		} else {
			opline.SetOpcode(ZEND_POST_DEC_STATIC_PROP)
		}
		ZendMakeTmpResult(result, opline)
	} else {
		var var_node Znode
		compiler.CompileVar(&var_node, var_ast, BP_VAR_RW, 0)
		ZendEmitOpTmp(result, lang.Cond(ast.Kind() == ZEND_AST_POST_INC, ZEND_POST_INC, ZEND_POST_DEC), &var_node, nil)
	}
}
func (compiler *Compiler) CompilePreIncdec(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.Child(0)
	b.Assert(ast.Kind() == ZEND_AST_PRE_INC || ast.Kind() == ZEND_AST_PRE_DEC)
	ZendEnsureWritableVariable(var_ast)
	if var_ast.Kind() == ZEND_AST_PROP {
		var opline *types.ZendOp = compiler.CompileProp(result, var_ast, BP_VAR_RW, 0)
		if ast.Kind() == ZEND_AST_PRE_INC {
			opline.SetOpcode(ZEND_PRE_INC_OBJ)
		} else {
			opline.SetOpcode(ZEND_PRE_DEC_OBJ)
		}
	} else if var_ast.Kind() == ZEND_AST_STATIC_PROP {
		var opline *types.ZendOp = compiler.CompileStaticProp(result, var_ast, BP_VAR_RW, 0, 0)
		if ast.Kind() == ZEND_AST_PRE_INC {
			opline.SetOpcode(ZEND_PRE_INC_STATIC_PROP)
		} else {
			opline.SetOpcode(ZEND_PRE_DEC_STATIC_PROP)
		}
	} else {
		var var_node Znode
		compiler.CompileVar(&var_node, var_ast, BP_VAR_RW, 0)
		ZendEmitOp(result, lang.Cond(ast.Kind() == ZEND_AST_PRE_INC, ZEND_PRE_INC, ZEND_PRE_DEC), &var_node, nil)
	}
}
func (compiler *Compiler) CompileCast(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.Child(0)
	var expr_node Znode
	var opline *types.ZendOp
	compiler.CompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOpTmp(result, ZEND_CAST, &expr_node, nil)
	opline.SetExtendedValue(ast.Attr())
	if ast.Attr() == types.IsNull {
		faults.Error(faults.E_DEPRECATED, "The (unset) cast is deprecated")
	}
}
func (compiler *Compiler) CompileShorthandConditional(result *Znode, ast *ZendAst) {
	var cond_ast *ZendAst = ast.Child(0)
	var false_ast *ZendAst = ast.Children()[2]
	var cond_node Znode
	var false_node Znode
	var opline_qm_assign *types.ZendOp
	var opnum_jmp_set uint32
	b.Assert(ast.Child(1) == nil)
	compiler.CompileExpr(&cond_node, cond_ast)
	opnum_jmp_set = GetNextOpNumber()
	ZendEmitOpTmp(result, ZEND_JMP_SET, &cond_node, nil)
	compiler.CompileExpr(&false_node, false_ast)
	opline_qm_assign = ZendEmitOpTmp(nil, ZEND_QM_ASSIGN, &false_node, nil)
	opline_qm_assign.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
		opline_qm_assign.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
	} else {
		opline_qm_assign.SetResult(result.GetOp())
	}
	ZendUpdateJumpTargetToNext(opnum_jmp_set)
}
func (compiler *Compiler) CompileConditional(result *Znode, ast *ZendAst) {
	var cond_ast *ZendAst = ast.Child(0)
	var true_ast *ZendAst = ast.Child(1)
	var false_ast *ZendAst = ast.Children()[2]
	var cond_node Znode
	var true_node Znode
	var false_node Znode
	var opline_qm_assign2 *types.ZendOp
	var opnum_jmpz uint32
	var opnum_jmp uint32
	if cond_ast.Kind() == ZEND_AST_CONDITIONAL && cond_ast.Attr() != ZEND_PARENTHESIZED_CONDITIONAL {
		if cond_ast.Child(1) != nil {
			if true_ast != nil {
				faults.Error(faults.E_DEPRECATED, "Unparenthesized `a ? b : c ? d : e` is deprecated. Use either `(a ? b : c) ? d : e` or `a ? b : (c ? d : e)`")
			} else {
				faults.Error(faults.E_DEPRECATED, "Unparenthesized `a ? b : c ?: d` is deprecated. Use either `(a ? b : c) ?: d` or `a ? b : (c ?: d)`")
			}
		} else {
			if true_ast != nil {
				faults.Error(faults.E_DEPRECATED, "Unparenthesized `a ?: b ? c : d` is deprecated. Use either `(a ?: b) ? c : d` or `a ?: (b ? c : d)`")
			}
		}
	}
	if true_ast == nil {
		compiler.CompileShorthandConditional(result, ast)
		return
	}
	compiler.CompileExpr(&cond_node, cond_ast)
	opnum_jmpz = ZendEmitCondJump(ZEND_JMPZ, &cond_node, 0)
	compiler.CompileExpr(&true_node, true_ast)
	ZendEmitOpTmp(result, ZEND_QM_ASSIGN, &true_node, nil)
	opnum_jmp = ZendEmitJump(0)
	ZendUpdateJumpTargetToNext(opnum_jmpz)
	compiler.CompileExpr(&false_node, false_ast)
	opline_qm_assign2 = ZendEmitOp(nil, ZEND_QM_ASSIGN, &false_node, nil)
	opline_qm_assign2.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
		opline_qm_assign2.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
	} else {
		opline_qm_assign2.SetResult(result.GetOp())
	}
	ZendUpdateJumpTargetToNext(opnum_jmp)
}
func (compiler *Compiler) CompileCoalesce(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.Children()[0]
	var default_ast *ZendAst = ast.Child(1)
	var expr_node Znode
	var default_node Znode
	var opline *types.ZendOp
	var opnum uint32
	compiler.CompileVar(&expr_node, expr_ast, BP_VAR_IS, 0)
	opnum = GetNextOpNumber()
	ZendEmitOpTmp(result, ZEND_COALESCE, &expr_node, nil)
	compiler.CompileExpr(&default_node, default_ast)
	opline = ZendEmitOpTmp(nil, ZEND_QM_ASSIGN, &default_node, nil)
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
		opline.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
	} else {
		opline.SetResult(result.GetOp())
	}
	opline = CG__().GetActiveOpArray().GetOpcodes()[opnum]
	opline.GetOp2().SetOplineNum(GetNextOpNumber())
}
func (compiler *Compiler) CompileAssignCoalesce(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.Children()[0]
	var default_ast *ZendAst = ast.Child(1)
	var var_node_is Znode
	var var_node_w Znode
	var default_node Znode
	var assign_node Znode
	var node *Znode
	var opline *types.ZendOp
	var coalesce_opnum uint32
	var need_frees bool = 0

	/* Remember expressions compiled during the initial BP_VAR_IS lookup,
	 * to avoid double-evaluation when we compile again with BP_VAR_W. */

	var orig_memoized_exprs *types.Array = CG__().GetMemoizedExprs()
	var orig_memoize_mode int = CG__().GetMemoizeMode()
	ZendEnsureWritableVariable(var_ast)
	if IsThisFetch(var_ast) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot re-assign $this")
	}
	CG__().SetMemoizedExprs(types.NewArray())
	CG__().SetMemoizeMode(ZEND_MEMOIZE_COMPILE)
	compiler.CompileVar(&var_node_is, var_ast, BP_VAR_IS, 0)
	coalesce_opnum = GetNextOpNumber()
	ZendEmitOpTmp(result, ZEND_COALESCE, &var_node_is, nil)
	CG__().SetMemoizeMode(ZEND_MEMOIZE_NONE)
	compiler.CompileExpr(&default_node, default_ast)
	CG__().SetMemoizeMode(ZEND_MEMOIZE_FETCH)
	compiler.CompileVar(&var_node_w, var_ast, BP_VAR_W, 0)

	/* Reproduce some of the zend_compile_assign() opcode fixup logic here. */

	opline = CG__().GetActiveOpArray().GetOpcodes()[CG__().GetActiveOpArray().GetLast()-1]
	switch var_ast.Kind() {
	case ZEND_AST_VAR:
		ZendEmitOp(&assign_node, ZEND_ASSIGN, &var_node_w, &default_node)
	case ZEND_AST_STATIC_PROP:
		opline.SetOpcode(ZEND_ASSIGN_STATIC_PROP)
		ZendEmitOpData(&default_node)
		assign_node = var_node_w
	case ZEND_AST_DIM:
		opline.SetOpcode(ZEND_ASSIGN_DIM)
		ZendEmitOpData(&default_node)
		assign_node = var_node_w
	case ZEND_AST_PROP:
		opline.SetOpcode(ZEND_ASSIGN_OBJ)
		ZendEmitOpData(&default_node)
		assign_node = var_node_w
	default:

	}
	opline = ZendEmitOpTmp(nil, ZEND_QM_ASSIGN, &assign_node, nil)
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
		opline.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
	} else {
		opline.SetResult(result.GetOp())
	}
	CG__().GetMemoizedExprs().ForeachEx(func(_ types.ArrayKey, value *types.Zval) bool {
		node = value.Ptr()
		if node.GetOpType() == IS_TMP_VAR || node.GetOpType() == IS_VAR {
			need_frees = 1
			return false
		}
		return true
	})

	/* Free DUPed expressions if there are any */
	if need_frees != 0 {
		var jump_opnum uint32 = ZendEmitJump(0)
		ZendUpdateJumpTargetToNext(coalesce_opnum)
		CG__().GetMemoizedExprs().Foreach(func(_ types.ArrayKey, value *types.Zval) {
			node = value.Ptr()
			if node.GetOpType() == IS_TMP_VAR || node.GetOpType() == IS_VAR {
				ZendEmitOp(nil, ZEND_FREE, node, nil)
			}
		})
		ZendUpdateJumpTargetToNext(jump_opnum)
	} else {
		ZendUpdateJumpTargetToNext(coalesce_opnum)
	}
	CG__().GetMemoizedExprs().Destroy()
	CG__().SetMemoizeMode(orig_memoize_mode)
}
func (compiler *Compiler) CompilePrint(result *Znode, ast *ZendAst) {
	var opline *types.ZendOp
	var expr_ast *ZendAst = ast.Child(0)
	var expr_node Znode
	compiler.CompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOp(nil, ZEND_ECHO, &expr_node, nil)
	opline.SetExtendedValue(1)
	result.SetOpType(IS_CONST)
	result.GetConstant().SetLong(1)
}
func (compiler *Compiler) CompileExit(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.Children()[0]
	if expr_ast != nil {
		var expr_node Znode
		compiler.CompileExpr(&expr_node, expr_ast)
		ZendEmitOp(nil, ZEND_EXIT, &expr_node, nil)
	} else {
		ZendEmitOp(nil, ZEND_EXIT, nil, nil)
	}
	result.SetOpType(IS_CONST)
	result.GetConstant().SetBool(true)
}
func (compiler *Compiler) CompileYield(result *Znode, ast *ZendAst) {
	var value_ast *ZendAst = ast.Children()[0]
	var key_ast *ZendAst = ast.Child(1)
	var value_node Znode
	var key_node Znode
	var value_node_ptr *Znode = nil
	var key_node_ptr *Znode = nil
	var opline *types.ZendOp
	var returns_by_ref bool = CG__().GetActiveOpArray().IsReturnReference()
	ZendMarkFunctionAsGenerator()
	if key_ast != nil {
		compiler.CompileExpr(&key_node, key_ast)
		key_node_ptr = &key_node
	}
	if value_ast != nil {
		if returns_by_ref != 0 && ZendIsVariable(value_ast) != 0 {
			compiler.CompileVar(&value_node, value_ast, BP_VAR_W, 1)
		} else {
			compiler.CompileExpr(&value_node, value_ast)
		}
		value_node_ptr = &value_node
	}
	opline = ZendEmitOp(result, ZEND_YIELD, value_node_ptr, key_node_ptr)
	if value_ast != nil && returns_by_ref != 0 && ZendIsCall(value_ast) != 0 {
		opline.SetExtendedValue(ZEND_RETURNS_FUNCTION)
	}
}
func (compiler *Compiler) CompileYieldFrom(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.Children()[0]
	var expr_node Znode
	ZendMarkFunctionAsGenerator()
	if CG__().GetActiveOpArray().IsReturnReference() {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, `Cannot use "yield from" inside a by-reference generator`)
	}
	compiler.CompileExpr(&expr_node, expr_ast)
	ZendEmitOpTmp(result, ZEND_YIELD_FROM, &expr_node, nil)
}
func (compiler *Compiler) CompileInstanceof(result *Znode, ast *ZendAst) {
	var obj_ast *ZendAst = ast.Children()[0]
	var class_ast *ZendAst = ast.Child(1)
	var obj_node Znode
	var class_node Znode
	var opline *types.ZendOp
	compiler.CompileExpr(&obj_node, obj_ast)
	if obj_node.GetOpType() == IS_CONST {
		ZendDoFree(&obj_node)
		result.SetOpType(IS_CONST)
		result.GetConstant().SetFalse()
		return
	}
	compiler.CompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_NO_AUTOLOAD|ZEND_FETCH_CLASS_EXCEPTION)
	opline = ZendEmitOpTmp(result, ZEND_INSTANCEOF, &obj_node, nil)
	if class_node.GetOpType() == IS_CONST {
		opline.SetOp2Type(IS_CONST)
		opline.GetOp2().SetConstant(ZendAddClassNameLiteral(class_node.GetConstant().StringEx()))
		opline.SetExtendedValue(ZendAllocCacheSlot())
	} else {
		opline.SetOp2Type(class_node.GetOpType())
		if class_node.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(class_node.GetConstant()))
		} else {
			opline.SetOp2(class_node.GetOp())
		}
	}
}
func (compiler *Compiler) CompileIncludeOrEval(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.Children()[0]
	var expr_node Znode
	var opline *types.ZendOp
	ZendDoExtendedFcallBegin()
	compiler.CompileExpr(&expr_node, expr_ast)
	opline = ZendEmitOp(result, ZEND_INCLUDE_OR_EVAL, &expr_node, nil)
	opline.SetExtendedValue(ast.Attr())
	ZendDoExtendedFcallEnd()
}
func (compiler *Compiler) CompileIssetOrEmpty(result *Znode, ast *ZendAst) {
	var varAst *ZendAst = ast.Children()[0]
	var varNode Znode
	var opline *types.ZendOp = nil
	b.Assert(ast.Kind() == ZEND_AST_ISSET || ast.Kind() == ZEND_AST_EMPTY)
	if !ZendIsVariable(varAst) {
		if ast.Kind() == ZEND_AST_EMPTY {
			/* empty(expr) can be transformed to !expr */
			var notAst = AstCreateEx(ZEND_AST_UNARY_OP, ZendAstAttr(ZEND_BOOL_NOT), varAst)
			compiler.CompileExpr(result, notAst)
			return
		} else {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, `Cannot use isset() on the result of an expression (you can use "null !== expression" instead)`)
		}
	}
	switch varAst.Kind() {
	case ZEND_AST_VAR:
		if IsThisFetch(varAst) {
			opline = ZendEmitOp(result, ZEND_ISSET_ISEMPTY_THIS, nil, nil)
			CG__().GetActiveOpArray().SetIsUsesThis(true)
		} else if ZendTryCompileCv(&varNode, varAst) == types.SUCCESS {
			opline = ZendEmitOp(result, ZEND_ISSET_ISEMPTY_CV, &varNode, nil)
		} else {
			opline = compiler.CompileSimpleVarNoCv(result, varAst, BP_VAR_IS, 0)
			opline.SetOpcode(ZEND_ISSET_ISEMPTY_VAR)
		}
	case ZEND_AST_DIM:
		opline = compiler.CompileDim(result, varAst, BP_VAR_IS)
		opline.SetOpcode(ZEND_ISSET_ISEMPTY_DIM_OBJ)
	case ZEND_AST_PROP:
		opline = compiler.CompileProp(result, varAst, BP_VAR_IS, 0)
		opline.SetOpcode(ZEND_ISSET_ISEMPTY_PROP_OBJ)
	case ZEND_AST_STATIC_PROP:
		opline = compiler.CompileStaticProp(result, varAst, BP_VAR_IS, 0, 0)
		opline.SetOpcode(ZEND_ISSET_ISEMPTY_STATIC_PROP)
	default:

	}
	opline.SetResultType(IS_TMP_VAR)
	result.SetOpType(opline.GetResultType())
	if ast.Kind() != ZEND_AST_ISSET {
		opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_ISEMPTY)
	}
}
func (compiler *Compiler) CompileSilence(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.Children()[0]
	var silence_node Znode
	ZendEmitOpTmp(&silence_node, ZEND_BEGIN_SILENCE, nil, nil)
	if expr_ast.Kind() == ZEND_AST_VAR {

		/* For @$var we need to force a FETCH instruction, otherwise the CV access will
		 * happen outside the silenced section. */

		compiler.CompileSimpleVarNoCv(result, expr_ast, BP_VAR_R, 0)

		/* For @$var we need to force a FETCH instruction, otherwise the CV access will
		 * happen outside the silenced section. */

	} else {
		compiler.CompileExpr(result, expr_ast)
	}
	ZendEmitOp(nil, ZEND_END_SILENCE, &silence_node, nil)
}
func (compiler *Compiler) CompileShellExec(result *Znode, ast *ZendAst) {
	var expr_ast *ZendAst = ast.Children()[0]
	var fn_name types.Zval
	var name_ast *ZendAst
	var args_ast *ZendAst
	var call_ast *ZendAst
	fn_name.SetString("shell_exec")
	name_ast = ZendAstCreateZval(&fn_name)
	args_ast = AstCreateList(ZEND_AST_ARG_LIST, expr_ast)
	call_ast = AstCreate(ZEND_AST_CALL, name_ast, args_ast)
	compiler.CompileExpr(result, call_ast)
}
func (compiler *Compiler) CompileArray(result *Znode, ast *ZendAst) {
	var list *ZendAstList = ast.AsAstList()
	var opline *types.ZendOp
	var i uint32
	var opnum_init uint32 = -1
	var packed bool = 1
	if compiler.TryCtEvalArray(result.GetConstant(), ast) != 0 {
		result.SetOpType(IS_CONST)
		return
	}

	/* Empty arrays are handled at compile-time */

	b.Assert(list.GetChildren() > 0)
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.Children()[i]
		var value_ast *ZendAst
		var key_ast *ZendAst
		var by_ref bool
		var value_node Znode
		var key_node Znode
		var key_node_ptr *Znode = nil
		if elem_ast == nil {
			faults.Error(faults.E_COMPILE_ERROR, "Cannot use empty array elements in arrays")
		}
		value_ast = elem_ast.Children()[0]
		if elem_ast.Kind() == ZEND_AST_UNPACK {
			compiler.CompileExpr(&value_node, value_ast)
			if i == 0 {
				opnum_init = GetNextOpNumber()
				opline = ZendEmitOpTmp(result, ZEND_INIT_ARRAY, nil, nil)
			}
			opline = ZendEmitOp(nil, ZEND_ADD_ARRAY_UNPACK, &value_node, nil)
			opline.SetResultType(result.GetOpType())
			if result.GetOpType() == IS_CONST {
				opline.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
			} else {
				opline.SetResult(result.GetOp())
			}
			continue
		}
		key_ast = elem_ast.Child(1)
		by_ref = elem_ast.Attr()
		if key_ast != nil {
			compiler.CompileExpr(&key_node, key_ast)
			ZendHandleNumericOp(&key_node)
			key_node_ptr = &key_node
		}
		if by_ref != 0 {
			ZendEnsureWritableVariable(value_ast)
			compiler.CompileVar(&value_node, value_ast, BP_VAR_W, 1)
		} else {
			compiler.CompileExpr(&value_node, value_ast)
		}
		if i == 0 {
			opnum_init = GetNextOpNumber()
			opline = ZendEmitOpTmp(result, ZEND_INIT_ARRAY, &value_node, key_node_ptr)
			opline.SetExtendedValue(list.GetChildren() << ZEND_ARRAY_SIZE_SHIFT)
		} else {
			opline = ZendEmitOp(nil, ZEND_ADD_ARRAY_ELEMENT, &value_node, key_node_ptr)
			opline.SetResultType(result.GetOpType())
			if result.GetOpType() == IS_CONST {
				opline.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
			} else {
				opline.SetResult(result.GetOp())
			}
		}
		opline.SetExtendedValue(opline.GetExtendedValue() | by_ref)
		if key_ast != nil && key_node.GetOpType() == IS_CONST && key_node.GetConstant().IsString() {
			packed = 0
		}
	}

	/* Add a flag to INIT_ARRAY if we know this array cannot be packed */

	if packed == 0 {
		b.Assert(opnum_init != uint32-1)
		opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_init]
		opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_ARRAY_NOT_PACKED)
	}

	/* Add a flag to INIT_ARRAY if we know this array cannot be packed */
}
func (compiler *Compiler) CompileConst(result *Znode, ast *ZendAst) {
	var name_ast *ZendAst = ast.Children()[0]
	var opline *types.ZendOp
	var is_fully_qualified bool
	var orig_name *types.String = ZendAstGetStr(name_ast)
	var resolved_name, isFullyQualified = ZendResolveConstName(orig_name, name_ast.Attr())
	if resolved_name == "__COMPILER_HALT_OFFSET__" || name_ast.Attr() != ZEND_NAME_RELATIVE && orig_name.GetStr() == "__COMPILER_HALT_OFFSET__" {
		var last *ZendAst = CG__().GetAst()
		for last != nil && last.Kind() == ZEND_AST_STMT_LIST {
			var list *ZendAstList = last.AsAstList()
			if list.GetChildren() == 0 {
				break
			}
			last = list.Children()[list.GetChildren()-1]
		}
		if last != nil && last.Kind() == ZEND_AST_HALT_COMPILER {
			result.SetOpType(IS_CONST)
			result.GetConstant().SetLong(ZendAstGetZval(last.Children()[0]).Long())
			// types.ZendStringReleaseEx(resolved_name, 0)
			return
		}
	}
	if ZendTryCtEvalConst(result.GetConstant(), resolved_name, isFullyQualified) != 0 {
		result.SetOpType(IS_CONST)
		// types.ZendStringReleaseEx(resolved_name, 0)
		return
	}
	opline = ZendEmitOpTmp(result, ZEND_FETCH_CONSTANT, nil, nil)
	opline.SetOp2Type(IS_CONST)
	if isFullyQualified {
		opline.GetOp2().SetConstant(ZendAddConstNameLiteral(resolved_name, 0))
	} else {
		opline.GetOp1().SetNum(IS_CONSTANT_UNQUALIFIED)
		if FC__().CurrentNamespace() != "" {
			opline.GetOp1().SetNum(opline.GetOp1().GetNum() | IS_CONSTANT_IN_NAMESPACE)
			opline.GetOp2().SetConstant(ZendAddConstNameLiteral(resolved_name.GetStr(), 1))
		} else {
			opline.GetOp2().SetConstant(ZendAddConstNameLiteral(resolved_name.GetStr(), 0))
		}
	}
	opline.SetExtendedValue(ZendAllocCacheSlot())
}
func (compiler *Compiler) CompileClassConst(result *Znode, ast *ZendAst) {
	var class_ast *ZendAst = ast.Children()[0]
	var const_ast *ZendAst = ast.Child(1)
	var class_node Znode
	var const_node Znode
	var opline *types.ZendOp
	compiler.EvalConstExpr(ast.Children()[0])
	compiler.EvalConstExpr(ast.Child(1))
	class_ast = ast.Children()[0]
	const_ast = ast.Child(1)
	if class_ast.Kind() == ZEND_AST_ZVAL {
		var resolved_name *types.String
		resolved_name = ZendResolveClassNameAst(class_ast)
		if const_ast.Kind() == ZEND_AST_ZVAL && ZendTryCtEvalClassConst(result.GetConstant(), resolved_name, ZendAstGetStr(const_ast)) != 0 {
			result.SetOpType(IS_CONST)
			// types.ZendStringReleaseEx(resolved_name, 0)
			return
		}
		// types.ZendStringReleaseEx(resolved_name, 0)
	}
	compiler.CompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_EXCEPTION)
	compiler.CompileExpr(&const_node, const_ast)
	opline = ZendEmitOpTmp(result, ZEND_FETCH_CLASS_CONSTANT, nil, &const_node)
	ZendSetClassNameOp1(opline, &class_node)
	opline.SetExtendedValue(ZendAllocCacheSlots(2))
}
func (compiler *Compiler) CompileClassName(result *Znode, ast *ZendAst) {
	var class_ast *ZendAst = ast.Children()[0]
	var opline *types.ZendOp
	if ZendTryCompileConstExprResolveClassName(result.GetConstant(), class_ast) != 0 {
		result.SetOpType(IS_CONST)
		return
	}
	opline = ZendEmitOpTmp(result, ZEND_FETCH_CLASS_NAME, nil, nil)
	opline.GetOp1().SetNum(ZendGetClassFetchType(ZendAstGetStrVal(class_ast)))
}
func (compiler *Compiler) CompileRopeAddEx(opline *types.ZendOp, result *Znode, num uint32, elem_node *Znode) *types.ZendOp {
	if num == 0 {
		result.SetOpType(IS_TMP_VAR)
		result.GetOp().SetVar(-1)
		opline.SetOpcode(ZEND_ROPE_INIT)
	} else {
		opline.SetOpcode(ZEND_ROPE_ADD)
		opline.SetOp1Type(result.GetOpType())
		if result.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(result.GetConstant()))
		} else {
			opline.SetOp1(result.GetOp())
		}
	}
	opline.SetOp2Type(elem_node.GetOpType())
	if elem_node.GetOpType() == IS_CONST {
		opline.GetOp2().SetConstant(ZendAddLiteral(elem_node.GetConstant()))
	} else {
		opline.SetOp2(elem_node.GetOp())
	}
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
		opline.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
	} else {
		opline.SetResult(result.GetOp())
	}
	opline.SetExtendedValue(num)
	return opline
}
func (compiler *Compiler) CompileRopeAdd(result *Znode, num uint32, elem_node *Znode) *types.ZendOp {
	var opline *types.ZendOp = GetNextOp()
	if num == 0 {
		result.SetOpType(IS_TMP_VAR)
		result.GetOp().SetVar(-1)
		opline.SetOpcode(ZEND_ROPE_INIT)
	} else {
		opline.SetOpcode(ZEND_ROPE_ADD)
		opline.SetOp1Type(result.GetOpType())
		if result.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(result.GetConstant()))
		} else {
			opline.SetOp1(result.GetOp())
		}
	}
	opline.SetOp2Type(elem_node.GetOpType())
	if elem_node.GetOpType() == IS_CONST {
		opline.GetOp2().SetConstant(ZendAddLiteral(elem_node.GetConstant()))
	} else {
		opline.SetOp2(elem_node.GetOp())
	}
	opline.SetResultType(result.GetOpType())
	if result.GetOpType() == IS_CONST {
		opline.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
	} else {
		opline.SetResult(result.GetOp())
	}
	opline.SetExtendedValue(num)
	return opline
}
func (compiler *Compiler) CompileEncapsList(result *Znode, ast *ZendAst) {
	var i uint32
	var j uint32
	var rope_init_lineno uint32 = -1
	var opline *types.ZendOp = nil
	var init_opline *types.ZendOp
	var elem_node Znode
	var last_const_node Znode
	var list *ZendAstList = ast.AsAstList()
	var reserved_op_number uint32 = -1
	b.Assert(list.GetChildren() > 0)
	j = 0
	last_const_node.SetOpType(IS_UNUSED)
	for i = 0; i < list.GetChildren(); i++ {
		compiler.CompileExpr(&elem_node, list.Children()[i])
		if elem_node.GetOpType() == IS_CONST {
			operators.ConvertToString(elem_node.GetConstant())
			if elem_node.GetConstant().String() == "" {
				// ZvalPtrDtor(elem_node.GetConstant())
			} else if last_const_node.GetOpType() == IS_CONST {
				operators.ConcatFunction(last_const_node.GetConstant(), last_const_node.GetConstant(), elem_node.GetConstant())
				// ZvalPtrDtor(elem_node.GetConstant())
			} else {
				last_const_node.SetOpType(IS_CONST)
				types.ZVAL_COPY_VALUE(last_const_node.GetConstant(), elem_node.GetConstant())

				/* Reserve place for ZEND_ROPE_ADD instruction */

				reserved_op_number = GetNextOpNumber()
				opline = GetNextOp()
				opline.SetOpcode(ZEND_NOP)
			}
			continue
		} else {
			if j == 0 {
				if last_const_node.GetOpType() == IS_CONST {
					rope_init_lineno = reserved_op_number
				} else {
					rope_init_lineno = GetNextOpNumber()
				}
			}
			if last_const_node.GetOpType() == IS_CONST {
				opline = CG__().GetActiveOpArray().GetOpcodes()[reserved_op_number]
				compiler.CompileRopeAddEx(opline, result, lang.PostInc(&j), &last_const_node)
				last_const_node.SetOpType(IS_UNUSED)
			}
			opline = compiler.CompileRopeAdd(result, lang.PostInc(&j), &elem_node)
		}
	}
	if j == 0 {
		result.SetOpType(IS_CONST)
		if last_const_node.GetOpType() == IS_CONST {
			types.ZVAL_COPY_VALUE(result.GetConstant(), last_const_node.GetConstant())
		} else {
			result.GetConstant().SetString("")
		}
		CG__().GetActiveOpArray().SetLast(reserved_op_number - 1)
		return
	} else if last_const_node.GetOpType() == IS_CONST {
		opline = CG__().GetActiveOpArray().GetOpcodes()[reserved_op_number]
		opline = compiler.CompileRopeAddEx(opline, result, lang.PostInc(&j), &last_const_node)
	}
	init_opline = CG__().GetActiveOpArray().GetOpcodes() + rope_init_lineno
	if j == 1 {
		if opline.GetOp2Type() == IS_CONST {
			result.SetOpType(opline.GetOp2Type())
			if result.GetOpType() == IS_CONST {
				types.ZVAL_COPY_VALUE(result.GetConstant(), CT_CONSTANT(opline.GetOp2()))
			} else {
				result.SetOp(opline.GetOp2())
			}
			opline.SetNop()
		} else {
			opline.SetOpcode(ZEND_CAST)
			opline.SetExtendedValue(types.IsString)
			opline.SetOp1Type(opline.GetOp2Type())
			opline.SetOp1(opline.GetOp2())
			opline.SetResultType(IS_TMP_VAR)
			opline.GetResult().SetVar(GetTemporaryVariable())
			opline.SetOp2Type(IS_UNUSED)
			result.SetOpType(opline.GetResultType())
			if result.GetOpType() == IS_CONST {
				types.ZVAL_COPY_VALUE(result.GetConstant(), CT_CONSTANT(opline.GetResult()))
			} else {
				result.SetOp(opline.GetResult())
			}
		}
	} else if j == 2 {
		opline.SetOpcode(ZEND_FAST_CONCAT)
		opline.SetExtendedValue(0)
		opline.SetOp1Type(init_opline.GetOp2Type())
		opline.SetOp1(init_opline.GetOp2())
		opline.SetResultType(IS_TMP_VAR)
		opline.GetResult().SetVar(GetTemporaryVariable())
		init_opline.SetNop()
		result.SetOpType(opline.GetResultType())
		if result.GetOpType() == IS_CONST {
			types.ZVAL_COPY_VALUE(result.GetConstant(), CT_CONSTANT(opline.GetResult()))
		} else {
			result.SetOp(opline.GetResult())
		}
	} else {
		var var_ uint32
		init_opline.SetExtendedValue(j)
		opline.SetOpcode(ZEND_ROPE_END)
		opline.GetResult().SetVar(GetTemporaryVariable())
		opline.GetOp1().SetVar(GetTemporaryVariable())
		var_ = opline.GetOp1().GetVar()
		result.SetOpType(opline.GetResultType())
		if result.GetOpType() == IS_CONST {
			types.ZVAL_COPY_VALUE(result.GetConstant(), CT_CONSTANT(opline.GetResult()))
		} else {
			result.SetOp(opline.GetResult())
		}

		/* Allocates the necessary number of zval slots to keep the rope */

		i = (j*b.SizeOf("zend_string *") + (b.SizeOf("zval") - 1)) / b.SizeOf("zval")
		for i > 1 {
			GetTemporaryVariable()
			i--
		}

		/* Update all the previous opcodes to use the same variable */

		for opline != init_opline {
			opline--
			if opline.GetOpcode() == ZEND_ROPE_ADD && opline.GetResult().GetVar() == uint32-1 {
				opline.GetOp1().SetVar(var_)
				opline.GetResult().SetVar(var_)
			} else if opline.GetOpcode() == ZEND_ROPE_INIT && opline.GetResult().GetVar() == uint32-1 {
				opline.GetResult().SetVar(var_)
			}
		}

		/* Update all the previous opcodes to use the same variable */

	}
}
