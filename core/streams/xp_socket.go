// <<generate>>

package streams

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/zend"
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

// # include "php.h"

// # include "ext/standard/file.h"

// # include "streams/php_streams_int.h"

// # include "php_network.h"

const MSG_DONTWAIT = 0
const MSG_PEEK = 0

func XP_SOCK_BUF_SIZE(sz __auto__) __auto__ { return sz }

/* {{{ Generic socket stream operations */

func PhpSockopWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.abstract)
	var didwrite ssize_t
	var ptimeout *__struct__timeval
	if sock == nil || sock.socket == -1 {
		return 0
	}
	if sock.timeout.tv_sec == -1 {
		ptimeout = nil
	} else {
		ptimeout = &sock.timeout
	}
retry:
	didwrite = send(sock.socket, buf, XP_SOCK_BUF_SIZE(count), b.Cond(sock.is_blocked && ptimeout != nil, MSG_DONTWAIT, 0))
	if didwrite <= 0 {
		var estr *byte
		var err int = core.PhpSocketErrno()
		if err == core.EWOULDBLOCK || err == EAGAIN {
			if sock.is_blocked {
				var retval int
				sock.timeout_event = 0
				for {
					retval = core.PhpPollfdFor(sock.socket, POLLOUT, ptimeout)
					if retval == 0 {
						sock.timeout_event = 1
						break
					}
					if retval > 0 {

						/* writable now; retry */

						goto retry

						/* writable now; retry */

					}
					err = core.PhpSocketErrno()
					if err != EINTR {
						break
					}
				}
			} else {

				/* EWOULDBLOCK/EAGAIN is not an error for a non-blocking stream.
				 * Report zero byte write instead. */

				return 0

				/* EWOULDBLOCK/EAGAIN is not an error for a non-blocking stream.
				 * Report zero byte write instead. */

			}
		}
		estr = core.PhpSocketStrerror(err, nil, 0)
		core.PhpErrorDocref(nil, zend.E_NOTICE, "send of "+zend.ZEND_LONG_FMT+" bytes failed with errno=%d %s", zend.ZendLong(count), err, estr)
		zend.Efree(estr)
	}
	if didwrite > 0 {
		PhpStreamNotifyProgressIncrement(core.PHP_STREAM_CONTEXT(stream), didwrite, 0)
	}
	return didwrite
}
func PhpSockStreamWaitForData(stream *core.PhpStream, sock *core.PhpNetstreamDataT) {
	var retval int
	var ptimeout *__struct__timeval
	if sock == nil || sock.socket == -1 {
		return
	}
	sock.timeout_event = 0
	if sock.timeout.tv_sec == -1 {
		ptimeout = nil
	} else {
		ptimeout = &sock.timeout
	}
	for true {
		retval = core.PhpPollfdFor(sock.socket, core.PHP_POLLREADABLE, ptimeout)
		if retval == 0 {
			sock.timeout_event = 1
		}
		if retval >= 0 {
			break
		}
		if core.PhpSocketErrno() != EINTR {
			break
		}
	}
}
func PhpSockopRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.abstract)
	var nr_bytes ssize_t = 0
	var err int
	if sock == nil || sock.socket == -1 {
		return -1
	}
	if sock.is_blocked {
		PhpSockStreamWaitForData(stream, sock)
		if sock.timeout_event {
			return 0
		}
	}
	nr_bytes = recv(sock.socket, buf, XP_SOCK_BUF_SIZE(count), b.Cond(sock.is_blocked && sock.timeout.tv_sec != -1, MSG_DONTWAIT, 0))
	err = core.PhpSocketErrno()
	if nr_bytes < 0 {
		if err == EAGAIN || err == core.EWOULDBLOCK {
			nr_bytes = 0
		} else {
			stream.eof = 1
		}
	} else if nr_bytes == 0 {
		stream.eof = 1
	}
	if nr_bytes > 0 {
		PhpStreamNotifyProgressIncrement(core.PHP_STREAM_CONTEXT(stream), nr_bytes, 0)
	}
	return nr_bytes
}
func PhpSockopClose(stream *core.PhpStream, close_handle int) int {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.abstract)
	if sock == nil {
		return 0
	}
	if close_handle != 0 {
		if sock.socket != core.SOCK_ERR {
			core.Closesocket(sock.socket)
			sock.socket = core.SOCK_ERR
		}
	}
	zend.Pefree(sock, core.PhpStreamIsPersistent(stream))
	return 0
}
func PhpSockopFlush(stream *core.PhpStream) int { return 0 }
func PhpSockopStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.abstract)
	return zend.ZendFstat(sock.socket, &ssb.sb)
}
func SockSendto(sock *core.PhpNetstreamDataT, buf *byte, buflen int, flags int, addr *__struct__sockaddr, addrlen socklen_t) int {
	var ret int
	if addr != nil {
		ret = sendto(sock.socket, buf, XP_SOCK_BUF_SIZE(buflen), flags, addr, XP_SOCK_BUF_SIZE(addrlen))
		if ret == core.SOCK_CONN_ERR {
			return -1
		} else {
			return ret
		}
	}
	if b.Assign(&ret, send(sock.socket, buf, buflen, flags)) == core.SOCK_CONN_ERR {
		return -1
	} else {
		return ret
	}
}
func SockRecvfrom(sock *core.PhpNetstreamDataT, buf *byte, buflen int, flags int, textaddr **zend.ZendString, addr **__struct__sockaddr, addrlen *socklen_t) int {
	var ret int
	var want_addr int = textaddr != nil || addr != nil
	if want_addr != 0 {
		var sa core.PhpSockaddrStorage
		var sl socklen_t = b.SizeOf("sa")
		ret = recvfrom(sock.socket, buf, XP_SOCK_BUF_SIZE(buflen), flags, (*__struct__sockaddr)(&sa), &sl)
		if ret == core.SOCK_CONN_ERR {
			ret = -1
		} else {
			ret = ret
		}
		if sl {
			core.PhpNetworkPopulateNameFromSockaddr((*__struct__sockaddr)(&sa), sl, textaddr, addr, addrlen)
		} else {
			if textaddr != nil {
				*textaddr = zend.ZSTR_EMPTY_ALLOC()
			}
			if addr != nil {
				*addr = nil
				*addrlen = 0
			}
		}
	} else {
		ret = recv(sock.socket, buf, XP_SOCK_BUF_SIZE(buflen), flags)
		if ret == core.SOCK_CONN_ERR {
			ret = -1
		} else {
			ret = ret
		}
	}
	return ret
}
func PhpSockopSetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var oldmode int
	var flags int
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.abstract)
	var xparam *PhpStreamXportParam
	if sock == nil {
		return core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	}
	switch option {
	case core.PHP_STREAM_OPTION_CHECK_LIVENESS:
		var tv __struct__timeval
		var buf byte
		var alive int = 1
		if value == -1 {
			if sock.timeout.tv_sec == -1 {
				tv.tv_sec = standard.FG(default_socket_timeout)
				tv.tv_usec = 0
			} else {
				tv = sock.timeout
			}
		} else {
			tv.tv_sec = value
			tv.tv_usec = 0
		}
		if sock.socket == -1 {
			alive = 0
		} else if core.PhpPollfdFor(sock.socket, core.PHP_POLLREADABLE|POLLPRI, &tv) > 0 {
			var ret ssize_t
			var err int
			ret = recv(sock.socket, &buf, b.SizeOf("buf"), MSG_PEEK)
			err = core.PhpSocketErrno()
			if 0 == ret || 0 > ret && err != core.EWOULDBLOCK && err != EAGAIN && err != EMSGSIZE {
				alive = 0
			}
		}
		if alive != 0 {
			return core.PHP_STREAM_OPTION_RETURN_OK
		} else {
			return core.PHP_STREAM_OPTION_RETURN_ERR
		}
	case core.PHP_STREAM_OPTION_BLOCKING:
		oldmode = sock.is_blocked
		if zend.SUCCESS == core.PhpSetSockBlocking(sock.socket, value) {
			sock.is_blocked = value
			return oldmode
		}
		return core.PHP_STREAM_OPTION_RETURN_ERR
	case core.PHP_STREAM_OPTION_READ_TIMEOUT:
		sock.timeout = *((*__struct__timeval)(ptrparam))
		sock.timeout_event = 0
		return core.PHP_STREAM_OPTION_RETURN_OK
	case core.PHP_STREAM_OPTION_META_DATA_API:
		zend.AddAssocBool((*zend.Zval)(ptrparam), "timed_out", sock.timeout_event)
		zend.AddAssocBool((*zend.Zval)(ptrparam), "blocked", sock.is_blocked)
		zend.AddAssocBool((*zend.Zval)(ptrparam), "eof", stream.eof)
		return core.PHP_STREAM_OPTION_RETURN_OK
	case core.PHP_STREAM_OPTION_XPORT_API:
		xparam = (*PhpStreamXportParam)(ptrparam)
		switch xparam.GetOp() {
		case STREAM_XPORT_OP_LISTEN:
			if listen(sock.socket, xparam.GetBacklog()) == 0 {
				xparam.SetReturncode(0)
			} else {
				xparam.SetReturncode(-1)
			}
			return core.PHP_STREAM_OPTION_RETURN_OK
		case STREAM_XPORT_OP_GET_NAME:
			xparam.SetReturncode(core.PhpNetworkGetSockName(sock.socket, b.CondF1(xparam.GetWantTextaddr() != 0, func() *zend.ZendString { return &xparam.outputs.textaddr }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return &xparam.outputs.addr }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return &xparam.outputs.addrlen }, nil)))
			return core.PHP_STREAM_OPTION_RETURN_OK
		case STREAM_XPORT_OP_GET_PEER_NAME:
			xparam.SetReturncode(core.PhpNetworkGetPeerName(sock.socket, b.CondF1(xparam.GetWantTextaddr() != 0, func() *zend.ZendString { return &xparam.outputs.textaddr }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return &xparam.outputs.addr }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return &xparam.outputs.addrlen }, nil)))
			return core.PHP_STREAM_OPTION_RETURN_OK
		case STREAM_XPORT_OP_SEND:
			flags = 0
			if (xparam.GetFlags() & STREAM_OOB) == STREAM_OOB {
				flags |= MSG_OOB
			}
			xparam.SetReturncode(SockSendto(sock, xparam.GetBuf(), xparam.GetBuflen(), flags, xparam.GetInputsAddr(), xparam.GetInputsAddrlen()))
			if xparam.GetReturncode() == -1 {
				var err *byte = core.PhpSocketStrerror(core.PhpSocketErrno(), nil, 0)
				core.PhpErrorDocref(nil, zend.E_WARNING, "%s\n", err)
				zend.Efree(err)
			}
			return core.PHP_STREAM_OPTION_RETURN_OK
		case STREAM_XPORT_OP_RECV:
			flags = 0
			if (xparam.GetFlags() & STREAM_OOB) == STREAM_OOB {
				flags |= MSG_OOB
			}
			if (xparam.GetFlags() & STREAM_PEEK) == STREAM_PEEK {
				flags |= MSG_PEEK
			}
			xparam.SetReturncode(SockRecvfrom(sock, xparam.GetBuf(), xparam.GetBuflen(), flags, b.CondF1(xparam.GetWantTextaddr() != 0, func() *zend.ZendString { return &xparam.outputs.textaddr }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return &xparam.outputs.addr }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return &xparam.outputs.addrlen }, nil)))
			return core.PHP_STREAM_OPTION_RETURN_OK
		case STREAM_XPORT_OP_SHUTDOWN:
			var shutdown_how []int = []int{core.SHUT_RD, core.SHUT_WR, core.SHUT_RDWR}
			xparam.SetReturncode(shutdown(sock.socket, shutdown_how[xparam.GetHow()]))
			return core.PHP_STREAM_OPTION_RETURN_OK
		default:
			return core.PHP_STREAM_OPTION_RETURN_NOTIMPL
		}
	default:
		return core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	}
}
func PhpSockopCast(stream *core.PhpStream, castas int, ret *any) int {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.abstract)
	if sock == nil {
		return zend.FAILURE
	}
	switch castas {
	case core.PHP_STREAM_AS_STDIO:
		if ret != nil {
			*((**r.FILE)(ret)) = fdopen(sock.socket, stream.mode)
			if *ret {
				return zend.SUCCESS
			}
			return zend.FAILURE
		}
		return zend.SUCCESS
	case core.PHP_STREAM_AS_FD_FOR_SELECT:

	case core.PHP_STREAM_AS_FD:

	case core.PHP_STREAM_AS_SOCKETD:
		if ret != nil {
			*((*core.PhpSocketT)(ret)) = sock.socket
		}
		return zend.SUCCESS
	default:
		return zend.FAILURE
	}
}

/* }}} */

var PhpStreamGenericSocketOps core.PhpStreamOps = core.PhpStreamOps{PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "generic_socket", nil, PhpSockopCast, PhpSockopStat, PhpSockopSetOption}
var PhpStreamSocketOps core.PhpStreamOps = core.PhpStreamOps{PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "tcp_socket", nil, PhpSockopCast, PhpSockopStat, PhpTcpSockopSetOption}
var PhpStreamUdpSocketOps core.PhpStreamOps = core.PhpStreamOps{PhpSockopWrite, PhpSockopRead, PhpSockopClose, PhpSockopFlush, "udp_socket", nil, PhpSockopCast, PhpSockopStat, PhpTcpSockopSetOption}

/* network socket operations */

func ParseIpAddressEx(str *byte, str_len int, portno *int, get_err int, err **zend.ZendString) *byte {
	var colon *byte
	var host *byte = nil
	var p *byte
	if (*str) == '[' && str_len > 1 {

		/* IPV6 notation to specify raw address with port (i.e. [fe80::1]:80) */

		p = memchr(str+1, ']', str_len-2)
		if p == nil || (*(p + 1)) != ':' {
			if get_err != 0 {
				*err = core.Strpprintf(0, "Failed to parse IPv6 address \"%s\"", str)
			}
			return nil
		}
		*portno = atoi(p + 2)
		return zend.Estrndup(str+1, p-str-1)
	}
	if str_len != 0 {
		colon = memchr(str, ':', str_len-1)
	} else {
		colon = nil
	}
	if colon != nil {
		*portno = atoi(colon + 1)
		host = zend.Estrndup(str, colon-str)
	} else {
		if get_err != 0 {
			*err = core.Strpprintf(0, "Failed to parse address \"%s\"", str)
		}
		return nil
	}
	return host
}
func ParseIpAddress(xparam *PhpStreamXportParam, portno *int) *byte {
	return ParseIpAddressEx(xparam.GetName(), xparam.GetNamelen(), portno, xparam.GetWantErrortext(), &xparam.outputs.error_text)
}
func PhpTcpSockopBind(stream *core.PhpStream, sock *core.PhpNetstreamDataT, xparam *PhpStreamXportParam) int {
	var host *byte = nil
	var portno int
	var err int
	var sockopts long = core.STREAM_SOCKOP_NONE
	var tmpzval *zend.Zval = nil
	host = ParseIpAddress(xparam, &portno)
	if host == nil {
		return -1
	}
	sock.socket = core.PhpNetworkBindSocketToLocalAddr(host, portno, b.Cond(stream.ops == &PhpStreamUdpSocketOps, SOCK_DGRAM, SOCK_STREAM), sockopts, b.CondF1(xparam.GetWantErrortext() != 0, func() *zend.ZendString { return &xparam.outputs.error_text }, nil), &err)
	if host != nil {
		zend.Efree(host)
	}
	if sock.socket == -1 {
		return -1
	} else {
		return 0
	}
}
func PhpTcpSockopConnect(stream *core.PhpStream, sock *core.PhpNetstreamDataT, xparam *PhpStreamXportParam) int {
	var host *byte = nil
	var bindto *byte = nil
	var portno int
	var bindport int = 0
	var err int = 0
	var ret int
	var tmpzval *zend.Zval = nil
	var sockopts long = core.STREAM_SOCKOP_NONE
	host = ParseIpAddress(xparam, &portno)
	if host == nil {
		return -1
	}
	if core.PHP_STREAM_CONTEXT(stream) != nil && b.Assign(&tmpzval, PhpStreamContextGetOption(core.PHP_STREAM_CONTEXT(stream), "socket", "bindto")) != nil {
		if zend.Z_TYPE_P(tmpzval) != zend.IS_STRING {
			if xparam.GetWantErrortext() != 0 {
				xparam.SetErrorText(core.Strpprintf(0, "local_addr context option is not a string."))
			}
			zend.Efree(host)
			return -1
		}
		bindto = ParseIpAddressEx(zend.Z_STRVAL_P(tmpzval), zend.Z_STRLEN_P(tmpzval), &bindport, xparam.GetWantErrortext(), &xparam.outputs.error_text)
	}
	if stream.ops != &PhpStreamUdpSocketOps && core.PHP_STREAM_CONTEXT(stream) != nil && b.Assign(&tmpzval, PhpStreamContextGetOption(core.PHP_STREAM_CONTEXT(stream), "socket", "tcp_nodelay")) != nil && zend.ZendIsTrue(tmpzval) != 0 {
		sockopts |= core.STREAM_SOCKOP_TCP_NODELAY
	}

	/* Note: the test here for php_stream_udp_socket_ops is important, because we
	 * want the default to be TCP sockets so that the openssl extension can
	 * re-use this code. */

	sock.socket = core.PhpNetworkConnectSocketToHost(host, portno, b.Cond(stream.ops == &PhpStreamUdpSocketOps, SOCK_DGRAM, SOCK_STREAM), xparam.GetOp() == STREAM_XPORT_OP_CONNECT_ASYNC, xparam.GetTimeout(), b.CondF1(xparam.GetWantErrortext() != 0, func() *zend.ZendString { return &xparam.outputs.error_text }, nil), &err, bindto, bindport, sockopts)
	if sock.socket == -1 {
		ret = -1
	} else {
		ret = 0
	}
	xparam.SetErrorCode(err)
	if host != nil {
		zend.Efree(host)
	}
	if bindto != nil {
		zend.Efree(bindto)
	}
	if ret >= 0 && xparam.GetOp() == STREAM_XPORT_OP_CONNECT_ASYNC && err == EINPROGRESS {

		/* indicates pending connection */

		return 1

		/* indicates pending connection */

	}
	return ret
}
func PhpTcpSockopAccept(stream *core.PhpStream, sock *core.PhpNetstreamDataT, xparam *PhpStreamXportParam) int {
	var clisock int
	var nodelay zend.ZendBool = 0
	var tmpzval *zend.Zval = nil
	xparam.SetClient(nil)
	if nil != core.PHP_STREAM_CONTEXT(stream) && b.Assign(&tmpzval, PhpStreamContextGetOption(core.PHP_STREAM_CONTEXT(stream), "socket", "tcp_nodelay")) != nil && zend.ZendIsTrue(tmpzval) != 0 {
		nodelay = 1
	}
	clisock = core.PhpNetworkAcceptIncoming(sock.socket, b.CondF1(xparam.GetWantTextaddr() != 0, func() *zend.ZendString { return &xparam.outputs.textaddr }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return &xparam.outputs.addr }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return &xparam.outputs.addrlen }, nil), xparam.GetTimeout(), b.CondF1(xparam.GetWantErrortext() != 0, func() *zend.ZendString { return &xparam.outputs.error_text }, nil), &xparam.outputs.error_code, nodelay)
	if clisock >= 0 {
		var clisockdata *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(zend.Emalloc(b.SizeOf("* clisockdata")))
		memcpy(clisockdata, sock, b.SizeOf("* clisockdata"))
		clisockdata.socket = clisock
		xparam.SetClient(core.PhpStreamAllocRel(stream.ops, clisockdata, nil, "r+"))
		if xparam.GetClient() != nil {
			xparam.GetClient().ctx = stream.ctx
			if stream.ctx != nil {
				zend.GC_ADDREF(stream.ctx)
			}
		}
	}
	if xparam.GetClient() == nil {
		return -1
	} else {
		return 0
	}
}
func PhpTcpSockopSetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.abstract)
	var xparam *PhpStreamXportParam
	switch option {
	case core.PHP_STREAM_OPTION_XPORT_API:
		xparam = (*PhpStreamXportParam)(ptrparam)
		switch xparam.GetOp() {
		case STREAM_XPORT_OP_CONNECT:

		case STREAM_XPORT_OP_CONNECT_ASYNC:
			xparam.SetReturncode(PhpTcpSockopConnect(stream, sock, xparam))
			return core.PHP_STREAM_OPTION_RETURN_OK
		case STREAM_XPORT_OP_BIND:
			xparam.SetReturncode(PhpTcpSockopBind(stream, sock, xparam))
			return core.PHP_STREAM_OPTION_RETURN_OK
		case STREAM_XPORT_OP_ACCEPT:
			xparam.SetReturncode(PhpTcpSockopAccept(stream, sock, xparam))
			return core.PHP_STREAM_OPTION_RETURN_OK
		default:

			/* fall through */

		}
	}
	return PhpSockopSetOption(stream, option, value, ptrparam)
}
func PhpStreamGenericSocketFactory(proto *byte, protolen int, resourcename *byte, resourcenamelen int, persistent_id *byte, options int, flags int, timeout *__struct__timeval, context *core.PhpStreamContext) *core.PhpStream {
	var stream *core.PhpStream = nil
	var sock *core.PhpNetstreamDataT
	var ops *core.PhpStreamOps

	/* which type of socket ? */

	if strncmp(proto, "tcp", protolen) == 0 {
		ops = &PhpStreamSocketOps
	} else if strncmp(proto, "udp", protolen) == 0 {
		ops = &PhpStreamUdpSocketOps
	} else {

		/* should never happen */

		return nil

		/* should never happen */

	}
	sock = zend.Pemalloc(b.SizeOf("php_netstream_data_t"), b.Cond(persistent_id != nil, 1, 0))
	memset(sock, 0, b.SizeOf("php_netstream_data_t"))
	sock.is_blocked = 1
	sock.timeout.tv_sec = standard.FG(default_socket_timeout)
	sock.timeout.tv_usec = 0

	/* we don't know the socket until we have determined if we are binding or
	 * connecting */

	sock.socket = -1
	stream = core.PhpStreamAllocRel(ops, sock, persistent_id, "r+")
	if stream == nil {
		zend.Pefree(sock, b.Cond(persistent_id != nil, 1, 0))
		return nil
	}
	if flags == 0 {
		return stream
	}
	return stream
}
