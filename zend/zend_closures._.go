package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

var ZendCeClosure *types2.ClassEntry

var ClosureHandlers ZendObjectHandlers
var ClosureFunctions []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", AccPrivate, zim_Closure___construct, nil),
	types2.MakeZendFunctionEntryEx("bind", AccPublic|AccStatic, zim_Closure_bind, []ArgInfo{MakeReturnArgInfo(2),
		MakeArgName("closure"),
		MakeArgName("newthis"),
		MakeArgName("newscope"),
	}),
	types2.MakeZendFunctionEntryEx("bindTo", AccPublic, zim_Closure_bind, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("newthis"),
		MakeArgName("newscope"),
	}),
	types2.MakeZendFunctionEntryEx("call", AccPublic, zim_Closure_call, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("newthis"),
		MakeArgVariadic("parameters"),
	}),
	types2.MakeZendFunctionEntryEx("fromCallable", AccPublic|AccStatic, zim_Closure_fromCallable, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("callable"),
	}),
}
