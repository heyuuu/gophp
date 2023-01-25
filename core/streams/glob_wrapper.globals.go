// <<generate>>

package streams

import (
	"sik/core"
)

const GLOB_ONLYDIR = 1 << 30
const GLOB_FLAGMASK = ^GLOB_ONLYDIR

var PhpGlobStreamOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpGlobStreamRead, PhpGlobStreamClose, nil, "glob", PhpGlobStreamRewind, nil, nil, nil}
var PhpGlobStreamWrapperOps core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{nil, nil, nil, nil, PhpGlobStreamOpener, "glob", nil, nil, nil, nil, nil}
var PhpGlobStreamWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&PhpGlobStreamWrapperOps, nil, 0}
