package core

var OutputGlobals ZendOutputGlobals
var PhpOutputDefaultHandlerName = "default output handler"

var PhpOutputDirect func(str *byte, str_len int) int = PhpOutputStderr
