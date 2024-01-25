package def

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
)

type FuncType = types.FunctionEntry
type FuncHandler = php.ZifHandler
type ArgInfo struct {
	Name     string
	ByRef    bool
	Variadic bool
}
type ReturnInfo = types.ArgInfo

func DefFunc(name string, minNumArgs uint32, maxNumArgs int32, argInfos []ArgInfo, handler FuncHandler) FuncType {
	var realArgInfos []types.ArgInfo
	if len(argInfos) > 0 {
		realArgInfos = make([]types.ArgInfo, len(argInfos))
		for i, argInfo := range argInfos {
			realArgInfos[i] = types.ArgInfo{
				Name:     argInfo.Name,
				ByRef:    argInfo.ByRef,
				Variadic: argInfo.Variadic,
			}
		}
	}

	return types.DefFunctionEntry(name, handler, realArgInfos)
}
