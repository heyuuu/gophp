package streams

import (
	"sik/core"
)

var PhpUrlDecode func(str *byte, len_ int) int

/* Memory streams use a dynamic memory buffer to emulate a stream.
 * You can use php_stream_memory_open to create a readonly stream
 * from an existing memory buffer.
 */

var PhpStreamMemoryOps core.PhpStreamOps = core.MakePhpStreamOps(PhpStreamMemoryWrite, PhpStreamMemoryRead, PhpStreamMemoryClose, PhpStreamMemoryFlush, "MEMORY", PhpStreamMemorySeek, PhpStreamMemoryCast, PhpStreamMemoryStat, PhpStreamMemorySetOption)

var PhpStreamTempOps core.PhpStreamOps = core.MakePhpStreamOps(PhpStreamTempWrite, PhpStreamTempRead, PhpStreamTempClose, PhpStreamTempFlush, "TEMP", PhpStreamTempSeek, PhpStreamTempCast, PhpStreamTempStat, PhpStreamTempSetOption)
var PhpStreamRfc2397Ops core.PhpStreamOps = core.MakePhpStreamOps(PhpStreamTempWrite, PhpStreamTempRead, PhpStreamTempClose, PhpStreamTempFlush, "RFC2397", PhpStreamTempSeek, PhpStreamTempCast, PhpStreamTempStat, PhpStreamTempSetOption)
var PhpStreamRfc2397Wops core.PhpStreamWrapperOps = core.MakePhpStreamWrapperOps(PhpStreamUrlWrapRfc2397, nil, nil, nil, nil, "RFC2397", nil, nil, nil, nil, nil)
var PhpStreamRfc2397Wrapper core.PhpStreamWrapper = core.MakePhpStreamWrapper(&PhpStreamRfc2397Wops, nil, 1)
