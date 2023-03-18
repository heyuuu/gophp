package zend

import (
	"log"
	"sik/zend/types"
	"strconv"
	"strings"
)

func strCaseEquals(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	}
	return strings.EqualFold(str1, str2)
}

//func memcmp(str1 []byte, str2 []byte, len_ int) ZendBool {
//	var str1_ = string(str1[:len_])
//	var str2_ = string(str2[:len_])
//	return intBool(str1_ == str2_)
//}

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

func castZendStringPtr(str *string) *types.ZendString {
	if str != nil {
		return types.NewZendString(*str)
	}
	return nil
}
func castStrPtr(str *types.ZendString) *string {
	if str != nil {
		var s = str.GetStr()
		return &s
	}
	return nil
}
