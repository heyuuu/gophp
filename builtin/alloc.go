package builtin

func Malloc(size int, args ...any) any {
	// todo 单纯作为申请内存的标识，后续需移除
	return nil
}

func Calloc(size int, args ...any) any {
	// todo 单纯作为申请内存的标识，后续需移除
	return nil
}

func Realloc(ptr any, size int, args ...any) any {
	// todo 单纯作为申请内存的标识，后续需移除
	return nil
}

func Free(ptr any) {
	// todo 单纯作为内存释放的标识，后续需移除
}

func Strdup(s *byte) *byte {
	// todo 单纯作为内存释放的标识，后续需移除
	return nil
}
func Strndup(s *byte, length int) *byte {
	// todo 单纯作为内存释放的标识，后续需移除
	return nil
}
