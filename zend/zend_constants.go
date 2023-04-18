package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * constants
 */
const (
	// 大小写敏感，默认是开启的，用户通过define()定义的始终是区分大小 写的，通过扩展定义的可以自由选择
	CONST_CS = 1 << 0

	// 持久化的，只有通过扩展、内核定义的才支持，这种常量不 会在request结束时清理掉
	CONST_PERSISTENT = 1 << 1

	// 允许编译时替换，编译时如果发现有地方在读取常量的值，那么编 译器会尝试直接替换为常量值，而不是在执行时再去读取，目前这个flag只有TRUE、 FALSE、NULL三个常量在使用
	CONST_CT_SUBST = 1 << 2

	CONST_NO_FILE_CACHE = 1 << 3
)

const PHP_USER_CONSTANT = 0x7fffff
const ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK = 0x1000
const IS_CONSTANT_VISITED_MARK = 0x80

/**
 * ZendConstant
 */
type ZendConstant struct {
	value types.Zval
	name  string
	// 标识符(原 value.u2.constant_flags)
	flags uint32
}

func NewConstant(name string, flags int, moduleNumber int) *ZendConstant {
	c := &ZendConstant{name: name}
	c.SetFlags(flags, moduleNumber)
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
func (c *ZendConstant) Flags() uint8       { return uint8(c.flags & 0xff) }
func (c *ZendConstant) ModuleNumber() int  { return int(c.flags >> 8) }

func (c *ZendConstant) GetName() *types.String { return types.NewString(c.name) }
func (c *ZendConstant) SetName(name string)    { c.name = name }
func (c *ZendConstant) SetFlags(flags int, moduleNumber int) {
	realFlags := uint32(flags&0xff | moduleNumber<<8)
	c.flags = realFlags
}

func (c *ZendConstant) IsCaseSensitive() bool {
	return (c.Flags() & CONST_CS) != 0
}
func (c *ZendConstant) IsPersistent() bool {
	return (c.Flags() & CONST_PERSISTENT) != 0
}
func (c *ZendConstant) IsCtSubst() bool {
	return (c.Flags() & CONST_CT_SUBST) != 0
}
func (c *ZendConstant) IsNoFileCache() bool {
	return (c.Flags() & CONST_NO_FILE_CACHE) != 0
}
