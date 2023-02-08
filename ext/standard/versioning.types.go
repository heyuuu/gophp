// <<generate>>

package standard

/**
 * SpecialFormsT
 */
type SpecialFormsT struct {
	name  *byte
	order int
}

// func NewSpecialFormsT(name *byte, order int) *SpecialFormsT {
//     return &SpecialFormsT{
//         name:name,
//         order:order,
//     }
// }
func MakeSpecialFormsT(name *byte, order int) SpecialFormsT {
	return SpecialFormsT{
		name:  name,
		order: order,
	}
}
func (this *SpecialFormsT) GetName() *byte { return this.name }

// func (this *SpecialFormsT) SetName(value *byte) { this.name = value }
func (this *SpecialFormsT) GetOrder() int { return this.order }

// func (this *SpecialFormsT) SetOrder(value int) { this.order = value }
