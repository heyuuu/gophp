// <<generate>>

package streams

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	"sik/zend"
)

// Source: <main/streams/php_stream_context.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Wez Furlong <wez@thebrainroom.com>                           |
   +----------------------------------------------------------------------+
*/

type PhpStreamNotificationFunc func(context *core.PhpStreamContext, notifycode int, severity int, xmsg *byte, xcode int, bytes_sofar int, bytes_max int, ptr any)

const PHP_STREAM_NOTIFIER_PROGRESS = 1

/* Attempt to fetch context from the zval passed,
   If no context was passed, use the default context
   The default context has not yet been created, do it now. */

func PhpStreamContextFromZval(zcontext *zend.Zval, nocontext int) __auto__ {
	if b.CondF2(b.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", standard.PhpLeStreamContext()) }, nocontext), nil, func() __auto__ { return standard.FG(default_context) }) {
		return standard.FG(default_context)
	} else {
		standard.FG(default_context) = PhpStreamContextAlloc()
		return standard.FG(default_context)
	}
}
func PhpStreamContextToZval(context *core.PhpStreamContext, zval *zend.Zval) {
	zend.ZVAL_RES(zval, context.GetRes())
	zend.GC_ADDREF(context.GetRes())
}

/* not all notification codes are implemented */

const PHP_STREAM_NOTIFY_RESOLVE = 1
const PHP_STREAM_NOTIFY_CONNECT = 2
const PHP_STREAM_NOTIFY_AUTH_REQUIRED = 3
const PHP_STREAM_NOTIFY_MIME_TYPE_IS = 4
const PHP_STREAM_NOTIFY_FILE_SIZE_IS = 5
const PHP_STREAM_NOTIFY_REDIRECTED = 6
const PHP_STREAM_NOTIFY_PROGRESS = 7
const PHP_STREAM_NOTIFY_COMPLETED = 8
const PHP_STREAM_NOTIFY_FAILURE = 9
const PHP_STREAM_NOTIFY_AUTH_RESULT = 10
const PHP_STREAM_NOTIFY_SEVERITY_INFO = 0
const PHP_STREAM_NOTIFY_SEVERITY_WARN = 1
const PHP_STREAM_NOTIFY_SEVERITY_ERR = 2

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
