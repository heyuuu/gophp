// <<generate>>

package core

import (
	"sik/core/streams"
	"sik/ext/standard"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/network.c>

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
   | Author: Stig Venaas <venaas@uninett.no>                              |
   | Streams work by Wez Furlong <wez@thebrainroom.com>                   |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include < stddef . h >

// # include < errno . h >

// # include < sys / param . h >

// # include < sys / types . h >

// # include < sys / socket . h >

// # include < fcntl . h >

// # include < sys / select . h >

// # include < poll . h >

// # include < netinet / in . h >

// # include < netdb . h >

// # include < arpa / inet . h >

// # include "php_network.h"

// # include "ext/standard/file.h"

// #define SOCK_ERR       - 1

// #define SOCK_CONN_ERR       - 1

// #define PHP_TIMEOUT_ERROR_VALUE       ETIMEDOUT

// #define PHP_GAI_STRERROR(x) ( gai_strerror ( x ) )

/* {{{ php_network_freeaddresses
 */

func PhpNetworkFreeaddresses(sal **__struct__sockaddr) {
	var sap **__struct__sockaddr
	if sal == nil {
		return
	}
	for sap = sal; (*sap) != nil; sap++ {
		zend._efree(*sap)
	}
	zend._efree(sal)
}

/* }}} */

func PhpNetworkGetaddresses(host *byte, socktype int, sal ***__struct__sockaddr, error_string **zend.ZendString) int {
	var sap **__struct__sockaddr
	var n int
	var ipv6_borked int = -1
	var hints __struct__addrinfo
	var res *__struct__addrinfo
	var sai *__struct__addrinfo
	if host == nil {
		return 0
	}
	memset(&hints, '0', g.SizeOf("hints"))
	hints.ai_family = AF_INET
	hints.ai_socktype = socktype

	/* probe for a working IPv6 stack; even if detected as having v6 at compile
	 * time, at runtime some stacks are slow to resolve or have other issues
	 * if they are not correctly configured.
	 * static variable use is safe here since simple store or fetch operations
	 * are atomic and because the actual probe process is not in danger of
	 * collisions or race conditions. */

	if ipv6_borked == -1 {
		var s int
		s = socket(PF_INET6, SOCK_DGRAM, 0)
		if s == -1 {
			ipv6_borked = 1
		} else {
			ipv6_borked = 0
			close(s)
		}
	}
	if ipv6_borked != 0 {
		hints.ai_family = AF_INET
	} else {
		hints.ai_family = AF_UNSPEC
	}
	if g.Assign(&n, getaddrinfo(host, nil, &hints, &res)) {
		if error_string != nil {

			/* free error string received during previous iteration (if any) */

			if (*error_string) != nil {
				zend.ZendStringReleaseEx(*error_string, 0)
			}
			*error_string = zend.ZendStrpprintf(0, "php_network_getaddresses: getaddrinfo failed: %s", gai_strerror(n))
			PhpErrorDocref(nil, 1<<1, "%s", (*error_string).val)
		} else {
			PhpErrorDocref(nil, 1<<1, "php_network_getaddresses: getaddrinfo failed: %s", gai_strerror(n))
		}
		return 0
	} else if res == nil {
		if error_string != nil {

			/* free error string received during previous iteration (if any) */

			if (*error_string) != nil {
				zend.ZendStringReleaseEx(*error_string, 0)
			}
			*error_string = zend.ZendStrpprintf(0, "php_network_getaddresses: getaddrinfo failed (null result pointer) errno=%d", errno)
			PhpErrorDocref(nil, 1<<1, "%s", (*error_string).val)
		} else {
			PhpErrorDocref(nil, 1<<1, "php_network_getaddresses: getaddrinfo failed (null result pointer)")
		}
		return 0
	}
	sai = res
	for n = 1; g.Assign(&sai, sai.ai_next) != nil; n++ {

	}
	*sal = zend._safeEmalloc(n+1, g.SizeOf("* sal"), 0)
	sai = res
	sap = *sal
	for {
		*sap = zend._emalloc(sai.ai_addrlen)
		memcpy(*sap, sai.ai_addr, sai.ai_addrlen)
		sap++
		if g.Assign(&sai, sai.ai_next) == nil {
			break
		}
	}
	freeaddrinfo(res)
	*sap = nil
	return n
}

/* }}} */

// #define O_NONBLOCK       O_NDELAY

type PhpNonBlockingFlagsT = int

// #define SET_SOCKET_BLOCKING_MODE(sock,save) save = fcntl ( sock , F_GETFL , 0 ) ; fcntl ( sock , F_SETFL , save | O_NONBLOCK )

// #define RESTORE_SOCKET_BLOCKING_MODE(sock,save) fcntl ( sock , F_SETFL , save )

/* Connect to a socket using an interruptible connect with optional timeout.
 * Optionally, the connect can be made asynchronously, which will implicitly
 * enable non-blocking mode on the socket.
 * */

func PhpNetworkConnectSocket(sockfd PhpSocketT, addr *__struct__sockaddr, addrlen socklen_t, asynchronous int, timeout *__struct__timeval, error_string **zend.ZendString, error_code *int) int {
	var orig_flags PhpNonBlockingFlagsT
	var n int
	var error int = 0
	var len_ socklen_t
	var ret int = 0
	orig_flags = fcntl(sockfd, F_GETFL, 0)
	fcntl(sockfd, F_SETFL, orig_flags|O_NDELAY)
	if g.Assign(&n, connect(sockfd, addr, addrlen)) != 0 {
		error = errno
		if error_code != nil {
			*error_code = error
		}
		if error != EINPROGRESS {
			if error_string != nil {
				*error_string = PhpSocketErrorStr(error)
			}
			return -1
		}
		if asynchronous != 0 && error == EINPROGRESS {

			/* this is fine by us */

			return 0

			/* this is fine by us */

		}
	}
	if n == 0 {
		goto ok
	}
	if g.Assign(&n, PhpPollfdFor(sockfd, POLLIN|POLLERR|POLLHUP|POLLOUT, timeout)) == 0 {
		error = ETIMEDOUT
	}
	if n > 0 {
		len_ = g.SizeOf("error")

		/*
		   BSD-derived systems set errno correctly
		   Solaris returns -1 from getsockopt in case of error
		*/

		if getsockopt(sockfd, SOL_SOCKET, SO_ERROR, (*byte)(&error), &len_) != 0 {
			ret = -1
		}

		/*
		   BSD-derived systems set errno correctly
		   Solaris returns -1 from getsockopt in case of error
		*/

	} else {

		/* whoops: sockfd has disappeared */

		ret = -1

		/* whoops: sockfd has disappeared */

	}
ok:
	if asynchronous == 0 {

		/* back to blocking mode */

		fcntl(sockfd, F_SETFL, orig_flags)

		/* back to blocking mode */

	}
	if error_code != nil {
		*error_code = error
	}
	if error != 0 {
		ret = -1
		if error_string != nil {
			*error_string = PhpSocketErrorStr(error)
		}
	}
	return ret
}

/* }}} */

func SubTimes(a __struct__timeval, b __struct__timeval, result *__struct__timeval) {
	result.tv_usec = a.tv_usec - b.tv_usec
	if result.tv_usec < 0 {
		a.tv_sec--
		result.tv_usec += 1000000
	}
	result.tv_sec = a.tv_sec - b.tv_sec
	if result.tv_sec < 0 {
		result.tv_sec++
		result.tv_usec -= 1000000
	}
}

/* }}} */

func PhpNetworkBindSocketToLocalAddr(host *byte, port unsigned, socktype int, sockopts long, error_string **zend.ZendString, error_code *int) PhpSocketT {
	var num_addrs int
	var n int
	var err int = 0
	var sock PhpSocketT
	var sal **__struct__sockaddr
	var psal ***__struct__sockaddr
	var sa **__struct__sockaddr
	var socklen socklen_t
	var sockoptval int = 1
	num_addrs = PhpNetworkGetaddresses(host, socktype, &psal, error_string)
	if num_addrs == 0 {

		/* could not resolve address(es) */

		return -1

		/* could not resolve address(es) */

	}
	for sal = psal; (*sal) != nil; sal++ {
		sa = *sal

		/* create a socket for this address */

		sock = socket(sa.sa_family, socktype, 0)
		if sock == -1 {
			continue
		}
		switch sa.sa_family {
		case AF_INET6:
			(*__struct__sockaddr_in6)(sa).sin6_family = sa.sa_family
			(*__struct__sockaddr_in6)(sa).sin6_port = htons(port)
			socklen = g.SizeOf("struct sockaddr_in6")
			break
		case AF_INET:
			(*__struct__sockaddr_in)(sa).sin_family = sa.sa_family
			(*__struct__sockaddr_in)(sa).sin_port = htons(port)
			socklen = g.SizeOf("struct sockaddr_in")
			break
		default:

			/* Unknown family */

			socklen = 0
			sa = nil
		}
		if sa != nil {

			/* attempt to bind */

			n = bind(sock, sa, socklen)
			if n != -1 {
				goto bound
			}
			err = errno
		}
		close(sock)
	}
	sock = -1
	if error_code != nil {
		*error_code = err
	}
	if error_string != nil {
		*error_string = PhpSocketErrorStr(err)
	}
bound:
	PhpNetworkFreeaddresses(psal)
	return sock
}

/* }}} */

func PhpNetworkParseNetworkAddressWithPort(addr *byte, addrlen zend.ZendLong, sa *__struct__sockaddr, sl *socklen_t) int {
	var colon *byte
	var tmp *byte
	var ret int = zend.FAILURE
	var port short
	var in4 *__struct__sockaddr_in = (*__struct__sockaddr_in)(sa)
	var psal **__struct__sockaddr
	var n int
	var errstr *zend.ZendString = nil
	var in6 *__struct__sockaddr_in6 = (*__struct__sockaddr_in6)(sa)
	memset(in6, 0, g.SizeOf("struct sockaddr_in6"))
	if (*addr) == '[' {
		colon = memchr(addr+1, ']', addrlen-1)
		if colon == nil || colon[1] != ':' {
			return zend.FAILURE
		}
		port = atoi(colon + 2)
		addr++
	} else {
		colon = memchr(addr, ':', addrlen)
		if colon == nil {
			return zend.FAILURE
		}
		port = atoi(colon + 1)
	}
	tmp = zend._estrndup(addr, colon-addr)

	/* first, try interpreting the address as a numeric address */

	if inet_pton(AF_INET6, tmp, &in6.sin6_addr) > 0 {
		in6.sin6_port = htons(port)
		in6.sin6_family = AF_INET6
		*sl = g.SizeOf("struct sockaddr_in6")
		ret = zend.SUCCESS
		goto out
	}
	if inet_aton(tmp, &in4.sin_addr) > 0 {
		in4.sin_port = htons(port)
		in4.sin_family = AF_INET
		*sl = g.SizeOf("struct sockaddr_in")
		ret = zend.SUCCESS
		goto out
	}

	/* looks like we'll need to resolve it */

	n = PhpNetworkGetaddresses(tmp, SOCK_DGRAM, &psal, &errstr)
	if n == 0 {
		if errstr != nil {
			PhpErrorDocref(nil, 1<<1, "Failed to resolve `%s': %s", tmp, errstr.val)
			zend.ZendStringReleaseEx(errstr, 0)
		}
		goto out
	}

	/* copy the details from the first item */

	switch (*psal).sa_family {
	case AF_INET6:
		*in6 = *(*((**__struct__sockaddr_in6)(psal)))
		in6.sin6_port = htons(port)
		*sl = g.SizeOf("struct sockaddr_in6")
		ret = zend.SUCCESS
		break
	case AF_INET:
		*in4 = *(*((**__struct__sockaddr_in)(psal)))
		in4.sin_port = htons(port)
		*sl = g.SizeOf("struct sockaddr_in")
		ret = zend.SUCCESS
		break
	}
	PhpNetworkFreeaddresses(psal)
out:
	zend._efree(tmp)
	return ret
}
func PhpNetworkPopulateNameFromSockaddr(sa *__struct__sockaddr, sl socklen_t, textaddr **zend.ZendString, addr **__struct__sockaddr, addrlen *socklen_t) {
	if addr != nil {
		*addr = zend._emalloc(sl)
		memcpy(*addr, sa, sl)
		*addrlen = sl
	}
	if textaddr != nil {
		var abuf []byte
		var buf *byte = nil
		switch sa.sa_family {
		case AF_INET:

			/* generally not thread safe, but it *is* thread safe under win32 */

			buf = inet_ntoa((*__struct__sockaddr_in)(sa).sin_addr)
			if buf != nil {
				*textaddr = zend.ZendStrpprintf(0, "%s:%d", buf, ntohs((*__struct__sockaddr_in)(sa).sin_port))
			}
			break
		case AF_INET6:
			buf = (*byte)(inet_ntop(sa.sa_family, &((*__struct__sockaddr_in6)(sa)).sin6_addr, (*byte)(&abuf), g.SizeOf("abuf")))
			if buf != nil {
				*textaddr = zend.ZendStrpprintf(0, "[%s]:%d", buf, ntohs((*__struct__sockaddr_in6)(sa).sin6_port))
			}
			break
		}
	}
}
func PhpNetworkGetPeerName(sock PhpSocketT, textaddr **zend.ZendString, addr **__struct__sockaddr, addrlen *socklen_t) int {
	var sa PhpSockaddrStorage
	var sl socklen_t = g.SizeOf("sa")
	memset(&sa, 0, g.SizeOf("sa"))
	if getpeername(sock, (*__struct__sockaddr)(&sa), &sl) == 0 {
		PhpNetworkPopulateNameFromSockaddr((*__struct__sockaddr)(&sa), sl, textaddr, addr, addrlen)
		return 0
	}
	return -1
}
func PhpNetworkGetSockName(sock PhpSocketT, textaddr **zend.ZendString, addr **__struct__sockaddr, addrlen *socklen_t) int {
	var sa PhpSockaddrStorage
	var sl socklen_t = g.SizeOf("sa")
	memset(&sa, 0, g.SizeOf("sa"))
	if getsockname(sock, (*__struct__sockaddr)(&sa), &sl) == 0 {
		PhpNetworkPopulateNameFromSockaddr((*__struct__sockaddr)(&sa), sl, textaddr, addr, addrlen)
		return 0
	}
	return -1
}

/* Accept a client connection from a server socket,
 * using an optional timeout.
 * Returns the peer address in addr/addrlen (it will emalloc
 * these, so be sure to efree the result).
 * If you specify textaddr, a text-printable
 * version of the address will be emalloc'd and returned.
 * */

func PhpNetworkAcceptIncoming(srvsock PhpSocketT, textaddr **zend.ZendString, addr **__struct__sockaddr, addrlen *socklen_t, timeout *__struct__timeval, error_string **zend.ZendString, error_code *int, tcp_nodelay int) PhpSocketT {
	var clisock PhpSocketT = -1
	var error int = 0
	var n int
	var sa PhpSockaddrStorage
	var sl socklen_t
	n = PhpPollfdFor(srvsock, POLLIN|POLLERR|POLLHUP, timeout)
	if n == 0 {
		error = ETIMEDOUT
	} else if n == -1 {
		error = errno
	} else {
		sl = g.SizeOf("sa")
		clisock = accept(srvsock, (*__struct__sockaddr)(&sa), &sl)
		if clisock != -1 {
			PhpNetworkPopulateNameFromSockaddr((*__struct__sockaddr)(&sa), sl, textaddr, addr, addrlen)
			if tcp_nodelay != 0 {

			}
		} else {
			error = errno
		}
	}
	if error_code != nil {
		*error_code = error
	}
	if error_string != nil {
		*error_string = PhpSocketErrorStr(error)
	}
	return clisock
}

/* }}} */

func PhpNetworkConnectSocketToHost(host *byte, port uint16, socktype int, asynchronous int, timeout *__struct__timeval, error_string **zend.ZendString, error_code *int, bindto *byte, bindport uint16, sockopts long) PhpSocketT {
	var num_addrs int
	var n int
	var fatal int = 0
	var sock PhpSocketT
	var sal **__struct__sockaddr
	var psal ***__struct__sockaddr
	var sa **__struct__sockaddr
	var working_timeout __struct__timeval
	var socklen socklen_t
	var limit_time __struct__timeval
	var time_now __struct__timeval
	num_addrs = PhpNetworkGetaddresses(host, socktype, &psal, error_string)
	if num_addrs == 0 {

		/* could not resolve address(es) */

		return -1

		/* could not resolve address(es) */

	}
	if timeout != nil {
		memcpy(&working_timeout, timeout, g.SizeOf("working_timeout"))
		gettimeofday(&limit_time, nil)
		limit_time.tv_sec += working_timeout.tv_sec
		limit_time.tv_usec += working_timeout.tv_usec
		if limit_time.tv_usec >= 1000000 {
			limit_time.tv_usec -= 1000000
			limit_time.tv_sec++
		}
	}
	for sal = psal; fatal == 0 && (*sal) != nil; sal++ {
		sa = *sal

		/* create a socket for this address */

		sock = socket(sa.sa_family, socktype, 0)
		if sock == -1 {
			continue
		}
		switch sa.sa_family {
		case AF_INET6:
			if bindto == nil || strchr(bindto, ':') {
				(*__struct__sockaddr_in6)(sa).sin6_family = sa.sa_family
				(*__struct__sockaddr_in6)(sa).sin6_port = htons(port)
				socklen = g.SizeOf("struct sockaddr_in6")
			} else {
				socklen = 0
				sa = nil
			}
			break
		case AF_INET:
			(*__struct__sockaddr_in)(sa).sin_family = sa.sa_family
			(*__struct__sockaddr_in)(sa).sin_port = htons(port)
			socklen = g.SizeOf("struct sockaddr_in")
			break
		default:

			/* Unknown family */

			socklen = 0
			sa = nil
		}
		if sa != nil {

			/* make a connection attempt */

			if bindto != nil {
				var local_address *__struct__sockaddr = nil
				var local_address_len int = 0
				if sa.sa_family == AF_INET {
					if strchr(bindto, ':') {
						goto skip_bind
					}
					var in4 *__struct__sockaddr_in = zend._emalloc(g.SizeOf("struct sockaddr_in"))
					local_address = (*__struct__sockaddr)(in4)
					local_address_len = g.SizeOf("struct sockaddr_in")
					in4.sin_family = sa.sa_family
					in4.sin_port = htons(bindport)
					if !(inet_aton(bindto, &in4.sin_addr)) {
						PhpErrorDocref(nil, 1<<1, "Invalid IP Address: %s", bindto)
						goto skip_bind
					}
					memset(&(in4.sin_zero), 0, g.SizeOf("in4 -> sin_zero"))
				} else {
					var in6 *__struct__sockaddr_in6 = zend._emalloc(g.SizeOf("struct sockaddr_in6"))
					local_address = (*__struct__sockaddr)(in6)
					local_address_len = g.SizeOf("struct sockaddr_in6")
					in6.sin6_family = sa.sa_family
					in6.sin6_port = htons(bindport)
					if inet_pton(AF_INET6, bindto, &in6.sin6_addr) < 1 {
						PhpErrorDocref(nil, 1<<1, "Invalid IP Address: %s", bindto)
						goto skip_bind
					}
				}
				if local_address == nil || bind(sock, local_address, local_address_len) {
					PhpErrorDocref(nil, 1<<1, "failed to bind to '%s:%d', system said: %s", bindto, bindport, strerror(errno))
				}
			skip_bind:
				if local_address != nil {
					zend._efree(local_address)
				}
			}

			/* free error string received during previous iteration (if any) */

			if error_string != nil && (*error_string) != nil {
				zend.ZendStringReleaseEx(*error_string, 0)
				*error_string = nil
			}
			n = PhpNetworkConnectSocket(sock, sa, socklen, asynchronous, g.Cond(timeout != nil, &working_timeout, nil), error_string, error_code)
			if n != -1 {
				goto connected
			}

			/* adjust timeout for next attempt */

			if timeout != nil {
				gettimeofday(&time_now, nil)
				if !(__timercmp(&time_now, &limit_time, "<")) {

					/* time limit expired; don't attempt any further connections */

					fatal = 1

					/* time limit expired; don't attempt any further connections */

				} else {

					/* work out remaining time */

					SubTimes(limit_time, time_now, &working_timeout)

					/* work out remaining time */

				}
			}

			/* adjust timeout for next attempt */

		}
		close(sock)
	}
	sock = -1
connected:
	PhpNetworkFreeaddresses(psal)
	return sock
}

/* }}} */

func PhpAnyAddr(family int, addr *PhpSockaddrStorage, port uint16) {
	memset(addr, 0, g.SizeOf("php_sockaddr_storage"))
	switch family {
	case AF_INET6:
		var sin6 *__struct__sockaddr_in6 = (*__struct__sockaddr_in6)(addr)
		sin6.sin6_family = AF_INET6
		sin6.sin6_port = htons(port)
		sin6.sin6_addr = in6addr_any
		break
	case AF_INET:
		var sin *__struct__sockaddr_in = (*__struct__sockaddr_in)(addr)
		sin.sin_family = AF_INET
		sin.sin_port = htons(port)
		sin.sin_addr.s_addr = htonl(INADDR_ANY)
		break
	}
}

/* }}} */

func PhpSockaddrSize(addr *PhpSockaddrStorage) int {
	switch (*__struct__sockaddr)(addr).sa_family {
	case AF_INET:
		return g.SizeOf("struct sockaddr_in")
	case AF_INET6:
		return g.SizeOf("struct sockaddr_in6")
	default:
		return 0
	}
}

/* }}} */

func PhpSocketStrerror(err long, buf *byte, bufsize int) *byte {
	var errstr *byte
	errstr = strerror(err)
	if buf == nil {
		buf = zend._estrdup(errstr)
	} else {
		strncpy(buf, errstr, bufsize)
		buf[g.Cond(bufsize != 0, bufsize-1, 0)] = 0
	}
	return buf
}

/* }}} */

func PhpSocketErrorStr(err long) *zend.ZendString {
	var errstr *byte
	errstr = strerror(err)
	return zend.ZendStringInit(errstr, strlen(errstr), 0)
}

/* }}} */

func _phpStreamSockOpenFromSocket(socket PhpSocketT, persistent_id *byte) *PhpStream {
	var stream *PhpStream
	var sock *PhpNetstreamDataT
	if g.Cond(persistent_id != nil, 1, 0) {
		sock = zend.__zendMalloc(g.SizeOf("php_netstream_data_t"))
	} else {
		sock = zend._emalloc(g.SizeOf("php_netstream_data_t"))
	}
	memset(sock, 0, g.SizeOf("php_netstream_data_t"))
	sock.SetIsBlocked(1)
	sock.timeout.tv_sec = standard.FileGlobals.default_socket_timeout
	sock.timeout.tv_usec = 0
	sock.SetSocket(socket)
	stream = _phpStreamAlloc(&PhpStreamGenericSocketOps, sock, persistent_id, "r+")
	if stream == nil {
		g.CondF(g.Cond(persistent_id != nil, 1, 0), func() { return zend.Free(sock) }, func() { return zend._efree(sock) })
	} else {
		stream.SetFlags(stream.GetFlags() | 0x10)
	}
	return stream
}
func _phpStreamSockOpenHost(host *byte, port uint16, socktype int, timeout *__struct__timeval, persistent_id *byte) *PhpStream {
	var res *byte
	var reslen zend.ZendLong
	var stream *PhpStream
	reslen = zend.ZendSpprintf(&res, 0, "tcp://%s:%d", host, port)
	stream = streams._phpStreamXportCreate(res, reslen, 0x8, 0|2, persistent_id, timeout, nil, nil, nil)
	zend._efree(res)
	return stream
}
func PhpSetSockBlocking(socketd PhpSocketT, block int) int {
	var ret int = zend.SUCCESS
	var myflag int = 0
	var flags int = fcntl(socketd, F_GETFL)
	myflag = O_NDELAY
	if block == 0 {
		flags |= myflag
	} else {
		flags &= ^myflag
	}
	if fcntl(socketd, F_SETFL, flags) == -1 {
		ret = zend.FAILURE
	}
	return ret
}
func _phpEmitFdSetsizeWarning(max_fd int) {
	PhpErrorDocref(nil, 1<<1, "You MUST recompile PHP with a larger value of FD_SETSIZE.\n"+"It is set to %d, but you have descriptors numbered at least as high as %d.\n"+" --enable-fd-setsize=%d is recommended, but you may want to set it\n"+"to equal the maximum number of open files supported by your system,\n"+"in order to avoid seeing this error again at a later date.", FD_SETSIZE, max_fd, max_fd + 1024 & ^1023)
}
func PhpNetworkGethostbyname(name *byte) *__struct__hostent { return gethostbyname(name) }
