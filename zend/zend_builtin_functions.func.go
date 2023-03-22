package zend

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
	"strings"
)

func ZmStartupCore(type_ int, module_number int) int {
	var class_entry types.ClassEntry
	memset(&class_entry, 0, b.SizeOf("zend_class_entry"))
	class_entry.SetName(types.ZendStringInitInterned("stdClass", b.SizeOf("\"stdClass\"")-1, 1))
	class_entry.SetBuiltinFunctions(nil)
	ZendStandardClassDef = ZendRegisterInternalClass(&class_entry)
	ZendRegisterDefaultClasses()
	return types.SUCCESS
}
func ZendStartupBuiltinFunctions() int {
	ZendBuiltinModule.SetModuleNumber(0)
	ZendBuiltinModule.SetType(MODULE_PERSISTENT)
	if b.Assign(&(EG__().GetCurrentModule()), ZendRegisterModuleEx(&ZendBuiltinModule)) == nil {
		return types.FAILURE
	} else {
		return types.SUCCESS
	}
}

func ZifZendVersion() string  { return ZEND_VERSION }
func ZifGcMemCaches() int     { return ZendMmGc(ZendMmGetHeap()) }
func ZifGcCollectCycles() int { return 0 }
func ZifGcEnabled() bool      { return true }
func ZifGcEnable() {
	ZendAlterIniEntryChars("zend.enable_gc", "1", ZEND_INI_USER, ZEND_INI_STAGE_RUNTIME)
}
func ZifGcDisable() {
	ZendAlterIniEntryChars("zend.enable_gc", "0", ZEND_INI_USER, ZEND_INI_STAGE_RUNTIME)
}

func ZifGcStatus(ret zpp.DefRet) {
	ArrayInitSize(ret, 3)
	AddAssocLongEx(ret, "runs", 0)
	AddAssocLongEx(ret, "collected", 0)
	AddAssocLongEx(ret, "threshold", 0)
	AddAssocLongEx(ret, "roots", 0)
}

func ZifFuncNumArgs(executeData zpp.DefEx) int {
	return IZifFuncNumArgs(executeData)
}

func IZifFuncNumArgs(executeData *ZendExecuteData) int {
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
func ZifFuncGetArg(executeData zpp.DefEx, returnValue zpp.DefRet, argNum int) {
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

	arg_count := ex.NumArgs()
	if argNum >= arg_count {
		faults.Error(faults.E_WARNING, "func_get_arg():  Argument "+ZEND_LONG_FMT+" not passed to function", argNum)
		returnValue.SetFalse()
		return
	}

	var arg *types.Zval
	first_extra_arg := int(ex.GetFunc().GetOpArray().GetNumArgs())
	if argNum >= first_extra_arg && ex.NumArgs() > first_extra_arg {
		arg = ex.VarNum(ex.GetFunc().GetOpArray().GetLastVar() + int(ex.GetFunc().GetOpArray().GetT()) + (argNum - first_extra_arg))
	} else {
		arg = ex.Arg(argNum + 1)
	}

	if !(arg.IsUndef()) {
		types.ZVAL_COPY_DEREF(returnValue, arg)
	}
}
func ZifFuncGetArgs(executeData zpp.DefEx) (*types.Array, bool) {
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
		return types.NewEmptyArray(), true
	}

	first_extra_arg := int(ex.GetFunc().GetOpArray().GetNumArgs())
	var values []*types.Zval
	if argCount <= first_extra_arg {
		values = executeData.Args(0, argCount)
	} else {
		values = executeData.Args(0, first_extra_arg)

		for i := 0; i < argCount-first_extra_arg; i++ {
			p := ex.VarNum(ex.GetFunc().GetOpArray().GetLastVar() + int(ex.GetFunc().GetOpArray().GetT()) + i)
			values = append(values, p)
		}
	}

	arr := types.NewZendArray(argCount)
	for _, zv := range values {
		if zv.IsUndef() {
			arr.NextIndexInsertNew(types.NewZvalNull())
		} else {
			zv = types.ZVAL_DEREF(zv)
			if zv.IsRefcounted() {
				zv.AddRefcount()
			}
			arr.NextIndexInsertNew(zv)
		}
	}
	return arr, true
}
func ZifStrlen(str *types.String) int        { return str.GetLen() }
func ZifStrcmp(str1 string, str2 string) int { return strings.Compare(str1, str2) }
func ZifStrncmp(str1 string, str2 string, len_ int) (int, bool) {
	if len_ < 0 {
		faults.Error(faults.E_WARNING, "Length must be greater than or equal to 0")
		return 0, false
	}
	return ZendBinaryStrncmp(str1, str2, len_), true
}
func ZifStrcasecmp(str1 string, str2 string) int { return ZendBinaryStrcasecmp(str1, str2) }
func ZifStrncasecmp(str1 string, str2 string, len_ int) (int, bool) {
	if len_ < 0 {
		faults.Error(faults.E_WARNING, "Length must be greater than or equal to 0")
		return 0, false
	}
	return ZendBinaryStrncasecmp(str1, str2, len_), true
}
func ZifEach(executeData *ZendExecuteData, return_value *types.Zval) {
	var array *types.Zval
	var entry *types.Zval
	var tmp types.Zval
	var num_key ZendUlong
	var target_hash *types.Array
	var key *types.String
	if ZendParseParameters(executeData.NumArgs(), "z/", &array) == types.FAILURE {
		return
	}
	if EG__().GetEachDeprecationThrown() == 0 {
		faults.Error(faults.E_DEPRECATED, "The each() function is deprecated. This message will be suppressed on further calls")
		EG__().SetEachDeprecationThrown(1)
	}
	target_hash = HASH_OF(array)
	if target_hash == nil {
		faults.Error(faults.E_WARNING, "Variable passed to each() is not an array or object")
		return
	}
	for true {
		entry = types.ZendHashGetCurrentData(target_hash)
		if entry == nil {
			return_value.SetFalse()
			return
		} else if entry.IsIndirect() {
			entry = entry.GetZv()
			if entry.IsUndef() {
				types.ZendHashMoveForward(target_hash)
				continue
			}
		}
		break
	}
	ArrayInitSize(return_value, 4)
	types.ZendHashRealInitMixed(return_value.GetArr())

	/* add value elements */

	entry = types.ZVAL_DEREF(entry)
	if entry.IsRefcounted() {
		entry.GetCounted().AddRefcountEx(2)
	}
	return_value.GetArr().IndexAddNewH(1, entry)
	return_value.GetArr().KeyAddNew(types.ZSTR_VALUE.GetStr(), entry)

	/* add the key elements */

	if types.ZendHashGetCurrentKey(target_hash, &key, &num_key) == types.HASH_KEY_IS_STRING {
		tmp.SetStringCopy(key)
		tmp.TryAddRefcount()
	} else {
		tmp.SetLong(num_key)
	}
	return_value.GetArr().IndexAddNewH(0, &tmp)
	return_value.GetArr().KeyAddNew(types.ZSTR_KEY.GetStr(), &tmp)
	types.ZendHashMoveForward(target_hash)
}
func ZifErrorReporting(ret zpp.DefRet, _ zpp.DefOpt, newErrorLevel *types.Zval) {
	var old_error_reporting int
	old_error_reporting = EG__().GetErrorReporting()
	if newErrorLevel != nil {
		var new_val = ZvalTryGetString(newErrorLevel)
		if new_val == nil {
			return
		}
		for {
			var p = EG__().GetErrorReportingIniEntry()
			if p == nil {
				var zv = EG__().GetIniDirectives().KeyFind(types.ZSTR_ERROR_REPORTING.GetStr())
				if zv != nil {
					EG__().SetErrorReportingIniEntry((*ZendIniEntry)(zv.GetPtr()))
					p = EG__().GetErrorReportingIniEntry()
				} else {
					break
				}
			}
			if p.GetModified() == 0 {
				if EG__().GetModifiedIniDirectives() == nil {
					ALLOC_HASHTABLE(EG__().GetModifiedIniDirectives())
					types.ZendHashInit(EG__().GetModifiedIniDirectives(), 8, nil, nil, 0)
				}
				if types.ZendHashAddPtr(EG__().GetModifiedIniDirectives(), types.ZSTR_ERROR_REPORTING, p) != nil {
					p.SetOrigValue(p.GetValue())
					p.SetOrigModifiable(p.GetModifiable())
					p.SetModified(1)
				}
			} else if p.GetOrigValue() != p.GetValue() {
				types.ZendStringReleaseEx(p.GetValue(), 0)
			}
			p.SetValue(new_val)
			if newErrorLevel.IsLong() {
				EG__().SetErrorReporting(newErrorLevel.GetLval())
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
	var val *types.Zval
	ht.ProtectRecursive()
	var __ht = ht
	for _, _p := range __ht.foreachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.GetZv()
			if _z.IsUndef() {
				continue
			}
		}
		val = _z
		val = types.ZVAL_DEREF(val)
		if val.IsRefcounted() {
			if val.IsArray() {
				if val.IsRefcounted() {
					if val.IsRecursive() {
						faults.Error(faults.E_WARNING, "Constants cannot be recursive arrays")
						ret = 0
						break
					} else if ValidateConstantArray(val.GetArr()) == 0 {
						ret = 0
						break
					}
				}
			} else if val.GetType() != types.IS_STRING && val.GetType() != types.IS_RESOURCE {
				faults.Error(faults.E_WARNING, "Constants may only evaluate to scalar values, arrays or resources")
				ret = 0
				break
			}
		}
	}
	ht.UnprotectRecursive()
	return ret
}
func CopyConstantArray(dst *types.Zval, src *types.Zval) {
	var key *types.String
	var idx ZendUlong
	var new_val *types.Zval
	var val *types.Zval
	ArrayInitSize(dst, types.Z_ARRVAL_P(src).GetNNumOfElements())
	var __ht = src.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.GetZv()
			if _z.IsUndef() {
				continue
			}
		}
		idx = _p.GetH()
		key = _p.GetKey()
		val = _z

		/* constant arrays can't contain references */

		val = types.ZVAL_DEREF(val)
		if key != nil {
			new_val = dst.GetArr().KeyAddNew(key.GetStr(), val)
		} else {
			new_val = dst.GetArr().IndexAddNewH(idx, val)
		}
		if val.IsArray() {
			if val.IsRefcounted() {
				CopyConstantArray(new_val, val)
			}
		} else {
			val.TryAddRefcount()
		}
	}
}

func ZifDefine(constantName string, value *types.Zval, _ zpp.DefOpt, caseInsensitive bool) bool {
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
	switch value.GetType() {
	case types.IS_NULL,
		types.IS_FALSE,
		types.IS_TRUE,
		types.IS_LONG,
		types.IS_DOUBLE,
		types.IS_STRING,
		types.IS_RESOURCE:
		// pass
	case types.IS_ARRAY:
		if value.IsRefcounted() {
			if ValidateConstantArray(value.GetArr()) == 0 {
				return false
			} else {
				CopyConstantArray(c.Value(), value)
				goto register_constant
			}
		}
	case types.IS_OBJECT:
		if val_free.IsUndef() {
			if types.Z_OBJ_HT_P(value).GetGet() != nil {
				value = types.Z_OBJ_HT_P(value).GetGet()(value, &val_free)
				goto repeat
			} else if types.Z_OBJ_HT_P(value).GetCastObject() != nil {
				if types.Z_OBJ_HT_P(value).GetCastObject()(value, &val_free, types.IS_STRING) == types.SUCCESS {
					value = &val_free
					break
				}
			}
		}
		fallthrough
	default:
		faults.Error(faults.E_WARNING, "Constants may only evaluate to scalar values, arrays or resources")
		ZvalPtrDtor(&val_free)
		return false
	}
	types.ZVAL_COPY(c.Value(), value)
	ZvalPtrDtor(&val_free)
register_constant:
	if caseInsensitive {
		faults.Error(faults.E_DEPRECATED, "define(): Declaration of case-insensitive constants is deprecated")
	}

	/* non persistent */

	ZEND_CONSTANT_SET_FLAGS(&c, caseSensitive, PHP_USER_CONSTANT)
	c.SetNameVal(constantName)
	if ZendRegisterConstant(&c) == types.SUCCESS {
		return true
	} else {
		return false
	}
}
func ZifDefined(constantName string) bool {
	if ZendGetConstantEx(types.NewString(constantName), ZendGetExecutedScope(), ZEND_FETCH_CLASS_SILENT|ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK) != nil {
		return true
	} else {
		return false
	}
}
func ZifGetClass(executeData *ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval = nil
	if ZendParseParameters(executeData.NumArgs(), "|o", &obj) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	if obj == nil {
		var scope = ZendGetExecutedScope()
		if scope != nil {
			return_value.SetStringCopy(scope.GetName())
			return
		} else {
			faults.Error(faults.E_WARNING, "get_class() called without object from outside a class")
			return_value.SetFalse()
			return
		}
	}
	return_value.SetStringCopy(types.Z_OBJCE_P(obj).GetName())
	return
}
func ZifGetCalledClass(ex zpp.DefEx) (string, bool) {
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
func ZifGetParentClass(executeData *ZendExecuteData, return_value *types.Zval) {
	var arg *types.Zval
	var ce *types.ClassEntry = nil
	if ZendParseParameters(executeData.NumArgs(), "|z", &arg) == types.FAILURE {
		return
	}
	if executeData.NumArgs() == 0 {
		ce = ZendGetExecutedScope()
		if ce != nil && ce.GetParent() {
			return_value.SetStringCopy(ce.GetParent().name)
			return
		} else {
			return_value.SetFalse()
			return
		}
	}
	if arg.IsObject() {
		ce = types.Z_OBJ_P(arg).GetCe()
	} else if arg.IsString() {
		ce = ZendLookupClass(arg.GetStr())
	}
	if ce != nil && ce.GetParent() {
		return_value.SetStringCopy(ce.GetParent().name)
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func IsAImpl(executeData *ZendExecuteData, return_value *types.Zval, only_subclass types.ZendBool) {
	var obj *types.Zval
	var class_name *types.String
	var instance_ce *types.ClassEntry
	var ce *types.ClassEntry
	var allow_string = only_subclass
	var retval types.ZendBool
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
		instance_ce = ZendLookupClass(obj.GetStr())
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
	if only_subclass == 0 && types.ZendStringEquals(instance_ce.GetName(), class_name) != 0 {
		retval = 1
	} else {
		ce = ZendLookupClassEx(class_name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if ce == nil {
			retval = 0
		} else {
			if only_subclass != 0 && instance_ce == ce {
				retval = 0
			} else {
				retval = InstanceofFunction(instance_ce, ce)
			}
		}
	}
	types.ZVAL_BOOL(return_value, retval != 0)
	return
}
func ZifIsSubclassOf(executeData *ZendExecuteData, return_value *types.Zval) {
	IsAImpl(executeData, return_value, 1)
}
func ZifIsA(executeData *ZendExecuteData, return_value *types.Zval) {
	IsAImpl(executeData, return_value, 0)
}
func AddClassVars(scope *types.ClassEntry, ce *types.ClassEntry, statics int, return_value *types.Zval) {
	var prop_info *ZendPropertyInfo
	var prop *types.Zval
	var prop_copy types.Zval
	var key *types.String
	var __ht *types.Array = ce.GetPropertiesInfo()
	for _, _p := range __ht.foreachData() {
		var _z = _p.GetVal()

		key = _p.GetKey()
		prop_info = _z.GetPtr()
		if prop_info.IsProtected() && ZendCheckProtected(prop_info.GetCe(), scope) == 0 || prop_info.IsPrivate() && prop_info.GetCe() != scope {
			continue
		}
		prop = nil
		if statics != 0 && prop_info.IsStatic() {
			prop = ce.GetDefaultStaticMembersTable()[prop_info.GetOffset()]
			prop = types.ZVAL_DEINDIRECT(prop)
		} else if statics == 0 && !prop_info.IsStatic() {
			prop = ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(prop_info.GetOffset())]
		}
		if prop == nil {
			continue
		}
		if prop.IsUndef() {

			/* Return uninitialized typed properties as a null value */

			prop_copy.SetNull()

			/* Return uninitialized typed properties as a null value */

		} else {

			/* copy: enforce read only access */

			types.ZVAL_COPY_OR_DUP(&prop_copy, prop)

			/* copy: enforce read only access */

		}
		prop = &prop_copy

		/* this is necessary to make it able to work with default array
		 * properties, returned to user */

		if prop.IsConstant() {
			if ZvalUpdateConstantEx(prop, nil) != types.SUCCESS {
				return
			}
		}
		return_value.GetArr().KeyAddNew(key.GetStr(), prop)
	}
}
func ZifGetClassVars(executeData *ZendExecuteData, return_value *types.Zval) {
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
func ZifGetObjectVars(executeData *ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var value *types.Zval
	var properties *types.Array
	var key *types.String
	var zobj *types.ZendObject
	var num_key ZendUlong
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
	properties = types.Z_OBJ_HT_P(obj).GetGetProperties()(obj)
	if properties == nil {
		return_value.SetFalse()
		return
	}
	zobj = obj.GetObj()
	if zobj.GetCe().GetDefaultPropertiesCount() == 0 && properties == zobj.GetProperties() && !(properties.IsRecursive()) {

		/* fast copy */

		if zobj.GetHandlers() == &StdObjectHandlers {
			return_value.SetArray(types.ZendProptableToSymtable(properties, 0))
			return
		}
		return_value.SetArray(types.ZendProptableToSymtable(properties, 1))
		return
	} else {
		ArrayInitSize(return_value, properties.GetNNumOfElements())
		var __ht = properties
		for _, _p := range __ht.foreachData() {
			var _z = _p.GetVal()

			num_key = _p.GetH()
			key = _p.GetKey()
			value = _z
			var is_dynamic = 1
			if value.IsIndirect() {
				value = value.GetZv()
				if value.IsUndef() {
					continue
				}
				is_dynamic = 0
			}
			if key != nil && ZendCheckPropertyAccess(zobj, key, is_dynamic) == types.FAILURE {
				continue
			}
			if value.IsReference() && value.GetRefcount() == 1 {
				value = types.Z_REFVAL_P(value)
			}
			value.TryAddRefcount()
			if key == nil {

				/* This case is only possible due to loopholes, e.g. ArrayObject */

				return_value.GetArr().IndexAddH(num_key, value)

				/* This case is only possible due to loopholes, e.g. ArrayObject */

			} else if is_dynamic == 0 && key.GetVal()[0] == 0 {
				var prop_name *byte
				var class_name *byte
				var prop_len int
				ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, &prop_len)

				/* We assume here that a mangled property name is never
				 * numeric. This is probably a safe assumption, but
				 * theoretically someone might write an extension with
				 * private, numeric properties. Well, too bad.
				 */

				return_value.GetArr().KeyAddNew(b.CastStr(prop_name, prop_len), value)

				/* We assume here that a mangled property name is never
				 * numeric. This is probably a safe assumption, but
				 * theoretically someone might write an extension with
				 * private, numeric properties. Well, too bad.
				 */

			} else {
				return_value.GetArr().SymtableAddNew(key.GetStr(), value)
			}
		}
	}
}
func ZifGetMangledObjectVars(executeData *ZendExecuteData, return_value *types.Zval) {
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
	properties = types.Z_OBJ_HT_P(obj).GetGetProperties()(obj)
	if properties == nil {
		types.ZVAL_EMPTY_ARRAY(return_value)
		return
	}
	properties = types.ZendProptableToSymtable(properties, types.Z_OBJCE_P(obj).GetDefaultPropertiesCount() != 0 || types.Z_OBJ_P(obj).GetHandlers() != &StdObjectHandlers || properties.IsRecursive())
	return_value.SetArray(properties)
	return
}
func SameName(key *types.String, name *types.String) int {
	var lcname *types.String
	var ret int
	if key == name {
		return 1
	}
	if key.GetLen() != name.GetLen() {
		return 0
	}
	lcname = ZendStringTolower(name)
	ret = memcmp(lcname.GetVal(), key.GetVal(), key.GetLen()) == 0
	types.ZendStringReleaseEx(lcname, 0)
	return ret
}
func ZifGetClassMethods(executeData *ZendExecuteData, return_value *types.Zval) {
	var klass *types.Zval
	var method_name types.Zval
	var ce *types.ClassEntry = nil
	var scope *types.ClassEntry
	var mptr *ZendFunction
	var key *types.String
	if ZendParseParameters(executeData.NumArgs(), "z", &klass) == types.FAILURE {
		return
	}
	if klass.IsObject() {
		ce = types.Z_OBJCE_P(klass)
	} else if klass.IsString() {
		ce = ZendLookupClass(klass.GetStr())
	}
	if ce == nil {
		return_value.SetNull()
		return
	}
	ArrayInit(return_value)
	scope = ZendGetExecutedScope()
	var __ht *types.Array = ce.GetFunctionTable()
	for _, _p := range __ht.foreachData() {
		var _z = _p.GetVal()

		key = _p.GetKey()
		mptr = _z.GetPtr()
		if mptr.IsPublic() || scope != nil && (mptr.IsProtected() && ZendCheckProtected(mptr.GetScope(), scope) != 0 || mptr.IsPrivate() && scope == mptr.GetScope()) {
			if mptr.GetType() == ZEND_USER_FUNCTION && (mptr.GetOpArray().GetRefcount() == nil || mptr.op_array.refcount > 1) && key != nil && SameName(key, mptr.GetFunctionName()) == 0 {
				method_name.SetStringCopy(ZendFindAliasName(mptr.GetScope(), key))
				return_value.GetArr().NextIndexInsertNew(&method_name)
			} else {
				method_name.SetStringCopy(mptr.GetFunctionName())
				return_value.GetArr().NextIndexInsertNew(&method_name)
			}
		}
	}
}
func ZifMethodExists(executeData *ZendExecuteData, return_value *types.Zval) {
	var klass *types.Zval
	var method_name *types.String
	var lcname *types.String
	var ce *types.ClassEntry
	var func_ *ZendFunction
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			klass = fp.ParseZval()
			method_name = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if klass.IsObject() {
		ce = types.Z_OBJCE_P(klass)
	} else if klass.IsString() {
		if b.Assign(&ce, ZendLookupClass(klass.GetStr())) == nil {
			return_value.SetFalse()
			return
		}
	} else {
		return_value.SetFalse()
		return
	}
	lcname = ZendStringTolower(method_name)
	func_ = types.ZendHashFindPtr(ce.GetFunctionTable(), lcname)
	types.ZendStringReleaseEx(lcname, 0)
	if func_ != nil {

		/* Exclude shadow properties when checking a method on a specific class. Include
		 * them when checking an object, as method_exists() generally ignores visibility.
		 * TODO: Should we use EG(scope) for the object case instead? */

		types.ZVAL_BOOL(return_value, klass.IsObject() || !func_.IsPrivate() || func_.GetScope() == ce)
		return
	}
	if klass.IsObject() {
		var obj = klass.GetObj()
		func_ = types.Z_OBJ_HT_P(klass).GetGetMethod()(&obj, method_name, nil)
		if func_ != nil {
			if func_.IsCallViaTrampoline() {

				/* Returns true to the fake Closure's __invoke */

				types.ZVAL_BOOL(return_value, func_.GetScope() == ZendCeClosure && types.ZendStringEqualsLiteral(method_name, ZEND_INVOKE_FUNC_NAME))
				types.ZendStringReleaseEx(func_.GetFunctionName(), 0)
				ZendFreeTrampoline(func_)
				return
			}
			return_value.SetTrue()
			return
		}
	}
	return_value.SetFalse()
	return
}
func ZifPropertyExists(executeData *ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval
	var property *types.String
	var ce *types.ClassEntry
	var property_info *ZendPropertyInfo
	var property_z types.Zval
	if ZendParseParameters(executeData.NumArgs(), "zS", &object, &property) == types.FAILURE {
		return
	}
	if property == nil {
		return_value.SetFalse()
		return
	}
	if object.IsString() {
		ce = ZendLookupClass(object.GetStr())
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
	property_info = types.ZendHashFindPtr(ce.GetPropertiesInfo(), property)
	if property_info != nil && (!property_info.IsPrivate() || property_info.GetCe() == ce) {
		return_value.SetTrue()
		return
	}
	property_z.SetString(property)
	if object.IsObject() && types.Z_OBJ_HT(*object).GetHasProperty()(object, &property_z, 2, nil) != 0 {
		return_value.SetTrue()
		return
	}
	return_value.SetFalse()
	return
}
func ClassExistsImpl(executeData *ZendExecuteData, return_value *types.Zval, flags int, skip_flags int) {
	var name *types.String
	var lcname *types.String
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
		if name.GetVal()[0] == '\\' {

			/* Ignore leading "\" */

			lcname = types.ZendStringAlloc(name.GetLen()-1, 0)
			ZendStrTolowerCopy(lcname.GetVal(), name.GetVal()+1, name.GetLen()-1)
		} else {
			lcname = ZendStringTolower(name)
		}
		ce = types.ZendHashFindPtr(EG__().GetClassTable(), lcname)
		types.ZendStringReleaseEx(lcname, 0)
	} else {
		ce = ZendLookupClass(name)
	}
	if ce != nil {
		types.ZVAL_BOOL(return_value, (ce.GetCeFlags()&flags) == flags && !ce.HasCeFlags(skip_flags))
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifClassExists(executeData *ZendExecuteData, return_value *types.Zval) {
	ClassExistsImpl(executeData, return_value, ZEND_ACC_LINKED, ZEND_ACC_INTERFACE|ZEND_ACC_TRAIT)
}
func ZifInterfaceExists(executeData *ZendExecuteData, return_value *types.Zval) {
	ClassExistsImpl(executeData, return_value, ZEND_ACC_LINKED|ZEND_ACC_INTERFACE, 0)
}
func ZifTraitExists(executeData *ZendExecuteData, return_value *types.Zval) {
	ClassExistsImpl(executeData, return_value, ZEND_ACC_TRAIT, 0)
}
func ZifFunctionExists(executeData *ZendExecuteData, return_value *types.Zval) {
	var name *types.String
	var func_ *ZendFunction
	var lcname *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			name = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if name.GetVal()[0] == '\\' {

		/* Ignore leading "\" */

		lcname = types.ZendStringAlloc(name.GetLen()-1, 0)
		ZendStrTolowerCopy(lcname.GetVal(), name.GetVal()+1, name.GetLen()-1)
	} else {
		lcname = ZendStringTolower(name)
	}
	func_ = types.ZendHashFindPtr(EG__().GetFunctionTable(), lcname)
	types.ZendStringReleaseEx(lcname, 0)

	/*
	 * A bit of a hack, but not a bad one: we see if the handler of the function
	 * is actually one that displays "function is disabled" message.
	 */

	types.ZVAL_BOOL(return_value, func_ != nil && (func_.GetType() != ZEND_INTERNAL_FUNCTION || func_.GetInternalFunction().GetHandler() != ZifDisplayDisabledFunction))
	return
}
func ZifClassAlias(executeData *ZendExecuteData, return_value *types.Zval) {
	var class_name *types.String
	var alias_name *byte
	var ce *types.ClassEntry
	var alias_name_len int
	var autoload = 1
	if ZendParseParameters(executeData.NumArgs(), "Ss|b", &class_name, &alias_name, &alias_name_len, &autoload) == types.FAILURE {
		return
	}
	ce = ZendLookupClassEx(class_name, nil, b.Cond(autoload == 0, ZEND_FETCH_CLASS_NO_AUTOLOAD, 0))
	if ce != nil {
		if ce.GetType() == ZEND_USER_CLASS {
			if ZendRegisterClassAliasEx(alias_name, alias_name_len, ce, 0) == types.SUCCESS {
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
func ZifGetIncludedFiles(executeData *ZendExecuteData, return_value *types.Zval) {
	var entry *types.String
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	ArrayInit(return_value)
	var __ht = EG__().GetIncludedFiles()
	for _, _p := range __ht.foreachData() {
		var _z = _p.GetVal()

		entry = _p.GetKey()
		if entry != nil {
			AddNextIndexStr(return_value, entry.Copy())
		}
	}
}
func ZifTriggerError(executeData *ZendExecuteData, return_value *types.Zval) {
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
func ZifSetErrorHandler(executeData *ZendExecuteData, return_value *types.Zval) {
	var error_handler *types.Zval
	var error_type = faults.E_ALL
	if ZendParseParameters(executeData.NumArgs(), "z|l", &error_handler, &error_type) == types.FAILURE {
		return
	}
	if error_handler.GetType() != types.IS_NULL {
		if ZendIsCallable(error_handler, 0, nil) == 0 {
			var error_handler_name = ZendGetCallableName(error_handler)
			faults.Error(faults.E_WARNING, "%s() expects the argument (%s) to be a valid callback", GetActiveFunctionName(), b.CondF1(error_handler_name != nil, func() []byte { return error_handler_name.GetVal() }, "unknown"))
			types.ZendStringReleaseEx(error_handler_name, 0)
			return
		}
	}
	if EG__().GetUserErrorHandler().GetType() != types.IS_UNDEF {
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
func ZifRestoreErrorHandler(executeData *ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if EG__().GetUserErrorHandler().GetType() != types.IS_UNDEF {
		var zeh types.Zval
		types.ZVAL_COPY_VALUE(&zeh, EG__().GetUserErrorHandler())
		EG__().GetUserErrorHandler().SetUndef()
		ZvalPtrDtor(&zeh)
	}
	if ZendStackIsEmpty(EG__().GetUserErrorHandlers()) != 0 {
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
func ZifSetExceptionHandler(executeData *ZendExecuteData, return_value *types.Zval) {
	var exception_handler *types.Zval
	if ZendParseParameters(executeData.NumArgs(), "z", &exception_handler) == types.FAILURE {
		return
	}
	if exception_handler.GetType() != types.IS_NULL {
		if ZendIsCallable(exception_handler, 0, nil) == 0 {
			var exception_handler_name = ZendGetCallableName(exception_handler)
			faults.Error(faults.E_WARNING, "%s() expects the argument (%s) to be a valid callback", GetActiveFunctionName(), b.CondF1(exception_handler_name != nil, func() []byte { return exception_handler_name.GetVal() }, "unknown"))
			types.ZendStringReleaseEx(exception_handler_name, 0)
			return
		}
	}
	if EG__().GetUserExceptionHandler().GetType() != types.IS_UNDEF {
		types.ZVAL_COPY(return_value, EG__().GetUserExceptionHandler())
	}
	EG__().GetUserExceptionHandlers().Push(EG__().GetUserExceptionHandler())
	if exception_handler.IsNull() {
		EG__().GetUserExceptionHandler().SetUndef()
		return
	}
	types.ZVAL_COPY(EG__().GetUserExceptionHandler(), exception_handler)
}
func ZifRestoreExceptionHandler(executeData *ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if EG__().GetUserExceptionHandler().GetType() != types.IS_UNDEF {
		ZvalPtrDtor(EG__().GetUserExceptionHandler())
	}
	if ZendStackIsEmpty(EG__().GetUserExceptionHandlers()) != 0 {
		EG__().GetUserExceptionHandler().SetUndef()
	} else {
		var tmp *types.Zval = EG__().GetUserExceptionHandlers().Top()
		types.ZVAL_COPY_VALUE(EG__().GetUserExceptionHandler(), tmp)
		EG__().GetUserExceptionHandlers().DelTop()
	}
	return_value.SetTrue()
	return
}
func CopyClassOrInterfaceName(array *types.Zval, key *types.String, ce *types.ClassEntry) {
	if ce.GetRefcount() == 1 && !ce.IsImmutable() || SameName(key, ce.GetName()) != 0 {
		key = ce.GetName()
	}
	AddNextIndexStr(array, key.Copy())
}
func GetDeclaredClassImpl(executeData *ZendExecuteData, return_value *types.Zval, flags int, skip_flags int) {
	var key *types.String
	var ce *types.ClassEntry
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	ArrayInit(return_value)
	var __ht = EG__().GetClassTable()
	for _, _p := range __ht.foreachData() {
		var _z = _p.GetVal()

		key = _p.GetKey()
		ce = _z.GetPtr()
		if key != nil && key.GetVal()[0] != 0 && ce.HasCeFlags(flags) && !ce.HasCeFlags(skip_flags) {
			CopyClassOrInterfaceName(return_value, key, ce)
		}
	}
}
func ZifGetDeclaredTraits(executeData *ZendExecuteData, return_value *types.Zval) {
	GetDeclaredClassImpl(executeData, return_value, ZEND_ACC_TRAIT, 0)
}
func ZifGetDeclaredClasses(executeData *ZendExecuteData, return_value *types.Zval) {
	GetDeclaredClassImpl(executeData, return_value, ZEND_ACC_LINKED, ZEND_ACC_INTERFACE|ZEND_ACC_TRAIT)
}
func ZifGetDeclaredInterfaces(executeData *ZendExecuteData, return_value *types.Zval) {
	GetDeclaredClassImpl(executeData, return_value, ZEND_ACC_INTERFACE, 0)
}
func ZifGetDefinedFunctions(executeData *ZendExecuteData, return_value *types.Zval) {
	var internal types.Zval
	var user types.Zval
	var key *types.String
	var func_ *ZendFunction
	var exclude_disabled = 0
	if ZendParseParameters(executeData.NumArgs(), "|b", &exclude_disabled) == types.FAILURE {
		return
	}
	ArrayInit(&internal)
	ArrayInit(&user)
	ArrayInit(return_value)
	var __ht = EG__().GetFunctionTable()
	for _, _p := range __ht.foreachData() {
		var _z = _p.GetVal()

		key = _p.GetKey()
		func_ = _z.GetPtr()
		if key != nil && key.GetVal()[0] != 0 {
			if func_.GetType() == ZEND_INTERNAL_FUNCTION && (exclude_disabled == 0 || func_.GetInternalFunction().GetHandler() != ZifDisplayDisabledFunction) {
				AddNextIndexStr(&internal, key.Copy())
			} else if func_.GetType() == ZEND_USER_FUNCTION {
				AddNextIndexStr(&user, key.Copy())
			}
		}
	}
	return_value.GetArr().KeyAddNew("internal", &internal)
	return_value.GetArr().KeyAddNew("user", &user)
}
func ZifGetDefinedVars(executeData *ZendExecuteData, return_value *types.Zval) {
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
func ZifCreateFunction(executeData *ZendExecuteData, return_value *types.Zval) {
	var function_name *types.String
	var eval_code *byte
	var function_args *byte
	var function_code *byte
	var eval_code_length int
	var function_args_len int
	var function_code_len int
	var retval int
	var eval_name *byte
	if ZendParseParameters(executeData.NumArgs(), "ss", &function_args, &function_args_len, &function_code, &function_code_len) == types.FAILURE {
		return
	}
	eval_code = (*byte)(Emalloc(b.SizeOf("\"function \" LAMBDA_TEMP_FUNCNAME") + function_args_len + 2 + 2 + function_code_len))
	eval_code_length = b.SizeOf("\"function \" LAMBDA_TEMP_FUNCNAME \"(\"") - 1
	memcpy(eval_code, "function "+LAMBDA_TEMP_FUNCNAME+"(", eval_code_length)
	memcpy(eval_code+eval_code_length, function_args, function_args_len)
	eval_code_length += function_args_len
	eval_code[b.PostInc(&eval_code_length)] = ')'
	eval_code[b.PostInc(&eval_code_length)] = '{'
	memcpy(eval_code+eval_code_length, function_code, function_code_len)
	eval_code_length += function_code_len
	eval_code[b.PostInc(&eval_code_length)] = '}'
	eval_code[eval_code_length] = '0'
	eval_name = ZendMakeCompiledStringDescription("runtime-created function")
	retval = ZendEvalStringl(eval_code, eval_code_length, nil, eval_name)
	Efree(eval_code)
	Efree(eval_name)
	if retval == types.SUCCESS {
		var func_ *ZendOpArray
		var static_variables *types.Array
		func_ = types.ZendHashStrFindPtr(EG__().GetFunctionTable(), LAMBDA_TEMP_FUNCNAME, b.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		if func_ == nil {
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Unexpected inconsistency in create_function()")
			return_value.SetFalse()
			return
		}
		if func_.GetRefcount() != nil {
			func_.refcount++
		}
		static_variables = func_.GetStaticVariables()
		func_.SetStaticVariables(nil)
		types.ZendHashStrDel(EG__().GetFunctionTable(), LAMBDA_TEMP_FUNCNAME, b.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		func_.SetStaticVariables(static_variables)
		function_name = types.ZendStringAlloc(b.SizeOf("\"0lambda_\"")+MAX_LENGTH_OF_LONG, 0)
		function_name.GetVal()[0] = '0'
		for {
			function_name.SetLen(core.Snprintf(function_name.GetVal()+1, b.SizeOf("\"lambda_\"")+MAX_LENGTH_OF_LONG, "lambda_%d", b.PreInc(&(EG__().GetLambdaCount()))) + 1)
			if types.ZendHashAddPtr(EG__().GetFunctionTable(), function_name, func_) != nil {
				break
			}
		}
		return_value.SetString(function_name)
		return
	} else {
		types.ZendHashStrDel(EG__().GetFunctionTable(), LAMBDA_TEMP_FUNCNAME, b.SizeOf("LAMBDA_TEMP_FUNCNAME")-1)
		return_value.SetFalse()
		return
	}
}
func ZifGetResourceType(executeData *ZendExecuteData, return_value *types.Zval) {
	var resource_type *byte
	var z_resource_type *types.Zval
	if ZendParseParameters(executeData.NumArgs(), "r", &z_resource_type) == types.FAILURE {
		return
	}
	resource_type = ZendRsrcListGetRsrcType(z_resource_type.GetRes())
	if resource_type != nil {
		return_value.SetRawString(b.CastStrAuto(resource_type))
		return
	} else {
		return_value.SetRawString("Unknown")
		return
	}
}
func ZifGetResources(executeData *ZendExecuteData, return_value *types.Zval) {
	var type_ *types.String = nil
	var key *types.String
	var index ZendUlong
	var val *types.Zval
	if ZendParseParameters(executeData.NumArgs(), "|S", &type_) == types.FAILURE {
		return
	}
	if type_ == nil {
		ArrayInit(return_value)
		var __ht *types.Array = EG__().GetRegularList()
		for _, _p := range __ht.foreachData() {
			var _z = _p.GetVal()

			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			if key == nil {
				val.AddRefcount()
				return_value.GetArr().IndexAddNewH(index, val)
			}
		}
	} else if types.ZendStringEqualsLiteral(type_, "Unknown") {
		ArrayInit(return_value)
		var __ht *types.Array = EG__().GetRegularList()
		for _, _p := range __ht.foreachData() {
			var _z = _p.GetVal()

			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			if key == nil && types.Z_RES_TYPE_P(val) <= 0 {
				val.AddRefcount()
				return_value.GetArr().IndexAddNewH(index, val)
			}
		}
	} else {
		var id = ZendFetchListDtorId(type_.GetVal())
		if id <= 0 {
			faults.Error(faults.E_WARNING, "get_resources():  Unknown resource type '%s'", type_.GetVal())
			return_value.SetFalse()
			return
		}
		ArrayInit(return_value)
		var __ht *types.Array = EG__().GetRegularList()
		for _, _p := range __ht.foreachData() {
			var _z = _p.GetVal()

			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			if key == nil && types.Z_RES_TYPE_P(val) == id {
				val.AddRefcount()
				return_value.GetArr().IndexAddNewH(index, val)
			}
		}
	}
}
func AddZendextInfo(ext *ZendExtension, arg any) int {
	var name_array = (*types.Zval)(arg)
	AddNextIndexString(name_array, ext.GetName())
	return 0
}
func ZifGetLoadedExtensions(executeData *ZendExecuteData, return_value *types.Zval) {
	var zendext = 0
	if ZendParseParameters(executeData.NumArgs(), "|b", &zendext) == types.FAILURE {
		return
	}
	ArrayInit(return_value)
	if zendext != 0 {
		ZendExtensions.ApplyWithArgument(LlistApplyWithArgFuncT(AddZendextInfo), return_value)
	} else {
		var module *ZendModuleEntry
		var __ht = &ModuleRegistry
		for _, _p := range __ht.foreachData() {
			var _z = _p.GetVal()

			module = _z.GetPtr()
			AddNextIndexString(return_value, module.GetName())
		}
	}
}
func ZifGetDefinedConstants(executeData *ZendExecuteData, return_value *types.Zval) {
	var categorize = 0
	if ZendParseParameters(executeData.NumArgs(), "|b", &categorize) == types.FAILURE {
		return
	}
	ArrayInit(return_value)
	if categorize != 0 {
		var val *ZendConstant
		var module_number int
		var modules *types.Zval
		var const_val types.Zval
		var module_names **byte
		var module *ZendModuleEntry
		var i = 1
		modules = Ecalloc(ModuleRegistry.GetNNumOfElements()+2, b.SizeOf("zval"))
		module_names = Emalloc((ModuleRegistry.GetNNumOfElements() + 2) * b.SizeOf("char *"))
		module_names[0] = "internal"
		var __ht = &ModuleRegistry
		for _, _p := range __ht.foreachData() {
			var _z = _p.GetVal()

			module = _z.GetPtr()
			module_names[module.GetModuleNumber()] = (*byte)(module.GetName())
			i++
		}
		module_names[i] = "user"
		var __ht__1 = EG__().GetZendConstants()
		for _, _p := range __ht__1.foreachData() {
			var _z = _p.GetVal()

			val = _z.GetPtr()
			if val.GetName() == nil {

				/* skip special constants */

				continue

				/* skip special constants */

			}
			if ZEND_CONSTANT_MODULE_NUMBER(val) == PHP_USER_CONSTANT {
				module_number = i
			} else if ZEND_CONSTANT_MODULE_NUMBER(val) > i {

				/* should not happen */

				continue

				/* should not happen */

			} else {
				module_number = ZEND_CONSTANT_MODULE_NUMBER(val)
			}
			if modules[module_number].IsUndef() {
				ArrayInit(&modules[module_number])
				AddAssocZval(return_value, module_names[module_number], &modules[module_number])
			}
			types.ZVAL_COPY_OR_DUP(&const_val, val.Value())
			modules[module_number].GetArr().KeyAddNew(val.GetName().GetStr(), &const_val)
		}
		Efree(module_names)
		Efree(modules)
	} else {
		var constant *ZendConstant
		var const_val types.Zval
		var __ht = EG__().GetZendConstants()
		for _, _p := range __ht.foreachData() {
			var _z = _p.GetVal()

			constant = _z.GetPtr()
			if constant.GetName() == nil {

				/* skip special constants */

				continue

				/* skip special constants */

			}
			types.ZVAL_COPY_OR_DUP(&const_val, constant.Value())
			return_value.GetArr().KeyAddNew(constant.GetName().GetStr(), &const_val)
		}
	}
}
func DebugBacktraceGetArgs(call *ZendExecuteData, arg_array *types.Zval) {
	var num_args uint32 = call.NumArgs()
	if num_args != 0 {
		var i uint32 = 0
		var p = call.Arg(1)
		ArrayInitSize(arg_array, num_args)
		types.ZendHashRealInitPacked(arg_array.GetArr())

		fillScope := types.PackedFillStart(arg_array.GetArr())
		if call.GetFunc().GetType() == ZEND_USER_FUNCTION {
			var first_extra_arg = b.Min(num_args, call.GetFunc().GetOpArray().GetNumArgs())
			if (ZEND_CALL_INFO(call) & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {

				/* In case of attached symbol_table, values on stack may be invalid
				 * and we have to access them through symbol_table
				 * See: https://bugs.php.net/bug.php?id=73156
				 */

				var arg_name *types.String
				var arg *types.Zval
				for i < first_extra_arg {
					arg_name = call.GetFunc().GetOpArray().GetVars()[i]
					arg = types.ZendHashFindExInd(call.GetSymbolTable(), arg_name, 1)
					if arg != nil {
						if arg.IsRefcounted() {
							arg.AddRefcount()
						}
						fillScope.FillSet(arg)
					} else {
						fillScope.FillSetNull()
					}
					fillScope.FillNext()
					i++
				}
			} else {
				for i < first_extra_arg {
					if p.GetTypeInfo() != types.IS_UNDEF {
						if p.IsRefcounted() {
							p.AddRefcount()
						}
						fillScope.FillSet(p)
					} else {
						fillScope.FillSetNull()
					}
					fillScope.FillNext()
					p++
					i++
				}
			}
			p = call.VarNum(call.GetFunc().GetOpArray().GetLastVar() + call.GetFunc().GetOpArray().GetT())
		}
		for i < num_args {
			if p.GetTypeInfo() != types.IS_UNDEF {
				if p.IsRefcounted() {
					p.AddRefcount()
				}
				fillScope.FillSet(p)
			} else {
				fillScope.FillSetNull()
			}
			fillScope.FillNext()
			p++
			i++
		}
		fillScope.FillEnd()
		types.Z_ARRVAL_P(arg_array).SetNNumOfElements(num_args)
	} else {
		types.ZVAL_EMPTY_ARRAY(arg_array)
	}
}
func DebugPrintBacktraceArgs(arg_array *types.Zval) {
	var tmp *types.Zval
	var i = 0
	var __ht = arg_array.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z = _p.GetVal()

		tmp = _z
		if b.PostInc(&i) {
			ZEND_PUTS(", ")
		}
		ZendPrintFlatZvalR(tmp)
	}
}
func SkipInternalHandler(skip *ZendExecuteData) types.ZendBool {
	return !(skip.GetFunc() != nil && ZEND_USER_CODE(skip.GetFunc().GetCommonType())) && skip.GetPrevExecuteData() != nil && skip.GetPrevExecuteData().GetFunc() != nil && ZEND_USER_CODE(skip.GetPrevExecuteData().GetFunc().GetCommonType()) && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_ICALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_UCALL && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL_BY_NAME && skip.GetPrevExecuteData().GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL
}
func ZifDebugPrintBacktrace(executeData *ZendExecuteData, return_value *types.Zval) {
	var call *ZendExecuteData
	var ptr *ZendExecuteData
	var skip *ZendExecuteData
	var object *types.ZendObject
	var lineno int
	var frameno = 0
	var func_ *ZendFunction
	var function_name *byte
	var filename *byte
	var class_name *types.String = nil
	var call_type *byte
	var include_filename *byte = nil
	var arg_array types.Zval
	var indent = 0
	var options = 0
	var limit = 0
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
		class_name = nil
		call_type = nil
		arg_array.SetUndef()
		ptr = ZendGeneratorCheckPlaceholderFrame(ptr)
		skip = ptr

		/* skip internal handler */

		if SkipInternalHandler(skip) != 0 {
			skip = skip.GetPrevExecuteData()
		}
		if skip.GetFunc() != nil && ZEND_USER_CODE(skip.GetFunc().GetCommonType()) {
			filename = skip.GetFunc().GetOpArray().GetFilename().GetVal()
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
			filename = nil
			lineno = 0
		}

		/* $this may be passed into regular internal functions */

		if call.GetThis().IsObject() {
			object = call.GetThis().GetObj()
		} else {
			object = nil
		}
		if call.GetFunc() != nil {
			var zend_function_name *types.String
			func_ = call.GetFunc()
			if func_.GetScope() != nil && func_.GetScope().GetTraitAliases() != nil {
				zend_function_name = ZendResolveMethodName(b.CondF(object != nil, func() *types.ClassEntry { return object.GetCe() }, func() *types.ClassEntry { return func_.GetScope() }), func_)
			} else {
				zend_function_name = func_.GetFunctionName()
			}
			if zend_function_name != nil {
				function_name = zend_function_name.GetVal()
			} else {
				function_name = nil
			}
		} else {
			func_ = nil
			function_name = nil
		}
		if function_name != nil {
			if object != nil {
				if func_.GetScope() != nil {
					class_name = func_.GetScope().GetName()
				} else if object.GetHandlers().GetGetClassName() == ZendStdGetClassName {
					class_name = object.GetCe().GetName()
				} else {
					class_name = object.GetHandlers().GetGetClassName()(object)
				}
				call_type = "->"
			} else if func_.GetScope() != nil {
				class_name = func_.GetScope().GetName()
				call_type = "::"
			} else {
				class_name = nil
				call_type = nil
			}
			if func_.GetType() != ZEND_EVAL_CODE {
				if (options & DEBUG_BACKTRACE_IGNORE_ARGS) == 0 {
					DebugBacktraceGetArgs(call, &arg_array)
				}
			}
		} else {

			/* i know this is kinda ugly, but i'm trying to avoid extra cycles in the main execution loop */

			var build_filename_arg = 1
			if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetCommonType())) || ptr.GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL {

				/* can happen when calling eval from a custom sapi */

				function_name = "unknown"
				build_filename_arg = 0
			} else {
				switch ptr.GetOpline().GetExtendedValue() {
				case ZEND_EVAL:
					function_name = "eval"
					build_filename_arg = 0
				case ZEND_INCLUDE:
					function_name = "include"
				case ZEND_REQUIRE:
					function_name = "require"
				case ZEND_INCLUDE_ONCE:
					function_name = "include_once"
				case ZEND_REQUIRE_ONCE:
					function_name = "require_once"
				default:

					/* this can actually happen if you use debug_backtrace() in your error_handler and
					 * you're in the top-scope */

					function_name = "unknown"
					build_filename_arg = 0
				}
			}
			if build_filename_arg != 0 && include_filename != nil {
				ArrayInit(&arg_array)
				AddNextIndexString(&arg_array, (*byte)(include_filename))
			}
			call_type = nil
		}
		ZendPrintf("#%-2d ", indent)
		if class_name != nil {
			ZEND_PUTS(class_name.GetStr())
			ZEND_PUTS(call_type)
			if object != nil && func_.GetScope() == nil && object.GetHandlers().GetGetClassName() != ZendStdGetClassName {
				types.ZendStringReleaseEx(class_name, 0)
			}
		}
		ZendPrintf("%s(", function_name)
		if arg_array.GetType() != types.IS_UNDEF {
			DebugPrintBacktraceArgs(&arg_array)
			ZvalPtrDtor(&arg_array)
		}
		if filename != nil {
			ZendPrintf(") called at [%s:%d]\n", filename, lineno)
		} else {
			var prev_call = skip
			var prev = skip.GetPrevExecuteData()
			for prev != nil {
				if prev_call != nil && prev_call.GetFunc() != nil && !(ZEND_USER_CODE(prev_call.GetFunc().GetCommonType())) {
					prev = nil
					break
				}
				if prev.GetFunc() != nil && ZEND_USER_CODE(prev.GetFunc().GetCommonType()) {
					ZendPrintf(") called at [%s:%d]\n", prev.GetFunc().GetOpArray().GetFilename().GetVal(), prev.GetOpline().GetLineno())
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
	var object *types.ZendObject
	var lineno int
	var frameno = 0
	var func_ *ZendFunction
	var function_name *types.String
	var filename *types.String
	var include_filename *types.String = nil
	var stack_frame types.Zval
	var tmp types.Zval
	ArrayInit(return_value)
	if !(b.Assign(&ptr, CurrEX())) {
		return
	}
	if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetCommonType())) {
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

			if ptr.GetFunc() != nil && ZEND_USER_CODE(ptr.GetFunc().GetCommonType()) && ptr.GetOpline().GetOpcode() == ZEND_NEW {
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
		if skip.GetFunc() != nil && ZEND_USER_CODE(skip.GetFunc().GetCommonType()) {
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
			tmp.SetStringCopy(filename)
			stack_frame.GetArr().KeyAddNew(types.ZSTR_FILE.GetStr(), &tmp)
			tmp.SetLong(lineno)
			stack_frame.GetArr().KeyAddNew(types.ZSTR_LINE.GetStr(), &tmp)
		} else {
			var prev_call = skip
			var prev = skip.GetPrevExecuteData()
			for prev != nil {
				if prev_call != nil && prev_call.GetFunc() != nil && !(ZEND_USER_CODE(prev_call.GetFunc().GetCommonType())) && !prev_call.GetFunc().IsCallViaTrampoline() {
					break
				}
				if prev.GetFunc() != nil && ZEND_USER_CODE(prev.GetFunc().GetCommonType()) {
					tmp.SetStringCopy(prev.GetFunc().GetOpArray().GetFilename())
					stack_frame.GetArr().KeyAddNew(types.ZSTR_FILE.GetStr(), &tmp)
					tmp.SetLong(prev.GetOpline().GetLineno())
					stack_frame.GetArr().KeyAddNew(types.ZSTR_LINE.GetStr(), &tmp)
					break
				}
				prev_call = prev
				prev = prev.GetPrevExecuteData()
			}
			filename = nil
		}

		/* $this may be passed into regular internal functions */

		if call != nil && call.GetThis().IsObject() {
			object = call.GetThis().GetObj()
		} else {
			object = nil
		}
		if call != nil && call.GetFunc() != nil {
			func_ = call.GetFunc()
			if func_.GetScope() != nil && func_.GetScope().GetTraitAliases() != nil {
				function_name = ZendResolveMethodName(b.CondF(object != nil, func() *types.ClassEntry { return object.GetCe() }, func() *types.ClassEntry { return func_.GetScope() }), func_)
			} else {
				function_name = func_.GetFunctionName()
			}
		} else {
			func_ = nil
			function_name = nil
		}
		if function_name != nil {
			tmp.SetStringCopy(function_name)
			stack_frame.GetArr().KeyAddNew(types.ZSTR_FUNCTION.GetStr(), &tmp)
			if object != nil {
				if func_.GetScope() != nil {
					tmp.SetStringCopy(func_.GetScope().GetName())
				} else if object.GetHandlers().GetGetClassName() == ZendStdGetClassName {
					tmp.SetStringCopy(object.GetCe().GetName())
				} else {
					tmp.SetString(object.GetHandlers().GetGetClassName()(object))
				}
				stack_frame.GetArr().KeyAddNew(types.ZSTR_CLASS.GetStr(), &tmp)
				if (options & DEBUG_BACKTRACE_PROVIDE_OBJECT) != 0 {
					tmp.SetObject(object)
					stack_frame.GetArr().KeyAddNew(types.ZSTR_OBJECT.GetStr(), &tmp)
					tmp.AddRefcount()
				}
				tmp.SetInternedString(types.ZSTR_OBJECT_OPERATOR)
				stack_frame.GetArr().KeyAddNew(types.ZSTR_TYPE.GetStr(), &tmp)
			} else if func_.GetScope() != nil {
				tmp.SetStringCopy(func_.GetScope().GetName())
				stack_frame.GetArr().KeyAddNew(types.ZSTR_CLASS.GetStr(), &tmp)
				tmp.SetInternedString(types.ZSTR_PAAMAYIM_NEKUDOTAYIM)
				stack_frame.GetArr().KeyAddNew(types.ZSTR_TYPE.GetStr(), &tmp)
			}
			if (options&DEBUG_BACKTRACE_IGNORE_ARGS) == 0 && func_.GetType() != ZEND_EVAL_CODE {
				DebugBacktraceGetArgs(call, &tmp)
				stack_frame.GetArr().KeyAddNew(types.ZSTR_ARGS.GetStr(), &tmp)
			}
		} else {

			/* i know this is kinda ugly, but i'm trying to avoid extra cycles in the main execution loop */

			var build_filename_arg = 1
			var pseudo_function_name *types.String
			if ptr.GetFunc() == nil || !(ZEND_USER_CODE(ptr.GetFunc().GetCommonType())) || ptr.GetOpline().GetOpcode() != ZEND_INCLUDE_OR_EVAL {

				/* can happen when calling eval from a custom sapi */

				pseudo_function_name = types.ZSTR_UNKNOWN
				build_filename_arg = 0
			} else {
				switch ptr.GetOpline().GetExtendedValue() {
				case ZEND_EVAL:
					pseudo_function_name = types.ZSTR_EVAL
					build_filename_arg = 0
				case ZEND_INCLUDE:
					pseudo_function_name = types.ZSTR_INCLUDE
				case ZEND_REQUIRE:
					pseudo_function_name = types.ZSTR_REQUIRE
				case ZEND_INCLUDE_ONCE:
					pseudo_function_name = types.ZSTR_INCLUDE_ONCE
				case ZEND_REQUIRE_ONCE:
					pseudo_function_name = types.ZSTR_REQUIRE_ONCE
				default:

					/* this can actually happen if you use debug_backtrace() in your error_handler and
					 * you're in the top-scope */

					pseudo_function_name = types.ZSTR_UNKNOWN
					build_filename_arg = 0
				}
			}
			if build_filename_arg != 0 && include_filename != nil {
				var arg_array types.Zval
				ArrayInit(&arg_array)

				/* include_filename always points to the last filename of the last last called-function.
				   if we have called include in the frame above - this is the file we have included.
				*/

				tmp.SetStringCopy(include_filename)
				arg_array.GetArr().NextIndexInsertNew(&tmp)
				stack_frame.GetArr().KeyAddNew(types.ZSTR_ARGS.GetStr(), &arg_array)
			}
			tmp.SetInternedString(pseudo_function_name)
			stack_frame.GetArr().KeyAddNew(types.ZSTR_FUNCTION.GetStr(), &tmp)
		}
		return_value.GetArr().NextIndexInsertNew(&stack_frame)
		include_filename = filename
		call = skip
		ptr = skip.GetPrevExecuteData()
	}
}
func ZifDebugBacktrace(executeData *ZendExecuteData, return_value *types.Zval) {
	var options = DEBUG_BACKTRACE_PROVIDE_OBJECT
	var limit = 0
	if ZendParseParameters(executeData.NumArgs(), "|ll", &options, &limit) == types.FAILURE {
		return
	}
	ZendFetchDebugBacktrace(return_value, 1, options, limit)
}
func ZifExtensionLoaded(executeData *ZendExecuteData, return_value *types.Zval) {
	var extension_name *types.String
	var lcname *types.String
	if ZendParseParameters(executeData.NumArgs(), "S", &extension_name) == types.FAILURE {
		return
	}
	lcname = ZendStringTolower(extension_name)
	if types.ZendHashExists(&ModuleRegistry, lcname) != 0 {
		return_value.SetTrue()
	} else {
		return_value.SetFalse()
	}
	types.ZendStringReleaseEx(lcname, 0)
}
func ZifGetExtensionFuncs(executeData *ZendExecuteData, return_value *types.Zval) {
	var extension_name *types.String
	var lcname *types.String
	var array int
	var module *ZendModuleEntry
	var zif *ZendFunction
	if ZendParseParameters(executeData.NumArgs(), "S", &extension_name) == types.FAILURE {
		return
	}
	if strncasecmp(extension_name.GetVal(), "zend", b.SizeOf("\"zend\"")) {
		lcname = ZendStringTolower(extension_name)
		module = types.ZendHashFindPtr(&ModuleRegistry, lcname)
		types.ZendStringReleaseEx(lcname, 0)
	} else {
		module = types.ZendHashStrFindPtr(&ModuleRegistry, "core", b.SizeOf("\"core\"")-1)
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
	var __ht = CG__().GetFunctionTable()
	for _, _p := range __ht.foreachData() {
		var _z = _p.GetVal()

		zif = _z.GetPtr()
		if zif.GetCommonType() == ZEND_INTERNAL_FUNCTION && zif.GetInternalFunction().GetModule() == module {
			if array == 0 {
				ArrayInit(return_value)
				array = 1
			}
			AddNextIndexStr(return_value, zif.GetFunctionName().Copy())
		}
	}
	if array == 0 {
		return_value.SetFalse()
		return
	}
}
