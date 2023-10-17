package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func (compiler *Compiler) CompileIf(ast *ZendAst) {
	var list *ZendAstList = ast.AsAstList()
	var i uint32
	var jmp_opnums *uint32 = nil
	if list.GetChildren() > 1 {
		jmp_opnums = SafeEmalloc(b.SizeOf("uint32_t"), list.GetChildren()-1, 0)
	}
	for i = 0; i < list.GetChildren(); i++ {
		var elem_ast *ZendAst = list.Children()[i]
		var cond_ast *ZendAst = elem_ast.Child(0)
		var stmt_ast *ZendAst = elem_ast.Child(1)
		if cond_ast != nil {
			var cond_node Znode
			var opnum_jmpz uint32
			compiler.CompileExpr(&cond_node, cond_ast)
			opnum_jmpz = ZendEmitCondJump(ZEND_JMPZ, &cond_node, 0)
			compiler.CompileStmt(stmt_ast)
			if i != list.GetChildren()-1 {
				jmp_opnums[i] = ZendEmitJump(0)
			}
			ZendUpdateJumpTargetToNext(opnum_jmpz)
		} else {

			/* "else" can only occur as last element. */

			b.Assert(i == list.GetChildren()-1)
			compiler.CompileStmt(stmt_ast)
		}
	}
	if list.GetChildren() > 1 {
		for i = 0; i < list.GetChildren()-1; i++ {
			ZendUpdateJumpTargetToNext(jmp_opnums[i])
		}
		Efree(jmp_opnums)
	}
}
func (compiler *Compiler) DetermineSwitchJumptableType(cases *ZendAstList) uint8 {
	var i uint32
	var common_type uint8 = types.IsUndef
	for i = 0; i < cases.GetChildren(); i++ {
		var case_ast *ZendAst = cases.Children()[i]
		var cond_ast **ZendAst = case_ast.Child(0)
		var cond_zv *types.Zval
		if case_ast.Child(0) == nil {

			/* Skip default clause */

			continue

			/* Skip default clause */

		}
		compiler.EvalConstExpr(cond_ast)
		if cond_ast.Kind() != ZEND_AST_ZVAL {

			/* Non-constant case */

			return types.IsUndef

			/* Non-constant case */

		}
		cond_zv = ZendAstGetZval(case_ast.Child(0))
		if !cond_zv.IsLong() && !cond_zv.IsString() {

			/* We only optimize switched on integers and strings */

			return types.IsUndef

			/* We only optimize switched on integers and strings */

		}
		if common_type == types.IsUndef {
			common_type = cond_zv.Type()
		} else if common_type != cond_zv.Type() {

			/* Non-uniform case types */

			return types.IsUndef

			/* Non-uniform case types */

		}
		if cond_zv.IsString() && operators.IsNumericString(cond_zv.String(), nil, nil, 0) != 0 {

			/* Numeric strings cannot be compared with a simple hash lookup */

			return types.IsUndef

			/* Numeric strings cannot be compared with a simple hash lookup */

		}
	}
	return common_type
}
func ShouldUseJumptable(cases *ZendAstList, jumptable_type uint8) bool {
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_NO_JUMPTABLES) != 0 {
		return 0
	}

	/* Thresholds are chosen based on when the average switch time for equidistributed
	 * input becomes smaller when using the jumptable optimization. */

	if jumptable_type == types.IsLong {
		return cases.GetChildren() >= 5
	} else {
		b.Assert(jumptable_type == types.IsString)
		return cases.GetChildren() >= 2
	}

	/* Thresholds are chosen based on when the average switch time for equidistributed
	 * input becomes smaller when using the jumptable optimization. */
}
func (compiler *Compiler) CompileSwitch(ast *ZendAst) {
	var expr_ast *ZendAst = ast.Child(0)
	var cases *ZendAstList = ast.Child(1).AsAstList()
	var i uint32
	var has_default_case bool = false
	var expr_node Znode
	var case_node Znode
	var opline *types.ZendOp
	var jmpnz_opnums *uint32
	var opnum_default_jmp uint32
	var opnum_switch uint32 = uint32(-1)
	var jumptable_type uint8
	var jumptable *types.Array = nil
	compiler.CompileExpr(&expr_node, expr_ast)
	ZendBeginLoop(ZEND_FREE, &expr_node, 1)
	case_node.SetOpType(IS_TMP_VAR)
	case_node.GetOp().SetVar(GetTemporaryVariable())
	jumptable_type = compiler.DetermineSwitchJumptableType(cases)
	if jumptable_type != types.IsUndef && ShouldUseJumptable(cases, jumptable_type) != 0 {
		var jumptable_op Znode
		jumptable = types.NewArrayCap(cases.GetChildren())
		jumptable_op.SetOpType(IS_CONST)
		jumptable_op.GetConstant().SetArray(jumptable)
		opline = ZendEmitOp(nil, lang.Cond(jumptable_type == types.IsLong, ZEND_SWITCH_LONG, ZEND_SWITCH_STRING), &expr_node, &jumptable_op)
		if opline.GetOp1Type() == IS_CONST {
			//CT_CONSTANT(opline.GetOp1()).TryAddRefcount()
		}
		opnum_switch = opline - CG__().GetActiveOpArray().GetOpcodes()
	}
	jmpnz_opnums = SafeEmalloc(b.SizeOf("uint32_t"), cases.GetChildren(), 0)
	for i = 0; i < cases.GetChildren(); i++ {
		var case_ast *ZendAst = cases.Children()[i]
		var cond_ast *ZendAst = case_ast.Child(0)
		var cond_node Znode
		if cond_ast == nil {
			if has_default_case != 0 {
				compiler.setLinenoByAst(case_ast)
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Switch statements may only contain one default clause")
			}
			has_default_case = 1
			continue
		}
		compiler.CompileExpr(&cond_node, cond_ast)
		if expr_node.GetOpType() == IS_CONST && expr_node.GetConstant().IsFalse() {
			jmpnz_opnums[i] = ZendEmitCondJump(ZEND_JMPZ, &cond_node, 0)
		} else if expr_node.GetOpType() == IS_CONST && expr_node.GetConstant().IsTrue() {
			jmpnz_opnums[i] = ZendEmitCondJump(ZEND_JMPNZ, &cond_node, 0)
		} else {
			opline = ZendEmitOp(nil, lang.Cond((expr_node.GetOpType()&(IS_VAR|IS_TMP_VAR)) != 0, ZEND_CASE, ZEND_IS_EQUAL), &expr_node, &cond_node)
			opline.SetResultType(case_node.GetOpType())
			if case_node.GetOpType() == IS_CONST {
				opline.GetResult().SetConstant(ZendAddLiteral(case_node.GetConstant()))
			} else {
				opline.SetResult(case_node.GetOp())
			}
			if opline.GetOp1Type() == IS_CONST {
				//CT_CONSTANT(opline.GetOp1()).TryAddRefcount()
			}
			jmpnz_opnums[i] = ZendEmitCondJump(ZEND_JMPNZ, &case_node, 0)
		}
	}
	opnum_default_jmp = ZendEmitJump(0)
	for i = 0; i < cases.GetChildren(); i++ {
		var case_ast *ZendAst = cases.Children()[i]
		var cond_ast *ZendAst = case_ast.Child(0)
		var stmt_ast *ZendAst = case_ast.Child(1)
		if cond_ast != nil {
			ZendUpdateJumpTargetToNext(jmpnz_opnums[i])
			if jumptable != nil {
				var cond_zv *types.Zval = cond_ast.Val()
				var jmp_target types.Zval
				jmp_target.SetLong(GetNextOpNumber())
				b.Assert(cond_zv.IsType(jumptable_type))
				if cond_zv.IsLong() {
					jumptable.IndexAdd(cond_zv.Long(), &jmp_target)
				} else {
					b.Assert(cond_zv.IsString())
					jumptable.KeyAdd(cond_zv.String(), &jmp_target)
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
		compiler.CompileStmt(stmt_ast)
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
		// ZvalPtrDtorNogc(expr_node.GetConstant())
	}
	Efree(jmpnz_opnums)
}
func (compiler *Compiler) CompileTry(ast *ZendAst) {
	var try_ast *ZendAst = ast.Child(0)
	var catches *ZendAstList = ast.Child(1).AsAstList()
	var finally_ast *ZendAst = ast.Children()[2]
	var i uint32
	var j uint32
	var opline *types.ZendOp
	var try_catch_offset uint32
	var jmp_opnums *uint32 = SafeEmalloc(b.SizeOf("uint32_t"), catches.GetChildren(), 0)
	var orig_fast_call_var uint32 = CG__().GetContext().GetFastCallVar()
	var orig_try_catch_offset uint32 = CG__().GetContext().GetTryCatchOffset()
	if catches.GetChildren() == 0 && finally_ast == nil {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use try without catch or finally")
	}

	/* label: try { } must not be equal to try { label: } */
	var label = CG__().GetContext().LastLabel()
	if label != nil {
		if label.GetOplineNum() == GetNextOpNumber() {
			ZendEmitOp(nil, ZEND_NOP, nil, nil)
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
		CG__().LoopVarStackPush(&fast_call)
	}
	CG__().GetContext().SetTryCatchOffset(try_catch_offset)
	compiler.CompileStmt(try_ast)
	if catches.GetChildren() != 0 {
		jmp_opnums[0] = ZendEmitJump(0)
	}
	for i = 0; i < catches.GetChildren(); i++ {
		var catch_ast *ZendAst = catches.Children()[i]
		var classes *ZendAstList = catch_ast.Child(0).AsAstList()
		var var_ast *ZendAst = catch_ast.Child(1)
		var stmt_ast *ZendAst = catch_ast.Children()[2]
		var var_name *types.String = var_ast.Val().StringEx()
		var is_last_catch bool = i+1 == catches.GetChildren()
		var jmp_multicatch *uint32 = SafeEmalloc(b.SizeOf("uint32_t"), classes.GetChildren()-1, 0)
		var opnum_catch uint32 = uint32 - 1
		compiler.setLinenoByAst(catch_ast)
		for j = 0; j < classes.GetChildren(); j++ {
			var class_ast *ZendAst = classes.Children()[j]
			var is_last_class bool = j+1 == classes.GetChildren()
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
			if var_name.GetStr() == "this" {
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
		compiler.CompileStmt(stmt_ast)
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

		CG__().LoopVarStackPop()

		/* Push DISCARD_EXCEPTION on unwind stack */

		discard_exception.SetOpcode(ZEND_DISCARD_EXCEPTION)
		discard_exception.SetVarType(IS_TMP_VAR)
		discard_exception.SetVarNum(CG__().GetContext().GetFastCallVar())
		CG__().LoopVarStackPush(&discard_exception)
		compiler.setLinenoByAst(finally_ast)
		opline = ZendEmitOp(nil, ZEND_FAST_CALL, nil, nil)
		opline.GetOp1().SetNum(try_catch_offset)
		opline.SetResultType(IS_TMP_VAR)
		opline.GetResult().SetVar(CG__().GetContext().GetFastCallVar())
		ZendEmitOp(nil, ZEND_JMP, nil, nil)
		compiler.CompileStmt(finally_ast)
		CG__().GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetFinallyOp(opnum_jmp + 1)
		CG__().GetActiveOpArray().GetTryCatchArray()[try_catch_offset].SetFinallyEnd(GetNextOpNumber())
		opline = ZendEmitOp(nil, ZEND_FAST_RET, nil, nil)
		opline.SetOp1Type(IS_TMP_VAR)
		opline.GetOp1().SetVar(CG__().GetContext().GetFastCallVar())
		opline.GetOp2().SetNum(orig_try_catch_offset)
		ZendUpdateJumpTargetToNext(opnum_jmp)
		CG__().GetContext().SetFastCallVar(orig_fast_call_var)

		/* Pop DISCARD_EXCEPTION from unwind stack */
		CG__().LoopVarStackPop()
	}
	CG__().GetContext().SetTryCatchOffset(orig_try_catch_offset)
	Efree(jmp_opnums)
}
func ZendDeclareIsFirstStatement(ast *ZendAst) int {
	var i uint32 = 0
	var file_ast *ZendAstList = CG__().GetAst().AsAstList()

	/* Check to see if this declare is preceded only by declare statements */

	for i < file_ast.GetChildren() {
		if file_ast.Children()[i] == ast {
			return types.SUCCESS
		} else if file_ast.Children()[i] == nil {

			/* Empty statements are not allowed prior to a declare */

			return types.FAILURE

			/* Empty statements are not allowed prior to a declare */

		} else if file_ast.Children()[i].Kind() != ZEND_AST_DECLARE {

			/* declares can only be preceded by other declares */

			return types.FAILURE

			/* declares can only be preceded by other declares */

		}
		i++
	}
	return types.FAILURE
}
func (compiler *Compiler) CompileDeclare(ast *ZendAst) {
	var declares *ZendAstList = ast.Child(0).AsAstList()
	var stmt_ast *ZendAst = ast.Child(1)
	var i uint32
	for i = 0; i < declares.GetChildren(); i++ {
		var declare_ast *ZendAst = declares.Children()[i]
		var name_ast *ZendAst = declare_ast.Child(0)
		var value_ast *ZendAst = declare_ast.Child(1)
		var name *types.String = ZendAstGetStr(name_ast)
		if value_ast.Kind() != ZEND_AST_ZVAL {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "declare(%s) value must be a literal", name.GetVal())
		}
		if ascii.StrCaseEquals(name.GetStr(), "ticks") {
			// todo 触发不支持 ticks 的 warning
		} else if ascii.StrCaseEquals(name.GetStr(), "encoding") {
			if types.FAILURE == ZendDeclareIsFirstStatement(ast) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Encoding declaration pragma must be the very first statement in the script")
			}
		} else if ascii.StrCaseEquals(name.GetStr(), "strict_types") {
			var value_zv types.Zval
			if types.FAILURE == ZendDeclareIsFirstStatement(ast) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "strict_types declaration must be the very first statement in the script")
			}
			if ast.Child(1) != nil {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "strict_types declaration must not use block mode")
			}
			compiler.ConstExprToZval(&value_zv, value_ast)
			if !value_zv.IsLong() || value_zv.Long() != 0 && value_zv.Long() != 1 {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "strict_types declaration must have 0 or 1 as its value")
			}
			if value_zv.Long() == 1 {
				CG__().GetActiveOpArray().SetIsStrictTypes(true)
			}
		} else {
			faults.Error(faults.E_COMPILE_WARNING, fmt.Sprintf("Unsupported declare '%s'", name.GetVal()))
		}
	}
	if stmt_ast != nil {
		compiler.CompileStmt(stmt_ast)
	}
}
func (compiler *Compiler) CompileStmtList(ast *ZendAst) {
	var list *ZendAstList = ast.AsAstList()
	var i uint32
	for i = 0; i < list.GetChildren(); i++ {
		compiler.CompileStmt(list.Children()[i])
	}
}
func (compiler *Compiler) CompileTypename(ast *ZendAst, force_allow_null bool) types.TypeHint {
	var allow_null bool = force_allow_null
	var orig_ast_attr ZendAstAttr = ast.Attr()
	var type_ types.TypeHint
	if (ast.Attr() & ZEND_TYPE_NULLABLE) != 0 {
		allow_null = true
		ast.SetAttr(ast.Attr() &^ ZEND_TYPE_NULLABLE)
	}
	if ast.Kind() == ZEND_AST_TYPE {
		return types.TypeHintCode(ast.Attr(), allow_null)
	} else {
		var class_name *types.String = ZendAstGetStr(ast)
		var type_code types.ZvalType = ZendLookupBuiltinTypeByName(class_name.GetStr())
		if type_code != 0 {
			if (ast.Attr() & ZEND_NAME_NOT_FQ) != ZEND_NAME_NOT_FQ {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Type declaration '%s' must be unqualified", operators.ZendStringTolower(class_name).GetVal())
			}
			type_ = types.TypeHintCode(type_code, allow_null)
		} else {
			var fetch_type uint32 = ZendGetClassFetchTypeAst(ast)
			if fetch_type == ZEND_FETCH_CLASS_DEFAULT {
				class_name = ZendResolveClassNameAst(ast)
				ZendAssertValidClassName(class_name.GetStr())
			} else {
				ZendEnsureValidClassFetchType(fetch_type)
				//class_name.AddRefcount()
			}
			type_ = types.TypeHintClassName(class_name, allow_null)
		}
	}
	ast.SetAttr(orig_ast_attr)
	return type_
}
func (compiler *Compiler) CompileParams(ast *ZendAst, return_type_ast *ZendAst) {
	var list *ZendAstList = ast.AsAstList()
	var i uint32
	var op_array *types.ZendOpArray = CG__().GetActiveOpArray()
	var arg_infos *ZendArgInfo
	if return_type_ast != nil {

		/* Use op_array->arg_info[-1] for return type */

		arg_infos = SafeEmalloc(b.SizeOf("zend_arg_info"), list.GetChildren()+1, 0)
		*arg_infos = MakeZendReturnArgInfo(
			compiler.CompileTypename(return_type_ast, 0),
			op_array.IsReturnReference(),
		)
		if arg_infos.GetType().Code() == types.IsVoid && arg_infos.GetType().AllowNull() {
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
		var param_ast *ZendAst = list.Children()[i]
		var type_ast *ZendAst = param_ast.Child(0)
		var var_ast *ZendAst = param_ast.Child(1)
		var default_ast *ZendAst = param_ast.Children()[2]
		var name *types.String = var_ast.Val().StringEx()
		var is_ref bool = (param_ast.Attr() & ZEND_PARAM_REF) != 0
		var is_variadic bool = (param_ast.Attr() & ZEND_PARAM_VARIADIC) != 0
		var var_node Znode
		var default_node Znode
		var opcode uint8
		var opline *types.ZendOp
		var arg_info *ZendArgInfo
		if ZendIsAutoGlobal(name.GetStr()) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot re-assign auto-global variable %s", name.GetVal())
		}
		var_node.SetOpType(IS_CV)
		var_node.GetOp().SetVar(LookupCv(name.GetStr()))
		if EX_VAR_TO_NUM(var_node.GetOp().GetVar()) != i {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Redefinition of parameter $%s", name.GetVal())
		} else if name.GetStr() == "this" {
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
			compiler.ConstExprToZval(default_node.GetConstant(), default_ast)
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
			types.TypeHintCode(0, 1),
			is_ref,
			is_variadic,
		)
		if type_ast != nil {
			var has_null_default bool = default_ast != nil && (default_node.GetConstant().IsNull() || default_node.GetConstant().IsConstantAst() && types.Z_ASTVAL(default_node.GetConstant()).Kind() == ZEND_AST_CONSTANT && ascii.StrCaseEquals(ZendAstGetConstantName(types.Z_ASTVAL(default_node.GetConstant())), "NULL"))
			op_array.SetIsHasTypeHints(true)
			arg_info.SetType(compiler.CompileTypename(type_ast, has_null_default))
			if arg_info.GetType().Code() == types.IsVoid {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "void cannot be used as a parameter type")
			}
			if type_ast.Kind() == ZEND_AST_TYPE {
				if arg_info.GetType().Code() == types.IsArray {
					if default_ast != nil && has_null_default == 0 && default_node.GetConstant().Type() != types.IsArray && default_node.GetConstant().Type() != types.IsConstantAst {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters with array type can only be an array or NULL")
					}
				} else if arg_info.GetType().Code() == types.IsCallable && default_ast != nil {
					if has_null_default == 0 && default_node.GetConstant().Type() != types.IsConstantAst {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters with callable type can only be NULL")
					}
				}
			} else {
				if default_ast != nil && has_null_default == 0 && default_node.GetConstant().Type() != types.IsConstantAst {
					if arg_info.GetType().IsClass() {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters with a class type can only be NULL")
					} else {
						switch arg_info.GetType().Code() {
						case types.IsDouble:
							if default_node.GetConstant().Type() != types.IsDouble && default_node.GetConstant().Type() != types.IsLong {
								faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters with a float type can only be float, integer, or NULL")
							}
							operators.ConvertToDouble(default_node.GetConstant())
						case types.IsIterable:
							if default_node.GetConstant().Type() != types.IsArray {
								faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters with iterable type can only be an array or NULL")
							}
						case types.IsObject:
							faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters with an object type can only be NULL")
						default:
							if !(types.ZEND_SAME_FAKE_TYPE(arg_info.GetType().Code(), default_node.GetConstant().Type())) {
								faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Default value for parameters with a %s type can only be %s or NULL", types.ZendGetTypeByConst(arg_info.GetType().Code()), types.ZendGetTypeByConst(arg_info.GetType().Code()))
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
