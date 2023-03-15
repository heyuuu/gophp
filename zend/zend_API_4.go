// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ObjectPropertiesInit(object *ZendObject, class_type *ZendClassEntry) {
	object.SetProperties(nil)
	_objectPropertiesInit(object, class_type)
}
func ObjectPropertiesInitEx(object *ZendObject, properties *HashTable) {
	object.SetProperties(properties)
	if object.GetCe().GetDefaultPropertiesCount() != 0 {
		var prop *Zval
		var key *ZendString
		var property_info *ZendPropertyInfo
		var __ht *HashTable = properties
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			key = _p.GetKey()
			prop = _z
			property_info = ZendGetPropertyInfo(object.GetCe(), key, 1)
			if property_info != ZEND_WRONG_PROPERTY_INFO && property_info != nil && !property_info.IsStatic() {
				var slot *Zval = OBJ_PROP(object, property_info.GetOffset())
				if property_info.GetType() != 0 {
					var tmp Zval
					ZVAL_COPY_VALUE(&tmp, prop)
					if ZendVerifyPropertyType(property_info, &tmp, 0) == 0 {
						continue
					}
					ZVAL_COPY_VALUE(slot, &tmp)
				} else {
					ZVAL_COPY_VALUE(slot, prop)
				}
				prop.SetIndirect(slot)
			}
		}
	}
}
func ObjectPropertiesLoad(object *ZendObject, properties *HashTable) {
	var prop *Zval
	var tmp Zval
	var key *ZendString
	var h ZendLong
	var property_info *ZendPropertyInfo
	var __ht *HashTable = properties
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		h = _p.GetH()
		key = _p.GetKey()
		prop = _z
		if key != nil {
			if key.GetVal()[0] == '0' {
				var class_name *byte
				var prop_name *byte
				var prop_name_len int
				if ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, &prop_name_len) == SUCCESS {
					var pname *ZendString = ZendStringInit(prop_name, prop_name_len, 0)
					var prev_scope *ZendClassEntry = EG__().GetFakeScope()
					if class_name != nil && class_name[0] != '*' {
						var cname *ZendString = ZendStringInit(class_name, strlen(class_name), 0)
						EG__().SetFakeScope(ZendLookupClass(cname))
						ZendStringReleaseEx(cname, 0)
					}
					property_info = ZendGetPropertyInfo(object.GetCe(), pname, 1)
					ZendStringReleaseEx(pname, 0)
					EG__().SetFakeScope(prev_scope)
				} else {
					property_info = ZEND_WRONG_PROPERTY_INFO
				}
			} else {
				property_info = ZendGetPropertyInfo(object.GetCe(), key, 1)
			}
			if property_info != ZEND_WRONG_PROPERTY_INFO && property_info != nil && !property_info.IsStatic() {
				var slot *Zval = OBJ_PROP(object, property_info.GetOffset())
				ZvalPtrDtor(slot)
				ZVAL_COPY_VALUE(slot, prop)
				ZvalAddRef(slot)
				if object.GetProperties() != nil {
					tmp.SetIndirect(slot)
					object.GetProperties().KeyUpdate(key.GetStr(), &tmp)
				}
			} else {
				if object.GetProperties() == nil {
					RebuildObjectProperties(object)
				}
				prop = object.GetProperties().KeyUpdate(key.GetStr(), prop)
				ZvalAddRef(prop)
			}
		} else {
			if object.GetProperties() == nil {
				RebuildObjectProperties(object)
			}
			prop = object.GetProperties().IndexUpdateH(h, prop)
			ZvalAddRef(prop)
		}
	}
}
func _objectAndPropertiesInit(arg *Zval, class_type *ZendClassEntry, properties *HashTable) int {
	if class_type.HasCeFlags(ZEND_ACC_INTERFACE | ZEND_ACC_TRAIT | ZEND_ACC_IMPLICIT_ABSTRACT_CLASS | ZEND_ACC_EXPLICIT_ABSTRACT_CLASS) {
		if class_type.IsInterface() {
			ZendThrowError(nil, "Cannot instantiate interface %s", class_type.GetName().GetVal())
		} else if class_type.IsTrait() {
			ZendThrowError(nil, "Cannot instantiate trait %s", class_type.GetName().GetVal())
		} else {
			ZendThrowError(nil, "Cannot instantiate abstract class %s", class_type.GetName().GetVal())
		}
		arg.SetNull()
		arg.SetObj(nil)
		return FAILURE
	}
	if !class_type.IsConstantsUpdated() {
		if ZendUpdateClassConstants(class_type) != SUCCESS {
			arg.SetNull()
			arg.SetObj(nil)
			return FAILURE
		}
	}
	if class_type.GetCreateObject() == nil {
		var obj *ZendObject = ZendObjectsNew(class_type)
		arg.SetObject(obj)
		if properties != nil {
			ObjectPropertiesInitEx(obj, properties)
		} else {
			_objectPropertiesInit(obj, class_type)
		}
	} else {
		arg.SetObject(class_type.GetCreateObject()(class_type))
	}
	return SUCCESS
}
func ObjectAndPropertiesInit(arg *Zval, class_type *ZendClassEntry, properties *HashTable) int {
	return _objectAndPropertiesInit(arg, class_type, properties)
}
func ObjectInitEx(arg *Zval, class_type *ZendClassEntry) int {
	return _objectAndPropertiesInit(arg, class_type, nil)
}
func ObjectInit(arg *Zval) int {
	arg.SetObject(ZendObjectsNew(ZendStandardClassDef))
	return SUCCESS
}
func AddAssocLongEx(arg *Zval, key string, n ZendLong) int {
	var tmp Zval
	tmp.SetLong(n)
	arg.GetArr().SymtableUpdate(key, &tmp)
	return SUCCESS
}
func AddAssocNullEx(arg *Zval, key string) int {
	var tmp Zval
	tmp.SetNull()
	arg.GetArr().SymtableUpdate(key, &tmp)
	return SUCCESS
}
func AddAssocBoolEx(arg *Zval, key string, b int) int {
	var tmp Zval
	tmp.SetBool(b != 0)
	arg.GetArr().SymtableUpdate(key, &tmp)
	return SUCCESS
}
func AddAssocDoubleEx(arg *Zval, key string, d float64) int {
	var tmp Zval
	tmp.SetDouble(d)
	arg.GetArr().SymtableUpdate(key, &tmp)
	return SUCCESS
}
func AddAssocStrEx(arg *Zval, key string, str string) int {
	arg.GetArr().SymtableUpdate(key, NewZvalString(str))
	return SUCCESS
}
func AddAssocStringlEx(arg *Zval, key string, str string) int {
	arg.GetArr().SymtableUpdate(key, NewZvalString(str))
	return SUCCESS
}
func AddAssocZvalEx(arg *Zval, key string, value *Zval) int {
	arg.GetArr().SymtableUpdate(key, value)
	return SUCCESS
}
func AddIndexLong(arg *Zval, index ZendUlong, n ZendLong) int {
	var tmp Zval
	tmp.SetLong(n)
	arg.GetArr().IndexUpdateH(index, &tmp)
	return SUCCESS
}
func AddIndexBool(arg *Zval, index ZendUlong, b int) int {
	var tmp Zval
	tmp.SetBool(b != 0)
	arg.GetArr().IndexUpdateH(index, &tmp)
	return SUCCESS
}
func AddIndexResource(arg *Zval, index ZendUlong, r *ZendResource) int {
	var tmp Zval
	tmp.SetResource(r)
	arg.GetArr().IndexUpdateH(index, &tmp)
	return SUCCESS
}
func AddIndexDouble(arg *Zval, index ZendUlong, d float64) int {
	var tmp Zval
	tmp.SetDouble(d)
	arg.GetArr().IndexUpdateH(index, &tmp)
	return SUCCESS
}
func AddIndexStr(arg *Zval, index ZendUlong, str *ZendString) int {
	zv := NewZvalString(str.GetStr())
	arg.GetArr().IndexUpdateH(index, zv)
	return SUCCESS
}
func AddIndexString(arg *Zval, index ZendUlong, str *byte) int {
	zv := NewZvalString(b.CastStrAuto(str))
	arg.GetArr().IndexUpdateH(index, zv)
	return SUCCESS
}
func AddIndexStringl(arg *Zval, index ZendUlong, str *byte, length int) int {
	zv := NewZvalString(b.CastStr(str, length))
	arg.GetArr().IndexUpdateH(index, zv)
	return SUCCESS
}
func AddNextIndexLong(arg *Zval, n ZendLong) int {
	if arg.GetArr().NextIndexInsert(NewZvalLong(n)) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexNull(arg *Zval) int {
	if arg.GetArr().NextIndexInsert(NewZvalNull()) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexBool(arg *Zval, b int) int {
	var tmp Zval
	tmp.SetBool(b != 0)
	if arg.GetArr().NextIndexInsert(&tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexResource(arg *Zval, r *ZendResource) int {
	var tmp Zval
	tmp.SetResource(r)
	if arg.GetArr().NextIndexInsert(&tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexDouble(arg *Zval, d float64) int {
	var tmp Zval
	tmp.SetDouble(d)
	if arg.GetArr().NextIndexInsert(&tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexStr(arg *Zval, str *ZendString) int {
	var tmp Zval
	tmp.SetString(str)
	if arg.GetArr().NextIndexInsert(&tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexString(arg *Zval, str *byte) int {
	var tmp Zval
	tmp.SetRawString(b.CastStrAuto(str))
	if arg.GetArr().NextIndexInsert(&tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddNextIndexStringl(arg *Zval, str *byte, length int) int {
	var tmp Zval
	tmp.SetRawString(b.CastStr(str, length))
	if arg.GetArr().NextIndexInsert(&tmp) != nil {
		return SUCCESS
	} else {
		return FAILURE
	}
}
func ArraySetZvalKey(ht *HashTable, key *Zval, value *Zval) int {
	var result *Zval
	switch key.GetType() {
	case IS_STRING:
		result = ht.SymtableUpdate(key.GetStr().GetStr(), value)
		break
	case IS_NULL:
		result = ht.SymtableUpdate(ZSTR_EMPTY_ALLOC().GetStr(), value)
		break
	case IS_RESOURCE:
		ZendError(E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", Z_RES_HANDLE_P(key), Z_RES_HANDLE_P(key))
		result = ht.IndexUpdate(Z_RES_HANDLE_P(key), value)
		break
	case IS_FALSE:
		result = ht.IndexUpdateH(0, value)
		break
	case IS_TRUE:
		result = ht.IndexUpdateH(1, value)
		break
	case IS_LONG:
		result = ht.IndexUpdate(key.GetLval(), value)
		break
	case IS_DOUBLE:
		result = ht.IndexUpdate(ZendDvalToLval(key.GetDval()), value)
		break
	default:
		ZendError(E_WARNING, "Illegal offset type")
		result = nil
	}
	if result != nil {
		result.TryAddRefcount()
		return SUCCESS
	} else {
		return FAILURE
	}
}
func AddPropertyLongEx(arg *Zval, key string, n ZendLong) int {
	return AddPropertyZvalEx(arg, key, NewZvalLong(n))
}
func AddPropertyNullEx(arg *Zval, key string) int {
	return AddPropertyZvalEx(arg, key, NewZvalNull())
}
func AddPropertyResourceEx(arg *Zval, key string, r *ZendResource) int {
	return AddPropertyZvalEx(arg, key, NewZvalResource(r))
}
func AddPropertyStrEx(arg *Zval, key string, str string) int {
	return AddPropertyZvalEx(arg, key, NewZvalString(str))
}
func AddPropertyZvalEx(arg *Zval, key string, value *Zval) int {
	zKey := NewZvalString(key)
	Z_OBJ_HT(*arg).GetWriteProperty()(arg, zKey, value, nil)
	return SUCCESS
}
func ZendStartupModuleEx(module *ZendModuleEntry) int {
	var name_len int
	var lcname *ZendString
	if module.GetModuleStarted() != 0 {
		return SUCCESS
	}
	module.SetModuleStarted(1)

	/* Check module dependencies */

	if module.GetDeps() != nil {
		var dep *ZendModuleDep = module.GetDeps()
		for dep.GetName() != nil {
			if dep.GetType() == MODULE_DEP_REQUIRED {
				var req_mod *ZendModuleEntry
				name_len = strlen(dep.GetName())
				lcname = ZendStringAlloc(name_len, 0)
				ZendStrTolowerCopy(lcname.GetVal(), dep.GetName(), name_len)
				if b.Assign(&req_mod, ZendHashFindPtr(&ModuleRegistry, lcname)) == nil || req_mod.GetModuleStarted() == 0 {
					ZendStringEfree(lcname)

					/* TODO: Check version relationship */

					ZendError(E_CORE_WARNING, "Cannot load module '%s' because required module '%s' is not loaded", module.GetName(), dep.GetName())
					module.SetModuleStarted(0)
					return FAILURE
				}
				ZendStringEfree(lcname)
			}
			dep++
		}
	}

	/* Initialize module globals */

	if module.GetGlobalsSize() != 0 {
		if module.GetGlobalsCtor() != nil {
			module.GetGlobalsCtor()(module.GetGlobalsPtr())
		}
	}
	if module.GetModuleStartupFunc() != nil {
		EG__().SetCurrentModule(module)
		if module.GetModuleStartupFunc()(module.GetType(), module.GetModuleNumber()) == FAILURE {
			ZendErrorNoreturn(E_CORE_ERROR, "Unable to start %s module", module.GetName())
			EG__().SetCurrentModule(nil)
			return FAILURE
		}
		EG__().SetCurrentModule(nil)
	}
	return SUCCESS
}
