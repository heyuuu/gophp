package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func (compiler *Compiler) CompileAssignRef(result *Znode, ast *ZendAst) {
	var target_ast *ZendAst = ast.Child(0)
	var source_ast *ZendAst = ast.Child(1)
	var target_node Znode
	var source_node Znode
	var opline *types.ZendOp
	var offset uint32
	var flags uint32
	if IsThisFetch(target_ast) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot re-assign $this")
	}
	ZendEnsureWritableVariable(target_ast)
	offset = ZendDelayedCompileBegin()
	compiler.DelayedCompileVar(&target_node, target_ast, BP_VAR_W, 1)
	compiler.CompileVar(&source_node, source_ast, BP_VAR_W, 1)
	if (target_ast.Kind() != ZEND_AST_VAR || target_ast.Child(0).Kind() != ZEND_AST_ZVAL) && source_node.GetOpType() != IS_CV {

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
func (compiler *Compiler) EmitAssignRefZnode(var_ast *ZendAst, value_node *Znode) {
	var dummyNode Znode
	var assignAst *ZendAst = AstCreate(ZEND_AST_ASSIGN_REF, var_ast, ZendAstCreateZnode(value_node))
	compiler.CompileAssignRef(&dummyNode, assignAst)
	ZendDoFree(&dummyNode)
}
func (compiler *Compiler) CompileCompoundAssign(result *Znode, ast *ZendAst) {
	var var_ast *ZendAst = ast.Child(0)
	var expr_ast *ZendAst = ast.Child(1)
	var opcode uint32 = ast.Attr()
	var var_node Znode
	var expr_node Znode
	var opline *types.ZendOp
	var offset uint32
	var cache_slot uint32
	ZendEnsureWritableVariable(var_ast)
	switch var_ast.Kind() {
	case ZEND_AST_VAR:
		offset = ZendDelayedCompileBegin()
		compiler.DelayedCompileVar(&var_node, var_ast, BP_VAR_RW, 0)
		compiler.CompileExpr(&expr_node, expr_ast)
		ZendDelayedCompileEnd(offset)
		opline = ZendEmitOp(result, ZEND_ASSIGN_OP, &var_node, &expr_node)
		opline.SetExtendedValue(opcode)
		return
	case ZEND_AST_STATIC_PROP:
		offset = ZendDelayedCompileBegin()
		compiler.DelayedCompileVar(result, var_ast, BP_VAR_RW, 0)
		compiler.CompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		cache_slot = opline.GetExtendedValue()
		opline.SetOpcode(ZEND_ASSIGN_STATIC_PROP_OP)
		opline.SetExtendedValue(opcode)
		opline = ZendEmitOpData(&expr_node)
		opline.SetExtendedValue(cache_slot)
		return
	case ZEND_AST_DIM:
		offset = ZendDelayedCompileBegin()
		compiler.DelayedCompileDim(result, var_ast, BP_VAR_RW)
		compiler.CompileExpr(&expr_node, expr_ast)
		opline = ZendDelayedCompileEnd(offset)
		opline.SetOpcode(ZEND_ASSIGN_DIM_OP)
		opline.SetExtendedValue(opcode)
		ZendEmitOpData(&expr_node)
		return
	case ZEND_AST_PROP:
		offset = ZendDelayedCompileBegin()
		compiler.DelayedCompileProp(result, var_ast, BP_VAR_RW)
		compiler.CompileExpr(&expr_node, expr_ast)
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
func (compiler *Compiler) CompileArgs(ast *ZendAst, fbc types.IFunction) uint32 {
	var args *ZendAstList = ast.AsAstList()
	var i uint32
	var uses_arg_unpack bool = 0
	var arg_count uint32 = 0
	for i = 0; i < args.GetChildren(); i++ {
		var arg *ZendAst = args.Children()[i]
		var arg_num uint32 = i + 1
		var arg_node Znode
		var opline *types.ZendOp
		var opcode OpCode
		if arg.Kind() == ZEND_AST_UNPACK {
			uses_arg_unpack = 1
			fbc = nil
			compiler.CompileExpr(&arg_node, arg.Children()[0])
			opline = ZendEmitOp(nil, ZEND_SEND_UNPACK, &arg_node, nil)
			opline.GetOp2().SetNum(arg_count)
			opline.GetResult().SetVar(uint32(types.ZendIntptrT((*ZendExecuteData)(nil).Arg(arg_count))))
			continue
		}
		if uses_arg_unpack != 0 {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use positional argument after argument unpacking")
		}
		arg_count++
		if ZendIsVariableOrCall(arg) != 0 {
			if ZendIsCall(arg) != 0 {
				compiler.CompileVar(&arg_node, arg, BP_VAR_R, 0)
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
					compiler.CompileVar(&arg_node, arg, BP_VAR_W, 1)
					opcode = ZEND_SEND_REF
				} else {
					compiler.CompileVar(&arg_node, arg, BP_VAR_R, 0)
					if arg_node.GetOpType() == IS_TMP_VAR {
						opcode = ZEND_SEND_VAL
					} else {
						opcode = ZEND_SEND_VAR
					}
				}
			} else {
				for {
					if arg.Kind() == ZEND_AST_VAR {
						compiler.setLinenoByAst(ast)
						if IsThisFetch(arg) {
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
					compiler.CompileVar(&arg_node, arg, BP_VAR_FUNC_ARG, 1)
					opcode = ZEND_SEND_FUNC_ARG
					break
				}
			}
		} else {
			compiler.CompileExpr(&arg_node, arg)
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
func ZendGetCallOp(init_op *types.ZendOp, fbc types.IFunction) uint8 {
	if fbc != nil {
		if fbc.GetType() == ZEND_INTERNAL_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS) == 0 {
			if init_op.GetOpcode() == ZEND_INIT_FCALL && ZendExecuteInternal == nil {
				if !fbc.HasFnFlags(types.AccAbstract | types.AccDeprecated | types.AccHasTypeHints | types.AccReturnReference) {
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
func (compiler *Compiler) CompileCallCommon(result *Znode, args_ast *ZendAst, fbc types.IFunction) {
	var opline *types.ZendOp
	var opnum_init uint32 = GetNextOpNumber() - 1
	var arg_count uint32
	arg_count = compiler.CompileArgs(args_ast, fbc)
	ZendDoExtendedFcallBegin()
	opline = CG__().GetActiveOpArray().GetOpcodes()[opnum_init]
	opline.SetExtendedValue(arg_count)
	if opline.GetOpcode() == ZEND_INIT_FCALL {
		opline.GetOp1().SetNum(ZendVmCalcUsedStack(arg_count, fbc))
	}
	opline = ZendEmitOp(result, ZendGetCallOp(opline, fbc), nil, nil)
	ZendDoExtendedFcallEnd()
}
func (compiler *Compiler) CompileFunctionName(name_node *Znode, name_ast *ZendAst) bool {
	orig_name := ZendAstGetStrVal(name_ast)
	resolveName, isFullyQualified := ZendResolveFunctionName(orig_name, name_ast.Attr())

	name_node.SetOpType(IS_CONST)
	name_node.GetConstant().SetString(resolveName)
	return !isFullyQualified && FC__().GetCurrentNamespace() != nil
}
func (compiler *Compiler) CompileNsCall(result *Znode, name_node *Znode, args_ast *ZendAst) {
	var opline *types.ZendOp = GetNextOp()
	opline.SetOpcode(ZEND_INIT_NS_FCALL_BY_NAME)
	opline.SetOp2Type(IS_CONST)
	opline.GetOp2().SetConstant(ZendAddNsFuncNameLiteral(name_node.GetConstant().String()))
	opline.GetResult().SetNum(ZendAllocCacheSlot())
	compiler.CompileCallCommon(result, args_ast, nil)
}
func (compiler *Compiler) CompileDynamicCall(result *Znode, name_node *Znode, args_ast *ZendAst) {
	if name_node.GetOpType() == IS_CONST && name_node.GetConstant().IsString() {
		var colon *byte
		var str *types.String = name_node.GetConstant().String()
		if lang.Assign(&colon, operators.ZendMemrchr(str.GetVal(), ':', str.GetLen())) != nil && colon > str.GetVal() && (*(colon - 1)) == ':' {
			var class *types.String = types.NewString(b.CastStr(str.GetVal(), colon-str.GetVal()-1))
			var method *types.String = types.NewString(b.CastStr(colon+1, str.GetLen()-(colon-str.GetVal())-1))
			var opline *types.ZendOp = GetNextOp()
			opline.SetOpcode(ZEND_INIT_STATIC_METHOD_CALL)
			opline.SetOp1Type(IS_CONST)
			opline.GetOp1().SetConstant(ZendAddClassNameLiteral(class))
			opline.SetOp2Type(IS_CONST)
			opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(method))

			/* 2 slots, for class and method */

			opline.GetResult().SetNum(ZendAllocCacheSlots(2))
			// ZvalPtrDtor(name_node.GetConstant())
		} else {
			var opline *types.ZendOp = GetNextOp()
			opline.SetOpcode(ZEND_INIT_FCALL_BY_NAME)
			opline.SetOp2Type(IS_CONST)
			opline.GetOp2().SetConstant(ZendAddFuncNameLiteral(str))
			opline.GetResult().SetNum(ZendAllocCacheSlot())
		}
	} else {
		ZendEmitOp(nil, ZEND_INIT_DYNAMIC_CALL, nil, name_node)
	}
	compiler.CompileCallCommon(result, args_ast, nil)
}
func ZendArgsContainUnpack(args *ZendAstList) bool {
	var i uint32
	for i = 0; i < args.GetChildren(); i++ {
		if args.Children()[i].Kind() == ZEND_AST_UNPACK {
			return 1
		}
	}
	return 0
}
func (compiler *Compiler) CompileFuncStrlen(result *Znode, args *ZendAstList) int {
	var arg_node Znode
	if (CG__().GetCompilerOptions()&ZEND_COMPILE_NO_BUILTIN_STRLEN) != 0 || args.GetChildren() != 1 {
		return types.FAILURE
	}
	compiler.CompileExpr(&arg_node, args.Children()[0])
	if arg_node.GetOpType() == IS_CONST && arg_node.GetConstant().IsString() {
		result.SetOpType(IS_CONST)
		result.GetConstant().SetLong(arg_node.GetConstant().String().GetLen())
	} else {
		ZendEmitOpTmp(result, ZEND_STRLEN, &arg_node, nil)
	}
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncTypecheck(result *Znode, args *ZendAstList, type_ uint32) int {
	var arg_node Znode
	var opline *types.ZendOp
	if args.GetChildren() != 1 {
		return types.FAILURE
	}
	compiler.CompileExpr(&arg_node, args.Children()[0])
	opline = ZendEmitOpTmp(result, ZEND_TYPE_CHECK, &arg_node, nil)
	if type_ != types.IsBool {
		opline.SetExtendedValue(1 << type_)
	} else {
		opline.SetExtendedValue(1<<types.IsFalse | 1<<types.IsTrue)
	}
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncCast(result *Znode, args *ZendAstList, type_ uint32) int {
	var arg_node Znode
	var opline *types.ZendOp
	if args.GetChildren() != 1 {
		return types.FAILURE
	}
	compiler.CompileExpr(&arg_node, args.Children()[0])
	opline = ZendEmitOpTmp(result, ZEND_CAST, &arg_node, nil)
	opline.SetExtendedValue(type_)
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncDefined(result *Znode, args *ZendAstList) int {
	var name *types.String
	var opline *types.ZendOp
	if args.GetChildren() != 1 || args.Children()[0].Kind() != ZEND_AST_ZVAL {
		return types.FAILURE
	}
	name = operators.ZvalGetString(args.Child(0).Val())
	if operators.ZendMemrchr(name.GetVal(), '\\', name.GetLen()) || operators.ZendMemrchr(name.GetVal(), ':', name.GetLen()) {
		// types.ZendStringReleaseEx(name, 0)
		return types.FAILURE
	}
	if ZendTryCtEvalConst(result.GetConstant(), name.GetStr(), false) != 0 {
		// types.ZendStringReleaseEx(name, 0)
		// ZvalPtrDtor(result.GetConstant())
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
	var lcname *types.String = operators.ZendStringTolower(name)
	c.SetStringEx(lcname)
	ZendAddLiteral(&c)
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncChr(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 1 && args.Children()[0].Kind() == ZEND_AST_ZVAL && ZendAstGetZval(args.Children()[0]).IsLong() {
		var c ZendLong = ZendAstGetZval(args.Children()[0]).Long() & 0xff
		result.SetOpType(IS_CONST)
		result.GetConstant().SetString(string(byte(c)))
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func (compiler *Compiler) CompileFuncOrd(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 1 && args.Children()[0].Kind() == ZEND_AST_ZVAL && ZendAstGetZval(args.Children()[0]).IsString() {
		result.SetOpType(IS_CONST)
		result.GetConstant().SetLong(uint8(ZendAstGetZval(args.Children()[0]).String().GetStr()[0]))
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func FbcIsFinalized(fbc types.IFunction) bool {
	return !(ZEND_USER_CODE(fbc.GetType())) || fbc.IsDonePassTwo()
}
func ZendTryCompileCtBoundInitUserFunc(name_ast *ZendAst, num_args uint32) int {
	var lcname *types.String
	var fbc types.IFunction
	var opline *types.ZendOp
	if name_ast.Kind() != ZEND_AST_ZVAL || name_ast.Val().Type() != types.IsString {
		return types.FAILURE
	}
	name := ZendAstGetStr(name_ast)
	lcname = operators.ZendStringTolower(name)
	fbc = CG__().FunctionTable().Get(lcname.GetStr())
	if fbc == nil || FbcIsFinalized(fbc) == 0 || fbc.GetType() == ZEND_INTERNAL_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_INTERNAL_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_USER_FUNCTIONS) != 0 || fbc.GetType() == ZEND_USER_FUNCTION && (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_OTHER_FILES) != 0 && fbc.GetOpArray().GetFilename() != CG__().GetActiveOpArray().GetFilename() {
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
func (compiler *Compiler) CompileInitUserFunc(name_ast *ZendAst, num_args uint32, orig_func_name *types.String) {
	var opline *types.ZendOp
	var name_node Znode
	if ZendTryCompileCtBoundInitUserFunc(name_ast, num_args) == types.SUCCESS {
		return
	}
	compiler.CompileExpr(&name_node, name_ast)
	opline = ZendEmitOp(nil, ZEND_INIT_USER_CALL, nil, &name_node)
	opline.SetOp1Type(IS_CONST)
	LITERAL_STR(opline.GetOp1(), orig_func_name.Copy())
	opline.SetExtendedValue(num_args)
}
func (compiler *Compiler) CompileFuncCufa(result *Znode, args *ZendAstList, lcname *types.String) int {
	var arg_node Znode
	if args.GetChildren() != 2 {
		return types.FAILURE
	}
	compiler.CompileInitUserFunc(args.Children()[0], 0, lcname)
	if args.Children()[1].Kind() == ZEND_AST_CALL && args.Children()[1].Children()[0].Kind() == ZEND_AST_ZVAL && ZendAstGetZval(args.Children()[1].Children()[0]).IsString() && args.Children()[1].Children()[1].Kind() == ZEND_AST_ARG_LIST {
		var orig_name *types.String = ZendAstGetStr(args.Children()[1].Children()[0])
		var list *ZendAstList = args.Children()[1].Children()[1].AsAstList()
		name, _ := ZendResolveFunctionName(orig_name.GetStr(), args.Children()[1].Children()[0].Attr())
		if ascii.StrCaseEquals(name, "array_slice") && list.GetChildren() == 3 && list.Children()[1].Kind() == ZEND_AST_ZVAL {
			var zv *types.Zval = ZendAstGetZval(list.Children()[1])
			if zv.IsLong() && zv.Long() >= 0 && zv.Long() <= 0x7fffffff {
				var opline *types.ZendOp
				var len_node Znode
				compiler.CompileExpr(&arg_node, list.Children()[0])
				compiler.CompileExpr(&len_node, list.Children()[2])
				opline = ZendEmitOp(nil, ZEND_SEND_ARRAY, &arg_node, &len_node)
				opline.SetExtendedValue(zv.Long())
				ZendEmitOp(result, ZEND_DO_FCALL, nil, nil)
				// types.ZendStringReleaseEx(name, 0)
				return types.SUCCESS
			}
		}
		// types.ZendStringReleaseEx(name, 0)
	}
	compiler.CompileExpr(&arg_node, args.Children()[1])
	ZendEmitOp(nil, ZEND_SEND_ARRAY, &arg_node, nil)
	ZendEmitOp(result, ZEND_DO_FCALL, nil, nil)
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncCuf(result *Znode, args *ZendAstList, lcname *types.String) int {
	var i uint32
	if args.GetChildren() < 1 {
		return types.FAILURE
	}
	compiler.CompileInitUserFunc(args.Children()[0], args.GetChildren()-1, lcname)
	for i = 1; i < args.GetChildren(); i++ {
		var arg_ast *ZendAst = args.Children()[i]
		var arg_node Znode
		var opline *types.ZendOp
		compiler.CompileExpr(&arg_node, arg_ast)
		opline = ZendEmitOp(nil, ZEND_SEND_USER, &arg_node, nil)
		opline.GetOp2().SetNum(i)
		opline.GetResult().SetVar(uint32(types.ZendIntptrT(nil.Arg(i))))
	}
	ZendEmitOp(result, ZEND_DO_FCALL, nil, nil)
	return types.SUCCESS
}
func (compiler *Compiler) CompileAssert(result *Znode, args *ZendAstList, name *types.String, fbc types.IFunction) {
	if EG__().GetAssertions() >= 0 {
		var name_node Znode
		var opline *types.ZendOp
		var check_op_number uint32 = GetNextOpNumber()
		ZendEmitOp(nil, ZEND_ASSERT_CHECK, nil, nil)
		if fbc != nil && FbcIsFinalized(fbc) != 0 {
			name_node.SetOpType(IS_CONST)
			name_node.GetConstant().SetString(name.GetStr())
			opline = ZendEmitOp(nil, ZEND_INIT_FCALL, nil, &name_node)
		} else {
			opline = ZendEmitOp(nil, ZEND_INIT_NS_FCALL_BY_NAME, nil, nil)
			opline.SetOp2Type(IS_CONST)
			opline.GetOp2().SetConstant(ZendAddNsFuncNameLiteral(name))
		}
		opline.GetResult().SetNum(ZendAllocCacheSlot())
		if args.GetChildren() == 1 && (args.Children()[0].Kind() != ZEND_AST_ZVAL || ZendAstGetZval(args.Children()[0]).Type() != types.IsString) {

			/* add "assert(condition) as assertion message */
			args.AddChild(ZendAstCreateZvalFromStr("assert()"))
		}
		compiler.CompileCallCommon(result, (*ZendAst)(args), fbc)
		opline = CG__().GetActiveOpArray().GetOpcodes()[check_op_number]
		opline.GetOp2().SetOplineNum(GetNextOpNumber())
		opline.SetResultType(result.GetOpType())
		if result.GetOpType() == IS_CONST {
			opline.GetResult().SetConstant(ZendAddLiteral(result.GetConstant()))
		} else {
			opline.SetResult(result.GetOp())
		}
	} else {
		result.SetOpType(IS_CONST)
		result.GetConstant().SetTrue()
	}
}
func (compiler *Compiler) CompileFuncInArray(result *Znode, args *ZendAstList) int {
	var strict bool = 0
	var array Znode
	var needly Znode
	var opline *types.ZendOp
	if args.GetChildren() == 3 {
		if args.Children()[2].Kind() == ZEND_AST_ZVAL {
			strict = operators.IZendIsTrue(ZendAstGetZval(args.Children()[2]))
		} else if args.Children()[2].Kind() == ZEND_AST_CONST {
			var value types.Zval
			var name_ast *ZendAst = args.Children()[2].Children()[0]
			var resolved_name, isFullyQualified = ZendResolveConstName(ZendAstGetStr(name_ast).GetStr(), name_ast.Attr())
			if ZendTryCtEvalConst(&value, resolved_name, isFullyQualified) == 0 {
				return types.FAILURE
			}
			strict = operators.IZendIsTrue(&value)
		} else {
			return types.FAILURE
		}
	} else if args.GetChildren() != 2 {
		return types.FAILURE
	}
	if args.Children()[1].Kind() != ZEND_AST_ARRAY || compiler.TryCtEvalArray(array.GetConstant(), args.Children()[1]) == 0 {
		return types.FAILURE
	}
	if array.GetConstant().Array().Len() > 0 {
		var ok bool = 1
		var val *types.Zval
		var tmp types.Zval
		var src *types.Array = array.GetConstant().Array()
		var dst *types.Array = types.NewArray(src.Len())
		tmp.SetTrue()
		if strict != 0 {
			src.ForeachEx(func(_ types.ArrayKey, val *types.Zval) bool {
				if val.IsString() {
					dst.KeyAdd(val.String().GetStr(), &tmp)
				} else if val.IsLong() {
					dst.IndexAdd(val.Long(), &tmp)
				} else {
					dst.Destroy()
					ok = 0
					return false
				}
				return true
			})
		} else {
			var __ht *types.Array = src
			for _, _p := range __ht.ForeachData() {
				var _z *types.Zval = _p.GetVal()

				val = _z
				if !val.IsString() || operators.IsNumericString(val.String().GetStr(), nil, nil, 0) != 0 {
					dst.Destroy()
					ok = 0
					break
				}
				dst.KeyAdd(val.String().GetStr(), &tmp)
			}
		}
		src.Destroy()
		if ok == 0 {
			return types.FAILURE
		}
		array.GetConstant().Array() = dst
	}
	array.SetOpType(IS_CONST)
	compiler.CompileExpr(&needly, args.Children()[0])
	opline = ZendEmitOpTmp(result, ZEND_IN_ARRAY, &needly, &array)
	opline.SetExtendedValue(strict)
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncCount(result *Znode, args *ZendAstList, lcname *types.String) int {
	var arg_node Znode
	var opline *types.ZendOp
	if args.GetChildren() != 1 {
		return types.FAILURE
	}
	compiler.CompileExpr(&arg_node, args.Children()[0])
	opline = ZendEmitOpTmp(result, ZEND_COUNT, &arg_node, nil)
	opline.SetExtendedValue(uint32(types.IntBool(lcname.GetStr() == "sizeof")))
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncGetClass(result *Znode, args *ZendAstList) int {
	if args.GetChildren() == 0 {
		ZendEmitOpTmp(result, ZEND_GET_CLASS, nil, nil)
	} else {
		var arg_node Znode
		if args.GetChildren() != 1 {
			return types.FAILURE
		}
		compiler.CompileExpr(&arg_node, args.Children()[0])
		ZendEmitOpTmp(result, ZEND_GET_CLASS, &arg_node, nil)
	}
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncGetCalledClass(result *Znode, args *ZendAstList) int {
	if args.GetChildren() != 0 {
		return types.FAILURE
	}
	ZendEmitOpTmp(result, ZEND_GET_CALLED_CLASS, nil, nil)
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncGettype(result *Znode, args *ZendAstList) int {
	var arg_node Znode
	if args.GetChildren() != 1 {
		return types.FAILURE
	}
	compiler.CompileExpr(&arg_node, args.Children()[0])
	ZendEmitOpTmp(result, ZEND_GET_TYPE, &arg_node, nil)
	return types.SUCCESS
}
func (compiler *Compiler) CompileFuncNumArgs(result *Znode, args *ZendAstList) int {
	if CG__().GetActiveOpArray().FunctionName() != "" && args.GetChildren() == 0 {
		ZendEmitOpTmp(result, ZEND_FUNC_NUM_ARGS, nil, nil)
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
