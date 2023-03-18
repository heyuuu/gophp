// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/types"
)

func ZendRethrowException(executeData *ZendExecuteData) {
	if executeData.GetOpline().opcode != ZEND_HANDLE_EXCEPTION {
		EG__().SetOplineBeforeException(executeData.GetOpline())
		executeData.GetOpline() = EG__().GetExceptionOp()
	}
}
func ZendImplementThrowable(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	if InstanceofFunction(class_type, ZendCeException) != 0 || InstanceofFunction(class_type, ZendCeError) != 0 {
		return types.SUCCESS
	}
	ZendErrorNoreturn(E_ERROR, "Class %s cannot implement interface %s, extend %s or %s instead", class_type.GetName().GetVal(), interface_.GetName().GetVal(), ZendCeException.GetName().GetVal(), ZendCeError.GetName().GetVal())
	return types.FAILURE
}
func IGetExceptionBase(object *types.Zval) *ZendClassEntry {
	if InstanceofFunction(types.Z_OBJCE_P(object), ZendCeException) != 0 {
		return ZendCeException
	} else {
		return ZendCeError
	}
}
func ZendGetExceptionBase(object *types.Zval) *ZendClassEntry { return IGetExceptionBase(object) }
func ZendExceptionSetPrevious(exception *types.ZendObject, add_previous *types.ZendObject) {
	var previous *types.Zval
	var ancestor *types.Zval
	var ex *types.Zval
	var pv types.Zval
	var zv types.Zval
	var rv types.Zval
	var base_ce *ZendClassEntry
	if exception == nil || add_previous == nil {
		return
	}
	if exception == add_previous {
		OBJ_RELEASE(add_previous)
		return
	}
	pv.SetObject(add_previous)
	if InstanceofFunction(types.Z_OBJCE(pv), ZendCeThrowable) == 0 {
		ZendErrorNoreturn(E_CORE_ERROR, "Previous exception must implement Throwable")
		return
	}
	zv.SetObject(exception)
	ex = &zv
	for {
		ancestor = ZendReadPropertyEx(IGetExceptionBase(&pv), &pv, types.ZSTR_PREVIOUS, 1, &rv)
		for ancestor.IsObject() {
			if ancestor.GetObj() == ex.GetObj() {
				OBJ_RELEASE(add_previous)
				return
			}
			ancestor = ZendReadPropertyEx(IGetExceptionBase(ancestor), ancestor, types.ZSTR_PREVIOUS, 1, &rv)
		}
		base_ce = IGetExceptionBase(ex)
		previous = ZendReadPropertyEx(base_ce, ex, types.ZSTR_PREVIOUS, 1, &rv)
		if previous.IsNull() {
			ZendUpdatePropertyEx(base_ce, ex, types.ZSTR_PREVIOUS, &pv)
			add_previous.DelRefcount()
			return
		}
		ex = previous
		if ex.GetObj() == add_previous {
			break
		}
	}
}
func ZendExceptionSave() {
	if EG__().GetPrevException() != nil {
		ZendExceptionSetPrevious(EG__().GetException(), EG__().GetPrevException())
	}
	if EG__().GetException() != nil {
		EG__().SetPrevException(EG__().GetException())
	}
	EG__().SetException(nil)
}
func ZendExceptionRestore() {
	if EG__().GetPrevException() != nil {
		if EG__().GetException() != nil {
			ZendExceptionSetPrevious(EG__().GetException(), EG__().GetPrevException())
		} else {
			EG__().SetException(EG__().GetPrevException())
		}
		EG__().SetPrevException(nil)
	}
}
func ZendThrowExceptionInternal(exception *types.Zval) {
	if exception != nil {
		var previous *types.ZendObject = EG__().GetException()
		ZendExceptionSetPrevious(exception.GetObj(), EG__().GetException())
		EG__().SetException(exception.GetObj())
		if previous != nil {
			return
		}
	}
	if CurrEX() == nil {
		if exception != nil && (types.Z_OBJCE_P(exception) == ZendCeParseError || types.Z_OBJCE_P(exception) == ZendCeCompileError) {
			return
		}
		if EG__().GetException() != nil {
			ZendExceptionError(EG__().GetException(), E_ERROR)
		}
		ZendErrorNoreturn(E_CORE_ERROR, "Exception thrown without a stack frame")
	}
	if ZendThrowExceptionHook != nil {
		ZendThrowExceptionHook(exception)
	}
	if CurrEX().GetFunc() == nil || !(ZEND_USER_CODE(CurrEX().GetFunc().GetCommonType())) || CurrEX().GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION {

		/* no need to rethrow the exception */

		return

		/* no need to rethrow the exception */

	}
	EG__().SetOplineBeforeException(CurrEX().GetOpline())
	CurrEX().SetOpline(EG__().GetExceptionOp())
}
func ZendClearException() {
	var exception *types.ZendObject
	if EG__().GetPrevException() != nil {
		OBJ_RELEASE(EG__().GetPrevException())
		EG__().SetPrevException(nil)
	}
	if EG__().GetException() == nil {
		return
	}

	/* exception may have destructor */

	exception = EG__().GetException()
	EG__().SetException(nil)
	OBJ_RELEASE(exception)
	if CurrEX() != nil {
		CurrEX().SetOpline(EG__().GetOplineBeforeException())
	}
}
func ZendDefaultExceptionNewEx(class_type *ZendClassEntry, skip_top_traces int) *types.ZendObject {
	var obj types.Zval
	var tmp types.Zval
	var object *types.ZendObject
	var trace types.Zval
	var base_ce *ZendClassEntry
	var filename *types.ZendString
	object = ZendObjectsNew(class_type)
	obj.SetObj(object)
	types.Z_OBJ_HT(obj) = &DefaultExceptionHandlers
	ObjectPropertiesInit(object, class_type)
	if CurrEX() != nil {
		ZendFetchDebugBacktrace(&trace, skip_top_traces, b.Cond(EG__().GetExceptionIgnoreArgs() != 0, DEBUG_BACKTRACE_IGNORE_ARGS, 0), 0)
	} else {
		ArrayInit(&trace)
	}
	trace.SetRefcount(0)
	base_ce = IGetExceptionBase(&obj)
	if class_type != ZendCeParseError && class_type != ZendCeCompileError || !(b.Assign(&filename, ZendGetCompiledFilename())) {
		tmp.SetRawString(b.CastStrAuto(ZendGetExecutedFilename()))
		ZendUpdatePropertyEx(base_ce, &obj, types.ZSTR_FILE, &tmp)
		ZvalPtrDtor(&tmp)
		tmp.SetLong(ZendGetExecutedLineno())
		ZendUpdatePropertyEx(base_ce, &obj, types.ZSTR_LINE, &tmp)
	} else {
		tmp.SetString(filename)
		ZendUpdatePropertyEx(base_ce, &obj, types.ZSTR_FILE, &tmp)
		tmp.SetLong(ZendGetCompiledLineno())
		ZendUpdatePropertyEx(base_ce, &obj, types.ZSTR_LINE, &tmp)
	}
	ZendUpdatePropertyEx(base_ce, &obj, types.ZSTR_TRACE, &trace)
	return object
}
func ZendDefaultExceptionNew(class_type *ZendClassEntry) *types.ZendObject {
	return ZendDefaultExceptionNewEx(class_type, 0)
}
func ZendErrorExceptionNew(class_type *ZendClassEntry) *types.ZendObject {
	return ZendDefaultExceptionNewEx(class_type, 2)
}
func ZimExceptionClone(executeData *ZendExecuteData, return_value *types.Zval) {
	/* Should never be executable */

	ZendThrowException(nil, "Cannot clone object using __clone()", 0)

	/* Should never be executable */
}
func ZimExceptionConstruct(executeData *ZendExecuteData, return_value *types.Zval) {
	var message *types.ZendString = nil
	var code ZendLong = 0
	var tmp types.Zval
	var object *types.Zval
	var previous *types.Zval = nil
	var base_ce *ZendClassEntry
	var argc int = executeData.NumArgs()
	object = ZEND_THIS(executeData)
	base_ce = IGetExceptionBase(object)
	if ZendParseParametersEx(ZEND_PARSE_PARAMS_QUIET, argc, "|SlO!", &message, &code, &previous, ZendCeThrowable) == types.FAILURE {
		var ce *ZendClassEntry
		if executeData.GetThis().u1.v.type_ == types.IS_OBJECT {
			ce = types.Z_OBJCE(executeData.GetThis())
		} else if executeData.GetThis().GetCe() != nil {
			ce = executeData.GetThis().GetCe()
		} else {
			ce = base_ce
		}
		ZendThrowError(nil, "Wrong parameters for %s([string $message [, long $code [, Throwable $previous = NULL]]])", ce.GetName().GetVal())
		return
	}
	if message != nil {
		tmp.SetString(message)
		ZendUpdatePropertyEx(base_ce, object, types.ZSTR_MESSAGE, &tmp)
	}
	if code != 0 {
		tmp.SetLong(code)
		ZendUpdatePropertyEx(base_ce, object, types.ZSTR_CODE, &tmp)
	}
	if previous != nil {
		ZendUpdatePropertyEx(base_ce, object, types.ZSTR_PREVIOUS, previous)
	}
}
func CHECK_EXC_TYPE(id types.ZendKnownStringId, type_ uint32) {
	pvalue = ZendReadPropertyEx(IGetExceptionBase(object), object, types.ZSTR_KNOWN(id), 1, &value)
	if pvalue.GetType() != types.IS_NULL && pvalue.GetType() != type_ {
		ZendUnsetProperty(IGetExceptionBase(object), object, types.ZSTR_KNOWN(id).GetVal(), types.ZSTR_KNOWN(id).GetLen())
	}
}
func ZimExceptionWakeup(executeData *ZendExecuteData, return_value *types.Zval) {
	var value types.Zval
	var pvalue *types.Zval
	var object *types.Zval = ZEND_THIS(executeData)
	CHECK_EXC_TYPE(types.ZEND_STR_MESSAGE, types.IS_STRING)
	CHECK_EXC_TYPE(types.ZEND_STR_STRING, types.IS_STRING)
	CHECK_EXC_TYPE(types.ZEND_STR_CODE, types.IS_LONG)
	CHECK_EXC_TYPE(types.ZEND_STR_FILE, types.IS_STRING)
	CHECK_EXC_TYPE(types.ZEND_STR_LINE, types.IS_LONG)
	CHECK_EXC_TYPE(types.ZEND_STR_TRACE, types.IS_ARRAY)
	pvalue = ZendReadProperty(IGetExceptionBase(object), object, "previous", b.SizeOf("\"previous\"")-1, 1, &value)
	if pvalue != nil && pvalue.GetType() != types.IS_NULL && (pvalue.GetType() != types.IS_OBJECT || InstanceofFunction(types.Z_OBJCE_P(pvalue), ZendCeThrowable) == 0 || pvalue == object) {
		ZendUnsetProperty(IGetExceptionBase(object), object, "previous", b.SizeOf("\"previous\"")-1)
	}
}
func ZimErrorExceptionConstruct(executeData *ZendExecuteData, return_value *types.Zval) {
	var message *types.ZendString = nil
	var filename *types.ZendString = nil
	var code ZendLong = 0
	var severity ZendLong = E_ERROR
	var lineno ZendLong
	var tmp types.Zval
	var object *types.Zval
	var previous *types.Zval = nil
	var argc int = executeData.NumArgs()
	if ZendParseParametersEx(ZEND_PARSE_PARAMS_QUIET, argc, "|SllSlO!", &message, &code, &severity, &filename, &lineno, &previous, ZendCeThrowable) == types.FAILURE {
		var ce *ZendClassEntry
		if executeData.GetThis().u1.v.type_ == types.IS_OBJECT {
			ce = types.Z_OBJCE(executeData.GetThis())
		} else if executeData.GetThis().GetCe() != nil {
			ce = executeData.GetThis().GetCe()
		} else {
			ce = ZendCeErrorException
		}
		ZendThrowError(nil, "Wrong parameters for %s([string $message [, long $code, [ long $severity, [ string $filename, [ long $lineno  [, Throwable $previous = NULL]]]]]])", ce.GetName().GetVal())
		return
	}
	object = ZEND_THIS(executeData)
	if message != nil {
		tmp.SetStringCopy(message)
		ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_MESSAGE, &tmp)
		ZvalPtrDtor(&tmp)
	}
	if code != 0 {
		tmp.SetLong(code)
		ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_CODE, &tmp)
	}
	if previous != nil {
		ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_PREVIOUS, previous)
	}
	tmp.SetLong(severity)
	ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_SEVERITY, &tmp)
	if argc >= 4 {
		tmp.SetStringCopy(filename)
		ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_FILE, &tmp)
		ZvalPtrDtor(&tmp)
		if argc < 5 {
			lineno = 0
		}
		tmp.SetLong(lineno)
		ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_LINE, &tmp)
	}
}
func GET_PROPERTY(object *types.Zval, id types.ZendKnownStringId) *types.Zval {
	return ZendReadPropertyEx(IGetExceptionBase(object), object, types.ZSTR_KNOWN(id), 0, &rv)
}
func GET_PROPERTY_SILENT(object *types.Zval, id types.ZendKnownStringId) *types.Zval {
	return ZendReadPropertyEx(IGetExceptionBase(object), object, types.ZSTR_KNOWN(id), 1, &rv)
}
func zim_exception_getFile(executeData *ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS(executeData), types.ZEND_STR_FILE)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getLine(executeData *ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS(executeData), types.ZEND_STR_LINE)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getMessage(executeData *ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS(executeData), types.ZEND_STR_MESSAGE)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getCode(executeData *ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS(executeData), types.ZEND_STR_CODE)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getTrace(executeData *ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS(executeData), types.ZEND_STR_TRACE)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_error_exception_getSeverity(executeData *ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	prop = GET_PROPERTY(ZEND_THIS(executeData), types.ZEND_STR_SEVERITY)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func TRACE_APPEND_KEY(key *types.ZendString) {
	tmp = ht.KeyFind(key.GetStr())
	if tmp {
		if tmp.GetType() != types.IS_STRING {
			ZendError(E_WARNING, "Value for %s is no string", key.GetVal())
			str.AppendString("[unknown]")
		} else {
			str.AppendString(b.CastStrAuto(tmp.GetStr().GetVal()))
		}
	}
}
func _buildTraceArgs(arg *types.Zval, str *SmartStr) {
	/* the trivial way would be to do
	 * convert_to_string_ex(arg);
	 * append it and kill the now tmp arg.
	 * but that could cause some E_NOTICE and also damn long lines.
	 */

	arg = types.ZVAL_DEREF(arg)
	switch arg.GetType() {
	case types.IS_NULL:
		str.AppendString("NULL, ")
	case types.IS_STRING:
		str.AppendByte('\'')
		SmartStrAppendEscaped(str, arg.GetStr().GetVal(), b.Min(arg.GetStr().GetLen(), 15))
		if arg.GetStr().GetLen() > 15 {
			str.AppendString("...', ")
		} else {
			str.AppendString("', ")
		}
	case types.IS_FALSE:
		str.AppendString("false, ")
	case types.IS_TRUE:
		str.AppendString("true, ")
	case types.IS_RESOURCE:
		str.AppendString("Resource id #")
		str.AppendLong(types.Z_RES_HANDLE_P(arg))
		str.AppendString(", ")
	case types.IS_LONG:
		str.AppendLong(arg.GetLval())
		str.AppendString(", ")
	case types.IS_DOUBLE:
		SmartStrAppendPrintf(str, "%.*G", int(EG__().GetPrecision()), arg.GetDval())
		str.AppendString(", ")
	case types.IS_ARRAY:
		str.AppendString("Array, ")
	case types.IS_OBJECT:
		var class_name *types.ZendString = types.Z_OBJ_HT(*arg).GetGetClassName()(types.Z_OBJ_P(arg))
		str.AppendString("Object(")
		str.AppendString(b.CastStrAuto(class_name.GetVal()))
		str.AppendString("), ")
		types.ZendStringReleaseEx(class_name, 0)
	}
}
func _buildTraceString(str *SmartStr, ht *types.HashTable, num uint32) {
	var file *types.Zval
	var tmp *types.Zval
	str.AppendByte('#')
	str.AppendLong(num)
	str.AppendByte(' ')
	file = ht.KeyFind(types.ZSTR_FILE.GetStr())
	if file != nil {
		if file.GetType() != types.IS_STRING {
			ZendError(E_WARNING, "Function name is no string")
			str.AppendString("[unknown function]")
		} else {
			var line ZendLong
			tmp = ht.KeyFind(types.ZSTR_LINE.GetStr())
			if tmp != nil {
				if tmp.IsLong() {
					line = tmp.GetLval()
				} else {
					ZendError(E_WARNING, "Line is no long")
					line = 0
				}
			} else {
				line = 0
			}
			str.AppendString(file.GetStr().GetStr())
			str.AppendByte('(')
			str.AppendLong(line)
			str.AppendString("): ")
		}
	} else {
		str.AppendString("[internal function]: ")
	}
	TRACE_APPEND_KEY(types.ZSTR_CLASS)
	TRACE_APPEND_KEY(types.ZSTR_TYPE)
	TRACE_APPEND_KEY(types.ZSTR_FUNCTION)
	str.AppendByte('(')
	tmp = ht.KeyFind(types.ZSTR_ARGS.GetStr())
	if tmp != nil {
		if tmp.IsArray() {
			var last_len int = str.GetS().GetLen()
			var arg *types.Zval
			var __ht *types.HashTable = tmp.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()

				arg = _z
				_buildTraceArgs(arg, str)
			}
			if last_len != str.GetS().GetLen() {
				str.GetS().GetLen() -= 2
			}
		} else {
			ZendError(E_WARNING, "args element is no array")
		}
	}
	str.AppendString(")\n")
}
func zim_exception_getTraceAsString(executeData *ZendExecuteData, return_value *types.Zval) {
	var trace *types.Zval
	var frame *types.Zval
	var rv types.Zval
	var index ZendUlong
	var object *types.Zval
	var base_ce *ZendClassEntry
	var str SmartStr = MakeSmartStr(0)
	var num uint32 = 0
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	object = ZEND_THIS(executeData)
	base_ce = IGetExceptionBase(object)
	trace = ZendReadPropertyEx(base_ce, object, types.ZSTR_TRACE, 1, &rv)
	if trace.GetType() != types.IS_ARRAY {
		return_value.SetFalse()
		return
	}
	var __ht *types.HashTable = trace.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		index = _p.GetH()
		frame = _z
		if frame.GetType() != types.IS_ARRAY {
			ZendError(E_WARNING, "Expected array for frame "+ZEND_ULONG_FMT, index)
			continue
		}
		_buildTraceString(&str, frame.GetArr(), b.PostInc(&num))
	}
	str.AppendByte('#')
	str.AppendLong(num)
	str.AppendString(" {main}")
	str.ZeroTail()
	return_value.SetString(str.GetS())
	return
}
func zim_exception_getPrevious(executeData *ZendExecuteData, return_value *types.Zval) {
	var rv types.Zval
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	types.ZVAL_COPY(return_value, GET_PROPERTY_SILENT(ZEND_THIS(executeData), types.ZEND_STR_PREVIOUS))
}
func zim_exception___toString(executeData *ZendExecuteData, return_value *types.Zval) {
	var trace types.Zval
	var exception *types.Zval
	var base_ce *ZendClassEntry
	var str *types.ZendString
	var fci ZendFcallInfo
	var rv types.Zval
	var tmp types.Zval
	var fname *types.ZendString
	if ZendParseParametersNone() == types.FAILURE {
		return
	}
	str = types.ZSTR_EMPTY_ALLOC()
	exception = ZEND_THIS(executeData)
	fname = types.ZendStringInit("gettraceasstring", b.SizeOf("\"gettraceasstring\"")-1, 0)
	for exception != nil && exception.IsObject() && InstanceofFunction(types.Z_OBJCE_P(exception), ZendCeThrowable) != 0 {
		var prev_str *types.ZendString = str
		var message *types.ZendString = ZvalGetString(GET_PROPERTY(exception, types.ZEND_STR_MESSAGE))
		var file *types.ZendString = ZvalGetString(GET_PROPERTY(exception, types.ZEND_STR_FILE))
		var line ZendLong = ZvalGetLong(GET_PROPERTY(exception, types.ZEND_STR_LINE))
		fci.SetSize(b.SizeOf("fci"))
		fci.GetFunctionName().SetString(fname)
		fci.SetObject(exception.GetObj())
		fci.SetRetval(&trace)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		ZendCallFunction(&fci, nil)
		if trace.GetType() != types.IS_STRING {
			ZvalPtrDtor(&trace)
			trace.SetUndef()
		}
		if (types.Z_OBJCE_P(exception) == ZendCeTypeError || types.Z_OBJCE_P(exception) == ZendCeArgumentCountError) && strstr(message.GetVal(), ", called in ") {
			var real_message *types.ZendString = ZendStrpprintf(0, "%s and defined", message.GetVal())
			types.ZendStringReleaseEx(message, 0)
			message = real_message
		}
		if message.GetLen() > 0 {
			str = ZendStrpprintf(0, "%s: %s in %s:"+ZEND_LONG_FMT+"\nStack trace:\n%s%s%s", types.Z_OBJCE_P(exception).GetName().GetVal(), message.GetVal(), file.GetVal(), line, b.CondF1(trace.IsString() && trace.GetStr().GetLen() != 0, func() []byte { return trace.GetStr().GetVal() }, "#0 {main}\n"), b.Cond(prev_str.GetLen() != 0, "\n\nNext ", ""), prev_str.GetVal())
		} else {
			str = ZendStrpprintf(0, "%s in %s:"+ZEND_LONG_FMT+"\nStack trace:\n%s%s%s", types.Z_OBJCE_P(exception).GetName().GetVal(), file.GetVal(), line, b.CondF1(trace.IsString() && trace.GetStr().GetLen() != 0, func() []byte { return trace.GetStr().GetVal() }, "#0 {main}\n"), b.Cond(prev_str.GetLen() != 0, "\n\nNext ", ""), prev_str.GetVal())
		}
		types.ZendStringReleaseEx(prev_str, 0)
		types.ZendStringReleaseEx(message, 0)
		types.ZendStringReleaseEx(file, 0)
		ZvalPtrDtor(&trace)
		exception.ProtectRecursive()
		exception = GET_PROPERTY(exception, types.ZEND_STR_PREVIOUS)
		if exception != nil && exception.IsObject() && exception.IsRecursive() {
			break
		}
	}
	types.ZendStringReleaseEx(fname, 0)
	exception = ZEND_THIS(executeData)

	/* Reset apply counts */

	for exception != nil && exception.IsObject() && b.Assign(&base_ce, IGetExceptionBase(exception)) && InstanceofFunction(types.Z_OBJCE_P(exception), base_ce) != 0 {
		if exception.IsRecursive() {
			exception.UnprotectRecursive()
		} else {
			break
		}
		exception = GET_PROPERTY(exception, types.ZEND_STR_PREVIOUS)
	}
	exception = ZEND_THIS(executeData)
	base_ce = IGetExceptionBase(exception)

	/* We store the result in the private property string so we can access
	 * the result in uncaught exception handlers without memleaks. */

	tmp.SetString(str)
	ZendUpdatePropertyEx(base_ce, exception, types.ZSTR_STRING, &tmp)
	return_value.SetString(str)
	return
}
func ZendRegisterDefaultException() {
	var ce zend_class_entry
	var ce ZendClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("Throwable", b.SizeOf("\"Throwable\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsThrowable)
	ZendCeThrowable = ZendRegisterInternalInterface(&ce)
	ZendCeThrowable.SetInterfaceGetsImplemented(ZendImplementThrowable)
	memcpy(&DefaultExceptionHandlers, &StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	DefaultExceptionHandlers.SetCloneObj(nil)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("Exception", b.SizeOf("\"Exception\"")-1, 1))
	ce.SetBuiltinFunctions(DefaultExceptionFunctions)
	ZendCeException = ZendRegisterInternalClassEx(&ce, nil)
	ZendCeException.SetCreateObject(ZendDefaultExceptionNew)
	ZendClassImplements(ZendCeException, 1, ZendCeThrowable)
	ZendDeclarePropertyString(ZendCeException, "message", b.SizeOf("\"message\"")-1, "", ZEND_ACC_PROTECTED)
	ZendDeclarePropertyString(ZendCeException, "string", b.SizeOf("\"string\"")-1, "", ZEND_ACC_PRIVATE)
	ZendDeclarePropertyLong(ZendCeException, "code", b.SizeOf("\"code\"")-1, 0, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeException, "file", b.SizeOf("\"file\"")-1, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeException, "line", b.SizeOf("\"line\"")-1, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeException, "trace", b.SizeOf("\"trace\"")-1, ZEND_ACC_PRIVATE)
	ZendDeclarePropertyNull(ZendCeException, "previous", b.SizeOf("\"previous\"")-1, ZEND_ACC_PRIVATE)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("ErrorException", b.SizeOf("\"ErrorException\"")-1, 1))
	ce.SetBuiltinFunctions(ErrorExceptionFunctions)
	ZendCeErrorException = ZendRegisterInternalClassEx(&ce, ZendCeException)
	ZendCeErrorException.SetCreateObject(ZendErrorExceptionNew)
	ZendDeclarePropertyLong(ZendCeErrorException, "severity", b.SizeOf("\"severity\"")-1, E_ERROR, ZEND_ACC_PROTECTED)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("Error", b.SizeOf("\"Error\"")-1, 1))
	ce.SetBuiltinFunctions(DefaultExceptionFunctions)
	ZendCeError = ZendRegisterInternalClassEx(&ce, nil)
	ZendCeError.SetCreateObject(ZendDefaultExceptionNew)
	ZendClassImplements(ZendCeError, 1, ZendCeThrowable)
	ZendDeclarePropertyString(ZendCeError, "message", b.SizeOf("\"message\"")-1, "", ZEND_ACC_PROTECTED)
	ZendDeclarePropertyString(ZendCeError, "string", b.SizeOf("\"string\"")-1, "", ZEND_ACC_PRIVATE)
	ZendDeclarePropertyLong(ZendCeError, "code", b.SizeOf("\"code\"")-1, 0, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeError, "file", b.SizeOf("\"file\"")-1, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeError, "line", b.SizeOf("\"line\"")-1, ZEND_ACC_PROTECTED)
	ZendDeclarePropertyNull(ZendCeError, "trace", b.SizeOf("\"trace\"")-1, ZEND_ACC_PRIVATE)
	ZendDeclarePropertyNull(ZendCeError, "previous", b.SizeOf("\"previous\"")-1, ZEND_ACC_PRIVATE)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("CompileError", b.SizeOf("\"CompileError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeCompileError = ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeCompileError.SetCreateObject(ZendDefaultExceptionNew)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("ParseError", b.SizeOf("\"ParseError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeParseError = ZendRegisterInternalClassEx(&ce, ZendCeCompileError)
	ZendCeParseError.SetCreateObject(ZendDefaultExceptionNew)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("TypeError", b.SizeOf("\"TypeError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeTypeError = ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeTypeError.SetCreateObject(ZendDefaultExceptionNew)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("ArgumentCountError", b.SizeOf("\"ArgumentCountError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeArgumentCountError = ZendRegisterInternalClassEx(&ce, ZendCeTypeError)
	ZendCeArgumentCountError.SetCreateObject(ZendDefaultExceptionNew)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("ArithmeticError", b.SizeOf("\"ArithmeticError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeArithmeticError = ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeArithmeticError.SetCreateObject(ZendDefaultExceptionNew)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("DivisionByZeroError", b.SizeOf("\"DivisionByZeroError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeDivisionByZeroError = ZendRegisterInternalClassEx(&ce, ZendCeArithmeticError)
	ZendCeDivisionByZeroError.SetCreateObject(ZendDefaultExceptionNew)
}
func ZendThrowException(exception_ce *ZendClassEntry, message string, code ZendLong) *types.ZendObject {
	var ex types.Zval
	var tmp types.Zval
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
		tmp.SetRawString(b.CastStrAuto(message))
		ZendUpdatePropertyEx(exception_ce, &ex, types.ZSTR_MESSAGE, &tmp)
		ZvalPtrDtor(&tmp)
	}
	if code != 0 {
		tmp.SetLong(code)
		ZendUpdatePropertyEx(exception_ce, &ex, types.ZSTR_CODE, &tmp)
	}
	ZendThrowExceptionInternal(&ex)
	return ex.GetObj()
}
func ZendThrowExceptionEx(exception_ce *ZendClassEntry, code ZendLong, format string, args ...any) *types.ZendObject {
	message := ZendSprintf(format, args)
	obj := ZendThrowException(exception_ce, message, code)
	return obj
}
func ZendThrowErrorException(exception_ce *ZendClassEntry, message string, code ZendLong, severity int) *types.ZendObject {
	var ex types.Zval
	var tmp types.Zval
	var obj = ZendThrowException(exception_ce, message, code)
	ex.SetObject(obj)
	tmp.SetLong(severity)
	ZendUpdatePropertyEx(ZendCeErrorException, &ex, types.ZSTR_SEVERITY, &tmp)
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
func ZendExceptionError(ex *types.ZendObject, severity int) {
	var exception types.Zval
	var rv types.Zval
	var ce_exception *ZendClassEntry
	exception.SetObject(ex)
	ce_exception = ex.GetCe()
	EG__().SetException(nil)
	if ce_exception == ZendCeParseError || ce_exception == ZendCeCompileError {
		var message *types.ZendString = ZvalGetString(GET_PROPERTY(&exception, types.ZEND_STR_MESSAGE))
		var file *types.ZendString = ZvalGetString(GET_PROPERTY_SILENT(&exception, types.ZEND_STR_FILE))
		var line ZendLong = ZvalGetLong(GET_PROPERTY_SILENT(&exception, types.ZEND_STR_LINE))
		ZendErrorHelper(b.Cond(ce_exception == ZendCeParseError, E_PARSE, E_COMPILE_ERROR), file.GetVal(), line, "%s", message.GetVal())
		types.ZendStringReleaseEx(file, 0)
		types.ZendStringReleaseEx(message, 0)
	} else if InstanceofFunction(ce_exception, ZendCeThrowable) != 0 {
		var tmp types.Zval
		var str *types.ZendString
		var file *types.ZendString = nil
		var line ZendLong = 0
		ZendCallMethodWith0Params(&exception, ce_exception, ex.GetCe().GetTostring(), "__tostring", &tmp)
		if EG__().GetException() == nil {
			if tmp.GetType() != types.IS_STRING {
				ZendError(E_WARNING, "%s::__toString() must return a string", ce_exception.GetName().GetVal())
			} else {
				ZendUpdatePropertyEx(IGetExceptionBase(&exception), &exception, types.ZSTR_STRING, &tmp)
			}
		}
		ZvalPtrDtor(&tmp)
		if EG__().GetException() != nil {
			var zv types.Zval
			zv.SetObject(EG__().GetException())

			/* do the best we can to inform about the inner exception */

			if InstanceofFunction(ce_exception, ZendCeException) != 0 || InstanceofFunction(ce_exception, ZendCeError) != 0 {
				file = ZvalGetString(GET_PROPERTY_SILENT(&zv, types.ZEND_STR_FILE))
				line = ZvalGetLong(GET_PROPERTY_SILENT(&zv, types.ZEND_STR_LINE))
			}
			ZendErrorVa(E_WARNING, b.CondF1(file != nil && file.GetLen() > 0, func() []byte { return file.GetVal() }, nil), line, "Uncaught %s in exception handling during call to %s::__tostring()", types.Z_OBJCE(zv).GetName().GetVal(), ce_exception.GetName().GetVal())
			if file != nil {
				types.ZendStringReleaseEx(file, 0)
			}
		}
		str = ZvalGetString(GET_PROPERTY_SILENT(&exception, types.ZEND_STR_STRING))
		file = ZvalGetString(GET_PROPERTY_SILENT(&exception, types.ZEND_STR_FILE))
		line = ZvalGetLong(GET_PROPERTY_SILENT(&exception, types.ZEND_STR_LINE))
		ZendErrorVa(severity, b.CondF1(file != nil && file.GetLen() > 0, func() []byte { return file.GetVal() }, nil), line, "Uncaught %s\n  thrown", str.GetVal())
		types.ZendStringReleaseEx(str, 0)
		types.ZendStringReleaseEx(file, 0)
	} else {
		ZendError(severity, "Uncaught exception '%s'", ce_exception.GetName().GetVal())
	}
	OBJ_RELEASE(ex)
}
func ZendThrowExceptionObject(exception *types.Zval) {
	var exception_ce *ZendClassEntry
	if exception == nil || exception.GetType() != types.IS_OBJECT {
		ZendErrorNoreturn(E_CORE_ERROR, "Need to supply an object when throwing an exception")
	}
	exception_ce = types.Z_OBJCE_P(exception)
	if exception_ce == nil || InstanceofFunction(exception_ce, ZendCeThrowable) == 0 {
		ZendThrowError(nil, "Cannot throw objects that do not implement Throwable")
		ZvalPtrDtor(exception)
		return
	}
	ZendThrowExceptionInternal(exception)
}
