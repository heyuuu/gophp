// <<generate>>

package zend

// Source: <Zend/zend_alloc_sizes.h>

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

const ZEND_MM_CHUNK_SIZE = 2 * 1024 * 1024
const ZEND_MM_PAGE_SIZE = 4 * 1024
const ZEND_MM_PAGES uint32 = ZEND_MM_CHUNK_SIZE / ZEND_MM_PAGE_SIZE
const ZEND_MM_FIRST_PAGE = 1
const ZEND_MM_MIN_SMALL_SIZE = 8
const ZEND_MM_MAX_SMALL_SIZE = 3072
const ZEND_MM_MAX_LARGE_SIZE = ZEND_MM_CHUNK_SIZE - ZEND_MM_PAGE_SIZE*ZEND_MM_FIRST_PAGE

/* num, size, count, pages */
