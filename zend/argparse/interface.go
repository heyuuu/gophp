package argparse

import "sik/zend/types"

type ExecuteData interface {
	CalleeName() string
	NumArgs() int
	Arg(pos int) *types.Zval
	IsArgUseStrictTypes() bool
}
