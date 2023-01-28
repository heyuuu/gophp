// <<generate>>

package zend

func TS_HASH(table *TsHashTable) *HashTable { return &(table.GetHash()) }
func ZendTsHashInit(ht *TsHashTable, nSize uint32, pHashFunction __auto__, pDestructor DtorFuncT, persistent ZendBool) {
	_zendTsHashInit(ht, nSize, pDestructor, persistent)
}
func ZendTsHashInitEx(ht *TsHashTable, nSize uint32, pHashFunction __auto__, pDestructor DtorFuncT, persistent ZendBool, bApplyProtection __auto__) {
	_zendTsHashInit(ht, nSize, pDestructor, persistent)
}
func ZendTsHashStrFindPtr(ht *TsHashTable, str *byte, len_ int) any {
	var zv *Zval
	zv = ZendTsHashStrFind(ht, str, len_)
	if zv != nil {
		return Z_PTR_P(zv)
	} else {
		return nil
	}
}
func ZendTsHashStrUpdatePtr(ht *TsHashTable, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendTsHashStrUpdate(ht, str, len_, &tmp)
	if zv != nil {
		return Z_PTR_P(zv)
	} else {
		return nil
	}
}
func ZendTsHashStrAddPtr(ht *TsHashTable, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	ZVAL_PTR(&tmp, pData)
	zv = ZendTsHashStrAdd(ht, str, len_, &tmp)
	if zv != nil {
		return Z_PTR_P(zv)
	} else {
		return nil
	}
}
func ZendTsHashExists(ht *TsHashTable, key *ZendString) int { return ZendTsHashFind(ht, key) != nil }
func ZendTsHashIndexExists(ht *TsHashTable, h ZendUlong) int {
	return ZendTsHashIndexFind(ht, h) != nil
}
func ZEND_TS_INIT_SYMTABLE(ht *TsHashTable) { ZEND_TS_INIT_SYMTABLE_EX(ht, 2, 0) }
func ZEND_TS_INIT_SYMTABLE_EX(ht *TsHashTable, n uint32, persistent ZendBool) {
	ZendTsHashInit(ht, n, nil, ZVAL_PTR_DTOR, persistent)
}
func BeginRead(ht *TsHashTable)  {}
func EndRead(ht *TsHashTable)    {}
func BeginWrite(ht *TsHashTable) {}
func EndWrite(ht *TsHashTable)   {}
func _zendTsHashInit(ht *TsHashTable, nSize uint32, pDestructor DtorFuncT, persistent ZendBool) {
	_zendHashInit(TS_HASH(ht), nSize, pDestructor, persistent)
}
func ZendTsHashDestroy(ht *TsHashTable) {
	BeginWrite(ht)
	ZendHashDestroy(TS_HASH(ht))
	EndWrite(ht)
}
func ZendTsHashClean(ht *TsHashTable) {
	ht.SetReader(0)
	BeginWrite(ht)
	ZendHashClean(TS_HASH(ht))
	EndWrite(ht)
}
func ZendTsHashAdd(ht *TsHashTable, key *ZendString, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashAdd(TS_HASH(ht), key, pData)
	EndWrite(ht)
	return retval
}
func ZendTsHashUpdate(ht *TsHashTable, key *ZendString, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashUpdate(TS_HASH(ht), key, pData)
	EndWrite(ht)
	return retval
}
func ZendTsHashNextIndexInsert(ht *TsHashTable, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashNextIndexInsert(TS_HASH(ht), pData)
	EndWrite(ht)
	return retval
}
func ZendTsHashIndexUpdate(ht *TsHashTable, h ZendUlong, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashIndexUpdate(TS_HASH(ht), h, pData)
	EndWrite(ht)
	return retval
}
func ZendTsHashAddEmptyElement(ht *TsHashTable, key *ZendString) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashAddEmptyElement(TS_HASH(ht), key)
	EndWrite(ht)
	return retval
}
func ZendTsHashGracefulDestroy(ht *TsHashTable) {
	BeginWrite(ht)
	ZendHashGracefulDestroy(TS_HASH(ht))
	EndWrite(ht)
}
func ZendTsHashApply(ht *TsHashTable, apply_func ApplyFuncT) {
	BeginWrite(ht)
	ZendHashApply(TS_HASH(ht), apply_func)
	EndWrite(ht)
}
func ZendTsHashApplyWithArgument(ht *TsHashTable, apply_func ApplyFuncArgT, argument any) {
	BeginWrite(ht)
	ZendHashApplyWithArgument(TS_HASH(ht), apply_func, argument)
	EndWrite(ht)
}
func ZendTsHashApplyWithArguments(ht *TsHashTable, apply_func ApplyFuncArgsT, num_args int, _ ...any) {
	var args va_list
	va_start(args, num_args)
	BeginWrite(ht)
	ZendHashApplyWithArguments(TS_HASH(ht), apply_func, num_args, args)
	EndWrite(ht)
	va_end(args)
}
func ZendTsHashReverseApply(ht *TsHashTable, apply_func ApplyFuncT) {
	BeginWrite(ht)
	ZendHashReverseApply(TS_HASH(ht), apply_func)
	EndWrite(ht)
}
func ZendTsHashDel(ht *TsHashTable, key *ZendString) int {
	var retval int
	BeginWrite(ht)
	retval = ZendHashDel(TS_HASH(ht), key)
	EndWrite(ht)
	return retval
}
func ZendTsHashIndexDel(ht *TsHashTable, h ZendUlong) int {
	var retval int
	BeginWrite(ht)
	retval = ZendHashIndexDel(TS_HASH(ht), h)
	EndWrite(ht)
	return retval
}
func ZendTsHashFind(ht *TsHashTable, key *ZendString) *Zval {
	var retval *Zval
	BeginRead(ht)
	retval = ZendHashFind(TS_HASH(ht), key)
	EndRead(ht)
	return retval
}
func ZendTsHashIndexFind(ht *TsHashTable, h ZendUlong) *Zval {
	var retval *Zval
	BeginRead(ht)
	retval = ZendHashIndexFind(TS_HASH(ht), h)
	EndRead(ht)
	return retval
}
func ZendTsHashCopy(target *TsHashTable, source *TsHashTable, pCopyConstructor CopyCtorFuncT) {
	BeginRead(source)
	BeginWrite(target)
	ZendHashCopy(TS_HASH(target), TS_HASH(source), pCopyConstructor)
	EndWrite(target)
	EndRead(source)
}
func ZendTsHashCopyToHash(target *HashTable, source *TsHashTable, pCopyConstructor CopyCtorFuncT) {
	BeginRead(source)
	ZendHashCopy(target, TS_HASH(source), pCopyConstructor)
	EndRead(source)
}
func ZendTsHashMerge(target *TsHashTable, source *TsHashTable, pCopyConstructor CopyCtorFuncT, overwrite int) {
	BeginRead(source)
	BeginWrite(target)
	ZendHashMerge(TS_HASH(target), TS_HASH(source), pCopyConstructor, overwrite)
	EndWrite(target)
	EndRead(source)
}
func ZendTsHashMergeEx(target *TsHashTable, source *TsHashTable, pCopyConstructor CopyCtorFuncT, pMergeSource MergeCheckerFuncT, pParam any) {
	BeginRead(source)
	BeginWrite(target)
	ZendHashMergeEx(TS_HASH(target), TS_HASH(source), pCopyConstructor, pMergeSource, pParam)
	EndWrite(target)
	EndRead(source)
}
func ZendTsHashSort(ht *TsHashTable, sort_func SortFuncT, compare_func CompareFuncT, renumber int) int {
	var retval int
	BeginWrite(ht)
	retval = ZendHashSortEx(TS_HASH(ht), sort_func, compare_func, renumber)
	EndWrite(ht)
	return retval
}
func ZendTsHashCompare(ht1 *TsHashTable, ht2 *TsHashTable, compar CompareFuncT, ordered ZendBool) int {
	var retval int
	BeginRead(ht1)
	BeginRead(ht2)
	retval = ZendHashCompare(TS_HASH(ht1), TS_HASH(ht2), compar, ordered)
	EndRead(ht2)
	EndRead(ht1)
	return retval
}
func ZendTsHashMinmax(ht *TsHashTable, compar CompareFuncT, flag int) *Zval {
	var retval *Zval
	BeginRead(ht)
	retval = ZendHashMinmax(TS_HASH(ht), compar, flag)
	EndRead(ht)
	return retval
}
func ZendTsHashNumElements(ht *TsHashTable) int {
	var retval int
	BeginRead(ht)
	retval = TS_HASH(ht).GetNNumOfElements()
	EndRead(ht)
	return retval
}
func ZendTsHashRehash(ht *TsHashTable) int {
	var retval int
	BeginWrite(ht)
	retval = ZendHashRehash(TS_HASH(ht))
	EndWrite(ht)
	return retval
}
func ZendTsHashStrFind(ht *TsHashTable, key *byte, len_ int) *Zval {
	var retval *Zval
	BeginRead(ht)
	retval = ZendHashStrFind(TS_HASH(ht), key, len_)
	EndRead(ht)
	return retval
}
func ZendTsHashStrUpdate(ht *TsHashTable, key *byte, len_ int, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashStrUpdate(TS_HASH(ht), key, len_, pData)
	EndWrite(ht)
	return retval
}
func ZendTsHashStrAdd(ht *TsHashTable, key *byte, len_ int, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashStrAdd(TS_HASH(ht), key, len_, pData)
	EndWrite(ht)
	return retval
}
