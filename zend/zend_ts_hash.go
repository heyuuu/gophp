// <<generate>>

package zend

// Source: <Zend/zend_ts_hash.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Harald Radi <harald.radi@nme.at>                            |
   +----------------------------------------------------------------------+
*/

// #define ZEND_TS_HASH_H

// # include "zend.h"

// @type TsHashTable struct
type _zendTsHashtable = TsHashTable

// #define TS_HASH(table) ( & ( table -> hash ) )

/* startup/shutdown */

// #define zend_ts_hash_init(ht,nSize,pHashFunction,pDestructor,persistent) _zend_ts_hash_init ( ht , nSize , pDestructor , persistent )

// #define zend_ts_hash_init_ex(ht,nSize,pHashFunction,pDestructor,persistent,bApplyProtection) _zend_ts_hash_init ( ht , nSize , pDestructor , persistent )

/* additions/updates/changes */

/* Deletes */

/* Data retrieval */

/* Copying, merging and sorting */

func ZendTsHashStrFindPtr(ht *TsHashTable, str *byte, len_ int) any {
	var zv *Zval
	zv = ZendTsHashStrFind(ht, str, len_)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendTsHashStrUpdatePtr(ht *TsHashTable, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendTsHashStrUpdate(ht, str, len_, &tmp)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendTsHashStrAddPtr(ht *TsHashTable, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendTsHashStrAdd(ht, str, len_, &tmp)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendTsHashExists(ht *TsHashTable, key *ZendString) int { return ZendTsHashFind(ht, key) != nil }
func ZendTsHashIndexExists(ht *TsHashTable, h ZendUlong) int {
	return ZendTsHashIndexFind(ht, h) != nil
}

// #define ZEND_TS_INIT_SYMTABLE(ht) ZEND_TS_INIT_SYMTABLE_EX ( ht , 2 , 0 )

// #define ZEND_TS_INIT_SYMTABLE_EX(ht,n,persistent) zend_ts_hash_init ( ht , n , NULL , ZVAL_PTR_DTOR , persistent )

// Source: <Zend/zend_ts_hash.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Harald Radi <harald.radi@nme.at>                            |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_ts_hash.h"

/* ts management functions */

func BeginRead(ht *TsHashTable)  {}
func EndRead(ht *TsHashTable)    {}
func BeginWrite(ht *TsHashTable) {}
func EndWrite(ht *TsHashTable)   {}

/* delegates */

func _zendTsHashInit(ht *TsHashTable, nSize uint32, pDestructor DtorFuncT, persistent ZendBool) {
	_zendHashInit(&(ht.GetHash()), nSize, pDestructor, persistent)
}
func ZendTsHashDestroy(ht *TsHashTable) {
	BeginWrite(ht)
	ZendHashDestroy(&(ht.GetHash()))
	EndWrite(ht)
}
func ZendTsHashClean(ht *TsHashTable) {
	ht.SetReader(0)
	BeginWrite(ht)
	ZendHashClean(&(ht.GetHash()))
	EndWrite(ht)
}
func ZendTsHashAdd(ht *TsHashTable, key *ZendString, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashAdd(&(ht.GetHash()), key, pData)
	EndWrite(ht)
	return retval
}
func ZendTsHashUpdate(ht *TsHashTable, key *ZendString, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashUpdate(&(ht.GetHash()), key, pData)
	EndWrite(ht)
	return retval
}
func ZendTsHashNextIndexInsert(ht *TsHashTable, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashNextIndexInsert(&(ht.GetHash()), pData)
	EndWrite(ht)
	return retval
}
func ZendTsHashIndexUpdate(ht *TsHashTable, h ZendUlong, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashIndexUpdate(&(ht.GetHash()), h, pData)
	EndWrite(ht)
	return retval
}
func ZendTsHashAddEmptyElement(ht *TsHashTable, key *ZendString) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashAddEmptyElement(&(ht.GetHash()), key)
	EndWrite(ht)
	return retval
}
func ZendTsHashGracefulDestroy(ht *TsHashTable) {
	BeginWrite(ht)
	ZendHashGracefulDestroy(&(ht.GetHash()))
	EndWrite(ht)
}
func ZendTsHashApply(ht *TsHashTable, apply_func ApplyFuncT) {
	BeginWrite(ht)
	ZendHashApply(&(ht.GetHash()), apply_func)
	EndWrite(ht)
}
func ZendTsHashApplyWithArgument(ht *TsHashTable, apply_func ApplyFuncArgT, argument any) {
	BeginWrite(ht)
	ZendHashApplyWithArgument(&(ht.GetHash()), apply_func, argument)
	EndWrite(ht)
}
func ZendTsHashApplyWithArguments(ht *TsHashTable, apply_func ApplyFuncArgsT, num_args int, _ ...any) {
	var args va_list
	va_start(args, num_args)
	BeginWrite(ht)
	ZendHashApplyWithArguments(&(ht.GetHash()), apply_func, num_args, args)
	EndWrite(ht)
	va_end(args)
}
func ZendTsHashReverseApply(ht *TsHashTable, apply_func ApplyFuncT) {
	BeginWrite(ht)
	ZendHashReverseApply(&(ht.GetHash()), apply_func)
	EndWrite(ht)
}
func ZendTsHashDel(ht *TsHashTable, key *ZendString) int {
	var retval int
	BeginWrite(ht)
	retval = ZendHashDel(&(ht.GetHash()), key)
	EndWrite(ht)
	return retval
}
func ZendTsHashIndexDel(ht *TsHashTable, h ZendUlong) int {
	var retval int
	BeginWrite(ht)
	retval = ZendHashIndexDel(&(ht.GetHash()), h)
	EndWrite(ht)
	return retval
}
func ZendTsHashFind(ht *TsHashTable, key *ZendString) *Zval {
	var retval *Zval
	BeginRead(ht)
	retval = ZendHashFind(&(ht.GetHash()), key)
	EndRead(ht)
	return retval
}
func ZendTsHashIndexFind(ht *TsHashTable, h ZendUlong) *Zval {
	var retval *Zval
	BeginRead(ht)
	retval = ZendHashIndexFind(&(ht.GetHash()), h)
	EndRead(ht)
	return retval
}
func ZendTsHashCopy(target *TsHashTable, source *TsHashTable, pCopyConstructor CopyCtorFuncT) {
	BeginRead(source)
	BeginWrite(target)
	ZendHashCopy(&(target.GetHash()), &(source.GetHash()), pCopyConstructor)
	EndWrite(target)
	EndRead(source)
}
func ZendTsHashCopyToHash(target *HashTable, source *TsHashTable, pCopyConstructor CopyCtorFuncT) {
	BeginRead(source)
	ZendHashCopy(target, &(source.GetHash()), pCopyConstructor)
	EndRead(source)
}
func ZendTsHashMerge(target *TsHashTable, source *TsHashTable, pCopyConstructor CopyCtorFuncT, overwrite int) {
	BeginRead(source)
	BeginWrite(target)
	ZendHashMerge(&(target.GetHash()), &(source.GetHash()), pCopyConstructor, overwrite)
	EndWrite(target)
	EndRead(source)
}
func ZendTsHashMergeEx(target *TsHashTable, source *TsHashTable, pCopyConstructor CopyCtorFuncT, pMergeSource MergeCheckerFuncT, pParam any) {
	BeginRead(source)
	BeginWrite(target)
	ZendHashMergeEx(&(target.GetHash()), &(source.GetHash()), pCopyConstructor, pMergeSource, pParam)
	EndWrite(target)
	EndRead(source)
}
func ZendTsHashSort(ht *TsHashTable, sort_func SortFuncT, compare_func CompareFuncT, renumber int) int {
	var retval int
	BeginWrite(ht)
	retval = ZendHashSortEx(&(ht.GetHash()), sort_func, compare_func, renumber)
	EndWrite(ht)
	return retval
}
func ZendTsHashCompare(ht1 *TsHashTable, ht2 *TsHashTable, compar CompareFuncT, ordered ZendBool) int {
	var retval int
	BeginRead(ht1)
	BeginRead(ht2)
	retval = ZendHashCompare(&(ht1.GetHash()), &(ht2.GetHash()), compar, ordered)
	EndRead(ht2)
	EndRead(ht1)
	return retval
}
func ZendTsHashMinmax(ht *TsHashTable, compar CompareFuncT, flag int) *Zval {
	var retval *Zval
	BeginRead(ht)
	retval = ZendHashMinmax(&(ht.GetHash()), compar, flag)
	EndRead(ht)
	return retval
}
func ZendTsHashNumElements(ht *TsHashTable) int {
	var retval int
	BeginRead(ht)
	retval = &(ht.GetHash()).GetNNumOfElements()
	EndRead(ht)
	return retval
}
func ZendTsHashRehash(ht *TsHashTable) int {
	var retval int
	BeginWrite(ht)
	retval = ZendHashRehash(&(ht.GetHash()))
	EndWrite(ht)
	return retval
}
func ZendTsHashStrFind(ht *TsHashTable, key *byte, len_ int) *Zval {
	var retval *Zval
	BeginRead(ht)
	retval = ZendHashStrFind(&(ht.GetHash()), key, len_)
	EndRead(ht)
	return retval
}
func ZendTsHashStrUpdate(ht *TsHashTable, key *byte, len_ int, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashStrUpdate(&(ht.GetHash()), key, len_, pData)
	EndWrite(ht)
	return retval
}
func ZendTsHashStrAdd(ht *TsHashTable, key *byte, len_ int, pData *Zval) *Zval {
	var retval *Zval
	BeginWrite(ht)
	retval = ZendHashStrAdd(&(ht.GetHash()), key, len_, pData)
	EndWrite(ht)
	return retval
}
