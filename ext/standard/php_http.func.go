package standard

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

func PhpUrlEncodeHash(ht *types.Array, formstr *zend.SmartStr) int {
	return PhpUrlEncodeHashEx(ht, formstr, nil, 0, nil, 0, nil, 0, nil)
}
