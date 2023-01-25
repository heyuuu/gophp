// <<generate>>

package streams

import (
	"sik/core"
)

var PhpUrlDecode func(str *byte, len_ int) int
var PhpStreamMemoryOps core.PhpStreamOps = core.PhpStreamOps{PhpStreamMemoryWrite, PhpStreamMemoryRead, PhpStreamMemoryClose, PhpStreamMemoryFlush, "MEMORY", PhpStreamMemorySeek, PhpStreamMemoryCast, PhpStreamMemoryStat, PhpStreamMemorySetOption}
var PhpStreamTempOps core.PhpStreamOps = core.PhpStreamOps{PhpStreamTempWrite, PhpStreamTempRead, PhpStreamTempClose, PhpStreamTempFlush, "TEMP", PhpStreamTempSeek, PhpStreamTempCast, PhpStreamTempStat, PhpStreamTempSetOption}
var PhpStreamRfc2397Ops core.PhpStreamOps = core.PhpStreamOps{PhpStreamTempWrite, PhpStreamTempRead, PhpStreamTempClose, PhpStreamTempFlush, "RFC2397", PhpStreamTempSeek, PhpStreamTempCast, PhpStreamTempStat, PhpStreamTempSetOption}
var PhpStreamRfc2397Wops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpStreamUrlWrapRfc2397, nil, nil, nil, nil, "RFC2397", nil, nil, nil, nil, nil}
var PhpStreamRfc2397Wrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&PhpStreamRfc2397Wops, nil, 1}
