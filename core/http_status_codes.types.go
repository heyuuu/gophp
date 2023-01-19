// <<generate>>

package core

/**
 * HttpResponseStatusCodePair
 */
type HttpResponseStatusCodePair struct {
	code int
	str  *byte
}

func (this HttpResponseStatusCodePair) GetCode() int        { return this.code }
func (this *HttpResponseStatusCodePair) SetCode(value int)  { this.code = value }
func (this HttpResponseStatusCodePair) GetStr() *byte       { return this.str }
func (this *HttpResponseStatusCodePair) SetStr(value *byte) { this.str = value }
