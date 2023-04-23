package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ASSERTG(v __auto__) __auto__ { return AssertGlobals.v }
func SAFE_STRING(s *byte) string {
	if s != nil {
		return s
	} else {
		return ""
	}
}
func OnChangeCallback(
	entry *zend.ZendIniEntry,
	new_value *types.String,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	if zend.CurrEX() != nil {
		if ASSERTG(callback).IsNotUndef() {
			// zend.ZvalPtrDtor(&(ASSERTG(callback)))
			ASSERTG(callback).SetUndef()
		}
		if new_value != nil && (ASSERTG(callback).IsNotUndef() || new_value.GetLen() != 0) {
			ASSERTG(callback).SetStringCopy(new_value)
		}
	} else {
		if ASSERTG(cb) {
			zend.Pefree(ASSERTG(cb), 1)
		}
		if new_value != nil && new_value.GetLen() != 0 {
			ASSERTG(cb) = zend.Pemalloc(new_value.GetLen() + 1)
			memcpy(ASSERTG(cb), new_value.GetVal(), new_value.GetLen())
			ASSERTG(cb)[new_value.GetLen()] = '0'
		} else {
			ASSERTG(cb) = nil
		}
	}
	return types.SUCCESS
}
func PhpAssertInitGlobals(assert_globals_p *ZendAssertGlobals) {
	assert_globals_p.GetCallback().SetUndef()
	assert_globals_p.SetCb(nil)
}
func ZmStartupAssert(type_ int, module_number int) int {
	var ce types.ClassEntry
	PhpAssertInitGlobals(&AssertGlobals)
	zend.REGISTER_INI_ENTRIES(module_number)
	zend.RegisterLongConstant("ASSERT_ACTIVE", ASSERT_ACTIVE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("ASSERT_CALLBACK", ASSERT_CALLBACK, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("ASSERT_BAIL", ASSERT_BAIL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("ASSERT_WARNING", ASSERT_WARNING, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("ASSERT_QUIET_EVAL", ASSERT_QUIET_EVAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("ASSERT_EXCEPTION", ASSERT_EXCEPTION, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.NewString("AssertionError"))
	ce.SetBuiltinFunctions(nil)
	AssertionErrorCe = zend.ZendRegisterInternalClassEx(&ce, faults.ZendCeError)
	return types.SUCCESS
}
func ZmShutdownAssert(type_ int, module_number int) int {
	if ASSERTG(cb) {
		zend.Pefree(ASSERTG(cb), 1)
		ASSERTG(cb) = nil
	}
	return types.SUCCESS
}
func ZmDeactivateAssert(type_ int, module_number int) int {
	if ASSERTG(callback).IsNotUndef() {
		// zend.ZvalPtrDtor(&(ASSERTG(callback)))
		ASSERTG(callback).SetUndef()
	}
	return types.SUCCESS
}
func ZmInfoAssert(zend_module *zend.ModuleEntry) { zend.DISPLAY_INI_ENTRIES() }
func ZifAssert(executeData zpp.Ex, return_value zpp.Ret, assertion *types.Zval, _ zpp.Opt, description *types.Zval) {
	var assertion *types.Zval
	var description *types.Zval = nil
	var val int
	var myeval *byte = nil
	var compiled_string_description *byte
	if !(ASSERTG(active)) {
		return_value.SetTrue()
		return
	}
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			assertion = fp.ParseZval()
			fp.StartOptional()
			description = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if assertion.IsString() {
		var retval types.Zval
		var old_error_reporting int = 0
		if zend.ZendForbidDynamicCall("assert() with string argument") == types.FAILURE {
			return_value.SetFalse()
			return
		}
		core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Calling assert() with a string argument is deprecated")
		myeval = assertion.String().GetVal()
		if ASSERTG(quiet_eval) {
			old_error_reporting = zend.EG__().GetErrorReporting()
			zend.EG__().SetErrorReporting(0)
		}
		compiled_string_description = zend.ZendMakeCompiledStringDescription("assert code")
		if zend.ZendEvalStringl(myeval, assertion.String().GetLen(), &retval, compiled_string_description) == types.FAILURE {
			zend.Efree(compiled_string_description)
			if description == nil {
				faults.ThrowError(nil, "Failure evaluating code: %s%s", core.PHP_EOL, myeval)
			} else {
				var str *types.String = zend.ZvalGetString(description)
				faults.ThrowError(nil, "Failure evaluating code: %s%s:\"%s\"", core.PHP_EOL, str.GetVal(), myeval)
				// types.ZendStringReleaseEx(str, 0)
			}
			if ASSERTG(bail) {
				faults.Bailout()
			}
			return_value.SetFalse()
			return
		}
		zend.Efree(compiled_string_description)
		if ASSERTG(quiet_eval) {
			zend.EG__().SetErrorReporting(old_error_reporting)
		}
		zend.ConvertToBoolean(&retval)
		val = retval.IsType(types.IS_TRUE)
	} else {
		val = zend.ZendIsTrue(assertion)
	}
	if val != 0 {
		return_value.SetTrue()
		return
	}
	if ASSERTG(callback).IsUndef() && ASSERTG(cb) {
		(ASSERTG(callback)).SetStringVal(b.CastStrAuto(ASSERTG(cb)))
	}
	if ASSERTG(callback).IsNotUndef() {
		var args []types.Zval
		var retval types.Zval
		var lineno uint32 = zend.ZendGetExecutedLineno()
		var filename *byte = zend.ZendGetExecutedFilename()
		args[0].SetStringVal(b.CastStrAuto(SAFE_STRING(filename)))
		args[1].SetLong(lineno)
		args[2].SetStringVal(b.CastStrAuto(SAFE_STRING(myeval)))
		retval.SetFalse()

		/* XXX do we want to check for error here? */

		if description == nil {
			zend.CallUserFunction(nil, &(ASSERTG(callback)), &retval, 3, args)
			// zend.ZvalPtrDtor(&args[2])
			// zend.ZvalPtrDtor(&args[0])
		} else {
			args[3].SetString(zend.ZvalGetString(description))
			zend.CallUserFunction(nil, &(ASSERTG(callback)), &retval, 4, args)
			// zend.ZvalPtrDtor(&args[3])
			// zend.ZvalPtrDtor(&args[2])
			// zend.ZvalPtrDtor(&args[0])
		}
		// zend.ZvalPtrDtor(&retval)
	}
	if ASSERTG(exception) {
		if description == nil {
			faults.ThrowException(AssertionErrorCe, nil, faults.E_ERROR)
		} else if description.IsType(types.IS_OBJECT) && zend.InstanceofFunction(types.Z_OBJCE_P(description), faults.ZendCeThrowable) != 0 {
			// 			description.AddRefcount()
			faults.ThrowExceptionObject(description)
		} else {
			var str *types.String = zend.ZvalGetString(description)
			faults.ThrowException(AssertionErrorCe, str.GetVal(), faults.E_ERROR)
			// types.ZendStringReleaseEx(str, 0)
		}
	} else if ASSERTG(warning) {
		if description == nil {
			if myeval != nil {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Assertion \"%s\" failed", myeval)
			} else {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Assertion failed")
			}
		} else {
			var str *types.String = zend.ZvalGetString(description)
			if myeval != nil {
				core.PhpErrorDocref(nil, faults.E_WARNING, "%s: \"%s\" failed", str.GetVal(), myeval)
			} else {
				core.PhpErrorDocref(nil, faults.E_WARNING, "%s failed", str.GetVal())
			}
			// types.ZendStringReleaseEx(str, 0)
		}
	}
	if ASSERTG(bail) {
		faults.Bailout()
	}
	return_value.SetFalse()
	return
}
func ZifAssertOptions(executeData zpp.Ex, return_value zpp.Ret, what *types.Zval, _ zpp.Opt, value *types.Zval) {
	var value *types.Zval = nil
	var what zend.ZendLong
	var oldint types.ZendBool
	var ac int = executeData.NumArgs()
	var key *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			what = fp.ParseLong()
			fp.StartOptional()
			value = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	switch what {
	case ASSERT_ACTIVE:
		oldint = ASSERTG(active)
		if ac == 2 {
			var value_str *types.String = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = types.NewString("assert.active")
			zend.ZendAlterIniEntryEx(key.GetStr(), value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			// types.ZendStringReleaseEx(key, 0)
			// types.ZendStringReleaseEx(value_str, 0)
		}
		return_value.SetLong(oldint)
		return
	case ASSERT_BAIL:
		oldint = ASSERTG(bail)
		if ac == 2 {
			var value_str *types.String = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = types.NewString("assert.bail")
			zend.ZendAlterIniEntryEx(key.GetStr(), value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			// types.ZendStringReleaseEx(key, 0)
			// types.ZendStringReleaseEx(value_str, 0)
		}
		return_value.SetLong(oldint)
		return
	case ASSERT_QUIET_EVAL:
		oldint = ASSERTG(quiet_eval)
		if ac == 2 {
			var value_str *types.String = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = types.NewString("assert.quiet_eval")
			zend.ZendAlterIniEntryEx(key.GetStr(), value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			// types.ZendStringReleaseEx(key, 0)
			// types.ZendStringReleaseEx(value_str, 0)
		}
		return_value.SetLong(oldint)
		return
	case ASSERT_WARNING:
		oldint = ASSERTG(warning)
		if ac == 2 {
			var value_str *types.String = zend.ZvalTryGetString(value)
			if value_str == nil {
				return
			}
			key = types.NewString("assert.warning")
			zend.ZendAlterIniEntryEx(key.GetStr(), value_str, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			// types.ZendStringReleaseEx(key, 0)
			// types.ZendStringReleaseEx(value_str, 0)
		}
		return_value.SetLong(oldint)
		return
	case ASSERT_CALLBACK:
		if ASSERTG(callback).IsNotUndef() {
			types.ZVAL_COPY(return_value, &(ASSERTG(callback)))
		} else if ASSERTG(cb) {
			return_value.SetStringVal(b.CastStrAuto(ASSERTG(cb)))
		} else {
			return_value.SetNull()
		}
		if ac == 2 {
			// zend.ZvalPtrDtor(&(ASSERTG(callback)))
			types.ZVAL_COPY(&(ASSERTG(callback)), value)
		}
		return
	case ASSERT_EXCEPTION:
		oldint = ASSERTG(exception)
		if ac == 2 {
			var val *types.String = zend.ZvalTryGetString(value)
			if val == nil {
				return
			}
			key = types.NewString("assert.exception")
			zend.ZendAlterIniEntryEx(key.GetStr(), val, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0)
			// types.ZendStringReleaseEx(val, 0)
			// types.ZendStringReleaseEx(key, 0)
		}
		return_value.SetLong(oldint)
		return
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unknown value "+zend.ZEND_LONG_FMT, what)
	}
	return_value.SetFalse()
	return
}
