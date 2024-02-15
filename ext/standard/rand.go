package standard

import (
	cryptoRand "crypto/rand"
	"fmt"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/zpp"
	"math"
	"math/big"
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

// @zif(alias="srand")
func ZifMtSrand(ctx *php.Context, _ zpp.Opt, seed_ *int, mode_ *int) {
	var seed int64
	if seed_ == nil {
		seed = randGenerateSeed()
	} else {
		seed = int64(*seed_)
	}
	BG(ctx).InitRandGenerator(seed)
}

// @zif(alias="getrandmax")
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

func PhpRandomStringSafe(size int) (string, bool) {
	assert.Assert(size >= 0)
	if size == 0 {
		return "", true
	}

	bytes := make([]byte, size)
	if _, err := cryptoRand.Read(bytes); err != nil {
		return "", false
	}

	return string(bytes), true
}

func PhpRandomIntSafe(min int, max int) (int, bool) {
	assert.Assert(min <= max)
	if min == max {
		return min, true
	}

	bigMin := big.NewInt(int64(min))
	bigMax := big.NewInt(int64(max))
	diff := bigMax.Sub(bigMax, bigMin)

	num, err := cryptoRand.Int(cryptoRand.Reader, diff)
	if err != nil {
		return 0, false
	}
	num = num.Add(num, bigMin)

	return int(num.Int64()), true
}

func ZifRandomBytes(ctx *php.Context, length int) string {
	if length < 1 {
		php.ThrowException(ctx, nil, "Length must be greater than 0", 0)
		return ""
	}

	str, ok := PhpRandomStringSafe(length)
	if !ok {
		return ""
	}
	return str
}
func ZifRandomInt(ctx *php.Context, min int, max int) int {
	if min > max {
		php.ThrowException(ctx, nil, "Minimum value must be less than or equal to the maximum value", 0)
		return 0
	}
	num, ok := PhpRandomIntSafe(min, max)
	if !ok {
		return 0
	}
	return num
}
