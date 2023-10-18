package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"math"
	"math/rand"
	"time"
)

const PHP_MT_RAND_MAX int = math.MaxInt32
const MT_RAND_MT19937 = 0
const MT_RAND_PHP = 1

func getRand() *rand.Rand {
	return BG__().GetRandGenerator()
}

func randGenerateSeed() int64 {
	return time.Now().UnixNano()
}

func PhpMtRand() uint32 {
	return BG__().GetRandGenerator().Uint32()
}
func PhpMtRandRange(min int, max int) int {
	b.Assert(min <= max)
	return min + getRand().Intn(max-min)
}

func ZmStartupMtRand(moduleNumber int) int {
	zend.RegisterLongConstant("MT_RAND_MT19937", MT_RAND_MT19937, zend.CONST_CS|zend.CONST_PERSISTENT, moduleNumber)
	zend.RegisterLongConstant("MT_RAND_PHP", MT_RAND_PHP, zend.CONST_CS|zend.CONST_PERSISTENT, moduleNumber)
	return types.SUCCESS
}

//@zif -alias srand
func ZifMtSrand(_ zpp.Opt, seed_ *int, mode_ *int) {
	var seed int64
	if seed_ == nil {
		seed = randGenerateSeed()
	} else {
		seed = int64(*seed_)
	}
	BG__().InitRandGenerator(seed)
}

//@zif -alias getrandmax
func ZifMtGetrandmax() int {
	return PHP_MT_RAND_MAX
}

func ZifMtRand(_ zpp.Opt, min_ *int, max_ *int) (int, bool) {
	var min = b.Option(min_, 0)
	var max = b.Option(max_, PHP_MT_RAND_MAX)
	if max < min {
		core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("max(%d) is smaller than min(%d)", max, min))
		return 0, false
	}
	return PhpMtRandRange(min, max), true
}

func ZifRand(_ zpp.Opt, min_ *int, max_ *int) int {
	var min = b.Option(min_, 0)
	var max = b.Option(max_, PHP_MT_RAND_MAX)
	if max < min {
		min, max = max, min
	}
	return PhpMtRandRange(min, max)
}
