package context

import "github.com/heyuuu/gophp/php/types"

type Context interface {
	RegisterConst(name string, value *types.Zval)
	RegisterClass(name string, class types.Class)
	RegisterFunction(name string, fun types.Function)

	Arg(idx int) any
	ArgOrDefault(idx int, defaultValue *types.Zval) any
}
