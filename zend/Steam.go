package zend

/**
 * Stream 抽象接口
 */
type IStream interface {
	Read(len_ int) []byte
	Close()
	Isatty() bool
	FileSize() int
}

