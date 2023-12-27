package types

import "github.com/heyuuu/gophp/php/perr"

func assert(cond bool) {
	perr.Assert(cond)
}
