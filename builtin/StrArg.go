package builtin

type StrArg struct {
	str  *byte
	len_ uint
}

func NewStrArg(str *byte, len_ uint) StrArg { return StrArg{str, len_} }

func (this StrArg) Len() uint     { return this.len_ }
func (this StrArg) StrPtr() *byte { return this.str }
func (this StrArg) Str() string   { return CastStr(this.str, this.len_) }

func (this StrArg) Hash() uint {
	return HashStr(this.Str())
}
