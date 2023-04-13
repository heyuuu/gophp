package core

func PHP_OUTPUT_HANDLER_INITBUF_SIZE(s int) __auto__ {
	if s > 1 {
		return s + PHP_OUTPUT_HANDLER_ALIGNTO_SIZE - s%PHP_OUTPUT_HANDLER_ALIGNTO_SIZE
	} else {
		return PHP_OUTPUT_HANDLER_DEFAULT_SIZE
	}
}
func OG__() *ZendOutputGlobals { return &OutputGlobals }
func PUTS(str string) int      { return PhpOutputWrite(str) }
func PUTS_H(str string)        { PhpOutputWriteUnbuffered(str) }
func PUTC(c byte)              { PhpOutputWrite(string(c)) }
func PhpOutputTeardown() {
	PhpOutputEndAll()
	PhpOutputDeactivate()
	PhpOutputShutdown()
}
