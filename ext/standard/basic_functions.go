package standard

import (
	"github.com/heyuuu/gophp/ext/standard/printer"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/lang"
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

func ZifIniGet(ctx *php.Context, varname string) (string, bool) {
	return ctx.INI().GetStr(varname)
}
func ZifIniGetAll(ctx *php.Context, returnValue zpp.Ret, _ zpp.Opt, extension *string, details_ *bool) {
	var details = lang.Option(details_, true)
	var moduleNumber = 0
	//if extension != nil {
	//	module := php.GetRegisteredModule(ctx.Engine(), *extension)
	//	if module == nil {
	//		php.ErrorDocRef(ctx, "", perr.E_WARNING, fmt.Sprintf("Unable to find extension '%s'", *extension))
	//		returnValue.SetFalse()
	//		return
	//	}
	//	moduleNumber = module.ModuleNumber()
	//}

	returnValue.SetEmptyArray()
	ctx.INI().EachIniDirective(true, func(global *php.IniEntry, local *php.IniEntry) {
		if moduleNumber != 0 && global.ModuleNumber() != moduleNumber {
			return
		}

		if details {
			options := types.NewArrayCap(3)
			if global.HasValue() {
				options.AddAssocStr("global_value", global.Value())
			} else {
				options.AddAssocNull("global_value")
			}
			if local.HasValue() {
				options.AddAssocStr("local_value", local.Value())
			} else {
				options.AddAssocNull("local_value")
			}
			options.AddAssocLong("access", int(global.Modifiable()))
			returnValue.Array().SymtableUpdate(global.Name(), types.ZvalArray(options))
		} else {
			if local.HasValue() {
				returnValue.Array().SymtableUpdate(local.Name(), types.ZvalString(local.Value()))
			} else {
				returnValue.Array().SymtableUpdate(local.Name(), php.UninitializedZval())
			}
		}
	})
}

// @zif(alias="ini_alter")
func ZifIniSet(ctx *php.Context, varname string, newvalue string) (string, bool) {
	oldValue, exists := ctx.INI().GetStr(varname)
	if !exists || !ctx.INI().AlterIni(varname, newvalue) {
		return "", false
	}
	return oldValue, true
}
func ZifIniRestore(ctx *php.Context, varName string) {
	ctx.INI().RestoreIni(varName, php.IniStageRuntime)
}
func ZifSetIncludePath(ctx *php.Context, newIncludePath zpp.Path) (string, bool) {
	oldValue, exists := ctx.INI().GetStr("include_path")
	if !exists || !ctx.INI().AlterIni("include_path", newIncludePath) {
		return "", false
	}
	return oldValue, true
}
func ZifGetIncludePath(ctx *php.Context) (string, bool) {
	return ctx.INI().GetStr("include_path")
}
func ZifRestoreIncludePath(ctx *php.Context) {
	key := "include_path"
	ctx.INI().RestoreIni(key, php.IniStageRuntime)
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
