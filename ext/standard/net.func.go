// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func PhpInetNtop(addr *__struct__sockaddr) *zend.ZendString {
	var addrlen socklen_t = b.SizeOf("struct sockaddr_in")
	if addr == nil {
		return nil
	}

	/* Prefer inet_ntop() as it's more task-specific and doesn't have to be demangled */

	switch addr.sa_family {
	case AF_INET:
		var ret *zend.ZendString = zend.ZendStringAlloc(INET_ADDRSTRLEN, 0)
		if inet_ntop(AF_INET, &((*__struct__sockaddr_in)(addr).sin_addr), zend.ZSTR_VAL(ret), INET_ADDRSTRLEN) {
			zend.ZSTR_LEN(ret) = strlen(zend.ZSTR_VAL(ret))
			return ret
		}
		zend.ZendStringEfree(ret)
		break
	}

	/* Fallback on getnameinfo() */

	switch addr.sa_family {
	case AF_INET:
		var ret *zend.ZendString = zend.ZendStringAlloc(NI_MAXHOST, 0)
		if getnameinfo(addr, addrlen, zend.ZSTR_VAL(ret), NI_MAXHOST, nil, 0, NI_NUMERICHOST) == zend.SUCCESS {

			/* Also demangle numeric host with %name suffix */

			var colon *byte = strchr(zend.ZSTR_VAL(ret), '%')
			if colon != nil {
				*colon = 0
			}
			zend.ZSTR_LEN(ret) = strlen(zend.ZSTR_VAL(ret))
			return ret
		}
		zend.ZendStringEfree(ret)
		break
	}
	return nil
}
func IfaceAppendUnicast(unicast *zend.Zval, flags zend.ZendLong, addr *__struct__sockaddr, netmask *__struct__sockaddr, broadcast *__struct__sockaddr, ptp *__struct__sockaddr) {
	var host *zend.ZendString
	var u zend.Zval
	zend.ArrayInit(&u)
	zend.AddAssocLong(&u, "flags", flags)
	if addr != nil {
		zend.AddAssocLong(&u, "family", addr.sa_family)
		if b.Assign(&host, PhpInetNtop(addr)) {
			zend.AddAssocStr(&u, "address", host)
		}
	}
	if b.Assign(&host, PhpInetNtop(netmask)) {
		zend.AddAssocStr(&u, "netmask", host)
	}
	if b.Assign(&host, PhpInetNtop(broadcast)) {
		zend.AddAssocStr(&u, "broadcast", host)
	}
	if b.Assign(&host, PhpInetNtop(ptp)) {
		zend.AddAssocStr(&u, "ptp", host)
	}
	zend.AddNextIndexZval(unicast, &u)
}
func ZifNetGetInterfaces(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var addrs *__struct__ifaddrs = nil
	var p *__struct__ifaddrs
	if zend.ZEND_NUM_ARGS() != 0 {
		zend.ZendWrongParametersNoneError()
		return
	}
	if getifaddrs(&addrs) {
		core.PhpError(zend.E_WARNING, "getifaddrs() failed %d: %s", errno, strerror(errno))
		zend.RETVAL_FALSE
		return
	}
	zend.ArrayInit(return_value)
	for p = addrs; p != nil; p = p.ifa_next {
		var iface *zend.Zval = zend.ZendHashStrFind(zend.Z_ARR_P(return_value), p.ifa_name, strlen(p.ifa_name))
		var unicast *zend.Zval
		var status *zend.Zval
		if iface == nil {
			var newif zend.Zval
			zend.ArrayInit(&newif)
			iface = zend.ZendHashStrAdd(zend.Z_ARR_P(return_value), p.ifa_name, strlen(p.ifa_name), &newif)
		}
		unicast = zend.ZendHashStrFind(zend.Z_ARR_P(iface), "unicast", b.SizeOf("\"unicast\"")-1)
		if unicast == nil {
			var newuni zend.Zval
			zend.ArrayInit(&newuni)
			unicast = zend.ZendHashStrAdd(zend.Z_ARR_P(iface), "unicast", b.SizeOf("\"unicast\"")-1, &newuni)
		}
		IfaceAppendUnicast(unicast, p.ifa_flags, p.ifa_addr, p.ifa_netmask, b.CondF1((p.ifa_flags&IFF_BROADCAST) != 0, func() __auto__ { return p.ifa_broadaddr }, nil), b.CondF1((p.ifa_flags&IFF_POINTOPOINT) != 0, func() __auto__ { return p.ifa_dstaddr }, nil))
		status = zend.ZendHashStrFind(zend.Z_ARR_P(iface), "up", b.SizeOf("\"up\"")-1)
		if status == nil {
			zend.AddAssocBool(iface, "up", (p.ifa_flags&IFF_UP) != 0)
		}
	}
	freeifaddrs(addrs)
}
