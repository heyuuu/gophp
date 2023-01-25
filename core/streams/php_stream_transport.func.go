// <<generate>>

package streams

import (
	"sik/core"
	"sik/zend"
)

func PhpStreamXportCreate(name *byte, namelen int, options int, flags int, persistent_id *byte, timeout *__struct__timeval, context *core.PhpStreamContext, estr **zend.ZendString, ecode *int) *core.PhpStream {
	return _phpStreamXportCreate(name, namelen, options, flags, persistent_id, timeout, context, estr, ecode)
}
