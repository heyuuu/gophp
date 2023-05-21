package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendIncludeOrEval(inc_filename *types.Zval, type_ int) *types.ZendOpArray {
	var new_op_array *types.ZendOpArray = nil
	var tmp_inc_filename types.Zval
	tmp_inc_filename.SetUndef()
	if inc_filename.GetType() != types.IS_STRING {
		var tmp *types.String = operators.ZvalTryGetString(inc_filename)
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
		var resolved_path *string
		resolved_path = core.PhpResolvePathForZend(inc_filename.String().GetStr())
		if resolved_path != nil {
			if EG__().GetIncludedFiles().KeyExists(*resolved_path) {
				goto already_compiled
			}
		} else if EG__().GetException() != nil {
			break
		} else if strlen(inc_filename.String().GetVal()) != inc_filename.String().GetLen() {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE_ONCE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), inc_filename.String().GetVal())
			break
		} else {
			*resolved_path = inc_filename.String().GetStr()
		}

		if fh := NewFileHandleByOpenStream(*resolved_path); fh != nil {
			if fh.GetOpenedPath() == "" {
				fh.SetOpenedPath(*resolved_path)
			}
			if types.ZendHashAddEmptyElement(EG__().GetIncludedFiles(), fh.GetOpenedPath().GetStr()) != nil {
				var op_array *types.ZendOpArray = CompileFile(&fh, b.Cond(type_ == ZEND_INCLUDE_ONCE, ZEND_INCLUDE, ZEND_REQUIRE))
				ZendDestroyFileHandle(&fh)
				if tmp_inc_filename.IsNotUndef() {

				}
				return op_array
			} else {
				fh.Destroy()
			already_compiled:
				new_op_array = ZEND_FAKE_OP_ARRAY
			}
		} else {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE_ONCE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), inc_filename.String().GetVal())
		}
	case ZEND_INCLUDE:
		fallthrough
	case ZEND_REQUIRE:
		if strlen(inc_filename.String().GetVal()) != inc_filename.String().GetLen() {
			ZendMessageDispatcher(b.Cond(type_ == ZEND_INCLUDE, ZMSG_FAILED_INCLUDE_FOPEN, ZMSG_FAILED_REQUIRE_FOPEN), inc_filename.String().GetVal())
			break
		}
		new_op_array = CompileFilename(type_, inc_filename)
	case ZEND_EVAL:
		var eval_desc *byte = ZendMakeCompiledStringDescription("eval()'d code")
		new_op_array = CompileString(inc_filename, eval_desc)
		Efree(eval_desc)
	default:

	}
	if tmp_inc_filename.IsNotUndef() {

	}
	return new_op_array
}
func ZendDoFcallOverloaded(call *ZendExecuteData, ret *types.Zval) int {
	var fbc types.IFunction = call.GetFunc()
	var object *types.ZendObject

	/* Not sure what should be done here if it's a static method */

	if call.GetThis().GetType() != types.IS_OBJECT {
		ZendVmStackFreeArgs(call)
		if fbc.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
			// types.ZendStringReleaseEx(fbc.GetFunctionName(), 0)
		}
		Efree(fbc)
		ZendVmStackFreeCallFrame(call)
		faults.ThrowError(nil, "Cannot call overloaded function for non-object")
		return 0
	}
	object = call.GetThis().Object()
	ret.SetNull()
	EG__().SetCurrentExecuteData(call)
	object.CallMethod(fbc.GetFunctionName(), object, call, ret)
	EG__().SetCurrentExecuteData(call.GetPrevExecuteData())
	ZendVmStackFreeArgs(call)
	if fbc.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
		// types.ZendStringReleaseEx(fbc.GetFunctionName(), 0)
	}
	Efree(fbc)
	return 1
}
func ZendFeResetIterator(array_ptr *types.Zval, by_ref int, opline *ZendOp, executeData *ZendExecuteData) types.ZendBool {
	var ce *types.ClassEntry = types.Z_OBJCE_P(array_ptr)
	var iter *ZendObjectIterator = ce.GetGetIterator()(ce, array_ptr, by_ref)
	var is_empty types.ZendBool
	if iter == nil || EG__().GetException() != nil {
		if iter != nil {
			// OBJ_RELEASE(iter.GetStd())
		}
		if EG__().GetException() == nil {
			faults.ThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ce.Name())
		}
		opline.Result().SetUndef()
		return 1
	}
	iter.SetIndex(0)
	if iter.GetFuncs().GetRewind() != nil {
		iter.GetFuncs().GetRewind()(iter)
		if EG__().GetException() != nil {
			// OBJ_RELEASE(iter.GetStd())
			opline.Result().SetUndef()
			return 1
		}
	}
	is_empty = iter.GetFuncs().GetValid()(iter) != types.SUCCESS
	if EG__().GetException() != nil {
		// OBJ_RELEASE(iter.GetStd())
		opline.Result().SetUndef()
		return 1
	}
	iter.SetIndex(-1)
	opline.Result().SetObject(iter.GetStd())
	opline.Result().SetFeIterIdx(uint32 - 1)
	return is_empty
}
func _zendQuickGetConstant(key *types.Zval, flags uint32, check_defined_only int, opline *ZendOp, executeData *ZendExecuteData) int {
	var orig_key *types.Zval = key
	var c *ZendConstant = nil
	c = EG__().ConstantTable().Get(key.String().GetStr())
	if c == nil {
		key++
		c = EG__().ConstantTable().Get(key.String().GetStr())
		if c == nil || c.IsCaseSensitive() {
			if (flags & (IS_CONSTANT_IN_NAMESPACE | IS_CONSTANT_UNQUALIFIED)) == (IS_CONSTANT_IN_NAMESPACE | IS_CONSTANT_UNQUALIFIED) {
				key++
				c = EG__().ConstantTable().Get(key.String().GetStr())
				if c == nil {
					key++
					c = EG__().ConstantTable().Get(key.String().GetStr())
					if c != nil && c.IsCaseSensitive() {
						c = nil
					}
				}
			}
		}
	}
	if c == nil {
		if check_defined_only == 0 {
			if (opline.GetOp1().GetNum() & IS_CONSTANT_UNQUALIFIED) != 0 {
				var actual *byte = (*byte)(operators.ZendMemrchr(opline.Const2().String().GetVal(), '\\', opline.Const2().String().GetLen()))
				if actual == nil {
					opline.Result().SetStringCopy(opline.Const2().String())
				} else {
					actual++
					opline.Result().SetStringVal(b.CastStr(actual, opline.Const2().String().GetLen()-(actual-opline.Const2().String().GetVal())))
				}

				faults.Error(faults.E_WARNING, "Use of undefined constant %s - assumed '%s' (this will throw an Error in a future version of PHP)", opline.Result().String().GetVal(), opline.Result().String().GetVal())

			} else {
				faults.ThrowError(nil, "Undefined constant '%s'", opline.Const2().String().GetVal())
				opline.Result().SetUndef()
			}
		}
		return types.FAILURE
	}
	if check_defined_only == 0 {
		types.ZVAL_COPY_OR_DUP(opline.Result(), c.Value())
		if !c.IsCaseSensitive() && !c.IsCtSubst() {
			var ns_sep *byte
			var shortname_offset int
			var shortname_len int
			var is_deprecated types.ZendBool
			if (flags & IS_CONSTANT_UNQUALIFIED) != 0 {
				var access_key *types.Zval
				if (flags & IS_CONSTANT_IN_NAMESPACE) == 0 {
					access_key = orig_key - 1
				} else {
					if key < orig_key+2 {
						goto check_short_name
					} else {
						access_key = orig_key + 2
					}
				}
				is_deprecated = types.IntBool(c.GetName().GetStr() != access_key.StringVal())
			} else {
			check_short_name:

				/* Namespaces are always case-insensitive. Only compare shortname. */

				ns_sep = operators.ZendMemrchr(c.GetName().GetVal(), '\\', c.GetName().GetLen())
				if ns_sep != nil {
					shortname_offset = ns_sep - c.GetName().GetVal() + 1
					shortname_len = c.GetName().GetLen() - shortname_offset
				} else {
					shortname_offset = 0
					shortname_len = c.GetName().GetLen()
				}
				is_deprecated = memcmp(c.GetName().GetVal()+shortname_offset, (orig_key-1).GetStr().GetVal()+shortname_offset, shortname_len) != 0
			}
			if is_deprecated != 0 {
				faults.Error(faults.E_DEPRECATED, "Case-insensitive constants are deprecated. "+"The correct casing for this constant is \"%s\"", c.GetName().GetVal())
				return types.SUCCESS
			}
		}
	}
	CACHE_PTR(opline.GetExtendedValue(), c)
	return types.SUCCESS
}
func ZendQuickGetConstant(key *types.Zval, flags uint32, opline *ZendOp, executeData *ZendExecuteData) {
	_zendQuickGetConstant(key, flags, 0, opline, executeData)
}
func ZendQuickCheckConstant(key *types.Zval, opline *ZendOp, executeData *ZendExecuteData) int {
	return _zendQuickGetConstant(key, 0, 1, opline, executeData)
}
func ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData *ZendExecuteData) int {
	OPLINE = executeData.GetOpline() + 1
	return 0
}
func ZEND_VM_NEXT_OPCODE(executeData *ZendExecuteData, opline *ZendOp) int {
	b.Assert(EG__().GetException() == nil)
	OPLINE = opline + 1
	return 0
}
func ZEND_VM_SET_RELATIVE_OPCODE(executeData *ZendExecuteData, opline *ZendOp, offset uint32) {
	OPLINE = ZEND_OFFSET_TO_OPLINE(opline, offset)
	ZEND_VM_INTERRUPT_CHECK(executeData)
}
func ZEND_VM_JMP_EX(executeData *ZendExecuteData, new_op *ZendOp, check_exception int) int {
	if check_exception != 0 && EG__().GetException() != nil {
		return 0
	}
	OPLINE = new_op
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_VM_JMP(executeData *ZendExecuteData, new_op *ZendOp) {
	ZEND_VM_JMP_EX(executeData, new_op, 1)
}
func ZEND_VM_INC_OPCODE(executeData *ZendExecuteData) int {
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
				ZEND_VM_INTERRUPT_CHECK(executeData)
			}
		} else if (opline + 1).opcode == ZEND_JMPNZ {
			if !_result {
				OPLINE = opline + 2
			} else {
				OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
				ZEND_VM_INTERRUPT_CHECK(executeData)
			}
		} else {
			break
		}
		return 0
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
			ZEND_VM_INTERRUPT_CHECK(executeData)
		}
		return 0
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
			ZEND_VM_INTERRUPT_CHECK(executeData)
		}
		return 0
		break
	}
}
func ZEND_VM_SMART_BRANCH_TRUE() {
	if (opline + 1).opcode == ZEND_JMPNZ {
		OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
		ZEND_VM_INTERRUPT_CHECK(executeData)
		return 0
	} else if (opline + 1).opcode == ZEND_JMPZ {
		OPLINE = opline + 2
		return 0
	}
}
func ZEND_VM_SMART_BRANCH_TRUE_JMPZ() {
	OPLINE = opline + 2
	return 0
}
func ZEND_VM_SMART_BRANCH_TRUE_JMPNZ() {
	OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_VM_SMART_BRANCH_FALSE() {
	if (opline + 1).opcode == ZEND_JMPNZ {
		OPLINE = opline + 2
		return 0
	} else if (opline + 1).opcode == ZEND_JMPZ {
		OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
		ZEND_VM_INTERRUPT_CHECK(executeData)
		return 0
	}
}
func ZEND_VM_SMART_BRANCH_FALSE_JMPZ() {
	OPLINE = OP_JMP_ADDR(opline+1, (opline + 1).op2)
	ZEND_VM_INTERRUPT_CHECK(executeData)
	return 0
}
func ZEND_VM_SMART_BRANCH_FALSE_JMPNZ() {
	OPLINE = opline + 2
	return 0
}
func UNDEF_RESULT() {
	if (opline.result_type & (IS_VAR | IS_TMP_VAR)) != 0 {
		EX_VAR(opline.result.var_).SetUndef()
	}
}
