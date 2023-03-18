// <<generate>>

package streams

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/zend"
	"sik/zend/types"
)

func XP_SOCK_BUF_SIZE(sz __auto__) __auto__ { return sz }
func PhpSockopWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.GetAbstract())
	var didwrite ssize_t
	var ptimeout *__struct__timeval
	if sock == nil || sock.GetSocket() == -1 {
		return 0
	}
	if sock.GetTimeout().tv_sec == -1 {
		ptimeout = nil
	} else {
		ptimeout = sock.GetTimeout()
	}
retry:
	didwrite = send(sock.GetSocket(), buf, XP_SOCK_BUF_SIZE(count), b.Cond(sock.GetIsBlocked() && ptimeout != nil, MSG_DONTWAIT, 0))
	if didwrite <= 0 {
		var estr *byte
		var err int = core.PhpSocketErrno()
		if err == core.EWOULDBLOCK || err == EAGAIN {
			if sock.GetIsBlocked() {
				var retval int
				sock.SetTimeoutEvent(0)
				for {
					retval = core.PhpPollfdFor(sock.GetSocket(), POLLOUT, ptimeout)
					if retval == 0 {
						sock.SetTimeoutEvent(1)
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
	if sock == nil || sock.GetSocket() == -1 {
		return
	}
	sock.SetTimeoutEvent(0)
	if sock.GetTimeout().tv_sec == -1 {
		ptimeout = nil
	} else {
		ptimeout = sock.GetTimeout()
	}
	for true {
		retval = core.PhpPollfdFor(sock.GetSocket(), core.PHP_POLLREADABLE, ptimeout)
		if retval == 0 {
			sock.SetTimeoutEvent(1)
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
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.GetAbstract())
	var nr_bytes ssize_t = 0
	var err int
	if sock == nil || sock.GetSocket() == -1 {
		return -1
	}
	if sock.GetIsBlocked() {
		PhpSockStreamWaitForData(stream, sock)
		if sock.GetTimeoutEvent() {
			return 0
		}
	}
	nr_bytes = recv(sock.GetSocket(), buf, XP_SOCK_BUF_SIZE(count), b.Cond(sock.GetIsBlocked() && sock.GetTimeout().tv_sec != -1, MSG_DONTWAIT, 0))
	err = core.PhpSocketErrno()
	if nr_bytes < 0 {
		if err == EAGAIN || err == core.EWOULDBLOCK {
			nr_bytes = 0
		} else {
			stream.SetEof(1)
		}
	} else if nr_bytes == 0 {
		stream.SetEof(1)
	}
	if nr_bytes > 0 {
		PhpStreamNotifyProgressIncrement(core.PHP_STREAM_CONTEXT(stream), nr_bytes, 0)
	}
	return nr_bytes
}
func PhpSockopClose(stream *core.PhpStream, close_handle int) int {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.GetAbstract())
	if sock == nil {
		return 0
	}
	if close_handle != 0 {
		if sock.GetSocket() != core.SOCK_ERR {
			core.Closesocket(sock.GetSocket())
			sock.SetSocket(core.SOCK_ERR)
		}
	}
	zend.Pefree(sock, stream.GetIsPersistent())
	return 0
}
func PhpSockopFlush(stream *core.PhpStream) int { return 0 }
func PhpSockopStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.GetAbstract())
	return zend.ZendFstat(sock.GetSocket(), ssb.GetSb())
}
func SockSendto(
	sock *core.PhpNetstreamDataT,
	buf *byte,
	buflen int,
	flags int,
	addr *__struct__sockaddr,
	addrlen socklen_t,
) int {
	var ret int
	if addr != nil {
		ret = sendto(sock.GetSocket(), buf, XP_SOCK_BUF_SIZE(buflen), flags, addr, XP_SOCK_BUF_SIZE(addrlen))
		if ret == core.SOCK_CONN_ERR {
			return -1
		} else {
			return ret
		}
	}
	if b.Assign(&ret, send(sock.GetSocket(), buf, buflen, flags)) == core.SOCK_CONN_ERR {
		return -1
	} else {
		return ret
	}
}
func SockRecvfrom(
	sock *core.PhpNetstreamDataT,
	buf *byte,
	buflen int,
	flags int,
	textaddr **types.ZendString,
	addr **__struct__sockaddr,
	addrlen *socklen_t,
) int {
	var ret int
	var want_addr int = textaddr != nil || addr != nil
	if want_addr != 0 {
		var sa core.PhpSockaddrStorage
		var sl socklen_t = b.SizeOf("sa")
		ret = recvfrom(sock.GetSocket(), buf, XP_SOCK_BUF_SIZE(buflen), flags, (*__struct__sockaddr)(&sa), &sl)
		if ret == core.SOCK_CONN_ERR {
			ret = -1
		} else {
			ret = ret
		}
		if sl {
			core.PhpNetworkPopulateNameFromSockaddr((*__struct__sockaddr)(&sa), sl, textaddr, addr, addrlen)
		} else {
			if textaddr != nil {
				*textaddr = types.ZSTR_EMPTY_ALLOC()
			}
			if addr != nil {
				*addr = nil
				*addrlen = 0
			}
		}
	} else {
		ret = recv(sock.GetSocket(), buf, XP_SOCK_BUF_SIZE(buflen), flags)
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
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.GetAbstract())
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
			if sock.GetTimeout().tv_sec == -1 {
				tv.tv_sec = standard.FG(default_socket_timeout)
				tv.tv_usec = 0
			} else {
				tv = sock.GetTimeout()
			}
		} else {
			tv.tv_sec = value
			tv.tv_usec = 0
		}
		if sock.GetSocket() == -1 {
			alive = 0
		} else if core.PhpPollfdFor(sock.GetSocket(), core.PHP_POLLREADABLE|POLLPRI, &tv) > 0 {
			var ret ssize_t
			var err int
			ret = recv(sock.GetSocket(), &buf, b.SizeOf("buf"), MSG_PEEK)
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
		fallthrough
	case core.PHP_STREAM_OPTION_BLOCKING:
		oldmode = sock.GetIsBlocked()
		if types.SUCCESS == core.PhpSetSockBlocking(sock.GetSocket(), value) {
			sock.SetIsBlocked(value)
			return oldmode
		}
		return core.PHP_STREAM_OPTION_RETURN_ERR
	case core.PHP_STREAM_OPTION_READ_TIMEOUT:
		sock.SetTimeout(*((*__struct__timeval)(ptrparam)))
		sock.SetTimeoutEvent(0)
		return core.PHP_STREAM_OPTION_RETURN_OK
	case core.PHP_STREAM_OPTION_META_DATA_API:
		zend.AddAssocBool((*types.Zval)(ptrparam), "timed_out", sock.GetTimeoutEvent())
		zend.AddAssocBool((*types.Zval)(ptrparam), "blocked", sock.GetIsBlocked())
		zend.AddAssocBool((*types.Zval)(ptrparam), "eof", stream.GetEof())
		return core.PHP_STREAM_OPTION_RETURN_OK
	case core.PHP_STREAM_OPTION_XPORT_API:
		xparam = (*PhpStreamXportParam)(ptrparam)
		switch xparam.GetOp() {
		case STREAM_XPORT_OP_LISTEN:
			if listen(sock.GetSocket(), xparam.GetBacklog()) == 0 {
				xparam.SetReturncode(0)
			} else {
				xparam.SetReturncode(-1)
			}
			return core.PHP_STREAM_OPTION_RETURN_OK
		case STREAM_XPORT_OP_GET_NAME:
			xparam.SetReturncode(core.PhpNetworkGetSockName(sock.GetSocket(), b.CondF1(xparam.GetWantTextaddr() != 0, func() *types.ZendString { return xparam.GetTextaddr() }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return xparam.GetOutputsAddr() }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return xparam.GetOutputsAddrlen() }, nil)))
			return core.PHP_STREAM_OPTION_RETURN_OK
		case STREAM_XPORT_OP_GET_PEER_NAME:
			xparam.SetReturncode(core.PhpNetworkGetPeerName(sock.GetSocket(), b.CondF1(xparam.GetWantTextaddr() != 0, func() *types.ZendString { return xparam.GetTextaddr() }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return xparam.GetOutputsAddr() }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return xparam.GetOutputsAddrlen() }, nil)))
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
			xparam.SetReturncode(SockRecvfrom(sock, xparam.GetBuf(), xparam.GetBuflen(), flags, b.CondF1(xparam.GetWantTextaddr() != 0, func() *types.ZendString { return xparam.GetTextaddr() }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return xparam.GetOutputsAddr() }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return xparam.GetOutputsAddrlen() }, nil)))
			return core.PHP_STREAM_OPTION_RETURN_OK
		case STREAM_XPORT_OP_SHUTDOWN:
			var shutdown_how []int = []int{core.SHUT_RD, core.SHUT_WR, core.SHUT_RDWR}
			xparam.SetReturncode(shutdown(sock.GetSocket(), shutdown_how[xparam.GetHow()]))
			return core.PHP_STREAM_OPTION_RETURN_OK
		default:
			return core.PHP_STREAM_OPTION_RETURN_NOTIMPL
		}
		fallthrough
	default:
		return core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	}
}
func PhpSockopCast(stream *core.PhpStream, castas int, ret *any) int {
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.GetAbstract())
	if sock == nil {
		return types.FAILURE
	}
	switch castas {
	case core.PHP_STREAM_AS_STDIO:
		if ret != nil {
			*((**r.FILE)(ret)) = fdopen(sock.GetSocket(), stream.GetMode())
			if *ret {
				return types.SUCCESS
			}
			return types.FAILURE
		}
		return types.SUCCESS
	case core.PHP_STREAM_AS_FD_FOR_SELECT:
		fallthrough
	case core.PHP_STREAM_AS_FD:
		fallthrough
	case core.PHP_STREAM_AS_SOCKETD:
		if ret != nil {
			*((*core.PhpSocketT)(ret)) = sock.GetSocket()
		}
		return types.SUCCESS
	default:
		return types.FAILURE
	}
}
func ParseIpAddressEx(str *byte, str_len int, portno *int, get_err int, err **types.ZendString) *byte {
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
	return ParseIpAddressEx(xparam.GetName(), xparam.GetNamelen(), portno, xparam.GetWantErrortext(), xparam.GetErrorText())
}
func PhpTcpSockopBind(stream *core.PhpStream, sock *core.PhpNetstreamDataT, xparam *PhpStreamXportParam) int {
	var host *byte = nil
	var portno int
	var err int
	var sockopts long = core.STREAM_SOCKOP_NONE
	var tmpzval *types.Zval = nil
	host = ParseIpAddress(xparam, &portno)
	if host == nil {
		return -1
	}
	sock.SetSocket(core.PhpNetworkBindSocketToLocalAddr(host, portno, b.Cond(stream.GetOps() == &PhpStreamUdpSocketOps, SOCK_DGRAM, SOCK_STREAM), sockopts, b.CondF1(xparam.GetWantErrortext() != 0, func() *types.ZendString { return xparam.GetErrorText() }, nil), &err))
	if host != nil {
		zend.Efree(host)
	}
	if sock.GetSocket() == -1 {
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
	var tmpzval *types.Zval = nil
	var sockopts long = core.STREAM_SOCKOP_NONE
	host = ParseIpAddress(xparam, &portno)
	if host == nil {
		return -1
	}
	if core.PHP_STREAM_CONTEXT(stream) != nil && b.Assign(&tmpzval, PhpStreamContextGetOption(core.PHP_STREAM_CONTEXT(stream), "socket", "bindto")) != nil {
		if tmpzval.GetType() != types.IS_STRING {
			if xparam.GetWantErrortext() != 0 {
				xparam.SetErrorText(core.Strpprintf(0, "local_addr context option is not a string."))
			}
			zend.Efree(host)
			return -1
		}
		bindto = ParseIpAddressEx(tmpzval.GetStr().GetVal(), tmpzval.GetStr().GetLen(), &bindport, xparam.GetWantErrortext(), xparam.GetErrorText())
	}
	if stream.GetOps() != &PhpStreamUdpSocketOps && core.PHP_STREAM_CONTEXT(stream) != nil && b.Assign(&tmpzval, PhpStreamContextGetOption(core.PHP_STREAM_CONTEXT(stream), "socket", "tcp_nodelay")) != nil && zend.ZendIsTrue(tmpzval) != 0 {
		sockopts |= core.STREAM_SOCKOP_TCP_NODELAY
	}

	/* Note: the test here for php_stream_udp_socket_ops is important, because we
	 * want the default to be TCP sockets so that the openssl extension can
	 * re-use this code. */

	sock.SetSocket(core.PhpNetworkConnectSocketToHost(host, portno, b.Cond(stream.GetOps() == &PhpStreamUdpSocketOps, SOCK_DGRAM, SOCK_STREAM), xparam.GetOp() == STREAM_XPORT_OP_CONNECT_ASYNC, xparam.GetTimeout(), b.CondF1(xparam.GetWantErrortext() != 0, func() *types.ZendString { return xparam.GetErrorText() }, nil), &err, bindto, bindport, sockopts))
	if sock.GetSocket() == -1 {
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
	var nodelay types.ZendBool = 0
	var tmpzval *types.Zval = nil
	xparam.SetClient(nil)
	if nil != core.PHP_STREAM_CONTEXT(stream) && b.Assign(&tmpzval, PhpStreamContextGetOption(core.PHP_STREAM_CONTEXT(stream), "socket", "tcp_nodelay")) != nil && zend.ZendIsTrue(tmpzval) != 0 {
		nodelay = 1
	}
	clisock = core.PhpNetworkAcceptIncoming(sock.GetSocket(), b.CondF1(xparam.GetWantTextaddr() != 0, func() *types.ZendString { return xparam.GetTextaddr() }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() *__struct__sockaddr { return xparam.GetOutputsAddr() }, nil), b.CondF1(xparam.GetWantAddr() != 0, func() socklen_t { return xparam.GetOutputsAddrlen() }, nil), xparam.GetTimeout(), b.CondF1(xparam.GetWantErrortext() != 0, func() *types.ZendString { return xparam.GetErrorText() }, nil), xparam.GetErrorCode(), nodelay)
	if clisock >= 0 {
		var clisockdata *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(zend.Emalloc(b.SizeOf("* clisockdata")))
		memcpy(clisockdata, sock, b.SizeOf("* clisockdata"))
		clisockdata.SetSocket(clisock)
		xparam.SetClient(core.PhpStreamAllocRel(stream.GetOps(), clisockdata, nil, "r+"))
		if xparam.GetClient() != nil {
			xparam.GetClient().SetCtx(stream.GetCtx())
			if stream.GetCtx() != nil {
				stream.GetCtx().AddRefcount()
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
	var sock *core.PhpNetstreamDataT = (*core.PhpNetstreamDataT)(stream.GetAbstract())
	var xparam *PhpStreamXportParam
	switch option {
	case core.PHP_STREAM_OPTION_XPORT_API:
		xparam = (*PhpStreamXportParam)(ptrparam)
		switch xparam.GetOp() {
		case STREAM_XPORT_OP_CONNECT:
			fallthrough
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
func PhpStreamGenericSocketFactory(
	proto *byte,
	protolen int,
	resourcename *byte,
	resourcenamelen int,
	persistent_id *byte,
	options int,
	flags int,
	timeout *__struct__timeval,
	context *core.PhpStreamContext,
) *core.PhpStream {
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
	sock.SetIsBlocked(1)
	sock.GetTimeout().tv_sec = standard.FG(default_socket_timeout)
	sock.GetTimeout().tv_usec = 0

	/* we don't know the socket until we have determined if we are binding or
	 * connecting */

	sock.SetSocket(-1)
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
