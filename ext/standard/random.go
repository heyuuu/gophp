package standard

import (
	cryptoRand "crypto/rand"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/assert"
	"math/big"
)

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

	num, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, false
	}

	return min + int(num.Int64()), true
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
