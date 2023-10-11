package core

import b "github.com/heyuuu/gophp/builtin"

/* handler ops */
const PHP_OUTPUT_HANDLER_WRITE = 0x0
const PHP_OUTPUT_HANDLER_START = 0x1
const PHP_OUTPUT_HANDLER_CLEAN = 0x2
const PHP_OUTPUT_HANDLER_FLUSH = 0x4
const PHP_OUTPUT_HANDLER_FINAL = 0x8
const PHP_OUTPUT_HANDLER_CONT = PHP_OUTPUT_HANDLER_WRITE
const PHP_OUTPUT_HANDLER_END = PHP_OUTPUT_HANDLER_FINAL

/* handler ability flags */
const PHP_OUTPUT_HANDLER_CLEANABLE = 0x10
const PHP_OUTPUT_HANDLER_FLUSHABLE = 0x20
const PHP_OUTPUT_HANDLER_REMOVABLE = 0x40
const PHP_OUTPUT_HANDLER_STDFLAGS = 0x70

/* handler status flags */
const PHP_OUTPUT_HANDLER_STARTED = 0x1000
const PHP_OUTPUT_HANDLER_DISABLED = 0x2000
const PHP_OUTPUT_HANDLER_PROCESSED = 0x4000

/* handler op return values */
type PhpOutputHandlerStatusT = int

const (
	PHP_OUTPUT_HANDLER_FAILURE = iota
	PHP_OUTPUT_HANDLER_SUCCESS
	PHP_OUTPUT_HANDLER_NO_DATA
)

/* php_output_stack_pop() flags */
const PHP_OUTPUT_POP_TRY = 0x0
const PHP_OUTPUT_POP_FORCE = 0x1
const PHP_OUTPUT_POP_DISCARD = 0x10
const PHP_OUTPUT_POP_SILENT = 0x100

/* handler hooks */
const PHP_OUTPUT_HANDLER_ALIGNTO_SIZE = 0x1000
const PHP_OUTPUT_HANDLER_DEFAULT_SIZE = 0x4000

/* old-style, stateless callback */
type PhpOutputHandlerFuncT func(output *byte, output_len int, handled_output **byte, handled_output_len *int, mode int)

/* new-style, opaque context callback */
type PhpOutputHandlerContextFuncT func(handlerContext any, outputContext *PhpOutputContext) int

type OutputHandlerT func(output string, mode int) (handledOutput string)

// todo temp, need remove
func wrapOutputHandler(h PhpOutputHandlerFuncT) OutputHandlerT {
	if h == nil {
		return nil
	}
	return func(output string, mode int) (handledOutput string) {
		var outputPtr *byte
		var outputLen int
		h(b.CastStrPtr(output), len(output), &outputPtr, &outputLen, mode)
		return b.CastStr(outputPtr, outputLen)
	}
}
