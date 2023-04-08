package streams

import (
	"github.com/heyuuu/gophp/core"
)

const MSG_DONTWAIT = 0
const MSG_PEEK = 0

var PhpStreamSocketOps core.PhpStreamOps = core.MakePhpStreamOps(PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "tcp_socket", nil, PhpSockopCast, PhpSockopStat, PhpTcpSockopSetOption)
var PhpStreamUdpSocketOps core.PhpStreamOps = core.MakePhpStreamOps(PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "udp_socket", nil, PhpSockopCast, PhpSockopStat, PhpTcpSockopSetOption)
