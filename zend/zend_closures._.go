package zend

import (
	"github.com/heyuuu/gophp/zend/types"
)

var ZendCeClosure *types.ClassEntry

var ClosureHandlers ZendObjectHandlers
var ClosureFunctions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", ZEND_ACC_PRIVATE, zim_Closure___construct, nil),
	types.MakeZendFunctionEntryEx("bind", ZEND_ACC_PUBLIC|ZEND_ACC_STATIC, zim_Closure_bind, []ArgInfo{MakeReturnArgInfo(2),
		MakeArgName("closure"),
		MakeArgName("newthis"),
		MakeArgName("newscope"),
	}),
	types.MakeZendFunctionEntryEx("bindTo", ZEND_ACC_PUBLIC, zim_Closure_bind, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("newthis"),
		MakeArgName("newscope"),
	}),
	types.MakeZendFunctionEntryEx("call", ZEND_ACC_PUBLIC, zim_Closure_call, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("newthis"),
		MakeArgVariadic("parameters"),
	}),
	types.MakeZendFunctionEntryEx("fromCallable", ZEND_ACC_PUBLIC|ZEND_ACC_STATIC, zim_Closure_fromCallable, []ArgInfo{MakeReturnArgInfo(1),
		MakeArgName("callable"),
	}),
}
