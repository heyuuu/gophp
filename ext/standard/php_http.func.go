package standard

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

func PhpUrlEncodeHash(ht *types.Array, formstr *zend.SmartStr) int {
	return PhpUrlEncodeHashEx(ht, formstr, nil, 0, nil, 0, nil, 0, nil)
}
