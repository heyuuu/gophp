// <<generate>>

package zend

import (
	b "sik/builtin"
)

func SIGG(v __auto__) __auto__ { return ZendSignalGlobals.v }
func SIGNAL_BEGIN_CRITICAL() {
	var oldmask sigset_t
	ZendSigprocmask(SIG_BLOCK, &GlobalSigmask, &oldmask)
}
func SIGNAL_END_CRITICAL() __auto__ {
	return ZendSigprocmask(SIG_SETMASK, &oldmask, nil)
}
func ZendSigprocmask(signo __auto__, set *sigset_t, oldset *sigset_t) __auto__ {
	return sigprocmask(signo, set, oldset)
}
func ZendSignalHandlerDefer(signo int, siginfo *siginfo_t, context any) {
	var errno_save int = errno
	var queue *ZendSignalQueueT
	var qtmp *ZendSignalQueueT
	if EXPECTED(SIGG(active)) {
		if UNEXPECTED(SIGG(depth) == 0) {
			if UNEXPECTED(SIGG(blocked)) {
				SIGG(blocked) = 0
			}
			if EXPECTED(SIGG(running) == 0) {
				SIGG(running) = 1
				ZendSignalHandler(signo, siginfo, context)
				queue = SIGG(phead)
				SIGG(phead) = nil
				for queue != nil {
					ZendSignalHandler(queue.GetZendSignal().GetSigno(), queue.GetZendSignal().GetSiginfo(), queue.GetZendSignal().GetContext())
					qtmp = queue.GetNext()
					queue.SetNext(SIGG(pavail))
					queue.GetZendSignal().SetSigno(0)
					SIGG(pavail) = queue
					queue = qtmp
				}
				SIGG(running) = 0
			}
		} else {
			SIGG(blocked) = 1
			if b.Assign(&queue, SIGG(pavail)) {
				SIGG(pavail) = queue.GetNext()
				queue.GetZendSignal().SetSigno(signo)
				queue.GetZendSignal().SetSiginfo(siginfo)
				queue.GetZendSignal().SetContext(context)
				queue.SetNext(nil)
				if SIGG(phead) && SIGG(ptail) {
					SIGG(ptail).next = queue
				} else {
					SIGG(phead) = queue
				}
				SIGG(ptail) = queue
			}
		}
	} else {

		/* need to just run handler if we're inactive and getting a signal */

		ZendSignalHandler(signo, siginfo, context)

		/* need to just run handler if we're inactive and getting a signal */

	}
	errno = errno_save
}
func ZendSignalHandlerUnblock() {
	var queue *ZendSignalQueueT
	var zend_signal ZendSignalT
	if EXPECTED(SIGG(active)) {
		SIGNAL_BEGIN_CRITICAL()
		queue = SIGG(phead)
		SIGG(phead) = queue.GetNext()
		zend_signal = queue.GetZendSignal()
		queue.SetNext(SIGG(pavail))
		queue.GetZendSignal().SetSigno(0)
		SIGG(pavail) = queue
		ZendSignalHandlerDefer(zend_signal.GetSigno(), zend_signal.GetSiginfo(), zend_signal.GetContext())
		SIGNAL_END_CRITICAL()
	}
}
func ZendSignalHandler(signo int, siginfo *siginfo_t, context any) {
	var errno_save int = errno
	var sa __struct__sigaction
	var sigset sigset_t
	var p_sig ZendSignalEntryT
	p_sig = SIGG(handlers)[signo-1]
	if p_sig.GetHandler() == SIG_DFL {
		if sigaction(signo, nil, &sa) == 0 {
			sa.sa_handler = SIG_DFL
			sigemptyset(&sa.sa_mask)
			sigemptyset(&sigset)
			sigaddset(&sigset, signo)
			if sigaction(signo, &sa, nil) == 0 {

				/* throw away any blocked signals */

				ZendSigprocmask(SIG_UNBLOCK, &sigset, nil)
				kill(getpid(), signo)
			}
		}
	} else if p_sig.GetHandler() != SIG_IGN {
		if p_sig.IsSiginfo() {
			if p_sig.IsResethand() {
				SIGG(handlers)[signo-1].flags = 0
				SIGG(handlers)[signo-1].handler = SIG_DFL
			}
			(*((func(int, *siginfo_t, any))(p_sig.GetHandler())))(signo, siginfo, context)
		} else {
			(*((func(int))(p_sig.GetHandler())))(signo)
		}
	}
	errno = errno_save
}
func ZendSigaction(signo int, act *__struct__sigaction, oldact *__struct__sigaction) int {
	var sa __struct__sigaction
	var sigset sigset_t
	if oldact != nil {
		oldact.sa_flags = SIGG(handlers)[signo-1].flags
		oldact.sa_handler = any(SIGG(handlers)[signo-1].handler)
		oldact.sa_mask = GlobalSigmask
	}
	if act != nil {
		SIGG(handlers)[signo-1].flags = act.sa_flags
		if (act.sa_flags & SA_SIGINFO) != 0 {
			SIGG(handlers)[signo-1].handler = any(act.sa_sigaction)
		} else {
			SIGG(handlers)[signo-1].handler = any(act.sa_handler)
		}
		memset(&sa, 0, b.SizeOf("sa"))
		if SIGG(handlers)[signo-1].handler == any(SIG_IGN) {
			sa.sa_sigaction = any(SIG_IGN)
		} else {
			sa.sa_flags = SA_SIGINFO | act.sa_flags&SA_FLAGS_MASK
			sa.sa_sigaction = ZendSignalHandlerDefer
			sa.sa_mask = GlobalSigmask
		}
		if sigaction(signo, &sa, nil) < 0 {
			ZendErrorNoreturn(E_ERROR, "Error installing signal handler for %d", signo)
		}

		/* unsure this signal is not blocked */

		sigemptyset(&sigset)
		sigaddset(&sigset, signo)
		ZendSigprocmask(SIG_UNBLOCK, &sigset, nil)
	}
	return SUCCESS
}
func ZendSignal(signo int, handler func(int)) int {
	var sa __struct__sigaction
	memset(&sa, 0, b.SizeOf("sa"))
	sa.sa_flags = 0
	sa.sa_handler = handler
	sa.sa_mask = GlobalSigmask
	return ZendSigaction(signo, &sa, nil)
}
func ZendSignalRegister(signo int, handler func(int, *siginfo_t, any)) int {
	var sa __struct__sigaction
	if sigaction(signo, nil, &sa) == 0 {
		if (sa.sa_flags&SA_SIGINFO) != 0 && sa.sa_sigaction == handler {
			return FAILURE
		}
		SIGG(handlers)[signo-1].flags = sa.sa_flags
		if (sa.sa_flags & SA_SIGINFO) != 0 {
			SIGG(handlers)[signo-1].handler = any(sa.sa_sigaction)
		} else {
			SIGG(handlers)[signo-1].handler = any(sa.sa_handler)
		}
		sa.sa_flags = SA_SIGINFO
		sa.sa_sigaction = handler
		sa.sa_mask = GlobalSigmask
		if sigaction(signo, &sa, nil) < 0 {
			ZendErrorNoreturn(E_ERROR, "Error installing signal handler for %d", signo)
		}
		return SUCCESS
	}
	return FAILURE
}
func ZendSignalActivate() {
	var x int
	memcpy(&SIGG(handlers), &GlobalOrigHandlers, b.SizeOf("global_orig_handlers"))
	if SIGG(reset) {
		for x = 0; x < b.SizeOf("zend_sigs")/b.SizeOf("* zend_sigs"); x++ {
			ZendSignalRegister(ZendSigs[x], ZendSignalHandlerDefer)
		}
	}
	SIGG(active) = 1
	SIGG(depth) = 0
}
func ZendSignalDeactivate() {
	if SIGG(check) {
		var x int
		var sa __struct__sigaction
		if SIGG(depth) != 0 {
			ZendError(E_CORE_WARNING, "zend_signal: shutdown with non-zero blocking depth (%d)", SIGG(depth))
		}

		/* did anyone steal our installed handler */

		for x = 0; x < b.SizeOf("zend_sigs")/b.SizeOf("* zend_sigs"); x++ {
			sigaction(ZendSigs[x], nil, &sa)
			if sa.sa_sigaction != ZendSignalHandlerDefer && sa.sa_sigaction != any(SIG_IGN) {
				ZendError(E_CORE_WARNING, "zend_signal: handler was replaced for signal (%d) after startup", ZendSigs[x])
			}
		}

		/* did anyone steal our installed handler */

	}

	/* After active=0 is set, signal handlers will be called directly and other
	 * state that is reset below will not be accessed. */

	*((*volatile__int)(&SIGG(active))) = 0
	SIGG(running) = 0
	SIGG(blocked) = 0
	SIGG(depth) = 0

	/* If there are any queued signals because of a missed unblock, drop them. */

	if SIGG(phead) && SIGG(ptail) {
		SIGG(ptail).next = SIGG(pavail)
		SIGG(pavail) = SIGG(phead)
		SIGG(phead) = nil
		SIGG(ptail) = nil
	}

	/* If there are any queued signals because of a missed unblock, drop them. */
}
func ZendSignalGlobalsCtor(zend_signal_globals *ZendSignalGlobalsT) {
	var x int
	memset(zend_signal_globals, 0, b.SizeOf("* zend_signal_globals"))
	zend_signal_globals.SetReset(1)
	for x = 0; x < b.SizeOf("zend_signal_globals -> pstorage")/b.SizeOf("* zend_signal_globals -> pstorage"); x++ {
		var queue *ZendSignalQueueT = &zend_signal_globals.pstorage[x]
		queue.GetZendSignal().SetSigno(0)
		queue.SetNext(zend_signal_globals.GetPavail())
		zend_signal_globals.SetPavail(queue)
	}
}
func ZendSignalInit() {
	var signo int
	var sa __struct__sigaction

	/* Save previously registered signal handlers into orig_handlers */

	memset(&GlobalOrigHandlers, 0, b.SizeOf("global_orig_handlers"))
	for signo = 1; signo < NSIG; signo++ {
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
