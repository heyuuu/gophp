// <<generate>>

package standard

import (
	"sik/zend"
)

func PhpUrlEncodeHash(ht *zend.HashTable, formstr *zend.SmartStr) int {
	return PhpUrlEncodeHashEx(ht, formstr, nil, 0, nil, 0, nil, 0, nil)
}
