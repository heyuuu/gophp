package builtin

func FlagMatch[T integer](flags T, flag T) bool { return flags&flag != 0 }
func FlagMatchNum[T integer](flags T, flag T) T { return flags & flag }

func HashStr(str string) uint {
	return HashBytes([]byte(str))
}

func EqualsAny[T comparable](value T, expected ...T) bool {
	for _, v := range expected {
		if value == v {
			return true
		}
	}
	return false
}

func CopySlice[T any](src []T) []T {
	if len(src) == 0 {
		return nil
	}

	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

/*
 * DJBX33A (Daniel J. Bernstein, Times 33 with Addition)
 *
 * This is Daniel J. Bernstein's popular `times 33' hash function as
 * posted by him years ago on comp.lang.c. It basically uses a function
 * like ``hash(i) = hash(i-1) * 33 + str[i]''. This is one of the best
 * known hash functions for strings. Because it is both computed very
 * fast and distributes very well.
 *
 * The magic of number 33, i.e. why it works better than many other
 * constants, prime or not, has never been adequately explained by
 * anyone. So I try an explanation: if one experimentally tests all
 * multipliers between 1 and 256 (as RSE did now) one detects that even
 * numbers are not usable at all. The remaining 128 odd numbers
 * (except for the number 1) work more or less all equally well. They
 * all distribute in an acceptable way and this way fill a hash table
 * with an average percent of approx. 86%.
 *
 * If one compares the Chi^2 values of the variants, the number 33 not
 * even has the best value. But the number 33 and a few other equally
 * good numbers like 17, 31, 63, 127 and 129 have nevertheless a great
 * advantage to the remaining numbers in the large set of possible
 * multipliers: their multiply operation can be replaced by a faster
 * operation based on just one shift plus either a single addition
 * or subtraction operation. And because a hash function has to both
 * distribute good _and_ has to be very fast to compute, those few
 * numbers should be preferred and seems to be the reason why Daniel J.
 * Bernstein also preferred it.
 *
 *
 *                  -- Ralf S. Engelschall <rse@engelschall.com>
 */
func HashBytes(bytes []byte) uint {
	var hash uint = 5381
	for _, c := range bytes {
		hash = hash<<5 + hash + uint(c)
	}
	/* Hash value can't be zero, so we always set the high bit */
	return hash | -0x8000000000000000
}

func EmptyString(len_ int) string {
	return string(make([]byte, len_))
}
