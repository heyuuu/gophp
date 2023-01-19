// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_bitset.h>

/*
   +----------------------------------------------------------------------+
   | Zend OPcache JIT                                                     |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define _ZEND_BITSET_H_

type ZendBitset *ZendUlong

// #define ZEND_BITSET_ELM_SIZE       sizeof ( zend_ulong )

// #define ZEND_BITSET_ELM_NUM(n) ( ( n ) >> 6 )

// #define ZEND_BITSET_BIT_NUM(n) ( ( zend_ulong ) ( n ) & Z_UL ( 0x3f ) )

// #define ZEND_BITSET_ALLOCA(n,use_heap) ( zend_bitset ) do_alloca ( ( n ) * ZEND_BITSET_ELM_SIZE , use_heap )

/* Number of trailing zero bits (0x01 -> 0; 0x40 -> 6; 0x00 -> LEN) */

func ZendUlongNtz(num ZendUlong) int {
	var n int
	if num == 0 {
		return 8 * 8
	}
	n = 1
	if (num & 0xffffffff) == 0 {
		n += 32
		num = num >> 32
	}
	if (num & 0xffff) == 0 {
		n += 16
		num = num >> 16
	}
	if (num & 0xff) == 0 {
		n += 8
		num = num >> 8
	}
	if (num & 0xf) == 0 {
		n += 4
		num = num >> 4
	}
	if (num & 0x3) == 0 {
		n += 2
		num = num >> 2
	}
	return n - (num & 1)
}

/* Returns the number of zend_ulong words needed to store a bitset that is N
   bits long.  */

func ZendBitsetLen(n uint32) uint32 {
	return (n + (g.SizeOf("zend_long")*8 - 1)) / (g.SizeOf("zend_long") * 8)
}
func ZendBitsetIn(set ZendBitset, n uint32) ZendBool {
	return set[n/(g.SizeOf("( set ) [ 0 ]")*8)] >> (n&g.SizeOf("( set ) [ 0 ]")*8 - 1) & 1
}
func ZendBitsetIncl(set ZendBitset, n uint32) {
	set[n>>6] |= 1 << (zend_ulong(n) & 0x3f)
}
func ZendBitsetExcl(set ZendBitset, n uint32) {
	set[n>>6] &= ^(1 << (zend_ulong(n) & 0x3f))
}
func ZendBitsetClear(set ZendBitset, len_ uint32) {
	memset(set, 0, len_*g.SizeOf("zend_ulong"))
}
func ZendBitsetEmpty(set ZendBitset, len_ uint32) int {
	var i uint32
	for i = 0; i < len_; i++ {
		if set[i] {
			return 0
		}
	}
	return 1
}
func ZendBitsetFill(set ZendBitset, len_ uint32) {
	memset(set, 0xff, len_*g.SizeOf("zend_ulong"))
}
func ZendBitsetEqual(set1 ZendBitset, set2 ZendBitset, len_ uint32) ZendBool {
	return memcmp(set1, set2, len_*g.SizeOf("zend_ulong")) == 0
}
func ZendBitsetCopy(set1 ZendBitset, set2 ZendBitset, len_ uint32) {
	memcpy(set1, set2, len_*g.SizeOf("zend_ulong"))
}
func ZendBitsetIntersection(set1 ZendBitset, set2 ZendBitset, len_ uint32) {
	var i uint32
	for i = 0; i < len_; i++ {
		set1[i] &= set2[i]
	}
}
func ZendBitsetUnion(set1 ZendBitset, set2 ZendBitset, len_ uint32) {
	var i uint32
	for i = 0; i < len_; i++ {
		set1[i] |= set2[i]
	}
}
func ZendBitsetDifference(set1 ZendBitset, set2 ZendBitset, len_ uint32) {
	var i uint32
	for i = 0; i < len_; i++ {
		set1[i] = set1[i] & ^(set2[i])
	}
}
func ZendBitsetUnionWithIntersection(set1 ZendBitset, set2 ZendBitset, set3 ZendBitset, set4 ZendBitset, len_ uint32) {
	var i uint32
	for i = 0; i < len_; i++ {
		set1[i] = set2[i] | set3[i]&set4[i]
	}
}
func ZendBitsetUnionWithDifference(set1 ZendBitset, set2 ZendBitset, set3 ZendBitset, set4 ZendBitset, len_ uint32) {
	var i uint32
	for i = 0; i < len_; i++ {
		set1[i] = set2[i] | set3[i] & ^(set4[i])
	}
}
func ZendBitsetSubset(set1 ZendBitset, set2 ZendBitset, len_ uint32) ZendBool {
	var i uint32
	for i = 0; i < len_; i++ {
		if (set1[i] & ^(set2[i])) != 0 {
			return 0
		}
	}
	return 1
}
func ZendBitsetFirst(set ZendBitset, len_ uint32) int {
	var i uint32
	for i = 0; i < len_; i++ {
		if set[i] {
			return g.SizeOf("zend_ulong")*8*i + ZendUlongNtz(set[i])
		}
	}
	return -1
}
func ZendBitsetLast(set ZendBitset, len_ uint32) int {
	var i uint32 = len_
	for i > 0 {
		i--
		if set[i] {
			var j int = g.SizeOf("zend_ulong")*8*i - 1
			var x ZendUlong = set[i]
			for x != 0 {
				x = x >> 1
				j++
			}
			return j
		}
	}
	return -1
}

// #define ZEND_BITSET_FOREACH(set,len,bit) do { zend_bitset _set = ( set ) ; uint32_t _i , _len = ( len ) ; for ( _i = 0 ; _i < _len ; _i ++ ) { zend_ulong _x = _set [ _i ] ; if ( _x ) { ( bit ) = ZEND_BITSET_ELM_SIZE * 8 * _i ; for ( ; _x != 0 ; _x >>= Z_UL ( 1 ) , ( bit ) ++ ) { if ( ! ( _x & Z_UL ( 1 ) ) ) continue ;

// #define ZEND_BITSET_REVERSE_FOREACH(set,len,bit) do { zend_bitset _set = ( set ) ; uint32_t _i = ( len ) ; zend_ulong _test = Z_UL ( 1 ) << ( ZEND_BITSET_ELM_SIZE * 8 - 1 ) ; while ( _i -- > 0 ) { zend_ulong _x = _set [ _i ] ; if ( _x ) { ( bit ) = ZEND_BITSET_ELM_SIZE * 8 * ( _i + 1 ) - 1 ; for ( ; _x != 0 ; _x <<= Z_UL ( 1 ) , ( bit ) -- ) { if ( ! ( _x & _test ) ) continue ;

// #define ZEND_BITSET_FOREACH_END() } } } } while ( 0 )

func ZendBitsetPopFirst(set ZendBitset, len_ uint32) int {
	var i int = ZendBitsetFirst(set, len_)
	if i >= 0 {
		ZendBitsetExcl(set, i)
	}
	return i
}
