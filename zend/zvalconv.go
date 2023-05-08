package zend

import "github.com/heyuuu/gophp/php/types"

func IZendIsTrue(op *types.Zval) int { return types.IntBool(ZvalIsTrue(op)) }

func ZvalIsTrue(op *types.Zval) bool {
again:
	switch op.GetType() {
	case types.IS_TRUE:
		return true
	case types.IS_LONG:
		return op.Long() != 0
	case types.IS_DOUBLE:
		return op.Double() != 0
	case types.IS_STRING:
		str := op.StringVal()
		return str != "" && str != "0"
	case types.IS_ARRAY:
		return op.Array().Len() != 0
	case types.IS_OBJECT:
		if types.Z_OBJ_HT_P(op).GetCastObject() == ZendStdCastObjectTostring {
			return true
		} else {
			return ZendObjectIsTrue(op)
		}
	case types.IS_RESOURCE:
		return types.Z_RES_HANDLE_P(op) != 0
	case types.IS_REFERENCE:
		op = types.Z_REFVAL_P(op)
		goto again
	}
	return false
}
