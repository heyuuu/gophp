package standard

import (
	"sik/zend"
)

func PhpRandomBytesThrow(b any, s int) int  { return PhpRandomBytes(b, s, 1) }
func PhpRandomBytesSilent(b any, s int) int { return PhpRandomBytes(b, s, 0) }
func PhpRandomIntThrow(min zend.ZendLong, max zend.ZendLong, result *zend.ZendLong) int {
	return PhpRandomInt(min, max, result, 1)
}
func PhpRandomIntSilent(min zend.ZendLong, max zend.ZendLong, result *zend.ZendLong) int {
	return PhpRandomInt(min, max, result, 0)
}
func RANDOM_G(v int) __auto__ { return RandomGlobals.v }
