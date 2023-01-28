// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func HT_FLAGS(ht *HashTable) uint32 { return ht.GetUFlags() }
func HT_INVALIDATE(ht *HashTable)   { ht.GetUFlags() = HASH_FLAG_UNINITIALIZED }
func HT_IS_INITIALIZED(ht *HashTable) bool {
	return (ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) == 0
}
func HT_IS_PACKED(ht *HashTable) bool {
	return (ht.GetUFlags() & HASH_FLAG_PACKED) != 0
}
func HT_IS_WITHOUT_HOLES(ht *HashTable) bool {
	return ht.GetNNumUsed() == ht.GetNNumOfElements()
}
func HT_HAS_STATIC_KEYS_ONLY(ht *HashTable) bool {
	return (ht.GetUFlags() & (HASH_FLAG_PACKED | HASH_FLAG_STATIC_KEYS)) != 0
}
func HT_ITERATORS_COUNT(ht *HashTable) ZendUchar      { return ht.GetNIteratorsCount() }
func HT_ITERATORS_OVERFLOW(ht *HashTable) bool        { return ht.GetNIteratorsCount() == 0xff }
func HT_HAS_ITERATORS(ht *HashTable) bool             { return ht.GetNIteratorsCount() != 0 }
func HT_SET_ITERATORS_COUNT(ht *HashTable, iters int) { ht.GetNIteratorsCount() = iters }
func HT_INC_ITERATORS_COUNT(ht *HashTable) {
	HT_SET_ITERATORS_COUNT(ht, ht.GetNIteratorsCount()+1)
}
func HT_DEC_ITERATORS_COUNT(ht *HashTable) {
	HT_SET_ITERATORS_COUNT(ht, ht.GetNIteratorsCount()-1)
}
func ZVAL_EMPTY_ARRAY(z *Zval) {
	var __z *Zval = z
	__z.GetArr() = (*ZendArray)(&ZendEmptyArray)
	__z.GetTypeInfo() = IS_ARRAY
}
func ZendHashInit(ht *HashTable, nSize uint32, pHashFunction __auto__, pDestructor DtorFuncT, persistent ZendBool) {
	_zendHashInit(ht, nSize, pDestructor, persistent)
}
func ZendHashInitEx(ht *HashTable, nSize uint32, pHashFunction __auto__, pDestructor DtorFuncT, persistent ZendBool, bApplyProtection int) {
	_zendHashInit(ht, nSize, pDestructor, persistent)
}
func ZendHashFindEx(ht *HashTable, key *ZendString, known_hash ZendBool) *Zval {
	if known_hash != 0 {
		return _zendHashFindKnownHash(ht, key)
	} else {
		return ZendHashFind(ht, key)
	}
}
func ZEND_HASH_INDEX_FIND(_ht *HashTable, _h ZendUlong, _ret *Zval, _not_found __auto__) {
	if (_ht.GetUFlags() & HASH_FLAG_PACKED) != 0 {
		if zend_ulong(_h) < zend_ulong(_ht).nNumUsed {
			_ret = _ht.GetArData()[_h].GetVal()
			if _ret.GetType() == IS_UNDEF {
				goto _not_found
			}
		} else {
			goto _not_found
		}
	} else {
		_ret = _zendHashIndexFind(_ht, _h)
		if _ret == nil {
			goto _not_found
		}
	}
}
func ZendHashExists(ht *HashTable, key *ZendString) ZendBool { return ZendHashFind(ht, key) != nil }
func ZendHashStrExists(ht *HashTable, str *byte, len_ int) ZendBool {
	return ZendHashStrFind(ht, str, len_) != nil
}
func ZendHashIndexExists(ht *HashTable, h ZendUlong) ZendBool { return ZendHashIndexFind(ht, h) != nil }
func ZendHashHasMoreElementsEx(ht *HashTable, pos *HashPosition) ZEND_RESULT_CODE {
	if ZendHashGetCurrentKeyTypeEx(ht, pos) == HASH_KEY_NON_EXISTENT {
		return FAILURE
	} else {
		return SUCCESS
	}
}
func ZendHashHasMoreElements(ht *HashTable) ZEND_RESULT_CODE {
	return ZendHashHasMoreElementsEx(ht, ht.GetNInternalPointer())
}
func ZendHashMoveForward(ht *HashTable) int {
	return ZendHashMoveForwardEx(ht, ht.GetNInternalPointer())
}
func ZendHashMoveBackwards(ht *HashTable) int {
	return ZendHashMoveBackwardsEx(ht, ht.GetNInternalPointer())
}
func ZendHashGetCurrentKey(ht *HashTable, str_index **ZendString, num_index *ZendUlong) int {
	return ZendHashGetCurrentKeyEx(ht, str_index, num_index, ht.GetNInternalPointer())
}
func ZendHashGetCurrentKeyZval(ht *HashTable, key *Zval) {
	ZendHashGetCurrentKeyZvalEx(ht, key, ht.GetNInternalPointer())
}
func ZendHashGetCurrentKeyType(ht *HashTable) int {
	return ZendHashGetCurrentKeyTypeEx(ht, ht.GetNInternalPointer())
}
func ZendHashGetCurrentData(ht *HashTable) *Zval {
	return ZendHashGetCurrentDataEx(ht, ht.GetNInternalPointer())
}
func ZendHashInternalPointerReset(ht *HashTable) {
	ZendHashInternalPointerResetEx(ht, ht.GetNInternalPointer())
}
func ZendHashInternalPointerEnd(ht *HashTable) {
	ZendHashInternalPointerEndEx(ht, ht.GetNInternalPointer())
}
func ZendHashSort(ht *HashTable, compare_func CompareFuncT, renumber ZendBool) int {
	return ZendHashSortEx(ht, ZendSort, compare_func, renumber)
}
func ZendHashNumElements(ht *HashTable) __auto__     { return ht.GetNNumOfElements() }
func ZendHashNextFreeElement(ht *HashTable) ZendLong { return ht.GetNNextFreeElement() }
func ZendNewArray(size uint32) *HashTable            { return _zendNewArray(size) }
func ZendHashIteratorsUpdate(ht *HashTable, from HashPosition, to HashPosition) {
	if HT_HAS_ITERATORS(ht) {
		_zendHashIteratorsUpdate(ht, from, to)
	}
}
func ZEND_INIT_SYMTABLE(ht *HashTable) { ZEND_INIT_SYMTABLE_EX(ht, 8, 0) }
func ZEND_INIT_SYMTABLE_EX(ht *HashTable, n uint32, persistent ZendBool) {
	ZendHashInit(ht, n, nil, ZVAL_PTR_DTOR, persistent)
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
func ZendHashFindInd(ht *HashTable, key *ZendString) *Zval {
	var zv *Zval
	zv = ZendHashFind(ht, key)
	if zv != nil && zv.GetType() == IS_INDIRECT {
		if zv.GetZv().GetType() != IS_UNDEF {
			return zv.GetZv()
		} else {
			return nil
		}
	} else {
		return zv
	}
}
func ZendHashFindExInd(ht *HashTable, key *ZendString, known_hash ZendBool) *Zval {
	var zv *Zval
	zv = ZendHashFindEx(ht, key, known_hash)
	if zv != nil && zv.GetType() == IS_INDIRECT {
		if zv.GetZv().GetType() != IS_UNDEF {
			return zv.GetZv()
		} else {
			return nil
		}
	} else {
		return zv
	}
}
func ZendHashExistsInd(ht *HashTable, key *ZendString) int {
	var zv *Zval
	zv = ZendHashFind(ht, key)
	return zv != nil && (zv.GetType() != IS_INDIRECT || zv.GetZv().GetType() != IS_UNDEF)
}
func ZendHashStrFindInd(ht *HashTable, str *byte, len_ int) *Zval {
	var zv *Zval
	zv = ZendHashStrFind(ht, str, len_)
	if zv != nil && zv.GetType() == IS_INDIRECT {
		if zv.GetZv().GetType() != IS_UNDEF {
			return zv.GetZv()
		} else {
			return nil
		}
	} else {
		return zv
	}
}
func ZendHashStrExistsInd(ht *HashTable, str string, len_ int) int {
	var zv *Zval
	zv = ZendHashStrFind(ht, str, len_)
	return zv != nil && (zv.GetType() != IS_INDIRECT || zv.GetZv().GetType() != IS_UNDEF)
}
func ZendSymtableAddNew(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ZendHashIndexAddNew(ht, idx, pData)
	} else {
		return ZendHashAddNew(ht, key, pData)
	}
}
func ZendSymtableUpdate(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashUpdate(ht, key, pData)
	}
}
func ZendSymtableUpdateInd(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashUpdateInd(ht, key, pData)
	}
}
func ZendSymtableDel(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ZendHashIndexDel(ht, idx)
	} else {
		return ZendHashDel(ht, key)
	}
}
func ZendSymtableDelInd(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ZendHashIndexDel(ht, idx)
	} else {
		return ZendHashDelInd(ht, key)
	}
}
func ZendSymtableFind(ht *HashTable, key *ZendString) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ZendHashIndexFind(ht, idx)
	} else {
		return ZendHashFind(ht, key)
	}
}
func ZendSymtableFindInd(ht *HashTable, key *ZendString) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ZendHashIndexFind(ht, idx)
	} else {
		return ZendHashFindInd(ht, key)
	}
}
func ZendSymtableExists(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ZendHashIndexExists(ht, idx)
	} else {
		return ZendHashExists(ht, key)
	}
}
func ZendSymtableExistsInd(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ZendHashIndexExists(ht, idx)
	} else {
		return ZendHashExistsInd(ht, key)
	}
}
func ZendSymtableStrUpdate(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashStrUpdate(ht, str, len_, pData)
	}
}
func ZendSymtableStrUpdateInd(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashStrUpdateInd(ht, str, len_, pData)
	}
}
func ZendSymtableStrDel(ht *HashTable, str *byte, len_ int) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ZendHashIndexDel(ht, idx)
	} else {
		return ZendHashStrDel(ht, str, len_)
	}
}
func ZendSymtableStrDelInd(ht *HashTable, str *byte, len_ int) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ZendHashIndexDel(ht, idx)
	} else {
		return ZendHashStrDelInd(ht, str, len_)
	}
}
func ZendSymtableStrFind(ht *HashTable, str *byte, len_ int) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ZendHashIndexFind(ht, idx)
	} else {
		return ZendHashStrFind(ht, str, len_)
	}
}
func ZendSymtableStrExists(ht *HashTable, str *byte, len_ int) int {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ZendHashIndexExists(ht, idx)
	} else {
		return ZendHashStrExists(ht, str, len_)
	}
}
func ZendHashAddPtr(ht *HashTable, key *ZendString, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendHashAdd(ht, key, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashAddNewPtr(ht *HashTable, key *ZendString, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendHashAddNew(ht, key, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashStrAddPtr(ht *HashTable, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendHashStrAdd(ht, str, len_, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashStrAddNewPtr(ht *HashTable, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendHashStrAddNew(ht, str, len_, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashUpdatePtr(ht *HashTable, key *ZendString, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendHashUpdate(ht, key, &tmp)
	return zv.GetPtr()
}
func ZendHashStrUpdatePtr(ht *HashTable, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendHashStrUpdate(ht, str, len_, &tmp)
	return zv.GetPtr()
}
func ZendHashAddMem(ht *HashTable, key *ZendString, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ZendHashAdd(ht, key, &tmp)) {
		zv.GetPtr() = Pemalloc(size, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashAddNewMem(ht *HashTable, key *ZendString, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ZendHashAddNew(ht, key, &tmp)) {
		zv.GetPtr() = Pemalloc(size, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashStrAddMem(ht *HashTable, str *byte, len_ int, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ZendHashStrAdd(ht, str, len_, &tmp)) {
		zv.GetPtr() = Pemalloc(size, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashStrAddNewMem(ht *HashTable, str *byte, len_ int, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ZendHashStrAddNew(ht, str, len_, &tmp)) {
		zv.GetPtr() = Pemalloc(size, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashUpdateMem(ht *HashTable, key *ZendString, pData any, size int) any {
	var p any
	p = Pemalloc(size, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashUpdatePtr(ht, key, p)
}
func ZendHashStrUpdateMem(ht *HashTable, str *byte, len_ int, pData any, size int) any {
	var p any
	p = Pemalloc(size, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashStrUpdatePtr(ht, str, len_, p)
}
func ZendHashIndexAddPtr(ht *HashTable, h ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendHashIndexAdd(ht, h, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexAddNewPtr(ht *HashTable, h ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendHashIndexAddNew(ht, h, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexUpdatePtr(ht *HashTable, h ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendHashIndexUpdate(ht, h, &tmp)
	return zv.GetPtr()
}
func ZendHashIndexAddMem(ht *HashTable, h ZendUlong, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ZendHashIndexAdd(ht, h, &tmp)) {
		zv.GetPtr() = Pemalloc(size, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashNextIndexInsertPtr(ht *HashTable, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendHashNextIndexInsert(ht, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexUpdateMem(ht *HashTable, h ZendUlong, pData any, size int) any {
	var p any
	p = Pemalloc(size, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashIndexUpdatePtr(ht, h, p)
}
func ZendHashNextIndexInsertMem(ht *HashTable, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ZendHashNextIndexInsert(ht, &tmp)) {
		zv.GetPtr() = Pemalloc(size, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashFindPtr(ht *HashTable, key *ZendString) any {
	var zv *Zval
	zv = ZendHashFind(ht, key)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashFindExPtr(ht *HashTable, key *ZendString, known_hash ZendBool) any {
	var zv *Zval
	zv = ZendHashFindEx(ht, key, known_hash)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashStrFindPtr(ht *HashTable, str string, len_ int) any {
	var zv *Zval
	zv = ZendHashStrFind(ht, str, len_)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexFindPtr(ht *HashTable, h ZendUlong) any {
	var zv *Zval
	zv = ZendHashIndexFind(ht, h)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexFindDeref(ht *HashTable, h ZendUlong) *Zval {
	var zv *Zval = ZendHashIndexFind(ht, h)
	if zv != nil {
		ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashFindDeref(ht *HashTable, str *ZendString) *Zval {
	var zv *Zval = ZendHashFind(ht, str)
	if zv != nil {
		ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashStrFindDeref(ht *HashTable, str string, len_ int) *Zval {
	var zv *Zval = ZendHashStrFind(ht, str, len_)
	if zv != nil {
		ZVAL_DEREF(zv)
	}
	return zv
}
func ZendSymtableStrFindPtr(ht *HashTable, str *byte, len_ int) any {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC_STR(str, len_, idx) != 0 {
		return ZendHashIndexFindPtr(ht, idx)
	} else {
		return ZendHashStrFindPtr(ht, str, len_)
	}
}
func ZendHashGetCurrentDataPtrEx(ht *HashTable, pos *HashPosition) any {
	var zv *Zval
	zv = ZendHashGetCurrentDataEx(ht, pos)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashGetCurrentDataPtr(ht *HashTable) any {
	return ZendHashGetCurrentDataPtrEx(ht, ht.GetNInternalPointer())
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
func _zendHashAppendEx(ht *HashTable, key *ZendString, zv *Zval, interned int) *Zval {
	var idx uint32 = b.PostInc(&(ht.GetNNumUsed()))
	var nIndex uint32
	var p *Bucket = ht.GetArData() + idx
	ZVAL_COPY_VALUE(p.GetVal(), zv)
	if interned == 0 {
		ht.GetUFlags() &= ^HASH_FLAG_STATIC_KEYS
		ZendStringAddref(key)
		ZendStringHashVal(key)
	}
	p.SetKey(key)
	p.SetH(key.GetH())
	nIndex = uint32(p.GetH() | ht.GetNTableMask())
	p.GetVal().GetNext() = HT_HASH(ht, nIndex)
	HT_HASH(ht, nIndex) = HT_IDX_TO_HASH(idx)
	ht.GetNNumOfElements()++
	return p.GetVal()
}
func _zendHashAppend(ht *HashTable, key *ZendString, zv *Zval) *Zval {
	return _zendHashAppendEx(ht, key, zv, 0)
}
func _zendHashAppendPtrEx(ht *HashTable, key *ZendString, ptr any, interned int) *Zval {
	var idx uint32 = b.PostInc(&(ht.GetNNumUsed()))
	var nIndex uint32
	var p *Bucket = ht.GetArData() + idx
	ZVAL_PTR(p.GetVal(), ptr)
	if interned == 0 {
		ht.GetUFlags() &= ^HASH_FLAG_STATIC_KEYS
		ZendStringAddref(key)
		ZendStringHashVal(key)
	}
	p.SetKey(key)
	p.SetH(key.GetH())
	nIndex = uint32(p.GetH() | ht.GetNTableMask())
	p.GetVal().GetNext() = HT_HASH(ht, nIndex)
	HT_HASH(ht, nIndex) = HT_IDX_TO_HASH(idx)
	ht.GetNNumOfElements()++
	return p.GetVal()
}
func _zendHashAppendPtr(ht *HashTable, key *ZendString, ptr any) *Zval {
	return _zendHashAppendPtrEx(ht, key, ptr, 0)
}
func _zendHashAppendInd(ht *HashTable, key *ZendString, ptr *Zval) {
	var idx uint32 = b.PostInc(&(ht.GetNNumUsed()))
	var nIndex uint32
	var p *Bucket = ht.GetArData() + idx
	ZVAL_INDIRECT(p.GetVal(), ptr)
	ht.GetUFlags() &= ^HASH_FLAG_STATIC_KEYS
	ZendStringAddref(key)
	ZendStringHashVal(key)
	p.SetKey(key)
	p.SetH(key.GetH())
	nIndex = uint32(p.GetH() | ht.GetNTableMask())
	p.GetVal().GetNext() = HT_HASH(ht, nIndex)
	HT_HASH(ht, nIndex) = HT_IDX_TO_HASH(idx)
	ht.GetNNumOfElements()++
}
func HT_ASSERT_RC1(ht *HashTable) {}
func ZEND_HASH_IF_FULL_DO_RESIZE(ht *HashTable) {
	if ht.GetNNumUsed() >= ht.GetNTableSize() {
		ZendHashDoResize(ht)
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
func ZendHashRealInitPackedEx(ht *HashTable) {
	var data any
	if (GC_FLAGS(ht) & IS_ARRAY_PERSISTENT) != 0 {
		data = Pemalloc(HT_SIZE_EX(ht.GetNTableSize(), HT_MIN_MASK), 1)
	} else if ht.GetNTableSize() == HT_MIN_SIZE {
		data = Emalloc(HT_SIZE_EX(HT_MIN_SIZE, HT_MIN_MASK))
	} else {
		data = Emalloc(HT_SIZE_EX(ht.GetNTableSize(), HT_MIN_MASK))
	}
	HT_SET_DATA_ADDR(ht, data)

	/* Don't overwrite iterator count. */

	ht.SetFlags(HASH_FLAG_PACKED | HASH_FLAG_STATIC_KEYS)
	HT_HASH_RESET_PACKED(ht)
}
func ZendHashRealInitMixedEx(ht *HashTable) {
	var data any
	var nSize uint32 = ht.GetNTableSize()
	if (GC_FLAGS(ht) & IS_ARRAY_PERSISTENT) != 0 {
		data = Pemalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)), 1)
	} else if nSize == HT_MIN_SIZE {
		data = Emalloc(HT_SIZE_EX(HT_MIN_SIZE, HT_SIZE_TO_MASK(HT_MIN_SIZE)))
		ht.SetNTableMask(HT_SIZE_TO_MASK(HT_MIN_SIZE))
		HT_SET_DATA_ADDR(ht, data)

		/* Don't overwrite iterator count. */

		ht.SetFlags(HASH_FLAG_STATIC_KEYS)
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
	ht.SetNTableMask(HT_SIZE_TO_MASK(nSize))
	HT_SET_DATA_ADDR(ht, data)
	ht.GetUFlags() = HASH_FLAG_STATIC_KEYS
	HT_HASH_RESET(ht)
}
func ZendHashRealInitEx(ht *HashTable, packed int) {
	HT_ASSERT_RC1(ht)
	ZEND_ASSERT((ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) != 0)
	if packed != 0 {
		ZendHashRealInitPackedEx(ht)
	} else {
		ZendHashRealInitMixedEx(ht)
	}
}
func _zendHashInitInt(ht *HashTable, nSize uint32, pDestructor DtorFuncT, persistent ZendBool) {
	GC_SET_REFCOUNT(ht, 1)
	GC_TYPE_INFO(ht) = IS_ARRAY | b.Cond(persistent != 0, GC_PERSISTENT<<GC_FLAGS_SHIFT, GC_COLLECTABLE<<GC_FLAGS_SHIFT)
	ht.GetUFlags() = HASH_FLAG_UNINITIALIZED
	ht.SetNTableMask(HT_MIN_MASK)
	HT_SET_DATA_ADDR(ht, &UninitializedBucket)
	ht.SetNNumUsed(0)
	ht.SetNNumOfElements(0)
	ht.SetNInternalPointer(0)
	ht.SetNNextFreeElement(0)
	ht.SetPDestructor(pDestructor)
	ht.SetNTableSize(ZendHashCheckSize(nSize))
}
func _zendHashInit(ht *HashTable, nSize uint32, pDestructor DtorFuncT, persistent ZendBool) {
	_zendHashInitInt(ht, nSize, pDestructor, persistent)
}
func _zendNewArray0() *HashTable {
	var ht *HashTable = Emalloc(b.SizeOf("HashTable"))
	_zendHashInitInt(ht, HT_MIN_SIZE, ZVAL_PTR_DTOR, 0)
	return ht
}
func _zendNewArray(nSize uint32) *HashTable {
	var ht *HashTable = Emalloc(b.SizeOf("HashTable"))
	_zendHashInitInt(ht, nSize, ZVAL_PTR_DTOR, 0)
	return ht
}
func ZendNewPair(val1 *Zval, val2 *Zval) *HashTable {
	var p *Bucket
	var ht *HashTable = Emalloc(b.SizeOf("HashTable"))
	_zendHashInitInt(ht, HT_MIN_SIZE, ZVAL_PTR_DTOR, 0)
	ht.SetNNextFreeElement(2)
	ht.SetNNumOfElements(ht.GetNNextFreeElement())
	ht.SetNNumUsed(ht.GetNNumOfElements())
	ZendHashRealInitPackedEx(ht)
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
func ZendHashPackedGrow(ht *HashTable) {
	HT_ASSERT_RC1(ht)
	if ht.GetNTableSize() >= HT_MAX_SIZE {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (%u * %zu + %zu)", ht.GetNTableSize()*2, b.SizeOf("Bucket"), b.SizeOf("Bucket"))
	}
	ht.SetNTableSize(ht.GetNTableSize() + ht.GetNTableSize())
	HT_SET_DATA_ADDR(ht, Perealloc2(HT_GET_DATA_ADDR(ht), HT_SIZE_EX(ht.GetNTableSize(), HT_MIN_MASK), HT_USED_SIZE(ht), GC_FLAGS(ht)&IS_ARRAY_PERSISTENT))
}
func ZendHashRealInit(ht *HashTable, packed ZendBool) {
	HT_ASSERT_RC1(ht)
	ZendHashRealInitEx(ht, packed)
}
func ZendHashRealInitPacked(ht *HashTable) {
	HT_ASSERT_RC1(ht)
	ZendHashRealInitPackedEx(ht)
}
func ZendHashRealInitMixed(ht *HashTable) {
	HT_ASSERT_RC1(ht)
	ZendHashRealInitMixedEx(ht)
}
func ZendHashPackedToHash(ht *HashTable) {
	var new_data any
	var old_data any = HT_GET_DATA_ADDR(ht)
	var old_buckets *Bucket = ht.GetArData()
	var nSize uint32 = ht.GetNTableSize()
	HT_ASSERT_RC1(ht)
	ht.GetUFlags() &= ^HASH_FLAG_PACKED
	new_data = Pemalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)), GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
	ht.SetNTableMask(HT_SIZE_TO_MASK(ht.GetNTableSize()))
	HT_SET_DATA_ADDR(ht, new_data)
	memcpy(ht.GetArData(), old_buckets, b.SizeOf("Bucket")*ht.GetNNumUsed())
	Pefree(old_data, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
	ZendHashRehash(ht)
}
func ZendHashToPacked(ht *HashTable) {
	var new_data any
	var old_data any = HT_GET_DATA_ADDR(ht)
	var old_buckets *Bucket = ht.GetArData()
	HT_ASSERT_RC1(ht)
	new_data = Pemalloc(HT_SIZE_EX(ht.GetNTableSize(), HT_MIN_MASK), GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
	ht.GetUFlags() |= HASH_FLAG_PACKED | HASH_FLAG_STATIC_KEYS
	ht.SetNTableMask(HT_MIN_MASK)
	HT_SET_DATA_ADDR(ht, new_data)
	HT_HASH_RESET_PACKED(ht)
	memcpy(ht.GetArData(), old_buckets, b.SizeOf("Bucket")*ht.GetNNumUsed())
	Pefree(old_data, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
}
func ZendHashExtend(ht *HashTable, nSize uint32, packed ZendBool) {
	HT_ASSERT_RC1(ht)
	if nSize == 0 {
		return
	}
	if (ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) != 0 {
		if nSize > ht.GetNTableSize() {
			ht.SetNTableSize(ZendHashCheckSize(nSize))
		}
		ZendHashRealInit(ht, packed)
	} else {
		if packed != 0 {
			ZEND_ASSERT((ht.GetUFlags() & HASH_FLAG_PACKED) != 0)
			if nSize > ht.GetNTableSize() {
				ht.SetNTableSize(ZendHashCheckSize(nSize))
				HT_SET_DATA_ADDR(ht, Perealloc2(HT_GET_DATA_ADDR(ht), HT_SIZE_EX(ht.GetNTableSize(), HT_MIN_MASK), HT_USED_SIZE(ht), GC_FLAGS(ht)&IS_ARRAY_PERSISTENT))
			}
		} else {
			ZEND_ASSERT((ht.GetUFlags() & HASH_FLAG_PACKED) == 0)
			if nSize > ht.GetNTableSize() {
				var new_data any
				var old_data any = HT_GET_DATA_ADDR(ht)
				var old_buckets *Bucket = ht.GetArData()
				nSize = ZendHashCheckSize(nSize)
				ht.SetNTableSize(nSize)
				new_data = Pemalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)), GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
				ht.SetNTableMask(HT_SIZE_TO_MASK(ht.GetNTableSize()))
				HT_SET_DATA_ADDR(ht, new_data)
				memcpy(ht.GetArData(), old_buckets, b.SizeOf("Bucket")*ht.GetNNumUsed())
				Pefree(old_data, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
				ZendHashRehash(ht)
			}
		}
	}
}
func ZendHashDiscard(ht *HashTable, nNumUsed uint32) {
	var p *Bucket
	var end *Bucket
	var arData *Bucket
	var nIndex uint32
	arData = ht.GetArData()
	p = arData + ht.GetNNumUsed()
	end = arData + nNumUsed
	ht.SetNNumUsed(nNumUsed)
	for p != end {
		p--
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		ht.GetNNumOfElements()--

		/* Collision pointers always directed from higher to lower buckets */

		nIndex = p.GetH() | ht.GetNTableMask()
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

			if _z.GetType() == IS_UNDEF {
				continue
			}
			val = _z
			if val.GetType() == IS_INDIRECT {
				if val.GetZv().GetType() == IS_UNDEF {
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
	if (ht.GetUFlags() & HASH_FLAG_HAS_EMPTY_IND) != 0 {
		num = ZendArrayRecalcElements(ht)
		if ht.GetNNumOfElements() == num {
			ht.GetUFlags() &= ^HASH_FLAG_HAS_EMPTY_IND
		}
	} else if ht == &(ExecutorGlobals.GetSymbolTable()) {
		num = ZendArrayRecalcElements(ht)
	} else {
		num = ht.GetNNumOfElements()
	}
	return num
}
func _zendHashGetValidPos(ht *HashTable, pos HashPosition) HashPosition {
	for pos < ht.GetNNumUsed() && Z_ISUNDEF(ht.GetArData()[pos].GetVal()) {
		pos++
	}
	return pos
}
func _zendHashGetCurrentPos(ht *HashTable) HashPosition {
	return _zendHashGetValidPos(ht, ht.GetNInternalPointer())
}
func ZendHashGetCurrentPos(ht *HashTable) HashPosition { return _zendHashGetCurrentPos(ht) }
func ZendHashIteratorAdd(ht *HashTable, pos HashPosition) uint32 {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators()
	var end *HashTableIterator = iter + ExecutorGlobals.GetHtIteratorsCount()
	var idx uint32
	if !(HT_ITERATORS_OVERFLOW(ht)) {
		HT_INC_ITERATORS_COUNT(ht)
	}
	for iter != end {
		if iter.GetHt() == nil {
			iter.SetHt(ht)
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
	iter.SetHt(ht)
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
		if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(HT_ITERATORS_OVERFLOW(iter.GetHt())) {
			HT_DEC_ITERATORS_COUNT(iter.GetHt())
		}
		if !(HT_ITERATORS_OVERFLOW(ht)) {
			HT_INC_ITERATORS_COUNT(ht)
		}
		iter.SetHt(ht)
		iter.SetPos(_zendHashGetCurrentPos(ht))
	}
	return iter.GetPos()
}
func ZendHashIteratorPosEx(idx uint32, array *Zval) HashPosition {
	var ht *HashTable = Z_ARRVAL_P(array)
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators() + idx
	ZEND_ASSERT(idx != uint32-1)
	if iter.GetHt() != ht {
		if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(HT_ITERATORS_OVERFLOW(ht)) {
			HT_DEC_ITERATORS_COUNT(iter.GetHt())
		}
		SEPARATE_ARRAY(array)
		ht = Z_ARRVAL_P(array)
		if !(HT_ITERATORS_OVERFLOW(ht)) {
			HT_INC_ITERATORS_COUNT(ht)
		}
		iter.SetHt(ht)
		iter.SetPos(_zendHashGetCurrentPos(ht))
	}
	return iter.GetPos()
}
func ZendHashIteratorDel(idx uint32) {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators() + idx
	ZEND_ASSERT(idx != uint32-1)
	if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(HT_ITERATORS_OVERFLOW(iter.GetHt())) {
		ZEND_ASSERT(iter.GetHt().GetNIteratorsCount() != 0)
		HT_DEC_ITERATORS_COUNT(iter.GetHt())
	}
	iter.SetHt(nil)
	if idx == ExecutorGlobals.GetHtIteratorsUsed()-1 {
		for idx > 0 && ExecutorGlobals.GetHtIterators()[idx-1].GetHt() == nil {
			idx--
		}
		ExecutorGlobals.SetHtIteratorsUsed(idx)
	}
}
func _zendHashIteratorsRemove(ht *HashTable) {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators()
	var end *HashTableIterator = iter + ExecutorGlobals.GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht {
			iter.SetHt(HT_POISONED_PTR)
		}
		iter++
	}
}
func ZendHashIteratorsRemove(ht *HashTable) {
	if HT_HAS_ITERATORS(ht) {
		_zendHashIteratorsRemove(ht)
	}
}
func ZendHashIteratorsLowerPos(ht *HashTable, start HashPosition) HashPosition {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators()
	var end *HashTableIterator = iter + ExecutorGlobals.GetHtIteratorsUsed()
	var res HashPosition = ht.GetNNumUsed()
	for iter != end {
		if iter.GetHt() == ht {
			if iter.GetPos() >= start && iter.GetPos() < res {
				res = iter.GetPos()
			}
		}
		iter++
	}
	return res
}
func _zendHashIteratorsUpdate(ht *HashTable, from HashPosition, to HashPosition) {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators()
	var end *HashTableIterator = iter + ExecutorGlobals.GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht && iter.GetPos() == from {
			iter.SetPos(to)
		}
		iter++
	}
}
func ZendHashIteratorsAdvance(ht *HashTable, step HashPosition) {
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators()
	var end *HashTableIterator = iter + ExecutorGlobals.GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht {
			iter.SetPos(iter.GetPos() + step)
		}
		iter++
	}
}
func ZendHashFindBucket(ht *HashTable, key *ZendString, known_hash ZendBool) *Bucket {
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
	arData = ht.GetArData()
	nIndex = h | ht.GetNTableMask()
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
func ZendHashStrFindBucket(ht *HashTable, str *byte, len_ int, h ZendUlong) *Bucket {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var arData *Bucket
	arData = ht.GetArData()
	nIndex = h | ht.GetNTableMask()
	idx = HT_HASH_EX(arData, nIndex)
	for idx != HT_INVALID_IDX {
		ZEND_ASSERT(idx < HT_IDX_TO_HASH(ht.GetNTableSize()))
		p = HT_HASH_TO_BUCKET_EX(arData, idx)
		if p.GetH() == h && p.GetKey() != nil && p.GetKey().GetLen() == len_ && !(memcmp(p.GetKey().GetVal(), str, len_)) {
			return p
		}
		idx = p.GetVal().GetNext()
	}
	return nil
}
func ZendHashIndexFindBucket(ht *HashTable, h ZendUlong) *Bucket {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var arData *Bucket
	arData = ht.GetArData()
	nIndex = h | ht.GetNTableMask()
	idx = HT_HASH_EX(arData, nIndex)
	for idx != HT_INVALID_IDX {
		ZEND_ASSERT(idx < HT_IDX_TO_HASH(ht.GetNTableSize()))
		p = HT_HASH_TO_BUCKET_EX(arData, idx)
		if p.GetH() == h && p.GetKey() == nil {
			return p
		}
		idx = p.GetVal().GetNext()
	}
	return nil
}
func _zendHashAddOrUpdateI(ht *HashTable, key *ZendString, pData *Zval, flag uint32) *Zval {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var arData *Bucket
	HT_ASSERT_RC1(ht)
	if (ht.GetUFlags() & (HASH_FLAG_UNINITIALIZED | HASH_FLAG_PACKED)) != 0 {
		if (ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) != 0 {
			ZendHashRealInitMixed(ht)
			ZendStringAddref(key)
			ht.GetUFlags() &= ^HASH_FLAG_STATIC_KEYS
			ZendStringHashVal(key)
			goto add_to_hash
		} else {
			ZendHashPackedToHash(ht)
			ZendStringAddref(key)
			ht.GetUFlags() &= ^HASH_FLAG_STATIC_KEYS
			ZendStringHashVal(key)
		}
	} else if (flag&HASH_ADD_NEW) == 0 || core.ZEND_DEBUG != 0 {
		p = ZendHashFindBucket(ht, key, 0)
		if p != nil {
			var data *Zval
			ZEND_ASSERT((flag & HASH_ADD_NEW) == 0)
			if (flag & HASH_ADD) != 0 {
				if (flag & HASH_UPDATE_INDIRECT) == 0 {
					return nil
				}
				ZEND_ASSERT(p.GetVal() != pData)
				data = p.GetVal()
				if data.GetType() == IS_INDIRECT {
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
				if (flag&HASH_UPDATE_INDIRECT) != 0 && data.GetType() == IS_INDIRECT {
					data = data.GetZv()
				}
			}
			if ht.GetPDestructor() != nil {
				ht.GetPDestructor()(data)
			}
			ZVAL_COPY_VALUE(data, pData)
			return data
		}
		ZendStringAddref(key)
		ht.GetUFlags() &= ^HASH_FLAG_STATIC_KEYS
	} else {
		ZendStringAddref(key)
		ht.GetUFlags() &= ^HASH_FLAG_STATIC_KEYS
		ZendStringHashVal(key)
	}
	ZEND_HASH_IF_FULL_DO_RESIZE(ht)
add_to_hash:
	ht.GetNNumUsed()++
	idx = ht.GetNNumUsed() - 1
	ht.GetNNumOfElements()++
	arData = ht.GetArData()
	p = arData + idx
	p.SetKey(key)
	h = key.GetH()
	p.SetH(h)
	nIndex = h | ht.GetNTableMask()
	p.GetVal().GetNext() = HT_HASH_EX(arData, nIndex)
	HT_HASH_EX(arData, nIndex) = HT_IDX_TO_HASH(idx)
	ZVAL_COPY_VALUE(p.GetVal(), pData)
	return p.GetVal()
}
func _zendHashStrAddOrUpdateI(ht *HashTable, str *byte, len_ int, h ZendUlong, pData *Zval, flag uint32) *Zval {
	var key *ZendString
	var nIndex uint32
	var idx uint32
	var p *Bucket
	HT_ASSERT_RC1(ht)
	if (ht.GetUFlags() & (HASH_FLAG_UNINITIALIZED | HASH_FLAG_PACKED)) != 0 {
		if (ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) != 0 {
			ZendHashRealInitMixed(ht)
			goto add_to_hash
		} else {
			ZendHashPackedToHash(ht)
		}
	} else if (flag & HASH_ADD_NEW) == 0 {
		p = ZendHashStrFindBucket(ht, str, len_, h)
		if p != nil {
			var data *Zval
			if (flag & HASH_ADD) != 0 {
				if (flag & HASH_UPDATE_INDIRECT) == 0 {
					return nil
				}
				ZEND_ASSERT(p.GetVal() != pData)
				data = p.GetVal()
				if data.GetType() == IS_INDIRECT {
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
				if (flag&HASH_UPDATE_INDIRECT) != 0 && data.GetType() == IS_INDIRECT {
					data = data.GetZv()
				}
			}
			if ht.GetPDestructor() != nil {
				ht.GetPDestructor()(data)
			}
			ZVAL_COPY_VALUE(data, pData)
			return data
		}
	}
	ZEND_HASH_IF_FULL_DO_RESIZE(ht)
add_to_hash:
	ht.GetNNumUsed()++
	idx = ht.GetNNumUsed() - 1
	ht.GetNNumOfElements()++
	p = ht.GetArData() + idx
	key = ZendStringInit(str, len_, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
	p.SetKey(key)
	key.GetH() = h
	p.SetH(key.GetH())
	ht.GetUFlags() &= ^HASH_FLAG_STATIC_KEYS
	ZVAL_COPY_VALUE(p.GetVal(), pData)
	nIndex = h | ht.GetNTableMask()
	p.GetVal().GetNext() = HT_HASH(ht, nIndex)
	HT_HASH(ht, nIndex) = HT_IDX_TO_HASH(idx)
	return p.GetVal()
}
func ZendHashAddOrUpdate(ht *HashTable, key *ZendString, pData *Zval, flag uint32) *Zval {
	if flag == HASH_ADD {
		return ZendHashAdd(ht, key, pData)
	} else if flag == HASH_ADD_NEW {
		return ZendHashAddNew(ht, key, pData)
	} else if flag == HASH_UPDATE {
		return ZendHashUpdate(ht, key, pData)
	} else {
		ZEND_ASSERT(flag == (HASH_UPDATE | HASH_UPDATE_INDIRECT))
		return ZendHashUpdateInd(ht, key, pData)
	}
}
func ZendHashAdd(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return _zendHashAddOrUpdateI(ht, key, pData, HASH_ADD)
}
func ZendHashUpdate(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return _zendHashAddOrUpdateI(ht, key, pData, HASH_UPDATE)
}
func ZendHashUpdateInd(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return _zendHashAddOrUpdateI(ht, key, pData, HASH_UPDATE|HASH_UPDATE_INDIRECT)
}
func ZendHashAddNew(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return _zendHashAddOrUpdateI(ht, key, pData, HASH_ADD_NEW)
}
func ZendHashStrAddOrUpdate(ht *HashTable, str *byte, len_ int, pData *Zval, flag uint32) *Zval {
	if flag == HASH_ADD {
		return ZendHashStrAdd(ht, str, len_, pData)
	} else if flag == HASH_ADD_NEW {
		return ZendHashStrAddNew(ht, str, len_, pData)
	} else if flag == HASH_UPDATE {
		return ZendHashStrUpdate(ht, str, len_, pData)
	} else {
		ZEND_ASSERT(flag == (HASH_UPDATE | HASH_UPDATE_INDIRECT))
		return ZendHashStrUpdateInd(ht, str, len_, pData)
	}
}
func ZendHashStrUpdate(ht *HashTable, str string, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return _zendHashStrAddOrUpdateI(ht, str, len_, h, pData, HASH_UPDATE)
}
func ZendHashStrUpdateInd(ht *HashTable, str string, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return _zendHashStrAddOrUpdateI(ht, str, len_, h, pData, HASH_UPDATE|HASH_UPDATE_INDIRECT)
}
func ZendHashStrAdd(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return _zendHashStrAddOrUpdateI(ht, str, len_, h, pData, HASH_ADD)
}
func ZendHashStrAddNew(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return _zendHashStrAddOrUpdateI(ht, str, len_, h, pData, HASH_ADD_NEW)
}
func ZendHashIndexAddEmptyElement(ht *HashTable, h ZendUlong) *Zval {
	var dummy Zval
	ZVAL_NULL(&dummy)
	return ZendHashIndexAdd(ht, h, &dummy)
}
func ZendHashAddEmptyElement(ht *HashTable, key *ZendString) *Zval {
	var dummy Zval
	ZVAL_NULL(&dummy)
	return ZendHashAdd(ht, key, &dummy)
}
func ZendHashStrAddEmptyElement(ht *HashTable, str *byte, len_ int) *Zval {
	var dummy Zval
	ZVAL_NULL(&dummy)
	return ZendHashStrAdd(ht, str, len_, &dummy)
}
func _zendHashIndexAddOrUpdateI(ht *HashTable, h ZendUlong, pData *Zval, flag uint32) *Zval {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	HT_ASSERT_RC1(ht)
	if (ht.GetUFlags() & HASH_FLAG_PACKED) != 0 {
		if h < ht.GetNNumUsed() {
			p = ht.GetArData() + h
			if p.GetVal().GetType() != IS_UNDEF {
			replace:
				if (flag & HASH_ADD) != 0 {
					return nil
				}
				if ht.GetPDestructor() != nil {
					ht.GetPDestructor()(p.GetVal())
				}
				ZVAL_COPY_VALUE(p.GetVal(), pData)
				return p.GetVal()
			} else {
				goto convert_to_hash
			}
		} else if h < ht.GetNTableSize() {
		add_to_packed:
			p = ht.GetArData() + h

			/* incremental initialization of empty Buckets */

			if (flag & (HASH_ADD_NEW | HASH_ADD_NEXT)) != (HASH_ADD_NEW | HASH_ADD_NEXT) {
				if h > ht.GetNNumUsed() {
					var q *Bucket = ht.GetArData() + ht.GetNNumUsed()
					for q != p {
						ZVAL_UNDEF(q.GetVal())
						q++
					}
				}
			}
			ht.SetNNumUsed(h + 1)
			ht.SetNNextFreeElement(ht.GetNNumUsed())
			goto add
		} else if h>>1 < ht.GetNTableSize() && ht.GetNTableSize()>>1 < ht.GetNNumOfElements() {
			ZendHashPackedGrow(ht)
			goto add_to_packed
		} else {
			if ht.GetNNumUsed() >= ht.GetNTableSize() {
				ht.SetNTableSize(ht.GetNTableSize() + ht.GetNTableSize())
			}
		convert_to_hash:
			ZendHashPackedToHash(ht)
		}
	} else if (ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) != 0 {
		if h < ht.GetNTableSize() {
			ZendHashRealInitPackedEx(ht)
			goto add_to_packed
		}
		ZendHashRealInitMixed(ht)
	} else {
		if (flag&HASH_ADD_NEW) == 0 || core.ZEND_DEBUG != 0 {
			p = ZendHashIndexFindBucket(ht, h)
			if p != nil {
				ZEND_ASSERT((flag & HASH_ADD_NEW) == 0)
				goto replace
			}
		}
		ZEND_HASH_IF_FULL_DO_RESIZE(ht)
	}
	ht.GetNNumUsed()++
	idx = ht.GetNNumUsed() - 1
	nIndex = h | ht.GetNTableMask()
	p = ht.GetArData() + idx
	p.GetVal().GetNext() = HT_HASH(ht, nIndex)
	HT_HASH(ht, nIndex) = HT_IDX_TO_HASH(idx)
	if ZendLong(h >= ZendLong(ht.GetNNextFreeElement())) != 0 {
		if h < ZEND_LONG_MAX {
			ht.SetNNextFreeElement(h + 1)
		} else {
			ht.SetNNextFreeElement(ZEND_LONG_MAX)
		}
	}
add:
	ht.GetNNumOfElements()++
	p.SetH(h)
	p.SetKey(nil)
	ZVAL_COPY_VALUE(p.GetVal(), pData)
	return p.GetVal()
}
func ZendHashIndexAddOrUpdate(ht *HashTable, h ZendUlong, pData *Zval, flag uint32) *Zval {
	if flag == HASH_ADD {
		return ZendHashIndexAdd(ht, h, pData)
	} else if flag == (HASH_ADD | HASH_ADD_NEW) {
		return ZendHashIndexAddNew(ht, h, pData)
	} else if flag == (HASH_ADD | HASH_ADD_NEXT) {
		ZEND_ASSERT(h == ht.GetNNextFreeElement())
		return ZendHashNextIndexInsert(ht, pData)
	} else if flag == (HASH_ADD | HASH_ADD_NEW | HASH_ADD_NEXT) {
		ZEND_ASSERT(h == ht.GetNNextFreeElement())
		return ZendHashNextIndexInsertNew(ht, pData)
	} else {
		ZEND_ASSERT(flag == HASH_UPDATE)
		return ZendHashIndexUpdate(ht, h, pData)
	}
}
func ZendHashIndexAdd(ht *HashTable, h ZendUlong, pData *Zval) *Zval {
	return _zendHashIndexAddOrUpdateI(ht, h, pData, HASH_ADD)
}
func ZendHashIndexAddNew(ht *HashTable, h ZendUlong, pData *Zval) *Zval {
	return _zendHashIndexAddOrUpdateI(ht, h, pData, HASH_ADD|HASH_ADD_NEW)
}
func ZendHashIndexUpdate(ht *HashTable, h ZendUlong, pData *Zval) *Zval {
	return _zendHashIndexAddOrUpdateI(ht, h, pData, HASH_UPDATE)
}
func ZendHashNextIndexInsert(ht *HashTable, pData *Zval) *Zval {
	return _zendHashIndexAddOrUpdateI(ht, ht.GetNNextFreeElement(), pData, HASH_ADD|HASH_ADD_NEXT)
}
func ZendHashNextIndexInsertNew(ht *HashTable, pData *Zval) *Zval {
	return _zendHashIndexAddOrUpdateI(ht, ht.GetNNextFreeElement(), pData, HASH_ADD|HASH_ADD_NEW|HASH_ADD_NEXT)
}
func ZendHashSetBucketKey(ht *HashTable, b *Bucket, key *ZendString) *Zval {
	var nIndex uint32
	var idx uint32
	var i uint32
	var p *Bucket
	var arData *Bucket
	HT_ASSERT_RC1(ht)
	ZEND_ASSERT((ht.GetUFlags() & HASH_FLAG_PACKED) == 0)
	p = ZendHashFindBucket(ht, key, 0)
	if p != nil {
		if p == b {
			return p.GetVal()
		} else {
			return nil
		}
	}
	ZendStringAddref(key)
	ht.GetUFlags() &= ^HASH_FLAG_STATIC_KEYS
	arData = ht.GetArData()

	/* del from hash */

	idx = HT_IDX_TO_HASH(b - arData)
	nIndex = b.GetH() | ht.GetNTableMask()
	i = HT_HASH_EX(arData, nIndex)
	if i == idx {
		HT_HASH_EX(arData, nIndex) = b.GetVal().GetNext()
	} else {
		p = HT_HASH_TO_BUCKET_EX(arData, i)
		for p.GetVal().GetNext() != idx {
			i = p.GetVal().GetNext()
			p = HT_HASH_TO_BUCKET_EX(arData, i)
		}
		p.GetVal().GetNext() = b.GetVal().GetNext()
	}
	ZendStringRelease(b.GetKey())

	/* add to hash */

	idx = b - arData
	b.SetKey(key)
	b.SetH(key.GetH())
	nIndex = b.GetH() | ht.GetNTableMask()
	idx = HT_IDX_TO_HASH(idx)
	i = HT_HASH_EX(arData, nIndex)
	if i == HT_INVALID_IDX || i < idx {
		b.GetVal().GetNext() = i
		HT_HASH_EX(arData, nIndex) = idx
	} else {
		p = HT_HASH_TO_BUCKET_EX(arData, i)
		for p.GetVal().GetNext() != HT_INVALID_IDX && p.GetVal().GetNext() > idx {
			i = p.GetVal().GetNext()
			p = HT_HASH_TO_BUCKET_EX(arData, i)
		}
		b.GetVal().GetNext() = p.GetVal().GetNext()
		p.GetVal().GetNext() = idx
	}
	return b.GetVal()
}
func ZendHashDoResize(ht *HashTable) {
	HT_ASSERT_RC1(ht)
	if ht.GetNNumUsed() > ht.GetNNumOfElements()+(ht.GetNNumOfElements()>>5) {
		ZendHashRehash(ht)
	} else if ht.GetNTableSize() < HT_MAX_SIZE {
		var new_data any
		var old_data any = HT_GET_DATA_ADDR(ht)
		var nSize uint32 = ht.GetNTableSize() + ht.GetNTableSize()
		var old_buckets *Bucket = ht.GetArData()
		ht.SetNTableSize(nSize)
		new_data = Pemalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)), GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
		ht.SetNTableMask(HT_SIZE_TO_MASK(ht.GetNTableSize()))
		HT_SET_DATA_ADDR(ht, new_data)
		memcpy(ht.GetArData(), old_buckets, b.SizeOf("Bucket")*ht.GetNNumUsed())
		Pefree(old_data, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
		ZendHashRehash(ht)
	} else {
		ZendErrorNoreturn(E_ERROR, "Possible integer overflow in memory allocation (%u * %zu + %zu)", ht.GetNTableSize()*2, b.SizeOf("Bucket")+b.SizeOf("uint32_t"), b.SizeOf("Bucket"))
	}
}
func ZendHashRehash(ht *HashTable) int {
	var p *Bucket
	var nIndex uint32
	var i uint32
	if ht.GetNNumOfElements() == 0 {
		if (ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) == 0 {
			ht.SetNNumUsed(0)
			HT_HASH_RESET(ht)
		}
		return SUCCESS
	}
	HT_HASH_RESET(ht)
	i = 0
	p = ht.GetArData()
	if HT_IS_WITHOUT_HOLES(ht) {
		for {
			nIndex = p.GetH() | ht.GetNTableMask()
			p.GetVal().GetNext() = HT_HASH(ht, nIndex)
			HT_HASH(ht, nIndex) = HT_IDX_TO_HASH(i)
			p++
			if b.PreInc(&i) >= ht.GetNNumUsed() {
				break
			}
		}
	} else {
		var old_num_used uint32 = ht.GetNNumUsed()
		for {
			if p.GetVal().IsType(IS_UNDEF) {
				var j uint32 = i
				var q *Bucket = p
				if !(HT_HAS_ITERATORS(ht)) {
					for b.PreInc(&i) < ht.GetNNumUsed() {
						p++
						if p.GetVal().GetTypeInfo() != IS_UNDEF {
							ZVAL_COPY_VALUE(q.GetVal(), p.GetVal())
							q.SetH(p.GetH())
							nIndex = q.GetH() | ht.GetNTableMask()
							q.SetKey(p.GetKey())
							q.GetVal().GetNext() = HT_HASH(ht, nIndex)
							HT_HASH(ht, nIndex) = HT_IDX_TO_HASH(j)
							if ht.GetNInternalPointer() == i {
								ht.SetNInternalPointer(j)
							}
							q++
							j++
						}
					}
				} else {
					var iter_pos uint32 = ZendHashIteratorsLowerPos(ht, 0)
					for b.PreInc(&i) < ht.GetNNumUsed() {
						p++
						if p.GetVal().GetTypeInfo() != IS_UNDEF {
							ZVAL_COPY_VALUE(q.GetVal(), p.GetVal())
							q.SetH(p.GetH())
							nIndex = q.GetH() | ht.GetNTableMask()
							q.SetKey(p.GetKey())
							q.GetVal().GetNext() = HT_HASH(ht, nIndex)
							HT_HASH(ht, nIndex) = HT_IDX_TO_HASH(j)
							if ht.GetNInternalPointer() == i {
								ht.SetNInternalPointer(j)
							}
							if i >= iter_pos {
								for {
									ZendHashIteratorsUpdate(ht, iter_pos, j)
									iter_pos = ZendHashIteratorsLowerPos(ht, iter_pos+1)
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
				ht.SetNNumUsed(j)
				break
			}
			nIndex = p.GetH() | ht.GetNTableMask()
			p.GetVal().GetNext() = HT_HASH(ht, nIndex)
			HT_HASH(ht, nIndex) = HT_IDX_TO_HASH(i)
			p++
			if b.PreInc(&i) >= ht.GetNNumUsed() {
				break
			}
		}

		/* Migrate pointer to one past the end of the array to the new one past the end, so that
		 * newly inserted elements are picked up correctly. */

		if HT_HAS_ITERATORS(ht) {
			_zendHashIteratorsUpdate(ht, old_num_used, ht.GetNNumUsed())
		}

		/* Migrate pointer to one past the end of the array to the new one past the end, so that
		 * newly inserted elements are picked up correctly. */

	}
	return SUCCESS
}
func _zendHashDelElEx(ht *HashTable, idx uint32, p *Bucket, prev *Bucket) {
	if (ht.GetUFlags() & HASH_FLAG_PACKED) == 0 {
		if prev != nil {
			prev.GetVal().GetNext() = p.GetVal().GetNext()
		} else {
			HT_HASH(ht, p.GetH()|ht.GetNTableMask()) = p.GetVal().GetNext()
		}
	}
	idx = HT_HASH_TO_IDX(idx)
	ht.GetNNumOfElements()--
	if ht.GetNInternalPointer() == idx || HT_HAS_ITERATORS(ht) {
		var new_idx uint32
		new_idx = idx
		for true {
			new_idx++
			if new_idx >= ht.GetNNumUsed() {
				break
			} else if ht.GetArData()[new_idx].GetVal().GetType() != IS_UNDEF {
				break
			}
		}
		if ht.GetNInternalPointer() == idx {
			ht.SetNInternalPointer(new_idx)
		}
		ZendHashIteratorsUpdate(ht, idx, new_idx)
	}
	if ht.GetNNumUsed()-1 == idx {
		for {
			ht.GetNNumUsed()--
			if !(ht.GetNNumUsed() > 0 && ht.GetArData()[ht.GetNNumUsed()-1].GetVal().IsType(IS_UNDEF)) {
				break
			}
		}
		ht.SetNInternalPointer(MIN(ht.GetNInternalPointer(), ht.GetNNumUsed()))
	}
	if p.GetKey() != nil {
		ZendStringRelease(p.GetKey())
	}
	if ht.GetPDestructor() != nil {
		var tmp Zval
		ZVAL_COPY_VALUE(&tmp, p.GetVal())
		ZVAL_UNDEF(p.GetVal())
		ht.GetPDestructor()(&tmp)
	} else {
		ZVAL_UNDEF(p.GetVal())
	}
}
func _zendHashDelEl(ht *HashTable, idx uint32, p *Bucket) {
	var prev *Bucket = nil
	if (ht.GetUFlags() & HASH_FLAG_PACKED) == 0 {
		var nIndex uint32 = p.GetH() | ht.GetNTableMask()
		var i uint32 = HT_HASH(ht, nIndex)
		if i != idx {
			prev = HT_HASH_TO_BUCKET(ht, i)
			for prev.GetVal().GetNext() != idx {
				i = prev.GetVal().GetNext()
				prev = HT_HASH_TO_BUCKET(ht, i)
			}
		}
	}
	_zendHashDelElEx(ht, idx, p, prev)
}
func ZendHashDelBucket(ht *HashTable, p *Bucket) {
	HT_ASSERT_RC1(ht)
	_zendHashDelEl(ht, HT_IDX_TO_HASH(p-ht.GetArData()), p)
}
func ZendHashDel(ht *HashTable, key *ZendString) int {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	HT_ASSERT_RC1(ht)
	h = ZendStringHashVal(key)
	nIndex = h | ht.GetNTableMask()
	idx = HT_HASH(ht, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(ht, idx)
		if p.GetKey() == key || p.GetH() == h && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), key) != 0 {
			_zendHashDelElEx(ht, idx, p, prev)
			return SUCCESS
		}
		prev = p
		idx = p.GetVal().GetNext()
	}
	return FAILURE
}
func ZendHashDelInd(ht *HashTable, key *ZendString) int {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	HT_ASSERT_RC1(ht)
	h = ZendStringHashVal(key)
	nIndex = h | ht.GetNTableMask()
	idx = HT_HASH(ht, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(ht, idx)
		if p.GetKey() == key || p.GetH() == h && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), key) != 0 {
			if p.GetVal().IsType(IS_INDIRECT) {
				var data *Zval = p.GetVal().GetZv()
				if data.GetType() == IS_UNDEF {
					return FAILURE
				} else {
					if ht.GetPDestructor() != nil {
						var tmp Zval
						ZVAL_COPY_VALUE(&tmp, data)
						ZVAL_UNDEF(data)
						ht.GetPDestructor()(&tmp)
					} else {
						ZVAL_UNDEF(data)
					}
					ht.GetUFlags() |= HASH_FLAG_HAS_EMPTY_IND
				}
			} else {
				_zendHashDelElEx(ht, idx, p, prev)
			}
			return SUCCESS
		}
		prev = p
		idx = p.GetVal().GetNext()
	}
	return FAILURE
}
func ZendHashStrDelInd(ht *HashTable, str *byte, len_ int) int {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	HT_ASSERT_RC1(ht)
	h = ZendInlineHashFunc(str, len_)
	nIndex = h | ht.GetNTableMask()
	idx = HT_HASH(ht, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(ht, idx)
		if p.GetH() == h && p.GetKey() != nil && p.GetKey().GetLen() == len_ && !(memcmp(p.GetKey().GetVal(), str, len_)) {
			if p.GetVal().IsType(IS_INDIRECT) {
				var data *Zval = p.GetVal().GetZv()
				if data.GetType() == IS_UNDEF {
					return FAILURE
				} else {
					if ht.GetPDestructor() != nil {
						ht.GetPDestructor()(data)
					}
					ZVAL_UNDEF(data)
					ht.GetUFlags() |= HASH_FLAG_HAS_EMPTY_IND
				}
			} else {
				_zendHashDelElEx(ht, idx, p, prev)
			}
			return SUCCESS
		}
		prev = p
		idx = p.GetVal().GetNext()
	}
	return FAILURE
}
func ZendHashStrDel(ht *HashTable, str *byte, len_ int) int {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	HT_ASSERT_RC1(ht)
	h = ZendInlineHashFunc(str, len_)
	nIndex = h | ht.GetNTableMask()
	idx = HT_HASH(ht, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(ht, idx)
		if p.GetH() == h && p.GetKey() != nil && p.GetKey().GetLen() == len_ && !(memcmp(p.GetKey().GetVal(), str, len_)) {
			_zendHashDelElEx(ht, idx, p, prev)
			return SUCCESS
		}
		prev = p
		idx = p.GetVal().GetNext()
	}
	return FAILURE
}
func ZendHashIndexDel(ht *HashTable, h ZendUlong) int {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	HT_ASSERT_RC1(ht)
	if (ht.GetUFlags() & HASH_FLAG_PACKED) != 0 {
		if h < ht.GetNNumUsed() {
			p = ht.GetArData() + h
			if p.GetVal().GetType() != IS_UNDEF {
				_zendHashDelElEx(ht, HT_IDX_TO_HASH(h), p, nil)
				return SUCCESS
			}
		}
		return FAILURE
	}
	nIndex = h | ht.GetNTableMask()
	idx = HT_HASH(ht, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(ht, idx)
		if p.GetH() == h && p.GetKey() == nil {
			_zendHashDelElEx(ht, idx, p, prev)
			return SUCCESS
		}
		prev = p
		idx = p.GetVal().GetNext()
	}
	return FAILURE
}
func ZendHashDestroy(ht *HashTable) {
	var p *Bucket
	var end *Bucket
	if ht.GetNNumUsed() != 0 {
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if ht.GetPDestructor() != nil {
			if HT_HAS_STATIC_KEYS_ONLY(ht) {
				if HT_IS_WITHOUT_HOLES(ht) {
					for {
						ht.GetPDestructor()(p.GetVal())
						if b.PreInc(&p) == end {
							break
						}
					}
				} else {
					for {
						if p.GetVal().GetType() != IS_UNDEF {
							ht.GetPDestructor()(p.GetVal())
						}
						if b.PreInc(&p) == end {
							break
						}
					}
				}
			} else if HT_IS_WITHOUT_HOLES(ht) {
				for {
					ht.GetPDestructor()(p.GetVal())
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
						ht.GetPDestructor()(p.GetVal())
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
			if !(HT_HAS_STATIC_KEYS_ONLY(ht)) {
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
		ZendHashIteratorsRemove(ht)
	} else if (ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) != 0 {
		return
	}
	Pefree(HT_GET_DATA_ADDR(ht), GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
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
			ZendHashDestroy(ht)
			goto free_ht
		}
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if HT_HAS_STATIC_KEYS_ONLY(ht) {
			for {
				IZvalPtrDtor(p.GetVal())
				if b.PreInc(&p) == end {
					break
				}
			}
		} else if HT_IS_WITHOUT_HOLES(ht) {
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
	} else if (ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) != 0 {
		goto free_ht
	}
	Efree(HT_GET_DATA_ADDR(ht))
free_ht:
	ZendHashIteratorsRemove(ht)
	FREE_HASHTABLE(ht)
}
func ZendHashClean(ht *HashTable) {
	var p *Bucket
	var end *Bucket
	HT_ASSERT_RC1(ht)
	if ht.GetNNumUsed() != 0 {
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if ht.GetPDestructor() != nil {
			if HT_HAS_STATIC_KEYS_ONLY(ht) {
				if HT_IS_WITHOUT_HOLES(ht) {
					for {
						ht.GetPDestructor()(p.GetVal())
						if b.PreInc(&p) == end {
							break
						}
					}
				} else {
					for {
						if p.GetVal().GetType() != IS_UNDEF {
							ht.GetPDestructor()(p.GetVal())
						}
						if b.PreInc(&p) == end {
							break
						}
					}
				}
			} else if HT_IS_WITHOUT_HOLES(ht) {
				for {
					ht.GetPDestructor()(p.GetVal())
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
						ht.GetPDestructor()(p.GetVal())
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
			if !(HT_HAS_STATIC_KEYS_ONLY(ht)) {
				if HT_IS_WITHOUT_HOLES(ht) {
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
		if (ht.GetUFlags() & HASH_FLAG_PACKED) == 0 {
			HT_HASH_RESET(ht)
		}
	}
	ht.SetNNumUsed(0)
	ht.SetNNumOfElements(0)
	ht.SetNNextFreeElement(0)
	ht.SetNInternalPointer(0)
}
func ZendSymtableClean(ht *HashTable) {
	var p *Bucket
	var end *Bucket
	HT_ASSERT_RC1(ht)
	if ht.GetNNumUsed() != 0 {
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if HT_HAS_STATIC_KEYS_ONLY(ht) {
			for {
				IZvalPtrDtor(p.GetVal())
				if b.PreInc(&p) == end {
					break
				}
			}
		} else if HT_IS_WITHOUT_HOLES(ht) {
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
		HT_HASH_RESET(ht)
	}
	ht.SetNNumUsed(0)
	ht.SetNNumOfElements(0)
	ht.SetNNextFreeElement(0)
	ht.SetNInternalPointer(0)
}
func ZendHashGracefulDestroy(ht *HashTable) {
	var idx uint32
	var p *Bucket
	HT_ASSERT_RC1(ht)
	p = ht.GetArData()
	for idx = 0; idx < ht.GetNNumUsed(); {
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		_zendHashDelEl(ht, HT_IDX_TO_HASH(idx), p)
		idx++
		p++
	}
	if (ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) == 0 {
		Pefree(HT_GET_DATA_ADDR(ht), GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
	}
}
func ZendHashGracefulReverseDestroy(ht *HashTable) {
	var idx uint32
	var p *Bucket
	HT_ASSERT_RC1(ht)
	idx = ht.GetNNumUsed()
	p = ht.GetArData() + ht.GetNNumUsed()
	for idx > 0 {
		idx--
		p--
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		_zendHashDelEl(ht, HT_IDX_TO_HASH(idx), p)
	}
	if (ht.GetUFlags() & HASH_FLAG_UNINITIALIZED) == 0 {
		Pefree(HT_GET_DATA_ADDR(ht), GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
	}
}
func ZendHashApply(ht *HashTable, apply_func ApplyFuncT) {
	var idx uint32
	var p *Bucket
	var result int
	for idx = 0; idx < ht.GetNNumUsed(); idx++ {
		p = ht.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		result = apply_func(p.GetVal())
		if (result & ZEND_HASH_APPLY_REMOVE) != 0 {
			HT_ASSERT_RC1(ht)
			_zendHashDelEl(ht, HT_IDX_TO_HASH(idx), p)
		}
		if (result & ZEND_HASH_APPLY_STOP) != 0 {
			break
		}
	}
}
func ZendHashApplyWithArgument(ht *HashTable, apply_func ApplyFuncArgT, argument any) {
	var idx uint32
	var p *Bucket
	var result int
	for idx = 0; idx < ht.GetNNumUsed(); idx++ {
		p = ht.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		result = apply_func(p.GetVal(), argument)
		if (result & ZEND_HASH_APPLY_REMOVE) != 0 {
			HT_ASSERT_RC1(ht)
			_zendHashDelEl(ht, HT_IDX_TO_HASH(idx), p)
		}
		if (result & ZEND_HASH_APPLY_STOP) != 0 {
			break
		}
	}
}
func ZendHashApplyWithArguments(ht *HashTable, apply_func ApplyFuncArgsT, num_args int, _ ...any) {
	var idx uint32
	var p *Bucket
	var args va_list
	var hash_key ZendHashKey
	var result int
	for idx = 0; idx < ht.GetNNumUsed(); idx++ {
		p = ht.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		va_start(args, num_args)
		hash_key.SetH(p.GetH())
		hash_key.SetKey(p.GetKey())
		result = apply_func(p.GetVal(), num_args, args, &hash_key)
		if (result & ZEND_HASH_APPLY_REMOVE) != 0 {
			HT_ASSERT_RC1(ht)
			_zendHashDelEl(ht, HT_IDX_TO_HASH(idx), p)
		}
		if (result & ZEND_HASH_APPLY_STOP) != 0 {
			va_end(args)
			break
		}
		va_end(args)
	}
}
func ZendHashReverseApply(ht *HashTable, apply_func ApplyFuncT) {
	var idx uint32
	var p *Bucket
	var result int
	idx = ht.GetNNumUsed()
	for idx > 0 {
		idx--
		p = ht.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		result = apply_func(p.GetVal())
		if (result & ZEND_HASH_APPLY_REMOVE) != 0 {
			HT_ASSERT_RC1(ht)
			_zendHashDelEl(ht, HT_IDX_TO_HASH(idx), p)
		}
		if (result & ZEND_HASH_APPLY_STOP) != 0 {
			break
		}
	}
}
func ZendHashCopy(target *HashTable, source *HashTable, pCopyConstructor CopyCtorFuncT) {
	var idx uint32
	var p *Bucket
	var new_entry *Zval
	var data *Zval
	HT_ASSERT_RC1(target)
	for idx = 0; idx < source.GetNNumUsed(); idx++ {
		p = source.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}

		/* INDIRECT element may point to UNDEF-ined slots */

		data = p.GetVal()
		if data.GetType() == IS_INDIRECT {
			data = data.GetZv()
			if data.GetType() == IS_UNDEF {
				continue
			}
		}
		if p.GetKey() != nil {
			new_entry = ZendHashUpdate(target, p.GetKey(), data)
		} else {
			new_entry = ZendHashIndexUpdate(target, p.GetH(), data)
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
			if Z_ISREF_P(data) && Z_REFCOUNT_P(data) == 1 && (Z_REFVAL_P(data).GetType() != IS_ARRAY || Z_ARRVAL_P(Z_REFVAL_P(data)) != source) {
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
		q.GetVal().GetNext() = HT_HASH(target, nIndex)
		HT_HASH(target, nIndex) = HT_IDX_TO_HASH(idx)
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
		target.GetUFlags() = HASH_FLAG_UNINITIALIZED
		target.SetNTableMask(HT_MIN_MASK)
		target.SetNNumUsed(0)
		target.SetNNumOfElements(0)
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNInternalPointer(0)
		target.SetNTableSize(HT_MIN_SIZE)
		HT_SET_DATA_ADDR(target, &UninitializedBucket)
	} else if (GC_FLAGS(source) & IS_ARRAY_IMMUTABLE) != 0 {
		target.GetUFlags() = source.GetUFlags() & HASH_FLAG_MASK
		target.SetNTableMask(source.GetNTableMask())
		target.SetNNumUsed(source.GetNNumUsed())
		target.SetNNumOfElements(source.GetNNumOfElements())
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNTableSize(source.GetNTableSize())
		HT_SET_DATA_ADDR(target, Emalloc(HT_SIZE(target)))
		target.SetNInternalPointer(source.GetNInternalPointer())
		memcpy(HT_GET_DATA_ADDR(target), HT_GET_DATA_ADDR(source), HT_USED_SIZE(source))
	} else if (source.GetUFlags() & HASH_FLAG_PACKED) != 0 {
		target.GetUFlags() = source.GetUFlags() & HASH_FLAG_MASK
		target.SetNTableMask(HT_MIN_MASK)
		target.SetNNumUsed(source.GetNNumUsed())
		target.SetNNumOfElements(source.GetNNumOfElements())
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNTableSize(source.GetNTableSize())
		HT_SET_DATA_ADDR(target, Emalloc(HT_SIZE_EX(target.GetNTableSize(), HT_MIN_MASK)))
		if source.GetNInternalPointer() < source.GetNNumUsed() {
			target.SetNInternalPointer(source.GetNInternalPointer())
		} else {
			target.SetNInternalPointer(0)
		}
		HT_HASH_RESET_PACKED(target)
		if HT_IS_WITHOUT_HOLES(target) {
			ZendArrayDupPackedElements(source, target, 0)
		} else {
			ZendArrayDupPackedElements(source, target, 1)
		}
	} else {
		target.GetUFlags() = source.GetUFlags() & HASH_FLAG_MASK
		target.SetNTableMask(source.GetNTableMask())
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		if source.GetNInternalPointer() < source.GetNNumUsed() {
			target.SetNInternalPointer(source.GetNInternalPointer())
		} else {
			target.SetNInternalPointer(0)
		}
		target.SetNTableSize(source.GetNTableSize())
		HT_SET_DATA_ADDR(target, Emalloc(HT_SIZE(target)))
		HT_HASH_RESET(target)
		if HT_HAS_STATIC_KEYS_ONLY(target) {
			if HT_IS_WITHOUT_HOLES(source) {
				idx = ZendArrayDupElements(source, target, 1, 0)
			} else {
				idx = ZendArrayDupElements(source, target, 1, 1)
			}
		} else {
			if HT_IS_WITHOUT_HOLES(source) {
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
func ZendHashMerge(target *HashTable, source *HashTable, pCopyConstructor CopyCtorFuncT, overwrite ZendBool) {
	var idx uint32
	var p *Bucket
	var t *Zval
	var s *Zval
	HT_ASSERT_RC1(target)
	if overwrite != 0 {
		for idx = 0; idx < source.GetNNumUsed(); idx++ {
			p = source.GetArData() + idx
			s = p.GetVal()
			if s.GetType() == IS_INDIRECT {
				s = s.GetZv()
			}
			if s.GetType() == IS_UNDEF {
				continue
			}
			if p.GetKey() != nil {
				t = _zendHashAddOrUpdateI(target, p.GetKey(), s, HASH_UPDATE|HASH_UPDATE_INDIRECT)
				if pCopyConstructor != nil {
					pCopyConstructor(t)
				}
			} else {
				t = ZendHashIndexUpdate(target, p.GetH(), s)
				if pCopyConstructor != nil {
					pCopyConstructor(t)
				}
			}
		}
	} else {
		for idx = 0; idx < source.GetNNumUsed(); idx++ {
			p = source.GetArData() + idx
			s = p.GetVal()
			if s.GetType() == IS_INDIRECT {
				s = s.GetZv()
			}
			if s.GetType() == IS_UNDEF {
				continue
			}
			if p.GetKey() != nil {
				t = _zendHashAddOrUpdateI(target, p.GetKey(), s, HASH_ADD|HASH_UPDATE_INDIRECT)
				if t != nil && pCopyConstructor != nil {
					pCopyConstructor(t)
				}
			} else {
				t = ZendHashIndexAdd(target, p.GetH(), s)
				if t != nil && pCopyConstructor != nil {
					pCopyConstructor(t)
				}
			}
		}
	}
}
func ZendHashReplaceCheckerWrapper(target *HashTable, source_data *Zval, p *Bucket, pParam any, merge_checker_func MergeCheckerFuncT) ZendBool {
	var hash_key ZendHashKey
	hash_key.SetH(p.GetH())
	hash_key.SetKey(p.GetKey())
	return merge_checker_func(target, source_data, &hash_key, pParam)
}
func ZendHashMergeEx(target *HashTable, source *HashTable, pCopyConstructor CopyCtorFuncT, pMergeSource MergeCheckerFuncT, pParam any) {
	var idx uint32
	var p *Bucket
	var t *Zval
	HT_ASSERT_RC1(target)
	for idx = 0; idx < source.GetNNumUsed(); idx++ {
		p = source.GetArData() + idx
		if p.GetVal().IsType(IS_UNDEF) {
			continue
		}
		if ZendHashReplaceCheckerWrapper(target, p.GetVal(), p, pParam, pMergeSource) != 0 {
			t = ZendHashUpdate(target, p.GetKey(), p.GetVal())
			if pCopyConstructor != nil {
				pCopyConstructor(t)
			}
		}
	}
}
func ZendHashFind(ht *HashTable, key *ZendString) *Zval {
	var p *Bucket
	p = ZendHashFindBucket(ht, key, 0)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func _zendHashFindKnownHash(ht *HashTable, key *ZendString) *Zval {
	var p *Bucket
	p = ZendHashFindBucket(ht, key, 1)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func ZendHashStrFind(ht *HashTable, str *byte, len_ int) *Zval {
	var h ZendUlong
	var p *Bucket
	h = ZendInlineHashFunc(str, len_)
	p = ZendHashStrFindBucket(ht, str, len_, h)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func ZendHashIndexFind(ht *HashTable, h ZendUlong) *Zval {
	var p *Bucket
	if (ht.GetUFlags() & HASH_FLAG_PACKED) != 0 {
		if h < ht.GetNNumUsed() {
			p = ht.GetArData() + h
			if p.GetVal().GetType() != IS_UNDEF {
				return p.GetVal()
			}
		}
		return nil
	}
	p = ZendHashIndexFindBucket(ht, h)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func _zendHashIndexFind(ht *HashTable, h ZendUlong) *Zval {
	var p *Bucket
	p = ZendHashIndexFindBucket(ht, h)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func ZendHashInternalPointerResetEx(ht *HashTable, pos *HashPosition) {
	*pos = _zendHashGetValidPos(ht, 0)
}
func ZendHashInternalPointerEndEx(ht *HashTable, pos *HashPosition) {
	var idx uint32
	idx = ht.GetNNumUsed()
	for idx > 0 {
		idx--
		if ht.GetArData()[idx].GetVal().GetType() != IS_UNDEF {
			*pos = idx
			return
		}
	}
	*pos = ht.GetNNumUsed()
}
func ZendHashMoveForwardEx(ht *HashTable, pos *HashPosition) int {
	var idx uint32
	idx = _zendHashGetValidPos(ht, *pos)
	if idx < ht.GetNNumUsed() {
		for true {
			idx++
			if idx >= ht.GetNNumUsed() {
				*pos = ht.GetNNumUsed()
				return SUCCESS
			}
			if ht.GetArData()[idx].GetVal().GetType() != IS_UNDEF {
				*pos = idx
				return SUCCESS
			}
		}
	} else {
		return FAILURE
	}
}
func ZendHashMoveBackwardsEx(ht *HashTable, pos *HashPosition) int {
	var idx uint32 = *pos
	if idx < ht.GetNNumUsed() {
		for idx > 0 {
			idx--
			if ht.GetArData()[idx].GetVal().GetType() != IS_UNDEF {
				*pos = idx
				return SUCCESS
			}
		}
		*pos = ht.GetNNumUsed()
		return SUCCESS
	} else {
		return FAILURE
	}
}
func ZendHashGetCurrentKeyEx(ht *HashTable, str_index **ZendString, num_index *ZendUlong, pos *HashPosition) int {
	var idx uint32
	var p *Bucket
	idx = _zendHashGetValidPos(ht, *pos)
	if idx < ht.GetNNumUsed() {
		p = ht.GetArData() + idx
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
func ZendHashGetCurrentKeyZvalEx(ht *HashTable, key *Zval, pos *HashPosition) {
	var idx uint32
	var p *Bucket
	idx = _zendHashGetValidPos(ht, *pos)
	if idx >= ht.GetNNumUsed() {
		ZVAL_NULL(key)
	} else {
		p = ht.GetArData() + idx
		if p.GetKey() != nil {
			ZVAL_STR_COPY(key, p.GetKey())
		} else {
			ZVAL_LONG(key, p.GetH())
		}
	}
}
func ZendHashGetCurrentKeyTypeEx(ht *HashTable, pos *HashPosition) int {
	var idx uint32
	var p *Bucket
	idx = _zendHashGetValidPos(ht, *pos)
	if idx < ht.GetNNumUsed() {
		p = ht.GetArData() + idx
		if p.GetKey() != nil {
			return HASH_KEY_IS_STRING
		} else {
			return HASH_KEY_IS_LONG
		}
	}
	return HASH_KEY_NON_EXISTENT
}
func ZendHashGetCurrentDataEx(ht *HashTable, pos *HashPosition) *Zval {
	var idx uint32
	var p *Bucket
	idx = _zendHashGetValidPos(ht, *pos)
	if idx < ht.GetNNumUsed() {
		p = ht.GetArData() + idx
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
func ZendHashSortEx(ht *HashTable, sort SortFuncT, compar CompareFuncT, renumber ZendBool) int {
	var p *Bucket
	var i uint32
	var j uint32
	HT_ASSERT_RC1(ht)
	if ht.GetNNumOfElements() <= 1 && !(renumber != 0 && ht.GetNNumOfElements() > 0) {
		return SUCCESS
	}
	if HT_IS_WITHOUT_HOLES(ht) {
		i = ht.GetNNumUsed()
	} else {
		j = 0
		i = 0
		for ; j < ht.GetNNumUsed(); j++ {
			p = ht.GetArData() + j
			if p.GetVal().IsType(IS_UNDEF) {
				continue
			}
			if i != j {
				ht.GetArData()[i] = *p
			}
			i++
		}
	}
	sort(any(ht.GetArData()), i, b.SizeOf("Bucket"), compar, swap_func_t(b.CondF2(renumber != 0, ZendHashBucketRenumSwap, func() __auto__ {
		if (ht.GetUFlags() & HASH_FLAG_PACKED) != 0 {
			return ZendHashBucketPackedSwap
		} else {
			return ZendHashBucketSwap
		}
	})))
	ht.SetNNumUsed(i)
	ht.SetNInternalPointer(0)
	if renumber != 0 {
		for j = 0; j < i; j++ {
			p = ht.GetArData() + j
			p.SetH(j)
			if p.GetKey() != nil {
				ZendStringRelease(p.GetKey())
				p.SetKey(nil)
			}
		}
		ht.SetNNextFreeElement(i)
	}
	if (ht.GetUFlags() & HASH_FLAG_PACKED) != 0 {
		if renumber == 0 {
			ZendHashPackedToHash(ht)
		}
	} else {
		if renumber != 0 {
			var new_data any
			var old_data any = HT_GET_DATA_ADDR(ht)
			var old_buckets *Bucket = ht.GetArData()
			new_data = Pemalloc(HT_SIZE_EX(ht.GetNTableSize(), HT_MIN_MASK), GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
			ht.GetUFlags() |= HASH_FLAG_PACKED | HASH_FLAG_STATIC_KEYS
			ht.SetNTableMask(HT_MIN_MASK)
			HT_SET_DATA_ADDR(ht, new_data)
			memcpy(ht.GetArData(), old_buckets, b.SizeOf("Bucket")*ht.GetNNumUsed())
			Pefree(old_data, GC_FLAGS(ht)&IS_ARRAY_PERSISTENT)
			HT_HASH_RESET_PACKED(ht)
		} else {
			ZendHashRehash(ht)
		}
	}
	return SUCCESS
}
func ZendHashCompareImpl(ht1 *HashTable, ht2 *HashTable, compar CompareFuncT, ordered ZendBool) int {
	var idx1 uint32
	var idx2 uint32
	if ht1.GetNNumOfElements() != ht2.GetNNumOfElements() {
		if ht1.GetNNumOfElements() > ht2.GetNNumOfElements() {
			return 1
		} else {
			return -1
		}
	}
	idx1 = 0
	idx2 = 0
	for ; idx1 < ht1.GetNNumUsed(); idx1++ {
		var p1 *Bucket = ht1.GetArData() + idx1
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
				pData2 = ZendHashIndexFind(ht2, p1.GetH())
				if pData2 == nil {
					return 1
				}
			} else {
				pData2 = ZendHashFind(ht2, p1.GetKey())
				if pData2 == nil {
					return 1
				}
			}
		}
		pData1 = p1.GetVal()
		if pData1.GetType() == IS_INDIRECT {
			pData1 = pData1.GetZv()
		}
		if pData2.GetType() == IS_INDIRECT {
			pData2 = pData2.GetZv()
		}
		if pData1.GetType() == IS_UNDEF {
			if pData2.GetType() != IS_UNDEF {
				return -1
			}
		} else if pData2.GetType() == IS_UNDEF {
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
func ZendHashCompare(ht1 *HashTable, ht2 *HashTable, compar CompareFuncT, ordered ZendBool) int {
	var result int
	if ht1 == ht2 {
		return 0
	}

	/* It's enough to protect only one of the arrays.
	 * The second one may be referenced from the first and this may cause
	 * false recursion detection.
	 */

	if GC_IS_RECURSIVE(ht1) != 0 {
		ZendErrorNoreturn(E_ERROR, "Nesting level too deep - recursive dependency?")
	}
	if (GC_FLAGS(ht1) & GC_IMMUTABLE) == 0 {
		GC_PROTECT_RECURSION(ht1)
	}
	result = ZendHashCompareImpl(ht1, ht2, compar, ordered)
	if (GC_FLAGS(ht1) & GC_IMMUTABLE) == 0 {
		GC_UNPROTECT_RECURSION(ht1)
	}
	return result
}
func ZendHashMinmax(ht *HashTable, compar CompareFuncT, flag uint32) *Zval {
	var idx uint32
	var p *Bucket
	var res *Bucket
	if ht.GetNNumOfElements() == 0 {
		return nil
	}
	idx = 0
	for true {
		if idx == ht.GetNNumUsed() {
			return nil
		}
		if ht.GetArData()[idx].GetVal().GetType() != IS_UNDEF {
			break
		}
		idx++
	}
	res = ht.GetArData() + idx
	for ; idx < ht.GetNNumUsed(); idx++ {
		p = ht.GetArData() + idx
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
	if HT_IS_PACKED(ht) {
		goto convert
	}
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = _p.GetVal()

			if _z.GetType() == IS_UNDEF {
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

			if _z.GetType() == IS_UNDEF {
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
			ZendHashUpdate(new_ht, str_key, zv)
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

			if _z.GetType() == IS_UNDEF {
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
			if _z.GetType() == IS_INDIRECT {
				_z = _z.GetZv()
			}
			if _z.GetType() == IS_UNDEF {
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
				ZendHashIndexUpdate(new_ht, num_key, zv)
			} else {
				ZendHashUpdate(new_ht, str_key, zv)
			}

			/* Again, thank ArrayObject for `!str_key ||`. */

		}
		break
	}
	return new_ht
}
