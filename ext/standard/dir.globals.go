// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
)

var DirGlobals PhpDirGlobals
var DirClassEntryPtr *zend.ZendClassEntry
var PhpDirClassFunctions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"close",
		ZifClosedir,
		ArginfoDir,
		uint32_t(b.SizeOf("arginfo_dir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rewind",
		ZifRewinddir,
		ArginfoDir,
		uint32_t(b.SizeOf("arginfo_dir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"read",
		PhpIfReaddir,
		ArginfoDir,
		uint32_t(b.SizeOf("arginfo_dir")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}
