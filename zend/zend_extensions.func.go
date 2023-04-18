package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"log"
)

func ZendLoadExtension(path *byte) int {
	var handle any
	handle = DL_LOAD(path)
	if !handle {
		log.Printf("Failed loading %s:  %s\n", path, DL_ERROR())
		return types2.FAILURE
	}
	return ZendLoadExtensionHandle(handle, path)
}
func ZendLoadExtensionHandle(handle any, path *byte) int {
	var new_extension *ZendExtension
	var extension_version_info *ZendExtensionVersionInfo
	extension_version_info = (*ZendExtensionVersionInfo)(DL_FETCH_SYMBOL(handle, "extension_version_info"))
	if extension_version_info == nil {
		extension_version_info = (*ZendExtensionVersionInfo)(DL_FETCH_SYMBOL(handle, "_extension_version_info"))
	}
	new_extension = (*ZendExtension)(DL_FETCH_SYMBOL(handle, "zend_extension_entry"))
	if new_extension == nil {
		new_extension = (*ZendExtension)(DL_FETCH_SYMBOL(handle, "_zend_extension_entry"))
	}
	if extension_version_info == nil || new_extension == nil {
		log.Printf("%s doesn't appear to be a valid Zend extension\n", path)

		/* See http://support.microsoft.com/kb/190351 */

		DL_UNLOAD(handle)
		return types2.FAILURE
	}

	/* allow extension to proclaim compatibility with any Zend version */

	if extension_version_info.GetZendExtensionApiNo() != ZEND_EXTENSION_API_NO && (new_extension.GetApiNoCheck() == nil || new_extension.GetApiNoCheck()(ZEND_EXTENSION_API_NO) != types2.SUCCESS) {
		if extension_version_info.GetZendExtensionApiNo() > ZEND_EXTENSION_API_NO {
			log.Printf("%s requires Zend Engine API version %d.\n"+"The Zend Engine API version %d which is installed, is outdated.\n\n", new_extension.GetName(), extension_version_info.GetZendExtensionApiNo(), ZEND_EXTENSION_API_NO)

			/* See http://support.microsoft.com/kb/190351 */

			DL_UNLOAD(handle)
			return types2.FAILURE
		} else if extension_version_info.GetZendExtensionApiNo() < ZEND_EXTENSION_API_NO {
			log.Printf("%s requires Zend Engine API version %d.\n"+"The Zend Engine API version %d which is installed, is newer.\n"+"Contact %s at %s for a later version of %s.\n\n", new_extension.GetName(), extension_version_info.GetZendExtensionApiNo(), ZEND_EXTENSION_API_NO, new_extension.GetAuthor(), new_extension.GetURL(), new_extension.GetName())

			/* See http://support.microsoft.com/kb/190351 */

			DL_UNLOAD(handle)
			return types2.FAILURE
		}
	} else if strcmp("API"+"ZEND_EXTENSION_API_NO"+ZEND_BUILD_TS, extension_version_info.GetBuildId()) && (new_extension.GetBuildIdCheck() == nil || new_extension.GetBuildIdCheck()("API"+"ZEND_EXTENSION_API_NO"+ZEND_BUILD_TS) != types2.SUCCESS) {
		log.Printf("Cannot load %s - it was built with configuration %s, whereas running engine is %s\n", new_extension.GetName(), extension_version_info.GetBuildId(), "API"+"ZEND_EXTENSION_API_NO"+ZEND_BUILD_TS)

		/* See http://support.microsoft.com/kb/190351 */

		DL_UNLOAD(handle)
		return types2.FAILURE
	} else if ZendGetExtension(new_extension.GetName()) != nil {
		log.Printf("Cannot load %s - it was already loaded\n", new_extension.GetName())

		/* See http://support.microsoft.com/kb/190351 */

		DL_UNLOAD(handle)
		return types2.FAILURE
	}
	return ZendRegisterExtension(new_extension, handle)
}
func ZendRegisterExtension(new_extension *ZendExtension, handle any) int {
	var extension ZendExtension
	extension = *new_extension
	extension.SetHandle(handle)
	ZendExtensionDispatchMessage(ZEND_EXTMSG_NEW_EXTENSION, &extension)
	ZendExtensions.AddElement(&extension)
	if extension.GetOpArrayCtor() != nil {
		ZendExtensionFlags |= ZEND_EXTENSIONS_HAVE_OP_ARRAY_CTOR
	}
	if extension.GetOpArrayDtor() != nil {
		ZendExtensionFlags |= ZEND_EXTENSIONS_HAVE_OP_ARRAY_DTOR
	}
	if extension.GetOpArrayHandler() != nil {
		ZendExtensionFlags |= ZEND_EXTENSIONS_HAVE_OP_ARRAY_HANDLER
	}
	if extension.GetOpArrayPersistCalc() != nil {
		ZendExtensionFlags |= ZEND_EXTENSIONS_HAVE_OP_ARRAY_PERSIST_CALC
	}
	if extension.GetOpArrayPersist() != nil {
		ZendExtensionFlags |= ZEND_EXTENSIONS_HAVE_OP_ARRAY_PERSIST
	}

	/*fprintf(stderr, "Loaded %s, version %s\n", extension.name, extension.version);*/

	return types2.SUCCESS

	/*fprintf(stderr, "Loaded %s, version %s\n", extension.name, extension.version);*/
}
func ZendExtensionShutdown(extension *ZendExtension) {
	if extension.GetShutdown() != nil {
		extension.GetShutdown()(extension)
	}
}
func ZendExtensionStartup(extension *ZendExtension) int {
	if extension.GetStartup() != nil {
		if extension.GetStartup()(extension) != types2.SUCCESS {
			return 1
		}
		ZendAppendVersionInfo(extension)
	}
	return 0
}
func ZendStartupExtensionsMechanism() int {
	/* Startup extensions mechanism */

	ZendExtensions.Init(b.SizeOf("zend_extension"), (func(any))(ZendExtensionDtor), 1)
	ZendOpArrayExtensionHandles = 0
	LastResourceNumber = 0
	return types2.SUCCESS
}
func ZendStartupExtensions() int {
	ZendLlistApplyWithDel(&ZendExtensions, (func(any) int)(ZendExtensionStartup))
	return types2.SUCCESS
}
func ZendShutdownExtensions() {
	ZendExtensions.Apply(LlistApplyFuncT(ZendExtensionShutdown))
	ZendExtensions.Destroy()
}
func ZendExtensionDtor(extension *ZendExtension) {
	if extension.GetHandle() && !(getenv("ZEND_DONT_UNLOAD_MODULES")) {
		DL_UNLOAD(extension.GetHandle())
	}
}
func ZendExtensionMessageDispatcher(extension *ZendExtension, num_args int, args ...any) {
	var message int
	var arg any
	if extension.GetMessageHandler() == nil || num_args != 2 {
		return
	}
	message = __va_arg(args, int(_))
	arg = __va_arg(args, any(_))
	extension.GetMessageHandler()(message, arg)
}
func ZendExtensionDispatchMessage(message int, arg any) {
	ZendLlistApplyWithArguments(&ZendExtensions, LlistApplyWithArgsFuncT(ZendExtensionMessageDispatcher), 2, message, arg)
}
func ZendGetResourceHandle(extension *ZendExtension) int {
	if LastResourceNumber < types2.ZEND_MAX_RESERVED_RESOURCES {
		extension.SetResourceNumber(LastResourceNumber)
		LastResourceNumber++
		return LastResourceNumber - 1
	} else {
		return -1
	}
}
func ZendGetOpArrayExtensionHandle() int {
	ZendOpArrayExtensionHandles++
	return ZendOpArrayExtensionHandles - 1
}
func ZendGetExtension(extension_name *byte) *ZendExtension {
	var element *ZendLlistElement
	for element = ZendExtensions.GetHead(); element != nil; element = element.GetNext() {
		var extension *ZendExtension = (*ZendExtension)(element.GetData())
		if !(strcmp(extension.GetName(), extension_name)) {
			return extension
		}
	}
	return nil
}
func ZendExtensionOpArrayPersistCalcHandler(extension *ZendExtension, data *ZendExtensionPersistData) {
	if extension.GetOpArrayPersistCalc() != nil {
		data.SetSize(data.GetSize() + extension.GetOpArrayPersistCalc()(data.GetOpArray()))
	}
}
func ZendExtensionOpArrayPersistHandler(extension *ZendExtension, data *ZendExtensionPersistData) {
	if extension.GetOpArrayPersist() != nil {
		var size int = extension.GetOpArrayPersist()(data.GetOpArray(), data.GetMem())
		if size != 0 {
			data.SetMem(any((*byte)(data.GetMem() + size)))
			data.SetSize(data.GetSize() + size)
		}
	}
}
func ZendExtensionsOpArrayPersistCalc(op_array *types2.ZendOpArray) int {
	if (ZendExtensionFlags & ZEND_EXTENSIONS_HAVE_OP_ARRAY_PERSIST_CALC) != 0 {
		var data ZendExtensionPersistData
		data.SetOpArray(op_array)
		data.SetSize(0)
		data.SetMem(nil)
		ZendExtensions.ApplyWithArgument(LlistApplyWithArgFuncT(ZendExtensionOpArrayPersistCalcHandler), &data)
		return data.GetSize()
	}
	return 0
}
func ZendExtensionsOpArrayPersist(op_array *types2.ZendOpArray, mem any) int {
	if (ZendExtensionFlags & ZEND_EXTENSIONS_HAVE_OP_ARRAY_PERSIST) != 0 {
		var data ZendExtensionPersistData
		data.SetOpArray(op_array)
		data.SetSize(0)
		data.SetMem(mem)
		ZendExtensions.ApplyWithArgument(LlistApplyWithArgFuncT(ZendExtensionOpArrayPersistHandler), &data)
		return data.GetSize()
	}
	return 0
}
