// <<generate>>

package standard

import (
	"sik/core"
	r "sik/runtime"
	"sik/zend"
)

var ZifFdSet func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var ZifFdIsset func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var ZmStartupUserStreams func(type_ int, module_number int) int
var PhpSetSockBlocking func(socketd core.PhpSocketT, block int) int

const PHP_CSV_NO_ESCAPE = r.EOF
const META_DEF_BUFSIZE = 8192
const PHP_FILE_USE_INCLUDE_PATH = 1
const PHP_FILE_IGNORE_NEW_LINES = 2
const PHP_FILE_SKIP_EMPTY_LINES = 4
const PHP_FILE_APPEND = 8
const PHP_FILE_NO_DEFAULT_CONTEXT = 16

type PhpMetaTagsToken = int

const (
	TOK_EOF = 0
	TOK_OPENTAG
	TOK_CLOSETAG
	TOK_SLASH
	TOK_EQUAL
	TOK_SPACE
	TOK_ID
	TOK_STRING
	TOK_OTHER
)

var FileGlobals PhpFileGlobals
var LeStreamContext int = zend.FAILURE
var FlockValues []int = []int{LOCK_SH, LOCK_EX, LOCK_UN}

const PHP_META_UNSAFE = ".\\+*?[^]$() "
const PHP_FILE_BUF_SIZE = 80
const PHP_META_HTML401_CHARS = "-_.:"
