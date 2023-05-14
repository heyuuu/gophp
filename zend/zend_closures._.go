package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

var ZendCeClosure *types.ClassEntry

var ClosureHandlers types.ObjectHandlers
var ClosureFunctions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", types.AccPrivate, zim_Closure___construct, nil),
	types.MakeZendFunctionEntryEx("bind", types.AccPublic|types.AccStatic, zim_Closure_bind, []ArgInfo{MakeReturnArgInfo(2),
		MakeArgName("closure"),
		MakeArgName("newthis"),
		MakeArgName("newscope"),
	}),
	types.MakeZendFunctionEntryEx("bindTo", types.AccPublic, zim_Closure_bind, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("newthis"),
		MakeArgName("newscope"),
	}),
	types.MakeZendFunctionEntryEx("call", types.AccPublic, zim_Closure_call, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("newthis"),
		MakeArgVariadic("parameters"),
	}),
	types.MakeZendFunctionEntryEx("fromCallable", types.AccPublic|types.AccStatic, zim_Closure_fromCallable, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("callable"),
	}),
}
