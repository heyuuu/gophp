// <<generate>>

package zend

import "sik/zend/types"

/**
 * ZendSignalEntryT
 */
type ZendSignalEntryT struct {
	flags   int
	handler any
}

// func MakeZendSignalEntryT(flags int, handler any) ZendSignalEntryT {
//     return ZendSignalEntryT{
//         flags:flags,
//         handler:handler,
//     }
// }
// func (this *ZendSignalEntryT)  GetFlags() int      { return this.flags }
func (this *ZendSignalEntryT) SetFlags(value int)   { this.flags = value }
func (this *ZendSignalEntryT) GetHandler() any      { return this.handler }
func (this *ZendSignalEntryT) SetHandler(value any) { this.handler = value }

/* ZendSignalEntryT.flags */
func (this *ZendSignalEntryT) AddFlags(value int)      { this.flags |= value }
func (this *ZendSignalEntryT) SubFlags(value int)      { this.flags &^= value }
func (this *ZendSignalEntryT) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *ZendSignalEntryT) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this ZendSignalEntryT) IsSiginfo() bool           { return this.HasFlags(SA_SIGINFO) }
func (this ZendSignalEntryT) IsResethand() bool         { return this.HasFlags(SA_RESETHAND) }
func (this *ZendSignalEntryT) SetIsSiginfo(cond bool)   { this.SwitchFlags(SA_SIGINFO, cond) }
func (this *ZendSignalEntryT) SetIsResethand(cond bool) { this.SwitchFlags(SA_RESETHAND, cond) }

/**
 * ZendSignalT
 */
type ZendSignalT struct {
	signo   int
	siginfo *siginfo_t
	context any
}

// func MakeZendSignalT(signo int, siginfo *siginfo_t, context any) ZendSignalT {
//     return ZendSignalT{
//         signo:signo,
//         siginfo:siginfo,
//         context:context,
//     }
// }
func (this *ZendSignalT) GetSigno() int               { return this.signo }
func (this *ZendSignalT) SetSigno(value int)          { this.signo = value }
func (this *ZendSignalT) GetSiginfo() *siginfo_t      { return this.siginfo }
func (this *ZendSignalT) SetSiginfo(value *siginfo_t) { this.siginfo = value }
func (this *ZendSignalT) GetContext() any             { return this.context }
func (this *ZendSignalT) SetContext(value any)        { this.context = value }

/**
 * ZendSignalQueueT
 */
type ZendSignalQueueT struct {
	zend_signal ZendSignalT
	next        *ZendSignalQueueT
}

// func MakeZendSignalQueueT(zend_signal ZendSignalT, next *ZendSignalQueueT) ZendSignalQueueT {
//     return ZendSignalQueueT{
//         zend_signal:zend_signal,
//         next:next,
//     }
// }
func (this *ZendSignalQueueT) GetZendSignal() ZendSignalT { return this.zend_signal }

// func (this *ZendSignalQueueT) SetZendSignal(value ZendSignalT) { this.zend_signal = value }
func (this *ZendSignalQueueT) GetNext() *ZendSignalQueueT      { return this.next }
func (this *ZendSignalQueueT) SetNext(value *ZendSignalQueueT) { this.next = value }

/**
 * ZendSignalGlobalsT
 */
type ZendSignalGlobalsT struct {
	depth    int
	blocked  int
	running  int
	active   int
	check    bool
	reset    types.ZendBool
	handlers []ZendSignalEntryT
	pstorage []ZendSignalQueueT
	phead    *ZendSignalQueueT
	ptail    *ZendSignalQueueT
	pavail   *ZendSignalQueueT
}

func (this *ZendSignalGlobalsT) GetCheck() types.ZendBool          { return types.IntBool(this.check) }
func (this *ZendSignalGlobalsT) SetReset(value types.ZendBool)     { this.reset = value }
func (this *ZendSignalGlobalsT) GetPstorage() []ZendSignalQueueT   { return this.pstorage }
func (this *ZendSignalGlobalsT) GetPavail() *ZendSignalQueueT      { return this.pavail }
func (this *ZendSignalGlobalsT) SetPavail(value *ZendSignalQueueT) { this.pavail = value }
