package faults

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core/pfmt"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/operators"
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

var DefaultExceptionHandlers types.ObjectHandlers

var ZendFuncsThrowable = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("getMessage", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getCode", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getFile", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getLine", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getTrace", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getPrevious", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("getTraceAsString", types.AccPublic|types.AccAbstract, nil, nil),
	types.MakeZendFunctionEntryEx("__toString", types.AccPublic|types.AccAbstract, nil, nil),
}
var DefaultExceptionFunctions = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__clone", types.AccPrivate|types.AccFinal, ZimExceptionClone, nil),
	types.MakeZendFunctionEntryEx("__construct", types.AccPublic, ZimExceptionConstruct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("message"),
		zend.MakeArgName("code"),
		zend.MakeArgName("previous"),
	}),
	types.MakeZendFunctionEntryEx("__wakeup", types.AccPublic, ZimExceptionWakeup, nil),
	types.MakeZendFunctionEntryEx("getMessage", types.AccPublic|types.AccFinal, zim_exception_getMessage, nil),
	types.MakeZendFunctionEntryEx("getCode", types.AccPublic|types.AccFinal, zim_exception_getCode, nil),
	types.MakeZendFunctionEntryEx("getFile", types.AccPublic|types.AccFinal, zim_exception_getFile, nil),
	types.MakeZendFunctionEntryEx("getLine", types.AccPublic|types.AccFinal, zim_exception_getLine, nil),
	types.MakeZendFunctionEntryEx("getTrace", types.AccPublic|types.AccFinal, zim_exception_getTrace, nil),
	types.MakeZendFunctionEntryEx("getPrevious", types.AccPublic|types.AccFinal, zim_exception_getPrevious, nil),
	types.MakeZendFunctionEntryEx("getTraceAsString", types.AccPublic|types.AccFinal, zim_exception_getTraceAsString, nil),
	types.MakeZendFunctionEntryEx("__toString", 0, zim_exception___toString, nil),
}
var ErrorExceptionFunctions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", types.AccPublic, ZimErrorExceptionConstruct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("message"),
		zend.MakeArgName("code"),
		zend.MakeArgName("severity"),
		zend.MakeArgName("filename"),
		zend.MakeArgName("lineno"),
		zend.MakeArgName("previous"),
	}),
	types.MakeZendFunctionEntryEx("getSeverity", types.AccPublic|types.AccFinal, zim_error_exception_getSeverity, nil),
}

/**
 * functions
 */
func RethrowException(executeData *zend.ZendExecuteData) {
	if executeData.GetOpline().GetOpcode() != zend.ZEND_HANDLE_EXCEPTION {
		zend.EG__().SetOplineBeforeException(executeData.GetOpline())
		executeData.SetOpline(zend.EG__().GetExceptionOp())
	}
}
func ZendImplementThrowable(interface_ *types.ClassEntry, classType *types.ClassEntry) int {
	if operators.InstanceofFunction(classType, ZendCeException) || operators.InstanceofFunction(classType, ZendCeError) {
		return types.SUCCESS
	}
	ErrorNoreturn(E_ERROR, "Class %s cannot implement interface %s, extend %s or %s instead", classType.Name(), interface_.Name(), ZendCeException.Name(), ZendCeError.Name())
	return types.FAILURE
}
func GetExceptionBase(object *types.Zval) *types.ClassEntry {
	if operators.InstanceofFunction(types.Z_OBJCE_P(object), ZendCeException) {
		return ZendCeException
	} else {
		return ZendCeError
	}
}
func ExceptionSetPrevious(exception *types.Object, add_previous *types.Object) {
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
		// zend.OBJ_RELEASE(add_previous)
		return
	}
	pv.SetObject(add_previous)
	if !operators.InstanceofFunction(types.Z_OBJCE(pv), ZendCeThrowable) {
		ErrorNoreturn(E_CORE_ERROR, "Previous exception must implement Throwable")
		return
	}
	zv.SetObject(exception)
	ex = &zv
	for {
		ancestor = zend.ZendReadProperty(GetExceptionBase(&pv), &pv, types.STR_PREVIOUS, true, &rv)
		for ancestor.IsObject() {
			if ancestor.Object() == ex.Object() {
				// zend.OBJ_RELEASE(add_previous)
				return
			}
			ancestor = zend.ZendReadProperty(GetExceptionBase(ancestor), ancestor, types.STR_PREVIOUS, true, &rv)
		}
		base_ce = GetExceptionBase(ex)
		previous = zend.ZendReadProperty(base_ce, ex, types.STR_PREVIOUS, 1, &rv)
		if previous.IsNull() {
			zend.ZendUpdatePropertyEx(base_ce, ex, types.STR_PREVIOUS, &pv)
			return
		}
		ex = previous
		if ex.Object() == add_previous {
			break
		}
	}
}
func ThrowExceptionInternal(exception *types.Zval) {
	if exception != nil {
		var previous *types.Object = zend.EG__().GetException()
		ExceptionSetPrevious(exception.Object(), zend.EG__().GetException())
		zend.EG__().SetException(exception.Object())
		if previous != nil {
			return
		}
	}
	if zend.CurrEX() == nil {
		if exception != nil && (types.Z_OBJCE_P(exception) == ZendCeParseError || types.Z_OBJCE_P(exception) == ZendCeCompileError) {
			return
		}
		if zend.EG__().HasException() {
			ExceptionError(zend.EG__().GetException(), E_ERROR)
		}
		ErrorNoreturn(E_CORE_ERROR, "Exception thrown without a stack frame")
	}
	if zend.CurrEX().GetFunc() == nil || !(zend.ZEND_USER_CODE(zend.CurrEX().GetFunc().GetType())) || zend.CurrEX().GetOpline().GetOpcode() == zend.ZEND_HANDLE_EXCEPTION {
		/* no need to rethrow the exception */
		return
	}
	zend.EG__().SetOplineBeforeException(zend.CurrEX().GetOpline())
	zend.CurrEX().SetOpline(zend.EG__().GetExceptionOp())
}
func DefaultExceptionNewEx(class_type *types.ClassEntry, skip_top_traces int) *types.Object {
	var obj types.Zval
	var tmp types.Zval
	var trace types.Zval
	var base_ce *types.ClassEntry
	var filename string

	object := types.NewObject(class_type, &DefaultExceptionHandlers)
	obj.SetObject(object)
	if zend.CurrEX() != nil {
		zend.ZendFetchDebugBacktrace(&trace, skip_top_traces, lang.Cond(zend.EG__().GetExceptionIgnoreArgs(), zend.DEBUG_BACKTRACE_IGNORE_ARGS, 0), 0)
	} else {
		zend.ArrayInit(&trace)
	}
	//trace.SetRefcount(0)
	base_ce = GetExceptionBase(&obj)
	if class_type != ZendCeParseError && class_type != ZendCeCompileError || !(lang.Assign(&filename, zend.ZendGetCompiledFilename())) {
		tmp.SetString(b.CastStrAuto(zend.ZendGetExecutedFilename()))
		zend.ZendUpdatePropertyEx(base_ce, &obj, types.STR_FILE, &tmp)
		// zend.ZvalPtrDtor(&tmp)
		tmp.SetLong(zend.ZendGetExecutedLineno())
		zend.ZendUpdatePropertyEx(base_ce, &obj, types.STR_LINE, &tmp)
	} else {
		tmp.SetString(filename)
		zend.ZendUpdatePropertyEx(base_ce, &obj, types.STR_FILE, &tmp)
		tmp.SetLong(zend.ZendGetCompiledLineno())
		zend.ZendUpdatePropertyEx(base_ce, &obj, types.STR_LINE, &tmp)
	}
	zend.ZendUpdatePropertyEx(base_ce, &obj, types.STR_TRACE, &trace)
	return object
}
func DefaultExceptionNew(class_type *types.ClassEntry) *types.Object {
	return DefaultExceptionNewEx(class_type, 0)
}
func ErrorExceptionNew(class_type *types.ClassEntry) *types.Object {
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
	object = executeData.ThisObjectZval()
	base_ce = GetExceptionBase(object)
	if zend.ZendParseParametersEx(zpp.FlagQuiet, argc, "|SlO!", &message, &code, &previous, ZendCeThrowable) == types.FAILURE {
		var ce *types.ClassEntry
		if executeData.InScope() {
			ce = executeData.ThisClass()
		} else {
			ce = base_ce
		}
		ThrowError(nil, "Wrong parameters for %s([string $message [, long $code [, Throwable $previous = NULL]]])", ce.Name())
		return
	}
	if message != nil {
		tmp.SetStringEx(message)
		zend.ZendUpdatePropertyEx(base_ce, object, types.STR_MESSAGE, &tmp)
	}
	if code != 0 {
		tmp.SetLong(code)
		zend.ZendUpdatePropertyEx(base_ce, object, types.STR_CODE, &tmp)
	}
	if previous != nil {
		zend.ZendUpdatePropertyEx(base_ce, object, types.STR_PREVIOUS, previous)
	}
}
func CHECK_EXC_TYPE(object *types.Zval, id string, type_ uint8, value *types.Zval) {
	pvalue := zend.ZendReadProperty(GetExceptionBase(object), object, id, 1, value)
	if !pvalue.IsNull() && pvalue.Type() != type_ {
		zend.ZendUnsetProperty(GetExceptionBase(object), object, id)
	}
}
func ZimExceptionWakeup(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value types.Zval
	var pvalue *types.Zval
	var object *types.Zval = executeData.ThisObjectZval()
	CHECK_EXC_TYPE(object, types.STR_MESSAGE, types.IsString, &value)
	CHECK_EXC_TYPE(object, types.STR_STRING, types.IsString, &value)
	CHECK_EXC_TYPE(object, types.STR_CODE, types.IsLong, &value)
	CHECK_EXC_TYPE(object, types.STR_FILE, types.IsString, &value)
	CHECK_EXC_TYPE(object, types.STR_LINE, types.IsLong, &value)
	CHECK_EXC_TYPE(object, types.STR_TRACE, types.IsArray, &value)
	pvalue = zend.ZendReadProperty(GetExceptionBase(object), object, "previous", 1, &value)
	if pvalue != nil && !pvalue.IsNull() && (!pvalue.IsObject() || operators.InstanceofFunction(types.Z_OBJCE_P(pvalue), ZendCeThrowable) == 0 || pvalue == object) {
		zend.ZendUnsetProperty(GetExceptionBase(object), object, "previous")
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
		if executeData.InScope() {
			ce = executeData.ThisClass()
		} else {
			ce = ZendCeErrorException
		}
		ThrowError(nil, "Wrong parameters for %s([string $message [, long $code, [ long $severity, [ string $filename, [ long $lineno  [, Throwable $previous = NULL]]]]]])", ce.Name())
		return
	}
	object = executeData.ThisObjectZval()
	if message != nil {
		tmp.SetString(message.GetStr())
		zend.ZendUpdatePropertyEx(ZendCeException, object, types.STR_MESSAGE, &tmp)
		// zend.ZvalPtrDtor(&tmp)
	}
	if code != 0 {
		tmp.SetLong(code)
		zend.ZendUpdatePropertyEx(ZendCeException, object, types.STR_CODE, &tmp)
	}
	if previous != nil {
		zend.ZendUpdatePropertyEx(ZendCeException, object, types.STR_PREVIOUS, previous)
	}
	tmp.SetLong(severity)
	zend.ZendUpdatePropertyEx(ZendCeException, object, types.STR_SEVERITY, &tmp)
	if argc >= 4 {
		tmp.SetString(filename.GetStr())
		zend.ZendUpdatePropertyEx(ZendCeException, object, types.STR_FILE, &tmp)
		// zend.ZvalPtrDtor(&tmp)
		if argc < 5 {
			lineno = 0
		}
		tmp.SetLong(lineno)
		zend.ZendUpdatePropertyEx(ZendCeException, object, types.STR_LINE, &tmp)
	}
}
func GET_PROPERTY(object *types.Zval, id string, rv *types.Zval) *types.Zval {
	return zend.ZendReadProperty(GetExceptionBase(object), object, id, false, rv)
}
func GET_PROPERTY_SILENT(object *types.Zval, id string, rv *types.Zval) *types.Zval {
	return zend.ZendReadProperty(GetExceptionBase(object), object, id, true, rv)
}

func zim_exception_getFile(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(executeData.ThisObjectZval(), types.STR_FILE, &rv)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getLine(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(executeData.ThisObjectZval(), types.STR_LINE, &rv)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getMessage(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(executeData.ThisObjectZval(), types.STR_MESSAGE, &rv)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getCode(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(executeData.ThisObjectZval(), types.STR_CODE, &rv)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_exception_getTrace(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(executeData.ThisObjectZval(), types.STR_TRACE, &rv)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func zim_error_exception_getSeverity(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prop *types.Zval
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	prop = GET_PROPERTY(executeData.ThisObjectZval(), types.STR_SEVERITY, &rv)
	prop = types.ZVAL_DEREF(prop)
	types.ZVAL_COPY(return_value, prop)
}
func traceAppendKey(ht *types.Array, key string) string {
	tmp := ht.KeyFind(key)
	if tmp != nil {
		if !tmp.IsString() {
			Error(E_WARNING, "Value for %s is no string", key)
			return "[unknown]"
		} else {
			return tmp.String()
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
	switch arg.Type() {
	case types.IsNull:
		str.WriteString("NULL, ")
	case types.IsString:
		argStr := arg.String()
		if len(argStr) > 15 {
			str.WriteByte('\'')
			str.WriteEscaped(argStr[:15])
			str.WriteString("...', ")
		} else {
			str.WriteByte('\'')
			str.WriteEscaped(argStr)
			str.WriteString("', ")
		}
	case types.IsFalse:
		str.WriteString("false, ")
	case types.IsTrue:
		str.WriteString("true, ")
	case types.IsResource:
		str.WriteString("Resource id #")
		str.WriteLong(arg.ResourceHandle())
		str.WriteString(", ")
	case types.IsLong:
		str.WriteLong(arg.Long())
		str.WriteString(", ")
	case types.IsDouble:
		str.WriteString(pfmt.Sprintf("%.*G", int(zend.EG__().GetPrecision()), arg.Double()))
		str.WriteString(", ")
	case types.IsArray:
		str.WriteString("Array, ")
	case types.IsObject:
		str.WriteString("Object(")
		str.WriteString(arg.Object().ClassName())
		str.WriteString("), ")
	}
}
func _buildTraceString(str *zend.SmartStr, ht *types.Array, num uint32) {
	var file *types.Zval
	var tmp *types.Zval
	str.WriteByte('#')
	str.WriteLong(num)
	str.WriteByte(' ')
	file = ht.KeyFind(types.STR_FILE)
	if file != nil {
		if !file.IsString() {
			Error(E_WARNING, "Function name is no string")
			str.WriteString("[unknown function]")
		} else {
			var line zend.ZendLong
			tmp = ht.KeyFind(types.STR_LINE)
			if tmp != nil {
				if tmp.IsLong() {
					line = tmp.Long()
				} else {
					Error(E_WARNING, "Line is no long")
					line = 0
				}
			} else {
				line = 0
			}
			str.WriteString(file.String())
			str.WriteByte('(')
			str.WriteLong(line)
			str.WriteString("): ")
		}
	} else {
		str.WriteString("[internal function]: ")
	}
	str.WriteString(traceAppendKey(ht, types.STR_CLASS))
	str.WriteString(traceAppendKey(ht, types.STR_TYPE))
	str.WriteString(traceAppendKey(ht, types.STR_FUNCTION))
	str.WriteByte('(')
	tmp = ht.KeyFind(types.STR_ARGS)
	if tmp != nil {
		if tmp.IsArray() {
			var last_len int = str.GetS().GetLen()
			tmp.Array().Foreach(func(_ types.ArrayKey, arg *types.Zval) {
				_buildTraceArgs(arg, str)
			})
			if last_len != str.GetS().GetLen() {
				str.GetS().GetLen() -= 2
			}
		} else {
			Error(E_WARNING, "args element is no array")
		}
	}
	str.WriteString(")\n")
}
func zim_exception_getTraceAsString(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var trace *types.Zval
	var rv types.Zval
	var object *types.Zval
	var base_ce *types.ClassEntry
	var str zend.SmartStr = MakeSmartStr(0)
	var num uint32 = 0
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	object = executeData.ThisObjectZval()
	base_ce = GetExceptionBase(object)
	trace = zend.ZendReadProperty(base_ce, object, types.STR_TRACE, 1, &rv)
	if !trace.IsArray() {
		return_value.SetFalse()
		return
	}
	trace.Array().Foreach(func(key types.ArrayKey, frame *types.Zval) {
		if !frame.IsArray() {
			Error(E_WARNING, "Expected array for frame "+zend.ZEND_ULONG_FMT, key.IdxKey())
			return
		}
		_buildTraceString(&str, frame.Array(), lang.PostInc(&num))
	})
	str.WriteByte('#')
	str.WriteLong(num)
	str.WriteString(" {main}")
	str.ZeroTail()
	return_value.SetStringEx(str.GetS())
	return
}
func zim_exception_getPrevious(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var rv types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	types.ZVAL_COPY(return_value, GET_PROPERTY_SILENT(executeData.ThisObjectZval(), types.STR_PREVIOUS, &rv))
}
func zim_exception___toString(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var trace types.Zval
	var exception *types.Zval
	var base_ce *types.ClassEntry
	var str *types.String
	var rv types.Zval
	var tmp types.Zval
	var fname string
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	str = types.NewString("")
	exception = executeData.ThisObjectZval()
	fname = "gettraceasstring"
	for exception != nil && exception.IsObject() && operators.InstanceofFunction(types.Z_OBJCE_P(exception), ZendCeThrowable) {
		var prev_str *types.String = str
		var message *types.String = operators.ZvalGetString(GET_PROPERTY(exception, types.STR_MESSAGE, &rv))
		var file *types.String = operators.ZvalGetString(GET_PROPERTY(exception, types.STR_FILE, &rv))
		var line zend.ZendLong = operators.ZvalGetLong(GET_PROPERTY(exception, types.STR_LINE, &rv))

		fci := types.InitFCallInfo(exception.Object(), &trace)
		fci.SetFunctionName(fname)
		zend.ZendCallFunction(fci, nil)
		if !trace.IsString() {
			trace.SetUndef()
		}
		if (types.Z_OBJCE_P(exception) == ZendCeTypeError || types.Z_OBJCE_P(exception) == ZendCeArgumentCountError) && strstr(message.GetVal(), ", called in ") {
			var real_message *types.String = zend.ZendSprintfZStr("%s and defined", message.GetVal())
			message = real_message
		}
		if message.GetLen() > 0 {
			str = zend.ZendSprintfZStr("%s: %s in %s:"+zend.ZEND_LONG_FMT+"\nStack trace:\n%s%s%s", types.Z_OBJCE_P(exception).Name(), message.GetVal(), file.GetVal(), line, lang.CondF1(trace.IsString() && trace.StringEx().GetLen() != 0, func() []byte { return trace.StringEx().GetVal() }, "#0 {main}\n"), lang.Cond(prev_str.GetLen() != 0, "\n\nNext ", ""), prev_str.GetVal())
		} else {
			str = zend.ZendSprintfZStr("%s in %s:"+zend.ZEND_LONG_FMT+"\nStack trace:\n%s%s%s", types.Z_OBJCE_P(exception).Name(), file.GetVal(), line, lang.CondF1(trace.IsString() && trace.StringEx().GetLen() != 0, func() []byte { return trace.StringEx().GetVal() }, "#0 {main}\n"), lang.Cond(prev_str.GetLen() != 0, "\n\nNext ", ""), prev_str.GetVal())
		}
		exception.Object().ProtectRecursive()
		exception = GET_PROPERTY(exception, types.STR_PREVIOUS, &rv)
		if exception != nil && exception.IsObject() && exception.Object().IsRecursive() {
			break
		}
	}
	// types.ZendStringReleaseEx(fname, 0)
	exception = executeData.ThisObjectZval()

	/* Reset apply counts */

	for exception != nil && exception.IsObject() && lang.Assign(&base_ce, GetExceptionBase(exception)) && operators.InstanceofFunction(types.Z_OBJCE_P(exception), base_ce) != 0 {
		if exception.Object().IsRecursive() {
			exception.Object().UnprotectRecursive()
		} else {
			break
		}
		exception = GET_PROPERTY(exception, types.STR_PREVIOUS, &rv)
	}
	exception = executeData.ThisObjectZval()
	base_ce = GetExceptionBase(exception)

	/* We store the result in the private property string so we can access
	 * the result in uncaught exception handlers without memleaks. */

	tmp.SetStringEx(str)
	zend.ZendUpdatePropertyEx(base_ce, exception, types.STR_STRING, &tmp)
	return_value.SetStringEx(str)
	return
}
func RegisterDefaultException() {
	ZendCeThrowable = zend.RegisterInternalInterface("Throwable", ZendFuncsThrowable)
	ZendCeThrowable.SetInterfaceGetsImplemented(ZendImplementThrowable)

	DefaultExceptionHandlers = *types.NewObjectHandlersEx(zend.StdObjectHandlersPtr, types.ObjectHandlersSetting{
		CloneObj: nil,
	})

	ZendCeException = zend.RegisterClass("Exception", DefaultExceptionNew, DefaultExceptionFunctions)
	zend.ZendClassImplements(ZendCeException, 1, ZendCeThrowable)
	zend.ZendDeclarePropertyString(ZendCeException, "message", b.SizeOf("\"message\"")-1, "", types.AccProtected)
	zend.ZendDeclarePropertyString(ZendCeException, "string", b.SizeOf("\"string\"")-1, "", types.AccPrivate)
	zend.ZendDeclarePropertyLong(ZendCeException, "code", b.SizeOf("\"code\"")-1, 0, types.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeException, "file", b.SizeOf("\"file\"")-1, types.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeException, "line", b.SizeOf("\"line\"")-1, types.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeException, "trace", b.SizeOf("\"trace\"")-1, types.AccPrivate)
	zend.ZendDeclarePropertyNull(ZendCeException, "previous", b.SizeOf("\"previous\"")-1, types.AccPrivate)

	ZendCeErrorException = zend.RegisterSubClass(ZendCeException, "ErrorException", ErrorExceptionNew, ErrorExceptionFunctions)
	zend.ZendDeclarePropertyLong(ZendCeErrorException, "severity", b.SizeOf("\"severity\"")-1, E_ERROR, types.AccProtected)

	ZendCeError = zend.RegisterClass("Error", DefaultExceptionNew, DefaultExceptionFunctions)
	zend.ZendClassImplements(ZendCeError, 1, ZendCeThrowable)
	zend.ZendDeclarePropertyString(ZendCeError, "message", b.SizeOf("\"message\"")-1, "", types.AccProtected)
	zend.ZendDeclarePropertyString(ZendCeError, "string", b.SizeOf("\"string\"")-1, "", types.AccPrivate)
	zend.ZendDeclarePropertyLong(ZendCeError, "code", b.SizeOf("\"code\"")-1, 0, types.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeError, "file", b.SizeOf("\"file\"")-1, types.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeError, "line", b.SizeOf("\"line\"")-1, types.AccProtected)
	zend.ZendDeclarePropertyNull(ZendCeError, "trace", b.SizeOf("\"trace\"")-1, types.AccPrivate)
	zend.ZendDeclarePropertyNull(ZendCeError, "previous", b.SizeOf("\"previous\"")-1, types.AccPrivate)

	ZendCeCompileError = zend.RegisterSubClass(ZendCeError, "CompileError", DefaultExceptionNew, nil)
	ZendCeParseError = zend.RegisterSubClass(ZendCeCompileError, "ParseError", DefaultExceptionNew, nil)
	ZendCeTypeError = zend.RegisterSubClass(ZendCeError, "TypeError", DefaultExceptionNew, nil)
	ZendCeArgumentCountError = zend.RegisterSubClass(ZendCeTypeError, "ArgumentCountError", DefaultExceptionNew, nil)
	ZendCeArithmeticError = zend.RegisterSubClass(ZendCeError, "ArithmeticError", DefaultExceptionNew, nil)
	ZendCeDivisionByZeroError = zend.RegisterSubClass(ZendCeArithmeticError, "DivisionByZeroError", DefaultExceptionNew, nil)
}
func ThrowException(exception_ce *types.ClassEntry, message string, code zend.ZendLong) *types.Object {
	var ex types.Zval
	var tmp types.Zval
	if exception_ce != nil {
		if operators.InstanceofFunction(exception_ce, ZendCeThrowable) == 0 {
			Error(E_NOTICE, "Exceptions must implement Throwable")
			exception_ce = ZendCeException
		}
	} else {
		exception_ce = ZendCeException
	}
	zend.ObjectInitEx(&ex, exception_ce)
	if message {
		tmp.SetString(b.CastStrAuto(message))
		zend.ZendUpdatePropertyEx(exception_ce, &ex, types.STR_MESSAGE, &tmp)
		// zend.ZvalPtrDtor(&tmp)
	}
	if code != 0 {
		tmp.SetLong(code)
		zend.ZendUpdatePropertyEx(exception_ce, &ex, types.STR_CODE, &tmp)
	}
	ThrowExceptionInternal(&ex)
	return ex.Object()
}
func ThrowExceptionEx(exceptionCe *types.ClassEntry, code zend.ZendLong, format string, args ...any) *types.Object {
	message := zend.ZendSprintf(format, args)
	obj := ThrowException(exceptionCe, message, code)
	return obj
}
func ThrowErrorException(exceptionCe *types.ClassEntry, message string, code zend.ZendLong, severity int) *types.Object {
	var ex types.Zval
	var tmp types.Zval
	var obj = ThrowException(exceptionCe, message, code)
	ex.SetObject(obj)
	tmp.SetLong(severity)
	zend.ZendUpdatePropertyEx(ZendCeErrorException, &ex, types.STR_SEVERITY, &tmp)
	return obj
}
func ExceptionError(ex *types.Object, severity int) {
	var exception types.Zval
	var rv types.Zval
	var ce_exception *types.ClassEntry
	exception.SetObject(ex)
	ce_exception = ex.GetCe()
	zend.EG__().SetException(nil)
	if ce_exception == ZendCeParseError || ce_exception == ZendCeCompileError {
		var message *types.String = operators.ZvalGetString(GET_PROPERTY(&exception, types.STR_MESSAGE, &rv))
		var file *types.String = operators.ZvalGetString(GET_PROPERTY_SILENT(&exception, types.STR_FILE, &rv))
		var line zend.ZendLong = operators.ZvalGetLong(GET_PROPERTY_SILENT(&exception, types.STR_LINE, &rv))
		if ce_exception == ZendCeParseError {
			errorCb(E_PARSE, file.GetStr(), uint32(line), message.GetStr())
		} else {
			errorCb(E_COMPILE_ERROR, file.GetStr(), uint32(line), message.GetStr())
		}
		// types.ZendStringReleaseEx(file, 0)
		// types.ZendStringReleaseEx(message, 0)
	} else if operators.InstanceofFunction(ce_exception, ZendCeThrowable) {
		var tmp types.Zval
		var str *types.String
		var file *types.String = nil
		var line zend.ZendLong = 0
		zend.ZendCallMethodWith0Params(&exception, ce_exception, ex.GetCe().GetTostring(), "__tostring", &tmp)
		if zend.EG__().NoException() {
			if !tmp.IsString() {
				Error(E_WARNING, "%s::__toString() must return a string", ce_exception.Name())
			} else {
				zend.ZendUpdatePropertyEx(GetExceptionBase(&exception), &exception, types.STR_STRING, &tmp)
			}
		}
		// zend.ZvalPtrDtor(&tmp)
		if zend.EG__().HasException() {
			var zv types.Zval
			zv.SetObject(zend.EG__().GetException())

			/* do the best we can to inform about the inner exception */

			if operators.InstanceofFunction(ce_exception, ZendCeException) != 0 || operators.InstanceofFunction(ce_exception, ZendCeError) != 0 {
				file = operators.ZvalGetString(GET_PROPERTY_SILENT(&zv, types.STR_FILE, &rv))
				line = operators.ZvalGetLong(GET_PROPERTY_SILENT(&zv, types.STR_LINE, &rv))
			}
			errMsg := fmt.Sprintf("Uncaught %s in exception handling during call to %s::__tostring()", types.Z_OBJCE(zv).Name(), ce_exception.Name())
			if file == nil {
				errorCb(E_WARNING, "", uint32(line), errMsg)
			} else {
				errorCb(E_WARNING, file.GetStr(), uint32(line), errMsg)
			}
			//if file != nil {
			//	types.ZendStringReleaseEx(file, 0)
			//}
		}
		str = operators.ZvalGetString(GET_PROPERTY_SILENT(&exception, types.STR_STRING, &rv))
		file = operators.ZvalGetString(GET_PROPERTY_SILENT(&exception, types.STR_FILE, &rv))
		line = operators.ZvalGetLong(GET_PROPERTY_SILENT(&exception, types.STR_LINE, &rv))
		if file == nil {
			errorCb(severity, "", uint32(line), fmt.Sprintf("Uncaught %s\n  thrown", str.GetStr()))
		} else {
			errorCb(severity, file.GetStr(), uint32(line), fmt.Sprintf("Uncaught %s\n  thrown", str.GetStr()))
		}
		// types.ZendStringReleaseEx(str, 0)
		// types.ZendStringReleaseEx(file, 0)
	} else {
		Error(severity, "Uncaught exception '%s'", ce_exception.Name())
	}
	// zend.OBJ_RELEASE(ex)
}
func ThrowExceptionObject(exception *types.Zval) {
	var exceptionCe *types.ClassEntry
	if exception == nil || !exception.IsObject() {
		ErrorNoreturn(E_CORE_ERROR, "Need to supply an object when throwing an exception")
	}
	exceptionCe = types.Z_OBJCE_P(exception)
	if exceptionCe == nil || !operators.InstanceofFunction(exceptionCe, ZendCeThrowable) {
		ThrowError(nil, "Cannot throw objects that do not implement Throwable")
		return
	}
	ThrowExceptionInternal(exception)
}
