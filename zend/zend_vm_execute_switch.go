package zend

func vmGetHandler(opcode OpCode, op *ZendOp) OpcodeHandlerT {
    var handler OpcodeHandlerT
    switch opcode {
    case ZEND_NOP:
        handler = ZEND_NOP_SPEC_HANDLER
    case ZEND_ADD:
        handler = vmGetAddHandler(op)
    case ZEND_SUB:
        handler = vmGetSubHandler(op)
    case ZEND_MUL:
        handler = vmGetMulHandler(op)
    case ZEND_DIV:
        handler = vmGetDivHandler(op)
    case ZEND_MOD:
        handler = vmGetModHandler(op)
    case ZEND_SL:
        handler = vmGetSlHandler(op)
    case ZEND_SR:
        handler = vmGetSrHandler(op)
    case ZEND_CONCAT:
        handler = vmGetConcatHandler(op)
    case ZEND_BW_OR:
        handler = vmGetBwOrHandler(op)
    case ZEND_BW_AND:
        handler = vmGetBwAndHandler(op)
    case ZEND_BW_XOR:
        handler = vmGetBwXorHandler(op)
    case ZEND_POW:
        handler = vmGetPowHandler(op)
    case ZEND_BW_NOT:
        handler = vmGetBwNotHandler(op)
    case ZEND_BOOL_NOT:
        handler = vmGetBoolNotHandler(op)
    case ZEND_BOOL_XOR:
        handler = vmGetBoolXorHandler(op)
    case ZEND_IS_IDENTICAL:
        handler = vmGetIsIdenticalHandler(op)
    case ZEND_IS_NOT_IDENTICAL:
        handler = vmGetIsNotIdenticalHandler(op)
    case ZEND_IS_EQUAL:
        handler = vmGetIsEqualHandler(op)
    case ZEND_IS_NOT_EQUAL:
        handler = vmGetIsNotEqualHandler(op)
    case ZEND_IS_SMALLER:
        handler = vmGetIsSmallerHandler(op)
    case ZEND_IS_SMALLER_OR_EQUAL:
        handler = vmGetIsSmallerOrEqualHandler(op)
    case ZEND_ASSIGN:
        handler = vmGetAssignHandler(op)
    case ZEND_ASSIGN_DIM:
        handler = vmGetAssignDimHandler(op)
    case ZEND_ASSIGN_OBJ:
        handler = vmGetAssignObjHandler(op)
    case ZEND_ASSIGN_STATIC_PROP:
        handler = vmGetAssignStaticPropHandler(op)
    case ZEND_ASSIGN_OP:
        handler = vmGetAssignOpHandler(op)
    case ZEND_ASSIGN_DIM_OP:
        handler = vmGetAssignDimOpHandler(op)
    case ZEND_ASSIGN_OBJ_OP:
        handler = vmGetAssignObjOpHandler(op)
    case ZEND_ASSIGN_STATIC_PROP_OP:
        handler = ZEND_ASSIGN_STATIC_PROP_OP_SPEC_HANDLER
    case ZEND_ASSIGN_REF:
        handler = vmGetAssignRefHandler(op)
    case ZEND_QM_ASSIGN:
        handler = vmGetQmAssignHandler(op)
    case ZEND_ASSIGN_OBJ_REF:
        handler = vmGetAssignObjRefHandler(op)
    case ZEND_ASSIGN_STATIC_PROP_REF:
        handler = ZEND_ASSIGN_STATIC_PROP_REF_SPEC_HANDLER
    case ZEND_PRE_INC:
        handler = vmGetPreIncHandler(op)
    case ZEND_PRE_DEC:
        handler = vmGetPreDecHandler(op)
    case ZEND_POST_INC:
        handler = vmGetPostIncHandler(op)
    case ZEND_POST_DEC:
        handler = vmGetPostDecHandler(op)
    case ZEND_PRE_INC_STATIC_PROP:
        handler = vmGetPreIncStaticPropHandler(op)
    case ZEND_PRE_DEC_STATIC_PROP:
        handler = ZEND_PRE_INC_STATIC_PROP_SPEC_HANDLER
    case ZEND_POST_INC_STATIC_PROP:
        handler = vmGetPostIncStaticPropHandler(op)
    case ZEND_POST_DEC_STATIC_PROP:
        handler = ZEND_POST_INC_STATIC_PROP_SPEC_HANDLER
    case ZEND_JMP:
        handler = ZEND_JMP_SPEC_HANDLER
    case ZEND_JMPZ:
        handler = vmGetJmpzHandler(op)
    case ZEND_JMPNZ:
        handler = vmGetJmpnzHandler(op)
    case ZEND_JMPZNZ:
        handler = vmGetJmpznzHandler(op)
    case ZEND_JMPZ_EX:
        handler = vmGetJmpzExHandler(op)
    case ZEND_JMPNZ_EX:
        handler = vmGetJmpnzExHandler(op)
    case ZEND_CASE:
        handler = vmGetCaseHandler(op)
    case ZEND_CHECK_VAR:
        handler = ZEND_CHECK_VAR_SPEC_CV_UNUSED_HANDLER
    case ZEND_SEND_VAR_NO_REF_EX:
        handler = vmGetSendVarNoRefExHandler(op)
    case ZEND_CAST:
        handler = vmGetCastHandler(op)
    case ZEND_BOOL:
        handler = vmGetBoolHandler(op)
    case ZEND_FAST_CONCAT:
        handler = vmGetFastConcatHandler(op)
    case ZEND_ROPE_INIT:
        handler = vmGetRopeInitHandler(op)
    case ZEND_ROPE_ADD:
        handler = vmGetRopeAddHandler(op)
    case ZEND_ROPE_END:
        handler = vmGetRopeEndHandler(op)
    case ZEND_BEGIN_SILENCE:
        handler = ZEND_BEGIN_SILENCE_SPEC_HANDLER
    case ZEND_END_SILENCE:
        handler = ZEND_END_SILENCE_SPEC_TMP_HANDLER
    case ZEND_INIT_FCALL_BY_NAME:
        handler = ZEND_INIT_FCALL_BY_NAME_SPEC_CONST_HANDLER
    case ZEND_DO_FCALL:
        handler = vmGetDoFcallHandler(op)
    case ZEND_INIT_FCALL:
        handler = ZEND_INIT_FCALL_SPEC_CONST_HANDLER
    case ZEND_RETURN:
        handler = vmGetReturnHandler(op)
    case ZEND_RECV:
        handler = ZEND_RECV_SPEC_UNUSED_HANDLER
    case ZEND_RECV_INIT:
        handler = ZEND_RECV_INIT_SPEC_CONST_HANDLER
    case ZEND_SEND_VAL:
        handler = vmGetSendValHandler(op)
    case ZEND_SEND_VAR_EX:
        handler = vmGetSendVarExHandler(op)
    case ZEND_SEND_REF:
        handler = vmGetSendRefHandler(op)
    case ZEND_NEW:
        handler = vmGetNewHandler(op)
    case ZEND_INIT_NS_FCALL_BY_NAME:
        handler = ZEND_INIT_NS_FCALL_BY_NAME_SPEC_CONST_HANDLER
    case ZEND_FREE:
        handler = ZEND_FREE_SPEC_TMPVAR_HANDLER
    case ZEND_INIT_ARRAY:
        handler = vmGetInitArrayHandler(op)
    case ZEND_ADD_ARRAY_ELEMENT:
        handler = vmGetAddArrayElementHandler(op)
    case ZEND_INCLUDE_OR_EVAL:
        handler = vmGetIncludeOrEvalHandler(op)
    case ZEND_UNSET_VAR:
        handler = vmGetUnsetVarHandler(op)
    case ZEND_UNSET_DIM:
        handler = vmGetUnsetDimHandler(op)
    case ZEND_UNSET_OBJ:
        handler = vmGetUnsetObjHandler(op)
    case ZEND_FE_RESET_R:
        handler = vmGetFeResetRHandler(op)
    case ZEND_FE_FETCH_R:
        handler = ZEND_FE_FETCH_R_SPEC_VAR_HANDLER
    case ZEND_EXIT:
        handler = ZEND_EXIT_SPEC_HANDLER
    case ZEND_FETCH_R:
        handler = vmGetFetchRHandler(op)
    case ZEND_FETCH_DIM_R:
        handler = vmGetFetchDimRHandler(op)
    case ZEND_FETCH_OBJ_R:
        handler = vmGetFetchObjRHandler(op)
    case ZEND_FETCH_W:
        handler = vmGetFetchWHandler(op)
    case ZEND_FETCH_DIM_W:
        handler = vmGetFetchDimWHandler(op)
    case ZEND_FETCH_OBJ_W:
        handler = vmGetFetchObjWHandler(op)
    case ZEND_FETCH_RW:
        handler = vmGetFetchRwHandler(op)
    case ZEND_FETCH_DIM_RW:
        handler = vmGetFetchDimRwHandler(op)
    case ZEND_FETCH_OBJ_RW:
        handler = vmGetFetchObjRwHandler(op)
    case ZEND_FETCH_IS:
        handler = vmGetFetchIsHandler(op)
    case ZEND_FETCH_DIM_IS:
        handler = vmGetFetchDimIsHandler(op)
    case ZEND_FETCH_OBJ_IS:
        handler = vmGetFetchObjIsHandler(op)
    case ZEND_FETCH_FUNC_ARG:
        handler = vmGetFetchFuncArgHandler(op)
    case ZEND_FETCH_DIM_FUNC_ARG:
        handler = vmGetFetchDimFuncArgHandler(op)
    case ZEND_FETCH_OBJ_FUNC_ARG:
        handler = vmGetFetchObjFuncArgHandler(op)
    case ZEND_FETCH_UNSET:
        handler = vmGetFetchUnsetHandler(op)
    case ZEND_FETCH_DIM_UNSET:
        handler = vmGetFetchDimUnsetHandler(op)
    case ZEND_FETCH_OBJ_UNSET:
        handler = vmGetFetchObjUnsetHandler(op)
    case ZEND_FETCH_LIST_R:
        handler = vmGetFetchListRHandler(op)
    case ZEND_FETCH_CONSTANT:
        handler = ZEND_FETCH_CONSTANT_SPEC_UNUSED_CONST_HANDLER
    case ZEND_CHECK_FUNC_ARG:
        handler = vmGetCheckFuncArgHandler(op)
    case ZEND_EXT_STMT:
        handler = ZEND_EXT_STMT_SPEC_HANDLER
    case ZEND_EXT_FCALL_BEGIN:
        handler = ZEND_EXT_FCALL_BEGIN_SPEC_HANDLER
    case ZEND_EXT_FCALL_END:
        handler = ZEND_EXT_FCALL_END_SPEC_HANDLER
    case ZEND_EXT_NOP:
        handler = ZEND_EXT_NOP_SPEC_HANDLER
    case ZEND_TICKS:
        handler = ZEND_TICKS_SPEC_HANDLER
    case ZEND_SEND_VAR_NO_REF:
        handler = ZEND_SEND_VAR_NO_REF_SPEC_VAR_HANDLER
    case ZEND_CATCH:
        handler = ZEND_CATCH_SPEC_CONST_HANDLER
    case ZEND_THROW:
        handler = vmGetThrowHandler(op)
    case ZEND_FETCH_CLASS:
        handler = vmGetFetchClassHandler(op)
    case ZEND_CLONE:
        handler = vmGetCloneHandler(op)
    case ZEND_RETURN_BY_REF:
        handler = vmGetReturnByRefHandler(op)
    case ZEND_INIT_METHOD_CALL:
        handler = vmGetInitMethodCallHandler(op)
    case ZEND_INIT_STATIC_METHOD_CALL:
        handler = vmGetInitStaticMethodCallHandler(op)
    case ZEND_ISSET_ISEMPTY_VAR:
        handler = vmGetIssetIsemptyVarHandler(op)
    case ZEND_ISSET_ISEMPTY_DIM_OBJ:
        handler = vmGetIssetIsemptyDimObjHandler(op)
    case ZEND_SEND_VAL_EX:
        handler = vmGetSendValExHandler(op)
    case ZEND_SEND_VAR:
        handler = vmGetSendVarHandler(op)
    case ZEND_INIT_USER_CALL:
        handler = vmGetInitUserCallHandler(op)
    case ZEND_SEND_ARRAY:
        handler = ZEND_SEND_ARRAY_SPEC_HANDLER
    case ZEND_SEND_USER:
        handler = vmGetSendUserHandler(op)
    case ZEND_STRLEN:
        handler = vmGetStrlenHandler(op)
    case ZEND_DEFINED:
        handler = ZEND_DEFINED_SPEC_CONST_HANDLER
    case ZEND_TYPE_CHECK:
        handler = vmGetTypeCheckHandler(op)
    case ZEND_VERIFY_RETURN_TYPE:
        handler = vmGetVerifyReturnTypeHandler(op)
    case ZEND_FE_RESET_RW:
        handler = vmGetFeResetRwHandler(op)
    case ZEND_FE_FETCH_RW:
        handler = ZEND_FE_FETCH_RW_SPEC_VAR_HANDLER
    case ZEND_FE_FREE:
        handler = ZEND_FE_FREE_SPEC_TMPVAR_HANDLER
    case ZEND_INIT_DYNAMIC_CALL:
        handler = vmGetInitDynamicCallHandler(op)
    case ZEND_DO_ICALL:
        handler = vmGetDoIcallHandler(op)
    case ZEND_DO_UCALL:
        handler = vmGetDoUcallHandler(op)
    case ZEND_DO_FCALL_BY_NAME:
        handler = vmGetDoFcallByNameHandler(op)
    case ZEND_PRE_INC_OBJ:
        handler = vmGetPreIncObjHandler(op)
    case ZEND_PRE_DEC_OBJ:
        handler = vmGetPreDecObjHandler(op)
    case ZEND_POST_INC_OBJ:
        handler = vmGetPostIncObjHandler(op)
    case ZEND_POST_DEC_OBJ:
        handler = vmGetPostDecObjHandler(op)
    case ZEND_ECHO:
        handler = vmGetEchoHandler(op)
    case ZEND_OP_DATA:
        handler = nil
    case ZEND_INSTANCEOF:
        handler = vmGetInstanceofHandler(op)
    case ZEND_GENERATOR_CREATE:
        handler = ZEND_GENERATOR_CREATE_SPEC_HANDLER
    case ZEND_MAKE_REF:
        handler = vmGetMakeRefHandler(op)
    case ZEND_DECLARE_FUNCTION:
        handler = ZEND_DECLARE_FUNCTION_SPEC_HANDLER
    case ZEND_DECLARE_LAMBDA_FUNCTION:
        handler = ZEND_DECLARE_LAMBDA_FUNCTION_SPEC_CONST_UNUSED_HANDLER
    case ZEND_DECLARE_CONST:
        handler = ZEND_DECLARE_CONST_SPEC_CONST_CONST_HANDLER
    case ZEND_DECLARE_CLASS:
        handler = ZEND_DECLARE_CLASS_SPEC_CONST_HANDLER
    case ZEND_DECLARE_CLASS_DELAYED:
        handler = ZEND_DECLARE_CLASS_DELAYED_SPEC_CONST_CONST_HANDLER
    case ZEND_DECLARE_ANON_CLASS:
        handler = ZEND_DECLARE_ANON_CLASS_SPEC_HANDLER
    case ZEND_ADD_ARRAY_UNPACK:
        handler = ZEND_ADD_ARRAY_UNPACK_SPEC_HANDLER
    case ZEND_ISSET_ISEMPTY_PROP_OBJ:
        handler = vmGetIssetIsemptyPropObjHandler(op)
    case ZEND_HANDLE_EXCEPTION:
        handler = ZEND_HANDLE_EXCEPTION_SPEC_HANDLER
    case ZEND_USER_OPCODE:
        handler = ZEND_USER_OPCODE_SPEC_HANDLER
    case ZEND_ASSERT_CHECK:
        handler = ZEND_ASSERT_CHECK_SPEC_HANDLER
    case ZEND_JMP_SET:
        handler = vmGetJmpSetHandler(op)
    case ZEND_UNSET_CV:
        handler = ZEND_UNSET_CV_SPEC_CV_UNUSED_HANDLER
    case ZEND_ISSET_ISEMPTY_CV:
        handler = vmGetIssetIsemptyCvHandler(op)
    case ZEND_FETCH_LIST_W:
        handler = vmGetFetchListWHandler(op)
    case ZEND_SEPARATE:
        handler = ZEND_SEPARATE_SPEC_VAR_UNUSED_HANDLER
    case ZEND_FETCH_CLASS_NAME:
        handler = ZEND_FETCH_CLASS_NAME_SPEC_UNUSED_HANDLER
    case ZEND_CALL_TRAMPOLINE:
        handler = ZEND_CALL_TRAMPOLINE_SPEC_HANDLER
    case ZEND_DISCARD_EXCEPTION:
        handler = ZEND_DISCARD_EXCEPTION_SPEC_HANDLER
    case ZEND_YIELD:
        handler = vmGetYieldHandler(op)
    case ZEND_GENERATOR_RETURN:
        handler = vmGetGeneratorReturnHandler(op)
    case ZEND_FAST_CALL:
        handler = ZEND_FAST_CALL_SPEC_HANDLER
    case ZEND_FAST_RET:
        handler = ZEND_FAST_RET_SPEC_HANDLER
    case ZEND_RECV_VARIADIC:
        handler = ZEND_RECV_VARIADIC_SPEC_UNUSED_HANDLER
    case ZEND_SEND_UNPACK:
        handler = ZEND_SEND_UNPACK_SPEC_HANDLER
    case ZEND_YIELD_FROM:
        handler = vmGetYieldFromHandler(op)
    case ZEND_COPY_TMP:
        handler = ZEND_COPY_TMP_SPEC_TMPVAR_UNUSED_HANDLER
    case ZEND_BIND_GLOBAL:
        handler = ZEND_BIND_GLOBAL_SPEC_CV_CONST_HANDLER
    case ZEND_COALESCE:
        handler = vmGetCoalesceHandler(op)
    case ZEND_SPACESHIP:
        handler = vmGetSpaceshipHandler(op)
    case ZEND_FUNC_NUM_ARGS:
        handler = ZEND_FUNC_NUM_ARGS_SPEC_UNUSED_UNUSED_HANDLER
    case ZEND_FUNC_GET_ARGS:
        handler = vmGetFuncGetArgsHandler(op)
    case ZEND_FETCH_STATIC_PROP_R:
        handler = ZEND_FETCH_STATIC_PROP_R_SPEC_HANDLER
    case ZEND_FETCH_STATIC_PROP_W:
        handler = ZEND_FETCH_STATIC_PROP_W_SPEC_HANDLER
    case ZEND_FETCH_STATIC_PROP_RW:
        handler = ZEND_FETCH_STATIC_PROP_RW_SPEC_HANDLER
    case ZEND_FETCH_STATIC_PROP_IS:
        handler = ZEND_FETCH_STATIC_PROP_IS_SPEC_HANDLER
    case ZEND_FETCH_STATIC_PROP_FUNC_ARG:
        handler = ZEND_FETCH_STATIC_PROP_FUNC_ARG_SPEC_HANDLER
    case ZEND_FETCH_STATIC_PROP_UNSET:
        handler = ZEND_FETCH_STATIC_PROP_UNSET_SPEC_HANDLER
    case ZEND_UNSET_STATIC_PROP:
        handler = ZEND_UNSET_STATIC_PROP_SPEC_HANDLER
    case ZEND_ISSET_ISEMPTY_STATIC_PROP:
        handler = ZEND_ISSET_ISEMPTY_STATIC_PROP_SPEC_HANDLER
    case ZEND_FETCH_CLASS_CONSTANT:
        handler = vmGetFetchClassConstantHandler(op)
    case ZEND_BIND_LEXICAL:
        handler = ZEND_BIND_LEXICAL_SPEC_TMP_CV_HANDLER
    case ZEND_BIND_STATIC:
        handler = ZEND_BIND_STATIC_SPEC_CV_UNUSED_HANDLER
    case ZEND_FETCH_THIS:
        handler = ZEND_FETCH_THIS_SPEC_UNUSED_UNUSED_HANDLER
    case ZEND_SEND_FUNC_ARG:
        handler = ZEND_SEND_FUNC_ARG_SPEC_VAR_HANDLER
    case ZEND_ISSET_ISEMPTY_THIS:
        handler = ZEND_ISSET_ISEMPTY_THIS_SPEC_UNUSED_UNUSED_HANDLER
    case ZEND_SWITCH_LONG:
        handler = vmGetSwitchLongHandler(op)
    case ZEND_SWITCH_STRING:
        handler = vmGetSwitchStringHandler(op)
    case ZEND_IN_ARRAY:
        handler = vmGetInArrayHandler(op)
    case ZEND_COUNT:
        handler = vmGetCountHandler(op)
    case ZEND_GET_CLASS:
        handler = vmGetGetClassHandler(op)
    case ZEND_GET_CALLED_CLASS:
        handler = ZEND_GET_CALLED_CLASS_SPEC_UNUSED_UNUSED_HANDLER
    case ZEND_GET_TYPE:
        handler = vmGetGetTypeHandler(op)
    case ZEND_ARRAY_KEY_EXISTS:
        handler = vmGetArrayKeyExistsHandler(op)
    }
    if handler == nil {
        return ZEND_NULL_HANDLER
    }
    return handler
}
func vmOpcodeIsCommutative(opcode OpCode) bool {
    switch opcode {
    case ZEND_MUL, ZEND_BW_OR, ZEND_BW_AND, ZEND_BW_XOR, ZEND_BOOL_XOR, ZEND_IS_IDENTICAL, ZEND_IS_NOT_IDENTICAL, ZEND_IS_EQUAL, ZEND_IS_NOT_EQUAL:
        return true
    }
    return false
}
func vmGetAddHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_ADD_SPEC_CONST_CONST_HANDLER,
        ZEND_ADD_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_ADD_SPEC_CONST_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_ADD_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ADD_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
    }
    return handlers[offset]
}
func vmGetSubHandler(op *ZendOp) OpcodeHandlerT {
    spec := 26 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_SUB_SPEC_CONST_CONST_HANDLER,
        ZEND_SUB_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SUB_SPEC_CONST_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SUB_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_SUB_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER,
    }
    return handlers[offset]
}
func vmGetMulHandler(op *ZendOp) OpcodeHandlerT {
    spec := 51 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_MUL_SPEC_CONST_CONST_HANDLER,
        nil,
        nil,
        nil,
        nil,
        ZEND_MUL_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_MUL_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
    }
    return handlers[offset]
}
func vmGetDivHandler(op *ZendOp) OpcodeHandlerT {
    spec := 76 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_DIV_SPEC_CONST_CONST_HANDLER,
        ZEND_DIV_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_DIV_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_DIV_SPEC_CONST_CV_HANDLER,
        ZEND_DIV_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_DIV_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_DIV_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_DIV_SPEC_TMPVAR_CV_HANDLER,
        ZEND_DIV_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_DIV_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_DIV_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_DIV_SPEC_TMPVAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_DIV_SPEC_CV_CONST_HANDLER,
        ZEND_DIV_SPEC_CV_TMPVAR_HANDLER,
        ZEND_DIV_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_DIV_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetModHandler(op *ZendOp) OpcodeHandlerT {
    spec := 101 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_MOD_SPEC_CONST_CONST_HANDLER,
        ZEND_MOD_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_MOD_SPEC_CONST_TMPVARCV_HANDLER,
        nil,
        ZEND_MOD_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_MOD_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MOD_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_MOD_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER,
    }
    return handlers[offset]
}
func vmGetSlHandler(op *ZendOp) OpcodeHandlerT {
    spec := 126 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_SL_SPEC_CONST_CONST_HANDLER,
        ZEND_SL_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SL_SPEC_CONST_TMPVARCV_HANDLER,
        nil,
        ZEND_SL_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SL_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SL_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_SL_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
    }
    return handlers[offset]
}
func vmGetSrHandler(op *ZendOp) OpcodeHandlerT {
    spec := 151 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_SR_SPEC_CONST_CONST_HANDLER,
        ZEND_SR_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SR_SPEC_CONST_TMPVARCV_HANDLER,
        nil,
        ZEND_SR_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SR_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SR_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_SR_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
    }
    return handlers[offset]
}
func vmGetConcatHandler(op *ZendOp) OpcodeHandlerT {
    spec := 176 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        ZEND_CONCAT_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_CONCAT_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_CONCAT_SPEC_CONST_CV_HANDLER,
        ZEND_CONCAT_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_CONCAT_SPEC_TMPVAR_CV_HANDLER,
        ZEND_CONCAT_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_CONCAT_SPEC_TMPVAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_CONCAT_SPEC_CV_CONST_HANDLER,
        ZEND_CONCAT_SPEC_CV_TMPVAR_HANDLER,
        ZEND_CONCAT_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_CONCAT_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetBwOrHandler(op *ZendOp) OpcodeHandlerT {
    spec := 201 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_BW_OR_SPEC_CONST_CONST_HANDLER,
        nil,
        nil,
        nil,
        nil,
        ZEND_BW_OR_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_OR_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_BW_OR_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
    }
    return handlers[offset]
}
func vmGetBwAndHandler(op *ZendOp) OpcodeHandlerT {
    spec := 226 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_BW_AND_SPEC_CONST_CONST_HANDLER,
        nil,
        nil,
        nil,
        nil,
        ZEND_BW_AND_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_AND_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_BW_AND_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER,
    }
    return handlers[offset]
}
func vmGetBwXorHandler(op *ZendOp) OpcodeHandlerT {
    spec := 251 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_BW_XOR_SPEC_CONST_CONST_HANDLER,
        nil,
        nil,
        nil,
        nil,
        ZEND_BW_XOR_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_XOR_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_BW_XOR_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER,
    }
    return handlers[offset]
}
func vmGetPowHandler(op *ZendOp) OpcodeHandlerT {
    spec := 276 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_POW_SPEC_CONST_CONST_HANDLER,
        ZEND_POW_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_POW_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_POW_SPEC_CONST_CV_HANDLER,
        ZEND_POW_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_POW_SPEC_TMPVAR_CV_HANDLER,
        ZEND_POW_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_POW_SPEC_TMPVAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_POW_SPEC_CV_CONST_HANDLER,
        ZEND_POW_SPEC_CV_TMPVAR_HANDLER,
        ZEND_POW_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_POW_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetBwNotHandler(op *ZendOp) OpcodeHandlerT {
    spec := 301 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_BW_NOT_SPEC_CONST_HANDLER,
        ZEND_BW_NOT_SPEC_TMPVAR_HANDLER,
        ZEND_BW_NOT_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_BW_NOT_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetBoolNotHandler(op *ZendOp) OpcodeHandlerT {
    spec := 306 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_BOOL_NOT_SPEC_CONST_HANDLER,
        ZEND_BOOL_NOT_SPEC_TMPVAR_HANDLER,
        ZEND_BOOL_NOT_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_BOOL_NOT_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetBoolXorHandler(op *ZendOp) OpcodeHandlerT {
    spec := 311 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_BOOL_XOR_SPEC_CONST_CONST_HANDLER,
        nil,
        nil,
        nil,
        nil,
        ZEND_BOOL_XOR_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_BOOL_XOR_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_BOOL_XOR_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        nil,
        ZEND_BOOL_XOR_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_BOOL_XOR_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_BOOL_XOR_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_BOOL_XOR_SPEC_CV_CONST_HANDLER,
        ZEND_BOOL_XOR_SPEC_CV_TMPVAR_HANDLER,
        ZEND_BOOL_XOR_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_BOOL_XOR_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetIsIdenticalHandler(op *ZendOp) OpcodeHandlerT {
    spec := 336 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_IS_IDENTICAL_SPEC_CONST_CONST_HANDLER,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_IDENTICAL_SPEC_TMP_CONST_HANDLER,
        ZEND_IS_IDENTICAL_SPEC_TMP_TMP_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_IDENTICAL_SPEC_VAR_CONST_HANDLER,
        ZEND_IS_IDENTICAL_SPEC_VAR_TMP_HANDLER,
        ZEND_IS_IDENTICAL_SPEC_VAR_VAR_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_IDENTICAL_SPEC_CV_CONST_HANDLER,
        ZEND_IS_IDENTICAL_SPEC_CV_TMP_HANDLER,
        ZEND_IS_IDENTICAL_SPEC_CV_VAR_HANDLER,
        nil,
        ZEND_IS_IDENTICAL_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetIsNotIdenticalHandler(op *ZendOp) OpcodeHandlerT {
    spec := 361 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_COMMUTATIVE
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_IS_NOT_IDENTICAL_SPEC_CONST_CONST_HANDLER,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_IDENTICAL_SPEC_TMP_CONST_HANDLER,
        ZEND_IS_NOT_IDENTICAL_SPEC_TMP_TMP_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_IDENTICAL_SPEC_VAR_CONST_HANDLER,
        ZEND_IS_NOT_IDENTICAL_SPEC_VAR_TMP_HANDLER,
        ZEND_IS_NOT_IDENTICAL_SPEC_VAR_VAR_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_IDENTICAL_SPEC_CV_CONST_HANDLER,
        ZEND_IS_NOT_IDENTICAL_SPEC_CV_TMP_HANDLER,
        ZEND_IS_NOT_IDENTICAL_SPEC_CV_VAR_HANDLER,
        nil,
        ZEND_IS_NOT_IDENTICAL_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetIsEqualHandler(op *ZendOp) OpcodeHandlerT {
    spec := 386 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH | SPEC_RULE_COMMUTATIVE
    offset := vmOffsetBySpec(spec, op)
    handlers := [75]OpcodeHandlerT{
        ZEND_IS_EQUAL_SPEC_CONST_CONST_HANDLER,
        ZEND_IS_EQUAL_SPEC_CONST_CONST_HANDLER,
        ZEND_IS_EQUAL_SPEC_CONST_CONST_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_SPEC_CV_CONST_HANDLER,
        ZEND_IS_EQUAL_SPEC_CV_CONST_JMPZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_CV_CONST_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_CV_TMPVAR_HANDLER,
        ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_CV_TMPVAR_HANDLER,
        ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_SPEC_CV_CV_HANDLER,
        ZEND_IS_EQUAL_SPEC_CV_CV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_SPEC_CV_CV_JMPNZ_HANDLER,
    }
    return handlers[offset]
}
func vmGetIsNotEqualHandler(op *ZendOp) OpcodeHandlerT {
    spec := 461 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH | SPEC_RULE_COMMUTATIVE
    offset := vmOffsetBySpec(spec, op)
    handlers := [75]OpcodeHandlerT{
        ZEND_IS_NOT_EQUAL_SPEC_CONST_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CONST_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CONST_CONST_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_SPEC_CV_CV_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CV_CV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_SPEC_CV_CV_JMPNZ_HANDLER,
    }
    return handlers[offset]
}
func vmGetIsSmallerHandler(op *ZendOp) OpcodeHandlerT {
    spec := 536 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH
    offset := vmOffsetBySpec(spec, op)
    handlers := [75]OpcodeHandlerT{
        ZEND_IS_SMALLER_SPEC_CONST_CONST_HANDLER,
        ZEND_IS_SMALLER_SPEC_CONST_CONST_HANDLER,
        ZEND_IS_SMALLER_SPEC_CONST_CONST_HANDLER,
        ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
    }
    return handlers[offset]
}
func vmGetIsSmallerOrEqualHandler(op *ZendOp) OpcodeHandlerT {
    spec := 611 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_SMART_BRANCH
    offset := vmOffsetBySpec(spec, op)
    handlers := [75]OpcodeHandlerT{
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
    }
    return handlers[offset]
}
func vmGetAssignHandler(op *ZendOp) OpcodeHandlerT {
    spec := 686 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_RETVAL
    offset := vmOffsetBySpec(spec, op)
    handlers := [50]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_SPEC_VAR_CONST_RETVAL_UNUSED_HANDLER,
        ZEND_ASSIGN_SPEC_VAR_CONST_RETVAL_USED_HANDLER,
        ZEND_ASSIGN_SPEC_VAR_TMP_RETVAL_UNUSED_HANDLER,
        ZEND_ASSIGN_SPEC_VAR_TMP_RETVAL_USED_HANDLER,
        ZEND_ASSIGN_SPEC_VAR_VAR_RETVAL_UNUSED_HANDLER,
        ZEND_ASSIGN_SPEC_VAR_VAR_RETVAL_USED_HANDLER,
        nil,
        nil,
        ZEND_ASSIGN_SPEC_VAR_CV_RETVAL_UNUSED_HANDLER,
        ZEND_ASSIGN_SPEC_VAR_CV_RETVAL_USED_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_SPEC_CV_CONST_RETVAL_UNUSED_HANDLER,
        ZEND_ASSIGN_SPEC_CV_CONST_RETVAL_USED_HANDLER,
        ZEND_ASSIGN_SPEC_CV_TMP_RETVAL_UNUSED_HANDLER,
        ZEND_ASSIGN_SPEC_CV_TMP_RETVAL_USED_HANDLER,
        ZEND_ASSIGN_SPEC_CV_VAR_RETVAL_UNUSED_HANDLER,
        ZEND_ASSIGN_SPEC_CV_VAR_RETVAL_USED_HANDLER,
        nil,
        nil,
        ZEND_ASSIGN_SPEC_CV_CV_RETVAL_UNUSED_HANDLER,
        ZEND_ASSIGN_SPEC_CV_CV_RETVAL_USED_HANDLER,
    }
    return handlers[offset]
}
func vmGetAssignDimHandler(op *ZendOp) OpcodeHandlerT {
    spec := 736 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_OP_DATA
    offset := vmOffsetBySpec(spec, op)
    handlers := [125]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetAssignObjHandler(op *ZendOp) OpcodeHandlerT {
    spec := 861 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_OP_DATA
    offset := vmOffsetBySpec(spec, op)
    handlers := [125]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetAssignStaticPropHandler(op *ZendOp) OpcodeHandlerT {
    spec := 986 | SPEC_RULE_OP_DATA
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_CONST_HANDLER,
        ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_TMP_HANDLER,
        ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetAssignOpHandler(op *ZendOp) OpcodeHandlerT {
    spec := 991 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OP_SPEC_VAR_CONST_HANDLER,
        ZEND_ASSIGN_OP_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_ASSIGN_OP_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_ASSIGN_OP_SPEC_VAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OP_SPEC_CV_CONST_HANDLER,
        ZEND_ASSIGN_OP_SPEC_CV_TMPVAR_HANDLER,
        ZEND_ASSIGN_OP_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_ASSIGN_OP_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetAssignDimOpHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1016 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_DIM_OP_SPEC_VAR_CONST_HANDLER,
        ZEND_ASSIGN_DIM_OP_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_ASSIGN_DIM_OP_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_ASSIGN_DIM_OP_SPEC_VAR_UNUSED_HANDLER,
        ZEND_ASSIGN_DIM_OP_SPEC_VAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_DIM_OP_SPEC_CV_CONST_HANDLER,
        ZEND_ASSIGN_DIM_OP_SPEC_CV_TMPVAR_HANDLER,
        ZEND_ASSIGN_DIM_OP_SPEC_CV_TMPVAR_HANDLER,
        ZEND_ASSIGN_DIM_OP_SPEC_CV_UNUSED_HANDLER,
        ZEND_ASSIGN_DIM_OP_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetAssignObjOpHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1041 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_OP_SPEC_VAR_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_OP_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_ASSIGN_OBJ_OP_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_OP_SPEC_VAR_CV_HANDLER,
        ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_CV_HANDLER,
        ZEND_ASSIGN_OBJ_OP_SPEC_CV_CONST_HANDLER,
        ZEND_ASSIGN_OBJ_OP_SPEC_CV_TMPVAR_HANDLER,
        ZEND_ASSIGN_OBJ_OP_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_OP_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetAssignRefHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1067 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_REF_SPEC_VAR_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_REF_SPEC_VAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_REF_SPEC_CV_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_REF_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetQmAssignHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1092 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_QM_ASSIGN_SPEC_CONST_HANDLER,
        ZEND_QM_ASSIGN_SPEC_TMP_HANDLER,
        ZEND_QM_ASSIGN_SPEC_VAR_HANDLER,
        nil,
        ZEND_QM_ASSIGN_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetAssignObjRefHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1097 | SPEC_RULE_OP1 | SPEC_RULE_OP2 | SPEC_RULE_OP_DATA
    offset := vmOffsetBySpec(spec, op)
    handlers := [125]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CONST_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CONST_OP_DATA_CV_HANDLER,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CV_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CV_OP_DATA_CV_HANDLER,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CONST_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CONST_OP_DATA_CV_HANDLER,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CV_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CV_OP_DATA_CV_HANDLER,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_CV_CONST_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_CV_CONST_OP_DATA_CV_HANDLER,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_CV_CV_OP_DATA_VAR_HANDLER,
        nil,
        ZEND_ASSIGN_OBJ_REF_SPEC_CV_CV_OP_DATA_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetPreIncHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1223 | SPEC_RULE_OP1 | SPEC_RULE_RETVAL
    offset := vmOffsetBySpec(spec, op)
    handlers := [10]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        ZEND_PRE_INC_SPEC_VAR_RETVAL_UNUSED_HANDLER,
        ZEND_PRE_INC_SPEC_VAR_RETVAL_USED_HANDLER,
        nil,
        nil,
        ZEND_PRE_INC_SPEC_CV_RETVAL_UNUSED_HANDLER,
        ZEND_PRE_INC_SPEC_CV_RETVAL_USED_HANDLER,
    }
    return handlers[offset]
}
func vmGetPreDecHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1233 | SPEC_RULE_OP1 | SPEC_RULE_RETVAL
    offset := vmOffsetBySpec(spec, op)
    handlers := [10]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        ZEND_PRE_DEC_SPEC_VAR_RETVAL_UNUSED_HANDLER,
        ZEND_PRE_DEC_SPEC_VAR_RETVAL_USED_HANDLER,
        nil,
        nil,
        ZEND_PRE_DEC_SPEC_CV_RETVAL_UNUSED_HANDLER,
        ZEND_PRE_DEC_SPEC_CV_RETVAL_USED_HANDLER,
    }
    return handlers[offset]
}
func vmGetPostIncHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1243 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        nil,
        nil,
        ZEND_POST_INC_SPEC_VAR_HANDLER,
        nil,
        ZEND_POST_INC_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetPostDecHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1248 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        nil,
        nil,
        ZEND_POST_DEC_SPEC_VAR_HANDLER,
        nil,
        ZEND_POST_DEC_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetPreIncStaticPropHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1253
    offset := vmOffsetBySpec(spec, op)
    handlers := [0]OpcodeHandlerT{
    }
    return handlers[offset]
}
func vmGetPostIncStaticPropHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1254
    offset := vmOffsetBySpec(spec, op)
    handlers := [0]OpcodeHandlerT{
    }
    return handlers[offset]
}
func vmGetJmpzHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1256 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_JMPZ_SPEC_CONST_HANDLER,
        ZEND_JMPZ_SPEC_TMPVAR_HANDLER,
        ZEND_JMPZ_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_JMPZ_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetJmpnzHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1261 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_JMPNZ_SPEC_CONST_HANDLER,
        ZEND_JMPNZ_SPEC_TMPVAR_HANDLER,
        ZEND_JMPNZ_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_JMPNZ_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetJmpznzHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1266 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_JMPZNZ_SPEC_CONST_HANDLER,
        ZEND_JMPZNZ_SPEC_TMPVAR_HANDLER,
        ZEND_JMPZNZ_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_JMPZNZ_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetJmpzExHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1271 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_JMPZ_EX_SPEC_CONST_HANDLER,
        ZEND_JMPZ_EX_SPEC_TMPVAR_HANDLER,
        ZEND_JMPZ_EX_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_JMPZ_EX_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetJmpnzExHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1276 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_JMPNZ_EX_SPEC_CONST_HANDLER,
        ZEND_JMPNZ_EX_SPEC_TMPVAR_HANDLER,
        ZEND_JMPNZ_EX_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_JMPNZ_EX_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetCaseHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1281 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_CASE_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_CASE_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_CASE_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_CASE_SPEC_TMPVAR_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetSendVarNoRefExHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1287 | SPEC_RULE_QUICK_ARG
    offset := vmOffsetBySpec(spec, op)
    handlers := [2]OpcodeHandlerT{
        ZEND_SEND_VAR_NO_REF_EX_SPEC_VAR_HANDLER,
        ZEND_SEND_VAR_NO_REF_EX_SPEC_VAR_QUICK_HANDLER,
    }
    return handlers[offset]
}
func vmGetCastHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1289 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_CAST_SPEC_CONST_HANDLER,
        ZEND_CAST_SPEC_TMP_HANDLER,
        ZEND_CAST_SPEC_VAR_HANDLER,
        nil,
        ZEND_CAST_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetBoolHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1294 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_BOOL_SPEC_CONST_HANDLER,
        ZEND_BOOL_SPEC_TMPVAR_HANDLER,
        ZEND_BOOL_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_BOOL_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFastConcatHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1299 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_FAST_CONCAT_SPEC_CONST_CONST_HANDLER,
        ZEND_FAST_CONCAT_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_FAST_CONCAT_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_FAST_CONCAT_SPEC_CONST_CV_HANDLER,
        ZEND_FAST_CONCAT_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_FAST_CONCAT_SPEC_TMPVAR_CV_HANDLER,
        ZEND_FAST_CONCAT_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_FAST_CONCAT_SPEC_TMPVAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FAST_CONCAT_SPEC_CV_CONST_HANDLER,
        ZEND_FAST_CONCAT_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FAST_CONCAT_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_FAST_CONCAT_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetRopeInitHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1324 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_ROPE_INIT_SPEC_UNUSED_CONST_HANDLER,
        ZEND_ROPE_INIT_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_ROPE_INIT_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_ROPE_INIT_SPEC_UNUSED_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetRopeAddHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1329 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_ROPE_ADD_SPEC_TMP_CONST_HANDLER,
        ZEND_ROPE_ADD_SPEC_TMP_TMPVAR_HANDLER,
        ZEND_ROPE_ADD_SPEC_TMP_TMPVAR_HANDLER,
        nil,
        ZEND_ROPE_ADD_SPEC_TMP_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetRopeEndHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1334 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_ROPE_END_SPEC_TMP_CONST_HANDLER,
        ZEND_ROPE_END_SPEC_TMP_TMPVAR_HANDLER,
        ZEND_ROPE_END_SPEC_TMP_TMPVAR_HANDLER,
        nil,
        ZEND_ROPE_END_SPEC_TMP_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetDoFcallHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1342 | SPEC_RULE_RETVAL
    offset := vmOffsetBySpec(spec, op)
    handlers := [2]OpcodeHandlerT{
        ZEND_DO_FCALL_SPEC_RETVAL_UNUSED_HANDLER,
        ZEND_DO_FCALL_SPEC_RETVAL_USED_HANDLER,
    }
    return handlers[offset]
}
func vmGetReturnHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1345 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_RETURN_SPEC_CONST_HANDLER,
        ZEND_RETURN_SPEC_TMP_HANDLER,
        ZEND_RETURN_SPEC_VAR_HANDLER,
        nil,
        ZEND_RETURN_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetSendValHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1352 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_SEND_VAL_SPEC_CONST_HANDLER,
        ZEND_SEND_VAL_SPEC_TMPVAR_HANDLER,
        ZEND_SEND_VAL_SPEC_TMPVAR_HANDLER,
        nil,
        nil,
    }
    return handlers[offset]
}
func vmGetSendVarExHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1357 | SPEC_RULE_OP1 | SPEC_RULE_QUICK_ARG
    offset := vmOffsetBySpec(spec, op)
    handlers := [10]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        ZEND_SEND_VAR_EX_SPEC_VAR_HANDLER,
        ZEND_SEND_VAR_EX_SPEC_VAR_QUICK_HANDLER,
        nil,
        nil,
        ZEND_SEND_VAR_EX_SPEC_CV_HANDLER,
        ZEND_SEND_VAR_EX_SPEC_CV_QUICK_HANDLER,
    }
    return handlers[offset]
}
func vmGetSendRefHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1367 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        nil,
        nil,
        ZEND_SEND_REF_SPEC_VAR_HANDLER,
        nil,
        ZEND_SEND_REF_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetNewHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1372 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_NEW_SPEC_CONST_UNUSED_HANDLER,
        nil,
        ZEND_NEW_SPEC_VAR_UNUSED_HANDLER,
        ZEND_NEW_SPEC_UNUSED_UNUSED_HANDLER,
        nil,
    }
    return handlers[offset]
}
func vmGetInitArrayHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1379 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_INIT_ARRAY_SPEC_CONST_CONST_HANDLER,
        ZEND_INIT_ARRAY_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_INIT_ARRAY_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_INIT_ARRAY_SPEC_CONST_UNUSED_HANDLER,
        ZEND_INIT_ARRAY_SPEC_CONST_CV_HANDLER,
        ZEND_INIT_ARRAY_SPEC_TMP_CONST_HANDLER,
        ZEND_INIT_ARRAY_SPEC_TMP_TMPVAR_HANDLER,
        ZEND_INIT_ARRAY_SPEC_TMP_TMPVAR_HANDLER,
        ZEND_INIT_ARRAY_SPEC_TMP_UNUSED_HANDLER,
        ZEND_INIT_ARRAY_SPEC_TMP_CV_HANDLER,
        ZEND_INIT_ARRAY_SPEC_VAR_CONST_HANDLER,
        ZEND_INIT_ARRAY_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_INIT_ARRAY_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_INIT_ARRAY_SPEC_VAR_UNUSED_HANDLER,
        ZEND_INIT_ARRAY_SPEC_VAR_CV_HANDLER,
        ZEND_INIT_ARRAY_SPEC_UNUSED_CONST_HANDLER,
        ZEND_INIT_ARRAY_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_INIT_ARRAY_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_INIT_ARRAY_SPEC_UNUSED_UNUSED_HANDLER,
        ZEND_INIT_ARRAY_SPEC_UNUSED_CV_HANDLER,
        ZEND_INIT_ARRAY_SPEC_CV_CONST_HANDLER,
        ZEND_INIT_ARRAY_SPEC_CV_TMPVAR_HANDLER,
        ZEND_INIT_ARRAY_SPEC_CV_TMPVAR_HANDLER,
        ZEND_INIT_ARRAY_SPEC_CV_UNUSED_HANDLER,
        ZEND_INIT_ARRAY_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetAddArrayElementHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1404 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CONST_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_UNUSED_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CV_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CONST_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_TMPVAR_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_TMPVAR_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_UNUSED_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CV_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CONST_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_UNUSED_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CONST_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_TMPVAR_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_TMPVAR_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_UNUSED_HANDLER,
        ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetIncludeOrEvalHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1429 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_INCLUDE_OR_EVAL_SPEC_CONST_HANDLER,
        ZEND_INCLUDE_OR_EVAL_SPEC_TMPVAR_HANDLER,
        ZEND_INCLUDE_OR_EVAL_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_INCLUDE_OR_EVAL_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetUnsetVarHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1434 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_UNSET_VAR_SPEC_CONST_UNUSED_HANDLER,
        ZEND_UNSET_VAR_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_UNSET_VAR_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        ZEND_UNSET_VAR_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetUnsetDimHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1439 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_UNSET_DIM_SPEC_VAR_CONST_HANDLER,
        ZEND_UNSET_DIM_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_UNSET_DIM_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_UNSET_DIM_SPEC_VAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_UNSET_DIM_SPEC_CV_CONST_HANDLER,
        ZEND_UNSET_DIM_SPEC_CV_TMPVAR_HANDLER,
        ZEND_UNSET_DIM_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_UNSET_DIM_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetUnsetObjHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1464 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_UNSET_OBJ_SPEC_VAR_CONST_HANDLER,
        ZEND_UNSET_OBJ_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_UNSET_OBJ_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_UNSET_OBJ_SPEC_VAR_CV_HANDLER,
        ZEND_UNSET_OBJ_SPEC_UNUSED_CONST_HANDLER,
        ZEND_UNSET_OBJ_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_UNSET_OBJ_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_UNSET_OBJ_SPEC_UNUSED_CV_HANDLER,
        ZEND_UNSET_OBJ_SPEC_CV_CONST_HANDLER,
        ZEND_UNSET_OBJ_SPEC_CV_TMPVAR_HANDLER,
        ZEND_UNSET_OBJ_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_UNSET_OBJ_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFeResetRHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1489 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FE_RESET_R_SPEC_CONST_HANDLER,
        ZEND_FE_RESET_R_SPEC_TMP_HANDLER,
        ZEND_FE_RESET_R_SPEC_VAR_HANDLER,
        nil,
        ZEND_FE_RESET_R_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchRHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1496 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FETCH_R_SPEC_CONST_UNUSED_HANDLER,
        ZEND_FETCH_R_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_FETCH_R_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        ZEND_FETCH_R_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchDimRHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1501 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_FETCH_DIM_R_SPEC_CONST_CONST_HANDLER,
        ZEND_FETCH_DIM_R_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_R_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_DIM_R_SPEC_CONST_CV_HANDLER,
        ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER,
        ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_DIM_R_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_DIM_R_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_R_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_DIM_R_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchObjRHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1526 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_FETCH_OBJ_R_SPEC_CONST_CONST_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_R_SPEC_CONST_CV_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_R_SPEC_UNUSED_CV_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_R_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_R_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchWHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1551 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FETCH_W_SPEC_CONST_UNUSED_HANDLER,
        ZEND_FETCH_W_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_FETCH_W_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        ZEND_FETCH_W_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchDimWHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1556 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_DIM_W_SPEC_VAR_CONST_HANDLER,
        ZEND_FETCH_DIM_W_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_W_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_W_SPEC_VAR_UNUSED_HANDLER,
        ZEND_FETCH_DIM_W_SPEC_VAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_DIM_W_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_DIM_W_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_W_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_W_SPEC_CV_UNUSED_HANDLER,
        ZEND_FETCH_DIM_W_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchObjWHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1581 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_OBJ_W_SPEC_VAR_CONST_HANDLER,
        ZEND_FETCH_OBJ_W_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_W_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_W_SPEC_VAR_CV_HANDLER,
        ZEND_FETCH_OBJ_W_SPEC_UNUSED_CONST_HANDLER,
        ZEND_FETCH_OBJ_W_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_W_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_W_SPEC_UNUSED_CV_HANDLER,
        ZEND_FETCH_OBJ_W_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_OBJ_W_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_W_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_W_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchRwHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1606 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FETCH_RW_SPEC_CONST_UNUSED_HANDLER,
        ZEND_FETCH_RW_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_FETCH_RW_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        ZEND_FETCH_RW_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchDimRwHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1611 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_DIM_RW_SPEC_VAR_CONST_HANDLER,
        ZEND_FETCH_DIM_RW_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_RW_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_RW_SPEC_VAR_UNUSED_HANDLER,
        ZEND_FETCH_DIM_RW_SPEC_VAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_DIM_RW_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_DIM_RW_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_RW_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_RW_SPEC_CV_UNUSED_HANDLER,
        ZEND_FETCH_DIM_RW_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchObjRwHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1636 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_OBJ_RW_SPEC_VAR_CONST_HANDLER,
        ZEND_FETCH_OBJ_RW_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_RW_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_RW_SPEC_VAR_CV_HANDLER,
        ZEND_FETCH_OBJ_RW_SPEC_UNUSED_CONST_HANDLER,
        ZEND_FETCH_OBJ_RW_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_RW_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_RW_SPEC_UNUSED_CV_HANDLER,
        ZEND_FETCH_OBJ_RW_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_OBJ_RW_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_RW_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_RW_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchIsHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1661 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FETCH_IS_SPEC_CONST_UNUSED_HANDLER,
        ZEND_FETCH_IS_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_FETCH_IS_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        ZEND_FETCH_IS_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchDimIsHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1666 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_FETCH_DIM_IS_SPEC_CONST_CONST_HANDLER,
        ZEND_FETCH_DIM_IS_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_IS_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_DIM_IS_SPEC_CONST_CV_HANDLER,
        ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CV_HANDLER,
        ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_DIM_IS_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_DIM_IS_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_IS_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_DIM_IS_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchObjIsHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1691 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_FETCH_OBJ_IS_SPEC_CONST_CONST_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_IS_SPEC_CONST_CV_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CV_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CV_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_UNUSED_CONST_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_IS_SPEC_UNUSED_CV_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_IS_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_IS_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchFuncArgHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1716 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FETCH_FUNC_ARG_SPEC_CONST_UNUSED_HANDLER,
        ZEND_FETCH_FUNC_ARG_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_FETCH_FUNC_ARG_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        ZEND_FETCH_FUNC_ARG_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchDimFuncArgHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1721 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_CONST_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_UNUSED_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_CV_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_CONST_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_UNUSED_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_CV_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_CONST_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_UNUSED_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_UNUSED_HANDLER,
        ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchObjFuncArgHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1746 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_CONST_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_CV_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_CONST_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_CV_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_CONST_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_CV_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_CONST_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_CV_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchUnsetHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1771 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FETCH_UNSET_SPEC_CONST_UNUSED_HANDLER,
        ZEND_FETCH_UNSET_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_FETCH_UNSET_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        ZEND_FETCH_UNSET_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchDimUnsetHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1776 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_DIM_UNSET_SPEC_VAR_CONST_HANDLER,
        ZEND_FETCH_DIM_UNSET_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_UNSET_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_DIM_UNSET_SPEC_VAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_DIM_UNSET_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_DIM_UNSET_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_DIM_UNSET_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_DIM_UNSET_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchObjUnsetHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1801 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_OBJ_UNSET_SPEC_VAR_CONST_HANDLER,
        ZEND_FETCH_OBJ_UNSET_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_UNSET_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_UNSET_SPEC_VAR_CV_HANDLER,
        ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_CONST_HANDLER,
        ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_CV_HANDLER,
        ZEND_FETCH_OBJ_UNSET_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_OBJ_UNSET_SPEC_CV_TMPVAR_HANDLER,
        ZEND_FETCH_OBJ_UNSET_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_OBJ_UNSET_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchListRHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1826 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_FETCH_LIST_R_SPEC_CONST_CONST_HANDLER,
        ZEND_FETCH_LIST_R_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_FETCH_LIST_R_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_LIST_R_SPEC_CONST_CV_HANDLER,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CV_HANDLER,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetCheckFuncArgHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1852 | SPEC_RULE_QUICK_ARG
    offset := vmOffsetBySpec(spec, op)
    handlers := [2]OpcodeHandlerT{
        ZEND_CHECK_FUNC_ARG_SPEC_UNUSED_HANDLER,
        ZEND_CHECK_FUNC_ARG_SPEC_UNUSED_QUICK_HANDLER,
    }
    return handlers[offset]
}
func vmGetThrowHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1861 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_THROW_SPEC_CONST_HANDLER,
        ZEND_THROW_SPEC_TMP_HANDLER,
        ZEND_THROW_SPEC_VAR_HANDLER,
        nil,
        ZEND_THROW_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchClassHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1866 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FETCH_CLASS_SPEC_UNUSED_CONST_HANDLER,
        ZEND_FETCH_CLASS_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_FETCH_CLASS_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_FETCH_CLASS_SPEC_UNUSED_UNUSED_HANDLER,
        ZEND_FETCH_CLASS_SPEC_UNUSED_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetCloneHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1871 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_CLONE_SPEC_CONST_HANDLER,
        ZEND_CLONE_SPEC_TMPVAR_HANDLER,
        ZEND_CLONE_SPEC_TMPVAR_HANDLER,
        ZEND_CLONE_SPEC_UNUSED_HANDLER,
        ZEND_CLONE_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetReturnByRefHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1876 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_RETURN_BY_REF_SPEC_CONST_HANDLER,
        ZEND_RETURN_BY_REF_SPEC_TMP_HANDLER,
        ZEND_RETURN_BY_REF_SPEC_VAR_HANDLER,
        nil,
        ZEND_RETURN_BY_REF_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetInitMethodCallHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1881 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_INIT_METHOD_CALL_SPEC_CONST_CONST_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_INIT_METHOD_CALL_SPEC_CONST_CV_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CV_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CV_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_UNUSED_CONST_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_INIT_METHOD_CALL_SPEC_UNUSED_CV_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_CV_CONST_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_CV_TMPVAR_HANDLER,
        ZEND_INIT_METHOD_CALL_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_INIT_METHOD_CALL_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetInitStaticMethodCallHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1906 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_CONST_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_UNUSED_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_CONST_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_UNUSED_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_CV_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_CONST_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_UNUSED_HANDLER,
        ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
    }
    return handlers[offset]
}
func vmGetIssetIsemptyVarHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1931 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_ISSET_ISEMPTY_VAR_SPEC_CONST_UNUSED_HANDLER,
        ZEND_ISSET_ISEMPTY_VAR_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_ISSET_ISEMPTY_VAR_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        ZEND_ISSET_ISEMPTY_VAR_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetIssetIsemptyDimObjHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1936 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_CONST_HANDLER,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_CV_HANDLER,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CV_HANDLER,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_CONST_HANDLER,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_TMPVAR_HANDLER,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetSendValExHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1961 | SPEC_RULE_OP1 | SPEC_RULE_QUICK_ARG
    offset := vmOffsetBySpec(spec, op)
    handlers := [10]OpcodeHandlerT{
        ZEND_SEND_VAL_EX_SPEC_CONST_HANDLER,
        ZEND_SEND_VAL_EX_SPEC_CONST_QUICK_HANDLER,
        ZEND_SEND_VAL_EX_SPEC_TMP_HANDLER,
        ZEND_SEND_VAL_EX_SPEC_TMP_QUICK_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
    }
    return handlers[offset]
}
func vmGetSendVarHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1971 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        nil,
        nil,
        ZEND_SEND_VAR_SPEC_VAR_HANDLER,
        nil,
        ZEND_SEND_VAR_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetInitUserCallHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1976 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_INIT_USER_CALL_SPEC_CONST_CONST_HANDLER,
        ZEND_INIT_USER_CALL_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_INIT_USER_CALL_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_INIT_USER_CALL_SPEC_CONST_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetSendUserHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1982 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_SEND_USER_SPEC_CONST_HANDLER,
        ZEND_SEND_USER_SPEC_TMP_HANDLER,
        ZEND_SEND_USER_SPEC_VAR_HANDLER,
        nil,
        ZEND_SEND_USER_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetStrlenHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1987 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_STRLEN_SPEC_CONST_HANDLER,
        ZEND_STRLEN_SPEC_TMPVAR_HANDLER,
        ZEND_STRLEN_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_STRLEN_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetTypeCheckHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1993 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_TYPE_CHECK_SPEC_CONST_HANDLER,
        ZEND_TYPE_CHECK_SPEC_TMPVAR_HANDLER,
        ZEND_TYPE_CHECK_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_TYPE_CHECK_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetVerifyReturnTypeHandler(op *ZendOp) OpcodeHandlerT {
    spec := 1998 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_VERIFY_RETURN_TYPE_SPEC_CONST_UNUSED_HANDLER,
        ZEND_VERIFY_RETURN_TYPE_SPEC_TMP_UNUSED_HANDLER,
        ZEND_VERIFY_RETURN_TYPE_SPEC_VAR_UNUSED_HANDLER,
        ZEND_VERIFY_RETURN_TYPE_SPEC_UNUSED_UNUSED_HANDLER,
        ZEND_VERIFY_RETURN_TYPE_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetFeResetRwHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2003 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FE_RESET_RW_SPEC_CONST_HANDLER,
        ZEND_FE_RESET_RW_SPEC_TMP_HANDLER,
        ZEND_FE_RESET_RW_SPEC_VAR_HANDLER,
        nil,
        ZEND_FE_RESET_RW_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetInitDynamicCallHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2010 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_INIT_DYNAMIC_CALL_SPEC_CONST_HANDLER,
        ZEND_INIT_DYNAMIC_CALL_SPEC_TMPVAR_HANDLER,
        ZEND_INIT_DYNAMIC_CALL_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_INIT_DYNAMIC_CALL_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetDoIcallHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2015 | SPEC_RULE_RETVAL
    offset := vmOffsetBySpec(spec, op)
    handlers := [2]OpcodeHandlerT{
        ZEND_DO_ICALL_SPEC_RETVAL_UNUSED_HANDLER,
        ZEND_DO_ICALL_SPEC_RETVAL_USED_HANDLER,
    }
    return handlers[offset]
}
func vmGetDoUcallHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2017 | SPEC_RULE_RETVAL
    offset := vmOffsetBySpec(spec, op)
    handlers := [2]OpcodeHandlerT{
        ZEND_DO_UCALL_SPEC_RETVAL_UNUSED_HANDLER,
        ZEND_DO_UCALL_SPEC_RETVAL_USED_HANDLER,
    }
    return handlers[offset]
}
func vmGetDoFcallByNameHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2019 | SPEC_RULE_RETVAL
    offset := vmOffsetBySpec(spec, op)
    handlers := [2]OpcodeHandlerT{
        ZEND_DO_FCALL_BY_NAME_SPEC_RETVAL_UNUSED_HANDLER,
        ZEND_DO_FCALL_BY_NAME_SPEC_RETVAL_USED_HANDLER,
    }
    return handlers[offset]
}
func vmGetPreIncObjHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2021 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [0]OpcodeHandlerT{
    }
    return handlers[offset]
}
func vmGetPreDecObjHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2021 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_PRE_INC_OBJ_SPEC_VAR_CONST_HANDLER,
        ZEND_PRE_INC_OBJ_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_PRE_INC_OBJ_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_PRE_INC_OBJ_SPEC_VAR_CV_HANDLER,
        ZEND_PRE_INC_OBJ_SPEC_UNUSED_CONST_HANDLER,
        ZEND_PRE_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_PRE_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_PRE_INC_OBJ_SPEC_UNUSED_CV_HANDLER,
        ZEND_PRE_INC_OBJ_SPEC_CV_CONST_HANDLER,
        ZEND_PRE_INC_OBJ_SPEC_CV_TMPVAR_HANDLER,
        ZEND_PRE_INC_OBJ_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_PRE_INC_OBJ_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetPostIncObjHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2046 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [0]OpcodeHandlerT{
    }
    return handlers[offset]
}
func vmGetPostDecObjHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2046 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_POST_INC_OBJ_SPEC_VAR_CONST_HANDLER,
        ZEND_POST_INC_OBJ_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_POST_INC_OBJ_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_POST_INC_OBJ_SPEC_VAR_CV_HANDLER,
        ZEND_POST_INC_OBJ_SPEC_UNUSED_CONST_HANDLER,
        ZEND_POST_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_POST_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_POST_INC_OBJ_SPEC_UNUSED_CV_HANDLER,
        ZEND_POST_INC_OBJ_SPEC_CV_CONST_HANDLER,
        ZEND_POST_INC_OBJ_SPEC_CV_TMPVAR_HANDLER,
        ZEND_POST_INC_OBJ_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_POST_INC_OBJ_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetEchoHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2071 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_ECHO_SPEC_CONST_HANDLER,
        ZEND_ECHO_SPEC_TMPVAR_HANDLER,
        ZEND_ECHO_SPEC_TMPVAR_HANDLER,
        nil,
        ZEND_ECHO_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetInstanceofHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2077 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_INSTANCEOF_SPEC_TMPVAR_CONST_HANDLER,
        nil,
        ZEND_INSTANCEOF_SPEC_TMPVAR_VAR_HANDLER,
        ZEND_INSTANCEOF_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        ZEND_INSTANCEOF_SPEC_TMPVAR_CONST_HANDLER,
        nil,
        ZEND_INSTANCEOF_SPEC_TMPVAR_VAR_HANDLER,
        ZEND_INSTANCEOF_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_INSTANCEOF_SPEC_CV_CONST_HANDLER,
        nil,
        ZEND_INSTANCEOF_SPEC_CV_VAR_HANDLER,
        ZEND_INSTANCEOF_SPEC_CV_UNUSED_HANDLER,
        nil,
    }
    return handlers[offset]
}
func vmGetMakeRefHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2103 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        nil,
        nil,
        ZEND_MAKE_REF_SPEC_VAR_UNUSED_HANDLER,
        nil,
        ZEND_MAKE_REF_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetIssetIsemptyPropObjHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2115 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_CONST_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_CV_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CV_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CV_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_CONST_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_TMPVAR_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_TMPVAR_HANDLER,
        nil,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_CV_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_CONST_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_TMPVAR_HANDLER,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetJmpSetHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2143 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_JMP_SET_SPEC_CONST_HANDLER,
        ZEND_JMP_SET_SPEC_TMP_HANDLER,
        ZEND_JMP_SET_SPEC_VAR_HANDLER,
        nil,
        ZEND_JMP_SET_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetIssetIsemptyCvHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2149 | SPEC_RULE_ISSET
    offset := vmOffsetBySpec(spec, op)
    handlers := [2]OpcodeHandlerT{
        ZEND_ISSET_ISEMPTY_CV_SPEC_CV_UNUSED_SET_HANDLER,
        ZEND_ISSET_ISEMPTY_CV_SPEC_CV_UNUSED_EMPTY_HANDLER,
    }
    return handlers[offset]
}
func vmGetFetchListWHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2151 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FETCH_LIST_W_SPEC_VAR_CONST_HANDLER,
        ZEND_FETCH_LIST_W_SPEC_VAR_TMPVAR_HANDLER,
        ZEND_FETCH_LIST_W_SPEC_VAR_TMPVAR_HANDLER,
        nil,
        ZEND_FETCH_LIST_W_SPEC_VAR_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetYieldHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2160 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_YIELD_SPEC_CONST_CONST_HANDLER,
        ZEND_YIELD_SPEC_CONST_TMP_HANDLER,
        ZEND_YIELD_SPEC_CONST_VAR_HANDLER,
        ZEND_YIELD_SPEC_CONST_UNUSED_HANDLER,
        ZEND_YIELD_SPEC_CONST_CV_HANDLER,
        ZEND_YIELD_SPEC_TMP_CONST_HANDLER,
        ZEND_YIELD_SPEC_TMP_TMP_HANDLER,
        ZEND_YIELD_SPEC_TMP_VAR_HANDLER,
        ZEND_YIELD_SPEC_TMP_UNUSED_HANDLER,
        ZEND_YIELD_SPEC_TMP_CV_HANDLER,
        ZEND_YIELD_SPEC_VAR_CONST_HANDLER,
        ZEND_YIELD_SPEC_VAR_TMP_HANDLER,
        ZEND_YIELD_SPEC_VAR_VAR_HANDLER,
        ZEND_YIELD_SPEC_VAR_UNUSED_HANDLER,
        ZEND_YIELD_SPEC_VAR_CV_HANDLER,
        ZEND_YIELD_SPEC_UNUSED_CONST_HANDLER,
        ZEND_YIELD_SPEC_UNUSED_TMP_HANDLER,
        ZEND_YIELD_SPEC_UNUSED_VAR_HANDLER,
        ZEND_YIELD_SPEC_UNUSED_UNUSED_HANDLER,
        ZEND_YIELD_SPEC_UNUSED_CV_HANDLER,
        ZEND_YIELD_SPEC_CV_CONST_HANDLER,
        ZEND_YIELD_SPEC_CV_TMP_HANDLER,
        ZEND_YIELD_SPEC_CV_VAR_HANDLER,
        ZEND_YIELD_SPEC_CV_UNUSED_HANDLER,
        ZEND_YIELD_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetGeneratorReturnHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2185 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_GENERATOR_RETURN_SPEC_CONST_HANDLER,
        ZEND_GENERATOR_RETURN_SPEC_TMP_HANDLER,
        ZEND_GENERATOR_RETURN_SPEC_VAR_HANDLER,
        nil,
        ZEND_GENERATOR_RETURN_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetYieldFromHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2194 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_YIELD_FROM_SPEC_CONST_HANDLER,
        ZEND_YIELD_FROM_SPEC_TMP_HANDLER,
        ZEND_YIELD_FROM_SPEC_VAR_HANDLER,
        nil,
        ZEND_YIELD_FROM_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetCoalesceHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2201 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_COALESCE_SPEC_CONST_HANDLER,
        ZEND_COALESCE_SPEC_TMP_HANDLER,
        ZEND_COALESCE_SPEC_VAR_HANDLER,
        nil,
        ZEND_COALESCE_SPEC_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetSpaceshipHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2206 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [25]OpcodeHandlerT{
        ZEND_SPACESHIP_SPEC_CONST_CONST_HANDLER,
        ZEND_SPACESHIP_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_SPACESHIP_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_SPACESHIP_SPEC_CONST_CV_HANDLER,
        ZEND_SPACESHIP_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_SPACESHIP_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_SPACESHIP_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_SPACESHIP_SPEC_TMPVAR_CV_HANDLER,
        ZEND_SPACESHIP_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_SPACESHIP_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_SPACESHIP_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_SPACESHIP_SPEC_TMPVAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_SPACESHIP_SPEC_CV_CONST_HANDLER,
        ZEND_SPACESHIP_SPEC_CV_TMPVAR_HANDLER,
        ZEND_SPACESHIP_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_SPACESHIP_SPEC_CV_CV_HANDLER,
    }
    return handlers[offset]
}
func vmGetFuncGetArgsHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2232 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FUNC_GET_ARGS_SPEC_CONST_UNUSED_HANDLER,
        nil,
        nil,
        ZEND_FUNC_GET_ARGS_SPEC_UNUSED_UNUSED_HANDLER,
        nil,
    }
    return handlers[offset]
}
func vmGetFetchClassConstantHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2245 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_FETCH_CLASS_CONSTANT_SPEC_CONST_CONST_HANDLER,
        nil,
        ZEND_FETCH_CLASS_CONSTANT_SPEC_VAR_CONST_HANDLER,
        ZEND_FETCH_CLASS_CONSTANT_SPEC_UNUSED_CONST_HANDLER,
        nil,
    }
    return handlers[offset]
}
func vmGetSwitchLongHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2255 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_SWITCH_LONG_SPEC_CONST_CONST_HANDLER,
        ZEND_SWITCH_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SWITCH_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        nil,
        ZEND_SWITCH_LONG_SPEC_TMPVARCV_CONST_HANDLER,
    }
    return handlers[offset]
}
func vmGetSwitchStringHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2260 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_SWITCH_STRING_SPEC_CONST_CONST_HANDLER,
        ZEND_SWITCH_STRING_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SWITCH_STRING_SPEC_TMPVARCV_CONST_HANDLER,
        nil,
        ZEND_SWITCH_STRING_SPEC_TMPVARCV_CONST_HANDLER,
    }
    return handlers[offset]
}
func vmGetInArrayHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2265 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_IN_ARRAY_SPEC_CONST_CONST_HANDLER,
        ZEND_IN_ARRAY_SPEC_TMP_CONST_HANDLER,
        ZEND_IN_ARRAY_SPEC_VAR_CONST_HANDLER,
        nil,
        ZEND_IN_ARRAY_SPEC_CV_CONST_HANDLER,
    }
    return handlers[offset]
}
func vmGetCountHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2270 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_COUNT_SPEC_CONST_UNUSED_HANDLER,
        ZEND_COUNT_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_COUNT_SPEC_TMPVAR_UNUSED_HANDLER,
        nil,
        ZEND_COUNT_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetGetClassHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2275 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_GET_CLASS_SPEC_CONST_UNUSED_HANDLER,
        ZEND_GET_CLASS_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_GET_CLASS_SPEC_TMPVAR_UNUSED_HANDLER,
        ZEND_GET_CLASS_SPEC_UNUSED_UNUSED_HANDLER,
        ZEND_GET_CLASS_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetGetTypeHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2281 | SPEC_RULE_OP1
    offset := vmOffsetBySpec(spec, op)
    handlers := [5]OpcodeHandlerT{
        ZEND_GET_TYPE_SPEC_CONST_UNUSED_HANDLER,
        ZEND_GET_TYPE_SPEC_TMP_UNUSED_HANDLER,
        ZEND_GET_TYPE_SPEC_VAR_UNUSED_HANDLER,
        nil,
        ZEND_GET_TYPE_SPEC_CV_UNUSED_HANDLER,
    }
    return handlers[offset]
}
func vmGetArrayKeyExistsHandler(op *ZendOp) OpcodeHandlerT {
    spec := 2286 | SPEC_RULE_OP1 | SPEC_RULE_OP2
    offset := vmOffsetBySpec(spec, op)
    handlers := [917]OpcodeHandlerT{
        ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CONST_HANDLER,
        ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_TMPVAR_HANDLER,
        ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_TMPVAR_HANDLER,
        nil,
        ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CV_HANDLER,
        ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CV_HANDLER,
        ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER,
        ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER,
        nil,
        ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CONST_HANDLER,
        ZEND_ARRAY_KEY_EXISTS_SPEC_CV_TMPVAR_HANDLER,
        ZEND_ARRAY_KEY_EXISTS_SPEC_CV_TMPVAR_HANDLER,
        nil,
        ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CV_HANDLER,
        ZEND_JMP_FORWARD_SPEC_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ADD_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ADD_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_CONST_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_LONG_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_SPEC_CONST_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_LONG_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_SUB_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SUB_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_MUL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_MUL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        nil,
        nil,
        nil,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER,
        ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER,
        ZEND_PRE_INC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_UNUSED_HANDLER,
        ZEND_PRE_INC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_USED_HANDLER,
        ZEND_PRE_INC_LONG_SPEC_CV_RETVAL_UNUSED_HANDLER,
        ZEND_PRE_INC_LONG_SPEC_CV_RETVAL_USED_HANDLER,
        ZEND_PRE_DEC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_UNUSED_HANDLER,
        ZEND_PRE_DEC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_USED_HANDLER,
        ZEND_PRE_DEC_LONG_SPEC_CV_RETVAL_UNUSED_HANDLER,
        ZEND_PRE_DEC_LONG_SPEC_CV_RETVAL_USED_HANDLER,
        ZEND_POST_INC_LONG_NO_OVERFLOW_SPEC_CV_HANDLER,
        ZEND_POST_INC_LONG_SPEC_CV_HANDLER,
        ZEND_POST_DEC_LONG_NO_OVERFLOW_SPEC_CV_HANDLER,
        ZEND_POST_DEC_LONG_SPEC_CV_HANDLER,
        ZEND_QM_ASSIGN_LONG_SPEC_CONST_HANDLER,
        ZEND_QM_ASSIGN_LONG_SPEC_TMPVARCV_HANDLER,
        ZEND_QM_ASSIGN_LONG_SPEC_TMPVARCV_HANDLER,
        nil,
        ZEND_QM_ASSIGN_LONG_SPEC_TMPVARCV_HANDLER,
        ZEND_QM_ASSIGN_DOUBLE_SPEC_CONST_HANDLER,
        ZEND_QM_ASSIGN_DOUBLE_SPEC_TMPVARCV_HANDLER,
        ZEND_QM_ASSIGN_DOUBLE_SPEC_TMPVARCV_HANDLER,
        nil,
        ZEND_QM_ASSIGN_DOUBLE_SPEC_TMPVARCV_HANDLER,
        ZEND_QM_ASSIGN_NOREF_SPEC_CONST_HANDLER,
        ZEND_QM_ASSIGN_NOREF_SPEC_TMPVARCV_HANDLER,
        ZEND_QM_ASSIGN_NOREF_SPEC_TMPVARCV_HANDLER,
        nil,
        ZEND_QM_ASSIGN_NOREF_SPEC_TMPVARCV_HANDLER,
        nil,
        ZEND_FETCH_DIM_R_INDEX_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_FETCH_DIM_R_INDEX_SPEC_CONST_TMPVARCV_HANDLER,
        nil,
        ZEND_FETCH_DIM_R_INDEX_SPEC_CONST_TMPVARCV_HANDLER,
        ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER,
        ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER,
        nil,
        ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER,
        ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_CONST_HANDLER,
        ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER,
        ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER,
        nil,
        ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER,
        nil,
        nil,
        nil,
        nil,
        nil,
        ZEND_FETCH_DIM_R_INDEX_SPEC_CV_CONST_HANDLER,
        ZEND_FETCH_DIM_R_INDEX_SPEC_CV_TMPVARCV_HANDLER,
        ZEND_FETCH_DIM_R_INDEX_SPEC_CV_TMPVARCV_HANDLER,
        nil,
        ZEND_FETCH_DIM_R_INDEX_SPEC_CV_TMPVARCV_HANDLER,
        nil,
        nil,
        ZEND_SEND_VAR_SIMPLE_SPEC_VAR_HANDLER,
        nil,
        ZEND_SEND_VAR_SIMPLE_SPEC_CV_HANDLER,
        nil,
        nil,
        ZEND_SEND_VAR_EX_SIMPLE_SPEC_VAR_HANDLER,
        nil,
        ZEND_SEND_VAR_EX_SIMPLE_SPEC_CV_HANDLER,
        ZEND_SEND_VAL_SIMPLE_SPEC_CONST_HANDLER,
        ZEND_SEND_VAL_EX_SIMPLE_SPEC_CONST_HANDLER,
        ZEND_FE_FETCH_R_SIMPLE_SPEC_VAR_CV_RETVAL_UNUSED_HANDLER,
        ZEND_FE_FETCH_R_SIMPLE_SPEC_VAR_CV_RETVAL_USED_HANDLER,
    }
    return handlers[offset]
}
