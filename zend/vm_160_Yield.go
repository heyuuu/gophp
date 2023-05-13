package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_YIELD_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Const1()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
				{

					// generator.GetValue().TryAddRefcount()

				}
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Const1()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)

				// generator.GetValue().TryAddRefcount()

			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var key *types.Zval = opline.Const2()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)

			// generator.GetKey().TryAddRefcount()

		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_CONST_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Const1()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
				{

					// generator.GetValue().TryAddRefcount()

				}
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Const1()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)

				// generator.GetValue().TryAddRefcount()

			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_CONST_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Const1()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
				{

					// generator.GetValue().TryAddRefcount()

				}
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Const1()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)

				// generator.GetValue().TryAddRefcount()

			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = opline.Op2()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Const1()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
				{

					// generator.GetValue().TryAddRefcount()

				}
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Const1()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)

				// generator.GetValue().TryAddRefcount()

			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	/* Consts, temporary variables and references need copying */

	{

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		generator.GetKey().SetLong(generator.GetLargestUsedIntegerKey())
	}
	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Const1()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
				{

					// generator.GetValue().TryAddRefcount()

				}
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Const1()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)

				// generator.GetValue().TryAddRefcount()

			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var key *types.Zval = opline.Cv2OrUndef()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var key *types.Zval = opline.Const2()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)

			// generator.GetKey().TryAddRefcount()

		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_TMP_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_TMP_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = opline.Op2()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	/* Consts, temporary variables and references need copying */

	{

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		generator.GetKey().SetLong(generator.GetLargestUsedIntegerKey())
	}
	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var key *types.Zval = opline.Cv2OrUndef()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Op1()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Op1()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var key *types.Zval = opline.Const2()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)

			// generator.GetKey().TryAddRefcount()

		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_VAR_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Op1()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Op1()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_VAR_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Op1()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Op1()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = opline.Op2()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Op1()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Op1()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	/* Consts, temporary variables and references need copying */

	{

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		generator.GetKey().SetLong(generator.GetLargestUsedIntegerKey())
	}
	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		var free_op1 ZendFreeOp
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Op1()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Op1()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var key *types.Zval = opline.Cv2OrUndef()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* Consts, temporary variables and references need copying */

	/* Consts, temporary variables and references need copying */

	{

		/* If no value was specified yield null */

		generator.GetValue().SetNull()

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	{
		var key *types.Zval = opline.Const2()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)

			// generator.GetKey().TryAddRefcount()

		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_UNUSED_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* Consts, temporary variables and references need copying */

	/* Consts, temporary variables and references need copying */

	{

		/* If no value was specified yield null */

		generator.GetValue().SetNull()

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_UNUSED_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* Consts, temporary variables and references need copying */

	/* Consts, temporary variables and references need copying */

	{

		/* If no value was specified yield null */

		generator.GetValue().SetNull()

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = opline.Op2()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* Consts, temporary variables and references need copying */

	/* Consts, temporary variables and references need copying */

	{

		/* If no value was specified yield null */

		generator.GetValue().SetNull()

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	/* Consts, temporary variables and references need copying */

	{

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		generator.GetKey().SetLong(generator.GetLargestUsedIntegerKey())
	}
	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* If a function call result is yielded and the function did
	 * not return by reference we throw a notice. */

	/* Constants and temporary variables aren't yieldable by reference,
	 * but we still allow them with a notice. */

	/* Consts, temporary variables and references need copying */

	/* Consts, temporary variables and references need copying */

	{

		/* If no value was specified yield null */

		generator.GetValue().SetNull()

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	{
		var key *types.Zval = opline.Cv2OrUndef()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Cv1OrUndef()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Cv1OrUndef()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var key *types.Zval = opline.Const2()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)

			// generator.GetKey().TryAddRefcount()

		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_CV_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Cv1OrUndef()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Cv1OrUndef()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, executeData)

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_CV_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Cv1OrUndef()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Cv1OrUndef()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var free_op2 ZendFreeOp
		var key *types.Zval = opline.Op2()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Cv1OrUndef()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Cv1OrUndef()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	/* Consts, temporary variables and references need copying */

	{

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		generator.GetKey().SetLong(generator.GetLargestUsedIntegerKey())
	}
	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
func ZEND_YIELD_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	if generator.IsForcedClose() {
		return zend_yield_in_closed_generator_helper_SPEC(executeData)
	}

	/* Destroy the previously yielded value */

	// ZvalPtrDtor(generator.GetValue())

	/* Destroy the previously yielded key */

	// ZvalPtrDtor(generator.GetKey())

	/* Set the new yielded value */

	{
		if (executeData.GetFunc().GetOpArray().GetFnFlags() & AccReturnReference) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			{
				var value *types.Zval
				faults.Error(faults.E_NOTICE, "Only variable references should be yielded by reference")
				value = opline.Cv1OrUndef()
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* If a function call result is yielded and the function did
			 * not return by reference we throw a notice. */

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *types.Zval = opline.Cv1OrUndef()

			/* Consts, temporary variables and references need copying */

			{
				types.ZVAL_COPY_VALUE(generator.GetValue(), value)
			}

			/* Consts, temporary variables and references need copying */

		}
	}

	/* If no value was specified yield null */

	/* If no value was specified yield null */

	/* Set the new yielded key */

	{
		var key *types.Zval = opline.Cv2OrUndef()

		/* Consts, temporary variables and references need copying */

		{
			types.ZVAL_COPY_VALUE(generator.GetKey(), key)
		}

		if generator.GetKey().IsLong() && generator.GetKey().Long() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().Long())
		}
	}

	/* If no key was specified we use auto-increment keys */

	if RETURN_VALUE_USED(opline) {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget(opline.Result())
		generator.GetSendTarget().SetNull()
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	ZEND_VM_INC_OPCODE(executeData)

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in executeData so we don't resume at an old position. */

	return -1
}
