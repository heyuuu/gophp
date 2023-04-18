package streams

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
)

func PhpStreamXportCreate(
	name *byte,
	namelen int,
	options int,
	flags int,
	persistent_id *byte,
	timeout *__struct__timeval,
	context *core.PhpStreamContext,
	estr **types.String,
	ecode *int,
) *core.PhpStream {
	return _phpStreamXportCreate(name, namelen, options, flags, persistent_id, timeout, context, estr, ecode)
}
