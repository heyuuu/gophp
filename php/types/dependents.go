package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
)

func assert(cond bool) {
	b.Assert(cond)
}

func triggerError(message string) {
	faults.ErrorNoreturn(faults.E_ERROR, message)
}
