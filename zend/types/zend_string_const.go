// <<generate>>

package types

import (
	"math"
)

var ZendOneCharString []*ZendString

var InternedStringsPermanent = NewInternedStrings()

const ZSTR_MAX_LEN = math.MaxInt

type ZendKnownStringId = string
type ZendKnownString = string

// ZendEmptyString
var ZSTR_EMPTY = initZendString("")

// ZendOneCharString
var oneCharStrings []*ZendString

func init() {
	for i := 0; i < 256; i++ {
		ZendOneCharString[i] = initZendString(string([]byte{byte(i)}))
	}
}

// ZendKnownString
var (
	ZSTR_FILE                 = initZendString("file")
	ZSTR_LINE                 = initZendString("line")
	ZSTR_FUNCTION             = initZendString("function")
	ZSTR_CLASS                = initZendString("class")
	ZSTR_OBJECT               = initZendString("object")
	ZSTR_TYPE                 = initZendString("type")
	ZSTR_OBJECT_OPERATOR      = initZendString("->")
	ZSTR_PAAMAYIM_NEKUDOTAYIM = initZendString("::")
	ZSTR_ARGS                 = initZendString("args")
	ZSTR_UNKNOWN              = initZendString("unknown")
	ZSTR_EVAL                 = initZendString("eval")
	ZSTR_INCLUDE              = initZendString("include")
	ZSTR_REQUIRE              = initZendString("require")
	ZSTR_INCLUDE_ONCE         = initZendString("include_once")
	ZSTR_REQUIRE_ONCE         = initZendString("require_once")
	ZSTR_SCALAR               = initZendString("scalar")
	ZSTR_ERROR_REPORTING      = initZendString("error_reporting")
	ZSTR_STATIC               = initZendString("static")
	ZSTR_THIS                 = initZendString("this")
	ZSTR_VALUE                = initZendString("value")
	ZSTR_KEY                  = initZendString("key")
	ZSTR_MAGIC_AUTOLOAD       = initZendString("__autoload")
	ZSTR_MAGIC_INVOKE         = initZendString("__invoke")
	ZSTR_PREVIOUS             = initZendString("previous")
	ZSTR_CODE                 = initZendString("code")
	ZSTR_MESSAGE              = initZendString("message")
	ZSTR_SEVERITY             = initZendString("severity")
	ZSTR_STRING               = initZendString("string")
	ZSTR_TRACE                = initZendString("trace")
	ZSTR_SCHEME               = initZendString("scheme")
	ZSTR_HOST                 = initZendString("host")
	ZSTR_PORT                 = initZendString("port")
	ZSTR_USER                 = initZendString("user")
	ZSTR_PASS                 = initZendString("pass")
	ZSTR_PATH                 = initZendString("path")
	ZSTR_QUERY                = initZendString("query")
	ZSTR_FRAGMENT             = initZendString("fragment")
	ZSTR_NULL                 = initZendString("NULL")
	ZSTR_BOOLEAN              = initZendString("boolean")
	ZSTR_INTEGER              = initZendString("integer")
	ZSTR_DOUBLE               = initZendString("double")
	ZSTR_ARRAY                = initZendString("array")
	ZSTR_RESOURCE             = initZendString("resource")
	ZSTR_CLOSED_RESOURCE      = initZendString("resource (closed)")
	ZSTR_NAME                 = initZendString("name")
	ZSTR_ARGV                 = initZendString("argv")
	ZSTR_ARGC                 = initZendString("argc")
	ZSTR_ARRAY_CAPITALIZED    = initZendString("Array")
)

const (
	ZEND_STR_FILE                 ZendKnownString = "file"
	ZEND_STR_LINE                                 = "line"
	ZEND_STR_FUNCTION                             = "function"
	ZEND_STR_CLASS                                = "class"
	ZEND_STR_OBJECT                               = "object"
	ZEND_STR_TYPE                                 = "type"
	ZEND_STR_OBJECT_OPERATOR                      = "->"
	ZEND_STR_PAAMAYIM_NEKUDOTAYIM                 = "::"
	ZEND_STR_ARGS                                 = "args"
	ZEND_STR_UNKNOWN                              = "unknown"
	ZEND_STR_EVAL                                 = "eval"
	ZEND_STR_INCLUDE                              = "include"
	ZEND_STR_REQUIRE                              = "require"
	ZEND_STR_INCLUDE_ONCE                         = "include_once"
	ZEND_STR_REQUIRE_ONCE                         = "require_once"
	ZEND_STR_SCALAR                               = "scalar"
	ZEND_STR_ERROR_REPORTING                      = "error_reporting"
	ZEND_STR_STATIC                               = "static"
	ZEND_STR_THIS                                 = "this"
	ZEND_STR_VALUE                                = "value"
	ZEND_STR_KEY                                  = "key"
	ZEND_STR_MAGIC_AUTOLOAD                       = "__autoload"
	ZEND_STR_MAGIC_INVOKE                         = "__invoke"
	ZEND_STR_PREVIOUS                             = "previous"
	ZEND_STR_CODE                                 = "code"
	ZEND_STR_MESSAGE                              = "message"
	ZEND_STR_SEVERITY                             = "severity"
	ZEND_STR_STRING                               = "string"
	ZEND_STR_TRACE                                = "trace"
	ZEND_STR_SCHEME                               = "scheme"
	ZEND_STR_HOST                                 = "host"
	ZEND_STR_PORT                                 = "port"
	ZEND_STR_USER                                 = "user"
	ZEND_STR_PASS                                 = "pass"
	ZEND_STR_PATH                                 = "path"
	ZEND_STR_QUERY                                = "query"
	ZEND_STR_FRAGMENT                             = "fragment"
	ZEND_STR_NULL                                 = "NULL"
	ZEND_STR_BOOLEAN                              = "boolean"
	ZEND_STR_INTEGER                              = "integer"
	ZEND_STR_DOUBLE                               = "double"
	ZEND_STR_ARRAY                                = "array"
	ZEND_STR_RESOURCE                             = "resource"
	ZEND_STR_CLOSED_RESOURCE                      = "resource (closed)"
	ZEND_STR_NAME                                 = "name"
	ZEND_STR_ARGV                                 = "argv"
	ZEND_STR_ARGC                                 = "argc"
	ZEND_STR_ARRAY_CAPITALIZED                    = "Array"
)
