package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/ext/standard/array"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/globals"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"math"
	"net"
	"os"
	"strings"
	"time"
)

func BG__() *PhpBasicGlobals { return &BasicGlobals }
func BasicGlobalsCtor(basic_globals_p *PhpBasicGlobals) {
	BG__().Ctor()
}
func BasicGlobalsDtor(basic_globals_p *PhpBasicGlobals) {
	basic_globals_p.Dtor()
}
func ZmStartupBasic(type_ int, module_number int) int {
	BasicGlobalsCtor(&BasicGlobals)
	IncompleteClassEntry = PhpCreateIncompleteClass()
	BG__().incomplete_class = IncompleteClassEntry
	zend.RegisterLongConstant("CONNECTION_ABORTED", core.PHP_CONNECTION_ABORTED, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("CONNECTION_NORMAL", core.PHP_CONNECTION_NORMAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("CONNECTION_TIMEOUT", core.PHP_CONNECTION_TIMEOUT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("INI_USER", zend.ZEND_INI_USER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("INI_PERDIR", zend.ZEND_INI_PERDIR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("INI_SYSTEM", zend.ZEND_INI_SYSTEM, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("INI_ALL", zend.ZEND_INI_ALL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("INI_SCANNER_NORMAL", zend.ZEND_INI_SCANNER_NORMAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("INI_SCANNER_RAW", zend.ZEND_INI_SCANNER_RAW, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("INI_SCANNER_TYPED", zend.ZEND_INI_SCANNER_TYPED, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_URL_SCHEME", PHP_URL_SCHEME, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_URL_HOST", PHP_URL_HOST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_URL_PORT", PHP_URL_PORT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_URL_USER", PHP_URL_USER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_URL_PASS", PHP_URL_PASS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_URL_PATH", PHP_URL_PATH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_URL_QUERY", PHP_URL_QUERY, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_URL_FRAGMENT", PHP_URL_FRAGMENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_QUERY_RFC1738", PHP_QUERY_RFC1738, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_QUERY_RFC3986", PHP_QUERY_RFC3986, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)

	// #define REGISTER_MATH_CONSTANT(x) REGISTER_DOUBLE_CONSTANT ( # x , x , CONST_CS | CONST_PERSISTENT )

	zend.RegisterDoubleConstant("M_E", M_E, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_LOG2E", M_LOG2E, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_LOG10E", M_LOG10E, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_LN2", M_LN2, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_LN10", M_LN10, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_PI", M_PI, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_PI_2", M_PI_2, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_PI_4", M_PI_4, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_1_PI", M_1_PI, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_2_PI", M_2_PI, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_SQRTPI", M_SQRTPI, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_2_SQRTPI", M_2_SQRTPI, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_LNPI", M_LNPI, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_EULER", M_EULER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_SQRT2", M_SQRT2, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_SQRT1_2", M_SQRT1_2, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("M_SQRT3", M_SQRT3, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("INF", math.Inf(1), zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterDoubleConstant("NAN", math.NaN(), zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_ROUND_HALF_UP", PHP_ROUND_HALF_UP, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_ROUND_HALF_DOWN", PHP_ROUND_HALF_DOWN, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_ROUND_HALF_EVEN", PHP_ROUND_HALF_EVEN, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PHP_ROUND_HALF_ODD", PHP_ROUND_HALF_ODD, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	RegisterPhpinfoConstants(type_, module_number)
	RegisterHtmlConstants(type_, module_number)
	str.RegisterStringConstants(type_, module_number)
	if ZmStartupVar(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupFile(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupPack(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupBrowscap(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupStandardFilters(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupUserFilters(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupPassword(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupMtRand(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupCrypt(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupDir(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupSyslog(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if array.ZmStartupArray(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupAssert(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupUrlScannerEx(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupUserStreams(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupImagetypes(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	streams.PhpRegisterUrlStreamWrapper("php", &PhpStreamPhpWrapper)
	streams.PhpRegisterUrlStreamWrapper("file", &PhpPlainFilesWrapper)
	streams.PhpRegisterUrlStreamWrapper("glob", &streams.PhpGlobStreamWrapper)
	streams.PhpRegisterUrlStreamWrapper("data", &streams.PhpStreamRfc2397Wrapper)
	streams.PhpRegisterUrlStreamWrapper("http", &PhpStreamHttpWrapper)
	streams.PhpRegisterUrlStreamWrapper("ftp", &PhpStreamFtpWrapper)
	if ZmStartupDns(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupHrtime(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	return types.SUCCESS
}
func ZmShutdownBasic(type_ int, module_number int) int {
	ZmShutdownSyslog(type_, module_number)
	BasicGlobalsDtor(&BasicGlobals)
	streams.PhpUnregisterUrlStreamWrapper("php")
	streams.PhpUnregisterUrlStreamWrapper("http")
	streams.PhpUnregisterUrlStreamWrapper("ftp")
	ZmShutdownBrowscap(type_, module_number)
	//array.ZmShutdownArray(type_, module_number)
	ZmShutdownAssert(type_, module_number)
	ZmShutdownUrlScannerEx(type_, module_number)
	ZmShutdownFile(type_, module_number)
	ZmShutdownStandardFilters(type_, module_number)
	ZmShutdownCrypt(type_, module_number)
	ZmShutdownPassword(type_, module_number)
	return types.SUCCESS
}
func ZmActivateBasic(type_ int, module_number int) int {
	BG__().Activate()
	ZmActivateDir(type_, module_number)
	ZmActivateUrlScannerEx(type_, module_number)

	/* Setup default context */
	FG__().default_context = nil

	/* Default to global wrappers only */
	FG__().SetStreamWrappers(nil)

	/* Default to global filters only */
	FG__().SetStreamFilters(nil)
	return types.SUCCESS
}
func ZmDeactivateBasic(type_ int, module_number int) int {
	BG__().ResetRandGenerator()
	if BG__().umask != -1 {
		umask(BG__().umask)
	}

	/* Check if locale was changed and change it back
	 * to the value in startup environment */

	if BG__().localeChanged {
		setlocale(LC_ALL, "C")
		setlocale(LC_CTYPE, "")
		BG__().localeString = nil
	}

	/* FG__().stream_wrappers and FG__().stream_filters are destroyed
	 * during php_request_shutdown() */

	ZmDeactivateFilestat(type_, module_number)
	ZmDeactivateAssert(type_, module_number)
	ZmDeactivateUrlScannerEx(type_, module_number)
	streams.ZmDeactivateStreams(type_, module_number)
	ZmDeactivateUserFilters(type_, module_number)
	return types.SUCCESS
}
func ZmInfoBasic(zend_module *zend.ModuleEntry) {
	PhpInfoPrintTableStart()
	ZmInfoDl(zend_module)
	ZmInfoMail(zend_module)
	PhpInfoPrintTableEnd()
	ZmInfoAssert(zend_module)
}
func ZifConstant(returnValue zpp.Ret, constName string) {
	scope := zend.ZendGetExecutedScope()
	c := zend.ZendGetConstantEx(constName, scope, zend.ZEND_FETCH_CLASS_SILENT)
	if c != nil {
		types.ZVAL_COPY_OR_DUP(returnValue, c)
		if returnValue.IsType(types.IsConstantAst) {
			if zend.ZvalUpdateConstantEx(returnValue, scope) != types.SUCCESS {
				return
			}
		}
	} else {
		if zend.EG__().GetException() == nil {
			core.PhpErrorDocref("", faults.E_WARNING, "Couldn't find constant %s", constName)
		}
		returnValue.SetNull()
		return
	}
}

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
	//return_value.SetLong(ntohl(ip.s_addr))
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
func ZifGetenv(_ zpp.Opt, varname_ *string, localOnly bool) *types.Zval {
	if varname_ == nil {
		arr := core.DupEnvVariables()
		return types.NewZvalArray(arr)
	}

	env := core.Env__()

	var varName = b.Option(varname_, "")
	if localOnly {
		ptr := core.SapiGetenv(varName)
		if ptr != nil {
			return types.NewZvalString(*ptr)
		}
	}

	/* system method */
	if val, ok := env.LookupEnv(varName); ok {
		return types.NewZvalString(val)
	}
	return types.NewZvalFalse()
}
func ZifPutenv(setting string) bool {
	if setting == "" || setting[0] == '=' {
		core.PhpErrorDocref("", faults.E_WARNING, "Invalid parameter syntax")
		return false
	}

	env := core.Env__()

	// parse key
	key := setting
	if pos := strings.IndexByte(key, '='); pos >= 0 {
		key = key[:pos]
	}

	// 记录环境变量
	if key == setting {
		// setenv 格式为 {key}，删除环境变量
		env.UnSetEnv(setting)
	} else {
		// setenv 格式为 {key}={value}，设置环境变量新值
		env.PutEnv(setting)
	}

	// todo 特殊环境变量处理
	if strings.HasPrefix(key, "TZ") {
		tzset()
	}

	return true
}
func ParseOpts(opts *byte, result **core.Opt) int {
	var paras *core.Opt = nil
	var i uint
	var count uint = 0
	var opts_len uint = uint(strlen(opts))
	for i = 0; i < opts_len; i++ {
		if opts[i] >= 48 && opts[i] <= 57 || opts[i] >= 65 && opts[i] <= 90 || opts[i] >= 97 && opts[i] <= 122 {
			count++
		}
	}
	paras = zend.SafeEmalloc(b.SizeOf("opt_struct"), count, 0)
	memset(paras, 0, b.SizeOf("opt_struct")*count)
	*result = paras
	for (*opts) >= 48 && (*opts) <= 57 || (*opts) >= 65 && (*opts) <= 90 || (*opts) >= 97 && (*opts) <= 122 {
		paras.SetOptChar(*opts)
		paras.SetNeedParam((*(lang.PreInc(&opts))) == ':')
		paras.SetOptName(nil)
		if paras.GetNeedParam() == 1 {
			opts++
			if (*opts) == ':' {
				paras.GetNeedParam()++
				opts++
			}
		}
		paras++
	}
	return count
}
func ZifGetopt(executeData zpp.Ex, return_value zpp.Ret, shortOptions string, _ zpp.Opt, longOptions *types.Array, optind zpp.RefZval) {
	var options *byte = shortOptions
	var options_len int = len(shortOptions)
	var p_longopts *types.Zval = longOptions
	var zoptind *types.Zval = optind

	var argv **byte = nil
	var opt []byte = []byte{'0'}
	var optname *byte
	var argc int = 0
	var o int
	var len_ int
	var php_optarg *byte = nil
	var php_optind int = 1
	var val types.Zval
	var args *types.Zval = nil
	var optname_len int = 0
	var opts *core.Opt
	var orig_opts *core.Opt

	/* Init zoptind to 1 */
	if zoptind != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zoptind, 1)
	}

	/* Get argv from the global symbol table. We calculate argc ourselves
	 * in order to be on the safe side, even though it is also available
	 * from the symbol table. */
	if (core.PG__().http_globals[core.TRACK_VARS_SERVER].Type() == types.IsArray || zend.ZendIsAutoGlobal("_SERVER")) && (lang.Assign(&args, types.ZendHashFindInd(core.PG__().http_globals[core.TRACK_VARS_SERVER].Array(), types.STR_ARGV)) != nil || lang.Assign(&args, types.ZendHashFindInd(zend.EG__().GetSymbolTable(), types.STR_ARGV)) != nil) {
		var pos int = 0
		if !args.IsArray() {
			return_value.SetFalse()
			return
		}
		argc = args.Array().Len()

		/* Attempt to allocate enough memory to hold all of the arguments
		 * and a trailing NULL */

		argv = (**byte)(zend.SafeEmalloc(b.SizeOf("char *"), argc+1, 0))

		/* Iterate over the hash to construct the argv array. */
		args.Array().Foreach(func(key types.ArrayKey, value *types.Zval) {
			var arg_str *types.String = operators.ZvalGetString(value)
			argv[lang.PostInc(&pos)] = zend.Estrdup(arg_str.GetVal())
		})

		/* The C Standard requires argv[argc] to be NULL - this might
		 * keep some getopt implementations happy. */

		argv[argc] = nil

		/* The C Standard requires argv[argc] to be NULL - this might
		 * keep some getopt implementations happy. */

	} else {

		/* Return false if we can't find argv. */

		return_value.SetFalse()
		return
	}
	len_ = ParseOpts(options, &opts)
	if p_longopts != nil {
		var count int
		var entry *types.Zval
		count = p_longopts.Array().Len()

		/* the first <len> slots are filled by the one short ops
		 * we now extend our array and jump to the new added structs */

		opts = (*core.Opt)(zend.Erealloc(opts, b.SizeOf("opt_struct")*(len_+count+1)))
		orig_opts = opts
		opts += len_
		memset(opts, 0, count*b.SizeOf("opt_struct"))

		/* Iterate over the hash to construct the argv array. */
		p_longopts.Array().Foreach(func(_ types.ArrayKey, entry *types.Zval) {
			var arg_str *types.String = operators.ZvalGetString(entry)
			opts.SetNeedParam(0)
			opts.SetOptName(zend.Estrdup(arg_str.GetVal()))
			len_ = strlen(opts.GetOptName())
			if len_ > 0 && opts.GetOptName()[len_-1] == ':' {
				opts.GetNeedParam()++
				opts.GetOptName()[len_-1] = '0'
				if len_ > 1 && opts.GetOptName()[len_-2] == ':' {
					opts.GetNeedParam()++
					opts.GetOptName()[len_-2] = '0'
				}
			}
			opts.SetOptChar(0)
			opts++
		})

		/* Iterate over the hash to construct the argv array. */

	} else {
		opts = (*core.Opt)(zend.Erealloc(opts, b.SizeOf("opt_struct")*(len_+1)))
		orig_opts = opts
		opts += len_
	}

	/* php_getopt want to identify the last param */

	opts.SetOptChar('-')
	opts.SetNeedParam(0)
	opts.SetOptName(nil)

	/* Initialize the return value as an array. */

	zend.ArrayInit(return_value)

	/* after our pointer arithmetic jump back to the first element */

	opts = orig_opts
	for lang.Assign(&o, core.PhpGetopt(argc, argv, opts, &php_optarg, &php_optind, 0, 1)) != -1 {

		/* Skip unknown arguments. */

		if o == core.PHP_GETOPT_INVALID_ARG {
			continue
		}

		/* Prepare the option character and the argument string. */

		if o == 0 {
			optname = opts[core.PhpOptidx].GetOptName()
		} else {
			if o == 1 {
				o = '-'
			}
			opt[0] = o
			optname = opt
		}
		if php_optarg != nil {

			/* keep the arg as binary, since the encoding is not known */

			val.SetString(b.CastStrAuto(php_optarg))

		} else {
			val.SetFalse()
		}

		/* Add this option / argument pair to the result hash. */

		optname_len = strlen(optname)
		if !(optname_len > 1 && optname[0] == '0') && operators.IsNumericString(b.CastStr(optname, optname_len), nil, nil, 0) == types.IsLong {

			/* numeric string */

			var optname_int int = atoi(optname)
			if lang.Assign(&args, return_value.Array().IndexFind(optname_int)) != nil {
				if !args.IsArray() {
					operators.ConvertToArrayEx(args)
				}
				args.Array().Append(&val)
			} else {
				return_value.Array().IndexUpdate(optname_int, &val)
			}
		} else {

			/* other strings */

			if lang.Assign(&args, return_value.Array().KeyFind(b.CastStrAuto(optname))) != nil {
				if !args.IsArray() {
					operators.ConvertToArrayEx(args)
				}
				args.Array().Append(&val)
			} else {
				return_value.Array().KeyAdd(b.CastStrAuto(optname), &val)
			}

			/* other strings */

		}
		php_optarg = nil
	}

	/* Set zoptind to php_optind */

	if zoptind != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zoptind, php_optind)
	}
	zend.Efree(orig_opts)
}
func ZifFlush() {
	core.SapiFlush()
}
func ZifSleep(seconds int) (int, bool) {
	if seconds < 0 {
		core.PhpErrorDocref("", faults.E_WARNING, "Number of seconds must be greater than or equal to 0")
		return 0, false
	}
	rest := zend.Sleep(time.Duration(seconds) * time.Second)
	return int(rest.Seconds()), true
}
func ZifUsleep(microSeconds int) *types.Zval {
	if microSeconds < 0 {
		core.PhpErrorDocref("", faults.E_WARNING, "Number of microseconds must be greater than or equal to 0")
		return types.NewZvalFalse()
	}
	zend.Sleep(time.Duration(microSeconds) * time.Microsecond)
	return types.NewZvalNull()
}
func ZifTimeNanosleep(seconds int, nanoseconds int) *types.Zval {
	if seconds < 0 {
		core.PhpErrorDocref("", faults.E_WARNING, "The seconds value must be greater than 0")
		return types.NewZvalFalse()
	}
	if nanoseconds < 0 {
		core.PhpErrorDocref("", faults.E_WARNING, "The nanoseconds value must be greater than 0")
		return types.NewZvalFalse()
	}

	duration := time.Duration(seconds)*time.Second + time.Duration(nanoseconds)*time.Nanosecond
	rest := zend.Sleep(duration)
	if rest == 0 {
		return types.NewZvalTrue()
	} else {
		arr := types.NewArrayCap(2)
		arr.KeyAdd("seconds", types.NewZvalLong(int(rest.Seconds())))
		arr.KeyAdd("nanoseconds", types.NewZvalLong(int(rest.Nanoseconds())%int(time.Second)))
		return types.NewZvalArray(arr)
	}
}
func ZifTimeSleepUntil(timestamp float64) *types.Zval {
	targetTime := time.UnixMilli(int64(timestamp * float64(time.Second)))

	if targetTime.Before(time.Now()) {
		core.PhpErrorDocref("", faults.E_WARNING, "Sleep until to time is less than current time")
		return types.NewZvalFalse()
	}

	rest := zend.SleepUtil(targetTime)
	if rest == 0 {
		return types.NewZvalTrue()
	} else {
		return types.NewZvalFalse()
	}
}
func ZifGetCurrentUser() string {
	return zend.CurrEntrance().UserName()
}

func parseConfigArray(hash *types.Array) *types.Array {
	arr := types.NewArray()
	hash.Foreach(func(key types.ArrayKey, value *types.Zval) {
		if value.IsString() {
			if key.IsStrKey() {
				arr.SymtableUpdate(key.StrKey(), types.NewZvalString(value.String()))
			} else {
				arr.IndexUpdate(key.IdxKey(), types.NewZvalString(value.String()))
			}
		} else if value.IsArray() {
			tmp := parseConfigArray(value.Array())
			arr.KeyAdd(key.StrKey(), types.NewZvalArray(tmp))
		}
	})
	return arr
}

func ZifGetCfgVar(optionName string) *types.Zval {
	retval := core.CfgGetEntry(optionName)
	if retval != nil {
		if retval.IsType(types.IsArray) {
			arr := parseConfigArray(retval.Array())
			return types.NewZvalArray(arr)
		} else {
			return types.NewZvalString(retval.String())
		}
	} else {
		return types.NewZvalFalse()
	}
}
func ZifGetMagicQuotesRuntime(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetFalse()
	return
}
func ZifGetMagicQuotesGpc(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetFalse()
	return
}
func ZifErrorLog(message string, _ zpp.Opt, messageType int, destination *zpp.Path, extraHeaders *string) bool {
	var stream *core.PhpStream = nil
	var nbytes int
	switch messageType {
	case 1: /*send an email */
		if PhpMail(destination, "PHP error_log message", message, extraHeaders, nil) == 0 {
			return false
		}
	case 2: /*send to an address */
		core.PhpErrorDocref("", faults.E_WARNING, "TCP/IP option not available!")
		return false
	case 3: /*save to a file */
		stream = core.PhpStreamOpenWrapper(destination, "a", core.IGNORE_URL_WIN|core.REPORT_ERRORS, nil)
		if stream == nil {
			return false
		}
		nbytes = core.PhpStreamWriteString(stream, message)
		core.PhpStreamClose(stream)
		if nbytes != len(message) {
			return false
		}
	case 4: /* send to SAPI */
		if core.SM__().GetLogMessage() != nil {
			core.SM__().GetLogMessage()(message, -1)
		} else {
			return false
		}
	default:
		core.PhpLogErrWithSeverity(message, LOG_NOTICE)
	}
	return true
}
func ZifErrorGetLast() *types.Zval {
	lastError := core.PG__().LastError()
	if lastError != nil {
		arr := types.NewArrayCap(4)
		arr.KeyAdd("type", types.NewZvalLong(lastError.Type))
		arr.KeyAdd("message", types.NewZvalString(lastError.Message))
		if lastError.File != "" {
			arr.KeyAdd("file", types.NewZvalString(lastError.File))
		} else {
			arr.KeyAdd("file", types.NewZvalString("-"))
		}
		arr.KeyAdd("line", types.NewZvalLong(lastError.Lineno))
		return types.NewZvalArray(arr)
	}
	return types.NewZvalNull()
}
func ZifErrorClearLast() {
	core.PG__().ClearLastError()
}
func ZifCallUserFunc(callback zpp.Callable, _ zpp.Opt, args []*types.Zval) *types.Zval {
	retval, ok := callback.Call(args...)
	if ok && retval.IsNotUndef() {
		if retval.IsRef() {
			retval = retval.DeRef().Copy()
		}
		return retval
	}
	return types.NewZvalUndef()
}
func ZifCallUserFuncArray(callback zpp.Callable, parameters *types.Array) *types.Zval {
	args := parameters.Values()
	retval, ok := callback.Call(args...)
	if ok && retval.IsNotUndef() {
		if retval.IsRef() {
			retval = retval.DeRef().Copy()
		}
		return retval
	}
	return types.NewZvalUndef()
}
func ZifForwardStaticCall(executeData zpp.Ex, return_value zpp.Ret, functionName *types.Zval, _ zpp.Opt, parameters []*types.Zval) {
	var retval types.Zval
	var fci types.ZendFcallInfo
	var fci_cache types.ZendFcallInfoCache
	var called_scope *types.ClassEntry
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.ParseFunc(&fci, &fci_cache)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				fci.SetParams(_real_arg + 1)
				fci.SetParamCount(_num_varargs)
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				fci.SetParams(nil)
				fci.SetParamCount(0)
			}
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if !(executeData.GetPrevExecuteData().func_.common.scope) {
		faults.ThrowError(nil, "Cannot call forward_static_call() when no class scope is active")
		return
	}
	fci.SetRetval(&retval)
	called_scope = zend.ZendGetCalledScope(executeData)
	if called_scope != nil && fci_cache.GetCallingScope() != nil && operators.InstanceofFunction(called_scope, fci_cache.GetCallingScope()) != 0 {
		fci_cache.SetCalledScope(called_scope)
	}
	if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS && retval.IsNotUndef() {
		if retval.IsRef() {
			operators.ZendUnwrapReference(&retval)
		}
		types.ZVAL_COPY_VALUE(return_value, &retval)
	}
}
func ZifForwardStaticCallArray(executeData zpp.Ex, return_value zpp.Ret, functionName *types.Zval, parameters *types.Zval) {
	var params *types.Zval
	var retval types.Zval
	var fci types.ZendFcallInfo
	var fci_cache types.ZendFcallInfoCache
	var called_scope *types.ClassEntry
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			fp.ParseFunc(&fci, &fci_cache)
			params = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	zend.ZendFcallInfoArgs(&fci, params)
	fci.SetRetval(&retval)
	called_scope = zend.ZendGetCalledScope(executeData)
	if called_scope != nil && fci_cache.GetCallingScope() != nil && operators.InstanceofFunction(called_scope, fci_cache.GetCallingScope()) != 0 {
		fci_cache.SetCalledScope(called_scope)
	}
	if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS && retval.IsNotUndef() {
		if retval.IsRef() {
			operators.ZendUnwrapReference(&retval)
		}
		types.ZVAL_COPY_VALUE(return_value, &retval)
	}
	zend.ZendFcallInfoArgsClear(&fci, 1)
}
func PhpCallShutdownFunctions() {
	if BG__().HasUserShutdownFunctions() {
		faults.Try(func() {
			BG__().EachUserShutdownFunction(func(shutdownFunctionEntry *PhpShutdownFunction) {
				var retval types.Zval
				if !zend.IsCallable(shutdownFunctionEntry.Fn(), nil, 0) {
					var functionName = zend.GetCallableName(shutdownFunctionEntry.Fn(), nil)
					core.PhpError(faults.E_WARNING, "(Registered shutdown functions) Unable to call %s() - function does not exist", functionName)
				} else {
					zend.CallUserFunction_Ex(nil, shutdownFunctionEntry.Fn(), &retval, shutdownFunctionEntry.Args())
				}
			})
		})
	}
}
func PhpFreeShutdownFunctions() {
	BG__().ResetUserShutdownFunctions()
}
func ZifRegisterShutdownFunction(functionName *types.Zval, _ zpp.Opt, parameters []*types.Zval) {
	/* Prevent entering of anything but valid callback (syntax check only!) */
	if !zend.IsCallable(functionName, nil, 0) {
		var callbackName = zend.GetCallableName(functionName, nil)
		core.PhpErrorDocref("", faults.E_WARNING, "Invalid shutdown callback '%s' passed", callbackName)
		return
	}

	var fn = functionName.CopyValue()
	var args = make([]types.Zval, len(parameters))
	for i, parameter := range parameters {
		args[i].CopyValueFrom(parameter)
	}
	BG__().AddUserShutdownFunction(*NewShutdownFunction(fn, args))
}

func GetHighlightIni() *zend.ZendSyntaxHighlighterIni {
	var syntaxHighlighterIni zend.ZendSyntaxHighlighterIni
	syntaxHighlighterIni.SetHighlightComment(zend.INI_STR("highlight.comment"))
	syntaxHighlighterIni.SetHighlightDefault(zend.INI_STR("highlight.default"))
	syntaxHighlighterIni.SetHighlightHtml(zend.INI_STR("highlight.html"))
	syntaxHighlighterIni.SetHighlightKeyword(zend.INI_STR("highlight.keyword"))
	syntaxHighlighterIni.SetHighlightString(zend.INI_STR("highlight.string"))
	return &syntaxHighlighterIni
}

//@zif -alias show_source
func ZifHighlightFile(fileName zpp.Path, _ zpp.Opt, return_ bool) *types.Zval {
	if core.PhpCheckOpenBasedir(fileName) != 0 {
		return types.NewZvalFalse()
	}
	if return_ {
		core.PhpOutputStartDefault()
	}
	if ret := zend.HighlightFile(fileName, GetHighlightIni()); ret == types.FAILURE {
		if return_ {
			core.PhpOutputEnd()
		}
		return types.NewZvalFalse()
	}
	if return_ {
		contents, ok := core.OG__().GetContents()
		core.PhpOutputDiscard()

		if ok {
			return types.NewZvalString(contents)
		} else {
			return types.NewZvalNull()
		}
	} else {
		return types.NewZvalTrue()
	}
}
func ZifPhpStripWhitespace(fileName string) *types.Zval {
	var originalLexState zend.ZendLexState
	core.PhpOutputStartDefault()

	fh := zend.NewFileHandleByFilename(fileName)
	zend.ZendSaveLexicalState(&originalLexState)
	if zend.OpenFileForScanning(fh) == types.FAILURE {
		zend.ZendRestoreLexicalState(&originalLexState)
		core.PhpOutputEnd()
		return types.NewZvalString("")
	}
	zend.ZendStrip()
	zend.ZendDestroyFileHandle(fh)
	zend.ZendRestoreLexicalState(&originalLexState)
	ret := core.OG__().GetContentsZval()
	core.PhpOutputDiscard()
	return ret
}
func ZifHighlightString(string string, _ zpp.Opt, return_ bool) *types.Zval {
	var oldErrorReporting int = zend.EG__().GetErrorReporting()
	if return_ {
		core.PhpOutputStartDefault()
	}
	zend.EG__().SetErrorReporting(faults.E_ERROR)
	hicompiledStringDescription := zend.ZendMakeCompiledStringDescription("highlighted code")
	if zend.HighlightString(types.NewZvalString(string), GetHighlightIni(), hicompiledStringDescription) == types.FAILURE {
		zend.Efree(hicompiledStringDescription)
		zend.EG__().SetErrorReporting(oldErrorReporting)
		if return_ {
			core.PhpOutputEnd()
		}
		return types.NewZvalFalse()
	}
	zend.Efree(hicompiledStringDescription)
	zend.EG__().SetErrorReporting(oldErrorReporting)
	if return_ {
		ret := core.OG__().GetContentsZval()
		core.PhpOutputDiscard()
		return ret
	} else {
		return types.NewZvalTrue()
	}
}
func ZifIniGet(varname string) (string, bool) {
	return zend.ZendIniGetValueEx(varname)
}
func ZifIniGetAll(returnValue zpp.Ret, _ zpp.Opt, extension *string, details_ *bool) {
	var details bool = b.Option(details_, true)
	var moduleNumber int = 0
	zend.ZendIniSortEntries()
	if extension != nil {
		module := globals.G().GetModule(*extension)
		if module == nil {
			core.PhpErrorDocref("", faults.E_WARNING, "Unable to find extension '%s'", *extension)
			returnValue.SetFalse()
			return
		}
		moduleNumber = module.GetModuleNumber()
	}

	zend.ArrayInit(returnValue)
	zend.EG__().IniDirectives().Foreach(func(key string, iniEntry *zend.ZendIniEntry) {
		var option types.Zval
		if moduleNumber != 0 && iniEntry.GetModuleNumber() != moduleNumber {
			return
		}
		if key != "" {
			if details {
				zend.ArrayInit(&option)
				if iniEntry.GetOrigValue() != nil {
					zend.AddAssocStr(&option, "global_value", iniEntry.GetOrigValue().GetStr())
				} else if iniEntry.GetValue() != nil {
					zend.AddAssocStr(&option, "global_value", iniEntry.GetValue().GetStr())
				} else {
					zend.AddAssocNull(&option, "global_value")
				}
				if iniEntry.GetValue() != nil {
					zend.AddAssocStr(&option, "local_value", iniEntry.GetValue().GetStr())
				} else {
					zend.AddAssocNull(&option, "local_value")
				}
				zend.AddAssocLong(&option, "access", iniEntry.GetModifiable())
				returnValue.Array().SymtableUpdate(iniEntry.GetName().GetStr(), &option)
			} else {
				if iniEntry.GetValue() != nil {
					var zv types.Zval
					zv.SetString(iniEntry.GetValue().GetStr())
					returnValue.Array().SymtableUpdate(iniEntry.GetName().GetStr(), &zv)
				} else {
					returnValue.Array().SymtableUpdate(iniEntry.GetName().GetStr(), zend.UninitializedZval())
				}
			}
		}
	})
}
func PhpIniCheckPathEx(option string, newOption string) bool {
	if len(option)+1 != len(newOption) {
		return false
	}
	return option == newOption[:len(option)]
}
func PhpIniCheckPath(option_name *byte, option_len int, new_option_name string, new_option_len int) int {
	if option_len+1 != new_option_len {
		return 0
	}
	return !(strncmp(option_name, new_option_name, option_len))
}

//@zif -alias ini_alter
func ZifIniSet(return_value zpp.Ret, varname string, newvalue string) (string, bool) {
	val := zend.ZendIniGetValue(varname)

	/* open basedir check */
	if core.PG__().GetOpenBasedir() != nil {
		if varname == "error_log" || varname == "java.class.path" || varname == "java.home" || varname == "mail.log" || varname == "java.library.path" || varname == "vpopmail.directory" {
			if core.PhpCheckOpenBasedir(newvalue) != 0 {
				return_value.SetFalse()
				return "", false
			}
		}
	}
	if !zend.ZendAlterIniEntryEx(varname, newvalue, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0) {
		return_value.SetFalse()
		return "", false
	}

	if val != nil {
		return val.GetStr(), true
	}
	return "", false
}
func ZifIniRestore(varName string) {
	zend.ZendRestoreIniEntry(varName, core.PHP_INI_STAGE_RUNTIME)
}
func ZifSetIncludePath(newIncludePath zpp.Path) (string, bool) {
	oldValue := zend.ZendIniGetValue("include_path")

	if !zend.ZendAlterIniEntryEx("include_path", newIncludePath, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0) {
		return "", false
	}

	if oldValue != nil {
		return oldValue.GetStr(), true
	} else {
		return "", false
	}
}
func ZifGetIncludePath() (string, bool) {
	return zend.ZendIniGetValueEx("include_path")
}
func ZifRestoreIncludePath() {
	key := "include_path"
	zend.ZendRestoreIniEntry(key, core.PHP_INI_STAGE_RUNTIME)
}
func ZifPrintR(var_ *types.Zval, _ zpp.Opt, return_ bool) *types.Zval {
	if return_ {
		s := zend.ZendPrintZvalRToStr(var_, 0).GetStr()
		return types.NewZvalString(s)
	} else {
		zend.ZendPrintZvalR(var_, 0)
		return types.NewZvalTrue()
	}
}
func ZifConnectionAborted() int {
	return int(core.PG__().GetConnectionStatus()) & core.PHP_CONNECTION_ABORTED
}
func ZifConnectionStatus() int {
	return int(core.PG__().GetConnectionStatus())
}
func ZifIgnoreUserAbort(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, enable *bool) int {
	oldSetting := core.PG__().GetIgnoreUserAbort()
	if enable != nil {
		zend.ZendAlterIniEntryChars("ignore_user_abort", lang.Cond(*enable, "1", "0"), core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME)
	}
	return lang.Cond(oldSetting, 1, 0)
}
func ZifGetservbyname(executeData zpp.Ex, return_value zpp.Ret, service *types.Zval, protocol *types.Zval) {
	var name *byte
	var proto *byte
	var name_len int
	var proto_len int
	var serv *__struct__servent
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			name, name_len = fp.ParseString()
			proto, proto_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* empty string behaves like NULL on windows implementation of
	   getservbyname. Let be portable instead. */

	serv = getservbyname(name, proto)
	if serv == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(ntohs(serv.s_port))
	return
}
func ZifGetservbyport(executeData zpp.Ex, return_value zpp.Ret, port *types.Zval, protocol *types.Zval) {
	var proto *byte
	var proto_len int
	var port zend.ZendLong
	var serv *__struct__servent
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			port = fp.ParseLong()
			proto, proto_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	serv = getservbyport(htons(uint16(port)), proto)
	if serv == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetString(b.CastStrAuto(serv.s_name))
	return
}
func ZifGetprotobyname(executeData zpp.Ex, return_value zpp.Ret, name *types.Zval) {
	var name *byte
	var name_len int
	var ent *__struct__protoent
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			name, name_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	ent = getprotobyname(name)
	if ent == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(ent.p_proto)
	return
}
func ZifGetprotobynumber(executeData zpp.Ex, return_value zpp.Ret, proto *types.Zval) {
	var proto zend.ZendLong
	var ent *__struct__protoent
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			proto = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	ent = getprotobynumber(int(proto))
	if ent == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetString(b.CastStrAuto(ent.p_name))
	return
}
func ZifRegisterTickFunction(executeData zpp.Ex, return_value zpp.Ret, functionName *types.Zval, _ zpp.Opt, parameters []*types.Zval) {
	// todo 触发 warning
}
func ZifUnregisterTickFunction(executeData zpp.Ex, return_value zpp.Ret, functionName *types.Zval) {
	// todo 触发 warning
}
func ZifIsUploadedFile(path string) bool {
	return core.SG__().ExistUploadFile(path)
}
func ZifMoveUploadedFile(executeData zpp.Ex, return_value zpp.Ret, path string, newPath zpp.Path) bool {
	var successful bool = 0
	var oldmask int
	var ret int
	if !core.SG__().ExistUploadFile(path) {
		return false
	}
	if core.PhpCheckOpenBasedir(newPath) != 0 {
		return_value.SetFalse()
		return
	}
	if err := os.Rename(path, newPath); err == nil {
		successful = 1
		oldmask = umask(077)
		umask(oldmask)
		ret = zend.VCWD_CHMOD(newPath, 0666 & ^oldmask)
		if ret == -1 {
			core.PhpErrorDocref("", faults.E_WARNING, "%s", strerror(errno))
		}
	} else if PhpCopyFileEx(path, newPath, core.STREAM_DISABLE_OPEN_BASEDIR) == types.SUCCESS {
		zend.VCWD_UNLINK(path)
		successful = 1
	}
	if successful != 0 {
		core.SG__().DeleteUploadFile(path)
	} else {
		core.PhpErrorDocref("", faults.E_WARNING, "Unable to move '%s' to '%s'", path, new_path)
	}
	return successful != 0
}
func PhpSimpleIniParserCb(arg1 *types.Zval, arg2 *types.Zval, arg3 *types.Zval, callback_type int, arr *types.Zval) {
	switch callback_type {
	case zend.ZEND_INI_PARSER_ENTRY:
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}
		// arg2.TryAddRefcount()
		arr.Array().SymtableUpdate(arg1.StringEx().GetStr(), arg2)
	case zend.ZEND_INI_PARSER_POP_ENTRY:
		var hash types.Zval
		var find_hash *types.Zval
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}
		if !(arg1.StringEx().GetLen() > 1 && arg1.StringEx().GetStr()[0] == '0') && operators.IsNumericString(arg1.StringEx().GetStr(), nil, nil, 0) == types.IsLong {
			var key = zend.StrToLongWithUnit(arg1.String())
			if lang.Assign(&find_hash, arr.Array().IndexFind(key)) == nil {
				zend.ArrayInit(&hash)
				find_hash = arr.Array().IndexAddNew(key, &hash)
			}
		} else {
			if lang.Assign(&find_hash, arr.Array().KeyFind(arg1.StringEx().GetStr())) == nil {
				zend.ArrayInit(&hash)
				find_hash = arr.Array().KeyAddNew(arg1.StringEx().GetStr(), &hash)
			}
		}
		if !find_hash.IsArray() {
			// zend.ZvalPtrDtorNogc(find_hash)
			zend.ArrayInit(find_hash)
		}
		if arg3 == nil || arg3.IsString() && arg3.StringEx().GetLen() == 0 {
			// arg2.TryAddRefcount()
			zend.AddNextIndexZval(find_hash, arg2)
		} else {
			zend.ArraySetZvalKey(find_hash.Array(), arg3, arg2)
		}
	case zend.ZEND_INI_PARSER_SECTION:

	}
}
func PhpIniParserCbWithSections(arg1 *types.Zval, arg2 *types.Zval, arg3 *types.Zval, callback_type int, arr *types.Zval) {
	if callback_type == zend.ZEND_INI_PARSER_SECTION {
		zend.ArrayInit(&(BG__().active_ini_file_section))
		arr.Array().SymtableUpdate(arg1.StringEx().GetStr(), &(BG__().active_ini_file_section))
	} else if arg2 != nil {
		var active_arr *types.Zval
		if BG__().active_ini_file_section.IsNotUndef() {
			active_arr = &(BG__().active_ini_file_section)
		} else {
			active_arr = arr
		}
		PhpSimpleIniParserCb(arg1, arg2, arg3, callback_type, active_arr)
	}
}
func ZifParseIniFile(executeData zpp.Ex, return_value zpp.Ret, filename string, _ zpp.Opt, processSections_ *types.Zval, scannerMode *types.Zval) {
	var processSections bool = 0
	var scanner_mode zend.ZendLong = zend.ZEND_INI_SCANNER_NORMAL
	var ini_parser_cb zend.ZendIniParserCbT
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			fp.StartOptional()
			processSections = fp.ParseBool()
			scanner_mode = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if filename == "" {
		core.PhpErrorDocref("", faults.E_WARNING, "Filename cannot be empty!")
		return_value.SetFalse()
		return
	}

	/* Set callback function */

	if processSections != 0 {
		BG__().active_ini_file_section.SetUndef()
		ini_parser_cb = zend.ZendIniParserCbT(PhpIniParserCbWithSections)
	} else {
		ini_parser_cb = zend.ZendIniParserCbT(PhpSimpleIniParserCb)
	}

	/* Setup filehandle */

	fh := zend.NewFileHandleByFilename(filename)
	zend.ArrayInit(return_value)
	if zend.ZendParseIniFile(fh, 0, int(scanner_mode), ini_parser_cb, return_value) == types.FAILURE {
		return_value.Array().Destroy()
		return_value.SetFalse()
		return
	}
}
func ZifParseIniString(executeData zpp.Ex, return_value zpp.Ret, iniString *types.Zval, _ zpp.Opt, processSections *types.Zval, scannerMode *types.Zval) {
	var string *byte = nil
	var str *byte = nil
	var str_len int = 0
	var process_sections bool = 0
	var scanner_mode zend.ZendLong = zend.ZEND_INI_SCANNER_NORMAL
	var ini_parser_cb zend.ZendIniParserCbT
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			str, str_len = fp.ParseString()
			fp.StartOptional()
			process_sections = fp.ParseBool()
			scanner_mode = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if core.INT_MAX-str_len < zend.ZEND_MMAP_AHEAD {
		return_value.SetFalse()
	}

	/* Set callback function */

	if process_sections != 0 {
		BG__().active_ini_file_section.SetUndef()
		ini_parser_cb = zend.ZendIniParserCbT(PhpIniParserCbWithSections)
	} else {
		ini_parser_cb = zend.ZendIniParserCbT(PhpSimpleIniParserCb)
	}

	/* Setup string */

	string = (*byte)(zend.Emalloc(str_len + zend.ZEND_MMAP_AHEAD))
	memcpy(string, str, str_len)
	memset(string+str_len, 0, zend.ZEND_MMAP_AHEAD)
	zend.ArrayInit(return_value)
	if zend.ZendParseIniString(string, 0, int(scanner_mode), ini_parser_cb, return_value) == types.FAILURE {
		return_value.Array().Destroy()
		return_value.SetFalse()
	}
	zend.Efree(string)
}
func ZifSysGetloadavg(executeData zpp.Ex, return_value zpp.Ret) {
	var load []float64
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if getloadavg(load, 3) == -1 {
		return_value.SetFalse()
		return
	} else {
		zend.ArrayInit(return_value)
		zend.AddIndexDouble(return_value, 0, load[0])
		zend.AddIndexDouble(return_value, 1, load[1])
		zend.AddIndexDouble(return_value, 2, load[2])
	}
}
