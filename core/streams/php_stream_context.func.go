package streams

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

func PhpStreamContextFromZval(zcontext *types.Zval, nocontext int) __auto__ {
	if b.CondF2(b.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", standard.PhpLeStreamContext()) }, nocontext), nil, func() __auto__ { return standard.FG__().default_context }) {
		return standard.FG__().default_context
	} else {
		standard.FG__().default_context = PhpStreamContextAlloc()
		return standard.FG__().default_context
	}
}
func PhpStreamContextToZval(context *core.PhpStreamContext, zval *types.Zval) {
	zval.SetResource(context.GetRes())
	//context.GetRes().AddRefcount()
}
func PhpStreamNotifyInfo(context *core.PhpStreamContext, code zend.ZendLong, xmsg *byte, xcode int) {
	if context != nil && context.GetNotifier() != nil {
		PhpStreamNotificationNotify(context, code, PHP_STREAM_NOTIFY_SEVERITY_INFO, xmsg, xcode, 0, 0, nil)
	}
}
func PhpStreamNotifyProgress(context *core.PhpStreamContext, bsofar int, bmax int) {
	if context != nil && context.GetNotifier() != nil {
		PhpStreamNotificationNotify(context, PHP_STREAM_NOTIFY_PROGRESS, PHP_STREAM_NOTIFY_SEVERITY_INFO, nil, 0, bsofar, bmax, nil)
	}
}
func PhpStreamNotifyProgressInit(context *core.PhpStreamContext, sofar int, bmax int) {
	if context != nil && context.GetNotifier() != nil {
		context.GetNotifier().SetProgress(sofar)
		context.GetNotifier().SetProgressMax(bmax)
		context.GetNotifier().SetMask(context.GetNotifier().GetMask() | PHP_STREAM_NOTIFIER_PROGRESS)
		PhpStreamNotifyProgress(context, sofar, bmax)
	}
}
func PhpStreamNotifyProgressIncrement(context *core.PhpStreamContext, dsofar ssize_t, dmax int) {
	if context != nil && context.GetNotifier() != nil && (context.GetNotifier().GetMask()&PHP_STREAM_NOTIFIER_PROGRESS) != 0 {
		context.GetNotifier().SetProgress(context.GetNotifier().GetProgress() + dsofar)
		context.GetNotifier().SetProgressMax(context.GetNotifier().GetProgressMax() + dmax)
		PhpStreamNotifyProgress(context, context.GetNotifier().GetProgress(), context.GetNotifier().GetProgressMax())
	}
}
func PhpStreamNotifyFileSize(context *core.PhpStreamContext, file_size int, xmsg *byte, xcode int) {
	if context != nil && context.GetNotifier() != nil {
		PhpStreamNotificationNotify(context, PHP_STREAM_NOTIFY_FILE_SIZE_IS, PHP_STREAM_NOTIFY_SEVERITY_INFO, xmsg, xcode, 0, file_size, nil)
	}
}
func PhpStreamNotifyError(context *core.PhpStreamContext, code zend.ZendLong, xmsg *byte, xcode int) {
	if context != nil && context.GetNotifier() != nil {
		PhpStreamNotificationNotify(context, code, PHP_STREAM_NOTIFY_SEVERITY_ERR, xmsg, xcode, 0, 0, nil)
	}
}
