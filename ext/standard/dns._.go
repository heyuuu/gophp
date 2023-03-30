package standard

import (
	"github.com/heyuuu/gophp/zend"
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
const PHP_DNS_ALL zend.ZendLong = PHP_DNS_A | PHP_DNS_NS | PHP_DNS_CNAME | PHP_DNS_SOA | PHP_DNS_PTR | PHP_DNS_HINFO | PHP_DNS_CAA | PHP_DNS_MX | PHP_DNS_TXT | PHP_DNS_A6 | PHP_DNS_SRV | PHP_DNS_NAPTR | PHP_DNS_AAAA

const HFIXEDSZ = 12
const QFIXEDSZ = 4
const MAXRESOURCERECORDS = 64
