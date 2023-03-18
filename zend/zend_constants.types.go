// <<generate>>

package zend

import "sik/zend/types"

/**
 * ZendConstant
 */
type ZendConstant struct {
	value types.Zval
	name  string
}

func NewZendConstant(name string) *ZendConstant {
	return &ZendConstant{name: name}
}

func (this *ZendConstant) Value() *types.Zval                 { return &this.value }
func (this *ZendConstant) Name() string                       { return this.name }
func (this *ZendConstant) SetNameStr(value *types.ZendString) { this.name = value.GetStr() }

func (this *ZendConstant) GetValue() types.Zval            { return this.value }
func (this *ZendConstant) GetName() *types.ZendString      { return types.NewZendString(this.name) }
func (this *ZendConstant) SetName(value *types.ZendString) { this.name = value.GetStr() }
