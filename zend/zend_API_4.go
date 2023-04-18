package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ObjectPropertiesInit(object *types2.ZendObject, class_type *types2.ClassEntry) {
	object.SetProperties(nil)
	_objectPropertiesInit(object, class_type)
}
func ObjectPropertiesInitEx(object *types2.ZendObject, properties *types2.Array) {
	object.SetProperties(properties)
	if object.GetCe().GetDefaultPropertiesCount() != 0 {
		var prop *types2.Zval
		var key *types2.String
		var property_info *ZendPropertyInfo
		var __ht *types2.Array = properties
		for _, _p := range __ht.ForeachData() {
			var _z *types2.Zval = _p.GetVal()

			key = _p.GetKey()
			prop = _z
			property_info = ZendGetPropertyInfo(object.GetCe(), key, 1)
			if property_info != ZEND_WRONG_PROPERTY_INFO && property_info != nil && !property_info.IsStatic() {
				var slot *types2.Zval = OBJ_PROP(object, property_info.GetOffset())
				if property_info.GetType() != 0 {
					var tmp types2.Zval
					types2.ZVAL_COPY_VALUE(&tmp, prop)
					if ZendVerifyPropertyType(property_info, &tmp, 0) == 0 {
						continue
					}
					types2.ZVAL_COPY_VALUE(slot, &tmp)
				} else {
					slot.CopyValueFrom(prop)
				}
				prop.SetIndirect(slot)
			}
		}
	}
}
func ObjectPropertiesLoad(object *types2.ZendObject, properties *types2.Array) {
	var prop *types2.Zval
	var tmp types2.Zval
	var key *types2.String
	var h ZendLong
	var property_info *ZendPropertyInfo
	var __ht *types2.Array = properties
	for _, _p := range __ht.ForeachData() {
		var _z *types2.Zval = _p.GetVal()

		h = _p.GetH()
		key = _p.GetKey()
		prop = _z
		if key != nil {
			if key.GetStr()[0] == '0' {
				var class_name *byte
				var prop_name *byte
				var prop_name_len int
				if ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, &prop_name_len) == types2.SUCCESS {
					var pname *types2.String = types2.NewString(b.CastStr(prop_name, prop_name_len))
					var prev_scope *types2.ClassEntry = EG__().GetFakeScope()
					if class_name != nil && class_name[0] != '*' {
						var cname *types2.String = types2.NewString(class_name)
						EG__().SetFakeScope(ZendLookupClass(cname))
						// types.ZendStringReleaseEx(cname, 0)
					}
					property_info = ZendGetPropertyInfo(object.GetCe(), pname, 1)
					// types.ZendStringReleaseEx(pname, 0)
					EG__().SetFakeScope(prev_scope)
				} else {
					property_info = ZEND_WRONG_PROPERTY_INFO
				}
			} else {
				property_info = ZendGetPropertyInfo(object.GetCe(), key, 1)
			}
			if property_info != ZEND_WRONG_PROPERTY_INFO && property_info != nil && !property_info.IsStatic() {
				var slot *types2.Zval = OBJ_PROP(object, property_info.GetOffset())
				// ZvalPtrDtor(slot)
				slot.CopyValueFrom(prop)
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
			prop = object.GetProperties().IndexUpdate(h, prop)
			ZvalAddRef(prop)
		}
	}
}
func _objectAndPropertiesInit(arg *types2.Zval, class_type *types2.ClassEntry, properties *types2.Array) int {
	if class_type.HasCeFlags(AccInterface | AccTrait | AccImplicitAbstractClass | AccExplicitAbstractClass) {
		if class_type.IsInterface() {
			faults.ThrowError(nil, "Cannot instantiate interface %s", class_type.GetName().GetVal())
		} else if class_type.IsTrait() {
			faults.ThrowError(nil, "Cannot instantiate trait %s", class_type.GetName().GetVal())
		} else {
			faults.ThrowError(nil, "Cannot instantiate abstract class %s", class_type.GetName().GetVal())
		}
		arg.SetNull()
		return types2.FAILURE
	}
	if !class_type.IsConstantsUpdated() {
		if ZendUpdateClassConstants(class_type) != types2.SUCCESS {
			arg.SetNull()
			return types2.FAILURE
		}
	}
	if class_type.GetCreateObject() == nil {
		var obj *types2.ZendObject = ZendObjectsNew(class_type)
		arg.SetObject(obj)
		if properties != nil {
			ObjectPropertiesInitEx(obj, properties)
		} else {
			_objectPropertiesInit(obj, class_type)
		}
	} else {
		arg.SetObject(class_type.GetCreateObject()(class_type))
	}
	return types2.SUCCESS
}
func ObjectAndPropertiesInit(arg *types2.Zval, class_type *types2.ClassEntry, properties *types2.Array) int {
	return _objectAndPropertiesInit(arg, class_type, properties)
}
func ObjectInitEx(arg *types2.Zval, class_type *types2.ClassEntry) int {
	return _objectAndPropertiesInit(arg, class_type, nil)
}
func ObjectInit(arg *types2.Zval) int {
	arg.SetObject(ZendObjectsNew(ZendStandardClassDef))
	return types2.SUCCESS
}
func AddAssocLongEx(arg *types2.Zval, key string, n ZendLong) int {
	var tmp types2.Zval
	tmp.SetLong(n)
	arg.Array().SymtableUpdate(key, &tmp)
	return types2.SUCCESS
}
func AddAssocNullEx(arg *types2.Zval, key string) int {
	var tmp types2.Zval
	tmp.SetNull()
	arg.Array().SymtableUpdate(key, &tmp)
	return types2.SUCCESS
}
func AddAssocBoolEx(arg *types2.Zval, key string, b int) int {
	var tmp types2.Zval
	tmp.SetBool(b != 0)
	arg.Array().SymtableUpdate(key, &tmp)
	return types2.SUCCESS
}
func AddAssocDoubleEx(arg *types2.Zval, key string, d float64) int {
	var tmp types2.Zval
	tmp.SetDouble(d)
	arg.Array().SymtableUpdate(key, &tmp)
	return types2.SUCCESS
}
func AddAssocStrEx(arg *types2.Zval, key string, str string) int {
	arg.Array().SymtableUpdate(key, types2.NewZvalString(str))
	return types2.SUCCESS
}
func AddAssocStringlEx(arg *types2.Zval, key string, str string) int {
	arg.Array().SymtableUpdate(key, types2.NewZvalString(str))
	return types2.SUCCESS
}
func AddAssocZvalEx(arg *types2.Zval, key string, value *types2.Zval) int {
	arg.Array().SymtableUpdate(key, value)
	return types2.SUCCESS
}
func AddIndexLong(arg *types2.Zval, index ZendUlong, n ZendLong) int {
	var tmp types2.Zval
	tmp.SetLong(n)
	arg.Array().IndexUpdate(index, &tmp)
	return types2.SUCCESS
}
func AddIndexBool(arg *types2.Zval, index ZendUlong, b int) int {
	var tmp types2.Zval
	tmp.SetBool(b != 0)
	arg.Array().IndexUpdate(index, &tmp)
	return types2.SUCCESS
}
func AddIndexResource(arg *types2.Zval, index ZendUlong, r *types2.ZendResource) int {
	var tmp types2.Zval
	tmp.SetResource(r)
	arg.Array().IndexUpdate(index, &tmp)
	return types2.SUCCESS
}
func AddIndexDouble(arg *types2.Zval, index ZendUlong, d float64) int {
	var tmp types2.Zval
	tmp.SetDouble(d)
	arg.Array().IndexUpdate(index, &tmp)
	return types2.SUCCESS
}
func AddIndexStr(arg *types2.Zval, index ZendUlong, str *types2.String) int {
	zv := types2.NewZvalString(str.GetStr())
	arg.Array().IndexUpdate(index, zv)
	return types2.SUCCESS
}
func AddIndexString(arg *types2.Zval, index ZendUlong, str *byte) int {
	zv := types2.NewZvalString(b.CastStrAuto(str))
	arg.Array().IndexUpdate(index, zv)
	return types2.SUCCESS
}
func AddIndexStringl(arg *types2.Zval, index ZendUlong, str *byte, length int) int {
	zv := types2.NewZvalString(b.CastStr(str, length))
	arg.Array().IndexUpdate(index, zv)
	return types2.SUCCESS
}
func AddNextIndexLong(arg *types2.Zval, n ZendLong) int {
	if arg.Array().NextIndexInsert(types2.NewZvalLong(n)) != nil {
		return types2.SUCCESS
	} else {
		return types2.FAILURE
	}
}
func AddNextIndexNull(arg *types2.Zval) int {
	if arg.Array().NextIndexInsert(types2.NewZvalNull()) != nil {
		return types2.SUCCESS
	} else {
		return types2.FAILURE
	}
}
func AddNextIndexBool(arg *types2.Zval, b int) int {
	var tmp types2.Zval
	tmp.SetBool(b != 0)
	if arg.Array().NextIndexInsert(&tmp) != nil {
		return types2.SUCCESS
	} else {
		return types2.FAILURE
	}
}
func AddNextIndexResource(arg *types2.Zval, r *types2.ZendResource) int {
	var tmp types2.Zval
	tmp.SetResource(r)
	if arg.Array().NextIndexInsert(&tmp) != nil {
		return types2.SUCCESS
	} else {
		return types2.FAILURE
	}
}
func AddNextIndexDouble(arg *types2.Zval, d float64) int {
	var tmp types2.Zval
	tmp.SetDouble(d)
	if arg.Array().NextIndexInsert(&tmp) != nil {
		return types2.SUCCESS
	} else {
		return types2.FAILURE
	}
}
func AddNextIndexStrEx(arg *types2.Zval, str string) int {
	var tmp types2.Zval
	tmp.SetStringVal(str)
	if arg.Array().NextIndexInsert(&tmp) != nil {
		return types2.SUCCESS
	} else {
		return types2.FAILURE
	}
}
func AddNextIndexStr(arg *types2.Zval, str *types2.String) int {
	var tmp types2.Zval
	tmp.SetString(str)
	if arg.Array().NextIndexInsert(&tmp) != nil {
		return types2.SUCCESS
	} else {
		return types2.FAILURE
	}
}
func AddNextIndexString(arg *types2.Zval, str string) int {
	var tmp types2.Zval
	tmp.SetStringVal(str)
	if arg.Array().NextIndexInsert(&tmp) != nil {
		return types2.SUCCESS
	} else {
		return types2.FAILURE
	}
}
func AddNextIndexStringl(arg *types2.Zval, str *byte, length int) int {
	var tmp types2.Zval
	tmp.SetStringVal(b.CastStr(str, length))
	if arg.Array().NextIndexInsert(&tmp) != nil {
		return types2.SUCCESS
	} else {
		return types2.FAILURE
	}
}
func ArraySetZvalKey(ht *types2.Array, key *types2.Zval, value *types2.Zval) int {
	var result *types2.Zval
	switch key.GetType() {
	case types2.IS_STRING:
		result = ht.SymtableUpdate(key.String().GetStr(), value)
		break
	case types2.IS_NULL:
		result = ht.SymtableUpdate(types2.NewString("").GetStr(), value)
		break
	case types2.IS_RESOURCE:
		faults.Error(faults.E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", types2.Z_RES_HANDLE_P(key), types2.Z_RES_HANDLE_P(key))
		result = ht.IndexUpdate(types2.Z_RES_HANDLE_P(key), value)
		break
	case types2.IS_FALSE:
		result = ht.IndexUpdate(0, value)
		break
	case types2.IS_TRUE:
		result = ht.IndexUpdate(1, value)
		break
	case types2.IS_LONG:
		result = ht.IndexUpdate(key.Long(), value)
		break
	case types2.IS_DOUBLE:
		result = ht.IndexUpdate(DvalToLval(key.Double()), value)
		break
	default:
		faults.Error(faults.E_WARNING, "Illegal offset type")
		result = nil
	}
	if result != nil {
		// result.TryAddRefcount()
		return types2.SUCCESS
	} else {
		return types2.FAILURE
	}
}
func AddPropertyLongEx(arg *types2.Zval, key string, n ZendLong) int {
	return AddPropertyZvalEx(arg, key, types2.NewZvalLong(n))
}
func AddPropertyNullEx(arg *types2.Zval, key string) int {
	return AddPropertyZvalEx(arg, key, types2.NewZvalNull())
}
func AddPropertyResourceEx(arg *types2.Zval, key string, r *types2.ZendResource) int {
	return AddPropertyZvalEx(arg, key, types2.NewZvalResource(r))
}
func AddPropertyStrEx(arg *types2.Zval, key string, str string) int {
	return AddPropertyZvalEx(arg, key, types2.NewZvalString(str))
}
func AddPropertyZvalEx(arg *types2.Zval, key string, value *types2.Zval) int {
	zKey := types2.NewZvalString(key)
	types2.Z_OBJ_HT(*arg).GetWriteProperty()(arg, zKey, value, nil)
	return types2.SUCCESS
}
func ZendStartupModuleEx(module *ModuleEntry) bool {
	if module.GetModuleStarted() != 0 {
		return true
	}
	module.SetModuleStarted(1)

	/* Initialize module globals */
	if module.GetGlobalsSize() != 0 {
		if module.GetGlobalsCtor() != nil {
			module.GetGlobalsCtor()(module.GetGlobalsPtr())
		}
	}
	if module.GetModuleStartupFunc() != nil {
		EG__().SetCurrentModule(module)
		if module.GetModuleStartupFunc()(module.GetType(), module.GetModuleNumber()) == types2.FAILURE {
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Unable to start %s module", module.GetName())
			EG__().SetCurrentModule(nil)
			return false
		}
		EG__().SetCurrentModule(nil)
	}
	return true
}
