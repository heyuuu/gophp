package standard

import (
	b "sik/builtin"
)

var COMMON = b.Cond(is_ref, "&", "")
