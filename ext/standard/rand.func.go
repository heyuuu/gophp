// <<generate>>

package standard

import (
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func PhpSrand(seed zend.ZendLong) { PhpMtSrand(seed) }
func PhpRand() zend.ZendLong      { return PhpMtRand() }
func ZifRand(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var min zend.ZendLong
	var max zend.ZendLong
	var argc int = executeData.NumArgs()
	if argc == 0 {
		return_value.SetLong(PhpMtRand() >> 1)
		return
	}
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
			zend.Z_PARAM_PROLOGUE(0, 0)
			if !zend.ZendParseArgLong00(_arg, &min) {
				_expected_type = argparse.Z_EXPECTED_LONG
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if !zend.ZendParseArgLong00(_arg, &max) {
				_expected_type = argparse.Z_EXPECTED_LONG
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if max < min {
		return_value.SetLong(PhpMtRandCommon(max, min))
		return
	}
	return_value.SetLong(PhpMtRandCommon(min, max))
	return
}
