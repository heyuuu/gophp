// <<generate>>

package core

// Source: <main/reentrancy.c>

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
   | Author: Sascha Schumann <sascha@schumann.cx>                         |
   +----------------------------------------------------------------------+
*/

// # include < sys / types . h >

// # include < string . h >

// # include < errno . h >

// # include "php_reentrancy.h"

// # include "ext/standard/php_rand.h"

const (
	LOCALTIME_R = iota
	CTIME_R
	ASCTIME_R
	GMTIME_R
	READDIR_R
	NUMBER_OF_LOCKS
)

// #define local_lock(x)

// #define local_unlock(x)
