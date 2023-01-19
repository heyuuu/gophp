// <<generate>>

package core

/**
 * Opt
 */
type Opt struct {
	opt_char   byte
	need_param int
	opt_name   *byte
}

func (this Opt) GetOptChar() byte        { return this.opt_char }
func (this *Opt) SetOptChar(value byte)  { this.opt_char = value }
func (this Opt) GetNeedParam() int       { return this.need_param }
func (this *Opt) SetNeedParam(value int) { this.need_param = value }
func (this Opt) GetOptName() *byte       { return this.opt_name }
func (this *Opt) SetOptName(value *byte) { this.opt_name = value }
