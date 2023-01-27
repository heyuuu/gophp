// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendRethrowException(execute_data *ZendExecuteData) {
	if EX(opline).opcode != ZEND_HANDLE_EXCEPTION {
		ExecutorGlobals.SetOplineBeforeException(EX(opline))
		EX(opline) = ExecutorGlobals.GetExceptionOp()
	}
}
func ZendImplementThrowable(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	if InstanceofFunction(class_type, ZendCeException) != 0 || InstanceofFunction(class_type, ZendCeError) != 0 {
		return SUCCESS
	}
	ZendErrorNoreturn(E_ERROR, "Class %s cannot implement interface %s, extend %s or %s instead", class_type.GetName().GetVal(), interface_.GetName().GetVal(), ZendCeException.GetName().GetVal(), ZendCeError.GetName().GetVal())
	return FAILURE
}
func IGetExceptionBase(object *Zval) *ZendClassEntry {
	if InstanceofFunction(Z_OBJCE_P(object), ZendCeException) != 0 {
		return ZendCeException
	} else {
		return ZendCeError
	}
}
func ZendGetExceptionBase(object *Zval) *ZendClassEntry { return IGetExceptionBase(object) }
func ZendExceptionSetPrevious(exception *ZendObject, add_previous *ZendObject) {
	var previous *Zval
	var ancestor *Zval
	var ex *Zval
	var pv Zval
	var zv Zval
	var rv Zval
	var base_ce *ZendClassEntry
	if exception == nil || add_previous == nil {
		return
	}
	if exception == add_previous {
		OBJ_RELEASE(add_previous)
		return
	}
	ZVAL_OBJ(&pv, add_previous)
	if InstanceofFunction(Z_OBJCE(pv), ZendCeThrowable) == 0 {
		ZendErrorNoreturn(E_CORE_ERROR, "Previous exception must implement Throwable")
		return
	}
	ZVAL_OBJ(&zv, exception)
	ex = &zv
	for {
		ancestor = ZendReadPropertyEx(IGetExceptionBase(&pv), &pv, ZSTR_KNOWN(ZEND_STR_PREVIOUS), 1, &rv)
		for ancestor.IsType(IS_OBJECT) {
			if Z_OBJ_P(ancestor) == Z_OBJ_P(ex) {
				OBJ_RELEASE(add_previous)
				return
			}
			ancestor = ZendReadPropertyEx(IGetExceptionBase(ancestor), ancestor, ZSTR_KNOWN(ZEND_STR_PREVIOUS), 1, &rv)
		}
		base_ce = IGetExceptionBase(ex)
		previous = ZendReadPropertyEx(base_ce, ex, ZSTR_KNOWN(ZEND_STR_PREVIOUS), 1, &rv)
		if previous.IsType(IS_NULL) {
			ZendUpdatePropertyEx(base_ce, ex, ZSTR_KNOWN(ZEND_STR_PREVIOUS), &pv)
			GC_DELREF(add_previous)
			return
		}
		ex = previous
		if Z_OBJ_P(ex) == add_previous {
			break
		}
	}
}
func ZendExceptionSave() {
	if ExecutorGlobals.GetPrevException() != nil {
		ZendExceptionSetPrevious(ExecutorGlobals.GetException(), ExecutorGlobals.GetPrevException())
	}
	if ExecutorGlobals.GetException() != nil {
		ExecutorGlobals.SetPrevException(ExecutorGlobals.GetException())
	}
	ExecutorGlobals.SetException(nil)
}
func ZendExceptionRestore() {
	if ExecutorGlobals.GetPrevException() != nil {
		if ExecutorGlobals.GetException() != nil {
			ZendExceptionSetPrevious(ExecutorGlobals.GetException(), ExecutorGlobals.GetPrevException())
		} else {
			ExecutorGlobals.SetException(ExecutorGlobals.GetPrevException())
		}
		ExecutorGlobals.SetPrevException(nil)
	}
}
func ZendThrowExceptionInternal(exception *Zval) {
	if exception != nil {
		var previous *ZendObject = ExecutorGlobals.GetException()
		ZendExceptionSetPrevious(Z_OBJ_P(exception), ExecutorGlobals.GetException())
		ExecutorGlobals.SetException(Z_OBJ_P(exception))
		if previous != nil {
			return
		}
	}
	if ExecutorGlobals.GetCurrentExecuteData() == nil {
		if exception != nil && (Z_OBJCE_P(exception) == ZendCeParseError || Z_OBJCE_P(exception) == ZendCeCompileError) {
			return
		}
		if ExecutorGlobals.GetException() != nil {
			ZendExceptionError(ExecutorGlobals.GetException(), E_ERROR)
		}
		ZendErrorNoreturn(E_CORE_ERROR, "Exception thrown without a stack frame")
	}
	if ZendThrowExceptionHook != nil {
		ZendThrowExceptionHook(exception)
	}
	if ExecutorGlobals.GetCurrentExecuteData().GetFunc() == nil || !(ZEND_USER_CODE(ExecutorGlobals.GetCurrentExecuteData().GetFunc().GetCommonType())) || ExecutorGlobals.GetCurrentExecuteData().GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION {

		/* no need to rethrow the exception */

		return

		/* no need to rethrow the exception */

	}
	ExecutorGlobals.SetOplineBeforeException(ExecutorGlobals.GetCurrentExecuteData().GetOpline())
	ExecutorGlobals.GetCurrentExecuteData().SetOpline(ExecutorGlobals.GetExceptionOp())
}
func ZendClearException() {
	var exception *ZendObject
	if ExecutorGlobals.GetPrevException() != nil {
		OBJ_RELEASE(ExecutorGlobals.GetPrevException())
		ExecutorGlobals.SetPrevException(nil)
	}
	if ExecutorGlobals.GetException() == nil {
		return
	}

	/* exception may have destructor */

	exception = ExecutorGlobals.GetException()
	ExecutorGlobals.SetException(nil)
	OBJ_RELEASE(exception)
	if ExecutorGlobals.GetCurrentExecuteData() != nil {
		ExecutorGlobals.GetCurrentExecuteData().SetOpline(ExecutorGlobals.GetOplineBeforeException())
	}
}
func ZendDefaultExceptionNewEx(class_type *ZendClassEntry, skip_top_traces int) *ZendObject {
	var obj Zval
	var tmp Zval
	var object *ZendObject
	var trace Zval
	var base_ce *ZendClassEntry
	var filename *ZendString
	object = ZendObjectsNew(class_type)
	obj.SetObj(object)
	Z_OBJ_HT(obj) = &DefaultExceptionHandlers
	ObjectPropertiesInit(object, class_type)
	if ExecutorGlobals.GetCurrentExecuteData() != nil {
		ZendFetchDebugBacktrace(&trace, skip_top_traces, b.Cond(ExecutorGlobals.GetExceptionIgnoreArgs() != 0, DEBUG_BACKTRACE_IGNORE_ARGS, 0), 0)
	} else {
		ArrayInit(&trace)
	}
	Z_SET_REFCOUNT(trace, 0)
	base_ce = IGetExceptionBase(&obj)
	if class_type != ZendCeParseError && class_type != ZendCeCompileError || !(b.Assign(&filename, ZendGetCompiledFilename())) {
		ZVAL_STRING(&tmp, ZendGetExecutedFilename())
		ZendUpdatePropertyEx(base_ce, &obj, ZSTR_KNOWN(ZEND_STR_FILE), &tmp)
		ZvalPtrDtor(&tmp)
		ZVAL_LONG(&tmp, ZendGetExecutedLineno())
		ZendUpdatePropertyEx(base_ce, &obj, ZSTR_KNOWN(ZEND_STR_LINE), &tmp)
	} else {
		ZVAL_STR(&tmp, filename)
		ZendUpdatePropertyEx(base_ce, &obj, ZSTR_KNOWN(ZEND_STR_FILE), &tmp)
		ZVAL_LONG(&tmp, ZendGetCompiledLineno())
		ZendUpdatePropertyEx(base_ce, &obj, ZSTR_KNOWN(ZEND_STR_LINE), &tmp)
	}
	ZendUpdatePropertyEx(base_ce, &obj, ZSTR_KNOWN(ZEND_STR_TRACE), &trace)
	return object
}
func ZendDefaultExceptionNew(class_type *ZendClassEntry) *ZendObject {
	return ZendDefaultExceptionNewEx(class_type, 0)
}
func ZendErrorExceptionNew(class_type *ZendClassEntry) *ZendObject {
	return ZendDefaultExceptionNewEx(class_type, 2)
}
func ZimExceptionClone(execute_data *ZendExecuteData, return_value *Zval) {
	/* Should never be executable */

	ZendThrowException(nil, "Cannot clone object using __clone()", 0)

	/* Should never be executable */
}
func ZimExceptionConstruct(execute_data *ZendExecuteData, return_value *Zval) {
	var message *ZendString = nil
	var code ZendLong = 0
	var tmp Zval
	var object *Zval
	var previous *Zval = nil
	var base_ce *ZendClassEntry
	var argc int = ZEND_NUM_ARGS()
	object = ZEND_THIS
	base_ce = IGetExceptionBase(object)
	if ZendParseParametersEx(ZEND_PARSE_PARAMS_QUIET, argc, "|SlO!", &message, &code, &previous, ZendCeThrowable) == FAILURE {
		var ce *ZendClassEntry
		if EX(This).u1.v.type_ == IS_OBJECT {
			ce = Z_OBJCE(EX(This))
		} else if EX(This).GetCe() != nil {
			ce = EX(This).GetCe()
		} else {
			ce = base_ce
		}
		ZendThrowError(nil, "Wrong parameters for %s([string $message [, long $code [, Throwable $previous = NULL]]])", ce.GetName().GetVal())
		return
	}
	if message != nil {
		ZVAL_STR(&tmp, message)
		ZendUpdatePropertyEx(base_ce, object, ZSTR_KNOWN(ZEND_STR_MESSAGE), &tmp)
	}
	if code != 0 {
		ZVAL_LONG(&tmp, code)
		ZendUpdatePropertyEx(base_ce, object, ZSTR_KNOWN(ZEND_STR_CODE), &tmp)
	}
	if previous != nil {
		ZendUpdatePropertyEx(base_ce, object, ZSTR_KNOWN(ZEND_STR_PREVIOUS), previous)
	}
}
func CHECK_EXC_TYPE(id ZendKnownStringId, type_ uint32) {
	pvalue = ZendReadPropertyEx(IGetExceptionBase(object), object, ZSTR_KNOWN(id), 1, &value)
	if pvalue.GetType() != IS_NULL && pvalue.GetType() != type_ {
		ZendUnsetProperty(IGetExceptionBase(object), object, ZSTR_KNOWN(id).GetVal(), ZSTR_KNOWN(id).GetLen())
	}
}
func ZimExceptionWakeup(execute_data *ZendExecuteData, return_value *Zval) {
	var value Zval
	var pvalue *Zval
	var object *Zval = ZEND_THIS
	CHECK_EXC_TYPE(ZEND_STR_MESSAGE, IS_STRING)
	CHECK_EXC_TYPE(ZEND_STR_STRING, IS_STRING)
	CHECK_EXC_TYPE(ZEND_STR_CODE, IS_LONG)
	CHECK_EXC_TYPE(ZEND_STR_FILE, IS_STRING)
	CHECK_EXC_TYPE(ZEND_STR_LINE, IS_LONG)
	CHECK_EXC_TYPE(ZEND_STR_TRACE, IS_ARRAY)
	pvalue = ZendReadProperty(IGetExceptionBase(object), object, "previous", b.SizeOf("\"previous\"")-1, 1, &value)
	if pvalue != nil && pvalue.GetType() != IS_NULL && (pvalue.GetType() != IS_OBJECT || InstanceofFunction(Z_OBJCE_P(pvalue), ZendCeThrowable) == 0 || pvalue == object) {
		ZendUnsetProperty(IGetExceptionBase(object), object, "previous", b.SizeOf("\"previous\"")-1)
	}
}
func ZimErrorExceptionConstruct(execute_data *ZendExecuteData, return_value *Zval) {
	var message *ZendString = nil
	var filename *ZendString = nil
	var code ZendLong = 0
	var severity ZendLong = E_ERROR
	var lineno ZendLong
	var tmp Zval
	var object *Zval
	var previous *Zval = nil
	var argc int = ZEND_NUM_ARGS()
	if ZendParseParametersEx(ZEND_PARSE_PARAMS_QUIET, argc, "|SllSlO!", &message, &code, &severity, &filename, &lineno, &previous, ZendCeThrowable) == FAILURE {
		var ce *ZendClassEntry
		if EX(This).u1.v.type_ == IS_OBJECT {
			ce = Z_OBJCE(EX(This))
		} else if EX(This).GetCe() != nil {
			ce = EX(This).GetCe()
		} else {
			ce = ZendCeErrorException
		}
		ZendThrowError(nil, "Wrong parameters for %s([string $message [, long $code, [ long $severity, [ string $filename, [ long $lineno  [, Throwable $previous = NULL]]]]]])", ce.GetName().GetVal())
		return
	}
	object = ZEND_THIS
	if message != nil {
		ZVAL_STR_COPY(&tmp, message)
		ZendUpdatePropertyEx(ZendCeException, object, ZSTR_KNOWN(ZEND_STR_MESSAGE), &tmp)
		ZvalPtrDtor(&tmp)
	}
	if code != 0 {
		ZVAL_LONG(&tmp, code)
		ZendUpdatePropertyEx(ZendCeException, object, ZSTR_KNOWN(ZEND_STR_CODE), &tmp)
	}
	if previous != nil {
		ZendUpdatePropertyEx(ZendCeException, object, ZSTR_KNOWN(ZEND_STR_PREVIOUS), previous)
	}
	ZVAL_LONG(&tmp, severity)
	ZendUpdatePropertyEx(ZendCeException, object, ZSTR_KNOWN(ZEND_STR_SEVERITY), &tmp)
	if argc >= 4 {
		ZVAL_STR_COPY(&tmp, filename)
		ZendUpdatePropertyEx(ZendCeException, object, ZSTR_KNOWN(ZEND_STR_FILE), &tmp)
		ZvalPtrDtor(&tmp)
		if argc < 5 {
			lineno = 0
		}
		ZVAL_LONG(&tmp, lineno)
		ZendUpdatePropertyEx(ZendCeException, object, ZSTR_KNOWN(ZEND_STR_LINE), &tmp)
	}
}
func GET_PROPERTY(object *Zval, id ZendKnownStringId) *Zval {
	return ZendReadPropertyEx(IGetExceptionBase(object), object, ZSTR_KNOWN(id), 0, &rv)
}
func GET_PROPERTY_SILENT(object *Zval, id ZendKnownStringId) *Zval {
	return ZendReadPropertyEx(IGetExceptionBase(object), object, ZSTR_KNOWN(id), 1, &rv)
}
func zim_exception_getFile(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if ZendParseParametersNone() == FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS, ZEND_STR_FILE)
	ZVAL_DEREF(prop)
	ZVAL_COPY(return_value, prop)
}
func zim_exception_getLine(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if ZendParseParametersNone() == FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS, ZEND_STR_LINE)
	ZVAL_DEREF(prop)
	ZVAL_COPY(return_value, prop)
}
func zim_exception_getMessage(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if ZendParseParametersNone() == FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS, ZEND_STR_MESSAGE)
	ZVAL_DEREF(prop)
	ZVAL_COPY(return_value, prop)
}
func zim_exception_getCode(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if ZendParseParametersNone() == FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS, ZEND_STR_CODE)
	ZVAL_DEREF(prop)
	ZVAL_COPY(return_value, prop)
}
func zim_exception_getTrace(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if ZendParseParametersNone() == FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS, ZEND_STR_TRACE)
	ZVAL_DEREF(prop)
	ZVAL_COPY(return_value, prop)
}
func zim_error_exception_getSeverity(execute_data *ZendExecuteData, return_value *Zval) {
	var prop *Zval
	var rv Zval
	if ZendParseParametersNone() == FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS, ZEND_STR_SEVERITY)
	ZVAL_DEREF(prop)
	ZVAL_COPY(return_value, prop)
}
func TRACE_APPEND_KEY(key *ZendString) {
	tmp = ht.Find(key)
	if tmp {
		if tmp.GetType() != IS_STRING {
			ZendError(E_WARNING, "Value for %s is no string", key.GetVal())
			SmartStrAppends(str, "[unknown]")
		} else {
			SmartStrAppends(str, Z_STRVAL_P(tmp))
		}
	}
}
func _buildTraceArgs(arg *Zval, str *SmartStr) {
	/* the trivial way would be to do
	 * convert_to_string_ex(arg);
	 * append it and kill the now tmp arg.
	 * but that could cause some E_NOTICE and also damn long lines.
	 */

	ZVAL_DEREF(arg)
	switch arg.GetType() {
	case IS_NULL:
		SmartStrAppends(str, "NULL, ")
		break
	case IS_STRING:
		SmartStrAppendc(str, '\'')
		SmartStrAppendEscaped(str, Z_STRVAL_P(arg), MIN(Z_STRLEN_P(arg), 15))
		if Z_STRLEN_P(arg) > 15 {
			SmartStrAppends(str, "...', ")
		} else {
			SmartStrAppends(str, "', ")
		}
		break
	case IS_FALSE:
		SmartStrAppends(str, "false, ")
		break
	case IS_TRUE:
		SmartStrAppends(str, "true, ")
		break
	case IS_RESOURCE:
		SmartStrAppends(str, "Resource id #")
		SmartStrAppendLong(str, Z_RES_HANDLE_P(arg))
		SmartStrAppends(str, ", ")
		break
	case IS_LONG:
		SmartStrAppendLong(str, Z_LVAL_P(arg))
		SmartStrAppends(str, ", ")
		break
	case IS_DOUBLE:
		SmartStrAppendPrintf(str, "%.*G", int(ExecutorGlobals.GetPrecision()), Z_DVAL_P(arg))
		SmartStrAppends(str, ", ")
		break
	case IS_ARRAY:
		SmartStrAppends(str, "Array, ")
		break
	case IS_OBJECT:
		var class_name *ZendString = Z_OBJ_HANDLER_P(arg, get_class_name)(Z_OBJ_P(arg))
		SmartStrAppends(str, "Object(")
		SmartStrAppends(str, class_name.GetVal())
		SmartStrAppends(str, "), ")
		ZendStringReleaseEx(class_name, 0)
		break
	}
}
func _buildTraceString(str *SmartStr, ht *HashTable, num uint32) {
	var file *Zval
	var tmp *Zval
	SmartStrAppendc(str, '#')
	SmartStrAppendLong(str, num)
	SmartStrAppendc(str, ' ')
	file = ht.FindEx(ZSTR_KNOWN(ZEND_STR_FILE), 1)
	if file != nil {
		if file.GetType() != IS_STRING {
			ZendError(E_WARNING, "Function name is no string")
			SmartStrAppends(str, "[unknown function]")
		} else {
			var line ZendLong
			tmp = ht.FindEx(ZSTR_KNOWN(ZEND_STR_LINE), 1)
			if tmp != nil {
				if tmp.IsType(IS_LONG) {
					line = Z_LVAL_P(tmp)
				} else {
					ZendError(E_WARNING, "Line is no long")
					line = 0
				}
			} else {
				line = 0
			}
			SmartStrAppend(str, Z_STR_P(file))
			SmartStrAppendc(str, '(')
			SmartStrAppendLong(str, line)
			SmartStrAppends(str, "): ")
		}
	} else {
		SmartStrAppends(str, "[internal function]: ")
	}
	TRACE_APPEND_KEY(ZSTR_KNOWN(ZEND_STR_CLASS))
	TRACE_APPEND_KEY(ZSTR_KNOWN(ZEND_STR_TYPE))
	TRACE_APPEND_KEY(ZSTR_KNOWN(ZEND_STR_FUNCTION))
	SmartStrAppendc(str, '(')
	tmp = ht.FindEx(ZSTR_KNOWN(ZEND_STR_ARGS), 1)
	if tmp != nil {
		if tmp.IsType(IS_ARRAY) {
			var last_len int = str.GetS().GetLen()
			var arg *Zval
			for {
				var __ht *HashTable = Z_ARRVAL_P(tmp)
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = _p.GetVal()

					if _z.IsType(IS_UNDEF) {
						continue
					}
					arg = _z
					_buildTraceArgs(arg, str)
				}
				break
			}
			if last_len != str.GetS().GetLen() {
				str.GetS().SetLen(str.GetS().GetLen() - 2)
			}
		} else {
			ZendError(E_WARNING, "args element is no array")
		}
	}
	SmartStrAppends(str, ")\n")
}
func zim_exception_getTraceAsString(execute_data *ZendExecuteData, return_value *Zval) {
	var trace *Zval
	var frame *Zval
	var rv Zval
	var index ZendUlong
	var object *Zval
	var base_ce *ZendClassEntry
	var str SmartStr = SmartStr{0}
	var num uint32 = 0
	if ZendParseParametersNone() == FAILURE {
		return
	}
	object = ZEND_THIS
	base_ce = IGetExceptionBase(object)
	trace = ZendReadPropertyEx(base_ce, object, ZSTR_KNOWN(ZEND_STR_TRACE), 1, &rv)
	if trace.GetType() != IS_ARRAY {
		RETVAL_FALSE
		return
	}
	for {
		var __ht *HashTable = Z_ARRVAL_P(trace)
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			index = _p.GetH()
			frame = _z
			if frame.GetType() != IS_ARRAY {
				ZendError(E_WARNING, "Expected array for frame "+ZEND_ULONG_FMT, index)
				continue
			}
			_buildTraceString(&str, Z_ARRVAL_P(frame), b.PostInc(&num))
		}
		break
	}
	SmartStrAppendc(&str, '#')
	SmartStrAppendLong(&str, num)
	SmartStrAppends(&str, " {main}")
	SmartStr0(&str)
	RETVAL_NEW_STR(str.GetS())
	return
}
func zim_exception_getPrevious(execute_data *ZendExecuteData, return_value *Zval) {
	var rv Zval
	if ZendParseParametersNone() == FAILURE {
		return
	}
	ZVAL_COPY(return_value, GET_PROPERTY_SILENT(ZEND_THIS, ZEND_STR_PREVIOUS))
}
func zim_exception___toString(execute_data *ZendExecuteData, return_value *Zval) {
	var trace Zval
	var exception *Zval
	var base_ce *ZendClassEntry
	var str *ZendString
	var fci ZendFcallInfo
	var rv Zval
	var tmp Zval
	var fname *ZendString
	if ZendParseParametersNone() == FAILURE {
		return
	}
	str = ZSTR_EMPTY_ALLOC()
	exception = ZEND_THIS
	fname = ZendStringInit("gettraceasstring", b.SizeOf("\"gettraceasstring\"")-1, 0)
	for exception != nil && exception.IsType(IS_OBJECT) && InstanceofFunction(Z_OBJCE_P(exception), ZendCeThrowable) != 0 {
		var prev_str *ZendString = str
		var message *ZendString = ZvalGetString(GET_PROPERTY(exception, ZEND_STR_MESSAGE))
		var file *ZendString = ZvalGetString(GET_PROPERTY(exception, ZEND_STR_FILE))
		var line ZendLong = ZvalGetLong(GET_PROPERTY(exception, ZEND_STR_LINE))
		fci.SetSize(b.SizeOf("fci"))
		ZVAL_STR(fci.GetFunctionName(), fname)
		fci.SetObject(Z_OBJ_P(exception))
		fci.SetRetval(&trace)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		ZendCallFunction(&fci, nil)
		if trace.GetType() != IS_STRING {
			ZvalPtrDtor(&trace)
			ZVAL_UNDEF(&trace)
		}
		if (Z_OBJCE_P(exception) == ZendCeTypeError || Z_OBJCE_P(exception) == ZendCeArgumentCountError) && strstr(message.GetVal(), ", called in ") {
			var real_message *ZendString = ZendStrpprintf(0, "%s and defined", message.GetVal())
			ZendStringReleaseEx(message, 0)
			message = real_message
		}
		if message.GetLen() > 0 {
			str = ZendStrpprintf(0, "%s: %s in %s:"+ZEND_LONG_FMT+"\nStack trace:\n%s%s%s", Z_OBJCE_P(exception).GetName().GetVal(), message.GetVal(), file.GetVal(), line, b.CondF1(trace.IsType(IS_STRING) && Z_STRLEN(trace) != 0, func() []byte { return Z_STRVAL(trace) }, "#0 {main}\n"), b.Cond(prev_str.GetLen() != 0, "\n\nNext ", ""), prev_str.GetVal())
		} else {
			str = ZendStrpprintf(0, "%s in %s:"+ZEND_LONG_FMT+"\nStack trace:\n%s%s%s", Z_OBJCE_P(exception).GetName().GetVal(), file.GetVal(), line, b.CondF1(trace.IsType(IS_STRING) && Z_STRLEN(trace) != 0, func() []byte { return Z_STRVAL(trace) }, "#0 {main}\n"), b.Cond(prev_str.GetLen() != 0, "\n\nNext ", ""), prev_str.GetVal())
		}
		ZendStringReleaseEx(prev_str, 0)
		ZendStringReleaseEx(message, 0)
		ZendStringReleaseEx(file, 0)
		ZvalPtrDtor(&trace)
		Z_PROTECT_RECURSION_P(exception)
		exception = GET_PROPERTY(exception, ZEND_STR_PREVIOUS)
		if exception != nil && exception.IsType(IS_OBJECT) && Z_IS_RECURSIVE_P(exception) != 0 {
			break
		}
	}
	ZendStringReleaseEx(fname, 0)
	exception = ZEND_THIS

	/* Reset apply counts */

	for exception != nil && exception.IsType(IS_OBJECT) && b.Assign(&base_ce, IGetExceptionBase(exception)) && InstanceofFunction(Z_OBJCE_P(exception), base_ce) != 0 {
		if Z_IS_RECURSIVE_P(exception) != 0 {
			Z_UNPROTECT_RECURSION_P(exception)
		} else {
			break
		}
		exception = GET_PROPERTY(exception, ZEND_STR_PREVIOUS)
	}
	exception = ZEND_THIS
	base_ce = IGetExceptionBase(exception)

	/* We store the result in the private property string so we can access
	 * the result in uncaught exception handlers without memleaks. */

	ZVAL_STR(&tmp, str)
	ZendUpdatePropertyEx(base_ce, exception, ZSTR_KNOWN(ZEND_STR_STRING), &tmp)
	RETVAL_STR(str)
	return
}
func ZendRegisterDefaultException() {
	var ce zend_class_entry
	var ce ZendClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Throwable", b.SizeOf("\"Throwable\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsThrowable)
	ZendCeThrowable = ZendRegisterInternalInterface(&ce)
	ZendCeThrowable.interface_gets_implemented = ZendImplementThrowable
	memcpy(&DefaultExceptionHandlers, &StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	DefaultExceptionHandlers.SetCloneObj(nil)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Exception", b.SizeOf("\"Exception\"")-1, 1))
	ce.SetBuiltinFunctions(DefaultExceptionFunctions)
	ZendCeException = ZendRegisterInternalClassEx(&ce, nil)
	ZendCeException.create_object = ZendDefaultExceptionNew
	ZendClassImplements(ZendCeException, 1, ZendCeThrowable)
	ZendDeclarePropertyString(ZendCeException, "message", b.SizeOf("\"message\"")-1, "", ZEND_ACC_PROTECTED)
	ZendDeclarePropertyString(ZendCeException, "string", b.SizeOf("\"string\"")-1, "", ZEND_ACC_PRIVATE)
	ZendDeclarePropertyLong(ZendCeException, "code", b.SizeOf("\"code\"")-1, 0, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeException, "file", b.SizeOf("\"file\"")-1, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeException, "line", b.SizeOf("\"line\"")-1, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeException, "trace", b.SizeOf("\"trace\"")-1, ZEND_ACC_PRIVATE)
	ZendDeclarePropertyNull(ZendCeException, "previous", b.SizeOf("\"previous\"")-1, ZEND_ACC_PRIVATE)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ErrorException", b.SizeOf("\"ErrorException\"")-1, 1))
	ce.SetBuiltinFunctions(ErrorExceptionFunctions)
	ZendCeErrorException = ZendRegisterInternalClassEx(&ce, ZendCeException)
	ZendCeErrorException.create_object = ZendErrorExceptionNew
	ZendDeclarePropertyLong(ZendCeErrorException, "severity", b.SizeOf("\"severity\"")-1, E_ERROR, ZEND_ACC_PROTECTED)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Error", b.SizeOf("\"Error\"")-1, 1))
	ce.SetBuiltinFunctions(DefaultExceptionFunctions)
	ZendCeError = ZendRegisterInternalClassEx(&ce, nil)
	ZendCeError.create_object = ZendDefaultExceptionNew
	ZendClassImplements(ZendCeError, 1, ZendCeThrowable)
	ZendDeclarePropertyString(ZendCeError, "message", b.SizeOf("\"message\"")-1, "", ZEND_ACC_PROTECTED)
	ZendDeclarePropertyString(ZendCeError, "string", b.SizeOf("\"string\"")-1, "", ZEND_ACC_PRIVATE)
	ZendDeclarePropertyLong(ZendCeError, "code", b.SizeOf("\"code\"")-1, 0, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeError, "file", b.SizeOf("\"file\"")-1, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeError, "line", b.SizeOf("\"line\"")-1, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeError, "trace", b.SizeOf("\"trace\"")-1, ZEND_ACC_PRIVATE)
	ZendDeclarePropertyNull(ZendCeError, "previous", b.SizeOf("\"previous\"")-1, ZEND_ACC_PRIVATE)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("CompileError", b.SizeOf("\"CompileError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeCompileError = ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeCompileError.create_object = ZendDefaultExceptionNew
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ParseError", b.SizeOf("\"ParseError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeParseError = ZendRegisterInternalClassEx(&ce, ZendCeCompileError)
	ZendCeParseError.create_object = ZendDefaultExceptionNew
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("TypeError", b.SizeOf("\"TypeError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeTypeError = ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeTypeError.create_object = ZendDefaultExceptionNew
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ArgumentCountError", b.SizeOf("\"ArgumentCountError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeArgumentCountError = ZendRegisterInternalClassEx(&ce, ZendCeTypeError)
	ZendCeArgumentCountError.create_object = ZendDefaultExceptionNew
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ArithmeticError", b.SizeOf("\"ArithmeticError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeArithmeticError = ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeArithmeticError.create_object = ZendDefaultExceptionNew
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("DivisionByZeroError", b.SizeOf("\"DivisionByZeroError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeDivisionByZeroError = ZendRegisterInternalClassEx(&ce, ZendCeArithmeticError)
	ZendCeDivisionByZeroError.create_object = ZendDefaultExceptionNew
}
func ZendExceptionGetDefault() *ZendClassEntry { return ZendCeException }
func ZendGetErrorException() *ZendClassEntry   { return ZendCeErrorException }
func ZendThrowException(exception_ce *ZendClassEntry, message string, code ZendLong) *ZendObject {
	var ex Zval
	var tmp Zval
	if exception_ce != nil {
		if InstanceofFunction(exception_ce, ZendCeThrowable) == 0 {
			ZendError(E_NOTICE, "Exceptions must implement Throwable")
			exception_ce = ZendCeException
		}
	} else {
		exception_ce = ZendCeException
	}
	ObjectInitEx(&ex, exception_ce)
	if message {
		ZVAL_STRING(&tmp, message)
		ZendUpdatePropertyEx(exception_ce, &ex, ZSTR_KNOWN(ZEND_STR_MESSAGE), &tmp)
		ZvalPtrDtor(&tmp)
	}
	if code != 0 {
		ZVAL_LONG(&tmp, code)
		ZendUpdatePropertyEx(exception_ce, &ex, ZSTR_KNOWN(ZEND_STR_CODE), &tmp)
	}
	ZendThrowExceptionInternal(&ex)
	return ex.GetObj()
}
func ZendThrowExceptionEx(exception_ce *ZendClassEntry, code ZendLong, format string, _ ...any) *ZendObject {
	var arg va_list
	var message *byte
	var obj *ZendObject
	va_start(arg, format)
	ZendVspprintf(&message, 0, format, arg)
	va_end(arg)
	obj = ZendThrowException(exception_ce, message, code)
	Efree(message)
	return obj
}
func ZendThrowErrorException(exception_ce *ZendClassEntry, message *byte, code ZendLong, severity int) *ZendObject {
	var ex Zval
	var tmp Zval
	var obj *ZendObject = ZendThrowException(exception_ce, message, code)
	ZVAL_OBJ(&ex, obj)
	ZVAL_LONG(&tmp, severity)
	ZendUpdatePropertyEx(ZendCeErrorException, &ex, ZSTR_KNOWN(ZEND_STR_SEVERITY), &tmp)
	return obj
}
func ZendErrorVa(type_ int, file *byte, lineno uint32, format string, _ ...any) {
	var args va_list
	va_start(args, format)
	ZendErrorCb(type_, file, lineno, format, args)
	va_end(args)
}
func ZendErrorHelper(type_ int, filename *byte, lineno uint32, format string, _ ...any) {
	var va va_list
	va_start(va, format)
	ZendErrorCb(type_, filename, lineno, format, va)
	va_end(va)
}
func ZendExceptionError(ex *ZendObject, severity int) {
	var exception Zval
	var rv Zval
	var ce_exception *ZendClassEntry
	ZVAL_OBJ(&exception, ex)
	ce_exception = ex.GetCe()
	ExecutorGlobals.SetException(nil)
	if ce_exception == ZendCeParseError || ce_exception == ZendCeCompileError {
		var message *ZendString = ZvalGetString(GET_PROPERTY(&exception, ZEND_STR_MESSAGE))
		var file *ZendString = ZvalGetString(GET_PROPERTY_SILENT(&exception, ZEND_STR_FILE))
		var line ZendLong = ZvalGetLong(GET_PROPERTY_SILENT(&exception, ZEND_STR_LINE))
		ZendErrorHelper(b.Cond(ce_exception == ZendCeParseError, E_PARSE, E_COMPILE_ERROR), file.GetVal(), line, "%s", message.GetVal())
		ZendStringReleaseEx(file, 0)
		ZendStringReleaseEx(message, 0)
	} else if InstanceofFunction(ce_exception, ZendCeThrowable) != 0 {
		var tmp Zval
		var str *ZendString
		var file *ZendString = nil
		var line ZendLong = 0
		ZendCallMethodWith0Params(&exception, ce_exception, ex.GetCe().GetTostring(), "__tostring", &tmp)
		if ExecutorGlobals.GetException() == nil {
			if tmp.GetType() != IS_STRING {
				ZendError(E_WARNING, "%s::__toString() must return a string", ce_exception.GetName().GetVal())
			} else {
				ZendUpdatePropertyEx(IGetExceptionBase(&exception), &exception, ZSTR_KNOWN(ZEND_STR_STRING), &tmp)
			}
		}
		ZvalPtrDtor(&tmp)
		if ExecutorGlobals.GetException() != nil {
			var zv Zval
			ZVAL_OBJ(&zv, ExecutorGlobals.GetException())

			/* do the best we can to inform about the inner exception */

			if InstanceofFunction(ce_exception, ZendCeException) != 0 || InstanceofFunction(ce_exception, ZendCeError) != 0 {
				file = ZvalGetString(GET_PROPERTY_SILENT(&zv, ZEND_STR_FILE))
				line = ZvalGetLong(GET_PROPERTY_SILENT(&zv, ZEND_STR_LINE))
			}
			ZendErrorVa(E_WARNING, b.CondF1(file != nil && file.GetLen() > 0, func() []byte { return file.GetVal() }, nil), line, "Uncaught %s in exception handling during call to %s::__tostring()", Z_OBJCE(zv).GetName().GetVal(), ce_exception.GetName().GetVal())
			if file != nil {
				ZendStringReleaseEx(file, 0)
			}
		}
		str = ZvalGetString(GET_PROPERTY_SILENT(&exception, ZEND_STR_STRING))
		file = ZvalGetString(GET_PROPERTY_SILENT(&exception, ZEND_STR_FILE))
		line = ZvalGetLong(GET_PROPERTY_SILENT(&exception, ZEND_STR_LINE))
		ZendErrorVa(severity, b.CondF1(file != nil && file.GetLen() > 0, func() []byte { return file.GetVal() }, nil), line, "Uncaught %s\n  thrown", str.GetVal())
		ZendStringReleaseEx(str, 0)
		ZendStringReleaseEx(file, 0)
	} else {
		ZendError(severity, "Uncaught exception '%s'", ce_exception.GetName().GetVal())
	}
	OBJ_RELEASE(ex)
}
func ZendThrowExceptionObject(exception *Zval) {
	var exception_ce *ZendClassEntry
	if exception == nil || exception.GetType() != IS_OBJECT {
		ZendErrorNoreturn(E_CORE_ERROR, "Need to supply an object when throwing an exception")
	}
	exception_ce = Z_OBJCE_P(exception)
	if exception_ce == nil || InstanceofFunction(exception_ce, ZendCeThrowable) == 0 {
		ZendThrowError(nil, "Cannot throw objects that do not implement Throwable")
		ZvalPtrDtor(exception)
		return
	}
	ZendThrowExceptionInternal(exception)
}
