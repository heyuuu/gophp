package zend

import "sik/zend/types"

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

func NewZendConstant(name string, flags int, moduleNumber int) *ZendConstant {
	c := &ZendConstant{name: name}

	realFlags := uint32(flags&0xff | moduleNumber<<8)
	c.value.SetConstantFlags(realFlags)

	return c
}

func (this *ZendConstant) Name() string       { return this.name }
func (this *ZendConstant) Value() *types.Zval { return &this.value }
func (this *ZendConstant) Flags() uint8       { return uint8(this.value.GetConstantFlags() & 0xff) }
func (this *ZendConstant) ModuleNumber() int  { return int(this.value.GetConstantFlags() << 8) }

func (this *ZendConstant) GetName() *types.ZendString      { return types.NewZendString(this.name) }
func (this *ZendConstant) SetName(value *types.ZendString) { this.name = value.GetStr() }

/**
 * functions
 */
