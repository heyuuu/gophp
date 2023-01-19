// <<generate>>

package core

// Source: <main/php_reentrancy.h>

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

// #define PHP_REENTRANCY_H

// # include "php.h"

// # include < sys / types . h >

// # include < dirent . h >

// # include < time . h >

/* currently, PHP does not check for these functions, but assumes
   that they are available on all systems. */

// #define HAVE_LOCALTIME       1

// #define HAVE_GMTIME       1

// #define HAVE_ASCTIME       1

// #define HAVE_CTIME       1

// #define php_localtime_r       localtime_r

// #define php_ctime_r       ctime_r

// #define php_asctime_r       asctime_r

// #define php_gmtime_r       gmtime_r

// #define php_strtok_r       strtok_r

// #define php_rand_r       rand_r

// #define reentrancy_startup()

// #define reentrancy_shutdown()
