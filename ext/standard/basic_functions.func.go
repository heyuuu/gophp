package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/sapi/cli"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/globals"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
	"math"
)

func BG__() *PhpBasicGlobals { return &BasicGlobals }
func PhpPutenvDestructor(zv *types.Zval) {
	var pe *PutenvEntry = zv.Ptr()
	if pe.GetPreviousValue() != nil {
		putenv(pe.GetPreviousValue())
	} else {
		unsetenv(pe.GetKey())
	}

	/* don't forget to reset the various libc globals that
	 * we might have changed by an earlier call to tzset(). */

	if !(strncmp(pe.GetKey(), "TZ", pe.GetKeyLen())) {
		tzset()
	}
	zend.Efree(pe.GetPutenvString())
	zend.Efree(pe.GetKey())
	zend.Efree(pe)
}
func BasicGlobalsCtor(basic_globals_p *PhpBasicGlobals) {
	BG__().mt_rand_is_seeded = 0
	BG__().mt_rand_mode = MT_RAND_MT19937
	BG__().umask = -1
	BG__().next = nil
	BG__().left = -1
	BG__().user_tick_functions = nil
	BG__().UserFilterMap = nil
	BG__().serialize_lock = 0
	memset(&(BG__().serialize), 0, b.SizeOf("BG ( serialize )"))
	memset(&(BG__().unserialize), 0, b.SizeOf("BG ( unserialize )"))
	memset(&(BG__().url_adapt_session_ex), 0, b.SizeOf("BG ( url_adapt_session_ex )"))
	memset(&(BG__().url_adapt_output_ex), 0, b.SizeOf("BG ( url_adapt_output_ex )"))
	BG__().url_adapt_session_ex.type_ = 1
	BG__().url_adapt_output_ex.type_ = 0
	BG__().url_adapt_session_hosts_ht.Init(0, nil)
	BG__().url_adapt_output_hosts_ht.Init(0, nil)
	BG__().incomplete_class = IncompleteClassEntry
	BG__().page_uid = -1
	BG__().page_gid = -1
}
func BasicGlobalsDtor(basic_globals_p *PhpBasicGlobals) {
	if basic_globals_p.GetUrlAdaptSessionEx().GetTags() != nil {
		basic_globals_p.GetUrlAdaptSessionEx().GetTags().Destroy()
		zend.Free(basic_globals_p.GetUrlAdaptSessionEx().GetTags())
	}
	if basic_globals_p.GetUrlAdaptOutputEx().GetTags() != nil {
		basic_globals_p.GetUrlAdaptOutputEx().GetTags().Destroy()
		zend.Free(basic_globals_p.GetUrlAdaptOutputEx().GetTags())
	}
	basic_globals_p.GetUrlAdaptSessionHostsHt().Destroy()
	basic_globals_p.GetUrlAdaptOutputHostsHt().Destroy()
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
	if zm_startup_nl_langinfo(type_, module_number) != types.SUCCESS {
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
	if ZmStartupArray(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupAssert(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupUrlScannerEx(type_, module_number) != types.SUCCESS {
		return types.FAILURE
	}
	if ZmStartupExec(type_, module_number) != types.SUCCESS {
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
	if ZmStartupRandom(type_, module_number) != types.SUCCESS {
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
	ZmShutdownArray(type_, module_number)
	ZmShutdownAssert(type_, module_number)
	ZmShutdownUrlScannerEx(type_, module_number)
	ZmShutdownFile(type_, module_number)
	ZmShutdownStandardFilters(type_, module_number)
	ZmShutdownCrypt(type_, module_number)
	ZmShutdownRandom(type_, module_number)
	ZmShutdownPassword(type_, module_number)
	return types.SUCCESS
}
func ZmActivateBasic(type_ int, module_number int) int {
	BG__().serialize_lock = 0
	memset(&(BG__().serialize), 0, b.SizeOf("BG ( serialize )"))
	memset(&(BG__().unserialize), 0, b.SizeOf("BG ( unserialize )"))
	BG__().locale_string = nil
	BG__().locale_changed = 0
	BG__().array_walk_fci = zend.EmptyFcallInfo
	BG__().array_walk_fci_cache = zend.EmptyFcallInfoCache
	BG__().user_compare_fci = zend.EmptyFcallInfo
	BG__().user_compare_fci_cache = zend.EmptyFcallInfoCache
	BG__().page_uid = -1
	BG__().page_gid = -1
	BG__().page_inode = -1
	BG__().page_mtime = -1
	BG__().putenv_ht.Init(1, PhpPutenvDestructor)
	BG__().user_shutdown_function_names = nil
	ZmActivateFilestat(type_, module_number)
	ZmActivateSyslog(type_, module_number)
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
	tsrm_env_lock()
	BG__().putenv_ht.Destroy()
	tsrm_env_unlock()
	BG__().mt_rand_is_seeded = 0
	if BG__().umask != -1 {
		umask(BG__().umask)
	}

	/* Check if locale was changed and change it back
	 * to the value in startup environment */

	if BG__().locale_changed {
		setlocale(LC_ALL, "C")
		setlocale(LC_CTYPE, "")
		if BG__().locale_string {
			// types.ZendStringReleaseEx(BG__().locale_string, 0)
			BG__().locale_string = nil
		}
	}

	/* FG__().stream_wrappers and FG__().stream_filters are destroyed
	 * during php_request_shutdown() */

	ZmDeactivateFilestat(type_, module_number)
	ZmDeactivateAssert(type_, module_number)
	ZmDeactivateUrlScannerEx(type_, module_number)
	streams.ZmDeactivateStreams(type_, module_number)
	if BG__().user_tick_functions {
		BG__().user_tick_functions.Destroy()
		zend.Efree(BG__().user_tick_functions)
		BG__().user_tick_functions = nil
	}
	ZmDeactivateUserFilters(type_, module_number)
	ZmDeactivateBrowscap(type_, module_number)
	BG__().page_uid = -1
	BG__().page_gid = -1
	return types.SUCCESS
}
func ZmInfoBasic(zend_module *zend.ModuleEntry) {
	PhpInfoPrintTableStart()
	ZmInfoDl(zend_module)
	ZmInfoMail(zend_module)
	PhpInfoPrintTableEnd()
	ZmInfoAssert(zend_module)
}
func ZifConstant(executeData zpp.Ex, return_value zpp.Ret, constName *types.Zval) {
	var const_name *types.String
	var c *types.Zval
	var scope *types.ClassEntry
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			const_name = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	scope = zend.ZendGetExecutedScope()
	c = zend.ZendGetConstantEx(const_name, scope, zend.ZEND_FETCH_CLASS_SILENT)
	if c != nil {
		types.ZVAL_COPY_OR_DUP(return_value, c)
		if return_value.IsType(types.IS_CONSTANT_AST) {
			if zend.ZvalUpdateConstantEx(return_value, scope) != types.SUCCESS {
				return
			}
		}
	} else {
		if zend.EG__().GetException() == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Couldn't find constant %s", const_name.GetVal())
		}
		return_value.SetNull()
		return
	}
}
func ZifInetNtop(executeData zpp.Ex, return_value zpp.Ret, inAddr *types.Zval) {
	var address *byte
	var address_len int
	var af int = AF_INET
	var buffer []byte
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			address, address_len = fp.ParseString()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if address_len == 16 {
		af = AF_INET6
	} else if address_len != 4 {
		return_value.SetFalse()
		return
	}
	if !(inet_ntop(af, address, buffer, b.SizeOf("buffer"))) {
		return_value.SetFalse()
		return
	}
	return_value.SetStringVal(b.CastStrAuto(buffer))
	return
}
func ZifInetPton(executeData zpp.Ex, return_value zpp.Ret, ipAddress string) {
	var ret int
	var af int = AF_INET
	var address *byte
	var address_len int
	var buffer []byte
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			address, address_len = fp.ParseString()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	memset(buffer, 0, b.SizeOf("buffer"))
	if strchr(address, ':') {
		af = AF_INET6
	} else if !(strchr(address, '.')) {
		return_value.SetFalse()
		return
	}
	ret = inet_pton(af, address, buffer)
	if ret <= 0 {
		return_value.SetFalse()
		return
	}
	return_value.SetStringVal(b.CastStr(buffer, b.Cond(af == AF_INET, 4, 16)))
	return
}
func ZifIp2long(executeData zpp.Ex, return_value zpp.Ret, ipAddress *types.Zval) {
	var addr *byte
	var addr_len int
	var ip __struct__in_addr
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
	if addr_len == 0 || inet_pton(AF_INET, addr, &ip) != 1 {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(ntohl(ip.s_addr))
	return
}
func ZifLong2ip(executeData zpp.Ex, return_value zpp.Ret, properAddress *types.Zval) {
	var ip zend.ZendUlong
	var sip zend.ZendLong
	var myaddr __struct__in_addr
	var str []byte
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			sip = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* autoboxes on 32bit platforms, but that's expected */

	ip = zend.ZendUlong(sip)
	myaddr.s_addr = htonl(ip)
	if inet_ntop(AF_INET, &myaddr, str, b.SizeOf("str")) {
		return_value.SetStringVal(b.CastStrAuto(str))
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifGetenv(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, varname *types.Zval, localOnly *types.Zval) {
	var ptr *byte
	var str *byte = nil
	var str_len int
	var local_only types.ZendBool = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 2, 0)
			fp.StartOptional()
			str, str_len = fp.ParseString()
			local_only = fp.ParseBool()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if str == nil {
		zend.ArrayInit(return_value)
		core.PhpImportEnvironmentVariables(return_value)
		return
	}
	if local_only == 0 {

		/* SAPI method returns an emalloc()'d string */

		ptr = core.SapiGetenv(b.CastStr(str, str_len))
		if ptr != nil {

			// TODO: avoid realocation ???

			return_value.SetStringVal(b.CastStrAuto(ptr))
			zend.Efree(ptr)
			return
		}
	}
	tsrm_env_lock()

	/* system method returns a const */

	ptr = getenv(str)
	if ptr != nil {
		return_value.SetStringVal(b.CastStrAuto(ptr))
	}
	tsrm_env_unlock()
	if ptr != nil {
		return
	}
	return_value.SetFalse()
	return
}
func ZifPutenv(executeData zpp.Ex, return_value zpp.Ret, setting *types.Zval) {
	var setting *byte
	var setting_len int
	var p *byte
	var env **byte
	var pe PutenvEntry
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			setting, setting_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if setting_len == 0 || setting[0] == '=' {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid parameter syntax")
		return_value.SetFalse()
		return
	}
	pe.SetPutenvString(zend.Estrndup(setting, setting_len))
	pe.SetKey(zend.Estrndup(setting, setting_len))
	if b.Assign(&p, strchr(pe.GetKey(), '=')) {
		*p = '0'
	}
	pe.SetKeyLen(strlen(pe.GetKey()))
	tsrm_env_lock()
	types.ZendHashStrDel(&(BG__().putenv_ht), b.CastStr(pe.GetKey(), pe.GetKeyLen()))

	/* find previous value */

	pe.SetPreviousValue(nil)
	for env = cli.Environ; env != nil && (*env) != nil; env++ {
		if !(strncmp(*env, pe.GetKey(), pe.GetKeyLen())) && (*env)[pe.GetKeyLen()] == '=' {
			pe.SetPreviousValue(*env)
			break
		}
	}
	if p == nil {
		unsetenv(pe.GetPutenvString())
	}
	if p == nil || putenv(pe.GetPutenvString()) == 0 {
		types.ZendHashAddMem(&(BG__().putenv_ht), b.CastStr(pe.GetKey(), pe.GetKeyLen()), &pe, b.SizeOf("putenv_entry"))
		if !(strncmp(pe.GetKey(), "TZ", pe.GetKeyLen())) {
			tzset()
		}
		tsrm_env_unlock()
		return_value.SetTrue()
		return
	} else {
		zend.Efree(pe.GetPutenvString())
		zend.Efree(pe.GetKey())
		return_value.SetFalse()
		return
	}
}
func FreeArgv(argv **byte, argc int) {
	var i int
	if argv != nil {
		for i = 0; i < argc; i++ {
			if argv[i] != nil {
				zend.Efree(argv[i])
			}
		}
		zend.Efree(argv)
	}
}
func FreeLongopts(longopts *core.Opt) {
	var p *core.Opt
	if longopts != nil {
		for p = longopts; p != nil && p.GetOptChar() != '-'; p++ {
			if p.GetOptName() != nil {
				zend.Efree((*byte)(p.GetOptName()))
			}
		}
	}
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
		paras.SetNeedParam((*(b.PreInc(&opts))) == ':')
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
func ZifGetopt(executeData zpp.Ex, return_value zpp.Ret, options *types.Zval, _ zpp.Opt, opts *types.Zval, optind zpp.RefZval) {
	var options *byte = nil
	var argv **byte = nil
	var opt []byte = []byte{'0'}
	var optname *byte
	var argc int = 0
	var o int
	var options_len int = 0
	var len_ int
	var php_optarg *byte = nil
	var php_optind int = 1
	var val types.Zval
	var args *types.Zval = nil
	var p_longopts *types.Zval = nil
	var zoptind *types.Zval = nil
	var optname_len int = 0
	var opts *core.Opt
	var orig_opts *core.Opt
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			options, options_len = fp.ParseString()
			fp.StartOptional()
			p_longopts = fp.ParseArray()
			zoptind = fp.ParseZval()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}

	/* Init zoptind to 1 */

	if zoptind != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zoptind, 1)
	}

	/* Get argv from the global symbol table. We calculate argc ourselves
	 * in order to be on the safe side, even though it is also available
	 * from the symbol table. */

	if (core.PG__().http_globals[core.TRACK_VARS_SERVER].GetType() == types.IS_ARRAY || zend.ZendIsAutoGlobalStr("_SERVER") != 0) && (b.Assign(&args, types.ZendHashFindInd(core.PG__().http_globals[core.TRACK_VARS_SERVER].Array(), types.STR_ARGV)) != nil || b.Assign(&args, types.ZendHashFindInd(zend.EG__().GetSymbolTable(), types.STR_ARGV)) != nil) {
		var pos int = 0
		var entry *types.Zval
		if args.GetType() != types.IS_ARRAY {
			return_value.SetFalse()
			return
		}
		argc = args.Array().Len()

		/* Attempt to allocate enough memory to hold all of the arguments
		 * and a trailing NULL */

		argv = (**byte)(zend.SafeEmalloc(b.SizeOf("char *"), argc+1, 0))

		/* Iterate over the hash to construct the argv array. */

		var __ht *types.Array = args.Array()
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			entry = _z
			var tmp_arg_str *types.String
			var arg_str *types.String = zend.ZvalGetTmpString(entry, &tmp_arg_str)
			argv[b.PostInc(&pos)] = zend.Estrdup(arg_str.GetVal())
			// zend.ZendTmpStringRelease(tmp_arg_str)
		}

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

		var __ht *types.Array = p_longopts.Array()
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			entry = _z
			var tmp_arg_str *types.String
			var arg_str *types.String = zend.ZvalGetTmpString(entry, &tmp_arg_str)
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
			// zend.ZendTmpStringRelease(tmp_arg_str)
		}

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
	for b.Assign(&o, core.PhpGetopt(argc, argv, opts, &php_optarg, &php_optind, 0, 1)) != -1 {

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

			val.SetStringVal(b.CastStrAuto(php_optarg))

		} else {
			val.SetFalse()
		}

		/* Add this option / argument pair to the result hash. */

		optname_len = strlen(optname)
		if !(optname_len > 1 && optname[0] == '0') && zend.IsNumericString(b.CastStr(optname, optname_len), nil, nil, 0) == types.IS_LONG {

			/* numeric string */

			var optname_int int = atoi(optname)
			if b.Assign(&args, return_value.Array().IndexFind(optname_int)) != nil {
				if args.GetType() != types.IS_ARRAY {
					zend.ConvertToArrayEx(args)
				}
				args.Array().NextIndexInsert(&val)
			} else {
				return_value.Array().IndexUpdate(optname_int, &val)
			}
		} else {

			/* other strings */

			if b.Assign(&args, return_value.Array().KeyFind(b.CastStrAuto(optname))) != nil {
				if args.GetType() != types.IS_ARRAY {
					zend.ConvertToArrayEx(args)
				}
				args.Array().NextIndexInsert(&val)
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
	FreeLongopts(orig_opts)
	zend.Efree(orig_opts)
	FreeArgv(argv, argc)
}
func ZifFlush(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	core.SapiFlush()
}
func ZifSleep(executeData zpp.Ex, return_value zpp.Ret, seconds *types.Zval) {
	var num zend.ZendLong
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			num = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if num < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Number of seconds must be greater than or equal to 0")
		return_value.SetFalse()
		return
	}
	return_value.SetLong(core.PhpSleep(uint(num)))
	return
}
func ZifUsleep(executeData zpp.Ex, return_value zpp.Ret, microSeconds *types.Zval) {
	var num zend.ZendLong
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			num = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if num < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Number of microseconds must be greater than or equal to 0")
		return_value.SetFalse()
		return
	}
	usleep(uint(num))
}
func ZifTimeNanosleep(executeData zpp.Ex, return_value zpp.Ret, seconds *types.Zval, nanoseconds *types.Zval) {
	var tv_sec zend.ZendLong
	var tv_nsec zend.ZendLong
	var php_req __struct__timespec
	var php_rem __struct__timespec
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			tv_sec = fp.ParseLong()
			tv_nsec = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if tv_sec < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The seconds value must be greater than 0")
		return_value.SetFalse()
		return
	}
	if tv_nsec < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The nanoseconds value must be greater than 0")
		return_value.SetFalse()
		return
	}
	php_req.tv_sec = int64(tv_sec)
	php_req.tv_nsec = long(tv_nsec)
	if !(nanosleep(&php_req, &php_rem)) {
		return_value.SetTrue()
		return
	} else if errno == EINTR {
		zend.ArrayInit(return_value)
		zend.AddAssocLongEx(return_value, "seconds", php_rem.tv_sec)
		zend.AddAssocLongEx(return_value, "nanoseconds", php_rem.tv_nsec)
		return
	} else if errno == EINVAL {
		core.PhpErrorDocref(nil, faults.E_WARNING, "nanoseconds was not in the range 0 to 999 999 999 or seconds was negative")
	}
	return_value.SetFalse()
	return
}
func ZifTimeSleepUntil(executeData zpp.Ex, return_value zpp.Ret, timestamp *types.Zval) {
	var target_secs float64
	var tm __struct__timeval
	var php_req __struct__timespec
	var php_rem __struct__timespec
	var current_ns uint64
	var target_ns uint64
	var diff_ns uint64
	var ns_per_sec uint64 = 1000000000
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			target_secs = fp.ParseDouble()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if gettimeofday((*__struct__timeval)(&tm), nil) != 0 {
		return_value.SetFalse()
		return
	}
	target_ns = uint64(target_secs * ns_per_sec)
	current_ns = uint64(tm.tv_sec)*ns_per_sec + uint64(tm.tv_usec)*1000
	if target_ns < current_ns {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Sleep until to time is less than current time")
		return_value.SetFalse()
		return
	}
	diff_ns = target_ns - current_ns
	php_req.tv_sec = time_t(diff_ns / ns_per_sec)
	php_req.tv_nsec = long(diff_ns % ns_per_sec)
	for nanosleep(&php_req, &php_rem) {
		if errno == EINTR {
			php_req.tv_sec = php_rem.tv_sec
			php_req.tv_nsec = php_rem.tv_nsec
		} else {
			return_value.SetFalse()
			return
		}
	}
	return_value.SetTrue()
	return
}
func ZifGetCurrentUser(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetStringVal(b.CastStrAuto(core.PhpGetCurrentUser()))
	return
}
func AddConfigEntry(h zend.ZendUlong, key *types.String, entry *types.Zval, retval *types.Zval) {
	if entry.IsType(types.IS_STRING) {
		var str = entry.String().Copy()
		if key != nil {
			zend.AddAssocStrEx(retval, key.GetStr(), str.GetStr())
		} else {
			zend.AddIndexStr(retval, h, str)
		}
	} else if entry.IsType(types.IS_ARRAY) {
		var tmp types.Zval
		zend.ArrayInit(&tmp)
		AddConfigEntries(entry.Array(), &tmp)
		retval.Array().KeyUpdate(key.GetStr(), &tmp)
	}
}
func AddConfigEntries(hash *types.Array, return_value *types.Zval) {
	var h zend.ZendUlong
	var key *types.String
	var zv *types.Zval
	var __ht *types.Array = hash
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		h = _p.GetH()
		key = _p.GetKey()
		zv = _z
		AddConfigEntry(h, key, zv, return_value)
	}
}
func ZifGetCfgVar(executeData zpp.Ex, return_value zpp.Ret, optionName *types.Zval) {
	var varname *byte
	var varname_len int
	var retval *types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			varname, varname_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	retval = core.CfgGetEntry(b.CastStr(varname, varname_len))
	if retval != nil {
		if retval.IsType(types.IS_ARRAY) {
			zend.ArrayInit(return_value)
			AddConfigEntries(retval.Array(), return_value)
			return
		} else {
			return_value.SetStringVal(b.CastStrAuto(retval.String().GetVal()))
			return
		}
	} else {
		return_value.SetFalse()
		return
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
func ZifErrorLog(executeData zpp.Ex, return_value zpp.Ret, message *types.Zval, _ zpp.Opt, messageType *types.Zval, destination *types.Zval, extraHeaders *types.Zval) {
	var message *byte
	var opt *byte = nil
	var headers *byte = nil
	var message_len int
	var opt_len int = 0
	var headers_len int = 0
	var opt_err int = 0
	var argc int = executeData.NumArgs()
	var erropt zend.ZendLong = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 4, 0)
			message, message_len = fp.ParseString()
			fp.StartOptional()
			erropt = fp.ParseLong()
			opt, opt_len = fp.ParsePath()
			headers, headers_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if argc > 1 {
		opt_err = int(erropt)
	}
	if _phpErrorLogEx(opt_err, message, message_len, opt, headers) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func _phpErrorLogEx(opt_err int, message *byte, message_len int, opt *byte, headers *byte) int {
	var stream *core.PhpStream = nil
	var nbytes int
	switch opt_err {
	case 1:
		if PhpMail(opt, "PHP error_log message", message, headers, nil) == 0 {
			return types.FAILURE
		}
	case 2:
		core.PhpErrorDocref(nil, faults.E_WARNING, "TCP/IP option not available!")
		return types.FAILURE
	case 3:
		stream = core.PhpStreamOpenWrapper(opt, "a", core.IGNORE_URL_WIN|core.REPORT_ERRORS, nil)
		if stream == nil {
			return types.FAILURE
		}
		nbytes = core.PhpStreamWrite(stream, message, message_len)
		core.PhpStreamClose(stream)
		if nbytes != message_len {
			return types.FAILURE
		}
	case 4:
		if core.SM__().GetLogMessage() != nil {
			core.SM__().GetLogMessage()(message, -1)
		} else {
			return types.FAILURE
		}
	default:
		core.PhpLogErrWithSeverity(message, LOG_NOTICE)
	}
	return types.SUCCESS
}
func ZifErrorGetLast(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if core.PG__().last_error_message {
		zend.ArrayInit(return_value)
		zend.AddAssocLongEx(return_value, "type", core.PG__().last_error_type)
		zend.AddAssocStr(return_value, "message", b.CastStrAuto(core.PG__().last_error_message))
		if core.PG__().last_error_file {
			zend.AddAssocStr(return_value, "file", b.CastStrAuto(core.PG__().last_error_file))
		} else {
			zend.AddAssocStr(return_value, "file", "-")
		}
		zend.AddAssocLongEx(return_value, "line", core.PG__().last_error_lineno)
	}
}
func ZifErrorClearLast(executeData zpp.Ex, return_value zpp.Ret) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if core.PG__().last_error_message {
		core.PG__().last_error_type = 0
		core.PG__().last_error_lineno = 0
		zend.Free(core.PG__().last_error_message)
		core.PG__().last_error_message = nil
		if core.PG__().last_error_file {
			zend.Free(core.PG__().last_error_file)
			core.PG__().last_error_file = nil
		}
	}
}
func ZifCallUserFunc(executeData zpp.Ex, return_value zpp.Ret, functionName *types.Zval, _ zpp.Opt, parameters []*types.Zval) {
	var retval types.Zval
	var fci types.ZendFcallInfo
	var fci_cache types.ZendFcallInfoCache
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.ParseFunc(&fci, &fci_cache)
			__arg, __arg_len := fp.ParseVariadic0()
			fci.SetParams(__arg)
			fci.SetParamCount(uint32(__arg_len))
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	fci.SetRetval(&retval)
	if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS && retval.IsNotUndef() {
		if retval.IsReference() {
			zend.ZendUnwrapReference(&retval)
		}
		types.ZVAL_COPY_VALUE(return_value, &retval)
	}
}
func ZifCallUserFuncArray(executeData zpp.Ex, return_value zpp.Ret, functionName *types.Zval, parameters *types.Zval) {
	var params *types.Zval
	var retval types.Zval
	var fci types.ZendFcallInfo
	var fci_cache types.ZendFcallInfoCache
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
	if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS && retval.IsNotUndef() {
		if retval.IsReference() {
			zend.ZendUnwrapReference(&retval)
		}
		types.ZVAL_COPY_VALUE(return_value, &retval)
	}
	zend.ZendFcallInfoArgsClear(&fci, 1)
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
	if called_scope != nil && fci_cache.GetCallingScope() != nil && zend.InstanceofFunction(called_scope, fci_cache.GetCallingScope()) != 0 {
		fci_cache.SetCalledScope(called_scope)
	}
	if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS && retval.IsNotUndef() {
		if retval.IsReference() {
			zend.ZendUnwrapReference(&retval)
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
	if called_scope != nil && fci_cache.GetCallingScope() != nil && zend.InstanceofFunction(called_scope, fci_cache.GetCallingScope()) != 0 {
		fci_cache.SetCalledScope(called_scope)
	}
	if zend.ZendCallFunction(&fci, &fci_cache) == types.SUCCESS && retval.IsNotUndef() {
		if retval.IsReference() {
			zend.ZendUnwrapReference(&retval)
		}
		types.ZVAL_COPY_VALUE(return_value, &retval)
	}
	zend.ZendFcallInfoArgsClear(&fci, 1)
}
func UserShutdownFunctionDtor(zv *types.Zval) {
	var i int
	var shutdown_function_entry *PhpShutdownFunctionEntry = zv.Ptr()
	for i = 0; i < shutdown_function_entry.GetArgCount(); i++ {
		// zend.ZvalPtrDtor(shutdown_function_entry.GetArguments()[i])
	}
	zend.Efree(shutdown_function_entry.GetArguments())
	zend.Efree(shutdown_function_entry)
}
func UserTickFunctionDtor(tick_function_entry *UserTickFunctionEntry) {
	var i int
	for i = 0; i < tick_function_entry.GetArgCount(); i++ {
		// zend.ZvalPtrDtor(tick_function_entry.GetArguments()[i])
	}
	zend.Efree(tick_function_entry.GetArguments())
}
func UserShutdownFunctionCall(zv *types.Zval) int {
	var shutdown_function_entry *PhpShutdownFunctionEntry = zv.Ptr()
	var retval types.Zval
	if zend.ZendIsCallable(shutdown_function_entry.GetArguments()[0], 0, nil) == 0 {
		var function_name *types.String = zend.ZendGetCallableName(shutdown_function_entry.GetArguments()[0])
		core.PhpError(faults.E_WARNING, "(Registered shutdown functions) Unable to call %s() - function does not exist", function_name.GetVal())
		// types.ZendStringReleaseEx(function_name, 0)
		return 0
	}
	if zend.CallUserFunction(nil, shutdown_function_entry.GetArguments()[0], &retval, shutdown_function_entry.GetArgCount()-1, shutdown_function_entry.GetArguments()+1) == types.SUCCESS {
		// zend.ZvalPtrDtor(&retval)
	}
	return 0
}
func UserTickFunctionCall(tick_fe *UserTickFunctionEntry) {
	var retval types.Zval
	var function *types.Zval = tick_fe.GetArguments()[0]

	/* Prevent reentrant calls to the same user ticks function */

	if tick_fe.GetCalling() == 0 {
		tick_fe.SetCalling(1)
		if zend.CallUserFunction(nil, function, &retval, tick_fe.GetArgCount()-1, tick_fe.GetArguments()+1) == types.SUCCESS {
			// zend.ZvalPtrDtor(&retval)
		} else {
			var obj *types.Zval
			var method *types.Zval
			if function.IsType(types.IS_STRING) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to call %s() - function does not exist", function.String().GetVal())
			} else if function.IsType(types.IS_ARRAY) && b.Assign(&obj, function.Array().IndexFind(0)) != nil && b.Assign(&method, function.Array().IndexFind(1)) != nil && obj.IsType(types.IS_OBJECT) && method.IsType(types.IS_STRING) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to call %s::%s() - function does not exist", types.Z_OBJCE_P(obj).GetName().GetVal(), method.String().GetVal())
			} else {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to call tick function")
			}
		}
		tick_fe.SetCalling(0)
	}

	/* Prevent reentrant calls to the same user ticks function */
}
func RunUserTickFunctions(tick_count int, arg any) {
	BG__().user_tick_functions.Apply(zend.LlistApplyFuncT(UserTickFunctionCall))
}
func UserTickFunctionCompare(tick_fe1 *UserTickFunctionEntry, tick_fe2 *UserTickFunctionEntry) int {
	var func1 *types.Zval = tick_fe1.GetArguments()[0]
	var func2 *types.Zval = tick_fe2.GetArguments()[0]
	var ret bool
	if func1.IsType(types.IS_STRING) && func2.IsType(types.IS_STRING) {
		ret = func1.StringVal() == func2.StringVal()
	} else if func1.IsType(types.IS_ARRAY) && func2.IsType(types.IS_ARRAY) {
		ret = zend.ZendCompareArrays(func1, func2) == 0
	} else if func1.IsType(types.IS_OBJECT) && func2.IsType(types.IS_OBJECT) {
		ret = zend.ZendCompareObjects(func1, func2) == 0
	} else {
		ret = false
	}
	if !ret && tick_fe1.GetCalling() != 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to delete tick function executed at the moment")
		return 0
	}
	return types.IntBool(ret)
}
func PhpCallShutdownFunctions() {
	if BG__().user_shutdown_function_names {
		faults.Try(func() {
			types.ZendHashApply(BG__().user_shutdown_function_names, UserShutdownFunctionCall)
		})
	}
}
func PhpFreeShutdownFunctions() {
	if BG__().user_shutdown_function_names {
		faults.TryCatch(func() {
			BG__().user_shutdown_function_names.Destroy()
			zend.FREE_HASHTABLE(BG__().user_shutdown_function_names)
			BG__().user_shutdown_function_names = nil
		}, func() {
			/* maybe shutdown method call exit, we just ignore it */
			zend.FREE_HASHTABLE(BG__().user_shutdown_function_names)
			BG__().user_shutdown_function_names = nil
		})
	}
}
func ZifRegisterShutdownFunction(executeData zpp.Ex, return_value zpp.Ret, functionName *types.Zval, _ zpp.Opt, parameters []*types.Zval) {
	var shutdown_function_entry PhpShutdownFunctionEntry
	var i int
	shutdown_function_entry.SetArgCount(executeData.NumArgs())
	if shutdown_function_entry.GetArgCount() < 1 {
		zend.ZendWrongParamCount()
		return
	}
	shutdown_function_entry.SetArguments((*types.Zval)(zend.SafeEmalloc(b.SizeOf("zval"), shutdown_function_entry.GetArgCount(), 0)))
	if zend.ZendGetParametersArray(executeData.NumArgs(), shutdown_function_entry.GetArgCount(), shutdown_function_entry.GetArguments()) == types.FAILURE {
		zend.Efree(shutdown_function_entry.GetArguments())
		return_value.SetFalse()
		return
	}

	/* Prevent entering of anything but valid callback (syntax check only!) */

	if zend.ZendIsCallable(shutdown_function_entry.GetArguments()[0], 0, nil) == 0 {
		var callback_name *types.String = zend.ZendGetCallableName(shutdown_function_entry.GetArguments()[0])
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid shutdown callback '%s' passed", callback_name.GetVal())
		zend.Efree(shutdown_function_entry.GetArguments())
		// types.ZendStringReleaseEx(callback_name, 0)
		return_value.SetFalse()
	} else {
		if !(BG__().user_shutdown_function_names) {
			zend.ALLOC_HASHTABLE(BG__().user_shutdown_function_names)
			BG__().user_shutdown_function_names.Init(0, UserShutdownFunctionDtor)
		}
		for i = 0; i < shutdown_function_entry.GetArgCount(); i++ {
			shutdown_function_entry.GetArguments()[i].TryAddRefcount()
		}
		types.ZendHashNextIndexInsertMem(BG__().user_shutdown_function_names, &shutdown_function_entry, b.SizeOf("php_shutdown_function_entry"))
	}

	/* Prevent entering of anything but valid callback (syntax check only!) */
}
func PhpGetHighlight(syntax_highlighter_ini *zend.ZendSyntaxHighlighterIni) {
	syntax_highlighter_ini.SetHighlightComment(zend.INI_STR("highlight.comment"))
	syntax_highlighter_ini.SetHighlightDefault(zend.INI_STR("highlight.default"))
	syntax_highlighter_ini.SetHighlightHtml(zend.INI_STR("highlight.html"))
	syntax_highlighter_ini.SetHighlightKeyword(zend.INI_STR("highlight.keyword"))
	syntax_highlighter_ini.SetHighlightString(zend.INI_STR("highlight.string"))
}

//@zif -alias show_source
func ZifHighlightFile(executeData zpp.Ex, return_value zpp.Ret, fileName *types.Zval, _ zpp.Opt, return_ *types.Zval) {
	var filename *byte
	var filename_len int
	var ret int
	var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
	var i types.ZendBool = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			filename, filename_len = fp.ParsePath()
			fp.StartOptional()
			i = fp.ParseBool()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if core.PhpCheckOpenBasedir(filename) != 0 {
		return_value.SetFalse()
		return
	}
	if i != 0 {
		core.PhpOutputStartDefault()
	}
	PhpGetHighlight(&syntax_highlighter_ini)
	ret = zend.HighlightFile(filename, &syntax_highlighter_ini)
	if ret == types.FAILURE {
		if i != 0 {
			core.PhpOutputEnd()
		}
		return_value.SetFalse()
		return
	}
	if i != 0 {
		core.PhpOutputGetContents(return_value)
		core.PhpOutputDiscard()
	} else {
		return_value.SetTrue()
		return
	}
}
func ZifPhpStripWhitespace(executeData zpp.Ex, return_value zpp.Ret, fileName *types.Zval) {
	var filename *byte
	var filename_len int
	var original_lex_state zend.ZendLexState
	var file_handle zend.ZendFileHandle
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			filename, filename_len = fp.ParsePath()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpOutputStartDefault()
	file_handle.InitFilename(filename)
	zend.ZendSaveLexicalState(&original_lex_state)
	if zend.OpenFileForScanning(&file_handle) == types.FAILURE {
		zend.ZendRestoreLexicalState(&original_lex_state)
		core.PhpOutputEnd()
		return_value.SetStringVal("")
		return
	}
	zend.ZendStrip()
	zend.ZendDestroyFileHandle(&file_handle)
	zend.ZendRestoreLexicalState(&original_lex_state)
	core.PhpOutputGetContents(return_value)
	core.PhpOutputDiscard()
}
func ZifHighlightString(executeData zpp.Ex, return_value zpp.Ret, string *types.Zval, _ zpp.Opt, return_ *types.Zval) {
	var expr *types.Zval
	var syntax_highlighter_ini zend.ZendSyntaxHighlighterIni
	var hicompiled_string_description *byte
	var i types.ZendBool = 0
	var old_error_reporting int = zend.EG__().GetErrorReporting()
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			expr = fp.ParseZval()
			fp.StartOptional()
			i = fp.ParseBool()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if zend.TryConvertToString(expr) == 0 {
		return
	}
	if i != 0 {
		core.PhpOutputStartDefault()
	}
	zend.EG__().SetErrorReporting(faults.E_ERROR)
	PhpGetHighlight(&syntax_highlighter_ini)
	hicompiled_string_description = zend.ZendMakeCompiledStringDescription("highlighted code")
	if zend.HighlightString(expr, &syntax_highlighter_ini, hicompiled_string_description) == types.FAILURE {
		zend.Efree(hicompiled_string_description)
		zend.EG__().SetErrorReporting(old_error_reporting)
		if i != 0 {
			core.PhpOutputEnd()
		}
		return_value.SetFalse()
		return
	}
	zend.Efree(hicompiled_string_description)
	zend.EG__().SetErrorReporting(old_error_reporting)
	if i != 0 {
		core.PhpOutputGetContents(return_value)
		core.PhpOutputDiscard()
	} else {
		return_value.SetTrue()
		return
	}
}
func ZifIniGet(varname string) (string, bool) {
	return zend.ZendIniGetValueEx(varname)
}
func ZifIniGetAll(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, extension *types.Zval, details *types.Zval) {
	var extname *byte = nil
	var extname_len int = 0
	var module_number int = 0
	var details types.ZendBool = 1
	var key *types.String
	var ini_entry *zend.ZendIniEntry
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 2, 0)
			fp.StartOptional()
			extname, extname_len = fp.ParseStringEx(true, false)
			details = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	zend.ZendIniSortEntries()
	if extname != nil {
		module := globals.G().GetModule(b.CastStr(extname, extname_len))
		if module == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to find extension '%s'", extname)
			return_value.SetFalse()
			return
		}
		module_number = module.GetModuleNumber()
	}
	zend.ArrayInit(return_value)

	zend.EG__().IniDirectives().Foreach(func(key string, ini_entry *zend.ZendIniEntry) {
		var option types.Zval
		if module_number != 0 && ini_entry.GetModuleNumber() != module_number {
			return
		}
		if key == nil || key != "" {
			if details != 0 {
				zend.ArrayInit(&option)
				if ini_entry.GetOrigValue() != nil {
					zend.AddAssocStr(&option, "global_value", ini_entry.GetOrigValue().GetStr())
				} else if ini_entry.GetValue() != nil {
					zend.AddAssocStr(&option, "global_value", ini_entry.GetValue().GetStr())
				} else {
					zend.AddAssocNull(&option, "global_value")
				}
				if ini_entry.GetValue() != nil {
					zend.AddAssocStr(&option, "local_value", ini_entry.GetValue().GetStr())
				} else {
					zend.AddAssocNull(&option, "local_value")
				}
				zend.AddAssocLong(&option, "access", ini_entry.GetModifiable())
				return_value.Array().SymtableUpdate(ini_entry.GetName().GetStr(), &option)
			} else {
				if ini_entry.GetValue() != nil {
					var zv types.Zval
					zv.SetStringCopy(ini_entry.GetValue())
					return_value.Array().SymtableUpdate(ini_entry.GetName().GetStr(), &zv)
				} else {
					return_value.Array().SymtableUpdate(ini_entry.GetName().GetStr(), zend.EG__().GetUninitializedZval())
				}
			}
		}
	})
}
func PhpIniCheckPath(option_name *byte, option_len int, new_option_name string, new_option_len int) int {
	if option_len+1 != new_option_len {
		return 0
	}
	return !(strncmp(option_name, new_option_name, option_len))
}

//@zif -alias ini_alter
func ZifIniSet(return_value zpp.Ret, varname string, newvalue string) {
	val := zend.ZendIniGetValue(varname)

	/* copy to return here, because alter might free it! */
	if val != nil {
		return_value.SetStringVal(val.GetStr())
	} else {
		return_value.SetFalse()
	}

	// #define _CHECK_PATH(var,var_len,ini) php_ini_check_path ( var , var_len , ini , sizeof ( ini ) )

	/* open basedir check */

	if core.PG__().open_basedir {
		if PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "error_log", b.SizeOf("\"error_log\"")) != 0 || PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "java.class.path", b.SizeOf("\"java.class.path\"")) != 0 || PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "java.home", b.SizeOf("\"java.home\"")) != 0 || PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "mail.log", b.SizeOf("\"mail.log\"")) != 0 || PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "java.library.path", b.SizeOf("\"java.library.path\"")) != 0 || PhpIniCheckPath(varname.GetVal(), varname.GetLen(), "vpopmail.directory", b.SizeOf("\"vpopmail.directory\"")) != 0 {
			if core.PhpCheckOpenBasedir(newvalue) != 0 {
				return_value.SetFalse()
				return
			}
		}
	}
	if !zend.ZendAlterIniEntryEx(varname.GetStr(), newvalue, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0) {
		return_value.SetFalse()
		return
	}
}
func ZifIniRestore(executeData zpp.Ex, return_value zpp.Ret, varname *types.Zval) {
	var varname *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			varname = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	zend.ZendRestoreIniEntry(varname, core.PHP_INI_STAGE_RUNTIME)
}
func ZifSetIncludePath(executeData zpp.Ex, return_value zpp.Ret, newIncludePath *types.Zval) {
	var new_value *types.String
	var old_value *byte
	var key *types.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			new_value = fp.ParsePathStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	old_value = zend.ZendIniString("include_path", 0)

	/* copy to return here, because alter might free it! */

	if old_value != nil {
		return_value.SetStringVal(b.CastStrAuto(old_value))
	} else {
		return_value.SetFalse()
	}
	key = types.NewString("include_path")
	if !zend.ZendAlterIniEntryEx(key.GetStr(), new_value, core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME, 0) {
		// types.ZendStringReleaseEx(key, 0)

		return_value.SetFalse()
		return
	}
	// types.ZendStringReleaseEx(key, 0)
}
func ZifGetIncludePath(executeData zpp.Ex, return_value zpp.Ret) {
	var str *byte
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	str = zend.ZendIniString("include_path", 0)
	if str == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetStringVal(b.CastStrAuto(str))
	return
}
func ZifRestoreIncludePath(executeData zpp.Ex, return_value zpp.Ret) {
	var key *types.String
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	key = types.NewString("include_path")
	zend.ZendRestoreIniEntry(key, core.PHP_INI_STAGE_RUNTIME)
	// types.ZendStringEfree(key)
}
func ZifPrintR(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval, _ zpp.Opt, return_ *types.Zval) {
	var var_ *types.Zval
	var do_return types.ZendBool = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			var_ = fp.ParseZval()
			fp.StartOptional()
			do_return = fp.ParseBool()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if do_return != 0 {
		return_value.SetString(zend.ZendPrintZvalRToStr(var_, 0))
		return
	} else {
		zend.ZendPrintZvalR(var_, 0)
		return_value.SetTrue()
		return
	}
}
func ZifConnectionAborted(executeData zpp.Ex, return_value zpp.Ret) {
	return_value.SetLong(core.PG__().connection_status & core.PHP_CONNECTION_ABORTED)
	return
}
func ZifConnectionStatus(executeData zpp.Ex, return_value zpp.Ret) {
	return_value.SetLong(core.PG__().connection_status)
	return
}
func ZifIgnoreUserAbort(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, value *types.Zval) {
	var arg types.ZendBool = 0
	var old_setting int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 1, 0)
			fp.StartOptional()
			arg = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	old_setting = uint16(core.PG__().ignore_user_abort)
	if executeData.NumArgs() != 0 {
		var key *types.String = types.NewString("ignore_user_abort")
		zend.ZendAlterIniEntryChars(key.GetStr(), b.CastStr(b.Cond(arg != 0, "1", "0"), 1), core.PHP_INI_USER, core.PHP_INI_STAGE_RUNTIME)
		// types.ZendStringReleaseEx(key, 0)
	}
	return_value.SetLong(old_setting)
	return
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
	return_value.SetStringVal(b.CastStrAuto(serv.s_name))
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
	return_value.SetStringVal(b.CastStrAuto(ent.p_name))
	return
}
func ZifRegisterTickFunction(executeData zpp.Ex, return_value zpp.Ret, functionName *types.Zval, _ zpp.Opt, parameters []*types.Zval) {
	var tick_fe UserTickFunctionEntry
	var i int
	var function_name *types.String = nil
	tick_fe.SetCalling(0)
	tick_fe.SetArgCount(executeData.NumArgs())
	if tick_fe.GetArgCount() < 1 {
		zend.ZendWrongParamCount()
		return
	}
	tick_fe.SetArguments((*types.Zval)(zend.SafeEmalloc(b.SizeOf("zval"), tick_fe.GetArgCount(), 0)))
	if zend.ZendGetParametersArray(executeData.NumArgs(), tick_fe.GetArgCount(), tick_fe.GetArguments()) == types.FAILURE {
		zend.Efree(tick_fe.GetArguments())
		return_value.SetFalse()
		return
	}
	if zend.ZendIsCallable(tick_fe.GetArguments()[0], 0, &function_name) == 0 {
		zend.Efree(tick_fe.GetArguments())
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid tick callback '%s' passed", function_name.GetVal())
		// types.ZendStringReleaseEx(function_name, 0)
		return_value.SetFalse()
		return
	} else if function_name != nil {
		// types.ZendStringReleaseEx(function_name, 0)
	}
	if tick_fe.GetArguments()[0].GetType() != types.IS_ARRAY && tick_fe.GetArguments()[0].GetType() != types.IS_OBJECT {
		zend.ConvertToStringEx(tick_fe.GetArguments()[0])
	}
	if !(BG__().user_tick_functions) {
		BG__().user_tick_functions = (*zend.ZendLlist)(zend.Emalloc(b.SizeOf("zend_llist")))
		BG__().user_tick_functions.Init(b.SizeOf("user_tick_function_entry"), zend.LlistDtorFuncT(UserTickFunctionDtor), 0)
		core.PhpAddTickFunction(RunUserTickFunctions, nil)
	}
	for i = 0; i < tick_fe.GetArgCount(); i++ {
		tick_fe.GetArguments()[i].TryAddRefcount()
	}
	BG__().user_tick_functions.AddElement(&tick_fe)
	return_value.SetTrue()
	return
}
func ZifUnregisterTickFunction(executeData zpp.Ex, return_value zpp.Ret, functionName *types.Zval) {
	var function *types.Zval
	var tick_fe UserTickFunctionEntry
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			function = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if !(BG__().user_tick_functions) {
		return
	}
	if function.GetType() != types.IS_ARRAY && function.GetType() != types.IS_OBJECT {
		zend.ConvertToString(function)
	}
	tick_fe.SetArguments((*types.Zval)(zend.Emalloc(b.SizeOf("zval"))))
	types.ZVAL_COPY_VALUE(tick_fe.GetArguments()[0], function)
	tick_fe.SetArgCount(1)
	zend.ZendLlistDelElement(BG__().user_tick_functions, &tick_fe, (func(any, any) int)(UserTickFunctionCompare))
	zend.Efree(tick_fe.GetArguments())
}
func ZifIsUploadedFile(executeData zpp.Ex, return_value zpp.Ret, path *types.Zval) {
	var path *byte
	var path_len int
	if !(core.SG__().rfc1867_uploaded_files) {
		return_value.SetFalse()
		return
	}
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			path, path_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if core.SG__().rfc1867_uploaded_files.KeyExists(b.CastStr(path, path_len)) {
		return_value.SetTrue()
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifMoveUploadedFile(executeData zpp.Ex, return_value zpp.Ret, path *types.Zval, newPath *types.Zval) {
	var path *byte
	var new_path *byte
	var path_len int
	var new_path_len int
	var successful types.ZendBool = 0
	var oldmask int
	var ret int
	if !(core.SG__().rfc1867_uploaded_files) {
		return_value.SetFalse()
		return
	}
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			path, path_len = fp.ParseString()
			new_path, new_path_len = fp.ParsePath()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if !core.SG__().rfc1867_uploaded_files.KeyExists(b.CastStr(path, path_len)) {
		return_value.SetFalse()
		return
	}
	if core.PhpCheckOpenBasedir(new_path) != 0 {
		return_value.SetFalse()
		return
	}
	if zend.VCWD_RENAME(path, new_path) == 0 {
		successful = 1
		oldmask = umask(077)
		umask(oldmask)
		ret = zend.VCWD_CHMOD(new_path, 0666 & ^oldmask)
		if ret == -1 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "%s", strerror(errno))
		}
	} else if PhpCopyFileEx(path, new_path, core.STREAM_DISABLE_OPEN_BASEDIR) == types.SUCCESS {
		zend.VCWD_UNLINK(path)
		successful = 1
	}
	if successful != 0 {
		types.ZendHashStrDel(core.SG__().rfc1867_uploaded_files, b.CastStr(path, path_len))
	} else {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to move '%s' to '%s'", path, new_path)
	}
	return_value.SetBool(successful != 0)
	return
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
		arr.Array().SymtableUpdate(arg1.String().GetStr(), arg2)
	case zend.ZEND_INI_PARSER_POP_ENTRY:
		var hash types.Zval
		var find_hash *types.Zval
		if arg2 == nil {

			/* bare string - nothing to do */

			break

			/* bare string - nothing to do */

		}
		if !(arg1.String().GetLen() > 1 && arg1.String().GetVal()[0] == '0') && zend.IsNumericString(arg1.String().GetStr(), nil, nil, 0) == types.IS_LONG {
			var key = zend.StrToLongWithUnit(arg1.StringVal())
			if b.Assign(&find_hash, arr.Array().IndexFind(key)) == nil {
				zend.ArrayInit(&hash)
				find_hash = arr.Array().IndexAddNew(key, &hash)
			}
		} else {
			if b.Assign(&find_hash, arr.Array().KeyFind(arg1.String().GetStr())) == nil {
				zend.ArrayInit(&hash)
				find_hash = arr.Array().KeyAddNew(arg1.String().GetStr(), &hash)
			}
		}
		if find_hash.GetType() != types.IS_ARRAY {
			// zend.ZvalPtrDtorNogc(find_hash)
			zend.ArrayInit(find_hash)
		}
		if arg3 == nil || arg3.IsType(types.IS_STRING) && arg3.String().GetLen() == 0 {
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
		arr.Array().SymtableUpdate(arg1.String().GetStr(), &(BG__().active_ini_file_section))
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
func ZifParseIniFile(executeData zpp.Ex, return_value zpp.Ret, filename *types.Zval, _ zpp.Opt, processSections *types.Zval, scannerMode *types.Zval) {
	var filename *byte = nil
	var filename_len int = 0
	var process_sections types.ZendBool = 0
	var scanner_mode zend.ZendLong = zend.ZEND_INI_SCANNER_NORMAL
	var fh zend.ZendFileHandle
	var ini_parser_cb zend.ZendIniParserCbT
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			filename, filename_len = fp.ParsePath()
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
	if filename_len == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Filename cannot be empty!")
		return_value.SetFalse()
		return
	}

	/* Set callback function */

	if process_sections != 0 {
		BG__().active_ini_file_section.SetUndef()
		ini_parser_cb = zend.ZendIniParserCbT(PhpIniParserCbWithSections)
	} else {
		ini_parser_cb = zend.ZendIniParserCbT(PhpSimpleIniParserCb)
	}

	/* Setup filehandle */

	fh.InitFilename(filename)
	zend.ArrayInit(return_value)
	if zend.ZendParseIniFile(&fh, 0, int(scanner_mode), ini_parser_cb, return_value) == types.FAILURE {
		return_value.Array().DestroyEx()
		return_value.SetFalse()
		return
	}
}
func ZifParseIniString(executeData zpp.Ex, return_value zpp.Ret, iniString *types.Zval, _ zpp.Opt, processSections *types.Zval, scannerMode *types.Zval) {
	var string *byte = nil
	var str *byte = nil
	var str_len int = 0
	var process_sections types.ZendBool = 0
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
		return_value.Array().DestroyEx()
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
