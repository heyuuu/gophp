// <<generate>>

package zend

// #define ZEND_HASH_H

// # include "zend.h"

// #define HT_ALLOW_COW_VIOLATION(ht)

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

// #define ZEND_HASH_FILL_PACKED(ht) do { HashTable * __fill_ht = ( ht ) ; Bucket * __fill_bkt = __fill_ht -> arData + __fill_ht -> nNumUsed ; uint32_t __fill_idx = __fill_ht -> nNumUsed ; ZEND_ASSERT ( HT_FLAGS ( __fill_ht ) & HASH_FLAG_PACKED ) ;

// #define ZEND_HASH_FILL_END() __fill_ht -> nNumUsed = __fill_idx ; __fill_ht -> nNumOfElements = __fill_idx ; __fill_ht -> nNextFreeElement = __fill_idx ; __fill_ht -> nInternalPointer = 0 ; } while ( 0 )

// # include "zend.h"

// # include "zend_globals.h"

// # include "zend_variables.h"

// #define HT_ASSERT(ht,expr)

// #define IS_CONSISTENT(a)

// #define SET_INCONSISTENT(n)
