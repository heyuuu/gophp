package ctype

func IsAscii(c byte) bool {
	return c <= 0x7f
}

func IsLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
