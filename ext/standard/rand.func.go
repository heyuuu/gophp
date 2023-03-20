// <<generate>>

package standard

import (
	"sik/zend"
	"sik/zend/types"
	"sik/zend/zpp"
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

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			min = fp.ParseLong()
			max = fp.ParseLong()
			if fp.HasError() {
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
