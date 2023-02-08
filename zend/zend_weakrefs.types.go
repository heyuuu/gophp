// <<generate>>

package zend

/**
 * ZendWeakref
 */
type ZendWeakref struct {
	referent *ZendObject
	std      ZendObject
}

func (this *ZendWeakref) GetReferent() *ZendObject      { return this.referent }
func (this *ZendWeakref) SetReferent(value *ZendObject) { this.referent = value }
func (this *ZendWeakref) GetStd() ZendObject            { return this.std }
