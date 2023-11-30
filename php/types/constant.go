package types

const (
	// Case Sensitive
	// 大小写敏感，默认是开启的，用户通过define()定义的始终是区分大小 写的，通过扩展定义的可以自由选择
	ConstCs = 1 << 0

	// Persistent
	// 持久化的，只有通过扩展、内核定义的才支持，这种常量不会在request结束时清理掉
	ConstPersistent = 1 << 1

	// Allow compile-time substitution
	// 允许编译时替换，编译时如果发现有地方在读取常量的值，那么编译器会尝试直接替换为常量值，而不是在执行时再去读取，目前这个flag只有TRUE、 FALSE、NULL三个常量在使用
	ConstCtSubst = 1 << 2

	// Can't be saved in file cache
	ConstNoFileCache = 1 << 3
)

type Constant struct {
	value Zval
	name  string
	flags uint32
}

func NewConstant(name string, value *Zval, flags uint32) *Constant {
	c := &Constant{name: name, value: *value, flags: flags}
	return c
}

func (c Constant) Value() Zval {
	return c.value
}
