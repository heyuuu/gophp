// <<generate>>

package streams

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

/* Open a client or server socket connection */

/* Bind the stream to a local address */

/* Connect to a remote address */

/* Prepare to listen */

/* Get the next client and their address as a string, or the underlying address
 * structure.  You must efree either of these if you request them */

/* Get the name of either the socket or it's peer */

/* Similar to recv() system call; read data from the stream, optionally
 * peeking, optionally retrieving OOB data */

/* Similar to send() system call; send data to the stream, optionally
 * sending it as OOB data */

/* Similar to shutdown() system call; shut down part of a full-duplex
 * connection */

/* Structure definition for the set_option interface that the above functions wrap */

/* Because both client and server streams use the same mechanisms
   for encryption we use the LSB to denote clients.
*/

/* These functions provide crypto support on the underlying transport */
