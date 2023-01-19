// <<generate>>

package streams

import (
	"sik/core"
	"sik/zend"
)

// Source: <main/streams/php_stream_context.h>

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

type PhpStreamNotificationFunc func(context *core.PhpStreamContext, notifycode int, severity int, xmsg *byte, xcode int, bytes_sofar int, bytes_max int, ptr any)

// #define PHP_STREAM_NOTIFIER_PROGRESS       1

/* Attempt to fetch context from the zval passed,
   If no context was passed, use the default context
   The default context has not yet been created, do it now. */

// #define php_stream_context_from_zval(zcontext,nocontext) ( ( zcontext ) ? zend_fetch_resource_ex ( zcontext , "Stream-Context" , php_le_stream_context ( ) ) : ( nocontext ) ? NULL : FG ( default_context ) ? FG ( default_context ) : ( FG ( default_context ) = php_stream_context_alloc ( ) ) )

// #define php_stream_context_to_zval(context,zval) { ZVAL_RES ( zval , ( context ) -> res ) ; GC_ADDREF ( ( context ) -> res ) ; }

// @type PhpStreamNotifier struct
// @type PhpStreamContext struct

/* not all notification codes are implemented */

// #define PHP_STREAM_NOTIFY_RESOLVE       1

// #define PHP_STREAM_NOTIFY_CONNECT       2

// #define PHP_STREAM_NOTIFY_AUTH_REQUIRED       3

// #define PHP_STREAM_NOTIFY_MIME_TYPE_IS       4

// #define PHP_STREAM_NOTIFY_FILE_SIZE_IS       5

// #define PHP_STREAM_NOTIFY_REDIRECTED       6

// #define PHP_STREAM_NOTIFY_PROGRESS       7

// #define PHP_STREAM_NOTIFY_COMPLETED       8

// #define PHP_STREAM_NOTIFY_FAILURE       9

// #define PHP_STREAM_NOTIFY_AUTH_RESULT       10

// #define PHP_STREAM_NOTIFY_SEVERITY_INFO       0

// #define PHP_STREAM_NOTIFY_SEVERITY_WARN       1

// #define PHP_STREAM_NOTIFY_SEVERITY_ERR       2

// #define php_stream_notify_info(context,code,xmsg,xcode) do { if ( ( context ) && ( context ) -> notifier ) { php_stream_notification_notify ( ( context ) , ( code ) , PHP_STREAM_NOTIFY_SEVERITY_INFO , ( xmsg ) , ( xcode ) , 0 , 0 , NULL ) ; } } while ( 0 )

// #define php_stream_notify_progress(context,bsofar,bmax) do { if ( ( context ) && ( context ) -> notifier ) { php_stream_notification_notify ( ( context ) , PHP_STREAM_NOTIFY_PROGRESS , PHP_STREAM_NOTIFY_SEVERITY_INFO , NULL , 0 , ( bsofar ) , ( bmax ) , NULL ) ; } } while ( 0 )

// #define php_stream_notify_progress_init(context,sofar,bmax) do { if ( ( context ) && ( context ) -> notifier ) { ( context ) -> notifier -> progress = ( sofar ) ; ( context ) -> notifier -> progress_max = ( bmax ) ; ( context ) -> notifier -> mask |= PHP_STREAM_NOTIFIER_PROGRESS ; php_stream_notify_progress ( ( context ) , ( sofar ) , ( bmax ) ) ; } } while ( 0 )

// #define php_stream_notify_progress_increment(context,dsofar,dmax) do { if ( ( context ) && ( context ) -> notifier && ( context ) -> notifier -> mask & PHP_STREAM_NOTIFIER_PROGRESS ) { ( context ) -> notifier -> progress += ( dsofar ) ; ( context ) -> notifier -> progress_max += ( dmax ) ; php_stream_notify_progress ( ( context ) , ( context ) -> notifier -> progress , ( context ) -> notifier -> progress_max ) ; } } while ( 0 )

// #define php_stream_notify_file_size(context,file_size,xmsg,xcode) do { if ( ( context ) && ( context ) -> notifier ) { php_stream_notification_notify ( ( context ) , PHP_STREAM_NOTIFY_FILE_SIZE_IS , PHP_STREAM_NOTIFY_SEVERITY_INFO , ( xmsg ) , ( xcode ) , 0 , ( file_size ) , NULL ) ; } } while ( 0 )

// #define php_stream_notify_error(context,code,xmsg,xcode) do { if ( ( context ) && ( context ) -> notifier ) { php_stream_notification_notify ( ( context ) , ( code ) , PHP_STREAM_NOTIFY_SEVERITY_ERR , ( xmsg ) , ( xcode ) , 0 , 0 , NULL ) ; } } while ( 0 )
