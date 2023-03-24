package standard

func Strnatcmp(a *byte, b *byte) int {
	return StrnatcmpEx(a, strlen(a), b, strlen(b), 0)
}
func Strnatcasecmp(a *byte, b *byte) int {
	return StrnatcmpEx(a, strlen(a), b, strlen(b), 1)
}
func PhpMblen(ptr *byte, len_ int) __auto__ { return mblen(ptr, len_) }
