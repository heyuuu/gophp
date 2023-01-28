// <<generate>>

package zend

import b "sik/builtin"

/**
 * ZendString
 */
type ZendString struct {
	ZendRefcounted
	h    ZendUlong
	len_ int
	val  []byte
}

var _ IRefcounted = &ZendString{}

func (this *ZendString) GetH() ZendUlong      { return this.h }
func (this *ZendString) SetH(value ZendUlong) { this.h = value }
func (this *ZendString) GetLen() int          { return this.len_ }
func (this *ZendString) SetLen(value int)     { this.len_ = value }
func (this *ZendString) GetVal() []byte       { return this.val }
func (this *ZendString) SetVal(value []byte)  { this.val = value }

func (this *ZendString) GetHash() ZendUlong {
	if this.h == 0 {
		this.h = b.HashBytes(this.val[:this.len_])
	}

	return this.h
}

func (this *ZendString) Copy() *ZendString {
	this.AddRefcount()
	return this
}

func (this *ZendString) Dup(persistent int) *ZendString {
	return ZendStringInit(this.GetVal(), this.GetLen(), persistent)
}
