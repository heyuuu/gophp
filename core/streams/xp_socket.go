// <<generate>>

package streams

import (
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	g "sik/runtime/grammar"
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

// #define MSG_DONTWAIT       0

// #define MSG_PEEK       0

// #define XP_SOCK_BUF_SIZE(sz) ( sz )

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
	didwrite = send(sock.socket, buf, count, g.Cond(sock.is_blocked && ptimeout != nil, 0, 0))
	if didwrite <= 0 {
		var estr *byte
		var err int = errno
		if err == EAGAIN || err == EAGAIN {
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
					err = errno
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
		core.PhpErrorDocref(nil, 1<<3, "send of "+"%"+"lld"+" bytes failed with errno=%d %s", zend.ZendLong(count), err, estr)
		zend._efree(estr)
	}
	if didwrite > 0 {
		if (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)) != nil && (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier() != nil && ((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().GetMask()&1) != 0 {
			(*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().SetProgress((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().GetProgress() + didwrite)
			(*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().SetProgressMax((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().GetProgressMax() + 0)
			if (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)) != nil && (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier() != nil {
				PhpStreamNotificationNotify((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)), 7, 0, nil, 0, (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().GetProgress(), (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().GetProgressMax(), nil)
			}
		}
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
		retval = core.PhpPollfdFor(sock.socket, POLLIN|POLLERR|POLLHUP, ptimeout)
		if retval == 0 {
			sock.timeout_event = 1
		}
		if retval >= 0 {
			break
		}
		if errno != EINTR {
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
	nr_bytes = recv(sock.socket, buf, count, g.Cond(sock.is_blocked && sock.timeout.tv_sec != -1, 0, 0))
	err = errno
	if nr_bytes < 0 {
		if err == EAGAIN || err == EAGAIN {
			nr_bytes = 0
		} else {
			stream.eof = 1
		}
	} else if nr_bytes == 0 {
		stream.eof = 1
	}
	if nr_bytes > 0 {
		if (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)) != nil && (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier() != nil && ((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().GetMask()&1) != 0 {
			(*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().SetProgress((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().GetProgress() + nr_bytes)
			(*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().SetProgressMax((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().GetProgressMax() + 0)
			if (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)) != nil && (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier() != nil {
				PhpStreamNotificationNotify((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)), 7, 0, nil, 0, (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().GetProgress(), (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)).GetNotifier().GetProgressMax(), nil)
			}
		}
	}
	return nr_bytes
}
func PhpSockopClose(stream *core.PhpStream, close_handle int) int {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.abstract)
	if sock == nil {
		return 0
	}
	if close_handle != 0 {
		if sock.socket != -1 {
			close(sock.socket)
			sock.socket = -1
		}
	}
	g.CondF(stream.is_persistent != 0, func() { return zend.Free(sock) }, func() { return zend._efree(sock) })
	return 0
}
func PhpSockopFlush(stream *core.PhpStream) int { return 0 }
func PhpSockopStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.abstract)
	return fstat(sock.socket, &ssb.sb)
}
func SockSendto(sock *core.PhpNetstreamDataT, buf *byte, buflen int, flags int, addr *__struct__sockaddr, addrlen socklen_t) int {
	var ret int
	if addr != nil {
		ret = sendto(sock.socket, buf, buflen, flags, addr, addrlen)
		if ret == -1 {
			return -1
		} else {
			return ret
		}
	}
	if g.Assign(&ret, send(sock.socket, buf, buflen, flags)) == -1 {
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
		var sl socklen_t = g.SizeOf("sa")
		ret = recvfrom(sock.socket, buf, buflen, flags, (*__struct__sockaddr)(&sa), &sl)
		if ret == -1 {
			ret = -1
		} else {
			ret = ret
		}
		if sl {
			core.PhpNetworkPopulateNameFromSockaddr((*__struct__sockaddr)(&sa), sl, textaddr, addr, addrlen)
		} else {
			if textaddr != nil {
				*textaddr = zend.ZendEmptyString
			}
			if addr != nil {
				*addr = nil
				*addrlen = 0
			}
		}
	} else {
		ret = recv(sock.socket, buf, buflen, flags)
		if ret == -1 {
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
		return -2
	}
	switch option {
	case 12:
		var tv __struct__timeval
		var buf byte
		var alive int = 1
		if value == -1 {
			if sock.timeout.tv_sec == -1 {
				tv.tv_sec = standard.FileGlobals.default_socket_timeout
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
		} else if core.PhpPollfdFor(sock.socket, POLLIN|POLLERR|POLLHUP|POLLPRI, &tv) > 0 {
			var ret ssize_t
			var err int
			ret = recv(sock.socket, &buf, g.SizeOf("buf"), 0)
			err = errno
			if 0 == ret || 0 > ret && err != EAGAIN && err != EAGAIN && err != EMSGSIZE {
				alive = 0
			}
		}
		if alive != 0 {
			return 0
		} else {
			return -1
		}
	case 1:
		oldmode = sock.is_blocked
		if zend.SUCCESS == core.PhpSetSockBlocking(sock.socket, value) {
			sock.is_blocked = value
			return oldmode
		}
		return -1
	case 4:
		sock.timeout = *((*__struct__timeval)(ptrparam))
		sock.timeout_event = 0
		return 0
	case 11:
		zend.AddAssocBoolEx((*zend.Zval)(ptrparam), "timed_out", strlen("timed_out"), sock.timeout_event)
		zend.AddAssocBoolEx((*zend.Zval)(ptrparam), "blocked", strlen("blocked"), sock.is_blocked)
		zend.AddAssocBoolEx((*zend.Zval)(ptrparam), "eof", strlen("eof"), stream.eof)
		return 0
	case 7:
		xparam = (*PhpStreamXportParam)(ptrparam)
		switch xparam.GetOp() {
		case STREAM_XPORT_OP_LISTEN:
			if listen(sock.socket, xparam.GetBacklog()) == 0 {
				xparam.SetReturncode(0)
			} else {
				xparam.SetReturncode(-1)
			}
			return 0
		case STREAM_XPORT_OP_GET_NAME:
			xparam.SetReturncode(core.PhpNetworkGetSockName(sock.socket, g.CondF1(xparam.GetWantTextaddr() != 0, func() *zend.ZendString { return &xparam.outputs.textaddr }, nil), g.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return &xparam.outputs.addr }, nil), g.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return &xparam.outputs.addrlen }, nil)))
			return 0
		case STREAM_XPORT_OP_GET_PEER_NAME:
			xparam.SetReturncode(core.PhpNetworkGetPeerName(sock.socket, g.CondF1(xparam.GetWantTextaddr() != 0, func() *zend.ZendString { return &xparam.outputs.textaddr }, nil), g.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return &xparam.outputs.addr }, nil), g.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return &xparam.outputs.addrlen }, nil)))
			return 0
		case STREAM_XPORT_OP_SEND:
			flags = 0
			if (xparam.GetFlags() & STREAM_OOB) == STREAM_OOB {
				flags |= MSG_OOB
			}
			xparam.SetReturncode(SockSendto(sock, xparam.GetBuf(), xparam.GetBuflen(), flags, xparam.GetInputsAddr(), xparam.GetInputsAddrlen()))
			if xparam.GetReturncode() == -1 {
				var err *byte = core.PhpSocketStrerror(errno, nil, 0)
				core.PhpErrorDocref(nil, 1<<1, "%s\n", err)
				zend._efree(err)
			}
			return 0
		case STREAM_XPORT_OP_RECV:
			flags = 0
			if (xparam.GetFlags() & STREAM_OOB) == STREAM_OOB {
				flags |= MSG_OOB
			}
			if (xparam.GetFlags() & STREAM_PEEK) == STREAM_PEEK {
				flags |= 0
			}
			xparam.SetReturncode(SockRecvfrom(sock, xparam.GetBuf(), xparam.GetBuflen(), flags, g.CondF1(xparam.GetWantTextaddr() != 0, func() *zend.ZendString { return &xparam.outputs.textaddr }, nil), g.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return &xparam.outputs.addr }, nil), g.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return &xparam.outputs.addrlen }, nil)))
			return 0
		case STREAM_XPORT_OP_SHUTDOWN:
			var shutdown_how []int = []int{0, 1, 2}
			xparam.SetReturncode(shutdown(sock.socket, shutdown_how[xparam.GetHow()]))
			return 0
		default:
			return -2
		}
	default:
		return -2
	}
}
func PhpSockopCast(stream *core.PhpStream, castas int, ret *any) int {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.abstract)
	if sock == nil {
		return zend.FAILURE
	}
	switch castas {
	case 0:
		if ret != nil {
			*((**r.FILE)(ret)) = fdopen(sock.socket, stream.mode)
			if *ret {
				return zend.SUCCESS
			}
			return zend.FAILURE
		}
		return zend.SUCCESS
	case 3:

	case 1:

	case 2:
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
				*err = zend.ZendStrpprintf(0, "Failed to parse IPv6 address \"%s\"", str)
			}
			return nil
		}
		*portno = atoi(p + 2)
		return zend._estrndup(str+1, p-str-1)
	}
	if str_len != 0 {
		colon = memchr(str, ':', str_len-1)
	} else {
		colon = nil
	}
	if colon != nil {
		*portno = atoi(colon + 1)
		host = zend._estrndup(str, colon-str)
	} else {
		if get_err != 0 {
			*err = zend.ZendStrpprintf(0, "Failed to parse address \"%s\"", str)
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
	var sockopts long = 1 << 0
	var tmpzval *zend.Zval = nil
	host = ParseIpAddress(xparam, &portno)
	if host == nil {
		return -1
	}
	sock.socket = core.PhpNetworkBindSocketToLocalAddr(host, portno, g.Cond(stream.ops == &PhpStreamUdpSocketOps, SOCK_DGRAM, SOCK_STREAM), sockopts, g.CondF1(xparam.GetWantErrortext() != 0, func() *zend.ZendString { return &xparam.outputs.error_text }, nil), &err)
	if host != nil {
		zend._efree(host)
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
	var sockopts long = 1 << 0
	host = ParseIpAddress(xparam, &portno)
	if host == nil {
		return -1
	}
	if (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)) != nil && g.Assign(&tmpzval, PhpStreamContextGetOption((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)), "socket", "bindto")) != nil {
		if tmpzval.u1.v.type_ != 6 {
			if xparam.GetWantErrortext() != 0 {
				xparam.SetErrorText(zend.ZendStrpprintf(0, "local_addr context option is not a string."))
			}
			zend._efree(host)
			return -1
		}
		bindto = ParseIpAddressEx(tmpzval.value.str.val, tmpzval.value.str.len_, &bindport, xparam.GetWantErrortext(), &xparam.outputs.error_text)
	}
	if stream.ops != &PhpStreamUdpSocketOps && (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)) != nil && g.Assign(&tmpzval, PhpStreamContextGetOption((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)), "socket", "tcp_nodelay")) != nil && zend.ZendIsTrue(tmpzval) != 0 {
		sockopts |= 1 << 5
	}

	/* Note: the test here for php_stream_udp_socket_ops is important, because we
	 * want the default to be TCP sockets so that the openssl extension can
	 * re-use this code. */

	sock.socket = core.PhpNetworkConnectSocketToHost(host, portno, g.Cond(stream.ops == &PhpStreamUdpSocketOps, SOCK_DGRAM, SOCK_STREAM), xparam.GetOp() == STREAM_XPORT_OP_CONNECT_ASYNC, xparam.GetTimeout(), g.CondF1(xparam.GetWantErrortext() != 0, func() *zend.ZendString { return &xparam.outputs.error_text }, nil), &err, bindto, bindport, sockopts)
	if sock.socket == -1 {
		ret = -1
	} else {
		ret = 0
	}
	xparam.SetErrorCode(err)
	if host != nil {
		zend._efree(host)
	}
	if bindto != nil {
		zend._efree(bindto)
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
	if nil != (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)) && g.Assign(&tmpzval, PhpStreamContextGetOption((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)), "socket", "tcp_nodelay")) != nil && zend.ZendIsTrue(tmpzval) != 0 {
		nodelay = 1
	}
	clisock = core.PhpNetworkAcceptIncoming(sock.socket, g.CondF1(xparam.GetWantTextaddr() != 0, func() *zend.ZendString { return &xparam.outputs.textaddr }, nil), g.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return &xparam.outputs.addr }, nil), g.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return &xparam.outputs.addrlen }, nil), xparam.GetTimeout(), g.CondF1(xparam.GetWantErrortext() != 0, func() *zend.ZendString { return &xparam.outputs.error_text }, nil), &xparam.outputs.error_code, nodelay)
	if clisock >= 0 {
		var clisockdata *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(zend._emalloc(g.SizeOf("* clisockdata")))
		memcpy(clisockdata, sock, g.SizeOf("* clisockdata"))
		clisockdata.socket = clisock
		xparam.SetClient(_phpStreamAlloc(stream.ops, clisockdata, nil, "r+"))
		if xparam.GetClient() != nil {
			xparam.GetClient().ctx = stream.ctx
			if stream.ctx != nil {
				zend.ZendGcAddref(&(stream.ctx).gc)
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
	case 7:
		xparam = (*PhpStreamXportParam)(ptrparam)
		switch xparam.GetOp() {
		case STREAM_XPORT_OP_CONNECT:

		case STREAM_XPORT_OP_CONNECT_ASYNC:
			xparam.SetReturncode(PhpTcpSockopConnect(stream, sock, xparam))
			return 0
		case STREAM_XPORT_OP_BIND:
			xparam.SetReturncode(PhpTcpSockopBind(stream, sock, xparam))
			return 0
		case STREAM_XPORT_OP_ACCEPT:
			xparam.SetReturncode(PhpTcpSockopAccept(stream, sock, xparam))
			return 0
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
	if g.Cond(persistent_id != nil, 1, 0) {
		sock = zend.__zendMalloc(g.SizeOf("php_netstream_data_t"))
	} else {
		sock = zend._emalloc(g.SizeOf("php_netstream_data_t"))
	}
	memset(sock, 0, g.SizeOf("php_netstream_data_t"))
	sock.is_blocked = 1
	sock.timeout.tv_sec = standard.FileGlobals.default_socket_timeout
	sock.timeout.tv_usec = 0

	/* we don't know the socket until we have determined if we are binding or
	 * connecting */

	sock.socket = -1
	stream = _phpStreamAlloc(ops, sock, persistent_id, "r+")
	if stream == nil {
		g.CondF(g.Cond(persistent_id != nil, 1, 0), func() { return zend.Free(sock) }, func() { return zend._efree(sock) })
		return nil
	}
	if flags == 0 {
		return stream
	}
	return stream
}
