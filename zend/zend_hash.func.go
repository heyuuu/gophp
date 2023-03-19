// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZVAL_EMPTY_ARRAY(z *types.Zval) {
	z.SetArr((*types.ZendArray)(&ZendEmptyArray))
	z.SetTypeInfo(types.IS_ARRAY)
}
func ZendHashInit(ht *types.HashTable, nSize uint32, pHashFunction any, pDestructor types.DtorFuncT, persistent types.ZendBool) {
	*ht = *types.NewZendArrayEx(nSize, pDestructor, persistent != 0)
}
func ZendHashInitEx(ht *types.HashTable, nSize uint32, pHashFunction any, pDestructor types.DtorFuncT, persistent types.ZendBool, bApplyProtection int) {
	*ht = *types.NewZendArrayEx(nSize, pDestructor, persistent != 0)
}
func ZEND_HASH_INDEX_FIND(_ht *types.HashTable, _h ZendUlong, _ret *types.Zval, _not_found __auto__) {
	_ret = _ht.IndexFindH(_h)
	if _ret == nil {
		goto _not_found
	}
}
func ZendHashExists(ht *types.HashTable, key *types.ZendString) types.ZendBool {
	var exists = ht.KeyExists(key.GetStr())
	return types.IntBool(exists)
}
func ZendHashStrExists(ht *types.HashTable, str *byte, len_ int) types.ZendBool {
	var exists = ht.KeyExists(b.CastStr(str, len_))
	return types.IntBool(exists)
}
func ZendHashIndexExists(ht *types.HashTable, h ZendUlong) types.ZendBool {
	var exists = ht.IndexExists(int(h))
	return types.IntBool(exists)
}
func ZendHashHasMoreElementsEx(ht *types.HashTable, pos *types.HashPosition) types.ZEND_RESULT_CODE {
	if ZendHashGetCurrentKeyTypeEx(ht, pos) == HASH_KEY_NON_EXISTENT {
		return types.FAILURE
	} else {
		return types.SUCCESS
	}
}
func ZendHashMoveForward(ht *types.HashTable) int {
	return ZendHashMoveForwardEx(ht, &ht.nInternalPointer)
}
func ZendHashMoveBackwards(ht *types.HashTable) int {
	return ZendHashMoveBackwardsEx(ht, &ht.nInternalPointer)
}
func ZendHashGetCurrentKey(ht *types.HashTable, str_index **types.ZendString, num_index *ZendUlong) int {
	return ZendHashGetCurrentKeyEx(ht, str_index, num_index, ht.GetNInternalPointer())
}
func ZendHashGetCurrentKeyZval(ht *types.HashTable, key *types.Zval) {
	ZendHashGetCurrentKeyZvalEx(ht, key, ht.GetNInternalPointer())
}
func ZendHashGetCurrentData(ht *types.HashTable) *types.Zval {
	return ZendHashGetCurrentDataEx(ht, ht.GetNInternalPointer())
}
func ZendHashInternalPointerReset(ht *types.HashTable) {
	ZendHashInternalPointerResetEx(ht, ht.GetNInternalPointer())
}
func ZendHashInternalPointerEnd(ht *types.HashTable) {
	ZendHashInternalPointerEndEx(ht, ht.GetNInternalPointer())
}
func ZendHashSort(ht *types.HashTable, compare_func types.CompareFuncT, renumber types.ZendBool) int {
	return ht.SortCompatible(compare_func, renumber)
}
func ZendNewArray(size uint32) *types.HashTable { return types.NewZendArray(size) }
func ZendHashIteratorsUpdate(ht *types.HashTable, from types.HashPosition, to types.HashPosition) {
	if ht.HasIterators() {
		_zendHashIteratorsUpdate(ht, from, to)
	}
}
func ZEND_HANDLE_NUMERIC_STR(key *byte, length int, idx *ZendUlong) bool {
	var str = b.CastStr(key, length)
	if number, ok := zendParseNumericStr(str); ok {
		*idx = ZendUlong(number)
		return true
	} else {
		return false
	}
}
func ZEND_HANDLE_NUMERIC(key *types.ZendString, idx *ZendUlong) bool {
	var str = key.GetStr()
	if number, ok := zendParseNumericStr(str); ok {
		*idx = ZendUlong(number)
		return true
	} else {
		return false
	}
}
func ZendHashFindInd(ht *types.HashTable, key *types.ZendString) *types.Zval {
	var zv *types.Zval
	zv = ht.KeyFind(key.GetStr())
	if zv != nil && zv.IsType(types.IS_INDIRECT) {
		if types.Z_INDIRECT_P(zv).GetType() != types.IS_UNDEF {
			return zv.GetZv()
		} else {
			return nil
		}
	} else {
		return zv
	}
}
func ZendHashFindExInd(ht *types.HashTable, key *types.ZendString, known_hash types.ZendBool) *types.Zval {
	var zv *types.Zval
	zv = ht.KeyFind(key.GetStr())
	if zv != nil && zv.IsType(types.IS_INDIRECT) {
		if types.Z_INDIRECT_P(zv).GetType() != types.IS_UNDEF {
			return zv.GetZv()
		} else {
			return nil
		}
	} else {
		return zv
	}
}

func ZendHashAddPtr(ht *types.HashTable, key *types.ZendString, pData any) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, pData)
	zv = ht.KeyAdd(key.GetStr(), &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashAddNewPtr(ht *types.HashTable, key *types.ZendString, pData any) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, pData)
	zv = ht.KeyAddNew(key.GetStr(), &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashStrAddPtr(ht *types.HashTable, str *byte, len_ int, pData any) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, pData)
	zv = ht.KeyAdd(b.CastStr(str, len_), &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashUpdatePtr(ht *types.HashTable, key *types.ZendString, pData any) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, pData)
	zv = ht.KeyUpdate(key.GetStr(), &tmp)
	return zv.GetPtr()
}
func ZendHashStrUpdatePtr(ht *types.HashTable, str *byte, len_ int, pData any) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, pData)
	zv = ht.KeyUpdate(b.CastStr(str, len_), &tmp)
	return zv.GetPtr()
}
func ZendHashAddMem(ht *types.HashTable, key *types.ZendString, pData any, size int) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ht.KeyAdd(key.GetStr(), &tmp)) {
		zv.SetPtr(Pemalloc(size, ht.GetGcFlags()&types.IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashStrAddMem(ht *types.HashTable, str *byte, len_ int, pData any, size int) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ht.KeyAdd(b.CastStr(str, len_), &tmp)) {
		zv.SetPtr(Pemalloc(size, ht.GetGcFlags()&types.IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashUpdateMem(ht *types.HashTable, key *types.ZendString, pData any, size int) any {
	var p any
	p = Pemalloc(size, ht.GetGcFlags()&types.IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashUpdatePtr(ht, key, p)
}
func ZendHashStrUpdateMem(ht *types.HashTable, str *byte, len_ int, pData any, size int) any {
	var p any
	p = Pemalloc(size, ht.GetGcFlags()&types.IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashStrUpdatePtr(ht, str, len_, p)
}
func ZendHashIndexAddPtr(ht *types.HashTable, h ZendUlong, pData any) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, pData)
	zv = ht.IndexAddH(h, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexAddNewPtr(ht *types.HashTable, h ZendUlong, pData any) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, pData)
	zv = ht.IndexAddNewH(h, &tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexUpdatePtr(ht *types.HashTable, h ZendUlong, pData any) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, pData)
	zv = ht.IndexUpdateH(h, &tmp)
	return zv.GetPtr()
}
func ZendHashIndexAddMem(ht *types.HashTable, h ZendUlong, pData any, size int) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ht.IndexAddH(h, &tmp)) {
		zv.SetPtr(Pemalloc(size, ht.GetGcFlags()&types.IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashNextIndexInsertPtr(ht *types.HashTable, pData any) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, pData)
	zv = ht.NextIndexInsert(&tmp)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexUpdateMem(ht *types.HashTable, h ZendUlong, pData any, size int) any {
	var p any
	p = Pemalloc(size, ht.GetGcFlags()&types.IS_ARRAY_PERSISTENT)
	memcpy(p, pData, size)
	return ZendHashIndexUpdatePtr(ht, h, p)
}
func ZendHashNextIndexInsertMem(ht *types.HashTable, pData any, size int) any {
	var tmp types.Zval
	var zv *types.Zval
	types.ZVAL_PTR(&tmp, nil)
	if b.Assign(&zv, ht.NextIndexInsert(&tmp)) {
		zv.SetPtr(Pemalloc(size, ht.GetGcFlags()&types.IS_ARRAY_PERSISTENT))
		memcpy(zv.GetPtr(), pData, size)
		return zv.GetPtr()
	}
	return nil
}
func ZendHashFindPtr(ht *types.HashTable, key *types.ZendString) any {
	return ht.KeyFindPtr(key.GetStr())
}
func ZendHashFindExPtr(ht *types.HashTable, key *types.ZendString, known_hash types.ZendBool) any {
	return ht.KeyFindPtr(key.GetStr())
}
func ZendHashStrFindPtr(ht *types.HashTable, str *byte, len_ int) any {
	var zv *types.Zval
	zv = ht.KeyFind(b.CastStr(str, len_))
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexFindPtr(ht *types.HashTable, h ZendUlong) any {
	var zv *types.Zval
	zv = ht.IndexFindH(h)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexFindDeref(ht *types.HashTable, h ZendUlong) *types.Zval {
	var zv = ht.IndexFindH(h)
	if zv != nil {
		zv = types.ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashFindDeref(ht *types.HashTable, str *types.ZendString) *types.Zval {
	var zv *types.Zval = ht.KeyFind(str.GetStr())
	if zv != nil {
		zv = types.ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashStrFindDeref(ht *types.HashTable, str *byte, len_ int) *types.Zval {
	var zv *types.Zval = ht.KeyFind(b.CastStr(str, len_))
	if zv != nil {
		zv = types.ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashGetCurrentDataPtrEx(ht *types.HashTable, pos *types.HashPosition) any {
	var zv *types.Zval
	zv = ZendHashGetCurrentDataEx(ht, pos)
	if zv != nil {
		return zv.GetPtr()
	} else {
		return nil
	}
}
func ZendHashGetCurrentDataPtr(ht *types.HashTable) any {
	return ZendHashGetCurrentDataPtrEx(ht, ht.GetNInternalPointer())
}

func _zendHashAppend(ht *types.HashTable, key *types.ZendString, zv *types.Zval) {
	var bucket = types.NewBucketStr(key.GetStr(), zv)
	ht.appendBucket(bucket)
}
func _zendHashAppendPtr(ht *types.HashTable, key *types.ZendString, ptr any) {
	var bucketKey = types.NewStrKey(key.GetStr())
	var bucket = types.NewBucketPtr(bucketKey, ptr)
	ht.appendBucket(bucket)
}
func _zendHashAppendInd(ht *types.HashTable, key *types.ZendString, ptr *types.Zval) {
	var bucketKey = types.NewStrKey(key.GetStr())
	var bucket = types.NewBucketIndirect(bucketKey, ptr)
	ht.appendBucket(bucket)
}
func ZendHashCheckSize(nSize uint32) uint32 {
	/* Use big enough power of 2 */

	if nSize <= types.HT_MIN_SIZE {
		return types.HT_MIN_SIZE
	} else if nSize >= types.HT_MAX_SIZE {
		faults.ZendErrorNoreturn(faults.E_ERROR, "Possible integer overflow in memory allocation (%u * %zu + %zu)", nSize, b.SizeOf("Bucket"), b.SizeOf("Bucket"))
	}
	nSize -= 1
	nSize |= nSize >> 1
	nSize |= nSize >> 2
	nSize |= nSize >> 4
	nSize |= nSize >> 8
	nSize |= nSize >> 16
	return nSize + 1
}

func ZendHashRealInit(ht *types.HashTable, packed types.ZendBool) { /* ignore simplify */ ht.RealInit() }
func ZendHashRealInitPacked(ht *types.HashTable)                  { /* ignore simplify */ ht.RealInit() }
func ZendHashRealInitMixed(ht *types.HashTable)                   { /* ignore simplify */ ht.RealInit() }
func ZendHashToPacked(ht *types.HashTable) {
	// todo 此函数不应被调用
	b.Assert(false)
}
func ZendHashIteratorAdd(ht *types.HashTable, pos types.HashPosition) uint32 {
	var iter *types.HashTableIterator = EG__().GetHtIterators()
	var end *types.HashTableIterator = iter + EG__().GetHtIteratorsCount()
	var idx uint32
	if !(ht.IsIteratorsOverflow()) {
		ht.IncNIteratorsCount()
	}

	for iter != end {
		if iter.GetHt() == nil {
			iter.SetHt(ht)
			iter.SetPos(pos)
			idx = iter - EG__().GetHtIterators()
			if idx+1 > EG__().GetHtIteratorsUsed() {
				EG__().SetHtIteratorsUsed(idx + 1)
			}
			return idx
		}
		iter++
	}
	if EG__().GetHtIterators() == EG__().GetHtIteratorsSlots() {
		EG__().SetHtIterators(Emalloc(b.SizeOf("HashTableIterator") * (EG__().GetHtIteratorsCount() + 8)))
		memcpy(EG__().GetHtIterators(), EG__().GetHtIteratorsSlots(), b.SizeOf("HashTableIterator")*EG__().GetHtIteratorsCount())
	} else {
		EG__().SetHtIterators(Erealloc(EG__().GetHtIterators(), b.SizeOf("HashTableIterator")*(EG__().GetHtIteratorsCount()+8)))
	}
	iter = EG__().GetHtIterators() + EG__().GetHtIteratorsCount()
	EG__().SetHtIteratorsCount(EG__().GetHtIteratorsCount() + 8)
	iter.SetHt(ht)
	iter.SetPos(pos)
	memset(iter+1, 0, b.SizeOf("HashTableIterator")*7)
	idx = iter - EG__().GetHtIterators()
	EG__().SetHtIteratorsUsed(idx + 1)
	return idx
}
func ZendHashIteratorPos(idx uint32, ht *types.HashTable) types.HashPosition {
	var iter *types.HashTableIterator = EG__().GetHtIterators() + idx
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
func ZendHashIteratorPosEx(idx uint32, array *types.Zval) types.HashPosition {
	var ht *types.HashTable = array.GetArr()
	var iter *types.HashTableIterator = EG__().GetHtIterators() + idx
	b.Assert(idx != uint32-1)
	if iter.GetHt() != ht {
		if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(ht.IsIteratorsOverflow()) {
			iter.GetHt().DecNIteratorsCount()
		}
		types.SEPARATE_ARRAY(array)
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
	var iter *types.HashTableIterator = EG__().GetHtIterators() + idx
	b.Assert(idx != uint32-1)
	if iter.GetHt() != nil && iter.GetHt() != HT_POISONED_PTR && !(iter.GetHt().IsIteratorsOverflow()) {
		b.Assert(iter.GetHt().GetNIteratorsCount() != 0)
		iter.GetHt().DecNIteratorsCount()
	}
	iter.SetHt(nil)
	if idx == EG__().GetHtIteratorsUsed()-1 {
		for idx > 0 && EG__().GetHtIterators()[idx-1].GetHt() == nil {
			idx--
		}
		EG__().SetHtIteratorsUsed(idx)
	}
}
func _zendHashIteratorsRemove(ht *types.HashTable) {
	var iter *types.HashTableIterator = EG__().GetHtIterators()
	var end *types.HashTableIterator = iter + EG__().GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht {
			iter.SetHt(HT_POISONED_PTR)
		}
		iter++
	}
}
func ZendHashIteratorsRemove(ht *types.HashTable) {
	if ht.HasIterators() {
		_zendHashIteratorsRemove(ht)
	}
}
func ZendHashIteratorsLowerPos(ht *types.HashTable, start types.HashPosition) types.HashPosition {
	var iter *types.HashTableIterator = EG__().GetHtIterators()
	var end *types.HashTableIterator = iter + EG__().GetHtIteratorsUsed()
	var res types.HashPosition = ht.GetNNumUsed()
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
func _zendHashIteratorsUpdate(ht *types.HashTable, from types.HashPosition, to types.HashPosition) {
	var iter *types.HashTableIterator = EG__().GetHtIterators()
	var end *types.HashTableIterator = iter + EG__().GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht && iter.GetPos() == from {
			iter.SetPos(to)
		}
		iter++
	}
}
func ZendHashIteratorsAdvance(ht *types.HashTable, step types.HashPosition) {
	var iter *types.HashTableIterator = EG__().GetHtIterators()
	var end *types.HashTableIterator = iter + EG__().GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht {
			iter.SetPos(iter.GetPos() + step)
		}
		iter++
	}
}

func ZendHashIndexAddEmptyElement(ht *types.HashTable, h ZendUlong) *types.Zval {
	var dummy types.Zval
	(&dummy).SetUndef()
	return ht.IndexAddH(h, &dummy)
}
func ZendHashAddEmptyElement(ht *types.HashTable, key *types.ZendString) *types.Zval {
	var dummy types.Zval
	(&dummy).SetUndef()
	return ht.KeyAdd(key.GetStr(), &dummy)
}
func ZendHashStrAddEmptyElement(ht *types.HashTable, str *byte, len_ int) *types.Zval {
	var dummy types.Zval
	(&dummy).SetUndef()
	return ht.KeyAdd(b.CastStr(str, len_), &dummy)
}
func ZendHashSetBucketKey(ht *types.HashTable, b *types.Bucket, key *types.ZendString) *types.Zval {
	return ht.SetBucketKey(b, key.GetStr())
}
func ZendHashDelBucket(ht *types.HashTable, p *types.Bucket) {
	ht.assertRc1()
	// todo 调整为传入 pos 更合适
	if pos, ok := ht.posBucket(p); ok {
		ht.deleteBucket(pos)
	}
}
func ZendHashDel(ht *types.HashTable, key *types.ZendString) int {
	var strKey = key.GetStr()
	if ht.KeyDelete(strKey) {
		return types.SUCCESS
	}
	return types.FAILURE
}
func ZendHashDelInd(ht *types.HashTable, key *types.ZendString) int {
	var strKey = key.GetStr()
	if ht.KeyDeleteIndirect(strKey) {
		return types.SUCCESS
	}
	return types.FAILURE
}
func ZendHashStrDel(ht *types.HashTable, str *byte, len_ int) int {
	var strKey = b.CastStr(str, len_)
	if ht.KeyDelete(strKey) {
		return types.SUCCESS
	}
	return types.FAILURE
}
func ZendHashIndexDel(ht *types.HashTable, h ZendUlong) int {
	var index = int(h)
	if ht.IndexDelete(index) {
		return types.SUCCESS
	}
	return types.FAILURE
}

func ZendHashApply(ht *types.HashTable, apply_func ApplyFuncT) {
	ht.applyValidBucket(func(p *types.Bucket) int {
		return apply_func(p.GetVal())
	})
}
func ZendHashApplyWithArgument(ht *types.HashTable, apply_func ApplyFuncArgT, argument any) {
	ht.applyValidBucket(func(p *types.Bucket) int {
		return apply_func(p.GetVal(), argument)
	})
}
func ZendHashApplyWithArguments(ht *types.HashTable, apply_func ApplyFuncArgsT, num_args int, args ...any) {
	ht.applyValidBucket(func(p *types.Bucket) int {
		var hash_key = p.key.GetZendHashKey()
		return apply_func(p.GetVal(), num_args, args, &hash_key)
	})
}
func ZendHashReverseApply(ht *types.HashTable, apply_func ApplyFuncT) {
	ht.applyValidBucketReserve(func(p *types.Bucket) int {
		return apply_func(p.GetVal())
	})
}
func ZendHashCopy(target *types.HashTable, source *types.HashTable, pCopyConstructor types.CopyCtorFuncT) {
	target.assertRc1()
	source.eachValidBucketIndirect(func(pos uint32, p *types.Bucket, data *types.Zval) {
		var newEntry = target.Update(p.key, data)
		if pCopyConstructor != nil {
			pCopyConstructor(newEntry)
		}
	})
}

func ZendArrayDupElements(source *types.HashTable, target *types.HashTable) {
	target.eachValidBucketIndirect(func(pos uint32, p *types.Bucket, data *types.Zval) {
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
		var newBucket = types.NewBucket(p.GetZendKey(), data)
		target.appendBucket(newBucket)

		// 更新内部指针
		if source.nInternalPointer == pos {
			target.nInternalPointer = target.LastPos()
		}
	})
}

func ZendArrayDup(source *types.HashTable) *types.HashTable {
	var target *types.HashTable = types.NewZendArray(source.nTableSize)
	target.AddGcFlags(types.GC_COLLECTABLE)
	target.nNextFreeElement = source.nNextFreeElement

	if source.GetNNumOfElements() == 0 {
		return target
	}

	target.SetFlags(source.GetFlags())

	if (source.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) != 0 {
		target.SetNNumUsed(source.GetNNumUsed())
		target.SetNNumOfElements(source.GetNNumOfElements())
		types.HT_SET_DATA_ADDR(target, Emalloc(types.HT_SIZE(target)))
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
func ZendHashMerge(target *types.HashTable, source *types.HashTable, pCopyConstructor types.CopyCtorFuncT, overwrite types.ZendBool) {
	target.assertRc1()
	if overwrite != 0 {
		source.eachValidBucketIndirect(func(pos uint32, p *types.Bucket, data *types.Zval) {
			var t = target.UpdateIndirect(p.GetZendKey(), data)
			if pCopyConstructor != nil {
				pCopyConstructor(t)
			}
		})
	} else {
		source.eachValidBucketIndirect(func(pos uint32, p *types.Bucket, s *types.Zval) {
			var t = target.AddIndirect(p.GetZendKey(), s)
			if t != nil && pCopyConstructor != nil {
				pCopyConstructor(t)
			}
		})
	}
}
func ZendHashInternalPointerResetEx(ht *types.HashTable, pos *types.HashPosition) {
	*pos = ht.validPosVal(0)
}
func ZendHashInternalPointerEndEx(ht *types.HashTable, pos *types.HashPosition) {
	var idx uint32
	idx = ht.GetNNumUsed()
	for idx > 0 {
		idx--
		if ht.data[idx].GetVal().GetType() != types.IS_UNDEF {
			*pos = idx
			return
		}
	}
	*pos = ht.GetNNumUsed()
}

// 查找下一个有效位置
func ZendHashMoveForwardEx(ht *types.HashTable, pos *types.HashPosition) int {
	if idx, ok := ht.validPos(*pos); ok {
		*pos, _ = ht.validPos(idx + 1)
		return types.SUCCESS
	}
	return types.FAILURE
}

func ZendHashMoveBackwardsEx(ht *types.HashTable, pos *types.HashPosition) int {
	var idx uint32 = *pos
	if idx < ht.GetNNumUsed() {
		for idx > 0 {
			idx--
			if ht.data[idx].GetVal().GetType() != types.IS_UNDEF {
				*pos = idx
				return types.SUCCESS
			}
		}
		*pos = ht.GetNNumUsed()
		return types.SUCCESS
	}
	return types.FAILURE
}
func ZendHashGetCurrentKeyEx(ht *types.HashTable, str_index **types.ZendString, num_index *ZendUlong, pos *types.HashPosition) int {
	var idx uint32
	var p *types.Bucket
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
func ZendHashGetCurrentKeyZvalEx(ht *types.HashTable, key *types.Zval, pos *types.HashPosition) {
	var idx uint32
	var p *types.Bucket
	idx = ht.validPosVal(*pos)
	if idx >= ht.GetNNumUsed() {
		key.SetNull()
	} else {
		p = ht.GetArData() + idx
		if p.GetKey() != nil {
			types.ZVAL_STR_COPY(key, p.GetKey())
		} else {
			key.SetLong(p.IndexKey())
		}
	}
}
func ZendHashGetCurrentKeyTypeEx(ht *types.HashTable, pos *types.HashPosition) int {
	var idx uint32
	var p *types.Bucket
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
func ZendHashGetCurrentDataEx(ht *types.HashTable, pos *types.HashPosition) *types.Zval {
	var idx uint32
	var p *types.Bucket
	idx = ht.validPosVal(*pos)
	if idx < ht.GetNNumUsed() {
		p = ht.GetArData() + idx
		return p.GetVal()
	} else {
		return nil
	}
}
func ZendHashBucketSwap(p *types.Bucket, q *types.Bucket) {
	var val types.Zval
	var h ZendUlong
	var key *types.ZendString
	types.ZVAL_COPY_VALUE(&val, p.GetVal())
	h = p.GetH()
	key = p.GetKey()
	types.ZVAL_COPY_VALUE(p.GetVal(), q.GetVal())
	p.SetH(q.GetH())
	p.SetKey(q.GetKey())
	types.ZVAL_COPY_VALUE(q.GetVal(), &val)
	q.SetH(h)
	q.SetKey(key)
}
func ZendHashCompareImpl(ht1 *types.HashTable, ht2 *types.HashTable, compar types.CompareFuncT, ordered types.ZendBool) int {
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
		var p1 *types.Bucket = ht1.GetArData() + idx1
		var p2 *types.Bucket
		var pData1 *types.Zval
		var pData2 *types.Zval
		var result int
		if p1.GetVal().IsUndef() {
			continue
		}
		if ordered != 0 {
			for true {
				b.Assert(idx2 != ht2.GetNNumUsed())
				p2 = ht2.GetArData() + idx2
				if p2.GetVal().GetType() != types.IS_UNDEF {
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
		if pData1.IsType(types.IS_INDIRECT) {
			pData1 = pData1.GetZv()
		}
		if pData2.IsType(types.IS_INDIRECT) {
			pData2 = pData2.GetZv()
		}
		if pData1.IsType(types.IS_UNDEF) {
			if pData2.GetType() != types.IS_UNDEF {
				return -1
			}
		} else if pData2.IsType(types.IS_UNDEF) {
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
func ZendHashCompare(ht1 *types.HashTable, ht2 *types.HashTable, compar types.CompareFuncT, ordered types.ZendBool) int {
	var result int
	if ht1 == ht2 {
		return 0
	}

	/* It's enough to protect only one of the arrays.
	 * The second one may be referenced from the first and this may cause
	 * false recursion detection.
	 */

	if ht1.IsRecursive() {
		faults.ZendErrorNoreturn(faults.E_ERROR, "Nesting level too deep - recursive dependency?")
	}

	ht1.TryProtectRecursive()

	result = ZendHashCompareImpl(ht1, ht2, compar, ordered)

	ht1.TryUnProtectRecursive()

	return result
}
func ZendHashMinmax(ht *types.HashTable, compar types.CompareFuncT, flag uint32) *types.Zval {
	var res *types.Bucket
	if ht.GetNNumOfElements() == 0 {
		return nil
	}

	ht.eachValidBucket(func(pos uint32, p *types.Bucket) {
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

func zendParseNumericStr(str string) (int, bool) {
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

func ZendSymtableToProptable(ht *types.HashTable) *types.HashTable {
	var num_key ZendUlong
	var str_key *types.ZendString
	var zv *types.Zval

	var __ht *types.HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		str_key = _p.GetKey()
		if str_key == nil {
			goto convert
		}
	}
	if (ht.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
		ht.AddRefcount()
	}
	return ht
convert:
	var new_ht *types.HashTable = ZendNewArray(ht.GetNNumOfElements())
	var __ht__1 *types.HashTable = ht
	for _, _p := range __ht__1.foreachData() {
		var _z *types.Zval = _p.GetVal()

		num_key = _p.GetH()
		str_key = _p.GetKey()
		zv = _z
		if str_key == nil {
			str_key = ZendLongToStr(num_key)
			str_key.DelRefcount()
		}
		for {
			if types.Z_OPT_REFCOUNTED_P(zv) {
				if types.Z_ISREF_P(zv) && types.Z_REFCOUNT_P(zv) == 1 {
					zv = types.Z_REFVAL_P(zv)
					if !(types.Z_OPT_REFCOUNTED_P(zv)) {
						break
					}
				}
				types.Z_ADDREF_P(zv)
			}
			break
		}
		new_ht.KeyUpdate(str_key.GetStr(), zv)
	}
	return new_ht
}
func ZendProptableToSymtable(ht *types.HashTable, always_duplicate types.ZendBool) *types.HashTable {
	var num_key ZendUlong
	var str_key *types.ZendString
	var zv *types.Zval
	var __ht *types.HashTable = ht
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

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
	if (ht.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
		ht.AddRefcount()
	}
	return ht
convert:
	var new_ht *types.HashTable = ZendNewArray(ht.GetNNumOfElements())
	var __ht__1 *types.HashTable = ht
	for _, _p := range __ht__1.foreachData() {
		var _z *types.Zval = _p.GetVal()
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
			if types.Z_OPT_REFCOUNTED_P(zv) {
				if types.Z_ISREF_P(zv) && types.Z_REFCOUNT_P(zv) == 1 {
					zv = types.Z_REFVAL_P(zv)
					if !(types.Z_OPT_REFCOUNTED_P(zv)) {
						break
					}
				}
				types.Z_ADDREF_P(zv)
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
