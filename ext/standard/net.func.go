package standard

import (
	"bytes"
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func PhpInetNtop(addr *__struct__sockaddr) (string, bool) {
	var addrlen socklen_t = b.SizeOf("struct sockaddr_in")
	if addr == nil {
		return nil
	}

	/* Prefer inet_ntop() as it's more task-specific and doesn't have to be demangled */

	switch addr.sa_family {
	case AF_INET:
		var buf = make([]byte, INET_ADDRSTRLEN)
		if inet_ntop(AF_INET, &((*__struct__sockaddr_in)(addr).sin_addr), buf, INET_ADDRSTRLEN) {
			return string(buf[:strlen(buf)]), true
		}
	}

	/* Fallback on getnameinfo() */
	switch addr.sa_family {
	case AF_INET:
		var buf = make([]byte, NI_MAXHOST)
		if getnameinfo(addr, addrlen, buf, NI_MAXHOST, nil, 0, NI_NUMERICHOST) == types.SUCCESS {
			/* Also demangle numeric host with %name suffix */
			if idx := bytes.IndexByte(buf, '%'); idx >= 0 {
				buf[idx] = 0
			}
			return string(buf[:strlen(buf)]), true
		}
	}
	return "", false
}
func IfaceAppendUnicast(
	unicast *types.Zval,
	flags zend.ZendLong,
	addr *__struct__sockaddr,
	netmask *__struct__sockaddr,
	broadcast *__struct__sockaddr,
	ptp *__struct__sockaddr,
) {
	var u types.Zval
	zend.ArrayInit(&u)
	zend.AddAssocLong(&u, "flags", flags)
	if addr != nil {
		zend.AddAssocLong(&u, "family", addr.sa_family)
		if host, ok := PhpInetNtop(addr); ok {
			zend.AddAssocStr(&u, "address", host)
		}
	}
	if host, ok := PhpInetNtop(addr); ok {
		zend.AddAssocStr(&u, "netmask", host)
	}
	if host, ok := PhpInetNtop(broadcast); ok {
		zend.AddAssocStr(&u, "broadcast", host)
	}
	if host, ok := PhpInetNtop(ptp); ok {
		zend.AddAssocStr(&u, "ptp", host)
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
		core.PhpError(faults.E_WARNING, fmt.Sprintf("getifaddrs() failed %d: %s", errno, strerror(errno)))
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	for p = addrs; p != nil; p = p.ifa_next {
		var iface *types.Zval = return_value.Array().KeyFind(b.CastStrAuto(p.ifa_name))
		var unicast *types.Zval
		var status *types.Zval
		if iface == nil {
			var newif types.Zval
			zend.ArrayInit(&newif)
			iface = return_value.Array().KeyAdd(b.CastStrAuto(p.ifa_name), &newif)
		}
		unicast = iface.Array().KeyFind("unicast")
		if unicast == nil {
			var newuni types.Zval
			zend.ArrayInit(&newuni)
			unicast = iface.Array().KeyAdd("unicast", &newuni)
		}
		IfaceAppendUnicast(unicast, p.ifa_flags, p.ifa_addr, p.ifa_netmask, lang.CondF1((p.ifa_flags&IFF_BROADCAST) != 0, func() __auto__ { return p.ifa_broadaddr }, nil), lang.CondF1((p.ifa_flags&IFF_POINTOPOINT) != 0, func() __auto__ { return p.ifa_dstaddr }, nil))
		status = iface.Array().KeyFind("up")
		if status == nil {
			zend.AddAssocBool(iface, "up", (p.ifa_flags&IFF_UP) != 0)
		}
	}
	freeifaddrs(addrs)
}
