package zpp

/* Parameter parsing API -- andrei */

const (
	FlagQuiet   = 1 << 1
	FlagThrow   = 1 << 2
	FlagOldMode = 1 << 3 // 标识兼容旧 TypeSpec 等价方式
)
