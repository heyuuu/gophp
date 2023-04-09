package types

import (
	"math"
)

var InternedStringsPermanent = NewInternedStrings()

const ZSTR_MAX_LEN = math.MaxInt

// ZendOneCharString
var oneCharStrings []*String

func init() {
	for i := 0; i < 256; i++ {
		oneCharStrings[i] = initString(string([]byte{byte(i)}))
	}
}

// ZendKnownString
var (
	STR_FILE                 = "file"
	STR_LINE                 = "line"
	STR_FUNCTION             = "function"
	STR_CLASS                = "class"
	STR_OBJECT               = "object"
	STR_TYPE                 = "type"
	STR_OBJECT_OPERATOR      = "->"
	STR_PAAMAYIM_NEKUDOTAYIM = "::"
	STR_ARGS                 = "args"
	STR_UNKNOWN              = "unknown"
	STR_EVAL                 = "eval"
	STR_INCLUDE              = "include"
	STR_REQUIRE              = "require"
	STR_INCLUDE_ONCE         = "include_once"
	STR_REQUIRE_ONCE         = "require_once"
	STR_SCALAR               = "scalar"
	STR_ERROR_REPORTING      = "error_reporting"
	STR_STATIC               = "static"
	STR_THIS                 = "this"
	STR_VALUE                = "value"
	STR_KEY                  = "key"
	STR_MAGIC_AUTOLOAD       = "__autoload"
	STR_MAGIC_INVOKE         = "__invoke"
	STR_PREVIOUS             = "previous"
	STR_CODE                 = "code"
	STR_MESSAGE              = "message"
	STR_SEVERITY             = "severity"
	STR_STRING               = "string"
	STR_TRACE                = "trace"
	STR_SCHEME               = "scheme"
	STR_HOST                 = "host"
	STR_PORT                 = "port"
	STR_USER                 = "user"
	STR_PASS                 = "pass"
	STR_PATH                 = "path"
	STR_QUERY                = "query"
	STR_FRAGMENT             = "fragment"
	STR_ARGV                 = "argv"
	STR_ARGC                 = "argc"
	STR_ARRAY_CAPITALIZED    = "Array"
)
