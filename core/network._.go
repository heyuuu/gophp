// <<generate>>

package core

// Source: <main/network.c>

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
   | Author: Stig Venaas <venaas@uninett.no>                              |
   | Streams work by Wez Furlong <wez@thebrainroom.com>                   |
   +----------------------------------------------------------------------+
*/

const SOCK_ERR = -1
const SOCK_CONN_ERR = -1
const PHP_TIMEOUT_ERROR_VALUE = ETIMEDOUT

/* {{{ php_network_freeaddresses
 */

/* }}} */

/* }}} */

const O_NONBLOCK = O_NDELAY

type PhpNonBlockingFlagsT = int

/* Connect to a socket using an interruptible connect with optional timeout.
 * Optionally, the connect can be made asynchronously, which will implicitly
 * enable non-blocking mode on the socket.
 * */

/* }}} */

/* }}} */

/* }}} */

/* Accept a client connection from a server socket,
 * using an optional timeout.
 * Returns the peer address in addr/addrlen (it will emalloc
 * these, so be sure to efree the result).
 * If you specify textaddr, a text-printable
 * version of the address will be emalloc'd and returned.
 * */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */
