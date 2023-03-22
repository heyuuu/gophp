package types

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/faults"
)

func NewEmptyArray() *Array {
	return NewZendArray(0)
}

func ZVAL_EMPTY_ARRAY(z *Zval) {
	z.SetArr((*Array)(&ZendEmptyArray))
	z.SetTypeInfo(IS_ARRAY)
}
func ZendHashInit(ht *Array, nSize uint32, pHashFunction any, pDestructor DtorFuncT, persistent ZendBool) {
	*ht = *NewZendArrayEx(nSize, pDestructor, persistent != 0)
}
func ZendHashInitEx(ht *Array, nSize uint32, pHashFunction any, pDestructor DtorFuncT, persistent ZendBool, bApplyProtection int) {
	*ht = *NewZendArrayEx(nSize, pDestructor, persistent != 0)
}
func ZEND_HASH_INDEX_FIND(_ht *Array, _h zend.ZendUlong, _ret *Zval, _not_found __auto__) {
	_ret = _ht.IndexFindH(_h)
	if _ret == nil {
		goto _not_found
	}
}
func ZendHashExists(ht *Array, key *String) ZendBool {
	var exists = ht.KeyExists(key.GetStr())
	return IntBool(exists)
}
func ZendHashStrExists(ht *Array, str *byte, len_ int) ZendBool {
	var exists = ht.KeyExists(b.CastStr(str, len_))
	return IntBool(exists)
}
func ZendHashIndexExists(ht *Array, h zend.ZendUlong) ZendBool {
	var exists = ht.IndexExists(int(h))
	return IntBool(exists)
}
func ZendHashHasMoreElementsEx(ht *Array, pos *HashPosition) ZEND_RESULT_CODE {
	if ZendHashGetCurrentKeyTypeEx(ht, pos) == HASH_KEY_NON_EXISTENT {
		return FAILURE
	} else {
		return SUCCESS
	}
}
func ZendHashMoveForward(ht *Array) int {
	return ZendHashMoveForwardEx(ht, &ht.nInternalPointer)
}
func ZendHashMoveBackwards(ht *Array) int {
	return ZendHashMoveBackwardsEx(ht, &ht.nInternalPointer)
}
func ZendHashGetCurrentKey(ht *Array, str_index **String, num_index *zend.ZendUlong) int {
	return ZendHashGetCurrentKeyEx(ht, str_index, num_index, ht.GetNInternalPointer())
}
func ZendHashGetCurrentKeyZval(ht *Array, key *Zval) {
	ZendHashGetCurrentKeyZvalEx(ht, key, ht.GetNInternalPointer())
}
func ZendHashGetCurrentData(ht *Array) *Zval {
	return ZendHashGetCurrentDataEx(ht, ht.GetNInternalPointer())
}
func ZendHashInternalPointerReset(ht *Array) {
	ZendHashInternalPointerResetEx(ht, ht.GetNInternalPointer())
}
func ZendHashInternalPointerEnd(ht *Array) {
	ZendHashInternalPointerEndEx(ht, ht.GetNInternalPointer())
}
func ZendHashSort(ht *Array, compare_func CompareFuncT, renumber ZendBool) int {
	return ht.SortCompatible(compare_func, renumber)
}
func ZendNewArray(size uint32) *Array { return NewZendArray(size) }
func ZendHashIteratorsUpdate(ht *Array, from HashPosition, to HashPosition) {
	if ht.HasIterators() {
		_zendHashIteratorsUpdate(ht, from, to)
	}
}
func ZEND_HANDLE_NUMERIC_STR(key *byte, length int, idx *zend.ZendUlong) bool {
	var str = b.CastStr(key, length)
	if number, ok := ZendParseNumericStr(str); ok {
		*idx = zend.ZendUlong(number)
		return true
	} else {
		return false
	}
}
func ZEND_HANDLE_NUMERIC(key *String, idx *zend.ZendUlong) bool {
	var str = key.GetStr()
	if number, ok := ZendParseNumericStr(str); ok {
		*idx = zend.ZendUlong(number)
		return true
	} else {
		return false
	}
}
func ZendHashFindInd(ht *Array, key *String) *Zval {
	var zv *Zval
	zv = ht.KeyFind(key.GetStr())
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
func ZendHashFindExInd(ht *Array, key *String, known_hash ZendBool) *Zval {
	var zv *Zval
	zv = ht.KeyFind(key.GetStr())
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

func ZendHashAddPtr(ht *Array, key *String, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.KeyAdd(key.GetStr(), &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashAddNewPtr(ht *Array, key *String, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.KeyAddNew(key.GetStr(), &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashStrAddPtr(ht *Array, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.KeyAdd(b.CastStr(str, len_), &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashUpdatePtr(ht *Array, key *String, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.KeyUpdate(key.GetStr(), &tmp)
	return zv.GetPtr()
}
func ZendHashStrUpdatePtr(ht *Array, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.KeyUpdate(b.CastStr(str, len_), &tmp)
	return zv.GetPtr()
}
func ZendHashAddMem(ht *Array, key *String, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ht.KeyAdd(key.GetStr(), &tmp)) {
		zv.SetPtr(zend.Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashStrAddMem(ht *Array, str *byte, len_ int, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ht.KeyAdd(b.CastStr(str, len_), &tmp)) {
		zv.SetPtr(zend.Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashUpdateMem(ht *Array, key *String, pData any, size int) any {
	var p any
	p = zend.Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashUpdatePtr(ht, key, p)
}
func ZendHashStrUpdateMem(ht *Array, str *byte, len_ int, pData any, size int) any {
	var p any
	p = zend.Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashStrUpdatePtr(ht, str, len_, p)
}
func ZendHashIndexAddPtr(ht *Array, h zend.ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.IndexAddH(h, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexAddNewPtr(ht *Array, h zend.ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.IndexAddNewH(h, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexUpdatePtr(ht *Array, h zend.ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.IndexUpdateH(h, &tmp)
	return zv.GetPtr()
}
func ZendHashIndexAddMem(ht *Array, h zend.ZendUlong, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ht.IndexAddH(h, &tmp)) {
		zv.SetPtr(zend.Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashNextIndexInsertPtr(ht *Array, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.NextIndexInsert(&tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexUpdateMem(ht *Array, h zend.ZendUlong, pData any, size int) any {
	var p any
	p = zend.Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashIndexUpdatePtr(ht, h, p)
}
func ZendHashNextIndexInsertMem(ht *Array, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ht.NextIndexInsert(&tmp)) {
		zv.SetPtr(zend.Pemalloc(size, ht.GetGcFlags()&IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashFindPtr(ht *Array, key *String) any {
	return ht.KeyFindPtr(key.GetStr())
}
func ZendHashFindExPtr(ht *Array, key *String, known_hash ZendBool) any {
	return ht.KeyFindPtr(key.GetStr())
}
func ZendHashStrFindPtr(ht *Array, str *byte, len_ int) any {
	var zv *Zval
	zv = ht.KeyFind(b.CastStr(str, len_))
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexFindPtr(ht *Array, h zend.ZendUlong) any {
	var zv *Zval
	zv = ht.IndexFindH(h)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexFindDeref(ht *Array, h zend.ZendUlong) *Zval {
	var zv = ht.IndexFindH(h)
	if zv != nil {
		zv = ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashFindDeref(ht *Array, str *String) *Zval {
	var zv *Zval = ht.KeyFind(str.GetStr())
	if zv != nil {
		zv = ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashStrFindDeref(ht *Array, str *byte, len_ int) *Zval {
	var zv *Zval = ht.KeyFind(b.CastStr(str, len_))
	if zv != nil {
		zv = ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashGetCurrentDataPtrEx(ht *Array, pos *HashPosition) any {
	var zv *Zval
	zv = ZendHashGetCurrentDataEx(ht, pos)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashGetCurrentDataPtr(ht *Array) any {
	return ZendHashGetCurrentDataPtrEx(ht, ht.GetNInternalPointer())
}

func _zendHashAppend(ht *Array, key *String, zv *Zval) {
	var bucket = NewBucketStr(key.GetStr(), zv)
	ht.appendBucket(bucket)
}
func _zendHashAppendPtr(ht *Array, key *String, ptr any) {
	var bucketKey = NewStrKey(key.GetStr())
	var bucket = NewBucketPtr(bucketKey, ptr)
	ht.appendBucket(bucket)
}
func _zendHashAppendInd(ht *Array, key *String, ptr *Zval) {
	var bucketKey = NewStrKey(key.GetStr())
	var bucket = NewBucketIndirect(bucketKey, ptr)
	ht.appendBucket(bucket)
}
func ZendHashCheckSize(nSize uint32) uint32 {
	/* Use big enough power of 2 */

	if nSize <= HT_MIN_SIZE {
		return HT_MIN_SIZE
	} else if nSize >= HT_MAX_SIZE {
		faults.ErrorNoreturn(faults.E_ERROR, "Possible integer overflow in memory allocation (%u * %zu + %zu)", nSize, b.SizeOf("Bucket"), b.SizeOf("Bucket"))
	}
	nSize -= 1
	nSize |= nSize >> 1
	nSize |= nSize >> 2
	nSize |= nSize >> 4
	nSize |= nSize >> 8
	nSize |= nSize >> 16
	return nSize + 1
}

func ZendHashRealInit(ht *Array, packed ZendBool) { /* ignore simplify */ ht.RealInit() }
func ZendHashRealInitPacked(ht *Array)            { /* ignore simplify */ ht.RealInit() }
func ZendHashRealInitMixed(ht *Array)             { /* ignore simplify */ ht.RealInit() }
func ZendHashToPacked(ht *Array) {
	// todo 此函数不应被调用
	b.Assert(false)
}
func ZendHashIteratorAdd(ht *Array, pos HashPosition) uint32 {
	var iter *HashTableIterator = zend.EG__().GetHtIterators()
	var end *HashTableIterator = iter + zend.EG__().GetHtIteratorsCount()
	var idx uint32
	if !(ht.IsIteratorsOverflow()) {
		ht.IncNIteratorsCount()
	}

	for iter != end {
		if iter.GetHt() == nil {
			iter.SetHt(ht)
			iter.SetPos(pos)
			idx = iter - zend.EG__().GetHtIterators()
			if idx+1 > zend.EG__().GetHtIteratorsUsed() {
				zend.EG__().SetHtIteratorsUsed(idx + 1)
			}
			return idx
		}
		iter++
	}
	if zend.EG__().GetHtIterators() == zend.EG__().GetHtIteratorsSlots() {
		zend.EG__().SetHtIterators(zend.Emalloc(b.SizeOf("HashTableIterator") * (zend.EG__().GetHtIteratorsCount() + 8)))
		memcpy(zend.EG__().GetHtIterators(), zend.EG__().GetHtIteratorsSlots(), b.SizeOf("HashTableIterator")*zend.EG__().GetHtIteratorsCount())
	} else {
		zend.EG__().SetHtIterators(zend.Erealloc(zend.EG__().GetHtIterators(), b.SizeOf("HashTableIterator")*(zend.EG__().GetHtIteratorsCount()+8)))
	}
	iter = zend.EG__().GetHtIterators() + zend.EG__().GetHtIteratorsCount()
	zend.EG__().SetHtIteratorsCount(zend.EG__().GetHtIteratorsCount() + 8)
	iter.SetHt(ht)
	iter.SetPos(pos)
	memset(iter+1, 0, b.SizeOf("HashTableIterator")*7)
	idx = iter - zend.EG__().GetHtIterators()
	zend.EG__().SetHtIteratorsUsed(idx + 1)
	return idx
}
func ZendHashIteratorPos(idx uint32, ht *Array) HashPosition {
	var iter *HashTableIterator = zend.EG__().GetHtIterators() + idx
	b.Assert(idx != uint32-1)
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
	var ht *Array = array.GetArr()
	var iter *HashTableIterator = zend.EG__().GetHtIterators() + idx
	b.Assert(idx != uint32-1)
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
	var iter *HashTableIterator = zend.EG__().GetHtIterators() + idx
	b.Assert(idx != uint32-1)
	if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(iter.GetHt().IsIteratorsOverflow()) {
		b.Assert(iter.GetHt().GetNIteratorsCount() != 0)
		iter.GetHt().DecNIteratorsCount()
	}
	iter.SetHt(nil)
	if idx == zend.EG__().GetHtIteratorsUsed()-1 {
		for idx > 0 && zend.EG__().GetHtIterators()[idx-1].GetHt() == nil {
			idx--
		}
		zend.EG__().SetHtIteratorsUsed(idx)
	}
}
func _zendHashIteratorsRemove(ht *Array) {
	var iter *HashTableIterator = zend.EG__().GetHtIterators()
	var end *HashTableIterator = iter + zend.EG__().GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht {
			iter.SetHt(HT_POISONED_PTR)
		}
		iter++
	}
}
func ZendHashIteratorsRemove(ht *Array) {
	if ht.HasIterators() {
		_zendHashIteratorsRemove(ht)
	}
}
func ZendHashIteratorsLowerPos(ht *Array, start HashPosition) HashPosition {
	var iter *HashTableIterator = zend.EG__().GetHtIterators()
	var end *HashTableIterator = iter + zend.EG__().GetHtIteratorsUsed()
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
func _zendHashIteratorsUpdate(ht *Array, from HashPosition, to HashPosition) {
	var iter *HashTableIterator = zend.EG__().GetHtIterators()
	var end *HashTableIterator = iter + zend.EG__().GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht && iter.GetPos() == from {
			iter.SetPos(to)
		}
		iter++
	}
}
func ZendHashIteratorsAdvance(ht *Array, step HashPosition) {
	var iter *HashTableIterator = zend.EG__().GetHtIterators()
	var end *HashTableIterator = iter + zend.EG__().GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht {
			iter.SetPos(iter.GetPos() + step)
		}
		iter++
	}
}

func ZendHashIndexAddEmptyElement(ht *Array, h zend.ZendUlong) *Zval {
	var dummy Zval
	(&dummy).SetUndef()
	return ht.IndexAddH(h, &dummy)
}
func ZendHashAddEmptyElement(ht *Array, key *String) *Zval {
	var dummy Zval
	(&dummy).SetUndef()
	return ht.KeyAdd(key.GetStr(), &dummy)
}
func ZendHashStrAddEmptyElement(ht *Array, str *byte, len_ int) *Zval {
	var dummy Zval
	(&dummy).SetUndef()
	return ht.KeyAdd(b.CastStr(str, len_), &dummy)
}
func ZendHashSetBucketKey(ht *Array, b *Bucket, key *String) *Zval {
	return ht.SetBucketKey(b, key.GetStr())
}
func ZendHashDelBucket(ht *Array, p *Bucket) {
	ht.assertRc1()
	// todo 调整为传入 pos 更合适
	if pos, ok := ht.posBucket(p); ok {
		ht.deleteBucket(pos)
	}
}
func ZendHashDel(ht *Array, key *String) int {
	var strKey = key.GetStr()
	if ht.KeyDelete(strKey) {
		return SUCCESS
	}
	return FAILURE
}
func ZendHashDelInd(ht *Array, key *String) int {
	var strKey = key.GetStr()
	if ht.KeyDeleteIndirect(strKey) {
		return SUCCESS
	}
	return FAILURE
}
func ZendHashStrDel(ht *Array, str *byte, len_ int) int {
	var strKey = b.CastStr(str, len_)
	if ht.KeyDelete(strKey) {
		return SUCCESS
	}
	return FAILURE
}
func ZendHashIndexDel(ht *Array, h zend.ZendUlong) int {
	var index = int(h)
	if ht.IndexDelete(index) {
		return SUCCESS
	}
	return FAILURE
}

func ZendHashApply(ht *Array, apply_func ApplyFuncT) {
	ht.applyValidBucket(func(p *Bucket) int {
		return apply_func(p.GetVal())
	})
}
func ZendHashApplyWithArgument(ht *Array, apply_func ApplyFuncArgT, argument any) {
	ht.applyValidBucket(func(p *Bucket) int {
		return apply_func(p.GetVal(), argument)
	})
}
func ZendHashApplyWithArguments(ht *Array, apply_func ApplyFuncArgsT, num_args int, args ...any) {
	ht.applyValidBucket(func(p *Bucket) int {
		var hash_key = p.key.GetZendHashKey()
		return apply_func(p.GetVal(), num_args, args, &hash_key)
	})
}
func ZendHashReverseApply(ht *Array, apply_func ApplyFuncT) {
	ht.applyValidBucketReserve(func(p *Bucket) int {
		return apply_func(p.GetVal())
	})
}
func ZendHashCopy(target *Array, source *Array, pCopyConstructor CopyCtorFuncT) {
	target.assertRc1()
	source.eachValidBucketIndirect(func(pos uint32, p *Bucket, data *Zval) {
		var newEntry = target.Update(p.key, data)
		if pCopyConstructor != nil {
			pCopyConstructor(newEntry)
		}
	})
}

func ZendArrayDupElements(source *Array, target *Array) {
	target.eachValidBucketIndirect(func(pos uint32, p *Bucket, data *Zval) {
		// 增加引用计数
		for {
			if data.IsRefcounted() {
				if data.IsReference() && data.GetRefcount() == 1 && (!data.GetRef().GetVal().IsArray() || data.GetRef().GetVal().GetArr() != source) {
					data = data.GetRef().GetVal()
					if !(data.IsRefcounted()) {
						break
					}
				}
				data.AddRefcount()
			}
			break
		}

		// 添加元素到新数组
		var newBucket = NewBucket(p.GetZendKey(), data)
		target.appendBucket(newBucket)

		// 更新内部指针
		if source.nInternalPointer == pos {
			target.nInternalPointer = target.LastPos()
		}
	})
}

func ZendArrayDup(source *Array) *Array {
	var target *Array = NewZendArray(source.nTableSize)
	target.AddGcFlags(GC_COLLECTABLE)
	target.nNextFreeElement = source.nNextFreeElement

	if source.GetNNumOfElements() == 0 {
		return target
	}

	target.SetFlags(source.GetFlags())

	if (source.GetGcFlags() & IS_ARRAY_IMMUTABLE) != 0 {
		target.SetNNumUsed(source.GetNNumUsed())
		target.SetNNumOfElements(source.GetNNumOfElements())
		HT_SET_DATA_ADDR(target, zend.Emalloc(HT_SIZE(target)))
		target.SetNInternalPointer(source.GetNInternalPointer())

		target.copyDataAndHash(source)
	} else {
		if source.nInternalPointer < source.DataSize() {
			target.nInternalPointer = source.nInternalPointer
		}
		target.resetDataAndHash(source.DataSize())

		ZendArrayDupElements(source, target)
		target.SetNNumOfElements(target.DataSize())
	}
	return target
}
func ZendHashMerge(target *Array, source *Array, pCopyConstructor CopyCtorFuncT, overwrite ZendBool) {
	target.assertRc1()
	if overwrite != 0 {
		source.eachValidBucketIndirect(func(pos uint32, p *Bucket, data *Zval) {
			var t = target.UpdateIndirect(p.GetZendKey(), data)
			if pCopyConstructor != nil {
				pCopyConstructor(t)
			}
		})
	} else {
		source.eachValidBucketIndirect(func(pos uint32, p *Bucket, s *Zval) {
			var t = target.AddIndirect(p.GetZendKey(), s)
			if t != nil && pCopyConstructor != nil {
				pCopyConstructor(t)
			}
		})
	}
}
func ZendHashInternalPointerResetEx(ht *Array, pos *HashPosition) {
	*pos = ht.validPosVal(0)
}
func ZendHashInternalPointerEndEx(ht *Array, pos *HashPosition) {
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
func ZendHashMoveForwardEx(ht *Array, pos *HashPosition) int {
	if idx, ok := ht.validPos(*pos); ok {
		*pos, _ = ht.validPos(idx + 1)
		return SUCCESS
	}
	return FAILURE
}

func ZendHashMoveBackwardsEx(ht *Array, pos *HashPosition) int {
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
func ZendHashGetCurrentKeyEx(ht *Array, str_index **String, num_index *zend.ZendUlong, pos *HashPosition) int {
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
func ZendHashGetCurrentKeyZvalEx(ht *Array, key *Zval, pos *HashPosition) {
	var idx uint32
	var p *Bucket
	idx = ht.validPosVal(*pos)
	if idx >= ht.GetNNumUsed() {
		key.SetNull()
	} else {
		p = ht.GetArData() + idx
		if p.GetKey() != nil {
			ZVAL_STR_COPY(key, p.GetKey())
		} else {
			key.SetLong(p.IndexKey())
		}
	}
}
func ZendHashGetCurrentKeyTypeEx(ht *Array, pos *HashPosition) int {
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
func ZendHashGetCurrentDataEx(ht *Array, pos *HashPosition) *Zval {
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
	var h zend.ZendUlong
	var key *String
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
func ZendHashCompareImpl(ht1 *Array, ht2 *Array, compar CompareFuncT, ordered ZendBool) int {
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
		if p1.GetVal().IsUndef() {
			continue
		}
		if ordered != 0 {
			for true {
				b.Assert(idx2 != ht2.GetNNumUsed())
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
				pData2 = ht2.IndexFindH(p1.GetH())
				if pData2 == nil {
					return 1
				}
			} else {
				pData2 = ht2.KeyFind(p1.GetKey().GetStr())
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
func ZendHashCompare(ht1 *Array, ht2 *Array, compar CompareFuncT, ordered ZendBool) int {
	var result int
	if ht1 == ht2 {
		return 0
	}

	/* It's enough to protect only one of the arrays.
	 * The second one may be referenced from the first and this may cause
	 * false recursion detection.
	 */

	if ht1.IsRecursive() {
		faults.ErrorNoreturn(faults.E_ERROR, "Nesting level too deep - recursive dependency?")
	}

	ht1.TryProtectRecursive()

	result = ZendHashCompareImpl(ht1, ht2, compar, ordered)

	ht1.TryUnProtectRecursive()

	return result
}
func ZendHashMinmax(ht *Array, compar CompareFuncT, flag uint32) *Zval {
	var res *Bucket
	if ht.GetNNumOfElements() == 0 {
		return nil
	}

	ht.eachValidBucket(func(pos uint32, p *Bucket) {
		if flag != 0 {
			if compar(res, p) < 0 {
				res = p
			}
		} else {
			if compar(res, p) > 0 {
				res = p
			}
		}
	})

	return res.GetVal()
}

func ZendParseNumericStr(str string) (int, bool) {
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
		(length-i > zend.MAX_LENGTH_OF_LONG-1) /* number too long */ {
		return 0, false
	}

	var number = 0
	for _, c := range str[i:] {
		if c >= '0' && c <= '9' {
			number = number*10 + int(c-'0')
		} else {
			return 0, false
		}
	}

	// 处理符号和 overflow
	if str[0] == '-' {
		if number-1 > zend.ZEND_LONG_MAX {
			return 0, false
		}
		number = -number
	} else {
		if number > zend.ZEND_LONG_MAX {
			return 0, false
		}
	}

	return number, true
}

func ZendSymtableToProptable(ht *Array) *Array {
	var num_key zend.ZendUlong
	var str_key *String
	var zv *Zval

	var __ht *Array = ht
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
	var new_ht *Array = ZendNewArray(ht.GetNNumOfElements())
	var __ht__1 *Array = ht
	for _, _p := range __ht__1.foreachData() {
		var _z *Zval = _p.GetVal()

		num_key = _p.GetH()
		str_key = _p.GetKey()
		zv = _z
		if str_key == nil {
			str_key = zend.ZendLongToStr(num_key)
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
		new_ht.KeyUpdate(str_key.GetStr(), zv)
	}
	return new_ht
}
func ZendProptableToSymtable(ht *Array, always_duplicate ZendBool) *Array {
	var num_key zend.ZendUlong
	var str_key *String
	var zv *Zval
	var __ht *Array = ht
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
	var new_ht *Array = ZendNewArray(ht.GetNNumOfElements())
	var __ht__1 *Array = ht
	for _, _p := range __ht__1.foreachData() {
		var _z *Zval = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.GetZv()
			if _z.IsUndef() {
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
			new_ht.IndexUpdateH(num_key, zv)
		} else {
			new_ht.KeyUpdate(str_key.GetStr(), zv)
		}

		/* Again, thank ArrayObject for `!str_key ||`. */

	}
	return new_ht
}
