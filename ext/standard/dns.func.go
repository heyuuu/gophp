package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func ZifGethostname(executeData zpp.Ex, return_value zpp.Ret) {
	var buf []byte
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if gethostname(buf, b.SizeOf("buf")) {
		core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("unable to fetch host [%d]: %s", errno, strerror(errno)))
		return_value.SetFalse()
		return
	}
	return_value.SetString(b.CastStrAuto(buf))
	return
}
func ZifGethostbyaddr(executeData zpp.Ex, return_value zpp.Ret, ipAddress *types.Zval) {
	var addr *byte
	var addr_len int
	var hostname *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			addr, addr_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	hostname = PhpGethostbyaddr(addr)
	if hostname == nil {
		core.PhpErrorDocref("", faults.E_WARNING, "Address is not a valid IPv4 or IPv6 address")
		return_value.SetFalse()
	} else {
		return_value.SetStringEx(hostname)
	}
}
func PhpGethostbyaddr(ip *byte) *types.String {
	var addr6 __struct__in6_addr
	var addr __struct__in_addr
	var hp *__struct__hostent
	if inet_pton(AF_INET6, ip, &addr6) {
		hp = gethostbyaddr((*byte)(&addr6), b.SizeOf("addr6"), AF_INET6)
	} else if inet_pton(AF_INET, ip, &addr) {
		hp = gethostbyaddr((*byte)(&addr), b.SizeOf("addr"), AF_INET)
	} else {
		return nil
	}
	if hp == nil || hp.h_name == nil || hp.h_name[0] == '0' {
		return types.NewString(ip)
	}
	return types.NewString(hp.h_name)
}
func ZifGethostbyname(executeData zpp.Ex, return_value zpp.Ret, hostname *types.Zval) {
	var hostname *byte
	var hostname_len int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			hostname, hostname_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if hostname_len > core.MAXFQDNLEN {

		/* name too long, protect from CVE-2015-0235 */

		core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Host name is too long, the limit is %d characters", core.MAXFQDNLEN))
		return_value.SetString(b.CastStr(hostname, hostname_len))
		return
	}
	return_value.SetStringEx(PhpGethostbyname(hostname))
	return
}
func ZifGethostbynamel(executeData zpp.Ex, return_value zpp.Ret, hostname *types.Zval) {
	var hostname *byte
	var hostname_len int
	var hp *__struct__hostent
	var in __struct__in_addr
	var i int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			hostname, hostname_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if hostname_len > core.MAXFQDNLEN {

		/* name too long, protect from CVE-2015-0235 */

		core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Host name is too long, the limit is %d characters", core.MAXFQDNLEN))
		return_value.SetFalse()
		return
	}
	hp = core.PhpNetworkGethostbyname(hostname)
	if hp == nil {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	for i = 0; ; i++ {

		/* On macos h_addr_list entries may be misaligned. */

		var h_addr_entry *__struct__in_addr
		memcpy(&h_addr_entry, hp.h_addr_list[i], b.SizeOf("struct in_addr *"))
		if h_addr_entry == nil {
			return
		}
		in = *h_addr_entry
		zend.AddNextIndexString(return_value, inet_ntoa(in))
	}
}
func PhpGethostbyname(name *byte) *types.String {
	var hp *__struct__hostent
	var h_addr_0 *__struct__in_addr
	var in __struct__in_addr
	var address *byte
	hp = core.PhpNetworkGethostbyname(name)
	if hp == nil {
		return types.NewString(name)
	}

	/* On macos h_addr_list entries may be misaligned. */

	memcpy(&h_addr_0, hp.h_addr_list[0], b.SizeOf("struct in_addr *"))
	if h_addr_0 == nil {
		return types.NewString(name)
	}
	memcpy(in.s_addr, h_addr_0, b.SizeOf("in . s_addr"))
	address = inet_ntoa(in)
	return types.NewString(address)
}

//@zif -alias checkdnsrr
func ZifDnsCheckRecord(executeData zpp.Ex, return_value zpp.Ret, host *types.Zval, _ zpp.Opt, type_ *types.Zval) {
	var hp *HEADER
	var answer Querybuf
	var hostname *byte
	var rectype *byte = nil
	var hostname_len int
	var rectype_len int = 0
	var type_ int = DNS_T_MX
	var i int
	var from __struct__sockaddr_storage
	var fromsize uint32 = b.SizeOf("from")
	var handle dns_handle_t
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			hostname, hostname_len = fp.ParseString()
			fp.StartOptional()
			rectype, rectype_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if hostname_len == 0 {
		core.PhpErrorDocref("", faults.E_WARNING, "Host cannot be empty")
		return_value.SetFalse()
		return
	}
	if rectype != nil {
		if !(strcasecmp("A", rectype)) {
			type_ = DNS_T_A
		} else if !(strcasecmp("NS", rectype)) {
			type_ = DNS_T_NS
		} else if !(strcasecmp("MX", rectype)) {
			type_ = DNS_T_MX
		} else if !(strcasecmp("PTR", rectype)) {
			type_ = DNS_T_PTR
		} else if !(strcasecmp("ANY", rectype)) {
			type_ = DNS_T_ANY
		} else if !(strcasecmp("SOA", rectype)) {
			type_ = DNS_T_SOA
		} else if !(strcasecmp("CAA", rectype)) {
			type_ = DNS_T_CAA
		} else if !(strcasecmp("TXT", rectype)) {
			type_ = DNS_T_TXT
		} else if !(strcasecmp("CNAME", rectype)) {
			type_ = DNS_T_CNAME
		} else if !(strcasecmp("AAAA", rectype)) {
			type_ = DNS_T_AAAA
		} else if !(strcasecmp("SRV", rectype)) {
			type_ = DNS_T_SRV
		} else if !(strcasecmp("NAPTR", rectype)) {
			type_ = DNS_T_NAPTR
		} else if !(strcasecmp("A6", rectype)) {
			type_ = DNS_T_A6
		} else {
			core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type '%s' not supported", rectype))
			return_value.SetFalse()
			return
		}
	}
	handle = dns_open(nil)
	if handle == nil {
		return_value.SetFalse()
		return
	}
	i = PhpDnsSearch(handle, hostname, C_IN, type_, answer.GetQb2(), b.SizeOf(answer))
	PhpDnsFreeHandle(handle)
	if i < 0 {
		return_value.SetFalse()
		return
	}
	hp = (*HEADER)(&answer)
	return_value.SetBool(ntohs(hp.ancount) != 0)
	return
}
func PhpParserr(
	cp *byte,
	end *byte,
	answer *Querybuf,
	type_to_fetch int,
	store int,
	raw int,
	subarray *types.Zval,
) *byte {
	var type_ u_short
	var class u_short
	var dlen u_short
	var ttl u_long
	var n long
	var i long
	var s u_short
	var tp *u_char
	var p *u_char
	var name []byte
	var have_v6_break int = 0
	var in_v6_break int = 0
	subarray.SetUndef()
	n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf("name")-2)
	if n < 0 {
		return nil
	}
	cp += n
	if cp+10 > end {
		return nil
	}
	GETSHORT(type_, cp)
	GETSHORT(class, cp)
	GETLONG(ttl, cp)
	GETSHORT(dlen, cp)
	if cp+dlen > end {
		return nil
	}
	if dlen == 0 {

		/* No data in the response - nothing to do */

		return nil

		/* No data in the response - nothing to do */

	}
	if type_to_fetch != DNS_T_ANY && type_ != type_to_fetch {
		cp += dlen
		return cp
	}
	if store == 0 {
		cp += dlen
		return cp
	}
	zend.ArrayInit(subarray)
	zend.AddAssocString(subarray, "host", name)
	zend.AddAssocString(subarray, "class", "IN")
	zend.AddAssocLong(subarray, "ttl", ttl)
	void(class)
	if raw != 0 {
		zend.AddAssocLong(subarray, "type", type_)
		zend.AddAssocStringl(subarray, "data", (*byte)(cp), uint32(dlen))
		cp += dlen
		return cp
	}
	switch type_ {
	case DNS_T_A:
		if cp+4 > end {
			return nil
		}
		zend.AddAssocString(subarray, "type", "A")
		core.Snprintf(name, b.SizeOf("name"), "%d.%d.%d.%d", cp[0], cp[1], cp[2], cp[3])
		zend.AddAssocString(subarray, "ip", name)
		cp += dlen
	case DNS_T_MX:
		if cp+2 > end {
			return nil
		}
		zend.AddAssocString(subarray, "type", "MX")
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "pri", n)
		fallthrough
	case DNS_T_CNAME:
		if type_ == DNS_T_CNAME {
			zend.AddAssocString(subarray, "type", "CNAME")
		}
		fallthrough
	case DNS_T_NS:
		if type_ == DNS_T_NS {
			zend.AddAssocString(subarray, "type", "NS")
		}
		fallthrough
	case DNS_T_PTR:
		if type_ == DNS_T_PTR {
			zend.AddAssocString(subarray, "type", "PTR")
		}
		n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocString(subarray, "target", name)
	case DNS_T_HINFO:

		/* See RFC 1010 for values */

		zend.AddAssocString(subarray, "type", "HINFO")
		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "cpu", (*byte)(cp), n)
		cp += n
		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "os", (*byte)(cp), n)
		cp += n
	case DNS_T_CAA:

		/* See RFC 6844 for values https://tools.ietf.org/html/rfc6844 */

		zend.AddAssocString(subarray, "type", "CAA")

		// 1 flag byte

		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		zend.AddAssocLong(subarray, "flags", n)
		cp++

		// Tag length (1 byte)

		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "tag", (*byte)(cp), n)
		cp += n
		if int(dlen < int(n)+2) != 0 {
			return nil
		}
		n = dlen - n - 2
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "value", (*byte)(cp), n)
		cp += n
	case DNS_T_TXT:
		var l1 int = 0
		var entries types.Zval
		zend.AddAssocStr(subarray, "type", "TXT")
		zend.ArrayInit(&entries)

		var buf strings.Builder
		for l1 < dlen {
			n = cp[l1]
			if l1+n >= dlen {
				// Invalid chunk length, truncate
				n = dlen - (l1 + 1)
			}
			if n != 0 {
				buf.WriteString(cp[l1+1 : l1+1+n])
				zend.AddNextIndexStringl(&entries, (*byte)(cp+l1+1), n)
			}
			l1 = l1 + n + 1
		}
		cp += dlen
		zend.AddAssocStr(subarray, "txt", buf.String())
		zend.AddAssocZval(subarray, "entries", &entries)
	case DNS_T_SOA:
		zend.AddAssocString(subarray, "type", "SOA")
		n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocString(subarray, "mname", name)
		n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocString(subarray, "rname", name)
		if cp+5*4 > end {
			return nil
		}
		GETLONG(n, cp)
		zend.AddAssocLong(subarray, "serial", n)
		GETLONG(n, cp)
		zend.AddAssocLong(subarray, "refresh", n)
		GETLONG(n, cp)
		zend.AddAssocLong(subarray, "retry", n)
		GETLONG(n, cp)
		zend.AddAssocLong(subarray, "expire", n)
		GETLONG(n, cp)
		zend.AddAssocLong(subarray, "minimum-ttl", n)
	case DNS_T_AAAA:
		tp = (*u_char)(name)
		if cp+8*2 > end {
			return nil
		}
		for i = 0; i < 8; i++ {
			GETSHORT(s, cp)
			if s != 0 {
				if tp > (*u_char)(name) {
					in_v6_break = 0
					tp[0] = ':'
					tp++
				}
				tp += sprintf((*byte)(tp), "%x", s)
			} else {
				if have_v6_break == 0 {
					have_v6_break = 1
					in_v6_break = 1
					tp[0] = ':'
					tp++
				} else if in_v6_break == 0 {
					tp[0] = ':'
					tp++
					tp[0] = '0'
					tp++
				}
			}
		}
		if have_v6_break != 0 && in_v6_break != 0 {
			tp[0] = ':'
			tp++
		}
		tp[0] = '0'
		zend.AddAssocString(subarray, "type", "AAAA")
		zend.AddAssocString(subarray, "ipv6", name)
	case DNS_T_A6:
		p = cp
		zend.AddAssocString(subarray, "type", "A6")
		if cp+1 > end {
			return nil
		}
		n = int(cp[0]) & 0xff
		cp++
		zend.AddAssocLong(subarray, "masklen", n)
		tp = (*u_char)(name)
		if n > 15 {
			have_v6_break = 1
			in_v6_break = 1
			tp[0] = ':'
			tp++
		}
		if n%16 > 8 {

			/* Partial short */

			if cp[0] != 0 {
				if tp > (*u_char)(name) {
					in_v6_break = 0
					tp[0] = ':'
					tp++
				}
				sprintf((*byte)(tp), "%x", cp[0]&0xff)
			} else {
				if have_v6_break == 0 {
					have_v6_break = 1
					in_v6_break = 1
					tp[0] = ':'
					tp++
				} else if in_v6_break == 0 {
					tp[0] = ':'
					tp++
					tp[0] = '0'
					tp++
				}
			}
			cp++
		}
		for i = (n + 8) / 16; i < 8; i++ {
			if cp+2 > end {
				return nil
			}
			GETSHORT(s, cp)
			if s != 0 {
				if tp > (*u_char)(name) {
					in_v6_break = 0
					tp[0] = ':'
					tp++
				}
				tp += sprintf((*byte)(tp), "%x", s)
			} else {
				if have_v6_break == 0 {
					have_v6_break = 1
					in_v6_break = 1
					tp[0] = ':'
					tp++
				} else if in_v6_break == 0 {
					tp[0] = ':'
					tp++
					tp[0] = '0'
					tp++
				}
			}
		}
		if have_v6_break != 0 && in_v6_break != 0 {
			tp[0] = ':'
			tp++
		}
		tp[0] = '0'
		zend.AddAssocString(subarray, "ipv6", name)
		if cp < p+dlen {
			n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
			if n < 0 {
				return nil
			}
			cp += n
			zend.AddAssocString(subarray, "chain", name)
		}
	case DNS_T_SRV:
		if cp+3*2 > end {
			return nil
		}
		zend.AddAssocString(subarray, "type", "SRV")
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "pri", n)
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "weight", n)
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "port", n)
		n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocString(subarray, "target", name)
	case DNS_T_NAPTR:
		if cp+2*2 > end {
			return nil
		}
		zend.AddAssocString(subarray, "type", "NAPTR")
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "order", n)
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "pref", n)
		if cp+1 > end {
			return nil
		}
		n = cp[0] & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "flags", (*byte)(cp), n)
		cp += n
		if cp+1 > end {
			return nil
		}
		n = cp[0] & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "services", (*byte)(cp), n)
		cp += n
		if cp+1 > end {
			return nil
		}
		n = cp[0] & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "regex", (*byte)(cp), n)
		cp += n
		n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocString(subarray, "replacement", name)
	default:
		// zend.ZvalPtrDtor(subarray)
		subarray.SetUndef()
		cp += dlen
	}
	return cp
}
func ZifDnsGetRecord(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var hostname *byte
	var hostname_len int
	var type_param zend.ZendLong = PHP_DNS_ANY
	var authns *types.Zval = nil
	var addtl *types.Zval = nil
	var type_to_fetch int
	var dns_errno int
	var from __struct__sockaddr_storage
	var fromsize uint32 = b.SizeOf("from")
	var handle dns_handle_t
	var hp *HEADER
	var answer Querybuf
	var cp *u_char = nil
	var end *u_char = nil
	var n int
	var qd int
	var an int
	var ns int = 0
	var ar int = 0
	var type_ int
	var first_query int = 1
	var store_results int = 1
	var raw bool = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 5, 0)
			hostname, hostname_len = fp.ParseString()
			fp.StartOptional()
			type_param = fp.ParseLong()
			authns = fp.ParseZval()
			addtl = fp.ParseZval()
			raw = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if authns != nil {
		authns = zend.ZendTryArrayInit(authns)
		if authns == nil {
			return
		}
	}
	if addtl != nil {
		addtl = zend.ZendTryArrayInit(addtl)
		if addtl == nil {
			return
		}
	}
	if raw == 0 {
		if (type_param & ^PHP_DNS_ALL) != 0 && type_param != PHP_DNS_ANY {
			core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Type '%d' not supported", type_param))
			return_value.SetFalse()
			return
		}
	} else {
		if type_param < 1 || type_param > 0xffff {
			core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("Numeric DNS record type must be between 1 and 65535, '%d' given", type_param))
			return_value.SetFalse()
			return
		}
	}

	/* Initialize the return array */

	zend.ArrayInit(return_value)

	/* - We emulate an or'ed type mask by querying type by type. (Steps 0 - NUMTYPES-1 )
	 *   If additional info is wanted we check again with DNS_T_ANY (step NUMTYPES / NUMTYPES+1 )
	 *   store_results is used to skip storing the results retrieved in step
	 *   NUMTYPES+1 when results were already fetched.
	 * - In case of PHP_DNS_ANY we use the directly fetch DNS_T_ANY. (step NUMTYPES+1 )
	 * - In case of raw mode, we query only the requestd type instead of looping type by type
	 *   before going with the additional info stuff.
	 */

	if raw != 0 {
		type_ = -1
	} else if type_param == PHP_DNS_ANY {
		type_ = PHP_DNS_NUM_TYPES + 1
	} else {
		type_ = 0
	}
	for ; type_ < lang.Cond(addtl != nil, PHP_DNS_NUM_TYPES+2, PHP_DNS_NUM_TYPES) || first_query != 0; type_++ {
		first_query = 0
		switch type_ {
		case -1:
			type_to_fetch = type_param

			/* skip over the rest and go directly to additional records */

			type_ = PHP_DNS_NUM_TYPES - 1
		case 0:
			if (type_param & PHP_DNS_A) != 0 {
				type_to_fetch = DNS_T_A
			} else {
				type_to_fetch = 0
			}
		case 1:
			if (type_param & PHP_DNS_NS) != 0 {
				type_to_fetch = DNS_T_NS
			} else {
				type_to_fetch = 0
			}
		case 2:
			if (type_param & PHP_DNS_CNAME) != 0 {
				type_to_fetch = DNS_T_CNAME
			} else {
				type_to_fetch = 0
			}
		case 3:
			if (type_param & PHP_DNS_SOA) != 0 {
				type_to_fetch = DNS_T_SOA
			} else {
				type_to_fetch = 0
			}
		case 4:
			if (type_param & PHP_DNS_PTR) != 0 {
				type_to_fetch = DNS_T_PTR
			} else {
				type_to_fetch = 0
			}
		case 5:
			if (type_param & PHP_DNS_HINFO) != 0 {
				type_to_fetch = DNS_T_HINFO
			} else {
				type_to_fetch = 0
			}
		case 6:
			if (type_param & PHP_DNS_MX) != 0 {
				type_to_fetch = DNS_T_MX
			} else {
				type_to_fetch = 0
			}
		case 7:
			if (type_param & PHP_DNS_TXT) != 0 {
				type_to_fetch = DNS_T_TXT
			} else {
				type_to_fetch = 0
			}
		case 8:
			if (type_param & PHP_DNS_AAAA) != 0 {
				type_to_fetch = DNS_T_AAAA
			} else {
				type_to_fetch = 0
			}
		case 9:
			if (type_param & PHP_DNS_SRV) != 0 {
				type_to_fetch = DNS_T_SRV
			} else {
				type_to_fetch = 0
			}
		case 10:
			if (type_param & PHP_DNS_NAPTR) != 0 {
				type_to_fetch = DNS_T_NAPTR
			} else {
				type_to_fetch = 0
			}
		case 11:
			if (type_param & PHP_DNS_A6) != 0 {
				type_to_fetch = DNS_T_A6
			} else {
				type_to_fetch = 0
			}
		case 12:
			if (type_param & PHP_DNS_CAA) != 0 {
				type_to_fetch = DNS_T_CAA
			} else {
				type_to_fetch = 0
			}
		case PHP_DNS_NUM_TYPES:
			store_results = 0
			continue
		default:
			fallthrough
		case PHP_DNS_NUM_TYPES + 1:
			type_to_fetch = DNS_T_ANY
		}
		if type_to_fetch != 0 {
			handle = dns_open(nil)
			if handle == nil {
				return_value.Array().Destroy()
				return_value.SetFalse()
				return
			}
			n = PhpDnsSearch(handle, hostname, C_IN, type_to_fetch, answer.GetQb2(), b.SizeOf(answer))
			if n < 0 {
				dns_errno = PhpDnsErrno(handle)
				PhpDnsFreeHandle(handle)
				switch dns_errno {
				case NO_DATA:
					fallthrough
				case HOST_NOT_FOUND:
					continue
				case NO_RECOVERY:
					core.PhpErrorDocref("", faults.E_WARNING, "An unexpected server failure occurred.")
				case TRY_AGAIN:
					core.PhpErrorDocref("", faults.E_WARNING, "A temporary server error occurred.")
				default:
					core.PhpErrorDocref("", faults.E_WARNING, "DNS Query failed")
				}
				return_value.Array().Destroy()
				return_value.SetFalse()
				return
			}
			cp = answer.GetQb2() + HFIXEDSZ
			end = answer.GetQb2() + n
			hp = (*HEADER)(&answer)
			qd = ntohs(hp.qdcount)
			an = ntohs(hp.ancount)
			ns = ntohs(hp.nscount)
			ar = ntohs(hp.arcount)

			/* Skip QD entries, they're only used by dn_expand later on */

			for lang.PostDec(&qd) > 0 {
				n = dn_skipname(cp, end)
				if n < 0 {
					core.PhpErrorDocref("", faults.E_WARNING, "Unable to parse DNS data received")
					return_value.Array().Destroy()
					PhpDnsFreeHandle(handle)
					return_value.SetFalse()
					return
				}
				cp += n + QFIXEDSZ
			}

			/* YAY! Our real answers! */

			for lang.PostDec(&an) && cp != nil && cp < end {
				var retval types.Zval
				cp = PhpParserr(cp, end, &answer, type_to_fetch, store_results, raw, &retval)
				if retval.IsNotUndef() && store_results != 0 {
					zend.AddNextIndexZval(return_value, &retval)
				}
			}
			if authns != nil || addtl != nil {

				/* List of Authoritative Name Servers
				 * Process when only requesting addtl so that we can skip through the section
				 */

				for lang.PostDec(&ns) > 0 && cp != nil && cp < end {
					var retval types.Zval
					cp = PhpParserr(cp, end, &answer, DNS_T_ANY, authns != nil, raw, &retval)
					if retval.IsNotUndef() {
						zend.AddNextIndexZval(authns, &retval)
					}
				}

				/* List of Authoritative Name Servers
				 * Process when only requesting addtl so that we can skip through the section
				 */

			}
			if addtl != nil {

				/* Additional records associated with authoritative name servers */

				for lang.PostDec(&ar) > 0 && cp != nil && cp < end {
					var retval types.Zval
					cp = PhpParserr(cp, end, &answer, DNS_T_ANY, 1, raw, &retval)
					if retval.IsNotUndef() {
						zend.AddNextIndexZval(addtl, &retval)
					}
				}

				/* Additional records associated with authoritative name servers */

			}
			PhpDnsFreeHandle(handle)
		}
	}
}

//@zif -alias getmxrr
func ZifDnsGetMx(executeData zpp.Ex, return_value zpp.Ret, hostname *types.Zval, mxhosts zpp.RefZval, _ zpp.Opt, weight zpp.RefZval) {
	var hostname *byte
	var hostname_len int
	var mx_list *types.Zval
	var weight_list *types.Zval = nil
	var count int
	var qdc int
	var type_ u_short
	var weight u_short
	var answer Querybuf
	var buf []byte
	var hp *HEADER
	var cp *u_char
	var end *u_char
	var i int
	var from __struct__sockaddr_storage
	var fromsize uint32 = b.SizeOf("from")
	var handle dns_handle_t
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			hostname, hostname_len = fp.ParseString()
			mx_list = fp.ParseZval()
			fp.StartOptional()
			weight_list = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	mx_list = zend.ZendTryArrayInit(mx_list)
	if mx_list == nil {
		return
	}
	if weight_list != nil {
		weight_list = zend.ZendTryArrayInit(weight_list)
		if weight_list == nil {
			return
		}
	}
	handle = dns_open(nil)
	if handle == nil {
		return_value.SetFalse()
		return
	}
	i = PhpDnsSearch(handle, hostname, C_IN, DNS_T_MX, answer.GetQb2(), b.SizeOf(answer))
	if i < 0 {
		PhpDnsFreeHandle(handle)
		return_value.SetFalse()
		return
	}
	hp = (*HEADER)(&answer)
	cp = answer.GetQb2() + HFIXEDSZ
	end = answer.GetQb2() + i
	for qdc = ntohs(uint16(hp.qdcount)); lang.PostDec(&qdc); cp += i + QFIXEDSZ {
		if lang.Assign(&i, dn_skipname(cp, end)) < 0 {
			PhpDnsFreeHandle(handle)
			return_value.SetFalse()
			return
		}
	}
	count = ntohs(uint16(hp.ancount))
	for lang.PreDec(&count) >= 0 && cp < end {
		if lang.Assign(&i, dn_skipname(cp, end)) < 0 {
			PhpDnsFreeHandle(handle)
			return_value.SetFalse()
			return
		}
		cp += i
		GETSHORT(type_, cp)
		cp += INT16SZ + INT32SZ
		GETSHORT(i, cp)
		if type_ != DNS_T_MX {
			cp += i
			continue
		}
		GETSHORT(weight, cp)
		if lang.Assign(&i, dn_expand(answer.GetQb2(), end, cp, buf, b.SizeOf("buf")-1)) < 0 {
			PhpDnsFreeHandle(handle)
			return_value.SetFalse()
			return
		}
		cp += i
		zend.AddNextIndexString(mx_list, buf)
		if weight_list != nil {
			zend.AddNextIndexLong(weight_list, weight)
		}
	}
	PhpDnsFreeHandle(handle)
	return_value.SetBool(mx_list.Array().Len() != 0)
	return
}
func ZmStartupDns(type_ int, module_number int) int {
	zend.RegisterLongConstant("DNS_A", PHP_DNS_A, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_NS", PHP_DNS_NS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_CNAME", PHP_DNS_CNAME, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_SOA", PHP_DNS_SOA, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_PTR", PHP_DNS_PTR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_HINFO", PHP_DNS_HINFO, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_CAA", PHP_DNS_CAA, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_MX", PHP_DNS_MX, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_TXT", PHP_DNS_TXT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_SRV", PHP_DNS_SRV, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_NAPTR", PHP_DNS_NAPTR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_AAAA", PHP_DNS_AAAA, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_A6", PHP_DNS_A6, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_ANY", PHP_DNS_ANY, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("DNS_ALL", PHP_DNS_ALL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	return types.SUCCESS
}
