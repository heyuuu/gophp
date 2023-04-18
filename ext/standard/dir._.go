package standard

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var DirGlobals PhpDirGlobals
var DirClassEntryPtr *types2.ClassEntry

/* {{{ arginfo */

var PhpDirClassFunctions []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("close", 0, ZifClosedir, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("directory"),
		zend.MakeArgName("context"),
	}),
	types2.MakeZendFunctionEntryEx("rewind", 0, ZifRewinddir, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("directory"),
		zend.MakeArgName("context"),
	}),
	types2.MakeZendFunctionEntryEx("read", 0, PhpIfReaddir, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("directory"),
		zend.MakeArgName("context"),
	}),
}
