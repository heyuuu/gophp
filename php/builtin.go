package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

var builtinModuleEntries []ModuleEntry

func AddBuiltinModule(m ModuleEntry) {
	builtinModuleEntries = append(builtinModuleEntries, m)
}

func ZendRegisterStandardConstants(ctx *Context) {
	RegisterConstant(ctx, 0, "E_ERROR", Long(int(perr.E_ERROR)))
	RegisterConstant(ctx, 0, "E_RECOVERABLE_ERROR", Long(int(perr.E_RECOVERABLE_ERROR)))
	RegisterConstant(ctx, 0, "E_WARNING", Long(int(perr.E_WARNING)))
	RegisterConstant(ctx, 0, "E_PARSE", Long(int(perr.E_PARSE)))
	RegisterConstant(ctx, 0, "E_NOTICE", Long(int(perr.E_NOTICE)))
	RegisterConstant(ctx, 0, "E_STRICT", Long(int(perr.E_STRICT)))
	RegisterConstant(ctx, 0, "E_DEPRECATED", Long(int(perr.E_DEPRECATED)))
	RegisterConstant(ctx, 0, "E_CORE_ERROR", Long(int(perr.E_CORE_ERROR)))
	RegisterConstant(ctx, 0, "E_CORE_WARNING", Long(int(perr.E_CORE_WARNING)))
	RegisterConstant(ctx, 0, "E_COMPILE_ERROR", Long(int(perr.E_COMPILE_ERROR)))
	RegisterConstant(ctx, 0, "E_COMPILE_WARNING", Long(int(perr.E_COMPILE_WARNING)))
	RegisterConstant(ctx, 0, "E_USER_ERROR", Long(int(perr.E_USER_ERROR)))
	RegisterConstant(ctx, 0, "E_USER_WARNING", Long(int(perr.E_USER_WARNING)))
	RegisterConstant(ctx, 0, "E_USER_NOTICE", Long(int(perr.E_USER_NOTICE)))
	RegisterConstant(ctx, 0, "E_USER_DEPRECATED", Long(int(perr.E_USER_DEPRECATED)))
	RegisterConstant(ctx, 0, "E_ALL", Long(int(perr.E_ALL)))
	//RegisterConstant(ctx, 0, "DEBUG_BACKTRACE_PROVIDE_OBJECT", Long(DEBUG_BACKTRACE_PROVIDE_OBJECT))
	//RegisterConstant(ctx, 0, "DEBUG_BACKTRACE_IGNORE_ARGS", Long(DEBUG_BACKTRACE_IGNORE_ARGS))

	/* true/false constants */
	RegisterConstant(ctx, 0, "ZEND_THREAD_SAFE", types.False)
	RegisterConstant(ctx, 0, "ZEND_DEBUG_BUILD", types.False)

	RegisterConstantEx(ctx, 0, "TRUE", types.True, types.ConstPersistent|types.ConstCtSubst)
	RegisterConstantEx(ctx, 0, "FALSE", types.False, types.ConstPersistent|types.ConstCtSubst)
	RegisterConstantEx(ctx, 0, "NULL", types.Null, types.ConstPersistent|types.ConstCtSubst)
}
