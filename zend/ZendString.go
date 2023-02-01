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

func NewZendString(str string) *ZendString {
	var zs = &ZendString{val: []byte(str), len_: len(str)}

	zs.SetRefcount(1)
	zs.SetGcTypeInfo(IS_STRING)

	return zs
}

func NewZendStringPersistent(str string, persistent bool) *ZendString {
	var zs = NewZendString(str)
	if persistent {
		zs.AddGcFlags(IS_STR_PERSISTENT)
	}
	return zs
}

func (this *ZendString) GetH() ZendUlong      { return this.h }
func (this *ZendString) SetH(value ZendUlong) { this.h = value }
func (this *ZendString) GetLen() int          { return this.len_ }
func (this *ZendString) SetLen(value int)     { this.len_ = value }
func (this *ZendString) GetVal() []byte       { return this.val }
func (this *ZendString) SetVal(value []byte)  { this.val = value }

func (this *ZendString) GetStr() string {
	return string(this.val[:this.len_])
}

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
	return NewZendStringPersistent(this.GetStr(), persistent != 0)
}
