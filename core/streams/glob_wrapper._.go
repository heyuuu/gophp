package streams

import (
	"github.com/heyuuu/gophp/core"
)

const GLOB_ONLYDIR = 1 << 30
const GLOB_FLAGMASK = ^GLOB_ONLYDIR

/* {{{ */

var PhpGlobStreamOps core.PhpStreamOps = core.MakePhpStreamOps(nil, PhpGlobStreamRead, PhpGlobStreamClose, nil, "glob", PhpGlobStreamRewind, nil, nil, nil)

/* {{{ php_glob_stream_opener */

var PhpGlobStreamWrapperOps core.PhpStreamWrapperOps = core.MakePhpStreamWrapperOps(nil, nil, nil, nil, PhpGlobStreamOpener, "glob", nil, nil, nil, nil, nil)
var PhpGlobStreamWrapper core.PhpStreamWrapper = core.MakePhpStreamWrapper(&PhpGlobStreamWrapperOps, nil, 0)
