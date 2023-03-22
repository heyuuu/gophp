package def

import (
	"log"
	"sik/zend"
	"sik/zend/types"
)

type FuncType = types.ZendFunctionEntry
type FuncHandler = zend.ZifHandler
type ArgInfo = zend.ArgInfo
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

	return types.DefFunctionEntry(
		name,
		handler,
		minNumArgs,
		argInfos,
		nil,
		0,
	)
}
