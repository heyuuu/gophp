package core

import (
	"github.com/heyuuu/gophp/zend/types"
)

var OutputGlobals ZendOutputGlobals
var PhpOutputDefaultHandlerName = "default output handler"
var PhpOutputDevnullHandlerName = "null output handler"

var PhpOutputHandlerAliases types.Array
var PhpOutputHandlerConflicts types.Array
var PhpOutputHandlerReverseConflicts types.Array
var PhpOutputDirect func(str *byte, str_len int) int = PhpOutputStderr
