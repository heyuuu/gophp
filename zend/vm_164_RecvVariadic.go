package zend

func ZEND_RECV_VARIADIC_SPEC_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var arg_num uint32 = opline.GetOp1().GetNum()
	var arg_count uint32 = executeData.NumArgs()
	var params *types.Zval
	params = opline.Result()
	if arg_num <= arg_count {
		var param *types.Zval
		ArrayInitSize(params, arg_count-arg_num+1)
		types.ZendHashRealInitPacked(params.GetArr())
		for {
			fillScope := types.PackedFillStart(params.GetArr())
			param = executeData.VarNum(executeData.GetFunc().GetOpArray().last_var + executeData.GetFunc().GetOpArray().T)
			if (executeData.GetFunc().GetOpArray().fn_flags & AccHasTypeHints) != 0 {
				ZEND_ADD_CALL_FLAG(executeData, ZEND_CALL_FREE_EXTRA_ARGS)
				for {
					ZendVerifyVariadicArgType(executeData.GetFunc(), arg_num, param, nil, CACHE_ADDR(opline.GetOp2().GetNum()))

					param.TryAddRefcount()

					fillScope.FillSet(param)
					fillScope.FillNext()
					param++
					if b.PreInc(&arg_num) > arg_count {
						break
					}
				}
			} else {
				for {

					param.TryAddRefcount()

					fillScope.FillSet(param)
					fillScope.FillNext()
					param++
					if b.PreInc(&arg_num) > arg_count {
						break
					}
				}
			}
			fillScope.FillEnd()
			break
		}
	} else {
		params.SetEmptyArray()
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
