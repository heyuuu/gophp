package zend

func ZendLoadExtension(path *byte) int {
	panic("暂不支持动态链接扩展")
	//var handle any
	//handle = DL_LOAD(path)
	//if !handle {
	//	log.Printf("Failed loading %s:  %s\n", path, DL_ERROR())
	//	return types.FAILURE
	//}
	//return ZendLoadExtensionHandle(handle, path)
}
func ZendLoadExtensionHandle(handle any, path *byte) int {
	panic("暂不支持动态链接扩展")
	//var new_extension *ZendExtension
	//new_extension = (*ZendExtension)(DL_FETCH_SYMBOL(handle, "zend_extension_entry"))
	//if new_extension == nil {
	//	new_extension = (*ZendExtension)(DL_FETCH_SYMBOL(handle, "_zend_extension_entry"))
	//}
	//
	//if ZendExtensions.Register(new_extension, handle) {
	//	return types.SUCCESS
	//} else {
	//	log.Printf("Cannot load %s - it was already loaded\n", new_extension.GetNameStr())
	//
	//	/* See http://support.microsoft.com/kb/190351 */
	//	DL_UNLOAD(handle)
	//	return types.FAILURE
	//}
}
