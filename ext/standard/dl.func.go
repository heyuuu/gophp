package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func GET_DL_ERROR() __auto__ { return zend.DL_ERROR() }
func PhpLoadShlib(path *byte, errp **byte) any {
	var handle any
	var err *byte
	handle = zend.DL_LOAD(path)
	if !handle {
		err = GET_DL_ERROR()
		*errp = zend.Estrdup(err)
		GET_DL_ERROR()
	}
	return handle
}
func PhpLoadExtension(filename *byte) int {
	var handle any
	var libpath *byte
	var module_entry *zend.ZendModuleEntry
	var get_module func() *zend.ZendModuleEntry
	var error_type int
	var slash_suffix int = 0
	var extension_dir *byte
	var err1 *byte
	var err2 *byte
	extension_dir = zend.INI_STR("extension_dir")
	error_type = faults.E_CORE_WARNING

	/* Check if passed filename contains directory separators */

	if strchr(filename, '/') != nil || strchr(filename, zend.DEFAULT_SLASH) != nil {

		/* Passing modules with full path is not supported for dynamically loaded extensions */

		libpath = zend.Estrdup(filename)
	} else if extension_dir != nil && extension_dir[0] {
		slash_suffix = zend.IS_SLASH(extension_dir[strlen(extension_dir)-1])

		/* Try as filename first */
		if slash_suffix != 0 {
			core.Spprintf(&libpath, 0, "%s%s", extension_dir, filename)
		} else {
			core.Spprintf(&libpath, 0, "%s%c%s", extension_dir, zend.DEFAULT_SLASH, filename)
		}
	} else {
		return types.FAILURE
	}
	handle = PhpLoadShlib(libpath, &err1)
	if !handle {

		/* Now, consider 'filename' as extension name and build file name */

		var orig_libpath *byte = libpath
		if slash_suffix != 0 {
			core.Spprintf(&libpath, 0, "%s"+core.PHP_SHLIB_EXT_PREFIX+"%s."+core.PHP_SHLIB_SUFFIX, extension_dir, filename)
		} else {
			core.Spprintf(&libpath, 0, "%s%c"+core.PHP_SHLIB_EXT_PREFIX+"%s."+core.PHP_SHLIB_SUFFIX, extension_dir, zend.DEFAULT_SLASH, filename)
		}
		handle = PhpLoadShlib(libpath, &err2)
		if !handle {
			core.PhpErrorDocref(nil, error_type, "Unable to load dynamic library '%s' (tried: %s (%s), %s (%s))", filename, orig_libpath, err1, libpath, err2)
			zend.Efree(orig_libpath)
			zend.Efree(err1)
			zend.Efree(libpath)
			zend.Efree(err2)
			return types.FAILURE
		}
		zend.Efree(orig_libpath)
		zend.Efree(err1)
	}
	zend.Efree(libpath)
	get_module = (func() *zend.ZendModuleEntry)(zend.DL_FETCH_SYMBOL(handle, "get_module"))

	/* Some OS prepend _ to symbol names while their dynamic linker
	 * does not do that automatically. Thus we check manually for
	 * _get_module. */

	if get_module == nil {
		get_module = (func() *zend.ZendModuleEntry)(zend.DL_FETCH_SYMBOL(handle, "_get_module"))
	}
	if get_module == nil {
		if zend.DL_FETCH_SYMBOL(handle, "zend_extension_entry") || zend.DL_FETCH_SYMBOL(handle, "_zend_extension_entry") {
			zend.DL_UNLOAD(handle)
			core.PhpErrorDocref(nil, error_type, "Invalid library (appears to be a Zend Extension, try loading using zend_extension=%s from php.ini)", filename)
			return types.FAILURE
		}
		zend.DL_UNLOAD(handle)
		core.PhpErrorDocref(nil, error_type, "Invalid library (maybe not a PHP library) '%s'", filename)
		return types.FAILURE
	}
	module_entry = get_module()
	module_entry.SetModuleNumber(zend.ZendNextFreeModule())
	module_entry.SetHandle(handle)
	if b.Assign(&module_entry, zend.ZendRegisterModuleEx(module_entry)) == nil {
		zend.DL_UNLOAD(handle)
		return types.FAILURE
	}
	return types.SUCCESS
}
func ZmInfoDl(zend_module *zend.ZendModuleEntry) {
	PhpInfoPrintTableRow(2, "Dynamic Library Support", "enabled")
}
