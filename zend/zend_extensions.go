// <<generate>>

package zend

import (
	b "sik/builtin"
	r "sik/runtime"
)

// Source: <Zend/zend_extensions.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define ZEND_EXTENSIONS_H

// # include "zend_compile.h"

// # include "zend_build.h"

/*
The constants below are derived from ext/opcache/ZendAccelerator.h

You can use the following macro to check the extension API version for compatibilities:

#define    ZEND_EXTENSION_API_NO_5_0_X __special__     220040412
#define    ZEND_EXTENSION_API_NO_5_1_X __special__     220051025
#define    ZEND_EXTENSION_API_NO_5_2_X __special__     220060519
#define    ZEND_EXTENSION_API_NO_5_3_X __special__     220090626
#define    ZEND_EXTENSION_API_NO_5_4_X __special__     220100525
#define    ZEND_EXTENSION_API_NO_5_5_X __special__     220121212
#define    ZEND_EXTENSION_API_NO_5_6_X __special__     220131226
#define    ZEND_EXTENSION_API_NO_7_0_X __special__     320151012

#if ZEND_EXTENSION_API_NO < ZEND_EXTENSION_API_NO_5_5_X
   // do something for php versions lower than 5.5.x
#endif
*/

const ZEND_EXTENSION_API_NO = 320190902

// #define ZEND_EXTENSION_BUILD_ID       "API" ZEND_TOSTR ( ZEND_EXTENSION_API_NO ) ZEND_BUILD_TS ZEND_BUILD_DEBUG ZEND_BUILD_SYSTEM ZEND_BUILD_EXTRA

/* Typedef's for zend_extension function pointers */

type StartupFuncT func(extension *ZendExtension) int
type ShutdownFuncT func(extension *ZendExtension)
type ActivateFuncT func()
type DeactivateFuncT func()
type MessageHandlerFuncT func(message int, arg any)
type OpArrayHandlerFuncT func(op_array *ZendOpArray)
type StatementHandlerFuncT func(frame *ZendExecuteData)
type FcallBeginHandlerFuncT func(frame *ZendExecuteData)
type FcallEndHandlerFuncT func(frame *ZendExecuteData)
type OpArrayCtorFuncT func(op_array *ZendOpArray)
type OpArrayDtorFuncT func(op_array *ZendOpArray)
type OpArrayPersistCalcFuncT func(op_array *ZendOpArray) int
type OpArrayPersistFuncT func(op_array *ZendOpArray, mem any) int

const ZEND_EXTMSG_NEW_EXTENSION = 1

// #define ZEND_EXTENSION() ZEND_EXT_API zend_extension_version_info extension_version_info = { ZEND_EXTENSION_API_NO , ZEND_EXTENSION_BUILD_ID }

// #define STANDARD_ZEND_EXTENSION_PROPERTIES       NULL , NULL , NULL , NULL , NULL , NULL , NULL , NULL , NULL , - 1

// #define COMPAT_ZEND_EXTENSION_PROPERTIES       NULL , NULL , NULL , NULL , NULL , NULL , NULL , NULL , - 1

// #define BUILD_COMPAT_ZEND_EXTENSION_PROPERTIES       NULL , NULL , NULL , NULL , NULL , NULL , NULL , - 1

var ZendExtensions ZendLlist

const ZEND_EXTENSIONS_HAVE_OP_ARRAY_CTOR = 1 << 0
const ZEND_EXTENSIONS_HAVE_OP_ARRAY_DTOR = 1 << 1
const ZEND_EXTENSIONS_HAVE_OP_ARRAY_HANDLER = 1 << 2
const ZEND_EXTENSIONS_HAVE_OP_ARRAY_PERSIST_CALC = 1 << 3
const ZEND_EXTENSIONS_HAVE_OP_ARRAY_PERSIST = 1 << 4

// Source: <Zend/zend_extensions.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// # include "zend_extensions.h"

var ZendExtensionFlags uint32 = 0
var ZendOpArrayExtensionHandles int = 0
var LastResourceNumber int

func ZendLoadExtension(path *byte) int {
	var handle any
	handle = DL_LOAD(path)
	if !handle {
		r.Fprintf(stderr, "Failed loading %s:  %s\n", path, DL_ERROR())
		return FAILURE
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
		r.Fprintf(stderr, "%s doesn't appear to be a valid Zend extension\n", path)

		/* See http://support.microsoft.com/kb/190351 */

		DL_UNLOAD(handle)
		return FAILURE
	}

	/* allow extension to proclaim compatibility with any Zend version */

	if extension_version_info.GetZendExtensionApiNo() != ZEND_EXTENSION_API_NO && (new_extension.GetApiNoCheck() == nil || new_extension.GetApiNoCheck()(ZEND_EXTENSION_API_NO) != SUCCESS) {
		if extension_version_info.GetZendExtensionApiNo() > ZEND_EXTENSION_API_NO {
			r.Fprintf(stderr, "%s requires Zend Engine API version %d.\n"+"The Zend Engine API version %d which is installed, is outdated.\n\n", new_extension.GetName(), extension_version_info.GetZendExtensionApiNo(), ZEND_EXTENSION_API_NO)

			/* See http://support.microsoft.com/kb/190351 */

			DL_UNLOAD(handle)
			return FAILURE
		} else if extension_version_info.GetZendExtensionApiNo() < ZEND_EXTENSION_API_NO {
			r.Fprintf(stderr, "%s requires Zend Engine API version %d.\n"+"The Zend Engine API version %d which is installed, is newer.\n"+"Contact %s at %s for a later version of %s.\n\n", new_extension.GetName(), extension_version_info.GetZendExtensionApiNo(), ZEND_EXTENSION_API_NO, new_extension.GetAuthor(), new_extension.GetURL(), new_extension.GetName())

			/* See http://support.microsoft.com/kb/190351 */

			DL_UNLOAD(handle)
			return FAILURE
		}
	} else if strcmp("API"+"ZEND_EXTENSION_API_NO"+ZEND_BUILD_TS, extension_version_info.GetBuildId()) && (new_extension.GetBuildIdCheck() == nil || new_extension.GetBuildIdCheck()("API"+"ZEND_EXTENSION_API_NO"+ZEND_BUILD_TS) != SUCCESS) {
		r.Fprintf(stderr, "Cannot load %s - it was built with configuration %s, whereas running engine is %s\n", new_extension.GetName(), extension_version_info.GetBuildId(), "API"+"ZEND_EXTENSION_API_NO"+ZEND_BUILD_TS)

		/* See http://support.microsoft.com/kb/190351 */

		DL_UNLOAD(handle)
		return FAILURE
	} else if ZendGetExtension(new_extension.GetName()) != nil {
		r.Fprintf(stderr, "Cannot load %s - it was already loaded\n", new_extension.GetName())

		/* See http://support.microsoft.com/kb/190351 */

		DL_UNLOAD(handle)
		return FAILURE
	}
	return ZendRegisterExtension(new_extension, handle)
}
func ZendRegisterExtension(new_extension *ZendExtension, handle any) int {
	var extension ZendExtension
	extension = *new_extension
	extension.SetHandle(handle)
	ZendExtensionDispatchMessage(ZEND_EXTMSG_NEW_EXTENSION, &extension)
	ZendLlistAddElement(&ZendExtensions, &extension)
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

	return SUCCESS

	/*fprintf(stderr, "Loaded %s, version %s\n", extension.name, extension.version);*/
}
func ZendExtensionShutdown(extension *ZendExtension) {
	if extension.GetShutdown() != nil {
		extension.GetShutdown()(extension)
	}
}
func ZendExtensionStartup(extension *ZendExtension) int {
	if extension.GetStartup() != nil {
		if extension.GetStartup()(extension) != SUCCESS {
			return 1
		}
		ZendAppendVersionInfo(extension)
	}
	return 0
}
func ZendStartupExtensionsMechanism() int {
	/* Startup extensions mechanism */

	ZendLlistInit(&ZendExtensions, b.SizeOf("zend_extension"), (func(any))(ZendExtensionDtor), 1)
	ZendOpArrayExtensionHandles = 0
	LastResourceNumber = 0
	return SUCCESS
}
func ZendStartupExtensions() int {
	ZendLlistApplyWithDel(&ZendExtensions, (func(any) int)(ZendExtensionStartup))
	return SUCCESS
}
func ZendShutdownExtensions() {
	ZendLlistApply(&ZendExtensions, LlistApplyFuncT(ZendExtensionShutdown))
	ZendLlistDestroy(&ZendExtensions)
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
	if LastResourceNumber < ZEND_MAX_RESERVED_RESOURCES {
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
func ZendExtensionsOpArrayPersistCalc(op_array *ZendOpArray) int {
	if (ZendExtensionFlags & ZEND_EXTENSIONS_HAVE_OP_ARRAY_PERSIST_CALC) != 0 {
		var data ZendExtensionPersistData
		data.SetOpArray(op_array)
		data.SetSize(0)
		data.SetMem(nil)
		ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionOpArrayPersistCalcHandler), &data)
		return data.GetSize()
	}
	return 0
}
func ZendExtensionsOpArrayPersist(op_array *ZendOpArray, mem any) int {
	if (ZendExtensionFlags & ZEND_EXTENSIONS_HAVE_OP_ARRAY_PERSIST) != 0 {
		var data ZendExtensionPersistData
		data.SetOpArray(op_array)
		data.SetSize(0)
		data.SetMem(mem)
		ZendLlistApplyWithArgument(&ZendExtensions, LlistApplyWithArgFuncT(ZendExtensionOpArrayPersistHandler), &data)
		return data.GetSize()
	}
	return 0
}
