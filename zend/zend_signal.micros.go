// <<generate>>

package zend

// #define ZEND_SIGNAL_H

// # include < signal . h >

// #define ZEND_SIGNAL_BLOCK_INTERRUPTIONS() SIGG ( depth ) ++ ;

// #define ZEND_SIGNAL_UNBLOCK_INTERRUPTIONS() if ( ( ( SIGG ( depth ) -- ) == SIGG ( blocked ) ) ) { zend_signal_handler_unblock ( ) ; }

// #define _GNU_SOURCE

// # include < string . h >

// # include "zend.h"

// # include "zend_globals.h"

// # include < signal . h >

// # include < unistd . h >

// # include "zend_signal.h"
