package types

import (
	b "sik/builtin"
	"sik/zend"
)

func (this *Bucket) GetH() uint {
	if this.IsStrKey() {
		return b.HashStr(this.key.KeyKey())
	} else {
		return uint(this.key.index)
	}
}
func (this *Bucket) GetKey() *String {
	if this.IsStrKey() {
		return NewString(this.key.KeyKey())
	} else {
		return nil
	}
}
func (this *Bucket) SetH(value zend.ZendUlong) {
	// todo remove
	b.Assert(false)
}
func (this *Bucket) SetKey(value *String) {
	// todo remove
	b.Assert(false)
}

func (this *Bucket) CopyFrom(from *Bucket) {
	this.SetVal(from.GetVal())
	this.key = from.key
}

func (this *Bucket) IsValid() bool {
	return !this.val.IsUndef()
}

func (this *Bucket) SetInvalid() {
	this.val.SetUndef()
}


