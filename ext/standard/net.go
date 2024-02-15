package standard

import (
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
	"net"
	"strconv"
	"strings"
)

var netmasks = [...]string{
	// ip-v4 (0-32)
	"0.0.0.0",
	"128.0.0.0", "192.0.0.0", "224.0.0.0", "240.0.0.0", "248.0.0.0", "252.0.0.0", "254.0.0.0", "255.0.0.0",
	"255.128.0.0", "255.192.0.0", "255.224.0.0", "255.240.0.0", "255.248.0.0", "255.252.0.0", "255.254.0.0", "255.255.0.0",
	"255.255.128.0", "255.255.192.0", "255.255.224.0", "255.255.240.0", "255.255.248.0", "255.255.252.0", "255.255.254.0", "255.255.255.0",
	"255.255.255.128", "255.255.255.192", "255.255.255.224", "255.255.255.240", "255.255.255.248", "255.255.255.252", "255.255.255.254", "255.255.255.255",
	// ip-v6 (33-128)
	//"8000::", "c000::", "e000::", "f000::", "f800::", "fc00::", "fe00::", "ff00::", "ff80::", "ffc0::", "ffe0::", "fff0::", "fff8::", "fffc::", "fffe::", "ffff::",
	//"ffff:8000::", "ffff:c000::", "ffff:e000::", "ffff:f000::", "ffff:f800::", "ffff:fc00::", "ffff:fe00::", "ffff:ff00::", "ffff:ff80::", "ffff:ffc0::", "ffff:ffe0::", "ffff:fff0::", "ffff:fff8::", "ffff:fffc::", "ffff:fffe::", "ffff:ffff::",
	"ffff:ffff:8000::", "ffff:ffff:c000::", "ffff:ffff:e000::", "ffff:ffff:f000::", "ffff:ffff:f800::", "ffff:ffff:fc00::", "ffff:ffff:fe00::", "ffff:ffff:ff00::", "ffff:ffff:ff80::", "ffff:ffff:ffc0::", "ffff:ffff:ffe0::", "ffff:ffff:fff0::", "ffff:ffff:fff8::", "ffff:ffff:fffc::", "ffff:ffff:fffe::", "ffff:ffff:ffff::",
	"ffff:ffff:ffff:8000::", "ffff:ffff:ffff:c000::", "ffff:ffff:ffff:e000::", "ffff:ffff:ffff:f000::", "ffff:ffff:ffff:f800::", "ffff:ffff:ffff:fc00::", "ffff:ffff:ffff:fe00::", "ffff:ffff:ffff:ff00::", "ffff:ffff:ffff:ff80::", "ffff:ffff:ffff:ffc0::", "ffff:ffff:ffff:ffe0::", "ffff:ffff:ffff:fff0::", "ffff:ffff:ffff:fff8::", "ffff:ffff:ffff:fffc::", "ffff:ffff:ffff:fffe::", "ffff:ffff:ffff:ffff::",
	"ffff:ffff:ffff:ffff:8000::", "ffff:ffff:ffff:ffff:c000::", "ffff:ffff:ffff:ffff:e000::", "ffff:ffff:ffff:ffff:f000::", "ffff:ffff:ffff:ffff:f800::", "ffff:ffff:ffff:ffff:fc00::", "ffff:ffff:ffff:ffff:fe00::", "ffff:ffff:ffff:ffff:ff00::", "ffff:ffff:ffff:ffff:ff80::", "ffff:ffff:ffff:ffff:ffc0::", "ffff:ffff:ffff:ffff:ffe0::", "ffff:ffff:ffff:ffff:fff0::", "ffff:ffff:ffff:ffff:fff8::", "ffff:ffff:ffff:ffff:fffc::", "ffff:ffff:ffff:ffff:fffe::", "ffff:ffff:ffff:ffff:ffff::",
	"ffff:ffff:ffff:ffff:ffff:8000::", "ffff:ffff:ffff:ffff:ffff:c000::", "ffff:ffff:ffff:ffff:ffff:e000::", "ffff:ffff:ffff:ffff:ffff:f000::", "ffff:ffff:ffff:ffff:ffff:f800::", "ffff:ffff:ffff:ffff:ffff:fc00::", "ffff:ffff:ffff:ffff:ffff:fe00::", "ffff:ffff:ffff:ffff:ffff:ff00::", "ffff:ffff:ffff:ffff:ffff:ff80::", "ffff:ffff:ffff:ffff:ffff:ffc0::", "ffff:ffff:ffff:ffff:ffff:ffe0::", "ffff:ffff:ffff:ffff:ffff:fff0::", "ffff:ffff:ffff:ffff:ffff:fff8::", "ffff:ffff:ffff:ffff:ffff:fffc::", "ffff:ffff:ffff:ffff:ffff:fffe::", "ffff:ffff:ffff:ffff:ffff:ffff::",
	"ffff:ffff:ffff:ffff:ffff:ffff:8000::", "ffff:ffff:ffff:ffff:ffff:ffff:c000::", "ffff:ffff:ffff:ffff:ffff:ffff:e000::", "ffff:ffff:ffff:ffff:ffff:ffff:f000::", "ffff:ffff:ffff:ffff:ffff:ffff:f800::", "ffff:ffff:ffff:ffff:ffff:ffff:fc00::", "ffff:ffff:ffff:ffff:ffff:ffff:fe00::", "ffff:ffff:ffff:ffff:ffff:ffff:ff00::", "ffff:ffff:ffff:ffff:ffff:ffff:ff80::", "ffff:ffff:ffff:ffff:ffff:ffff:ffc0::", "ffff:ffff:ffff:ffff:ffff:ffff:ffe0::", "ffff:ffff:ffff:ffff:ffff:ffff:fff0::", "ffff:ffff:ffff:ffff:ffff:ffff:fff8::", "ffff:ffff:ffff:ffff:ffff:ffff:fffc::", "ffff:ffff:ffff:ffff:ffff:ffff:fffe::", "ffff:ffff:ffff:ffff:ffff:ffff:ffff::",
	"ffff:ffff:ffff:ffff:ffff:ffff:ffff:8000", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:c000", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:e000", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:f000", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:f800", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:fc00", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:fe00", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ff00", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ff80", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffc0", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffe0", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:fff0", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:fff8", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:fffc", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:fffe", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff",
}

func addrGetString(addr net.Addr) (address string, netmask string) {
	address, netmask, _ = strings.Cut(addr.String(), "/")
	if netmaskBits, err := strconv.Atoi(netmask); err == nil {
		if 0 <= netmaskBits && netmaskBits <= 128 {
			netmask = netmasks[netmaskBits]
		}
	}
	return
}

func ZifNetGetInterfaces() (*types.Array, bool) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, false
	}

	arr := types.NewArrayCap(len(ifaces))
	for _, iface := range ifaces {
		unicast := types.NewArray()
		if addrs, err := iface.Addrs(); err == nil {
			for _, addr := range addrs {
				address, netmask := addrGetString(addr)
				addrArr := types.NewArrayOfPairs([]types.ArrayPair{
					{Key: types.StrKey("flags"), Val: php.Long(32841)},
					{Key: types.StrKey("family"), Val: php.Long(2)},
					{Key: types.StrKey("address"), Val: php.String(address)},
					{Key: types.StrKey("netmask"), Val: php.String(netmask)},
				})
				unicast.Append(types.ZvalArray(addrArr))
			}
		}

		info := types.NewArray()
		info.KeyAdd("unicast", types.ZvalArray(unicast))
		info.KeyAdd("up", types.ZvalBool(iface.Flags&net.FlagUp != 0))

		arr.KeyUpdate(iface.Name, types.ZvalArray(info))
	}
	return arr, true
}
