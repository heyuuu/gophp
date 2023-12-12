package operators

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

func ConvertObjectToType(obj *types.Object, ctype types.ZvalType) Val {
	if result, ok := obj.Cast(ctype); ok {
		return result
	} else if obj.CanCast() {
		php.Error(perr.E_RECOVERABLE_ERROR, fmt.Sprintf("Object of class %s could not be converted to %s", obj.CeName(), types.ZendGetTypeByConst(ctype)))
	}
	return nil
}
