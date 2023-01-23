// <<generate>>

package streams

import (
	"sik/core"
)

// Source: <main/streams/php_stream_transport.h>

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

// # include < sys / socket . h >

type PhpStreamTransportFactoryFunc func(proto *byte, protolen int, resourcename *byte, resourcenamelen int, persistent_id *byte, options int, flags int, timeout *__struct__timeval, context *core.PhpStreamContext) *core.PhpStream
type PhpStreamTransportFactory *PhpStreamTransportFactoryFunc

// #define STREAM_XPORT_CLIENT       0

// #define STREAM_XPORT_SERVER       1

// #define STREAM_XPORT_CONNECT       2

// #define STREAM_XPORT_BIND       4

// #define STREAM_XPORT_LISTEN       8

// #define STREAM_XPORT_CONNECT_ASYNC       16

/* Open a client or server socket connection */

// #define php_stream_xport_create(name,namelen,options,flags,persistent_id,timeout,context,estr,ecode) _php_stream_xport_create ( name , namelen , options , flags , persistent_id , timeout , context , estr , ecode STREAMS_CC )

/* Bind the stream to a local address */

/* Connect to a remote address */

/* Prepare to listen */

/* Get the next client and their address as a string, or the underlying address
 * structure.  You must efree either of these if you request them */

/* Get the name of either the socket or it's peer */

type PhpStreamXportSendRecvFlags = int

const (
	STREAM_OOB  = 1
	STREAM_PEEK = 2
)

/* Similar to recv() system call; read data from the stream, optionally
 * peeking, optionally retrieving OOB data */

/* Similar to send() system call; send data to the stream, optionally
 * sending it as OOB data */

type StreamShutdownT = int

const (
	STREAM_SHUT_RD = iota
	STREAM_SHUT_WR
	STREAM_SHUT_RDWR
)

/* Similar to shutdown() system call; shut down part of a full-duplex
 * connection */

/* Structure definition for the set_option interface that the above functions wrap */

/* Because both client and server streams use the same mechanisms
   for encryption we use the LSB to denote clients.
*/

type PhpStreamXportCryptMethodT = int

const (
	STREAM_CRYPTO_METHOD_SSLv2_CLIENT   PhpStreamXportCryptMethodT = 1<<1 | 1
	STREAM_CRYPTO_METHOD_SSLv3_CLIENT   PhpStreamXportCryptMethodT = 1<<2 | 1
	STREAM_CRYPTO_METHOD_SSLv23_CLIENT  PhpStreamXportCryptMethodT = 1<<3 | 1<<4 | 1<<5 | 1
	STREAM_CRYPTO_METHOD_TLSv1_0_CLIENT PhpStreamXportCryptMethodT = 1<<3 | 1
	STREAM_CRYPTO_METHOD_TLSv1_1_CLIENT PhpStreamXportCryptMethodT = 1<<4 | 1
	STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT PhpStreamXportCryptMethodT = 1<<5 | 1
	STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT PhpStreamXportCryptMethodT = 1<<6 | 1
	STREAM_CRYPTO_METHOD_TLS_CLIENT     PhpStreamXportCryptMethodT = 1<<3 | 1<<4 | 1<<5 | 1<<6 | 1
	STREAM_CRYPTO_METHOD_TLS_ANY_CLIENT PhpStreamXportCryptMethodT = 1<<3 | 1<<4 | 1<<5 | 1<<6 | 1
	STREAM_CRYPTO_METHOD_ANY_CLIENT     PhpStreamXportCryptMethodT = 1<<1 | 1<<2 | 1<<3 | 1<<4 | 1<<5 | 1<<6 | 1
	STREAM_CRYPTO_METHOD_SSLv2_SERVER   PhpStreamXportCryptMethodT = 1 << 1
	STREAM_CRYPTO_METHOD_SSLv3_SERVER   PhpStreamXportCryptMethodT = 1 << 2
	STREAM_CRYPTO_METHOD_SSLv23_SERVER  PhpStreamXportCryptMethodT = 1<<3 | 1<<4 | 1<<5 | 1<<6
	STREAM_CRYPTO_METHOD_TLSv1_0_SERVER PhpStreamXportCryptMethodT = 1 << 3
	STREAM_CRYPTO_METHOD_TLSv1_1_SERVER PhpStreamXportCryptMethodT = 1 << 4
	STREAM_CRYPTO_METHOD_TLSv1_2_SERVER PhpStreamXportCryptMethodT = 1 << 5
	STREAM_CRYPTO_METHOD_TLSv1_3_SERVER PhpStreamXportCryptMethodT = 1 << 6
	STREAM_CRYPTO_METHOD_TLS_SERVER     PhpStreamXportCryptMethodT = 1<<3 | 1<<4 | 1<<5 | 1<<6
	STREAM_CRYPTO_METHOD_TLS_ANY_SERVER PhpStreamXportCryptMethodT = 1<<3 | 1<<4 | 1<<5 | 1<<6
	STREAM_CRYPTO_METHOD_ANY_SERVER     PhpStreamXportCryptMethodT = 1<<1 | 1<<2 | 1<<3 | 1<<4 | 1<<5 | 1<<6
)

/* These functions provide crypto support on the underlying transport */
