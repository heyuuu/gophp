package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/plog"
	"github.com/heyuuu/gophp/php/types"
	"log"
	"strings"
)

// functions
func Error(ctx *Context, typ perr.ErrorType, message string) {
	filename, lineno := errorGetFilenameLineno(ctx, typ)
	ErrorTrigger(ctx, perr.NewAt(typ, message, filename, lineno))
}
func ErrorNoreturn(ctx *Context, typ perr.ErrorType, message string) {
	Error(ctx, typ, message)
	/* Should never reach this. */
	panic(perr.Unreachable())
}
func ErrorAt(ctx *Context, typ perr.ErrorType, lineno uint32, message string) {
	filename, _ := errorGetFilenameLineno(ctx, typ)
	ErrorTrigger(ctx, perr.NewAt(typ, message, filename, lineno))
}
func ErrorAtNoreturn(ctx *Context, typ perr.ErrorType, lineno uint32, message string) {
	ErrorAt(ctx, typ, lineno, message)
	/* Should never reach this. */
	panic(perr.Unreachable())
}

func ErrorDocRef(ctx *Context, docRef string, typ perr.ErrorType, buffer string, params ...string) {
	// todo
	message := fmt.Sprintf("%s: %s", buffer, strings.Join(params, ","))
	Error(ctx, typ, message)
}

func ThrowError(ctx *Context, exceptionCe *types.Class, message string) {
	// todo
	Error(ctx, perr.E_ERROR, message)
}

func errorGetFilenameLineno(ctx *Context, typ perr.ErrorType) (string, uint32) {
	// todo
	return "", 0
}

func ErrorTrigger(ctx *Context, err perr.Error) {
	errorVaList(ctx, err.Type, err.Filename, err.Lineno, err.Message)
}

func errorVaList(ctx *Context, typ perr.ErrorType, errorFilename string, errorLineno uint32, message string) {
	ErrorCb(ctx, typ, errorFilename, errorLineno, message)
}

func ErrorCb(ctx *Context, typ perr.ErrorType, errorFilename string, errorLineno uint32, message string) {
	var error_type_str string
	var syslog_type_int plog.Level = plog.Notice
	switch typ {
	case perr.E_ERROR,
		perr.E_CORE_ERROR,
		perr.E_COMPILE_ERROR,
		perr.E_USER_ERROR:
		error_type_str = "Fatal error"
		syslog_type_int = plog.Error
	case perr.E_RECOVERABLE_ERROR:
		error_type_str = "Recoverable fatal error"
		syslog_type_int = plog.Error
	case perr.E_WARNING,
		perr.E_CORE_WARNING,
		perr.E_COMPILE_WARNING,
		perr.E_USER_WARNING:
		error_type_str = "Warning"
		syslog_type_int = plog.Warning
	case perr.E_PARSE:
		error_type_str = "Parse error"
		syslog_type_int = plog.Error
	case perr.E_NOTICE,
		perr.E_USER_NOTICE:
		error_type_str = "Notice"
		syslog_type_int = plog.Notice
	case perr.E_STRICT:
		error_type_str = "Strict Standards"
		syslog_type_int = plog.Info
	case perr.E_DEPRECATED,
		perr.E_USER_DEPRECATED:
		error_type_str = "Deprecated"
		syslog_type_int = plog.Info
	default:
		error_type_str = "Unknown error"
	}
	// todo 待修复
	if errorFilename == "" {
		errorFilename = "__UNKNOWN_FILE__"
	}

	//logBuffer := fmt.Sprintf("PHP %s:  %s in %s on line %d", error_type_str, message, errorFilename, errorLineno)
	logBuffer := fmt.Sprintf("\n%s: %s in %s on line %d", error_type_str, message, errorFilename, errorLineno)
	PhpLogErrWithSeverity(ctx, logBuffer, syslog_type_int)
}

func PhpLogErrWithSeverity(ctx *Context, logMessage string, syslogTypeInt plog.Level) {
	if !ctx.EG().ErrorSuppress() {
		ctx.WriteString(logMessage + "\n")
	}
	log.Println(logMessage)
}
