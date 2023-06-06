package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"log"
)

func ZendLoadExtension(path *byte) int {
	var handle any
	handle = DL_LOAD(path)
	if !handle {
		log.Printf("Failed loading %s:  %s\n", path, DL_ERROR())
		return types.FAILURE
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
		return types.FAILURE
	}

	/* allow extension to proclaim compatibility with any Zend version */
	if extension_version_info.GetZendExtensionApiNo() != ZEND_EXTENSION_API_NO && (new_extension.GetApiNoCheck() == nil || new_extension.GetApiNoCheck()(ZEND_EXTENSION_API_NO) != types.SUCCESS) {
		if extension_version_info.GetZendExtensionApiNo() > ZEND_EXTENSION_API_NO {
			log.Printf("%s requires Zend Engine API version %d.\n"+"The Zend Engine API version %d which is installed, is outdated.\n\n", new_extension.GetNameStr(), extension_version_info.GetZendExtensionApiNo(), ZEND_EXTENSION_API_NO)

			/* See http://support.microsoft.com/kb/190351 */
			DL_UNLOAD(handle)
			return types.FAILURE
		} else if extension_version_info.GetZendExtensionApiNo() < ZEND_EXTENSION_API_NO {
			log.Printf("%s requires Zend Engine API version %d.\nThe Zend Engine API version %d which is installed, is newer.\nContact %s at %s for a later version of %s.\n\n", new_extension.GetNameStr(), extension_version_info.GetZendExtensionApiNo(), ZEND_EXTENSION_API_NO, new_extension.GetAuthor(), new_extension.GetURL(), new_extension.GetName())

			/* See http://support.microsoft.com/kb/190351 */

			DL_UNLOAD(handle)
			return types.FAILURE
		}
	} else if strcmp("API"+"ZEND_EXTENSION_API_NO"+ZEND_BUILD_TS, extension_version_info.GetBuildId()) && (new_extension.GetBuildIdCheck() == nil || new_extension.GetBuildIdCheck()("API"+"ZEND_EXTENSION_API_NO"+ZEND_BUILD_TS) != types.SUCCESS) {
		log.Printf("Cannot load %s - it was built with configuration %s, whereas running engine is %s\n", new_extension.GetNameStr(), extension_version_info.GetBuildId(), "API"+"ZEND_EXTENSION_API_NO"+ZEND_BUILD_TS)

		/* See http://support.microsoft.com/kb/190351 */

		DL_UNLOAD(handle)
		return types.FAILURE
	} else if ZendExtensions.Get(new_extension.GetNameStr()) != nil {
		log.Printf("Cannot load %s - it was already loaded\n", new_extension.GetNameStr())

		/* See http://support.microsoft.com/kb/190351 */
		DL_UNLOAD(handle)
		return types.FAILURE
	}
	ZendExtensions.Register(new_extension, handle)
	return types.SUCCESS
}
