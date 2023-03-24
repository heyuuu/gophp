package standard

func PhpDnsSearch(
	res dns_handle_t,
	dname *byte,
	class __auto__,
	type_ int,
	answer []u_char,
	anslen __auto__,
) int {
	return int(dns_search(res, dname, class, type_, (*byte)(answer), anslen, (*__struct__sockaddr)(&from), &fromsize))
}
func PhpDnsFreeHandle(res dns_handle_t) __auto__ { return dns_free(res) }
func PhpDnsErrno(handle dns_handle_t) __auto__   { return h_errno }
