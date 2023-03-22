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
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
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
