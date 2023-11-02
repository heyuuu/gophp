package operators

import (
	"fmt"
	"github.com/heyuuu/gophp/php/faults"
	"github.com/heyuuu/gophp/php/types"
)

func ConvertObjectToType(obj *types.Object, ctype types.ZvalType) Val {
	if result, ok := obj.Cast(ctype); ok {
		return result
	} else if obj.CanCast() {
		faults.Error(faults.E_RECOVERABLE_ERROR, fmt.Sprintf("Object of class %s could not be converted to %s", obj.Ce().Name(), types.ZendGetTypeByConst(ctype)))
	}
	return nil
}
