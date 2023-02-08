// <<generate>>

package streams

import (
	"sik/core"
)

// Source: <main/streams/xp_socket.c>

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

const MSG_DONTWAIT = 0
const MSG_PEEK = 0

/* {{{ Generic socket stream operations */

var PhpStreamGenericSocketOps core.PhpStreamOps = core.MakePhpStreamOps(PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "generic_socket", nil, PhpSockopCast, PhpSockopStat, PhpSockopSetOption)
var PhpStreamSocketOps core.PhpStreamOps = core.MakePhpStreamOps(PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "tcp_socket", nil, PhpSockopCast, PhpSockopStat, PhpTcpSockopSetOption)
var PhpStreamUdpSocketOps core.PhpStreamOps = core.MakePhpStreamOps(PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "udp_socket", nil, PhpSockopCast, PhpSockopStat, PhpTcpSockopSetOption)

/* network socket operations */
