package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifConstant
var DefZifConstant = def.DefFunc("constant", 1, 1, []def.ArgInfo{{name: "const_name"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	const_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifConstant(executeData, returnValue, const_name)
})

// generate by ZifInetNtop
var DefZifInetNtop = def.DefFunc("inet_ntop", 1, 1, []def.ArgInfo{{name: "in_addr"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	in_addr := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifInetNtop(executeData, returnValue, in_addr)
})

// generate by ZifIp2long
var DefZifIp2long = def.DefFunc("ip2long", 1, 1, []def.ArgInfo{{name: "ip_address"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	ip_address := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIp2long(executeData, returnValue, ip_address)
})

// generate by ZifLong2ip
var DefZifLong2ip = def.DefFunc("long2ip", 1, 1, []def.ArgInfo{{name: "proper_address"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	proper_address := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLong2ip(executeData, returnValue, proper_address)
})

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

// generate by ZifPutenv
var DefZifPutenv = def.DefFunc("putenv", 1, 1, []def.ArgInfo{{name: "setting"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	setting := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPutenv(executeData, returnValue, setting)
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

// generate by ZifFlush
var DefZifFlush = def.DefFunc("flush", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifFlush(executeData, returnValue)
})

// generate by ZifSleep
var DefZifSleep = def.DefFunc("sleep", 1, 1, []def.ArgInfo{{name: "seconds"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	seconds := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSleep(executeData, returnValue, seconds)
})

// generate by ZifUsleep
var DefZifUsleep = def.DefFunc("usleep", 1, 1, []def.ArgInfo{{name: "micro_seconds"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	micro_seconds := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUsleep(executeData, returnValue, micro_seconds)
})

// generate by ZifTimeNanosleep
var DefZifTimeNanosleep = def.DefFunc("time_nanosleep", 2, 2, []def.ArgInfo{{name: "seconds"}, {name: "nanoseconds"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	seconds := fp.ParseZval()
	nanoseconds := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTimeNanosleep(executeData, returnValue, seconds, nanoseconds)
})

// generate by ZifTimeSleepUntil
var DefZifTimeSleepUntil = def.DefFunc("time_sleep_until", 1, 1, []def.ArgInfo{{name: "timestamp"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	timestamp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTimeSleepUntil(executeData, returnValue, timestamp)
})

// generate by ZifGetCurrentUser
var DefZifGetCurrentUser = def.DefFunc("get_current_user", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetCurrentUser(executeData, returnValue)
})

// generate by ZifGetCfgVar
var DefZifGetCfgVar = def.DefFunc("get_cfg_var", 1, 1, []def.ArgInfo{{name: "option_name"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	option_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetCfgVar(executeData, returnValue, option_name)
})

// generate by ZifGetMagicQuotesRuntime
var DefZifGetMagicQuotesRuntime = def.DefFunc("get_magic_quotes_runtime", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetMagicQuotesRuntime(executeData, returnValue)
})

// generate by ZifGetMagicQuotesGpc
var DefZifGetMagicQuotesGpc = def.DefFunc("get_magic_quotes_gpc", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetMagicQuotesGpc(executeData, returnValue)
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

// generate by ZifPhpStripWhitespace
var DefZifPhpStripWhitespace = def.DefFunc("php_strip_whitespace", 1, 1, []def.ArgInfo{{name: "file_name"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	file_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPhpStripWhitespace(executeData, returnValue, file_name)
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

// generate by ZifIniGet
var DefZifIniGet = def.DefFunc("ini_get", 1, 1, []def.ArgInfo{{name: "varname"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	varname := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIniGet(executeData, returnValue, varname)
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

// generate by ZifIniSet
var DefZifIniSet = def.DefFunc("ini_set", 2, 2, []def.ArgInfo{{name: "varname"}, {name: "newvalue"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	varname := fp.ParseZval()
	newvalue := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIniSet(executeData, returnValue, varname, newvalue)
})

// generate by ZifIniRestore
var DefZifIniRestore = def.DefFunc("ini_restore", 1, 1, []def.ArgInfo{{name: "varname"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	varname := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIniRestore(executeData, returnValue, varname)
})

// generate by ZifSetIncludePath
var DefZifSetIncludePath = def.DefFunc("set_include_path", 1, 1, []def.ArgInfo{{name: "new_include_path"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	new_include_path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSetIncludePath(executeData, returnValue, new_include_path)
})

// generate by ZifGetIncludePath
var DefZifGetIncludePath = def.DefFunc("get_include_path", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetIncludePath(executeData, returnValue)
})

// generate by ZifRestoreIncludePath
var DefZifRestoreIncludePath = def.DefFunc("restore_include_path", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifRestoreIncludePath(executeData, returnValue)
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

// generate by ZifConnectionAborted
var DefZifConnectionAborted = def.DefFunc("connection_aborted", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifConnectionAborted(executeData, returnValue)
})

// generate by ZifConnectionStatus
var DefZifConnectionStatus = def.DefFunc("connection_status", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifConnectionStatus(executeData, returnValue)
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

// generate by ZifGetservbyname
var DefZifGetservbyname = def.DefFunc("getservbyname", 2, 2, []def.ArgInfo{{name: "service"}, {name: "protocol"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	service := fp.ParseZval()
	protocol := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetservbyname(executeData, returnValue, service, protocol)
})

// generate by ZifGetservbyport
var DefZifGetservbyport = def.DefFunc("getservbyport", 2, 2, []def.ArgInfo{{name: "port"}, {name: "protocol"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	port := fp.ParseZval()
	protocol := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetservbyport(executeData, returnValue, port, protocol)
})

// generate by ZifGetprotobyname
var DefZifGetprotobyname = def.DefFunc("getprotobyname", 1, 1, []def.ArgInfo{{name: "name"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetprotobyname(executeData, returnValue, name)
})

// generate by ZifGetprotobynumber
var DefZifGetprotobynumber = def.DefFunc("getprotobynumber", 1, 1, []def.ArgInfo{{name: "proto"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	proto := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetprotobynumber(executeData, returnValue, proto)
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

// generate by ZifUnregisterTickFunction
var DefZifUnregisterTickFunction = def.DefFunc("unregister_tick_function", 1, 1, []def.ArgInfo{{name: "function_name"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	function_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUnregisterTickFunction(executeData, returnValue, function_name)
})

// generate by ZifIsUploadedFile
var DefZifIsUploadedFile = def.DefFunc("is_uploaded_file", 1, 1, []def.ArgInfo{{name: "path"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsUploadedFile(executeData, returnValue, path)
})

// generate by ZifMoveUploadedFile
var DefZifMoveUploadedFile = def.DefFunc("move_uploaded_file", 2, 2, []def.ArgInfo{{name: "path"}, {name: "new_path"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	path := fp.ParseZval()
	new_path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMoveUploadedFile(executeData, returnValue, path, new_path)
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

// generate by ZifSysGetloadavg
var DefZifSysGetloadavg = def.DefFunc("sys_getloadavg", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifSysGetloadavg(executeData, returnValue)
})
