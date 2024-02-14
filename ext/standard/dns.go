package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"net"
	"os"
)

const DNS_T_A = 1
const DNS_T_NS = 2
const DNS_T_CNAME = 5
const DNS_T_SOA = 6
const DNS_T_PTR = 12
const DNS_T_HINFO = 13
const DNS_T_MINFO = 14
const DNS_T_MX = 15
const DNS_T_TXT = 16
const DNS_T_AAAA = 28
const DNS_T_SRV = 33
const DNS_T_NAPTR = 35
const DNS_T_A6 = 38
const DNS_T_CAA = 257
const DNS_T_ANY = 255

const PHP_DNS_NUM_TYPES = 13
const PHP_DNS_A = 0x1
const PHP_DNS_NS = 0x2
const PHP_DNS_CNAME = 0x10
const PHP_DNS_SOA = 0x20
const PHP_DNS_PTR = 0x800
const PHP_DNS_HINFO = 0x1000
const PHP_DNS_CAA = 0x2000
const PHP_DNS_MX = 0x4000
const PHP_DNS_TXT = 0x8000
const PHP_DNS_A6 = 0x1000000
const PHP_DNS_SRV = 0x2000000
const PHP_DNS_NAPTR = 0x4000000
const PHP_DNS_AAAA = 0x8000000
const PHP_DNS_ANY = 0x10000000
const PHP_DNS_ALL int = PHP_DNS_A | PHP_DNS_NS | PHP_DNS_CNAME | PHP_DNS_SOA | PHP_DNS_PTR | PHP_DNS_HINFO | PHP_DNS_CAA | PHP_DNS_MX | PHP_DNS_TXT | PHP_DNS_A6 | PHP_DNS_SRV | PHP_DNS_NAPTR | PHP_DNS_AAAA

const HFIXEDSZ = 12
const QFIXEDSZ = 4
const MAXRESOURCERECORDS = 64

const MaxFqdnLen = 255

func ZifGethostname() (string, bool) {
	name, err := os.Hostname()
	if err != nil {
		return "", false
	}
	return name, true
}

func ZifGethostbyaddr(ipAddress string) (string, bool) {
	names, err := net.LookupAddr(ipAddress)
	if err != nil || len(names) == 0 {
		return "", false
	}
	return names[0], true
}

func ZifGethostbyname(ctx *php.Context, hostname string) string {
	if len(hostname) > MaxFqdnLen {
		/* name too long, protect from CVE-2015-0235 */
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Host name is too long, the limit is %d characters", MaxFqdnLen))
		return hostname
	}
	addrs, err := net.LookupHost(hostname)
	if err != nil || len(addrs) == 0 {
		return hostname
	}
	return addrs[0]
}

func ZifGethostbynamel(ctx *php.Context, hostname string) ([]string, bool) {
	if len(hostname) > MaxFqdnLen {
		/* name too long, protect from CVE-2015-0235 */
		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Host name is too long, the limit is %d characters", MaxFqdnLen))
		return nil, false
	}

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, false
	}
	return addrs, true
}

func dnsGetMx(hostname string) (hosts []string, weights []int, ok bool) {
	mxes, err := net.LookupMX(hostname)
	if err != nil {
		return nil, nil, false
	}
	for _, mx := range mxes {
		hosts = append(hosts, mx.Host)
		weights = append(weights, int(mx.Pref))
	}
	return hosts, weights, true
}

// @zif(alias="getmxrr")
func ZifDnsGetMx(hostname string, hosts zpp.RefZval, _ zpp.Opt, weights zpp.RefZval) bool {
	hostValues, weightValues, ok := dnsGetMx(hostname)
	hosts.SetVal(types.ZvalArrayOfString(hostValues))
	if weights != nil {
		hosts.SetVal(types.ZvalArrayOfInt(weightValues))
	}
	return ok
}

func ZmStartupDns(ctx *php.Context, moduleNumber int) {
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_A", PHP_DNS_A)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_NS", PHP_DNS_NS)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_CNAME", PHP_DNS_CNAME)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_SOA", PHP_DNS_SOA)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_PTR", PHP_DNS_PTR)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_HINFO", PHP_DNS_HINFO)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_CAA", PHP_DNS_CAA)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_MX", PHP_DNS_MX)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_TXT", PHP_DNS_TXT)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_SRV", PHP_DNS_SRV)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_NAPTR", PHP_DNS_NAPTR)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_AAAA", PHP_DNS_AAAA)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_A6", PHP_DNS_A6)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_ANY", PHP_DNS_ANY)
	php.RegisterLongConstant(ctx, moduleNumber, "DNS_ALL", PHP_DNS_ALL)
}
