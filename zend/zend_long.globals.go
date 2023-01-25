// <<generate>>

package zend

const ZEND_ENABLE_ZVAL_LONG64 = 1

type ZendLong = int64
type ZendUlong = uint64
type ZendOffT = int64

const ZEND_LONG_MAX ZendLong = INT64_MAX
const ZEND_LONG_MIN float = INT64_MIN
const ZEND_ULONG_MAX = UINT64_MAX
const SIZEOF_ZEND_LONG = 8
const ZEND_LTOA_BUF_LEN = 65
const ZEND_LONG_FMT string = "%" + "lld"
const ZEND_ULONG_FMT *byte = "%" + "llu"
const ZEND_XLONG_FMT = "%" + PRIx64
const ZEND_LONG_FMT_SPEC = "lld"
const ZEND_ULONG_FMT_SPEC = "llu"
const ZEND_STRTOL_PTR = strtoll
const ZEND_STRTOUL_PTR = strtoull
const ZEND_ABS = imaxabs
const MAX_LENGTH_OF_LONG = 20
const LONG_MIN_DIGITS = "9223372036854775808"

var LongMinDigits []byte = LONG_MIN_DIGITS

const ZEND_ADDR_FMT = "0x%016zx"
