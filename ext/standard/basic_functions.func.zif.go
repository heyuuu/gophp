package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifGetenv
var DefZifGetenv = def.DefFunc("getenv", 0, 2, []def.ArgInfo{{name: "varname"}, {name: "local_only"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	varname := fp.ParseZval()
	local_only := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetenv(executeData, returnValue, nil, varname, local_only)
})

// generate by ZifGetopt
var DefZifGetopt = def.DefFunc("getopt", 1, 3, []def.ArgInfo{{name: "options"}, {name: "opts"}, {name: "optind"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	options := fp.ParseZval()
	fp.StartOptional()
	opts := fp.ParseZval()
	optind := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifGetopt(executeData, returnValue, options, nil, opts, optind)
})

// generate by ZifErrorLog
var DefZifErrorLog = def.DefFunc("error_log", 1, 4, []def.ArgInfo{{name: "message"}, {name: "message_type"}, {name: "destination"}, {name: "extra_headers"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	message := fp.ParseZval()
	fp.StartOptional()
	message_type := fp.ParseZval()
	destination := fp.ParseZval()
	extra_headers := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifErrorLog(executeData, returnValue, message, nil, message_type, destination, extra_headers)
})

// generate by ZifErrorGetLast
var DefZifErrorGetLast = def.DefFunc("error_get_last", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifErrorGetLast(executeData, returnValue)
})

// generate by ZifErrorClearLast
var DefZifErrorClearLast = def.DefFunc("error_clear_last", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifErrorClearLast(executeData, returnValue)
})

// generate by ZifCallUserFunc
var DefZifCallUserFunc = def.DefFunc("call_user_func", 1, -1, []def.ArgInfo{{name: "function_name"}, {name: "parameters"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	function_name := fp.ParseZval()
	fp.StartOptional()
	parameters := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifCallUserFunc(executeData, returnValue, function_name, nil, parameters)
})

// generate by ZifCallUserFuncArray
var DefZifCallUserFuncArray = def.DefFunc("call_user_func_array", 2, 2, []def.ArgInfo{{name: "function_name"}, {name: "parameters"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	function_name := fp.ParseZval()
	parameters := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCallUserFuncArray(executeData, returnValue, function_name, parameters)
})

// generate by ZifForwardStaticCall
var DefZifForwardStaticCall = def.DefFunc("forward_static_call", 1, -1, []def.ArgInfo{{name: "function_name"}, {name: "parameters"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	function_name := fp.ParseZval()
	fp.StartOptional()
	parameters := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifForwardStaticCall(executeData, returnValue, function_name, nil, parameters)
})

// generate by ZifForwardStaticCallArray
var DefZifForwardStaticCallArray = def.DefFunc("forward_static_call_array", 2, 2, []def.ArgInfo{{name: "function_name"}, {name: "parameters"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	function_name := fp.ParseZval()
	parameters := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifForwardStaticCallArray(executeData, returnValue, function_name, parameters)
})

// generate by ZifRegisterShutdownFunction
var DefZifRegisterShutdownFunction = def.DefFunc("register_shutdown_function", 1, -1, []def.ArgInfo{{name: "function_name"}, {name: "parameters"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	function_name := fp.ParseZval()
	fp.StartOptional()
	parameters := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifRegisterShutdownFunction(executeData, returnValue, function_name, nil, parameters)
})

// generate by ZifHighlightFile
var DefZifHighlightFile = def.DefFunc("highlight_file", 1, 2, []def.ArgInfo{{name: "file_name"}, {name: "return_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	file_name := fp.ParseZval()
	fp.StartOptional()
	return_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHighlightFile(executeData, returnValue, file_name, nil, return_)
})

// generate by ZifHighlightString
var DefZifHighlightString = def.DefFunc("highlight_string", 1, 2, []def.ArgInfo{{name: "string"}, {name: "return_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	string := fp.ParseZval()
	fp.StartOptional()
	return_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHighlightString(executeData, returnValue, string, nil, return_)
})

// generate by ZifIniGetAll
var DefZifIniGetAll = def.DefFunc("ini_get_all", 0, 2, []def.ArgInfo{{name: "extension"}, {name: "details"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	extension := fp.ParseZval()
	details := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIniGetAll(executeData, returnValue, nil, extension, details)
})

// generate by ZifPrintR
var DefZifPrintR = def.DefFunc("print_r", 1, 2, []def.ArgInfo{{name: "var_"}, {name: "return_"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	return_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPrintR(executeData, returnValue, var_, nil, return_)
})

// generate by ZifIgnoreUserAbort
var DefZifIgnoreUserAbort = def.DefFunc("ignore_user_abort", 0, 1, []def.ArgInfo{{name: "value"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIgnoreUserAbort(executeData, returnValue, nil, value)
})

// generate by ZifRegisterTickFunction
var DefZifRegisterTickFunction = def.DefFunc("register_tick_function", 1, -1, []def.ArgInfo{{name: "function_name"}, {name: "parameters"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	function_name := fp.ParseZval()
	fp.StartOptional()
	parameters := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifRegisterTickFunction(executeData, returnValue, function_name, nil, parameters)
})

// generate by ZifParseIniFile
var DefZifParseIniFile = def.DefFunc("parse_ini_file", 1, 3, []def.ArgInfo{{name: "filename"}, {name: "process_sections"}, {name: "scanner_mode"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	process_sections := fp.ParseZval()
	scanner_mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifParseIniFile(executeData, returnValue, filename, nil, process_sections, scanner_mode)
})

// generate by ZifParseIniString
var DefZifParseIniString = def.DefFunc("parse_ini_string", 1, 3, []def.ArgInfo{{name: "ini_string"}, {name: "process_sections"}, {name: "scanner_mode"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	ini_string := fp.ParseZval()
	fp.StartOptional()
	process_sections := fp.ParseZval()
	scanner_mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifParseIniString(executeData, returnValue, ini_string, nil, process_sections, scanner_mode)
})
