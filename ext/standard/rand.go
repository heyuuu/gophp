package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/zpp"
	"math"
	"math/rand"
	"time"
)

const PHP_MT_RAND_MAX int = math.MaxInt32
const MT_RAND_MT19937 = 0
const MT_RAND_PHP = 1

func RegisterMtRandConstants(ctx *php.Context, moduleNumber int) {
	php.RegisterConstant(ctx, moduleNumber, "MT_RAND_MT19937", php.Long(MT_RAND_MT19937))
	php.RegisterConstant(ctx, moduleNumber, "MT_RAND_PHP", php.Long(MT_RAND_PHP))
}

func getRand(ctx *php.Context) *rand.Rand {
	return BG(ctx).GetRandGenerator()
}

func randGenerateSeed() int64 {
	return time.Now().UnixNano()
}

func PhpMtRand(ctx *php.Context) uint32 {
	return BG(ctx).GetRandGenerator().Uint32()
}
func PhpMtRandRange(ctx *php.Context, min int, max int) int {
	php.Assert(min <= max)
	return min + getRand(ctx).Intn(max-min+1)
}

//@zif(alias="srand")
func ZifMtSrand(ctx *php.Context, _ zpp.Opt, seed_ *int, mode_ *int) {
	var seed int64
	if seed_ == nil {
		seed = randGenerateSeed()
	} else {
		seed = int64(*seed_)
	}
	BG(ctx).InitRandGenerator(seed)
}

//@zif(alias="getrandmax")
func ZifMtGetrandmax() int {
	return PHP_MT_RAND_MAX
}

func ZifMtRand(ctx *php.Context, _ zpp.Opt, min_ *int, max_ *int) (int, bool) {
	var min = lang.Option(min_, 0)
	var max = lang.Option(max_, PHP_MT_RAND_MAX)
	if max < min {
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("max(%d) is smaller than min(%d)", max, min))
		return 0, false
	}
	return PhpMtRandRange(ctx, min, max), true
}

func ZifRand(ctx *php.Context, _ zpp.Opt, min_ *int, max_ *int) int {
	var min = lang.Option(min_, 0)
	var max = lang.Option(max_, PHP_MT_RAND_MAX)
	if max < min {
		min, max = max, min
	}
	return PhpMtRandRange(ctx, min, max)
}
