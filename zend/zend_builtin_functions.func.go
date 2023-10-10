package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/globals"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func ZmStartupCore(type_ int, module_number int) int {
	ZendStandardClassDef = RegisterClass("stdClass", nil, nil)
	ZendRegisterDefaultClasses()
	return types.SUCCESS
}
func NewStdClassObject(properties *types.Array) *types.Object {
	obj := types.NewStdObject(ZendStandardClassDef)
	obj.SetProperties(properties)
	return obj
}

func ZendStartupBuiltinFunctions() {
	module := ZendRegisterModuleEx(&ZendBuiltinModule)
	EG__().SetCurrentModule(module)
}

func ZifZendVersion() string  { return ZEND_VERSION }
func ZifGcMemCaches() int     { return 0 }
func ZifGcCollectCycles() int { return 0 }
func ZifGcEnabled() bool      { return true }
func ZifGcEnable() {
	ZendAlterIniEntryChars("zend.enable_gc", "1", ZEND_INI_USER, ZEND_INI_STAGE_RUNTIME)
}
func ZifGcDisable() {
	ZendAlterIniEntryChars("zend.enable_gc", "0", ZEND_INI_USER, ZEND_INI_STAGE_RUNTIME)
}

func ZifGcStatus(ret zpp.Ret) {
	ArrayInitSize(ret, 3)
	AddAssocLongEx(ret, "runs", 0)
	AddAssocLongEx(ret, "collected", 0)
	AddAssocLongEx(ret, "threshold", 0)
	AddAssocLongEx(ret, "roots", 0)
}

func ZifFuncNumArgs(executeData zpp.Ex) int {
	var ex = executeData.GetPrevExecuteData()
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_CODE) != 0 {
		faults.Error(faults.E_WARNING, "func_num_args():  Called from the global scope - no function context")
		return -1
	}
	if ZendForbidDynamicCall("func_num_args()") == types.FAILURE {
		return -1
	}
	return ex.NumArgs()
}

//@zif -old
func ZifFuncGetArg(executeData zpp.Ex, returnValue zpp.Ret, argNum int) {
	if argNum < 0 {
		faults.Error(faults.E_WARNING, "func_get_arg():  The argument number should be >= 0")
		returnValue.SetFalse()
		return
	}

	ex := executeData.GetPrevExecuteData()
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_CODE) != 0 {
		faults.Error(faults.E_WARNING, "func_get_arg():  Called from the global scope - no function context")
		returnValue.SetFalse()
		return
	}
	if ZendForbidDynamicCall("func_get_arg()") == types.FAILURE {
		returnValue.SetFalse()
		return
	}

	argCount := ex.NumArgs()
	if argNum >= argCount {
		faults.Error(faults.E_WARNING, "func_get_arg():  Argument "+ZEND_LONG_FMT+" not passed to function", argNum)
		returnValue.SetFalse()
		return
	}

	var arg *types.Zval
	firstExtraArg := int(ex.GetFunc().GetOpArray().GetNumArgs())
	if argNum >= firstExtraArg && ex.NumArgs() > firstExtraArg {
		arg = ex.VarNum(ex.GetFunc().GetOpArray().GetLastVar() + int(ex.GetFunc().GetOpArray().GetT()) + (argNum - firstExtraArg))
	} else {
		arg = ex.Arg(argNum + 1)
	}

	if !(arg.IsUndef()) {
		types.ZVAL_COPY_DEREF(returnValue, arg)
	}
}
func ZifFuncGetArgs(executeData zpp.Ex) (*types.Array, bool) {
	ex := executeData.GetPrevExecuteData()
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_CODE) != 0 {
		faults.Error(faults.E_WARNING, "func_get_args():  Called from the global scope - no function context")
		return nil, false
	}
	if ZendForbidDynamicCall("func_get_args()") == types.FAILURE {
		return nil, false
	}

	argCount := ex.NumArgs()
	if argCount == 0 {
		return types.NewArray(), true
	}

	firstExtraArg := int(ex.GetFunc().GetOpArray().GetNumArgs())
	var values []*types.Zval
	if argCount <= firstExtraArg {
		values = executeData.Args(0, argCount)
	} else {
		values = executeData.Args(0, firstExtraArg)

		for i := 0; i < argCount-firstExtraArg; i++ {
			p := ex.VarNum(ex.GetFunc().GetOpArray().GetLastVar() + int(ex.GetFunc().GetOpArray().GetT()) + i)
			values = append(values, p)
		}
	}

	arr := types.NewArrayCap(argCount)
	for _, zv := range values {
		if zv.IsUndef() {
			arr.AppendNew(types.NewZvalNull())
		} else {
			zv = types.ZVAL_DEREF(zv)

			// zv.TryAddRefcount()

			arr.AppendNew(zv)
		}
	}
	return arr, true
}
func ZifStrlen(str string) int               { return len(str) }
func ZifStrcmp(str1 string, str2 string) int { return strings.Compare(str1, str2) }
func ZifStrncmp(str1 string, str2 string, len_ int) (int, bool) {
	if len_ < 0 {
		faults.Error(faults.E_WARNING, "Length must be greater than or equal to 0")
		return 0, false
	}
	return operators.ZendBinaryStrncmp(str1, str2, len_), true
}
func ZifStrcasecmp(str1 string, str2 string) int { return operators.ZendBinaryStrcasecmp(str1, str2) }
func ZifStrncasecmp(str1 string, str2 string, len_ int) (int, bool) {
	if len_ < 0 {
		faults.Error(faults.E_WARNING, "Length must be greater than or equal to 0")
		return 0, false
	}
	return operators.ZendBinaryStrncasecmp(str1, str2, len_), true
}

//@zif -old "z/"
func ZifEach(array zpp.RefZval) (*types.Array, bool) {
	if EG__().GetEachDeprecationThrown() == 0 {
		faults.Error(faults.E_DEPRECATED, "The each() function is deprecated. This message will be suppressed on further calls")
		EG__().SetEachDeprecationThrown(1)
	}
	targetHash := HASH_OF(array)
	if targetHash == nil {
		faults.Error(faults.E_WARNING, "Variable passed to each() is not an array or object")
		return nil, false
	}

	pair := targetHash.Current()
	if pair == nil {
		return nil, false
	}
	key := pair.GetKey()
	val := pair.GetVal().DeRef()

	result := types.NewArrayCap(4)

	/* add value elements */
	result.IndexAddNew(1, val)
	result.KeyAddNew(types.STR_VALUE, val)

	/* add the key elements */
	var tmp types.Zval
	if key.IsStrKey() {
		tmp.SetString(key.StrKey())
	} else {
		tmp.SetLong(key.IdxKey())
	}

	result.IndexAddNew(0, &tmp)
	result.KeyAddNew(types.STR_KEY, &tmp)

	result.MoveNext()

	return result, true
}
func ZifErrorReporting(ret zpp.Ret, _ zpp.Opt, newErrorLevel *types.Zval) {
	var old_error_reporting int
	old_error_reporting = EG__().GetErrorReporting()
	if newErrorLevel != nil {
		var new_val = operators.ZvalTryGetString(newErrorLevel)
		if new_val == nil {
			return
		}
		for {
			var p = EG__().GetErrorReportingIniEntry()
			if p == nil {
				var iniEntry = EG__().IniDirectives().Get(types.STR_ERROR_REPORTING)
				if iniEntry != nil {
					EG__().SetErrorReportingIniEntry(iniEntry)
					p = iniEntry
				} else {
					break
				}
			}
			if p.GetModified() == 0 {
				if EG__().ModifiedIniDirectives() == nil {
					EG__().InitModifiedIniDirectives()
				}
				if EG__().ModifiedIniDirectives().Add(types.STR_ERROR_REPORTING, p) {
					p.SetOrigValue(p.GetValue())
					p.SetOrigModifiable(p.GetModifiable())
					p.SetModified(1)
				}
			} else if p.GetOrigValue() != p.GetValue() {
				// types.ZendStringReleaseEx(p.GetValue(), 0)
			}
			p.SetValue(new_val)
			if newErrorLevel.IsLong() {
				EG__().SetErrorReporting(newErrorLevel.Long())
			} else {
				EG__().SetErrorReporting(atoi(p.GetValue().GetVal()))
			}
			break
		}
	}
	ret.SetLong(old_error_reporting)
}
func ValidateConstantArray(ht *types.Array) int {
	var ret = 1
	ht.ProtectRecursive()
	ht.ForeachIndirectEx(func(_ types.ArrayKey, val *types.Zval) bool {
		val = types.ZVAL_DEREF(val)
		if val.IsRefcounted() {
			if val.IsArray() {
				if val.Array().IsRecursive() {
					faults.Error(faults.E_WARNING, "Constants cannot be recursive arrays")
					ret = 0
					return false
				} else if ValidateConstantArray(val.Array()) == 0 {
					ret = 0
					return false
				}
			} else if !val.IsString() && !val.IsResource() {
				faults.Error(faults.E_WARNING, "Constants may only evaluate to scalar values, arrays or resources")
				ret = 0
				return false
			}
		}
		return true
	})
	ht.UnprotectRecursive()
	return ret
}
func CopyConstantArray(dst *types.Zval, src *types.Zval) {
	dstArr := types.NewArrayCap(src.Array().Len())
	src.Array().Foreach(func(key types.ArrayKey, val *types.Zval) {
		/* constant arrays can't contain references */
		val = types.ZVAL_DEREF(val)
		newVal := dstArr.Add(key, val)
		if val.IsArray() {
			if val.IsRefcounted() {
				CopyConstantArray(newVal, val)
			}
		}
	})
	dst.SetArray(dstArr)
}

func ZifDefine(constantName string, value *types.Zval, _ zpp.Opt, caseInsensitive bool) bool {
	var val_free types.Zval
	var caseSensitive = CONST_CS
	var c ZendConstant
	if caseInsensitive {
		caseSensitive = 0
	}
	if strings.Contains(constantName, "::") {
		faults.Error(faults.E_WARNING, "Class constants cannot be defined or redefined")
		return false
	}
	val_free.SetUndef()
repeat:
	switch value.Type() {
	case types.IsNull,
		types.IsFalse,
		types.IsTrue,
		types.IsLong,
		types.IsDouble,
		types.IsString,
		types.IsResource:
		// pass
	case types.IsArray:
		if value.IsRefcounted() {
			if ValidateConstantArray(value.Array()) == 0 {
				return false
			} else {
				CopyConstantArray(c.Value(), value)
				goto register_constant
			}
		}
	case types.IsObject:
		if val_free.IsUndef() {
			if value.Object().CanGet() {
				value = value.Object().Get(&val_free)
				goto repeat
			} else if value.Object().CanCast() {
				if value.Object().Cast(&val_free, types.IsString) == types.SUCCESS {
					value = &val_free
					break
				}
			}
		}
		fallthrough
	default:
		faults.Error(faults.E_WARNING, "Constants may only evaluate to scalar values, arrays or resources")
		// ZvalPtrDtor(&val_free)
		return false
	}
	types.ZVAL_COPY(c.Value(), value)
	// ZvalPtrDtor(&val_free)
register_constant:
	if caseInsensitive {
		faults.Error(faults.E_DEPRECATED, "define(): Declaration of case-insensitive constants is deprecated")
	}

	/* non persistent */

	c.SetFlags(caseSensitive, PHP_USER_CONSTANT)
	c.SetName(constantName)
	return ZendRegisterConstant(&c)
}
func ZifDefined(constantName string) bool {
	if ZendGetConstantEx(constantName, ZendGetExecutedScope(), ZEND_FETCH_CLASS_SILENT|ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK) != nil {
		return true
	} else {
		return false
	}
}
func ZifGetClass(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, object *types.Zval) {
	var obj *types.Zval = nil
	if ZendParseParameters(executeData.NumArgs(), "|o", &obj) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	if obj == nil {
		var scope = ZendGetExecutedScope()
		if scope != nil {
			return_value.SetString(scope.Name())
			return
		} else {
			faults.Error(faults.E_WARNING, "get_class() called without object from outside a class")
			return_value.SetFalse()
			return
		}
	}
	return_value.SetString(obj.Object().GetCe().Name())
	return
}
func ZifGetCalledClass(ex zpp.Ex) (string, bool) {
	var called_scope *types.ClassEntry
	called_scope = ZendGetCalledScope(ex)
	if called_scope != nil {
		return called_scope.Name(), true
	} else {
		var scope = ZendGetExecutedScope()
		if scope == nil {
			faults.Error(faults.E_WARNING, "get_called_class() called from outside a class")
		}
	}
	return "", false
}
func ZifGetParentClass(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, object *types.Zval) (string, bool) {
	var ce *types.ClassEntry = nil
	if executeData.NumArgs() == 0 {
		ce = ZendGetExecutedScope()
	} else {
		if object.IsObject() {
			ce = object.Object().GetCe()
		} else if object.IsString() {
			ce = ZendLookupClassString(object.String())
		}
	}
	if ce != nil && ce.GetParent() != nil {
		return ce.GetParent().Name(), true
	} else {
		return_value.SetFalse()
		return "", false
	}
}
func IsAImpl(executeData *ZendExecuteData, return_value *types.Zval, only_subclass bool) {
	var obj *types.Zval
	var class_name *types.String
	var instance_ce *types.ClassEntry
	var ce *types.ClassEntry
	var allow_string = only_subclass
	var retval bool
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			obj = fp.ParseZval()
			class_name = fp.ParseStr()
			fp.StartOptional()
			allow_string = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/*
	 * allow_string - is_a default is no, is_subclass_of is yes.
	 *   if it's allowed, then the autoloader will be called if the class does not exist.
	 *   default behaviour is different, as 'is_a' used to be used to test mixed return values
	 *   and there is no easy way to deprecate this.
	 */

	if allow_string != 0 && obj.IsString() {
		instance_ce = ZendLookupClass(obj.StringEx())
		if instance_ce == nil {
			return_value.SetFalse()
			return
		}
	} else if obj.IsObject() {
		instance_ce = types.Z_OBJCE_P(obj)
	} else {
		return_value.SetFalse()
		return
	}
	if only_subclass == 0 && instance_ce.Name() == class_name.GetStr() {
		retval = 1
	} else {
		ce = ZendLookupClassEx(class_name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if ce == nil {
			retval = 0
		} else {
			if only_subclass != 0 && instance_ce == ce {
				retval = 0
			} else {
				retval = operators.InstanceofFunction(instance_ce, ce)
			}
		}
	}
	return_value.SetBool(retval != 0)
	return
}
func ZifIsSubclassOf(executeData zpp.Ex, return_value zpp.Ret, object *types.Zval, className *types.Zval, _ zpp.Opt, allowString *types.Zval) {
	IsAImpl(executeData, return_value, 1)
}
func ZifIsA(executeData zpp.Ex, return_value zpp.Ret, object *types.Zval, className *types.Zval, _ zpp.Opt, allowString *types.Zval) {
	IsAImpl(executeData, return_value, 0)
}
func AddClassVars(scope *types.ClassEntry, ce *types.ClassEntry, statics int, return_value *types.Zval) {
	ce.PropertyTable().ForeachEx(func(key string, propInfo *types.PropertyInfo) bool {
		if propInfo.IsProtected() && !ZendCheckProtected(propInfo.GetCe(), scope) || propInfo.IsPrivate() && propInfo.GetCe() != scope {
			return true
		}
		var prop *types.Zval = nil
		if statics != 0 && propInfo.IsStatic() {
			prop = ce.GetDefaultStaticMembersTable()[propInfo.GetOffset()]
			prop = types.ZVAL_DEINDIRECT(prop)
		} else if statics == 0 && !propInfo.IsStatic() {
			prop = ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(propInfo.GetOffset())]
		}
		if prop == nil {
			return true
		}

		var prop_copy types.Zval
		if prop.IsUndef() {
			/* Return uninitialized typed properties as a null value */
			prop_copy.SetNull()
		} else {
			/* copy: enforce read only access */
			types.ZVAL_COPY_OR_DUP(&prop_copy, prop)

		}
		prop = &prop_copy

		/* this is necessary to make it able to work with default array
		 * properties, returned to user */

		if prop.IsConstantAst() {
			if ZvalUpdateConstantEx(prop, nil) != types.SUCCESS {
				return false
			}
		}
		return_value.Array().KeyAddNew(key, prop)
		return true
	})
}
func ZifGetClassVars(executeData zpp.Ex, return_value zpp.Ret, className *types.Zval) {
	var class_name *types.String
	var ce *types.ClassEntry
	var scope *types.ClassEntry
	if ZendParseParameters(executeData.NumArgs(), "S", &class_name) == types.FAILURE {
		return
	}
	ce = ZendLookupClass(class_name)
	if ce == nil {
		return_value.SetFalse()
		return
	} else {
		ArrayInit(return_value)
		if !ce.IsConstantsUpdated() {
			if ZendUpdateClassConstants(ce) != types.SUCCESS {
				return
			}
		}
		scope = ZendGetExecutedScope()
		AddClassVars(scope, ce, 0, return_value)
		AddClassVars(scope, ce, 1, return_value)
	}
}
func ZifGetObjectVars(obj zpp.Object) (*types.Array, bool) {
	properties := obj.Object().GetPropertiesArray()
	if properties == nil {
		return nil, false
	}

	zobj := obj.Object()
	if zobj.GetCe().GetDefaultPropertiesCount() == 0 && properties == zobj.GetProperties() && !(properties.IsRecursive()) {
		/* fast copy */
		if zobj.GetHandlers() == StdObjectHandlersPtr {
			return types.ZendProptableToSymtable(properties, 0), true
		}
		return types.ZendProptableToSymtable(properties, 1), true
	} else {
		retArr := types.NewArrayCap(properties.Len())
		properties.Foreach(func(key types.ArrayKey, value *types.Zval) {
			var isDynamic = true
			if value.IsIndirect() {
				value = value.Indirect()
				if value.IsUndef() {
					return
				}
				isDynamic = false
			}

			if key.IsStrKey() && ZendCheckPropertyAccess(zobj, key.StrKey(), isDynamic) == types.FAILURE {
				return
			}

			if !key.IsStrKey() {
				/* This case is only possible due to loopholes, e.g. ArrayObject */
				retArr.IndexAdd(key.IdxKey(), value)
			} else if !isDynamic && key.StrKey()[0] == 0 {
				/* We assume here that a mangled property name is never
				 * numeric. This is probably a safe assumption, but
				 * theoretically someone might write an extension with
				 * private, numeric properties. Well, too bad.
				 */

				_, propName, _ := ZendUnmanglePropertyName_Ex(key.StrKey())
				retArr.KeyAddNew(propName, value)
			} else {
				retArr.SymtableAddNew(key.StrKey(), value)
			}
		})
		return retArr, true
	}
}
func ZifGetMangledObjectVars(executeData zpp.Ex, return_value zpp.Ret, obj *types.Zval) {
	var obj *types.Zval
	var properties *types.Array
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			obj = fp.ParseObject()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	properties = obj.Object().GetPropertiesArray()
	if properties == nil {
		return_value.SetEmptyArray()
		return
	}
	properties = types.ZendProptableToSymtable(properties, types.Z_OBJCE_P(obj).GetDefaultPropertiesCount() != 0 || obj.Object().GetHandlers() != StdObjectHandlersPtr || properties.IsRecursive())
	return_value.SetArray(properties)
	return
}
func SameNameEx(key string, name string) bool {
	return key == ascii.StrToLower(name)
}
func ZifGetClassMethods(class *types.Zval) *types.Zval {
	var ce *types.ClassEntry = nil
	if class.IsObject() {
		ce = class.Object().GetCe()
	} else if class.IsString() {
		ce = ZendLookupClass(class.StringEx())
	}
	if ce == nil {
		return types.NewZvalNull()
	}

	var names []string
	scope := ZendGetExecutedScope()
	ce.FunctionTable().Foreach(func(key string, mptr types.IFunction) {
		if mptr.IsPublic() || scope != nil && (mptr.IsProtected() && ZendCheckProtected(mptr.GetScope(), scope) || mptr.IsPrivate() && scope == mptr.GetScope()) {
			if mptr.GetType() == ZEND_USER_FUNCTION && (mptr.GetOpArray().GetRefcount() == nil || *mptr.GetOpArray().GetRefcount() > 1) && key != "" && !SameNameEx(key, mptr.FunctionName()) {
				names = append(names, ZendFindAliasName(mptr.GetScope(), key))
			} else {
				names = append(names, mptr.FunctionName())
			}
		}
	})
	return types.NewZvalArray(types.NewArrayOfString(names))
}
func ZifMethodExists(object *types.Zval, method string) bool {
	var ce *types.ClassEntry
	if object.IsObject() {
		ce = object.Object().GetCe()
	} else if object.IsString() {
		ce = ZendLookupClassString(object.String())
	}
	if ce == nil {
		return false
	}

	func_ := ce.FunctionTable().Get(method)
	if func_ != nil {
		/* Exclude shadow properties when checking a method on a specific class. Include
		 * them when checking an object, as method_exists() generally ignores visibility.
		 * TODO: Should we use EG(scope) for the object case instead? */
		return object.IsObject() || !func_.IsPrivate() || func_.GetScope() == ce
	}
	if object.IsObject() {
		func_ = object.Object().GetMethod(method, nil)
		if func_ != nil {
			if func_.IsCallViaTrampoline() {
				/* Returns true to the fake Closure's __invoke */
				ret := func_.GetScope() == ZendCeClosure && method == ZEND_INVOKE_FUNC_NAME
				ZendFreeTrampoline(func_)
				return ret
			}
			return true
		}
	}
	return false
}
func ZifPropertyExists(executeData zpp.Ex, return_value zpp.Ret, objectOrClass *types.Zval, propertyName *types.Zval) {
	var object *types.Zval
	var property *types.String
	var ce *types.ClassEntry
	var property_info *types.PropertyInfo
	var property_z types.Zval
	if ZendParseParameters(executeData.NumArgs(), "zS", &object, &property) == types.FAILURE {
		return
	}
	if property == nil {
		return_value.SetFalse()
		return
	}
	if object.IsString() {
		ce = ZendLookupClass(object.StringEx())
		if ce == nil {
			return_value.SetFalse()
			return
		}
	} else if object.IsObject() {
		ce = types.Z_OBJCE_P(object)
	} else {
		faults.Error(faults.E_WARNING, "First parameter must either be an object or the name of an existing class")
		return_value.SetNull()
		return
	}
	property_info = ce.PropertyTable().Get(property.GetStr())
	if property_info != nil && (!property_info.IsPrivate() || property_info.GetCe() == ce) {
		return_value.SetTrue()
		return
	}
	property_z.SetStringEx(property)
	if object.IsObject() && object.Object().HasPropertyEx(&property_z, 2) {
		return_value.SetTrue()
		return
	}
	return_value.SetFalse()
	return
}
func ClassExistsImpl(executeData *ZendExecuteData, return_value *types.Zval, flags int, skip_flags int) {
	var name *types.String
	var lcname string
	var ce *types.ClassEntry
	var autoload = 1
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			name = fp.ParseStr()
			fp.StartOptional()
			autoload = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if autoload == 0 {
		if name.GetStr()[0] == '\\' {
			/* Ignore leading "\" */
			lcname = ascii.StrToLower(name.GetStr()[1:])
		} else {
			lcname = ascii.StrToLower(name.GetStr())
		}
		ce = EG__().ClassTable().Get(lcname)
	} else {
		ce = ZendLookupClass(name)
	}
	if ce != nil {
		return_value.SetBool((ce.GetCeFlags()&flags) == flags && !ce.HasCeFlags(skip_flags))
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifClassExists(executeData zpp.Ex, return_value zpp.Ret, classname *types.Zval, _ zpp.Opt, autoload *types.Zval) {
	ClassExistsImpl(executeData, return_value, types.AccLinked, types.AccInterface|types.AccTrait)
}
func ZifInterfaceExists(executeData zpp.Ex, return_value zpp.Ret, classname *types.Zval, _ zpp.Opt, autoload *types.Zval) {
	ClassExistsImpl(executeData, return_value, types.AccLinked|types.AccInterface, 0)
}
func ZifTraitExists(executeData zpp.Ex, return_value zpp.Ret, traitname *types.Zval, _ zpp.Opt, autoload *types.Zval) {
	ClassExistsImpl(executeData, return_value, types.AccTrait, 0)
}
func ZifFunctionExists(functionName string) bool {
	var func_ types.IFunction
	var lcname string
	if functionName[0] == '\\' {
		lcname = ascii.StrToLower(functionName[1:])
	} else {
		lcname = ascii.StrToLower(functionName)
	}

	func_ = EG__().FunctionTable().Get(lcname)

	/*
	 * A bit of a hack, but not a bad one: we see if the handler of the function
	 * is actually one that displays "function is disabled" message.
	 */
	return func_ != nil && (func_.GetType() != ZEND_INTERNAL_FUNCTION || func_.GetInternalFunction().GetHandler() != ZifDisplayDisabledFunction)
}
func ZifClassAlias(executeData zpp.Ex, return_value zpp.Ret, userClassName *types.Zval, aliasName *types.Zval, _ zpp.Opt, autoload *types.Zval) {
	var class_name *types.String
	var alias_name *byte
	var ce *types.ClassEntry
	var alias_name_len int
	var autoload = 1
	if ZendParseParameters(executeData.NumArgs(), "Ss|b", &class_name, &alias_name, &alias_name_len, &autoload) == types.FAILURE {
		return
	}
	ce = ZendLookupClassEx(class_name, nil, lang.Cond(autoload == 0, ZEND_FETCH_CLASS_NO_AUTOLOAD, 0))
	if ce != nil {
		if ce.IsUserClass() {
			if ZendRegisterClassAliasEx(b.CastStr(alias_name, alias_name_len), ce, 0) == types.SUCCESS {
				return_value.SetTrue()
				return
			} else {
				faults.Error(faults.E_WARNING, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), alias_name)
				return_value.SetFalse()
				return
			}
		} else {
			faults.Error(faults.E_WARNING, "First argument of class_alias() must be a name of user defined class")
			return_value.SetFalse()
			return
		}
	} else {
		faults.Error(faults.E_WARNING, "Class '%s' not found", class_name.GetVal())
		return_value.SetFalse()
		return
	}
}

//@zif -alias get_required_files
func ZifGetIncludedFiles() *types.Array {
	retArr := types.NewArray()
	EG__().GetIncludedFiles().Foreach(func(key types.ArrayKey, value *types.Zval) {
		if key.IsStrKey() {
			retArr.Append(types.NewZvalString(key.StrKey()))
		}
	})
	return retArr
}

//@zif -alias user_error
func ZifTriggerError(executeData zpp.Ex, return_value zpp.Ret, message *types.Zval, _ zpp.Opt, errorType *types.Zval) {
	var error_type = faults.E_USER_NOTICE
	var message *byte
	var message_len int
	if ZendParseParameters(executeData.NumArgs(), "s|l", &message, &message_len, &error_type) == types.FAILURE {
		return
	}
	switch error_type {
	case faults.E_USER_ERROR:
		fallthrough
	case faults.E_USER_WARNING:
		fallthrough
	case faults.E_USER_NOTICE:
		fallthrough
	case faults.E_USER_DEPRECATED:

	default:
		faults.Error(faults.E_WARNING, "Invalid error type specified")
		return_value.SetFalse()
		return
	}
	faults.Error(int(error_type), "%s", message)
	return_value.SetTrue()
	return
}
func ZifSetErrorHandler(executeData zpp.Ex, return_value zpp.Ret, errorHandler *types.Zval, _ zpp.Opt, errorTypes *types.Zval) {
	var error_handler *types.Zval
	var error_type = faults.E_ALL
	if ZendParseParameters(executeData.NumArgs(), "z|l", &error_handler, &error_type) == types.FAILURE {
		return
	}
	if !error_handler.IsNull() {
		if ZendIsCallable(error_handler, 0, nil) == 0 {
			var errorHandlerName = ZendGetCallableName(error_handler)
			faults.Error(faults.E_WARNING, "%s() expects the argument (%s) to be a valid callback", GetActiveFunctionName(), lang.CondF1(errorHandlerName != "", func() string { return errorHandlerName }, "unknown"))
			return
		}
	}
	if EG__().GetUserErrorHandler().IsNotUndef() {
		types.ZVAL_COPY(return_value, EG__().GetUserErrorHandler())
	}
	EG__().GetUserErrorHandlersErrorReporting().Push(EG__().GetUserErrorHandlerErrorReporting())
	EG__().GetUserErrorHandlers().Push(EG__().GetUserErrorHandler())
	if error_handler.IsNull() {
		EG__().GetUserErrorHandler().SetUndef()
		return
	}
	types.ZVAL_COPY(EG__().GetUserErrorHandler(), error_handler)
	EG__().SetUserErrorHandlerErrorReporting(int(error_type))
}
func ZifRestoreErrorHandler(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if EG__().GetUserErrorHandler().IsNotUndef() {
		var zeh types.Zval
		types.ZVAL_COPY_VALUE(&zeh, EG__().GetUserErrorHandler())
		EG__().GetUserErrorHandler().SetUndef()
		// ZvalPtrDtor(&zeh)
	}
	if EG__().GetUserErrorHandlers().IsEmpty() {
		EG__().GetUserErrorHandler().SetUndef()
	} else {
		var tmp *types.Zval
		EG__().SetUserErrorHandlerErrorReporting(ZendStackIntTop(EG__().GetUserErrorHandlersErrorReporting()))
		EG__().GetUserErrorHandlersErrorReporting().DelTop()
		tmp = EG__().GetUserErrorHandlers().Top()
		types.ZVAL_COPY_VALUE(EG__().GetUserErrorHandler(), tmp)
		EG__().GetUserErrorHandlers().DelTop()
	}
	return_value.SetTrue()
	return
}
func ZifSetExceptionHandler(executeData zpp.Ex, return_value zpp.Ret, exceptionHandler *types.Zval) {
	var exception_handler *types.Zval
	if ZendParseParameters(executeData.NumArgs(), "z", &exception_handler) == types.FAILURE {
		return
	}
	if !exception_handler.IsNull() {
		if !ZendIsCallable(exception_handler, 0, nil) {
			var exceptionHandlerName = ZendGetCallableName(exception_handler)
			faults.Error(faults.E_WARNING, "%s() expects the argument (%s) to be a valid callback", GetActiveFunctionName(), lang.CondF1(exceptionHandlerName != "", func() string { return exceptionHandlerName }, "unknown"))
			return
		}
	}
	if EG__().GetUserExceptionHandler().IsNotUndef() {
		types.ZVAL_COPY(return_value, EG__().GetUserExceptionHandler())
	}
	EG__().GetUserExceptionHandlers().Push(EG__().GetUserExceptionHandler())
	if exception_handler.IsNull() {
		EG__().GetUserExceptionHandler().SetUndef()
		return
	}
	types.ZVAL_COPY(EG__().GetUserExceptionHandler(), exception_handler)
}
func ZifRestoreExceptionHandler(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if EG__().GetUserExceptionHandler().IsNotUndef() {
		// ZvalPtrDtor(EG__().GetUserExceptionHandler())
	}
	if EG__().GetUserExceptionHandlers().IsEmpty() {
		EG__().GetUserExceptionHandler().SetUndef()
	} else {
		var tmp *types.Zval = EG__().GetUserExceptionHandlers().Top()
		types.ZVAL_COPY_VALUE(EG__().GetUserExceptionHandler(), tmp)
		EG__().GetUserExceptionHandlers().DelTop()
	}
	return_value.SetTrue()
	return
}

func GetDeclaredClassImpl(flags uint32, skipFlags uint32) *types.Array {
	arr := types.NewArrayCap(EG__().ClassTable().Len())
	EG__().ClassTable().Foreach(func(key string, ce *types.ClassEntry) {
		if key != "" && ce.HasCeFlags(flags) && !ce.HasCeFlags(skipFlags) {
			// 非别名创建的 ce 使用真实类名；class_alias 别名创建的 ce 使用 key 值，此时类名为小写
			if SameNameEx(key, ce.Name()) {
				key = ce.Name()
			}
			// 添加到数组
			arr.Append(types.NewZvalString(key))
		}
	})
	return arr
}
func ZifGetDeclaredTraits() *types.Array {
	return GetDeclaredClassImpl(types.AccTrait, 0)
}
func ZifGetDeclaredClasses() *types.Array {
	return GetDeclaredClassImpl(types.AccLinked, types.AccInterface|types.AccTrait)
}
func ZifGetDeclaredInterfaces() *types.Array {
	return GetDeclaredClassImpl(types.AccInterface, 0)
}
func ZifGetDefinedFunctions(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, excludeDisabled *types.Zval) {
	var internal types.Zval
	var user types.Zval
	var exclude_disabled = 0
	if ZendParseParameters(executeData.NumArgs(), "|b", &exclude_disabled) == types.FAILURE {
		return
	}
	ArrayInit(&internal)
	ArrayInit(&user)
	ArrayInit(return_value)

	EG__().FunctionTable().Foreach(func(key string, func_ types.IFunction) {
		if key != "" {
			if func_.GetType() == ZEND_INTERNAL_FUNCTION && (exclude_disabled == 0 || func_.GetInternalFunction().GetHandler() != ZifDisplayDisabledFunction) {
				AddNextIndexStrEx(&internal, key)
			} else if func_.GetType() == ZEND_USER_FUNCTION {
				AddNextIndexStrEx(&user, key)
			}
		}
	})

	return_value.Array().KeyAddNew("internal", &internal)
	return_value.Array().KeyAddNew("user", &user)
}
func ZifGetDefinedVars(executeData zpp.Ex, return_value zpp.Ret) {
	var symbol_table *types.Array
	if ZendForbidDynamicCall("get_defined_vars()") == types.FAILURE {
		return
	}
	symbol_table = ZendRebuildSymbolTable()
	if symbol_table == nil {
		return
	}
	return_value.SetArray(types.ZendArrayDup(symbol_table))
	return
}
func ZifGetResourceType(executeData zpp.Ex, return_value zpp.Ret, res *types.Zval) {
	var resource_type *byte
	var z_resource_type *types.Zval
	if ZendParseParameters(executeData.NumArgs(), "r", &z_resource_type) == types.FAILURE {
		return
	}
	resource_type = ZendRsrcListGetRsrcType(z_resource_type.Resource())
	if resource_type != nil {
		return_value.SetString(b.CastStrAuto(resource_type))
		return
	} else {
		return_value.SetString("Unknown")
		return
	}
}
func ZifGetResources(_ zpp.Opt, type_ *string) (*types.Array, bool) {
	retArr := types.NewArray()
	if type_ == nil {
		// skip 跳过获取全局resouce的逻辑
	} else if *type_ == "Unknown" {
		// pass
	} else {
		var id = ZendFetchListDtorId(*type_)
		if id <= 0 {
			faults.Error(faults.E_WARNING, "get_resources():  Unknown resource type '%s'", *type_)
			return nil, false
		}
		// pass
	}
	return retArr, true
}
func AddZendextInfo(ext *ZendExtension, arg any) int {
	var name_array = (*types.Zval)(arg)
	AddNextIndexString(name_array, ext.GetName())
	return 0
}
func ZifGetLoadedExtensions(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, zendExtensions *types.Zval) {
	var zendext = 0
	if ZendParseParameters(executeData.NumArgs(), "|b", &zendext) == types.FAILURE {
		return
	}
	ArrayInit(return_value)
	if zendext != 0 {
		ZendExtensions.Apply(func(ext *ZendExtension) {
			AddZendextInfo(ext, return_value)
		})

	} else {
		globals.G().EachModule(func(module *ModuleEntry) {
			AddNextIndexString(return_value, module.GetName())
		})
	}
}
func ZifGetDefinedConstants(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, categorize *types.Zval) {
	var categorize = 0
	if ZendParseParameters(executeData.NumArgs(), "|b", &categorize) == types.FAILURE {
		return
	}
	ArrayInit(return_value)
	if categorize != 0 {
		var val *ZendConstant
		var module_number int
		var modules []types.Zval
		var const_val types.Zval
		var module_names []string
		var i = 1
		modules = make([]types.Zval, globals.G().CountModules()+2)
		module_names = make([]string, globals.G().CountModules()+2)
		module_names[0] = "internal"
		globals.G().EachModule(func(module *ModuleEntry) {
			module_names[module.GetModuleNumber()] = module.GetName()
			i++
		})
		module_names[i] = "user"
		EG__().ConstantTable().Foreach(func(_ string, val *ZendConstant) {
			if val.GetName() == nil {
				/* skip special constants */
				return
			}
			if val.ModuleNumber() == PHP_USER_CONSTANT {
				module_number = i
			} else if val.ModuleNumber() > i {
				/* should not happen */
				return

			} else {
				module_number = val.ModuleNumber()
			}
			if modules[module_number].IsUndef() {
				ArrayInit(&modules[module_number])
				AddAssocZval(return_value, module_names[module_number], &modules[module_number])
			}
			types.ZVAL_COPY_OR_DUP(&const_val, val.Value())
			modules[module_number].Array().KeyAddNew(val.GetName().GetStr(), &const_val)
		})

		Efree(modules)
	} else {
		EG__().ConstantTable().Foreach(func(_ string, constant *ZendConstant) {
			if constant.GetName() == nil {
				/* skip special constants */
				return
			}

			var constVal types.Zval
			types.ZVAL_COPY_OR_DUP(&constVal, constant.Value())
			return_value.Array().KeyAddNew(constant.GetName().GetStr(), &constVal)
		})
	}
}
func DebugBacktraceGetArgs(call *ZendExecuteData, arg_array *types.Zval) {
	var numArgs = call.NumArgs()
	if numArgs != 0 {

		var i = 0
		var p = call.Arg(1)
		ArrayInitSize(arg_array, numArgs)

		arr := types.NewArrayCap(numArgs)
		if call.GetFunc().GetType() == ZEND_USER_FUNCTION {
			var firstExtraArg = b.Min(numArgs, call.GetFunc().GetOpArray().GetNumArgs())
			if (ZEND_CALL_INFO(call) & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
				/* In case of attached symbol_table, values on stack may be invalid
				 * and we have to access them through symbol_table
				 * See: https://bugs.php.net/bug.php?id=73156
				 */
				var arg *types.Zval
				for ; i < firstExtraArg; i++ {
					argName := call.GetFunc().GetOpArray().VarName(i)
					arg = types.ZendHashFindInd(call.GetSymbolTable(), argName)
					if arg != nil {
						arr.Append(arg)
					} else {
						arr.Append(types.NewZvalNull())
					}
				}
			} else {
				for ; i < firstExtraArg; i++ {
					if !p.IsUndef() {
						arr.Append(p)
					} else {
						arr.Append(types.NewZvalNull())
					}
					p++
				}
			}
			p = call.VarNum(call.GetFunc().GetOpArray().GetLastVar() + call.GetFunc().GetOpArray().GetT())
		}
		for ; i < numArgs; i++ {
			if !p.IsUndef() {
				arr.Append(p)
			} else {
				arr.Append(types.NewZvalNull())
			}
			p++
		}
		arg_array.SetArray(arr)
	} else {
		arg_array.SetEmptyArray()
	}
}
func DebugPrintBacktraceArgs(argArray *types.Zval) {
	var i = 0
	argArray.Array().Foreach(func(key types.ArrayKey, value *types.Zval) {
		if i != 0 {
			ZEND_PUTS(", ")
		}
		i++
		ZendPrintFlatZvalR(value)
	})
}
func SkipInternalHandler(skip *ZendExecuteData) bool {
	return !(skip.GetFunc() != nil && ZEND_USER_CODE(skip.GetFunc().GetType())) && skip.GetPrevExecuteData() != nil && skip.GetPrevExecuteData().GetFunc() != nil && ZEND_USER_CODE(skip.GetPrevExecuteData().GetFunc().GetType()) && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_ICALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_UCALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL_BY_NAME && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL
}
func ZifDebugPrintBacktrace(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, options *types.Zval, limit *types.Zval) {
	var call *ZendExecuteData
	var ptr *ZendExecuteData
	var skip *ZendExecuteData
	var object *types.Object
	var lineno int
	var frameno = 0
	var func_ types.IFunction
	var functionName string
	var filename string = ""
	var class_name string = ""
	var call_type string = ""
	var include_filename string = ""
	var arg_array types.Zval
	var indent = 0
	if ZendParseParameters(executeData.NumArgs(), "|ll", &options, &limit) == types.FAILURE {
		return
	}
	arg_array.SetUndef()
	ptr = executeData.GetPrevExecuteData()

	/* skip debug_backtrace() */

	call = ptr
	ptr = ptr.GetPrevExecuteData()
	for ptr != nil && (limit == 0 || frameno < limit) {
		frameno++
		class_name = ""
		call_type = nil
		arg_array.SetUndef()
		ptr = ZendGeneratorCheckPlaceholderFrame(ptr)
		skip = ptr

		/* skip internal handler */

		if SkipInternalHandler(skip) != 0 {
			skip = skip.GetPrevExecuteData()
		}
		if skip.GetFunc() != nil && ZEND_USER_CODE(skip.GetFunc().GetType()) {
			filename = skip.GetFunc().GetOpArray().GetFilename()
			if skip.GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION {
				if EG__().GetOplineBeforeException() != nil {
					lineno = EG__().GetOplineBeforeException().GetLineno()
				} else {
					lineno = skip.GetFunc().GetOpArray().GetLineEnd()
				}
			} else {
				lineno = skip.GetOpline().GetLineno()
			}
		} else {
			filename = ""
			lineno = 0
		}

		/* $this may be passed into regular internal functions */

		object = call.ThisObject()
		if call.GetFunc() != nil {
			func_ = call.GetFunc()
			if func_.GetScope() != nil && func_.GetScope().GetTraitAliases() != nil {
				if object != nil {
					functionName = ZendResolveMethodName(object.GetCe(), func_)
				} else {
					functionName = ZendResolveMethodName(func_.GetScope(), func_)
				}
			} else {
				functionName = func_.FunctionName()
			}
		} else {
			func_ = nil
			functionName = ""
		}
		if functionName != "" {
			if object != nil {
				if func_.GetScope() != nil {
					class_name = func_.GetScope().Name()
				} else {
					class_name = object.ClassName()
				}
				call_type = "->"
			} else if func_.GetScope() != nil {
				class_name = func_.GetScope().Name()
				call_type = "::"
			} else {
				class_name = ""
				call_type = ""
			}
			if func_.GetType() != ZEND_EVAL_CODE {
				if (options & DEBUG_BACKTRACE_IGNORE_ARGS) == 0 {
					DebugBacktraceGetArgs(call, &arg_array)
				}
			}
		} else {

			/* i know this is kinda ugly, but i'm trying to avoid extra cycles in the main execution loop */

			var build_filename_arg = 1
			if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetType())) || ptr.GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL {

				/* can happen when calling eval from a custom sapi */

				functionName = "unknown"
				build_filename_arg = 0
			} else {
				switch ptr.GetOpline().GetExtendedValue() {
				case ZEND_EVAL:
					functionName = "eval"
					build_filename_arg = 0
				case ZEND_INCLUDE:
					functionName = "include"
				case ZEND_REQUIRE:
					functionName = "require"
				case ZEND_INCLUDE_ONCE:
					functionName = "include_once"
				case ZEND_REQUIRE_ONCE:
					functionName = "require_once"
				default:

					/* this can actually happen if you use debug_backtrace() in your error_handler and
					 * you're in the top-scope */

					functionName = "unknown"
					build_filename_arg = 0
				}
			}
			if build_filename_arg != 0 && include_filename != "" {
				ArrayInit(&arg_array)
				arg_array.Array().Append(types.NewZvalString(include_filename))
			}
			call_type = nil
		}
		ZendPrintf("#%-2d ", indent)
		if class_name != "" {
			ZEND_PUTS(class_name)
			ZEND_PUTS(call_type)
		}
		ZendPrintf("%s(", functionName)
		if arg_array.IsNotUndef() {
			DebugPrintBacktraceArgs(&arg_array)
			// ZvalPtrDtor(&arg_array)
		}
		if filename != "" {
			ZendPrintf(") called at [%s:%d]\n", filename, lineno)
		} else {
			var prev_call = skip
			var prev = skip.GetPrevExecuteData()
			for prev != nil {
				if prev_call != nil && prev_call.GetFunc() != nil && !(ZEND_USER_CODE(prev_call.GetFunc().GetType())) {
					prev = nil
					break
				}
				if prev.GetFunc() != nil && ZEND_USER_CODE(prev.GetFunc().GetType()) {
					ZendPrintf(") called at [%s:%d]\n", prev.GetFunc().GetOpArray().GetFilename(), prev.GetOpline().GetLineno())
					break
				}
				prev_call = prev
				prev = prev.GetPrevExecuteData()
			}
			if prev == nil {
				ZEND_PUTS(")\n")
			}
		}
		include_filename = filename
		call = skip
		ptr = skip.GetPrevExecuteData()
		indent++
	}
}
func ZendFetchDebugBacktrace(return_value *types.Zval, skip_last int, options int, limit int) {
	var ptr *ZendExecuteData
	var skip *ZendExecuteData
	var call *ZendExecuteData = nil
	var object *types.Object
	var lineno int
	var frameno = 0
	var func_ types.IFunction
	var functionName string
	var filename string
	var include_filename string = ""
	var stack_frame types.Zval
	var tmp types.Zval
	ArrayInit(return_value)
	if !(lang.Assign(&ptr, CurrEX())) {
		return
	}
	if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetType())) {
		call = ptr
		ptr = ptr.GetPrevExecuteData()
	}
	if ptr != nil {
		if skip_last != 0 {

			/* skip debug_backtrace() */

			call = ptr
			ptr = ptr.GetPrevExecuteData()
		} else {

			/* skip "new Exception()" */

			if ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetType()) && ptr.GetOpline().GetOpcode() == ZEND_NEW {
				call = ptr
				ptr = ptr.GetPrevExecuteData()
			}

			/* skip "new Exception()" */

		}
		if call == nil {
			call = ptr
			ptr = ptr.GetPrevExecuteData()
		}
	}
	for ptr != nil && (limit == 0 || frameno < limit) {
		frameno++
		ArrayInit(&stack_frame)
		ptr = ZendGeneratorCheckPlaceholderFrame(ptr)
		skip = ptr

		/* skip internal handler */

		if SkipInternalHandler(skip) != 0 {
			skip = skip.GetPrevExecuteData()
		}
		if skip.GetFunc() != nil && ZEND_USER_CODE(skip.GetFunc().GetType()) {
			filename = skip.GetFunc().GetOpArray().GetFilename()
			if skip.GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION {
				if EG__().GetOplineBeforeException() != nil {
					lineno = EG__().GetOplineBeforeException().GetLineno()
				} else {
					lineno = skip.GetFunc().GetOpArray().GetLineEnd()
				}
			} else {
				lineno = skip.GetOpline().GetLineno()
			}
			tmp.SetString(filename)
			stack_frame.Array().KeyAddNew(types.STR_FILE, &tmp)
			tmp.SetLong(lineno)
			stack_frame.Array().KeyAddNew(types.STR_LINE, &tmp)
		} else {
			var prev_call = skip
			var prev = skip.GetPrevExecuteData()
			for prev != nil {
				if prev_call != nil && prev_call.GetFunc() != nil && !(ZEND_USER_CODE(prev_call.GetFunc().GetType())) && !prev_call.GetFunc().IsCallViaTrampoline() {
					break
				}
				if prev.GetFunc() != nil && ZEND_USER_CODE(prev.GetFunc().GetType()) {
					tmp.SetString(prev.GetFunc().GetOpArray().GetFilename())
					stack_frame.Array().KeyAddNew(types.STR_FILE, &tmp)
					tmp.SetLong(prev.GetOpline().GetLineno())
					stack_frame.Array().KeyAddNew(types.STR_LINE, &tmp)
					break
				}
				prev_call = prev
				prev = prev.GetPrevExecuteData()
			}
			filename = ""
		}

		/* $this may be passed into regular internal functions */

		if call != nil {
			object = call.ThisObject()
		} else {
			object = nil
		}
		if call != nil && call.GetFunc() != nil {
			func_ = call.GetFunc()
			if func_.GetScope() != nil && func_.GetScope().GetTraitAliases() != nil {
				if object != nil {
					functionName = ZendResolveMethodName(object.GetCe(), func_)
				} else {
					functionName = ZendResolveMethodName(func_.GetScope(), func_)
				}
			} else {
				functionName = func_.FunctionName()
			}
		} else {
			func_ = nil
			functionName = ""
		}
		if functionName != "" {
			tmp.SetString(functionName)
			stack_frame.Array().KeyAddNew(types.STR_FUNCTION, &tmp)
			if object != nil {
				if func_.GetScope() != nil {
					tmp.SetString(func_.GetScope().GetName().GetStr())
				} else {
					tmp.SetString(object.ClassName())
				}
				stack_frame.Array().KeyAddNew(types.STR_CLASS, &tmp)
				if (options & DEBUG_BACKTRACE_PROVIDE_OBJECT) != 0 {
					tmp.SetObject(object)
					stack_frame.Array().KeyAddNew(types.STR_OBJECT, &tmp)
					// 					tmp.AddRefcount()
				}
				tmp.SetString(types.STR_OBJECT_OPERATOR)
				stack_frame.Array().KeyAddNew(types.STR_TYPE, &tmp)
			} else if func_.GetScope() != nil {
				tmp.SetString(func_.GetScope().GetName().GetStr())
				stack_frame.Array().KeyAddNew(types.STR_CLASS, &tmp)
				tmp.SetString(types.STR_PAAMAYIM_NEKUDOTAYIM)
				stack_frame.Array().KeyAddNew(types.STR_TYPE, &tmp)
			}
			if (options&DEBUG_BACKTRACE_IGNORE_ARGS) == 0 && func_.GetType() != ZEND_EVAL_CODE {
				DebugBacktraceGetArgs(call, &tmp)
				stack_frame.Array().KeyAddNew(types.STR_ARGS, &tmp)
			}
		} else {
			/* i know this is kinda ugly, but i'm trying to avoid extra cycles in the main execution loop */
			var build_filename_arg = 1
			var pseudo_function_name string
			if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetType())) || ptr.GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL {
				/* can happen when calling eval from a custom sapi */
				pseudo_function_name = types.STR_UNKNOWN
				build_filename_arg = 0
			} else {
				switch ptr.GetOpline().GetExtendedValue() {
				case ZEND_EVAL:
					pseudo_function_name = types.STR_EVAL
					build_filename_arg = 0
				case ZEND_INCLUDE:
					pseudo_function_name = types.STR_INCLUDE
				case ZEND_REQUIRE:
					pseudo_function_name = types.STR_REQUIRE
				case ZEND_INCLUDE_ONCE:
					pseudo_function_name = types.STR_INCLUDE_ONCE
				case ZEND_REQUIRE_ONCE:
					pseudo_function_name = types.STR_REQUIRE_ONCE
				default:

					/* this can actually happen if you use debug_backtrace() in your error_handler and
					 * you're in the top-scope */

					pseudo_function_name = types.STR_UNKNOWN
					build_filename_arg = 0
				}
			}
			if build_filename_arg != 0 && include_filename != nil {
				var arg_array types.Zval
				ArrayInit(&arg_array)

				/* include_filename always points to the last filename of the last last called-function.
				   if we have called include in the frame above - this is the file we have included.
				*/

				tmp.SetString(include_filename.GetStr())
				arg_array.Array().AppendNew(&tmp)
				stack_frame.Array().KeyAddNew(types.STR_ARGS, &arg_array)
			}
			tmp.SetString(pseudo_function_name)
			stack_frame.Array().KeyAddNew(types.STR_FUNCTION, &tmp)
		}
		return_value.Array().AppendNew(&stack_frame)
		include_filename = filename
		call = skip
		ptr = skip.GetPrevExecuteData()
	}
}
func ZifDebugBacktrace(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, options *types.Zval, limit *types.Zval) {
	var options = DEBUG_BACKTRACE_PROVIDE_OBJECT
	var limit = 0
	if ZendParseParameters(executeData.NumArgs(), "|ll", &options, &limit) == types.FAILURE {
		return
	}
	ZendFetchDebugBacktrace(return_value, 1, options, limit)
}
func ZifExtensionLoaded(extensionName string) bool {
	return globals.G().GetModule(extensionName) != nil
}
func ZifGetExtensionFuncs(executeData zpp.Ex, return_value zpp.Ret, extensionName *types.Zval) {
	var extension_name *types.String
	var array int
	var module *ModuleEntry
	if ZendParseParameters(executeData.NumArgs(), "S", &extension_name) == types.FAILURE {
		return
	}

	if !ascii.StrCaseEquals(extension_name.GetStr(), "zend") {
		module = globals.G().GetModule(extension_name.GetStr())
	} else {
		module = globals.G().GetModule("core")
	}
	if module == nil {
		return_value.SetFalse()
		return
	}
	if module.GetFunctions() != nil {

		/* avoid BC break, if functions list is empty, will return an empty array */

		ArrayInit(return_value)
		array = 1
	} else {
		array = 0
	}
	CG__().FunctionTable().Foreach(func(_ string, f types.IFunction) {
		if f.GetType() == ZEND_INTERNAL_FUNCTION && f.GetInternalFunction().GetModule() == module {
			if array == 0 {
				ArrayInit(return_value)
				array = 1
			}
			return_value.Array().Append(types.NewZvalString(f.FunctionName()))
		}
	})
	if array == 0 {
		return_value.SetFalse()
		return
	}
}
