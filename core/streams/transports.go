// <<generate>>

package streams

import (
	"sik/core"
	"sik/ext/standard"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/streams/transports.c>

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

// # include "php_streams_int.h"

// # include "ext/standard/file.h"

var XportHash zend.HashTable

func PhpStreamXportGetHash() *zend.HashTable { return &XportHash }
func PhpStreamXportRegister(protocol string, factory PhpStreamTransportFactory) int {
	var str *zend.ZendString = zend.ZendStringInitInterned(protocol, strlen(protocol), 1)
	zend.ZendHashUpdatePtr(&XportHash, str, factory)
	zend.ZendStringReleaseEx(str, 1)
	return zend.SUCCESS
}
func PhpStreamXportUnregister(protocol *byte) int {
	return zend.ZendHashStrDel(&XportHash, protocol, strlen(protocol))
}

// #define ERR_REPORT(out_err,fmt,arg) if ( out_err ) { * out_err = strpprintf ( 0 , fmt , arg ) ; } else { php_error_docref ( NULL , E_WARNING , fmt , arg ) ; }

// #define ERR_RETURN(out_err,local_err,fmt) if ( out_err ) { * out_err = local_err ; } else { php_error_docref ( NULL , E_WARNING , fmt , local_err ? ZSTR_VAL ( local_err ) : "Unspecified error" ) ; if ( local_err ) { zend_string_release_ex ( local_err , 0 ) ; local_err = NULL ; } }

func _phpStreamXportCreate(name *byte, namelen int, options int, flags int, persistent_id *byte, timeout *__struct__timeval, context *core.PhpStreamContext, error_string **zend.ZendString, error_code *int) *core.PhpStream {
	var stream *core.PhpStream = nil
	var factory PhpStreamTransportFactory = nil
	var p *byte
	var protocol *byte = nil
	var n int = 0
	var failed int = 0
	var error_text *zend.ZendString = nil
	var default_timeout __struct__timeval = __struct__timeval{0, 0}
	default_timeout.tv_sec = standard.FileGlobals.default_socket_timeout
	if timeout == nil {
		timeout = &default_timeout
	}

	/* check for a cached persistent socket */

	if persistent_id != nil {
		switch PhpStreamFromPersistentId(persistent_id, &stream) {
		case 0:

			/* use a 0 second timeout when checking if the socket
			 * has already died */

			if 0 == _phpStreamSetOption(stream, 12, 0, nil) {
				return stream
			}

			/* dead - kill it */

			_phpStreamFree(stream, 1|2|16)
			stream = nil
		case 1:

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
		if nil == g.Assign(&factory, zend.ZendHashStrFindPtr(&XportHash, protocol, n)) {
			var wrapper_name []byte
			if n >= g.SizeOf("wrapper_name") {
				n = g.SizeOf("wrapper_name") - 1
			}
			var php_str_len int
			if n >= g.SizeOf("wrapper_name") {
				php_str_len = g.SizeOf("wrapper_name") - 1
			} else {
				php_str_len = n
			}
			memcpy(wrapper_name, protocol, php_str_len)
			wrapper_name[php_str_len] = '0'
			if error_string != nil {
				*error_string = zend.ZendStrpprintf(0, "Unable to find the socket transport \"%s\" - did you forget to enable it when you configured PHP?", wrapper_name)
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Unable to find the socket transport \"%s\" - did you forget to enable it when you configured PHP?", wrapper_name)
			}
			return nil
		}
	}
	if factory == nil {

		/* should never happen */

		core.PhpErrorDocref(nil, 1<<1, "Could not find a factory !?")
		return nil
	}
	stream = factory(protocol, n, (*byte)(name), namelen, persistent_id, options, flags, timeout, context)
	if stream != nil {
		PhpStreamContextSet(stream, context)
		if (flags & 1) == 0 {

			/* client */

			if (flags & (2 | 16)) != 0 {
				if -1 == PhpStreamXportConnect(stream, name, namelen, g.Cond((flags&16) != 0, 1, 0), timeout, &error_text, error_code) {
					if error_string != nil {
						*error_string = error_text
					} else {
						core.PhpErrorDocref(nil, 1<<1, "connect() failed: %s", g.CondF1(error_text != nil, func() []byte { return error_text.val }, "Unspecified error"))
						if error_text != nil {
							zend.ZendStringReleaseEx(error_text, 0)
							error_text = nil
						}
					}
					failed = 1
				}
			}

			/* client */

		} else {

			/* server */

			if (flags & 4) != 0 {
				if 0 != PhpStreamXportBind(stream, name, namelen, &error_text) {
					if error_string != nil {
						*error_string = error_text
					} else {
						core.PhpErrorDocref(nil, 1<<1, "bind() failed: %s", g.CondF1(error_text != nil, func() []byte { return error_text.val }, "Unspecified error"))
						if error_text != nil {
							zend.ZendStringReleaseEx(error_text, 0)
							error_text = nil
						}
					}
					failed = 1
				} else if (flags & 8) != 0 {
					var zbacklog *zend.Zval = nil
					var backlog int = 32
					if (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)) != nil && g.Assign(&zbacklog, PhpStreamContextGetOption((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)), "socket", "backlog")) != nil {
						backlog = zend.ZvalGetLong(zbacklog)
					}
					if 0 != PhpStreamXportListen(stream, backlog, &error_text) {
						if error_string != nil {
							*error_string = error_text
						} else {
							core.PhpErrorDocref(nil, 1<<1, "listen() failed: %s", g.CondF1(error_text != nil, func() []byte { return error_text.val }, "Unspecified error"))
							if error_text != nil {
								zend.ZendStringReleaseEx(error_text, 0)
								error_text = nil
							}
						}
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
			_phpStreamFree(stream, 1|2|16)
		} else {
			_phpStreamFree(stream, 1|2)
		}
		stream = nil
	}
	return stream
}

/* Bind the stream to a local address */

func PhpStreamXportBind(stream *core.PhpStream, name *byte, namelen int, error_text **zend.ZendString) int {
	var param PhpStreamXportParam
	var ret int
	memset(&param, 0, g.SizeOf("param"))
	param.SetOp(STREAM_XPORT_OP_BIND)
	param.SetName((*byte)(name))
	param.SetNamelen(namelen)
	if error_text != nil {
		param.SetWantErrortext(1)
	} else {
		param.SetWantErrortext(0)
	}
	ret = _phpStreamSetOption(stream, 7, 0, &param)
	if ret == 0 {
		if error_text != nil {
			*error_text = param.GetErrorText()
		}
		return param.GetReturncode()
	}
	return ret
}

/* Connect to a remote address */

func PhpStreamXportConnect(stream *core.PhpStream, name *byte, namelen int, asynchronous int, timeout *__struct__timeval, error_text **zend.ZendString, error_code *int) int {
	var param PhpStreamXportParam
	var ret int
	memset(&param, 0, g.SizeOf("param"))
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
	ret = _phpStreamSetOption(stream, 7, 0, &param)
	if ret == 0 {
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

/* Prepare to listen */

func PhpStreamXportListen(stream *core.PhpStream, backlog int, error_text **zend.ZendString) int {
	var param PhpStreamXportParam
	var ret int
	memset(&param, 0, g.SizeOf("param"))
	param.SetOp(STREAM_XPORT_OP_LISTEN)
	param.SetBacklog(backlog)
	if error_text != nil {
		param.SetWantErrortext(1)
	} else {
		param.SetWantErrortext(0)
	}
	ret = _phpStreamSetOption(stream, 7, 0, &param)
	if ret == 0 {
		if error_text != nil {
			*error_text = param.GetErrorText()
		}
		return param.GetReturncode()
	}
	return ret
}

/* Get the next client and their address (as a string) */

func PhpStreamXportAccept(stream *core.PhpStream, client **core.PhpStream, textaddr **zend.ZendString, addr *any, addrlen *socklen_t, timeout *__struct__timeval, error_text **zend.ZendString) int {
	var param PhpStreamXportParam
	var ret int
	memset(&param, 0, g.SizeOf("param"))
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
	ret = _phpStreamSetOption(stream, 7, 0, &param)
	if ret == 0 {
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
func PhpStreamXportGetName(stream *core.PhpStream, want_peer int, textaddr **zend.ZendString, addr *any, addrlen *socklen_t) int {
	var param PhpStreamXportParam
	var ret int
	memset(&param, 0, g.SizeOf("param"))
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
	ret = _phpStreamSetOption(stream, 7, 0, &param)
	if ret == 0 {
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
	memset(&param, 0, g.SizeOf("param"))
	param.SetOp(STREAM_XPORT_CRYPTO_OP_SETUP)
	param.SetMethod(crypto_method)
	param.SetSession(session_stream)
	ret = _phpStreamSetOption(stream, 8, 0, &param)
	if ret == 0 {
		return param.GetReturncode()
	}
	core.PhpErrorDocref("streams.crypto", 1<<1, "this stream does not support SSL/crypto")
	return ret
}
func PhpStreamXportCryptoEnable(stream *core.PhpStream, activate int) int {
	var param PhpStreamXportCryptoParam
	var ret int
	memset(&param, 0, g.SizeOf("param"))
	param.SetOp(STREAM_XPORT_CRYPTO_OP_ENABLE)
	param.SetActivate(activate)
	ret = _phpStreamSetOption(stream, 8, 0, &param)
	if ret == 0 {
		return param.GetReturncode()
	}
	core.PhpErrorDocref("streams.crypto", 1<<1, "this stream does not support SSL/crypto")
	return ret
}

/* Similar to recv() system call; read data from the stream, optionally
 * peeking, optionally retrieving OOB data */

func PhpStreamXportRecvfrom(stream *core.PhpStream, buf *byte, buflen int, flags int, addr *any, addrlen *socklen_t, textaddr **zend.ZendString) int {
	var param PhpStreamXportParam
	var ret int = 0
	var recvd_len int = 0

	/* otherwise, we are going to bypass the buffer */

	memset(&param, 0, g.SizeOf("param"))
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
	ret = _phpStreamSetOption(stream, 7, 0, &param)
	if ret == 0 {
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

/* Similar to send() system call; send data to the stream, optionally
 * sending it as OOB data */

func PhpStreamXportSendto(stream *core.PhpStream, buf *byte, buflen int, flags int, addr any, addrlen socklen_t) int {
	var param PhpStreamXportParam
	var ret int = 0
	var oob int
	oob = (flags & STREAM_OOB) == STREAM_OOB
	if (oob != 0 || addr) && stream.writefilters.GetHead() != nil {
		core.PhpErrorDocref(nil, 1<<1, "cannot write OOB data, or data to a targeted address on a filtered stream")
		return -1
	}
	memset(&param, 0, g.SizeOf("param"))
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
	ret = _phpStreamSetOption(stream, 7, 0, &param)
	if ret == 0 {
		return param.GetReturncode()
	}
	return -1
}

/* Similar to shutdown() system call; shut down part of a full-duplex
 * connection */

func PhpStreamXportShutdown(stream *core.PhpStream, how StreamShutdownT) int {
	var param PhpStreamXportParam
	var ret int = 0
	memset(&param, 0, g.SizeOf("param"))
	param.SetOp(STREAM_XPORT_OP_SHUTDOWN)
	param.SetHow(how)
	ret = _phpStreamSetOption(stream, 7, 0, &param)
	if ret == 0 {
		return param.GetReturncode()
	}
	return -1
}
