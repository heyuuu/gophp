// <<generate>>

package streams

import (
	"sik/core"
)

const MSG_DONTWAIT = 0
const MSG_PEEK = 0

var PhpStreamGenericSocketOps core.PhpStreamOps = core.PhpStreamOps{PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "generic_socket", nil, PhpSockopCast, PhpSockopStat, PhpSockopSetOption}
var PhpStreamSocketOps core.PhpStreamOps = core.PhpStreamOps{PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "tcp_socket", nil, PhpSockopCast, PhpSockopStat, PhpTcpSockopSetOption}
var PhpStreamUdpSocketOps core.PhpStreamOps = core.PhpStreamOps{PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "udp_socket", nil, PhpSockopCast, PhpSockopStat, PhpTcpSockopSetOption}
