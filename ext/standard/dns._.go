// <<generate>>

package standard

import (
	"sik/zend"
)

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

const BIND_8_COMPAT = 1
const MAXHOSTNAMELEN = 255

/* For the local hostname obtained via gethostname which is different from the
   dns-related MAXHOSTNAMELEN constant above */

const HOST_NAME_MAX = 255

/* type compat */

const DNS_T_A = 1
const DNS_T_NS = 2
const DNS_T_CNAME = 5
const DNS_T_SOA = 6
const DNS_T_PTR = 12
const DNS_T_HINFO = 13
const DNS_T_MINFO = 14
const DNS_T_MX = 15
const DNS_T_TXT = 16
const DNS_T_AAAA = 28
const DNS_T_SRV = 33
const DNS_T_NAPTR = 35
const DNS_T_A6 = 38
const DNS_T_CAA = 257
const DNS_T_ANY = 255

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

const PHP_DNS_NUM_TYPES = 13
const PHP_DNS_A = 0x1
const PHP_DNS_NS = 0x2
const PHP_DNS_CNAME = 0x10
const PHP_DNS_SOA = 0x20
const PHP_DNS_PTR = 0x800
const PHP_DNS_HINFO = 0x1000
const PHP_DNS_CAA = 0x2000
const PHP_DNS_MX = 0x4000
const PHP_DNS_TXT = 0x8000
const PHP_DNS_A6 = 0x1000000
const PHP_DNS_SRV = 0x2000000
const PHP_DNS_NAPTR = 0x4000000
const PHP_DNS_AAAA = 0x8000000
const PHP_DNS_ANY = 0x10000000
const PHP_DNS_ALL zend.ZendLong = PHP_DNS_A | PHP_DNS_NS | PHP_DNS_CNAME | PHP_DNS_SOA | PHP_DNS_PTR | PHP_DNS_HINFO | PHP_DNS_CAA | PHP_DNS_MX | PHP_DNS_TXT | PHP_DNS_A6 | PHP_DNS_SRV | PHP_DNS_NAPTR | PHP_DNS_AAAA

/* Note: These functions are defined in ext/standard/dns_win32.c for Windows! */

const HFIXEDSZ = 12
const QFIXEDSZ = 4
const MAXRESOURCERECORDS = 64

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
