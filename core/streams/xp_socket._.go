package streams

import (
	"sik/core"
)

const MSG_DONTWAIT = 0
const MSG_PEEK = 0

/* {{{ Generic socket stream operations */

var PhpStreamSocketOps core.PhpStreamOps = core.MakePhpStreamOps(PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "tcp_socket", nil, PhpSockopCast, PhpSockopStat, PhpTcpSockopSetOption)
var PhpStreamUdpSocketOps core.PhpStreamOps = core.MakePhpStreamOps(PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "udp_socket", nil, PhpSockopCast, PhpSockopStat, PhpTcpSockopSetOption)

/* network socket operations */
