// <<generate>>

package standard

// Source: <ext/standard/php_dns.h>

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
   |          Marcus Boerger <helly@php.net>                              |
   |          Pollita <pollita@php.net>                                   |
   +----------------------------------------------------------------------+
*/

// #define PHP_DNS_H

func PhpDnsSearch(res dns_handle_t, dname *byte, class __auto__, type_ int, answer []u_char, anslen __auto__) int {
	return int(dns_search(res, dname, class, type_, (*byte)(answer), anslen, (*__struct__sockaddr)(&from), &fromsize))
}
func PhpDnsFreeHandle(res dns_handle_t) __auto__ { return dns_free(res) }
func PhpDnsErrno(handle dns_handle_t) __auto__   { return h_errno }

const HAVE_DNS_SEARCH_FUNC = 1
const HAVE_FULL_DNS_FUNCS = 1
const INT16SZ = 2
const INT32SZ = 4
