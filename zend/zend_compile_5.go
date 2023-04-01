package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZendCompileIf(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var jmp_opnums *uint32 = nil
	if list.GetChildren() > 1 {
		jmp_opnums = SafeEmalloc(b.SizeOf("uint32_t"), list.GetChildren()-1, 0)
	}
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.GetChild()[i]
		var cond_ast *ZendAst = elem_ast.GetChild()[0]
		var stmt_ast *ZendAst = elem_ast.GetChild()[1]
		if cond_ast != nil {
			var cond_node Znode
			var opnum_jmpz uint32
			ZendCompileExpr(&cond_node, cond_ast)
			opnum_jmpz = ZendEmitCondJump(ZEND_JMPZ, &cond_node, 0)
			ZendCompileStmt(stmt_ast)
			if i != list.GetChildren()-1 {
				jmp_opnums[i] = ZendEmitJump(0)
			}
			ZendUpdateJumpTargetToNext(opnum_jmpz)
		} else {

			/* "else" can only occur as last element. */

			b.Assert(i == list.GetChildren()-1)
			ZendCompileStmt(stmt_ast)
		}
	}
	if list.GetChildren() > 1 {
		for i = 0; i < list.GetChildren()-1; i++ {
			ZendUpdateJumpTargetToNext(jmp_opnums[i])
		}
		Efree(jmp_opnums)
	}
}
func DetermineSwitchJumptableType(cases *ZendAstList) types.ZendUchar {
	var i uint32
	var common_type types.ZendUchar = types.IS_UNDEF
	for i = 0; i < cases.GetChildren(); i++ {
		var case_ast *ZendAst = cases.GetChild()[i]
		var cond_ast **ZendAst = case_ast.GetChild()[0]
		var cond_zv *types.Zval
		if case_ast.GetChild()[0] == nil {

			/* Skip default clause */

			continue

			/* Skip default clause */

		}
		ZendEvalConstExpr(cond_ast)
		if cond_ast.GetKind() != ZEND_AST_ZVAL {

			/* Non-constant case */

			return types.IS_UNDEF

			/* Non-constant case */

		}
		cond_zv = ZendAstGetZval(case_ast.GetChild()[0])
		if cond_zv.GetType() != types.IS_LONG && cond_zv.GetType() != types.IS_STRING {

			/* We only optimize switched on integers and strings */

			return types.IS_UNDEF

			/* We only optimize switched on integers and strings */

		}
		if common_type == types.IS_UNDEF {
			common_type = cond_zv.GetType()
		} else if common_type != cond_zv.GetType() {

			/* Non-uniform case types */

			return types.IS_UNDEF

			/* Non-uniform case types */

		}
		if cond_zv.IsString() && IsNumericString(cond_zv.GetStr().GetStr(), nil, nil, 0) != 0 {

			/* Numeric strings cannot be compared with a simple hash lookup */

			return types.IS_UNDEF

			/* Numeric strings cannot be compared with a simple hash lookup */

		}
	}
	return common_type
}
func ShouldUseJumptable(cases *ZendAstList, jumptable_type types.ZendUchar) types.ZendBool {
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_NO_JUMPTABLES) != 0 {
		return 0
	}

	/* Thresholds are chosen based on when the average switch time for equidistributed
	 * input becomes smaller when using the jumptable optimization. */

	if jumptable_type == types.IS_LONG {
		return cases.GetChildren() >= 5
	} else {
		b.Assert(jumptable_type == types.IS_STRING)
		return cases.GetChildren() >= 2
	}

	/* Thresholds are chosen based on when the average switch time for equidistributed
	 * input becomes smaller when using the jumptable optimization. */
}
func ZendCompileSwitch(ast *ZendAst) {
	var expr_ast *ZendAst = ast.GetChild()[0]
	var cases *ZendAstList = ZendAstGetList(ast.GetChild()[1])
	var i uint32
	var has_default_case types.ZendBool = 0
	var expr_node Znode
	var case_node Znode
	var opline *ZendOp
	var jmpnz_opnums *uint32
	var opnum_default_jmp uint32
	var opnum_switch uint32 = uint32 - 1
	var jumptable_type types.ZendUchar
	var jumptable *types.Array = nil
	ZendCompileExpr(&expr_node, expr_ast)
	ZendBeginLoop(ZEND_FREE, &expr_node, 1)
	case_node.SetOpType(IS_TMP_VAR)
	case_node.GetOp().SetVar(GetTemporaryVariable())
	jumptable_type = DetermineSwitchJumptableType(cases)
	if jumptable_type != types.IS_UNDEF && ShouldUseJumptable(cases, jumptable_type) != 0 {
		var jumptable_op Znode
		ALLOC_HASHTABLE(jumptable)
		jumptable = types.MakeArrayEx(cases.GetChildren(), nil, 0)
		jumptable_op.SetOpType(IS_CONST)
		jumptable_op.GetConstant().SetArray(jumptable)
		opline = ZendEmitOp(nil, b.Cond(jumptable_type == types.IS_LONG, ZEND_SWITCH_LONG, ZEND_SWITCH_STRING), &expr_node, &jumptable_op)
		if opline.GetOp1Type() == IS_CONST {
			CT_CONSTANT(opline.GetOp1()).TryAddRefcount()
		}
		opnum_switch = opline - CG__().GetActiveOpArray().GetOpcodes()
	}
	jmpnz_opnums = SafeEmalloc(b.SizeOf("uint32_t"), cases.GetChildren(), 0)
	for i = 0; i < cases.GetChildren(); i++ {
		var case_ast *ZendAst = cases.GetChild()[i]
		var cond_ast *ZendAst = case_ast.GetChild()[0]
		var cond_node Znode
		if cond_ast == nil {
			if has_default_case != 0 {
				CG__().SetZendLineno(case_ast.GetLineno())
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Switch statements may only contain one default clause")
			}
			has_default_case = 1
			continue
		}
		ZendCompileExpr(&cond_node, cond_ast)
		if expr_node.GetOpType() == IS_CONST && expr_node.GetConstant().IsFalse() {
			jmpnz_opnums[i] = ZendEmitCondJump(ZEND_JMPZ, &cond_node, 0)
		} else if expr_node.GetOpType() == IS_CONST && expr_node.GetConstant().IsTrue() {
			jmpnz_opnums[i] = ZendEmitCondJump(ZEND_JMPNZ, &cond_node, 0)
		} else {
			opline = ZendEmitOp(nil, b.Cond((expr_node.GetOpType()&(IS_VAR|IS_TMP_VAR)) != 0, ZEND_CASE, ZEND_IS_EQUAL), &expr_node, &cond_node)
			opline.SetResultType(case_node.GetOpType())
			if case_node.GetOpType() == IS_CONST {
				opline.GetResult().SetConstant(ZendAddLiteral(case_node.GetConstant()))
			} else {
				opline.SetResult(case_node.GetOp())
			}
			if opline.GetOp1Type() == IS_CONST {
				CT_CONSTANT(opline.GetOp1()).TryAddRefcount()
			}
			jmpnz_opnums[i] = ZendEmitCondJump(ZEND_JMPNZ, &case_node, 0)
		}
	}
	opnum_default_jmp = ZendEmitJump(0)
	for i = 0; i < cases.GetChildren(); i++ {
		var case_ast *ZendAst = cases.GetChild()[i]
		var cond_ast *ZendAst = case_ast.GetChild()[0]
		var stmt_ast *ZendAst = case_ast.GetChild()[1]
		if cond_ast != nil {
			ZendUpdateJumpTargetToNext(jmpnz_opnums[i])
			if jumptable != nil {
				var cond_zv *types.Zval = ZendAstGetZval(cond_ast)
				var jmp_target types.Zval
				jmp_target.SetLong(GetNextOpNumber())
				b.Assert(cond_zv.IsType(jumptable_type))
				if cond_zv.IsLong() {
					jumptable.IndexAdd(cond_zv.GetLval(), &jmp_target)
				} else {
					b.Assert(cond_zv.IsString())
					jumptable.KeyAdd(cond_zv.GetStr().GetStr(), &jmp_target)
				}
			}
		} else {
			ZendUpdateJumpTargetToNext(opnum_default_jmp)
			if jumptable != nil {
				b.Assert(opnum_switch != uint32-1)
				opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_switch]
				opline.SetExtendedValue(GetNextOpNumber())
			}
		}
		ZendCompileStmt(stmt_ast)
	}
	if has_default_case == 0 {
		ZendUpdateJumpTargetToNext(opnum_default_jmp)
		if jumptable != nil {
			opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_switch]
			opline.SetExtendedValue(GetNextOpNumber())
		}
	}
	ZendEndLoop(GetNextOpNumber(), &expr_node)
	if (expr_node.GetOpType() & (IS_VAR | IS_TMP_VAR)) != 0 {
		opline = ZendEmitOp(nil, ZEND_FREE, &expr_node, nil)
		opline.SetExtendedValue(ZEND_FREE_SWITCH)
	} else if expr_node.GetOpType() == IS_CONST {
		ZvalPtrDtorNogc(expr_node.GetConstant())
	}
	Efree(jmpnz_opnums)
}
func ZendCompileTry(ast *ZendAst) {
	var try_ast *ZendAst = ast.GetChild()[0]
	var catches *ZendAstList = ZendAstGetList(ast.GetChild()[1])
	var finally_ast *ZendAst = ast.GetChild()[2]
	var i uint32
	var j uint32
	var opline *ZendOp
	var try_catch_offset uint32
	var jmp_opnums *uint32 = SafeEmalloc(b.SizeOf("uint32_t"), catches.GetChildren(), 0)
	var orig_fast_call_var uint32 = CG__().GetContext().GetFastCallVar()
	var orig_try_catch_offset uint32 = CG__().GetContext().GetTryCatchOffset()
	if catches.GetChildren() == 0 && finally_ast == nil {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use try without catch or finally")
	}

	/* label: try { } must not be equal to try { label: } */

	if CG__().GetContext().GetLabels() != nil {
		var label *ZendLabel
		var __ht *types.Array = CG__().GetContext().GetLabels()
		for _, _p := range __ht.ForeachDataReserve() {
			var _z types.Zval = _p.GetVal()

			label = _z.GetPtr()
			if label.GetOplineNum() == GetNextOpNumber() {
				ZendEmitOp(nil, ZEND_NOP, nil, nil)
			}
			break
		}
	}
	try_catch_offset = ZendAddTryElement(GetNextOpNumber())
	if finally_ast != nil {
		var fast_call ZendLoopVar
		if !CG__().GetActiveOpArray().IsHasFinallyBlock() {
			CG__().GetActiveOpArray().SetIsHasFinallyBlock(true)
		}
		CG__().GetContext().SetFastCallVar(GetTemporaryVariable())

		/* Push FAST_CALL on unwind stack */

		fast_call.SetOpcode(ZEND_FAST_CALL)
		fast_call.SetVarType(IS_TMP_VAR)
		fast_call.SetVarNum(CG__().GetContext().GetFastCallVar())
		fast_call.SetTryCatchOffset(try_catch_offset)
		CG__().GetLoopVarStack().Push(&fast_call)
	}
	CG__().GetContext().SetTryCatchOffset(try_catch_offset)
	ZendCompileStmt(try_ast)
	if catches.GetChildren() != 0 {
		jmp_opnums[0] = ZendEmitJump(0)
	}
	for i = 0; i < catches.GetChildren(); i++ {
		var catch_ast *ZendAst = catches.GetChild()[i]
		var classes *ZendAstList = ZendAstGetList(catch_ast.GetChild()[0])
		var var_ast *ZendAst = catch_ast.GetChild()[1]
		var stmt_ast *ZendAst = catch_ast.GetChild()[2]
		var var_name *types.String = ZvalMakeInternedString(ZendAstGetZval(var_ast))
		var is_last_catch types.ZendBool = i+1 == catches.GetChildren()
		var jmp_multicatch *uint32 = SafeEmalloc(b.SizeOf("uint32_t"), classes.GetChildren()-1, 0)
		var opnum_catch uint32 = uint32 - 1
		CG__().SetZendLineno(catch_ast.GetLineno())
		for j = 0; j < classes.GetChildren(); j++ {
			var class_ast *ZendAst = classes.GetChild()[j]
			var is_last_class types.ZendBool = j+1 == classes.GetChildren()
			if ZendIsConstDefaultClassRef(class_ast) == 0 {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Bad class name in the catch statement")
			}
			opnum_catch = GetNextOpNumber()
			if i == 0 && j == 0 {
				CG__().GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetCatchOp(opnum_catch)
			}
			opline = GetNextOp()
			opline.SetOpcode(ZEND_CATCH)
			opline.SetOp1Type(IS_CONST)
			opline.GetOp1().SetConstant(ZendAddClassNameLiteral(ZendResolveClassNameAst(class_ast)))
			opline.SetExtendedValue(ZendAllocCacheSlot())
			if types.ZendStringEqualsLiteral(var_name, "this") {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot re-assign $this")
			}
			opline.SetResultType(IS_CV)
			opline.GetResult().SetVar(LookupCv(var_name))
			if is_last_catch != 0 && is_last_class != 0 {
				opline.SetExtendedValue(opline.GetExtendedValue() | ZEND_LAST_CATCH)
			}
			if is_last_class == 0 {
				jmp_multicatch[j] = ZendEmitJump(0)
				opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_catch]
				opline.GetOp2().SetOplineNum(GetNextOpNumber())
			}
		}
		for j = 0; j < classes.GetChildren()-1; j++ {
			ZendUpdateJumpTargetToNext(jmp_multicatch[j])
		}
		Efree(jmp_multicatch)
		ZendCompileStmt(stmt_ast)
		if is_last_catch == 0 {
			jmp_opnums[i+1] = ZendEmitJump(0)
		}
		b.Assert(opnum_catch != uint32-1 && "Should have at least one class")
		opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_catch]
		if is_last_catch == 0 {
			opline.GetOp2().SetOplineNum(GetNextOpNumber())
		}
	}
	for i = 0; i < catches.GetChildren(); i++ {
		ZendUpdateJumpTargetToNext(jmp_opnums[i])
	}
	if finally_ast != nil {
		var discard_exception ZendLoopVar
		var opnum_jmp uint32 = GetNextOpNumber() + 1

		/* Pop FAST_CALL from unwind stack */

		CG__().GetLoopVarStack().DelTop()

		/* Push DISCARD_EXCEPTION on unwind stack */

		discard_exception.SetOpcode(ZEND_DISCARD_EXCEPTION)
		discard_exception.SetVarType(IS_TMP_VAR)
		discard_exception.SetVarNum(CG__().GetContext().GetFastCallVar())
		CG__().GetLoopVarStack().Push(&discard_exception)
		CG__().SetZendLineno(finally_ast.GetLineno())
		opline = ZendEmitOp(nil, ZEND_FAST_CALL, nil, nil)
		opline.GetOp1().SetNum(try_catch_offset)
		opline.SetResultType(IS_TMP_VAR)
		opline.GetResult().SetVar(CG__().GetContext().GetFastCallVar())
		ZendEmitOp(nil, ZEND_JMP, nil, nil)
		ZendCompileStmt(finally_ast)
		CG__().GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetFinallyOp(opnum_jmp + 1)
		CG__().GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetFinallyEnd(GetNextOpNumber())
		opline = ZendEmitOp(nil, ZEND_FAST_RET, nil, nil)
		opline.SetOp1Type(IS_TMP_VAR)
		opline.GetOp1().SetVar(CG__().GetContext().GetFastCallVar())
		opline.GetOp2().SetNum(orig_try_catch_offset)
		ZendUpdateJumpTargetToNext(opnum_jmp)
		CG__().GetContext().SetFastCallVar(orig_fast_call_var)

		/* Pop DISCARD_EXCEPTION from unwind stack */

		CG__().GetLoopVarStack().DelTop()

		/* Pop DISCARD_EXCEPTION from unwind stack */

	}
	CG__().GetContext().SetTryCatchOffset(orig_try_catch_offset)
	Efree(jmp_opnums)
}
func ZendDeclareIsFirstStatement(ast *ZendAst) int {
	var i uint32 = 0
	var file_ast *ZendAstList = ZendAstGetList(CG__().GetAst())

	/* Check to see if this declare is preceded only by declare statements */

	for i < file_ast.GetChildren() {
		if file_ast.GetChild()[i] == ast {
			return types.SUCCESS
		} else if file_ast.GetChild()[i] == nil {

			/* Empty statements are not allowed prior to a declare */

			return types.FAILURE

			/* Empty statements are not allowed prior to a declare */

		} else if file_ast.GetChild()[i].GetKind() != ZEND_AST_DECLARE {

			/* declares can only be preceded by other declares */

			return types.FAILURE

			/* declares can only be preceded by other declares */

		}
		i++
	}
	return types.FAILURE
}
func ZendCompileDeclare(ast *ZendAst) {
	var declares *ZendAstList = ZendAstGetList(ast.GetChild()[0])
	var stmt_ast *ZendAst = ast.GetChild()[1]
	var orig_declarables ZendDeclarables = FC__().GetDeclarables()
	var i uint32
	for i = 0; i < declares.GetChildren(); i++ {
		var declare_ast *ZendAst = declares.GetChild()[i]
		var name_ast *ZendAst = declare_ast.GetChild()[0]
		var value_ast *ZendAst = declare_ast.GetChild()[1]
		var name *types.String = ZendAstGetStr(name_ast)
		if value_ast.GetKind() != ZEND_AST_ZVAL {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "declare(%s) value must be a literal", name.GetVal())
		}
		if types.ZendStringEqualsLiteralCi(name, "ticks") {
			var value_zv types.Zval
			ZendConstExprToZval(&value_zv, value_ast)
			FC__().GetDeclarables().SetTicks(ZvalGetLong(&value_zv))
			ZvalPtrDtorNogc(&value_zv)
		} else if types.ZendStringEqualsLiteralCi(name, "encoding") {
			if types.FAILURE == ZendDeclareIsFirstStatement(ast) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Encoding declaration pragma must be "+"the very first statement in the script")
			}
		} else if types.ZendStringEqualsLiteralCi(name, "strict_types") {
			var value_zv types.Zval
			if types.FAILURE == ZendDeclareIsFirstStatement(ast) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "strict_types declaration must be "+"the very first statement in the script")
			}
			if ast.GetChild()[1] != nil {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "strict_types declaration must not "+"use block mode")
			}
			ZendConstExprToZval(&value_zv, value_ast)
			if value_zv.GetType() != types.IS_LONG || value_zv.GetLval() != 0 && value_zv.GetLval() != 1 {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "strict_types declaration must have 0 or 1 as its value")
			}
			if value_zv.GetLval() == 1 {
				CG__().GetActiveOpArray().SetIsStrictTypes(true)
			}
		} else {
			faults.Error(faults.E_COMPILE_WARNING, "Unsupported declare '%s'", name.GetVal())
		}
	}
	if stmt_ast != nil {
		ZendCompileStmt(stmt_ast)
		FC__().SetDeclarables(orig_declarables)
	}
}
func ZendCompileStmtList(ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		ZendCompileStmt(list.GetChild()[i])
	}
}
func ZendCompileTypename(ast *ZendAst, force_allow_null types.ZendBool) types.ZendType {
	var allow_null types.ZendBool = force_allow_null
	var orig_ast_attr ZendAstAttr = ast.GetAttr()
	var type_ types.ZendType
	if (ast.GetAttr() & ZEND_TYPE_NULLABLE) != 0 {
		allow_null = 1
		ast.SetAttr(ast.GetAttr() &^ ZEND_TYPE_NULLABLE)
	}
	if ast.GetKind() == ZEND_AST_TYPE {
		return types.ZEND_TYPE_ENCODE(ast.GetAttr(), allow_null)
	} else {
		var class_name *types.String = ZendAstGetStr(ast)
		var type_code types.ZendUchar = ZendLookupBuiltinTypeByName(class_name)
		if type_code != 0 {
			if (ast.GetAttr() & ZEND_NAME_NOT_FQ) != ZEND_NAME_NOT_FQ {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Type declaration '%s' must be unqualified", ZendStringTolower(class_name).GetVal())
			}
			type_ = types.ZEND_TYPE_ENCODE(type_code, allow_null)
		} else {
			var fetch_type uint32 = ZendGetClassFetchTypeAst(ast)
			if fetch_type == ZEND_FETCH_CLASS_DEFAULT {
				class_name = ZendResolveClassNameAst(ast)
				ZendAssertValidClassName(class_name.GetStr())
			} else {
				ZendEnsureValidClassFetchType(fetch_type)
				class_name.AddRefcount()
			}
			type_ = types.ZEND_TYPE_ENCODE_CLASS(class_name, allow_null)
		}
	}
	ast.SetAttr(orig_ast_attr)
	return type_
}
func ZendCompileParams(ast *ZendAst, return_type_ast *ZendAst) {
	var list *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var op_array *types.ZendOpArray = CG__().GetActiveOpArray()
	var arg_infos *ZendArgInfo
	if return_type_ast != nil {

		/* Use op_array->arg_info[-1] for return type */

		arg_infos = SafeEmalloc(b.SizeOf("zend_arg_info"), list.GetChildren()+1, 0)
		*arg_infos = MakeZendReturnArgInfo(
			ZendCompileTypename(return_type_ast, 0),
			op_array.IsReturnReference(),
		)
		if arg_infos.GetType().Code() == types.IS_VOID && arg_infos.GetType().AllowNull() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Void type cannot be nullable")
		}
		arg_infos++
		op_array.SetIsHasReturnType(true)
	} else {
		if list.GetChildren() == 0 {
			return
		}
		arg_infos = SafeEmalloc(b.SizeOf("zend_arg_info"), list.GetChildren(), 0)
	}
	for i = 0; i < list.GetChildren(); i++ {
		var param_ast *ZendAst = list.GetChild()[i]
		var type_ast *ZendAst = param_ast.GetChild()[0]
		var var_ast *ZendAst = param_ast.GetChild()[1]
		var default_ast *ZendAst = param_ast.GetChild()[2]
		var name *types.String = ZvalMakeInternedString(ZendAstGetZval(var_ast))
		var is_ref types.ZendBool = (param_ast.GetAttr() & ZEND_PARAM_REF) != 0
		var is_variadic types.ZendBool = (param_ast.GetAttr() & ZEND_PARAM_VARIADIC) != 0
		var var_node Znode
		var default_node Znode
		var opcode types.ZendUchar
		var opline *ZendOp
		var arg_info *ZendArgInfo
		if ZendIsAutoGlobal(name) != 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot re-assign auto-global variable %s", name.GetVal())
		}
		var_node.SetOpType(IS_CV)
		var_node.GetOp().SetVar(LookupCv(name))
		if EX_VAR_TO_NUM(var_node.GetOp().GetVar()) != i {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Redefinition of parameter $%s", name.GetVal())
		} else if types.ZendStringEqualsLiteral(name, "this") {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use $this as parameter")
		}
		if op_array.IsVariadic() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Only the last parameter can be variadic")
		}
		if is_variadic != 0 {
			opcode = ZEND_RECV_VARIADIC
			default_node.SetOpType(IS_UNUSED)
			op_array.SetIsVariadic(true)
			if default_ast != nil {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Variadic parameter cannot have a default value")
			}
		} else if default_ast != nil {

			/* we cannot substitute constants here or it will break ReflectionParameter::getDefaultValueConstantName() and ReflectionParameter::isDefaultValueConstant() */

			var cops uint32 = CG__().GetCompilerOptions()
			CG__().SetCompilerOptions(CG__().GetCompilerOptions() | ZEND_COMPILE_NO_CONSTANT_SUBSTITUTION | ZEND_COMPILE_NO_PERSISTENT_CONSTANT_SUBSTITUTION)
			opcode = ZEND_RECV_INIT
			default_node.SetOpType(IS_CONST)
			ZendConstExprToZval(default_node.GetConstant(), default_ast)
			CG__().SetCompilerOptions(cops)
		} else {
			opcode = ZEND_RECV
			default_node.SetOpType(IS_UNUSED)
			op_array.SetRequiredNumArgs(i + 1)
		}
		arg_info = &arg_infos[i]
		*arg_info = MakeZendArgInfo(
			name.Copy(),
			/* TODO: Keep compatibility, but may be better reset "allow_null" ??? */
			types.ZEND_TYPE_ENCODE(0, 1),
			is_ref,
			is_variadic,
		)
		if type_ast != nil {
			var has_null_default types.ZendBool = default_ast != nil && (default_node.GetConstant().IsNull() || default_node.GetConstant().IsConstant() && types.Z_ASTVAL(default_node.GetConstant()).GetKind() == ZEND_AST_CONSTANT && strcasecmp(ZendAstGetConstantName(types.Z_ASTVAL(default_node.GetConstant())).GetVal(), "NULL") == 0)
			op_array.SetIsHasTypeHints(true)
			arg_info.SetType(ZendCompileTypename(type_ast, has_null_default))
			if arg_info.GetType().Code() == types.IS_VOID {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "void cannot be used as a parameter type")
			}
			if type_ast.GetKind() == ZEND_AST_TYPE {
				if arg_info.GetType().Code() == types.IS_ARRAY {
					if default_ast != nil && has_null_default == 0 && default_node.GetConstant().GetType() != types.IS_ARRAY && default_node.GetConstant().GetType() != types.IS_CONSTANT_AST {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters "+"with array type can only be an array or NULL")
					}
				} else if arg_info.GetType().Code() == types.IS_CALLABLE && default_ast != nil {
					if has_null_default == 0 && default_node.GetConstant().GetType() != types.IS_CONSTANT_AST {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters "+"with callable type can only be NULL")
					}
				}
			} else {
				if default_ast != nil && has_null_default == 0 && default_node.GetConstant().GetType() != types.IS_CONSTANT_AST {
					if arg_info.GetType().IsClass() {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters "+"with a class type can only be NULL")
					} else {
						switch arg_info.GetType().Code() {
						case types.IS_DOUBLE:
							if default_node.GetConstant().GetType() != types.IS_DOUBLE && default_node.GetConstant().GetType() != types.IS_LONG {
								faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters "+"with a float type can only be float, integer, or NULL")
							}
							ConvertToDouble(default_node.GetConstant())
						case types.IS_ITERABLE:
							if default_node.GetConstant().GetType() != types.IS_ARRAY {
								faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters "+"with iterable type can only be an array or NULL")
							}
						case types.IS_OBJECT:
							faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters "+"with an object type can only be NULL")
						default:
							if !(types.ZEND_SAME_FAKE_TYPE(arg_info.GetType().Code(), default_node.GetConstant().GetType())) {
								faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters "+"with a %s type can only be %s or NULL", types.ZendGetTypeByConst(arg_info.GetType().Code()), types.ZendGetTypeByConst(arg_info.GetType().Code()))
							}
						}
					}
				}
			}
		}
		opline = ZendEmitOp(nil, opcode, nil, &default_node)
		opline.SetResultType(var_node.GetOpType())
		if var_node.GetOpType() == IS_CONST {
			opline.GetResult().SetConstant(ZendAddLiteral(var_node.GetConstant()))
		} else {
			opline.SetResult(var_node.GetOp())
		}
		opline.GetOp1().SetNum(i + 1)
		if type_ast != nil {

			/* Allocate cache slot to speed-up run-time class resolution */

			if opline.GetOpcode() == ZEND_RECV_INIT {
				if arg_info.GetType().IsClass() {
					opline.SetExtendedValue(ZendAllocCacheSlot())
				}
			} else {
				if arg_info.GetType().IsClass() {
					opline.GetOp2().SetNum(op_array.GetCacheSize())
					op_array.SetCacheSize(op_array.GetCacheSize() + b.SizeOf("void *"))
				} else {
					opline.GetOp2().SetNum(-1)
				}
			}

			/* Allocate cache slot to speed-up run-time class resolution */

		} else {
			if opline.GetOpcode() != ZEND_RECV_INIT {
				opline.GetOp2().SetNum(-1)
			}
		}
	}

	/* These are assigned at the end to avoid uninitialized memory in case of an error */

	op_array.SetNumArgs(list.GetChildren())
	op_array.SetArgInfo(arg_infos)

	/* Don't count the variadic argument */

	if op_array.IsVariadic() {
		op_array.GetNumArgs()--
	}
}
