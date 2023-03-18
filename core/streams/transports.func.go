// <<generate>>

package streams

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	"sik/zend"
	"sik/zend/types"
)

func PhpStreamXportGetHash() *types.HashTable { return &XportHash }
func PhpStreamXportRegister(protocol string, factory PhpStreamTransportFactory) int {
	var str *types.ZendString = types.ZendStringInitInterned(protocol, strlen(protocol), 1)
	zend.ZendHashUpdatePtr(&XportHash, str, factory)
	types.ZendStringReleaseEx(str, 1)
	return types.SUCCESS
}
func PhpStreamXportUnregister(protocol *byte) int {
	return zend.ZendHashStrDel(&XportHash, protocol, strlen(protocol))
}
func ERR_REPORT(out_err **types.ZendString, fmt string, arg []byte) {
	if out_err != nil {
		*out_err = core.Strpprintf(0, fmt, arg)
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, fmt, arg)
	}
}
func ERR_RETURN(out_err **types.ZendString, local_err *types.ZendString, fmt string) {
	if out_err != nil {
		*out_err = local_err
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, fmt, b.CondF1(local_err != nil, func() []byte { return local_err.GetVal() }, "Unspecified error"))
		if local_err != nil {
			types.ZendStringReleaseEx(local_err, 0)
			local_err = nil
		}
	}
}
func _phpStreamXportCreate(
	name *byte,
	namelen int,
	options int,
	flags int,
	persistent_id *byte,
	timeout *__struct__timeval,
	context *core.PhpStreamContext,
	error_string **types.ZendString,
	error_code *int,
) *core.PhpStream {
	var stream *core.PhpStream = nil
	var factory PhpStreamTransportFactory = nil
	var p *byte
	var protocol *byte = nil
	var n int = 0
	var failed int = 0
	var error_text *types.ZendString = nil
	var default_timeout __struct__timeval = __struct__timeval{0, 0}
	default_timeout.tv_sec = standard.FG(default_socket_timeout)
	if timeout == nil {
		timeout = &default_timeout
	}

	/* check for a cached persistent socket */

	if persistent_id != nil {
		switch PhpStreamFromPersistentId(persistent_id, &stream) {
		case core.PHP_STREAM_PERSISTENT_SUCCESS:

			/* use a 0 second timeout when checking if the socket
			 * has already died */

			if core.PHP_STREAM_OPTION_RETURN_OK == core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_CHECK_LIVENESS, 0, nil) {
				return stream
			}

			/* dead - kill it */

			core.PhpStreamPclose(stream)
			stream = nil
			fallthrough
		case core.PHP_STREAM_PERSISTENT_FAILURE:
			fallthrough
		default:

			/* failed; get a new one */

		}
	}
	for p = name; isalnum(int(*p)) || (*p) == '+' || (*p) == '-' || (*p) == '.'; p++ {
		n++
	}
	if (*p) == ':' && n > 1 && !(strncmp("://", p, 3)) {
		protocol = name
		name = p + 3
		namelen -= n + 3
	} else {
		protocol = "tcp"
		n = 3
	}
	if protocol != nil {
		if nil == b.Assign(&factory, zend.ZendHashStrFindPtr(&XportHash, protocol, n)) {
			var wrapper_name []byte
			if n >= b.SizeOf("wrapper_name") {
				n = b.SizeOf("wrapper_name") - 1
			}
			core.PHP_STRLCPY(wrapper_name, protocol, b.SizeOf("wrapper_name"), n)
			ERR_REPORT(error_string, "Unable to find the socket transport \"%s\" - did you forget to enable it when you configured PHP?", wrapper_name)
			return nil
		}
	}
	if factory == nil {

		/* should never happen */

		core.PhpErrorDocref(nil, zend.E_WARNING, "Could not find a factory !?")
		return nil
	}
	stream = factory(protocol, n, (*byte)(name), namelen, persistent_id, options, flags, timeout, context)
	if stream != nil {
		PhpStreamContextSet(stream, context)
		if (flags & STREAM_XPORT_SERVER) == 0 {

			/* client */

			if (flags & (STREAM_XPORT_CONNECT | STREAM_XPORT_CONNECT_ASYNC)) != 0 {
				if -1 == PhpStreamXportConnect(stream, name, namelen, b.Cond((flags&STREAM_XPORT_CONNECT_ASYNC) != 0, 1, 0), timeout, &error_text, error_code) {
					ERR_RETURN(error_string, error_text, "connect() failed: %s")
					failed = 1
				}
			}

			/* client */

		} else {

			/* server */

			if (flags & STREAM_XPORT_BIND) != 0 {
				if 0 != PhpStreamXportBind(stream, name, namelen, &error_text) {
					ERR_RETURN(error_string, error_text, "bind() failed: %s")
					failed = 1
				} else if (flags & STREAM_XPORT_LISTEN) != 0 {
					var zbacklog *types.Zval = nil
					var backlog int = 32
					if core.PHP_STREAM_CONTEXT(stream) != nil && b.Assign(&zbacklog, PhpStreamContextGetOption(core.PHP_STREAM_CONTEXT(stream), "socket", "backlog")) != nil {
						backlog = zend.ZvalGetLong(zbacklog)
					}
					if 0 != PhpStreamXportListen(stream, backlog, &error_text) {
						ERR_RETURN(error_string, error_text, "listen() failed: %s")
						failed = 1
					}
				}
			}

			/* server */

		}
	}
	if failed != 0 {

		/* failure means that they don't get a stream to play with */

		if persistent_id != nil {
			core.PhpStreamPclose(stream)
		} else {
			core.PhpStreamClose(stream)
		}
		stream = nil
	}
	return stream
}
func PhpStreamXportBind(stream *core.PhpStream, name *byte, namelen int, error_text **types.ZendString) int {
	var param PhpStreamXportParam
	var ret int
	memset(&param, 0, b.SizeOf("param"))
	param.SetOp(STREAM_XPORT_OP_BIND)
	param.SetName((*byte)(name))
	param.SetNamelen(namelen)
	if error_text != nil {
		param.SetWantErrortext(1)
	} else {
		param.SetWantErrortext(0)
	}
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_XPORT_API, 0, &param)
	if ret == core.PHP_STREAM_OPTION_RETURN_OK {
		if error_text != nil {
			*error_text = param.GetErrorText()
		}
		return param.GetReturncode()
	}
	return ret
}
func PhpStreamXportConnect(
	stream *core.PhpStream,
	name *byte,
	namelen int,
	asynchronous int,
	timeout *__struct__timeval,
	error_text **types.ZendString,
	error_code *int,
) int {
	var param PhpStreamXportParam
	var ret int
	memset(&param, 0, b.SizeOf("param"))
	if asynchronous != 0 {
		param.SetOp(STREAM_XPORT_OP_CONNECT_ASYNC)
	} else {
		param.SetOp(STREAM_XPORT_OP_CONNECT)
	}
	param.SetName((*byte)(name))
	param.SetNamelen(namelen)
	param.SetTimeout(timeout)
	if error_text != nil {
		param.SetWantErrortext(1)
	} else {
		param.SetWantErrortext(0)
	}
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_XPORT_API, 0, &param)
	if ret == core.PHP_STREAM_OPTION_RETURN_OK {
		if error_text != nil {
			*error_text = param.GetErrorText()
		}
		if error_code != nil {
			*error_code = param.GetErrorCode()
		}
		return param.GetReturncode()
	}
	return ret
}
func PhpStreamXportListen(stream *core.PhpStream, backlog int, error_text **types.ZendString) int {
	var param PhpStreamXportParam
	var ret int
	memset(&param, 0, b.SizeOf("param"))
	param.SetOp(STREAM_XPORT_OP_LISTEN)
	param.SetBacklog(backlog)
	if error_text != nil {
		param.SetWantErrortext(1)
	} else {
		param.SetWantErrortext(0)
	}
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_XPORT_API, 0, &param)
	if ret == core.PHP_STREAM_OPTION_RETURN_OK {
		if error_text != nil {
			*error_text = param.GetErrorText()
		}
		return param.GetReturncode()
	}
	return ret
}
func PhpStreamXportAccept(
	stream *core.PhpStream,
	client **core.PhpStream,
	textaddr **types.ZendString,
	addr *any,
	addrlen *socklen_t,
	timeout *__struct__timeval,
	error_text **types.ZendString,
) int {
	var param PhpStreamXportParam
	var ret int
	memset(&param, 0, b.SizeOf("param"))
	param.SetOp(STREAM_XPORT_OP_ACCEPT)
	param.SetTimeout(timeout)
	if addr != nil {
		param.SetWantAddr(1)
	} else {
		param.SetWantAddr(0)
	}
	if textaddr != nil {
		param.SetWantTextaddr(1)
	} else {
		param.SetWantTextaddr(0)
	}
	if error_text != nil {
		param.SetWantErrortext(1)
	} else {
		param.SetWantErrortext(0)
	}
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_XPORT_API, 0, &param)
	if ret == core.PHP_STREAM_OPTION_RETURN_OK {
		*client = param.GetClient()
		if addr != nil {
			*addr = param.GetOutputsAddr()
			*addrlen = param.GetOutputsAddrlen()
		}
		if textaddr != nil {
			*textaddr = param.GetTextaddr()
		}
		if error_text != nil {
			*error_text = param.GetErrorText()
		}
		return param.GetReturncode()
	}
	return ret
}
func PhpStreamXportGetName(stream *core.PhpStream, want_peer int, textaddr **types.ZendString, addr *any, addrlen *socklen_t) int {
	var param PhpStreamXportParam
	var ret int
	memset(&param, 0, b.SizeOf("param"))
	if want_peer != 0 {
		param.SetOp(STREAM_XPORT_OP_GET_PEER_NAME)
	} else {
		param.SetOp(STREAM_XPORT_OP_GET_NAME)
	}
	if addr != nil {
		param.SetWantAddr(1)
	} else {
		param.SetWantAddr(0)
	}
	if textaddr != nil {
		param.SetWantTextaddr(1)
	} else {
		param.SetWantTextaddr(0)
	}
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_XPORT_API, 0, &param)
	if ret == core.PHP_STREAM_OPTION_RETURN_OK {
		if addr != nil {
			*addr = param.GetOutputsAddr()
			*addrlen = param.GetOutputsAddrlen()
		}
		if textaddr != nil {
			*textaddr = param.GetTextaddr()
		}
		return param.GetReturncode()
	}
	return ret
}
func PhpStreamXportCryptoSetup(stream *core.PhpStream, crypto_method PhpStreamXportCryptMethodT, session_stream *core.PhpStream) int {
	var param PhpStreamXportCryptoParam
	var ret int
	memset(&param, 0, b.SizeOf("param"))
	param.SetOp(STREAM_XPORT_CRYPTO_OP_SETUP)
	param.SetMethod(crypto_method)
	param.SetSession(session_stream)
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_CRYPTO_API, 0, &param)
	if ret == core.PHP_STREAM_OPTION_RETURN_OK {
		return param.GetReturncode()
	}
	core.PhpErrorDocref("streams.crypto", zend.E_WARNING, "this stream does not support SSL/crypto")
	return ret
}
func PhpStreamXportCryptoEnable(stream *core.PhpStream, activate int) int {
	var param PhpStreamXportCryptoParam
	var ret int
	memset(&param, 0, b.SizeOf("param"))
	param.SetOp(STREAM_XPORT_CRYPTO_OP_ENABLE)
	param.SetActivate(activate)
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_CRYPTO_API, 0, &param)
	if ret == core.PHP_STREAM_OPTION_RETURN_OK {
		return param.GetReturncode()
	}
	core.PhpErrorDocref("streams.crypto", zend.E_WARNING, "this stream does not support SSL/crypto")
	return ret
}
func PhpStreamXportRecvfrom(
	stream *core.PhpStream,
	buf *byte,
	buflen int,
	flags int,
	addr *any,
	addrlen *socklen_t,
	textaddr **types.ZendString,
) int {
	var param PhpStreamXportParam
	var ret int = 0
	var recvd_len int = 0

	/* otherwise, we are going to bypass the buffer */

	memset(&param, 0, b.SizeOf("param"))
	param.SetOp(STREAM_XPORT_OP_RECV)
	if addr != nil {
		param.SetWantAddr(1)
	} else {
		param.SetWantAddr(0)
	}
	if textaddr != nil {
		param.SetWantTextaddr(1)
	} else {
		param.SetWantTextaddr(0)
	}
	param.SetBuf(buf)
	param.SetBuflen(buflen)
	param.SetFlags(flags)
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_XPORT_API, 0, &param)
	if ret == core.PHP_STREAM_OPTION_RETURN_OK {
		if addr != nil {
			*addr = param.GetOutputsAddr()
			*addrlen = param.GetOutputsAddrlen()
		}
		if textaddr != nil {
			*textaddr = param.GetTextaddr()
		}
		return recvd_len + param.GetReturncode()
	}
	if recvd_len != 0 {
		return recvd_len
	} else {
		return -1
	}
}
func PhpStreamXportSendto(
	stream *core.PhpStream,
	buf *byte,
	buflen int,
	flags int,
	addr any,
	addrlen socklen_t,
) int {
	var param PhpStreamXportParam
	var ret int = 0
	var oob int
	oob = (flags & STREAM_OOB) == STREAM_OOB
	if (oob != 0 || addr) && stream.GetWritefilters().GetHead() != nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "cannot write OOB data, or data to a targeted address on a filtered stream")
		return -1
	}
	memset(&param, 0, b.SizeOf("param"))
	param.SetOp(STREAM_XPORT_OP_SEND)
	if addr {
		param.SetWantAddr(1)
	} else {
		param.SetWantAddr(0)
	}
	param.SetBuf((*byte)(buf))
	param.SetBuflen(buflen)
	param.SetFlags(flags)
	param.SetInputsAddr(addr)
	param.SetInputsAddrlen(addrlen)
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_XPORT_API, 0, &param)
	if ret == core.PHP_STREAM_OPTION_RETURN_OK {
		return param.GetReturncode()
	}
	return -1
}
func PhpStreamXportShutdown(stream *core.PhpStream, how StreamShutdownT) int {
	var param PhpStreamXportParam
	var ret int = 0
	memset(&param, 0, b.SizeOf("param"))
	param.SetOp(STREAM_XPORT_OP_SHUTDOWN)
	param.SetHow(how)
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_XPORT_API, 0, &param)
	if ret == core.PHP_STREAM_OPTION_RETURN_OK {
		return param.GetReturncode()
	}
	return -1
}
