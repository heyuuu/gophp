// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

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

// #define ZEND_SIGNAL_H

// # include < signal . h >

// #define NSIG       65

// #define ZEND_SIGNAL_QUEUE_SIZE       64

/* Signal structs */

// @type ZendSignalEntryT struct

// @type ZendSignalT struct

// @type ZendSignalQueueT struct

/* Signal Globals */

// @type ZendSignalGlobalsT struct

// #define SIGG(v) ( zend_signal_globals . v )

var ZendSignalGlobals ZendSignalGlobalsT

// #define ZEND_SIGNAL_BLOCK_INTERRUPTIONS() SIGG ( depth ) ++ ;

// #define ZEND_SIGNAL_UNBLOCK_INTERRUPTIONS() if ( ( ( SIGG ( depth ) -- ) == SIGG ( blocked ) ) ) { zend_signal_handler_unblock ( ) ; }

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

// #define _GNU_SOURCE

// # include < string . h >

// # include "zend.h"

// # include "zend_globals.h"

// # include < signal . h >

// # include < unistd . h >

// # include "zend_signal.h"

// #define SIGNAL_BEGIN_CRITICAL() sigset_t oldmask ; zend_sigprocmask ( SIG_BLOCK , & global_sigmask , & oldmask ) ;

// #define SIGNAL_END_CRITICAL() zend_sigprocmask ( SIG_SETMASK , & oldmask , NULL ) ;

// #define zend_sigprocmask(signo,set,oldset) sigprocmask ( ( signo ) , ( set ) , ( oldset ) )

// #define TIMEOUT_SIG       SIGPROF

var ZendSigs []int = []int{SIGPROF, SIGHUP, SIGINT, SIGQUIT, SIGTERM, SIGUSR1, SIGUSR2}

// #define SA_FLAGS_MASK       ~ ( SA_NODEFER | SA_RESETHAND )

/* True globals, written only at process startup */

var GlobalOrigHandlers []ZendSignalEntryT
var GlobalSigmask sigset_t

/* {{{ zend_signal_handler_defer
 *  Blocks signals if in critical section */

func ZendSignalHandlerDefer(signo int, siginfo *siginfo_t, context any) {
	var errno_save int = errno
	var queue *ZendSignalQueueT
	var qtmp *ZendSignalQueueT
	if ZendSignalGlobals.GetActive() != 0 {
		if ZendSignalGlobals.GetDepth() == 0 {
			if ZendSignalGlobals.GetBlocked() != 0 {
				ZendSignalGlobals.SetBlocked(0)
			}
			if ZendSignalGlobals.GetRunning() == 0 {
				ZendSignalGlobals.SetRunning(1)
				ZendSignalHandler(signo, siginfo, context)
				queue = ZendSignalGlobals.GetPhead()
				ZendSignalGlobals.SetPhead(nil)
				for queue != nil {
					ZendSignalHandler(queue.GetZendSignal().GetSigno(), queue.GetZendSignal().GetSiginfo(), queue.GetZendSignal().GetContext())
					qtmp = queue.GetNext()
					queue.SetNext(ZendSignalGlobals.GetPavail())
					queue.GetZendSignal().SetSigno(0)
					ZendSignalGlobals.SetPavail(queue)
					queue = qtmp
				}
				ZendSignalGlobals.SetRunning(0)
			}
		} else {
			ZendSignalGlobals.SetBlocked(1)
			if g.Assign(&queue, ZendSignalGlobals.GetPavail()) {
				ZendSignalGlobals.SetPavail(queue.GetNext())
				queue.GetZendSignal().SetSigno(signo)
				queue.GetZendSignal().SetSiginfo(siginfo)
				queue.GetZendSignal().SetContext(context)
				queue.SetNext(nil)
				if ZendSignalGlobals.GetPhead() != nil && ZendSignalGlobals.GetPtail() != nil {
					ZendSignalGlobals.GetPtail().SetNext(queue)
				} else {
					ZendSignalGlobals.SetPhead(queue)
				}
				ZendSignalGlobals.SetPtail(queue)
			}
		}
	} else {

		/* need to just run handler if we're inactive and getting a signal */

		ZendSignalHandler(signo, siginfo, context)

		/* need to just run handler if we're inactive and getting a signal */

	}
	errno = errno_save
}

/* {{{ zend_signal_handler_unblock
 * Handle deferred signal from HANDLE_UNBLOCK_ALARMS */

func ZendSignalHandlerUnblock() {
	var queue *ZendSignalQueueT
	var zend_signal ZendSignalT
	if ZendSignalGlobals.GetActive() != 0 {
		var oldmask sigset_t
		sigprocmask(SIG_BLOCK, &GlobalSigmask, &oldmask)
		queue = ZendSignalGlobals.GetPhead()
		ZendSignalGlobals.SetPhead(queue.GetNext())
		zend_signal = queue.GetZendSignal()
		queue.SetNext(ZendSignalGlobals.GetPavail())
		queue.GetZendSignal().SetSigno(0)
		ZendSignalGlobals.SetPavail(queue)
		ZendSignalHandlerDefer(zend_signal.GetSigno(), zend_signal.GetSiginfo(), zend_signal.GetContext())
		sigprocmask(SIG_SETMASK, &oldmask, nil)
	}
}

/* }}} */

func ZendSignalHandler(signo int, siginfo *siginfo_t, context any) {
	var errno_save int = errno
	var sa __struct__sigaction
	var sigset sigset_t
	var p_sig ZendSignalEntryT
	p_sig = ZendSignalGlobals.GetHandlers()[signo-1]
	if p_sig.GetHandler() == SIG_DFL {
		if sigaction(signo, nil, &sa) == 0 {
			sa.sa_handler = SIG_DFL
			sigemptyset(&sa.sa_mask)
			sigemptyset(&sigset)
			sigaddset(&sigset, signo)
			if sigaction(signo, &sa, nil) == 0 {

				/* throw away any blocked signals */

				sigprocmask(SIG_UNBLOCK, &sigset, nil)
				kill(getpid(), signo)
			}
		}
	} else if p_sig.GetHandler() != SIG_IGN {
		if (p_sig.GetFlags() & SA_SIGINFO) != 0 {
			if (p_sig.GetFlags() & SA_RESETHAND) != 0 {
				ZendSignalGlobals.GetHandlers()[signo-1].SetFlags(0)
				ZendSignalGlobals.GetHandlers()[signo-1].SetHandler(SIG_DFL)
			}
			(*((func(int, *siginfo_t, any))(p_sig.GetHandler())))(signo, siginfo, context)
		} else {
			(*((func(int))(p_sig.GetHandler())))(signo)
		}
	}
	errno = errno_save
}

/* {{{ zend_sigaction
 *  Register a signal handler that will be deferred in critical sections */

func ZendSigaction(signo int, act *__struct__sigaction, oldact *__struct__sigaction) int {
	var sa __struct__sigaction
	var sigset sigset_t
	if oldact != nil {
		oldact.sa_flags = ZendSignalGlobals.GetHandlers()[signo-1].GetFlags()
		oldact.sa_handler = any(ZendSignalGlobals.GetHandlers()[signo-1].GetHandler())
		oldact.sa_mask = GlobalSigmask
	}
	if act != nil {
		ZendSignalGlobals.GetHandlers()[signo-1].SetFlags(act.sa_flags)
		if (act.sa_flags & SA_SIGINFO) != 0 {
			ZendSignalGlobals.GetHandlers()[signo-1].SetHandler(any(act.sa_sigaction))
		} else {
			ZendSignalGlobals.GetHandlers()[signo-1].SetHandler(any(act.sa_handler))
		}
		memset(&sa, 0, g.SizeOf("sa"))
		if ZendSignalGlobals.GetHandlers()[signo-1].GetHandler() == any(SIG_IGN) {
			sa.sa_sigaction = any(SIG_IGN)
		} else {
			sa.sa_flags = SA_SIGINFO | act.sa_flags & ^(SA_NODEFER|SA_RESETHAND)
			sa.sa_sigaction = ZendSignalHandlerDefer
			sa.sa_mask = GlobalSigmask
		}
		if sigaction(signo, &sa, nil) < 0 {
			ZendErrorNoreturn(1<<0, "Error installing signal handler for %d", signo)
		}

		/* unsure this signal is not blocked */

		sigemptyset(&sigset)
		sigaddset(&sigset, signo)
		sigprocmask(SIG_UNBLOCK, &sigset, nil)
	}
	return SUCCESS
}

/* }}} */

func ZendSignal(signo int, handler func(int)) int {
	var sa __struct__sigaction
	memset(&sa, 0, g.SizeOf("sa"))
	sa.sa_flags = 0
	sa.sa_handler = handler
	sa.sa_mask = GlobalSigmask
	return ZendSigaction(signo, &sa, nil)
}

/* }}} */

func ZendSignalRegister(signo int, handler func(int, *siginfo_t, any)) int {
	var sa __struct__sigaction
	if sigaction(signo, nil, &sa) == 0 {
		if (sa.sa_flags&SA_SIGINFO) != 0 && sa.sa_sigaction == handler {
			return FAILURE
		}
		ZendSignalGlobals.GetHandlers()[signo-1].SetFlags(sa.sa_flags)
		if (sa.sa_flags & SA_SIGINFO) != 0 {
			ZendSignalGlobals.GetHandlers()[signo-1].SetHandler(any(sa.sa_sigaction))
		} else {
			ZendSignalGlobals.GetHandlers()[signo-1].SetHandler(any(sa.sa_handler))
		}
		sa.sa_flags = SA_SIGINFO
		sa.sa_sigaction = handler
		sa.sa_mask = GlobalSigmask
		if sigaction(signo, &sa, nil) < 0 {
			ZendErrorNoreturn(1<<0, "Error installing signal handler for %d", signo)
		}
		return SUCCESS
	}
	return FAILURE
}

/* {{{ zend_signal_activate
 *  Install our signal handlers, per request */

func ZendSignalActivate() {
	var x int
	memcpy(&(ZendSignalGlobals.GetHandlers()), &GlobalOrigHandlers, g.SizeOf("global_orig_handlers"))
	if ZendSignalGlobals.GetReset() != 0 {
		for x = 0; x < g.SizeOf("zend_sigs")/g.SizeOf("* zend_sigs"); x++ {
			ZendSignalRegister(ZendSigs[x], ZendSignalHandlerDefer)
		}
	}
	ZendSignalGlobals.SetActive(1)
	ZendSignalGlobals.SetDepth(0)
}

/* {{{ zend_signal_deactivate
 * */

func ZendSignalDeactivate() {
	if ZendSignalGlobals.GetCheck() != 0 {
		var x int
		var sa __struct__sigaction
		if ZendSignalGlobals.GetDepth() != 0 {
			ZendError(1<<5, "zend_signal: shutdown with non-zero blocking depth (%d)", ZendSignalGlobals.GetDepth())
		}

		/* did anyone steal our installed handler */

		for x = 0; x < g.SizeOf("zend_sigs")/g.SizeOf("* zend_sigs"); x++ {
			sigaction(ZendSigs[x], nil, &sa)
			if sa.sa_sigaction != ZendSignalHandlerDefer && sa.sa_sigaction != any(SIG_IGN) {
				ZendError(1<<5, "zend_signal: handler was replaced for signal (%d) after startup", ZendSigs[x])
			}
		}

		/* did anyone steal our installed handler */

	}

	/* After active=0 is set, signal handlers will be called directly and other
	 * state that is reset below will not be accessed. */

	*((*volatile__int)(&(ZendSignalGlobals.GetActive()))) = 0
	ZendSignalGlobals.SetRunning(0)
	ZendSignalGlobals.SetBlocked(0)
	ZendSignalGlobals.SetDepth(0)

	/* If there are any queued signals because of a missed unblock, drop them. */

	if ZendSignalGlobals.GetPhead() != nil && ZendSignalGlobals.GetPtail() != nil {
		ZendSignalGlobals.GetPtail().SetNext(ZendSignalGlobals.GetPavail())
		ZendSignalGlobals.SetPavail(ZendSignalGlobals.GetPhead())
		ZendSignalGlobals.SetPhead(nil)
		ZendSignalGlobals.SetPtail(nil)
	}

	/* If there are any queued signals because of a missed unblock, drop them. */
}

/* }}} */

func ZendSignalGlobalsCtor(zend_signal_globals *ZendSignalGlobalsT) {
	var x int
	memset(zend_signal_globals, 0, g.SizeOf("* zend_signal_globals"))
	zend_signal_globals.SetReset(1)
	for x = 0; x < g.SizeOf("zend_signal_globals -> pstorage")/g.SizeOf("* zend_signal_globals -> pstorage"); x++ {
		var queue *ZendSignalQueueT = &zend_signal_globals.pstorage[x]
		queue.GetZendSignal().SetSigno(0)
		queue.SetNext(zend_signal_globals.GetPavail())
		zend_signal_globals.SetPavail(queue)
	}
}

/* }}} */

func ZendSignalInit() {
	var signo int
	var sa __struct__sigaction

	/* Save previously registered signal handlers into orig_handlers */

	memset(&GlobalOrigHandlers, 0, g.SizeOf("global_orig_handlers"))
	for signo = 1; signo < 65; signo++ {
		if sigaction(signo, nil, &sa) == 0 {
			GlobalOrigHandlers[signo-1].SetFlags(sa.sa_flags)
			if (sa.sa_flags & SA_SIGINFO) != 0 {
				GlobalOrigHandlers[signo-1].SetHandler(any(sa.sa_sigaction))
			} else {
				GlobalOrigHandlers[signo-1].SetHandler(any(sa.sa_handler))
			}
		}
	}
}

/* }}} */

func ZendSignalStartup() {
	ZendSignalGlobalsCtor(&ZendSignalGlobals)

	/* Used to block signals during execution of signal handlers */

	sigfillset(&GlobalSigmask)
	sigdelset(&GlobalSigmask, SIGILL)
	sigdelset(&GlobalSigmask, SIGABRT)
	sigdelset(&GlobalSigmask, SIGFPE)
	sigdelset(&GlobalSigmask, SIGKILL)
	sigdelset(&GlobalSigmask, SIGSEGV)
	sigdelset(&GlobalSigmask, SIGCONT)
	sigdelset(&GlobalSigmask, SIGSTOP)
	sigdelset(&GlobalSigmask, SIGTSTP)
	sigdelset(&GlobalSigmask, SIGTTIN)
	sigdelset(&GlobalSigmask, SIGTTOU)
	ZendSignalInit()
}

/* }}} */
