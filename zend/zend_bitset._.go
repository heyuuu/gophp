package zend

import (
	b "sik/builtin"
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

type ZendBitset *ZendUlong

const ZEND_BITSET_ELM_SIZE = b.SizeOf("zend_ulong")

/* Number of trailing zero bits (0x01 -> 0; 0x40 -> 6; 0x00 -> LEN) */

/* Returns the number of zend_ulong words needed to store a bitset that is N
   bits long.  */
