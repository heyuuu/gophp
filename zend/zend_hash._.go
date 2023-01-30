// <<generate>>

package zend

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

const HASH_KEY_IS_STRING = 1
const HASH_KEY_IS_LONG = 2
const HASH_KEY_NON_EXISTENT = 3
const HASH_UPDATE uint32 = 1 << 0
const HASH_ADD uint32 = 1 << 1
const HASH_UPDATE_INDIRECT uint32 = 1 << 2
const HASH_ADD_NEW uint32 = 1 << 3
const HASH_ADD_NEXT = 1 << 4

/* Only the low byte are real flags */

const HASH_FLAG_MASK = 0xff

var ZendEmptyArray HashTable

type MergeCheckerFuncT func(target_ht *HashTable, source_data *Zval, hash_key *ZendHashKey, pParam any) ZendBool

/* startup/shutdown */

/* additions/updates/changes */

const ZEND_HASH_APPLY_KEEP = 0
const ZEND_HASH_APPLY_REMOVE = 1 << 0
const ZEND_HASH_APPLY_STOP = 1 << 1

type ApplyFuncT func(pDest *Zval) int
type ApplyFuncArgT func(pDest *Zval, argument any) int
type ApplyFuncArgsT func(pDest *Zval, num_args int, args []any, hash_key *ZendHashKey) int

/* This function should be used with special care (in other words,
 * it should usually not be used).  When used with the ZEND_HASH_APPLY_STOP
 * return value, it assumes things about the order of the elements in the hash.
 * Also, it does not provide the same kind of reentrancy protection that
 * the standard apply functions do.
 */

/* Deletes */

/* Data retrieval */

/* The same as zend_hash_find(), but hash value of the key must be already calculated */

/* Misc */

/* traversing */

/* Copying, merging and sorting */

/* The following macros are useful to insert a sequence of new elements
 * of packed array. They may be used instead of series of
 * zend_hash_next_index_insert_new()
 * (HashTable must have enough free buckets).
 */

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

const HT_POISONED_PTR *HashTable = (*HashTable)(intptr_t - 1)

var UninitializedBucket []uint32 = []uint32{HT_INVALID_IDX, HT_INVALID_IDX}

//ZEND_API const HashTable zend_empty_array = {

/* This is used to recurse elements and selectively delete certain entries
 * from a hashtable. apply_func() receives the data and decides if the entry
 * should be deleted or recursion should be stopped. The following three
 * return codes are possible:
 * ZEND_HASH_APPLY_KEEP   - continue
 * ZEND_HASH_APPLY_STOP   - stop iteration
 * ZEND_HASH_APPLY_REMOVE - delete the element, combineable with the former
 */

/* Returns the hash table data if found and NULL if not. */

/* This function will be extremely optimized by remembering
 * the end of the list
 */

/* This function should be made binary safe  */

/* Takes a "symtable" hashtable (contains integer and non-numeric string keys)
 * and converts it to a "proptable" (contains only string keys).
 * If the symtable didn't need duplicating, its refcount is incremented.
 */

/* Takes a "proptable" hashtable (contains only string keys) and converts it to
 * a "symtable" (contains integer and non-numeric string keys).
 * If the proptable didn't need duplicating, its refcount is incremented.
 */
