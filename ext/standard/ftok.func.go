// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func ZifFtok(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var pathname *byte
	var proj *byte
	var pathname_len int
	var proj_len int
	var k key_t
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			pathname, pathname_len = fp.ParsePath()
			proj, proj_len = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if pathname_len == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Pathname is invalid")
		return_value.SetLong(-1)
		return
	}
	if proj_len != 1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Project identifier is invalid")
		return_value.SetLong(-1)
		return
	}
	if core.PhpCheckOpenBasedir(pathname) != 0 {
		return_value.SetLong(-1)
		return
	}
	k = ftok(pathname, proj[0])
	if k == -1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "ftok() failed - %s", strerror(errno))
	}
	return_value.SetLong(k)
	return
}
