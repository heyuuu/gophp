package zend

import "strings"

func strCaseEquals(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	}
	return strings.EqualFold(str1, str2)
}

func memcmp(str1 []byte, str2 []byte, len_ int) ZendBool {
	var str1_ = string(str1[:len_])
	var str2_ = string(str2[:len_])
	return intBool(str1_ == str2_)
}
