package perr

//go:generate stringer -type=ErrorType
type ErrorType uint32

const (
	E_ERROR             ErrorType = 1 << 0
	E_WARNING           ErrorType = 1 << 1
	E_PARSE             ErrorType = 1 << 2
	E_NOTICE            ErrorType = 1 << 3
	E_CORE_ERROR        ErrorType = 1 << 4
	E_CORE_WARNING      ErrorType = 1 << 5
	E_COMPILE_ERROR     ErrorType = 1 << 6
	E_COMPILE_WARNING   ErrorType = 1 << 7
	E_USER_ERROR        ErrorType = 1 << 8
	E_USER_WARNING      ErrorType = 1 << 9
	E_USER_NOTICE       ErrorType = 1 << 10
	E_STRICT            ErrorType = 1 << 11
	E_RECOVERABLE_ERROR ErrorType = 1 << 12
	E_DEPRECATED        ErrorType = 1 << 13
	E_USER_DEPRECATED   ErrorType = 1 << 14
	E_ALL               ErrorType = E_ERROR | E_WARNING | E_PARSE | E_NOTICE | E_CORE_ERROR | E_CORE_WARNING | E_COMPILE_ERROR | E_COMPILE_WARNING | E_USER_ERROR | E_USER_WARNING | E_USER_NOTICE | E_RECOVERABLE_ERROR | E_DEPRECATED | E_USER_DEPRECATED | E_STRICT
	E_CORE              ErrorType = E_CORE_ERROR | E_CORE_WARNING
)

type Error struct {
	Message  string
	Filename string
	Type     ErrorType
	Lineno   uint32
}

func (e Error) Error() string { return e.Message }

func New(typ ErrorType, message string) Error {
	return Error{Type: typ, Message: message, Filename: "", Lineno: 0}
}

func NewAt(typ ErrorType, message string, filename string, lineno uint32) Error {
	return Error{Type: typ, Message: message, Filename: filename, Lineno: lineno}
}
