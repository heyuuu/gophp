package replace

import "testing"

func Test_replaceMakeFunctionEntryDef(t *testing.T) {
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
