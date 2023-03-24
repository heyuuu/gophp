package core

/**
 * Opt
 */
type Opt struct {
	Char      byte
	NeedParam bool
	Name      string
}

func MakeOpt(opt_char byte, need_param int, opt_name string) Opt {
	return Opt{
		Char:      opt_char,
		NeedParam: need_param != 0,
		Name:      opt_name,
	}
}
func (this *Opt) GetOptChar() byte       { return this.Char }
func (this *Opt) SetOptChar(value byte)  { this.Char = value }
func (this *Opt) GetNeedParam() int      { return this.NeedParam }
func (this *Opt) SetNeedParam(value int) { this.NeedParam = value }
func (this *Opt) GetOptName() *byte      { return this.Name }
func (this *Opt) SetOptName(value *byte) { this.Name = value }
