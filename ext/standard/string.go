package standard

import (
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/zend/types"
)

func PhpStringToupper(s *types.String) *types.String {
	return types.NewString(ascii.StrToUpper(s.GetStr()))
}

func PhpStringTolower(s *types.String) *types.String {
	return types.NewString(ascii.StrToLower(s.GetStr()))
}
