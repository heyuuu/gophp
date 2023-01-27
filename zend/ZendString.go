// <<generate>>

package zend

import b "sik/builtin"

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

func ZendStringNew(str string, persistent bool) *ZendString {
	var zs = &ZendString{len_: len(str), val: []byte(str)}

	zs.SetGcRefcount(1)
	zs.SetGcTypeInfo(IS_STRING)
	if persistent {
		zs.AddGcFlags(IS_STR_PERSISTENT)
	}

	return zs
}

func ZendStringInit(str *byte, len_ int, persistent int) *ZendString {
	var str_ = b.NewStrArg(str, uint(len_)).Str()
	return ZendStringNew(str_, persistent != 0)
}

func ZendStringAlloc(len_ int, persistent int) *ZendString {
	var str_ = string(make([]byte, len_))
	return ZendStringNew(str_, persistent != 0)
}

func (this *ZendString) GetH() ZendUlong      { return this.h }
func (this *ZendString) SetH(value ZendUlong) { this.h = value }
func (this *ZendString) GetLen() int          { return this.len_ }
func (this *ZendString) SetLen(value int)     { this.len_ = value }
func (this *ZendString) GetVal() []byte       { return this.val }
func (this *ZendString) SetVal(value []byte)  { this.val = value }

func (this *ZendString) GetStr() string { return string(this.val[:this.len_]) }
func (this *ZendString) GetHash() ZendUlong {
	if this.h == 0 {
		this.h = b.HashBytes(this.val[:this.len_])
	}

	return this.h
}

func (this *ZendString) IsPersistent() bool {
	return b.FlagMatch(this.GetGcFlags(), IS_STR_PERSISTENT)
}

func (this *ZendString) Copy() *ZendString {
	return ZendStringNew(this.GetStr(), this.IsPersistent())
}

func (this *ZendString) Dup(persistent int) *ZendString {
	return ZendStringNew(this.GetStr(), persistent != 0)
}
