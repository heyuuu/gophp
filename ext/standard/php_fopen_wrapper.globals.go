// <<generate>>

package standard

import (
	"sik/core"
)

var PhpStreamOutputOps core.PhpStreamOps = core.PhpStreamOps{PhpStreamOutputWrite, PhpStreamOutputRead, PhpStreamOutputClose, nil, "Output", nil, nil, nil, nil}

type PhpStreamInput = PhpStreamInputT

var PhpStreamInputOps core.PhpStreamOps = core.PhpStreamOps{PhpStreamInputWrite, PhpStreamInputRead, PhpStreamInputClose, PhpStreamInputFlush, "Input", PhpStreamInputSeek, nil, nil, nil}
var PhpStdioWops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpStreamUrlWrapPhp, nil, nil, nil, nil, "PHP", nil, nil, nil, nil, nil}
var PhpStreamPhpWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&PhpStdioWops, nil, 0}
