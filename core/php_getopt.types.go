// <<generate>>

package core

import b "sik/builtin"

/**
 * Opt
 */
type Opt struct {
	opt_char   byte
	need_param int
	opt_name   *byte
}

// func NewOpt(opt_char byte, need_param int, opt_name *byte) *Opt {
//     return &Opt{
//         opt_char:opt_char,
//         need_param:need_param,
//         opt_name:opt_name,
//     }
// }
func MakeOpt(opt_char byte, need_param int, opt_name string) Opt {
	return Opt{
		opt_char:   opt_char,
		need_param: need_param,
		opt_name:   b.CastStrPtr(opt_name),
	}
}
func (this *Opt) GetOptChar() byte       { return this.opt_char }
func (this *Opt) SetOptChar(value byte)  { this.opt_char = value }
func (this *Opt) GetNeedParam() int      { return this.need_param }
func (this *Opt) SetNeedParam(value int) { this.need_param = value }
func (this *Opt) GetOptName() *byte      { return this.opt_name }
func (this *Opt) SetOptName(value *byte) { this.opt_name = value }
