package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

var ZendCeGenerator *types2.ClassEntry
var zend_ce_ClosedGeneratorException *types2.ClassEntry

var ZEND_GENERATOR_CURRENTLY_RUNNING types2.ZendUchar = 0x1
var ZEND_GENERATOR_FORCED_CLOSE types2.ZendUchar = 0x2
var ZEND_GENERATOR_AT_FIRST_YIELD types2.ZendUchar = 0x4
var ZEND_GENERATOR_DO_INIT types2.ZendUchar = 0x8

var ZendGeneratorHandlers ZendObjectHandlers

var ZendGeneratorIteratorFunctions ZendObjectIteratorFuncs = MakeZendObjectIteratorFuncs(ZendGeneratorIteratorDtor, ZendGeneratorIteratorValid, ZendGeneratorIteratorGetData, ZendGeneratorIteratorGetKey, ZendGeneratorIteratorMoveForward, ZendGeneratorIteratorRewind, nil)
var GeneratorFunctions []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("rewind", AccPublic, zim_Generator_rewind, []ArgInfo{MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("valid", AccPublic, zim_Generator_valid, []ArgInfo{MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("current", AccPublic, zim_Generator_current, []ArgInfo{MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("key", AccPublic, zim_Generator_key, []ArgInfo{MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("next", AccPublic, zim_Generator_next, []ArgInfo{MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("send", AccPublic, zim_Generator_send, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("value"),
	}),
	types2.MakeZendFunctionEntryEx("throw", AccPublic, zim_Generator_throw, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("exception"),
	}),
	types2.MakeZendFunctionEntryEx("getReturn", AccPublic, zim_Generator_getReturn, []ArgInfo{MakeReturnArgInfo(-1)}),
}
