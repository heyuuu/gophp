package core

type BoolInt = int
type BooleanE = int

const (
	NO  = 0
	YES = 1
)

const Slprintf = ApPhpSlprintf
const Snprintf = ApPhpSnprintf

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

type WideInt = int64
type UWideInt = uint64

const FORMAT_CONV_MAX_PRECISION = 500

const LCONV_DECIMAL_POINT = lconv.decimal_point

const FALSE = 0
const NUL = '0'
const S_NULL = "(null)"
const S_NULL_LEN = 6
const FLOAT_DIGITS = 6
const EXPONENT_LENGTH = 10

const NDIG = 320

const NUM_BUF_SIZE = 2048

type Buffy = BufArea
