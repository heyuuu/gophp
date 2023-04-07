package faults

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

/**
 * constants and global variables
 */
var ZendCeThrowable *types.ClassEntry
var ZendCeException *types.ClassEntry
var ZendCeErrorException *types.ClassEntry
var ZendCeError *types.ClassEntry
var ZendCeCompileError *types.ClassEntry
var ZendCeParseError *types.ClassEntry
var ZendCeTypeError *types.ClassEntry
var ZendCeArgumentCountError *types.ClassEntry
var ZendCeArithmeticError *types.ClassEntry
var ZendCeDivisionByZeroError *types.ClassEntry
var ZendThrowExceptionHook func(ex *types.Zval)

var DefaultExceptionHandlers zend.ZendObjectHandlers

var ZendFuncsThrowable []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("getMessage", zend.AccPublic|zend.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getCode", zend.AccPublic|zend.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getFile", zend.AccPublic|zend.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getLine", zend.AccPublic|zend.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getTrace", zend.AccPublic|zend.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getPrevious", zend.AccPublic|zend.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getTraceAsString", zend.AccPublic|zend.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("__toString", zend.AccPublic|zend.AccAbstract, nil, nil),
}
var DefaultExceptionFunctions = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__clone", zend.AccPrivate|zend.AccFinal, ZimExceptionClone, nil),
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, ZimExceptionConstruct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("message"),
		zend.MakeArgName("code"),
		zend.MakeArgName("previous"),
	}),
	types.MakeZendFunctionEntryEx("__wakeup", zend.AccPublic, ZimExceptionWakeup, nil),
	types.MakeZendFunctionEntryEx("getMessage", zend.AccPublic|zend.AccFinal, zim_exception_getMessage, nil),
	types.MakeZendFunctionEntryEx("getCode", zend.AccPublic|zend.AccFinal, zim_exception_getCode, nil),
	types.MakeZendFunctionEntryEx("getFile", zend.AccPublic|zend.AccFinal, zim_exception_getFile, nil),
	types.MakeZendFunctionEntryEx("getLine", zend.AccPublic|zend.AccFinal, zim_exception_getLine, nil),
	types.MakeZendFunctionEntryEx("getTrace", zend.AccPublic|zend.AccFinal, zim_exception_getTrace, nil),
	types.MakeZendFunctionEntryEx("getPrevious", zend.AccPublic|zend.AccFinal, zim_exception_getPrevious, nil),
	types.MakeZendFunctionEntryEx("getTraceAsString", zend.AccPublic|zend.AccFinal, zim_exception_getTraceAsString, nil),
	types.MakeZendFunctionEntryEx("__toString", 0, zim_exception___toString, nil),
}
var ErrorExceptionFunctions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.AccPublic, ZimErrorExceptionConstruct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("message"),
		zend.MakeArgName("code"),
		zend.MakeArgName("severity"),
		zend.MakeArgName("filename"),
		zend.MakeArgName("lineno"),
		zend.MakeArgName("previous"),
	}),
	types.MakeZendFunctionEntryEx("getSeverity", zend.AccPublic|zend.AccFinal, zim_error_exception_getSeverity, nil),
}

/**
 * functions
 */
func RethrowException(executeData *zend.ZendExecuteData) {
	if executeData.GetOpline().opcode != zend.ZEND_HANDLE_EXCEPTION {
		zend.EG__().SetOplineBeforeException(executeData.GetOpline())
		executeData.GetOpline() = zend.EG__().GetExceptionOp()
	}
}
func ZendImplementThrowable(interface_ *types.ClassEntry, class_type *types.ClassEntry) int {
	if zend.InstanceofFunction(class_type, ZendCeException) != 0 || zend.InstanceofFunction(class_type, ZendCeError) != 0 {
		return types.SUCCESS
	}
	ErrorNoreturn(E_ERROR, "Class %s cannot implement interface %s, extend %s or %s instead", class_type.GetName().GetVal(), interface_.GetName().GetVal(), ZendCeException.GetName().GetVal(), ZendCeError.GetName().GetVal())
	return types.FAILURE
}
func GetExceptionBase(object *types.Zval) *types.ClassEntry {
	if zend.InstanceofFunction(types.Z_OBJCE_P(object), ZendCeException) != 0 {
		return ZendCeException
	} else {
		return ZendCeError
	}
}
func ExceptionSetPrevious(exception *types.ZendObject, add_previous *types.ZendObject) {
	var previous *types.Zval
	var ancestor *types.Zval
	var ex *types.Zval
	var pv types.Zval
	var zv types.Zval
	var rv types.Zval
	var base_ce *types.ClassEntry
	if exception == nil || add_previous == nil {
		return
	}
	if exception == add_previous {
		zend.OBJ_RELEASE(add_previous)
		return
	}
	pv.SetObject(add_previous)
	if zend.InstanceofFunction(types.Z_OBJCE(pv), ZendCeThrowable) == 0 {
		ErrorNoreturn(E_CORE_ERROR, "Previous exception must implement Throwable")
		return
	}
	zv.SetObject(exception)
	ex = &zv
	for {
		ancestor = zend.ZendReadPropertyEx(GetExceptionBase(&pv), &pv, types.ZSTR_PREVIOUS, 1, &rv)
		for ancestor.IsObject() {
			if ancestor.GetObj() == ex.GetObj() {
				zend.OBJ_RELEASE(add_previous)
				return
			}
			ancestor = zend.ZendReadPropertyEx(GetExceptionBase(ancestor), ancestor, types.ZSTR_PREVIOUS, 1, &rv)
		}
		base_ce = GetExceptionBase(ex)
		previous = zend.ZendReadPropertyEx(base_ce, ex, types.ZSTR_PREVIOUS, 1, &rv)
		if previous.IsNull() {
			zend.ZendUpdatePropertyEx(base_ce, ex, types.ZSTR_PREVIOUS, &pv)
			add_previous.DelRefcount()
			return
		}
		ex = previous
		if ex.GetObj() == add_previous {
			break
		}
	}
}
func ExceptionSave() {
	if zend.EG__().GetPrevException() != nil {
		ExceptionSetPrevious(zend.EG__().GetException(), zend.EG__().GetPrevException())
	}
	if zend.EG__().GetException() != nil {
		zend.EG__().SetPrevException(zend.EG__().GetException())
	}
	zend.EG__().SetException(nil)
}
func ExceptionRestore() {
	if zend.EG__().GetPrevException() != nil {
		if zend.EG__().GetException() != nil {
			ExceptionSetPrevious(zend.EG__().GetException(), zend.EG__().GetPrevException())
		} else {
			zend.EG__().SetException(zend.EG__().GetPrevException())
		}
		zend.EG__().SetPrevException(nil)
	}
}
func ThrowExceptionInternal(exception *types.Zval) {
	if exception != nil {
		var previous *types.ZendObject = zend.EG__().GetException()
		ExceptionSetPrevious(exception.GetObj(), zend.EG__().GetException())
		zend.EG__().SetException(exception.GetObj())
		if previous != nil {
			return
		}
	}
	if zend.CurrEX() == nil {
		if exception != nil && (types.Z_OBJCE_P(exception) == ZendCeParseError || types.Z_OBJCE_P(exception) == ZendCeCompileError) {
			return
		}
		if zend.EG__().GetException() != nil {
			ExceptionError(zend.EG__().GetException(), E_ERROR)
		}
		ErrorNoreturn(E_CORE_ERROR, "Exception thrown without a stack frame")
	}
	if ZendThrowExceptionHook != nil {
		ZendThrowExceptionHook(exception)
	}
	if zend.CurrEX().GetFunc() == nil || !(zend.ZEND_USER_CODE(zend.CurrEX().GetFunc().GetType())) || zend.CurrEX().GetOpline().GetOpcode() == zend.ZEND_HANDLE_EXCEPTION {

		/* no need to rethrow the exception */

		return

		/* no need to rethrow the exception */

	}
	zend.EG__().SetOplineBeforeException(zend.CurrEX().GetOpline())
	zend.CurrEX().SetOpline(zend.EG__().GetExceptionOp())
}
func ClearException() {
	var exception *types.ZendObject
	if zend.EG__().GetPrevException() != nil {
		zend.OBJ_RELEASE(zend.EG__().GetPrevException())
		zend.EG__().SetPrevException(nil)
	}
	if zend.EG__().GetException() == nil {
		return
	}

	/* exception may have destructor */

	exception = zend.EG__().GetException()
	zend.EG__().SetException(nil)
	zend.OBJ_RELEASE(exception)
	if zend.CurrEX() != nil {
		zend.CurrEX().SetOpline(zend.EG__().GetOplineBeforeException())
	}
}
func DefaultExceptionNewEx(class_type *types.ClassEntry, skip_top_traces int) *types.ZendObject {
	var obj types.Zval
	var tmp types.Zval
	var object *types.ZendObject
	var trace types.Zval
	var base_ce *types.ClassEntry
	var filename *types.String
	object = zend.ZendObjectsNew(class_type)
	obj.SetObj(object)
	types.Z_OBJ_HT(obj) = &DefaultExceptionHandlers
	zend.ObjectPropertiesInit(object, class_type)
	if zend.CurrEX() != nil {
		zend.ZendFetchDebugBacktrace(&trace, skip_top_traces, b.Cond(zend.EG__().GetExceptionIgnoreArgs() != 0, zend.DEBUG_BACKTRACE_IGNORE_ARGS, 0), 0)
	} else {
		zend.ArrayInit(&trace)
	}
	trace.SetRefcount(0)
	base_ce = GetExceptionBase(&obj)
	if class_type != ZendCeParseError && class_type != ZendCeCompileError || !(b.Assign(&filename, zend.ZendGetCompiledFilename())) {
		tmp.SetStringVal(b.CastStrAuto(zend.ZendGetExecutedFilename()))
		zend.ZendUpdatePropertyEx(base_ce, &obj, types.ZSTR_FILE, &tmp)
		zend.ZvalPtrDtor(&tmp)
		tmp.SetLong(zend.ZendGetExecutedLineno())
		zend.ZendUpdatePropertyEx(base_ce, &obj, types.ZSTR_LINE, &tmp)
	} else {
		tmp.SetString(filename)
		zend.ZendUpdatePropertyEx(base_ce, &obj, types.ZSTR_FILE, &tmp)
		tmp.SetLong(zend.ZendGetCompiledLineno())
		zend.ZendUpdatePropertyEx(base_ce, &obj, types.ZSTR_LINE, &tmp)
	}
	zend.ZendUpdatePropertyEx(base_ce, &obj, types.ZSTR_TRACE, &trace)
	return object
}
func DefaultExceptionNew(class_type *types.ClassEntry) *types.ZendObject {
	return DefaultExceptionNewEx(class_type, 0)
}
func ErrorExceptionNew(class_type *types.ClassEntry) *types.ZendObject {
	return DefaultExceptionNewEx(class_type, 2)
}
func ZimExceptionClone(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	/* Should never be executable */
	ThrowException(nil, "Cannot clone object using __clone()", 0)
}
func ZimExceptionConstruct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var message *types.String = nil
	var code zend.ZendLong = 0
	var tmp types.Zval
	var object *types.Zval
	var previous *types.Zval = nil
	var base_ce *types.ClassEntry
	var argc int = executeData.NumArgs()
	object = zend.ZEND_THIS(executeData)
	base_ce = GetExceptionBase(object)
	if zend.ZendParseParametersEx(zpp.FlagQuiet, argc, "|SlO!", &message, &code, &previous, ZendCeThrowable) == types.FAILURE {
		var ce *types.ClassEntry
		if executeData.GetThis().IsObject() {
			ce = types.Z_OBJCE(executeData.GetThis())
		} else if executeData.GetThis().GetCe() != nil {
			ce = executeData.GetThis().GetCe()
		} else {
			ce = base_ce
		}
		ThrowError(nil, "Wrong parameters for %s([string $message [, long $code [, Throwable $previous = NULL]]])", ce.GetName().GetVal())
		return
	}
	if message != nil {
		tmp.SetString(message)
		zend.ZendUpdatePropertyEx(base_ce, object, types.ZSTR_MESSAGE, &tmp)
	}
	if code != 0 {
		tmp.SetLong(code)
		zend.ZendUpdatePropertyEx(base_ce, object, types.ZSTR_CODE, &tmp)
	}
	if previous != nil {
		zend.ZendUpdatePropertyEx(base_ce, object, types.ZSTR_PREVIOUS, previous)
	}
}
func CHECK_EXC_TYPE(object *types.Zval, id *types.String, type_ types.ZendUchar, value *types.Zval) {
	pvalue := zend.ZendReadPropertyEx(GetExceptionBase(object), object, id, 1, value)
	if pvalue.GetType() != types.IS_NULL && pvalue.GetType() != type_ {
		zend.ZendUnsetProperty(GetExceptionBase(object), object, id.GetVal(), id.GetLen())
	}
}
func ZimExceptionWakeup(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value types.Zval
	var pvalue *types.Zval
	var object *types.Zval = zend.ZEND_THIS(executeData)
	CHECK_EXC_TYPE(object, types.ZSTR_MESSAGE, types.IS_STRING, &value)
	CHECK_EXC_TYPE(object, types.ZSTR_STRING, types.IS_STRING, &value)
	CHECK_EXC_TYPE(object, types.ZSTR_CODE, types.IS_LONG, &value)
	CHECK_EXC_TYPE(object, types.ZSTR_FILE, types.IS_STRING, &value)
	CHECK_EXC_TYPE(object, types.ZSTR_LINE, types.IS_LONG, &value)
	CHECK_EXC_TYPE(object, types.ZSTR_TRACE, types.IS_ARRAY, &value)
	pvalue = zend.ZendReadProperty(GetExceptionBase(object), object, "previous", b.SizeOf("\"previous\"")-1, 1, &value)
	if pvalue != nil && pvalue.GetType() != types.IS_NULL && (pvalue.GetType() != types.IS_OBJECT || zend.InstanceofFunction(types.Z_OBJCE_P(pvalue), ZendCeThrowable) == 0 || pvalue == object) {
		zend.ZendUnsetProperty(GetExceptionBase(object), object, "previous", b.SizeOf("\"previous\"")-1)
	}
}
func ZimErrorExceptionConstruct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var message *types.String = nil
	var filename *types.String = nil
	var code zend.ZendLong = 0
	var severity zend.ZendLong = E_ERROR
	var lineno zend.ZendLong
	var tmp types.Zval
	var object *types.Zval
	var previous *types.Zval = nil
	var argc int = executeData.NumArgs()
	if zend.ZendParseParametersEx(zpp.FlagQuiet, argc, "|SllSlO!", &message, &code, &severity, &filename, &lineno, &previous, ZendCeThrowable) == types.FAILURE {
		var ce *types.ClassEntry
		if executeData.GetThis().IsObject() {
			ce = types.Z_OBJCE(executeData.GetThis())
		} else if executeData.GetThis().GetCe() != nil {
			ce = executeData.GetThis().GetCe()
		} else {
			ce = ZendCeErrorException
		}
		ThrowError(nil, "Wrong parameters for %s([string $message [, long $code, [ long $severity, [ string $filename, [ long $lineno  [, Throwable $previous = NULL]]]]]])", ce.GetName().GetVal())
		return
	}
	object = zend.ZEND_THIS(executeData)
	if message != nil {
		tmp.SetStringCopy(message)
		zend.ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_MESSAGE, &tmp)
		zend.ZvalPtrDtor(&tmp)
	}
	if code != 0 {
		tmp.SetLong(code)
		zend.ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_CODE, &tmp)
	}
	if previous != nil {
		zend.ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_PREVIOUS, previous)
	}
	tmp.SetLong(severity)
	zend.ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_SEVERITY, &tmp)
	if argc >= 4 {
		tmp.SetStringCopy(filename)
		zend.ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_FILE, &tmp)
		zend.ZvalPtrDtor(&tmp)
		if argc < 5 {
			lineno = 0
		}
		tmp.SetLong(lineno)
		zend.ZendUpdatePropertyEx(ZendCeException, object, types.ZSTR_LINE, &tmp)
	}
}
func GET_PROPERTY(object *types.Zval, id *types.String) *types.Zval {
	return zend.ZendReadPropertyEx(GetExceptionBase(object), object, id, 0, &rv)
}
func GET_PROPERTY_SILENT(object *types.Zval, id *types.String) *types.Zval {
	return zend.ZendReadPropertyEx(GetExceptionBase(object), object, id, 1, &rv)
}
func zim_exception_getFile(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(zend.ZEND_THIS(executeData), types.ZSTR_FILE)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getLine(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(zend.ZEND_THIS(executeData), types.ZSTR_LINE)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getMessage(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(zend.ZEND_THIS(executeData), types.ZSTR_MESSAGE)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getCode(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(zend.ZEND_THIS(executeData), types.ZSTR_CODE)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getTrace(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(zend.ZEND_THIS(executeData), types.ZSTR_TRACE)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_error_exception_getSeverity(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(zend.ZEND_THIS(executeData), types.ZSTR_SEVERITY)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func TRACE_APPEND_KEY(key *types.String) {
	tmp = ht.KeyFind(key.GetStr())
	if tmp {
		if tmp.GetType() != types.IS_STRING {
			Error(E_WARNING, "Value for %s is no string", key.GetVal())
			str.AppendString("[unknown]")
		} else {
			str.AppendString(b.CastStrAuto(tmp.GetStr().GetVal()))
		}
	}
}
func _buildTraceArgs(arg *types.Zval, str *zend.SmartStr) {
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
		zend.SmartStrAppendEscaped(str, arg.GetStr().GetVal(), b.Min(arg.GetStr().GetLen(), 15))
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
		zend.SmartStrAppendPrintf(str, "%.*G", int(zend.EG__().GetPrecision()), arg.GetDval())
		str.AppendString(", ")
	case types.IS_ARRAY:
		str.AppendString("Array, ")
	case types.IS_OBJECT:
		var class_name *types.String = types.Z_OBJ_HT(*arg).GetGetClassName()(types.Z_OBJ_P(arg))
		str.AppendString("Object(")
		str.AppendString(b.CastStrAuto(class_name.GetVal()))
		str.AppendString("), ")
		// types.ZendStringReleaseEx(class_name, 0)
	}
}
func _buildTraceString(str *zend.SmartStr, ht *types.Array, num uint32) {
	var file *types.Zval
	var tmp *types.Zval
	str.AppendByte('#')
	str.AppendLong(num)
	str.AppendByte(' ')
	file = ht.KeyFind(types.ZSTR_FILE.GetStr())
	if file != nil {
		if file.GetType() != types.IS_STRING {
			Error(E_WARNING, "Function name is no string")
			str.AppendString("[unknown function]")
		} else {
			var line zend.ZendLong
			tmp = ht.KeyFind(types.ZSTR_LINE.GetStr())
			if tmp != nil {
				if tmp.IsLong() {
					line = tmp.GetLval()
				} else {
					Error(E_WARNING, "Line is no long")
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
			var __ht *types.Array = tmp.GetArr()
			for _, _p := range __ht.ForeachData() {
				var _z *types.Zval = _p.GetVal()

				arg = _z
				_buildTraceArgs(arg, str)
			}
			if last_len != str.GetS().GetLen() {
				str.GetS().GetLen() -= 2
			}
		} else {
			Error(E_WARNING, "args element is no array")
		}
	}
	str.AppendString(")\n")
}
func zim_exception_getTraceAsString(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var trace *types.Zval
	var frame *types.Zval
	var rv types.Zval
	var index zend.ZendUlong
	var object *types.Zval
	var base_ce *types.ClassEntry
	var str zend.SmartStr = MakeSmartStr(0)
	var num uint32 = 0
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	object = zend.ZEND_THIS(executeData)
	base_ce = GetExceptionBase(object)
	trace = zend.ZendReadPropertyEx(base_ce, object, types.ZSTR_TRACE, 1, &rv)
	if trace.GetType() != types.IS_ARRAY {
		return_value.SetFalse()
		return
	}
	var __ht *types.Array = trace.GetArr()
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		index = _p.GetH()
		frame = _z
		if frame.GetType() != types.IS_ARRAY {
			Error(E_WARNING, "Expected array for frame "+zend.ZEND_ULONG_FMT, index)
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
func zim_exception_getPrevious(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	types.ZVAL_COPY(return_value, GET_PROPERTY_SILENT(zend.ZEND_THIS(executeData), types.ZSTR_PREVIOUS))
}
func zim_exception___toString(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var trace types.Zval
	var exception *types.Zval
	var base_ce *types.ClassEntry
	var str *types.String
	var fci types.ZendFcallInfo
	var rv types.Zval
	var tmp types.Zval
	var fname *types.String
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	str = types.NewString("")
	exception = zend.ZEND_THIS(executeData)
	fname = types.NewString("gettraceasstring")
	for exception != nil && exception.IsObject() && zend.InstanceofFunction(types.Z_OBJCE_P(exception), ZendCeThrowable) != 0 {
		var prev_str *types.String = str
		var message *types.String = zend.ZvalGetString(GET_PROPERTY(exception, types.ZSTR_MESSAGE))
		var file *types.String = zend.ZvalGetString(GET_PROPERTY(exception, types.ZSTR_FILE))
		var line zend.ZendLong = zend.ZvalGetLong(GET_PROPERTY(exception, types.ZSTR_LINE))
		fci.SetSize(b.SizeOf("fci"))
		fci.GetFunctionName().SetString(fname)
		fci.SetObject(exception.GetObj())
		fci.SetRetval(&trace)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		zend.ZendCallFunction(&fci, nil)
		if trace.GetType() != types.IS_STRING {
			zend.ZvalPtrDtor(&trace)
			trace.SetUndef()
		}
		if (types.Z_OBJCE_P(exception) == ZendCeTypeError || types.Z_OBJCE_P(exception) == ZendCeArgumentCountError) && strstr(message.GetVal(), ", called in ") {
			var real_message *types.String = zend.ZendStrpprintf(0, "%s and defined", message.GetVal())
			// types.ZendStringReleaseEx(message, 0)
			message = real_message
		}
		if message.GetLen() > 0 {
			str = zend.ZendStrpprintf(0, "%s: %s in %s:"+zend.ZEND_LONG_FMT+"\nStack trace:\n%s%s%s", types.Z_OBJCE_P(exception).GetName().GetVal(), message.GetVal(), file.GetVal(), line, b.CondF1(trace.IsString() && trace.GetStr().GetLen() != 0, func() []byte { return trace.GetStr().GetVal() }, "#0 {main}\n"), b.Cond(prev_str.GetLen() != 0, "\n\nNext ", ""), prev_str.GetVal())
		} else {
			str = zend.ZendStrpprintf(0, "%s in %s:"+zend.ZEND_LONG_FMT+"\nStack trace:\n%s%s%s", types.Z_OBJCE_P(exception).GetName().GetVal(), file.GetVal(), line, b.CondF1(trace.IsString() && trace.GetStr().GetLen() != 0, func() []byte { return trace.GetStr().GetVal() }, "#0 {main}\n"), b.Cond(prev_str.GetLen() != 0, "\n\nNext ", ""), prev_str.GetVal())
		}
		// types.ZendStringReleaseEx(prev_str, 0)
		// types.ZendStringReleaseEx(message, 0)
		// types.ZendStringReleaseEx(file, 0)
		zend.ZvalPtrDtor(&trace)
		exception.ProtectRecursive()
		exception = GET_PROPERTY(exception, types.ZSTR_PREVIOUS)
		if exception != nil && exception.IsObject() && exception.IsRecursive() {
			break
		}
	}
	// types.ZendStringReleaseEx(fname, 0)
	exception = zend.ZEND_THIS(executeData)

	/* Reset apply counts */

	for exception != nil && exception.IsObject() && b.Assign(&base_ce, GetExceptionBase(exception)) && zend.InstanceofFunction(types.Z_OBJCE_P(exception), base_ce) != 0 {
		if exception.IsRecursive() {
			exception.UnprotectRecursive()
		} else {
			break
		}
		exception = GET_PROPERTY(exception, types.ZSTR_PREVIOUS)
	}
	exception = zend.ZEND_THIS(executeData)
	base_ce = GetExceptionBase(exception)

	/* We store the result in the private property string so we can access
	 * the result in uncaught exception handlers without memleaks. */

	tmp.SetString(str)
	zend.ZendUpdatePropertyEx(base_ce, exception, types.ZSTR_STRING, &tmp)
	return_value.SetString(str)
	return
}
func RegisterDefaultException() {
	var ce types.ClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("Throwable", b.SizeOf("\"Throwable\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsThrowable)
	ZendCeThrowable = zend.ZendRegisterInternalInterface(&ce)
	ZendCeThrowable.SetInterfaceGetsImplemented(ZendImplementThrowable)
	memcpy(&DefaultExceptionHandlers, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	DefaultExceptionHandlers.SetCloneObj(nil)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("Exception", b.SizeOf("\"Exception\"")-1, 1))
	ce.SetBuiltinFunctions(DefaultExceptionFunctions)
	ZendCeException = zend.ZendRegisterInternalClassEx(&ce, nil)
	ZendCeException.SetCreateObject(DefaultExceptionNew)
	zend.ZendClassImplements(ZendCeException, 1, ZendCeThrowable)
	zend.ZendDeclarePropertyString(ZendCeException, "message", b.SizeOf("\"message\"")-1, "", zend.AccProtected)
	zend.ZendDeclarePropertyString(ZendCeException, "string", b.SizeOf("\"string\"")-1, "", zend.AccPrivate)
	zend.ZendDeclarePropertyLong(ZendCeException, "code", b.SizeOf("\"code\"")-1, 0, zend.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeException, "file", b.SizeOf("\"file\"")-1, zend.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeException, "line", b.SizeOf("\"line\"")-1, zend.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeException, "trace", b.SizeOf("\"trace\"")-1, zend.AccPrivate)
	zend.ZendDeclarePropertyNull(ZendCeException, "previous", b.SizeOf("\"previous\"")-1, zend.AccPrivate)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("ErrorException", b.SizeOf("\"ErrorException\"")-1, 1))
	ce.SetBuiltinFunctions(ErrorExceptionFunctions)
	ZendCeErrorException = zend.ZendRegisterInternalClassEx(&ce, ZendCeException)
	ZendCeErrorException.SetCreateObject(ErrorExceptionNew)
	zend.ZendDeclarePropertyLong(ZendCeErrorException, "severity", b.SizeOf("\"severity\"")-1, E_ERROR, zend.AccProtected)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("Error", b.SizeOf("\"Error\"")-1, 1))
	ce.SetBuiltinFunctions(DefaultExceptionFunctions)
	ZendCeError = zend.ZendRegisterInternalClassEx(&ce, nil)
	ZendCeError.SetCreateObject(DefaultExceptionNew)
	zend.ZendClassImplements(ZendCeError, 1, ZendCeThrowable)
	zend.ZendDeclarePropertyString(ZendCeError, "message", b.SizeOf("\"message\"")-1, "", zend.AccProtected)
	zend.ZendDeclarePropertyString(ZendCeError, "string", b.SizeOf("\"string\"")-1, "", zend.AccPrivate)
	zend.ZendDeclarePropertyLong(ZendCeError, "code", b.SizeOf("\"code\"")-1, 0, zend.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeError, "file", b.SizeOf("\"file\"")-1, zend.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeError, "line", b.SizeOf("\"line\"")-1, zend.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeError, "trace", b.SizeOf("\"trace\"")-1, zend.AccPrivate)
	zend.ZendDeclarePropertyNull(ZendCeError, "previous", b.SizeOf("\"previous\"")-1, zend.AccPrivate)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("CompileError", b.SizeOf("\"CompileError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeCompileError = zend.ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeCompileError.SetCreateObject(DefaultExceptionNew)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("ParseError", b.SizeOf("\"ParseError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeParseError = zend.ZendRegisterInternalClassEx(&ce, ZendCeCompileError)
	ZendCeParseError.SetCreateObject(DefaultExceptionNew)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("TypeError", b.SizeOf("\"TypeError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeTypeError = zend.ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeTypeError.SetCreateObject(DefaultExceptionNew)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("ArgumentCountError", b.SizeOf("\"ArgumentCountError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeArgumentCountError = zend.ZendRegisterInternalClassEx(&ce, ZendCeTypeError)
	ZendCeArgumentCountError.SetCreateObject(DefaultExceptionNew)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("ArithmeticError", b.SizeOf("\"ArithmeticError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeArithmeticError = zend.ZendRegisterInternalClassEx(&ce, ZendCeError)
	ZendCeArithmeticError.SetCreateObject(DefaultExceptionNew)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.ZendStringInitInterned("DivisionByZeroError", b.SizeOf("\"DivisionByZeroError\"")-1, 1))
	ce.SetBuiltinFunctions(nil)
	ZendCeDivisionByZeroError = zend.ZendRegisterInternalClassEx(&ce, ZendCeArithmeticError)
	ZendCeDivisionByZeroError.SetCreateObject(DefaultExceptionNew)
}
func ThrowException(exception_ce *types.ClassEntry, message string, code zend.ZendLong) *types.ZendObject {
	var ex types.Zval
	var tmp types.Zval
	if exception_ce != nil {
		if zend.InstanceofFunction(exception_ce, ZendCeThrowable) == 0 {
			Error(E_NOTICE, "Exceptions must implement Throwable")
			exception_ce = ZendCeException
		}
	} else {
		exception_ce = ZendCeException
	}
	zend.ObjectInitEx(&ex, exception_ce)
	if message {
		tmp.SetStringVal(b.CastStrAuto(message))
		zend.ZendUpdatePropertyEx(exception_ce, &ex, types.ZSTR_MESSAGE, &tmp)
		zend.ZvalPtrDtor(&tmp)
	}
	if code != 0 {
		tmp.SetLong(code)
		zend.ZendUpdatePropertyEx(exception_ce, &ex, types.ZSTR_CODE, &tmp)
	}
	ThrowExceptionInternal(&ex)
	return ex.GetObj()
}
func ThrowExceptionEx(exception_ce *types.ClassEntry, code zend.ZendLong, format string, args ...any) *types.ZendObject {
	message := zend.ZendSprintf(format, args)
	obj := ThrowException(exception_ce, message, code)
	return obj
}
func ThrowErrorException(exception_ce *types.ClassEntry, message string, code zend.ZendLong, severity int) *types.ZendObject {
	var ex types.Zval
	var tmp types.Zval
	var obj = ThrowException(exception_ce, message, code)
	ex.SetObject(obj)
	tmp.SetLong(severity)
	zend.ZendUpdatePropertyEx(ZendCeErrorException, &ex, types.ZSTR_SEVERITY, &tmp)
	return obj
}
func ErrorVa(type_ int, file *byte, lineno uint32, format string, _ ...any) {
	var args va_list
	va_start(args, format)
	zend.ZendErrorCb(type_, file, lineno, format, args)
	va_end(args)
}
func ErrorHelper(type_ int, filename *byte, lineno uint32, format string, _ ...any) {
	var va va_list
	va_start(va, format)
	zend.ZendErrorCb(type_, filename, lineno, format, va)
	va_end(va)
}
func ExceptionError(ex *types.ZendObject, severity int) {
	var exception types.Zval
	var rv types.Zval
	var ce_exception *types.ClassEntry
	exception.SetObject(ex)
	ce_exception = ex.GetCe()
	zend.EG__().SetException(nil)
	if ce_exception == ZendCeParseError || ce_exception == ZendCeCompileError {
		var message *types.String = zend.ZvalGetString(GET_PROPERTY(&exception, types.ZSTR_MESSAGE))
		var file *types.String = zend.ZvalGetString(GET_PROPERTY_SILENT(&exception, types.ZSTR_FILE))
		var line zend.ZendLong = zend.ZvalGetLong(GET_PROPERTY_SILENT(&exception, types.ZSTR_LINE))
		ErrorHelper(b.Cond(ce_exception == ZendCeParseError, E_PARSE, E_COMPILE_ERROR), file.GetVal(), line, "%s", message.GetVal())
		// types.ZendStringReleaseEx(file, 0)
		// types.ZendStringReleaseEx(message, 0)
	} else if zend.InstanceofFunction(ce_exception, ZendCeThrowable) != 0 {
		var tmp types.Zval
		var str *types.String
		var file *types.String = nil
		var line zend.ZendLong = 0
		zend.ZendCallMethodWith0Params(&exception, ce_exception, ex.GetCe().GetTostring(), "__tostring", &tmp)
		if zend.EG__().GetException() == nil {
			if tmp.GetType() != types.IS_STRING {
				Error(E_WARNING, "%s::__toString() must return a string", ce_exception.GetName().GetVal())
			} else {
				zend.ZendUpdatePropertyEx(GetExceptionBase(&exception), &exception, types.ZSTR_STRING, &tmp)
			}
		}
		zend.ZvalPtrDtor(&tmp)
		if zend.EG__().GetException() != nil {
			var zv types.Zval
			zv.SetObject(zend.EG__().GetException())

			/* do the best we can to inform about the inner exception */

			if zend.InstanceofFunction(ce_exception, ZendCeException) != 0 || zend.InstanceofFunction(ce_exception, ZendCeError) != 0 {
				file = zend.ZvalGetString(GET_PROPERTY_SILENT(&zv, types.ZSTR_FILE))
				line = zend.ZvalGetLong(GET_PROPERTY_SILENT(&zv, types.ZSTR_LINE))
			}
			ErrorVa(E_WARNING, b.CondF1(file != nil && file.GetLen() > 0, func() []byte { return file.GetVal() }, nil), line, "Uncaught %s in exception handling during call to %s::__tostring()", types.Z_OBJCE(zv).GetName().GetVal(), ce_exception.GetName().GetVal())
			if file != nil {
				// types.ZendStringReleaseEx(file, 0)
			}
		}
		str = zend.ZvalGetString(GET_PROPERTY_SILENT(&exception, types.ZSTR_STRING))
		file = zend.ZvalGetString(GET_PROPERTY_SILENT(&exception, types.ZSTR_FILE))
		line = zend.ZvalGetLong(GET_PROPERTY_SILENT(&exception, types.ZSTR_LINE))
		ErrorVa(severity, b.CondF1(file != nil && file.GetLen() > 0, func() []byte { return file.GetVal() }, nil), line, "Uncaught %s\n  thrown", str.GetVal())
		// types.ZendStringReleaseEx(str, 0)
		// types.ZendStringReleaseEx(file, 0)
	} else {
		Error(severity, "Uncaught exception '%s'", ce_exception.GetName().GetVal())
	}
	zend.OBJ_RELEASE(ex)
}
func ThrowExceptionObject(exception *types.Zval) {
	var exception_ce *types.ClassEntry
	if exception == nil || exception.GetType() != types.IS_OBJECT {
		ErrorNoreturn(E_CORE_ERROR, "Need to supply an object when throwing an exception")
	}
	exception_ce = types.Z_OBJCE_P(exception)
	if exception_ce == nil || zend.InstanceofFunction(exception_ce, ZendCeThrowable) == 0 {
		ThrowError(nil, "Cannot throw objects that do not implement Throwable")
		zend.ZvalPtrDtor(exception)
		return
	}
	ThrowExceptionInternal(exception)
}
