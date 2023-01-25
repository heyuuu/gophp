// <<generate>>

package zend

type ZendStringCopyStorageFuncT func()
type ZendNewInternedStringFuncT func(str *ZendString) *ZendString
type ZendStringInitInternedFuncT func(str *byte, size int, permanent int) *ZendString

var ZendNewInternedString ZendNewInternedStringFuncT
var ZendStringInitInterned ZendStringInitInternedFuncT
var ZendOneCharString []*ZendString

const _STR_HEADER_SIZE = _ZSTR_HEADER_SIZE
const _ZSTR_HEADER_SIZE = zend_long((*byte)(&((*ZendString)(nil).GetVal())) - (*byte)(nil))
const ZSTR_MAX_OVERHEAD = ZEND_MM_ALIGNED_SIZE(_ZSTR_HEADER_SIZE + 1)
const ZSTR_MAX_LEN = SIZE_MAX - ZSTR_MAX_OVERHEAD

type ZendKnownStringId = int

const (
	ZEND_STR_FILE = iota
	ZEND_STR_LINE
	ZEND_STR_FUNCTION
	ZEND_STR_CLASS
	ZEND_STR_OBJECT
	ZEND_STR_TYPE
	ZEND_STR_OBJECT_OPERATOR
	ZEND_STR_PAAMAYIM_NEKUDOTAYIM
	ZEND_STR_ARGS
	ZEND_STR_UNKNOWN
	ZEND_STR_EVAL
	ZEND_STR_INCLUDE
	ZEND_STR_REQUIRE
	ZEND_STR_INCLUDE_ONCE
	ZEND_STR_REQUIRE_ONCE
	ZEND_STR_SCALAR
	ZEND_STR_ERROR_REPORTING
	ZEND_STR_STATIC
	ZEND_STR_THIS
	ZEND_STR_VALUE
	ZEND_STR_KEY
	ZEND_STR_MAGIC_AUTOLOAD
	ZEND_STR_MAGIC_INVOKE
	ZEND_STR_PREVIOUS
	ZEND_STR_CODE
	ZEND_STR_MESSAGE
	ZEND_STR_SEVERITY
	ZEND_STR_STRING
	ZEND_STR_TRACE
	ZEND_STR_SCHEME
	ZEND_STR_HOST
	ZEND_STR_PORT
	ZEND_STR_USER
	ZEND_STR_PASS
	ZEND_STR_PATH
	ZEND_STR_QUERY
	ZEND_STR_FRAGMENT
	ZEND_STR_NULL
	ZEND_STR_BOOLEAN
	ZEND_STR_INTEGER
	ZEND_STR_DOUBLE
	ZEND_STR_ARRAY
	ZEND_STR_RESOURCE
	ZEND_STR_CLOSED_RESOURCE
	ZEND_STR_NAME
	ZEND_STR_ARGV
	ZEND_STR_ARGC
	ZEND_STR_ARRAY_CAPITALIZED
	ZEND_STR_LAST_KNOWN
)

var InternedStringsPermanent HashTable
var InternedStringRequestHandler ZendNewInternedStringFuncT = ZendNewInternedStringRequest
var InternedStringInitRequestHandler ZendStringInitInternedFuncT = ZendStringInitInternedRequest
var ZendEmptyString *ZendString = nil
var ZendKnownStrings **ZendString = nil
var KnownStrings []*byte = []*byte{"file", "line", "function", "class", "object", "type", "->", "::", "args", "unknown", "eval", "include", "require", "include_once", "require_once", "scalar", "error_reporting", "static", "this", "value", "key", "__autoload", "__invoke", "previous", "code", "message", "severity", "string", "trace", "scheme", "host", "port", "user", "pass", "path", "query", "fragment", "NULL", "boolean", "integer", "double", "array", "resource", "resource (closed)", "name", "argv", "argc", "Array", nil}
