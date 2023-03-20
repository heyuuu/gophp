package zend

import (
	"log"
	"sik/zend/types"
	"strconv"
	"strings"
)

func strEscape(str string) string {
	var replacer = strings.NewReplacer("\\\\", "\\", "\\'", "'")
	return replacer.Replace(str)
}

func strToD(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return value
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
