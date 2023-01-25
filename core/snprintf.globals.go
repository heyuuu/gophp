// <<generate>>

package core

type BoolInt = int
type BooleanE = int

const (
	NO  = 0
	YES = 1
)

var Php0cvt func(value float64, ndigit int, dec_point byte, exponent byte, buf *byte) *byte

const Slprintf = ApPhpSlprintf
const Vslprintf = ApPhpVslprintf
const Snprintf = ApPhpSnprintf
const Vsnprintf = ApPhpVsnprintf

type LengthModifierE = int

const (
	LM_STD = 0
	LM_INTMAX_T
	LM_PTRDIFF_T
	LM_LONG_LONG
	LM_SIZE_T
	LM_LONG
	LM_LONG_DOUBLE
	LM_PHP_INT_T
)

type WideInt = long__long
type UWideInt = unsigned__long__long

const FORMAT_CONV_MAX_PRECISION = 500
const LCONV_DECIMAL_POINT = (*lconv).decimal_point
const FALSE = 0
const TRUE = 1
const NUL = '0'
const INT_NULL = (*int)(0)
const S_NULL = "(null)"
const S_NULL_LEN = 6
const FLOAT_DIGITS = 6
const EXPONENT_LENGTH = 10
const NDIG = 320
const NUM_BUF_SIZE = 2048

type Buffy = BufArea
