package zend

import "github.com/heyuuu/gophp/zend/types"

var ZendCeGenerator *types.ClassEntry
var zend_ce_ClosedGeneratorException *types.ClassEntry

var ZEND_GENERATOR_CURRENTLY_RUNNING types.ZendUchar = 0x1
var ZEND_GENERATOR_FORCED_CLOSE types.ZendUchar = 0x2
var ZEND_GENERATOR_AT_FIRST_YIELD types.ZendUchar = 0x4
var ZEND_GENERATOR_DO_INIT types.ZendUchar = 0x8

var ZendGeneratorHandlers ZendObjectHandlers

var ZendGeneratorIteratorFunctions ZendObjectIteratorFuncs = MakeZendObjectIteratorFuncs(ZendGeneratorIteratorDtor, ZendGeneratorIteratorValid, ZendGeneratorIteratorGetData, ZendGeneratorIteratorGetKey, ZendGeneratorIteratorMoveForward, ZendGeneratorIteratorRewind, nil)
var GeneratorFunctions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("rewind", AccPublic, zim_Generator_rewind, []ArgInfo{MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", AccPublic, zim_Generator_valid, []ArgInfo{MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", AccPublic, zim_Generator_current, []ArgInfo{MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", AccPublic, zim_Generator_key, []ArgInfo{MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", AccPublic, zim_Generator_next, []ArgInfo{MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("send", AccPublic, zim_Generator_send, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("value"),
	}),
	types.MakeZendFunctionEntryEx("throw", AccPublic, zim_Generator_throw, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("exception"),
	}),
	types.MakeZendFunctionEntryEx("getReturn", AccPublic, zim_Generator_getReturn, []ArgInfo{MakeReturnArgInfo(-1)}),
}
