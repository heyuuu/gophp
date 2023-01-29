// <<generate>>

package zend

import (
	b "sik/builtin"
)

func HT_ITERATORS_OVERFLOW(ht *HashTable) bool { return ht.GetNIteratorsCount() == 0xff }
func HT_HAS_ITERATORS(ht *HashTable) bool      { return ht.GetNIteratorsCount() != 0 }
func HT_INC_ITERATORS_COUNT(ht *HashTable) {
	ht.SetNIteratorsCount(ht.GetNIteratorsCount() + 1)
}
func HT_DEC_ITERATORS_COUNT(ht *HashTable) {
	ht.SetNIteratorsCount(ht.GetNIteratorsCount() - 1)
}
func ZVAL_EMPTY_ARRAY(z *Zval) {
	z.SetArr((*ZendArray)(&ZendEmptyArray))
	z.SetTypeInfo(IS_ARRAY)
}
func ZendHashInit(ht *HashTable, nSize uint32, pHashFunction any, pDestructor DtorFuncT, persistent ZendBool) {
	*ht = *NewZendArrayEx(nSize, pDestructor, persistent != 0)
}
func ZendHashInitEx(ht *HashTable, nSize uint32, pHashFunction any, pDestructor DtorFuncT, persistent ZendBool, bApplyProtection int) {
	*ht = *NewZendArrayEx(nSize, pDestructor, persistent != 0)
}
func ZendHashFindEx(ht *HashTable, key *ZendString, known_hash ZendBool) *Zval {
	if known_hash != 0 {
		return _zendHashFindKnownHash(ht, key)
	} else {
		return ZendHashFind(ht, key)
	}
}
func ZEND_HASH_INDEX_FIND(_ht *HashTable, _h ZendUlong, _ret *Zval, _not_found __auto__) {
	_ret = _zendHashIndexFind(_ht, _h)
	if _ret == nil {
		goto _not_found
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
func ZendNewArray(size uint32) *HashTable { return NewZendArray(size) }
func ZendHashIteratorsUpdate(ht *HashTable, from HashPosition, to HashPosition) {
	if HT_HAS_ITERATORS(ht) {
		_zendHashIteratorsUpdate(ht, from, to)
	}
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
func ZendHashFindExInd(ht *HashTable, key *ZendString, known_hash ZendBool) *Zval {
	var zv *Zval
	zv = ZendHashFindEx(ht, key, known_hash)
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
func ZendHashExistsInd(ht *HashTable, key *ZendString) int {
	var zv *Zval
	zv = ZendHashFind(ht, key)
	return zv != nil && (zv.GetType() != IS_INDIRECT || Z_INDIRECT_P(zv).GetType() != IS_UNDEF)
}
func ZendHashStrExistsInd(ht *HashTable, str string, len_ int) int {
	var zv *Zval
	zv = ZendHashStrFind(ht, str, len_)
	return zv != nil && (zv.GetType() != IS_INDIRECT || Z_INDIRECT_P(zv).GetType() != IS_UNDEF)
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
func ZendSymtableFind(ht *HashTable, key *ZendString) *Zval {
	var idx ZendUlong
	if ZEND_HANDLE_NUMERIC(key, idx) != 0 {
		return ZendHashIndexFind(ht, idx)
	} else {
		return ZendHashFind(ht, key)
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
		zv.SetPtr(Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT))
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
		zv.SetPtr(Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashUpdateMem(ht *HashTable, key *ZendString, pData any, size int) any {
	var p any
	p = Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashUpdatePtr(ht, key, p)
}
func ZendHashStrUpdateMem(ht *HashTable, str *byte, len_ int, pData any, size int) any {
	var p any
	p = Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
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
		zv.SetPtr(Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT))
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
	p = Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashIndexUpdatePtr(ht, h, p)
}
func ZendHashNextIndexInsertMem(ht *HashTable, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ZendHashNextIndexInsert(ht, &tmp)) {
		zv.SetPtr(Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT))
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
		ht.SubUFlags(HASH_FLAG_STATIC_KEYS)
		key.AddRefcount()
		key.GetHash()
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
		ht.SubUFlags(HASH_FLAG_STATIC_KEYS)
		key.AddRefcount()
		key.GetHash()
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
	ht.SubUFlags(HASH_FLAG_STATIC_KEYS)
	key.AddRefcount()
	key.GetHash()
	p.SetKey(key)
	p.SetH(key.GetH())
	nIndex = uint32(p.GetH() | ht.GetNTableMask())
	p.GetVal().GetNext() = HT_HASH(ht, nIndex)
	HT_HASH(ht, nIndex) = HT_IDX_TO_HASH(idx)
	ht.GetNNumOfElements()++
}
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

func ZendHashRealInit(ht *HashTable, packed ZendBool) { ht.RealInit() }
func ZendHashRealInitPacked(ht *HashTable)            { ht.RealInit() }
func ZendHashRealInitMixed(ht *HashTable)             { ht.RealInit() }
func ZendHashToPacked(ht *HashTable) {
	// todo 此函数不应被调用
	ZEND_ASSERT(false)
}
func ZendHashExtend(ht *HashTable, nSize uint32, packed ZendBool) {
	ht.assertRc1()
	if nSize == 0 {
		return
	}
	if nSize > ht.GetNTableSize() {
		var new_data any
		var old_data any = HT_GET_DATA_ADDR(ht)
		var old_buckets *Bucket = ht.GetArData()
		nSize = ZendHashCheckSize(nSize)
		ht.SetNTableSize(nSize)
		new_data = Pemalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)), ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
		ht.SetNTableMask(HT_SIZE_TO_MASK(ht.GetNTableSize()))
		HT_SET_DATA_ADDR(ht, new_data)
		memcpy(ht.GetArData(), old_buckets, b.SizeOf("Bucket")*ht.GetNNumUsed())
		Pefree(old_data, ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
		ZendHashRehash(ht)
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
	var ht *HashTable = array.GetArr()
	var iter *HashTableIterator = ExecutorGlobals.GetHtIterators() + idx
	ZEND_ASSERT(idx != uint32-1)
	if iter.GetHt() != ht {
		if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(HT_ITERATORS_OVERFLOW(ht)) {
			HT_DEC_ITERATORS_COUNT(iter.GetHt())
		}
		SEPARATE_ARRAY(array)
		ht = array.GetArr()
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
func ZendHashFindBucket(ht *HashTable, key *ZendString) *Bucket {
	var key_ = NewStrKey(key.GetStr())
	return ht.FindBucket(key_)
}
func ZendHashStrFindBucket(ht *HashTable, str *byte, len_ int) *Bucket {
	var key_ = NewStrKey(b.CastStr(str, len_))
	return ht.FindBucket(key_)
}
func ZendHashIndexFindBucket(ht *HashTable, h ZendUlong) *Bucket {
	var key_ = NewIndexKey(int(h))
	return ht.FindBucket(key_)
}
func _zendHashAddOrUpdateI(ht *HashTable, key *ZendString, pData *Zval, flag uint32) *Zval {
	var strKey = key.GetStr()
	return ht.addOrUpdate(strKey, pData, flag)
}

func _zendHashStrAddOrUpdateI(ht *HashTable, str *byte, len_ int, pData *Zval, flag uint32) *Zval {
	var strKey = b.CastStr(str, len_)
	return ht.addOrUpdate(strKey, pData, flag)
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
func ZendHashStrUpdate(ht *HashTable, str string, len_ int, pData *Zval) *Zval {
	return _zendHashStrAddOrUpdateI(ht, str, len_, pData, HASH_UPDATE)
}
func ZendHashStrUpdateInd(ht *HashTable, str string, len_ int, pData *Zval) *Zval {
	return _zendHashStrAddOrUpdateI(ht, str, len_, pData, HASH_UPDATE|HASH_UPDATE_INDIRECT)
}
func ZendHashStrAdd(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	return _zendHashStrAddOrUpdateI(ht, str, len_, pData, HASH_ADD)
}
func ZendHashStrAddNew(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	return _zendHashStrAddOrUpdateI(ht, str, len_, pData, HASH_ADD_NEW)
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
	return ht.indexAddOrUpdate(int(h), pData, flag)
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
	ht.assertRc1()
	p = ZendHashFindBucket(ht, key)
	if p != nil {
		if p == b {
			return p.GetVal()
		} else {
			return nil
		}
	}
	key.AddRefcount()
	ht.SubUFlags(HASH_FLAG_STATIC_KEYS)
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
	ht.assertRc1()
	if ht.GetNNumUsed() > ht.GetNNumOfElements()+(ht.GetNNumOfElements()>>5) {
		ZendHashRehash(ht)
	} else if ht.GetNTableSize() < HT_MAX_SIZE {
		var new_data any
		var old_data any = HT_GET_DATA_ADDR(ht)
		var nSize uint32 = ht.GetNTableSize() + ht.GetNTableSize()
		var old_buckets *Bucket = ht.GetArData()
		ht.SetNTableSize(nSize)
		new_data = Pemalloc(HT_SIZE_EX(nSize, HT_SIZE_TO_MASK(nSize)), ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
		ht.SetNTableMask(HT_SIZE_TO_MASK(ht.GetNTableSize()))
		HT_SET_DATA_ADDR(ht, new_data)
		memcpy(ht.GetArData(), old_buckets, b.SizeOf("Bucket")*ht.GetNNumUsed())
		Pefree(old_data, ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
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
		ht.SetNNumUsed(0)
		HT_HASH_RESET(ht)
		return SUCCESS
	}
	HT_HASH_RESET(ht)
	i = 0
	p = ht.GetArData()
	if ht.IsWithoutHoles() {
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
	if prev != nil {
		prev.GetVal().GetNext() = p.GetVal().GetNext()
	} else {
		HT_HASH(ht, p.GetH()|ht.GetNTableMask()) = p.GetVal().GetNext()
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
	var nIndex uint32 = p.GetH() | ht.GetNTableMask()
	var i uint32 = HT_HASH(ht, nIndex)
	if i != idx {
		prev = HT_HASH_TO_BUCKET(ht, i)
		for prev.GetVal().GetNext() != idx {
			i = prev.GetVal().GetNext()
			prev = HT_HASH_TO_BUCKET(ht, i)
		}
	}
	_zendHashDelElEx(ht, idx, p, prev)
}
func ZendHashDelBucket(ht *HashTable, p *Bucket) {
	ht.assertRc1()
	_zendHashDelEl(ht, HT_IDX_TO_HASH(p-ht.GetArData()), p)
}
func ZendHashDel(ht *HashTable, key *ZendString) int {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	ht.assertRc1()
	h = key.GetHash()
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
	ht.assertRc1()
	h = key.GetHash()
	nIndex = h | ht.GetNTableMask()
	idx = HT_HASH(ht, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(ht, idx)
		if p.GetKey() == key || p.GetH() == h && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), key) != 0 {
			if p.GetVal().IsType(IS_INDIRECT) {
				var data *Zval = p.GetVal().GetZv()
				if data.IsType(IS_UNDEF) {
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
					ht.AddUFlags(HASH_FLAG_HAS_EMPTY_IND)
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
	ht.assertRc1()
	h = ZendInlineHashFunc(str, len_)
	nIndex = h | ht.GetNTableMask()
	idx = HT_HASH(ht, nIndex)
	for idx != HT_INVALID_IDX {
		p = HT_HASH_TO_BUCKET(ht, idx)
		if p.GetH() == h && p.GetKey() != nil && p.GetKey().GetLen() == len_ && !(memcmp(p.GetKey().GetVal(), str, len_)) {
			if p.GetVal().IsType(IS_INDIRECT) {
				var data *Zval = p.GetVal().GetZv()
				if data.IsType(IS_UNDEF) {
					return FAILURE
				} else {
					if ht.GetPDestructor() != nil {
						ht.GetPDestructor()(data)
					}
					ZVAL_UNDEF(data)
					ht.AddUFlags(HASH_FLAG_HAS_EMPTY_IND)
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
	ht.assertRc1()
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
	ht.assertRc1()
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
			if ht.IsStaticKeys() {
				if ht.IsWithoutHoles() {
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
			} else if ht.IsWithoutHoles() {
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
			if !(ht.IsStaticKeys()) {
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
	}
	Pefree(HT_GET_DATA_ADDR(ht), ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
}
func ZendArrayDestroy(ht *HashTable) {
	var p *Bucket
	var end *Bucket

	/* break possible cycles */

	GC_REMOVE_FROM_BUFFER(ht)
	ht.GetGcTypeInfo() = IS_NULL
	if ht.GetNNumUsed() != 0 {

		/* In some rare cases destructors of regular arrays may be changed */

		if ht.GetPDestructor() != ZVAL_PTR_DTOR {
			ZendHashDestroy(ht)
			goto free_ht
		}
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if ht.IsStaticKeys() {
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
	}
	Efree(HT_GET_DATA_ADDR(ht))
free_ht:
	ZendHashIteratorsRemove(ht)
	FREE_HASHTABLE(ht)
}
func ZendHashClean(ht *HashTable) {
	var p *Bucket
	var end *Bucket
	ht.assertRc1()
	if ht.GetNNumUsed() != 0 {
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if ht.GetPDestructor() != nil {
			if ht.IsStaticKeys() {
				if ht.IsWithoutHoles() {
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
			} else if ht.IsWithoutHoles() {
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
			if !(ht.IsStaticKeys()) {
				if ht.IsWithoutHoles() {
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
		HT_HASH_RESET(ht)
	}
	ht.SetNNumUsed(0)
	ht.SetNNumOfElements(0)
	ht.SetNNextFreeElement(0)
	ht.SetNInternalPointer(0)
}
func ZendSymtableClean(ht *HashTable) {
	var p *Bucket
	var end *Bucket
	ht.assertRc1()
	if ht.GetNNumUsed() != 0 {
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if ht.IsStaticKeys() {
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
		HT_HASH_RESET(ht)
	}
	ht.SetNNumUsed(0)
	ht.SetNNumOfElements(0)
	ht.SetNNextFreeElement(0)
	ht.SetNInternalPointer(0)
}
func ZendHashGracefulReverseDestroy(ht *HashTable) {
	var idx uint32
	var p *Bucket
	ht.assertRc1()
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
			ht.assertRc1()
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
			ht.assertRc1()
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
			ht.assertRc1()
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
			ht.assertRc1()
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
	target.assertRc1()
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
			q.GetKey().AddRefcount()
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
	target.SetRefcount(1)
	target.GetGcTypeInfo() = IS_ARRAY | GC_COLLECTABLE<<GC_FLAGS_SHIFT
	target.SetPDestructor(ZVAL_PTR_DTOR)
	if source.GetNNumOfElements() == 0 {
		target.SetNTableMask(HT_MIN_MASK)
		target.SetNNumUsed(0)
		target.SetNNumOfElements(0)
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNInternalPointer(0)
		target.SetNTableSize(HT_MIN_SIZE)
		HT_SET_DATA_ADDR(target, &UninitializedBucket)
	} else if (source.GetGcFlags() & IS_ARRAY_IMMUTABLE) != 0 {
		target.SetUFlags(source.GetUFlags() & HASH_FLAG_MASK)
		target.SetNNumUsed(source.GetNNumUsed())
		target.SetNNumOfElements(source.GetNNumOfElements())
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNTableSize(source.GetNTableSize())
		HT_SET_DATA_ADDR(target, Emalloc(HT_SIZE(target)))
		target.SetNInternalPointer(source.GetNInternalPointer())
		memcpy(HT_GET_DATA_ADDR(target), HT_GET_DATA_ADDR(source), HT_USED_SIZE(source))
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
		HT_SET_DATA_ADDR(target, Emalloc(HT_SIZE(target)))
		HT_HASH_RESET(target)
		if target.IsStaticKeys() {
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
func ZendHashMerge(target *HashTable, source *HashTable, pCopyConstructor CopyCtorFuncT, overwrite ZendBool) {
	var idx uint32
	var p *Bucket
	var t *Zval
	var s *Zval
	target.assertRc1()
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
			if s.IsType(IS_INDIRECT) {
				s = s.GetZv()
			}
			if s.IsType(IS_UNDEF) {
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
	target.assertRc1()
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
	p = ZendHashFindBucket(ht, key)
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
	var p *Bucket
	p = ZendHashStrFindBucket(ht, str, len_)
	if p != nil {
		return p.GetVal()
	} else {
		return nil
	}
}
func ZendHashIndexFind(ht *HashTable, h ZendUlong) *Zval {
	var p *Bucket
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
	ht.assertRc1()
	if ht.GetNNumOfElements() <= 1 && !(renumber != 0 && ht.GetNNumOfElements() > 0) {
		return SUCCESS
	}
	if ht.IsWithoutHoles() {
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
		if ht.HasUFlags(HASH_FLAG_PACKED) {
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
	if renumber != 0 {
		var new_data any
		var old_data any = HT_GET_DATA_ADDR(ht)
		var old_buckets *Bucket = ht.GetArData()
		new_data = Pemalloc(HT_SIZE_EX(ht.GetNTableSize(), HT_MIN_MASK), ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
		ht.AddUFlags(HASH_FLAG_PACKED | HASH_FLAG_STATIC_KEYS)
		ht.SetNTableMask(HT_MIN_MASK)
		HT_SET_DATA_ADDR(ht, new_data)
		memcpy(ht.GetArData(), old_buckets, b.SizeOf("Bucket")*ht.GetNNumUsed())
		Pefree(old_data, ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
		HT_HASH_RESET_PACKED(ht)
	} else {
		ZendHashRehash(ht)
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
	if (ht1.GetGcFlags() & GC_IMMUTABLE) == 0 {
		GC_PROTECT_RECURSION(ht1)
	}
	result = ZendHashCompareImpl(ht1, ht2, compar, ordered)
	if (ht1.GetGcFlags() & GC_IMMUTABLE) == 0 {
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
	if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
		ht.AddRefcount()
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
				str_key.DelRefcount()
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
	if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
		ht.AddRefcount()
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
