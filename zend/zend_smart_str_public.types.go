// <<generate>>

package zend

/**
 * SmartStr
 */
type SmartStr struct {
	s *ZendString
	a int
}

func (this SmartStr) GetS() *ZendString       { return this.s }
func (this *SmartStr) SetS(value *ZendString) { this.s = value }
func (this SmartStr) GetA() int               { return this.a }
func (this *SmartStr) SetA(value int)         { this.a = value }
