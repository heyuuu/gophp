package core

func PhpSocketErrno() __auto__                                 { return errno }
func PhpPoll2(ufds *PhpPollfd, nfds int, timeout int) __auto__ { return poll(ufds, nfds, timeout) }
func PhpTvtoto(timeouttv *__struct__timeval) int {
	if timeouttv != nil {
		return timeouttv.tv_sec*1000 + timeouttv.tv_usec/1000
	}
	return -1
}
func PhpPollfdFor(fd PhpSocketT, events int, timeouttv *__struct__timeval) int {
	var p PhpPollfd
	var n int
	p.fd = fd
	p.events = events
	p.revents = 0
	n = PhpPoll2(&p, 1, PhpTvtoto(timeouttv))
	if n > 0 {
		return p.revents
	}
	return n
}
func PHP_SAFE_FD_SET(fd PhpSocketT, set *fd_set) {
	if fd < FD_SETSIZE {
		FD_SET(fd, set)
	}
}
func PHP_SAFE_FD_ISSET(fd PhpSocketT, set *fd_set) bool {
	return fd < FD_SETSIZE && FD_ISSET(fd, set)
}
func PHP_SAFE_MAX_FD(m PhpSocketT, n int) {
	if m >= FD_SETSIZE {
		_phpEmitFdSetsizeWarning(m)
		m = FD_SETSIZE - 1
	}
}
func PhpStreamSockOpenFromSocket(socket PhpSocketT, persistent int) *PhpStream {
	return _phpStreamSockOpenFromSocket(socket, persistent)
}
func PhpStreamSockOpenHost(host *byte, port uint16, socktype int, timeout int, persistent int) *PhpStream {
	return _phpStreamSockOpenHost(host, port, socktype, timeout, persistent)
}
