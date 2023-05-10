package spl

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func ZmStartupSplExceptions(type_ int, module_number int) int {
	spl_ce_LogicException = zend.SplRegisterSubClass(spl_ce_Exception, "LogicException", nil, nil)
	spl_ce_BadFunctionCallException = zend.SplRegisterSubClass(spl_ce_LogicException, "BadFunctionCallException", nil, nil)
	spl_ce_BadMethodCallException = zend.SplRegisterSubClass(spl_ce_BadFunctionCallException, "BadMethodCallException", nil, nil)
	spl_ce_DomainException = zend.SplRegisterSubClass(spl_ce_LogicException, "DomainException", nil, nil)
	spl_ce_InvalidArgumentException = zend.SplRegisterSubClass(spl_ce_LogicException, "InvalidArgumentException", nil, nil)
	spl_ce_LengthException = zend.SplRegisterSubClass(spl_ce_LogicException, "LengthException", nil, nil)
	spl_ce_OutOfRangeException = zend.SplRegisterSubClass(spl_ce_LogicException, "OutOfRangeException", nil, nil)
	spl_ce_RuntimeException = zend.SplRegisterSubClass(spl_ce_Exception, "RuntimeException", nil, nil)
	spl_ce_OutOfBoundsException = zend.SplRegisterSubClass(spl_ce_RuntimeException, "OutOfBoundsException", nil, nil)
	spl_ce_OverflowException = zend.SplRegisterSubClass(spl_ce_RuntimeException, "OverflowException", nil, nil)
	spl_ce_RangeException = zend.SplRegisterSubClass(spl_ce_RuntimeException, "RangeException", nil, nil)
	spl_ce_UnderflowException = zend.SplRegisterSubClass(spl_ce_RuntimeException, "UnderflowException", nil, nil)
	spl_ce_UnexpectedValueException = zend.SplRegisterSubClass(spl_ce_RuntimeException, "UnexpectedValueException", nil, nil)
	return types.SUCCESS
}
