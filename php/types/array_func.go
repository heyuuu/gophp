package types

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"strconv"
)

func ArrayLazyDup(arr *Array) *Array {
	// todo 涉及数组的写时复制机制，待完成
	/*
	 *	此处功能要求为:
	 *  - 若原数组数据非只读，标记原数组数组为只读，设置原数组指向数组的一个只读 reader
	 *  - 此方法返回底层数组数据的只读 reader
	 *  只读 reader 要求
	 *  - 读操作时直接读取数据
	 *  - 写操作时，复制底层数组数据后指向新数据，在新数据上操作
	 *  - pos 对应的数据不会发生改变
	 */
	return arr
}

func ArrayRealDup(arr *Array) *Array {
	// todo 涉及数组的写时复制机制，待完成
	/**
	 * 此处功能要求为:
	 * - 确认当前是一个只读 reader
	 * - 若指向的数组上只有唯一一个 reader, 直接使用此数组，并标记为非只读
	 * - 否则，复制此数组作为真实数组使用
	 */
	return ZendArrayDup(arr)
}

func NewArrayOfInt(items []int) *Array {
	// todo
	arr := NewArray(0)
	for _, item := range items {
		arr.Append(NewZvalLong(item))
	}
	return arr
}
func NewArrayOfString(items []string) *Array {
	// todo
	arr := NewArray(0)
	for _, item := range items {
		arr.Append(NewZvalString(item))
	}
	return arr
}
func NewArrayOfZval(items []*Zval) *Array {
	// todo
	arr := NewArray(0)
	for _, item := range items {
		arr.Append(item)
	}
	return arr
}

func ZendHashHasMoreElementsEx(ht *Array, pos *ArrayPosition) bool {
	_, ok := ht.validPos(*pos)
	return ok
}
func ZendHashMoveForward(ht *Array) int {
	return ZendHashMoveForwardEx(ht, &ht.internalPointer)
}

// 查找下一个有效位置
func ZendHashMoveForwardEx(ht *Array, pos *ArrayPosition) int {
	if idx, ok := ht.validPos(*pos); ok {
		*pos, _ = ht.validPos(idx + 1)
		return SUCCESS
	}
	return FAILURE
}

func ZendHashMoveBackwards(ht *Array) int {
	return ZendHashMoveBackwardsEx(ht, &ht.internalPointer)
}
func ZendHashMoveBackwardsEx(ht *Array, pos *ArrayPosition) int {
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

func (ht *Array) Current(indirect bool) (ArrayKey, *Zval, bool) {
	return ht.CurrentEx(ht.internalPointer, indirect)
}

func (ht *Array) CurrentEx(pos uint32, indirect bool) (ArrayKey, *Zval, bool) {
	realPos, ok := ht.validPosEx(pos, indirect)
	if ok {
		p := ht.Bucket(realPos)
		return p.GetArrayKey(), p.GetVal(), true
	} else {
		return ArrayKey{}, nil, false
	}
}

func ZendHashGetCurrentKeyExEx(ht *Array, pos ArrayPosition) *ArrayKey {
	key, _, ok := ht.CurrentEx(pos, false)
	if ok {
		return &key
	}
	return nil
}

func ZendHashGetCurrentKeyZval(ht *Array, key *Zval) {
	ZendHashGetCurrentKeyZvalEx(ht, key, ht.GetNInternalPointer())
}
func ZendHashGetCurrentKeyZvalEx(ht *Array, key *Zval, pos *ArrayPosition) {
	var idx uint32
	var p *Bucket
	idx = ht.validPosVal(*pos)
	if idx >= ht.GetNNumUsed() {
		key.SetNull()
	} else {
		p = ht.Bucket(idx)
		if p.GetKey() != nil {
			key.SetStringVal(p.StrKey())
		} else {
			key.SetLong(p.IndexKey())
		}
	}
}

func ZendHashGetCurrentData(ht *Array) *Zval {
	return ZendHashGetCurrentDataEx(ht, ht.GetNInternalPointer())
}
func ZendHashGetCurrentDataEx(ht *Array, pos *ArrayPosition) *Zval {
	_, val, ok := ht.CurrentEx(*pos, false)
	if !ok {
		return nil
	}
	return val
}

func ZendHashInternalPointerReset(ht *Array) {
	ZendHashInternalPointerResetEx(ht, ht.GetNInternalPointer())
}
func ZendHashInternalPointerResetEx(ht *Array, pos *ArrayPosition) {
	*pos = ht.validPosVal(0)
}

func ZendHashInternalPointerEnd(ht *Array) {
	ZendHashInternalPointerEndEx(ht, ht.GetNInternalPointer())
}
func ZendHashInternalPointerEndEx(ht *Array, pos *ArrayPosition) {
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

func HandleNumericStr(key string, idx *zend.ZendUlong) bool {
	if number, ok := ParseNumericStr(key); ok {
		*idx = zend.ZendUlong(number)
		return true
	} else {
		return false
	}
}
func ZendHashFindInd(ht *Array, key string) *Zval {
	var zv *Zval
	zv = ht.KeyFind(key)
	if zv != nil && zv.IsType(IS_INDIRECT) {
		if Z_INDIRECT_P(zv).GetType() != IS_UNDEF {
			return zv.Indirect()
		} else {
			return nil
		}
	} else {
		return zv
	}
}

func ZendHashUpdatePtr(ht *Array, key string, pData any) any {
	zv := ht.KeyUpdate(key, NewZvalPtr(pData))
	return zv.Ptr()
}

func ZendHashAddMem(ht *Array, key string, pData any, size int) any {
	zv := ht.KeyAdd(key, NewZvalPtr(nil))
	if zv != nil {
		zv.SetPtr(zend.Pemalloc(size))
		memcpy(zv.Ptr(), pData, size)
		return zv.Ptr()
	}
	return nil
}
func ZendHashUpdateMem(ht *Array, key string, pData any, size int) any {
	var p any
	p = zend.Pemalloc(size)
	memcpy(p, pData, size)
	return ZendHashUpdatePtr(ht, key, p)
}
func ZendHashStrUpdateMem(ht *Array, str string, pData any, size int) any {
	var p any
	p = zend.Pemalloc(size)
	memcpy(p, pData, size)
	return ZendHashUpdatePtr(ht, str, p)
}
func ZendHashIndexAddPtr(ht *Array, index int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.IndexAdd(index, &tmp)
	if zv != nil {
		return zv.Ptr()
	} else {
		return nil
	}
}
func ZendHashIndexAddNewPtr(ht *Array, index int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.IndexAddNew(index, &tmp)
	if zv != nil {
		return zv.Ptr()
	} else {
		return nil
	}
}
func ZendHashIndexUpdatePtr(ht *Array, index int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ht.IndexUpdate(index, &tmp)
	return zv.Ptr()
}
func ZendHashNextIndexInsertPtr(ht *Array, pData any) any {
	tmp := NewZvalPtr(pData)
	zv := ht.Append(tmp)
	if zv != nil {
		return zv.Ptr()
	} else {
		return nil
	}
}
func ZendHashIndexUpdateMem(ht *Array, index int, pData any, size int) any {
	var p any
	p = zend.Pemalloc(size)
	memcpy(p, pData, size)
	return ZendHashIndexUpdatePtr(ht, index, p)
}
func ZendHashNextIndexInsertMem(ht *Array, pData any, size int) any {
	tmp := NewZvalPtr(nil)
	zv := ht.Append(tmp)
	if zv != nil {
		zv.SetPtr(zend.Pemalloc(size))
		memcpy(zv.Ptr(), pData, size)
		return zv.Ptr()
	}
	return nil
}
func ZendHashFindPtr(ht *Array, key string) any {
	return ht.KeyFindPtr(key)
}
func ZendHashStrFindPtr(ht *Array, key string) any {
	return ht.KeyFindPtr(key)
}
func ZendHashIndexFindPtr(ht *Array, h int) any {
	var zv *Zval
	zv = ht.IndexFind(h)
	if zv != nil {
		return zv.Ptr()
	} else {
		return nil
	}
}
func ZendHashIndexFindDeref(ht *Array, h int) *Zval {
	var zv = ht.IndexFind(h)
	if zv != nil {
		zv = ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashFindDeref(ht *Array, key string) *Zval {
	var zv *Zval = ht.KeyFind(key)
	if zv != nil {
		zv = ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashStrFindDeref(ht *Array, key string) *Zval {
	var zv *Zval = ht.KeyFind(key)
	if zv != nil {
		zv = ZVAL_DEREF(zv)
	}
	return zv
}
func ZendHashGetCurrentDataPtrEx(ht *Array, pos *ArrayPosition) any {
	var zv *Zval
	zv = ZendHashGetCurrentDataEx(ht, pos)
	if zv != nil {
		return zv.Ptr()
	} else {
		return nil
	}
}
func ZendHashGetCurrentDataPtr(ht *Array) any {
	return ZendHashGetCurrentDataPtrEx(ht, ht.GetNInternalPointer())
}
func ZendHashIteratorPos(idx uint32, ht *Array) ArrayPosition {
	var iter *ArrayIterator = zend.EG__().GetArrayIterator(idx)
	if iter.GetHt() != ht {
		zend.EG__().SetArrayIterator(idx, ht.IteratorEx(ht.currentPosVal()))
	}
	return iter.GetPos()
}
func ZendHashIteratorsLowerPos(ht *Array, start ArrayPosition) ArrayPosition {
	var iter *ArrayIterator = zend.EG__().GetHtIterators()
	var end *ArrayIterator = iter + zend.EG__().GetHtIteratorsUsed()
	var res ArrayPosition = ht.GetNNumUsed()
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

func ZendHashAddEmptyElement(ht *Array, key string) *Zval {
	var dummy Zval
	(&dummy).SetUndef()
	return ht.KeyAdd(key, &dummy)
}
func ZendHashStrAddEmptyElement(ht *Array, str string) *Zval {
	var dummy Zval
	(&dummy).SetUndef()
	return ht.KeyAdd(str, &dummy)
}
func ZendHashDelBucket(ht *Array, p *Bucket) {
	ht.assertRc1()
	// todo 调整为传入 pos 更合适
	if pos, ok := ht.posBucket(p); ok {
		ht.deleteBucket(pos)
	}
}
func ZendHashDel(ht *Array, key string) int {
	if ht.KeyDelete(key) {
		return SUCCESS
	}
	return FAILURE
}
func ZendHashDelInd(ht *Array, key string) int {
	if ht.KeyDeleteIndirect(key) {
		return SUCCESS
	}
	return FAILURE
}
func ZendHashStrDel(ht *Array, key string) int {
	if ht.KeyDelete(key) {
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

func ZendHashCopy(target *Array, source *Array) {
	target.assertRc1()
	source.ForeachIndirect(func(key ArrayKey, value *Zval) {
		target.Update(key, value)
	})
}

func ZendArrayDup(source *Array) *Array {
	// 空数组单独处理
	if source.elementsCount == 0 {
		target := NewArray(0)
		target.nextFreeElement = source.nextFreeElement
		return target
	}

	var target *Array = NewArray(source.Cap())
	target.flags = source.flags
	target.dupData0(source)

	return target
}
func ZendHashMerge(target *Array, source *Array, overwrite bool) {
	target.assertRc1()
	if overwrite {
		source.EachValidBucketIndirect(func(pos uint32, p *Bucket, data *Zval) {
			target.UpdateIndirect(p.GetArrayKey(), data)
		})
	} else {
		source.EachValidBucketIndirect(func(pos uint32, p *Bucket, s *Zval) {
			target.AddIndirect(p.GetArrayKey(), s)
		})
	}
}

func ZendHashCompareImpl(ht1 *Array, ht2 *Array, compar CompareFuncT, ordered ZendBool) int {
	var idx1 uint32
	var idx2 uint32
	if ht1.Len() != ht2.Len() {
		if ht1.Len() > ht2.Len() {
			return 1
		} else {
			return -1
		}
	}
	idx1 = 0
	idx2 = 0
	for ; idx1 < ht1.GetNNumUsed(); idx1++ {
		var p1 *Bucket = ht1.Bucket(idx1)
		var p2 *Bucket
		var pData1 *Zval
		var pData2 *Zval
		var result int
		if p1.GetVal().IsUndef() {
			continue
		}
		if ordered != 0 {
			for true {
				assert(idx2 != ht2.GetNNumUsed())
				p2 = ht2.Bucket(idx2)
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
				pData2 = ht2.KeyFind(p1.GetKey().GetStr())
				if pData2 == nil {
					return 1
				}
			}
		}
		pData1 = p1.GetVal()
		if pData1.IsType(IS_INDIRECT) {
			pData1 = pData1.Indirect()
		}
		if pData2.IsType(IS_INDIRECT) {
			pData2 = pData2.Indirect()
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

func ParseNumericStr(str string) (int, bool) {
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
	for _, _p := range __ht.ForeachData() {
		var _z *Zval = _p.GetVal()

		str_key = _p.GetKey()
		if str_key == nil {
			goto convert
		}
	}
	if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
		// 		ht.AddRefcount()
	}
	return ht
convert:
	var newHt *Array = NewArray(ht.Len())
	ht.Foreach(func(key ArrayKey, zv *Zval) {
		var strKey string
		if key.IsStrKey() {
			strKey = key.StrKey()
		} else {
			strKey = strconv.Itoa(key.IdxKey())
		}
		for {
			if zv.IsRefcounted() {
				if zv.IsReference() && Z_REFCOUNT_P(zv) == 1 {
					zv = Z_REFVAL_P(zv)
					if !(zv.IsRefcounted()) {
						break
					}
				}
				//Z_ADDREF_P(zv)
			}
			break
		}
		newHt.KeyUpdate(strKey, zv)
	})
	return newHt
}
func ZendProptableToSymtable(ht *Array, always_duplicate ZendBool) *Array {
	var num_key zend.ZendUlong
	var str_key *String
	var zv *Zval
	var __ht *Array = ht
	for _, _p := range __ht.ForeachData() {
		var _z *Zval = _p.GetVal()

		str_key = _p.GetKey()

		/* The `str_key &&` here might seem redundant: property tables should
		 * only have string keys. Unfortunately, this isn't true, at the very
		 * least because of ArrayObject, which stores a symtable where the
		 * property table should be.
		 */

		if str_key != nil && HandleNumericStr(str_key, &num_key) {
			goto convert
		}
	}
	if always_duplicate != 0 {
		return ZendArrayDup(ht)
	}
	if (ht.GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
		// 		ht.AddRefcount()
	}
	return ht
convert:
	var new_ht *Array = NewArray(ht.Len())
	var __ht__1 *Array = ht
	for _, _p := range __ht__1.ForeachData() {
		var _z *Zval = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		str_key = _p.GetKey()
		zv = _z
		for {
			if zv.IsRefcounted() {
				if zv.IsReference() && Z_REFCOUNT_P(zv) == 1 {
					zv = Z_REFVAL_P(zv)
					if !(zv.IsRefcounted()) {
						break
					}
				}
				//Z_ADDREF_P(zv)
			}
			break
		}

		/* Again, thank ArrayObject for `!str_key ||`. */

		if str_key == nil || HandleNumericStr(str_key, &num_key) {
			new_ht.IndexUpdate(num_key, zv)
		} else {
			new_ht.KeyUpdate(str_key.GetStr(), zv)
		}

		/* Again, thank ArrayObject for `!str_key ||`. */

	}
	return new_ht
}
