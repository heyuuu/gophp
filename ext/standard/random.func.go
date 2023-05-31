package standard

import (
	cryptoRand "crypto/rand"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"math/big"
)

func PhpRandomStringSafe(size int) (string, bool) {
	b.Assert(size >= 0)
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
	b.Assert(min <= max)
	if min == max {
		return min, true
	}

	num, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, false
	}

	return min + int(num.Int64()), true
}

func ZifRandomBytes(length int) string {
	if length < 1 {
		faults.ThrowException(faults.ZendCeError, "Length must be greater than 0", 0)
		return ""
	}

	str, ok := PhpRandomStringSafe(length)
	if !ok {
		return ""
	}
	return str
}
func ZifRandomInt(min int, max int) int {
	if min > max {
		faults.ThrowException(faults.ZendCeError, "Minimum value must be less than or equal to the maximum value", 0)
		return 0
	}
	num, ok := PhpRandomIntSafe(min, max)
	if !ok {
		return 0
	}
	return num
}
