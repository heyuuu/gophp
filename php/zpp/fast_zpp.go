package zpp

import (
	"github.com/heyuuu/gophp/php/types"
)

func NewFastParser(ex IExecuteData, minNumArgs int, maxNumArgs int, flags int) IParser {
	return &FastParser{}
}

var _ IParser = (*FastParser)(nil)

type FastParser struct {
	err error
}

func (f *FastParser) HasError() bool { return f.err != nil }
func (f *FastParser) ParseZval() *types.Zval {
	//TODO implement me
	return types.NewZvalNull()
}
