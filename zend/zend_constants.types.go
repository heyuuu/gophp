// <<generate>>

package zend

/**
 * ZendConstant
 */
type ZendConstant struct {
	value Zval
	name  *ZendString
}

// func MakeZendConstant(value Zval, name *ZendString) ZendConstant {
//     return ZendConstant{
//         value:value,
//         name:name,
//     }
// }
func (this *ZendConstant) GetValue() Zval { return this.value }

// func (this *ZendConstant) SetValue(value Zval) { this.value = value }
func (this *ZendConstant) GetName() *ZendString      { return this.name }
func (this *ZendConstant) SetName(value *ZendString) { this.name = value }
