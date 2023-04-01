package standard

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

var DirGlobals PhpDirGlobals
var DirClassEntryPtr *types.ClassEntry

/* {{{ arginfo */

var PhpDirClassFunctions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("close", 0, ZifClosedir, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("directory"),
		zend.MakeArgName("context"),
	}),
	types.MakeZendFunctionEntryEx("rewind", 0, ZifRewinddir, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("directory"),
		zend.MakeArgName("context"),
	}),
	types.MakeZendFunctionEntryEx("read", 0, PhpIfReaddir, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("directory"),
		zend.MakeArgName("context"),
	}),
}
