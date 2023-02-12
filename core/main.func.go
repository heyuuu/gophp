// <<generate>>

package core

import (
	b "sik/builtin"
	"sik/core/streams"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/zend"
)

func SAFE_FILENAME(f __auto__) string {
	if f {
		return f
	} else {
		return "-"
	}
}
func GetSafeCharsetHint() *byte {
	var lastHint *byte = nil
	var lastCodeset *byte = nil
	var hint *byte = SG__().default_charset
	var len_ int = strlen(hint)
	var i int = 0
	if lastHint == SG__().default_charset {
		return lastCodeset
	}
	lastHint = hint
	lastCodeset = nil
	for i = 0; i < b.SizeOf("charset_map")/b.SizeOf("charset_map [ 0 ]"); i++ {
		if len_ == standard.CharsetMap[i].codeset_len && zend.ZendBinaryStrcasecmp(hint, len_, standard.CharsetMap[i].codeset, len_) == 0 {
			lastCodeset = (*byte)(standard.CharsetMap[i].codeset)
			break
		}
	}
	return lastCodeset
}
func OnSetFacility(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var facility *byte = new_value.GetVal()
	return zend.FAILURE
}
func OnSetPrecision(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var i zend.ZendLong
	zend.ZEND_ATOL(i, new_value.GetVal())
	if i >= -1 {
		zend.EG__().SetPrecision(i)
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func OnSetSerializePrecision(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var i zend.ZendLong
	zend.ZEND_ATOL(i, new_value.GetVal())
	if i >= -1 {
		PG(serialize_precision) = i
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func OnChangeMemoryLimit(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var value int
	if new_value != nil {
		value = zend.ZendAtol(new_value.GetVal(), new_value.GetLen())
	} else {
		value = int64(1) << 30
	}
	if zend.ZendSetMemoryLimit(value) == zend.FAILURE {

		/* When the memory limit is reset to the original level during deactivation, we may be
		 * using more memory than the original limit while shutdown is still in progress.
		 * Ignore a failure for now, and set the memory limit when the memory manager has been
		 * shut down and the minimal amount of memory is used. */

		if stage != zend.ZEND_INI_STAGE_DEACTIVATE {
			zend.ZendError(zend.E_WARNING, "Failed to set memory limit to %zd bytes (Current memory usage is %zd bytes)", value, zend.ZendMemoryUsage(true))
			return zend.FAILURE
		}

		/* When the memory limit is reset to the original level during deactivation, we may be
		 * using more memory than the original limit while shutdown is still in progress.
		 * Ignore a failure for now, and set the memory limit when the memory manager has been
		 * shut down and the minimal amount of memory is used. */

	}
	PG(memory_limit) = value
	return zend.SUCCESS
}
func OnSetLogFilter(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var filter *byte = new_value.GetVal()
	if !(strcmp(filter, "all")) {
		PG(syslog_filter) = PHP_SYSLOG_FILTER_ALL
		return zend.SUCCESS
	}
	if !(strcmp(filter, "no-ctrl")) {
		PG(syslog_filter) = PHP_SYSLOG_FILTER_NO_CTRL
		return zend.SUCCESS
	}
	if !(strcmp(filter, "ascii")) {
		PG(syslog_filter) = PHP_SYSLOG_FILTER_ASCII
		return zend.SUCCESS
	}
	if !(strcmp(filter, "raw")) {
		PG(syslog_filter) = PHP_SYSLOG_FILTER_RAW
		return zend.SUCCESS
	}
	return zend.FAILURE
}
func PhpDisableFunctions() {
	var s *byte = nil
	var e *byte
	if !(*(zend.INI_STR("disable_functions"))) {
		return
	}
	PG(disable_functions) = strdup(zend.INI_STR("disable_functions"))
	e = PG(disable_functions)
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
				zend.ZendDisableFunction(s, e-s)
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
		zend.ZendDisableFunction(s, e-s)
	}
}
func PhpDisableClasses() {
	var s *byte = nil
	var e *byte
	if !(*(zend.INI_STR("disable_classes"))) {
		return
	}
	PG(disable_classes) = strdup(zend.INI_STR("disable_classes"))
	e = PG(disable_classes)
	for *e {
		switch *e {
		case ' ':
			fallthrough
		case ',':
			if s != nil {
				*e = '0'
				zend.ZendDisableClass(s, e-s)
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
		zend.ZendDisableClass(s, e-s)
	}
}
func PhpBinaryInit() {
	var binary_location *byte = nil
	if sapi_module.GetExecutableLocation() != nil {
		binary_location = (*byte)(zend.Malloc(MAXPATHLEN))
		if binary_location != nil && !(strchr(sapi_module.GetExecutableLocation(), '/')) {
			var envpath *byte
			var path *byte
			var found int = 0
			if b.Assign(&envpath, getenv("PATH")) != nil {
				var search_dir *byte
				var search_path []*byte
				var last *byte = nil
				var s zend.ZendStatT
				path = zend.Estrdup(envpath)
				search_dir = PhpStrtokR(path, ":", &last)
				for search_dir != nil {
					Snprintf(search_path, MAXPATHLEN, "%s/%s", search_dir, sapi_module.GetExecutableLocation())
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
		} else if zend.VCWD_REALPATH(sapi_module.GetExecutableLocation(), binary_location) == nil || zend.VCWD_ACCESS(binary_location, X_OK) {
			zend.Free(binary_location)
			binary_location = nil
		}
	}
	PG(php_binary) = binary_location
}
func OnUpdateTimeout(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	if stage == PHP_INI_STAGE_STARTUP {

		/* Don't set a timeout on startup, only per-request */

		zend.ZEND_ATOL(zend.EG__().GetTimeoutSeconds(), new_value.GetVal())
		return zend.SUCCESS
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
	return zend.SUCCESS
}
func PhpGetDisplayErrorsMode(value *byte, value_length int) int {
	var mode int
	if value == nil {
		return PHP_DISPLAY_ERRORS_STDOUT
	}
	if value_length == 2 && !(strcasecmp("on", value)) {
		mode = PHP_DISPLAY_ERRORS_STDOUT
	} else if value_length == 3 && !(strcasecmp("yes", value)) {
		mode = PHP_DISPLAY_ERRORS_STDOUT
	} else if value_length == 4 && !(strcasecmp("true", value)) {
		mode = PHP_DISPLAY_ERRORS_STDOUT
	} else if value_length == 6 && !(strcasecmp(value, "stderr")) {
		mode = PHP_DISPLAY_ERRORS_STDERR
	} else if value_length == 6 && !(strcasecmp(value, "stdout")) {
		mode = PHP_DISPLAY_ERRORS_STDOUT
	} else {
		zend.ZEND_ATOL(mode, value)
		if mode != 0 && mode != PHP_DISPLAY_ERRORS_STDOUT && mode != PHP_DISPLAY_ERRORS_STDERR {
			mode = PHP_DISPLAY_ERRORS_STDOUT
		}
	}
	return mode
}
func OnUpdateDisplayErrors(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	PG(display_errors) = zend.ZendBool(PhpGetDisplayErrorsMode(new_value.GetVal(), new_value.GetLen()))
	return zend.SUCCESS
}
func DisplayErrorsMode(ini_entry *zend.ZendIniEntry, type_ int) {
	var mode int
	var cgi_or_cli int
	var tmp_value_length int
	var tmp_value *byte
	if type_ == zend.ZEND_INI_DISPLAY_ORIG && ini_entry.GetModified() != 0 {
		if ini_entry.GetOrigValue() != nil {
			tmp_value = ini_entry.GetOrigValue().GetVal()
		} else {
			tmp_value = nil
		}
		if ini_entry.GetOrigValue() != nil {
			tmp_value_length = ini_entry.GetOrigValue().GetLen()
		} else {
			tmp_value_length = 0
		}
	} else if ini_entry.GetValue() != nil {
		tmp_value = ini_entry.GetValue().GetVal()
		tmp_value_length = ini_entry.GetValue().GetLen()
	} else {
		tmp_value = nil
		tmp_value_length = 0
	}
	mode = PhpGetDisplayErrorsMode(tmp_value, tmp_value_length)

	/* Display 'On' for other SAPIs instead of STDOUT or STDERR */

	cgi_or_cli = !(strcmp(sapi_module.GetName(), "cli")) || !(strcmp(sapi_module.GetName(), "cgi")) || !(strcmp(sapi_module.GetName(), "phpdbg"))
	switch mode {
	case PHP_DISPLAY_ERRORS_STDERR:
		if cgi_or_cli != 0 {
			PUTS("STDERR")
		} else {
			PUTS("On")
		}
	case PHP_DISPLAY_ERRORS_STDOUT:
		if cgi_or_cli != 0 {
			PUTS("STDOUT")
		} else {
			PUTS("On")
		}
	default:
		PUTS("Off")
	}
}
func PhpGetInternalEncoding() *byte {
	if PG(internal_encoding) && PG(internal_encoding)[0] {
		return PG(internal_encoding)
	} else if SG__().default_charset {
		return SG__().default_charset
	}
	return ""
}
func PhpGetInputEncoding() *byte {
	if PG(input_encoding) && PG(input_encoding)[0] {
		return PG(input_encoding)
	} else if SG__().default_charset {
		return SG__().default_charset
	}
	return ""
}
func PhpGetOutputEncoding() *byte {
	if PG(output_encoding) && PG(output_encoding)[0] {
		return PG(output_encoding)
	} else if SG__().default_charset {
		return SG__().default_charset
	}
	return ""
}
func OnUpdateDefaultCharset(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	if memchr(new_value.GetVal(), '0', new_value.GetLen()) || strpbrk(new_value.GetVal(), "\r\n") {
		return zend.FAILURE
	}
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	if PhpInternalEncodingChanged != nil {
		PhpInternalEncodingChanged()
	}
	if new_value != nil {

	}
	return zend.SUCCESS
}
func OnUpdateDefaultMimeTye(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	if memchr(new_value.GetVal(), '0', new_value.GetLen()) || strpbrk(new_value.GetVal(), "\r\n") {
		return zend.FAILURE
	}
	return zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
}
func OnUpdateInternalEncoding(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
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
	return zend.SUCCESS
}
func OnUpdateInputEncoding(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
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
	return zend.SUCCESS
}
func OnUpdateOutputEncoding(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
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
	return zend.SUCCESS
}
func OnUpdateErrorLog(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	/* Only do the safemode/open_basedir check at runtime */

	if (stage == PHP_INI_STAGE_RUNTIME || stage == PHP_INI_STAGE_HTACCESS) && new_value != nil && strcmp(new_value.GetVal(), "syslog") {
		if PG(open_basedir) && PhpCheckOpenBasedir(new_value.GetVal()) != 0 {
			return zend.FAILURE
		}
	}
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	return zend.SUCCESS
}
func OnUpdateMailLog(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	/* Only do the safemode/open_basedir check at runtime */

	if (stage == PHP_INI_STAGE_RUNTIME || stage == PHP_INI_STAGE_HTACCESS) && new_value != nil {
		if PG(open_basedir) && PhpCheckOpenBasedir(new_value.GetVal()) != 0 {
			return zend.FAILURE
		}
	}
	zend.OnUpdateString(entry, new_value, mh_arg1, mh_arg2, mh_arg3, stage)
	return zend.SUCCESS
}
func OnChangeMailForceExtra(
	entry *zend.ZendIniEntry,
	new_value *zend.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	/* Don't allow changing it in htaccess */

	if stage == PHP_INI_STAGE_HTACCESS {
		return zend.FAILURE
	}
	return zend.SUCCESS
}
func PhpDuringModuleStartup() int  { return ModuleStartup }
func PhpDuringModuleShutdown() int { return ModuleShutdown }
func PhpGetModuleInitialized() int { return ModuleInitialized }
func PhpLogErrWithSeverity(log_message *byte, syslog_type_int int) {
	var fd int = -1
	var error_time int64
	if PG(in_error_log) {

		/* prevent recursive invocation */

		return

		/* prevent recursive invocation */

	}
	PG(in_error_log) = 1

	/* Try to use the specified logging location. */

	if PG(error_log) != nil {
		if !(strcmp(PG(error_log), "syslog")) {
			PhpSyslog(syslog_type_int, "%s", log_message)
			PG(in_error_log) = 0
			return
		}
		fd = zend.VCWD_OPEN_MODE(PG(error_log), O_CREAT|O_APPEND|O_WRONLY, 0644)
		if fd != -1 {
			var tmp *byte
			var len_ int
			var error_time_str *zend.ZendString
			time(&error_time)
			error_time_str = php_format_date("d-M-Y H:i:s e", 13, error_time, 1)
			len_ = Spprintf(&tmp, 0, "[%s] %s%s", error_time_str.GetVal(), log_message, PHP_EOL)
			PhpIgnoreValue(write(fd, tmp, len_))
			zend.Efree(tmp)
			zend.ZendStringFree(error_time_str)
			close(fd)
			PG(in_error_log) = 0
			return
		}
	}

	/* Otherwise fall back to the default logging location, if we have one */

	if sapi_module.GetLogMessage() != nil {
		sapi_module.GetLogMessage()(log_message, syslog_type_int)
	}
	PG(in_error_log) = 0
}
func PhpWrite(buf any, size int) int { return PHPWRITE(buf, size) }
func PhpPrintf(format string, _ ...any) int {
	var args va_list
	var ret int
	var buffer *byte
	var size int
	va_start(args, format)
	size = Vspprintf(&buffer, 0, format, args)
	ret = PHPWRITE(buffer, size)
	zend.Efree(buffer)
	va_end(args)
	return ret
}
func PhpVerror(docref *byte, params *byte, type_ int, format *byte, args ...any) {
	var replace_buffer *zend.ZendString = nil
	var replace_origin *zend.ZendString = nil
	var buffer *byte = nil
	var docref_buf *byte = nil
	var target *byte = nil
	var docref_target *byte = ""
	var docref_root *byte = ""
	var p *byte
	var buffer_len int = 0
	var space *byte = ""
	var class_name *byte = ""
	var function *byte
	var origin_len int
	var origin *byte
	var message *byte
	var is_function int = 0

	/* get error text into buffer and escape for html if necessary */

	buffer_len = int(Vspprintf(&buffer, 0, format, args))
	if PG(html_errors) {
		replace_buffer = standard.PhpEscapeHtmlEntities((*uint8)(buffer), buffer_len, 0, standard.ENT_COMPAT, GetSafeCharsetHint())

		/* Retry with substituting invalid chars on fail. */

		if replace_buffer == nil || replace_buffer.GetLen() < 1 {
			replace_buffer = standard.PhpEscapeHtmlEntities((*uint8)(buffer), buffer_len, 0, standard.ENT_COMPAT|standard.ENT_HTML_SUBSTITUTE_ERRORS, GetSafeCharsetHint())
		}
		zend.Efree(buffer)
		if replace_buffer != nil {
			buffer = replace_buffer.GetVal()
			buffer_len = int(replace_buffer.GetLen())
		} else {
			buffer = ""
			buffer_len = 0
		}
	}

	/* which function caused the problem if any at all */

	if PhpDuringModuleStartup() != 0 {
		function = "PHP Startup"
	} else if PhpDuringModuleShutdown() != 0 {
		function = "PHP Shutdown"
	} else if zend.EG__().GetCurrentExecuteData() != nil && zend.EG__().GetCurrentExecuteData().GetFunc() != nil && zend.ZEND_USER_CODE(zend.EG__().GetCurrentExecuteData().GetFunc().GetCommonType()) && zend.EG__().GetCurrentExecuteData().GetOpline() != nil && zend.EG__().GetCurrentExecuteData().GetOpline().GetOpcode() == zend.ZEND_INCLUDE_OR_EVAL {
		switch zend.EG__().GetCurrentExecuteData().GetOpline().GetExtendedValue() {
		case zend.ZEND_EVAL:
			function = "eval"
			is_function = 1
		case zend.ZEND_INCLUDE:
			function = "include"
			is_function = 1
		case zend.ZEND_INCLUDE_ONCE:
			function = "include_once"
			is_function = 1
		case zend.ZEND_REQUIRE:
			function = "require"
			is_function = 1
		case zend.ZEND_REQUIRE_ONCE:
			function = "require_once"
			is_function = 1
		default:
			function = "Unknown"
		}
	} else {
		function = zend.GetActiveFunctionName()
		if function == nil || !(strlen(function)) {
			function = "Unknown"
		} else {
			is_function = 1
			class_name = zend.GetActiveClassName(&space)
		}
	}

	/* if we still have memory then format the origin */

	if is_function != 0 {
		origin_len = int(Spprintf(&origin, 0, "%s%s%s(%s)", class_name, space, function, params))
	} else {
		origin_len = int(Spprintf(&origin, 0, "%s", function))
	}
	if PG(html_errors) {
		replace_origin = standard.PhpEscapeHtmlEntities((*uint8)(origin), origin_len, 0, standard.ENT_COMPAT, GetSafeCharsetHint())
		zend.Efree(origin)
		origin = replace_origin.GetVal()
	}

	/* origin and buffer available, so lets come up with the error message */

	if docref != nil && docref[0] == '#' {
		docref_target = strchr(docref, '#')
		docref = nil
	}

	/* no docref given but function is known (the default) */

	if docref == nil && is_function != 0 {
		var doclen int
		for (*function) == '_' {
			function++
		}
		if space[0] == '0' {
			doclen = int(Spprintf(&docref_buf, 0, "function.%s", function))
		} else {
			doclen = int(Spprintf(&docref_buf, 0, "%s.%s", class_name, function))
		}
		for b.Assign(&p, strchr(docref_buf, '_')) != nil {
			*p = '-'
		}
		docref = standard.PhpStrtolower(docref_buf, doclen)
	}

	/* we have a docref for a function AND
	 * - we show errors in html mode AND
	 * - the user wants to see the links
	 */

	if docref != nil && is_function != 0 && PG(html_errors) && strlen(PG(docref_root)) {
		if strncmp(docref, "http://", 7) {

			/* We don't have 'http://' so we use docref_root */

			var ref *byte
			docref_root = PG(docref_root)
			ref = zend.Estrdup(docref)
			if docref_buf != nil {
				zend.Efree(docref_buf)
			}
			docref_buf = ref

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

			if PG(docref_ext) && strlen(PG(docref_ext)) {
				Spprintf(&docref_buf, 0, "%s%s", ref, PG(docref_ext))
				zend.Efree(ref)
			}
			docref = docref_buf
		}

		/* display html formatted or only show the additional links */

		if PG(html_errors) {
			Spprintf(&message, 0, "%s [<a href='%s%s%s'>%s</a>]: %s", origin, docref_root, docref, docref_target, docref, buffer)
		} else {
			Spprintf(&message, 0, "%s [%s%s%s]: %s", origin, docref_root, docref, docref_target, buffer)
		}
		if target != nil {
			zend.Efree(target)
		}
	} else {
		Spprintf(&message, 0, "%s: %s", origin, buffer)
	}
	if replace_origin != nil {
		zend.ZendStringFree(replace_origin)
	} else {
		zend.Efree(origin)
	}
	if docref_buf != nil {
		zend.Efree(docref_buf)
	}
	if PG(track_errors) && ModuleInitialized != 0 && zend.EG__().GetActive() != 0 && (zend.EG__().GetUserErrorHandler().IsType(zend.IS_UNDEF) || (zend.EG__().GetUserErrorHandlerErrorReporting()&type_) == 0) {
		var tmp zend.Zval
		zend.ZVAL_STRINGL(&tmp, buffer, buffer_len)
		if zend.EG__().GetCurrentExecuteData() != nil {
			if zend.ZendSetLocalVarStr("php_errormsg", b.SizeOf("\"php_errormsg\"")-1, &tmp, 0) == zend.FAILURE {
				zend.ZvalPtrDtor(&tmp)
			}
		} else {
			zend.EG__().GetSymbolTable().KeyUpdateIndirect("php_errormsg", &tmp)
		}
	}
	if replace_buffer != nil {
		zend.ZendStringFree(replace_buffer)
	} else {
		zend.Efree(buffer)
	}
	PhpError(type_, "%s", message)
	zend.Efree(message)
}
func PhpErrorDocref(docref string, type_ int, format string, _ ...any) {
	var args va_list
	va_start(args, format)
	PhpVerror(docref, "", type_, format, args)
	va_end(args)
}
func PhpErrorDocref1(docref *byte, param1 *byte, type_ int, format string, _ ...any) {
	var args va_list
	va_start(args, format)
	PhpVerror(docref, param1, type_, format, args)
	va_end(args)
}
func PhpErrorDocref2(
	docref *byte,
	param1 *byte,
	param2 *byte,
	type_ int,
	format string,
	_ ...any,
) {
	var params *byte
	var args va_list
	Spprintf(&params, 0, "%s,%s", param1, param2)
	va_start(args, format)
	PhpVerror(docref, b.Cond(params != nil, params, "..."), type_, format, args)
	va_end(args)
	if params != nil {
		zend.Efree(params)
	}
}
func PhpHtmlPuts(str *byte, size int) { zend.ZendHtmlPuts(str, size) }
func PhpErrorCb(type_ int, error_filename *byte, error_lineno uint32, format *byte, args ...any) {
	var buffer *byte
	var buffer_len int
	var display int
	buffer_len = int(Vspprintf(&buffer, PG(log_errors_max_len), format, args))

	/* check for repeated errors to be ignored */

	if PG(ignore_repeated_errors) && PG(last_error_message) {

		/* no check for PG(last_error_file) is needed since it cannot
		 * be NULL if PG(last_error_message) is not NULL */

		if strcmp(PG(last_error_message), buffer) || !(PG(ignore_repeated_source)) && (PG(last_error_lineno) != int(error_lineno) || strcmp(PG(last_error_file), error_filename)) {
			display = 1
		} else {
			display = 0
		}

		/* no check for PG(last_error_file) is needed since it cannot
		 * be NULL if PG(last_error_message) is not NULL */

	} else {
		display = 1
	}

	/* according to error handling mode, throw exception or show it */

	if zend.EG__().GetErrorHandling() == zend.EH_THROW {
		switch type_ {
		case zend.E_ERROR:
			fallthrough
		case zend.E_CORE_ERROR:
			fallthrough
		case zend.E_COMPILE_ERROR:
			fallthrough
		case zend.E_USER_ERROR:
			fallthrough
		case zend.E_PARSE:

		/* fatal errors are real errors and cannot be made exceptions */

		case zend.E_STRICT:
			fallthrough
		case zend.E_DEPRECATED:
			fallthrough
		case zend.E_USER_DEPRECATED:

		/* for the sake of BC to old damaged code */

		case zend.E_NOTICE:
			fallthrough
		case zend.E_USER_NOTICE:

			/* notices are no errors and are not treated as such like E_WARNINGS */

		default:

			/* throw an exception if we are in EH_THROW mode
			 * but DO NOT overwrite a pending exception
			 */

			if zend.EG__().GetException() == nil {
				zend.ZendThrowErrorException(zend.EG__().GetExceptionClass(), buffer, 0, type_)
			}
			zend.Efree(buffer)
			return
		}
	}

	/* store the error if it has changed */

	if display != 0 {
		if PG(last_error_message) {
			var s *byte = PG(last_error_message)
			PG(last_error_message) = nil
			zend.Free(s)
		}
		if PG(last_error_file) {
			var s *byte = PG(last_error_file)
			PG(last_error_file) = nil
			zend.Free(s)
		}
		if error_filename == nil {
			error_filename = "Unknown"
		}
		PG(last_error_type) = type_
		PG(last_error_message) = strdup(buffer)
		PG(last_error_file) = strdup(error_filename)
		PG(last_error_lineno) = error_lineno
	}

	/* display/log the error if necessary */

	if display != 0 && ((zend.EG__().GetErrorReporting()&type_) != 0 || (type_&zend.E_CORE) != 0) && (PG(log_errors) || PG(display_errors) || ModuleInitialized == 0) {
		var error_type_str *byte
		var syslog_type_int int = LOG_NOTICE
		switch type_ {
		case zend.E_ERROR:
			fallthrough
		case zend.E_CORE_ERROR:
			fallthrough
		case zend.E_COMPILE_ERROR:
			fallthrough
		case zend.E_USER_ERROR:
			error_type_str = "Fatal error"
			syslog_type_int = LOG_ERR
		case zend.E_RECOVERABLE_ERROR:
			error_type_str = "Recoverable fatal error"
			syslog_type_int = LOG_ERR
		case zend.E_WARNING:
			fallthrough
		case zend.E_CORE_WARNING:
			fallthrough
		case zend.E_COMPILE_WARNING:
			fallthrough
		case zend.E_USER_WARNING:
			error_type_str = "Warning"
			syslog_type_int = LOG_WARNING
		case zend.E_PARSE:
			error_type_str = "Parse error"
			syslog_type_int = LOG_ERR
		case zend.E_NOTICE:
			fallthrough
		case zend.E_USER_NOTICE:
			error_type_str = "Notice"
			syslog_type_int = LOG_NOTICE
		case zend.E_STRICT:
			error_type_str = "Strict Standards"
			syslog_type_int = LOG_INFO
		case zend.E_DEPRECATED:
			fallthrough
		case zend.E_USER_DEPRECATED:
			error_type_str = "Deprecated"
			syslog_type_int = LOG_INFO
		default:
			error_type_str = "Unknown error"
		}
		if ModuleInitialized == 0 || PG(log_errors) {
			var log_buffer *byte
			Spprintf(&log_buffer, 0, "PHP %s:  %s in %s on line %"+"u", error_type_str, buffer, error_filename, error_lineno)
			PhpLogErrWithSeverity(log_buffer, syslog_type_int)
			zend.Efree(log_buffer)
		}
		if PG(display_errors) && (ModuleInitialized != 0 && !(PG(during_request_startup)) || PG(display_startup_errors)) {
			if PG(xmlrpc_errors) {
				PhpPrintf("<?xml version=\"1.0\"?><methodResponse><fault><value><struct><member><name>faultCode</name><value><int>"+zend.ZEND_LONG_FMT+"</int></value></member><member><name>faultString</name><value><string>%s:%s in %s on line %"+"u"+"</string></value></member></struct></value></fault></methodResponse>", PG(xmlrpc_error_number), error_type_str, buffer, error_filename, error_lineno)
			} else {
				var prepend_string *byte = zend.INI_STR("error_prepend_string")
				var append_string *byte = zend.INI_STR("error_append_string")
				if PG(html_errors) {
					if type_ == zend.E_ERROR || type_ == zend.E_PARSE {
						var buf *zend.ZendString = standard.PhpEscapeHtmlEntities((*uint8)(buffer), buffer_len, 0, standard.ENT_COMPAT, GetSafeCharsetHint())
						PhpPrintf("%s<br />\n<b>%s</b>:  %s in <b>%s</b> on line <b>%"+"u"+"</b><br />\n%s", STR_PRINT(prepend_string), error_type_str, buf.GetVal(), error_filename, error_lineno, STR_PRINT(append_string))
						zend.ZendStringFree(buf)
					} else {
						PhpPrintf("%s<br />\n<b>%s</b>:  %s in <b>%s</b> on line <b>%"+"u"+"</b><br />\n%s", STR_PRINT(prepend_string), error_type_str, buffer, error_filename, error_lineno, STR_PRINT(append_string))
					}
				} else {

					/* Write CLI/CGI errors to stderr if display_errors = "stderr" */

					if (!(strcmp(sapi_module.GetName(), "cli")) || !(strcmp(sapi_module.GetName(), "cgi")) || !(strcmp(sapi_module.GetName(), "phpdbg"))) && PG(display_errors) == PHP_DISPLAY_ERRORS_STDERR {
						r.Fprintf(stderr, "%s: %s in %s on line %"+"u"+"\n", error_type_str, buffer, error_filename, error_lineno)
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
	case zend.E_CORE_ERROR:
		if ModuleInitialized == 0 {

			/* bad error in module startup - no way we can live with this */

			exit(-2)

			/* bad error in module startup - no way we can live with this */

		}
		fallthrough
	case zend.E_ERROR:
		fallthrough
	case zend.E_RECOVERABLE_ERROR:
		fallthrough
	case zend.E_PARSE:
		fallthrough
	case zend.E_COMPILE_ERROR:
		fallthrough
	case zend.E_USER_ERROR:
		zend.EG__().SetExitStatus(255)
		if ModuleInitialized != 0 {
			if !(PG(display_errors)) && !(SG__().headers_sent) && SG__().sapi_headers.http_response_code == 200 {
				var ctr SapiHeaderLine = MakeSapiHeaderLine(0)
				ctr.SetLine("HTTP/1.0 500 Internal Server Error")
				ctr.SetLineLen(b.SizeOf("\"HTTP/1.0 500 Internal Server Error\"") - 1)
				SapiHeaderOp(SAPI_HEADER_REPLACE, &ctr)
			}

			/* the parser would return 1 (failure), we can bail out nicely */

			if type_ != zend.E_PARSE {

				/* restore memory limit */

				zend.ZendSetMemoryLimit(PG(memory_limit))
				zend.Efree(buffer)
				zend.ZendObjectsStoreMarkDestructed(zend.EG__().GetObjectsStore())
				zend.ZendBailout()
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
	if PG(track_errors) && ModuleInitialized != 0 && zend.EG__().GetActive() != 0 {
		var tmp zend.Zval
		zend.ZVAL_STRINGL(&tmp, buffer, buffer_len)
		if zend.EG__().GetCurrentExecuteData() != nil {
			if zend.ZendSetLocalVarStr("php_errormsg", b.SizeOf("\"php_errormsg\"")-1, &tmp, 0) == zend.FAILURE {
				zend.ZvalPtrDtor(&tmp)
			}
		} else {
			zend.EG__().GetSymbolTable().KeyUpdateIndirect("php_errormsg", &tmp)
		}
	}
	zend.Efree(buffer)
}
func PhpGetCurrentUser() *byte {
	var pstat *zend.ZendStatT
	if SG__().request_info.current_user {
		return SG__().request_info.current_user
	}

	/* FIXME: I need to have this somehow handled if
	   USE_SAPI is defined, because cgi will also be
	   interfaced in USE_SAPI */

	pstat = SapiGetStat()
	if pstat == nil {
		return ""
	} else {
		var pwd *__struct__passwd
		if b.Assign(&pwd, getpwuid(pstat.st_uid)) == nil {
			return ""
		}
		SG__().request_info.current_user_length = strlen(pwd.pw_name)
		SG__().request_info.current_user = zend.Estrndup(pwd.pw_name, SG__().request_info.current_user_length)
		return SG__().request_info.current_user
	}
}
func ZifSetTimeLimit(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var new_timeout zend.ZendLong
	var new_timeout_str *byte
	var new_timeout_strlen int
	var key *zend.ZendString
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &new_timeout) == zend.FAILURE {
		return
	}
	new_timeout_strlen = int(zend.ZendSpprintf(&new_timeout_str, 0, zend.ZEND_LONG_FMT, new_timeout))
	key = zend.ZendStringInit("max_execution_time", b.SizeOf("\"max_execution_time\"")-1, 0)
	if zend.ZendAlterIniEntryCharsEx(key, new_timeout_str, new_timeout_strlen, PHP_INI_USER, PHP_INI_STAGE_RUNTIME, 0) == zend.SUCCESS {
		return_value.SetTrue()
	} else {
		return_value.SetFalse()
	}
	zend.ZendStringReleaseEx(key, 0)
	zend.Efree(new_timeout_str)
}
func PhpFopenWrapperForZend(filename *byte, opened_path **zend.ZendString) *r.FILE {
	return streams.PhpStreamOpenWrapperAsFile((*byte)(filename), "rb", USE_PATH|IGNORE_URL_WIN|REPORT_ERRORS|STREAM_OPEN_FOR_INCLUDE, opened_path)
}
func PhpZendStreamCloser(handle any) { PhpStreamClose((*PhpStream)(handle)) }
func PhpZendStreamFsizer(handle any) int {
	var stream *PhpStream = handle
	var ssb PhpStreamStatbuf

	/* File size reported by stat() may be inaccurate if stream filters are used.
	 * TODO: Should stat() be generally disabled if filters are used? */

	if stream.GetReadfilters().GetHead() != nil {
		return 0
	}
	if PhpStreamStat(stream, &ssb) == 0 {
		return ssb.GetSb().st_size
	}
	return 0
}
func PhpStreamOpenForZend(filename *byte, handle *zend.ZendFileHandle) int {
	return PhpStreamOpenForZendEx(filename, handle, USE_PATH|REPORT_ERRORS|STREAM_OPEN_FOR_INCLUDE)
}
func PhpStreamOpenForZendEx(filename *byte, handle *zend.ZendFileHandle, mode int) int {
	var opened_path *zend.ZendString
	var stream *PhpStream = PhpStreamOpenWrapper((*byte)(filename), "rb", mode, &opened_path)
	if stream != nil {
		memset(handle, 0, b.SizeOf("zend_file_handle"))
		handle.SetType(zend.ZEND_HANDLE_STREAM)
		handle.SetFilename((*byte)(filename))
		handle.SetOpenedPath(opened_path)
		handle.SetStream(zend.MakeZendStream(
			stream,
			0,
			zend.ZendStreamReaderT(_phpStreamRead),
			PhpZendStreamFsizer,
			PhpZendStreamCloser,
		))

		/* suppress warning if this stream is not explicitly closed */

		PhpStreamAutoCleanup(stream)

		/* Disable buffering to avoid double buffering between PHP and Zend streams. */

		PhpStreamSetOption(stream, PHP_STREAM_OPTION_READ_BUFFER, PHP_STREAM_BUFFER_NONE, nil)
		return zend.SUCCESS
	}
	return zend.FAILURE
}
func PhpResolvePathForZend(filename *byte, filename_len int) *zend.ZendString {
	return PhpResolvePath(filename, filename_len, PG(include_path))
}
func PhpGetConfigurationDirectiveForZend(name *zend.ZendString) *zend.Zval {
	return CfgGetEntryEx(name)
}
func PhpFreeRequestGlobals() {
	if PG(last_error_message) {
		zend.Free(PG(last_error_message))
		PG(last_error_message) = nil
	}
	if PG(last_error_file) {
		zend.Free(PG(last_error_file))
		PG(last_error_file) = nil
	}
	if PG(php_sys_temp_dir) {
		zend.Efree(PG(php_sys_temp_dir))
		PG(php_sys_temp_dir) = nil
	}
}
func PhpMessageHandlerForZend(message zend.ZendLong, data any) {
	switch message {
	case zend.ZMSG_FAILED_INCLUDE_FOPEN:
		PhpErrorDocref("function.include", zend.E_WARNING, "Failed opening '%s' for inclusion (include_path='%s')", PhpStripUrlPasswd((*byte)(data)), STR_PRINT(PG(include_path)))
	case zend.ZMSG_FAILED_REQUIRE_FOPEN:
		PhpErrorDocref("function.require", zend.E_COMPILE_ERROR, "Failed opening required '%s' (include_path='%s')", PhpStripUrlPasswd((*byte)(data)), STR_PRINT(PG(include_path)))
	case zend.ZMSG_FAILED_HIGHLIGHT_FOPEN:
		PhpErrorDocref(nil, zend.E_WARNING, "Failed opening '%s' for highlighting", PhpStripUrlPasswd((*byte)(data)))
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
			Snprintf(memory_leak_buf, b.SizeOf("memory_leak_buf"), "[%s]  Script:  '%s'\n", datetime_str, SAFE_FILENAME(SG__().request_info.path_translated))
		} else {
			Snprintf(memory_leak_buf, b.SizeOf("memory_leak_buf"), "[null]  Script:  '%s'\n", SAFE_FILENAME(SG__().request_info.path_translated))
		}
		r.Fprintf(stderr, "%s", memory_leak_buf)
	}
}
func PhpOnTimeout(seconds int) {
	PG(connection_status) |= PHP_CONNECTION_TIMEOUT
}
func PhpRequestStartup() int {
	var retval int = zend.SUCCESS
	zend.ZendInternedStringsActivate()
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		PG(in_error_log) = 0
		PG(during_request_startup) = 1
		PhpOutputActivate()

		/* initialize global variables */

		PG(modules_activated) = 0
		PG(header_is_being_sent) = 0
		PG(connection_status) = PHP_CONNECTION_NORMAL
		PG(in_user_include) = 0
		zend.ZendActivate()
		SapiActivate()
		zend.ZendSignalActivate()
		if PG(max_input_time) == -1 {
			zend.ZendSetTimeout(zend.EG__().GetTimeoutSeconds(), 1)
		} else {
			zend.ZendSetTimeout(PG(max_input_time), 1)
		}

		/* Disable realpath cache if an open_basedir is set */

		if PG(open_basedir) && (*PG)(open_basedir) {
			zend.CWDG(realpath_cache_size_limit) = 0
		}
		if PG(expose_php) {
			SapiAddHeader(SAPI_PHP_VERSION_HEADER)
		}
		if PG(output_handler) && PG(output_handler)[0] {
			var oh zend.Zval
			zend.ZVAL_STRING(&oh, PG(output_handler))
			PhpOutputStartUser(&oh, 0, PHP_OUTPUT_HANDLER_STDFLAGS)
			zend.ZvalPtrDtor(&oh)
		} else if PG(output_buffering) {
			PhpOutputStartUser(nil, b.CondF1(PG(output_buffering) > 1, func() __auto__ { return PG(output_buffering) }, 0), PHP_OUTPUT_HANDLER_STDFLAGS)
		} else if PG(implicit_flush) {
			PhpOutputSetImplicitFlush(1)
		}

		/* We turn this off in php_execute_script() */

		PhpHashEnvironment()
		zend.ZendActivateModules()
		PG(modules_activated) = 1
	} else {
		zend.EG__().SetBailout(__orig_bailout)
		retval = zend.FAILURE
	}
	zend.EG__().SetBailout(__orig_bailout)
	SG__().sapi_started = 1
	return retval
}
func PhpRequestShutdown(dummy any) {
	var report_memleaks zend.ZendBool
	zend.EG__().AddFlags(zend.EG_FLAGS_IN_SHUTDOWN)
	report_memleaks = PG(report_memleaks)

	/* EG(current_execute_data) points into nirvana and therefore cannot be safely accessed
	 * inside zend_executor callback functions.
	 */

	zend.EG__().SetCurrentExecuteData(nil)
	PhpDeactivateTicks()

	/* 1. Call all possible shutdown functions registered with register_shutdown_function() */

	if PG(modules_activated) {
		var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
		var __bailout JMP_BUF
		zend.EG__().SetBailout(&__bailout)
		if zend.SETJMP(__bailout) == 0 {
			standard.PhpCallShutdownFunctions()
		}
		zend.EG__().SetBailout(__orig_bailout)
	}

	/* 2. Call all possible __destruct() functions */

	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		zend.ZendCallDestructors()
	}
	zend.EG__().SetBailout(__orig_bailout)

	/* 3. Flush all output buffers */

	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		var send_buffer zend.ZendBool = b.Cond(SG__().request_info.headers_only, 0, 1)
		if zend.CG__().GetUncleanShutdown() != 0 && PG(last_error_type) == zend.E_ERROR && int(PG(memory_limit) < zend.ZendMemoryUsage(1)) != 0 {
			send_buffer = 0
		}
		if send_buffer == 0 {
			PhpOutputDiscardAll()
		} else {
			PhpOutputEndAll()
		}
	}
	zend.EG__().SetBailout(__orig_bailout)

	/* 4. Reset max_execution_time (no longer executing php code after response sent) */

	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		zend.ZendUnsetTimeout()
	}
	zend.EG__().SetBailout(__orig_bailout)

	/* 5. Call all extensions RSHUTDOWN functions */

	if PG(modules_activated) {
		zend.ZendDeactivateModules()
	}

	/* 6. Shutdown output layer (send the set HTTP headers, cleanup output handlers, etc.) */

	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		PhpOutputDeactivate()
	}
	zend.EG__().SetBailout(__orig_bailout)

	/* 7. Free shutdown functions */

	if PG(modules_activated) {
		standard.PhpFreeShutdownFunctions()
	}

	/* 8. Destroy super-globals */

	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		var i int
		for i = 0; i < NUM_TRACK_VARS; i++ {
			zend.ZvalPtrDtor(&PG(http_globals)[i])
		}
	}
	zend.EG__().SetBailout(__orig_bailout)

	/* 9. free request-bound globals */

	PhpFreeRequestGlobals()

	/* 10. Shutdown scanner/executor/compiler and restore ini entries */

	zend.ZendDeactivate()

	/* 11. Call all extensions post-RSHUTDOWN functions */

	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		zend.ZendPostDeactivateModules()
	}
	zend.EG__().SetBailout(__orig_bailout)

	/* 12. SAPI related shutdown (free stuff) */

	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		SapiDeactivate()
	}
	zend.EG__().SetBailout(__orig_bailout)

	/* 13. free virtual CWD memory */

	zend.VirtualCwdDeactivate()

	/* 14. Destroy stream hashes */

	var __orig_bailout *JMP_BUF = EG__().bailout
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		PhpShutdownStreamHashes()
	}
	zend.EG__().SetBailout(__orig_bailout)

	/* 15. Free Willy (here be crashes) */

	zend.ZendInternedStringsDeactivate()
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		zend.ShutdownMemoryManager(zend.CG__().GetUncleanShutdown() != 0 || report_memleaks == 0, 0)
	}
	zend.EG__().SetBailout(__orig_bailout)

	/* Reset memory limit, as the reset during INI_STAGE_DEACTIVATE may have failed.
	 * At this point, no memory beyond a single chunk should be in use. */

	zend.ZendSetMemoryLimit(PG(memory_limit))

	/* 16. Deactivate Zend signals */

	zend.ZendSignalDeactivate()

	/* 16. Deactivate Zend signals */
}
func PhpComInitialize() {}
func CoreGlobalsDtor(core_globals *PhpCoreGlobals) {
	if core_globals.GetLastErrorMessage() != nil {
		zend.Free(core_globals.GetLastErrorMessage())
	}
	if core_globals.GetLastErrorFile() != nil {
		zend.Free(core_globals.GetLastErrorFile())
	}
	if core_globals.GetDisableFunctions() != nil {
		zend.Free(core_globals.GetDisableFunctions())
	}
	if core_globals.GetDisableClasses() != nil {
		zend.Free(core_globals.GetDisableClasses())
	}
	if core_globals.GetPhpBinary() != nil {
		zend.Free(core_globals.GetPhpBinary())
	}
	PhpShutdownTicks()
}
func ZmInfoPhpCore(zend_module *zend.ZendModuleEntry) {
	standard.PhpInfoPrintTableStart()
	standard.PhpInfoPrintTableRow(2, "PHP Version", PHP_VERSION)
	standard.PhpInfoPrintTableEnd()
	zend.DISPLAY_INI_ENTRIES()
}
func PhpRegisterExtensions(ptr **zend.ZendModuleEntry, count int) int {
	var end **zend.ZendModuleEntry = ptr + count
	for ptr < end {
		if (*ptr) != nil {
			if zend.ZendRegisterInternalModule(*ptr) == nil {
				return zend.FAILURE
			}
		}
		ptr++
	}
	return zend.SUCCESS
}
func PhpRegisterExtensionsBc(ptr *zend.ZendModuleEntry, count int) int {
	for b.PostDec(&count) {
		if zend.ZendRegisterInternalModule(b.PostInc(&ptr)) == nil {
			return zend.FAILURE
		}
	}
	return zend.SUCCESS
}
func PhpModuleStartup(sf *SapiModule, additional_modules *zend.ZendModuleEntry, num_additional_modules uint32) int {
	var zuf zend.ZendUtilityFunctions
	var zuv zend.ZendUtilityValues
	var retval int = zend.SUCCESS
	var module_number int = 0
	var php_os *byte
	var module *zend.ZendModuleEntry
	php_os = PHP_OS
	ModuleShutdown = 0
	ModuleStartup = 1
	SapiInitializeEmptyRequest()
	SapiActivate()
	if ModuleInitialized != 0 {
		return zend.SUCCESS
	}
	sapi_module = *sf
	PhpOutputStartup()
	memset(&CoreGlobals, 0, b.SizeOf("core_globals"))
	PhpStartupTicks()
	//zend.GcGlobalsCtor()
	zuf.SetErrorFunction(PhpErrorCb)
	zuf.SetPrintfFunction(PhpPrintf)
	zuf.SetWriteFunction(PhpOutputWrite)
	zuf.SetFopenFunction(PhpFopenWrapperForZend)
	zuf.SetMessageHandler(PhpMessageHandlerForZend)
	zuf.SetGetConfigurationDirective(PhpGetConfigurationDirectiveForZend)
	zuf.SetTicksFunction(PhpRunTicks)
	zuf.SetOnTimeout(PhpOnTimeout)
	zuf.SetStreamOpenFunction(PhpStreamOpenForZend)
	zuf.SetPrintfToSmartStringFunction(PhpPrintfToSmartString)
	zuf.SetPrintfToSmartStrFunction(PhpPrintfToSmartStr)
	zuf.SetGetenvFunction(SapiGetenv)
	zuf.SetResolvePathFunction(PhpResolvePathForZend)
	zend.ZendStartup(&zuf)
	setlocale(LC_CTYPE, "")
	tzset()
	zend.LeIndexPtr = zend.ZendRegisterListDestructorsEx(nil, nil, "index pointer", 0)

	/* Register constants */

	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_VERSION", PHP_VERSION, b.SizeOf("PHP_VERSION")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_MAJOR_VERSION", PHP_MAJOR_VERSION, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_MINOR_VERSION", PHP_MINOR_VERSION, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_RELEASE_VERSION", PHP_RELEASE_VERSION, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_EXTRA_VERSION", PHP_EXTRA_VERSION, b.SizeOf("PHP_EXTRA_VERSION")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_VERSION_ID", PHP_VERSION_ID, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_ZTS", 0, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_DEBUG", 0, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_OS", php_os, strlen(php_os), zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_OS_FAMILY", PHP_OS_FAMILY, b.SizeOf("PHP_OS_FAMILY")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_SAPI", sapi_module.GetName(), strlen(sapi_module.GetName()), zend.CONST_PERSISTENT|zend.CONST_CS|zend.CONST_NO_FILE_CACHE)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("DEFAULT_INCLUDE_PATH", PHP_INCLUDE_PATH, b.SizeOf("PHP_INCLUDE_PATH")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PEAR_INSTALL_DIR", PEAR_INSTALLDIR, b.SizeOf("PEAR_INSTALLDIR")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PEAR_EXTENSION_DIR", PHP_EXTENSION_DIR, b.SizeOf("PHP_EXTENSION_DIR")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_EXTENSION_DIR", PHP_EXTENSION_DIR, b.SizeOf("PHP_EXTENSION_DIR")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_PREFIX", PHP_PREFIX, b.SizeOf("PHP_PREFIX")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_BINDIR", PHP_BINDIR, b.SizeOf("PHP_BINDIR")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_MANDIR", PHP_MANDIR, b.SizeOf("PHP_MANDIR")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_LIBDIR", PHP_LIBDIR, b.SizeOf("PHP_LIBDIR")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_DATADIR", PHP_DATADIR, b.SizeOf("PHP_DATADIR")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_SYSCONFDIR", PHP_SYSCONFDIR, b.SizeOf("PHP_SYSCONFDIR")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_LOCALSTATEDIR", PHP_LOCALSTATEDIR, b.SizeOf("PHP_LOCALSTATEDIR")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_CONFIG_FILE_PATH", PHP_CONFIG_FILE_PATH, strlen(PHP_CONFIG_FILE_PATH), zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_CONFIG_FILE_SCAN_DIR", PHP_CONFIG_FILE_SCAN_DIR, b.SizeOf("PHP_CONFIG_FILE_SCAN_DIR")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_SHLIB_SUFFIX", PHP_SHLIB_SUFFIX, b.SizeOf("PHP_SHLIB_SUFFIX")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_EOL", PHP_EOL, b.SizeOf("PHP_EOL")-1, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_MAXPATHLEN", MAXPATHLEN, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_INT_MAX", zend.ZEND_LONG_MAX, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_INT_MIN", zend.ZEND_LONG_MIN, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_INT_SIZE", zend.SIZEOF_ZEND_LONG, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_FD_SETSIZE", FD_SETSIZE, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_LONG_CONSTANT("PHP_FLOAT_DIG", DBL_DIG, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_DOUBLE_CONSTANT("PHP_FLOAT_EPSILON", DBL_EPSILON, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_DOUBLE_CONSTANT("PHP_FLOAT_MAX", DBL_MAX, zend.CONST_PERSISTENT|zend.CONST_CS)
	zend.REGISTER_MAIN_DOUBLE_CONSTANT("PHP_FLOAT_MIN", DBL_MIN, zend.CONST_PERSISTENT|zend.CONST_CS)
	PhpBinaryInit()
	if PG(php_binary) {
		zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_BINARY", PG(php_binary), strlen(PG(php_binary)), zend.CONST_PERSISTENT|zend.CONST_CS|zend.CONST_NO_FILE_CACHE)
	} else {
		zend.REGISTER_MAIN_STRINGL_CONSTANT("PHP_BINARY", "", 0, zend.CONST_PERSISTENT|zend.CONST_CS|zend.CONST_NO_FILE_CACHE)
	}
	PhpOutputRegisterConstants()
	PhpRfc1867RegisterConstants()

	/* this will read in php.ini, set up the configuration parameters,
	   load zend extensions and register php function extensions
	   to be loaded later */

	if PhpInitConfig() == zend.FAILURE {
		return zend.FAILURE
	}

	/* Register PHP core ini entries */

	zend.REGISTER_INI_ENTRIES()

	/* Register Zend ini entries */

	zend.ZendRegisterStandardIniEntries()

	/* Disable realpath cache if an open_basedir is set */

	if PG(open_basedir) && (*PG)(open_basedir) {
		zend.CWDG(realpath_cache_size_limit) = 0
	}
	PG(have_called_openlog) = 0

	/* initialize stream wrappers registry
	 * (this uses configuration parameters from php.ini)
	 */

	if PhpInitStreamWrappers(module_number) == zend.FAILURE {
		PhpPrintf("PHP:  Unable to initialize stream url wrappers.\n")
		return zend.FAILURE
	}
	zuv.SetHtmlErrors(1)
	PhpStartupAutoGlobals()
	zend.ZendSetUtilityValues(&zuv)
	PhpStartupSapiContentTypes()

	/* startup extensions statically compiled in */

	if PhpRegisterInternalExtensionsFunc() == zend.FAILURE {
		PhpPrintf("Unable to start builtin modules\n")
		return zend.FAILURE
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

	zend.ZendStartupExtensions()
	zend.ZendCollectModuleHandlers()

	/* register additional functions */

	if sapi_module.GetAdditionalFunctions() != nil {
		if b.Assign(&module, zend.ZendHashStrFindPtr(&zend.ModuleRegistry, "standard", b.SizeOf("\"standard\"")-1)) != nil {
			zend.EG__().SetCurrentModule(module)
			zend.ZendRegisterFunctions(nil, sapi_module.GetAdditionalFunctions(), nil, zend.MODULE_PERSISTENT)
			zend.EG__().SetCurrentModule(nil)
		}
	}

	/* disable certain classes and functions as requested by php.ini */

	PhpDisableFunctions()
	PhpDisableClasses()

	/* make core report what it should */

	if b.Assign(&module, zend.ZendHashStrFindPtr(&zend.ModuleRegistry, "core", b.SizeOf("\"core\"")-1)) != nil {
		module.SetVersion(PHP_VERSION)
		module.SetInfoFunc(ZmInfoPhpCore)
	}
	ModuleInitialized = 1
	if zend.ZendPostStartup() != zend.SUCCESS {
		return zend.FAILURE
	}

	/* Check for deprecated directives */

	var directives []struct {
		error_level long
		phrase      *byte
		directives  []*byte
	} = []struct {
		error_level long
		phrase      *byte
		directives  []*byte
	}{
		{
			zend.E_DEPRECATED,
			"Directive '%s' is deprecated",
			{"track_errors", "allow_url_include", nil},
		},
		{
			zend.E_CORE_ERROR,
			"Directive '%s' is no longer available in PHP",
			{"allow_call_time_pass_reference", "asp_tags", "define_syslog_variables", "highlight.bg", "magic_quotes_gpc", "magic_quotes_runtime", "magic_quotes_sybase", "register_globals", "register_long_arrays", "safe_mode", "safe_mode_gid", "safe_mode_include_dir", "safe_mode_exec_dir", "safe_mode_allowed_env_vars", "safe_mode_protected_env_vars", "zend.ze1_compatibility_mode", nil},
		},
	}
	var i uint
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {

		/* 2 = Count of deprecation structs */

		for i = 0; i < 2; i++ {
			var p **byte = directives[i].directives
			for (*p) != nil {
				var value zend.ZendLong
				if CfgGetLong((*byte)(*p), &value) == zend.SUCCESS && value != 0 {
					zend.ZendError(directives[i].error_level, directives[i].phrase, *p)
				}
				p++
			}
		}

		/* 2 = Count of deprecation structs */

	} else {
		zend.EG__().SetBailout(__orig_bailout)
		retval = zend.FAILURE
	}
	zend.EG__().SetBailout(__orig_bailout)
	zend.VirtualCwdDeactivate()
	SapiDeactivate()
	ModuleStartup = 0
	zend.ShutdownMemoryManager(1, 0)
	zend.VirtualCwdActivate()
	zend.ZendInternedStringsSwitchStorage(1)

	/* we're done */

	return retval

	/* we're done */
}
func PhpModuleShutdownWrapper(sapi_globals *SapiModule) int {
	PhpModuleShutdown()
	return zend.SUCCESS
}
func PhpModuleShutdown() {
	var module_number int = 0
	ModuleShutdown = 1
	if ModuleInitialized == 0 {
		return
	}
	zend.ZendInternedStringsSwitchStorage(0)
	SapiFlush()
	zend.ZendShutdown()

	/* Destroys filter & transport registries too */

	PhpShutdownStreamWrappers(module_number)
	zend.UNREGISTER_INI_ENTRIES()

	/* close down the ini config */

	PhpShutdownConfig()
	zend.ZendIniShutdown()
	zend.ShutdownMemoryManager(zend.CG__().GetUncleanShutdown(), 1)
	PhpOutputShutdown()
	zend.ZendInternedStringsDtor()
	if zend.ZendPostShutdownCb != nil {
		var cb func() = zend.ZendPostShutdownCb
		zend.ZendPostShutdownCb = nil
		cb()
	}
	ModuleInitialized = 0
	CoreGlobalsDtor(&CoreGlobals)
	//zend.GcGlobalsDtor()
}
func PhpExecuteScript(primary_file *zend.ZendFileHandle) int {
	var prepend_file_p *zend.ZendFileHandle
	var append_file_p *zend.ZendFileHandle
	var prepend_file zend.ZendFileHandle
	var append_file zend.ZendFileHandle
	var old_cwd *byte
	var retval int = 0
	zend.EG__().SetExitStatus(0)
	const OLD_CWD_SIZE = 4096
	old_cwd = zend.DoAlloca(OLD_CWD_SIZE, use_heap)
	old_cwd[0] = '0'
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		var realfile []byte
		PG(during_request_startup) = 0
		if primary_file.GetFilename() != nil && (SG__().options&SAPI_OPTION_NO_CHDIR) == 0 {
			PhpIgnoreValue(zend.VCWD_GETCWD(old_cwd, OLD_CWD_SIZE-1))
			zend.VCWD_CHDIR_FILE(primary_file.GetFilename())
		}

		/* Only lookup the real file path and add it to the included_files list if already opened
		 *   otherwise it will get opened and added to the included_files list in zend_execute_scripts
		 */

		if primary_file.GetFilename() != nil && strcmp("Standard input code", primary_file.GetFilename()) && primary_file.GetOpenedPath() == nil && primary_file.GetType() != zend.ZEND_HANDLE_FILENAME {
			if ExpandFilepath(primary_file.GetFilename(), realfile) != nil {
				primary_file.SetOpenedPath(zend.ZendStringInit(realfile, strlen(realfile), 0))
				zend.ZendHashAddEmptyElement(zend.EG__().GetIncludedFiles(), primary_file.GetOpenedPath())
			}
		}
		if PG(auto_prepend_file) && PG(auto_prepend_file)[0] {
			prepend_file.InitFilename(PG(auto_prepend_file))
			prepend_file_p = &prepend_file
		} else {
			prepend_file_p = nil
		}
		if PG(auto_append_file) && PG(auto_append_file)[0] {
			append_file.InitFilename(PG(auto_append_file))
			append_file_p = &append_file
		} else {
			append_file_p = nil
		}
		if PG(max_input_time) != -1 {
			zend.ZendSetTimeout(zend.INI_INT("max_execution_time"), 0)
		}

		/*
		   If cli primary file has shabang line and there is a prepend file,
		   the `skip_shebang` will be used by prepend file but not primary file,
		   save it and restore after prepend file been executed.
		*/

		if zend.CG__().GetSkipShebang() != 0 && prepend_file_p != nil {
			zend.CG__().SetSkipShebang(0)
			if zend.ZendExecuteScripts(zend.ZEND_REQUIRE, nil, 1, prepend_file_p) == zend.SUCCESS {
				zend.CG__().SetSkipShebang(1)
				retval = zend.ZendExecuteScripts(zend.ZEND_REQUIRE, nil, 2, primary_file, append_file_p) == zend.SUCCESS
			}
		} else {
			retval = zend.ZendExecuteScripts(zend.ZEND_REQUIRE, nil, 3, prepend_file_p, primary_file, append_file_p) == zend.SUCCESS
		}

		/*
		   If cli primary file has shabang line and there is a prepend file,
		   the `skip_shebang` will be used by prepend file but not primary file,
		   save it and restore after prepend file been executed.
		*/

	}
	zend.EG__().SetBailout(__orig_bailout)
	if zend.EG__().GetException() != nil {
		var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
		var __bailout JMP_BUF
		zend.EG__().SetBailout(&__bailout)
		if zend.SETJMP(__bailout) == 0 {
			zend.ZendExceptionError(zend.EG__().GetException(), zend.E_ERROR)
		}
		zend.EG__().SetBailout(__orig_bailout)
	}
	if old_cwd[0] != '0' {
		PhpIgnoreValue(zend.VCWD_CHDIR(old_cwd))
	}
	zend.FreeAlloca(old_cwd, use_heap)
	return retval
}
func PhpExecuteSimpleScript(primary_file *zend.ZendFileHandle, ret *zend.Zval) int {
	var old_cwd *byte
	zend.EG__().SetExitStatus(0)
	const OLD_CWD_SIZE = 4096
	old_cwd = zend.DoAlloca(OLD_CWD_SIZE, use_heap)
	old_cwd[0] = '0'
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		PG(during_request_startup) = 0
		if primary_file.GetFilename() != nil && (SG__().options&SAPI_OPTION_NO_CHDIR) == 0 {
			PhpIgnoreValue(zend.VCWD_GETCWD(old_cwd, OLD_CWD_SIZE-1))
			zend.VCWD_CHDIR_FILE(primary_file.GetFilename())
		}
		zend.ZendExecuteScripts(zend.ZEND_REQUIRE, ret, 1, primary_file)
	}
	zend.EG__().SetBailout(__orig_bailout)
	if old_cwd[0] != '0' {
		PhpIgnoreValue(zend.VCWD_CHDIR(old_cwd))
	}
	zend.FreeAlloca(old_cwd, use_heap)
	return zend.EG__().GetExitStatus()
}
func PhpHandleAbortedConnection() {
	PG(connection_status) = PHP_CONNECTION_ABORTED
	PhpOutputSetStatus(PHP_OUTPUT_DISABLED)
	if !(PG(ignore_user_abort)) {
		zend.ZendBailout()
	}
}
func PhpHandleAuthData(auth *byte) int {
	var ret int = -1
	var auth_len int = b.CondF1(auth != nil, func() __auto__ { return strlen(auth) }, 0)
	if auth != nil && auth_len > 0 && zend.ZendBinaryStrncasecmp(auth, auth_len, "Basic ", b.SizeOf("\"Basic \"")-1, b.SizeOf("\"Basic \"")-1) == 0 {
		var pass *byte
		var user *zend.ZendString
		user = standard.PhpBase64Decode((*uint8)(auth+6), auth_len-6)
		if user != nil {
			pass = strchr(user.GetVal(), ':')
			if pass != nil {
				b.PostInc(&(*pass)) = '0'
				SG__().request_info.auth_user = zend.Estrndup(user.GetVal(), user.GetLen())
				SG__().request_info.auth_password = zend.Estrdup(pass)
				ret = 0
			}
			zend.ZendStringFree(user)
		}
	}
	if ret == -1 {
		SG__().request_info.auth_password = nil
		SG__().request_info.auth_user = SG__().request_info.auth_password
	} else {
		SG__().request_info.auth_digest = nil
	}
	if ret == -1 && auth != nil && auth_len > 0 && zend.ZendBinaryStrncasecmp(auth, auth_len, "Digest ", b.SizeOf("\"Digest \"")-1, b.SizeOf("\"Digest \"")-1) == 0 {
		SG__().request_info.auth_digest = zend.Estrdup(auth + 7)
		ret = 0
	}
	if ret == -1 {
		SG__().request_info.auth_digest = nil
	}
	return ret
}
func PhpLintScript(file *zend.ZendFileHandle) int {
	var op_array *zend.ZendOpArray
	var retval int = zend.FAILURE
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		op_array = zend.ZendCompileFile(file, zend.ZEND_INCLUDE)
		zend.ZendDestroyFileHandle(file)
		if op_array != nil {
			zend.DestroyOpArray(op_array)
			zend.Efree(op_array)
			retval = zend.SUCCESS
		}
	}
	zend.EG__().SetBailout(__orig_bailout)
	if zend.EG__().GetException() != nil {
		zend.ZendExceptionError(zend.EG__().GetException(), zend.E_ERROR)
	}
	return retval
}
