// <<generate>>

package types

import (
	"math"
)

var InternedStringsPermanent = NewInternedStrings()

const ZSTR_MAX_LEN = math.MaxInt

/**
 * 预初始化常用 ZendString
 */
// ZendEmptyString
var ZSTR_EMPTY = initZendString("")

// ZendOneCharString
var oneCharStrings []*ZendString

func init() {
	for i := 0; i < 256; i++ {
		oneCharStrings[i] = initZendString(string([]byte{byte(i)}))
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
