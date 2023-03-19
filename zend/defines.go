package zend

import (
	"log"
	"sik/zend/types"
)

/**
 * 函数定义相关
 */
type DefFuncType = types.ZendFunctionEntry
type DefFuncOpts struct {
	name       string
	handler    ZifHandler
	minNumArgs int32
	maxNumArgs int32
	argNames   []string
	argInfos   []ArgInfo
	returnInfo *ArgInfo
	flags      uint32
}

func DefFunc(opts DefFuncOpts) DefFuncType {
	if len(opts.name) == 0 {
		log.Fatalf("DefFunc() 参数 name 不可为空")
	} else if opts.handler == nil {
		log.Fatalf("DefFunc() 参数 handler 不可为空: name=" + opts.name)
	}

	var argInfos []ArgInfo
	if opts.argInfos != nil {
		argInfos = opts.argInfos
	} else if opts.argNames != nil {
		for _, argName := range opts.argNames {
			argInfos = append(argInfos, MakeArgInfo(argName))
		}
	}

	var requiredNumArgs uint32
	if opts.minNumArgs < 0 {
		requiredNumArgs = uint32(len(argInfos))
	} else {
		requiredNumArgs = uint32(opts.minNumArgs)
	}

	return DefFuncType{
		funcName:        opts.name,
		handler:         opts.handler,
		requiredNumArgs: requiredNumArgs,
		argInfos:        argInfos,
		returnArgInfo:   opts.returnInfo,
		flags:           opts.flags,
	}
}
