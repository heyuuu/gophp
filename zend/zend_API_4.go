package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ObjectPropertiesInit(object *types.ZendObject, class_type *types.ClassEntry) {
	object.SetProperties(nil)
	_objectPropertiesInit(object, class_type)
}
func ObjectPropertiesInitEx(object *types.ZendObject, properties *types.Array) {
	object.SetProperties(properties)
	if object.GetCe().GetDefaultPropertiesCount() != 0 {
		properties.Foreach(func(key_ types.ArrayKey, prop *types.Zval) {
			propertyInfo := ZendGetPropertyInfo(object.GetCe(), key_.StrKey(), 1)
			if propertyInfo != ZEND_WRONG_PROPERTY_INFO && propertyInfo != nil && !propertyInfo.IsStatic() {
				var slot *types.Zval = OBJ_PROP(object, propertyInfo.GetOffset())
				if propertyInfo.GetType() != 0 {
					var tmp types.Zval
					types.ZVAL_COPY_VALUE(&tmp, prop)
					if ZendVerifyPropertyType(propertyInfo, &tmp, 0) == 0 {
						return
					}
					types.ZVAL_COPY_VALUE(slot, &tmp)
				} else {
					slot.CopyValueFrom(prop)
				}
				prop.SetIndirect(slot)
			}
		})
	}
}
func ObjectPropertiesLoad(object *types.ZendObject, properties *types.Array) {
	var prop *types.Zval
	var tmp types.Zval
	var key *types.String
	var h ZendLong
	var property_info *ZendPropertyInfo
	var __ht *types.Array = properties
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		h = _p.GetH()
		key = _p.GetKey()
		prop = _z
		if key != nil {
			if key.GetStr()[0] == '0' {
				var class_name *byte
				var prop_name *byte
				var prop_name_len int
				if ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, &prop_name_len) == types.SUCCESS {
					var pname *types.String = types.NewString(b.CastStr(prop_name, prop_name_len))
					var prev_scope *types.ClassEntry = EG__().GetFakeScope()
					if class_name != nil && class_name[0] != '*' {
						var cname *types.String = types.NewString(class_name)
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
				var slot *types.Zval = OBJ_PROP(object, property_info.GetOffset())
				// ZvalPtrDtor(slot)
				slot.CopyValueFrom(prop)
				//ZvalAddRef(slot)
				if object.GetProperties() != nil {
					tmp.SetIndirect(slot)
					object.GetProperties().KeyUpdate(key.GetStr(), &tmp)
				}
			} else {
				if object.GetProperties() == nil {
					RebuildObjectProperties(object)
				}
				prop = object.GetProperties().KeyUpdate(key.GetStr(), prop)
				//ZvalAddRef(prop)
			}
		} else {
			if object.GetProperties() == nil {
				RebuildObjectProperties(object)
			}
			prop = object.GetProperties().IndexUpdate(h, prop)
			//ZvalAddRef(prop)
		}
	}
}
func AddAssocLongEx(arg *types.Zval, key string, n ZendLong) int {
	var tmp types.Zval
	tmp.SetLong(n)
	arg.Array().SymtableUpdate(key, &tmp)
	return types.SUCCESS
}
func AddAssocNullEx(arg *types.Zval, key string) int {
	var tmp types.Zval
	tmp.SetNull()
	arg.Array().SymtableUpdate(key, &tmp)
	return types.SUCCESS
}
func AddAssocBoolEx(arg *types.Zval, key string, b int) int {
	var tmp types.Zval
	tmp.SetBool(b != 0)
	arg.Array().SymtableUpdate(key, &tmp)
	return types.SUCCESS
}
func AddAssocDoubleEx(arg *types.Zval, key string, d float64) int {
	var tmp types.Zval
	tmp.SetDouble(d)
	arg.Array().SymtableUpdate(key, &tmp)
	return types.SUCCESS
}
func AddAssocStrEx(arg *types.Zval, key string, str string) int {
	arg.Array().SymtableUpdate(key, types.NewZvalString(str))
	return types.SUCCESS
}
func AddAssocStringlEx(arg *types.Zval, key string, str string) int {
	arg.Array().SymtableUpdate(key, types.NewZvalString(str))
	return types.SUCCESS
}
func AddAssocZvalEx(arg *types.Zval, key string, value *types.Zval) int {
	arg.Array().SymtableUpdate(key, value)
	return types.SUCCESS
}
func AddIndexLong(arg *types.Zval, index ZendUlong, n ZendLong) int {
	var tmp types.Zval
	tmp.SetLong(n)
	arg.Array().IndexUpdate(index, &tmp)
	return types.SUCCESS
}
func AddIndexBool(arg *types.Zval, index ZendUlong, b int) int {
	var tmp types.Zval
	tmp.SetBool(b != 0)
	arg.Array().IndexUpdate(index, &tmp)
	return types.SUCCESS
}
func AddIndexDouble(arg *types.Zval, index ZendUlong, d float64) int {
	var tmp types.Zval
	tmp.SetDouble(d)
	arg.Array().IndexUpdate(index, &tmp)
	return types.SUCCESS
}
func AddIndexStr(arg *types.Zval, index int, str *types.String) int {
	zv := types.NewZvalString(str.GetStr())
	arg.Array().IndexUpdate(index, zv)
	return types.SUCCESS
}
func AddIndexStrEx(arg *types.Zval, index int, str string) int {
	zv := types.NewZvalString(str)
	arg.Array().IndexUpdate(index, zv)
	return types.SUCCESS
}
func AddIndexString(arg *types.Zval, index ZendUlong, str *byte) int {
	zv := types.NewZvalString(b.CastStrAuto(str))
	arg.Array().IndexUpdate(index, zv)
	return types.SUCCESS
}
func AddIndexStringl(arg *types.Zval, index ZendUlong, str *byte, length int) int {
	zv := types.NewZvalString(b.CastStr(str, length))
	arg.Array().IndexUpdate(index, zv)
	return types.SUCCESS
}
func AddNextIndexLong(arg *types.Zval, n ZendLong) int {
	if arg.Array().Append(types.NewZvalLong(n)) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexNull(arg *types.Zval) int {
	if arg.Array().Append(types.NewZvalNull()) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexBool(arg *types.Zval, b int) int {
	var tmp types.Zval
	tmp.SetBool(b != 0)
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexResource(arg *types.Zval, r *types.ZendResource) int {
	var tmp types.Zval
	tmp.SetResource(r)
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexDouble(arg *types.Zval, d float64) int {
	var tmp types.Zval
	tmp.SetDouble(d)
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexStrEx(arg *types.Zval, str string) int {
	var tmp types.Zval
	tmp.SetStringVal(str)
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexStr(arg *types.Zval, str *types.String) int {
	var tmp types.Zval
	tmp.SetString(str)
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexString(arg *types.Zval, str string) int {
	var tmp types.Zval
	tmp.SetStringVal(str)
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexStringl(arg *types.Zval, str *byte, length int) int {
	var tmp types.Zval
	tmp.SetStringVal(b.CastStr(str, length))
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func ArraySetZvalKey(ht *types.Array, key *types.Zval, value *types.Zval) int {
	var result *types.Zval
	switch key.GetType() {
	case types.IS_STRING:
		result = ht.SymtableUpdate(key.String().GetStr(), value)
		break
	case types.IS_NULL:
		result = ht.SymtableUpdate(types.NewString("").GetStr(), value)
		break
	case types.IS_RESOURCE:
		faults.Error(faults.E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", key.ResourceHandle(), key.ResourceHandle())
		result = ht.IndexUpdate(key.ResourceHandle(), value)
		break
	case types.IS_FALSE:
		result = ht.IndexUpdate(0, value)
		break
	case types.IS_TRUE:
		result = ht.IndexUpdate(1, value)
		break
	case types.IS_LONG:
		result = ht.IndexUpdate(key.Long(), value)
		break
	case types.IS_DOUBLE:
		result = ht.IndexUpdate(operators.DvalToLval(key.Double()), value)
		break
	default:
		faults.Error(faults.E_WARNING, "Illegal offset type")
		result = nil
	}
	if result != nil {
		// result.TryAddRefcount()
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddPropertyLongEx(arg *types.Zval, key string, n ZendLong) int {
	return AddPropertyZvalEx(arg, key, types.NewZvalLong(n))
}
func AddPropertyNullEx(arg *types.Zval, key string) int {
	return AddPropertyZvalEx(arg, key, types.NewZvalNull())
}
func AddPropertyResourceEx(arg *types.Zval, key string, r *types.ZendResource) int {
	return AddPropertyZvalEx(arg, key, types.NewZvalResource(r))
}
func AddPropertyStrEx(arg *types.Zval, key string, str string) int {
	return AddPropertyZvalEx(arg, key, types.NewZvalString(str))
}
func AddPropertyZvalEx(arg *types.Zval, key string, value *types.Zval) int {
	zKey := types.NewZvalString(key)
	arg.Object().Handlers().GetWriteProperty()(arg, zKey, value, nil)
	return types.SUCCESS
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
		if module.GetModuleStartupFunc()(module.GetType(), module.GetModuleNumber()) == types.FAILURE {
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Unable to start %s module", module.GetName())
			EG__().SetCurrentModule(nil)
			return false
		}
		EG__().SetCurrentModule(nil)
	}
	return true
}
