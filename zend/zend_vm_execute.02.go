package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func zend_cannot_pass_by_ref_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg *types.Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	faults.ThrowError(nil, "Cannot pass parameter %d by reference", arg_num)
	FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
	arg = ZEND_CALL_VAR(executeData.GetCall(), opline.GetResult().GetVar())
	arg.SetUndef()
	return 0
}
func ZEND_SEND_UNPACK_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var args *types.Zval
	var arg_num int
	args = GetZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
	arg_num = executeData.GetCall().NumArgs() + 1
send_again:
	if args.IsArray() {
		var ht *types.Array = args.GetArr()
		var arg *types.Zval
		var top *types.Zval
		var name *types.String
		ZendVmStackExtendCallFrame(&(executeData.GetCall()), arg_num-1, ht.Len())
		if (opline.GetOp1Type()&(IS_VAR|IS_CV)) != 0 && args.GetRefcount() > 1 {
			var i uint32
			var separate int = 0

			/* check if any of arguments are going to be passed by reference */

			for i = 0; i < ht.Len(); i++ {
				if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num+i) != 0 {
					separate = 1
					break
				}
			}
			if separate != 0 {
				types.SEPARATE_ARRAY(args)
				ht = args.GetArr()
			}
		}
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			name = _p.GetKey()
			arg = _z
			if name != nil {
				faults.ThrowError(nil, "Cannot unpack array with string keys")
				FREE_OP(free_op1)
				return 0
			}
			top = executeData.GetCall().Arg(arg_num)
			if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
				if arg.IsReference() {
					arg.AddRefcount()
					top.SetReference(arg.GetRef())
				} else if (opline.GetOp1Type() & (IS_VAR | IS_CV)) != 0 {

					/* array is already separated above */

					types.ZVAL_MAKE_REF_EX(arg, 2)
					top.SetReference(arg.GetRef())
				} else {
					arg.TryAddRefcount()
					top.SetNewRef(arg)
				}
			} else {
				types.ZVAL_COPY_DEREF(top, arg)
			}
			executeData.GetCall().
				NumArgs()++
			arg_num++
		}
	} else if args.IsObject() {
		var ce *types.ClassEntry = types.Z_OBJCE_P(args)
		var iter *ZendObjectIterator
		if ce == nil || ce.GetGetIterator() == nil {
			faults.Error(faults.E_WARNING, "Only arrays and Traversables can be unpacked")
		} else {
			iter = ce.GetGetIterator()(ce, args, 0)
			if iter == nil {
				FREE_OP(free_op1)
				if EG__().GetException() == nil {
					faults.ThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
				}
				return 0
			}
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
			}
			for ; iter.GetFuncs().GetValid()(iter) == types.SUCCESS; arg_num++ {
				var arg *types.Zval
				var top *types.Zval
				if EG__().GetException() != nil {
					break
				}
				arg = iter.GetFuncs().GetGetCurrentData()(iter)
				if EG__().GetException() != nil {
					break
				}
				if iter.GetFuncs().GetGetCurrentKey() != nil {
					var key types.Zval
					iter.GetFuncs().GetGetCurrentKey()(iter, &key)
					if EG__().GetException() != nil {
						break
					}
					if key.GetType() != types.IS_LONG {
						faults.ThrowError(nil, b.Cond(key.IsString(), "Cannot unpack Traversable with string keys", "Cannot unpack Traversable with non-integer keys"))
						ZvalPtrDtor(&key)
						break
					}
				}
				if ARG_MUST_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
					faults.Error(faults.E_WARNING, "Cannot pass by-reference argument %d of %s%s%s()"+" by unpacking a Traversable, passing by-value instead", arg_num, b.CondF1(executeData.GetCall().func_.common.scope, func() []byte { return executeData.GetCall().func_.common.scope.name.GetVal() }, ""), b.Cond(executeData.GetCall().func_.common.scope, "::", ""), executeData.GetCall().func_.common.function_name.GetVal())
				}
				arg = types.ZVAL_DEREF(arg)
				arg.TryAddRefcount()
				ZendVmStackExtendCallFrame(&(executeData.GetCall()), arg_num-1, 1)
				top = executeData.GetCall().Arg(arg_num)
				types.ZVAL_COPY_VALUE(top, arg)
				executeData.GetCall().
					NumArgs()++
				iter.GetFuncs().GetMoveForward()(iter)
			}
			ZendIteratorDtor(iter)
		}
	} else if args.IsReference() {
		args = types.Z_REFVAL_P(args)
		goto send_again
	} else {
		if opline.GetOp1Type() == IS_CV && args.IsUndef() {
			ZVAL_UNDEFINED_OP1()
		}
		faults.Error(faults.E_WARNING, "Only arrays and Traversables can be unpacked")
	}
	FREE_OP(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_SEND_ARRAY_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var args *types.Zval
	args = GetZvalPtr(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
	if args.GetType() != types.IS_ARRAY {
		if (opline.GetOp1Type()&(IS_VAR|IS_CV)) != 0 && args.IsReference() {
			args = types.Z_REFVAL_P(args)
			if args.IsArray() {
				goto send_array
			}
		}
		faults.InternalTypeError(executeData.IsCallUseStrictTypes(), "call_user_func_array() expects parameter 2 to be array, %s given", types.ZendGetTypeByConst(args.GetType()))
		if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_CLOSURE) != 0 {
			OBJ_RELEASE(ZEND_CLOSURE_OBJECT(executeData.GetCall().func_))
		} else if (ZEND_CALL_INFO(executeData.GetCall()) & ZEND_CALL_RELEASE_THIS) != 0 {
			OBJ_RELEASE(executeData.GetCall().This.GetObj())
		}
		executeData.GetCall().
			func_ = (types.IFunction)(&ZendPassFunction)
		executeData.GetCall().
			This.GetObj() = nil
		ZEND_CALL_INFO(executeData.GetCall()) &= ^(ZEND_CALL_RELEASE_THIS | ZEND_CALL_HAS_THIS)
		FREE_UNFETCHED_OP(opline.GetOp2Type(), opline.GetOp2().GetVar())
	} else {
		var arg_num uint32
		var ht *types.Array
		var arg *types.Zval
		var param *types.Zval
	send_array:
		ht = args.GetArr()
		if opline.GetOp2Type() != IS_UNUSED {
			var free_op2 ZendFreeOp
			var op2 *types.Zval = GetZvalPtr(opline.GetOp2Type(), opline.GetOp2(), &free_op2, BP_VAR_R)
			var skip uint32 = opline.GetExtendedValue()
			var count uint32 = ht.Len()
			var len_ ZendLong = ZvalGetLong(op2)
			if len_ < 0 {
				len_ += zend_long(count - skip)
			}
			if skip < count && len_ > 0 {
				if len_ > zend_long(count-skip) {
					len_ = zend_long(count - skip)
				}
				ZendVmStackExtendCallFrame(&(executeData.GetCall()), 0, len_)
				arg_num = 1
				param = executeData.GetCall().Arg(1)
				var __ht *types.Array = ht
				for _, _p := range __ht.ForeachData() {
					var _z *types.Zval = _p.GetVal()

					arg = _z
					var must_wrap types.ZendBool = 0
					if skip > 0 {
						skip--
						continue
					} else if zend_long(arg_num-1) >= len_ {
						break
					} else if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
						if !(arg.IsReference()) {
							if ARG_MAY_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) == 0 {

								/* By-value send is not allowed -- emit a warning,
								 * but still perform the call. */

								ZendParamMustBeRef(executeData.GetCall().func_, arg_num)
								must_wrap = 1
							}
						}
					} else {
						if arg.IsReference() && (executeData.GetCall().func_.common.fn_flags&AccCallViaTrampoline) == 0 {

							/* don't separate references for __call */

							arg = types.Z_REFVAL_P(arg)

							/* don't separate references for __call */

						}
					}
					if must_wrap == 0 {
						types.ZVAL_COPY(param, arg)
					} else {
						arg.TryAddRefcount()
						param.SetNewRef(arg)
					}
					executeData.GetCall().
						NumArgs()++
					arg_num++
					param++
				}
			}
			FREE_OP(free_op2)
		} else {
			ZendVmStackExtendCallFrame(&(executeData.GetCall()), 0, ht.Len())
			arg_num = 1
			param = executeData.GetCall().Arg(1)
			var __ht *types.Array = ht
			for _, _p := range __ht.ForeachData() {
				var _z *types.Zval = _p.GetVal()

				arg = _z
				var must_wrap types.ZendBool = 0
				if ARG_SHOULD_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) != 0 {
					if !(arg.IsReference()) {
						if ARG_MAY_BE_SENT_BY_REF(executeData.GetCall().func_, arg_num) == 0 {

							/* By-value send is not allowed -- emit a warning,
							 * but still perform the call. */

							ZendParamMustBeRef(executeData.GetCall().func_, arg_num)
							must_wrap = 1
						}
					}
				} else {
					if arg.IsReference() && (executeData.GetCall().func_.common.fn_flags&AccCallViaTrampoline) == 0 {

						/* don't separate references for __call */

						arg = types.Z_REFVAL_P(arg)

						/* don't separate references for __call */

					}
				}
				if must_wrap == 0 {
					types.ZVAL_COPY(param, arg)
				} else {
					arg.TryAddRefcount()
					param.SetNewRef(arg)
				}
				executeData.GetCall().
					NumArgs()++
				arg_num++
				param++
			}
		}
	}
	FREE_OP(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func zend_case_helper_SPEC(op_1 *types.Zval, op_2 *types.Zval, executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if op_1.IsUndef() {
		op_1 = ZVAL_UNDEFINED_OP1()
	}
	if op_2.IsUndef() {
		op_2 = ZVAL_UNDEFINED_OP2()
	}
	CompareFunction(opline.GetResultZval(), op_1, op_2)
	if (opline.GetOp2Type() & (IS_TMP_VAR | IS_VAR)) != 0 {
		ZvalPtrDtorNogc(op_2)
	}
	if EG__().GetException() != nil {
		return 0
	}
	if opline.GetResultZval().GetLval() == 0 {
		ZEND_VM_SMART_BRANCH_TRUE()
		opline.GetResultZval().SetTrue()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	} else {
		ZEND_VM_SMART_BRANCH_FALSE()
		opline.GetResultZval().SetFalse()
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func ZEND_ADD_ARRAY_UNPACK_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	op1 = GetZvalPtr(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
add_unpack_again:
	if op1.IsArray() {
		var ht *types.Array = op1.GetArr()
		var val *types.Zval
		var key *types.String
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			val = _z
			if key != nil {
				faults.ThrowError(nil, "Cannot unpack array with string keys")
				FREE_OP(free_op1)
				return 0
			} else {
				if val.IsReference() && val.GetRefcount() == 1 {
					val = types.Z_REFVAL_P(val)
				}
				val.TryAddRefcount()
				if opline.GetResultZval().GetArr().NextIndexInsert(val) == nil {
					ZendCannotAddElement()
					ZvalPtrDtorNogc(val)
					break
				}
			}
		}
	} else if op1.IsObject() {
		var ce *types.ClassEntry = types.Z_OBJCE_P(op1)
		var iter *ZendObjectIterator
		if ce == nil || ce.GetGetIterator() == nil {
			faults.ThrowError(nil, "Only arrays and Traversables can be unpacked")
		} else {
			iter = ce.GetGetIterator()(ce, op1, 0)
			if iter == nil {
				FREE_OP(free_op1)
				if EG__().GetException() == nil {
					faults.ThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
				}
				return 0
			}
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
			}
			for iter.GetFuncs().GetValid()(iter) == types.SUCCESS {
				var val *types.Zval
				if EG__().GetException() != nil {
					break
				}
				val = iter.GetFuncs().GetGetCurrentData()(iter)
				if EG__().GetException() != nil {
					break
				}
				if iter.GetFuncs().GetGetCurrentKey() != nil {
					var key types.Zval
					iter.GetFuncs().GetGetCurrentKey()(iter, &key)
					if EG__().GetException() != nil {
						break
					}
					if key.GetType() != types.IS_LONG {
						faults.ThrowError(nil, b.Cond(key.IsString(), "Cannot unpack Traversable with string keys", "Cannot unpack Traversable with non-integer keys"))
						ZvalPtrDtor(&key)
						break
					}
				}
				val = types.ZVAL_DEREF(val)
				val.TryAddRefcount()
				if opline.GetResultZval().GetArr().NextIndexInsert(val) == nil {
					ZendCannotAddElement()
					ZvalPtrDtorNogc(val)
				}
				iter.GetFuncs().GetMoveForward()(iter)
			}
			ZendIteratorDtor(iter)
		}
	} else if op1.IsReference() {
		op1 = types.Z_REFVAL_P(op1)
		goto add_unpack_again
	} else {
		faults.ThrowError(nil, "Only arrays and Traversables can be unpacked")
	}
	FREE_OP(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_UNSET_STATIC_PROP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.Zval
	var name *types.String
	var tmp_name *types.String = nil
	var ce *types.ClassEntry
	var free_op1 ZendFreeOp
	if opline.GetOp2Type() == IS_CONST {
		ce = CACHED_PTR(opline.GetExtendedValue())
		if ce == nil {
			ce = ZendFetchClassByName(RT_CONSTANT(opline, opline.GetOp2()).GetStr(), (RT_CONSTANT(opline, opline.GetOp2()) + 1).GetStr(), ZEND_FETCH_CLASS_DEFAULT|ZEND_FETCH_CLASS_EXCEPTION)
			if ce == nil {
				b.Assert(EG__().GetException() != nil)
				FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
				return 0
			}
		}
	} else if opline.GetOp2Type() == IS_UNUSED {
		ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
		if ce == nil {
			b.Assert(EG__().GetException() != nil)
			FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
			return 0
		}
	} else {
		ce = opline.GetOp2Zval().GetCe()
	}
	varname = GetZvalPtrUndef(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
	if opline.GetOp1Type() == IS_CONST {
		name = varname.GetStr()
	} else if varname.IsString() {
		name = varname.GetStr()
	} else {
		if opline.GetOp1Type() == IS_CV && varname.IsUndef() {
			varname = ZVAL_UNDEFINED_OP1()
		}
		name = ZvalGetTmpString(varname, &tmp_name)
	}
	ZendStdUnsetStaticProperty(ce, name)
	ZendTmpStringRelease(tmp_name)
	FREE_OP(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_ISSET_ISEMPTY_STATIC_PROP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var value *types.Zval
	var result int
	result = ZendFetchStaticPropertyAddress(&value, nil, opline.GetExtendedValue() & ^ZEND_ISEMPTY, BP_VAR_IS, 0, opline, executeData)
	if (opline.GetExtendedValue() & ZEND_ISEMPTY) == 0 {
		result = result == types.SUCCESS && value.GetType() > types.IS_NULL && (!(value.IsReference()) || types.Z_REFVAL_P(value).GetType() != types.IS_NULL)
	} else {
		result = result != types.SUCCESS || IZendIsTrue(value) == 0
	}
	ZEND_VM_SMART_BRANCH(result, 1)
	types.ZVAL_BOOL(opline.GetResultZval(), result != 0)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_EXIT_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if opline.GetOp1Type() != IS_UNUSED {
		var free_op1 ZendFreeOp
		var ptr *types.Zval = GetZvalPtr(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
		for {
			if ptr.IsLong() {
				EG__().SetExitStatus(ptr.GetLval())
			} else {
				if (opline.GetOp1Type()&(IS_VAR|IS_CV)) != 0 && ptr.IsReference() {
					ptr = types.Z_REFVAL_P(ptr)
					if ptr.IsLong() {
						EG__().SetExitStatus(ptr.GetLval())
						break
					}
				}
				ZendPrintZval(ptr, 0)
			}
			break
		}
		FREE_OP(free_op1)
	}
	faults.Bailout()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_BEGIN_SILENCE_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	opline.GetResultZval().SetLong(EG__().GetErrorReporting())
	if EG__().GetErrorReporting() != 0 {
		for {
			EG__().SetErrorReporting(0)
			if EG__().GetErrorReportingIniEntry() == nil {
				var zv *types.Zval = EG__().GetIniDirectives().KeyFind(types.ZSTR_ERROR_REPORTING.GetStr())
				if zv != nil {
					EG__().SetErrorReportingIniEntry((*ZendIniEntry)(zv.GetPtr()))
				} else {
					break
				}
			}
			if EG__().GetErrorReportingIniEntry().GetModified() == 0 {
				if EG__().GetModifiedIniDirectives() == nil {
					ALLOC_HASHTABLE(EG__().GetModifiedIniDirectives())
					EG__().GetModifiedIniDirectives() = types.MakeArrayEx(8, nil, 0)
				}
				if types.ZendHashAddPtr(EG__().GetModifiedIniDirectives(), types.ZSTR_ERROR_REPORTING.GetStr(), EG__().GetErrorReportingIniEntry()) != nil {
					EG__().GetErrorReportingIniEntry().SetOrigValue(EG__().GetErrorReportingIniEntry().GetValue())
					EG__().GetErrorReportingIniEntry().SetOrigModifiable(EG__().GetErrorReportingIniEntry().GetModifiable())
					EG__().GetErrorReportingIniEntry().SetModified(1)
				}
			}
			break
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_EXT_STMT_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if EG__().GetNoExtensions() == 0 {
		ZendExtensions.ApplyWithArgument(LlistApplyWithArgFuncT(ZendExtensionStatementHandler), executeData)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_EXT_FCALL_BEGIN_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if EG__().GetNoExtensions() == 0 {
		ZendExtensions.ApplyWithArgument(LlistApplyWithArgFuncT(ZendExtensionFcallBeginHandler), executeData)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_EXT_FCALL_END_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if EG__().GetNoExtensions() == 0 {
		ZendExtensions.ApplyWithArgument(LlistApplyWithArgFuncT(ZendExtensionFcallEndHandler), executeData)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_DECLARE_ANON_CLASS_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var zv *types.Zval
	var ce *types.ClassEntry
	var opline *ZendOp = executeData.GetOpline()
	ce = CACHED_PTR(opline.GetExtendedValue())
	if ce == nil {
		var rtd_key *types.String = RT_CONSTANT(opline, opline.GetOp1()).GetStr()
		zv = EG__().GetClassTable().KeyFind(rtd_key.GetStr())
		if zv == nil {
			for {
				b.Assert((executeData.GetFunc().op_array.fn_flags & AccPreloaded) != 0)
				if ZendPreloadAutoload != nil && ZendPreloadAutoload(executeData.GetFunc().op_array.filename) == types.SUCCESS {
					zv = EG__().GetClassTable().KeyFind(rtd_key.GetStr())
					if zv != nil {
						break
					}
				}
				faults.ErrorNoreturn(faults.E_ERROR, "Anonymous class wasn't preloaded")
				break
			}
		}
		b.Assert(zv != nil)
		ce = zv.GetCe()
		if !ce.IsLinked() {
			if ZendDoLinkClass(ce, b.CondF1(opline.GetOp2Type() == IS_CONST, func() *types.String { return RT_CONSTANT(opline, opline.GetOp2()).GetStr() }, nil)) == types.FAILURE {
				return 0
			}
		}
		CACHE_PTR(opline.GetExtendedValue(), ce)
	}
	opline.GetResultZval().SetCe(ce)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_DECLARE_FUNCTION_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	DoBindFunction(RT_CONSTANT(opline, opline.GetOp1()))
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_TICKS_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if uint32(b.PreInc(&(EG__().GetTicksCount())) >= opline.GetExtendedValue()) != 0 {
		EG__().SetTicksCount(0)
		if ZendTicksFunction != nil {
			ZendTicksFunction(opline.GetExtendedValue())
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_EXT_NOP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_NOP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func zend_dispatch_try_catch_finally_helper_SPEC(try_catch_offset uint32, op_num uint32, executeData *ZendExecuteData) int {
	/* May be NULL during generator closing (only finally blocks are executed) */

	var ex *types.ZendObject = EG__().GetException()

	/* Walk try/catch/finally structures upwards, performing the necessary actions */

	for try_catch_offset != uint32-1 {
		var try_catch *ZendTryCatchElement = executeData.GetFunc().op_array.try_catch_array[try_catch_offset]
		if op_num < try_catch.GetCatchOp() && ex != nil {

			/* Go to catch block */

			CleanupLiveVars(executeData, op_num, try_catch.GetCatchOp())
			return ZEND_VM_JMP_EX(executeData, executeData.GetFunc().op_array.opcodes[try_catch.GetCatchOp()], 0)
		} else if op_num < try_catch.GetFinallyOp() {

			/* Go to finally block */

			var fast_call *types.Zval = EX_VAR(executeData.GetFunc().op_array.opcodes[try_catch.GetFinallyEnd()].op1.var_)
			CleanupLiveVars(executeData, op_num, try_catch.GetFinallyOp())
			fast_call.SetObj(EG__().GetException())
			EG__().SetException(nil)
			fast_call.SetOplineNum(uint32 - 1)
			return ZEND_VM_JMP_EX(executeData, executeData.GetFunc().op_array.opcodes[try_catch.GetFinallyOp()], 0)
		} else if op_num < try_catch.GetFinallyEnd() {
			var fast_call *types.Zval = EX_VAR(executeData.GetFunc().op_array.opcodes[try_catch.GetFinallyEnd()].op1.var_)

			/* cleanup incomplete RETURN statement */

			if fast_call.GetOplineNum() != uint32-1 && (executeData.GetFunc().op_array.opcodes[fast_call.GetOplineNum()].op2_type&(IS_TMP_VAR|IS_VAR)) != 0 {
				var return_value *types.Zval = EX_VAR(executeData.GetFunc().op_array.opcodes[fast_call.GetOplineNum()].op2.var_)
				ZvalPtrDtor(return_value)
			}

			/* Chain potential exception from wrapping finally block */

			if fast_call.GetObj() != nil {
				if ex != nil {
					faults.ExceptionSetPrevious(ex, fast_call.GetObj())
				} else {
					EG__().SetException(fast_call.GetObj())
				}
				ex = fast_call.GetObj()
			}

			/* Chain potential exception from wrapping finally block */

		}
		try_catch_offset--
	}

	/* Uncaught exception */

	CleanupLiveVars(executeData, op_num, 0)
	if (EX_CALL_INFO() & ZEND_CALL_GENERATOR) != 0 {
		var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
		ZendGeneratorClose(generator, 1)
		return -1
	} else {

		/* We didn't execute RETURN, and have to initialize return_value */

		if executeData.GetReturnValue() {
			executeData.GetReturnValue().
				SetUndef()
		}
		return zend_leave_helper_SPEC(executeData)
	}
}
func ZEND_HANDLE_EXCEPTION_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var throw_op *ZendOp = EG__().GetOplineBeforeException()
	var throw_op_num uint32 = throw_op - executeData.GetFunc().op_array.opcodes
	var i int
	var current_try_catch_offset int = -1
	if (throw_op.GetOpcode() == ZEND_FREE || throw_op.GetOpcode() == ZEND_FE_FREE) && (throw_op.GetExtendedValue()&ZEND_FREE_ON_RETURN) != 0 {

		/* exceptions thrown because of loop var destruction on return/break/...
		 * are logically thrown at the end of the foreach loop, so adjust the
		 * throw_op_num.
		 */

		var range_ *ZendLiveRange = FindLiveRange(executeData.GetFunc().op_array, throw_op_num, throw_op.GetOp1().GetVar())
		throw_op_num = range_.GetEnd()
	}

	/* Find the innermost try/catch/finally the exception was thrown in */

	for i = 0; i < executeData.GetFunc().op_array.last_try_catch; i++ {
		var try_catch *ZendTryCatchElement = executeData.GetFunc().op_array.try_catch_array[i]
		if try_catch.GetTryOp() > throw_op_num {

			/* further blocks will not be relevant... */

			break

			/* further blocks will not be relevant... */

		}
		if throw_op_num < try_catch.GetCatchOp() || throw_op_num < try_catch.GetFinallyEnd() {
			current_try_catch_offset = i
		}
	}
	CleanupUnfinishedCalls(executeData, throw_op_num)
	if (throw_op.GetResultType() & (IS_VAR | IS_TMP_VAR)) != 0 {
		switch throw_op.GetOpcode() {
		case ZEND_ADD_ARRAY_ELEMENT:
			fallthrough
		case ZEND_ADD_ARRAY_UNPACK:
			fallthrough
		case ZEND_ROPE_INIT:
			fallthrough
		case ZEND_ROPE_ADD:

		case ZEND_FETCH_CLASS:
			fallthrough
		case ZEND_DECLARE_ANON_CLASS:

		default:
			ZvalPtrDtorNogc(EX_VAR(throw_op.GetResult().GetVar()))
		}
	}
	return zend_dispatch_try_catch_finally_helper_SPEC(current_try_catch_offset, throw_op_num, executeData)
}
func ZEND_USER_OPCODE_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ret int
	//ret = ZendUserOpcodeHandlers[opline.GetOpcode()](executeData)
	ret = UserOpcodeHandlerT(nil)(executeData)
	opline = executeData.GetOpline()
	switch ret {
	case ZEND_USER_OPCODE_CONTINUE:
		return 0
	case ZEND_USER_OPCODE_RETURN:
		if (EX_CALL_INFO() & ZEND_CALL_GENERATOR) != 0 {
			var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
			ZendGeneratorClose(generator, 1)
			return -1
		} else {
			return zend_leave_helper_SPEC(executeData)
		}
		fallthrough
	case ZEND_USER_OPCODE_ENTER:
		return 1
	case ZEND_USER_OPCODE_LEAVE:
		return 2
	case ZEND_USER_OPCODE_DISPATCH:
		ZEND_VM_DISPATCH(executeData, opline.GetOpcode(), opline)
		fallthrough
	default:
		ZEND_VM_DISPATCH(executeData, zend_uchar(ret&0xff), opline)
	}
}
func zend_yield_in_closed_generator_helper_SPEC(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	faults.ThrowError(nil, "Cannot yield from finally in a force-closed generator")
	FREE_UNFETCHED_OP(opline.GetOp2Type(), opline.GetOp2().GetVar())
	FREE_UNFETCHED_OP(opline.GetOp1Type(), opline.GetOp1().GetVar())
	UNDEF_RESULT()
	return 0
}
func ZEND_DISCARD_EXCEPTION_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fast_call *types.Zval = opline.GetOp1Zval()

	/* cleanup incomplete RETURN statement */

	if fast_call.GetOplineNum() != uint32-1 && (executeData.GetFunc().op_array.opcodes[fast_call.GetOplineNum()].op2_type&(IS_TMP_VAR|IS_VAR)) != 0 {
		var return_value *types.Zval = EX_VAR(executeData.GetFunc().op_array.opcodes[fast_call.GetOplineNum()].op2.var_)
		ZvalPtrDtor(return_value)
	}

	/* cleanup delayed exception */

	if fast_call.GetObj() != nil {

		/* discard the previously thrown exception */

		OBJ_RELEASE(fast_call.GetObj())
		fast_call.SetObj(nil)
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FAST_CALL_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fast_call *types.Zval = opline.GetResultZval()
	fast_call.SetObj(nil)

	/* set return address */

	fast_call.SetOplineNum(opline - executeData.GetFunc().op_array.opcodes)
	return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp1()), 0)
}
func ZEND_FAST_RET_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fast_call *types.Zval = opline.GetOp1Zval()
	var current_try_catch_offset uint32
	var current_op_num uint32
	if fast_call.GetOplineNum() != uint32-1 {
		var fast_ret *ZendOp = executeData.GetFunc().op_array.opcodes + fast_call.GetOplineNum()
		return ZEND_VM_JMP_EX(executeData, fast_ret+1, 0)
	}

	/* special case for unhandled exceptions */

	EG__().SetException(fast_call.GetObj())
	fast_call.SetObj(nil)
	current_try_catch_offset = opline.GetOp2().GetNum()
	current_op_num = opline - executeData.GetFunc().op_array.opcodes
	return zend_dispatch_try_catch_finally_helper_SPEC(current_try_catch_offset, current_op_num, executeData)
}
func ZEND_ASSERT_CHECK_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	if EG__().GetAssertions() <= 0 {
		var target *ZendOp = OP_JMP_ADDR(opline, opline.GetOp2())
		if RETURN_VALUE_USED(opline) {
			opline.GetResultZval().SetTrue()
		}
		return ZEND_VM_JMP_EX(executeData, target, 0)
	} else {
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func ZEND_CALL_TRAMPOLINE_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var args *types.Array = nil
	var fbc types.IFunction = executeData.GetFunc()
	var ret *types.Zval = executeData.GetReturnValue()
	var call_info uint32 = EX_CALL_INFO() & (ZEND_CALL_NESTED | ZEND_CALL_TOP | ZEND_CALL_RELEASE_THIS)
	var num_args uint32 = executeData.NumArgs()
	var call *ZendExecuteData
	if num_args != 0 {
		var p *types.Zval = executeData.Arg(1)
		var end *types.Zval = p + num_args
		args = types.NewArray(num_args)
		types.ZendHashRealInitPacked(args)
		for {
			fillScope := types.PackedFillStart(args)
			for {
				fillScope.FillSet(p)
				fillScope.FillNext()
				p++
				if p == end {
					break
				}
			}
			fillScope.FillEnd()
			break
		}
	}
	call = executeData
	EG__().SetCurrentExecuteData(executeData.GetPrevExecuteData())
	executeData = CurrEX()
	if fbc.GetOpArray().IsStatic() {
		call.SetFunc(fbc.GetOpArray().GetScope().GetCallstatic())
	} else {
		call.SetFunc(fbc.GetOpArray().GetScope().GetCall())
	}
	b.Assert(ZendVmCalcUsedStack(2, call.GetFunc()) <= size_t((*byte)(EG__().GetVmStackEnd())-(*byte)(call)))
	call.NumArgs() = 2
	call.Arg(1).SetString(fbc.GetFunctionName())
	if args != nil {
		call.Arg(2).SetArray(args)
	} else {
		call.Arg(2).SetEmptyArray()
	}
	ZendFreeTrampoline(fbc)
	fbc = call.GetFunc()
	if fbc.GetType() == ZEND_USER_FUNCTION {
		if !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		executeData = call
		IInitFuncExecuteData(fbc.GetOpArray(), ret, 0, executeData)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			executeData = executeData.GetPrevExecuteData()
			ZEND_ADD_CALL_FLAG(call, ZEND_CALL_TOP)
			ZendExecuteEx(call)
		}
	} else {
		var retval types.Zval
		b.Assert(fbc.GetType() == ZEND_INTERNAL_FUNCTION)
		EG__().SetCurrentExecuteData(call)
		if fbc.IsHasTypeHints() && ZendVerifyInternalArgTypes(fbc, call) == 0 {
			ZendVmStackFreeCallFrame(call)
			if ret != nil {
				ret.SetUndef()
			}
			goto call_trampoline_end
		}
		if ret == nil {
			ret = &retval
		}
		ret.SetNull()
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			fbc.GetInternalFunction().GetHandler()(call, ret)

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, ret)
		}
		EG__().SetCurrentExecuteData(call.GetPrevExecuteData())
	call_trampoline_end:
		ZendVmStackFreeArgs(call)
		if ret == &retval {
			ZvalPtrDtor(ret)
		}
	}
	executeData = CurrEX()
	if !(executeData.GetFunc()) || !(ZEND_USER_CODE(executeData.GetFunc().type_)) || (call_info&ZEND_CALL_TOP) != 0 {
		return -1
	}
	if (call_info & ZEND_CALL_RELEASE_THIS) != 0 {
		var object *types.ZendObject = call.GetThis().GetObj()
		OBJ_RELEASE(object)
	}
	ZendVmStackFreeCallFrame(call)
	if EG__().GetException() != nil {
		faults.RethrowException(executeData)
		return 2
	}
	ZEND_VM_INC_OPCODE(executeData)
	return 2
}
func zend_interrupt_helper_SPEC(executeData *ZendExecuteData) int {
	EG__().SetVmInterrupt(0)
	if EG__().GetTimedOut() != 0 {
		ZendTimeout(0)
	} else if ZendInterruptFunction != nil {
		ZendInterruptFunction(executeData)
		return 1
	}
	return 0
}
func ZEND_INIT_FCALL_BY_NAME_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var fbc types.IFunction
	var function_name *types.Zval
	var func_ *types.Zval
	var call *ZendExecuteData
	fbc = CACHED_PTR(opline.GetResult().GetNum())
	if fbc == nil {
		function_name = (*types.Zval)(RT_CONSTANT(opline, opline.GetOp2()))
		func_ = EG__().GetFunctionTable().KeyFind((function_name + 1).GetStr().GetStr())
		if func_ == nil {
			return zend_undefined_function_helper_SPEC(executeData)
		}
		fbc = func_.GetFunc()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		CACHE_PTR(opline.GetResult().GetNum(), fbc)
	}
	call = _zendVmStackPushCallFrame(ZEND_CALL_NESTED_FUNCTION, fbc, opline.GetExtendedValue(), nil)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
func ZEND_INIT_DYNAMIC_CALL_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var function_name *types.Zval
	var call *ZendExecuteData
	function_name = RT_CONSTANT(opline, opline.GetOp2())
try_function_name:

	if function_name.IsArray() {
		call = ZendInitDynamicCallArray(function_name.GetArr(), opline.GetExtendedValue())
	} else {
		faults.ThrowError(nil, "Function name must be a string")
		call = nil
	}
	if call == nil {
		return 0
	}
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_INIT_NS_FCALL_BY_NAME_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var func_name *types.Zval
	var func_ *types.Zval
	var fbc types.IFunction
	var call *ZendExecuteData
	fbc = CACHED_PTR(opline.GetResult().GetNum())
	if fbc == nil {
		func_name = (*types.Zval)(RT_CONSTANT(opline, opline.GetOp2()))
		func_ = EG__().GetFunctionTable().KeyFind((func_name + 1).GetStr().GetStr())
		if func_ == nil {
			func_ = EG__().GetFunctionTable().KeyFind((func_name + 2).GetStr().GetStr())
			if func_ == nil {
				return zend_undefined_function_helper_SPEC(executeData)
			}
		}
		fbc = func_.GetFunc()
		if fbc.GetType() == ZEND_USER_FUNCTION && !(RUN_TIME_CACHE(fbc.GetOpArray())) {
			InitFuncRunTimeCache(fbc.GetOpArray())
		}
		CACHE_PTR(opline.GetResult().GetNum(), fbc)
	}
	call = _zendVmStackPushCallFrame(ZEND_CALL_NESTED_FUNCTION, fbc, opline.GetExtendedValue(), nil)
	call.SetPrevExecuteData(executeData.GetCall())
	executeData.GetCall() = call
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
