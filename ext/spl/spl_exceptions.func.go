// <<generate>>

package spl

import (
	"sik/zend/types"
)

func ZmStartupSplExceptions(type_ int, module_number int) int {
	SplRegisterSubClass(&spl_ce_LogicException, spl_ce_Exception, "LogicException", nil, nil)
	SplRegisterSubClass(&spl_ce_BadFunctionCallException, spl_ce_LogicException, "BadFunctionCallException", nil, nil)
	SplRegisterSubClass(&spl_ce_BadMethodCallException, spl_ce_BadFunctionCallException, "BadMethodCallException", nil, nil)
	SplRegisterSubClass(&spl_ce_DomainException, spl_ce_LogicException, "DomainException", nil, nil)
	SplRegisterSubClass(&spl_ce_InvalidArgumentException, spl_ce_LogicException, "InvalidArgumentException", nil, nil)
	SplRegisterSubClass(&spl_ce_LengthException, spl_ce_LogicException, "LengthException", nil, nil)
	SplRegisterSubClass(&spl_ce_OutOfRangeException, spl_ce_LogicException, "OutOfRangeException", nil, nil)
	SplRegisterSubClass(&spl_ce_RuntimeException, spl_ce_Exception, "RuntimeException", nil, nil)
	SplRegisterSubClass(&spl_ce_OutOfBoundsException, spl_ce_RuntimeException, "OutOfBoundsException", nil, nil)
	SplRegisterSubClass(&spl_ce_OverflowException, spl_ce_RuntimeException, "OverflowException", nil, nil)
	SplRegisterSubClass(&spl_ce_RangeException, spl_ce_RuntimeException, "RangeException", nil, nil)
	SplRegisterSubClass(&spl_ce_UnderflowException, spl_ce_RuntimeException, "UnderflowException", nil, nil)
	SplRegisterSubClass(&spl_ce_UnexpectedValueException, spl_ce_RuntimeException, "UnexpectedValueException", nil, nil)
	return types.SUCCESS
}
