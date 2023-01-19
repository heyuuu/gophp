// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_hash.h>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define ZEND_HASH_H

// # include "zend.h"

// #define HASH_KEY_IS_STRING       1

// #define HASH_KEY_IS_LONG       2

// #define HASH_KEY_NON_EXISTENT       3

// #define HASH_UPDATE       ( 1 << 0 )

// #define HASH_ADD       ( 1 << 1 )

// #define HASH_UPDATE_INDIRECT       ( 1 << 2 )

// #define HASH_ADD_NEW       ( 1 << 3 )

// #define HASH_ADD_NEXT       ( 1 << 4 )

// #define HASH_FLAG_CONSISTENCY       ( ( 1 << 0 ) | ( 1 << 1 ) )

// #define HASH_FLAG_PACKED       ( 1 << 2 )

// #define HASH_FLAG_UNINITIALIZED       ( 1 << 3 )

// #define HASH_FLAG_STATIC_KEYS       ( 1 << 4 )

// #define HASH_FLAG_HAS_EMPTY_IND       ( 1 << 5 )

// #define HASH_FLAG_ALLOW_COW_VIOLATION       ( 1 << 6 )

/* Only the low byte are real flags */

// #define HASH_FLAG_MASK       0xff

// #define HT_FLAGS(ht) ( ht ) -> u . flags

// #define HT_INVALIDATE(ht) do { HT_FLAGS ( ht ) = HASH_FLAG_UNINITIALIZED ; } while ( 0 )

// #define HT_IS_INITIALIZED(ht) ( ( HT_FLAGS ( ht ) & HASH_FLAG_UNINITIALIZED ) == 0 )

// #define HT_IS_PACKED(ht) ( ( HT_FLAGS ( ht ) & HASH_FLAG_PACKED ) != 0 )

// #define HT_IS_WITHOUT_HOLES(ht) ( ( ht ) -> nNumUsed == ( ht ) -> nNumOfElements )

// #define HT_HAS_STATIC_KEYS_ONLY(ht) ( ( HT_FLAGS ( ht ) & ( HASH_FLAG_PACKED | HASH_FLAG_STATIC_KEYS ) ) != 0 )

// #define HT_ALLOW_COW_VIOLATION(ht)

// #define HT_ITERATORS_COUNT(ht) ( ht ) -> u . v . nIteratorsCount

// #define HT_ITERATORS_OVERFLOW(ht) ( HT_ITERATORS_COUNT ( ht ) == 0xff )

// #define HT_HAS_ITERATORS(ht) ( HT_ITERATORS_COUNT ( ht ) != 0 )

// #define HT_SET_ITERATORS_COUNT(ht,iters) do { HT_ITERATORS_COUNT ( ht ) = ( iters ) ; } while ( 0 )

// #define HT_INC_ITERATORS_COUNT(ht) HT_SET_ITERATORS_COUNT ( ht , HT_ITERATORS_COUNT ( ht ) + 1 )

// #define HT_DEC_ITERATORS_COUNT(ht) HT_SET_ITERATORS_COUNT ( ht , HT_ITERATORS_COUNT ( ht ) - 1 )

var ZendEmptyArray HashTable

// #define ZVAL_EMPTY_ARRAY(z) do { zval * __z = ( z ) ; Z_ARR_P ( __z ) = ( zend_array * ) & zend_empty_array ; Z_TYPE_INFO_P ( __z ) = IS_ARRAY ; } while ( 0 )

// @type ZendHashKey struct

type MergeCheckerFuncT func(target_ht *HashTable, source_data *Zval, hash_key *ZendHashKey, pParam any) ZendBool

/* startup/shutdown */

// #define zend_hash_init(ht,nSize,pHashFunction,pDestructor,persistent) _zend_hash_init ( ( ht ) , ( nSize ) , ( pDestructor ) , ( persistent ) )

// #define zend_hash_init_ex(ht,nSize,pHashFunction,pDestructor,persistent,bApplyProtection) _zend_hash_init ( ( ht ) , ( nSize ) , ( pDestructor ) , ( persistent ) )

/* additions/updates/changes */

// #define ZEND_HASH_APPLY_KEEP       0

// #define ZEND_HASH_APPLY_REMOVE       1 << 0

// #define ZEND_HASH_APPLY_STOP       1 << 1

type ApplyFuncT func(pDest *Zval) int
type ApplyFuncArgT func(pDest *Zval, argument any) int
type ApplyFuncArgsT func(pDest *Zval, num_args int, args va_list, hash_key *ZendHashKey) int

/* This function should be used with special care (in other words,
 * it should usually not be used).  When used with the ZEND_HASH_APPLY_STOP
 * return value, it assumes things about the order of the elements in the hash.
 * Also, it does not provide the same kind of reentrancy protection that
 * the standard apply functions do.
 */

/* Deletes */

/* Data retrieval */

/* The same as zend_hash_find(), but hash value of the key must be already calculated */

func ZendHashFindEx(ht *HashTable, key *ZendString, known_hash ZendBool) *Zval {
	if known_hash != 0 {
		return _zendHashFindKnownHash(ht, key)
	} else {
		return ZendHashFind(ht, key)
	}
}

// #define ZEND_HASH_INDEX_FIND(_ht,_h,_ret,_not_found) do { if ( EXPECTED ( HT_FLAGS ( _ht ) & HASH_FLAG_PACKED ) ) { if ( EXPECTED ( ( zend_ulong ) ( _h ) < ( zend_ulong ) ( _ht ) -> nNumUsed ) ) { _ret = & _ht -> arData [ _h ] . val ; if ( UNEXPECTED ( Z_TYPE_P ( _ret ) == IS_UNDEF ) ) { goto _not_found ; } } else { goto _not_found ; } } else { _ret = _zend_hash_index_find ( _ht , _h ) ; if ( UNEXPECTED ( _ret == NULL ) ) { goto _not_found ; } } } while ( 0 )

/* Misc */

func ZendHashExists(ht *HashTable, key *ZendString) ZendBool { return ZendHashFind(ht, key) != nil }
func ZendHashStrExists(ht *HashTable, str *byte, len_ int) ZendBool {
	return ZendHashStrFind(ht, str, len_) != nil
}
func ZendHashIndexExists(ht *HashTable, h ZendUlong) ZendBool { return ZendHashIndexFind(ht, h) != nil }

/* traversing */

// #define zend_hash_has_more_elements_ex(ht,pos) ( zend_hash_get_current_key_type_ex ( ht , pos ) == HASH_KEY_NON_EXISTENT ? FAILURE : SUCCESS )

// #define zend_hash_has_more_elements(ht) zend_hash_has_more_elements_ex ( ht , & ( ht ) -> nInternalPointer )

// #define zend_hash_move_forward(ht) zend_hash_move_forward_ex ( ht , & ( ht ) -> nInternalPointer )

// #define zend_hash_move_backwards(ht) zend_hash_move_backwards_ex ( ht , & ( ht ) -> nInternalPointer )

// #define zend_hash_get_current_key(ht,str_index,num_index) zend_hash_get_current_key_ex ( ht , str_index , num_index , & ( ht ) -> nInternalPointer )

// #define zend_hash_get_current_key_zval(ht,key) zend_hash_get_current_key_zval_ex ( ht , key , & ( ht ) -> nInternalPointer )

// #define zend_hash_get_current_key_type(ht) zend_hash_get_current_key_type_ex ( ht , & ( ht ) -> nInternalPointer )

// #define zend_hash_get_current_data(ht) zend_hash_get_current_data_ex ( ht , & ( ht ) -> nInternalPointer )

// #define zend_hash_internal_pointer_reset(ht) zend_hash_internal_pointer_reset_ex ( ht , & ( ht ) -> nInternalPointer )

// #define zend_hash_internal_pointer_end(ht) zend_hash_internal_pointer_end_ex ( ht , & ( ht ) -> nInternalPointer )

/* Copying, merging and sorting */

// #define zend_hash_sort(ht,compare_func,renumber) zend_hash_sort_ex ( ht , zend_sort , compare_func , renumber )

// #define zend_hash_num_elements(ht) ( ht ) -> nNumOfElements

// #define zend_hash_next_free_element(ht) ( ht ) -> nNextFreeElement

// #define zend_new_array(size) _zend_new_array ( size )

func ZendHashIteratorsUpdate(ht *HashTable, from HashPosition, to HashPosition) {
	if ht.GetNIteratorsCount() != 0 {
		_zendHashIteratorsUpdate(ht, from, to)
	}
}

// #define ZEND_INIT_SYMTABLE(ht) ZEND_INIT_SYMTABLE_EX ( ht , 8 , 0 )

// #define ZEND_INIT_SYMTABLE_EX(ht,n,persistent) zend_hash_init ( ht , n , NULL , ZVAL_PTR_DTOR , persistent )

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

// #define ZEND_HANDLE_NUMERIC_STR(key,length,idx) _zend_handle_numeric_str ( key , length , & idx )

// #define ZEND_HANDLE_NUMERIC(key,idx) ZEND_HANDLE_NUMERIC_STR ( ZSTR_VAL ( key ) , ZSTR_LEN ( key ) , idx )

func ZendHashFindInd(ht *HashTable, key *ZendString) *Zval {
	var zv *Zval
	zv = ZendHashFind(ht, key)
	if zv != nil && zv.GetType() == 13 {
		if zv.GetValue().GetZv().GetType() != 0 {
			return zv.GetValue().GetZv()
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
	if zv != nil && zv.GetType() == 13 {
		if zv.GetValue().GetZv().GetType() != 0 {
			return zv.GetValue().GetZv()
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
	return zv != nil && (zv.GetType() != 13 || zv.GetValue().GetZv().GetType() != 0)
}
func ZendHashStrFindInd(ht *HashTable, str *byte, len_ int) *Zval {
	var zv *Zval
	zv = ZendHashStrFind(ht, str, len_)
	if zv != nil && zv.GetType() == 13 {
		if zv.GetValue().GetZv().GetType() != 0 {
			return zv.GetValue().GetZv()
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
	return zv != nil && (zv.GetType() != 13 || zv.GetValue().GetZv().GetType() != 0)
}
func ZendSymtableAddNew(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	var idx ZendUlong
	if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &idx) != 0 {
		return ZendHashIndexAddNew(ht, idx, pData)
	} else {
		return ZendHashAddNew(ht, key, pData)
	}
}
func ZendSymtableUpdate(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	var idx ZendUlong
	if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &idx) != 0 {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashUpdate(ht, key, pData)
	}
}
func ZendSymtableUpdateInd(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	var idx ZendUlong
	if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &idx) != 0 {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashUpdateInd(ht, key, pData)
	}
}
func ZendSymtableDel(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &idx) != 0 {
		return ZendHashIndexDel(ht, idx)
	} else {
		return ZendHashDel(ht, key)
	}
}
func ZendSymtableDelInd(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &idx) != 0 {
		return ZendHashIndexDel(ht, idx)
	} else {
		return ZendHashDelInd(ht, key)
	}
}
func ZendSymtableFind(ht *HashTable, key *ZendString) *Zval {
	var idx ZendUlong
	if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &idx) != 0 {
		return ZendHashIndexFind(ht, idx)
	} else {
		return ZendHashFind(ht, key)
	}
}
func ZendSymtableFindInd(ht *HashTable, key *ZendString) *Zval {
	var idx ZendUlong
	if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &idx) != 0 {
		return ZendHashIndexFind(ht, idx)
	} else {
		return ZendHashFindInd(ht, key)
	}
}
func ZendSymtableExists(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &idx) != 0 {
		return ZendHashIndexExists(ht, idx)
	} else {
		return ZendHashExists(ht, key)
	}
}
func ZendSymtableExistsInd(ht *HashTable, key *ZendString) int {
	var idx ZendUlong
	if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &idx) != 0 {
		return ZendHashIndexExists(ht, idx)
	} else {
		return ZendHashExistsInd(ht, key)
	}
}
func ZendSymtableStrUpdate(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var idx ZendUlong
	if _zendHandleNumericStr(str, len_, &idx) != 0 {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashStrUpdate(ht, str, len_, pData)
	}
}
func ZendSymtableStrUpdateInd(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var idx ZendUlong
	if _zendHandleNumericStr(str, len_, &idx) != 0 {
		return ZendHashIndexUpdate(ht, idx, pData)
	} else {
		return ZendHashStrUpdateInd(ht, str, len_, pData)
	}
}
func ZendSymtableStrDel(ht *HashTable, str *byte, len_ int) int {
	var idx ZendUlong
	if _zendHandleNumericStr(str, len_, &idx) != 0 {
		return ZendHashIndexDel(ht, idx)
	} else {
		return ZendHashStrDel(ht, str, len_)
	}
}
func ZendSymtableStrDelInd(ht *HashTable, str *byte, len_ int) int {
	var idx ZendUlong
	if _zendHandleNumericStr(str, len_, &idx) != 0 {
		return ZendHashIndexDel(ht, idx)
	} else {
		return ZendHashStrDelInd(ht, str, len_)
	}
}
func ZendSymtableStrFind(ht *HashTable, str *byte, len_ int) *Zval {
	var idx ZendUlong
	if _zendHandleNumericStr(str, len_, &idx) != 0 {
		return ZendHashIndexFind(ht, idx)
	} else {
		return ZendHashStrFind(ht, str, len_)
	}
}
func ZendSymtableStrExists(ht *HashTable, str *byte, len_ int) int {
	var idx ZendUlong
	if _zendHandleNumericStr(str, len_, &idx) != 0 {
		return ZendHashIndexExists(ht, idx)
	} else {
		return ZendHashStrExists(ht, str, len_)
	}
}
func ZendHashAddPtr(ht *HashTable, key *ZendString, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendHashAdd(ht, key, &tmp)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashAddNewPtr(ht *HashTable, key *ZendString, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendHashAddNew(ht, key, &tmp)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashStrAddPtr(ht *HashTable, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendHashStrAdd(ht, str, len_, &tmp)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashStrAddNewPtr(ht *HashTable, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendHashStrAddNew(ht, str, len_, &tmp)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashUpdatePtr(ht *HashTable, key *ZendString, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendHashUpdate(ht, key, &tmp)
	return zv.GetValue().GetPtr()
}
func ZendHashStrUpdatePtr(ht *HashTable, str *byte, len_ int, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendHashStrUpdate(ht, str, len_, &tmp)
	return zv.GetValue().GetPtr()
}
func ZendHashAddMem(ht *HashTable, key *ZendString, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(nil)
	&tmp.SetTypeInfo(14)
	if g.Assign(&zv, ZendHashAdd(ht, key, &tmp)) {
		if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
			zv.GetValue().SetPtr(__zendMalloc(size))
		} else {
			zv.GetValue().SetPtr(_emalloc(size))
		}
		memcpy(zv.GetValue().GetPtr(), pData, size)
		return zv.GetValue().GetPtr()
	}
	return nil
}
func ZendHashAddNewMem(ht *HashTable, key *ZendString, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(nil)
	&tmp.SetTypeInfo(14)
	if g.Assign(&zv, ZendHashAddNew(ht, key, &tmp)) {
		if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
			zv.GetValue().SetPtr(__zendMalloc(size))
		} else {
			zv.GetValue().SetPtr(_emalloc(size))
		}
		memcpy(zv.GetValue().GetPtr(), pData, size)
		return zv.GetValue().GetPtr()
	}
	return nil
}
func ZendHashStrAddMem(ht *HashTable, str *byte, len_ int, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(nil)
	&tmp.SetTypeInfo(14)
	if g.Assign(&zv, ZendHashStrAdd(ht, str, len_, &tmp)) {
		if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
			zv.GetValue().SetPtr(__zendMalloc(size))
		} else {
			zv.GetValue().SetPtr(_emalloc(size))
		}
		memcpy(zv.GetValue().GetPtr(), pData, size)
		return zv.GetValue().GetPtr()
	}
	return nil
}
func ZendHashStrAddNewMem(ht *HashTable, str *byte, len_ int, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(nil)
	&tmp.SetTypeInfo(14)
	if g.Assign(&zv, ZendHashStrAddNew(ht, str, len_, &tmp)) {
		if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
			zv.GetValue().SetPtr(__zendMalloc(size))
		} else {
			zv.GetValue().SetPtr(_emalloc(size))
		}
		memcpy(zv.GetValue().GetPtr(), pData, size)
		return zv.GetValue().GetPtr()
	}
	return nil
}
func ZendHashUpdateMem(ht *HashTable, key *ZendString, pData any, size int) any {
	var p any
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
		p = __zendMalloc(size)
	} else {
		p = _emalloc(size)
	}
	memcpy(p, pData, size)
	return ZendHashUpdatePtr(ht, key, p)
}
func ZendHashStrUpdateMem(ht *HashTable, str *byte, len_ int, pData any, size int) any {
	var p any
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
		p = __zendMalloc(size)
	} else {
		p = _emalloc(size)
	}
	memcpy(p, pData, size)
	return ZendHashStrUpdatePtr(ht, str, len_, p)
}
func ZendHashIndexAddPtr(ht *HashTable, h ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendHashIndexAdd(ht, h, &tmp)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexAddNewPtr(ht *HashTable, h ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendHashIndexAddNew(ht, h, &tmp)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexUpdatePtr(ht *HashTable, h ZendUlong, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendHashIndexUpdate(ht, h, &tmp)
	return zv.GetValue().GetPtr()
}
func ZendHashIndexAddMem(ht *HashTable, h ZendUlong, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(nil)
	&tmp.SetTypeInfo(14)
	if g.Assign(&zv, ZendHashIndexAdd(ht, h, &tmp)) {
		if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
			zv.GetValue().SetPtr(__zendMalloc(size))
		} else {
			zv.GetValue().SetPtr(_emalloc(size))
		}
		memcpy(zv.GetValue().GetPtr(), pData, size)
		return zv.GetValue().GetPtr()
	}
	return nil
}
func ZendHashNextIndexInsertPtr(ht *HashTable, pData any) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(pData)
	&tmp.SetTypeInfo(14)
	zv = ZendHashNextIndexInsert(ht, &tmp)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexUpdateMem(ht *HashTable, h ZendUlong, pData any, size int) any {
	var p any
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
		p = __zendMalloc(size)
	} else {
		p = _emalloc(size)
	}
	memcpy(p, pData, size)
	return ZendHashIndexUpdatePtr(ht, h, p)
}
func ZendHashNextIndexInsertMem(ht *HashTable, pData any, size int) any {
	var tmp Zval
	var zv *Zval
	&tmp.GetValue().SetPtr(nil)
	&tmp.SetTypeInfo(14)
	if g.Assign(&zv, ZendHashNextIndexInsert(ht, &tmp)) {
		if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
			zv.GetValue().SetPtr(__zendMalloc(size))
		} else {
			zv.GetValue().SetPtr(_emalloc(size))
		}
		memcpy(zv.GetValue().GetPtr(), pData, size)
		return zv.GetValue().GetPtr()
	}
	return nil
}
func ZendHashFindPtr(ht *HashTable, key *ZendString) any {
	var zv *Zval
	zv = ZendHashFind(ht, key)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashFindExPtr(ht *HashTable, key *ZendString, known_hash ZendBool) any {
	var zv *Zval
	zv = ZendHashFindEx(ht, key, known_hash)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashStrFindPtr(ht *HashTable, str string, len_ int) any {
	var zv *Zval
	zv = ZendHashStrFind(ht, str, len_)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexFindPtr(ht *HashTable, h ZendUlong) any {
	var zv *Zval
	zv = ZendHashIndexFind(ht, h)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}
func ZendHashIndexFindDeref(ht *HashTable, h ZendUlong) *Zval {
	var zv *Zval = ZendHashIndexFind(ht, h)
	if zv != nil {
		if zv.GetType() == 10 {
			zv = &(*zv).value.GetRef().GetVal()
		}
	}
	return zv
}
func ZendHashFindDeref(ht *HashTable, str *ZendString) *Zval {
	var zv *Zval = ZendHashFind(ht, str)
	if zv != nil {
		if zv.GetType() == 10 {
			zv = &(*zv).value.GetRef().GetVal()
		}
	}
	return zv
}
func ZendHashStrFindDeref(ht *HashTable, str string, len_ int) *Zval {
	var zv *Zval = ZendHashStrFind(ht, str, len_)
	if zv != nil {
		if zv.GetType() == 10 {
			zv = &(*zv).value.GetRef().GetVal()
		}
	}
	return zv
}
func ZendSymtableStrFindPtr(ht *HashTable, str *byte, len_ int) any {
	var idx ZendUlong
	if _zendHandleNumericStr(str, len_, &idx) != 0 {
		return ZendHashIndexFindPtr(ht, idx)
	} else {
		return ZendHashStrFindPtr(ht, str, len_)
	}
}
func ZendHashGetCurrentDataPtrEx(ht *HashTable, pos *HashPosition) any {
	var zv *Zval
	zv = ZendHashGetCurrentDataEx(ht, pos)
	if zv != nil {
		return zv.GetValue().GetPtr()
	} else {
		return nil
	}
}

// #define zend_hash_get_current_data_ptr(ht) zend_hash_get_current_data_ptr_ex ( ht , & ( ht ) -> nInternalPointer )

// #define ZEND_HASH_FOREACH(_ht,indirect) do { HashTable * __ht = ( _ht ) ; Bucket * _p = __ht -> arData ; Bucket * _end = _p + __ht -> nNumUsed ; for ( ; _p != _end ; _p ++ ) { zval * _z = & _p -> val ; if ( indirect && Z_TYPE_P ( _z ) == IS_INDIRECT ) { _z = Z_INDIRECT_P ( _z ) ; } if ( UNEXPECTED ( Z_TYPE_P ( _z ) == IS_UNDEF ) ) continue ;

// #define ZEND_HASH_REVERSE_FOREACH(_ht,indirect) do { HashTable * __ht = ( _ht ) ; uint32_t _idx = __ht -> nNumUsed ; Bucket * _p = __ht -> arData + _idx ; zval * _z ; for ( _idx = __ht -> nNumUsed ; _idx > 0 ; _idx -- ) { _p -- ; _z = & _p -> val ; if ( indirect && Z_TYPE_P ( _z ) == IS_INDIRECT ) { _z = Z_INDIRECT_P ( _z ) ; } if ( UNEXPECTED ( Z_TYPE_P ( _z ) == IS_UNDEF ) ) continue ;

// #define ZEND_HASH_FOREACH_END() } } while ( 0 )

// #define ZEND_HASH_FOREACH_END_DEL() __ht -> nNumOfElements -- ; do { uint32_t j = HT_IDX_TO_HASH ( _idx - 1 ) ; uint32_t nIndex = _p -> h | __ht -> nTableMask ; uint32_t i = HT_HASH ( __ht , nIndex ) ; if ( UNEXPECTED ( j != i ) ) { Bucket * prev = HT_HASH_TO_BUCKET ( __ht , i ) ; while ( Z_NEXT ( prev -> val ) != j ) { i = Z_NEXT ( prev -> val ) ; prev = HT_HASH_TO_BUCKET ( __ht , i ) ; } Z_NEXT ( prev -> val ) = Z_NEXT ( _p -> val ) ; } else { HT_HASH ( __ht , nIndex ) = Z_NEXT ( _p -> val ) ; } } while ( 0 ) ; } __ht -> nNumUsed = _idx ; } while ( 0 )

// #define ZEND_HASH_FOREACH_BUCKET(ht,_bucket) ZEND_HASH_FOREACH ( ht , 0 ) ; _bucket = _p ;

// #define ZEND_HASH_FOREACH_VAL(ht,_val) ZEND_HASH_FOREACH ( ht , 0 ) ; _val = _z ;

// #define ZEND_HASH_FOREACH_VAL_IND(ht,_val) ZEND_HASH_FOREACH ( ht , 1 ) ; _val = _z ;

// #define ZEND_HASH_FOREACH_PTR(ht,_ptr) ZEND_HASH_FOREACH ( ht , 0 ) ; _ptr = Z_PTR_P ( _z ) ;

// #define ZEND_HASH_FOREACH_NUM_KEY(ht,_h) ZEND_HASH_FOREACH ( ht , 0 ) ; _h = _p -> h ;

// #define ZEND_HASH_FOREACH_STR_KEY(ht,_key) ZEND_HASH_FOREACH ( ht , 0 ) ; _key = _p -> key ;

// #define ZEND_HASH_FOREACH_KEY(ht,_h,_key) ZEND_HASH_FOREACH ( ht , 0 ) ; _h = _p -> h ; _key = _p -> key ;

// #define ZEND_HASH_FOREACH_NUM_KEY_VAL(ht,_h,_val) ZEND_HASH_FOREACH ( ht , 0 ) ; _h = _p -> h ; _val = _z ;

// #define ZEND_HASH_FOREACH_STR_KEY_VAL(ht,_key,_val) ZEND_HASH_FOREACH ( ht , 0 ) ; _key = _p -> key ; _val = _z ;

// #define ZEND_HASH_FOREACH_KEY_VAL(ht,_h,_key,_val) ZEND_HASH_FOREACH ( ht , 0 ) ; _h = _p -> h ; _key = _p -> key ; _val = _z ;

// #define ZEND_HASH_FOREACH_STR_KEY_VAL_IND(ht,_key,_val) ZEND_HASH_FOREACH ( ht , 1 ) ; _key = _p -> key ; _val = _z ;

// #define ZEND_HASH_FOREACH_KEY_VAL_IND(ht,_h,_key,_val) ZEND_HASH_FOREACH ( ht , 1 ) ; _h = _p -> h ; _key = _p -> key ; _val = _z ;

// #define ZEND_HASH_FOREACH_NUM_KEY_PTR(ht,_h,_ptr) ZEND_HASH_FOREACH ( ht , 0 ) ; _h = _p -> h ; _ptr = Z_PTR_P ( _z ) ;

// #define ZEND_HASH_FOREACH_STR_KEY_PTR(ht,_key,_ptr) ZEND_HASH_FOREACH ( ht , 0 ) ; _key = _p -> key ; _ptr = Z_PTR_P ( _z ) ;

// #define ZEND_HASH_FOREACH_KEY_PTR(ht,_h,_key,_ptr) ZEND_HASH_FOREACH ( ht , 0 ) ; _h = _p -> h ; _key = _p -> key ; _ptr = Z_PTR_P ( _z ) ;

// #define ZEND_HASH_REVERSE_FOREACH_BUCKET(ht,_bucket) ZEND_HASH_REVERSE_FOREACH ( ht , 0 ) ; _bucket = _p ;

// #define ZEND_HASH_REVERSE_FOREACH_VAL(ht,_val) ZEND_HASH_REVERSE_FOREACH ( ht , 0 ) ; _val = _z ;

// #define ZEND_HASH_REVERSE_FOREACH_PTR(ht,_ptr) ZEND_HASH_REVERSE_FOREACH ( ht , 0 ) ; _ptr = Z_PTR_P ( _z ) ;

// #define ZEND_HASH_REVERSE_FOREACH_VAL_IND(ht,_val) ZEND_HASH_REVERSE_FOREACH ( ht , 1 ) ; _val = _z ;

// #define ZEND_HASH_REVERSE_FOREACH_STR_KEY_VAL(ht,_key,_val) ZEND_HASH_REVERSE_FOREACH ( ht , 0 ) ; _key = _p -> key ; _val = _z ;

// #define ZEND_HASH_REVERSE_FOREACH_KEY_VAL(ht,_h,_key,_val) ZEND_HASH_REVERSE_FOREACH ( ht , 0 ) ; _h = _p -> h ; _key = _p -> key ; _val = _z ;

// #define ZEND_HASH_REVERSE_FOREACH_KEY_VAL_IND(ht,_h,_key,_val) ZEND_HASH_REVERSE_FOREACH ( ht , 1 ) ; _h = _p -> h ; _key = _p -> key ; _val = _z ;

/* The following macros are useful to insert a sequence of new elements
 * of packed array. They may be used instead of series of
 * zend_hash_next_index_insert_new()
 * (HashTable must have enough free buckets).
 */

// #define ZEND_HASH_FILL_PACKED(ht) do { HashTable * __fill_ht = ( ht ) ; Bucket * __fill_bkt = __fill_ht -> arData + __fill_ht -> nNumUsed ; uint32_t __fill_idx = __fill_ht -> nNumUsed ; ZEND_ASSERT ( HT_FLAGS ( __fill_ht ) & HASH_FLAG_PACKED ) ;

// #define ZEND_HASH_FILL_SET(_val) ZVAL_COPY_VALUE ( & __fill_bkt -> val , _val )

// #define ZEND_HASH_FILL_SET_NULL() ZVAL_NULL ( & __fill_bkt -> val )

// #define ZEND_HASH_FILL_SET_LONG(_val) ZVAL_LONG ( & __fill_bkt -> val , _val )

// #define ZEND_HASH_FILL_SET_DOUBLE(_val) ZVAL_DOUBLE ( & __fill_bkt -> val , _val )

// #define ZEND_HASH_FILL_SET_STR(_val) ZVAL_STR ( & __fill_bkt -> val , _val )

// #define ZEND_HASH_FILL_SET_STR_COPY(_val) ZVAL_STR_COPY ( & __fill_bkt -> val , _val )

// #define ZEND_HASH_FILL_SET_INTERNED_STR(_val) ZVAL_INTERNED_STR ( & __fill_bkt -> val , _val )

// #define ZEND_HASH_FILL_NEXT() do { __fill_bkt -> h = ( __fill_idx ) ; __fill_bkt -> key = NULL ; __fill_bkt ++ ; __fill_idx ++ ; } while ( 0 )

// #define ZEND_HASH_FILL_ADD(_val) do { ZEND_HASH_FILL_SET ( _val ) ; ZEND_HASH_FILL_NEXT ( ) ; } while ( 0 )

// #define ZEND_HASH_FILL_END() __fill_ht -> nNumUsed = __fill_idx ; __fill_ht -> nNumOfElements = __fill_idx ; __fill_ht -> nNextFreeElement = __fill_idx ; __fill_ht -> nInternalPointer = 0 ; } while ( 0 )

func _zendHashAppendEx(ht *HashTable, key *ZendString, zv *Zval, interned int) *Zval {
	var idx uint32 = g.PostInc(&(ht.GetNNumUsed()))
	var nIndex uint32
	var p *Bucket = ht.GetArData() + idx
	var _z1 *Zval = &p.val
	var _z2 *Zval = zv
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if interned == 0 && (ZvalGcFlags(key.GetGc().GetTypeInfo())&1<<6) == 0 {
		ht.SetUFlags(ht.GetUFlags() &^ (1 << 4))
		ZendStringAddref(key)
		ZendStringHashVal(key)
	}
	p.SetKey(key)
	p.SetH(key.GetH())
	nIndex = uint32(p.GetH() | ht.GetNTableMask())
	p.GetVal().SetNext((*uint32)(ht.GetArData())[int32(nIndex)])
	(*uint32)(ht.GetArData())[int32(nIndex)] = idx
	ht.GetNNumOfElements()++
	return &p.val
}
func _zendHashAppend(ht *HashTable, key *ZendString, zv *Zval) *Zval {
	return _zendHashAppendEx(ht, key, zv, 0)
}
func _zendHashAppendPtrEx(ht *HashTable, key *ZendString, ptr any, interned int) *Zval {
	var idx uint32 = g.PostInc(&(ht.GetNNumUsed()))
	var nIndex uint32
	var p *Bucket = ht.GetArData() + idx
	&p.val.value.ptr = ptr
	&p.val.u1.type_info = 14
	if interned == 0 && (ZvalGcFlags(key.GetGc().GetTypeInfo())&1<<6) == 0 {
		ht.SetUFlags(ht.GetUFlags() &^ (1 << 4))
		ZendStringAddref(key)
		ZendStringHashVal(key)
	}
	p.SetKey(key)
	p.SetH(key.GetH())
	nIndex = uint32(p.GetH() | ht.GetNTableMask())
	p.GetVal().SetNext((*uint32)(ht.GetArData())[int32(nIndex)])
	(*uint32)(ht.GetArData())[int32(nIndex)] = idx
	ht.GetNNumOfElements()++
	return &p.val
}
func _zendHashAppendPtr(ht *HashTable, key *ZendString, ptr any) *Zval {
	return _zendHashAppendPtrEx(ht, key, ptr, 0)
}
func _zendHashAppendInd(ht *HashTable, key *ZendString, ptr *Zval) {
	var idx uint32 = g.PostInc(&(ht.GetNNumUsed()))
	var nIndex uint32
	var p *Bucket = ht.GetArData() + idx
	&p.val.value.zv = ptr
	&p.val.u1.type_info = 13
	if (ZvalGcFlags(key.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ht.SetUFlags(ht.GetUFlags() &^ (1 << 4))
		ZendStringAddref(key)
		ZendStringHashVal(key)
	}
	p.SetKey(key)
	p.SetH(key.GetH())
	nIndex = uint32(p.GetH() | ht.GetNTableMask())
	p.GetVal().SetNext((*uint32)(ht.GetArData())[int32(nIndex)])
	(*uint32)(ht.GetArData())[int32(nIndex)] = idx
	ht.GetNNumOfElements()++
}

// Source: <Zend/zend_hash.c>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_globals.h"

// # include "zend_variables.h"

// #define HT_ASSERT(ht,expr)

// #define HT_ASSERT_RC1(ht) HT_ASSERT ( ht , GC_REFCOUNT ( ht ) == 1 )

// #define HT_POISONED_PTR       ( ( HashTable * ) ( intptr_t ) - 1 )

// #define IS_CONSISTENT(a)

// #define SET_INCONSISTENT(n)

// #define ZEND_HASH_IF_FULL_DO_RESIZE(ht) if ( ( ht ) -> nNumUsed >= ( ht ) -> nTableSize ) { zend_hash_do_resize ( ht ) ; }

func ZendHashCheckSize(nSize uint32) uint32 {
	/* Use big enough power of 2 */

	if nSize <= 8 {
		return 8
	} else if nSize >= 0x80000000 {
		ZendErrorNoreturn(1<<0, "Possible integer overflow in memory allocation (%u * %zu + %zu)", nSize, g.SizeOf("Bucket"), g.SizeOf("Bucket"))
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
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
		data = __zendMalloc(size_t(ht.GetNTableSize())*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"))
	} else if ht.GetNTableSize() == 8 {
		data = _emalloc(size_t(8)*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"))
	} else {
		data = _emalloc(size_t(ht.GetNTableSize())*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"))
	}
	ht.SetArData((*Bucket)((*byte)(data) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))

	/* Don't overwrite iterator count. */

	ht.SetUVFlags(1<<2 | 1<<4)
	(*uint32)(ht.GetArData())[int32(-2)] = uint32 - 1
	(*uint32)(ht.GetArData())[int32(-1)] = uint32 - 1
}
func ZendHashRealInitMixedEx(ht *HashTable) {
	var data any
	var nSize uint32 = ht.GetNTableSize()
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
		data = __zendMalloc(size_t(nSize)*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32(-(nSize+nSize))))*g.SizeOf("uint32_t"))
	} else if nSize == 8 {
		data = _emalloc(size_t(8)*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32(-(8+8))))*g.SizeOf("uint32_t"))
		ht.SetNTableMask(uint32(-(8 + 8)))
		ht.SetArData((*Bucket)((*byte)(data) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))

		/* Don't overwrite iterator count. */

		ht.SetUVFlags(1 << 4)
		(*uint32)(data)[int32(0)] = -1
		(*uint32)(data)[int32(1)] = -1
		(*uint32)(data)[int32(2)] = -1
		(*uint32)(data)[int32(3)] = -1
		(*uint32)(data)[int32(4)] = -1
		(*uint32)(data)[int32(5)] = -1
		(*uint32)(data)[int32(6)] = -1
		(*uint32)(data)[int32(7)] = -1
		(*uint32)(data)[int32(8)] = -1
		(*uint32)(data)[int32(9)] = -1
		(*uint32)(data)[int32(10)] = -1
		(*uint32)(data)[int32(11)] = -1
		(*uint32)(data)[int32(12)] = -1
		(*uint32)(data)[int32(13)] = -1
		(*uint32)(data)[int32(14)] = -1
		(*uint32)(data)[int32(15)] = -1
		return
	} else {
		data = _emalloc(size_t(nSize)*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32(-(nSize+nSize))))*g.SizeOf("uint32_t"))
	}
	ht.SetNTableMask(uint32(-(nSize + nSize)))
	ht.SetArData((*Bucket)((*byte)(data) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
	ht.SetUFlags(1 << 4)
	memset(&(*uint32)(ht.GetArData())[int32(ht.GetNTableMask())], uint32-1, (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t"))
}
func ZendHashRealInitEx(ht *HashTable, packed int) {
	assert((ht.GetUFlags() & 1 << 3) != 0)
	if packed != 0 {
		ZendHashRealInitPackedEx(ht)
	} else {
		ZendHashRealInitMixedEx(ht)
	}
}

var UninitializedBucket []uint32 = []uint32{uint32 - 1, uint32 - 1}

//ZEND_API const HashTable zend_empty_array = {

func _zendHashInitInt(ht *HashTable, nSize uint32, pDestructor DtorFuncT, persistent ZendBool) {
	ZendGcSetRefcount(&ht.gc, 1)
	ht.GetGc().SetTypeInfo(7 | g.Cond(persistent != 0, 1<<7<<0, 1<<4<<0))
	ht.SetUFlags(1 << 3)
	ht.SetNTableMask(uint32 - 2)
	ht.SetArData((*Bucket)((*byte)(&UninitializedBucket) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
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
	var ht *HashTable = _emalloc(g.SizeOf("HashTable"))
	_zendHashInitInt(ht, 8, ZvalPtrDtor, 0)
	return ht
}
func _zendNewArray(nSize uint32) *HashTable {
	var ht *HashTable = _emalloc(g.SizeOf("HashTable"))
	_zendHashInitInt(ht, nSize, ZvalPtrDtor, 0)
	return ht
}
func ZendNewPair(val1 *Zval, val2 *Zval) *HashTable {
	var p *Bucket
	var ht *HashTable = _emalloc(g.SizeOf("HashTable"))
	_zendHashInitInt(ht, 8, ZvalPtrDtor, 0)
	ht.SetNNextFreeElement(2)
	ht.SetNNumOfElements(ht.GetNNextFreeElement())
	ht.SetNNumUsed(ht.GetNNumOfElements())
	ZendHashRealInitPackedEx(ht)
	p = ht.GetArData()
	var _z1 *Zval = &p.val
	var _z2 *Zval = val1
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	p.SetH(0)
	p.SetKey(nil)
	p++
	var _z1 *Zval = &p.val
	var _z2 *Zval = val2
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	p.SetH(1)
	p.SetKey(nil)
	return ht
}
func ZendHashPackedGrow(ht *HashTable) {
	if ht.GetNTableSize() >= 0x80000000 {
		ZendErrorNoreturn(1<<0, "Possible integer overflow in memory allocation (%u * %zu + %zu)", ht.GetNTableSize()*2, g.SizeOf("Bucket"), g.SizeOf("Bucket"))
	}
	ht.SetNTableSize(ht.GetNTableSize() + ht.GetNTableSize())
	ht.SetArData((*Bucket)((*byte)(g.CondF((ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7) != 0, func() any {
		return __zendRealloc((*byte)(ht.GetArData()-(size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")), size_t(ht.GetNTableSize())*g.SizeOf("Bucket")+(size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"))
	}, func() any {
		return _erealloc2((*byte)(ht.GetArData()-(size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")), size_t(ht.GetNTableSize())*g.SizeOf("Bucket")+(size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"), (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")+size_t(ht).nNumUsed*g.SizeOf("Bucket"))
	})) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
}
func ZendHashRealInit(ht *HashTable, packed ZendBool) { ZendHashRealInitEx(ht, packed) }
func ZendHashRealInitPacked(ht *HashTable)            { ZendHashRealInitPackedEx(ht) }
func ZendHashRealInitMixed(ht *HashTable)             { ZendHashRealInitMixedEx(ht) }
func ZendHashPackedToHash(ht *HashTable) {
	var new_data any
	var old_data any = (*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t"))
	var old_buckets *Bucket = ht.GetArData()
	var nSize uint32 = ht.GetNTableSize()
	ht.SetUFlags(ht.GetUFlags() &^ (1 << 2))
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
		new_data = __zendMalloc(size_t(nSize)*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32(-(nSize+nSize))))*g.SizeOf("uint32_t"))
	} else {
		new_data = _emalloc(size_t(nSize)*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32(-(nSize+nSize))))*g.SizeOf("uint32_t"))
	}
	ht.SetNTableMask(uint32(-(ht.GetNTableSize() + ht.GetNTableSize())))
	ht.SetArData((*Bucket)((*byte)(new_data) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
	memcpy(ht.GetArData(), old_buckets, g.SizeOf("Bucket")*ht.GetNNumUsed())
	g.CondF((ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7) != 0, func() { return Free(old_data) }, func() { return _efree(old_data) })
	ZendHashRehash(ht)
}
func ZendHashToPacked(ht *HashTable) {
	var new_data any
	var old_data any = (*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t"))
	var old_buckets *Bucket = ht.GetArData()
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
		new_data = __zendMalloc(size_t(ht.GetNTableSize())*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"))
	} else {
		new_data = _emalloc(size_t(ht.GetNTableSize())*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"))
	}
	ht.SetUFlags(ht.GetUFlags() | 1<<2 | 1<<4)
	ht.SetNTableMask(uint32 - 2)
	ht.SetArData((*Bucket)((*byte)(new_data) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
	(*uint32)(ht.GetArData())[int32(-2)] = uint32 - 1
	(*uint32)(ht.GetArData())[int32(-1)] = uint32 - 1
	memcpy(ht.GetArData(), old_buckets, g.SizeOf("Bucket")*ht.GetNNumUsed())
	g.CondF((ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7) != 0, func() { return Free(old_data) }, func() { return _efree(old_data) })
}
func ZendHashExtend(ht *HashTable, nSize uint32, packed ZendBool) {
	if nSize == 0 {
		return
	}
	if (ht.GetUFlags() & 1 << 3) != 0 {
		if nSize > ht.GetNTableSize() {
			ht.SetNTableSize(ZendHashCheckSize(nSize))
		}
		ZendHashRealInit(ht, packed)
	} else {
		if packed != 0 {
			assert((ht.GetUFlags() & 1 << 2) != 0)
			if nSize > ht.GetNTableSize() {
				ht.SetNTableSize(ZendHashCheckSize(nSize))
				ht.SetArData((*Bucket)((*byte)(g.CondF((ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7) != 0, func() any {
					return __zendRealloc((*byte)(ht.GetArData()-(size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")), size_t(ht.GetNTableSize())*g.SizeOf("Bucket")+(size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"))
				}, func() any {
					return _erealloc2((*byte)(ht.GetArData()-(size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")), size_t(ht.GetNTableSize())*g.SizeOf("Bucket")+(size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"), (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")+size_t(ht).nNumUsed*g.SizeOf("Bucket"))
				})) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
			}
		} else {
			assert((ht.GetUFlags() & 1 << 2) == 0)
			if nSize > ht.GetNTableSize() {
				var new_data any
				var old_data any = (*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t"))
				var old_buckets *Bucket = ht.GetArData()
				nSize = ZendHashCheckSize(nSize)
				ht.SetNTableSize(nSize)
				if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
					new_data = __zendMalloc(size_t(nSize)*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32(-(nSize+nSize))))*g.SizeOf("uint32_t"))
				} else {
					new_data = _emalloc(size_t(nSize)*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32(-(nSize+nSize))))*g.SizeOf("uint32_t"))
				}
				ht.SetNTableMask(uint32(-(ht.GetNTableSize() + ht.GetNTableSize())))
				ht.SetArData((*Bucket)((*byte)(new_data) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
				memcpy(ht.GetArData(), old_buckets, g.SizeOf("Bucket")*ht.GetNNumUsed())
				g.CondF((ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7) != 0, func() { return Free(old_data) }, func() { return _efree(old_data) })
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
		if p.GetVal().GetType() == 0 {
			continue
		}
		ht.GetNNumOfElements()--

		/* Collision pointers always directed from higher to lower buckets */

		nIndex = p.GetH() | ht.GetNTableMask()
		(*uint32)(arData)[int32(nIndex)] = p.GetVal().GetNext()
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
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			val = _z
			if val.GetType() == 13 {
				if val.GetValue().GetZv().GetType() == 0 {
					num--
				}
			}
		}
		break
	}
	return num
}

/* }}} */

func ZendArrayCount(ht *HashTable) uint32 {
	var num uint32
	if (ht.GetUFlags() & 1 << 5) != 0 {
		num = ZendArrayRecalcElements(ht)
		if ht.GetNNumOfElements() == num {
			ht.SetUFlags(ht.GetUFlags() &^ (1 << 5))
		}
	} else if ht == &EG.symbol_table {
		num = ZendArrayRecalcElements(ht)
	} else {
		num = ht.GetNNumOfElements()
	}
	return num
}

/* }}} */

func _zendHashGetValidPos(ht *HashTable, pos HashPosition) HashPosition {
	for pos < ht.GetNNumUsed() && ht.GetArData()[pos].GetVal().GetType() == 0 {
		pos++
	}
	return pos
}
func _zendHashGetCurrentPos(ht *HashTable) HashPosition {
	return _zendHashGetValidPos(ht, ht.GetNInternalPointer())
}
func ZendHashGetCurrentPos(ht *HashTable) HashPosition { return _zendHashGetCurrentPos(ht) }
func ZendHashIteratorAdd(ht *HashTable, pos HashPosition) uint32 {
	var iter *HashTableIterator = EG.GetHtIterators()
	var end *HashTableIterator = iter + EG.GetHtIteratorsCount()
	var idx uint32
	if ht.GetNIteratorsCount() != 0xff {
		ht.SetNIteratorsCount(ht.GetNIteratorsCount() + 1)
	}
	for iter != end {
		if iter.GetHt() == nil {
			iter.SetHt(ht)
			iter.SetPos(pos)
			idx = iter - EG.GetHtIterators()
			if idx+1 > EG.GetHtIteratorsUsed() {
				EG.SetHtIteratorsUsed(idx + 1)
			}
			return idx
		}
		iter++
	}
	if EG.GetHtIterators() == EG.GetHtIteratorsSlots() {
		EG.SetHtIterators(_emalloc(g.SizeOf("HashTableIterator") * (EG.GetHtIteratorsCount() + 8)))
		memcpy(EG.GetHtIterators(), EG.GetHtIteratorsSlots(), g.SizeOf("HashTableIterator")*EG.GetHtIteratorsCount())
	} else {
		EG.SetHtIterators(_erealloc(EG.GetHtIterators(), g.SizeOf("HashTableIterator")*(EG.GetHtIteratorsCount()+8)))
	}
	iter = EG.GetHtIterators() + EG.GetHtIteratorsCount()
	EG.SetHtIteratorsCount(EG.GetHtIteratorsCount() + 8)
	iter.SetHt(ht)
	iter.SetPos(pos)
	memset(iter+1, 0, g.SizeOf("HashTableIterator")*7)
	idx = iter - EG.GetHtIterators()
	EG.SetHtIteratorsUsed(idx + 1)
	return idx
}
func ZendHashIteratorPos(idx uint32, ht *HashTable) HashPosition {
	var iter *HashTableIterator = EG.GetHtIterators() + idx
	assert(idx != uint32-1)
	if iter.GetHt() != ht {
		if iter.GetHt() != nil && iter.GetHt() != (*HashTable)(intptr_t-1) && iter.GetHt().GetNIteratorsCount() != 0xff {
			iter.GetHt().SetNIteratorsCount(iter.GetHt().GetNIteratorsCount() - 1)
		}
		if ht.GetNIteratorsCount() != 0xff {
			ht.SetNIteratorsCount(ht.GetNIteratorsCount() + 1)
		}
		iter.SetHt(ht)
		iter.SetPos(_zendHashGetCurrentPos(ht))
	}
	return iter.GetPos()
}
func ZendHashIteratorPosEx(idx uint32, array *Zval) HashPosition {
	var ht *HashTable = array.GetValue().GetArr()
	var iter *HashTableIterator = EG.GetHtIterators() + idx
	assert(idx != uint32-1)
	if iter.GetHt() != ht {
		if iter.GetHt() != nil && iter.GetHt() != (*HashTable)(intptr_t-1) && ht.GetNIteratorsCount() != 0xff {
			iter.GetHt().SetNIteratorsCount(iter.GetHt().GetNIteratorsCount() - 1)
		}
		var _zv *Zval = array
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		ht = array.GetValue().GetArr()
		if ht.GetNIteratorsCount() != 0xff {
			ht.SetNIteratorsCount(ht.GetNIteratorsCount() + 1)
		}
		iter.SetHt(ht)
		iter.SetPos(_zendHashGetCurrentPos(ht))
	}
	return iter.GetPos()
}
func ZendHashIteratorDel(idx uint32) {
	var iter *HashTableIterator = EG.GetHtIterators() + idx
	assert(idx != uint32-1)
	if iter.GetHt() != nil && iter.GetHt() != (*HashTable)(intptr_t-1) && iter.GetHt().GetNIteratorsCount() != 0xff {
		assert(iter.GetHt().GetNIteratorsCount() != 0)
		iter.GetHt().SetNIteratorsCount(iter.GetHt().GetNIteratorsCount() - 1)
	}
	iter.SetHt(nil)
	if idx == EG.GetHtIteratorsUsed()-1 {
		for idx > 0 && EG.GetHtIterators()[idx-1].GetHt() == nil {
			idx--
		}
		EG.SetHtIteratorsUsed(idx)
	}
}
func _zendHashIteratorsRemove(ht *HashTable) {
	var iter *HashTableIterator = EG.GetHtIterators()
	var end *HashTableIterator = iter + EG.GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht {
			iter.SetHt((*HashTable)(intptr_t - 1))
		}
		iter++
	}
}
func ZendHashIteratorsRemove(ht *HashTable) {
	if ht.GetNIteratorsCount() != 0 {
		_zendHashIteratorsRemove(ht)
	}
}
func ZendHashIteratorsLowerPos(ht *HashTable, start HashPosition) HashPosition {
	var iter *HashTableIterator = EG.GetHtIterators()
	var end *HashTableIterator = iter + EG.GetHtIteratorsUsed()
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
	var iter *HashTableIterator = EG.GetHtIterators()
	var end *HashTableIterator = iter + EG.GetHtIteratorsUsed()
	for iter != end {
		if iter.GetHt() == ht && iter.GetPos() == from {
			iter.SetPos(to)
		}
		iter++
	}
}
func ZendHashIteratorsAdvance(ht *HashTable, step HashPosition) {
	var iter *HashTableIterator = EG.GetHtIterators()
	var end *HashTableIterator = iter + EG.GetHtIteratorsUsed()
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
	idx = (*uint32)(arData)[int32(nIndex)]
	if idx == uint32-1 {
		return nil
	}
	p = arData + idx
	if p.GetKey() == key {
		return p
	}
	for true {
		if p.GetH() == key.GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), key) != 0 {
			return p
		}
		idx = p.GetVal().GetNext()
		if idx == uint32-1 {
			return nil
		}
		p = arData + idx
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
	idx = (*uint32)(arData)[int32(nIndex)]
	for idx != uint32-1 {
		assert(idx < ht.GetNTableSize())
		p = arData + idx
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
	idx = (*uint32)(arData)[int32(nIndex)]
	for idx != uint32-1 {
		assert(idx < ht.GetNTableSize())
		p = arData + idx
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
	if (ht.GetUFlags() & (1<<3 | 1<<2)) != 0 {
		if (ht.GetUFlags() & 1 << 3) != 0 {
			ZendHashRealInitMixed(ht)
			if (ZvalGcFlags(key.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
				ZendStringAddref(key)
				ht.SetUFlags(ht.GetUFlags() &^ (1 << 4))
				ZendStringHashVal(key)
			}
			goto add_to_hash
		} else {
			ZendHashPackedToHash(ht)
			if (ZvalGcFlags(key.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
				ZendStringAddref(key)
				ht.SetUFlags(ht.GetUFlags() &^ (1 << 4))
				ZendStringHashVal(key)
			}
		}
	} else if (flag & 1 << 3) == 0 {
		p = ZendHashFindBucket(ht, key, 0)
		if p != nil {
			var data *Zval
			assert((flag & 1 << 3) == 0)
			if (flag & 1 << 1) != 0 {
				if (flag & 1 << 2) == 0 {
					return nil
				}
				assert(&p.val != pData)
				data = &p.val
				if data.GetType() == 13 {
					data = data.GetValue().GetZv()
					if data.GetType() != 0 {
						return nil
					}
				} else {
					return nil
				}
			} else {
				assert(&p.val != pData)
				data = &p.val
				if (flag&1<<2) != 0 && data.GetType() == 13 {
					data = data.GetValue().GetZv()
				}
			}
			if ht.GetPDestructor() != nil {
				ht.GetPDestructor()(data)
			}
			var _z1 *Zval = data
			var _z2 *Zval = pData
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			return data
		}
		if (ZvalGcFlags(key.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
			ZendStringAddref(key)
			ht.SetUFlags(ht.GetUFlags() &^ (1 << 4))
		}
	} else if (ZvalGcFlags(key.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendStringAddref(key)
		ht.SetUFlags(ht.GetUFlags() &^ (1 << 4))
		ZendStringHashVal(key)
	}
	if ht.GetNNumUsed() >= ht.GetNTableSize() {
		ZendHashDoResize(ht)
	}
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
	p.GetVal().SetNext((*uint32)(arData)[int32(nIndex)])
	(*uint32)(arData)[int32(nIndex)] = idx
	var _z1 *Zval = &p.val
	var _z2 *Zval = pData
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	return &p.val
}
func _zendHashStrAddOrUpdateI(ht *HashTable, str *byte, len_ int, h ZendUlong, pData *Zval, flag uint32) *Zval {
	var key *ZendString
	var nIndex uint32
	var idx uint32
	var p *Bucket
	if (ht.GetUFlags() & (1<<3 | 1<<2)) != 0 {
		if (ht.GetUFlags() & 1 << 3) != 0 {
			ZendHashRealInitMixed(ht)
			goto add_to_hash
		} else {
			ZendHashPackedToHash(ht)
		}
	} else if (flag & 1 << 3) == 0 {
		p = ZendHashStrFindBucket(ht, str, len_, h)
		if p != nil {
			var data *Zval
			if (flag & 1 << 1) != 0 {
				if (flag & 1 << 2) == 0 {
					return nil
				}
				assert(&p.val != pData)
				data = &p.val
				if data.GetType() == 13 {
					data = data.GetValue().GetZv()
					if data.GetType() != 0 {
						return nil
					}
				} else {
					return nil
				}
			} else {
				assert(&p.val != pData)
				data = &p.val
				if (flag&1<<2) != 0 && data.GetType() == 13 {
					data = data.GetValue().GetZv()
				}
			}
			if ht.GetPDestructor() != nil {
				ht.GetPDestructor()(data)
			}
			var _z1 *Zval = data
			var _z2 *Zval = pData
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			return data
		}
	}
	if ht.GetNNumUsed() >= ht.GetNTableSize() {
		ZendHashDoResize(ht)
	}
add_to_hash:
	ht.GetNNumUsed()++
	idx = ht.GetNNumUsed() - 1
	ht.GetNNumOfElements()++
	p = ht.GetArData() + idx
	key = ZendStringInit(str, len_, ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7)
	p.SetKey(key)
	key.SetH(h)
	p.SetH(key.GetH())
	ht.SetUFlags(ht.GetUFlags() &^ (1 << 4))
	var _z1 *Zval = &p.val
	var _z2 *Zval = pData
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	nIndex = h | ht.GetNTableMask()
	p.GetVal().SetNext((*uint32)(ht.GetArData())[int32(nIndex)])
	(*uint32)(ht.GetArData())[int32(nIndex)] = idx
	return &p.val
}
func ZendHashAddOrUpdate(ht *HashTable, key *ZendString, pData *Zval, flag uint32) *Zval {
	if flag == 1<<1 {
		return ZendHashAdd(ht, key, pData)
	} else if flag == 1<<3 {
		return ZendHashAddNew(ht, key, pData)
	} else if flag == 1<<0 {
		return ZendHashUpdate(ht, key, pData)
	} else {
		assert(flag == (1<<0 | 1<<2))
		return ZendHashUpdateInd(ht, key, pData)
	}
}
func ZendHashAdd(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return _zendHashAddOrUpdateI(ht, key, pData, 1<<1)
}
func ZendHashUpdate(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return _zendHashAddOrUpdateI(ht, key, pData, 1<<0)
}
func ZendHashUpdateInd(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return _zendHashAddOrUpdateI(ht, key, pData, 1<<0|1<<2)
}
func ZendHashAddNew(ht *HashTable, key *ZendString, pData *Zval) *Zval {
	return _zendHashAddOrUpdateI(ht, key, pData, 1<<3)
}
func ZendHashStrAddOrUpdate(ht *HashTable, str *byte, len_ int, pData *Zval, flag uint32) *Zval {
	if flag == 1<<1 {
		return ZendHashStrAdd(ht, str, len_, pData)
	} else if flag == 1<<3 {
		return ZendHashStrAddNew(ht, str, len_, pData)
	} else if flag == 1<<0 {
		return ZendHashStrUpdate(ht, str, len_, pData)
	} else {
		assert(flag == (1<<0 | 1<<2))
		return ZendHashStrUpdateInd(ht, str, len_, pData)
	}
}
func ZendHashStrUpdate(ht *HashTable, str string, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return _zendHashStrAddOrUpdateI(ht, str, len_, h, pData, 1<<0)
}
func ZendHashStrUpdateInd(ht *HashTable, str string, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return _zendHashStrAddOrUpdateI(ht, str, len_, h, pData, 1<<0|1<<2)
}
func ZendHashStrAdd(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return _zendHashStrAddOrUpdateI(ht, str, len_, h, pData, 1<<1)
}
func ZendHashStrAddNew(ht *HashTable, str *byte, len_ int, pData *Zval) *Zval {
	var h ZendUlong = ZendHashFunc(str, len_)
	return _zendHashStrAddOrUpdateI(ht, str, len_, h, pData, 1<<3)
}
func ZendHashIndexAddEmptyElement(ht *HashTable, h ZendUlong) *Zval {
	var dummy Zval
	&dummy.SetTypeInfo(1)
	return ZendHashIndexAdd(ht, h, &dummy)
}
func ZendHashAddEmptyElement(ht *HashTable, key *ZendString) *Zval {
	var dummy Zval
	&dummy.SetTypeInfo(1)
	return ZendHashAdd(ht, key, &dummy)
}
func ZendHashStrAddEmptyElement(ht *HashTable, str *byte, len_ int) *Zval {
	var dummy Zval
	&dummy.SetTypeInfo(1)
	return ZendHashStrAdd(ht, str, len_, &dummy)
}
func _zendHashIndexAddOrUpdateI(ht *HashTable, h ZendUlong, pData *Zval, flag uint32) *Zval {
	var nIndex uint32
	var idx uint32
	var p *Bucket
	if (ht.GetUFlags() & 1 << 2) != 0 {
		if h < ht.GetNNumUsed() {
			p = ht.GetArData() + h
			if p.GetVal().GetType() != 0 {
			replace:
				if (flag & 1 << 1) != 0 {
					return nil
				}
				if ht.GetPDestructor() != nil {
					ht.GetPDestructor()(&p.val)
				}
				var _z1 *Zval = &p.val
				var _z2 *Zval = pData
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				return &p.val
			} else {
				goto convert_to_hash
			}
		} else if h < ht.GetNTableSize() {
		add_to_packed:
			p = ht.GetArData() + h

			/* incremental initialization of empty Buckets */

			if (flag & (1<<3 | 1<<4)) != (1<<3 | 1<<4) {
				if h > ht.GetNNumUsed() {
					var q *Bucket = ht.GetArData() + ht.GetNNumUsed()
					for q != p {
						&q.val.u1.type_info = 0
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
	} else if (ht.GetUFlags() & 1 << 3) != 0 {
		if h < ht.GetNTableSize() {
			ZendHashRealInitPackedEx(ht)
			goto add_to_packed
		}
		ZendHashRealInitMixed(ht)
	} else {
		if (flag & 1 << 3) == 0 {
			p = ZendHashIndexFindBucket(ht, h)
			if p != nil {
				assert((flag & 1 << 3) == 0)
				goto replace
			}
		}
		if ht.GetNNumUsed() >= ht.GetNTableSize() {
			ZendHashDoResize(ht)
		}
	}
	ht.GetNNumUsed()++
	idx = ht.GetNNumUsed() - 1
	nIndex = h | ht.GetNTableMask()
	p = ht.GetArData() + idx
	p.GetVal().SetNext((*uint32)(ht.GetArData())[int32(nIndex)])
	(*uint32)(ht.GetArData())[int32(nIndex)] = idx
	if ZendLong(h >= ZendLong(ht.GetNNextFreeElement())) != 0 {
		if h < INT64_MAX {
			ht.SetNNextFreeElement(h + 1)
		} else {
			ht.SetNNextFreeElement(INT64_MAX)
		}
	}
add:
	ht.GetNNumOfElements()++
	p.SetH(h)
	p.SetKey(nil)
	var _z1 *Zval = &p.val
	var _z2 *Zval = pData
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	return &p.val
}
func ZendHashIndexAddOrUpdate(ht *HashTable, h ZendUlong, pData *Zval, flag uint32) *Zval {
	if flag == 1<<1 {
		return ZendHashIndexAdd(ht, h, pData)
	} else if flag == (1<<1 | 1<<3) {
		return ZendHashIndexAddNew(ht, h, pData)
	} else if flag == (1<<1 | 1<<4) {
		assert(h == ht.GetNNextFreeElement())
		return ZendHashNextIndexInsert(ht, pData)
	} else if flag == (1<<1 | 1<<3 | 1<<4) {
		assert(h == ht.GetNNextFreeElement())
		return ZendHashNextIndexInsertNew(ht, pData)
	} else {
		assert(flag == 1<<0)
		return ZendHashIndexUpdate(ht, h, pData)
	}
}
func ZendHashIndexAdd(ht *HashTable, h ZendUlong, pData *Zval) *Zval {
	return _zendHashIndexAddOrUpdateI(ht, h, pData, 1<<1)
}
func ZendHashIndexAddNew(ht *HashTable, h ZendUlong, pData *Zval) *Zval {
	return _zendHashIndexAddOrUpdateI(ht, h, pData, 1<<1|1<<3)
}
func ZendHashIndexUpdate(ht *HashTable, h ZendUlong, pData *Zval) *Zval {
	return _zendHashIndexAddOrUpdateI(ht, h, pData, 1<<0)
}
func ZendHashNextIndexInsert(ht *HashTable, pData *Zval) *Zval {
	return _zendHashIndexAddOrUpdateI(ht, ht.GetNNextFreeElement(), pData, 1<<1|1<<4)
}
func ZendHashNextIndexInsertNew(ht *HashTable, pData *Zval) *Zval {
	return _zendHashIndexAddOrUpdateI(ht, ht.GetNNextFreeElement(), pData, 1<<1|1<<3|1<<4)
}
func ZendHashSetBucketKey(ht *HashTable, b *Bucket, key *ZendString) *Zval {
	var nIndex uint32
	var idx uint32
	var i uint32
	var p *Bucket
	var arData *Bucket
	assert((ht.GetUFlags() & 1 << 2) == 0)
	p = ZendHashFindBucket(ht, key, 0)
	if p != nil {
		if p == b {
			return &p.val
		} else {
			return nil
		}
	}
	if (ZvalGcFlags(key.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendStringAddref(key)
		ht.SetUFlags(ht.GetUFlags() &^ (1 << 4))
	}
	arData = ht.GetArData()

	/* del from hash */

	idx = b - arData
	nIndex = b.GetH() | ht.GetNTableMask()
	i = (*uint32)(arData)[int32(nIndex)]
	if i == idx {
		(*uint32)(arData)[int32(nIndex)] = b.GetVal().GetNext()
	} else {
		p = arData + i
		for p.GetVal().GetNext() != idx {
			i = p.GetVal().GetNext()
			p = arData + i
		}
		p.GetVal().SetNext(b.GetVal().GetNext())
	}
	ZendStringRelease(b.GetKey())

	/* add to hash */

	idx = b - arData
	b.SetKey(key)
	b.SetH(key.GetH())
	nIndex = b.GetH() | ht.GetNTableMask()
	idx = idx
	i = (*uint32)(arData)[int32(nIndex)]
	if i == uint32-1 || i < idx {
		b.GetVal().SetNext(i)
		(*uint32)(arData)[int32(nIndex)] = idx
	} else {
		p = arData + i
		for p.GetVal().GetNext() != uint32-1 && p.GetVal().GetNext() > idx {
			i = p.GetVal().GetNext()
			p = arData + i
		}
		b.GetVal().SetNext(p.GetVal().GetNext())
		p.GetVal().SetNext(idx)
	}
	return &b.val
}
func ZendHashDoResize(ht *HashTable) {
	if ht.GetNNumUsed() > ht.GetNNumOfElements()+(ht.GetNNumOfElements()>>5) {
		ZendHashRehash(ht)
	} else if ht.GetNTableSize() < 0x80000000 {
		var new_data any
		var old_data any = (*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t"))
		var nSize uint32 = ht.GetNTableSize() + ht.GetNTableSize()
		var old_buckets *Bucket = ht.GetArData()
		ht.SetNTableSize(nSize)
		if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
			new_data = __zendMalloc(size_t(nSize)*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32(-(nSize+nSize))))*g.SizeOf("uint32_t"))
		} else {
			new_data = _emalloc(size_t(nSize)*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32(-(nSize+nSize))))*g.SizeOf("uint32_t"))
		}
		ht.SetNTableMask(uint32(-(ht.GetNTableSize() + ht.GetNTableSize())))
		ht.SetArData((*Bucket)((*byte)(new_data) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
		memcpy(ht.GetArData(), old_buckets, g.SizeOf("Bucket")*ht.GetNNumUsed())
		g.CondF((ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7) != 0, func() { return Free(old_data) }, func() { return _efree(old_data) })
		ZendHashRehash(ht)
	} else {
		ZendErrorNoreturn(1<<0, "Possible integer overflow in memory allocation (%u * %zu + %zu)", ht.GetNTableSize()*2, g.SizeOf("Bucket")+g.SizeOf("uint32_t"), g.SizeOf("Bucket"))
	}
}
func ZendHashRehash(ht *HashTable) int {
	var p *Bucket
	var nIndex uint32
	var i uint32
	if ht.GetNNumOfElements() == 0 {
		if (ht.GetUFlags() & 1 << 3) == 0 {
			ht.SetNNumUsed(0)
			memset(&(*uint32)(ht.GetArData())[int32(ht.GetNTableMask())], uint32-1, (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t"))
		}
		return SUCCESS
	}
	memset(&(*uint32)(ht.GetArData())[int32(ht.GetNTableMask())], uint32-1, (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t"))
	i = 0
	p = ht.GetArData()
	if ht.GetNNumUsed() == ht.GetNNumOfElements() {
		for {
			nIndex = p.GetH() | ht.GetNTableMask()
			p.GetVal().SetNext((*uint32)(ht.GetArData())[int32(nIndex)])
			(*uint32)(ht.GetArData())[int32(nIndex)] = i
			p++
			if g.PreInc(&i) >= ht.GetNNumUsed() {
				break
			}
		}
	} else {
		var old_num_used uint32 = ht.GetNNumUsed()
		for {
			if p.GetVal().GetType() == 0 {
				var j uint32 = i
				var q *Bucket = p
				if ht.GetNIteratorsCount() == 0 {
					for g.PreInc(&i) < ht.GetNNumUsed() {
						p++
						if p.GetVal().GetTypeInfo() != 0 {
							var _z1 *Zval = &q.val
							var _z2 *Zval = &p.val
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							q.SetH(p.GetH())
							nIndex = q.GetH() | ht.GetNTableMask()
							q.SetKey(p.GetKey())
							q.GetVal().SetNext((*uint32)(ht.GetArData())[int32(nIndex)])
							(*uint32)(ht.GetArData())[int32(nIndex)] = j
							if ht.GetNInternalPointer() == i {
								ht.SetNInternalPointer(j)
							}
							q++
							j++
						}
					}
				} else {
					var iter_pos uint32 = ZendHashIteratorsLowerPos(ht, 0)
					for g.PreInc(&i) < ht.GetNNumUsed() {
						p++
						if p.GetVal().GetTypeInfo() != 0 {
							var _z1 *Zval = &q.val
							var _z2 *Zval = &p.val
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							q.SetH(p.GetH())
							nIndex = q.GetH() | ht.GetNTableMask()
							q.SetKey(p.GetKey())
							q.GetVal().SetNext((*uint32)(ht.GetArData())[int32(nIndex)])
							(*uint32)(ht.GetArData())[int32(nIndex)] = j
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
			p.GetVal().SetNext((*uint32)(ht.GetArData())[int32(nIndex)])
			(*uint32)(ht.GetArData())[int32(nIndex)] = i
			p++
			if g.PreInc(&i) >= ht.GetNNumUsed() {
				break
			}
		}

		/* Migrate pointer to one past the end of the array to the new one past the end, so that
		 * newly inserted elements are picked up correctly. */

		if ht.GetNIteratorsCount() != 0 {
			_zendHashIteratorsUpdate(ht, old_num_used, ht.GetNNumUsed())
		}

		/* Migrate pointer to one past the end of the array to the new one past the end, so that
		 * newly inserted elements are picked up correctly. */

	}
	return SUCCESS
}
func _zendHashDelElEx(ht *HashTable, idx uint32, p *Bucket, prev *Bucket) {
	if (ht.GetUFlags() & 1 << 2) == 0 {
		if prev != nil {
			prev.GetVal().SetNext(p.GetVal().GetNext())
		} else {
			(*uint32)(ht.GetArData())[int32(p.GetH()|ht.GetNTableMask())] = p.GetVal().GetNext()
		}
	}
	idx = idx
	ht.GetNNumOfElements()--
	if ht.GetNInternalPointer() == idx || ht.GetNIteratorsCount() != 0 {
		var new_idx uint32
		new_idx = idx
		for true {
			new_idx++
			if new_idx >= ht.GetNNumUsed() {
				break
			} else if ht.GetArData()[new_idx].GetVal().GetType() != 0 {
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
			if !(ht.GetNNumUsed() > 0 && ht.GetArData()[ht.GetNNumUsed()-1].GetVal().GetType() == 0) {
				break
			}
		}
		if ht.GetNInternalPointer() < ht.GetNNumUsed() {
			ht.SetNInternalPointer(ht.GetNInternalPointer())
		} else {
			ht.SetNInternalPointer(ht.GetNNumUsed())
		}
	}
	if p.GetKey() != nil {
		ZendStringRelease(p.GetKey())
	}
	if ht.GetPDestructor() != nil {
		var tmp Zval
		var _z1 *Zval = &tmp
		var _z2 *Zval = &p.val
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		&p.val.u1.type_info = 0
		ht.GetPDestructor()(&tmp)
	} else {
		&p.val.u1.type_info = 0
	}
}
func _zendHashDelEl(ht *HashTable, idx uint32, p *Bucket) {
	var prev *Bucket = nil
	if (ht.GetUFlags() & 1 << 2) == 0 {
		var nIndex uint32 = p.GetH() | ht.GetNTableMask()
		var i uint32 = (*uint32)(ht.GetArData())[int32(nIndex)]
		if i != idx {
			prev = ht.GetArData() + i
			for prev.GetVal().GetNext() != idx {
				i = prev.GetVal().GetNext()
				prev = ht.GetArData() + i
			}
		}
	}
	_zendHashDelElEx(ht, idx, p, prev)
}
func ZendHashDelBucket(ht *HashTable, p *Bucket) {
	_zendHashDelEl(ht, p-ht.GetArData(), p)
}
func ZendHashDel(ht *HashTable, key *ZendString) int {
	var h ZendUlong
	var nIndex uint32
	var idx uint32
	var p *Bucket
	var prev *Bucket = nil
	h = ZendStringHashVal(key)
	nIndex = h | ht.GetNTableMask()
	idx = (*uint32)(ht.GetArData())[int32(nIndex)]
	for idx != uint32-1 {
		p = ht.GetArData() + idx
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
	h = ZendStringHashVal(key)
	nIndex = h | ht.GetNTableMask()
	idx = (*uint32)(ht.GetArData())[int32(nIndex)]
	for idx != uint32-1 {
		p = ht.GetArData() + idx
		if p.GetKey() == key || p.GetH() == h && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), key) != 0 {
			if p.GetVal().GetType() == 13 {
				var data *Zval = p.GetVal().GetValue().GetZv()
				if data.GetType() == 0 {
					return FAILURE
				} else {
					if ht.GetPDestructor() != nil {
						var tmp Zval
						var _z1 *Zval = &tmp
						var _z2 *Zval = data
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						data.SetTypeInfo(0)
						ht.GetPDestructor()(&tmp)
					} else {
						data.SetTypeInfo(0)
					}
					ht.SetUFlags(ht.GetUFlags() | 1<<5)
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
	h = ZendInlineHashFunc(str, len_)
	nIndex = h | ht.GetNTableMask()
	idx = (*uint32)(ht.GetArData())[int32(nIndex)]
	for idx != uint32-1 {
		p = ht.GetArData() + idx
		if p.GetH() == h && p.GetKey() != nil && p.GetKey().GetLen() == len_ && !(memcmp(p.GetKey().GetVal(), str, len_)) {
			if p.GetVal().GetType() == 13 {
				var data *Zval = p.GetVal().GetValue().GetZv()
				if data.GetType() == 0 {
					return FAILURE
				} else {
					if ht.GetPDestructor() != nil {
						ht.GetPDestructor()(data)
					}
					data.SetTypeInfo(0)
					ht.SetUFlags(ht.GetUFlags() | 1<<5)
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
	h = ZendInlineHashFunc(str, len_)
	nIndex = h | ht.GetNTableMask()
	idx = (*uint32)(ht.GetArData())[int32(nIndex)]
	for idx != uint32-1 {
		p = ht.GetArData() + idx
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
	if (ht.GetUFlags() & 1 << 2) != 0 {
		if h < ht.GetNNumUsed() {
			p = ht.GetArData() + h
			if p.GetVal().GetType() != 0 {
				_zendHashDelElEx(ht, h, p, nil)
				return SUCCESS
			}
		}
		return FAILURE
	}
	nIndex = h | ht.GetNTableMask()
	idx = (*uint32)(ht.GetArData())[int32(nIndex)]
	for idx != uint32-1 {
		p = ht.GetArData() + idx
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
			if (ht.GetUFlags() & (1<<2 | 1<<4)) != 0 {
				if ht.GetNNumUsed() == ht.GetNNumOfElements() {
					for {
						ht.GetPDestructor()(&p.val)
						if g.PreInc(&p) == end {
							break
						}
					}
				} else {
					for {
						if p.GetVal().GetType() != 0 {
							ht.GetPDestructor()(&p.val)
						}
						if g.PreInc(&p) == end {
							break
						}
					}
				}
			} else if ht.GetNNumUsed() == ht.GetNNumOfElements() {
				for {
					ht.GetPDestructor()(&p.val)
					if p.GetKey() != nil {
						ZendStringRelease(p.GetKey())
					}
					if g.PreInc(&p) == end {
						break
					}
				}
			} else {
				for {
					if p.GetVal().GetType() != 0 {
						ht.GetPDestructor()(&p.val)
						if p.GetKey() != nil {
							ZendStringRelease(p.GetKey())
						}
					}
					if g.PreInc(&p) == end {
						break
					}
				}
			}
		} else {
			if (ht.GetUFlags() & (1<<2 | 1<<4)) == 0 {
				for {
					if p.GetVal().GetType() != 0 {
						if p.GetKey() != nil {
							ZendStringRelease(p.GetKey())
						}
					}
					if g.PreInc(&p) == end {
						break
					}
				}
			}
		}
		ZendHashIteratorsRemove(ht)
	} else if (ht.GetUFlags() & 1 << 3) != 0 {
		return
	}
	g.CondF((ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7) != 0, func() {
		return Free((*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
	}, func() {
		return _efree((*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
	})
}
func ZendArrayDestroy(ht *HashTable) {
	var p *Bucket
	var end *Bucket

	/* break possible cycles */

	var _p *ZendRefcounted = (*ZendRefcounted)(ht)
	if (_p.GetGc().GetTypeInfo() & 0xfffffc00) != 0 {
		GcRemoveFromBuffer(_p)
	}
	ht.GetGc().SetTypeInfo(1)
	if ht.GetNNumUsed() != 0 {

		/* In some rare cases destructors of regular arrays may be changed */

		if ht.GetPDestructor() != ZvalPtrDtor {
			ZendHashDestroy(ht)
			goto free_ht
		}
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if (ht.GetUFlags() & (1<<2 | 1<<4)) != 0 {
			for {
				IZvalPtrDtor(&p.val)
				if g.PreInc(&p) == end {
					break
				}
			}
		} else if ht.GetNNumUsed() == ht.GetNNumOfElements() {
			for {
				IZvalPtrDtor(&p.val)
				if p.GetKey() != nil {
					ZendStringReleaseEx(p.GetKey(), 0)
				}
				if g.PreInc(&p) == end {
					break
				}
			}
		} else {
			for {
				if p.GetVal().GetType() != 0 {
					IZvalPtrDtor(&p.val)
					if p.GetKey() != nil {
						ZendStringReleaseEx(p.GetKey(), 0)
					}
				}
				if g.PreInc(&p) == end {
					break
				}
			}
		}
	} else if (ht.GetUFlags() & 1 << 3) != 0 {
		goto free_ht
	}
	_efree((*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
free_ht:
	ZendHashIteratorsRemove(ht)
	_efree(ht)
}
func ZendHashClean(ht *HashTable) {
	var p *Bucket
	var end *Bucket
	if ht.GetNNumUsed() != 0 {
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if ht.GetPDestructor() != nil {
			if (ht.GetUFlags() & (1<<2 | 1<<4)) != 0 {
				if ht.GetNNumUsed() == ht.GetNNumOfElements() {
					for {
						ht.GetPDestructor()(&p.val)
						if g.PreInc(&p) == end {
							break
						}
					}
				} else {
					for {
						if p.GetVal().GetType() != 0 {
							ht.GetPDestructor()(&p.val)
						}
						if g.PreInc(&p) == end {
							break
						}
					}
				}
			} else if ht.GetNNumUsed() == ht.GetNNumOfElements() {
				for {
					ht.GetPDestructor()(&p.val)
					if p.GetKey() != nil {
						ZendStringRelease(p.GetKey())
					}
					if g.PreInc(&p) == end {
						break
					}
				}
			} else {
				for {
					if p.GetVal().GetType() != 0 {
						ht.GetPDestructor()(&p.val)
						if p.GetKey() != nil {
							ZendStringRelease(p.GetKey())
						}
					}
					if g.PreInc(&p) == end {
						break
					}
				}
			}
		} else {
			if (ht.GetUFlags() & (1<<2 | 1<<4)) == 0 {
				if ht.GetNNumUsed() == ht.GetNNumOfElements() {
					for {
						if p.GetKey() != nil {
							ZendStringRelease(p.GetKey())
						}
						if g.PreInc(&p) == end {
							break
						}
					}
				} else {
					for {
						if p.GetVal().GetType() != 0 {
							if p.GetKey() != nil {
								ZendStringRelease(p.GetKey())
							}
						}
						if g.PreInc(&p) == end {
							break
						}
					}
				}
			}
		}
		if (ht.GetUFlags() & 1 << 2) == 0 {
			memset(&(*uint32)(ht.GetArData())[int32(ht.GetNTableMask())], uint32-1, (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t"))
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
	if ht.GetNNumUsed() != 0 {
		p = ht.GetArData()
		end = p + ht.GetNNumUsed()
		if (ht.GetUFlags() & (1<<2 | 1<<4)) != 0 {
			for {
				IZvalPtrDtor(&p.val)
				if g.PreInc(&p) == end {
					break
				}
			}
		} else if ht.GetNNumUsed() == ht.GetNNumOfElements() {
			for {
				IZvalPtrDtor(&p.val)
				if p.GetKey() != nil {
					ZendStringRelease(p.GetKey())
				}
				if g.PreInc(&p) == end {
					break
				}
			}
		} else {
			for {
				if p.GetVal().GetType() != 0 {
					IZvalPtrDtor(&p.val)
					if p.GetKey() != nil {
						ZendStringRelease(p.GetKey())
					}
				}
				if g.PreInc(&p) == end {
					break
				}
			}
		}
		memset(&(*uint32)(ht.GetArData())[int32(ht.GetNTableMask())], uint32-1, (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t"))
	}
	ht.SetNNumUsed(0)
	ht.SetNNumOfElements(0)
	ht.SetNNextFreeElement(0)
	ht.SetNInternalPointer(0)
}
func ZendHashGracefulDestroy(ht *HashTable) {
	var idx uint32
	var p *Bucket
	p = ht.GetArData()
	for idx = 0; idx < ht.GetNNumUsed(); {
		if p.GetVal().GetType() == 0 {
			continue
		}
		_zendHashDelEl(ht, idx, p)
		idx++
		p++
	}
	if (ht.GetUFlags() & 1 << 3) == 0 {
		g.CondF((ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7) != 0, func() {
			return Free((*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
		}, func() {
			return _efree((*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
		})
	}
}
func ZendHashGracefulReverseDestroy(ht *HashTable) {
	var idx uint32
	var p *Bucket
	idx = ht.GetNNumUsed()
	p = ht.GetArData() + ht.GetNNumUsed()
	for idx > 0 {
		idx--
		p--
		if p.GetVal().GetType() == 0 {
			continue
		}
		_zendHashDelEl(ht, idx, p)
	}
	if (ht.GetUFlags() & 1 << 3) == 0 {
		g.CondF((ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7) != 0, func() {
			return Free((*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
		}, func() {
			return _efree((*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
		})
	}
}

/* This is used to recurse elements and selectively delete certain entries
 * from a hashtable. apply_func() receives the data and decides if the entry
 * should be deleted or recursion should be stopped. The following three
 * return codes are possible:
 * ZEND_HASH_APPLY_KEEP   - continue
 * ZEND_HASH_APPLY_STOP   - stop iteration
 * ZEND_HASH_APPLY_REMOVE - delete the element, combineable with the former
 */

func ZendHashApply(ht *HashTable, apply_func ApplyFuncT) {
	var idx uint32
	var p *Bucket
	var result int
	for idx = 0; idx < ht.GetNNumUsed(); idx++ {
		p = ht.GetArData() + idx
		if p.GetVal().GetType() == 0 {
			continue
		}
		result = apply_func(&p.val)
		if (result & 1 << 0) != 0 {
			_zendHashDelEl(ht, idx, p)
		}
		if (result & 1 << 1) != 0 {
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
		if p.GetVal().GetType() == 0 {
			continue
		}
		result = apply_func(&p.val, argument)
		if (result & 1 << 0) != 0 {
			_zendHashDelEl(ht, idx, p)
		}
		if (result & 1 << 1) != 0 {
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
		if p.GetVal().GetType() == 0 {
			continue
		}
		va_start(args, num_args)
		hash_key.SetH(p.GetH())
		hash_key.SetKey(p.GetKey())
		result = apply_func(&p.val, num_args, args, &hash_key)
		if (result & 1 << 0) != 0 {
			_zendHashDelEl(ht, idx, p)
		}
		if (result & 1 << 1) != 0 {
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
		if p.GetVal().GetType() == 0 {
			continue
		}
		result = apply_func(&p.val)
		if (result & 1 << 0) != 0 {
			_zendHashDelEl(ht, idx, p)
		}
		if (result & 1 << 1) != 0 {
			break
		}
	}
}
func ZendHashCopy(target *HashTable, source *HashTable, pCopyConstructor CopyCtorFuncT) {
	var idx uint32
	var p *Bucket
	var new_entry *Zval
	var data *Zval
	for idx = 0; idx < source.GetNNumUsed(); idx++ {
		p = source.GetArData() + idx
		if p.GetVal().GetType() == 0 {
			continue
		}

		/* INDIRECT element may point to UNDEF-ined slots */

		data = &p.val
		if data.GetType() == 13 {
			data = data.GetValue().GetZv()
			if data.GetType() == 0 {
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
	var data *Zval = &p.val
	if with_holes != 0 {
		if packed == 0 && data.GetTypeInfo() == 13 {
			data = data.GetValue().GetZv()
		}
		if data.GetTypeInfo() == 0 {
			return 0
		}
	} else if packed == 0 {

		/* INDIRECT element may point to UNDEF-ined slots */

		if data.GetTypeInfo() == 13 {
			data = data.GetValue().GetZv()
			if data.GetTypeInfo() == 0 {
				return 0
			}
		}

		/* INDIRECT element may point to UNDEF-ined slots */

	}
	for {
		if (data.GetTypeInfo() & 0xff00) != 0 {
			if data.GetType() == 10 && ZvalRefcountP(data) == 1 && (&(*data).value.GetRef().GetVal().u1.v.type_ != 7 || &(*data).value.GetRef().GetVal().value.arr != source) {
				data = &(*data).value.GetRef().GetVal()
				if (data.GetTypeInfo() & 0xff00) == 0 {
					break
				}
			}
			ZvalAddrefP(data)
		}
		break
	}
	var _z1 *Zval = &q.val
	var _z2 *Zval = data
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
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
		q.GetVal().SetNext((*uint32)(target.GetArData())[int32(nIndex)])
		(*uint32)(target.GetArData())[int32(nIndex)] = idx
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
				&q.val.u1.type_info = 0
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
	target = (*HashTable)(_emalloc(g.SizeOf("HashTable")))
	ZendGcSetRefcount(&target.gc, 1)
	target.GetGc().SetTypeInfo(7 | 1<<4<<0)
	target.SetPDestructor(ZvalPtrDtor)
	if source.GetNNumOfElements() == 0 {
		target.SetUFlags(1 << 3)
		target.SetNTableMask(uint32 - 2)
		target.SetNNumUsed(0)
		target.SetNNumOfElements(0)
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNInternalPointer(0)
		target.SetNTableSize(8)
		target.SetArData((*Bucket)((*byte)(&UninitializedBucket) + (size_t(uint32)-int32(target.GetNTableMask()))*g.SizeOf("uint32_t")))
	} else if (ZvalGcFlags(source.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		target.SetUFlags(source.GetUFlags() & 0xff)
		target.SetNTableMask(source.GetNTableMask())
		target.SetNNumUsed(source.GetNNumUsed())
		target.SetNNumOfElements(source.GetNNumOfElements())
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNTableSize(source.GetNTableSize())
		target.SetArData((*Bucket)((*byte)(_emalloc(size_t(target.GetNTableSize())*g.SizeOf("Bucket")+(size_t(uint32)-int32(target.GetNTableMask()))*g.SizeOf("uint32_t"))) + (size_t(uint32)-int32(target.GetNTableMask()))*g.SizeOf("uint32_t")))
		target.SetNInternalPointer(source.GetNInternalPointer())
		memcpy((*byte)(target.GetArData()-(size_t(uint32)-int32(target.GetNTableMask()))*g.SizeOf("uint32_t")), (*byte)(source.GetArData()-(size_t(uint32)-int32(source.GetNTableMask()))*g.SizeOf("uint32_t")), (size_t(uint32)-int32(source.GetNTableMask()))*g.SizeOf("uint32_t")+size_t(source).nNumUsed*g.SizeOf("Bucket"))
	} else if (source.GetUFlags() & 1 << 2) != 0 {
		target.SetUFlags(source.GetUFlags() & 0xff)
		target.SetNTableMask(uint32 - 2)
		target.SetNNumUsed(source.GetNNumUsed())
		target.SetNNumOfElements(source.GetNNumOfElements())
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		target.SetNTableSize(source.GetNTableSize())
		target.SetArData((*Bucket)((*byte)(_emalloc(size_t(target.GetNTableSize())*g.SizeOf("Bucket")+(size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"))) + (size_t(uint32)-int32(target.GetNTableMask()))*g.SizeOf("uint32_t")))
		if source.GetNInternalPointer() < source.GetNNumUsed() {
			target.SetNInternalPointer(source.GetNInternalPointer())
		} else {
			target.SetNInternalPointer(0)
		}
		(*uint32)(target.GetArData())[int32(-2)] = uint32 - 1
		(*uint32)(target.GetArData())[int32(-1)] = uint32 - 1
		if target.GetNNumUsed() == target.GetNNumOfElements() {
			ZendArrayDupPackedElements(source, target, 0)
		} else {
			ZendArrayDupPackedElements(source, target, 1)
		}
	} else {
		target.SetUFlags(source.GetUFlags() & 0xff)
		target.SetNTableMask(source.GetNTableMask())
		target.SetNNextFreeElement(source.GetNNextFreeElement())
		if source.GetNInternalPointer() < source.GetNNumUsed() {
			target.SetNInternalPointer(source.GetNInternalPointer())
		} else {
			target.SetNInternalPointer(0)
		}
		target.SetNTableSize(source.GetNTableSize())
		target.SetArData((*Bucket)((*byte)(_emalloc(size_t(target.GetNTableSize())*g.SizeOf("Bucket")+(size_t(uint32)-int32(target.GetNTableMask()))*g.SizeOf("uint32_t"))) + (size_t(uint32)-int32(target.GetNTableMask()))*g.SizeOf("uint32_t")))
		memset(&(*uint32)(target.GetArData())[int32(target.GetNTableMask())], uint32-1, (size_t(uint32)-int32(target.GetNTableMask()))*g.SizeOf("uint32_t"))
		if (target.GetUFlags() & (1<<2 | 1<<4)) != 0 {
			if source.GetNNumUsed() == source.GetNNumOfElements() {
				idx = ZendArrayDupElements(source, target, 1, 0)
			} else {
				idx = ZendArrayDupElements(source, target, 1, 1)
			}
		} else {
			if source.GetNNumUsed() == source.GetNNumOfElements() {
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
	if overwrite != 0 {
		for idx = 0; idx < source.GetNNumUsed(); idx++ {
			p = source.GetArData() + idx
			s = &p.val
			if s.GetType() == 13 {
				s = s.GetValue().GetZv()
			}
			if s.GetType() == 0 {
				continue
			}
			if p.GetKey() != nil {
				t = _zendHashAddOrUpdateI(target, p.GetKey(), s, 1<<0|1<<2)
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
			s = &p.val
			if s.GetType() == 13 {
				s = s.GetValue().GetZv()
			}
			if s.GetType() == 0 {
				continue
			}
			if p.GetKey() != nil {
				t = _zendHashAddOrUpdateI(target, p.GetKey(), s, 1<<1|1<<2)
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
	for idx = 0; idx < source.GetNNumUsed(); idx++ {
		p = source.GetArData() + idx
		if p.GetVal().GetType() == 0 {
			continue
		}
		if ZendHashReplaceCheckerWrapper(target, &p.val, p, pParam, pMergeSource) != 0 {
			t = ZendHashUpdate(target, p.GetKey(), &p.val)
			if pCopyConstructor != nil {
				pCopyConstructor(t)
			}
		}
	}
}

/* Returns the hash table data if found and NULL if not. */

func ZendHashFind(ht *HashTable, key *ZendString) *Zval {
	var p *Bucket
	p = ZendHashFindBucket(ht, key, 0)
	if p != nil {
		return &p.val
	} else {
		return nil
	}
}
func _zendHashFindKnownHash(ht *HashTable, key *ZendString) *Zval {
	var p *Bucket
	p = ZendHashFindBucket(ht, key, 1)
	if p != nil {
		return &p.val
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
		return &p.val
	} else {
		return nil
	}
}
func ZendHashIndexFind(ht *HashTable, h ZendUlong) *Zval {
	var p *Bucket
	if (ht.GetUFlags() & 1 << 2) != 0 {
		if h < ht.GetNNumUsed() {
			p = ht.GetArData() + h
			if p.GetVal().GetType() != 0 {
				return &p.val
			}
		}
		return nil
	}
	p = ZendHashIndexFindBucket(ht, h)
	if p != nil {
		return &p.val
	} else {
		return nil
	}
}
func _zendHashIndexFind(ht *HashTable, h ZendUlong) *Zval {
	var p *Bucket
	p = ZendHashIndexFindBucket(ht, h)
	if p != nil {
		return &p.val
	} else {
		return nil
	}
}
func ZendHashInternalPointerResetEx(ht *HashTable, pos *HashPosition) {
	*pos = _zendHashGetValidPos(ht, 0)
}

/* This function will be extremely optimized by remembering
 * the end of the list
 */

func ZendHashInternalPointerEndEx(ht *HashTable, pos *HashPosition) {
	var idx uint32
	idx = ht.GetNNumUsed()
	for idx > 0 {
		idx--
		if ht.GetArData()[idx].GetVal().GetType() != 0 {
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
			if ht.GetArData()[idx].GetVal().GetType() != 0 {
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
			if ht.GetArData()[idx].GetVal().GetType() != 0 {
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

/* This function should be made binary safe  */

func ZendHashGetCurrentKeyEx(ht *HashTable, str_index **ZendString, num_index *ZendUlong, pos *HashPosition) int {
	var idx uint32
	var p *Bucket
	idx = _zendHashGetValidPos(ht, *pos)
	if idx < ht.GetNNumUsed() {
		p = ht.GetArData() + idx
		if p.GetKey() != nil {
			*str_index = p.GetKey()
			return 1
		} else {
			*num_index = p.GetH()
			return 2
		}
	}
	return 3
}
func ZendHashGetCurrentKeyZvalEx(ht *HashTable, key *Zval, pos *HashPosition) {
	var idx uint32
	var p *Bucket
	idx = _zendHashGetValidPos(ht, *pos)
	if idx >= ht.GetNNumUsed() {
		key.SetTypeInfo(1)
	} else {
		p = ht.GetArData() + idx
		if p.GetKey() != nil {
			var __z *Zval = key
			var __s *ZendString = p.GetKey()
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
		} else {
			var __z *Zval = key
			__z.GetValue().SetLval(p.GetH())
			__z.SetTypeInfo(4)
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
			return 1
		} else {
			return 2
		}
	}
	return 3
}
func ZendHashGetCurrentDataEx(ht *HashTable, pos *HashPosition) *Zval {
	var idx uint32
	var p *Bucket
	idx = _zendHashGetValidPos(ht, *pos)
	if idx < ht.GetNNumUsed() {
		p = ht.GetArData() + idx
		return &p.val
	} else {
		return nil
	}
}
func ZendHashBucketSwap(p *Bucket, q *Bucket) {
	var val Zval
	var h ZendUlong
	var key *ZendString
	var _z1 *Zval = &val
	var _z2 *Zval = &p.val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	h = p.GetH()
	key = p.GetKey()
	var _z1 *Zval = &p.val
	var _z2 *Zval = &q.val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	p.SetH(q.GetH())
	p.SetKey(q.GetKey())
	var _z1 *Zval = &q.val
	var _z2 *Zval = &val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	q.SetH(h)
	q.SetKey(key)
}
func ZendHashBucketRenumSwap(p *Bucket, q *Bucket) {
	var val Zval
	var _z1 *Zval = &val
	var _z2 *Zval = &p.val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	var _z1 *Zval = &p.val
	var _z2 *Zval = &q.val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	var _z1 *Zval = &q.val
	var _z2 *Zval = &val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
}
func ZendHashBucketPackedSwap(p *Bucket, q *Bucket) {
	var val Zval
	var h ZendUlong
	var _z1 *Zval = &val
	var _z2 *Zval = &p.val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	h = p.GetH()
	var _z1 *Zval = &p.val
	var _z2 *Zval = &q.val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	p.SetH(q.GetH())
	var _z1 *Zval = &q.val
	var _z2 *Zval = &val
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	q.SetH(h)
}
func ZendHashSortEx(ht *HashTable, sort SortFuncT, compar CompareFuncT, renumber ZendBool) int {
	var p *Bucket
	var i uint32
	var j uint32
	if ht.GetNNumOfElements() <= 1 && !(renumber != 0 && ht.GetNNumOfElements() > 0) {
		return SUCCESS
	}
	if ht.GetNNumUsed() == ht.GetNNumOfElements() {
		i = ht.GetNNumUsed()
	} else {
		j = 0
		i = 0
		for ; j < ht.GetNNumUsed(); j++ {
			p = ht.GetArData() + j
			if p.GetVal().GetType() == 0 {
				continue
			}
			if i != j {
				ht.GetArData()[i] = *p
			}
			i++
		}
	}
	sort(any(ht.GetArData()), i, g.SizeOf("Bucket"), compar, swap_func_t(g.CondF2(renumber != 0, ZendHashBucketRenumSwap, func() __auto__ {
		if (ht.GetUFlags() & 1 << 2) != 0 {
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
	if (ht.GetUFlags() & 1 << 2) != 0 {
		if renumber == 0 {
			ZendHashPackedToHash(ht)
		}
	} else {
		if renumber != 0 {
			var new_data any
			var old_data any = (*byte)(ht.GetArData() - (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t"))
			var old_buckets *Bucket = ht.GetArData()
			if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
				new_data = __zendMalloc(size_t(ht.GetNTableSize())*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"))
			} else {
				new_data = _emalloc(size_t(ht.GetNTableSize())*g.SizeOf("Bucket") + (size_t(uint32)-int32(uint32-2))*g.SizeOf("uint32_t"))
			}
			ht.SetUFlags(ht.GetUFlags() | 1<<2 | 1<<4)
			ht.SetNTableMask(uint32 - 2)
			ht.SetArData((*Bucket)((*byte)(new_data) + (size_t(uint32)-int32(ht.GetNTableMask()))*g.SizeOf("uint32_t")))
			memcpy(ht.GetArData(), old_buckets, g.SizeOf("Bucket")*ht.GetNNumUsed())
			g.CondF((ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<7) != 0, func() { return Free(old_data) }, func() { return _efree(old_data) })
			(*uint32)(ht.GetArData())[int32(-2)] = uint32 - 1
			(*uint32)(ht.GetArData())[int32(-1)] = uint32 - 1
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
		if p1.GetVal().GetType() == 0 {
			continue
		}
		if ordered != 0 {
			for true {
				assert(idx2 != ht2.GetNNumUsed())
				p2 = ht2.GetArData() + idx2
				if p2.GetVal().GetType() != 0 {
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
			pData2 = &p2.val
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
		pData1 = &p1.val
		if pData1.GetType() == 13 {
			pData1 = pData1.GetValue().GetZv()
		}
		if pData2.GetType() == 13 {
			pData2 = pData2.GetValue().GetZv()
		}
		if pData1.GetType() == 0 {
			if pData2.GetType() != 0 {
				return -1
			}
		} else if pData2.GetType() == 0 {
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

	if (ZvalGcFlags(ht1.GetGc().GetTypeInfo()) & 1 << 5) != 0 {
		ZendErrorNoreturn(1<<0, "Nesting level too deep - recursive dependency?")
	}
	if (ZvalGcFlags(ht1.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ht1.GetGc().SetTypeInfo(ht1.GetGc().GetTypeInfo() | 1<<5<<0)
	}
	result = ZendHashCompareImpl(ht1, ht2, compar, ordered)
	if (ZvalGcFlags(ht1.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ht1.GetGc().SetTypeInfo(ht1.GetGc().GetTypeInfo() &^ (1 << 5 << 0))
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
		if ht.GetArData()[idx].GetVal().GetType() != 0 {
			break
		}
		idx++
	}
	res = ht.GetArData() + idx
	for ; idx < ht.GetNNumUsed(); idx++ {
		p = ht.GetArData() + idx
		if p.GetVal().GetType() == 0 {
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
	return &res.val
}
func _zendHandleNumericStrEx(key *byte, length int, idx *ZendUlong) int {
	var tmp *byte = key
	var end *byte = key + length
	if (*tmp) == '-' {
		tmp++
	}
	if (*tmp) == '0' && length > 1 || end-tmp > 20-1 {
		return 0
	}
	*idx = (*tmp) - '0'
	for true {
		tmp++
		if tmp == end {
			if (*key) == '-' {
				if (*idx)-1 > INT64_MAX {
					return 0
				}
				*idx = 0 - (*idx)
			} else if (*idx) > INT64_MAX {
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

/* Takes a "symtable" hashtable (contains integer and non-numeric string keys)
 * and converts it to a "proptable" (contains only string keys).
 * If the symtable didn't need duplicating, its refcount is incremented.
 */

func ZendSymtableToProptable(ht *HashTable) *HashTable {
	var num_key ZendUlong
	var str_key *ZendString
	var zv *Zval
	if (ht.GetUFlags() & 1 << 2) != 0 {
		goto convert
	}
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			str_key = _p.GetKey()
			if str_key == nil {
				goto convert
			}
		}
		break
	}
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendGcAddref(&ht.gc)
	}
	return ht
convert:
	var new_ht *HashTable = _zendNewArray(ht.GetNNumOfElements())
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
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
				if (zv.GetTypeInfo() & 0xff00) != 0 {
					if zv.GetType() == 10 && ZvalRefcountP(zv) == 1 {
						zv = &(*zv).value.GetRef().GetVal()
						if (zv.GetTypeInfo() & 0xff00) == 0 {
							break
						}
					}
					ZvalAddrefP(zv)
				}
				break
			}
			ZendHashUpdate(new_ht, str_key, zv)
		}
		break
	}
	return new_ht
}

/* Takes a "proptable" hashtable (contains only string keys) and converts it to
 * a "symtable" (contains integer and non-numeric string keys).
 * If the proptable didn't need duplicating, its refcount is incremented.
 */

func ZendProptableToSymtable(ht *HashTable, always_duplicate ZendBool) *HashTable {
	var num_key ZendUlong
	var str_key *ZendString
	var zv *Zval
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			str_key = _p.GetKey()

			/* The `str_key &&` here might seem redundant: property tables should
			 * only have string keys. Unfortunately, this isn't true, at the very
			 * least because of ArrayObject, which stores a symtable where the
			 * property table should be.
			 */

			if str_key != nil && _zendHandleNumericStr(str_key.GetVal(), str_key.GetLen(), &num_key) != 0 {
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
	if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendGcAddref(&ht.gc)
	}
	return ht
convert:
	var new_ht *HashTable = _zendNewArray(ht.GetNNumOfElements())
	for {
		var __ht *HashTable = ht
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val
			if _z.GetType() == 13 {
				_z = _z.GetValue().GetZv()
			}
			if _z.GetType() == 0 {
				continue
			}
			num_key = _p.GetH()
			str_key = _p.GetKey()
			zv = _z
			for {
				if (zv.GetTypeInfo() & 0xff00) != 0 {
					if zv.GetType() == 10 && ZvalRefcountP(zv) == 1 {
						zv = &(*zv).value.GetRef().GetVal()
						if (zv.GetTypeInfo() & 0xff00) == 0 {
							break
						}
					}
					ZvalAddrefP(zv)
				}
				break
			}

			/* Again, thank ArrayObject for `!str_key ||`. */

			if str_key == nil || _zendHandleNumericStr(str_key.GetVal(), str_key.GetLen(), &num_key) != 0 {
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
