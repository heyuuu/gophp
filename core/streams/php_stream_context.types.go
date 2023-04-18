package streams

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

/**
 * PhpStreamNotifier
 */
type PhpStreamNotifier struct {
	func_        PhpStreamNotificationFunc
	dtor         func(notifier *PhpStreamNotifier)
	ptr          types2.Zval
	mask         int
	progress     int
	progress_max int
}

func (this *PhpStreamNotifier) GetFunc() PhpStreamNotificationFunc              { return this.func_ }
func (this *PhpStreamNotifier) SetFunc(value PhpStreamNotificationFunc)         { this.func_ = value }
func (this *PhpStreamNotifier) GetDtor() func(notifier *PhpStreamNotifier)      { return this.dtor }
func (this *PhpStreamNotifier) SetDtor(value func(notifier *PhpStreamNotifier)) { this.dtor = value }
func (this *PhpStreamNotifier) GetPtr() types2.Zval                             { return this.ptr }

func (this *PhpStreamNotifier) GetMask() int             { return this.mask }
func (this *PhpStreamNotifier) SetMask(value int)        { this.mask = value }
func (this *PhpStreamNotifier) GetProgress() int         { return this.progress }
func (this *PhpStreamNotifier) SetProgress(value int)    { this.progress = value }
func (this *PhpStreamNotifier) GetProgressMax() int      { return this.progress_max }
func (this *PhpStreamNotifier) SetProgressMax(value int) { this.progress_max = value }

/**
 * PhpStreamContext
 */
type PhpStreamContext struct {
	notifier *PhpStreamNotifier
	options  types2.Zval
	res      *types2.ZendResource
}

func (this *PhpStreamContext) GetNotifier() *PhpStreamNotifier      { return this.notifier }
func (this *PhpStreamContext) SetNotifier(value *PhpStreamNotifier) { this.notifier = value }
func (this *PhpStreamContext) GetOptions() types2.Zval              { return this.options }

func (this *PhpStreamContext) GetRes() *types2.ZendResource      { return this.res }
func (this *PhpStreamContext) SetRes(value *types2.ZendResource) { this.res = value }
