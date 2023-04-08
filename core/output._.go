package core

var OutputGlobals ZendOutputGlobals
var PhpOutputDefaultHandlerName = "default output handler"

var PhpOutputDirect func(str string) int = PhpOutputStderr
