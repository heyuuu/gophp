package types

type ConstFlag uint8

const (
	// Case Sensitive
	// 大小写敏感，默认是开启的，用户通过define()定义的始终是区分大小 写的，通过扩展定义的可以自由选择
	ConstCs ConstFlag = 1 << 0

	// Persistent
	// 持久化的，只有通过扩展、内核定义的才支持，这种常量不会在request结束时清理掉
	ConstPersistent ConstFlag = 1 << 1

	// Allow compile-time substitution
	// 允许编译时替换，编译时如果发现有地方在读取常量的值，那么编译器会尝试直接替换为常量值，而不是在执行时再去读取，目前这个flag只有TRUE、 FALSE、NULL三个常量在使用
	ConstCtSubst ConstFlag = 1 << 2

	// Can't be saved in file cache
	ConstNoFileCache ConstFlag = 1 << 3
)

// Constant
type Constant struct {
	name         string
	value        Zval
	flags        ConstFlag
	moduleNumber int
}

func NewConstant(name string, value Zval, flags ConstFlag, moduleNumber int) *Constant {
	return &Constant{name: name, value: value, flags: flags, moduleNumber: moduleNumber}
}

func (c Constant) Name() string      { return c.name }
func (c Constant) Value() Zval       { return c.value }
func (c Constant) ModuleNumber() int { return c.moduleNumber }

func (c Constant) IsCaseSensitive() bool { return c.flags&ConstCs != 0 }
func (c Constant) IsPersistent() bool    { return c.flags&ConstPersistent != 0 }
func (c Constant) IsCtSubst() bool       { return c.flags&ConstCtSubst != 0 }
func (c Constant) IsNoFileCache() bool   { return c.flags&ConstNoFileCache != 0 }

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
