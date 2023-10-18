package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendMakeVarResult(result *Znode, opline *types.ZendOp) {
	opline.SetResultType(IS_VAR)
	opline.GetResult().SetVar(GetTemporaryVariable())
	result.SetOpType(opline.GetResultType())
	if result.GetOpType() == IS_CONST {
		types.ZVAL_COPY_VALUE(result.GetConstant(), CT_CONSTANT(opline.GetResult()))
	} else {
		result.SetOp(opline.GetResult())
	}
}
func ZendMakeTmpResult(result *Znode, opline *types.ZendOp) {
	opline.SetResultType(IS_TMP_VAR)
	opline.GetResult().SetVar(GetTemporaryVariable())
	result.SetOpType(opline.GetResultType())
	if result.GetOpType() == IS_CONST {
		types.ZVAL_COPY_VALUE(result.GetConstant(), CT_CONSTANT(opline.GetResult()))
	} else {
		result.SetOp(opline.GetResult())
	}
}
func ZendEmitOp(result *Znode, opcode OpCode, op1 *Znode, op2 *Znode) *types.ZendOp {
	var opline *types.ZendOp = GetNextOp()
	opline.SetOpcode(opcode)
	if op1 != nil {
		opline.SetOp1Type(op1.GetOpType())
		if op1.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(op1.GetConstant()))
		} else {
			opline.SetOp1(op1.GetOp())
		}
	}
	if op2 != nil {
		opline.SetOp2Type(op2.GetOpType())
		if op2.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(op2.GetConstant()))
		} else {
			opline.SetOp2(op2.GetOp())
		}
	}
	if result != nil {
		ZendMakeVarResult(result, opline)
	}
	return opline
}
func ZendEmitOpTmp(result *Znode, opcode OpCode, op1 *Znode, op2 *Znode) *types.ZendOp {
	var opline *types.ZendOp = GetNextOp()
	opline.SetOpcode(opcode)
	if op1 != nil {
		opline.SetOp1Type(op1.GetOpType())
		if op1.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(op1.GetConstant()))
		} else {
			opline.SetOp1(op1.GetOp())
		}
	}
	if op2 != nil {
		opline.SetOp2Type(op2.GetOpType())
		if op2.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(op2.GetConstant()))
		} else {
			opline.SetOp2(op2.GetOp())
		}
	}
	if result != nil {
		ZendMakeTmpResult(result, opline)
	}
	return opline
}
func ZendEmitOpData(value *Znode) *types.ZendOp {
	return ZendEmitOp(nil, ZEND_OP_DATA, value, nil)
}
func ZendEmitJump(opnum_target uint32) uint32 {
	var opnum uint32 = GetNextOpNumber()
	var opline *types.ZendOp = ZendEmitOp(nil, ZEND_JMP, nil, nil)
	opline.GetOp1().SetOplineNum(opnum_target)
	return opnum
}
func ZendIsSmartBranch(opline *types.ZendOp) int {
	switch opline.GetOpcode() {
	case ZEND_IS_IDENTICAL:
		fallthrough
	case ZEND_IS_NOT_IDENTICAL:
		fallthrough
	case ZEND_IS_EQUAL:
		fallthrough
	case ZEND_IS_NOT_EQUAL:
		fallthrough
	case ZEND_IS_SMALLER:
		fallthrough
	case ZEND_IS_SMALLER_OR_EQUAL:
		fallthrough
	case ZEND_CASE:
		fallthrough
	case ZEND_ISSET_ISEMPTY_CV:
		fallthrough
	case ZEND_ISSET_ISEMPTY_VAR:
		fallthrough
	case ZEND_ISSET_ISEMPTY_DIM_OBJ:
		fallthrough
	case ZEND_ISSET_ISEMPTY_PROP_OBJ:
		fallthrough
	case ZEND_ISSET_ISEMPTY_STATIC_PROP:
		fallthrough
	case ZEND_INSTANCEOF:
		fallthrough
	case ZEND_TYPE_CHECK:
		fallthrough
	case ZEND_DEFINED:
		fallthrough
	case ZEND_IN_ARRAY:
		fallthrough
	case ZEND_ARRAY_KEY_EXISTS:
		return 1
	default:
		return 0
	}
}
func ZendEmitCondJump(opcode uint8, cond *Znode, opnum_target uint32) uint32 {
	var opnum uint32 = GetNextOpNumber()
	var opline *types.ZendOp
	if (cond.GetOpType()&(IS_CV|IS_CONST)) != 0 && opnum > 0 && ZendIsSmartBranch(CG__().GetActiveOpArray().GetOpcodes()+opnum-1) != 0 {

		/* emit extra NOP to avoid incorrect SMART_BRANCH in very rare cases */

		ZendEmitOp(nil, ZEND_NOP, nil, nil)
		opnum = GetNextOpNumber()
	}
	opline = ZendEmitOp(nil, opcode, cond, nil)
	opline.GetOp2().SetOplineNum(opnum_target)
	return opnum
}
func ZendUpdateJumpTarget(opnum_jump uint32, opnum_target uint32) {
	var opline *types.ZendOp = CG__().GetActiveOpArray().GetOpcodes()[opnum_jump]
	switch opline.GetOpcode() {
	case ZEND_JMP:
		opline.GetOp1().SetOplineNum(opnum_target)
	case ZEND_JMPZ:
		fallthrough
	case ZEND_JMPNZ:
		fallthrough
	case ZEND_JMPZ_EX:
		fallthrough
	case ZEND_JMPNZ_EX:
		fallthrough
	case ZEND_JMP_SET:
		fallthrough
	case ZEND_COALESCE:
		opline.GetOp2().SetOplineNum(opnum_target)
	default:

	}
}
func ZendUpdateJumpTargetToNext(opnum_jump uint32) {
	ZendUpdateJumpTarget(opnum_jump, GetNextOpNumber())
}
func ZendDelayedEmitOp(result *Znode, opcode OpCode, op1 *Znode, op2 *Znode) *types.ZendOp {
	var tmp_opline *types.ZendOp = CurrCompiler().initOp(nil)
	tmp_opline.SetOpcode(opcode)
	if op1 != nil {
		tmp_opline.SetOp1Type(op1.GetOpType())
		if op1.GetOpType() == IS_CONST {
			tmp_opline.GetOp1().SetConstant(ZendAddLiteral(op1.GetConstant()))
		} else {
			tmp_opline.SetOp1(op1.GetOp())
		}
	}
	if op2 != nil {
		tmp_opline.SetOp2Type(op2.GetOpType())
		if op2.GetOpType() == IS_CONST {
			tmp_opline.GetOp2().SetConstant(ZendAddLiteral(op2.GetConstant()))
		} else {
			tmp_opline.SetOp2(op2.GetOp())
		}
	}
	if result != nil {
		ZendMakeVarResult(result, tmp_opline)
	}
	CG__().DelayedOplinesStackPush(tmp_opline)
	return CG__().DelayedOplinesStackTop()
}

func ZendDelayedCompileBlock(block func()) *types.ZendOp {
	offset := CG__().DelayedOplinesStackDepth()

	block()

	var opline *types.ZendOp = nil
	for _, delayedOpline := range CG__().DelayedOplinesStackCut(offset) {
		*opline = *delayedOpline
	}
	return opline
}

func (compiler *Compiler) CompileMemoizedExpr(result *Znode, expr *ZendAst) {
	var memoize_mode int = CG__().GetMemoizeMode()
	if memoize_mode == ZEND_MEMOIZE_COMPILE {
		var memoized_result Znode

		/* Go through normal compilation */

		CG__().SetMemoizeMode(ZEND_MEMOIZE_NONE)
		compiler.CompileExpr(result, expr)
		CG__().SetMemoizeMode(ZEND_MEMOIZE_COMPILE)
		if result.GetOpType() == IS_VAR {
			ZendEmitOp(&memoized_result, ZEND_COPY_TMP, result, nil)
		} else if result.GetOpType() == IS_TMP_VAR {
			ZendEmitOpTmp(&memoized_result, ZEND_COPY_TMP, result, nil)
		} else {
			if result.GetOpType() == IS_CONST {
				//result.GetConstant().TryAddRefcount()
			}
			memoized_result = *result
		}
		types.ZendHashIndexUpdateMem(CG__().GetMemoizedExprs(), uintPtr(expr), &memoized_result, b.SizeOf("znode"))
	} else if memoize_mode == ZEND_MEMOIZE_FETCH {
		var memoized_result *Znode = types.ZendHashIndexFindPtr(CG__().GetMemoizedExprs(), uintPtr(expr))
		*result = *memoized_result
		if result.GetOpType() == IS_CONST {
			//result.GetConstant().TryAddRefcount()
		}
	} else {
		b.Assert(false)
	}
}
func ZendEmitReturnTypeCheck(expr *Znode, return_info *ZendArgInfo, implicit bool) {
	if return_info.GetType().IsSet() {
		var opline *types.ZendOp

		/* `return ...;` is illegal in a void function (but `return;` isn't) */

		if return_info.GetType().Code() == types.IsVoid {
			if expr != nil {
				if expr.GetOpType() == IS_CONST && expr.GetConstant().IsNull() {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, `A void function must not return a value (did you mean "return;" instead of "return null;"?)`)
				} else {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "A void function must not return a value")
				}
			}

			/* we don't need run-time check */

			return

			/* we don't need run-time check */

		}
		if expr == nil && implicit == 0 {
			if return_info.GetType().AllowNull() {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, `A function with return type must return a value (did you mean "return null;" instead of "return;"?)`)
			} else {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "A function with return type must return a value")
			}
		}
		if expr != nil && expr.GetOpType() == IS_CONST {
			if return_info.GetType().Code() == expr.GetConstant().Type() || return_info.GetType().Code() == types.IsBool && (expr.GetConstant().IsFalse() || expr.GetConstant().IsTrue()) || return_info.GetType().AllowNull() && expr.GetConstant().IsNull() {

				/* we don't need run-time check */

				return
			}
		}
		opline = ZendEmitOp(nil, ZEND_VERIFY_RETURN_TYPE, expr, nil)
		if expr != nil && expr.GetOpType() == IS_CONST {
			expr.SetOpType(IS_TMP_VAR)
			opline.SetResultType(expr.GetOpType())
			expr.GetOp().SetVar(GetTemporaryVariable())
			opline.GetResult().SetVar(expr.GetOp().GetVar())
		}
		if return_info.GetType().IsClass() {
			opline.GetOp2().SetNum(CG__().GetActiveOpArray().GetCacheSize())
			CG__().GetActiveOpArray().SetCacheSize(CG__().GetActiveOpArray().GetCacheSize() + b.SizeOf("void *"))
		} else {
			opline.GetOp2().SetNum(-1)
		}
	}
}
func ZendEmitFinalReturn(returnOne bool) {
	var zn Znode
	var ret *types.ZendOp
	var returnsReference = CG__().GetActiveOpArray().IsReturnReference()
	if CG__().GetActiveOpArray().IsHasReturnType() && !CG__().GetActiveOpArray().IsGenerator() {
		ZendEmitReturnTypeCheck(nil, CG__().GetActiveOpArray().GetArgInfo()-1, 1)
	}
	zn.SetOpType(IS_CONST)
	if returnOne {
		zn.GetConstant().SetLong(1)
	} else {
		zn.GetConstant().SetNull()
	}
	ret = ZendEmitOp(nil, lang.Cond(returnsReference, ZEND_RETURN_BY_REF, ZEND_RETURN), &zn, nil)
	ret.SetExtendedValue(-1)
}
func ZendIsVariable(ast *ZendAst) bool {
	return ast.Kind() == ZEND_AST_VAR || ast.Kind() == ZEND_AST_DIM || ast.Kind() == ZEND_AST_PROP || ast.Kind() == ZEND_AST_STATIC_PROP
}
func ZendIsCall(ast *ZendAst) bool {
	return ast.Kind() == ZEND_AST_CALL || ast.Kind() == ZEND_AST_METHOD_CALL || ast.Kind() == ZEND_AST_STATIC_CALL
}
func ZendIsVariableOrCall(ast *ZendAst) bool {
	return ZendIsVariable(ast) != 0 || ZendIsCall(ast) != 0
}
func ZendIsUntickedStmt(ast *ZendAst) bool {
	return ast.Kind() == ZEND_AST_STMT_LIST || ast.Kind() == ZEND_AST_LABEL || ast.Kind() == ZEND_AST_PROP_DECL || ast.Kind() == ZEND_AST_CLASS_CONST_DECL || ast.Kind() == ZEND_AST_USE_TRAIT || ast.Kind() == ZEND_AST_METHOD
}
func ZendCanWriteToVariable(ast *ZendAst) bool {
	for ast.Kind() == ZEND_AST_DIM || ast.Kind() == ZEND_AST_PROP {
		ast = ast.Child(0)
	}
	return ZendIsVariableOrCall(ast)
}
func ZendIsConstDefaultClassRef(name_ast *ZendAst) bool {
	if name_ast.Kind() != ZEND_AST_ZVAL {
		return 0
	}
	return ZEND_FETCH_CLASS_DEFAULT == ZendGetClassFetchTypeAst(name_ast)
}
func ZendHandleNumericOp(node *Znode) {
	if node.GetOpType() == IS_CONST && node.GetConstant().IsString() {
		var index ZendUlong
		if types.HandleNumericStr(node.GetConstant().String(), &index) {
			node.GetConstant().SetLong(index)
		}
	}
}
func ZendHandleNumericDim(opline *types.ZendOp, dim_node *Znode) {
	if dim_node.GetConstant().IsString() {
		var index ZendUlong
		if types.HandleNumericStr(dim_node.GetConstant().String(), &index) {

			/* For numeric indexes we also keep the original value to use by ArrayAccess
			 * See bug #63217
			 */

			var c int = ZendAddLiteral(dim_node.GetConstant())
			b.Assert(opline.GetOp2().GetConstant()+1 == c)
			CT_CONSTANT(opline.GetOp2()).SetLong(index)
			CT_CONSTANT(opline.GetOp2()).GetU2Extra() = ZEND_EXTRA_VALUE
			return
		}
	}
}
func ZendSetClassNameOp1(opline *types.ZendOp, class_node *Znode) {
	if class_node.GetOpType() == IS_CONST {
		opline.SetOp1Type(IS_CONST)
		opline.GetOp1().SetConstant(ZendAddClassNameLiteral(class_node.GetConstant().String()))
	} else {
		opline.SetOp1Type(class_node.GetOpType())
		if class_node.GetOpType() == IS_CONST {
			opline.GetOp1().SetConstant(ZendAddLiteral(class_node.GetConstant()))
		} else {
			opline.SetOp1(class_node.GetOp())
		}
	}
}
func (compiler *Compiler) CompileClassRef(result *Znode, name_ast *ZendAst, fetch_flags uint32) {
	var fetch_type uint32
	if name_ast.Kind() != ZEND_AST_ZVAL {
		var name_node Znode
		compiler.CompileExpr(&name_node, name_ast)
		if name_node.GetOpType() == IS_CONST {
			var name string
			if name_node.GetConstant().Type() != types.IsString {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Illegal class name")
			}
			name = name_node.GetConstant().String()
			fetch_type = ZendGetClassFetchType(name)
			if fetch_type == ZEND_FETCH_CLASS_DEFAULT {
				result.SetOpType(IS_CONST)
				result.GetConstant().SetString(ZendResolveClassName(name, ZEND_NAME_FQ))
			} else {
				ZendEnsureValidClassFetchType(fetch_type)
				result.SetOpType(IS_UNUSED)
				result.GetOp().SetNum(fetch_type | fetch_flags)
			}
			// types.ZendStringReleaseEx(name, 0)
		} else {
			var opline *types.ZendOp = ZendEmitOp(result, ZEND_FETCH_CLASS, nil, &name_node)
			opline.GetOp1().SetNum(ZEND_FETCH_CLASS_DEFAULT | fetch_flags)
		}
		return
	}

	/* Fully qualified names are always default refs */

	if name_ast.Attr() == ZEND_NAME_FQ {
		result.SetOpType(IS_CONST)
		result.GetConstant().SetString(ZendResolveClassNameAst(name_ast))
		return
	}
	fetch_type = ZendGetClassFetchType(ZendAstGetStrVal(name_ast))
	if ZEND_FETCH_CLASS_DEFAULT == fetch_type {
		result.SetOpType(IS_CONST)
		result.GetConstant().SetString(ZendResolveClassNameAst(name_ast))
	} else {
		ZendEnsureValidClassFetchType(fetch_type)
		result.SetOpType(IS_UNUSED)
		result.GetOp().SetNum(fetch_type | fetch_flags)
	}
}
func ZendTryCompileCv(result *Znode, ast *ZendAst) int {
	var nameAst *ZendAst = ast.Child(0)
	if nameAst.Kind() == ZEND_AST_ZVAL {
		var zv *types.Zval = nameAst.Val()
		var name string
		if zv.IsString() {
			name = zv.String()
		} else {
			name = operators.ZvalGetStrVal(zv)
		}
		if ZendIsAutoGlobal(name) {
			return types.FAILURE
		}
		result.SetOpType(IS_CV)
		result.GetOp().SetVar(LookupCv(name))
		return types.SUCCESS
	}
	return types.FAILURE
}
func (compiler *Compiler) CompileSimpleVarNoCv(result *Znode, ast *ZendAst, type_ uint32, delayed int) *types.ZendOp {
	var name_ast *ZendAst = ast.Child(0)
	var name_node Znode
	var opline *types.ZendOp
	compiler.CompileExpr(&name_node, name_ast)
	if name_node.GetOpType() == IS_CONST {
		operators.ConvertToString(name_node.GetConstant())
	}
	if delayed != 0 {
		opline = ZendDelayedEmitOp(result, ZEND_FETCH_R, &name_node, nil)
	} else {
		opline = ZendEmitOp(result, ZEND_FETCH_R, &name_node, nil)
	}
	if name_node.GetOpType() == IS_CONST && ZendIsAutoGlobal(name_node.GetConstant().String()) {
		opline.SetExtendedValue(ZEND_FETCH_GLOBAL)
	} else {
		opline.SetExtendedValue(ZEND_FETCH_LOCAL)
	}
	ZendAdjustForFetchType(opline, result, type_)
	return opline
}
func IsThisFetch(ast *ZendAst) bool {
	if ast.Kind() == ZEND_AST_VAR && ast.Child(0).Kind() == ZEND_AST_ZVAL {
		var name *types.Zval = ast.Child(0).Val()
		return name.IsString() && name.String() == "this"
	}
	return false
}
func (compiler *Compiler) CompileSimpleVar(result *Znode, ast *ZendAst, typ uint32, delayed int) *types.ZendOp {
	if IsThisFetch(ast) {
		var opline *types.ZendOp = ZendEmitOp(result, ZEND_FETCH_THIS, nil, nil)
		if typ == BP_VAR_R || typ == BP_VAR_IS {
			opline.SetResultType(IS_TMP_VAR)
			result.SetOpType(IS_TMP_VAR)
		}
		CG__().GetActiveOpArray().SetIsUsesThis(true)
		return opline
	} else if ZendTryCompileCv(result, ast) == types.FAILURE {
		return compiler.CompileSimpleVarNoCv(result, ast, typ, delayed)
	}
	return nil
}
func ZendSeparateIfCallAndWrite(node *Znode, ast *ZendAst, type_ uint32) {
	if type_ != BP_VAR_R && type_ != BP_VAR_IS && ZendIsCall(ast) != 0 {
		if node.GetOpType() == IS_VAR {
			var opline *types.ZendOp = ZendEmitOp(nil, ZEND_SEPARATE, node, nil)
			opline.SetResultType(IS_VAR)
			opline.GetResult().SetVar(opline.GetOp1().GetVar())
		} else {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use result of built-in function in write context")
		}
	}
}
func (compiler *Compiler) EmitAssignZnode(var_ast *ZendAst, value_node *Znode) {
	var dummyNode Znode
	var assignAst *ZendAst = AstCreate(ZEND_AST_ASSIGN, var_ast, ZendAstCreateZnode(value_node))
	compiler.CompileAssign(&dummyNode, assignAst)
	ZendDoFree(&dummyNode)
}
func (compiler *Compiler) DelayedCompileDim(result *Znode, ast *ZendAst, type_ uint32) *types.ZendOp {
	if ast.Attr() == ZEND_DIM_ALTERNATIVE_SYNTAX {
		faults.Error(faults.E_DEPRECATED, "Array and string offset access syntax with curly braces is deprecated")
	}
	var var_ast *ZendAst = ast.Child(0)
	var dim_ast *ZendAst = ast.Child(1)
	var opline *types.ZendOp
	var var_node Znode
	var dim_node Znode
	opline = compiler.DelayedCompileVar(&var_node, var_ast, type_, 0)
	if opline != nil && type_ == BP_VAR_W && (opline.GetOpcode() == ZEND_FETCH_STATIC_PROP_W || opline.GetOpcode() == ZEND_FETCH_OBJ_W) {
		opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_DIM_WRITE)
	}
	ZendSeparateIfCallAndWrite(&var_node, var_ast, type_)
	if dim_ast == nil {
		if type_ == BP_VAR_R || type_ == BP_VAR_IS {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use [] for reading")
		}
		if type_ == BP_VAR_UNSET {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use [] for unsetting")
		}
		dim_node.SetOpType(IS_UNUSED)
	} else {
		compiler.CompileExpr(&dim_node, dim_ast)
	}
	opline = ZendDelayedEmitOp(result, ZEND_FETCH_DIM_R, &var_node, &dim_node)
	ZendAdjustForFetchType(opline, result, type_)
	if dim_node.GetOpType() == IS_CONST {
		ZendHandleNumericDim(opline, &dim_node)
	}
	return opline
}
func (compiler *Compiler) CompileDim(result *Znode, ast *ZendAst, type_ uint32) *types.ZendOp {
	return ZendDelayedCompileBlock(func() {
		compiler.DelayedCompileDim(result, ast, type_)
	})
}
func (compiler *Compiler) DelayedCompileProp(result *Znode, ast *ZendAst, type_ uint32) *types.ZendOp {
	var obj_ast *ZendAst = ast.Child(0)
	var prop_ast *ZendAst = ast.Child(1)
	var obj_node Znode
	var prop_node Znode
	var opline *types.ZendOp
	if IsThisFetch(obj_ast) {
		obj_node.SetOpType(IS_UNUSED)
		CG__().GetActiveOpArray().SetIsUsesThis(true)
	} else {
		opline = compiler.DelayedCompileVar(&obj_node, obj_ast, type_, 0)
		if opline != nil && type_ == BP_VAR_W && (opline.GetOpcode() == ZEND_FETCH_STATIC_PROP_W || opline.GetOpcode() == ZEND_FETCH_OBJ_W) {
			opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_OBJ_WRITE)
		}
		ZendSeparateIfCallAndWrite(&obj_node, obj_ast, type_)
	}
	compiler.CompileExpr(&prop_node, prop_ast)
	opline = ZendDelayedEmitOp(result, ZEND_FETCH_OBJ_R, &obj_node, &prop_node)
	if opline.GetOp2Type() == IS_CONST {
		operators.ConvertToString(CT_CONSTANT(opline.GetOp2()))
		opline.SetExtendedValue(ZendAllocCacheSlots(3))
	}
	ZendAdjustForFetchType(opline, result, type_)
	return opline
}
func (compiler *Compiler) CompileProp(result *Znode, ast *ZendAst, type_ uint32, by_ref int) *types.ZendOp {
	return ZendDelayedCompileBlock(func() {
		var opline *types.ZendOp = compiler.DelayedCompileProp(result, ast, type_)
		if by_ref != 0 {
			opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_REF)
		}
	})
}
func (compiler *Compiler) CompileStaticProp(result *Znode, ast *ZendAst, type_ uint32, by_ref int, delayed int) *types.ZendOp {
	var class_ast *ZendAst = ast.Child(0)
	var prop_ast *ZendAst = ast.Child(1)
	var class_node Znode
	var prop_node Znode
	var opline *types.ZendOp
	compiler.CompileClassRef(&class_node, class_ast, ZEND_FETCH_CLASS_EXCEPTION)
	compiler.CompileExpr(&prop_node, prop_ast)
	if delayed != 0 {
		opline = ZendDelayedEmitOp(result, ZEND_FETCH_STATIC_PROP_R, &prop_node, nil)
	} else {
		opline = ZendEmitOp(result, ZEND_FETCH_STATIC_PROP_R, &prop_node, nil)
	}
	if opline.GetOp1Type() == IS_CONST {
		operators.ConvertToString(CT_CONSTANT(opline.GetOp1()))
		opline.SetExtendedValue(ZendAllocCacheSlots(3))
	}
	if class_node.GetOpType() == IS_CONST {
		opline.SetOp2Type(IS_CONST)
		opline.GetOp2().SetConstant(ZendAddClassNameLiteral(class_node.GetConstant().String()))
		if opline.GetOp1Type() != IS_CONST {
			opline.SetExtendedValue(ZendAllocCacheSlot())
		}
	} else {
		opline.SetOp2Type(class_node.GetOpType())
		if class_node.GetOpType() == IS_CONST {
			opline.GetOp2().SetConstant(ZendAddLiteral(class_node.GetConstant()))
		} else {
			opline.SetOp2(class_node.GetOp())
		}
	}
	if by_ref != 0 && (type_ == BP_VAR_W || type_ == BP_VAR_FUNC_ARG) {
		opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_FETCH_REF)
	}
	ZendAdjustForFetchType(opline, result, type_)
	return opline
}
func ZendVerifyListAssignTarget(var_ast *ZendAst, old_style bool) {
	if var_ast.Kind() == ZEND_AST_ARRAY {
		if var_ast.Attr() == ZEND_ARRAY_SYNTAX_LONG {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot assign to array(), use [] instead")
		}
		if old_style != var_ast.Attr() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot mix [] and list()")
		}
	} else if ZendCanWriteToVariable(var_ast) == 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Assignments can only happen to writable values")
	}
}
func ZendPropagateListRefs(ast *ZendAst) bool {
	var list *ZendAstList = ast.AsAstList()
	var has_refs bool = 0
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.Children()[i]
		if elem_ast != nil {
			var var_ast *ZendAst = elem_ast.Child(0)
			if var_ast.Kind() == ZEND_AST_ARRAY {
				elem_ast.SetAttr(ZendPropagateListRefs(var_ast))
			}
			has_refs |= elem_ast.Attr()
		}
	}
	return has_refs
}
func (compiler *Compiler) CompileListAssign(result *Znode, ast *ZendAst, expr_node *Znode, old_style bool) {
	var list *ZendAstList = ast.AsAstList()
	var i uint32
	var has_elems bool = 0
	var is_keyed bool = list.GetChildren() > 0 && list.Children()[0] != nil && list.Children()[0].Children()[1] != nil
	//if list.GetChildren() != 0 && expr_node.GetOpType() == IS_CONST && expr_node.GetConstant().IsString() {
	//	ZvalMakeInternedString(expr_node.GetConstant())
	//}
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.Children()[i]
		var var_ast *ZendAst
		var key_ast *ZendAst
		var fetch_result Znode
		var dim_node Znode
		var opline *types.ZendOp
		if elem_ast == nil {
			if is_keyed != 0 {
				faults.Error(faults.E_COMPILE_ERROR, "Cannot use empty array entries in keyed array assignment")
			} else {
				continue
			}
		}
		if elem_ast.Kind() == ZEND_AST_UNPACK {
			faults.Error(faults.E_COMPILE_ERROR, "Spread operator is not supported in assignments")
		}
		var_ast = elem_ast.Child(0)
		key_ast = elem_ast.Child(1)
		has_elems = 1
		if is_keyed != 0 {
			if key_ast == nil {
				faults.Error(faults.E_COMPILE_ERROR, "Cannot mix keyed and unkeyed array entries in assignments")
			}
			compiler.CompileExpr(&dim_node, key_ast)
		} else {
			if key_ast != nil {
				faults.Error(faults.E_COMPILE_ERROR, "Cannot mix keyed and unkeyed array entries in assignments")
			}
			dim_node.SetOpType(IS_CONST)
			dim_node.GetConstant().SetLong(i)
		}
		if expr_node.GetOpType() == IS_CONST {
			//expr_node.GetConstant().TryAddRefcount()
		}
		ZendVerifyListAssignTarget(var_ast, old_style)
		opline = ZendEmitOp(&fetch_result, lang.CondF1(elem_ast.Attr() != 0, func() __auto__ {
			if expr_node.GetOpType() == IS_CV {
				return ZEND_FETCH_DIM_W
			} else {
				return ZEND_FETCH_LIST_W
			}
		}, ZEND_FETCH_LIST_R), expr_node, &dim_node)
		if dim_node.GetOpType() == IS_CONST {
			ZendHandleNumericDim(opline, &dim_node)
		}
		if var_ast.Kind() == ZEND_AST_ARRAY {
			if elem_ast.Attr() != 0 {
				ZendEmitOp(&fetch_result, ZEND_MAKE_REF, &fetch_result, nil)
			}
			compiler.CompileListAssign(nil, var_ast, &fetch_result, var_ast.Attr())
		} else if elem_ast.Attr() != 0 {
			compiler.EmitAssignRefZnode(var_ast, &fetch_result)
		} else {
			compiler.EmitAssignZnode(var_ast, &fetch_result)
		}
	}
	if has_elems == 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use empty list")
	}
	if result != nil {
		*result = *expr_node
	} else {
		ZendDoFree(expr_node)
	}
}
func ZendEnsureWritableVariable(ast *ZendAst) {
	if ast.Kind() == ZEND_AST_CALL {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Can't use function return value in write context")
	}
	if ast.Kind() == ZEND_AST_METHOD_CALL || ast.Kind() == ZEND_AST_STATIC_CALL {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Can't use method return value in write context")
	}
}
func ZendIsAssignToSelf(var_ast *ZendAst, expr_ast *ZendAst) bool {
	if expr_ast.Kind() != ZEND_AST_VAR || expr_ast.Child(0).Kind() != ZEND_AST_ZVAL {
		return 0
	}
	for ZendIsVariable(var_ast) != 0 && var_ast.Kind() != ZEND_AST_VAR {
		var_ast = var_ast.Children()[0]
	}
	if var_ast.Kind() != ZEND_AST_VAR || var_ast.Children()[0].Kind() != ZEND_AST_ZVAL {
		return 0
	}
	var name1 *types.String = operators.ZvalGetString(ZendAstGetZval(var_ast.Children()[0]))
	var name2 *types.String = operators.ZvalGetString(ZendAstGetZval(expr_ast.Children()[0]))
	var result = name1.GetStr() == name2.GetStr()
	return types.IntBool(result)
}
func (compiler *Compiler) CompileAssign(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.Children()[0]
	var expr_ast *ZendAst = ast.Child(1)
	var var_node Znode
	var expr_node Znode
	var opline *types.ZendOp
	var offset uint32
	if IsThisFetch(var_ast) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot re-assign $this")
	}
	ZendEnsureWritableVariable(var_ast)
	switch var_ast.Kind() {
	case ZEND_AST_VAR:
		ZendDelayedCompileBlock(func() {
			compiler.DelayedCompileVar(&var_node, var_ast, BP_VAR_W, 0)
			compiler.CompileExpr(&expr_node, expr_ast)
		})
		ZendEmitOp(result, ZEND_ASSIGN, &var_node, &expr_node)
		return
	case ZEND_AST_STATIC_PROP:
		opline = ZendDelayedCompileBlock(func() {
			compiler.DelayedCompileVar(result, var_ast, BP_VAR_W, 0)
			compiler.CompileExpr(&expr_node, expr_ast)
		})
		opline.SetOpcode(ZEND_ASSIGN_STATIC_PROP)
		ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_DIM:
		opline = ZendDelayedCompileBlock(func() {
			compiler.DelayedCompileDim(result, var_ast, BP_VAR_W)
			if ZendIsAssignToSelf(var_ast, expr_ast) != 0 && !IsThisFetch(expr_ast) {
				/* $a[0] = $a should evaluate the right $a first */
				var cv_node Znode
				if ZendTryCompileCv(&cv_node, expr_ast) == types.FAILURE {
					compiler.CompileSimpleVarNoCv(&expr_node, expr_ast, BP_VAR_R, 0)
				} else {
					ZendEmitOpTmp(&expr_node, ZEND_QM_ASSIGN, &cv_node, nil)
				}
			} else {
				compiler.CompileExpr(&expr_node, expr_ast)
			}
		})
		opline.SetOpcode(ZEND_ASSIGN_DIM)
		opline = ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_PROP:
		opline = ZendDelayedCompileBlock(func() {
			compiler.DelayedCompileProp(result, var_ast, BP_VAR_W)
			compiler.CompileExpr(&expr_node, expr_ast)
		})
		opline.SetOpcode(ZEND_ASSIGN_OBJ)
		ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_ARRAY:
		if ZendPropagateListRefs(var_ast) != 0 {
			if ZendIsVariableOrCall(expr_ast) == 0 {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot assign reference to non referencable value")
			}
			compiler.CompileVar(&expr_node, expr_ast, BP_VAR_W, 1)

			/* MAKE_REF is usually not necessary for CVs. However, if there are
			 * self-assignments, this forces the RHS to evaluate first. */

			ZendEmitOp(&expr_node, ZEND_MAKE_REF, &expr_node, nil)

			/* MAKE_REF is usually not necessary for CVs. However, if there are
			 * self-assignments, this forces the RHS to evaluate first. */

		} else {
			if expr_ast.Kind() == ZEND_AST_VAR {

				/* list($a, $b) = $a should evaluate the right $a first */

				var cv_node Znode
				if ZendTryCompileCv(&cv_node, expr_ast) == types.FAILURE {
					compiler.CompileSimpleVarNoCv(&expr_node, expr_ast, BP_VAR_R, 0)
				} else {
					ZendEmitOpTmp(&expr_node, ZEND_QM_ASSIGN, &cv_node, nil)
				}
			} else {
				compiler.CompileExpr(&expr_node, expr_ast)
			}
		}
		compiler.CompileListAssign(result, var_ast, &expr_node, var_ast.Attr())
		return
	default:

	}
}
