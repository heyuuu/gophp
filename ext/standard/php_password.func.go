package standard

import (
	"github.com/heyuuu/gophp/php/types"
)

func PhpPasswordAlgoIdentify(hash *types.String) *PhpPasswordAlgo {
	return PhpPasswordAlgoIdentifyEx(hash, PhpPasswordAlgoDefault())
}
