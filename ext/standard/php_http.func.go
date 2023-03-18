// <<generate>>

package standard

import (
	"sik/zend"
	"sik/zend/types"
)

func PhpUrlEncodeHash(ht *types.HashTable, formstr *zend.SmartStr) int {
	return PhpUrlEncodeHashEx(ht, formstr, nil, 0, nil, 0, nil, 0, nil)
}
