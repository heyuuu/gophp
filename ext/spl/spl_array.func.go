package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
)

func SplArrayFromObj(obj *types.ZendObject) *SplArrayObject {
	return (*SplArrayObject)((*byte)(obj - zend_long((*byte)(&((*SplArrayObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLARRAY_P(zv *types.Zval) *SplArrayObject {
	return SplArrayFromObj(zv.Object())
}
func SplArrayGetHashTablePtr(intern *SplArrayObject) **types.Array {
	if intern.IsIsSelf() {
		if intern.GetStd().GetProperties() == nil {
			zend.RebuildObjectProperties(intern.GetStd())
		}
		return intern.GetStd().GetProperties()
	} else if intern.IsUseOther() {
		var other *SplArrayObject = Z_SPLARRAY_P(intern.GetArray())
		return SplArrayGetHashTablePtr(other)
	} else if intern.GetArray().IsType(types.IS_ARRAY) {
		return &(intern.GetArray().Array())
	} else {
		var obj *types.ZendObject = intern.GetArray().Object()
		if obj.GetProperties() == nil {
			zend.RebuildObjectProperties(obj)
		} else if obj.GetProperties().GetRefcount() > 1 {
			obj.SetProperties(types.ZendArrayDup(obj.GetProperties()))
		}
		return obj.GetProperties()
	}

	//??? TODO: Delay duplication for arrays; only duplicate for write operations
}
func SplArrayGetHashTable(intern *SplArrayObject) *types.Array {
	return (*SplArrayGetHashTablePtr)(intern)
}
func SplArrayReplaceHashTable(intern *SplArrayObject, ht *types.Array) {
	var ht_ptr **types.Array = SplArrayGetHashTablePtr(intern)
	ht_ptr.DestroyEx()
	*ht_ptr = ht
}
func SplArrayIsObject(intern *SplArrayObject) types.ZendBool {
	for intern.IsUseOther() {
		intern = Z_SPLARRAY_P(intern.GetArray())
	}
	return intern.IsIsSelf() || intern.GetArray().IsType(types.IS_OBJECT)
}
func SplArrayCreateHtIter(ht *types.Array, intern *SplArrayObject) {
	intern.SetHtIter(zend.EG__().AddArrayIterator(ht))
	SplArraySkipProtected(intern, ht)
}
func SplArrayGetPosPtr(ht *types.Array, intern *SplArrayObject) *uint32 {
	if intern.GetHtIter() == uint32-1 {
		SplArrayCreateHtIter(ht, intern)
	}
	return zend.EG__().GetHtIterators()[intern.GetHtIter()].GetPos()
}
func SplArrayObjectFreeStorage(object *types.ZendObject) {
	var intern *SplArrayObject = SplArrayFromObj(object)
	if intern.GetHtIter() != uint32-1 {
		zend.EG__().DelArrayIterator(intern.GetHtIter())
	}
	zend.ZendObjectStdDtor(intern.GetStd())
}
func SplArrayObjectNewEx(classType *types.ClassEntry, orig *types.Zval, cloneOrig int) *types.ZendObject {
	var intern *SplArrayObject = new(SplArrayObject)
	var parent *types.ClassEntry = classType
	var inherited int = 0
	zend.ZendObjectStdInit(intern.GetStd(), classType)
	zend.ObjectPropertiesInit(intern.GetStd(), classType)
	intern.SetArFlags(0)
	intern.SetCeGetIterator(spl_ce_ArrayIterator)
	if orig != nil {
		var other *SplArrayObject = Z_SPLARRAY_P(orig)
		intern.SetIsCloneMask(false)
		intern.AddArFlags(other.GetArFlags() & SPL_ARRAY_CLONE_MASK)
		intern.SetCeGetIterator(other.GetCeGetIterator())
		if cloneOrig != 0 {
			if other.IsIsSelf() {
				intern.GetArray().SetUndef()
			} else if orig.Object().Handlers() == &spl_handler_ArrayObject {
				intern.GetArray().SetArray(types.ZendArrayDup(SplArrayGetHashTable(other)))
			} else {
				b.Assert(orig.Object().Handlers() == &spl_handler_ArrayIterator)
				types.ZVAL_COPY(intern.GetArray(), orig)
				intern.SetIsUseOther(true)
			}
		} else {
			types.ZVAL_COPY(intern.GetArray(), orig)
			intern.SetIsUseOther(true)
		}
	} else {
		zend.ArrayInit(intern.GetArray())
	}
	for parent != nil {
		if parent == spl_ce_ArrayIterator || parent == spl_ce_RecursiveArrayIterator {
			intern.GetStd().SetHandlers(&spl_handler_ArrayIterator)
			break
		} else if parent == spl_ce_ArrayObject {
			intern.GetStd().SetHandlers(&spl_handler_ArrayObject)
			break
		}
		parent = parent.GetParent()
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, faults.E_COMPILE_ERROR, "Internal compiler error, Class is not child of ArrayObject or ArrayIterator")
	}
	if inherited != 0 {
		intern.SetFptrOffsetGet(classType.FunctionTable().Get("offsetget"))
		if intern.GetFptrOffsetGet().GetScope() == parent {
			intern.SetFptrOffsetGet(nil)
		}
		intern.SetFptrOffsetSet(classType.FunctionTable().Get("offsetset"))
		if intern.GetFptrOffsetSet().GetScope() == parent {
			intern.SetFptrOffsetSet(nil)
		}
		intern.SetFptrOffsetHas(classType.FunctionTable().Get("offsetexists"))
		if intern.GetFptrOffsetHas().GetScope() == parent {
			intern.SetFptrOffsetHas(nil)
		}
		intern.SetFptrOffsetDel(classType.FunctionTable().Get("offsetunset"))
		if intern.GetFptrOffsetDel().GetScope() == parent {
			intern.SetFptrOffsetDel(nil)
		}
		intern.SetFptrCount(classType.FunctionTable().Get("count"))
		if intern.GetFptrCount().GetScope() == parent {
			intern.SetFptrCount(nil)
		}
	}

	/* Cache iterator functions if ArrayIterator or derived. Check current's */

	if intern.GetStd().GetHandlers() == &spl_handler_ArrayIterator {
		var funcs_ptr *zend.ZendClassIteratorFuncs = classType.GetIteratorFuncsPtr()
		if funcs_ptr.GetZfCurrent() == nil {
			funcs_ptr.SetZfRewind(classType.FunctionTable().Get("rewind"))
			funcs_ptr.SetZfValid(classType.FunctionTable().Get("valid"))
			funcs_ptr.SetZfKey(classType.FunctionTable().Get("key"))
			funcs_ptr.SetZfCurrent(classType.FunctionTable().Get("current"))
			funcs_ptr.SetZfNext(classType.FunctionTable().Get("next"))
		}
		if inherited != 0 {
			if funcs_ptr.GetZfRewind().GetScope() != parent {
				intern.SetIsOverloadedRewind(true)
			}
			if funcs_ptr.GetZfValid().GetScope() != parent {
				intern.SetIsOverloadedValid(true)
			}
			if funcs_ptr.GetZfKey().GetScope() != parent {
				intern.SetIsOverloadedKey(true)
			}
			if funcs_ptr.GetZfCurrent().GetScope() != parent {
				intern.SetIsOverloadedCurrent(true)
			}
			if funcs_ptr.GetZfNext().GetScope() != parent {
				intern.SetIsOverloadedNext(true)
			}
		}
	}
	intern.SetHtIter(uint32 - 1)
	return intern.GetStd()
}
func SplArrayObjectNew(class_type *types.ClassEntry) *types.ZendObject {
	return SplArrayObjectNewEx(class_type, nil, 0)
}
func SplArrayObjectClone(zobject *types.Zval) *types.ZendObject {
	var old_object *types.ZendObject
	var new_object *types.ZendObject
	old_object = zobject.Object()
	new_object = SplArrayObjectNewEx(old_object.GetCe(), zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
func SplArrayGetDimensionPtr(check_inherited int, intern *SplArrayObject, offset *types.Zval, type_ int) *types.Zval {
	var retval *types.Zval
	var index zend.ZendLong
	var offset_key *types.String
	var ht *types.Array = SplArrayGetHashTable(intern)
	if offset == nil || offset.IsUndef() || ht == nil {
		return zend.EG__().GetUninitializedZval()
	}
	if (type_ == zend.BP_VAR_W || type_ == zend.BP_VAR_RW) && intern.GetNApplyCount() > 0 {
		faults.Error(faults.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return zend.EG__().GetErrorZval()
	}
try_again:
	switch offset.GetType() {
	case types.IS_NULL:
		offset_key = types.NewString("")
		goto fetch_dim_string
	case types.IS_STRING:
		offset_key = offset.String()
	fetch_dim_string:
		retval = ht.SymtableFind(offset_key.GetStr())
		if retval != nil {
			if retval.IsIndirect() {
				retval = retval.Indirect()
				if retval.IsUndef() {
					switch type_ {
					case zend.BP_VAR_R:
						faults.Error(faults.E_NOTICE, "Undefined index: %s", offset_key.GetVal())
						fallthrough
					case zend.BP_VAR_UNSET:
						fallthrough
					case zend.BP_VAR_IS:
						retval = zend.EG__().GetUninitializedZval()
					case zend.BP_VAR_RW:
						faults.Error(faults.E_NOTICE, "Undefined index: %s", offset_key.GetVal())
						fallthrough
					case zend.BP_VAR_W:
						retval.SetNull()
					}
				}
			}
		} else {
			switch type_ {
			case zend.BP_VAR_R:
				faults.Error(faults.E_NOTICE, "Undefined index: %s", offset_key.GetVal())
				fallthrough
			case zend.BP_VAR_UNSET:
				fallthrough
			case zend.BP_VAR_IS:
				retval = zend.EG__().GetUninitializedZval()
			case zend.BP_VAR_RW:
				faults.Error(faults.E_NOTICE, "Undefined index: %s", offset_key.GetVal())
				fallthrough
			case zend.BP_VAR_W:
				var value types.Zval
				value.SetNull()
				retval = ht.SymtableUpdate(offset_key.GetStr(), &value)
			}
		}
		return retval
	case types.IS_RESOURCE:
		faults.Error(faults.E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", offset.ResourceHandle(), offset.ResourceHandle())
		index = offset.ResourceHandle()
		goto num_index
	case types.IS_DOUBLE:
		index = zend.ZendLong(offset.Double())
		goto num_index
	case types.IS_FALSE:
		index = 0
		goto num_index
	case types.IS_TRUE:
		index = 1
		goto num_index
	case types.IS_LONG:
		index = offset.Long()
	num_index:
		if b.Assign(&retval, ht.IndexFind(index)) == nil {
			switch type_ {
			case zend.BP_VAR_R:
				faults.Error(faults.E_NOTICE, "Undefined offset: "+zend.ZEND_LONG_FMT, index)
				fallthrough
			case zend.BP_VAR_UNSET:
				fallthrough
			case zend.BP_VAR_IS:
				retval = zend.EG__().GetUninitializedZval()
			case zend.BP_VAR_RW:
				faults.Error(faults.E_NOTICE, "Undefined offset: "+zend.ZEND_LONG_FMT, index)
				fallthrough
			case zend.BP_VAR_W:
				var value types.Zval
				value.SetUndef()
				retval = ht.IndexUpdate(index, &value)
			}
		}
		return retval
	case types.IS_REFERENCE:
		offset = types.ZVAL_DEREF(offset)
		goto try_again
	default:
		faults.Error(faults.E_WARNING, "Illegal offset type")
		if type_ == zend.BP_VAR_W || type_ == zend.BP_VAR_RW {
			return zend.EG__().GetErrorZval()
		} else {
			return zend.EG__().GetUninitializedZval()
		}
	}
}
func SplArrayReadDimensionEx(check_inherited int, object *types.Zval, offset *types.Zval, type_ int, rv *types.Zval) *types.Zval {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var ret *types.Zval
	if check_inherited != 0 && (intern.GetFptrOffsetGet() != nil || type_ == zend.BP_VAR_IS && intern.GetFptrOffsetHas() != nil) {
		if type_ == zend.BP_VAR_IS {
			if SplArrayHasDimension(object, offset, 0) == 0 {
				return zend.EG__().GetUninitializedZval()
			}
		}
		if intern.GetFptrOffsetGet() != nil {
			var tmp types.Zval
			if offset == nil {
				tmp.SetUndef()
				offset = &tmp
			} else {
				offset = types.SEPARATE_ARG_IF_REF(offset)
			}
			zend.ZendCallMethodWith1Params(object, types.Z_OBJCE_P(object), intern.GetFptrOffsetGet(), "offsetGet", rv, offset)
			// zend.ZvalPtrDtor(offset)
			if !(rv.IsUndef()) {
				return rv
			}
			return zend.EG__().GetUninitializedZval()
		}
	}
	ret = SplArrayGetDimensionPtr(check_inherited, intern, offset, type_)

	/* When in a write context,
	 * ZE has to be fooled into thinking this is in a reference set
	 * by separating (if necessary) and returning as IS_REFERENCE (with refcount == 1)
	 */

	if (type_ == zend.BP_VAR_W || type_ == zend.BP_VAR_RW || type_ == zend.BP_VAR_UNSET) && !(ret.IsReference()) && ret != zend.EG__().GetUninitializedZval() {
		ret.SetNewRef(ret)
	}
	return ret
}
func SplArrayReadDimension(object *types.Zval, offset *types.Zval, type_ int, rv *types.Zval) *types.Zval {
	return SplArrayReadDimensionEx(1, object, offset, type_, rv)
}
func SplArrayWriteDimensionEx(check_inherited int, object *types.Zval, offset *types.Zval, value *types.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var index zend.ZendLong
	var ht *types.Array
	if check_inherited != 0 && intern.GetFptrOffsetSet() != nil {
		var tmp types.Zval
		if offset == nil {
			tmp.SetNull()
			offset = &tmp
		} else {
			offset = types.SEPARATE_ARG_IF_REF(offset)
		}
		zend.ZendCallMethodWith2Params(object, types.Z_OBJCE_P(object), intern.GetFptrOffsetSet(), "offsetSet", nil, offset, value)
		// zend.ZvalPtrDtor(offset)
		return
	}
	if intern.GetNApplyCount() > 0 {
		faults.Error(faults.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return
	}
	//value.TryAddRefcount()
	if offset == nil {
		ht = SplArrayGetHashTable(intern)
		ht.Append(value)
		return
	}
try_again:
	switch offset.GetType() {
	case types.IS_STRING:
		ht = SplArrayGetHashTable(intern)
		ht.SymtableUpdateInd(offset.String().GetStr(), value)
		return
	case types.IS_DOUBLE:
		index = zend.ZendLong(offset.Double())
		goto num_index
	case types.IS_RESOURCE:
		index = offset.ResourceHandle()
		goto num_index
	case types.IS_FALSE:
		index = 0
		goto num_index
	case types.IS_TRUE:
		index = 1
		goto num_index
	case types.IS_LONG:
		index = offset.Long()
	num_index:
		ht = SplArrayGetHashTable(intern)
		ht.IndexUpdate(index, value)
		return
	case types.IS_NULL:
		ht = SplArrayGetHashTable(intern)
		ht.Append(value)
		return
	case types.IS_REFERENCE:
		offset = types.ZVAL_DEREF(offset)
		goto try_again
	default:
		faults.Error(faults.E_WARNING, "Illegal offset type")
		// zend.ZvalPtrDtor(value)
		return
	}
}
func SplArrayWriteDimension(object *types.Zval, offset *types.Zval, value *types.Zval) {
	SplArrayWriteDimensionEx(1, object, offset, value)
}
func SplArrayUnsetDimensionEx(check_inherited int, object *types.Zval, offset *types.Zval) {
	var index zend.ZendLong
	var ht *types.Array
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if check_inherited != 0 && intern.GetFptrOffsetDel() != nil {
		offset = types.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, types.Z_OBJCE_P(object), intern.GetFptrOffsetDel(), "offsetUnset", nil, offset)
		// zend.ZvalPtrDtor(offset)
		return
	}
	if intern.GetNApplyCount() > 0 {
		faults.Error(faults.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return
	}
try_again:
	switch offset.GetType() {
	case types.IS_STRING:
		ht = SplArrayGetHashTable(intern)
		if ht == zend.EG__().GetSymbolTable() {
			if zend.ZendDeleteGlobalVariable(offset.String()) != 0 {
				faults.Error(faults.E_NOTICE, "Undefined index: %s", offset.String().GetVal())
			}
		} else {
			var data *types.Zval = ht.SymtableFind(offset.String().GetStr())
			if data != nil {
				if data.IsIndirect() {
					data = data.Indirect()
					if data.IsUndef() {
						faults.Error(faults.E_NOTICE, "Undefined index: %s", offset.String().GetVal())
					} else {
						// zend.ZvalPtrDtor(data)
						data.SetUndef()
						ht.MarkHasEmptyIndex()
						types.ZendHashMoveForwardEx(ht, SplArrayGetPosPtr(ht, intern))
						if SplArrayIsObject(intern) != 0 {
							SplArraySkipProtected(intern, ht)
						}
					}
				} else if ht.SymtableDel(offset.String().GetStr()) == false {
					faults.Error(faults.E_NOTICE, "Undefined index: %s", offset.String().GetVal())
				}
			} else {
				faults.Error(faults.E_NOTICE, "Undefined index: %s", offset.String().GetVal())
			}
		}
	case types.IS_DOUBLE:
		index = zend.ZendLong(offset.Double())
		goto num_index
	case types.IS_RESOURCE:
		index = offset.ResourceHandle()
		goto num_index
	case types.IS_FALSE:
		index = 0
		goto num_index
	case types.IS_TRUE:
		index = 1
		goto num_index
	case types.IS_LONG:
		index = offset.Long()
	num_index:
		ht = SplArrayGetHashTable(intern)
		if types.ZendHashIndexDel(ht, index) == types.FAILURE {
			faults.Error(faults.E_NOTICE, "Undefined offset: "+zend.ZEND_LONG_FMT, index)
		}
	case types.IS_REFERENCE:
		offset = types.ZVAL_DEREF(offset)
		goto try_again
	default:
		faults.Error(faults.E_WARNING, "Illegal offset type")
		return
	}
}
func SplArrayUnsetDimension(object *types.Zval, offset *types.Zval) {
	SplArrayUnsetDimensionEx(1, object, offset)
}
func SplArrayHasDimensionEx(check_inherited int, object *types.Zval, offset *types.Zval, check_empty int) int {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var index zend.ZendLong
	var rv types.Zval
	var value *types.Zval = nil
	var tmp *types.Zval
	if check_inherited != 0 && intern.GetFptrOffsetHas() != nil {
		offset = types.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, types.Z_OBJCE_P(object), intern.GetFptrOffsetHas(), "offsetExists", &rv, offset)
		// zend.ZvalPtrDtor(offset)
		if operators.ZvalIsTrue(&rv) {
			// zend.ZvalPtrDtor(&rv)
			if check_empty != 1 {
				return 1
			} else if intern.GetFptrOffsetGet() != nil {
				value = SplArrayReadDimensionEx(1, object, offset, zend.BP_VAR_R, &rv)
			}
		} else {
			// zend.ZvalPtrDtor(&rv)
			return 0
		}
	}
	if value == nil {
		var ht *types.Array = SplArrayGetHashTable(intern)
	try_again:
		switch offset.GetType() {
		case types.IS_STRING:
			if b.Assign(&tmp, ht.SymtableFind(offset.String().GetStr())) != nil {
				if check_empty == 2 {
					return 1
				}
			} else {
				return 0
			}
		case types.IS_DOUBLE:
			index = zend.ZendLong(offset.Double())
			goto num_index
		case types.IS_RESOURCE:
			index = offset.ResourceHandle()
			goto num_index
		case types.IS_FALSE:
			index = 0
			goto num_index
		case types.IS_TRUE:
			index = 1
			goto num_index
		case types.IS_LONG:
			index = offset.Long()
		num_index:
			if b.Assign(&tmp, ht.IndexFind(index)) != nil {
				if check_empty == 2 {
					return 1
				}
			} else {
				return 0
			}
		case types.IS_REFERENCE:
			offset = types.ZVAL_DEREF(offset)
			goto try_again
		default:
			faults.Error(faults.E_WARNING, "Illegal offset type")
			return 0
		}
		if check_empty != 0 && check_inherited != 0 && intern.GetFptrOffsetGet() != nil {
			value = SplArrayReadDimensionEx(1, object, offset, zend.BP_VAR_R, &rv)
		} else {
			value = tmp
		}
	}
	var result types.ZendBool = b.CondF(check_empty != 0, func() int { return operators.IZendIsTrue(value) }, func() bool { return value.GetType() != types.IS_NULL })
	if value == &rv {
		// zend.ZvalPtrDtor(&rv)
	}
	return result
}
func SplArrayHasDimension(object *types.Zval, offset *types.Zval, check_empty int) int {
	return SplArrayHasDimensionEx(1, object, offset, check_empty)
}
func zim_spl_Array_offsetExists(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var index *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &index) == types.FAILURE {
		return
	}
	return_value.SetBool(SplArrayHasDimensionEx(0, zend.ZEND_THIS(executeData), index, 2) != 0)
	return
}
func zim_spl_Array_offsetGet(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value *types.Zval
	var index *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &index) == types.FAILURE {
		return
	}
	value = SplArrayReadDimensionEx(0, zend.ZEND_THIS(executeData), index, zend.BP_VAR_R, return_value)
	if value != return_value {
		types.ZVAL_COPY_DEREF(return_value, value)
	}
}
func zim_spl_Array_offsetSet(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var index *types.Zval
	var value *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "zz", &index, &value) == types.FAILURE {
		return
	}
	SplArrayWriteDimensionEx(0, zend.ZEND_THIS(executeData), index, value)
}
func SplArrayIteratorAppend(object *types.Zval, append_value *types.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if SplArrayIsObject(intern) != 0 {
		faults.ThrowError(nil, "Cannot append properties to objects, use %s::offsetSet() instead", types.Z_OBJCE_P(object).GetName().GetVal())
		return
	}
	SplArrayWriteDimension(object, nil, append_value)
}
func zim_spl_Array_append(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var value *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &value) == types.FAILURE {
		return
	}
	SplArrayIteratorAppend(zend.ZEND_THIS(executeData), value)
}
func zim_spl_Array_offsetUnset(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var index *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &index) == types.FAILURE {
		return
	}
	SplArrayUnsetDimensionEx(0, zend.ZEND_THIS(executeData), index)
}
func zim_spl_Array_getArrayCopy(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	return_value.SetArray(types.ZendArrayDup(SplArrayGetHashTable(intern)))
	return
}
func SplArrayGetPropertiesFor(object *types.Zval, purpose zend.ZendPropPurpose) *types.Array {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var ht *types.Array
	var dup types.ZendBool
	if intern.IsStdPropList() {
		return zend.ZendStdGetPropertiesFor(object, purpose)
	}

	/* We are supposed to be the only owner of the internal hashtable.
	 * The "dup" flag decides whether this is a "long-term" use where
	 * we need to duplicate, or a "temporary" one, where we can expect
	 * that no operations on the ArrayObject will be performed in the
	 * meantime. */

	switch purpose {
	case zend.ZEND_PROP_PURPOSE_ARRAY_CAST:
		dup = 1
	case zend.ZEND_PROP_PURPOSE_VAR_EXPORT:
		fallthrough
	case zend.ZEND_PROP_PURPOSE_JSON:
		fallthrough
	case zend._ZEND_PROP_PURPOSE_ARRAY_KEY_EXISTS:
		dup = 0
	default:
		return zend.ZendStdGetPropertiesFor(object, purpose)
	}
	ht = SplArrayGetHashTable(intern)
	if dup != 0 {
		ht = types.ZendArrayDup(ht)
	} else {
		// 		ht.AddRefcount()
	}
	return ht
}
func SplArrayGetDebugInfo(obj *types.Zval) *types.Array {
	var storage *types.Zval
	var zname *types.String
	var base *types.ClassEntry
	var intern *SplArrayObject = Z_SPLARRAY_P(obj)
	if intern.GetStd().GetProperties() == nil {
		zend.RebuildObjectProperties(intern.GetStd())
	}
	if intern.IsIsSelf() {
		return types.ZendArrayDup(intern.GetStd().GetProperties())
	} else {
		var debug_info *types.Array
		debug_info = types.NewArray(intern.GetStd().GetProperties().Len() + 1)
		types.ZendHashCopy(debug_info, intern.GetStd().GetProperties())
		storage = intern.GetArray()
		//storage.TryAddRefcount()
		if obj.Object().Handlers() == &spl_handler_ArrayIterator {
			base = spl_ce_ArrayIterator
		} else {
			base = spl_ce_ArrayObject
		}
		zname = SplGenPrivatePropName(base, "storage")
		debug_info.SymtableUpdate(zname.GetStr(), storage)
		// types.ZendStringReleaseEx(zname, 0)
		return debug_info
	}
}
func SplArrayGetGc(obj *types.Zval, gc_data **types.Zval, gc_data_count *int) *types.Array {
	var intern *SplArrayObject = Z_SPLARRAY_P(obj)
	*gc_data = intern.GetArray()
	*gc_data_count = 1
	return zend.ZendStdGetProperties(obj)
}
func SplArrayReadProperty(object *types.Zval, member *types.Zval, type_ int, cache_slot *any, rv *types.Zval) *types.Zval {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.IsArrayAsProps() && zend.ZendStdHasProperty(object, member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		return SplArrayReadDimension(object, member, type_, rv)
	}
	return zend.ZendStdReadProperty(object, member, type_, cache_slot, rv)
}
func SplArrayWriteProperty(object *types.Zval, member *types.Zval, value *types.Zval, cache_slot *any) *types.Zval {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.IsArrayAsProps() && zend.ZendStdHasProperty(object, member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		SplArrayWriteDimension(object, member, value)
		return value
	}
	return zend.ZendStdWriteProperty(object, member, value, cache_slot)
}
func SplArrayGetPropertyPtrPtr(object *types.Zval, member *types.Zval, type_ int, cache_slot *any) *types.Zval {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.IsArrayAsProps() && zend.ZendStdHasProperty(object, member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {

		/* If object has offsetGet() overridden, then fallback to read_property,
		 * which will call offsetGet(). */

		if intern.GetFptrOffsetGet() != nil {
			return nil
		}
		return SplArrayGetDimensionPtr(1, intern, member, type_)
	}
	return zend.ZendStdGetPropertyPtrPtr(object, member, type_, cache_slot)
}
func SplArrayHasProperty(object *types.Zval, member *types.Zval, has_set_exists int, cache_slot *any) int {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.IsArrayAsProps() && zend.ZendStdHasProperty(object, member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		return SplArrayHasDimension(object, member, has_set_exists)
	}
	return zend.ZendStdHasProperty(object, member, has_set_exists, cache_slot)
}
func SplArrayUnsetProperty(object *types.Zval, member *types.Zval, cache_slot *any) {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.IsArrayAsProps() && zend.ZendStdHasProperty(object, member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		SplArrayUnsetDimension(object, member)
		return
	}
	zend.ZendStdUnsetProperty(object, member, cache_slot)
}
func SplArrayCompareObjects(o1 *types.Zval, o2 *types.Zval) int {
	var ht1 *types.Array
	var ht2 *types.Array
	var intern1 *SplArrayObject
	var intern2 *SplArrayObject
	var result int = 0
	intern1 = Z_SPLARRAY_P(o1)
	intern2 = Z_SPLARRAY_P(o2)
	ht1 = SplArrayGetHashTable(intern1)
	ht2 = SplArrayGetHashTable(intern2)
	result = operators.ZendCompareSymbolTables(ht1, ht2)

	/* if we just compared std.properties, don't do it again */

	if result == 0 && !(ht1 == intern1.GetStd().GetProperties() && ht2 == intern2.GetStd().GetProperties()) {
		result = zend.ZendStdCompareObjects(o1, o2)
	}
	return result
}
func SplArraySkipProtected(intern *SplArrayObject, aht *types.Array) int {
	var data *types.Zval
	if SplArrayIsObject(intern) != 0 {
		var pos_ptr *uint32 = SplArrayGetPosPtr(aht, intern)
		for {
			if key := types.ZendHashGetCurrentKeyExEx(aht, *pos_ptr); key != nil && key.IsStrKey() {
				data = types.ZendHashGetCurrentDataEx(aht, pos_ptr)
				if data != nil && data.IsIndirect() && b.Assign(&data, data.Indirect()).GetType() == types.IS_UNDEF {

				} else if key.StrKey() == "" {
					return types.SUCCESS
				}
			} else {
				return types.SUCCESS
			}
			if !types.ZendHashHasMoreElementsEx(aht, pos_ptr) {
				return types.FAILURE
			}
			types.ZendHashMoveForwardEx(aht, pos_ptr)

		}
	}
	return types.FAILURE
}
func SplArrayNextEx(intern *SplArrayObject, aht *types.Array) int {
	var pos_ptr *uint32 = SplArrayGetPosPtr(aht, intern)
	types.ZendHashMoveForwardEx(aht, pos_ptr)
	if SplArrayIsObject(intern) != 0 {
		return SplArraySkipProtected(intern, aht)
	} else {
		return types.ResultCode(types.ZendHashHasMoreElementsEx(aht, pos_ptr))
	}
}
func SplArrayNext(intern *SplArrayObject) int {
	var aht *types.Array = SplArrayGetHashTable(intern)
	return SplArrayNextEx(intern, aht)
}
func SplArrayItDtor(iter *zend.ZendObjectIterator) {
	zend.ZendUserItInvalidateCurrent(iter)
}
func SplArrayItValid(iter *zend.ZendObjectIterator) int {
	var object *SplArrayObject = Z_SPLARRAY_P(iter.GetData())
	var aht *types.Array = SplArrayGetHashTable(object)
	if object.IsOverloadedValid() {
		return zend.ZendUserItValid(iter)
	} else {
		return types.ResultCode(types.ZendHashHasMoreElementsEx(aht, SplArrayGetPosPtr(aht, object)))
	}
}
func SplArrayItGetCurrentData(iter *zend.ZendObjectIterator) *types.Zval {
	var object *SplArrayObject = Z_SPLARRAY_P(iter.GetData())
	var aht *types.Array = SplArrayGetHashTable(object)
	if object.IsOverloadedCurrent() {
		return zend.ZendUserItGetCurrentData(iter)
	} else {
		var data *types.Zval = types.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, object))
		if data != nil && data.IsIndirect() {
			data = data.Indirect()
		}
		return data
	}
}
func SplArrayItGetCurrentKey(iter *zend.ZendObjectIterator, key *types.Zval) {
	var object *SplArrayObject = Z_SPLARRAY_P(iter.GetData())
	var aht *types.Array = SplArrayGetHashTable(object)
	if object.IsOverloadedKey() {
		zend.ZendUserItGetCurrentKey(iter, key)
	} else {
		types.ZendHashGetCurrentKeyZvalEx(aht, key, SplArrayGetPosPtr(aht, object))
	}
}
func SplArrayItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplArrayObject = Z_SPLARRAY_P(iter.GetData())
	var aht *types.Array = SplArrayGetHashTable(object)
	if object.IsOverloadedNext() {
		zend.ZendUserItMoveForward(iter)
	} else {
		zend.ZendUserItInvalidateCurrent(iter)
		SplArrayNextEx(object, aht)
	}
}
func SplArrayRewind(intern *SplArrayObject) {
	var aht *types.Array = SplArrayGetHashTable(intern)
	if intern.GetHtIter() == uint32-1 {
		SplArrayGetPosPtr(aht, intern)
	} else {
		types.ZendHashInternalPointerResetEx(aht, SplArrayGetPosPtr(aht, intern))
		SplArraySkipProtected(intern, aht)
	}
}
func SplArrayItRewind(iter *zend.ZendObjectIterator) {
	var object *SplArrayObject = Z_SPLARRAY_P(iter.GetData())
	if object.IsOverloadedRewind() {
		zend.ZendUserItRewind(iter)
	} else {
		zend.ZendUserItInvalidateCurrent(iter)
		SplArrayRewind(object)
	}
}
func SplArraySetArray(object *types.Zval, intern *SplArrayObject, array *types.Zval, ar_flags zend.ZendLong, just_array int) {
	if array.GetType() != types.IS_OBJECT && array.GetType() != types.IS_ARRAY {
		faults.ThrowException(spl_ce_InvalidArgumentException, "Passed variable is not an array or object", 0)
		return
	}
	if array.IsType(types.IS_ARRAY) {
		if array.GetRefcount() == 1 {
			types.ZVAL_COPY(intern.GetArray(), array)
		} else {
			//??? TODO: try to avoid array duplication
			intern.GetArray().SetArray(types.ZendArrayDup(array.Array()))
		}
	} else {
		if array.Object().Handlers() == &spl_handler_ArrayObject || array.Object().Handlers() == &spl_handler_ArrayIterator {
			// zend.ZvalPtrDtor(intern.GetArray())
			if just_array != 0 {
				var other *SplArrayObject = Z_SPLARRAY_P(array)
				ar_flags = other.GetArFlags() & ^SPL_ARRAY_INT_MASK
			}
			if object.Object() == array.Object() {
				ar_flags |= SPL_ARRAY_IS_SELF
				intern.GetArray().SetUndef()
			} else {
				ar_flags |= SPL_ARRAY_USE_OTHER
				types.ZVAL_COPY(intern.GetArray(), array)
			}
		} else {
			var handler zend.ZendObjectGetPropertiesT = array.Object().Handlers().GetGetProperties()
			if handler != zend.ZendStdGetProperties {
				faults.ThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "Overloaded object of type %s is not compatible with %s", types.Z_OBJCE_P(array).GetName().GetVal(), intern.GetStd().GetCe().GetName().GetVal())
				return
			}
			// zend.ZvalPtrDtor(intern.GetArray())
			types.ZVAL_COPY(intern.GetArray(), array)
		}
	}
	intern.SetArFlags(intern.GetArFlags() &^ SPL_ARRAY_IS_SELF & ^SPL_ARRAY_USE_OTHER)
	intern.AddArFlags(ar_flags)
	intern.SetHtIter(uint32 - 1)
}
func SplArrayGetIterator(ce *types.ClassEntry, object *types.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *zend.ZendUserIterator
	var array_object *SplArrayObject = Z_SPLARRAY_P(object)
	if by_ref != 0 && array_object.IsOverloadedCurrent() {
		faults.ThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("zend_user_iterator"))
	zend.ZendIteratorInit(iterator.GetIt())
	// 	object.AddRefcount()
	iterator.GetIt().GetData().SetObject(object.Object())
	iterator.GetIt().SetFuncs(&SplArrayItFuncs)
	iterator.SetCe(ce)
	iterator.GetValue().SetUndef()
	return iterator.GetIt()
}
func zim_spl_Array___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject
	var array *types.Zval
	var ar_flags zend.ZendLong = 0
	var ce_get_iterator *types.ClassEntry = spl_ce_ArrayIterator
	if executeData.NumArgs() == 0 {
		return
	}
	if zend.ZendParseParametersThrow(executeData.NumArgs(), "z|lC", &array, &ar_flags, &ce_get_iterator) == types.FAILURE {
		return
	}
	intern = Z_SPLARRAY_P(object)
	if executeData.NumArgs() > 2 {
		intern.SetCeGetIterator(ce_get_iterator)
	}
	ar_flags &= ^SPL_ARRAY_INT_MASK
	SplArraySetArray(object, intern, array, ar_flags, executeData.NumArgs() == 1)
}
func zim_spl_ArrayIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject
	var array *types.Zval
	var ar_flags zend.ZendLong = 0
	if executeData.NumArgs() == 0 {
		return
	}
	if zend.ZendParseParametersThrow(executeData.NumArgs(), "z|l", &array, &ar_flags) == types.FAILURE {
		return
	}
	intern = Z_SPLARRAY_P(object)
	ar_flags &= ^SPL_ARRAY_INT_MASK
	SplArraySetArray(object, intern, array, ar_flags, executeData.NumArgs() == 1)
}
func zim_spl_Array_setIteratorClass(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)

	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	ce_get_iterator := fp.ParseClass(spl_ce_ArrayIterator)
	if fp.HasError() {
		return
	}
	intern.SetCeGetIterator(ce_get_iterator)
}
func zim_spl_Array_getIteratorClass(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	//intern.GetCeGetIterator().GetName().AddRefcount()
	return_value.SetString(intern.GetCeGetIterator().GetName())
	return
}
func zim_spl_Array_getFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetLong(intern.GetArFlags() & ^SPL_ARRAY_INT_MASK)
	return
}
func zim_spl_Array_setFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var ar_flags zend.ZendLong = 0
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &ar_flags) == types.FAILURE {
		return
	}
	intern.SetArFlags(intern.GetArFlags()&SPL_ARRAY_INT_MASK | ar_flags & ^SPL_ARRAY_INT_MASK)
}
func zim_spl_Array_exchangeArray(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var array *types.Zval
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if zend.ZendParseParameters(executeData.NumArgs(), "z", &array) == types.FAILURE {
		return
	}
	if intern.GetNApplyCount() > 0 {
		faults.Error(faults.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return
	}
	return_value.SetArray(types.ZendArrayDup(SplArrayGetHashTable(intern)))
	SplArraySetArray(object, intern, array, 0, 1)
}
func zim_spl_Array_getIterator(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetObject(SplArrayObjectNewEx(intern.GetCeGetIterator(), object, 0))
}
func zim_spl_Array_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	SplArrayRewind(intern)
}
func zim_spl_Array_seek(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var opos zend.ZendLong
	var position zend.ZendLong
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *types.Array = SplArrayGetHashTable(intern)
	var result int
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &position) == types.FAILURE {
		return
	}
	opos = position
	if position >= 0 {
		SplArrayRewind(intern)
		result = types.SUCCESS
		for b.PostDec(&position) > 0 && b.Assign(&result, SplArrayNext(intern)) == types.SUCCESS {

		}
		if result == types.SUCCESS && types.ZendHashHasMoreElementsEx(aht, SplArrayGetPosPtr(aht, intern)) {
			return
		}
	}
	faults.ThrowExceptionEx(spl_ce_OutOfBoundsException, 0, "Seek position "+zend.ZEND_LONG_FMT+" is out of range", opos)
}
func SplArrayObjectCountElementsHelper(intern *SplArrayObject) zend.ZendLong {
	var aht *types.Array = SplArrayGetHashTable(intern)
	if SplArrayIsObject(intern) != 0 {
		var count zend.ZendLong = 0
		/* Count public/dynamic properties */
		aht.Foreach(func(key types.ArrayKey, value *types.Zval) {
			if value.IsIndirect() {
				if types.Z_INDIRECT_P(value).IsUndef() {
					return
				}
				if key.IsStrKey() && key.StrKey() == "" {
					return
				}
			}
			count++
		})

		return count
	} else {
		return aht.Len()
	}
}
func SplArrayObjectCountElements(object *types.Zval, count *zend.ZendLong) int {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.GetFptrCount() != nil {
		var rv types.Zval
		zend.ZendCallMethodWith0Params(object, intern.GetStd().GetCe(), intern.GetFptrCount(), "count", &rv)
		if rv.IsNotUndef() {
			*count = operators.ZvalGetLong(&rv)
			// zend.ZvalPtrDtor(&rv)
			return types.SUCCESS
		}
		*count = 0
		return types.FAILURE
	}
	*count = SplArrayObjectCountElementsHelper(intern)
	return types.SUCCESS
}
func zim_spl_Array_count(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(zend.ZEND_THIS(executeData))
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetLong(SplArrayObjectCountElementsHelper(intern))
	return
}
func SplArrayMethod(executeData *zend.ZendExecuteData, return_value *types.Zval, fname string, fname_len int, use_arg int) {
	var intern *SplArrayObject = Z_SPLARRAY_P(zend.ZEND_THIS(executeData))
	var aht *types.Array = SplArrayGetHashTable(intern)
	var function_name types.Zval
	var params []types.Zval
	var arg *types.Zval = nil
	function_name.SetStringVal(b.CastStr(fname, fname_len))
	params[0].SetNewEmptyRef()
	types.Z_REFVAL(params[0]).SetArray(aht)
	// 	aht.AddRefcount()
	if use_arg == 0 {
		intern.GetNApplyCount()++
		zend.CallUserFunction(nil, &function_name, return_value, 1, params)
		intern.GetNApplyCount()--
	} else if use_arg == SPL_ARRAY_METHOD_MAY_USER_ARG {
		if zend.ZendParseParametersEx(zpp.FlagQuiet, executeData.NumArgs(), "|z", &arg) == types.FAILURE {
			faults.ThrowException(spl_ce_BadMethodCallException, "Function expects one argument at most", 0)
			goto exit
		}
		if arg != nil {
			types.ZVAL_COPY_VALUE(&params[1], arg)
		}
		intern.GetNApplyCount()++
		zend.CallUserFunction(nil, &function_name, return_value, b.Cond(arg != nil, 2, 1), params)
		intern.GetNApplyCount()--
	} else {
		if executeData.NumArgs() != 1 || zend.ZendParseParametersEx(zpp.FlagQuiet, executeData.NumArgs(), "z", &arg) == types.FAILURE {
			faults.ThrowException(spl_ce_BadMethodCallException, "Function expects exactly one argument", 0)
			goto exit
		}
		types.ZVAL_COPY_VALUE(&params[1], arg)
		intern.GetNApplyCount()++
		zend.CallUserFunction(nil, &function_name, return_value, 2, params)
		intern.GetNApplyCount()--
	}
exit:
	var new_ht *types.Array = types.Z_REFVAL(params[0]).Array()
	if aht != new_ht {
		SplArrayReplaceHashTable(intern, new_ht)
	} else {
		//aht.DelRefcount()
	}
	types.Z_REFVAL(params[0]).SetNull()
	// zend.ZvalPtrDtor(&params[0])
	//types.ZendStringFree(function_name.String())
}
func zim_spl_Array_asort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplArrayMethod(executeData, return_value, "asort", b.SizeOf("\"asort\"")-1, SPL_ARRAY_METHOD_MAY_USER_ARG)
}
func zim_spl_Array_ksort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplArrayMethod(executeData, return_value, "ksort", b.SizeOf("\"ksort\"")-1, SPL_ARRAY_METHOD_MAY_USER_ARG)
}
func zim_spl_Array_uasort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplArrayMethod(executeData, return_value, "uasort", b.SizeOf("\"uasort\"")-1, SPL_ARRAY_METHOD_USE_ARG)
}
func zim_spl_Array_uksort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplArrayMethod(executeData, return_value, "uksort", b.SizeOf("\"uksort\"")-1, SPL_ARRAY_METHOD_USE_ARG)
}
func zim_spl_Array_natsort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplArrayMethod(executeData, return_value, "natsort", b.SizeOf("\"natsort\"")-1, SPL_ARRAY_METHOD_NO_ARG)
}
func zim_spl_Array_natcasesort(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplArrayMethod(executeData, return_value, "natcasesort", b.SizeOf("\"natcasesort\"")-1, SPL_ARRAY_METHOD_NO_ARG)
}
func zim_spl_Array_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var entry *types.Zval
	var aht *types.Array = SplArrayGetHashTable(intern)
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if b.Assign(&entry, types.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, intern))) == nil {
		return
	}
	if entry.IsIndirect() {
		entry = entry.Indirect()
		if entry.IsUndef() {
			return
		}
	}
	types.ZVAL_COPY_DEREF(return_value, entry)
}
func zim_spl_Array_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	SplArrayIteratorKey(zend.ZEND_THIS(executeData), return_value)
}
func SplArrayIteratorKey(object *types.Zval, return_value *types.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *types.Array = SplArrayGetHashTable(intern)
	types.ZendHashGetCurrentKeyZvalEx(aht, return_value, SplArrayGetPosPtr(aht, intern))
}
func zim_spl_Array_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *types.Array = SplArrayGetHashTable(intern)
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	SplArrayNextEx(intern, aht)
}
func zim_spl_Array_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *types.Array = SplArrayGetHashTable(intern)
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetBool(types.ZendHashHasMoreElementsEx(aht, SplArrayGetPosPtr(aht, intern)))
	return
}
func zim_spl_Array_hasChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var entry *types.Zval
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *types.Array = SplArrayGetHashTable(intern)
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if b.Assign(&entry, types.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, intern))) == nil {
		return_value.SetFalse()
		return
	}
	if entry.IsIndirect() {
		entry = entry.Indirect()
	}
	entry = types.ZVAL_DEREF(entry)
	return_value.SetBool(entry.IsType(types.IS_ARRAY) || entry.IsType(types.IS_OBJECT) && !intern.IsChildArraysOnly())
	return
}
func zim_spl_Array_getChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var entry *types.Zval
	var flags types.Zval
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *types.Array = SplArrayGetHashTable(intern)
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if b.Assign(&entry, types.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, intern))) == nil {
		return
	}
	if entry.IsIndirect() {
		entry = entry.Indirect()
	}
	entry = types.ZVAL_DEREF(entry)
	if entry.IsType(types.IS_OBJECT) {
		if intern.IsChildArraysOnly() {
			return
		}
		if operators.InstanceofFunction(types.Z_OBJCE_P(entry), types.Z_OBJCE_P(zend.ZEND_THIS(executeData))) != 0 {
			return_value.SetObject(entry.Object())
			// 			return_value.AddRefcount()
			return
		}
	}
	flags.SetLong(intern.GetArFlags())
	SplInstantiateArgEx2(types.Z_OBJCE_P(zend.ZEND_THIS(executeData)), return_value, entry, &flags)
}
func zim_spl_Array_serialize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var members types.Zval
	var flags types.Zval
	var var_hash standard.PhpSerializeDataT
	var buf zend.SmartStr = zend.MakeSmartStr(0)
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	standard.PHP_VAR_SERIALIZE_INIT(var_hash)
	flags.SetLong(intern.GetArFlags() & SPL_ARRAY_CLONE_MASK)

	/* storage */

	buf.AppendString("x:")
	standard.PhpVarSerialize(&buf, &flags, &var_hash)
	if !intern.IsIsSelf() {
		standard.PhpVarSerialize(&buf, intern.GetArray(), &var_hash)
		buf.AppendByte(';')
	}

	/* members */

	buf.AppendString("m:")
	if intern.GetStd().GetProperties() == nil {
		zend.RebuildObjectProperties(intern.GetStd())
	}
	members.SetArray(intern.GetStd().GetProperties())
	standard.PhpVarSerialize(&buf, &members, &var_hash)

	/* done */

	standard.PHP_VAR_SERIALIZE_DESTROY(var_hash)
	if buf.GetS() != nil {
		return_value.SetString(buf.GetS())
		return
	}
	return_value.SetNull()
	return
}
func zim_spl_Array_unserialize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var buf *byte
	var buf_len int
	var p *uint8
	var s *uint8
	var var_hash standard.PhpUnserializeDataT
	var members *types.Zval
	var zflags *types.Zval
	var array *types.Zval
	var flags zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "s", &buf, &buf_len) == types.FAILURE {
		return
	}
	if buf_len == 0 {
		return
	}
	if intern.GetNApplyCount() > 0 {
		faults.Error(faults.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return
	}

	/* storage */

	p = (*uint8)(buf)
	s = p
	standard.PHP_VAR_UNSERIALIZE_INIT(var_hash)
	if (*p) != 'x' || (*(b.PreInc(&p))) != ':' {
		goto outexcept
	}
	p++
	zflags = standard.VarTmpVar(&var_hash)
	if standard.PhpVarUnserialize(zflags, &p, s+buf_len, &var_hash) == 0 || zflags.GetType() != types.IS_LONG {
		goto outexcept
	}
	p--
	flags = zflags.Long()

	/* flags needs to be verified and we also need to verify whether the next
	 * thing we get is ';'. After that we require an 'm' or something else
	 * where 'm' stands for members and anything else should be an array. If
	 * neither 'a' or 'm' follows we have an error. */

	if (*p) != ';' {
		goto outexcept
	}
	p++
	if (flags & SPL_ARRAY_IS_SELF) != 0 {

		/* If IS_SELF is used, the flags are not followed by an array/object */

		intern.SetIsCloneMask(false)
		intern.AddArFlags(flags & SPL_ARRAY_CLONE_MASK)
		// zend.ZvalPtrDtor(intern.GetArray())
		intern.GetArray().SetUndef()
	} else {
		if (*p) != 'a' && (*p) != 'O' && (*p) != 'C' && (*p) != 'r' {
			goto outexcept
		}
		array = standard.VarTmpVar(&var_hash)
		if standard.PhpVarUnserialize(array, &p, s+buf_len, &var_hash) == 0 || array.GetType() != types.IS_ARRAY && array.GetType() != types.IS_OBJECT {
			goto outexcept
		}
		intern.SetIsCloneMask(false)
		intern.AddArFlags(flags & SPL_ARRAY_CLONE_MASK)
		if array.IsType(types.IS_ARRAY) {
			// zend.ZvalPtrDtor(intern.GetArray())
			types.ZVAL_COPY_VALUE(intern.GetArray(), array)
			array.SetNull()
			types.SeparateArray(intern.GetArray())
		} else {
			SplArraySetArray(object, intern, array, 0, 1)
		}
		if (*p) != ';' {
			goto outexcept
		}
		p++
	}

	/* members */

	if (*p) != 'm' || (*(b.PreInc(&p))) != ':' {
		goto outexcept
	}
	p++
	members = standard.VarTmpVar(&var_hash)
	if standard.PhpVarUnserialize(members, &p, s+buf_len, &var_hash) == 0 || members.GetType() != types.IS_ARRAY {
		goto outexcept
	}

	/* copy members */

	zend.ObjectPropertiesLoad(intern.GetStd(), members.Array())

	/* done reading $serialized */

	standard.PHP_VAR_UNSERIALIZE_DESTROY(var_hash)
	return
outexcept:
	standard.PHP_VAR_UNSERIALIZE_DESTROY(var_hash)
	faults.ThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Error at offset "+zend.ZEND_LONG_FMT+" of %zd bytes", zend_long((*byte)(p-buf)), buf_len)
	return
}
func zim_spl_Array___serialize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(zend.ZEND_THIS(executeData))
	var tmp types.Zval
	if !executeData.CheckNumArgsNone(true) {
		return
	}
	zend.ArrayInit(return_value)

	/* flags */

	tmp.SetLong(intern.GetArFlags() & SPL_ARRAY_CLONE_MASK)
	return_value.Array().Append(&tmp)

	/* storage */

	if intern.IsIsSelf() {
		tmp.SetNull()
	} else {
		types.ZVAL_COPY(&tmp, intern.GetArray())
	}
	return_value.Array().Append(&tmp)

	/* members */

	tmp.SetArray(zend.ZendStdGetProperties(zend.ZEND_THIS(executeData)))
	//tmp.TryAddRefcount()
	return_value.Array().Append(&tmp)

	/* iterator class */

	if intern.GetCeGetIterator() == spl_ce_ArrayIterator {
		tmp.SetNull()
	} else {
		tmp.SetStringCopy(intern.GetCeGetIterator().GetName())
	}
	return_value.Array().Append(&tmp)
}
func zim_spl_Array___unserialize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(zend.ZEND_THIS(executeData))
	var data *types.Array
	var flags_zv *types.Zval
	var storage_zv *types.Zval
	var members_zv *types.Zval
	var iterator_class_zv *types.Zval
	var flags zend.ZendLong
	if zend.ZendParseParametersThrow(executeData.NumArgs(), "h", &data) == types.FAILURE {
		return
	}
	flags_zv = data.IndexFind(0)
	storage_zv = data.IndexFind(1)
	members_zv = data.IndexFind(2)
	iterator_class_zv = data.IndexFind(3)
	if flags_zv == nil || storage_zv == nil || members_zv == nil || flags_zv.GetType() != types.IS_LONG || members_zv.GetType() != types.IS_ARRAY || iterator_class_zv != nil && (iterator_class_zv.GetType() != types.IS_NULL && iterator_class_zv.GetType() != types.IS_STRING) {
		faults.ThrowException(spl_ce_UnexpectedValueException, "Incomplete or ill-typed serialization data", 0)
		return
	}
	flags = flags_zv.Long()
	intern.SetIsCloneMask(false)
	intern.AddArFlags(flags & SPL_ARRAY_CLONE_MASK)
	if (flags & SPL_ARRAY_IS_SELF) != 0 {
		// zend.ZvalPtrDtor(intern.GetArray())
		intern.GetArray().SetUndef()
	} else {
		SplArraySetArray(zend.ZEND_THIS(executeData), intern, storage_zv, 0, 1)
	}
	zend.ObjectPropertiesLoad(intern.GetStd(), members_zv.Array())
	if iterator_class_zv != nil && iterator_class_zv.IsString() {
		var ce *types.ClassEntry = zend.ZendLookupClass(iterator_class_zv.String())
		if ce == nil {
			faults.ThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Cannot deserialize ArrayObject with iterator class '%s'; no such class exists", iterator_class_zv.String().GetVal())
			return
		} else if operators.InstanceofFunction(ce, spl_ce_Iterator) == 0 {
			faults.ThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Cannot deserialize ArrayObject with iterator class '%s'; this class does not implement the Iterator interface", iterator_class_zv.String().GetVal())
			return
		} else {
			intern.SetCeGetIterator(ce)
		}
	}
}
func zim_spl_Array___debugInfo(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetArray(SplArrayGetDebugInfo(zend.getThis()))
	return
}
func ZmStartupSplArray(type_ int, module_number int) int {
	spl_ce_ArrayObject = zend.RegisterClass("ArrayObject", SplArrayObjectNew, spl_funcs_ArrayObject)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, spl_ce_Aggregate)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, spl_ce_ArrayAccess)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, spl_ce_Serializable)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, spl_ce_Countable)

	spl_handler_ArrayObject = *types.NewObjectHandlersEx(zend.StdObjectHandlersPtr, types.ObjectHandlersSetting{
		Offset:            int((*byte)(&((*SplArrayObject)(nil).GetStd())) - (*byte)(nil)),
		CloneObj:          SplArrayObjectClone,
		ReadDimension:     SplArrayReadDimension,
		WriteDimension:    SplArrayWriteDimension,
		UnsetDimension:    SplArrayUnsetDimension,
		HasDimension:      SplArrayHasDimension,
		CountElements:     SplArrayObjectCountElements,
		GetPropertiesFor:  SplArrayGetPropertiesFor,
		ReadProperty:      SplArrayReadProperty,
		WriteProperty:     SplArrayWriteProperty,
		GetPropertyPtrPtr: SplArrayGetPropertyPtrPtr,
		HasProperty:       SplArrayHasProperty,
		UnsetProperty:     SplArrayUnsetProperty,
		CompareObjects:    SplArrayCompareObjects,
		FreeObj:           SplArrayObjectFreeStorage,
	})

	spl_ce_ArrayIterator = zend.RegisterClass("ArrayIterator", SplArrayObjectNew, spl_funcs_ArrayIterator)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_ArrayAccess)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_SeekableIterator)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_Serializable)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_Countable)

	memcpy(&spl_handler_ArrayIterator, &spl_handler_ArrayObject, b.SizeOf("zend_object_handlers"))
	spl_ce_ArrayIterator.SetGetIterator(SplArrayGetIterator)
	spl_ce_ArrayIterator.AddCeFlags(zend.AccReuseGetIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayObject, "STD_PROP_LIST", zend.ZendLong(SPL_ARRAY_STD_PROP_LIST))
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayObject, "ARRAY_AS_PROPS", zend.ZendLong(SPL_ARRAY_ARRAY_AS_PROPS))
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayIterator, "STD_PROP_LIST", zend.ZendLong(SPL_ARRAY_STD_PROP_LIST))
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayIterator, "ARRAY_AS_PROPS", zend.ZendLong(SPL_ARRAY_ARRAY_AS_PROPS))

	spl_ce_RecursiveArrayIterator = zend.RegisterSubClass(spl_ce_ArrayIterator, "RecursiveArrayIterator", SplArrayObjectNew, spl_funcs_RecursiveArrayIterator)
	zend.ZendClassImplements(spl_ce_RecursiveArrayIterator, 1, spl_ce_RecursiveIterator)
	spl_ce_RecursiveArrayIterator.SetGetIterator(SplArrayGetIterator)
	spl_ce_RecursiveArrayIterator.AddCeFlags(zend.AccReuseGetIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveArrayIterator, "CHILD_ARRAYS_ONLY", zend.ZendLong(SPL_ARRAY_CHILD_ARRAYS_ONLY))

	return types.SUCCESS
}
