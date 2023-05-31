package standard

import (
	cryptoRand "crypto/rand"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
)

func PhpRandomBytesThrow(b any, s int) int { return PhpRandomBytes(b, s, 1) }
func PhpRandomIntThrow(min zend.ZendLong, max zend.ZendLong, result *zend.ZendLong) int {
	return PhpRandomInt(min, max, result, 1)
}
func RANDOM_G(v int) __auto__ { return RandomGlobals.v }

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
