package spl

import (
	"sik/zend/faults"
	"sik/zend/types"
)

var spl_ce_LogicException *types.ClassEntry
var spl_ce_BadFunctionCallException *types.ClassEntry
var spl_ce_BadMethodCallException *types.ClassEntry
var spl_ce_DomainException *types.ClassEntry
var spl_ce_InvalidArgumentException *types.ClassEntry
var spl_ce_LengthException *types.ClassEntry
var spl_ce_OutOfRangeException *types.ClassEntry
var spl_ce_RuntimeException *types.ClassEntry
var spl_ce_OutOfBoundsException *types.ClassEntry
var spl_ce_OverflowException *types.ClassEntry
var spl_ce_RangeException *types.ClassEntry
var spl_ce_UnderflowException *types.ClassEntry
var spl_ce_UnexpectedValueException *types.ClassEntry

var spl_ce_Exception *types.ClassEntry = faults.ZendCeException
