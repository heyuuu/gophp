package zpp

import (
	"sik/zend"
	"sik/zend/types"
)

/**
 * 此处的类型用于 sikgen 脚本生成代码
 */
type (
	DefOpt any
	DefEx  = *zend.ZendExecuteData
	DefRet = *types.Zval
	// Type: 'L', FAST_ZPP: Z_PARAM_STRICT_LONG
	DefStrictLong = int

	DefRefArray = *types.Zval // fp.ParseArrayEx(false, true)
)
