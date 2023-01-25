// <<generate>>

package zend

const NSIG = 65
const ZEND_SIGNAL_QUEUE_SIZE = 64

var ZendSignalGlobals ZendSignalGlobalsT

const TIMEOUT_SIG = SIGPROF

var ZendSigs []int = []int{TIMEOUT_SIG, SIGHUP, SIGINT, SIGQUIT, SIGTERM, SIGUSR1, SIGUSR2}

const SA_FLAGS_MASK = ^(SA_NODEFER | SA_RESETHAND)

var GlobalOrigHandlers []ZendSignalEntryT
var GlobalSigmask sigset_t
