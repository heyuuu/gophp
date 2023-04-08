package core

const Closesocket = close
const EWOULDBLOCK = EAGAIN

const SHUT_RD = 0
const SHUT_WR = 1
const SHUT_RDWR = 2

type PhpSocketT = int

const STREAM_SOCKOP_NONE = 1 << 0
const STREAM_SOCKOP_TCP_NODELAY = 1 << 5

type PhpPollfd = __struct__pollfd

const PHP_POLLREADABLE = POLLIN | POLLERR | POLLHUP

const PHP_SOCK_CHUNK_SIZE = 8192

type PhpSockaddrStorage = __struct__sockaddr_storage

var PhpStreamGenericSocketOps PhpStreamOps

const MAXFQDNLEN = 255
