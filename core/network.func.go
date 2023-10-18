package core

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

func PHP_GAI_STRERROR(x int) __auto__ { return gai_strerror(x) }
func PhpNetworkFreeaddresses(sal **__struct__sockaddr) {
	var sap **__struct__sockaddr
	if sal == nil {
		return
	}
	for sap = sal; (*sap) != nil; sap++ {
		zend.Efree(*sap)
	}
	zend.Efree(sal)
}
func PhpNetworkGetaddresses(host *byte, socktype int, sal ***__struct__sockaddr, error_string **types.String) int {
	var sap **__struct__sockaddr
	var n int
	var ipv6_borked int = -1
	var hints __struct__addrinfo
	var res *__struct__addrinfo
	var sai *__struct__addrinfo
	if host == nil {
		return 0
	}
	memset(&hints, '0', b.SizeOf("hints"))
	hints.ai_family = AF_INET
	hints.ai_socktype = socktype

	var errorString string
	if error_string != nil {
		defer func() {
			if errorString != "" {
				*error_string = types.NewString(errorString)
			} else {
				*error_string = nil
			}
		}()
	}

	/* probe for a working IPv6 stack; even if detected as having v6 at compile
	 * time, at runtime some stacks are slow to resolve or have other issues
	 * if they are not correctly configured.
	 * static variable use is safe here since simple store or fetch operations
	 * are atomic and because the actual probe process is not in danger of
	 * collisions or race conditions. */

	if ipv6_borked == -1 {
		var s int
		s = socket(PF_INET6, SOCK_DGRAM, 0)
		if s == SOCK_ERR {
			ipv6_borked = 1
		} else {
			ipv6_borked = 0
			Closesocket(s)
		}
	}
	if ipv6_borked != 0 {
		hints.ai_family = AF_INET
	} else {
		hints.ai_family = AF_UNSPEC
	}
	if lang.Assign(&n, getaddrinfo(host, nil, &hints, &res)) {
		errorString = fmt.Sprintf("php_network_getaddresses: getaddrinfo failed: %s", PHP_GAI_STRERROR(n))
		PhpErrorDocref("", faults.E_WARNING, errorString)
		return 0
	} else if res == nil {
		if error_string != nil {
			errorString = fmt.Sprintf("php_network_getaddresses: getaddrinfo failed (null result pointer) errno=%d", errno)
		} else {
			errorString = "php_network_getaddresses: getaddrinfo failed (null result pointer)"
		}
		PhpErrorDocref("", faults.E_WARNING, errorString)
		return 0
	}
	sai = res
	for n = 1; lang.Assign(&sai, sai.ai_next) != nil; n++ {

	}
	*sal = zend.SafeEmalloc(n+1, b.SizeOf("* sal"), 0)
	sai = res
	sap = *sal
	for {
		*sap = zend.Emalloc(sai.ai_addrlen)
		memcpy(*sap, sai.ai_addr, sai.ai_addrlen)
		sap++
		if lang.Assign(&sai, sai.ai_next) == nil {
			break
		}
	}
	freeaddrinfo(res)
	*sap = nil
	return n
}
func SET_SOCKET_BLOCKING_MODE(sock PhpSocketT, save PhpNonBlockingFlagsT) {
	save = fcntl(sock, F_GETFL, 0)
	fcntl(sock, F_SETFL, save|O_NONBLOCK)
}
func RESTORE_SOCKET_BLOCKING_MODE(sock PhpSocketT, save PhpNonBlockingFlagsT) __auto__ {
	return fcntl(sock, F_SETFL, save)
}
func PhpNetworkConnectSocket(
	sockfd PhpSocketT,
	addr *__struct__sockaddr,
	addrlen socklen_t,
	asynchronous int,
	timeout *__struct__timeval,
	error_string **types.String,
	error_code *int,
) int {
	var orig_flags PhpNonBlockingFlagsT
	var n int
	var error int = 0
	var len_ socklen_t
	var ret int = 0
	SET_SOCKET_BLOCKING_MODE(sockfd, orig_flags)
	if lang.Assign(&n, connect(sockfd, addr, addrlen)) != 0 {
		error = PhpSocketErrno()
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
	if lang.Assign(&n, PhpPollfdFor(sockfd, PHP_POLLREADABLE|POLLOUT, timeout)) == 0 {
		error = PHP_TIMEOUT_ERROR_VALUE
	}
	if n > 0 {
		len_ = b.SizeOf("error")

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

		RESTORE_SOCKET_BLOCKING_MODE(sockfd, orig_flags)

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
func PhpNetworkBindSocketToLocalAddr(
	host *byte,
	port unsigned,
	socktype int,
	sockopts long,
	error_string **types.String,
	error_code *int,
) PhpSocketT {
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
		if sock == SOCK_ERR {
			continue
		}
		switch sa.sa_family {
		case AF_INET6:
			(*__struct__sockaddr_in6)(sa).sin6_family = sa.sa_family
			(*__struct__sockaddr_in6)(sa).sin6_port = htons(port)
			socklen = b.SizeOf("struct sockaddr_in6")
		case AF_INET:
			(*__struct__sockaddr_in)(sa).sin_family = sa.sa_family
			(*__struct__sockaddr_in)(sa).sin_port = htons(port)
			socklen = b.SizeOf("struct sockaddr_in")
		default:

			/* Unknown family */

			socklen = 0
			sa = nil
		}
		if sa != nil {

			/* attempt to bind */

			n = bind(sock, sa, socklen)
			if n != SOCK_CONN_ERR {
				goto bound
			}
			err = PhpSocketErrno()
		}
		Closesocket(sock)
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
func PhpNetworkParseNetworkAddressWithPort(addr *byte, addrlen zend.ZendLong, sa *__struct__sockaddr, sl *socklen_t) int {
	var colon *byte
	var tmp *byte
	var ret int = types.FAILURE
	var port short
	var in4 *__struct__sockaddr_in = (*__struct__sockaddr_in)(sa)
	var psal **__struct__sockaddr
	var n int
	var errstr *types.String = nil
	var in6 *__struct__sockaddr_in6 = (*__struct__sockaddr_in6)(sa)
	memset(in6, 0, b.SizeOf("struct sockaddr_in6"))
	if (*addr) == '[' {
		colon = memchr(addr+1, ']', addrlen-1)
		if colon == nil || colon[1] != ':' {
			return types.FAILURE
		}
		port = atoi(colon + 2)
		addr++
	} else {
		colon = memchr(addr, ':', addrlen)
		if colon == nil {
			return types.FAILURE
		}
		port = atoi(colon + 1)
	}
	tmp = zend.Estrndup(addr, colon-addr)

	/* first, try interpreting the address as a numeric address */

	if inet_pton(AF_INET6, tmp, in6.sin6_addr) > 0 {
		in6.sin6_port = htons(port)
		in6.sin6_family = AF_INET6
		*sl = b.SizeOf("struct sockaddr_in6")
		ret = types.SUCCESS
		goto out
	}
	if inet_aton(tmp, in4.sin_addr) > 0 {
		in4.sin_port = htons(port)
		in4.sin_family = AF_INET
		*sl = b.SizeOf("struct sockaddr_in")
		ret = types.SUCCESS
		goto out
	}

	/* looks like we'll need to resolve it */

	n = PhpNetworkGetaddresses(tmp, SOCK_DGRAM, &psal, &errstr)
	if n == 0 {
		if errstr != nil {
			PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Failed to resolve `%s': %s", tmp, errstr.GetStr()))
			// types.ZendStringReleaseEx(errstr, 0)
		}
		goto out
	}

	/* copy the details from the first item */

	switch psal.sa_family {
	case AF_INET6:
		*in6 = *(*((**__struct__sockaddr_in6)(psal)))
		in6.sin6_port = htons(port)
		*sl = b.SizeOf("struct sockaddr_in6")
		ret = types.SUCCESS
	case AF_INET:
		*in4 = *(*((**__struct__sockaddr_in)(psal)))
		in4.sin_port = htons(port)
		*sl = b.SizeOf("struct sockaddr_in")
		ret = types.SUCCESS
	}
	PhpNetworkFreeaddresses(psal)
out:
	zend.Efree(tmp)
	return ret
}
func PhpNetworkPopulateNameFromSockaddr(sa *__struct__sockaddr, sl socklen_t, textaddr **types.String, addr **__struct__sockaddr, addrlen *socklen_t) {
	if addr != nil {
		*addr = zend.Emalloc(sl)
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
				*textaddr = types.NewString(fmt.Sprintf("%s:%d", buf, ntohs((*__struct__sockaddr_in)(sa).sin_port)))
			}
		case AF_INET6:
			buf = (*byte)(inet_ntop(sa.sa_family, (*__struct__sockaddr_in6)(sa).sin6_addr, (*byte)(&abuf), b.SizeOf("abuf")))
			if buf != nil {
				*textaddr = types.NewString(fmt.Sprintf("[%s]:%d", buf, ntohs((*__struct__sockaddr_in6)(sa).sin6_port)))
			}
		}
	}
}
func PhpNetworkGetPeerName(sock PhpSocketT, textaddr **types.String, addr **__struct__sockaddr, addrlen *socklen_t) int {
	var sa PhpSockaddrStorage
	var sl socklen_t = b.SizeOf("sa")
	memset(&sa, 0, b.SizeOf("sa"))
	if getpeername(sock, (*__struct__sockaddr)(&sa), &sl) == 0 {
		PhpNetworkPopulateNameFromSockaddr((*__struct__sockaddr)(&sa), sl, textaddr, addr, addrlen)
		return 0
	}
	return -1
}
func PhpNetworkGetSockName(sock PhpSocketT, textaddr **types.String, addr **__struct__sockaddr, addrlen *socklen_t) int {
	var sa PhpSockaddrStorage
	var sl socklen_t = b.SizeOf("sa")
	memset(&sa, 0, b.SizeOf("sa"))
	if getsockname(sock, (*__struct__sockaddr)(&sa), &sl) == 0 {
		PhpNetworkPopulateNameFromSockaddr((*__struct__sockaddr)(&sa), sl, textaddr, addr, addrlen)
		return 0
	}
	return -1
}
func PhpNetworkAcceptIncoming(
	srvsock PhpSocketT,
	textaddr **types.String,
	addr **__struct__sockaddr,
	addrlen *socklen_t,
	timeout *__struct__timeval,
	error_string **types.String,
	error_code *int,
	tcp_nodelay int,
) PhpSocketT {
	var clisock PhpSocketT = -1
	var error int = 0
	var n int
	var sa PhpSockaddrStorage
	var sl socklen_t
	n = PhpPollfdFor(srvsock, PHP_POLLREADABLE, timeout)
	if n == 0 {
		error = PHP_TIMEOUT_ERROR_VALUE
	} else if n == -1 {
		error = PhpSocketErrno()
	} else {
		sl = b.SizeOf("sa")
		clisock = accept(srvsock, (*__struct__sockaddr)(&sa), &sl)
		if clisock != SOCK_ERR {
			PhpNetworkPopulateNameFromSockaddr((*__struct__sockaddr)(&sa), sl, textaddr, addr, addrlen)
			if tcp_nodelay != 0 {

			}
		} else {
			error = PhpSocketErrno()
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
func PhpNetworkConnectSocketToHost(
	host *byte,
	port uint16,
	socktype int,
	asynchronous int,
	timeout *__struct__timeval,
	error_string **types.String,
	error_code *int,
	bindto *byte,
	bindport uint16,
	sockopts long,
) PhpSocketT {
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
		memcpy(&working_timeout, timeout, b.SizeOf("working_timeout"))
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
		if sock == SOCK_ERR {
			continue
		}
		switch sa.sa_family {
		case AF_INET6:
			if bindto == nil || strchr(bindto, ':') {
				(*__struct__sockaddr_in6)(sa).sin6_family = sa.sa_family
				(*__struct__sockaddr_in6)(sa).sin6_port = htons(port)
				socklen = b.SizeOf("struct sockaddr_in6")
			} else {
				socklen = 0
				sa = nil
			}
		case AF_INET:
			(*__struct__sockaddr_in)(sa).sin_family = sa.sa_family
			(*__struct__sockaddr_in)(sa).sin_port = htons(port)
			socklen = b.SizeOf("struct sockaddr_in")
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
					var in4 *__struct__sockaddr_in = zend.Emalloc(b.SizeOf("struct sockaddr_in"))
					local_address = (*__struct__sockaddr)(in4)
					local_address_len = b.SizeOf("struct sockaddr_in")
					in4.sin_family = sa.sa_family
					in4.sin_port = htons(bindport)
					if !(inet_aton(bindto, in4.sin_addr)) {
						PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Invalid IP Address: %s", bindto))
						goto skip_bind
					}
					memset(&(in4.sin_zero), 0, b.SizeOf("in4 -> sin_zero"))
				} else {
					var in6 *__struct__sockaddr_in6 = zend.Emalloc(b.SizeOf("struct sockaddr_in6"))
					local_address = (*__struct__sockaddr)(in6)
					local_address_len = b.SizeOf("struct sockaddr_in6")
					in6.sin6_family = sa.sa_family
					in6.sin6_port = htons(bindport)
					if inet_pton(AF_INET6, bindto, in6.sin6_addr) < 1 {
						PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Invalid IP Address: %s", bindto))
						goto skip_bind
					}
				}
				if local_address == nil || bind(sock, local_address, local_address_len) {
					PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("failed to bind to '%s:%d', system said: %s", bindto, bindport, strerror(errno)))
				}
			skip_bind:
				if local_address != nil {
					zend.Efree(local_address)
				}
			}

			/* free error string received during previous iteration (if any) */

			if error_string != nil && (*error_string) != nil {
				// types.ZendStringReleaseEx(*error_string, 0)
				*error_string = nil
			}
			n = PhpNetworkConnectSocket(sock, sa, socklen, asynchronous, lang.Cond(timeout != nil, &working_timeout, nil), error_string, error_code)
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
		Closesocket(sock)
	}
	sock = -1
connected:
	PhpNetworkFreeaddresses(psal)
	return sock
}
func PhpAnyAddr(family int, addr *PhpSockaddrStorage, port uint16) {
	memset(addr, 0, b.SizeOf("php_sockaddr_storage"))
	switch family {
	case AF_INET6:
		var sin6 *__struct__sockaddr_in6 = (*__struct__sockaddr_in6)(addr)
		sin6.sin6_family = AF_INET6
		sin6.sin6_port = htons(port)
		sin6.sin6_addr = in6addr_any
	case AF_INET:
		var sin *__struct__sockaddr_in = (*__struct__sockaddr_in)(addr)
		sin.sin_family = AF_INET
		sin.sin_port = htons(port)
		sin.sin_addr.s_addr = htonl(INADDR_ANY)
	}
}
func PhpSockaddrSize(addr *PhpSockaddrStorage) int {
	switch (*__struct__sockaddr)(addr).sa_family {
	case AF_INET:
		return b.SizeOf("struct sockaddr_in")
	case AF_INET6:
		return b.SizeOf("struct sockaddr_in6")
	default:
		return 0
	}
}
func PhpSocketStrerror(err long, buf *byte, bufsize int) *byte {
	var errstr *byte
	errstr = strerror(err)
	if buf == nil {
		buf = zend.Estrdup(errstr)
	} else {
		strncpy(buf, errstr, bufsize)
		buf[lang.Cond(bufsize != 0, bufsize-1, 0)] = 0
	}
	return buf
}
func PhpSocketErrorStr(err long) *types.String {
	var errstr *byte
	errstr = strerror(err)
	return types.NewString(errstr)
}
func _phpStreamSockOpenFromSocket(socket PhpSocketT, persistent_id *byte) *PhpStream {
	var stream *PhpStream
	var sock *PhpNetstreamDataT
	sock = zend.Pemalloc(b.SizeOf("php_netstream_data_t"))
	memset(sock, 0, b.SizeOf("php_netstream_data_t"))
	sock.SetIsBlocked(1)
	sock.GetTimeout().tv_sec = standard.FG__().default_socket_timeout
	sock.GetTimeout().tv_usec = 0
	sock.SetSocket(socket)
	stream = PhpStreamAllocRel(&PhpStreamGenericSocketOps, sock, persistent_id, "r+")
	if stream == nil {
		zend.Pefree(sock)
	} else {
		stream.SetIsAvoidBlocking(true)
	}
	return stream
}
func _phpStreamSockOpenHost(host *byte, port uint16, socktype int, timeout *__struct__timeval, persistent_id *byte) *PhpStream {
	res := fmt.Sprintf("tcp://%s:%d", host, port)
	return streams.PhpStreamXportCreate(res, len(res), REPORT_ERRORS, streams.STREAM_XPORT_CLIENT|streams.STREAM_XPORT_CONNECT, persistent_id, timeout, nil, nil, nil)
}
func PhpSetSockBlocking(socketd PhpSocketT, block int) int {
	var ret int = types.SUCCESS
	var myflag int = 0
	var flags int = fcntl(socketd, F_GETFL)
	myflag = O_NONBLOCK
	if block == 0 {
		flags |= myflag
	} else {
		flags &= ^myflag
	}
	if fcntl(socketd, F_SETFL, flags) == -1 {
		ret = types.FAILURE
	}
	return ret
}
func _phpEmitFdSetsizeWarning(max_fd int) {
	PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("You MUST recompile PHP with a larger value of FD_SETSIZE.\nIt is set to %d, but you have descriptors numbered at least as high as %d.\n --enable-fd-setsize=%d is recommended, but you may want to set it\nto equal the maximum number of open files supported by your system,\nin order to avoid seeing this error again at a later date.", FD_SETSIZE, max_fd, max_fd + 1024 & ^1023))
}
func PhpNetworkGethostbyname(name *byte) *__struct__hostent { return gethostbyname(name) }
