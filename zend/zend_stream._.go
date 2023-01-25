// <<generate>>

package zend

import (
	r "sik/runtime"
)

// Source: <Zend/zend_stream.h>

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
   | Authors: Wez Furlong <wez@thebrainroom.com>                          |
   |          Scott MacVicar <scottmac@php.net>                           |
   |          Nuno Lopes <nlopess@php.net>                                |
   |          Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* Lightweight stream implementation for the ZE scanners.
 * These functions are private to the engine.
 * */

type ZendStreamFsizerT func(handle any) int
type ZendStreamReaderT func(handle any, buf *byte, len_ int) ssize_t
type ZendStreamCloserT func(handle any)

const ZEND_MMAP_AHEAD = 32

type ZendStreamType = int

const (
	ZEND_HANDLE_FILENAME = iota
	ZEND_HANDLE_FP
	ZEND_HANDLE_STREAM
)

type ZendStatT = __struct__stat

const ZendFseek = r.Fseek
const ZendFtell = r.Ftell
const ZendLseek = lseek
const ZendFstat = fstat
const ZendStat = stat

// Source: <Zend/zend_stream.c>

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
   | Authors: Wez Furlong <wez@thebrainroom.com>                          |
   |          Scott MacVicar <scottmac@php.net>                           |
   |          Nuno Lopes <nlopess@php.net>                                |
   |          Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

var Isatty func(fd int) int

/* }}} */
