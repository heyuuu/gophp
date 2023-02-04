// <<generate>>

package zend

import "math"

type ZendNewInternedStringFuncT func(str *ZendString) *ZendString
type ZendStringInitInternedFuncT func(str *byte, size int, permanent int) *ZendString

var ZendNewInternedString ZendNewInternedStringFuncT
var ZendStringInitInterned ZendStringInitInternedFuncT
var ZendOneCharString []*ZendString

const ZSTR_MAX_LEN = math.MaxInt

type ZendKnownStringId = string
type ZendKnownString = string

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

var InternedStringsPermanent HashTable
var InternedStringRequestHandler ZendNewInternedStringFuncT = ZendNewInternedStringRequest
var InternedStringInitRequestHandler ZendStringInitInternedFuncT = ZendStringInitInternedRequest
