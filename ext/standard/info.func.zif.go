package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifPhpversion
var DefZifPhpversion = def.DefFunc("phpversion", 0, 1, []def.ArgInfo{{name: "extension"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	extension := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPhpversion(executeData, returnValue, nil, extension)
})

// generate by ZifPhpcredits
var DefZifPhpcredits = def.DefFunc("phpcredits", 0, 1, []def.ArgInfo{{name: "flag"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	flag := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPhpcredits(executeData, returnValue, nil, flag)
})

// generate by ZifPhpSapiName
var DefZifPhpSapiName = def.DefFunc("php_sapi_name", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifPhpSapiName(executeData, returnValue)
})

// generate by ZifPhpUname
var DefZifPhpUname = def.DefFunc("php_uname", 0, 1, []def.ArgInfo{{name: "mode"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPhpUname(executeData, returnValue, nil, mode)
})

// generate by ZifPhpIniScannedFiles
var DefZifPhpIniScannedFiles = def.DefFunc("php_ini_scanned_files", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifPhpIniScannedFiles(executeData, returnValue)
})

// generate by ZifPhpIniLoadedFile
var DefZifPhpIniLoadedFile = def.DefFunc("php_ini_loaded_file", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifPhpIniLoadedFile(executeData, returnValue)
})
