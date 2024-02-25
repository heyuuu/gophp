package streams

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
)

const PHP_STREAM_NOTIFIER_PROGRESS = 1

/* Attempt to fetch context from the zval passed,
   If no context was passed, use the default context
   The default context has not yet been created, do it now. */

/* not all notification codes are implemented */

const (
	StreamNotifyResolve      = 1
	StreamNotifyConnect      = 2
	StreamNotifyAuthRequired = 3
	StreamNotifyMimeTypeIs   = 4
	StreamNotifyFileSizeIs   = 5
	StreamNotifyRedirected   = 6
	StreamNotifyProgress     = 7
	StreamNotifyCompleted    = 8
	StreamNotifyFailure      = 9
	StreamNotifyAuthResult   = 10

	StreamNotifySeverityInfo = 0
	StreamNotifySeverityWarn = 1
	StreamNotifySeverityErr  = 2
)

// -- api

func StreamNotifyInfo(context *StreamContext, code int, xmsg string, xcode int) {
	if context != nil && context.Notifier() != nil {
		context.Notifier().Notify(context, code, StreamNotifySeverityInfo, xmsg, xcode, 0, 0, nil)
	}
}

func StreamNotifyProgressInit(context *StreamContext, sofar int, bmax int) {
	if context != nil && context.Notifier() != nil {
		context.Notifier().Init(context, sofar, bmax)
	}
}

func PhpStreamNotifyProgressIncrement(context *StreamContext, dsofar int, dmax int) {
	if context != nil && context.Notifier() != nil {
		context.Notifier().Increment(context, dsofar, dmax)
	}
}
func PhpStreamNotifyFileSize(context *StreamContext, fileSize int, xmsg string, xcode int) {
	if context != nil && context.Notifier() != nil {
		context.Notifier().Notify(context, StreamNotifyFileSizeIs, StreamNotifySeverityInfo, xmsg, xcode, 0, fileSize, nil)
	}
}
func PhpStreamNotifyError(context *StreamContext, code int, xmsg string, xcode int) {
	if context != nil && context.Notifier() != nil {
		context.Notifier().Notify(context, code, StreamNotifySeverityErr, xmsg, xcode, 0, 0, nil)
	}
}

// -- types

// StreamContext
type StreamContext struct {
	notifier *StreamNotifier `prop:""`
	resource *types.Resource `get:""`
	options  *types.Table[*types.Table[*types.Zval]]
}

func NewStreamContext(ctx *php.Context) *StreamContext {
	context := &StreamContext{}
	context.resource = php.RegisterResource(ctx, context, LeStreamContext)
	context.options = types.NewTable[*types.Table[*types.Zval]]()
	return context
}
func (c *StreamContext) GetOption(wrapperName string, optionName string) *types.Zval {
	if c == nil {
		return nil
	}
	wrapperOptions := c.options.Get(wrapperName)
	if wrapperOptions == nil {
		return nil
	}
	return wrapperOptions.Get(optionName)
}

func (c *StreamContext) SetOption(wrapperName string, optionName string, optionValue types.Zval) {
	wrapperOptions := c.options.Get(wrapperName)
	if wrapperOptions == nil {
		wrapperOptions = types.NewTable[*types.Zval]()
		c.options.Set(wrapperName, wrapperOptions)
	}
	wrapperOptions.Set(optionName, &optionValue)
}

func (c *StreamContext) OptionsArray() *types.Array {
	// todo 优化成可复用形式？
	arr := types.NewArrayCap(c.options.Len())
	c.options.Each(func(wrapperName string, wrapperOptions *types.Table[*types.Zval]) {
		wrapperArr := types.NewArrayCap(wrapperOptions.Len())
		wrapperOptions.Each(func(optionName string, optionValue *types.Zval) {
			if optionValue != nil {
				wrapperArr.KeyAdd(optionName, *optionValue)
			}
		})

		arr.KeyAdd(wrapperName, types.ZvalArray(wrapperArr))
	})
	return arr
}

type StreamNotificationFunc func(context *StreamContext, notifycode int, severity int, xmsg string, xcode int, bytesSofar int, bytesMax int, ptr any)

// StreamNotifier
type StreamNotifier struct {
	fn          StreamNotificationFunc
	ptr         types.Zval
	mask        int
	progress    int
	progressMax int
}

func NewPhpStreamNotifier(fn StreamNotificationFunc, ptr *types.Zval) *StreamNotifier {
	notifier := &StreamNotifier{fn: fn}
	//notifier.ptr.CopyValueFrom(ptr)
	return notifier
}

// PhpStreamNotificationNotify
func (n *StreamNotifier) Notify(context *StreamContext, notifycode int, severity int, xmsg string, xcode int, bytesSofar int, bytesMax int, ptr any) {
	n.fn(context, notifycode, severity, xmsg, xcode, bytesSofar, bytesMax, ptr)
}

// StreamNotifyProgressInit
func (n *StreamNotifier) Init(context *StreamContext, sofar int, bmax int) {
	n.progress = sofar
	n.progressMax = bmax
	n.mask |= PHP_STREAM_NOTIFIER_PROGRESS

	n.doProgress(context)
}

// PhpStreamNotifyProgressIncrement
func (n *StreamNotifier) Increment(context *StreamContext, dsofar int, dmax int) {
	if n.mask&PHP_STREAM_NOTIFIER_PROGRESS != 0 {
		n.progress += dsofar
		n.progressMax += dmax

		n.doProgress(context)
	}
}

// PhpStreamNotifyProgress
func (n *StreamNotifier) doProgress(context *StreamContext) {
	bsofar := n.progress
	bmax := n.progressMax
	n.Notify(context, StreamNotifyProgress, StreamNotifySeverityInfo, "", 0, bsofar, bmax, nil)
}

func (n *StreamNotifier) GetFunc() StreamNotificationFunc { return n.fn }
func (n *StreamNotifier) GetPtr() *types.Zval             { return &n.ptr }
