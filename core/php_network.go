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

// #define closesocket       close

// # include < netinet / tcp . h >

// #define EWOULDBLOCK       EAGAIN

// #define php_socket_errno() errno

/* like strerror, but caller must efree the returned string,
 * unless buf is not NULL.
 * Also works sensibly for win32 */

// # include < netinet / in . h >

// # include < sys / socket . h >

/* These are here, rather than with the win32 counterparts above,
 * since <sys/socket.h> defines them. */

// #define SHUT_RD       0

// #define SHUT_WR       1

// #define SHUT_RDWR       2

// # include < sys / time . h >

// # include < stddef . h >

type PhpSocketT = int

// #define SOCK_ERR       - 1

// #define SOCK_CONN_ERR       - 1

// #define SOCK_RECV_ERR       - 1

// #define STREAM_SOCKOP_NONE       ( 1 << 0 )

// #define STREAM_SOCKOP_SO_REUSEPORT       ( 1 << 1 )

// #define STREAM_SOCKOP_SO_BROADCAST       ( 1 << 2 )

// #define STREAM_SOCKOP_IPV6_V6ONLY       ( 1 << 3 )

// #define STREAM_SOCKOP_IPV6_V6ONLY_ENABLED       ( 1 << 4 )

// #define STREAM_SOCKOP_TCP_NODELAY       ( 1 << 5 )

/* uncomment this to debug poll(2) emulation on systems that have poll(2) */

// # include < poll . h >

type PhpPollfd = __struct__pollfd

// #define PHP_POLLREADABLE       ( POLLIN | POLLERR | POLLHUP )

// #define php_poll2(ufds,nfds,timeout) poll ( ufds , nfds , timeout )

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
	n = poll(&p, 1, PhpTvtoto(timeouttv))
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
	n = poll(&p, 1, timeout)
	if n > 0 {
		return p.revents
	}
	return n
}

/* emit warning and suggestion for unsafe select(2) usage */

// #define PHP_SAFE_FD_SET(fd,set) do { if ( fd < FD_SETSIZE ) FD_SET ( fd , set ) ; } while ( 0 )

// #define PHP_SAFE_FD_CLR(fd,set) do { if ( fd < FD_SETSIZE ) FD_CLR ( fd , set ) ; } while ( 0 )

// #define PHP_SAFE_FD_ISSET(fd,set) ( ( fd < FD_SETSIZE ) && FD_ISSET ( fd , set ) )

// #define PHP_SAFE_MAX_FD(m,n) do { if ( m >= FD_SETSIZE ) { _php_emit_fd_setsize_warning ( m ) ; m = FD_SETSIZE - 1 ; } } while ( 0 )

// #define PHP_SOCK_CHUNK_SIZE       8192

type PhpSockaddrStorage = __struct__sockaddr_storage

// #define php_connect_nonb(sock,addr,addrlen,timeout) php_network_connect_socket ( ( sock ) , ( addr ) , ( addrlen ) , 0 , ( timeout ) , NULL , NULL )

// @type PhpNetstreamDataT struct

var PhpStreamSocketOps PhpStreamOps
var PhpStreamGenericSocketOps PhpStreamOps

// #define PHP_STREAM_IS_SOCKET       ( & php_stream_socket_ops )

/* open a connection to a host using php_hostconnect and return a stream */

// #define php_stream_sock_open_from_socket(socket,persistent) _php_stream_sock_open_from_socket ( ( socket ) , ( persistent ) STREAMS_CC )

// #define php_stream_sock_open_host(host,port,socktype,timeout,persistent) _php_stream_sock_open_host ( ( host ) , ( port ) , ( socktype ) , ( timeout ) , ( persistent ) STREAMS_CC )

/* {{{ memory debug */

// #define php_stream_sock_open_from_socket_rel(socket,persistent) _php_stream_sock_open_from_socket ( ( socket ) , ( persistent ) STREAMS_REL_CC )

// #define php_stream_sock_open_host_rel(host,port,socktype,timeout,persistent) _php_stream_sock_open_host ( ( host ) , ( port ) , ( socktype ) , ( timeout ) , ( persistent ) STREAMS_REL_CC )

// #define php_stream_sock_open_unix_rel(path,pathlen,persistent,timeval) _php_stream_sock_open_unix ( ( path ) , ( pathlen ) , ( persistent ) , ( timeval ) STREAMS_REL_CC )

/* }}} */

// #define MAXFQDNLEN       255
