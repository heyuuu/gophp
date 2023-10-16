package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ObjectPropertiesLoad(object *types.Object, properties *types.Array) {
	var prop *types.Zval
	var tmp types.Zval
	var key *types.String
	var h ZendLong
	var property_info *types.PropertyInfo
	var __ht *types.Array = properties
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		h = _p.GetH()
		key = _p.GetKey()
		prop = _z
		if key != nil {
			if key.GetStr()[0] == '0' {
				if className, propName, ok := ZendUnmanglePropertyName_Ex(key.GetStr()); ok {
					var prevScope *types.ClassEntry = EG__().GetFakeScope()
					if className != "" && className[0] != '*' {
						EG__().SetFakeScope(ZendLookupClass(className))
					}
					property_info = ZendGetPropertyInfo(object.GetCe(), propName)
					EG__().SetFakeScope(prevScope)
				} else {
					property_info = nil
				}
			} else {
				property_info = ZendGetPropertyInfo(object.GetCe(), key.GetStr())
			}
			if property_info != nil && !property_info.IsStatic() {
				var slot *types.Zval = OBJ_PROP(object, property_info.GetOffset())
				slot.CopyValueFrom(prop)
				if object.GetProperties() != nil {
					tmp.SetIndirect(slot)
					object.GetProperties().KeyUpdate(key.GetStr(), &tmp)
				}
			} else {
				if object.GetProperties() == nil {
					RebuildObjectProperties(object)
				}
				prop = object.GetProperties().KeyUpdate(key.GetStr(), prop)
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
func AddIndexDouble(arg *types.Zval, index ZendUlong, d float64) int {
	var tmp types.Zval
	tmp.SetDouble(d)
	arg.Array().IndexUpdate(index, &tmp)
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
func AddNextIndexResource(arg *types.Zval, r *types.Resource) int {
	var tmp types.Zval
	tmp.SetResource(r)
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexStrEx(arg *types.Zval, str string) int {
	var tmp types.Zval
	tmp.SetString(str)
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexStr(arg *types.Zval, str *types.String) int {
	var tmp types.Zval
	tmp.SetStringEx(str)
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexString(arg *types.Zval, str string) int {
	var tmp types.Zval
	tmp.SetString(str)
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func AddNextIndexStringl(arg *types.Zval, str *byte, length int) int {
	var tmp types.Zval
	tmp.SetString(b.CastStr(str, length))
	if arg.Array().Append(&tmp) != nil {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func ArraySetZvalKey(ht *types.Array, key *types.Zval, value *types.Zval) int {
	var result *types.Zval
	switch key.Type() {
	case types.IsString:
		result = ht.SymtableUpdate(key.String(), value)
		break
	case types.IsNull:
		result = ht.SymtableUpdate(types.NewString("").GetStr(), value)
		break
	case types.IsResource:
		faults.Error(faults.E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", key.ResourceHandle(), key.ResourceHandle())
		result = ht.IndexUpdate(key.ResourceHandle(), value)
		break
	case types.IsFalse:
		result = ht.IndexUpdate(0, value)
		break
	case types.IsTrue:
		result = ht.IndexUpdate(1, value)
		break
	case types.IsLong:
		result = ht.IndexUpdate(key.Long(), value)
		break
	case types.IsDouble:
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
func AddPropertyResourceEx(arg *types.Zval, key string, r *types.Resource) int {
	return AddPropertyZvalEx(arg, key, types.NewZvalResource(r))
}
func AddPropertyStrEx(arg *types.Zval, key string, str string) int {
	return AddPropertyZvalEx(arg, key, types.NewZvalString(str))
}
func AddPropertyZvalEx(arg *types.Zval, key string, value *types.Zval) int {
	zKey := types.NewZvalString(key)
	arg.Object().WritePropertyEx(zKey, value)
	return types.SUCCESS
}
func ZendStartupModuleEx(module *ModuleEntry) bool {
	if module.IsModuleStarted() {
		return true
	}
	module.SetModuleStarted(true)

	/* Initialize module globals */
	if module.GetGlobalsSize() != 0 {
		if module.GetGlobalsCtor() != nil {
			module.GetGlobalsCtor()(module.GetGlobalsPtr())
		}
	}

	EG__().SetCurrentModule(module)
	if !module.ModuleStartup() {
		faults.ErrorNoreturn(faults.E_CORE_ERROR, "Unable to start %s module", module.Name())
		EG__().SetCurrentModule(nil)
		return false
	}
	EG__().SetCurrentModule(nil)
	return true
}
