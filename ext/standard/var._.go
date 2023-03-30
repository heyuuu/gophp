package standard

import (
	b "github.com/heyuuu/gophp/builtin"
)

var COMMON = b.Cond(is_ref, "&", "")
