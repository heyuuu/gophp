// <<generate>>

package standard

// Source: <ext/standard/proc_open.h>

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
   | Author: Wez Furlong <wez@thebrainroom.com>                           |
   +----------------------------------------------------------------------+
*/

type PhpFileDescriptorT = int
type PhpProcessIdT = pid_t

/* Environment block under win32 is a NUL terminated sequence of NUL terminated
 * name=value strings.
 * Under unix, it is an argv style array.
 * */

type _phpProcessEnv = PhpProcessEnvT

// Source: <ext/standard/proc_open.c>

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
   | Author: Wez Furlong <wez@thebrainroom.com>                           |
   +----------------------------------------------------------------------+
*/

/* This symbol is defined in ext/standard/config.m4.
 * Essentially, it is set if you HAVE_FORK || PHP_WIN32
 * Other platforms may modify that configure check and add suitable #ifdefs
 * around the alternate code.
 * */

var LeProcOpen int

/* {{{ _php_array_to_envp */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

const DESC_PIPE = 1
const DESC_FILE = 2
const DESC_REDIRECT = 3
const DESC_PARENT_MODE_WRITE = 8

/* }}} */

/* {{{ proto resource proc_open(string|array command, array descriptorspec, array &pipes [, string cwd [, array env [, array other_options]]])
   Run a process with more control over it's file descriptors */

/* }}} */
