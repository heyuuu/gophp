// <<generate>>

package zend

// Source: <Zend/zend_signal.h>

/*
  +----------------------------------------------------------------------+
  | Zend Signal Handling                                                 |
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
  | Authors: Lucas Nealan <lucas@php.net>                                |
  |          Arnaud Le Blanc <lbarnaud@php.net>                          |
  +----------------------------------------------------------------------+

*/

const NSIG = 65
const ZEND_SIGNAL_QUEUE_SIZE = 64

/* Signal structs */

/* Signal Globals */

var ZendSignalGlobals ZendSignalGlobalsT

// Source: <Zend/zend_signal.c>

/*
  +----------------------------------------------------------------------+
  | Zend Signal Handling                                                 |
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
  | Authors: Lucas Nealan <lucas@php.net>                                |
  |          Arnaud Le Blanc <lbarnaud@php.net>                          |
  +----------------------------------------------------------------------+

   This software was contributed to PHP by Facebook Inc. in 2008.

   Future revisions and derivatives of this source code must acknowledge
   Facebook Inc. as the original contributor of this module by leaving
   this note intact in the source code.

   All other licensing and usage conditions are those of the PHP Group.
*/

const TIMEOUT_SIG = SIGPROF

var ZendSigs []int = []int{TIMEOUT_SIG, SIGHUP, SIGINT, SIGQUIT, SIGTERM, SIGUSR1, SIGUSR2}

const SA_FLAGS_MASK = ^(SA_NODEFER | SA_RESETHAND)

/* True globals, written only at process startup */

var GlobalOrigHandlers []ZendSignalEntryT
var GlobalSigmask sigset_t

/* {{{ zend_signal_handler_defer
 *  Blocks signals if in critical section */

/* {{{ zend_signal_handler_unblock
 * Handle deferred signal from HANDLE_UNBLOCK_ALARMS */

/* }}} */

/* {{{ zend_sigaction
 *  Register a signal handler that will be deferred in critical sections */

/* }}} */

/* }}} */

/* {{{ zend_signal_activate
 *  Install our signal handlers, per request */

/* {{{ zend_signal_deactivate
 * */

/* }}} */

/* }}} */

/* }}} */

/* }}} */
