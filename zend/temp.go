package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func Z_OBJ_HANDLE_P(obj *types.Zval) uint32 {
	return obj.Object().GetHandle()
}
