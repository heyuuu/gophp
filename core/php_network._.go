package core

const Closesocket = close
const EWOULDBLOCK = EAGAIN

/* like strerror, but caller must efree the returned string,
 * unless buf is not NULL.
 * Also works sensibly for win32 */

/* These are here, rather than with the win32 counterparts above,
 * since <sys/socket.h> defines them. */

const SHUT_RD = 0
const SHUT_WR = 1
const SHUT_RDWR = 2

type PhpSocketT = int

const SOCK_RECV_ERR = -1
const STREAM_SOCKOP_NONE long = 1 << 0
const STREAM_SOCKOP_SO_REUSEPORT = 1 << 1
const STREAM_SOCKOP_SO_BROADCAST = 1 << 2
const STREAM_SOCKOP_IPV6_V6ONLY = 1 << 3
const STREAM_SOCKOP_IPV6_V6ONLY_ENABLED = 1 << 4
const STREAM_SOCKOP_TCP_NODELAY = 1 << 5

/* uncomment this to debug poll(2) emulation on systems that have poll(2) */

type PhpPollfd = __struct__pollfd

const PHP_POLLREADABLE = POLLIN | POLLERR | POLLHUP

/* timeval-to-timeout (for poll(2)) */

/* hybrid select(2)/poll(2) for a single descriptor.
 * timeouttv follows same rules as select(2), but is reduced to millisecond accuracy.
 * Returns 0 on timeout, -1 on error, or the event mask (ala poll(2)).
 */

/* emit warning and suggestion for unsafe select(2) usage */

const PHP_SOCK_CHUNK_SIZE = 8192

type PhpSockaddrStorage = __struct__sockaddr_storage

var PhpStreamSocketOps PhpStreamOps
var PhpStreamGenericSocketOps PhpStreamOps

const PHP_STREAM_IS_SOCKET = &PhpStreamSocketOps

/* open a connection to a host using php_hostconnect and return a stream */

/* {{{ memory debug */

const MAXFQDNLEN = 255
