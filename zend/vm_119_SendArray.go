package zend

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
			OBJ_RELEASE(executeData.GetCall().This.Object())
		}
		executeData.GetCall().
			func_ = (types.IFunction)(&ZendPassFunction)
		executeData.GetCall().
			This.Object() = nil
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
						// arg.TryAddRefcount()
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
					// arg.TryAddRefcount()
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
