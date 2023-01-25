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

/* Only the low byte are real flags */

/* startup/shutdown */

/* additions/updates/changes */

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

//ZEND_API const HashTable zend_empty_array = {

/* }}} */

/* }}} */

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
