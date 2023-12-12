package operators

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

// OperatorHandler
type OperatorHandler interface {
	Precision() int
	Error(level perr.ErrorType, message string)
	ThrowError(exceptionCe *types.Class, message string)
	ThrowException(exceptionCe *types.Class, message string)
	NewObject(properties *types.Array) *types.Object
	ObjectGetArray(obj *types.Object) *types.Array
	HasException() bool
}
