package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
	"math/rand"
	"time"
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
	strTokState   StrTokState
	randGenerator *rand.Rand
}

func (bg *BasicGlobals) StrTokState() *StrTokState {
	return &bg.strTokState
}

func (bg *BasicGlobals) ResetRandGenerator() {
	bg.randGenerator = nil
}
func (bg *BasicGlobals) InitRandGenerator(seed int64) {
	bg.randGenerator = rand.New(rand.NewSource(seed))
}
func (bg *BasicGlobals) GetRandGenerator() *rand.Rand {
	if bg.randGenerator == nil {
		seed := time.Now().UnixNano()
		bg.InitRandGenerator(seed)
	}
	return bg.randGenerator
}
