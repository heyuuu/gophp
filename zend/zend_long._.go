package zend

import "math"

type ZendLong = int
type ZendUlong = uint
type ZendOffT = int

const ZEND_LONG_MAX = math.MaxInt
const ZEND_LONG_MIN = math.MinInt
const ZEND_ULONG_MAX = math.MaxUint
const SIZEOF_ZEND_LONG = 8

/* Conversion macros. */

const ZEND_LONG_FMT string = "%lld"
const ZEND_ULONG_FMT string = "%llu"
const ZEND_LONG_FMT_SPEC = "lld"
const ZEND_STRTOL_PTR = strtoll
const ZEND_STRTOUL_PTR = strtoull
const ZEND_ABS = imaxabs
const MAX_LENGTH_OF_LONG = 20
const LONG_MIN_DIGITS = "9223372036854775808"

var LongMinDigits []byte = LONG_MIN_DIGITS
