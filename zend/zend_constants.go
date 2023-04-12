package zend

import "github.com/heyuuu/gophp/zend/types"

/**
 * constants
 */
const CONST_CS = 1 << 0
const CONST_PERSISTENT = 1 << 1
const CONST_CT_SUBST = 1 << 2
const CONST_NO_FILE_CACHE = 1 << 3

const PHP_USER_CONSTANT = 0x7fffff
const ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK = 0x1000

var ZEND_CONSTANT_DTOR types.DtorFuncT = FreeZendConstant

const IS_CONSTANT_VISITED_MARK = 0x80

/**
 * ZendConstant
 */
type ZendConstant struct {
	value types.Zval
	name  string
}

func NewConstant(name string, flags int, moduleNumber int) *ZendConstant {
	c := &ZendConstant{name: name}

	realFlags := uint32(flags&0xff | moduleNumber<<8)
	c.value.SetConstantFlags(realFlags)

	return c
}

func CopyConstant(c *ZendConstant) *ZendConstant {
	return &ZendConstant{
		value: c.value,
		name:  c.name,
	}
}

func (c *ZendConstant) Name() string       { return c.name }
func (c *ZendConstant) Value() *types.Zval { return &c.value }
func (c *ZendConstant) Flags() uint8       { return uint8(c.value.GetConstantFlags() & 0xff) }
func (c *ZendConstant) ModuleNumber() int  { return int(c.value.GetConstantFlags() << 8) }

func (c *ZendConstant) GetName() *types.String     { return types.NewString(c.name) }
func (c *ZendConstant) SetName(name *types.String) { c.name = name.GetStr() }
func (c *ZendConstant) SetNameVal(name string)     { c.name = name }

func (c *ZendConstant) IsPersistent() bool {
	return (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) != 0
}

/**
 * functions
 */
