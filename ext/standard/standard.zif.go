package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifAssert
var DefZifAssert = def.DefFunc("assert", 1, 2, []def.ArgInfo{{Name: "assertion"}, {Name: "description"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	assertion := fp.ParseZval()
	fp.StartOptional()
	description := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAssert(executeData, returnValue, assertion, nil, description)
})

// generate by ZifAssertOptions
var DefZifAssertOptions = def.DefFunc("assert_options", 1, 2, []def.ArgInfo{{Name: "what"}, {Name: "value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	what := fp.ParseZval()
	fp.StartOptional()
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAssertOptions(executeData, returnValue, what, nil, value)
})

// generate by ZifBase64Encode
var DefZifBase64Encode = def.DefFunc("base64_encode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifBase64Encode(str)
	returnValue.SetString(ret)
})

// generate by ZifBase64Decode
var DefZifBase64Decode = def.DefFunc("base64_decode", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "strict"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	strict := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifBase64Decode(str, nil, strict)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifConstant
var DefZifConstant = def.DefFunc("constant", 1, 1, []def.ArgInfo{{Name: "const_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	const_name := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ZifConstant(returnValue, const_name)
})

// generate by ZifInetNtop
var DefZifInetNtop = def.DefFunc("inet_ntop", 1, 1, []def.ArgInfo{{Name: "ip"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	ip := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifInetNtop(ip)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifInetPton
var DefZifInetPton = def.DefFunc("inet_pton", 1, 1, []def.ArgInfo{{Name: "ip_address"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	ip_address := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifInetPton(ip_address)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifIp2long
var DefZifIp2long = def.DefFunc("ip2long", 1, 1, []def.ArgInfo{{Name: "ip_address"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	ip_address := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifIp2long(ip_address)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifLong2ip
var DefZifLong2ip = def.DefFunc("long2ip", 1, 1, []def.ArgInfo{{Name: "ip_address"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	ip_address := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifLong2ip(ip_address)
	returnValue.SetString(ret)
})

// generate by ZifGetenv
var DefZifGetenv = def.DefFunc("getenv", 0, 2, []def.ArgInfo{{Name: "varname"}, {Name: "local_only"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	varname_ := fp.ParseStringValNullable()
	local_only := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifGetenv(nil, varname_, local_only)
	returnValue.SetBy(ret)
})

// generate by ZifPutenv
var DefZifPutenv = def.DefFunc("putenv", 1, 1, []def.ArgInfo{{Name: "setting"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	setting := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifPutenv(setting)
	returnValue.SetBool(ret)
})

// generate by ZifGetopt
var DefZifGetopt = def.DefFunc("getopt", 1, 3, []def.ArgInfo{{Name: "short_options"}, {Name: "long_options"}, {Name: "optind"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	short_options := fp.ParseStringVal()
	fp.StartOptional()
	long_options := fp.ParseArrayHt()
	optind := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifGetopt(executeData, returnValue, short_options, nil, long_options, optind)
})

// generate by ZifFlush
var DefZifFlush = def.DefFunc("flush", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifFlush()
})

// generate by ZifSleep
var DefZifSleep = def.DefFunc("sleep", 1, 1, []def.ArgInfo{{Name: "seconds"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	seconds := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSleep(seconds)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifUsleep
var DefZifUsleep = def.DefFunc("usleep", 1, 1, []def.ArgInfo{{Name: "micro_seconds"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	micro_seconds := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifUsleep(micro_seconds)
	returnValue.SetBy(ret)
})

// generate by ZifTimeNanosleep
var DefZifTimeNanosleep = def.DefFunc("time_nanosleep", 2, 2, []def.ArgInfo{{Name: "seconds"}, {Name: "nanoseconds"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	seconds := fp.ParseLong()
	nanoseconds := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifTimeNanosleep(seconds, nanoseconds)
	returnValue.SetBy(ret)
})

// generate by ZifTimeSleepUntil
var DefZifTimeSleepUntil = def.DefFunc("time_sleep_until", 1, 1, []def.ArgInfo{{Name: "timestamp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	timestamp := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifTimeSleepUntil(timestamp)
	returnValue.SetBy(ret)
})

// generate by ZifGetCurrentUser
var DefZifGetCurrentUser = def.DefFunc("get_current_user", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGetCurrentUser()
	returnValue.SetString(ret)
})

// generate by ZifGetCfgVar
var DefZifGetCfgVar = def.DefFunc("get_cfg_var", 1, 1, []def.ArgInfo{{Name: "option_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	option_name := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifGetCfgVar(option_name)
	returnValue.SetBy(ret)
})

// generate by ZifGetMagicQuotesRuntime
var DefZifGetMagicQuotesRuntime = def.DefFunc("get_magic_quotes_runtime", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetMagicQuotesRuntime(executeData, returnValue)
})

// generate by ZifGetMagicQuotesGpc
var DefZifGetMagicQuotesGpc = def.DefFunc("get_magic_quotes_gpc", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetMagicQuotesGpc(executeData, returnValue)
})

// generate by ZifErrorLog
var DefZifErrorLog = def.DefFunc("error_log", 1, 4, []def.ArgInfo{{Name: "message"}, {Name: "message_type"}, {Name: "destination"}, {Name: "extra_headers"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	message := fp.ParseStringVal()
	fp.StartOptional()
	message_type := fp.ParseLong()
	destination := fp.ParsePathValNullable()
	extra_headers := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifErrorLog(message, nil, message_type, destination, extra_headers)
	returnValue.SetBool(ret)
})

// generate by ZifErrorGetLast
var DefZifErrorGetLast = def.DefFunc("error_get_last", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifErrorGetLast()
	returnValue.SetBy(ret)
})

// generate by ZifErrorClearLast
var DefZifErrorClearLast = def.DefFunc("error_clear_last", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifErrorClearLast()
})

// generate by ZifCallUserFunc
var DefZifCallUserFunc = def.DefFunc("call_user_func", 1, -1, []def.ArgInfo{{Name: "callback"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	callback := fp.ParseCallable()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret := ZifCallUserFunc(callback, nil, args)
	returnValue.SetBy(ret)
})

// generate by ZifCallUserFuncArray
var DefZifCallUserFuncArray = def.DefFunc("call_user_func_array", 2, 2, []def.ArgInfo{{Name: "callback"}, {Name: "parameters"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	callback := fp.ParseCallable()
	parameters := fp.ParseArrayHt()
	if fp.HasError() {
		return
	}
	ret := ZifCallUserFuncArray(callback, parameters)
	returnValue.SetBy(ret)
})

// generate by ZifForwardStaticCall
var DefZifForwardStaticCall = def.DefFunc("forward_static_call", 1, -1, []def.ArgInfo{{Name: "function_name"}, {Name: "parameters"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifForwardStaticCallArray = def.DefFunc("forward_static_call_array", 2, 2, []def.ArgInfo{{Name: "function_name"}, {Name: "parameters"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	function_name := fp.ParseZval()
	parameters := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifForwardStaticCallArray(executeData, returnValue, function_name, parameters)
})

// generate by ZifRegisterShutdownFunction
var DefZifRegisterShutdownFunction = def.DefFunc("register_shutdown_function", 1, -1, []def.ArgInfo{{Name: "function_name"}, {Name: "parameters"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	function_name := fp.ParseZval()
	fp.StartOptional()
	parameters := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifRegisterShutdownFunction(function_name, nil, parameters)
})

// generate by ZifHighlightFile
var DefZifHighlightFile = def.DefFunc("highlight_file", 1, 2, []def.ArgInfo{{Name: "file_name"}, {Name: "return"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	file_name := fp.ParsePathVal()
	fp.StartOptional()
	return_ := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifHighlightFile(file_name, nil, return_)
	returnValue.SetBy(ret)
})

// generate by ZifHighlightFile
var DefZifShowSource = def.DefFunc("show_source", 1, 2, []def.ArgInfo{{Name: "file_name"}, {Name: "return"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	file_name := fp.ParsePathVal()
	fp.StartOptional()
	return_ := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifHighlightFile(file_name, nil, return_)
	returnValue.SetBy(ret)
})

// generate by ZifPhpStripWhitespace
var DefZifPhpStripWhitespace = def.DefFunc("php_strip_whitespace", 1, 1, []def.ArgInfo{{Name: "file_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	file_name := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifPhpStripWhitespace(file_name)
	returnValue.SetBy(ret)
})

// generate by ZifHighlightString
var DefZifHighlightString = def.DefFunc("highlight_string", 1, 2, []def.ArgInfo{{Name: "string"}, {Name: "return"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	string := fp.ParseStringVal()
	fp.StartOptional()
	return_ := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifHighlightString(string, nil, return_)
	returnValue.SetBy(ret)
})

// generate by ZifIniGet
var DefZifIniGet = def.DefFunc("ini_get", 1, 1, []def.ArgInfo{{Name: "varname"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	varname := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifIniGet(varname)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifIniGetAll
var DefZifIniGetAll = def.DefFunc("ini_get_all", 0, 2, []def.ArgInfo{{Name: "extension"}, {Name: "details"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	extension := fp.ParseStringValNullable()
	details_ := fp.ParseBoolValNullable()
	if fp.HasError() {
		return
	}
	ZifIniGetAll(returnValue, nil, extension, details_)
})

// generate by ZifIniSet
var DefZifIniSet = def.DefFunc("ini_set", 2, 2, []def.ArgInfo{{Name: "varname"}, {Name: "newvalue"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	varname := fp.ParseStringVal()
	newvalue := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifIniSet(returnValue, varname, newvalue)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifIniSet
var DefZifIniAlter = def.DefFunc("ini_alter", 2, 2, []def.ArgInfo{{Name: "varname"}, {Name: "newvalue"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	varname := fp.ParseStringVal()
	newvalue := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifIniSet(returnValue, varname, newvalue)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifIniRestore
var DefZifIniRestore = def.DefFunc("ini_restore", 1, 1, []def.ArgInfo{{Name: "var_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_name := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ZifIniRestore(var_name)
})

// generate by ZifSetIncludePath
var DefZifSetIncludePath = def.DefFunc("set_include_path", 1, 1, []def.ArgInfo{{Name: "new_include_path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	new_include_path := fp.ParsePathVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSetIncludePath(new_include_path)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifGetIncludePath
var DefZifGetIncludePath = def.DefFunc("get_include_path", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret, ok := ZifGetIncludePath()
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifRestoreIncludePath
var DefZifRestoreIncludePath = def.DefFunc("restore_include_path", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifRestoreIncludePath()
})

// generate by ZifPrintR
var DefZifPrintR = def.DefFunc("print_r", 1, 2, []def.ArgInfo{{Name: "var"}, {Name: "return"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	return_ := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifPrintR(var_, nil, return_)
	returnValue.SetBy(ret)
})

// generate by ZifConnectionAborted
var DefZifConnectionAborted = def.DefFunc("connection_aborted", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifConnectionAborted()
	returnValue.SetLong(ret)
})

// generate by ZifConnectionStatus
var DefZifConnectionStatus = def.DefFunc("connection_status", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifConnectionStatus()
	returnValue.SetLong(ret)
})

// generate by ZifIgnoreUserAbort
var DefZifIgnoreUserAbort = def.DefFunc("ignore_user_abort", 0, 1, []def.ArgInfo{{Name: "enable"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	enable := fp.ParseBoolValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifIgnoreUserAbort(executeData, returnValue, nil, enable)
	returnValue.SetLong(ret)
})

// generate by ZifGetservbyname
var DefZifGetservbyname = def.DefFunc("getservbyname", 2, 2, []def.ArgInfo{{Name: "service"}, {Name: "protocol"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	service := fp.ParseZval()
	protocol := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetservbyname(executeData, returnValue, service, protocol)
})

// generate by ZifGetservbyport
var DefZifGetservbyport = def.DefFunc("getservbyport", 2, 2, []def.ArgInfo{{Name: "port"}, {Name: "protocol"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	port := fp.ParseZval()
	protocol := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetservbyport(executeData, returnValue, port, protocol)
})

// generate by ZifGetprotobyname
var DefZifGetprotobyname = def.DefFunc("getprotobyname", 1, 1, []def.ArgInfo{{Name: "name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetprotobyname(executeData, returnValue, name)
})

// generate by ZifGetprotobynumber
var DefZifGetprotobynumber = def.DefFunc("getprotobynumber", 1, 1, []def.ArgInfo{{Name: "proto"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	proto := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetprotobynumber(executeData, returnValue, proto)
})

// generate by ZifRegisterTickFunction
var DefZifRegisterTickFunction = def.DefFunc("register_tick_function", 1, -1, []def.ArgInfo{{Name: "function_name"}, {Name: "parameters"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifUnregisterTickFunction = def.DefFunc("unregister_tick_function", 1, 1, []def.ArgInfo{{Name: "function_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	function_name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUnregisterTickFunction(executeData, returnValue, function_name)
})

// generate by ZifIsUploadedFile
var DefZifIsUploadedFile = def.DefFunc("is_uploaded_file", 1, 1, []def.ArgInfo{{Name: "path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	path := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifIsUploadedFile(path)
	returnValue.SetBool(ret)
})

// generate by ZifMoveUploadedFile
var DefZifMoveUploadedFile = def.DefFunc("move_uploaded_file", 2, 2, []def.ArgInfo{{Name: "path"}, {Name: "new_path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	path := fp.ParseStringVal()
	new_path := fp.ParsePathVal()
	if fp.HasError() {
		return
	}
	ret := ZifMoveUploadedFile(executeData, returnValue, path, new_path)
	returnValue.SetBool(ret)
})

// generate by ZifParseIniFile
var DefZifParseIniFile = def.DefFunc("parse_ini_file", 1, 3, []def.ArgInfo{{Name: "filename"}, {Name: "process_sections"}, {Name: "scanner_mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	filename := fp.ParseStringVal()
	fp.StartOptional()
	process_sections_ := fp.ParseZval()
	scanner_mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifParseIniFile(executeData, returnValue, filename, nil, process_sections_, scanner_mode)
})

// generate by ZifParseIniString
var DefZifParseIniString = def.DefFunc("parse_ini_string", 1, 3, []def.ArgInfo{{Name: "ini_string"}, {Name: "process_sections"}, {Name: "scanner_mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifSysGetloadavg = def.DefFunc("sys_getloadavg", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifSysGetloadavg(executeData, returnValue)
})

// generate by ZifGetBrowser
var DefZifGetBrowser = def.DefFunc("get_browser", 0, 2, []def.ArgInfo{{Name: "browser_name"}, {Name: "return_array"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	browser_name := fp.ParseStringValNullable()
	return_array := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ZifGetBrowser(executeData, returnValue, nil, browser_name, return_array)
})

// generate by ZifCrc32
var DefZifCrc32 = def.DefFunc("crc32", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifCrc32(str)
	returnValue.SetLong(ret)
})

// generate by ZifCrypt
var DefZifCrypt = def.DefFunc("crypt", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "salt"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str_ := fp.ParseStringVal()
	fp.StartOptional()
	salt_ := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifCrypt(str_, nil, salt_)
	returnValue.SetString(ret)
})

// generate by ZifConvertCyrString
var DefZifConvertCyrString = def.DefFunc("convert_cyr_string", 3, 3, []def.ArgInfo{{Name: "str"}, {Name: "from"}, {Name: "to"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	str := fp.ParseZval()
	from := fp.ParseZval()
	to := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifConvertCyrString(executeData, returnValue, str, from, to)
})

// generate by ZifStrptime
var DefZifStrptime = def.DefFunc("strptime", 2, 2, []def.ArgInfo{{Name: "timestamp"}, {Name: "format"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	timestamp := fp.ParseZval()
	format := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrptime(executeData, returnValue, timestamp, format)
})

// generate by ZifOpendir
var DefZifOpendir = def.DefFunc("opendir", 1, 2, []def.ArgInfo{{Name: "path"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	path := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifOpendir(executeData, returnValue, path, nil, context)
})

// generate by ZifGetdir
var DefZifGetdir = def.DefFunc("getdir", 1, 2, []def.ArgInfo{{Name: "directory"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	directory := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetdir(executeData, returnValue, directory, nil, context)
})

// generate by ZifClosedir
var DefZifClosedir = def.DefFunc("closedir", 0, 1, []def.ArgInfo{{Name: "dir_handle"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	dir_handle := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClosedir(executeData, returnValue, nil, dir_handle)
})

// generate by ZifChroot
var DefZifChroot = def.DefFunc("chroot", 1, 1, []def.ArgInfo{{Name: "directory"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	directory := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChroot(executeData, returnValue, directory)
})

// generate by ZifChdir
var DefZifChdir = def.DefFunc("chdir", 1, 1, []def.ArgInfo{{Name: "directory"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	directory := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChdir(executeData, returnValue, directory)
})

// generate by ZifGetcwd
var DefZifGetcwd = def.DefFunc("getcwd", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGetcwd(executeData, returnValue)
})

// generate by ZifRewinddir
var DefZifRewinddir = def.DefFunc("rewinddir", 0, 1, []def.ArgInfo{{Name: "dir_handle"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	dir_handle := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRewinddir(executeData, returnValue, nil, dir_handle)
})

// generate by ZifGlob
var DefZifGlob = def.DefFunc("glob", 1, 2, []def.ArgInfo{{Name: "pattern"}, {Name: "flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	pattern := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGlob(executeData, returnValue, pattern, nil, flags)
})

// generate by ZifScandir
var DefZifScandir = def.DefFunc("scandir", 1, 3, []def.ArgInfo{{Name: "dir"}, {Name: "sorting_order"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	dir := fp.ParseZval()
	fp.StartOptional()
	sorting_order := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifScandir(executeData, returnValue, dir, nil, sorting_order, context)
})

// generate by ZifGethostname
var DefZifGethostname = def.DefFunc("gethostname", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifGethostname(executeData, returnValue)
})

// generate by ZifGethostbyaddr
var DefZifGethostbyaddr = def.DefFunc("gethostbyaddr", 1, 1, []def.ArgInfo{{Name: "ip_address"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	ip_address := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGethostbyaddr(executeData, returnValue, ip_address)
})

// generate by ZifGethostbyname
var DefZifGethostbyname = def.DefFunc("gethostbyname", 1, 1, []def.ArgInfo{{Name: "hostname"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hostname := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGethostbyname(executeData, returnValue, hostname)
})

// generate by ZifGethostbynamel
var DefZifGethostbynamel = def.DefFunc("gethostbynamel", 1, 1, []def.ArgInfo{{Name: "hostname"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hostname := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGethostbynamel(executeData, returnValue, hostname)
})

// generate by ZifDnsCheckRecord
var DefZifDnsCheckRecord = def.DefFunc("dns_check_record", 1, 2, []def.ArgInfo{{Name: "host"}, {Name: "type"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	host := fp.ParseZval()
	fp.StartOptional()
	type_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDnsCheckRecord(executeData, returnValue, host, nil, type_)
})

// generate by ZifDnsCheckRecord
var DefZifCheckdnsrr = def.DefFunc("checkdnsrr", 1, 2, []def.ArgInfo{{Name: "host"}, {Name: "type"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	host := fp.ParseZval()
	fp.StartOptional()
	type_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDnsCheckRecord(executeData, returnValue, host, nil, type_)
})

// generate by ZifDnsGetMx
var DefZifDnsGetMx = def.DefFunc("dns_get_mx", 2, 3, []def.ArgInfo{{Name: "hostname"}, {Name: "mxhosts"}, {Name: "weight"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	hostname := fp.ParseZval()
	mxhosts := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	weight := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifDnsGetMx(executeData, returnValue, hostname, mxhosts, nil, weight)
})

// generate by ZifDnsGetMx
var DefZifGetmxrr = def.DefFunc("getmxrr", 2, 3, []def.ArgInfo{{Name: "hostname"}, {Name: "mxhosts"}, {Name: "weight"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	hostname := fp.ParseZval()
	mxhosts := fp.ParseZvalEx(false, true)
	fp.StartOptional()
	weight := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifDnsGetMx(executeData, returnValue, hostname, mxhosts, nil, weight)
})

// generate by ZifExec
var DefZifExec = def.DefFunc("exec", 1, 3, []def.ArgInfo{{Name: "command"}, {Name: "output"}, {Name: "return_value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	command := fp.ParseZval()
	fp.StartOptional()
	output := fp.ParseZvalEx(false, true)
	return_value := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifExec(executeData, returnValue, command, nil, output, return_value)
})

// generate by ZifSystem
var DefZifSystem = def.DefFunc("system", 1, 2, []def.ArgInfo{{Name: "command"}, {Name: "return_value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	command := fp.ParseZval()
	fp.StartOptional()
	return_value := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifSystem(executeData, returnValue, command, nil, return_value)
})

// generate by ZifPassthru
var DefZifPassthru = def.DefFunc("passthru", 1, 2, []def.ArgInfo{{Name: "command"}, {Name: "return_value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	command := fp.ParseZval()
	fp.StartOptional()
	return_value := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifPassthru(executeData, returnValue, command, nil, return_value)
})

// generate by ZifEscapeshellcmd
var DefZifEscapeshellcmd = def.DefFunc("escapeshellcmd", 1, 1, []def.ArgInfo{{Name: "command"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	command := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifEscapeshellcmd(command)
	returnValue.SetString(ret)
})

// generate by ZifEscapeshellarg
var DefZifEscapeshellarg = def.DefFunc("escapeshellarg", 1, 1, []def.ArgInfo{{Name: "arg"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	arg := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifEscapeshellarg(arg)
	returnValue.SetString(ret)
})

// generate by ZifShellExec
var DefZifShellExec = def.DefFunc("shell_exec", 1, 1, []def.ArgInfo{{Name: "cmd"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	cmd := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifShellExec(executeData, returnValue, cmd)
})

// generate by ZifFlock
var DefZifFlock = def.DefFunc("flock", 2, 3, []def.ArgInfo{{Name: "fp"}, {Name: "operation"}, {Name: "wouldblock"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	fp := fp.ParseZval()
	operation := fp.ParseZval()
	fp.StartOptional()
	wouldblock := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifFlock(executeData, returnValue, fp, operation, nil, wouldblock)
})

// generate by ZifGetMetaTags
var DefZifGetMetaTags = def.DefFunc("get_meta_tags", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "use_include_path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	use_include_path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetMetaTags(executeData, returnValue, filename, nil, use_include_path)
})

// generate by ZifFileGetContents
var DefZifFileGetContents = def.DefFunc("file_get_contents", 1, 5, []def.ArgInfo{{Name: "filename"}, {Name: "flags"}, {Name: "context"}, {Name: "offset"}, {Name: "maxlen"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 5, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	context := fp.ParseZval()
	offset := fp.ParseZval()
	maxlen := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileGetContents(executeData, returnValue, filename, nil, flags, context, offset, maxlen)
})

// generate by ZifFilePutContents
var DefZifFilePutContents = def.DefFunc("file_put_contents", 2, 4, []def.ArgInfo{{Name: "filename"}, {Name: "data"}, {Name: "flags"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	filename := fp.ParseZval()
	data := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFilePutContents(executeData, returnValue, filename, data, nil, flags, context)
})

// generate by ZifFile
var DefZifFile = def.DefFunc("file", 1, 3, []def.ArgInfo{{Name: "filename"}, {Name: "flags"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFile(executeData, returnValue, filename, nil, flags, context)
})

// generate by ZifTempnam
var DefZifTempnam = def.DefFunc("tempnam", 2, 2, []def.ArgInfo{{Name: "dir"}, {Name: "prefix"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	dir := fp.ParseZval()
	prefix := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTempnam(executeData, returnValue, dir, prefix)
})

// generate by ZifTmpfile
var DefZifTmpfile = def.DefFunc("tmpfile", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifTmpfile(executeData, returnValue)
})

// generate by ZifOpen
var DefZifOpen = def.DefFunc("open", 2, 4, []def.ArgInfo{{Name: "filename"}, {Name: "mode"}, {Name: "use_include_path"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	filename := fp.ParseZval()
	mode := fp.ParseZval()
	fp.StartOptional()
	use_include_path := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifOpen(executeData, returnValue, filename, mode, nil, use_include_path, context)
})

// generate by ZifFclose
var DefZifFclose = def.DefFunc("fclose", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFclose(executeData, returnValue, fp)
})

// generate by ZifPopen
var DefZifPopen = def.DefFunc("popen", 2, 2, []def.ArgInfo{{Name: "command"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	command := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPopen(executeData, returnValue, command, mode)
})

// generate by ZifPclose
var DefZifPclose = def.DefFunc("pclose", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPclose(executeData, returnValue, fp)
})

// generate by ZifFeof
var DefZifFeof = def.DefFunc("feof", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFeof(executeData, returnValue, fp)
})

// generate by ZifFgets
var DefZifFgets = def.DefFunc("fgets", 1, 2, []def.ArgInfo{{Name: "fp"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	fp := fp.ParseResource()
	fp.StartOptional()
	length := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifFgets(fp, nil, length)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifFgetc
var DefZifFgetc = def.DefFunc("fgetc", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFgetc(executeData, returnValue, fp)
})

// generate by ZifFgetss
var DefZifFgetss = def.DefFunc("fgetss", 1, 3, []def.ArgInfo{{Name: "fp"}, {Name: "length"}, {Name: "allowable_tags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	fp := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseLongNullable()
	allowable_tags := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifFgetss(fp, nil, length, allowable_tags)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifFwrite
var DefZifFwrite = def.DefFunc("fwrite", 2, 3, []def.ArgInfo{{Name: "fp"}, {Name: "str"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	fp := fp.ParseZval()
	str := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFwrite(executeData, returnValue, fp, str, nil, length)
})

// generate by ZifFwrite
var DefZifFputs = def.DefFunc("fputs", 2, 3, []def.ArgInfo{{Name: "fp"}, {Name: "str"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	fp := fp.ParseZval()
	str := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFwrite(executeData, returnValue, fp, str, nil, length)
})

// generate by ZifFflush
var DefZifFflush = def.DefFunc("fflush", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFflush(executeData, returnValue, fp)
})

// generate by ZifRewind
var DefZifRewind = def.DefFunc("rewind", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRewind(executeData, returnValue, fp)
})

// generate by ZifFtell
var DefZifFtell = def.DefFunc("ftell", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFtell(executeData, returnValue, fp)
})

// generate by ZifFseek
var DefZifFseek = def.DefFunc("fseek", 2, 3, []def.ArgInfo{{Name: "fp"}, {Name: "offset"}, {Name: "whence"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	fp := fp.ParseZval()
	offset := fp.ParseZval()
	fp.StartOptional()
	whence := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFseek(executeData, returnValue, fp, offset, nil, whence)
})

// generate by ZifMkdir
var DefZifMkdir = def.DefFunc("mkdir", 1, 4, []def.ArgInfo{{Name: "pathname"}, {Name: "mode"}, {Name: "recursive"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	pathname := fp.ParseZval()
	fp.StartOptional()
	mode := fp.ParseZval()
	recursive := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMkdir(executeData, returnValue, pathname, nil, mode, recursive, context)
})

// generate by ZifRmdir
var DefZifRmdir = def.DefFunc("rmdir", 1, 2, []def.ArgInfo{{Name: "dirname"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	dirname := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRmdir(executeData, returnValue, dirname, nil, context)
})

// generate by ZifReadfile
var DefZifReadfile = def.DefFunc("readfile", 1, 3, []def.ArgInfo{{Name: "filename"}, {Name: "flags"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifReadfile(executeData, returnValue, filename, nil, flags, context)
})

// generate by ZifUmask
var DefZifUmask = def.DefFunc("umask", 0, 1, []def.ArgInfo{{Name: "mask"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	mask := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUmask(executeData, returnValue, nil, mask)
})

// generate by ZifFpassthru
var DefZifFpassthru = def.DefFunc("fpassthru", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFpassthru(executeData, returnValue, fp)
})

// generate by ZifRename
var DefZifRename = def.DefFunc("rename", 2, 3, []def.ArgInfo{{Name: "old_name"}, {Name: "new_name"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	old_name := fp.ParseZval()
	new_name := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRename(executeData, returnValue, old_name, new_name, nil, context)
})

// generate by ZifUnlink
var DefZifUnlink = def.DefFunc("unlink", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUnlink(executeData, returnValue, filename, nil, context)
})

// generate by ZifFtruncate
var DefZifFtruncate = def.DefFunc("ftruncate", 2, 2, []def.ArgInfo{{Name: "fp"}, {Name: "size"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	size := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFtruncate(executeData, returnValue, fp, size)
})

// generate by ZifFstat
var DefZifFstat = def.DefFunc("fstat", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFstat(executeData, returnValue, fp)
})

// generate by ZifCopy
var DefZifCopy = def.DefFunc("copy", 2, 3, []def.ArgInfo{{Name: "source_file"}, {Name: "destination_file"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	source_file := fp.ParseZval()
	destination_file := fp.ParseZval()
	fp.StartOptional()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCopy(executeData, returnValue, source_file, destination_file, nil, context)
})

// generate by ZifFread
var DefZifFread = def.DefFunc("fread", 2, 2, []def.ArgInfo{{Name: "fp"}, {Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	length := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFread(executeData, returnValue, fp, length)
})

// generate by ZifFputcsv
var DefZifFputcsv = def.DefFunc("fputcsv", 2, 5, []def.ArgInfo{{Name: "fp"}, {Name: "fields"}, {Name: "delimiter"}, {Name: "enclosure"}, {Name: "escape_char"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 5, 0)
	fp := fp.ParseZval()
	fields := fp.ParseZval()
	fp.StartOptional()
	delimiter := fp.ParseZval()
	enclosure := fp.ParseZval()
	escape_char := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFputcsv(executeData, returnValue, fp, fields, nil, delimiter, enclosure, escape_char)
})

// generate by ZifFgetcsv
var DefZifFgetcsv = def.DefFunc("fgetcsv", 1, 5, []def.ArgInfo{{Name: "fp"}, {Name: "length"}, {Name: "delimiter"}, {Name: "enclosure"}, {Name: "escape"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 5, 0)
	fp := fp.ParseZval()
	fp.StartOptional()
	length := fp.ParseZval()
	delimiter := fp.ParseZval()
	enclosure := fp.ParseZval()
	escape := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFgetcsv(executeData, returnValue, fp, nil, length, delimiter, enclosure, escape)
})

// generate by ZifRealpath
var DefZifRealpath = def.DefFunc("realpath", 1, 1, []def.ArgInfo{{Name: "path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRealpath(executeData, returnValue, path)
})

// generate by ZifFnmatch
var DefZifFnmatch = def.DefFunc("fnmatch", 2, 3, []def.ArgInfo{{Name: "pattern"}, {Name: "filename"}, {Name: "flags"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	pattern := fp.ParseZval()
	filename := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFnmatch(executeData, returnValue, pattern, filename, nil, flags)
})

// generate by ZifSysGetTempDir
var DefZifSysGetTempDir = def.DefFunc("sys_get_temp_dir", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifSysGetTempDir(executeData, returnValue)
})

// generate by ZifDiskTotalSpace
var DefZifDiskTotalSpace = def.DefFunc("disk_total_space", 1, 1, []def.ArgInfo{{Name: "path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDiskTotalSpace(executeData, returnValue, path)
})

// generate by ZifDiskFreeSpace
var DefZifDiskFreeSpace = def.DefFunc("disk_free_space", 1, 1, []def.ArgInfo{{Name: "path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDiskFreeSpace(executeData, returnValue, path)
})

// generate by ZifDiskFreeSpace
var DefZifDiskfreespace = def.DefFunc("diskfreespace", 1, 1, []def.ArgInfo{{Name: "path"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	path := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDiskFreeSpace(executeData, returnValue, path)
})

// generate by ZifChgrp
var DefZifChgrp = def.DefFunc("chgrp", 2, 2, []def.ArgInfo{{Name: "filename"}, {Name: "group"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filename := fp.ParseZval()
	group := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChgrp(executeData, returnValue, filename, group)
})

// generate by ZifLchgrp
var DefZifLchgrp = def.DefFunc("lchgrp", 2, 2, []def.ArgInfo{{Name: "filename"}, {Name: "group"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filename := fp.ParseZval()
	group := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLchgrp(executeData, returnValue, filename, group)
})

// generate by ZifChown
var DefZifChown = def.DefFunc("chown", 2, 2, []def.ArgInfo{{Name: "filename"}, {Name: "user"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filename := fp.ParseZval()
	user := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChown(executeData, returnValue, filename, user)
})

// generate by ZifLchown
var DefZifLchown = def.DefFunc("lchown", 2, 2, []def.ArgInfo{{Name: "filename"}, {Name: "user"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filename := fp.ParseZval()
	user := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLchown(executeData, returnValue, filename, user)
})

// generate by ZifChmod
var DefZifChmod = def.DefFunc("chmod", 2, 2, []def.ArgInfo{{Name: "filename"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filename := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifChmod(executeData, returnValue, filename, mode)
})

// generate by ZifTouch
var DefZifTouch = def.DefFunc("touch", 1, 3, []def.ArgInfo{{Name: "filename"}, {Name: "time"}, {Name: "atime"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	filename := fp.ParseZval()
	fp.StartOptional()
	time := fp.ParseZval()
	atime := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTouch(executeData, returnValue, filename, nil, time, atime)
})

// generate by ZifClearstatcache
var DefZifClearstatcache = def.DefFunc("clearstatcache", 0, 2, []def.ArgInfo{{Name: "clear_realpath_cache"}, {Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	clear_realpath_cache := fp.ParseZval()
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifClearstatcache(executeData, returnValue, nil, clear_realpath_cache, filename)
})

// generate by ZifFileperms
var DefZifFileperms = def.DefFunc("fileperms", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileperms(executeData, returnValue, filename)
})

// generate by ZifFileinode
var DefZifFileinode = def.DefFunc("fileinode", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileinode(executeData, returnValue, filename)
})

// generate by ZifFilesize
var DefZifFilesize = def.DefFunc("filesize", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFilesize(executeData, returnValue, filename)
})

// generate by ZifFileowner
var DefZifFileowner = def.DefFunc("fileowner", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileowner(executeData, returnValue, filename)
})

// generate by ZifFilegroup
var DefZifFilegroup = def.DefFunc("filegroup", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFilegroup(executeData, returnValue, filename)
})

// generate by ZifFileatime
var DefZifFileatime = def.DefFunc("fileatime", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileatime(executeData, returnValue, filename)
})

// generate by ZifFilemtime
var DefZifFilemtime = def.DefFunc("filemtime", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFilemtime(executeData, returnValue, filename)
})

// generate by ZifFilectime
var DefZifFilectime = def.DefFunc("filectime", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFilectime(executeData, returnValue, filename)
})

// generate by ZifFiletype
var DefZifFiletype = def.DefFunc("filetype", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFiletype(executeData, returnValue, filename)
})

// generate by ZifIsWritable
var DefZifIsWritable = def.DefFunc("is_writable", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsWritable(executeData, returnValue, filename)
})

// generate by ZifIsWritable
var DefZifIsWriteable = def.DefFunc("is_writeable", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsWritable(executeData, returnValue, filename)
})

// generate by ZifIsReadable
var DefZifIsReadable = def.DefFunc("is_readable", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsReadable(executeData, returnValue, filename)
})

// generate by ZifIsExecutable
var DefZifIsExecutable = def.DefFunc("is_executable", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsExecutable(executeData, returnValue, filename)
})

// generate by ZifIsFile
var DefZifIsFile = def.DefFunc("is_file", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsFile(executeData, returnValue, filename)
})

// generate by ZifIsDir
var DefZifIsDir = def.DefFunc("is_dir", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsDir(executeData, returnValue, filename)
})

// generate by ZifIsLink
var DefZifIsLink = def.DefFunc("is_link", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsLink(executeData, returnValue, filename)
})

// generate by ZifFileExists
var DefZifFileExists = def.DefFunc("file_exists", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFileExists(executeData, returnValue, filename)
})

// generate by ZifLstat
var DefZifLstat = def.DefFunc("lstat", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifLstat(executeData, returnValue)
})

// generate by ZifStat
var DefZifStat = def.DefFunc("stat", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifStat(executeData, returnValue)
})

// generate by ZifRealpathCacheSize
var DefZifRealpathCacheSize = def.DefFunc("realpath_cache_size", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifRealpathCacheSize(executeData, returnValue)
})

// generate by ZifRealpathCacheGet
var DefZifRealpathCacheGet = def.DefFunc("realpath_cache_get", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifRealpathCacheGet(executeData, returnValue)
})

// generate by ZifUserSprintf
var DefZifSprintf = def.DefFunc("sprintf", 1, -1, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret, ok := ZifUserSprintf(format, nil, args)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifVsprintf
var DefZifVsprintf = def.DefFunc("vsprintf", 2, 2, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	format := fp.ParseZval()
	args := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifVsprintf(executeData, returnValue, format, args)
})

// generate by ZifUserPrintf
var DefZifPrintf = def.DefFunc("printf", 1, -1, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifUserPrintf(executeData, returnValue, format, nil, args)
})

// generate by ZifVprintf
var DefZifVprintf = def.DefFunc("vprintf", 2, 2, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	format := fp.ParseZval()
	args := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifVprintf(executeData, returnValue, format, args)
})

// generate by ZifFprintf
var DefZifFprintf = def.DefFunc("fprintf", 2, -1, []def.ArgInfo{{Name: "stream"}, {Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, -1, 0)
	stream := fp.ParseZval()
	format := fp.ParseZval()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifFprintf(executeData, returnValue, stream, format, nil, args)
})

// generate by ZifVfprintf
var DefZifVfprintf = def.DefFunc("vfprintf", 3, 3, []def.ArgInfo{{Name: "stream"}, {Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	stream := fp.ParseZval()
	format := fp.ParseZval()
	args := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifVfprintf(executeData, returnValue, stream, format, args)
})

// generate by ZifFsockopen
var DefZifFsockopen = def.DefFunc("fsockopen", 1, 5, []def.ArgInfo{{Name: "hostname"}, {Name: "port"}, {Name: "errno"}, {Name: "errstr"}, {Name: "timeout"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 5, 0)
	hostname := fp.ParseZval()
	fp.StartOptional()
	port := fp.ParseZval()
	errno := fp.ParseZvalEx(false, true)
	errstr := fp.ParseZvalEx(false, true)
	timeout := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFsockopen(executeData, returnValue, hostname, nil, port, errno, errstr, timeout)
})

// generate by ZifPfsockopen
var DefZifPfsockopen = def.DefFunc("pfsockopen", 1, 5, []def.ArgInfo{{Name: "hostname"}, {Name: "port"}, {Name: "errno"}, {Name: "errstr"}, {Name: "timeout"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 5, 0)
	hostname := fp.ParseZval()
	fp.StartOptional()
	port := fp.ParseZval()
	errno := fp.ParseZvalEx(false, true)
	errstr := fp.ParseZvalEx(false, true)
	timeout := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPfsockopen(executeData, returnValue, hostname, nil, port, errno, errstr, timeout)
})

// generate by ZifFtok
var DefZifFtok = def.DefFunc("ftok", 2, 2, []def.ArgInfo{{Name: "pathname"}, {Name: "proj"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	pathname := fp.ParseZval()
	proj := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFtok(executeData, returnValue, pathname, proj)
})

// generate by ZifHeader
var DefZifHeader = def.DefFunc("header", 1, 3, []def.ArgInfo{{Name: "header"}, {Name: "replace"}, {Name: "http_response_code"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	header := fp.ParseZval()
	fp.StartOptional()
	replace := fp.ParseZval()
	http_response_code := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHeader(executeData, returnValue, header, nil, replace, http_response_code)
})

// generate by ZifHeaderRemove
var DefZifHeaderRemove = def.DefFunc("header_remove", 0, 1, []def.ArgInfo{{Name: "name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	name := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHeaderRemove(executeData, returnValue, nil, name)
})

// generate by ZifSetcookie
var DefZifSetcookie = def.DefFunc("setcookie", 1, 7, []def.ArgInfo{{Name: "name"}, {Name: "value"}, {Name: "expires_or_options"}, {Name: "path"}, {Name: "domain"}, {Name: "secure"}, {Name: "httponly"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 7, 0)
	name := fp.ParseZval()
	fp.StartOptional()
	value := fp.ParseZval()
	expires_or_options := fp.ParseZval()
	path := fp.ParseZval()
	domain := fp.ParseZval()
	secure := fp.ParseZval()
	httponly := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSetcookie(executeData, returnValue, name, nil, value, expires_or_options, path, domain, secure, httponly)
})

// generate by ZifSetrawcookie
var DefZifSetrawcookie = def.DefFunc("setrawcookie", 1, 7, []def.ArgInfo{{Name: "name"}, {Name: "value"}, {Name: "expires_or_options"}, {Name: "path"}, {Name: "domain"}, {Name: "secure"}, {Name: "httponly"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 7, 0)
	name := fp.ParseZval()
	fp.StartOptional()
	value := fp.ParseZval()
	expires_or_options := fp.ParseZval()
	path := fp.ParseZval()
	domain := fp.ParseZval()
	secure := fp.ParseZval()
	httponly := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSetrawcookie(executeData, returnValue, name, nil, value, expires_or_options, path, domain, secure, httponly)
})

// generate by ZifHeadersSent
var DefZifHeadersSent = def.DefFunc("headers_sent", 0, 2, []def.ArgInfo{{Name: "file"}, {Name: "line"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	file_ := fp.ParseZvalEx(false, true)
	line_ := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifHeadersSent(executeData, returnValue, nil, file_, line_)
})

// generate by ZifHeadersList
var DefZifHeadersList = def.DefFunc("headers_list", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifHeadersList(executeData, returnValue)
})

// generate by ZifHttpResponseCode
var DefZifHttpResponseCode = def.DefFunc("http_response_code", 0, 1, []def.ArgInfo{{Name: "response_code"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	response_code := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHttpResponseCode(executeData, returnValue, nil, response_code)
})

// generate by ZifHrtime
var DefZifHrtime = def.DefFunc("hrtime", 1, 1, []def.ArgInfo{{Name: "get_as_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	get_as_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHrtime(executeData, returnValue, get_as_number)
})

// generate by ZifHtmlspecialchars
var DefZifHtmlspecialchars = def.DefFunc("htmlspecialchars", 1, 4, []def.ArgInfo{{Name: "string"}, {Name: "quote_style"}, {Name: "encoding"}, {Name: "double_encode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	string := fp.ParseZval()
	fp.StartOptional()
	quote_style := fp.ParseZval()
	encoding := fp.ParseZval()
	double_encode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHtmlspecialchars(executeData, returnValue, string, nil, quote_style, encoding, double_encode)
})

// generate by ZifHtmlspecialcharsDecode
var DefZifHtmlspecialcharsDecode = def.DefFunc("htmlspecialchars_decode", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "quote_style"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	quote_style_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifHtmlspecialcharsDecode(str, nil, quote_style_)
	returnValue.SetString(ret)
})

// generate by ZifHtmlEntityDecode
var DefZifHtmlEntityDecode = def.DefFunc("html_entity_decode", 1, 3, []def.ArgInfo{{Name: "str"}, {Name: "quote_style"}, {Name: "encoding"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	quote_style_ := fp.ParseLongNullable()
	encoding := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifHtmlEntityDecode(str, nil, quote_style_, encoding)
	returnValue.SetString(ret)
})

// generate by ZifHtmlentities
var DefZifHtmlentities = def.DefFunc("htmlentities", 1, 4, []def.ArgInfo{{Name: "string"}, {Name: "quote_style"}, {Name: "encoding"}, {Name: "double_encode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	string := fp.ParseZval()
	fp.StartOptional()
	quote_style := fp.ParseZval()
	encoding := fp.ParseZval()
	double_encode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHtmlentities(executeData, returnValue, string, nil, quote_style, encoding, double_encode)
})

// generate by ZifGetHtmlTranslationTable
var DefZifGetHtmlTranslationTable = def.DefFunc("get_html_translation_table", 0, 3, []def.ArgInfo{{Name: "table"}, {Name: "quote_style"}, {Name: "encoding"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 3, 0)
	fp.StartOptional()
	table := fp.ParseZval()
	quote_style := fp.ParseZval()
	encoding := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetHtmlTranslationTable(executeData, returnValue, nil, table, quote_style, encoding)
})

// generate by ZifHttpBuildQuery
var DefZifHttpBuildQuery = def.DefFunc("http_build_query", 1, 4, []def.ArgInfo{{Name: "formdata"}, {Name: "prefix"}, {Name: "arg_separator"}, {Name: "enc_type"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	formdata := fp.ParseZval()
	fp.StartOptional()
	prefix := fp.ParseZval()
	arg_separator := fp.ParseZval()
	enc_type := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHttpBuildQuery(executeData, returnValue, formdata, nil, prefix, arg_separator, enc_type)
})

// generate by ZifImageTypeToMimeType
var DefZifImageTypeToMimeType = def.DefFunc("image_type_to_mime_type", 1, 1, []def.ArgInfo{{Name: "imagetype"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	imagetype := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifImageTypeToMimeType(executeData, returnValue, imagetype)
})

// generate by ZifImageTypeToExtension
var DefZifImageTypeToExtension = def.DefFunc("image_type_to_extension", 1, 2, []def.ArgInfo{{Name: "imagetype"}, {Name: "include_dot"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	imagetype := fp.ParseZval()
	fp.StartOptional()
	include_dot := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifImageTypeToExtension(executeData, returnValue, imagetype, nil, include_dot)
})

// generate by ZifGetimagesize
var DefZifGetimagesize = def.DefFunc("getimagesize", 1, 2, []def.ArgInfo{{Name: "imagefile"}, {Name: "info"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	imagefile := fp.ParseZval()
	fp.StartOptional()
	info := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifGetimagesize(executeData, returnValue, imagefile, nil, info)
})

// generate by ZifGetimagesizefromstring
var DefZifGetimagesizefromstring = def.DefFunc("getimagesizefromstring", 1, 2, []def.ArgInfo{{Name: "imagefile"}, {Name: "info"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	imagefile := fp.ParseZval()
	fp.StartOptional()
	info := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifGetimagesizefromstring(executeData, returnValue, imagefile, nil, info)
})

// generate by ZifPhpversion
var DefZifPhpversion = def.DefFunc("phpversion", 0, 1, []def.ArgInfo{{Name: "extension"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	extension := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifPhpversion(nil, extension)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifPhpcredits
var DefZifPhpcredits = def.DefFunc("phpcredits", 0, 1, []def.ArgInfo{{Name: "flag"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	flag := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPhpcredits(executeData, returnValue, nil, flag)
})

// generate by ZifPhpSapiName
var DefZifPhpSapiName = def.DefFunc("php_sapi_name", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifPhpSapiName(executeData, returnValue)
})

// generate by ZifPhpUname
var DefZifPhpUname = def.DefFunc("php_uname", 0, 1, []def.ArgInfo{{Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPhpUname(executeData, returnValue, nil, mode)
})

// generate by ZifPhpIniScannedFiles
var DefZifPhpIniScannedFiles = def.DefFunc("php_ini_scanned_files", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifPhpIniScannedFiles(executeData, returnValue)
})

// generate by ZifPhpIniLoadedFile
var DefZifPhpIniLoadedFile = def.DefFunc("php_ini_loaded_file", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifPhpIniLoadedFile(executeData, returnValue)
})

// generate by ZifIptcembed
var DefZifIptcembed = def.DefFunc("iptcembed", 2, 3, []def.ArgInfo{{Name: "iptcdata"}, {Name: "jpeg_file_name"}, {Name: "spool"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	iptcdata := fp.ParseZval()
	jpeg_file_name := fp.ParseZval()
	fp.StartOptional()
	spool := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIptcembed(executeData, returnValue, iptcdata, jpeg_file_name, nil, spool)
})

// generate by ZifIptcparse
var DefZifIptcparse = def.DefFunc("iptcparse", 1, 1, []def.ArgInfo{{Name: "iptcdata"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	iptcdata := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIptcparse(executeData, returnValue, iptcdata)
})

// generate by ZifLcgValue
var DefZifLcgValue = def.DefFunc("lcg_value", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifLcgValue()
	returnValue.SetDouble(ret)
})

// generate by ZifLevenshtein
var DefZifLevenshtein = def.DefFunc("levenshtein", 2, 5, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}, {Name: "cost_ins"}, {Name: "cost_rep"}, {Name: "cost_del"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 5, 0)
	str1 := fp.ParseZval()
	str2 := fp.ParseZval()
	fp.StartOptional()
	cost_ins := fp.ParseZval()
	cost_rep := fp.ParseZval()
	cost_del := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLevenshtein(executeData, returnValue, str1, str2, nil, cost_ins, cost_rep, cost_del)
})

// generate by ZifReadlink
var DefZifReadlink = def.DefFunc("readlink", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifReadlink(executeData, returnValue, filename)
})

// generate by ZifLinkinfo
var DefZifLinkinfo = def.DefFunc("linkinfo", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParsePathVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifLinkinfo(returnValue, filename)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSymlink
var DefZifSymlink = def.DefFunc("symlink", 2, 2, []def.ArgInfo{{Name: "target"}, {Name: "link"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	target := fp.ParsePathVal()
	link := fp.ParsePathVal()
	if fp.HasError() {
		return
	}
	ret := ZifSymlink(target, link)
	returnValue.SetBool(ret)
})

// generate by ZifLink
var DefZifLink = def.DefFunc("link", 2, 2, []def.ArgInfo{{Name: "target"}, {Name: "link"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	target := fp.ParseZval()
	link := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLink(executeData, returnValue, target, link)
})

// generate by ZifEzmlmHash
var DefZifEzmlmHash = def.DefFunc("ezmlm_hash", 1, 1, []def.ArgInfo{{Name: "addr"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	addr := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifEzmlmHash(executeData, returnValue, addr)
})

// generate by ZifMail
var DefZifMail = def.DefFunc("mail", 3, 5, []def.ArgInfo{{Name: "to"}, {Name: "subject"}, {Name: "message"}, {Name: "additional_headers"}, {Name: "additional_parameters"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 5, 0)
	to := fp.ParseZval()
	subject := fp.ParseZval()
	message := fp.ParseZval()
	fp.StartOptional()
	additional_headers := fp.ParseZval()
	additional_parameters := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMail(executeData, returnValue, to, subject, message, nil, additional_headers, additional_parameters)
})

// generate by ZifAbs
var DefZifAbs = def.DefFunc("abs", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifAbs(number)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifCeil
var DefZifCeil = def.DefFunc("ceil", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifCeil(number)
	if ok {
		returnValue.SetDouble(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifFloor
var DefZifFloor = def.DefFunc("floor", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifFloor(number)
	if ok {
		returnValue.SetDouble(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifRound
var DefZifRound = def.DefFunc("round", 1, 3, []def.ArgInfo{{Name: "number"}, {Name: "precision"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	number := fp.ParseZval()
	fp.StartOptional()
	precision := fp.ParseLong()
	mode_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifRound(number, nil, precision, mode_)
	if ok {
		returnValue.SetDouble(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSin
var DefZifSin = def.DefFunc("sin", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifSin(number)
	returnValue.SetDouble(ret)
})

// generate by ZifCos
var DefZifCos = def.DefFunc("cos", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifCos(number)
	returnValue.SetDouble(ret)
})

// generate by ZifTan
var DefZifTan = def.DefFunc("tan", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifTan(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAsin
var DefZifAsin = def.DefFunc("asin", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAsin(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAcos
var DefZifAcos = def.DefFunc("acos", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAcos(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAtan
var DefZifAtan = def.DefFunc("atan", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAtan(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAtan2
var DefZifAtan2 = def.DefFunc("atan2", 2, 2, []def.ArgInfo{{Name: "y"}, {Name: "x"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	y := fp.ParseDouble()
	x := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAtan2(y, x)
	returnValue.SetDouble(ret)
})

// generate by ZifSinh
var DefZifSinh = def.DefFunc("sinh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifSinh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifCosh
var DefZifCosh = def.DefFunc("cosh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifCosh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifTanh
var DefZifTanh = def.DefFunc("tanh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifTanh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAsinh
var DefZifAsinh = def.DefFunc("asinh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAsinh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAcosh
var DefZifAcosh = def.DefFunc("acosh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAcosh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifAtanh
var DefZifAtanh = def.DefFunc("atanh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifAtanh(number)
	returnValue.SetDouble(ret)
})

// generate by ZifPi
var DefZifPi = def.DefFunc("pi", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifPi()
	returnValue.SetDouble(ret)
})

// generate by ZifIsFinite
var DefZifIsFinite = def.DefFunc("is_finite", 1, 1, []def.ArgInfo{{Name: "val"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	val := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifIsFinite(val)
	returnValue.SetBool(ret)
})

// generate by ZifIsInfinite
var DefZifIsInfinite = def.DefFunc("is_infinite", 1, 1, []def.ArgInfo{{Name: "val"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	val := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifIsInfinite(val)
	returnValue.SetBool(ret)
})

// generate by ZifIsNan
var DefZifIsNan = def.DefFunc("is_nan", 1, 1, []def.ArgInfo{{Name: "val"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	val := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifIsNan(val)
	returnValue.SetBool(ret)
})

// generate by ZifPow
var DefZifPow = def.DefFunc("pow", 2, 2, []def.ArgInfo{{Name: "base"}, {Name: "exponent"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	base := fp.ParseZval()
	exponent := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPow(returnValue, base, exponent)
})

// generate by ZifExp
var DefZifExp = def.DefFunc("exp", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifExp(number)
	returnValue.SetDouble(ret)
})

// generate by ZifExpm1
var DefZifExpm1 = def.DefFunc("expm1", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifExpm1(number)
	returnValue.SetDouble(ret)
})

// generate by ZifLog1p
var DefZifLog1p = def.DefFunc("log1p", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifLog1p(number)
	returnValue.SetDouble(ret)
})

// generate by ZifLog
var DefZifLog = def.DefFunc("log", 1, 2, []def.ArgInfo{{Name: "number"}, {Name: "base"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	number := fp.ParseDouble()
	fp.StartOptional()
	base := fp.ParseDoubleNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifLog(number, nil, base)
	if ok {
		returnValue.SetDouble(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifLog10
var DefZifLog10 = def.DefFunc("log10", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifLog10(number)
	returnValue.SetDouble(ret)
})

// generate by ZifSqrt
var DefZifSqrt = def.DefFunc("sqrt", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifSqrt(number)
	returnValue.SetDouble(ret)
})

// generate by ZifHypot
var DefZifHypot = def.DefFunc("hypot", 2, 2, []def.ArgInfo{{Name: "num1"}, {Name: "num2"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	num1 := fp.ParseDouble()
	num2 := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifHypot(num1, num2)
	returnValue.SetDouble(ret)
})

// generate by ZifDeg2rad
var DefZifDeg2rad = def.DefFunc("deg2rad", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifDeg2rad(number)
	returnValue.SetDouble(ret)
})

// generate by ZifRad2deg
var DefZifRad2deg = def.DefFunc("rad2deg", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifRad2deg(number)
	returnValue.SetDouble(ret)
})

// generate by ZifBindec
var DefZifBindec = def.DefFunc("bindec", 1, 1, []def.ArgInfo{{Name: "binary_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	binary_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifBindec(binary_number)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifHexdec
var DefZifHexdec = def.DefFunc("hexdec", 1, 1, []def.ArgInfo{{Name: "hexadecimal_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hexadecimal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifHexdec(hexadecimal_number)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifOctdec
var DefZifOctdec = def.DefFunc("octdec", 1, 1, []def.ArgInfo{{Name: "octal_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	octal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifOctdec(octal_number)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifDecbin
var DefZifDecbin = def.DefFunc("decbin", 1, 1, []def.ArgInfo{{Name: "decimal_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	decimal_number := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifDecbin(decimal_number)
	returnValue.SetString(ret)
})

// generate by ZifDecoct
var DefZifDecoct = def.DefFunc("decoct", 1, 1, []def.ArgInfo{{Name: "decimal_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	decimal_number := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifDecoct(decimal_number)
	returnValue.SetString(ret)
})

// generate by ZifDechex
var DefZifDechex = def.DefFunc("dechex", 1, 1, []def.ArgInfo{{Name: "decimal_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	decimal_number := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifDechex(decimal_number)
	returnValue.SetString(ret)
})

// generate by ZifBaseConvert
var DefZifBaseConvert = def.DefFunc("base_convert", 3, 3, []def.ArgInfo{{Name: "number"}, {Name: "frombase"}, {Name: "tobase"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	number := fp.ParseZval()
	frombase := fp.ParseLong()
	tobase := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret, ok := ZifBaseConvert(number, frombase, tobase)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifNumberFormat
var DefZifNumberFormat = def.DefFunc("number_format", 1, 4, []def.ArgInfo{{Name: "number"}, {Name: "num_decimal_places"}, {Name: "dec_separator"}, {Name: "thousands_separator"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	number := fp.ParseDouble()
	fp.StartOptional()
	num_decimal_places := fp.ParseLong()
	dec_separator_ := fp.ParseStringValNullable()
	thousands_separator := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifNumberFormat(number, nil, num_decimal_places, dec_separator_, thousands_separator)
	returnValue.SetString(ret)
})

// generate by ZifFmod
var DefZifFmod = def.DefFunc("fmod", 2, 2, []def.ArgInfo{{Name: "x"}, {Name: "y"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	x := fp.ParseDouble()
	y := fp.ParseDouble()
	if fp.HasError() {
		return
	}
	ret := ZifFmod(x, y)
	returnValue.SetDouble(ret)
})

// generate by ZifIntdiv
var DefZifIntdiv = def.DefFunc("intdiv", 2, 2, []def.ArgInfo{{Name: "dividend"}, {Name: "divisor"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	dividend := fp.ParseLong()
	divisor := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifIntdiv(dividend, divisor)
	returnValue.SetLong(ret)
})

// generate by ZifMd5
var DefZifMd5 = def.DefFunc("md5", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "raw_output"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	raw_output := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifMd5(str, nil, raw_output)
	returnValue.SetString(ret)
})

// generate by ZifMd5File
var DefZifMd5File = def.DefFunc("md5_file", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "raw_output"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	filename := fp.ParseStringVal()
	fp.StartOptional()
	raw_output := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifMd5File(filename, nil, raw_output)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifMicrotime
var DefZifMicrotime = def.DefFunc("microtime", 0, 1, []def.ArgInfo{{Name: "get_as_float"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	get_as_float := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifMicrotime(executeData, returnValue, nil, get_as_float)
})

// generate by ZifGettimeofday
var DefZifGettimeofday = def.DefFunc("gettimeofday", 0, 1, []def.ArgInfo{{Name: "get_as_float"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	get_as_float := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGettimeofday(executeData, returnValue, nil, get_as_float)
})

// generate by ZifGetrusage
var DefZifGetrusage = def.DefFunc("getrusage", 0, 1, []def.ArgInfo{{Name: "who"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	who := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetrusage(executeData, returnValue, nil, who)
})

// generate by ZifNetGetInterfaces
var DefZifNetGetInterfaces = def.DefFunc("net_get_interfaces", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifNetGetInterfaces(executeData, returnValue)
})

// generate by ZifPack
var DefZifPack = def.DefFunc("pack", 1, -1, []def.ArgInfo{{Name: "format"}, {Name: "args"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, -1, 0)
	format_ := fp.ParseStringVal()
	fp.StartOptional()
	args := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ret, ok := ZifPack(format_, nil, args)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifUnpack
var DefZifUnpack = def.DefFunc("unpack", 2, 3, []def.ArgInfo{{Name: "format"}, {Name: "input"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	format := fp.ParseZval()
	input := fp.ParseZval()
	fp.StartOptional()
	offset := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUnpack(executeData, returnValue, format, input, nil, offset)
})

// generate by ZifGetmyuid
var DefZifGetmyuid = def.DefFunc("getmyuid", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret, ok := ZifGetmyuid()
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifGetmygid
var DefZifGetmygid = def.DefFunc("getmygid", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret, ok := ZifGetmygid()
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifGetmypid
var DefZifGetmypid = def.DefFunc("getmypid", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifGetmypid()
	returnValue.SetLong(ret)
})

// generate by ZifGetmyinode
var DefZifGetmyinode = def.DefFunc("getmyinode", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret, ok := ZifGetmyinode()
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifGetlastmod
var DefZifGetlastmod = def.DefFunc("getlastmod", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret, ok := ZifGetlastmod(executeData, returnValue)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifPasswordGetInfo
var DefZifPasswordGetInfo = def.DefFunc("password_get_info", 1, 1, []def.ArgInfo{{Name: "hash"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hash_ := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifPasswordGetInfo(hash_)
	returnValue.SetBy(ret)
})

// generate by ZifPasswordNeedsRehash
var DefZifPasswordNeedsRehash = def.DefFunc("password_needs_rehash", 2, 3, []def.ArgInfo{{Name: "hash"}, {Name: "algo"}, {Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	hash_ := fp.ParseStringVal()
	algo_ := fp.ParseZval()
	fp.StartOptional()
	options := fp.ParseArrayOrObjectHt()
	if fp.HasError() {
		return
	}
	ret := ZifPasswordNeedsRehash(hash_, algo_, nil, options)
	returnValue.SetBool(ret)
})

// generate by ZifPasswordVerify
var DefZifPasswordVerify = def.DefFunc("password_verify", 2, 2, []def.ArgInfo{{Name: "password"}, {Name: "hash"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	password := fp.ParseStringVal()
	hash := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifPasswordVerify(password, hash)
	returnValue.SetBool(ret)
})

// generate by ZifPasswordHash
var DefZifPasswordHash = def.DefFunc("password_hash", 2, 3, []def.ArgInfo{{Name: "password"}, {Name: "algo"}, {Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	password := fp.ParseStringVal()
	algo_ := fp.ParseZval()
	fp.StartOptional()
	options := fp.ParseArrayOrObjectHt()
	if fp.HasError() {
		return
	}
	ret := ZifPasswordHash(password, algo_, nil, options)
	returnValue.SetBy(ret)
})

// generate by ZifPasswordAlgos
var DefZifPasswordAlgos = def.DefFunc("password_algos", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifPasswordAlgos()
	returnValue.SetArrayOfString(ret)
})

// generate by ZifQuotedPrintableDecode
var DefZifQuotedPrintableDecode = def.DefFunc("quoted_printable_decode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifQuotedPrintableDecode(str)
	returnValue.SetString(ret)
})

// generate by ZifQuotedPrintableEncode
var DefZifQuotedPrintableEncode = def.DefFunc("quoted_printable_encode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifQuotedPrintableEncode(str)
	returnValue.SetString(ret)
})

// generate by ZifMtSrand
var DefZifMtSrand = def.DefFunc("mt_srand", 0, 2, []def.ArgInfo{{Name: "seed"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	seed_ := fp.ParseLongNullable()
	mode_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ZifMtSrand(nil, seed_, mode_)
})

// generate by ZifMtSrand
var DefZifSrand = def.DefFunc("srand", 0, 2, []def.ArgInfo{{Name: "seed"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	seed_ := fp.ParseLongNullable()
	mode_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ZifMtSrand(nil, seed_, mode_)
})

// generate by ZifMtGetrandmax
var DefZifMtGetrandmax = def.DefFunc("mt_getrandmax", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifMtGetrandmax()
	returnValue.SetLong(ret)
})

// generate by ZifMtGetrandmax
var DefZifGetrandmax = def.DefFunc("getrandmax", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ret := ZifMtGetrandmax()
	returnValue.SetLong(ret)
})

// generate by ZifMtRand
var DefZifMtRand = def.DefFunc("mt_rand", 0, 2, []def.ArgInfo{{Name: "min"}, {Name: "max"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	min_ := fp.ParseLongNullable()
	max_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret, ok := ZifMtRand(nil, min_, max_)
	if ok {
		returnValue.SetLong(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifRand
var DefZifRand = def.DefFunc("rand", 0, 2, []def.ArgInfo{{Name: "min"}, {Name: "max"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	min_ := fp.ParseLongNullable()
	max_ := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifRand(nil, min_, max_)
	returnValue.SetLong(ret)
})

// generate by ZifRandomBytes
var DefZifRandomBytes = def.DefFunc("random_bytes", 1, 1, []def.ArgInfo{{Name: "length"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	length := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifRandomBytes(length)
	returnValue.SetString(ret)
})

// generate by ZifRandomInt
var DefZifRandomInt = def.DefFunc("random_int", 2, 2, []def.ArgInfo{{Name: "min"}, {Name: "max"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	min := fp.ParseLong()
	max := fp.ParseLong()
	if fp.HasError() {
		return
	}
	ret := ZifRandomInt(min, max)
	returnValue.SetLong(ret)
})

// generate by ZifSha1
var DefZifSha1 = def.DefFunc("sha1", 1, 2, []def.ArgInfo{{Name: "str"}, {Name: "raw_output"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	str := fp.ParseStringVal()
	fp.StartOptional()
	raw_output := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifSha1(str, nil, raw_output)
	returnValue.SetString(ret)
})

// generate by ZifSha1File
var DefZifSha1File = def.DefFunc("sha1_file", 1, 2, []def.ArgInfo{{Name: "filename"}, {Name: "raw_output"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	filename := fp.ParseStringVal()
	fp.StartOptional()
	raw_output := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret, ok := ZifSha1File(filename, nil, raw_output)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifSoundex
var DefZifSoundex = def.DefFunc("soundex", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSoundex(executeData, returnValue, str)
})

// generate by ZifStreamSocketPair
var DefZifStreamSocketPair = def.DefFunc("stream_socket_pair", 3, 3, []def.ArgInfo{{Name: "domain"}, {Name: "type"}, {Name: "protocol"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	domain := fp.ParseZval()
	type_ := fp.ParseZval()
	protocol := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketPair(executeData, returnValue, domain, type_, protocol)
})

// generate by ZifStreamSocketClient
var DefZifStreamSocketClient = def.DefFunc("stream_socket_client", 1, 6, []def.ArgInfo{{Name: "remoteaddress"}, {Name: "errcode"}, {Name: "errstring"}, {Name: "timeout"}, {Name: "flags"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 6, 0)
	remoteaddress := fp.ParseZval()
	fp.StartOptional()
	errcode := fp.ParseZvalEx(false, true)
	errstring := fp.ParseZvalEx(false, true)
	timeout := fp.ParseZval()
	flags := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketClient(executeData, returnValue, remoteaddress, nil, errcode, errstring, timeout, flags, context)
})

// generate by ZifStreamSocketServer
var DefZifStreamSocketServer = def.DefFunc("stream_socket_server", 1, 5, []def.ArgInfo{{Name: "localaddress"}, {Name: "errcode"}, {Name: "errstring"}, {Name: "flags"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 5, 0)
	localaddress := fp.ParseZval()
	fp.StartOptional()
	errcode := fp.ParseZvalEx(false, true)
	errstring := fp.ParseZvalEx(false, true)
	flags := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketServer(executeData, returnValue, localaddress, nil, errcode, errstring, flags, context)
})

// generate by ZifStreamSocketAccept
var DefZifStreamSocketAccept = def.DefFunc("stream_socket_accept", 1, 3, []def.ArgInfo{{Name: "serverstream"}, {Name: "timeout"}, {Name: "peername"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	serverstream := fp.ParseZval()
	fp.StartOptional()
	timeout := fp.ParseZval()
	peername := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifStreamSocketAccept(executeData, returnValue, serverstream, nil, timeout, peername)
})

// generate by ZifStreamSocketGetName
var DefZifStreamSocketGetName = def.DefFunc("stream_socket_get_name", 2, 2, []def.ArgInfo{{Name: "stream"}, {Name: "want_peer"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	stream := fp.ParseZval()
	want_peer := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketGetName(executeData, returnValue, stream, want_peer)
})

// generate by ZifStreamSocketSendto
var DefZifStreamSocketSendto = def.DefFunc("stream_socket_sendto", 2, 4, []def.ArgInfo{{Name: "stream"}, {Name: "data"}, {Name: "flags"}, {Name: "target_addr"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream := fp.ParseZval()
	data := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	target_addr := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketSendto(executeData, returnValue, stream, data, nil, flags, target_addr)
})

// generate by ZifStreamSocketRecvfrom
var DefZifStreamSocketRecvfrom = def.DefFunc("stream_socket_recvfrom", 2, 4, []def.ArgInfo{{Name: "stream"}, {Name: "amount"}, {Name: "flags"}, {Name: "remote_addr"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream_ := fp.ParseResource()
	amount := fp.ParseLong()
	fp.StartOptional()
	flags := fp.ParseLong()
	remote_addr := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ret, ok := ZifStreamSocketRecvfrom(executeData, returnValue, stream_, amount, nil, flags, remote_addr)
	if ok {
		returnValue.SetString(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifStreamGetContents
var DefZifStreamGetContents = def.DefFunc("stream_get_contents", 1, 3, []def.ArgInfo{{Name: "source"}, {Name: "maxlen"}, {Name: "offset"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	source := fp.ParseZval()
	fp.StartOptional()
	maxlen := fp.ParseZval()
	offset := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamGetContents(executeData, returnValue, source, nil, maxlen, offset)
})

// generate by ZifStreamCopyToStream
var DefZifStreamCopyToStream = def.DefFunc("stream_copy_to_stream", 2, 4, []def.ArgInfo{{Name: "source"}, {Name: "dest"}, {Name: "maxlen"}, {Name: "pos"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	source := fp.ParseZval()
	dest := fp.ParseZval()
	fp.StartOptional()
	maxlen := fp.ParseZval()
	pos := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamCopyToStream(executeData, returnValue, source, dest, nil, maxlen, pos)
})

// generate by ZifStreamGetMetaData
var DefZifStreamGetMetaData = def.DefFunc("stream_get_meta_data", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamGetMetaData(executeData, returnValue, fp)
})

// generate by ZifStreamGetMetaData
var DefZifSocketGetStatus = def.DefFunc("socket_get_status", 1, 1, []def.ArgInfo{{Name: "fp"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamGetMetaData(executeData, returnValue, fp)
})

// generate by ZifStreamGetTransports
var DefZifStreamGetTransports = def.DefFunc("stream_get_transports", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifStreamGetTransports(executeData, returnValue)
})

// generate by ZifStreamGetWrappers
var DefZifStreamGetWrappers = def.DefFunc("stream_get_wrappers", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifStreamGetWrappers(executeData, returnValue)
})

// generate by ZifStreamSelect
var DefZifStreamSelect = def.DefFunc("stream_select", 4, 5, []def.ArgInfo{{Name: "read_streams"}, {Name: "write_streams"}, {Name: "except_streams"}, {Name: "tv_sec"}, {Name: "tv_usec"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 4, 5, 0)
	read_streams := fp.ParseZvalEx(false, true)
	write_streams := fp.ParseZvalEx(false, true)
	except_streams := fp.ParseZvalEx(false, true)
	tv_sec := fp.ParseZval()
	fp.StartOptional()
	tv_usec := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSelect(executeData, returnValue, read_streams, write_streams, except_streams, tv_sec, nil, tv_usec)
})

// generate by ZifStreamContextGetOptions
var DefZifStreamContextGetOptions = def.DefFunc("stream_context_get_options", 1, 1, []def.ArgInfo{{Name: "stream_or_context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream_or_context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextGetOptions(executeData, returnValue, stream_or_context)
})

// generate by ZifStreamContextSetOption
var DefZifStreamContextSetOption = def.DefFunc("stream_context_set_option", 2, 4, []def.ArgInfo{{Name: "stream_or_context"}, {Name: "wrappername"}, {Name: "optionname"}, {Name: "value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream_or_context := fp.ParseZval()
	wrappername := fp.ParseZval()
	fp.StartOptional()
	optionname := fp.ParseZval()
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextSetOption(executeData, returnValue, stream_or_context, wrappername, nil, optionname, value)
})

// generate by ZifStreamContextSetParams
var DefZifStreamContextSetParams = def.DefFunc("stream_context_set_params", 2, 2, []def.ArgInfo{{Name: "stream_or_context"}, {Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	stream_or_context := fp.ParseZval()
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextSetParams(executeData, returnValue, stream_or_context, options)
})

// generate by ZifStreamContextGetParams
var DefZifStreamContextGetParams = def.DefFunc("stream_context_get_params", 1, 1, []def.ArgInfo{{Name: "stream_or_context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream_or_context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextGetParams(executeData, returnValue, stream_or_context)
})

// generate by ZifStreamContextGetDefault
var DefZifStreamContextGetDefault = def.DefFunc("stream_context_get_default", 0, 1, []def.ArgInfo{{Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextGetDefault(executeData, returnValue, nil, options)
})

// generate by ZifStreamContextSetDefault
var DefZifStreamContextSetDefault = def.DefFunc("stream_context_set_default", 1, 1, []def.ArgInfo{{Name: "options"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextSetDefault(executeData, returnValue, options)
})

// generate by ZifStreamContextCreate
var DefZifStreamContextCreate = def.DefFunc("stream_context_create", 0, 2, []def.ArgInfo{{Name: "options"}, {Name: "params"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	options := fp.ParseZval()
	params := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextCreate(executeData, returnValue, nil, options, params)
})

// generate by ZifStreamFilterPrepend
var DefZifStreamFilterPrepend = def.DefFunc("stream_filter_prepend", 2, 4, []def.ArgInfo{{Name: "stream"}, {Name: "filtername"}, {Name: "read_write"}, {Name: "filterparams"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream := fp.ParseZval()
	filtername := fp.ParseZval()
	fp.StartOptional()
	read_write := fp.ParseZval()
	filterparams := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamFilterPrepend(executeData, returnValue, stream, filtername, nil, read_write, filterparams)
})

// generate by ZifStreamFilterAppend
var DefZifStreamFilterAppend = def.DefFunc("stream_filter_append", 2, 4, []def.ArgInfo{{Name: "stream"}, {Name: "filtername"}, {Name: "read_write"}, {Name: "filterparams"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream := fp.ParseZval()
	filtername := fp.ParseZval()
	fp.StartOptional()
	read_write := fp.ParseZval()
	filterparams := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamFilterAppend(executeData, returnValue, stream, filtername, nil, read_write, filterparams)
})

// generate by ZifStreamFilterRemove
var DefZifStreamFilterRemove = def.DefFunc("stream_filter_remove", 1, 1, []def.ArgInfo{{Name: "stream_filter"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream_filter := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamFilterRemove(executeData, returnValue, stream_filter)
})

// generate by ZifStreamGetLine
var DefZifStreamGetLine = def.DefFunc("stream_get_line", 2, 3, []def.ArgInfo{{Name: "stream"}, {Name: "maxlen"}, {Name: "ending"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	stream := fp.ParseZval()
	maxlen := fp.ParseZval()
	fp.StartOptional()
	ending := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamGetLine(executeData, returnValue, stream, maxlen, nil, ending)
})

// generate by ZifStreamSetBlocking
var DefZifStreamSetBlocking = def.DefFunc("stream_set_blocking", 2, 2, []def.ArgInfo{{Name: "socket"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	socket := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetBlocking(executeData, returnValue, socket, mode)
})

// generate by ZifStreamSetBlocking
var DefZifSocketSetBlocking = def.DefFunc("socket_set_blocking", 2, 2, []def.ArgInfo{{Name: "socket"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	socket := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetBlocking(executeData, returnValue, socket, mode)
})

// generate by ZifStreamSetTimeout
var DefZifStreamSetTimeout = def.DefFunc("stream_set_timeout", 2, 3, []def.ArgInfo{{Name: "stream"}, {Name: "seconds"}, {Name: "microseconds"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	stream := fp.ParseZval()
	seconds := fp.ParseZval()
	fp.StartOptional()
	microseconds := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetTimeout(executeData, returnValue, stream, seconds, nil, microseconds)
})

// generate by ZifStreamSetTimeout
var DefZifSocketSetTimeout = def.DefFunc("socket_set_timeout", 2, 3, []def.ArgInfo{{Name: "stream"}, {Name: "seconds"}, {Name: "microseconds"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	stream := fp.ParseZval()
	seconds := fp.ParseZval()
	fp.StartOptional()
	microseconds := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetTimeout(executeData, returnValue, stream, seconds, nil, microseconds)
})

// generate by ZifStreamSetWriteBuffer
var DefZifStreamSetWriteBuffer = def.DefFunc("stream_set_write_buffer", 2, 2, []def.ArgInfo{{Name: "fp"}, {Name: "buffer"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	buffer := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetWriteBuffer(executeData, returnValue, fp, buffer)
})

// generate by ZifStreamSetWriteBuffer
var DefZifSetFileBuffer = def.DefFunc("set_file_buffer", 2, 2, []def.ArgInfo{{Name: "fp"}, {Name: "buffer"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	buffer := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetWriteBuffer(executeData, returnValue, fp, buffer)
})

// generate by ZifStreamSetChunkSize
var DefZifStreamSetChunkSize = def.DefFunc("stream_set_chunk_size", 2, 2, []def.ArgInfo{{Name: "fp"}, {Name: "chunk_size"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	chunk_size := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetChunkSize(executeData, returnValue, fp, chunk_size)
})

// generate by ZifStreamSetReadBuffer
var DefZifStreamSetReadBuffer = def.DefFunc("stream_set_read_buffer", 2, 2, []def.ArgInfo{{Name: "fp"}, {Name: "buffer"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	buffer := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetReadBuffer(executeData, returnValue, fp, buffer)
})

// generate by ZifStreamSocketEnableCrypto
var DefZifStreamSocketEnableCrypto = def.DefFunc("stream_socket_enable_crypto", 2, 4, []def.ArgInfo{{Name: "stream"}, {Name: "enable"}, {Name: "cryptokind"}, {Name: "sessionstream"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream := fp.ParseZval()
	enable := fp.ParseZval()
	fp.StartOptional()
	cryptokind := fp.ParseZval()
	sessionstream := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketEnableCrypto(executeData, returnValue, stream, enable, nil, cryptokind, sessionstream)
})

// generate by ZifStreamResolveIncludePath
var DefZifStreamResolveIncludePath = def.DefFunc("stream_resolve_include_path", 1, 1, []def.ArgInfo{{Name: "filename"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamResolveIncludePath(executeData, returnValue, filename)
})

// generate by ZifStreamIsLocal
var DefZifStreamIsLocal = def.DefFunc("stream_is_local", 1, 1, []def.ArgInfo{{Name: "stream"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamIsLocal(executeData, returnValue, stream)
})

// generate by ZifStreamSupportsLock
var DefZifStreamSupportsLock = def.DefFunc("stream_supports_lock", 1, 1, []def.ArgInfo{{Name: "stream"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSupportsLock(executeData, returnValue, stream)
})

// generate by ZifStreamIsatty
var DefZifStreamIsatty = def.DefFunc("stream_isatty", 1, 1, []def.ArgInfo{{Name: "stream"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamIsatty(executeData, returnValue, stream)
})

// generate by ZifStreamSocketShutdown
var DefZifStreamSocketShutdown = def.DefFunc("stream_socket_shutdown", 2, 2, []def.ArgInfo{{Name: "stream"}, {Name: "how"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	stream := fp.ParseZval()
	how := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketShutdown(executeData, returnValue, stream, how)
})

// generate by ZifOpenlog
var DefZifOpenlog = def.DefFunc("openlog", 3, 3, []def.ArgInfo{{Name: "ident"}, {Name: "option"}, {Name: "facility"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	ident := fp.ParseZval()
	option := fp.ParseZval()
	facility := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifOpenlog(executeData, returnValue, ident, option, facility)
})

// generate by ZifCloselog
var DefZifCloselog = def.DefFunc("closelog", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifCloselog(executeData, returnValue)
})

// generate by ZifSyslog
var DefZifSyslog = def.DefFunc("syslog", 2, 2, []def.ArgInfo{{Name: "priority"}, {Name: "message"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	priority := fp.ParseZval()
	message := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSyslog(executeData, returnValue, priority, message)
})

// generate by ZifGettype
var DefZifGettype = def.DefFunc("gettype", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifGettype(var_)
	returnValue.SetString(ret)
})

// generate by ZifSettype
var DefZifSettype = def.DefFunc("settype", 2, 2, []def.ArgInfo{{Name: "var"}, {Name: "typ"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	var_ := fp.ParseZvalEx(false, true)
	typ := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifSettype(var_, typ)
	returnValue.SetBool(ret)
})

// generate by ZifIntval
var DefZifIntval = def.DefFunc("intval", 1, 2, []def.ArgInfo{{Name: "var"}, {Name: "base"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	base := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIntval(executeData, returnValue, var_, nil, base)
})

// generate by ZifFloatval
var DefZifFloatval = def.DefFunc("floatval", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFloatval(executeData, returnValue, var_)
})

// generate by ZifFloatval
var DefZifDoubleval = def.DefFunc("doubleval", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFloatval(executeData, returnValue, var_)
})

// generate by ZifBoolval
var DefZifBoolval = def.DefFunc("boolval", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifBoolval(executeData, returnValue, var_)
})

// generate by ZifStrval
var DefZifStrval = def.DefFunc("strval", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStrval(executeData, returnValue, var_)
})

// generate by ZifIsNull
var DefZifIsNull = def.DefFunc("is_null", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsNull(executeData, returnValue, var_)
})

// generate by ZifIsResource
var DefZifIsResource = def.DefFunc("is_resource", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsResource(executeData, returnValue, var_)
})

// generate by ZifIsBool
var DefZifIsBool = def.DefFunc("is_bool", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsBool(executeData, returnValue, var_)
})

// generate by ZifIsInt
var DefZifIsInt = def.DefFunc("is_int", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsInt(executeData, returnValue, var_)
})

// generate by ZifIsInt
var DefZifIsInteger = def.DefFunc("is_integer", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsInt(executeData, returnValue, var_)
})

// generate by ZifIsInt
var DefZifIsLong = def.DefFunc("is_long", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsInt(executeData, returnValue, var_)
})

// generate by ZifIsFloat
var DefZifIsFloat = def.DefFunc("is_float", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsFloat(executeData, returnValue, var_)
})

// generate by ZifIsFloat
var DefZifIsDouble = def.DefFunc("is_double", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsFloat(executeData, returnValue, var_)
})

// generate by ZifIsString
var DefZifIsString = def.DefFunc("is_string", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsString(executeData, returnValue, var_)
})

// generate by ZifIsArray
var DefZifIsArray = def.DefFunc("is_array", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsArray(executeData, returnValue, var_)
})

// generate by ZifIsObject
var DefZifIsObject = def.DefFunc("is_object", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsObject(executeData, returnValue, var_)
})

// generate by ZifIsNumeric
var DefZifIsNumeric = def.DefFunc("is_numeric", 1, 1, []def.ArgInfo{{Name: "value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsNumeric(executeData, returnValue, value)
})

// generate by ZifIsScalar
var DefZifIsScalar = def.DefFunc("is_scalar", 1, 1, []def.ArgInfo{{Name: "value"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsScalar(executeData, returnValue, value)
})

// generate by ZifIsCallable
var DefZifIsCallable = def.DefFunc("is_callable", 1, 3, []def.ArgInfo{{Name: "var"}, {Name: "syntax_only"}, {Name: "callable_name"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	var_ := fp.ParseZval()
	fp.StartOptional()
	syntax_only := fp.ParseZval()
	callable_name := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifIsCallable(executeData, returnValue, var_, nil, syntax_only, callable_name)
})

// generate by ZifIsIterable
var DefZifIsIterable = def.DefFunc("is_iterable", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsIterable(executeData, returnValue, var_)
})

// generate by ZifIsCountable
var DefZifIsCountable = def.DefFunc("is_countable", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsCountable(executeData, returnValue, var_)
})

// generate by ZifUniqid
var DefZifUniqid = def.DefFunc("uniqid", 0, 2, []def.ArgInfo{{Name: "prefix"}, {Name: "more_entropy"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	prefix := fp.ParseStringVal()
	more_entropy := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifUniqid(nil, prefix, more_entropy)
	returnValue.SetString(ret)
})

// generate by ZifParseUrl
var DefZifParseUrl = def.DefFunc("parse_url", 1, 2, []def.ArgInfo{{Name: "url"}, {Name: "component"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	url := fp.ParseStringVal()
	fp.StartOptional()
	component := fp.ParseLongNullable()
	if fp.HasError() {
		return
	}
	ret := ZifParseUrl(url, nil, component)
	returnValue.SetBy(ret)
})

// generate by ZifUrlencode
var DefZifUrlencode = def.DefFunc("urlencode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifUrlencode(str)
	returnValue.SetString(ret)
})

// generate by ZifUrldecode
var DefZifUrldecode = def.DefFunc("urldecode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifUrldecode(str)
	returnValue.SetString(ret)
})

// generate by ZifRawurlencode
var DefZifRawurlencode = def.DefFunc("rawurlencode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRawurlencode(executeData, returnValue, str)
})

// generate by ZifRawurldecode
var DefZifRawurldecode = def.DefFunc("rawurldecode", 1, 1, []def.ArgInfo{{Name: "str"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	str := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifRawurldecode(str)
	returnValue.SetString(ret)
})

// generate by ZifGetHeaders
var DefZifGetHeaders = def.DefFunc("get_headers", 1, 3, []def.ArgInfo{{Name: "url"}, {Name: "format"}, {Name: "context"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	url := fp.ParseZval()
	fp.StartOptional()
	format := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifGetHeaders(executeData, returnValue, url, nil, format, context)
})

// generate by ZifUserFilterNop
var DefZifUserFilterNop = def.DefFunc("user_filter_nop", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifUserFilterNop(executeData, returnValue)
})

// generate by ZifStreamBucketMakeWriteable
var DefZifStreamBucketMakeWriteable = def.DefFunc("stream_bucket_make_writeable", 1, 1, []def.ArgInfo{{Name: "brigade"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	brigade := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamBucketMakeWriteable(executeData, returnValue, brigade)
})

// generate by ZifStreamBucketPrepend
var DefZifStreamBucketPrepend = def.DefFunc("stream_bucket_prepend", 2, 2, []def.ArgInfo{{Name: "brigade"}, {Name: "bucket"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	brigade := fp.ParseZval()
	bucket := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamBucketPrepend(executeData, returnValue, brigade, bucket)
})

// generate by ZifStreamBucketAppend
var DefZifStreamBucketAppend = def.DefFunc("stream_bucket_append", 2, 2, []def.ArgInfo{{Name: "brigade"}, {Name: "bucket"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	brigade := fp.ParseZval()
	bucket := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamBucketAppend(executeData, returnValue, brigade, bucket)
})

// generate by ZifStreamBucketNew
var DefZifStreamBucketNew = def.DefFunc("stream_bucket_new", 2, 2, []def.ArgInfo{{Name: "stream"}, {Name: "buffer"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	stream := fp.ParseZval()
	buffer := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamBucketNew(executeData, returnValue, stream, buffer)
})

// generate by ZifStreamGetFilters
var DefZifStreamGetFilters = def.DefFunc("stream_get_filters", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifStreamGetFilters(executeData, returnValue)
})

// generate by ZifStreamFilterRegister
var DefZifStreamFilterRegister = def.DefFunc("stream_filter_register", 2, 2, []def.ArgInfo{{Name: "filtername"}, {Name: "classname"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filtername := fp.ParseStringVal()
	classname := fp.ParseStringVal()
	if fp.HasError() {
		return
	}
	ret := ZifStreamFilterRegister(filtername, classname)
	returnValue.SetBool(ret)
})

// generate by ZifConvertUuencode
var DefZifConvertUuencode = def.DefFunc("convert_uuencode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifConvertUuencode(executeData, returnValue, data)
})

// generate by ZifConvertUudecode
var DefZifConvertUudecode = def.DefFunc("convert_uudecode", 1, 1, []def.ArgInfo{{Name: "data"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	data := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifConvertUudecode(executeData, returnValue, data)
})

// generate by ZifVarDump
var DefZifVarDump = def.DefFunc("var_dump", 0, -1, []def.ArgInfo{{Name: "vars"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	vars := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifVarDump(vars)
})

// generate by ZifDebugZvalDump
var DefZifDebugZvalDump = def.DefFunc("debug_zval_dump", 0, -1, []def.ArgInfo{{Name: "vars"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, -1, 0)
	vars := fp.ParseVariadic()
	if fp.HasError() {
		return
	}
	ZifDebugZvalDump(vars)
})

// generate by ZifVarExport
var DefZifVarExport = def.DefFunc("var_export", 1, 2, []def.ArgInfo{{Name: "var"}, {Name: "return"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	var__ := fp.ParseZval()
	fp.StartOptional()
	return_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifVarExport(executeData, returnValue, var__, nil, return_)
})

// generate by ZifSerialize
var DefZifSerialize = def.DefFunc("serialize", 1, 1, []def.ArgInfo{{Name: "var"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	var_ := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret := ZifSerialize(var_)
	returnValue.SetBy(ret)
})

// generate by ZifUnserialize
var DefZifUnserialize = def.DefFunc("unserialize", 1, 2, []def.ArgInfo{{Name: "variable_representation"}, {Name: "allowed_classes"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	variable_representation := fp.ParseZval()
	fp.StartOptional()
	allowed_classes := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifUnserialize(executeData, returnValue, variable_representation, nil, allowed_classes)
})

// generate by ZifMemoryGetUsage
var DefZifMemoryGetUsage = def.DefFunc("memory_get_usage", 0, 1, []def.ArgInfo{{Name: "real_usage"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	real_usage := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifMemoryGetUsage(nil, real_usage)
	returnValue.SetLong(ret)
})

// generate by ZifMemoryGetPeakUsage
var DefZifMemoryGetPeakUsage = def.DefFunc("memory_get_peak_usage", 0, 1, []def.ArgInfo{{Name: "real_usage"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	real_usage := fp.ParseBoolVal()
	if fp.HasError() {
		return
	}
	ret := ZifMemoryGetPeakUsage(nil, real_usage)
	returnValue.SetLong(ret)
})

// generate by ZifVersionCompare
var DefZifVersionCompare = def.DefFunc("version_compare", 2, 3, []def.ArgInfo{{Name: "ver1"}, {Name: "ver2"}, {Name: "oper"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	ver1 := fp.ParseZval()
	ver2 := fp.ParseZval()
	fp.StartOptional()
	oper := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifVersionCompare(ver1, ver2, nil, oper)
	returnValue.SetBy(ret)
})
