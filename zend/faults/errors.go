package faults

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/types"
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
func Bailout() { _zendBailout(__FILE__, __LINE__) }

func _zendBailout(filename *byte, lineno uint32) {
	if zend.EG__().GetBailout() == nil {
		OutputDebugString(1, "%s(%d) : Bailed out without a bailout address!", filename, lineno)
		exit(-1)
	}
	zend.GcProtect(1)
	zend.CG__().SetUncleanShutdown(1)
	zend.CG__().SetActiveClassEntry(nil)
	zend.CG__().SetInCompilation(0)
	zend.EG__().SetCurrentExecuteData(nil)
	zend.LONGJMP(*zend.EG__().GetBailout(), types.FAILURE)
}

func ErrorVaList(type_ int, error_filename *byte, error_lineno uint32, format string, args ...any) {
	var usr_copy va_list
	var params []types.Zval
	var retval types.Zval
	var orig_user_error_handler types.Zval
	var in_compilation types.ZendBool
	var saved_class_entry *types.ClassEntry
	var loop_var_stack zend.ZendStack
	var delayed_oplines_stack zend.ZendStack
	var symbol_table *types.Array
	var orig_fake_scope *types.ClassEntry

	/* Report about uncaught exception in case of fatal errors */

	if zend.EG__().GetException() != nil {
		var ex *zend.ZendExecuteData
		var opline *zend.ZendOp
		switch type_ {
		case E_CORE_ERROR:

		case E_ERROR:

		case E_RECOVERABLE_ERROR:

		case E_PARSE:

		case E_COMPILE_ERROR:

		case E_USER_ERROR:
			ex = zend.CurrEX()
			opline = nil
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
			break
		default:
			break
		}
	}

	/* if we don't have a user defined error handler */

	if zend.EG__().GetUserErrorHandler().IsUndef() || (zend.EG__().GetUserErrorHandlerErrorReporting()&type_) == 0 || zend.EG__().GetErrorHandling() != zend.EH_NORMAL {
		zend.ZendErrorCb(type_, error_filename, error_lineno, format, args)
	} else {
		switch type_ {
		case E_ERROR:

		case E_PARSE:

		case E_CORE_ERROR:

		case E_CORE_WARNING:

		case E_COMPILE_ERROR:

		case E_COMPILE_WARNING:

			/* The error may not be safe to handle in user-space */

			zend.ZendErrorCb(type_, error_filename, error_lineno, format, args)
			break
		default:

			/* Handle the error in user space */

			zend.VaCopy(usr_copy, args)
			params[1].SetString(zend.ZendStrpprintf(0, format, usr_copy))
			va_end(usr_copy)
			params[0].SetLong(type_)
			if error_filename != nil {
				params[2].SetRawString(b.CastStrAuto(error_filename))
			} else {
				params[2].SetNull()
			}
			params[3].SetLong(error_lineno)
			symbol_table = zend.ZendRebuildSymbolTable()

			/* during shutdown the symbol table table can be still null */

			if symbol_table == nil {
				params[4].SetNull()
			} else {
				params[4].SetArray(types.ZendArrayDup(symbol_table))
			}
			types.ZVAL_COPY_VALUE(&orig_user_error_handler, zend.EG__().GetUserErrorHandler())
			zend.EG__().GetUserErrorHandler().SetUndef()

			/* User error handler may include() additinal PHP files.
			 * If an error was generated during comilation PHP will compile
			 * such scripts recursively, but some CG() variables may be
			 * inconsistent. */

			in_compilation = zend.CG__().GetInCompilation()
			if in_compilation != 0 {
				saved_class_entry = zend.CG__().GetActiveClassEntry()
				zend.CG__().SetActiveClassEntry(nil)
				zend.SAVE_STACK(loop_var_stack)
				zend.SAVE_STACK(delayed_oplines_stack)
				zend.CG__().SetInCompilation(0)
			}
			orig_fake_scope = zend.EG__().GetFakeScope()
			zend.EG__().SetFakeScope(nil)
			if zend.CallUserFunction(zend.CG__().GetFunctionTable(), nil, &orig_user_error_handler, &retval, 5, params) == types.SUCCESS {
				if retval.IsNotUndef() {
					if retval.IsFalse() {
						zend.ZendErrorCb(type_, error_filename, error_lineno, format, args)
					}
					zend.ZvalPtrDtor(&retval)
				}
			} else if zend.EG__().GetException() == nil {

				/* The user error handler failed, use built-in error handler */

				zend.ZendErrorCb(type_, error_filename, error_lineno, format, args)

				/* The user error handler failed, use built-in error handler */

			}
			zend.EG__().SetFakeScope(orig_fake_scope)
			if in_compilation != 0 {
				zend.CG__().SetActiveClassEntry(saved_class_entry)
				zend.RESTORE_STACK(loop_var_stack)
				zend.RESTORE_STACK(delayed_oplines_stack)
				zend.CG__().SetInCompilation(1)
			}
			zend.ZvalPtrDtor(&params[4])
			zend.ZvalPtrDtor(&params[2])
			zend.ZvalPtrDtor(&params[1])
			if zend.EG__().GetUserErrorHandler().IsUndef() {
				types.ZVAL_COPY_VALUE(zend.EG__().GetUserErrorHandler(), &orig_user_error_handler)
			} else {
				zend.ZvalPtrDtor(&orig_user_error_handler)
			}
			break
		}
	}
	if type_ == E_PARSE {

		/* eval() errors do not affect exit_status */

		if !(zend.CurrEX() != nil && zend.CurrEX().GetFunc() != nil && zend.ZEND_USER_CODE(zend.CurrEX().GetFunc().GetType()) && zend.CurrEX().GetOpline().GetOpcode() == zend.ZEND_INCLUDE_OR_EVAL && zend.CurrEX().GetOpline().GetExtendedValue() == zend.ZEND_EVAL) {
			zend.EG__().SetExitStatus(255)
		}

		/* eval() errors do not affect exit_status */

	}
}

func GetFilenameLineno(type_ int, filename **byte, lineno *uint32) {
	/* Obtain relevant filename and lineno */

	switch type_ {
	case E_CORE_ERROR:

	case E_CORE_WARNING:
		*filename = nil
		*lineno = 0
		break
	case E_PARSE:

	case E_COMPILE_ERROR:

	case E_COMPILE_WARNING:

	case E_ERROR:

	case E_NOTICE:

	case E_STRICT:

	case E_DEPRECATED:

	case E_WARNING:

	case E_USER_ERROR:

	case E_USER_WARNING:

	case E_USER_NOTICE:

	case E_USER_DEPRECATED:

	case E_RECOVERABLE_ERROR:
		if zend.ZendIsCompiling() != 0 {
			*filename = zend.ZendGetCompiledFilename().GetVal()
			*lineno = zend.ZendGetCompiledLineno()
		} else if zend.ZendIsExecuting() != 0 {
			*filename = zend.ZendGetExecutedFilename()
			if (*filename)[0] == '[' {
				*filename = nil
				*lineno = 0
			} else {
				*lineno = zend.ZendGetExecutedLineno()
			}
		} else {
			*filename = nil
			*lineno = 0
		}
		break
	default:
		*filename = nil
		*lineno = 0
		break
	}
	if (*filename) == nil {
		*filename = "Unknown"
	}
}
func ErrorAt(type_ int, filename *byte, lineno uint32, format string, _ ...any) {
	var args va_list
	if filename == nil {
		var dummy_lineno uint32
		GetFilenameLineno(type_, &filename, &dummy_lineno)
	}
	va_start(args, format)
	ErrorVaList(type_, filename, lineno, format, args)
	va_end(args)
}
func ErrorEx(typ int, message string) {
	// todo
}
func Error(type_ int, format string, args ...any) {
	var filename *byte
	var lineno uint32
	GetFilenameLineno(type_, &filename, &lineno)
	ErrorVaList(type_, filename, lineno, format, args)
}
func ErrorAtNoreturn(type_ int, filename *byte, lineno uint32, format string, _ ...any) {
	var args va_list
	if filename == nil {
		var dummy_lineno uint32
		GetFilenameLineno(type_, &filename, &dummy_lineno)
	}
	va_start(args, format)
	ErrorVaList(type_, filename, lineno, format, args)
	va_end(args)

	/* Should never reach this. */

	abort()

	/* Should never reach this. */
}
func ErrorNoreturn(type_ int, format string, _ ...any) {
	var filename *byte
	var lineno uint32
	var args va_list
	GetFilenameLineno(type_, &filename, &lineno)
	va_start(args, format)
	ErrorVaList(type_, filename, lineno, format, args)
	va_end(args)

	/* Should never reach this. */

	abort()

	/* Should never reach this. */
}
func ThrowErrorEx(exception_ce *types.ClassEntry, message string) {
	if exception_ce != nil {
		if zend.InstanceofFunction(exception_ce, ZendCeError) == 0 {
			Error(E_NOTICE, "Error exceptions must be derived from Error")
			exception_ce = ZendCeError
		}
	} else {
		exception_ce = ZendCeError
	}

	/* Marker used to disable exception generation during preloading. */

	if zend.EG__().GetException() == any(uintPtr-1) {
		return
	}

	//TODO: we can't convert compile-time errors to exceptions yet???

	if zend.CurrEX() != nil && zend.CG__().GetInCompilation() == 0 {
		ThrowException(exception_ce, message, 0)
	} else {
		Error(E_ERROR, "%s", message)
	}
}
func ThrowError(exception_ce *types.ClassEntry, format string, args ...any) {
	message := zend.ZendSprintf(format, args...)
	ThrowErrorEx(exception_ce, message)
}
func TypeError(format string, args ...any) {
	message := zend.ZendSprintf(format, args...)
	ThrowException(ZendCeTypeError, message, 0)
	zend.Efree(message)
}

func InternalTypeErrorEx(throwException bool, message string) {
	if throwException {
		ThrowException(ZendCeTypeError, message, 0)
	} else {
		Error(E_WARNING, "%s", message)
	}
}
func InternalTypeError(throw_exception bool, format string, args ...any) {
	message := zend.ZendSprintf(format, args...)
	if throw_exception {
		ThrowException(ZendCeTypeError, message, 0)
	} else {
		Error(E_WARNING, "%s", message)
	}
}
func InternalArgumentCountError(throw_exception bool, format string, args ...any) {
	message := zend.ZendSprintf(format, args...)
	if throw_exception {
		ThrowException(ZendCeArgumentCountError, message, 0)
	} else {
		Error(E_WARNING, "%s", message)
	}
}
func OutputDebugString(trigger_break types.ZendBool, format string, _ ...any) {}
