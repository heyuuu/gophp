// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	"sik/zend"
)

func SplArrayFromObj(obj *zend.ZendObject) *SplArrayObject {
	return (*SplArrayObject)((*byte)(obj - zend_long((*byte)(&((*SplArrayObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLARRAY_P(zv *zend.Zval) *SplArrayObject { return SplArrayFromObj(zend.Z_OBJ_P(zv)) }
func SplArrayGetHashTablePtr(intern *SplArrayObject) **zend.HashTable {
	//??? TODO: Delay duplication for arrays; only duplicate for write operations

	if intern.IsIsSelf() {
		if intern.GetStd().GetProperties() == nil {
			zend.RebuildObjectProperties(&intern.GetStd())
		}
		return &intern.GetStd().GetProperties()
	} else if intern.IsUseOther() {
		var other *SplArrayObject = Z_SPLARRAY_P(&intern.GetArray())
		return SplArrayGetHashTablePtr(other)
	} else if intern.GetArray().IsType(zend.IS_ARRAY) {
		return &zend.Z_ARRVAL(intern.GetArray())
	} else {
		var obj *zend.ZendObject = zend.Z_OBJ(intern.GetArray())
		if obj.GetProperties() == nil {
			zend.RebuildObjectProperties(obj)
		} else if zend.GC_REFCOUNT(obj.GetProperties()) > 1 {
			if (zend.GC_FLAGS(obj.GetProperties()) & zend.IS_ARRAY_IMMUTABLE) == 0 {
				zend.GC_DELREF(obj.GetProperties())
			}
			obj.SetProperties(zend.ZendArrayDup(obj.GetProperties()))
		}
		return &obj.GetProperties()
	}

	//??? TODO: Delay duplication for arrays; only duplicate for write operations
}
func SplArrayGetHashTable(intern *SplArrayObject) *zend.HashTable {
	return (*SplArrayGetHashTablePtr)(intern)
}
func SplArrayReplaceHashTable(intern *SplArrayObject, ht *zend.HashTable) {
	var ht_ptr **zend.HashTable = SplArrayGetHashTablePtr(intern)
	zend.ZendArrayDestroy(*ht_ptr)
	*ht_ptr = ht
}
func SplArrayIsObject(intern *SplArrayObject) zend.ZendBool {
	for intern.IsUseOther() {
		intern = Z_SPLARRAY_P(&intern.GetArray())
	}
	return intern.IsIsSelf() || intern.GetArray().IsType(zend.IS_OBJECT)
}
func SplArrayCreateHtIter(ht *zend.HashTable, intern *SplArrayObject) {
	intern.SetHtIter(zend.ZendHashIteratorAdd(ht, zend.ZendHashGetCurrentPos(ht)))
	zend.ZendHashInternalPointerResetEx(ht, &zend.ExecutorGlobals.GetHtIterators()[intern.GetHtIter()].GetPos())
	SplArraySkipProtected(intern, ht)
}
func SplArrayGetPosPtr(ht *zend.HashTable, intern *SplArrayObject) *uint32 {
	if intern.GetHtIter() == uint32-1 {
		SplArrayCreateHtIter(ht, intern)
	}
	return &zend.ExecutorGlobals.GetHtIterators()[intern.GetHtIter()].GetPos()
}
func SplArrayObjectFreeStorage(object *zend.ZendObject) {
	var intern *SplArrayObject = SplArrayFromObj(object)
	if intern.GetHtIter() != uint32-1 {
		zend.ZendHashIteratorDel(intern.GetHtIter())
	}
	zend.ZendObjectStdDtor(&intern.GetStd())
	zend.ZvalPtrDtor(&intern.GetArray())
}
func SplArrayObjectNewEx(class_type *zend.ZendClassEntry, orig *zend.Zval, clone_orig int) *zend.ZendObject {
	var intern *SplArrayObject
	var parent *zend.ZendClassEntry = class_type
	var inherited int = 0
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_array_object"), parent)
	zend.ZendObjectStdInit(&intern.GetStd(), class_type)
	zend.ObjectPropertiesInit(&intern.GetStd(), class_type)
	intern.SetArFlags(0)
	intern.SetCeGetIterator(spl_ce_ArrayIterator)
	if orig != nil {
		var other *SplArrayObject = Z_SPLARRAY_P(orig)
		intern.SetIsCloneMask(false)
		intern.AddArFlags(other.GetArFlags() & SPL_ARRAY_CLONE_MASK)
		intern.SetCeGetIterator(other.GetCeGetIterator())
		if clone_orig != 0 {
			if other.IsIsSelf() {
				zend.ZVAL_UNDEF(&intern.GetArray())
			} else if zend.Z_OBJ_HT_P(orig) == &spl_handler_ArrayObject {
				zend.ZVAL_ARR(&intern.GetArray(), zend.ZendArrayDup(SplArrayGetHashTable(other)))
			} else {
				zend.ZEND_ASSERT(zend.Z_OBJ_HT_P(orig) == &spl_handler_ArrayIterator)
				zend.ZVAL_COPY(&intern.GetArray(), orig)
				intern.SetIsUseOther(true)
			}
		} else {
			zend.ZVAL_COPY(&intern.GetArray(), orig)
			intern.SetIsUseOther(true)
		}
	} else {
		zend.ArrayInit(&intern.GetArray())
	}
	for parent != nil {
		if parent == spl_ce_ArrayIterator || parent == spl_ce_RecursiveArrayIterator {
			intern.GetStd().SetHandlers(&spl_handler_ArrayIterator)
			break
		} else if parent == spl_ce_ArrayObject {
			intern.GetStd().SetHandlers(&spl_handler_ArrayObject)
			break
		}
		parent = parent.parent
		inherited = 1
	}
	if parent == nil {
		core.PhpErrorDocref(nil, zend.E_COMPILE_ERROR, "Internal compiler error, Class is not child of ArrayObject or ArrayIterator")
	}
	if inherited != 0 {
		intern.SetFptrOffsetGet(zend.ZendHashStrFindPtr(&class_type.GetFunctionTable(), "offsetget", b.SizeOf("\"offsetget\"")-1))
		if intern.GetFptrOffsetGet().GetScope() == parent {
			intern.SetFptrOffsetGet(nil)
		}
		intern.SetFptrOffsetSet(zend.ZendHashStrFindPtr(&class_type.GetFunctionTable(), "offsetset", b.SizeOf("\"offsetset\"")-1))
		if intern.GetFptrOffsetSet().GetScope() == parent {
			intern.SetFptrOffsetSet(nil)
		}
		intern.SetFptrOffsetHas(zend.ZendHashStrFindPtr(&class_type.GetFunctionTable(), "offsetexists", b.SizeOf("\"offsetexists\"")-1))
		if intern.GetFptrOffsetHas().GetScope() == parent {
			intern.SetFptrOffsetHas(nil)
		}
		intern.SetFptrOffsetDel(zend.ZendHashStrFindPtr(&class_type.GetFunctionTable(), "offsetunset", b.SizeOf("\"offsetunset\"")-1))
		if intern.GetFptrOffsetDel().GetScope() == parent {
			intern.SetFptrOffsetDel(nil)
		}
		intern.SetFptrCount(zend.ZendHashStrFindPtr(&class_type.GetFunctionTable(), "count", b.SizeOf("\"count\"")-1))
		if intern.GetFptrCount().GetScope() == parent {
			intern.SetFptrCount(nil)
		}
	}

	/* Cache iterator functions if ArrayIterator or derived. Check current's */

	if intern.GetStd().GetHandlers() == &spl_handler_ArrayIterator {
		var funcs_ptr *zend.ZendClassIteratorFuncs = class_type.GetIteratorFuncsPtr()
		if funcs_ptr.GetZfCurrent() == nil {
			funcs_ptr.SetZfRewind(zend.ZendHashStrFindPtr(&class_type.GetFunctionTable(), "rewind", b.SizeOf("\"rewind\"")-1))
			funcs_ptr.SetZfValid(zend.ZendHashStrFindPtr(&class_type.GetFunctionTable(), "valid", b.SizeOf("\"valid\"")-1))
			funcs_ptr.SetZfKey(zend.ZendHashStrFindPtr(&class_type.GetFunctionTable(), "key", b.SizeOf("\"key\"")-1))
			funcs_ptr.SetZfCurrent(zend.ZendHashStrFindPtr(&class_type.GetFunctionTable(), "current", b.SizeOf("\"current\"")-1))
			funcs_ptr.SetZfNext(zend.ZendHashStrFindPtr(&class_type.GetFunctionTable(), "next", b.SizeOf("\"next\"")-1))
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
	return &intern.GetStd()
}
func SplArrayObjectNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplArrayObjectNewEx(class_type, nil, 0)
}
func SplArrayObjectClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	old_object = zend.Z_OBJ_P(zobject)
	new_object = SplArrayObjectNewEx(old_object.GetCe(), zobject, 1)
	zend.ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
func SplArrayGetDimensionPtr(check_inherited int, intern *SplArrayObject, offset *zend.Zval, type_ int) *zend.Zval {
	var retval *zend.Zval
	var index zend.ZendLong
	var offset_key *zend.ZendString
	var ht *zend.HashTable = SplArrayGetHashTable(intern)
	if offset == nil || zend.Z_ISUNDEF_P(offset) || ht == nil {
		return &(zend.ExecutorGlobals.GetUninitializedZval())
	}
	if (type_ == zend.BP_VAR_W || type_ == zend.BP_VAR_RW) && intern.GetNApplyCount() > 0 {
		zend.ZendError(zend.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return &(zend.ExecutorGlobals.GetErrorZval())
	}
try_again:
	switch zend.Z_TYPE_P(offset) {
	case zend.IS_NULL:
		offset_key = zend.ZSTR_EMPTY_ALLOC()
		goto fetch_dim_string
	case zend.IS_STRING:
		offset_key = zend.Z_STR_P(offset)
	fetch_dim_string:
		retval = zend.ZendSymtableFind(ht, offset_key)
		if retval != nil {
			if zend.Z_TYPE_P(retval) == zend.IS_INDIRECT {
				retval = zend.Z_INDIRECT_P(retval)
				if zend.Z_TYPE_P(retval) == zend.IS_UNDEF {
					switch type_ {
					case zend.BP_VAR_R:
						zend.ZendError(zend.E_NOTICE, "Undefined index: %s", zend.ZSTR_VAL(offset_key))
					case zend.BP_VAR_UNSET:

					case zend.BP_VAR_IS:
						retval = &(zend.ExecutorGlobals.GetUninitializedZval())
						break
					case zend.BP_VAR_RW:
						zend.ZendError(zend.E_NOTICE, "Undefined index: %s", zend.ZSTR_VAL(offset_key))
					case zend.BP_VAR_W:
						zend.ZVAL_NULL(retval)
					}
				}
			}
		} else {
			switch type_ {
			case zend.BP_VAR_R:
				zend.ZendError(zend.E_NOTICE, "Undefined index: %s", zend.ZSTR_VAL(offset_key))
			case zend.BP_VAR_UNSET:

			case zend.BP_VAR_IS:
				retval = &(zend.ExecutorGlobals.GetUninitializedZval())
				break
			case zend.BP_VAR_RW:
				zend.ZendError(zend.E_NOTICE, "Undefined index: %s", zend.ZSTR_VAL(offset_key))
			case zend.BP_VAR_W:
				var value zend.Zval
				zend.ZVAL_NULL(&value)
				retval = zend.ZendSymtableUpdate(ht, offset_key, &value)
			}
		}
		return retval
	case zend.IS_RESOURCE:
		zend.ZendError(zend.E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", zend.Z_RES_P(offset).GetHandle(), zend.Z_RES_P(offset).GetHandle())
		index = zend.Z_RES_P(offset).GetHandle()
		goto num_index
	case zend.IS_DOUBLE:
		index = zend.ZendLong(zend.Z_DVAL_P(offset))
		goto num_index
	case zend.IS_FALSE:
		index = 0
		goto num_index
	case zend.IS_TRUE:
		index = 1
		goto num_index
	case zend.IS_LONG:
		index = zend.Z_LVAL_P(offset)
	num_index:
		if b.Assign(&retval, zend.ZendHashIndexFind(ht, index)) == nil {
			switch type_ {
			case zend.BP_VAR_R:
				zend.ZendError(zend.E_NOTICE, "Undefined offset: "+zend.ZEND_LONG_FMT, index)
			case zend.BP_VAR_UNSET:

			case zend.BP_VAR_IS:
				retval = &(zend.ExecutorGlobals.GetUninitializedZval())
				break
			case zend.BP_VAR_RW:
				zend.ZendError(zend.E_NOTICE, "Undefined offset: "+zend.ZEND_LONG_FMT, index)
			case zend.BP_VAR_W:
				var value zend.Zval
				zend.ZVAL_UNDEF(&value)
				retval = zend.ZendHashIndexUpdate(ht, index, &value)
			}
		}
		return retval
	case zend.IS_REFERENCE:
		zend.ZVAL_DEREF(offset)
		goto try_again
	default:
		zend.ZendError(zend.E_WARNING, "Illegal offset type")
		if type_ == zend.BP_VAR_W || type_ == zend.BP_VAR_RW {
			return &(zend.ExecutorGlobals.GetErrorZval())
		} else {
			return &(zend.ExecutorGlobals.GetUninitializedZval())
		}
	}
}
func SplArrayReadDimensionEx(check_inherited int, object *zend.Zval, offset *zend.Zval, type_ int, rv *zend.Zval) *zend.Zval {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var ret *zend.Zval
	if check_inherited != 0 && (intern.GetFptrOffsetGet() != nil || type_ == zend.BP_VAR_IS && intern.GetFptrOffsetHas() != nil) {
		if type_ == zend.BP_VAR_IS {
			if SplArrayHasDimension(object, offset, 0) == 0 {
				return &(zend.ExecutorGlobals.GetUninitializedZval())
			}
		}
		if intern.GetFptrOffsetGet() != nil {
			var tmp zend.Zval
			if offset == nil {
				zend.ZVAL_UNDEF(&tmp)
				offset = &tmp
			} else {
				zend.SEPARATE_ARG_IF_REF(offset)
			}
			zend.ZendCallMethodWith1Params(object, zend.Z_OBJCE_P(object), &intern.GetFptrOffsetGet(), "offsetGet", rv, offset)
			zend.ZvalPtrDtor(offset)
			if !(zend.Z_ISUNDEF_P(rv)) {
				return rv
			}
			return &(zend.ExecutorGlobals.GetUninitializedZval())
		}
	}
	ret = SplArrayGetDimensionPtr(check_inherited, intern, offset, type_)

	/* When in a write context,
	 * ZE has to be fooled into thinking this is in a reference set
	 * by separating (if necessary) and returning as IS_REFERENCE (with refcount == 1)
	 */

	if (type_ == zend.BP_VAR_W || type_ == zend.BP_VAR_RW || type_ == zend.BP_VAR_UNSET) && !(zend.Z_ISREF_P(ret)) && ret != &(zend.ExecutorGlobals.GetUninitializedZval()) {
		zend.ZVAL_NEW_REF(ret, ret)
	}
	return ret
}
func SplArrayReadDimension(object *zend.Zval, offset *zend.Zval, type_ int, rv *zend.Zval) *zend.Zval {
	return SplArrayReadDimensionEx(1, object, offset, type_, rv)
}
func SplArrayWriteDimensionEx(check_inherited int, object *zend.Zval, offset *zend.Zval, value *zend.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var index zend.ZendLong
	var ht *zend.HashTable
	if check_inherited != 0 && intern.GetFptrOffsetSet() != nil {
		var tmp zend.Zval
		if offset == nil {
			zend.ZVAL_NULL(&tmp)
			offset = &tmp
		} else {
			zend.SEPARATE_ARG_IF_REF(offset)
		}
		zend.ZendCallMethodWith2Params(object, zend.Z_OBJCE_P(object), &intern.GetFptrOffsetSet(), "offsetSet", nil, offset, value)
		zend.ZvalPtrDtor(offset)
		return
	}
	if intern.GetNApplyCount() > 0 {
		zend.ZendError(zend.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return
	}
	zend.Z_TRY_ADDREF_P(value)
	if offset == nil {
		ht = SplArrayGetHashTable(intern)
		zend.ZendHashNextIndexInsert(ht, value)
		return
	}
try_again:
	switch zend.Z_TYPE_P(offset) {
	case zend.IS_STRING:
		ht = SplArrayGetHashTable(intern)
		zend.ZendSymtableUpdateInd(ht, zend.Z_STR_P(offset), value)
		return
	case zend.IS_DOUBLE:
		index = zend.ZendLong(zend.Z_DVAL_P(offset))
		goto num_index
	case zend.IS_RESOURCE:
		index = zend.Z_RES_HANDLE_P(offset)
		goto num_index
	case zend.IS_FALSE:
		index = 0
		goto num_index
	case zend.IS_TRUE:
		index = 1
		goto num_index
	case zend.IS_LONG:
		index = zend.Z_LVAL_P(offset)
	num_index:
		ht = SplArrayGetHashTable(intern)
		zend.ZendHashIndexUpdate(ht, index, value)
		return
	case zend.IS_NULL:
		ht = SplArrayGetHashTable(intern)
		zend.ZendHashNextIndexInsert(ht, value)
		return
	case zend.IS_REFERENCE:
		zend.ZVAL_DEREF(offset)
		goto try_again
	default:
		zend.ZendError(zend.E_WARNING, "Illegal offset type")
		zend.ZvalPtrDtor(value)
		return
	}
}
func SplArrayWriteDimension(object *zend.Zval, offset *zend.Zval, value *zend.Zval) {
	SplArrayWriteDimensionEx(1, object, offset, value)
}
func SplArrayUnsetDimensionEx(check_inherited int, object *zend.Zval, offset *zend.Zval) {
	var index zend.ZendLong
	var ht *zend.HashTable
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if check_inherited != 0 && intern.GetFptrOffsetDel() != nil {
		zend.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, zend.Z_OBJCE_P(object), &intern.GetFptrOffsetDel(), "offsetUnset", nil, offset)
		zend.ZvalPtrDtor(offset)
		return
	}
	if intern.GetNApplyCount() > 0 {
		zend.ZendError(zend.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return
	}
try_again:
	switch zend.Z_TYPE_P(offset) {
	case zend.IS_STRING:
		ht = SplArrayGetHashTable(intern)
		if ht == &(zend.ExecutorGlobals.GetSymbolTable()) {
			if zend.ZendDeleteGlobalVariable(zend.Z_STR_P(offset)) != 0 {
				zend.ZendError(zend.E_NOTICE, "Undefined index: %s", zend.Z_STRVAL_P(offset))
			}
		} else {
			var data *zend.Zval = zend.ZendSymtableFind(ht, zend.Z_STR_P(offset))
			if data != nil {
				if zend.Z_TYPE_P(data) == zend.IS_INDIRECT {
					data = zend.Z_INDIRECT_P(data)
					if zend.Z_TYPE_P(data) == zend.IS_UNDEF {
						zend.ZendError(zend.E_NOTICE, "Undefined index: %s", zend.Z_STRVAL_P(offset))
					} else {
						zend.ZvalPtrDtor(data)
						zend.ZVAL_UNDEF(data)
						zend.HT_FLAGS(ht) |= zend.HASH_FLAG_HAS_EMPTY_IND
						zend.ZendHashMoveForwardEx(ht, SplArrayGetPosPtr(ht, intern))
						if SplArrayIsObject(intern) != 0 {
							SplArraySkipProtected(intern, ht)
						}
					}
				} else if zend.ZendSymtableDel(ht, zend.Z_STR_P(offset)) == zend.FAILURE {
					zend.ZendError(zend.E_NOTICE, "Undefined index: %s", zend.Z_STRVAL_P(offset))
				}
			} else {
				zend.ZendError(zend.E_NOTICE, "Undefined index: %s", zend.Z_STRVAL_P(offset))
			}
		}
		break
	case zend.IS_DOUBLE:
		index = zend.ZendLong(zend.Z_DVAL_P(offset))
		goto num_index
	case zend.IS_RESOURCE:
		index = zend.Z_RES_HANDLE_P(offset)
		goto num_index
	case zend.IS_FALSE:
		index = 0
		goto num_index
	case zend.IS_TRUE:
		index = 1
		goto num_index
	case zend.IS_LONG:
		index = zend.Z_LVAL_P(offset)
	num_index:
		ht = SplArrayGetHashTable(intern)
		if zend.ZendHashIndexDel(ht, index) == zend.FAILURE {
			zend.ZendError(zend.E_NOTICE, "Undefined offset: "+zend.ZEND_LONG_FMT, index)
		}
		break
	case zend.IS_REFERENCE:
		zend.ZVAL_DEREF(offset)
		goto try_again
	default:
		zend.ZendError(zend.E_WARNING, "Illegal offset type")
		return
	}
}
func SplArrayUnsetDimension(object *zend.Zval, offset *zend.Zval) {
	SplArrayUnsetDimensionEx(1, object, offset)
}
func SplArrayHasDimensionEx(check_inherited int, object *zend.Zval, offset *zend.Zval, check_empty int) int {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var index zend.ZendLong
	var rv zend.Zval
	var value *zend.Zval = nil
	var tmp *zend.Zval
	if check_inherited != 0 && intern.GetFptrOffsetHas() != nil {
		zend.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(object, zend.Z_OBJCE_P(object), &intern.GetFptrOffsetHas(), "offsetExists", &rv, offset)
		zend.ZvalPtrDtor(offset)
		if zend.ZendIsTrue(&rv) != 0 {
			zend.ZvalPtrDtor(&rv)
			if check_empty != 1 {
				return 1
			} else if intern.GetFptrOffsetGet() != nil {
				value = SplArrayReadDimensionEx(1, object, offset, zend.BP_VAR_R, &rv)
			}
		} else {
			zend.ZvalPtrDtor(&rv)
			return 0
		}
	}
	if value == nil {
		var ht *zend.HashTable = SplArrayGetHashTable(intern)
	try_again:
		switch zend.Z_TYPE_P(offset) {
		case zend.IS_STRING:
			if b.Assign(&tmp, zend.ZendSymtableFind(ht, zend.Z_STR_P(offset))) != nil {
				if check_empty == 2 {
					return 1
				}
			} else {
				return 0
			}
			break
		case zend.IS_DOUBLE:
			index = zend.ZendLong(zend.Z_DVAL_P(offset))
			goto num_index
		case zend.IS_RESOURCE:
			index = zend.Z_RES_HANDLE_P(offset)
			goto num_index
		case zend.IS_FALSE:
			index = 0
			goto num_index
		case zend.IS_TRUE:
			index = 1
			goto num_index
		case zend.IS_LONG:
			index = zend.Z_LVAL_P(offset)
		num_index:
			if b.Assign(&tmp, zend.ZendHashIndexFind(ht, index)) != nil {
				if check_empty == 2 {
					return 1
				}
			} else {
				return 0
			}
			break
		case zend.IS_REFERENCE:
			zend.ZVAL_DEREF(offset)
			goto try_again
		default:
			zend.ZendError(zend.E_WARNING, "Illegal offset type")
			return 0
		}
		if check_empty != 0 && check_inherited != 0 && intern.GetFptrOffsetGet() != nil {
			value = SplArrayReadDimensionEx(1, object, offset, zend.BP_VAR_R, &rv)
		} else {
			value = tmp
		}
	}
	var result zend.ZendBool = b.CondF(check_empty != 0, func() int { return zend.ZendIsTrue(value) }, func() bool { return zend.Z_TYPE_P(value) != zend.IS_NULL })
	if value == &rv {
		zend.ZvalPtrDtor(&rv)
	}
	return result
}
func SplArrayHasDimension(object *zend.Zval, offset *zend.Zval, check_empty int) int {
	return SplArrayHasDimensionEx(1, object, offset, check_empty)
}
func zim_spl_Array_offsetExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var index *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &index) == zend.FAILURE {
		return
	}
	zend.RETVAL_BOOL(SplArrayHasDimensionEx(0, zend.ZEND_THIS, index, 2) != 0)
	return
}
func zim_spl_Array_offsetGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	var index *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &index) == zend.FAILURE {
		return
	}
	value = SplArrayReadDimensionEx(0, zend.ZEND_THIS, index, zend.BP_VAR_R, return_value)
	if value != return_value {
		zend.ZVAL_COPY_DEREF(return_value, value)
	}
}
func zim_spl_Array_offsetSet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var index *zend.Zval
	var value *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "zz", &index, &value) == zend.FAILURE {
		return
	}
	SplArrayWriteDimensionEx(0, zend.ZEND_THIS, index, value)
}
func SplArrayIteratorAppend(object *zend.Zval, append_value *zend.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if SplArrayIsObject(intern) != 0 {
		zend.ZendThrowError(nil, "Cannot append properties to objects, use %s::offsetSet() instead", zend.ZSTR_VAL(zend.Z_OBJCE_P(object).GetName()))
		return
	}
	SplArrayWriteDimension(object, nil, append_value)
}
func zim_spl_Array_append(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var value *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &value) == zend.FAILURE {
		return
	}
	SplArrayIteratorAppend(zend.ZEND_THIS, value)
}
func zim_spl_Array_offsetUnset(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var index *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &index) == zend.FAILURE {
		return
	}
	SplArrayUnsetDimensionEx(0, zend.ZEND_THIS, index)
}
func zim_spl_Array_getArrayCopy(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	zend.RETVAL_ARR(zend.ZendArrayDup(SplArrayGetHashTable(intern)))
	return
}
func SplArrayGetPropertiesFor(object *zend.Zval, purpose zend.ZendPropPurpose) *zend.HashTable {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var ht *zend.HashTable
	var dup zend.ZendBool
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
		break
	case zend.ZEND_PROP_PURPOSE_VAR_EXPORT:

	case zend.ZEND_PROP_PURPOSE_JSON:

	case zend._ZEND_PROP_PURPOSE_ARRAY_KEY_EXISTS:
		dup = 0
		break
	default:
		return zend.ZendStdGetPropertiesFor(object, purpose)
	}
	ht = SplArrayGetHashTable(intern)
	if dup != 0 {
		ht = zend.ZendArrayDup(ht)
	} else {
		zend.GC_ADDREF(ht)
	}
	return ht
}
func SplArrayGetDebugInfo(obj *zend.Zval) *zend.HashTable {
	var storage *zend.Zval
	var zname *zend.ZendString
	var base *zend.ZendClassEntry
	var intern *SplArrayObject = Z_SPLARRAY_P(obj)
	if intern.GetStd().GetProperties() == nil {
		zend.RebuildObjectProperties(&intern.GetStd())
	}
	if intern.IsIsSelf() {
		return zend.ZendArrayDup(intern.GetStd().GetProperties())
	} else {
		var debug_info *zend.HashTable
		debug_info = zend.ZendNewArray(zend.ZendHashNumElements(intern.GetStd().GetProperties()) + 1)
		zend.ZendHashCopy(debug_info, intern.GetStd().GetProperties(), zend.CopyCtorFuncT(zend.ZvalAddRef))
		storage = &intern.GetArray()
		zend.Z_TRY_ADDREF_P(storage)
		if zend.Z_OBJ_HT_P(obj) == &spl_handler_ArrayIterator {
			base = spl_ce_ArrayIterator
		} else {
			base = spl_ce_ArrayObject
		}
		zname = SplGenPrivatePropName(base, "storage", b.SizeOf("\"storage\"")-1)
		zend.ZendSymtableUpdate(debug_info, zname, storage)
		zend.ZendStringReleaseEx(zname, 0)
		return debug_info
	}
}
func SplArrayGetGc(obj *zend.Zval, gc_data **zend.Zval, gc_data_count *int) *zend.HashTable {
	var intern *SplArrayObject = Z_SPLARRAY_P(obj)
	*gc_data = &intern.GetArray()
	*gc_data_count = 1
	return zend.ZendStdGetProperties(obj)
}
func SplArrayReadProperty(object *zend.Zval, member *zend.Zval, type_ int, cache_slot *any, rv *zend.Zval) *zend.Zval {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.IsArrayAsProps() && zend.ZendStdHasProperty(object, member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		return SplArrayReadDimension(object, member, type_, rv)
	}
	return zend.ZendStdReadProperty(object, member, type_, cache_slot, rv)
}
func SplArrayWriteProperty(object *zend.Zval, member *zend.Zval, value *zend.Zval, cache_slot *any) *zend.Zval {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.IsArrayAsProps() && zend.ZendStdHasProperty(object, member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		SplArrayWriteDimension(object, member, value)
		return value
	}
	return zend.ZendStdWriteProperty(object, member, value, cache_slot)
}
func SplArrayGetPropertyPtrPtr(object *zend.Zval, member *zend.Zval, type_ int, cache_slot *any) *zend.Zval {
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
func SplArrayHasProperty(object *zend.Zval, member *zend.Zval, has_set_exists int, cache_slot *any) int {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.IsArrayAsProps() && zend.ZendStdHasProperty(object, member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		return SplArrayHasDimension(object, member, has_set_exists)
	}
	return zend.ZendStdHasProperty(object, member, has_set_exists, cache_slot)
}
func SplArrayUnsetProperty(object *zend.Zval, member *zend.Zval, cache_slot *any) {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.IsArrayAsProps() && zend.ZendStdHasProperty(object, member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		SplArrayUnsetDimension(object, member)
		return
	}
	zend.ZendStdUnsetProperty(object, member, cache_slot)
}
func SplArrayCompareObjects(o1 *zend.Zval, o2 *zend.Zval) int {
	var ht1 *zend.HashTable
	var ht2 *zend.HashTable
	var intern1 *SplArrayObject
	var intern2 *SplArrayObject
	var result int = 0
	intern1 = Z_SPLARRAY_P(o1)
	intern2 = Z_SPLARRAY_P(o2)
	ht1 = SplArrayGetHashTable(intern1)
	ht2 = SplArrayGetHashTable(intern2)
	result = zend.ZendCompareSymbolTables(ht1, ht2)

	/* if we just compared std.properties, don't do it again */

	if result == 0 && !(ht1 == intern1.GetStd().GetProperties() && ht2 == intern2.GetStd().GetProperties()) {
		result = zend.ZendStdCompareObjects(o1, o2)
	}
	return result
}
func SplArraySkipProtected(intern *SplArrayObject, aht *zend.HashTable) int {
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	var data *zend.Zval
	if SplArrayIsObject(intern) != 0 {
		var pos_ptr *uint32 = SplArrayGetPosPtr(aht, intern)
		for {
			if zend.ZendHashGetCurrentKeyEx(aht, &string_key, &num_key, pos_ptr) == zend.HASH_KEY_IS_STRING {
				data = zend.ZendHashGetCurrentDataEx(aht, pos_ptr)
				if data != nil && zend.Z_TYPE_P(data) == zend.IS_INDIRECT && zend.Z_TYPE_P(b.Assign(&data, zend.Z_INDIRECT_P(data))) == zend.IS_UNDEF {

				} else if zend.ZSTR_LEN(string_key) == 0 || zend.ZSTR_VAL(string_key)[0] {
					return zend.SUCCESS
				}
			} else {
				return zend.SUCCESS
			}
			if zend.ZendHashHasMoreElementsEx(aht, pos_ptr) != zend.SUCCESS {
				return zend.FAILURE
			}
			zend.ZendHashMoveForwardEx(aht, pos_ptr)

		}
	}
	return zend.FAILURE
}
func SplArrayNextEx(intern *SplArrayObject, aht *zend.HashTable) int {
	var pos_ptr *uint32 = SplArrayGetPosPtr(aht, intern)
	zend.ZendHashMoveForwardEx(aht, pos_ptr)
	if SplArrayIsObject(intern) != 0 {
		return SplArraySkipProtected(intern, aht)
	} else {
		return zend.ZendHashHasMoreElementsEx(aht, pos_ptr)
	}
}
func SplArrayNext(intern *SplArrayObject) int {
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	return SplArrayNextEx(intern, aht)
}
func SplArrayItDtor(iter *zend.ZendObjectIterator) {
	zend.ZendUserItInvalidateCurrent(iter)
	zend.ZvalPtrDtor(&iter.GetData())
}
func SplArrayItValid(iter *zend.ZendObjectIterator) int {
	var object *SplArrayObject = Z_SPLARRAY_P(&iter.GetData())
	var aht *zend.HashTable = SplArrayGetHashTable(object)
	if object.IsOverloadedValid() {
		return zend.ZendUserItValid(iter)
	} else {
		return zend.ZendHashHasMoreElementsEx(aht, SplArrayGetPosPtr(aht, object))
	}
}
func SplArrayItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var object *SplArrayObject = Z_SPLARRAY_P(&iter.GetData())
	var aht *zend.HashTable = SplArrayGetHashTable(object)
	if object.IsOverloadedCurrent() {
		return zend.ZendUserItGetCurrentData(iter)
	} else {
		var data *zend.Zval = zend.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, object))
		if data != nil && zend.Z_TYPE_P(data) == zend.IS_INDIRECT {
			data = zend.Z_INDIRECT_P(data)
		}
		return data
	}
}
func SplArrayItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplArrayObject = Z_SPLARRAY_P(&iter.GetData())
	var aht *zend.HashTable = SplArrayGetHashTable(object)
	if object.IsOverloadedKey() {
		zend.ZendUserItGetCurrentKey(iter, key)
	} else {
		zend.ZendHashGetCurrentKeyZvalEx(aht, key, SplArrayGetPosPtr(aht, object))
	}
}
func SplArrayItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplArrayObject = Z_SPLARRAY_P(&iter.GetData())
	var aht *zend.HashTable = SplArrayGetHashTable(object)
	if object.IsOverloadedNext() {
		zend.ZendUserItMoveForward(iter)
	} else {
		zend.ZendUserItInvalidateCurrent(iter)
		SplArrayNextEx(object, aht)
	}
}
func SplArrayRewind(intern *SplArrayObject) {
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if intern.GetHtIter() == uint32-1 {
		SplArrayGetPosPtr(aht, intern)
	} else {
		zend.ZendHashInternalPointerResetEx(aht, SplArrayGetPosPtr(aht, intern))
		SplArraySkipProtected(intern, aht)
	}
}
func SplArrayItRewind(iter *zend.ZendObjectIterator) {
	var object *SplArrayObject = Z_SPLARRAY_P(&iter.GetData())
	if object.IsOverloadedRewind() {
		zend.ZendUserItRewind(iter)
	} else {
		zend.ZendUserItInvalidateCurrent(iter)
		SplArrayRewind(object)
	}
}
func SplArraySetArray(object *zend.Zval, intern *SplArrayObject, array *zend.Zval, ar_flags zend.ZendLong, just_array int) {
	if zend.Z_TYPE_P(array) != zend.IS_OBJECT && zend.Z_TYPE_P(array) != zend.IS_ARRAY {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Passed variable is not an array or object", 0)
		return
	}
	if zend.Z_TYPE_P(array) == zend.IS_ARRAY {
		zend.ZvalPtrDtor(&intern.GetArray())
		if zend.Z_REFCOUNT_P(array) == 1 {
			zend.ZVAL_COPY(&intern.GetArray(), array)
		} else {

			//??? TODO: try to avoid array duplication

			zend.ZVAL_ARR(&intern.GetArray(), zend.ZendArrayDup(zend.Z_ARR_P(array)))

			//??? TODO: try to avoid array duplication

		}
	} else {
		if zend.Z_OBJ_HT_P(array) == &spl_handler_ArrayObject || zend.Z_OBJ_HT_P(array) == &spl_handler_ArrayIterator {
			zend.ZvalPtrDtor(&intern.GetArray())
			if just_array != 0 {
				var other *SplArrayObject = Z_SPLARRAY_P(array)
				ar_flags = other.GetArFlags() & ^SPL_ARRAY_INT_MASK
			}
			if zend.Z_OBJ_P(object) == zend.Z_OBJ_P(array) {
				ar_flags |= SPL_ARRAY_IS_SELF
				zend.ZVAL_UNDEF(&intern.GetArray())
			} else {
				ar_flags |= SPL_ARRAY_USE_OTHER
				zend.ZVAL_COPY(&intern.GetArray(), array)
			}
		} else {
			var handler zend.ZendObjectGetPropertiesT = zend.Z_OBJ_HANDLER_P(array, get_properties)
			if handler != zend.ZendStdGetProperties {
				zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "Overloaded object of type %s is not compatible with %s", zend.ZSTR_VAL(zend.Z_OBJCE_P(array).GetName()), zend.ZSTR_VAL(intern.GetStd().GetCe().GetName()))
				return
			}
			zend.ZvalPtrDtor(&intern.GetArray())
			zend.ZVAL_COPY(&intern.GetArray(), array)
		}
	}
	intern.SetArFlags(intern.GetArFlags() &^ SPL_ARRAY_IS_SELF & ^SPL_ARRAY_USE_OTHER)
	intern.AddArFlags(ar_flags)
	intern.SetHtIter(uint32 - 1)
}
func SplArrayGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *zend.ZendUserIterator
	var array_object *SplArrayObject = Z_SPLARRAY_P(object)
	if by_ref != 0 && array_object.IsOverloadedCurrent() {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("zend_user_iterator"))
	zend.ZendIteratorInit(&iterator.GetIt())
	zend.Z_ADDREF_P(object)
	zend.ZVAL_OBJ(&iterator.GetIt().GetData(), zend.Z_OBJ_P(object))
	iterator.GetIt().SetFuncs(&SplArrayItFuncs)
	iterator.SetCe(ce)
	zend.ZVAL_UNDEF(&iterator.GetValue())
	return &iterator.GetIt()
}
func zim_spl_Array___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject
	var array *zend.Zval
	var ar_flags zend.ZendLong = 0
	var ce_get_iterator *zend.ZendClassEntry = spl_ce_ArrayIterator
	if zend.ZEND_NUM_ARGS() == 0 {
		return
	}
	if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "z|lC", &array, &ar_flags, &ce_get_iterator) == zend.FAILURE {
		return
	}
	intern = Z_SPLARRAY_P(object)
	if zend.ZEND_NUM_ARGS() > 2 {
		intern.SetCeGetIterator(ce_get_iterator)
	}
	ar_flags &= ^SPL_ARRAY_INT_MASK
	SplArraySetArray(object, intern, array, ar_flags, zend.ZEND_NUM_ARGS() == 1)
}
func zim_spl_ArrayIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject
	var array *zend.Zval
	var ar_flags zend.ZendLong = 0
	if zend.ZEND_NUM_ARGS() == 0 {
		return
	}
	if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "z|l", &array, &ar_flags) == zend.FAILURE {
		return
	}
	intern = Z_SPLARRAY_P(object)
	ar_flags &= ^SPL_ARRAY_INT_MASK
	SplArraySetArray(object, intern, array, ar_flags, zend.ZEND_NUM_ARGS() == 1)
}
func zim_spl_Array_setIteratorClass(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var ce_get_iterator *zend.ZendClassEntry = spl_ce_ArrayIterator
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgClass(_arg, &ce_get_iterator, _i, 0) == 0 {
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	intern.SetCeGetIterator(ce_get_iterator)
}
func zim_spl_Array_getIteratorClass(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendStringAddref(intern.GetCeGetIterator().GetName())
	zend.RETVAL_STR(intern.GetCeGetIterator().GetName())
	return
}
func zim_spl_Array_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(intern.GetArFlags() & ^SPL_ARRAY_INT_MASK)
	return
}
func zim_spl_Array_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var ar_flags zend.ZendLong = 0
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &ar_flags) == zend.FAILURE {
		return
	}
	intern.SetArFlags(intern.GetArFlags()&SPL_ARRAY_INT_MASK | ar_flags & ^SPL_ARRAY_INT_MASK)
}
func zim_spl_Array_exchangeArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var array *zend.Zval
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &array) == zend.FAILURE {
		return
	}
	if intern.GetNApplyCount() > 0 {
		zend.ZendError(zend.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return
	}
	zend.RETVAL_ARR(zend.ZendArrayDup(SplArrayGetHashTable(intern)))
	SplArraySetArray(object, intern, array, 0, 1)
}
func zim_spl_Array_getIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZVAL_OBJ(return_value, SplArrayObjectNewEx(intern.GetCeGetIterator(), object, 0))
}
func zim_spl_Array_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplArrayRewind(intern)
}
func zim_spl_Array_seek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var opos zend.ZendLong
	var position zend.ZendLong
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	var result int
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &position) == zend.FAILURE {
		return
	}
	opos = position
	if position >= 0 {
		SplArrayRewind(intern)
		result = zend.SUCCESS
		for b.PostDec(&position) > 0 && b.Assign(&result, SplArrayNext(intern)) == zend.SUCCESS {

		}
		if result == zend.SUCCESS && zend.ZendHashHasMoreElementsEx(aht, SplArrayGetPosPtr(aht, intern)) == zend.SUCCESS {
			return
		}
	}
	zend.ZendThrowExceptionEx(spl_ce_OutOfBoundsException, 0, "Seek position "+zend.ZEND_LONG_FMT+" is out of range", opos)
}
func SplArrayObjectCountElementsHelper(intern *SplArrayObject) zend.ZendLong {
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if SplArrayIsObject(intern) != 0 {
		var count zend.ZendLong = 0
		var key *zend.ZendString
		var val *zend.Zval

		/* Count public/dynamic properties */

		for {
			var __ht *zend.HashTable = aht
			var _p *zend.Bucket = __ht.GetArData()
			var _end *zend.Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.GetVal()

				if zend.Z_TYPE_P(_z) == zend.IS_UNDEF {
					continue
				}
				key = _p.GetKey()
				val = _z
				if zend.Z_TYPE_P(val) == zend.IS_INDIRECT {
					if zend.Z_TYPE_P(zend.Z_INDIRECT_P(val)) == zend.IS_UNDEF {
						continue
					}
					if key != nil && zend.ZSTR_VAL(key)[0] == '0' {
						continue
					}
				}
				count++
			}
			break
		}
		return count
	} else {
		return zend.ZendHashNumElements(aht)
	}
}
func SplArrayObjectCountElements(object *zend.Zval, count *zend.ZendLong) int {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	if intern.GetFptrCount() != nil {
		var rv zend.Zval
		zend.ZendCallMethodWith0Params(object, intern.GetStd().GetCe(), &intern.GetFptrCount(), "count", &rv)
		if rv.GetType() != zend.IS_UNDEF {
			*count = zend.ZvalGetLong(&rv)
			zend.ZvalPtrDtor(&rv)
			return zend.SUCCESS
		}
		*count = 0
		return zend.FAILURE
	}
	*count = SplArrayObjectCountElementsHelper(intern)
	return zend.SUCCESS
}
func zim_spl_Array_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(SplArrayObjectCountElementsHelper(intern))
	return
}
func SplArrayMethod(execute_data *zend.ZendExecuteData, return_value *zend.Zval, fname string, fname_len int, use_arg int) {
	var intern *SplArrayObject = Z_SPLARRAY_P(zend.ZEND_THIS)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	var function_name zend.Zval
	var params []zend.Zval
	var arg *zend.Zval = nil
	zend.ZVAL_STRINGL(&function_name, fname, fname_len)
	zend.ZVAL_NEW_EMPTY_REF(&params[0])
	zend.ZVAL_ARR(zend.Z_REFVAL(params[0]), aht)
	zend.GC_ADDREF(aht)
	if use_arg == 0 {
		intern.GetNApplyCount()++
		zend.CallUserFunction(zend.ExecutorGlobals.GetFunctionTable(), nil, &function_name, return_value, 1, params)
		intern.GetNApplyCount()--
	} else if use_arg == SPL_ARRAY_METHOD_MAY_USER_ARG {
		if zend.ZendParseParametersEx(zend.ZEND_PARSE_PARAMS_QUIET, zend.ZEND_NUM_ARGS(), "|z", &arg) == zend.FAILURE {
			zend.ZendThrowException(spl_ce_BadMethodCallException, "Function expects one argument at most", 0)
			goto exit
		}
		if arg != nil {
			zend.ZVAL_COPY_VALUE(&params[1], arg)
		}
		intern.GetNApplyCount()++
		zend.CallUserFunction(zend.ExecutorGlobals.GetFunctionTable(), nil, &function_name, return_value, b.Cond(arg != nil, 2, 1), params)
		intern.GetNApplyCount()--
	} else {
		if zend.ZEND_NUM_ARGS() != 1 || zend.ZendParseParametersEx(zend.ZEND_PARSE_PARAMS_QUIET, zend.ZEND_NUM_ARGS(), "z", &arg) == zend.FAILURE {
			zend.ZendThrowException(spl_ce_BadMethodCallException, "Function expects exactly one argument", 0)
			goto exit
		}
		zend.ZVAL_COPY_VALUE(&params[1], arg)
		intern.GetNApplyCount()++
		zend.CallUserFunction(zend.ExecutorGlobals.GetFunctionTable(), nil, &function_name, return_value, 2, params)
		intern.GetNApplyCount()--
	}
exit:
	var new_ht *zend.HashTable = zend.Z_ARRVAL_P(zend.Z_REFVAL(params[0]))
	if aht != new_ht {
		SplArrayReplaceHashTable(intern, new_ht)
	} else {
		zend.GC_DELREF(aht)
	}
	zend.ZVAL_NULL(zend.Z_REFVAL(params[0]))
	zend.ZvalPtrDtor(&params[0])
	zend.ZendStringFree(zend.Z_STR(function_name))
}
func zim_spl_Array_asort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "asort", b.SizeOf("\"asort\"")-1, SPL_ARRAY_METHOD_MAY_USER_ARG)
}
func zim_spl_Array_ksort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "ksort", b.SizeOf("\"ksort\"")-1, SPL_ARRAY_METHOD_MAY_USER_ARG)
}
func zim_spl_Array_uasort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "uasort", b.SizeOf("\"uasort\"")-1, SPL_ARRAY_METHOD_USE_ARG)
}
func zim_spl_Array_uksort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "uksort", b.SizeOf("\"uksort\"")-1, SPL_ARRAY_METHOD_USE_ARG)
}
func zim_spl_Array_natsort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "natsort", b.SizeOf("\"natsort\"")-1, SPL_ARRAY_METHOD_NO_ARG)
}
func zim_spl_Array_natcasesort(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplArrayMethod(execute_data, return_value, "natcasesort", b.SizeOf("\"natcasesort\"")-1, SPL_ARRAY_METHOD_NO_ARG)
}
func zim_spl_Array_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var entry *zend.Zval
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if b.Assign(&entry, zend.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, intern))) == nil {
		return
	}
	if zend.Z_TYPE_P(entry) == zend.IS_INDIRECT {
		entry = zend.Z_INDIRECT_P(entry)
		if zend.Z_TYPE_P(entry) == zend.IS_UNDEF {
			return
		}
	}
	zend.ZVAL_COPY_DEREF(return_value, entry)
}
func zim_spl_Array_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplArrayIteratorKey(zend.ZEND_THIS, return_value)
}
func SplArrayIteratorKey(object *zend.Zval, return_value *zend.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	zend.ZendHashGetCurrentKeyZvalEx(aht, return_value, SplArrayGetPosPtr(aht, intern))
}
func zim_spl_Array_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplArrayNextEx(intern, aht)
}
func zim_spl_Array_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_BOOL(zend.ZendHashHasMoreElementsEx(aht, SplArrayGetPosPtr(aht, intern)) == zend.SUCCESS)
	return
}
func zim_spl_Array_hasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var entry *zend.Zval
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if b.Assign(&entry, zend.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, intern))) == nil {
		zend.RETVAL_FALSE
		return
	}
	if zend.Z_TYPE_P(entry) == zend.IS_INDIRECT {
		entry = zend.Z_INDIRECT_P(entry)
	}
	zend.ZVAL_DEREF(entry)
	zend.RETVAL_BOOL(zend.Z_TYPE_P(entry) == zend.IS_ARRAY || zend.Z_TYPE_P(entry) == zend.IS_OBJECT && !intern.IsChildArraysOnly())
	return
}
func zim_spl_Array_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var entry *zend.Zval
	var flags zend.Zval
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var aht *zend.HashTable = SplArrayGetHashTable(intern)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if b.Assign(&entry, zend.ZendHashGetCurrentDataEx(aht, SplArrayGetPosPtr(aht, intern))) == nil {
		return
	}
	if zend.Z_TYPE_P(entry) == zend.IS_INDIRECT {
		entry = zend.Z_INDIRECT_P(entry)
	}
	zend.ZVAL_DEREF(entry)
	if zend.Z_TYPE_P(entry) == zend.IS_OBJECT {
		if intern.IsChildArraysOnly() {
			return
		}
		if zend.InstanceofFunction(zend.Z_OBJCE_P(entry), zend.Z_OBJCE_P(zend.ZEND_THIS)) != 0 {
			zend.ZVAL_OBJ(return_value, zend.Z_OBJ_P(entry))
			zend.Z_ADDREF_P(return_value)
			return
		}
	}
	zend.ZVAL_LONG(&flags, intern.GetArFlags())
	SplInstantiateArgEx2(zend.Z_OBJCE_P(zend.ZEND_THIS), return_value, entry, &flags)
}
func zim_spl_Array_serialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var members zend.Zval
	var flags zend.Zval
	var var_hash standard.PhpSerializeDataT
	var buf zend.SmartStr = zend.SmartStr{0}
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	standard.PHP_VAR_SERIALIZE_INIT(var_hash)
	zend.ZVAL_LONG(&flags, intern.GetArFlags()&SPL_ARRAY_CLONE_MASK)

	/* storage */

	zend.SmartStrAppendl(&buf, "x:", 2)
	standard.PhpVarSerialize(&buf, &flags, &var_hash)
	if !intern.IsIsSelf() {
		standard.PhpVarSerialize(&buf, &intern.GetArray(), &var_hash)
		zend.SmartStrAppendc(&buf, ';')
	}

	/* members */

	zend.SmartStrAppendl(&buf, "m:", 2)
	if intern.GetStd().GetProperties() == nil {
		zend.RebuildObjectProperties(&intern.GetStd())
	}
	zend.ZVAL_ARR(&members, intern.GetStd().GetProperties())
	standard.PhpVarSerialize(&buf, &members, &var_hash)

	/* done */

	standard.PHP_VAR_SERIALIZE_DESTROY(var_hash)
	if buf.GetS() != nil {
		zend.RETVAL_NEW_STR(buf.GetS())
		return
	}
	zend.RETVAL_NULL()
	return
}
func zim_spl_Array_unserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
	var buf *byte
	var buf_len int
	var p *uint8
	var s *uint8
	var var_hash standard.PhpUnserializeDataT
	var members *zend.Zval
	var zflags *zend.Zval
	var array *zend.Zval
	var flags zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "s", &buf, &buf_len) == zend.FAILURE {
		return
	}
	if buf_len == 0 {
		return
	}
	if intern.GetNApplyCount() > 0 {
		zend.ZendError(zend.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
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
	if standard.PhpVarUnserialize(zflags, &p, s+buf_len, &var_hash) == 0 || zend.Z_TYPE_P(zflags) != zend.IS_LONG {
		goto outexcept
	}
	p--
	flags = zend.Z_LVAL_P(zflags)

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
		zend.ZvalPtrDtor(&intern.GetArray())
		zend.ZVAL_UNDEF(&intern.GetArray())
	} else {
		if (*p) != 'a' && (*p) != 'O' && (*p) != 'C' && (*p) != 'r' {
			goto outexcept
		}
		array = standard.VarTmpVar(&var_hash)
		if standard.PhpVarUnserialize(array, &p, s+buf_len, &var_hash) == 0 || zend.Z_TYPE_P(array) != zend.IS_ARRAY && zend.Z_TYPE_P(array) != zend.IS_OBJECT {
			goto outexcept
		}
		intern.SetIsCloneMask(false)
		intern.AddArFlags(flags & SPL_ARRAY_CLONE_MASK)
		if zend.Z_TYPE_P(array) == zend.IS_ARRAY {
			zend.ZvalPtrDtor(&intern.GetArray())
			zend.ZVAL_COPY_VALUE(&intern.GetArray(), array)
			zend.ZVAL_NULL(array)
			zend.SEPARATE_ARRAY(&intern.GetArray())
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
	if standard.PhpVarUnserialize(members, &p, s+buf_len, &var_hash) == 0 || zend.Z_TYPE_P(members) != zend.IS_ARRAY {
		goto outexcept
	}

	/* copy members */

	zend.ObjectPropertiesLoad(&intern.GetStd(), zend.Z_ARRVAL_P(members))

	/* done reading $serialized */

	standard.PHP_VAR_UNSERIALIZE_DESTROY(var_hash)
	return
outexcept:
	standard.PHP_VAR_UNSERIALIZE_DESTROY(var_hash)
	zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Error at offset "+zend.ZEND_LONG_FMT+" of %zd bytes", zend_long((*byte)(p-buf)), buf_len)
	return
}
func zim_spl_Array___serialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(zend.ZEND_THIS)
	var tmp zend.Zval
	if zend.ZendParseParametersNoneThrow() == zend.FAILURE {
		return
	}
	zend.ArrayInit(return_value)

	/* flags */

	zend.ZVAL_LONG(&tmp, intern.GetArFlags()&SPL_ARRAY_CLONE_MASK)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &tmp)

	/* storage */

	if intern.IsIsSelf() {
		zend.ZVAL_NULL(&tmp)
	} else {
		zend.ZVAL_COPY(&tmp, &intern.GetArray())
	}
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &tmp)

	/* members */

	zend.ZVAL_ARR(&tmp, zend.ZendStdGetProperties(zend.ZEND_THIS))
	zend.Z_TRY_ADDREF(tmp)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &tmp)

	/* iterator class */

	if intern.GetCeGetIterator() == spl_ce_ArrayIterator {
		zend.ZVAL_NULL(&tmp)
	} else {
		zend.ZVAL_STR_COPY(&tmp, intern.GetCeGetIterator().GetName())
	}
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &tmp)
}
func zim_spl_Array___unserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(zend.ZEND_THIS)
	var data *zend.HashTable
	var flags_zv *zend.Zval
	var storage_zv *zend.Zval
	var members_zv *zend.Zval
	var iterator_class_zv *zend.Zval
	var flags zend.ZendLong
	if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "h", &data) == zend.FAILURE {
		return
	}
	flags_zv = zend.ZendHashIndexFind(data, 0)
	storage_zv = zend.ZendHashIndexFind(data, 1)
	members_zv = zend.ZendHashIndexFind(data, 2)
	iterator_class_zv = zend.ZendHashIndexFind(data, 3)
	if flags_zv == nil || storage_zv == nil || members_zv == nil || zend.Z_TYPE_P(flags_zv) != zend.IS_LONG || zend.Z_TYPE_P(members_zv) != zend.IS_ARRAY || iterator_class_zv != nil && (zend.Z_TYPE_P(iterator_class_zv) != zend.IS_NULL && zend.Z_TYPE_P(iterator_class_zv) != zend.IS_STRING) {
		zend.ZendThrowException(spl_ce_UnexpectedValueException, "Incomplete or ill-typed serialization data", 0)
		return
	}
	flags = zend.Z_LVAL_P(flags_zv)
	intern.SetIsCloneMask(false)
	intern.AddArFlags(flags & SPL_ARRAY_CLONE_MASK)
	if (flags & SPL_ARRAY_IS_SELF) != 0 {
		zend.ZvalPtrDtor(&intern.GetArray())
		zend.ZVAL_UNDEF(&intern.GetArray())
	} else {
		SplArraySetArray(zend.ZEND_THIS, intern, storage_zv, 0, 1)
	}
	zend.ObjectPropertiesLoad(&intern.GetStd(), zend.Z_ARRVAL_P(members_zv))
	if iterator_class_zv != nil && zend.Z_TYPE_P(iterator_class_zv) == zend.IS_STRING {
		var ce *zend.ZendClassEntry = zend.ZendLookupClass(zend.Z_STR_P(iterator_class_zv))
		if ce == nil {
			zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Cannot deserialize ArrayObject with iterator class '%s'; no such class exists", zend.ZSTR_VAL(zend.Z_STR_P(iterator_class_zv)))
			return
		} else if zend.InstanceofFunction(ce, spl_ce_Iterator) == 0 {
			zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Cannot deserialize ArrayObject with iterator class '%s'; this class does not implement the Iterator interface", zend.ZSTR_VAL(zend.Z_STR_P(iterator_class_zv)))
			return
		} else {
			intern.SetCeGetIterator(ce)
		}
	}
}
func zim_spl_Array___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_ARR(SplArrayGetDebugInfo(zend.getThis()))
	return
}
func ZmStartupSplArray(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_ArrayObject, "ArrayObject", SplArrayObjectNew, spl_funcs_ArrayObject)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, spl_ce_Aggregate)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, spl_ce_ArrayAccess)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, spl_ce_Serializable)
	zend.ZendClassImplements(spl_ce_ArrayObject, 1, spl_ce_Countable)
	memcpy(&spl_handler_ArrayObject, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	spl_handler_ArrayObject.SetOffset(zend_long((*byte)(&((*SplArrayObject)(nil).GetStd())) - (*byte)(nil)))
	spl_handler_ArrayObject.SetCloneObj(SplArrayObjectClone)
	spl_handler_ArrayObject.SetReadDimension(SplArrayReadDimension)
	spl_handler_ArrayObject.SetWriteDimension(SplArrayWriteDimension)
	spl_handler_ArrayObject.SetUnsetDimension(SplArrayUnsetDimension)
	spl_handler_ArrayObject.SetHasDimension(SplArrayHasDimension)
	spl_handler_ArrayObject.SetCountElements(SplArrayObjectCountElements)
	spl_handler_ArrayObject.SetGetPropertiesFor(SplArrayGetPropertiesFor)
	spl_handler_ArrayObject.SetGetGc(SplArrayGetGc)
	spl_handler_ArrayObject.SetReadProperty(SplArrayReadProperty)
	spl_handler_ArrayObject.SetWriteProperty(SplArrayWriteProperty)
	spl_handler_ArrayObject.SetGetPropertyPtrPtr(SplArrayGetPropertyPtrPtr)
	spl_handler_ArrayObject.SetHasProperty(SplArrayHasProperty)
	spl_handler_ArrayObject.SetUnsetProperty(SplArrayUnsetProperty)
	spl_handler_ArrayObject.SetCompareObjects(SplArrayCompareObjects)
	spl_handler_ArrayObject.SetDtorObj(zend.ZendObjectsDestroyObject)
	spl_handler_ArrayObject.SetFreeObj(SplArrayObjectFreeStorage)
	SplRegisterStdClass(&spl_ce_ArrayIterator, "ArrayIterator", SplArrayObjectNew, spl_funcs_ArrayIterator)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_Iterator)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_ArrayAccess)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_SeekableIterator)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_Serializable)
	zend.ZendClassImplements(spl_ce_ArrayIterator, 1, spl_ce_Countable)
	memcpy(&spl_handler_ArrayIterator, &spl_handler_ArrayObject, b.SizeOf("zend_object_handlers"))
	spl_ce_ArrayIterator.SetGetIterator(SplArrayGetIterator)
	spl_ce_ArrayIterator.AddCeFlags(zend.ZEND_ACC_REUSE_GET_ITERATOR)
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayObject, "STD_PROP_LIST", b.SizeOf("\"STD_PROP_LIST\"")-1, zend.ZendLong(SPL_ARRAY_STD_PROP_LIST))
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayObject, "ARRAY_AS_PROPS", b.SizeOf("\"ARRAY_AS_PROPS\"")-1, zend.ZendLong(SPL_ARRAY_ARRAY_AS_PROPS))
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayIterator, "STD_PROP_LIST", b.SizeOf("\"STD_PROP_LIST\"")-1, zend.ZendLong(SPL_ARRAY_STD_PROP_LIST))
	zend.ZendDeclareClassConstantLong(spl_ce_ArrayIterator, "ARRAY_AS_PROPS", b.SizeOf("\"ARRAY_AS_PROPS\"")-1, zend.ZendLong(SPL_ARRAY_ARRAY_AS_PROPS))
	SplRegisterSubClass(&spl_ce_RecursiveArrayIterator, spl_ce_ArrayIterator, "RecursiveArrayIterator", SplArrayObjectNew, spl_funcs_RecursiveArrayIterator)
	zend.ZendClassImplements(spl_ce_RecursiveArrayIterator, 1, spl_ce_RecursiveIterator)
	spl_ce_RecursiveArrayIterator.SetGetIterator(SplArrayGetIterator)
	spl_ce_RecursiveArrayIterator.AddCeFlags(zend.ZEND_ACC_REUSE_GET_ITERATOR)
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveArrayIterator, "CHILD_ARRAYS_ONLY", b.SizeOf("\"CHILD_ARRAYS_ONLY\"")-1, zend.ZendLong(SPL_ARRAY_CHILD_ARRAYS_ONLY))
	return zend.SUCCESS
}
