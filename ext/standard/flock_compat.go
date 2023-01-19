// <<generate>>

package standard

import g "sik/runtime/grammar"

// Source: <ext/standard/flock_compat.h>

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

// #define FLOCK_COMPAT_H

/* php_flock internally uses fcntl whether or not flock is available
 * This way our php_flock even works on NFS files.
 * More info: /usr/src/linux/Documentation
 */

/* Userland LOCK_* constants */

// #define PHP_LOCK_SH       1

// #define PHP_LOCK_EX       2

// #define PHP_LOCK_UN       3

// #define PHP_LOCK_NB       4

// Source: <ext/standard/flock_compat.c>

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

// # include "php.h"

// # include < errno . h >

// # include "ext/standard/flock_compat.h"

// # include < unistd . h >

// # include < fcntl . h >

// # include < sys / file . h >

func PhpFlock(fd int, operation int) int {
	var flck __struct__flock
	var ret int
	flck.l_len = 0
	flck.l_start = flck.l_len
	flck.l_whence = SEEK_SET
	if (operation & LOCK_SH) != 0 {
		flck.l_type = F_RDLCK
	} else if (operation & LOCK_EX) != 0 {
		flck.l_type = F_WRLCK
	} else if (operation & LOCK_UN) != 0 {
		flck.l_type = F_UNLCK
	} else {
		errno = EINVAL
		return -1
	}
	ret = fcntl(fd, g.Cond((operation&LOCK_NB) != 0, F_SETLK, F_SETLKW), &flck)
	if (operation&LOCK_NB) != 0 && ret == -1 && (errno == EACCES || errno == EAGAIN) {
		errno = EWOULDBLOCK
	}
	if ret != -1 {
		ret = 0
	}
	return ret
}

/* }}} */
