package core

func PHP_OUTPUT_HANDLER_INITBUF_SIZE(s int) __auto__ {
	if s > 1 {
		return s + PHP_OUTPUT_HANDLER_ALIGNTO_SIZE - s%PHP_OUTPUT_HANDLER_ALIGNTO_SIZE
	} else {
		return PHP_OUTPUT_HANDLER_DEFAULT_SIZE
	}
}
func OG__() *ZendOutputGlobals            { return &OutputGlobals }
func PHPWRITE(str *byte, str_len int) int { return PhpOutputWrite(str, str_len) }
func PHPWRITE_H(str *byte, str_len int) int {
	return PhpOutputWriteUnbuffered(str, str_len)
}
func PUTC(c __auto__) int { return PhpOutputWrite((*byte)(&c), 1) }
func PUTS(str string) {
	var __str *byte = str
	PhpOutputWrite(__str, strlen(__str))
}
func PhpOutputTeardown() {
	PhpOutputEndAll()
	PhpOutputDeactivate()
	PhpOutputShutdown()
}
