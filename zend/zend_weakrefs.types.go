// <<generate>>

package zend

/**
 * ZendWeakref
 */
type ZendWeakref struct {
	referent *ZendObject
	std      ZendObject
}

// func NewZendWeakref(referent *ZendObject, std ZendObject) *ZendWeakref {
//     return &ZendWeakref{
//         referent:referent,
//         std:std,
//     }
// }
// func MakeZendWeakref(referent *ZendObject, std ZendObject) ZendWeakref {
//     return ZendWeakref{
//         referent:referent,
//         std:std,
//     }
// }
func (this *ZendWeakref) GetReferent() *ZendObject      { return this.referent }
func (this *ZendWeakref) SetReferent(value *ZendObject) { this.referent = value }
func (this *ZendWeakref) GetStd() ZendObject            { return this.std }

// func (this *ZendWeakref) SetStd(value ZendObject) { this.std = value }
