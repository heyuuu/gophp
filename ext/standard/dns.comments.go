// <<generate>>

package standard

// Source: <ext/standard/dns.c>

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
   | Authors: The typical suspects                                        |
   |          Pollita <pollita@php.net>                                   |
   |          Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* For the local hostname obtained via gethostname which is different from the
   dns-related MAXHOSTNAMELEN constant above */

/* type compat */

/* }}} */

/* {{{ proto string gethostname()
   Get the host name of the current machine */

/* }}} */

/* TODO: Reimplement the gethostby* functions using the new winxp+ API, in dns_win32.c, then
we can have a dns.c, dns_unix.c and dns_win32.c instead of a messy dns.c full of #ifdef
*/

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* Note: These functions are defined in ext/standard/dns_win32.c for Windows! */

/* just a hack to free resources allocated by glibc in __res_nsend()
 * See also:
 *   res_thread_freeres() in glibc/resolv/res_init.c
 *   __libc_res_nsend()   in resolv/res_send.c
 * */

/* {{{ proto bool dns_check_record(string host [, string type])
   Check DNS records corresponding to a given Internet host name or IP address */

/* }}} */

/* {{{ php_parserr */

/* }}} */

/* }}} */

/* }}} */
