// <<generate>>

package zend

/**
 * ZendConstant
 */
type ZendConstant struct {
	value Zval
	name  string
}

func NewZendConstant(name string) *ZendConstant {
	return &ZendConstant{name: name}
}

func (this *ZendConstant) Value() *Zval                 { return &this.value }
func (this *ZendConstant) Name() string                 { return this.name }
func (this *ZendConstant) SetNameStr(value *ZendString) { this.name = value.GetStr() }

func (this *ZendConstant) GetValue() Zval            { return this.value }
func (this *ZendConstant) GetName() *ZendString      { return NewZendString(this.name) }
func (this *ZendConstant) SetName(value *ZendString) { this.name = value.GetStr() }
