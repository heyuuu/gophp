package faults

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/operators"
)

/**
 * constants and global variables
 */
const E_ERROR = 1 << 0
const E_WARNING = 1 << 1
const E_PARSE zend.ZendLong = 1 << 2
const E_NOTICE = 1 << 3
const E_CORE_ERROR = 1 << 4
const E_CORE_WARNING = 1 << 5
const E_COMPILE_ERROR = 1 << 6
const E_COMPILE_WARNING = 1 << 7
const E_USER_ERROR zend.ZendLong = 1 << 8
const E_USER_WARNING zend.ZendLong = 1 << 9
const E_USER_NOTICE zend.ZendLong = 1 << 10
const E_STRICT zend.ZendLong = 1 << 11
const E_RECOVERABLE_ERROR zend.ZendLong = 1 << 12
const E_DEPRECATED = 1 << 13
const E_USER_DEPRECATED zend.ZendLong = 1 << 14
const E_ALL zend.ZendLong = E_ERROR | E_WARNING | E_PARSE | E_NOTICE | E_CORE_ERROR | E_CORE_WARNING | E_COMPILE_ERROR | E_COMPILE_WARNING | E_USER_ERROR | E_USER_WARNING | E_USER_NOTICE | E_RECOVERABLE_ERROR | E_DEPRECATED | E_USER_DEPRECATED | E_STRICT
const E_CORE = E_CORE_ERROR | E_CORE_WARNING

/**
 * functions
 */
func Bailout() {
	zend.CG__().SetUncleanShutdown(1)
	zend.CG__().SetActiveClassEntry(nil)
	zend.CG__().SetInCompilation(0)
	zend.EG__().SetCurrentExecuteData(nil)
	throw()
}

func GetException() *types.Object {
	return zend.EG__().GetException()
}
func HasException() bool {
	return GetException() != nil
}
func GetUserErrorHandler() *types.Zval {
	return zend.EG__().GetUserErrorHandler()
}
func GetUserErrorHandlerErrorReporting() int {
	return zend.EG__().GetUserErrorHandlerErrorReporting()
}

func errorCb(typ int, errorFilename string, errorLineno uint32, message string) {
	core.PhpErrorCb(typ, errorFilename, errorLineno, message)
}

func errorVaList(typ int, errorFilename string, errorLineno uint32, message string) {
	/* Report about uncaught exception in case of fatal errors */
	if HasException() {
		switch typ {
		case E_CORE_ERROR,
			E_ERROR,
			E_RECOVERABLE_ERROR,
			E_PARSE,
			E_COMPILE_ERROR,
			E_USER_ERROR:
			var ex *zend.ZendExecuteData = zend.CurrEX()
			var opline *types.ZendOp = nil
			for ex != nil && (ex.GetFunc() == nil || !(zend.ZEND_USER_CODE(ex.GetFunc().GetType()))) {
				ex = ex.GetPrevExecuteData()
			}
			if ex != nil && ex.GetOpline().GetOpcode() == zend.ZEND_HANDLE_EXCEPTION && zend.EG__().GetOplineBeforeException() != nil {
				opline = zend.EG__().GetOplineBeforeException()
			}
			ExceptionError(zend.EG__().GetException(), E_WARNING)
			zend.EG__().SetException(nil)
			if opline != nil {
				ex.SetOpline(opline)
			}
		}
	}

	/* if we don't have a user defined error handler */
	if GetUserErrorHandler().IsUndef() || GetUserErrorHandlerErrorReporting()&typ == 0 || zend.EG__().GetErrorHandling() != zend.EH_NORMAL {
		errorCb(typ, errorFilename, errorLineno, message)
	} else {
		switch typ {
		case E_ERROR,
			E_PARSE,
			E_CORE_ERROR,
			E_CORE_WARNING,
			E_COMPILE_ERROR,
			E_COMPILE_WARNING:
			errorCb(typ, errorFilename, errorLineno, message)
		default:
			var params [5]types.Zval
			params[0].SetLong(typ)
			params[1].SetString(message)
			if errorFilename != "" {
				params[2].SetString(errorFilename)
			} else {
				params[2].SetNull()
			}
			params[3].SetLong(int(errorLineno))

			symbolTable := zend.ZendRebuildSymbolTable()
			/* during shutdown the symbol table table can be still null */
			if symbolTable == nil {
				params[4].SetNull()
			} else {
				params[4].SetArray(types.ZendArrayDup(symbolTable))
			}

			var retval types.Zval
			var orig_user_error_handler types.Zval
			var orig_fake_scope *types.ClassEntry
			var cgBackup *zend.CGBackupStack

			types.ZVAL_COPY_VALUE(&orig_user_error_handler, zend.EG__().GetUserErrorHandler())
			zend.EG__().GetUserErrorHandler().SetUndef()

			/* User error handler may include() additinal PHP files.
			 * If an error was generated during comilation PHP will compile
			 * such scripts recursively, but some CG() variables may be
			 * inconsistent. */

			cgBackup = zend.CG__().BackupStack()
			orig_fake_scope = zend.EG__().GetFakeScope()
			zend.EG__().SetFakeScope(nil)
			if zend.CallUserFunction(nil, &orig_user_error_handler, &retval, 5, params[:]) == types.SUCCESS {
				if retval.IsNotUndef() {
					if retval.IsFalse() {
						errorCb(typ, errorFilename, errorLineno, message)
					}
					// zend.ZvalPtrDtor(&retval)
				}
			} else if zend.EG__().GetException() == nil {
				/* The user error handler failed, use built-in error handler */
				errorCb(typ, errorFilename, errorLineno, message)
			}
			zend.EG__().SetFakeScope(orig_fake_scope)
			zend.CG__().RestorePostError(cgBackup)

			if zend.EG__().GetUserErrorHandler().IsUndef() {
				types.ZVAL_COPY_VALUE(zend.EG__().GetUserErrorHandler(), &orig_user_error_handler)
			} else {
				// zend.ZvalPtrDtor(&orig_user_error_handler)
			}

		}
	}
	if typ == E_PARSE {
		/* eval() errors do not affect exit_status */
		if !(zend.CurrEX() != nil && zend.CurrEX().GetFunc() != nil && zend.ZEND_USER_CODE(zend.CurrEX().GetFunc().GetType()) && zend.CurrEX().GetOpline().GetOpcode() == zend.ZEND_INCLUDE_OR_EVAL && zend.CurrEX().GetOpline().GetExtendedValue() == zend.ZEND_EVAL) {
			zend.EG__().SetExitStatus(255)
		}
	}
}

func getFilenameLineno(typ int) (string, uint32) {
	/* Obtain relevant filename and lineno */
	switch typ {
	case E_CORE_ERROR, E_CORE_WARNING:
		return "Unknown", 0
	case E_PARSE, E_COMPILE_ERROR, E_COMPILE_WARNING, E_ERROR, E_NOTICE, E_STRICT, E_DEPRECATED,
		E_WARNING, E_USER_ERROR, E_USER_WARNING, E_USER_NOTICE, E_USER_DEPRECATED, E_RECOVERABLE_ERROR:
		if zend.ZendIsCompiling() {
			return zend.ZendGetCompiledFilename(), uint32(zend.ZendGetCompiledLineno())
		} else if zend.ZendIsExecuting() {
			filename := zend.ZendGetExecutedFilename()
			if filename != "" && filename[0] == '[' {
				return "Unknown", 0
			} else {
				return filename, zend.ZendGetExecutedLineno()
			}
		} else {
			return "Unknown", 0
		}
	default:
		return "Unknown", 0
	}
}

func ErrorAt(typ int, filename *string, lineno uint32, format string, args ...any) {
	var filenameStr string
	if filename == nil {
		filenameStr, _ = getFilenameLineno(typ)
	} else {
		filenameStr = *filename
	}
	message := zend.ZendSprintf(format, args...)
	errorVaList(typ, filenameStr, lineno, message)
}
func Error(typ int, format string, args ...any) {
	filename, lineno := getFilenameLineno(typ)
	message := zend.ZendSprintf(format, args...)
	errorVaList(typ, filename, lineno, message)
}
func ErrorAtNoreturn(typ int, filename *string, lineno uint32, format string, args ...any) {
	var filenameStr string
	if filename == nil {
		filenameStr, _ = getFilenameLineno(typ)
	} else {
		filenameStr = *filename
	}
	message := zend.ZendSprintf(format, args...)
	errorVaList(typ, filenameStr, lineno, message)

	/* Should never reach this. */
	panic("unreachable")
}
func ErrorNoreturn(typ int, format string, args ...any) {
	filename, lineno := getFilenameLineno(typ)
	message := zend.ZendSprintf(format, args...)
	errorVaList(typ, filename, lineno, message)

	/* Should never reach this. */
	panic("unreachable")
}
func ThrowErrorEx(exceptionCe *types.ClassEntry, message string) {
	if exceptionCe != nil {
		if operators.InstanceofFunction(exceptionCe, ZendCeError) == 0 {
			Error(E_NOTICE, "Error exceptions must be derived from Error")
			exceptionCe = ZendCeError
		}
	} else {
		exceptionCe = ZendCeError
	}

	/* Marker used to disable exception generation during preloading. */
	if zend.EG__().GetException() == nil {
		return
	}

	//TODO: we can't convert compile-time errors to exceptions yet???
	if zend.CurrEX() != nil && zend.CG__().GetInCompilation() == 0 {
		ThrowException(exceptionCe, message, 0)
	} else {
		Error(E_ERROR, message)
	}
}
func ThrowError(exceptionCe *types.ClassEntry, format string, args ...any) {
	message := zend.ZendSprintf(format, args...)
	ThrowErrorEx(exceptionCe, message)
}
func TypeError(format string, args ...any) {
	message := zend.ZendSprintf(format, args...)
	ThrowException(ZendCeTypeError, message, 0)
}

func InternalTypeErrorEx(throwException bool, message string) {
	if throwException {
		ThrowException(ZendCeTypeError, message, 0)
	} else {
		Error(E_WARNING, message)
	}
}
func InternalTypeError(throwException bool, format string, args ...any) {
	message := zend.ZendSprintf(format, args...)
	if throwException {
		ThrowException(ZendCeTypeError, message, 0)
	} else {
		Error(E_WARNING, message)
	}
}
func InternalArgumentCountError(throwException bool, format string, args ...any) {
	message := zend.ZendSprintf(format, args...)
	if throwException {
		ThrowException(ZendCeArgumentCountError, message, 0)
	} else {
		Error(E_WARNING, message)
	}
}
