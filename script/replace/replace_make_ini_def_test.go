package replace

import (
	"fmt"
	"sik/script/util"
	"testing"
)

func Test_replaceMakeIniEntryDef(t *testing.T) {
	files := []string{
		"/Users/heyu/Code/sik/sik-go-gen-2/zend/zend._.go",
		"/Users/heyu/Code/sik/sik-go-gen-2/ext/standard/assert._.go",
		"/Users/heyu/Code/sik/sik-go-gen-2/core/main._.go",
	}
	for _, file := range files {
		code := util.MustReadFileString(file)
		result := replaceMakeIniEntryDef(code)
		if result != code {
			util.MustWriteFileString(file, result)
		}
	}
}

func Test_replaceMakeIniEntryDef_Dev(t *testing.T) {
	code := `
var IniEntries = []ZendIniEntryDef{
	MakeZendIniEntryDef("error_reporting", OnUpdateErrorReporting, nil, nil, nil, nil, nil, ZEND_INI_ALL),
	MakeZendIniEntryDef("zend.assertions", OnUpdateAssertions, any(zend_long((*byte)(&((*ZendExecutorGlobals)(nil).GetAssertions()))-(*byte)(nil))), any(&ExecutorGlobals), nil, "1", nil, ZEND_INI_ALL),
	MakeZendIniEntryDef("zend.enable_gc", OnUpdateGCEnabled, nil, nil, nil, "1", ZendGcEnabledDisplayerCb, ZEND_INI_ALL),
	MakeZendIniEntryDef("zend.multibyte", OnUpdateBool, any(zend_long((*byte)(&((*ZendCompilerGlobals)(nil).GetMultibyte()))-(*byte)(nil))), any(&CompilerGlobals), nil, "0", ZendIniBooleanDisplayerCb, ZEND_INI_PERDIR),
	MakeZendIniEntryDef("zend.script_encoding", OnUpdateScriptEncoding, nil, nil, nil, nil, nil, ZEND_INI_ALL),
	MakeZendIniEntryDef("zend.detect_unicode", OnUpdateBool, any(zend_long((*byte)(&((*ZendCompilerGlobals)(nil).GetDetectUnicode()))-(*byte)(nil))), any(&CompilerGlobals), nil, "1", ZendIniBooleanDisplayerCb, ZEND_INI_ALL),
	MakeZendIniEntryDef("zend.signal_check", OnUpdateBool, any(zend_long((*byte)(&((*ZendSignalGlobalsT)(nil).GetCheck()))-(*byte)(nil))), any(&ZendSignalGlobals), nil, SIGNAL_CHECK_DEFAULT, ZendIniBooleanDisplayerCb, ZEND_INI_SYSTEM),
	MakeZendIniEntryDef("zend.exception_ignore_args", OnUpdateBool, any(zend_long((*byte)(&((*ZendExecutorGlobals)(nil).GetExceptionIgnoreArgs()))-(*byte)(nil))), any(&ExecutorGlobals), nil, "0", ZendIniBooleanDisplayerCb, ZEND_INI_ALL),
}
`
	result := replaceMakeIniEntryDef(code)
	fmt.Println(result)
}
