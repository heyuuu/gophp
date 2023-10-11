package builtin

import "unsafe"

func Cast[T any, P any](ptr *P) *T {
	return (*T)(unsafe.Pointer(ptr))
}

func CastPtr[T any, N integer](ptr N) *T {
	return (*T)(unsafe.Pointer(uintptr(ptr)))
}

func CastUintptr[T any](ptr *T) uintptr {
	return uintptr(unsafe.Pointer(ptr))
}

func CastBytes[I integer](str *byte, len_ I) []byte {
	// todo 此段代码仅表意，实际不应依赖此实现
	return []byte(CastStr(str, len_))
}

func CastStr[I integer](str *byte, len_ I) string {
	// todo 此段代码仅表意，实际不应依赖此实现 (因为无法保证 *byte 后续内存有效)
	if str == nil || len_ == 0 {
		return ""
	}

	var bytes = make([]byte, len_)
	var ptr = uintptr(unsafe.Pointer(str))
	for i := uint(0); i < uint(len_); i++ {
		bytes[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
	}

	return string(bytes)
}

func CastStrAuto(str *byte) string {
	// todo 此段代码仅做占位，后续应替换代码
	return CastStr(str, StrPtrlen(str))
}

func StrPtrlen(str *byte) int {
	// todo 此段代码仅做占位，后续应替换代码
	// C 中通过 *byte 获取字符串长度
	return 0
}

func CastStrPtr(str string) *byte {
	// todo 此段代码仅做占位，表示获取字符串头部指针，后续应替换代码
	return nil
}

func CastSlice[T any, I integer](start *T, len_ I) []T {
	// todo 此段代码仅表意，实际不应依赖此实现
	return *Cast[[]T](start)
}
