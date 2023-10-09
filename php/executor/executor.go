package executor

import "github.com/heyuuu/gophp/php/types"

// Executor
type Executor interface {
	RegisterClass(name string, typ any)
	RegisterClassConst(name string, value any)
	RegisterClassProperty(name string, value any)

	Arg(idx int) any
	ArgOrDefault(idx int, defaultValue *types.Zval) any
}
