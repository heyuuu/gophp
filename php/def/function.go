package def

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
)

type FuncType = types.FunctionEntry
type FuncHandler = php.ZifHandler
type ArgInfo struct {
	Name string
}
type ReturnInfo = types.ArgInfo

func DefFunc(name string, minNumArgs uint32, maxNumArgs int32, argInfos []ArgInfo, handler FuncHandler) FuncType {

	return types.DefFunctionEntry(name, handler)
}
