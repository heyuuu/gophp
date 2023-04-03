package zend

func ZEND_GENERATOR_RETURN_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval *types.Zval
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	retval = RT_CONSTANT(opline, opline.GetOp1())

	/* Copy return value into generator->retval */

	{
		types.ZVAL_COPY_VALUE(generator.GetRetval(), retval)
		{

			generator.GetRetval().TryAddRefcount()

		}
	}

	/* Close the generator to free up resources */

	ZendGeneratorClose(generator, 1)

	/* Pass execution back to handling code */

	return -1

	/* Pass execution back to handling code */
}
func ZEND_GENERATOR_RETURN_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval *types.Zval
	var free_op1 ZendFreeOp
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	retval = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)

	/* Copy return value into generator->retval */

	{
		types.ZVAL_COPY_VALUE(generator.GetRetval(), retval)
	}

	/* Close the generator to free up resources */

	ZendGeneratorClose(generator, 1)

	/* Pass execution back to handling code */

	return -1

	/* Pass execution back to handling code */
}
func ZEND_GENERATOR_RETURN_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval *types.Zval
	var free_op1 ZendFreeOp
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	retval = opline.Op1Ptr(&free_op1)

	/* Copy return value into generator->retval */

	{
		types.ZVAL_COPY_VALUE(generator.GetRetval(), retval)
	}

	/* Close the generator to free up resources */

	ZendGeneratorClose(generator, 1)

	/* Pass execution back to handling code */

	return -1

	/* Pass execution back to handling code */
}
func ZEND_GENERATOR_RETURN_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var retval *types.Zval
	var generator *ZendGenerator = ZendGetRunningGenerator(executeData)
	retval = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), executeData)

	/* Copy return value into generator->retval */

	{
		types.ZVAL_COPY_VALUE(generator.GetRetval(), retval)
	}

	/* Close the generator to free up resources */

	ZendGeneratorClose(generator, 1)

	/* Pass execution back to handling code */

	return -1

	/* Pass execution back to handling code */
}
