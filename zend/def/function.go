package def

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
	"log"
)

type FuncType = types.ZendFunctionEntry
type FuncHandler = zend.ZifHandler
type ArgInfo struct {
	Name string
}
type ReturnInfo = zend.ArgInfo

/**
 * 函数定义相关
 */
type DefFuncOpts struct {
	Handler    FuncHandler
	ArgInfos   []ArgInfo
	ReturnInfo *ReturnInfo
	Flags      uint32
}

func DefFunc(name string, minNumArgs uint32, maxNumArgs int32, argInfos []ArgInfo, handler FuncHandler) FuncType {
	if len(name) == 0 {
		log.Fatalf("DefFunc() 参数 name 不可为空")
	} else if handler == nil {
		log.Fatalf("DefFunc() 参数 handler 不可为空: name=" + name)
	}

	var realArgs []zend.ArgInfo
	for _, info := range argInfos {
		realArgs = append(realArgs, zend.DefArgInfo(info.Name))
	}

	return types.DefFunctionEntry(
		name,
		handler,
		minNumArgs,
		realArgs,
		nil,
		0,
	)
}
