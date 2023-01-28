package builtin

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func FlagMatch[T integer](flags T, flag T) bool { return flags&flag != 0 }
func FlagMatchNum[T integer](flags T, flag T) T { return flags & flag }

func HashStr(str string) uint {
	return HashBytes([]byte(str))
}

func HashBytes(bytes []byte) uint {
	var hash uint = 5381
	for _, c := range bytes {
		hash = hash<<5 + hash + uint(c)
	}
	/* Hash value can't be zero, so we always set the high bit */
	return hash | -0x8000000000000000
}

func Free(ptr any) {
	// todo 单纯作为内存失败的标识
}

func EmptyString(len_ int) string {
	return string(make([]byte, len_))
}
