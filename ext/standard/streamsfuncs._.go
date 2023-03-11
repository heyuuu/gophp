// <<generate>>

package standard

import (
	"sik/zend"
)

// Source: <ext/standard/streamsfuncs.h>

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
  | Authors: Wez Furlong <wez@thebrainroom.com>                          |
  +----------------------------------------------------------------------+
*/

const PHP_STREAM_CLIENT_PERSISTENT = 1
const PHP_STREAM_CLIENT_ASYNC_CONNECT = 2
const PHP_STREAM_CLIENT_CONNECT = 4

var ZifStreamWrapperRegister func(executeData *zend.ZendExecuteData, return_value *zend.Zval)
var ZifStreamWrapperUnregister func(executeData *zend.ZendExecuteData, return_value *zend.Zval)
var ZifStreamWrapperRestore func(executeData *zend.ZendExecuteData, return_value *zend.Zval)

// Source: <ext/standard/streamsfuncs.c>

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
  | Authors: Wez Furlong <wez@thebrainroom.com>                          |
  |          Sara Golemon <pollita@php.net>                              |
  +----------------------------------------------------------------------+
*/

type PhpTimeoutUll = unsigned__long__long

/* Streams based network functions */

/* {{{ proto array stream_socket_pair(int domain, int type, int protocol)
   Creates a pair of connected, indistinguishable socket streams */

/* {{{ proto resource stream_socket_client(string remoteaddress [, int &errcode [, string &errstring [, double timeout [, int flags [, resource context]]]]])
   Open a client connection to a remote address */

/* given a zval which is either a stream or a context, return the underlying
 * stream_context.  If it is a stream that does not have a context assigned, it
 * will create and assign a context and return that.  */

/* {{{ proto bool stream_isatty(resource stream)
Check if a stream is a TTY.
*/

/* {{{ proto int stream_socket_shutdown(resource stream, int how)
   causes all or part of a full-duplex connection on the socket associated
   with stream to be shut down.  If how is SHUT_RD,  further receptions will
   be disallowed. If how is SHUT_WR, further transmissions will be disallowed.
   If how is SHUT_RDWR,  further  receptions and transmissions will be
   disallowed. */
