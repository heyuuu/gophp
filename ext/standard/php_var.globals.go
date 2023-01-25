// <<generate>>

package standard

import (
	"sik/zend"
)

type PhpSerializeDataT *PhpSerializeData
type PhpUnserializeDataT *PhpUnserializeData

var PhpVarUnserializeRef func(rval *zend.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT) int
var PhpVarUnserializeIntern func(rval *zend.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT) int
