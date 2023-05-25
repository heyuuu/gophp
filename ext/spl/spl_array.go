package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

var _ types.IObject = (*ZicArrayObject)(nil)

type ZicArrayObject struct {
	types.ObjectStd

	array           types.Zval
	htIter          *types.ArrayIterator
	ht_iter         uint32
	ar_flags        int
	nApplyCount     uint8
	fptr_offset_get types.IFunction
	fptr_offset_set types.IFunction
	fptr_offset_has types.IFunction
	fptr_offset_del types.IFunction
	fptr_count      types.IFunction
	ce_get_iterator *types.ClassEntry
}

func NewZicArrayObject(ce *types.ClassEntry) *ZicArrayObject {
	// todo
	return &ZicArrayObject{}
}

func (o *ZicArrayObject) obj() *types.ZendObject {
	// todo
}
func (o *ZicArrayObject) zv() *types.Zval {
	// todo
}

func (o *ZicArrayObject) CanClone() bool { return true }
func (o *ZicArrayObject) Clone() *types.ZendObject {
	// todo
	// SplArrayObjectClone
	panic("implement me")
}

func (o *ZicArrayObject) ReadProperty(member *types.Zval, typ int, cacheSlot *any, rv *types.Zval) *types.Zval {
	// SplArrayReadProperty()
	if o.IsArrayAsProps() && o.ObjectStd.HasProperty(member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		return o.ReadDimension(member, typ, rv)
	}
	return o.ObjectStd.ReadProperty(member, typ, cacheSlot, rv)
}

func (o *ZicArrayObject) WriteProperty(member *types.Zval, value *types.Zval, cacheSlot *any) *types.Zval {
	// SplArrayWriteProperty
	if o.IsArrayAsProps() && o.ObjectStd.HasProperty(member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		o.WriteDimension(member, value)
		return value
	}
	return o.ObjectStd.WriteProperty(member, value, cacheSlot)
}

func (o *ZicArrayObject) HasProperty(member *types.Zval, hasSetExists int, cacheSlot *any) int {
	// SplArrayHasProperty
	if o.IsArrayAsProps() && o.ObjectStd.HasProperty(member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		return o.HasDimension(member, hasSetExists)
	}
	return o.ObjectStd.HasProperty(member, hasSetExists, cacheSlot)
}

func (o *ZicArrayObject) UnsetProperty(member *types.Zval, cacheSlot *any) {
	// SplArrayUnsetProperty
	if o.IsArrayAsProps() && o.ObjectStd.HasProperty(member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		o.UnsetDimension(member)
		return
	}
	o.ObjectStd.UnsetProperty(member, cacheSlot)
}

func (o *ZicArrayObject) GetPropertyPtr(member *types.Zval, typ int, cacheSlot *any) *types.Zval {
	// SplArrayGetPropertyPtrPtr

	if o.IsArrayAsProps() && o.HasProperty(member, zend.ZEND_PROPERTY_EXISTS, nil) == 0 {
		/* If object has offsetGet() overridden, then fallback to read_property,
		 * which will call offsetGet(). */
		if o.GetFptrOffsetGet() != nil {
			return nil
		}
		return o.getDimensionPtr(member, typ)
	}
	// parent::GetPropertyPtr
	return o.ObjectStd.GetPropertyPtr(member, typ, cacheSlot)
}

func (o *ZicArrayObject) CanGetPropertiesFor() bool { return true }
func (o *ZicArrayObject) GetPropertiesFor(purpose zend.ZendPropPurpose) *types.Array {
	// SplArrayGetPropertiesFor
	if o.IsStdPropList() {
		//return zend.ZendStdGetPropertiesFor(object, purpose)
		return o.ObjectStd.GetPropertiesArray()
	}

	/* We are supposed to be the only owner of the internal hashtable.
	 * The "dup" flag decides whether this is a "long-term" use where
	 * we need to duplicate, or a "temporary" one, where we can expect
	 * that no operations on the ArrayObject will be performed in the
	 * meantime. */
	var dup bool
	switch purpose {
	case zend.ZEND_PROP_PURPOSE_ARRAY_CAST:
		dup = true
	case zend.ZEND_PROP_PURPOSE_VAR_EXPORT,
		zend.ZEND_PROP_PURPOSE_JSON,
		zend.ZEND_PROP_PURPOSE_ARRAY_KEY_EXISTS:
		dup = false
	default:
		//return zend.ZendStdGetPropertiesFor(object, purpose)
		return o.ObjectStd.GetPropertiesArray()
	}
	ht := o.array.Array()
	if dup {
		ht = types.ZendArrayDup(ht)
	}
	return ht
}

func (o *ZicArrayObject) ReadDimension(offset *types.Zval, typ int, rv *types.Zval) *types.Zval {
	return o.readDimensionEx(true, offset, typ, rv)
}

func (o *ZicArrayObject) readDimensionEx(checkInherited bool, offset *types.Zval, typ int, rv *types.Zval) *types.Zval {
	// SplArrayReadDimensionEx
	if checkInherited && (o.fptr_offset_get != nil || typ == zend.BP_VAR_IS && o.fptr_offset_has != nil) {
		if typ == zend.BP_VAR_IS {
			if o.HasDimension(offset, 0) == 0 {
				return zend.UninitializedZval()
			}
		}
		if o.fptr_offset_get != nil {
			if offset == nil {
				offset = types.NewZvalUndef()
			} else {
				offset = types.SEPARATE_ARG_IF_REF(offset)
			}
			zend.ZendCallMethodWith1Params(o.zv(), o.GetCe(), o.fptr_offset_get, "offsetGet", rv, offset)
			if rv.IsNotUndef() {
				return rv
			}
			return zend.UninitializedZval()
		}
	}
	ret := o.getDimensionPtr(offset, typ)

	/* When in a write context,
	 * ZE has to be fooled into thinking this is in a reference set
	 * by separating (if necessary) and returning as IS_REFERENCE (with refcount == 1)
	 */
	if (typ == zend.BP_VAR_W || typ == zend.BP_VAR_RW || typ == zend.BP_VAR_UNSET) && !(ret.IsReference()) && ret != zend.UninitializedZval() {
		ret.SetNewRef(ret)
	}
	return ret
}

func (o *ZicArrayObject) WriteDimension(offset *types.Zval, value *types.Zval) {
	// SplArrayWriteDimension
	o.writeDimensionEx(true, offset, value)
}

func (o *ZicArrayObject) writeDimensionEx(checkInherited bool, offset *types.Zval, value *types.Zval) {
	var index zend.ZendLong
	var ht *types.Array
	if checkInherited && o.GetFptrOffsetSet() != nil {
		var tmp types.Zval
		if offset == nil {
			tmp.SetNull()
			offset = &tmp
		} else {
			offset = types.SEPARATE_ARG_IF_REF(offset)
		}
		zend.ZendCallMethodWith2Params(o.zv(), o.GetCe(), o.GetFptrOffsetSet(), "offsetSet", nil, offset, value)
		return
	}
	if o.GetNApplyCount() > 0 {
		faults.Error(faults.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return
	}
	if offset == nil {
		ht = o.array.Array()
		ht.Append(value)
		return
	}

	// 解析 offset 为 string or int
	isStrKey, isAppend, index, strKey := false, false, 0, ""
	offset = offset.DeRef()
	switch offset.GetType() {
	case types.IS_STRING:
		isStrKey = true
		strKey = offset.StringVal()
	case types.IS_DOUBLE:
		index = zend.ZendLong(offset.Double())
	case types.IS_RESOURCE:
		index = offset.ResourceHandle()
	case types.IS_FALSE:
		index = 0
	case types.IS_TRUE:
		index = 1
	case types.IS_LONG:
		index = offset.Long()
	case types.IS_NULL:
		isAppend = true
	default:
		faults.Error(faults.E_WARNING, "Illegal offset type")
		// zend.ZvalPtrDtor(value)
		return
	}

	// 更新
	ht = o.array.Array()
	if isAppend {
		ht.Append(value)
	} else if isStrKey {
		ht.SymtableUpdateInd(strKey, value)
	} else {
		ht.IndexUpdate(index, value)
	}
}

func (o *ZicArrayObject) HasDimension(offset *types.Zval, checkEmpty int) int {
	// SplArrayHasDimension
	return o.hasDimensionEx(true, offset, checkEmpty)
}

func (o *ZicArrayObject) hasDimensionEx(checkInherited bool, offset *types.Zval, checkEmpty int) int {
	// SplArrayHasDimensionEx
	var rv types.Zval
	var value *types.Zval = nil
	if checkInherited && o.GetFptrOffsetHas() != nil {
		offset = types.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(o.zv(), o.GetCe(), o.GetFptrOffsetHas(), "offsetExists", &rv, offset)
		if operators.ZvalIsTrue(&rv) {
			if checkEmpty != 1 {
				return 1
			} else if o.GetFptrOffsetGet() != nil {
				value = o.readDimensionEx(true, offset, zend.BP_VAR_R, &rv)
			}
		} else {
			return 0
		}
	}
	if value == nil {
		// 解析 offset 为 string or int
		isStrKey, index, strKey := false, 0, ""

		offset = offset.DeRef()
		switch offset.GetType() {
		case types.IS_STRING:
			isStrKey = true
			strKey = offset.StringVal()
		case types.IS_DOUBLE:
			index = zend.ZendLong(offset.Double())
		case types.IS_RESOURCE:
			index = offset.ResourceHandle()
		case types.IS_FALSE:
			index = 0
		case types.IS_TRUE:
			index = 1
		case types.IS_LONG:
			index = offset.Long()
		default:
			faults.Error(faults.E_WARNING, "Illegal offset type")
			return 0
		}

		var ht = o.array.Array()
		var tmp *types.Zval
		if isStrKey {
			tmp = ht.SymtableFind(strKey)
		} else {
			tmp = ht.IndexFind(index)
		}

		if tmp != nil {
			if checkEmpty == 2 {
				return 1
			}
		} else {
			return 0
		}

		if checkEmpty != 0 && checkInherited && o.GetFptrOffsetGet() != nil {
			value = o.readDimensionEx(true, offset, zend.BP_VAR_R, &rv)
		} else {
			value = tmp
		}
	}

	var result bool
	if checkEmpty != 0 {
		result = operators.ZvalIsTrue(value)
	} else {
		result = value.IsNull()
	}
	return types.IntBool(result)
}

func (o *ZicArrayObject) UnsetDimension(offset *types.Zval) {
	// SplArrayUnsetDimension
	o.unsetDimensionEx(true, offset)
}

func (o *ZicArrayObject) unsetDimensionEx(checkInherited bool, offset *types.Zval) {
	if checkInherited && o.GetFptrOffsetDel() != nil {
		offset = types.SEPARATE_ARG_IF_REF(offset)
		zend.ZendCallMethodWith1Params(o.zv(), o.GetCe(), o.GetFptrOffsetDel(), "offsetUnset", nil, offset)
		return
	}
	if o.GetNApplyCount() > 0 {
		faults.Error(faults.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return
	}

	// 解析 offset 为 string or int
	isStrKey, index, strKey := false, 0, ""
	offset = offset.DeRef()
	switch offset.GetType() {
	case types.IS_STRING:
		isStrKey = true
		strKey = offset.StringVal()
	case types.IS_DOUBLE:
		index = zend.ZendLong(offset.Double())
	case types.IS_RESOURCE:
		index = offset.ResourceHandle()
	case types.IS_FALSE:
		index = 0
	case types.IS_TRUE:
		index = 1
	case types.IS_LONG:
		index = offset.Long()
	default:
		faults.Error(faults.E_WARNING, "Illegal offset type")
		return
	}

	ht := o.array.Array()
	if isStrKey {
		if ht == zend.EG__().GetSymbolTable() {
			if !zend.ZendDeleteGlobalVariableEx(strKey) {
				faults.Error(faults.E_NOTICE, "Undefined index: %s", strKey)
			}
		} else {
			numericKey := types.NumericKey(strKey)
			var data = ht.Find(numericKey)
			if data != nil {
				if data.IsIndirect() {
					data = data.Indirect()
					if data.IsUndef() {
						faults.Error(faults.E_NOTICE, "Undefined index: %s", strKey)
					} else {
						data.SetUndef()
						ht.MarkHasEmptyIndex()
						types.ZendHashMoveForwardEx(ht, o.getPosPtr(ht))
						if o.isObject() {
							o.skipProtected(ht)
						}
					}
				} else if ht.SymtableDel(strKey) == false {
					faults.Error(faults.E_NOTICE, "Undefined index: %s", strKey)
				}
			} else {
				faults.Error(faults.E_NOTICE, "Undefined index: %s", strKey)
			}
		}
	} else {
		if !ht.IndexDelete(index) {
			faults.Error(faults.E_NOTICE, "Undefined offset: "+zend.ZEND_LONG_FMT, index)
		}
	}
}

func (o *ZicArrayObject) CanCountElements() bool { return true }
func (o *ZicArrayObject) CountElements() (int, bool) {
	// SplArrayObjectCountElements
	if o.GetFptrCount() != nil {
		var rv types.Zval
		zend.ZendCallMethodWith0Params(o.zv(), o.GetCe(), o.GetFptrCount(), "count", &rv)
		if rv.IsNotUndef() {
			return operators.ZvalGetLong(&rv), true
		}
		return 0, false
	}
	return o.countElementsHelper(), true
}

func (o *ZicArrayObject) countElementsHelper() int {
	// SplArrayObjectCountElementsHelper
	var aht = o.array.Array()
	if o.isObject() {
		var count = 0
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

func (o *ZicArrayObject) CanCompareObjectsTo(another *types.ZendObject) bool {
	_, ok := another.GetData().(*ZicArrayObject)
	return ok
}

func (o *ZicArrayObject) CompareObjectsTo(another *types.ZendObject) int {
	// SplArrayCompareObjects
	anotherData, ok := another.GetData().(*ZicArrayObject)
	if !ok {
		return 0
	}
	ht1 := o.array.Array()
	ht2 := anotherData.array.Array()
	result := operators.ZendCompareSymbolTables(ht1, ht2)

	if result == 0 && !(ht1 == o.ObjectStd.GetProperties() && ht2 == anotherData.ObjectStd.GetProperties()) {
		// parent::CompareObjectsTo()
		result = o.ObjectStd.CompareObjectsTo(another)
	}
	return result
}

func (o *ZicArrayObject) getDimensionPtr(offset *types.Zval, typ int) *types.Zval {
	var retval *types.Zval
	var ht = o.array.Array()
	if offset == nil || offset.IsUndef() || ht == nil {
		return zend.UninitializedZval()
	}
	if (typ == zend.BP_VAR_W || typ == zend.BP_VAR_RW) && o.nApplyCount > 0 {
		faults.Error(faults.E_WARNING, "Modification of ArrayObject during sorting is prohibited")
		return zend.EG__().GetErrorZval()
	}

	// 解析 offset 为 string or int
	isStrKey, index, strKey := false, 0, ""
	offset = offset.DeRef()
	switch offset.GetType() {
	case types.IS_NULL:
		isStrKey = true
		strKey = ""
	case types.IS_STRING:
		isStrKey = true
		strKey = offset.StringVal()
	case types.IS_RESOURCE:
		faults.Error(faults.E_NOTICE, "Resource ID#%d used as offset, casting to integer (%d)", offset.ResourceHandle(), offset.ResourceHandle())
		index = offset.ResourceHandle()
	case types.IS_DOUBLE:
		index = zend.ZendLong(offset.Double())
	case types.IS_FALSE:
		index = 0
	case types.IS_TRUE:
		index = 1
	case types.IS_LONG:
		index = offset.Long()
	default:
		faults.Error(faults.E_WARNING, "Illegal offset type")
		if typ == zend.BP_VAR_W || typ == zend.BP_VAR_RW {
			return zend.EG__().GetErrorZval()
		} else {
			return zend.UninitializedZval()
		}
	}

	if isStrKey {
		retval = ht.SymtableFind(strKey)
		if retval != nil {
			if retval.IsIndirect() {
				retval = retval.Indirect()
				if retval.IsUndef() {
					switch typ {
					case zend.BP_VAR_R:
						faults.Error(faults.E_NOTICE, "Undefined index: %s", strKey)
						fallthrough
					case zend.BP_VAR_UNSET:
						fallthrough
					case zend.BP_VAR_IS:
						retval = zend.UninitializedZval()
					case zend.BP_VAR_RW:
						faults.Error(faults.E_NOTICE, "Undefined index: %s", strKey)
						fallthrough
					case zend.BP_VAR_W:
						retval.SetNull()
					}
				}
			}
		} else {
			switch typ {
			case zend.BP_VAR_R:
				faults.Error(faults.E_NOTICE, "Undefined index: %s", strKey)
				fallthrough
			case zend.BP_VAR_UNSET:
				fallthrough
			case zend.BP_VAR_IS:
				retval = zend.UninitializedZval()
			case zend.BP_VAR_RW:
				faults.Error(faults.E_NOTICE, "Undefined index: %s", strKey)
				fallthrough
			case zend.BP_VAR_W:
				var value types.Zval
				value.SetNull()
				retval = ht.SymtableUpdate(strKey, &value)
			}
		}
		return retval
	} else {
		if b.Assign(&retval, ht.IndexFind(index)) == nil {
			switch typ {
			case zend.BP_VAR_R:
				faults.Error(faults.E_NOTICE, "Undefined offset: "+zend.ZEND_LONG_FMT, index)
				fallthrough
			case zend.BP_VAR_UNSET:
				fallthrough
			case zend.BP_VAR_IS:
				retval = zend.UninitializedZval()
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
	}
}

func (o *ZicArrayObject) isObject() bool {
	// SplArrayIsObject
	for o.IsUseOther() {
		o = Z_SPLARRAY_P(o.array) // ??
	}
	return o.IsIsSelf() || o.array.IsObject()
}

func (o *ZicArrayObject) skipProtected(aht *types.Array) int {
	// SplArraySkipProtected
	if o.isObject() {
		var posPtr = o.getPosPtr(aht)
		for {
			pair, nextPos := aht.NextEx(*posPtr)
			if pair == nil || !pair.GetKey().IsStrKey() {
				return types.SUCCESS
			}
			strKey := pair.GetKey().StrKey()
			if strKey == "" {
				return types.SUCCESS
			}

			*posPtr = nextPos
		}
	}
	return types.FAILURE
}

func (o *ZicArrayObject) getPosPtr(ht *types.Array) *uint32 {
	// SplArrayGetPosPtr
	if o.htIter == nil {
		o.createHtIter(ht)
	}
	return o.htIter.GetPos()
}

func (o *ZicArrayObject) createHtIter(ht *types.Array) {
	// SplArrayCreateHtIter
	o.htIter = ht.Iterator()
	o.skipProtected(ht)
}

//
func (o *ZicArrayObject) GetArray() *types.Zval                  { return &o.array }
func (o *ZicArrayObject) GetHtIter() uint32                      { return o.ht_iter }
func (o *ZicArrayObject) SetHtIter(value uint32)                 { o.ht_iter = value }
func (o *ZicArrayObject) GetArFlags() int                        { return o.ar_flags }
func (o *ZicArrayObject) SetArFlags(value int)                   { o.ar_flags = value }
func (o *ZicArrayObject) GetNApplyCount() uint8                  { return o.nApplyCount }
func (o *ZicArrayObject) GetFptrOffsetGet() types.IFunction      { return o.fptr_offset_get }
func (o *ZicArrayObject) SetFptrOffsetGet(value types.IFunction) { o.fptr_offset_get = value }
func (o *ZicArrayObject) GetFptrOffsetSet() types.IFunction      { return o.fptr_offset_set }
func (o *ZicArrayObject) SetFptrOffsetSet(value types.IFunction) { o.fptr_offset_set = value }
func (o *ZicArrayObject) GetFptrOffsetHas() types.IFunction      { return o.fptr_offset_has }
func (o *ZicArrayObject) SetFptrOffsetHas(value types.IFunction) { o.fptr_offset_has = value }
func (o *ZicArrayObject) GetFptrOffsetDel() types.IFunction      { return o.fptr_offset_del }
func (o *ZicArrayObject) SetFptrOffsetDel(value types.IFunction) { o.fptr_offset_del = value }
func (o *ZicArrayObject) GetFptrCount() types.IFunction          { return o.fptr_count }
func (o *ZicArrayObject) SetFptrCount(value types.IFunction)     { o.fptr_count = value }
func (o *ZicArrayObject) GetCeGetIterator() *types.ClassEntry    { return o.ce_get_iterator }
func (o *ZicArrayObject) SetCeGetIterator(value *types.ClassEntry) {
	o.ce_get_iterator = value
}

/* SplArrayObject.ar_flags */
func (o *ZicArrayObject) AddArFlags(value int)      { o.ar_flags |= value }
func (o *ZicArrayObject) SubArFlags(value int)      { o.ar_flags &^= value }
func (o *ZicArrayObject) HasArFlags(value int) bool { return o.ar_flags&value != 0 }
func (o *ZicArrayObject) SwitchArFlags(value int, cond bool) {
	if cond {
		o.AddArFlags(value)
	} else {
		o.SubArFlags(value)
	}
}
func (o *ZicArrayObject) IsIsSelf() bool       { return o.HasArFlags(SPL_ARRAY_IS_SELF) }
func (o *ZicArrayObject) IsUseOther() bool     { return o.HasArFlags(SPL_ARRAY_USE_OTHER) }
func (o *ZicArrayObject) IsStdPropList() bool  { return o.HasArFlags(SPL_ARRAY_STD_PROP_LIST) }
func (o *ZicArrayObject) IsArrayAsProps() bool { return o.HasArFlags(SPL_ARRAY_ARRAY_AS_PROPS) }
func (o *ZicArrayObject) IsOverloadedValid() bool {
	return o.HasArFlags(SPL_ARRAY_OVERLOADED_VALID)
}
func (o *ZicArrayObject) IsOverloadedCurrent() bool {
	return o.HasArFlags(SPL_ARRAY_OVERLOADED_CURRENT)
}
func (o *ZicArrayObject) IsOverloadedKey() bool  { return o.HasArFlags(SPL_ARRAY_OVERLOADED_KEY) }
func (o *ZicArrayObject) IsOverloadedNext() bool { return o.HasArFlags(SPL_ARRAY_OVERLOADED_NEXT) }
func (o *ZicArrayObject) IsOverloadedRewind() bool {
	return o.HasArFlags(SPL_ARRAY_OVERLOADED_REWIND)
}
func (o *ZicArrayObject) IsChildArraysOnly() bool {
	return o.HasArFlags(SPL_ARRAY_CHILD_ARRAYS_ONLY)
}
func (o *ZicArrayObject) IsCloneMask() bool       { return o.HasArFlags(SPL_ARRAY_CLONE_MASK) }
func (o *ZicArrayObject) SetIsIsSelf(cond bool)   { o.SwitchArFlags(SPL_ARRAY_IS_SELF, cond) }
func (o *ZicArrayObject) SetIsUseOther(cond bool) { o.SwitchArFlags(SPL_ARRAY_USE_OTHER, cond) }
func (o *ZicArrayObject) SetIsStdPropList(cond bool) {
	o.SwitchArFlags(SPL_ARRAY_STD_PROP_LIST, cond)
}
func (o *ZicArrayObject) SetIsArrayAsProps(cond bool) {
	o.SwitchArFlags(SPL_ARRAY_ARRAY_AS_PROPS, cond)
}
func (o *ZicArrayObject) SetIsOverloadedValid(cond bool) {
	o.SwitchArFlags(SPL_ARRAY_OVERLOADED_VALID, cond)
}
func (o *ZicArrayObject) SetIsOverloadedCurrent(cond bool) {
	o.SwitchArFlags(SPL_ARRAY_OVERLOADED_CURRENT, cond)
}
func (o *ZicArrayObject) SetIsOverloadedKey(cond bool) {
	o.SwitchArFlags(SPL_ARRAY_OVERLOADED_KEY, cond)
}
func (o *ZicArrayObject) SetIsOverloadedNext(cond bool) {
	o.SwitchArFlags(SPL_ARRAY_OVERLOADED_NEXT, cond)
}
func (o *ZicArrayObject) SetIsOverloadedRewind(cond bool) {
	o.SwitchArFlags(SPL_ARRAY_OVERLOADED_REWIND, cond)
}
func (o *ZicArrayObject) SetIsChildArraysOnly(cond bool) {
	o.SwitchArFlags(SPL_ARRAY_CHILD_ARRAYS_ONLY, cond)
}
func (o *ZicArrayObject) SetIsCloneMask(cond bool) { o.SwitchArFlags(SPL_ARRAY_CLONE_MASK, cond) }
