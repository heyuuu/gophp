package streams

import (
	"sik/core"
)

type PhpStreamNotificationFunc func(
	context *core.PhpStreamContext,
	notifycode int,
	severity int,
	xmsg *byte,
	xcode int,
	bytes_sofar int,
	bytes_max int,
	ptr any,
)

const PHP_STREAM_NOTIFIER_PROGRESS = 1

/* Attempt to fetch context from the zval passed,
   If no context was passed, use the default context
   The default context has not yet been created, do it now. */

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
