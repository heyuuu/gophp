// <<generate>>

package zend

/**
 * ZendString
 */
type ZendString struct {
	baseZendRefcounted
	h    ZendUlong
	len_ int
	val  []byte
}

var _ ZendRefcounted = &ZendString{}

func (this ZendString) GetH() ZendUlong       { return this.h }
func (this *ZendString) SetH(value ZendUlong) { this.h = value }
func (this ZendString) GetLen() int           { return this.len_ }
func (this *ZendString) SetLen(value int)     { this.len_ = value }
func (this ZendString) GetVal() []byte        { return this.val }
func (this *ZendString) SetVal(value []byte)  { this.val = value }
