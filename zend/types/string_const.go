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
	ZSTR_FILE                 = initString("file")
	ZSTR_LINE                 = initString("line")
	ZSTR_FUNCTION             = initString("function")
	ZSTR_CLASS                = initString("class")
	ZSTR_OBJECT               = initString("object")
	ZSTR_TYPE                 = initString("type")
	ZSTR_OBJECT_OPERATOR      = initString("->")
	ZSTR_PAAMAYIM_NEKUDOTAYIM = initString("::")
	ZSTR_ARGS                 = initString("args")
	ZSTR_UNKNOWN              = initString("unknown")
	ZSTR_EVAL                 = initString("eval")
	ZSTR_INCLUDE              = initString("include")
	ZSTR_REQUIRE              = initString("require")
	ZSTR_INCLUDE_ONCE         = initString("include_once")
	ZSTR_REQUIRE_ONCE         = initString("require_once")
	ZSTR_SCALAR               = initString("scalar")
	ZSTR_ERROR_REPORTING      = initString("error_reporting")
	ZSTR_STATIC               = initString("static")
	ZSTR_THIS                 = initString("this")
	ZSTR_VALUE                = initString("value")
	ZSTR_KEY                  = initString("key")
	ZSTR_MAGIC_AUTOLOAD       = initString("__autoload")
	ZSTR_MAGIC_INVOKE         = initString("__invoke")
	ZSTR_PREVIOUS             = initString("previous")
	ZSTR_CODE                 = initString("code")
	ZSTR_MESSAGE              = initString("message")
	ZSTR_SEVERITY             = initString("severity")
	ZSTR_STRING               = initString("string")
	ZSTR_TRACE                = initString("trace")
	ZSTR_SCHEME               = initString("scheme")
	ZSTR_HOST                 = initString("host")
	ZSTR_PORT                 = initString("port")
	ZSTR_USER                 = initString("user")
	ZSTR_PASS                 = initString("pass")
	ZSTR_PATH                 = initString("path")
	ZSTR_QUERY                = initString("query")
	ZSTR_FRAGMENT             = initString("fragment")
	ZSTR_NULL                 = initString("NULL")
	ZSTR_BOOLEAN              = initString("boolean")
	ZSTR_INTEGER              = initString("integer")
	ZSTR_DOUBLE               = initString("double")
	ZSTR_ARRAY                = initString("array")
	ZSTR_RESOURCE             = initString("resource")
	ZSTR_CLOSED_RESOURCE      = initString("resource (closed)")
	ZSTR_NAME                 = initString("name")
	ZSTR_ARGV                 = initString("argv")
	ZSTR_ARGC                 = initString("argc")
	ZSTR_ARRAY_CAPITALIZED    = initString("Array")
)
