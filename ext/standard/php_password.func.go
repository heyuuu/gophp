package standard

import (
	"sik/zend/types"
)

func PhpPasswordAlgoIdentify(hash *types.String) *PhpPasswordAlgo {
	return PhpPasswordAlgoIdentifyEx(hash, PhpPasswordAlgoDefault())
}
