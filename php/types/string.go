package types

import (
	"math"
)

/**
 * String
 */
type String string

func NewString(str string) *String { tmp := String(str); return &tmp }

func (zs *String) GetStr() string {
	if zs == nil {
		return ""
	}
	return string(*zs)
}
func (zs *String) GetLen() int {
	if zs == nil {
		return 0
	}
	return len(*zs)
}

/**
 * String Constants
 */
const STR_MAX_LEN = math.MaxInt

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
	STR_ARGV                 = "argv"
	STR_ARGC                 = "argc"
)
