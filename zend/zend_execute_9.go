// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendIncludeOrEval(inc_filename *Zval, type_ int) *ZendOpArray {
	var new_op_array *ZendOpArray = nil
	var tmp_inc_filename Zval
	tmp_inc_filename.SetUndef()
	if inc_filename.GetType() != IS_STRING {
		var tmp *ZendString = ZvalTryGetStringFunc(inc_filename)
		if tmp == nil {
			return nil
		}
		tmp_inc_filename.SetString(tmp)
		inc_filename = &tmp_inc_filename
	}
	switch type_ {
	case ZEND_INCLUDE_ONCE:
		fallthrough
	case ZEND_REQUIRE_ONCE:
		var file_handle ZendFileHandle
		var resolved_path *ZendString
		resolved_path = ZendResolvePath(Z_STRVAL_P(inc_filename), Z_STRLEN_P(inc_filename))
		if resolved_path != nil {
			if ZendHashExists(EG__().GetIncludedFiles(), resolved_path) != 0 {
				goto already_compiled
			}
		} else if EG__().GetException() != nil {
			break
		} else if strlen(Z_STRVAL_P(inc_filename)) != Z_STRLEN_P(inc_filename) {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE_ONCE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), Z_STRVAL_P(inc_filename))
			break
		} else {
			resolved_path = inc_filename.GetStr().Copy()
		}
		if SUCCESS == ZendStreamOpen(resolved_path.GetVal(), &file_handle) {
			if file_handle.GetOpenedPath() == nil {
				file_handle.SetOpenedPath(resolved_path.Copy())
			}
			if ZendHashAddEmptyElement(EG__().GetIncludedFiles(), file_handle.GetOpenedPath()) != nil {
				var op_array *ZendOpArray = ZendCompileFile(&file_handle, b.Cond(type_ == ZEND_INCLUDE_ONCE, ZEND_INCLUDE, ZEND_REQUIRE))
				ZendDestroyFileHandle(&file_handle)
				ZendStringReleaseEx(resolved_path, 0)
				if tmp_inc_filename.GetType() != IS_UNDEF {
					ZvalPtrDtorStr(&tmp_inc_filename)
				}
				return op_array
			} else {
				file_handle.Destroy()
			already_compiled:
				new_op_array = ZEND_FAKE_OP_ARRAY
			}
		} else {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE_ONCE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), Z_STRVAL_P(inc_filename))
		}
		ZendStringReleaseEx(resolved_path, 0)
	case ZEND_INCLUDE:
		fallthrough
	case ZEND_REQUIRE:
		if strlen(Z_STRVAL_P(inc_filename)) != Z_STRLEN_P(inc_filename) {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), Z_STRVAL_P(inc_filename))
			break
		}
		new_op_array = CompileFilename(type_, inc_filename)
	case ZEND_EVAL:
		var eval_desc *byte = ZendMakeCompiledStringDescription("eval()'d code")
		new_op_array = ZendCompileString(inc_filename, eval_desc)
		Efree(eval_desc)
	default:

	}
	if tmp_inc_filename.GetType() != IS_UNDEF {
		ZvalPtrDtorStr(&tmp_inc_filename)
	}
	return new_op_array
}
func ZendDoFcallOverloaded(call *ZendExecuteData, ret *Zval) int {
	var fbc *ZendFunction = call.GetFunc()
	var object *ZendObject

	/* Not sure what should be done here if it's a static method */

	if call.GetThis().GetType() != IS_OBJECT {
		ZendVmStackFreeArgs(call)
		if fbc.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
			ZendStringReleaseEx(fbc.GetFunctionName(), 0)
		}
		Efree(fbc)
		ZendVmStackFreeCallFrame(call)
		ZendThrowError(nil, "Cannot call overloaded function for non-object")
		return 0
	}
	object = call.GetThis().GetObj()
	ret.SetNull()
	EG__().SetCurrentExecuteData(call)
	object.GetHandlers().GetCallMethod()(fbc.GetFunctionName(), object, call, ret)
	EG__().SetCurrentExecuteData(call.GetPrevExecuteData())
	ZendVmStackFreeArgs(call)
	if fbc.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
		ZendStringReleaseEx(fbc.GetFunctionName(), 0)
	}
	Efree(fbc)
	return 1
}
func ZendFeResetIterator(array_ptr *Zval, by_ref int, opline *ZendOp, _ EXECUTE_DATA_D) ZendBool {
	var ce *ZendClassEntry = Z_OBJCE_P(array_ptr)
	var iter *ZendObjectIterator = ce.GetGetIterator()(ce, array_ptr, by_ref)
	var is_empty ZendBool
	if iter == nil || EG__().GetException() != nil {
		if iter != nil {
			OBJ_RELEASE(iter.GetStd())
		}
		if EG__().GetException() == nil {
			ZendThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
		}
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		return 1
	}
	iter.SetIndex(0)
	if iter.GetFuncs().GetRewind() != nil {
		iter.GetFuncs().GetRewind()(iter)
		if EG__().GetException() != nil {
			OBJ_RELEASE(iter.GetStd())
			EX_VAR(opline.GetResult().GetVar()).SetUndef()
			return 1
		}
	}
	is_empty = iter.GetFuncs().GetValid()(iter) != SUCCESS
	if EG__().GetException() != nil {
		OBJ_RELEASE(iter.GetStd())
		EX_VAR(opline.GetResult().GetVar()).SetUndef()
		return 1
	}
	iter.SetIndex(-1)
	EX_VAR(opline.GetResult().GetVar()).SetObject(iter.GetStd())
	EX_VAR(opline.GetResult().GetVar()).SetFeIterIdx(uint32 - 1)
	return is_empty
}
func _zendQuickGetConstant(key *Zval, flags uint32, check_defined_only int, opline *ZendOp, _ EXECUTE_DATA_D) int {
	var zv *Zval
	var orig_key *Zval = key
	var c *ZendConstant = nil
	zv = EG__().GetZendConstants().KeyFind(key.GetStr().GetStr())
	if zv != nil {
		c = (*ZendConstant)(zv.GetPtr())
	} else {
		key++
		zv = EG__().GetZendConstants().KeyFind(key.GetStr().GetStr())
		if zv != nil && (ZEND_CONSTANT_FLAGS((*ZendConstant)(zv.GetPtr()))&CONST_CS) == 0 {
			c = (*ZendConstant)(zv.GetPtr())
		} else {
			if (flags & (IS_CONSTANT_IN_NAMESPACE | IS_CONSTANT_UNQUALIFIED)) == (IS_CONSTANT_IN_NAMESPACE | IS_CONSTANT_UNQUALIFIED) {
				key++
				zv = EG__().GetZendConstants().KeyFind(key.GetStr().GetStr())
				if zv != nil {
					c = (*ZendConstant)(zv.GetPtr())
				} else {
					key++
					zv = EG__().GetZendConstants().KeyFind(key.GetStr().GetStr())
					if zv != nil && (ZEND_CONSTANT_FLAGS((*ZendConstant)(zv.GetPtr()))&CONST_CS) == 0 {
						c = (*ZendConstant)(zv.GetPtr())
					}
				}
			}
		}
	}
	if c == nil {
		if check_defined_only == 0 {
			if (opline.GetOp1().GetNum() & IS_CONSTANT_UNQUALIFIED) != 0 {
				var actual *byte = (*byte)(ZendMemrchr(Z_STRVAL_P(RT_CONSTANT(opline, opline.GetOp2())), '\\', Z_STRLEN_P(RT_CONSTANT(opline, opline.GetOp2()))))
				if actual == nil {
					EX_VAR(opline.GetResult().GetVar()).SetStringCopy(RT_CONSTANT(opline, opline.GetOp2()).GetStr())
				} else {
					actual++
					ZVAL_STRINGL(EX_VAR(opline.GetResult().GetVar()), actual, Z_STRLEN_P(RT_CONSTANT(opline, opline.GetOp2()))-(actual-Z_STRVAL_P(RT_CONSTANT(opline, opline.GetOp2()))))
				}

				/* non-qualified constant - allow text substitution */

				ZendError(E_WARNING, "Use of undefined constant %s - assumed '%s' (this will throw an Error in a future version of PHP)", Z_STRVAL_P(EX_VAR(opline.GetResult().GetVar())), Z_STRVAL_P(EX_VAR(opline.GetResult().GetVar())))

				/* non-qualified constant - allow text substitution */

			} else {
				ZendThrowError(nil, "Undefined constant '%s'", Z_STRVAL_P(RT_CONSTANT(opline, opline.GetOp2())))
				EX_VAR(opline.GetResult().GetVar()).SetUndef()
			}
		}
		return FAILURE
	}
	if check_defined_only == 0 {
		ZVAL_COPY_OR_DUP(EX_VAR(opline.GetResult().GetVar()), c.GetValue())
		if (ZEND_CONSTANT_FLAGS(c) & (CONST_CS | CONST_CT_SUBST)) == 0 {
			var ns_sep *byte
			var shortname_offset int
			var shortname_len int
			var is_deprecated ZendBool
			if (flags & IS_CONSTANT_UNQUALIFIED) != 0 {
				var access_key *Zval
				if (flags & IS_CONSTANT_IN_NAMESPACE) == 0 {
					access_key = orig_key - 1
				} else {
					if key < orig_key+2 {
						goto check_short_name
					} else {
						access_key = orig_key + 2
					}
				}
				is_deprecated = !(ZendStringEquals(c.GetName(), access_key.GetStr()))
			} else {
			check_short_name:

				/* Namespaces are always case-insensitive. Only compare shortname. */

				ns_sep = ZendMemrchr(c.GetName().GetVal(), '\\', c.GetName().GetLen())
				if ns_sep != nil {
					shortname_offset = ns_sep - c.GetName().GetVal() + 1
					shortname_len = c.GetName().GetLen() - shortname_offset
				} else {
					shortname_offset = 0
					shortname_len = c.GetName().GetLen()
				}
				is_deprecated = memcmp(c.GetName().GetVal()+shortname_offset, Z_STRVAL_P(orig_key-1)+shortname_offset, shortname_len) != 0
			}
			if is_deprecated != 0 {
				ZendError(E_DEPRECATED, "Case-insensitive constants are deprecated. "+"The correct casing for this constant is \"%s\"", c.GetName().GetVal())
				return SUCCESS
			}
		}
	}
	CACHE_PTR(opline.GetExtendedValue(), c)
	return SUCCESS
}
func ZendQuickGetConstant(key *Zval, flags uint32, opline *ZendOp, _ EXECUTE_DATA_D) {
	_zendQuickGetConstant(key, flags, 0, OPLINE_C, EXECUTE_DATA_C)
}
func ZendQuickCheckConstant(key *Zval, opline *ZendOp, _ EXECUTE_DATA_D) int {
	return _zendQuickGetConstant(key, 0, 1, OPLINE_C, EXECUTE_DATA_C)
}
func ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION() {
	OPLINE = EX(opline) + 1
	ZEND_VM_CONTINUE()
}
func ZEND_VM_NEXT_OPCODE() {
	ZEND_ASSERT(EG__().GetException() == nil)
	OPLINE = opline + 1
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SET_RELATIVE_OPCODE(opline *ZendOp, offset uint32) {
	OPLINE = ZEND_OFFSET_TO_OPLINE(opline, offset)
	ZEND_VM_INTERRUPT_CHECK()
}
func ZEND_VM_JMP_EX(new_op *ZendOp, check_exception int) {
	if check_exception != 0 && EG__().GetException() != nil {
		HANDLE_EXCEPTION()
	}
	OPLINE = new_op
	ZEND_VM_INTERRUPT_CHECK()
	ZEND_VM_CONTINUE()
}
func ZEND_VM_JMP(new_op *ZendOp) { ZEND_VM_JMP_EX(new_op, 1) }
func ZEND_VM_INC_OPCODE() int {
	OPLINE++
	return OPLINE - 1
}
func ZEND_VM_SMART_BRANCH(_result __auto__, _check int) {
	for {
		if _check != 0 && EG__().GetException() != nil {
			break
		}
		if (opline + 1).opcode == ZEND_JMPZ {
			if _result {
				OPLINE = opline + 2
			} else {
				OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
				ZEND_VM_INTERRUPT_CHECK()
			}
		} else if (opline + 1).opcode == ZEND_JMPNZ {
			if !_result {
				OPLINE = opline + 2
			} else {
				OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
				ZEND_VM_INTERRUPT_CHECK()
			}
		} else {
			break
		}
		ZEND_VM_CONTINUE()
		break
	}
}
func ZEND_VM_SMART_BRANCH_JMPZ(_result int, _check int) {
	for {
		if _check != 0 && EG__().GetException() != nil {
			break
		}
		if _result != 0 {
			OPLINE = opline + 2
		} else {
			OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
			ZEND_VM_INTERRUPT_CHECK()
		}
		ZEND_VM_CONTINUE()
		break
	}
}
func ZEND_VM_SMART_BRANCH_JMPNZ(_result int, _check int) {
	for {
		if _check != 0 && EG__().GetException() != nil {
			break
		}
		if _result == 0 {
			OPLINE = opline + 2
		} else {
			OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
			ZEND_VM_INTERRUPT_CHECK()
		}
		ZEND_VM_CONTINUE()
		break
	}
}
func ZEND_VM_SMART_BRANCH_TRUE() {
	if (opline + 1).opcode == ZEND_JMPNZ {
		OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
		ZEND_VM_INTERRUPT_CHECK()
		ZEND_VM_CONTINUE()
	} else if (opline + 1).opcode == ZEND_JMPZ {
		OPLINE = opline + 2
		ZEND_VM_CONTINUE()
	}
}
func ZEND_VM_SMART_BRANCH_TRUE_JMPZ() {
	OPLINE = opline + 2
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SMART_BRANCH_TRUE_JMPNZ() {
	OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
	ZEND_VM_INTERRUPT_CHECK()
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SMART_BRANCH_FALSE() {
	if (opline + 1).opcode == ZEND_JMPNZ {
		OPLINE = opline + 2
		ZEND_VM_CONTINUE()
	} else if (opline + 1).opcode == ZEND_JMPZ {
		OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
		ZEND_VM_INTERRUPT_CHECK()
		ZEND_VM_CONTINUE()
	}
}
func ZEND_VM_SMART_BRANCH_FALSE_JMPZ() {
	OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
	ZEND_VM_INTERRUPT_CHECK()
	ZEND_VM_CONTINUE()
}
func ZEND_VM_SMART_BRANCH_FALSE_JMPNZ() {
	OPLINE = opline + 2
	ZEND_VM_CONTINUE()
}
func UNDEF_RESULT() {
	if (opline.result_type & (IS_VAR | IS_TMP_VAR)) != 0 {
		EX_VAR(opline.result.var_).SetUndef()
	}
}
func ZendSetUserOpcodeHandler(opcode ZendUchar, handler UserOpcodeHandlerT) int {
	if opcode != ZEND_USER_OPCODE {
		if handler == nil {

			/* restore the original handler */

			ZendUserOpcodes[opcode] = opcode

			/* restore the original handler */

		} else {
			ZendUserOpcodes[opcode] = ZEND_USER_OPCODE
		}
		ZendUserOpcodeHandlers[opcode] = handler
		return SUCCESS
	}
	return FAILURE
}
func ZendGetUserOpcodeHandler(opcode ZendUchar) UserOpcodeHandlerT {
	return ZendUserOpcodeHandlers[opcode]
}
func ZendGetZvalPtr(
	opline *ZendOp,
	op_type int,
	node *ZnodeOp,
	execute_data *ZendExecuteData,
	should_free *ZendFreeOp,
	type_ int,
) *Zval {
	var ret *Zval
	switch op_type {
	case IS_CONST:
		ret = RT_CONSTANT(opline, *node)
		*should_free = nil
	case IS_TMP_VAR:
		fallthrough
	case IS_VAR:
		ret = EX_VAR(node.GetVar())
		*should_free = ret
	case IS_CV:
		ret = EX_VAR(node.GetVar())
		*should_free = nil
	default:
		ret = nil
		*should_free = ret
	}
	return ret
}
