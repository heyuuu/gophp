package spl

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func ZmStartupSplExceptions(type_ int, module_number int) int {
	spl_ce_LogicException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "LogicException",
		Parent:       spl_ce_Exception,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_BadFunctionCallException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "BadFunctionCallException",
		Parent:       spl_ce_LogicException,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_BadMethodCallException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "BadMethodCallException",
		Parent:       spl_ce_BadFunctionCallException,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_DomainException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "DomainException",
		Parent:       spl_ce_LogicException,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_InvalidArgumentException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "InvalidArgumentException",
		Parent:       spl_ce_LogicException,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_LengthException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "LengthException",
		Parent:       spl_ce_LogicException,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_OutOfRangeException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "OutOfRangeException",
		Parent:       spl_ce_LogicException,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_RuntimeException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "RuntimeException",
		Parent:       spl_ce_Exception,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_OutOfBoundsException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "OutOfBoundsException",
		Parent:       spl_ce_RuntimeException,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_OverflowException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "OverflowException",
		Parent:       spl_ce_RuntimeException,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_RangeException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "RangeException",
		Parent:       spl_ce_RuntimeException,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_UnderflowException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "UnderflowException",
		Parent:       spl_ce_RuntimeException,
		Functions:    nil,
		CreateObject: nil,
	})
	spl_ce_UnexpectedValueException = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "UnexpectedValueException",
		Parent:       spl_ce_RuntimeException,
		Functions:    nil,
		CreateObject: nil,
	})
	return types.SUCCESS
}
