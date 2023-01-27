// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func (this *HashTable) Flags() uint32 { return this.GetUFlags() }
func (this *HashTable) Invalidate()   { this.SetUFlags(HASH_FLAG_UNINITIALIZED) }
func (this *HashTable) IsInitialized() bool {
	return !this.HasUFlags(HASH_FLAG_UNINITIALIZED)
}
func (this *HashTable) IsPacked() bool { return this.HasUFlags(HASH_FLAG_PACKED) }
func (this *HashTable) IsWithoutHoles() bool {
	return this.GetNNumUsed() == this.GetNNumOfElements()
}
func (this *HashTable) HasStaticKeysOnly() bool {
	return this.HasUFlags(HASH_FLAG_PACKED | HASH_FLAG_STATIC_KEYS)
}
func (this *HashTable) IteratorsCount() ZendUchar   { return this.GetNIteratorsCount() }
func (this *HashTable) IteratorsOverflow() bool     { return this.GetNIteratorsCount() == 0xff }
func (this *HashTable) HasIterators() bool          { return this.GetNIteratorsCount() != 0 }
func (this *HashTable) SetIteratorsCount(iters int) { this.SetNIteratorsCount(iters) }
func (this *HashTable) IncIteratorsCount() {
	this.SetIteratorsCount(this.GetNIteratorsCount() + 1)
}
func (this *HashTable) DecIteratorsCount() {
	this.SetIteratorsCount(this.GetNIteratorsCount() - 1)
}
func ZVAL_EMPTY_ARRAY(z *Zval) {
	var __z *Zval = z
	__z.SetArr((*ZendArray)(&ZendEmptyArray))
	__z.SetTypeInfo(IS_ARRAY)
}
func (this *HashTable) Init(nSize uint32, pHashFunction __auto__, pDestructor DtorFuncT, persistent ZendBool) {
	this._init(nSize, pDestructor, persistent)
}
func (this *HashTable) InitEx(nSize uint32, pHashFunction __auto__, pDestructor DtorFuncT, persistent ZendBool, bApplyProtection int) {
	this._init(nSize, pDestructor, persistent)
}
func (this *HashTable) FindEx(key *ZendString, known_hash ZendBool) *Zval {
	if known_hash != 0 {
		return this._findKnownHash(key)
	} else {
		return this.Find(key)
	}
}
func ZEND_HASH_INDEX_FIND(_ht *HashTable, _h ZendUlong, _ret *Zval, _not_found __auto__) {
	if _ht.HasUFlags(HASH_FLAG_PACKED) {
		if zend_ulong(_h) < zend_ulong(_ht).nNumUsed {
			_ret = _ht.GetArData()[_h].GetVal()
			if _ret.IsType(IS_UNDEF) {
				goto _not_found
			}
		} else {
			goto _not_found
		}
	} else {
		_ret = _ht._indexFind(_h)
		if _ret == nil {
			goto _not_found
		}
	}
}
func (this *HashTable) Exists(key *ZendString) ZendBool        { return this.Find(key) != nil }
func (this *HashTable) StrExists(str *byte, len_ int) ZendBool { return this.StrFind(str, len_) != nil }
func (this *HashTable) IndexExists(h ZendUlong) ZendBool       { return this.IndexFind(h) != nil }
func (this *HashTable) HasMoreElementsEx(pos *HashPosition) ZEND_RESULT_CODE {
	if this.GetCurrentKeyTypeEx(pos) == HASH_KEY_NON_EXISTENT {
		return FAILURE
	} else {
		return SUCCESS
	}
}
func (this *HashTable) HasMoreElements() ZEND_RESULT_CODE {
	return this.HasMoreElementsEx(this.GetNInternalPointer())
}
func (this *HashTable) MoveForward() int {
	return this.MoveForwardEx(this.GetNInternalPointer())
}
func (this *HashTable) MoveBackwards() int {
	return this.MoveBackwardsEx(this.GetNInternalPointer())
}
func (this *HashTable) GetCurrentKey(str_index **ZendString, num_index *ZendUlong) int {
	return this.GetCurrentKeyEx(str_index, num_index, this.GetNInternalPointer())
}
func (this *HashTable) GetCurrentKeyZval(key *Zval) {
	this.GetCurrentKeyZvalEx(key, this.GetNInternalPointer())
}
func (this *HashTable) GetCurrentKeyType() int {
	return this.GetCurrentKeyTypeEx(this.GetNInternalPointer())
}
func (this *HashTable) GetCurrentData() *Zval {
	return this.GetCurrentDataEx(this.GetNInternalPointer())
}
func (this *HashTable) InternalPointerReset() {
	this.InternalPointerResetEx(this.GetNInternalPointer())
}
func (this *HashTable) InternalPointerEnd() {
	this.InternalPointerEndEx(this.GetNInternalPointer())
}
func (this *HashTable) Sort(compare_func CompareFuncT, renumber ZendBool) int {
	return this.SortEx(ZendSort, compare_func, renumber)
}
func (this *HashTable) NumElements() __auto__     { return this.GetNNumOfElements() }
func (this *HashTable) NextFreeElement() ZendLong { return this.GetNNextFreeElement() }
func ZendNewArray(size uint32) *HashTable         { return _zendNewArray(size) }
func (this *HashTable) IteratorsUpdate(from HashPosition, to HashPosition) {
	if this.HasIterators() {
		this._iteratorsUpdate(from, to)
	}
}
func ZEND_INIT_SYMTABLE(ht *HashTable) { ZEND_INIT_SYMTABLE_EX(ht, 8, 0) }
func ZEND_INIT_SYMTABLE_EX(ht *HashTable, n uint32, persistent ZendBool) {
	ht.Init(n, nil, ZVAL_PTR_DTOR, persistent)
}
func _zendHandleNumericStr(key *byte, length int, idx *ZendUlong) int {
	var tmp *byte = key
	if (*tmp) > '9' {
		return 0
	} else if (*tmp) < '0' {
		if (*tmp) != '-' {
			return 0
		}
		tmp++
		if (*tmp) > '9' || (*tmp) < '0' {
			return 0
		}
	}
	return _zendHandleNumericStrEx(key, length, idx)
}
func ZEND_HANDLE_NUMERIC_STR(key *byte, length int, idx ZendUlong) int {
	return _zendHandleNumericStr(key, length, &idx)
}
func ZEND_HANDLE_NUMERIC(key *ZendString, idx ZendUlong) int {
	return ZEND_HANDLE_NUMERIC_STR(key.GetVal(), key.GetLen(), idx)
}
func (this *HashTable) FindInd(key *ZendString) *Zval {
	var zv *Zval
	zv = this.Find(key)
	if zv != nil && zv.IsType(IS_INDIRECT) {
		if Z_INDIRECT_P(zv).GetType() != IS_UNDEF {
			return zv.GetZv()
		} else {
			return nil
		}
	} else {
		return zv
	}
}
func (this *HashTable) FindExInd(key *ZendString, known_hash ZendBool) *Zval {
	var zv *Zval
	zv = this.FindEx(key, known_hash)
	if zv != nil && zv.IsType(IS_INDIRECT) {
		if Z_INDIRECT_P(zv).GetType() != IS_UNDEF {
			return zv.GetZv()
		} else {
			return nil
		}
	} else {
		return zv
	}
}
func (this *HashTable) ExistsInd(key *ZendString) int {
	var zv *Zval
	zv = this.Find(key)
	return zv != nil && (zv.GetType() != IS_INDIRECT || Z_INDIRECT_P(zv).GetType() != IS_UNDEF)
}
func (this *HashTable) StrFindInd(str *byte, len_ int) *Zval {
	var zv *Zval
	zv = this.StrFind(str, len_)
	if zv != nil && zv.IsType(IS_INDIRECT) {
		if Z_INDIRECT_P(zv).GetType() != IS_UNDEF {
			return zv.GetZv()
		} else {
			return nil
		}
	} else {
		return zv
	}
}
func (this *HashTable) StrExistsInd(str string, len_ int) int {
	var zv *Zval
	zv = this.StrFind(str, len_)
	return zv != nil && (zv.GetType() != IS_INDIRECT || Z_INDIRECT_P(zv).GetType() != IS_UNDEF)
}
func ZendSymtableAddNew(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ht.IndexAddNew(idx, pData)
	} else {
		return ht.AddNew(key, pData)
	}
}
func ZendSymtableUpdate(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ht.IndexUpdate(idx, pData)
	} else {
		return ht.Update(key, pData)
	}
}
func ZendSymtableUpdateInd(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ht.IndexUpdate(idx, pData)
	} else {
		return ht.UpdateInd(key, pData)
	}
}
func ZendSymtableDel(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ht.IndexDel(idx)
	} else {
		return ht.Del(key)
	}
}
func ZendSymtableDelInd(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ht.IndexDel(idx)
	} else {
		return ht.DelInd(key)
	}
}
func ZendSymtableFind(ht *HashTable, key *ZendString) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ht.IndexFind(idx)
	} else {
		return ht.Find(key)
	}
}
func ZendSymtableFindInd(ht *HashTable, key *ZendString) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ht.IndexFind(idx)
	} else {
		return ht.FindInd(key)
	}
}
func ZendSymtableExists(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ht.IndexExists(idx)
	} else {
		return ht.Exists(key)
	}
}
func ZendSymtableExistsInd(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ht.IndexExists(idx)
	} else {
		return ht.ExistsInd(key)
	}
}
func ZendSymtableStrUpdate(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ht.IndexUpdate(idx, pData)
	} else {
		return ht.StrUpdate(str, len_, pData)
	}
}
func ZendSymtableStrUpdateInd(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ht.IndexUpdate(idx, pData)
	} else {
		return ht.StrUpdateInd(str, len_, pData)
	}
}
func ZendSymtableStrDel(ht *HashTable, str *byte, len_ int) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ht.IndexDel(idx)
	} else {
		return ht.StrDel(str, len_)
	}
}
func ZendSymtableStrDelInd(ht *HashTable, str *byte, len_ int) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ht.IndexDel(idx)
	} else {
		return ht.StrDelInd(str, len_)
	}
}
func ZendSymtableStrFind(ht *HashTable, str *byte, len_ int) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ht.IndexFind(idx)
	} else {
		return ht.StrFind(str, len_)
	}
}
func ZendSymtableStrExists(ht *HashTable, str *byte, len_ int) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ht.IndexExists(idx)
	} else {
		return ht.StrExists(str, len_)
	}
}
func (this *HashTable) AddPtr(key *ZendString, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = this.Add(key, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) AddNewPtr(key *ZendString, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = this.AddNew(key, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) StrAddPtr(str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = this.StrAdd(str, len_, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) StrAddNewPtr(str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = this.StrAddNew(str, len_, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) UpdatePtr(key *ZendString, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = this.Update(key, &tmp)
	return zv.GetPtr()
}
func (this *HashTable) StrUpdatePtr(str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = this.StrUpdate(str, len_, &tmp)
	return zv.GetPtr()
}
func (this *HashTable) AddMem(key *ZendString, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, this.Add(key, &tmp)) {
		zv.SetPtr(Pemalloc(size, GC_FLAGS(this)&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func (this *HashTable) AddNewMem(key *ZendString, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, this.AddNew(key, &tmp)) {
		zv.SetPtr(Pemalloc(size, GC_FLAGS(this)&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func (this *HashTable) StrAddMem(str *byte, len_ int, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, this.StrAdd(str, len_, &tmp)) {
		zv.SetPtr(Pemalloc(size, GC_FLAGS(this)&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func (this *HashTable) StrAddNewMem(str *byte, len_ int, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, this.StrAddNew(str, len_, &tmp)) {
		zv.SetPtr(Pemalloc(size, GC_FLAGS(this)&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func (this *HashTable) UpdateMem(key *ZendString, pData any, size int) any {
	var p any
	p = Pemalloc(size, GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return this.UpdatePtr(key, p)
}
func (this *HashTable) StrUpdateMem(str *byte, len_ int, pData any, size int) any {
	var p any
	p = Pemalloc(size, GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return this.StrUpdatePtr(str, len_, p)
}
func (this *HashTable) IndexAddPtr(h ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = this.IndexAdd(h, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) IndexAddNewPtr(h ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = this.IndexAddNew(h, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) IndexUpdatePtr(h ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = this.IndexUpdate(h, &tmp)
	return zv.GetPtr()
}
func (this *HashTable) IndexAddMem(h ZendUlong, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, this.IndexAdd(h, &tmp)) {
		zv.SetPtr(Pemalloc(size, GC_FLAGS(this)&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func (this *HashTable) NextIndexInsertPtr(pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = this.NextIndexInsert(&tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) IndexUpdateMem(h ZendUlong, pData any, size int) any {
	var p any
	p = Pemalloc(size, GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return this.IndexUpdatePtr(h, p)
}
func (this *HashTable) NextIndexInsertMem(pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, this.NextIndexInsert(&tmp)) {
		zv.SetPtr(Pemalloc(size, GC_FLAGS(this)&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func (this *HashTable) FindPtr(key *ZendString) any {
	var zv *Zval
	zv = this.Find(key)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) FindExPtr(key *ZendString, known_hash ZendBool) any {
	var zv *Zval
	zv = this.FindEx(key, known_hash)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) StrFindPtr(str string, len_ int) any {
	var zv *Zval
	zv = this.StrFind(str, len_)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) IndexFindPtr(h ZendUlong) any {
	var zv *Zval
	zv = this.IndexFind(h)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) IndexFindDeref(h ZendUlong) *Zval {
	var zv *Zval = this.IndexFind(h)
	if zv != nil {
		ZVAL_DEREF(zv)
	}
	return zv
}
func (this *HashTable) FindDeref(str *ZendString) *Zval {
	var zv *Zval = this.Find(str)
	if zv != nil {
		ZVAL_DEREF(zv)
	}
	return zv
}
func (this *HashTable) StrFindDeref(str string, len_ int) *Zval {
	var zv *Zval = this.StrFind(str, len_)
	if zv != nil {
		ZVAL_DEREF(zv)
	}
	return zv
}
func ZendSymtableStrFindPtr(ht *HashTable, str *byte, len_ int) any {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ht.IndexFindPtr(idx)
	} else {
		return ht.StrFindPtr(str, len_)
	}
}
func (this *HashTable) GetCurrentDataPtrEx(pos *HashPosition) any {
	var zv *Zval
	zv = this.GetCurrentDataEx(pos)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func (this *HashTable) GetCurrentDataPtr() any {
	return this.GetCurrentDataPtrEx(this.GetNInternalPointer())
}
func ZEND_HASH_FILL_SET(_val *Zval)                    { ZVAL_COPY_VALUE(__fill_bkt.val, _val) }
func ZEND_HASH_FILL_SET_NULL()                         { ZVAL_NULL(__fill_bkt.val) }
func ZEND_HASH_FILL_SET_LONG(_val ZendLong)            { ZVAL_LONG(__fill_bkt.val, _val) }
func ZEND_HASH_FILL_SET_DOUBLE(_val float64)           { ZVAL_DOUBLE(__fill_bkt.val, _val) }
func ZEND_HASH_FILL_SET_STR(_val *ZendString)          { ZVAL_STR(__fill_bkt.val, _val) }
func ZEND_HASH_FILL_SET_STR_COPY(_val *ZendString)     { ZVAL_STR_COPY(__fill_bkt.val, _val) }
func ZEND_HASH_FILL_SET_INTERNED_STR(_val *ZendString) { ZVAL_INTERNED_STR(__fill_bkt.val, _val) }
func ZEND_HASH_FILL_NEXT() {
	__fill_bkt.h = __fill_idx
	__fill_bkt.key = nil
	__fill_bkt++
	__fill_idx++
}
func ZEND_HASH_FILL_ADD(_val *Zval) {
	ZEND_HASH_FILL_SET(_val)
	ZEND_HASH_FILL_NEXT()
}
func (this *HashTable) _appendEx(key *ZendString, zv *Zval, interned int) *Zval {
	var idx uint32 = b.PostInc(&(this.GetNNumUsed()))
	var nIndex uint32
	var p *Bucket = this.GetArData() + idx
	ZVAL_COPY_VALUE(p.GetVal(), zv)
	if interned == 0 && ZSTR_IS_INTERNED(key) == 0 {
		this.SubUFlags(HASH_FLAG_STATIC_KEYS)
		ZendStringAddref(key)
		ZendStringHashVal(key)
	}
	p.SetKey(key)
	p.SetH(key.GetH())
	nIndex = uint32(p.GetH() | this.GetNTableMask())
	p.GetVal().SetNext(this.Hash(nIndex))
	this.Hash(nIndex) = HT_IDX_TO_HASH(idx)
	this.GetNNumOfElements()++
	return p.GetVal()
}
func (this *HashTable) _append(key *ZendString, zv *Zval) *Zval { return this._appendEx(key, zv, 0) }
func (this *HashTable) _appendPtrEx(key *ZendString, ptr any, interned int) *Zval {
	var idx uint32 = b.PostInc(&(this.GetNNumUsed()))
	var nIndex uint32
	var p *Bucket = this.GetArData() + idx
	ZVAL_PTR(p.GetVal(), ptr)
	if interned == 0 && ZSTR_IS_INTERNED(key) == 0 {
		this.SubUFlags(HASH_FLAG_STATIC_KEYS)
		ZendStringAddref(key)
		ZendStringHashVal(key)
	}
	p.SetKey(key)
	p.SetH(key.GetH())
	nIndex = uint32(p.GetH() | this.GetNTableMask())
	p.GetVal().SetNext(this.Hash(nIndex))
	this.Hash(nIndex) = HT_IDX_TO_HASH(idx)
	this.GetNNumOfElements()++
	return p.GetVal()
}
func (this *HashTable) _appendPtr(key *ZendString, ptr any) *Zval {
	return this._appendPtrEx(key, ptr, 0)
}
func (this *HashTable) _appendInd(key *ZendString, ptr *Zval) {
	var idx uint32 = b.PostInc(&(this.GetNNumUsed()))
	var nIndex uint32
	var p *Bucket = this.GetArData() + idx
	ZVAL_INDIRECT(p.GetVal(), ptr)
	if ZSTR_IS_INTERNED(key) == 0 {
		this.SubUFlags(HASH_FLAG_STATIC_KEYS)
		ZendStringAddref(key)
		ZendStringHashVal(key)
	}
	p.SetKey(key)
	p.SetH(key.GetH())
	nIndex = uint32(p.GetH() | this.GetNTableMask())
	p.GetVal().SetNext(this.Hash(nIndex))
	this.Hash(nIndex) = HT_IDX_TO_HASH(idx)
	this.GetNNumOfElements()++
}
func (this *HashTable) AssertRc1() {}
func ZEND_HASH_IF_FULL_DO_RESIZE(ht *HashTable) {
	if ht.GetNNumUsed() >= ht.GetNTableSize() {
		ht.DoResize()
	}
}
func ZendHashCheckSize(nSize uint32) uint32 {
	/* Use big enough power of 2 */

	if nSize <= HT_MIN_SIZE {
		return HT_MIN_SIZE
	} else if nSize >= HT_MAX_SIZE {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (%u * %zu + %zu)", nSize, b.SizeOf("Bucket"), b.SizeOf("Bucket"))
	}
	nSize -= 1
	nSize |= nSize >> 1
	nSize |= nSize >> 2
	nSize |= nSize >> 4
	nSize |= nSize >> 8
	nSize |= nSize >> 16
	return nSize + 1
}
func (this *HashTable) RealInitPackedEx() {
	var data any
	if (GC_FLAGS(this) & IS_ARRAY_PERSISTENT) != 0 {
		data = Pemalloc(HT_SIZE_EX(this.GetNTableSize(), HT_MIN_MASK), 1)
	} else if this.GetNTableSize() == HT_MIN_SIZE {
		data = Emalloc(HT_SIZE_EX(HT_MIN_SIZE, HT_MIN_MASK))
	} else {
		data = Emalloc(HT_SIZE_EX(this.GetNTableSize(), HT_MIN_MASK))
	}
	this.SetDataAddr(data)

	/* Don't overwrite iterator count. */

	this.SetFlags(HASH_FLAG_PACKED | HASH_FLAG_STATIC_KEYS)
	this.HashResetPacked()
}
func (this *HashTable) RealInitMixedEx() {
	var data any
	var nSize uint32 = this.GetNTableSize()
	if (GC_FLAGS(this) & IS_ARRAY_PERSISTENT) != 0 {
		data = Pemalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)), 1)
	} else if nSize == HT_MIN_SIZE {
		data = Emalloc(HT_SIZE_EX(HT_MIN_SIZE, HT_SIZE_TO_MASK(HT_MIN_SIZE)))
		this.SetNTableMask(HT_SIZE_TO_MASK(HT_MIN_SIZE))
		this.SetDataAddr(data)

		/* Don't overwrite iterator count. */

		this.SetFlags(HASH_FLAG_STATIC_KEYS)
		HT_HASH_EX(data, 0) = -1
		HT_HASH_EX(data, 1) = -1
		HT_HASH_EX(data, 2) = -1
		HT_HASH_EX(data, 3) = -1
		HT_HASH_EX(data, 4) = -1
		HT_HASH_EX(data, 5) = -1
		HT_HASH_EX(data, 6) = -1
		HT_HASH_EX(data, 7) = -1
		HT_HASH_EX(data, 8) = -1
		HT_HASH_EX(data, 9) = -1
		HT_HASH_EX(data, 10) = -1
		HT_HASH_EX(data, 11) = -1
		HT_HASH_EX(data, 12) = -1
		HT_HASH_EX(data, 13) = -1
		HT_HASH_EX(data, 14) = -1
		HT_HASH_EX(data, 15) = -1
		return
	} else {
		data = Emalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)))
	}
	this.SetNTableMask(HT_SIZE_TO_MASK(nSize))
	this.SetDataAddr(data)
	this.SetUFlags(HASH_FLAG_STATIC_KEYS)
	this.HashReset()
}
func (this *HashTable) RealInitEx(packed int) {
	this.AssertRc1()
	ZEND_ASSERT(this.HasUFlags(HASH_FLAG_UNINITIALIZED))
	if packed != 0 {
		this.RealInitPackedEx()
	} else {
		this.RealInitMixedEx()
	}
}
func (this *HashTable) _initInt(nSize uint32, pDestructor DtorFuncT, persistent ZendBool) {
	GC_SET_REFCOUNT(this, 1)
	GC_TYPE_INFO(this) = IS_ARRAY | b.Cond(persistent != 0, GC_PERSISTENT<<GC_FLAGS_SHIFT, GC_COLLECTABLE<<GC_FLAGS_SHIFT)
	this.SetUFlags(HASH_FLAG_UNINITIALIZED)
	this.SetNTableMask(HT_MIN_MASK)
	this.SetDataAddr(&UninitializedBucket)
	this.SetNNumUsed(0)
	this.SetNNumOfElements(0)
	this.SetNInternalPointer(0)
	this.SetNNextFreeElement(0)
	this.SetPDestructor(pDestructor)
	this.SetNTableSize(ZendHashCheckSize(nSize))
}
func (this *HashTable) _init(nSize uint32, pDestructor DtorFuncT, persistent ZendBool) {
	this._initInt(nSize, pDestructor, persistent)
}
func _zendNewArray0() *HashTable {
	var ht *HashTable = Emalloc(b.SizeOf("HashTable"))
	ht._initInt(HT_MIN_SIZE, ZVAL_PTR_DTOR, 0)
	return ht
}
func _zendNewArray(nSize uint32) *HashTable {
	var ht *HashTable = Emalloc(b.SizeOf("HashTable"))
	ht._initInt(nSize, ZVAL_PTR_DTOR, 0)
	return ht
}
func ZendNewPair(val1 *Zval, val2 *Zval) *HashTable {
	var p *Bucket
	var ht *HashTable = Emalloc(b.SizeOf("HashTable"))
	ht._initInt(HT_MIN_SIZE, ZVAL_PTR_DTOR, 0)
	ht.SetNNextFreeElement(2)
	ht.SetNNumOfElements(ht.GetNNextFreeElement())
	ht.SetNNumUsed(ht.GetNNumOfElements())
	ht.RealInitPackedEx()
	p = ht.GetArData()
	ZVAL_COPY_VALUE(p.GetVal(), val1)
	p.SetH(0)
	p.SetKey(nil)
	p++
	ZVAL_COPY_VALUE(p.GetVal(), val2)
	p.SetH(1)
	p.SetKey(nil)
	return ht
}
func (this *HashTable) PackedGrow() {
	this.AssertRc1()
	if this.GetNTableSize() >= HT_MAX_SIZE {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (%u * %zu + %zu)", this.GetNTableSize()*2, b.SizeOf("Bucket"), b.SizeOf("Bucket"))
	}
	this.SetNTableSize(this.GetNTableSize() + this.GetNTableSize())
	this.SetDataAddr(Perealloc2(this.GetDataAddr(), HT_SIZE_EX(this.GetNTableSize(), HT_MIN_MASK), this.UsedSize(), GC_FLAGS(this)&IS_ARRAY_PERSISTENT))
}
func (this *HashTable) RealInit(packed ZendBool) {
	this.AssertRc1()
	this.RealInitEx(packed)
}
func (this *HashTable) RealInitPacked() {
	this.AssertRc1()
	this.RealInitPackedEx()
}
func (this *HashTable) RealInitMixed() {
	this.AssertRc1()
	this.RealInitMixedEx()
}
func (this *HashTable) PackedToHash() {
	var new_data any
	var old_data any = this.GetDataAddr()
	var old_buckets *Bucket = this.GetArData()
	var nSize uint32 = this.GetNTableSize()
	this.AssertRc1()
	this.SubUFlags(HASH_FLAG_PACKED)
	new_data = Pemalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)), GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
	this.SetNTableMask(HT_SIZE_TO_MASK(this.GetNTableSize()))
	this.SetDataAddr(new_data)
	memcpy(this.GetArData(), old_buckets, b.SizeOf("Bucket")*this.GetNNumUsed())
	Pefree(old_data, GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
	this.Rehash()
}
func (this *HashTable) ToPacked() {
	var new_data any
	var old_data any = this.GetDataAddr()
	var old_buckets *Bucket = this.GetArData()
	this.AssertRc1()
	new_data = Pemalloc(HT_SIZE_EX(this.GetNTableSize(), HT_MIN_MASK), GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
	this.AddUFlags(HASH_FLAG_PACKED | HASH_FLAG_STATIC_KEYS)
	this.SetNTableMask(HT_MIN_MASK)
	this.SetDataAddr(new_data)
	this.HashResetPacked()
	memcpy(this.GetArData(), old_buckets, b.SizeOf("Bucket")*this.GetNNumUsed())
	Pefree(old_data, GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
}
func (this *HashTable) Extend(nSize uint32, packed ZendBool) {
	this.AssertRc1()
	if nSize == 0 {
		return
	}
	if this.HasUFlags(HASH_FLAG_UNINITIALIZED) {
		if nSize > this.GetNTableSize() {
			this.SetNTableSize(ZendHashCheckSize(nSize))
		}
		this.RealInit(packed)
	} else {
		if packed != 0 {
			ZEND_ASSERT(this.HasUFlags(HASH_FLAG_PACKED))
			if nSize > this.GetNTableSize() {
				this.SetNTableSize(ZendHashCheckSize(nSize))
				this.SetDataAddr(Perealloc2(this.GetDataAddr(), HT_SIZE_EX(this.GetNTableSize(), HT_MIN_MASK), this.UsedSize(), GC_FLAGS(this)&IS_ARRAY_PERSISTENT))
			}
		} else {
			ZEND_ASSERT(!this.HasUFlags(HASH_FLAG_PACKED))
			if nSize > this.GetNTableSize() {
				var new_data any
				var old_data any = this.GetDataAddr()
				var old_buckets *Bucket = this.GetArData()
				nSize = ZendHashCheckSize(nSize)
				this.SetNTableSize(nSize)
				new_data = Pemalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)), GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
				this.SetNTableMask(HT_SIZE_TO_MASK(this.GetNTableSize()))
				this.SetDataAddr(new_data)
				memcpy(this.GetArData(), old_buckets, b.SizeOf("Bucket")*this.GetNNumUsed())
				Pefree(old_data, GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
				this.Rehash()
			}
		}
	}
}
func (this *HashTable) Discard(nNumUsed uint32) {
	var p *Bucket
	var end *Bucket
	var arData *Bucket
	var nIndex uint32
	arData = this.GetArData()
	p = arData + this.GetNNumUsed()
	end = arData + nNumUsed
	this.SetNNumUsed(nNumUsed)
	for p != end {
		p--
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		this.GetNNumOfElements()--

		/* Collision pointers always directed from higher to lower buckets */

		nIndex = p.GetH() | this.GetNTableMask()
		HT_HASH_EX(arData, nIndex) = p.GetVal().GetNext()
	}
}
func ZendArrayRecalcElements(ht *HashTable) uint32 {
	var val *Zval
	var num uint32 = ht.GetNNumOfElements()
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			val = _z
			if val.IsType(IS_INDIRECT) {
				if Z_INDIRECT_P(val).IsType(IS_UNDEF) {
					num--
				}
			}
		}
		break
	}
	return num
}
func ZendArrayCount(ht *HashTable) uint32 {
	var num uint32
	if ht.HasUFlags(HASH_FLAG_HAS_EMPTY_IND) {
		num = ZendArrayRecalcElements(ht)
		if ht.GetNNumOfElements() == num {
			ht.SubUFlags(HASH_FLAG_HAS_EMPTY_IND)
		}
	} else if ht == &(ExecutorGlobals.GetSymbolTable()) {
		num = ZendArrayRecalcElements(ht)
	} else {
		num = ht.GetNNumOfElements()
	}
	return num
}
func (this *HashTable) _getValidPos(pos HashPosition) HashPosition {
	for pos < this.GetNNumUsed() && Z_ISUNDEF(this.GetArData()[pos].GetVal()) {
		pos++
	}
	return pos
}
func (this *HashTable) _getCurrentPos() HashPosition {
	return this._getValidPos(this.GetNInternalPointer())
}
func (this *HashTable) GetCurrentPos() HashPosition { return this._getCurrentPos() }
func (this *HashTable) IteratorAdd(pos HashPosition) uint32 {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators()
	var end *HashTableIterator = iter + ExecutorGlobals.GetHtIteratorsCount()
	var idx uint32
	if !(this.IteratorsOverflow()) {
		this.IncIteratorsCount()
	}
	for iter != end {
		if iter.GetHt() == nil {
			iter.SetHt(this)
			iter.SetPos(pos)
			idx = iter - ExecutorGlobals.GetHtIterators()
			if idx+1 > ExecutorGlobals.GetHtIteratorsUsed() {
				ExecutorGlobals.SetHtIteratorsUsed(idx + 1)
			}
			return idx
		}
		iter++
	}
	if ExecutorGlobals.GetHtIterators() == ExecutorGlobals.GetHtIteratorsSlots() {
		ExecutorGlobals.SetHtIterators(Emalloc(b.SizeOf("HashTableIterator") * (ExecutorGlobals.GetHtIteratorsCount() + 8)))
		memcpy(ExecutorGlobals.GetHtIterators(), ExecutorGlobals.GetHtIteratorsSlots(), b.SizeOf("HashTableIterator")*ExecutorGlobals.GetHtIteratorsCount())
	} else {
		ExecutorGlobals.SetHtIterators(Erealloc(ExecutorGlobals.GetHtIterators(), b.SizeOf("HashTableIterator")*(ExecutorGlobals.GetHtIteratorsCount()+8)))
	}
	iter = ExecutorGlobals.GetHtIterators() + ExecutorGlobals.GetHtIteratorsCount()
	ExecutorGlobals.SetHtIteratorsCount(ExecutorGlobals.GetHtIteratorsCount() + 8)
	iter.SetHt(this)
	iter.SetPos(pos)
	memset(iter+1, 0, b.SizeOf("HashTableIterator")*7)
	idx = iter - ExecutorGlobals.GetHtIterators()
	ExecutorGlobals.SetHtIteratorsUsed(idx + 1)
	return idx
}
func ZendHashIteratorPos(idx uint32, ht *HashTable) HashPosition {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators() + idx
	ZEND_ASSERT(idx != uint32-1)
	if iter.GetHt() != ht {
		if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(iter.GetHt().IteratorsOverflow()) {
			iter.GetHt().DecIteratorsCount()
		}
		if !(ht.IteratorsOverflow()) {
			ht.IncIteratorsCount()
		}
		iter.SetHt(ht)
		iter.SetPos(ht._getCurrentPos())
	}
	return iter.GetPos()
}
func ZendHashIteratorPosEx(idx uint32, array *Zval) HashPosition {
	var ht *HashTable = array.GetArr()
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators() + idx
	ZEND_ASSERT(idx != uint32-1)
	if iter.GetHt() != ht {
		if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(ht.IteratorsOverflow()) {
			iter.GetHt().DecIteratorsCount()
		}
		SEPARATE_ARRAY(array)
		ht = array.GetArr()
		if !(ht.IteratorsOverflow()) {
			ht.IncIteratorsCount()
		}
		iter.SetHt(ht)
		iter.SetPos(ht._getCurrentPos())
	}
	return iter.GetPos()
}
func ZendHashIteratorDel(idx uint32) {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators() + idx
	ZEND_ASSERT(idx != uint32-1)
	if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(iter.GetHt().IteratorsOverflow()) {
		ZEND_ASSERT(iter.GetHt().GetNIteratorsCount() != 0)
		iter.GetHt().DecIteratorsCount()
	}
	iter.SetHt(nil)
	if idx == ExecutorGlobals.GetHtIteratorsUsed()-1 {
		for idx > 0 && ExecutorGlobals.GetHtIterators()[idx-1].GetHt() == nil {
			idx--
		}
		ExecutorGlobals.SetHtIteratorsUsed(idx)
	}
}
func (this *HashTable) _iteratorsRemove() {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators()
	var end *HashTableIterator = iter + ExecutorGlobals.GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == this {
			iter.SetHt(HT_POISONED_PTR)
		}
		iter++
	}
}
func (this *HashTable) IteratorsRemove() {
	if this.HasIterators() {
		this._iteratorsRemove()
	}
}
func (this *HashTable) IteratorsLowerPos(start HashPosition) HashPosition {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators()
	var end *HashTableIterator = iter + ExecutorGlobals.GetHtIteratorsUsed()
	var res HashPosition = this.GetNNumUsed()
	for iter != end {
		if iter.GetHt() == this {
			if iter.GetPos() >= start && iter.GetPos() < res {
				res = iter.GetPos()
			}
		}
		iter++
	}
	return res
}
func (this *HashTable) _iteratorsUpdate(from HashPosition, to HashPosition) {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators()
	var end *HashTableIterator = iter + ExecutorGlobals.GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == this && iter.GetPos() == from {
			iter.SetPos(to)
		}
		iter++
	}
}
func (this *HashTable) IteratorsAdvance(step HashPosition) {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators()
	var end *HashTableIterator = iter + ExecutorGlobals.GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == this {
			iter.SetPos(iter.GetPos() + step)
		}
		iter++
	}
}
func (this *HashTable) FindBucket(key *ZendString, known_hash ZendBool) *Bucket {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var arData *Bucket
	if known_hash != 0 {
		h = key.GetH()
	} else {
		h = ZendStringHashVal(key)
	}
	arData = this.GetArData()
	nIndex = h | this.GetNTableMask()
	idx = HT_HASH_EX(arData, nIndex)
	if idx == HT_INVALID_IDX {
		return nil
	}
	p = HT_HASH_TO_BUCKET_EX(arData, idx)
	if p.GetKey() == key {
		return p
	}
	for true {
		if p.GetH() == key.GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), key) != 0 {
			return p
		}
		idx = p.GetVal().GetNext()
		if idx == HT_INVALID_IDX {
			return nil
		}
		p = HT_HASH_TO_BUCKET_EX(arData, idx)
		if p.GetKey() == key {
			return p
		}
	}
}
func (this *HashTable) StrFindBucket(str *byte, len_ int, h ZendUlong) *Bucket {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var arData *Bucket
	arData = this.GetArData()
	nIndex = h | this.GetNTableMask()
	idx = HT_HASH_EX(arData, nIndex)
	for idx != HT_INVALID_IDX {
		ZEND_ASSERT(idx < HT_IDX_TO_HASH(this.GetNTableSize()))
		p = HT_HASH_TO_BUCKET_EX(arData, idx)
		if p.GetH() == h && p.GetKey() != nil && p.GetKey().GetLen() == len_ && !(memcmp(p.GetKey().GetVal(), str, len_)) {
			return p
		}
		idx = p.GetVal().GetNext()
	}
	return nil
}
func (this *HashTable) IndexFindBucket(h ZendUlong) *Bucket {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var arData *Bucket
	arData = this.GetArData()
	nIndex = h | this.GetNTableMask()
	idx = HT_HASH_EX(arData, nIndex)
	for idx != HT_INVALID_IDX {
		ZEND_ASSERT(idx < HT_IDX_TO_HASH(this.GetNTableSize()))
		p = HT_HASH_TO_BUCKET_EX(arData, idx)
		if p.GetH() == h && p.GetKey() == nil {
			return p
		}
		idx = p.GetVal().GetNext()
	}
	return nil
}
func (this *HashTable) _addOrUpdateI(key *ZendString, pData *Zval, flag uint32) *Zval {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var arData *Bucket
	this.AssertRc1()
	if this.HasUFlags(HASH_FLAG_UNINITIALIZED | HASH_FLAG_PACKED) {
		if this.HasUFlags(HASH_FLAG_UNINITIALIZED) {
			this.RealInitMixed()
			if ZSTR_IS_INTERNED(key) == 0 {
				ZendStringAddref(key)
				this.SubUFlags(HASH_FLAG_STATIC_KEYS)
				ZendStringHashVal(key)
			}
			goto add_to_hash
		} else {
			this.PackedToHash()
			if ZSTR_IS_INTERNED(key) == 0 {
				ZendStringAddref(key)
				this.SubUFlags(HASH_FLAG_STATIC_KEYS)
				ZendStringHashVal(key)
			}
		}
	} else if (flag&HASH_ADD_NEW) == 0 || core.ZEND_DEBUG != 0 {
		p = this.FindBucket(key, 0)
		if p != nil {
			var data *Zval
			ZEND_ASSERT((flag & HASH_ADD_NEW) == 0)
			if (flag & HASH_ADD) != 0 {
				if (flag & HASH_UPDATE_INDIRECT) == 0 {
					return nil
				}
				ZEND_ASSERT(p.GetVal() != pData)
				data = p.GetVal()
				if data.IsType(IS_INDIRECT) {
					data = data.GetZv()
					if data.GetType() != IS_UNDEF {
						return nil
					}
				} else {
					return nil
				}
			} else {
				ZEND_ASSERT(p.GetVal() != pData)
				data = p.GetVal()
				if (flag&HASH_UPDATE_INDIRECT) != 0 && data.IsType(IS_INDIRECT) {
					data = data.GetZv()
				}
			}
			if this.GetPDestructor() != nil {
				this.GetPDestructor()(data)
			}
			ZVAL_COPY_VALUE(data, pData)
			return data
		}
		if ZSTR_IS_INTERNED(key) == 0 {
			ZendStringAddref(key)
			this.SubUFlags(HASH_FLAG_STATIC_KEYS)
		}
	} else if ZSTR_IS_INTERNED(key) == 0 {
		ZendStringAddref(key)
		this.SubUFlags(HASH_FLAG_STATIC_KEYS)
		ZendStringHashVal(key)
	}
	ZEND_HASH_IF_FULL_DO_RESIZE(this)
add_to_hash:
	this.GetNNumUsed()++
	idx = this.GetNNumUsed() - 1
	this.GetNNumOfElements()++
	arData = this.GetArData()
	p = arData + idx
	p.SetKey(key)
	h = key.GetH()
	p.SetH(h)
	nIndex = h | this.GetNTableMask()
	p.GetVal().SetNext(HT_HASH_EX(arData, nIndex))
	HT_HASH_EX(arData, nIndex) = HT_IDX_TO_HASH(idx)
	ZVAL_COPY_VALUE(p.GetVal(), pData)
	return p.GetVal()
}
func (this *HashTable) _strAddOrUpdateI(str *byte, len_ int, h ZendUlong, pData *Zval, flag uint32) *Zval {
	var key *ZendString
	var nIndex uint32
	var idx uint32
	var p *Bucket
	this.AssertRc1()
	if this.HasUFlags(HASH_FLAG_UNINITIALIZED | HASH_FLAG_PACKED) {
		if this.HasUFlags(HASH_FLAG_UNINITIALIZED) {
			this.RealInitMixed()
			goto add_to_hash
		} else {
			this.PackedToHash()
		}
	} else if (flag & HASH_ADD_NEW) == 0 {
		p = this.StrFindBucket(str, len_, h)
		if p != nil {
			var data *Zval
			if (flag & HASH_ADD) != 0 {
				if (flag & HASH_UPDATE_INDIRECT) == 0 {
					return nil
				}
				ZEND_ASSERT(p.GetVal() != pData)
				data = p.GetVal()
				if data.IsType(IS_INDIRECT) {
					data = data.GetZv()
					if data.GetType() != IS_UNDEF {
						return nil
					}
				} else {
					return nil
				}
			} else {
				ZEND_ASSERT(p.GetVal() != pData)
				data = p.GetVal()
				if (flag&HASH_UPDATE_INDIRECT) != 0 && data.IsType(IS_INDIRECT) {
					data = data.GetZv()
				}
			}
			if this.GetPDestructor() != nil {
				this.GetPDestructor()(data)
			}
			ZVAL_COPY_VALUE(data, pData)
			return data
		}
	}
	ZEND_HASH_IF_FULL_DO_RESIZE(this)
add_to_hash:
	this.GetNNumUsed()++
	idx = this.GetNNumUsed() - 1
	this.GetNNumOfElements()++
	p = this.GetArData() + idx
	key = ZendStringInit(str, len_, GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
	p.SetKey(key)
	key.SetH(h)
	p.SetH(key.GetH())
	this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	ZVAL_COPY_VALUE(p.GetVal(), pData)
	nIndex = h | this.GetNTableMask()
	p.GetVal().SetNext(this.Hash(nIndex))
	this.Hash(nIndex) = HT_IDX_TO_HASH(idx)
	return p.GetVal()
}
func (this *HashTable) AddOrUpdate(key *ZendString, pData *Zval, flag uint32) *Zval {
	if flag == HASH_ADD {
		return this.Add(key, pData)
	} else if flag == HASH_ADD_NEW {
		return this.AddNew(key, pData)
	} else if flag == HASH_UPDATE {
		return this.Update(key, pData)
	} else {
		ZEND_ASSERT(flag == (HASH_UPDATE | HASH_UPDATE_INDIRECT))
		return this.UpdateInd(key, pData)
	}
}
func (this *HashTable) Add(key *ZendString, pData *Zval) *Zval {
	return this._addOrUpdateI(key, pData, HASH_ADD)
}
func (this *HashTable) Update(key *ZendString, pData *Zval) *Zval {
	return this._addOrUpdateI(key, pData, HASH_UPDATE)
}
func (this *HashTable) UpdateInd(key *ZendString, pData *Zval) *Zval {
	return this._addOrUpdateI(key, pData, HASH_UPDATE|HASH_UPDATE_INDIRECT)
}
func (this *HashTable) AddNew(key *ZendString, pData *Zval) *Zval {
	return this._addOrUpdateI(key, pData, HASH_ADD_NEW)
}
func (this *HashTable) StrAddOrUpdate(str *byte, len_ int, pData *Zval, flag uint32) *Zval {
	if flag == HASH_ADD {
		return this.StrAdd(str, len_, pData)
	} else if flag == HASH_ADD_NEW {
		return this.StrAddNew(str, len_, pData)
	} else if flag == HASH_UPDATE {
		return this.StrUpdate(str, len_, pData)
	} else {
		ZEND_ASSERT(flag == (HASH_UPDATE | HASH_UPDATE_INDIRECT))
		return this.StrUpdateInd(str, len_, pData)
	}
}
func (this *HashTable) StrUpdate(str string, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return this._strAddOrUpdateI(str, len_, h, pData, HASH_UPDATE)
}
func (this *HashTable) StrUpdateInd(str string, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return this._strAddOrUpdateI(str, len_, h, pData, HASH_UPDATE|HASH_UPDATE_INDIRECT)
}
func (this *HashTable) StrAdd(str *byte, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return this._strAddOrUpdateI(str, len_, h, pData, HASH_ADD)
}
func (this *HashTable) StrAddNew(str *byte, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return this._strAddOrUpdateI(str, len_, h, pData, HASH_ADD_NEW)
}
func (this *HashTable) IndexAddEmptyElement(h ZendUlong) *Zval {
	var dummy Zval
	ZVAL_NULL(&dummy)
	return this.IndexAdd(h, &dummy)
}
func (this *HashTable) AddEmptyElement(key *ZendString) *Zval {
	var dummy Zval
	ZVAL_NULL(&dummy)
	return this.Add(key, &dummy)
}
func (this *HashTable) StrAddEmptyElement(str *byte, len_ int) *Zval {
	var dummy Zval
	ZVAL_NULL(&dummy)
	return this.StrAdd(str, len_, &dummy)
}
func (this *HashTable) _indexAddOrUpdateI(h ZendUlong, pData *Zval, flag uint32) *Zval {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	this.AssertRc1()
	if this.HasUFlags(HASH_FLAG_PACKED) {
		if h < this.GetNNumUsed() {
			p = this.GetArData() + h
			if p.GetVal().GetType() != IS_UNDEF {
			replace:
				if (flag & HASH_ADD) != 0 {
					return nil
				}
				if this.GetPDestructor() != nil {
					this.GetPDestructor()(p.GetVal())
				}
				ZVAL_COPY_VALUE(p.GetVal(), pData)
				return p.GetVal()
			} else {
				goto convert_to_hash
			}
		} else if h < this.GetNTableSize() {
		add_to_packed:
			p = this.GetArData() + h

			/* incremental initialization of empty Buckets */

			if (flag & (HASH_ADD_NEW | HASH_ADD_NEXT)) != (HASH_ADD_NEW | HASH_ADD_NEXT) {
				if h > this.GetNNumUsed() {
					var q *Bucket = this.GetArData() + this.GetNNumUsed()
					for q != p {
						ZVAL_UNDEF(q.GetVal())
						q++
					}
				}
			}
			this.SetNNumUsed(h + 1)
			this.SetNNextFreeElement(this.GetNNumUsed())
			goto add
		} else if h>>1 < this.GetNTableSize() && this.GetNTableSize()>>1 < this.GetNNumOfElements() {
			this.PackedGrow()
			goto add_to_packed
		} else {
			if this.GetNNumUsed() >= this.GetNTableSize() {
				this.SetNTableSize(this.GetNTableSize() + this.GetNTableSize())
			}
		convert_to_hash:
			this.PackedToHash()
		}
	} else if this.HasUFlags(HASH_FLAG_UNINITIALIZED) {
		if h < this.GetNTableSize() {
			this.RealInitPackedEx()
			goto add_to_packed
		}
		this.RealInitMixed()
	} else {
		if (flag&HASH_ADD_NEW) == 0 || core.ZEND_DEBUG != 0 {
			p = this.IndexFindBucket(h)
			if p != nil {
				ZEND_ASSERT((flag & HASH_ADD_NEW) == 0)
				goto replace
			}
		}
		ZEND_HASH_IF_FULL_DO_RESIZE(this)
	}
	this.GetNNumUsed()++
	idx = this.GetNNumUsed() - 1
	nIndex = h | this.GetNTableMask()
	p = this.GetArData() + idx
	p.GetVal().SetNext(this.Hash(nIndex))
	this.Hash(nIndex) = HT_IDX_TO_HASH(idx)
	if ZendLong(h >= ZendLong(this.GetNNextFreeElement())) != 0 {
		if h < ZEND_LONG_MAX {
			this.SetNNextFreeElement(h + 1)
		} else {
			this.SetNNextFreeElement(ZEND_LONG_MAX)
		}
	}
add:
	this.GetNNumOfElements()++
	p.SetH(h)
	p.SetKey(nil)
	ZVAL_COPY_VALUE(p.GetVal(), pData)
	return p.GetVal()
}
func (this *HashTable) IndexAddOrUpdate(h ZendUlong, pData *Zval, flag uint32) *Zval {
	if flag == HASH_ADD {
		return this.IndexAdd(h, pData)
	} else if flag == (HASH_ADD | HASH_ADD_NEW) {
		return this.IndexAddNew(h, pData)
	} else if flag == (HASH_ADD | HASH_ADD_NEXT) {
		ZEND_ASSERT(h == this.GetNNextFreeElement())
		return this.NextIndexInsert(pData)
	} else if flag == (HASH_ADD | HASH_ADD_NEW | HASH_ADD_NEXT) {
		ZEND_ASSERT(h == this.GetNNextFreeElement())
		return this.NextIndexInsertNew(pData)
	} else {
		ZEND_ASSERT(flag == HASH_UPDATE)
		return this.IndexUpdate(h, pData)
	}
}
func (this *HashTable) IndexAdd(h ZendUlong, pData *Zval) *Zval {
	return this._indexAddOrUpdateI(h, pData, HASH_ADD)
}
func (this *HashTable) IndexAddNew(h ZendUlong, pData *Zval) *Zval {
	return this._indexAddOrUpdateI(h, pData, HASH_ADD|HASH_ADD_NEW)
}
func (this *HashTable) IndexUpdate(h ZendUlong, pData *Zval) *Zval {
	return this._indexAddOrUpdateI(h, pData, HASH_UPDATE)
}
func (this *HashTable) NextIndexInsert(pData *Zval) *Zval {
	return this._indexAddOrUpdateI(this.GetNNextFreeElement(), pData, HASH_ADD|HASH_ADD_NEXT)
}
func (this *HashTable) NextIndexInsertNew(pData *Zval) *Zval {
	return this._indexAddOrUpdateI(this.GetNNextFreeElement(), pData, HASH_ADD|HASH_ADD_NEW|HASH_ADD_NEXT)
}
func (this *HashTable) SetBucketKey(b *Bucket, key *ZendString) *Zval {
	var nIndex uint32
	var idx uint32
	var i uint32
	var p *Bucket
	var arData *Bucket
	this.AssertRc1()
	ZEND_ASSERT(!this.HasUFlags(HASH_FLAG_PACKED))
	p = this.FindBucket(key, 0)
	if p != nil {
		if p == b {
			return p.GetVal()
		} else {
			return nil
		}
	}
	if ZSTR_IS_INTERNED(key) == 0 {
		ZendStringAddref(key)
		this.SubUFlags(HASH_FLAG_STATIC_KEYS)
	}
	arData = this.GetArData()

	/* del from hash */

	idx = HT_IDX_TO_HASH(b - arData)
	nIndex = b.GetH() | this.GetNTableMask()
	i = HT_HASH_EX(arData, nIndex)
	if i == idx {
		HT_HASH_EX(arData, nIndex) = b.GetVal().GetNext()
	} else {
		p = HT_HASH_TO_BUCKET_EX(arData, i)
		for p.GetVal().GetNext() != idx {
			i = p.GetVal().GetNext()
			p = HT_HASH_TO_BUCKET_EX(arData, i)
		}
		p.GetVal().SetNext(b.GetVal().GetNext())
	}
	ZendStringRelease(b.GetKey())

	/* add to hash */

	idx = b - arData
	b.SetKey(key)
	b.SetH(key.GetH())
	nIndex = b.GetH() | this.GetNTableMask()
	idx = HT_IDX_TO_HASH(idx)
	i = HT_HASH_EX(arData, nIndex)
	if i == HT_INVALID_IDX || i < idx {
		b.GetVal().SetNext(i)
		HT_HASH_EX(arData, nIndex) = idx
	} else {
		p = HT_HASH_TO_BUCKET_EX(arData, i)
		for p.GetVal().GetNext() != HT_INVALID_IDX && p.GetVal().GetNext() > idx {
			i = p.GetVal().GetNext()
			p = HT_HASH_TO_BUCKET_EX(arData, i)
		}
		b.GetVal().SetNext(p.GetVal().GetNext())
		p.GetVal().SetNext(idx)
	}
	return b.GetVal()
}
func (this *HashTable) DoResize() {
	this.AssertRc1()
	if this.GetNNumUsed() > this.GetNNumOfElements()+(this.GetNNumOfElements()>>5) {
		this.Rehash()
	} else if this.GetNTableSize() < HT_MAX_SIZE {
		var new_data any
		var old_data any = this.GetDataAddr()
		var nSize uint32 = this.GetNTableSize() + this.GetNTableSize()
		var old_buckets *Bucket = this.GetArData()
		this.SetNTableSize(nSize)
		new_data = Pemalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)), GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
		this.SetNTableMask(HT_SIZE_TO_MASK(this.GetNTableSize()))
		this.SetDataAddr(new_data)
		memcpy(this.GetArData(), old_buckets, b.SizeOf("Bucket")*this.GetNNumUsed())
		Pefree(old_data, GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
		this.Rehash()
	} else {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (%u * %zu + %zu)", this.GetNTableSize()*2, b.SizeOf("Bucket")+b.SizeOf("uint32_t"), b.SizeOf("Bucket"))
	}
}
func (this *HashTable) Rehash() int {
	var p *Bucket
	var nIndex uint32
	var i uint32
	if this.GetNNumOfElements() == 0 {
		if !this.HasUFlags(HASH_FLAG_UNINITIALIZED) {
			this.SetNNumUsed(0)
			this.HashReset()
		}
		return SUCCESS
	}
	this.HashReset()
	i = 0
	p = this.GetArData()
	if this.IsWithoutHoles() {
		for {
			nIndex = p.GetH() | this.GetNTableMask()
			p.GetVal().SetNext(this.Hash(nIndex))
			this.Hash(nIndex) = HT_IDX_TO_HASH(i)
			p++
			if b.PreInc(&i) >= this.GetNNumUsed() {
				break
			}
		}
	} else {
		var old_num_used uint32 = this.GetNNumUsed()
		for {
			if p.GetVal().IsType(IS_UNDEF) {
				var j uint32 = i
				var q *Bucket = p
				if !(this.HasIterators()) {
					for b.PreInc(&i) < this.GetNNumUsed() {
						p++
						if p.GetVal().GetTypeInfo() != IS_UNDEF {
							ZVAL_COPY_VALUE(q.GetVal(), p.GetVal())
							q.SetH(p.GetH())
							nIndex = q.GetH() | this.GetNTableMask()
							q.SetKey(p.GetKey())
							q.GetVal().SetNext(this.Hash(nIndex))
							this.Hash(nIndex) = HT_IDX_TO_HASH(j)
							if this.GetNInternalPointer() == i {
								this.SetNInternalPointer(j)
							}
							q++
							j++
						}
					}
				} else {
					var iter_pos uint32 = this.IteratorsLowerPos(0)
					for b.PreInc(&i) < this.GetNNumUsed() {
						p++
						if p.GetVal().GetTypeInfo() != IS_UNDEF {
							ZVAL_COPY_VALUE(q.GetVal(), p.GetVal())
							q.SetH(p.GetH())
							nIndex = q.GetH() | this.GetNTableMask()
							q.SetKey(p.GetKey())
							q.GetVal().SetNext(this.Hash(nIndex))
							this.Hash(nIndex) = HT_IDX_TO_HASH(j)
							if this.GetNInternalPointer() == i {
								this.SetNInternalPointer(j)
							}
							if i >= iter_pos {
								for {
									this.IteratorsUpdate(iter_pos, j)
									iter_pos = this.IteratorsLowerPos(iter_pos + 1)
									if iter_pos >= i {
										break
									}
								}
							}
							q++
							j++
						}
					}
				}
				this.SetNNumUsed(j)
				break
			}
			nIndex = p.GetH() | this.GetNTableMask()
			p.GetVal().SetNext(this.Hash(nIndex))
			this.Hash(nIndex) = HT_IDX_TO_HASH(i)
			p++
			if b.PreInc(&i) >= this.GetNNumUsed() {
				break
			}
		}

		/* Migrate pointer to one past the end of the array to the new one past the end, so that
		 * newly inserted elements are picked up correctly. */

		if this.HasIterators() {
			this._iteratorsUpdate(old_num_used, this.GetNNumUsed())
		}

		/* Migrate pointer to one past the end of the array to the new one past the end, so that
		 * newly inserted elements are picked up correctly. */

	}
	return SUCCESS
}
func (this *HashTable) _delElEx(idx uint32, p *Bucket, prev *Bucket) {
	if !this.HasUFlags(HASH_FLAG_PACKED) {
		if prev != nil {
			prev.GetVal().SetNext(p.GetVal().GetNext())
		} else {
			this.Hash(p.GetH() | this.GetNTableMask()) = p.GetVal().GetNext()
		}
	}
	idx = HT_HASH_TO_IDX(idx)
	this.GetNNumOfElements()--
	if this.GetNInternalPointer() == idx || this.HasIterators() {
		var new_idx uint32
		new_idx = idx
		for true {
			new_idx++
			if new_idx >= this.GetNNumUsed() {
				break
			} else if this.GetArData()[new_idx].GetVal().GetType() != IS_UNDEF {
				break
			}
		}
		if this.GetNInternalPointer() == idx {
			this.SetNInternalPointer(new_idx)
		}
		this.IteratorsUpdate(idx, new_idx)
	}
	if this.GetNNumUsed()-1 == idx {
		for {
			this.GetNNumUsed()--
			if !(this.GetNNumUsed() > 0 && this.GetArData()[this.GetNNumUsed()-1].GetVal().IsType(IS_UNDEF)) {
				break
			}
		}
		this.SetNInternalPointer(MIN(this.GetNInternalPointer(), this.GetNNumUsed()))
	}
	if p.GetKey() != nil {
		ZendStringRelease(p.GetKey())
	}
	if this.GetPDestructor() != nil {
		var tmp Zval
		ZVAL_COPY_VALUE(&tmp, p.GetVal())
		ZVAL_UNDEF(p.GetVal())
		this.GetPDestructor()(&tmp)
	} else {
		ZVAL_UNDEF(p.GetVal())
	}
}
func (this *HashTable) _delEl(idx uint32, p *Bucket) {
	var prev *Bucket = nil
	if !this.HasUFlags(HASH_FLAG_PACKED) {
		var nIndex uint32 = p.GetH() | this.GetNTableMask()
		var i uint32 = this.Hash(nIndex)
		if i != idx {
			prev = this.HashToBucket(i)
			for prev.GetVal().GetNext() != idx {
				i = prev.GetVal().GetNext()
				prev = this.HashToBucket(i)
			}
		}
	}
	this._delElEx(idx, p, prev)
}
func (this *HashTable) DelBucket(p *Bucket) {
	this.AssertRc1()
	this._delEl(HT_IDX_TO_HASH(p-this.GetArData()), p)
}
func (this *HashTable) Del(key *ZendString) int {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	this.AssertRc1()
	h = ZendStringHashVal(key)
	nIndex = h | this.GetNTableMask()
	idx = this.Hash(nIndex)
	for idx != HT_INVALID_IDX {
		p = this.HashToBucket(idx)
		if p.GetKey() == key || p.GetH() == h && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), key) != 0 {
			this._delElEx(idx, p, prev)
			return SUCCESS
		}
		prev = p
		idx = p.GetVal().GetNext()
	}
	return FAILURE
}
func (this *HashTable) DelInd(key *ZendString) int {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	this.AssertRc1()
	h = ZendStringHashVal(key)
	nIndex = h | this.GetNTableMask()
	idx = this.Hash(nIndex)
	for idx != HT_INVALID_IDX {
		p = this.HashToBucket(idx)
		if p.GetKey() == key || p.GetH() == h && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), key) != 0 {
			if p.GetVal().IsType(IS_INDIRECT) {
				var data *Zval = p.GetVal().GetZv()
				if data.IsType(IS_UNDEF) {
					return FAILURE
				} else {
					if this.GetPDestructor() != nil {
						var tmp Zval
						ZVAL_COPY_VALUE(&tmp, data)
						ZVAL_UNDEF(data)
						this.GetPDestructor()(&tmp)
					} else {
						ZVAL_UNDEF(data)
					}
					this.AddUFlags(HASH_FLAG_HAS_EMPTY_IND)
				}
			} else {
				this._delElEx(idx, p, prev)
			}
			return SUCCESS
		}
		prev = p
		idx = p.GetVal().GetNext()
	}
	return FAILURE
}
func (this *HashTable) StrDelInd(str *byte, len_ int) int {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	this.AssertRc1()
	h = ZendInlineHashFunc(str, len_)
	nIndex = h | this.GetNTableMask()
	idx = this.Hash(nIndex)
	for idx != HT_INVALID_IDX {
		p = this.HashToBucket(idx)
		if p.GetH() == h && p.GetKey() != nil && p.GetKey().GetLen() == len_ && !(memcmp(p.GetKey().GetVal(), str, len_)) {
			if p.GetVal().IsType(IS_INDIRECT) {
				var data *Zval = p.GetVal().GetZv()
				if data.IsType(IS_UNDEF) {
					return FAILURE
				} else {
					if this.GetPDestructor() != nil {
						this.GetPDestructor()(data)
					}
					ZVAL_UNDEF(data)
					this.AddUFlags(HASH_FLAG_HAS_EMPTY_IND)
				}
			} else {
				this._delElEx(idx, p, prev)
			}
			return SUCCESS
		}
		prev = p
		idx = p.GetVal().GetNext()
	}
	return FAILURE
}
func (this *HashTable) StrDel(str *byte, len_ int) int {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	this.AssertRc1()
	h = ZendInlineHashFunc(str, len_)
	nIndex = h | this.GetNTableMask()
	idx = this.Hash(nIndex)
	for idx != HT_INVALID_IDX {
		p = this.HashToBucket(idx)
		if p.GetH() == h && p.GetKey() != nil && p.GetKey().GetLen() == len_ && !(memcmp(p.GetKey().GetVal(), str, len_)) {
			this._delElEx(idx, p, prev)
			return SUCCESS
		}
		prev = p
		idx = p.GetVal().GetNext()
	}
	return FAILURE
}
func (this *HashTable) IndexDel(h ZendUlong) int {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	this.AssertRc1()
	if this.HasUFlags(HASH_FLAG_PACKED) {
		if h < this.GetNNumUsed() {
			p = this.GetArData() + h
			if p.GetVal().GetType() != IS_UNDEF {
				this._delElEx(HT_IDX_TO_HASH(h), p, nil)
				return SUCCESS
			}
		}
		return FAILURE
	}
	nIndex = h | this.GetNTableMask()
	idx = this.Hash(nIndex)
	for idx != HT_INVALID_IDX {
		p = this.HashToBucket(idx)
		if p.GetH() == h && p.GetKey() == nil {
			this._delElEx(idx, p, prev)
			return SUCCESS
		}
		prev = p
		idx = p.GetVal().GetNext()
	}
	return FAILURE
}
func (this *HashTable) Destroy() {
	var p *Bucket
	var end *Bucket
	if this.GetNNumUsed() != 0 {
		p = this.GetArData()
		end = p + this.GetNNumUsed()
		if this.GetPDestructor() != nil {
			if this.HasStaticKeysOnly() {
				if this.IsWithoutHoles() {
					for {
						this.GetPDestructor()(p.GetVal())
						if b.PreInc(&p) == end {
							break
						}
					}
				} else {
					for {
						if p.GetVal().GetType() != IS_UNDEF {
							this.GetPDestructor()(p.GetVal())
						}
						if b.PreInc(&p) == end {
							break
						}
					}
				}
			} else if this.IsWithoutHoles() {
				for {
					this.GetPDestructor()(p.GetVal())
					if p.GetKey() != nil {
						ZendStringRelease(p.GetKey())
					}
					if b.PreInc(&p) == end {
						break
					}
				}
			} else {
				for {
					if p.GetVal().GetType() != IS_UNDEF {
						this.GetPDestructor()(p.GetVal())
						if p.GetKey() != nil {
							ZendStringRelease(p.GetKey())
						}
					}
					if b.PreInc(&p) == end {
						break
					}
				}
			}
		} else {
			if !(this.HasStaticKeysOnly()) {
				for {
					if p.GetVal().GetType() != IS_UNDEF {
						if p.GetKey() != nil {
							ZendStringRelease(p.GetKey())
						}
					}
					if b.PreInc(&p) == end {
						break
					}
				}
			}
		}
		this.IteratorsRemove()
	} else if this.HasUFlags(HASH_FLAG_UNINITIALIZED) {
		return
	}
	Pefree(this.GetDataAddr(), GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
}
func ZendArrayDestroy(ht *HashTable) {
	var p *Bucket
	var end *Bucket

	/* break possible cycles */

	GC_REMOVE_FROM_BUFFER(ht)
	GC_TYPE_INFO(ht) = IS_NULL
	if ht.GetNNumUsed() != 0 {

		/* In some rare cases destructors of regular arrays may be changed */

		if ht.GetPDestructor() != ZVAL_PTR_DTOR {
			ht.Destroy()
			goto free_ht
		}
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if ht.HasStaticKeysOnly() {
			for {
				IZvalPtrDtor(p.GetVal())
				if b.PreInc(&p) == end {
					break
				}
			}
		} else if ht.IsWithoutHoles() {
			for {
				IZvalPtrDtor(p.GetVal())
				if p.GetKey() != nil {
					ZendStringReleaseEx(p.GetKey(), 0)
				}
				if b.PreInc(&p) == end {
					break
				}
			}
		} else {
			for {
				if p.GetVal().GetType() != IS_UNDEF {
					IZvalPtrDtor(p.GetVal())
					if p.GetKey() != nil {
						ZendStringReleaseEx(p.GetKey(), 0)
					}
				}
				if b.PreInc(&p) == end {
					break
				}
			}
		}
	} else if ht.HasUFlags(HASH_FLAG_UNINITIALIZED) {
		goto free_ht
	}
	Efree(ht.GetDataAddr())
free_ht:
	ht.IteratorsRemove()
	FREE_HASHTABLE(ht)
}
func (this *HashTable) Clean() {
	var p *Bucket
	var end *Bucket
	this.AssertRc1()
	if this.GetNNumUsed() != 0 {
		p = this.GetArData()
		end = p + this.GetNNumUsed()
		if this.GetPDestructor() != nil {
			if this.HasStaticKeysOnly() {
				if this.IsWithoutHoles() {
					for {
						this.GetPDestructor()(p.GetVal())
						if b.PreInc(&p) == end {
							break
						}
					}
				} else {
					for {
						if p.GetVal().GetType() != IS_UNDEF {
							this.GetPDestructor()(p.GetVal())
						}
						if b.PreInc(&p) == end {
							break
						}
					}
				}
			} else if this.IsWithoutHoles() {
				for {
					this.GetPDestructor()(p.GetVal())
					if p.GetKey() != nil {
						ZendStringRelease(p.GetKey())
					}
					if b.PreInc(&p) == end {
						break
					}
				}
			} else {
				for {
					if p.GetVal().GetType() != IS_UNDEF {
						this.GetPDestructor()(p.GetVal())
						if p.GetKey() != nil {
							ZendStringRelease(p.GetKey())
						}
					}
					if b.PreInc(&p) == end {
						break
					}
				}
			}
		} else {
			if !(this.HasStaticKeysOnly()) {
				if this.IsWithoutHoles() {
					for {
						if p.GetKey() != nil {
							ZendStringRelease(p.GetKey())
						}
						if b.PreInc(&p) == end {
							break
						}
					}
				} else {
					for {
						if p.GetVal().GetType() != IS_UNDEF {
							if p.GetKey() != nil {
								ZendStringRelease(p.GetKey())
							}
						}
						if b.PreInc(&p) == end {
							break
						}
					}
				}
			}
		}
		if !this.HasUFlags(HASH_FLAG_PACKED) {
			this.HashReset()
		}
	}
	this.SetNNumUsed(0)
	this.SetNNumOfElements(0)
	this.SetNNextFreeElement(0)
	this.SetNInternalPointer(0)
}
func ZendSymtableClean(ht *HashTable) {
	var p *Bucket
	var end *Bucket
	ht.AssertRc1()
	if ht.GetNNumUsed() != 0 {
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if ht.HasStaticKeysOnly() {
			for {
				IZvalPtrDtor(p.GetVal())
				if b.PreInc(&p) == end {
					break
				}
			}
		} else if ht.IsWithoutHoles() {
			for {
				IZvalPtrDtor(p.GetVal())
				if p.GetKey() != nil {
					ZendStringRelease(p.GetKey())
				}
				if b.PreInc(&p) == end {
					break
				}
			}
		} else {
			for {
				if p.GetVal().GetType() != IS_UNDEF {
					IZvalPtrDtor(p.GetVal())
					if p.GetKey() != nil {
						ZendStringRelease(p.GetKey())
					}
				}
				if b.PreInc(&p) == end {
					break
				}
			}
		}
		ht.HashReset()
	}
	ht.SetNNumUsed(0)
	ht.SetNNumOfElements(0)
	ht.SetNNextFreeElement(0)
	ht.SetNInternalPointer(0)
}
func (this *HashTable) GracefulDestroy() {
	var idx uint32
	var p *Bucket
	this.AssertRc1()
	p = this.GetArData()
	for idx = 0; idx < this.GetNNumUsed(); {
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		this._delEl(HT_IDX_TO_HASH(idx), p)
		idx++
		p++
	}
	if !this.HasUFlags(HASH_FLAG_UNINITIALIZED) {
		Pefree(this.GetDataAddr(), GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
	}
}
func (this *HashTable) GracefulReverseDestroy() {
	var idx uint32
	var p *Bucket
	this.AssertRc1()
	idx = this.GetNNumUsed()
	p = this.GetArData() + this.GetNNumUsed()
	for idx > 0 {
		idx--
		p--
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		this._delEl(HT_IDX_TO_HASH(idx), p)
	}
	if !this.HasUFlags(HASH_FLAG_UNINITIALIZED) {
		Pefree(this.GetDataAddr(), GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
	}
}
func (this *HashTable) Apply(apply_func ApplyFuncT) {
	var idx uint32
	var p *Bucket
	var result int
	for idx = 0; idx < this.GetNNumUsed(); idx++ {
		p = this.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		result = apply_func(p.GetVal())
		if (result & ZEND_HASH_APPLY_REMOVE) != 0 {
			this.AssertRc1()
			this._delEl(HT_IDX_TO_HASH(idx), p)
		}
		if (result & ZEND_HASH_APPLY_STOP) != 0 {
			break
		}
	}
}
func (this *HashTable) ApplyWithArgument(apply_func ApplyFuncArgT, argument any) {
	var idx uint32
	var p *Bucket
	var result int
	for idx = 0; idx < this.GetNNumUsed(); idx++ {
		p = this.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		result = apply_func(p.GetVal(), argument)
		if (result & ZEND_HASH_APPLY_REMOVE) != 0 {
			this.AssertRc1()
			this._delEl(HT_IDX_TO_HASH(idx), p)
		}
		if (result & ZEND_HASH_APPLY_STOP) != 0 {
			break
		}
	}
}
func (this *HashTable) ApplyWithArguments(apply_func ApplyFuncArgsT, num_args int, _ ...any) {
	var idx uint32
	var p *Bucket
	var args va_list
	var hash_key ZendHashKey
	var result int
	for idx = 0; idx < this.GetNNumUsed(); idx++ {
		p = this.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		va_start(args, num_args)
		hash_key.SetH(p.GetH())
		hash_key.SetKey(p.GetKey())
		result = apply_func(p.GetVal(), num_args, args, &hash_key)
		if (result & ZEND_HASH_APPLY_REMOVE) != 0 {
			this.AssertRc1()
			this._delEl(HT_IDX_TO_HASH(idx), p)
		}
		if (result & ZEND_HASH_APPLY_STOP) != 0 {
			va_end(args)
			break
		}
		va_end(args)
	}
}
func (this *HashTable) ReverseApply(apply_func ApplyFuncT) {
	var idx uint32
	var p *Bucket
	var result int
	idx = this.GetNNumUsed()
	for idx > 0 {
		idx--
		p = this.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		result = apply_func(p.GetVal())
		if (result & ZEND_HASH_APPLY_REMOVE) != 0 {
			this.AssertRc1()
			this._delEl(HT_IDX_TO_HASH(idx), p)
		}
		if (result & ZEND_HASH_APPLY_STOP) != 0 {
			break
		}
	}
}
func (this *HashTable) Copy(source *HashTable, pCopyConstructor CopyCtorFuncT) {
	var idx uint32
	var p *Bucket
	var new_entry *Zval
	var data *Zval
	this.AssertRc1()
	for idx = 0; idx < source.GetNNumUsed(); idx++ {
		p = source.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}

		/* INDIRECT element may point to UNDEF-ined slots */

		data = p.GetVal()
		if data.IsType(IS_INDIRECT) {
			data = data.GetZv()
			if data.IsType(IS_UNDEF) {
				continue
			}
		}
		if p.GetKey() != nil {
			new_entry = this.Update(p.GetKey(), data)
		} else {
			new_entry = this.IndexUpdate(p.GetH(), data)
		}
		if pCopyConstructor != nil {
			pCopyConstructor(new_entry)
		}
	}
}
func ZendArrayDupElement(source *HashTable, target *HashTable, idx uint32, p *Bucket, q *Bucket, packed int, static_keys int, with_holes int) int {
	var data *Zval = p.GetVal()
	if with_holes != 0 {
		if packed == 0 && data.GetTypeInfo() == IS_INDIRECT {
			data = data.GetZv()
		}
		if data.GetTypeInfo() == IS_UNDEF {
			return 0
		}
	} else if packed == 0 {

		/* INDIRECT element may point to UNDEF-ined slots */

		if data.GetTypeInfo() == IS_INDIRECT {
			data = data.GetZv()
			if data.GetTypeInfo() == IS_UNDEF {
				return 0
			}
		}

		/* INDIRECT element may point to UNDEF-ined slots */

	}
	for {
		if Z_OPT_REFCOUNTED_P(data) {
			if Z_ISREF_P(data) && Z_REFCOUNT_P(data) == 1 && (Z_REFVAL_P(data).GetType() != IS_ARRAY || Z_REFVAL_P(data).GetArr() != source) {
				data = Z_REFVAL_P(data)
				if !(Z_OPT_REFCOUNTED_P(data)) {
					break
				}
			}
			Z_ADDREF_P(data)
		}
		break
	}
	ZVAL_COPY_VALUE(q.GetVal(), data)
	q.SetH(p.GetH())
	if packed != 0 {
		q.SetKey(nil)
	} else {
		var nIndex uint32
		q.SetKey(p.GetKey())
		if static_keys == 0 && q.GetKey() != nil {
			ZendStringAddref(q.GetKey())
		}
		nIndex = q.GetH() | target.GetNTableMask()
		q.GetVal().SetNext(target.Hash(nIndex))
		target.Hash(nIndex) = HT_IDX_TO_HASH(idx)
	}
	return 1
}
func ZendArrayDupPackedElements(source *HashTable, target *HashTable, with_holes int) {
	var p *Bucket = source.GetArData()
	var q *Bucket = target.GetArData()
	var end *Bucket = p + source.GetNNumUsed()
	for {
		if ZendArrayDupElement(source, target, 0, p, q, 1, 1, with_holes) == 0 {
			if with_holes != 0 {
				ZVAL_UNDEF(q.GetVal())
			}
		}
		p++
		q++
		if p == end {
			break
		}
	}
}
func ZendArrayDupElements(source *HashTable, target *HashTable, static_keys int, with_holes int) uint32 {
	var idx uint32 = 0
	var p *Bucket = source.GetArData()
	var q *Bucket = target.GetArData()
	var end *Bucket = p + source.GetNNumUsed()
	for {
		if ZendArrayDupElement(source, target, idx, p, q, 0, static_keys, with_holes) == 0 {
			var target_idx uint32 = idx
			idx++
			p++
			for p != end {
				if ZendArrayDupElement(source, target, target_idx, p, q, 0, static_keys, with_holes) != 0 {
					if source.GetNInternalPointer() == idx {
						target.SetNInternalPointer(target_idx)
					}
					target_idx++
					q++
				}
				idx++
				p++
			}
			return target_idx
		}
		idx++
		p++
		q++
		if p == end {
			break
		}
	}
	return idx
}
func ZendArrayDup(source *HashTable) *HashTable {
	var idx uint32
	var target *HashTable
	ALLOC_HASHTABLE(target)
	GC_SET_REFCOUNT(target, 1)
	GC_TYPE_INFO(target) = IS_ARRAY | GC_COLLECTABLE<<GC_FLAGS_SHIFT
	target.SetPDestructor(ZVAL_PTR_DTOR)
	if source.GetNNumOfElements() == 0 {
		target.SetUFlags(HASH_FLAG_UNINITIALIZED)
		target.SetNTableMask(HT_MIN_MASK)
		target.SetNNumUsed(0)
		target.SetNNumOfElements(0)
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNInternalPointer(0)
		target.SetNTableSize(HT_MIN_SIZE)
		target.SetDataAddr(&UninitializedBucket)
	} else if (GC_FLAGS(source) & IS_ARRAY_IMMUTABLE) != 0 {
		target.SetUFlags(source.GetUFlags() & HASH_FLAG_MASK)
		target.SetNTableMask(source.GetNTableMask())
		target.SetNNumUsed(source.GetNNumUsed())
		target.SetNNumOfElements(source.GetNNumOfElements())
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNTableSize(source.GetNTableSize())
		target.SetDataAddr(Emalloc(target.Size()))
		target.SetNInternalPointer(source.GetNInternalPointer())
		memcpy(target.GetDataAddr(), source.GetDataAddr(), source.UsedSize())
	} else if source.HasUFlags(HASH_FLAG_PACKED) {
		target.SetUFlags(source.GetUFlags() & HASH_FLAG_MASK)
		target.SetNTableMask(HT_MIN_MASK)
		target.SetNNumUsed(source.GetNNumUsed())
		target.SetNNumOfElements(source.GetNNumOfElements())
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNTableSize(source.GetNTableSize())
		target.SetDataAddr(Emalloc(HT_SIZE_EX(target.GetNTableSize(), HT_MIN_MASK)))
		if source.GetNInternalPointer() < source.GetNNumUsed() {
			target.SetNInternalPointer(source.GetNInternalPointer())
		} else {
			target.SetNInternalPointer(0)
		}
		target.HashResetPacked()
		if target.IsWithoutHoles() {
			ZendArrayDupPackedElements(source, target, 0)
		} else {
			ZendArrayDupPackedElements(source, target, 1)
		}
	} else {
		target.SetUFlags(source.GetUFlags() & HASH_FLAG_MASK)
		target.SetNTableMask(source.GetNTableMask())
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		if source.GetNInternalPointer() < source.GetNNumUsed() {
			target.SetNInternalPointer(source.GetNInternalPointer())
		} else {
			target.SetNInternalPointer(0)
		}
		target.SetNTableSize(source.GetNTableSize())
		target.SetDataAddr(Emalloc(target.Size()))
		target.HashReset()
		if target.HasStaticKeysOnly() {
			if source.IsWithoutHoles() {
				idx = ZendArrayDupElements(source, target, 1, 0)
			} else {
				idx = ZendArrayDupElements(source, target, 1, 1)
			}
		} else {
			if source.IsWithoutHoles() {
				idx = ZendArrayDupElements(source, target, 0, 0)
			} else {
				idx = ZendArrayDupElements(source, target, 0, 1)
			}
		}
		target.SetNNumUsed(idx)
		target.SetNNumOfElements(idx)
	}
	return target
}
func (this *HashTable) Merge(source *HashTable, pCopyConstructor CopyCtorFuncT, overwrite ZendBool) {
	var idx uint32
	var p *Bucket
	var t *Zval
	var s *Zval
	this.AssertRc1()
	if overwrite != 0 {
		for idx = 0; idx < source.GetNNumUsed(); idx++ {
			p = source.GetArData() + idx
			s = p.GetVal()
			if s.IsType(IS_INDIRECT) {
				s = s.GetZv()
			}
			if s.IsType(IS_UNDEF) {
				continue
			}
			if p.GetKey() != nil {
				t = this._addOrUpdateI(p.GetKey(), s, HASH_UPDATE|HASH_UPDATE_INDIRECT)
				if pCopyConstructor != nil {
					pCopyConstructor(t)
				}
			} else {
				t = this.IndexUpdate(p.GetH(), s)
				if pCopyConstructor != nil {
					pCopyConstructor(t)
				}
			}
		}
	} else {
		for idx = 0; idx < source.GetNNumUsed(); idx++ {
			p = source.GetArData() + idx
			s = p.GetVal()
			if s.IsType(IS_INDIRECT) {
				s = s.GetZv()
			}
			if s.IsType(IS_UNDEF) {
				continue
			}
			if p.GetKey() != nil {
				t = this._addOrUpdateI(p.GetKey(), s, HASH_ADD|HASH_UPDATE_INDIRECT)
				if t != nil && pCopyConstructor != nil {
					pCopyConstructor(t)
				}
			} else {
				t = this.IndexAdd(p.GetH(), s)
				if t != nil && pCopyConstructor != nil {
					pCopyConstructor(t)
				}
			}
		}
	}
}
func (this *HashTable) ReplaceCheckerWrapper(source_data *Zval, p *Bucket, pParam any, merge_checker_func MergeCheckerFuncT) ZendBool {
	var hash_key ZendHashKey
	hash_key.SetH(p.GetH())
	hash_key.SetKey(p.GetKey())
	return merge_checker_func(this, source_data, &hash_key, pParam)
}
func (this *HashTable) MergeEx(source *HashTable, pCopyConstructor CopyCtorFuncT, pMergeSource MergeCheckerFuncT, pParam any) {
	var idx uint32
	var p *Bucket
	var t *Zval
	this.AssertRc1()
	for idx = 0; idx < source.GetNNumUsed(); idx++ {
		p = source.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		if this.ReplaceCheckerWrapper(p.GetVal(), p, pParam, pMergeSource) != 0 {
			t = this.Update(p.GetKey(), p.GetVal())
			if pCopyConstructor != nil {
				pCopyConstructor(t)
			}
		}
	}
}
func (this *HashTable) Find(key *ZendString) *Zval {
	var p *Bucket
	p = this.FindBucket(key, 0)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func (this *HashTable) _findKnownHash(key *ZendString) *Zval {
	var p *Bucket
	p = this.FindBucket(key, 1)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func (this *HashTable) StrFind(str *byte, len_ int) *Zval {
	var h ZendUlong
	var p *Bucket
	h = ZendInlineHashFunc(str, len_)
	p = this.StrFindBucket(str, len_, h)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func (this *HashTable) IndexFind(h ZendUlong) *Zval {
	var p *Bucket
	if this.HasUFlags(HASH_FLAG_PACKED) {
		if h < this.GetNNumUsed() {
			p = this.GetArData() + h
			if p.GetVal().GetType() != IS_UNDEF {
				return p.GetVal()
			}
		}
		return nil
	}
	p = this.IndexFindBucket(h)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func (this *HashTable) _indexFind(h ZendUlong) *Zval {
	var p *Bucket
	p = this.IndexFindBucket(h)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func (this *HashTable) InternalPointerResetEx(pos *HashPosition) { *pos = this._getValidPos(0) }
func (this *HashTable) InternalPointerEndEx(pos *HashPosition) {
	var idx uint32
	idx = this.GetNNumUsed()
	for idx > 0 {
		idx--
		if this.GetArData()[idx].GetVal().GetType() != IS_UNDEF {
			*pos = idx
			return
		}
	}
	*pos = this.GetNNumUsed()
}
func (this *HashTable) MoveForwardEx(pos *HashPosition) int {
	var idx uint32
	idx = this._getValidPos(*pos)
	if idx < this.GetNNumUsed() {
		for true {
			idx++
			if idx >= this.GetNNumUsed() {
				*pos = this.GetNNumUsed()
				return SUCCESS
			}
			if this.GetArData()[idx].GetVal().GetType() != IS_UNDEF {
				*pos = idx
				return SUCCESS
			}
		}
	} else {
		return FAILURE
	}
}
func (this *HashTable) MoveBackwardsEx(pos *HashPosition) int {
	var idx uint32 = *pos
	if idx < this.GetNNumUsed() {
		for idx > 0 {
			idx--
			if this.GetArData()[idx].GetVal().GetType() != IS_UNDEF {
				*pos = idx
				return SUCCESS
			}
		}
		*pos = this.GetNNumUsed()
		return SUCCESS
	} else {
		return FAILURE
	}
}
func (this *HashTable) GetCurrentKeyEx(str_index **ZendString, num_index *ZendUlong, pos *HashPosition) int {
	var idx uint32
	var p *Bucket
	idx = this._getValidPos(*pos)
	if idx < this.GetNNumUsed() {
		p = this.GetArData() + idx
		if p.GetKey() != nil {
			*str_index = p.GetKey()
			return HASH_KEY_IS_STRING
		} else {
			*num_index = p.GetH()
			return HASH_KEY_IS_LONG
		}
	}
	return HASH_KEY_NON_EXISTENT
}
func (this *HashTable) GetCurrentKeyZvalEx(key *Zval, pos *HashPosition) {
	var idx uint32
	var p *Bucket
	idx = this._getValidPos(*pos)
	if idx >= this.GetNNumUsed() {
		ZVAL_NULL(key)
	} else {
		p = this.GetArData() + idx
		if p.GetKey() != nil {
			ZVAL_STR_COPY(key, p.GetKey())
		} else {
			ZVAL_LONG(key, p.GetH())
		}
	}
}
func (this *HashTable) GetCurrentKeyTypeEx(pos *HashPosition) int {
	var idx uint32
	var p *Bucket
	idx = this._getValidPos(*pos)
	if idx < this.GetNNumUsed() {
		p = this.GetArData() + idx
		if p.GetKey() != nil {
			return HASH_KEY_IS_STRING
		} else {
			return HASH_KEY_IS_LONG
		}
	}
	return HASH_KEY_NON_EXISTENT
}
func (this *HashTable) GetCurrentDataEx(pos *HashPosition) *Zval {
	var idx uint32
	var p *Bucket
	idx = this._getValidPos(*pos)
	if idx < this.GetNNumUsed() {
		p = this.GetArData() + idx
		return p.GetVal()
	} else {
		return nil
	}
}
func ZendHashBucketSwap(p *Bucket, q *Bucket) {
	var val Zval
	var h ZendUlong
	var key *ZendString
	ZVAL_COPY_VALUE(&val, p.GetVal())
	h = p.GetH()
	key = p.GetKey()
	ZVAL_COPY_VALUE(p.GetVal(), q.GetVal())
	p.SetH(q.GetH())
	p.SetKey(q.GetKey())
	ZVAL_COPY_VALUE(q.GetVal(), &val)
	q.SetH(h)
	q.SetKey(key)
}
func ZendHashBucketRenumSwap(p *Bucket, q *Bucket) {
	var val Zval
	ZVAL_COPY_VALUE(&val, p.GetVal())
	ZVAL_COPY_VALUE(p.GetVal(), q.GetVal())
	ZVAL_COPY_VALUE(q.GetVal(), &val)
}
func ZendHashBucketPackedSwap(p *Bucket, q *Bucket) {
	var val Zval
	var h ZendUlong
	ZVAL_COPY_VALUE(&val, p.GetVal())
	h = p.GetH()
	ZVAL_COPY_VALUE(p.GetVal(), q.GetVal())
	p.SetH(q.GetH())
	ZVAL_COPY_VALUE(q.GetVal(), &val)
	q.SetH(h)
}
func (this *HashTable) SortEx(sort SortFuncT, compar CompareFuncT, renumber ZendBool) int {
	var p *Bucket
	var i uint32
	var j uint32
	this.AssertRc1()
	if this.GetNNumOfElements() <= 1 && !(renumber != 0 && this.GetNNumOfElements() > 0) {
		return SUCCESS
	}
	if this.IsWithoutHoles() {
		i = this.GetNNumUsed()
	} else {
		j = 0
		i = 0
		for ; j < this.GetNNumUsed(); j++ {
			p = this.GetArData() + j
			if p.GetVal().IsType(IS_UNDEF) {
				continue
			}
			if i != j {
				this.GetArData()[i] = *p
			}
			i++
		}
	}
	sort(any(this.GetArData()), i, b.SizeOf("Bucket"), compar, swap_func_t(b.CondF2(renumber != 0, ZendHashBucketRenumSwap, func() __auto__ {
		if this.HasUFlags(HASH_FLAG_PACKED) {
			return ZendHashBucketPackedSwap
		} else {
			return ZendHashBucketSwap
		}
	})))
	this.SetNNumUsed(i)
	this.SetNInternalPointer(0)
	if renumber != 0 {
		for j = 0; j < i; j++ {
			p = this.GetArData() + j
			p.SetH(j)
			if p.GetKey() != nil {
				ZendStringRelease(p.GetKey())
				p.SetKey(nil)
			}
		}
		this.SetNNextFreeElement(i)
	}
	if this.HasUFlags(HASH_FLAG_PACKED) {
		if renumber == 0 {
			this.PackedToHash()
		}
	} else {
		if renumber != 0 {
			var new_data any
			var old_data any = this.GetDataAddr()
			var old_buckets *Bucket = this.GetArData()
			new_data = Pemalloc(HT_SIZE_EX(this.GetNTableSize(), HT_MIN_MASK), GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
			this.AddUFlags(HASH_FLAG_PACKED | HASH_FLAG_STATIC_KEYS)
			this.SetNTableMask(HT_MIN_MASK)
			this.SetDataAddr(new_data)
			memcpy(this.GetArData(), old_buckets, b.SizeOf("Bucket")*this.GetNNumUsed())
			Pefree(old_data, GC_FLAGS(this)&IS_ARRAY_PERSISTENT)
			this.HashResetPacked()
		} else {
			this.Rehash()
		}
	}
	return SUCCESS
}
func (this *HashTable) CompareImpl(ht2 *HashTable, compar CompareFuncT, ordered ZendBool) int {
	var idx1 uint32
	var idx2 uint32
	if this.GetNNumOfElements() != ht2.GetNNumOfElements() {
		if this.GetNNumOfElements() > ht2.GetNNumOfElements() {
			return 1
		} else {
			return -1
		}
	}
	idx1 = 0
	idx2 = 0
	for ; idx1 < this.GetNNumUsed(); idx1++ {
		var p1 *Bucket = this.GetArData() + idx1
		var p2 *Bucket
		var pData1 *Zval
		var pData2 *Zval
		var result int
		if p1.GetVal().IsType(IS_UNDEF) {
			continue
		}
		if ordered != 0 {
			for true {
				ZEND_ASSERT(idx2 != ht2.GetNNumUsed())
				p2 = ht2.GetArData() + idx2
				if p2.GetVal().GetType() != IS_UNDEF {
					break
				}
				idx2++
			}
			if p1.GetKey() == nil && p2.GetKey() == nil {
				if p1.GetH() != p2.GetH() {
					if p1.GetH() > p2.GetH() {
						return 1
					} else {
						return -1
					}
				}
			} else if p1.GetKey() != nil && p2.GetKey() != nil {
				if p1.GetKey().GetLen() != p2.GetKey().GetLen() {
					if p1.GetKey().GetLen() > p2.GetKey().GetLen() {
						return 1
					} else {
						return -1
					}
				}
				result = memcmp(p1.GetKey().GetVal(), p2.GetKey().GetVal(), p1.GetKey().GetLen())
				if result != 0 {
					return result
				}
			} else {

				/* Mixed key types: A string key is considered as larger */

				if p1.GetKey() != nil {
					return 1
				} else {
					return -1
				}

				/* Mixed key types: A string key is considered as larger */

			}
			pData2 = p2.GetVal()
			idx2++
		} else {
			if p1.GetKey() == nil {
				pData2 = ht2.IndexFind(p1.GetH())
				if pData2 == nil {
					return 1
				}
			} else {
				pData2 = ht2.Find(p1.GetKey())
				if pData2 == nil {
					return 1
				}
			}
		}
		pData1 = p1.GetVal()
		if pData1.IsType(IS_INDIRECT) {
			pData1 = pData1.GetZv()
		}
		if pData2.IsType(IS_INDIRECT) {
			pData2 = pData2.GetZv()
		}
		if pData1.IsType(IS_UNDEF) {
			if pData2.GetType() != IS_UNDEF {
				return -1
			}
		} else if pData2.IsType(IS_UNDEF) {
			return 1
		} else {
			result = compar(pData1, pData2)
			if result != 0 {
				return result
			}
		}
	}
	return 0
}
func (this *HashTable) Compare(ht2 *HashTable, compar CompareFuncT, ordered ZendBool) int {
	var result int
	if this == ht2 {
		return 0
	}

	/* It's enough to protect only one of the arrays.
	 * The second one may be referenced from the first and this may cause
	 * false recursion detection.
	 */

	if GC_IS_RECURSIVE(this) != 0 {
		ZendErrorNoreturn(E_ERROR, "Nesting level too deep - recursive dependency?")
	}
	if (GC_FLAGS(this) & GC_IMMUTABLE) == 0 {
		GC_PROTECT_RECURSION(this)
	}
	result = this.CompareImpl(ht2, compar, ordered)
	if (GC_FLAGS(this) & GC_IMMUTABLE) == 0 {
		GC_UNPROTECT_RECURSION(this)
	}
	return result
}
func (this *HashTable) Minmax(compar CompareFuncT, flag uint32) *Zval {
	var idx uint32
	var p *Bucket
	var res *Bucket
	if this.GetNNumOfElements() == 0 {
		return nil
	}
	idx = 0
	for true {
		if idx == this.GetNNumUsed() {
			return nil
		}
		if this.GetArData()[idx].GetVal().GetType() != IS_UNDEF {
			break
		}
		idx++
	}
	res = this.GetArData() + idx
	for ; idx < this.GetNNumUsed(); idx++ {
		p = this.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		if flag != 0 {
			if compar(res, p) < 0 {
				res = p
			}
		} else {
			if compar(res, p) > 0 {
				res = p
			}
		}
	}
	return res.GetVal()
}
func _zendHandleNumericStrEx(key *byte, length int, idx *ZendUlong) int {
	var tmp *byte = key
	var end *byte = key + length
	if (*tmp) == '-' {
		tmp++
	}
	if (*tmp) == '0' && length > 1 || end-tmp > MAX_LENGTH_OF_LONG-1 || SIZEOF_ZEND_LONG == 4 && end-tmp == MAX_LENGTH_OF_LONG-1 && (*tmp) > '2' {
		return 0
	}
	*idx = (*tmp) - '0'
	for true {
		tmp++
		if tmp == end {
			if (*key) == '-' {
				if (*idx)-1 > ZEND_LONG_MAX {
					return 0
				}
				*idx = 0 - (*idx)
			} else if (*idx) > ZEND_LONG_MAX {
				return 0
			}
			return 1
		}
		if (*tmp) <= '9' && (*tmp) >= '0' {
			*idx = (*idx)*10 + ((*tmp) - '0')
		} else {
			return 0
		}
	}
}
func ZendSymtableToProptable(ht *HashTable) *HashTable {
	var num_key ZendUlong
	var str_key *ZendString
	var zv *Zval
	if ht.IsPacked() {
		goto convert
	}
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			str_key = _p.GetKey()
			if str_key == nil {
				goto convert
			}
		}
		break
	}
	if (GC_FLAGS(ht) & IS_ARRAY_IMMUTABLE) == 0 {
		GC_ADDREF(ht)
	}
	return ht
convert:
	var new_ht *HashTable = ZendNewArray(ht.GetNNumOfElements())
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			num_key = _p.GetH()
			str_key = _p.GetKey()
			zv = _z
			if str_key == nil {
				str_key = ZendLongToStr(num_key)
				ZendStringDelref(str_key)
			}
			for {
				if Z_OPT_REFCOUNTED_P(zv) {
					if Z_ISREF_P(zv) && Z_REFCOUNT_P(zv) == 1 {
						zv = Z_REFVAL_P(zv)
						if !(Z_OPT_REFCOUNTED_P(zv)) {
							break
						}
					}
					Z_ADDREF_P(zv)
				}
				break
			}
			new_ht.Update(str_key, zv)
		}
		break
	}
	return new_ht
}
func ZendProptableToSymtable(ht *HashTable, always_duplicate ZendBool) *HashTable {
	var num_key ZendUlong
	var str_key *ZendString
	var zv *Zval
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.IsType(IS_UNDEF) {
				continue
			}
			str_key = _p.GetKey()

			/* The `str_key &&` here might seem redundant: property tables should
			 * only have string keys. Unfortunately, this isn't true, at the very
			 * least because of ArrayObject, which stores a symtable where the
			 * property table should be.
			 */

			if str_key != nil && ZEND_HANDLE_NUMERIC(str_key, num_key) != 0 {
				goto convert
			}

			/* The `str_key &&` here might seem redundant: property tables should
			 * only have string keys. Unfortunately, this isn't true, at the very
			 * least because of ArrayObject, which stores a symtable where the
			 * property table should be.
			 */

		}
		break
	}
	if always_duplicate != 0 {
		return ZendArrayDup(ht)
	}
	if (GC_FLAGS(ht) & IS_ARRAY_IMMUTABLE) == 0 {
		GC_ADDREF(ht)
	}
	return ht
convert:
	var new_ht *HashTable = ZendNewArray(ht.GetNNumOfElements())
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()
			if _z.IsType(IS_INDIRECT) {
				_z = _z.GetZv()
			}
			if _z.IsType(IS_UNDEF) {
				continue
			}
			num_key = _p.GetH()
			str_key = _p.GetKey()
			zv = _z
			for {
				if Z_OPT_REFCOUNTED_P(zv) {
					if Z_ISREF_P(zv) && Z_REFCOUNT_P(zv) == 1 {
						zv = Z_REFVAL_P(zv)
						if !(Z_OPT_REFCOUNTED_P(zv)) {
							break
						}
					}
					Z_ADDREF_P(zv)
				}
				break
			}

			/* Again, thank ArrayObject for `!str_key ||`. */

			if str_key == nil || ZEND_HANDLE_NUMERIC(str_key, num_key) != 0 {
				new_ht.IndexUpdate(num_key, zv)
			} else {
				new_ht.Update(str_key, zv)
			}

			/* Again, thank ArrayObject for `!str_key ||`. */

		}
		break
	}
	return new_ht
}
