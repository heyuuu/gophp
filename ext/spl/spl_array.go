package spl

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

var _ types.IObject = (*ZicArrayObject)(nil)

type ZicArrayObject struct {
	types.ObjectStd

	array           types.Zval
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

func (o *ZicArrayObject) Free() {
	if o.ht_iter != -1 {
		zend.EG__().DelArrayIterator(o.ht_iter)
	}
	// parent::Free()
	o.ObjectStd.Free()
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
	// todo
	// SplArrayGetPropertyPtrPtr
	panic("implement me")
}

func (o *ZicArrayObject) CanGetPropertiesFor() bool { return true }
func (o *ZicArrayObject) GetPropertiesFor(purpose zend.ZendPropPurpose) *types.Array {
	// todo
	// SplArrayGetPropertiesFor
	panic("implement me")
}

func (o *ZicArrayObject) ReadDimension(offset *types.Zval, typ int, rv *types.Zval) *types.Zval {
	return o.readDimensionEx(true, offset, typ, rv)
}

func (o *ZicArrayObject) WriteDimension(offset *types.Zval, value *types.Zval) {
	// SplArrayWriteDimension
	o.writeDimensionEx(true, offset, value)
}

func (o *ZicArrayObject) HasDimension(offset *types.Zval, checkEmpty int) int {
	// todo
	// SplArrayHasDimension
	panic("implement me")
}

func (o *ZicArrayObject) UnsetDimension(offset *types.Zval) {
	// todo
	// SplArrayUnsetDimension
	panic("implement me")
}

func (o *ZicArrayObject) CanCountElements() bool { return true }
func (o *ZicArrayObject) CountElements(count *int) int {
	// todo
	// SplArrayObjectCountElements
	panic("implement me")
}

func (o *ZicArrayObject) CanCompareObjectsTo(obj2 *types.ZendObject) bool {
	// todo
	panic("implement me")
}

func (o *ZicArrayObject) CompareObjectsTo(another *types.ZendObject) int {
	// todo
	// SplArrayCompareObjects
	panic("implement me")
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

func (o *ZicArrayObject) getDimensionPtr(offset *types.Zval, typ int) *types.Zval {
	var retval *types.Zval
	var ht *types.Array = o.array.Array()
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

func (o *ZicArrayObject) unsetDimensionEx(checkInherited bool, object *types.Zval, offset *types.Zval) {
	var intern *SplArrayObject = Z_SPLARRAY_P(object)
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
						types.ZendHashMoveForwardEx(ht, SplArrayGetPosPtr(ht, intern))
						if o.isObject() {
							SplArraySkipProtected(intern, ht)
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

func (o *ZicArrayObject) isObject() bool {
	// SplArrayIsObject
	for o.IsUseOther() {
		o = Z_SPLARRAY_P(o.array) // ??
	}
	return o.IsIsSelf() || o.array.IsObject()
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
