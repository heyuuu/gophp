package builtin

import "unsafe"

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func FlagMatch[T integer](flags T, flag T) bool { return flags&flag != 0 }
func FlagMatchNum[T integer](flags T, flag T) T { return flags & flag }

func ForceCastPtr[T any, N integer](ptr N) *T {
	return (*T)(unsafe.Pointer(uintptr(ptr)))
}

func ForceUintPtr[T any](ptr *T) uintptr {
	return uintptr(unsafe.Pointer(ptr))
}
