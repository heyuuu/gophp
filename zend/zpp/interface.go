package zpp

import (
	"github.com/heyuuu/gophp/php/types"
)

type ExecuteData interface {
	CalleeName() string
	NumArgs() int
	Arg(pos int) *types.Zval
	IsArgUseStrictTypes() bool
}
