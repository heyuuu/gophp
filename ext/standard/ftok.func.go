package standard

import (
	"sik/core"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZifFtok(executeData zpp.DefEx, return_value zpp.DefReturn, pathname *types.Zval, proj *types.Zval) {
	var pathname *byte
	var proj *byte
	var pathname_len int
	var proj_len int
	var k key_t
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			pathname, pathname_len = fp.ParsePath()
			proj, proj_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if pathname_len == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Pathname is invalid")
		return_value.SetLong(-1)
		return
	}
	if proj_len != 1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Project identifier is invalid")
		return_value.SetLong(-1)
		return
	}
	if core.PhpCheckOpenBasedir(pathname) != 0 {
		return_value.SetLong(-1)
		return
	}
	k = ftok(pathname, proj[0])
	if k == -1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "ftok() failed - %s", strerror(errno))
	}
	return_value.SetLong(k)
	return
}
