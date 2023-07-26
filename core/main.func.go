package core

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core/date"
	"github.com/heyuuu/gophp/core/pfmt"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/globals"
	"github.com/heyuuu/gophp/zend/operators"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func SAFE_FILENAME(f __auto__) string {
	if f {
		return f
	} else {
		return "-"
	}
}
func GetSafeCharsetHint() string {
	var hint = SG__().defaultCharset
	codeSet, _ := standard.CheckCodeSet(hint)
	return codeSet
}
func OnSetFacility(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var facility *byte = new_value.GetVal()
	return types.FAILURE
}
func OnSetPrecision(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var i zend.ZendLong
	zend.ZEND_ATOL(i, new_value.GetVal())
	if i >= -1 {
		zend.EG__().SetPrecision(i)
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func OnSetSerializePrecision(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var i zend.ZendLong
	zend.ZEND_ATOL(i, new_value.GetVal())
	if i >= -1 {
		PG__().serialize_precision = i
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func OnChangeMemoryLimit(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var value int
	if new_value != nil {
		value = zend.StrToLongWithUnit(new_value.GetStr())
	} else {
		value = 1 << 30
	}
	if zend.ZendSetMemoryLimit(value) == types.FAILURE {

		/* When the memory limit is reset to the original level during deactivation, we may be
		 * using more memory than the original limit while shutdown is still in progress.
		 * Ignore a failure for now, and set the memory limit when the memory manager has been
		 * shut down and the minimal amount of memory is used. */

		if stage != zend.ZEND_INI_STAGE_DEACTIVATE {
			faults.Error(faults.E_WARNING, "Failed to set memory limit to %zd bytes (Current memory usage is %zd bytes)", value, zend.ZendMemoryUsage(true))
			return types.FAILURE
		}

		/* When the memory limit is reset to the original level during deactivation, we may be
		 * using more memory than the original limit while shutdown is still in progress.
		 * Ignore a failure for now, and set the memory limit when the memory manager has been
		 * shut down and the minimal amount of memory is used. */

	}
	PG__().memory_limit = value
	return types.SUCCESS
}
func OnSetLogFilter(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var filter *byte = new_value.GetVal()
	if !(strcmp(filter, "all")) {
		PG__().syslog_filter = PHP_SYSLOG_FILTER_ALL
		return types.SUCCESS
	}
	if !(strcmp(filter, "no-ctrl")) {
		PG__().syslog_filter = PHP_SYSLOG_FILTER_NO_CTRL
		return types.SUCCESS
	}
	if !(strcmp(filter, "ascii")) {
		PG__().syslog_filter = PHP_SYSLOG_FILTER_ASCII
		return types.SUCCESS
	}
	if !(strcmp(filter, "raw")) {
		PG__().syslog_filter = PHP_SYSLOG_FILTER_RAW
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpDisableFunctions() {
	var s *byte = nil
	var e *byte
	if !(*(zend.INI_STR("disable_functions"))) {
		return
	}
	PG__().disable_functions = strdup(zend.INI_STR("disable_functions"))
	e = PG__().disable_functions
	if e == nil {
		return
	}
	for *e {
		switch *e {
		case ' ':
			fallthrough
		case ',':
			if s != nil {
				*e = '0'
				zend.ZendDisableFunction(b.CastStr(s, e-s))
				s = nil
			}
		default:
			if s == nil {
				s = e
			}
		}
		e++
	}
	if s != nil {
		zend.ZendDisableFunction(b.CastStr(s, e-s))
	}
}
func PhpDisableClasses() {
	var s *byte = nil
	var e *byte
	if !(*(zend.INI_STR("disable_classes"))) {
		return
	}
	PG__().disable_classes = strdup(zend.INI_STR("disable_classes"))
	e = PG__().disable_classes
	for *e {
		switch *e {
		case ' ':
			fallthrough
		case ',':
			if s != nil {
				*e = '0'
				zend.ZendDisableClass(b.CastStr(s, e-s))
				s = nil
			}
		default:
			if s == nil {
				s = e
			}
		}
		e++
	}
	if s != nil {
		zend.ZendDisableClass(b.CastStr(s, e-s))
	}
}
func PhpBinaryInit() {
	var binary_location *byte = nil
	if SM__().GetExecutableLocation() != nil {
		binary_location = (*byte)(zend.Malloc(MAXPATHLEN))
		if binary_location != nil && !(strchr(SM__().GetExecutableLocation(), '/')) {
			var envpath *byte
			var path *byte
			var found = 0
			if lang.Assign(&envpath, getenv("PATH")) != nil {
				var search_dir *byte
				var search_path []*byte
				var last *byte = nil
				var s zend.ZendStatT
				path = zend.Estrdup(envpath)
				search_dir = PhpStrtokR(path, ":", &last)
				for search_dir != nil {
					Snprintf(search_path, MAXPATHLEN, "%s/%s", search_dir, SM__().GetExecutableLocation())
					if zend.VCWD_REALPATH(search_path, binary_location) != nil && !(zend.VCWD_ACCESS(binary_location, X_OK)) && zend.VCWD_STAT(binary_location, &s) == 0 && zend.S_ISREG(s.st_mode) {
						found = 1
						break
					}
					search_dir = PhpStrtokR(nil, ":", &last)
				}
				zend.Efree(path)
			}
			if found == 0 {
				zend.Free(binary_location)
				binary_location = nil
			}
		} else if zend.VCWD_REALPATH(SM__().GetExecutableLocation(), binary_location) == nil || zend.VCWD_ACCESS(binary_location, X_OK) {
			zend.Free(binary_location)
			binary_location = nil
		}
	}
	PG__().php_binary = binary_location
}
func OnUpdateTimeout(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	if stage == PHP_INI_STAGE_STARTUP {

		/* Don't set a timeout on startup, only per-request */

		zend.ZEND_ATOL(zend.EG__().GetTimeoutSeconds(), new_value.GetVal())
		return types.SUCCESS
	}
	zend.ZendUnsetTimeout()
	zend.ZEND_ATOL(zend.EG__().GetTimeoutSeconds(), new_value.GetVal())
	if stage != PHP_INI_STAGE_DEACTIVATE {

		/*
		 * If we're restoring INI values, we shouldn't reset the timer.
		 * Otherwise, the timer is active when PHP is idle, such as the
		 * the CLI web server or CGI. Running a script will re-activate
		 * the timeout, so it's not needed to do so at script end.
		 */

		zend.ZendSetTimeout(zend.EG__().GetTimeoutSeconds(), 0)

		/*
		 * If we're restoring INI values, we shouldn't reset the timer.
		 * Otherwise, the timer is active when PHP is idle, such as the
		 * the CLI web server or CGI. Running a script will re-activate
		 * the timeout, so it's not needed to do so at script end.
		 */

	}
	return types.SUCCESS
}
func PhpGetDisplayErrorsMode(value string) int {
	var mode int
	if value == "" {
		return PHP_DISPLAY_ERRORS_STDOUT
	}
	lcValue := ascii.StrToLower(value)
	switch lcValue {
	case "on", "yes", "true", "stdout":
		mode = PHP_DISPLAY_ERRORS_STDOUT
	case "stderr":
		mode = PHP_DISPLAY_ERRORS_STDERR
	default:
		zend.ZEND_ATOL(mode, value)
		if mode != 0 && mode != PHP_DISPLAY_ERRORS_STDOUT && mode != PHP_DISPLAY_ERRORS_STDERR {
			mode = PHP_DISPLAY_ERRORS_STDOUT
		}
	}
	return mode
}
func OnUpdateDisplayErrors(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	PG__().display_errors = PhpGetDisplayErrorsMode(new_value.GetStr()) != 0
	return types.SUCCESS
}
func DisplayErrorsMode(iniEntry *zend.ZendIniEntry, typ int) {
	var mode int
	var tmpValue string
	if typ == zend.ZEND_INI_DISPLAY_ORIG && iniEntry.GetModified() != 0 {
		if iniEntry.GetOrigValue() != nil {
			tmpValue = iniEntry.GetOrigValue().GetStr()
		} else {
			tmpValue = ""
		}
	} else if iniEntry.GetValue() != nil {
		tmpValue = iniEntry.GetValue().GetStr()
	} else {
		tmpValue = ""
	}
	mode = PhpGetDisplayErrorsMode(tmpValue)

	/* Display 'On' for other SAPIs instead of STDOUT or STDERR */
	switch mode {
	case PHP_DISPLAY_ERRORS_STDERR:
		PUTS("STDERR")
	case PHP_DISPLAY_ERRORS_STDOUT:
		PUTS("STDOUT")
	default:
		PUTS("Off")
	}
}
func OnUpdateDefaultCharset(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	if memchr(new_value.GetVal(), '0', new_value.GetLen()) || strpbrk(new_value.GetVal(), "\r\n") {
		return types.FAILURE
	}
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	if new_value != nil {

	}
	return types.SUCCESS
}
func OnUpdateDefaultMimeTye(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	if memchr(new_value.GetVal(), '0', new_value.GetLen()) || strpbrk(new_value.GetVal(), "\r\n") {
		return types.FAILURE
	}
	return zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
}
func OnUpdateInternalEncoding(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	if new_value != nil {

	}
	return types.SUCCESS
}
func OnUpdateInputEncoding(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	if new_value != nil {

	}
	return types.SUCCESS
}
func OnUpdateOutputEncoding(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	if new_value != nil {

	}
	return types.SUCCESS
}
func OnUpdateErrorLog(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	/* Only do the safemode/open_basedir check at runtime */

	if (stage == PHP_INI_STAGE_RUNTIME || stage == PHP_INI_STAGE_HTACCESS) && new_value != nil && strcmp(new_value.GetVal(), "syslog") {
		if PG__().open_basedir && PhpCheckOpenBasedir(new_value.GetVal()) != 0 {
			return types.FAILURE
		}
	}
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	return types.SUCCESS
}
func OnUpdateMailLog(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	/* Only do the safemode/open_basedir check at runtime */

	if (stage == PHP_INI_STAGE_RUNTIME || stage == PHP_INI_STAGE_HTACCESS) && new_value != nil {
		if PG__().open_basedir && PhpCheckOpenBasedir(new_value.GetVal()) != 0 {
			return types.FAILURE
		}
	}
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	return types.SUCCESS
}
func OnChangeMailForceExtra(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	/* Don't allow changing it in htaccess */

	if stage == PHP_INI_STAGE_HTACCESS {
		return types.FAILURE
	}
	return types.SUCCESS
}
func PhpDuringModuleStartup() int  { return ModuleStartup }
func PhpDuringModuleShutdown() int { return ModuleShutdown }
func PhpLogErrWithSeverity(logMessage string, syslogTypeInt int) {
	if PG__().in_error_log != 0 {
		/* prevent recursive invocation */
		return
	}
	PG__().in_error_log = 1
	defer func() {
		PG__().in_error_log = 0
	}()

	/* Try to use the specified logging location. */
	if PG__().error_log != nil {
		errorLog := b.CastStrAuto(PG__().error_log)
		if errorLog == "syslog" {
			PhpSyslog(syslogTypeInt, "%s", logMessage)
			PG__().in_error_log = 0
			return
		}

		f, err := os.OpenFile(errorLog, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err == nil {
			return
		}
		defer f.Close()

		errorTimeStr := date.Format("d-M-Y H:i:s e", time.Now(), false)
		tmp := fmt.Sprintf("[%s] %s%s", errorTimeStr, logMessage, PHP_EOL)
		_, _ = f.WriteString(tmp)
		return
	}

	/* Otherwise fall back to the default logging location, if we have one */
	if SM__().GetLogMessage() != nil {
		SM__().GetLogMessage()(logMessage, syslogTypeInt)
	}
}
func PhpPrintf(format string, args ...any) int {
	return PUTS(pfmt.Sprintf(format, args))
}
func PhpVerror(docref string, params string, type_ int, format string, args ...any) {
	var docrefBuf = ""
	var target = ""
	var docref_target = ""
	var docref_root = ""
	var p *byte
	var space string
	var class_name string
	var function string
	var origin string
	var message string
	var isFunction = false

	/* get error text into buffer and escape for html if necessary */
	buffer := pfmt.Sprintf(format, args...)
	if PG__().html_errors {
		replaceBuffer := standard.PhpEscapeHtmlEntities_Ex(buffer, 0, standard.ENT_COMPAT, GetSafeCharsetHint())
		/* Retry with substituting invalid chars on fail. */
		if replaceBuffer == "" {
			replaceBuffer = standard.PhpEscapeHtmlEntities_Ex(buffer, 0, standard.ENT_COMPAT|standard.ENT_HTML_SUBSTITUTE_ERRORS, GetSafeCharsetHint())
		}
		buffer = replaceBuffer
	}

	/* which function caused the problem if any at all */

	if PhpDuringModuleStartup() != 0 {
		function = "PHP Startup"
	} else if PhpDuringModuleShutdown() != 0 {
		function = "PHP Shutdown"
	} else if zend.CurrEX() != nil && zend.CurrEX().GetFunc() != nil && zend.ZEND_USER_CODE(zend.CurrEX().GetFunc().GetType()) && zend.CurrEX().GetOpline() != nil && zend.CurrEX().GetOpline().GetOpcode() == zend.ZEND_INCLUDE_OR_EVAL {
		switch zend.CurrEX().GetOpline().GetExtendedValue() {
		case zend.ZEND_EVAL:
			function = "eval"
			isFunction = true
		case zend.ZEND_INCLUDE:
			function = "include"
			isFunction = true
		case zend.ZEND_INCLUDE_ONCE:
			function = "include_once"
			isFunction = true
		case zend.ZEND_REQUIRE:
			function = "require"
			isFunction = true
		case zend.ZEND_REQUIRE_ONCE:
			function = "require_once"
			isFunction = true
		default:
			function = "Unknown"
		}
	} else {
		function = zend.CurrEX().FunctionName()
		if function == "" {
			function = "Unknown"
		} else {
			isFunction = true
			class_name = zend.CurrEX().ClassName()
			if class_name != "" {
				space = "::"
			}
		}
	}

	/* if we still have memory then format the origin */
	if isFunction {
		origin = fmt.Sprintf("%s%s%s(%s)", class_name, space, function, params)
	} else {
		origin = function
	}
	if PG__().html_errors {
		origin = standard.PhpEscapeHtmlEntities_Ex(origin, 0, standard.ENT_COMPAT, GetSafeCharsetHint())
	}

	/* origin and buffer available, so lets come up with the error message */
	if docref != "" && docref[0] == '#' {
		docref_target = docref
		docref = ""
	}

	/* no docref given but function is known (the default) */
	if docref == "" && isFunction {
		function = strings.TrimLeft(function, "_")
		if space[0] == '0' {
			docrefBuf = fmt.Sprintf("function.%s", function)
		} else {
			docrefBuf = fmt.Sprintf("%s.%s", class_name, function)
		}
		docrefBuf = strings.ReplaceAll(docrefBuf, "_", "-")
		docref = ascii.StrToLower(docrefBuf)
	}

	/* we have a docref for a function AND
	 * - we show errors in html mode AND
	 * - the user wants to see the links
	 */

	if docref != "" && isFunction && PG__().html_errors && strlen(PG__().docref_root) {
		if strncmp(docref, "http://", 7) {

			/* We don't have 'http://' so we use docref_root */

			var ref *byte
			docref_root = PG__().docref_root
			ref = zend.Estrdup(docref)
			if docrefBuf != nil {
				zend.Efree(docrefBuf)
			}
			docrefBuf = ref

			/* strip of the target if any */

			p = strrchr(ref, '#')
			if p != nil {
				target = zend.Estrdup(p)
				if target != nil {
					docref_target = target
					*p = '0'
				}
			}

			/* add the extension if it is set in ini */

			if PG__().docref_ext && strlen(PG__().docref_ext) {
				Spprintf(&docrefBuf, 0, "%s%s", ref, PG__().docref_ext)
				zend.Efree(ref)
			}
			docref = docrefBuf
		}

		/* display html formatted or only show the additional links */

		if PG__().html_errors {
			message = fmt.Sprintf("%s [<a href='%s%s%s'>%s</a>]: %s", origin, docref_root, docref, docref_target, docref, buffer)
		} else {
			message = fmt.Sprintf("%s [%s%s%s]: %s", origin, docref_root, docref, docref_target, buffer)
		}
	} else {
		message = fmt.Sprintf("%s: %s", origin, buffer)
	}
	if PG__().track_errors && ModuleInitialized != 0 && zend.EG__().GetActive() != 0 && (zend.EG__().GetUserErrorHandler().IsUndef() || (zend.EG__().GetUserErrorHandlerErrorReporting()&type_) == 0) {
		var tmp types.Zval
		tmp.SetStringVal(buffer)
		if zend.CurrEX() != nil {
			if zend.ZendSetLocalVarStr("php_errormsg", &tmp, 0) == types.FAILURE {
				// zend.ZvalPtrDtor(&tmp)
			}
		} else {
			zend.EG__().GetSymbolTable().KeyUpdateIndirect("php_errormsg", &tmp)
		}
	}
	PhpError(type_, message)
}

func PhpErrorDocref(docRef_ *string, typ int, format string, args ...any) {
	docRef := b.Option(docRef_, "")
	PhpVerror(docRef, "", typ, format, args)
}
func PhpErrorDocref1(docRef_ *string, param1_ *string, type_ int, format string, args ...any) {
	docRef := b.Option(docRef_, "")
	params := b.Option(param1_, "")
	PhpVerror(docRef, params, type_, format, args)
}
func PhpErrorDocref2(docRef_ *string, param1_ *string, param2_ *string, type_ int, format string, args ...any) {
	docRef := b.Option(docRef_, "")
	params := pfmt.Sprintf("%s,%s", b.Option(param1_, ""), b.Option(param2_, ""))
	if params == "" {
		params = "..."
	}
	PhpVerror(docRef, params, type_, format, args)
}
func PhpHtmlPuts(str *byte, size int) { zend.ZendHtmlPuts(str, size) }
func PhpErrorCb(type_ int, error_filename string, error_lineno uint32, format string, args ...any) {
	var buffer *byte
	var buffer_len int
	var display int
	buffer_len = int(Vspprintf(&buffer, PG__().log_errors_max_len, format, args))

	/* check for repeated errors to be ignored */

	lastError := PG__().LastError()
	if PG__().ignore_repeated_errors && lastError != nil {
		/* no check for PG__().last_error_file is needed since it cannot
		 * be NULL if PG__().last_error_message is not NULL */
		if lastError.Message != buffer || (!PG__().ignore_repeated_source && lastError.Lineno != int(error_lineno)) || lastError.File != error_filename {
			display = 1
		} else {
			display = 0
		}
	} else {
		display = 1
	}

	/* according to error handling mode, throw exception or show it */

	if zend.EG__().GetErrorHandling() == zend.EH_THROW {
		switch type_ {
		case faults.E_ERROR:
			fallthrough
		case faults.E_CORE_ERROR:
			fallthrough
		case faults.E_COMPILE_ERROR:
			fallthrough
		case faults.E_USER_ERROR:
			fallthrough
		case faults.E_PARSE:

		/* fatal errors are real errors and cannot be made exceptions */

		case faults.E_STRICT:
			fallthrough
		case faults.E_DEPRECATED:
			fallthrough
		case faults.E_USER_DEPRECATED:

		/* for the sake of BC to old damaged code */

		case faults.E_NOTICE:
			fallthrough
		case faults.E_USER_NOTICE:

			/* notices are no errors and are not treated as such like E_WARNINGS */

		default:

			/* throw an exception if we are in EH_THROW mode
			 * but DO NOT overwrite a pending exception
			 */

			if zend.EG__().GetException() == nil {
				faults.ThrowErrorException(zend.EG__().GetExceptionClass(), buffer, 0, type_)
			}
			zend.Efree(buffer)
			return
		}
	}

	/* store the error if it has changed */

	if display != 0 {
		if error_filename == nil {
			error_filename = "Unknown"
		}
		PG__().AddLastError(type_, strdup(buffer), strdup(error_filename), error_lineno)
	}

	/* display/log the error if necessary */

	if display != 0 && ((zend.EG__().GetErrorReporting()&type_) != 0 || (type_&faults.E_CORE) != 0) && (PG__().log_errors || PG__().display_errors || ModuleInitialized == 0) {
		var error_type_str *byte
		var syslog_type_int int = LOG_NOTICE
		switch type_ {
		case faults.E_ERROR:
			fallthrough
		case faults.E_CORE_ERROR:
			fallthrough
		case faults.E_COMPILE_ERROR:
			fallthrough
		case faults.E_USER_ERROR:
			error_type_str = "Fatal error"
			syslog_type_int = LOG_ERR
		case faults.E_RECOVERABLE_ERROR:
			error_type_str = "Recoverable fatal error"
			syslog_type_int = LOG_ERR
		case faults.E_WARNING:
			fallthrough
		case faults.E_CORE_WARNING:
			fallthrough
		case faults.E_COMPILE_WARNING:
			fallthrough
		case faults.E_USER_WARNING:
			error_type_str = "Warning"
			syslog_type_int = LOG_WARNING
		case faults.E_PARSE:
			error_type_str = "Parse error"
			syslog_type_int = LOG_ERR
		case faults.E_NOTICE:
			fallthrough
		case faults.E_USER_NOTICE:
			error_type_str = "Notice"
			syslog_type_int = LOG_NOTICE
		case faults.E_STRICT:
			error_type_str = "Strict Standards"
			syslog_type_int = LOG_INFO
		case faults.E_DEPRECATED:
			fallthrough
		case faults.E_USER_DEPRECATED:
			error_type_str = "Deprecated"
			syslog_type_int = LOG_INFO
		default:
			error_type_str = "Unknown error"
		}
		if ModuleInitialized == 0 || PG__().log_errors {
			var log_buffer *byte
			Spprintf(&log_buffer, 0, "PHP %s:  %s in %s on line %"+"u", error_type_str, buffer, error_filename, error_lineno)
			PhpLogErrWithSeverity(log_buffer, syslog_type_int)
			zend.Efree(log_buffer)
		}
		if PG__().display_errors && (ModuleInitialized != 0 && !(PG__().during_request_startup) || PG__().display_startup_errors) {
			if PG__().xmlrpc_errors {
				PhpPrintf("<?xml version=\"1.0\"?><methodResponse><fault><value><struct><member><name>faultCode</name><value><int>"+zend.ZEND_LONG_FMT+"</int></value></member><member><name>faultString</name><value><string>%s:%s in %s on line %"+"u"+"</string></value></member></struct></value></fault></methodResponse>", PG__().xmlrpc_error_number, error_type_str, buffer, error_filename, error_lineno)
			} else {
				var prepend_string = zend.INI_STR("error_prepend_string")
				var append_string = zend.INI_STR("error_append_string")
				if PG__().html_errors {
					if type_ == faults.E_ERROR || type_ == faults.E_PARSE {
						var buf = standard.PhpEscapeHtmlEntities((*uint8)(buffer), buffer_len, 0, standard.ENT_COMPAT, GetSafeCharsetHint())
						PhpPrintf("%s<br />\n<b>%s</b>:  %s in <b>%s</b> on line <b>%"+"u"+"</b><br />\n%s", STR_PRINT(prepend_string), error_type_str, buf.GetVal(), error_filename, error_lineno, STR_PRINT(append_string))
						//types.ZendStringFree(buf)
					} else {
						PhpPrintf("%s<br />\n<b>%s</b>:  %s in <b>%s</b> on line <b>%"+"u"+"</b><br />\n%s", STR_PRINT(prepend_string), error_type_str, buffer, error_filename, error_lineno, STR_PRINT(append_string))
					}
				} else {

					/* Write CLI/CGI errors to stderr if display_errors = "stderr" */

					if PG__().display_errors == PHP_DISPLAY_ERRORS_STDERR {
						log.Printf("%s: %s in %s on line %"+"u"+"\n", error_type_str, buffer, error_filename, error_lineno)
					} else {
						PhpPrintf("%s\n%s: %s in %s on line %"+"u"+"\n%s", STR_PRINT(prepend_string), error_type_str, buffer, error_filename, error_lineno, STR_PRINT(append_string))
					}

					/* Write CLI/CGI errors to stderr if display_errors = "stderr" */

				}
			}
		}
	}

	/* Bail out if we can't recover */

	switch type_ {
	case faults.E_CORE_ERROR:
		if ModuleInitialized == 0 {

			/* bad error in module startup - no way we can live with this */

			exit(-2)

			/* bad error in module startup - no way we can live with this */

		}
		fallthrough
	case faults.E_ERROR:
		fallthrough
	case faults.E_RECOVERABLE_ERROR:
		fallthrough
	case faults.E_PARSE:
		fallthrough
	case faults.E_COMPILE_ERROR:
		fallthrough
	case faults.E_USER_ERROR:
		zend.EG__().SetExitStatus(255)
		if ModuleInitialized != 0 {
			if !PG__().display_errors && !SG__().headersSent && SG__().sapiHeaders.httpResponseCode == 200 {
				var ctr = MakeSapiHeaderLine(0)
				ctr.SetLine("HTTP/1.0 500 Internal Server Error")
				ctr.SetLineLen(b.SizeOf("\"HTTP/1.0 500 Internal Server Error\"") - 1)
				SapiHeaderOp(SAPI_HEADER_REPLACE, &ctr)
			}

			/* the parser would return 1 (failure), we can bail out nicely */

			if type_ != faults.E_PARSE {

				/* restore memory limit */

				zend.ZendSetMemoryLimit(PG__().memory_limit)
				zend.Efree(buffer)
				//zend.ZendObjectsStoreMarkDestructed(zend.EG__().GetObjectsStore())
				faults.Bailout()
				return
			}

			/* the parser would return 1 (failure), we can bail out nicely */

		}
	}

	/* Log if necessary */

	if display == 0 {
		zend.Efree(buffer)
		return
	}
	if PG__().track_errors && ModuleInitialized != 0 && zend.EG__().GetActive() != 0 {
		var tmp types.Zval
		tmp.SetStringVal(b.CastStr(buffer, buffer_len))
		if zend.CurrEX() != nil {
			if zend.ZendSetLocalVarStr("php_errormsg", &tmp, 0) == types.FAILURE {
				// zend.ZvalPtrDtor(&tmp)
			}
		} else {
			zend.EG__().GetSymbolTable().KeyUpdateIndirect("php_errormsg", &tmp)
		}
	}
	zend.Efree(buffer)
}

//@alias -old
func ZifSetTimeLimit(seconds int) bool {
	return zend.ZendAlterIniEntryChars("max_execution_time", strconv.Itoa(seconds), PHP_INI_USER, PHP_INI_STAGE_RUNTIME)
}

func PhpStreamOpenForZend(filename string) (*PhpStreamForZend, string) {
	return PhpStreamOpenForZendExEx(filename, USE_PATH|REPORT_ERRORS|STREAM_OPEN_FOR_INCLUDE)
}

func PhpStreamOpenForZendEx(filename string, mode int) *zend.FileHandle {
	stream, openedPath := PhpStreamOpenForZendExEx(filename, mode)
	if stream == nil {
		return nil
	}
	return zend.NewFileHandleByStream(filename, openedPath, stream)
}

func PhpStreamOpenForZendExEx(filename string, mode int) (*PhpStreamForZend, string) {
	var openedPath *types.String
	var stream = PhpStreamOpenWrapper(filename, "rb", mode, &openedPath)
	if stream != nil {
		/* suppress warning if this stream is not explicitly closed */
		PhpStreamAutoCleanup(stream)

		/* Disable buffering to avoid double buffering between PHP and Zend streams. */
		PhpStreamSetOption(stream, PHP_STREAM_OPTION_READ_BUFFER, PHP_STREAM_BUFFER_NONE, nil)

		return NewPhpStreamForZend(stream), openedPath.GetStr()
	}

	return nil, ""
}

func PhpResolvePathForZend(filename string) *string {
	var result string
	zstr := PhpResolvePath(filename, b.CastStrPtr(filename), len(filename), PG__().include_path)
	if zstr == nil {
		return nil
	}
	result = zstr.GetStr()
	return &result
}
func PhpFreeRequestGlobals() {
	PG__().ClearLastError()
	if PG__().php_sys_temp_dir {
		zend.Efree(PG__().php_sys_temp_dir)
		PG__().php_sys_temp_dir = nil
	}
}
func PhpMessageHandlerForZend(message zend.ZendLong, data any) {
	switch message {
	case zend.ZMSG_FAILED_INCLUDE_FOPEN:
		PhpErrorDocref("function.include", faults.E_WARNING, "Failed opening '%s' for inclusion (include_path='%s')", PhpStripUrlPasswd((*byte)(data)), STR_PRINT(PG__().include_path))
	case zend.ZMSG_FAILED_REQUIRE_FOPEN:
		PhpErrorDocref("function.require", faults.E_COMPILE_ERROR, "Failed opening required '%s' (include_path='%s')", PhpStripUrlPasswd((*byte)(data)), STR_PRINT(PG__().include_path))
	case zend.ZMSG_FAILED_HIGHLIGHT_FOPEN:
		PhpErrorDocref(nil, faults.E_WARNING, "Failed opening '%s' for highlighting", PhpStripUrlPasswd((*byte)(data)))
	case zend.ZMSG_MEMORY_LEAK_DETECTED:
		fallthrough
	case zend.ZMSG_MEMORY_LEAK_REPEATED:

	case zend.ZMSG_MEMORY_LEAKS_GRAND_TOTAL:

	case zend.ZMSG_LOG_SCRIPT_NAME:
		var ta *__struct__tm
		var tmbuf __struct__tm
		var curtime int64
		var datetime_str *byte
		var asctimebuf []*byte
		var memory_leak_buf []byte
		time(&curtime)
		ta = PhpLocaltimeR(&curtime, &tmbuf)
		datetime_str = PhpAsctimeR(ta, asctimebuf)
		if datetime_str != nil {
			datetime_str[strlen(datetime_str)-1] = 0
			Snprintf(memory_leak_buf, b.SizeOf("memory_leak_buf"), "[%s]  Script:  '%s'\n", datetime_str, SAFE_FILENAME(SG__().RequestInfo.pathTranslated))
		} else {
			Snprintf(memory_leak_buf, b.SizeOf("memory_leak_buf"), "[null]  Script:  '%s'\n", SAFE_FILENAME(SG__().RequestInfo.pathTranslated))
		}
		log.Printf("%s", memory_leak_buf)
	}
}
func PhpOnTimeout(seconds int) {
	PG__().connection_status |= PHP_CONNECTION_TIMEOUT
}
func PhpRequestStartup() int {
	retVal := faults.Try(func() {
		PG__().in_error_log = 0
		PG__().during_request_startup = 1
		PhpOutputActivate()

		/* initialize global variables */

		PG__().modules_activated = 0
		PG__().header_is_being_sent = 0
		PG__().connection_status = PHP_CONNECTION_NORMAL
		PG__().in_user_include = 0
		zend.ZendActivate()
		SapiActivate()
		zend.ZendSignalActivate()
		if PG__().max_input_time == -1 {
			zend.ZendSetTimeout(zend.EG__().GetTimeoutSeconds(), 1)
		} else {
			zend.ZendSetTimeout(PG__().max_input_time, 1)
		}

		/* Disable realpath cache if an open_basedir is set */

		if PG__().open_basedir && *PG__().open_basedir {
			zend.CWDG__().SetRealpathCacheSizeLimit(0)
		}
		if PG__().expose_php {
			SapiAddHeader(SAPI_PHP_VERSION_HEADER)
		}
		if PG__().output_handler && PG__().output_handler[0] {
			var oh types.Zval
			oh.SetStringVal(b.CastStrAuto(PG__().output_handler))
			PhpOutputStartUser(&oh, 0, PHP_OUTPUT_HANDLER_STDFLAGS)
			// zend.ZvalPtrDtor(&oh)
		} else if PG__().output_buffering {
			PhpOutputStartUser(nil, lang.CondF1(PG__().output_buffering > 1, func() __auto__ { return PG__().output_buffering }, 0), PHP_OUTPUT_HANDLER_STDFLAGS)
		} else if PG__().implicit_flush {
			PhpOutputSetImplicitFlush(1)
		}

		/* We turn this off in php_execute_script() */

		PhpHashEnvironment()
		zend.ZendActivateModules()
		PG__().modules_activated = 1
	})

	SG__().sapiStarted = 1
	return types.IntBool(retVal)
}
func PhpRequestShutdown() {
	var report_memleaks bool
	zend.EG__().AddFlags(zend.EG_FLAGS_IN_SHUTDOWN)
	report_memleaks = PG__().report_memleaks

	/* EG(current_execute_data) points into nirvana and therefore cannot be safely accessed
	 * inside zend_executor callback functions.
	 */

	zend.EG__().SetCurrentExecuteData(nil)

	/* 1. Call all possible shutdown functions registered with register_shutdown_function() */
	if PG__().modules_activated {
		faults.Try(func() {
			standard.PhpCallShutdownFunctions()
		})
	}

	/* 2. Call all possible __destruct() functions */
	faults.Try(func() {
		zend.ZendCallDestructors()
	})

	/* 3. Flush all output buffers */
	faults.Try(func() {
		var sendBuffer = lang.Cond(SG__().RequestInfo.headersOnly, 0, 1)
		if zend.CG__().GetUncleanShutdown() && PG__().LastError() != nil && PG__().LastError().Type == faults.E_ERROR && PG__().memory_limit < zend.ZendMemoryUsage(true) {
			sendBuffer = 0
		}
		if sendBuffer == 0 {
			PhpOutputDiscardAll()
		} else {
			PhpOutputEndAll()
		}
	})

	/* 4. Reset max_execution_time (no longer executing php code after response sent) */
	faults.Try(func() {
		zend.ZendUnsetTimeout()
	})

	/* 5. Call all extensions RSHUTDOWN functions */
	if PG__().modules_activated {
		zend.ZendDeactivateModules()
	}

	/* 6. Shutdown output layer (send the set HTTP headers, cleanup output handlers, etc.) */
	faults.Try(func() {
		PhpOutputDeactivate()
	})

	/* 7. Free shutdown functions */
	if PG__().modules_activated {
		standard.PhpFreeShutdownFunctions()
	}

	/* 8. Destroy super-globals */
	faults.Try(func() {
		var i int
		for i = 0; i < NUM_TRACK_VARS; i++ {
			// zend.ZvalPtrDtor(&PG__().http_globals[i])
		}
	})

	/* 9. free request-bound globals */

	PhpFreeRequestGlobals()

	/* 10. Shutdown scanner/executor/compiler and restore ini entries */

	zend.ZendDeactivate()

	/* 11. Call all extensions post-RSHUTDOWN functions */
	//faults.Try(func() {
	//zend.ZendPostDeactivateModules()
	//})

	/* 12. SAPI related shutdown (free stuff) */
	faults.Try(func() {
		SapiDeactivate()
	})

	/* 13. free virtual CWD memory */

	zend.VirtualCwdDeactivate()

	/* 14. Destroy stream hashes */
	faults.Try(func() {
		PhpShutdownStreamHashes()
	})

	/* 15. Free Willy (here be crashes) */
	faults.Try(func() {
		zend.ShutdownMemoryManager(zend.CG__().GetUncleanShutdown() || !report_memleaks, false)
	})

	/* Reset memory limit, as the reset during INI_STAGE_DEACTIVATE may have failed.
	 * At this point, no memory beyond a single chunk should be in use. */

	zend.ZendSetMemoryLimit(PG__().memory_limit)

	/* 16. Deactivate Zend signals */

	zend.ZendSignalDeactivate()

	/* 16. Deactivate Zend signals */
}
func CoreGlobalsDtor(core_globals *PhpCoreGlobals) {
	core_globals.ClearLastError()
	if core_globals.GetDisableFunctions() != nil {
		zend.Free(core_globals.GetDisableFunctions())
	}
	if core_globals.GetDisableClasses() != nil {
		zend.Free(core_globals.GetDisableClasses())
	}
	if core_globals.GetPhpBinary() != nil {
		zend.Free(core_globals.GetPhpBinary())
	}
}
func ZmInfoPhpCore(zend_module *zend.ModuleEntry) {
	standard.PhpInfoPrintTableStart()
	standard.PhpInfoPrintTableRow(2, "PHP Version", PHP_VERSION)
	standard.PhpInfoPrintTableEnd()
	zend.DISPLAY_INI_ENTRIES()
}
func PhpRegisterExtensions(ptrs []*zend.ModuleEntry) bool {
	for _, ptr := range ptrs {
		if zend.ZendRegisterInternalModule(ptr) == nil {
			return false
		}
	}
	return true
}
func PhpRegisterExtensionsBc(ptr *zend.ModuleEntry, count int) int {
	for lang.PostDec(&count) {
		if zend.ZendRegisterInternalModule(lang.PostInc(&ptr)) == nil {
			return types.FAILURE
		}
	}
	return types.SUCCESS
}

func PhpModuleStartupEx(sf ISapiModule, additional_modules []zend.ModuleEntry) bool {
	retval := PhpModuleStartup(sf, additional_modules, len(additional_modules))
	return retval != types.FAILURE
}
func PhpModuleStartup(sf ISapiModule, additional_modules *zend.ModuleEntry, num_additional_modules uint32) int {
	var zuv zend.ZendUtilityValues
	var module_number = 0
	var php_os = PHP_OS
	ModuleShutdown = 0
	ModuleStartup = 1
	SapiInitializeEmptyRequest()
	SapiActivate()
	if ModuleInitialized != 0 {
		return types.SUCCESS
	}
	SetSM__(sf)
	PhpOutputStartup()
	memset(&CoreGlobals, 0, b.SizeOf("core_globals"))

	zend.ZendStartup()
	setlocale(LC_CTYPE, "")
	tzset()

	/* Register constants */

	zend.RegisterMainStringConstant("PHP_VERSION", PHP_VERSION, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_MAJOR_VERSION", PHP_MAJOR_VERSION, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_MINOR_VERSION", PHP_MINOR_VERSION, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_RELEASE_VERSION", PHP_RELEASE_VERSION, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_EXTRA_VERSION", PHP_EXTRA_VERSION, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_VERSION_ID", PHP_VERSION_ID, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_ZTS", 0, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_DEBUG", 0, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_OS", php_os, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_OS_FAMILY", PHP_OS_FAMILY, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_SAPI", SM__().Name(), zend.CONST_PERSISTENT|zend.CONST_CS|zend.CONST_NO_FILE_CACHE)
	zend.RegisterMainStringConstant("DEFAULT_INCLUDE_PATH", PHP_INCLUDE_PATH, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PEAR_INSTALL_DIR", PEAR_INSTALLDIR, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PEAR_EXTENSION_DIR", PHP_EXTENSION_DIR, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_EXTENSION_DIR", PHP_EXTENSION_DIR, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_PREFIX", PHP_PREFIX, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_BINDIR", PHP_BINDIR, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_MANDIR", PHP_MANDIR, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_LIBDIR", PHP_LIBDIR, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_DATADIR", PHP_DATADIR, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_SYSCONFDIR", PHP_SYSCONFDIR, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_LOCALSTATEDIR", PHP_LOCALSTATEDIR, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_CONFIG_FILE_PATH", PHP_CONFIG_FILE_PATH, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_CONFIG_FILE_SCAN_DIR", PHP_CONFIG_FILE_SCAN_DIR, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_SHLIB_SUFFIX", PHP_SHLIB_SUFFIX, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainStringConstant("PHP_EOL", PHP_EOL, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_MAXPATHLEN", MAXPATHLEN, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_INT_MAX", zend.ZEND_LONG_MAX, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_INT_MIN", zend.ZEND_LONG_MIN, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_INT_SIZE", zend.SIZEOF_ZEND_LONG, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_FD_SETSIZE", FD_SETSIZE, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainLongConstant("PHP_FLOAT_DIG", DBL_DIG, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainDoubleConstant("PHP_FLOAT_EPSILON", DBL_EPSILON, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainDoubleConstant("PHP_FLOAT_MAX", DBL_MAX, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.RegisterMainDoubleConstant("PHP_FLOAT_MIN", DBL_MIN, zend.CONST_PERSISTENT|zend.CONST_CS)
	PhpBinaryInit()
	if PG__().php_binary {
		zend.RegisterMainStringConstant("PHP_BINARY", PG__().php_binary, zend.CONST_PERSISTENT|zend.CONST_CS|zend.CONST_NO_FILE_CACHE)
	} else {
		zend.RegisterMainStringConstant("PHP_BINARY", "", zend.CONST_PERSISTENT|zend.CONST_CS|zend.CONST_NO_FILE_CACHE)
	}
	PhpOutputRegisterConstants()
	PhpRfc1867RegisterConstants()

	/* this will read in php.ini, set up the configuration parameters,
	   load zend extensions and register php function extensions
	   to be loaded later */

	if PhpInitConfig() == types.FAILURE {
		return types.FAILURE
	}

	/* Register PHP core ini entries */

	zend.REGISTER_INI_ENTRIES(module_number)

	/* Register Zend ini entries */

	zend.ZendRegisterStandardIniEntries()

	/* Disable realpath cache if an open_basedir is set */

	if PG__().open_basedir && *PG__().open_basedir {
		zend.CWDG__().SetRealpathCacheSizeLimit(0)
	}
	PG__().have_called_openlog = 0

	/* initialize stream wrappers registry
	 * (this uses configuration parameters from php.ini)
	 */

	if PhpInitStreamWrappers(module_number) == types.FAILURE {
		PhpPrintf("PHP:  Unable to initialize stream url wrappers.\n")
		return types.FAILURE
	}
	zuv.SetHtmlErrors(1)
	PhpStartupAutoGlobals()
	zend.ZendSetUtilityValues(&zuv)
	PhpStartupSapiContentTypes()

	/* startup extensions statically compiled in */

	if PhpRegisterInternalExtensions() == false {
		PhpPrintf("Unable to start builtin modules\n")
		return types.FAILURE
	}

	/* start additional PHP extensions */

	PhpRegisterExtensionsBc(additional_modules, num_additional_modules)

	/* load and startup extensions compiled as shared objects (aka DLLs)
	   as requested by php.ini entries
	   these are loaded after initialization of internal extensions
	   as extensions *might* rely on things from ext/standard
	   which is always an internal extension and to be initialized
	   ahead of all other internals
	*/

	PhpIniRegisterExtensions()
	zend.ZendStartupModules()

	/* start Zend extensions */
	zend.ZendExtensions.Startup()
	zend.ZendCollectModuleHandlers()

	/* register additional functions */

	if SM__().GetAdditionalFunctions() != nil {
		if module := globals.G().GetModule("standard"); module != nil {
			zend.EG__().SetCurrentModule(module)
			zend.ZendRegisterFunctions(nil, SM__().GetAdditionalFunctions(), nil, zend.MODULE_PERSISTENT)
			zend.EG__().SetCurrentModule(nil)
		}
	}

	/* disable certain classes and functions as requested by php.ini */

	PhpDisableFunctions()
	PhpDisableClasses()

	/* make core report what it should */
	if module := globals.G().GetModule("core"); module != nil {
		module.SetInfoFunc(ZmInfoPhpCore)
	}
	ModuleInitialized = 1
	if zend.ZendPostStartup() != types.SUCCESS {
		return types.FAILURE
	}

	/* Check for deprecated directives */
	var directives = []struct {
		error_level int
		phrase      string
		directives  []string
	}{
		{
			faults.E_DEPRECATED,
			"Directive '%s' is deprecated",
			[]string{"track_errors", "allow_url_include"},
		},
		{
			faults.E_CORE_ERROR,
			"Directive '%s' is no longer available in PHP",
			[]string{"allow_call_time_pass_reference", "asp_tags", "define_syslog_variables", "highlight.bg", "magic_quotes_gpc", "magic_quotes_runtime", "magic_quotes_sybase", "register_globals", "register_long_arrays", "safe_mode", "safe_mode_gid", "safe_mode_include_dir", "safe_mode_exec_dir", "safe_mode_allowed_env_vars", "safe_mode_protected_env_vars", "zend.ze1_compatibility_mode"},
		},
	}

	retval := faults.Try(func() {
		for _, directive := range directives {
			for _, p := range directive.directives {
				var value zend.ZendLong
				if CfgGetLong(p, &value) == types.SUCCESS && value != 0 {
					faults.Error(directive.error_level, directive.phrase, p)
				}
			}
		}
	})

	zend.VirtualCwdDeactivate()
	SapiDeactivate()
	ModuleStartup = 0
	zend.ShutdownMemoryManager(true, false)
	zend.VirtualCwdActivate()

	/* we're done */
	return types.IntBool(retval)
}
func PhpModuleShutdown() {
	var module_number = 0
	ModuleShutdown = 1
	if ModuleInitialized == 0 {
		return
	}
	SapiFlush()
	zend.ZendShutdown()

	/* Destroys filter & transport registries too */

	PhpShutdownStreamWrappers(module_number)
	zend.UNREGISTER_INI_ENTRIES(module_number)

	/* close down the ini config */

	PhpShutdownConfig()
	zend.ZendIniShutdown()
	zend.ShutdownMemoryManager(zend.CG__().GetUncleanShutdown(), true)
	PhpOutputShutdown()
	ModuleInitialized = 0
	CoreGlobalsDtor(&CoreGlobals)
	//zend.GcGlobalsDtor()
}
func PhpExecuteScript(primaryFile *zend.FileHandle) bool {
	var prepend_file_p *zend.FileHandle
	var append_file_p *zend.FileHandle
	var prepend_file zend.FileHandle
	var append_file zend.FileHandle
	var old_cwd *byte
	var retval = false
	zend.EG__().SetExitStatus(0)
	const OLD_CWD_SIZE = 4096
	old_cwd = zend.DoAlloca(OLD_CWD_SIZE, use_heap)
	old_cwd[0] = '0'

	faults.Try(func() {
		var realfile []byte
		PG__().during_request_startup = 0
		if primaryFile.GetFilename() != nil && (SG__().options&SAPI_OPTION_NO_CHDIR) == 0 {
			PhpIgnoreValue(zend.VCWD_GETCWD(old_cwd, OLD_CWD_SIZE-1))
			zend.VCWD_CHDIR_FILE(primaryFile.GetFilename())
		}

		/* Only lookup the real file path and add it to the included_files list if already opened
		 *   otherwise it will get opened and added to the included_files list in zend_execute_scripts
		 */

		if primaryFile.GetFilename() != nil && strcmp("Standard input code", primaryFile.GetFilename()) && primaryFile.GetOpenedPath() == nil && !primaryFile.IsTypeHandleFileName() {
			if ExpandFilepath(primaryFile.GetFilename(), realfile) != nil {
				primaryFile.SetOpenedPath(realfile)
				types.ZendHashAddEmptyElement(zend.EG__().GetIncludedFiles(), primaryFile.GetOpenedPath().GetStr())
			}
		}
		if PG__().auto_prepend_file && PG__().auto_prepend_file[0] {
			prepend_file_p = zend.NewFileHandleByFilename(PG__().auto_prepend_file)
		} else {
			prepend_file_p = nil
		}
		if PG__().auto_append_file && PG__().auto_append_file[0] {
			append_file_p = zend.NewFileHandleByFilename(PG__().auto_append_file)
		} else {
			append_file_p = nil
		}
		if PG__().max_input_time != -1 {
			zend.ZendSetTimeout(zend.INI_INT("max_execution_time"), 0)
		}

		/*
		   If cli primary file has shabang line and there is a prepend file,
		   the `skip_shebang` will be used by prepend file but not primary file,
		   save it and restore after prepend file been executed.
		*/

		if zend.CG__().GetSkipShebang() != 0 && prepend_file_p != nil {
			zend.CG__().SetSkipShebang(0)
			if zend.ZendExecuteScriptsEx(zend.ZEND_REQUIRE, nil, prepend_file_p) {
				zend.CG__().SetSkipShebang(1)
				retval = zend.ZendExecuteScriptsEx(zend.ZEND_REQUIRE, nil, primaryFile, append_file_p)
			}
		} else {
			retval = zend.ZendExecuteScriptsEx(zend.ZEND_REQUIRE, nil, prepend_file_p, primaryFile, append_file_p)
		}
	})

	if zend.EG__().GetException() != nil {
		faults.Try(func() {
			faults.ExceptionError(zend.EG__().GetException(), faults.E_ERROR)
		})
	}
	if old_cwd[0] != '0' {
		PhpIgnoreValue(zend.VCWD_CHDIR(old_cwd))
	}
	zend.FreeAlloca(old_cwd, use_heap)
	return retval
}
func PhpHandleAbortedConnection() {
	PG__().connection_status = PHP_CONNECTION_ABORTED
	OG__().SetStatusDisabled()
	if !(PG__().ignore_user_abort) {
		faults.Bailout()
	}
}
func PhpHandleAuthData(auth *byte) int {
	var ret = -1
	var auth_len int = lang.CondF1(auth != nil, func() __auto__ { return strlen(auth) }, 0)
	if auth != nil && auth_len > 0 && operators.ZendBinaryStrncasecmp(b.CastStr(auth, auth_len), "Basic ", b.SizeOf("\"Basic \"")-1) == 0 {
		var pass *byte
		var user *types.String
		user = types.NewString(standard.PhpBase64Decode(b.CastStr((*uint8)(auth+6), auth_len-6)))
		if user != nil {
			pass = strchr(user.GetVal(), ':')
			if pass != nil {
				lang.PostInc(&(*pass)) = '0'
				SG__().RequestInfo.authUser = zend.Estrndup(user.GetVal(), user.GetLen())
				SG__().RequestInfo.authPassword = zend.Estrdup(pass)
				ret = 0
			}
			//types.ZendStringFree(user)
		}
	}
	if ret == -1 {
		SG__().RequestInfo.authPassword = nil
		SG__().RequestInfo.authUser = SG__().RequestInfo.authPassword
	} else {
		SG__().RequestInfo.authDigest = nil
	}
	if ret == -1 && auth != nil && auth_len > 0 && operators.ZendBinaryStrncasecmp(b.CastStr(auth, auth_len), "Digest ", b.SizeOf("\"Digest \"")-1) == 0 {
		SG__().RequestInfo.authDigest = zend.Estrdup(auth + 7)
		ret = 0
	}
	if ret == -1 {
		SG__().RequestInfo.authDigest = nil
	}
	return ret
}
func PhpLintScript(file *zend.FileHandle) int {
	var op_array *types.ZendOpArray
	var retval = types.FAILURE

	faults.Try(func() {
		op_array = zend.CompileFile(file, zend.ZEND_INCLUDE)
		zend.ZendDestroyFileHandle(file)
		if op_array != nil {
			//zend.DestroyOpArray(op_array)
			//zend.Efree(op_array)
			retval = types.SUCCESS
		}
	})
	if zend.EG__().GetException() != nil {
		faults.ExceptionError(zend.EG__().GetException(), faults.E_ERROR)
	}
	return retval
}
