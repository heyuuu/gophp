// <<generate>>

package standard

import (
	"sik/zend"
)

func PhpPasswordAlgoIdentify(hash *zend.ZendString) *PhpPasswordAlgo {
	return PhpPasswordAlgoIdentifyEx(hash, PhpPasswordAlgoDefault())
}
