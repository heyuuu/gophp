package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func PhpInetNtop(addr *__struct__sockaddr) *types2.String {
	var addrlen socklen_t = b.SizeOf("struct sockaddr_in")
	if addr == nil {
		return nil
	}

	/* Prefer inet_ntop() as it's more task-specific and doesn't have to be demangled */

	switch addr.sa_family {
	case AF_INET:
		var ret *types2.String = types2.ZendStringAlloc(INET_ADDRSTRLEN, 0)
		if inet_ntop(AF_INET, &((*__struct__sockaddr_in)(addr).sin_addr), ret.GetVal(), INET_ADDRSTRLEN) {
			return ret.Cutoff(strlen(ret.GetVal()))
		}
	}

	/* Fallback on getnameinfo() */

	switch addr.sa_family {
	case AF_INET:
		var ret *types2.String = types2.ZendStringAlloc(NI_MAXHOST, 0)
		if getnameinfo(addr, addrlen, ret.GetVal(), NI_MAXHOST, nil, 0, NI_NUMERICHOST) == types2.SUCCESS {

			/* Also demangle numeric host with %name suffix */

			var colon *byte = strchr(ret.GetVal(), '%')
			if colon != nil {
				*colon = 0
			}
			return ret.Cutoff(strlen(ret.GetVal()))
		}
		// types.ZendStringEfree(ret)
	}
	return nil
}
func IfaceAppendUnicast(
	unicast *types2.Zval,
	flags zend.ZendLong,
	addr *__struct__sockaddr,
	netmask *__struct__sockaddr,
	broadcast *__struct__sockaddr,
	ptp *__struct__sockaddr,
) {
	var host *types2.String
	var u types2.Zval
	zend.ArrayInit(&u)
	zend.AddAssocLong(&u, "flags", flags)
	if addr != nil {
		zend.AddAssocLong(&u, "family", addr.sa_family)
		if b.Assign(&host, PhpInetNtop(addr)) {
			zend.AddAssocStr(&u, "address", host.GetStr())
		}
	}
	if b.Assign(&host, PhpInetNtop(netmask)) {
		zend.AddAssocStr(&u, "netmask", host.GetStr())
	}
	if b.Assign(&host, PhpInetNtop(broadcast)) {
		zend.AddAssocStr(&u, "broadcast", host.GetStr())
	}
	if b.Assign(&host, PhpInetNtop(ptp)) {
		zend.AddAssocStr(&u, "ptp", host.GetStr())
	}
	zend.AddNextIndexZval(unicast, &u)
}
func ZifNetGetInterfaces(executeData zpp.Ex, return_value zpp.Ret) {
	var addrs *__struct__ifaddrs = nil
	var p *__struct__ifaddrs
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if getifaddrs(&addrs) {
		core.PhpError(faults.E_WARNING, "getifaddrs() failed %d: %s", errno, strerror(errno))
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	for p = addrs; p != nil; p = p.ifa_next {
		var iface *types2.Zval = return_value.Array().KeyFind(b.CastStrAuto(p.ifa_name))
		var unicast *types2.Zval
		var status *types2.Zval
		if iface == nil {
			var newif types2.Zval
			zend.ArrayInit(&newif)
			iface = return_value.Array().KeyAdd(b.CastStrAuto(p.ifa_name), &newif)
		}
		unicast = iface.Array().KeyFind("unicast")
		if unicast == nil {
			var newuni types2.Zval
			zend.ArrayInit(&newuni)
			unicast = iface.Array().KeyAdd("unicast", &newuni)
		}
		IfaceAppendUnicast(unicast, p.ifa_flags, p.ifa_addr, p.ifa_netmask, b.CondF1((p.ifa_flags&IFF_BROADCAST) != 0, func() __auto__ { return p.ifa_broadaddr }, nil), b.CondF1((p.ifa_flags&IFF_POINTOPOINT) != 0, func() __auto__ { return p.ifa_dstaddr }, nil))
		status = iface.Array().KeyFind("up")
		if status == nil {
			zend.AddAssocBool(iface, "up", (p.ifa_flags&IFF_UP) != 0)
		}
	}
	freeifaddrs(addrs)
}
