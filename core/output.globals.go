// <<generate>>

package core

import (
	"sik/zend"
)

const PHP_OUTPUT_DEBUG = 0
const PHP_OUTPUT_NOINLINE = 0

var OutputGlobals ZendOutputGlobals
var PhpOutputDefaultHandlerName []byte = "default output handler"
var PhpOutputDevnullHandlerName []byte = "null output handler"
var PhpOutputHandlerAliases zend.HashTable
var PhpOutputHandlerConflicts zend.HashTable
var PhpOutputHandlerReverseConflicts zend.HashTable
var PhpOutputDirect func(str *byte, str_len int) int = PhpOutputStderr
