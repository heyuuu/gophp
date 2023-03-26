package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZifReadlink(executeData zpp.Ex, return_value zpp.Ret, filename *types.Zval) {
	var link *byte
	var link_len int
	var buff []byte
	var ret ssize_t
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			link, link_len = fp.ParsePath()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if core.PhpCheckOpenBasedir(link) != 0 {
		return_value.SetFalse()
		return
	}
	ret = zend.PhpSysReadlink(link, buff, core.MAXPATHLEN-1)
	if ret == -1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s", strerror(errno))
		return_value.SetFalse()
		return
	}

	/* Append NULL to the end of the string */

	buff[ret] = '0'
	return_value.SetStringVal(b.CastStr(buff, ret))
	return
}
func ZifLinkinfo(executeData zpp.Ex, return_value zpp.Ret, filename *types.Zval) {
	var link *byte
	var dirname *byte
	var link_len int
	var sb zend.ZendStatT
	var ret int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			link, link_len = fp.ParsePath()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	dirname = zend.Estrndup(link, link_len)
	PhpDirname(dirname, link_len)
	if core.PhpCheckOpenBasedir(dirname) != 0 {
		zend.Efree(dirname)
		return_value.SetFalse()
		return
	}
	ret = zend.VCWD_LSTAT(link, &sb)
	if ret == -1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s", strerror(errno))
		zend.Efree(dirname)
		return_value.SetLong(int64(-1))
		return
	}
	zend.Efree(dirname)
	return_value.SetLong(zend.ZendLong(sb.st_dev))
	return
}
func ZifSymlink(executeData zpp.Ex, return_value zpp.Ret, target *types.Zval, link *types.Zval) {
	var topath *byte
	var frompath *byte
	var topath_len int
	var frompath_len int
	var ret int
	var source_p []byte
	var dest_p []byte
	var dirname []byte
	var len_ int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			topath, topath_len = fp.ParsePath()
			frompath, frompath_len = fp.ParsePath()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if core.ExpandFilepath(frompath, source_p) == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "No such file or directory")
		return_value.SetFalse()
		return
	}
	memcpy(dirname, source_p, b.SizeOf("source_p"))
	len_ = PhpDirname(dirname, strlen(dirname))
	if core.ExpandFilepathEx(topath, dest_p, dirname, len_) == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "No such file or directory")
		return_value.SetFalse()
		return
	}
	if streams.PhpStreamLocateUrlWrapper(source_p, nil, core.STREAM_LOCATE_WRAPPERS_ONLY) != nil || streams.PhpStreamLocateUrlWrapper(dest_p, nil, core.STREAM_LOCATE_WRAPPERS_ONLY) != nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to symlink to a URL")
		return_value.SetFalse()
		return
	}
	if core.PhpCheckOpenBasedir(dest_p) != 0 {
		return_value.SetFalse()
		return
	}
	if core.PhpCheckOpenBasedir(source_p) != 0 {
		return_value.SetFalse()
		return
	}

	/* For the source, an expanded path must be used (in ZTS an other thread could have changed the CWD).
	 * For the target the exact string given by the user must be used, relative or not, existing or not.
	 * The target is relative to the link itself, not to the CWD. */

	ret = zend.PhpSysSymlink(topath, source_p)
	if ret == -1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s", strerror(errno))
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifLink(executeData zpp.Ex, return_value zpp.Ret, target *types.Zval, link *types.Zval) {
	var topath *byte
	var frompath *byte
	var topath_len int
	var frompath_len int
	var ret int
	var source_p []byte
	var dest_p []byte
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			topath, topath_len = fp.ParsePath()
			frompath, frompath_len = fp.ParsePath()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if core.ExpandFilepath(frompath, source_p) == nil || core.ExpandFilepath(topath, dest_p) == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "No such file or directory")
		return_value.SetFalse()
		return
	}
	if streams.PhpStreamLocateUrlWrapper(source_p, nil, core.STREAM_LOCATE_WRAPPERS_ONLY) != nil || streams.PhpStreamLocateUrlWrapper(dest_p, nil, core.STREAM_LOCATE_WRAPPERS_ONLY) != nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to link to a URL")
		return_value.SetFalse()
		return
	}
	if core.PhpCheckOpenBasedir(dest_p) != 0 {
		return_value.SetFalse()
		return
	}
	if core.PhpCheckOpenBasedir(source_p) != 0 {
		return_value.SetFalse()
		return
	}
	ret = zend.PhpSysLink(topath, frompath)
	if ret == -1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s", strerror(errno))
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
