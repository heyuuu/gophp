// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZendCompileAssignRef(result *Znode, ast *ZendAst) {
	var target_ast *ZendAst = ast.GetChild()[0]
	var source_ast *ZendAst = ast.GetChild()[1]
	var target_node Znode
	var source_node Znode
	var opline *ZendOp
	var offset uint32
	var flags uint32
	if IsThisFetch(target_ast) != 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot re-assign $this")
	}
	ZendEnsureWritableVariable(target_ast)
	offset = ZendDelayedCompileBegin()
	ZendDelayedCompileVar(&target_node, target_ast, BP_VAR_W, 1)
	ZendCompileVar(&source_node, source_ast, BP_VAR_W, 1)
	if (target_ast.GetKind() != ZEND_AST_VAR || target_ast.GetChild()[0].GetKind() != ZEND_AST_ZVAL) && source_node.GetOpType() != IS_CV {

		/* Both LHS and RHS expressions may modify the same data structure,
		 * and the modification during RHS evaluation may dangle the pointer
		 * to the result of the LHS evaluation.
		 * Use MAKE_REF instruction to replace direct pointer with REFERENCE.
		 * See: Bug #71539
		 */

		ZendEmitOp(&source_node, ZEND_MAKE_REF, &source_node, nil)

		/* Both LHS and RHS expressions may modify the same data structure,
		 * and the modification during RHS evaluation may dangle the pointer
		 * to the result of the LHS evaluation.
		 * Use MAKE_REF instruction to replace direct pointer with REFERENCE.
		 * See: Bug #71539
		 */

	}
	opline = ZendDelayedCompileEnd(offset)
	if source_node.GetOpType() != IS_VAR && ZendIsCall(source_ast) != 0 {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use result of built-in function in write context")
	}
	if ZendIsCall(source_ast) != 0 {
		flags = ZEND_RETURNS_FUNCTION
	} else {
		flags = 0
	}
	if opline != nil && opline.GetOpcode() == ZEND_FETCH_OBJ_W {
		opline.SetOpcode(ZEND_ASSIGN_OBJ_REF)
		opline.SetExtendedValue(opline.GetExtendedValue() &^ ZEND_FETCH_REF)
		opline.SetExtendedValue(opline.GetExtendedValue() | flags)
		ZendEmitOpData(&source_node)
		*result = target_node
	} else if opline != nil && opline.GetOpcode() == ZEND_FETCH_STATIC_PROP_W {
		opline.SetOpcode(ZEND_ASSIGN_STATIC_PROP_REF)
		opline.SetExtendedValue(opline.GetExtendedValue() &^ ZEND_FETCH_REF)
		opline.SetExtendedValue(opline.GetExtendedValue() | flags)
		ZendEmitOpData(&source_node)
		*result = target_node
	} else {
		opline = ZendEmitOp(result, ZEND_ASSIGN_REF, &target_node, &source_node)
		opline.SetExtendedValue(flags)
	}
}
func ZendEmitAssignRefZnode(var_ast *ZendAst, value_node *Znode) {
	var dummy_node Znode
	var assign_ast *ZendAst = ZendAstCreate(ZEND_AST_ASSIGN_REF, var_ast, ZendAstCreateZnode(value_node))
	ZendCompileAssignRef(&dummy_node, assign_ast)
	ZendDoFree(&dummy_node)
}
func ZendCompileCompoundAssign(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.GetChild()[0]
	var expr_ast *ZendAst = ast.GetChild()[1]
	var opcode uint32 = ast.GetAttr()
	var var_node Znode
	var expr_node Znode
	var opline *ZendOp
	var offset uint32
	var cache_slot uint32
	ZendEnsureWritableVariable(var_ast)
	switch var_ast.GetKind() {
	case ZEND_AST_VAR:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileVar(&var_node, var_ast, BP_VAR_RW, 0)
		ZendCompileExpr(&expr_node, expr_ast)
		ZendDelayedCompileEnd(offset)
		opline = ZendEmitOp(result, ZEND_ASSIGN_OP, &var_node, &expr_node)
		opline.SetExtendedValue(opcode)
		return
	case ZEND_AST_STATIC_PROP:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileVar(result, var_ast, BP_VAR_RW, 0)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		cache_slot = opline.GetExtendedValue()
		opline.SetOpcode(ZEND_ASSIGN_STATIC_PROP_OP)
		opline.SetExtendedValue(opcode)
		opline = ZendEmitOpData(&expr_node)
		opline.SetExtendedValue(cache_slot)
		return
	case ZEND_AST_DIM:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileDim(result, var_ast, BP_VAR_RW)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		opline.SetOpcode(ZEND_ASSIGN_DIM_OP)
		opline.SetExtendedValue(opcode)
		ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_PROP:
		offset = ZendDelayedCompileBegin()
		ZendDelayedCompileProp(result, var_ast, BP_VAR_RW)
		ZendCompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		cache_slot = opline.GetExtendedValue()
		opline.SetOpcode(ZEND_ASSIGN_OBJ_OP)
		opline.SetExtendedValue(opcode)
		opline = ZendEmitOpData(&expr_node)
		opline.SetExtendedValue(cache_slot)
		return
	default:

	}
}
func ZendCompileArgs(ast *ZendAst, fbc *ZendFunction) uint32 {
	var args *ZendAstList = ZendAstGetList(ast)
	var i uint32
	var uses_arg_unpack types.ZendBool = 0
	var arg_count uint32 = 0
	for i = 0; i < args.GetChildren(); i++ {
		var arg *ZendAst = args.GetChild()[i]
		var arg_num uint32 = i + 1
		var arg_node Znode
		var opline *ZendOp
		var opcode types.ZendUchar
		if arg.GetKind() == ZEND_AST_UNPACK {
			uses_arg_unpack = 1
			fbc = nil
			ZendCompileExpr(&arg_node, arg.GetChild()[0])
			opline = ZendEmitOp(nil, ZEND_SEND_UNPACK, &arg_node, nil)
			opline.GetOp2().SetNum(arg_count)
			opline.GetResult().SetVar(uint32(types.ZendIntptrT(nil.Arg(arg_count))))
			continue
		}
		if uses_arg_unpack != 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use positional argument after argument unpacking")
		}
		arg_count++
		if ZendIsVariableOrCall(arg) != 0 {
			if ZendIsCall(arg) != 0 {
				ZendCompileVar(&arg_node, arg, BP_VAR_R, 0)
				if (arg_node.GetOpType() & (IS_CONST | IS_TMP_VAR)) != 0 {

					/* Function call was converted into builtin instruction */

					if fbc == nil || ARG_MUST_BE_SENT_BY_REF(fbc, arg_num) != 0 {
						opcode = ZEND_SEND_VAL_EX
					} else {
						opcode = ZEND_SEND_VAL
					}

					/* Function call was converted into builtin instruction */

				} else {
					if fbc != nil {
						if ARG_MUST_BE_SENT_BY_REF(fbc, arg_num) != 0 {
							opcode = ZEND_SEND_VAR_NO_REF
						} else if ARG_MAY_BE_SENT_BY_REF(fbc, arg_num) != 0 {
							opcode = ZEND_SEND_VAL
						} else {
							opcode = ZEND_SEND_VAR
						}
					} else {
						opcode = ZEND_SEND_VAR_NO_REF_EX
					}
				}
			} else if fbc != nil {
				if ARG_SHOULD_BE_SENT_BY_REF(fbc, arg_num) != 0 {
					ZendCompileVar(&arg_node, arg, BP_VAR_W, 1)
					opcode = ZEND_SEND_REF
				} else {
					ZendCompileVar(&arg_node, arg, BP_VAR_R, 0)
					if arg_node.GetOpType() == IS_TMP_VAR {
						opcode = ZEND_SEND_VAL
					} else {
						opcode = ZEND_SEND_VAR
					}
				}
			} else {
				for {
					if arg.GetKind() == ZEND_AST_VAR {
						CG__().SetZendLineno(ZendAstGetLineno(ast))
						if IsThisFetch(arg) != 0 {
							ZendEmitOp(&arg_node, ZEND_FETCH_THIS, nil, nil)
							opcode = ZEND_SEND_VAR_EX
							CG__().GetActiveOpArray().SetIsUsesThis(true)
							break
						} else if ZendTryCompileCv(&arg_node, arg) == types.SUCCESS {
							opcode = ZEND_SEND_VAR_EX
							break
						}
					}
					opline = ZendEmitOp(nil, ZEND_CHECK_FUNC_ARG, nil, nil)
					opline.GetOp2().SetNum(arg_num)
					ZendCompileVar(&arg_node, arg, BP_VAR_FUNC_ARG, 1)
					opcode = ZEND_SEND_FUNC_ARG
					break
				}
			}
		} else {
			ZendCompileExpr(&arg_node, arg)
			if arg_node.GetOpType() == IS_VAR {

				/* pass ++$a or something similar */

				if fbc != nil {
					if ARG_MUST_BE_SENT_BY_REF(fbc, arg_num) != 0 {
						opcode = ZEND_SEND_VAR_NO_REF
					} else if ARG_MAY_BE_SENT_BY_REF(fbc, arg_num) != 0 {
						opcode = ZEND_SEND_VAL
					} else {
						opcode = ZEND_SEND_VAR
					}
				} else {
					opcode = ZEND_SEND_VAR_NO_REF_EX
				}

				/* pass ++$a or something similar */

			} else if arg_node.GetOpType() == IS_CV {
				if fbc != nil {
					if ARG_SHOULD_BE_SENT_BY_REF(fbc, arg_num) != 0 {
						opcode = ZEND_SEND_REF
					} else {
						opcode = ZEND_SEND_VAR
					}
				} else {
					opcode = ZEND_SEND_VAR_EX
				}
			} else {
				if fbc != nil {
					opcode = ZEND_SEND_VAL
					if ARG_MUST_BE_SENT_BY_REF(fbc, arg_num) != 0 {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Only variables can be passed by reference")
					}
				} else {
					opcode = ZEND_SEND_VAL_EX
				}
			}
		}
		opline = ZendEmitOp(nil, opcode, &arg_node, nil)
		opline.GetOp2().SetOplineNum(arg_num)
		opline.GetResult().SetVar(uint32(types.ZendIntptrT(nil.Arg(arg_num))))
	}
	return arg_count
}
func ZendGetCallOp(init_op *ZendOp, fbc *ZendFunction) types.ZendUchar {
	if fbc != nil {
		if fbc.GetType() == ZEND_INTERNAL_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS) == 0 {
			if init_op.GetOpcode() == ZEND_INIT_FCALL && ZendExecuteInternal == nil {
				if !fbc.HasFnFlags(ZEND_ACC_ABSTRACT | ZEND_ACC_DEPRECATED | ZEND_ACC_HAS_TYPE_HINTS | ZEND_ACC_RETURN_REFERENCE) {
					return ZEND_DO_ICALL
				} else {
					return ZEND_DO_FCALL_BY_NAME
				}
			}
		} else if (CG__().GetCompilerOptions() & ZEND_COMPILE_IGNORE_USER_FUNCTIONS) == 0 {
			if ZendExecuteEx == ExecuteEx && !fbc.IsAbstract() {
				return ZEND_DO_UCALL
			}
		}
	} else if ZendExecuteEx == ExecuteEx && ZendExecuteInternal == nil && (init_op.GetOpcode() == ZEND_INIT_FCALL_BY_NAME || init_op.GetOpcode() == ZEND_INIT_NS_FCALL_BY_NAME) {
		return ZEND_DO_FCALL_BY_NAME
	}
	return ZEND_DO_FCALL
}
func ZendCompileCallCommon(result *Znode, args_ast *ZendAst, fbc *ZendFunction) {
	var opline *ZendOp
	var opnum_init uint32 = GetNextOpNumber() - 1
	var arg_count uint32
	arg_count = ZendCompileArgs(args_ast, fbc)
	ZendDoExtendedFcallBegin()
	opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_init]
	opline.SetExtendedValue(arg_count)
	if opline.GetOpcode() == ZEND_INIT_FCALL {
		opline.GetOp1().SetNum(ZendVmCalcUsedStack(arg_count, fbc))
	}
	opline = ZendEmitOp(result, ZendGetCallOp(opline, fbc), nil, nil)
	ZendDoExtendedFcallEnd()
}
func ZendCompileFunctionName(name_node *Znode, name_ast *ZendAst) types.ZendBool {
	var orig_name *types.String = ZendAstGetStr(name_ast)
	var is_fully_qualified types.ZendBool
	name_node.SetOpType(IS_CONST)
	name_node.GetConstant().SetString(ZendResolveFunctionName(orig_name, name_ast.GetAttr(), &is_fully_qualified))
	return is_fully_qualified == 0 && FC__().GetCurrentNamespace() != nil
}
func ZendCompileNsCall(result *Znode, name_node *Znode, args_ast *ZendAst) {
	var opline *ZendOp = GetNextOp()
	opline.SetOpcode(ZEND_INIT_NS_FCALL_BY_NAME)
	opline.SetOp2Type(IS_CONST)
	opline.GetOp2().SetConstant(ZendAddNsFuncNameLiteral(name_node.GetConstant().GetStr()))
	opline.GetResult().SetNum(ZendAllocCacheSlot())
	ZendCompileCallCommon(result, args_ast, nil)
}
func ZendCompileDynamicCall(result *Znode, name_node *Znode, args_ast *ZendAst) {
	if name_node.GetOpType() == IS_CONST && name_node.GetConstant().IsString() {
		var colon *byte
		var str *types.String = name_node.GetConstant().GetStr()
		if b.Assign(&colon, ZendMemrchr(str.GetVal(), ':', str.GetLen())) != nil && colon > str.GetVal() && (*(colon - 1)) == ':' {
			var class *types.String = types.NewString(b.CastStr(str.GetVal(), colon-str.GetVal()-1))
			var method *types.String = types.NewString(b.CastStr(colon+1, str.GetLen()-(colon-str.GetVal())-1))
			var opline *ZendOp = GetNextOp()
			opline.SetOpcode(ZEND_INIT_STATIC_METHOD_CALL)
			opline.SetOp1Type(IS_CONST)
			opline.GetOp1().SetConstant(ZendAddClassNameLiteral(class))
			opline.SetOp2Type(IS_CONST)
			opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(method))

			/* 2 slots, for class and method */

			opline.GetResult().SetNum(ZendAllocCacheSlots(2))
			ZvalPtrDtor(name_node.GetConstant())
		} else {
			var opline *ZendOp = GetNextOp()
			opline.SetOpcode(ZEND_INIT_FCALL_BY_NAME)
			opline.SetOp2Type(IS_CONST)
			opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(str))
			opline.GetResult().SetNum(ZendAllocCacheSlot())
		}
	} else {
		ZendEmitOp(nil, ZEND_INIT_DYNAMIC_CALL, nil, name_node)
	}
	ZendCompileCallCommon(result, args_ast, nil)
}
func ZendArgsContainUnpack(args *ZendAstList) types.ZendBool {
	var i uint32
	for i = 0; i < args.GetChildren(); i++ {
		if args.GetChild()[i].GetKind() == ZEND_AST_UNPACK {
			return 1
		}
	}
	return 0
}
func ZendCompileFuncStrlen(result *Znode, args *ZendAstList) int {
	var arg_node Znode
	if (CG__().GetCompilerOptions()&ZEND_COMPILE_NO_BUILTIN_STRLEN) != 0 || args.GetChildren() != 1 {
		return types.FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	if arg_node.GetOpType() == IS_CONST && arg_node.GetConstant().IsString() {
		result.SetOpType(IS_CONST)
		result.GetConstant().SetLong(arg_node.GetConstant().GetStr().GetLen())
	} else {
		ZendEmitOpTmp(result, ZEND_STRLEN, &arg_node, nil)
	}
	return types.SUCCESS
}
func ZendCompileFuncTypecheck(result *Znode, args *ZendAstList, type_ uint32) int {
	var arg_node Znode
	var opline *ZendOp
	if args.GetChildren() != 1 {
		return types.FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	opline = ZendEmitOpTmp(result, ZEND_TYPE_CHECK, &arg_node, nil)
	if type_ != types.IS_BOOL {
		opline.SetExtendedValue(1 << type_)
	} else {
		opline.SetExtendedValue(1<<types.IS_FALSE | 1<<types.IS_TRUE)
	}
	return types.SUCCESS
}
func ZendCompileFuncCast(result *Znode, args *ZendAstList, type_ uint32) int {
	var arg_node Znode
	var opline *ZendOp
	if args.GetChildren() != 1 {
		return types.FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	opline = ZendEmitOpTmp(result, ZEND_CAST, &arg_node, nil)
	opline.SetExtendedValue(type_)
	return types.SUCCESS
}
func ZendCompileFuncDefined(result *Znode, args *ZendAstList) int {
	var name *types.String
	var opline *ZendOp
	if args.GetChildren() != 1 || args.GetChild()[0].GetKind() != ZEND_AST_ZVAL {
		return types.FAILURE
	}
	name = ZvalGetString(ZendAstGetZval(args.GetChild()[0]))
	if ZendMemrchr(name.GetVal(), '\\', name.GetLen()) || ZendMemrchr(name.GetVal(), ':', name.GetLen()) {
		types.ZendStringReleaseEx(name, 0)
		return types.FAILURE
	}
	if ZendTryCtEvalConst(result.GetConstant(), name, 0) != 0 {
		types.ZendStringReleaseEx(name, 0)
		ZvalPtrDtor(result.GetConstant())
		result.GetConstant().SetTrue()
		result.SetOpType(IS_CONST)
		return types.SUCCESS
	}
	opline = ZendEmitOpTmp(result, ZEND_DEFINED, nil, nil)
	opline.SetOp1Type(IS_CONST)
	LITERAL_STR(opline.GetOp1(), name)
	opline.SetExtendedValue(ZendAllocCacheSlot())

	/* Lowercase constant name in a separate literal */

	var c types.Zval
	var lcname *types.String = ZendStringTolower(name)
	c.SetString(lcname)
	ZendAddLiteral(&c)
	return types.SUCCESS
}
func ZendCompileFuncChr(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 1 && args.GetChild()[0].GetKind() == ZEND_AST_ZVAL && ZendAstGetZval(args.GetChild()[0]).IsLong() {
		var c ZendLong = ZendAstGetZval(args.GetChild()[0]).GetLval() & 0xff
		result.SetOpType(IS_CONST)
		result.GetConstant().SetInternedString(types.ZSTR_CHAR(c))
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func ZendCompileFuncOrd(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 1 && args.GetChild()[0].GetKind() == ZEND_AST_ZVAL && ZendAstGetZval(args.GetChild()[0]).IsString() {
		result.SetOpType(IS_CONST)
		result.GetConstant().SetLong(uint8(ZendAstGetZval(args.GetChild()[0]).GetStr().GetVal()[0]))
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func FbcIsFinalized(fbc *ZendFunction) types.ZendBool {
	return !(ZEND_USER_CODE(fbc.GetType())) || fbc.IsDonePassTwo()
}
func ZendTryCompileCtBoundInitUserFunc(name_ast *ZendAst, num_args uint32) int {
	var name *types.String
	var lcname *types.String
	var fbc *ZendFunction
	var opline *ZendOp
	if name_ast.GetKind() != ZEND_AST_ZVAL || ZendAstGetZval(name_ast).GetType() != types.IS_STRING {
		return types.FAILURE
	}
	name = ZendAstGetStr(name_ast)
	lcname = ZendStringTolower(name)
	fbc = types.ZendHashFindPtr(CG__().GetFunctionTable(), lcname.GetStr())
	if fbc == nil || FbcIsFinalized(fbc) == 0 || fbc.GetType() == ZEND_INTERNAL_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_USER_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_OTHER_FILES) != 0 && fbc.GetOpArray().GetFilename() != CG__().GetActiveOpArray().GetFilename() {
		types.ZendStringReleaseEx(lcname, 0)
		return types.FAILURE
	}
	opline = ZendEmitOp(nil, ZEND_INIT_FCALL, nil, nil)
	opline.SetExtendedValue(num_args)
	opline.GetOp1().SetNum(ZendVmCalcUsedStack(num_args, fbc))
	opline.SetOp2Type(IS_CONST)
	LITERAL_STR(opline.GetOp2(), lcname)
	opline.GetResult().SetNum(ZendAllocCacheSlot())
	return types.SUCCESS
}
func ZendCompileInitUserFunc(name_ast *ZendAst, num_args uint32, orig_func_name *types.String) {
	var opline *ZendOp
	var name_node Znode
	if ZendTryCompileCtBoundInitUserFunc(name_ast, num_args) == types.SUCCESS {
		return
	}
	ZendCompileExpr(&name_node, name_ast)
	opline = ZendEmitOp(nil, ZEND_INIT_USER_CALL, nil, &name_node)
	opline.SetOp1Type(IS_CONST)
	LITERAL_STR(opline.GetOp1(), orig_func_name.Copy())
	opline.SetExtendedValue(num_args)
}
func ZendCompileFuncCufa(result *Znode, args *ZendAstList, lcname *types.String) int {
	var arg_node Znode
	if args.GetChildren() != 2 {
		return types.FAILURE
	}
	ZendCompileInitUserFunc(args.GetChild()[0], 0, lcname)
	if args.GetChild()[1].GetKind() == ZEND_AST_CALL && args.GetChild()[1].GetChild()[0].GetKind() == ZEND_AST_ZVAL && ZendAstGetZval(args.GetChild()[1].GetChild()[0]).IsString() && args.GetChild()[1].GetChild()[1].GetKind() == ZEND_AST_ARG_LIST {
		var orig_name *types.String = ZendAstGetStr(args.GetChild()[1].GetChild()[0])
		var list *ZendAstList = ZendAstGetList(args.GetChild()[1].GetChild()[1])
		var is_fully_qualified types.ZendBool
		var name *types.String = ZendResolveFunctionName(orig_name, args.GetChild()[1].GetChild()[0].GetAttr(), &is_fully_qualified)
		if types.ZendStringEqualsLiteralCi(name, "array_slice") && list.GetChildren() == 3 && list.GetChild()[1].GetKind() == ZEND_AST_ZVAL {
			var zv *types.Zval = ZendAstGetZval(list.GetChild()[1])
			if zv.IsLong() && zv.GetLval() >= 0 && zv.GetLval() <= 0x7fffffff {
				var opline *ZendOp
				var len_node Znode
				ZendCompileExpr(&arg_node, list.GetChild()[0])
				ZendCompileExpr(&len_node, list.GetChild()[2])
				opline = ZendEmitOp(nil, ZEND_SEND_ARRAY, &arg_node, &len_node)
				opline.SetExtendedValue(zv.GetLval())
				ZendEmitOp(result, ZEND_DO_FCALL, nil, nil)
				types.ZendStringReleaseEx(name, 0)
				return types.SUCCESS
			}
		}
		types.ZendStringReleaseEx(name, 0)
	}
	ZendCompileExpr(&arg_node, args.GetChild()[1])
	ZendEmitOp(nil, ZEND_SEND_ARRAY, &arg_node, nil)
	ZendEmitOp(result, ZEND_DO_FCALL, nil, nil)
	return types.SUCCESS
}
func ZendCompileFuncCuf(result *Znode, args *ZendAstList, lcname *types.String) int {
	var i uint32
	if args.GetChildren() < 1 {
		return types.FAILURE
	}
	ZendCompileInitUserFunc(args.GetChild()[0], args.GetChildren()-1, lcname)
	for i = 1; i < args.GetChildren(); i++ {
		var arg_ast *ZendAst = args.GetChild()[i]
		var arg_node Znode
		var opline *ZendOp
		ZendCompileExpr(&arg_node, arg_ast)
		opline = ZendEmitOp(nil, ZEND_SEND_USER, &arg_node, nil)
		opline.GetOp2().SetNum(i)
		opline.GetResult().SetVar(uint32(types.ZendIntptrT(nil.Arg(i))))
	}
	ZendEmitOp(result, ZEND_DO_FCALL, nil, nil)
	return types.SUCCESS
}
func ZendCompileAssert(result *Znode, args *ZendAstList, name *types.String, fbc *ZendFunction) {
	if EG__().GetAssertions() >= 0 {
		var name_node Znode
		var opline *ZendOp
		var check_op_number uint32 = GetNextOpNumber()
		ZendEmitOp(nil, ZEND_ASSERT_CHECK, nil, nil)
		if fbc != nil && FbcIsFinalized(fbc) != 0 {
			name_node.SetOpType(IS_CONST)
			name_node.GetConstant().SetStringCopy(name)
			opline = ZendEmitOp(nil, ZEND_INIT_FCALL, nil, &name_node)
		} else {
			opline = ZendEmitOp(nil, ZEND_INIT_NS_FCALL_BY_NAME, nil, nil)
			opline.SetOp2Type(IS_CONST)
			opline.GetOp2().SetConstant(ZendAddNsFuncNameLiteral(name))
		}
		opline.GetResult().SetNum(ZendAllocCacheSlot())
		if args.GetChildren() == 1 && (args.GetChild()[0].GetKind() != ZEND_AST_ZVAL || ZendAstGetZval(args.GetChild()[0]).GetType() != types.IS_STRING) {

			/* add "assert(condition) as assertion message */

			ZendAstListAdd((*ZendAst)(args), ZendAstCreateZvalFromStr(ZendAstExport("assert(", args.GetChild()[0], ")")))

			/* add "assert(condition) as assertion message */

		}
		ZendCompileCallCommon(result, (*ZendAst)(args), fbc)
		opline = CG__().GetActiveOpArray().GetOpcodes()[check_op_number]
		opline.GetOp2().SetOplineNum(GetNextOpNumber())
		opline.SetResultType(result.GetOpType())
		if result.GetOpType() == IS_CONST {
			opline.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
		} else {
			opline.SetResult(result.GetOp())
		}
	} else {
		if fbc == nil {
			types.ZendStringReleaseEx(name, 0)
		}
		result.SetOpType(IS_CONST)
		result.GetConstant().SetTrue()
	}
}
func ZendCompileFuncInArray(result *Znode, args *ZendAstList) int {
	var strict types.ZendBool = 0
	var array Znode
	var needly Znode
	var opline *ZendOp
	if args.GetChildren() == 3 {
		if args.GetChild()[2].GetKind() == ZEND_AST_ZVAL {
			strict = ZendIsTrue(ZendAstGetZval(args.GetChild()[2]))
		} else if args.GetChild()[2].GetKind() == ZEND_AST_CONST {
			var value types.Zval
			var name_ast *ZendAst = args.GetChild()[2].GetChild()[0]
			var is_fully_qualified types.ZendBool
			var resolved_name *types.String = ZendResolveConstName(ZendAstGetStr(name_ast), name_ast.GetAttr(), &is_fully_qualified)
			if ZendTryCtEvalConst(&value, resolved_name, is_fully_qualified) == 0 {
				types.ZendStringReleaseEx(resolved_name, 0)
				return types.FAILURE
			}
			types.ZendStringReleaseEx(resolved_name, 0)
			strict = ZendIsTrue(&value)
			ZvalPtrDtor(&value)
		} else {
			return types.FAILURE
		}
	} else if args.GetChildren() != 2 {
		return types.FAILURE
	}
	if args.GetChild()[1].GetKind() != ZEND_AST_ARRAY || ZendTryCtEvalArray(array.GetConstant(), args.GetChild()[1]) == 0 {
		return types.FAILURE
	}
	if types.Z_ARRVAL(array.GetConstant()).GetNNumOfElements() > 0 {
		var ok types.ZendBool = 1
		var val *types.Zval
		var tmp types.Zval
		var src *types.Array = array.GetConstant().GetArr()
		var dst *types.Array = types.NewZendArray(src.GetNNumOfElements())
		tmp.SetTrue()
		if strict != 0 {
			var __ht *types.Array = src
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()

				val = _z
				if val.IsString() {
					dst.KeyAdd(val.GetStr().GetStr(), &tmp)
				} else if val.IsLong() {
					dst.IndexAddH(val.GetLval(), &tmp)
				} else {
					dst.DestroyEx()
					ok = 0
					break
				}
			}
		} else {
			var __ht *types.Array = src
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()

				val = _z
				if val.GetType() != types.IS_STRING || IsNumericString(val.GetStr().GetStr(), nil, nil, 0) != 0 {
					dst.DestroyEx()
					ok = 0
					break
				}
				dst.KeyAdd(val.GetStr().GetStr(), &tmp)
			}
		}
		src.DestroyEx()
		if ok == 0 {
			return types.FAILURE
		}
		array.GetConstant().GetArr() = dst
	}
	array.SetOpType(IS_CONST)
	ZendCompileExpr(&needly, args.GetChild()[0])
	opline = ZendEmitOpTmp(result, ZEND_IN_ARRAY, &needly, &array)
	opline.SetExtendedValue(strict)
	return types.SUCCESS
}
func ZendCompileFuncCount(result *Znode, args *ZendAstList, lcname *types.String) int {
	var arg_node Znode
	var opline *ZendOp
	if args.GetChildren() != 1 {
		return types.FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	opline = ZendEmitOpTmp(result, ZEND_COUNT, &arg_node, nil)
	opline.SetExtendedValue(types.ZendStringEqualsLiteral(lcname, "sizeof"))
	return types.SUCCESS
}
func ZendCompileFuncGetClass(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 0 {
		ZendEmitOpTmp(result, ZEND_GET_CLASS, nil, nil)
	} else {
		var arg_node Znode
		if args.GetChildren() != 1 {
			return types.FAILURE
		}
		ZendCompileExpr(&arg_node, args.GetChild()[0])
		ZendEmitOpTmp(result, ZEND_GET_CLASS, &arg_node, nil)
	}
	return types.SUCCESS
}
func ZendCompileFuncGetCalledClass(result *Znode, args *ZendAstList) int {
	if args.GetChildren() != 0 {
		return types.FAILURE
	}
	ZendEmitOpTmp(result, ZEND_GET_CALLED_CLASS, nil, nil)
	return types.SUCCESS
}
func ZendCompileFuncGettype(result *Znode, args *ZendAstList) int {
	var arg_node Znode
	if args.GetChildren() != 1 {
		return types.FAILURE
	}
	ZendCompileExpr(&arg_node, args.GetChild()[0])
	ZendEmitOpTmp(result, ZEND_GET_TYPE, &arg_node, nil)
	return types.SUCCESS
}
func ZendCompileFuncNumArgs(result *Znode, args *ZendAstList) int {
	if CG__().GetActiveOpArray().GetFunctionName() != nil && args.GetChildren() == 0 {
		ZendEmitOpTmp(result, ZEND_FUNC_NUM_ARGS, nil, nil)
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
