// <<generate>>

package standard

import (
	"sik/zend/types"
)

func PhpPasswordAlgoIdentify(hash *types.ZendString) *PhpPasswordAlgo {
	return PhpPasswordAlgoIdentifyEx(hash, PhpPasswordAlgoDefault())
}
