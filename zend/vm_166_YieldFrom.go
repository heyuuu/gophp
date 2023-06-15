package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_YIELD_FROM_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	var val *types.Zval
	val = opline.Const1()
	if generator.IsForcedClose() {
		faults.ThrowError(nil, "Cannot use \"yield from\" in a force-closed generator")
		UNDEF_RESULT()
		return 0
	}
	if val.IsArray() {
		types.ZVAL_COPY_VALUE(generator.GetValues(), val)
		generator.SetValuesFePos(0)
	} else {
		faults.ThrowError(nil, "Can use \"yield from\" only with arrays and Traversables")
		UNDEF_RESULT()
		return 0
	}

	/* This is the default return value
	 * when the expression is a Generator, it will be overwritten in zend_generator_resume() */

	if RETURN_VALUE_USED(opline) {
		opline.Result().SetNull()
	}

	/* This generator has no send target (though the generator we delegate to might have one) */

	generator.SetSendTarget(nil)

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_FROM_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	var val *types.Zval
	var free_op1 ZendFreeOp
	val = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	if generator.IsForcedClose() {
		faults.ThrowError(nil, "Cannot use \"yield from\" in a force-closed generator")
		UNDEF_RESULT()
		return 0
	}
	if val.IsArray() {
		types.ZVAL_COPY_VALUE(generator.GetValues(), val)
		generator.SetValuesFePos(0)
	} else if val.IsObject() && types.Z_OBJCE_P(val).GetGetIterator() != nil {
		var ce *types.ClassEntry = types.Z_OBJCE_P(val)
		if ce == ZendCeGenerator {
			var new_gen *ZendGenerator = (*ZendGenerator)(val.GetObj())
			if new_gen.GetRetval().IsUndef() {
				if ZendGeneratorGetCurrent(new_gen) == generator {
					faults.ThrowError(nil, "Impossible to yield from the Generator being currently run")
					// ZvalPtrDtor(val)
					UNDEF_RESULT()
					return 0
				} else {
					ZendGeneratorYieldFrom(generator, new_gen)
				}
			} else if new_gen.GetExecuteData() == nil {
				faults.ThrowError(nil, "Generator passed to yield from was aborted without proper return and is unable to continue")
				// ZvalPtrDtor(val)
				UNDEF_RESULT()
				return 0
			} else {
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(opline.Result(), new_gen.GetRetval())
				}
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else {
			var iter *ZendObjectIterator = ce.GetGetIterator()(ce, val, 0)
			// ZvalPtrDtorNogc(free_op1)
			if iter == nil || EG__().GetException() != nil {
				if EG__().GetException() == nil {
					faults.ThrowError(nil, "Object of type %s did not create an Iterator", ce.Name())
				}
				UNDEF_RESULT()
				return 0
			}
			iter.SetIndex(0)
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
				if EG__().GetException() != nil {
					// OBJ_RELEASE(iter.GetStd())
					UNDEF_RESULT()
					return 0
				}
			}
			generator.GetValues().SetObject(iter.GetStd())
		}
	} else {
		faults.ThrowError(nil, "Can use \"yield from\" only with arrays and Traversables")
		// ZvalPtrDtorNogc(free_op1)
		UNDEF_RESULT()
		return 0
	}

	/* This is the default return value
	 * when the expression is a Generator, it will be overwritten in zend_generator_resume() */

	if RETURN_VALUE_USED(opline) {
		opline.Result().SetNull()
	}

	/* This generator has no send target (though the generator we delegate to might have one) */

	generator.SetSendTarget(nil)

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_FROM_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	var val *types.Zval
	var free_op1 ZendFreeOp
	val = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, executeData)
	if generator.IsForcedClose() {
		faults.ThrowError(nil, "Cannot use \"yield from\" in a force-closed generator")
		// ZvalPtrDtorNogc(free_op1)
		UNDEF_RESULT()
		return 0
	}
	if val.IsArray() {
		types.ZVAL_COPY_VALUE(generator.GetValues(), val)
		generator.SetValuesFePos(0)
		// ZvalPtrDtorNogc(free_op1)
	} else if val.IsObject() && types.Z_OBJCE_P(val).GetGetIterator() != nil {
		var ce *types.ClassEntry = types.Z_OBJCE_P(val)
		if ce == ZendCeGenerator {
			var new_gen *ZendGenerator = (*ZendGenerator)(val.GetObj())
			// ZvalPtrDtorNogc(free_op1)
			if new_gen.GetRetval().IsUndef() {
				if ZendGeneratorGetCurrent(new_gen) == generator {
					faults.ThrowError(nil, "Impossible to yield from the Generator being currently run")
					// ZvalPtrDtor(val)
					UNDEF_RESULT()
					return 0
				} else {
					ZendGeneratorYieldFrom(generator, new_gen)
				}
			} else if new_gen.GetExecuteData() == nil {
				faults.ThrowError(nil, "Generator passed to yield from was aborted without proper return and is unable to continue")
				// ZvalPtrDtor(val)
				UNDEF_RESULT()
				return 0
			} else {
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(opline.Result(), new_gen.GetRetval())
				}
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else {
			var iter *ZendObjectIterator = ce.GetGetIterator()(ce, val, 0)
			// ZvalPtrDtorNogc(free_op1)
			if iter == nil || EG__().GetException() != nil {
				if EG__().GetException() == nil {
					faults.ThrowError(nil, "Object of type %s did not create an Iterator", ce.Name())
				}
				UNDEF_RESULT()
				return 0
			}
			iter.SetIndex(0)
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
				if EG__().GetException() != nil {
					// OBJ_RELEASE(iter.GetStd())
					UNDEF_RESULT()
					return 0
				}
			}
			generator.GetValues().SetObject(iter.GetStd())
		}
	} else {
		faults.ThrowError(nil, "Can use \"yield from\" only with arrays and Traversables")
		// ZvalPtrDtorNogc(free_op1)
		UNDEF_RESULT()
		return 0
	}

	/* This is the default return value
	 * when the expression is a Generator, it will be overwritten in zend_generator_resume() */

	if RETURN_VALUE_USED(opline) {
		opline.Result().SetNull()
	}

	/* This generator has no send target (though the generator we delegate to might have one) */

	generator.SetSendTarget(nil)

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_FROM_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	var val *types.Zval
	val = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), executeData)
	if generator.IsForcedClose() {
		faults.ThrowError(nil, "Cannot use \"yield from\" in a force-closed generator")
		UNDEF_RESULT()
		return 0
	}
	if val.IsArray() {
		types.ZVAL_COPY_VALUE(generator.GetValues(), val)
		generator.SetValuesFePos(0)
	} else if val.IsObject() && types.Z_OBJCE_P(val).GetGetIterator() != nil {
		var ce *types.ClassEntry = types.Z_OBJCE_P(val)
		if ce == ZendCeGenerator {
			var new_gen *ZendGenerator = (*ZendGenerator)(val.GetObj())
			if new_gen.GetRetval().IsUndef() {
				if ZendGeneratorGetCurrent(new_gen) == generator {
					faults.ThrowError(nil, "Impossible to yield from the Generator being currently run")
					// ZvalPtrDtor(val)
					UNDEF_RESULT()
					return 0
				} else {
					ZendGeneratorYieldFrom(generator, new_gen)
				}
			} else if new_gen.GetExecuteData() == nil {
				faults.ThrowError(nil, "Generator passed to yield from was aborted without proper return and is unable to continue")
				// ZvalPtrDtor(val)
				UNDEF_RESULT()
				return 0
			} else {
				if RETURN_VALUE_USED(opline) {
					types.ZVAL_COPY(opline.Result(), new_gen.GetRetval())
				}
				return ZEND_VM_NEXT_OPCODE(executeData, opline)
			}
		} else {
			var iter *ZendObjectIterator = ce.GetGetIterator()(ce, val, 0)
			if iter == nil || EG__().GetException() != nil {
				if EG__().GetException() == nil {
					faults.ThrowError(nil, "Object of type %s did not create an Iterator", ce.Name())
				}
				UNDEF_RESULT()
				return 0
			}
			iter.SetIndex(0)
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
				if EG__().GetException() != nil {
					// OBJ_RELEASE(iter.GetStd())
					UNDEF_RESULT()
					return 0
				}
			}
			generator.GetValues().SetObject(iter.GetStd())
		}
	} else {
		faults.ThrowError(nil, "Can use \"yield from\" only with arrays and Traversables")
		UNDEF_RESULT()
		return 0
	}

	/* This is the default return value
	 * when the expression is a Generator, it will be overwritten in zend_generator_resume() */

	if RETURN_VALUE_USED(opline) {
		opline.Result().SetNull()
	}

	/* This generator has no send target (though the generator we delegate to might have one) */

	generator.SetSendTarget(nil)

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
