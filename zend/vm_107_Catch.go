package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZEND_CATCH_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var ce *types.ClassEntry
	var catch_ce *types.ClassEntry
	var exception *types.ZendObject
	var ex *types.Zval

	/* Check whether an exception has been thrown, if not, jump over code */

	faults.ExceptionRestore()
	if EG__().GetException() == nil {
		return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
	}
	catch_ce = CACHED_PTR(opline.GetExtendedValue() & ^ZEND_LAST_CATCH)
	if catch_ce == nil {
		catch_ce = ZendFetchClassByName(opline.Const1().String(), (opline.Const1() + 1).GetStr(), ZEND_FETCH_CLASS_NO_AUTOLOAD)
		CACHE_PTR(opline.GetExtendedValue() & ^ZEND_LAST_CATCH, catch_ce)
	}
	ce = EG__().GetException().GetCe()
	if ce != catch_ce {
		if catch_ce == nil || operators.InstanceofFunction(ce, catch_ce) == 0 {
			if (opline.GetExtendedValue() & ZEND_LAST_CATCH) != 0 {
				faults.RethrowException(executeData)
				return 0
			}
			return ZEND_VM_JMP_EX(executeData, OP_JMP_ADDR(opline, opline.GetOp2()), 0)
		}
	}
	exception = EG__().GetException()
	ex = opline.Result()

	/* Always perform a strict assignment. There is a reasonable expectation that if you
	 * write "catch (Exception $e)" then $e will actually be instanceof Exception. As such,
	 * we should not permit coercion to string here. */

	var tmp types.Zval
	tmp.SetObject(exception)
	EG__().SetException(nil)
	ZendAssignToVariable(ex, &tmp, 1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
