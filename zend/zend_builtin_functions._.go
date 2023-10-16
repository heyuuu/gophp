package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

var BuiltinFunctions = []types.FunctionEntry{
	DefZifZendVersion,
	DefZifFuncNumArgs,
	DefZifFuncGetArg,
	DefZifFuncGetArgs,
	DefZifStrlen,
	DefZifStrcmp,
	DefZifStrncmp,
	DefZifStrcasecmp,
	DefZifStrncasecmp,
	DefZifEach,
	DefZifErrorReporting,
	DefZifDefine,
	DefZifDefined,
	DefZifGetClass,
	DefZifGetCalledClass,
	DefZifGetParentClass,
	DefZifMethodExists,
	DefZifPropertyExists,
	DefZifClassExists,
	DefZifInterfaceExists,
	DefZifTraitExists,
	DefZifFunctionExists,
	DefZifClassAlias,
	DefZifGetIncludedFiles,
	DefZifGetRequiredFiles,
	DefZifIsSubclassOf,
	DefZifIsA,
	DefZifGetClassVars,
	DefZifGetObjectVars,
	DefZifGetMangledObjectVars,
	DefZifGetClassMethods,
	DefZifTriggerError,
	DefZifUserError,
	DefZifSetErrorHandler,
	DefZifRestoreErrorHandler,
	DefZifSetExceptionHandler,
	DefZifRestoreExceptionHandler,
	DefZifGetDeclaredClasses,
	DefZifGetDeclaredTraits,
	DefZifGetDeclaredInterfaces,
	DefZifGetDefinedFunctions,
	DefZifGetDefinedVars,
	DefZifGetResourceType,
	DefZifGetResources,
	DefZifGetLoadedExtensions,
	DefZifExtensionLoaded,
	DefZifGetExtensionFuncs,
	DefZifGetDefinedConstants,
	DefZifDebugBacktrace,
	DefZifDebugPrintBacktrace,
	DefZifGcMemCaches,
	DefZifGcCollectCycles,
	DefZifGcEnabled,
	DefZifGcEnable,
	DefZifGcDisable,
	DefZifGcStatus,
}

// BasicModuleData
type BasicModuleData struct{}

var _ ModuleData = (*BasicModuleData)(nil)

func (d *BasicModuleData) Name() string                     { return "Core" }
func (d *BasicModuleData) Version() string                  { return ZEND_VERSION }
func (d *BasicModuleData) Functions() []types.FunctionEntry { return BuiltinFunctions }
func (d *BasicModuleData) ModuleStartup(moduleNumber int) bool {
	return ZmStartupCore(0, moduleNumber) == types.SUCCESS
}
func (d *BasicModuleData) ModuleShutdown(moduleNumber int) bool {
	return true
}
func (d *BasicModuleData) RequestStartup(moduleNumber int) bool {
	return true
}
func (d *BasicModuleData) RequestShutdown(moduleNumber int) bool {
	return true
}

var ZendBuiltinModule = MakeZendModuleEntry(&BasicModuleData{}, nil)
