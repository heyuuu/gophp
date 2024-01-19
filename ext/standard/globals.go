package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
)

const globalKey = "ext.standard.globals"

func BG(ctx *php.Context) *BasicGlobals {
	bg, ok := php.ContextGetOrInit(ctx, globalKey, func() *BasicGlobals {
		return &BasicGlobals{}
	})
	if !ok {
		panic(perr.Internalf("php.ContextGetOrInit() fail"))
	}
	return bg
}

type StrTokState struct {
	str string
}

type BasicGlobals struct {
	strTokState StrTokState
}

func (bg *BasicGlobals) StrTokState() *StrTokState {
	return &bg.strTokState
}
