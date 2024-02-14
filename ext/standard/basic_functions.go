package standard

import (
	"github.com/heyuuu/gophp/ext/standard/printer"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"net"
)

/**
 * 压缩IP转人类可读IP.
 * e.g.
 *     inet_ntop("\x7f\x00\x00\x01") == "127.0.0.1"
 * 	   inet_ntop("\x7f\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01") == "7f00::1"
 */
func ZifInetNtop(ip string) (string, bool) {
	if len(ip) == 4 || len(ip) == 16 {
		return net.IP(ip).String(), true
	} else {
		return "", false
	}
}

/**
 * 人类可读IP转压缩IP，inet_ntop 的逆操作
 */
func ZifInetPton(ipAddress string) (string, bool) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return "", false
	}
	if ipV4 := ip.To4(); ipV4 != nil {
		return string(ipV4), true
	} else {
		return string(ip), true
	}
}
func ZifIp2long(ipAddress string) (int, bool) {
	ipV4 := net.ParseIP(ipAddress).To4()
	if ipV4 == nil {
		return 0, false
	}

	bytes := []byte(ipV4)
	num := int(bytes[0])<<24 + int(bytes[1])<<16 + int(bytes[2])<<8 + int(bytes[3])
	return num, true
}
func ZifLong2ip(ipAddress int) string {
	/* autoboxes on 32bit platforms, but that's expected */
	ip := net.IPv4(
		byte(ipAddress>>24),
		byte(ipAddress>>16),
		byte(ipAddress>>8),
		byte(ipAddress),
	)
	return ip.String()
}

// @zif(onError=1)
func ZifPrintR(ctx *php.Context, var_ types.Zval, _ zpp.Opt, return_ bool) *types.Zval {
	s := printer.PrintR(ctx, var_)
	if return_ {
		return types.NewZvalString(s)
	} else {
		ctx.WriteString(s)
		return types.NewZvalTrue()
	}
}
