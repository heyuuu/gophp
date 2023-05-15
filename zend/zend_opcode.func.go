package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"sort"
)

func ZendExtensionOpArrayCtorHandler(extension *ZendExtension, op_array *types.ZendOpArray) {
	if extension.GetOpArrayCtor() != nil {
		extension.GetOpArrayCtor()(op_array)
	}
}
func ZendExtensionOpArrayDtorHandler(extension *ZendExtension, op_array *types.ZendOpArray) {
	if extension.GetOpArrayDtor() != nil {
		extension.GetOpArrayDtor()(op_array)
	}
}
func InitOpArray(op_array *types.ZendOpArray, initial_ops_size int) {
	op_array.Init()
	op_array.SetRefcount((*uint32)(Emalloc(b.SizeOf("uint32_t"))))
	op_array.refcount = 1
	op_array.SetLast(0)
	op_array.SetOpcodes(Emalloc(initial_ops_size * b.SizeOf("zend_op")))
	op_array.SetLastVar(0)
	op_array.SetVars(nil)
	op_array.SetT(0)
	op_array.SetFunctionName(nil)
	op_array.SetFilename(ZendGetCompiledFilename())
	op_array.SetDocComment(nil)
	op_array.SetArgInfo(nil)
	op_array.SetNumArgs(0)
	op_array.SetRequiredNumArgs(0)
	op_array.SetScope(nil)
	op_array.SetPrototype(nil)
	op_array.SetLiveRange(nil)
	op_array.SetTryCatchArray(nil)
	op_array.SetLastLiveRange(0)
	op_array.SetStaticVariables(nil)
	ZEND_MAP_PTR_INIT(op_array.static_variables_ptr, op_array.GetStaticVariables())
	op_array.SetLastTryCatch(0)
	op_array.SetFnFlags(0)
	op_array.SetLastLiteral(0)
	op_array.SetLiterals(nil)
	ZEND_MAP_PTR_INIT(op_array.run_time_cache, nil)
	op_array.SetCacheSize(ZendOpArrayExtensionHandles * b.SizeOf("void *"))
	memset(op_array.GetReserved(), 0, types.ZEND_MAX_RESERVED_RESOURCES*b.SizeOf("void *"))
	if (ZendExtensionFlags & ZEND_EXTENSIONS_HAVE_OP_ARRAY_CTOR) != 0 {
		ZendExtensions.ApplyWithArgument(LlistApplyWithArgFuncT(ZendExtensionOpArrayCtorHandler), op_array)
	}
}
func ZendFunctionDtor(zv *types.Zval) {
	var function types.IFunction = zv.Ptr()
	if function.GetType() == ZEND_USER_FUNCTION {
		b.Assert(function.GetFunctionName() != nil)
		DestroyOpArray(function.GetOpArray())
	} else {
		b.Assert(function.GetType() == ZEND_INTERNAL_FUNCTION)
		b.Assert(function.GetFunctionName() != nil)
		// types.ZendStringReleaseEx(function.GetFunctionName(), 1)

		/* For methods this will be called explicitly. */

		if function.GetScope() == nil {
			//ZendFreeInternalArgInfo(function.GetInternalFunction())
		}
		if !function.IsArenaAllocated() {
			Pefree(function)
		}
	}
}
func ZendFunctionDtorEx(function types.IFunction) {
	//if function.GetType() == ZEND_USER_FUNCTION {
	//	b.Assert(function.GetFunctionName() != nil)
	//	DestroyOpArray(function.GetOpArray())
	//} else {
	//	b.Assert(function.GetType() == ZEND_INTERNAL_FUNCTION)
	//	b.Assert(function.GetFunctionName() != nil)
	//}
}
func DestroyZendClass(zv *types.Zval) {
	DestroyZendClassEntry(zv.Ptr().(*types.ClassEntry))
}
func DestroyZendClassEntry(ce *types.ClassEntry) {
	Efree(ce)
}
func DestroyOpArray(op_array *types.ZendOpArray) {
	Efree(op_array)
}
func ZendUpdateExtendedStmts(op_array *types.ZendOpArray) {
	var opline *ZendOp = op_array.GetOpcodes()
	var end *ZendOp = opline + op_array.GetLast()
	for opline < end {
		if opline.GetOpcode() == ZEND_EXT_STMT {
			if opline+1 < end {
				if (opline + 1).GetOpcode() == ZEND_EXT_STMT {
					opline.SetOpcode(ZEND_NOP)
					opline++
					continue
				}
				if opline+1 < end {
					opline.SetLineno((opline + 1).GetLineno())
				}
			} else {
				opline.SetOpcode(ZEND_NOP)
			}
		}
		opline++
	}
}
func ZendExtensionOpArrayHandler(extension *ZendExtension, op_array *types.ZendOpArray) {
	if extension.GetOpArrayHandler() != nil {
		extension.GetOpArrayHandler()(op_array)
	}
}
func ZendCheckFinallyBreakout(op_array *types.ZendOpArray, op_num uint32, dst_num uint32) {
	var i int
	for i = 0; i < op_array.GetLastTryCatch(); i++ {
		if (op_num < op_array.GetTryCatchArray()[i].GetFinallyOp() || op_num >= op_array.GetTryCatchArray()[i].GetFinallyEnd()) && (dst_num >= op_array.GetTryCatchArray()[i].GetFinallyOp() && dst_num <= op_array.GetTryCatchArray()[i].GetFinallyEnd()) {
			CG__().SetInCompilation(1)
			CG__().SetActiveOpArray(op_array)
			CG__().SetZendLineno(op_array.GetOpcodes()[op_num].GetLineno())
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "jump into a finally block is disallowed")
		} else if op_num >= op_array.GetTryCatchArray()[i].GetFinallyOp() && op_num <= op_array.GetTryCatchArray()[i].GetFinallyEnd() && (dst_num > op_array.GetTryCatchArray()[i].GetFinallyEnd() || dst_num < op_array.GetTryCatchArray()[i].GetFinallyOp()) {
			CG__().SetInCompilation(1)
			CG__().SetActiveOpArray(op_array)
			CG__().SetZendLineno(op_array.GetOpcodes()[op_num].GetLineno())
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "jump out of a finally block is disallowed")
		}
	}
}
func ZendGetBrkContTarget(op_array *types.ZendOpArray, opline *ZendOp) uint32 {
	var nest_levels int = opline.GetOp2().GetNum()
	var array_offset int = opline.GetOp1().GetNum()
	var jmp_to *ZendBrkContElement
	for {
		jmp_to = CG__().GetContext().GetBrkContArray()[array_offset]
		if nest_levels > 1 {
			array_offset = jmp_to.GetParent()
		}
		if b.PreDec(&nest_levels) <= 0 {
			break
		}
	}
	if opline.GetOpcode() == ZEND_BRK {
		return jmp_to.GetBrk()
	} else {
		return jmp_to.GetCont()
	}
}
func EmitLiveRangeRaw(op_array *types.ZendOpArray, var_num uint32, kind uint32, start uint32, end uint32) {
	var range_ *ZendLiveRange
	op_array.GetLastLiveRange()++
	op_array.SetLiveRange(Erealloc(op_array.GetLiveRange(), b.SizeOf("zend_live_range")*op_array.GetLastLiveRange()))
	b.Assert(start < end)
	range_ = op_array.GetLiveRange()[op_array.GetLastLiveRange()-1]
	range_.SetVar(uint32(intPtr(nil.VarNum(op_array.GetLastVar() + var_num))))
	range_.SetVar(range_.GetVar() | kind)
	range_.SetStart(start)
	range_.SetEnd(end)
}
func EmitLiveRange(op_array *types.ZendOpArray, var_num uint32, start uint32, end uint32, needs_live_range ZendNeedsLiveRangeCb) {
	var def_opline *ZendOp = op_array.GetOpcodes()[start]
	var orig_def_opline *ZendOp = def_opline
	var use_opline *ZendOp = op_array.GetOpcodes()[end]
	var kind uint32
	switch def_opline.GetOpcode() {
	case ZEND_ADD_ARRAY_ELEMENT:
		fallthrough
	case ZEND_ADD_ARRAY_UNPACK:
		fallthrough
	case ZEND_ROPE_ADD:
		b.Assert(false)
		return
	case ZEND_JMPZ_EX:
		fallthrough
	case ZEND_JMPNZ_EX:
		fallthrough
	case ZEND_BOOL:
		fallthrough
	case ZEND_BOOL_NOT:
		fallthrough
	case ZEND_FETCH_CLASS:
		fallthrough
	case ZEND_DECLARE_ANON_CLASS:
		fallthrough
	case ZEND_FAST_CALL:
		return
	case ZEND_BEGIN_SILENCE:
		kind = ZEND_LIVE_SILENCE
		start++
	case ZEND_ROPE_INIT:
		kind = ZEND_LIVE_ROPE

		/* ROPE live ranges include the generating opcode. */

		def_opline--
	case ZEND_FE_RESET_R:
		fallthrough
	case ZEND_FE_RESET_RW:
		kind = ZEND_LIVE_LOOP
		start++
	case ZEND_NEW:
		var level int = 0
		var orig_start uint32 = start
		for def_opline+1 < use_opline {
			def_opline++
			start++
			if def_opline.GetOpcode() == ZEND_DO_FCALL {
				if level == 0 {
					break
				}
				level--
			} else {
				switch def_opline.GetOpcode() {
				case ZEND_INIT_FCALL:
					fallthrough
				case ZEND_INIT_FCALL_BY_NAME:
					fallthrough
				case ZEND_INIT_NS_FCALL_BY_NAME:
					fallthrough
				case ZEND_INIT_DYNAMIC_CALL:
					fallthrough
				case ZEND_INIT_USER_CALL:
					fallthrough
				case ZEND_INIT_METHOD_CALL:
					fallthrough
				case ZEND_INIT_STATIC_METHOD_CALL:
					fallthrough
				case ZEND_NEW:
					level++
				case ZEND_DO_ICALL:
					fallthrough
				case ZEND_DO_UCALL:
					fallthrough
				case ZEND_DO_FCALL_BY_NAME:
					level--
				}
			}
		}
		EmitLiveRangeRaw(op_array, var_num, ZEND_LIVE_NEW, orig_start+1, start+1)
		if start+1 == end {

			/* Trivial live-range, no need to store it. */

			return

			/* Trivial live-range, no need to store it. */

		}
		fallthrough
	default:
		start++
		kind = ZEND_LIVE_TMPVAR

		/* Check hook to determine whether a live range is necessary,
		 * e.g. based on type info. */

		if needs_live_range != nil && needs_live_range(op_array, orig_def_opline) == 0 {
			return
		}
	case ZEND_COPY_TMP:

		/* COPY_TMP has a split live-range: One from the definition until the use in
		 * "null" branch, and another from the start of the "non-null" branch to the
		 * FREE opcode. */

		var rt_var_num uint32 = uint32(intPtr(nil.VarNum(op_array.GetLastVar() + var_num)))
		var block_start_op *ZendOp = use_opline
		if needs_live_range != nil && needs_live_range(op_array, orig_def_opline) == 0 {
			return
		}
		for (block_start_op - 1).opcode == ZEND_FREE {
			block_start_op--
		}
		kind = ZEND_LIVE_TMPVAR
		start = block_start_op - op_array.GetOpcodes()
		if start != end {
			EmitLiveRangeRaw(op_array, var_num, kind, start, end)
		}
		for {
			use_opline--
			if (use_opline.GetOp1Type()&(IS_TMP_VAR|IS_VAR)) != 0 && use_opline.GetOp1().GetVar() == rt_var_num || (use_opline.GetOp2Type()&(IS_TMP_VAR|IS_VAR)) != 0 && use_opline.GetOp2().GetVar() == rt_var_num {
				break
			}
		}
		start = def_opline + 1 - op_array.GetOpcodes()
		end = use_opline - op_array.GetOpcodes()
		EmitLiveRangeRaw(op_array, var_num, kind, start, end)
		return
	}
	EmitLiveRangeRaw(op_array, var_num, kind, start, end)
}
func IsFakeDef(opline *ZendOp) types.ZendBool {
	/* These opcodes only modify the result, not create it. */

	return opline.GetOpcode() == ZEND_ROPE_ADD || opline.GetOpcode() == ZEND_ADD_ARRAY_ELEMENT || opline.GetOpcode() == ZEND_ADD_ARRAY_UNPACK

	/* These opcodes only modify the result, not create it. */
}
func KeepsOp1Alive(opline *ZendOp) types.ZendBool {
	/* These opcodes don't consume their OP1 operand,
	 * it is later freed by something else. */

	if opline.GetOpcode() == ZEND_CASE || opline.GetOpcode() == ZEND_SWITCH_LONG || opline.GetOpcode() == ZEND_FETCH_LIST_R || opline.GetOpcode() == ZEND_COPY_TMP {
		return 1
	}
	b.Assert(opline.GetOpcode() != ZEND_SWITCH_STRING && opline.GetOpcode() != ZEND_FE_FETCH_R && opline.GetOpcode() != ZEND_FE_FETCH_RW && opline.GetOpcode() != ZEND_FETCH_LIST_W && opline.GetOpcode() != ZEND_VERIFY_RETURN_TYPE && opline.GetOpcode() != ZEND_BIND_LEXICAL && opline.GetOpcode() != ZEND_ROPE_ADD)
	return 0
}
func CmpLiveRange(a *ZendLiveRange, b *ZendLiveRange) int { return a.GetStart() - b.GetStart() }
func SwapLiveRange(a *ZendLiveRange, b *ZendLiveRange) {
	*a, *b = *b, *a
}
func ZendCalcLiveRanges(op_array *types.ZendOpArray, needs_live_range ZendNeedsLiveRangeCb) {
	var opnum uint32 = op_array.GetLast()
	var opline *ZendOp = op_array.GetOpcodes()[opnum]
	var var_offset uint32 = op_array.GetLastVar()
	var last_use *uint32 = DoAlloca(b.SizeOf("uint32_t")*op_array.GetT(), use_heap)
	memset(last_use, -1, b.SizeOf("uint32_t")*op_array.GetT())
	b.Assert(op_array.GetLiveRange() == nil)
	for opnum > 0 {
		opnum--
		opline--
		if (opline.GetResultType()&(IS_TMP_VAR|IS_VAR)) != 0 && IsFakeDef(opline) == 0 {
			var var_num uint32 = EX_VAR_TO_NUM(opline.GetResult().GetVar()) - var_offset

			/* Defs without uses can occur for two reasons: Either because the result is
			 * genuinely unused (e.g. omitted FREE opcode for an unused boolean result), or
			 * because there are multiple defining opcodes (e.g. JMPZ_EX and QM_ASSIGN), in
			 * which case the last one starts the live range. As such, we can simply ignore
			 * missing uses here. */

			if last_use[var_num] != uint32-1 {

				/* Skip trivial live-range */

				if opnum+1 != last_use[var_num] {
					var num uint32

					/* OP_DATA uses only op1 operand */

					b.Assert(opline.GetOpcode() != ZEND_OP_DATA)
					num = opnum
					EmitLiveRange(op_array, var_num, num, last_use[var_num], needs_live_range)
				}
				last_use[var_num] = uint32 - 1
			}

			/* Defs without uses can occur for two reasons: Either because the result is
			 * genuinely unused (e.g. omitted FREE opcode for an unused boolean result), or
			 * because there are multiple defining opcodes (e.g. JMPZ_EX and QM_ASSIGN), in
			 * which case the last one starts the live range. As such, we can simply ignore
			 * missing uses here. */

		}
		if (opline.GetOp1Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
			var var_num uint32 = EX_VAR_TO_NUM(opline.GetOp1().GetVar()) - var_offset
			if last_use[var_num] == uint32-1 {
				if KeepsOp1Alive(opline) == 0 {

					/* OP_DATA is really part of the previous opcode. */

					last_use[var_num] = opnum - (opline.GetOpcode() == ZEND_OP_DATA)

					/* OP_DATA is really part of the previous opcode. */

				}
			}
		}
		if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
			var var_num uint32 = EX_VAR_TO_NUM(opline.GetOp2().GetVar()) - var_offset
			if opline.GetOpcode() == ZEND_FE_FETCH_R || opline.GetOpcode() == ZEND_FE_FETCH_RW {

				/* OP2 of FE_FETCH is actually a def, not a use. */

				if last_use[var_num] != uint32-1 {
					if opnum+1 != last_use[var_num] {
						EmitLiveRange(op_array, var_num, opnum, last_use[var_num], needs_live_range)
					}
					last_use[var_num] = uint32 - 1
				}

				/* OP2 of FE_FETCH is actually a def, not a use. */

			} else if last_use[var_num] == uint32-1 {

				/* OP_DATA uses only op1 operand */

				b.Assert(opline.GetOpcode() != ZEND_OP_DATA)
				last_use[var_num] = opnum
			}
		}
	}
	if op_array.GetLastLiveRange() > 1 {
		var r1 *ZendLiveRange = op_array.GetLiveRange()
		var r2 *ZendLiveRange = r1 + op_array.GetLastLiveRange() - 1

		/* In most cases we need just revert the array */
		for r1 < r2 {
			SwapLiveRange(r1, r2)
			r1++
			r2--
		}
		r1 = op_array.GetLiveRange()
		r2 = r1 + op_array.GetLastLiveRange() - 1
		for r1 < r2 {
			if r1.GetStart() > (r1 + 1).GetStart() {
				//ZendSort(r1, r2-r1+1, b.SizeOf("zend_live_range"), types.CompareFuncT(CmpLiveRange), types.SwapFuncT(SwapLiveRange))
				var r []ZendLiveRange = r1[:r2-r1+1]
				sort.Slice(r, func(i, j int) bool {
					return r[i].GetStart() < r[j].GetStart()
				})
				break
			}
			r1++
		}
	}
	FreeAlloca(last_use, use_heap)
}
func PassTwo(op_array *types.ZendOpArray) int {
	var opline *ZendOp
	var end *ZendOp
	if !(ZEND_USER_CODE(op_array.GetType())) {
		return 0
	}
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_EXTENDED_STMT) != 0 {
		ZendUpdateExtendedStmts(op_array)
	}
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_HANDLE_OP_ARRAY) != 0 {
		if (ZendExtensionFlags & ZEND_EXTENSIONS_HAVE_OP_ARRAY_HANDLER) != 0 {
			ZendExtensions.ApplyWithArgument(LlistApplyWithArgFuncT(ZendExtensionOpArrayHandler), op_array)
		}
	}
	if CG__().GetContext().GetVarsSize() != op_array.GetLastVar() {
		op_array.SetVars((**types.String)(Erealloc(op_array.GetVars(), b.SizeOf("zend_string *")*op_array.GetLastVar())))
		CG__().GetContext().SetVarsSize(op_array.GetLastVar())
	}
	op_array.SetOpcodes((*ZendOp)(Erealloc(op_array.GetOpcodes(), ZEND_MM_ALIGNED_SIZE_EX(b.SizeOf("zend_op")*op_array.GetLast(), 16)+b.SizeOf("zval")*op_array.GetLastLiteral())))
	if op_array.GetLiterals() != nil {
		memcpy((*byte)(op_array.GetOpcodes())+ZEND_MM_ALIGNED_SIZE_EX(b.SizeOf("zend_op")*op_array.GetLast(), 16), op_array.GetLiterals(), b.SizeOf("zval")*op_array.GetLastLiteral())
		Efree(op_array.GetLiterals())
		op_array.SetLiterals((*types.Zval)((*byte)(op_array.GetOpcodes()) + ZEND_MM_ALIGNED_SIZE_EX(b.SizeOf("zend_op")*op_array.GetLast(), 16)))
	}
	CG__().GetContext().SetOpcodesSize(op_array.GetLast())
	CG__().GetContext().SetLiteralsSize(op_array.GetLastLiteral())

	/* Needs to be set directly after the opcode/literal reallocation, to ensure destruction
	 * happens correctly if any of the following fixups generate a fatal error. */

	op_array.SetIsDonePassTwo(true)
	opline = op_array.GetOpcodes()
	end = opline + op_array.GetLast()
	for opline < end {
		switch opline.GetOpcode() {
		case ZEND_RECV_INIT:
			var val *types.Zval = CT_CONSTANT(opline.GetOp2())
			if val.IsConstantAst() {
				var slot uint32 = ZEND_MM_ALIGNED_SIZE_EX(op_array.GetCacheSize(), 8)
				val.SetCacheSlot(slot)
				op_array.SetCacheSize(op_array.GetCacheSize() + b.SizeOf("zval"))
			}
		case ZEND_FAST_CALL:
			opline.GetOp1().SetOplineNum(op_array.GetTryCatchArray()[opline.GetOp1().GetNum()].GetFinallyOp())
			ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp1())
		case ZEND_BRK:
			fallthrough
		case ZEND_CONT:
			var jmp_target uint32 = ZendGetBrkContTarget(op_array, opline)
			if op_array.IsHasFinallyBlock() {
				ZendCheckFinallyBreakout(op_array, opline-op_array.GetOpcodes(), jmp_target)
			}
			opline.SetOpcode(ZEND_JMP)
			opline.GetOp1().SetOplineNum(jmp_target)
			opline.GetOp2().SetNum(0)
			ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp1())
		case ZEND_GOTO:
			ZendResolveGotoLabel(op_array, opline)
			if op_array.IsHasFinallyBlock() {
				ZendCheckFinallyBreakout(op_array, opline-op_array.GetOpcodes(), opline.GetOp1().GetOplineNum())
			}
			fallthrough
		case ZEND_JMP:
			ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp1())
		case ZEND_JMPZNZ:

			/* absolute index to relative offset */

			opline.SetExtendedValue(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, opline.GetExtendedValue()))
			fallthrough
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
			fallthrough
		case ZEND_FE_RESET_R:
			fallthrough
		case ZEND_FE_RESET_RW:
			ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp2())
		case ZEND_ASSERT_CHECK:

			/* If result of assert is unused, result of check is unused as well */

			var call *ZendOp = op_array.GetOpcodes()[opline.GetOp2().GetOplineNum()-1]
			if call.GetOpcode() == ZEND_EXT_FCALL_END {
				call--
			}
			if call.GetResultType() == IS_UNUSED {
				opline.SetResultType(IS_UNUSED)
			}
			ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp2())
		case ZEND_FE_FETCH_R:
			fallthrough
		case ZEND_FE_FETCH_RW:

			/* absolute index to relative offset */

			opline.SetExtendedValue(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, opline.GetExtendedValue()))
		case ZEND_CATCH:
			if (opline.GetExtendedValue() & ZEND_LAST_CATCH) == 0 {
				ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array, opline, opline.GetOp2())
			}
		case ZEND_RETURN:
			fallthrough
		case ZEND_RETURN_BY_REF:
			if op_array.IsGenerator() {
				opline.SetOpcode(ZEND_GENERATOR_RETURN)
			}
		case ZEND_SWITCH_LONG:
			fallthrough
		case ZEND_SWITCH_STRING:

			/* absolute indexes to relative offsets */

			var jumptable *types.Array = CT_CONSTANT(opline.GetOp2()).Array()
			jumptable.Foreach(func(_ types.ArrayKey, zv *types.Zval) {
				zv.SetLong(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, zv.Long()))
			})
			opline.SetExtendedValue(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, opline.GetExtendedValue()))
		}
		if opline.GetOp1Type() == IS_CONST {
			ZEND_PASS_TWO_UPDATE_CONSTANT(op_array, opline, opline.GetOp1())
		} else if (opline.GetOp1Type() & (IS_VAR | IS_TMP_VAR)) != 0 {
			opline.GetOp1().SetVar(uint32(types.ZendIntptrT(nil.VarNum(op_array.GetLastVar() + opline.GetOp1().GetVar()))))
		}
		if opline.GetOp2Type() == IS_CONST {
			ZEND_PASS_TWO_UPDATE_CONSTANT(op_array, opline, opline.GetOp2())
		} else if (opline.GetOp2Type() & (IS_VAR | IS_TMP_VAR)) != 0 {
			opline.GetOp2().SetVar(uint32(types.ZendIntptrT(nil.VarNum(op_array.GetLastVar() + opline.GetOp2().GetVar()))))
		}
		if (opline.GetResultType() & (IS_VAR | IS_TMP_VAR)) != 0 {
			opline.GetResult().SetVar(uint32(types.ZendIntptrT(nil.VarNum(op_array.GetLastVar() + opline.GetResult().GetVar()))))
		}
		ZendVmSetOpcodeHandler(opline)
		opline++
	}
	ZendCalcLiveRanges(op_array, nil)
	return 0
}
func GetUnaryOp(opcode int) UnaryOpType {
	switch opcode {
	case ZEND_BW_NOT:
		return UnaryOpType(operators.BitwiseNotFunction)
	case ZEND_BOOL_NOT:
		return UnaryOpType(operators.BooleanNotFunction)
	default:
		return UnaryOpType(nil)
	}
}
func GetBinaryOp(opcode int) BinaryOpType {
	switch opcode {
	case ZEND_ADD:
		return BinaryOpType(operators.AddFunction)
	case ZEND_SUB:
		return BinaryOpType(operators.SubFunction)
	case ZEND_MUL:
		return BinaryOpType(operators.MulFunction)
	case ZEND_POW:
		return BinaryOpType(operators.PowFunction)
	case ZEND_DIV:
		return BinaryOpType(operators.DivFunction)
	case ZEND_MOD:
		return BinaryOpType(operators.ModFunction)
	case ZEND_SL:
		return BinaryOpType(operators.ShiftLeftFunction)
	case ZEND_SR:
		return BinaryOpType(operators.ShiftRightFunction)
	case ZEND_PARENTHESIZED_CONCAT:
		fallthrough
	case ZEND_FAST_CONCAT:
		fallthrough
	case ZEND_CONCAT:
		return BinaryOpType(operators.ConcatFunction)
	case ZEND_IS_IDENTICAL:
		return BinaryOpType(operators.IsIdenticalFunction)
	case ZEND_IS_NOT_IDENTICAL:
		return BinaryOpType(operators.IsNotIdenticalFunction)
	case ZEND_IS_EQUAL:
		fallthrough
	case ZEND_CASE:
		return BinaryOpType(operators.IsEqualFunction)
	case ZEND_IS_NOT_EQUAL:
		return BinaryOpType(operators.IsNotEqualFunction)
	case ZEND_IS_SMALLER:
		return BinaryOpType(operators.IsSmallerFunction)
	case ZEND_IS_SMALLER_OR_EQUAL:
		return BinaryOpType(operators.IsSmallerOrEqualFunction)
	case ZEND_SPACESHIP:
		return BinaryOpType(operators.CompareFunction)
	case ZEND_BW_OR:
		return BinaryOpType(operators.BitwiseOrFunction)
	case ZEND_BW_AND:
		return BinaryOpType(operators.BitwiseAndFunction)
	case ZEND_BW_XOR:
		return BinaryOpType(operators.BitwiseXorFunction)
	case ZEND_BOOL_XOR:
		return BinaryOpType(operators.BooleanXorFunction)
	default:
		b.Assert(false)
		return BinaryOpType(nil)
	}
}
