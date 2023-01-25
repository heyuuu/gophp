// <<generate>>

package zend

// Source: <Zend/zend_map_ptr.h>

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
   | Authors: Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define ZEND_MAP_PTR_H

// # include "zend_portability.h"

const ZEND_MAP_PTR_KIND_PTR = 0
const ZEND_MAP_PTR_KIND_PTR_OR_OFFSET = 1

//#if defined(ZTS) || defined(TSRM_WIN32)

const ZEND_MAP_PTR_KIND = ZEND_MAP_PTR_KIND_PTR_OR_OFFSET

//#else

// #define ZEND_MAP_PTR(ptr) ptr ## __ptr

// #define ZEND_MAP_PTR_DEF(type,name) type * ZEND_MAP_PTR ( name )

func ZEND_MAP_PTR_IS_OFFSET(ptr __auto__) int { return uintPtr(ptr__ptr) & 1 }
func ZEND_MAP_PTR_OFFSET2PTR(ptr __auto__) *any {
	return (*any)((*byte)(CG(map_ptr_base) + uintPtr(ptr__ptr-1)))
}
func ZEND_MAP_PTR_PTR2OFFSET(ptr *any) any {
	return any(uintptr_t((*byte)(ptr)-(*byte)(CG(map_ptr_base))) | 1)
}
func ZEND_MAP_PTR_GET(ptr __auto__) any {
	if ZEND_MAP_PTR_IS_OFFSET(ptr) != 0 {
		return *(ZEND_MAP_PTR_OFFSET2PTR(ptr))
	} else {
		return any(*ptr__ptr)
	}
}
func ZEND_MAP_PTR_SET(ptr __auto__, val any) {
	if ZEND_MAP_PTR_IS_OFFSET(ptr) != 0 {
		*(ZEND_MAP_PTR_OFFSET2PTR(ptr)) = val
	} else {
		*ptr__ptr = val
	}
}
func ZEND_MAP_PTR_INIT(ptr __auto__, val __auto__) { ptr__ptr = val }
func ZEND_MAP_PTR_NEW(ptr __auto__)                { ptr__ptr = ZendMapPtrNew() }
