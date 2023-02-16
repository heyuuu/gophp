package replace

import (
	"sik/script/util"
	"sik/script/util/finder"
	"testing"
)

func Test_replaceMakeFunctionEntryDef(t *testing.T) {
	f := finder.DefaultProjectFinder()
	f.Walk(func(file finder.FileInfo) {
		code := util.MustReadFileString(file.Path)
		result := replaceMakeFunctionEntryDef(code)
		if result != code {
			util.MustWriteFileString(file.Path, result)
		}
	})
}

func Test_replaceMakeFunctionEntryDef_Dev(t *testing.T) {
	code := `
var BuiltinFunctions []ZendFunctionEntry = []ZendFunctionEntry{
	MakeZendFunctionEntry("zend_version", ZifZendVersion, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("gc_disable", ZifGcDisable, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry("gc_status", ZifGcStatus, ArginfoZendVoid, uint32(b.SizeOf("arginfo_zend__void")/b.SizeOf("struct _zend_internal_arg_info")-1), 0),
	MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
`
	result := replaceMakeFunctionEntryDef(code)
	println(result)
}
