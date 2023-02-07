// <<generate>>

package streams

import (
	"sik/core"
	"sik/zend"
)

/**
 * PhpStreamXportParam
 */
type PhpStreamXportParam struct {
	op             int
	want_addr      uint
	want_textaddr  uint
	want_errortext uint
	how            uint
	inputs         struct {
		name    *byte
		namelen int
		timeout *__struct__timeval
		addr    *__struct__sockaddr
		buf     *byte
		buflen  int
		addrlen socklen_t
		backlog int
		flags   int
	}
	outputs struct {
		client     *core.PhpStream
		addr       *__struct__sockaddr
		addrlen    socklen_t
		textaddr   *zend.ZendString
		error_text *zend.ZendString
		returncode int
		error_code int
	}
}

func (this *PhpStreamXportParam) GetOp() int                              { return this.op }
func (this *PhpStreamXportParam) SetOp(value int)                         { this.op = value }
func (this *PhpStreamXportParam) GetWantAddr() uint                       { return this.want_addr }
func (this *PhpStreamXportParam) SetWantAddr(value uint)                  { this.want_addr = value }
func (this *PhpStreamXportParam) GetWantTextaddr() uint                   { return this.want_textaddr }
func (this *PhpStreamXportParam) SetWantTextaddr(value uint)              { this.want_textaddr = value }
func (this *PhpStreamXportParam) GetWantErrortext() uint                  { return this.want_errortext }
func (this *PhpStreamXportParam) SetWantErrortext(value uint)             { this.want_errortext = value }
func (this *PhpStreamXportParam) GetHow() uint                            { return this.how }
func (this *PhpStreamXportParam) SetHow(value uint)                       { this.how = value }
func (this *PhpStreamXportParam) GetName() *byte                          { return this.inputs.name }
func (this *PhpStreamXportParam) SetName(value *byte)                     { this.inputs.name = value }
func (this *PhpStreamXportParam) GetNamelen() int                         { return this.inputs.namelen }
func (this *PhpStreamXportParam) SetNamelen(value int)                    { this.inputs.namelen = value }
func (this *PhpStreamXportParam) GetTimeout() *__struct__timeval          { return this.inputs.timeout }
func (this *PhpStreamXportParam) SetTimeout(value *__struct__timeval)     { this.inputs.timeout = value }
func (this *PhpStreamXportParam) GetInputsAddr() *__struct__sockaddr      { return this.inputs.addr }
func (this *PhpStreamXportParam) SetInputsAddr(value *__struct__sockaddr) { this.inputs.addr = value }
func (this *PhpStreamXportParam) GetBuf() *byte                           { return this.inputs.buf }
func (this *PhpStreamXportParam) SetBuf(value *byte)                      { this.inputs.buf = value }
func (this *PhpStreamXportParam) GetBuflen() int                          { return this.inputs.buflen }
func (this *PhpStreamXportParam) SetBuflen(value int)                     { this.inputs.buflen = value }
func (this *PhpStreamXportParam) GetInputsAddrlen() socklen_t             { return this.inputs.addrlen }
func (this *PhpStreamXportParam) SetInputsAddrlen(value socklen_t)        { this.inputs.addrlen = value }
func (this *PhpStreamXportParam) GetBacklog() int                         { return this.inputs.backlog }
func (this *PhpStreamXportParam) SetBacklog(value int)                    { this.inputs.backlog = value }
func (this *PhpStreamXportParam) GetFlags() int                           { return this.inputs.flags }
func (this *PhpStreamXportParam) SetFlags(value int)                      { this.inputs.flags = value }
func (this *PhpStreamXportParam) GetClient() *core.PhpStream              { return this.outputs.client }
func (this *PhpStreamXportParam) SetClient(value *core.PhpStream)         { this.outputs.client = value }
func (this *PhpStreamXportParam) GetOutputsAddr() *__struct__sockaddr     { return this.outputs.addr }

// func (this *PhpStreamXportParam) SetOutputsAddr(value *__struct__sockaddr) { this.outputs.addr = value }
func (this *PhpStreamXportParam) GetOutputsAddrlen() socklen_t { return this.outputs.addrlen }

// func (this *PhpStreamXportParam) SetOutputsAddrlen(value socklen_t) { this.outputs.addrlen = value }
func (this *PhpStreamXportParam) GetTextaddr() *zend.ZendString { return this.outputs.textaddr }

// func (this *PhpStreamXportParam) SetTextaddr(value *zend.ZendString) { this.outputs.textaddr = value }
func (this *PhpStreamXportParam) GetErrorText() *zend.ZendString { return this.outputs.error_text }
func (this *PhpStreamXportParam) SetErrorText(value *zend.ZendString) {
	this.outputs.error_text = value
}
func (this *PhpStreamXportParam) GetReturncode() int      { return this.outputs.returncode }
func (this *PhpStreamXportParam) SetReturncode(value int) { this.outputs.returncode = value }
func (this *PhpStreamXportParam) GetErrorCode() int       { return this.outputs.error_code }
func (this *PhpStreamXportParam) SetErrorCode(value int)  { this.outputs.error_code = value }

/* PhpStreamXportParam.inputs.flags */
func (this *PhpStreamXportParam) AddFlags(value int)      { this.inputs.flags |= value }
func (this *PhpStreamXportParam) SubFlags(value int)      { this.inputs.flags &^= value }
func (this *PhpStreamXportParam) HasFlags(value int) bool { return this.inputs.flags&value != 0 }
func (this *PhpStreamXportParam) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}

/**
 * PhpStreamXportCryptoParam
 */
type PhpStreamXportCryptoParam struct {
	inputs struct {
		session  *core.PhpStream
		activate int
		method   PhpStreamXportCryptMethodT
	}
	outputs struct {
		returncode int
	}
	op int
}

// func (this *PhpStreamXportCryptoParam)  GetSession() *core.PhpStream      { return this.inputs.session }
func (this *PhpStreamXportCryptoParam) SetSession(value *core.PhpStream) { this.inputs.session = value }

// func (this *PhpStreamXportCryptoParam)  GetActivate() int      { return this.inputs.activate }
func (this *PhpStreamXportCryptoParam) SetActivate(value int) { this.inputs.activate = value }

// func (this *PhpStreamXportCryptoParam)  GetMethod() PhpStreamXportCryptMethodT      { return this.inputs.method }
func (this *PhpStreamXportCryptoParam) SetMethod(value PhpStreamXportCryptMethodT) {
	this.inputs.method = value
}
func (this *PhpStreamXportCryptoParam) GetReturncode() int { return this.outputs.returncode }

// func (this *PhpStreamXportCryptoParam) SetReturncode(value int) { this.outputs.returncode = value }
// func (this *PhpStreamXportCryptoParam)  GetOp() int      { return this.op }
func (this *PhpStreamXportCryptoParam) SetOp(value int) { this.op = value }
