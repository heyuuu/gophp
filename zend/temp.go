package zend

import "sik/zend/types"

func Z_OBJ_HANDLE_P(obj *types.Zval) uint32 {
	return obj.GetObj().GetHandle()
}
