// <<generate>>

package standard

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/net.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
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
   | Authors: Sara Golemon <pollita@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_network.h"

// # include < arpa / inet . h >

// # include < net / if . h >

// # include < ifaddrs . h >

// # include < netdb . h >

func PhpInetNtop(addr *__struct__sockaddr) *zend.ZendString {
	var addrlen socklen_t = g.SizeOf("struct sockaddr_in")
	if addr == nil {
		return nil
	}

	/* Prefer inet_ntop() as it's more task-specific and doesn't have to be demangled */

	switch addr.sa_family {
	case AF_INET:
		var ret *zend.ZendString = zend.ZendStringAlloc(INET_ADDRSTRLEN, 0)
		if inet_ntop(AF_INET, &((*__struct__sockaddr_in)(addr).sin_addr), ret.val, INET_ADDRSTRLEN) {
			ret.len_ = strlen(ret.val)
			return ret
		}
		zend.ZendStringEfree(ret)
		break
	}

	/* Fallback on getnameinfo() */

	switch addr.sa_family {
	case AF_INET:
		var ret *zend.ZendString = zend.ZendStringAlloc(NI_MAXHOST, 0)
		if getnameinfo(addr, addrlen, ret.val, NI_MAXHOST, nil, 0, NI_NUMERICHOST) == zend.SUCCESS {

			/* Also demangle numeric host with %name suffix */

			var colon *byte = strchr(ret.val, '%')
			if colon != nil {
				*colon = 0
			}
			ret.len_ = strlen(ret.val)
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
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &u
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.AddAssocLongEx(&u, "flags", strlen("flags"), flags)
	if addr != nil {
		zend.AddAssocLongEx(&u, "family", strlen("family"), addr.sa_family)
		if g.Assign(&host, PhpInetNtop(addr)) {
			zend.AddAssocStrEx(&u, "address", strlen("address"), host)
		}
	}
	if g.Assign(&host, PhpInetNtop(netmask)) {
		zend.AddAssocStrEx(&u, "netmask", strlen("netmask"), host)
	}
	if g.Assign(&host, PhpInetNtop(broadcast)) {
		zend.AddAssocStrEx(&u, "broadcast", strlen("broadcast"), host)
	}
	if g.Assign(&host, PhpInetNtop(ptp)) {
		zend.AddAssocStrEx(&u, "ptp", strlen("ptp"), host)
	}
	zend.AddNextIndexZval(unicast, &u)
}

/* {{{ proto array|false net_get_interfaces()
Returns an array in the form:
array(
  'ifacename' => array(
    'description' => 'Awesome interface', // Win32 only
    'mac' => '00:11:22:33:44:55',         // Win32 only
    'mtu' => 1234,                        // Win32 only
    'unicast' => array(
      0 => array(
        'family' => 2,                    // e.g. AF_INET, AF_INET6, AF_PACKET
        'address' => '127.0.0.1',
        'netmnask' => '255.0.0.0',
        'broadcast' => '127.255.255.255', // POSIX only
        'ptp' => '127.0.0.2',             // POSIX only
      ), // etc...
    ),
  ), // etc...
)
*/

func ZifNetGetInterfaces(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var addrs *__struct__ifaddrs = nil
	var p *__struct__ifaddrs
	if execute_data.This.u2.num_args != 0 {
		zend.ZendWrongParametersNoneError()
		return
	}
	if getifaddrs(&addrs) {
		zend.ZendError(1<<1, "getifaddrs() failed %d: %s", errno, strerror(errno))
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for p = addrs; p != nil; p = p.ifa_next {
		var iface *zend.Zval = zend.ZendHashStrFind(return_value.value.arr, p.ifa_name, strlen(p.ifa_name))
		var unicast *zend.Zval
		var status *zend.Zval
		if iface == nil {
			var newif zend.Zval
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = &newif
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			iface = zend.ZendHashStrAdd(return_value.value.arr, p.ifa_name, strlen(p.ifa_name), &newif)
		}
		unicast = zend.ZendHashStrFind(iface.value.arr, "unicast", g.SizeOf("\"unicast\"")-1)
		if unicast == nil {
			var newuni zend.Zval
			var __arr *zend.ZendArray = zend._zendNewArray(0)
			var __z *zend.Zval = &newuni
			__z.value.arr = __arr
			__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
			unicast = zend.ZendHashStrAdd(iface.value.arr, "unicast", g.SizeOf("\"unicast\"")-1, &newuni)
		}
		IfaceAppendUnicast(unicast, p.ifa_flags, p.ifa_addr, p.ifa_netmask, g.CondF1((p.ifa_flags&IFF_BROADCAST) != 0, func() __auto__ { return p.ifa_broadaddr }, nil), g.CondF1((p.ifa_flags&IFF_POINTOPOINT) != 0, func() __auto__ { return p.ifa_dstaddr }, nil))
		status = zend.ZendHashStrFind(iface.value.arr, "up", g.SizeOf("\"up\"")-1)
		if status == nil {
			zend.AddAssocBoolEx(iface, "up", strlen("up"), (p.ifa_flags&IFF_UP) != 0)
		}
	}
	freeifaddrs(addrs)
}

/* }}} */
