// <<generate>>

package zend

/**
 * SmartString
 */
type SmartString struct {
	c    *byte
	len_ int
	a    int
}

func (this *SmartString) GetC() *byte      { return this.c }
func (this *SmartString) SetC(value *byte) { this.c = value }
func (this *SmartString) GetLen() int      { return this.len_ }
func (this *SmartString) SetLen(value int) { this.len_ = value }
func (this *SmartString) GetA() int        { return this.a }
func (this *SmartString) SetA(value int)   { this.a = value }
