// <<generate>>

package core

// Source: <main/explicit_bzero.c>

/*
  +----------------------------------------------------------------------+
  | PHP Version 7                                                        |
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
  | Author:                                                              |
  +----------------------------------------------------------------------+
*/

// # include "php.h"

/*    $OpenBSD: explicit_bzero.c,v 1.4 2015/08/31 02:53:57 guenther Exp $ */

// # include < string . h >

func PhpExplicitBzero(dst any, siz int) {
	var i int = 0
	var buf *uint8 = (*volatile__unsigned__char)(dst)
	for ; i < siz; i++ {
		buf[i] = 0
	}
}
