package standard

import (
	"github.com/heyuuu/gophp/zend/types"
)

type PhpSerializeDataT *PhpSerializeData
type PhpUnserializeDataT *PhpUnserializeData

var PhpVarUnserializeRef func(rval *types.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT) int
var PhpVarUnserializeIntern func(rval *types.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT) int
