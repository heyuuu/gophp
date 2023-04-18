package zend

import (
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

func strEscape(str string) string {
	var replacer = strings.NewReplacer("\\\\", "\\", "\\'", "'")
	return replacer.Replace(str)
}

func castZendStringPtr(str *string) *types.String {
	if str != nil {
		return types.NewString(*str)
	}
	return nil
}
func castStrPtr(str *types.String) *string {
	if str != nil {
		var s = str.GetStr()
		return &s
	}
	return nil
}
