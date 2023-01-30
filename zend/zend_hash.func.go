// <<generate>>

package zend

import (
	b "sik/builtin"
)

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
func ZEND_HASH_INDEX_FIND(_ht *HashTable, _h ZendUlong, _ret *Zval, _not_found __auto__) {
	_ret = ZendHashIndexFind(_ht, _h)
	if _ret == nil {
		goto _not_found
	}
}
func ZendHashExists(ht *HashTable, key *ZendString) ZendBool {
	var exists = ht.ExistsByZendString(key)
	return intBool(exists)
}
func ZendHashStrExists(ht *HashTable, str *byte, len_ int) ZendBool {
	var exists = ht.ExistsByStrPtr(str, len_)
	return intBool(exists)
}
func ZendHashIndexExists(ht *HashTable, h ZendUlong) ZendBool {
	var exists = ht.ExistsByIndex(int(h))
	return intBool(exists)
}
func ZendHashHasMoreElementsEx(ht *HashTable, pos *HashPosition) ZEND_RESULT_CODE {
	if ZendHashGetCurrentKeyTypeEx(ht, pos) == HASH_KEY_NON_EXISTENT {
		return FAILURE
	} else {
		return SUCCESS
	}
}
func ZendHashMoveForward(ht *HashTable) int {
	return ZendHashMoveForwardEx(ht, &ht.nInternalPointer)
}
func ZendHashMoveBackwards(ht *HashTable) int {
	return ZendHashMoveBackwardsEx(ht, &ht.nInternalPointer)
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
	if ht.HasIterators() {
		_zendHashIteratorsUpdate(ht, from, to)
	}
}
func ZEND_HANDLE_NUMERIC_STR(key *byte, length int, idx *ZendUlong) bool {
	var str = b.CastStr(key, length)
	if number, ok := zendParseNumericStr(str); ok {
		*idx = number
		return true
	} else {
		return false
	}
}
func ZEND_HANDLE_NUMERIC(key *ZendString, idx *ZendUlong) bool {
	var str = key.GetStr()
	if number, ok := zendParseNumericStr(str); ok {
		*idx = number
		return true
	} else {
		return false
	}
}
func ZendHashFindInd(ht *HashTable, key *ZendString) *Zval {
	var zv *Zval
	zv = ht.FindByZendString(key)
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
	zv = ht.FindByZendString(key)
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
	zv = ht.FindByZendString(key)
	return zv != nil && (zv.GetType() != IS_INDIRECT || Z_INDIRECT_P(zv).GetType() != IS_UNDEF)
}
func ZendHashStrExistsInd(ht *HashTable, str string, len_ int) int {
	var zv *Zval
	zv = ht.FindByStrPtr(str, len_)
	return zv != nil && (zv.GetType() != IS_INDIRECT || Z_INDIRECT_P(zv).GetType() != IS_UNDEF)
}
func ZendSymtableAddNew(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	if idx, ok := zendParseNumericStr(key.GetStr()); ok {
		return ZendHashIndexAddNew(ht, idx, pData)
	} else {
		return ZendHashAddNew(ht, key, pData)
	}
}
func ZendSymtableUpdate(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	if idx, ok := zendParseNumericStr(key.GetStr()); ok {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashUpdate(ht, key, pData)
	}
}
func ZendSymtableUpdateInd(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	if idx, ok := zendParseNumericStr(key.GetStr()); ok {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashUpdateInd(ht, key, pData)
	}
}
func ZendSymtableDel(ht *HashTable, key *ZendString) int {
	if idx, ok := zendParseNumericStr(key.GetStr()); ok {
		return ZendHashIndexDel(ht, idx)
	} else {
		return ZendHashDel(ht, key)
	}
}
func ZendSymtableFind(ht *HashTable, key *ZendString) *Zval {
	if idx, ok := zendParseNumericStr(key.GetStr()); ok {
		return ZendHashIndexFind(ht, idx)
	} else {

		return ht.FindByZendString(key)

	}
}
func ZendSymtableExists(ht *HashTable, key *ZendString) int {
	if idx, ok := zendParseNumericStr(key.GetStr()); ok {
		return ZendHashIndexExists(ht, idx)
	} else {
		return ZendHashExists(ht, key)
	}
}
func ZendSymtableExistsInd(ht *HashTable, key *ZendString) int {
	if idx, ok := zendParseNumericStr(key.GetStr()); ok {
		return ZendHashIndexExists(ht, idx)
	} else {
		return ZendHashExistsInd(ht, key)
	}
}
func ZendSymtableStrUpdate(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var str_ = b.CastStr(str, len_)
	if idx, ok := zendParseNumericStr(str_); ok {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashStrUpdate(ht, str, len_, pData)
	}
}
func ZendSymtableStrUpdateInd(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var str_ = b.CastStr(str, len_)
	if idx, ok := zendParseNumericStr(str_); ok {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashStrUpdateInd(ht, str, len_, pData)
	}
}
func ZendSymtableStrDel(ht *HashTable, str *byte, len_ int) int {
	var str_ = b.CastStr(str, len_)
	if idx, ok := zendParseNumericStr(str_); ok {
		return ZendHashIndexDel(ht, idx)
	} else {
		return ZendHashStrDel(ht, str, len_)
	}
}
func ZendSymtableStrFind(ht *HashTable, str *byte, len_ int) *Zval {
	var str_ = b.CastStr(str, len_)
	if idx, ok := zendParseNumericStr(str_); ok {
		return ZendHashIndexFind(ht, idx)
	} else {
		return ht.FindByStrPtr(str, len_)
	}
}
func ZendSymtableStrExists(ht *HashTable, str *byte, len_ int) int {
	var str_ = b.CastStr(str, len_)
	if idx, ok := zendParseNumericStr(str_); ok {
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
	zv = ht.FindByZendString(key)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashFindExPtr(ht *HashTable, key *ZendString, known_hash ZendBool) any {
	var zv *Zval
	zv = ht.FindByZendString(key)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashStrFindPtr(ht *HashTable, str string) any {
	var zv *Zval
	zv = ht.FindByStr(str)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexFindPtr(ht *HashTable, h ZendUlong) any {
	var zv = ht.FindByIndex(int(h))
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexFindDeref(ht *HashTable, h ZendUlong) *Zval {
	var zv = ht.FindByIndex(int(h))
	if zv != nil {
		ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashFindDeref(ht *HashTable, str *ZendString) *Zval {
	var zv *Zval = ht.FindByZendString(str)
	if zv != nil {
		ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashStrFindDeref(ht *HashTable, str string, len_ int) *Zval {
	var zv *Zval = ht.FindByStrPtr(str, len_)
	if zv != nil {
		ZVAL_DEREF(zv)
	}
	return zv
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

func _zendHashAppend(ht *HashTable, key *ZendString, zv *Zval) {
	var bucket = NewBucketStr(key.GetStr(), zv)
	ht.appendBucket(bucket)
}
func _zendHashAppendPtr(ht *HashTable, key *ZendString, ptr any) {
	var bucketKey = NewStrKey(key.GetStr())
	var bucket = NewBucketPtr(bucketKey, ptr)
	ht.appendBucket(bucket)
}
func _zendHashAppendInd(ht *HashTable, key *ZendString, ptr *Zval) {
	var bucketKey = NewStrKey(key.GetStr())
	var bucket = NewBucketIndirect(bucketKey, ptr)
	ht.appendBucket(bucket)
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

func ZendHashRealInit(ht *HashTable, packed ZendBool) { /* ignore simplify */ ht.RealInit() }
func ZendHashRealInitPacked(ht *HashTable)            { /* ignore simplify */ ht.RealInit() }
func ZendHashRealInitMixed(ht *HashTable)             { /* ignore simplify */ ht.RealInit() }
func ZendHashToPacked(ht *HashTable) {
	// todo 此函数不应被调用
	ZEND_ASSERT(false)
}
func _zendHashGetValidPos(ht *HashTable, pos HashPosition) HashPosition { return ht.validPosVal(pos) }
func _zendHashGetCurrentPos(ht *HashTable) HashPosition                 { return ht.currentPosVal() }
func ZendHashGetCurrentPos(ht *HashTable) HashPosition                  { return ht.currentPosVal() }
func ZendHashIteratorAdd(ht *HashTable, pos HashPosition) uint32 {
	var iter *HashTableIterator = __EG().GetHtIterators()
	var end *HashTableIterator = iter + __EG().GetHtIteratorsCount()
	var idx uint32
	if !(ht.IsIteratorsOverflow()) {
		ht.IncNIteratorsCount()
	}
	for iter != end {
		if iter.GetHt() == nil {
			iter.SetHt(ht)
			iter.SetPos(pos)
			idx = iter - __EG().GetHtIterators()
			if idx+1 > __EG().GetHtIteratorsUsed() {
				__EG().SetHtIteratorsUsed(idx + 1)
			}
			return idx
		}
		iter++
	}
	if __EG().GetHtIterators() == __EG().GetHtIteratorsSlots() {
		__EG().SetHtIterators(Emalloc(b.SizeOf("HashTableIterator") * (__EG().GetHtIteratorsCount() + 8)))
		memcpy(__EG().GetHtIterators(), __EG().GetHtIteratorsSlots(), b.SizeOf("HashTableIterator")*__EG().GetHtIteratorsCount())
	} else {
		__EG().SetHtIterators(Erealloc(__EG().GetHtIterators(), b.SizeOf("HashTableIterator")*(__EG().GetHtIteratorsCount()+8)))
	}
	iter = __EG().GetHtIterators() + __EG().GetHtIteratorsCount()
	__EG().SetHtIteratorsCount(__EG().GetHtIteratorsCount() + 8)
	iter.SetHt(ht)
	iter.SetPos(pos)
	memset(iter+1, 0, b.SizeOf("HashTableIterator")*7)
	idx = iter - __EG().GetHtIterators()
	__EG().SetHtIteratorsUsed(idx + 1)
	return idx
}
func ZendHashIteratorPos(idx uint32, ht *HashTable) HashPosition {
	var iter *HashTableIterator = __EG().GetHtIterators() + idx
	ZEND_ASSERT(idx != uint32-1)
	if iter.GetHt() != ht {
		if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(iter.GetHt().IsIteratorsOverflow()) {
			iter.GetHt().DecNIteratorsCount()
		}
		if !(ht.IsIteratorsOverflow()) {
			ht.IncNIteratorsCount()
		}
		iter.SetHt(ht)
		iter.SetPos(ht.currentPosVal())
	}
	return iter.GetPos()
}
func ZendHashIteratorPosEx(idx uint32, array *Zval) HashPosition {
	var ht *HashTable = array.GetArr()
	var iter *HashTableIterator = __EG().GetHtIterators() + idx
	ZEND_ASSERT(idx != uint32-1)
	if iter.GetHt() != ht {
		if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(ht.IsIteratorsOverflow()) {
			iter.GetHt().DecNIteratorsCount()
		}
		SEPARATE_ARRAY(array)
		ht = array.GetArr()
		if !(ht.IsIteratorsOverflow()) {
			ht.IncNIteratorsCount()
		}
		iter.SetHt(ht)
		iter.SetPos(ht.currentPosVal())
	}
	return iter.GetPos()
}
func ZendHashIteratorDel(idx uint32) {
	var iter *HashTableIterator = __EG().GetHtIterators() + idx
	ZEND_ASSERT(idx != uint32-1)
	if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(iter.GetHt().IsIteratorsOverflow()) {
		ZEND_ASSERT(iter.GetHt().GetNIteratorsCount() != 0)
		iter.GetHt().DecNIteratorsCount()
	}
	iter.SetHt(nil)
	if idx == __EG().GetHtIteratorsUsed()-1 {
		for idx > 0 && __EG().GetHtIterators()[idx-1].GetHt() == nil {
			idx--
		}
		__EG().SetHtIteratorsUsed(idx)
	}
}
func _zendHashIteratorsRemove(ht *HashTable) {
	var iter *HashTableIterator = __EG().GetHtIterators()
	var end *HashTableIterator = iter + __EG().GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht {
			iter.SetHt(HT_POISONED_PTR)
		}
		iter++
	}
}
func ZendHashIteratorsRemove(ht *HashTable) {
	if ht.HasIterators() {
		_zendHashIteratorsRemove(ht)
	}
}
func ZendHashIteratorsLowerPos(ht *HashTable, start HashPosition) HashPosition {
	var iter *HashTableIterator = __EG().GetHtIterators()
	var end *HashTableIterator = iter + __EG().GetHtIteratorsUsed()
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
	var iter *HashTableIterator = __EG().GetHtIterators()
	var end *HashTableIterator = iter + __EG().GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht && iter.GetPos() == from {
			iter.SetPos(to)
		}
		iter++
	}
}
func ZendHashIteratorsAdvance(ht *HashTable, step HashPosition) {
	var iter *HashTableIterator = __EG().GetHtIterators()
	var end *HashTableIterator = iter + __EG().GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht {
			iter.SetPos(iter.GetPos() + step)
		}
		iter++
	}
}

func ZendHashAdd(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return ht.addOrUpdateByZendString(key, pData, HASH_ADD)
}
func ZendHashUpdate(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return ht.addOrUpdateByZendString(key, pData, HASH_UPDATE)
}
func ZendHashUpdateInd(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return ht.addOrUpdateByZendString(key, pData, HASH_UPDATE|HASH_UPDATE_INDIRECT)
}
func ZendHashAddNew(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return ht.addOrUpdateByZendString(key, pData, HASH_ADD_NEW)
}

func ZendHashStrUpdate(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	return ht.addOrUpdateByStrPtr(str, len_, pData, HASH_UPDATE)
}
func ZendHashStrUpdateInd(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	return ht.addOrUpdateByStrPtr(str, len_, pData, HASH_UPDATE|HASH_UPDATE_INDIRECT)
}
func ZendHashStrAdd(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	return ht.addOrUpdateByStrPtr(str, len_, pData, HASH_ADD)
}
func ZendHashStrAddNew(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	return ht.addOrUpdateByStrPtr(str, len_, pData, HASH_ADD_NEW)
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
func ZendHashIndexAdd(ht *HashTable, h ZendUlong, pData *Zval) *Zval {
	return ht.IndexAddH(h, pData)
}
func ZendHashIndexAddNew(ht *HashTable, h ZendUlong, pData *Zval) *Zval {
	return ht.IndexAddNewH(h, pData)
}
func ZendHashIndexUpdate(ht *HashTable, h ZendUlong, pData *Zval) *Zval {
	return ht.IndexUpdateH(h, pData)
}
func ZendHashNextIndexInsert(ht *HashTable, pData *Zval) *Zval {
	return ht.NextIndexInsert(pData)
}
func ZendHashNextIndexInsertNew(ht *HashTable, pData *Zval) *Zval {
	return ht.NextIndexInsertNew(pData)
}
func ZendHashSetBucketKey(ht *HashTable, b *Bucket, key *ZendString) *Zval {
	var nIndex uint32
	var idx uint32
	var i uint32
	var p *Bucket
	var arData *Bucket
	ht.assertRc1()
	p = ht.FindBucketByZendString(key)
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
func ZendHashRehash(ht *HashTable) { ht.Rehash() }
func _zendHashDelElEx(ht *HashTable, idx uint32, p *Bucket, prev *Bucket) {
	if prev != nil {
		prev.GetVal().GetNext() = p.GetVal().GetNext()
	} else {
		HT_HASH(ht, p.GetH()|ht.GetNTableMask()) = p.GetVal().GetNext()
	}
	idx = HT_HASH_TO_IDX(idx)
	ht.GetNNumOfElements()--
	if ht.GetNInternalPointer() == idx || ht.HasIterators() {
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
		ht.resetHash()
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
		ht.resetHash()
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
		target.resetHash()
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
			if p.IsStrKey() {
				t = target.KeyUpdateIndirect(p.StrKey(), s)
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
			if p.IsStrKey() {
				t = target.KeyAddIndirect(p.StrKey(), s)
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
func ZendHashStrFind(ht *HashTable, str *byte, len_ int) *Zval { return ht.FindByStrPtr(str, len_) }
func ZendHashIndexFind(ht *HashTable, h ZendUlong) *Zval       { return ht.FindByIndex(int(h)) }
func ZendHashInternalPointerResetEx(ht *HashTable, pos *HashPosition) {
	*pos = _zendHashGetValidPos(ht, 0)
}
func ZendHashInternalPointerEndEx(ht *HashTable, pos *HashPosition) {
	var idx uint32
	idx = ht.GetNNumUsed()
	for idx > 0 {
		idx--
		if ht.data[idx].GetVal().GetType() != IS_UNDEF {
			*pos = idx
			return
		}
	}
	*pos = ht.GetNNumUsed()
}

// 查找下一个有效位置
func ZendHashMoveForwardEx(ht *HashTable, pos *HashPosition) int {
	if idx, ok := ht.validPos(*pos); ok {
		*pos, _ = ht.validPos(idx + 1)
		return SUCCESS
	}
	return FAILURE
}

func ZendHashMoveBackwardsEx(ht *HashTable, pos *HashPosition) int {
	var idx uint32 = *pos
	if idx < ht.GetNNumUsed() {
		for idx > 0 {
			idx--
			if ht.data[idx].GetVal().GetType() != IS_UNDEF {
				*pos = idx
				return SUCCESS
			}
		}
		*pos = ht.GetNNumUsed()
		return SUCCESS
	}
	return FAILURE
}
func ZendHashGetCurrentKeyEx(ht *HashTable, str_index **ZendString, num_index *ZendUlong, pos *HashPosition) int {
	var idx uint32
	var p *Bucket
	idx = ht.validPosVal(*pos)
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
	idx = ht.validPosVal(*pos)
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
	idx = ht.validPosVal(*pos)
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
	idx = ht.validPosVal(*pos)
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
				pData2 = ht2.FindByZendString(p1.GetKey())
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
	idx = ht.validPosVal(0)
	res = ht.GetArData() + idx
	for ; idx < ht.GetNNumUsed(); idx++ {
		p = ht.GetArData() + idx
		if p.IsValid() {
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

func zendParseNumericStr(str string) (ZendUlong, bool) {
	// 首字符非数字快速失败
	if len(str) == 0 {
		return 0, false
	}
	if (str[0] < '9' || str[0] > '0') && str[0] != '-' {
		return 0, false
	}

	// 字符串转数字
	var length = len(str)
	var i = 0
	if str[i] == '-' {
		i++
	}
	if (length > 1 && str[i] == '0') /* numbers with leading zeros */ ||
		(length-i > MAX_LENGTH_OF_LONG-1) /* number too long */ {
		return 0, false
	}

	var number uint = 0
	for _, c := range str[i:] {
		if c >= '0' && c <= '9' {
			number = number*10 + uint(c-'0')
		} else {
			return 0, false
		}
	}

	// 处理符号和 overflow
	if str[0] == '-' {
		if number-1 > ZEND_LONG_MAX {
			return 0, false
		}
		number = -number
	} else {
		if number > ZEND_LONG_MAX {
			return 0, false
		}
	}

	return number, true
}

func ZendSymtableToProptable(ht *HashTable) *HashTable {
	var num_key ZendUlong
	var str_key *ZendString
	var zv *Zval

	var __ht *HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		str_key = _p.GetKey()
		if str_key == nil {
			goto convert
		}
	}
	if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
		ht.AddRefcount()
	}
	return ht
convert:
	var new_ht *HashTable = ZendNewArray(ht.GetNNumOfElements())
	var __ht__1 *HashTable = ht
	for _, _p := range __ht__1.foreachData() {
		var _z *Zval = _p.GetVal()

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
	return new_ht
}
func ZendProptableToSymtable(ht *HashTable, always_duplicate ZendBool) *HashTable {
	var num_key ZendUlong
	var str_key *ZendString
	var zv *Zval
	var __ht *HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		str_key = _p.GetKey()

		/* The `str_key &&` here might seem redundant: property tables should
		 * only have string keys. Unfortunately, this isn't true, at the very
		 * least because of ArrayObject, which stores a symtable where the
		 * property table should be.
		 */

		if str_key != nil && ZEND_HANDLE_NUMERIC(str_key, &num_key) {
			goto convert
		}

		/* The `str_key &&` here might seem redundant: property tables should
		 * only have string keys. Unfortunately, this isn't true, at the very
		 * least because of ArrayObject, which stores a symtable where the
		 * property table should be.
		 */

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
	var __ht__1 *HashTable = ht
	for _, _p := range __ht__1.foreachData() {
		var _z *Zval = _p.GetVal()
		if _z.IsType(IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(IS_UNDEF) {
				continue
			}
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

		if str_key == nil || ZEND_HANDLE_NUMERIC(str_key, &num_key) {
			ZendHashIndexUpdate(new_ht, num_key, zv)
		} else {
			ZendHashUpdate(new_ht, str_key, zv)
		}

		/* Again, thank ArrayObject for `!str_key ||`. */

	}
	return new_ht
}
