package standard

import (
	"github.com/heyuuu/gophp/core"
)

var PhpStreamOutputOps core.PhpStreamOps = core.MakePhpStreamOps(PhpStreamOutputWrite, PhpStreamOutputRead, PhpStreamOutputClose, nil, "Output", nil, nil, nil, nil)

type PhpStreamInput = PhpStreamInputT

var PhpStreamInputOps core.PhpStreamOps = core.MakePhpStreamOps(PhpStreamInputWrite, PhpStreamInputRead, PhpStreamInputClose, PhpStreamInputFlush, "Input", PhpStreamInputSeek, nil, nil, nil)
var PhpStdioWops core.PhpStreamWrapperOps = core.MakePhpStreamWrapperOps(PhpStreamUrlWrapPhp, nil, nil, nil, nil, "PHP", nil, nil, nil, nil, nil)
var PhpStreamPhpWrapper core.PhpStreamWrapper = core.MakePhpStreamWrapper(&PhpStdioWops, nil, 0)
