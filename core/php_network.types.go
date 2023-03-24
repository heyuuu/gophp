package core

/**
 * PhpNetstreamDataT
 */
type PhpNetstreamDataT struct {
	socket        PhpSocketT
	is_blocked    byte
	timeout       __struct__timeval
	timeout_event byte
	ownsize       int
}

// func MakePhpNetstreamDataT(socket PhpSocketT, is_blocked byte, timeout __struct__timeval, timeout_event byte, ownsize int) PhpNetstreamDataT {
//     return PhpNetstreamDataT{
//         socket:socket,
//         is_blocked:is_blocked,
//         timeout:timeout,
//         timeout_event:timeout_event,
//         ownsize:ownsize,
//     }
// }
func (this *PhpNetstreamDataT) GetSocket() PhpSocketT              { return this.socket }
func (this *PhpNetstreamDataT) SetSocket(value PhpSocketT)         { this.socket = value }
func (this *PhpNetstreamDataT) GetIsBlocked() byte                 { return this.is_blocked }
func (this *PhpNetstreamDataT) SetIsBlocked(value byte)            { this.is_blocked = value }
func (this *PhpNetstreamDataT) GetTimeout() __struct__timeval      { return this.timeout }
func (this *PhpNetstreamDataT) SetTimeout(value __struct__timeval) { this.timeout = value }
func (this *PhpNetstreamDataT) GetTimeoutEvent() byte              { return this.timeout_event }
func (this *PhpNetstreamDataT) SetTimeoutEvent(value byte)         { this.timeout_event = value }

// func (this *PhpNetstreamDataT)  GetOwnsize() int      { return this.ownsize }
// func (this *PhpNetstreamDataT) SetOwnsize(value int) { this.ownsize = value }
