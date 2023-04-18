package standard

import b "github.com/heyuuu/gophp/builtin"

func PhpMblen(ptr *byte, len_ int) int {
	// 返回第一个多字节字符的长度
	str := b.CastStr(ptr, len_)
	return PhpMblenEx(str)
}

func PhpMblenEx(str string) int {
	// 返回第一个多字节字符的长度
	for i, _ := range str {
		if i != 0 {
			return i
		}
	}
	return len(str)
}
