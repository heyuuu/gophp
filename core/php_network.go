// <<generate>>

package core

// Source: <main/php_network.h>

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
   +----------------------------------------------------------------------+
*/

// #define _PHP_NETWORK_H

// # include < php . h >

const Closesocket = close

// # include < netinet / tcp . h >

const EWOULDBLOCK = EAGAIN

func PhpSocketErrno() __auto__ { return errno }

/* like strerror, but caller must efree the returned string,
 * unless buf is not NULL.
 * Also works sensibly for win32 */

// # include < netinet / in . h >

// # include < sys / socket . h >

/* These are here, rather than with the win32 counterparts above,
 * since <sys/socket.h> defines them. */

const SHUT_RD = 0
const SHUT_WR = 1
const SHUT_RDWR = 2

// # include < sys / time . h >

// # include < stddef . h >

type PhpSocketT = int

const SOCK_RECV_ERR = -1
const STREAM_SOCKOP_NONE long = 1 << 0
const STREAM_SOCKOP_SO_REUSEPORT = 1 << 1
const STREAM_SOCKOP_SO_BROADCAST = 1 << 2
const STREAM_SOCKOP_IPV6_V6ONLY = 1 << 3
const STREAM_SOCKOP_IPV6_V6ONLY_ENABLED = 1 << 4
const STREAM_SOCKOP_TCP_NODELAY = 1 << 5

/* uncomment this to debug poll(2) emulation on systems that have poll(2) */

// # include < poll . h >

type PhpPollfd = __struct__pollfd

const PHP_POLLREADABLE = POLLIN | POLLERR | POLLHUP

func PhpPoll2(ufds *PhpPollfd, nfds int, timeout int) __auto__ { return poll(ufds, nfds, timeout) }

/* timeval-to-timeout (for poll(2)) */

func PhpTvtoto(timeouttv *__struct__timeval) int {
	if timeouttv != nil {
		return timeouttv.tv_sec*1000 + timeouttv.tv_usec/1000
	}
	return -1
}

/* hybrid select(2)/poll(2) for a single descriptor.
 * timeouttv follows same rules as select(2), but is reduced to millisecond accuracy.
 * Returns 0 on timeout, -1 on error, or the event mask (ala poll(2)).
 */

func PhpPollfdFor(fd PhpSocketT, events int, timeouttv *__struct__timeval) int {
	var p PhpPollfd
	var n int
	p.fd = fd
	p.events = events
	p.revents = 0
	n = PhpPoll2(&p, 1, PhpTvtoto(timeouttv))
	if n > 0 {
		return p.revents
	}
	return n
}
func PhpPollfdForMs(fd PhpSocketT, events int, timeout int) int {
	var p PhpPollfd
	var n int
	p.fd = fd
	p.events = events
	p.revents = 0
	n = PhpPoll2(&p, 1, timeout)
	if n > 0 {
		return p.revents
	}
	return n
}

/* emit warning and suggestion for unsafe select(2) usage */

func PHP_SAFE_FD_SET(fd PhpSocketT, set *fd_set) {
	if fd < FD_SETSIZE {
		FD_SET(fd, set)
	}
}
func PHP_SAFE_FD_CLR(fd PhpSocketT, set fd_set) {
	if fd < FD_SETSIZE {
		FD_CLR(fd, set)
	}
}
func PHP_SAFE_FD_ISSET(fd PhpSocketT, set *fd_set) bool {
	return fd < FD_SETSIZE && FD_ISSET(fd, set)
}
func PHP_SAFE_MAX_FD(m PhpSocketT, n int) {
	if m >= FD_SETSIZE {
		_phpEmitFdSetsizeWarning(m)
		m = FD_SETSIZE - 1
	}
}

const PHP_SOCK_CHUNK_SIZE = 8192

type PhpSockaddrStorage = __struct__sockaddr_storage

func PhpConnectNonb(sock PhpSocketT, addr *__struct__sockaddr, addrlen socklen_t, timeout *__struct__timeval) int {
	return PhpNetworkConnectSocket(sock, addr, addrlen, 0, timeout, nil, nil)
}

var PhpStreamSocketOps PhpStreamOps
var PhpStreamGenericSocketOps PhpStreamOps

const PHP_STREAM_IS_SOCKET = &PhpStreamSocketOps

/* open a connection to a host using php_hostconnect and return a stream */

func PhpStreamSockOpenFromSocket(socket PhpSocketT, persistent int) *PhpStream {
	return _phpStreamSockOpenFromSocket(socket, persistent)
}
func PhpStreamSockOpenHost(host *byte, port uint16, socktype int, timeout int, persistent int) *PhpStream {
	return _phpStreamSockOpenHost(host, port, socktype, timeout, persistent)
}

/* {{{ memory debug */

func PhpStreamSockOpenFromSocketRel(socket PhpSocketT, persistent *byte) *PhpStream {
	return _phpStreamSockOpenFromSocket(socket, persistent)
}
func PhpStreamSockOpenHostRel(host *byte, port uint16, socktype int, timeout *__struct__timeval, persistent *byte) *PhpStream {
	return _phpStreamSockOpenHost(host, port, socktype, timeout, persistent)
}
func PhpStreamSockOpenUnixRel(path __auto__, pathlen __auto__, persistent __auto__, timeval __auto__) __auto__ {
	return _php_stream_sock_open_unix(path, pathlen, persistent, timeval)
}

/* }}} */

const MAXFQDNLEN = 255
